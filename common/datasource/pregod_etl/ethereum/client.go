package ethereum

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/pregod/common/database"
)

const (
	DefaultTransactionLimit = 30000
	DefaultTransferLimit    = 10000
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
}

// Deprecated: use kurora query getAssestTransactionAndLogs instead
func GetAssetTransfers(ctx context.Context, parameter GetAssetTransfersParameter) (*GetAssetTransfersResult, error) {
	result := GetAssetTransfersResult{}
	fromAddress := common.HexToAddress(parameter.FromAddress)
	blockNum := parameter.FromBlock

	err := database.EthDb().
		Raw("select * from "+
			"("+
			"(select value,block_number,hash,from_address,to_address from ethereum.transactions where block_number >= ? and from_address = ? limit ?) "+
			"UNION ALL "+
			"(select value,block_number,transaction_hash as hash, from_address,to_address from ethereum.token_transfers where block_number >= ? and to_address = ? limit ?)"+
			") "+
			"as foo order by foo.block_number",
			blockNum, fromAddress, DefaultTransactionLimit, blockNum, fromAddress, DefaultTransferLimit).
		Scan(&result.Transfers).
		WithContext(ctx).
		Error
	if err != nil {
		return nil, err
	}
	return &result, err
}
