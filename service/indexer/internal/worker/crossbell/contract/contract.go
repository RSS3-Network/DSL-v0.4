package contract

import (
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

//go:generate abigen --abi ./profile/profile.abi --pkg profile --type Profile --out ./profile/profile.go
//go:generate abigen --abi ./character/character.abi --pkg character --type Character --out ./character/character.go
//go:generate abigen --abi ./periphery/periphery.abi --pkg periphery --type Periphery --out ./periphery/periphery.go
//go:generate abigen --abi ./linklist/linklist.abi --pkg linklist --type LinkList --out ./linklist/linklist.go

const (
	// BrokenBlockNumber is the block height of the contract non-compatible upgrade
	BrokenBlockNumber = 6552927
)

var (
	AddressCharacter = common.HexToAddress("0xa6f969045641Cf486a747A2688F3a5A6d43cd0D8")
	AddressLinkList  = common.HexToAddress("0xFc8C75bD5c26F50798758f387B698f207a016b6A")
	AddressPeriphery = common.HexToAddress("0x96e96b7AF62D628cE7eb2016D2c1D2786614eA73")

	EventNameTransfer         = "Transfer"
	EventNameSetHandle        = "SetHandle"
	EventNameCharacterCreated = "CharacterCreated"
	EventNamePostNote         = "PostNote"
	EventNameLinkCharacter    = "LinkCharacter"
	EventNameUnlinkCharacter  = "UnlinkCharacter"
	EventNameSetProfileUri    = "SetProfileUri"
	EventNameSetCharacterUri  = "SetCharacterUri"

	EventHashProfileCreated        = common.BytesToHash(crypto.Keccak256([]byte("ProfileCreated(uint256,address,address,string,uint256)")))
	EventHashCharacterCreated      = common.BytesToHash(crypto.Keccak256([]byte("CharacterCreated(uint256,address,address,string,uint256)")))
	EventHashSetHandle             = common.BytesToHash(crypto.Keccak256([]byte("SetHandle(address,uint256,string)")))
	EventHashSetPrimaryProfileId   = common.BytesToHash(crypto.Keccak256([]byte("SetPrimaryProfileId(address,uint256)")))
	EventHashSetPrimaryCharacterId = common.BytesToHash(crypto.Keccak256([]byte("SetPrimaryCharacterId(address,uint256)")))
	EventHashPostNote              = common.BytesToHash(crypto.Keccak256([]byte("PostNote(uint256,uint256,bytes32,bytes32,bytes)")))
	EventHashLinkProfile           = common.BytesToHash(crypto.Keccak256([]byte("LinkProfile(address,uint256,uint256,bytes32,uint256)")))
	EventHashLinkCharacter         = common.BytesToHash(crypto.Keccak256([]byte("LinkCharacter(address,uint256,uint256,bytes32,uint256)")))
	EventHashUnlinkProfile         = common.BytesToHash(crypto.Keccak256([]byte("UnlinkProfile(address,uint256,uint256,bytes32)")))
	EventHashUnlinkCharacter       = common.BytesToHash(crypto.Keccak256([]byte("UnlinkCharacter(address,uint256,uint256,bytes32)")))
	EventHashSetProfileUri         = common.BytesToHash(crypto.Keccak256([]byte("SetProfileUri(uint256,string)")))
	EventHashSetCharacterUri       = common.BytesToHash(crypto.Keccak256([]byte("SetCharacterUri(uint256,string)")))

	LinkItemTypeAnyUri = common.BytesToHash(common.RightPadBytes([]byte("AnyUri"), common.HashLength))

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
	ErrorUnknownEvent           = errors.New("unknown event")
	ErrorUnknownContractAddress = errors.New("unknown contract address")
)
