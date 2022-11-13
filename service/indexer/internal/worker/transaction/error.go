package transaction

import "errors"

var (
	ErrorNativeTokenTransferValueIsZero = errors.New("native token transfer value is zero")
	ErrorInvalidContractEvent           = errors.New("invalid contract event")
	ErrorUnsupportedContractEvent       = errors.New("unsupported contract event")
	ErrorInvalidTopicsLength            = errors.New("invalid topics length")
)
