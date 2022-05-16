package model

import "github.com/shopspring/decimal"

type Transaction struct {
	BlockNumber decimal.Decimal `gorm:"column:block_number"`
	Hash        string          `gorm:"column:hash"`
	AddressFrom string          `gorm:"column:address_from"`
	AddressTo   string          `gorm:"column:address_to"`
	Network     string          `gorm:"column:network"`
}
