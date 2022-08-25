package model

import (
	"time"

	"github.com/lib/pq"
)

type Address struct {
	Address          string         `gorm:"column:address;primaryKey" json:"address"`
	Status           bool           `gorm:"column:status;default:true" json:"status"`
	DoneNetworks     pq.StringArray `gorm:"column:done_networks;type:text[]" json:"done_networks"`
	IndexingNetworks pq.StringArray `gorm:"column:indexing_networks;type:text[]" json:"indexing_networks"`
	CreatedAt        time.Time      `gorm:"column:created_at;autoCreateTime;not null;default:now();index" json:"-"`
	UpdatedAt        time.Time      `gorm:"column:updated_at;autoUpdateTime;not null;default:now();index" json:"updated_at"`
}

func (Address) TableName() string {
	return "address"
}
