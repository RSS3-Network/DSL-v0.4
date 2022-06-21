// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// DataTypesCreateProfileData is an auto generated low-level Go binding around an user-defined struct.
type DataTypesCreateProfileData struct {
	To                 common.Address
	Handle             string
	Uri                string
	LinkModule         common.Address
	LinkModuleInitData []byte
}

// DataTypesERC721Struct is an auto generated low-level Go binding around an user-defined struct.
type DataTypesERC721Struct struct {
	TokenAddress  common.Address
	Erc721TokenId *big.Int
}

// DataTypesMintNoteData is an auto generated low-level Go binding around an user-defined struct.
type DataTypesMintNoteData struct {
	ProfileId      *big.Int
	NoteId         *big.Int
	To             common.Address
	MintModuleData []byte
}

// DataTypesNote is an auto generated low-level Go binding around an user-defined struct.
type DataTypesNote struct {
	LinkItemType [32]byte
	LinkKey      [32]byte
	ContentUri   string
	LinkModule   common.Address
	MintModule   common.Address
	MintNFT      common.Address
	Deleted      bool
	Locked       bool
}

// DataTypesNoteStruct is an auto generated low-level Go binding around an user-defined struct.
type DataTypesNoteStruct struct {
	ProfileId *big.Int
	NoteId    *big.Int
}

// DataTypesPostNoteData is an auto generated low-level Go binding around an user-defined struct.
type DataTypesPostNoteData struct {
	ProfileId          *big.Int
	ContentUri         string
	LinkModule         common.Address
	LinkModuleInitData []byte
	MintModule         common.Address
	MintModuleInitData []byte
	Locked             bool
}

// DataTypesProfile is an auto generated low-level Go binding around an user-defined struct.
type DataTypesProfile struct {
	ProfileId   *big.Int
	Handle      string
	Uri         string
	NoteCount   *big.Int
	SocialToken common.Address
	LinkModule  common.Address
}

// DataTypesProfileLinkStruct is an auto generated low-level Go binding around an user-defined struct.
type DataTypesProfileLinkStruct struct {
	FromProfileId *big.Int
	ToProfileId   *big.Int
	LinkType      [32]byte
}

// DataTypescreateThenLinkProfileData is an auto generated low-level Go binding around an user-defined struct.
type DataTypescreateThenLinkProfileData struct {
	FromProfileId *big.Int
	To            common.Address
	LinkType      [32]byte
}

// DataTypeslinkAddressData is an auto generated low-level Go binding around an user-defined struct.
type DataTypeslinkAddressData struct {
	FromProfileId *big.Int
	EthAddress    common.Address
	LinkType      [32]byte
	Data          []byte
}

// DataTypeslinkAnyUriData is an auto generated low-level Go binding around an user-defined struct.
type DataTypeslinkAnyUriData struct {
	FromProfileId *big.Int
	ToUri         string
	LinkType      [32]byte
	Data          []byte
}

// DataTypeslinkERC721Data is an auto generated low-level Go binding around an user-defined struct.
type DataTypeslinkERC721Data struct {
	FromProfileId *big.Int
	TokenAddress  common.Address
	TokenId       *big.Int
	LinkType      [32]byte
	Data          []byte
}

// DataTypeslinkLinklistData is an auto generated low-level Go binding around an user-defined struct.
type DataTypeslinkLinklistData struct {
	FromProfileId *big.Int
	ToLinkListId  *big.Int
	LinkType      [32]byte
	Data          []byte
}

// DataTypeslinkNoteData is an auto generated low-level Go binding around an user-defined struct.
type DataTypeslinkNoteData struct {
	FromProfileId *big.Int
	ToProfileId   *big.Int
	ToNoteId      *big.Int
	LinkType      [32]byte
	Data          []byte
}

// DataTypeslinkProfileData is an auto generated low-level Go binding around an user-defined struct.
type DataTypeslinkProfileData struct {
	FromProfileId *big.Int
	ToProfileId   *big.Int
	LinkType      [32]byte
	Data          []byte
}

// DataTypessetLinkModule4AddressData is an auto generated low-level Go binding around an user-defined struct.
type DataTypessetLinkModule4AddressData struct {
	Account            common.Address
	LinkModule         common.Address
	LinkModuleInitData []byte
}

// DataTypessetLinkModule4ERC721Data is an auto generated low-level Go binding around an user-defined struct.
type DataTypessetLinkModule4ERC721Data struct {
	TokenAddress       common.Address
	TokenId            *big.Int
	LinkModule         common.Address
	LinkModuleInitData []byte
}

// DataTypessetLinkModule4LinklistData is an auto generated low-level Go binding around an user-defined struct.
type DataTypessetLinkModule4LinklistData struct {
	LinklistId         *big.Int
	LinkModule         common.Address
	LinkModuleInitData []byte
}

// DataTypessetLinkModule4NoteData is an auto generated low-level Go binding around an user-defined struct.
type DataTypessetLinkModule4NoteData struct {
	ProfileId          *big.Int
	NoteId             *big.Int
	LinkModule         common.Address
	LinkModuleInitData []byte
}

// DataTypessetLinkModule4ProfileData is an auto generated low-level Go binding around an user-defined struct.
type DataTypessetLinkModule4ProfileData struct {
	ProfileId          *big.Int
	LinkModule         common.Address
	LinkModuleInitData []byte
}

// DataTypessetMintModule4NoteData is an auto generated low-level Go binding around an user-defined struct.
type DataTypessetMintModule4NoteData struct {
	ProfileId          *big.Int
	NoteId             *big.Int
	MintModule         common.Address
	MintModuleInitData []byte
}

// DataTypesunlinkAnyUriData is an auto generated low-level Go binding around an user-defined struct.
type DataTypesunlinkAnyUriData struct {
	FromProfileId *big.Int
	ToUri         string
	LinkType      [32]byte
}

// DataTypesunlinkERC721Data is an auto generated low-level Go binding around an user-defined struct.
type DataTypesunlinkERC721Data struct {
	FromProfileId *big.Int
	TokenAddress  common.Address
	TokenId       *big.Int
	LinkType      [32]byte
}

// DataTypesunlinkNoteData is an auto generated low-level Go binding around an user-defined struct.
type DataTypesunlinkNoteData struct {
	FromProfileId *big.Int
	ToProfileId   *big.Int
	ToNoteId      *big.Int
	LinkType      [32]byte
}

// DataTypesunlinkProfileData is an auto generated low-level Go binding around an user-defined struct.
type DataTypesunlinkProfileData struct {
	FromProfileId *big.Int
	ToProfileId   *big.Int
	LinkType      [32]byte
}

// ERC721MetaData contains all meta data concerning the ERC721 contract.
var ERC721MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"type\":\"address\",\"name\":\"owner\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"address\",\"name\":\"approved\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\",\"indexed\":true}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ApprovalForAll\",\"inputs\":[{\"type\":\"address\",\"name\":\"owner\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"address\",\"name\":\"operator\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"bool\",\"name\":\"approved\",\"internalType\":\"bool\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"type\":\"address\",\"name\":\"from\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"address\",\"name\":\"to\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\",\"indexed\":true}],\"anonymous\":false},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"approve\",\"inputs\":[{\"type\":\"address\",\"name\":\"to\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"attachLinklist\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"linklistId\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"profileId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"balanceOf\",\"inputs\":[{\"type\":\"address\",\"name\":\"owner\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"burn\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"canCreate\",\"inputs\":[{\"type\":\"string\",\"name\":\"handle\",\"internalType\":\"string\"},{\"type\":\"address\",\"name\":\"account\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"createProfile\",\"inputs\":[{\"type\":\"tuple\",\"name\":\"vars\",\"internalType\":\"structDataTypes.CreateProfileData\",\"components\":[{\"type\":\"address\",\"name\":\"to\",\"internalType\":\"address\"},{\"type\":\"string\",\"name\":\"handle\",\"internalType\":\"string\"},{\"type\":\"string\",\"name\":\"uri\",\"internalType\":\"string\"},{\"type\":\"address\",\"name\":\"linkModule\",\"internalType\":\"address\"},{\"type\":\"bytes\",\"name\":\"linkModuleInitData\",\"internalType\":\"bytes\"}]}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"createThenLinkProfile\",\"inputs\":[{\"type\":\"tuple\",\"name\":\"vars\",\"internalType\":\"structDataTypes.createThenLinkProfileData\",\"components\":[{\"type\":\"uint256\",\"name\":\"fromProfileId\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"to\",\"internalType\":\"address\"},{\"type\":\"bytes32\",\"name\":\"linkType\",\"internalType\":\"bytes32\"}]}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"deleteNote\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"profileId\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"noteId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"detachLinklist\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"linklistId\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"profileId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"getApproved\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"string\",\"name\":\"\",\"internalType\":\"string\"}],\"name\":\"getHandle\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"profileId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"getLinkModule4Address\",\"inputs\":[{\"type\":\"address\",\"name\":\"account\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"getLinkModule4ERC721\",\"inputs\":[{\"type\":\"address\",\"name\":\"tokenAddress\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"getLinkModule4Linklist\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"getLinklistContract\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"getLinklistId\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"profileId\",\"internalType\":\"uint256\"},{\"type\":\"bytes32\",\"name\":\"linkType\",\"internalType\":\"bytes32\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bytes32\",\"name\":\"\",\"internalType\":\"bytes32\"}],\"name\":\"getLinklistType\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"linkListId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"string\",\"name\":\"\",\"internalType\":\"string\"}],\"name\":\"getLinklistUri\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"tuple\",\"name\":\"\",\"internalType\":\"structDataTypes.Note\",\"components\":[{\"type\":\"bytes32\",\"name\":\"linkItemType\",\"internalType\":\"bytes32\"},{\"type\":\"bytes32\",\"name\":\"linkKey\",\"internalType\":\"bytes32\"},{\"type\":\"string\",\"name\":\"contentUri\",\"internalType\":\"string\"},{\"type\":\"address\",\"name\":\"linkModule\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"mintModule\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"mintNFT\",\"internalType\":\"address\"},{\"type\":\"bool\",\"name\":\"deleted\",\"internalType\":\"bool\"},{\"type\":\"bool\",\"name\":\"locked\",\"internalType\":\"bool\"}]}],\"name\":\"getNote\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"profileId\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"noteId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"getOperator\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"profileId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"getPrimaryProfileId\",\"inputs\":[{\"type\":\"address\",\"name\":\"account\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"tuple\",\"name\":\"\",\"internalType\":\"structDataTypes.Profile\",\"components\":[{\"type\":\"uint256\",\"name\":\"profileId\",\"internalType\":\"uint256\"},{\"type\":\"string\",\"name\":\"handle\",\"internalType\":\"string\"},{\"type\":\"string\",\"name\":\"uri\",\"internalType\":\"string\"},{\"type\":\"uint256\",\"name\":\"noteCount\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"socialToken\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"linkModule\",\"internalType\":\"address\"}]}],\"name\":\"getProfile\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"profileId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"tuple\",\"name\":\"\",\"internalType\":\"structDataTypes.Profile\",\"components\":[{\"type\":\"uint256\",\"name\":\"profileId\",\"internalType\":\"uint256\"},{\"type\":\"string\",\"name\":\"handle\",\"internalType\":\"string\"},{\"type\":\"string\",\"name\":\"uri\",\"internalType\":\"string\"},{\"type\":\"uint256\",\"name\":\"noteCount\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"socialToken\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"linkModule\",\"internalType\":\"address\"}]}],\"name\":\"getProfileByHandle\",\"inputs\":[{\"type\":\"string\",\"name\":\"handle\",\"internalType\":\"string\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"string\",\"name\":\"\",\"internalType\":\"string\"}],\"name\":\"getProfileUri\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"profileId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"pure\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"getRevision\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"initialize\",\"inputs\":[{\"type\":\"string\",\"name\":\"_name\",\"internalType\":\"string\"},{\"type\":\"string\",\"name\":\"_symbol\",\"internalType\":\"string\"},{\"type\":\"address\",\"name\":\"_linklistContract\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"_mintNFTImpl\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"_periphery\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"_resolver\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"isApprovedForAll\",\"inputs\":[{\"type\":\"address\",\"name\":\"owner\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"operator\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"isPrimaryProfile\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"profileId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"linkAddress\",\"inputs\":[{\"type\":\"tuple\",\"name\":\"vars\",\"internalType\":\"structDataTypes.linkAddressData\",\"components\":[{\"type\":\"uint256\",\"name\":\"fromProfileId\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"ethAddress\",\"internalType\":\"address\"},{\"type\":\"bytes32\",\"name\":\"linkType\",\"internalType\":\"bytes32\"},{\"type\":\"bytes\",\"name\":\"data\",\"internalType\":\"bytes\"}]}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"linkAnyUri\",\"inputs\":[{\"type\":\"tuple\",\"name\":\"vars\",\"internalType\":\"structDataTypes.linkAnyUriData\",\"components\":[{\"type\":\"uint256\",\"name\":\"fromProfileId\",\"internalType\":\"uint256\"},{\"type\":\"string\",\"name\":\"toUri\",\"internalType\":\"string\"},{\"type\":\"bytes32\",\"name\":\"linkType\",\"internalType\":\"bytes32\"},{\"type\":\"bytes\",\"name\":\"data\",\"internalType\":\"bytes\"}]}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"linkERC721\",\"inputs\":[{\"type\":\"tuple\",\"name\":\"vars\",\"internalType\":\"structDataTypes.linkERC721Data\",\"components\":[{\"type\":\"uint256\",\"name\":\"fromProfileId\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"tokenAddress\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"},{\"type\":\"bytes32\",\"name\":\"linkType\",\"internalType\":\"bytes32\"},{\"type\":\"bytes\",\"name\":\"data\",\"internalType\":\"bytes\"}]}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"linkLinklist\",\"inputs\":[{\"type\":\"tuple\",\"name\":\"vars\",\"internalType\":\"structDataTypes.linkLinklistData\",\"components\":[{\"type\":\"uint256\",\"name\":\"fromProfileId\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"toLinkListId\",\"internalType\":\"uint256\"},{\"type\":\"bytes32\",\"name\":\"linkType\",\"internalType\":\"bytes32\"},{\"type\":\"bytes\",\"name\":\"data\",\"internalType\":\"bytes\"}]}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"linkNote\",\"inputs\":[{\"type\":\"tuple\",\"name\":\"vars\",\"internalType\":\"structDataTypes.linkNoteData\",\"components\":[{\"type\":\"uint256\",\"name\":\"fromProfileId\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"toProfileId\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"toNoteId\",\"internalType\":\"uint256\"},{\"type\":\"bytes32\",\"name\":\"linkType\",\"internalType\":\"bytes32\"},{\"type\":\"bytes\",\"name\":\"data\",\"internalType\":\"bytes\"}]}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"linkProfile\",\"inputs\":[{\"type\":\"tuple\",\"name\":\"vars\",\"internalType\":\"structDataTypes.linkProfileData\",\"components\":[{\"type\":\"uint256\",\"name\":\"fromProfileId\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"toProfileId\",\"internalType\":\"uint256\"},{\"type\":\"bytes32\",\"name\":\"linkType\",\"internalType\":\"bytes32\"},{\"type\":\"bytes\",\"name\":\"data\",\"internalType\":\"bytes\"}]}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"linkProfileLink\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"fromProfileId\",\"internalType\":\"uint256\"},{\"type\":\"tuple\",\"name\":\"linkData\",\"internalType\":\"structDataTypes.ProfileLinkStruct\",\"components\":[{\"type\":\"uint256\",\"name\":\"fromProfileId\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"toProfileId\",\"internalType\":\"uint256\"},{\"type\":\"bytes32\",\"name\":\"linkType\",\"internalType\":\"bytes32\"}]},{\"type\":\"bytes32\",\"name\":\"linkType\",\"internalType\":\"bytes32\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"lockNote\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"profileId\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"noteId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"mintNote\",\"inputs\":[{\"type\":\"tuple\",\"name\":\"vars\",\"internalType\":\"structDataTypes.MintNoteData\",\"components\":[{\"type\":\"uint256\",\"name\":\"profileId\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"noteId\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"to\",\"internalType\":\"address\"},{\"type\":\"bytes\",\"name\":\"mintModuleData\",\"internalType\":\"bytes\"}]}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"string\",\"name\":\"\",\"internalType\":\"string\"}],\"name\":\"name\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"ownerOf\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"postNote\",\"inputs\":[{\"type\":\"tuple\",\"name\":\"vars\",\"internalType\":\"structDataTypes.PostNoteData\",\"components\":[{\"type\":\"uint256\",\"name\":\"profileId\",\"internalType\":\"uint256\"},{\"type\":\"string\",\"name\":\"contentUri\",\"internalType\":\"string\"},{\"type\":\"address\",\"name\":\"linkModule\",\"internalType\":\"address\"},{\"type\":\"bytes\",\"name\":\"linkModuleInitData\",\"internalType\":\"bytes\"},{\"type\":\"address\",\"name\":\"mintModule\",\"internalType\":\"address\"},{\"type\":\"bytes\",\"name\":\"mintModuleInitData\",\"internalType\":\"bytes\"},{\"type\":\"bool\",\"name\":\"locked\",\"internalType\":\"bool\"}]}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"postNote4Address\",\"inputs\":[{\"type\":\"tuple\",\"name\":\"noteData\",\"internalType\":\"structDataTypes.PostNoteData\",\"components\":[{\"type\":\"uint256\",\"name\":\"profileId\",\"internalType\":\"uint256\"},{\"type\":\"string\",\"name\":\"contentUri\",\"internalType\":\"string\"},{\"type\":\"address\",\"name\":\"linkModule\",\"internalType\":\"address\"},{\"type\":\"bytes\",\"name\":\"linkModuleInitData\",\"internalType\":\"bytes\"},{\"type\":\"address\",\"name\":\"mintModule\",\"internalType\":\"address\"},{\"type\":\"bytes\",\"name\":\"mintModuleInitData\",\"internalType\":\"bytes\"},{\"type\":\"bool\",\"name\":\"locked\",\"internalType\":\"bool\"}]},{\"type\":\"address\",\"name\":\"ethAddress\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"postNote4AnyUri\",\"inputs\":[{\"type\":\"tuple\",\"name\":\"postNoteData\",\"internalType\":\"structDataTypes.PostNoteData\",\"components\":[{\"type\":\"uint256\",\"name\":\"profileId\",\"internalType\":\"uint256\"},{\"type\":\"string\",\"name\":\"contentUri\",\"internalType\":\"string\"},{\"type\":\"address\",\"name\":\"linkModule\",\"internalType\":\"address\"},{\"type\":\"bytes\",\"name\":\"linkModuleInitData\",\"internalType\":\"bytes\"},{\"type\":\"address\",\"name\":\"mintModule\",\"internalType\":\"address\"},{\"type\":\"bytes\",\"name\":\"mintModuleInitData\",\"internalType\":\"bytes\"},{\"type\":\"bool\",\"name\":\"locked\",\"internalType\":\"bool\"}]},{\"type\":\"string\",\"name\":\"uri\",\"internalType\":\"string\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"postNote4ERC721\",\"inputs\":[{\"type\":\"tuple\",\"name\":\"postNoteData\",\"internalType\":\"structDataTypes.PostNoteData\",\"components\":[{\"type\":\"uint256\",\"name\":\"profileId\",\"internalType\":\"uint256\"},{\"type\":\"string\",\"name\":\"contentUri\",\"internalType\":\"string\"},{\"type\":\"address\",\"name\":\"linkModule\",\"internalType\":\"address\"},{\"type\":\"bytes\",\"name\":\"linkModuleInitData\",\"internalType\":\"bytes\"},{\"type\":\"address\",\"name\":\"mintModule\",\"internalType\":\"address\"},{\"type\":\"bytes\",\"name\":\"mintModuleInitData\",\"internalType\":\"bytes\"},{\"type\":\"bool\",\"name\":\"locked\",\"internalType\":\"bool\"}]},{\"type\":\"tuple\",\"name\":\"erc721\",\"internalType\":\"structDataTypes.ERC721Struct\",\"components\":[{\"type\":\"address\",\"name\":\"tokenAddress\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"erc721TokenId\",\"internalType\":\"uint256\"}]}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"postNote4Linklist\",\"inputs\":[{\"type\":\"tuple\",\"name\":\"noteData\",\"internalType\":\"structDataTypes.PostNoteData\",\"components\":[{\"type\":\"uint256\",\"name\":\"profileId\",\"internalType\":\"uint256\"},{\"type\":\"string\",\"name\":\"contentUri\",\"internalType\":\"string\"},{\"type\":\"address\",\"name\":\"linkModule\",\"internalType\":\"address\"},{\"type\":\"bytes\",\"name\":\"linkModuleInitData\",\"internalType\":\"bytes\"},{\"type\":\"address\",\"name\":\"mintModule\",\"internalType\":\"address\"},{\"type\":\"bytes\",\"name\":\"mintModuleInitData\",\"internalType\":\"bytes\"},{\"type\":\"bool\",\"name\":\"locked\",\"internalType\":\"bool\"}]},{\"type\":\"uint256\",\"name\":\"toLinklistId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"postNote4Note\",\"inputs\":[{\"type\":\"tuple\",\"name\":\"postNoteData\",\"internalType\":\"structDataTypes.PostNoteData\",\"components\":[{\"type\":\"uint256\",\"name\":\"profileId\",\"internalType\":\"uint256\"},{\"type\":\"string\",\"name\":\"contentUri\",\"internalType\":\"string\"},{\"type\":\"address\",\"name\":\"linkModule\",\"internalType\":\"address\"},{\"type\":\"bytes\",\"name\":\"linkModuleInitData\",\"internalType\":\"bytes\"},{\"type\":\"address\",\"name\":\"mintModule\",\"internalType\":\"address\"},{\"type\":\"bytes\",\"name\":\"mintModuleInitData\",\"internalType\":\"bytes\"},{\"type\":\"bool\",\"name\":\"locked\",\"internalType\":\"bool\"}]},{\"type\":\"tuple\",\"name\":\"note\",\"internalType\":\"structDataTypes.NoteStruct\",\"components\":[{\"type\":\"uint256\",\"name\":\"profileId\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"noteId\",\"internalType\":\"uint256\"}]}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"postNote4Profile\",\"inputs\":[{\"type\":\"tuple\",\"name\":\"postNoteData\",\"internalType\":\"structDataTypes.PostNoteData\",\"components\":[{\"type\":\"uint256\",\"name\":\"profileId\",\"internalType\":\"uint256\"},{\"type\":\"string\",\"name\":\"contentUri\",\"internalType\":\"string\"},{\"type\":\"address\",\"name\":\"linkModule\",\"internalType\":\"address\"},{\"type\":\"bytes\",\"name\":\"linkModuleInitData\",\"internalType\":\"bytes\"},{\"type\":\"address\",\"name\":\"mintModule\",\"internalType\":\"address\"},{\"type\":\"bytes\",\"name\":\"mintModuleInitData\",\"internalType\":\"bytes\"},{\"type\":\"bool\",\"name\":\"locked\",\"internalType\":\"bool\"}]},{\"type\":\"uint256\",\"name\":\"toProfileId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"resolver\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"safeTransferFrom\",\"inputs\":[{\"type\":\"address\",\"name\":\"from\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"to\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"safeTransferFrom\",\"inputs\":[{\"type\":\"address\",\"name\":\"from\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"to\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"},{\"type\":\"bytes\",\"name\":\"_data\",\"internalType\":\"bytes\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setApprovalForAll\",\"inputs\":[{\"type\":\"address\",\"name\":\"operator\",\"internalType\":\"address\"},{\"type\":\"bool\",\"name\":\"approved\",\"internalType\":\"bool\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setHandle\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"profileId\",\"internalType\":\"uint256\"},{\"type\":\"string\",\"name\":\"newHandle\",\"internalType\":\"string\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setLinkModule4Address\",\"inputs\":[{\"type\":\"tuple\",\"name\":\"vars\",\"internalType\":\"structDataTypes.setLinkModule4AddressData\",\"components\":[{\"type\":\"address\",\"name\":\"account\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"linkModule\",\"internalType\":\"address\"},{\"type\":\"bytes\",\"name\":\"linkModuleInitData\",\"internalType\":\"bytes\"}]}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setLinkModule4ERC721\",\"inputs\":[{\"type\":\"tuple\",\"name\":\"vars\",\"internalType\":\"structDataTypes.setLinkModule4ERC721Data\",\"components\":[{\"type\":\"address\",\"name\":\"tokenAddress\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"linkModule\",\"internalType\":\"address\"},{\"type\":\"bytes\",\"name\":\"linkModuleInitData\",\"internalType\":\"bytes\"}]}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setLinkModule4Linklist\",\"inputs\":[{\"type\":\"tuple\",\"name\":\"vars\",\"internalType\":\"structDataTypes.setLinkModule4LinklistData\",\"components\":[{\"type\":\"uint256\",\"name\":\"linklistId\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"linkModule\",\"internalType\":\"address\"},{\"type\":\"bytes\",\"name\":\"linkModuleInitData\",\"internalType\":\"bytes\"}]}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setLinkModule4Note\",\"inputs\":[{\"type\":\"tuple\",\"name\":\"vars\",\"internalType\":\"structDataTypes.setLinkModule4NoteData\",\"components\":[{\"type\":\"uint256\",\"name\":\"profileId\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"noteId\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"linkModule\",\"internalType\":\"address\"},{\"type\":\"bytes\",\"name\":\"linkModuleInitData\",\"internalType\":\"bytes\"}]}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setLinkModule4Profile\",\"inputs\":[{\"type\":\"tuple\",\"name\":\"vars\",\"internalType\":\"structDataTypes.setLinkModule4ProfileData\",\"components\":[{\"type\":\"uint256\",\"name\":\"profileId\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"linkModule\",\"internalType\":\"address\"},{\"type\":\"bytes\",\"name\":\"linkModuleInitData\",\"internalType\":\"bytes\"}]}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setLinklistUri\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"linklistId\",\"internalType\":\"uint256\"},{\"type\":\"string\",\"name\":\"uri\",\"internalType\":\"string\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setMintModule4Note\",\"inputs\":[{\"type\":\"tuple\",\"name\":\"vars\",\"internalType\":\"structDataTypes.setMintModule4NoteData\",\"components\":[{\"type\":\"uint256\",\"name\":\"profileId\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"noteId\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"mintModule\",\"internalType\":\"address\"},{\"type\":\"bytes\",\"name\":\"mintModuleInitData\",\"internalType\":\"bytes\"}]}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setNoteUri\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"profileId\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"noteId\",\"internalType\":\"uint256\"},{\"type\":\"string\",\"name\":\"newUri\",\"internalType\":\"string\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setOperator\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"profileId\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"operator\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setPrimaryProfileId\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"profileId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setProfileUri\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"profileId\",\"internalType\":\"uint256\"},{\"type\":\"string\",\"name\":\"newUri\",\"internalType\":\"string\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setSocialToken\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"profileId\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"tokenAddress\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"supportsInterface\",\"inputs\":[{\"type\":\"bytes4\",\"name\":\"interfaceId\",\"internalType\":\"bytes4\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"string\",\"name\":\"\",\"internalType\":\"string\"}],\"name\":\"symbol\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"tokenByIndex\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"index\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"inputs\":[{\"type\":\"address\",\"name\":\"owner\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"index\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"string\",\"name\":\"\",\"internalType\":\"string\"}],\"name\":\"tokenURI\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"profileId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"totalSupply\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"transferFrom\",\"inputs\":[{\"type\":\"address\",\"name\":\"from\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"to\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"unlinkAddress\",\"inputs\":[{\"type\":\"tuple\",\"name\":\"vars\",\"internalType\":\"structDataTypes.linkAddressData\",\"components\":[{\"type\":\"uint256\",\"name\":\"fromProfileId\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"ethAddress\",\"internalType\":\"address\"},{\"type\":\"bytes32\",\"name\":\"linkType\",\"internalType\":\"bytes32\"},{\"type\":\"bytes\",\"name\":\"data\",\"internalType\":\"bytes\"}]}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"unlinkAnyUri\",\"inputs\":[{\"type\":\"tuple\",\"name\":\"vars\",\"internalType\":\"structDataTypes.unlinkAnyUriData\",\"components\":[{\"type\":\"uint256\",\"name\":\"fromProfileId\",\"internalType\":\"uint256\"},{\"type\":\"string\",\"name\":\"toUri\",\"internalType\":\"string\"},{\"type\":\"bytes32\",\"name\":\"linkType\",\"internalType\":\"bytes32\"}]}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"unlinkERC721\",\"inputs\":[{\"type\":\"tuple\",\"name\":\"vars\",\"internalType\":\"structDataTypes.unlinkERC721Data\",\"components\":[{\"type\":\"uint256\",\"name\":\"fromProfileId\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"tokenAddress\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"},{\"type\":\"bytes32\",\"name\":\"linkType\",\"internalType\":\"bytes32\"}]}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"unlinkLinklist\",\"inputs\":[{\"type\":\"tuple\",\"name\":\"vars\",\"internalType\":\"structDataTypes.linkLinklistData\",\"components\":[{\"type\":\"uint256\",\"name\":\"fromProfileId\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"toLinkListId\",\"internalType\":\"uint256\"},{\"type\":\"bytes32\",\"name\":\"linkType\",\"internalType\":\"bytes32\"},{\"type\":\"bytes\",\"name\":\"data\",\"internalType\":\"bytes\"}]}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"unlinkNote\",\"inputs\":[{\"type\":\"tuple\",\"name\":\"vars\",\"internalType\":\"structDataTypes.unlinkNoteData\",\"components\":[{\"type\":\"uint256\",\"name\":\"fromProfileId\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"toProfileId\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"toNoteId\",\"internalType\":\"uint256\"},{\"type\":\"bytes32\",\"name\":\"linkType\",\"internalType\":\"bytes32\"}]}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"unlinkProfile\",\"inputs\":[{\"type\":\"tuple\",\"name\":\"vars\",\"internalType\":\"structDataTypes.unlinkProfileData\",\"components\":[{\"type\":\"uint256\",\"name\":\"fromProfileId\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"toProfileId\",\"internalType\":\"uint256\"},{\"type\":\"bytes32\",\"name\":\"linkType\",\"internalType\":\"bytes32\"}]}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"unlinkProfileLink\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"fromProfileId\",\"internalType\":\"uint256\"},{\"type\":\"tuple\",\"name\":\"linkData\",\"internalType\":\"structDataTypes.ProfileLinkStruct\",\"components\":[{\"type\":\"uint256\",\"name\":\"fromProfileId\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"toProfileId\",\"internalType\":\"uint256\"},{\"type\":\"bytes32\",\"name\":\"linkType\",\"internalType\":\"bytes32\"}]},{\"type\":\"bytes32\",\"name\":\"linkType\",\"internalType\":\"bytes32\"}]}]",
}

// ERC721ABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC721MetaData.ABI instead.
var ERC721ABI = ERC721MetaData.ABI

// ERC721 is an auto generated Go binding around an Ethereum contract.
type ERC721 struct {
	ERC721Caller     // Read-only binding to the contract
	ERC721Transactor // Write-only binding to the contract
	ERC721Filterer   // Log filterer for contract events
}

// ERC721Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC721Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC721Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC721Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC721Session struct {
	Contract     *ERC721           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC721CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC721CallerSession struct {
	Contract *ERC721Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC721TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC721TransactorSession struct {
	Contract     *ERC721Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC721Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC721Raw struct {
	Contract *ERC721 // Generic contract binding to access the raw methods on
}

// ERC721CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC721CallerRaw struct {
	Contract *ERC721Caller // Generic read-only contract binding to access the raw methods on
}

// ERC721TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC721TransactorRaw struct {
	Contract *ERC721Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC721 creates a new instance of ERC721, bound to a specific deployed contract.
func NewERC721(address common.Address, backend bind.ContractBackend) (*ERC721, error) {
	contract, err := bindERC721(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC721{ERC721Caller: ERC721Caller{contract: contract}, ERC721Transactor: ERC721Transactor{contract: contract}, ERC721Filterer: ERC721Filterer{contract: contract}}, nil
}

// NewERC721Caller creates a new read-only instance of ERC721, bound to a specific deployed contract.
func NewERC721Caller(address common.Address, caller bind.ContractCaller) (*ERC721Caller, error) {
	contract, err := bindERC721(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721Caller{contract: contract}, nil
}

// NewERC721Transactor creates a new write-only instance of ERC721, bound to a specific deployed contract.
func NewERC721Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC721Transactor, error) {
	contract, err := bindERC721(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721Transactor{contract: contract}, nil
}

// NewERC721Filterer creates a new log filterer instance of ERC721, bound to a specific deployed contract.
func NewERC721Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC721Filterer, error) {
	contract, err := bindERC721(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC721Filterer{contract: contract}, nil
}

// bindERC721 binds a generic wrapper to an already deployed contract.
func bindERC721(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721 *ERC721Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC721.Contract.ERC721Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721 *ERC721Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721.Contract.ERC721Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721 *ERC721Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721.Contract.ERC721Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721 *ERC721CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC721.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721 *ERC721TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721 *ERC721TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ERC721 *ERC721Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ERC721 *ERC721Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ERC721.Contract.BalanceOf(&_ERC721.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ERC721 *ERC721CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ERC721.Contract.BalanceOf(&_ERC721.CallOpts, owner)
}

// CanCreate is a free data retrieval call binding the contract method 0x7daca6d0.
//
// Solidity: function canCreate(string handle, address account) view returns(bool)
func (_ERC721 *ERC721Caller) CanCreate(opts *bind.CallOpts, handle string, account common.Address) (bool, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "canCreate", handle, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanCreate is a free data retrieval call binding the contract method 0x7daca6d0.
//
// Solidity: function canCreate(string handle, address account) view returns(bool)
func (_ERC721 *ERC721Session) CanCreate(handle string, account common.Address) (bool, error) {
	return _ERC721.Contract.CanCreate(&_ERC721.CallOpts, handle, account)
}

// CanCreate is a free data retrieval call binding the contract method 0x7daca6d0.
//
// Solidity: function canCreate(string handle, address account) view returns(bool)
func (_ERC721 *ERC721CallerSession) CanCreate(handle string, account common.Address) (bool, error) {
	return _ERC721.Contract.CanCreate(&_ERC721.CallOpts, handle, account)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ERC721 *ERC721Caller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ERC721 *ERC721Session) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ERC721.Contract.GetApproved(&_ERC721.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ERC721 *ERC721CallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ERC721.Contract.GetApproved(&_ERC721.CallOpts, tokenId)
}

// GetHandle is a free data retrieval call binding the contract method 0xec81d194.
//
// Solidity: function getHandle(uint256 profileId) view returns(string)
func (_ERC721 *ERC721Caller) GetHandle(opts *bind.CallOpts, profileId *big.Int) (string, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "getHandle", profileId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetHandle is a free data retrieval call binding the contract method 0xec81d194.
//
// Solidity: function getHandle(uint256 profileId) view returns(string)
func (_ERC721 *ERC721Session) GetHandle(profileId *big.Int) (string, error) {
	return _ERC721.Contract.GetHandle(&_ERC721.CallOpts, profileId)
}

// GetHandle is a free data retrieval call binding the contract method 0xec81d194.
//
// Solidity: function getHandle(uint256 profileId) view returns(string)
func (_ERC721 *ERC721CallerSession) GetHandle(profileId *big.Int) (string, error) {
	return _ERC721.Contract.GetHandle(&_ERC721.CallOpts, profileId)
}

// GetLinkModule4Address is a free data retrieval call binding the contract method 0x31b9d08c.
//
// Solidity: function getLinkModule4Address(address account) view returns(address)
func (_ERC721 *ERC721Caller) GetLinkModule4Address(opts *bind.CallOpts, account common.Address) (common.Address, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "getLinkModule4Address", account)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLinkModule4Address is a free data retrieval call binding the contract method 0x31b9d08c.
//
// Solidity: function getLinkModule4Address(address account) view returns(address)
func (_ERC721 *ERC721Session) GetLinkModule4Address(account common.Address) (common.Address, error) {
	return _ERC721.Contract.GetLinkModule4Address(&_ERC721.CallOpts, account)
}

// GetLinkModule4Address is a free data retrieval call binding the contract method 0x31b9d08c.
//
// Solidity: function getLinkModule4Address(address account) view returns(address)
func (_ERC721 *ERC721CallerSession) GetLinkModule4Address(account common.Address) (common.Address, error) {
	return _ERC721.Contract.GetLinkModule4Address(&_ERC721.CallOpts, account)
}

// GetLinkModule4ERC721 is a free data retrieval call binding the contract method 0x2209d145.
//
// Solidity: function getLinkModule4ERC721(address tokenAddress, uint256 tokenId) view returns(address)
func (_ERC721 *ERC721Caller) GetLinkModule4ERC721(opts *bind.CallOpts, tokenAddress common.Address, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "getLinkModule4ERC721", tokenAddress, tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLinkModule4ERC721 is a free data retrieval call binding the contract method 0x2209d145.
//
// Solidity: function getLinkModule4ERC721(address tokenAddress, uint256 tokenId) view returns(address)
func (_ERC721 *ERC721Session) GetLinkModule4ERC721(tokenAddress common.Address, tokenId *big.Int) (common.Address, error) {
	return _ERC721.Contract.GetLinkModule4ERC721(&_ERC721.CallOpts, tokenAddress, tokenId)
}

// GetLinkModule4ERC721 is a free data retrieval call binding the contract method 0x2209d145.
//
// Solidity: function getLinkModule4ERC721(address tokenAddress, uint256 tokenId) view returns(address)
func (_ERC721 *ERC721CallerSession) GetLinkModule4ERC721(tokenAddress common.Address, tokenId *big.Int) (common.Address, error) {
	return _ERC721.Contract.GetLinkModule4ERC721(&_ERC721.CallOpts, tokenAddress, tokenId)
}

// GetLinkModule4Linklist is a free data retrieval call binding the contract method 0xfe9299fb.
//
// Solidity: function getLinkModule4Linklist(uint256 tokenId) view returns(address)
func (_ERC721 *ERC721Caller) GetLinkModule4Linklist(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "getLinkModule4Linklist", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLinkModule4Linklist is a free data retrieval call binding the contract method 0xfe9299fb.
//
// Solidity: function getLinkModule4Linklist(uint256 tokenId) view returns(address)
func (_ERC721 *ERC721Session) GetLinkModule4Linklist(tokenId *big.Int) (common.Address, error) {
	return _ERC721.Contract.GetLinkModule4Linklist(&_ERC721.CallOpts, tokenId)
}

// GetLinkModule4Linklist is a free data retrieval call binding the contract method 0xfe9299fb.
//
// Solidity: function getLinkModule4Linklist(uint256 tokenId) view returns(address)
func (_ERC721 *ERC721CallerSession) GetLinkModule4Linklist(tokenId *big.Int) (common.Address, error) {
	return _ERC721.Contract.GetLinkModule4Linklist(&_ERC721.CallOpts, tokenId)
}

// GetLinklistContract is a free data retrieval call binding the contract method 0xc053f6b8.
//
// Solidity: function getLinklistContract() view returns(address)
func (_ERC721 *ERC721Caller) GetLinklistContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "getLinklistContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLinklistContract is a free data retrieval call binding the contract method 0xc053f6b8.
//
// Solidity: function getLinklistContract() view returns(address)
func (_ERC721 *ERC721Session) GetLinklistContract() (common.Address, error) {
	return _ERC721.Contract.GetLinklistContract(&_ERC721.CallOpts)
}

// GetLinklistContract is a free data retrieval call binding the contract method 0xc053f6b8.
//
// Solidity: function getLinklistContract() view returns(address)
func (_ERC721 *ERC721CallerSession) GetLinklistContract() (common.Address, error) {
	return _ERC721.Contract.GetLinklistContract(&_ERC721.CallOpts)
}

// GetLinklistId is a free data retrieval call binding the contract method 0xd70e10c6.
//
// Solidity: function getLinklistId(uint256 profileId, bytes32 linkType) view returns(uint256)
func (_ERC721 *ERC721Caller) GetLinklistId(opts *bind.CallOpts, profileId *big.Int, linkType [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "getLinklistId", profileId, linkType)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLinklistId is a free data retrieval call binding the contract method 0xd70e10c6.
//
// Solidity: function getLinklistId(uint256 profileId, bytes32 linkType) view returns(uint256)
func (_ERC721 *ERC721Session) GetLinklistId(profileId *big.Int, linkType [32]byte) (*big.Int, error) {
	return _ERC721.Contract.GetLinklistId(&_ERC721.CallOpts, profileId, linkType)
}

// GetLinklistId is a free data retrieval call binding the contract method 0xd70e10c6.
//
// Solidity: function getLinklistId(uint256 profileId, bytes32 linkType) view returns(uint256)
func (_ERC721 *ERC721CallerSession) GetLinklistId(profileId *big.Int, linkType [32]byte) (*big.Int, error) {
	return _ERC721.Contract.GetLinklistId(&_ERC721.CallOpts, profileId, linkType)
}

// GetLinklistType is a free data retrieval call binding the contract method 0x8b4ca06a.
//
// Solidity: function getLinklistType(uint256 linkListId) view returns(bytes32)
func (_ERC721 *ERC721Caller) GetLinklistType(opts *bind.CallOpts, linkListId *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "getLinklistType", linkListId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetLinklistType is a free data retrieval call binding the contract method 0x8b4ca06a.
//
// Solidity: function getLinklistType(uint256 linkListId) view returns(bytes32)
func (_ERC721 *ERC721Session) GetLinklistType(linkListId *big.Int) ([32]byte, error) {
	return _ERC721.Contract.GetLinklistType(&_ERC721.CallOpts, linkListId)
}

// GetLinklistType is a free data retrieval call binding the contract method 0x8b4ca06a.
//
// Solidity: function getLinklistType(uint256 linkListId) view returns(bytes32)
func (_ERC721 *ERC721CallerSession) GetLinklistType(linkListId *big.Int) ([32]byte, error) {
	return _ERC721.Contract.GetLinklistType(&_ERC721.CallOpts, linkListId)
}

// GetLinklistUri is a free data retrieval call binding the contract method 0xdca27135.
//
// Solidity: function getLinklistUri(uint256 tokenId) view returns(string)
func (_ERC721 *ERC721Caller) GetLinklistUri(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "getLinklistUri", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetLinklistUri is a free data retrieval call binding the contract method 0xdca27135.
//
// Solidity: function getLinklistUri(uint256 tokenId) view returns(string)
func (_ERC721 *ERC721Session) GetLinklistUri(tokenId *big.Int) (string, error) {
	return _ERC721.Contract.GetLinklistUri(&_ERC721.CallOpts, tokenId)
}

// GetLinklistUri is a free data retrieval call binding the contract method 0xdca27135.
//
// Solidity: function getLinklistUri(uint256 tokenId) view returns(string)
func (_ERC721 *ERC721CallerSession) GetLinklistUri(tokenId *big.Int) (string, error) {
	return _ERC721.Contract.GetLinklistUri(&_ERC721.CallOpts, tokenId)
}

// GetNote is a free data retrieval call binding the contract method 0xdb491e80.
//
// Solidity: function getNote(uint256 profileId, uint256 noteId) view returns((bytes32,bytes32,string,address,address,address,bool,bool))
func (_ERC721 *ERC721Caller) GetNote(opts *bind.CallOpts, profileId *big.Int, noteId *big.Int) (DataTypesNote, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "getNote", profileId, noteId)

	if err != nil {
		return *new(DataTypesNote), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesNote)).(*DataTypesNote)

	return out0, err

}

// GetNote is a free data retrieval call binding the contract method 0xdb491e80.
//
// Solidity: function getNote(uint256 profileId, uint256 noteId) view returns((bytes32,bytes32,string,address,address,address,bool,bool))
func (_ERC721 *ERC721Session) GetNote(profileId *big.Int, noteId *big.Int) (DataTypesNote, error) {
	return _ERC721.Contract.GetNote(&_ERC721.CallOpts, profileId, noteId)
}

// GetNote is a free data retrieval call binding the contract method 0xdb491e80.
//
// Solidity: function getNote(uint256 profileId, uint256 noteId) view returns((bytes32,bytes32,string,address,address,address,bool,bool))
func (_ERC721 *ERC721CallerSession) GetNote(profileId *big.Int, noteId *big.Int) (DataTypesNote, error) {
	return _ERC721.Contract.GetNote(&_ERC721.CallOpts, profileId, noteId)
}

// GetOperator is a free data retrieval call binding the contract method 0x05f63c8a.
//
// Solidity: function getOperator(uint256 profileId) view returns(address)
func (_ERC721 *ERC721Caller) GetOperator(opts *bind.CallOpts, profileId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "getOperator", profileId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOperator is a free data retrieval call binding the contract method 0x05f63c8a.
//
// Solidity: function getOperator(uint256 profileId) view returns(address)
func (_ERC721 *ERC721Session) GetOperator(profileId *big.Int) (common.Address, error) {
	return _ERC721.Contract.GetOperator(&_ERC721.CallOpts, profileId)
}

// GetOperator is a free data retrieval call binding the contract method 0x05f63c8a.
//
// Solidity: function getOperator(uint256 profileId) view returns(address)
func (_ERC721 *ERC721CallerSession) GetOperator(profileId *big.Int) (common.Address, error) {
	return _ERC721.Contract.GetOperator(&_ERC721.CallOpts, profileId)
}

// GetPrimaryProfileId is a free data retrieval call binding the contract method 0x1bb4d231.
//
// Solidity: function getPrimaryProfileId(address account) view returns(uint256)
func (_ERC721 *ERC721Caller) GetPrimaryProfileId(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "getPrimaryProfileId", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrimaryProfileId is a free data retrieval call binding the contract method 0x1bb4d231.
//
// Solidity: function getPrimaryProfileId(address account) view returns(uint256)
func (_ERC721 *ERC721Session) GetPrimaryProfileId(account common.Address) (*big.Int, error) {
	return _ERC721.Contract.GetPrimaryProfileId(&_ERC721.CallOpts, account)
}

// GetPrimaryProfileId is a free data retrieval call binding the contract method 0x1bb4d231.
//
// Solidity: function getPrimaryProfileId(address account) view returns(uint256)
func (_ERC721 *ERC721CallerSession) GetPrimaryProfileId(account common.Address) (*big.Int, error) {
	return _ERC721.Contract.GetPrimaryProfileId(&_ERC721.CallOpts, account)
}

// GetProfile is a free data retrieval call binding the contract method 0xf08f4f64.
//
// Solidity: function getProfile(uint256 profileId) view returns((uint256,string,string,uint256,address,address))
func (_ERC721 *ERC721Caller) GetProfile(opts *bind.CallOpts, profileId *big.Int) (DataTypesProfile, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "getProfile", profileId)

	if err != nil {
		return *new(DataTypesProfile), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesProfile)).(*DataTypesProfile)

	return out0, err

}

// GetProfile is a free data retrieval call binding the contract method 0xf08f4f64.
//
// Solidity: function getProfile(uint256 profileId) view returns((uint256,string,string,uint256,address,address))
func (_ERC721 *ERC721Session) GetProfile(profileId *big.Int) (DataTypesProfile, error) {
	return _ERC721.Contract.GetProfile(&_ERC721.CallOpts, profileId)
}

// GetProfile is a free data retrieval call binding the contract method 0xf08f4f64.
//
// Solidity: function getProfile(uint256 profileId) view returns((uint256,string,string,uint256,address,address))
func (_ERC721 *ERC721CallerSession) GetProfile(profileId *big.Int) (DataTypesProfile, error) {
	return _ERC721.Contract.GetProfile(&_ERC721.CallOpts, profileId)
}

// GetProfileByHandle is a free data retrieval call binding the contract method 0x0c16de10.
//
// Solidity: function getProfileByHandle(string handle) view returns((uint256,string,string,uint256,address,address))
func (_ERC721 *ERC721Caller) GetProfileByHandle(opts *bind.CallOpts, handle string) (DataTypesProfile, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "getProfileByHandle", handle)

	if err != nil {
		return *new(DataTypesProfile), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesProfile)).(*DataTypesProfile)

	return out0, err

}

// GetProfileByHandle is a free data retrieval call binding the contract method 0x0c16de10.
//
// Solidity: function getProfileByHandle(string handle) view returns((uint256,string,string,uint256,address,address))
func (_ERC721 *ERC721Session) GetProfileByHandle(handle string) (DataTypesProfile, error) {
	return _ERC721.Contract.GetProfileByHandle(&_ERC721.CallOpts, handle)
}

// GetProfileByHandle is a free data retrieval call binding the contract method 0x0c16de10.
//
// Solidity: function getProfileByHandle(string handle) view returns((uint256,string,string,uint256,address,address))
func (_ERC721 *ERC721CallerSession) GetProfileByHandle(handle string) (DataTypesProfile, error) {
	return _ERC721.Contract.GetProfileByHandle(&_ERC721.CallOpts, handle)
}

// GetProfileUri is a free data retrieval call binding the contract method 0xcba4f5cc.
//
// Solidity: function getProfileUri(uint256 profileId) view returns(string)
func (_ERC721 *ERC721Caller) GetProfileUri(opts *bind.CallOpts, profileId *big.Int) (string, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "getProfileUri", profileId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetProfileUri is a free data retrieval call binding the contract method 0xcba4f5cc.
//
// Solidity: function getProfileUri(uint256 profileId) view returns(string)
func (_ERC721 *ERC721Session) GetProfileUri(profileId *big.Int) (string, error) {
	return _ERC721.Contract.GetProfileUri(&_ERC721.CallOpts, profileId)
}

// GetProfileUri is a free data retrieval call binding the contract method 0xcba4f5cc.
//
// Solidity: function getProfileUri(uint256 profileId) view returns(string)
func (_ERC721 *ERC721CallerSession) GetProfileUri(profileId *big.Int) (string, error) {
	return _ERC721.Contract.GetProfileUri(&_ERC721.CallOpts, profileId)
}

// GetRevision is a free data retrieval call binding the contract method 0x1316529d.
//
// Solidity: function getRevision() pure returns(uint256)
func (_ERC721 *ERC721Caller) GetRevision(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "getRevision")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRevision is a free data retrieval call binding the contract method 0x1316529d.
//
// Solidity: function getRevision() pure returns(uint256)
func (_ERC721 *ERC721Session) GetRevision() (*big.Int, error) {
	return _ERC721.Contract.GetRevision(&_ERC721.CallOpts)
}

// GetRevision is a free data retrieval call binding the contract method 0x1316529d.
//
// Solidity: function getRevision() pure returns(uint256)
func (_ERC721 *ERC721CallerSession) GetRevision() (*big.Int, error) {
	return _ERC721.Contract.GetRevision(&_ERC721.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ERC721 *ERC721Caller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ERC721 *ERC721Session) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ERC721.Contract.IsApprovedForAll(&_ERC721.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ERC721 *ERC721CallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ERC721.Contract.IsApprovedForAll(&_ERC721.CallOpts, owner, operator)
}

// IsPrimaryProfile is a free data retrieval call binding the contract method 0xc387c453.
//
// Solidity: function isPrimaryProfile(uint256 profileId) view returns(bool)
func (_ERC721 *ERC721Caller) IsPrimaryProfile(opts *bind.CallOpts, profileId *big.Int) (bool, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "isPrimaryProfile", profileId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPrimaryProfile is a free data retrieval call binding the contract method 0xc387c453.
//
// Solidity: function isPrimaryProfile(uint256 profileId) view returns(bool)
func (_ERC721 *ERC721Session) IsPrimaryProfile(profileId *big.Int) (bool, error) {
	return _ERC721.Contract.IsPrimaryProfile(&_ERC721.CallOpts, profileId)
}

// IsPrimaryProfile is a free data retrieval call binding the contract method 0xc387c453.
//
// Solidity: function isPrimaryProfile(uint256 profileId) view returns(bool)
func (_ERC721 *ERC721CallerSession) IsPrimaryProfile(profileId *big.Int) (bool, error) {
	return _ERC721.Contract.IsPrimaryProfile(&_ERC721.CallOpts, profileId)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC721 *ERC721Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC721 *ERC721Session) Name() (string, error) {
	return _ERC721.Contract.Name(&_ERC721.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC721 *ERC721CallerSession) Name() (string, error) {
	return _ERC721.Contract.Name(&_ERC721.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ERC721 *ERC721Caller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ERC721 *ERC721Session) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ERC721.Contract.OwnerOf(&_ERC721.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ERC721 *ERC721CallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ERC721.Contract.OwnerOf(&_ERC721.CallOpts, tokenId)
}

// Resolver is a free data retrieval call binding the contract method 0x04f3bcec.
//
// Solidity: function resolver() view returns(address)
func (_ERC721 *ERC721Caller) Resolver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "resolver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Resolver is a free data retrieval call binding the contract method 0x04f3bcec.
//
// Solidity: function resolver() view returns(address)
func (_ERC721 *ERC721Session) Resolver() (common.Address, error) {
	return _ERC721.Contract.Resolver(&_ERC721.CallOpts)
}

// Resolver is a free data retrieval call binding the contract method 0x04f3bcec.
//
// Solidity: function resolver() view returns(address)
func (_ERC721 *ERC721CallerSession) Resolver() (common.Address, error) {
	return _ERC721.Contract.Resolver(&_ERC721.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC721 *ERC721Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC721 *ERC721Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC721.Contract.SupportsInterface(&_ERC721.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC721 *ERC721CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC721.Contract.SupportsInterface(&_ERC721.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC721 *ERC721Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC721 *ERC721Session) Symbol() (string, error) {
	return _ERC721.Contract.Symbol(&_ERC721.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC721 *ERC721CallerSession) Symbol() (string, error) {
	return _ERC721.Contract.Symbol(&_ERC721.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_ERC721 *ERC721Caller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_ERC721 *ERC721Session) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _ERC721.Contract.TokenByIndex(&_ERC721.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_ERC721 *ERC721CallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _ERC721.Contract.TokenByIndex(&_ERC721.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_ERC721 *ERC721Caller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_ERC721 *ERC721Session) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _ERC721.Contract.TokenOfOwnerByIndex(&_ERC721.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_ERC721 *ERC721CallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _ERC721.Contract.TokenOfOwnerByIndex(&_ERC721.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 profileId) view returns(string)
func (_ERC721 *ERC721Caller) TokenURI(opts *bind.CallOpts, profileId *big.Int) (string, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "tokenURI", profileId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 profileId) view returns(string)
func (_ERC721 *ERC721Session) TokenURI(profileId *big.Int) (string, error) {
	return _ERC721.Contract.TokenURI(&_ERC721.CallOpts, profileId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 profileId) view returns(string)
func (_ERC721 *ERC721CallerSession) TokenURI(profileId *big.Int) (string, error) {
	return _ERC721.Contract.TokenURI(&_ERC721.CallOpts, profileId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC721 *ERC721Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC721 *ERC721Session) TotalSupply() (*big.Int, error) {
	return _ERC721.Contract.TotalSupply(&_ERC721.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC721 *ERC721CallerSession) TotalSupply() (*big.Int, error) {
	return _ERC721.Contract.TotalSupply(&_ERC721.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ERC721 *ERC721Transactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ERC721 *ERC721Session) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.Approve(&_ERC721.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ERC721 *ERC721TransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.Approve(&_ERC721.TransactOpts, to, tokenId)
}

// AttachLinklist is a paid mutator transaction binding the contract method 0x515d42d2.
//
// Solidity: function attachLinklist(uint256 linklistId, uint256 profileId) returns()
func (_ERC721 *ERC721Transactor) AttachLinklist(opts *bind.TransactOpts, linklistId *big.Int, profileId *big.Int) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "attachLinklist", linklistId, profileId)
}

// AttachLinklist is a paid mutator transaction binding the contract method 0x515d42d2.
//
// Solidity: function attachLinklist(uint256 linklistId, uint256 profileId) returns()
func (_ERC721 *ERC721Session) AttachLinklist(linklistId *big.Int, profileId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.AttachLinklist(&_ERC721.TransactOpts, linklistId, profileId)
}

// AttachLinklist is a paid mutator transaction binding the contract method 0x515d42d2.
//
// Solidity: function attachLinklist(uint256 linklistId, uint256 profileId) returns()
func (_ERC721 *ERC721TransactorSession) AttachLinklist(linklistId *big.Int, profileId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.AttachLinklist(&_ERC721.TransactOpts, linklistId, profileId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_ERC721 *ERC721Transactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_ERC721 *ERC721Session) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.Burn(&_ERC721.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_ERC721 *ERC721TransactorSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.Burn(&_ERC721.TransactOpts, tokenId)
}

// CreateProfile is a paid mutator transaction binding the contract method 0xbd5f69cb.
//
// Solidity: function createProfile((address,string,string,address,bytes) vars) returns()
func (_ERC721 *ERC721Transactor) CreateProfile(opts *bind.TransactOpts, vars DataTypesCreateProfileData) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "createProfile", vars)
}

// CreateProfile is a paid mutator transaction binding the contract method 0xbd5f69cb.
//
// Solidity: function createProfile((address,string,string,address,bytes) vars) returns()
func (_ERC721 *ERC721Session) CreateProfile(vars DataTypesCreateProfileData) (*types.Transaction, error) {
	return _ERC721.Contract.CreateProfile(&_ERC721.TransactOpts, vars)
}

// CreateProfile is a paid mutator transaction binding the contract method 0xbd5f69cb.
//
// Solidity: function createProfile((address,string,string,address,bytes) vars) returns()
func (_ERC721 *ERC721TransactorSession) CreateProfile(vars DataTypesCreateProfileData) (*types.Transaction, error) {
	return _ERC721.Contract.CreateProfile(&_ERC721.TransactOpts, vars)
}

// CreateThenLinkProfile is a paid mutator transaction binding the contract method 0x0ab6beba.
//
// Solidity: function createThenLinkProfile((uint256,address,bytes32) vars) returns()
func (_ERC721 *ERC721Transactor) CreateThenLinkProfile(opts *bind.TransactOpts, vars DataTypescreateThenLinkProfileData) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "createThenLinkProfile", vars)
}

// CreateThenLinkProfile is a paid mutator transaction binding the contract method 0x0ab6beba.
//
// Solidity: function createThenLinkProfile((uint256,address,bytes32) vars) returns()
func (_ERC721 *ERC721Session) CreateThenLinkProfile(vars DataTypescreateThenLinkProfileData) (*types.Transaction, error) {
	return _ERC721.Contract.CreateThenLinkProfile(&_ERC721.TransactOpts, vars)
}

// CreateThenLinkProfile is a paid mutator transaction binding the contract method 0x0ab6beba.
//
// Solidity: function createThenLinkProfile((uint256,address,bytes32) vars) returns()
func (_ERC721 *ERC721TransactorSession) CreateThenLinkProfile(vars DataTypescreateThenLinkProfileData) (*types.Transaction, error) {
	return _ERC721.Contract.CreateThenLinkProfile(&_ERC721.TransactOpts, vars)
}

// DeleteNote is a paid mutator transaction binding the contract method 0xc2a6fe3b.
//
// Solidity: function deleteNote(uint256 profileId, uint256 noteId) returns()
func (_ERC721 *ERC721Transactor) DeleteNote(opts *bind.TransactOpts, profileId *big.Int, noteId *big.Int) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "deleteNote", profileId, noteId)
}

// DeleteNote is a paid mutator transaction binding the contract method 0xc2a6fe3b.
//
// Solidity: function deleteNote(uint256 profileId, uint256 noteId) returns()
func (_ERC721 *ERC721Session) DeleteNote(profileId *big.Int, noteId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.DeleteNote(&_ERC721.TransactOpts, profileId, noteId)
}

// DeleteNote is a paid mutator transaction binding the contract method 0xc2a6fe3b.
//
// Solidity: function deleteNote(uint256 profileId, uint256 noteId) returns()
func (_ERC721 *ERC721TransactorSession) DeleteNote(profileId *big.Int, noteId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.DeleteNote(&_ERC721.TransactOpts, profileId, noteId)
}

// DetachLinklist is a paid mutator transaction binding the contract method 0x7936da0e.
//
// Solidity: function detachLinklist(uint256 linklistId, uint256 profileId) returns()
func (_ERC721 *ERC721Transactor) DetachLinklist(opts *bind.TransactOpts, linklistId *big.Int, profileId *big.Int) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "detachLinklist", linklistId, profileId)
}

// DetachLinklist is a paid mutator transaction binding the contract method 0x7936da0e.
//
// Solidity: function detachLinklist(uint256 linklistId, uint256 profileId) returns()
func (_ERC721 *ERC721Session) DetachLinklist(linklistId *big.Int, profileId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.DetachLinklist(&_ERC721.TransactOpts, linklistId, profileId)
}

// DetachLinklist is a paid mutator transaction binding the contract method 0x7936da0e.
//
// Solidity: function detachLinklist(uint256 linklistId, uint256 profileId) returns()
func (_ERC721 *ERC721TransactorSession) DetachLinklist(linklistId *big.Int, profileId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.DetachLinklist(&_ERC721.TransactOpts, linklistId, profileId)
}

// Initialize is a paid mutator transaction binding the contract method 0xe56f2fe4.
//
// Solidity: function initialize(string _name, string _symbol, address _linklistContract, address _mintNFTImpl, address _periphery, address _resolver) returns()
func (_ERC721 *ERC721Transactor) Initialize(opts *bind.TransactOpts, _name string, _symbol string, _linklistContract common.Address, _mintNFTImpl common.Address, _periphery common.Address, _resolver common.Address) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "initialize", _name, _symbol, _linklistContract, _mintNFTImpl, _periphery, _resolver)
}

// Initialize is a paid mutator transaction binding the contract method 0xe56f2fe4.
//
// Solidity: function initialize(string _name, string _symbol, address _linklistContract, address _mintNFTImpl, address _periphery, address _resolver) returns()
func (_ERC721 *ERC721Session) Initialize(_name string, _symbol string, _linklistContract common.Address, _mintNFTImpl common.Address, _periphery common.Address, _resolver common.Address) (*types.Transaction, error) {
	return _ERC721.Contract.Initialize(&_ERC721.TransactOpts, _name, _symbol, _linklistContract, _mintNFTImpl, _periphery, _resolver)
}

// Initialize is a paid mutator transaction binding the contract method 0xe56f2fe4.
//
// Solidity: function initialize(string _name, string _symbol, address _linklistContract, address _mintNFTImpl, address _periphery, address _resolver) returns()
func (_ERC721 *ERC721TransactorSession) Initialize(_name string, _symbol string, _linklistContract common.Address, _mintNFTImpl common.Address, _periphery common.Address, _resolver common.Address) (*types.Transaction, error) {
	return _ERC721.Contract.Initialize(&_ERC721.TransactOpts, _name, _symbol, _linklistContract, _mintNFTImpl, _periphery, _resolver)
}

// LinkAddress is a paid mutator transaction binding the contract method 0x388f5083.
//
// Solidity: function linkAddress((uint256,address,bytes32,bytes) vars) returns()
func (_ERC721 *ERC721Transactor) LinkAddress(opts *bind.TransactOpts, vars DataTypeslinkAddressData) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "linkAddress", vars)
}

// LinkAddress is a paid mutator transaction binding the contract method 0x388f5083.
//
// Solidity: function linkAddress((uint256,address,bytes32,bytes) vars) returns()
func (_ERC721 *ERC721Session) LinkAddress(vars DataTypeslinkAddressData) (*types.Transaction, error) {
	return _ERC721.Contract.LinkAddress(&_ERC721.TransactOpts, vars)
}

// LinkAddress is a paid mutator transaction binding the contract method 0x388f5083.
//
// Solidity: function linkAddress((uint256,address,bytes32,bytes) vars) returns()
func (_ERC721 *ERC721TransactorSession) LinkAddress(vars DataTypeslinkAddressData) (*types.Transaction, error) {
	return _ERC721.Contract.LinkAddress(&_ERC721.TransactOpts, vars)
}

// LinkAnyUri is a paid mutator transaction binding the contract method 0x5fb88183.
//
// Solidity: function linkAnyUri((uint256,string,bytes32,bytes) vars) returns()
func (_ERC721 *ERC721Transactor) LinkAnyUri(opts *bind.TransactOpts, vars DataTypeslinkAnyUriData) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "linkAnyUri", vars)
}

// LinkAnyUri is a paid mutator transaction binding the contract method 0x5fb88183.
//
// Solidity: function linkAnyUri((uint256,string,bytes32,bytes) vars) returns()
func (_ERC721 *ERC721Session) LinkAnyUri(vars DataTypeslinkAnyUriData) (*types.Transaction, error) {
	return _ERC721.Contract.LinkAnyUri(&_ERC721.TransactOpts, vars)
}

// LinkAnyUri is a paid mutator transaction binding the contract method 0x5fb88183.
//
// Solidity: function linkAnyUri((uint256,string,bytes32,bytes) vars) returns()
func (_ERC721 *ERC721TransactorSession) LinkAnyUri(vars DataTypeslinkAnyUriData) (*types.Transaction, error) {
	return _ERC721.Contract.LinkAnyUri(&_ERC721.TransactOpts, vars)
}

// LinkERC721 is a paid mutator transaction binding the contract method 0xcb8e757e.
//
// Solidity: function linkERC721((uint256,address,uint256,bytes32,bytes) vars) returns()
func (_ERC721 *ERC721Transactor) LinkERC721(opts *bind.TransactOpts, vars DataTypeslinkERC721Data) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "linkERC721", vars)
}

// LinkERC721 is a paid mutator transaction binding the contract method 0xcb8e757e.
//
// Solidity: function linkERC721((uint256,address,uint256,bytes32,bytes) vars) returns()
func (_ERC721 *ERC721Session) LinkERC721(vars DataTypeslinkERC721Data) (*types.Transaction, error) {
	return _ERC721.Contract.LinkERC721(&_ERC721.TransactOpts, vars)
}

// LinkERC721 is a paid mutator transaction binding the contract method 0xcb8e757e.
//
// Solidity: function linkERC721((uint256,address,uint256,bytes32,bytes) vars) returns()
func (_ERC721 *ERC721TransactorSession) LinkERC721(vars DataTypeslinkERC721Data) (*types.Transaction, error) {
	return _ERC721.Contract.LinkERC721(&_ERC721.TransactOpts, vars)
}

// LinkLinklist is a paid mutator transaction binding the contract method 0x9864c307.
//
// Solidity: function linkLinklist((uint256,uint256,bytes32,bytes) vars) returns()
func (_ERC721 *ERC721Transactor) LinkLinklist(opts *bind.TransactOpts, vars DataTypeslinkLinklistData) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "linkLinklist", vars)
}

// LinkLinklist is a paid mutator transaction binding the contract method 0x9864c307.
//
// Solidity: function linkLinklist((uint256,uint256,bytes32,bytes) vars) returns()
func (_ERC721 *ERC721Session) LinkLinklist(vars DataTypeslinkLinklistData) (*types.Transaction, error) {
	return _ERC721.Contract.LinkLinklist(&_ERC721.TransactOpts, vars)
}

// LinkLinklist is a paid mutator transaction binding the contract method 0x9864c307.
//
// Solidity: function linkLinklist((uint256,uint256,bytes32,bytes) vars) returns()
func (_ERC721 *ERC721TransactorSession) LinkLinklist(vars DataTypeslinkLinklistData) (*types.Transaction, error) {
	return _ERC721.Contract.LinkLinklist(&_ERC721.TransactOpts, vars)
}

// LinkNote is a paid mutator transaction binding the contract method 0xb9d32845.
//
// Solidity: function linkNote((uint256,uint256,uint256,bytes32,bytes) vars) returns()
func (_ERC721 *ERC721Transactor) LinkNote(opts *bind.TransactOpts, vars DataTypeslinkNoteData) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "linkNote", vars)
}

// LinkNote is a paid mutator transaction binding the contract method 0xb9d32845.
//
// Solidity: function linkNote((uint256,uint256,uint256,bytes32,bytes) vars) returns()
func (_ERC721 *ERC721Session) LinkNote(vars DataTypeslinkNoteData) (*types.Transaction, error) {
	return _ERC721.Contract.LinkNote(&_ERC721.TransactOpts, vars)
}

// LinkNote is a paid mutator transaction binding the contract method 0xb9d32845.
//
// Solidity: function linkNote((uint256,uint256,uint256,bytes32,bytes) vars) returns()
func (_ERC721 *ERC721TransactorSession) LinkNote(vars DataTypeslinkNoteData) (*types.Transaction, error) {
	return _ERC721.Contract.LinkNote(&_ERC721.TransactOpts, vars)
}

// LinkProfile is a paid mutator transaction binding the contract method 0xa914c76e.
//
// Solidity: function linkProfile((uint256,uint256,bytes32,bytes) vars) returns()
func (_ERC721 *ERC721Transactor) LinkProfile(opts *bind.TransactOpts, vars DataTypeslinkProfileData) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "linkProfile", vars)
}

// LinkProfile is a paid mutator transaction binding the contract method 0xa914c76e.
//
// Solidity: function linkProfile((uint256,uint256,bytes32,bytes) vars) returns()
func (_ERC721 *ERC721Session) LinkProfile(vars DataTypeslinkProfileData) (*types.Transaction, error) {
	return _ERC721.Contract.LinkProfile(&_ERC721.TransactOpts, vars)
}

// LinkProfile is a paid mutator transaction binding the contract method 0xa914c76e.
//
// Solidity: function linkProfile((uint256,uint256,bytes32,bytes) vars) returns()
func (_ERC721 *ERC721TransactorSession) LinkProfile(vars DataTypeslinkProfileData) (*types.Transaction, error) {
	return _ERC721.Contract.LinkProfile(&_ERC721.TransactOpts, vars)
}

// LinkProfileLink is a paid mutator transaction binding the contract method 0xc7009c6b.
//
// Solidity: function linkProfileLink(uint256 fromProfileId, (uint256,uint256,bytes32) linkData, bytes32 linkType) returns()
func (_ERC721 *ERC721Transactor) LinkProfileLink(opts *bind.TransactOpts, fromProfileId *big.Int, linkData DataTypesProfileLinkStruct, linkType [32]byte) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "linkProfileLink", fromProfileId, linkData, linkType)
}

// LinkProfileLink is a paid mutator transaction binding the contract method 0xc7009c6b.
//
// Solidity: function linkProfileLink(uint256 fromProfileId, (uint256,uint256,bytes32) linkData, bytes32 linkType) returns()
func (_ERC721 *ERC721Session) LinkProfileLink(fromProfileId *big.Int, linkData DataTypesProfileLinkStruct, linkType [32]byte) (*types.Transaction, error) {
	return _ERC721.Contract.LinkProfileLink(&_ERC721.TransactOpts, fromProfileId, linkData, linkType)
}

// LinkProfileLink is a paid mutator transaction binding the contract method 0xc7009c6b.
//
// Solidity: function linkProfileLink(uint256 fromProfileId, (uint256,uint256,bytes32) linkData, bytes32 linkType) returns()
func (_ERC721 *ERC721TransactorSession) LinkProfileLink(fromProfileId *big.Int, linkData DataTypesProfileLinkStruct, linkType [32]byte) (*types.Transaction, error) {
	return _ERC721.Contract.LinkProfileLink(&_ERC721.TransactOpts, fromProfileId, linkData, linkType)
}

// LockNote is a paid mutator transaction binding the contract method 0x74f345cf.
//
// Solidity: function lockNote(uint256 profileId, uint256 noteId) returns()
func (_ERC721 *ERC721Transactor) LockNote(opts *bind.TransactOpts, profileId *big.Int, noteId *big.Int) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "lockNote", profileId, noteId)
}

// LockNote is a paid mutator transaction binding the contract method 0x74f345cf.
//
// Solidity: function lockNote(uint256 profileId, uint256 noteId) returns()
func (_ERC721 *ERC721Session) LockNote(profileId *big.Int, noteId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.LockNote(&_ERC721.TransactOpts, profileId, noteId)
}

// LockNote is a paid mutator transaction binding the contract method 0x74f345cf.
//
// Solidity: function lockNote(uint256 profileId, uint256 noteId) returns()
func (_ERC721 *ERC721TransactorSession) LockNote(profileId *big.Int, noteId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.LockNote(&_ERC721.TransactOpts, profileId, noteId)
}

// MintNote is a paid mutator transaction binding the contract method 0xa7ccb4bf.
//
// Solidity: function mintNote((uint256,uint256,address,bytes) vars) returns(uint256)
func (_ERC721 *ERC721Transactor) MintNote(opts *bind.TransactOpts, vars DataTypesMintNoteData) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "mintNote", vars)
}

// MintNote is a paid mutator transaction binding the contract method 0xa7ccb4bf.
//
// Solidity: function mintNote((uint256,uint256,address,bytes) vars) returns(uint256)
func (_ERC721 *ERC721Session) MintNote(vars DataTypesMintNoteData) (*types.Transaction, error) {
	return _ERC721.Contract.MintNote(&_ERC721.TransactOpts, vars)
}

// MintNote is a paid mutator transaction binding the contract method 0xa7ccb4bf.
//
// Solidity: function mintNote((uint256,uint256,address,bytes) vars) returns(uint256)
func (_ERC721 *ERC721TransactorSession) MintNote(vars DataTypesMintNoteData) (*types.Transaction, error) {
	return _ERC721.Contract.MintNote(&_ERC721.TransactOpts, vars)
}

// PostNote is a paid mutator transaction binding the contract method 0x29c301c2.
//
// Solidity: function postNote((uint256,string,address,bytes,address,bytes,bool) vars) returns(uint256)
func (_ERC721 *ERC721Transactor) PostNote(opts *bind.TransactOpts, vars DataTypesPostNoteData) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "postNote", vars)
}

// PostNote is a paid mutator transaction binding the contract method 0x29c301c2.
//
// Solidity: function postNote((uint256,string,address,bytes,address,bytes,bool) vars) returns(uint256)
func (_ERC721 *ERC721Session) PostNote(vars DataTypesPostNoteData) (*types.Transaction, error) {
	return _ERC721.Contract.PostNote(&_ERC721.TransactOpts, vars)
}

// PostNote is a paid mutator transaction binding the contract method 0x29c301c2.
//
// Solidity: function postNote((uint256,string,address,bytes,address,bytes,bool) vars) returns(uint256)
func (_ERC721 *ERC721TransactorSession) PostNote(vars DataTypesPostNoteData) (*types.Transaction, error) {
	return _ERC721.Contract.PostNote(&_ERC721.TransactOpts, vars)
}

// PostNote4Address is a paid mutator transaction binding the contract method 0x92f7070b.
//
// Solidity: function postNote4Address((uint256,string,address,bytes,address,bytes,bool) noteData, address ethAddress) returns(uint256)
func (_ERC721 *ERC721Transactor) PostNote4Address(opts *bind.TransactOpts, noteData DataTypesPostNoteData, ethAddress common.Address) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "postNote4Address", noteData, ethAddress)
}

// PostNote4Address is a paid mutator transaction binding the contract method 0x92f7070b.
//
// Solidity: function postNote4Address((uint256,string,address,bytes,address,bytes,bool) noteData, address ethAddress) returns(uint256)
func (_ERC721 *ERC721Session) PostNote4Address(noteData DataTypesPostNoteData, ethAddress common.Address) (*types.Transaction, error) {
	return _ERC721.Contract.PostNote4Address(&_ERC721.TransactOpts, noteData, ethAddress)
}

// PostNote4Address is a paid mutator transaction binding the contract method 0x92f7070b.
//
// Solidity: function postNote4Address((uint256,string,address,bytes,address,bytes,bool) noteData, address ethAddress) returns(uint256)
func (_ERC721 *ERC721TransactorSession) PostNote4Address(noteData DataTypesPostNoteData, ethAddress common.Address) (*types.Transaction, error) {
	return _ERC721.Contract.PostNote4Address(&_ERC721.TransactOpts, noteData, ethAddress)
}

// PostNote4AnyUri is a paid mutator transaction binding the contract method 0xf316bacd.
//
// Solidity: function postNote4AnyUri((uint256,string,address,bytes,address,bytes,bool) postNoteData, string uri) returns(uint256)
func (_ERC721 *ERC721Transactor) PostNote4AnyUri(opts *bind.TransactOpts, postNoteData DataTypesPostNoteData, uri string) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "postNote4AnyUri", postNoteData, uri)
}

// PostNote4AnyUri is a paid mutator transaction binding the contract method 0xf316bacd.
//
// Solidity: function postNote4AnyUri((uint256,string,address,bytes,address,bytes,bool) postNoteData, string uri) returns(uint256)
func (_ERC721 *ERC721Session) PostNote4AnyUri(postNoteData DataTypesPostNoteData, uri string) (*types.Transaction, error) {
	return _ERC721.Contract.PostNote4AnyUri(&_ERC721.TransactOpts, postNoteData, uri)
}

// PostNote4AnyUri is a paid mutator transaction binding the contract method 0xf316bacd.
//
// Solidity: function postNote4AnyUri((uint256,string,address,bytes,address,bytes,bool) postNoteData, string uri) returns(uint256)
func (_ERC721 *ERC721TransactorSession) PostNote4AnyUri(postNoteData DataTypesPostNoteData, uri string) (*types.Transaction, error) {
	return _ERC721.Contract.PostNote4AnyUri(&_ERC721.TransactOpts, postNoteData, uri)
}

// PostNote4ERC721 is a paid mutator transaction binding the contract method 0x327b2a03.
//
// Solidity: function postNote4ERC721((uint256,string,address,bytes,address,bytes,bool) postNoteData, (address,uint256) erc721) returns(uint256)
func (_ERC721 *ERC721Transactor) PostNote4ERC721(opts *bind.TransactOpts, postNoteData DataTypesPostNoteData, erc721 DataTypesERC721Struct) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "postNote4ERC721", postNoteData, erc721)
}

// PostNote4ERC721 is a paid mutator transaction binding the contract method 0x327b2a03.
//
// Solidity: function postNote4ERC721((uint256,string,address,bytes,address,bytes,bool) postNoteData, (address,uint256) erc721) returns(uint256)
func (_ERC721 *ERC721Session) PostNote4ERC721(postNoteData DataTypesPostNoteData, erc721 DataTypesERC721Struct) (*types.Transaction, error) {
	return _ERC721.Contract.PostNote4ERC721(&_ERC721.TransactOpts, postNoteData, erc721)
}

// PostNote4ERC721 is a paid mutator transaction binding the contract method 0x327b2a03.
//
// Solidity: function postNote4ERC721((uint256,string,address,bytes,address,bytes,bool) postNoteData, (address,uint256) erc721) returns(uint256)
func (_ERC721 *ERC721TransactorSession) PostNote4ERC721(postNoteData DataTypesPostNoteData, erc721 DataTypesERC721Struct) (*types.Transaction, error) {
	return _ERC721.Contract.PostNote4ERC721(&_ERC721.TransactOpts, postNoteData, erc721)
}

// PostNote4Linklist is a paid mutator transaction binding the contract method 0x44b82a24.
//
// Solidity: function postNote4Linklist((uint256,string,address,bytes,address,bytes,bool) noteData, uint256 toLinklistId) returns(uint256)
func (_ERC721 *ERC721Transactor) PostNote4Linklist(opts *bind.TransactOpts, noteData DataTypesPostNoteData, toLinklistId *big.Int) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "postNote4Linklist", noteData, toLinklistId)
}

// PostNote4Linklist is a paid mutator transaction binding the contract method 0x44b82a24.
//
// Solidity: function postNote4Linklist((uint256,string,address,bytes,address,bytes,bool) noteData, uint256 toLinklistId) returns(uint256)
func (_ERC721 *ERC721Session) PostNote4Linklist(noteData DataTypesPostNoteData, toLinklistId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.PostNote4Linklist(&_ERC721.TransactOpts, noteData, toLinklistId)
}

// PostNote4Linklist is a paid mutator transaction binding the contract method 0x44b82a24.
//
// Solidity: function postNote4Linklist((uint256,string,address,bytes,address,bytes,bool) noteData, uint256 toLinklistId) returns(uint256)
func (_ERC721 *ERC721TransactorSession) PostNote4Linklist(noteData DataTypesPostNoteData, toLinklistId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.PostNote4Linklist(&_ERC721.TransactOpts, noteData, toLinklistId)
}

// PostNote4Note is a paid mutator transaction binding the contract method 0x9a4dec18.
//
// Solidity: function postNote4Note((uint256,string,address,bytes,address,bytes,bool) postNoteData, (uint256,uint256) note) returns(uint256)
func (_ERC721 *ERC721Transactor) PostNote4Note(opts *bind.TransactOpts, postNoteData DataTypesPostNoteData, note DataTypesNoteStruct) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "postNote4Note", postNoteData, note)
}

// PostNote4Note is a paid mutator transaction binding the contract method 0x9a4dec18.
//
// Solidity: function postNote4Note((uint256,string,address,bytes,address,bytes,bool) postNoteData, (uint256,uint256) note) returns(uint256)
func (_ERC721 *ERC721Session) PostNote4Note(postNoteData DataTypesPostNoteData, note DataTypesNoteStruct) (*types.Transaction, error) {
	return _ERC721.Contract.PostNote4Note(&_ERC721.TransactOpts, postNoteData, note)
}

// PostNote4Note is a paid mutator transaction binding the contract method 0x9a4dec18.
//
// Solidity: function postNote4Note((uint256,string,address,bytes,address,bytes,bool) postNoteData, (uint256,uint256) note) returns(uint256)
func (_ERC721 *ERC721TransactorSession) PostNote4Note(postNoteData DataTypesPostNoteData, note DataTypesNoteStruct) (*types.Transaction, error) {
	return _ERC721.Contract.PostNote4Note(&_ERC721.TransactOpts, postNoteData, note)
}

// PostNote4Profile is a paid mutator transaction binding the contract method 0x2ff5b07e.
//
// Solidity: function postNote4Profile((uint256,string,address,bytes,address,bytes,bool) postNoteData, uint256 toProfileId) returns(uint256)
func (_ERC721 *ERC721Transactor) PostNote4Profile(opts *bind.TransactOpts, postNoteData DataTypesPostNoteData, toProfileId *big.Int) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "postNote4Profile", postNoteData, toProfileId)
}

// PostNote4Profile is a paid mutator transaction binding the contract method 0x2ff5b07e.
//
// Solidity: function postNote4Profile((uint256,string,address,bytes,address,bytes,bool) postNoteData, uint256 toProfileId) returns(uint256)
func (_ERC721 *ERC721Session) PostNote4Profile(postNoteData DataTypesPostNoteData, toProfileId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.PostNote4Profile(&_ERC721.TransactOpts, postNoteData, toProfileId)
}

// PostNote4Profile is a paid mutator transaction binding the contract method 0x2ff5b07e.
//
// Solidity: function postNote4Profile((uint256,string,address,bytes,address,bytes,bool) postNoteData, uint256 toProfileId) returns(uint256)
func (_ERC721 *ERC721TransactorSession) PostNote4Profile(postNoteData DataTypesPostNoteData, toProfileId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.PostNote4Profile(&_ERC721.TransactOpts, postNoteData, toProfileId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721 *ERC721Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721 *ERC721Session) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.SafeTransferFrom(&_ERC721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721 *ERC721TransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.SafeTransferFrom(&_ERC721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_ERC721 *ERC721Transactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_ERC721 *ERC721Session) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721.Contract.SafeTransferFrom0(&_ERC721.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_ERC721 *ERC721TransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721.Contract.SafeTransferFrom0(&_ERC721.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ERC721 *ERC721Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ERC721 *ERC721Session) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ERC721.Contract.SetApprovalForAll(&_ERC721.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ERC721 *ERC721TransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ERC721.Contract.SetApprovalForAll(&_ERC721.TransactOpts, operator, approved)
}

// SetHandle is a paid mutator transaction binding the contract method 0xa6e6178d.
//
// Solidity: function setHandle(uint256 profileId, string newHandle) returns()
func (_ERC721 *ERC721Transactor) SetHandle(opts *bind.TransactOpts, profileId *big.Int, newHandle string) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "setHandle", profileId, newHandle)
}

// SetHandle is a paid mutator transaction binding the contract method 0xa6e6178d.
//
// Solidity: function setHandle(uint256 profileId, string newHandle) returns()
func (_ERC721 *ERC721Session) SetHandle(profileId *big.Int, newHandle string) (*types.Transaction, error) {
	return _ERC721.Contract.SetHandle(&_ERC721.TransactOpts, profileId, newHandle)
}

// SetHandle is a paid mutator transaction binding the contract method 0xa6e6178d.
//
// Solidity: function setHandle(uint256 profileId, string newHandle) returns()
func (_ERC721 *ERC721TransactorSession) SetHandle(profileId *big.Int, newHandle string) (*types.Transaction, error) {
	return _ERC721.Contract.SetHandle(&_ERC721.TransactOpts, profileId, newHandle)
}

// SetLinkModule4Address is a paid mutator transaction binding the contract method 0x08cb68ff.
//
// Solidity: function setLinkModule4Address((address,address,bytes) vars) returns()
func (_ERC721 *ERC721Transactor) SetLinkModule4Address(opts *bind.TransactOpts, vars DataTypessetLinkModule4AddressData) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "setLinkModule4Address", vars)
}

// SetLinkModule4Address is a paid mutator transaction binding the contract method 0x08cb68ff.
//
// Solidity: function setLinkModule4Address((address,address,bytes) vars) returns()
func (_ERC721 *ERC721Session) SetLinkModule4Address(vars DataTypessetLinkModule4AddressData) (*types.Transaction, error) {
	return _ERC721.Contract.SetLinkModule4Address(&_ERC721.TransactOpts, vars)
}

// SetLinkModule4Address is a paid mutator transaction binding the contract method 0x08cb68ff.
//
// Solidity: function setLinkModule4Address((address,address,bytes) vars) returns()
func (_ERC721 *ERC721TransactorSession) SetLinkModule4Address(vars DataTypessetLinkModule4AddressData) (*types.Transaction, error) {
	return _ERC721.Contract.SetLinkModule4Address(&_ERC721.TransactOpts, vars)
}

// SetLinkModule4ERC721 is a paid mutator transaction binding the contract method 0x69492c97.
//
// Solidity: function setLinkModule4ERC721((address,uint256,address,bytes) vars) returns()
func (_ERC721 *ERC721Transactor) SetLinkModule4ERC721(opts *bind.TransactOpts, vars DataTypessetLinkModule4ERC721Data) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "setLinkModule4ERC721", vars)
}

// SetLinkModule4ERC721 is a paid mutator transaction binding the contract method 0x69492c97.
//
// Solidity: function setLinkModule4ERC721((address,uint256,address,bytes) vars) returns()
func (_ERC721 *ERC721Session) SetLinkModule4ERC721(vars DataTypessetLinkModule4ERC721Data) (*types.Transaction, error) {
	return _ERC721.Contract.SetLinkModule4ERC721(&_ERC721.TransactOpts, vars)
}

// SetLinkModule4ERC721 is a paid mutator transaction binding the contract method 0x69492c97.
//
// Solidity: function setLinkModule4ERC721((address,uint256,address,bytes) vars) returns()
func (_ERC721 *ERC721TransactorSession) SetLinkModule4ERC721(vars DataTypessetLinkModule4ERC721Data) (*types.Transaction, error) {
	return _ERC721.Contract.SetLinkModule4ERC721(&_ERC721.TransactOpts, vars)
}

// SetLinkModule4Linklist is a paid mutator transaction binding the contract method 0x0c4dd5f2.
//
// Solidity: function setLinkModule4Linklist((uint256,address,bytes) vars) returns()
func (_ERC721 *ERC721Transactor) SetLinkModule4Linklist(opts *bind.TransactOpts, vars DataTypessetLinkModule4LinklistData) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "setLinkModule4Linklist", vars)
}

// SetLinkModule4Linklist is a paid mutator transaction binding the contract method 0x0c4dd5f2.
//
// Solidity: function setLinkModule4Linklist((uint256,address,bytes) vars) returns()
func (_ERC721 *ERC721Session) SetLinkModule4Linklist(vars DataTypessetLinkModule4LinklistData) (*types.Transaction, error) {
	return _ERC721.Contract.SetLinkModule4Linklist(&_ERC721.TransactOpts, vars)
}

// SetLinkModule4Linklist is a paid mutator transaction binding the contract method 0x0c4dd5f2.
//
// Solidity: function setLinkModule4Linklist((uint256,address,bytes) vars) returns()
func (_ERC721 *ERC721TransactorSession) SetLinkModule4Linklist(vars DataTypessetLinkModule4LinklistData) (*types.Transaction, error) {
	return _ERC721.Contract.SetLinkModule4Linklist(&_ERC721.TransactOpts, vars)
}

// SetLinkModule4Note is a paid mutator transaction binding the contract method 0xdb8c198d.
//
// Solidity: function setLinkModule4Note((uint256,uint256,address,bytes) vars) returns()
func (_ERC721 *ERC721Transactor) SetLinkModule4Note(opts *bind.TransactOpts, vars DataTypessetLinkModule4NoteData) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "setLinkModule4Note", vars)
}

// SetLinkModule4Note is a paid mutator transaction binding the contract method 0xdb8c198d.
//
// Solidity: function setLinkModule4Note((uint256,uint256,address,bytes) vars) returns()
func (_ERC721 *ERC721Session) SetLinkModule4Note(vars DataTypessetLinkModule4NoteData) (*types.Transaction, error) {
	return _ERC721.Contract.SetLinkModule4Note(&_ERC721.TransactOpts, vars)
}

// SetLinkModule4Note is a paid mutator transaction binding the contract method 0xdb8c198d.
//
// Solidity: function setLinkModule4Note((uint256,uint256,address,bytes) vars) returns()
func (_ERC721 *ERC721TransactorSession) SetLinkModule4Note(vars DataTypessetLinkModule4NoteData) (*types.Transaction, error) {
	return _ERC721.Contract.SetLinkModule4Note(&_ERC721.TransactOpts, vars)
}

// SetLinkModule4Profile is a paid mutator transaction binding the contract method 0x5b507cfd.
//
// Solidity: function setLinkModule4Profile((uint256,address,bytes) vars) returns()
func (_ERC721 *ERC721Transactor) SetLinkModule4Profile(opts *bind.TransactOpts, vars DataTypessetLinkModule4ProfileData) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "setLinkModule4Profile", vars)
}

// SetLinkModule4Profile is a paid mutator transaction binding the contract method 0x5b507cfd.
//
// Solidity: function setLinkModule4Profile((uint256,address,bytes) vars) returns()
func (_ERC721 *ERC721Session) SetLinkModule4Profile(vars DataTypessetLinkModule4ProfileData) (*types.Transaction, error) {
	return _ERC721.Contract.SetLinkModule4Profile(&_ERC721.TransactOpts, vars)
}

// SetLinkModule4Profile is a paid mutator transaction binding the contract method 0x5b507cfd.
//
// Solidity: function setLinkModule4Profile((uint256,address,bytes) vars) returns()
func (_ERC721 *ERC721TransactorSession) SetLinkModule4Profile(vars DataTypessetLinkModule4ProfileData) (*types.Transaction, error) {
	return _ERC721.Contract.SetLinkModule4Profile(&_ERC721.TransactOpts, vars)
}

// SetLinklistUri is a paid mutator transaction binding the contract method 0x33f06ee6.
//
// Solidity: function setLinklistUri(uint256 linklistId, string uri) returns()
func (_ERC721 *ERC721Transactor) SetLinklistUri(opts *bind.TransactOpts, linklistId *big.Int, uri string) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "setLinklistUri", linklistId, uri)
}

// SetLinklistUri is a paid mutator transaction binding the contract method 0x33f06ee6.
//
// Solidity: function setLinklistUri(uint256 linklistId, string uri) returns()
func (_ERC721 *ERC721Session) SetLinklistUri(linklistId *big.Int, uri string) (*types.Transaction, error) {
	return _ERC721.Contract.SetLinklistUri(&_ERC721.TransactOpts, linklistId, uri)
}

// SetLinklistUri is a paid mutator transaction binding the contract method 0x33f06ee6.
//
// Solidity: function setLinklistUri(uint256 linklistId, string uri) returns()
func (_ERC721 *ERC721TransactorSession) SetLinklistUri(linklistId *big.Int, uri string) (*types.Transaction, error) {
	return _ERC721.Contract.SetLinklistUri(&_ERC721.TransactOpts, linklistId, uri)
}

// SetMintModule4Note is a paid mutator transaction binding the contract method 0xd23b320b.
//
// Solidity: function setMintModule4Note((uint256,uint256,address,bytes) vars) returns()
func (_ERC721 *ERC721Transactor) SetMintModule4Note(opts *bind.TransactOpts, vars DataTypessetMintModule4NoteData) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "setMintModule4Note", vars)
}

// SetMintModule4Note is a paid mutator transaction binding the contract method 0xd23b320b.
//
// Solidity: function setMintModule4Note((uint256,uint256,address,bytes) vars) returns()
func (_ERC721 *ERC721Session) SetMintModule4Note(vars DataTypessetMintModule4NoteData) (*types.Transaction, error) {
	return _ERC721.Contract.SetMintModule4Note(&_ERC721.TransactOpts, vars)
}

// SetMintModule4Note is a paid mutator transaction binding the contract method 0xd23b320b.
//
// Solidity: function setMintModule4Note((uint256,uint256,address,bytes) vars) returns()
func (_ERC721 *ERC721TransactorSession) SetMintModule4Note(vars DataTypessetMintModule4NoteData) (*types.Transaction, error) {
	return _ERC721.Contract.SetMintModule4Note(&_ERC721.TransactOpts, vars)
}

// SetNoteUri is a paid mutator transaction binding the contract method 0x628b644a.
//
// Solidity: function setNoteUri(uint256 profileId, uint256 noteId, string newUri) returns()
func (_ERC721 *ERC721Transactor) SetNoteUri(opts *bind.TransactOpts, profileId *big.Int, noteId *big.Int, newUri string) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "setNoteUri", profileId, noteId, newUri)
}

// SetNoteUri is a paid mutator transaction binding the contract method 0x628b644a.
//
// Solidity: function setNoteUri(uint256 profileId, uint256 noteId, string newUri) returns()
func (_ERC721 *ERC721Session) SetNoteUri(profileId *big.Int, noteId *big.Int, newUri string) (*types.Transaction, error) {
	return _ERC721.Contract.SetNoteUri(&_ERC721.TransactOpts, profileId, noteId, newUri)
}

// SetNoteUri is a paid mutator transaction binding the contract method 0x628b644a.
//
// Solidity: function setNoteUri(uint256 profileId, uint256 noteId, string newUri) returns()
func (_ERC721 *ERC721TransactorSession) SetNoteUri(profileId *big.Int, noteId *big.Int, newUri string) (*types.Transaction, error) {
	return _ERC721.Contract.SetNoteUri(&_ERC721.TransactOpts, profileId, noteId, newUri)
}

// SetOperator is a paid mutator transaction binding the contract method 0xe7a1c1c0.
//
// Solidity: function setOperator(uint256 profileId, address operator) returns()
func (_ERC721 *ERC721Transactor) SetOperator(opts *bind.TransactOpts, profileId *big.Int, operator common.Address) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "setOperator", profileId, operator)
}

// SetOperator is a paid mutator transaction binding the contract method 0xe7a1c1c0.
//
// Solidity: function setOperator(uint256 profileId, address operator) returns()
func (_ERC721 *ERC721Session) SetOperator(profileId *big.Int, operator common.Address) (*types.Transaction, error) {
	return _ERC721.Contract.SetOperator(&_ERC721.TransactOpts, profileId, operator)
}

// SetOperator is a paid mutator transaction binding the contract method 0xe7a1c1c0.
//
// Solidity: function setOperator(uint256 profileId, address operator) returns()
func (_ERC721 *ERC721TransactorSession) SetOperator(profileId *big.Int, operator common.Address) (*types.Transaction, error) {
	return _ERC721.Contract.SetOperator(&_ERC721.TransactOpts, profileId, operator)
}

// SetPrimaryProfileId is a paid mutator transaction binding the contract method 0x295cb43e.
//
// Solidity: function setPrimaryProfileId(uint256 profileId) returns()
func (_ERC721 *ERC721Transactor) SetPrimaryProfileId(opts *bind.TransactOpts, profileId *big.Int) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "setPrimaryProfileId", profileId)
}

// SetPrimaryProfileId is a paid mutator transaction binding the contract method 0x295cb43e.
//
// Solidity: function setPrimaryProfileId(uint256 profileId) returns()
func (_ERC721 *ERC721Session) SetPrimaryProfileId(profileId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.SetPrimaryProfileId(&_ERC721.TransactOpts, profileId)
}

// SetPrimaryProfileId is a paid mutator transaction binding the contract method 0x295cb43e.
//
// Solidity: function setPrimaryProfileId(uint256 profileId) returns()
func (_ERC721 *ERC721TransactorSession) SetPrimaryProfileId(profileId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.SetPrimaryProfileId(&_ERC721.TransactOpts, profileId)
}

// SetProfileUri is a paid mutator transaction binding the contract method 0x7c392b51.
//
// Solidity: function setProfileUri(uint256 profileId, string newUri) returns()
func (_ERC721 *ERC721Transactor) SetProfileUri(opts *bind.TransactOpts, profileId *big.Int, newUri string) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "setProfileUri", profileId, newUri)
}

// SetProfileUri is a paid mutator transaction binding the contract method 0x7c392b51.
//
// Solidity: function setProfileUri(uint256 profileId, string newUri) returns()
func (_ERC721 *ERC721Session) SetProfileUri(profileId *big.Int, newUri string) (*types.Transaction, error) {
	return _ERC721.Contract.SetProfileUri(&_ERC721.TransactOpts, profileId, newUri)
}

// SetProfileUri is a paid mutator transaction binding the contract method 0x7c392b51.
//
// Solidity: function setProfileUri(uint256 profileId, string newUri) returns()
func (_ERC721 *ERC721TransactorSession) SetProfileUri(profileId *big.Int, newUri string) (*types.Transaction, error) {
	return _ERC721.Contract.SetProfileUri(&_ERC721.TransactOpts, profileId, newUri)
}

// SetSocialToken is a paid mutator transaction binding the contract method 0x95d9fa7d.
//
// Solidity: function setSocialToken(uint256 profileId, address tokenAddress) returns()
func (_ERC721 *ERC721Transactor) SetSocialToken(opts *bind.TransactOpts, profileId *big.Int, tokenAddress common.Address) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "setSocialToken", profileId, tokenAddress)
}

// SetSocialToken is a paid mutator transaction binding the contract method 0x95d9fa7d.
//
// Solidity: function setSocialToken(uint256 profileId, address tokenAddress) returns()
func (_ERC721 *ERC721Session) SetSocialToken(profileId *big.Int, tokenAddress common.Address) (*types.Transaction, error) {
	return _ERC721.Contract.SetSocialToken(&_ERC721.TransactOpts, profileId, tokenAddress)
}

// SetSocialToken is a paid mutator transaction binding the contract method 0x95d9fa7d.
//
// Solidity: function setSocialToken(uint256 profileId, address tokenAddress) returns()
func (_ERC721 *ERC721TransactorSession) SetSocialToken(profileId *big.Int, tokenAddress common.Address) (*types.Transaction, error) {
	return _ERC721.Contract.SetSocialToken(&_ERC721.TransactOpts, profileId, tokenAddress)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721 *ERC721Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721 *ERC721Session) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.TransferFrom(&_ERC721.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721 *ERC721TransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.TransferFrom(&_ERC721.TransactOpts, from, to, tokenId)
}

// UnlinkAddress is a paid mutator transaction binding the contract method 0xb9ad7c7f.
//
// Solidity: function unlinkAddress((uint256,address,bytes32,bytes) vars) returns()
func (_ERC721 *ERC721Transactor) UnlinkAddress(opts *bind.TransactOpts, vars DataTypeslinkAddressData) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "unlinkAddress", vars)
}

// UnlinkAddress is a paid mutator transaction binding the contract method 0xb9ad7c7f.
//
// Solidity: function unlinkAddress((uint256,address,bytes32,bytes) vars) returns()
func (_ERC721 *ERC721Session) UnlinkAddress(vars DataTypeslinkAddressData) (*types.Transaction, error) {
	return _ERC721.Contract.UnlinkAddress(&_ERC721.TransactOpts, vars)
}

// UnlinkAddress is a paid mutator transaction binding the contract method 0xb9ad7c7f.
//
// Solidity: function unlinkAddress((uint256,address,bytes32,bytes) vars) returns()
func (_ERC721 *ERC721TransactorSession) UnlinkAddress(vars DataTypeslinkAddressData) (*types.Transaction, error) {
	return _ERC721.Contract.UnlinkAddress(&_ERC721.TransactOpts, vars)
}

// UnlinkAnyUri is a paid mutator transaction binding the contract method 0xef0828ab.
//
// Solidity: function unlinkAnyUri((uint256,string,bytes32) vars) returns()
func (_ERC721 *ERC721Transactor) UnlinkAnyUri(opts *bind.TransactOpts, vars DataTypesunlinkAnyUriData) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "unlinkAnyUri", vars)
}

// UnlinkAnyUri is a paid mutator transaction binding the contract method 0xef0828ab.
//
// Solidity: function unlinkAnyUri((uint256,string,bytes32) vars) returns()
func (_ERC721 *ERC721Session) UnlinkAnyUri(vars DataTypesunlinkAnyUriData) (*types.Transaction, error) {
	return _ERC721.Contract.UnlinkAnyUri(&_ERC721.TransactOpts, vars)
}

// UnlinkAnyUri is a paid mutator transaction binding the contract method 0xef0828ab.
//
// Solidity: function unlinkAnyUri((uint256,string,bytes32) vars) returns()
func (_ERC721 *ERC721TransactorSession) UnlinkAnyUri(vars DataTypesunlinkAnyUriData) (*types.Transaction, error) {
	return _ERC721.Contract.UnlinkAnyUri(&_ERC721.TransactOpts, vars)
}

// UnlinkERC721 is a paid mutator transaction binding the contract method 0x867884e6.
//
// Solidity: function unlinkERC721((uint256,address,uint256,bytes32) vars) returns()
func (_ERC721 *ERC721Transactor) UnlinkERC721(opts *bind.TransactOpts, vars DataTypesunlinkERC721Data) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "unlinkERC721", vars)
}

// UnlinkERC721 is a paid mutator transaction binding the contract method 0x867884e6.
//
// Solidity: function unlinkERC721((uint256,address,uint256,bytes32) vars) returns()
func (_ERC721 *ERC721Session) UnlinkERC721(vars DataTypesunlinkERC721Data) (*types.Transaction, error) {
	return _ERC721.Contract.UnlinkERC721(&_ERC721.TransactOpts, vars)
}

// UnlinkERC721 is a paid mutator transaction binding the contract method 0x867884e6.
//
// Solidity: function unlinkERC721((uint256,address,uint256,bytes32) vars) returns()
func (_ERC721 *ERC721TransactorSession) UnlinkERC721(vars DataTypesunlinkERC721Data) (*types.Transaction, error) {
	return _ERC721.Contract.UnlinkERC721(&_ERC721.TransactOpts, vars)
}

// UnlinkLinklist is a paid mutator transaction binding the contract method 0x86833c88.
//
// Solidity: function unlinkLinklist((uint256,uint256,bytes32,bytes) vars) returns()
func (_ERC721 *ERC721Transactor) UnlinkLinklist(opts *bind.TransactOpts, vars DataTypeslinkLinklistData) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "unlinkLinklist", vars)
}

// UnlinkLinklist is a paid mutator transaction binding the contract method 0x86833c88.
//
// Solidity: function unlinkLinklist((uint256,uint256,bytes32,bytes) vars) returns()
func (_ERC721 *ERC721Session) UnlinkLinklist(vars DataTypeslinkLinklistData) (*types.Transaction, error) {
	return _ERC721.Contract.UnlinkLinklist(&_ERC721.TransactOpts, vars)
}

// UnlinkLinklist is a paid mutator transaction binding the contract method 0x86833c88.
//
// Solidity: function unlinkLinklist((uint256,uint256,bytes32,bytes) vars) returns()
func (_ERC721 *ERC721TransactorSession) UnlinkLinklist(vars DataTypeslinkLinklistData) (*types.Transaction, error) {
	return _ERC721.Contract.UnlinkLinklist(&_ERC721.TransactOpts, vars)
}

// UnlinkNote is a paid mutator transaction binding the contract method 0x40ad34d8.
//
// Solidity: function unlinkNote((uint256,uint256,uint256,bytes32) vars) returns()
func (_ERC721 *ERC721Transactor) UnlinkNote(opts *bind.TransactOpts, vars DataTypesunlinkNoteData) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "unlinkNote", vars)
}

// UnlinkNote is a paid mutator transaction binding the contract method 0x40ad34d8.
//
// Solidity: function unlinkNote((uint256,uint256,uint256,bytes32) vars) returns()
func (_ERC721 *ERC721Session) UnlinkNote(vars DataTypesunlinkNoteData) (*types.Transaction, error) {
	return _ERC721.Contract.UnlinkNote(&_ERC721.TransactOpts, vars)
}

// UnlinkNote is a paid mutator transaction binding the contract method 0x40ad34d8.
//
// Solidity: function unlinkNote((uint256,uint256,uint256,bytes32) vars) returns()
func (_ERC721 *ERC721TransactorSession) UnlinkNote(vars DataTypesunlinkNoteData) (*types.Transaction, error) {
	return _ERC721.Contract.UnlinkNote(&_ERC721.TransactOpts, vars)
}

// UnlinkProfile is a paid mutator transaction binding the contract method 0x251a99cf.
//
// Solidity: function unlinkProfile((uint256,uint256,bytes32) vars) returns()
func (_ERC721 *ERC721Transactor) UnlinkProfile(opts *bind.TransactOpts, vars DataTypesunlinkProfileData) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "unlinkProfile", vars)
}

// UnlinkProfile is a paid mutator transaction binding the contract method 0x251a99cf.
//
// Solidity: function unlinkProfile((uint256,uint256,bytes32) vars) returns()
func (_ERC721 *ERC721Session) UnlinkProfile(vars DataTypesunlinkProfileData) (*types.Transaction, error) {
	return _ERC721.Contract.UnlinkProfile(&_ERC721.TransactOpts, vars)
}

// UnlinkProfile is a paid mutator transaction binding the contract method 0x251a99cf.
//
// Solidity: function unlinkProfile((uint256,uint256,bytes32) vars) returns()
func (_ERC721 *ERC721TransactorSession) UnlinkProfile(vars DataTypesunlinkProfileData) (*types.Transaction, error) {
	return _ERC721.Contract.UnlinkProfile(&_ERC721.TransactOpts, vars)
}

// UnlinkProfileLink is a paid mutator transaction binding the contract method 0x570477fb.
//
// Solidity: function unlinkProfileLink(uint256 fromProfileId, (uint256,uint256,bytes32) linkData, bytes32 linkType) returns()
func (_ERC721 *ERC721Transactor) UnlinkProfileLink(opts *bind.TransactOpts, fromProfileId *big.Int, linkData DataTypesProfileLinkStruct, linkType [32]byte) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "unlinkProfileLink", fromProfileId, linkData, linkType)
}

// UnlinkProfileLink is a paid mutator transaction binding the contract method 0x570477fb.
//
// Solidity: function unlinkProfileLink(uint256 fromProfileId, (uint256,uint256,bytes32) linkData, bytes32 linkType) returns()
func (_ERC721 *ERC721Session) UnlinkProfileLink(fromProfileId *big.Int, linkData DataTypesProfileLinkStruct, linkType [32]byte) (*types.Transaction, error) {
	return _ERC721.Contract.UnlinkProfileLink(&_ERC721.TransactOpts, fromProfileId, linkData, linkType)
}

// UnlinkProfileLink is a paid mutator transaction binding the contract method 0x570477fb.
//
// Solidity: function unlinkProfileLink(uint256 fromProfileId, (uint256,uint256,bytes32) linkData, bytes32 linkType) returns()
func (_ERC721 *ERC721TransactorSession) UnlinkProfileLink(fromProfileId *big.Int, linkData DataTypesProfileLinkStruct, linkType [32]byte) (*types.Transaction, error) {
	return _ERC721.Contract.UnlinkProfileLink(&_ERC721.TransactOpts, fromProfileId, linkData, linkType)
}

// ERC721ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC721 contract.
type ERC721ApprovalIterator struct {
	Event *ERC721Approval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC721ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721Approval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC721Approval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC721ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721Approval represents a Approval event raised by the ERC721 contract.
type ERC721Approval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ERC721 *ERC721Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*ERC721ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ERC721.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC721ApprovalIterator{contract: _ERC721.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ERC721 *ERC721Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC721Approval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ERC721.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721Approval)
				if err := _ERC721.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ERC721 *ERC721Filterer) ParseApproval(log types.Log) (*ERC721Approval, error) {
	event := new(ERC721Approval)
	if err := _ERC721.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC721ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ERC721 contract.
type ERC721ApprovalForAllIterator struct {
	Event *ERC721ApprovalForAll // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC721ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721ApprovalForAll)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC721ApprovalForAll)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC721ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721ApprovalForAll represents a ApprovalForAll event raised by the ERC721 contract.
type ERC721ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ERC721 *ERC721Filterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ERC721ApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ERC721.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ERC721ApprovalForAllIterator{contract: _ERC721.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ERC721 *ERC721Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ERC721ApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ERC721.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721ApprovalForAll)
				if err := _ERC721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ERC721 *ERC721Filterer) ParseApprovalForAll(log types.Log) (*ERC721ApprovalForAll, error) {
	event := new(ERC721ApprovalForAll)
	if err := _ERC721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC721TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC721 contract.
type ERC721TransferIterator struct {
	Event *ERC721Transfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC721TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721Transfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC721Transfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC721TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721Transfer represents a Transfer event raised by the ERC721 contract.
type ERC721Transfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ERC721 *ERC721Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*ERC721TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ERC721.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC721TransferIterator{contract: _ERC721.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ERC721 *ERC721Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC721Transfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ERC721.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721Transfer)
				if err := _ERC721.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ERC721 *ERC721Filterer) ParseTransfer(log types.Log) (*ERC721Transfer, error) {
	event := new(ERC721Transfer)
	if err := _ERC721.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
