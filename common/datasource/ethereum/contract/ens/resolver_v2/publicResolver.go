// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package resolverv2

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

// ResolverV2MetaData contains all meta data concerning the ResolverV2 contract.
var ResolverV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractENS\",\"name\":\"_ens\",\"type\":\"address\"},{\"internalType\":\"contractINameWrapper\",\"name\":\"wrapperAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_trustedETHController\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_trustedReverseRegistrar\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"contentType\",\"type\":\"uint256\"}],\"name\":\"ABIChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"AddrChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"coinType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"newAddress\",\"type\":\"bytes\"}],\"name\":\"AddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"Approved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"hash\",\"type\":\"bytes\"}],\"name\":\"ContenthashChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"resource\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"record\",\"type\":\"bytes\"}],\"name\":\"DNSRecordChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"resource\",\"type\":\"uint16\"}],\"name\":\"DNSRecordDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"lastzonehash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"zonehash\",\"type\":\"bytes\"}],\"name\":\"DNSZonehashChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes4\",\"name\":\"interfaceID\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"implementer\",\"type\":\"address\"}],\"name\":\"InterfaceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"NameChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"x\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"y\",\"type\":\"bytes32\"}],\"name\":\"PubkeyChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"indexedKey\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"TextChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newVersion\",\"type\":\"uint64\"}],\"name\":\"VersionChanged\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"contentTypes\",\"type\":\"uint256\"}],\"name\":\"ABI\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"addr\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"coinType\",\"type\":\"uint256\"}],\"name\":\"addr\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"clearRecords\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"contenthash\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"resource\",\"type\":\"uint16\"}],\"name\":\"dnsRecord\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"}],\"name\":\"hasDNSRecords\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"bytes4\",\"name\":\"interfaceID\",\"type\":\"bytes4\"}],\"name\":\"interfaceImplementer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"}],\"name\":\"isApprovedFor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nodehash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicallWithNodeCheck\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"pubkey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"x\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"y\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"recordVersions\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"contentType\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"setABI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"coinType\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"a\",\"type\":\"bytes\"}],\"name\":\"setAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"setAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"hash\",\"type\":\"bytes\"}],\"name\":\"setContenthash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"setDNSRecords\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"bytes4\",\"name\":\"interfaceID\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"implementer\",\"type\":\"address\"}],\"name\":\"setInterface\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"newName\",\"type\":\"string\"}],\"name\":\"setName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"x\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"y\",\"type\":\"bytes32\"}],\"name\":\"setPubkey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"setText\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"hash\",\"type\":\"bytes\"}],\"name\":\"setZonehash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceID\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"text\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"zonehash\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ResolverV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use ResolverV2MetaData.ABI instead.
var ResolverV2ABI = ResolverV2MetaData.ABI

// ResolverV2 is an auto generated Go binding around an Ethereum contract.
type ResolverV2 struct {
	ResolverV2Caller     // Read-only binding to the contract
	ResolverV2Transactor // Write-only binding to the contract
	ResolverV2Filterer   // Log filterer for contract events
}

// ResolverV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type ResolverV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ResolverV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ResolverV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ResolverV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ResolverV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ResolverV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ResolverV2Session struct {
	Contract     *ResolverV2       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ResolverV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ResolverV2CallerSession struct {
	Contract *ResolverV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ResolverV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ResolverV2TransactorSession struct {
	Contract     *ResolverV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ResolverV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type ResolverV2Raw struct {
	Contract *ResolverV2 // Generic contract binding to access the raw methods on
}

// ResolverV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ResolverV2CallerRaw struct {
	Contract *ResolverV2Caller // Generic read-only contract binding to access the raw methods on
}

// ResolverV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ResolverV2TransactorRaw struct {
	Contract *ResolverV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewResolverV2 creates a new instance of ResolverV2, bound to a specific deployed contract.
func NewResolverV2(address common.Address, backend bind.ContractBackend) (*ResolverV2, error) {
	contract, err := bindResolverV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ResolverV2{ResolverV2Caller: ResolverV2Caller{contract: contract}, ResolverV2Transactor: ResolverV2Transactor{contract: contract}, ResolverV2Filterer: ResolverV2Filterer{contract: contract}}, nil
}

// NewResolverV2Caller creates a new read-only instance of ResolverV2, bound to a specific deployed contract.
func NewResolverV2Caller(address common.Address, caller bind.ContractCaller) (*ResolverV2Caller, error) {
	contract, err := bindResolverV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ResolverV2Caller{contract: contract}, nil
}

// NewResolverV2Transactor creates a new write-only instance of ResolverV2, bound to a specific deployed contract.
func NewResolverV2Transactor(address common.Address, transactor bind.ContractTransactor) (*ResolverV2Transactor, error) {
	contract, err := bindResolverV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ResolverV2Transactor{contract: contract}, nil
}

// NewResolverV2Filterer creates a new log filterer instance of ResolverV2, bound to a specific deployed contract.
func NewResolverV2Filterer(address common.Address, filterer bind.ContractFilterer) (*ResolverV2Filterer, error) {
	contract, err := bindResolverV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ResolverV2Filterer{contract: contract}, nil
}

// bindResolverV2 binds a generic wrapper to an already deployed contract.
func bindResolverV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ResolverV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ResolverV2 *ResolverV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ResolverV2.Contract.ResolverV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ResolverV2 *ResolverV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ResolverV2.Contract.ResolverV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ResolverV2 *ResolverV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ResolverV2.Contract.ResolverV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ResolverV2 *ResolverV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ResolverV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ResolverV2 *ResolverV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ResolverV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ResolverV2 *ResolverV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ResolverV2.Contract.contract.Transact(opts, method, params...)
}

// ABI is a free data retrieval call binding the contract method 0x2203ab56.
//
// Solidity: function ABI(bytes32 node, uint256 contentTypes) view returns(uint256, bytes)
func (_ResolverV2 *ResolverV2Caller) ABI(opts *bind.CallOpts, node [32]byte, contentTypes *big.Int) (*big.Int, []byte, error) {
	var out []interface{}
	err := _ResolverV2.contract.Call(opts, &out, "ABI", node, contentTypes)

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
func (_ResolverV2 *ResolverV2Session) ABI(node [32]byte, contentTypes *big.Int) (*big.Int, []byte, error) {
	return _ResolverV2.Contract.ABI(&_ResolverV2.CallOpts, node, contentTypes)
}

// ABI is a free data retrieval call binding the contract method 0x2203ab56.
//
// Solidity: function ABI(bytes32 node, uint256 contentTypes) view returns(uint256, bytes)
func (_ResolverV2 *ResolverV2CallerSession) ABI(node [32]byte, contentTypes *big.Int) (*big.Int, []byte, error) {
	return _ResolverV2.Contract.ABI(&_ResolverV2.CallOpts, node, contentTypes)
}

// Addr is a free data retrieval call binding the contract method 0x3b3b57de.
//
// Solidity: function addr(bytes32 node) view returns(address)
func (_ResolverV2 *ResolverV2Caller) Addr(opts *bind.CallOpts, node [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ResolverV2.contract.Call(opts, &out, "addr", node)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Addr is a free data retrieval call binding the contract method 0x3b3b57de.
//
// Solidity: function addr(bytes32 node) view returns(address)
func (_ResolverV2 *ResolverV2Session) Addr(node [32]byte) (common.Address, error) {
	return _ResolverV2.Contract.Addr(&_ResolverV2.CallOpts, node)
}

// Addr is a free data retrieval call binding the contract method 0x3b3b57de.
//
// Solidity: function addr(bytes32 node) view returns(address)
func (_ResolverV2 *ResolverV2CallerSession) Addr(node [32]byte) (common.Address, error) {
	return _ResolverV2.Contract.Addr(&_ResolverV2.CallOpts, node)
}

// Addr0 is a free data retrieval call binding the contract method 0xf1cb7e06.
//
// Solidity: function addr(bytes32 node, uint256 coinType) view returns(bytes)
func (_ResolverV2 *ResolverV2Caller) Addr0(opts *bind.CallOpts, node [32]byte, coinType *big.Int) ([]byte, error) {
	var out []interface{}
	err := _ResolverV2.contract.Call(opts, &out, "addr0", node, coinType)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Addr0 is a free data retrieval call binding the contract method 0xf1cb7e06.
//
// Solidity: function addr(bytes32 node, uint256 coinType) view returns(bytes)
func (_ResolverV2 *ResolverV2Session) Addr0(node [32]byte, coinType *big.Int) ([]byte, error) {
	return _ResolverV2.Contract.Addr0(&_ResolverV2.CallOpts, node, coinType)
}

// Addr0 is a free data retrieval call binding the contract method 0xf1cb7e06.
//
// Solidity: function addr(bytes32 node, uint256 coinType) view returns(bytes)
func (_ResolverV2 *ResolverV2CallerSession) Addr0(node [32]byte, coinType *big.Int) ([]byte, error) {
	return _ResolverV2.Contract.Addr0(&_ResolverV2.CallOpts, node, coinType)
}

// Contenthash is a free data retrieval call binding the contract method 0xbc1c58d1.
//
// Solidity: function contenthash(bytes32 node) view returns(bytes)
func (_ResolverV2 *ResolverV2Caller) Contenthash(opts *bind.CallOpts, node [32]byte) ([]byte, error) {
	var out []interface{}
	err := _ResolverV2.contract.Call(opts, &out, "contenthash", node)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Contenthash is a free data retrieval call binding the contract method 0xbc1c58d1.
//
// Solidity: function contenthash(bytes32 node) view returns(bytes)
func (_ResolverV2 *ResolverV2Session) Contenthash(node [32]byte) ([]byte, error) {
	return _ResolverV2.Contract.Contenthash(&_ResolverV2.CallOpts, node)
}

// Contenthash is a free data retrieval call binding the contract method 0xbc1c58d1.
//
// Solidity: function contenthash(bytes32 node) view returns(bytes)
func (_ResolverV2 *ResolverV2CallerSession) Contenthash(node [32]byte) ([]byte, error) {
	return _ResolverV2.Contract.Contenthash(&_ResolverV2.CallOpts, node)
}

// DnsRecord is a free data retrieval call binding the contract method 0xa8fa5682.
//
// Solidity: function dnsRecord(bytes32 node, bytes32 name, uint16 resource) view returns(bytes)
func (_ResolverV2 *ResolverV2Caller) DnsRecord(opts *bind.CallOpts, node [32]byte, name [32]byte, resource uint16) ([]byte, error) {
	var out []interface{}
	err := _ResolverV2.contract.Call(opts, &out, "dnsRecord", node, name, resource)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// DnsRecord is a free data retrieval call binding the contract method 0xa8fa5682.
//
// Solidity: function dnsRecord(bytes32 node, bytes32 name, uint16 resource) view returns(bytes)
func (_ResolverV2 *ResolverV2Session) DnsRecord(node [32]byte, name [32]byte, resource uint16) ([]byte, error) {
	return _ResolverV2.Contract.DnsRecord(&_ResolverV2.CallOpts, node, name, resource)
}

// DnsRecord is a free data retrieval call binding the contract method 0xa8fa5682.
//
// Solidity: function dnsRecord(bytes32 node, bytes32 name, uint16 resource) view returns(bytes)
func (_ResolverV2 *ResolverV2CallerSession) DnsRecord(node [32]byte, name [32]byte, resource uint16) ([]byte, error) {
	return _ResolverV2.Contract.DnsRecord(&_ResolverV2.CallOpts, node, name, resource)
}

// HasDNSRecords is a free data retrieval call binding the contract method 0x4cbf6ba4.
//
// Solidity: function hasDNSRecords(bytes32 node, bytes32 name) view returns(bool)
func (_ResolverV2 *ResolverV2Caller) HasDNSRecords(opts *bind.CallOpts, node [32]byte, name [32]byte) (bool, error) {
	var out []interface{}
	err := _ResolverV2.contract.Call(opts, &out, "hasDNSRecords", node, name)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasDNSRecords is a free data retrieval call binding the contract method 0x4cbf6ba4.
//
// Solidity: function hasDNSRecords(bytes32 node, bytes32 name) view returns(bool)
func (_ResolverV2 *ResolverV2Session) HasDNSRecords(node [32]byte, name [32]byte) (bool, error) {
	return _ResolverV2.Contract.HasDNSRecords(&_ResolverV2.CallOpts, node, name)
}

// HasDNSRecords is a free data retrieval call binding the contract method 0x4cbf6ba4.
//
// Solidity: function hasDNSRecords(bytes32 node, bytes32 name) view returns(bool)
func (_ResolverV2 *ResolverV2CallerSession) HasDNSRecords(node [32]byte, name [32]byte) (bool, error) {
	return _ResolverV2.Contract.HasDNSRecords(&_ResolverV2.CallOpts, node, name)
}

// InterfaceImplementer is a free data retrieval call binding the contract method 0x124a319c.
//
// Solidity: function interfaceImplementer(bytes32 node, bytes4 interfaceID) view returns(address)
func (_ResolverV2 *ResolverV2Caller) InterfaceImplementer(opts *bind.CallOpts, node [32]byte, interfaceID [4]byte) (common.Address, error) {
	var out []interface{}
	err := _ResolverV2.contract.Call(opts, &out, "interfaceImplementer", node, interfaceID)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InterfaceImplementer is a free data retrieval call binding the contract method 0x124a319c.
//
// Solidity: function interfaceImplementer(bytes32 node, bytes4 interfaceID) view returns(address)
func (_ResolverV2 *ResolverV2Session) InterfaceImplementer(node [32]byte, interfaceID [4]byte) (common.Address, error) {
	return _ResolverV2.Contract.InterfaceImplementer(&_ResolverV2.CallOpts, node, interfaceID)
}

// InterfaceImplementer is a free data retrieval call binding the contract method 0x124a319c.
//
// Solidity: function interfaceImplementer(bytes32 node, bytes4 interfaceID) view returns(address)
func (_ResolverV2 *ResolverV2CallerSession) InterfaceImplementer(node [32]byte, interfaceID [4]byte) (common.Address, error) {
	return _ResolverV2.Contract.InterfaceImplementer(&_ResolverV2.CallOpts, node, interfaceID)
}

// IsApprovedFor is a free data retrieval call binding the contract method 0xa9784b3e.
//
// Solidity: function isApprovedFor(address owner, bytes32 node, address delegate) view returns(bool)
func (_ResolverV2 *ResolverV2Caller) IsApprovedFor(opts *bind.CallOpts, owner common.Address, node [32]byte, delegate common.Address) (bool, error) {
	var out []interface{}
	err := _ResolverV2.contract.Call(opts, &out, "isApprovedFor", owner, node, delegate)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedFor is a free data retrieval call binding the contract method 0xa9784b3e.
//
// Solidity: function isApprovedFor(address owner, bytes32 node, address delegate) view returns(bool)
func (_ResolverV2 *ResolverV2Session) IsApprovedFor(owner common.Address, node [32]byte, delegate common.Address) (bool, error) {
	return _ResolverV2.Contract.IsApprovedFor(&_ResolverV2.CallOpts, owner, node, delegate)
}

// IsApprovedFor is a free data retrieval call binding the contract method 0xa9784b3e.
//
// Solidity: function isApprovedFor(address owner, bytes32 node, address delegate) view returns(bool)
func (_ResolverV2 *ResolverV2CallerSession) IsApprovedFor(owner common.Address, node [32]byte, delegate common.Address) (bool, error) {
	return _ResolverV2.Contract.IsApprovedFor(&_ResolverV2.CallOpts, owner, node, delegate)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_ResolverV2 *ResolverV2Caller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _ResolverV2.contract.Call(opts, &out, "isApprovedForAll", account, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_ResolverV2 *ResolverV2Session) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _ResolverV2.Contract.IsApprovedForAll(&_ResolverV2.CallOpts, account, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_ResolverV2 *ResolverV2CallerSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _ResolverV2.Contract.IsApprovedForAll(&_ResolverV2.CallOpts, account, operator)
}

// Name is a free data retrieval call binding the contract method 0x691f3431.
//
// Solidity: function name(bytes32 node) view returns(string)
func (_ResolverV2 *ResolverV2Caller) Name(opts *bind.CallOpts, node [32]byte) (string, error) {
	var out []interface{}
	err := _ResolverV2.contract.Call(opts, &out, "name", node)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x691f3431.
//
// Solidity: function name(bytes32 node) view returns(string)
func (_ResolverV2 *ResolverV2Session) Name(node [32]byte) (string, error) {
	return _ResolverV2.Contract.Name(&_ResolverV2.CallOpts, node)
}

// Name is a free data retrieval call binding the contract method 0x691f3431.
//
// Solidity: function name(bytes32 node) view returns(string)
func (_ResolverV2 *ResolverV2CallerSession) Name(node [32]byte) (string, error) {
	return _ResolverV2.Contract.Name(&_ResolverV2.CallOpts, node)
}

// Pubkey is a free data retrieval call binding the contract method 0xc8690233.
//
// Solidity: function pubkey(bytes32 node) view returns(bytes32 x, bytes32 y)
func (_ResolverV2 *ResolverV2Caller) Pubkey(opts *bind.CallOpts, node [32]byte) (struct {
	X [32]byte
	Y [32]byte
}, error) {
	var out []interface{}
	err := _ResolverV2.contract.Call(opts, &out, "pubkey", node)

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
func (_ResolverV2 *ResolverV2Session) Pubkey(node [32]byte) (struct {
	X [32]byte
	Y [32]byte
}, error) {
	return _ResolverV2.Contract.Pubkey(&_ResolverV2.CallOpts, node)
}

// Pubkey is a free data retrieval call binding the contract method 0xc8690233.
//
// Solidity: function pubkey(bytes32 node) view returns(bytes32 x, bytes32 y)
func (_ResolverV2 *ResolverV2CallerSession) Pubkey(node [32]byte) (struct {
	X [32]byte
	Y [32]byte
}, error) {
	return _ResolverV2.Contract.Pubkey(&_ResolverV2.CallOpts, node)
}

// RecordVersions is a free data retrieval call binding the contract method 0xd700ff33.
//
// Solidity: function recordVersions(bytes32 ) view returns(uint64)
func (_ResolverV2 *ResolverV2Caller) RecordVersions(opts *bind.CallOpts, arg0 [32]byte) (uint64, error) {
	var out []interface{}
	err := _ResolverV2.contract.Call(opts, &out, "recordVersions", arg0)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// RecordVersions is a free data retrieval call binding the contract method 0xd700ff33.
//
// Solidity: function recordVersions(bytes32 ) view returns(uint64)
func (_ResolverV2 *ResolverV2Session) RecordVersions(arg0 [32]byte) (uint64, error) {
	return _ResolverV2.Contract.RecordVersions(&_ResolverV2.CallOpts, arg0)
}

// RecordVersions is a free data retrieval call binding the contract method 0xd700ff33.
//
// Solidity: function recordVersions(bytes32 ) view returns(uint64)
func (_ResolverV2 *ResolverV2CallerSession) RecordVersions(arg0 [32]byte) (uint64, error) {
	return _ResolverV2.Contract.RecordVersions(&_ResolverV2.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) view returns(bool)
func (_ResolverV2 *ResolverV2Caller) SupportsInterface(opts *bind.CallOpts, interfaceID [4]byte) (bool, error) {
	var out []interface{}
	err := _ResolverV2.contract.Call(opts, &out, "supportsInterface", interfaceID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) view returns(bool)
func (_ResolverV2 *ResolverV2Session) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _ResolverV2.Contract.SupportsInterface(&_ResolverV2.CallOpts, interfaceID)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) view returns(bool)
func (_ResolverV2 *ResolverV2CallerSession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _ResolverV2.Contract.SupportsInterface(&_ResolverV2.CallOpts, interfaceID)
}

// Text is a free data retrieval call binding the contract method 0x59d1d43c.
//
// Solidity: function text(bytes32 node, string key) view returns(string)
func (_ResolverV2 *ResolverV2Caller) Text(opts *bind.CallOpts, node [32]byte, key string) (string, error) {
	var out []interface{}
	err := _ResolverV2.contract.Call(opts, &out, "text", node, key)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Text is a free data retrieval call binding the contract method 0x59d1d43c.
//
// Solidity: function text(bytes32 node, string key) view returns(string)
func (_ResolverV2 *ResolverV2Session) Text(node [32]byte, key string) (string, error) {
	return _ResolverV2.Contract.Text(&_ResolverV2.CallOpts, node, key)
}

// Text is a free data retrieval call binding the contract method 0x59d1d43c.
//
// Solidity: function text(bytes32 node, string key) view returns(string)
func (_ResolverV2 *ResolverV2CallerSession) Text(node [32]byte, key string) (string, error) {
	return _ResolverV2.Contract.Text(&_ResolverV2.CallOpts, node, key)
}

// Zonehash is a free data retrieval call binding the contract method 0x5c98042b.
//
// Solidity: function zonehash(bytes32 node) view returns(bytes)
func (_ResolverV2 *ResolverV2Caller) Zonehash(opts *bind.CallOpts, node [32]byte) ([]byte, error) {
	var out []interface{}
	err := _ResolverV2.contract.Call(opts, &out, "zonehash", node)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Zonehash is a free data retrieval call binding the contract method 0x5c98042b.
//
// Solidity: function zonehash(bytes32 node) view returns(bytes)
func (_ResolverV2 *ResolverV2Session) Zonehash(node [32]byte) ([]byte, error) {
	return _ResolverV2.Contract.Zonehash(&_ResolverV2.CallOpts, node)
}

// Zonehash is a free data retrieval call binding the contract method 0x5c98042b.
//
// Solidity: function zonehash(bytes32 node) view returns(bytes)
func (_ResolverV2 *ResolverV2CallerSession) Zonehash(node [32]byte) ([]byte, error) {
	return _ResolverV2.Contract.Zonehash(&_ResolverV2.CallOpts, node)
}

// Approve is a paid mutator transaction binding the contract method 0xa4b91a01.
//
// Solidity: function approve(bytes32 node, address delegate, bool approved) returns()
func (_ResolverV2 *ResolverV2Transactor) Approve(opts *bind.TransactOpts, node [32]byte, delegate common.Address, approved bool) (*types.Transaction, error) {
	return _ResolverV2.contract.Transact(opts, "approve", node, delegate, approved)
}

// Approve is a paid mutator transaction binding the contract method 0xa4b91a01.
//
// Solidity: function approve(bytes32 node, address delegate, bool approved) returns()
func (_ResolverV2 *ResolverV2Session) Approve(node [32]byte, delegate common.Address, approved bool) (*types.Transaction, error) {
	return _ResolverV2.Contract.Approve(&_ResolverV2.TransactOpts, node, delegate, approved)
}

// Approve is a paid mutator transaction binding the contract method 0xa4b91a01.
//
// Solidity: function approve(bytes32 node, address delegate, bool approved) returns()
func (_ResolverV2 *ResolverV2TransactorSession) Approve(node [32]byte, delegate common.Address, approved bool) (*types.Transaction, error) {
	return _ResolverV2.Contract.Approve(&_ResolverV2.TransactOpts, node, delegate, approved)
}

// ClearRecords is a paid mutator transaction binding the contract method 0x3603d758.
//
// Solidity: function clearRecords(bytes32 node) returns()
func (_ResolverV2 *ResolverV2Transactor) ClearRecords(opts *bind.TransactOpts, node [32]byte) (*types.Transaction, error) {
	return _ResolverV2.contract.Transact(opts, "clearRecords", node)
}

// ClearRecords is a paid mutator transaction binding the contract method 0x3603d758.
//
// Solidity: function clearRecords(bytes32 node) returns()
func (_ResolverV2 *ResolverV2Session) ClearRecords(node [32]byte) (*types.Transaction, error) {
	return _ResolverV2.Contract.ClearRecords(&_ResolverV2.TransactOpts, node)
}

// ClearRecords is a paid mutator transaction binding the contract method 0x3603d758.
//
// Solidity: function clearRecords(bytes32 node) returns()
func (_ResolverV2 *ResolverV2TransactorSession) ClearRecords(node [32]byte) (*types.Transaction, error) {
	return _ResolverV2.Contract.ClearRecords(&_ResolverV2.TransactOpts, node)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_ResolverV2 *ResolverV2Transactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _ResolverV2.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_ResolverV2 *ResolverV2Session) Multicall(data [][]byte) (*types.Transaction, error) {
	return _ResolverV2.Contract.Multicall(&_ResolverV2.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_ResolverV2 *ResolverV2TransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _ResolverV2.Contract.Multicall(&_ResolverV2.TransactOpts, data)
}

// MulticallWithNodeCheck is a paid mutator transaction binding the contract method 0xe32954eb.
//
// Solidity: function multicallWithNodeCheck(bytes32 nodehash, bytes[] data) returns(bytes[] results)
func (_ResolverV2 *ResolverV2Transactor) MulticallWithNodeCheck(opts *bind.TransactOpts, nodehash [32]byte, data [][]byte) (*types.Transaction, error) {
	return _ResolverV2.contract.Transact(opts, "multicallWithNodeCheck", nodehash, data)
}

// MulticallWithNodeCheck is a paid mutator transaction binding the contract method 0xe32954eb.
//
// Solidity: function multicallWithNodeCheck(bytes32 nodehash, bytes[] data) returns(bytes[] results)
func (_ResolverV2 *ResolverV2Session) MulticallWithNodeCheck(nodehash [32]byte, data [][]byte) (*types.Transaction, error) {
	return _ResolverV2.Contract.MulticallWithNodeCheck(&_ResolverV2.TransactOpts, nodehash, data)
}

// MulticallWithNodeCheck is a paid mutator transaction binding the contract method 0xe32954eb.
//
// Solidity: function multicallWithNodeCheck(bytes32 nodehash, bytes[] data) returns(bytes[] results)
func (_ResolverV2 *ResolverV2TransactorSession) MulticallWithNodeCheck(nodehash [32]byte, data [][]byte) (*types.Transaction, error) {
	return _ResolverV2.Contract.MulticallWithNodeCheck(&_ResolverV2.TransactOpts, nodehash, data)
}

// SetABI is a paid mutator transaction binding the contract method 0x623195b0.
//
// Solidity: function setABI(bytes32 node, uint256 contentType, bytes data) returns()
func (_ResolverV2 *ResolverV2Transactor) SetABI(opts *bind.TransactOpts, node [32]byte, contentType *big.Int, data []byte) (*types.Transaction, error) {
	return _ResolverV2.contract.Transact(opts, "setABI", node, contentType, data)
}

// SetABI is a paid mutator transaction binding the contract method 0x623195b0.
//
// Solidity: function setABI(bytes32 node, uint256 contentType, bytes data) returns()
func (_ResolverV2 *ResolverV2Session) SetABI(node [32]byte, contentType *big.Int, data []byte) (*types.Transaction, error) {
	return _ResolverV2.Contract.SetABI(&_ResolverV2.TransactOpts, node, contentType, data)
}

// SetABI is a paid mutator transaction binding the contract method 0x623195b0.
//
// Solidity: function setABI(bytes32 node, uint256 contentType, bytes data) returns()
func (_ResolverV2 *ResolverV2TransactorSession) SetABI(node [32]byte, contentType *big.Int, data []byte) (*types.Transaction, error) {
	return _ResolverV2.Contract.SetABI(&_ResolverV2.TransactOpts, node, contentType, data)
}

// SetAddr is a paid mutator transaction binding the contract method 0x8b95dd71.
//
// Solidity: function setAddr(bytes32 node, uint256 coinType, bytes a) returns()
func (_ResolverV2 *ResolverV2Transactor) SetAddr(opts *bind.TransactOpts, node [32]byte, coinType *big.Int, a []byte) (*types.Transaction, error) {
	return _ResolverV2.contract.Transact(opts, "setAddr", node, coinType, a)
}

// SetAddr is a paid mutator transaction binding the contract method 0x8b95dd71.
//
// Solidity: function setAddr(bytes32 node, uint256 coinType, bytes a) returns()
func (_ResolverV2 *ResolverV2Session) SetAddr(node [32]byte, coinType *big.Int, a []byte) (*types.Transaction, error) {
	return _ResolverV2.Contract.SetAddr(&_ResolverV2.TransactOpts, node, coinType, a)
}

// SetAddr is a paid mutator transaction binding the contract method 0x8b95dd71.
//
// Solidity: function setAddr(bytes32 node, uint256 coinType, bytes a) returns()
func (_ResolverV2 *ResolverV2TransactorSession) SetAddr(node [32]byte, coinType *big.Int, a []byte) (*types.Transaction, error) {
	return _ResolverV2.Contract.SetAddr(&_ResolverV2.TransactOpts, node, coinType, a)
}

// SetAddr0 is a paid mutator transaction binding the contract method 0xd5fa2b00.
//
// Solidity: function setAddr(bytes32 node, address a) returns()
func (_ResolverV2 *ResolverV2Transactor) SetAddr0(opts *bind.TransactOpts, node [32]byte, a common.Address) (*types.Transaction, error) {
	return _ResolverV2.contract.Transact(opts, "setAddr0", node, a)
}

// SetAddr0 is a paid mutator transaction binding the contract method 0xd5fa2b00.
//
// Solidity: function setAddr(bytes32 node, address a) returns()
func (_ResolverV2 *ResolverV2Session) SetAddr0(node [32]byte, a common.Address) (*types.Transaction, error) {
	return _ResolverV2.Contract.SetAddr0(&_ResolverV2.TransactOpts, node, a)
}

// SetAddr0 is a paid mutator transaction binding the contract method 0xd5fa2b00.
//
// Solidity: function setAddr(bytes32 node, address a) returns()
func (_ResolverV2 *ResolverV2TransactorSession) SetAddr0(node [32]byte, a common.Address) (*types.Transaction, error) {
	return _ResolverV2.Contract.SetAddr0(&_ResolverV2.TransactOpts, node, a)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ResolverV2 *ResolverV2Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _ResolverV2.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ResolverV2 *ResolverV2Session) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ResolverV2.Contract.SetApprovalForAll(&_ResolverV2.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ResolverV2 *ResolverV2TransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ResolverV2.Contract.SetApprovalForAll(&_ResolverV2.TransactOpts, operator, approved)
}

// SetContenthash is a paid mutator transaction binding the contract method 0x304e6ade.
//
// Solidity: function setContenthash(bytes32 node, bytes hash) returns()
func (_ResolverV2 *ResolverV2Transactor) SetContenthash(opts *bind.TransactOpts, node [32]byte, hash []byte) (*types.Transaction, error) {
	return _ResolverV2.contract.Transact(opts, "setContenthash", node, hash)
}

// SetContenthash is a paid mutator transaction binding the contract method 0x304e6ade.
//
// Solidity: function setContenthash(bytes32 node, bytes hash) returns()
func (_ResolverV2 *ResolverV2Session) SetContenthash(node [32]byte, hash []byte) (*types.Transaction, error) {
	return _ResolverV2.Contract.SetContenthash(&_ResolverV2.TransactOpts, node, hash)
}

// SetContenthash is a paid mutator transaction binding the contract method 0x304e6ade.
//
// Solidity: function setContenthash(bytes32 node, bytes hash) returns()
func (_ResolverV2 *ResolverV2TransactorSession) SetContenthash(node [32]byte, hash []byte) (*types.Transaction, error) {
	return _ResolverV2.Contract.SetContenthash(&_ResolverV2.TransactOpts, node, hash)
}

// SetDNSRecords is a paid mutator transaction binding the contract method 0x0af179d7.
//
// Solidity: function setDNSRecords(bytes32 node, bytes data) returns()
func (_ResolverV2 *ResolverV2Transactor) SetDNSRecords(opts *bind.TransactOpts, node [32]byte, data []byte) (*types.Transaction, error) {
	return _ResolverV2.contract.Transact(opts, "setDNSRecords", node, data)
}

// SetDNSRecords is a paid mutator transaction binding the contract method 0x0af179d7.
//
// Solidity: function setDNSRecords(bytes32 node, bytes data) returns()
func (_ResolverV2 *ResolverV2Session) SetDNSRecords(node [32]byte, data []byte) (*types.Transaction, error) {
	return _ResolverV2.Contract.SetDNSRecords(&_ResolverV2.TransactOpts, node, data)
}

// SetDNSRecords is a paid mutator transaction binding the contract method 0x0af179d7.
//
// Solidity: function setDNSRecords(bytes32 node, bytes data) returns()
func (_ResolverV2 *ResolverV2TransactorSession) SetDNSRecords(node [32]byte, data []byte) (*types.Transaction, error) {
	return _ResolverV2.Contract.SetDNSRecords(&_ResolverV2.TransactOpts, node, data)
}

// SetInterface is a paid mutator transaction binding the contract method 0xe59d895d.
//
// Solidity: function setInterface(bytes32 node, bytes4 interfaceID, address implementer) returns()
func (_ResolverV2 *ResolverV2Transactor) SetInterface(opts *bind.TransactOpts, node [32]byte, interfaceID [4]byte, implementer common.Address) (*types.Transaction, error) {
	return _ResolverV2.contract.Transact(opts, "setInterface", node, interfaceID, implementer)
}

// SetInterface is a paid mutator transaction binding the contract method 0xe59d895d.
//
// Solidity: function setInterface(bytes32 node, bytes4 interfaceID, address implementer) returns()
func (_ResolverV2 *ResolverV2Session) SetInterface(node [32]byte, interfaceID [4]byte, implementer common.Address) (*types.Transaction, error) {
	return _ResolverV2.Contract.SetInterface(&_ResolverV2.TransactOpts, node, interfaceID, implementer)
}

// SetInterface is a paid mutator transaction binding the contract method 0xe59d895d.
//
// Solidity: function setInterface(bytes32 node, bytes4 interfaceID, address implementer) returns()
func (_ResolverV2 *ResolverV2TransactorSession) SetInterface(node [32]byte, interfaceID [4]byte, implementer common.Address) (*types.Transaction, error) {
	return _ResolverV2.Contract.SetInterface(&_ResolverV2.TransactOpts, node, interfaceID, implementer)
}

// SetName is a paid mutator transaction binding the contract method 0x77372213.
//
// Solidity: function setName(bytes32 node, string newName) returns()
func (_ResolverV2 *ResolverV2Transactor) SetName(opts *bind.TransactOpts, node [32]byte, newName string) (*types.Transaction, error) {
	return _ResolverV2.contract.Transact(opts, "setName", node, newName)
}

// SetName is a paid mutator transaction binding the contract method 0x77372213.
//
// Solidity: function setName(bytes32 node, string newName) returns()
func (_ResolverV2 *ResolverV2Session) SetName(node [32]byte, newName string) (*types.Transaction, error) {
	return _ResolverV2.Contract.SetName(&_ResolverV2.TransactOpts, node, newName)
}

// SetName is a paid mutator transaction binding the contract method 0x77372213.
//
// Solidity: function setName(bytes32 node, string newName) returns()
func (_ResolverV2 *ResolverV2TransactorSession) SetName(node [32]byte, newName string) (*types.Transaction, error) {
	return _ResolverV2.Contract.SetName(&_ResolverV2.TransactOpts, node, newName)
}

// SetPubkey is a paid mutator transaction binding the contract method 0x29cd62ea.
//
// Solidity: function setPubkey(bytes32 node, bytes32 x, bytes32 y) returns()
func (_ResolverV2 *ResolverV2Transactor) SetPubkey(opts *bind.TransactOpts, node [32]byte, x [32]byte, y [32]byte) (*types.Transaction, error) {
	return _ResolverV2.contract.Transact(opts, "setPubkey", node, x, y)
}

// SetPubkey is a paid mutator transaction binding the contract method 0x29cd62ea.
//
// Solidity: function setPubkey(bytes32 node, bytes32 x, bytes32 y) returns()
func (_ResolverV2 *ResolverV2Session) SetPubkey(node [32]byte, x [32]byte, y [32]byte) (*types.Transaction, error) {
	return _ResolverV2.Contract.SetPubkey(&_ResolverV2.TransactOpts, node, x, y)
}

// SetPubkey is a paid mutator transaction binding the contract method 0x29cd62ea.
//
// Solidity: function setPubkey(bytes32 node, bytes32 x, bytes32 y) returns()
func (_ResolverV2 *ResolverV2TransactorSession) SetPubkey(node [32]byte, x [32]byte, y [32]byte) (*types.Transaction, error) {
	return _ResolverV2.Contract.SetPubkey(&_ResolverV2.TransactOpts, node, x, y)
}

// SetText is a paid mutator transaction binding the contract method 0x10f13a8c.
//
// Solidity: function setText(bytes32 node, string key, string value) returns()
func (_ResolverV2 *ResolverV2Transactor) SetText(opts *bind.TransactOpts, node [32]byte, key string, value string) (*types.Transaction, error) {
	return _ResolverV2.contract.Transact(opts, "setText", node, key, value)
}

// SetText is a paid mutator transaction binding the contract method 0x10f13a8c.
//
// Solidity: function setText(bytes32 node, string key, string value) returns()
func (_ResolverV2 *ResolverV2Session) SetText(node [32]byte, key string, value string) (*types.Transaction, error) {
	return _ResolverV2.Contract.SetText(&_ResolverV2.TransactOpts, node, key, value)
}

// SetText is a paid mutator transaction binding the contract method 0x10f13a8c.
//
// Solidity: function setText(bytes32 node, string key, string value) returns()
func (_ResolverV2 *ResolverV2TransactorSession) SetText(node [32]byte, key string, value string) (*types.Transaction, error) {
	return _ResolverV2.Contract.SetText(&_ResolverV2.TransactOpts, node, key, value)
}

// SetZonehash is a paid mutator transaction binding the contract method 0xce3decdc.
//
// Solidity: function setZonehash(bytes32 node, bytes hash) returns()
func (_ResolverV2 *ResolverV2Transactor) SetZonehash(opts *bind.TransactOpts, node [32]byte, hash []byte) (*types.Transaction, error) {
	return _ResolverV2.contract.Transact(opts, "setZonehash", node, hash)
}

// SetZonehash is a paid mutator transaction binding the contract method 0xce3decdc.
//
// Solidity: function setZonehash(bytes32 node, bytes hash) returns()
func (_ResolverV2 *ResolverV2Session) SetZonehash(node [32]byte, hash []byte) (*types.Transaction, error) {
	return _ResolverV2.Contract.SetZonehash(&_ResolverV2.TransactOpts, node, hash)
}

// SetZonehash is a paid mutator transaction binding the contract method 0xce3decdc.
//
// Solidity: function setZonehash(bytes32 node, bytes hash) returns()
func (_ResolverV2 *ResolverV2TransactorSession) SetZonehash(node [32]byte, hash []byte) (*types.Transaction, error) {
	return _ResolverV2.Contract.SetZonehash(&_ResolverV2.TransactOpts, node, hash)
}

// ResolverV2ABIChangedIterator is returned from FilterABIChanged and is used to iterate over the raw logs and unpacked data for ABIChanged events raised by the ResolverV2 contract.
type ResolverV2ABIChangedIterator struct {
	Event *ResolverV2ABIChanged // Event containing the contract specifics and raw log

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
func (it *ResolverV2ABIChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResolverV2ABIChanged)
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
		it.Event = new(ResolverV2ABIChanged)
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
func (it *ResolverV2ABIChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResolverV2ABIChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResolverV2ABIChanged represents a ABIChanged event raised by the ResolverV2 contract.
type ResolverV2ABIChanged struct {
	Node        [32]byte
	ContentType *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterABIChanged is a free log retrieval operation binding the contract event 0xaa121bbeef5f32f5961a2a28966e769023910fc9479059ee3495d4c1a696efe3.
//
// Solidity: event ABIChanged(bytes32 indexed node, uint256 indexed contentType)
func (_ResolverV2 *ResolverV2Filterer) FilterABIChanged(opts *bind.FilterOpts, node [][32]byte, contentType []*big.Int) (*ResolverV2ABIChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var contentTypeRule []interface{}
	for _, contentTypeItem := range contentType {
		contentTypeRule = append(contentTypeRule, contentTypeItem)
	}

	logs, sub, err := _ResolverV2.contract.FilterLogs(opts, "ABIChanged", nodeRule, contentTypeRule)
	if err != nil {
		return nil, err
	}
	return &ResolverV2ABIChangedIterator{contract: _ResolverV2.contract, event: "ABIChanged", logs: logs, sub: sub}, nil
}

// WatchABIChanged is a free log subscription operation binding the contract event 0xaa121bbeef5f32f5961a2a28966e769023910fc9479059ee3495d4c1a696efe3.
//
// Solidity: event ABIChanged(bytes32 indexed node, uint256 indexed contentType)
func (_ResolverV2 *ResolverV2Filterer) WatchABIChanged(opts *bind.WatchOpts, sink chan<- *ResolverV2ABIChanged, node [][32]byte, contentType []*big.Int) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var contentTypeRule []interface{}
	for _, contentTypeItem := range contentType {
		contentTypeRule = append(contentTypeRule, contentTypeItem)
	}

	logs, sub, err := _ResolverV2.contract.WatchLogs(opts, "ABIChanged", nodeRule, contentTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResolverV2ABIChanged)
				if err := _ResolverV2.contract.UnpackLog(event, "ABIChanged", log); err != nil {
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
func (_ResolverV2 *ResolverV2Filterer) ParseABIChanged(log types.Log) (*ResolverV2ABIChanged, error) {
	event := new(ResolverV2ABIChanged)
	if err := _ResolverV2.contract.UnpackLog(event, "ABIChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ResolverV2AddrChangedIterator is returned from FilterAddrChanged and is used to iterate over the raw logs and unpacked data for AddrChanged events raised by the ResolverV2 contract.
type ResolverV2AddrChangedIterator struct {
	Event *ResolverV2AddrChanged // Event containing the contract specifics and raw log

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
func (it *ResolverV2AddrChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResolverV2AddrChanged)
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
		it.Event = new(ResolverV2AddrChanged)
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
func (it *ResolverV2AddrChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResolverV2AddrChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResolverV2AddrChanged represents a AddrChanged event raised by the ResolverV2 contract.
type ResolverV2AddrChanged struct {
	Node [32]byte
	A    common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAddrChanged is a free log retrieval operation binding the contract event 0x52d7d861f09ab3d26239d492e8968629f95e9e318cf0b73bfddc441522a15fd2.
//
// Solidity: event AddrChanged(bytes32 indexed node, address a)
func (_ResolverV2 *ResolverV2Filterer) FilterAddrChanged(opts *bind.FilterOpts, node [][32]byte) (*ResolverV2AddrChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ResolverV2.contract.FilterLogs(opts, "AddrChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &ResolverV2AddrChangedIterator{contract: _ResolverV2.contract, event: "AddrChanged", logs: logs, sub: sub}, nil
}

// WatchAddrChanged is a free log subscription operation binding the contract event 0x52d7d861f09ab3d26239d492e8968629f95e9e318cf0b73bfddc441522a15fd2.
//
// Solidity: event AddrChanged(bytes32 indexed node, address a)
func (_ResolverV2 *ResolverV2Filterer) WatchAddrChanged(opts *bind.WatchOpts, sink chan<- *ResolverV2AddrChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ResolverV2.contract.WatchLogs(opts, "AddrChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResolverV2AddrChanged)
				if err := _ResolverV2.contract.UnpackLog(event, "AddrChanged", log); err != nil {
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
func (_ResolverV2 *ResolverV2Filterer) ParseAddrChanged(log types.Log) (*ResolverV2AddrChanged, error) {
	event := new(ResolverV2AddrChanged)
	if err := _ResolverV2.contract.UnpackLog(event, "AddrChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ResolverV2AddressChangedIterator is returned from FilterAddressChanged and is used to iterate over the raw logs and unpacked data for AddressChanged events raised by the ResolverV2 contract.
type ResolverV2AddressChangedIterator struct {
	Event *ResolverV2AddressChanged // Event containing the contract specifics and raw log

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
func (it *ResolverV2AddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResolverV2AddressChanged)
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
		it.Event = new(ResolverV2AddressChanged)
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
func (it *ResolverV2AddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResolverV2AddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResolverV2AddressChanged represents a AddressChanged event raised by the ResolverV2 contract.
type ResolverV2AddressChanged struct {
	Node       [32]byte
	CoinType   *big.Int
	NewAddress []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAddressChanged is a free log retrieval operation binding the contract event 0x65412581168e88a1e60c6459d7f44ae83ad0832e670826c05a4e2476b57af752.
//
// Solidity: event AddressChanged(bytes32 indexed node, uint256 coinType, bytes newAddress)
func (_ResolverV2 *ResolverV2Filterer) FilterAddressChanged(opts *bind.FilterOpts, node [][32]byte) (*ResolverV2AddressChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ResolverV2.contract.FilterLogs(opts, "AddressChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &ResolverV2AddressChangedIterator{contract: _ResolverV2.contract, event: "AddressChanged", logs: logs, sub: sub}, nil
}

// WatchAddressChanged is a free log subscription operation binding the contract event 0x65412581168e88a1e60c6459d7f44ae83ad0832e670826c05a4e2476b57af752.
//
// Solidity: event AddressChanged(bytes32 indexed node, uint256 coinType, bytes newAddress)
func (_ResolverV2 *ResolverV2Filterer) WatchAddressChanged(opts *bind.WatchOpts, sink chan<- *ResolverV2AddressChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ResolverV2.contract.WatchLogs(opts, "AddressChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResolverV2AddressChanged)
				if err := _ResolverV2.contract.UnpackLog(event, "AddressChanged", log); err != nil {
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
func (_ResolverV2 *ResolverV2Filterer) ParseAddressChanged(log types.Log) (*ResolverV2AddressChanged, error) {
	event := new(ResolverV2AddressChanged)
	if err := _ResolverV2.contract.UnpackLog(event, "AddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ResolverV2ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ResolverV2 contract.
type ResolverV2ApprovalForAllIterator struct {
	Event *ResolverV2ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ResolverV2ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResolverV2ApprovalForAll)
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
		it.Event = new(ResolverV2ApprovalForAll)
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
func (it *ResolverV2ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResolverV2ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResolverV2ApprovalForAll represents a ApprovalForAll event raised by the ResolverV2 contract.
type ResolverV2ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ResolverV2 *ResolverV2Filterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ResolverV2ApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ResolverV2.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ResolverV2ApprovalForAllIterator{contract: _ResolverV2.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ResolverV2 *ResolverV2Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ResolverV2ApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ResolverV2.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResolverV2ApprovalForAll)
				if err := _ResolverV2.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_ResolverV2 *ResolverV2Filterer) ParseApprovalForAll(log types.Log) (*ResolverV2ApprovalForAll, error) {
	event := new(ResolverV2ApprovalForAll)
	if err := _ResolverV2.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ResolverV2ApprovedIterator is returned from FilterApproved and is used to iterate over the raw logs and unpacked data for Approved events raised by the ResolverV2 contract.
type ResolverV2ApprovedIterator struct {
	Event *ResolverV2Approved // Event containing the contract specifics and raw log

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
func (it *ResolverV2ApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResolverV2Approved)
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
		it.Event = new(ResolverV2Approved)
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
func (it *ResolverV2ApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResolverV2ApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResolverV2Approved represents a Approved event raised by the ResolverV2 contract.
type ResolverV2Approved struct {
	Owner    common.Address
	Node     [32]byte
	Delegate common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproved is a free log retrieval operation binding the contract event 0xf0ddb3b04746704017f9aa8bd728fcc2c1d11675041205350018915f5e4750a0.
//
// Solidity: event Approved(address owner, bytes32 indexed node, address indexed delegate, bool indexed approved)
func (_ResolverV2 *ResolverV2Filterer) FilterApproved(opts *bind.FilterOpts, node [][32]byte, delegate []common.Address, approved []bool) (*ResolverV2ApprovedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var delegateRule []interface{}
	for _, delegateItem := range delegate {
		delegateRule = append(delegateRule, delegateItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}

	logs, sub, err := _ResolverV2.contract.FilterLogs(opts, "Approved", nodeRule, delegateRule, approvedRule)
	if err != nil {
		return nil, err
	}
	return &ResolverV2ApprovedIterator{contract: _ResolverV2.contract, event: "Approved", logs: logs, sub: sub}, nil
}

// WatchApproved is a free log subscription operation binding the contract event 0xf0ddb3b04746704017f9aa8bd728fcc2c1d11675041205350018915f5e4750a0.
//
// Solidity: event Approved(address owner, bytes32 indexed node, address indexed delegate, bool indexed approved)
func (_ResolverV2 *ResolverV2Filterer) WatchApproved(opts *bind.WatchOpts, sink chan<- *ResolverV2Approved, node [][32]byte, delegate []common.Address, approved []bool) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var delegateRule []interface{}
	for _, delegateItem := range delegate {
		delegateRule = append(delegateRule, delegateItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}

	logs, sub, err := _ResolverV2.contract.WatchLogs(opts, "Approved", nodeRule, delegateRule, approvedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResolverV2Approved)
				if err := _ResolverV2.contract.UnpackLog(event, "Approved", log); err != nil {
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

// ParseApproved is a log parse operation binding the contract event 0xf0ddb3b04746704017f9aa8bd728fcc2c1d11675041205350018915f5e4750a0.
//
// Solidity: event Approved(address owner, bytes32 indexed node, address indexed delegate, bool indexed approved)
func (_ResolverV2 *ResolverV2Filterer) ParseApproved(log types.Log) (*ResolverV2Approved, error) {
	event := new(ResolverV2Approved)
	if err := _ResolverV2.contract.UnpackLog(event, "Approved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ResolverV2ContenthashChangedIterator is returned from FilterContenthashChanged and is used to iterate over the raw logs and unpacked data for ContenthashChanged events raised by the ResolverV2 contract.
type ResolverV2ContenthashChangedIterator struct {
	Event *ResolverV2ContenthashChanged // Event containing the contract specifics and raw log

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
func (it *ResolverV2ContenthashChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResolverV2ContenthashChanged)
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
		it.Event = new(ResolverV2ContenthashChanged)
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
func (it *ResolverV2ContenthashChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResolverV2ContenthashChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResolverV2ContenthashChanged represents a ContenthashChanged event raised by the ResolverV2 contract.
type ResolverV2ContenthashChanged struct {
	Node [32]byte
	Hash []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterContenthashChanged is a free log retrieval operation binding the contract event 0xe379c1624ed7e714cc0937528a32359d69d5281337765313dba4e081b72d7578.
//
// Solidity: event ContenthashChanged(bytes32 indexed node, bytes hash)
func (_ResolverV2 *ResolverV2Filterer) FilterContenthashChanged(opts *bind.FilterOpts, node [][32]byte) (*ResolverV2ContenthashChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ResolverV2.contract.FilterLogs(opts, "ContenthashChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &ResolverV2ContenthashChangedIterator{contract: _ResolverV2.contract, event: "ContenthashChanged", logs: logs, sub: sub}, nil
}

// WatchContenthashChanged is a free log subscription operation binding the contract event 0xe379c1624ed7e714cc0937528a32359d69d5281337765313dba4e081b72d7578.
//
// Solidity: event ContenthashChanged(bytes32 indexed node, bytes hash)
func (_ResolverV2 *ResolverV2Filterer) WatchContenthashChanged(opts *bind.WatchOpts, sink chan<- *ResolverV2ContenthashChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ResolverV2.contract.WatchLogs(opts, "ContenthashChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResolverV2ContenthashChanged)
				if err := _ResolverV2.contract.UnpackLog(event, "ContenthashChanged", log); err != nil {
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
func (_ResolverV2 *ResolverV2Filterer) ParseContenthashChanged(log types.Log) (*ResolverV2ContenthashChanged, error) {
	event := new(ResolverV2ContenthashChanged)
	if err := _ResolverV2.contract.UnpackLog(event, "ContenthashChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ResolverV2DNSRecordChangedIterator is returned from FilterDNSRecordChanged and is used to iterate over the raw logs and unpacked data for DNSRecordChanged events raised by the ResolverV2 contract.
type ResolverV2DNSRecordChangedIterator struct {
	Event *ResolverV2DNSRecordChanged // Event containing the contract specifics and raw log

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
func (it *ResolverV2DNSRecordChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResolverV2DNSRecordChanged)
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
		it.Event = new(ResolverV2DNSRecordChanged)
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
func (it *ResolverV2DNSRecordChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResolverV2DNSRecordChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResolverV2DNSRecordChanged represents a DNSRecordChanged event raised by the ResolverV2 contract.
type ResolverV2DNSRecordChanged struct {
	Node     [32]byte
	Name     []byte
	Resource uint16
	Record   []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDNSRecordChanged is a free log retrieval operation binding the contract event 0x52a608b3303a48862d07a73d82fa221318c0027fbbcfb1b2329bface3f19ff2b.
//
// Solidity: event DNSRecordChanged(bytes32 indexed node, bytes name, uint16 resource, bytes record)
func (_ResolverV2 *ResolverV2Filterer) FilterDNSRecordChanged(opts *bind.FilterOpts, node [][32]byte) (*ResolverV2DNSRecordChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ResolverV2.contract.FilterLogs(opts, "DNSRecordChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &ResolverV2DNSRecordChangedIterator{contract: _ResolverV2.contract, event: "DNSRecordChanged", logs: logs, sub: sub}, nil
}

// WatchDNSRecordChanged is a free log subscription operation binding the contract event 0x52a608b3303a48862d07a73d82fa221318c0027fbbcfb1b2329bface3f19ff2b.
//
// Solidity: event DNSRecordChanged(bytes32 indexed node, bytes name, uint16 resource, bytes record)
func (_ResolverV2 *ResolverV2Filterer) WatchDNSRecordChanged(opts *bind.WatchOpts, sink chan<- *ResolverV2DNSRecordChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ResolverV2.contract.WatchLogs(opts, "DNSRecordChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResolverV2DNSRecordChanged)
				if err := _ResolverV2.contract.UnpackLog(event, "DNSRecordChanged", log); err != nil {
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
func (_ResolverV2 *ResolverV2Filterer) ParseDNSRecordChanged(log types.Log) (*ResolverV2DNSRecordChanged, error) {
	event := new(ResolverV2DNSRecordChanged)
	if err := _ResolverV2.contract.UnpackLog(event, "DNSRecordChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ResolverV2DNSRecordDeletedIterator is returned from FilterDNSRecordDeleted and is used to iterate over the raw logs and unpacked data for DNSRecordDeleted events raised by the ResolverV2 contract.
type ResolverV2DNSRecordDeletedIterator struct {
	Event *ResolverV2DNSRecordDeleted // Event containing the contract specifics and raw log

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
func (it *ResolverV2DNSRecordDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResolverV2DNSRecordDeleted)
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
		it.Event = new(ResolverV2DNSRecordDeleted)
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
func (it *ResolverV2DNSRecordDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResolverV2DNSRecordDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResolverV2DNSRecordDeleted represents a DNSRecordDeleted event raised by the ResolverV2 contract.
type ResolverV2DNSRecordDeleted struct {
	Node     [32]byte
	Name     []byte
	Resource uint16
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDNSRecordDeleted is a free log retrieval operation binding the contract event 0x03528ed0c2a3ebc993b12ce3c16bb382f9c7d88ef7d8a1bf290eaf35955a1207.
//
// Solidity: event DNSRecordDeleted(bytes32 indexed node, bytes name, uint16 resource)
func (_ResolverV2 *ResolverV2Filterer) FilterDNSRecordDeleted(opts *bind.FilterOpts, node [][32]byte) (*ResolverV2DNSRecordDeletedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ResolverV2.contract.FilterLogs(opts, "DNSRecordDeleted", nodeRule)
	if err != nil {
		return nil, err
	}
	return &ResolverV2DNSRecordDeletedIterator{contract: _ResolverV2.contract, event: "DNSRecordDeleted", logs: logs, sub: sub}, nil
}

// WatchDNSRecordDeleted is a free log subscription operation binding the contract event 0x03528ed0c2a3ebc993b12ce3c16bb382f9c7d88ef7d8a1bf290eaf35955a1207.
//
// Solidity: event DNSRecordDeleted(bytes32 indexed node, bytes name, uint16 resource)
func (_ResolverV2 *ResolverV2Filterer) WatchDNSRecordDeleted(opts *bind.WatchOpts, sink chan<- *ResolverV2DNSRecordDeleted, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ResolverV2.contract.WatchLogs(opts, "DNSRecordDeleted", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResolverV2DNSRecordDeleted)
				if err := _ResolverV2.contract.UnpackLog(event, "DNSRecordDeleted", log); err != nil {
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
func (_ResolverV2 *ResolverV2Filterer) ParseDNSRecordDeleted(log types.Log) (*ResolverV2DNSRecordDeleted, error) {
	event := new(ResolverV2DNSRecordDeleted)
	if err := _ResolverV2.contract.UnpackLog(event, "DNSRecordDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ResolverV2DNSZonehashChangedIterator is returned from FilterDNSZonehashChanged and is used to iterate over the raw logs and unpacked data for DNSZonehashChanged events raised by the ResolverV2 contract.
type ResolverV2DNSZonehashChangedIterator struct {
	Event *ResolverV2DNSZonehashChanged // Event containing the contract specifics and raw log

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
func (it *ResolverV2DNSZonehashChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResolverV2DNSZonehashChanged)
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
		it.Event = new(ResolverV2DNSZonehashChanged)
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
func (it *ResolverV2DNSZonehashChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResolverV2DNSZonehashChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResolverV2DNSZonehashChanged represents a DNSZonehashChanged event raised by the ResolverV2 contract.
type ResolverV2DNSZonehashChanged struct {
	Node         [32]byte
	Lastzonehash []byte
	Zonehash     []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDNSZonehashChanged is a free log retrieval operation binding the contract event 0x8f15ed4b723ef428f250961da8315675b507046737e19319fc1a4d81bfe87f85.
//
// Solidity: event DNSZonehashChanged(bytes32 indexed node, bytes lastzonehash, bytes zonehash)
func (_ResolverV2 *ResolverV2Filterer) FilterDNSZonehashChanged(opts *bind.FilterOpts, node [][32]byte) (*ResolverV2DNSZonehashChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ResolverV2.contract.FilterLogs(opts, "DNSZonehashChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &ResolverV2DNSZonehashChangedIterator{contract: _ResolverV2.contract, event: "DNSZonehashChanged", logs: logs, sub: sub}, nil
}

// WatchDNSZonehashChanged is a free log subscription operation binding the contract event 0x8f15ed4b723ef428f250961da8315675b507046737e19319fc1a4d81bfe87f85.
//
// Solidity: event DNSZonehashChanged(bytes32 indexed node, bytes lastzonehash, bytes zonehash)
func (_ResolverV2 *ResolverV2Filterer) WatchDNSZonehashChanged(opts *bind.WatchOpts, sink chan<- *ResolverV2DNSZonehashChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ResolverV2.contract.WatchLogs(opts, "DNSZonehashChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResolverV2DNSZonehashChanged)
				if err := _ResolverV2.contract.UnpackLog(event, "DNSZonehashChanged", log); err != nil {
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

// ParseDNSZonehashChanged is a log parse operation binding the contract event 0x8f15ed4b723ef428f250961da8315675b507046737e19319fc1a4d81bfe87f85.
//
// Solidity: event DNSZonehashChanged(bytes32 indexed node, bytes lastzonehash, bytes zonehash)
func (_ResolverV2 *ResolverV2Filterer) ParseDNSZonehashChanged(log types.Log) (*ResolverV2DNSZonehashChanged, error) {
	event := new(ResolverV2DNSZonehashChanged)
	if err := _ResolverV2.contract.UnpackLog(event, "DNSZonehashChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ResolverV2InterfaceChangedIterator is returned from FilterInterfaceChanged and is used to iterate over the raw logs and unpacked data for InterfaceChanged events raised by the ResolverV2 contract.
type ResolverV2InterfaceChangedIterator struct {
	Event *ResolverV2InterfaceChanged // Event containing the contract specifics and raw log

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
func (it *ResolverV2InterfaceChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResolverV2InterfaceChanged)
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
		it.Event = new(ResolverV2InterfaceChanged)
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
func (it *ResolverV2InterfaceChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResolverV2InterfaceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResolverV2InterfaceChanged represents a InterfaceChanged event raised by the ResolverV2 contract.
type ResolverV2InterfaceChanged struct {
	Node        [32]byte
	InterfaceID [4]byte
	Implementer common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterInterfaceChanged is a free log retrieval operation binding the contract event 0x7c69f06bea0bdef565b709e93a147836b0063ba2dd89f02d0b7e8d931e6a6daa.
//
// Solidity: event InterfaceChanged(bytes32 indexed node, bytes4 indexed interfaceID, address implementer)
func (_ResolverV2 *ResolverV2Filterer) FilterInterfaceChanged(opts *bind.FilterOpts, node [][32]byte, interfaceID [][4]byte) (*ResolverV2InterfaceChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var interfaceIDRule []interface{}
	for _, interfaceIDItem := range interfaceID {
		interfaceIDRule = append(interfaceIDRule, interfaceIDItem)
	}

	logs, sub, err := _ResolverV2.contract.FilterLogs(opts, "InterfaceChanged", nodeRule, interfaceIDRule)
	if err != nil {
		return nil, err
	}
	return &ResolverV2InterfaceChangedIterator{contract: _ResolverV2.contract, event: "InterfaceChanged", logs: logs, sub: sub}, nil
}

// WatchInterfaceChanged is a free log subscription operation binding the contract event 0x7c69f06bea0bdef565b709e93a147836b0063ba2dd89f02d0b7e8d931e6a6daa.
//
// Solidity: event InterfaceChanged(bytes32 indexed node, bytes4 indexed interfaceID, address implementer)
func (_ResolverV2 *ResolverV2Filterer) WatchInterfaceChanged(opts *bind.WatchOpts, sink chan<- *ResolverV2InterfaceChanged, node [][32]byte, interfaceID [][4]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var interfaceIDRule []interface{}
	for _, interfaceIDItem := range interfaceID {
		interfaceIDRule = append(interfaceIDRule, interfaceIDItem)
	}

	logs, sub, err := _ResolverV2.contract.WatchLogs(opts, "InterfaceChanged", nodeRule, interfaceIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResolverV2InterfaceChanged)
				if err := _ResolverV2.contract.UnpackLog(event, "InterfaceChanged", log); err != nil {
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
func (_ResolverV2 *ResolverV2Filterer) ParseInterfaceChanged(log types.Log) (*ResolverV2InterfaceChanged, error) {
	event := new(ResolverV2InterfaceChanged)
	if err := _ResolverV2.contract.UnpackLog(event, "InterfaceChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ResolverV2NameChangedIterator is returned from FilterNameChanged and is used to iterate over the raw logs and unpacked data for NameChanged events raised by the ResolverV2 contract.
type ResolverV2NameChangedIterator struct {
	Event *ResolverV2NameChanged // Event containing the contract specifics and raw log

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
func (it *ResolverV2NameChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResolverV2NameChanged)
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
		it.Event = new(ResolverV2NameChanged)
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
func (it *ResolverV2NameChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResolverV2NameChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResolverV2NameChanged represents a NameChanged event raised by the ResolverV2 contract.
type ResolverV2NameChanged struct {
	Node [32]byte
	Name string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterNameChanged is a free log retrieval operation binding the contract event 0xb7d29e911041e8d9b843369e890bcb72c9388692ba48b65ac54e7214c4c348f7.
//
// Solidity: event NameChanged(bytes32 indexed node, string name)
func (_ResolverV2 *ResolverV2Filterer) FilterNameChanged(opts *bind.FilterOpts, node [][32]byte) (*ResolverV2NameChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ResolverV2.contract.FilterLogs(opts, "NameChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &ResolverV2NameChangedIterator{contract: _ResolverV2.contract, event: "NameChanged", logs: logs, sub: sub}, nil
}

// WatchNameChanged is a free log subscription operation binding the contract event 0xb7d29e911041e8d9b843369e890bcb72c9388692ba48b65ac54e7214c4c348f7.
//
// Solidity: event NameChanged(bytes32 indexed node, string name)
func (_ResolverV2 *ResolverV2Filterer) WatchNameChanged(opts *bind.WatchOpts, sink chan<- *ResolverV2NameChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ResolverV2.contract.WatchLogs(opts, "NameChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResolverV2NameChanged)
				if err := _ResolverV2.contract.UnpackLog(event, "NameChanged", log); err != nil {
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
func (_ResolverV2 *ResolverV2Filterer) ParseNameChanged(log types.Log) (*ResolverV2NameChanged, error) {
	event := new(ResolverV2NameChanged)
	if err := _ResolverV2.contract.UnpackLog(event, "NameChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ResolverV2PubkeyChangedIterator is returned from FilterPubkeyChanged and is used to iterate over the raw logs and unpacked data for PubkeyChanged events raised by the ResolverV2 contract.
type ResolverV2PubkeyChangedIterator struct {
	Event *ResolverV2PubkeyChanged // Event containing the contract specifics and raw log

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
func (it *ResolverV2PubkeyChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResolverV2PubkeyChanged)
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
		it.Event = new(ResolverV2PubkeyChanged)
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
func (it *ResolverV2PubkeyChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResolverV2PubkeyChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResolverV2PubkeyChanged represents a PubkeyChanged event raised by the ResolverV2 contract.
type ResolverV2PubkeyChanged struct {
	Node [32]byte
	X    [32]byte
	Y    [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPubkeyChanged is a free log retrieval operation binding the contract event 0x1d6f5e03d3f63eb58751986629a5439baee5079ff04f345becb66e23eb154e46.
//
// Solidity: event PubkeyChanged(bytes32 indexed node, bytes32 x, bytes32 y)
func (_ResolverV2 *ResolverV2Filterer) FilterPubkeyChanged(opts *bind.FilterOpts, node [][32]byte) (*ResolverV2PubkeyChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ResolverV2.contract.FilterLogs(opts, "PubkeyChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &ResolverV2PubkeyChangedIterator{contract: _ResolverV2.contract, event: "PubkeyChanged", logs: logs, sub: sub}, nil
}

// WatchPubkeyChanged is a free log subscription operation binding the contract event 0x1d6f5e03d3f63eb58751986629a5439baee5079ff04f345becb66e23eb154e46.
//
// Solidity: event PubkeyChanged(bytes32 indexed node, bytes32 x, bytes32 y)
func (_ResolverV2 *ResolverV2Filterer) WatchPubkeyChanged(opts *bind.WatchOpts, sink chan<- *ResolverV2PubkeyChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ResolverV2.contract.WatchLogs(opts, "PubkeyChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResolverV2PubkeyChanged)
				if err := _ResolverV2.contract.UnpackLog(event, "PubkeyChanged", log); err != nil {
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
func (_ResolverV2 *ResolverV2Filterer) ParsePubkeyChanged(log types.Log) (*ResolverV2PubkeyChanged, error) {
	event := new(ResolverV2PubkeyChanged)
	if err := _ResolverV2.contract.UnpackLog(event, "PubkeyChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ResolverV2TextChangedIterator is returned from FilterTextChanged and is used to iterate over the raw logs and unpacked data for TextChanged events raised by the ResolverV2 contract.
type ResolverV2TextChangedIterator struct {
	Event *ResolverV2TextChanged // Event containing the contract specifics and raw log

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
func (it *ResolverV2TextChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResolverV2TextChanged)
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
		it.Event = new(ResolverV2TextChanged)
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
func (it *ResolverV2TextChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResolverV2TextChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResolverV2TextChanged represents a TextChanged event raised by the ResolverV2 contract.
type ResolverV2TextChanged struct {
	Node       [32]byte
	IndexedKey common.Hash
	Key        string
	Value      string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTextChanged is a free log retrieval operation binding the contract event 0x448bc014f1536726cf8d54ff3d6481ed3cbc683c2591ca204274009afa09b1a1.
//
// Solidity: event TextChanged(bytes32 indexed node, string indexed indexedKey, string key, string value)
func (_ResolverV2 *ResolverV2Filterer) FilterTextChanged(opts *bind.FilterOpts, node [][32]byte, indexedKey []string) (*ResolverV2TextChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var indexedKeyRule []interface{}
	for _, indexedKeyItem := range indexedKey {
		indexedKeyRule = append(indexedKeyRule, indexedKeyItem)
	}

	logs, sub, err := _ResolverV2.contract.FilterLogs(opts, "TextChanged", nodeRule, indexedKeyRule)
	if err != nil {
		return nil, err
	}
	return &ResolverV2TextChangedIterator{contract: _ResolverV2.contract, event: "TextChanged", logs: logs, sub: sub}, nil
}

// WatchTextChanged is a free log subscription operation binding the contract event 0x448bc014f1536726cf8d54ff3d6481ed3cbc683c2591ca204274009afa09b1a1.
//
// Solidity: event TextChanged(bytes32 indexed node, string indexed indexedKey, string key, string value)
func (_ResolverV2 *ResolverV2Filterer) WatchTextChanged(opts *bind.WatchOpts, sink chan<- *ResolverV2TextChanged, node [][32]byte, indexedKey []string) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var indexedKeyRule []interface{}
	for _, indexedKeyItem := range indexedKey {
		indexedKeyRule = append(indexedKeyRule, indexedKeyItem)
	}

	logs, sub, err := _ResolverV2.contract.WatchLogs(opts, "TextChanged", nodeRule, indexedKeyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResolverV2TextChanged)
				if err := _ResolverV2.contract.UnpackLog(event, "TextChanged", log); err != nil {
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

// ParseTextChanged is a log parse operation binding the contract event 0x448bc014f1536726cf8d54ff3d6481ed3cbc683c2591ca204274009afa09b1a1.
//
// Solidity: event TextChanged(bytes32 indexed node, string indexed indexedKey, string key, string value)
func (_ResolverV2 *ResolverV2Filterer) ParseTextChanged(log types.Log) (*ResolverV2TextChanged, error) {
	event := new(ResolverV2TextChanged)
	if err := _ResolverV2.contract.UnpackLog(event, "TextChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ResolverV2VersionChangedIterator is returned from FilterVersionChanged and is used to iterate over the raw logs and unpacked data for VersionChanged events raised by the ResolverV2 contract.
type ResolverV2VersionChangedIterator struct {
	Event *ResolverV2VersionChanged // Event containing the contract specifics and raw log

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
func (it *ResolverV2VersionChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResolverV2VersionChanged)
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
		it.Event = new(ResolverV2VersionChanged)
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
func (it *ResolverV2VersionChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResolverV2VersionChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResolverV2VersionChanged represents a VersionChanged event raised by the ResolverV2 contract.
type ResolverV2VersionChanged struct {
	Node       [32]byte
	NewVersion uint64
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVersionChanged is a free log retrieval operation binding the contract event 0xc6621ccb8f3f5a04bb6502154b2caf6adf5983fe76dfef1cfc9c42e3579db444.
//
// Solidity: event VersionChanged(bytes32 indexed node, uint64 newVersion)
func (_ResolverV2 *ResolverV2Filterer) FilterVersionChanged(opts *bind.FilterOpts, node [][32]byte) (*ResolverV2VersionChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ResolverV2.contract.FilterLogs(opts, "VersionChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &ResolverV2VersionChangedIterator{contract: _ResolverV2.contract, event: "VersionChanged", logs: logs, sub: sub}, nil
}

// WatchVersionChanged is a free log subscription operation binding the contract event 0xc6621ccb8f3f5a04bb6502154b2caf6adf5983fe76dfef1cfc9c42e3579db444.
//
// Solidity: event VersionChanged(bytes32 indexed node, uint64 newVersion)
func (_ResolverV2 *ResolverV2Filterer) WatchVersionChanged(opts *bind.WatchOpts, sink chan<- *ResolverV2VersionChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ResolverV2.contract.WatchLogs(opts, "VersionChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResolverV2VersionChanged)
				if err := _ResolverV2.contract.UnpackLog(event, "VersionChanged", log); err != nil {
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

// ParseVersionChanged is a log parse operation binding the contract event 0xc6621ccb8f3f5a04bb6502154b2caf6adf5983fe76dfef1cfc9c42e3579db444.
//
// Solidity: event VersionChanged(bytes32 indexed node, uint64 newVersion)
func (_ResolverV2 *ResolverV2Filterer) ParseVersionChanged(log types.Log) (*ResolverV2VersionChanged, error) {
	event := new(ResolverV2VersionChanged)
	if err := _ResolverV2.contract.UnpackLog(event, "VersionChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
