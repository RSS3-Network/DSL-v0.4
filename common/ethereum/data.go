package ethereum

import (
	"errors"

	"github.com/ethereum/go-ethereum/core/types"
)

var ErrorUnsupportedTransactionType = errors.New("unsupported transaction type")

type SourceData struct {
	Transaction *types.Transaction `json:"transaction,omitempty"`
	Receipt     *types.Receipt     `json:"receipt,omitempty"`
}
