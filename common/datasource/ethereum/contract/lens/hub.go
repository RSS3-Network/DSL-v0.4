// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lens

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

// DataTypesCollectWithSigData is an auto generated low-level Go binding around an user-defined struct.
type DataTypesCollectWithSigData struct {
	Collector common.Address
	ProfileId *big.Int
	PubId     *big.Int
	Data      []byte
	Sig       DataTypesEIP712Signature
}

// DataTypesCommentData is an auto generated low-level Go binding around an user-defined struct.
type DataTypesCommentData struct {
	ProfileId               *big.Int
	ContentURI              string
	ProfileIdPointed        *big.Int
	PubIdPointed            *big.Int
	ReferenceModuleData     []byte
	CollectModule           common.Address
	CollectModuleInitData   []byte
	ReferenceModule         common.Address
	ReferenceModuleInitData []byte
}

// DataTypesCommentWithSigData is an auto generated low-level Go binding around an user-defined struct.
type DataTypesCommentWithSigData struct {
	ProfileId               *big.Int
	ContentURI              string
	ProfileIdPointed        *big.Int
	PubIdPointed            *big.Int
	ReferenceModuleData     []byte
	CollectModule           common.Address
	CollectModuleInitData   []byte
	ReferenceModule         common.Address
	ReferenceModuleInitData []byte
	Sig                     DataTypesEIP712Signature
}

// DataTypesCreateProfileData is an auto generated low-level Go binding around an user-defined struct.
type DataTypesCreateProfileData struct {
	To                   common.Address
	Handle               string
	ImageURI             string
	FollowModule         common.Address
	FollowModuleInitData []byte
	FollowNFTURI         string
}

// DataTypesEIP712Signature is an auto generated low-level Go binding around an user-defined struct.
type DataTypesEIP712Signature struct {
	V        uint8
	R        [32]byte
	S        [32]byte
	Deadline *big.Int
}

// DataTypesFollowWithSigData is an auto generated low-level Go binding around an user-defined struct.
type DataTypesFollowWithSigData struct {
	Follower   common.Address
	ProfileIds []*big.Int
	Datas      [][]byte
	Sig        DataTypesEIP712Signature
}

// DataTypesMirrorData is an auto generated low-level Go binding around an user-defined struct.
type DataTypesMirrorData struct {
	ProfileId               *big.Int
	ProfileIdPointed        *big.Int
	PubIdPointed            *big.Int
	ReferenceModuleData     []byte
	ReferenceModule         common.Address
	ReferenceModuleInitData []byte
}

// DataTypesMirrorWithSigData is an auto generated low-level Go binding around an user-defined struct.
type DataTypesMirrorWithSigData struct {
	ProfileId               *big.Int
	ProfileIdPointed        *big.Int
	PubIdPointed            *big.Int
	ReferenceModuleData     []byte
	ReferenceModule         common.Address
	ReferenceModuleInitData []byte
	Sig                     DataTypesEIP712Signature
}

// DataTypesPostData is an auto generated low-level Go binding around an user-defined struct.
type DataTypesPostData struct {
	ProfileId               *big.Int
	ContentURI              string
	CollectModule           common.Address
	CollectModuleInitData   []byte
	ReferenceModule         common.Address
	ReferenceModuleInitData []byte
}

// DataTypesPostWithSigData is an auto generated low-level Go binding around an user-defined struct.
type DataTypesPostWithSigData struct {
	ProfileId               *big.Int
	ContentURI              string
	CollectModule           common.Address
	CollectModuleInitData   []byte
	ReferenceModule         common.Address
	ReferenceModuleInitData []byte
	Sig                     DataTypesEIP712Signature
}

// DataTypesProfileStruct is an auto generated low-level Go binding around an user-defined struct.
type DataTypesProfileStruct struct {
	PubCount     *big.Int
	FollowModule common.Address
	FollowNFT    common.Address
	Handle       string
	ImageURI     string
	FollowNFTURI string
}

// DataTypesPublicationStruct is an auto generated low-level Go binding around an user-defined struct.
type DataTypesPublicationStruct struct {
	ProfileIdPointed *big.Int
	PubIdPointed     *big.Int
	ContentURI       string
	ReferenceModule  common.Address
	CollectModule    common.Address
	CollectNFT       common.Address
}

// DataTypesSetDefaultProfileWithSigData is an auto generated low-level Go binding around an user-defined struct.
type DataTypesSetDefaultProfileWithSigData struct {
	Wallet    common.Address
	ProfileId *big.Int
	Sig       DataTypesEIP712Signature
}

// DataTypesSetDispatcherWithSigData is an auto generated low-level Go binding around an user-defined struct.
type DataTypesSetDispatcherWithSigData struct {
	ProfileId  *big.Int
	Dispatcher common.Address
	Sig        DataTypesEIP712Signature
}

// DataTypesSetFollowModuleWithSigData is an auto generated low-level Go binding around an user-defined struct.
type DataTypesSetFollowModuleWithSigData struct {
	ProfileId            *big.Int
	FollowModule         common.Address
	FollowModuleInitData []byte
	Sig                  DataTypesEIP712Signature
}

// DataTypesSetFollowNFTURIWithSigData is an auto generated low-level Go binding around an user-defined struct.
type DataTypesSetFollowNFTURIWithSigData struct {
	ProfileId    *big.Int
	FollowNFTURI string
	Sig          DataTypesEIP712Signature
}

// DataTypesSetProfileImageURIWithSigData is an auto generated low-level Go binding around an user-defined struct.
type DataTypesSetProfileImageURIWithSigData struct {
	ProfileId *big.Int
	ImageURI  string
	Sig       DataTypesEIP712Signature
}

// IERC721TimeTokenData is an auto generated low-level Go binding around an user-defined struct.
type IERC721TimeTokenData struct {
	Owner         common.Address
	MintTimestamp *big.Int
}

// HubMetaData contains all meta data concerning the Hub contract.
var HubMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"followNFTImpl\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collectNFTImpl\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CallerNotCollectNFT\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotFollowNFT\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotInitImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmergencyAdminCannotUnpause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitParamsInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Initialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotGovernance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotGovernanceOrEmergencyAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotOwnerOrApproved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotProfileOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotProfileOwnerOrDispatcher\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Paused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProfileCreatorNotWhitelisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProfileImageURILengthInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PublicationDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PublishingPaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignatureExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignatureInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroSpender\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.EIP712Signature\",\"name\":\"sig\",\"type\":\"tuple\"}],\"name\":\"burnWithSig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pubId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"collect\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"collector\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pubId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.EIP712Signature\",\"name\":\"sig\",\"type\":\"tuple\"}],\"internalType\":\"structDataTypes.CollectWithSigData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"collectWithSig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contentURI\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"profileIdPointed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pubIdPointed\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"referenceModuleData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"collectModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"collectModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"referenceModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"referenceModuleInitData\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.CommentData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"comment\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contentURI\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"profileIdPointed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pubIdPointed\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"referenceModuleData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"collectModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"collectModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"referenceModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"referenceModuleInitData\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.EIP712Signature\",\"name\":\"sig\",\"type\":\"tuple\"}],\"internalType\":\"structDataTypes.CommentWithSigData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"commentWithSig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"handle\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"imageURI\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"followModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"followModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"followNFTURI\",\"type\":\"string\"}],\"internalType\":\"structDataTypes.CreateProfileData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"createProfile\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"defaultProfile\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pubId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collectNFTId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"emitCollectNFTTransferEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"followNFTId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"emitFollowNFTTransferEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"exists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"profileIds\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"datas\",\"type\":\"bytes[]\"}],\"name\":\"follow\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"follower\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"profileIds\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"datas\",\"type\":\"bytes[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.EIP712Signature\",\"name\":\"sig\",\"type\":\"tuple\"}],\"internalType\":\"structDataTypes.FollowWithSigData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"followWithSig\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pubId\",\"type\":\"uint256\"}],\"name\":\"getCollectModule\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pubId\",\"type\":\"uint256\"}],\"name\":\"getCollectNFT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCollectNFTImpl\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pubId\",\"type\":\"uint256\"}],\"name\":\"getContentURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"}],\"name\":\"getDispatcher\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDomainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"}],\"name\":\"getFollowModule\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"}],\"name\":\"getFollowNFT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFollowNFTImpl\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"}],\"name\":\"getFollowNFTURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGovernance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"}],\"name\":\"getHandle\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"}],\"name\":\"getProfile\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pubCount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"followModule\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"followNFT\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"handle\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"imageURI\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"followNFTURI\",\"type\":\"string\"}],\"internalType\":\"structDataTypes.ProfileStruct\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"handle\",\"type\":\"string\"}],\"name\":\"getProfileIdByHandle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pubId\",\"type\":\"uint256\"}],\"name\":\"getPub\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"profileIdPointed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pubIdPointed\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contentURI\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"referenceModule\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collectModule\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collectNFT\",\"type\":\"address\"}],\"internalType\":\"structDataTypes.PublicationStruct\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"}],\"name\":\"getPubCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pubId\",\"type\":\"uint256\"}],\"name\":\"getPubPointer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pubId\",\"type\":\"uint256\"}],\"name\":\"getPubType\",\"outputs\":[{\"internalType\":\"enumDataTypes.PubType\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pubId\",\"type\":\"uint256\"}],\"name\":\"getReferenceModule\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getState\",\"outputs\":[{\"internalType\":\"enumDataTypes.ProtocolState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"newGovernance\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collectModule\",\"type\":\"address\"}],\"name\":\"isCollectModuleWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"followModule\",\"type\":\"address\"}],\"name\":\"isFollowModuleWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"profileCreator\",\"type\":\"address\"}],\"name\":\"isProfileCreatorWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"referenceModule\",\"type\":\"address\"}],\"name\":\"isReferenceModuleWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"mintTimestampOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"profileIdPointed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pubIdPointed\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"referenceModuleData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"referenceModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"referenceModuleInitData\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.MirrorData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"mirror\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"profileIdPointed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pubIdPointed\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"referenceModuleData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"referenceModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"referenceModuleInitData\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.EIP712Signature\",\"name\":\"sig\",\"type\":\"tuple\"}],\"internalType\":\"structDataTypes.MirrorWithSigData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"mirrorWithSig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.EIP712Signature\",\"name\":\"sig\",\"type\":\"tuple\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.EIP712Signature\",\"name\":\"sig\",\"type\":\"tuple\"}],\"name\":\"permitForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contentURI\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"collectModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"collectModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"referenceModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"referenceModuleInitData\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.PostData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"post\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contentURI\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"collectModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"collectModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"referenceModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"referenceModuleInitData\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.EIP712Signature\",\"name\":\"sig\",\"type\":\"tuple\"}],\"internalType\":\"structDataTypes.PostWithSigData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"postWithSig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"}],\"name\":\"setDefaultProfile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.EIP712Signature\",\"name\":\"sig\",\"type\":\"tuple\"}],\"internalType\":\"structDataTypes.SetDefaultProfileWithSigData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"setDefaultProfileWithSig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"dispatcher\",\"type\":\"address\"}],\"name\":\"setDispatcher\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"dispatcher\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.EIP712Signature\",\"name\":\"sig\",\"type\":\"tuple\"}],\"internalType\":\"structDataTypes.SetDispatcherWithSigData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"setDispatcherWithSig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newEmergencyAdmin\",\"type\":\"address\"}],\"name\":\"setEmergencyAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"followModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"followModuleInitData\",\"type\":\"bytes\"}],\"name\":\"setFollowModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"followModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"followModuleInitData\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.EIP712Signature\",\"name\":\"sig\",\"type\":\"tuple\"}],\"internalType\":\"structDataTypes.SetFollowModuleWithSigData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"setFollowModuleWithSig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"followNFTURI\",\"type\":\"string\"}],\"name\":\"setFollowNFTURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"followNFTURI\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.EIP712Signature\",\"name\":\"sig\",\"type\":\"tuple\"}],\"internalType\":\"structDataTypes.SetFollowNFTURIWithSigData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"setFollowNFTURIWithSig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newGovernance\",\"type\":\"address\"}],\"name\":\"setGovernance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"imageURI\",\"type\":\"string\"}],\"name\":\"setProfileImageURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"imageURI\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.EIP712Signature\",\"name\":\"sig\",\"type\":\"tuple\"}],\"internalType\":\"structDataTypes.SetProfileImageURIWithSigData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"setProfileImageURIWithSig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumDataTypes.ProtocolState\",\"name\":\"newState\",\"type\":\"uint8\"}],\"name\":\"setState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"sigNonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenDataOf\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"mintTimestamp\",\"type\":\"uint96\"}],\"internalType\":\"structIERC721Time.TokenData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collectModule\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"whitelist\",\"type\":\"bool\"}],\"name\":\"whitelistCollectModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"followModule\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"whitelist\",\"type\":\"bool\"}],\"name\":\"whitelistFollowModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"profileCreator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"whitelist\",\"type\":\"bool\"}],\"name\":\"whitelistProfileCreator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"referenceModule\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"whitelist\",\"type\":\"bool\"}],\"name\":\"whitelistReferenceModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// HubABI is the input ABI used to generate the binding from.
// Deprecated: Use HubMetaData.ABI instead.
var HubABI = HubMetaData.ABI

// Hub is an auto generated Go binding around an Ethereum contract.
type Hub struct {
	HubCaller     // Read-only binding to the contract
	HubTransactor // Write-only binding to the contract
	HubFilterer   // Log filterer for contract events
}

// HubCaller is an auto generated read-only Go binding around an Ethereum contract.
type HubCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HubTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HubTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HubFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HubFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HubSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HubSession struct {
	Contract     *Hub              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HubCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HubCallerSession struct {
	Contract *HubCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// HubTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HubTransactorSession struct {
	Contract     *HubTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HubRaw is an auto generated low-level Go binding around an Ethereum contract.
type HubRaw struct {
	Contract *Hub // Generic contract binding to access the raw methods on
}

// HubCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HubCallerRaw struct {
	Contract *HubCaller // Generic read-only contract binding to access the raw methods on
}

// HubTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HubTransactorRaw struct {
	Contract *HubTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHub creates a new instance of Hub, bound to a specific deployed contract.
func NewHub(address common.Address, backend bind.ContractBackend) (*Hub, error) {
	contract, err := bindHub(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Hub{HubCaller: HubCaller{contract: contract}, HubTransactor: HubTransactor{contract: contract}, HubFilterer: HubFilterer{contract: contract}}, nil
}

// NewHubCaller creates a new read-only instance of Hub, bound to a specific deployed contract.
func NewHubCaller(address common.Address, caller bind.ContractCaller) (*HubCaller, error) {
	contract, err := bindHub(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HubCaller{contract: contract}, nil
}

// NewHubTransactor creates a new write-only instance of Hub, bound to a specific deployed contract.
func NewHubTransactor(address common.Address, transactor bind.ContractTransactor) (*HubTransactor, error) {
	contract, err := bindHub(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HubTransactor{contract: contract}, nil
}

// NewHubFilterer creates a new log filterer instance of Hub, bound to a specific deployed contract.
func NewHubFilterer(address common.Address, filterer bind.ContractFilterer) (*HubFilterer, error) {
	contract, err := bindHub(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HubFilterer{contract: contract}, nil
}

// bindHub binds a generic wrapper to an already deployed contract.
func bindHub(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HubABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hub *HubRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Hub.Contract.HubCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hub *HubRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hub.Contract.HubTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hub *HubRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hub.Contract.HubTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hub *HubCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Hub.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hub *HubTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hub.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hub *HubTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hub.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Hub *HubCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Hub.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Hub *HubSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Hub.Contract.BalanceOf(&_Hub.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Hub *HubCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Hub.Contract.BalanceOf(&_Hub.CallOpts, owner)
}

// DefaultProfile is a free data retrieval call binding the contract method 0x92254a62.
//
// Solidity: function defaultProfile(address wallet) view returns(uint256)
func (_Hub *HubCaller) DefaultProfile(opts *bind.CallOpts, wallet common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Hub.contract.Call(opts, &out, "defaultProfile", wallet)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultProfile is a free data retrieval call binding the contract method 0x92254a62.
//
// Solidity: function defaultProfile(address wallet) view returns(uint256)
func (_Hub *HubSession) DefaultProfile(wallet common.Address) (*big.Int, error) {
	return _Hub.Contract.DefaultProfile(&_Hub.CallOpts, wallet)
}

// DefaultProfile is a free data retrieval call binding the contract method 0x92254a62.
//
// Solidity: function defaultProfile(address wallet) view returns(uint256)
func (_Hub *HubCallerSession) DefaultProfile(wallet common.Address) (*big.Int, error) {
	return _Hub.Contract.DefaultProfile(&_Hub.CallOpts, wallet)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 tokenId) view returns(bool)
func (_Hub *HubCaller) Exists(opts *bind.CallOpts, tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _Hub.contract.Call(opts, &out, "exists", tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 tokenId) view returns(bool)
func (_Hub *HubSession) Exists(tokenId *big.Int) (bool, error) {
	return _Hub.Contract.Exists(&_Hub.CallOpts, tokenId)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 tokenId) view returns(bool)
func (_Hub *HubCallerSession) Exists(tokenId *big.Int) (bool, error) {
	return _Hub.Contract.Exists(&_Hub.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Hub *HubCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Hub.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Hub *HubSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Hub.Contract.GetApproved(&_Hub.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Hub *HubCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Hub.Contract.GetApproved(&_Hub.CallOpts, tokenId)
}

// GetCollectModule is a free data retrieval call binding the contract method 0x57ff49f6.
//
// Solidity: function getCollectModule(uint256 profileId, uint256 pubId) view returns(address)
func (_Hub *HubCaller) GetCollectModule(opts *bind.CallOpts, profileId *big.Int, pubId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Hub.contract.Call(opts, &out, "getCollectModule", profileId, pubId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCollectModule is a free data retrieval call binding the contract method 0x57ff49f6.
//
// Solidity: function getCollectModule(uint256 profileId, uint256 pubId) view returns(address)
func (_Hub *HubSession) GetCollectModule(profileId *big.Int, pubId *big.Int) (common.Address, error) {
	return _Hub.Contract.GetCollectModule(&_Hub.CallOpts, profileId, pubId)
}

// GetCollectModule is a free data retrieval call binding the contract method 0x57ff49f6.
//
// Solidity: function getCollectModule(uint256 profileId, uint256 pubId) view returns(address)
func (_Hub *HubCallerSession) GetCollectModule(profileId *big.Int, pubId *big.Int) (common.Address, error) {
	return _Hub.Contract.GetCollectModule(&_Hub.CallOpts, profileId, pubId)
}

// GetCollectNFT is a free data retrieval call binding the contract method 0x52aaef55.
//
// Solidity: function getCollectNFT(uint256 profileId, uint256 pubId) view returns(address)
func (_Hub *HubCaller) GetCollectNFT(opts *bind.CallOpts, profileId *big.Int, pubId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Hub.contract.Call(opts, &out, "getCollectNFT", profileId, pubId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCollectNFT is a free data retrieval call binding the contract method 0x52aaef55.
//
// Solidity: function getCollectNFT(uint256 profileId, uint256 pubId) view returns(address)
func (_Hub *HubSession) GetCollectNFT(profileId *big.Int, pubId *big.Int) (common.Address, error) {
	return _Hub.Contract.GetCollectNFT(&_Hub.CallOpts, profileId, pubId)
}

// GetCollectNFT is a free data retrieval call binding the contract method 0x52aaef55.
//
// Solidity: function getCollectNFT(uint256 profileId, uint256 pubId) view returns(address)
func (_Hub *HubCallerSession) GetCollectNFT(profileId *big.Int, pubId *big.Int) (common.Address, error) {
	return _Hub.Contract.GetCollectNFT(&_Hub.CallOpts, profileId, pubId)
}

// GetCollectNFTImpl is a free data retrieval call binding the contract method 0x77349a5f.
//
// Solidity: function getCollectNFTImpl() view returns(address)
func (_Hub *HubCaller) GetCollectNFTImpl(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Hub.contract.Call(opts, &out, "getCollectNFTImpl")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCollectNFTImpl is a free data retrieval call binding the contract method 0x77349a5f.
//
// Solidity: function getCollectNFTImpl() view returns(address)
func (_Hub *HubSession) GetCollectNFTImpl() (common.Address, error) {
	return _Hub.Contract.GetCollectNFTImpl(&_Hub.CallOpts)
}

// GetCollectNFTImpl is a free data retrieval call binding the contract method 0x77349a5f.
//
// Solidity: function getCollectNFTImpl() view returns(address)
func (_Hub *HubCallerSession) GetCollectNFTImpl() (common.Address, error) {
	return _Hub.Contract.GetCollectNFTImpl(&_Hub.CallOpts)
}

// GetContentURI is a free data retrieval call binding the contract method 0xb5a31496.
//
// Solidity: function getContentURI(uint256 profileId, uint256 pubId) view returns(string)
func (_Hub *HubCaller) GetContentURI(opts *bind.CallOpts, profileId *big.Int, pubId *big.Int) (string, error) {
	var out []interface{}
	err := _Hub.contract.Call(opts, &out, "getContentURI", profileId, pubId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetContentURI is a free data retrieval call binding the contract method 0xb5a31496.
//
// Solidity: function getContentURI(uint256 profileId, uint256 pubId) view returns(string)
func (_Hub *HubSession) GetContentURI(profileId *big.Int, pubId *big.Int) (string, error) {
	return _Hub.Contract.GetContentURI(&_Hub.CallOpts, profileId, pubId)
}

// GetContentURI is a free data retrieval call binding the contract method 0xb5a31496.
//
// Solidity: function getContentURI(uint256 profileId, uint256 pubId) view returns(string)
func (_Hub *HubCallerSession) GetContentURI(profileId *big.Int, pubId *big.Int) (string, error) {
	return _Hub.Contract.GetContentURI(&_Hub.CallOpts, profileId, pubId)
}

// GetDispatcher is a free data retrieval call binding the contract method 0x540528b9.
//
// Solidity: function getDispatcher(uint256 profileId) view returns(address)
func (_Hub *HubCaller) GetDispatcher(opts *bind.CallOpts, profileId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Hub.contract.Call(opts, &out, "getDispatcher", profileId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetDispatcher is a free data retrieval call binding the contract method 0x540528b9.
//
// Solidity: function getDispatcher(uint256 profileId) view returns(address)
func (_Hub *HubSession) GetDispatcher(profileId *big.Int) (common.Address, error) {
	return _Hub.Contract.GetDispatcher(&_Hub.CallOpts, profileId)
}

// GetDispatcher is a free data retrieval call binding the contract method 0x540528b9.
//
// Solidity: function getDispatcher(uint256 profileId) view returns(address)
func (_Hub *HubCallerSession) GetDispatcher(profileId *big.Int) (common.Address, error) {
	return _Hub.Contract.GetDispatcher(&_Hub.CallOpts, profileId)
}

// GetDomainSeparator is a free data retrieval call binding the contract method 0xed24911d.
//
// Solidity: function getDomainSeparator() view returns(bytes32)
func (_Hub *HubCaller) GetDomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Hub.contract.Call(opts, &out, "getDomainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetDomainSeparator is a free data retrieval call binding the contract method 0xed24911d.
//
// Solidity: function getDomainSeparator() view returns(bytes32)
func (_Hub *HubSession) GetDomainSeparator() ([32]byte, error) {
	return _Hub.Contract.GetDomainSeparator(&_Hub.CallOpts)
}

// GetDomainSeparator is a free data retrieval call binding the contract method 0xed24911d.
//
// Solidity: function getDomainSeparator() view returns(bytes32)
func (_Hub *HubCallerSession) GetDomainSeparator() ([32]byte, error) {
	return _Hub.Contract.GetDomainSeparator(&_Hub.CallOpts)
}

// GetFollowModule is a free data retrieval call binding the contract method 0xd923d20c.
//
// Solidity: function getFollowModule(uint256 profileId) view returns(address)
func (_Hub *HubCaller) GetFollowModule(opts *bind.CallOpts, profileId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Hub.contract.Call(opts, &out, "getFollowModule", profileId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFollowModule is a free data retrieval call binding the contract method 0xd923d20c.
//
// Solidity: function getFollowModule(uint256 profileId) view returns(address)
func (_Hub *HubSession) GetFollowModule(profileId *big.Int) (common.Address, error) {
	return _Hub.Contract.GetFollowModule(&_Hub.CallOpts, profileId)
}

// GetFollowModule is a free data retrieval call binding the contract method 0xd923d20c.
//
// Solidity: function getFollowModule(uint256 profileId) view returns(address)
func (_Hub *HubCallerSession) GetFollowModule(profileId *big.Int) (common.Address, error) {
	return _Hub.Contract.GetFollowModule(&_Hub.CallOpts, profileId)
}

// GetFollowNFT is a free data retrieval call binding the contract method 0xa9ec6563.
//
// Solidity: function getFollowNFT(uint256 profileId) view returns(address)
func (_Hub *HubCaller) GetFollowNFT(opts *bind.CallOpts, profileId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Hub.contract.Call(opts, &out, "getFollowNFT", profileId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFollowNFT is a free data retrieval call binding the contract method 0xa9ec6563.
//
// Solidity: function getFollowNFT(uint256 profileId) view returns(address)
func (_Hub *HubSession) GetFollowNFT(profileId *big.Int) (common.Address, error) {
	return _Hub.Contract.GetFollowNFT(&_Hub.CallOpts, profileId)
}

// GetFollowNFT is a free data retrieval call binding the contract method 0xa9ec6563.
//
// Solidity: function getFollowNFT(uint256 profileId) view returns(address)
func (_Hub *HubCallerSession) GetFollowNFT(profileId *big.Int) (common.Address, error) {
	return _Hub.Contract.GetFollowNFT(&_Hub.CallOpts, profileId)
}

// GetFollowNFTImpl is a free data retrieval call binding the contract method 0x3502ac4b.
//
// Solidity: function getFollowNFTImpl() view returns(address)
func (_Hub *HubCaller) GetFollowNFTImpl(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Hub.contract.Call(opts, &out, "getFollowNFTImpl")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFollowNFTImpl is a free data retrieval call binding the contract method 0x3502ac4b.
//
// Solidity: function getFollowNFTImpl() view returns(address)
func (_Hub *HubSession) GetFollowNFTImpl() (common.Address, error) {
	return _Hub.Contract.GetFollowNFTImpl(&_Hub.CallOpts)
}

// GetFollowNFTImpl is a free data retrieval call binding the contract method 0x3502ac4b.
//
// Solidity: function getFollowNFTImpl() view returns(address)
func (_Hub *HubCallerSession) GetFollowNFTImpl() (common.Address, error) {
	return _Hub.Contract.GetFollowNFTImpl(&_Hub.CallOpts)
}

// GetFollowNFTURI is a free data retrieval call binding the contract method 0x374c9473.
//
// Solidity: function getFollowNFTURI(uint256 profileId) view returns(string)
func (_Hub *HubCaller) GetFollowNFTURI(opts *bind.CallOpts, profileId *big.Int) (string, error) {
	var out []interface{}
	err := _Hub.contract.Call(opts, &out, "getFollowNFTURI", profileId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetFollowNFTURI is a free data retrieval call binding the contract method 0x374c9473.
//
// Solidity: function getFollowNFTURI(uint256 profileId) view returns(string)
func (_Hub *HubSession) GetFollowNFTURI(profileId *big.Int) (string, error) {
	return _Hub.Contract.GetFollowNFTURI(&_Hub.CallOpts, profileId)
}

// GetFollowNFTURI is a free data retrieval call binding the contract method 0x374c9473.
//
// Solidity: function getFollowNFTURI(uint256 profileId) view returns(string)
func (_Hub *HubCallerSession) GetFollowNFTURI(profileId *big.Int) (string, error) {
	return _Hub.Contract.GetFollowNFTURI(&_Hub.CallOpts, profileId)
}

// GetGovernance is a free data retrieval call binding the contract method 0x289b3c0d.
//
// Solidity: function getGovernance() view returns(address)
func (_Hub *HubCaller) GetGovernance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Hub.contract.Call(opts, &out, "getGovernance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetGovernance is a free data retrieval call binding the contract method 0x289b3c0d.
//
// Solidity: function getGovernance() view returns(address)
func (_Hub *HubSession) GetGovernance() (common.Address, error) {
	return _Hub.Contract.GetGovernance(&_Hub.CallOpts)
}

// GetGovernance is a free data retrieval call binding the contract method 0x289b3c0d.
//
// Solidity: function getGovernance() view returns(address)
func (_Hub *HubCallerSession) GetGovernance() (common.Address, error) {
	return _Hub.Contract.GetGovernance(&_Hub.CallOpts)
}

// GetHandle is a free data retrieval call binding the contract method 0xec81d194.
//
// Solidity: function getHandle(uint256 profileId) view returns(string)
func (_Hub *HubCaller) GetHandle(opts *bind.CallOpts, profileId *big.Int) (string, error) {
	var out []interface{}
	err := _Hub.contract.Call(opts, &out, "getHandle", profileId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetHandle is a free data retrieval call binding the contract method 0xec81d194.
//
// Solidity: function getHandle(uint256 profileId) view returns(string)
func (_Hub *HubSession) GetHandle(profileId *big.Int) (string, error) {
	return _Hub.Contract.GetHandle(&_Hub.CallOpts, profileId)
}

// GetHandle is a free data retrieval call binding the contract method 0xec81d194.
//
// Solidity: function getHandle(uint256 profileId) view returns(string)
func (_Hub *HubCallerSession) GetHandle(profileId *big.Int) (string, error) {
	return _Hub.Contract.GetHandle(&_Hub.CallOpts, profileId)
}

// GetProfile is a free data retrieval call binding the contract method 0xf08f4f64.
//
// Solidity: function getProfile(uint256 profileId) view returns((uint256,address,address,string,string,string))
func (_Hub *HubCaller) GetProfile(opts *bind.CallOpts, profileId *big.Int) (DataTypesProfileStruct, error) {
	var out []interface{}
	err := _Hub.contract.Call(opts, &out, "getProfile", profileId)

	if err != nil {
		return *new(DataTypesProfileStruct), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesProfileStruct)).(*DataTypesProfileStruct)

	return out0, err

}

// GetProfile is a free data retrieval call binding the contract method 0xf08f4f64.
//
// Solidity: function getProfile(uint256 profileId) view returns((uint256,address,address,string,string,string))
func (_Hub *HubSession) GetProfile(profileId *big.Int) (DataTypesProfileStruct, error) {
	return _Hub.Contract.GetProfile(&_Hub.CallOpts, profileId)
}

// GetProfile is a free data retrieval call binding the contract method 0xf08f4f64.
//
// Solidity: function getProfile(uint256 profileId) view returns((uint256,address,address,string,string,string))
func (_Hub *HubCallerSession) GetProfile(profileId *big.Int) (DataTypesProfileStruct, error) {
	return _Hub.Contract.GetProfile(&_Hub.CallOpts, profileId)
}

// GetProfileIdByHandle is a free data retrieval call binding the contract method 0x20fa728a.
//
// Solidity: function getProfileIdByHandle(string handle) view returns(uint256)
func (_Hub *HubCaller) GetProfileIdByHandle(opts *bind.CallOpts, handle string) (*big.Int, error) {
	var out []interface{}
	err := _Hub.contract.Call(opts, &out, "getProfileIdByHandle", handle)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProfileIdByHandle is a free data retrieval call binding the contract method 0x20fa728a.
//
// Solidity: function getProfileIdByHandle(string handle) view returns(uint256)
func (_Hub *HubSession) GetProfileIdByHandle(handle string) (*big.Int, error) {
	return _Hub.Contract.GetProfileIdByHandle(&_Hub.CallOpts, handle)
}

// GetProfileIdByHandle is a free data retrieval call binding the contract method 0x20fa728a.
//
// Solidity: function getProfileIdByHandle(string handle) view returns(uint256)
func (_Hub *HubCallerSession) GetProfileIdByHandle(handle string) (*big.Int, error) {
	return _Hub.Contract.GetProfileIdByHandle(&_Hub.CallOpts, handle)
}

// GetPub is a free data retrieval call binding the contract method 0xc736c990.
//
// Solidity: function getPub(uint256 profileId, uint256 pubId) view returns((uint256,uint256,string,address,address,address))
func (_Hub *HubCaller) GetPub(opts *bind.CallOpts, profileId *big.Int, pubId *big.Int) (DataTypesPublicationStruct, error) {
	var out []interface{}
	err := _Hub.contract.Call(opts, &out, "getPub", profileId, pubId)

	if err != nil {
		return *new(DataTypesPublicationStruct), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesPublicationStruct)).(*DataTypesPublicationStruct)

	return out0, err

}

// GetPub is a free data retrieval call binding the contract method 0xc736c990.
//
// Solidity: function getPub(uint256 profileId, uint256 pubId) view returns((uint256,uint256,string,address,address,address))
func (_Hub *HubSession) GetPub(profileId *big.Int, pubId *big.Int) (DataTypesPublicationStruct, error) {
	return _Hub.Contract.GetPub(&_Hub.CallOpts, profileId, pubId)
}

// GetPub is a free data retrieval call binding the contract method 0xc736c990.
//
// Solidity: function getPub(uint256 profileId, uint256 pubId) view returns((uint256,uint256,string,address,address,address))
func (_Hub *HubCallerSession) GetPub(profileId *big.Int, pubId *big.Int) (DataTypesPublicationStruct, error) {
	return _Hub.Contract.GetPub(&_Hub.CallOpts, profileId, pubId)
}

// GetPubCount is a free data retrieval call binding the contract method 0x3a15ff07.
//
// Solidity: function getPubCount(uint256 profileId) view returns(uint256)
func (_Hub *HubCaller) GetPubCount(opts *bind.CallOpts, profileId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Hub.contract.Call(opts, &out, "getPubCount", profileId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPubCount is a free data retrieval call binding the contract method 0x3a15ff07.
//
// Solidity: function getPubCount(uint256 profileId) view returns(uint256)
func (_Hub *HubSession) GetPubCount(profileId *big.Int) (*big.Int, error) {
	return _Hub.Contract.GetPubCount(&_Hub.CallOpts, profileId)
}

// GetPubCount is a free data retrieval call binding the contract method 0x3a15ff07.
//
// Solidity: function getPubCount(uint256 profileId) view returns(uint256)
func (_Hub *HubCallerSession) GetPubCount(profileId *big.Int) (*big.Int, error) {
	return _Hub.Contract.GetPubCount(&_Hub.CallOpts, profileId)
}

// GetPubPointer is a free data retrieval call binding the contract method 0x5ca3eebf.
//
// Solidity: function getPubPointer(uint256 profileId, uint256 pubId) view returns(uint256, uint256)
func (_Hub *HubCaller) GetPubPointer(opts *bind.CallOpts, profileId *big.Int, pubId *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Hub.contract.Call(opts, &out, "getPubPointer", profileId, pubId)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetPubPointer is a free data retrieval call binding the contract method 0x5ca3eebf.
//
// Solidity: function getPubPointer(uint256 profileId, uint256 pubId) view returns(uint256, uint256)
func (_Hub *HubSession) GetPubPointer(profileId *big.Int, pubId *big.Int) (*big.Int, *big.Int, error) {
	return _Hub.Contract.GetPubPointer(&_Hub.CallOpts, profileId, pubId)
}

// GetPubPointer is a free data retrieval call binding the contract method 0x5ca3eebf.
//
// Solidity: function getPubPointer(uint256 profileId, uint256 pubId) view returns(uint256, uint256)
func (_Hub *HubCallerSession) GetPubPointer(profileId *big.Int, pubId *big.Int) (*big.Int, *big.Int, error) {
	return _Hub.Contract.GetPubPointer(&_Hub.CallOpts, profileId, pubId)
}

// GetPubType is a free data retrieval call binding the contract method 0x31fff07c.
//
// Solidity: function getPubType(uint256 profileId, uint256 pubId) view returns(uint8)
func (_Hub *HubCaller) GetPubType(opts *bind.CallOpts, profileId *big.Int, pubId *big.Int) (uint8, error) {
	var out []interface{}
	err := _Hub.contract.Call(opts, &out, "getPubType", profileId, pubId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetPubType is a free data retrieval call binding the contract method 0x31fff07c.
//
// Solidity: function getPubType(uint256 profileId, uint256 pubId) view returns(uint8)
func (_Hub *HubSession) GetPubType(profileId *big.Int, pubId *big.Int) (uint8, error) {
	return _Hub.Contract.GetPubType(&_Hub.CallOpts, profileId, pubId)
}

// GetPubType is a free data retrieval call binding the contract method 0x31fff07c.
//
// Solidity: function getPubType(uint256 profileId, uint256 pubId) view returns(uint8)
func (_Hub *HubCallerSession) GetPubType(profileId *big.Int, pubId *big.Int) (uint8, error) {
	return _Hub.Contract.GetPubType(&_Hub.CallOpts, profileId, pubId)
}

// GetReferenceModule is a free data retrieval call binding the contract method 0xb7984c05.
//
// Solidity: function getReferenceModule(uint256 profileId, uint256 pubId) view returns(address)
func (_Hub *HubCaller) GetReferenceModule(opts *bind.CallOpts, profileId *big.Int, pubId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Hub.contract.Call(opts, &out, "getReferenceModule", profileId, pubId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetReferenceModule is a free data retrieval call binding the contract method 0xb7984c05.
//
// Solidity: function getReferenceModule(uint256 profileId, uint256 pubId) view returns(address)
func (_Hub *HubSession) GetReferenceModule(profileId *big.Int, pubId *big.Int) (common.Address, error) {
	return _Hub.Contract.GetReferenceModule(&_Hub.CallOpts, profileId, pubId)
}

// GetReferenceModule is a free data retrieval call binding the contract method 0xb7984c05.
//
// Solidity: function getReferenceModule(uint256 profileId, uint256 pubId) view returns(address)
func (_Hub *HubCallerSession) GetReferenceModule(profileId *big.Int, pubId *big.Int) (common.Address, error) {
	return _Hub.Contract.GetReferenceModule(&_Hub.CallOpts, profileId, pubId)
}

// GetState is a free data retrieval call binding the contract method 0x1865c57d.
//
// Solidity: function getState() view returns(uint8)
func (_Hub *HubCaller) GetState(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Hub.contract.Call(opts, &out, "getState")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetState is a free data retrieval call binding the contract method 0x1865c57d.
//
// Solidity: function getState() view returns(uint8)
func (_Hub *HubSession) GetState() (uint8, error) {
	return _Hub.Contract.GetState(&_Hub.CallOpts)
}

// GetState is a free data retrieval call binding the contract method 0x1865c57d.
//
// Solidity: function getState() view returns(uint8)
func (_Hub *HubCallerSession) GetState() (uint8, error) {
	return _Hub.Contract.GetState(&_Hub.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Hub *HubCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Hub.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Hub *HubSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Hub.Contract.IsApprovedForAll(&_Hub.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Hub *HubCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Hub.Contract.IsApprovedForAll(&_Hub.CallOpts, owner, operator)
}

// IsCollectModuleWhitelisted is a free data retrieval call binding the contract method 0xad3e72bf.
//
// Solidity: function isCollectModuleWhitelisted(address collectModule) view returns(bool)
func (_Hub *HubCaller) IsCollectModuleWhitelisted(opts *bind.CallOpts, collectModule common.Address) (bool, error) {
	var out []interface{}
	err := _Hub.contract.Call(opts, &out, "isCollectModuleWhitelisted", collectModule)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCollectModuleWhitelisted is a free data retrieval call binding the contract method 0xad3e72bf.
//
// Solidity: function isCollectModuleWhitelisted(address collectModule) view returns(bool)
func (_Hub *HubSession) IsCollectModuleWhitelisted(collectModule common.Address) (bool, error) {
	return _Hub.Contract.IsCollectModuleWhitelisted(&_Hub.CallOpts, collectModule)
}

// IsCollectModuleWhitelisted is a free data retrieval call binding the contract method 0xad3e72bf.
//
// Solidity: function isCollectModuleWhitelisted(address collectModule) view returns(bool)
func (_Hub *HubCallerSession) IsCollectModuleWhitelisted(collectModule common.Address) (bool, error) {
	return _Hub.Contract.IsCollectModuleWhitelisted(&_Hub.CallOpts, collectModule)
}

// IsFollowModuleWhitelisted is a free data retrieval call binding the contract method 0x1cbb2620.
//
// Solidity: function isFollowModuleWhitelisted(address followModule) view returns(bool)
func (_Hub *HubCaller) IsFollowModuleWhitelisted(opts *bind.CallOpts, followModule common.Address) (bool, error) {
	var out []interface{}
	err := _Hub.contract.Call(opts, &out, "isFollowModuleWhitelisted", followModule)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFollowModuleWhitelisted is a free data retrieval call binding the contract method 0x1cbb2620.
//
// Solidity: function isFollowModuleWhitelisted(address followModule) view returns(bool)
func (_Hub *HubSession) IsFollowModuleWhitelisted(followModule common.Address) (bool, error) {
	return _Hub.Contract.IsFollowModuleWhitelisted(&_Hub.CallOpts, followModule)
}

// IsFollowModuleWhitelisted is a free data retrieval call binding the contract method 0x1cbb2620.
//
// Solidity: function isFollowModuleWhitelisted(address followModule) view returns(bool)
func (_Hub *HubCallerSession) IsFollowModuleWhitelisted(followModule common.Address) (bool, error) {
	return _Hub.Contract.IsFollowModuleWhitelisted(&_Hub.CallOpts, followModule)
}

// IsProfileCreatorWhitelisted is a free data retrieval call binding the contract method 0xaf05dd22.
//
// Solidity: function isProfileCreatorWhitelisted(address profileCreator) view returns(bool)
func (_Hub *HubCaller) IsProfileCreatorWhitelisted(opts *bind.CallOpts, profileCreator common.Address) (bool, error) {
	var out []interface{}
	err := _Hub.contract.Call(opts, &out, "isProfileCreatorWhitelisted", profileCreator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsProfileCreatorWhitelisted is a free data retrieval call binding the contract method 0xaf05dd22.
//
// Solidity: function isProfileCreatorWhitelisted(address profileCreator) view returns(bool)
func (_Hub *HubSession) IsProfileCreatorWhitelisted(profileCreator common.Address) (bool, error) {
	return _Hub.Contract.IsProfileCreatorWhitelisted(&_Hub.CallOpts, profileCreator)
}

// IsProfileCreatorWhitelisted is a free data retrieval call binding the contract method 0xaf05dd22.
//
// Solidity: function isProfileCreatorWhitelisted(address profileCreator) view returns(bool)
func (_Hub *HubCallerSession) IsProfileCreatorWhitelisted(profileCreator common.Address) (bool, error) {
	return _Hub.Contract.IsProfileCreatorWhitelisted(&_Hub.CallOpts, profileCreator)
}

// IsReferenceModuleWhitelisted is a free data retrieval call binding the contract method 0x8e204fb4.
//
// Solidity: function isReferenceModuleWhitelisted(address referenceModule) view returns(bool)
func (_Hub *HubCaller) IsReferenceModuleWhitelisted(opts *bind.CallOpts, referenceModule common.Address) (bool, error) {
	var out []interface{}
	err := _Hub.contract.Call(opts, &out, "isReferenceModuleWhitelisted", referenceModule)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsReferenceModuleWhitelisted is a free data retrieval call binding the contract method 0x8e204fb4.
//
// Solidity: function isReferenceModuleWhitelisted(address referenceModule) view returns(bool)
func (_Hub *HubSession) IsReferenceModuleWhitelisted(referenceModule common.Address) (bool, error) {
	return _Hub.Contract.IsReferenceModuleWhitelisted(&_Hub.CallOpts, referenceModule)
}

// IsReferenceModuleWhitelisted is a free data retrieval call binding the contract method 0x8e204fb4.
//
// Solidity: function isReferenceModuleWhitelisted(address referenceModule) view returns(bool)
func (_Hub *HubCallerSession) IsReferenceModuleWhitelisted(referenceModule common.Address) (bool, error) {
	return _Hub.Contract.IsReferenceModuleWhitelisted(&_Hub.CallOpts, referenceModule)
}

// MintTimestampOf is a free data retrieval call binding the contract method 0x50ddf35c.
//
// Solidity: function mintTimestampOf(uint256 tokenId) view returns(uint256)
func (_Hub *HubCaller) MintTimestampOf(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Hub.contract.Call(opts, &out, "mintTimestampOf", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintTimestampOf is a free data retrieval call binding the contract method 0x50ddf35c.
//
// Solidity: function mintTimestampOf(uint256 tokenId) view returns(uint256)
func (_Hub *HubSession) MintTimestampOf(tokenId *big.Int) (*big.Int, error) {
	return _Hub.Contract.MintTimestampOf(&_Hub.CallOpts, tokenId)
}

// MintTimestampOf is a free data retrieval call binding the contract method 0x50ddf35c.
//
// Solidity: function mintTimestampOf(uint256 tokenId) view returns(uint256)
func (_Hub *HubCallerSession) MintTimestampOf(tokenId *big.Int) (*big.Int, error) {
	return _Hub.Contract.MintTimestampOf(&_Hub.CallOpts, tokenId)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Hub *HubCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Hub.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Hub *HubSession) Name() (string, error) {
	return _Hub.Contract.Name(&_Hub.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Hub *HubCallerSession) Name() (string, error) {
	return _Hub.Contract.Name(&_Hub.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Hub *HubCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Hub.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Hub *HubSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Hub.Contract.OwnerOf(&_Hub.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Hub *HubCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Hub.Contract.OwnerOf(&_Hub.CallOpts, tokenId)
}

// SigNonces is a free data retrieval call binding the contract method 0xf990ccd7.
//
// Solidity: function sigNonces(address ) view returns(uint256)
func (_Hub *HubCaller) SigNonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Hub.contract.Call(opts, &out, "sigNonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SigNonces is a free data retrieval call binding the contract method 0xf990ccd7.
//
// Solidity: function sigNonces(address ) view returns(uint256)
func (_Hub *HubSession) SigNonces(arg0 common.Address) (*big.Int, error) {
	return _Hub.Contract.SigNonces(&_Hub.CallOpts, arg0)
}

// SigNonces is a free data retrieval call binding the contract method 0xf990ccd7.
//
// Solidity: function sigNonces(address ) view returns(uint256)
func (_Hub *HubCallerSession) SigNonces(arg0 common.Address) (*big.Int, error) {
	return _Hub.Contract.SigNonces(&_Hub.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Hub *HubCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Hub.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Hub *HubSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Hub.Contract.SupportsInterface(&_Hub.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Hub *HubCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Hub.Contract.SupportsInterface(&_Hub.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Hub *HubCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Hub.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Hub *HubSession) Symbol() (string, error) {
	return _Hub.Contract.Symbol(&_Hub.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Hub *HubCallerSession) Symbol() (string, error) {
	return _Hub.Contract.Symbol(&_Hub.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Hub *HubCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Hub.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Hub *HubSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Hub.Contract.TokenByIndex(&_Hub.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Hub *HubCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Hub.Contract.TokenByIndex(&_Hub.CallOpts, index)
}

// TokenDataOf is a free data retrieval call binding the contract method 0xc0da9bcd.
//
// Solidity: function tokenDataOf(uint256 tokenId) view returns((address,uint96))
func (_Hub *HubCaller) TokenDataOf(opts *bind.CallOpts, tokenId *big.Int) (IERC721TimeTokenData, error) {
	var out []interface{}
	err := _Hub.contract.Call(opts, &out, "tokenDataOf", tokenId)

	if err != nil {
		return *new(IERC721TimeTokenData), err
	}

	out0 := *abi.ConvertType(out[0], new(IERC721TimeTokenData)).(*IERC721TimeTokenData)

	return out0, err

}

// TokenDataOf is a free data retrieval call binding the contract method 0xc0da9bcd.
//
// Solidity: function tokenDataOf(uint256 tokenId) view returns((address,uint96))
func (_Hub *HubSession) TokenDataOf(tokenId *big.Int) (IERC721TimeTokenData, error) {
	return _Hub.Contract.TokenDataOf(&_Hub.CallOpts, tokenId)
}

// TokenDataOf is a free data retrieval call binding the contract method 0xc0da9bcd.
//
// Solidity: function tokenDataOf(uint256 tokenId) view returns((address,uint96))
func (_Hub *HubCallerSession) TokenDataOf(tokenId *big.Int) (IERC721TimeTokenData, error) {
	return _Hub.Contract.TokenDataOf(&_Hub.CallOpts, tokenId)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Hub *HubCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Hub.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Hub *HubSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Hub.Contract.TokenOfOwnerByIndex(&_Hub.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Hub *HubCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Hub.Contract.TokenOfOwnerByIndex(&_Hub.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Hub *HubCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Hub.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Hub *HubSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Hub.Contract.TokenURI(&_Hub.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Hub *HubCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Hub.Contract.TokenURI(&_Hub.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Hub *HubCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hub.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Hub *HubSession) TotalSupply() (*big.Int, error) {
	return _Hub.Contract.TotalSupply(&_Hub.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Hub *HubCallerSession) TotalSupply() (*big.Int, error) {
	return _Hub.Contract.TotalSupply(&_Hub.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Hub *HubTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Hub.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Hub *HubSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Hub.Contract.Approve(&_Hub.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Hub *HubTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Hub.Contract.Approve(&_Hub.TransactOpts, to, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_Hub *HubTransactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Hub.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_Hub *HubSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _Hub.Contract.Burn(&_Hub.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_Hub *HubTransactorSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _Hub.Contract.Burn(&_Hub.TransactOpts, tokenId)
}

// BurnWithSig is a paid mutator transaction binding the contract method 0xdd69cdb1.
//
// Solidity: function burnWithSig(uint256 tokenId, (uint8,bytes32,bytes32,uint256) sig) returns()
func (_Hub *HubTransactor) BurnWithSig(opts *bind.TransactOpts, tokenId *big.Int, sig DataTypesEIP712Signature) (*types.Transaction, error) {
	return _Hub.contract.Transact(opts, "burnWithSig", tokenId, sig)
}

// BurnWithSig is a paid mutator transaction binding the contract method 0xdd69cdb1.
//
// Solidity: function burnWithSig(uint256 tokenId, (uint8,bytes32,bytes32,uint256) sig) returns()
func (_Hub *HubSession) BurnWithSig(tokenId *big.Int, sig DataTypesEIP712Signature) (*types.Transaction, error) {
	return _Hub.Contract.BurnWithSig(&_Hub.TransactOpts, tokenId, sig)
}

// BurnWithSig is a paid mutator transaction binding the contract method 0xdd69cdb1.
//
// Solidity: function burnWithSig(uint256 tokenId, (uint8,bytes32,bytes32,uint256) sig) returns()
func (_Hub *HubTransactorSession) BurnWithSig(tokenId *big.Int, sig DataTypesEIP712Signature) (*types.Transaction, error) {
	return _Hub.Contract.BurnWithSig(&_Hub.TransactOpts, tokenId, sig)
}

// Collect is a paid mutator transaction binding the contract method 0x84114ad4.
//
// Solidity: function collect(uint256 profileId, uint256 pubId, bytes data) returns(uint256)
func (_Hub *HubTransactor) Collect(opts *bind.TransactOpts, profileId *big.Int, pubId *big.Int, data []byte) (*types.Transaction, error) {
	return _Hub.contract.Transact(opts, "collect", profileId, pubId, data)
}

// Collect is a paid mutator transaction binding the contract method 0x84114ad4.
//
// Solidity: function collect(uint256 profileId, uint256 pubId, bytes data) returns(uint256)
func (_Hub *HubSession) Collect(profileId *big.Int, pubId *big.Int, data []byte) (*types.Transaction, error) {
	return _Hub.Contract.Collect(&_Hub.TransactOpts, profileId, pubId, data)
}

// Collect is a paid mutator transaction binding the contract method 0x84114ad4.
//
// Solidity: function collect(uint256 profileId, uint256 pubId, bytes data) returns(uint256)
func (_Hub *HubTransactorSession) Collect(profileId *big.Int, pubId *big.Int, data []byte) (*types.Transaction, error) {
	return _Hub.Contract.Collect(&_Hub.TransactOpts, profileId, pubId, data)
}

// CollectWithSig is a paid mutator transaction binding the contract method 0xb48951e4.
//
// Solidity: function collectWithSig((address,uint256,uint256,bytes,(uint8,bytes32,bytes32,uint256)) vars) returns(uint256)
func (_Hub *HubTransactor) CollectWithSig(opts *bind.TransactOpts, vars DataTypesCollectWithSigData) (*types.Transaction, error) {
	return _Hub.contract.Transact(opts, "collectWithSig", vars)
}

// CollectWithSig is a paid mutator transaction binding the contract method 0xb48951e4.
//
// Solidity: function collectWithSig((address,uint256,uint256,bytes,(uint8,bytes32,bytes32,uint256)) vars) returns(uint256)
func (_Hub *HubSession) CollectWithSig(vars DataTypesCollectWithSigData) (*types.Transaction, error) {
	return _Hub.Contract.CollectWithSig(&_Hub.TransactOpts, vars)
}

// CollectWithSig is a paid mutator transaction binding the contract method 0xb48951e4.
//
// Solidity: function collectWithSig((address,uint256,uint256,bytes,(uint8,bytes32,bytes32,uint256)) vars) returns(uint256)
func (_Hub *HubTransactorSession) CollectWithSig(vars DataTypesCollectWithSigData) (*types.Transaction, error) {
	return _Hub.Contract.CollectWithSig(&_Hub.TransactOpts, vars)
}

// Comment is a paid mutator transaction binding the contract method 0xb6f32d2b.
//
// Solidity: function comment((uint256,string,uint256,uint256,bytes,address,bytes,address,bytes) vars) returns(uint256)
func (_Hub *HubTransactor) Comment(opts *bind.TransactOpts, vars DataTypesCommentData) (*types.Transaction, error) {
	return _Hub.contract.Transact(opts, "comment", vars)
}

// Comment is a paid mutator transaction binding the contract method 0xb6f32d2b.
//
// Solidity: function comment((uint256,string,uint256,uint256,bytes,address,bytes,address,bytes) vars) returns(uint256)
func (_Hub *HubSession) Comment(vars DataTypesCommentData) (*types.Transaction, error) {
	return _Hub.Contract.Comment(&_Hub.TransactOpts, vars)
}

// Comment is a paid mutator transaction binding the contract method 0xb6f32d2b.
//
// Solidity: function comment((uint256,string,uint256,uint256,bytes,address,bytes,address,bytes) vars) returns(uint256)
func (_Hub *HubTransactorSession) Comment(vars DataTypesCommentData) (*types.Transaction, error) {
	return _Hub.Contract.Comment(&_Hub.TransactOpts, vars)
}

// CommentWithSig is a paid mutator transaction binding the contract method 0x7a375716.
//
// Solidity: function commentWithSig((uint256,string,uint256,uint256,bytes,address,bytes,address,bytes,(uint8,bytes32,bytes32,uint256)) vars) returns(uint256)
func (_Hub *HubTransactor) CommentWithSig(opts *bind.TransactOpts, vars DataTypesCommentWithSigData) (*types.Transaction, error) {
	return _Hub.contract.Transact(opts, "commentWithSig", vars)
}

// CommentWithSig is a paid mutator transaction binding the contract method 0x7a375716.
//
// Solidity: function commentWithSig((uint256,string,uint256,uint256,bytes,address,bytes,address,bytes,(uint8,bytes32,bytes32,uint256)) vars) returns(uint256)
func (_Hub *HubSession) CommentWithSig(vars DataTypesCommentWithSigData) (*types.Transaction, error) {
	return _Hub.Contract.CommentWithSig(&_Hub.TransactOpts, vars)
}

// CommentWithSig is a paid mutator transaction binding the contract method 0x7a375716.
//
// Solidity: function commentWithSig((uint256,string,uint256,uint256,bytes,address,bytes,address,bytes,(uint8,bytes32,bytes32,uint256)) vars) returns(uint256)
func (_Hub *HubTransactorSession) CommentWithSig(vars DataTypesCommentWithSigData) (*types.Transaction, error) {
	return _Hub.Contract.CommentWithSig(&_Hub.TransactOpts, vars)
}

// CreateProfile is a paid mutator transaction binding the contract method 0xffea138e.
//
// Solidity: function createProfile((address,string,string,address,bytes,string) vars) returns(uint256)
func (_Hub *HubTransactor) CreateProfile(opts *bind.TransactOpts, vars DataTypesCreateProfileData) (*types.Transaction, error) {
	return _Hub.contract.Transact(opts, "createProfile", vars)
}

// CreateProfile is a paid mutator transaction binding the contract method 0xffea138e.
//
// Solidity: function createProfile((address,string,string,address,bytes,string) vars) returns(uint256)
func (_Hub *HubSession) CreateProfile(vars DataTypesCreateProfileData) (*types.Transaction, error) {
	return _Hub.Contract.CreateProfile(&_Hub.TransactOpts, vars)
}

// CreateProfile is a paid mutator transaction binding the contract method 0xffea138e.
//
// Solidity: function createProfile((address,string,string,address,bytes,string) vars) returns(uint256)
func (_Hub *HubTransactorSession) CreateProfile(vars DataTypesCreateProfileData) (*types.Transaction, error) {
	return _Hub.Contract.CreateProfile(&_Hub.TransactOpts, vars)
}

// EmitCollectNFTTransferEvent is a paid mutator transaction binding the contract method 0x86e2947b.
//
// Solidity: function emitCollectNFTTransferEvent(uint256 profileId, uint256 pubId, uint256 collectNFTId, address from, address to) returns()
func (_Hub *HubTransactor) EmitCollectNFTTransferEvent(opts *bind.TransactOpts, profileId *big.Int, pubId *big.Int, collectNFTId *big.Int, from common.Address, to common.Address) (*types.Transaction, error) {
	return _Hub.contract.Transact(opts, "emitCollectNFTTransferEvent", profileId, pubId, collectNFTId, from, to)
}

// EmitCollectNFTTransferEvent is a paid mutator transaction binding the contract method 0x86e2947b.
//
// Solidity: function emitCollectNFTTransferEvent(uint256 profileId, uint256 pubId, uint256 collectNFTId, address from, address to) returns()
func (_Hub *HubSession) EmitCollectNFTTransferEvent(profileId *big.Int, pubId *big.Int, collectNFTId *big.Int, from common.Address, to common.Address) (*types.Transaction, error) {
	return _Hub.Contract.EmitCollectNFTTransferEvent(&_Hub.TransactOpts, profileId, pubId, collectNFTId, from, to)
}

// EmitCollectNFTTransferEvent is a paid mutator transaction binding the contract method 0x86e2947b.
//
// Solidity: function emitCollectNFTTransferEvent(uint256 profileId, uint256 pubId, uint256 collectNFTId, address from, address to) returns()
func (_Hub *HubTransactorSession) EmitCollectNFTTransferEvent(profileId *big.Int, pubId *big.Int, collectNFTId *big.Int, from common.Address, to common.Address) (*types.Transaction, error) {
	return _Hub.Contract.EmitCollectNFTTransferEvent(&_Hub.TransactOpts, profileId, pubId, collectNFTId, from, to)
}

// EmitFollowNFTTransferEvent is a paid mutator transaction binding the contract method 0xbdfdd4bc.
//
// Solidity: function emitFollowNFTTransferEvent(uint256 profileId, uint256 followNFTId, address from, address to) returns()
func (_Hub *HubTransactor) EmitFollowNFTTransferEvent(opts *bind.TransactOpts, profileId *big.Int, followNFTId *big.Int, from common.Address, to common.Address) (*types.Transaction, error) {
	return _Hub.contract.Transact(opts, "emitFollowNFTTransferEvent", profileId, followNFTId, from, to)
}

// EmitFollowNFTTransferEvent is a paid mutator transaction binding the contract method 0xbdfdd4bc.
//
// Solidity: function emitFollowNFTTransferEvent(uint256 profileId, uint256 followNFTId, address from, address to) returns()
func (_Hub *HubSession) EmitFollowNFTTransferEvent(profileId *big.Int, followNFTId *big.Int, from common.Address, to common.Address) (*types.Transaction, error) {
	return _Hub.Contract.EmitFollowNFTTransferEvent(&_Hub.TransactOpts, profileId, followNFTId, from, to)
}

// EmitFollowNFTTransferEvent is a paid mutator transaction binding the contract method 0xbdfdd4bc.
//
// Solidity: function emitFollowNFTTransferEvent(uint256 profileId, uint256 followNFTId, address from, address to) returns()
func (_Hub *HubTransactorSession) EmitFollowNFTTransferEvent(profileId *big.Int, followNFTId *big.Int, from common.Address, to common.Address) (*types.Transaction, error) {
	return _Hub.Contract.EmitFollowNFTTransferEvent(&_Hub.TransactOpts, profileId, followNFTId, from, to)
}

// Follow is a paid mutator transaction binding the contract method 0xfb78ae6c.
//
// Solidity: function follow(uint256[] profileIds, bytes[] datas) returns(uint256[])
func (_Hub *HubTransactor) Follow(opts *bind.TransactOpts, profileIds []*big.Int, datas [][]byte) (*types.Transaction, error) {
	return _Hub.contract.Transact(opts, "follow", profileIds, datas)
}

// Follow is a paid mutator transaction binding the contract method 0xfb78ae6c.
//
// Solidity: function follow(uint256[] profileIds, bytes[] datas) returns(uint256[])
func (_Hub *HubSession) Follow(profileIds []*big.Int, datas [][]byte) (*types.Transaction, error) {
	return _Hub.Contract.Follow(&_Hub.TransactOpts, profileIds, datas)
}

// Follow is a paid mutator transaction binding the contract method 0xfb78ae6c.
//
// Solidity: function follow(uint256[] profileIds, bytes[] datas) returns(uint256[])
func (_Hub *HubTransactorSession) Follow(profileIds []*big.Int, datas [][]byte) (*types.Transaction, error) {
	return _Hub.Contract.Follow(&_Hub.TransactOpts, profileIds, datas)
}

// FollowWithSig is a paid mutator transaction binding the contract method 0x8e4fd6a9.
//
// Solidity: function followWithSig((address,uint256[],bytes[],(uint8,bytes32,bytes32,uint256)) vars) returns(uint256[])
func (_Hub *HubTransactor) FollowWithSig(opts *bind.TransactOpts, vars DataTypesFollowWithSigData) (*types.Transaction, error) {
	return _Hub.contract.Transact(opts, "followWithSig", vars)
}

// FollowWithSig is a paid mutator transaction binding the contract method 0x8e4fd6a9.
//
// Solidity: function followWithSig((address,uint256[],bytes[],(uint8,bytes32,bytes32,uint256)) vars) returns(uint256[])
func (_Hub *HubSession) FollowWithSig(vars DataTypesFollowWithSigData) (*types.Transaction, error) {
	return _Hub.Contract.FollowWithSig(&_Hub.TransactOpts, vars)
}

// FollowWithSig is a paid mutator transaction binding the contract method 0x8e4fd6a9.
//
// Solidity: function followWithSig((address,uint256[],bytes[],(uint8,bytes32,bytes32,uint256)) vars) returns(uint256[])
func (_Hub *HubTransactorSession) FollowWithSig(vars DataTypesFollowWithSigData) (*types.Transaction, error) {
	return _Hub.Contract.FollowWithSig(&_Hub.TransactOpts, vars)
}

// Initialize is a paid mutator transaction binding the contract method 0x077f224a.
//
// Solidity: function initialize(string name, string symbol, address newGovernance) returns()
func (_Hub *HubTransactor) Initialize(opts *bind.TransactOpts, name string, symbol string, newGovernance common.Address) (*types.Transaction, error) {
	return _Hub.contract.Transact(opts, "initialize", name, symbol, newGovernance)
}

// Initialize is a paid mutator transaction binding the contract method 0x077f224a.
//
// Solidity: function initialize(string name, string symbol, address newGovernance) returns()
func (_Hub *HubSession) Initialize(name string, symbol string, newGovernance common.Address) (*types.Transaction, error) {
	return _Hub.Contract.Initialize(&_Hub.TransactOpts, name, symbol, newGovernance)
}

// Initialize is a paid mutator transaction binding the contract method 0x077f224a.
//
// Solidity: function initialize(string name, string symbol, address newGovernance) returns()
func (_Hub *HubTransactorSession) Initialize(name string, symbol string, newGovernance common.Address) (*types.Transaction, error) {
	return _Hub.Contract.Initialize(&_Hub.TransactOpts, name, symbol, newGovernance)
}

// Mirror is a paid mutator transaction binding the contract method 0x2faeed81.
//
// Solidity: function mirror((uint256,uint256,uint256,bytes,address,bytes) vars) returns(uint256)
func (_Hub *HubTransactor) Mirror(opts *bind.TransactOpts, vars DataTypesMirrorData) (*types.Transaction, error) {
	return _Hub.contract.Transact(opts, "mirror", vars)
}

// Mirror is a paid mutator transaction binding the contract method 0x2faeed81.
//
// Solidity: function mirror((uint256,uint256,uint256,bytes,address,bytes) vars) returns(uint256)
func (_Hub *HubSession) Mirror(vars DataTypesMirrorData) (*types.Transaction, error) {
	return _Hub.Contract.Mirror(&_Hub.TransactOpts, vars)
}

// Mirror is a paid mutator transaction binding the contract method 0x2faeed81.
//
// Solidity: function mirror((uint256,uint256,uint256,bytes,address,bytes) vars) returns(uint256)
func (_Hub *HubTransactorSession) Mirror(vars DataTypesMirrorData) (*types.Transaction, error) {
	return _Hub.Contract.Mirror(&_Hub.TransactOpts, vars)
}

// MirrorWithSig is a paid mutator transaction binding the contract method 0xdf457c34.
//
// Solidity: function mirrorWithSig((uint256,uint256,uint256,bytes,address,bytes,(uint8,bytes32,bytes32,uint256)) vars) returns(uint256)
func (_Hub *HubTransactor) MirrorWithSig(opts *bind.TransactOpts, vars DataTypesMirrorWithSigData) (*types.Transaction, error) {
	return _Hub.contract.Transact(opts, "mirrorWithSig", vars)
}

// MirrorWithSig is a paid mutator transaction binding the contract method 0xdf457c34.
//
// Solidity: function mirrorWithSig((uint256,uint256,uint256,bytes,address,bytes,(uint8,bytes32,bytes32,uint256)) vars) returns(uint256)
func (_Hub *HubSession) MirrorWithSig(vars DataTypesMirrorWithSigData) (*types.Transaction, error) {
	return _Hub.Contract.MirrorWithSig(&_Hub.TransactOpts, vars)
}

// MirrorWithSig is a paid mutator transaction binding the contract method 0xdf457c34.
//
// Solidity: function mirrorWithSig((uint256,uint256,uint256,bytes,address,bytes,(uint8,bytes32,bytes32,uint256)) vars) returns(uint256)
func (_Hub *HubTransactorSession) MirrorWithSig(vars DataTypesMirrorWithSigData) (*types.Transaction, error) {
	return _Hub.Contract.MirrorWithSig(&_Hub.TransactOpts, vars)
}

// Permit is a paid mutator transaction binding the contract method 0x7ef67f99.
//
// Solidity: function permit(address spender, uint256 tokenId, (uint8,bytes32,bytes32,uint256) sig) returns()
func (_Hub *HubTransactor) Permit(opts *bind.TransactOpts, spender common.Address, tokenId *big.Int, sig DataTypesEIP712Signature) (*types.Transaction, error) {
	return _Hub.contract.Transact(opts, "permit", spender, tokenId, sig)
}

// Permit is a paid mutator transaction binding the contract method 0x7ef67f99.
//
// Solidity: function permit(address spender, uint256 tokenId, (uint8,bytes32,bytes32,uint256) sig) returns()
func (_Hub *HubSession) Permit(spender common.Address, tokenId *big.Int, sig DataTypesEIP712Signature) (*types.Transaction, error) {
	return _Hub.Contract.Permit(&_Hub.TransactOpts, spender, tokenId, sig)
}

// Permit is a paid mutator transaction binding the contract method 0x7ef67f99.
//
// Solidity: function permit(address spender, uint256 tokenId, (uint8,bytes32,bytes32,uint256) sig) returns()
func (_Hub *HubTransactorSession) Permit(spender common.Address, tokenId *big.Int, sig DataTypesEIP712Signature) (*types.Transaction, error) {
	return _Hub.Contract.Permit(&_Hub.TransactOpts, spender, tokenId, sig)
}

// PermitForAll is a paid mutator transaction binding the contract method 0x89028a13.
//
// Solidity: function permitForAll(address owner, address operator, bool approved, (uint8,bytes32,bytes32,uint256) sig) returns()
func (_Hub *HubTransactor) PermitForAll(opts *bind.TransactOpts, owner common.Address, operator common.Address, approved bool, sig DataTypesEIP712Signature) (*types.Transaction, error) {
	return _Hub.contract.Transact(opts, "permitForAll", owner, operator, approved, sig)
}

// PermitForAll is a paid mutator transaction binding the contract method 0x89028a13.
//
// Solidity: function permitForAll(address owner, address operator, bool approved, (uint8,bytes32,bytes32,uint256) sig) returns()
func (_Hub *HubSession) PermitForAll(owner common.Address, operator common.Address, approved bool, sig DataTypesEIP712Signature) (*types.Transaction, error) {
	return _Hub.Contract.PermitForAll(&_Hub.TransactOpts, owner, operator, approved, sig)
}

// PermitForAll is a paid mutator transaction binding the contract method 0x89028a13.
//
// Solidity: function permitForAll(address owner, address operator, bool approved, (uint8,bytes32,bytes32,uint256) sig) returns()
func (_Hub *HubTransactorSession) PermitForAll(owner common.Address, operator common.Address, approved bool, sig DataTypesEIP712Signature) (*types.Transaction, error) {
	return _Hub.Contract.PermitForAll(&_Hub.TransactOpts, owner, operator, approved, sig)
}

// Post is a paid mutator transaction binding the contract method 0x963ff141.
//
// Solidity: function post((uint256,string,address,bytes,address,bytes) vars) returns(uint256)
func (_Hub *HubTransactor) Post(opts *bind.TransactOpts, vars DataTypesPostData) (*types.Transaction, error) {
	return _Hub.contract.Transact(opts, "post", vars)
}

// Post is a paid mutator transaction binding the contract method 0x963ff141.
//
// Solidity: function post((uint256,string,address,bytes,address,bytes) vars) returns(uint256)
func (_Hub *HubSession) Post(vars DataTypesPostData) (*types.Transaction, error) {
	return _Hub.Contract.Post(&_Hub.TransactOpts, vars)
}

// Post is a paid mutator transaction binding the contract method 0x963ff141.
//
// Solidity: function post((uint256,string,address,bytes,address,bytes) vars) returns(uint256)
func (_Hub *HubTransactorSession) Post(vars DataTypesPostData) (*types.Transaction, error) {
	return _Hub.Contract.Post(&_Hub.TransactOpts, vars)
}

// PostWithSig is a paid mutator transaction binding the contract method 0x3b508132.
//
// Solidity: function postWithSig((uint256,string,address,bytes,address,bytes,(uint8,bytes32,bytes32,uint256)) vars) returns(uint256)
func (_Hub *HubTransactor) PostWithSig(opts *bind.TransactOpts, vars DataTypesPostWithSigData) (*types.Transaction, error) {
	return _Hub.contract.Transact(opts, "postWithSig", vars)
}

// PostWithSig is a paid mutator transaction binding the contract method 0x3b508132.
//
// Solidity: function postWithSig((uint256,string,address,bytes,address,bytes,(uint8,bytes32,bytes32,uint256)) vars) returns(uint256)
func (_Hub *HubSession) PostWithSig(vars DataTypesPostWithSigData) (*types.Transaction, error) {
	return _Hub.Contract.PostWithSig(&_Hub.TransactOpts, vars)
}

// PostWithSig is a paid mutator transaction binding the contract method 0x3b508132.
//
// Solidity: function postWithSig((uint256,string,address,bytes,address,bytes,(uint8,bytes32,bytes32,uint256)) vars) returns(uint256)
func (_Hub *HubTransactorSession) PostWithSig(vars DataTypesPostWithSigData) (*types.Transaction, error) {
	return _Hub.Contract.PostWithSig(&_Hub.TransactOpts, vars)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Hub *HubTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Hub.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Hub *HubSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Hub.Contract.SafeTransferFrom(&_Hub.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Hub *HubTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Hub.Contract.SafeTransferFrom(&_Hub.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Hub *HubTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Hub.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Hub *HubSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Hub.Contract.SafeTransferFrom0(&_Hub.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Hub *HubTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Hub.Contract.SafeTransferFrom0(&_Hub.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Hub *HubTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Hub.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Hub *HubSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Hub.Contract.SetApprovalForAll(&_Hub.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Hub *HubTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Hub.Contract.SetApprovalForAll(&_Hub.TransactOpts, operator, approved)
}

// SetDefaultProfile is a paid mutator transaction binding the contract method 0xf1b2f8bc.
//
// Solidity: function setDefaultProfile(uint256 profileId) returns()
func (_Hub *HubTransactor) SetDefaultProfile(opts *bind.TransactOpts, profileId *big.Int) (*types.Transaction, error) {
	return _Hub.contract.Transact(opts, "setDefaultProfile", profileId)
}

// SetDefaultProfile is a paid mutator transaction binding the contract method 0xf1b2f8bc.
//
// Solidity: function setDefaultProfile(uint256 profileId) returns()
func (_Hub *HubSession) SetDefaultProfile(profileId *big.Int) (*types.Transaction, error) {
	return _Hub.Contract.SetDefaultProfile(&_Hub.TransactOpts, profileId)
}

// SetDefaultProfile is a paid mutator transaction binding the contract method 0xf1b2f8bc.
//
// Solidity: function setDefaultProfile(uint256 profileId) returns()
func (_Hub *HubTransactorSession) SetDefaultProfile(profileId *big.Int) (*types.Transaction, error) {
	return _Hub.Contract.SetDefaultProfile(&_Hub.TransactOpts, profileId)
}

// SetDefaultProfileWithSig is a paid mutator transaction binding the contract method 0xdc217253.
//
// Solidity: function setDefaultProfileWithSig((address,uint256,(uint8,bytes32,bytes32,uint256)) vars) returns()
func (_Hub *HubTransactor) SetDefaultProfileWithSig(opts *bind.TransactOpts, vars DataTypesSetDefaultProfileWithSigData) (*types.Transaction, error) {
	return _Hub.contract.Transact(opts, "setDefaultProfileWithSig", vars)
}

// SetDefaultProfileWithSig is a paid mutator transaction binding the contract method 0xdc217253.
//
// Solidity: function setDefaultProfileWithSig((address,uint256,(uint8,bytes32,bytes32,uint256)) vars) returns()
func (_Hub *HubSession) SetDefaultProfileWithSig(vars DataTypesSetDefaultProfileWithSigData) (*types.Transaction, error) {
	return _Hub.Contract.SetDefaultProfileWithSig(&_Hub.TransactOpts, vars)
}

// SetDefaultProfileWithSig is a paid mutator transaction binding the contract method 0xdc217253.
//
// Solidity: function setDefaultProfileWithSig((address,uint256,(uint8,bytes32,bytes32,uint256)) vars) returns()
func (_Hub *HubTransactorSession) SetDefaultProfileWithSig(vars DataTypesSetDefaultProfileWithSigData) (*types.Transaction, error) {
	return _Hub.Contract.SetDefaultProfileWithSig(&_Hub.TransactOpts, vars)
}

// SetDispatcher is a paid mutator transaction binding the contract method 0xbfd24f47.
//
// Solidity: function setDispatcher(uint256 profileId, address dispatcher) returns()
func (_Hub *HubTransactor) SetDispatcher(opts *bind.TransactOpts, profileId *big.Int, dispatcher common.Address) (*types.Transaction, error) {
	return _Hub.contract.Transact(opts, "setDispatcher", profileId, dispatcher)
}

// SetDispatcher is a paid mutator transaction binding the contract method 0xbfd24f47.
//
// Solidity: function setDispatcher(uint256 profileId, address dispatcher) returns()
func (_Hub *HubSession) SetDispatcher(profileId *big.Int, dispatcher common.Address) (*types.Transaction, error) {
	return _Hub.Contract.SetDispatcher(&_Hub.TransactOpts, profileId, dispatcher)
}

// SetDispatcher is a paid mutator transaction binding the contract method 0xbfd24f47.
//
// Solidity: function setDispatcher(uint256 profileId, address dispatcher) returns()
func (_Hub *HubTransactorSession) SetDispatcher(profileId *big.Int, dispatcher common.Address) (*types.Transaction, error) {
	return _Hub.Contract.SetDispatcher(&_Hub.TransactOpts, profileId, dispatcher)
}

// SetDispatcherWithSig is a paid mutator transaction binding the contract method 0xbfbf0b4b.
//
// Solidity: function setDispatcherWithSig((uint256,address,(uint8,bytes32,bytes32,uint256)) vars) returns()
func (_Hub *HubTransactor) SetDispatcherWithSig(opts *bind.TransactOpts, vars DataTypesSetDispatcherWithSigData) (*types.Transaction, error) {
	return _Hub.contract.Transact(opts, "setDispatcherWithSig", vars)
}

// SetDispatcherWithSig is a paid mutator transaction binding the contract method 0xbfbf0b4b.
//
// Solidity: function setDispatcherWithSig((uint256,address,(uint8,bytes32,bytes32,uint256)) vars) returns()
func (_Hub *HubSession) SetDispatcherWithSig(vars DataTypesSetDispatcherWithSigData) (*types.Transaction, error) {
	return _Hub.Contract.SetDispatcherWithSig(&_Hub.TransactOpts, vars)
}

// SetDispatcherWithSig is a paid mutator transaction binding the contract method 0xbfbf0b4b.
//
// Solidity: function setDispatcherWithSig((uint256,address,(uint8,bytes32,bytes32,uint256)) vars) returns()
func (_Hub *HubTransactorSession) SetDispatcherWithSig(vars DataTypesSetDispatcherWithSigData) (*types.Transaction, error) {
	return _Hub.Contract.SetDispatcherWithSig(&_Hub.TransactOpts, vars)
}

// SetEmergencyAdmin is a paid mutator transaction binding the contract method 0x35da3394.
//
// Solidity: function setEmergencyAdmin(address newEmergencyAdmin) returns()
func (_Hub *HubTransactor) SetEmergencyAdmin(opts *bind.TransactOpts, newEmergencyAdmin common.Address) (*types.Transaction, error) {
	return _Hub.contract.Transact(opts, "setEmergencyAdmin", newEmergencyAdmin)
}

// SetEmergencyAdmin is a paid mutator transaction binding the contract method 0x35da3394.
//
// Solidity: function setEmergencyAdmin(address newEmergencyAdmin) returns()
func (_Hub *HubSession) SetEmergencyAdmin(newEmergencyAdmin common.Address) (*types.Transaction, error) {
	return _Hub.Contract.SetEmergencyAdmin(&_Hub.TransactOpts, newEmergencyAdmin)
}

// SetEmergencyAdmin is a paid mutator transaction binding the contract method 0x35da3394.
//
// Solidity: function setEmergencyAdmin(address newEmergencyAdmin) returns()
func (_Hub *HubTransactorSession) SetEmergencyAdmin(newEmergencyAdmin common.Address) (*types.Transaction, error) {
	return _Hub.Contract.SetEmergencyAdmin(&_Hub.TransactOpts, newEmergencyAdmin)
}

// SetFollowModule is a paid mutator transaction binding the contract method 0x6dea40b3.
//
// Solidity: function setFollowModule(uint256 profileId, address followModule, bytes followModuleInitData) returns()
func (_Hub *HubTransactor) SetFollowModule(opts *bind.TransactOpts, profileId *big.Int, followModule common.Address, followModuleInitData []byte) (*types.Transaction, error) {
	return _Hub.contract.Transact(opts, "setFollowModule", profileId, followModule, followModuleInitData)
}

// SetFollowModule is a paid mutator transaction binding the contract method 0x6dea40b3.
//
// Solidity: function setFollowModule(uint256 profileId, address followModule, bytes followModuleInitData) returns()
func (_Hub *HubSession) SetFollowModule(profileId *big.Int, followModule common.Address, followModuleInitData []byte) (*types.Transaction, error) {
	return _Hub.Contract.SetFollowModule(&_Hub.TransactOpts, profileId, followModule, followModuleInitData)
}

// SetFollowModule is a paid mutator transaction binding the contract method 0x6dea40b3.
//
// Solidity: function setFollowModule(uint256 profileId, address followModule, bytes followModuleInitData) returns()
func (_Hub *HubTransactorSession) SetFollowModule(profileId *big.Int, followModule common.Address, followModuleInitData []byte) (*types.Transaction, error) {
	return _Hub.Contract.SetFollowModule(&_Hub.TransactOpts, profileId, followModule, followModuleInitData)
}

// SetFollowModuleWithSig is a paid mutator transaction binding the contract method 0x3b28b89f.
//
// Solidity: function setFollowModuleWithSig((uint256,address,bytes,(uint8,bytes32,bytes32,uint256)) vars) returns()
func (_Hub *HubTransactor) SetFollowModuleWithSig(opts *bind.TransactOpts, vars DataTypesSetFollowModuleWithSigData) (*types.Transaction, error) {
	return _Hub.contract.Transact(opts, "setFollowModuleWithSig", vars)
}

// SetFollowModuleWithSig is a paid mutator transaction binding the contract method 0x3b28b89f.
//
// Solidity: function setFollowModuleWithSig((uint256,address,bytes,(uint8,bytes32,bytes32,uint256)) vars) returns()
func (_Hub *HubSession) SetFollowModuleWithSig(vars DataTypesSetFollowModuleWithSigData) (*types.Transaction, error) {
	return _Hub.Contract.SetFollowModuleWithSig(&_Hub.TransactOpts, vars)
}

// SetFollowModuleWithSig is a paid mutator transaction binding the contract method 0x3b28b89f.
//
// Solidity: function setFollowModuleWithSig((uint256,address,bytes,(uint8,bytes32,bytes32,uint256)) vars) returns()
func (_Hub *HubTransactorSession) SetFollowModuleWithSig(vars DataTypesSetFollowModuleWithSigData) (*types.Transaction, error) {
	return _Hub.Contract.SetFollowModuleWithSig(&_Hub.TransactOpts, vars)
}

// SetFollowNFTURI is a paid mutator transaction binding the contract method 0xc6b5d06c.
//
// Solidity: function setFollowNFTURI(uint256 profileId, string followNFTURI) returns()
func (_Hub *HubTransactor) SetFollowNFTURI(opts *bind.TransactOpts, profileId *big.Int, followNFTURI string) (*types.Transaction, error) {
	return _Hub.contract.Transact(opts, "setFollowNFTURI", profileId, followNFTURI)
}

// SetFollowNFTURI is a paid mutator transaction binding the contract method 0xc6b5d06c.
//
// Solidity: function setFollowNFTURI(uint256 profileId, string followNFTURI) returns()
func (_Hub *HubSession) SetFollowNFTURI(profileId *big.Int, followNFTURI string) (*types.Transaction, error) {
	return _Hub.Contract.SetFollowNFTURI(&_Hub.TransactOpts, profileId, followNFTURI)
}

// SetFollowNFTURI is a paid mutator transaction binding the contract method 0xc6b5d06c.
//
// Solidity: function setFollowNFTURI(uint256 profileId, string followNFTURI) returns()
func (_Hub *HubTransactorSession) SetFollowNFTURI(profileId *big.Int, followNFTURI string) (*types.Transaction, error) {
	return _Hub.Contract.SetFollowNFTURI(&_Hub.TransactOpts, profileId, followNFTURI)
}

// SetFollowNFTURIWithSig is a paid mutator transaction binding the contract method 0xbd12d3f0.
//
// Solidity: function setFollowNFTURIWithSig((uint256,string,(uint8,bytes32,bytes32,uint256)) vars) returns()
func (_Hub *HubTransactor) SetFollowNFTURIWithSig(opts *bind.TransactOpts, vars DataTypesSetFollowNFTURIWithSigData) (*types.Transaction, error) {
	return _Hub.contract.Transact(opts, "setFollowNFTURIWithSig", vars)
}

// SetFollowNFTURIWithSig is a paid mutator transaction binding the contract method 0xbd12d3f0.
//
// Solidity: function setFollowNFTURIWithSig((uint256,string,(uint8,bytes32,bytes32,uint256)) vars) returns()
func (_Hub *HubSession) SetFollowNFTURIWithSig(vars DataTypesSetFollowNFTURIWithSigData) (*types.Transaction, error) {
	return _Hub.Contract.SetFollowNFTURIWithSig(&_Hub.TransactOpts, vars)
}

// SetFollowNFTURIWithSig is a paid mutator transaction binding the contract method 0xbd12d3f0.
//
// Solidity: function setFollowNFTURIWithSig((uint256,string,(uint8,bytes32,bytes32,uint256)) vars) returns()
func (_Hub *HubTransactorSession) SetFollowNFTURIWithSig(vars DataTypesSetFollowNFTURIWithSigData) (*types.Transaction, error) {
	return _Hub.Contract.SetFollowNFTURIWithSig(&_Hub.TransactOpts, vars)
}

// SetGovernance is a paid mutator transaction binding the contract method 0xab033ea9.
//
// Solidity: function setGovernance(address newGovernance) returns()
func (_Hub *HubTransactor) SetGovernance(opts *bind.TransactOpts, newGovernance common.Address) (*types.Transaction, error) {
	return _Hub.contract.Transact(opts, "setGovernance", newGovernance)
}

// SetGovernance is a paid mutator transaction binding the contract method 0xab033ea9.
//
// Solidity: function setGovernance(address newGovernance) returns()
func (_Hub *HubSession) SetGovernance(newGovernance common.Address) (*types.Transaction, error) {
	return _Hub.Contract.SetGovernance(&_Hub.TransactOpts, newGovernance)
}

// SetGovernance is a paid mutator transaction binding the contract method 0xab033ea9.
//
// Solidity: function setGovernance(address newGovernance) returns()
func (_Hub *HubTransactorSession) SetGovernance(newGovernance common.Address) (*types.Transaction, error) {
	return _Hub.Contract.SetGovernance(&_Hub.TransactOpts, newGovernance)
}

// SetProfileImageURI is a paid mutator transaction binding the contract method 0x054871b8.
//
// Solidity: function setProfileImageURI(uint256 profileId, string imageURI) returns()
func (_Hub *HubTransactor) SetProfileImageURI(opts *bind.TransactOpts, profileId *big.Int, imageURI string) (*types.Transaction, error) {
	return _Hub.contract.Transact(opts, "setProfileImageURI", profileId, imageURI)
}

// SetProfileImageURI is a paid mutator transaction binding the contract method 0x054871b8.
//
// Solidity: function setProfileImageURI(uint256 profileId, string imageURI) returns()
func (_Hub *HubSession) SetProfileImageURI(profileId *big.Int, imageURI string) (*types.Transaction, error) {
	return _Hub.Contract.SetProfileImageURI(&_Hub.TransactOpts, profileId, imageURI)
}

// SetProfileImageURI is a paid mutator transaction binding the contract method 0x054871b8.
//
// Solidity: function setProfileImageURI(uint256 profileId, string imageURI) returns()
func (_Hub *HubTransactorSession) SetProfileImageURI(profileId *big.Int, imageURI string) (*types.Transaction, error) {
	return _Hub.Contract.SetProfileImageURI(&_Hub.TransactOpts, profileId, imageURI)
}

// SetProfileImageURIWithSig is a paid mutator transaction binding the contract method 0x9b84a14c.
//
// Solidity: function setProfileImageURIWithSig((uint256,string,(uint8,bytes32,bytes32,uint256)) vars) returns()
func (_Hub *HubTransactor) SetProfileImageURIWithSig(opts *bind.TransactOpts, vars DataTypesSetProfileImageURIWithSigData) (*types.Transaction, error) {
	return _Hub.contract.Transact(opts, "setProfileImageURIWithSig", vars)
}

// SetProfileImageURIWithSig is a paid mutator transaction binding the contract method 0x9b84a14c.
//
// Solidity: function setProfileImageURIWithSig((uint256,string,(uint8,bytes32,bytes32,uint256)) vars) returns()
func (_Hub *HubSession) SetProfileImageURIWithSig(vars DataTypesSetProfileImageURIWithSigData) (*types.Transaction, error) {
	return _Hub.Contract.SetProfileImageURIWithSig(&_Hub.TransactOpts, vars)
}

// SetProfileImageURIWithSig is a paid mutator transaction binding the contract method 0x9b84a14c.
//
// Solidity: function setProfileImageURIWithSig((uint256,string,(uint8,bytes32,bytes32,uint256)) vars) returns()
func (_Hub *HubTransactorSession) SetProfileImageURIWithSig(vars DataTypesSetProfileImageURIWithSigData) (*types.Transaction, error) {
	return _Hub.Contract.SetProfileImageURIWithSig(&_Hub.TransactOpts, vars)
}

// SetState is a paid mutator transaction binding the contract method 0x56de96db.
//
// Solidity: function setState(uint8 newState) returns()
func (_Hub *HubTransactor) SetState(opts *bind.TransactOpts, newState uint8) (*types.Transaction, error) {
	return _Hub.contract.Transact(opts, "setState", newState)
}

// SetState is a paid mutator transaction binding the contract method 0x56de96db.
//
// Solidity: function setState(uint8 newState) returns()
func (_Hub *HubSession) SetState(newState uint8) (*types.Transaction, error) {
	return _Hub.Contract.SetState(&_Hub.TransactOpts, newState)
}

// SetState is a paid mutator transaction binding the contract method 0x56de96db.
//
// Solidity: function setState(uint8 newState) returns()
func (_Hub *HubTransactorSession) SetState(newState uint8) (*types.Transaction, error) {
	return _Hub.Contract.SetState(&_Hub.TransactOpts, newState)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Hub *HubTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Hub.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Hub *HubSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Hub.Contract.TransferFrom(&_Hub.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Hub *HubTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Hub.Contract.TransferFrom(&_Hub.TransactOpts, from, to, tokenId)
}

// WhitelistCollectModule is a paid mutator transaction binding the contract method 0x31dcebe3.
//
// Solidity: function whitelistCollectModule(address collectModule, bool whitelist) returns()
func (_Hub *HubTransactor) WhitelistCollectModule(opts *bind.TransactOpts, collectModule common.Address, whitelist bool) (*types.Transaction, error) {
	return _Hub.contract.Transact(opts, "whitelistCollectModule", collectModule, whitelist)
}

// WhitelistCollectModule is a paid mutator transaction binding the contract method 0x31dcebe3.
//
// Solidity: function whitelistCollectModule(address collectModule, bool whitelist) returns()
func (_Hub *HubSession) WhitelistCollectModule(collectModule common.Address, whitelist bool) (*types.Transaction, error) {
	return _Hub.Contract.WhitelistCollectModule(&_Hub.TransactOpts, collectModule, whitelist)
}

// WhitelistCollectModule is a paid mutator transaction binding the contract method 0x31dcebe3.
//
// Solidity: function whitelistCollectModule(address collectModule, bool whitelist) returns()
func (_Hub *HubTransactorSession) WhitelistCollectModule(collectModule common.Address, whitelist bool) (*types.Transaction, error) {
	return _Hub.Contract.WhitelistCollectModule(&_Hub.TransactOpts, collectModule, whitelist)
}

// WhitelistFollowModule is a paid mutator transaction binding the contract method 0xa6d8e1e4.
//
// Solidity: function whitelistFollowModule(address followModule, bool whitelist) returns()
func (_Hub *HubTransactor) WhitelistFollowModule(opts *bind.TransactOpts, followModule common.Address, whitelist bool) (*types.Transaction, error) {
	return _Hub.contract.Transact(opts, "whitelistFollowModule", followModule, whitelist)
}

// WhitelistFollowModule is a paid mutator transaction binding the contract method 0xa6d8e1e4.
//
// Solidity: function whitelistFollowModule(address followModule, bool whitelist) returns()
func (_Hub *HubSession) WhitelistFollowModule(followModule common.Address, whitelist bool) (*types.Transaction, error) {
	return _Hub.Contract.WhitelistFollowModule(&_Hub.TransactOpts, followModule, whitelist)
}

// WhitelistFollowModule is a paid mutator transaction binding the contract method 0xa6d8e1e4.
//
// Solidity: function whitelistFollowModule(address followModule, bool whitelist) returns()
func (_Hub *HubTransactorSession) WhitelistFollowModule(followModule common.Address, whitelist bool) (*types.Transaction, error) {
	return _Hub.Contract.WhitelistFollowModule(&_Hub.TransactOpts, followModule, whitelist)
}

// WhitelistProfileCreator is a paid mutator transaction binding the contract method 0x20905506.
//
// Solidity: function whitelistProfileCreator(address profileCreator, bool whitelist) returns()
func (_Hub *HubTransactor) WhitelistProfileCreator(opts *bind.TransactOpts, profileCreator common.Address, whitelist bool) (*types.Transaction, error) {
	return _Hub.contract.Transact(opts, "whitelistProfileCreator", profileCreator, whitelist)
}

// WhitelistProfileCreator is a paid mutator transaction binding the contract method 0x20905506.
//
// Solidity: function whitelistProfileCreator(address profileCreator, bool whitelist) returns()
func (_Hub *HubSession) WhitelistProfileCreator(profileCreator common.Address, whitelist bool) (*types.Transaction, error) {
	return _Hub.Contract.WhitelistProfileCreator(&_Hub.TransactOpts, profileCreator, whitelist)
}

// WhitelistProfileCreator is a paid mutator transaction binding the contract method 0x20905506.
//
// Solidity: function whitelistProfileCreator(address profileCreator, bool whitelist) returns()
func (_Hub *HubTransactorSession) WhitelistProfileCreator(profileCreator common.Address, whitelist bool) (*types.Transaction, error) {
	return _Hub.Contract.WhitelistProfileCreator(&_Hub.TransactOpts, profileCreator, whitelist)
}

// WhitelistReferenceModule is a paid mutator transaction binding the contract method 0x4187e4c5.
//
// Solidity: function whitelistReferenceModule(address referenceModule, bool whitelist) returns()
func (_Hub *HubTransactor) WhitelistReferenceModule(opts *bind.TransactOpts, referenceModule common.Address, whitelist bool) (*types.Transaction, error) {
	return _Hub.contract.Transact(opts, "whitelistReferenceModule", referenceModule, whitelist)
}

// WhitelistReferenceModule is a paid mutator transaction binding the contract method 0x4187e4c5.
//
// Solidity: function whitelistReferenceModule(address referenceModule, bool whitelist) returns()
func (_Hub *HubSession) WhitelistReferenceModule(referenceModule common.Address, whitelist bool) (*types.Transaction, error) {
	return _Hub.Contract.WhitelistReferenceModule(&_Hub.TransactOpts, referenceModule, whitelist)
}

// WhitelistReferenceModule is a paid mutator transaction binding the contract method 0x4187e4c5.
//
// Solidity: function whitelistReferenceModule(address referenceModule, bool whitelist) returns()
func (_Hub *HubTransactorSession) WhitelistReferenceModule(referenceModule common.Address, whitelist bool) (*types.Transaction, error) {
	return _Hub.Contract.WhitelistReferenceModule(&_Hub.TransactOpts, referenceModule, whitelist)
}

// HubApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Hub contract.
type HubApprovalIterator struct {
	Event *HubApproval // Event containing the contract specifics and raw log

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
func (it *HubApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HubApproval)
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
		it.Event = new(HubApproval)
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
func (it *HubApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HubApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HubApproval represents a Approval event raised by the Hub contract.
type HubApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Hub *HubFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*HubApprovalIterator, error) {

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

	logs, sub, err := _Hub.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &HubApprovalIterator{contract: _Hub.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Hub *HubFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *HubApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Hub.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HubApproval)
				if err := _Hub.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Hub *HubFilterer) ParseApproval(log types.Log) (*HubApproval, error) {
	event := new(HubApproval)
	if err := _Hub.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HubApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Hub contract.
type HubApprovalForAllIterator struct {
	Event *HubApprovalForAll // Event containing the contract specifics and raw log

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
func (it *HubApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HubApprovalForAll)
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
		it.Event = new(HubApprovalForAll)
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
func (it *HubApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HubApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HubApprovalForAll represents a ApprovalForAll event raised by the Hub contract.
type HubApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Hub *HubFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*HubApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Hub.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &HubApprovalForAllIterator{contract: _Hub.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Hub *HubFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *HubApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Hub.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HubApprovalForAll)
				if err := _Hub.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_Hub *HubFilterer) ParseApprovalForAll(log types.Log) (*HubApprovalForAll, error) {
	event := new(HubApprovalForAll)
	if err := _Hub.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HubTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Hub contract.
type HubTransferIterator struct {
	Event *HubTransfer // Event containing the contract specifics and raw log

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
func (it *HubTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HubTransfer)
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
		it.Event = new(HubTransfer)
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
func (it *HubTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HubTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HubTransfer represents a Transfer event raised by the Hub contract.
type HubTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Hub *HubFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*HubTransferIterator, error) {

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

	logs, sub, err := _Hub.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &HubTransferIterator{contract: _Hub.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Hub *HubFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *HubTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Hub.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HubTransfer)
				if err := _Hub.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Hub *HubFilterer) ParseTransfer(log types.Log) (*HubTransfer, error) {
	event := new(HubTransfer)
	if err := _Hub.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
