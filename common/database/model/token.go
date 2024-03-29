package model

import (
	"time"
)

type Token struct {
	ID              string    `gorm:"column:id;not null;primaryKey" json:"-"`
	Name            string    `gorm:"column:name;not null" json:"name"`
	Symbol          string    `gorm:"column:symbol;not null" json:"symbol"`
	Logo            string    `gorm:"column:logo" json:"logo"`
	Standard        string    `gorm:"column:standard" json:"standard"`
	Network         string    `gorm:"column:network;not null;primaryKey" json:"network"`
	Decimal         uint8     `gorm:"column:decimal;not null" json:"decimal"`
	ContractAddress string    `gorm:"column:contract_address;not null" json:"contract_address"`
	CreatedAt       time.Time `gorm:"column:created_at;autoCreateTime;not null;default:now();index" json:"-"`
	UpdatedAt       time.Time `gorm:"column:updated_at;autoUpdateTime;not null;default:now();index" json:"-"`
}
