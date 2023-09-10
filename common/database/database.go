package database

import (
	"log"
	"os"
	"sync"
	"time"

	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/collectibe"
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
	&collectibe.FriendTech{},
}

var (
	client               *gorm.DB
	globalLocker         sync.RWMutex
	socialLocker         sync.RWMutex
	globalDatabaseClient *gorm.DB
	socialDatabaseClient *gorm.DB
	ethDbLocker          sync.RWMutex
	ethDbClient          *gorm.DB
)

func EthDb() *gorm.DB {
	ethDbLocker.RLock()

	defer ethDbLocker.RUnlock()

	return ethDbClient
}

func ReplaceEthDb(db *gorm.DB) {
	ethDbLocker.Lock()

	defer ethDbLocker.Unlock()

	ethDbClient = db
}

func Global() *gorm.DB {
	globalLocker.RLock()

	defer globalLocker.RUnlock()

	return globalDatabaseClient
}

func Social() *gorm.DB {
	socialLocker.RLock()

	defer socialLocker.RUnlock()

	return socialDatabaseClient
}

func ReplaceGlobal(db *gorm.DB) {
	globalLocker.Lock()

	defer globalLocker.Unlock()

	globalDatabaseClient = db
}

func ReplaceSocial(db *gorm.DB) {
	socialLocker.Lock()

	defer socialLocker.Unlock()

	socialDatabaseClient = db
}

func Dial(dsn string, autoMigrate bool) (*gorm.DB, error) {
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
		return nil, err
	}

	if autoMigrate {
		if err := client.AutoMigrate(tables...); err != nil {
			return nil, err
		}
	}
	sqlDB, _ := client.DB()

	sqlDB.SetMaxIdleConns(50)
	sqlDB.SetMaxOpenConns(300)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return client, nil
}
