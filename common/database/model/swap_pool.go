package model

import (
	"time"
)

type SwapPool struct {
	Address   string    `gorm:"column:pool_address;primaryKey"`
	Name      string    `gorm:"column:name"`
	Network   string    `gorm:"column:network"`
	Protocol  string    `gorm:"column:protocol"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime;not null;default:now();index"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime;not null;default:now();index"`
}
