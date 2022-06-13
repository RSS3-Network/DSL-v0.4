package snapshot

import "encoding/json"

type Proposal struct {
	ID       string          `gorm:"column:id;primaryKey" json:"id"`
	SpaceID  string          `gorm:"column:space_id" json:"space_id"`
	Metadata json.RawMessage `gorm:"column:metadata;type:jsonb" json:"metadata"`
}
