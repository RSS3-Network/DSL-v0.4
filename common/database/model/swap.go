package model

import (
	"time"
)

type Swap struct {
	ContractAddress string    `gorm:"column:contract_address;primaryKey"`
	Name            string    `gorm:"column:name"`
	Source          string    `gorm:"column:source"`
	CreatedAt       time.Time `gorm:"column:created_at;autoCreateTime;not null;default:now();index"`
	UpdatedAt       time.Time `gorm:"column:updated_at;autoUpdateTime;not null;default:now();index"`
}
