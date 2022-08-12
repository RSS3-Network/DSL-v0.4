package model

import "time"

const (
	AddressStatusDone    = "done"
	AddressStatusRunning = "running"
)

type Address struct {
	Address   string    `gorm:"column:address;primaryKey" json:"address"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime;not null;default:now();index" json:"-"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime;not null;default:now();index" json:"updated_at"`
}

func (Address) TableName() string {
	return "address"
}
