package ethereum

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/pregod/common/database"
)

type GetAssetTransfersParameter struct {
	FromBlock   int64  `json:"fromBlock,omitempty"`
	FromAddress string `json:"fromAddress,omitempty"`
}

type GetAssetTransfersResult struct {
	Transfers []Transfer
}

type Transfer struct {
	BlockNum int64          `gorm:"column:block_number"`
	Hash     common.Hash    `gorm:"column:hash"`
	From     common.Address `gorm:"column:from_address"`
	To       []byte         `gorm:"column:to_address"`
	Value    float64        `gorm:"column:value"`
	Category string         `gorm:"column:category"`
}

func GetAssetTransfers(ctx context.Context, parameter GetAssetTransfersParameter) (*GetAssetTransfersResult, error) {
	result := GetAssetTransfersResult{}
	fromAddress := common.HexToAddress(parameter.FromAddress)
	blockNum := parameter.FromBlock

	err := database.EthDb().
		Raw("select 'external' as category,value,block_number,hash,from_address,to_address from ethereum.transactions where block_number >= ? and from_address = ?", blockNum, fromAddress).
		Scan(&result.Transfers).
		WithContext(ctx).
		Error
	if err != nil {
		return nil, err
	}
	return &result, err
}
