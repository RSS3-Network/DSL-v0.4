package gitcoin

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var EventHashDonation = common.BytesToHash(crypto.Keccak256([]byte("DonationSent(address,uint256,address,address)")))
