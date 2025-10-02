package main

import (
	"log"
	rds "mercury_algo_backend/infra/mysql"
	api "mercury_algo_backend/kitex_gen/api/mercuryalgobackendservice"
)

func main() {
	svr := api.NewServer(new(MercuryAlgoBackendServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}

	// Initialize database connection
	_, err = rds.Initialize()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer func() {
		if err = rds.Close(); err != nil {
			log.Printf("Failed to close database connection: %v", err)
		}
	}()
}
