// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package registrar

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

// RegistrarMetaData contains all meta data concerning the Registrar contract.
var RegistrarMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractBaseRegistrar\",\"name\":\"_base\",\"type\":\"address\"},{\"internalType\":\"contractPriceOracle\",\"name\":\"_prices\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_minCommitmentAge\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxCommitmentAge\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"label\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expires\",\"type\":\"uint256\"}],\"name\":\"NameRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"label\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expires\",\"type\":\"uint256\"}],\"name\":\"NameRenewed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"}],\"name\":\"NewPriceOracle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"MIN_REGISTRATION_DURATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"available\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"name\":\"commit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"commitments\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"secret\",\"type\":\"bytes32\"}],\"name\":\"makeCommitment\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"secret\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"makeCommitmentWithConfig\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxCommitmentAge\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minCommitmentAge\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"secret\",\"type\":\"bytes32\"}],\"name\":\"register\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"secret\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"registerWithConfig\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"renew\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"rentPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minCommitmentAge\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxCommitmentAge\",\"type\":\"uint256\"}],\"name\":\"setCommitmentAges\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractPriceOracle\",\"name\":\"_prices\",\"type\":\"address\"}],\"name\":\"setPriceOracle\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceID\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"valid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// RegistrarABI is the input ABI used to generate the binding from.
// Deprecated: Use RegistrarMetaData.ABI instead.
var RegistrarABI = RegistrarMetaData.ABI

// Registrar is an auto generated Go binding around an Ethereum contract.
type Registrar struct {
	RegistrarCaller     // Read-only binding to the contract
	RegistrarTransactor // Write-only binding to the contract
	RegistrarFilterer   // Log filterer for contract events
}

// RegistrarCaller is an auto generated read-only Go binding around an Ethereum contract.
type RegistrarCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrarTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RegistrarTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrarFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RegistrarFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrarSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RegistrarSession struct {
	Contract     *Registrar        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RegistrarCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RegistrarCallerSession struct {
	Contract *RegistrarCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// RegistrarTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RegistrarTransactorSession struct {
	Contract     *RegistrarTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// RegistrarRaw is an auto generated low-level Go binding around an Ethereum contract.
type RegistrarRaw struct {
	Contract *Registrar // Generic contract binding to access the raw methods on
}

// RegistrarCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RegistrarCallerRaw struct {
	Contract *RegistrarCaller // Generic read-only contract binding to access the raw methods on
}

// RegistrarTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RegistrarTransactorRaw struct {
	Contract *RegistrarTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRegistrar creates a new instance of Registrar, bound to a specific deployed contract.
func NewRegistrar(address common.Address, backend bind.ContractBackend) (*Registrar, error) {
	contract, err := bindRegistrar(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Registrar{RegistrarCaller: RegistrarCaller{contract: contract}, RegistrarTransactor: RegistrarTransactor{contract: contract}, RegistrarFilterer: RegistrarFilterer{contract: contract}}, nil
}

// NewRegistrarCaller creates a new read-only instance of Registrar, bound to a specific deployed contract.
func NewRegistrarCaller(address common.Address, caller bind.ContractCaller) (*RegistrarCaller, error) {
	contract, err := bindRegistrar(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RegistrarCaller{contract: contract}, nil
}

// NewRegistrarTransactor creates a new write-only instance of Registrar, bound to a specific deployed contract.
func NewRegistrarTransactor(address common.Address, transactor bind.ContractTransactor) (*RegistrarTransactor, error) {
	contract, err := bindRegistrar(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RegistrarTransactor{contract: contract}, nil
}

// NewRegistrarFilterer creates a new log filterer instance of Registrar, bound to a specific deployed contract.
func NewRegistrarFilterer(address common.Address, filterer bind.ContractFilterer) (*RegistrarFilterer, error) {
	contract, err := bindRegistrar(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RegistrarFilterer{contract: contract}, nil
}

// bindRegistrar binds a generic wrapper to an already deployed contract.
func bindRegistrar(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RegistrarMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registrar *RegistrarRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Registrar.Contract.RegistrarCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registrar *RegistrarRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registrar.Contract.RegistrarTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registrar *RegistrarRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registrar.Contract.RegistrarTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registrar *RegistrarCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Registrar.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registrar *RegistrarTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registrar.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registrar *RegistrarTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registrar.Contract.contract.Transact(opts, method, params...)
}

// MINREGISTRATIONDURATION is a free data retrieval call binding the contract method 0x8a95b09f.
//
// Solidity: function MIN_REGISTRATION_DURATION() view returns(uint256)
func (_Registrar *RegistrarCaller) MINREGISTRATIONDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Registrar.contract.Call(opts, &out, "MIN_REGISTRATION_DURATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINREGISTRATIONDURATION is a free data retrieval call binding the contract method 0x8a95b09f.
//
// Solidity: function MIN_REGISTRATION_DURATION() view returns(uint256)
func (_Registrar *RegistrarSession) MINREGISTRATIONDURATION() (*big.Int, error) {
	return _Registrar.Contract.MINREGISTRATIONDURATION(&_Registrar.CallOpts)
}

// MINREGISTRATIONDURATION is a free data retrieval call binding the contract method 0x8a95b09f.
//
// Solidity: function MIN_REGISTRATION_DURATION() view returns(uint256)
func (_Registrar *RegistrarCallerSession) MINREGISTRATIONDURATION() (*big.Int, error) {
	return _Registrar.Contract.MINREGISTRATIONDURATION(&_Registrar.CallOpts)
}

// Available is a free data retrieval call binding the contract method 0xaeb8ce9b.
//
// Solidity: function available(string name) view returns(bool)
func (_Registrar *RegistrarCaller) Available(opts *bind.CallOpts, name string) (bool, error) {
	var out []interface{}
	err := _Registrar.contract.Call(opts, &out, "available", name)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Available is a free data retrieval call binding the contract method 0xaeb8ce9b.
//
// Solidity: function available(string name) view returns(bool)
func (_Registrar *RegistrarSession) Available(name string) (bool, error) {
	return _Registrar.Contract.Available(&_Registrar.CallOpts, name)
}

// Available is a free data retrieval call binding the contract method 0xaeb8ce9b.
//
// Solidity: function available(string name) view returns(bool)
func (_Registrar *RegistrarCallerSession) Available(name string) (bool, error) {
	return _Registrar.Contract.Available(&_Registrar.CallOpts, name)
}

// Commitments is a free data retrieval call binding the contract method 0x839df945.
//
// Solidity: function commitments(bytes32 ) view returns(uint256)
func (_Registrar *RegistrarCaller) Commitments(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Registrar.contract.Call(opts, &out, "commitments", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Commitments is a free data retrieval call binding the contract method 0x839df945.
//
// Solidity: function commitments(bytes32 ) view returns(uint256)
func (_Registrar *RegistrarSession) Commitments(arg0 [32]byte) (*big.Int, error) {
	return _Registrar.Contract.Commitments(&_Registrar.CallOpts, arg0)
}

// Commitments is a free data retrieval call binding the contract method 0x839df945.
//
// Solidity: function commitments(bytes32 ) view returns(uint256)
func (_Registrar *RegistrarCallerSession) Commitments(arg0 [32]byte) (*big.Int, error) {
	return _Registrar.Contract.Commitments(&_Registrar.CallOpts, arg0)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Registrar *RegistrarCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Registrar.contract.Call(opts, &out, "isOwner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Registrar *RegistrarSession) IsOwner() (bool, error) {
	return _Registrar.Contract.IsOwner(&_Registrar.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Registrar *RegistrarCallerSession) IsOwner() (bool, error) {
	return _Registrar.Contract.IsOwner(&_Registrar.CallOpts)
}

// MakeCommitment is a free data retrieval call binding the contract method 0xf49826be.
//
// Solidity: function makeCommitment(string name, address owner, bytes32 secret) pure returns(bytes32)
func (_Registrar *RegistrarCaller) MakeCommitment(opts *bind.CallOpts, name string, owner common.Address, secret [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Registrar.contract.Call(opts, &out, "makeCommitment", name, owner, secret)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MakeCommitment is a free data retrieval call binding the contract method 0xf49826be.
//
// Solidity: function makeCommitment(string name, address owner, bytes32 secret) pure returns(bytes32)
func (_Registrar *RegistrarSession) MakeCommitment(name string, owner common.Address, secret [32]byte) ([32]byte, error) {
	return _Registrar.Contract.MakeCommitment(&_Registrar.CallOpts, name, owner, secret)
}

// MakeCommitment is a free data retrieval call binding the contract method 0xf49826be.
//
// Solidity: function makeCommitment(string name, address owner, bytes32 secret) pure returns(bytes32)
func (_Registrar *RegistrarCallerSession) MakeCommitment(name string, owner common.Address, secret [32]byte) ([32]byte, error) {
	return _Registrar.Contract.MakeCommitment(&_Registrar.CallOpts, name, owner, secret)
}

// MakeCommitmentWithConfig is a free data retrieval call binding the contract method 0x3d86c52f.
//
// Solidity: function makeCommitmentWithConfig(string name, address owner, bytes32 secret, address resolver, address addr) pure returns(bytes32)
func (_Registrar *RegistrarCaller) MakeCommitmentWithConfig(opts *bind.CallOpts, name string, owner common.Address, secret [32]byte, resolver common.Address, addr common.Address) ([32]byte, error) {
	var out []interface{}
	err := _Registrar.contract.Call(opts, &out, "makeCommitmentWithConfig", name, owner, secret, resolver, addr)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MakeCommitmentWithConfig is a free data retrieval call binding the contract method 0x3d86c52f.
//
// Solidity: function makeCommitmentWithConfig(string name, address owner, bytes32 secret, address resolver, address addr) pure returns(bytes32)
func (_Registrar *RegistrarSession) MakeCommitmentWithConfig(name string, owner common.Address, secret [32]byte, resolver common.Address, addr common.Address) ([32]byte, error) {
	return _Registrar.Contract.MakeCommitmentWithConfig(&_Registrar.CallOpts, name, owner, secret, resolver, addr)
}

// MakeCommitmentWithConfig is a free data retrieval call binding the contract method 0x3d86c52f.
//
// Solidity: function makeCommitmentWithConfig(string name, address owner, bytes32 secret, address resolver, address addr) pure returns(bytes32)
func (_Registrar *RegistrarCallerSession) MakeCommitmentWithConfig(name string, owner common.Address, secret [32]byte, resolver common.Address, addr common.Address) ([32]byte, error) {
	return _Registrar.Contract.MakeCommitmentWithConfig(&_Registrar.CallOpts, name, owner, secret, resolver, addr)
}

// MaxCommitmentAge is a free data retrieval call binding the contract method 0xce1e09c0.
//
// Solidity: function maxCommitmentAge() view returns(uint256)
func (_Registrar *RegistrarCaller) MaxCommitmentAge(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Registrar.contract.Call(opts, &out, "maxCommitmentAge")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxCommitmentAge is a free data retrieval call binding the contract method 0xce1e09c0.
//
// Solidity: function maxCommitmentAge() view returns(uint256)
func (_Registrar *RegistrarSession) MaxCommitmentAge() (*big.Int, error) {
	return _Registrar.Contract.MaxCommitmentAge(&_Registrar.CallOpts)
}

// MaxCommitmentAge is a free data retrieval call binding the contract method 0xce1e09c0.
//
// Solidity: function maxCommitmentAge() view returns(uint256)
func (_Registrar *RegistrarCallerSession) MaxCommitmentAge() (*big.Int, error) {
	return _Registrar.Contract.MaxCommitmentAge(&_Registrar.CallOpts)
}

// MinCommitmentAge is a free data retrieval call binding the contract method 0x8d839ffe.
//
// Solidity: function minCommitmentAge() view returns(uint256)
func (_Registrar *RegistrarCaller) MinCommitmentAge(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Registrar.contract.Call(opts, &out, "minCommitmentAge")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinCommitmentAge is a free data retrieval call binding the contract method 0x8d839ffe.
//
// Solidity: function minCommitmentAge() view returns(uint256)
func (_Registrar *RegistrarSession) MinCommitmentAge() (*big.Int, error) {
	return _Registrar.Contract.MinCommitmentAge(&_Registrar.CallOpts)
}

// MinCommitmentAge is a free data retrieval call binding the contract method 0x8d839ffe.
//
// Solidity: function minCommitmentAge() view returns(uint256)
func (_Registrar *RegistrarCallerSession) MinCommitmentAge() (*big.Int, error) {
	return _Registrar.Contract.MinCommitmentAge(&_Registrar.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Registrar *RegistrarCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Registrar.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Registrar *RegistrarSession) Owner() (common.Address, error) {
	return _Registrar.Contract.Owner(&_Registrar.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Registrar *RegistrarCallerSession) Owner() (common.Address, error) {
	return _Registrar.Contract.Owner(&_Registrar.CallOpts)
}

// RentPrice is a free data retrieval call binding the contract method 0x83e7f6ff.
//
// Solidity: function rentPrice(string name, uint256 duration) view returns(uint256)
func (_Registrar *RegistrarCaller) RentPrice(opts *bind.CallOpts, name string, duration *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Registrar.contract.Call(opts, &out, "rentPrice", name, duration)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RentPrice is a free data retrieval call binding the contract method 0x83e7f6ff.
//
// Solidity: function rentPrice(string name, uint256 duration) view returns(uint256)
func (_Registrar *RegistrarSession) RentPrice(name string, duration *big.Int) (*big.Int, error) {
	return _Registrar.Contract.RentPrice(&_Registrar.CallOpts, name, duration)
}

// RentPrice is a free data retrieval call binding the contract method 0x83e7f6ff.
//
// Solidity: function rentPrice(string name, uint256 duration) view returns(uint256)
func (_Registrar *RegistrarCallerSession) RentPrice(name string, duration *big.Int) (*big.Int, error) {
	return _Registrar.Contract.RentPrice(&_Registrar.CallOpts, name, duration)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) pure returns(bool)
func (_Registrar *RegistrarCaller) SupportsInterface(opts *bind.CallOpts, interfaceID [4]byte) (bool, error) {
	var out []interface{}
	err := _Registrar.contract.Call(opts, &out, "supportsInterface", interfaceID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) pure returns(bool)
func (_Registrar *RegistrarSession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _Registrar.Contract.SupportsInterface(&_Registrar.CallOpts, interfaceID)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) pure returns(bool)
func (_Registrar *RegistrarCallerSession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _Registrar.Contract.SupportsInterface(&_Registrar.CallOpts, interfaceID)
}

// Valid is a free data retrieval call binding the contract method 0x9791c097.
//
// Solidity: function valid(string name) pure returns(bool)
func (_Registrar *RegistrarCaller) Valid(opts *bind.CallOpts, name string) (bool, error) {
	var out []interface{}
	err := _Registrar.contract.Call(opts, &out, "valid", name)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Valid is a free data retrieval call binding the contract method 0x9791c097.
//
// Solidity: function valid(string name) pure returns(bool)
func (_Registrar *RegistrarSession) Valid(name string) (bool, error) {
	return _Registrar.Contract.Valid(&_Registrar.CallOpts, name)
}

// Valid is a free data retrieval call binding the contract method 0x9791c097.
//
// Solidity: function valid(string name) pure returns(bool)
func (_Registrar *RegistrarCallerSession) Valid(name string) (bool, error) {
	return _Registrar.Contract.Valid(&_Registrar.CallOpts, name)
}

// Commit is a paid mutator transaction binding the contract method 0xf14fcbc8.
//
// Solidity: function commit(bytes32 commitment) returns()
func (_Registrar *RegistrarTransactor) Commit(opts *bind.TransactOpts, commitment [32]byte) (*types.Transaction, error) {
	return _Registrar.contract.Transact(opts, "commit", commitment)
}

// Commit is a paid mutator transaction binding the contract method 0xf14fcbc8.
//
// Solidity: function commit(bytes32 commitment) returns()
func (_Registrar *RegistrarSession) Commit(commitment [32]byte) (*types.Transaction, error) {
	return _Registrar.Contract.Commit(&_Registrar.TransactOpts, commitment)
}

// Commit is a paid mutator transaction binding the contract method 0xf14fcbc8.
//
// Solidity: function commit(bytes32 commitment) returns()
func (_Registrar *RegistrarTransactorSession) Commit(commitment [32]byte) (*types.Transaction, error) {
	return _Registrar.Contract.Commit(&_Registrar.TransactOpts, commitment)
}

// Register is a paid mutator transaction binding the contract method 0x85f6d155.
//
// Solidity: function register(string name, address owner, uint256 duration, bytes32 secret) payable returns()
func (_Registrar *RegistrarTransactor) Register(opts *bind.TransactOpts, name string, owner common.Address, duration *big.Int, secret [32]byte) (*types.Transaction, error) {
	return _Registrar.contract.Transact(opts, "register", name, owner, duration, secret)
}

// Register is a paid mutator transaction binding the contract method 0x85f6d155.
//
// Solidity: function register(string name, address owner, uint256 duration, bytes32 secret) payable returns()
func (_Registrar *RegistrarSession) Register(name string, owner common.Address, duration *big.Int, secret [32]byte) (*types.Transaction, error) {
	return _Registrar.Contract.Register(&_Registrar.TransactOpts, name, owner, duration, secret)
}

// Register is a paid mutator transaction binding the contract method 0x85f6d155.
//
// Solidity: function register(string name, address owner, uint256 duration, bytes32 secret) payable returns()
func (_Registrar *RegistrarTransactorSession) Register(name string, owner common.Address, duration *big.Int, secret [32]byte) (*types.Transaction, error) {
	return _Registrar.Contract.Register(&_Registrar.TransactOpts, name, owner, duration, secret)
}

// RegisterWithConfig is a paid mutator transaction binding the contract method 0xf7a16963.
//
// Solidity: function registerWithConfig(string name, address owner, uint256 duration, bytes32 secret, address resolver, address addr) payable returns()
func (_Registrar *RegistrarTransactor) RegisterWithConfig(opts *bind.TransactOpts, name string, owner common.Address, duration *big.Int, secret [32]byte, resolver common.Address, addr common.Address) (*types.Transaction, error) {
	return _Registrar.contract.Transact(opts, "registerWithConfig", name, owner, duration, secret, resolver, addr)
}

// RegisterWithConfig is a paid mutator transaction binding the contract method 0xf7a16963.
//
// Solidity: function registerWithConfig(string name, address owner, uint256 duration, bytes32 secret, address resolver, address addr) payable returns()
func (_Registrar *RegistrarSession) RegisterWithConfig(name string, owner common.Address, duration *big.Int, secret [32]byte, resolver common.Address, addr common.Address) (*types.Transaction, error) {
	return _Registrar.Contract.RegisterWithConfig(&_Registrar.TransactOpts, name, owner, duration, secret, resolver, addr)
}

// RegisterWithConfig is a paid mutator transaction binding the contract method 0xf7a16963.
//
// Solidity: function registerWithConfig(string name, address owner, uint256 duration, bytes32 secret, address resolver, address addr) payable returns()
func (_Registrar *RegistrarTransactorSession) RegisterWithConfig(name string, owner common.Address, duration *big.Int, secret [32]byte, resolver common.Address, addr common.Address) (*types.Transaction, error) {
	return _Registrar.Contract.RegisterWithConfig(&_Registrar.TransactOpts, name, owner, duration, secret, resolver, addr)
}

// Renew is a paid mutator transaction binding the contract method 0xacf1a841.
//
// Solidity: function renew(string name, uint256 duration) payable returns()
func (_Registrar *RegistrarTransactor) Renew(opts *bind.TransactOpts, name string, duration *big.Int) (*types.Transaction, error) {
	return _Registrar.contract.Transact(opts, "renew", name, duration)
}

// Renew is a paid mutator transaction binding the contract method 0xacf1a841.
//
// Solidity: function renew(string name, uint256 duration) payable returns()
func (_Registrar *RegistrarSession) Renew(name string, duration *big.Int) (*types.Transaction, error) {
	return _Registrar.Contract.Renew(&_Registrar.TransactOpts, name, duration)
}

// Renew is a paid mutator transaction binding the contract method 0xacf1a841.
//
// Solidity: function renew(string name, uint256 duration) payable returns()
func (_Registrar *RegistrarTransactorSession) Renew(name string, duration *big.Int) (*types.Transaction, error) {
	return _Registrar.Contract.Renew(&_Registrar.TransactOpts, name, duration)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Registrar *RegistrarTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registrar.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Registrar *RegistrarSession) RenounceOwnership() (*types.Transaction, error) {
	return _Registrar.Contract.RenounceOwnership(&_Registrar.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Registrar *RegistrarTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Registrar.Contract.RenounceOwnership(&_Registrar.TransactOpts)
}

// SetCommitmentAges is a paid mutator transaction binding the contract method 0x7e324479.
//
// Solidity: function setCommitmentAges(uint256 _minCommitmentAge, uint256 _maxCommitmentAge) returns()
func (_Registrar *RegistrarTransactor) SetCommitmentAges(opts *bind.TransactOpts, _minCommitmentAge *big.Int, _maxCommitmentAge *big.Int) (*types.Transaction, error) {
	return _Registrar.contract.Transact(opts, "setCommitmentAges", _minCommitmentAge, _maxCommitmentAge)
}

// SetCommitmentAges is a paid mutator transaction binding the contract method 0x7e324479.
//
// Solidity: function setCommitmentAges(uint256 _minCommitmentAge, uint256 _maxCommitmentAge) returns()
func (_Registrar *RegistrarSession) SetCommitmentAges(_minCommitmentAge *big.Int, _maxCommitmentAge *big.Int) (*types.Transaction, error) {
	return _Registrar.Contract.SetCommitmentAges(&_Registrar.TransactOpts, _minCommitmentAge, _maxCommitmentAge)
}

// SetCommitmentAges is a paid mutator transaction binding the contract method 0x7e324479.
//
// Solidity: function setCommitmentAges(uint256 _minCommitmentAge, uint256 _maxCommitmentAge) returns()
func (_Registrar *RegistrarTransactorSession) SetCommitmentAges(_minCommitmentAge *big.Int, _maxCommitmentAge *big.Int) (*types.Transaction, error) {
	return _Registrar.Contract.SetCommitmentAges(&_Registrar.TransactOpts, _minCommitmentAge, _maxCommitmentAge)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address _prices) returns()
func (_Registrar *RegistrarTransactor) SetPriceOracle(opts *bind.TransactOpts, _prices common.Address) (*types.Transaction, error) {
	return _Registrar.contract.Transact(opts, "setPriceOracle", _prices)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address _prices) returns()
func (_Registrar *RegistrarSession) SetPriceOracle(_prices common.Address) (*types.Transaction, error) {
	return _Registrar.Contract.SetPriceOracle(&_Registrar.TransactOpts, _prices)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address _prices) returns()
func (_Registrar *RegistrarTransactorSession) SetPriceOracle(_prices common.Address) (*types.Transaction, error) {
	return _Registrar.Contract.SetPriceOracle(&_Registrar.TransactOpts, _prices)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Registrar *RegistrarTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Registrar.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Registrar *RegistrarSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Registrar.Contract.TransferOwnership(&_Registrar.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Registrar *RegistrarTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Registrar.Contract.TransferOwnership(&_Registrar.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Registrar *RegistrarTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registrar.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Registrar *RegistrarSession) Withdraw() (*types.Transaction, error) {
	return _Registrar.Contract.Withdraw(&_Registrar.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Registrar *RegistrarTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Registrar.Contract.Withdraw(&_Registrar.TransactOpts)
}

// RegistrarNameRegisteredIterator is returned from FilterNameRegistered and is used to iterate over the raw logs and unpacked data for NameRegistered events raised by the Registrar contract.
type RegistrarNameRegisteredIterator struct {
	Event *RegistrarNameRegistered // Event containing the contract specifics and raw log

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
func (it *RegistrarNameRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrarNameRegistered)
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
		it.Event = new(RegistrarNameRegistered)
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
func (it *RegistrarNameRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrarNameRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrarNameRegistered represents a NameRegistered event raised by the Registrar contract.
type RegistrarNameRegistered struct {
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
func (_Registrar *RegistrarFilterer) FilterNameRegistered(opts *bind.FilterOpts, label [][32]byte, owner []common.Address) (*RegistrarNameRegisteredIterator, error) {

	var labelRule []interface{}
	for _, labelItem := range label {
		labelRule = append(labelRule, labelItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Registrar.contract.FilterLogs(opts, "NameRegistered", labelRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &RegistrarNameRegisteredIterator{contract: _Registrar.contract, event: "NameRegistered", logs: logs, sub: sub}, nil
}

// WatchNameRegistered is a free log subscription operation binding the contract event 0xca6abbe9d7f11422cb6ca7629fbf6fe9efb1c621f71ce8f02b9f2a230097404f.
//
// Solidity: event NameRegistered(string name, bytes32 indexed label, address indexed owner, uint256 cost, uint256 expires)
func (_Registrar *RegistrarFilterer) WatchNameRegistered(opts *bind.WatchOpts, sink chan<- *RegistrarNameRegistered, label [][32]byte, owner []common.Address) (event.Subscription, error) {

	var labelRule []interface{}
	for _, labelItem := range label {
		labelRule = append(labelRule, labelItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Registrar.contract.WatchLogs(opts, "NameRegistered", labelRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrarNameRegistered)
				if err := _Registrar.contract.UnpackLog(event, "NameRegistered", log); err != nil {
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
func (_Registrar *RegistrarFilterer) ParseNameRegistered(log types.Log) (*RegistrarNameRegistered, error) {
	event := new(RegistrarNameRegistered)
	if err := _Registrar.contract.UnpackLog(event, "NameRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrarNameRenewedIterator is returned from FilterNameRenewed and is used to iterate over the raw logs and unpacked data for NameRenewed events raised by the Registrar contract.
type RegistrarNameRenewedIterator struct {
	Event *RegistrarNameRenewed // Event containing the contract specifics and raw log

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
func (it *RegistrarNameRenewedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrarNameRenewed)
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
		it.Event = new(RegistrarNameRenewed)
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
func (it *RegistrarNameRenewedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrarNameRenewedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrarNameRenewed represents a NameRenewed event raised by the Registrar contract.
type RegistrarNameRenewed struct {
	Name    string
	Label   [32]byte
	Cost    *big.Int
	Expires *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNameRenewed is a free log retrieval operation binding the contract event 0x3da24c024582931cfaf8267d8ed24d13a82a8068d5bd337d30ec45cea4e506ae.
//
// Solidity: event NameRenewed(string name, bytes32 indexed label, uint256 cost, uint256 expires)
func (_Registrar *RegistrarFilterer) FilterNameRenewed(opts *bind.FilterOpts, label [][32]byte) (*RegistrarNameRenewedIterator, error) {

	var labelRule []interface{}
	for _, labelItem := range label {
		labelRule = append(labelRule, labelItem)
	}

	logs, sub, err := _Registrar.contract.FilterLogs(opts, "NameRenewed", labelRule)
	if err != nil {
		return nil, err
	}
	return &RegistrarNameRenewedIterator{contract: _Registrar.contract, event: "NameRenewed", logs: logs, sub: sub}, nil
}

// WatchNameRenewed is a free log subscription operation binding the contract event 0x3da24c024582931cfaf8267d8ed24d13a82a8068d5bd337d30ec45cea4e506ae.
//
// Solidity: event NameRenewed(string name, bytes32 indexed label, uint256 cost, uint256 expires)
func (_Registrar *RegistrarFilterer) WatchNameRenewed(opts *bind.WatchOpts, sink chan<- *RegistrarNameRenewed, label [][32]byte) (event.Subscription, error) {

	var labelRule []interface{}
	for _, labelItem := range label {
		labelRule = append(labelRule, labelItem)
	}

	logs, sub, err := _Registrar.contract.WatchLogs(opts, "NameRenewed", labelRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrarNameRenewed)
				if err := _Registrar.contract.UnpackLog(event, "NameRenewed", log); err != nil {
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
func (_Registrar *RegistrarFilterer) ParseNameRenewed(log types.Log) (*RegistrarNameRenewed, error) {
	event := new(RegistrarNameRenewed)
	if err := _Registrar.contract.UnpackLog(event, "NameRenewed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrarNewPriceOracleIterator is returned from FilterNewPriceOracle and is used to iterate over the raw logs and unpacked data for NewPriceOracle events raised by the Registrar contract.
type RegistrarNewPriceOracleIterator struct {
	Event *RegistrarNewPriceOracle // Event containing the contract specifics and raw log

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
func (it *RegistrarNewPriceOracleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrarNewPriceOracle)
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
		it.Event = new(RegistrarNewPriceOracle)
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
func (it *RegistrarNewPriceOracleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrarNewPriceOracleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrarNewPriceOracle represents a NewPriceOracle event raised by the Registrar contract.
type RegistrarNewPriceOracle struct {
	Oracle common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNewPriceOracle is a free log retrieval operation binding the contract event 0xf261845a790fe29bbd6631e2ca4a5bdc83e6eed7c3271d9590d97287e00e9123.
//
// Solidity: event NewPriceOracle(address indexed oracle)
func (_Registrar *RegistrarFilterer) FilterNewPriceOracle(opts *bind.FilterOpts, oracle []common.Address) (*RegistrarNewPriceOracleIterator, error) {

	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _Registrar.contract.FilterLogs(opts, "NewPriceOracle", oracleRule)
	if err != nil {
		return nil, err
	}
	return &RegistrarNewPriceOracleIterator{contract: _Registrar.contract, event: "NewPriceOracle", logs: logs, sub: sub}, nil
}

// WatchNewPriceOracle is a free log subscription operation binding the contract event 0xf261845a790fe29bbd6631e2ca4a5bdc83e6eed7c3271d9590d97287e00e9123.
//
// Solidity: event NewPriceOracle(address indexed oracle)
func (_Registrar *RegistrarFilterer) WatchNewPriceOracle(opts *bind.WatchOpts, sink chan<- *RegistrarNewPriceOracle, oracle []common.Address) (event.Subscription, error) {

	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _Registrar.contract.WatchLogs(opts, "NewPriceOracle", oracleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrarNewPriceOracle)
				if err := _Registrar.contract.UnpackLog(event, "NewPriceOracle", log); err != nil {
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
func (_Registrar *RegistrarFilterer) ParseNewPriceOracle(log types.Log) (*RegistrarNewPriceOracle, error) {
	event := new(RegistrarNewPriceOracle)
	if err := _Registrar.contract.UnpackLog(event, "NewPriceOracle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrarOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Registrar contract.
type RegistrarOwnershipTransferredIterator struct {
	Event *RegistrarOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RegistrarOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrarOwnershipTransferred)
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
		it.Event = new(RegistrarOwnershipTransferred)
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
func (it *RegistrarOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrarOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrarOwnershipTransferred represents a OwnershipTransferred event raised by the Registrar contract.
type RegistrarOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Registrar *RegistrarFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RegistrarOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Registrar.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RegistrarOwnershipTransferredIterator{contract: _Registrar.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Registrar *RegistrarFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RegistrarOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Registrar.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrarOwnershipTransferred)
				if err := _Registrar.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Registrar *RegistrarFilterer) ParseOwnershipTransferred(log types.Log) (*RegistrarOwnershipTransferred, error) {
	event := new(RegistrarOwnershipTransferred)
	if err := _Registrar.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
