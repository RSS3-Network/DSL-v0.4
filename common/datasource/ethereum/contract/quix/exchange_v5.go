// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package quix

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

// ExchangeV5MetaData contains all meta data concerning the ExchangeV5 contract.
var ExchangeV5MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"addresspayable\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"SellOrderFilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endPrice\",\"type\":\"uint256\"}],\"name\":\"calculateCurrentPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"cancelPreviousSellOrders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAtBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentERC20\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"addresspayable\",\"name\":\"buyer\",\"type\":\"address\"}],\"name\":\"fillSellOrder\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"getRoyaltyPayoutAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"getRoyaltyPayoutRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_newMakerWallet\",\"type\":\"address\"}],\"name\":\"setMakerWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_royaltyRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_cancellationRegistry\",\"type\":\"address\"}],\"name\":\"setRegistryContracts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_payoutAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_payoutPerMille\",\"type\":\"uint256\"}],\"name\":\"setRoyalty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ExchangeV5ABI is the input ABI used to generate the binding from.
// Deprecated: Use ExchangeV5MetaData.ABI instead.
var ExchangeV5ABI = ExchangeV5MetaData.ABI

// ExchangeV5 is an auto generated Go binding around an Ethereum contract.
type ExchangeV5 struct {
	ExchangeV5Caller     // Read-only binding to the contract
	ExchangeV5Transactor // Write-only binding to the contract
	ExchangeV5Filterer   // Log filterer for contract events
}

// ExchangeV5Caller is an auto generated read-only Go binding around an Ethereum contract.
type ExchangeV5Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeV5Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ExchangeV5Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeV5Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExchangeV5Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeV5Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExchangeV5Session struct {
	Contract     *ExchangeV5       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExchangeV5CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExchangeV5CallerSession struct {
	Contract *ExchangeV5Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ExchangeV5TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExchangeV5TransactorSession struct {
	Contract     *ExchangeV5Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ExchangeV5Raw is an auto generated low-level Go binding around an Ethereum contract.
type ExchangeV5Raw struct {
	Contract *ExchangeV5 // Generic contract binding to access the raw methods on
}

// ExchangeV5CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExchangeV5CallerRaw struct {
	Contract *ExchangeV5Caller // Generic read-only contract binding to access the raw methods on
}

// ExchangeV5TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExchangeV5TransactorRaw struct {
	Contract *ExchangeV5Transactor // Generic write-only contract binding to access the raw methods on
}

// NewExchangeV5 creates a new instance of ExchangeV5, bound to a specific deployed contract.
func NewExchangeV5(address common.Address, backend bind.ContractBackend) (*ExchangeV5, error) {
	contract, err := bindExchangeV5(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExchangeV5{ExchangeV5Caller: ExchangeV5Caller{contract: contract}, ExchangeV5Transactor: ExchangeV5Transactor{contract: contract}, ExchangeV5Filterer: ExchangeV5Filterer{contract: contract}}, nil
}

// NewExchangeV5Caller creates a new read-only instance of ExchangeV5, bound to a specific deployed contract.
func NewExchangeV5Caller(address common.Address, caller bind.ContractCaller) (*ExchangeV5Caller, error) {
	contract, err := bindExchangeV5(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExchangeV5Caller{contract: contract}, nil
}

// NewExchangeV5Transactor creates a new write-only instance of ExchangeV5, bound to a specific deployed contract.
func NewExchangeV5Transactor(address common.Address, transactor bind.ContractTransactor) (*ExchangeV5Transactor, error) {
	contract, err := bindExchangeV5(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExchangeV5Transactor{contract: contract}, nil
}

// NewExchangeV5Filterer creates a new log filterer instance of ExchangeV5, bound to a specific deployed contract.
func NewExchangeV5Filterer(address common.Address, filterer bind.ContractFilterer) (*ExchangeV5Filterer, error) {
	contract, err := bindExchangeV5(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExchangeV5Filterer{contract: contract}, nil
}

// bindExchangeV5 binds a generic wrapper to an already deployed contract.
func bindExchangeV5(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ExchangeV5ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExchangeV5 *ExchangeV5Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExchangeV5.Contract.ExchangeV5Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExchangeV5 *ExchangeV5Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExchangeV5.Contract.ExchangeV5Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExchangeV5 *ExchangeV5Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExchangeV5.Contract.ExchangeV5Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExchangeV5 *ExchangeV5CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExchangeV5.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExchangeV5 *ExchangeV5TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExchangeV5.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExchangeV5 *ExchangeV5TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExchangeV5.Contract.contract.Transact(opts, method, params...)
}

// CalculateCurrentPrice is a free data retrieval call binding the contract method 0xc3fbdb7c.
//
// Solidity: function calculateCurrentPrice(uint256 startTime, uint256 endTime, uint256 startPrice, uint256 endPrice) view returns(uint256)
func (_ExchangeV5 *ExchangeV5Caller) CalculateCurrentPrice(opts *bind.CallOpts, startTime *big.Int, endTime *big.Int, startPrice *big.Int, endPrice *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ExchangeV5.contract.Call(opts, &out, "calculateCurrentPrice", startTime, endTime, startPrice, endPrice)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateCurrentPrice is a free data retrieval call binding the contract method 0xc3fbdb7c.
//
// Solidity: function calculateCurrentPrice(uint256 startTime, uint256 endTime, uint256 startPrice, uint256 endPrice) view returns(uint256)
func (_ExchangeV5 *ExchangeV5Session) CalculateCurrentPrice(startTime *big.Int, endTime *big.Int, startPrice *big.Int, endPrice *big.Int) (*big.Int, error) {
	return _ExchangeV5.Contract.CalculateCurrentPrice(&_ExchangeV5.CallOpts, startTime, endTime, startPrice, endPrice)
}

// CalculateCurrentPrice is a free data retrieval call binding the contract method 0xc3fbdb7c.
//
// Solidity: function calculateCurrentPrice(uint256 startTime, uint256 endTime, uint256 startPrice, uint256 endPrice) view returns(uint256)
func (_ExchangeV5 *ExchangeV5CallerSession) CalculateCurrentPrice(startTime *big.Int, endTime *big.Int, startPrice *big.Int, endPrice *big.Int) (*big.Int, error) {
	return _ExchangeV5.Contract.CalculateCurrentPrice(&_ExchangeV5.CallOpts, startTime, endTime, startPrice, endPrice)
}

// GetRoyaltyPayoutAddress is a free data retrieval call binding the contract method 0x2df5cb23.
//
// Solidity: function getRoyaltyPayoutAddress(address contractAddress) view returns(address)
func (_ExchangeV5 *ExchangeV5Caller) GetRoyaltyPayoutAddress(opts *bind.CallOpts, contractAddress common.Address) (common.Address, error) {
	var out []interface{}
	err := _ExchangeV5.contract.Call(opts, &out, "getRoyaltyPayoutAddress", contractAddress)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoyaltyPayoutAddress is a free data retrieval call binding the contract method 0x2df5cb23.
//
// Solidity: function getRoyaltyPayoutAddress(address contractAddress) view returns(address)
func (_ExchangeV5 *ExchangeV5Session) GetRoyaltyPayoutAddress(contractAddress common.Address) (common.Address, error) {
	return _ExchangeV5.Contract.GetRoyaltyPayoutAddress(&_ExchangeV5.CallOpts, contractAddress)
}

// GetRoyaltyPayoutAddress is a free data retrieval call binding the contract method 0x2df5cb23.
//
// Solidity: function getRoyaltyPayoutAddress(address contractAddress) view returns(address)
func (_ExchangeV5 *ExchangeV5CallerSession) GetRoyaltyPayoutAddress(contractAddress common.Address) (common.Address, error) {
	return _ExchangeV5.Contract.GetRoyaltyPayoutAddress(&_ExchangeV5.CallOpts, contractAddress)
}

// GetRoyaltyPayoutRate is a free data retrieval call binding the contract method 0xec4210ff.
//
// Solidity: function getRoyaltyPayoutRate(address contractAddress) view returns(uint256)
func (_ExchangeV5 *ExchangeV5Caller) GetRoyaltyPayoutRate(opts *bind.CallOpts, contractAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ExchangeV5.contract.Call(opts, &out, "getRoyaltyPayoutRate", contractAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoyaltyPayoutRate is a free data retrieval call binding the contract method 0xec4210ff.
//
// Solidity: function getRoyaltyPayoutRate(address contractAddress) view returns(uint256)
func (_ExchangeV5 *ExchangeV5Session) GetRoyaltyPayoutRate(contractAddress common.Address) (*big.Int, error) {
	return _ExchangeV5.Contract.GetRoyaltyPayoutRate(&_ExchangeV5.CallOpts, contractAddress)
}

// GetRoyaltyPayoutRate is a free data retrieval call binding the contract method 0xec4210ff.
//
// Solidity: function getRoyaltyPayoutRate(address contractAddress) view returns(uint256)
func (_ExchangeV5 *ExchangeV5CallerSession) GetRoyaltyPayoutRate(contractAddress common.Address) (*big.Int, error) {
	return _ExchangeV5.Contract.GetRoyaltyPayoutRate(&_ExchangeV5.CallOpts, contractAddress)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ExchangeV5 *ExchangeV5Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExchangeV5.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ExchangeV5 *ExchangeV5Session) Owner() (common.Address, error) {
	return _ExchangeV5.Contract.Owner(&_ExchangeV5.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ExchangeV5 *ExchangeV5CallerSession) Owner() (common.Address, error) {
	return _ExchangeV5.Contract.Owner(&_ExchangeV5.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ExchangeV5 *ExchangeV5Caller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ExchangeV5.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ExchangeV5 *ExchangeV5Session) Paused() (bool, error) {
	return _ExchangeV5.Contract.Paused(&_ExchangeV5.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ExchangeV5 *ExchangeV5CallerSession) Paused() (bool, error) {
	return _ExchangeV5.Contract.Paused(&_ExchangeV5.CallOpts)
}

// CancelPreviousSellOrders is a paid mutator transaction binding the contract method 0x083d089d.
//
// Solidity: function cancelPreviousSellOrders(address addr, address tokenAddr, uint256 tokenId) returns()
func (_ExchangeV5 *ExchangeV5Transactor) CancelPreviousSellOrders(opts *bind.TransactOpts, addr common.Address, tokenAddr common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ExchangeV5.contract.Transact(opts, "cancelPreviousSellOrders", addr, tokenAddr, tokenId)
}

// CancelPreviousSellOrders is a paid mutator transaction binding the contract method 0x083d089d.
//
// Solidity: function cancelPreviousSellOrders(address addr, address tokenAddr, uint256 tokenId) returns()
func (_ExchangeV5 *ExchangeV5Session) CancelPreviousSellOrders(addr common.Address, tokenAddr common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ExchangeV5.Contract.CancelPreviousSellOrders(&_ExchangeV5.TransactOpts, addr, tokenAddr, tokenId)
}

// CancelPreviousSellOrders is a paid mutator transaction binding the contract method 0x083d089d.
//
// Solidity: function cancelPreviousSellOrders(address addr, address tokenAddr, uint256 tokenId) returns()
func (_ExchangeV5 *ExchangeV5TransactorSession) CancelPreviousSellOrders(addr common.Address, tokenAddr common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ExchangeV5.Contract.CancelPreviousSellOrders(&_ExchangeV5.TransactOpts, addr, tokenAddr, tokenId)
}

// FillSellOrder is a paid mutator transaction binding the contract method 0x912d97fc.
//
// Solidity: function fillSellOrder(address seller, address contractAddress, uint256 tokenId, uint256 startTime, uint256 expiration, uint256 price, uint256 quantity, uint256 createdAtBlockNumber, address paymentERC20, bytes signature, address buyer) payable returns()
func (_ExchangeV5 *ExchangeV5Transactor) FillSellOrder(opts *bind.TransactOpts, seller common.Address, contractAddress common.Address, tokenId *big.Int, startTime *big.Int, expiration *big.Int, price *big.Int, quantity *big.Int, createdAtBlockNumber *big.Int, paymentERC20 common.Address, signature []byte, buyer common.Address) (*types.Transaction, error) {
	return _ExchangeV5.contract.Transact(opts, "fillSellOrder", seller, contractAddress, tokenId, startTime, expiration, price, quantity, createdAtBlockNumber, paymentERC20, signature, buyer)
}

// FillSellOrder is a paid mutator transaction binding the contract method 0x912d97fc.
//
// Solidity: function fillSellOrder(address seller, address contractAddress, uint256 tokenId, uint256 startTime, uint256 expiration, uint256 price, uint256 quantity, uint256 createdAtBlockNumber, address paymentERC20, bytes signature, address buyer) payable returns()
func (_ExchangeV5 *ExchangeV5Session) FillSellOrder(seller common.Address, contractAddress common.Address, tokenId *big.Int, startTime *big.Int, expiration *big.Int, price *big.Int, quantity *big.Int, createdAtBlockNumber *big.Int, paymentERC20 common.Address, signature []byte, buyer common.Address) (*types.Transaction, error) {
	return _ExchangeV5.Contract.FillSellOrder(&_ExchangeV5.TransactOpts, seller, contractAddress, tokenId, startTime, expiration, price, quantity, createdAtBlockNumber, paymentERC20, signature, buyer)
}

// FillSellOrder is a paid mutator transaction binding the contract method 0x912d97fc.
//
// Solidity: function fillSellOrder(address seller, address contractAddress, uint256 tokenId, uint256 startTime, uint256 expiration, uint256 price, uint256 quantity, uint256 createdAtBlockNumber, address paymentERC20, bytes signature, address buyer) payable returns()
func (_ExchangeV5 *ExchangeV5TransactorSession) FillSellOrder(seller common.Address, contractAddress common.Address, tokenId *big.Int, startTime *big.Int, expiration *big.Int, price *big.Int, quantity *big.Int, createdAtBlockNumber *big.Int, paymentERC20 common.Address, signature []byte, buyer common.Address) (*types.Transaction, error) {
	return _ExchangeV5.Contract.FillSellOrder(&_ExchangeV5.TransactOpts, seller, contractAddress, tokenId, startTime, expiration, price, quantity, createdAtBlockNumber, paymentERC20, signature, buyer)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ExchangeV5 *ExchangeV5Transactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExchangeV5.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ExchangeV5 *ExchangeV5Session) Pause() (*types.Transaction, error) {
	return _ExchangeV5.Contract.Pause(&_ExchangeV5.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ExchangeV5 *ExchangeV5TransactorSession) Pause() (*types.Transaction, error) {
	return _ExchangeV5.Contract.Pause(&_ExchangeV5.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ExchangeV5 *ExchangeV5Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExchangeV5.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ExchangeV5 *ExchangeV5Session) RenounceOwnership() (*types.Transaction, error) {
	return _ExchangeV5.Contract.RenounceOwnership(&_ExchangeV5.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ExchangeV5 *ExchangeV5TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ExchangeV5.Contract.RenounceOwnership(&_ExchangeV5.TransactOpts)
}

// SetMakerWallet is a paid mutator transaction binding the contract method 0x2b812a34.
//
// Solidity: function setMakerWallet(address _newMakerWallet) returns()
func (_ExchangeV5 *ExchangeV5Transactor) SetMakerWallet(opts *bind.TransactOpts, _newMakerWallet common.Address) (*types.Transaction, error) {
	return _ExchangeV5.contract.Transact(opts, "setMakerWallet", _newMakerWallet)
}

// SetMakerWallet is a paid mutator transaction binding the contract method 0x2b812a34.
//
// Solidity: function setMakerWallet(address _newMakerWallet) returns()
func (_ExchangeV5 *ExchangeV5Session) SetMakerWallet(_newMakerWallet common.Address) (*types.Transaction, error) {
	return _ExchangeV5.Contract.SetMakerWallet(&_ExchangeV5.TransactOpts, _newMakerWallet)
}

// SetMakerWallet is a paid mutator transaction binding the contract method 0x2b812a34.
//
// Solidity: function setMakerWallet(address _newMakerWallet) returns()
func (_ExchangeV5 *ExchangeV5TransactorSession) SetMakerWallet(_newMakerWallet common.Address) (*types.Transaction, error) {
	return _ExchangeV5.Contract.SetMakerWallet(&_ExchangeV5.TransactOpts, _newMakerWallet)
}

// SetRegistryContracts is a paid mutator transaction binding the contract method 0xb1216b79.
//
// Solidity: function setRegistryContracts(address _royaltyRegistry, address _cancellationRegistry) returns()
func (_ExchangeV5 *ExchangeV5Transactor) SetRegistryContracts(opts *bind.TransactOpts, _royaltyRegistry common.Address, _cancellationRegistry common.Address) (*types.Transaction, error) {
	return _ExchangeV5.contract.Transact(opts, "setRegistryContracts", _royaltyRegistry, _cancellationRegistry)
}

// SetRegistryContracts is a paid mutator transaction binding the contract method 0xb1216b79.
//
// Solidity: function setRegistryContracts(address _royaltyRegistry, address _cancellationRegistry) returns()
func (_ExchangeV5 *ExchangeV5Session) SetRegistryContracts(_royaltyRegistry common.Address, _cancellationRegistry common.Address) (*types.Transaction, error) {
	return _ExchangeV5.Contract.SetRegistryContracts(&_ExchangeV5.TransactOpts, _royaltyRegistry, _cancellationRegistry)
}

// SetRegistryContracts is a paid mutator transaction binding the contract method 0xb1216b79.
//
// Solidity: function setRegistryContracts(address _royaltyRegistry, address _cancellationRegistry) returns()
func (_ExchangeV5 *ExchangeV5TransactorSession) SetRegistryContracts(_royaltyRegistry common.Address, _cancellationRegistry common.Address) (*types.Transaction, error) {
	return _ExchangeV5.Contract.SetRegistryContracts(&_ExchangeV5.TransactOpts, _royaltyRegistry, _cancellationRegistry)
}

// SetRoyalty is a paid mutator transaction binding the contract method 0x55c338aa.
//
// Solidity: function setRoyalty(address contractAddress, address _payoutAddress, uint256 _payoutPerMille) returns()
func (_ExchangeV5 *ExchangeV5Transactor) SetRoyalty(opts *bind.TransactOpts, contractAddress common.Address, _payoutAddress common.Address, _payoutPerMille *big.Int) (*types.Transaction, error) {
	return _ExchangeV5.contract.Transact(opts, "setRoyalty", contractAddress, _payoutAddress, _payoutPerMille)
}

// SetRoyalty is a paid mutator transaction binding the contract method 0x55c338aa.
//
// Solidity: function setRoyalty(address contractAddress, address _payoutAddress, uint256 _payoutPerMille) returns()
func (_ExchangeV5 *ExchangeV5Session) SetRoyalty(contractAddress common.Address, _payoutAddress common.Address, _payoutPerMille *big.Int) (*types.Transaction, error) {
	return _ExchangeV5.Contract.SetRoyalty(&_ExchangeV5.TransactOpts, contractAddress, _payoutAddress, _payoutPerMille)
}

// SetRoyalty is a paid mutator transaction binding the contract method 0x55c338aa.
//
// Solidity: function setRoyalty(address contractAddress, address _payoutAddress, uint256 _payoutPerMille) returns()
func (_ExchangeV5 *ExchangeV5TransactorSession) SetRoyalty(contractAddress common.Address, _payoutAddress common.Address, _payoutPerMille *big.Int) (*types.Transaction, error) {
	return _ExchangeV5.Contract.SetRoyalty(&_ExchangeV5.TransactOpts, contractAddress, _payoutAddress, _payoutPerMille)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ExchangeV5 *ExchangeV5Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ExchangeV5.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ExchangeV5 *ExchangeV5Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ExchangeV5.Contract.TransferOwnership(&_ExchangeV5.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ExchangeV5 *ExchangeV5TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ExchangeV5.Contract.TransferOwnership(&_ExchangeV5.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ExchangeV5 *ExchangeV5Transactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExchangeV5.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ExchangeV5 *ExchangeV5Session) Unpause() (*types.Transaction, error) {
	return _ExchangeV5.Contract.Unpause(&_ExchangeV5.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ExchangeV5 *ExchangeV5TransactorSession) Unpause() (*types.Transaction, error) {
	return _ExchangeV5.Contract.Unpause(&_ExchangeV5.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_ExchangeV5 *ExchangeV5Transactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExchangeV5.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_ExchangeV5 *ExchangeV5Session) Withdraw() (*types.Transaction, error) {
	return _ExchangeV5.Contract.Withdraw(&_ExchangeV5.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_ExchangeV5 *ExchangeV5TransactorSession) Withdraw() (*types.Transaction, error) {
	return _ExchangeV5.Contract.Withdraw(&_ExchangeV5.TransactOpts)
}

// ExchangeV5OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ExchangeV5 contract.
type ExchangeV5OwnershipTransferredIterator struct {
	Event *ExchangeV5OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ExchangeV5OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeV5OwnershipTransferred)
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
		it.Event = new(ExchangeV5OwnershipTransferred)
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
func (it *ExchangeV5OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeV5OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeV5OwnershipTransferred represents a OwnershipTransferred event raised by the ExchangeV5 contract.
type ExchangeV5OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ExchangeV5 *ExchangeV5Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ExchangeV5OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ExchangeV5.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeV5OwnershipTransferredIterator{contract: _ExchangeV5.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ExchangeV5 *ExchangeV5Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ExchangeV5OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ExchangeV5.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeV5OwnershipTransferred)
				if err := _ExchangeV5.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ExchangeV5 *ExchangeV5Filterer) ParseOwnershipTransferred(log types.Log) (*ExchangeV5OwnershipTransferred, error) {
	event := new(ExchangeV5OwnershipTransferred)
	if err := _ExchangeV5.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeV5PausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ExchangeV5 contract.
type ExchangeV5PausedIterator struct {
	Event *ExchangeV5Paused // Event containing the contract specifics and raw log

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
func (it *ExchangeV5PausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeV5Paused)
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
		it.Event = new(ExchangeV5Paused)
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
func (it *ExchangeV5PausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeV5PausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeV5Paused represents a Paused event raised by the ExchangeV5 contract.
type ExchangeV5Paused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ExchangeV5 *ExchangeV5Filterer) FilterPaused(opts *bind.FilterOpts) (*ExchangeV5PausedIterator, error) {

	logs, sub, err := _ExchangeV5.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &ExchangeV5PausedIterator{contract: _ExchangeV5.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ExchangeV5 *ExchangeV5Filterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ExchangeV5Paused) (event.Subscription, error) {

	logs, sub, err := _ExchangeV5.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeV5Paused)
				if err := _ExchangeV5.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_ExchangeV5 *ExchangeV5Filterer) ParsePaused(log types.Log) (*ExchangeV5Paused, error) {
	event := new(ExchangeV5Paused)
	if err := _ExchangeV5.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeV5SellOrderFilledIterator is returned from FilterSellOrderFilled and is used to iterate over the raw logs and unpacked data for SellOrderFilled events raised by the ExchangeV5 contract.
type ExchangeV5SellOrderFilledIterator struct {
	Event *ExchangeV5SellOrderFilled // Event containing the contract specifics and raw log

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
func (it *ExchangeV5SellOrderFilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeV5SellOrderFilled)
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
		it.Event = new(ExchangeV5SellOrderFilled)
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
func (it *ExchangeV5SellOrderFilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeV5SellOrderFilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeV5SellOrderFilled represents a SellOrderFilled event raised by the ExchangeV5 contract.
type ExchangeV5SellOrderFilled struct {
	Seller          common.Address
	Buyer           common.Address
	ContractAddress common.Address
	TokenId         *big.Int
	Price           *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSellOrderFilled is a free log retrieval operation binding the contract event 0x70ba0d31158674eea8365d0f7b9ac70e552cc28b8bb848664e4feb939c6578f8.
//
// Solidity: event SellOrderFilled(address indexed seller, address buyer, address indexed contractAddress, uint256 indexed tokenId, uint256 price)
func (_ExchangeV5 *ExchangeV5Filterer) FilterSellOrderFilled(opts *bind.FilterOpts, seller []common.Address, contractAddress []common.Address, tokenId []*big.Int) (*ExchangeV5SellOrderFilledIterator, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ExchangeV5.contract.FilterLogs(opts, "SellOrderFilled", sellerRule, contractAddressRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeV5SellOrderFilledIterator{contract: _ExchangeV5.contract, event: "SellOrderFilled", logs: logs, sub: sub}, nil
}

// WatchSellOrderFilled is a free log subscription operation binding the contract event 0x70ba0d31158674eea8365d0f7b9ac70e552cc28b8bb848664e4feb939c6578f8.
//
// Solidity: event SellOrderFilled(address indexed seller, address buyer, address indexed contractAddress, uint256 indexed tokenId, uint256 price)
func (_ExchangeV5 *ExchangeV5Filterer) WatchSellOrderFilled(opts *bind.WatchOpts, sink chan<- *ExchangeV5SellOrderFilled, seller []common.Address, contractAddress []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ExchangeV5.contract.WatchLogs(opts, "SellOrderFilled", sellerRule, contractAddressRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeV5SellOrderFilled)
				if err := _ExchangeV5.contract.UnpackLog(event, "SellOrderFilled", log); err != nil {
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

// ParseSellOrderFilled is a log parse operation binding the contract event 0x70ba0d31158674eea8365d0f7b9ac70e552cc28b8bb848664e4feb939c6578f8.
//
// Solidity: event SellOrderFilled(address indexed seller, address buyer, address indexed contractAddress, uint256 indexed tokenId, uint256 price)
func (_ExchangeV5 *ExchangeV5Filterer) ParseSellOrderFilled(log types.Log) (*ExchangeV5SellOrderFilled, error) {
	event := new(ExchangeV5SellOrderFilled)
	if err := _ExchangeV5.contract.UnpackLog(event, "SellOrderFilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeV5UnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ExchangeV5 contract.
type ExchangeV5UnpausedIterator struct {
	Event *ExchangeV5Unpaused // Event containing the contract specifics and raw log

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
func (it *ExchangeV5UnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeV5Unpaused)
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
		it.Event = new(ExchangeV5Unpaused)
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
func (it *ExchangeV5UnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeV5UnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeV5Unpaused represents a Unpaused event raised by the ExchangeV5 contract.
type ExchangeV5Unpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ExchangeV5 *ExchangeV5Filterer) FilterUnpaused(opts *bind.FilterOpts) (*ExchangeV5UnpausedIterator, error) {

	logs, sub, err := _ExchangeV5.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &ExchangeV5UnpausedIterator{contract: _ExchangeV5.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ExchangeV5 *ExchangeV5Filterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ExchangeV5Unpaused) (event.Subscription, error) {

	logs, sub, err := _ExchangeV5.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeV5Unpaused)
				if err := _ExchangeV5.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_ExchangeV5 *ExchangeV5Filterer) ParseUnpaused(log types.Log) (*ExchangeV5Unpaused, error) {
	event := new(ExchangeV5Unpaused)
	if err := _ExchangeV5.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
