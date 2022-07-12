package model

import (
	"time"

	"github.com/lib/pq"
)

type Asset struct {
	TokenAddress string         `gorm:"column:token_address;primaryKey" json:"token_address"`
	TokenID      string         `gorm:"column:token_id;primaryKey" json:"token_id"`
	Network      string         `gorm:"column:network;primaryKey" json:"network"`
	Owner        string         `gorm:"column:owner;index" json:"owner"`
	Platform     string         `gorm:"column:platform;index" json:"platform"`
	Source       string         `gorm:"column:source" json:"-"`
	RelatedURLs  pq.StringArray `gorm:"column:related_urls;type:text[]" json:"related_urls"`
	Title        string         `gorm:"column:title" json:"title"`
	Summary      string         `gorm:"column:summary" json:"summary"`
	Description  string         `gorm:"column:description" json:"description"`
	FileURL      string         `gorm:"column:file_url" json:"file_url"`
	CreatedAt    time.Time      `gorm:"column:created_at;autoCreateTime;not null;default:now();index" json:"-"`
	UpdatedAt    time.Time      `gorm:"column:updated_at;autoUpdateTime;not null;default:now();index" json:"-"`
}

func (Asset) TableName() string {
	return "assets"
}
