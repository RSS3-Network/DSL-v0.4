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
	_ = abi.ConvertType
)

// ArtistMetaData contains all meta data concerning the Artist contract.
var ArtistMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"editionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"fundingRecipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"quantity\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"royaltyBPS\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"startTime\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"endTime\",\"type\":\"uint32\"}],\"name\":\"EditionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"editionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"numSold\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"}],\"name\":\"EditionPurchased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_editionId\",\"type\":\"uint256\"}],\"name\":\"buyEdition\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_fundingRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_quantity\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_royaltyBPS\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_startTime\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_endTime\",\"type\":\"uint32\"}],\"name\":\"createEdition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"depositedForEdition\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"editions\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"fundingRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"numSold\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"quantity\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"royaltyBPS\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"startTime\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"endTime\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_editionId\",\"type\":\"uint256\"}],\"name\":\"getOwnersOfEdition\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_editionId\",\"type\":\"uint256\"}],\"name\":\"getTokenIdsOfEdition\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_artistId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_baseURI\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_editionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_salePrice\",\"type\":\"uint256\"}],\"name\":\"royaltyInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"fundingRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"royaltyAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_editionId\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_endTime\",\"type\":\"uint32\"}],\"name\":\"setEndTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_editionId\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_startTime\",\"type\":\"uint32\"}],\"name\":\"setStartTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenToEdition\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_editionId\",\"type\":\"uint256\"}],\"name\":\"withdrawFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"withdrawnForEdition\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ArtistABI is the input ABI used to generate the binding from.
// Deprecated: Use ArtistMetaData.ABI instead.
var ArtistABI = ArtistMetaData.ABI

// Artist is an auto generated Go binding around an Ethereum contract.
type Artist struct {
	ArtistCaller     // Read-only binding to the contract
	ArtistTransactor // Write-only binding to the contract
	ArtistFilterer   // Log filterer for contract events
}

// ArtistCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArtistCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArtistTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArtistTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArtistFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArtistFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArtistSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArtistSession struct {
	Contract     *Artist           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArtistCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArtistCallerSession struct {
	Contract *ArtistCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ArtistTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArtistTransactorSession struct {
	Contract     *ArtistTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArtistRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArtistRaw struct {
	Contract *Artist // Generic contract binding to access the raw methods on
}

// ArtistCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArtistCallerRaw struct {
	Contract *ArtistCaller // Generic read-only contract binding to access the raw methods on
}

// ArtistTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArtistTransactorRaw struct {
	Contract *ArtistTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArtist creates a new instance of Artist, bound to a specific deployed contract.
func NewArtist(address common.Address, backend bind.ContractBackend) (*Artist, error) {
	contract, err := bindArtist(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Artist{ArtistCaller: ArtistCaller{contract: contract}, ArtistTransactor: ArtistTransactor{contract: contract}, ArtistFilterer: ArtistFilterer{contract: contract}}, nil
}

// NewArtistCaller creates a new read-only instance of Artist, bound to a specific deployed contract.
func NewArtistCaller(address common.Address, caller bind.ContractCaller) (*ArtistCaller, error) {
	contract, err := bindArtist(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArtistCaller{contract: contract}, nil
}

// NewArtistTransactor creates a new write-only instance of Artist, bound to a specific deployed contract.
func NewArtistTransactor(address common.Address, transactor bind.ContractTransactor) (*ArtistTransactor, error) {
	contract, err := bindArtist(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArtistTransactor{contract: contract}, nil
}

// NewArtistFilterer creates a new log filterer instance of Artist, bound to a specific deployed contract.
func NewArtistFilterer(address common.Address, filterer bind.ContractFilterer) (*ArtistFilterer, error) {
	contract, err := bindArtist(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArtistFilterer{contract: contract}, nil
}

// bindArtist binds a generic wrapper to an already deployed contract.
func bindArtist(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ArtistMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Artist *ArtistRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Artist.Contract.ArtistCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Artist *ArtistRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Artist.Contract.ArtistTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Artist *ArtistRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Artist.Contract.ArtistTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Artist *ArtistCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Artist.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Artist *ArtistTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Artist.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Artist *ArtistTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Artist.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Artist *ArtistCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Artist.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Artist *ArtistSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Artist.Contract.BalanceOf(&_Artist.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Artist *ArtistCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Artist.Contract.BalanceOf(&_Artist.CallOpts, owner)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Artist *ArtistCaller) ContractURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Artist.contract.Call(opts, &out, "contractURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Artist *ArtistSession) ContractURI() (string, error) {
	return _Artist.Contract.ContractURI(&_Artist.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Artist *ArtistCallerSession) ContractURI() (string, error) {
	return _Artist.Contract.ContractURI(&_Artist.CallOpts)
}

// DepositedForEdition is a free data retrieval call binding the contract method 0xe1a3d573.
//
// Solidity: function depositedForEdition(uint256 ) view returns(uint256)
func (_Artist *ArtistCaller) DepositedForEdition(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Artist.contract.Call(opts, &out, "depositedForEdition", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositedForEdition is a free data retrieval call binding the contract method 0xe1a3d573.
//
// Solidity: function depositedForEdition(uint256 ) view returns(uint256)
func (_Artist *ArtistSession) DepositedForEdition(arg0 *big.Int) (*big.Int, error) {
	return _Artist.Contract.DepositedForEdition(&_Artist.CallOpts, arg0)
}

// DepositedForEdition is a free data retrieval call binding the contract method 0xe1a3d573.
//
// Solidity: function depositedForEdition(uint256 ) view returns(uint256)
func (_Artist *ArtistCallerSession) DepositedForEdition(arg0 *big.Int) (*big.Int, error) {
	return _Artist.Contract.DepositedForEdition(&_Artist.CallOpts, arg0)
}

// Editions is a free data retrieval call binding the contract method 0x279c806e.
//
// Solidity: function editions(uint256 ) view returns(address fundingRecipient, uint256 price, uint32 numSold, uint32 quantity, uint32 royaltyBPS, uint32 startTime, uint32 endTime)
func (_Artist *ArtistCaller) Editions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	FundingRecipient common.Address
	Price            *big.Int
	NumSold          uint32
	Quantity         uint32
	RoyaltyBPS       uint32
	StartTime        uint32
	EndTime          uint32
}, error) {
	var out []interface{}
	err := _Artist.contract.Call(opts, &out, "editions", arg0)

	outstruct := new(struct {
		FundingRecipient common.Address
		Price            *big.Int
		NumSold          uint32
		Quantity         uint32
		RoyaltyBPS       uint32
		StartTime        uint32
		EndTime          uint32
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

	return *outstruct, err

}

// Editions is a free data retrieval call binding the contract method 0x279c806e.
//
// Solidity: function editions(uint256 ) view returns(address fundingRecipient, uint256 price, uint32 numSold, uint32 quantity, uint32 royaltyBPS, uint32 startTime, uint32 endTime)
func (_Artist *ArtistSession) Editions(arg0 *big.Int) (struct {
	FundingRecipient common.Address
	Price            *big.Int
	NumSold          uint32
	Quantity         uint32
	RoyaltyBPS       uint32
	StartTime        uint32
	EndTime          uint32
}, error) {
	return _Artist.Contract.Editions(&_Artist.CallOpts, arg0)
}

// Editions is a free data retrieval call binding the contract method 0x279c806e.
//
// Solidity: function editions(uint256 ) view returns(address fundingRecipient, uint256 price, uint32 numSold, uint32 quantity, uint32 royaltyBPS, uint32 startTime, uint32 endTime)
func (_Artist *ArtistCallerSession) Editions(arg0 *big.Int) (struct {
	FundingRecipient common.Address
	Price            *big.Int
	NumSold          uint32
	Quantity         uint32
	RoyaltyBPS       uint32
	StartTime        uint32
	EndTime          uint32
}, error) {
	return _Artist.Contract.Editions(&_Artist.CallOpts, arg0)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Artist *ArtistCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Artist.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Artist *ArtistSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Artist.Contract.GetApproved(&_Artist.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Artist *ArtistCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Artist.Contract.GetApproved(&_Artist.CallOpts, tokenId)
}

// GetOwnersOfEdition is a free data retrieval call binding the contract method 0x13dd2960.
//
// Solidity: function getOwnersOfEdition(uint256 _editionId) view returns(address[])
func (_Artist *ArtistCaller) GetOwnersOfEdition(opts *bind.CallOpts, _editionId *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _Artist.contract.Call(opts, &out, "getOwnersOfEdition", _editionId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOwnersOfEdition is a free data retrieval call binding the contract method 0x13dd2960.
//
// Solidity: function getOwnersOfEdition(uint256 _editionId) view returns(address[])
func (_Artist *ArtistSession) GetOwnersOfEdition(_editionId *big.Int) ([]common.Address, error) {
	return _Artist.Contract.GetOwnersOfEdition(&_Artist.CallOpts, _editionId)
}

// GetOwnersOfEdition is a free data retrieval call binding the contract method 0x13dd2960.
//
// Solidity: function getOwnersOfEdition(uint256 _editionId) view returns(address[])
func (_Artist *ArtistCallerSession) GetOwnersOfEdition(_editionId *big.Int) ([]common.Address, error) {
	return _Artist.Contract.GetOwnersOfEdition(&_Artist.CallOpts, _editionId)
}

// GetTokenIdsOfEdition is a free data retrieval call binding the contract method 0x74e79189.
//
// Solidity: function getTokenIdsOfEdition(uint256 _editionId) view returns(uint256[])
func (_Artist *ArtistCaller) GetTokenIdsOfEdition(opts *bind.CallOpts, _editionId *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Artist.contract.Call(opts, &out, "getTokenIdsOfEdition", _editionId)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetTokenIdsOfEdition is a free data retrieval call binding the contract method 0x74e79189.
//
// Solidity: function getTokenIdsOfEdition(uint256 _editionId) view returns(uint256[])
func (_Artist *ArtistSession) GetTokenIdsOfEdition(_editionId *big.Int) ([]*big.Int, error) {
	return _Artist.Contract.GetTokenIdsOfEdition(&_Artist.CallOpts, _editionId)
}

// GetTokenIdsOfEdition is a free data retrieval call binding the contract method 0x74e79189.
//
// Solidity: function getTokenIdsOfEdition(uint256 _editionId) view returns(uint256[])
func (_Artist *ArtistCallerSession) GetTokenIdsOfEdition(_editionId *big.Int) ([]*big.Int, error) {
	return _Artist.Contract.GetTokenIdsOfEdition(&_Artist.CallOpts, _editionId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Artist *ArtistCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Artist.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Artist *ArtistSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Artist.Contract.IsApprovedForAll(&_Artist.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Artist *ArtistCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Artist.Contract.IsApprovedForAll(&_Artist.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Artist *ArtistCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Artist.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Artist *ArtistSession) Name() (string, error) {
	return _Artist.Contract.Name(&_Artist.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Artist *ArtistCallerSession) Name() (string, error) {
	return _Artist.Contract.Name(&_Artist.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Artist *ArtistCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Artist.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Artist *ArtistSession) Owner() (common.Address, error) {
	return _Artist.Contract.Owner(&_Artist.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Artist *ArtistCallerSession) Owner() (common.Address, error) {
	return _Artist.Contract.Owner(&_Artist.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Artist *ArtistCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Artist.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Artist *ArtistSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Artist.Contract.OwnerOf(&_Artist.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Artist *ArtistCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Artist.Contract.OwnerOf(&_Artist.CallOpts, tokenId)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 _editionId, uint256 _salePrice) view returns(address fundingRecipient, uint256 royaltyAmount)
func (_Artist *ArtistCaller) RoyaltyInfo(opts *bind.CallOpts, _editionId *big.Int, _salePrice *big.Int) (struct {
	FundingRecipient common.Address
	RoyaltyAmount    *big.Int
}, error) {
	var out []interface{}
	err := _Artist.contract.Call(opts, &out, "royaltyInfo", _editionId, _salePrice)

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
// Solidity: function royaltyInfo(uint256 _editionId, uint256 _salePrice) view returns(address fundingRecipient, uint256 royaltyAmount)
func (_Artist *ArtistSession) RoyaltyInfo(_editionId *big.Int, _salePrice *big.Int) (struct {
	FundingRecipient common.Address
	RoyaltyAmount    *big.Int
}, error) {
	return _Artist.Contract.RoyaltyInfo(&_Artist.CallOpts, _editionId, _salePrice)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 _editionId, uint256 _salePrice) view returns(address fundingRecipient, uint256 royaltyAmount)
func (_Artist *ArtistCallerSession) RoyaltyInfo(_editionId *big.Int, _salePrice *big.Int) (struct {
	FundingRecipient common.Address
	RoyaltyAmount    *big.Int
}, error) {
	return _Artist.Contract.RoyaltyInfo(&_Artist.CallOpts, _editionId, _salePrice)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (_Artist *ArtistCaller) SupportsInterface(opts *bind.CallOpts, _interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Artist.contract.Call(opts, &out, "supportsInterface", _interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (_Artist *ArtistSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _Artist.Contract.SupportsInterface(&_Artist.CallOpts, _interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (_Artist *ArtistCallerSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _Artist.Contract.SupportsInterface(&_Artist.CallOpts, _interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Artist *ArtistCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Artist.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Artist *ArtistSession) Symbol() (string, error) {
	return _Artist.Contract.Symbol(&_Artist.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Artist *ArtistCallerSession) Symbol() (string, error) {
	return _Artist.Contract.Symbol(&_Artist.CallOpts)
}

// TokenToEdition is a free data retrieval call binding the contract method 0x602787ed.
//
// Solidity: function tokenToEdition(uint256 ) view returns(uint256)
func (_Artist *ArtistCaller) TokenToEdition(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Artist.contract.Call(opts, &out, "tokenToEdition", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenToEdition is a free data retrieval call binding the contract method 0x602787ed.
//
// Solidity: function tokenToEdition(uint256 ) view returns(uint256)
func (_Artist *ArtistSession) TokenToEdition(arg0 *big.Int) (*big.Int, error) {
	return _Artist.Contract.TokenToEdition(&_Artist.CallOpts, arg0)
}

// TokenToEdition is a free data retrieval call binding the contract method 0x602787ed.
//
// Solidity: function tokenToEdition(uint256 ) view returns(uint256)
func (_Artist *ArtistCallerSession) TokenToEdition(arg0 *big.Int) (*big.Int, error) {
	return _Artist.Contract.TokenToEdition(&_Artist.CallOpts, arg0)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_Artist *ArtistCaller) TokenURI(opts *bind.CallOpts, _tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Artist.contract.Call(opts, &out, "tokenURI", _tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_Artist *ArtistSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _Artist.Contract.TokenURI(&_Artist.CallOpts, _tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_Artist *ArtistCallerSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _Artist.Contract.TokenURI(&_Artist.CallOpts, _tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Artist *ArtistCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Artist.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Artist *ArtistSession) TotalSupply() (*big.Int, error) {
	return _Artist.Contract.TotalSupply(&_Artist.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Artist *ArtistCallerSession) TotalSupply() (*big.Int, error) {
	return _Artist.Contract.TotalSupply(&_Artist.CallOpts)
}

// WithdrawnForEdition is a free data retrieval call binding the contract method 0xd3bb0528.
//
// Solidity: function withdrawnForEdition(uint256 ) view returns(uint256)
func (_Artist *ArtistCaller) WithdrawnForEdition(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Artist.contract.Call(opts, &out, "withdrawnForEdition", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawnForEdition is a free data retrieval call binding the contract method 0xd3bb0528.
//
// Solidity: function withdrawnForEdition(uint256 ) view returns(uint256)
func (_Artist *ArtistSession) WithdrawnForEdition(arg0 *big.Int) (*big.Int, error) {
	return _Artist.Contract.WithdrawnForEdition(&_Artist.CallOpts, arg0)
}

// WithdrawnForEdition is a free data retrieval call binding the contract method 0xd3bb0528.
//
// Solidity: function withdrawnForEdition(uint256 ) view returns(uint256)
func (_Artist *ArtistCallerSession) WithdrawnForEdition(arg0 *big.Int) (*big.Int, error) {
	return _Artist.Contract.WithdrawnForEdition(&_Artist.CallOpts, arg0)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Artist *ArtistTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Artist.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Artist *ArtistSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Artist.Contract.Approve(&_Artist.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Artist *ArtistTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Artist.Contract.Approve(&_Artist.TransactOpts, to, tokenId)
}

// BuyEdition is a paid mutator transaction binding the contract method 0xbd8616ec.
//
// Solidity: function buyEdition(uint256 _editionId) payable returns()
func (_Artist *ArtistTransactor) BuyEdition(opts *bind.TransactOpts, _editionId *big.Int) (*types.Transaction, error) {
	return _Artist.contract.Transact(opts, "buyEdition", _editionId)
}

// BuyEdition is a paid mutator transaction binding the contract method 0xbd8616ec.
//
// Solidity: function buyEdition(uint256 _editionId) payable returns()
func (_Artist *ArtistSession) BuyEdition(_editionId *big.Int) (*types.Transaction, error) {
	return _Artist.Contract.BuyEdition(&_Artist.TransactOpts, _editionId)
}

// BuyEdition is a paid mutator transaction binding the contract method 0xbd8616ec.
//
// Solidity: function buyEdition(uint256 _editionId) payable returns()
func (_Artist *ArtistTransactorSession) BuyEdition(_editionId *big.Int) (*types.Transaction, error) {
	return _Artist.Contract.BuyEdition(&_Artist.TransactOpts, _editionId)
}

// CreateEdition is a paid mutator transaction binding the contract method 0x3fafef29.
//
// Solidity: function createEdition(address _fundingRecipient, uint256 _price, uint32 _quantity, uint32 _royaltyBPS, uint32 _startTime, uint32 _endTime) returns()
func (_Artist *ArtistTransactor) CreateEdition(opts *bind.TransactOpts, _fundingRecipient common.Address, _price *big.Int, _quantity uint32, _royaltyBPS uint32, _startTime uint32, _endTime uint32) (*types.Transaction, error) {
	return _Artist.contract.Transact(opts, "createEdition", _fundingRecipient, _price, _quantity, _royaltyBPS, _startTime, _endTime)
}

// CreateEdition is a paid mutator transaction binding the contract method 0x3fafef29.
//
// Solidity: function createEdition(address _fundingRecipient, uint256 _price, uint32 _quantity, uint32 _royaltyBPS, uint32 _startTime, uint32 _endTime) returns()
func (_Artist *ArtistSession) CreateEdition(_fundingRecipient common.Address, _price *big.Int, _quantity uint32, _royaltyBPS uint32, _startTime uint32, _endTime uint32) (*types.Transaction, error) {
	return _Artist.Contract.CreateEdition(&_Artist.TransactOpts, _fundingRecipient, _price, _quantity, _royaltyBPS, _startTime, _endTime)
}

// CreateEdition is a paid mutator transaction binding the contract method 0x3fafef29.
//
// Solidity: function createEdition(address _fundingRecipient, uint256 _price, uint32 _quantity, uint32 _royaltyBPS, uint32 _startTime, uint32 _endTime) returns()
func (_Artist *ArtistTransactorSession) CreateEdition(_fundingRecipient common.Address, _price *big.Int, _quantity uint32, _royaltyBPS uint32, _startTime uint32, _endTime uint32) (*types.Transaction, error) {
	return _Artist.Contract.CreateEdition(&_Artist.TransactOpts, _fundingRecipient, _price, _quantity, _royaltyBPS, _startTime, _endTime)
}

// Initialize is a paid mutator transaction binding the contract method 0xabfc83a0.
//
// Solidity: function initialize(address _owner, uint256 _artistId, string _name, string _symbol, string _baseURI) returns()
func (_Artist *ArtistTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address, _artistId *big.Int, _name string, _symbol string, _baseURI string) (*types.Transaction, error) {
	return _Artist.contract.Transact(opts, "initialize", _owner, _artistId, _name, _symbol, _baseURI)
}

// Initialize is a paid mutator transaction binding the contract method 0xabfc83a0.
//
// Solidity: function initialize(address _owner, uint256 _artistId, string _name, string _symbol, string _baseURI) returns()
func (_Artist *ArtistSession) Initialize(_owner common.Address, _artistId *big.Int, _name string, _symbol string, _baseURI string) (*types.Transaction, error) {
	return _Artist.Contract.Initialize(&_Artist.TransactOpts, _owner, _artistId, _name, _symbol, _baseURI)
}

// Initialize is a paid mutator transaction binding the contract method 0xabfc83a0.
//
// Solidity: function initialize(address _owner, uint256 _artistId, string _name, string _symbol, string _baseURI) returns()
func (_Artist *ArtistTransactorSession) Initialize(_owner common.Address, _artistId *big.Int, _name string, _symbol string, _baseURI string) (*types.Transaction, error) {
	return _Artist.Contract.Initialize(&_Artist.TransactOpts, _owner, _artistId, _name, _symbol, _baseURI)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Artist *ArtistTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Artist.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Artist *ArtistSession) RenounceOwnership() (*types.Transaction, error) {
	return _Artist.Contract.RenounceOwnership(&_Artist.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Artist *ArtistTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Artist.Contract.RenounceOwnership(&_Artist.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Artist *ArtistTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Artist.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Artist *ArtistSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Artist.Contract.SafeTransferFrom(&_Artist.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Artist *ArtistTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Artist.Contract.SafeTransferFrom(&_Artist.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Artist *ArtistTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Artist.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Artist *ArtistSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Artist.Contract.SafeTransferFrom0(&_Artist.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Artist *ArtistTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Artist.Contract.SafeTransferFrom0(&_Artist.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Artist *ArtistTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Artist.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Artist *ArtistSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Artist.Contract.SetApprovalForAll(&_Artist.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Artist *ArtistTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Artist.Contract.SetApprovalForAll(&_Artist.TransactOpts, operator, approved)
}

// SetEndTime is a paid mutator transaction binding the contract method 0xbb314ca1.
//
// Solidity: function setEndTime(uint256 _editionId, uint32 _endTime) returns()
func (_Artist *ArtistTransactor) SetEndTime(opts *bind.TransactOpts, _editionId *big.Int, _endTime uint32) (*types.Transaction, error) {
	return _Artist.contract.Transact(opts, "setEndTime", _editionId, _endTime)
}

// SetEndTime is a paid mutator transaction binding the contract method 0xbb314ca1.
//
// Solidity: function setEndTime(uint256 _editionId, uint32 _endTime) returns()
func (_Artist *ArtistSession) SetEndTime(_editionId *big.Int, _endTime uint32) (*types.Transaction, error) {
	return _Artist.Contract.SetEndTime(&_Artist.TransactOpts, _editionId, _endTime)
}

// SetEndTime is a paid mutator transaction binding the contract method 0xbb314ca1.
//
// Solidity: function setEndTime(uint256 _editionId, uint32 _endTime) returns()
func (_Artist *ArtistTransactorSession) SetEndTime(_editionId *big.Int, _endTime uint32) (*types.Transaction, error) {
	return _Artist.Contract.SetEndTime(&_Artist.TransactOpts, _editionId, _endTime)
}

// SetStartTime is a paid mutator transaction binding the contract method 0xfbab9e04.
//
// Solidity: function setStartTime(uint256 _editionId, uint32 _startTime) returns()
func (_Artist *ArtistTransactor) SetStartTime(opts *bind.TransactOpts, _editionId *big.Int, _startTime uint32) (*types.Transaction, error) {
	return _Artist.contract.Transact(opts, "setStartTime", _editionId, _startTime)
}

// SetStartTime is a paid mutator transaction binding the contract method 0xfbab9e04.
//
// Solidity: function setStartTime(uint256 _editionId, uint32 _startTime) returns()
func (_Artist *ArtistSession) SetStartTime(_editionId *big.Int, _startTime uint32) (*types.Transaction, error) {
	return _Artist.Contract.SetStartTime(&_Artist.TransactOpts, _editionId, _startTime)
}

// SetStartTime is a paid mutator transaction binding the contract method 0xfbab9e04.
//
// Solidity: function setStartTime(uint256 _editionId, uint32 _startTime) returns()
func (_Artist *ArtistTransactorSession) SetStartTime(_editionId *big.Int, _startTime uint32) (*types.Transaction, error) {
	return _Artist.Contract.SetStartTime(&_Artist.TransactOpts, _editionId, _startTime)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Artist *ArtistTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Artist.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Artist *ArtistSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Artist.Contract.TransferFrom(&_Artist.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Artist *ArtistTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Artist.Contract.TransferFrom(&_Artist.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Artist *ArtistTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Artist.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Artist *ArtistSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Artist.Contract.TransferOwnership(&_Artist.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Artist *ArtistTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Artist.Contract.TransferOwnership(&_Artist.TransactOpts, newOwner)
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0x155dd5ee.
//
// Solidity: function withdrawFunds(uint256 _editionId) returns()
func (_Artist *ArtistTransactor) WithdrawFunds(opts *bind.TransactOpts, _editionId *big.Int) (*types.Transaction, error) {
	return _Artist.contract.Transact(opts, "withdrawFunds", _editionId)
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0x155dd5ee.
//
// Solidity: function withdrawFunds(uint256 _editionId) returns()
func (_Artist *ArtistSession) WithdrawFunds(_editionId *big.Int) (*types.Transaction, error) {
	return _Artist.Contract.WithdrawFunds(&_Artist.TransactOpts, _editionId)
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0x155dd5ee.
//
// Solidity: function withdrawFunds(uint256 _editionId) returns()
func (_Artist *ArtistTransactorSession) WithdrawFunds(_editionId *big.Int) (*types.Transaction, error) {
	return _Artist.Contract.WithdrawFunds(&_Artist.TransactOpts, _editionId)
}

// ArtistApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Artist contract.
type ArtistApprovalIterator struct {
	Event *ArtistApproval // Event containing the contract specifics and raw log

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
func (it *ArtistApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArtistApproval)
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
		it.Event = new(ArtistApproval)
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
func (it *ArtistApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArtistApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArtistApproval represents a Approval event raised by the Artist contract.
type ArtistApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Artist *ArtistFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*ArtistApprovalIterator, error) {

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

	logs, sub, err := _Artist.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ArtistApprovalIterator{contract: _Artist.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Artist *ArtistFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ArtistApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Artist.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArtistApproval)
				if err := _Artist.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Artist *ArtistFilterer) ParseApproval(log types.Log) (*ArtistApproval, error) {
	event := new(ArtistApproval)
	if err := _Artist.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArtistApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Artist contract.
type ArtistApprovalForAllIterator struct {
	Event *ArtistApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ArtistApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArtistApprovalForAll)
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
		it.Event = new(ArtistApprovalForAll)
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
func (it *ArtistApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArtistApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArtistApprovalForAll represents a ApprovalForAll event raised by the Artist contract.
type ArtistApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Artist *ArtistFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ArtistApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Artist.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ArtistApprovalForAllIterator{contract: _Artist.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Artist *ArtistFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ArtistApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Artist.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArtistApprovalForAll)
				if err := _Artist.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_Artist *ArtistFilterer) ParseApprovalForAll(log types.Log) (*ArtistApprovalForAll, error) {
	event := new(ArtistApprovalForAll)
	if err := _Artist.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArtistEditionCreatedIterator is returned from FilterEditionCreated and is used to iterate over the raw logs and unpacked data for EditionCreated events raised by the Artist contract.
type ArtistEditionCreatedIterator struct {
	Event *ArtistEditionCreated // Event containing the contract specifics and raw log

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
func (it *ArtistEditionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArtistEditionCreated)
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
		it.Event = new(ArtistEditionCreated)
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
func (it *ArtistEditionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArtistEditionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArtistEditionCreated represents a EditionCreated event raised by the Artist contract.
type ArtistEditionCreated struct {
	EditionId        *big.Int
	FundingRecipient common.Address
	Price            *big.Int
	Quantity         uint32
	RoyaltyBPS       uint32
	StartTime        uint32
	EndTime          uint32
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterEditionCreated is a free log retrieval operation binding the contract event 0xb3131d7d301f8caeb40981cffc627b1fdf324b5e4a23845b61c1a6ad2a25f385.
//
// Solidity: event EditionCreated(uint256 indexed editionId, address fundingRecipient, uint256 price, uint32 quantity, uint32 royaltyBPS, uint32 startTime, uint32 endTime)
func (_Artist *ArtistFilterer) FilterEditionCreated(opts *bind.FilterOpts, editionId []*big.Int) (*ArtistEditionCreatedIterator, error) {

	var editionIdRule []interface{}
	for _, editionIdItem := range editionId {
		editionIdRule = append(editionIdRule, editionIdItem)
	}

	logs, sub, err := _Artist.contract.FilterLogs(opts, "EditionCreated", editionIdRule)
	if err != nil {
		return nil, err
	}
	return &ArtistEditionCreatedIterator{contract: _Artist.contract, event: "EditionCreated", logs: logs, sub: sub}, nil
}

// WatchEditionCreated is a free log subscription operation binding the contract event 0xb3131d7d301f8caeb40981cffc627b1fdf324b5e4a23845b61c1a6ad2a25f385.
//
// Solidity: event EditionCreated(uint256 indexed editionId, address fundingRecipient, uint256 price, uint32 quantity, uint32 royaltyBPS, uint32 startTime, uint32 endTime)
func (_Artist *ArtistFilterer) WatchEditionCreated(opts *bind.WatchOpts, sink chan<- *ArtistEditionCreated, editionId []*big.Int) (event.Subscription, error) {

	var editionIdRule []interface{}
	for _, editionIdItem := range editionId {
		editionIdRule = append(editionIdRule, editionIdItem)
	}

	logs, sub, err := _Artist.contract.WatchLogs(opts, "EditionCreated", editionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArtistEditionCreated)
				if err := _Artist.contract.UnpackLog(event, "EditionCreated", log); err != nil {
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

// ParseEditionCreated is a log parse operation binding the contract event 0xb3131d7d301f8caeb40981cffc627b1fdf324b5e4a23845b61c1a6ad2a25f385.
//
// Solidity: event EditionCreated(uint256 indexed editionId, address fundingRecipient, uint256 price, uint32 quantity, uint32 royaltyBPS, uint32 startTime, uint32 endTime)
func (_Artist *ArtistFilterer) ParseEditionCreated(log types.Log) (*ArtistEditionCreated, error) {
	event := new(ArtistEditionCreated)
	if err := _Artist.contract.UnpackLog(event, "EditionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArtistEditionPurchasedIterator is returned from FilterEditionPurchased and is used to iterate over the raw logs and unpacked data for EditionPurchased events raised by the Artist contract.
type ArtistEditionPurchasedIterator struct {
	Event *ArtistEditionPurchased // Event containing the contract specifics and raw log

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
func (it *ArtistEditionPurchasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArtistEditionPurchased)
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
		it.Event = new(ArtistEditionPurchased)
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
func (it *ArtistEditionPurchasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArtistEditionPurchasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArtistEditionPurchased represents a EditionPurchased event raised by the Artist contract.
type ArtistEditionPurchased struct {
	EditionId *big.Int
	TokenId   *big.Int
	NumSold   uint32
	Buyer     common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEditionPurchased is a free log retrieval operation binding the contract event 0xe38cb07a52e5d88a83de7c9d29c2841118103e462d20f8c526b35872f9977785.
//
// Solidity: event EditionPurchased(uint256 indexed editionId, uint256 indexed tokenId, uint32 numSold, address indexed buyer)
func (_Artist *ArtistFilterer) FilterEditionPurchased(opts *bind.FilterOpts, editionId []*big.Int, tokenId []*big.Int, buyer []common.Address) (*ArtistEditionPurchasedIterator, error) {

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

	logs, sub, err := _Artist.contract.FilterLogs(opts, "EditionPurchased", editionIdRule, tokenIdRule, buyerRule)
	if err != nil {
		return nil, err
	}
	return &ArtistEditionPurchasedIterator{contract: _Artist.contract, event: "EditionPurchased", logs: logs, sub: sub}, nil
}

// WatchEditionPurchased is a free log subscription operation binding the contract event 0xe38cb07a52e5d88a83de7c9d29c2841118103e462d20f8c526b35872f9977785.
//
// Solidity: event EditionPurchased(uint256 indexed editionId, uint256 indexed tokenId, uint32 numSold, address indexed buyer)
func (_Artist *ArtistFilterer) WatchEditionPurchased(opts *bind.WatchOpts, sink chan<- *ArtistEditionPurchased, editionId []*big.Int, tokenId []*big.Int, buyer []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Artist.contract.WatchLogs(opts, "EditionPurchased", editionIdRule, tokenIdRule, buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArtistEditionPurchased)
				if err := _Artist.contract.UnpackLog(event, "EditionPurchased", log); err != nil {
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
func (_Artist *ArtistFilterer) ParseEditionPurchased(log types.Log) (*ArtistEditionPurchased, error) {
	event := new(ArtistEditionPurchased)
	if err := _Artist.contract.UnpackLog(event, "EditionPurchased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArtistOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Artist contract.
type ArtistOwnershipTransferredIterator struct {
	Event *ArtistOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ArtistOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArtistOwnershipTransferred)
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
		it.Event = new(ArtistOwnershipTransferred)
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
func (it *ArtistOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArtistOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArtistOwnershipTransferred represents a OwnershipTransferred event raised by the Artist contract.
type ArtistOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Artist *ArtistFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ArtistOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Artist.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ArtistOwnershipTransferredIterator{contract: _Artist.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Artist *ArtistFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ArtistOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Artist.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArtistOwnershipTransferred)
				if err := _Artist.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Artist *ArtistFilterer) ParseOwnershipTransferred(log types.Log) (*ArtistOwnershipTransferred, error) {
	event := new(ArtistOwnershipTransferred)
	if err := _Artist.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArtistTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Artist contract.
type ArtistTransferIterator struct {
	Event *ArtistTransfer // Event containing the contract specifics and raw log

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
func (it *ArtistTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArtistTransfer)
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
		it.Event = new(ArtistTransfer)
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
func (it *ArtistTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArtistTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArtistTransfer represents a Transfer event raised by the Artist contract.
type ArtistTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Artist *ArtistFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*ArtistTransferIterator, error) {

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

	logs, sub, err := _Artist.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ArtistTransferIterator{contract: _Artist.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Artist *ArtistFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ArtistTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Artist.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArtistTransfer)
				if err := _Artist.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Artist *ArtistFilterer) ParseTransfer(log types.Log) (*ArtistTransfer, error) {
	event := new(ArtistTransfer)
	if err := _Artist.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
