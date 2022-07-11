package transaction

import (
	"encoding/json"

	"github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type CoinMarketCapCoinInfo struct {
	Name        string `gorm:"column:name;primaryKey"`
	Symbol      string `gorm:"column:symbol;index"`
	Logo        string `gorm:"column:logo"`
	Slug        string `gorm:"column:slug"`
	Category    string `gorm:"column:category"`
	Description string `gorm:"column:description"`
	Decimals    uint8  `gorm:"column:decimals"`
	// "contract_address":[{"contract_address":"0xc98d64da73a6616c42117b582e832812e7b8d57f","platform":{"name":"Ethereum","coin":{"id":"1027","name":"Ethereum","symbol":"ETH","slug":"ethereum"}}}]
	ContractAddress json.RawMessage `json:"contract_address" gorm:"column:contract_address;type:jsonb"`

	// addition fields

	Addresses pq.StringArray `gorm:"column:addresses;type:text[];index"`
}

func (c *CoinMarketCapCoinInfo) FillFields() {
	if len(c.ContractAddress) == 0 {
		return
	}

	// Addresses
	var cas []struct {
		ContractAddress string `json:"contract_address"`
	}
	if err := json.Unmarshal(c.ContractAddress, &cas); err != nil {
		logrus.Error(err)
		return
	}
	c.Addresses = []string{}
	for _, ca := range cas {
		c.Addresses = append(c.Addresses, ca.ContractAddress)
	}
}
