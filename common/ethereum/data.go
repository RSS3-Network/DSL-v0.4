package ethereum

import (
	"errors"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

var (
	ErrorUnsupportedTransactionType = errors.New("unsupported transaction type")
)

type SourceData struct {
	Transaction *Transaction   `json:"transaction,omitempty"`
	Receipt     *types.Receipt `json:"receipt,omitempty"`
}

type Transaction struct {
	BlockNumber *big.Int        `json:"block_number"`
	Hash        common.Hash     `json:"hash"`
	From        common.Address  `json:"from"`
	To          *common.Address `json:"to,omitempty"`
	ChainID     *big.Int        `json:"chain_id"`
	Type        uint8           `json:"type"`
	Index       int             `json:"index"`
	Value       *big.Int        `json:"value"`
	Gas         uint64          `json:"gas"`
	GasPrice    *big.Int        `json:"gas_price"`
	GasFeeCap   *big.Int        `json:"gas_fee_cap"`
	GasTipCap   *big.Int        `json:"gas_tip_cap"`
	Data        []byte          `json:"data,omitempty"`
	Nonce       uint64          `json:"nonce"`
	Timestamp   time.Time       `json:"timestamp"`
}

func ConvertTransaction(transaction types.Transaction, block types.Block) (*Transaction, error) {
	var (
		transactionMessage types.Message
		err                error
	)

	switch transaction.Type() {
	case types.LegacyTxType:
		transactionMessage, err = transaction.AsMessage(types.NewEIP155Signer(transaction.ChainId()), nil)
	case types.DynamicFeeTxType:
		transactionMessage, err = transaction.AsMessage(types.LatestSignerForChainID(transaction.ChainId()), nil)
	default:
		err = ErrorUnsupportedTransactionType
	}

	if err != nil {
		return nil, err
	}

	index := 0

	for i, internalTransaction := range block.Transactions() {
		if internalTransaction.Hash() == transaction.Hash() {
			index = i

			break
		}
	}

	return &Transaction{
		BlockNumber: block.Number(),
		Hash:        transaction.Hash(),
		From:        transactionMessage.From(),
		To:          transaction.To(),
		ChainID:     transaction.ChainId(),
		Index:       index,
		Value:       transaction.Value(),
		Type:        transaction.Type(),
		Gas:         transaction.Gas(),
		GasPrice:    transaction.GasPrice(),
		GasFeeCap:   transaction.GasFeeCap(),
		GasTipCap:   transaction.GasTipCap(),
		Data:        transaction.Data(),
		Nonce:       transaction.Nonce(),
		Timestamp:   time.Unix(int64(block.Time()), 0),
	}, nil
}
