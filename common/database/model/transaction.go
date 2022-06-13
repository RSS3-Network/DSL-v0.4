package model

import (
	"encoding/json"
	"time"

	"github.com/shopspring/decimal"
)

type Transaction struct {
	BlockNumber decimal.Decimal `gorm:"column:block_number"`
	Timestamp   time.Time       `gorm:"column:timestamp"`
	Hash        string          `gorm:"column:hash"`
	AddressFrom string          `gorm:"column:address_from"`
	AddressTo   string          `gorm:"column:address_to"`
	Network     string          `gorm:"column:network"`
	Source      string          `gorm:"column:source"`
	SourceData  json.RawMessage `gorm:"column:source_data;type:jsonb"`
	CreatedAt   time.Time       `gorm:"column:created_at;autoCreateTime;not null;default:now();index"`
	UpdatedAt   time.Time       `gorm:"column:updated_at;autoUpdateTime;not null;default:now();index"`

	Transfers []Transfer `gorm:"-:all"`
}
