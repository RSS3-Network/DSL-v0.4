// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cyberconnect

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
	_ = abi.ConvertType
)

// DataTypesCollectParams is an auto generated low-level Go binding around an user-defined struct.
type DataTypesCollectParams struct {
	Collector common.Address
	ProfileId *big.Int
	EssenceId *big.Int
}

// DataTypesCreateProfileParams is an auto generated low-level Go binding around an user-defined struct.
type DataTypesCreateProfileParams struct {
	To       common.Address
	Handle   string
	Avatar   string
	Metadata string
	Operator common.Address
}

// DataTypesEIP712Signature is an auto generated low-level Go binding around an user-defined struct.
type DataTypesEIP712Signature struct {
	V        uint8
	R        [32]byte
	S        [32]byte
	Deadline *big.Int
}

// DataTypesRegisterEssenceParams is an auto generated low-level Go binding around an user-defined struct.
type DataTypesRegisterEssenceParams struct {
	ProfileId        *big.Int
	Name             string
	Symbol           string
	EssenceTokenURI  string
	EssenceMw        common.Address
	Transferable     bool
	DeployAtRegister bool
}

// DataTypesSubscribeParams is an auto generated low-level Go binding around an user-defined struct.
type DataTypesSubscribeParams struct {
	ProfileIds []*big.Int
}

// CyberConnectMetaData contains all meta data concerning the CyberConnect contract.
var CyberConnectMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collector\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"essenceId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"preData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"postData\",\"type\":\"bytes\"}],\"name\":\"CollectEssence\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"handle\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"avatar\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"}],\"name\":\"CreateProfile\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"essenceId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"essenceNFT\",\"type\":\"address\"}],\"name\":\"DeployEssenceNFT\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"subscribeNFT\",\"type\":\"address\"}],\"name\":\"DeploySubscribeNFT\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"name\":\"Initialize\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"essenceId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"essenceTokenURI\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"essenceMw\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"prepareReturnData\",\"type\":\"bytes\"}],\"name\":\"RegisterEssence\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newAvatar\",\"type\":\"string\"}],\"name\":\"SetAvatar\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"essenceId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"mw\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"prepareReturnData\",\"type\":\"bytes\"}],\"name\":\"SetEssenceData\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newMetadata\",\"type\":\"string\"}],\"name\":\"SetMetadata\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newDescriptor\",\"type\":\"address\"}],\"name\":\"SetNFTDescriptor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"preOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"SetNamespaceOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"prevApproved\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"SetOperatorApproval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"}],\"name\":\"SetPrimaryProfile\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"mw\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"prepareReturnData\",\"type\":\"bytes\"}],\"name\":\"SetSubscribeData\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"profileIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"preDatas\",\"type\":\"bytes[]\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"postDatas\",\"type\":\"bytes[]\"}],\"name\":\"Subscribe\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ENGINE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ESSENCE_BEACON\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SUBSCRIBE_BEACON\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"collector\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"essenceId\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.CollectParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"preData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"postData\",\"type\":\"bytes\"}],\"name\":\"collect\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"collector\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"essenceId\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.CollectParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"preData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"postData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.EIP712Signature\",\"name\":\"sig\",\"type\":\"tuple\"}],\"name\":\"collectWithSig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"handle\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"avatar\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"internalType\":\"structDataTypes.CreateProfileParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"preData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"postData\",\"type\":\"bytes\"}],\"name\":\"createProfile\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"}],\"name\":\"getAvatar\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"essenceId\",\"type\":\"uint256\"}],\"name\":\"getEssenceMw\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"essenceId\",\"type\":\"uint256\"}],\"name\":\"getEssenceNFT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"essenceId\",\"type\":\"uint256\"}],\"name\":\"getEssenceNFTTokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"}],\"name\":\"getHandleByProfileId\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"}],\"name\":\"getMetadata\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNFTDescriptor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNamespaceOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"getOperatorApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getPrimaryProfile\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"handle\",\"type\":\"string\"}],\"name\":\"getProfileIdByHandle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"}],\"name\":\"getSubscribeMw\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"}],\"name\":\"getSubscribeNFT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"}],\"name\":\"getSubscribeNFTTokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"toPause\",\"type\":\"bool\"}],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.EIP712Signature\",\"name\":\"sig\",\"type\":\"tuple\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"essenceTokenURI\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"essenceMw\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"transferable\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"deployAtRegister\",\"type\":\"bool\"}],\"internalType\":\"structDataTypes.RegisterEssenceParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"initData\",\"type\":\"bytes\"}],\"name\":\"registerEssence\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"essenceTokenURI\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"essenceMw\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"transferable\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"deployAtRegister\",\"type\":\"bool\"}],\"internalType\":\"structDataTypes.RegisterEssenceParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"initData\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.EIP712Signature\",\"name\":\"sig\",\"type\":\"tuple\"}],\"name\":\"registerEssenceWithSig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"avatar\",\"type\":\"string\"}],\"name\":\"setAvatar\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"avatar\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.EIP712Signature\",\"name\":\"sig\",\"type\":\"tuple\"}],\"name\":\"setAvatarWithSig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"essenceId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"mw\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"setEssenceData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"essenceId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"mw\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.EIP712Signature\",\"name\":\"sig\",\"type\":\"tuple\"}],\"name\":\"setEssenceDataWithSig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"}],\"name\":\"setMetadata\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.EIP712Signature\",\"name\":\"sig\",\"type\":\"tuple\"}],\"name\":\"setMetadataWithSig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"descriptor\",\"type\":\"address\"}],\"name\":\"setNFTDescriptor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"setNamespaceOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setOperatorApproval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"}],\"name\":\"setPrimaryProfile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.EIP712Signature\",\"name\":\"sig\",\"type\":\"tuple\"}],\"name\":\"setPrimaryProfileWithSig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"mw\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"setSubscribeData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"mw\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.EIP712Signature\",\"name\":\"sig\",\"type\":\"tuple\"}],\"name\":\"setSubscribeDataWithSig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"profileIds\",\"type\":\"uint256[]\"}],\"internalType\":\"structDataTypes.SubscribeParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"preDatas\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"postDatas\",\"type\":\"bytes[]\"}],\"name\":\"subscribe\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"profileIds\",\"type\":\"uint256[]\"}],\"internalType\":\"structDataTypes.SubscribeParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"preDatas\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"postDatas\",\"type\":\"bytes[]\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.EIP712Signature\",\"name\":\"sig\",\"type\":\"tuple\"}],\"name\":\"subscribeWithSig\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBurned\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalMinted\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// CyberConnectABI is the input ABI used to generate the binding from.
// Deprecated: Use CyberConnectMetaData.ABI instead.
var CyberConnectABI = CyberConnectMetaData.ABI

// CyberConnect is an auto generated Go binding around an Ethereum contract.
type CyberConnect struct {
	CyberConnectCaller     // Read-only binding to the contract
	CyberConnectTransactor // Write-only binding to the contract
	CyberConnectFilterer   // Log filterer for contract events
}

// CyberConnectCaller is an auto generated read-only Go binding around an Ethereum contract.
type CyberConnectCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CyberConnectTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CyberConnectTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CyberConnectFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CyberConnectFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CyberConnectSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CyberConnectSession struct {
	Contract     *CyberConnect     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CyberConnectCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CyberConnectCallerSession struct {
	Contract *CyberConnectCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// CyberConnectTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CyberConnectTransactorSession struct {
	Contract     *CyberConnectTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// CyberConnectRaw is an auto generated low-level Go binding around an Ethereum contract.
type CyberConnectRaw struct {
	Contract *CyberConnect // Generic contract binding to access the raw methods on
}

// CyberConnectCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CyberConnectCallerRaw struct {
	Contract *CyberConnectCaller // Generic read-only contract binding to access the raw methods on
}

// CyberConnectTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CyberConnectTransactorRaw struct {
	Contract *CyberConnectTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCyberConnect creates a new instance of CyberConnect, bound to a specific deployed contract.
func NewCyberConnect(address common.Address, backend bind.ContractBackend) (*CyberConnect, error) {
	contract, err := bindCyberConnect(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CyberConnect{CyberConnectCaller: CyberConnectCaller{contract: contract}, CyberConnectTransactor: CyberConnectTransactor{contract: contract}, CyberConnectFilterer: CyberConnectFilterer{contract: contract}}, nil
}

// NewCyberConnectCaller creates a new read-only instance of CyberConnect, bound to a specific deployed contract.
func NewCyberConnectCaller(address common.Address, caller bind.ContractCaller) (*CyberConnectCaller, error) {
	contract, err := bindCyberConnect(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CyberConnectCaller{contract: contract}, nil
}

// NewCyberConnectTransactor creates a new write-only instance of CyberConnect, bound to a specific deployed contract.
func NewCyberConnectTransactor(address common.Address, transactor bind.ContractTransactor) (*CyberConnectTransactor, error) {
	contract, err := bindCyberConnect(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CyberConnectTransactor{contract: contract}, nil
}

// NewCyberConnectFilterer creates a new log filterer instance of CyberConnect, bound to a specific deployed contract.
func NewCyberConnectFilterer(address common.Address, filterer bind.ContractFilterer) (*CyberConnectFilterer, error) {
	contract, err := bindCyberConnect(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CyberConnectFilterer{contract: contract}, nil
}

// bindCyberConnect binds a generic wrapper to an already deployed contract.
func bindCyberConnect(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CyberConnectMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CyberConnect *CyberConnectRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CyberConnect.Contract.CyberConnectCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CyberConnect *CyberConnectRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CyberConnect.Contract.CyberConnectTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CyberConnect *CyberConnectRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CyberConnect.Contract.CyberConnectTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CyberConnect *CyberConnectCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CyberConnect.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CyberConnect *CyberConnectTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CyberConnect.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CyberConnect *CyberConnectTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CyberConnect.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_CyberConnect *CyberConnectCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CyberConnect.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_CyberConnect *CyberConnectSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _CyberConnect.Contract.DOMAINSEPARATOR(&_CyberConnect.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_CyberConnect *CyberConnectCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _CyberConnect.Contract.DOMAINSEPARATOR(&_CyberConnect.CallOpts)
}

// ENGINE is a free data retrieval call binding the contract method 0x4785e8d4.
//
// Solidity: function ENGINE() view returns(address)
func (_CyberConnect *CyberConnectCaller) ENGINE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CyberConnect.contract.Call(opts, &out, "ENGINE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ENGINE is a free data retrieval call binding the contract method 0x4785e8d4.
//
// Solidity: function ENGINE() view returns(address)
func (_CyberConnect *CyberConnectSession) ENGINE() (common.Address, error) {
	return _CyberConnect.Contract.ENGINE(&_CyberConnect.CallOpts)
}

// ENGINE is a free data retrieval call binding the contract method 0x4785e8d4.
//
// Solidity: function ENGINE() view returns(address)
func (_CyberConnect *CyberConnectCallerSession) ENGINE() (common.Address, error) {
	return _CyberConnect.Contract.ENGINE(&_CyberConnect.CallOpts)
}

// ESSENCEBEACON is a free data retrieval call binding the contract method 0x1ec7d106.
//
// Solidity: function ESSENCE_BEACON() view returns(address)
func (_CyberConnect *CyberConnectCaller) ESSENCEBEACON(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CyberConnect.contract.Call(opts, &out, "ESSENCE_BEACON")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ESSENCEBEACON is a free data retrieval call binding the contract method 0x1ec7d106.
//
// Solidity: function ESSENCE_BEACON() view returns(address)
func (_CyberConnect *CyberConnectSession) ESSENCEBEACON() (common.Address, error) {
	return _CyberConnect.Contract.ESSENCEBEACON(&_CyberConnect.CallOpts)
}

// ESSENCEBEACON is a free data retrieval call binding the contract method 0x1ec7d106.
//
// Solidity: function ESSENCE_BEACON() view returns(address)
func (_CyberConnect *CyberConnectCallerSession) ESSENCEBEACON() (common.Address, error) {
	return _CyberConnect.Contract.ESSENCEBEACON(&_CyberConnect.CallOpts)
}

// SUBSCRIBEBEACON is a free data retrieval call binding the contract method 0x5b996c0f.
//
// Solidity: function SUBSCRIBE_BEACON() view returns(address)
func (_CyberConnect *CyberConnectCaller) SUBSCRIBEBEACON(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CyberConnect.contract.Call(opts, &out, "SUBSCRIBE_BEACON")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SUBSCRIBEBEACON is a free data retrieval call binding the contract method 0x5b996c0f.
//
// Solidity: function SUBSCRIBE_BEACON() view returns(address)
func (_CyberConnect *CyberConnectSession) SUBSCRIBEBEACON() (common.Address, error) {
	return _CyberConnect.Contract.SUBSCRIBEBEACON(&_CyberConnect.CallOpts)
}

// SUBSCRIBEBEACON is a free data retrieval call binding the contract method 0x5b996c0f.
//
// Solidity: function SUBSCRIBE_BEACON() view returns(address)
func (_CyberConnect *CyberConnectCallerSession) SUBSCRIBEBEACON() (common.Address, error) {
	return _CyberConnect.Contract.SUBSCRIBEBEACON(&_CyberConnect.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_CyberConnect *CyberConnectCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CyberConnect.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_CyberConnect *CyberConnectSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _CyberConnect.Contract.BalanceOf(&_CyberConnect.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_CyberConnect *CyberConnectCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _CyberConnect.Contract.BalanceOf(&_CyberConnect.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 ) view returns(address)
func (_CyberConnect *CyberConnectCaller) GetApproved(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CyberConnect.contract.Call(opts, &out, "getApproved", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 ) view returns(address)
func (_CyberConnect *CyberConnectSession) GetApproved(arg0 *big.Int) (common.Address, error) {
	return _CyberConnect.Contract.GetApproved(&_CyberConnect.CallOpts, arg0)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 ) view returns(address)
func (_CyberConnect *CyberConnectCallerSession) GetApproved(arg0 *big.Int) (common.Address, error) {
	return _CyberConnect.Contract.GetApproved(&_CyberConnect.CallOpts, arg0)
}

// GetAvatar is a free data retrieval call binding the contract method 0x1328ec9b.
//
// Solidity: function getAvatar(uint256 profileId) view returns(string)
func (_CyberConnect *CyberConnectCaller) GetAvatar(opts *bind.CallOpts, profileId *big.Int) (string, error) {
	var out []interface{}
	err := _CyberConnect.contract.Call(opts, &out, "getAvatar", profileId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetAvatar is a free data retrieval call binding the contract method 0x1328ec9b.
//
// Solidity: function getAvatar(uint256 profileId) view returns(string)
func (_CyberConnect *CyberConnectSession) GetAvatar(profileId *big.Int) (string, error) {
	return _CyberConnect.Contract.GetAvatar(&_CyberConnect.CallOpts, profileId)
}

// GetAvatar is a free data retrieval call binding the contract method 0x1328ec9b.
//
// Solidity: function getAvatar(uint256 profileId) view returns(string)
func (_CyberConnect *CyberConnectCallerSession) GetAvatar(profileId *big.Int) (string, error) {
	return _CyberConnect.Contract.GetAvatar(&_CyberConnect.CallOpts, profileId)
}

// GetEssenceMw is a free data retrieval call binding the contract method 0x5e316de4.
//
// Solidity: function getEssenceMw(uint256 profileId, uint256 essenceId) view returns(address)
func (_CyberConnect *CyberConnectCaller) GetEssenceMw(opts *bind.CallOpts, profileId *big.Int, essenceId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CyberConnect.contract.Call(opts, &out, "getEssenceMw", profileId, essenceId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEssenceMw is a free data retrieval call binding the contract method 0x5e316de4.
//
// Solidity: function getEssenceMw(uint256 profileId, uint256 essenceId) view returns(address)
func (_CyberConnect *CyberConnectSession) GetEssenceMw(profileId *big.Int, essenceId *big.Int) (common.Address, error) {
	return _CyberConnect.Contract.GetEssenceMw(&_CyberConnect.CallOpts, profileId, essenceId)
}

// GetEssenceMw is a free data retrieval call binding the contract method 0x5e316de4.
//
// Solidity: function getEssenceMw(uint256 profileId, uint256 essenceId) view returns(address)
func (_CyberConnect *CyberConnectCallerSession) GetEssenceMw(profileId *big.Int, essenceId *big.Int) (common.Address, error) {
	return _CyberConnect.Contract.GetEssenceMw(&_CyberConnect.CallOpts, profileId, essenceId)
}

// GetEssenceNFT is a free data retrieval call binding the contract method 0xa92dc7e8.
//
// Solidity: function getEssenceNFT(uint256 profileId, uint256 essenceId) view returns(address)
func (_CyberConnect *CyberConnectCaller) GetEssenceNFT(opts *bind.CallOpts, profileId *big.Int, essenceId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CyberConnect.contract.Call(opts, &out, "getEssenceNFT", profileId, essenceId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEssenceNFT is a free data retrieval call binding the contract method 0xa92dc7e8.
//
// Solidity: function getEssenceNFT(uint256 profileId, uint256 essenceId) view returns(address)
func (_CyberConnect *CyberConnectSession) GetEssenceNFT(profileId *big.Int, essenceId *big.Int) (common.Address, error) {
	return _CyberConnect.Contract.GetEssenceNFT(&_CyberConnect.CallOpts, profileId, essenceId)
}

// GetEssenceNFT is a free data retrieval call binding the contract method 0xa92dc7e8.
//
// Solidity: function getEssenceNFT(uint256 profileId, uint256 essenceId) view returns(address)
func (_CyberConnect *CyberConnectCallerSession) GetEssenceNFT(profileId *big.Int, essenceId *big.Int) (common.Address, error) {
	return _CyberConnect.Contract.GetEssenceNFT(&_CyberConnect.CallOpts, profileId, essenceId)
}

// GetEssenceNFTTokenURI is a free data retrieval call binding the contract method 0x206f8e86.
//
// Solidity: function getEssenceNFTTokenURI(uint256 profileId, uint256 essenceId) view returns(string)
func (_CyberConnect *CyberConnectCaller) GetEssenceNFTTokenURI(opts *bind.CallOpts, profileId *big.Int, essenceId *big.Int) (string, error) {
	var out []interface{}
	err := _CyberConnect.contract.Call(opts, &out, "getEssenceNFTTokenURI", profileId, essenceId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetEssenceNFTTokenURI is a free data retrieval call binding the contract method 0x206f8e86.
//
// Solidity: function getEssenceNFTTokenURI(uint256 profileId, uint256 essenceId) view returns(string)
func (_CyberConnect *CyberConnectSession) GetEssenceNFTTokenURI(profileId *big.Int, essenceId *big.Int) (string, error) {
	return _CyberConnect.Contract.GetEssenceNFTTokenURI(&_CyberConnect.CallOpts, profileId, essenceId)
}

// GetEssenceNFTTokenURI is a free data retrieval call binding the contract method 0x206f8e86.
//
// Solidity: function getEssenceNFTTokenURI(uint256 profileId, uint256 essenceId) view returns(string)
func (_CyberConnect *CyberConnectCallerSession) GetEssenceNFTTokenURI(profileId *big.Int, essenceId *big.Int) (string, error) {
	return _CyberConnect.Contract.GetEssenceNFTTokenURI(&_CyberConnect.CallOpts, profileId, essenceId)
}

// GetHandleByProfileId is a free data retrieval call binding the contract method 0x2307c987.
//
// Solidity: function getHandleByProfileId(uint256 profileId) view returns(string)
func (_CyberConnect *CyberConnectCaller) GetHandleByProfileId(opts *bind.CallOpts, profileId *big.Int) (string, error) {
	var out []interface{}
	err := _CyberConnect.contract.Call(opts, &out, "getHandleByProfileId", profileId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetHandleByProfileId is a free data retrieval call binding the contract method 0x2307c987.
//
// Solidity: function getHandleByProfileId(uint256 profileId) view returns(string)
func (_CyberConnect *CyberConnectSession) GetHandleByProfileId(profileId *big.Int) (string, error) {
	return _CyberConnect.Contract.GetHandleByProfileId(&_CyberConnect.CallOpts, profileId)
}

// GetHandleByProfileId is a free data retrieval call binding the contract method 0x2307c987.
//
// Solidity: function getHandleByProfileId(uint256 profileId) view returns(string)
func (_CyberConnect *CyberConnectCallerSession) GetHandleByProfileId(profileId *big.Int) (string, error) {
	return _CyberConnect.Contract.GetHandleByProfileId(&_CyberConnect.CallOpts, profileId)
}

// GetMetadata is a free data retrieval call binding the contract method 0xa574cea4.
//
// Solidity: function getMetadata(uint256 profileId) view returns(string)
func (_CyberConnect *CyberConnectCaller) GetMetadata(opts *bind.CallOpts, profileId *big.Int) (string, error) {
	var out []interface{}
	err := _CyberConnect.contract.Call(opts, &out, "getMetadata", profileId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetMetadata is a free data retrieval call binding the contract method 0xa574cea4.
//
// Solidity: function getMetadata(uint256 profileId) view returns(string)
func (_CyberConnect *CyberConnectSession) GetMetadata(profileId *big.Int) (string, error) {
	return _CyberConnect.Contract.GetMetadata(&_CyberConnect.CallOpts, profileId)
}

// GetMetadata is a free data retrieval call binding the contract method 0xa574cea4.
//
// Solidity: function getMetadata(uint256 profileId) view returns(string)
func (_CyberConnect *CyberConnectCallerSession) GetMetadata(profileId *big.Int) (string, error) {
	return _CyberConnect.Contract.GetMetadata(&_CyberConnect.CallOpts, profileId)
}

// GetNFTDescriptor is a free data retrieval call binding the contract method 0x17c4d5d2.
//
// Solidity: function getNFTDescriptor() view returns(address)
func (_CyberConnect *CyberConnectCaller) GetNFTDescriptor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CyberConnect.contract.Call(opts, &out, "getNFTDescriptor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetNFTDescriptor is a free data retrieval call binding the contract method 0x17c4d5d2.
//
// Solidity: function getNFTDescriptor() view returns(address)
func (_CyberConnect *CyberConnectSession) GetNFTDescriptor() (common.Address, error) {
	return _CyberConnect.Contract.GetNFTDescriptor(&_CyberConnect.CallOpts)
}

// GetNFTDescriptor is a free data retrieval call binding the contract method 0x17c4d5d2.
//
// Solidity: function getNFTDescriptor() view returns(address)
func (_CyberConnect *CyberConnectCallerSession) GetNFTDescriptor() (common.Address, error) {
	return _CyberConnect.Contract.GetNFTDescriptor(&_CyberConnect.CallOpts)
}

// GetNamespaceOwner is a free data retrieval call binding the contract method 0x42cfa815.
//
// Solidity: function getNamespaceOwner() view returns(address)
func (_CyberConnect *CyberConnectCaller) GetNamespaceOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CyberConnect.contract.Call(opts, &out, "getNamespaceOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetNamespaceOwner is a free data retrieval call binding the contract method 0x42cfa815.
//
// Solidity: function getNamespaceOwner() view returns(address)
func (_CyberConnect *CyberConnectSession) GetNamespaceOwner() (common.Address, error) {
	return _CyberConnect.Contract.GetNamespaceOwner(&_CyberConnect.CallOpts)
}

// GetNamespaceOwner is a free data retrieval call binding the contract method 0x42cfa815.
//
// Solidity: function getNamespaceOwner() view returns(address)
func (_CyberConnect *CyberConnectCallerSession) GetNamespaceOwner() (common.Address, error) {
	return _CyberConnect.Contract.GetNamespaceOwner(&_CyberConnect.CallOpts)
}

// GetOperatorApproval is a free data retrieval call binding the contract method 0xf0ab9600.
//
// Solidity: function getOperatorApproval(uint256 profileId, address operator) view returns(bool)
func (_CyberConnect *CyberConnectCaller) GetOperatorApproval(opts *bind.CallOpts, profileId *big.Int, operator common.Address) (bool, error) {
	var out []interface{}
	err := _CyberConnect.contract.Call(opts, &out, "getOperatorApproval", profileId, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetOperatorApproval is a free data retrieval call binding the contract method 0xf0ab9600.
//
// Solidity: function getOperatorApproval(uint256 profileId, address operator) view returns(bool)
func (_CyberConnect *CyberConnectSession) GetOperatorApproval(profileId *big.Int, operator common.Address) (bool, error) {
	return _CyberConnect.Contract.GetOperatorApproval(&_CyberConnect.CallOpts, profileId, operator)
}

// GetOperatorApproval is a free data retrieval call binding the contract method 0xf0ab9600.
//
// Solidity: function getOperatorApproval(uint256 profileId, address operator) view returns(bool)
func (_CyberConnect *CyberConnectCallerSession) GetOperatorApproval(profileId *big.Int, operator common.Address) (bool, error) {
	return _CyberConnect.Contract.GetOperatorApproval(&_CyberConnect.CallOpts, profileId, operator)
}

// GetPrimaryProfile is a free data retrieval call binding the contract method 0xdfa52f07.
//
// Solidity: function getPrimaryProfile(address user) view returns(uint256)
func (_CyberConnect *CyberConnectCaller) GetPrimaryProfile(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CyberConnect.contract.Call(opts, &out, "getPrimaryProfile", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrimaryProfile is a free data retrieval call binding the contract method 0xdfa52f07.
//
// Solidity: function getPrimaryProfile(address user) view returns(uint256)
func (_CyberConnect *CyberConnectSession) GetPrimaryProfile(user common.Address) (*big.Int, error) {
	return _CyberConnect.Contract.GetPrimaryProfile(&_CyberConnect.CallOpts, user)
}

// GetPrimaryProfile is a free data retrieval call binding the contract method 0xdfa52f07.
//
// Solidity: function getPrimaryProfile(address user) view returns(uint256)
func (_CyberConnect *CyberConnectCallerSession) GetPrimaryProfile(user common.Address) (*big.Int, error) {
	return _CyberConnect.Contract.GetPrimaryProfile(&_CyberConnect.CallOpts, user)
}

// GetProfileIdByHandle is a free data retrieval call binding the contract method 0x20fa728a.
//
// Solidity: function getProfileIdByHandle(string handle) view returns(uint256)
func (_CyberConnect *CyberConnectCaller) GetProfileIdByHandle(opts *bind.CallOpts, handle string) (*big.Int, error) {
	var out []interface{}
	err := _CyberConnect.contract.Call(opts, &out, "getProfileIdByHandle", handle)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProfileIdByHandle is a free data retrieval call binding the contract method 0x20fa728a.
//
// Solidity: function getProfileIdByHandle(string handle) view returns(uint256)
func (_CyberConnect *CyberConnectSession) GetProfileIdByHandle(handle string) (*big.Int, error) {
	return _CyberConnect.Contract.GetProfileIdByHandle(&_CyberConnect.CallOpts, handle)
}

// GetProfileIdByHandle is a free data retrieval call binding the contract method 0x20fa728a.
//
// Solidity: function getProfileIdByHandle(string handle) view returns(uint256)
func (_CyberConnect *CyberConnectCallerSession) GetProfileIdByHandle(handle string) (*big.Int, error) {
	return _CyberConnect.Contract.GetProfileIdByHandle(&_CyberConnect.CallOpts, handle)
}

// GetSubscribeMw is a free data retrieval call binding the contract method 0x56bfd88c.
//
// Solidity: function getSubscribeMw(uint256 profileId) view returns(address)
func (_CyberConnect *CyberConnectCaller) GetSubscribeMw(opts *bind.CallOpts, profileId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CyberConnect.contract.Call(opts, &out, "getSubscribeMw", profileId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSubscribeMw is a free data retrieval call binding the contract method 0x56bfd88c.
//
// Solidity: function getSubscribeMw(uint256 profileId) view returns(address)
func (_CyberConnect *CyberConnectSession) GetSubscribeMw(profileId *big.Int) (common.Address, error) {
	return _CyberConnect.Contract.GetSubscribeMw(&_CyberConnect.CallOpts, profileId)
}

// GetSubscribeMw is a free data retrieval call binding the contract method 0x56bfd88c.
//
// Solidity: function getSubscribeMw(uint256 profileId) view returns(address)
func (_CyberConnect *CyberConnectCallerSession) GetSubscribeMw(profileId *big.Int) (common.Address, error) {
	return _CyberConnect.Contract.GetSubscribeMw(&_CyberConnect.CallOpts, profileId)
}

// GetSubscribeNFT is a free data retrieval call binding the contract method 0xf4954aef.
//
// Solidity: function getSubscribeNFT(uint256 profileId) view returns(address)
func (_CyberConnect *CyberConnectCaller) GetSubscribeNFT(opts *bind.CallOpts, profileId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CyberConnect.contract.Call(opts, &out, "getSubscribeNFT", profileId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSubscribeNFT is a free data retrieval call binding the contract method 0xf4954aef.
//
// Solidity: function getSubscribeNFT(uint256 profileId) view returns(address)
func (_CyberConnect *CyberConnectSession) GetSubscribeNFT(profileId *big.Int) (common.Address, error) {
	return _CyberConnect.Contract.GetSubscribeNFT(&_CyberConnect.CallOpts, profileId)
}

// GetSubscribeNFT is a free data retrieval call binding the contract method 0xf4954aef.
//
// Solidity: function getSubscribeNFT(uint256 profileId) view returns(address)
func (_CyberConnect *CyberConnectCallerSession) GetSubscribeNFT(profileId *big.Int) (common.Address, error) {
	return _CyberConnect.Contract.GetSubscribeNFT(&_CyberConnect.CallOpts, profileId)
}

// GetSubscribeNFTTokenURI is a free data retrieval call binding the contract method 0x303c14f8.
//
// Solidity: function getSubscribeNFTTokenURI(uint256 profileId) view returns(string)
func (_CyberConnect *CyberConnectCaller) GetSubscribeNFTTokenURI(opts *bind.CallOpts, profileId *big.Int) (string, error) {
	var out []interface{}
	err := _CyberConnect.contract.Call(opts, &out, "getSubscribeNFTTokenURI", profileId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetSubscribeNFTTokenURI is a free data retrieval call binding the contract method 0x303c14f8.
//
// Solidity: function getSubscribeNFTTokenURI(uint256 profileId) view returns(string)
func (_CyberConnect *CyberConnectSession) GetSubscribeNFTTokenURI(profileId *big.Int) (string, error) {
	return _CyberConnect.Contract.GetSubscribeNFTTokenURI(&_CyberConnect.CallOpts, profileId)
}

// GetSubscribeNFTTokenURI is a free data retrieval call binding the contract method 0x303c14f8.
//
// Solidity: function getSubscribeNFTTokenURI(uint256 profileId) view returns(string)
func (_CyberConnect *CyberConnectCallerSession) GetSubscribeNFTTokenURI(profileId *big.Int) (string, error) {
	return _CyberConnect.Contract.GetSubscribeNFTTokenURI(&_CyberConnect.CallOpts, profileId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address , address ) view returns(bool)
func (_CyberConnect *CyberConnectCaller) IsApprovedForAll(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _CyberConnect.contract.Call(opts, &out, "isApprovedForAll", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address , address ) view returns(bool)
func (_CyberConnect *CyberConnectSession) IsApprovedForAll(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _CyberConnect.Contract.IsApprovedForAll(&_CyberConnect.CallOpts, arg0, arg1)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address , address ) view returns(bool)
func (_CyberConnect *CyberConnectCallerSession) IsApprovedForAll(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _CyberConnect.Contract.IsApprovedForAll(&_CyberConnect.CallOpts, arg0, arg1)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CyberConnect *CyberConnectCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CyberConnect.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CyberConnect *CyberConnectSession) Name() (string, error) {
	return _CyberConnect.Contract.Name(&_CyberConnect.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CyberConnect *CyberConnectCallerSession) Name() (string, error) {
	return _CyberConnect.Contract.Name(&_CyberConnect.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_CyberConnect *CyberConnectCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CyberConnect.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_CyberConnect *CyberConnectSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _CyberConnect.Contract.Nonces(&_CyberConnect.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_CyberConnect *CyberConnectCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _CyberConnect.Contract.Nonces(&_CyberConnect.CallOpts, arg0)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 id) view returns(address owner)
func (_CyberConnect *CyberConnectCaller) OwnerOf(opts *bind.CallOpts, id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CyberConnect.contract.Call(opts, &out, "ownerOf", id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 id) view returns(address owner)
func (_CyberConnect *CyberConnectSession) OwnerOf(id *big.Int) (common.Address, error) {
	return _CyberConnect.Contract.OwnerOf(&_CyberConnect.CallOpts, id)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 id) view returns(address owner)
func (_CyberConnect *CyberConnectCallerSession) OwnerOf(id *big.Int) (common.Address, error) {
	return _CyberConnect.Contract.OwnerOf(&_CyberConnect.CallOpts, id)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CyberConnect *CyberConnectCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CyberConnect.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CyberConnect *CyberConnectSession) Paused() (bool, error) {
	return _CyberConnect.Contract.Paused(&_CyberConnect.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CyberConnect *CyberConnectCallerSession) Paused() (bool, error) {
	return _CyberConnect.Contract.Paused(&_CyberConnect.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_CyberConnect *CyberConnectCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CyberConnect.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_CyberConnect *CyberConnectSession) ProxiableUUID() ([32]byte, error) {
	return _CyberConnect.Contract.ProxiableUUID(&_CyberConnect.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_CyberConnect *CyberConnectCallerSession) ProxiableUUID() ([32]byte, error) {
	return _CyberConnect.Contract.ProxiableUUID(&_CyberConnect.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CyberConnect *CyberConnectCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _CyberConnect.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CyberConnect *CyberConnectSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CyberConnect.Contract.SupportsInterface(&_CyberConnect.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CyberConnect *CyberConnectCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CyberConnect.Contract.SupportsInterface(&_CyberConnect.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CyberConnect *CyberConnectCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CyberConnect.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CyberConnect *CyberConnectSession) Symbol() (string, error) {
	return _CyberConnect.Contract.Symbol(&_CyberConnect.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CyberConnect *CyberConnectCallerSession) Symbol() (string, error) {
	return _CyberConnect.Contract.Symbol(&_CyberConnect.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_CyberConnect *CyberConnectCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _CyberConnect.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_CyberConnect *CyberConnectSession) TokenURI(tokenId *big.Int) (string, error) {
	return _CyberConnect.Contract.TokenURI(&_CyberConnect.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_CyberConnect *CyberConnectCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _CyberConnect.Contract.TokenURI(&_CyberConnect.CallOpts, tokenId)
}

// TotalBurned is a free data retrieval call binding the contract method 0xd89135cd.
//
// Solidity: function totalBurned() view returns(uint256)
func (_CyberConnect *CyberConnectCaller) TotalBurned(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CyberConnect.contract.Call(opts, &out, "totalBurned")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBurned is a free data retrieval call binding the contract method 0xd89135cd.
//
// Solidity: function totalBurned() view returns(uint256)
func (_CyberConnect *CyberConnectSession) TotalBurned() (*big.Int, error) {
	return _CyberConnect.Contract.TotalBurned(&_CyberConnect.CallOpts)
}

// TotalBurned is a free data retrieval call binding the contract method 0xd89135cd.
//
// Solidity: function totalBurned() view returns(uint256)
func (_CyberConnect *CyberConnectCallerSession) TotalBurned() (*big.Int, error) {
	return _CyberConnect.Contract.TotalBurned(&_CyberConnect.CallOpts)
}

// TotalMinted is a free data retrieval call binding the contract method 0xa2309ff8.
//
// Solidity: function totalMinted() view returns(uint256)
func (_CyberConnect *CyberConnectCaller) TotalMinted(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CyberConnect.contract.Call(opts, &out, "totalMinted")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalMinted is a free data retrieval call binding the contract method 0xa2309ff8.
//
// Solidity: function totalMinted() view returns(uint256)
func (_CyberConnect *CyberConnectSession) TotalMinted() (*big.Int, error) {
	return _CyberConnect.Contract.TotalMinted(&_CyberConnect.CallOpts)
}

// TotalMinted is a free data retrieval call binding the contract method 0xa2309ff8.
//
// Solidity: function totalMinted() view returns(uint256)
func (_CyberConnect *CyberConnectCallerSession) TotalMinted() (*big.Int, error) {
	return _CyberConnect.Contract.TotalMinted(&_CyberConnect.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CyberConnect *CyberConnectCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CyberConnect.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CyberConnect *CyberConnectSession) TotalSupply() (*big.Int, error) {
	return _CyberConnect.Contract.TotalSupply(&_CyberConnect.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CyberConnect *CyberConnectCallerSession) TotalSupply() (*big.Int, error) {
	return _CyberConnect.Contract.TotalSupply(&_CyberConnect.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint256)
func (_CyberConnect *CyberConnectCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CyberConnect.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint256)
func (_CyberConnect *CyberConnectSession) Version() (*big.Int, error) {
	return _CyberConnect.Contract.Version(&_CyberConnect.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint256)
func (_CyberConnect *CyberConnectCallerSession) Version() (*big.Int, error) {
	return _CyberConnect.Contract.Version(&_CyberConnect.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 id) returns()
func (_CyberConnect *CyberConnectTransactor) Approve(opts *bind.TransactOpts, spender common.Address, id *big.Int) (*types.Transaction, error) {
	return _CyberConnect.contract.Transact(opts, "approve", spender, id)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 id) returns()
func (_CyberConnect *CyberConnectSession) Approve(spender common.Address, id *big.Int) (*types.Transaction, error) {
	return _CyberConnect.Contract.Approve(&_CyberConnect.TransactOpts, spender, id)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 id) returns()
func (_CyberConnect *CyberConnectTransactorSession) Approve(spender common.Address, id *big.Int) (*types.Transaction, error) {
	return _CyberConnect.Contract.Approve(&_CyberConnect.TransactOpts, spender, id)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_CyberConnect *CyberConnectTransactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _CyberConnect.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_CyberConnect *CyberConnectSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _CyberConnect.Contract.Burn(&_CyberConnect.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_CyberConnect *CyberConnectTransactorSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _CyberConnect.Contract.Burn(&_CyberConnect.TransactOpts, tokenId)
}

// Collect is a paid mutator transaction binding the contract method 0xbe10bc61.
//
// Solidity: function collect((address,uint256,uint256) params, bytes preData, bytes postData) returns(uint256 tokenId)
func (_CyberConnect *CyberConnectTransactor) Collect(opts *bind.TransactOpts, params DataTypesCollectParams, preData []byte, postData []byte) (*types.Transaction, error) {
	return _CyberConnect.contract.Transact(opts, "collect", params, preData, postData)
}

// Collect is a paid mutator transaction binding the contract method 0xbe10bc61.
//
// Solidity: function collect((address,uint256,uint256) params, bytes preData, bytes postData) returns(uint256 tokenId)
func (_CyberConnect *CyberConnectSession) Collect(params DataTypesCollectParams, preData []byte, postData []byte) (*types.Transaction, error) {
	return _CyberConnect.Contract.Collect(&_CyberConnect.TransactOpts, params, preData, postData)
}

// Collect is a paid mutator transaction binding the contract method 0xbe10bc61.
//
// Solidity: function collect((address,uint256,uint256) params, bytes preData, bytes postData) returns(uint256 tokenId)
func (_CyberConnect *CyberConnectTransactorSession) Collect(params DataTypesCollectParams, preData []byte, postData []byte) (*types.Transaction, error) {
	return _CyberConnect.Contract.Collect(&_CyberConnect.TransactOpts, params, preData, postData)
}

// CollectWithSig is a paid mutator transaction binding the contract method 0x54036644.
//
// Solidity: function collectWithSig((address,uint256,uint256) params, bytes preData, bytes postData, address sender, (uint8,bytes32,bytes32,uint256) sig) returns(uint256 tokenId)
func (_CyberConnect *CyberConnectTransactor) CollectWithSig(opts *bind.TransactOpts, params DataTypesCollectParams, preData []byte, postData []byte, sender common.Address, sig DataTypesEIP712Signature) (*types.Transaction, error) {
	return _CyberConnect.contract.Transact(opts, "collectWithSig", params, preData, postData, sender, sig)
}

// CollectWithSig is a paid mutator transaction binding the contract method 0x54036644.
//
// Solidity: function collectWithSig((address,uint256,uint256) params, bytes preData, bytes postData, address sender, (uint8,bytes32,bytes32,uint256) sig) returns(uint256 tokenId)
func (_CyberConnect *CyberConnectSession) CollectWithSig(params DataTypesCollectParams, preData []byte, postData []byte, sender common.Address, sig DataTypesEIP712Signature) (*types.Transaction, error) {
	return _CyberConnect.Contract.CollectWithSig(&_CyberConnect.TransactOpts, params, preData, postData, sender, sig)
}

// CollectWithSig is a paid mutator transaction binding the contract method 0x54036644.
//
// Solidity: function collectWithSig((address,uint256,uint256) params, bytes preData, bytes postData, address sender, (uint8,bytes32,bytes32,uint256) sig) returns(uint256 tokenId)
func (_CyberConnect *CyberConnectTransactorSession) CollectWithSig(params DataTypesCollectParams, preData []byte, postData []byte, sender common.Address, sig DataTypesEIP712Signature) (*types.Transaction, error) {
	return _CyberConnect.Contract.CollectWithSig(&_CyberConnect.TransactOpts, params, preData, postData, sender, sig)
}

// CreateProfile is a paid mutator transaction binding the contract method 0x166fad6f.
//
// Solidity: function createProfile((address,string,string,string,address) params, bytes preData, bytes postData) payable returns(uint256 tokenID)
func (_CyberConnect *CyberConnectTransactor) CreateProfile(opts *bind.TransactOpts, params DataTypesCreateProfileParams, preData []byte, postData []byte) (*types.Transaction, error) {
	return _CyberConnect.contract.Transact(opts, "createProfile", params, preData, postData)
}

// CreateProfile is a paid mutator transaction binding the contract method 0x166fad6f.
//
// Solidity: function createProfile((address,string,string,string,address) params, bytes preData, bytes postData) payable returns(uint256 tokenID)
func (_CyberConnect *CyberConnectSession) CreateProfile(params DataTypesCreateProfileParams, preData []byte, postData []byte) (*types.Transaction, error) {
	return _CyberConnect.Contract.CreateProfile(&_CyberConnect.TransactOpts, params, preData, postData)
}

// CreateProfile is a paid mutator transaction binding the contract method 0x166fad6f.
//
// Solidity: function createProfile((address,string,string,string,address) params, bytes preData, bytes postData) payable returns(uint256 tokenID)
func (_CyberConnect *CyberConnectTransactorSession) CreateProfile(params DataTypesCreateProfileParams, preData []byte, postData []byte) (*types.Transaction, error) {
	return _CyberConnect.Contract.CreateProfile(&_CyberConnect.TransactOpts, params, preData, postData)
}

// Initialize is a paid mutator transaction binding the contract method 0x90657147.
//
// Solidity: function initialize(address _owner, string name, string symbol) returns()
func (_CyberConnect *CyberConnectTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address, name string, symbol string) (*types.Transaction, error) {
	return _CyberConnect.contract.Transact(opts, "initialize", _owner, name, symbol)
}

// Initialize is a paid mutator transaction binding the contract method 0x90657147.
//
// Solidity: function initialize(address _owner, string name, string symbol) returns()
func (_CyberConnect *CyberConnectSession) Initialize(_owner common.Address, name string, symbol string) (*types.Transaction, error) {
	return _CyberConnect.Contract.Initialize(&_CyberConnect.TransactOpts, _owner, name, symbol)
}

// Initialize is a paid mutator transaction binding the contract method 0x90657147.
//
// Solidity: function initialize(address _owner, string name, string symbol) returns()
func (_CyberConnect *CyberConnectTransactorSession) Initialize(_owner common.Address, name string, symbol string) (*types.Transaction, error) {
	return _CyberConnect.Contract.Initialize(&_CyberConnect.TransactOpts, _owner, name, symbol)
}

// Pause is a paid mutator transaction binding the contract method 0x02329a29.
//
// Solidity: function pause(bool toPause) returns()
func (_CyberConnect *CyberConnectTransactor) Pause(opts *bind.TransactOpts, toPause bool) (*types.Transaction, error) {
	return _CyberConnect.contract.Transact(opts, "pause", toPause)
}

// Pause is a paid mutator transaction binding the contract method 0x02329a29.
//
// Solidity: function pause(bool toPause) returns()
func (_CyberConnect *CyberConnectSession) Pause(toPause bool) (*types.Transaction, error) {
	return _CyberConnect.Contract.Pause(&_CyberConnect.TransactOpts, toPause)
}

// Pause is a paid mutator transaction binding the contract method 0x02329a29.
//
// Solidity: function pause(bool toPause) returns()
func (_CyberConnect *CyberConnectTransactorSession) Pause(toPause bool) (*types.Transaction, error) {
	return _CyberConnect.Contract.Pause(&_CyberConnect.TransactOpts, toPause)
}

// Permit is a paid mutator transaction binding the contract method 0x7ef67f99.
//
// Solidity: function permit(address spender, uint256 tokenId, (uint8,bytes32,bytes32,uint256) sig) returns()
func (_CyberConnect *CyberConnectTransactor) Permit(opts *bind.TransactOpts, spender common.Address, tokenId *big.Int, sig DataTypesEIP712Signature) (*types.Transaction, error) {
	return _CyberConnect.contract.Transact(opts, "permit", spender, tokenId, sig)
}

// Permit is a paid mutator transaction binding the contract method 0x7ef67f99.
//
// Solidity: function permit(address spender, uint256 tokenId, (uint8,bytes32,bytes32,uint256) sig) returns()
func (_CyberConnect *CyberConnectSession) Permit(spender common.Address, tokenId *big.Int, sig DataTypesEIP712Signature) (*types.Transaction, error) {
	return _CyberConnect.Contract.Permit(&_CyberConnect.TransactOpts, spender, tokenId, sig)
}

// Permit is a paid mutator transaction binding the contract method 0x7ef67f99.
//
// Solidity: function permit(address spender, uint256 tokenId, (uint8,bytes32,bytes32,uint256) sig) returns()
func (_CyberConnect *CyberConnectTransactorSession) Permit(spender common.Address, tokenId *big.Int, sig DataTypesEIP712Signature) (*types.Transaction, error) {
	return _CyberConnect.Contract.Permit(&_CyberConnect.TransactOpts, spender, tokenId, sig)
}

// RegisterEssence is a paid mutator transaction binding the contract method 0x71fd0a98.
//
// Solidity: function registerEssence((uint256,string,string,string,address,bool,bool) params, bytes initData) returns(uint256)
func (_CyberConnect *CyberConnectTransactor) RegisterEssence(opts *bind.TransactOpts, params DataTypesRegisterEssenceParams, initData []byte) (*types.Transaction, error) {
	return _CyberConnect.contract.Transact(opts, "registerEssence", params, initData)
}

// RegisterEssence is a paid mutator transaction binding the contract method 0x71fd0a98.
//
// Solidity: function registerEssence((uint256,string,string,string,address,bool,bool) params, bytes initData) returns(uint256)
func (_CyberConnect *CyberConnectSession) RegisterEssence(params DataTypesRegisterEssenceParams, initData []byte) (*types.Transaction, error) {
	return _CyberConnect.Contract.RegisterEssence(&_CyberConnect.TransactOpts, params, initData)
}

// RegisterEssence is a paid mutator transaction binding the contract method 0x71fd0a98.
//
// Solidity: function registerEssence((uint256,string,string,string,address,bool,bool) params, bytes initData) returns(uint256)
func (_CyberConnect *CyberConnectTransactorSession) RegisterEssence(params DataTypesRegisterEssenceParams, initData []byte) (*types.Transaction, error) {
	return _CyberConnect.Contract.RegisterEssence(&_CyberConnect.TransactOpts, params, initData)
}

// RegisterEssenceWithSig is a paid mutator transaction binding the contract method 0xe5fee479.
//
// Solidity: function registerEssenceWithSig((uint256,string,string,string,address,bool,bool) params, bytes initData, (uint8,bytes32,bytes32,uint256) sig) returns(uint256 tokenId)
func (_CyberConnect *CyberConnectTransactor) RegisterEssenceWithSig(opts *bind.TransactOpts, params DataTypesRegisterEssenceParams, initData []byte, sig DataTypesEIP712Signature) (*types.Transaction, error) {
	return _CyberConnect.contract.Transact(opts, "registerEssenceWithSig", params, initData, sig)
}

// RegisterEssenceWithSig is a paid mutator transaction binding the contract method 0xe5fee479.
//
// Solidity: function registerEssenceWithSig((uint256,string,string,string,address,bool,bool) params, bytes initData, (uint8,bytes32,bytes32,uint256) sig) returns(uint256 tokenId)
func (_CyberConnect *CyberConnectSession) RegisterEssenceWithSig(params DataTypesRegisterEssenceParams, initData []byte, sig DataTypesEIP712Signature) (*types.Transaction, error) {
	return _CyberConnect.Contract.RegisterEssenceWithSig(&_CyberConnect.TransactOpts, params, initData, sig)
}

// RegisterEssenceWithSig is a paid mutator transaction binding the contract method 0xe5fee479.
//
// Solidity: function registerEssenceWithSig((uint256,string,string,string,address,bool,bool) params, bytes initData, (uint8,bytes32,bytes32,uint256) sig) returns(uint256 tokenId)
func (_CyberConnect *CyberConnectTransactorSession) RegisterEssenceWithSig(params DataTypesRegisterEssenceParams, initData []byte, sig DataTypesEIP712Signature) (*types.Transaction, error) {
	return _CyberConnect.Contract.RegisterEssenceWithSig(&_CyberConnect.TransactOpts, params, initData, sig)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id) returns()
func (_CyberConnect *CyberConnectTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _CyberConnect.contract.Transact(opts, "safeTransferFrom", from, to, id)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id) returns()
func (_CyberConnect *CyberConnectSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _CyberConnect.Contract.SafeTransferFrom(&_CyberConnect.TransactOpts, from, to, id)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id) returns()
func (_CyberConnect *CyberConnectTransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _CyberConnect.Contract.SafeTransferFrom(&_CyberConnect.TransactOpts, from, to, id)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, bytes data) returns()
func (_CyberConnect *CyberConnectTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, data []byte) (*types.Transaction, error) {
	return _CyberConnect.contract.Transact(opts, "safeTransferFrom0", from, to, id, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, bytes data) returns()
func (_CyberConnect *CyberConnectSession) SafeTransferFrom0(from common.Address, to common.Address, id *big.Int, data []byte) (*types.Transaction, error) {
	return _CyberConnect.Contract.SafeTransferFrom0(&_CyberConnect.TransactOpts, from, to, id, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, bytes data) returns()
func (_CyberConnect *CyberConnectTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, id *big.Int, data []byte) (*types.Transaction, error) {
	return _CyberConnect.Contract.SafeTransferFrom0(&_CyberConnect.TransactOpts, from, to, id, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_CyberConnect *CyberConnectTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _CyberConnect.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_CyberConnect *CyberConnectSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _CyberConnect.Contract.SetApprovalForAll(&_CyberConnect.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_CyberConnect *CyberConnectTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _CyberConnect.Contract.SetApprovalForAll(&_CyberConnect.TransactOpts, operator, approved)
}

// SetAvatar is a paid mutator transaction binding the contract method 0x87beeb2b.
//
// Solidity: function setAvatar(uint256 profileId, string avatar) returns()
func (_CyberConnect *CyberConnectTransactor) SetAvatar(opts *bind.TransactOpts, profileId *big.Int, avatar string) (*types.Transaction, error) {
	return _CyberConnect.contract.Transact(opts, "setAvatar", profileId, avatar)
}

// SetAvatar is a paid mutator transaction binding the contract method 0x87beeb2b.
//
// Solidity: function setAvatar(uint256 profileId, string avatar) returns()
func (_CyberConnect *CyberConnectSession) SetAvatar(profileId *big.Int, avatar string) (*types.Transaction, error) {
	return _CyberConnect.Contract.SetAvatar(&_CyberConnect.TransactOpts, profileId, avatar)
}

// SetAvatar is a paid mutator transaction binding the contract method 0x87beeb2b.
//
// Solidity: function setAvatar(uint256 profileId, string avatar) returns()
func (_CyberConnect *CyberConnectTransactorSession) SetAvatar(profileId *big.Int, avatar string) (*types.Transaction, error) {
	return _CyberConnect.Contract.SetAvatar(&_CyberConnect.TransactOpts, profileId, avatar)
}

// SetAvatarWithSig is a paid mutator transaction binding the contract method 0x6463788d.
//
// Solidity: function setAvatarWithSig(uint256 profileId, string avatar, (uint8,bytes32,bytes32,uint256) sig) returns()
func (_CyberConnect *CyberConnectTransactor) SetAvatarWithSig(opts *bind.TransactOpts, profileId *big.Int, avatar string, sig DataTypesEIP712Signature) (*types.Transaction, error) {
	return _CyberConnect.contract.Transact(opts, "setAvatarWithSig", profileId, avatar, sig)
}

// SetAvatarWithSig is a paid mutator transaction binding the contract method 0x6463788d.
//
// Solidity: function setAvatarWithSig(uint256 profileId, string avatar, (uint8,bytes32,bytes32,uint256) sig) returns()
func (_CyberConnect *CyberConnectSession) SetAvatarWithSig(profileId *big.Int, avatar string, sig DataTypesEIP712Signature) (*types.Transaction, error) {
	return _CyberConnect.Contract.SetAvatarWithSig(&_CyberConnect.TransactOpts, profileId, avatar, sig)
}

// SetAvatarWithSig is a paid mutator transaction binding the contract method 0x6463788d.
//
// Solidity: function setAvatarWithSig(uint256 profileId, string avatar, (uint8,bytes32,bytes32,uint256) sig) returns()
func (_CyberConnect *CyberConnectTransactorSession) SetAvatarWithSig(profileId *big.Int, avatar string, sig DataTypesEIP712Signature) (*types.Transaction, error) {
	return _CyberConnect.Contract.SetAvatarWithSig(&_CyberConnect.TransactOpts, profileId, avatar, sig)
}

// SetEssenceData is a paid mutator transaction binding the contract method 0x0eee16f4.
//
// Solidity: function setEssenceData(uint256 profileId, uint256 essenceId, string uri, address mw, bytes data) returns()
func (_CyberConnect *CyberConnectTransactor) SetEssenceData(opts *bind.TransactOpts, profileId *big.Int, essenceId *big.Int, uri string, mw common.Address, data []byte) (*types.Transaction, error) {
	return _CyberConnect.contract.Transact(opts, "setEssenceData", profileId, essenceId, uri, mw, data)
}

// SetEssenceData is a paid mutator transaction binding the contract method 0x0eee16f4.
//
// Solidity: function setEssenceData(uint256 profileId, uint256 essenceId, string uri, address mw, bytes data) returns()
func (_CyberConnect *CyberConnectSession) SetEssenceData(profileId *big.Int, essenceId *big.Int, uri string, mw common.Address, data []byte) (*types.Transaction, error) {
	return _CyberConnect.Contract.SetEssenceData(&_CyberConnect.TransactOpts, profileId, essenceId, uri, mw, data)
}

// SetEssenceData is a paid mutator transaction binding the contract method 0x0eee16f4.
//
// Solidity: function setEssenceData(uint256 profileId, uint256 essenceId, string uri, address mw, bytes data) returns()
func (_CyberConnect *CyberConnectTransactorSession) SetEssenceData(profileId *big.Int, essenceId *big.Int, uri string, mw common.Address, data []byte) (*types.Transaction, error) {
	return _CyberConnect.Contract.SetEssenceData(&_CyberConnect.TransactOpts, profileId, essenceId, uri, mw, data)
}

// SetEssenceDataWithSig is a paid mutator transaction binding the contract method 0x5ea97079.
//
// Solidity: function setEssenceDataWithSig(uint256 profileId, uint256 essenceId, string uri, address mw, bytes data, (uint8,bytes32,bytes32,uint256) sig) returns()
func (_CyberConnect *CyberConnectTransactor) SetEssenceDataWithSig(opts *bind.TransactOpts, profileId *big.Int, essenceId *big.Int, uri string, mw common.Address, data []byte, sig DataTypesEIP712Signature) (*types.Transaction, error) {
	return _CyberConnect.contract.Transact(opts, "setEssenceDataWithSig", profileId, essenceId, uri, mw, data, sig)
}

// SetEssenceDataWithSig is a paid mutator transaction binding the contract method 0x5ea97079.
//
// Solidity: function setEssenceDataWithSig(uint256 profileId, uint256 essenceId, string uri, address mw, bytes data, (uint8,bytes32,bytes32,uint256) sig) returns()
func (_CyberConnect *CyberConnectSession) SetEssenceDataWithSig(profileId *big.Int, essenceId *big.Int, uri string, mw common.Address, data []byte, sig DataTypesEIP712Signature) (*types.Transaction, error) {
	return _CyberConnect.Contract.SetEssenceDataWithSig(&_CyberConnect.TransactOpts, profileId, essenceId, uri, mw, data, sig)
}

// SetEssenceDataWithSig is a paid mutator transaction binding the contract method 0x5ea97079.
//
// Solidity: function setEssenceDataWithSig(uint256 profileId, uint256 essenceId, string uri, address mw, bytes data, (uint8,bytes32,bytes32,uint256) sig) returns()
func (_CyberConnect *CyberConnectTransactorSession) SetEssenceDataWithSig(profileId *big.Int, essenceId *big.Int, uri string, mw common.Address, data []byte, sig DataTypesEIP712Signature) (*types.Transaction, error) {
	return _CyberConnect.Contract.SetEssenceDataWithSig(&_CyberConnect.TransactOpts, profileId, essenceId, uri, mw, data, sig)
}

// SetMetadata is a paid mutator transaction binding the contract method 0x593aa283.
//
// Solidity: function setMetadata(uint256 profileId, string metadata) returns()
func (_CyberConnect *CyberConnectTransactor) SetMetadata(opts *bind.TransactOpts, profileId *big.Int, metadata string) (*types.Transaction, error) {
	return _CyberConnect.contract.Transact(opts, "setMetadata", profileId, metadata)
}

// SetMetadata is a paid mutator transaction binding the contract method 0x593aa283.
//
// Solidity: function setMetadata(uint256 profileId, string metadata) returns()
func (_CyberConnect *CyberConnectSession) SetMetadata(profileId *big.Int, metadata string) (*types.Transaction, error) {
	return _CyberConnect.Contract.SetMetadata(&_CyberConnect.TransactOpts, profileId, metadata)
}

// SetMetadata is a paid mutator transaction binding the contract method 0x593aa283.
//
// Solidity: function setMetadata(uint256 profileId, string metadata) returns()
func (_CyberConnect *CyberConnectTransactorSession) SetMetadata(profileId *big.Int, metadata string) (*types.Transaction, error) {
	return _CyberConnect.Contract.SetMetadata(&_CyberConnect.TransactOpts, profileId, metadata)
}

// SetMetadataWithSig is a paid mutator transaction binding the contract method 0xd0e56c38.
//
// Solidity: function setMetadataWithSig(uint256 profileId, string metadata, (uint8,bytes32,bytes32,uint256) sig) returns()
func (_CyberConnect *CyberConnectTransactor) SetMetadataWithSig(opts *bind.TransactOpts, profileId *big.Int, metadata string, sig DataTypesEIP712Signature) (*types.Transaction, error) {
	return _CyberConnect.contract.Transact(opts, "setMetadataWithSig", profileId, metadata, sig)
}

// SetMetadataWithSig is a paid mutator transaction binding the contract method 0xd0e56c38.
//
// Solidity: function setMetadataWithSig(uint256 profileId, string metadata, (uint8,bytes32,bytes32,uint256) sig) returns()
func (_CyberConnect *CyberConnectSession) SetMetadataWithSig(profileId *big.Int, metadata string, sig DataTypesEIP712Signature) (*types.Transaction, error) {
	return _CyberConnect.Contract.SetMetadataWithSig(&_CyberConnect.TransactOpts, profileId, metadata, sig)
}

// SetMetadataWithSig is a paid mutator transaction binding the contract method 0xd0e56c38.
//
// Solidity: function setMetadataWithSig(uint256 profileId, string metadata, (uint8,bytes32,bytes32,uint256) sig) returns()
func (_CyberConnect *CyberConnectTransactorSession) SetMetadataWithSig(profileId *big.Int, metadata string, sig DataTypesEIP712Signature) (*types.Transaction, error) {
	return _CyberConnect.Contract.SetMetadataWithSig(&_CyberConnect.TransactOpts, profileId, metadata, sig)
}

// SetNFTDescriptor is a paid mutator transaction binding the contract method 0x7cad6cd1.
//
// Solidity: function setNFTDescriptor(address descriptor) returns()
func (_CyberConnect *CyberConnectTransactor) SetNFTDescriptor(opts *bind.TransactOpts, descriptor common.Address) (*types.Transaction, error) {
	return _CyberConnect.contract.Transact(opts, "setNFTDescriptor", descriptor)
}

// SetNFTDescriptor is a paid mutator transaction binding the contract method 0x7cad6cd1.
//
// Solidity: function setNFTDescriptor(address descriptor) returns()
func (_CyberConnect *CyberConnectSession) SetNFTDescriptor(descriptor common.Address) (*types.Transaction, error) {
	return _CyberConnect.Contract.SetNFTDescriptor(&_CyberConnect.TransactOpts, descriptor)
}

// SetNFTDescriptor is a paid mutator transaction binding the contract method 0x7cad6cd1.
//
// Solidity: function setNFTDescriptor(address descriptor) returns()
func (_CyberConnect *CyberConnectTransactorSession) SetNFTDescriptor(descriptor common.Address) (*types.Transaction, error) {
	return _CyberConnect.Contract.SetNFTDescriptor(&_CyberConnect.TransactOpts, descriptor)
}

// SetNamespaceOwner is a paid mutator transaction binding the contract method 0xf7727c9f.
//
// Solidity: function setNamespaceOwner(address owner) returns()
func (_CyberConnect *CyberConnectTransactor) SetNamespaceOwner(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _CyberConnect.contract.Transact(opts, "setNamespaceOwner", owner)
}

// SetNamespaceOwner is a paid mutator transaction binding the contract method 0xf7727c9f.
//
// Solidity: function setNamespaceOwner(address owner) returns()
func (_CyberConnect *CyberConnectSession) SetNamespaceOwner(owner common.Address) (*types.Transaction, error) {
	return _CyberConnect.Contract.SetNamespaceOwner(&_CyberConnect.TransactOpts, owner)
}

// SetNamespaceOwner is a paid mutator transaction binding the contract method 0xf7727c9f.
//
// Solidity: function setNamespaceOwner(address owner) returns()
func (_CyberConnect *CyberConnectTransactorSession) SetNamespaceOwner(owner common.Address) (*types.Transaction, error) {
	return _CyberConnect.Contract.SetNamespaceOwner(&_CyberConnect.TransactOpts, owner)
}

// SetOperatorApproval is a paid mutator transaction binding the contract method 0xafa83e68.
//
// Solidity: function setOperatorApproval(uint256 profileId, address operator, bool approved) returns()
func (_CyberConnect *CyberConnectTransactor) SetOperatorApproval(opts *bind.TransactOpts, profileId *big.Int, operator common.Address, approved bool) (*types.Transaction, error) {
	return _CyberConnect.contract.Transact(opts, "setOperatorApproval", profileId, operator, approved)
}

// SetOperatorApproval is a paid mutator transaction binding the contract method 0xafa83e68.
//
// Solidity: function setOperatorApproval(uint256 profileId, address operator, bool approved) returns()
func (_CyberConnect *CyberConnectSession) SetOperatorApproval(profileId *big.Int, operator common.Address, approved bool) (*types.Transaction, error) {
	return _CyberConnect.Contract.SetOperatorApproval(&_CyberConnect.TransactOpts, profileId, operator, approved)
}

// SetOperatorApproval is a paid mutator transaction binding the contract method 0xafa83e68.
//
// Solidity: function setOperatorApproval(uint256 profileId, address operator, bool approved) returns()
func (_CyberConnect *CyberConnectTransactorSession) SetOperatorApproval(profileId *big.Int, operator common.Address, approved bool) (*types.Transaction, error) {
	return _CyberConnect.Contract.SetOperatorApproval(&_CyberConnect.TransactOpts, profileId, operator, approved)
}

// SetPrimaryProfile is a paid mutator transaction binding the contract method 0x7c08d9fd.
//
// Solidity: function setPrimaryProfile(uint256 profileId) returns()
func (_CyberConnect *CyberConnectTransactor) SetPrimaryProfile(opts *bind.TransactOpts, profileId *big.Int) (*types.Transaction, error) {
	return _CyberConnect.contract.Transact(opts, "setPrimaryProfile", profileId)
}

// SetPrimaryProfile is a paid mutator transaction binding the contract method 0x7c08d9fd.
//
// Solidity: function setPrimaryProfile(uint256 profileId) returns()
func (_CyberConnect *CyberConnectSession) SetPrimaryProfile(profileId *big.Int) (*types.Transaction, error) {
	return _CyberConnect.Contract.SetPrimaryProfile(&_CyberConnect.TransactOpts, profileId)
}

// SetPrimaryProfile is a paid mutator transaction binding the contract method 0x7c08d9fd.
//
// Solidity: function setPrimaryProfile(uint256 profileId) returns()
func (_CyberConnect *CyberConnectTransactorSession) SetPrimaryProfile(profileId *big.Int) (*types.Transaction, error) {
	return _CyberConnect.Contract.SetPrimaryProfile(&_CyberConnect.TransactOpts, profileId)
}

// SetPrimaryProfileWithSig is a paid mutator transaction binding the contract method 0xa95c1cbf.
//
// Solidity: function setPrimaryProfileWithSig(uint256 profileId, (uint8,bytes32,bytes32,uint256) sig) returns()
func (_CyberConnect *CyberConnectTransactor) SetPrimaryProfileWithSig(opts *bind.TransactOpts, profileId *big.Int, sig DataTypesEIP712Signature) (*types.Transaction, error) {
	return _CyberConnect.contract.Transact(opts, "setPrimaryProfileWithSig", profileId, sig)
}

// SetPrimaryProfileWithSig is a paid mutator transaction binding the contract method 0xa95c1cbf.
//
// Solidity: function setPrimaryProfileWithSig(uint256 profileId, (uint8,bytes32,bytes32,uint256) sig) returns()
func (_CyberConnect *CyberConnectSession) SetPrimaryProfileWithSig(profileId *big.Int, sig DataTypesEIP712Signature) (*types.Transaction, error) {
	return _CyberConnect.Contract.SetPrimaryProfileWithSig(&_CyberConnect.TransactOpts, profileId, sig)
}

// SetPrimaryProfileWithSig is a paid mutator transaction binding the contract method 0xa95c1cbf.
//
// Solidity: function setPrimaryProfileWithSig(uint256 profileId, (uint8,bytes32,bytes32,uint256) sig) returns()
func (_CyberConnect *CyberConnectTransactorSession) SetPrimaryProfileWithSig(profileId *big.Int, sig DataTypesEIP712Signature) (*types.Transaction, error) {
	return _CyberConnect.Contract.SetPrimaryProfileWithSig(&_CyberConnect.TransactOpts, profileId, sig)
}

// SetSubscribeData is a paid mutator transaction binding the contract method 0x8d54db1b.
//
// Solidity: function setSubscribeData(uint256 profileId, string uri, address mw, bytes data) returns()
func (_CyberConnect *CyberConnectTransactor) SetSubscribeData(opts *bind.TransactOpts, profileId *big.Int, uri string, mw common.Address, data []byte) (*types.Transaction, error) {
	return _CyberConnect.contract.Transact(opts, "setSubscribeData", profileId, uri, mw, data)
}

// SetSubscribeData is a paid mutator transaction binding the contract method 0x8d54db1b.
//
// Solidity: function setSubscribeData(uint256 profileId, string uri, address mw, bytes data) returns()
func (_CyberConnect *CyberConnectSession) SetSubscribeData(profileId *big.Int, uri string, mw common.Address, data []byte) (*types.Transaction, error) {
	return _CyberConnect.Contract.SetSubscribeData(&_CyberConnect.TransactOpts, profileId, uri, mw, data)
}

// SetSubscribeData is a paid mutator transaction binding the contract method 0x8d54db1b.
//
// Solidity: function setSubscribeData(uint256 profileId, string uri, address mw, bytes data) returns()
func (_CyberConnect *CyberConnectTransactorSession) SetSubscribeData(profileId *big.Int, uri string, mw common.Address, data []byte) (*types.Transaction, error) {
	return _CyberConnect.Contract.SetSubscribeData(&_CyberConnect.TransactOpts, profileId, uri, mw, data)
}

// SetSubscribeDataWithSig is a paid mutator transaction binding the contract method 0x2229f28d.
//
// Solidity: function setSubscribeDataWithSig(uint256 profileId, string uri, address mw, bytes data, (uint8,bytes32,bytes32,uint256) sig) returns()
func (_CyberConnect *CyberConnectTransactor) SetSubscribeDataWithSig(opts *bind.TransactOpts, profileId *big.Int, uri string, mw common.Address, data []byte, sig DataTypesEIP712Signature) (*types.Transaction, error) {
	return _CyberConnect.contract.Transact(opts, "setSubscribeDataWithSig", profileId, uri, mw, data, sig)
}

// SetSubscribeDataWithSig is a paid mutator transaction binding the contract method 0x2229f28d.
//
// Solidity: function setSubscribeDataWithSig(uint256 profileId, string uri, address mw, bytes data, (uint8,bytes32,bytes32,uint256) sig) returns()
func (_CyberConnect *CyberConnectSession) SetSubscribeDataWithSig(profileId *big.Int, uri string, mw common.Address, data []byte, sig DataTypesEIP712Signature) (*types.Transaction, error) {
	return _CyberConnect.Contract.SetSubscribeDataWithSig(&_CyberConnect.TransactOpts, profileId, uri, mw, data, sig)
}

// SetSubscribeDataWithSig is a paid mutator transaction binding the contract method 0x2229f28d.
//
// Solidity: function setSubscribeDataWithSig(uint256 profileId, string uri, address mw, bytes data, (uint8,bytes32,bytes32,uint256) sig) returns()
func (_CyberConnect *CyberConnectTransactorSession) SetSubscribeDataWithSig(profileId *big.Int, uri string, mw common.Address, data []byte, sig DataTypesEIP712Signature) (*types.Transaction, error) {
	return _CyberConnect.Contract.SetSubscribeDataWithSig(&_CyberConnect.TransactOpts, profileId, uri, mw, data, sig)
}

// Subscribe is a paid mutator transaction binding the contract method 0x4940e2d1.
//
// Solidity: function subscribe((uint256[]) params, bytes[] preDatas, bytes[] postDatas) returns(uint256[])
func (_CyberConnect *CyberConnectTransactor) Subscribe(opts *bind.TransactOpts, params DataTypesSubscribeParams, preDatas [][]byte, postDatas [][]byte) (*types.Transaction, error) {
	return _CyberConnect.contract.Transact(opts, "subscribe", params, preDatas, postDatas)
}

// Subscribe is a paid mutator transaction binding the contract method 0x4940e2d1.
//
// Solidity: function subscribe((uint256[]) params, bytes[] preDatas, bytes[] postDatas) returns(uint256[])
func (_CyberConnect *CyberConnectSession) Subscribe(params DataTypesSubscribeParams, preDatas [][]byte, postDatas [][]byte) (*types.Transaction, error) {
	return _CyberConnect.Contract.Subscribe(&_CyberConnect.TransactOpts, params, preDatas, postDatas)
}

// Subscribe is a paid mutator transaction binding the contract method 0x4940e2d1.
//
// Solidity: function subscribe((uint256[]) params, bytes[] preDatas, bytes[] postDatas) returns(uint256[])
func (_CyberConnect *CyberConnectTransactorSession) Subscribe(params DataTypesSubscribeParams, preDatas [][]byte, postDatas [][]byte) (*types.Transaction, error) {
	return _CyberConnect.Contract.Subscribe(&_CyberConnect.TransactOpts, params, preDatas, postDatas)
}

// SubscribeWithSig is a paid mutator transaction binding the contract method 0x026103e9.
//
// Solidity: function subscribeWithSig((uint256[]) params, bytes[] preDatas, bytes[] postDatas, address sender, (uint8,bytes32,bytes32,uint256) sig) returns(uint256[])
func (_CyberConnect *CyberConnectTransactor) SubscribeWithSig(opts *bind.TransactOpts, params DataTypesSubscribeParams, preDatas [][]byte, postDatas [][]byte, sender common.Address, sig DataTypesEIP712Signature) (*types.Transaction, error) {
	return _CyberConnect.contract.Transact(opts, "subscribeWithSig", params, preDatas, postDatas, sender, sig)
}

// SubscribeWithSig is a paid mutator transaction binding the contract method 0x026103e9.
//
// Solidity: function subscribeWithSig((uint256[]) params, bytes[] preDatas, bytes[] postDatas, address sender, (uint8,bytes32,bytes32,uint256) sig) returns(uint256[])
func (_CyberConnect *CyberConnectSession) SubscribeWithSig(params DataTypesSubscribeParams, preDatas [][]byte, postDatas [][]byte, sender common.Address, sig DataTypesEIP712Signature) (*types.Transaction, error) {
	return _CyberConnect.Contract.SubscribeWithSig(&_CyberConnect.TransactOpts, params, preDatas, postDatas, sender, sig)
}

// SubscribeWithSig is a paid mutator transaction binding the contract method 0x026103e9.
//
// Solidity: function subscribeWithSig((uint256[]) params, bytes[] preDatas, bytes[] postDatas, address sender, (uint8,bytes32,bytes32,uint256) sig) returns(uint256[])
func (_CyberConnect *CyberConnectTransactorSession) SubscribeWithSig(params DataTypesSubscribeParams, preDatas [][]byte, postDatas [][]byte, sender common.Address, sig DataTypesEIP712Signature) (*types.Transaction, error) {
	return _CyberConnect.Contract.SubscribeWithSig(&_CyberConnect.TransactOpts, params, preDatas, postDatas, sender, sig)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 id) returns()
func (_CyberConnect *CyberConnectTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _CyberConnect.contract.Transact(opts, "transferFrom", from, to, id)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 id) returns()
func (_CyberConnect *CyberConnectSession) TransferFrom(from common.Address, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _CyberConnect.Contract.TransferFrom(&_CyberConnect.TransactOpts, from, to, id)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 id) returns()
func (_CyberConnect *CyberConnectTransactorSession) TransferFrom(from common.Address, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _CyberConnect.Contract.TransferFrom(&_CyberConnect.TransactOpts, from, to, id)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_CyberConnect *CyberConnectTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _CyberConnect.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_CyberConnect *CyberConnectSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _CyberConnect.Contract.UpgradeTo(&_CyberConnect.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_CyberConnect *CyberConnectTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _CyberConnect.Contract.UpgradeTo(&_CyberConnect.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_CyberConnect *CyberConnectTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _CyberConnect.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_CyberConnect *CyberConnectSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _CyberConnect.Contract.UpgradeToAndCall(&_CyberConnect.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_CyberConnect *CyberConnectTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _CyberConnect.Contract.UpgradeToAndCall(&_CyberConnect.TransactOpts, newImplementation, data)
}

// CyberConnectAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the CyberConnect contract.
type CyberConnectAdminChangedIterator struct {
	Event *CyberConnectAdminChanged // Event containing the contract specifics and raw log

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
func (it *CyberConnectAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CyberConnectAdminChanged)
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
		it.Event = new(CyberConnectAdminChanged)
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
func (it *CyberConnectAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CyberConnectAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CyberConnectAdminChanged represents a AdminChanged event raised by the CyberConnect contract.
type CyberConnectAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_CyberConnect *CyberConnectFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*CyberConnectAdminChangedIterator, error) {

	logs, sub, err := _CyberConnect.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &CyberConnectAdminChangedIterator{contract: _CyberConnect.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_CyberConnect *CyberConnectFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *CyberConnectAdminChanged) (event.Subscription, error) {

	logs, sub, err := _CyberConnect.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CyberConnectAdminChanged)
				if err := _CyberConnect.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_CyberConnect *CyberConnectFilterer) ParseAdminChanged(log types.Log) (*CyberConnectAdminChanged, error) {
	event := new(CyberConnectAdminChanged)
	if err := _CyberConnect.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CyberConnectApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the CyberConnect contract.
type CyberConnectApprovalIterator struct {
	Event *CyberConnectApproval // Event containing the contract specifics and raw log

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
func (it *CyberConnectApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CyberConnectApproval)
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
		it.Event = new(CyberConnectApproval)
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
func (it *CyberConnectApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CyberConnectApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CyberConnectApproval represents a Approval event raised by the CyberConnect contract.
type CyberConnectApproval struct {
	Owner   common.Address
	Spender common.Address
	Id      *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 indexed id)
func (_CyberConnect *CyberConnectFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address, id []*big.Int) (*CyberConnectApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _CyberConnect.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule, idRule)
	if err != nil {
		return nil, err
	}
	return &CyberConnectApprovalIterator{contract: _CyberConnect.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 indexed id)
func (_CyberConnect *CyberConnectFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CyberConnectApproval, owner []common.Address, spender []common.Address, id []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _CyberConnect.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CyberConnectApproval)
				if err := _CyberConnect.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 indexed id)
func (_CyberConnect *CyberConnectFilterer) ParseApproval(log types.Log) (*CyberConnectApproval, error) {
	event := new(CyberConnectApproval)
	if err := _CyberConnect.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CyberConnectApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the CyberConnect contract.
type CyberConnectApprovalForAllIterator struct {
	Event *CyberConnectApprovalForAll // Event containing the contract specifics and raw log

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
func (it *CyberConnectApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CyberConnectApprovalForAll)
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
		it.Event = new(CyberConnectApprovalForAll)
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
func (it *CyberConnectApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CyberConnectApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CyberConnectApprovalForAll represents a ApprovalForAll event raised by the CyberConnect contract.
type CyberConnectApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_CyberConnect *CyberConnectFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*CyberConnectApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _CyberConnect.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &CyberConnectApprovalForAllIterator{contract: _CyberConnect.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_CyberConnect *CyberConnectFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *CyberConnectApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _CyberConnect.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CyberConnectApprovalForAll)
				if err := _CyberConnect.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_CyberConnect *CyberConnectFilterer) ParseApprovalForAll(log types.Log) (*CyberConnectApprovalForAll, error) {
	event := new(CyberConnectApprovalForAll)
	if err := _CyberConnect.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CyberConnectBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the CyberConnect contract.
type CyberConnectBeaconUpgradedIterator struct {
	Event *CyberConnectBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *CyberConnectBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CyberConnectBeaconUpgraded)
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
		it.Event = new(CyberConnectBeaconUpgraded)
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
func (it *CyberConnectBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CyberConnectBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CyberConnectBeaconUpgraded represents a BeaconUpgraded event raised by the CyberConnect contract.
type CyberConnectBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_CyberConnect *CyberConnectFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*CyberConnectBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _CyberConnect.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &CyberConnectBeaconUpgradedIterator{contract: _CyberConnect.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_CyberConnect *CyberConnectFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *CyberConnectBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _CyberConnect.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CyberConnectBeaconUpgraded)
				if err := _CyberConnect.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_CyberConnect *CyberConnectFilterer) ParseBeaconUpgraded(log types.Log) (*CyberConnectBeaconUpgraded, error) {
	event := new(CyberConnectBeaconUpgraded)
	if err := _CyberConnect.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CyberConnectCollectEssenceIterator is returned from FilterCollectEssence and is used to iterate over the raw logs and unpacked data for CollectEssence events raised by the CyberConnect contract.
type CyberConnectCollectEssenceIterator struct {
	Event *CyberConnectCollectEssence // Event containing the contract specifics and raw log

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
func (it *CyberConnectCollectEssenceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CyberConnectCollectEssence)
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
		it.Event = new(CyberConnectCollectEssence)
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
func (it *CyberConnectCollectEssenceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CyberConnectCollectEssenceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CyberConnectCollectEssence represents a CollectEssence event raised by the CyberConnect contract.
type CyberConnectCollectEssence struct {
	Collector common.Address
	ProfileId *big.Int
	EssenceId *big.Int
	TokenId   *big.Int
	PreData   []byte
	PostData  []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCollectEssence is a free log retrieval operation binding the contract event 0xd4665c34529353804143794023e073d96230b57acf608ae514de33d10a09354f.
//
// Solidity: event CollectEssence(address indexed collector, uint256 indexed profileId, uint256 indexed essenceId, uint256 tokenId, bytes preData, bytes postData)
func (_CyberConnect *CyberConnectFilterer) FilterCollectEssence(opts *bind.FilterOpts, collector []common.Address, profileId []*big.Int, essenceId []*big.Int) (*CyberConnectCollectEssenceIterator, error) {

	var collectorRule []interface{}
	for _, collectorItem := range collector {
		collectorRule = append(collectorRule, collectorItem)
	}
	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var essenceIdRule []interface{}
	for _, essenceIdItem := range essenceId {
		essenceIdRule = append(essenceIdRule, essenceIdItem)
	}

	logs, sub, err := _CyberConnect.contract.FilterLogs(opts, "CollectEssence", collectorRule, profileIdRule, essenceIdRule)
	if err != nil {
		return nil, err
	}
	return &CyberConnectCollectEssenceIterator{contract: _CyberConnect.contract, event: "CollectEssence", logs: logs, sub: sub}, nil
}

// WatchCollectEssence is a free log subscription operation binding the contract event 0xd4665c34529353804143794023e073d96230b57acf608ae514de33d10a09354f.
//
// Solidity: event CollectEssence(address indexed collector, uint256 indexed profileId, uint256 indexed essenceId, uint256 tokenId, bytes preData, bytes postData)
func (_CyberConnect *CyberConnectFilterer) WatchCollectEssence(opts *bind.WatchOpts, sink chan<- *CyberConnectCollectEssence, collector []common.Address, profileId []*big.Int, essenceId []*big.Int) (event.Subscription, error) {

	var collectorRule []interface{}
	for _, collectorItem := range collector {
		collectorRule = append(collectorRule, collectorItem)
	}
	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var essenceIdRule []interface{}
	for _, essenceIdItem := range essenceId {
		essenceIdRule = append(essenceIdRule, essenceIdItem)
	}

	logs, sub, err := _CyberConnect.contract.WatchLogs(opts, "CollectEssence", collectorRule, profileIdRule, essenceIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CyberConnectCollectEssence)
				if err := _CyberConnect.contract.UnpackLog(event, "CollectEssence", log); err != nil {
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

// ParseCollectEssence is a log parse operation binding the contract event 0xd4665c34529353804143794023e073d96230b57acf608ae514de33d10a09354f.
//
// Solidity: event CollectEssence(address indexed collector, uint256 indexed profileId, uint256 indexed essenceId, uint256 tokenId, bytes preData, bytes postData)
func (_CyberConnect *CyberConnectFilterer) ParseCollectEssence(log types.Log) (*CyberConnectCollectEssence, error) {
	event := new(CyberConnectCollectEssence)
	if err := _CyberConnect.contract.UnpackLog(event, "CollectEssence", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CyberConnectCreateProfileIterator is returned from FilterCreateProfile and is used to iterate over the raw logs and unpacked data for CreateProfile events raised by the CyberConnect contract.
type CyberConnectCreateProfileIterator struct {
	Event *CyberConnectCreateProfile // Event containing the contract specifics and raw log

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
func (it *CyberConnectCreateProfileIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CyberConnectCreateProfile)
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
		it.Event = new(CyberConnectCreateProfile)
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
func (it *CyberConnectCreateProfileIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CyberConnectCreateProfileIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CyberConnectCreateProfile represents a CreateProfile event raised by the CyberConnect contract.
type CyberConnectCreateProfile struct {
	To        common.Address
	ProfileId *big.Int
	Handle    string
	Avatar    string
	Metadata  string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCreateProfile is a free log retrieval operation binding the contract event 0x0d0246441d2185882e392b0a85e33212f2e1f668cf78c11c1808d421b2549fa6.
//
// Solidity: event CreateProfile(address indexed to, uint256 indexed profileId, string handle, string avatar, string metadata)
func (_CyberConnect *CyberConnectFilterer) FilterCreateProfile(opts *bind.FilterOpts, to []common.Address, profileId []*big.Int) (*CyberConnectCreateProfileIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _CyberConnect.contract.FilterLogs(opts, "CreateProfile", toRule, profileIdRule)
	if err != nil {
		return nil, err
	}
	return &CyberConnectCreateProfileIterator{contract: _CyberConnect.contract, event: "CreateProfile", logs: logs, sub: sub}, nil
}

// WatchCreateProfile is a free log subscription operation binding the contract event 0x0d0246441d2185882e392b0a85e33212f2e1f668cf78c11c1808d421b2549fa6.
//
// Solidity: event CreateProfile(address indexed to, uint256 indexed profileId, string handle, string avatar, string metadata)
func (_CyberConnect *CyberConnectFilterer) WatchCreateProfile(opts *bind.WatchOpts, sink chan<- *CyberConnectCreateProfile, to []common.Address, profileId []*big.Int) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _CyberConnect.contract.WatchLogs(opts, "CreateProfile", toRule, profileIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CyberConnectCreateProfile)
				if err := _CyberConnect.contract.UnpackLog(event, "CreateProfile", log); err != nil {
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

// ParseCreateProfile is a log parse operation binding the contract event 0x0d0246441d2185882e392b0a85e33212f2e1f668cf78c11c1808d421b2549fa6.
//
// Solidity: event CreateProfile(address indexed to, uint256 indexed profileId, string handle, string avatar, string metadata)
func (_CyberConnect *CyberConnectFilterer) ParseCreateProfile(log types.Log) (*CyberConnectCreateProfile, error) {
	event := new(CyberConnectCreateProfile)
	if err := _CyberConnect.contract.UnpackLog(event, "CreateProfile", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CyberConnectDeployEssenceNFTIterator is returned from FilterDeployEssenceNFT and is used to iterate over the raw logs and unpacked data for DeployEssenceNFT events raised by the CyberConnect contract.
type CyberConnectDeployEssenceNFTIterator struct {
	Event *CyberConnectDeployEssenceNFT // Event containing the contract specifics and raw log

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
func (it *CyberConnectDeployEssenceNFTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CyberConnectDeployEssenceNFT)
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
		it.Event = new(CyberConnectDeployEssenceNFT)
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
func (it *CyberConnectDeployEssenceNFTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CyberConnectDeployEssenceNFTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CyberConnectDeployEssenceNFT represents a DeployEssenceNFT event raised by the CyberConnect contract.
type CyberConnectDeployEssenceNFT struct {
	ProfileId  *big.Int
	EssenceId  *big.Int
	EssenceNFT common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDeployEssenceNFT is a free log retrieval operation binding the contract event 0x9582a6b88daf141f38fa06c06d493aa8a09546d1b508289c833a9d48498ff189.
//
// Solidity: event DeployEssenceNFT(uint256 indexed profileId, uint256 indexed essenceId, address indexed essenceNFT)
func (_CyberConnect *CyberConnectFilterer) FilterDeployEssenceNFT(opts *bind.FilterOpts, profileId []*big.Int, essenceId []*big.Int, essenceNFT []common.Address) (*CyberConnectDeployEssenceNFTIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var essenceIdRule []interface{}
	for _, essenceIdItem := range essenceId {
		essenceIdRule = append(essenceIdRule, essenceIdItem)
	}
	var essenceNFTRule []interface{}
	for _, essenceNFTItem := range essenceNFT {
		essenceNFTRule = append(essenceNFTRule, essenceNFTItem)
	}

	logs, sub, err := _CyberConnect.contract.FilterLogs(opts, "DeployEssenceNFT", profileIdRule, essenceIdRule, essenceNFTRule)
	if err != nil {
		return nil, err
	}
	return &CyberConnectDeployEssenceNFTIterator{contract: _CyberConnect.contract, event: "DeployEssenceNFT", logs: logs, sub: sub}, nil
}

// WatchDeployEssenceNFT is a free log subscription operation binding the contract event 0x9582a6b88daf141f38fa06c06d493aa8a09546d1b508289c833a9d48498ff189.
//
// Solidity: event DeployEssenceNFT(uint256 indexed profileId, uint256 indexed essenceId, address indexed essenceNFT)
func (_CyberConnect *CyberConnectFilterer) WatchDeployEssenceNFT(opts *bind.WatchOpts, sink chan<- *CyberConnectDeployEssenceNFT, profileId []*big.Int, essenceId []*big.Int, essenceNFT []common.Address) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var essenceIdRule []interface{}
	for _, essenceIdItem := range essenceId {
		essenceIdRule = append(essenceIdRule, essenceIdItem)
	}
	var essenceNFTRule []interface{}
	for _, essenceNFTItem := range essenceNFT {
		essenceNFTRule = append(essenceNFTRule, essenceNFTItem)
	}

	logs, sub, err := _CyberConnect.contract.WatchLogs(opts, "DeployEssenceNFT", profileIdRule, essenceIdRule, essenceNFTRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CyberConnectDeployEssenceNFT)
				if err := _CyberConnect.contract.UnpackLog(event, "DeployEssenceNFT", log); err != nil {
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

// ParseDeployEssenceNFT is a log parse operation binding the contract event 0x9582a6b88daf141f38fa06c06d493aa8a09546d1b508289c833a9d48498ff189.
//
// Solidity: event DeployEssenceNFT(uint256 indexed profileId, uint256 indexed essenceId, address indexed essenceNFT)
func (_CyberConnect *CyberConnectFilterer) ParseDeployEssenceNFT(log types.Log) (*CyberConnectDeployEssenceNFT, error) {
	event := new(CyberConnectDeployEssenceNFT)
	if err := _CyberConnect.contract.UnpackLog(event, "DeployEssenceNFT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CyberConnectDeploySubscribeNFTIterator is returned from FilterDeploySubscribeNFT and is used to iterate over the raw logs and unpacked data for DeploySubscribeNFT events raised by the CyberConnect contract.
type CyberConnectDeploySubscribeNFTIterator struct {
	Event *CyberConnectDeploySubscribeNFT // Event containing the contract specifics and raw log

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
func (it *CyberConnectDeploySubscribeNFTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CyberConnectDeploySubscribeNFT)
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
		it.Event = new(CyberConnectDeploySubscribeNFT)
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
func (it *CyberConnectDeploySubscribeNFTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CyberConnectDeploySubscribeNFTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CyberConnectDeploySubscribeNFT represents a DeploySubscribeNFT event raised by the CyberConnect contract.
type CyberConnectDeploySubscribeNFT struct {
	ProfileId    *big.Int
	SubscribeNFT common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDeploySubscribeNFT is a free log retrieval operation binding the contract event 0x8fb25da3d7c3a4880851ad04d6fca5a6f663ca562d932462e1444544321de175.
//
// Solidity: event DeploySubscribeNFT(uint256 indexed profileId, address indexed subscribeNFT)
func (_CyberConnect *CyberConnectFilterer) FilterDeploySubscribeNFT(opts *bind.FilterOpts, profileId []*big.Int, subscribeNFT []common.Address) (*CyberConnectDeploySubscribeNFTIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var subscribeNFTRule []interface{}
	for _, subscribeNFTItem := range subscribeNFT {
		subscribeNFTRule = append(subscribeNFTRule, subscribeNFTItem)
	}

	logs, sub, err := _CyberConnect.contract.FilterLogs(opts, "DeploySubscribeNFT", profileIdRule, subscribeNFTRule)
	if err != nil {
		return nil, err
	}
	return &CyberConnectDeploySubscribeNFTIterator{contract: _CyberConnect.contract, event: "DeploySubscribeNFT", logs: logs, sub: sub}, nil
}

// WatchDeploySubscribeNFT is a free log subscription operation binding the contract event 0x8fb25da3d7c3a4880851ad04d6fca5a6f663ca562d932462e1444544321de175.
//
// Solidity: event DeploySubscribeNFT(uint256 indexed profileId, address indexed subscribeNFT)
func (_CyberConnect *CyberConnectFilterer) WatchDeploySubscribeNFT(opts *bind.WatchOpts, sink chan<- *CyberConnectDeploySubscribeNFT, profileId []*big.Int, subscribeNFT []common.Address) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var subscribeNFTRule []interface{}
	for _, subscribeNFTItem := range subscribeNFT {
		subscribeNFTRule = append(subscribeNFTRule, subscribeNFTItem)
	}

	logs, sub, err := _CyberConnect.contract.WatchLogs(opts, "DeploySubscribeNFT", profileIdRule, subscribeNFTRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CyberConnectDeploySubscribeNFT)
				if err := _CyberConnect.contract.UnpackLog(event, "DeploySubscribeNFT", log); err != nil {
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

// ParseDeploySubscribeNFT is a log parse operation binding the contract event 0x8fb25da3d7c3a4880851ad04d6fca5a6f663ca562d932462e1444544321de175.
//
// Solidity: event DeploySubscribeNFT(uint256 indexed profileId, address indexed subscribeNFT)
func (_CyberConnect *CyberConnectFilterer) ParseDeploySubscribeNFT(log types.Log) (*CyberConnectDeploySubscribeNFT, error) {
	event := new(CyberConnectDeploySubscribeNFT)
	if err := _CyberConnect.contract.UnpackLog(event, "DeploySubscribeNFT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CyberConnectInitializeIterator is returned from FilterInitialize and is used to iterate over the raw logs and unpacked data for Initialize events raised by the CyberConnect contract.
type CyberConnectInitializeIterator struct {
	Event *CyberConnectInitialize // Event containing the contract specifics and raw log

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
func (it *CyberConnectInitializeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CyberConnectInitialize)
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
		it.Event = new(CyberConnectInitialize)
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
func (it *CyberConnectInitializeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CyberConnectInitializeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CyberConnectInitialize represents a Initialize event raised by the CyberConnect contract.
type CyberConnectInitialize struct {
	Owner  common.Address
	Name   string
	Symbol string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterInitialize is a free log retrieval operation binding the contract event 0x5b98d40fc5d761472c5e9e4eeda87b2e7d3d58f72ed64216d9350fe5bc1b5bf1.
//
// Solidity: event Initialize(address indexed owner, string name, string symbol)
func (_CyberConnect *CyberConnectFilterer) FilterInitialize(opts *bind.FilterOpts, owner []common.Address) (*CyberConnectInitializeIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _CyberConnect.contract.FilterLogs(opts, "Initialize", ownerRule)
	if err != nil {
		return nil, err
	}
	return &CyberConnectInitializeIterator{contract: _CyberConnect.contract, event: "Initialize", logs: logs, sub: sub}, nil
}

// WatchInitialize is a free log subscription operation binding the contract event 0x5b98d40fc5d761472c5e9e4eeda87b2e7d3d58f72ed64216d9350fe5bc1b5bf1.
//
// Solidity: event Initialize(address indexed owner, string name, string symbol)
func (_CyberConnect *CyberConnectFilterer) WatchInitialize(opts *bind.WatchOpts, sink chan<- *CyberConnectInitialize, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _CyberConnect.contract.WatchLogs(opts, "Initialize", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CyberConnectInitialize)
				if err := _CyberConnect.contract.UnpackLog(event, "Initialize", log); err != nil {
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

// ParseInitialize is a log parse operation binding the contract event 0x5b98d40fc5d761472c5e9e4eeda87b2e7d3d58f72ed64216d9350fe5bc1b5bf1.
//
// Solidity: event Initialize(address indexed owner, string name, string symbol)
func (_CyberConnect *CyberConnectFilterer) ParseInitialize(log types.Log) (*CyberConnectInitialize, error) {
	event := new(CyberConnectInitialize)
	if err := _CyberConnect.contract.UnpackLog(event, "Initialize", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CyberConnectInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the CyberConnect contract.
type CyberConnectInitializedIterator struct {
	Event *CyberConnectInitialized // Event containing the contract specifics and raw log

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
func (it *CyberConnectInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CyberConnectInitialized)
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
		it.Event = new(CyberConnectInitialized)
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
func (it *CyberConnectInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CyberConnectInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CyberConnectInitialized represents a Initialized event raised by the CyberConnect contract.
type CyberConnectInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_CyberConnect *CyberConnectFilterer) FilterInitialized(opts *bind.FilterOpts) (*CyberConnectInitializedIterator, error) {

	logs, sub, err := _CyberConnect.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &CyberConnectInitializedIterator{contract: _CyberConnect.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_CyberConnect *CyberConnectFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *CyberConnectInitialized) (event.Subscription, error) {

	logs, sub, err := _CyberConnect.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CyberConnectInitialized)
				if err := _CyberConnect.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_CyberConnect *CyberConnectFilterer) ParseInitialized(log types.Log) (*CyberConnectInitialized, error) {
	event := new(CyberConnectInitialized)
	if err := _CyberConnect.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CyberConnectPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the CyberConnect contract.
type CyberConnectPausedIterator struct {
	Event *CyberConnectPaused // Event containing the contract specifics and raw log

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
func (it *CyberConnectPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CyberConnectPaused)
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
		it.Event = new(CyberConnectPaused)
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
func (it *CyberConnectPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CyberConnectPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CyberConnectPaused represents a Paused event raised by the CyberConnect contract.
type CyberConnectPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_CyberConnect *CyberConnectFilterer) FilterPaused(opts *bind.FilterOpts) (*CyberConnectPausedIterator, error) {

	logs, sub, err := _CyberConnect.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &CyberConnectPausedIterator{contract: _CyberConnect.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_CyberConnect *CyberConnectFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *CyberConnectPaused) (event.Subscription, error) {

	logs, sub, err := _CyberConnect.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CyberConnectPaused)
				if err := _CyberConnect.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_CyberConnect *CyberConnectFilterer) ParsePaused(log types.Log) (*CyberConnectPaused, error) {
	event := new(CyberConnectPaused)
	if err := _CyberConnect.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CyberConnectRegisterEssenceIterator is returned from FilterRegisterEssence and is used to iterate over the raw logs and unpacked data for RegisterEssence events raised by the CyberConnect contract.
type CyberConnectRegisterEssenceIterator struct {
	Event *CyberConnectRegisterEssence // Event containing the contract specifics and raw log

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
func (it *CyberConnectRegisterEssenceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CyberConnectRegisterEssence)
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
		it.Event = new(CyberConnectRegisterEssence)
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
func (it *CyberConnectRegisterEssenceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CyberConnectRegisterEssenceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CyberConnectRegisterEssence represents a RegisterEssence event raised by the CyberConnect contract.
type CyberConnectRegisterEssence struct {
	ProfileId         *big.Int
	EssenceId         *big.Int
	Name              string
	Symbol            string
	EssenceTokenURI   string
	EssenceMw         common.Address
	PrepareReturnData []byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRegisterEssence is a free log retrieval operation binding the contract event 0x7f8f05cb8699d58aa2b69fa7a1b1ab0db87ab881493e58a6842cec13b2982011.
//
// Solidity: event RegisterEssence(uint256 indexed profileId, uint256 indexed essenceId, string name, string symbol, string essenceTokenURI, address essenceMw, bytes prepareReturnData)
func (_CyberConnect *CyberConnectFilterer) FilterRegisterEssence(opts *bind.FilterOpts, profileId []*big.Int, essenceId []*big.Int) (*CyberConnectRegisterEssenceIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var essenceIdRule []interface{}
	for _, essenceIdItem := range essenceId {
		essenceIdRule = append(essenceIdRule, essenceIdItem)
	}

	logs, sub, err := _CyberConnect.contract.FilterLogs(opts, "RegisterEssence", profileIdRule, essenceIdRule)
	if err != nil {
		return nil, err
	}
	return &CyberConnectRegisterEssenceIterator{contract: _CyberConnect.contract, event: "RegisterEssence", logs: logs, sub: sub}, nil
}

// WatchRegisterEssence is a free log subscription operation binding the contract event 0x7f8f05cb8699d58aa2b69fa7a1b1ab0db87ab881493e58a6842cec13b2982011.
//
// Solidity: event RegisterEssence(uint256 indexed profileId, uint256 indexed essenceId, string name, string symbol, string essenceTokenURI, address essenceMw, bytes prepareReturnData)
func (_CyberConnect *CyberConnectFilterer) WatchRegisterEssence(opts *bind.WatchOpts, sink chan<- *CyberConnectRegisterEssence, profileId []*big.Int, essenceId []*big.Int) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var essenceIdRule []interface{}
	for _, essenceIdItem := range essenceId {
		essenceIdRule = append(essenceIdRule, essenceIdItem)
	}

	logs, sub, err := _CyberConnect.contract.WatchLogs(opts, "RegisterEssence", profileIdRule, essenceIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CyberConnectRegisterEssence)
				if err := _CyberConnect.contract.UnpackLog(event, "RegisterEssence", log); err != nil {
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

// ParseRegisterEssence is a log parse operation binding the contract event 0x7f8f05cb8699d58aa2b69fa7a1b1ab0db87ab881493e58a6842cec13b2982011.
//
// Solidity: event RegisterEssence(uint256 indexed profileId, uint256 indexed essenceId, string name, string symbol, string essenceTokenURI, address essenceMw, bytes prepareReturnData)
func (_CyberConnect *CyberConnectFilterer) ParseRegisterEssence(log types.Log) (*CyberConnectRegisterEssence, error) {
	event := new(CyberConnectRegisterEssence)
	if err := _CyberConnect.contract.UnpackLog(event, "RegisterEssence", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CyberConnectSetAvatarIterator is returned from FilterSetAvatar and is used to iterate over the raw logs and unpacked data for SetAvatar events raised by the CyberConnect contract.
type CyberConnectSetAvatarIterator struct {
	Event *CyberConnectSetAvatar // Event containing the contract specifics and raw log

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
func (it *CyberConnectSetAvatarIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CyberConnectSetAvatar)
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
		it.Event = new(CyberConnectSetAvatar)
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
func (it *CyberConnectSetAvatarIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CyberConnectSetAvatarIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CyberConnectSetAvatar represents a SetAvatar event raised by the CyberConnect contract.
type CyberConnectSetAvatar struct {
	ProfileId *big.Int
	NewAvatar string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSetAvatar is a free log retrieval operation binding the contract event 0x60f52aedb1db78381295534bb468e00f2962aea6c3bbb89eb141a27599ddac0e.
//
// Solidity: event SetAvatar(uint256 indexed profileId, string newAvatar)
func (_CyberConnect *CyberConnectFilterer) FilterSetAvatar(opts *bind.FilterOpts, profileId []*big.Int) (*CyberConnectSetAvatarIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _CyberConnect.contract.FilterLogs(opts, "SetAvatar", profileIdRule)
	if err != nil {
		return nil, err
	}
	return &CyberConnectSetAvatarIterator{contract: _CyberConnect.contract, event: "SetAvatar", logs: logs, sub: sub}, nil
}

// WatchSetAvatar is a free log subscription operation binding the contract event 0x60f52aedb1db78381295534bb468e00f2962aea6c3bbb89eb141a27599ddac0e.
//
// Solidity: event SetAvatar(uint256 indexed profileId, string newAvatar)
func (_CyberConnect *CyberConnectFilterer) WatchSetAvatar(opts *bind.WatchOpts, sink chan<- *CyberConnectSetAvatar, profileId []*big.Int) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _CyberConnect.contract.WatchLogs(opts, "SetAvatar", profileIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CyberConnectSetAvatar)
				if err := _CyberConnect.contract.UnpackLog(event, "SetAvatar", log); err != nil {
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

// ParseSetAvatar is a log parse operation binding the contract event 0x60f52aedb1db78381295534bb468e00f2962aea6c3bbb89eb141a27599ddac0e.
//
// Solidity: event SetAvatar(uint256 indexed profileId, string newAvatar)
func (_CyberConnect *CyberConnectFilterer) ParseSetAvatar(log types.Log) (*CyberConnectSetAvatar, error) {
	event := new(CyberConnectSetAvatar)
	if err := _CyberConnect.contract.UnpackLog(event, "SetAvatar", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CyberConnectSetEssenceDataIterator is returned from FilterSetEssenceData and is used to iterate over the raw logs and unpacked data for SetEssenceData events raised by the CyberConnect contract.
type CyberConnectSetEssenceDataIterator struct {
	Event *CyberConnectSetEssenceData // Event containing the contract specifics and raw log

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
func (it *CyberConnectSetEssenceDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CyberConnectSetEssenceData)
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
		it.Event = new(CyberConnectSetEssenceData)
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
func (it *CyberConnectSetEssenceDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CyberConnectSetEssenceDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CyberConnectSetEssenceData represents a SetEssenceData event raised by the CyberConnect contract.
type CyberConnectSetEssenceData struct {
	ProfileId         *big.Int
	EssenceId         *big.Int
	TokenURI          string
	Mw                common.Address
	PrepareReturnData []byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSetEssenceData is a free log retrieval operation binding the contract event 0x1644f271ba3edff28cd2563431920a960fcfebffd068680e682b2de30ca31694.
//
// Solidity: event SetEssenceData(uint256 indexed profileId, uint256 indexed essenceId, string tokenURI, address mw, bytes prepareReturnData)
func (_CyberConnect *CyberConnectFilterer) FilterSetEssenceData(opts *bind.FilterOpts, profileId []*big.Int, essenceId []*big.Int) (*CyberConnectSetEssenceDataIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var essenceIdRule []interface{}
	for _, essenceIdItem := range essenceId {
		essenceIdRule = append(essenceIdRule, essenceIdItem)
	}

	logs, sub, err := _CyberConnect.contract.FilterLogs(opts, "SetEssenceData", profileIdRule, essenceIdRule)
	if err != nil {
		return nil, err
	}
	return &CyberConnectSetEssenceDataIterator{contract: _CyberConnect.contract, event: "SetEssenceData", logs: logs, sub: sub}, nil
}

// WatchSetEssenceData is a free log subscription operation binding the contract event 0x1644f271ba3edff28cd2563431920a960fcfebffd068680e682b2de30ca31694.
//
// Solidity: event SetEssenceData(uint256 indexed profileId, uint256 indexed essenceId, string tokenURI, address mw, bytes prepareReturnData)
func (_CyberConnect *CyberConnectFilterer) WatchSetEssenceData(opts *bind.WatchOpts, sink chan<- *CyberConnectSetEssenceData, profileId []*big.Int, essenceId []*big.Int) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var essenceIdRule []interface{}
	for _, essenceIdItem := range essenceId {
		essenceIdRule = append(essenceIdRule, essenceIdItem)
	}

	logs, sub, err := _CyberConnect.contract.WatchLogs(opts, "SetEssenceData", profileIdRule, essenceIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CyberConnectSetEssenceData)
				if err := _CyberConnect.contract.UnpackLog(event, "SetEssenceData", log); err != nil {
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

// ParseSetEssenceData is a log parse operation binding the contract event 0x1644f271ba3edff28cd2563431920a960fcfebffd068680e682b2de30ca31694.
//
// Solidity: event SetEssenceData(uint256 indexed profileId, uint256 indexed essenceId, string tokenURI, address mw, bytes prepareReturnData)
func (_CyberConnect *CyberConnectFilterer) ParseSetEssenceData(log types.Log) (*CyberConnectSetEssenceData, error) {
	event := new(CyberConnectSetEssenceData)
	if err := _CyberConnect.contract.UnpackLog(event, "SetEssenceData", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CyberConnectSetMetadataIterator is returned from FilterSetMetadata and is used to iterate over the raw logs and unpacked data for SetMetadata events raised by the CyberConnect contract.
type CyberConnectSetMetadataIterator struct {
	Event *CyberConnectSetMetadata // Event containing the contract specifics and raw log

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
func (it *CyberConnectSetMetadataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CyberConnectSetMetadata)
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
		it.Event = new(CyberConnectSetMetadata)
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
func (it *CyberConnectSetMetadataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CyberConnectSetMetadataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CyberConnectSetMetadata represents a SetMetadata event raised by the CyberConnect contract.
type CyberConnectSetMetadata struct {
	ProfileId   *big.Int
	NewMetadata string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetMetadata is a free log retrieval operation binding the contract event 0xcf080687acfb4df18bbb9959fd7284daa7103f9d1c9164e70999fdedac2a3030.
//
// Solidity: event SetMetadata(uint256 indexed profileId, string newMetadata)
func (_CyberConnect *CyberConnectFilterer) FilterSetMetadata(opts *bind.FilterOpts, profileId []*big.Int) (*CyberConnectSetMetadataIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _CyberConnect.contract.FilterLogs(opts, "SetMetadata", profileIdRule)
	if err != nil {
		return nil, err
	}
	return &CyberConnectSetMetadataIterator{contract: _CyberConnect.contract, event: "SetMetadata", logs: logs, sub: sub}, nil
}

// WatchSetMetadata is a free log subscription operation binding the contract event 0xcf080687acfb4df18bbb9959fd7284daa7103f9d1c9164e70999fdedac2a3030.
//
// Solidity: event SetMetadata(uint256 indexed profileId, string newMetadata)
func (_CyberConnect *CyberConnectFilterer) WatchSetMetadata(opts *bind.WatchOpts, sink chan<- *CyberConnectSetMetadata, profileId []*big.Int) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _CyberConnect.contract.WatchLogs(opts, "SetMetadata", profileIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CyberConnectSetMetadata)
				if err := _CyberConnect.contract.UnpackLog(event, "SetMetadata", log); err != nil {
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

// ParseSetMetadata is a log parse operation binding the contract event 0xcf080687acfb4df18bbb9959fd7284daa7103f9d1c9164e70999fdedac2a3030.
//
// Solidity: event SetMetadata(uint256 indexed profileId, string newMetadata)
func (_CyberConnect *CyberConnectFilterer) ParseSetMetadata(log types.Log) (*CyberConnectSetMetadata, error) {
	event := new(CyberConnectSetMetadata)
	if err := _CyberConnect.contract.UnpackLog(event, "SetMetadata", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CyberConnectSetNFTDescriptorIterator is returned from FilterSetNFTDescriptor and is used to iterate over the raw logs and unpacked data for SetNFTDescriptor events raised by the CyberConnect contract.
type CyberConnectSetNFTDescriptorIterator struct {
	Event *CyberConnectSetNFTDescriptor // Event containing the contract specifics and raw log

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
func (it *CyberConnectSetNFTDescriptorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CyberConnectSetNFTDescriptor)
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
		it.Event = new(CyberConnectSetNFTDescriptor)
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
func (it *CyberConnectSetNFTDescriptorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CyberConnectSetNFTDescriptorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CyberConnectSetNFTDescriptor represents a SetNFTDescriptor event raised by the CyberConnect contract.
type CyberConnectSetNFTDescriptor struct {
	NewDescriptor common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetNFTDescriptor is a free log retrieval operation binding the contract event 0xb72b36bb50852d5d293221b9d1357433e4df37b01e636d3587c2bab7e8531ead.
//
// Solidity: event SetNFTDescriptor(address indexed newDescriptor)
func (_CyberConnect *CyberConnectFilterer) FilterSetNFTDescriptor(opts *bind.FilterOpts, newDescriptor []common.Address) (*CyberConnectSetNFTDescriptorIterator, error) {

	var newDescriptorRule []interface{}
	for _, newDescriptorItem := range newDescriptor {
		newDescriptorRule = append(newDescriptorRule, newDescriptorItem)
	}

	logs, sub, err := _CyberConnect.contract.FilterLogs(opts, "SetNFTDescriptor", newDescriptorRule)
	if err != nil {
		return nil, err
	}
	return &CyberConnectSetNFTDescriptorIterator{contract: _CyberConnect.contract, event: "SetNFTDescriptor", logs: logs, sub: sub}, nil
}

// WatchSetNFTDescriptor is a free log subscription operation binding the contract event 0xb72b36bb50852d5d293221b9d1357433e4df37b01e636d3587c2bab7e8531ead.
//
// Solidity: event SetNFTDescriptor(address indexed newDescriptor)
func (_CyberConnect *CyberConnectFilterer) WatchSetNFTDescriptor(opts *bind.WatchOpts, sink chan<- *CyberConnectSetNFTDescriptor, newDescriptor []common.Address) (event.Subscription, error) {

	var newDescriptorRule []interface{}
	for _, newDescriptorItem := range newDescriptor {
		newDescriptorRule = append(newDescriptorRule, newDescriptorItem)
	}

	logs, sub, err := _CyberConnect.contract.WatchLogs(opts, "SetNFTDescriptor", newDescriptorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CyberConnectSetNFTDescriptor)
				if err := _CyberConnect.contract.UnpackLog(event, "SetNFTDescriptor", log); err != nil {
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

// ParseSetNFTDescriptor is a log parse operation binding the contract event 0xb72b36bb50852d5d293221b9d1357433e4df37b01e636d3587c2bab7e8531ead.
//
// Solidity: event SetNFTDescriptor(address indexed newDescriptor)
func (_CyberConnect *CyberConnectFilterer) ParseSetNFTDescriptor(log types.Log) (*CyberConnectSetNFTDescriptor, error) {
	event := new(CyberConnectSetNFTDescriptor)
	if err := _CyberConnect.contract.UnpackLog(event, "SetNFTDescriptor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CyberConnectSetNamespaceOwnerIterator is returned from FilterSetNamespaceOwner and is used to iterate over the raw logs and unpacked data for SetNamespaceOwner events raised by the CyberConnect contract.
type CyberConnectSetNamespaceOwnerIterator struct {
	Event *CyberConnectSetNamespaceOwner // Event containing the contract specifics and raw log

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
func (it *CyberConnectSetNamespaceOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CyberConnectSetNamespaceOwner)
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
		it.Event = new(CyberConnectSetNamespaceOwner)
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
func (it *CyberConnectSetNamespaceOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CyberConnectSetNamespaceOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CyberConnectSetNamespaceOwner represents a SetNamespaceOwner event raised by the CyberConnect contract.
type CyberConnectSetNamespaceOwner struct {
	PreOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetNamespaceOwner is a free log retrieval operation binding the contract event 0x46e05b7612d0b60892e73d72895fb0752deda3fa69984aca10c3cc72889ef7df.
//
// Solidity: event SetNamespaceOwner(address indexed preOwner, address indexed newOwner)
func (_CyberConnect *CyberConnectFilterer) FilterSetNamespaceOwner(opts *bind.FilterOpts, preOwner []common.Address, newOwner []common.Address) (*CyberConnectSetNamespaceOwnerIterator, error) {

	var preOwnerRule []interface{}
	for _, preOwnerItem := range preOwner {
		preOwnerRule = append(preOwnerRule, preOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CyberConnect.contract.FilterLogs(opts, "SetNamespaceOwner", preOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CyberConnectSetNamespaceOwnerIterator{contract: _CyberConnect.contract, event: "SetNamespaceOwner", logs: logs, sub: sub}, nil
}

// WatchSetNamespaceOwner is a free log subscription operation binding the contract event 0x46e05b7612d0b60892e73d72895fb0752deda3fa69984aca10c3cc72889ef7df.
//
// Solidity: event SetNamespaceOwner(address indexed preOwner, address indexed newOwner)
func (_CyberConnect *CyberConnectFilterer) WatchSetNamespaceOwner(opts *bind.WatchOpts, sink chan<- *CyberConnectSetNamespaceOwner, preOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var preOwnerRule []interface{}
	for _, preOwnerItem := range preOwner {
		preOwnerRule = append(preOwnerRule, preOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CyberConnect.contract.WatchLogs(opts, "SetNamespaceOwner", preOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CyberConnectSetNamespaceOwner)
				if err := _CyberConnect.contract.UnpackLog(event, "SetNamespaceOwner", log); err != nil {
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

// ParseSetNamespaceOwner is a log parse operation binding the contract event 0x46e05b7612d0b60892e73d72895fb0752deda3fa69984aca10c3cc72889ef7df.
//
// Solidity: event SetNamespaceOwner(address indexed preOwner, address indexed newOwner)
func (_CyberConnect *CyberConnectFilterer) ParseSetNamespaceOwner(log types.Log) (*CyberConnectSetNamespaceOwner, error) {
	event := new(CyberConnectSetNamespaceOwner)
	if err := _CyberConnect.contract.UnpackLog(event, "SetNamespaceOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CyberConnectSetOperatorApprovalIterator is returned from FilterSetOperatorApproval and is used to iterate over the raw logs and unpacked data for SetOperatorApproval events raised by the CyberConnect contract.
type CyberConnectSetOperatorApprovalIterator struct {
	Event *CyberConnectSetOperatorApproval // Event containing the contract specifics and raw log

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
func (it *CyberConnectSetOperatorApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CyberConnectSetOperatorApproval)
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
		it.Event = new(CyberConnectSetOperatorApproval)
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
func (it *CyberConnectSetOperatorApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CyberConnectSetOperatorApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CyberConnectSetOperatorApproval represents a SetOperatorApproval event raised by the CyberConnect contract.
type CyberConnectSetOperatorApproval struct {
	ProfileId    *big.Int
	Operator     common.Address
	PrevApproved bool
	Approved     bool
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSetOperatorApproval is a free log retrieval operation binding the contract event 0x470544bc4035b5a73e4bb39628742f459c69af3a00b7568beff4bcb76aa51b6e.
//
// Solidity: event SetOperatorApproval(uint256 indexed profileId, address indexed operator, bool prevApproved, bool approved)
func (_CyberConnect *CyberConnectFilterer) FilterSetOperatorApproval(opts *bind.FilterOpts, profileId []*big.Int, operator []common.Address) (*CyberConnectSetOperatorApprovalIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _CyberConnect.contract.FilterLogs(opts, "SetOperatorApproval", profileIdRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &CyberConnectSetOperatorApprovalIterator{contract: _CyberConnect.contract, event: "SetOperatorApproval", logs: logs, sub: sub}, nil
}

// WatchSetOperatorApproval is a free log subscription operation binding the contract event 0x470544bc4035b5a73e4bb39628742f459c69af3a00b7568beff4bcb76aa51b6e.
//
// Solidity: event SetOperatorApproval(uint256 indexed profileId, address indexed operator, bool prevApproved, bool approved)
func (_CyberConnect *CyberConnectFilterer) WatchSetOperatorApproval(opts *bind.WatchOpts, sink chan<- *CyberConnectSetOperatorApproval, profileId []*big.Int, operator []common.Address) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _CyberConnect.contract.WatchLogs(opts, "SetOperatorApproval", profileIdRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CyberConnectSetOperatorApproval)
				if err := _CyberConnect.contract.UnpackLog(event, "SetOperatorApproval", log); err != nil {
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

// ParseSetOperatorApproval is a log parse operation binding the contract event 0x470544bc4035b5a73e4bb39628742f459c69af3a00b7568beff4bcb76aa51b6e.
//
// Solidity: event SetOperatorApproval(uint256 indexed profileId, address indexed operator, bool prevApproved, bool approved)
func (_CyberConnect *CyberConnectFilterer) ParseSetOperatorApproval(log types.Log) (*CyberConnectSetOperatorApproval, error) {
	event := new(CyberConnectSetOperatorApproval)
	if err := _CyberConnect.contract.UnpackLog(event, "SetOperatorApproval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CyberConnectSetPrimaryProfileIterator is returned from FilterSetPrimaryProfile and is used to iterate over the raw logs and unpacked data for SetPrimaryProfile events raised by the CyberConnect contract.
type CyberConnectSetPrimaryProfileIterator struct {
	Event *CyberConnectSetPrimaryProfile // Event containing the contract specifics and raw log

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
func (it *CyberConnectSetPrimaryProfileIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CyberConnectSetPrimaryProfile)
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
		it.Event = new(CyberConnectSetPrimaryProfile)
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
func (it *CyberConnectSetPrimaryProfileIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CyberConnectSetPrimaryProfileIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CyberConnectSetPrimaryProfile represents a SetPrimaryProfile event raised by the CyberConnect contract.
type CyberConnectSetPrimaryProfile struct {
	User      common.Address
	ProfileId *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSetPrimaryProfile is a free log retrieval operation binding the contract event 0xdf187b01e55e0f335a75956346e6d4417e6ea27aed3bc1bf2106d9d60bdc50d0.
//
// Solidity: event SetPrimaryProfile(address indexed user, uint256 indexed profileId)
func (_CyberConnect *CyberConnectFilterer) FilterSetPrimaryProfile(opts *bind.FilterOpts, user []common.Address, profileId []*big.Int) (*CyberConnectSetPrimaryProfileIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _CyberConnect.contract.FilterLogs(opts, "SetPrimaryProfile", userRule, profileIdRule)
	if err != nil {
		return nil, err
	}
	return &CyberConnectSetPrimaryProfileIterator{contract: _CyberConnect.contract, event: "SetPrimaryProfile", logs: logs, sub: sub}, nil
}

// WatchSetPrimaryProfile is a free log subscription operation binding the contract event 0xdf187b01e55e0f335a75956346e6d4417e6ea27aed3bc1bf2106d9d60bdc50d0.
//
// Solidity: event SetPrimaryProfile(address indexed user, uint256 indexed profileId)
func (_CyberConnect *CyberConnectFilterer) WatchSetPrimaryProfile(opts *bind.WatchOpts, sink chan<- *CyberConnectSetPrimaryProfile, user []common.Address, profileId []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _CyberConnect.contract.WatchLogs(opts, "SetPrimaryProfile", userRule, profileIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CyberConnectSetPrimaryProfile)
				if err := _CyberConnect.contract.UnpackLog(event, "SetPrimaryProfile", log); err != nil {
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

// ParseSetPrimaryProfile is a log parse operation binding the contract event 0xdf187b01e55e0f335a75956346e6d4417e6ea27aed3bc1bf2106d9d60bdc50d0.
//
// Solidity: event SetPrimaryProfile(address indexed user, uint256 indexed profileId)
func (_CyberConnect *CyberConnectFilterer) ParseSetPrimaryProfile(log types.Log) (*CyberConnectSetPrimaryProfile, error) {
	event := new(CyberConnectSetPrimaryProfile)
	if err := _CyberConnect.contract.UnpackLog(event, "SetPrimaryProfile", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CyberConnectSetSubscribeDataIterator is returned from FilterSetSubscribeData and is used to iterate over the raw logs and unpacked data for SetSubscribeData events raised by the CyberConnect contract.
type CyberConnectSetSubscribeDataIterator struct {
	Event *CyberConnectSetSubscribeData // Event containing the contract specifics and raw log

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
func (it *CyberConnectSetSubscribeDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CyberConnectSetSubscribeData)
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
		it.Event = new(CyberConnectSetSubscribeData)
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
func (it *CyberConnectSetSubscribeDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CyberConnectSetSubscribeDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CyberConnectSetSubscribeData represents a SetSubscribeData event raised by the CyberConnect contract.
type CyberConnectSetSubscribeData struct {
	ProfileId         *big.Int
	TokenURI          string
	Mw                common.Address
	PrepareReturnData []byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSetSubscribeData is a free log retrieval operation binding the contract event 0x56e2b6a9becdb5836c1184738ad1cddd0096556aa7ada94493a4a4feb117bd86.
//
// Solidity: event SetSubscribeData(uint256 indexed profileId, string tokenURI, address mw, bytes prepareReturnData)
func (_CyberConnect *CyberConnectFilterer) FilterSetSubscribeData(opts *bind.FilterOpts, profileId []*big.Int) (*CyberConnectSetSubscribeDataIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _CyberConnect.contract.FilterLogs(opts, "SetSubscribeData", profileIdRule)
	if err != nil {
		return nil, err
	}
	return &CyberConnectSetSubscribeDataIterator{contract: _CyberConnect.contract, event: "SetSubscribeData", logs: logs, sub: sub}, nil
}

// WatchSetSubscribeData is a free log subscription operation binding the contract event 0x56e2b6a9becdb5836c1184738ad1cddd0096556aa7ada94493a4a4feb117bd86.
//
// Solidity: event SetSubscribeData(uint256 indexed profileId, string tokenURI, address mw, bytes prepareReturnData)
func (_CyberConnect *CyberConnectFilterer) WatchSetSubscribeData(opts *bind.WatchOpts, sink chan<- *CyberConnectSetSubscribeData, profileId []*big.Int) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _CyberConnect.contract.WatchLogs(opts, "SetSubscribeData", profileIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CyberConnectSetSubscribeData)
				if err := _CyberConnect.contract.UnpackLog(event, "SetSubscribeData", log); err != nil {
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

// ParseSetSubscribeData is a log parse operation binding the contract event 0x56e2b6a9becdb5836c1184738ad1cddd0096556aa7ada94493a4a4feb117bd86.
//
// Solidity: event SetSubscribeData(uint256 indexed profileId, string tokenURI, address mw, bytes prepareReturnData)
func (_CyberConnect *CyberConnectFilterer) ParseSetSubscribeData(log types.Log) (*CyberConnectSetSubscribeData, error) {
	event := new(CyberConnectSetSubscribeData)
	if err := _CyberConnect.contract.UnpackLog(event, "SetSubscribeData", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CyberConnectSubscribeIterator is returned from FilterSubscribe and is used to iterate over the raw logs and unpacked data for Subscribe events raised by the CyberConnect contract.
type CyberConnectSubscribeIterator struct {
	Event *CyberConnectSubscribe // Event containing the contract specifics and raw log

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
func (it *CyberConnectSubscribeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CyberConnectSubscribe)
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
		it.Event = new(CyberConnectSubscribe)
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
func (it *CyberConnectSubscribeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CyberConnectSubscribeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CyberConnectSubscribe represents a Subscribe event raised by the CyberConnect contract.
type CyberConnectSubscribe struct {
	Sender     common.Address
	ProfileIds []*big.Int
	PreDatas   [][]byte
	PostDatas  [][]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSubscribe is a free log retrieval operation binding the contract event 0xcda1d04cbb59c6063db6682ab282a88dc856e841cd4d5562471f2f8bbcbd6d6b.
//
// Solidity: event Subscribe(address indexed sender, uint256[] profileIds, bytes[] preDatas, bytes[] postDatas)
func (_CyberConnect *CyberConnectFilterer) FilterSubscribe(opts *bind.FilterOpts, sender []common.Address) (*CyberConnectSubscribeIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CyberConnect.contract.FilterLogs(opts, "Subscribe", senderRule)
	if err != nil {
		return nil, err
	}
	return &CyberConnectSubscribeIterator{contract: _CyberConnect.contract, event: "Subscribe", logs: logs, sub: sub}, nil
}

// WatchSubscribe is a free log subscription operation binding the contract event 0xcda1d04cbb59c6063db6682ab282a88dc856e841cd4d5562471f2f8bbcbd6d6b.
//
// Solidity: event Subscribe(address indexed sender, uint256[] profileIds, bytes[] preDatas, bytes[] postDatas)
func (_CyberConnect *CyberConnectFilterer) WatchSubscribe(opts *bind.WatchOpts, sink chan<- *CyberConnectSubscribe, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CyberConnect.contract.WatchLogs(opts, "Subscribe", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CyberConnectSubscribe)
				if err := _CyberConnect.contract.UnpackLog(event, "Subscribe", log); err != nil {
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

// ParseSubscribe is a log parse operation binding the contract event 0xcda1d04cbb59c6063db6682ab282a88dc856e841cd4d5562471f2f8bbcbd6d6b.
//
// Solidity: event Subscribe(address indexed sender, uint256[] profileIds, bytes[] preDatas, bytes[] postDatas)
func (_CyberConnect *CyberConnectFilterer) ParseSubscribe(log types.Log) (*CyberConnectSubscribe, error) {
	event := new(CyberConnectSubscribe)
	if err := _CyberConnect.contract.UnpackLog(event, "Subscribe", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CyberConnectTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the CyberConnect contract.
type CyberConnectTransferIterator struct {
	Event *CyberConnectTransfer // Event containing the contract specifics and raw log

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
func (it *CyberConnectTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CyberConnectTransfer)
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
		it.Event = new(CyberConnectTransfer)
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
func (it *CyberConnectTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CyberConnectTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CyberConnectTransfer represents a Transfer event raised by the CyberConnect contract.
type CyberConnectTransfer struct {
	From common.Address
	To   common.Address
	Id   *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed id)
func (_CyberConnect *CyberConnectFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, id []*big.Int) (*CyberConnectTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _CyberConnect.contract.FilterLogs(opts, "Transfer", fromRule, toRule, idRule)
	if err != nil {
		return nil, err
	}
	return &CyberConnectTransferIterator{contract: _CyberConnect.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed id)
func (_CyberConnect *CyberConnectFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CyberConnectTransfer, from []common.Address, to []common.Address, id []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _CyberConnect.contract.WatchLogs(opts, "Transfer", fromRule, toRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CyberConnectTransfer)
				if err := _CyberConnect.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed id)
func (_CyberConnect *CyberConnectFilterer) ParseTransfer(log types.Log) (*CyberConnectTransfer, error) {
	event := new(CyberConnectTransfer)
	if err := _CyberConnect.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CyberConnectUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the CyberConnect contract.
type CyberConnectUnpausedIterator struct {
	Event *CyberConnectUnpaused // Event containing the contract specifics and raw log

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
func (it *CyberConnectUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CyberConnectUnpaused)
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
		it.Event = new(CyberConnectUnpaused)
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
func (it *CyberConnectUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CyberConnectUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CyberConnectUnpaused represents a Unpaused event raised by the CyberConnect contract.
type CyberConnectUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_CyberConnect *CyberConnectFilterer) FilterUnpaused(opts *bind.FilterOpts) (*CyberConnectUnpausedIterator, error) {

	logs, sub, err := _CyberConnect.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &CyberConnectUnpausedIterator{contract: _CyberConnect.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_CyberConnect *CyberConnectFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *CyberConnectUnpaused) (event.Subscription, error) {

	logs, sub, err := _CyberConnect.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CyberConnectUnpaused)
				if err := _CyberConnect.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_CyberConnect *CyberConnectFilterer) ParseUnpaused(log types.Log) (*CyberConnectUnpaused, error) {
	event := new(CyberConnectUnpaused)
	if err := _CyberConnect.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CyberConnectUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the CyberConnect contract.
type CyberConnectUpgradedIterator struct {
	Event *CyberConnectUpgraded // Event containing the contract specifics and raw log

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
func (it *CyberConnectUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CyberConnectUpgraded)
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
		it.Event = new(CyberConnectUpgraded)
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
func (it *CyberConnectUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CyberConnectUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CyberConnectUpgraded represents a Upgraded event raised by the CyberConnect contract.
type CyberConnectUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_CyberConnect *CyberConnectFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*CyberConnectUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _CyberConnect.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &CyberConnectUpgradedIterator{contract: _CyberConnect.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_CyberConnect *CyberConnectFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *CyberConnectUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _CyberConnect.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CyberConnectUpgraded)
				if err := _CyberConnect.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_CyberConnect *CyberConnectFilterer) ParseUpgraded(log types.Log) (*CyberConnectUpgraded, error) {
	event := new(CyberConnectUpgraded)
	if err := _CyberConnect.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
