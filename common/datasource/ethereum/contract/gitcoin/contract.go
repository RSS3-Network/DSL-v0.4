package gitcoin

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	EventHashDonationSent = common.BytesToHash(crypto.Keccak256([]byte("DonationSent(address,uint256,address,address)")))

	AddressRound = common.HexToAddress("0x9C3B81967EafBA0a451E324417DD4F3F353b997b")

	AddressQuadraticFundingVotingStrategy = common.HexToAddress("0x7B00EB3175a98CD2055a1c72589776a392cCA9c5")

	EventVoted = crypto.Keccak256Hash([]byte("Voted(address,uint256,address,address,bytes32,uint256,address)"))
)
