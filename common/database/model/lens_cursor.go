package model

type LensCursor struct {
	Address string `gorm:"column:address;primaryKey"`
	Cursor  string `gorm:"column:cursor"`
}
