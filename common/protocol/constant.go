package protocol

import "github.com/shopspring/decimal"

const (
	Version         = ""
	ExchangeJob     = "jobs"
	ContentTypeJSON = "application/json"
)

var IndexVirtual = decimal.NewFromInt(-1)
