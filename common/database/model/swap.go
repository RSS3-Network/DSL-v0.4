package model

import (
	"time"
)

type Swap struct {
	ContractAddress string    `gorm:"column:contract_address;primaryKey"`
	Source          string    `gorm:"column:source"`
	Network         string    `gorm:"column:network;primaryKey"`
	Protocol        string    `gorm:"column:protocol"`
	CreatedAt       time.Time `gorm:"column:created_at;autoCreateTime;not null;default:now();index"`
	UpdatedAt       time.Time `gorm:"column:updated_at;autoUpdateTime;not null;default:now();index"`
}
