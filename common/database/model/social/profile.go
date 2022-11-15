package social

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
	URL         string          `gorm:"column:url" json:"url,omitempty"`
	Action      string          `gorm:"column:action" json:"action,omitempty"`
	ExpireAt    *time.Time      `gorm:"column:expire_at;type:TIMESTAMP NULL" json:"expire_at,omitempty"`
	ProfileUris pq.StringArray  `gorm:"column:profile_uri;type:text[]" json:"profile_uri,omitempty"`
	BannerUris  pq.StringArray  `gorm:"column:banner_uri;type:text[]" json:"banner_uri,omitempty"`
	SocialUris  pq.StringArray  `gorm:"column:social_uri;type:text[]" json:"social_uri,omitempty"`
	SourceData  json.RawMessage `gorm:"column:source_data;type:jsonb" json:"-"`
	CreatedAt   time.Time       `gorm:"column:created_at;autoCreateTime;not null;default:now();index" json:"-"`
	UpdatedAt   time.Time       `gorm:"column:updated_at;autoUpdateTime;not null;default:now();index" json:"-"`
}
