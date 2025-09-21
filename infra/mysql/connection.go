package mysql

import (
	"fmt"
	"os"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Config represents the connection parameters required to initialize a MySQL
// connection using GORM.
type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
	Params   string

	// Pool configuration options. A zero value keeps the default behaviour from
	// the underlying database driver.
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
}

// Open creates a new GORM connection using the provided configuration.
func Open(cfg Config) (*gorm.DB, error) {
	if cfg.Database == "" {
		return nil, fmt.Errorf("mysql database name is required")
	}

	dsn := buildDSN(cfg)
	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	}

	db, err := gorm.Open(mysql.Open(dsn), gormConfig)
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	if cfg.MaxIdleConns > 0 {
		sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	}
	if cfg.MaxOpenConns > 0 {
		sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	}
	if cfg.ConnMaxLifetime > 0 {
		sqlDB.SetConnMaxLifetime(cfg.ConnMaxLifetime)
	}

	return db, nil
}

// OpenFromEnv is a convenience helper that constructs the configuration from
// standard environment variables and then attempts to open the connection.
func OpenFromEnv() (*gorm.DB, error) {
	cfg := Config{
		Host:     os.Getenv("MYSQL_HOST"),
		Port:     os.Getenv("MYSQL_PORT"),
		Username: os.Getenv("MYSQL_USER"),
		Password: os.Getenv("MYSQL_PASSWORD"),
		Database: os.Getenv("MYSQL_DATABASE"),
		Params:   os.Getenv("MYSQL_PARAMS"),
	}

	if cfg.Host == "" {
		cfg.Host = "127.0.0.1"
	}
	if cfg.Port == "" {
		cfg.Port = "3306"
	}
	if cfg.Params == "" {
		cfg.Params = "parseTime=true&loc=Local&charset=utf8mb4"
	}

	return Open(cfg)
}

func buildDSN(cfg Config) string {
	host := cfg.Host
	if host == "" {
		host = "127.0.0.1"
	}

	port := cfg.Port
	if port == "" {
		port = "3306"
	}

	params := cfg.Params
	if strings.TrimSpace(params) == "" {
		params = "parseTime=true&loc=Local&charset=utf8mb4"
	}

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", cfg.Username, cfg.Password, host, port, cfg.Database, params)
}
