package model

import (
	"math/big"
	"time"
)

type Token struct {
	Name            string    `gorm:"column:name;not null" json:"name"`
	Symbol          string    `gorm:"column:symbol;not null" json:"symbol"`
	ChainID         *big.Int  `gorm:"column:network;not null;primaryKey" json:"network"`
	Decimal         uint8     `gorm:"column:decimal;not null;" json:"decimal"`
	ContractAddress string    `gorm:"column:contract_address;not null;primaryKey" json:"contract_address"`
	CreatedAt       time.Time `gorm:"column:created_at;autoCreateTime;not null;default:now();index" json:"-"`
	UpdatedAt       time.Time `gorm:"column:updated_at;autoUpdateTime;not null;default:now();index" json:"-"`
}
