package database

import (
	"sync"
	"time"

	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/exchange"
	"github.com/naturalselectionlabs/pregod/common/database/model/governance"
	"github.com/naturalselectionlabs/pregod/common/database/model/social"
	"github.com/naturalselectionlabs/pregod/common/database/model/transaction"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
	client               *gorm.DB
	globalLocker         sync.RWMutex
	globalDatabaseClient *gorm.DB
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

func ReplaceGlobal(db *gorm.DB) {
	globalLocker.Lock()

	defer globalLocker.Unlock()

	globalDatabaseClient = db
}

func Dial(dsn string, autoMigrate bool) (*gorm.DB, error) {
	var err error
	client, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		//Logger: logger.New(
		//	log.New(os.Stdout, "\r\n", log.LstdFlags),
		//	logger.Config{
		//		SlowThreshold:             time.Second * 3,
		//		LogLevel:                  logger.Error,
		//		IgnoreRecordNotFoundError: true,
		//	},
		//),
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
