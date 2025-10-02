package models

import (
	gormModels "mercury_algo_backend/infra/mysql/models"
)

// TradingBot represents the domain model for a trading bot
type TradingBot struct {
	ID                         int32
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
		ID:                         t.ID,
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
		ID:                         gormTradingBot.ID,
		Name:                       gormTradingBot.Name,
		TradingPair:                gormTradingBot.TradingPair,
		ArithmeticAnnualizedReturn: gormTradingBot.ArithmeticAnnualizedReturn,
		SharpeRatio:                gormTradingBot.SharpeRatio,
		MaximumDrawdown:            gormTradingBot.MaximumDrawdown,
		UserCount:                  gormTradingBot.UserCount,
	}
}

// NewTradingBot creates a new TradingBot domain model
func NewTradingBot(name, tradingPair string) *TradingBot {
	return &TradingBot{
		Name:        name,
		TradingPair: tradingPair,
	}
}
