package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"mercury_algo_backend/handler"
	mysqlinfra "mercury_algo_backend/infra/mysql"
)

func main() {
	cfg := mysqlinfra.Config{
		Host:     getenv("MYSQL_HOST", "127.0.0.1"),
		Port:     getenv("MYSQL_PORT", "3306"),
		Username: getenv("MYSQL_USER", "root"),
		Password: os.Getenv("MYSQL_PASSWORD"),
		Database: os.Getenv("MYSQL_DATABASE"),
		Params:   getenv("MYSQL_PARAMS", "parseTime=true&loc=Local&charset=utf8mb4"),
	}

	db, err := mysqlinfra.Open(cfg)
	if err != nil {
		log.Fatalf("failed to connect to mysql: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to retrieve sql DB handle: %v", err)
	}
	defer sqlDB.Close()

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	mux := http.NewServeMux()
	mux.Handle("/healthz", handler.Health())

	server := &http.Server{
		Addr:         ":" + getenv("PORT", "8080"),
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	shutdownCtx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		log.Printf("http server listening on %s", server.Addr)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("http server terminated: %v", err)
		}
	}()

	<-shutdownCtx.Done()
	log.Println("shutdown signal received")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("graceful shutdown failed: %v", err)
	}
	log.Println("server stopped")
}

func getenv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
