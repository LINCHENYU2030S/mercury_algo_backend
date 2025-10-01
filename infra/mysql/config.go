package rds

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// DBConfig holds database connection parameters
type DBConfig struct {
	Host            string
	Port            int
	User            string
	Password        string
	DBName          string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
}

// LoadConfig loads configuration from environment variables
func LoadConfig() (*DBConfig, error) {
	dbPort, err := strconv.Atoi(getEnv("DB_PORT", "3306"))
	if err != nil {
		return nil, fmt.Errorf("invalid DB_PORT: %v", err)
	}

	maxIdleConns, err := strconv.Atoi(getEnv("DB_MAX_IDLE_CONNS", "10"))
	if err != nil {
		return nil, fmt.Errorf("invalid DB_MAX_IDLE_CONNS: %v", err)
	}

	maxOpenConns, err := strconv.Atoi(getEnv("DB_MAX_OPEN_CONNS", "100"))
	if err != nil {
		return nil, fmt.Errorf("invalid DB_MAX_OPEN_CONNS: %v", err)
	}

	connMaxLifetime, err := time.ParseDuration(getEnv("DB_CONN_MAX_LIFETIME", "1h"))
	if err != nil {
		return nil, fmt.Errorf("invalid DB_CONN_MAX_LIFETIME: %v", err)
	}

	return &DBConfig{
		Host:            getEnv("DB_HOST", "localhost"),
		Port:            dbPort,
		User:            getEnv("DB_USER", "root"),
		Password:        getEnv("DB_PASSWORD", ""),
		DBName:          getEnv("DB_NAME", "mercuryalgodb"),
		MaxIdleConns:    maxIdleConns,
		MaxOpenConns:    maxOpenConns,
		ConnMaxLifetime: connMaxLifetime,
	}, nil
}

// getEnv retrieves an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
