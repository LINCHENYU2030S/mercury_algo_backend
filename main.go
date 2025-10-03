package main

import (
	"log"
	rds "mercury_algo_backend/infra/mysql"
	rds_query "mercury_algo_backend/infra/mysql/models/query"
	api "mercury_algo_backend/kitex_gen/api/mercuryalgobackendservice"
	"mercury_algo_backend/server/http"
	"os"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/utils"
	"github.com/cloudwego/kitex/server"
)

func main() {
	// Initialize database connection
	db, err := rds.Initialize()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer func() {
		if err = rds.Close(); err != nil {
			log.Printf("Failed to close database connection: %v", err)
		}
	}()
	rds_query.SetDefault(db)

	go http.Init()

	addr := os.Getenv("KITEX_ADDR")
	if addr == "" {
		addr = ":9000"
	}
	svr := api.NewServer(
		new(MercuryAlgoBackendServiceImpl),
		server.WithServiceAddr(utils.NewNetAddr("tcp", addr)),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "mercury_algo"}),
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
