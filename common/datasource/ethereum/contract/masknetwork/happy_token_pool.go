// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package masknetwork

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

// HappyTokenPoolMetaData contains all meta data concerning the HappyTokenPool contract.
var HappyTokenPoolMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"to_value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token_address\",\"type\":\"address\"}],\"name\":\"ClaimSuccess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token_address\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"remaining_balance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint128[]\",\"name\":\"exchanged_values\",\"type\":\"uint128[]\"}],\"name\":\"DestructSuccess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"creation_time\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token_address\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"exchange_addrs\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint128[]\",\"name\":\"ratios\",\"type\":\"uint128[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"qualification\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"FillSuccess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"swapper\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from_address\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to_address\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"from_value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"to_value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"input_total\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"claimed\",\"type\":\"bool\"}],\"name\":\"SwapSuccess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token_address\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdraw_balance\",\"type\":\"uint256\"}],\"name\":\"WithdrawSuccess\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"base_time\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"check_availability\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"exchange_addrs\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"remaining\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"started\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"expired\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"destructed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"unlock_time\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapped\",\"type\":\"uint256\"},{\"internalType\":\"uint128[]\",\"name\":\"exchanged_tokens\",\"type\":\"uint128[]\"},{\"internalType\":\"bool\",\"name\":\"claimed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"start_time\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end_time\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"qualification_addr\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"ito_ids\",\"type\":\"bytes32[]\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"destruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_end\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_message\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"_exchange_addrs\",\"type\":\"address[]\"},{\"internalType\":\"uint128[]\",\"name\":\"_ratios\",\"type\":\"uint128[]\"},{\"internalType\":\"uint256\",\"name\":\"_unlock_time\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_token_addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_total_tokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_limit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_qualification\",\"type\":\"address\"}],\"name\":\"fill_pool\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_base_time\",\"type\":\"uint64\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_unlock_time\",\"type\":\"uint32\"}],\"name\":\"setUnlockTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"verification\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"exchange_addr_i\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"input_total\",\"type\":\"uint128\"},{\"internalType\":\"bytes32[]\",\"name\":\"data\",\"type\":\"bytes32[]\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"swapped\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// HappyTokenPoolABI is the input ABI used to generate the binding from.
// Deprecated: Use HappyTokenPoolMetaData.ABI instead.
var HappyTokenPoolABI = HappyTokenPoolMetaData.ABI

// HappyTokenPool is an auto generated Go binding around an Ethereum contract.
type HappyTokenPool struct {
	HappyTokenPoolCaller     // Read-only binding to the contract
	HappyTokenPoolTransactor // Write-only binding to the contract
	HappyTokenPoolFilterer   // Log filterer for contract events
}

// HappyTokenPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type HappyTokenPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HappyTokenPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HappyTokenPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HappyTokenPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HappyTokenPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HappyTokenPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HappyTokenPoolSession struct {
	Contract     *HappyTokenPool   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HappyTokenPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HappyTokenPoolCallerSession struct {
	Contract *HappyTokenPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// HappyTokenPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HappyTokenPoolTransactorSession struct {
	Contract     *HappyTokenPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// HappyTokenPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type HappyTokenPoolRaw struct {
	Contract *HappyTokenPool // Generic contract binding to access the raw methods on
}

// HappyTokenPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HappyTokenPoolCallerRaw struct {
	Contract *HappyTokenPoolCaller // Generic read-only contract binding to access the raw methods on
}

// HappyTokenPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HappyTokenPoolTransactorRaw struct {
	Contract *HappyTokenPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHappyTokenPool creates a new instance of HappyTokenPool, bound to a specific deployed contract.
func NewHappyTokenPool(address common.Address, backend bind.ContractBackend) (*HappyTokenPool, error) {
	contract, err := bindHappyTokenPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HappyTokenPool{HappyTokenPoolCaller: HappyTokenPoolCaller{contract: contract}, HappyTokenPoolTransactor: HappyTokenPoolTransactor{contract: contract}, HappyTokenPoolFilterer: HappyTokenPoolFilterer{contract: contract}}, nil
}

// NewHappyTokenPoolCaller creates a new read-only instance of HappyTokenPool, bound to a specific deployed contract.
func NewHappyTokenPoolCaller(address common.Address, caller bind.ContractCaller) (*HappyTokenPoolCaller, error) {
	contract, err := bindHappyTokenPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HappyTokenPoolCaller{contract: contract}, nil
}

// NewHappyTokenPoolTransactor creates a new write-only instance of HappyTokenPool, bound to a specific deployed contract.
func NewHappyTokenPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*HappyTokenPoolTransactor, error) {
	contract, err := bindHappyTokenPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HappyTokenPoolTransactor{contract: contract}, nil
}

// NewHappyTokenPoolFilterer creates a new log filterer instance of HappyTokenPool, bound to a specific deployed contract.
func NewHappyTokenPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*HappyTokenPoolFilterer, error) {
	contract, err := bindHappyTokenPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HappyTokenPoolFilterer{contract: contract}, nil
}

// bindHappyTokenPool binds a generic wrapper to an already deployed contract.
func bindHappyTokenPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HappyTokenPoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HappyTokenPool *HappyTokenPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HappyTokenPool.Contract.HappyTokenPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HappyTokenPool *HappyTokenPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HappyTokenPool.Contract.HappyTokenPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HappyTokenPool *HappyTokenPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HappyTokenPool.Contract.HappyTokenPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HappyTokenPool *HappyTokenPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HappyTokenPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HappyTokenPool *HappyTokenPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HappyTokenPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HappyTokenPool *HappyTokenPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HappyTokenPool.Contract.contract.Transact(opts, method, params...)
}

// BaseTime is a free data retrieval call binding the contract method 0x1686a9d1.
//
// Solidity: function base_time() view returns(uint64)
func (_HappyTokenPool *HappyTokenPoolCaller) BaseTime(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _HappyTokenPool.contract.Call(opts, &out, "base_time")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// BaseTime is a free data retrieval call binding the contract method 0x1686a9d1.
//
// Solidity: function base_time() view returns(uint64)
func (_HappyTokenPool *HappyTokenPoolSession) BaseTime() (uint64, error) {
	return _HappyTokenPool.Contract.BaseTime(&_HappyTokenPool.CallOpts)
}

// BaseTime is a free data retrieval call binding the contract method 0x1686a9d1.
//
// Solidity: function base_time() view returns(uint64)
func (_HappyTokenPool *HappyTokenPoolCallerSession) BaseTime() (uint64, error) {
	return _HappyTokenPool.Contract.BaseTime(&_HappyTokenPool.CallOpts)
}

// CheckAvailability is a free data retrieval call binding the contract method 0x6bfdaece.
//
// Solidity: function check_availability(bytes32 id) view returns(address[] exchange_addrs, uint256 remaining, bool started, bool expired, bool destructed, uint256 unlock_time, uint256 swapped, uint128[] exchanged_tokens, bool claimed, uint256 start_time, uint256 end_time, address qualification_addr)
func (_HappyTokenPool *HappyTokenPoolCaller) CheckAvailability(opts *bind.CallOpts, id [32]byte) (struct {
	ExchangeAddrs     []common.Address
	Remaining         *big.Int
	Started           bool
	Expired           bool
	Destructed        bool
	UnlockTime        *big.Int
	Swapped           *big.Int
	ExchangedTokens   []*big.Int
	Claimed           bool
	StartTime         *big.Int
	EndTime           *big.Int
	QualificationAddr common.Address
}, error) {
	var out []interface{}
	err := _HappyTokenPool.contract.Call(opts, &out, "check_availability", id)

	outstruct := new(struct {
		ExchangeAddrs     []common.Address
		Remaining         *big.Int
		Started           bool
		Expired           bool
		Destructed        bool
		UnlockTime        *big.Int
		Swapped           *big.Int
		ExchangedTokens   []*big.Int
		Claimed           bool
		StartTime         *big.Int
		EndTime           *big.Int
		QualificationAddr common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ExchangeAddrs = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Remaining = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Started = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.Expired = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.Destructed = *abi.ConvertType(out[4], new(bool)).(*bool)
	outstruct.UnlockTime = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Swapped = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.ExchangedTokens = *abi.ConvertType(out[7], new([]*big.Int)).(*[]*big.Int)
	outstruct.Claimed = *abi.ConvertType(out[8], new(bool)).(*bool)
	outstruct.StartTime = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	outstruct.EndTime = *abi.ConvertType(out[10], new(*big.Int)).(**big.Int)
	outstruct.QualificationAddr = *abi.ConvertType(out[11], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// CheckAvailability is a free data retrieval call binding the contract method 0x6bfdaece.
//
// Solidity: function check_availability(bytes32 id) view returns(address[] exchange_addrs, uint256 remaining, bool started, bool expired, bool destructed, uint256 unlock_time, uint256 swapped, uint128[] exchanged_tokens, bool claimed, uint256 start_time, uint256 end_time, address qualification_addr)
func (_HappyTokenPool *HappyTokenPoolSession) CheckAvailability(id [32]byte) (struct {
	ExchangeAddrs     []common.Address
	Remaining         *big.Int
	Started           bool
	Expired           bool
	Destructed        bool
	UnlockTime        *big.Int
	Swapped           *big.Int
	ExchangedTokens   []*big.Int
	Claimed           bool
	StartTime         *big.Int
	EndTime           *big.Int
	QualificationAddr common.Address
}, error) {
	return _HappyTokenPool.Contract.CheckAvailability(&_HappyTokenPool.CallOpts, id)
}

// CheckAvailability is a free data retrieval call binding the contract method 0x6bfdaece.
//
// Solidity: function check_availability(bytes32 id) view returns(address[] exchange_addrs, uint256 remaining, bool started, bool expired, bool destructed, uint256 unlock_time, uint256 swapped, uint128[] exchanged_tokens, bool claimed, uint256 start_time, uint256 end_time, address qualification_addr)
func (_HappyTokenPool *HappyTokenPoolCallerSession) CheckAvailability(id [32]byte) (struct {
	ExchangeAddrs     []common.Address
	Remaining         *big.Int
	Started           bool
	Expired           bool
	Destructed        bool
	UnlockTime        *big.Int
	Swapped           *big.Int
	ExchangedTokens   []*big.Int
	Claimed           bool
	StartTime         *big.Int
	EndTime           *big.Int
	QualificationAddr common.Address
}, error) {
	return _HappyTokenPool.Contract.CheckAvailability(&_HappyTokenPool.CallOpts, id)
}

// Claim is a paid mutator transaction binding the contract method 0xb391c508.
//
// Solidity: function claim(bytes32[] ito_ids) returns()
func (_HappyTokenPool *HappyTokenPoolTransactor) Claim(opts *bind.TransactOpts, ito_ids [][32]byte) (*types.Transaction, error) {
	return _HappyTokenPool.contract.Transact(opts, "claim", ito_ids)
}

// Claim is a paid mutator transaction binding the contract method 0xb391c508.
//
// Solidity: function claim(bytes32[] ito_ids) returns()
func (_HappyTokenPool *HappyTokenPoolSession) Claim(ito_ids [][32]byte) (*types.Transaction, error) {
	return _HappyTokenPool.Contract.Claim(&_HappyTokenPool.TransactOpts, ito_ids)
}

// Claim is a paid mutator transaction binding the contract method 0xb391c508.
//
// Solidity: function claim(bytes32[] ito_ids) returns()
func (_HappyTokenPool *HappyTokenPoolTransactorSession) Claim(ito_ids [][32]byte) (*types.Transaction, error) {
	return _HappyTokenPool.Contract.Claim(&_HappyTokenPool.TransactOpts, ito_ids)
}

// Destruct is a paid mutator transaction binding the contract method 0xcc0cab4c.
//
// Solidity: function destruct(bytes32 id) returns()
func (_HappyTokenPool *HappyTokenPoolTransactor) Destruct(opts *bind.TransactOpts, id [32]byte) (*types.Transaction, error) {
	return _HappyTokenPool.contract.Transact(opts, "destruct", id)
}

// Destruct is a paid mutator transaction binding the contract method 0xcc0cab4c.
//
// Solidity: function destruct(bytes32 id) returns()
func (_HappyTokenPool *HappyTokenPoolSession) Destruct(id [32]byte) (*types.Transaction, error) {
	return _HappyTokenPool.Contract.Destruct(&_HappyTokenPool.TransactOpts, id)
}

// Destruct is a paid mutator transaction binding the contract method 0xcc0cab4c.
//
// Solidity: function destruct(bytes32 id) returns()
func (_HappyTokenPool *HappyTokenPoolTransactorSession) Destruct(id [32]byte) (*types.Transaction, error) {
	return _HappyTokenPool.Contract.Destruct(&_HappyTokenPool.TransactOpts, id)
}

// FillPool is a paid mutator transaction binding the contract method 0xef65dc7e.
//
// Solidity: function fill_pool(bytes32 _hash, uint256 _start, uint256 _end, string _message, address[] _exchange_addrs, uint128[] _ratios, uint256 _unlock_time, address _token_addr, uint256 _total_tokens, uint256 _limit, address _qualification) payable returns()
func (_HappyTokenPool *HappyTokenPoolTransactor) FillPool(opts *bind.TransactOpts, _hash [32]byte, _start *big.Int, _end *big.Int, _message string, _exchange_addrs []common.Address, _ratios []*big.Int, _unlock_time *big.Int, _token_addr common.Address, _total_tokens *big.Int, _limit *big.Int, _qualification common.Address) (*types.Transaction, error) {
	return _HappyTokenPool.contract.Transact(opts, "fill_pool", _hash, _start, _end, _message, _exchange_addrs, _ratios, _unlock_time, _token_addr, _total_tokens, _limit, _qualification)
}

// FillPool is a paid mutator transaction binding the contract method 0xef65dc7e.
//
// Solidity: function fill_pool(bytes32 _hash, uint256 _start, uint256 _end, string _message, address[] _exchange_addrs, uint128[] _ratios, uint256 _unlock_time, address _token_addr, uint256 _total_tokens, uint256 _limit, address _qualification) payable returns()
func (_HappyTokenPool *HappyTokenPoolSession) FillPool(_hash [32]byte, _start *big.Int, _end *big.Int, _message string, _exchange_addrs []common.Address, _ratios []*big.Int, _unlock_time *big.Int, _token_addr common.Address, _total_tokens *big.Int, _limit *big.Int, _qualification common.Address) (*types.Transaction, error) {
	return _HappyTokenPool.Contract.FillPool(&_HappyTokenPool.TransactOpts, _hash, _start, _end, _message, _exchange_addrs, _ratios, _unlock_time, _token_addr, _total_tokens, _limit, _qualification)
}

// FillPool is a paid mutator transaction binding the contract method 0xef65dc7e.
//
// Solidity: function fill_pool(bytes32 _hash, uint256 _start, uint256 _end, string _message, address[] _exchange_addrs, uint128[] _ratios, uint256 _unlock_time, address _token_addr, uint256 _total_tokens, uint256 _limit, address _qualification) payable returns()
func (_HappyTokenPool *HappyTokenPoolTransactorSession) FillPool(_hash [32]byte, _start *big.Int, _end *big.Int, _message string, _exchange_addrs []common.Address, _ratios []*big.Int, _unlock_time *big.Int, _token_addr common.Address, _total_tokens *big.Int, _limit *big.Int, _qualification common.Address) (*types.Transaction, error) {
	return _HappyTokenPool.Contract.FillPool(&_HappyTokenPool.TransactOpts, _hash, _start, _end, _message, _exchange_addrs, _ratios, _unlock_time, _token_addr, _total_tokens, _limit, _qualification)
}

// Initialize is a paid mutator transaction binding the contract method 0xeb92db27.
//
// Solidity: function initialize(uint64 _base_time) returns()
func (_HappyTokenPool *HappyTokenPoolTransactor) Initialize(opts *bind.TransactOpts, _base_time uint64) (*types.Transaction, error) {
	return _HappyTokenPool.contract.Transact(opts, "initialize", _base_time)
}

// Initialize is a paid mutator transaction binding the contract method 0xeb92db27.
//
// Solidity: function initialize(uint64 _base_time) returns()
func (_HappyTokenPool *HappyTokenPoolSession) Initialize(_base_time uint64) (*types.Transaction, error) {
	return _HappyTokenPool.Contract.Initialize(&_HappyTokenPool.TransactOpts, _base_time)
}

// Initialize is a paid mutator transaction binding the contract method 0xeb92db27.
//
// Solidity: function initialize(uint64 _base_time) returns()
func (_HappyTokenPool *HappyTokenPoolTransactorSession) Initialize(_base_time uint64) (*types.Transaction, error) {
	return _HappyTokenPool.Contract.Initialize(&_HappyTokenPool.TransactOpts, _base_time)
}

// SetUnlockTime is a paid mutator transaction binding the contract method 0x69c12f10.
//
// Solidity: function setUnlockTime(bytes32 id, uint32 _unlock_time) returns()
func (_HappyTokenPool *HappyTokenPoolTransactor) SetUnlockTime(opts *bind.TransactOpts, id [32]byte, _unlock_time uint32) (*types.Transaction, error) {
	return _HappyTokenPool.contract.Transact(opts, "setUnlockTime", id, _unlock_time)
}

// SetUnlockTime is a paid mutator transaction binding the contract method 0x69c12f10.
//
// Solidity: function setUnlockTime(bytes32 id, uint32 _unlock_time) returns()
func (_HappyTokenPool *HappyTokenPoolSession) SetUnlockTime(id [32]byte, _unlock_time uint32) (*types.Transaction, error) {
	return _HappyTokenPool.Contract.SetUnlockTime(&_HappyTokenPool.TransactOpts, id, _unlock_time)
}

// SetUnlockTime is a paid mutator transaction binding the contract method 0x69c12f10.
//
// Solidity: function setUnlockTime(bytes32 id, uint32 _unlock_time) returns()
func (_HappyTokenPool *HappyTokenPoolTransactorSession) SetUnlockTime(id [32]byte, _unlock_time uint32) (*types.Transaction, error) {
	return _HappyTokenPool.Contract.SetUnlockTime(&_HappyTokenPool.TransactOpts, id, _unlock_time)
}

// Swap is a paid mutator transaction binding the contract method 0x17039d3b.
//
// Solidity: function swap(bytes32 id, bytes32 verification, uint256 exchange_addr_i, uint128 input_total, bytes32[] data) payable returns(uint256 swapped)
func (_HappyTokenPool *HappyTokenPoolTransactor) Swap(opts *bind.TransactOpts, id [32]byte, verification [32]byte, exchange_addr_i *big.Int, input_total *big.Int, data [][32]byte) (*types.Transaction, error) {
	return _HappyTokenPool.contract.Transact(opts, "swap", id, verification, exchange_addr_i, input_total, data)
}

// Swap is a paid mutator transaction binding the contract method 0x17039d3b.
//
// Solidity: function swap(bytes32 id, bytes32 verification, uint256 exchange_addr_i, uint128 input_total, bytes32[] data) payable returns(uint256 swapped)
func (_HappyTokenPool *HappyTokenPoolSession) Swap(id [32]byte, verification [32]byte, exchange_addr_i *big.Int, input_total *big.Int, data [][32]byte) (*types.Transaction, error) {
	return _HappyTokenPool.Contract.Swap(&_HappyTokenPool.TransactOpts, id, verification, exchange_addr_i, input_total, data)
}

// Swap is a paid mutator transaction binding the contract method 0x17039d3b.
//
// Solidity: function swap(bytes32 id, bytes32 verification, uint256 exchange_addr_i, uint128 input_total, bytes32[] data) payable returns(uint256 swapped)
func (_HappyTokenPool *HappyTokenPoolTransactorSession) Swap(id [32]byte, verification [32]byte, exchange_addr_i *big.Int, input_total *big.Int, data [][32]byte) (*types.Transaction, error) {
	return _HappyTokenPool.Contract.Swap(&_HappyTokenPool.TransactOpts, id, verification, exchange_addr_i, input_total, data)
}

// HappyTokenPoolClaimSuccessIterator is returned from FilterClaimSuccess and is used to iterate over the raw logs and unpacked data for ClaimSuccess events raised by the HappyTokenPool contract.
type HappyTokenPoolClaimSuccessIterator struct {
	Event *HappyTokenPoolClaimSuccess // Event containing the contract specifics and raw log

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
func (it *HappyTokenPoolClaimSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HappyTokenPoolClaimSuccess)
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
		it.Event = new(HappyTokenPoolClaimSuccess)
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
func (it *HappyTokenPoolClaimSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HappyTokenPoolClaimSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HappyTokenPoolClaimSuccess represents a ClaimSuccess event raised by the HappyTokenPool contract.
type HappyTokenPoolClaimSuccess struct {
	Id           [32]byte
	Claimer      common.Address
	Timestamp    *big.Int
	ToValue      *big.Int
	TokenAddress common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterClaimSuccess is a free log retrieval operation binding the contract event 0x0b8bdb11ddec1f8dd879bd98afac7c33c30ce590fcbc26ed559df98f9a381119.
//
// Solidity: event ClaimSuccess(bytes32 indexed id, address indexed claimer, uint256 timestamp, uint256 to_value, address token_address)
func (_HappyTokenPool *HappyTokenPoolFilterer) FilterClaimSuccess(opts *bind.FilterOpts, id [][32]byte, claimer []common.Address) (*HappyTokenPoolClaimSuccessIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}

	logs, sub, err := _HappyTokenPool.contract.FilterLogs(opts, "ClaimSuccess", idRule, claimerRule)
	if err != nil {
		return nil, err
	}
	return &HappyTokenPoolClaimSuccessIterator{contract: _HappyTokenPool.contract, event: "ClaimSuccess", logs: logs, sub: sub}, nil
}

// WatchClaimSuccess is a free log subscription operation binding the contract event 0x0b8bdb11ddec1f8dd879bd98afac7c33c30ce590fcbc26ed559df98f9a381119.
//
// Solidity: event ClaimSuccess(bytes32 indexed id, address indexed claimer, uint256 timestamp, uint256 to_value, address token_address)
func (_HappyTokenPool *HappyTokenPoolFilterer) WatchClaimSuccess(opts *bind.WatchOpts, sink chan<- *HappyTokenPoolClaimSuccess, id [][32]byte, claimer []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}

	logs, sub, err := _HappyTokenPool.contract.WatchLogs(opts, "ClaimSuccess", idRule, claimerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HappyTokenPoolClaimSuccess)
				if err := _HappyTokenPool.contract.UnpackLog(event, "ClaimSuccess", log); err != nil {
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

// ParseClaimSuccess is a log parse operation binding the contract event 0x0b8bdb11ddec1f8dd879bd98afac7c33c30ce590fcbc26ed559df98f9a381119.
//
// Solidity: event ClaimSuccess(bytes32 indexed id, address indexed claimer, uint256 timestamp, uint256 to_value, address token_address)
func (_HappyTokenPool *HappyTokenPoolFilterer) ParseClaimSuccess(log types.Log) (*HappyTokenPoolClaimSuccess, error) {
	event := new(HappyTokenPoolClaimSuccess)
	if err := _HappyTokenPool.contract.UnpackLog(event, "ClaimSuccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HappyTokenPoolDestructSuccessIterator is returned from FilterDestructSuccess and is used to iterate over the raw logs and unpacked data for DestructSuccess events raised by the HappyTokenPool contract.
type HappyTokenPoolDestructSuccessIterator struct {
	Event *HappyTokenPoolDestructSuccess // Event containing the contract specifics and raw log

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
func (it *HappyTokenPoolDestructSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HappyTokenPoolDestructSuccess)
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
		it.Event = new(HappyTokenPoolDestructSuccess)
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
func (it *HappyTokenPoolDestructSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HappyTokenPoolDestructSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HappyTokenPoolDestructSuccess represents a DestructSuccess event raised by the HappyTokenPool contract.
type HappyTokenPoolDestructSuccess struct {
	Id               [32]byte
	TokenAddress     common.Address
	RemainingBalance *big.Int
	ExchangedValues  []*big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterDestructSuccess is a free log retrieval operation binding the contract event 0x01a1fd9d8a4231d7bd8926ce16b753ba03e0310e0a776f13eba8e0c3cba7ea89.
//
// Solidity: event DestructSuccess(bytes32 indexed id, address indexed token_address, uint256 remaining_balance, uint128[] exchanged_values)
func (_HappyTokenPool *HappyTokenPoolFilterer) FilterDestructSuccess(opts *bind.FilterOpts, id [][32]byte, token_address []common.Address) (*HappyTokenPoolDestructSuccessIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var token_addressRule []interface{}
	for _, token_addressItem := range token_address {
		token_addressRule = append(token_addressRule, token_addressItem)
	}

	logs, sub, err := _HappyTokenPool.contract.FilterLogs(opts, "DestructSuccess", idRule, token_addressRule)
	if err != nil {
		return nil, err
	}
	return &HappyTokenPoolDestructSuccessIterator{contract: _HappyTokenPool.contract, event: "DestructSuccess", logs: logs, sub: sub}, nil
}

// WatchDestructSuccess is a free log subscription operation binding the contract event 0x01a1fd9d8a4231d7bd8926ce16b753ba03e0310e0a776f13eba8e0c3cba7ea89.
//
// Solidity: event DestructSuccess(bytes32 indexed id, address indexed token_address, uint256 remaining_balance, uint128[] exchanged_values)
func (_HappyTokenPool *HappyTokenPoolFilterer) WatchDestructSuccess(opts *bind.WatchOpts, sink chan<- *HappyTokenPoolDestructSuccess, id [][32]byte, token_address []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var token_addressRule []interface{}
	for _, token_addressItem := range token_address {
		token_addressRule = append(token_addressRule, token_addressItem)
	}

	logs, sub, err := _HappyTokenPool.contract.WatchLogs(opts, "DestructSuccess", idRule, token_addressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HappyTokenPoolDestructSuccess)
				if err := _HappyTokenPool.contract.UnpackLog(event, "DestructSuccess", log); err != nil {
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

// ParseDestructSuccess is a log parse operation binding the contract event 0x01a1fd9d8a4231d7bd8926ce16b753ba03e0310e0a776f13eba8e0c3cba7ea89.
//
// Solidity: event DestructSuccess(bytes32 indexed id, address indexed token_address, uint256 remaining_balance, uint128[] exchanged_values)
func (_HappyTokenPool *HappyTokenPoolFilterer) ParseDestructSuccess(log types.Log) (*HappyTokenPoolDestructSuccess, error) {
	event := new(HappyTokenPoolDestructSuccess)
	if err := _HappyTokenPool.contract.UnpackLog(event, "DestructSuccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HappyTokenPoolFillSuccessIterator is returned from FilterFillSuccess and is used to iterate over the raw logs and unpacked data for FillSuccess events raised by the HappyTokenPool contract.
type HappyTokenPoolFillSuccessIterator struct {
	Event *HappyTokenPoolFillSuccess // Event containing the contract specifics and raw log

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
func (it *HappyTokenPoolFillSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HappyTokenPoolFillSuccess)
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
		it.Event = new(HappyTokenPoolFillSuccess)
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
func (it *HappyTokenPoolFillSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HappyTokenPoolFillSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HappyTokenPoolFillSuccess represents a FillSuccess event raised by the HappyTokenPool contract.
type HappyTokenPoolFillSuccess struct {
	Creator       common.Address
	Id            [32]byte
	Total         *big.Int
	CreationTime  *big.Int
	TokenAddress  common.Address
	Message       string
	Start         *big.Int
	End           *big.Int
	ExchangeAddrs []common.Address
	Ratios        []*big.Int
	Qualification common.Address
	Limit         *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFillSuccess is a free log retrieval operation binding the contract event 0xbd9071dfd55143304371945d01e62c79ded9810323a79809ba3244e7b25856c2.
//
// Solidity: event FillSuccess(address indexed creator, bytes32 indexed id, uint256 total, uint256 creation_time, address token_address, string message, uint256 start, uint256 end, address[] exchange_addrs, uint128[] ratios, address qualification, uint256 limit)
func (_HappyTokenPool *HappyTokenPoolFilterer) FilterFillSuccess(opts *bind.FilterOpts, creator []common.Address, id [][32]byte) (*HappyTokenPoolFillSuccessIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _HappyTokenPool.contract.FilterLogs(opts, "FillSuccess", creatorRule, idRule)
	if err != nil {
		return nil, err
	}
	return &HappyTokenPoolFillSuccessIterator{contract: _HappyTokenPool.contract, event: "FillSuccess", logs: logs, sub: sub}, nil
}

// WatchFillSuccess is a free log subscription operation binding the contract event 0xbd9071dfd55143304371945d01e62c79ded9810323a79809ba3244e7b25856c2.
//
// Solidity: event FillSuccess(address indexed creator, bytes32 indexed id, uint256 total, uint256 creation_time, address token_address, string message, uint256 start, uint256 end, address[] exchange_addrs, uint128[] ratios, address qualification, uint256 limit)
func (_HappyTokenPool *HappyTokenPoolFilterer) WatchFillSuccess(opts *bind.WatchOpts, sink chan<- *HappyTokenPoolFillSuccess, creator []common.Address, id [][32]byte) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _HappyTokenPool.contract.WatchLogs(opts, "FillSuccess", creatorRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HappyTokenPoolFillSuccess)
				if err := _HappyTokenPool.contract.UnpackLog(event, "FillSuccess", log); err != nil {
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

// ParseFillSuccess is a log parse operation binding the contract event 0xbd9071dfd55143304371945d01e62c79ded9810323a79809ba3244e7b25856c2.
//
// Solidity: event FillSuccess(address indexed creator, bytes32 indexed id, uint256 total, uint256 creation_time, address token_address, string message, uint256 start, uint256 end, address[] exchange_addrs, uint128[] ratios, address qualification, uint256 limit)
func (_HappyTokenPool *HappyTokenPoolFilterer) ParseFillSuccess(log types.Log) (*HappyTokenPoolFillSuccess, error) {
	event := new(HappyTokenPoolFillSuccess)
	if err := _HappyTokenPool.contract.UnpackLog(event, "FillSuccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HappyTokenPoolSwapSuccessIterator is returned from FilterSwapSuccess and is used to iterate over the raw logs and unpacked data for SwapSuccess events raised by the HappyTokenPool contract.
type HappyTokenPoolSwapSuccessIterator struct {
	Event *HappyTokenPoolSwapSuccess // Event containing the contract specifics and raw log

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
func (it *HappyTokenPoolSwapSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HappyTokenPoolSwapSuccess)
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
		it.Event = new(HappyTokenPoolSwapSuccess)
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
func (it *HappyTokenPoolSwapSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HappyTokenPoolSwapSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HappyTokenPoolSwapSuccess represents a SwapSuccess event raised by the HappyTokenPool contract.
type HappyTokenPoolSwapSuccess struct {
	Id          [32]byte
	Swapper     common.Address
	FromAddress common.Address
	ToAddress   common.Address
	FromValue   *big.Int
	ToValue     *big.Int
	InputTotal  *big.Int
	Claimed     bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSwapSuccess is a free log retrieval operation binding the contract event 0x4c086e5fed134e5604e0e8e75908ae2d7db743ec0d4ee3f0d9a279ce2213279c.
//
// Solidity: event SwapSuccess(bytes32 indexed id, address indexed swapper, address from_address, address to_address, uint256 from_value, uint256 to_value, uint128 input_total, bool claimed)
func (_HappyTokenPool *HappyTokenPoolFilterer) FilterSwapSuccess(opts *bind.FilterOpts, id [][32]byte, swapper []common.Address) (*HappyTokenPoolSwapSuccessIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var swapperRule []interface{}
	for _, swapperItem := range swapper {
		swapperRule = append(swapperRule, swapperItem)
	}

	logs, sub, err := _HappyTokenPool.contract.FilterLogs(opts, "SwapSuccess", idRule, swapperRule)
	if err != nil {
		return nil, err
	}
	return &HappyTokenPoolSwapSuccessIterator{contract: _HappyTokenPool.contract, event: "SwapSuccess", logs: logs, sub: sub}, nil
}

// WatchSwapSuccess is a free log subscription operation binding the contract event 0x4c086e5fed134e5604e0e8e75908ae2d7db743ec0d4ee3f0d9a279ce2213279c.
//
// Solidity: event SwapSuccess(bytes32 indexed id, address indexed swapper, address from_address, address to_address, uint256 from_value, uint256 to_value, uint128 input_total, bool claimed)
func (_HappyTokenPool *HappyTokenPoolFilterer) WatchSwapSuccess(opts *bind.WatchOpts, sink chan<- *HappyTokenPoolSwapSuccess, id [][32]byte, swapper []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var swapperRule []interface{}
	for _, swapperItem := range swapper {
		swapperRule = append(swapperRule, swapperItem)
	}

	logs, sub, err := _HappyTokenPool.contract.WatchLogs(opts, "SwapSuccess", idRule, swapperRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HappyTokenPoolSwapSuccess)
				if err := _HappyTokenPool.contract.UnpackLog(event, "SwapSuccess", log); err != nil {
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

// ParseSwapSuccess is a log parse operation binding the contract event 0x4c086e5fed134e5604e0e8e75908ae2d7db743ec0d4ee3f0d9a279ce2213279c.
//
// Solidity: event SwapSuccess(bytes32 indexed id, address indexed swapper, address from_address, address to_address, uint256 from_value, uint256 to_value, uint128 input_total, bool claimed)
func (_HappyTokenPool *HappyTokenPoolFilterer) ParseSwapSuccess(log types.Log) (*HappyTokenPoolSwapSuccess, error) {
	event := new(HappyTokenPoolSwapSuccess)
	if err := _HappyTokenPool.contract.UnpackLog(event, "SwapSuccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HappyTokenPoolWithdrawSuccessIterator is returned from FilterWithdrawSuccess and is used to iterate over the raw logs and unpacked data for WithdrawSuccess events raised by the HappyTokenPool contract.
type HappyTokenPoolWithdrawSuccessIterator struct {
	Event *HappyTokenPoolWithdrawSuccess // Event containing the contract specifics and raw log

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
func (it *HappyTokenPoolWithdrawSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HappyTokenPoolWithdrawSuccess)
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
		it.Event = new(HappyTokenPoolWithdrawSuccess)
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
func (it *HappyTokenPoolWithdrawSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HappyTokenPoolWithdrawSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HappyTokenPoolWithdrawSuccess represents a WithdrawSuccess event raised by the HappyTokenPool contract.
type HappyTokenPoolWithdrawSuccess struct {
	Id              [32]byte
	TokenAddress    common.Address
	WithdrawBalance *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterWithdrawSuccess is a free log retrieval operation binding the contract event 0xab5315d16ef405a1f4e8c34017af486f59f27097e003d2a1981ae682c2f36731.
//
// Solidity: event WithdrawSuccess(bytes32 indexed id, address indexed token_address, uint256 withdraw_balance)
func (_HappyTokenPool *HappyTokenPoolFilterer) FilterWithdrawSuccess(opts *bind.FilterOpts, id [][32]byte, token_address []common.Address) (*HappyTokenPoolWithdrawSuccessIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var token_addressRule []interface{}
	for _, token_addressItem := range token_address {
		token_addressRule = append(token_addressRule, token_addressItem)
	}

	logs, sub, err := _HappyTokenPool.contract.FilterLogs(opts, "WithdrawSuccess", idRule, token_addressRule)
	if err != nil {
		return nil, err
	}
	return &HappyTokenPoolWithdrawSuccessIterator{contract: _HappyTokenPool.contract, event: "WithdrawSuccess", logs: logs, sub: sub}, nil
}

// WatchWithdrawSuccess is a free log subscription operation binding the contract event 0xab5315d16ef405a1f4e8c34017af486f59f27097e003d2a1981ae682c2f36731.
//
// Solidity: event WithdrawSuccess(bytes32 indexed id, address indexed token_address, uint256 withdraw_balance)
func (_HappyTokenPool *HappyTokenPoolFilterer) WatchWithdrawSuccess(opts *bind.WatchOpts, sink chan<- *HappyTokenPoolWithdrawSuccess, id [][32]byte, token_address []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var token_addressRule []interface{}
	for _, token_addressItem := range token_address {
		token_addressRule = append(token_addressRule, token_addressItem)
	}

	logs, sub, err := _HappyTokenPool.contract.WatchLogs(opts, "WithdrawSuccess", idRule, token_addressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HappyTokenPoolWithdrawSuccess)
				if err := _HappyTokenPool.contract.UnpackLog(event, "WithdrawSuccess", log); err != nil {
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

// ParseWithdrawSuccess is a log parse operation binding the contract event 0xab5315d16ef405a1f4e8c34017af486f59f27097e003d2a1981ae682c2f36731.
//
// Solidity: event WithdrawSuccess(bytes32 indexed id, address indexed token_address, uint256 withdraw_balance)
func (_HappyTokenPool *HappyTokenPoolFilterer) ParseWithdrawSuccess(log types.Log) (*HappyTokenPoolWithdrawSuccess, error) {
	event := new(HappyTokenPoolWithdrawSuccess)
	if err := _HappyTokenPool.contract.UnpackLog(event, "WithdrawSuccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
