package database

import (
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var tables = []any{
	&model.Transaction{},
	&model.Transfer{},
	&model.Swap{},
}

func Dial(dsn string, autoMigrate bool) (*gorm.DB, error) {
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if autoMigrate {
		if err := database.AutoMigrate(tables...); err != nil {
			return nil, err
		}
	}

	return database, nil
}
