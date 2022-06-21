package protocol

import "github.com/shopspring/decimal"

const (
	Version         = ""
	ExchangeJob     = "jobs"
	ContentTypeJSON = "application/json"
)

var LogIndexVirtual = decimal.NewFromInt(-1)
