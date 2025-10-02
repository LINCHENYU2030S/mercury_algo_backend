package main

import (
	"context"
	"mercury_algo_backend/domain_service/bots"
	"mercury_algo_backend/kitex_gen/api"
)

type MercuryAlgoBackendServiceImpl struct{}

// ListAllBots implements the MercuryAlgoBackendServiceImpl interface.
func (s *MercuryAlgoBackendServiceImpl) ListBots(ctx context.Context, req *api.ListBotsRequest) (resp *api.ListBotsResponse, err error) {
	resp, err := bots.
}
