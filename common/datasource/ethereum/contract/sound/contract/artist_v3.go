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

// ArtistV3MetaData contains all meta data concerning the ArtistV3 contract.
var ArtistV3MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumArtistV3.TimeType\",\"name\":\"timeType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"editionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"newTime\",\"type\":\"uint32\"}],\"name\":\"AuctionTimeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"editionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"fundingRecipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"quantity\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"royaltyBPS\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"startTime\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"endTime\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"permissionedQuantity\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signerAddress\",\"type\":\"address\"}],\"name\":\"EditionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"editionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"numSold\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"}],\"name\":\"EditionPurchased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"editionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"permissionedQuantity\",\"type\":\"uint32\"}],\"name\":\"PermissionedQuantitySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"editionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signerAddress\",\"type\":\"address\"}],\"name\":\"SignerAddressSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_editionId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"buyEdition\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_fundingRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_quantity\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_royaltyBPS\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_startTime\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_endTime\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_permissionedQuantity\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_signerAddress\",\"type\":\"address\"}],\"name\":\"createEdition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"depositedForEdition\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"editionCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"editions\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"fundingRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"numSold\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"quantity\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"royaltyBPS\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"startTime\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"endTime\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"permissionedQuantity\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"signerAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_artistId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_baseURI\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"ownersOfTokenIds\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_salePrice\",\"type\":\"uint256\"}],\"name\":\"royaltyInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"fundingRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"royaltyAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_editionId\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_endTime\",\"type\":\"uint32\"}],\"name\":\"setEndTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_editionId\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_permissionedQuantity\",\"type\":\"uint32\"}],\"name\":\"setPermissionedQuantity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_editionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_newSignerAddress\",\"type\":\"address\"}],\"name\":\"setSignerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_editionId\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_startTime\",\"type\":\"uint32\"}],\"name\":\"setStartTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenToEdition\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_editionId\",\"type\":\"uint256\"}],\"name\":\"withdrawFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"withdrawnForEdition\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ArtistV3ABI is the input ABI used to generate the binding from.
// Deprecated: Use ArtistV3MetaData.ABI instead.
var ArtistV3ABI = ArtistV3MetaData.ABI

// ArtistV3 is an auto generated Go binding around an Ethereum contract.
type ArtistV3 struct {
	ArtistV3Caller     // Read-only binding to the contract
	ArtistV3Transactor // Write-only binding to the contract
	ArtistV3Filterer   // Log filterer for contract events
}

// ArtistV3Caller is an auto generated read-only Go binding around an Ethereum contract.
type ArtistV3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArtistV3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ArtistV3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArtistV3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArtistV3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArtistV3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArtistV3Session struct {
	Contract     *ArtistV3         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArtistV3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArtistV3CallerSession struct {
	Contract *ArtistV3Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ArtistV3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArtistV3TransactorSession struct {
	Contract     *ArtistV3Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ArtistV3Raw is an auto generated low-level Go binding around an Ethereum contract.
type ArtistV3Raw struct {
	Contract *ArtistV3 // Generic contract binding to access the raw methods on
}

// ArtistV3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArtistV3CallerRaw struct {
	Contract *ArtistV3Caller // Generic read-only contract binding to access the raw methods on
}

// ArtistV3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArtistV3TransactorRaw struct {
	Contract *ArtistV3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewArtistV3 creates a new instance of ArtistV3, bound to a specific deployed contract.
func NewArtistV3(address common.Address, backend bind.ContractBackend) (*ArtistV3, error) {
	contract, err := bindArtistV3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ArtistV3{ArtistV3Caller: ArtistV3Caller{contract: contract}, ArtistV3Transactor: ArtistV3Transactor{contract: contract}, ArtistV3Filterer: ArtistV3Filterer{contract: contract}}, nil
}

// NewArtistV3Caller creates a new read-only instance of ArtistV3, bound to a specific deployed contract.
func NewArtistV3Caller(address common.Address, caller bind.ContractCaller) (*ArtistV3Caller, error) {
	contract, err := bindArtistV3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArtistV3Caller{contract: contract}, nil
}

// NewArtistV3Transactor creates a new write-only instance of ArtistV3, bound to a specific deployed contract.
func NewArtistV3Transactor(address common.Address, transactor bind.ContractTransactor) (*ArtistV3Transactor, error) {
	contract, err := bindArtistV3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArtistV3Transactor{contract: contract}, nil
}

// NewArtistV3Filterer creates a new log filterer instance of ArtistV3, bound to a specific deployed contract.
func NewArtistV3Filterer(address common.Address, filterer bind.ContractFilterer) (*ArtistV3Filterer, error) {
	contract, err := bindArtistV3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArtistV3Filterer{contract: contract}, nil
}

// bindArtistV3 binds a generic wrapper to an already deployed contract.
func bindArtistV3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ArtistV3ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArtistV3 *ArtistV3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArtistV3.Contract.ArtistV3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArtistV3 *ArtistV3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArtistV3.Contract.ArtistV3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArtistV3 *ArtistV3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArtistV3.Contract.ArtistV3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArtistV3 *ArtistV3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArtistV3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArtistV3 *ArtistV3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArtistV3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArtistV3 *ArtistV3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArtistV3.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ArtistV3 *ArtistV3Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ArtistV3.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ArtistV3 *ArtistV3Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ArtistV3.Contract.BalanceOf(&_ArtistV3.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ArtistV3 *ArtistV3CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ArtistV3.Contract.BalanceOf(&_ArtistV3.CallOpts, owner)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_ArtistV3 *ArtistV3Caller) ContractURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ArtistV3.contract.Call(opts, &out, "contractURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_ArtistV3 *ArtistV3Session) ContractURI() (string, error) {
	return _ArtistV3.Contract.ContractURI(&_ArtistV3.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_ArtistV3 *ArtistV3CallerSession) ContractURI() (string, error) {
	return _ArtistV3.Contract.ContractURI(&_ArtistV3.CallOpts)
}

// DepositedForEdition is a free data retrieval call binding the contract method 0xe1a3d573.
//
// Solidity: function depositedForEdition(uint256 ) view returns(uint256)
func (_ArtistV3 *ArtistV3Caller) DepositedForEdition(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ArtistV3.contract.Call(opts, &out, "depositedForEdition", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositedForEdition is a free data retrieval call binding the contract method 0xe1a3d573.
//
// Solidity: function depositedForEdition(uint256 ) view returns(uint256)
func (_ArtistV3 *ArtistV3Session) DepositedForEdition(arg0 *big.Int) (*big.Int, error) {
	return _ArtistV3.Contract.DepositedForEdition(&_ArtistV3.CallOpts, arg0)
}

// DepositedForEdition is a free data retrieval call binding the contract method 0xe1a3d573.
//
// Solidity: function depositedForEdition(uint256 ) view returns(uint256)
func (_ArtistV3 *ArtistV3CallerSession) DepositedForEdition(arg0 *big.Int) (*big.Int, error) {
	return _ArtistV3.Contract.DepositedForEdition(&_ArtistV3.CallOpts, arg0)
}

// EditionCount is a free data retrieval call binding the contract method 0x4bf44026.
//
// Solidity: function editionCount() view returns(uint256)
func (_ArtistV3 *ArtistV3Caller) EditionCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ArtistV3.contract.Call(opts, &out, "editionCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EditionCount is a free data retrieval call binding the contract method 0x4bf44026.
//
// Solidity: function editionCount() view returns(uint256)
func (_ArtistV3 *ArtistV3Session) EditionCount() (*big.Int, error) {
	return _ArtistV3.Contract.EditionCount(&_ArtistV3.CallOpts)
}

// EditionCount is a free data retrieval call binding the contract method 0x4bf44026.
//
// Solidity: function editionCount() view returns(uint256)
func (_ArtistV3 *ArtistV3CallerSession) EditionCount() (*big.Int, error) {
	return _ArtistV3.Contract.EditionCount(&_ArtistV3.CallOpts)
}

// Editions is a free data retrieval call binding the contract method 0x279c806e.
//
// Solidity: function editions(uint256 ) view returns(address fundingRecipient, uint256 price, uint32 numSold, uint32 quantity, uint32 royaltyBPS, uint32 startTime, uint32 endTime, uint32 permissionedQuantity, address signerAddress)
func (_ArtistV3 *ArtistV3Caller) Editions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	FundingRecipient     common.Address
	Price                *big.Int
	NumSold              uint32
	Quantity             uint32
	RoyaltyBPS           uint32
	StartTime            uint32
	EndTime              uint32
	PermissionedQuantity uint32
	SignerAddress        common.Address
}, error) {
	var out []interface{}
	err := _ArtistV3.contract.Call(opts, &out, "editions", arg0)

	outstruct := new(struct {
		FundingRecipient     common.Address
		Price                *big.Int
		NumSold              uint32
		Quantity             uint32
		RoyaltyBPS           uint32
		StartTime            uint32
		EndTime              uint32
		PermissionedQuantity uint32
		SignerAddress        common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FundingRecipient = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Price = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.NumSold = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.Quantity = *abi.ConvertType(out[3], new(uint32)).(*uint32)
	outstruct.RoyaltyBPS = *abi.ConvertType(out[4], new(uint32)).(*uint32)
	outstruct.StartTime = *abi.ConvertType(out[5], new(uint32)).(*uint32)
	outstruct.EndTime = *abi.ConvertType(out[6], new(uint32)).(*uint32)
	outstruct.PermissionedQuantity = *abi.ConvertType(out[7], new(uint32)).(*uint32)
	outstruct.SignerAddress = *abi.ConvertType(out[8], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Editions is a free data retrieval call binding the contract method 0x279c806e.
//
// Solidity: function editions(uint256 ) view returns(address fundingRecipient, uint256 price, uint32 numSold, uint32 quantity, uint32 royaltyBPS, uint32 startTime, uint32 endTime, uint32 permissionedQuantity, address signerAddress)
func (_ArtistV3 *ArtistV3Session) Editions(arg0 *big.Int) (struct {
	FundingRecipient     common.Address
	Price                *big.Int
	NumSold              uint32
	Quantity             uint32
	RoyaltyBPS           uint32
	StartTime            uint32
	EndTime              uint32
	PermissionedQuantity uint32
	SignerAddress        common.Address
}, error) {
	return _ArtistV3.Contract.Editions(&_ArtistV3.CallOpts, arg0)
}

// Editions is a free data retrieval call binding the contract method 0x279c806e.
//
// Solidity: function editions(uint256 ) view returns(address fundingRecipient, uint256 price, uint32 numSold, uint32 quantity, uint32 royaltyBPS, uint32 startTime, uint32 endTime, uint32 permissionedQuantity, address signerAddress)
func (_ArtistV3 *ArtistV3CallerSession) Editions(arg0 *big.Int) (struct {
	FundingRecipient     common.Address
	Price                *big.Int
	NumSold              uint32
	Quantity             uint32
	RoyaltyBPS           uint32
	StartTime            uint32
	EndTime              uint32
	PermissionedQuantity uint32
	SignerAddress        common.Address
}, error) {
	return _ArtistV3.Contract.Editions(&_ArtistV3.CallOpts, arg0)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ArtistV3 *ArtistV3Caller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ArtistV3.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ArtistV3 *ArtistV3Session) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ArtistV3.Contract.GetApproved(&_ArtistV3.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ArtistV3 *ArtistV3CallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ArtistV3.Contract.GetApproved(&_ArtistV3.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ArtistV3 *ArtistV3Caller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _ArtistV3.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ArtistV3 *ArtistV3Session) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ArtistV3.Contract.IsApprovedForAll(&_ArtistV3.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ArtistV3 *ArtistV3CallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ArtistV3.Contract.IsApprovedForAll(&_ArtistV3.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ArtistV3 *ArtistV3Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ArtistV3.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ArtistV3 *ArtistV3Session) Name() (string, error) {
	return _ArtistV3.Contract.Name(&_ArtistV3.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ArtistV3 *ArtistV3CallerSession) Name() (string, error) {
	return _ArtistV3.Contract.Name(&_ArtistV3.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ArtistV3 *ArtistV3Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ArtistV3.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ArtistV3 *ArtistV3Session) Owner() (common.Address, error) {
	return _ArtistV3.Contract.Owner(&_ArtistV3.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ArtistV3 *ArtistV3CallerSession) Owner() (common.Address, error) {
	return _ArtistV3.Contract.Owner(&_ArtistV3.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ArtistV3 *ArtistV3Caller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ArtistV3.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ArtistV3 *ArtistV3Session) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ArtistV3.Contract.OwnerOf(&_ArtistV3.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ArtistV3 *ArtistV3CallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ArtistV3.Contract.OwnerOf(&_ArtistV3.CallOpts, tokenId)
}

// OwnersOfTokenIds is a free data retrieval call binding the contract method 0x52f5c2e4.
//
// Solidity: function ownersOfTokenIds(uint256[] _tokenIds) view returns(address[])
func (_ArtistV3 *ArtistV3Caller) OwnersOfTokenIds(opts *bind.CallOpts, _tokenIds []*big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _ArtistV3.contract.Call(opts, &out, "ownersOfTokenIds", _tokenIds)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// OwnersOfTokenIds is a free data retrieval call binding the contract method 0x52f5c2e4.
//
// Solidity: function ownersOfTokenIds(uint256[] _tokenIds) view returns(address[])
func (_ArtistV3 *ArtistV3Session) OwnersOfTokenIds(_tokenIds []*big.Int) ([]common.Address, error) {
	return _ArtistV3.Contract.OwnersOfTokenIds(&_ArtistV3.CallOpts, _tokenIds)
}

// OwnersOfTokenIds is a free data retrieval call binding the contract method 0x52f5c2e4.
//
// Solidity: function ownersOfTokenIds(uint256[] _tokenIds) view returns(address[])
func (_ArtistV3 *ArtistV3CallerSession) OwnersOfTokenIds(_tokenIds []*big.Int) ([]common.Address, error) {
	return _ArtistV3.Contract.OwnersOfTokenIds(&_ArtistV3.CallOpts, _tokenIds)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 _tokenId, uint256 _salePrice) view returns(address fundingRecipient, uint256 royaltyAmount)
func (_ArtistV3 *ArtistV3Caller) RoyaltyInfo(opts *bind.CallOpts, _tokenId *big.Int, _salePrice *big.Int) (struct {
	FundingRecipient common.Address
	RoyaltyAmount    *big.Int
}, error) {
	var out []interface{}
	err := _ArtistV3.contract.Call(opts, &out, "royaltyInfo", _tokenId, _salePrice)

	outstruct := new(struct {
		FundingRecipient common.Address
		RoyaltyAmount    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FundingRecipient = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.RoyaltyAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 _tokenId, uint256 _salePrice) view returns(address fundingRecipient, uint256 royaltyAmount)
func (_ArtistV3 *ArtistV3Session) RoyaltyInfo(_tokenId *big.Int, _salePrice *big.Int) (struct {
	FundingRecipient common.Address
	RoyaltyAmount    *big.Int
}, error) {
	return _ArtistV3.Contract.RoyaltyInfo(&_ArtistV3.CallOpts, _tokenId, _salePrice)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 _tokenId, uint256 _salePrice) view returns(address fundingRecipient, uint256 royaltyAmount)
func (_ArtistV3 *ArtistV3CallerSession) RoyaltyInfo(_tokenId *big.Int, _salePrice *big.Int) (struct {
	FundingRecipient common.Address
	RoyaltyAmount    *big.Int
}, error) {
	return _ArtistV3.Contract.RoyaltyInfo(&_ArtistV3.CallOpts, _tokenId, _salePrice)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (_ArtistV3 *ArtistV3Caller) SupportsInterface(opts *bind.CallOpts, _interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ArtistV3.contract.Call(opts, &out, "supportsInterface", _interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (_ArtistV3 *ArtistV3Session) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _ArtistV3.Contract.SupportsInterface(&_ArtistV3.CallOpts, _interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (_ArtistV3 *ArtistV3CallerSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _ArtistV3.Contract.SupportsInterface(&_ArtistV3.CallOpts, _interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ArtistV3 *ArtistV3Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ArtistV3.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ArtistV3 *ArtistV3Session) Symbol() (string, error) {
	return _ArtistV3.Contract.Symbol(&_ArtistV3.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ArtistV3 *ArtistV3CallerSession) Symbol() (string, error) {
	return _ArtistV3.Contract.Symbol(&_ArtistV3.CallOpts)
}

// TokenToEdition is a free data retrieval call binding the contract method 0x602787ed.
//
// Solidity: function tokenToEdition(uint256 _tokenId) view returns(uint256)
func (_ArtistV3 *ArtistV3Caller) TokenToEdition(opts *bind.CallOpts, _tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ArtistV3.contract.Call(opts, &out, "tokenToEdition", _tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenToEdition is a free data retrieval call binding the contract method 0x602787ed.
//
// Solidity: function tokenToEdition(uint256 _tokenId) view returns(uint256)
func (_ArtistV3 *ArtistV3Session) TokenToEdition(_tokenId *big.Int) (*big.Int, error) {
	return _ArtistV3.Contract.TokenToEdition(&_ArtistV3.CallOpts, _tokenId)
}

// TokenToEdition is a free data retrieval call binding the contract method 0x602787ed.
//
// Solidity: function tokenToEdition(uint256 _tokenId) view returns(uint256)
func (_ArtistV3 *ArtistV3CallerSession) TokenToEdition(_tokenId *big.Int) (*big.Int, error) {
	return _ArtistV3.Contract.TokenToEdition(&_ArtistV3.CallOpts, _tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_ArtistV3 *ArtistV3Caller) TokenURI(opts *bind.CallOpts, _tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _ArtistV3.contract.Call(opts, &out, "tokenURI", _tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_ArtistV3 *ArtistV3Session) TokenURI(_tokenId *big.Int) (string, error) {
	return _ArtistV3.Contract.TokenURI(&_ArtistV3.CallOpts, _tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_ArtistV3 *ArtistV3CallerSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _ArtistV3.Contract.TokenURI(&_ArtistV3.CallOpts, _tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ArtistV3 *ArtistV3Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ArtistV3.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ArtistV3 *ArtistV3Session) TotalSupply() (*big.Int, error) {
	return _ArtistV3.Contract.TotalSupply(&_ArtistV3.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ArtistV3 *ArtistV3CallerSession) TotalSupply() (*big.Int, error) {
	return _ArtistV3.Contract.TotalSupply(&_ArtistV3.CallOpts)
}

// WithdrawnForEdition is a free data retrieval call binding the contract method 0xd3bb0528.
//
// Solidity: function withdrawnForEdition(uint256 ) view returns(uint256)
func (_ArtistV3 *ArtistV3Caller) WithdrawnForEdition(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ArtistV3.contract.Call(opts, &out, "withdrawnForEdition", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawnForEdition is a free data retrieval call binding the contract method 0xd3bb0528.
//
// Solidity: function withdrawnForEdition(uint256 ) view returns(uint256)
func (_ArtistV3 *ArtistV3Session) WithdrawnForEdition(arg0 *big.Int) (*big.Int, error) {
	return _ArtistV3.Contract.WithdrawnForEdition(&_ArtistV3.CallOpts, arg0)
}

// WithdrawnForEdition is a free data retrieval call binding the contract method 0xd3bb0528.
//
// Solidity: function withdrawnForEdition(uint256 ) view returns(uint256)
func (_ArtistV3 *ArtistV3CallerSession) WithdrawnForEdition(arg0 *big.Int) (*big.Int, error) {
	return _ArtistV3.Contract.WithdrawnForEdition(&_ArtistV3.CallOpts, arg0)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ArtistV3 *ArtistV3Transactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ArtistV3.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ArtistV3 *ArtistV3Session) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ArtistV3.Contract.Approve(&_ArtistV3.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ArtistV3 *ArtistV3TransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ArtistV3.Contract.Approve(&_ArtistV3.TransactOpts, to, tokenId)
}

// BuyEdition is a paid mutator transaction binding the contract method 0xabccf3dd.
//
// Solidity: function buyEdition(uint256 _editionId, bytes _signature) payable returns()
func (_ArtistV3 *ArtistV3Transactor) BuyEdition(opts *bind.TransactOpts, _editionId *big.Int, _signature []byte) (*types.Transaction, error) {
	return _ArtistV3.contract.Transact(opts, "buyEdition", _editionId, _signature)
}

// BuyEdition is a paid mutator transaction binding the contract method 0xabccf3dd.
//
// Solidity: function buyEdition(uint256 _editionId, bytes _signature) payable returns()
func (_ArtistV3 *ArtistV3Session) BuyEdition(_editionId *big.Int, _signature []byte) (*types.Transaction, error) {
	return _ArtistV3.Contract.BuyEdition(&_ArtistV3.TransactOpts, _editionId, _signature)
}

// BuyEdition is a paid mutator transaction binding the contract method 0xabccf3dd.
//
// Solidity: function buyEdition(uint256 _editionId, bytes _signature) payable returns()
func (_ArtistV3 *ArtistV3TransactorSession) BuyEdition(_editionId *big.Int, _signature []byte) (*types.Transaction, error) {
	return _ArtistV3.Contract.BuyEdition(&_ArtistV3.TransactOpts, _editionId, _signature)
}

// CreateEdition is a paid mutator transaction binding the contract method 0x73aaf879.
//
// Solidity: function createEdition(address _fundingRecipient, uint256 _price, uint32 _quantity, uint32 _royaltyBPS, uint32 _startTime, uint32 _endTime, uint32 _permissionedQuantity, address _signerAddress) returns()
func (_ArtistV3 *ArtistV3Transactor) CreateEdition(opts *bind.TransactOpts, _fundingRecipient common.Address, _price *big.Int, _quantity uint32, _royaltyBPS uint32, _startTime uint32, _endTime uint32, _permissionedQuantity uint32, _signerAddress common.Address) (*types.Transaction, error) {
	return _ArtistV3.contract.Transact(opts, "createEdition", _fundingRecipient, _price, _quantity, _royaltyBPS, _startTime, _endTime, _permissionedQuantity, _signerAddress)
}

// CreateEdition is a paid mutator transaction binding the contract method 0x73aaf879.
//
// Solidity: function createEdition(address _fundingRecipient, uint256 _price, uint32 _quantity, uint32 _royaltyBPS, uint32 _startTime, uint32 _endTime, uint32 _permissionedQuantity, address _signerAddress) returns()
func (_ArtistV3 *ArtistV3Session) CreateEdition(_fundingRecipient common.Address, _price *big.Int, _quantity uint32, _royaltyBPS uint32, _startTime uint32, _endTime uint32, _permissionedQuantity uint32, _signerAddress common.Address) (*types.Transaction, error) {
	return _ArtistV3.Contract.CreateEdition(&_ArtistV3.TransactOpts, _fundingRecipient, _price, _quantity, _royaltyBPS, _startTime, _endTime, _permissionedQuantity, _signerAddress)
}

// CreateEdition is a paid mutator transaction binding the contract method 0x73aaf879.
//
// Solidity: function createEdition(address _fundingRecipient, uint256 _price, uint32 _quantity, uint32 _royaltyBPS, uint32 _startTime, uint32 _endTime, uint32 _permissionedQuantity, address _signerAddress) returns()
func (_ArtistV3 *ArtistV3TransactorSession) CreateEdition(_fundingRecipient common.Address, _price *big.Int, _quantity uint32, _royaltyBPS uint32, _startTime uint32, _endTime uint32, _permissionedQuantity uint32, _signerAddress common.Address) (*types.Transaction, error) {
	return _ArtistV3.Contract.CreateEdition(&_ArtistV3.TransactOpts, _fundingRecipient, _price, _quantity, _royaltyBPS, _startTime, _endTime, _permissionedQuantity, _signerAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xabfc83a0.
//
// Solidity: function initialize(address _owner, uint256 _artistId, string _name, string _symbol, string _baseURI) returns()
func (_ArtistV3 *ArtistV3Transactor) Initialize(opts *bind.TransactOpts, _owner common.Address, _artistId *big.Int, _name string, _symbol string, _baseURI string) (*types.Transaction, error) {
	return _ArtistV3.contract.Transact(opts, "initialize", _owner, _artistId, _name, _symbol, _baseURI)
}

// Initialize is a paid mutator transaction binding the contract method 0xabfc83a0.
//
// Solidity: function initialize(address _owner, uint256 _artistId, string _name, string _symbol, string _baseURI) returns()
func (_ArtistV3 *ArtistV3Session) Initialize(_owner common.Address, _artistId *big.Int, _name string, _symbol string, _baseURI string) (*types.Transaction, error) {
	return _ArtistV3.Contract.Initialize(&_ArtistV3.TransactOpts, _owner, _artistId, _name, _symbol, _baseURI)
}

// Initialize is a paid mutator transaction binding the contract method 0xabfc83a0.
//
// Solidity: function initialize(address _owner, uint256 _artistId, string _name, string _symbol, string _baseURI) returns()
func (_ArtistV3 *ArtistV3TransactorSession) Initialize(_owner common.Address, _artistId *big.Int, _name string, _symbol string, _baseURI string) (*types.Transaction, error) {
	return _ArtistV3.Contract.Initialize(&_ArtistV3.TransactOpts, _owner, _artistId, _name, _symbol, _baseURI)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ArtistV3 *ArtistV3Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArtistV3.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ArtistV3 *ArtistV3Session) RenounceOwnership() (*types.Transaction, error) {
	return _ArtistV3.Contract.RenounceOwnership(&_ArtistV3.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ArtistV3 *ArtistV3TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ArtistV3.Contract.RenounceOwnership(&_ArtistV3.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ArtistV3 *ArtistV3Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ArtistV3.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ArtistV3 *ArtistV3Session) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ArtistV3.Contract.SafeTransferFrom(&_ArtistV3.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ArtistV3 *ArtistV3TransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ArtistV3.Contract.SafeTransferFrom(&_ArtistV3.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_ArtistV3 *ArtistV3Transactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ArtistV3.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_ArtistV3 *ArtistV3Session) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ArtistV3.Contract.SafeTransferFrom0(&_ArtistV3.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_ArtistV3 *ArtistV3TransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ArtistV3.Contract.SafeTransferFrom0(&_ArtistV3.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ArtistV3 *ArtistV3Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _ArtistV3.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ArtistV3 *ArtistV3Session) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ArtistV3.Contract.SetApprovalForAll(&_ArtistV3.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ArtistV3 *ArtistV3TransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ArtistV3.Contract.SetApprovalForAll(&_ArtistV3.TransactOpts, operator, approved)
}

// SetEndTime is a paid mutator transaction binding the contract method 0xbb314ca1.
//
// Solidity: function setEndTime(uint256 _editionId, uint32 _endTime) returns()
func (_ArtistV3 *ArtistV3Transactor) SetEndTime(opts *bind.TransactOpts, _editionId *big.Int, _endTime uint32) (*types.Transaction, error) {
	return _ArtistV3.contract.Transact(opts, "setEndTime", _editionId, _endTime)
}

// SetEndTime is a paid mutator transaction binding the contract method 0xbb314ca1.
//
// Solidity: function setEndTime(uint256 _editionId, uint32 _endTime) returns()
func (_ArtistV3 *ArtistV3Session) SetEndTime(_editionId *big.Int, _endTime uint32) (*types.Transaction, error) {
	return _ArtistV3.Contract.SetEndTime(&_ArtistV3.TransactOpts, _editionId, _endTime)
}

// SetEndTime is a paid mutator transaction binding the contract method 0xbb314ca1.
//
// Solidity: function setEndTime(uint256 _editionId, uint32 _endTime) returns()
func (_ArtistV3 *ArtistV3TransactorSession) SetEndTime(_editionId *big.Int, _endTime uint32) (*types.Transaction, error) {
	return _ArtistV3.Contract.SetEndTime(&_ArtistV3.TransactOpts, _editionId, _endTime)
}

// SetPermissionedQuantity is a paid mutator transaction binding the contract method 0x52e25bf2.
//
// Solidity: function setPermissionedQuantity(uint256 _editionId, uint32 _permissionedQuantity) returns()
func (_ArtistV3 *ArtistV3Transactor) SetPermissionedQuantity(opts *bind.TransactOpts, _editionId *big.Int, _permissionedQuantity uint32) (*types.Transaction, error) {
	return _ArtistV3.contract.Transact(opts, "setPermissionedQuantity", _editionId, _permissionedQuantity)
}

// SetPermissionedQuantity is a paid mutator transaction binding the contract method 0x52e25bf2.
//
// Solidity: function setPermissionedQuantity(uint256 _editionId, uint32 _permissionedQuantity) returns()
func (_ArtistV3 *ArtistV3Session) SetPermissionedQuantity(_editionId *big.Int, _permissionedQuantity uint32) (*types.Transaction, error) {
	return _ArtistV3.Contract.SetPermissionedQuantity(&_ArtistV3.TransactOpts, _editionId, _permissionedQuantity)
}

// SetPermissionedQuantity is a paid mutator transaction binding the contract method 0x52e25bf2.
//
// Solidity: function setPermissionedQuantity(uint256 _editionId, uint32 _permissionedQuantity) returns()
func (_ArtistV3 *ArtistV3TransactorSession) SetPermissionedQuantity(_editionId *big.Int, _permissionedQuantity uint32) (*types.Transaction, error) {
	return _ArtistV3.Contract.SetPermissionedQuantity(&_ArtistV3.TransactOpts, _editionId, _permissionedQuantity)
}

// SetSignerAddress is a paid mutator transaction binding the contract method 0x56dee996.
//
// Solidity: function setSignerAddress(uint256 _editionId, address _newSignerAddress) returns()
func (_ArtistV3 *ArtistV3Transactor) SetSignerAddress(opts *bind.TransactOpts, _editionId *big.Int, _newSignerAddress common.Address) (*types.Transaction, error) {
	return _ArtistV3.contract.Transact(opts, "setSignerAddress", _editionId, _newSignerAddress)
}

// SetSignerAddress is a paid mutator transaction binding the contract method 0x56dee996.
//
// Solidity: function setSignerAddress(uint256 _editionId, address _newSignerAddress) returns()
func (_ArtistV3 *ArtistV3Session) SetSignerAddress(_editionId *big.Int, _newSignerAddress common.Address) (*types.Transaction, error) {
	return _ArtistV3.Contract.SetSignerAddress(&_ArtistV3.TransactOpts, _editionId, _newSignerAddress)
}

// SetSignerAddress is a paid mutator transaction binding the contract method 0x56dee996.
//
// Solidity: function setSignerAddress(uint256 _editionId, address _newSignerAddress) returns()
func (_ArtistV3 *ArtistV3TransactorSession) SetSignerAddress(_editionId *big.Int, _newSignerAddress common.Address) (*types.Transaction, error) {
	return _ArtistV3.Contract.SetSignerAddress(&_ArtistV3.TransactOpts, _editionId, _newSignerAddress)
}

// SetStartTime is a paid mutator transaction binding the contract method 0xfbab9e04.
//
// Solidity: function setStartTime(uint256 _editionId, uint32 _startTime) returns()
func (_ArtistV3 *ArtistV3Transactor) SetStartTime(opts *bind.TransactOpts, _editionId *big.Int, _startTime uint32) (*types.Transaction, error) {
	return _ArtistV3.contract.Transact(opts, "setStartTime", _editionId, _startTime)
}

// SetStartTime is a paid mutator transaction binding the contract method 0xfbab9e04.
//
// Solidity: function setStartTime(uint256 _editionId, uint32 _startTime) returns()
func (_ArtistV3 *ArtistV3Session) SetStartTime(_editionId *big.Int, _startTime uint32) (*types.Transaction, error) {
	return _ArtistV3.Contract.SetStartTime(&_ArtistV3.TransactOpts, _editionId, _startTime)
}

// SetStartTime is a paid mutator transaction binding the contract method 0xfbab9e04.
//
// Solidity: function setStartTime(uint256 _editionId, uint32 _startTime) returns()
func (_ArtistV3 *ArtistV3TransactorSession) SetStartTime(_editionId *big.Int, _startTime uint32) (*types.Transaction, error) {
	return _ArtistV3.Contract.SetStartTime(&_ArtistV3.TransactOpts, _editionId, _startTime)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ArtistV3 *ArtistV3Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ArtistV3.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ArtistV3 *ArtistV3Session) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ArtistV3.Contract.TransferFrom(&_ArtistV3.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ArtistV3 *ArtistV3TransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ArtistV3.Contract.TransferFrom(&_ArtistV3.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ArtistV3 *ArtistV3Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ArtistV3.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ArtistV3 *ArtistV3Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ArtistV3.Contract.TransferOwnership(&_ArtistV3.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ArtistV3 *ArtistV3TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ArtistV3.Contract.TransferOwnership(&_ArtistV3.TransactOpts, newOwner)
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0x155dd5ee.
//
// Solidity: function withdrawFunds(uint256 _editionId) returns()
func (_ArtistV3 *ArtistV3Transactor) WithdrawFunds(opts *bind.TransactOpts, _editionId *big.Int) (*types.Transaction, error) {
	return _ArtistV3.contract.Transact(opts, "withdrawFunds", _editionId)
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0x155dd5ee.
//
// Solidity: function withdrawFunds(uint256 _editionId) returns()
func (_ArtistV3 *ArtistV3Session) WithdrawFunds(_editionId *big.Int) (*types.Transaction, error) {
	return _ArtistV3.Contract.WithdrawFunds(&_ArtistV3.TransactOpts, _editionId)
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0x155dd5ee.
//
// Solidity: function withdrawFunds(uint256 _editionId) returns()
func (_ArtistV3 *ArtistV3TransactorSession) WithdrawFunds(_editionId *big.Int) (*types.Transaction, error) {
	return _ArtistV3.Contract.WithdrawFunds(&_ArtistV3.TransactOpts, _editionId)
}

// ArtistV3ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ArtistV3 contract.
type ArtistV3ApprovalIterator struct {
	Event *ArtistV3Approval // Event containing the contract specifics and raw log

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
func (it *ArtistV3ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArtistV3Approval)
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
		it.Event = new(ArtistV3Approval)
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
func (it *ArtistV3ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArtistV3ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArtistV3Approval represents a Approval event raised by the ArtistV3 contract.
type ArtistV3Approval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ArtistV3 *ArtistV3Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*ArtistV3ApprovalIterator, error) {

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

	logs, sub, err := _ArtistV3.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ArtistV3ApprovalIterator{contract: _ArtistV3.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ArtistV3 *ArtistV3Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ArtistV3Approval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _ArtistV3.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArtistV3Approval)
				if err := _ArtistV3.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_ArtistV3 *ArtistV3Filterer) ParseApproval(log types.Log) (*ArtistV3Approval, error) {
	event := new(ArtistV3Approval)
	if err := _ArtistV3.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArtistV3ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ArtistV3 contract.
type ArtistV3ApprovalForAllIterator struct {
	Event *ArtistV3ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ArtistV3ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArtistV3ApprovalForAll)
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
		it.Event = new(ArtistV3ApprovalForAll)
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
func (it *ArtistV3ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArtistV3ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArtistV3ApprovalForAll represents a ApprovalForAll event raised by the ArtistV3 contract.
type ArtistV3ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ArtistV3 *ArtistV3Filterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ArtistV3ApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ArtistV3.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ArtistV3ApprovalForAllIterator{contract: _ArtistV3.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ArtistV3 *ArtistV3Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ArtistV3ApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ArtistV3.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArtistV3ApprovalForAll)
				if err := _ArtistV3.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_ArtistV3 *ArtistV3Filterer) ParseApprovalForAll(log types.Log) (*ArtistV3ApprovalForAll, error) {
	event := new(ArtistV3ApprovalForAll)
	if err := _ArtistV3.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArtistV3AuctionTimeSetIterator is returned from FilterAuctionTimeSet and is used to iterate over the raw logs and unpacked data for AuctionTimeSet events raised by the ArtistV3 contract.
type ArtistV3AuctionTimeSetIterator struct {
	Event *ArtistV3AuctionTimeSet // Event containing the contract specifics and raw log

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
func (it *ArtistV3AuctionTimeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArtistV3AuctionTimeSet)
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
		it.Event = new(ArtistV3AuctionTimeSet)
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
func (it *ArtistV3AuctionTimeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArtistV3AuctionTimeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArtistV3AuctionTimeSet represents a AuctionTimeSet event raised by the ArtistV3 contract.
type ArtistV3AuctionTimeSet struct {
	TimeType  uint8
	EditionId *big.Int
	NewTime   uint32
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAuctionTimeSet is a free log retrieval operation binding the contract event 0x494264d1227744d2d86690c5355813b007a13955626b261c2e02901a73f6f90c.
//
// Solidity: event AuctionTimeSet(uint8 timeType, uint256 editionId, uint32 indexed newTime)
func (_ArtistV3 *ArtistV3Filterer) FilterAuctionTimeSet(opts *bind.FilterOpts, newTime []uint32) (*ArtistV3AuctionTimeSetIterator, error) {

	var newTimeRule []interface{}
	for _, newTimeItem := range newTime {
		newTimeRule = append(newTimeRule, newTimeItem)
	}

	logs, sub, err := _ArtistV3.contract.FilterLogs(opts, "AuctionTimeSet", newTimeRule)
	if err != nil {
		return nil, err
	}
	return &ArtistV3AuctionTimeSetIterator{contract: _ArtistV3.contract, event: "AuctionTimeSet", logs: logs, sub: sub}, nil
}

// WatchAuctionTimeSet is a free log subscription operation binding the contract event 0x494264d1227744d2d86690c5355813b007a13955626b261c2e02901a73f6f90c.
//
// Solidity: event AuctionTimeSet(uint8 timeType, uint256 editionId, uint32 indexed newTime)
func (_ArtistV3 *ArtistV3Filterer) WatchAuctionTimeSet(opts *bind.WatchOpts, sink chan<- *ArtistV3AuctionTimeSet, newTime []uint32) (event.Subscription, error) {

	var newTimeRule []interface{}
	for _, newTimeItem := range newTime {
		newTimeRule = append(newTimeRule, newTimeItem)
	}

	logs, sub, err := _ArtistV3.contract.WatchLogs(opts, "AuctionTimeSet", newTimeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArtistV3AuctionTimeSet)
				if err := _ArtistV3.contract.UnpackLog(event, "AuctionTimeSet", log); err != nil {
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

// ParseAuctionTimeSet is a log parse operation binding the contract event 0x494264d1227744d2d86690c5355813b007a13955626b261c2e02901a73f6f90c.
//
// Solidity: event AuctionTimeSet(uint8 timeType, uint256 editionId, uint32 indexed newTime)
func (_ArtistV3 *ArtistV3Filterer) ParseAuctionTimeSet(log types.Log) (*ArtistV3AuctionTimeSet, error) {
	event := new(ArtistV3AuctionTimeSet)
	if err := _ArtistV3.contract.UnpackLog(event, "AuctionTimeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArtistV3EditionCreatedIterator is returned from FilterEditionCreated and is used to iterate over the raw logs and unpacked data for EditionCreated events raised by the ArtistV3 contract.
type ArtistV3EditionCreatedIterator struct {
	Event *ArtistV3EditionCreated // Event containing the contract specifics and raw log

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
func (it *ArtistV3EditionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArtistV3EditionCreated)
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
		it.Event = new(ArtistV3EditionCreated)
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
func (it *ArtistV3EditionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArtistV3EditionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArtistV3EditionCreated represents a EditionCreated event raised by the ArtistV3 contract.
type ArtistV3EditionCreated struct {
	EditionId            *big.Int
	FundingRecipient     common.Address
	Price                *big.Int
	Quantity             uint32
	RoyaltyBPS           uint32
	StartTime            uint32
	EndTime              uint32
	PermissionedQuantity uint32
	SignerAddress        common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterEditionCreated is a free log retrieval operation binding the contract event 0xb56f9ba6a8af17a142f8ad372c6fc283e81d8c6193b86afe7f6ca901acd698f3.
//
// Solidity: event EditionCreated(uint256 indexed editionId, address fundingRecipient, uint256 price, uint32 quantity, uint32 royaltyBPS, uint32 startTime, uint32 endTime, uint32 permissionedQuantity, address signerAddress)
func (_ArtistV3 *ArtistV3Filterer) FilterEditionCreated(opts *bind.FilterOpts, editionId []*big.Int) (*ArtistV3EditionCreatedIterator, error) {

	var editionIdRule []interface{}
	for _, editionIdItem := range editionId {
		editionIdRule = append(editionIdRule, editionIdItem)
	}

	logs, sub, err := _ArtistV3.contract.FilterLogs(opts, "EditionCreated", editionIdRule)
	if err != nil {
		return nil, err
	}
	return &ArtistV3EditionCreatedIterator{contract: _ArtistV3.contract, event: "EditionCreated", logs: logs, sub: sub}, nil
}

// WatchEditionCreated is a free log subscription operation binding the contract event 0xb56f9ba6a8af17a142f8ad372c6fc283e81d8c6193b86afe7f6ca901acd698f3.
//
// Solidity: event EditionCreated(uint256 indexed editionId, address fundingRecipient, uint256 price, uint32 quantity, uint32 royaltyBPS, uint32 startTime, uint32 endTime, uint32 permissionedQuantity, address signerAddress)
func (_ArtistV3 *ArtistV3Filterer) WatchEditionCreated(opts *bind.WatchOpts, sink chan<- *ArtistV3EditionCreated, editionId []*big.Int) (event.Subscription, error) {

	var editionIdRule []interface{}
	for _, editionIdItem := range editionId {
		editionIdRule = append(editionIdRule, editionIdItem)
	}

	logs, sub, err := _ArtistV3.contract.WatchLogs(opts, "EditionCreated", editionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArtistV3EditionCreated)
				if err := _ArtistV3.contract.UnpackLog(event, "EditionCreated", log); err != nil {
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

// ParseEditionCreated is a log parse operation binding the contract event 0xb56f9ba6a8af17a142f8ad372c6fc283e81d8c6193b86afe7f6ca901acd698f3.
//
// Solidity: event EditionCreated(uint256 indexed editionId, address fundingRecipient, uint256 price, uint32 quantity, uint32 royaltyBPS, uint32 startTime, uint32 endTime, uint32 permissionedQuantity, address signerAddress)
func (_ArtistV3 *ArtistV3Filterer) ParseEditionCreated(log types.Log) (*ArtistV3EditionCreated, error) {
	event := new(ArtistV3EditionCreated)
	if err := _ArtistV3.contract.UnpackLog(event, "EditionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArtistV3EditionPurchasedIterator is returned from FilterEditionPurchased and is used to iterate over the raw logs and unpacked data for EditionPurchased events raised by the ArtistV3 contract.
type ArtistV3EditionPurchasedIterator struct {
	Event *ArtistV3EditionPurchased // Event containing the contract specifics and raw log

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
func (it *ArtistV3EditionPurchasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArtistV3EditionPurchased)
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
		it.Event = new(ArtistV3EditionPurchased)
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
func (it *ArtistV3EditionPurchasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArtistV3EditionPurchasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArtistV3EditionPurchased represents a EditionPurchased event raised by the ArtistV3 contract.
type ArtistV3EditionPurchased struct {
	EditionId *big.Int
	TokenId   *big.Int
	NumSold   uint32
	Buyer     common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEditionPurchased is a free log retrieval operation binding the contract event 0xe38cb07a52e5d88a83de7c9d29c2841118103e462d20f8c526b35872f9977785.
//
// Solidity: event EditionPurchased(uint256 indexed editionId, uint256 indexed tokenId, uint32 numSold, address indexed buyer)
func (_ArtistV3 *ArtistV3Filterer) FilterEditionPurchased(opts *bind.FilterOpts, editionId []*big.Int, tokenId []*big.Int, buyer []common.Address) (*ArtistV3EditionPurchasedIterator, error) {

	var editionIdRule []interface{}
	for _, editionIdItem := range editionId {
		editionIdRule = append(editionIdRule, editionIdItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _ArtistV3.contract.FilterLogs(opts, "EditionPurchased", editionIdRule, tokenIdRule, buyerRule)
	if err != nil {
		return nil, err
	}
	return &ArtistV3EditionPurchasedIterator{contract: _ArtistV3.contract, event: "EditionPurchased", logs: logs, sub: sub}, nil
}

// WatchEditionPurchased is a free log subscription operation binding the contract event 0xe38cb07a52e5d88a83de7c9d29c2841118103e462d20f8c526b35872f9977785.
//
// Solidity: event EditionPurchased(uint256 indexed editionId, uint256 indexed tokenId, uint32 numSold, address indexed buyer)
func (_ArtistV3 *ArtistV3Filterer) WatchEditionPurchased(opts *bind.WatchOpts, sink chan<- *ArtistV3EditionPurchased, editionId []*big.Int, tokenId []*big.Int, buyer []common.Address) (event.Subscription, error) {

	var editionIdRule []interface{}
	for _, editionIdItem := range editionId {
		editionIdRule = append(editionIdRule, editionIdItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _ArtistV3.contract.WatchLogs(opts, "EditionPurchased", editionIdRule, tokenIdRule, buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArtistV3EditionPurchased)
				if err := _ArtistV3.contract.UnpackLog(event, "EditionPurchased", log); err != nil {
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

// ParseEditionPurchased is a log parse operation binding the contract event 0xe38cb07a52e5d88a83de7c9d29c2841118103e462d20f8c526b35872f9977785.
//
// Solidity: event EditionPurchased(uint256 indexed editionId, uint256 indexed tokenId, uint32 numSold, address indexed buyer)
func (_ArtistV3 *ArtistV3Filterer) ParseEditionPurchased(log types.Log) (*ArtistV3EditionPurchased, error) {
	event := new(ArtistV3EditionPurchased)
	if err := _ArtistV3.contract.UnpackLog(event, "EditionPurchased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArtistV3OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ArtistV3 contract.
type ArtistV3OwnershipTransferredIterator struct {
	Event *ArtistV3OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ArtistV3OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArtistV3OwnershipTransferred)
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
		it.Event = new(ArtistV3OwnershipTransferred)
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
func (it *ArtistV3OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArtistV3OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArtistV3OwnershipTransferred represents a OwnershipTransferred event raised by the ArtistV3 contract.
type ArtistV3OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ArtistV3 *ArtistV3Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ArtistV3OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ArtistV3.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ArtistV3OwnershipTransferredIterator{contract: _ArtistV3.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ArtistV3 *ArtistV3Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ArtistV3OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ArtistV3.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArtistV3OwnershipTransferred)
				if err := _ArtistV3.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ArtistV3 *ArtistV3Filterer) ParseOwnershipTransferred(log types.Log) (*ArtistV3OwnershipTransferred, error) {
	event := new(ArtistV3OwnershipTransferred)
	if err := _ArtistV3.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArtistV3PermissionedQuantitySetIterator is returned from FilterPermissionedQuantitySet and is used to iterate over the raw logs and unpacked data for PermissionedQuantitySet events raised by the ArtistV3 contract.
type ArtistV3PermissionedQuantitySetIterator struct {
	Event *ArtistV3PermissionedQuantitySet // Event containing the contract specifics and raw log

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
func (it *ArtistV3PermissionedQuantitySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArtistV3PermissionedQuantitySet)
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
		it.Event = new(ArtistV3PermissionedQuantitySet)
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
func (it *ArtistV3PermissionedQuantitySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArtistV3PermissionedQuantitySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArtistV3PermissionedQuantitySet represents a PermissionedQuantitySet event raised by the ArtistV3 contract.
type ArtistV3PermissionedQuantitySet struct {
	EditionId            *big.Int
	PermissionedQuantity uint32
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterPermissionedQuantitySet is a free log retrieval operation binding the contract event 0x1de856603e1e748b6f6726e7a51520fe7c63e9b801b8e4b2f8de1a02ae0a8dae.
//
// Solidity: event PermissionedQuantitySet(uint256 editionId, uint32 permissionedQuantity)
func (_ArtistV3 *ArtistV3Filterer) FilterPermissionedQuantitySet(opts *bind.FilterOpts) (*ArtistV3PermissionedQuantitySetIterator, error) {

	logs, sub, err := _ArtistV3.contract.FilterLogs(opts, "PermissionedQuantitySet")
	if err != nil {
		return nil, err
	}
	return &ArtistV3PermissionedQuantitySetIterator{contract: _ArtistV3.contract, event: "PermissionedQuantitySet", logs: logs, sub: sub}, nil
}

// WatchPermissionedQuantitySet is a free log subscription operation binding the contract event 0x1de856603e1e748b6f6726e7a51520fe7c63e9b801b8e4b2f8de1a02ae0a8dae.
//
// Solidity: event PermissionedQuantitySet(uint256 editionId, uint32 permissionedQuantity)
func (_ArtistV3 *ArtistV3Filterer) WatchPermissionedQuantitySet(opts *bind.WatchOpts, sink chan<- *ArtistV3PermissionedQuantitySet) (event.Subscription, error) {

	logs, sub, err := _ArtistV3.contract.WatchLogs(opts, "PermissionedQuantitySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArtistV3PermissionedQuantitySet)
				if err := _ArtistV3.contract.UnpackLog(event, "PermissionedQuantitySet", log); err != nil {
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

// ParsePermissionedQuantitySet is a log parse operation binding the contract event 0x1de856603e1e748b6f6726e7a51520fe7c63e9b801b8e4b2f8de1a02ae0a8dae.
//
// Solidity: event PermissionedQuantitySet(uint256 editionId, uint32 permissionedQuantity)
func (_ArtistV3 *ArtistV3Filterer) ParsePermissionedQuantitySet(log types.Log) (*ArtistV3PermissionedQuantitySet, error) {
	event := new(ArtistV3PermissionedQuantitySet)
	if err := _ArtistV3.contract.UnpackLog(event, "PermissionedQuantitySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArtistV3SignerAddressSetIterator is returned from FilterSignerAddressSet and is used to iterate over the raw logs and unpacked data for SignerAddressSet events raised by the ArtistV3 contract.
type ArtistV3SignerAddressSetIterator struct {
	Event *ArtistV3SignerAddressSet // Event containing the contract specifics and raw log

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
func (it *ArtistV3SignerAddressSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArtistV3SignerAddressSet)
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
		it.Event = new(ArtistV3SignerAddressSet)
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
func (it *ArtistV3SignerAddressSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArtistV3SignerAddressSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArtistV3SignerAddressSet represents a SignerAddressSet event raised by the ArtistV3 contract.
type ArtistV3SignerAddressSet struct {
	EditionId     *big.Int
	SignerAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSignerAddressSet is a free log retrieval operation binding the contract event 0x73b6ab337f7463563f035a0b9f885872e16ad1b5043b679cdf802cfcbaa3a534.
//
// Solidity: event SignerAddressSet(uint256 editionId, address indexed signerAddress)
func (_ArtistV3 *ArtistV3Filterer) FilterSignerAddressSet(opts *bind.FilterOpts, signerAddress []common.Address) (*ArtistV3SignerAddressSetIterator, error) {

	var signerAddressRule []interface{}
	for _, signerAddressItem := range signerAddress {
		signerAddressRule = append(signerAddressRule, signerAddressItem)
	}

	logs, sub, err := _ArtistV3.contract.FilterLogs(opts, "SignerAddressSet", signerAddressRule)
	if err != nil {
		return nil, err
	}
	return &ArtistV3SignerAddressSetIterator{contract: _ArtistV3.contract, event: "SignerAddressSet", logs: logs, sub: sub}, nil
}

// WatchSignerAddressSet is a free log subscription operation binding the contract event 0x73b6ab337f7463563f035a0b9f885872e16ad1b5043b679cdf802cfcbaa3a534.
//
// Solidity: event SignerAddressSet(uint256 editionId, address indexed signerAddress)
func (_ArtistV3 *ArtistV3Filterer) WatchSignerAddressSet(opts *bind.WatchOpts, sink chan<- *ArtistV3SignerAddressSet, signerAddress []common.Address) (event.Subscription, error) {

	var signerAddressRule []interface{}
	for _, signerAddressItem := range signerAddress {
		signerAddressRule = append(signerAddressRule, signerAddressItem)
	}

	logs, sub, err := _ArtistV3.contract.WatchLogs(opts, "SignerAddressSet", signerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArtistV3SignerAddressSet)
				if err := _ArtistV3.contract.UnpackLog(event, "SignerAddressSet", log); err != nil {
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

// ParseSignerAddressSet is a log parse operation binding the contract event 0x73b6ab337f7463563f035a0b9f885872e16ad1b5043b679cdf802cfcbaa3a534.
//
// Solidity: event SignerAddressSet(uint256 editionId, address indexed signerAddress)
func (_ArtistV3 *ArtistV3Filterer) ParseSignerAddressSet(log types.Log) (*ArtistV3SignerAddressSet, error) {
	event := new(ArtistV3SignerAddressSet)
	if err := _ArtistV3.contract.UnpackLog(event, "SignerAddressSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArtistV3TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ArtistV3 contract.
type ArtistV3TransferIterator struct {
	Event *ArtistV3Transfer // Event containing the contract specifics and raw log

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
func (it *ArtistV3TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArtistV3Transfer)
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
		it.Event = new(ArtistV3Transfer)
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
func (it *ArtistV3TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArtistV3TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArtistV3Transfer represents a Transfer event raised by the ArtistV3 contract.
type ArtistV3Transfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ArtistV3 *ArtistV3Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*ArtistV3TransferIterator, error) {

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

	logs, sub, err := _ArtistV3.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ArtistV3TransferIterator{contract: _ArtistV3.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ArtistV3 *ArtistV3Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ArtistV3Transfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _ArtistV3.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArtistV3Transfer)
				if err := _ArtistV3.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_ArtistV3 *ArtistV3Filterer) ParseTransfer(log types.Log) (*ArtistV3Transfer, error) {
	event := new(ArtistV3Transfer)
	if err := _ArtistV3.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
