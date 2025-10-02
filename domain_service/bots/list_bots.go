package bots

import (
	"context"
	domain_models "mercury_algo_backend/domain/models"
	rds_models "mercury_algo_backend/infra/mysql/models"
	"mercury_algo_backend/kitex_gen/api"
	func_utils "mercury_algo_backend/utils/functor"
	"strings"
)

func ListBots(ctx context.Context, req *api.ListBotsRequest) ([]*domain_models.TradingBot, error) {
	//TODO
	var botsDAL []*rds_models.TradingBot
	var err error
	if req.TradingPair != nil {
		botsDAL, err = GetBotsByTradingPair(ctx, strings.TrimSpace(*(req.TradingPair)))
	} else {
		botsDAL, err = GetAllBots(ctx)
	}
	if err != nil {
		return nil, err
	}
	bots := func_utils.Map(botsDAL, domain_models.ConvertBotDALToDO)

	return bots, nil
}
