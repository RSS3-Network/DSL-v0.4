package collectibe

import "time"

type FriendTech struct {
	ID              int64     `gorm:"column:id;not null;primaryKey" json:"-"`
	Address         string    `gorm:"column:address;not null;primaryKey" json:"-"`
	TwitterUsername string    `gorm:"column:twitterUsername;not null" json:"twitterUsername"`
	TwitterName     string    `gorm:"column:twitterName;not null" json:"twitterName"`
	TwitterPfpUrl   string    `gorm:"column:twitterPfpUrl" json:"twitterPfpUrl"`
	TwitterUserId   string    `gorm:"column:twitterUserId" json:"twitterUserId"`
	CreatedAt       time.Time `gorm:"column:created_at;autoCreateTime;not null;default:now();index" json:"-"`
	UpdatedAt       time.Time `gorm:"column:updated_at;autoUpdateTime;not null;default:now();index" json:"-"`
}

func (FriendTech) TableName() string {
	return "friend_tech"
}
