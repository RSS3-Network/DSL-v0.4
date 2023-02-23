package rss3

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	AddressStaking = common.HexToAddress("0x5301CbBeDc048AbaC7e213184132cf982d593563")

	EventHashWithdrawn      = crypto.Keccak256Hash([]byte("Withdrawn(uint256,address,address,uint256)"))
	EventHashDeposited      = crypto.Keccak256Hash([]byte("Deposited(address,uint256,uint256,uint256)"))
	EventHashRewardsClaimed = crypto.Keccak256Hash([]byte("RewardsClaimed(address,address,uint256)"))
)
