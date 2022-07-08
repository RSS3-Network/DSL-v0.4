package database

import (
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/exchange"
	"github.com/naturalselectionlabs/pregod/common/database/model/governance"
	"github.com/naturalselectionlabs/pregod/common/database/model/transaction"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
