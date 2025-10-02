package main

import (
	"log"

	rds "mercury_algo_backend/infra/mysql"
)

func main() {
	// Initialize database connection
	_, err := rds.Initialize()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer func() {
		err = rds.Close()
		if err != nil {
			log.Fatalf("Failed to close database connection")
		}
	}()
	log.Println("Starting Mercury Algo backend service...")

	// Set up and start your API server here
	// Example: server := setupServer()
	// server.Start()
}
