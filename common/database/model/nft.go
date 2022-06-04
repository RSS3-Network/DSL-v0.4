package model

import (
	"encoding/json"
	"time"
)

type NFT struct {
	Name            string          `gorm:"column:name"`
	Symbol          string          `gorm:"column:symbol"`
	Amount          string          `gorm:"column:amount"`
	Standard        string          `gorm:"column:standard"`
	ContractAddress string          `gorm:"column:contract_address;primaryKey"`
	Owner           string          `gorm:"column:owner;primaryKey"`
	TokenId         int             `gorm:"column:token_id;primaryKey"`
	URI             string          `gorm:"column:uri"`
	Metadata        json.RawMessage `gorm:"column:metadata;type:jsonb"`
	Network         string          `gorm:"column:network"`
	Source          string          `gorm:"column:source"`
	SourceData      json.RawMessage `gorm:"column:source_data;type:jsonb"`
	CreatedAt       time.Time       `gorm:"column:created_at;autoCreateTime;not null;default:now();index"`
	UpdatedAt       time.Time       `gorm:"column:updated_at;autoUpdateTime;not null;default:now();index"`
}
