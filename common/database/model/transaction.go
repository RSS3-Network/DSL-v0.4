package model

import (
	"encoding/json"
	"time"
)

type Transaction struct {
	BlockNumber int64           `gorm:"column:block_number" json:"block_number"`
	Timestamp   time.Time       `gorm:"column:timestamp" json:"timestamp"`
	Hash        string          `gorm:"column:hash;primaryKey" json:"hash"`
	Index       int64           `gorm:"column:index;index;default:0" json:"index"`
	AddressFrom string          `gorm:"column:address_from" json:"address_from"`
	AddressTo   string          `gorm:"column:address_to" json:"address_to"`
	Network     string          `gorm:"column:network;primaryKey" json:"network"`
	Platform    string          `gorm:"column:platform" json:"platform"`
	Source      string          `gorm:"column:source;primaryKey" json:"-"`
	SourceData  json.RawMessage `gorm:"column:source_data;type:jsonb" json:"-"`
	CreatedAt   time.Time       `gorm:"column:created_at;autoCreateTime;not null;default:now();index" json:"-"`
	UpdatedAt   time.Time       `gorm:"column:updated_at;autoUpdateTime;not null;default:now();index" json:"-"`

	Transfers []*Transfer `gorm:"-:all" json:"actions"`
}
