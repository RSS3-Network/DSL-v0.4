package model

import (
	"encoding/json"
	"time"

	"github.com/lib/pq"
)

type Profile struct {
	Address     string          `gorm:"column:address;primaryKey" json:"address"`
	Network     string          `gorm:"column:network;index" json:"network"`
	Platform    string          `gorm:"column:platform;index;primaryKey" json:"platform"`
	Source      string          `gorm:"column:source" json:"source"`
	Name        string          `gorm:"column:name" json:"name"`
	Handle      string          `gorm:"column:handle;primaryKey" json:"handle"`
	Bio         string          `gorm:"column:bio" json:"bio"`
	URL         string          `gorm:"column:url" json:"url"`
	ExpireAt    time.Time       `gorm:"column:expire_at" json:"expire_at"`
	ProfileUris pq.StringArray  `gorm:"column:profile_uri;type:text[]" json:"profile_uri"`
	BannerUris  pq.StringArray  `gorm:"column:banner_uri;type:text[]" json:"banner_uri"`
	SocialUris  pq.StringArray  `gorm:"column:social_uri;type:text[]" json:"social_uri"`
	SourceData  json.RawMessage `gorm:"column:source_data;type:jsonb" json:"-"`
	CreatedAt   time.Time       `gorm:"column:created_at;autoCreateTime;not null;default:now();index" json:"created_at"`
	UpdatedAt   time.Time       `gorm:"column:updated_at;autoUpdateTime;not null;default:now();index" json:"updated_at"`
}
