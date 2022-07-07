// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mrc20

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

// MRC20MetaData contains all meta data concerning the MRC20 contract.
var MRC20MetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"setParent\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"data\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferWithSig\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_childChain\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"parent\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"parentOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"ecrecovery\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"result\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"networkId\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"EIP712_TOKEN_TRANSFER_ORDER_SCHEMA_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"disabledHashes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenIdOrAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"data\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"}],\"name\":\"getTokenTransferOrderHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CHAINID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"EIP712_DOMAIN_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"EIP712_DOMAIN_SCHEMA_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"transaction\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transaction\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"input1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"output1\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transaction\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"input1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"output1\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transaction\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"input1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"input2\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"output1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"output2\",\"type\":\"uint256\"}],\"name\":\"LogTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transaction\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"input1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"input2\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"output1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"output2\",\"type\":\"uint256\"}],\"name\":\"LogFeeTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]",
}

// MRC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use MRC20MetaData.ABI instead.
var MRC20ABI = MRC20MetaData.ABI

// MRC20 is an auto generated Go binding around an Ethereum contract.
type MRC20 struct {
	MRC20Caller     // Read-only binding to the contract
	MRC20Transactor // Write-only binding to the contract
	MRC20Filterer   // Log filterer for contract events
}

// MRC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type MRC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MRC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type MRC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MRC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MRC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MRC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MRC20Session struct {
	Contract     *MRC20            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MRC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MRC20CallerSession struct {
	Contract *MRC20Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MRC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MRC20TransactorSession struct {
	Contract     *MRC20Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MRC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type MRC20Raw struct {
	Contract *MRC20 // Generic contract binding to access the raw methods on
}

// MRC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MRC20CallerRaw struct {
	Contract *MRC20Caller // Generic read-only contract binding to access the raw methods on
}

// MRC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MRC20TransactorRaw struct {
	Contract *MRC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewMRC20 creates a new instance of MRC20, bound to a specific deployed contract.
func NewMRC20(address common.Address, backend bind.ContractBackend) (*MRC20, error) {
	contract, err := bindMRC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MRC20{MRC20Caller: MRC20Caller{contract: contract}, MRC20Transactor: MRC20Transactor{contract: contract}, MRC20Filterer: MRC20Filterer{contract: contract}}, nil
}

// NewMRC20Caller creates a new read-only instance of MRC20, bound to a specific deployed contract.
func NewMRC20Caller(address common.Address, caller bind.ContractCaller) (*MRC20Caller, error) {
	contract, err := bindMRC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MRC20Caller{contract: contract}, nil
}

// NewMRC20Transactor creates a new write-only instance of MRC20, bound to a specific deployed contract.
func NewMRC20Transactor(address common.Address, transactor bind.ContractTransactor) (*MRC20Transactor, error) {
	contract, err := bindMRC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MRC20Transactor{contract: contract}, nil
}

// NewMRC20Filterer creates a new log filterer instance of MRC20, bound to a specific deployed contract.
func NewMRC20Filterer(address common.Address, filterer bind.ContractFilterer) (*MRC20Filterer, error) {
	contract, err := bindMRC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MRC20Filterer{contract: contract}, nil
}

// bindMRC20 binds a generic wrapper to an already deployed contract.
func bindMRC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MRC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MRC20 *MRC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MRC20.Contract.MRC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MRC20 *MRC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MRC20.Contract.MRC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MRC20 *MRC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MRC20.Contract.MRC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MRC20 *MRC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MRC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MRC20 *MRC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MRC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MRC20 *MRC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MRC20.Contract.contract.Transact(opts, method, params...)
}

// CHAINID is a free data retrieval call binding the contract method 0xcc79f97b.
//
// Solidity: function CHAINID() view returns(uint256)
func (_MRC20 *MRC20Caller) CHAINID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MRC20.contract.Call(opts, &out, "CHAINID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CHAINID is a free data retrieval call binding the contract method 0xcc79f97b.
//
// Solidity: function CHAINID() view returns(uint256)
func (_MRC20 *MRC20Session) CHAINID() (*big.Int, error) {
	return _MRC20.Contract.CHAINID(&_MRC20.CallOpts)
}

// CHAINID is a free data retrieval call binding the contract method 0xcc79f97b.
//
// Solidity: function CHAINID() view returns(uint256)
func (_MRC20 *MRC20CallerSession) CHAINID() (*big.Int, error) {
	return _MRC20.Contract.CHAINID(&_MRC20.CallOpts)
}

// EIP712DOMAINHASH is a free data retrieval call binding the contract method 0xe306f779.
//
// Solidity: function EIP712_DOMAIN_HASH() view returns(bytes32)
func (_MRC20 *MRC20Caller) EIP712DOMAINHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MRC20.contract.Call(opts, &out, "EIP712_DOMAIN_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EIP712DOMAINHASH is a free data retrieval call binding the contract method 0xe306f779.
//
// Solidity: function EIP712_DOMAIN_HASH() view returns(bytes32)
func (_MRC20 *MRC20Session) EIP712DOMAINHASH() ([32]byte, error) {
	return _MRC20.Contract.EIP712DOMAINHASH(&_MRC20.CallOpts)
}

// EIP712DOMAINHASH is a free data retrieval call binding the contract method 0xe306f779.
//
// Solidity: function EIP712_DOMAIN_HASH() view returns(bytes32)
func (_MRC20 *MRC20CallerSession) EIP712DOMAINHASH() ([32]byte, error) {
	return _MRC20.Contract.EIP712DOMAINHASH(&_MRC20.CallOpts)
}

// EIP712DOMAINSCHEMAHASH is a free data retrieval call binding the contract method 0xe614d0d6.
//
// Solidity: function EIP712_DOMAIN_SCHEMA_HASH() view returns(bytes32)
func (_MRC20 *MRC20Caller) EIP712DOMAINSCHEMAHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MRC20.contract.Call(opts, &out, "EIP712_DOMAIN_SCHEMA_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EIP712DOMAINSCHEMAHASH is a free data retrieval call binding the contract method 0xe614d0d6.
//
// Solidity: function EIP712_DOMAIN_SCHEMA_HASH() view returns(bytes32)
func (_MRC20 *MRC20Session) EIP712DOMAINSCHEMAHASH() ([32]byte, error) {
	return _MRC20.Contract.EIP712DOMAINSCHEMAHASH(&_MRC20.CallOpts)
}

// EIP712DOMAINSCHEMAHASH is a free data retrieval call binding the contract method 0xe614d0d6.
//
// Solidity: function EIP712_DOMAIN_SCHEMA_HASH() view returns(bytes32)
func (_MRC20 *MRC20CallerSession) EIP712DOMAINSCHEMAHASH() ([32]byte, error) {
	return _MRC20.Contract.EIP712DOMAINSCHEMAHASH(&_MRC20.CallOpts)
}

// EIP712TOKENTRANSFERORDERSCHEMAHASH is a free data retrieval call binding the contract method 0xabceeba2.
//
// Solidity: function EIP712_TOKEN_TRANSFER_ORDER_SCHEMA_HASH() view returns(bytes32)
func (_MRC20 *MRC20Caller) EIP712TOKENTRANSFERORDERSCHEMAHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MRC20.contract.Call(opts, &out, "EIP712_TOKEN_TRANSFER_ORDER_SCHEMA_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EIP712TOKENTRANSFERORDERSCHEMAHASH is a free data retrieval call binding the contract method 0xabceeba2.
//
// Solidity: function EIP712_TOKEN_TRANSFER_ORDER_SCHEMA_HASH() view returns(bytes32)
func (_MRC20 *MRC20Session) EIP712TOKENTRANSFERORDERSCHEMAHASH() ([32]byte, error) {
	return _MRC20.Contract.EIP712TOKENTRANSFERORDERSCHEMAHASH(&_MRC20.CallOpts)
}

// EIP712TOKENTRANSFERORDERSCHEMAHASH is a free data retrieval call binding the contract method 0xabceeba2.
//
// Solidity: function EIP712_TOKEN_TRANSFER_ORDER_SCHEMA_HASH() view returns(bytes32)
func (_MRC20 *MRC20CallerSession) EIP712TOKENTRANSFERORDERSCHEMAHASH() ([32]byte, error) {
	return _MRC20.Contract.EIP712TOKENTRANSFERORDERSCHEMAHASH(&_MRC20.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MRC20 *MRC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MRC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MRC20 *MRC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _MRC20.Contract.BalanceOf(&_MRC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MRC20 *MRC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _MRC20.Contract.BalanceOf(&_MRC20.CallOpts, account)
}

// CurrentSupply is a free data retrieval call binding the contract method 0x771282f6.
//
// Solidity: function currentSupply() view returns(uint256)
func (_MRC20 *MRC20Caller) CurrentSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MRC20.contract.Call(opts, &out, "currentSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentSupply is a free data retrieval call binding the contract method 0x771282f6.
//
// Solidity: function currentSupply() view returns(uint256)
func (_MRC20 *MRC20Session) CurrentSupply() (*big.Int, error) {
	return _MRC20.Contract.CurrentSupply(&_MRC20.CallOpts)
}

// CurrentSupply is a free data retrieval call binding the contract method 0x771282f6.
//
// Solidity: function currentSupply() view returns(uint256)
func (_MRC20 *MRC20CallerSession) CurrentSupply() (*big.Int, error) {
	return _MRC20.Contract.CurrentSupply(&_MRC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MRC20 *MRC20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _MRC20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MRC20 *MRC20Session) Decimals() (uint8, error) {
	return _MRC20.Contract.Decimals(&_MRC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MRC20 *MRC20CallerSession) Decimals() (uint8, error) {
	return _MRC20.Contract.Decimals(&_MRC20.CallOpts)
}

// DisabledHashes is a free data retrieval call binding the contract method 0xacd06cb3.
//
// Solidity: function disabledHashes(bytes32 ) view returns(bool)
func (_MRC20 *MRC20Caller) DisabledHashes(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _MRC20.contract.Call(opts, &out, "disabledHashes", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DisabledHashes is a free data retrieval call binding the contract method 0xacd06cb3.
//
// Solidity: function disabledHashes(bytes32 ) view returns(bool)
func (_MRC20 *MRC20Session) DisabledHashes(arg0 [32]byte) (bool, error) {
	return _MRC20.Contract.DisabledHashes(&_MRC20.CallOpts, arg0)
}

// DisabledHashes is a free data retrieval call binding the contract method 0xacd06cb3.
//
// Solidity: function disabledHashes(bytes32 ) view returns(bool)
func (_MRC20 *MRC20CallerSession) DisabledHashes(arg0 [32]byte) (bool, error) {
	return _MRC20.Contract.DisabledHashes(&_MRC20.CallOpts, arg0)
}

// Ecrecovery is a free data retrieval call binding the contract method 0x77d32e94.
//
// Solidity: function ecrecovery(bytes32 hash, bytes sig) view returns(address result)
func (_MRC20 *MRC20Caller) Ecrecovery(opts *bind.CallOpts, hash [32]byte, sig []byte) (common.Address, error) {
	var out []interface{}
	err := _MRC20.contract.Call(opts, &out, "ecrecovery", hash, sig)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Ecrecovery is a free data retrieval call binding the contract method 0x77d32e94.
//
// Solidity: function ecrecovery(bytes32 hash, bytes sig) view returns(address result)
func (_MRC20 *MRC20Session) Ecrecovery(hash [32]byte, sig []byte) (common.Address, error) {
	return _MRC20.Contract.Ecrecovery(&_MRC20.CallOpts, hash, sig)
}

// Ecrecovery is a free data retrieval call binding the contract method 0x77d32e94.
//
// Solidity: function ecrecovery(bytes32 hash, bytes sig) view returns(address result)
func (_MRC20 *MRC20CallerSession) Ecrecovery(hash [32]byte, sig []byte) (common.Address, error) {
	return _MRC20.Contract.Ecrecovery(&_MRC20.CallOpts, hash, sig)
}

// GetTokenTransferOrderHash is a free data retrieval call binding the contract method 0xb789543c.
//
// Solidity: function getTokenTransferOrderHash(address spender, uint256 tokenIdOrAmount, bytes32 data, uint256 expiration) view returns(bytes32 orderHash)
func (_MRC20 *MRC20Caller) GetTokenTransferOrderHash(opts *bind.CallOpts, spender common.Address, tokenIdOrAmount *big.Int, data [32]byte, expiration *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _MRC20.contract.Call(opts, &out, "getTokenTransferOrderHash", spender, tokenIdOrAmount, data, expiration)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetTokenTransferOrderHash is a free data retrieval call binding the contract method 0xb789543c.
//
// Solidity: function getTokenTransferOrderHash(address spender, uint256 tokenIdOrAmount, bytes32 data, uint256 expiration) view returns(bytes32 orderHash)
func (_MRC20 *MRC20Session) GetTokenTransferOrderHash(spender common.Address, tokenIdOrAmount *big.Int, data [32]byte, expiration *big.Int) ([32]byte, error) {
	return _MRC20.Contract.GetTokenTransferOrderHash(&_MRC20.CallOpts, spender, tokenIdOrAmount, data, expiration)
}

// GetTokenTransferOrderHash is a free data retrieval call binding the contract method 0xb789543c.
//
// Solidity: function getTokenTransferOrderHash(address spender, uint256 tokenIdOrAmount, bytes32 data, uint256 expiration) view returns(bytes32 orderHash)
func (_MRC20 *MRC20CallerSession) GetTokenTransferOrderHash(spender common.Address, tokenIdOrAmount *big.Int, data [32]byte, expiration *big.Int) ([32]byte, error) {
	return _MRC20.Contract.GetTokenTransferOrderHash(&_MRC20.CallOpts, spender, tokenIdOrAmount, data, expiration)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_MRC20 *MRC20Caller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MRC20.contract.Call(opts, &out, "isOwner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_MRC20 *MRC20Session) IsOwner() (bool, error) {
	return _MRC20.Contract.IsOwner(&_MRC20.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_MRC20 *MRC20CallerSession) IsOwner() (bool, error) {
	return _MRC20.Contract.IsOwner(&_MRC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MRC20 *MRC20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MRC20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MRC20 *MRC20Session) Name() (string, error) {
	return _MRC20.Contract.Name(&_MRC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MRC20 *MRC20CallerSession) Name() (string, error) {
	return _MRC20.Contract.Name(&_MRC20.CallOpts)
}

// NetworkId is a free data retrieval call binding the contract method 0x9025e64c.
//
// Solidity: function networkId() view returns(bytes)
func (_MRC20 *MRC20Caller) NetworkId(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _MRC20.contract.Call(opts, &out, "networkId")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// NetworkId is a free data retrieval call binding the contract method 0x9025e64c.
//
// Solidity: function networkId() view returns(bytes)
func (_MRC20 *MRC20Session) NetworkId() ([]byte, error) {
	return _MRC20.Contract.NetworkId(&_MRC20.CallOpts)
}

// NetworkId is a free data retrieval call binding the contract method 0x9025e64c.
//
// Solidity: function networkId() view returns(bytes)
func (_MRC20 *MRC20CallerSession) NetworkId() ([]byte, error) {
	return _MRC20.Contract.NetworkId(&_MRC20.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MRC20 *MRC20Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MRC20.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MRC20 *MRC20Session) Owner() (common.Address, error) {
	return _MRC20.Contract.Owner(&_MRC20.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MRC20 *MRC20CallerSession) Owner() (common.Address, error) {
	return _MRC20.Contract.Owner(&_MRC20.CallOpts)
}

// Parent is a free data retrieval call binding the contract method 0x60f96a8f.
//
// Solidity: function parent() view returns(address)
func (_MRC20 *MRC20Caller) Parent(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MRC20.contract.Call(opts, &out, "parent")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Parent is a free data retrieval call binding the contract method 0x60f96a8f.
//
// Solidity: function parent() view returns(address)
func (_MRC20 *MRC20Session) Parent() (common.Address, error) {
	return _MRC20.Contract.Parent(&_MRC20.CallOpts)
}

// Parent is a free data retrieval call binding the contract method 0x60f96a8f.
//
// Solidity: function parent() view returns(address)
func (_MRC20 *MRC20CallerSession) Parent() (common.Address, error) {
	return _MRC20.Contract.Parent(&_MRC20.CallOpts)
}

// ParentOwner is a free data retrieval call binding the contract method 0x7019d41a.
//
// Solidity: function parentOwner() view returns(address)
func (_MRC20 *MRC20Caller) ParentOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MRC20.contract.Call(opts, &out, "parentOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ParentOwner is a free data retrieval call binding the contract method 0x7019d41a.
//
// Solidity: function parentOwner() view returns(address)
func (_MRC20 *MRC20Session) ParentOwner() (common.Address, error) {
	return _MRC20.Contract.ParentOwner(&_MRC20.CallOpts)
}

// ParentOwner is a free data retrieval call binding the contract method 0x7019d41a.
//
// Solidity: function parentOwner() view returns(address)
func (_MRC20 *MRC20CallerSession) ParentOwner() (common.Address, error) {
	return _MRC20.Contract.ParentOwner(&_MRC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MRC20 *MRC20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MRC20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MRC20 *MRC20Session) Symbol() (string, error) {
	return _MRC20.Contract.Symbol(&_MRC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MRC20 *MRC20CallerSession) Symbol() (string, error) {
	return _MRC20.Contract.Symbol(&_MRC20.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function transaction() view returns(address)
func (_MRC20 *MRC20Caller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MRC20.contract.Call(opts, &out, "transaction")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function transaction() view returns(address)
func (_MRC20 *MRC20Session) Token() (common.Address, error) {
	return _MRC20.Contract.Token(&_MRC20.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function transaction() view returns(address)
func (_MRC20 *MRC20CallerSession) Token() (common.Address, error) {
	return _MRC20.Contract.Token(&_MRC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MRC20 *MRC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MRC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MRC20 *MRC20Session) TotalSupply() (*big.Int, error) {
	return _MRC20.Contract.TotalSupply(&_MRC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MRC20 *MRC20CallerSession) TotalSupply() (*big.Int, error) {
	return _MRC20.Contract.TotalSupply(&_MRC20.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address user, uint256 amount) returns()
func (_MRC20 *MRC20Transactor) Deposit(opts *bind.TransactOpts, user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MRC20.contract.Transact(opts, "deposit", user, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address user, uint256 amount) returns()
func (_MRC20 *MRC20Session) Deposit(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MRC20.Contract.Deposit(&_MRC20.TransactOpts, user, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address user, uint256 amount) returns()
func (_MRC20 *MRC20TransactorSession) Deposit(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MRC20.Contract.Deposit(&_MRC20.TransactOpts, user, amount)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _childChain, address _token) returns()
func (_MRC20 *MRC20Transactor) Initialize(opts *bind.TransactOpts, _childChain common.Address, _token common.Address) (*types.Transaction, error) {
	return _MRC20.contract.Transact(opts, "initialize", _childChain, _token)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _childChain, address _token) returns()
func (_MRC20 *MRC20Session) Initialize(_childChain common.Address, _token common.Address) (*types.Transaction, error) {
	return _MRC20.Contract.Initialize(&_MRC20.TransactOpts, _childChain, _token)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _childChain, address _token) returns()
func (_MRC20 *MRC20TransactorSession) Initialize(_childChain common.Address, _token common.Address) (*types.Transaction, error) {
	return _MRC20.Contract.Initialize(&_MRC20.TransactOpts, _childChain, _token)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MRC20 *MRC20Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MRC20.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MRC20 *MRC20Session) RenounceOwnership() (*types.Transaction, error) {
	return _MRC20.Contract.RenounceOwnership(&_MRC20.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MRC20 *MRC20TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MRC20.Contract.RenounceOwnership(&_MRC20.TransactOpts)
}

// SetParent is a paid mutator transaction binding the contract method 0x1499c592.
//
// Solidity: function setParent(address ) returns()
func (_MRC20 *MRC20Transactor) SetParent(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _MRC20.contract.Transact(opts, "setParent", arg0)
}

// SetParent is a paid mutator transaction binding the contract method 0x1499c592.
//
// Solidity: function setParent(address ) returns()
func (_MRC20 *MRC20Session) SetParent(arg0 common.Address) (*types.Transaction, error) {
	return _MRC20.Contract.SetParent(&_MRC20.TransactOpts, arg0)
}

// SetParent is a paid mutator transaction binding the contract method 0x1499c592.
//
// Solidity: function setParent(address ) returns()
func (_MRC20 *MRC20TransactorSession) SetParent(arg0 common.Address) (*types.Transaction, error) {
	return _MRC20.Contract.SetParent(&_MRC20.TransactOpts, arg0)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) payable returns(bool)
func (_MRC20 *MRC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MRC20.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) payable returns(bool)
func (_MRC20 *MRC20Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MRC20.Contract.Transfer(&_MRC20.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) payable returns(bool)
func (_MRC20 *MRC20TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MRC20.Contract.Transfer(&_MRC20.TransactOpts, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MRC20 *MRC20Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MRC20.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MRC20 *MRC20Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MRC20.Contract.TransferOwnership(&_MRC20.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MRC20 *MRC20TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MRC20.Contract.TransferOwnership(&_MRC20.TransactOpts, newOwner)
}

// TransferWithSig is a paid mutator transaction binding the contract method 0x19d27d9c.
//
// Solidity: function transferWithSig(bytes sig, uint256 amount, bytes32 data, uint256 expiration, address to) returns(address from)
func (_MRC20 *MRC20Transactor) TransferWithSig(opts *bind.TransactOpts, sig []byte, amount *big.Int, data [32]byte, expiration *big.Int, to common.Address) (*types.Transaction, error) {
	return _MRC20.contract.Transact(opts, "transferWithSig", sig, amount, data, expiration, to)
}

// TransferWithSig is a paid mutator transaction binding the contract method 0x19d27d9c.
//
// Solidity: function transferWithSig(bytes sig, uint256 amount, bytes32 data, uint256 expiration, address to) returns(address from)
func (_MRC20 *MRC20Session) TransferWithSig(sig []byte, amount *big.Int, data [32]byte, expiration *big.Int, to common.Address) (*types.Transaction, error) {
	return _MRC20.Contract.TransferWithSig(&_MRC20.TransactOpts, sig, amount, data, expiration, to)
}

// TransferWithSig is a paid mutator transaction binding the contract method 0x19d27d9c.
//
// Solidity: function transferWithSig(bytes sig, uint256 amount, bytes32 data, uint256 expiration, address to) returns(address from)
func (_MRC20 *MRC20TransactorSession) TransferWithSig(sig []byte, amount *big.Int, data [32]byte, expiration *big.Int, to common.Address) (*types.Transaction, error) {
	return _MRC20.Contract.TransferWithSig(&_MRC20.TransactOpts, sig, amount, data, expiration, to)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) payable returns()
func (_MRC20 *MRC20Transactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _MRC20.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) payable returns()
func (_MRC20 *MRC20Session) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _MRC20.Contract.Withdraw(&_MRC20.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) payable returns()
func (_MRC20 *MRC20TransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _MRC20.Contract.Withdraw(&_MRC20.TransactOpts, amount)
}

// MRC20DepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the MRC20 contract.
type MRC20DepositIterator struct {
	Event *MRC20Deposit // Event containing the contract specifics and raw log

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
func (it *MRC20DepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MRC20Deposit)
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
		it.Event = new(MRC20Deposit)
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
func (it *MRC20DepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MRC20DepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MRC20Deposit represents a Deposit event raised by the MRC20 contract.
type MRC20Deposit struct {
	Token   common.Address
	From    common.Address
	Amount  *big.Int
	Input1  *big.Int
	Output1 *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x4e2ca0515ed1aef1395f66b5303bb5d6f1bf9d61a353fa53f73f8ac9973fa9f6.
//
// Solidity: event Deposit(address indexed transaction, address indexed from, uint256 amount, uint256 input1, uint256 output1)
func (_MRC20 *MRC20Filterer) FilterDeposit(opts *bind.FilterOpts, token []common.Address, from []common.Address) (*MRC20DepositIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _MRC20.contract.FilterLogs(opts, "Deposit", tokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &MRC20DepositIterator{contract: _MRC20.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x4e2ca0515ed1aef1395f66b5303bb5d6f1bf9d61a353fa53f73f8ac9973fa9f6.
//
// Solidity: event Deposit(address indexed transaction, address indexed from, uint256 amount, uint256 input1, uint256 output1)
func (_MRC20 *MRC20Filterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *MRC20Deposit, token []common.Address, from []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _MRC20.contract.WatchLogs(opts, "Deposit", tokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MRC20Deposit)
				if err := _MRC20.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x4e2ca0515ed1aef1395f66b5303bb5d6f1bf9d61a353fa53f73f8ac9973fa9f6.
//
// Solidity: event Deposit(address indexed transaction, address indexed from, uint256 amount, uint256 input1, uint256 output1)
func (_MRC20 *MRC20Filterer) ParseDeposit(log types.Log) (*MRC20Deposit, error) {
	event := new(MRC20Deposit)
	if err := _MRC20.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MRC20LogFeeTransferIterator is returned from FilterLogFeeTransfer and is used to iterate over the raw logs and unpacked data for LogFeeTransfer events raised by the MRC20 contract.
type MRC20LogFeeTransferIterator struct {
	Event *MRC20LogFeeTransfer // Event containing the contract specifics and raw log

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
func (it *MRC20LogFeeTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MRC20LogFeeTransfer)
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
		it.Event = new(MRC20LogFeeTransfer)
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
func (it *MRC20LogFeeTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MRC20LogFeeTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MRC20LogFeeTransfer represents a LogFeeTransfer event raised by the MRC20 contract.
type MRC20LogFeeTransfer struct {
	Token   common.Address
	From    common.Address
	To      common.Address
	Amount  *big.Int
	Input1  *big.Int
	Input2  *big.Int
	Output1 *big.Int
	Output2 *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLogFeeTransfer is a free log retrieval operation binding the contract event 0x4dfe1bbbcf077ddc3e01291eea2d5c70c2b422b415d95645b9adcfd678cb1d63.
//
// Solidity: event LogFeeTransfer(address indexed transaction, address indexed from, address indexed to, uint256 amount, uint256 input1, uint256 input2, uint256 output1, uint256 output2)
func (_MRC20 *MRC20Filterer) FilterLogFeeTransfer(opts *bind.FilterOpts, token []common.Address, from []common.Address, to []common.Address) (*MRC20LogFeeTransferIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MRC20.contract.FilterLogs(opts, "LogFeeTransfer", tokenRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MRC20LogFeeTransferIterator{contract: _MRC20.contract, event: "LogFeeTransfer", logs: logs, sub: sub}, nil
}

// WatchLogFeeTransfer is a free log subscription operation binding the contract event 0x4dfe1bbbcf077ddc3e01291eea2d5c70c2b422b415d95645b9adcfd678cb1d63.
//
// Solidity: event LogFeeTransfer(address indexed transaction, address indexed from, address indexed to, uint256 amount, uint256 input1, uint256 input2, uint256 output1, uint256 output2)
func (_MRC20 *MRC20Filterer) WatchLogFeeTransfer(opts *bind.WatchOpts, sink chan<- *MRC20LogFeeTransfer, token []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MRC20.contract.WatchLogs(opts, "LogFeeTransfer", tokenRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MRC20LogFeeTransfer)
				if err := _MRC20.contract.UnpackLog(event, "LogFeeTransfer", log); err != nil {
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

// ParseLogFeeTransfer is a log parse operation binding the contract event 0x4dfe1bbbcf077ddc3e01291eea2d5c70c2b422b415d95645b9adcfd678cb1d63.
//
// Solidity: event LogFeeTransfer(address indexed transaction, address indexed from, address indexed to, uint256 amount, uint256 input1, uint256 input2, uint256 output1, uint256 output2)
func (_MRC20 *MRC20Filterer) ParseLogFeeTransfer(log types.Log) (*MRC20LogFeeTransfer, error) {
	event := new(MRC20LogFeeTransfer)
	if err := _MRC20.contract.UnpackLog(event, "LogFeeTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MRC20LogTransferIterator is returned from FilterLogTransfer and is used to iterate over the raw logs and unpacked data for LogTransfer events raised by the MRC20 contract.
type MRC20LogTransferIterator struct {
	Event *MRC20LogTransfer // Event containing the contract specifics and raw log

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
func (it *MRC20LogTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MRC20LogTransfer)
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
		it.Event = new(MRC20LogTransfer)
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
func (it *MRC20LogTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MRC20LogTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MRC20LogTransfer represents a LogTransfer event raised by the MRC20 contract.
type MRC20LogTransfer struct {
	Token   common.Address
	From    common.Address
	To      common.Address
	Amount  *big.Int
	Input1  *big.Int
	Input2  *big.Int
	Output1 *big.Int
	Output2 *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLogTransfer is a free log retrieval operation binding the contract event 0xe6497e3ee548a3372136af2fcb0696db31fc6cf20260707645068bd3fe97f3c4.
//
// Solidity: event LogTransfer(address indexed transaction, address indexed from, address indexed to, uint256 amount, uint256 input1, uint256 input2, uint256 output1, uint256 output2)
func (_MRC20 *MRC20Filterer) FilterLogTransfer(opts *bind.FilterOpts, token []common.Address, from []common.Address, to []common.Address) (*MRC20LogTransferIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MRC20.contract.FilterLogs(opts, "LogTransfer", tokenRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MRC20LogTransferIterator{contract: _MRC20.contract, event: "LogTransfer", logs: logs, sub: sub}, nil
}

// WatchLogTransfer is a free log subscription operation binding the contract event 0xe6497e3ee548a3372136af2fcb0696db31fc6cf20260707645068bd3fe97f3c4.
//
// Solidity: event LogTransfer(address indexed transaction, address indexed from, address indexed to, uint256 amount, uint256 input1, uint256 input2, uint256 output1, uint256 output2)
func (_MRC20 *MRC20Filterer) WatchLogTransfer(opts *bind.WatchOpts, sink chan<- *MRC20LogTransfer, token []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MRC20.contract.WatchLogs(opts, "LogTransfer", tokenRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MRC20LogTransfer)
				if err := _MRC20.contract.UnpackLog(event, "LogTransfer", log); err != nil {
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

// ParseLogTransfer is a log parse operation binding the contract event 0xe6497e3ee548a3372136af2fcb0696db31fc6cf20260707645068bd3fe97f3c4.
//
// Solidity: event LogTransfer(address indexed transaction, address indexed from, address indexed to, uint256 amount, uint256 input1, uint256 input2, uint256 output1, uint256 output2)
func (_MRC20 *MRC20Filterer) ParseLogTransfer(log types.Log) (*MRC20LogTransfer, error) {
	event := new(MRC20LogTransfer)
	if err := _MRC20.contract.UnpackLog(event, "LogTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MRC20OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MRC20 contract.
type MRC20OwnershipTransferredIterator struct {
	Event *MRC20OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MRC20OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MRC20OwnershipTransferred)
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
		it.Event = new(MRC20OwnershipTransferred)
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
func (it *MRC20OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MRC20OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MRC20OwnershipTransferred represents a OwnershipTransferred event raised by the MRC20 contract.
type MRC20OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MRC20 *MRC20Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MRC20OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MRC20.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MRC20OwnershipTransferredIterator{contract: _MRC20.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MRC20 *MRC20Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MRC20OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MRC20.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MRC20OwnershipTransferred)
				if err := _MRC20.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_MRC20 *MRC20Filterer) ParseOwnershipTransferred(log types.Log) (*MRC20OwnershipTransferred, error) {
	event := new(MRC20OwnershipTransferred)
	if err := _MRC20.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MRC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the MRC20 contract.
type MRC20TransferIterator struct {
	Event *MRC20Transfer // Event containing the contract specifics and raw log

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
func (it *MRC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MRC20Transfer)
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
		it.Event = new(MRC20Transfer)
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
func (it *MRC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MRC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MRC20Transfer represents a Transfer event raised by the MRC20 contract.
type MRC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MRC20 *MRC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MRC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MRC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MRC20TransferIterator{contract: _MRC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MRC20 *MRC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MRC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MRC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MRC20Transfer)
				if err := _MRC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MRC20 *MRC20Filterer) ParseTransfer(log types.Log) (*MRC20Transfer, error) {
	event := new(MRC20Transfer)
	if err := _MRC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MRC20WithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the MRC20 contract.
type MRC20WithdrawIterator struct {
	Event *MRC20Withdraw // Event containing the contract specifics and raw log

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
func (it *MRC20WithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MRC20Withdraw)
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
		it.Event = new(MRC20Withdraw)
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
func (it *MRC20WithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MRC20WithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MRC20Withdraw represents a Withdraw event raised by the MRC20 contract.
type MRC20Withdraw struct {
	Token   common.Address
	From    common.Address
	Amount  *big.Int
	Input1  *big.Int
	Output1 *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xebff2602b3f468259e1e99f613fed6691f3a6526effe6ef3e768ba7ae7a36c4f.
//
// Solidity: event Withdraw(address indexed transaction, address indexed from, uint256 amount, uint256 input1, uint256 output1)
func (_MRC20 *MRC20Filterer) FilterWithdraw(opts *bind.FilterOpts, token []common.Address, from []common.Address) (*MRC20WithdrawIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _MRC20.contract.FilterLogs(opts, "Withdraw", tokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &MRC20WithdrawIterator{contract: _MRC20.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xebff2602b3f468259e1e99f613fed6691f3a6526effe6ef3e768ba7ae7a36c4f.
//
// Solidity: event Withdraw(address indexed transaction, address indexed from, uint256 amount, uint256 input1, uint256 output1)
func (_MRC20 *MRC20Filterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *MRC20Withdraw, token []common.Address, from []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _MRC20.contract.WatchLogs(opts, "Withdraw", tokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MRC20Withdraw)
				if err := _MRC20.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xebff2602b3f468259e1e99f613fed6691f3a6526effe6ef3e768ba7ae7a36c4f.
//
// Solidity: event Withdraw(address indexed transaction, address indexed from, uint256 amount, uint256 input1, uint256 output1)
func (_MRC20 *MRC20Filterer) ParseWithdraw(log types.Log) (*MRC20Withdraw, error) {
	event := new(MRC20Withdraw)
	if err := _MRC20.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
