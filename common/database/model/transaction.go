package model

import (
	"encoding/json"
	"time"

	"github.com/lib/pq"
)

type Transaction struct {
	BlockNumber int64           `gorm:"column:block_number" json:"-"`
	Timestamp   time.Time       `gorm:"column:timestamp" json:"timestamp"`
	Hash        string          `gorm:"column:hash;primaryKey" json:"hash"`
	Index       int64           `gorm:"column:index;index;default:0" json:"index"`
	Owner       string          `gorm:"column:owner;index;primaryKey" json:"owner"`
	AddressFrom string          `gorm:"column:address_from;index" json:"address_from"`
	AddressTo   string          `gorm:"column:address_to;index" json:"address_to,omitempty"`
	Addresses   pq.StringArray  `gorm:"column:addresses;type:text[];index" json:"addresses,omitempty"`
	Network     string          `gorm:"column:network;primaryKey" json:"network"`
	Platform    string          `gorm:"column:platform;index" json:"platform,omitempty"`
	Source      string          `gorm:"column:source;primaryKey" json:"-"`
	Tag         string          `gorm:"column:tag;index" json:"tag"`
	Type        string          `gorm:"column:type;index" json:"type"`
	Success     *bool           `gorm:"column:success;default:true" json:"success"`
	SourceData  json.RawMessage `gorm:"column:source_data;type:jsonb" json:"-"`
	CreatedAt   time.Time       `gorm:"column:created_at;autoCreateTime;not null;default:now();index" json:"-"`
	UpdatedAt   time.Time       `gorm:"column:updated_at;autoUpdateTime;not null;default:now();index" json:"-"`

	Transfers []Transfer `gorm:"-:all" json:"actions"`
}
