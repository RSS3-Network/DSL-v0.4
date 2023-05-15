package blend

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

const Blend = "Blend"

var (
	AddressBlend = common.HexToAddress("0x29469395eAf6f95920E59F858042f0e28D98a20B")

	// EventLoanOfferTaken 0x06a333c2d6fe967ca967f7a35be2eb45e8caeb6cf05e16f55d42b91b5fe31255
	EventLoanOfferTaken = crypto.Keccak256Hash([]byte("LoanOfferTaken(bytes32,uint256,address,address,address,uint256,uint256,uint256,uint256)"))

	// EventRepay 0x2469cc9e12e74c63438d5b1117b318cd3a4cdaf9d659d9eac6d975d14d963254
	EventRepay = crypto.Keccak256Hash([]byte("Repay(uint256,address)"))

	// EventStartAuction 0xe5095dc360d1a56740c946cccc76520c1a1a57381c950520062adeda68dbf572
	EventStartAuction = crypto.Keccak256Hash([]byte("StartAuction(uint256,address)"))

	// EventRefinance 0x558a9295c62e9e1b12a21c8fe816f4816a2e0269a53157edbfa16017b11b9ac9
	EventRefinance = crypto.Keccak256Hash([]byte("Refinance(uint256,address,address,uint256,uint256,uint256)"))

	// EventSeize 0xb71caf41fe0e019dbe21a1ae3493f11a729c31548ed1e304ae7f6e8c8df275de
	EventSeize = crypto.Keccak256Hash([]byte("Seize(uint256,address)"))

	// EventBuyLocked 0x7ffb5bd9cdc79a6f9bc6e00c82f43836e0afbb204d47972001f6e853764a8ef1
	EventBuyLocked = crypto.Keccak256Hash([]byte("BuyLocked(uint256,address,address,address,uint256)"))
)
