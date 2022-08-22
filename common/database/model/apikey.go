package model

import "time"

type APIKey struct {
	Address string `gorm:"column:address;primaryKey" json:"address"`
	UUID    string `gorm:"column:uuid;index" json:"key"`
	Type    int    `gorm:"column:type;index;not null;default:1" json:"-"`
	Status  bool   `gorm:"column:status;not null;default:true" json:"-"`

	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime;not null;default:now();index" json:"-"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime;not null;default:now();index" json:"-"`
}

func (APIKey) TableName() string {
	return "apikey"
}
