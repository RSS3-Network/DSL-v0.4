// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package registrarv1

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

// RegistrarV1MetaData contains all meta data concerning the RegistrarV1 contract.
var RegistrarV1MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractBaseRegistrar\",\"name\":\"_base\",\"type\":\"address\"},{\"internalType\":\"contractPriceOracle\",\"name\":\"_prices\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_minCommitmentAge\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxCommitmentAge\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"label\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expires\",\"type\":\"uint256\"}],\"name\":\"NameRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"label\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expires\",\"type\":\"uint256\"}],\"name\":\"NameRenewed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"}],\"name\":\"NewPriceOracle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"MIN_REGISTRATION_DURATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"available\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"name\":\"commit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"commitments\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"secret\",\"type\":\"bytes32\"}],\"name\":\"makeCommitment\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"secret\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"makeCommitmentWithConfig\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxCommitmentAge\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minCommitmentAge\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"secret\",\"type\":\"bytes32\"}],\"name\":\"register\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"secret\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"registerWithConfig\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"renew\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"rentPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minCommitmentAge\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxCommitmentAge\",\"type\":\"uint256\"}],\"name\":\"setCommitmentAges\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractPriceOracle\",\"name\":\"_prices\",\"type\":\"address\"}],\"name\":\"setPriceOracle\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceID\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"valid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// RegistrarV1ABI is the input ABI used to generate the binding from.
// Deprecated: Use RegistrarV1MetaData.ABI instead.
var RegistrarV1ABI = RegistrarV1MetaData.ABI

// RegistrarV1 is an auto generated Go binding around an Ethereum contract.
type RegistrarV1 struct {
	RegistrarV1Caller     // Read-only binding to the contract
	RegistrarV1Transactor // Write-only binding to the contract
	RegistrarV1Filterer   // Log filterer for contract events
}

// RegistrarV1Caller is an auto generated read-only Go binding around an Ethereum contract.
type RegistrarV1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrarV1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type RegistrarV1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrarV1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RegistrarV1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrarV1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RegistrarV1Session struct {
	Contract     *RegistrarV1      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RegistrarV1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RegistrarV1CallerSession struct {
	Contract *RegistrarV1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// RegistrarV1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RegistrarV1TransactorSession struct {
	Contract     *RegistrarV1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// RegistrarV1Raw is an auto generated low-level Go binding around an Ethereum contract.
type RegistrarV1Raw struct {
	Contract *RegistrarV1 // Generic contract binding to access the raw methods on
}

// RegistrarV1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RegistrarV1CallerRaw struct {
	Contract *RegistrarV1Caller // Generic read-only contract binding to access the raw methods on
}

// RegistrarV1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RegistrarV1TransactorRaw struct {
	Contract *RegistrarV1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewRegistrarV1 creates a new instance of RegistrarV1, bound to a specific deployed contract.
func NewRegistrarV1(address common.Address, backend bind.ContractBackend) (*RegistrarV1, error) {
	contract, err := bindRegistrarV1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RegistrarV1{RegistrarV1Caller: RegistrarV1Caller{contract: contract}, RegistrarV1Transactor: RegistrarV1Transactor{contract: contract}, RegistrarV1Filterer: RegistrarV1Filterer{contract: contract}}, nil
}

// NewRegistrarV1Caller creates a new read-only instance of RegistrarV1, bound to a specific deployed contract.
func NewRegistrarV1Caller(address common.Address, caller bind.ContractCaller) (*RegistrarV1Caller, error) {
	contract, err := bindRegistrarV1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RegistrarV1Caller{contract: contract}, nil
}

// NewRegistrarV1Transactor creates a new write-only instance of RegistrarV1, bound to a specific deployed contract.
func NewRegistrarV1Transactor(address common.Address, transactor bind.ContractTransactor) (*RegistrarV1Transactor, error) {
	contract, err := bindRegistrarV1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RegistrarV1Transactor{contract: contract}, nil
}

// NewRegistrarV1Filterer creates a new log filterer instance of RegistrarV1, bound to a specific deployed contract.
func NewRegistrarV1Filterer(address common.Address, filterer bind.ContractFilterer) (*RegistrarV1Filterer, error) {
	contract, err := bindRegistrarV1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RegistrarV1Filterer{contract: contract}, nil
}

// bindRegistrarV1 binds a generic wrapper to an already deployed contract.
func bindRegistrarV1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RegistrarV1MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RegistrarV1 *RegistrarV1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RegistrarV1.Contract.RegistrarV1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RegistrarV1 *RegistrarV1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RegistrarV1.Contract.RegistrarV1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RegistrarV1 *RegistrarV1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RegistrarV1.Contract.RegistrarV1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RegistrarV1 *RegistrarV1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RegistrarV1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RegistrarV1 *RegistrarV1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RegistrarV1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RegistrarV1 *RegistrarV1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RegistrarV1.Contract.contract.Transact(opts, method, params...)
}

// MINREGISTRATIONDURATION is a free data retrieval call binding the contract method 0x8a95b09f.
//
// Solidity: function MIN_REGISTRATION_DURATION() view returns(uint256)
func (_RegistrarV1 *RegistrarV1Caller) MINREGISTRATIONDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RegistrarV1.contract.Call(opts, &out, "MIN_REGISTRATION_DURATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINREGISTRATIONDURATION is a free data retrieval call binding the contract method 0x8a95b09f.
//
// Solidity: function MIN_REGISTRATION_DURATION() view returns(uint256)
func (_RegistrarV1 *RegistrarV1Session) MINREGISTRATIONDURATION() (*big.Int, error) {
	return _RegistrarV1.Contract.MINREGISTRATIONDURATION(&_RegistrarV1.CallOpts)
}

// MINREGISTRATIONDURATION is a free data retrieval call binding the contract method 0x8a95b09f.
//
// Solidity: function MIN_REGISTRATION_DURATION() view returns(uint256)
func (_RegistrarV1 *RegistrarV1CallerSession) MINREGISTRATIONDURATION() (*big.Int, error) {
	return _RegistrarV1.Contract.MINREGISTRATIONDURATION(&_RegistrarV1.CallOpts)
}

// Available is a free data retrieval call binding the contract method 0xaeb8ce9b.
//
// Solidity: function available(string name) view returns(bool)
func (_RegistrarV1 *RegistrarV1Caller) Available(opts *bind.CallOpts, name string) (bool, error) {
	var out []interface{}
	err := _RegistrarV1.contract.Call(opts, &out, "available", name)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Available is a free data retrieval call binding the contract method 0xaeb8ce9b.
//
// Solidity: function available(string name) view returns(bool)
func (_RegistrarV1 *RegistrarV1Session) Available(name string) (bool, error) {
	return _RegistrarV1.Contract.Available(&_RegistrarV1.CallOpts, name)
}

// Available is a free data retrieval call binding the contract method 0xaeb8ce9b.
//
// Solidity: function available(string name) view returns(bool)
func (_RegistrarV1 *RegistrarV1CallerSession) Available(name string) (bool, error) {
	return _RegistrarV1.Contract.Available(&_RegistrarV1.CallOpts, name)
}

// Commitments is a free data retrieval call binding the contract method 0x839df945.
//
// Solidity: function commitments(bytes32 ) view returns(uint256)
func (_RegistrarV1 *RegistrarV1Caller) Commitments(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _RegistrarV1.contract.Call(opts, &out, "commitments", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Commitments is a free data retrieval call binding the contract method 0x839df945.
//
// Solidity: function commitments(bytes32 ) view returns(uint256)
func (_RegistrarV1 *RegistrarV1Session) Commitments(arg0 [32]byte) (*big.Int, error) {
	return _RegistrarV1.Contract.Commitments(&_RegistrarV1.CallOpts, arg0)
}

// Commitments is a free data retrieval call binding the contract method 0x839df945.
//
// Solidity: function commitments(bytes32 ) view returns(uint256)
func (_RegistrarV1 *RegistrarV1CallerSession) Commitments(arg0 [32]byte) (*big.Int, error) {
	return _RegistrarV1.Contract.Commitments(&_RegistrarV1.CallOpts, arg0)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_RegistrarV1 *RegistrarV1Caller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _RegistrarV1.contract.Call(opts, &out, "isOwner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_RegistrarV1 *RegistrarV1Session) IsOwner() (bool, error) {
	return _RegistrarV1.Contract.IsOwner(&_RegistrarV1.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_RegistrarV1 *RegistrarV1CallerSession) IsOwner() (bool, error) {
	return _RegistrarV1.Contract.IsOwner(&_RegistrarV1.CallOpts)
}

// MakeCommitment is a free data retrieval call binding the contract method 0xf49826be.
//
// Solidity: function makeCommitment(string name, address owner, bytes32 secret) pure returns(bytes32)
func (_RegistrarV1 *RegistrarV1Caller) MakeCommitment(opts *bind.CallOpts, name string, owner common.Address, secret [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _RegistrarV1.contract.Call(opts, &out, "makeCommitment", name, owner, secret)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MakeCommitment is a free data retrieval call binding the contract method 0xf49826be.
//
// Solidity: function makeCommitment(string name, address owner, bytes32 secret) pure returns(bytes32)
func (_RegistrarV1 *RegistrarV1Session) MakeCommitment(name string, owner common.Address, secret [32]byte) ([32]byte, error) {
	return _RegistrarV1.Contract.MakeCommitment(&_RegistrarV1.CallOpts, name, owner, secret)
}

// MakeCommitment is a free data retrieval call binding the contract method 0xf49826be.
//
// Solidity: function makeCommitment(string name, address owner, bytes32 secret) pure returns(bytes32)
func (_RegistrarV1 *RegistrarV1CallerSession) MakeCommitment(name string, owner common.Address, secret [32]byte) ([32]byte, error) {
	return _RegistrarV1.Contract.MakeCommitment(&_RegistrarV1.CallOpts, name, owner, secret)
}

// MakeCommitmentWithConfig is a free data retrieval call binding the contract method 0x3d86c52f.
//
// Solidity: function makeCommitmentWithConfig(string name, address owner, bytes32 secret, address resolver, address addr) pure returns(bytes32)
func (_RegistrarV1 *RegistrarV1Caller) MakeCommitmentWithConfig(opts *bind.CallOpts, name string, owner common.Address, secret [32]byte, resolver common.Address, addr common.Address) ([32]byte, error) {
	var out []interface{}
	err := _RegistrarV1.contract.Call(opts, &out, "makeCommitmentWithConfig", name, owner, secret, resolver, addr)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MakeCommitmentWithConfig is a free data retrieval call binding the contract method 0x3d86c52f.
//
// Solidity: function makeCommitmentWithConfig(string name, address owner, bytes32 secret, address resolver, address addr) pure returns(bytes32)
func (_RegistrarV1 *RegistrarV1Session) MakeCommitmentWithConfig(name string, owner common.Address, secret [32]byte, resolver common.Address, addr common.Address) ([32]byte, error) {
	return _RegistrarV1.Contract.MakeCommitmentWithConfig(&_RegistrarV1.CallOpts, name, owner, secret, resolver, addr)
}

// MakeCommitmentWithConfig is a free data retrieval call binding the contract method 0x3d86c52f.
//
// Solidity: function makeCommitmentWithConfig(string name, address owner, bytes32 secret, address resolver, address addr) pure returns(bytes32)
func (_RegistrarV1 *RegistrarV1CallerSession) MakeCommitmentWithConfig(name string, owner common.Address, secret [32]byte, resolver common.Address, addr common.Address) ([32]byte, error) {
	return _RegistrarV1.Contract.MakeCommitmentWithConfig(&_RegistrarV1.CallOpts, name, owner, secret, resolver, addr)
}

// MaxCommitmentAge is a free data retrieval call binding the contract method 0xce1e09c0.
//
// Solidity: function maxCommitmentAge() view returns(uint256)
func (_RegistrarV1 *RegistrarV1Caller) MaxCommitmentAge(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RegistrarV1.contract.Call(opts, &out, "maxCommitmentAge")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxCommitmentAge is a free data retrieval call binding the contract method 0xce1e09c0.
//
// Solidity: function maxCommitmentAge() view returns(uint256)
func (_RegistrarV1 *RegistrarV1Session) MaxCommitmentAge() (*big.Int, error) {
	return _RegistrarV1.Contract.MaxCommitmentAge(&_RegistrarV1.CallOpts)
}

// MaxCommitmentAge is a free data retrieval call binding the contract method 0xce1e09c0.
//
// Solidity: function maxCommitmentAge() view returns(uint256)
func (_RegistrarV1 *RegistrarV1CallerSession) MaxCommitmentAge() (*big.Int, error) {
	return _RegistrarV1.Contract.MaxCommitmentAge(&_RegistrarV1.CallOpts)
}

// MinCommitmentAge is a free data retrieval call binding the contract method 0x8d839ffe.
//
// Solidity: function minCommitmentAge() view returns(uint256)
func (_RegistrarV1 *RegistrarV1Caller) MinCommitmentAge(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RegistrarV1.contract.Call(opts, &out, "minCommitmentAge")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinCommitmentAge is a free data retrieval call binding the contract method 0x8d839ffe.
//
// Solidity: function minCommitmentAge() view returns(uint256)
func (_RegistrarV1 *RegistrarV1Session) MinCommitmentAge() (*big.Int, error) {
	return _RegistrarV1.Contract.MinCommitmentAge(&_RegistrarV1.CallOpts)
}

// MinCommitmentAge is a free data retrieval call binding the contract method 0x8d839ffe.
//
// Solidity: function minCommitmentAge() view returns(uint256)
func (_RegistrarV1 *RegistrarV1CallerSession) MinCommitmentAge() (*big.Int, error) {
	return _RegistrarV1.Contract.MinCommitmentAge(&_RegistrarV1.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RegistrarV1 *RegistrarV1Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RegistrarV1.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RegistrarV1 *RegistrarV1Session) Owner() (common.Address, error) {
	return _RegistrarV1.Contract.Owner(&_RegistrarV1.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RegistrarV1 *RegistrarV1CallerSession) Owner() (common.Address, error) {
	return _RegistrarV1.Contract.Owner(&_RegistrarV1.CallOpts)
}

// RentPrice is a free data retrieval call binding the contract method 0x83e7f6ff.
//
// Solidity: function rentPrice(string name, uint256 duration) view returns(uint256)
func (_RegistrarV1 *RegistrarV1Caller) RentPrice(opts *bind.CallOpts, name string, duration *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RegistrarV1.contract.Call(opts, &out, "rentPrice", name, duration)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RentPrice is a free data retrieval call binding the contract method 0x83e7f6ff.
//
// Solidity: function rentPrice(string name, uint256 duration) view returns(uint256)
func (_RegistrarV1 *RegistrarV1Session) RentPrice(name string, duration *big.Int) (*big.Int, error) {
	return _RegistrarV1.Contract.RentPrice(&_RegistrarV1.CallOpts, name, duration)
}

// RentPrice is a free data retrieval call binding the contract method 0x83e7f6ff.
//
// Solidity: function rentPrice(string name, uint256 duration) view returns(uint256)
func (_RegistrarV1 *RegistrarV1CallerSession) RentPrice(name string, duration *big.Int) (*big.Int, error) {
	return _RegistrarV1.Contract.RentPrice(&_RegistrarV1.CallOpts, name, duration)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) pure returns(bool)
func (_RegistrarV1 *RegistrarV1Caller) SupportsInterface(opts *bind.CallOpts, interfaceID [4]byte) (bool, error) {
	var out []interface{}
	err := _RegistrarV1.contract.Call(opts, &out, "supportsInterface", interfaceID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) pure returns(bool)
func (_RegistrarV1 *RegistrarV1Session) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _RegistrarV1.Contract.SupportsInterface(&_RegistrarV1.CallOpts, interfaceID)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) pure returns(bool)
func (_RegistrarV1 *RegistrarV1CallerSession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _RegistrarV1.Contract.SupportsInterface(&_RegistrarV1.CallOpts, interfaceID)
}

// Valid is a free data retrieval call binding the contract method 0x9791c097.
//
// Solidity: function valid(string name) pure returns(bool)
func (_RegistrarV1 *RegistrarV1Caller) Valid(opts *bind.CallOpts, name string) (bool, error) {
	var out []interface{}
	err := _RegistrarV1.contract.Call(opts, &out, "valid", name)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Valid is a free data retrieval call binding the contract method 0x9791c097.
//
// Solidity: function valid(string name) pure returns(bool)
func (_RegistrarV1 *RegistrarV1Session) Valid(name string) (bool, error) {
	return _RegistrarV1.Contract.Valid(&_RegistrarV1.CallOpts, name)
}

// Valid is a free data retrieval call binding the contract method 0x9791c097.
//
// Solidity: function valid(string name) pure returns(bool)
func (_RegistrarV1 *RegistrarV1CallerSession) Valid(name string) (bool, error) {
	return _RegistrarV1.Contract.Valid(&_RegistrarV1.CallOpts, name)
}

// Commit is a paid mutator transaction binding the contract method 0xf14fcbc8.
//
// Solidity: function commit(bytes32 commitment) returns()
func (_RegistrarV1 *RegistrarV1Transactor) Commit(opts *bind.TransactOpts, commitment [32]byte) (*types.Transaction, error) {
	return _RegistrarV1.contract.Transact(opts, "commit", commitment)
}

// Commit is a paid mutator transaction binding the contract method 0xf14fcbc8.
//
// Solidity: function commit(bytes32 commitment) returns()
func (_RegistrarV1 *RegistrarV1Session) Commit(commitment [32]byte) (*types.Transaction, error) {
	return _RegistrarV1.Contract.Commit(&_RegistrarV1.TransactOpts, commitment)
}

// Commit is a paid mutator transaction binding the contract method 0xf14fcbc8.
//
// Solidity: function commit(bytes32 commitment) returns()
func (_RegistrarV1 *RegistrarV1TransactorSession) Commit(commitment [32]byte) (*types.Transaction, error) {
	return _RegistrarV1.Contract.Commit(&_RegistrarV1.TransactOpts, commitment)
}

// Register is a paid mutator transaction binding the contract method 0x85f6d155.
//
// Solidity: function register(string name, address owner, uint256 duration, bytes32 secret) payable returns()
func (_RegistrarV1 *RegistrarV1Transactor) Register(opts *bind.TransactOpts, name string, owner common.Address, duration *big.Int, secret [32]byte) (*types.Transaction, error) {
	return _RegistrarV1.contract.Transact(opts, "register", name, owner, duration, secret)
}

// Register is a paid mutator transaction binding the contract method 0x85f6d155.
//
// Solidity: function register(string name, address owner, uint256 duration, bytes32 secret) payable returns()
func (_RegistrarV1 *RegistrarV1Session) Register(name string, owner common.Address, duration *big.Int, secret [32]byte) (*types.Transaction, error) {
	return _RegistrarV1.Contract.Register(&_RegistrarV1.TransactOpts, name, owner, duration, secret)
}

// Register is a paid mutator transaction binding the contract method 0x85f6d155.
//
// Solidity: function register(string name, address owner, uint256 duration, bytes32 secret) payable returns()
func (_RegistrarV1 *RegistrarV1TransactorSession) Register(name string, owner common.Address, duration *big.Int, secret [32]byte) (*types.Transaction, error) {
	return _RegistrarV1.Contract.Register(&_RegistrarV1.TransactOpts, name, owner, duration, secret)
}

// RegisterWithConfig is a paid mutator transaction binding the contract method 0xf7a16963.
//
// Solidity: function registerWithConfig(string name, address owner, uint256 duration, bytes32 secret, address resolver, address addr) payable returns()
func (_RegistrarV1 *RegistrarV1Transactor) RegisterWithConfig(opts *bind.TransactOpts, name string, owner common.Address, duration *big.Int, secret [32]byte, resolver common.Address, addr common.Address) (*types.Transaction, error) {
	return _RegistrarV1.contract.Transact(opts, "registerWithConfig", name, owner, duration, secret, resolver, addr)
}

// RegisterWithConfig is a paid mutator transaction binding the contract method 0xf7a16963.
//
// Solidity: function registerWithConfig(string name, address owner, uint256 duration, bytes32 secret, address resolver, address addr) payable returns()
func (_RegistrarV1 *RegistrarV1Session) RegisterWithConfig(name string, owner common.Address, duration *big.Int, secret [32]byte, resolver common.Address, addr common.Address) (*types.Transaction, error) {
	return _RegistrarV1.Contract.RegisterWithConfig(&_RegistrarV1.TransactOpts, name, owner, duration, secret, resolver, addr)
}

// RegisterWithConfig is a paid mutator transaction binding the contract method 0xf7a16963.
//
// Solidity: function registerWithConfig(string name, address owner, uint256 duration, bytes32 secret, address resolver, address addr) payable returns()
func (_RegistrarV1 *RegistrarV1TransactorSession) RegisterWithConfig(name string, owner common.Address, duration *big.Int, secret [32]byte, resolver common.Address, addr common.Address) (*types.Transaction, error) {
	return _RegistrarV1.Contract.RegisterWithConfig(&_RegistrarV1.TransactOpts, name, owner, duration, secret, resolver, addr)
}

// Renew is a paid mutator transaction binding the contract method 0xacf1a841.
//
// Solidity: function renew(string name, uint256 duration) payable returns()
func (_RegistrarV1 *RegistrarV1Transactor) Renew(opts *bind.TransactOpts, name string, duration *big.Int) (*types.Transaction, error) {
	return _RegistrarV1.contract.Transact(opts, "renew", name, duration)
}

// Renew is a paid mutator transaction binding the contract method 0xacf1a841.
//
// Solidity: function renew(string name, uint256 duration) payable returns()
func (_RegistrarV1 *RegistrarV1Session) Renew(name string, duration *big.Int) (*types.Transaction, error) {
	return _RegistrarV1.Contract.Renew(&_RegistrarV1.TransactOpts, name, duration)
}

// Renew is a paid mutator transaction binding the contract method 0xacf1a841.
//
// Solidity: function renew(string name, uint256 duration) payable returns()
func (_RegistrarV1 *RegistrarV1TransactorSession) Renew(name string, duration *big.Int) (*types.Transaction, error) {
	return _RegistrarV1.Contract.Renew(&_RegistrarV1.TransactOpts, name, duration)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RegistrarV1 *RegistrarV1Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RegistrarV1.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RegistrarV1 *RegistrarV1Session) RenounceOwnership() (*types.Transaction, error) {
	return _RegistrarV1.Contract.RenounceOwnership(&_RegistrarV1.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RegistrarV1 *RegistrarV1TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RegistrarV1.Contract.RenounceOwnership(&_RegistrarV1.TransactOpts)
}

// SetCommitmentAges is a paid mutator transaction binding the contract method 0x7e324479.
//
// Solidity: function setCommitmentAges(uint256 _minCommitmentAge, uint256 _maxCommitmentAge) returns()
func (_RegistrarV1 *RegistrarV1Transactor) SetCommitmentAges(opts *bind.TransactOpts, _minCommitmentAge *big.Int, _maxCommitmentAge *big.Int) (*types.Transaction, error) {
	return _RegistrarV1.contract.Transact(opts, "setCommitmentAges", _minCommitmentAge, _maxCommitmentAge)
}

// SetCommitmentAges is a paid mutator transaction binding the contract method 0x7e324479.
//
// Solidity: function setCommitmentAges(uint256 _minCommitmentAge, uint256 _maxCommitmentAge) returns()
func (_RegistrarV1 *RegistrarV1Session) SetCommitmentAges(_minCommitmentAge *big.Int, _maxCommitmentAge *big.Int) (*types.Transaction, error) {
	return _RegistrarV1.Contract.SetCommitmentAges(&_RegistrarV1.TransactOpts, _minCommitmentAge, _maxCommitmentAge)
}

// SetCommitmentAges is a paid mutator transaction binding the contract method 0x7e324479.
//
// Solidity: function setCommitmentAges(uint256 _minCommitmentAge, uint256 _maxCommitmentAge) returns()
func (_RegistrarV1 *RegistrarV1TransactorSession) SetCommitmentAges(_minCommitmentAge *big.Int, _maxCommitmentAge *big.Int) (*types.Transaction, error) {
	return _RegistrarV1.Contract.SetCommitmentAges(&_RegistrarV1.TransactOpts, _minCommitmentAge, _maxCommitmentAge)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address _prices) returns()
func (_RegistrarV1 *RegistrarV1Transactor) SetPriceOracle(opts *bind.TransactOpts, _prices common.Address) (*types.Transaction, error) {
	return _RegistrarV1.contract.Transact(opts, "setPriceOracle", _prices)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address _prices) returns()
func (_RegistrarV1 *RegistrarV1Session) SetPriceOracle(_prices common.Address) (*types.Transaction, error) {
	return _RegistrarV1.Contract.SetPriceOracle(&_RegistrarV1.TransactOpts, _prices)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address _prices) returns()
func (_RegistrarV1 *RegistrarV1TransactorSession) SetPriceOracle(_prices common.Address) (*types.Transaction, error) {
	return _RegistrarV1.Contract.SetPriceOracle(&_RegistrarV1.TransactOpts, _prices)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RegistrarV1 *RegistrarV1Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _RegistrarV1.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RegistrarV1 *RegistrarV1Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RegistrarV1.Contract.TransferOwnership(&_RegistrarV1.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RegistrarV1 *RegistrarV1TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RegistrarV1.Contract.TransferOwnership(&_RegistrarV1.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_RegistrarV1 *RegistrarV1Transactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RegistrarV1.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_RegistrarV1 *RegistrarV1Session) Withdraw() (*types.Transaction, error) {
	return _RegistrarV1.Contract.Withdraw(&_RegistrarV1.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_RegistrarV1 *RegistrarV1TransactorSession) Withdraw() (*types.Transaction, error) {
	return _RegistrarV1.Contract.Withdraw(&_RegistrarV1.TransactOpts)
}

// RegistrarV1NameRegisteredIterator is returned from FilterNameRegistered and is used to iterate over the raw logs and unpacked data for NameRegistered events raised by the RegistrarV1 contract.
type RegistrarV1NameRegisteredIterator struct {
	Event *RegistrarV1NameRegistered // Event containing the contract specifics and raw log

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
func (it *RegistrarV1NameRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrarV1NameRegistered)
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
		it.Event = new(RegistrarV1NameRegistered)
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
func (it *RegistrarV1NameRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrarV1NameRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrarV1NameRegistered represents a NameRegistered event raised by the RegistrarV1 contract.
type RegistrarV1NameRegistered struct {
	Name    string
	Label   [32]byte
	Owner   common.Address
	Cost    *big.Int
	Expires *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNameRegistered is a free log retrieval operation binding the contract event 0xca6abbe9d7f11422cb6ca7629fbf6fe9efb1c621f71ce8f02b9f2a230097404f.
//
// Solidity: event NameRegistered(string name, bytes32 indexed label, address indexed owner, uint256 cost, uint256 expires)
func (_RegistrarV1 *RegistrarV1Filterer) FilterNameRegistered(opts *bind.FilterOpts, label [][32]byte, owner []common.Address) (*RegistrarV1NameRegisteredIterator, error) {

	var labelRule []interface{}
	for _, labelItem := range label {
		labelRule = append(labelRule, labelItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _RegistrarV1.contract.FilterLogs(opts, "NameRegistered", labelRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &RegistrarV1NameRegisteredIterator{contract: _RegistrarV1.contract, event: "NameRegistered", logs: logs, sub: sub}, nil
}

// WatchNameRegistered is a free log subscription operation binding the contract event 0xca6abbe9d7f11422cb6ca7629fbf6fe9efb1c621f71ce8f02b9f2a230097404f.
//
// Solidity: event NameRegistered(string name, bytes32 indexed label, address indexed owner, uint256 cost, uint256 expires)
func (_RegistrarV1 *RegistrarV1Filterer) WatchNameRegistered(opts *bind.WatchOpts, sink chan<- *RegistrarV1NameRegistered, label [][32]byte, owner []common.Address) (event.Subscription, error) {

	var labelRule []interface{}
	for _, labelItem := range label {
		labelRule = append(labelRule, labelItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _RegistrarV1.contract.WatchLogs(opts, "NameRegistered", labelRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrarV1NameRegistered)
				if err := _RegistrarV1.contract.UnpackLog(event, "NameRegistered", log); err != nil {
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

// ParseNameRegistered is a log parse operation binding the contract event 0xca6abbe9d7f11422cb6ca7629fbf6fe9efb1c621f71ce8f02b9f2a230097404f.
//
// Solidity: event NameRegistered(string name, bytes32 indexed label, address indexed owner, uint256 cost, uint256 expires)
func (_RegistrarV1 *RegistrarV1Filterer) ParseNameRegistered(log types.Log) (*RegistrarV1NameRegistered, error) {
	event := new(RegistrarV1NameRegistered)
	if err := _RegistrarV1.contract.UnpackLog(event, "NameRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrarV1NameRenewedIterator is returned from FilterNameRenewed and is used to iterate over the raw logs and unpacked data for NameRenewed events raised by the RegistrarV1 contract.
type RegistrarV1NameRenewedIterator struct {
	Event *RegistrarV1NameRenewed // Event containing the contract specifics and raw log

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
func (it *RegistrarV1NameRenewedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrarV1NameRenewed)
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
		it.Event = new(RegistrarV1NameRenewed)
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
func (it *RegistrarV1NameRenewedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrarV1NameRenewedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrarV1NameRenewed represents a NameRenewed event raised by the RegistrarV1 contract.
type RegistrarV1NameRenewed struct {
	Name    string
	Label   [32]byte
	Cost    *big.Int
	Expires *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNameRenewed is a free log retrieval operation binding the contract event 0x3da24c024582931cfaf8267d8ed24d13a82a8068d5bd337d30ec45cea4e506ae.
//
// Solidity: event NameRenewed(string name, bytes32 indexed label, uint256 cost, uint256 expires)
func (_RegistrarV1 *RegistrarV1Filterer) FilterNameRenewed(opts *bind.FilterOpts, label [][32]byte) (*RegistrarV1NameRenewedIterator, error) {

	var labelRule []interface{}
	for _, labelItem := range label {
		labelRule = append(labelRule, labelItem)
	}

	logs, sub, err := _RegistrarV1.contract.FilterLogs(opts, "NameRenewed", labelRule)
	if err != nil {
		return nil, err
	}
	return &RegistrarV1NameRenewedIterator{contract: _RegistrarV1.contract, event: "NameRenewed", logs: logs, sub: sub}, nil
}

// WatchNameRenewed is a free log subscription operation binding the contract event 0x3da24c024582931cfaf8267d8ed24d13a82a8068d5bd337d30ec45cea4e506ae.
//
// Solidity: event NameRenewed(string name, bytes32 indexed label, uint256 cost, uint256 expires)
func (_RegistrarV1 *RegistrarV1Filterer) WatchNameRenewed(opts *bind.WatchOpts, sink chan<- *RegistrarV1NameRenewed, label [][32]byte) (event.Subscription, error) {

	var labelRule []interface{}
	for _, labelItem := range label {
		labelRule = append(labelRule, labelItem)
	}

	logs, sub, err := _RegistrarV1.contract.WatchLogs(opts, "NameRenewed", labelRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrarV1NameRenewed)
				if err := _RegistrarV1.contract.UnpackLog(event, "NameRenewed", log); err != nil {
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

// ParseNameRenewed is a log parse operation binding the contract event 0x3da24c024582931cfaf8267d8ed24d13a82a8068d5bd337d30ec45cea4e506ae.
//
// Solidity: event NameRenewed(string name, bytes32 indexed label, uint256 cost, uint256 expires)
func (_RegistrarV1 *RegistrarV1Filterer) ParseNameRenewed(log types.Log) (*RegistrarV1NameRenewed, error) {
	event := new(RegistrarV1NameRenewed)
	if err := _RegistrarV1.contract.UnpackLog(event, "NameRenewed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrarV1NewPriceOracleIterator is returned from FilterNewPriceOracle and is used to iterate over the raw logs and unpacked data for NewPriceOracle events raised by the RegistrarV1 contract.
type RegistrarV1NewPriceOracleIterator struct {
	Event *RegistrarV1NewPriceOracle // Event containing the contract specifics and raw log

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
func (it *RegistrarV1NewPriceOracleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrarV1NewPriceOracle)
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
		it.Event = new(RegistrarV1NewPriceOracle)
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
func (it *RegistrarV1NewPriceOracleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrarV1NewPriceOracleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrarV1NewPriceOracle represents a NewPriceOracle event raised by the RegistrarV1 contract.
type RegistrarV1NewPriceOracle struct {
	Oracle common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNewPriceOracle is a free log retrieval operation binding the contract event 0xf261845a790fe29bbd6631e2ca4a5bdc83e6eed7c3271d9590d97287e00e9123.
//
// Solidity: event NewPriceOracle(address indexed oracle)
func (_RegistrarV1 *RegistrarV1Filterer) FilterNewPriceOracle(opts *bind.FilterOpts, oracle []common.Address) (*RegistrarV1NewPriceOracleIterator, error) {

	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _RegistrarV1.contract.FilterLogs(opts, "NewPriceOracle", oracleRule)
	if err != nil {
		return nil, err
	}
	return &RegistrarV1NewPriceOracleIterator{contract: _RegistrarV1.contract, event: "NewPriceOracle", logs: logs, sub: sub}, nil
}

// WatchNewPriceOracle is a free log subscription operation binding the contract event 0xf261845a790fe29bbd6631e2ca4a5bdc83e6eed7c3271d9590d97287e00e9123.
//
// Solidity: event NewPriceOracle(address indexed oracle)
func (_RegistrarV1 *RegistrarV1Filterer) WatchNewPriceOracle(opts *bind.WatchOpts, sink chan<- *RegistrarV1NewPriceOracle, oracle []common.Address) (event.Subscription, error) {

	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _RegistrarV1.contract.WatchLogs(opts, "NewPriceOracle", oracleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrarV1NewPriceOracle)
				if err := _RegistrarV1.contract.UnpackLog(event, "NewPriceOracle", log); err != nil {
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

// ParseNewPriceOracle is a log parse operation binding the contract event 0xf261845a790fe29bbd6631e2ca4a5bdc83e6eed7c3271d9590d97287e00e9123.
//
// Solidity: event NewPriceOracle(address indexed oracle)
func (_RegistrarV1 *RegistrarV1Filterer) ParseNewPriceOracle(log types.Log) (*RegistrarV1NewPriceOracle, error) {
	event := new(RegistrarV1NewPriceOracle)
	if err := _RegistrarV1.contract.UnpackLog(event, "NewPriceOracle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrarV1OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RegistrarV1 contract.
type RegistrarV1OwnershipTransferredIterator struct {
	Event *RegistrarV1OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RegistrarV1OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrarV1OwnershipTransferred)
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
		it.Event = new(RegistrarV1OwnershipTransferred)
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
func (it *RegistrarV1OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrarV1OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrarV1OwnershipTransferred represents a OwnershipTransferred event raised by the RegistrarV1 contract.
type RegistrarV1OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RegistrarV1 *RegistrarV1Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RegistrarV1OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RegistrarV1.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RegistrarV1OwnershipTransferredIterator{contract: _RegistrarV1.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RegistrarV1 *RegistrarV1Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RegistrarV1OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RegistrarV1.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrarV1OwnershipTransferred)
				if err := _RegistrarV1.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_RegistrarV1 *RegistrarV1Filterer) ParseOwnershipTransferred(log types.Log) (*RegistrarV1OwnershipTransferred, error) {
	event := new(RegistrarV1OwnershipTransferred)
	if err := _RegistrarV1.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
