package exchange

import (
	"time"
)

type SwapPool struct {
	ContractAddress string    `gorm:"column:contract_address;primaryKey"`
	Source          string    `gorm:"column:source"`
	Network         string    `gorm:"column:network;primaryKey"`
	Token0          string    `gorm:"column:token0"`
	Token1          string    `gorm:"column:token1"`
	Protocol        string    `gorm:"column:protocol"`
	CreatedAt       time.Time `gorm:"column:created_at;autoCreateTime;not null;default:now();index"`
	UpdatedAt       time.Time `gorm:"column:updated_at;autoUpdateTime;not null;default:now();index"`
}
