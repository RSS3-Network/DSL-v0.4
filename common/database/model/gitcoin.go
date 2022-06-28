package model

type GitcoinProject struct {
	Active          bool   `gorm:"column:active" json:"active"`
	ID              int    `gorm:"column:id" json:"id"`
	Title           string `gorm:"column:title" json:"title"`
	Slug            string `gorm:"column:slug" json:"slug"`
	Description     string `gorm:"column:description" json:"description"`
	ReferUrl        string `gorm:"column:reference_url" json:"reference_url"`
	Logo            string `gorm:"column:logo" json:"logo"`
	AdminAddress    string `gorm:"column:admin_address" json:"admin_address"`
	TokenAddress    string `gorm:"column:token_address" json:"token_address"`
	TokenSymbol     string `gorm:"column:token_symbol" json:"token_symbol"`
	ContractAddress string `gorm:"column:contract_address" json:"contract_address"`
}

func (GitcoinProject) TableName() string {
	return "reptile-gitcoin.data"
}
