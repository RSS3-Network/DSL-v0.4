/*
Cache transaction info from api.zksync.io
*/
package model

import (
	"encoding/json"

	"github.com/sirupsen/logrus"
)

// https://docs.zksync.io/apiv02-docs/#tokens-api-v0.2-tokens-tokenlike
type GetTokenInfo struct {
	ID             *int64 `json:"id" gorm:"column:id;primaryKey"`
	Address        string `json:"address" gorm:"column:address"`
	Symbol         string `json:"symbol" gorm:"column:symbol"`
	Decimals       uint8  `json:"decimals" gorm:"column:decimals"`
	EnabledForFees bool   `json:"enabledForFees" gorm:"column:enabledForFees"`
}

// https://docs.zksync.io/apiv02-docs/#tokens-api-v0.2-tokens-nft-id
type GetNFTTokenInfo struct {
	ID               *int64 `json:"id" gorm:"column:id;primaryKey"`
	ContentHash      string `json:"contentHash" gorm:"column:contentHash"`
	CreatorID        int64  `json:"creatorId" gorm:"column:creatorId"`
	CreatorAddress   string `json:"creatorAddress" gorm:"column:creatorAddress"`
	SerialID         int    `json:"serialId" gorm:"column:serialId"`
	Address          string `json:"address" gorm:"column:address"`
	Symbol           string `json:"symbol" gorm:"column:symbol"`
	CurrentFactory   string `json:"currentFactory" gorm:"column:currentFactory"`
	WithdrawnFactory string `json:"withdrawnFactory" gorm:"column:withdrawnFactory"`
}

func (i *GetNFTTokenInfo) Bytes() []byte {
	result, err := json.Marshal(i)
	if err != nil {
		logrus.Error(err)
	}
	return result
}
