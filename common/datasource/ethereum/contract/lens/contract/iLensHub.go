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

// ILensHubMetaData contains all meta data concerning the ILensHub contract.
var ILensHubMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pubId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"collect\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"collector\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pubId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.EIP712Signature\",\"name\":\"sig\",\"type\":\"tuple\"}],\"internalType\":\"structDataTypes.CollectWithSigData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"collectWithSig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contentURI\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"profileIdPointed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pubIdPointed\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"referenceModuleData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"collectModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"collectModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"referenceModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"referenceModuleInitData\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.CommentData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"comment\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contentURI\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"profileIdPointed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pubIdPointed\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"referenceModuleData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"collectModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"collectModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"referenceModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"referenceModuleInitData\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.EIP712Signature\",\"name\":\"sig\",\"type\":\"tuple\"}],\"internalType\":\"structDataTypes.CommentWithSigData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"commentWithSig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"handle\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"imageURI\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"followModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"followModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"followNFTURI\",\"type\":\"string\"}],\"internalType\":\"structDataTypes.CreateProfileData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"createProfile\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"defaultProfile\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pubId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collectNFTId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"emitCollectNFTTransferEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"followNFTId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"emitFollowNFTTransferEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"profileIds\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"datas\",\"type\":\"bytes[]\"}],\"name\":\"follow\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"follower\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"profileIds\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"datas\",\"type\":\"bytes[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.EIP712Signature\",\"name\":\"sig\",\"type\":\"tuple\"}],\"internalType\":\"structDataTypes.FollowWithSigData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"followWithSig\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pubId\",\"type\":\"uint256\"}],\"name\":\"getCollectModule\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pubId\",\"type\":\"uint256\"}],\"name\":\"getCollectNFT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCollectNFTImpl\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pubId\",\"type\":\"uint256\"}],\"name\":\"getContentURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"}],\"name\":\"getDispatcher\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"}],\"name\":\"getFollowModule\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"}],\"name\":\"getFollowNFT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFollowNFTImpl\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"}],\"name\":\"getFollowNFTURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGovernance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"}],\"name\":\"getHandle\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"}],\"name\":\"getProfile\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pubCount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"followModule\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"followNFT\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"handle\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"imageURI\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"followNFTURI\",\"type\":\"string\"}],\"internalType\":\"structDataTypes.ProfileStruct\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"handle\",\"type\":\"string\"}],\"name\":\"getProfileIdByHandle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pubId\",\"type\":\"uint256\"}],\"name\":\"getPub\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"profileIdPointed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pubIdPointed\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contentURI\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"referenceModule\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collectModule\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collectNFT\",\"type\":\"address\"}],\"internalType\":\"structDataTypes.PublicationStruct\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"}],\"name\":\"getPubCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pubId\",\"type\":\"uint256\"}],\"name\":\"getPubPointer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pubId\",\"type\":\"uint256\"}],\"name\":\"getPubType\",\"outputs\":[{\"internalType\":\"enumDataTypes.PubType\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pubId\",\"type\":\"uint256\"}],\"name\":\"getReferenceModule\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"newGovernance\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collectModule\",\"type\":\"address\"}],\"name\":\"isCollectModuleWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"followModule\",\"type\":\"address\"}],\"name\":\"isFollowModuleWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"profileCreator\",\"type\":\"address\"}],\"name\":\"isProfileCreatorWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"referenceModule\",\"type\":\"address\"}],\"name\":\"isReferenceModuleWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"profileIdPointed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pubIdPointed\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"referenceModuleData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"referenceModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"referenceModuleInitData\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.MirrorData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"mirror\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"profileIdPointed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pubIdPointed\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"referenceModuleData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"referenceModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"referenceModuleInitData\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.EIP712Signature\",\"name\":\"sig\",\"type\":\"tuple\"}],\"internalType\":\"structDataTypes.MirrorWithSigData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"mirrorWithSig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contentURI\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"collectModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"collectModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"referenceModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"referenceModuleInitData\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.PostData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"post\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contentURI\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"collectModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"collectModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"referenceModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"referenceModuleInitData\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.EIP712Signature\",\"name\":\"sig\",\"type\":\"tuple\"}],\"internalType\":\"structDataTypes.PostWithSigData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"postWithSig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"}],\"name\":\"setDefaultProfile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.EIP712Signature\",\"name\":\"sig\",\"type\":\"tuple\"}],\"internalType\":\"structDataTypes.SetDefaultProfileWithSigData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"setDefaultProfileWithSig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"dispatcher\",\"type\":\"address\"}],\"name\":\"setDispatcher\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"dispatcher\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.EIP712Signature\",\"name\":\"sig\",\"type\":\"tuple\"}],\"internalType\":\"structDataTypes.SetDispatcherWithSigData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"setDispatcherWithSig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newEmergencyAdmin\",\"type\":\"address\"}],\"name\":\"setEmergencyAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"followModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"followModuleInitData\",\"type\":\"bytes\"}],\"name\":\"setFollowModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"followModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"followModuleInitData\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.EIP712Signature\",\"name\":\"sig\",\"type\":\"tuple\"}],\"internalType\":\"structDataTypes.SetFollowModuleWithSigData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"setFollowModuleWithSig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"followNFTURI\",\"type\":\"string\"}],\"name\":\"setFollowNFTURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"followNFTURI\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.EIP712Signature\",\"name\":\"sig\",\"type\":\"tuple\"}],\"internalType\":\"structDataTypes.SetFollowNFTURIWithSigData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"setFollowNFTURIWithSig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newGovernance\",\"type\":\"address\"}],\"name\":\"setGovernance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"imageURI\",\"type\":\"string\"}],\"name\":\"setProfileImageURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"imageURI\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.EIP712Signature\",\"name\":\"sig\",\"type\":\"tuple\"}],\"internalType\":\"structDataTypes.SetProfileImageURIWithSigData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"setProfileImageURIWithSig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumDataTypes.ProtocolState\",\"name\":\"newState\",\"type\":\"uint8\"}],\"name\":\"setState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collectModule\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"whitelist\",\"type\":\"bool\"}],\"name\":\"whitelistCollectModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"followModule\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"whitelist\",\"type\":\"bool\"}],\"name\":\"whitelistFollowModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"profileCreator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"whitelist\",\"type\":\"bool\"}],\"name\":\"whitelistProfileCreator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"referenceModule\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"whitelist\",\"type\":\"bool\"}],\"name\":\"whitelistReferenceModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ILensHubABI is the input ABI used to generate the binding from.
// Deprecated: Use ILensHubMetaData.ABI instead.
var ILensHubABI = ILensHubMetaData.ABI

// ILensHub is an auto generated Go binding around an Ethereum contract.
type ILensHub struct {
	ILensHubCaller     // Read-only binding to the contract
	ILensHubTransactor // Write-only binding to the contract
	ILensHubFilterer   // Log filterer for contract events
}

// ILensHubCaller is an auto generated read-only Go binding around an Ethereum contract.
type ILensHubCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILensHubTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ILensHubTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILensHubFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ILensHubFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILensHubSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ILensHubSession struct {
	Contract     *ILensHub         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ILensHubCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ILensHubCallerSession struct {
	Contract *ILensHubCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ILensHubTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ILensHubTransactorSession struct {
	Contract     *ILensHubTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ILensHubRaw is an auto generated low-level Go binding around an Ethereum contract.
type ILensHubRaw struct {
	Contract *ILensHub // Generic contract binding to access the raw methods on
}

// ILensHubCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ILensHubCallerRaw struct {
	Contract *ILensHubCaller // Generic read-only contract binding to access the raw methods on
}

// ILensHubTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ILensHubTransactorRaw struct {
	Contract *ILensHubTransactor // Generic write-only contract binding to access the raw methods on
}

// NewILensHub creates a new instance of ILensHub, bound to a specific deployed contract.
func NewILensHub(address common.Address, backend bind.ContractBackend) (*ILensHub, error) {
	contract, err := bindILensHub(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ILensHub{ILensHubCaller: ILensHubCaller{contract: contract}, ILensHubTransactor: ILensHubTransactor{contract: contract}, ILensHubFilterer: ILensHubFilterer{contract: contract}}, nil
}

// NewILensHubCaller creates a new read-only instance of ILensHub, bound to a specific deployed contract.
func NewILensHubCaller(address common.Address, caller bind.ContractCaller) (*ILensHubCaller, error) {
	contract, err := bindILensHub(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ILensHubCaller{contract: contract}, nil
}

// NewILensHubTransactor creates a new write-only instance of ILensHub, bound to a specific deployed contract.
func NewILensHubTransactor(address common.Address, transactor bind.ContractTransactor) (*ILensHubTransactor, error) {
	contract, err := bindILensHub(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ILensHubTransactor{contract: contract}, nil
}

// NewILensHubFilterer creates a new log filterer instance of ILensHub, bound to a specific deployed contract.
func NewILensHubFilterer(address common.Address, filterer bind.ContractFilterer) (*ILensHubFilterer, error) {
	contract, err := bindILensHub(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ILensHubFilterer{contract: contract}, nil
}

// bindILensHub binds a generic wrapper to an already deployed contract.
func bindILensHub(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ILensHubABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ILensHub *ILensHubRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ILensHub.Contract.ILensHubCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ILensHub *ILensHubRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ILensHub.Contract.ILensHubTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ILensHub *ILensHubRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ILensHub.Contract.ILensHubTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ILensHub *ILensHubCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ILensHub.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ILensHub *ILensHubTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ILensHub.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ILensHub *ILensHubTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ILensHub.Contract.contract.Transact(opts, method, params...)
}

// DefaultProfile is a free data retrieval call binding the contract method 0x92254a62.
//
// Solidity: function defaultProfile(address wallet) view returns(uint256)
func (_ILensHub *ILensHubCaller) DefaultProfile(opts *bind.CallOpts, wallet common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ILensHub.contract.Call(opts, &out, "defaultProfile", wallet)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultProfile is a free data retrieval call binding the contract method 0x92254a62.
//
// Solidity: function defaultProfile(address wallet) view returns(uint256)
func (_ILensHub *ILensHubSession) DefaultProfile(wallet common.Address) (*big.Int, error) {
	return _ILensHub.Contract.DefaultProfile(&_ILensHub.CallOpts, wallet)
}

// DefaultProfile is a free data retrieval call binding the contract method 0x92254a62.
//
// Solidity: function defaultProfile(address wallet) view returns(uint256)
func (_ILensHub *ILensHubCallerSession) DefaultProfile(wallet common.Address) (*big.Int, error) {
	return _ILensHub.Contract.DefaultProfile(&_ILensHub.CallOpts, wallet)
}

// GetCollectModule is a free data retrieval call binding the contract method 0x57ff49f6.
//
// Solidity: function getCollectModule(uint256 profileId, uint256 pubId) view returns(address)
func (_ILensHub *ILensHubCaller) GetCollectModule(opts *bind.CallOpts, profileId *big.Int, pubId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ILensHub.contract.Call(opts, &out, "getCollectModule", profileId, pubId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCollectModule is a free data retrieval call binding the contract method 0x57ff49f6.
//
// Solidity: function getCollectModule(uint256 profileId, uint256 pubId) view returns(address)
func (_ILensHub *ILensHubSession) GetCollectModule(profileId *big.Int, pubId *big.Int) (common.Address, error) {
	return _ILensHub.Contract.GetCollectModule(&_ILensHub.CallOpts, profileId, pubId)
}

// GetCollectModule is a free data retrieval call binding the contract method 0x57ff49f6.
//
// Solidity: function getCollectModule(uint256 profileId, uint256 pubId) view returns(address)
func (_ILensHub *ILensHubCallerSession) GetCollectModule(profileId *big.Int, pubId *big.Int) (common.Address, error) {
	return _ILensHub.Contract.GetCollectModule(&_ILensHub.CallOpts, profileId, pubId)
}

// GetCollectNFT is a free data retrieval call binding the contract method 0x52aaef55.
//
// Solidity: function getCollectNFT(uint256 profileId, uint256 pubId) view returns(address)
func (_ILensHub *ILensHubCaller) GetCollectNFT(opts *bind.CallOpts, profileId *big.Int, pubId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ILensHub.contract.Call(opts, &out, "getCollectNFT", profileId, pubId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCollectNFT is a free data retrieval call binding the contract method 0x52aaef55.
//
// Solidity: function getCollectNFT(uint256 profileId, uint256 pubId) view returns(address)
func (_ILensHub *ILensHubSession) GetCollectNFT(profileId *big.Int, pubId *big.Int) (common.Address, error) {
	return _ILensHub.Contract.GetCollectNFT(&_ILensHub.CallOpts, profileId, pubId)
}

// GetCollectNFT is a free data retrieval call binding the contract method 0x52aaef55.
//
// Solidity: function getCollectNFT(uint256 profileId, uint256 pubId) view returns(address)
func (_ILensHub *ILensHubCallerSession) GetCollectNFT(profileId *big.Int, pubId *big.Int) (common.Address, error) {
	return _ILensHub.Contract.GetCollectNFT(&_ILensHub.CallOpts, profileId, pubId)
}

// GetCollectNFTImpl is a free data retrieval call binding the contract method 0x77349a5f.
//
// Solidity: function getCollectNFTImpl() view returns(address)
func (_ILensHub *ILensHubCaller) GetCollectNFTImpl(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ILensHub.contract.Call(opts, &out, "getCollectNFTImpl")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCollectNFTImpl is a free data retrieval call binding the contract method 0x77349a5f.
//
// Solidity: function getCollectNFTImpl() view returns(address)
func (_ILensHub *ILensHubSession) GetCollectNFTImpl() (common.Address, error) {
	return _ILensHub.Contract.GetCollectNFTImpl(&_ILensHub.CallOpts)
}

// GetCollectNFTImpl is a free data retrieval call binding the contract method 0x77349a5f.
//
// Solidity: function getCollectNFTImpl() view returns(address)
func (_ILensHub *ILensHubCallerSession) GetCollectNFTImpl() (common.Address, error) {
	return _ILensHub.Contract.GetCollectNFTImpl(&_ILensHub.CallOpts)
}

// GetContentURI is a free data retrieval call binding the contract method 0xb5a31496.
//
// Solidity: function getContentURI(uint256 profileId, uint256 pubId) view returns(string)
func (_ILensHub *ILensHubCaller) GetContentURI(opts *bind.CallOpts, profileId *big.Int, pubId *big.Int) (string, error) {
	var out []interface{}
	err := _ILensHub.contract.Call(opts, &out, "getContentURI", profileId, pubId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetContentURI is a free data retrieval call binding the contract method 0xb5a31496.
//
// Solidity: function getContentURI(uint256 profileId, uint256 pubId) view returns(string)
func (_ILensHub *ILensHubSession) GetContentURI(profileId *big.Int, pubId *big.Int) (string, error) {
	return _ILensHub.Contract.GetContentURI(&_ILensHub.CallOpts, profileId, pubId)
}

// GetContentURI is a free data retrieval call binding the contract method 0xb5a31496.
//
// Solidity: function getContentURI(uint256 profileId, uint256 pubId) view returns(string)
func (_ILensHub *ILensHubCallerSession) GetContentURI(profileId *big.Int, pubId *big.Int) (string, error) {
	return _ILensHub.Contract.GetContentURI(&_ILensHub.CallOpts, profileId, pubId)
}

// GetDispatcher is a free data retrieval call binding the contract method 0x540528b9.
//
// Solidity: function getDispatcher(uint256 profileId) view returns(address)
func (_ILensHub *ILensHubCaller) GetDispatcher(opts *bind.CallOpts, profileId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ILensHub.contract.Call(opts, &out, "getDispatcher", profileId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetDispatcher is a free data retrieval call binding the contract method 0x540528b9.
//
// Solidity: function getDispatcher(uint256 profileId) view returns(address)
func (_ILensHub *ILensHubSession) GetDispatcher(profileId *big.Int) (common.Address, error) {
	return _ILensHub.Contract.GetDispatcher(&_ILensHub.CallOpts, profileId)
}

// GetDispatcher is a free data retrieval call binding the contract method 0x540528b9.
//
// Solidity: function getDispatcher(uint256 profileId) view returns(address)
func (_ILensHub *ILensHubCallerSession) GetDispatcher(profileId *big.Int) (common.Address, error) {
	return _ILensHub.Contract.GetDispatcher(&_ILensHub.CallOpts, profileId)
}

// GetFollowModule is a free data retrieval call binding the contract method 0xd923d20c.
//
// Solidity: function getFollowModule(uint256 profileId) view returns(address)
func (_ILensHub *ILensHubCaller) GetFollowModule(opts *bind.CallOpts, profileId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ILensHub.contract.Call(opts, &out, "getFollowModule", profileId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFollowModule is a free data retrieval call binding the contract method 0xd923d20c.
//
// Solidity: function getFollowModule(uint256 profileId) view returns(address)
func (_ILensHub *ILensHubSession) GetFollowModule(profileId *big.Int) (common.Address, error) {
	return _ILensHub.Contract.GetFollowModule(&_ILensHub.CallOpts, profileId)
}

// GetFollowModule is a free data retrieval call binding the contract method 0xd923d20c.
//
// Solidity: function getFollowModule(uint256 profileId) view returns(address)
func (_ILensHub *ILensHubCallerSession) GetFollowModule(profileId *big.Int) (common.Address, error) {
	return _ILensHub.Contract.GetFollowModule(&_ILensHub.CallOpts, profileId)
}

// GetFollowNFT is a free data retrieval call binding the contract method 0xa9ec6563.
//
// Solidity: function getFollowNFT(uint256 profileId) view returns(address)
func (_ILensHub *ILensHubCaller) GetFollowNFT(opts *bind.CallOpts, profileId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ILensHub.contract.Call(opts, &out, "getFollowNFT", profileId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFollowNFT is a free data retrieval call binding the contract method 0xa9ec6563.
//
// Solidity: function getFollowNFT(uint256 profileId) view returns(address)
func (_ILensHub *ILensHubSession) GetFollowNFT(profileId *big.Int) (common.Address, error) {
	return _ILensHub.Contract.GetFollowNFT(&_ILensHub.CallOpts, profileId)
}

// GetFollowNFT is a free data retrieval call binding the contract method 0xa9ec6563.
//
// Solidity: function getFollowNFT(uint256 profileId) view returns(address)
func (_ILensHub *ILensHubCallerSession) GetFollowNFT(profileId *big.Int) (common.Address, error) {
	return _ILensHub.Contract.GetFollowNFT(&_ILensHub.CallOpts, profileId)
}

// GetFollowNFTImpl is a free data retrieval call binding the contract method 0x3502ac4b.
//
// Solidity: function getFollowNFTImpl() view returns(address)
func (_ILensHub *ILensHubCaller) GetFollowNFTImpl(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ILensHub.contract.Call(opts, &out, "getFollowNFTImpl")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFollowNFTImpl is a free data retrieval call binding the contract method 0x3502ac4b.
//
// Solidity: function getFollowNFTImpl() view returns(address)
func (_ILensHub *ILensHubSession) GetFollowNFTImpl() (common.Address, error) {
	return _ILensHub.Contract.GetFollowNFTImpl(&_ILensHub.CallOpts)
}

// GetFollowNFTImpl is a free data retrieval call binding the contract method 0x3502ac4b.
//
// Solidity: function getFollowNFTImpl() view returns(address)
func (_ILensHub *ILensHubCallerSession) GetFollowNFTImpl() (common.Address, error) {
	return _ILensHub.Contract.GetFollowNFTImpl(&_ILensHub.CallOpts)
}

// GetFollowNFTURI is a free data retrieval call binding the contract method 0x374c9473.
//
// Solidity: function getFollowNFTURI(uint256 profileId) view returns(string)
func (_ILensHub *ILensHubCaller) GetFollowNFTURI(opts *bind.CallOpts, profileId *big.Int) (string, error) {
	var out []interface{}
	err := _ILensHub.contract.Call(opts, &out, "getFollowNFTURI", profileId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetFollowNFTURI is a free data retrieval call binding the contract method 0x374c9473.
//
// Solidity: function getFollowNFTURI(uint256 profileId) view returns(string)
func (_ILensHub *ILensHubSession) GetFollowNFTURI(profileId *big.Int) (string, error) {
	return _ILensHub.Contract.GetFollowNFTURI(&_ILensHub.CallOpts, profileId)
}

// GetFollowNFTURI is a free data retrieval call binding the contract method 0x374c9473.
//
// Solidity: function getFollowNFTURI(uint256 profileId) view returns(string)
func (_ILensHub *ILensHubCallerSession) GetFollowNFTURI(profileId *big.Int) (string, error) {
	return _ILensHub.Contract.GetFollowNFTURI(&_ILensHub.CallOpts, profileId)
}

// GetGovernance is a free data retrieval call binding the contract method 0x289b3c0d.
//
// Solidity: function getGovernance() view returns(address)
func (_ILensHub *ILensHubCaller) GetGovernance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ILensHub.contract.Call(opts, &out, "getGovernance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetGovernance is a free data retrieval call binding the contract method 0x289b3c0d.
//
// Solidity: function getGovernance() view returns(address)
func (_ILensHub *ILensHubSession) GetGovernance() (common.Address, error) {
	return _ILensHub.Contract.GetGovernance(&_ILensHub.CallOpts)
}

// GetGovernance is a free data retrieval call binding the contract method 0x289b3c0d.
//
// Solidity: function getGovernance() view returns(address)
func (_ILensHub *ILensHubCallerSession) GetGovernance() (common.Address, error) {
	return _ILensHub.Contract.GetGovernance(&_ILensHub.CallOpts)
}

// GetHandle is a free data retrieval call binding the contract method 0xec81d194.
//
// Solidity: function getHandle(uint256 profileId) view returns(string)
func (_ILensHub *ILensHubCaller) GetHandle(opts *bind.CallOpts, profileId *big.Int) (string, error) {
	var out []interface{}
	err := _ILensHub.contract.Call(opts, &out, "getHandle", profileId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetHandle is a free data retrieval call binding the contract method 0xec81d194.
//
// Solidity: function getHandle(uint256 profileId) view returns(string)
func (_ILensHub *ILensHubSession) GetHandle(profileId *big.Int) (string, error) {
	return _ILensHub.Contract.GetHandle(&_ILensHub.CallOpts, profileId)
}

// GetHandle is a free data retrieval call binding the contract method 0xec81d194.
//
// Solidity: function getHandle(uint256 profileId) view returns(string)
func (_ILensHub *ILensHubCallerSession) GetHandle(profileId *big.Int) (string, error) {
	return _ILensHub.Contract.GetHandle(&_ILensHub.CallOpts, profileId)
}

// GetProfile is a free data retrieval call binding the contract method 0xf08f4f64.
//
// Solidity: function getProfile(uint256 profileId) view returns((uint256,address,address,string,string,string))
func (_ILensHub *ILensHubCaller) GetProfile(opts *bind.CallOpts, profileId *big.Int) (DataTypesProfileStruct, error) {
	var out []interface{}
	err := _ILensHub.contract.Call(opts, &out, "getProfile", profileId)

	if err != nil {
		return *new(DataTypesProfileStruct), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesProfileStruct)).(*DataTypesProfileStruct)

	return out0, err

}

// GetProfile is a free data retrieval call binding the contract method 0xf08f4f64.
//
// Solidity: function getProfile(uint256 profileId) view returns((uint256,address,address,string,string,string))
func (_ILensHub *ILensHubSession) GetProfile(profileId *big.Int) (DataTypesProfileStruct, error) {
	return _ILensHub.Contract.GetProfile(&_ILensHub.CallOpts, profileId)
}

// GetProfile is a free data retrieval call binding the contract method 0xf08f4f64.
//
// Solidity: function getProfile(uint256 profileId) view returns((uint256,address,address,string,string,string))
func (_ILensHub *ILensHubCallerSession) GetProfile(profileId *big.Int) (DataTypesProfileStruct, error) {
	return _ILensHub.Contract.GetProfile(&_ILensHub.CallOpts, profileId)
}

// GetProfileIdByHandle is a free data retrieval call binding the contract method 0x20fa728a.
//
// Solidity: function getProfileIdByHandle(string handle) view returns(uint256)
func (_ILensHub *ILensHubCaller) GetProfileIdByHandle(opts *bind.CallOpts, handle string) (*big.Int, error) {
	var out []interface{}
	err := _ILensHub.contract.Call(opts, &out, "getProfileIdByHandle", handle)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProfileIdByHandle is a free data retrieval call binding the contract method 0x20fa728a.
//
// Solidity: function getProfileIdByHandle(string handle) view returns(uint256)
func (_ILensHub *ILensHubSession) GetProfileIdByHandle(handle string) (*big.Int, error) {
	return _ILensHub.Contract.GetProfileIdByHandle(&_ILensHub.CallOpts, handle)
}

// GetProfileIdByHandle is a free data retrieval call binding the contract method 0x20fa728a.
//
// Solidity: function getProfileIdByHandle(string handle) view returns(uint256)
func (_ILensHub *ILensHubCallerSession) GetProfileIdByHandle(handle string) (*big.Int, error) {
	return _ILensHub.Contract.GetProfileIdByHandle(&_ILensHub.CallOpts, handle)
}

// GetPub is a free data retrieval call binding the contract method 0xc736c990.
//
// Solidity: function getPub(uint256 profileId, uint256 pubId) view returns((uint256,uint256,string,address,address,address))
func (_ILensHub *ILensHubCaller) GetPub(opts *bind.CallOpts, profileId *big.Int, pubId *big.Int) (DataTypesPublicationStruct, error) {
	var out []interface{}
	err := _ILensHub.contract.Call(opts, &out, "getPub", profileId, pubId)

	if err != nil {
		return *new(DataTypesPublicationStruct), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesPublicationStruct)).(*DataTypesPublicationStruct)

	return out0, err

}

// GetPub is a free data retrieval call binding the contract method 0xc736c990.
//
// Solidity: function getPub(uint256 profileId, uint256 pubId) view returns((uint256,uint256,string,address,address,address))
func (_ILensHub *ILensHubSession) GetPub(profileId *big.Int, pubId *big.Int) (DataTypesPublicationStruct, error) {
	return _ILensHub.Contract.GetPub(&_ILensHub.CallOpts, profileId, pubId)
}

// GetPub is a free data retrieval call binding the contract method 0xc736c990.
//
// Solidity: function getPub(uint256 profileId, uint256 pubId) view returns((uint256,uint256,string,address,address,address))
func (_ILensHub *ILensHubCallerSession) GetPub(profileId *big.Int, pubId *big.Int) (DataTypesPublicationStruct, error) {
	return _ILensHub.Contract.GetPub(&_ILensHub.CallOpts, profileId, pubId)
}

// GetPubCount is a free data retrieval call binding the contract method 0x3a15ff07.
//
// Solidity: function getPubCount(uint256 profileId) view returns(uint256)
func (_ILensHub *ILensHubCaller) GetPubCount(opts *bind.CallOpts, profileId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ILensHub.contract.Call(opts, &out, "getPubCount", profileId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPubCount is a free data retrieval call binding the contract method 0x3a15ff07.
//
// Solidity: function getPubCount(uint256 profileId) view returns(uint256)
func (_ILensHub *ILensHubSession) GetPubCount(profileId *big.Int) (*big.Int, error) {
	return _ILensHub.Contract.GetPubCount(&_ILensHub.CallOpts, profileId)
}

// GetPubCount is a free data retrieval call binding the contract method 0x3a15ff07.
//
// Solidity: function getPubCount(uint256 profileId) view returns(uint256)
func (_ILensHub *ILensHubCallerSession) GetPubCount(profileId *big.Int) (*big.Int, error) {
	return _ILensHub.Contract.GetPubCount(&_ILensHub.CallOpts, profileId)
}

// GetPubPointer is a free data retrieval call binding the contract method 0x5ca3eebf.
//
// Solidity: function getPubPointer(uint256 profileId, uint256 pubId) view returns(uint256, uint256)
func (_ILensHub *ILensHubCaller) GetPubPointer(opts *bind.CallOpts, profileId *big.Int, pubId *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _ILensHub.contract.Call(opts, &out, "getPubPointer", profileId, pubId)

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
func (_ILensHub *ILensHubSession) GetPubPointer(profileId *big.Int, pubId *big.Int) (*big.Int, *big.Int, error) {
	return _ILensHub.Contract.GetPubPointer(&_ILensHub.CallOpts, profileId, pubId)
}

// GetPubPointer is a free data retrieval call binding the contract method 0x5ca3eebf.
//
// Solidity: function getPubPointer(uint256 profileId, uint256 pubId) view returns(uint256, uint256)
func (_ILensHub *ILensHubCallerSession) GetPubPointer(profileId *big.Int, pubId *big.Int) (*big.Int, *big.Int, error) {
	return _ILensHub.Contract.GetPubPointer(&_ILensHub.CallOpts, profileId, pubId)
}

// GetPubType is a free data retrieval call binding the contract method 0x31fff07c.
//
// Solidity: function getPubType(uint256 profileId, uint256 pubId) view returns(uint8)
func (_ILensHub *ILensHubCaller) GetPubType(opts *bind.CallOpts, profileId *big.Int, pubId *big.Int) (uint8, error) {
	var out []interface{}
	err := _ILensHub.contract.Call(opts, &out, "getPubType", profileId, pubId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetPubType is a free data retrieval call binding the contract method 0x31fff07c.
//
// Solidity: function getPubType(uint256 profileId, uint256 pubId) view returns(uint8)
func (_ILensHub *ILensHubSession) GetPubType(profileId *big.Int, pubId *big.Int) (uint8, error) {
	return _ILensHub.Contract.GetPubType(&_ILensHub.CallOpts, profileId, pubId)
}

// GetPubType is a free data retrieval call binding the contract method 0x31fff07c.
//
// Solidity: function getPubType(uint256 profileId, uint256 pubId) view returns(uint8)
func (_ILensHub *ILensHubCallerSession) GetPubType(profileId *big.Int, pubId *big.Int) (uint8, error) {
	return _ILensHub.Contract.GetPubType(&_ILensHub.CallOpts, profileId, pubId)
}

// GetReferenceModule is a free data retrieval call binding the contract method 0xb7984c05.
//
// Solidity: function getReferenceModule(uint256 profileId, uint256 pubId) view returns(address)
func (_ILensHub *ILensHubCaller) GetReferenceModule(opts *bind.CallOpts, profileId *big.Int, pubId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ILensHub.contract.Call(opts, &out, "getReferenceModule", profileId, pubId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetReferenceModule is a free data retrieval call binding the contract method 0xb7984c05.
//
// Solidity: function getReferenceModule(uint256 profileId, uint256 pubId) view returns(address)
func (_ILensHub *ILensHubSession) GetReferenceModule(profileId *big.Int, pubId *big.Int) (common.Address, error) {
	return _ILensHub.Contract.GetReferenceModule(&_ILensHub.CallOpts, profileId, pubId)
}

// GetReferenceModule is a free data retrieval call binding the contract method 0xb7984c05.
//
// Solidity: function getReferenceModule(uint256 profileId, uint256 pubId) view returns(address)
func (_ILensHub *ILensHubCallerSession) GetReferenceModule(profileId *big.Int, pubId *big.Int) (common.Address, error) {
	return _ILensHub.Contract.GetReferenceModule(&_ILensHub.CallOpts, profileId, pubId)
}

// IsCollectModuleWhitelisted is a free data retrieval call binding the contract method 0xad3e72bf.
//
// Solidity: function isCollectModuleWhitelisted(address collectModule) view returns(bool)
func (_ILensHub *ILensHubCaller) IsCollectModuleWhitelisted(opts *bind.CallOpts, collectModule common.Address) (bool, error) {
	var out []interface{}
	err := _ILensHub.contract.Call(opts, &out, "isCollectModuleWhitelisted", collectModule)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCollectModuleWhitelisted is a free data retrieval call binding the contract method 0xad3e72bf.
//
// Solidity: function isCollectModuleWhitelisted(address collectModule) view returns(bool)
func (_ILensHub *ILensHubSession) IsCollectModuleWhitelisted(collectModule common.Address) (bool, error) {
	return _ILensHub.Contract.IsCollectModuleWhitelisted(&_ILensHub.CallOpts, collectModule)
}

// IsCollectModuleWhitelisted is a free data retrieval call binding the contract method 0xad3e72bf.
//
// Solidity: function isCollectModuleWhitelisted(address collectModule) view returns(bool)
func (_ILensHub *ILensHubCallerSession) IsCollectModuleWhitelisted(collectModule common.Address) (bool, error) {
	return _ILensHub.Contract.IsCollectModuleWhitelisted(&_ILensHub.CallOpts, collectModule)
}

// IsFollowModuleWhitelisted is a free data retrieval call binding the contract method 0x1cbb2620.
//
// Solidity: function isFollowModuleWhitelisted(address followModule) view returns(bool)
func (_ILensHub *ILensHubCaller) IsFollowModuleWhitelisted(opts *bind.CallOpts, followModule common.Address) (bool, error) {
	var out []interface{}
	err := _ILensHub.contract.Call(opts, &out, "isFollowModuleWhitelisted", followModule)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFollowModuleWhitelisted is a free data retrieval call binding the contract method 0x1cbb2620.
//
// Solidity: function isFollowModuleWhitelisted(address followModule) view returns(bool)
func (_ILensHub *ILensHubSession) IsFollowModuleWhitelisted(followModule common.Address) (bool, error) {
	return _ILensHub.Contract.IsFollowModuleWhitelisted(&_ILensHub.CallOpts, followModule)
}

// IsFollowModuleWhitelisted is a free data retrieval call binding the contract method 0x1cbb2620.
//
// Solidity: function isFollowModuleWhitelisted(address followModule) view returns(bool)
func (_ILensHub *ILensHubCallerSession) IsFollowModuleWhitelisted(followModule common.Address) (bool, error) {
	return _ILensHub.Contract.IsFollowModuleWhitelisted(&_ILensHub.CallOpts, followModule)
}

// IsProfileCreatorWhitelisted is a free data retrieval call binding the contract method 0xaf05dd22.
//
// Solidity: function isProfileCreatorWhitelisted(address profileCreator) view returns(bool)
func (_ILensHub *ILensHubCaller) IsProfileCreatorWhitelisted(opts *bind.CallOpts, profileCreator common.Address) (bool, error) {
	var out []interface{}
	err := _ILensHub.contract.Call(opts, &out, "isProfileCreatorWhitelisted", profileCreator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsProfileCreatorWhitelisted is a free data retrieval call binding the contract method 0xaf05dd22.
//
// Solidity: function isProfileCreatorWhitelisted(address profileCreator) view returns(bool)
func (_ILensHub *ILensHubSession) IsProfileCreatorWhitelisted(profileCreator common.Address) (bool, error) {
	return _ILensHub.Contract.IsProfileCreatorWhitelisted(&_ILensHub.CallOpts, profileCreator)
}

// IsProfileCreatorWhitelisted is a free data retrieval call binding the contract method 0xaf05dd22.
//
// Solidity: function isProfileCreatorWhitelisted(address profileCreator) view returns(bool)
func (_ILensHub *ILensHubCallerSession) IsProfileCreatorWhitelisted(profileCreator common.Address) (bool, error) {
	return _ILensHub.Contract.IsProfileCreatorWhitelisted(&_ILensHub.CallOpts, profileCreator)
}

// IsReferenceModuleWhitelisted is a free data retrieval call binding the contract method 0x8e204fb4.
//
// Solidity: function isReferenceModuleWhitelisted(address referenceModule) view returns(bool)
func (_ILensHub *ILensHubCaller) IsReferenceModuleWhitelisted(opts *bind.CallOpts, referenceModule common.Address) (bool, error) {
	var out []interface{}
	err := _ILensHub.contract.Call(opts, &out, "isReferenceModuleWhitelisted", referenceModule)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsReferenceModuleWhitelisted is a free data retrieval call binding the contract method 0x8e204fb4.
//
// Solidity: function isReferenceModuleWhitelisted(address referenceModule) view returns(bool)
func (_ILensHub *ILensHubSession) IsReferenceModuleWhitelisted(referenceModule common.Address) (bool, error) {
	return _ILensHub.Contract.IsReferenceModuleWhitelisted(&_ILensHub.CallOpts, referenceModule)
}

// IsReferenceModuleWhitelisted is a free data retrieval call binding the contract method 0x8e204fb4.
//
// Solidity: function isReferenceModuleWhitelisted(address referenceModule) view returns(bool)
func (_ILensHub *ILensHubCallerSession) IsReferenceModuleWhitelisted(referenceModule common.Address) (bool, error) {
	return _ILensHub.Contract.IsReferenceModuleWhitelisted(&_ILensHub.CallOpts, referenceModule)
}

// Collect is a paid mutator transaction binding the contract method 0x84114ad4.
//
// Solidity: function collect(uint256 profileId, uint256 pubId, bytes data) returns(uint256)
func (_ILensHub *ILensHubTransactor) Collect(opts *bind.TransactOpts, profileId *big.Int, pubId *big.Int, data []byte) (*types.Transaction, error) {
	return _ILensHub.contract.Transact(opts, "collect", profileId, pubId, data)
}

// Collect is a paid mutator transaction binding the contract method 0x84114ad4.
//
// Solidity: function collect(uint256 profileId, uint256 pubId, bytes data) returns(uint256)
func (_ILensHub *ILensHubSession) Collect(profileId *big.Int, pubId *big.Int, data []byte) (*types.Transaction, error) {
	return _ILensHub.Contract.Collect(&_ILensHub.TransactOpts, profileId, pubId, data)
}

// Collect is a paid mutator transaction binding the contract method 0x84114ad4.
//
// Solidity: function collect(uint256 profileId, uint256 pubId, bytes data) returns(uint256)
func (_ILensHub *ILensHubTransactorSession) Collect(profileId *big.Int, pubId *big.Int, data []byte) (*types.Transaction, error) {
	return _ILensHub.Contract.Collect(&_ILensHub.TransactOpts, profileId, pubId, data)
}

// CollectWithSig is a paid mutator transaction binding the contract method 0xb48951e4.
//
// Solidity: function collectWithSig((address,uint256,uint256,bytes,(uint8,bytes32,bytes32,uint256)) vars) returns(uint256)
func (_ILensHub *ILensHubTransactor) CollectWithSig(opts *bind.TransactOpts, vars DataTypesCollectWithSigData) (*types.Transaction, error) {
	return _ILensHub.contract.Transact(opts, "collectWithSig", vars)
}

// CollectWithSig is a paid mutator transaction binding the contract method 0xb48951e4.
//
// Solidity: function collectWithSig((address,uint256,uint256,bytes,(uint8,bytes32,bytes32,uint256)) vars) returns(uint256)
func (_ILensHub *ILensHubSession) CollectWithSig(vars DataTypesCollectWithSigData) (*types.Transaction, error) {
	return _ILensHub.Contract.CollectWithSig(&_ILensHub.TransactOpts, vars)
}

// CollectWithSig is a paid mutator transaction binding the contract method 0xb48951e4.
//
// Solidity: function collectWithSig((address,uint256,uint256,bytes,(uint8,bytes32,bytes32,uint256)) vars) returns(uint256)
func (_ILensHub *ILensHubTransactorSession) CollectWithSig(vars DataTypesCollectWithSigData) (*types.Transaction, error) {
	return _ILensHub.Contract.CollectWithSig(&_ILensHub.TransactOpts, vars)
}

// Comment is a paid mutator transaction binding the contract method 0xb6f32d2b.
//
// Solidity: function comment((uint256,string,uint256,uint256,bytes,address,bytes,address,bytes) vars) returns(uint256)
func (_ILensHub *ILensHubTransactor) Comment(opts *bind.TransactOpts, vars DataTypesCommentData) (*types.Transaction, error) {
	return _ILensHub.contract.Transact(opts, "comment", vars)
}

// Comment is a paid mutator transaction binding the contract method 0xb6f32d2b.
//
// Solidity: function comment((uint256,string,uint256,uint256,bytes,address,bytes,address,bytes) vars) returns(uint256)
func (_ILensHub *ILensHubSession) Comment(vars DataTypesCommentData) (*types.Transaction, error) {
	return _ILensHub.Contract.Comment(&_ILensHub.TransactOpts, vars)
}

// Comment is a paid mutator transaction binding the contract method 0xb6f32d2b.
//
// Solidity: function comment((uint256,string,uint256,uint256,bytes,address,bytes,address,bytes) vars) returns(uint256)
func (_ILensHub *ILensHubTransactorSession) Comment(vars DataTypesCommentData) (*types.Transaction, error) {
	return _ILensHub.Contract.Comment(&_ILensHub.TransactOpts, vars)
}

// CommentWithSig is a paid mutator transaction binding the contract method 0x7a375716.
//
// Solidity: function commentWithSig((uint256,string,uint256,uint256,bytes,address,bytes,address,bytes,(uint8,bytes32,bytes32,uint256)) vars) returns(uint256)
func (_ILensHub *ILensHubTransactor) CommentWithSig(opts *bind.TransactOpts, vars DataTypesCommentWithSigData) (*types.Transaction, error) {
	return _ILensHub.contract.Transact(opts, "commentWithSig", vars)
}

// CommentWithSig is a paid mutator transaction binding the contract method 0x7a375716.
//
// Solidity: function commentWithSig((uint256,string,uint256,uint256,bytes,address,bytes,address,bytes,(uint8,bytes32,bytes32,uint256)) vars) returns(uint256)
func (_ILensHub *ILensHubSession) CommentWithSig(vars DataTypesCommentWithSigData) (*types.Transaction, error) {
	return _ILensHub.Contract.CommentWithSig(&_ILensHub.TransactOpts, vars)
}

// CommentWithSig is a paid mutator transaction binding the contract method 0x7a375716.
//
// Solidity: function commentWithSig((uint256,string,uint256,uint256,bytes,address,bytes,address,bytes,(uint8,bytes32,bytes32,uint256)) vars) returns(uint256)
func (_ILensHub *ILensHubTransactorSession) CommentWithSig(vars DataTypesCommentWithSigData) (*types.Transaction, error) {
	return _ILensHub.Contract.CommentWithSig(&_ILensHub.TransactOpts, vars)
}

// CreateProfile is a paid mutator transaction binding the contract method 0xffea138e.
//
// Solidity: function createProfile((address,string,string,address,bytes,string) vars) returns(uint256)
func (_ILensHub *ILensHubTransactor) CreateProfile(opts *bind.TransactOpts, vars DataTypesCreateProfileData) (*types.Transaction, error) {
	return _ILensHub.contract.Transact(opts, "createProfile", vars)
}

// CreateProfile is a paid mutator transaction binding the contract method 0xffea138e.
//
// Solidity: function createProfile((address,string,string,address,bytes,string) vars) returns(uint256)
func (_ILensHub *ILensHubSession) CreateProfile(vars DataTypesCreateProfileData) (*types.Transaction, error) {
	return _ILensHub.Contract.CreateProfile(&_ILensHub.TransactOpts, vars)
}

// CreateProfile is a paid mutator transaction binding the contract method 0xffea138e.
//
// Solidity: function createProfile((address,string,string,address,bytes,string) vars) returns(uint256)
func (_ILensHub *ILensHubTransactorSession) CreateProfile(vars DataTypesCreateProfileData) (*types.Transaction, error) {
	return _ILensHub.Contract.CreateProfile(&_ILensHub.TransactOpts, vars)
}

// EmitCollectNFTTransferEvent is a paid mutator transaction binding the contract method 0x86e2947b.
//
// Solidity: function emitCollectNFTTransferEvent(uint256 profileId, uint256 pubId, uint256 collectNFTId, address from, address to) returns()
func (_ILensHub *ILensHubTransactor) EmitCollectNFTTransferEvent(opts *bind.TransactOpts, profileId *big.Int, pubId *big.Int, collectNFTId *big.Int, from common.Address, to common.Address) (*types.Transaction, error) {
	return _ILensHub.contract.Transact(opts, "emitCollectNFTTransferEvent", profileId, pubId, collectNFTId, from, to)
}

// EmitCollectNFTTransferEvent is a paid mutator transaction binding the contract method 0x86e2947b.
//
// Solidity: function emitCollectNFTTransferEvent(uint256 profileId, uint256 pubId, uint256 collectNFTId, address from, address to) returns()
func (_ILensHub *ILensHubSession) EmitCollectNFTTransferEvent(profileId *big.Int, pubId *big.Int, collectNFTId *big.Int, from common.Address, to common.Address) (*types.Transaction, error) {
	return _ILensHub.Contract.EmitCollectNFTTransferEvent(&_ILensHub.TransactOpts, profileId, pubId, collectNFTId, from, to)
}

// EmitCollectNFTTransferEvent is a paid mutator transaction binding the contract method 0x86e2947b.
//
// Solidity: function emitCollectNFTTransferEvent(uint256 profileId, uint256 pubId, uint256 collectNFTId, address from, address to) returns()
func (_ILensHub *ILensHubTransactorSession) EmitCollectNFTTransferEvent(profileId *big.Int, pubId *big.Int, collectNFTId *big.Int, from common.Address, to common.Address) (*types.Transaction, error) {
	return _ILensHub.Contract.EmitCollectNFTTransferEvent(&_ILensHub.TransactOpts, profileId, pubId, collectNFTId, from, to)
}

// EmitFollowNFTTransferEvent is a paid mutator transaction binding the contract method 0xbdfdd4bc.
//
// Solidity: function emitFollowNFTTransferEvent(uint256 profileId, uint256 followNFTId, address from, address to) returns()
func (_ILensHub *ILensHubTransactor) EmitFollowNFTTransferEvent(opts *bind.TransactOpts, profileId *big.Int, followNFTId *big.Int, from common.Address, to common.Address) (*types.Transaction, error) {
	return _ILensHub.contract.Transact(opts, "emitFollowNFTTransferEvent", profileId, followNFTId, from, to)
}

// EmitFollowNFTTransferEvent is a paid mutator transaction binding the contract method 0xbdfdd4bc.
//
// Solidity: function emitFollowNFTTransferEvent(uint256 profileId, uint256 followNFTId, address from, address to) returns()
func (_ILensHub *ILensHubSession) EmitFollowNFTTransferEvent(profileId *big.Int, followNFTId *big.Int, from common.Address, to common.Address) (*types.Transaction, error) {
	return _ILensHub.Contract.EmitFollowNFTTransferEvent(&_ILensHub.TransactOpts, profileId, followNFTId, from, to)
}

// EmitFollowNFTTransferEvent is a paid mutator transaction binding the contract method 0xbdfdd4bc.
//
// Solidity: function emitFollowNFTTransferEvent(uint256 profileId, uint256 followNFTId, address from, address to) returns()
func (_ILensHub *ILensHubTransactorSession) EmitFollowNFTTransferEvent(profileId *big.Int, followNFTId *big.Int, from common.Address, to common.Address) (*types.Transaction, error) {
	return _ILensHub.Contract.EmitFollowNFTTransferEvent(&_ILensHub.TransactOpts, profileId, followNFTId, from, to)
}

// Follow is a paid mutator transaction binding the contract method 0xfb78ae6c.
//
// Solidity: function follow(uint256[] profileIds, bytes[] datas) returns(uint256[])
func (_ILensHub *ILensHubTransactor) Follow(opts *bind.TransactOpts, profileIds []*big.Int, datas [][]byte) (*types.Transaction, error) {
	return _ILensHub.contract.Transact(opts, "follow", profileIds, datas)
}

// Follow is a paid mutator transaction binding the contract method 0xfb78ae6c.
//
// Solidity: function follow(uint256[] profileIds, bytes[] datas) returns(uint256[])
func (_ILensHub *ILensHubSession) Follow(profileIds []*big.Int, datas [][]byte) (*types.Transaction, error) {
	return _ILensHub.Contract.Follow(&_ILensHub.TransactOpts, profileIds, datas)
}

// Follow is a paid mutator transaction binding the contract method 0xfb78ae6c.
//
// Solidity: function follow(uint256[] profileIds, bytes[] datas) returns(uint256[])
func (_ILensHub *ILensHubTransactorSession) Follow(profileIds []*big.Int, datas [][]byte) (*types.Transaction, error) {
	return _ILensHub.Contract.Follow(&_ILensHub.TransactOpts, profileIds, datas)
}

// FollowWithSig is a paid mutator transaction binding the contract method 0x8e4fd6a9.
//
// Solidity: function followWithSig((address,uint256[],bytes[],(uint8,bytes32,bytes32,uint256)) vars) returns(uint256[])
func (_ILensHub *ILensHubTransactor) FollowWithSig(opts *bind.TransactOpts, vars DataTypesFollowWithSigData) (*types.Transaction, error) {
	return _ILensHub.contract.Transact(opts, "followWithSig", vars)
}

// FollowWithSig is a paid mutator transaction binding the contract method 0x8e4fd6a9.
//
// Solidity: function followWithSig((address,uint256[],bytes[],(uint8,bytes32,bytes32,uint256)) vars) returns(uint256[])
func (_ILensHub *ILensHubSession) FollowWithSig(vars DataTypesFollowWithSigData) (*types.Transaction, error) {
	return _ILensHub.Contract.FollowWithSig(&_ILensHub.TransactOpts, vars)
}

// FollowWithSig is a paid mutator transaction binding the contract method 0x8e4fd6a9.
//
// Solidity: function followWithSig((address,uint256[],bytes[],(uint8,bytes32,bytes32,uint256)) vars) returns(uint256[])
func (_ILensHub *ILensHubTransactorSession) FollowWithSig(vars DataTypesFollowWithSigData) (*types.Transaction, error) {
	return _ILensHub.Contract.FollowWithSig(&_ILensHub.TransactOpts, vars)
}

// Initialize is a paid mutator transaction binding the contract method 0x077f224a.
//
// Solidity: function initialize(string name, string symbol, address newGovernance) returns()
func (_ILensHub *ILensHubTransactor) Initialize(opts *bind.TransactOpts, name string, symbol string, newGovernance common.Address) (*types.Transaction, error) {
	return _ILensHub.contract.Transact(opts, "initialize", name, symbol, newGovernance)
}

// Initialize is a paid mutator transaction binding the contract method 0x077f224a.
//
// Solidity: function initialize(string name, string symbol, address newGovernance) returns()
func (_ILensHub *ILensHubSession) Initialize(name string, symbol string, newGovernance common.Address) (*types.Transaction, error) {
	return _ILensHub.Contract.Initialize(&_ILensHub.TransactOpts, name, symbol, newGovernance)
}

// Initialize is a paid mutator transaction binding the contract method 0x077f224a.
//
// Solidity: function initialize(string name, string symbol, address newGovernance) returns()
func (_ILensHub *ILensHubTransactorSession) Initialize(name string, symbol string, newGovernance common.Address) (*types.Transaction, error) {
	return _ILensHub.Contract.Initialize(&_ILensHub.TransactOpts, name, symbol, newGovernance)
}

// Mirror is a paid mutator transaction binding the contract method 0x2faeed81.
//
// Solidity: function mirror((uint256,uint256,uint256,bytes,address,bytes) vars) returns(uint256)
func (_ILensHub *ILensHubTransactor) Mirror(opts *bind.TransactOpts, vars DataTypesMirrorData) (*types.Transaction, error) {
	return _ILensHub.contract.Transact(opts, "mirror", vars)
}

// Mirror is a paid mutator transaction binding the contract method 0x2faeed81.
//
// Solidity: function mirror((uint256,uint256,uint256,bytes,address,bytes) vars) returns(uint256)
func (_ILensHub *ILensHubSession) Mirror(vars DataTypesMirrorData) (*types.Transaction, error) {
	return _ILensHub.Contract.Mirror(&_ILensHub.TransactOpts, vars)
}

// Mirror is a paid mutator transaction binding the contract method 0x2faeed81.
//
// Solidity: function mirror((uint256,uint256,uint256,bytes,address,bytes) vars) returns(uint256)
func (_ILensHub *ILensHubTransactorSession) Mirror(vars DataTypesMirrorData) (*types.Transaction, error) {
	return _ILensHub.Contract.Mirror(&_ILensHub.TransactOpts, vars)
}

// MirrorWithSig is a paid mutator transaction binding the contract method 0xdf457c34.
//
// Solidity: function mirrorWithSig((uint256,uint256,uint256,bytes,address,bytes,(uint8,bytes32,bytes32,uint256)) vars) returns(uint256)
func (_ILensHub *ILensHubTransactor) MirrorWithSig(opts *bind.TransactOpts, vars DataTypesMirrorWithSigData) (*types.Transaction, error) {
	return _ILensHub.contract.Transact(opts, "mirrorWithSig", vars)
}

// MirrorWithSig is a paid mutator transaction binding the contract method 0xdf457c34.
//
// Solidity: function mirrorWithSig((uint256,uint256,uint256,bytes,address,bytes,(uint8,bytes32,bytes32,uint256)) vars) returns(uint256)
func (_ILensHub *ILensHubSession) MirrorWithSig(vars DataTypesMirrorWithSigData) (*types.Transaction, error) {
	return _ILensHub.Contract.MirrorWithSig(&_ILensHub.TransactOpts, vars)
}

// MirrorWithSig is a paid mutator transaction binding the contract method 0xdf457c34.
//
// Solidity: function mirrorWithSig((uint256,uint256,uint256,bytes,address,bytes,(uint8,bytes32,bytes32,uint256)) vars) returns(uint256)
func (_ILensHub *ILensHubTransactorSession) MirrorWithSig(vars DataTypesMirrorWithSigData) (*types.Transaction, error) {
	return _ILensHub.Contract.MirrorWithSig(&_ILensHub.TransactOpts, vars)
}

// Post is a paid mutator transaction binding the contract method 0x963ff141.
//
// Solidity: function post((uint256,string,address,bytes,address,bytes) vars) returns(uint256)
func (_ILensHub *ILensHubTransactor) Post(opts *bind.TransactOpts, vars DataTypesPostData) (*types.Transaction, error) {
	return _ILensHub.contract.Transact(opts, "post", vars)
}

// Post is a paid mutator transaction binding the contract method 0x963ff141.
//
// Solidity: function post((uint256,string,address,bytes,address,bytes) vars) returns(uint256)
func (_ILensHub *ILensHubSession) Post(vars DataTypesPostData) (*types.Transaction, error) {
	return _ILensHub.Contract.Post(&_ILensHub.TransactOpts, vars)
}

// Post is a paid mutator transaction binding the contract method 0x963ff141.
//
// Solidity: function post((uint256,string,address,bytes,address,bytes) vars) returns(uint256)
func (_ILensHub *ILensHubTransactorSession) Post(vars DataTypesPostData) (*types.Transaction, error) {
	return _ILensHub.Contract.Post(&_ILensHub.TransactOpts, vars)
}

// PostWithSig is a paid mutator transaction binding the contract method 0x3b508132.
//
// Solidity: function postWithSig((uint256,string,address,bytes,address,bytes,(uint8,bytes32,bytes32,uint256)) vars) returns(uint256)
func (_ILensHub *ILensHubTransactor) PostWithSig(opts *bind.TransactOpts, vars DataTypesPostWithSigData) (*types.Transaction, error) {
	return _ILensHub.contract.Transact(opts, "postWithSig", vars)
}

// PostWithSig is a paid mutator transaction binding the contract method 0x3b508132.
//
// Solidity: function postWithSig((uint256,string,address,bytes,address,bytes,(uint8,bytes32,bytes32,uint256)) vars) returns(uint256)
func (_ILensHub *ILensHubSession) PostWithSig(vars DataTypesPostWithSigData) (*types.Transaction, error) {
	return _ILensHub.Contract.PostWithSig(&_ILensHub.TransactOpts, vars)
}

// PostWithSig is a paid mutator transaction binding the contract method 0x3b508132.
//
// Solidity: function postWithSig((uint256,string,address,bytes,address,bytes,(uint8,bytes32,bytes32,uint256)) vars) returns(uint256)
func (_ILensHub *ILensHubTransactorSession) PostWithSig(vars DataTypesPostWithSigData) (*types.Transaction, error) {
	return _ILensHub.Contract.PostWithSig(&_ILensHub.TransactOpts, vars)
}

// SetDefaultProfile is a paid mutator transaction binding the contract method 0xf1b2f8bc.
//
// Solidity: function setDefaultProfile(uint256 profileId) returns()
func (_ILensHub *ILensHubTransactor) SetDefaultProfile(opts *bind.TransactOpts, profileId *big.Int) (*types.Transaction, error) {
	return _ILensHub.contract.Transact(opts, "setDefaultProfile", profileId)
}

// SetDefaultProfile is a paid mutator transaction binding the contract method 0xf1b2f8bc.
//
// Solidity: function setDefaultProfile(uint256 profileId) returns()
func (_ILensHub *ILensHubSession) SetDefaultProfile(profileId *big.Int) (*types.Transaction, error) {
	return _ILensHub.Contract.SetDefaultProfile(&_ILensHub.TransactOpts, profileId)
}

// SetDefaultProfile is a paid mutator transaction binding the contract method 0xf1b2f8bc.
//
// Solidity: function setDefaultProfile(uint256 profileId) returns()
func (_ILensHub *ILensHubTransactorSession) SetDefaultProfile(profileId *big.Int) (*types.Transaction, error) {
	return _ILensHub.Contract.SetDefaultProfile(&_ILensHub.TransactOpts, profileId)
}

// SetDefaultProfileWithSig is a paid mutator transaction binding the contract method 0xdc217253.
//
// Solidity: function setDefaultProfileWithSig((address,uint256,(uint8,bytes32,bytes32,uint256)) vars) returns()
func (_ILensHub *ILensHubTransactor) SetDefaultProfileWithSig(opts *bind.TransactOpts, vars DataTypesSetDefaultProfileWithSigData) (*types.Transaction, error) {
	return _ILensHub.contract.Transact(opts, "setDefaultProfileWithSig", vars)
}

// SetDefaultProfileWithSig is a paid mutator transaction binding the contract method 0xdc217253.
//
// Solidity: function setDefaultProfileWithSig((address,uint256,(uint8,bytes32,bytes32,uint256)) vars) returns()
func (_ILensHub *ILensHubSession) SetDefaultProfileWithSig(vars DataTypesSetDefaultProfileWithSigData) (*types.Transaction, error) {
	return _ILensHub.Contract.SetDefaultProfileWithSig(&_ILensHub.TransactOpts, vars)
}

// SetDefaultProfileWithSig is a paid mutator transaction binding the contract method 0xdc217253.
//
// Solidity: function setDefaultProfileWithSig((address,uint256,(uint8,bytes32,bytes32,uint256)) vars) returns()
func (_ILensHub *ILensHubTransactorSession) SetDefaultProfileWithSig(vars DataTypesSetDefaultProfileWithSigData) (*types.Transaction, error) {
	return _ILensHub.Contract.SetDefaultProfileWithSig(&_ILensHub.TransactOpts, vars)
}

// SetDispatcher is a paid mutator transaction binding the contract method 0xbfd24f47.
//
// Solidity: function setDispatcher(uint256 profileId, address dispatcher) returns()
func (_ILensHub *ILensHubTransactor) SetDispatcher(opts *bind.TransactOpts, profileId *big.Int, dispatcher common.Address) (*types.Transaction, error) {
	return _ILensHub.contract.Transact(opts, "setDispatcher", profileId, dispatcher)
}

// SetDispatcher is a paid mutator transaction binding the contract method 0xbfd24f47.
//
// Solidity: function setDispatcher(uint256 profileId, address dispatcher) returns()
func (_ILensHub *ILensHubSession) SetDispatcher(profileId *big.Int, dispatcher common.Address) (*types.Transaction, error) {
	return _ILensHub.Contract.SetDispatcher(&_ILensHub.TransactOpts, profileId, dispatcher)
}

// SetDispatcher is a paid mutator transaction binding the contract method 0xbfd24f47.
//
// Solidity: function setDispatcher(uint256 profileId, address dispatcher) returns()
func (_ILensHub *ILensHubTransactorSession) SetDispatcher(profileId *big.Int, dispatcher common.Address) (*types.Transaction, error) {
	return _ILensHub.Contract.SetDispatcher(&_ILensHub.TransactOpts, profileId, dispatcher)
}

// SetDispatcherWithSig is a paid mutator transaction binding the contract method 0xbfbf0b4b.
//
// Solidity: function setDispatcherWithSig((uint256,address,(uint8,bytes32,bytes32,uint256)) vars) returns()
func (_ILensHub *ILensHubTransactor) SetDispatcherWithSig(opts *bind.TransactOpts, vars DataTypesSetDispatcherWithSigData) (*types.Transaction, error) {
	return _ILensHub.contract.Transact(opts, "setDispatcherWithSig", vars)
}

// SetDispatcherWithSig is a paid mutator transaction binding the contract method 0xbfbf0b4b.
//
// Solidity: function setDispatcherWithSig((uint256,address,(uint8,bytes32,bytes32,uint256)) vars) returns()
func (_ILensHub *ILensHubSession) SetDispatcherWithSig(vars DataTypesSetDispatcherWithSigData) (*types.Transaction, error) {
	return _ILensHub.Contract.SetDispatcherWithSig(&_ILensHub.TransactOpts, vars)
}

// SetDispatcherWithSig is a paid mutator transaction binding the contract method 0xbfbf0b4b.
//
// Solidity: function setDispatcherWithSig((uint256,address,(uint8,bytes32,bytes32,uint256)) vars) returns()
func (_ILensHub *ILensHubTransactorSession) SetDispatcherWithSig(vars DataTypesSetDispatcherWithSigData) (*types.Transaction, error) {
	return _ILensHub.Contract.SetDispatcherWithSig(&_ILensHub.TransactOpts, vars)
}

// SetEmergencyAdmin is a paid mutator transaction binding the contract method 0x35da3394.
//
// Solidity: function setEmergencyAdmin(address newEmergencyAdmin) returns()
func (_ILensHub *ILensHubTransactor) SetEmergencyAdmin(opts *bind.TransactOpts, newEmergencyAdmin common.Address) (*types.Transaction, error) {
	return _ILensHub.contract.Transact(opts, "setEmergencyAdmin", newEmergencyAdmin)
}

// SetEmergencyAdmin is a paid mutator transaction binding the contract method 0x35da3394.
//
// Solidity: function setEmergencyAdmin(address newEmergencyAdmin) returns()
func (_ILensHub *ILensHubSession) SetEmergencyAdmin(newEmergencyAdmin common.Address) (*types.Transaction, error) {
	return _ILensHub.Contract.SetEmergencyAdmin(&_ILensHub.TransactOpts, newEmergencyAdmin)
}

// SetEmergencyAdmin is a paid mutator transaction binding the contract method 0x35da3394.
//
// Solidity: function setEmergencyAdmin(address newEmergencyAdmin) returns()
func (_ILensHub *ILensHubTransactorSession) SetEmergencyAdmin(newEmergencyAdmin common.Address) (*types.Transaction, error) {
	return _ILensHub.Contract.SetEmergencyAdmin(&_ILensHub.TransactOpts, newEmergencyAdmin)
}

// SetFollowModule is a paid mutator transaction binding the contract method 0x6dea40b3.
//
// Solidity: function setFollowModule(uint256 profileId, address followModule, bytes followModuleInitData) returns()
func (_ILensHub *ILensHubTransactor) SetFollowModule(opts *bind.TransactOpts, profileId *big.Int, followModule common.Address, followModuleInitData []byte) (*types.Transaction, error) {
	return _ILensHub.contract.Transact(opts, "setFollowModule", profileId, followModule, followModuleInitData)
}

// SetFollowModule is a paid mutator transaction binding the contract method 0x6dea40b3.
//
// Solidity: function setFollowModule(uint256 profileId, address followModule, bytes followModuleInitData) returns()
func (_ILensHub *ILensHubSession) SetFollowModule(profileId *big.Int, followModule common.Address, followModuleInitData []byte) (*types.Transaction, error) {
	return _ILensHub.Contract.SetFollowModule(&_ILensHub.TransactOpts, profileId, followModule, followModuleInitData)
}

// SetFollowModule is a paid mutator transaction binding the contract method 0x6dea40b3.
//
// Solidity: function setFollowModule(uint256 profileId, address followModule, bytes followModuleInitData) returns()
func (_ILensHub *ILensHubTransactorSession) SetFollowModule(profileId *big.Int, followModule common.Address, followModuleInitData []byte) (*types.Transaction, error) {
	return _ILensHub.Contract.SetFollowModule(&_ILensHub.TransactOpts, profileId, followModule, followModuleInitData)
}

// SetFollowModuleWithSig is a paid mutator transaction binding the contract method 0x3b28b89f.
//
// Solidity: function setFollowModuleWithSig((uint256,address,bytes,(uint8,bytes32,bytes32,uint256)) vars) returns()
func (_ILensHub *ILensHubTransactor) SetFollowModuleWithSig(opts *bind.TransactOpts, vars DataTypesSetFollowModuleWithSigData) (*types.Transaction, error) {
	return _ILensHub.contract.Transact(opts, "setFollowModuleWithSig", vars)
}

// SetFollowModuleWithSig is a paid mutator transaction binding the contract method 0x3b28b89f.
//
// Solidity: function setFollowModuleWithSig((uint256,address,bytes,(uint8,bytes32,bytes32,uint256)) vars) returns()
func (_ILensHub *ILensHubSession) SetFollowModuleWithSig(vars DataTypesSetFollowModuleWithSigData) (*types.Transaction, error) {
	return _ILensHub.Contract.SetFollowModuleWithSig(&_ILensHub.TransactOpts, vars)
}

// SetFollowModuleWithSig is a paid mutator transaction binding the contract method 0x3b28b89f.
//
// Solidity: function setFollowModuleWithSig((uint256,address,bytes,(uint8,bytes32,bytes32,uint256)) vars) returns()
func (_ILensHub *ILensHubTransactorSession) SetFollowModuleWithSig(vars DataTypesSetFollowModuleWithSigData) (*types.Transaction, error) {
	return _ILensHub.Contract.SetFollowModuleWithSig(&_ILensHub.TransactOpts, vars)
}

// SetFollowNFTURI is a paid mutator transaction binding the contract method 0xc6b5d06c.
//
// Solidity: function setFollowNFTURI(uint256 profileId, string followNFTURI) returns()
func (_ILensHub *ILensHubTransactor) SetFollowNFTURI(opts *bind.TransactOpts, profileId *big.Int, followNFTURI string) (*types.Transaction, error) {
	return _ILensHub.contract.Transact(opts, "setFollowNFTURI", profileId, followNFTURI)
}

// SetFollowNFTURI is a paid mutator transaction binding the contract method 0xc6b5d06c.
//
// Solidity: function setFollowNFTURI(uint256 profileId, string followNFTURI) returns()
func (_ILensHub *ILensHubSession) SetFollowNFTURI(profileId *big.Int, followNFTURI string) (*types.Transaction, error) {
	return _ILensHub.Contract.SetFollowNFTURI(&_ILensHub.TransactOpts, profileId, followNFTURI)
}

// SetFollowNFTURI is a paid mutator transaction binding the contract method 0xc6b5d06c.
//
// Solidity: function setFollowNFTURI(uint256 profileId, string followNFTURI) returns()
func (_ILensHub *ILensHubTransactorSession) SetFollowNFTURI(profileId *big.Int, followNFTURI string) (*types.Transaction, error) {
	return _ILensHub.Contract.SetFollowNFTURI(&_ILensHub.TransactOpts, profileId, followNFTURI)
}

// SetFollowNFTURIWithSig is a paid mutator transaction binding the contract method 0xbd12d3f0.
//
// Solidity: function setFollowNFTURIWithSig((uint256,string,(uint8,bytes32,bytes32,uint256)) vars) returns()
func (_ILensHub *ILensHubTransactor) SetFollowNFTURIWithSig(opts *bind.TransactOpts, vars DataTypesSetFollowNFTURIWithSigData) (*types.Transaction, error) {
	return _ILensHub.contract.Transact(opts, "setFollowNFTURIWithSig", vars)
}

// SetFollowNFTURIWithSig is a paid mutator transaction binding the contract method 0xbd12d3f0.
//
// Solidity: function setFollowNFTURIWithSig((uint256,string,(uint8,bytes32,bytes32,uint256)) vars) returns()
func (_ILensHub *ILensHubSession) SetFollowNFTURIWithSig(vars DataTypesSetFollowNFTURIWithSigData) (*types.Transaction, error) {
	return _ILensHub.Contract.SetFollowNFTURIWithSig(&_ILensHub.TransactOpts, vars)
}

// SetFollowNFTURIWithSig is a paid mutator transaction binding the contract method 0xbd12d3f0.
//
// Solidity: function setFollowNFTURIWithSig((uint256,string,(uint8,bytes32,bytes32,uint256)) vars) returns()
func (_ILensHub *ILensHubTransactorSession) SetFollowNFTURIWithSig(vars DataTypesSetFollowNFTURIWithSigData) (*types.Transaction, error) {
	return _ILensHub.Contract.SetFollowNFTURIWithSig(&_ILensHub.TransactOpts, vars)
}

// SetGovernance is a paid mutator transaction binding the contract method 0xab033ea9.
//
// Solidity: function setGovernance(address newGovernance) returns()
func (_ILensHub *ILensHubTransactor) SetGovernance(opts *bind.TransactOpts, newGovernance common.Address) (*types.Transaction, error) {
	return _ILensHub.contract.Transact(opts, "setGovernance", newGovernance)
}

// SetGovernance is a paid mutator transaction binding the contract method 0xab033ea9.
//
// Solidity: function setGovernance(address newGovernance) returns()
func (_ILensHub *ILensHubSession) SetGovernance(newGovernance common.Address) (*types.Transaction, error) {
	return _ILensHub.Contract.SetGovernance(&_ILensHub.TransactOpts, newGovernance)
}

// SetGovernance is a paid mutator transaction binding the contract method 0xab033ea9.
//
// Solidity: function setGovernance(address newGovernance) returns()
func (_ILensHub *ILensHubTransactorSession) SetGovernance(newGovernance common.Address) (*types.Transaction, error) {
	return _ILensHub.Contract.SetGovernance(&_ILensHub.TransactOpts, newGovernance)
}

// SetProfileImageURI is a paid mutator transaction binding the contract method 0x054871b8.
//
// Solidity: function setProfileImageURI(uint256 profileId, string imageURI) returns()
func (_ILensHub *ILensHubTransactor) SetProfileImageURI(opts *bind.TransactOpts, profileId *big.Int, imageURI string) (*types.Transaction, error) {
	return _ILensHub.contract.Transact(opts, "setProfileImageURI", profileId, imageURI)
}

// SetProfileImageURI is a paid mutator transaction binding the contract method 0x054871b8.
//
// Solidity: function setProfileImageURI(uint256 profileId, string imageURI) returns()
func (_ILensHub *ILensHubSession) SetProfileImageURI(profileId *big.Int, imageURI string) (*types.Transaction, error) {
	return _ILensHub.Contract.SetProfileImageURI(&_ILensHub.TransactOpts, profileId, imageURI)
}

// SetProfileImageURI is a paid mutator transaction binding the contract method 0x054871b8.
//
// Solidity: function setProfileImageURI(uint256 profileId, string imageURI) returns()
func (_ILensHub *ILensHubTransactorSession) SetProfileImageURI(profileId *big.Int, imageURI string) (*types.Transaction, error) {
	return _ILensHub.Contract.SetProfileImageURI(&_ILensHub.TransactOpts, profileId, imageURI)
}

// SetProfileImageURIWithSig is a paid mutator transaction binding the contract method 0x9b84a14c.
//
// Solidity: function setProfileImageURIWithSig((uint256,string,(uint8,bytes32,bytes32,uint256)) vars) returns()
func (_ILensHub *ILensHubTransactor) SetProfileImageURIWithSig(opts *bind.TransactOpts, vars DataTypesSetProfileImageURIWithSigData) (*types.Transaction, error) {
	return _ILensHub.contract.Transact(opts, "setProfileImageURIWithSig", vars)
}

// SetProfileImageURIWithSig is a paid mutator transaction binding the contract method 0x9b84a14c.
//
// Solidity: function setProfileImageURIWithSig((uint256,string,(uint8,bytes32,bytes32,uint256)) vars) returns()
func (_ILensHub *ILensHubSession) SetProfileImageURIWithSig(vars DataTypesSetProfileImageURIWithSigData) (*types.Transaction, error) {
	return _ILensHub.Contract.SetProfileImageURIWithSig(&_ILensHub.TransactOpts, vars)
}

// SetProfileImageURIWithSig is a paid mutator transaction binding the contract method 0x9b84a14c.
//
// Solidity: function setProfileImageURIWithSig((uint256,string,(uint8,bytes32,bytes32,uint256)) vars) returns()
func (_ILensHub *ILensHubTransactorSession) SetProfileImageURIWithSig(vars DataTypesSetProfileImageURIWithSigData) (*types.Transaction, error) {
	return _ILensHub.Contract.SetProfileImageURIWithSig(&_ILensHub.TransactOpts, vars)
}

// SetState is a paid mutator transaction binding the contract method 0x56de96db.
//
// Solidity: function setState(uint8 newState) returns()
func (_ILensHub *ILensHubTransactor) SetState(opts *bind.TransactOpts, newState uint8) (*types.Transaction, error) {
	return _ILensHub.contract.Transact(opts, "setState", newState)
}

// SetState is a paid mutator transaction binding the contract method 0x56de96db.
//
// Solidity: function setState(uint8 newState) returns()
func (_ILensHub *ILensHubSession) SetState(newState uint8) (*types.Transaction, error) {
	return _ILensHub.Contract.SetState(&_ILensHub.TransactOpts, newState)
}

// SetState is a paid mutator transaction binding the contract method 0x56de96db.
//
// Solidity: function setState(uint8 newState) returns()
func (_ILensHub *ILensHubTransactorSession) SetState(newState uint8) (*types.Transaction, error) {
	return _ILensHub.Contract.SetState(&_ILensHub.TransactOpts, newState)
}

// WhitelistCollectModule is a paid mutator transaction binding the contract method 0x31dcebe3.
//
// Solidity: function whitelistCollectModule(address collectModule, bool whitelist) returns()
func (_ILensHub *ILensHubTransactor) WhitelistCollectModule(opts *bind.TransactOpts, collectModule common.Address, whitelist bool) (*types.Transaction, error) {
	return _ILensHub.contract.Transact(opts, "whitelistCollectModule", collectModule, whitelist)
}

// WhitelistCollectModule is a paid mutator transaction binding the contract method 0x31dcebe3.
//
// Solidity: function whitelistCollectModule(address collectModule, bool whitelist) returns()
func (_ILensHub *ILensHubSession) WhitelistCollectModule(collectModule common.Address, whitelist bool) (*types.Transaction, error) {
	return _ILensHub.Contract.WhitelistCollectModule(&_ILensHub.TransactOpts, collectModule, whitelist)
}

// WhitelistCollectModule is a paid mutator transaction binding the contract method 0x31dcebe3.
//
// Solidity: function whitelistCollectModule(address collectModule, bool whitelist) returns()
func (_ILensHub *ILensHubTransactorSession) WhitelistCollectModule(collectModule common.Address, whitelist bool) (*types.Transaction, error) {
	return _ILensHub.Contract.WhitelistCollectModule(&_ILensHub.TransactOpts, collectModule, whitelist)
}

// WhitelistFollowModule is a paid mutator transaction binding the contract method 0xa6d8e1e4.
//
// Solidity: function whitelistFollowModule(address followModule, bool whitelist) returns()
func (_ILensHub *ILensHubTransactor) WhitelistFollowModule(opts *bind.TransactOpts, followModule common.Address, whitelist bool) (*types.Transaction, error) {
	return _ILensHub.contract.Transact(opts, "whitelistFollowModule", followModule, whitelist)
}

// WhitelistFollowModule is a paid mutator transaction binding the contract method 0xa6d8e1e4.
//
// Solidity: function whitelistFollowModule(address followModule, bool whitelist) returns()
func (_ILensHub *ILensHubSession) WhitelistFollowModule(followModule common.Address, whitelist bool) (*types.Transaction, error) {
	return _ILensHub.Contract.WhitelistFollowModule(&_ILensHub.TransactOpts, followModule, whitelist)
}

// WhitelistFollowModule is a paid mutator transaction binding the contract method 0xa6d8e1e4.
//
// Solidity: function whitelistFollowModule(address followModule, bool whitelist) returns()
func (_ILensHub *ILensHubTransactorSession) WhitelistFollowModule(followModule common.Address, whitelist bool) (*types.Transaction, error) {
	return _ILensHub.Contract.WhitelistFollowModule(&_ILensHub.TransactOpts, followModule, whitelist)
}

// WhitelistProfileCreator is a paid mutator transaction binding the contract method 0x20905506.
//
// Solidity: function whitelistProfileCreator(address profileCreator, bool whitelist) returns()
func (_ILensHub *ILensHubTransactor) WhitelistProfileCreator(opts *bind.TransactOpts, profileCreator common.Address, whitelist bool) (*types.Transaction, error) {
	return _ILensHub.contract.Transact(opts, "whitelistProfileCreator", profileCreator, whitelist)
}

// WhitelistProfileCreator is a paid mutator transaction binding the contract method 0x20905506.
//
// Solidity: function whitelistProfileCreator(address profileCreator, bool whitelist) returns()
func (_ILensHub *ILensHubSession) WhitelistProfileCreator(profileCreator common.Address, whitelist bool) (*types.Transaction, error) {
	return _ILensHub.Contract.WhitelistProfileCreator(&_ILensHub.TransactOpts, profileCreator, whitelist)
}

// WhitelistProfileCreator is a paid mutator transaction binding the contract method 0x20905506.
//
// Solidity: function whitelistProfileCreator(address profileCreator, bool whitelist) returns()
func (_ILensHub *ILensHubTransactorSession) WhitelistProfileCreator(profileCreator common.Address, whitelist bool) (*types.Transaction, error) {
	return _ILensHub.Contract.WhitelistProfileCreator(&_ILensHub.TransactOpts, profileCreator, whitelist)
}

// WhitelistReferenceModule is a paid mutator transaction binding the contract method 0x4187e4c5.
//
// Solidity: function whitelistReferenceModule(address referenceModule, bool whitelist) returns()
func (_ILensHub *ILensHubTransactor) WhitelistReferenceModule(opts *bind.TransactOpts, referenceModule common.Address, whitelist bool) (*types.Transaction, error) {
	return _ILensHub.contract.Transact(opts, "whitelistReferenceModule", referenceModule, whitelist)
}

// WhitelistReferenceModule is a paid mutator transaction binding the contract method 0x4187e4c5.
//
// Solidity: function whitelistReferenceModule(address referenceModule, bool whitelist) returns()
func (_ILensHub *ILensHubSession) WhitelistReferenceModule(referenceModule common.Address, whitelist bool) (*types.Transaction, error) {
	return _ILensHub.Contract.WhitelistReferenceModule(&_ILensHub.TransactOpts, referenceModule, whitelist)
}

// WhitelistReferenceModule is a paid mutator transaction binding the contract method 0x4187e4c5.
//
// Solidity: function whitelistReferenceModule(address referenceModule, bool whitelist) returns()
func (_ILensHub *ILensHubTransactorSession) WhitelistReferenceModule(referenceModule common.Address, whitelist bool) (*types.Transaction, error) {
	return _ILensHub.Contract.WhitelistReferenceModule(&_ILensHub.TransactOpts, referenceModule, whitelist)
}
