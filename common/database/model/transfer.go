package model

import (
	"encoding/json"
	"time"

	"github.com/lib/pq"
	"github.com/shopspring/decimal"
)

type Transfer struct {
	TransactionHash     string          `gorm:"column:transaction_hash;primaryKey"`
	Timestamp           time.Time       `gorm:"column:timestamp"`
	Type                string          `gorm:"column:type"`
	Tags                pq.StringArray  `gorm:"column:tags;type:text[]"`
	TransactionLogIndex decimal.Decimal `gorm:"column:transaction_log_index;primaryKey"`
	AddressFrom         string          `gorm:"column:address_from"`
	AddressTo           string          `gorm:"column:address_to"`
	// Attachments         json.RawMessage `gorm:"column:attachments;type:jsonb;default:'[]'"`
	Metadata   json.RawMessage `gorm:"column:metadata;type:jsonb;default:'{}'"`
	Network    string          `gorm:"column:network;primaryKey"`
	Source     string          `gorm:"column:source"`
	SourceData json.RawMessage `gorm:"column:source_data;type:jsonb"`
	CreatedAt  time.Time       `gorm:"column:created_at;autoCreateTime;not null;default:now();index"`
	UpdatedAt  time.Time       `gorm:"column:updated_at;autoUpdateTime;not null;default:now();index"`
}
