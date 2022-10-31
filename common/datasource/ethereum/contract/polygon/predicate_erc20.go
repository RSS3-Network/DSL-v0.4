// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package polygon

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

// ERC20PredicateMetaData contains all meta data concerning the ERC20Predicate contract.
var ERC20PredicateMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"depositReceiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"rootToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LockedERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MANAGER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN_TYPE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANSFER_EVENT_SIG\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"withdrawer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rootToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"log\",\"type\":\"bytes\"}],\"name\":\"exitTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"depositReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rootToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"depositData\",\"type\":\"bytes\"}],\"name\":\"lockTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ERC20PredicateABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC20PredicateMetaData.ABI instead.
var ERC20PredicateABI = ERC20PredicateMetaData.ABI

// ERC20Predicate is an auto generated Go binding around an Ethereum contract.
type ERC20Predicate struct {
	ERC20PredicateCaller     // Read-only binding to the contract
	ERC20PredicateTransactor // Write-only binding to the contract
	ERC20PredicateFilterer   // Log filterer for contract events
}

// ERC20PredicateCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20PredicateCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20PredicateTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20PredicateTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20PredicateFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20PredicateFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20PredicateSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20PredicateSession struct {
	Contract     *ERC20Predicate   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20PredicateCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20PredicateCallerSession struct {
	Contract *ERC20PredicateCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ERC20PredicateTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20PredicateTransactorSession struct {
	Contract     *ERC20PredicateTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ERC20PredicateRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20PredicateRaw struct {
	Contract *ERC20Predicate // Generic contract binding to access the raw methods on
}

// ERC20PredicateCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20PredicateCallerRaw struct {
	Contract *ERC20PredicateCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20PredicateTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20PredicateTransactorRaw struct {
	Contract *ERC20PredicateTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Predicate creates a new instance of ERC20Predicate, bound to a specific deployed contract.
func NewERC20Predicate(address common.Address, backend bind.ContractBackend) (*ERC20Predicate, error) {
	contract, err := bindERC20Predicate(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Predicate{ERC20PredicateCaller: ERC20PredicateCaller{contract: contract}, ERC20PredicateTransactor: ERC20PredicateTransactor{contract: contract}, ERC20PredicateFilterer: ERC20PredicateFilterer{contract: contract}}, nil
}

// NewERC20PredicateCaller creates a new read-only instance of ERC20Predicate, bound to a specific deployed contract.
func NewERC20PredicateCaller(address common.Address, caller bind.ContractCaller) (*ERC20PredicateCaller, error) {
	contract, err := bindERC20Predicate(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20PredicateCaller{contract: contract}, nil
}

// NewERC20PredicateTransactor creates a new write-only instance of ERC20Predicate, bound to a specific deployed contract.
func NewERC20PredicateTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20PredicateTransactor, error) {
	contract, err := bindERC20Predicate(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20PredicateTransactor{contract: contract}, nil
}

// NewERC20PredicateFilterer creates a new log filterer instance of ERC20Predicate, bound to a specific deployed contract.
func NewERC20PredicateFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20PredicateFilterer, error) {
	contract, err := bindERC20Predicate(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20PredicateFilterer{contract: contract}, nil
}

// bindERC20Predicate binds a generic wrapper to an already deployed contract.
func bindERC20Predicate(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20PredicateABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Predicate *ERC20PredicateRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Predicate.Contract.ERC20PredicateCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Predicate *ERC20PredicateRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Predicate.Contract.ERC20PredicateTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Predicate *ERC20PredicateRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Predicate.Contract.ERC20PredicateTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Predicate *ERC20PredicateCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Predicate.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Predicate *ERC20PredicateTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Predicate.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Predicate *ERC20PredicateTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Predicate.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ERC20Predicate *ERC20PredicateCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC20Predicate.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ERC20Predicate *ERC20PredicateSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ERC20Predicate.Contract.DEFAULTADMINROLE(&_ERC20Predicate.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ERC20Predicate *ERC20PredicateCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ERC20Predicate.Contract.DEFAULTADMINROLE(&_ERC20Predicate.CallOpts)
}

// MANAGERROLE is a free data retrieval call binding the contract method 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (_ERC20Predicate *ERC20PredicateCaller) MANAGERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC20Predicate.contract.Call(opts, &out, "MANAGER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MANAGERROLE is a free data retrieval call binding the contract method 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (_ERC20Predicate *ERC20PredicateSession) MANAGERROLE() ([32]byte, error) {
	return _ERC20Predicate.Contract.MANAGERROLE(&_ERC20Predicate.CallOpts)
}

// MANAGERROLE is a free data retrieval call binding the contract method 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (_ERC20Predicate *ERC20PredicateCallerSession) MANAGERROLE() ([32]byte, error) {
	return _ERC20Predicate.Contract.MANAGERROLE(&_ERC20Predicate.CallOpts)
}

// TOKENTYPE is a free data retrieval call binding the contract method 0x609c92b8.
//
// Solidity: function TOKEN_TYPE() view returns(bytes32)
func (_ERC20Predicate *ERC20PredicateCaller) TOKENTYPE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC20Predicate.contract.Call(opts, &out, "TOKEN_TYPE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TOKENTYPE is a free data retrieval call binding the contract method 0x609c92b8.
//
// Solidity: function TOKEN_TYPE() view returns(bytes32)
func (_ERC20Predicate *ERC20PredicateSession) TOKENTYPE() ([32]byte, error) {
	return _ERC20Predicate.Contract.TOKENTYPE(&_ERC20Predicate.CallOpts)
}

// TOKENTYPE is a free data retrieval call binding the contract method 0x609c92b8.
//
// Solidity: function TOKEN_TYPE() view returns(bytes32)
func (_ERC20Predicate *ERC20PredicateCallerSession) TOKENTYPE() ([32]byte, error) {
	return _ERC20Predicate.Contract.TOKENTYPE(&_ERC20Predicate.CallOpts)
}

// TRANSFEREVENTSIG is a free data retrieval call binding the contract method 0xb017a30f.
//
// Solidity: function TRANSFER_EVENT_SIG() view returns(bytes32)
func (_ERC20Predicate *ERC20PredicateCaller) TRANSFEREVENTSIG(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC20Predicate.contract.Call(opts, &out, "TRANSFER_EVENT_SIG")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TRANSFEREVENTSIG is a free data retrieval call binding the contract method 0xb017a30f.
//
// Solidity: function TRANSFER_EVENT_SIG() view returns(bytes32)
func (_ERC20Predicate *ERC20PredicateSession) TRANSFEREVENTSIG() ([32]byte, error) {
	return _ERC20Predicate.Contract.TRANSFEREVENTSIG(&_ERC20Predicate.CallOpts)
}

// TRANSFEREVENTSIG is a free data retrieval call binding the contract method 0xb017a30f.
//
// Solidity: function TRANSFER_EVENT_SIG() view returns(bytes32)
func (_ERC20Predicate *ERC20PredicateCallerSession) TRANSFEREVENTSIG() ([32]byte, error) {
	return _ERC20Predicate.Contract.TRANSFEREVENTSIG(&_ERC20Predicate.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ERC20Predicate *ERC20PredicateCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ERC20Predicate.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ERC20Predicate *ERC20PredicateSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ERC20Predicate.Contract.GetRoleAdmin(&_ERC20Predicate.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ERC20Predicate *ERC20PredicateCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ERC20Predicate.Contract.GetRoleAdmin(&_ERC20Predicate.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_ERC20Predicate *ERC20PredicateCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ERC20Predicate.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_ERC20Predicate *ERC20PredicateSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _ERC20Predicate.Contract.GetRoleMember(&_ERC20Predicate.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_ERC20Predicate *ERC20PredicateCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _ERC20Predicate.Contract.GetRoleMember(&_ERC20Predicate.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_ERC20Predicate *ERC20PredicateCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Predicate.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_ERC20Predicate *ERC20PredicateSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _ERC20Predicate.Contract.GetRoleMemberCount(&_ERC20Predicate.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_ERC20Predicate *ERC20PredicateCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _ERC20Predicate.Contract.GetRoleMemberCount(&_ERC20Predicate.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ERC20Predicate *ERC20PredicateCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _ERC20Predicate.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ERC20Predicate *ERC20PredicateSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ERC20Predicate.Contract.HasRole(&_ERC20Predicate.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ERC20Predicate *ERC20PredicateCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ERC20Predicate.Contract.HasRole(&_ERC20Predicate.CallOpts, role, account)
}

// ExitTokens is a paid mutator transaction binding the contract method 0x8274664f.
//
// Solidity: function exitTokens(address withdrawer, address rootToken, bytes log) returns()
func (_ERC20Predicate *ERC20PredicateTransactor) ExitTokens(opts *bind.TransactOpts, withdrawer common.Address, rootToken common.Address, log []byte) (*types.Transaction, error) {
	return _ERC20Predicate.contract.Transact(opts, "exitTokens", withdrawer, rootToken, log)
}

// ExitTokens is a paid mutator transaction binding the contract method 0x8274664f.
//
// Solidity: function exitTokens(address withdrawer, address rootToken, bytes log) returns()
func (_ERC20Predicate *ERC20PredicateSession) ExitTokens(withdrawer common.Address, rootToken common.Address, log []byte) (*types.Transaction, error) {
	return _ERC20Predicate.Contract.ExitTokens(&_ERC20Predicate.TransactOpts, withdrawer, rootToken, log)
}

// ExitTokens is a paid mutator transaction binding the contract method 0x8274664f.
//
// Solidity: function exitTokens(address withdrawer, address rootToken, bytes log) returns()
func (_ERC20Predicate *ERC20PredicateTransactorSession) ExitTokens(withdrawer common.Address, rootToken common.Address, log []byte) (*types.Transaction, error) {
	return _ERC20Predicate.Contract.ExitTokens(&_ERC20Predicate.TransactOpts, withdrawer, rootToken, log)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ERC20Predicate *ERC20PredicateTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ERC20Predicate.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ERC20Predicate *ERC20PredicateSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ERC20Predicate.Contract.GrantRole(&_ERC20Predicate.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ERC20Predicate *ERC20PredicateTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ERC20Predicate.Contract.GrantRole(&_ERC20Predicate.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_ERC20Predicate *ERC20PredicateTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _ERC20Predicate.contract.Transact(opts, "initialize", _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_ERC20Predicate *ERC20PredicateSession) Initialize(_owner common.Address) (*types.Transaction, error) {
	return _ERC20Predicate.Contract.Initialize(&_ERC20Predicate.TransactOpts, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_ERC20Predicate *ERC20PredicateTransactorSession) Initialize(_owner common.Address) (*types.Transaction, error) {
	return _ERC20Predicate.Contract.Initialize(&_ERC20Predicate.TransactOpts, _owner)
}

// LockTokens is a paid mutator transaction binding the contract method 0xe375b64e.
//
// Solidity: function lockTokens(address depositor, address depositReceiver, address rootToken, bytes depositData) returns()
func (_ERC20Predicate *ERC20PredicateTransactor) LockTokens(opts *bind.TransactOpts, depositor common.Address, depositReceiver common.Address, rootToken common.Address, depositData []byte) (*types.Transaction, error) {
	return _ERC20Predicate.contract.Transact(opts, "lockTokens", depositor, depositReceiver, rootToken, depositData)
}

// LockTokens is a paid mutator transaction binding the contract method 0xe375b64e.
//
// Solidity: function lockTokens(address depositor, address depositReceiver, address rootToken, bytes depositData) returns()
func (_ERC20Predicate *ERC20PredicateSession) LockTokens(depositor common.Address, depositReceiver common.Address, rootToken common.Address, depositData []byte) (*types.Transaction, error) {
	return _ERC20Predicate.Contract.LockTokens(&_ERC20Predicate.TransactOpts, depositor, depositReceiver, rootToken, depositData)
}

// LockTokens is a paid mutator transaction binding the contract method 0xe375b64e.
//
// Solidity: function lockTokens(address depositor, address depositReceiver, address rootToken, bytes depositData) returns()
func (_ERC20Predicate *ERC20PredicateTransactorSession) LockTokens(depositor common.Address, depositReceiver common.Address, rootToken common.Address, depositData []byte) (*types.Transaction, error) {
	return _ERC20Predicate.Contract.LockTokens(&_ERC20Predicate.TransactOpts, depositor, depositReceiver, rootToken, depositData)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ERC20Predicate *ERC20PredicateTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ERC20Predicate.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ERC20Predicate *ERC20PredicateSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ERC20Predicate.Contract.RenounceRole(&_ERC20Predicate.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ERC20Predicate *ERC20PredicateTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ERC20Predicate.Contract.RenounceRole(&_ERC20Predicate.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ERC20Predicate *ERC20PredicateTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ERC20Predicate.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ERC20Predicate *ERC20PredicateSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ERC20Predicate.Contract.RevokeRole(&_ERC20Predicate.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ERC20Predicate *ERC20PredicateTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ERC20Predicate.Contract.RevokeRole(&_ERC20Predicate.TransactOpts, role, account)
}

// ERC20PredicateLockedERC20Iterator is returned from FilterLockedERC20 and is used to iterate over the raw logs and unpacked data for LockedERC20 events raised by the ERC20Predicate contract.
type ERC20PredicateLockedERC20Iterator struct {
	Event *ERC20PredicateLockedERC20 // Event containing the contract specifics and raw log

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
func (it *ERC20PredicateLockedERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20PredicateLockedERC20)
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
		it.Event = new(ERC20PredicateLockedERC20)
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
func (it *ERC20PredicateLockedERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20PredicateLockedERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20PredicateLockedERC20 represents a LockedERC20 event raised by the ERC20Predicate contract.
type ERC20PredicateLockedERC20 struct {
	Depositor       common.Address
	DepositReceiver common.Address
	RootToken       common.Address
	Amount          *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLockedERC20 is a free log retrieval operation binding the contract event 0x9b217a401a5ddf7c4d474074aff9958a18d48690d77cc2151c4706aa7348b401.
//
// Solidity: event LockedERC20(address indexed depositor, address indexed depositReceiver, address indexed rootToken, uint256 amount)
func (_ERC20Predicate *ERC20PredicateFilterer) FilterLockedERC20(opts *bind.FilterOpts, depositor []common.Address, depositReceiver []common.Address, rootToken []common.Address) (*ERC20PredicateLockedERC20Iterator, error) {

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}
	var depositReceiverRule []interface{}
	for _, depositReceiverItem := range depositReceiver {
		depositReceiverRule = append(depositReceiverRule, depositReceiverItem)
	}
	var rootTokenRule []interface{}
	for _, rootTokenItem := range rootToken {
		rootTokenRule = append(rootTokenRule, rootTokenItem)
	}

	logs, sub, err := _ERC20Predicate.contract.FilterLogs(opts, "LockedERC20", depositorRule, depositReceiverRule, rootTokenRule)
	if err != nil {
		return nil, err
	}
	return &ERC20PredicateLockedERC20Iterator{contract: _ERC20Predicate.contract, event: "LockedERC20", logs: logs, sub: sub}, nil
}

// WatchLockedERC20 is a free log subscription operation binding the contract event 0x9b217a401a5ddf7c4d474074aff9958a18d48690d77cc2151c4706aa7348b401.
//
// Solidity: event LockedERC20(address indexed depositor, address indexed depositReceiver, address indexed rootToken, uint256 amount)
func (_ERC20Predicate *ERC20PredicateFilterer) WatchLockedERC20(opts *bind.WatchOpts, sink chan<- *ERC20PredicateLockedERC20, depositor []common.Address, depositReceiver []common.Address, rootToken []common.Address) (event.Subscription, error) {

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}
	var depositReceiverRule []interface{}
	for _, depositReceiverItem := range depositReceiver {
		depositReceiverRule = append(depositReceiverRule, depositReceiverItem)
	}
	var rootTokenRule []interface{}
	for _, rootTokenItem := range rootToken {
		rootTokenRule = append(rootTokenRule, rootTokenItem)
	}

	logs, sub, err := _ERC20Predicate.contract.WatchLogs(opts, "LockedERC20", depositorRule, depositReceiverRule, rootTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20PredicateLockedERC20)
				if err := _ERC20Predicate.contract.UnpackLog(event, "LockedERC20", log); err != nil {
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

// ParseLockedERC20 is a log parse operation binding the contract event 0x9b217a401a5ddf7c4d474074aff9958a18d48690d77cc2151c4706aa7348b401.
//
// Solidity: event LockedERC20(address indexed depositor, address indexed depositReceiver, address indexed rootToken, uint256 amount)
func (_ERC20Predicate *ERC20PredicateFilterer) ParseLockedERC20(log types.Log) (*ERC20PredicateLockedERC20, error) {
	event := new(ERC20PredicateLockedERC20)
	if err := _ERC20Predicate.contract.UnpackLog(event, "LockedERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20PredicateRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the ERC20Predicate contract.
type ERC20PredicateRoleAdminChangedIterator struct {
	Event *ERC20PredicateRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *ERC20PredicateRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20PredicateRoleAdminChanged)
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
		it.Event = new(ERC20PredicateRoleAdminChanged)
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
func (it *ERC20PredicateRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20PredicateRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20PredicateRoleAdminChanged represents a RoleAdminChanged event raised by the ERC20Predicate contract.
type ERC20PredicateRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ERC20Predicate *ERC20PredicateFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ERC20PredicateRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _ERC20Predicate.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ERC20PredicateRoleAdminChangedIterator{contract: _ERC20Predicate.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ERC20Predicate *ERC20PredicateFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ERC20PredicateRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _ERC20Predicate.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20PredicateRoleAdminChanged)
				if err := _ERC20Predicate.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ERC20Predicate *ERC20PredicateFilterer) ParseRoleAdminChanged(log types.Log) (*ERC20PredicateRoleAdminChanged, error) {
	event := new(ERC20PredicateRoleAdminChanged)
	if err := _ERC20Predicate.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20PredicateRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the ERC20Predicate contract.
type ERC20PredicateRoleGrantedIterator struct {
	Event *ERC20PredicateRoleGranted // Event containing the contract specifics and raw log

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
func (it *ERC20PredicateRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20PredicateRoleGranted)
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
		it.Event = new(ERC20PredicateRoleGranted)
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
func (it *ERC20PredicateRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20PredicateRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20PredicateRoleGranted represents a RoleGranted event raised by the ERC20Predicate contract.
type ERC20PredicateRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ERC20Predicate *ERC20PredicateFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ERC20PredicateRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ERC20Predicate.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20PredicateRoleGrantedIterator{contract: _ERC20Predicate.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ERC20Predicate *ERC20PredicateFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ERC20PredicateRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ERC20Predicate.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20PredicateRoleGranted)
				if err := _ERC20Predicate.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ERC20Predicate *ERC20PredicateFilterer) ParseRoleGranted(log types.Log) (*ERC20PredicateRoleGranted, error) {
	event := new(ERC20PredicateRoleGranted)
	if err := _ERC20Predicate.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20PredicateRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the ERC20Predicate contract.
type ERC20PredicateRoleRevokedIterator struct {
	Event *ERC20PredicateRoleRevoked // Event containing the contract specifics and raw log

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
func (it *ERC20PredicateRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20PredicateRoleRevoked)
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
		it.Event = new(ERC20PredicateRoleRevoked)
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
func (it *ERC20PredicateRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20PredicateRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20PredicateRoleRevoked represents a RoleRevoked event raised by the ERC20Predicate contract.
type ERC20PredicateRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ERC20Predicate *ERC20PredicateFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ERC20PredicateRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ERC20Predicate.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20PredicateRoleRevokedIterator{contract: _ERC20Predicate.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ERC20Predicate *ERC20PredicateFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ERC20PredicateRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ERC20Predicate.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20PredicateRoleRevoked)
				if err := _ERC20Predicate.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ERC20Predicate *ERC20PredicateFilterer) ParseRoleRevoked(log types.Log) (*ERC20PredicateRoleRevoked, error) {
	event := new(ERC20PredicateRoleRevoked)
	if err := _ERC20Predicate.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
