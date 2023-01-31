package aptos

import "github.com/naturalselectionlabs/pregod/common/database/model/metadata"

const (
	CoinTransferFunc    = "0x1::coin::transfer"
	AccountTransferFunc = "0x1::aptos_account::transfer"

	AptosCoin = "0x1::aptos_coin::AptosCoin"
)

var coinMetadata = map[string]metadata.Token{
	AptosCoin: {
		Name:     "Aptos",
		Symbol:   "APT",
		Decimals: 8,
		Image:    "https://assets.coingecko.com/coins/images/26455/large/aptos_round.png",
		Standard: "Native",
	},
}
