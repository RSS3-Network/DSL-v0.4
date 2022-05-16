package model

import "github.com/shopspring/decimal"

type Transfer struct {
	TransactionHash     string          `gorm:"column:transaction_hash;primaryKey"`
	TransactionLogIndex int             `gorm:"column:transaction_log_index;primaryKey"`
	AddressFrom         string          `gorm:"column:address_from"`
	AddressTo           string          `gorm:"column:address_to"`
	TokenAddress        string          `gorm:"column:token_address"`
	TokenStandard       string          `gorm:"column:token_standard"`
	TokenValue          decimal.Decimal `gorm:"column:value"`
	TokenID             decimal.Decimal `gorm:"column:token_id"`
	Network             string          `gorm:"column:network;primaryKey"`
}
