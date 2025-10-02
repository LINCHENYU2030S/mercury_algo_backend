package models

import (
	gormModels "mercury_algo_backend/infra/mysql/models"
	"mercury_algo_backend/kitex_gen/api"
	func_utils "mercury_algo_backend/utils/functor"
)

// TradingBot represents the domain model for a trading bot
type TradingBot struct {
	Name                       string
	TradingPair                string
	ArithmeticAnnualizedReturn *float32
	SharpeRatio                *float32
	MaximumDrawdown            *float32
	UserCount                  *int32
}

// ToGorm converts the domain model to a GORM model
func (t *TradingBot) ConvertBotDOToDAL() *gormModels.TradingBot {
	return &gormModels.TradingBot{
		Name:                       t.Name,
		TradingPair:                t.TradingPair,
		ArithmeticAnnualizedReturn: t.ArithmeticAnnualizedReturn,
		SharpeRatio:                t.SharpeRatio,
		MaximumDrawdown:            t.MaximumDrawdown,
		UserCount:                  t.UserCount,
	}
}

// FromGorm converts a GORM model to a domain model
func ConvertBotDALToDO(gormTradingBot *gormModels.TradingBot) *TradingBot {
	return &TradingBot{
		Name:                       gormTradingBot.Name,
		TradingPair:                gormTradingBot.TradingPair,
		ArithmeticAnnualizedReturn: gormTradingBot.ArithmeticAnnualizedReturn,
		SharpeRatio:                gormTradingBot.SharpeRatio,
		MaximumDrawdown:            gormTradingBot.MaximumDrawdown,
		UserCount:                  gormTradingBot.UserCount,
	}
}

func ConvertBotDOToApi(tradingBot *TradingBot) *api.TradingBot {
	return &api.TradingBot{
		Name:                       tradingBot.Name,
		TradingPair:                tradingBot.TradingPair,
		ArithmeticAnnualizedReturn: func_utils.Ptr(float64((*tradingBot.ArithmeticAnnualizedReturn))),
		SharpeRatio:                func_utils.Ptr(float64(*tradingBot.SharpeRatio)),
		MaximumDrawdown:            func_utils.Ptr(float64(*tradingBot.MaximumDrawdown)),
		UserCount:                  tradingBot.UserCount,
	}
}

// NewTradingBot creates a new TradingBot domain model
func NewTradingBot(name, tradingPair string) *TradingBot {
	return &TradingBot{
		Name:        name,
		TradingPair: tradingPair,
	}
}
