package model

import (
	"encoding/json"
	"time"

	"github.com/lib/pq"
	"github.com/shopspring/decimal"
)

type Transfer struct {
	TransactionHash     string          `gorm:"column:transaction_hash;primaryKey" json:"-"`
	Timestamp           time.Time       `gorm:"column:timestamp" json:"-"`
	Type                string          `gorm:"column:type" json:"type"`
	Tags                pq.StringArray  `gorm:"column:tags;type:text[]" json:"tags"`
	TransactionLogIndex decimal.Decimal `gorm:"column:transaction_log_index;primaryKey" json:"transaction_log_index"`
	AddressFrom         string          `gorm:"column:address_from" json:"address_from"`
	AddressTo           string          `gorm:"column:address_to" json:"address_to"`
	Metadata            json.RawMessage `gorm:"column:metadata;type:jsonb;default:'{}'" json:"metadata"`
	Network             string          `gorm:"column:network;primaryKey" json:"-"`
	Source              string          `gorm:"column:source" json:"-"`
	Platform            string          `gorm:"column:platform" json:"-"`
	SourceData          json.RawMessage `gorm:"column:source_data;type:jsonb" json:"-"`
	RelatedUrls         pq.StringArray  `gorm:"column:related_urls;type:text[]" json:"related_urls"`
	CreatedAt           time.Time       `gorm:"column:created_at;autoCreateTime;not null;default:now();index" json:"-"`
	UpdatedAt           time.Time       `gorm:"column:updated_at;autoUpdateTime;not null;default:now();index" json:"-"`
}
