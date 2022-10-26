package game

type GameContract struct {
	ContractAddress string `gorm:"column:contract_address;primaryKey"`
	Network         string `gorm:"column:network;primaryKey"`
	Platform        string `gorm:"column:platform"`
}
