package database

import (
	"log"
	"os"
	"time"

	"github.com/naturalselectionlabs/pregod/common/database/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var tables = []any{
	&model.Transaction{},
	&model.Transfer{},
	&model.SwapPool{},
	&model.CexWallet{},
	&model.Token{},
	&model.GetTokenInfo{},
	&model.GetNFTTokenInfo{},
	&model.CoinMarketCapCoinInfo{},
	&model.SnapshotSpace{},
	&model.SnapshotProposal{},
	&model.SnapshotVote{},
}

var Client *gorm.DB

func Dial(dsn string, autoMigrate bool) (*gorm.DB, error) {
	var err error
	Client, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Second * 3,
				LogLevel:                  logger.Error,
				IgnoreRecordNotFoundError: true,
			},
		),
	})
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
