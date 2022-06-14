package model

import "encoding/json"

type SnapshotSpace struct {
	SerialID int             `gorm:"column:serial_id;" json:"serial_id" `
	ID       string          `gorm:"column:id;primaryKey" json:"id"`
	Metadata json.RawMessage `gorm:"column:metadata;type:jsonb" json:"metadata"`
}

func (SnapshotSpace) TableName() string {
	return "snapshot-space"
}

type SnapshotProposal struct {
	SerialID int             `gorm:"column:serial_id;" json:"serial_id" `
	ID       string          `gorm:"column:id;primaryKey" json:"id"`
	SpaceID  string          `gorm:"column:space_id" json:"space_id"`
	Metadata json.RawMessage `gorm:"column:metadata;type:jsonb" json:"metadata"`
}

func (SnapshotProposal) TableName() string {
	return "snapshot-proposal"
}

type SnapshotVote struct {
	SerialID   int             `gorm:"column:serial_id;" json:"serial_id" `
	ID         string          `gorm:"column:id;primaryKey" json:"id"`
	ProposalID string          `gorm:"column:proposal_id" json:"proposal_id"`
	Metadata   json.RawMessage `gorm:"column:metadata;type:jsonb" json:"metadata"`
}

func (SnapshotVote) TableName() string {
	return "snapshot-vote"
}
