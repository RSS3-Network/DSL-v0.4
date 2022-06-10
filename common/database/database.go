package database

import (
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var tables = []any{
	&model.Transaction{},
	&model.Transfer{},
	&model.SwapPool{},
	&model.ExchangeWallet{},
	&model.Token{},
	&model.GetTokenInfo{},
	&model.GetNFTTokenInfo{},
	&model.CoinMarketCapCoinInfo{},
}

var Client *gorm.DB

func Dial(dsn string, autoMigrate bool) (*gorm.DB, error) {
	var err error
	Client, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if autoMigrate {
		if err := Client.AutoMigrate(tables...); err != nil {
			return nil, err
		}
	}

	return Client, nil
}
