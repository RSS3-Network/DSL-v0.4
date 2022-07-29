package model

import (
	"encoding/json"
	"time"

	"github.com/lib/pq"
)

type Asset struct {
	Network       string          `gorm:"column:network;primaryKey" json:"network"`
	TokenAddress  string          `gorm:"column:token_address;primaryKey" json:"token_address"`
	TokenID       string          `gorm:"column:token_id;primaryKey" json:"token_id"`
	TokenStandard string          `gorm:"column:token_standard" json:"token_standard"`
	Owner         string          `gorm:"column:owner;index" json:"owner"`
	Source        string          `gorm:"column:source" json:"-"`
	Title         string          `gorm:"column:title" json:"title"`
	Description   string          `gorm:"column:description" json:"description"`
	Attributes    json.RawMessage `gorm:"column:attributes;type:jsonb" json:"attributes,omitempty"`
	Image         string          `gorm:"column:image;type:text" json:"image"`
	RelatedUrls   pq.StringArray  `gorm:"column:related_urls;type:text[]" json:"related_urls"`
	Timestamp     time.Time       `gorm:"column:timestamp;index" json:"timestamp"`
	CreatedAt     time.Time       `gorm:"column:created_at;autoCreateTime;not null;default:now();index" json:"-"`
	UpdatedAt     time.Time       `gorm:"column:updated_at;autoUpdateTime;not null;default:now();index" json:"-"`
}

func (Asset) TableName() string {
	return "assets"
}
