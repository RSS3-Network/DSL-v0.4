package snapshot

import "encoding/json"

type Space struct {
	ID       string          `gorm:"column:id;primaryKey" json:"id"`
	Metadata json.RawMessage `gorm:"column:metadata;type:jsonb" json:"metadata"`
}
