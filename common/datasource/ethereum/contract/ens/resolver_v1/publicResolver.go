// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package resolverv1

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

// ResolverV1MetaData contains all meta data concerning the ResolverV1 contract.
var ResolverV1MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractENS\",\"name\":\"_ens\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"contentType\",\"type\":\"uint256\"}],\"name\":\"ABIChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"AddrChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"coinType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"newAddress\",\"type\":\"bytes\"}],\"name\":\"AddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isAuthorised\",\"type\":\"bool\"}],\"name\":\"AuthorisationChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"hash\",\"type\":\"bytes\"}],\"name\":\"ContenthashChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"resource\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"record\",\"type\":\"bytes\"}],\"name\":\"DNSRecordChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"resource\",\"type\":\"uint16\"}],\"name\":\"DNSRecordDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"DNSZoneCleared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes4\",\"name\":\"interfaceID\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"implementer\",\"type\":\"address\"}],\"name\":\"InterfaceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"NameChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"x\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"y\",\"type\":\"bytes32\"}],\"name\":\"PubkeyChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"indexedKey\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"TextChanged\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"contentTypes\",\"type\":\"uint256\"}],\"name\":\"ABI\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"addr\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"coinType\",\"type\":\"uint256\"}],\"name\":\"addr\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"authorisations\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"clearDNSZone\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"contenthash\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"resource\",\"type\":\"uint16\"}],\"name\":\"dnsRecord\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"}],\"name\":\"hasDNSRecords\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"bytes4\",\"name\":\"interfaceID\",\"type\":\"bytes4\"}],\"name\":\"interfaceImplementer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"pubkey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"x\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"y\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"contentType\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"setABI\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"coinType\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"a\",\"type\":\"bytes\"}],\"name\":\"setAddr\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"setAddr\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isAuthorised\",\"type\":\"bool\"}],\"name\":\"setAuthorisation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"hash\",\"type\":\"bytes\"}],\"name\":\"setContenthash\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"setDNSRecords\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"bytes4\",\"name\":\"interfaceID\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"implementer\",\"type\":\"address\"}],\"name\":\"setInterface\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"setName\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"x\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"y\",\"type\":\"bytes32\"}],\"name\":\"setPubkey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"setText\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceID\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"text\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ResolverV1ABI is the input ABI used to generate the binding from.
// Deprecated: Use ResolverV1MetaData.ABI instead.
var ResolverV1ABI = ResolverV1MetaData.ABI

// ResolverV1 is an auto generated Go binding around an Ethereum contract.
type ResolverV1 struct {
	ResolverV1Caller     // Read-only binding to the contract
	ResolverV1Transactor // Write-only binding to the contract
	ResolverV1Filterer   // Log filterer for contract events
}

// ResolverV1Caller is an auto generated read-only Go binding around an Ethereum contract.
type ResolverV1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ResolverV1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ResolverV1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ResolverV1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ResolverV1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ResolverV1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ResolverV1Session struct {
	Contract     *ResolverV1       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ResolverV1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ResolverV1CallerSession struct {
	Contract *ResolverV1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ResolverV1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ResolverV1TransactorSession struct {
	Contract     *ResolverV1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ResolverV1Raw is an auto generated low-level Go binding around an Ethereum contract.
type ResolverV1Raw struct {
	Contract *ResolverV1 // Generic contract binding to access the raw methods on
}

// ResolverV1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ResolverV1CallerRaw struct {
	Contract *ResolverV1Caller // Generic read-only contract binding to access the raw methods on
}

// ResolverV1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ResolverV1TransactorRaw struct {
	Contract *ResolverV1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewResolverV1 creates a new instance of ResolverV1, bound to a specific deployed contract.
func NewResolverV1(address common.Address, backend bind.ContractBackend) (*ResolverV1, error) {
	contract, err := bindResolverV1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ResolverV1{ResolverV1Caller: ResolverV1Caller{contract: contract}, ResolverV1Transactor: ResolverV1Transactor{contract: contract}, ResolverV1Filterer: ResolverV1Filterer{contract: contract}}, nil
}

// NewResolverV1Caller creates a new read-only instance of ResolverV1, bound to a specific deployed contract.
func NewResolverV1Caller(address common.Address, caller bind.ContractCaller) (*ResolverV1Caller, error) {
	contract, err := bindResolverV1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ResolverV1Caller{contract: contract}, nil
}

// NewResolverV1Transactor creates a new write-only instance of ResolverV1, bound to a specific deployed contract.
func NewResolverV1Transactor(address common.Address, transactor bind.ContractTransactor) (*ResolverV1Transactor, error) {
	contract, err := bindResolverV1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ResolverV1Transactor{contract: contract}, nil
}

// NewResolverV1Filterer creates a new log filterer instance of ResolverV1, bound to a specific deployed contract.
func NewResolverV1Filterer(address common.Address, filterer bind.ContractFilterer) (*ResolverV1Filterer, error) {
	contract, err := bindResolverV1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ResolverV1Filterer{contract: contract}, nil
}

// bindResolverV1 binds a generic wrapper to an already deployed contract.
func bindResolverV1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ResolverV1MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ResolverV1 *ResolverV1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ResolverV1.Contract.ResolverV1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ResolverV1 *ResolverV1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ResolverV1.Contract.ResolverV1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ResolverV1 *ResolverV1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ResolverV1.Contract.ResolverV1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ResolverV1 *ResolverV1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ResolverV1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ResolverV1 *ResolverV1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ResolverV1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ResolverV1 *ResolverV1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ResolverV1.Contract.contract.Transact(opts, method, params...)
}

// ABI is a free data retrieval call binding the contract method 0x2203ab56.
//
// Solidity: function ABI(bytes32 node, uint256 contentTypes) view returns(uint256, bytes)
func (_ResolverV1 *ResolverV1Caller) ABI(opts *bind.CallOpts, node [32]byte, contentTypes *big.Int) (*big.Int, []byte, error) {
	var out []interface{}
	err := _ResolverV1.contract.Call(opts, &out, "ABI", node, contentTypes)

	if err != nil {
		return *new(*big.Int), *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return out0, out1, err

}

// ABI is a free data retrieval call binding the contract method 0x2203ab56.
//
// Solidity: function ABI(bytes32 node, uint256 contentTypes) view returns(uint256, bytes)
func (_ResolverV1 *ResolverV1Session) ABI(node [32]byte, contentTypes *big.Int) (*big.Int, []byte, error) {
	return _ResolverV1.Contract.ABI(&_ResolverV1.CallOpts, node, contentTypes)
}

// ABI is a free data retrieval call binding the contract method 0x2203ab56.
//
// Solidity: function ABI(bytes32 node, uint256 contentTypes) view returns(uint256, bytes)
func (_ResolverV1 *ResolverV1CallerSession) ABI(node [32]byte, contentTypes *big.Int) (*big.Int, []byte, error) {
	return _ResolverV1.Contract.ABI(&_ResolverV1.CallOpts, node, contentTypes)
}

// Addr is a free data retrieval call binding the contract method 0x3b3b57de.
//
// Solidity: function addr(bytes32 node) view returns(address)
func (_ResolverV1 *ResolverV1Caller) Addr(opts *bind.CallOpts, node [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ResolverV1.contract.Call(opts, &out, "addr", node)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Addr is a free data retrieval call binding the contract method 0x3b3b57de.
//
// Solidity: function addr(bytes32 node) view returns(address)
func (_ResolverV1 *ResolverV1Session) Addr(node [32]byte) (common.Address, error) {
	return _ResolverV1.Contract.Addr(&_ResolverV1.CallOpts, node)
}

// Addr is a free data retrieval call binding the contract method 0x3b3b57de.
//
// Solidity: function addr(bytes32 node) view returns(address)
func (_ResolverV1 *ResolverV1CallerSession) Addr(node [32]byte) (common.Address, error) {
	return _ResolverV1.Contract.Addr(&_ResolverV1.CallOpts, node)
}

// Addr0 is a free data retrieval call binding the contract method 0xf1cb7e06.
//
// Solidity: function addr(bytes32 node, uint256 coinType) view returns(bytes)
func (_ResolverV1 *ResolverV1Caller) Addr0(opts *bind.CallOpts, node [32]byte, coinType *big.Int) ([]byte, error) {
	var out []interface{}
	err := _ResolverV1.contract.Call(opts, &out, "addr0", node, coinType)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Addr0 is a free data retrieval call binding the contract method 0xf1cb7e06.
//
// Solidity: function addr(bytes32 node, uint256 coinType) view returns(bytes)
func (_ResolverV1 *ResolverV1Session) Addr0(node [32]byte, coinType *big.Int) ([]byte, error) {
	return _ResolverV1.Contract.Addr0(&_ResolverV1.CallOpts, node, coinType)
}

// Addr0 is a free data retrieval call binding the contract method 0xf1cb7e06.
//
// Solidity: function addr(bytes32 node, uint256 coinType) view returns(bytes)
func (_ResolverV1 *ResolverV1CallerSession) Addr0(node [32]byte, coinType *big.Int) ([]byte, error) {
	return _ResolverV1.Contract.Addr0(&_ResolverV1.CallOpts, node, coinType)
}

// Authorisations is a free data retrieval call binding the contract method 0xf86bc879.
//
// Solidity: function authorisations(bytes32 , address , address ) view returns(bool)
func (_ResolverV1 *ResolverV1Caller) Authorisations(opts *bind.CallOpts, arg0 [32]byte, arg1 common.Address, arg2 common.Address) (bool, error) {
	var out []interface{}
	err := _ResolverV1.contract.Call(opts, &out, "authorisations", arg0, arg1, arg2)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Authorisations is a free data retrieval call binding the contract method 0xf86bc879.
//
// Solidity: function authorisations(bytes32 , address , address ) view returns(bool)
func (_ResolverV1 *ResolverV1Session) Authorisations(arg0 [32]byte, arg1 common.Address, arg2 common.Address) (bool, error) {
	return _ResolverV1.Contract.Authorisations(&_ResolverV1.CallOpts, arg0, arg1, arg2)
}

// Authorisations is a free data retrieval call binding the contract method 0xf86bc879.
//
// Solidity: function authorisations(bytes32 , address , address ) view returns(bool)
func (_ResolverV1 *ResolverV1CallerSession) Authorisations(arg0 [32]byte, arg1 common.Address, arg2 common.Address) (bool, error) {
	return _ResolverV1.Contract.Authorisations(&_ResolverV1.CallOpts, arg0, arg1, arg2)
}

// Contenthash is a free data retrieval call binding the contract method 0xbc1c58d1.
//
// Solidity: function contenthash(bytes32 node) view returns(bytes)
func (_ResolverV1 *ResolverV1Caller) Contenthash(opts *bind.CallOpts, node [32]byte) ([]byte, error) {
	var out []interface{}
	err := _ResolverV1.contract.Call(opts, &out, "contenthash", node)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Contenthash is a free data retrieval call binding the contract method 0xbc1c58d1.
//
// Solidity: function contenthash(bytes32 node) view returns(bytes)
func (_ResolverV1 *ResolverV1Session) Contenthash(node [32]byte) ([]byte, error) {
	return _ResolverV1.Contract.Contenthash(&_ResolverV1.CallOpts, node)
}

// Contenthash is a free data retrieval call binding the contract method 0xbc1c58d1.
//
// Solidity: function contenthash(bytes32 node) view returns(bytes)
func (_ResolverV1 *ResolverV1CallerSession) Contenthash(node [32]byte) ([]byte, error) {
	return _ResolverV1.Contract.Contenthash(&_ResolverV1.CallOpts, node)
}

// DnsRecord is a free data retrieval call binding the contract method 0xa8fa5682.
//
// Solidity: function dnsRecord(bytes32 node, bytes32 name, uint16 resource) view returns(bytes)
func (_ResolverV1 *ResolverV1Caller) DnsRecord(opts *bind.CallOpts, node [32]byte, name [32]byte, resource uint16) ([]byte, error) {
	var out []interface{}
	err := _ResolverV1.contract.Call(opts, &out, "dnsRecord", node, name, resource)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// DnsRecord is a free data retrieval call binding the contract method 0xa8fa5682.
//
// Solidity: function dnsRecord(bytes32 node, bytes32 name, uint16 resource) view returns(bytes)
func (_ResolverV1 *ResolverV1Session) DnsRecord(node [32]byte, name [32]byte, resource uint16) ([]byte, error) {
	return _ResolverV1.Contract.DnsRecord(&_ResolverV1.CallOpts, node, name, resource)
}

// DnsRecord is a free data retrieval call binding the contract method 0xa8fa5682.
//
// Solidity: function dnsRecord(bytes32 node, bytes32 name, uint16 resource) view returns(bytes)
func (_ResolverV1 *ResolverV1CallerSession) DnsRecord(node [32]byte, name [32]byte, resource uint16) ([]byte, error) {
	return _ResolverV1.Contract.DnsRecord(&_ResolverV1.CallOpts, node, name, resource)
}

// HasDNSRecords is a free data retrieval call binding the contract method 0x4cbf6ba4.
//
// Solidity: function hasDNSRecords(bytes32 node, bytes32 name) view returns(bool)
func (_ResolverV1 *ResolverV1Caller) HasDNSRecords(opts *bind.CallOpts, node [32]byte, name [32]byte) (bool, error) {
	var out []interface{}
	err := _ResolverV1.contract.Call(opts, &out, "hasDNSRecords", node, name)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasDNSRecords is a free data retrieval call binding the contract method 0x4cbf6ba4.
//
// Solidity: function hasDNSRecords(bytes32 node, bytes32 name) view returns(bool)
func (_ResolverV1 *ResolverV1Session) HasDNSRecords(node [32]byte, name [32]byte) (bool, error) {
	return _ResolverV1.Contract.HasDNSRecords(&_ResolverV1.CallOpts, node, name)
}

// HasDNSRecords is a free data retrieval call binding the contract method 0x4cbf6ba4.
//
// Solidity: function hasDNSRecords(bytes32 node, bytes32 name) view returns(bool)
func (_ResolverV1 *ResolverV1CallerSession) HasDNSRecords(node [32]byte, name [32]byte) (bool, error) {
	return _ResolverV1.Contract.HasDNSRecords(&_ResolverV1.CallOpts, node, name)
}

// InterfaceImplementer is a free data retrieval call binding the contract method 0x124a319c.
//
// Solidity: function interfaceImplementer(bytes32 node, bytes4 interfaceID) view returns(address)
func (_ResolverV1 *ResolverV1Caller) InterfaceImplementer(opts *bind.CallOpts, node [32]byte, interfaceID [4]byte) (common.Address, error) {
	var out []interface{}
	err := _ResolverV1.contract.Call(opts, &out, "interfaceImplementer", node, interfaceID)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InterfaceImplementer is a free data retrieval call binding the contract method 0x124a319c.
//
// Solidity: function interfaceImplementer(bytes32 node, bytes4 interfaceID) view returns(address)
func (_ResolverV1 *ResolverV1Session) InterfaceImplementer(node [32]byte, interfaceID [4]byte) (common.Address, error) {
	return _ResolverV1.Contract.InterfaceImplementer(&_ResolverV1.CallOpts, node, interfaceID)
}

// InterfaceImplementer is a free data retrieval call binding the contract method 0x124a319c.
//
// Solidity: function interfaceImplementer(bytes32 node, bytes4 interfaceID) view returns(address)
func (_ResolverV1 *ResolverV1CallerSession) InterfaceImplementer(node [32]byte, interfaceID [4]byte) (common.Address, error) {
	return _ResolverV1.Contract.InterfaceImplementer(&_ResolverV1.CallOpts, node, interfaceID)
}

// Name is a free data retrieval call binding the contract method 0x691f3431.
//
// Solidity: function name(bytes32 node) view returns(string)
func (_ResolverV1 *ResolverV1Caller) Name(opts *bind.CallOpts, node [32]byte) (string, error) {
	var out []interface{}
	err := _ResolverV1.contract.Call(opts, &out, "name", node)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x691f3431.
//
// Solidity: function name(bytes32 node) view returns(string)
func (_ResolverV1 *ResolverV1Session) Name(node [32]byte) (string, error) {
	return _ResolverV1.Contract.Name(&_ResolverV1.CallOpts, node)
}

// Name is a free data retrieval call binding the contract method 0x691f3431.
//
// Solidity: function name(bytes32 node) view returns(string)
func (_ResolverV1 *ResolverV1CallerSession) Name(node [32]byte) (string, error) {
	return _ResolverV1.Contract.Name(&_ResolverV1.CallOpts, node)
}

// Pubkey is a free data retrieval call binding the contract method 0xc8690233.
//
// Solidity: function pubkey(bytes32 node) view returns(bytes32 x, bytes32 y)
func (_ResolverV1 *ResolverV1Caller) Pubkey(opts *bind.CallOpts, node [32]byte) (struct {
	X [32]byte
	Y [32]byte
}, error) {
	var out []interface{}
	err := _ResolverV1.contract.Call(opts, &out, "pubkey", node)

	outstruct := new(struct {
		X [32]byte
		Y [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.X = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Y = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// Pubkey is a free data retrieval call binding the contract method 0xc8690233.
//
// Solidity: function pubkey(bytes32 node) view returns(bytes32 x, bytes32 y)
func (_ResolverV1 *ResolverV1Session) Pubkey(node [32]byte) (struct {
	X [32]byte
	Y [32]byte
}, error) {
	return _ResolverV1.Contract.Pubkey(&_ResolverV1.CallOpts, node)
}

// Pubkey is a free data retrieval call binding the contract method 0xc8690233.
//
// Solidity: function pubkey(bytes32 node) view returns(bytes32 x, bytes32 y)
func (_ResolverV1 *ResolverV1CallerSession) Pubkey(node [32]byte) (struct {
	X [32]byte
	Y [32]byte
}, error) {
	return _ResolverV1.Contract.Pubkey(&_ResolverV1.CallOpts, node)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) pure returns(bool)
func (_ResolverV1 *ResolverV1Caller) SupportsInterface(opts *bind.CallOpts, interfaceID [4]byte) (bool, error) {
	var out []interface{}
	err := _ResolverV1.contract.Call(opts, &out, "supportsInterface", interfaceID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) pure returns(bool)
func (_ResolverV1 *ResolverV1Session) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _ResolverV1.Contract.SupportsInterface(&_ResolverV1.CallOpts, interfaceID)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) pure returns(bool)
func (_ResolverV1 *ResolverV1CallerSession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _ResolverV1.Contract.SupportsInterface(&_ResolverV1.CallOpts, interfaceID)
}

// Text is a free data retrieval call binding the contract method 0x59d1d43c.
//
// Solidity: function text(bytes32 node, string key) view returns(string)
func (_ResolverV1 *ResolverV1Caller) Text(opts *bind.CallOpts, node [32]byte, key string) (string, error) {
	var out []interface{}
	err := _ResolverV1.contract.Call(opts, &out, "text", node, key)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Text is a free data retrieval call binding the contract method 0x59d1d43c.
//
// Solidity: function text(bytes32 node, string key) view returns(string)
func (_ResolverV1 *ResolverV1Session) Text(node [32]byte, key string) (string, error) {
	return _ResolverV1.Contract.Text(&_ResolverV1.CallOpts, node, key)
}

// Text is a free data retrieval call binding the contract method 0x59d1d43c.
//
// Solidity: function text(bytes32 node, string key) view returns(string)
func (_ResolverV1 *ResolverV1CallerSession) Text(node [32]byte, key string) (string, error) {
	return _ResolverV1.Contract.Text(&_ResolverV1.CallOpts, node, key)
}

// ClearDNSZone is a paid mutator transaction binding the contract method 0xad5780af.
//
// Solidity: function clearDNSZone(bytes32 node) returns()
func (_ResolverV1 *ResolverV1Transactor) ClearDNSZone(opts *bind.TransactOpts, node [32]byte) (*types.Transaction, error) {
	return _ResolverV1.contract.Transact(opts, "clearDNSZone", node)
}

// ClearDNSZone is a paid mutator transaction binding the contract method 0xad5780af.
//
// Solidity: function clearDNSZone(bytes32 node) returns()
func (_ResolverV1 *ResolverV1Session) ClearDNSZone(node [32]byte) (*types.Transaction, error) {
	return _ResolverV1.Contract.ClearDNSZone(&_ResolverV1.TransactOpts, node)
}

// ClearDNSZone is a paid mutator transaction binding the contract method 0xad5780af.
//
// Solidity: function clearDNSZone(bytes32 node) returns()
func (_ResolverV1 *ResolverV1TransactorSession) ClearDNSZone(node [32]byte) (*types.Transaction, error) {
	return _ResolverV1.Contract.ClearDNSZone(&_ResolverV1.TransactOpts, node)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_ResolverV1 *ResolverV1Transactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _ResolverV1.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_ResolverV1 *ResolverV1Session) Multicall(data [][]byte) (*types.Transaction, error) {
	return _ResolverV1.Contract.Multicall(&_ResolverV1.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_ResolverV1 *ResolverV1TransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _ResolverV1.Contract.Multicall(&_ResolverV1.TransactOpts, data)
}

// SetABI is a paid mutator transaction binding the contract method 0x623195b0.
//
// Solidity: function setABI(bytes32 node, uint256 contentType, bytes data) returns()
func (_ResolverV1 *ResolverV1Transactor) SetABI(opts *bind.TransactOpts, node [32]byte, contentType *big.Int, data []byte) (*types.Transaction, error) {
	return _ResolverV1.contract.Transact(opts, "setABI", node, contentType, data)
}

// SetABI is a paid mutator transaction binding the contract method 0x623195b0.
//
// Solidity: function setABI(bytes32 node, uint256 contentType, bytes data) returns()
func (_ResolverV1 *ResolverV1Session) SetABI(node [32]byte, contentType *big.Int, data []byte) (*types.Transaction, error) {
	return _ResolverV1.Contract.SetABI(&_ResolverV1.TransactOpts, node, contentType, data)
}

// SetABI is a paid mutator transaction binding the contract method 0x623195b0.
//
// Solidity: function setABI(bytes32 node, uint256 contentType, bytes data) returns()
func (_ResolverV1 *ResolverV1TransactorSession) SetABI(node [32]byte, contentType *big.Int, data []byte) (*types.Transaction, error) {
	return _ResolverV1.Contract.SetABI(&_ResolverV1.TransactOpts, node, contentType, data)
}

// SetAddr is a paid mutator transaction binding the contract method 0x8b95dd71.
//
// Solidity: function setAddr(bytes32 node, uint256 coinType, bytes a) returns()
func (_ResolverV1 *ResolverV1Transactor) SetAddr(opts *bind.TransactOpts, node [32]byte, coinType *big.Int, a []byte) (*types.Transaction, error) {
	return _ResolverV1.contract.Transact(opts, "setAddr", node, coinType, a)
}

// SetAddr is a paid mutator transaction binding the contract method 0x8b95dd71.
//
// Solidity: function setAddr(bytes32 node, uint256 coinType, bytes a) returns()
func (_ResolverV1 *ResolverV1Session) SetAddr(node [32]byte, coinType *big.Int, a []byte) (*types.Transaction, error) {
	return _ResolverV1.Contract.SetAddr(&_ResolverV1.TransactOpts, node, coinType, a)
}

// SetAddr is a paid mutator transaction binding the contract method 0x8b95dd71.
//
// Solidity: function setAddr(bytes32 node, uint256 coinType, bytes a) returns()
func (_ResolverV1 *ResolverV1TransactorSession) SetAddr(node [32]byte, coinType *big.Int, a []byte) (*types.Transaction, error) {
	return _ResolverV1.Contract.SetAddr(&_ResolverV1.TransactOpts, node, coinType, a)
}

// SetAddr0 is a paid mutator transaction binding the contract method 0xd5fa2b00.
//
// Solidity: function setAddr(bytes32 node, address a) returns()
func (_ResolverV1 *ResolverV1Transactor) SetAddr0(opts *bind.TransactOpts, node [32]byte, a common.Address) (*types.Transaction, error) {
	return _ResolverV1.contract.Transact(opts, "setAddr0", node, a)
}

// SetAddr0 is a paid mutator transaction binding the contract method 0xd5fa2b00.
//
// Solidity: function setAddr(bytes32 node, address a) returns()
func (_ResolverV1 *ResolverV1Session) SetAddr0(node [32]byte, a common.Address) (*types.Transaction, error) {
	return _ResolverV1.Contract.SetAddr0(&_ResolverV1.TransactOpts, node, a)
}

// SetAddr0 is a paid mutator transaction binding the contract method 0xd5fa2b00.
//
// Solidity: function setAddr(bytes32 node, address a) returns()
func (_ResolverV1 *ResolverV1TransactorSession) SetAddr0(node [32]byte, a common.Address) (*types.Transaction, error) {
	return _ResolverV1.Contract.SetAddr0(&_ResolverV1.TransactOpts, node, a)
}

// SetAuthorisation is a paid mutator transaction binding the contract method 0x3e9ce794.
//
// Solidity: function setAuthorisation(bytes32 node, address target, bool isAuthorised) returns()
func (_ResolverV1 *ResolverV1Transactor) SetAuthorisation(opts *bind.TransactOpts, node [32]byte, target common.Address, isAuthorised bool) (*types.Transaction, error) {
	return _ResolverV1.contract.Transact(opts, "setAuthorisation", node, target, isAuthorised)
}

// SetAuthorisation is a paid mutator transaction binding the contract method 0x3e9ce794.
//
// Solidity: function setAuthorisation(bytes32 node, address target, bool isAuthorised) returns()
func (_ResolverV1 *ResolverV1Session) SetAuthorisation(node [32]byte, target common.Address, isAuthorised bool) (*types.Transaction, error) {
	return _ResolverV1.Contract.SetAuthorisation(&_ResolverV1.TransactOpts, node, target, isAuthorised)
}

// SetAuthorisation is a paid mutator transaction binding the contract method 0x3e9ce794.
//
// Solidity: function setAuthorisation(bytes32 node, address target, bool isAuthorised) returns()
func (_ResolverV1 *ResolverV1TransactorSession) SetAuthorisation(node [32]byte, target common.Address, isAuthorised bool) (*types.Transaction, error) {
	return _ResolverV1.Contract.SetAuthorisation(&_ResolverV1.TransactOpts, node, target, isAuthorised)
}

// SetContenthash is a paid mutator transaction binding the contract method 0x304e6ade.
//
// Solidity: function setContenthash(bytes32 node, bytes hash) returns()
func (_ResolverV1 *ResolverV1Transactor) SetContenthash(opts *bind.TransactOpts, node [32]byte, hash []byte) (*types.Transaction, error) {
	return _ResolverV1.contract.Transact(opts, "setContenthash", node, hash)
}

// SetContenthash is a paid mutator transaction binding the contract method 0x304e6ade.
//
// Solidity: function setContenthash(bytes32 node, bytes hash) returns()
func (_ResolverV1 *ResolverV1Session) SetContenthash(node [32]byte, hash []byte) (*types.Transaction, error) {
	return _ResolverV1.Contract.SetContenthash(&_ResolverV1.TransactOpts, node, hash)
}

// SetContenthash is a paid mutator transaction binding the contract method 0x304e6ade.
//
// Solidity: function setContenthash(bytes32 node, bytes hash) returns()
func (_ResolverV1 *ResolverV1TransactorSession) SetContenthash(node [32]byte, hash []byte) (*types.Transaction, error) {
	return _ResolverV1.Contract.SetContenthash(&_ResolverV1.TransactOpts, node, hash)
}

// SetDNSRecords is a paid mutator transaction binding the contract method 0x0af179d7.
//
// Solidity: function setDNSRecords(bytes32 node, bytes data) returns()
func (_ResolverV1 *ResolverV1Transactor) SetDNSRecords(opts *bind.TransactOpts, node [32]byte, data []byte) (*types.Transaction, error) {
	return _ResolverV1.contract.Transact(opts, "setDNSRecords", node, data)
}

// SetDNSRecords is a paid mutator transaction binding the contract method 0x0af179d7.
//
// Solidity: function setDNSRecords(bytes32 node, bytes data) returns()
func (_ResolverV1 *ResolverV1Session) SetDNSRecords(node [32]byte, data []byte) (*types.Transaction, error) {
	return _ResolverV1.Contract.SetDNSRecords(&_ResolverV1.TransactOpts, node, data)
}

// SetDNSRecords is a paid mutator transaction binding the contract method 0x0af179d7.
//
// Solidity: function setDNSRecords(bytes32 node, bytes data) returns()
func (_ResolverV1 *ResolverV1TransactorSession) SetDNSRecords(node [32]byte, data []byte) (*types.Transaction, error) {
	return _ResolverV1.Contract.SetDNSRecords(&_ResolverV1.TransactOpts, node, data)
}

// SetInterface is a paid mutator transaction binding the contract method 0xe59d895d.
//
// Solidity: function setInterface(bytes32 node, bytes4 interfaceID, address implementer) returns()
func (_ResolverV1 *ResolverV1Transactor) SetInterface(opts *bind.TransactOpts, node [32]byte, interfaceID [4]byte, implementer common.Address) (*types.Transaction, error) {
	return _ResolverV1.contract.Transact(opts, "setInterface", node, interfaceID, implementer)
}

// SetInterface is a paid mutator transaction binding the contract method 0xe59d895d.
//
// Solidity: function setInterface(bytes32 node, bytes4 interfaceID, address implementer) returns()
func (_ResolverV1 *ResolverV1Session) SetInterface(node [32]byte, interfaceID [4]byte, implementer common.Address) (*types.Transaction, error) {
	return _ResolverV1.Contract.SetInterface(&_ResolverV1.TransactOpts, node, interfaceID, implementer)
}

// SetInterface is a paid mutator transaction binding the contract method 0xe59d895d.
//
// Solidity: function setInterface(bytes32 node, bytes4 interfaceID, address implementer) returns()
func (_ResolverV1 *ResolverV1TransactorSession) SetInterface(node [32]byte, interfaceID [4]byte, implementer common.Address) (*types.Transaction, error) {
	return _ResolverV1.Contract.SetInterface(&_ResolverV1.TransactOpts, node, interfaceID, implementer)
}

// SetName is a paid mutator transaction binding the contract method 0x77372213.
//
// Solidity: function setName(bytes32 node, string name) returns()
func (_ResolverV1 *ResolverV1Transactor) SetName(opts *bind.TransactOpts, node [32]byte, name string) (*types.Transaction, error) {
	return _ResolverV1.contract.Transact(opts, "setName", node, name)
}

// SetName is a paid mutator transaction binding the contract method 0x77372213.
//
// Solidity: function setName(bytes32 node, string name) returns()
func (_ResolverV1 *ResolverV1Session) SetName(node [32]byte, name string) (*types.Transaction, error) {
	return _ResolverV1.Contract.SetName(&_ResolverV1.TransactOpts, node, name)
}

// SetName is a paid mutator transaction binding the contract method 0x77372213.
//
// Solidity: function setName(bytes32 node, string name) returns()
func (_ResolverV1 *ResolverV1TransactorSession) SetName(node [32]byte, name string) (*types.Transaction, error) {
	return _ResolverV1.Contract.SetName(&_ResolverV1.TransactOpts, node, name)
}

// SetPubkey is a paid mutator transaction binding the contract method 0x29cd62ea.
//
// Solidity: function setPubkey(bytes32 node, bytes32 x, bytes32 y) returns()
func (_ResolverV1 *ResolverV1Transactor) SetPubkey(opts *bind.TransactOpts, node [32]byte, x [32]byte, y [32]byte) (*types.Transaction, error) {
	return _ResolverV1.contract.Transact(opts, "setPubkey", node, x, y)
}

// SetPubkey is a paid mutator transaction binding the contract method 0x29cd62ea.
//
// Solidity: function setPubkey(bytes32 node, bytes32 x, bytes32 y) returns()
func (_ResolverV1 *ResolverV1Session) SetPubkey(node [32]byte, x [32]byte, y [32]byte) (*types.Transaction, error) {
	return _ResolverV1.Contract.SetPubkey(&_ResolverV1.TransactOpts, node, x, y)
}

// SetPubkey is a paid mutator transaction binding the contract method 0x29cd62ea.
//
// Solidity: function setPubkey(bytes32 node, bytes32 x, bytes32 y) returns()
func (_ResolverV1 *ResolverV1TransactorSession) SetPubkey(node [32]byte, x [32]byte, y [32]byte) (*types.Transaction, error) {
	return _ResolverV1.Contract.SetPubkey(&_ResolverV1.TransactOpts, node, x, y)
}

// SetText is a paid mutator transaction binding the contract method 0x10f13a8c.
//
// Solidity: function setText(bytes32 node, string key, string value) returns()
func (_ResolverV1 *ResolverV1Transactor) SetText(opts *bind.TransactOpts, node [32]byte, key string, value string) (*types.Transaction, error) {
	return _ResolverV1.contract.Transact(opts, "setText", node, key, value)
}

// SetText is a paid mutator transaction binding the contract method 0x10f13a8c.
//
// Solidity: function setText(bytes32 node, string key, string value) returns()
func (_ResolverV1 *ResolverV1Session) SetText(node [32]byte, key string, value string) (*types.Transaction, error) {
	return _ResolverV1.Contract.SetText(&_ResolverV1.TransactOpts, node, key, value)
}

// SetText is a paid mutator transaction binding the contract method 0x10f13a8c.
//
// Solidity: function setText(bytes32 node, string key, string value) returns()
func (_ResolverV1 *ResolverV1TransactorSession) SetText(node [32]byte, key string, value string) (*types.Transaction, error) {
	return _ResolverV1.Contract.SetText(&_ResolverV1.TransactOpts, node, key, value)
}

// ResolverV1ABIChangedIterator is returned from FilterABIChanged and is used to iterate over the raw logs and unpacked data for ABIChanged events raised by the ResolverV1 contract.
type ResolverV1ABIChangedIterator struct {
	Event *ResolverV1ABIChanged // Event containing the contract specifics and raw log

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
func (it *ResolverV1ABIChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResolverV1ABIChanged)
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
		it.Event = new(ResolverV1ABIChanged)
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
func (it *ResolverV1ABIChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResolverV1ABIChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResolverV1ABIChanged represents a ABIChanged event raised by the ResolverV1 contract.
type ResolverV1ABIChanged struct {
	Node        [32]byte
	ContentType *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterABIChanged is a free log retrieval operation binding the contract event 0xaa121bbeef5f32f5961a2a28966e769023910fc9479059ee3495d4c1a696efe3.
//
// Solidity: event ABIChanged(bytes32 indexed node, uint256 indexed contentType)
func (_ResolverV1 *ResolverV1Filterer) FilterABIChanged(opts *bind.FilterOpts, node [][32]byte, contentType []*big.Int) (*ResolverV1ABIChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var contentTypeRule []interface{}
	for _, contentTypeItem := range contentType {
		contentTypeRule = append(contentTypeRule, contentTypeItem)
	}

	logs, sub, err := _ResolverV1.contract.FilterLogs(opts, "ABIChanged", nodeRule, contentTypeRule)
	if err != nil {
		return nil, err
	}
	return &ResolverV1ABIChangedIterator{contract: _ResolverV1.contract, event: "ABIChanged", logs: logs, sub: sub}, nil
}

// WatchABIChanged is a free log subscription operation binding the contract event 0xaa121bbeef5f32f5961a2a28966e769023910fc9479059ee3495d4c1a696efe3.
//
// Solidity: event ABIChanged(bytes32 indexed node, uint256 indexed contentType)
func (_ResolverV1 *ResolverV1Filterer) WatchABIChanged(opts *bind.WatchOpts, sink chan<- *ResolverV1ABIChanged, node [][32]byte, contentType []*big.Int) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var contentTypeRule []interface{}
	for _, contentTypeItem := range contentType {
		contentTypeRule = append(contentTypeRule, contentTypeItem)
	}

	logs, sub, err := _ResolverV1.contract.WatchLogs(opts, "ABIChanged", nodeRule, contentTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResolverV1ABIChanged)
				if err := _ResolverV1.contract.UnpackLog(event, "ABIChanged", log); err != nil {
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

// ParseABIChanged is a log parse operation binding the contract event 0xaa121bbeef5f32f5961a2a28966e769023910fc9479059ee3495d4c1a696efe3.
//
// Solidity: event ABIChanged(bytes32 indexed node, uint256 indexed contentType)
func (_ResolverV1 *ResolverV1Filterer) ParseABIChanged(log types.Log) (*ResolverV1ABIChanged, error) {
	event := new(ResolverV1ABIChanged)
	if err := _ResolverV1.contract.UnpackLog(event, "ABIChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ResolverV1AddrChangedIterator is returned from FilterAddrChanged and is used to iterate over the raw logs and unpacked data for AddrChanged events raised by the ResolverV1 contract.
type ResolverV1AddrChangedIterator struct {
	Event *ResolverV1AddrChanged // Event containing the contract specifics and raw log

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
func (it *ResolverV1AddrChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResolverV1AddrChanged)
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
		it.Event = new(ResolverV1AddrChanged)
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
func (it *ResolverV1AddrChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResolverV1AddrChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResolverV1AddrChanged represents a AddrChanged event raised by the ResolverV1 contract.
type ResolverV1AddrChanged struct {
	Node [32]byte
	A    common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAddrChanged is a free log retrieval operation binding the contract event 0x52d7d861f09ab3d26239d492e8968629f95e9e318cf0b73bfddc441522a15fd2.
//
// Solidity: event AddrChanged(bytes32 indexed node, address a)
func (_ResolverV1 *ResolverV1Filterer) FilterAddrChanged(opts *bind.FilterOpts, node [][32]byte) (*ResolverV1AddrChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ResolverV1.contract.FilterLogs(opts, "AddrChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &ResolverV1AddrChangedIterator{contract: _ResolverV1.contract, event: "AddrChanged", logs: logs, sub: sub}, nil
}

// WatchAddrChanged is a free log subscription operation binding the contract event 0x52d7d861f09ab3d26239d492e8968629f95e9e318cf0b73bfddc441522a15fd2.
//
// Solidity: event AddrChanged(bytes32 indexed node, address a)
func (_ResolverV1 *ResolverV1Filterer) WatchAddrChanged(opts *bind.WatchOpts, sink chan<- *ResolverV1AddrChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ResolverV1.contract.WatchLogs(opts, "AddrChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResolverV1AddrChanged)
				if err := _ResolverV1.contract.UnpackLog(event, "AddrChanged", log); err != nil {
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

// ParseAddrChanged is a log parse operation binding the contract event 0x52d7d861f09ab3d26239d492e8968629f95e9e318cf0b73bfddc441522a15fd2.
//
// Solidity: event AddrChanged(bytes32 indexed node, address a)
func (_ResolverV1 *ResolverV1Filterer) ParseAddrChanged(log types.Log) (*ResolverV1AddrChanged, error) {
	event := new(ResolverV1AddrChanged)
	if err := _ResolverV1.contract.UnpackLog(event, "AddrChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ResolverV1AddressChangedIterator is returned from FilterAddressChanged and is used to iterate over the raw logs and unpacked data for AddressChanged events raised by the ResolverV1 contract.
type ResolverV1AddressChangedIterator struct {
	Event *ResolverV1AddressChanged // Event containing the contract specifics and raw log

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
func (it *ResolverV1AddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResolverV1AddressChanged)
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
		it.Event = new(ResolverV1AddressChanged)
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
func (it *ResolverV1AddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResolverV1AddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResolverV1AddressChanged represents a AddressChanged event raised by the ResolverV1 contract.
type ResolverV1AddressChanged struct {
	Node       [32]byte
	CoinType   *big.Int
	NewAddress []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAddressChanged is a free log retrieval operation binding the contract event 0x65412581168e88a1e60c6459d7f44ae83ad0832e670826c05a4e2476b57af752.
//
// Solidity: event AddressChanged(bytes32 indexed node, uint256 coinType, bytes newAddress)
func (_ResolverV1 *ResolverV1Filterer) FilterAddressChanged(opts *bind.FilterOpts, node [][32]byte) (*ResolverV1AddressChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ResolverV1.contract.FilterLogs(opts, "AddressChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &ResolverV1AddressChangedIterator{contract: _ResolverV1.contract, event: "AddressChanged", logs: logs, sub: sub}, nil
}

// WatchAddressChanged is a free log subscription operation binding the contract event 0x65412581168e88a1e60c6459d7f44ae83ad0832e670826c05a4e2476b57af752.
//
// Solidity: event AddressChanged(bytes32 indexed node, uint256 coinType, bytes newAddress)
func (_ResolverV1 *ResolverV1Filterer) WatchAddressChanged(opts *bind.WatchOpts, sink chan<- *ResolverV1AddressChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ResolverV1.contract.WatchLogs(opts, "AddressChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResolverV1AddressChanged)
				if err := _ResolverV1.contract.UnpackLog(event, "AddressChanged", log); err != nil {
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

// ParseAddressChanged is a log parse operation binding the contract event 0x65412581168e88a1e60c6459d7f44ae83ad0832e670826c05a4e2476b57af752.
//
// Solidity: event AddressChanged(bytes32 indexed node, uint256 coinType, bytes newAddress)
func (_ResolverV1 *ResolverV1Filterer) ParseAddressChanged(log types.Log) (*ResolverV1AddressChanged, error) {
	event := new(ResolverV1AddressChanged)
	if err := _ResolverV1.contract.UnpackLog(event, "AddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ResolverV1AuthorisationChangedIterator is returned from FilterAuthorisationChanged and is used to iterate over the raw logs and unpacked data for AuthorisationChanged events raised by the ResolverV1 contract.
type ResolverV1AuthorisationChangedIterator struct {
	Event *ResolverV1AuthorisationChanged // Event containing the contract specifics and raw log

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
func (it *ResolverV1AuthorisationChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResolverV1AuthorisationChanged)
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
		it.Event = new(ResolverV1AuthorisationChanged)
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
func (it *ResolverV1AuthorisationChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResolverV1AuthorisationChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResolverV1AuthorisationChanged represents a AuthorisationChanged event raised by the ResolverV1 contract.
type ResolverV1AuthorisationChanged struct {
	Node         [32]byte
	Owner        common.Address
	Target       common.Address
	IsAuthorised bool
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAuthorisationChanged is a free log retrieval operation binding the contract event 0xe1c5610a6e0cbe10764ecd182adcef1ec338dc4e199c99c32ce98f38e12791df.
//
// Solidity: event AuthorisationChanged(bytes32 indexed node, address indexed owner, address indexed target, bool isAuthorised)
func (_ResolverV1 *ResolverV1Filterer) FilterAuthorisationChanged(opts *bind.FilterOpts, node [][32]byte, owner []common.Address, target []common.Address) (*ResolverV1AuthorisationChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _ResolverV1.contract.FilterLogs(opts, "AuthorisationChanged", nodeRule, ownerRule, targetRule)
	if err != nil {
		return nil, err
	}
	return &ResolverV1AuthorisationChangedIterator{contract: _ResolverV1.contract, event: "AuthorisationChanged", logs: logs, sub: sub}, nil
}

// WatchAuthorisationChanged is a free log subscription operation binding the contract event 0xe1c5610a6e0cbe10764ecd182adcef1ec338dc4e199c99c32ce98f38e12791df.
//
// Solidity: event AuthorisationChanged(bytes32 indexed node, address indexed owner, address indexed target, bool isAuthorised)
func (_ResolverV1 *ResolverV1Filterer) WatchAuthorisationChanged(opts *bind.WatchOpts, sink chan<- *ResolverV1AuthorisationChanged, node [][32]byte, owner []common.Address, target []common.Address) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _ResolverV1.contract.WatchLogs(opts, "AuthorisationChanged", nodeRule, ownerRule, targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResolverV1AuthorisationChanged)
				if err := _ResolverV1.contract.UnpackLog(event, "AuthorisationChanged", log); err != nil {
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

// ParseAuthorisationChanged is a log parse operation binding the contract event 0xe1c5610a6e0cbe10764ecd182adcef1ec338dc4e199c99c32ce98f38e12791df.
//
// Solidity: event AuthorisationChanged(bytes32 indexed node, address indexed owner, address indexed target, bool isAuthorised)
func (_ResolverV1 *ResolverV1Filterer) ParseAuthorisationChanged(log types.Log) (*ResolverV1AuthorisationChanged, error) {
	event := new(ResolverV1AuthorisationChanged)
	if err := _ResolverV1.contract.UnpackLog(event, "AuthorisationChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ResolverV1ContenthashChangedIterator is returned from FilterContenthashChanged and is used to iterate over the raw logs and unpacked data for ContenthashChanged events raised by the ResolverV1 contract.
type ResolverV1ContenthashChangedIterator struct {
	Event *ResolverV1ContenthashChanged // Event containing the contract specifics and raw log

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
func (it *ResolverV1ContenthashChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResolverV1ContenthashChanged)
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
		it.Event = new(ResolverV1ContenthashChanged)
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
func (it *ResolverV1ContenthashChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResolverV1ContenthashChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResolverV1ContenthashChanged represents a ContenthashChanged event raised by the ResolverV1 contract.
type ResolverV1ContenthashChanged struct {
	Node [32]byte
	Hash []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterContenthashChanged is a free log retrieval operation binding the contract event 0xe379c1624ed7e714cc0937528a32359d69d5281337765313dba4e081b72d7578.
//
// Solidity: event ContenthashChanged(bytes32 indexed node, bytes hash)
func (_ResolverV1 *ResolverV1Filterer) FilterContenthashChanged(opts *bind.FilterOpts, node [][32]byte) (*ResolverV1ContenthashChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ResolverV1.contract.FilterLogs(opts, "ContenthashChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &ResolverV1ContenthashChangedIterator{contract: _ResolverV1.contract, event: "ContenthashChanged", logs: logs, sub: sub}, nil
}

// WatchContenthashChanged is a free log subscription operation binding the contract event 0xe379c1624ed7e714cc0937528a32359d69d5281337765313dba4e081b72d7578.
//
// Solidity: event ContenthashChanged(bytes32 indexed node, bytes hash)
func (_ResolverV1 *ResolverV1Filterer) WatchContenthashChanged(opts *bind.WatchOpts, sink chan<- *ResolverV1ContenthashChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ResolverV1.contract.WatchLogs(opts, "ContenthashChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResolverV1ContenthashChanged)
				if err := _ResolverV1.contract.UnpackLog(event, "ContenthashChanged", log); err != nil {
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

// ParseContenthashChanged is a log parse operation binding the contract event 0xe379c1624ed7e714cc0937528a32359d69d5281337765313dba4e081b72d7578.
//
// Solidity: event ContenthashChanged(bytes32 indexed node, bytes hash)
func (_ResolverV1 *ResolverV1Filterer) ParseContenthashChanged(log types.Log) (*ResolverV1ContenthashChanged, error) {
	event := new(ResolverV1ContenthashChanged)
	if err := _ResolverV1.contract.UnpackLog(event, "ContenthashChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ResolverV1DNSRecordChangedIterator is returned from FilterDNSRecordChanged and is used to iterate over the raw logs and unpacked data for DNSRecordChanged events raised by the ResolverV1 contract.
type ResolverV1DNSRecordChangedIterator struct {
	Event *ResolverV1DNSRecordChanged // Event containing the contract specifics and raw log

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
func (it *ResolverV1DNSRecordChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResolverV1DNSRecordChanged)
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
		it.Event = new(ResolverV1DNSRecordChanged)
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
func (it *ResolverV1DNSRecordChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResolverV1DNSRecordChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResolverV1DNSRecordChanged represents a DNSRecordChanged event raised by the ResolverV1 contract.
type ResolverV1DNSRecordChanged struct {
	Node     [32]byte
	Name     []byte
	Resource uint16
	Record   []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDNSRecordChanged is a free log retrieval operation binding the contract event 0x52a608b3303a48862d07a73d82fa221318c0027fbbcfb1b2329bface3f19ff2b.
//
// Solidity: event DNSRecordChanged(bytes32 indexed node, bytes name, uint16 resource, bytes record)
func (_ResolverV1 *ResolverV1Filterer) FilterDNSRecordChanged(opts *bind.FilterOpts, node [][32]byte) (*ResolverV1DNSRecordChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ResolverV1.contract.FilterLogs(opts, "DNSRecordChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &ResolverV1DNSRecordChangedIterator{contract: _ResolverV1.contract, event: "DNSRecordChanged", logs: logs, sub: sub}, nil
}

// WatchDNSRecordChanged is a free log subscription operation binding the contract event 0x52a608b3303a48862d07a73d82fa221318c0027fbbcfb1b2329bface3f19ff2b.
//
// Solidity: event DNSRecordChanged(bytes32 indexed node, bytes name, uint16 resource, bytes record)
func (_ResolverV1 *ResolverV1Filterer) WatchDNSRecordChanged(opts *bind.WatchOpts, sink chan<- *ResolverV1DNSRecordChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ResolverV1.contract.WatchLogs(opts, "DNSRecordChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResolverV1DNSRecordChanged)
				if err := _ResolverV1.contract.UnpackLog(event, "DNSRecordChanged", log); err != nil {
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

// ParseDNSRecordChanged is a log parse operation binding the contract event 0x52a608b3303a48862d07a73d82fa221318c0027fbbcfb1b2329bface3f19ff2b.
//
// Solidity: event DNSRecordChanged(bytes32 indexed node, bytes name, uint16 resource, bytes record)
func (_ResolverV1 *ResolverV1Filterer) ParseDNSRecordChanged(log types.Log) (*ResolverV1DNSRecordChanged, error) {
	event := new(ResolverV1DNSRecordChanged)
	if err := _ResolverV1.contract.UnpackLog(event, "DNSRecordChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ResolverV1DNSRecordDeletedIterator is returned from FilterDNSRecordDeleted and is used to iterate over the raw logs and unpacked data for DNSRecordDeleted events raised by the ResolverV1 contract.
type ResolverV1DNSRecordDeletedIterator struct {
	Event *ResolverV1DNSRecordDeleted // Event containing the contract specifics and raw log

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
func (it *ResolverV1DNSRecordDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResolverV1DNSRecordDeleted)
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
		it.Event = new(ResolverV1DNSRecordDeleted)
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
func (it *ResolverV1DNSRecordDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResolverV1DNSRecordDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResolverV1DNSRecordDeleted represents a DNSRecordDeleted event raised by the ResolverV1 contract.
type ResolverV1DNSRecordDeleted struct {
	Node     [32]byte
	Name     []byte
	Resource uint16
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDNSRecordDeleted is a free log retrieval operation binding the contract event 0x03528ed0c2a3ebc993b12ce3c16bb382f9c7d88ef7d8a1bf290eaf35955a1207.
//
// Solidity: event DNSRecordDeleted(bytes32 indexed node, bytes name, uint16 resource)
func (_ResolverV1 *ResolverV1Filterer) FilterDNSRecordDeleted(opts *bind.FilterOpts, node [][32]byte) (*ResolverV1DNSRecordDeletedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ResolverV1.contract.FilterLogs(opts, "DNSRecordDeleted", nodeRule)
	if err != nil {
		return nil, err
	}
	return &ResolverV1DNSRecordDeletedIterator{contract: _ResolverV1.contract, event: "DNSRecordDeleted", logs: logs, sub: sub}, nil
}

// WatchDNSRecordDeleted is a free log subscription operation binding the contract event 0x03528ed0c2a3ebc993b12ce3c16bb382f9c7d88ef7d8a1bf290eaf35955a1207.
//
// Solidity: event DNSRecordDeleted(bytes32 indexed node, bytes name, uint16 resource)
func (_ResolverV1 *ResolverV1Filterer) WatchDNSRecordDeleted(opts *bind.WatchOpts, sink chan<- *ResolverV1DNSRecordDeleted, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ResolverV1.contract.WatchLogs(opts, "DNSRecordDeleted", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResolverV1DNSRecordDeleted)
				if err := _ResolverV1.contract.UnpackLog(event, "DNSRecordDeleted", log); err != nil {
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

// ParseDNSRecordDeleted is a log parse operation binding the contract event 0x03528ed0c2a3ebc993b12ce3c16bb382f9c7d88ef7d8a1bf290eaf35955a1207.
//
// Solidity: event DNSRecordDeleted(bytes32 indexed node, bytes name, uint16 resource)
func (_ResolverV1 *ResolverV1Filterer) ParseDNSRecordDeleted(log types.Log) (*ResolverV1DNSRecordDeleted, error) {
	event := new(ResolverV1DNSRecordDeleted)
	if err := _ResolverV1.contract.UnpackLog(event, "DNSRecordDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ResolverV1DNSZoneClearedIterator is returned from FilterDNSZoneCleared and is used to iterate over the raw logs and unpacked data for DNSZoneCleared events raised by the ResolverV1 contract.
type ResolverV1DNSZoneClearedIterator struct {
	Event *ResolverV1DNSZoneCleared // Event containing the contract specifics and raw log

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
func (it *ResolverV1DNSZoneClearedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResolverV1DNSZoneCleared)
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
		it.Event = new(ResolverV1DNSZoneCleared)
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
func (it *ResolverV1DNSZoneClearedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResolverV1DNSZoneClearedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResolverV1DNSZoneCleared represents a DNSZoneCleared event raised by the ResolverV1 contract.
type ResolverV1DNSZoneCleared struct {
	Node [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterDNSZoneCleared is a free log retrieval operation binding the contract event 0xb757169b8492ca2f1c6619d9d76ce22803035c3b1d5f6930dffe7b127c1a1983.
//
// Solidity: event DNSZoneCleared(bytes32 indexed node)
func (_ResolverV1 *ResolverV1Filterer) FilterDNSZoneCleared(opts *bind.FilterOpts, node [][32]byte) (*ResolverV1DNSZoneClearedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ResolverV1.contract.FilterLogs(opts, "DNSZoneCleared", nodeRule)
	if err != nil {
		return nil, err
	}
	return &ResolverV1DNSZoneClearedIterator{contract: _ResolverV1.contract, event: "DNSZoneCleared", logs: logs, sub: sub}, nil
}

// WatchDNSZoneCleared is a free log subscription operation binding the contract event 0xb757169b8492ca2f1c6619d9d76ce22803035c3b1d5f6930dffe7b127c1a1983.
//
// Solidity: event DNSZoneCleared(bytes32 indexed node)
func (_ResolverV1 *ResolverV1Filterer) WatchDNSZoneCleared(opts *bind.WatchOpts, sink chan<- *ResolverV1DNSZoneCleared, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ResolverV1.contract.WatchLogs(opts, "DNSZoneCleared", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResolverV1DNSZoneCleared)
				if err := _ResolverV1.contract.UnpackLog(event, "DNSZoneCleared", log); err != nil {
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

// ParseDNSZoneCleared is a log parse operation binding the contract event 0xb757169b8492ca2f1c6619d9d76ce22803035c3b1d5f6930dffe7b127c1a1983.
//
// Solidity: event DNSZoneCleared(bytes32 indexed node)
func (_ResolverV1 *ResolverV1Filterer) ParseDNSZoneCleared(log types.Log) (*ResolverV1DNSZoneCleared, error) {
	event := new(ResolverV1DNSZoneCleared)
	if err := _ResolverV1.contract.UnpackLog(event, "DNSZoneCleared", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ResolverV1InterfaceChangedIterator is returned from FilterInterfaceChanged and is used to iterate over the raw logs and unpacked data for InterfaceChanged events raised by the ResolverV1 contract.
type ResolverV1InterfaceChangedIterator struct {
	Event *ResolverV1InterfaceChanged // Event containing the contract specifics and raw log

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
func (it *ResolverV1InterfaceChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResolverV1InterfaceChanged)
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
		it.Event = new(ResolverV1InterfaceChanged)
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
func (it *ResolverV1InterfaceChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResolverV1InterfaceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResolverV1InterfaceChanged represents a InterfaceChanged event raised by the ResolverV1 contract.
type ResolverV1InterfaceChanged struct {
	Node        [32]byte
	InterfaceID [4]byte
	Implementer common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterInterfaceChanged is a free log retrieval operation binding the contract event 0x7c69f06bea0bdef565b709e93a147836b0063ba2dd89f02d0b7e8d931e6a6daa.
//
// Solidity: event InterfaceChanged(bytes32 indexed node, bytes4 indexed interfaceID, address implementer)
func (_ResolverV1 *ResolverV1Filterer) FilterInterfaceChanged(opts *bind.FilterOpts, node [][32]byte, interfaceID [][4]byte) (*ResolverV1InterfaceChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var interfaceIDRule []interface{}
	for _, interfaceIDItem := range interfaceID {
		interfaceIDRule = append(interfaceIDRule, interfaceIDItem)
	}

	logs, sub, err := _ResolverV1.contract.FilterLogs(opts, "InterfaceChanged", nodeRule, interfaceIDRule)
	if err != nil {
		return nil, err
	}
	return &ResolverV1InterfaceChangedIterator{contract: _ResolverV1.contract, event: "InterfaceChanged", logs: logs, sub: sub}, nil
}

// WatchInterfaceChanged is a free log subscription operation binding the contract event 0x7c69f06bea0bdef565b709e93a147836b0063ba2dd89f02d0b7e8d931e6a6daa.
//
// Solidity: event InterfaceChanged(bytes32 indexed node, bytes4 indexed interfaceID, address implementer)
func (_ResolverV1 *ResolverV1Filterer) WatchInterfaceChanged(opts *bind.WatchOpts, sink chan<- *ResolverV1InterfaceChanged, node [][32]byte, interfaceID [][4]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var interfaceIDRule []interface{}
	for _, interfaceIDItem := range interfaceID {
		interfaceIDRule = append(interfaceIDRule, interfaceIDItem)
	}

	logs, sub, err := _ResolverV1.contract.WatchLogs(opts, "InterfaceChanged", nodeRule, interfaceIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResolverV1InterfaceChanged)
				if err := _ResolverV1.contract.UnpackLog(event, "InterfaceChanged", log); err != nil {
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

// ParseInterfaceChanged is a log parse operation binding the contract event 0x7c69f06bea0bdef565b709e93a147836b0063ba2dd89f02d0b7e8d931e6a6daa.
//
// Solidity: event InterfaceChanged(bytes32 indexed node, bytes4 indexed interfaceID, address implementer)
func (_ResolverV1 *ResolverV1Filterer) ParseInterfaceChanged(log types.Log) (*ResolverV1InterfaceChanged, error) {
	event := new(ResolverV1InterfaceChanged)
	if err := _ResolverV1.contract.UnpackLog(event, "InterfaceChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ResolverV1NameChangedIterator is returned from FilterNameChanged and is used to iterate over the raw logs and unpacked data for NameChanged events raised by the ResolverV1 contract.
type ResolverV1NameChangedIterator struct {
	Event *ResolverV1NameChanged // Event containing the contract specifics and raw log

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
func (it *ResolverV1NameChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResolverV1NameChanged)
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
		it.Event = new(ResolverV1NameChanged)
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
func (it *ResolverV1NameChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResolverV1NameChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResolverV1NameChanged represents a NameChanged event raised by the ResolverV1 contract.
type ResolverV1NameChanged struct {
	Node [32]byte
	Name string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterNameChanged is a free log retrieval operation binding the contract event 0xb7d29e911041e8d9b843369e890bcb72c9388692ba48b65ac54e7214c4c348f7.
//
// Solidity: event NameChanged(bytes32 indexed node, string name)
func (_ResolverV1 *ResolverV1Filterer) FilterNameChanged(opts *bind.FilterOpts, node [][32]byte) (*ResolverV1NameChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ResolverV1.contract.FilterLogs(opts, "NameChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &ResolverV1NameChangedIterator{contract: _ResolverV1.contract, event: "NameChanged", logs: logs, sub: sub}, nil
}

// WatchNameChanged is a free log subscription operation binding the contract event 0xb7d29e911041e8d9b843369e890bcb72c9388692ba48b65ac54e7214c4c348f7.
//
// Solidity: event NameChanged(bytes32 indexed node, string name)
func (_ResolverV1 *ResolverV1Filterer) WatchNameChanged(opts *bind.WatchOpts, sink chan<- *ResolverV1NameChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ResolverV1.contract.WatchLogs(opts, "NameChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResolverV1NameChanged)
				if err := _ResolverV1.contract.UnpackLog(event, "NameChanged", log); err != nil {
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

// ParseNameChanged is a log parse operation binding the contract event 0xb7d29e911041e8d9b843369e890bcb72c9388692ba48b65ac54e7214c4c348f7.
//
// Solidity: event NameChanged(bytes32 indexed node, string name)
func (_ResolverV1 *ResolverV1Filterer) ParseNameChanged(log types.Log) (*ResolverV1NameChanged, error) {
	event := new(ResolverV1NameChanged)
	if err := _ResolverV1.contract.UnpackLog(event, "NameChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ResolverV1PubkeyChangedIterator is returned from FilterPubkeyChanged and is used to iterate over the raw logs and unpacked data for PubkeyChanged events raised by the ResolverV1 contract.
type ResolverV1PubkeyChangedIterator struct {
	Event *ResolverV1PubkeyChanged // Event containing the contract specifics and raw log

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
func (it *ResolverV1PubkeyChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResolverV1PubkeyChanged)
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
		it.Event = new(ResolverV1PubkeyChanged)
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
func (it *ResolverV1PubkeyChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResolverV1PubkeyChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResolverV1PubkeyChanged represents a PubkeyChanged event raised by the ResolverV1 contract.
type ResolverV1PubkeyChanged struct {
	Node [32]byte
	X    [32]byte
	Y    [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPubkeyChanged is a free log retrieval operation binding the contract event 0x1d6f5e03d3f63eb58751986629a5439baee5079ff04f345becb66e23eb154e46.
//
// Solidity: event PubkeyChanged(bytes32 indexed node, bytes32 x, bytes32 y)
func (_ResolverV1 *ResolverV1Filterer) FilterPubkeyChanged(opts *bind.FilterOpts, node [][32]byte) (*ResolverV1PubkeyChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ResolverV1.contract.FilterLogs(opts, "PubkeyChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &ResolverV1PubkeyChangedIterator{contract: _ResolverV1.contract, event: "PubkeyChanged", logs: logs, sub: sub}, nil
}

// WatchPubkeyChanged is a free log subscription operation binding the contract event 0x1d6f5e03d3f63eb58751986629a5439baee5079ff04f345becb66e23eb154e46.
//
// Solidity: event PubkeyChanged(bytes32 indexed node, bytes32 x, bytes32 y)
func (_ResolverV1 *ResolverV1Filterer) WatchPubkeyChanged(opts *bind.WatchOpts, sink chan<- *ResolverV1PubkeyChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ResolverV1.contract.WatchLogs(opts, "PubkeyChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResolverV1PubkeyChanged)
				if err := _ResolverV1.contract.UnpackLog(event, "PubkeyChanged", log); err != nil {
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

// ParsePubkeyChanged is a log parse operation binding the contract event 0x1d6f5e03d3f63eb58751986629a5439baee5079ff04f345becb66e23eb154e46.
//
// Solidity: event PubkeyChanged(bytes32 indexed node, bytes32 x, bytes32 y)
func (_ResolverV1 *ResolverV1Filterer) ParsePubkeyChanged(log types.Log) (*ResolverV1PubkeyChanged, error) {
	event := new(ResolverV1PubkeyChanged)
	if err := _ResolverV1.contract.UnpackLog(event, "PubkeyChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ResolverV1TextChangedIterator is returned from FilterTextChanged and is used to iterate over the raw logs and unpacked data for TextChanged events raised by the ResolverV1 contract.
type ResolverV1TextChangedIterator struct {
	Event *ResolverV1TextChanged // Event containing the contract specifics and raw log

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
func (it *ResolverV1TextChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResolverV1TextChanged)
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
		it.Event = new(ResolverV1TextChanged)
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
func (it *ResolverV1TextChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResolverV1TextChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResolverV1TextChanged represents a TextChanged event raised by the ResolverV1 contract.
type ResolverV1TextChanged struct {
	Node       [32]byte
	IndexedKey common.Hash
	Key        string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTextChanged is a free log retrieval operation binding the contract event 0xd8c9334b1a9c2f9da342a0a2b32629c1a229b6445dad78947f674b44444a7550.
//
// Solidity: event TextChanged(bytes32 indexed node, string indexed indexedKey, string key)
func (_ResolverV1 *ResolverV1Filterer) FilterTextChanged(opts *bind.FilterOpts, node [][32]byte, indexedKey []string) (*ResolverV1TextChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var indexedKeyRule []interface{}
	for _, indexedKeyItem := range indexedKey {
		indexedKeyRule = append(indexedKeyRule, indexedKeyItem)
	}

	logs, sub, err := _ResolverV1.contract.FilterLogs(opts, "TextChanged", nodeRule, indexedKeyRule)
	if err != nil {
		return nil, err
	}
	return &ResolverV1TextChangedIterator{contract: _ResolverV1.contract, event: "TextChanged", logs: logs, sub: sub}, nil
}

// WatchTextChanged is a free log subscription operation binding the contract event 0xd8c9334b1a9c2f9da342a0a2b32629c1a229b6445dad78947f674b44444a7550.
//
// Solidity: event TextChanged(bytes32 indexed node, string indexed indexedKey, string key)
func (_ResolverV1 *ResolverV1Filterer) WatchTextChanged(opts *bind.WatchOpts, sink chan<- *ResolverV1TextChanged, node [][32]byte, indexedKey []string) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var indexedKeyRule []interface{}
	for _, indexedKeyItem := range indexedKey {
		indexedKeyRule = append(indexedKeyRule, indexedKeyItem)
	}

	logs, sub, err := _ResolverV1.contract.WatchLogs(opts, "TextChanged", nodeRule, indexedKeyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResolverV1TextChanged)
				if err := _ResolverV1.contract.UnpackLog(event, "TextChanged", log); err != nil {
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

// ParseTextChanged is a log parse operation binding the contract event 0xd8c9334b1a9c2f9da342a0a2b32629c1a229b6445dad78947f674b44444a7550.
//
// Solidity: event TextChanged(bytes32 indexed node, string indexed indexedKey, string key)
func (_ResolverV1 *ResolverV1Filterer) ParseTextChanged(log types.Log) (*ResolverV1TextChanged, error) {
	event := new(ResolverV1TextChanged)
	if err := _ResolverV1.contract.UnpackLog(event, "TextChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
