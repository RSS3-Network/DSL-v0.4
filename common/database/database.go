package database

import (
	"log"
	"os"
	"sync"
	"time"

	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/exchange"
	"github.com/naturalselectionlabs/pregod/common/database/model/governance"
	"github.com/naturalselectionlabs/pregod/common/database/model/social"
	"github.com/naturalselectionlabs/pregod/common/database/model/transaction"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var tables = []any{
	&model.Transaction{},
	&model.Transfer{},
	&model.GetTokenInfo{},
	&model.GetNFTTokenInfo{},
	&social.Profile{},
	&exchange.SwapPool{},
	&exchange.CexWallet{},
	&transaction.Token{},
	&transaction.CoinMarketCapCoinInfo{},
	&governance.SnapshotSpace{},
	&governance.SnapshotProposal{},
	&governance.SnapshotVote{},
	&model.Asset{},
	&model.Token{},
	&model.Domain{},
	&model.APIKey{},
	&model.Address{},
}

var (
	client       *gorm.DB
	globalLocker sync.RWMutex
)

func Global() *gorm.DB {
	globalLocker.RLock()

	defer globalLocker.RUnlock()

	return client
}

func ReplaceGlobal(db *gorm.DB) {
	globalLocker.Lock()

	defer globalLocker.Unlock()

	client = db
}

func Dial(dsn string, autoMigrate bool) error {
	var err error
	client, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
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
		return err
	}

	if autoMigrate {
		if err := client.AutoMigrate(tables...); err != nil {
			return nil, err
		}
	}

	return nil
}
