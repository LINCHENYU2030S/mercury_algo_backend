package http

import (
	"log"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/cors"
)

// StartHTTPServer initializes and starts the HTTP server with automatic route mapping
func Init() {
	// Create a new Hertz server
	h := server.Default(server.WithHostPorts(":8888"))

	// Add CORS middleware
	h.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * 3600,
	}))

	// Register routes based on IDL annotations
	registerRoutes(h)

	// Start the server
	log.Println("Starting HTTP server on :8888")
	h.Spin()
}

func registerRoutes(h *server.Hertz) {
	// Register ListBots endpoint (using the IDL annotation path)
	// This directly maps to the endpoint defined in the Thrift IDL: (api.get="/api/v1/list_bots")
	h.GET("/api/v1/list_bots", ListBots)

	// You can add more endpoints here as needed
}
