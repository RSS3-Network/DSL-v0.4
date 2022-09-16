// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package party

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

// PartyProxyMetaData contains all meta data concerning the PartyProxy contract.
var PartyProxyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"logic\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// PartyProxyABI is the input ABI used to generate the binding from.
// Deprecated: Use PartyProxyMetaData.ABI instead.
var PartyProxyABI = PartyProxyMetaData.ABI

// PartyProxy is an auto generated Go binding around an Ethereum contract.
type PartyProxy struct {
	PartyProxyCaller     // Read-only binding to the contract
	PartyProxyTransactor // Write-only binding to the contract
	PartyProxyFilterer   // Log filterer for contract events
}

// PartyProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type PartyProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PartyProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PartyProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PartyProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PartyProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PartyProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PartyProxySession struct {
	Contract     *PartyProxy       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PartyProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PartyProxyCallerSession struct {
	Contract *PartyProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// PartyProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PartyProxyTransactorSession struct {
	Contract     *PartyProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// PartyProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type PartyProxyRaw struct {
	Contract *PartyProxy // Generic contract binding to access the raw methods on
}

// PartyProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PartyProxyCallerRaw struct {
	Contract *PartyProxyCaller // Generic read-only contract binding to access the raw methods on
}

// PartyProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PartyProxyTransactorRaw struct {
	Contract *PartyProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPartyProxy creates a new instance of PartyProxy, bound to a specific deployed contract.
func NewPartyProxy(address common.Address, backend bind.ContractBackend) (*PartyProxy, error) {
	contract, err := bindPartyProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PartyProxy{PartyProxyCaller: PartyProxyCaller{contract: contract}, PartyProxyTransactor: PartyProxyTransactor{contract: contract}, PartyProxyFilterer: PartyProxyFilterer{contract: contract}}, nil
}

// NewPartyProxyCaller creates a new read-only instance of PartyProxy, bound to a specific deployed contract.
func NewPartyProxyCaller(address common.Address, caller bind.ContractCaller) (*PartyProxyCaller, error) {
	contract, err := bindPartyProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PartyProxyCaller{contract: contract}, nil
}

// NewPartyProxyTransactor creates a new write-only instance of PartyProxy, bound to a specific deployed contract.
func NewPartyProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*PartyProxyTransactor, error) {
	contract, err := bindPartyProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PartyProxyTransactor{contract: contract}, nil
}

// NewPartyProxyFilterer creates a new log filterer instance of PartyProxy, bound to a specific deployed contract.
func NewPartyProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*PartyProxyFilterer, error) {
	contract, err := bindPartyProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PartyProxyFilterer{contract: contract}, nil
}

// bindPartyProxy binds a generic wrapper to an already deployed contract.
func bindPartyProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PartyProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PartyProxy *PartyProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PartyProxy.Contract.PartyProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PartyProxy *PartyProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PartyProxy.Contract.PartyProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PartyProxy *PartyProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PartyProxy.Contract.PartyProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PartyProxy *PartyProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PartyProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PartyProxy *PartyProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PartyProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PartyProxy *PartyProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PartyProxy.Contract.contract.Transact(opts, method, params...)
}

// Logic is a free data retrieval call binding the contract method 0xd7dfa0dd.
//
// Solidity: function logic() view returns(address)
func (_PartyProxy *PartyProxyCaller) Logic(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PartyProxy.contract.Call(opts, &out, "logic")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Logic is a free data retrieval call binding the contract method 0xd7dfa0dd.
//
// Solidity: function logic() view returns(address)
func (_PartyProxy *PartyProxySession) Logic() (common.Address, error) {
	return _PartyProxy.Contract.Logic(&_PartyProxy.CallOpts)
}

// Logic is a free data retrieval call binding the contract method 0xd7dfa0dd.
//
// Solidity: function logic() view returns(address)
func (_PartyProxy *PartyProxyCallerSession) Logic() (common.Address, error) {
	return _PartyProxy.Contract.Logic(&_PartyProxy.CallOpts)
}
