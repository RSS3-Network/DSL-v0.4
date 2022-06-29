package contract

import (
	"errors"
	"math"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	// BrokenBlockNumber is the block height of the contract non-compatible upgrade
	BrokenBlockNumber = math.MaxInt64
)

var (
	AddressCharacter = common.HexToAddress("0xa6f969045641Cf486a747A2688F3a5A6d43cd0D8")
	AddressLinkList  = common.HexToAddress("0xFc8C75bD5c26F50798758f387B698f207a016b6A")

	EventNameTransfer       = "Transfer"
	EventNameProfileCreated = "ProfileCreated"
	EventNamePostNote       = "PostNote"

	EventHashProfileCreated = common.BytesToHash(crypto.Keccak256([]byte("ProfileCreated(uint256,address,address,string,uint256)")))
	EventHashPostNote       = common.BytesToHash(crypto.Keccak256([]byte("PostNote(uint256,uint256,bytes32,bytes32,bytes)")))

	LinkTypeFollow  = "follow"
	LinkTypeLike    = "like"
	LinkTypeComment = "comment"

	LinkTypeMap = map[common.Hash]string{
		common.BytesToHash(common.RightPadBytes([]byte(LinkTypeFollow), common.HashLength)):  LinkTypeFollow,
		common.BytesToHash(common.RightPadBytes([]byte(LinkTypeLike), common.HashLength)):    LinkTypeLike,
		common.BytesToHash(common.RightPadBytes([]byte(LinkTypeComment), common.HashLength)): LinkTypeComment,
	}
)

var (
	ErrorUnknownUnknownEvent    = errors.New("unknown event")
	ErrorUnknownContractAddress = errors.New("unknown contract address")
)
