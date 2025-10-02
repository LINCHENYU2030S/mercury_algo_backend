package application

import (
	"context"
	"mercury_algo_backend/domain/models"
	"mercury_algo_backend/domain_service/bots"
	"mercury_algo_backend/kitex_gen/api"
	api_utils "mercury_algo_backend/utils/api"
	func_utils "mercury_algo_backend/utils/functor"
)

func ListBots(ctx context.Context, req *api.ListBotsRequest) (*api.ListBotsResponse, error) {
	resp := &api.ListBotsResponse{}
	botsDO, err := bots.ListBots(ctx, req)
	if err == nil {
		resp.SetBots(func_utils.Map(botsDO, models.ConvertBotDOToApi))
	}
	resp.SetBase(api_utils.GenerateBaseResp(err))
	return resp, err
}
