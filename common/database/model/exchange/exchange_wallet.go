package exchange

import (
	"time"
)

type CexWallet struct {
	WalletAddress string    `gorm:"column:wallet_address;primaryKey"`
	Name          string    `gorm:"column:name"`
	Network       string    `gorm:"column:network;primaryKey"`
	Source        string    `gorm:"column:source"`
	CreatedAt     time.Time `gorm:"column:created_at;autoCreateTime;not null;default:now();index"`
	UpdatedAt     time.Time `gorm:"column:updated_at;autoUpdateTime;not null;default:now();index"`
}
