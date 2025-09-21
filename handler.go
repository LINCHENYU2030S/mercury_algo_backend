package handler

import (
	"encoding/json"
	"net/http"
	"time"
)

// HealthResponse represents the payload returned by the health check endpoint.
type HealthResponse struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
}

// Health exposes a simple readiness endpoint that can be used by load balancers
// or orchestration platforms to verify that the service is operational.
func Health() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := HealthResponse{
			Status:    "ok",
			Timestamp: time.Now().UTC(),
		}
		_ = json.NewEncoder(w).Encode(response)
	})
}
