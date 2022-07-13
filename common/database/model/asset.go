package model

import (
	"time"
)

type Asset struct {
	Network       string    `gorm:"column:network;primaryKey" json:"network"`
	TokenAddress  string    `gorm:"column:token_address;primaryKey" json:"token_address"`
	TokenID       string    `gorm:"column:token_id;primaryKey" json:"token_id"`
	TokenStandard string    `gorm:"column:token_standard" json:"token_standard"`
	Owner         string    `gorm:"column:owner;index" json:"owner"`
	Source        string    `gorm:"column:source" json:"-"`
	Title         string    `gorm:"column:title" json:"title"`
	Description   string    `gorm:"column:description" json:"description"`
	FileURL       string    `gorm:"column:file_url" json:"file_url"`
	Timestamp     time.Time `gorm:"column:timestamp;index" json:"timestamp"`
	CreatedAt     time.Time `gorm:"column:created_at;autoCreateTime;not null;default:now();index" json:"-"`
	UpdatedAt     time.Time `gorm:"column:updated_at;autoUpdateTime;not null;default:now();index" json:"-"`
}

func (Asset) TableName() string {
	return "assets"
}
