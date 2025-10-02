package bots

import (
	"context"
	"fmt"
	rds "mercury_algo_backend/infra/mysql"
	rds_models "mercury_algo_backend/infra/mysql/models"
	rds_query "mercury_algo_backend/infra/mysql/models/query"
)

func GetBotsByTradingPair(ctx context.Context, trading_pair string) ([]*rds_models.TradingBot, error) {
	db := rds.GetDB()
	if db == nil {
		return nil, fmt.Errorf("mysql connection is not initialized")
	}
	bots, err := rds_query.Use(db).TradingBot.WithContext(ctx).
		Where(rds_query.TradingBot.TradingPair.Eq(trading_pair)).
		Find()
	if err != nil {
		return nil, err
	}
	return bots, nil
}

func GetAllBots(ctx context.Context) ([]*rds_models.TradingBot, error) {
	db := rds.GetDB()
	if db == nil {
		return nil, fmt.Errorf("mysql connection is not initialized")
	}
	bots, err := rds_query.Use(db).TradingBot.Find()
	if err != nil {
		return nil, err
	}
	return bots, nil
}
