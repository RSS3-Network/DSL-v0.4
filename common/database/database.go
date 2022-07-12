package database

import (
	"log"
	"os"
	"time"

	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/exchange"
	"github.com/naturalselectionlabs/pregod/common/database/model/governance"
	"github.com/naturalselectionlabs/pregod/common/database/model/transaction"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var tables = []any{
	&model.Transaction{},
	&model.Transfer{},
	&exchange.SwapPool{},
	&exchange.CexWallet{},
	&transaction.Token{},
	&model.GetTokenInfo{},
	&model.GetNFTTokenInfo{},
	&transaction.CoinMarketCapCoinInfo{},
	&governance.SnapshotSpace{},
	&governance.SnapshotProposal{},
	&governance.SnapshotVote{},
}

var Client *gorm.DB

func Dial(dsn string, autoMigrate bool) (*gorm.DB, error) {
	var err error
	Client, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
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
