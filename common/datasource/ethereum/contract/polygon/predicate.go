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

// PredicateMetaData contains all meta data concerning the Predicate contract.
var PredicateMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"exitor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ExitedEther\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"depositReceiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LockedEther\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MANAGER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN_TYPE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANSFER_EVENT_SIG\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"log\",\"type\":\"bytes\"}],\"name\":\"exitTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"depositReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"depositData\",\"type\":\"bytes\"}],\"name\":\"lockTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// PredicateABI is the input ABI used to generate the binding from.
// Deprecated: Use PredicateMetaData.ABI instead.
var PredicateABI = PredicateMetaData.ABI

// Predicate is an auto generated Go binding around an Ethereum contract.
type Predicate struct {
	PredicateCaller     // Read-only binding to the contract
	PredicateTransactor // Write-only binding to the contract
	PredicateFilterer   // Log filterer for contract events
}

// PredicateCaller is an auto generated read-only Go binding around an Ethereum contract.
type PredicateCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PredicateTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PredicateTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PredicateFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PredicateFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PredicateSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PredicateSession struct {
	Contract     *Predicate        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PredicateCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PredicateCallerSession struct {
	Contract *PredicateCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// PredicateTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PredicateTransactorSession struct {
	Contract     *PredicateTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// PredicateRaw is an auto generated low-level Go binding around an Ethereum contract.
type PredicateRaw struct {
	Contract *Predicate // Generic contract binding to access the raw methods on
}

// PredicateCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PredicateCallerRaw struct {
	Contract *PredicateCaller // Generic read-only contract binding to access the raw methods on
}

// PredicateTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PredicateTransactorRaw struct {
	Contract *PredicateTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPredicate creates a new instance of Predicate, bound to a specific deployed contract.
func NewPredicate(address common.Address, backend bind.ContractBackend) (*Predicate, error) {
	contract, err := bindPredicate(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Predicate{PredicateCaller: PredicateCaller{contract: contract}, PredicateTransactor: PredicateTransactor{contract: contract}, PredicateFilterer: PredicateFilterer{contract: contract}}, nil
}

// NewPredicateCaller creates a new read-only instance of Predicate, bound to a specific deployed contract.
func NewPredicateCaller(address common.Address, caller bind.ContractCaller) (*PredicateCaller, error) {
	contract, err := bindPredicate(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PredicateCaller{contract: contract}, nil
}

// NewPredicateTransactor creates a new write-only instance of Predicate, bound to a specific deployed contract.
func NewPredicateTransactor(address common.Address, transactor bind.ContractTransactor) (*PredicateTransactor, error) {
	contract, err := bindPredicate(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PredicateTransactor{contract: contract}, nil
}

// NewPredicateFilterer creates a new log filterer instance of Predicate, bound to a specific deployed contract.
func NewPredicateFilterer(address common.Address, filterer bind.ContractFilterer) (*PredicateFilterer, error) {
	contract, err := bindPredicate(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PredicateFilterer{contract: contract}, nil
}

// bindPredicate binds a generic wrapper to an already deployed contract.
func bindPredicate(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PredicateABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Predicate *PredicateRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Predicate.Contract.PredicateCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Predicate *PredicateRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Predicate.Contract.PredicateTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Predicate *PredicateRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Predicate.Contract.PredicateTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Predicate *PredicateCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Predicate.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Predicate *PredicateTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Predicate.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Predicate *PredicateTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Predicate.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Predicate *PredicateCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Predicate.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Predicate *PredicateSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Predicate.Contract.DEFAULTADMINROLE(&_Predicate.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Predicate *PredicateCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Predicate.Contract.DEFAULTADMINROLE(&_Predicate.CallOpts)
}

// MANAGERROLE is a free data retrieval call binding the contract method 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (_Predicate *PredicateCaller) MANAGERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Predicate.contract.Call(opts, &out, "MANAGER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MANAGERROLE is a free data retrieval call binding the contract method 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (_Predicate *PredicateSession) MANAGERROLE() ([32]byte, error) {
	return _Predicate.Contract.MANAGERROLE(&_Predicate.CallOpts)
}

// MANAGERROLE is a free data retrieval call binding the contract method 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (_Predicate *PredicateCallerSession) MANAGERROLE() ([32]byte, error) {
	return _Predicate.Contract.MANAGERROLE(&_Predicate.CallOpts)
}

// TOKENTYPE is a free data retrieval call binding the contract method 0x609c92b8.
//
// Solidity: function TOKEN_TYPE() view returns(bytes32)
func (_Predicate *PredicateCaller) TOKENTYPE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Predicate.contract.Call(opts, &out, "TOKEN_TYPE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TOKENTYPE is a free data retrieval call binding the contract method 0x609c92b8.
//
// Solidity: function TOKEN_TYPE() view returns(bytes32)
func (_Predicate *PredicateSession) TOKENTYPE() ([32]byte, error) {
	return _Predicate.Contract.TOKENTYPE(&_Predicate.CallOpts)
}

// TOKENTYPE is a free data retrieval call binding the contract method 0x609c92b8.
//
// Solidity: function TOKEN_TYPE() view returns(bytes32)
func (_Predicate *PredicateCallerSession) TOKENTYPE() ([32]byte, error) {
	return _Predicate.Contract.TOKENTYPE(&_Predicate.CallOpts)
}

// TRANSFEREVENTSIG is a free data retrieval call binding the contract method 0xb017a30f.
//
// Solidity: function TRANSFER_EVENT_SIG() view returns(bytes32)
func (_Predicate *PredicateCaller) TRANSFEREVENTSIG(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Predicate.contract.Call(opts, &out, "TRANSFER_EVENT_SIG")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TRANSFEREVENTSIG is a free data retrieval call binding the contract method 0xb017a30f.
//
// Solidity: function TRANSFER_EVENT_SIG() view returns(bytes32)
func (_Predicate *PredicateSession) TRANSFEREVENTSIG() ([32]byte, error) {
	return _Predicate.Contract.TRANSFEREVENTSIG(&_Predicate.CallOpts)
}

// TRANSFEREVENTSIG is a free data retrieval call binding the contract method 0xb017a30f.
//
// Solidity: function TRANSFER_EVENT_SIG() view returns(bytes32)
func (_Predicate *PredicateCallerSession) TRANSFEREVENTSIG() ([32]byte, error) {
	return _Predicate.Contract.TRANSFEREVENTSIG(&_Predicate.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Predicate *PredicateCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Predicate.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Predicate *PredicateSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Predicate.Contract.GetRoleAdmin(&_Predicate.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Predicate *PredicateCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Predicate.Contract.GetRoleAdmin(&_Predicate.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Predicate *PredicateCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Predicate.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Predicate *PredicateSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Predicate.Contract.GetRoleMember(&_Predicate.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Predicate *PredicateCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Predicate.Contract.GetRoleMember(&_Predicate.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Predicate *PredicateCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Predicate.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Predicate *PredicateSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Predicate.Contract.GetRoleMemberCount(&_Predicate.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Predicate *PredicateCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Predicate.Contract.GetRoleMemberCount(&_Predicate.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Predicate *PredicateCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Predicate.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Predicate *PredicateSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Predicate.Contract.HasRole(&_Predicate.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Predicate *PredicateCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Predicate.Contract.HasRole(&_Predicate.CallOpts, role, account)
}

// ExitTokens is a paid mutator transaction binding the contract method 0x8274664f.
//
// Solidity: function exitTokens(address , address , bytes log) returns()
func (_Predicate *PredicateTransactor) ExitTokens(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, log []byte) (*types.Transaction, error) {
	return _Predicate.contract.Transact(opts, "exitTokens", arg0, arg1, log)
}

// ExitTokens is a paid mutator transaction binding the contract method 0x8274664f.
//
// Solidity: function exitTokens(address , address , bytes log) returns()
func (_Predicate *PredicateSession) ExitTokens(arg0 common.Address, arg1 common.Address, log []byte) (*types.Transaction, error) {
	return _Predicate.Contract.ExitTokens(&_Predicate.TransactOpts, arg0, arg1, log)
}

// ExitTokens is a paid mutator transaction binding the contract method 0x8274664f.
//
// Solidity: function exitTokens(address , address , bytes log) returns()
func (_Predicate *PredicateTransactorSession) ExitTokens(arg0 common.Address, arg1 common.Address, log []byte) (*types.Transaction, error) {
	return _Predicate.Contract.ExitTokens(&_Predicate.TransactOpts, arg0, arg1, log)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Predicate *PredicateTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Predicate.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Predicate *PredicateSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Predicate.Contract.GrantRole(&_Predicate.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Predicate *PredicateTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Predicate.Contract.GrantRole(&_Predicate.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_Predicate *PredicateTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _Predicate.contract.Transact(opts, "initialize", _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_Predicate *PredicateSession) Initialize(_owner common.Address) (*types.Transaction, error) {
	return _Predicate.Contract.Initialize(&_Predicate.TransactOpts, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_Predicate *PredicateTransactorSession) Initialize(_owner common.Address) (*types.Transaction, error) {
	return _Predicate.Contract.Initialize(&_Predicate.TransactOpts, _owner)
}

// LockTokens is a paid mutator transaction binding the contract method 0xe375b64e.
//
// Solidity: function lockTokens(address depositor, address depositReceiver, address , bytes depositData) returns()
func (_Predicate *PredicateTransactor) LockTokens(opts *bind.TransactOpts, depositor common.Address, depositReceiver common.Address, arg2 common.Address, depositData []byte) (*types.Transaction, error) {
	return _Predicate.contract.Transact(opts, "lockTokens", depositor, depositReceiver, arg2, depositData)
}

// LockTokens is a paid mutator transaction binding the contract method 0xe375b64e.
//
// Solidity: function lockTokens(address depositor, address depositReceiver, address , bytes depositData) returns()
func (_Predicate *PredicateSession) LockTokens(depositor common.Address, depositReceiver common.Address, arg2 common.Address, depositData []byte) (*types.Transaction, error) {
	return _Predicate.Contract.LockTokens(&_Predicate.TransactOpts, depositor, depositReceiver, arg2, depositData)
}

// LockTokens is a paid mutator transaction binding the contract method 0xe375b64e.
//
// Solidity: function lockTokens(address depositor, address depositReceiver, address , bytes depositData) returns()
func (_Predicate *PredicateTransactorSession) LockTokens(depositor common.Address, depositReceiver common.Address, arg2 common.Address, depositData []byte) (*types.Transaction, error) {
	return _Predicate.Contract.LockTokens(&_Predicate.TransactOpts, depositor, depositReceiver, arg2, depositData)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Predicate *PredicateTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Predicate.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Predicate *PredicateSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Predicate.Contract.RenounceRole(&_Predicate.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Predicate *PredicateTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Predicate.Contract.RenounceRole(&_Predicate.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Predicate *PredicateTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Predicate.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Predicate *PredicateSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Predicate.Contract.RevokeRole(&_Predicate.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Predicate *PredicateTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Predicate.Contract.RevokeRole(&_Predicate.TransactOpts, role, account)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Predicate *PredicateTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Predicate.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Predicate *PredicateSession) Receive() (*types.Transaction, error) {
	return _Predicate.Contract.Receive(&_Predicate.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Predicate *PredicateTransactorSession) Receive() (*types.Transaction, error) {
	return _Predicate.Contract.Receive(&_Predicate.TransactOpts)
}

// PredicateExitedEtherIterator is returned from FilterExitedEther and is used to iterate over the raw logs and unpacked data for ExitedEther events raised by the Predicate contract.
type PredicateExitedEtherIterator struct {
	Event *PredicateExitedEther // Event containing the contract specifics and raw log

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
func (it *PredicateExitedEtherIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PredicateExitedEther)
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
		it.Event = new(PredicateExitedEther)
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
func (it *PredicateExitedEtherIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PredicateExitedEtherIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PredicateExitedEther represents a ExitedEther event raised by the Predicate contract.
type PredicateExitedEther struct {
	Exitor common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExitedEther is a free log retrieval operation binding the contract event 0x0fc0eed41f72d3da77d0f53b9594fc7073acd15ee9d7c536819a70a67c57ef3c.
//
// Solidity: event ExitedEther(address indexed exitor, uint256 amount)
func (_Predicate *PredicateFilterer) FilterExitedEther(opts *bind.FilterOpts, exitor []common.Address) (*PredicateExitedEtherIterator, error) {

	var exitorRule []interface{}
	for _, exitorItem := range exitor {
		exitorRule = append(exitorRule, exitorItem)
	}

	logs, sub, err := _Predicate.contract.FilterLogs(opts, "ExitedEther", exitorRule)
	if err != nil {
		return nil, err
	}
	return &PredicateExitedEtherIterator{contract: _Predicate.contract, event: "ExitedEther", logs: logs, sub: sub}, nil
}

// WatchExitedEther is a free log subscription operation binding the contract event 0x0fc0eed41f72d3da77d0f53b9594fc7073acd15ee9d7c536819a70a67c57ef3c.
//
// Solidity: event ExitedEther(address indexed exitor, uint256 amount)
func (_Predicate *PredicateFilterer) WatchExitedEther(opts *bind.WatchOpts, sink chan<- *PredicateExitedEther, exitor []common.Address) (event.Subscription, error) {

	var exitorRule []interface{}
	for _, exitorItem := range exitor {
		exitorRule = append(exitorRule, exitorItem)
	}

	logs, sub, err := _Predicate.contract.WatchLogs(opts, "ExitedEther", exitorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PredicateExitedEther)
				if err := _Predicate.contract.UnpackLog(event, "ExitedEther", log); err != nil {
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
func (_Predicate *PredicateFilterer) ParseExitedEther(log types.Log) (*PredicateExitedEther, error) {
	event := new(PredicateExitedEther)
	if err := _Predicate.contract.UnpackLog(event, "ExitedEther", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PredicateLockedEtherIterator is returned from FilterLockedEther and is used to iterate over the raw logs and unpacked data for LockedEther events raised by the Predicate contract.
type PredicateLockedEtherIterator struct {
	Event *PredicateLockedEther // Event containing the contract specifics and raw log

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
func (it *PredicateLockedEtherIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PredicateLockedEther)
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
		it.Event = new(PredicateLockedEther)
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
func (it *PredicateLockedEtherIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PredicateLockedEtherIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PredicateLockedEther represents a LockedEther event raised by the Predicate contract.
type PredicateLockedEther struct {
	Depositor       common.Address
	DepositReceiver common.Address
	Amount          *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLockedEther is a free log retrieval operation binding the contract event 0x3e799b2d61372379e767ef8f04d65089179b7a6f63f9be3065806456c7309f1b.
//
// Solidity: event LockedEther(address indexed depositor, address indexed depositReceiver, uint256 amount)
func (_Predicate *PredicateFilterer) FilterLockedEther(opts *bind.FilterOpts, depositor []common.Address, depositReceiver []common.Address) (*PredicateLockedEtherIterator, error) {

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}
	var depositReceiverRule []interface{}
	for _, depositReceiverItem := range depositReceiver {
		depositReceiverRule = append(depositReceiverRule, depositReceiverItem)
	}

	logs, sub, err := _Predicate.contract.FilterLogs(opts, "LockedEther", depositorRule, depositReceiverRule)
	if err != nil {
		return nil, err
	}
	return &PredicateLockedEtherIterator{contract: _Predicate.contract, event: "LockedEther", logs: logs, sub: sub}, nil
}

// WatchLockedEther is a free log subscription operation binding the contract event 0x3e799b2d61372379e767ef8f04d65089179b7a6f63f9be3065806456c7309f1b.
//
// Solidity: event LockedEther(address indexed depositor, address indexed depositReceiver, uint256 amount)
func (_Predicate *PredicateFilterer) WatchLockedEther(opts *bind.WatchOpts, sink chan<- *PredicateLockedEther, depositor []common.Address, depositReceiver []common.Address) (event.Subscription, error) {

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}
	var depositReceiverRule []interface{}
	for _, depositReceiverItem := range depositReceiver {
		depositReceiverRule = append(depositReceiverRule, depositReceiverItem)
	}

	logs, sub, err := _Predicate.contract.WatchLogs(opts, "LockedEther", depositorRule, depositReceiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PredicateLockedEther)
				if err := _Predicate.contract.UnpackLog(event, "LockedEther", log); err != nil {
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
func (_Predicate *PredicateFilterer) ParseLockedEther(log types.Log) (*PredicateLockedEther, error) {
	event := new(PredicateLockedEther)
	if err := _Predicate.contract.UnpackLog(event, "LockedEther", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PredicateRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Predicate contract.
type PredicateRoleAdminChangedIterator struct {
	Event *PredicateRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *PredicateRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PredicateRoleAdminChanged)
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
		it.Event = new(PredicateRoleAdminChanged)
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
func (it *PredicateRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PredicateRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PredicateRoleAdminChanged represents a RoleAdminChanged event raised by the Predicate contract.
type PredicateRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Predicate *PredicateFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*PredicateRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Predicate.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &PredicateRoleAdminChangedIterator{contract: _Predicate.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Predicate *PredicateFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *PredicateRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Predicate.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PredicateRoleAdminChanged)
				if err := _Predicate.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Predicate *PredicateFilterer) ParseRoleAdminChanged(log types.Log) (*PredicateRoleAdminChanged, error) {
	event := new(PredicateRoleAdminChanged)
	if err := _Predicate.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PredicateRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Predicate contract.
type PredicateRoleGrantedIterator struct {
	Event *PredicateRoleGranted // Event containing the contract specifics and raw log

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
func (it *PredicateRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PredicateRoleGranted)
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
		it.Event = new(PredicateRoleGranted)
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
func (it *PredicateRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PredicateRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PredicateRoleGranted represents a RoleGranted event raised by the Predicate contract.
type PredicateRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Predicate *PredicateFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PredicateRoleGrantedIterator, error) {

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

	logs, sub, err := _Predicate.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PredicateRoleGrantedIterator{contract: _Predicate.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Predicate *PredicateFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *PredicateRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Predicate.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PredicateRoleGranted)
				if err := _Predicate.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Predicate *PredicateFilterer) ParseRoleGranted(log types.Log) (*PredicateRoleGranted, error) {
	event := new(PredicateRoleGranted)
	if err := _Predicate.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PredicateRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Predicate contract.
type PredicateRoleRevokedIterator struct {
	Event *PredicateRoleRevoked // Event containing the contract specifics and raw log

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
func (it *PredicateRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PredicateRoleRevoked)
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
		it.Event = new(PredicateRoleRevoked)
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
func (it *PredicateRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PredicateRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PredicateRoleRevoked represents a RoleRevoked event raised by the Predicate contract.
type PredicateRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Predicate *PredicateFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PredicateRoleRevokedIterator, error) {

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

	logs, sub, err := _Predicate.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PredicateRoleRevokedIterator{contract: _Predicate.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Predicate *PredicateFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *PredicateRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Predicate.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PredicateRoleRevoked)
				if err := _Predicate.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Predicate *PredicateFilterer) ParseRoleRevoked(log types.Log) (*PredicateRoleRevoked, error) {
	event := new(PredicateRoleRevoked)
	if err := _Predicate.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
