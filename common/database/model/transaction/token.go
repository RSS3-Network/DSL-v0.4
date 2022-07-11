package transaction

import (
	"encoding/json"
	"time"
)

type Token struct {
	Name            string          `gorm:"column:name"`
	Symbol          string          `gorm:"column:symbol"`
	Decimal         int             `gorm:"column:decimal"`
	Standard        string          `gorm:"column:standard"`
	ContractAddress string          `gorm:"column:contract_address"`
	Metadata        json.RawMessage `gorm:"column:metadata;type:jsonb"`
	Network         string          `gorm:"column:network"`
	Source          string          `gorm:"column:source"`
	CreatedAt       time.Time       `gorm:"column:created_at;autoCreateTime;not null;default:now();index"`
	UpdatedAt       time.Time       `gorm:"column:updated_at;autoUpdateTime;not null;default:now();index"`
}
