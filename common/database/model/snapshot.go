package model

import (
	"encoding/json"
	"time"
)

type SnapshotSpace struct {
	ID        string          `gorm:"column:id;primaryKey" json:"id"`
	Metadata  json.RawMessage `gorm:"column:metadata;type:jsonb" json:"metadata"`
	Network   string          `gorm:"column:network" json:"network"`
	CreatedAt time.Time       `gorm:"column:created_at;autoCreateTime;not null;default:now();index"`
	UpdatedAt time.Time       `gorm:"column:updated_at;autoUpdateTime;not null;default:now();index"`
}

func (SnapshotSpace) TableName() string {
	return "snapshot_space"
}

type SnapshotProposal struct {
	ID          string          `gorm:"column:id;primaryKey" json:"id"`
	SpaceID     string          `gorm:"column:space_id" json:"space_id"`
	Author      string          `gorm:"column:author" json:"author"`
	Metadata    json.RawMessage `gorm:"column:metadata;type:jsonb" json:"metadata"`
	DateCreated time.Time       `gorm:"column:date_created;index:index_note_date_created" json:"date_created"`
	CreatedAt   time.Time       `gorm:"column:created_at;autoCreateTime;not null;default:now();index"`
	UpdatedAt   time.Time       `gorm:"column:updated_at;autoUpdateTime;not null;default:now();index"`
}

func (SnapshotProposal) TableName() string {
	return "snapshot_proposal"
}

type SnapshotVote struct {
	ID          string          `gorm:"column:id;primaryKey" json:"id"`
	Voter       string          `gorm:"column:voter" json:"voter"`
	Choice      json.RawMessage `gorm:"column:choice;type:jsonb" json:"choice"`
	ProposalID  string          `gorm:"column:proposal_id" json:"proposal_id"`
	SpaceID     string          `gorm:"column:space_id" json:"space_id"`
	DateCreated time.Time       `gorm:"column:date_created;index:index_note_date_created" json:"date_created"`
	CreatedAt   time.Time       `gorm:"column:created_at;autoCreateTime;not null;default:now();index"`
	UpdatedAt   time.Time       `gorm:"column:updated_at;autoUpdateTime;not null;default:now();index"`
}

func (SnapshotVote) TableName() string {
	return "snapshot_vote"
}
