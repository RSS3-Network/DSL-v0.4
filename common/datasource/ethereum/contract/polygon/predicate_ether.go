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

// EtherPredicateMetaData contains all meta data concerning the EtherPredicate contract.
var EtherPredicateMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"exitor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ExitedEther\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"depositReceiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LockedEther\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MANAGER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN_TYPE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANSFER_EVENT_SIG\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"log\",\"type\":\"bytes\"}],\"name\":\"exitTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"depositReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"depositData\",\"type\":\"bytes\"}],\"name\":\"lockTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// EtherPredicateABI is the input ABI used to generate the binding from.
// Deprecated: Use EtherPredicateMetaData.ABI instead.
var EtherPredicateABI = EtherPredicateMetaData.ABI

// EtherPredicate is an auto generated Go binding around an Ethereum contract.
type EtherPredicate struct {
	EtherPredicateCaller     // Read-only binding to the contract
	EtherPredicateTransactor // Write-only binding to the contract
	EtherPredicateFilterer   // Log filterer for contract events
}

// EtherPredicateCaller is an auto generated read-only Go binding around an Ethereum contract.
type EtherPredicateCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherPredicateTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EtherPredicateTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherPredicateFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EtherPredicateFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherPredicateSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EtherPredicateSession struct {
	Contract     *EtherPredicate   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EtherPredicateCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EtherPredicateCallerSession struct {
	Contract *EtherPredicateCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// EtherPredicateTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EtherPredicateTransactorSession struct {
	Contract     *EtherPredicateTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// EtherPredicateRaw is an auto generated low-level Go binding around an Ethereum contract.
type EtherPredicateRaw struct {
	Contract *EtherPredicate // Generic contract binding to access the raw methods on
}

// EtherPredicateCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EtherPredicateCallerRaw struct {
	Contract *EtherPredicateCaller // Generic read-only contract binding to access the raw methods on
}

// EtherPredicateTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EtherPredicateTransactorRaw struct {
	Contract *EtherPredicateTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEtherPredicate creates a new instance of EtherPredicate, bound to a specific deployed contract.
func NewEtherPredicate(address common.Address, backend bind.ContractBackend) (*EtherPredicate, error) {
	contract, err := bindEtherPredicate(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EtherPredicate{EtherPredicateCaller: EtherPredicateCaller{contract: contract}, EtherPredicateTransactor: EtherPredicateTransactor{contract: contract}, EtherPredicateFilterer: EtherPredicateFilterer{contract: contract}}, nil
}

// NewEtherPredicateCaller creates a new read-only instance of EtherPredicate, bound to a specific deployed contract.
func NewEtherPredicateCaller(address common.Address, caller bind.ContractCaller) (*EtherPredicateCaller, error) {
	contract, err := bindEtherPredicate(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EtherPredicateCaller{contract: contract}, nil
}

// NewEtherPredicateTransactor creates a new write-only instance of EtherPredicate, bound to a specific deployed contract.
func NewEtherPredicateTransactor(address common.Address, transactor bind.ContractTransactor) (*EtherPredicateTransactor, error) {
	contract, err := bindEtherPredicate(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EtherPredicateTransactor{contract: contract}, nil
}

// NewEtherPredicateFilterer creates a new log filterer instance of EtherPredicate, bound to a specific deployed contract.
func NewEtherPredicateFilterer(address common.Address, filterer bind.ContractFilterer) (*EtherPredicateFilterer, error) {
	contract, err := bindEtherPredicate(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EtherPredicateFilterer{contract: contract}, nil
}

// bindEtherPredicate binds a generic wrapper to an already deployed contract.
func bindEtherPredicate(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EtherPredicateABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EtherPredicate *EtherPredicateRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EtherPredicate.Contract.EtherPredicateCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EtherPredicate *EtherPredicateRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherPredicate.Contract.EtherPredicateTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EtherPredicate *EtherPredicateRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EtherPredicate.Contract.EtherPredicateTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EtherPredicate *EtherPredicateCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EtherPredicate.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EtherPredicate *EtherPredicateTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherPredicate.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EtherPredicate *EtherPredicateTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EtherPredicate.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_EtherPredicate *EtherPredicateCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EtherPredicate.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_EtherPredicate *EtherPredicateSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _EtherPredicate.Contract.DEFAULTADMINROLE(&_EtherPredicate.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_EtherPredicate *EtherPredicateCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _EtherPredicate.Contract.DEFAULTADMINROLE(&_EtherPredicate.CallOpts)
}

// MANAGERROLE is a free data retrieval call binding the contract method 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (_EtherPredicate *EtherPredicateCaller) MANAGERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EtherPredicate.contract.Call(opts, &out, "MANAGER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MANAGERROLE is a free data retrieval call binding the contract method 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (_EtherPredicate *EtherPredicateSession) MANAGERROLE() ([32]byte, error) {
	return _EtherPredicate.Contract.MANAGERROLE(&_EtherPredicate.CallOpts)
}

// MANAGERROLE is a free data retrieval call binding the contract method 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (_EtherPredicate *EtherPredicateCallerSession) MANAGERROLE() ([32]byte, error) {
	return _EtherPredicate.Contract.MANAGERROLE(&_EtherPredicate.CallOpts)
}

// TOKENTYPE is a free data retrieval call binding the contract method 0x609c92b8.
//
// Solidity: function TOKEN_TYPE() view returns(bytes32)
func (_EtherPredicate *EtherPredicateCaller) TOKENTYPE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EtherPredicate.contract.Call(opts, &out, "TOKEN_TYPE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TOKENTYPE is a free data retrieval call binding the contract method 0x609c92b8.
//
// Solidity: function TOKEN_TYPE() view returns(bytes32)
func (_EtherPredicate *EtherPredicateSession) TOKENTYPE() ([32]byte, error) {
	return _EtherPredicate.Contract.TOKENTYPE(&_EtherPredicate.CallOpts)
}

// TOKENTYPE is a free data retrieval call binding the contract method 0x609c92b8.
//
// Solidity: function TOKEN_TYPE() view returns(bytes32)
func (_EtherPredicate *EtherPredicateCallerSession) TOKENTYPE() ([32]byte, error) {
	return _EtherPredicate.Contract.TOKENTYPE(&_EtherPredicate.CallOpts)
}

// TRANSFEREVENTSIG is a free data retrieval call binding the contract method 0xb017a30f.
//
// Solidity: function TRANSFER_EVENT_SIG() view returns(bytes32)
func (_EtherPredicate *EtherPredicateCaller) TRANSFEREVENTSIG(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EtherPredicate.contract.Call(opts, &out, "TRANSFER_EVENT_SIG")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TRANSFEREVENTSIG is a free data retrieval call binding the contract method 0xb017a30f.
//
// Solidity: function TRANSFER_EVENT_SIG() view returns(bytes32)
func (_EtherPredicate *EtherPredicateSession) TRANSFEREVENTSIG() ([32]byte, error) {
	return _EtherPredicate.Contract.TRANSFEREVENTSIG(&_EtherPredicate.CallOpts)
}

// TRANSFEREVENTSIG is a free data retrieval call binding the contract method 0xb017a30f.
//
// Solidity: function TRANSFER_EVENT_SIG() view returns(bytes32)
func (_EtherPredicate *EtherPredicateCallerSession) TRANSFEREVENTSIG() ([32]byte, error) {
	return _EtherPredicate.Contract.TRANSFEREVENTSIG(&_EtherPredicate.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_EtherPredicate *EtherPredicateCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _EtherPredicate.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_EtherPredicate *EtherPredicateSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _EtherPredicate.Contract.GetRoleAdmin(&_EtherPredicate.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_EtherPredicate *EtherPredicateCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _EtherPredicate.Contract.GetRoleAdmin(&_EtherPredicate.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_EtherPredicate *EtherPredicateCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _EtherPredicate.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_EtherPredicate *EtherPredicateSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _EtherPredicate.Contract.GetRoleMember(&_EtherPredicate.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_EtherPredicate *EtherPredicateCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _EtherPredicate.Contract.GetRoleMember(&_EtherPredicate.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_EtherPredicate *EtherPredicateCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _EtherPredicate.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_EtherPredicate *EtherPredicateSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _EtherPredicate.Contract.GetRoleMemberCount(&_EtherPredicate.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_EtherPredicate *EtherPredicateCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _EtherPredicate.Contract.GetRoleMemberCount(&_EtherPredicate.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_EtherPredicate *EtherPredicateCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _EtherPredicate.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_EtherPredicate *EtherPredicateSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _EtherPredicate.Contract.HasRole(&_EtherPredicate.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_EtherPredicate *EtherPredicateCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _EtherPredicate.Contract.HasRole(&_EtherPredicate.CallOpts, role, account)
}

// ExitTokens is a paid mutator transaction binding the contract method 0x8274664f.
//
// Solidity: function exitTokens(address , address , bytes log) returns()
func (_EtherPredicate *EtherPredicateTransactor) ExitTokens(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, log []byte) (*types.Transaction, error) {
	return _EtherPredicate.contract.Transact(opts, "exitTokens", arg0, arg1, log)
}

// ExitTokens is a paid mutator transaction binding the contract method 0x8274664f.
//
// Solidity: function exitTokens(address , address , bytes log) returns()
func (_EtherPredicate *EtherPredicateSession) ExitTokens(arg0 common.Address, arg1 common.Address, log []byte) (*types.Transaction, error) {
	return _EtherPredicate.Contract.ExitTokens(&_EtherPredicate.TransactOpts, arg0, arg1, log)
}

// ExitTokens is a paid mutator transaction binding the contract method 0x8274664f.
//
// Solidity: function exitTokens(address , address , bytes log) returns()
func (_EtherPredicate *EtherPredicateTransactorSession) ExitTokens(arg0 common.Address, arg1 common.Address, log []byte) (*types.Transaction, error) {
	return _EtherPredicate.Contract.ExitTokens(&_EtherPredicate.TransactOpts, arg0, arg1, log)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_EtherPredicate *EtherPredicateTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _EtherPredicate.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_EtherPredicate *EtherPredicateSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _EtherPredicate.Contract.GrantRole(&_EtherPredicate.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_EtherPredicate *EtherPredicateTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _EtherPredicate.Contract.GrantRole(&_EtherPredicate.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_EtherPredicate *EtherPredicateTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _EtherPredicate.contract.Transact(opts, "initialize", _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_EtherPredicate *EtherPredicateSession) Initialize(_owner common.Address) (*types.Transaction, error) {
	return _EtherPredicate.Contract.Initialize(&_EtherPredicate.TransactOpts, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_EtherPredicate *EtherPredicateTransactorSession) Initialize(_owner common.Address) (*types.Transaction, error) {
	return _EtherPredicate.Contract.Initialize(&_EtherPredicate.TransactOpts, _owner)
}

// LockTokens is a paid mutator transaction binding the contract method 0xe375b64e.
//
// Solidity: function lockTokens(address depositor, address depositReceiver, address , bytes depositData) returns()
func (_EtherPredicate *EtherPredicateTransactor) LockTokens(opts *bind.TransactOpts, depositor common.Address, depositReceiver common.Address, arg2 common.Address, depositData []byte) (*types.Transaction, error) {
	return _EtherPredicate.contract.Transact(opts, "lockTokens", depositor, depositReceiver, arg2, depositData)
}

// LockTokens is a paid mutator transaction binding the contract method 0xe375b64e.
//
// Solidity: function lockTokens(address depositor, address depositReceiver, address , bytes depositData) returns()
func (_EtherPredicate *EtherPredicateSession) LockTokens(depositor common.Address, depositReceiver common.Address, arg2 common.Address, depositData []byte) (*types.Transaction, error) {
	return _EtherPredicate.Contract.LockTokens(&_EtherPredicate.TransactOpts, depositor, depositReceiver, arg2, depositData)
}

// LockTokens is a paid mutator transaction binding the contract method 0xe375b64e.
//
// Solidity: function lockTokens(address depositor, address depositReceiver, address , bytes depositData) returns()
func (_EtherPredicate *EtherPredicateTransactorSession) LockTokens(depositor common.Address, depositReceiver common.Address, arg2 common.Address, depositData []byte) (*types.Transaction, error) {
	return _EtherPredicate.Contract.LockTokens(&_EtherPredicate.TransactOpts, depositor, depositReceiver, arg2, depositData)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_EtherPredicate *EtherPredicateTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _EtherPredicate.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_EtherPredicate *EtherPredicateSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _EtherPredicate.Contract.RenounceRole(&_EtherPredicate.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_EtherPredicate *EtherPredicateTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _EtherPredicate.Contract.RenounceRole(&_EtherPredicate.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_EtherPredicate *EtherPredicateTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _EtherPredicate.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_EtherPredicate *EtherPredicateSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _EtherPredicate.Contract.RevokeRole(&_EtherPredicate.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_EtherPredicate *EtherPredicateTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _EtherPredicate.Contract.RevokeRole(&_EtherPredicate.TransactOpts, role, account)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_EtherPredicate *EtherPredicateTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherPredicate.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_EtherPredicate *EtherPredicateSession) Receive() (*types.Transaction, error) {
	return _EtherPredicate.Contract.Receive(&_EtherPredicate.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_EtherPredicate *EtherPredicateTransactorSession) Receive() (*types.Transaction, error) {
	return _EtherPredicate.Contract.Receive(&_EtherPredicate.TransactOpts)
}

// EtherPredicateExitedEtherIterator is returned from FilterExitedEther and is used to iterate over the raw logs and unpacked data for ExitedEther events raised by the EtherPredicate contract.
type EtherPredicateExitedEtherIterator struct {
	Event *EtherPredicateExitedEther // Event containing the contract specifics and raw log

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
func (it *EtherPredicateExitedEtherIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherPredicateExitedEther)
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
		it.Event = new(EtherPredicateExitedEther)
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
func (it *EtherPredicateExitedEtherIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherPredicateExitedEtherIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherPredicateExitedEther represents a ExitedEther event raised by the EtherPredicate contract.
type EtherPredicateExitedEther struct {
	Exitor common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExitedEther is a free log retrieval operation binding the contract event 0x0fc0eed41f72d3da77d0f53b9594fc7073acd15ee9d7c536819a70a67c57ef3c.
//
// Solidity: event ExitedEther(address indexed exitor, uint256 amount)
func (_EtherPredicate *EtherPredicateFilterer) FilterExitedEther(opts *bind.FilterOpts, exitor []common.Address) (*EtherPredicateExitedEtherIterator, error) {

	var exitorRule []interface{}
	for _, exitorItem := range exitor {
		exitorRule = append(exitorRule, exitorItem)
	}

	logs, sub, err := _EtherPredicate.contract.FilterLogs(opts, "ExitedEther", exitorRule)
	if err != nil {
		return nil, err
	}
	return &EtherPredicateExitedEtherIterator{contract: _EtherPredicate.contract, event: "ExitedEther", logs: logs, sub: sub}, nil
}

// WatchExitedEther is a free log subscription operation binding the contract event 0x0fc0eed41f72d3da77d0f53b9594fc7073acd15ee9d7c536819a70a67c57ef3c.
//
// Solidity: event ExitedEther(address indexed exitor, uint256 amount)
func (_EtherPredicate *EtherPredicateFilterer) WatchExitedEther(opts *bind.WatchOpts, sink chan<- *EtherPredicateExitedEther, exitor []common.Address) (event.Subscription, error) {

	var exitorRule []interface{}
	for _, exitorItem := range exitor {
		exitorRule = append(exitorRule, exitorItem)
	}

	logs, sub, err := _EtherPredicate.contract.WatchLogs(opts, "ExitedEther", exitorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherPredicateExitedEther)
				if err := _EtherPredicate.contract.UnpackLog(event, "ExitedEther", log); err != nil {
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

// ParseExitedEther is a log parse operation binding the contract event 0x0fc0eed41f72d3da77d0f53b9594fc7073acd15ee9d7c536819a70a67c57ef3c.
//
// Solidity: event ExitedEther(address indexed exitor, uint256 amount)
func (_EtherPredicate *EtherPredicateFilterer) ParseExitedEther(log types.Log) (*EtherPredicateExitedEther, error) {
	event := new(EtherPredicateExitedEther)
	if err := _EtherPredicate.contract.UnpackLog(event, "ExitedEther", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherPredicateLockedEtherIterator is returned from FilterLockedEther and is used to iterate over the raw logs and unpacked data for LockedEther events raised by the EtherPredicate contract.
type EtherPredicateLockedEtherIterator struct {
	Event *EtherPredicateLockedEther // Event containing the contract specifics and raw log

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
func (it *EtherPredicateLockedEtherIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherPredicateLockedEther)
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
		it.Event = new(EtherPredicateLockedEther)
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
func (it *EtherPredicateLockedEtherIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherPredicateLockedEtherIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherPredicateLockedEther represents a LockedEther event raised by the EtherPredicate contract.
type EtherPredicateLockedEther struct {
	Depositor       common.Address
	DepositReceiver common.Address
	Amount          *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLockedEther is a free log retrieval operation binding the contract event 0x3e799b2d61372379e767ef8f04d65089179b7a6f63f9be3065806456c7309f1b.
//
// Solidity: event LockedEther(address indexed depositor, address indexed depositReceiver, uint256 amount)
func (_EtherPredicate *EtherPredicateFilterer) FilterLockedEther(opts *bind.FilterOpts, depositor []common.Address, depositReceiver []common.Address) (*EtherPredicateLockedEtherIterator, error) {

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}
	var depositReceiverRule []interface{}
	for _, depositReceiverItem := range depositReceiver {
		depositReceiverRule = append(depositReceiverRule, depositReceiverItem)
	}

	logs, sub, err := _EtherPredicate.contract.FilterLogs(opts, "LockedEther", depositorRule, depositReceiverRule)
	if err != nil {
		return nil, err
	}
	return &EtherPredicateLockedEtherIterator{contract: _EtherPredicate.contract, event: "LockedEther", logs: logs, sub: sub}, nil
}

// WatchLockedEther is a free log subscription operation binding the contract event 0x3e799b2d61372379e767ef8f04d65089179b7a6f63f9be3065806456c7309f1b.
//
// Solidity: event LockedEther(address indexed depositor, address indexed depositReceiver, uint256 amount)
func (_EtherPredicate *EtherPredicateFilterer) WatchLockedEther(opts *bind.WatchOpts, sink chan<- *EtherPredicateLockedEther, depositor []common.Address, depositReceiver []common.Address) (event.Subscription, error) {

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}
	var depositReceiverRule []interface{}
	for _, depositReceiverItem := range depositReceiver {
		depositReceiverRule = append(depositReceiverRule, depositReceiverItem)
	}

	logs, sub, err := _EtherPredicate.contract.WatchLogs(opts, "LockedEther", depositorRule, depositReceiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherPredicateLockedEther)
				if err := _EtherPredicate.contract.UnpackLog(event, "LockedEther", log); err != nil {
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

// ParseLockedEther is a log parse operation binding the contract event 0x3e799b2d61372379e767ef8f04d65089179b7a6f63f9be3065806456c7309f1b.
//
// Solidity: event LockedEther(address indexed depositor, address indexed depositReceiver, uint256 amount)
func (_EtherPredicate *EtherPredicateFilterer) ParseLockedEther(log types.Log) (*EtherPredicateLockedEther, error) {
	event := new(EtherPredicateLockedEther)
	if err := _EtherPredicate.contract.UnpackLog(event, "LockedEther", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherPredicateRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the EtherPredicate contract.
type EtherPredicateRoleAdminChangedIterator struct {
	Event *EtherPredicateRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *EtherPredicateRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherPredicateRoleAdminChanged)
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
		it.Event = new(EtherPredicateRoleAdminChanged)
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
func (it *EtherPredicateRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherPredicateRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherPredicateRoleAdminChanged represents a RoleAdminChanged event raised by the EtherPredicate contract.
type EtherPredicateRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_EtherPredicate *EtherPredicateFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*EtherPredicateRoleAdminChangedIterator, error) {

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

	logs, sub, err := _EtherPredicate.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &EtherPredicateRoleAdminChangedIterator{contract: _EtherPredicate.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_EtherPredicate *EtherPredicateFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *EtherPredicateRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _EtherPredicate.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherPredicateRoleAdminChanged)
				if err := _EtherPredicate.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_EtherPredicate *EtherPredicateFilterer) ParseRoleAdminChanged(log types.Log) (*EtherPredicateRoleAdminChanged, error) {
	event := new(EtherPredicateRoleAdminChanged)
	if err := _EtherPredicate.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherPredicateRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the EtherPredicate contract.
type EtherPredicateRoleGrantedIterator struct {
	Event *EtherPredicateRoleGranted // Event containing the contract specifics and raw log

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
func (it *EtherPredicateRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherPredicateRoleGranted)
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
		it.Event = new(EtherPredicateRoleGranted)
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
func (it *EtherPredicateRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherPredicateRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherPredicateRoleGranted represents a RoleGranted event raised by the EtherPredicate contract.
type EtherPredicateRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_EtherPredicate *EtherPredicateFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*EtherPredicateRoleGrantedIterator, error) {

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

	logs, sub, err := _EtherPredicate.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &EtherPredicateRoleGrantedIterator{contract: _EtherPredicate.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_EtherPredicate *EtherPredicateFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *EtherPredicateRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _EtherPredicate.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherPredicateRoleGranted)
				if err := _EtherPredicate.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_EtherPredicate *EtherPredicateFilterer) ParseRoleGranted(log types.Log) (*EtherPredicateRoleGranted, error) {
	event := new(EtherPredicateRoleGranted)
	if err := _EtherPredicate.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherPredicateRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the EtherPredicate contract.
type EtherPredicateRoleRevokedIterator struct {
	Event *EtherPredicateRoleRevoked // Event containing the contract specifics and raw log

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
func (it *EtherPredicateRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherPredicateRoleRevoked)
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
		it.Event = new(EtherPredicateRoleRevoked)
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
func (it *EtherPredicateRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherPredicateRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherPredicateRoleRevoked represents a RoleRevoked event raised by the EtherPredicate contract.
type EtherPredicateRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_EtherPredicate *EtherPredicateFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*EtherPredicateRoleRevokedIterator, error) {

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

	logs, sub, err := _EtherPredicate.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &EtherPredicateRoleRevokedIterator{contract: _EtherPredicate.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_EtherPredicate *EtherPredicateFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *EtherPredicateRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _EtherPredicate.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherPredicateRoleRevoked)
				if err := _EtherPredicate.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_EtherPredicate *EtherPredicateFilterer) ParseRoleRevoked(log types.Log) (*EtherPredicateRoleRevoked, error) {
	event := new(EtherPredicateRoleRevoked)
	if err := _EtherPredicate.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
