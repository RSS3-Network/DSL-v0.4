package model

import "time"

type Domains struct {
	BlockTimestamp      time.Time `gorm:"column:block_timestamp" json:"block_timestamp"`
	TransactionHash     []byte    `gorm:"column:transaction_hash;type:bytea;index" json:"transaction_hash"`
	TransactionLogIndex int       `gorm:"column:transaction_log_index;type:bigint" json:"transaction_log_index"`
	Type                string    `gorm:"column:type;type:text" json:"type"`
	Name                string    `gorm:"column:name;type:text;index" json:"name"`
	AddressOwner        []byte    `gorm:"column:address_owner;type:bytea;" json:"address_owner"`
	AddressTarget       []byte    `gorm:"column:address_target;type:bytea" json:"address_target"`
	ExpiredAt           time.Time `gorm:"column:expired_at"`
	Source              string    `gorm:"column:source;type:text;index" json:"source"`
	CreatedAt           time.Time `gorm:"column:created_at;autoCreateTime;not null;default:now();index" json:"created_at"`
	UpdatedAt           time.Time `gorm:"column:updated_at;autoUpdateTime;not null;default:now();index" json:"updated_at"`
}

func (Domains) TableName() string {
	return "domains"
}
