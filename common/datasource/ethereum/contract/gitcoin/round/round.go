// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package round

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

// IRoundImplementationApplicationStatus is an auto generated low-level Go binding around an user-defined struct.
type IRoundImplementationApplicationStatus struct {
	Index     *big.Int
	StatusRow *big.Int
}

// MetaPtr is an auto generated low-level Go binding around an user-defined struct.
type MetaPtr struct {
	Protocol *big.Int
	Pointer  string
}

// RoundMetaData contains all meta data concerning the Round contract.
var RoundMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"protocol\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"pointer\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structMetaPtr\",\"name\":\"oldMetaPtr\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"protocol\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"pointer\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structMetaPtr\",\"name\":\"newMetaPtr\",\"type\":\"tuple\"}],\"name\":\"ApplicationMetaPtrUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"}],\"name\":\"ApplicationStatusesUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTime\",\"type\":\"uint256\"}],\"name\":\"ApplicationsEndTimeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTime\",\"type\":\"uint256\"}],\"name\":\"ApplicationsStartTimeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newAmount\",\"type\":\"uint256\"}],\"name\":\"MatchAmountUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"projectID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"applicationIndex\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"protocol\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"pointer\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structMetaPtr\",\"name\":\"applicationMetaPtr\",\"type\":\"tuple\"}],\"name\":\"NewProjectApplication\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"matchAmountAfterFees\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"protocolFeeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"roundFeeAmount\",\"type\":\"uint256\"}],\"name\":\"PayFeeAndEscrowFundsToPayoutContract\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"protocol\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"pointer\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structMetaPtr\",\"name\":\"oldMetaPtr\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"protocol\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"pointer\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structMetaPtr\",\"name\":\"newMetaPtr\",\"type\":\"tuple\"}],\"name\":\"ProjectsMetaPtrUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTime\",\"type\":\"uint256\"}],\"name\":\"RoundEndTimeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"roundFeeAddress\",\"type\":\"address\"}],\"name\":\"RoundFeeAddressUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"roundFeePercentage\",\"type\":\"uint32\"}],\"name\":\"RoundFeePercentageUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"protocol\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"pointer\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structMetaPtr\",\"name\":\"oldMetaPtr\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"protocol\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"pointer\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structMetaPtr\",\"name\":\"newMetaPtr\",\"type\":\"tuple\"}],\"name\":\"RoundMetaPtrUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTime\",\"type\":\"uint256\"}],\"name\":\"RoundStartTimeUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ROUND_OPERATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"alloSettings\",\"outputs\":[{\"internalType\":\"contractAlloSettings\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"applicationMetaPtr\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"protocol\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"pointer\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"applicationStatusesBitMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"applications\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"projectID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"applicationIndex\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"protocol\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"pointer\",\"type\":\"string\"}],\"internalType\":\"structMetaPtr\",\"name\":\"metaPtr\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"applicationsEndTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"applicationsIndexesByProjectID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"applicationsStartTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"projectID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"protocol\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"pointer\",\"type\":\"string\"}],\"internalType\":\"structMetaPtr\",\"name\":\"newApplicationMetaPtr\",\"type\":\"tuple\"}],\"name\":\"applyToRound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"projectID\",\"type\":\"bytes32\"}],\"name\":\"getApplicationIndexesByProjectID\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"applicationIndex\",\"type\":\"uint256\"}],\"name\":\"getApplicationStatus\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedParameters\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_alloSettings\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"matchAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextApplicationIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"payoutStrategy\",\"outputs\":[{\"internalType\":\"contractIPayoutStrategy\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"roundEndTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"roundFeeAddress\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"roundFeePercentage\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"roundMetaPtr\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"protocol\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"pointer\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"roundStartTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"statusRow\",\"type\":\"uint256\"}],\"internalType\":\"structIRoundImplementation.ApplicationStatus[]\",\"name\":\"statuses\",\"type\":\"tuple[]\"}],\"name\":\"setApplicationStatuses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setReadyForPayout\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"protocol\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"pointer\",\"type\":\"string\"}],\"internalType\":\"structMetaPtr\",\"name\":\"newApplicationMetaPtr\",\"type\":\"tuple\"}],\"name\":\"updateApplicationMetaPtr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newAmount\",\"type\":\"uint256\"}],\"name\":\"updateMatchAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newFeeAddress\",\"type\":\"address\"}],\"name\":\"updateRoundFeeAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"newFeePercentage\",\"type\":\"uint32\"}],\"name\":\"updateRoundFeePercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"protocol\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"pointer\",\"type\":\"string\"}],\"internalType\":\"structMetaPtr\",\"name\":\"newRoundMetaPtr\",\"type\":\"tuple\"}],\"name\":\"updateRoundMetaPtr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newApplicationsStartTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newApplicationsEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newRoundStartTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newRoundEndTime\",\"type\":\"uint256\"}],\"name\":\"updateStartAndEndTimes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"encodedVotes\",\"type\":\"bytes[]\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"votingStrategy\",\"outputs\":[{\"internalType\":\"contractIVotingStrategy\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"recipent\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// RoundABI is the input ABI used to generate the binding from.
// Deprecated: Use RoundMetaData.ABI instead.
var RoundABI = RoundMetaData.ABI

// Round is an auto generated Go binding around an Ethereum contract.
type Round struct {
	RoundCaller     // Read-only binding to the contract
	RoundTransactor // Write-only binding to the contract
	RoundFilterer   // Log filterer for contract events
}

// RoundCaller is an auto generated read-only Go binding around an Ethereum contract.
type RoundCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoundTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RoundTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoundFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RoundFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoundSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RoundSession struct {
	Contract     *Round            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RoundCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RoundCallerSession struct {
	Contract *RoundCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RoundTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RoundTransactorSession struct {
	Contract     *RoundTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RoundRaw is an auto generated low-level Go binding around an Ethereum contract.
type RoundRaw struct {
	Contract *Round // Generic contract binding to access the raw methods on
}

// RoundCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RoundCallerRaw struct {
	Contract *RoundCaller // Generic read-only contract binding to access the raw methods on
}

// RoundTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RoundTransactorRaw struct {
	Contract *RoundTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRound creates a new instance of Round, bound to a specific deployed contract.
func NewRound(address common.Address, backend bind.ContractBackend) (*Round, error) {
	contract, err := bindRound(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Round{RoundCaller: RoundCaller{contract: contract}, RoundTransactor: RoundTransactor{contract: contract}, RoundFilterer: RoundFilterer{contract: contract}}, nil
}

// NewRoundCaller creates a new read-only instance of Round, bound to a specific deployed contract.
func NewRoundCaller(address common.Address, caller bind.ContractCaller) (*RoundCaller, error) {
	contract, err := bindRound(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RoundCaller{contract: contract}, nil
}

// NewRoundTransactor creates a new write-only instance of Round, bound to a specific deployed contract.
func NewRoundTransactor(address common.Address, transactor bind.ContractTransactor) (*RoundTransactor, error) {
	contract, err := bindRound(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RoundTransactor{contract: contract}, nil
}

// NewRoundFilterer creates a new log filterer instance of Round, bound to a specific deployed contract.
func NewRoundFilterer(address common.Address, filterer bind.ContractFilterer) (*RoundFilterer, error) {
	contract, err := bindRound(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RoundFilterer{contract: contract}, nil
}

// bindRound binds a generic wrapper to an already deployed contract.
func bindRound(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RoundMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Round *RoundRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Round.Contract.RoundCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Round *RoundRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Round.Contract.RoundTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Round *RoundRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Round.Contract.RoundTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Round *RoundCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Round.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Round *RoundTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Round.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Round *RoundTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Round.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Round *RoundCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Round.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Round *RoundSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Round.Contract.DEFAULTADMINROLE(&_Round.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Round *RoundCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Round.Contract.DEFAULTADMINROLE(&_Round.CallOpts)
}

// ROUNDOPERATORROLE is a free data retrieval call binding the contract method 0xd97f3dcf.
//
// Solidity: function ROUND_OPERATOR_ROLE() view returns(bytes32)
func (_Round *RoundCaller) ROUNDOPERATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Round.contract.Call(opts, &out, "ROUND_OPERATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ROUNDOPERATORROLE is a free data retrieval call binding the contract method 0xd97f3dcf.
//
// Solidity: function ROUND_OPERATOR_ROLE() view returns(bytes32)
func (_Round *RoundSession) ROUNDOPERATORROLE() ([32]byte, error) {
	return _Round.Contract.ROUNDOPERATORROLE(&_Round.CallOpts)
}

// ROUNDOPERATORROLE is a free data retrieval call binding the contract method 0xd97f3dcf.
//
// Solidity: function ROUND_OPERATOR_ROLE() view returns(bytes32)
func (_Round *RoundCallerSession) ROUNDOPERATORROLE() ([32]byte, error) {
	return _Round.Contract.ROUNDOPERATORROLE(&_Round.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Round *RoundCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Round.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Round *RoundSession) VERSION() (string, error) {
	return _Round.Contract.VERSION(&_Round.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Round *RoundCallerSession) VERSION() (string, error) {
	return _Round.Contract.VERSION(&_Round.CallOpts)
}

// AlloSettings is a free data retrieval call binding the contract method 0x9c3b68ce.
//
// Solidity: function alloSettings() view returns(address)
func (_Round *RoundCaller) AlloSettings(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Round.contract.Call(opts, &out, "alloSettings")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AlloSettings is a free data retrieval call binding the contract method 0x9c3b68ce.
//
// Solidity: function alloSettings() view returns(address)
func (_Round *RoundSession) AlloSettings() (common.Address, error) {
	return _Round.Contract.AlloSettings(&_Round.CallOpts)
}

// AlloSettings is a free data retrieval call binding the contract method 0x9c3b68ce.
//
// Solidity: function alloSettings() view returns(address)
func (_Round *RoundCallerSession) AlloSettings() (common.Address, error) {
	return _Round.Contract.AlloSettings(&_Round.CallOpts)
}

// ApplicationMetaPtr is a free data retrieval call binding the contract method 0x685bd21b.
//
// Solidity: function applicationMetaPtr() view returns(uint256 protocol, string pointer)
func (_Round *RoundCaller) ApplicationMetaPtr(opts *bind.CallOpts) (struct {
	Protocol *big.Int
	Pointer  string
}, error) {
	var out []interface{}
	err := _Round.contract.Call(opts, &out, "applicationMetaPtr")

	outstruct := new(struct {
		Protocol *big.Int
		Pointer  string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Protocol = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Pointer = *abi.ConvertType(out[1], new(string)).(*string)

	return *outstruct, err

}

// ApplicationMetaPtr is a free data retrieval call binding the contract method 0x685bd21b.
//
// Solidity: function applicationMetaPtr() view returns(uint256 protocol, string pointer)
func (_Round *RoundSession) ApplicationMetaPtr() (struct {
	Protocol *big.Int
	Pointer  string
}, error) {
	return _Round.Contract.ApplicationMetaPtr(&_Round.CallOpts)
}

// ApplicationMetaPtr is a free data retrieval call binding the contract method 0x685bd21b.
//
// Solidity: function applicationMetaPtr() view returns(uint256 protocol, string pointer)
func (_Round *RoundCallerSession) ApplicationMetaPtr() (struct {
	Protocol *big.Int
	Pointer  string
}, error) {
	return _Round.Contract.ApplicationMetaPtr(&_Round.CallOpts)
}

// ApplicationStatusesBitMap is a free data retrieval call binding the contract method 0xc03fd662.
//
// Solidity: function applicationStatusesBitMap(uint256 ) view returns(uint256)
func (_Round *RoundCaller) ApplicationStatusesBitMap(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Round.contract.Call(opts, &out, "applicationStatusesBitMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ApplicationStatusesBitMap is a free data retrieval call binding the contract method 0xc03fd662.
//
// Solidity: function applicationStatusesBitMap(uint256 ) view returns(uint256)
func (_Round *RoundSession) ApplicationStatusesBitMap(arg0 *big.Int) (*big.Int, error) {
	return _Round.Contract.ApplicationStatusesBitMap(&_Round.CallOpts, arg0)
}

// ApplicationStatusesBitMap is a free data retrieval call binding the contract method 0xc03fd662.
//
// Solidity: function applicationStatusesBitMap(uint256 ) view returns(uint256)
func (_Round *RoundCallerSession) ApplicationStatusesBitMap(arg0 *big.Int) (*big.Int, error) {
	return _Round.Contract.ApplicationStatusesBitMap(&_Round.CallOpts, arg0)
}

// Applications is a free data retrieval call binding the contract method 0xdfefadff.
//
// Solidity: function applications(uint256 ) view returns(bytes32 projectID, uint256 applicationIndex, (uint256,string) metaPtr)
func (_Round *RoundCaller) Applications(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ProjectID        [32]byte
	ApplicationIndex *big.Int
	MetaPtr          MetaPtr
}, error) {
	var out []interface{}
	err := _Round.contract.Call(opts, &out, "applications", arg0)

	outstruct := new(struct {
		ProjectID        [32]byte
		ApplicationIndex *big.Int
		MetaPtr          MetaPtr
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ProjectID = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.ApplicationIndex = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.MetaPtr = *abi.ConvertType(out[2], new(MetaPtr)).(*MetaPtr)

	return *outstruct, err

}

// Applications is a free data retrieval call binding the contract method 0xdfefadff.
//
// Solidity: function applications(uint256 ) view returns(bytes32 projectID, uint256 applicationIndex, (uint256,string) metaPtr)
func (_Round *RoundSession) Applications(arg0 *big.Int) (struct {
	ProjectID        [32]byte
	ApplicationIndex *big.Int
	MetaPtr          MetaPtr
}, error) {
	return _Round.Contract.Applications(&_Round.CallOpts, arg0)
}

// Applications is a free data retrieval call binding the contract method 0xdfefadff.
//
// Solidity: function applications(uint256 ) view returns(bytes32 projectID, uint256 applicationIndex, (uint256,string) metaPtr)
func (_Round *RoundCallerSession) Applications(arg0 *big.Int) (struct {
	ProjectID        [32]byte
	ApplicationIndex *big.Int
	MetaPtr          MetaPtr
}, error) {
	return _Round.Contract.Applications(&_Round.CallOpts, arg0)
}

// ApplicationsEndTime is a free data retrieval call binding the contract method 0x1e949c9e.
//
// Solidity: function applicationsEndTime() view returns(uint256)
func (_Round *RoundCaller) ApplicationsEndTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Round.contract.Call(opts, &out, "applicationsEndTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ApplicationsEndTime is a free data retrieval call binding the contract method 0x1e949c9e.
//
// Solidity: function applicationsEndTime() view returns(uint256)
func (_Round *RoundSession) ApplicationsEndTime() (*big.Int, error) {
	return _Round.Contract.ApplicationsEndTime(&_Round.CallOpts)
}

// ApplicationsEndTime is a free data retrieval call binding the contract method 0x1e949c9e.
//
// Solidity: function applicationsEndTime() view returns(uint256)
func (_Round *RoundCallerSession) ApplicationsEndTime() (*big.Int, error) {
	return _Round.Contract.ApplicationsEndTime(&_Round.CallOpts)
}

// ApplicationsIndexesByProjectID is a free data retrieval call binding the contract method 0x10a11a6d.
//
// Solidity: function applicationsIndexesByProjectID(bytes32 , uint256 ) view returns(uint256)
func (_Round *RoundCaller) ApplicationsIndexesByProjectID(opts *bind.CallOpts, arg0 [32]byte, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Round.contract.Call(opts, &out, "applicationsIndexesByProjectID", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ApplicationsIndexesByProjectID is a free data retrieval call binding the contract method 0x10a11a6d.
//
// Solidity: function applicationsIndexesByProjectID(bytes32 , uint256 ) view returns(uint256)
func (_Round *RoundSession) ApplicationsIndexesByProjectID(arg0 [32]byte, arg1 *big.Int) (*big.Int, error) {
	return _Round.Contract.ApplicationsIndexesByProjectID(&_Round.CallOpts, arg0, arg1)
}

// ApplicationsIndexesByProjectID is a free data retrieval call binding the contract method 0x10a11a6d.
//
// Solidity: function applicationsIndexesByProjectID(bytes32 , uint256 ) view returns(uint256)
func (_Round *RoundCallerSession) ApplicationsIndexesByProjectID(arg0 [32]byte, arg1 *big.Int) (*big.Int, error) {
	return _Round.Contract.ApplicationsIndexesByProjectID(&_Round.CallOpts, arg0, arg1)
}

// ApplicationsStartTime is a free data retrieval call binding the contract method 0xd195bba1.
//
// Solidity: function applicationsStartTime() view returns(uint256)
func (_Round *RoundCaller) ApplicationsStartTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Round.contract.Call(opts, &out, "applicationsStartTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ApplicationsStartTime is a free data retrieval call binding the contract method 0xd195bba1.
//
// Solidity: function applicationsStartTime() view returns(uint256)
func (_Round *RoundSession) ApplicationsStartTime() (*big.Int, error) {
	return _Round.Contract.ApplicationsStartTime(&_Round.CallOpts)
}

// ApplicationsStartTime is a free data retrieval call binding the contract method 0xd195bba1.
//
// Solidity: function applicationsStartTime() view returns(uint256)
func (_Round *RoundCallerSession) ApplicationsStartTime() (*big.Int, error) {
	return _Round.Contract.ApplicationsStartTime(&_Round.CallOpts)
}

// GetApplicationIndexesByProjectID is a free data retrieval call binding the contract method 0xc5069b61.
//
// Solidity: function getApplicationIndexesByProjectID(bytes32 projectID) view returns(uint256[])
func (_Round *RoundCaller) GetApplicationIndexesByProjectID(opts *bind.CallOpts, projectID [32]byte) ([]*big.Int, error) {
	var out []interface{}
	err := _Round.contract.Call(opts, &out, "getApplicationIndexesByProjectID", projectID)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetApplicationIndexesByProjectID is a free data retrieval call binding the contract method 0xc5069b61.
//
// Solidity: function getApplicationIndexesByProjectID(bytes32 projectID) view returns(uint256[])
func (_Round *RoundSession) GetApplicationIndexesByProjectID(projectID [32]byte) ([]*big.Int, error) {
	return _Round.Contract.GetApplicationIndexesByProjectID(&_Round.CallOpts, projectID)
}

// GetApplicationIndexesByProjectID is a free data retrieval call binding the contract method 0xc5069b61.
//
// Solidity: function getApplicationIndexesByProjectID(bytes32 projectID) view returns(uint256[])
func (_Round *RoundCallerSession) GetApplicationIndexesByProjectID(projectID [32]byte) ([]*big.Int, error) {
	return _Round.Contract.GetApplicationIndexesByProjectID(&_Round.CallOpts, projectID)
}

// GetApplicationStatus is a free data retrieval call binding the contract method 0xd32cca47.
//
// Solidity: function getApplicationStatus(uint256 applicationIndex) view returns(uint256)
func (_Round *RoundCaller) GetApplicationStatus(opts *bind.CallOpts, applicationIndex *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Round.contract.Call(opts, &out, "getApplicationStatus", applicationIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetApplicationStatus is a free data retrieval call binding the contract method 0xd32cca47.
//
// Solidity: function getApplicationStatus(uint256 applicationIndex) view returns(uint256)
func (_Round *RoundSession) GetApplicationStatus(applicationIndex *big.Int) (*big.Int, error) {
	return _Round.Contract.GetApplicationStatus(&_Round.CallOpts, applicationIndex)
}

// GetApplicationStatus is a free data retrieval call binding the contract method 0xd32cca47.
//
// Solidity: function getApplicationStatus(uint256 applicationIndex) view returns(uint256)
func (_Round *RoundCallerSession) GetApplicationStatus(applicationIndex *big.Int) (*big.Int, error) {
	return _Round.Contract.GetApplicationStatus(&_Round.CallOpts, applicationIndex)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Round *RoundCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Round.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Round *RoundSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Round.Contract.GetRoleAdmin(&_Round.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Round *RoundCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Round.Contract.GetRoleAdmin(&_Round.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Round *RoundCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Round.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Round *RoundSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Round.Contract.GetRoleMember(&_Round.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Round *RoundCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Round.Contract.GetRoleMember(&_Round.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Round *RoundCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Round.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Round *RoundSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Round.Contract.GetRoleMemberCount(&_Round.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Round *RoundCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Round.Contract.GetRoleMemberCount(&_Round.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Round *RoundCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Round.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Round *RoundSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Round.Contract.HasRole(&_Round.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Round *RoundCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Round.Contract.HasRole(&_Round.CallOpts, role, account)
}

// MatchAmount is a free data retrieval call binding the contract method 0xf7f8ebca.
//
// Solidity: function matchAmount() view returns(uint256)
func (_Round *RoundCaller) MatchAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Round.contract.Call(opts, &out, "matchAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MatchAmount is a free data retrieval call binding the contract method 0xf7f8ebca.
//
// Solidity: function matchAmount() view returns(uint256)
func (_Round *RoundSession) MatchAmount() (*big.Int, error) {
	return _Round.Contract.MatchAmount(&_Round.CallOpts)
}

// MatchAmount is a free data retrieval call binding the contract method 0xf7f8ebca.
//
// Solidity: function matchAmount() view returns(uint256)
func (_Round *RoundCallerSession) MatchAmount() (*big.Int, error) {
	return _Round.Contract.MatchAmount(&_Round.CallOpts)
}

// NextApplicationIndex is a free data retrieval call binding the contract method 0x44b278a6.
//
// Solidity: function nextApplicationIndex() view returns(uint256)
func (_Round *RoundCaller) NextApplicationIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Round.contract.Call(opts, &out, "nextApplicationIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextApplicationIndex is a free data retrieval call binding the contract method 0x44b278a6.
//
// Solidity: function nextApplicationIndex() view returns(uint256)
func (_Round *RoundSession) NextApplicationIndex() (*big.Int, error) {
	return _Round.Contract.NextApplicationIndex(&_Round.CallOpts)
}

// NextApplicationIndex is a free data retrieval call binding the contract method 0x44b278a6.
//
// Solidity: function nextApplicationIndex() view returns(uint256)
func (_Round *RoundCallerSession) NextApplicationIndex() (*big.Int, error) {
	return _Round.Contract.NextApplicationIndex(&_Round.CallOpts)
}

// PayoutStrategy is a free data retrieval call binding the contract method 0x0c0a56b6.
//
// Solidity: function payoutStrategy() view returns(address)
func (_Round *RoundCaller) PayoutStrategy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Round.contract.Call(opts, &out, "payoutStrategy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PayoutStrategy is a free data retrieval call binding the contract method 0x0c0a56b6.
//
// Solidity: function payoutStrategy() view returns(address)
func (_Round *RoundSession) PayoutStrategy() (common.Address, error) {
	return _Round.Contract.PayoutStrategy(&_Round.CallOpts)
}

// PayoutStrategy is a free data retrieval call binding the contract method 0x0c0a56b6.
//
// Solidity: function payoutStrategy() view returns(address)
func (_Round *RoundCallerSession) PayoutStrategy() (common.Address, error) {
	return _Round.Contract.PayoutStrategy(&_Round.CallOpts)
}

// RoundEndTime is a free data retrieval call binding the contract method 0xe40205d6.
//
// Solidity: function roundEndTime() view returns(uint256)
func (_Round *RoundCaller) RoundEndTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Round.contract.Call(opts, &out, "roundEndTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RoundEndTime is a free data retrieval call binding the contract method 0xe40205d6.
//
// Solidity: function roundEndTime() view returns(uint256)
func (_Round *RoundSession) RoundEndTime() (*big.Int, error) {
	return _Round.Contract.RoundEndTime(&_Round.CallOpts)
}

// RoundEndTime is a free data retrieval call binding the contract method 0xe40205d6.
//
// Solidity: function roundEndTime() view returns(uint256)
func (_Round *RoundCallerSession) RoundEndTime() (*big.Int, error) {
	return _Round.Contract.RoundEndTime(&_Round.CallOpts)
}

// RoundFeeAddress is a free data retrieval call binding the contract method 0x76bacc99.
//
// Solidity: function roundFeeAddress() view returns(address)
func (_Round *RoundCaller) RoundFeeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Round.contract.Call(opts, &out, "roundFeeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RoundFeeAddress is a free data retrieval call binding the contract method 0x76bacc99.
//
// Solidity: function roundFeeAddress() view returns(address)
func (_Round *RoundSession) RoundFeeAddress() (common.Address, error) {
	return _Round.Contract.RoundFeeAddress(&_Round.CallOpts)
}

// RoundFeeAddress is a free data retrieval call binding the contract method 0x76bacc99.
//
// Solidity: function roundFeeAddress() view returns(address)
func (_Round *RoundCallerSession) RoundFeeAddress() (common.Address, error) {
	return _Round.Contract.RoundFeeAddress(&_Round.CallOpts)
}

// RoundFeePercentage is a free data retrieval call binding the contract method 0x8e9745dc.
//
// Solidity: function roundFeePercentage() view returns(uint32)
func (_Round *RoundCaller) RoundFeePercentage(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Round.contract.Call(opts, &out, "roundFeePercentage")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// RoundFeePercentage is a free data retrieval call binding the contract method 0x8e9745dc.
//
// Solidity: function roundFeePercentage() view returns(uint32)
func (_Round *RoundSession) RoundFeePercentage() (uint32, error) {
	return _Round.Contract.RoundFeePercentage(&_Round.CallOpts)
}

// RoundFeePercentage is a free data retrieval call binding the contract method 0x8e9745dc.
//
// Solidity: function roundFeePercentage() view returns(uint32)
func (_Round *RoundCallerSession) RoundFeePercentage() (uint32, error) {
	return _Round.Contract.RoundFeePercentage(&_Round.CallOpts)
}

// RoundMetaPtr is a free data retrieval call binding the contract method 0x2698f3fb.
//
// Solidity: function roundMetaPtr() view returns(uint256 protocol, string pointer)
func (_Round *RoundCaller) RoundMetaPtr(opts *bind.CallOpts) (struct {
	Protocol *big.Int
	Pointer  string
}, error) {
	var out []interface{}
	err := _Round.contract.Call(opts, &out, "roundMetaPtr")

	outstruct := new(struct {
		Protocol *big.Int
		Pointer  string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Protocol = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Pointer = *abi.ConvertType(out[1], new(string)).(*string)

	return *outstruct, err

}

// RoundMetaPtr is a free data retrieval call binding the contract method 0x2698f3fb.
//
// Solidity: function roundMetaPtr() view returns(uint256 protocol, string pointer)
func (_Round *RoundSession) RoundMetaPtr() (struct {
	Protocol *big.Int
	Pointer  string
}, error) {
	return _Round.Contract.RoundMetaPtr(&_Round.CallOpts)
}

// RoundMetaPtr is a free data retrieval call binding the contract method 0x2698f3fb.
//
// Solidity: function roundMetaPtr() view returns(uint256 protocol, string pointer)
func (_Round *RoundCallerSession) RoundMetaPtr() (struct {
	Protocol *big.Int
	Pointer  string
}, error) {
	return _Round.Contract.RoundMetaPtr(&_Round.CallOpts)
}

// RoundStartTime is a free data retrieval call binding the contract method 0xdd4f8f74.
//
// Solidity: function roundStartTime() view returns(uint256)
func (_Round *RoundCaller) RoundStartTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Round.contract.Call(opts, &out, "roundStartTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RoundStartTime is a free data retrieval call binding the contract method 0xdd4f8f74.
//
// Solidity: function roundStartTime() view returns(uint256)
func (_Round *RoundSession) RoundStartTime() (*big.Int, error) {
	return _Round.Contract.RoundStartTime(&_Round.CallOpts)
}

// RoundStartTime is a free data retrieval call binding the contract method 0xdd4f8f74.
//
// Solidity: function roundStartTime() view returns(uint256)
func (_Round *RoundCallerSession) RoundStartTime() (*big.Int, error) {
	return _Round.Contract.RoundStartTime(&_Round.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Round *RoundCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Round.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Round *RoundSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Round.Contract.SupportsInterface(&_Round.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Round *RoundCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Round.Contract.SupportsInterface(&_Round.CallOpts, interfaceId)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Round *RoundCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Round.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Round *RoundSession) Token() (common.Address, error) {
	return _Round.Contract.Token(&_Round.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Round *RoundCallerSession) Token() (common.Address, error) {
	return _Round.Contract.Token(&_Round.CallOpts)
}

// VotingStrategy is a free data retrieval call binding the contract method 0xfb5c8bfd.
//
// Solidity: function votingStrategy() view returns(address)
func (_Round *RoundCaller) VotingStrategy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Round.contract.Call(opts, &out, "votingStrategy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VotingStrategy is a free data retrieval call binding the contract method 0xfb5c8bfd.
//
// Solidity: function votingStrategy() view returns(address)
func (_Round *RoundSession) VotingStrategy() (common.Address, error) {
	return _Round.Contract.VotingStrategy(&_Round.CallOpts)
}

// VotingStrategy is a free data retrieval call binding the contract method 0xfb5c8bfd.
//
// Solidity: function votingStrategy() view returns(address)
func (_Round *RoundCallerSession) VotingStrategy() (common.Address, error) {
	return _Round.Contract.VotingStrategy(&_Round.CallOpts)
}

// ApplyToRound is a paid mutator transaction binding the contract method 0xb02f18e7.
//
// Solidity: function applyToRound(bytes32 projectID, (uint256,string) newApplicationMetaPtr) returns()
func (_Round *RoundTransactor) ApplyToRound(opts *bind.TransactOpts, projectID [32]byte, newApplicationMetaPtr MetaPtr) (*types.Transaction, error) {
	return _Round.contract.Transact(opts, "applyToRound", projectID, newApplicationMetaPtr)
}

// ApplyToRound is a paid mutator transaction binding the contract method 0xb02f18e7.
//
// Solidity: function applyToRound(bytes32 projectID, (uint256,string) newApplicationMetaPtr) returns()
func (_Round *RoundSession) ApplyToRound(projectID [32]byte, newApplicationMetaPtr MetaPtr) (*types.Transaction, error) {
	return _Round.Contract.ApplyToRound(&_Round.TransactOpts, projectID, newApplicationMetaPtr)
}

// ApplyToRound is a paid mutator transaction binding the contract method 0xb02f18e7.
//
// Solidity: function applyToRound(bytes32 projectID, (uint256,string) newApplicationMetaPtr) returns()
func (_Round *RoundTransactorSession) ApplyToRound(projectID [32]byte, newApplicationMetaPtr MetaPtr) (*types.Transaction, error) {
	return _Round.Contract.ApplyToRound(&_Round.TransactOpts, projectID, newApplicationMetaPtr)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Round *RoundTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Round.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Round *RoundSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Round.Contract.GrantRole(&_Round.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Round *RoundTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Round.Contract.GrantRole(&_Round.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xcce2df03.
//
// Solidity: function initialize(bytes encodedParameters, address _alloSettings) returns()
func (_Round *RoundTransactor) Initialize(opts *bind.TransactOpts, encodedParameters []byte, _alloSettings common.Address) (*types.Transaction, error) {
	return _Round.contract.Transact(opts, "initialize", encodedParameters, _alloSettings)
}

// Initialize is a paid mutator transaction binding the contract method 0xcce2df03.
//
// Solidity: function initialize(bytes encodedParameters, address _alloSettings) returns()
func (_Round *RoundSession) Initialize(encodedParameters []byte, _alloSettings common.Address) (*types.Transaction, error) {
	return _Round.Contract.Initialize(&_Round.TransactOpts, encodedParameters, _alloSettings)
}

// Initialize is a paid mutator transaction binding the contract method 0xcce2df03.
//
// Solidity: function initialize(bytes encodedParameters, address _alloSettings) returns()
func (_Round *RoundTransactorSession) Initialize(encodedParameters []byte, _alloSettings common.Address) (*types.Transaction, error) {
	return _Round.Contract.Initialize(&_Round.TransactOpts, encodedParameters, _alloSettings)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Round *RoundTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Round.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Round *RoundSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Round.Contract.RenounceRole(&_Round.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Round *RoundTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Round.Contract.RenounceRole(&_Round.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Round *RoundTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Round.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Round *RoundSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Round.Contract.RevokeRole(&_Round.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Round *RoundTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Round.Contract.RevokeRole(&_Round.TransactOpts, role, account)
}

// SetApplicationStatuses is a paid mutator transaction binding the contract method 0xfa1910e8.
//
// Solidity: function setApplicationStatuses((uint256,uint256)[] statuses) returns()
func (_Round *RoundTransactor) SetApplicationStatuses(opts *bind.TransactOpts, statuses []IRoundImplementationApplicationStatus) (*types.Transaction, error) {
	return _Round.contract.Transact(opts, "setApplicationStatuses", statuses)
}

// SetApplicationStatuses is a paid mutator transaction binding the contract method 0xfa1910e8.
//
// Solidity: function setApplicationStatuses((uint256,uint256)[] statuses) returns()
func (_Round *RoundSession) SetApplicationStatuses(statuses []IRoundImplementationApplicationStatus) (*types.Transaction, error) {
	return _Round.Contract.SetApplicationStatuses(&_Round.TransactOpts, statuses)
}

// SetApplicationStatuses is a paid mutator transaction binding the contract method 0xfa1910e8.
//
// Solidity: function setApplicationStatuses((uint256,uint256)[] statuses) returns()
func (_Round *RoundTransactorSession) SetApplicationStatuses(statuses []IRoundImplementationApplicationStatus) (*types.Transaction, error) {
	return _Round.Contract.SetApplicationStatuses(&_Round.TransactOpts, statuses)
}

// SetReadyForPayout is a paid mutator transaction binding the contract method 0xfe169f76.
//
// Solidity: function setReadyForPayout() payable returns()
func (_Round *RoundTransactor) SetReadyForPayout(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Round.contract.Transact(opts, "setReadyForPayout")
}

// SetReadyForPayout is a paid mutator transaction binding the contract method 0xfe169f76.
//
// Solidity: function setReadyForPayout() payable returns()
func (_Round *RoundSession) SetReadyForPayout() (*types.Transaction, error) {
	return _Round.Contract.SetReadyForPayout(&_Round.TransactOpts)
}

// SetReadyForPayout is a paid mutator transaction binding the contract method 0xfe169f76.
//
// Solidity: function setReadyForPayout() payable returns()
func (_Round *RoundTransactorSession) SetReadyForPayout() (*types.Transaction, error) {
	return _Round.Contract.SetReadyForPayout(&_Round.TransactOpts)
}

// UpdateApplicationMetaPtr is a paid mutator transaction binding the contract method 0x8926a888.
//
// Solidity: function updateApplicationMetaPtr((uint256,string) newApplicationMetaPtr) returns()
func (_Round *RoundTransactor) UpdateApplicationMetaPtr(opts *bind.TransactOpts, newApplicationMetaPtr MetaPtr) (*types.Transaction, error) {
	return _Round.contract.Transact(opts, "updateApplicationMetaPtr", newApplicationMetaPtr)
}

// UpdateApplicationMetaPtr is a paid mutator transaction binding the contract method 0x8926a888.
//
// Solidity: function updateApplicationMetaPtr((uint256,string) newApplicationMetaPtr) returns()
func (_Round *RoundSession) UpdateApplicationMetaPtr(newApplicationMetaPtr MetaPtr) (*types.Transaction, error) {
	return _Round.Contract.UpdateApplicationMetaPtr(&_Round.TransactOpts, newApplicationMetaPtr)
}

// UpdateApplicationMetaPtr is a paid mutator transaction binding the contract method 0x8926a888.
//
// Solidity: function updateApplicationMetaPtr((uint256,string) newApplicationMetaPtr) returns()
func (_Round *RoundTransactorSession) UpdateApplicationMetaPtr(newApplicationMetaPtr MetaPtr) (*types.Transaction, error) {
	return _Round.Contract.UpdateApplicationMetaPtr(&_Round.TransactOpts, newApplicationMetaPtr)
}

// UpdateMatchAmount is a paid mutator transaction binding the contract method 0x697e5f69.
//
// Solidity: function updateMatchAmount(uint256 newAmount) returns()
func (_Round *RoundTransactor) UpdateMatchAmount(opts *bind.TransactOpts, newAmount *big.Int) (*types.Transaction, error) {
	return _Round.contract.Transact(opts, "updateMatchAmount", newAmount)
}

// UpdateMatchAmount is a paid mutator transaction binding the contract method 0x697e5f69.
//
// Solidity: function updateMatchAmount(uint256 newAmount) returns()
func (_Round *RoundSession) UpdateMatchAmount(newAmount *big.Int) (*types.Transaction, error) {
	return _Round.Contract.UpdateMatchAmount(&_Round.TransactOpts, newAmount)
}

// UpdateMatchAmount is a paid mutator transaction binding the contract method 0x697e5f69.
//
// Solidity: function updateMatchAmount(uint256 newAmount) returns()
func (_Round *RoundTransactorSession) UpdateMatchAmount(newAmount *big.Int) (*types.Transaction, error) {
	return _Round.Contract.UpdateMatchAmount(&_Round.TransactOpts, newAmount)
}

// UpdateRoundFeeAddress is a paid mutator transaction binding the contract method 0xb9ff5ebe.
//
// Solidity: function updateRoundFeeAddress(address newFeeAddress) returns()
func (_Round *RoundTransactor) UpdateRoundFeeAddress(opts *bind.TransactOpts, newFeeAddress common.Address) (*types.Transaction, error) {
	return _Round.contract.Transact(opts, "updateRoundFeeAddress", newFeeAddress)
}

// UpdateRoundFeeAddress is a paid mutator transaction binding the contract method 0xb9ff5ebe.
//
// Solidity: function updateRoundFeeAddress(address newFeeAddress) returns()
func (_Round *RoundSession) UpdateRoundFeeAddress(newFeeAddress common.Address) (*types.Transaction, error) {
	return _Round.Contract.UpdateRoundFeeAddress(&_Round.TransactOpts, newFeeAddress)
}

// UpdateRoundFeeAddress is a paid mutator transaction binding the contract method 0xb9ff5ebe.
//
// Solidity: function updateRoundFeeAddress(address newFeeAddress) returns()
func (_Round *RoundTransactorSession) UpdateRoundFeeAddress(newFeeAddress common.Address) (*types.Transaction, error) {
	return _Round.Contract.UpdateRoundFeeAddress(&_Round.TransactOpts, newFeeAddress)
}

// UpdateRoundFeePercentage is a paid mutator transaction binding the contract method 0x28ab75c7.
//
// Solidity: function updateRoundFeePercentage(uint32 newFeePercentage) returns()
func (_Round *RoundTransactor) UpdateRoundFeePercentage(opts *bind.TransactOpts, newFeePercentage uint32) (*types.Transaction, error) {
	return _Round.contract.Transact(opts, "updateRoundFeePercentage", newFeePercentage)
}

// UpdateRoundFeePercentage is a paid mutator transaction binding the contract method 0x28ab75c7.
//
// Solidity: function updateRoundFeePercentage(uint32 newFeePercentage) returns()
func (_Round *RoundSession) UpdateRoundFeePercentage(newFeePercentage uint32) (*types.Transaction, error) {
	return _Round.Contract.UpdateRoundFeePercentage(&_Round.TransactOpts, newFeePercentage)
}

// UpdateRoundFeePercentage is a paid mutator transaction binding the contract method 0x28ab75c7.
//
// Solidity: function updateRoundFeePercentage(uint32 newFeePercentage) returns()
func (_Round *RoundTransactorSession) UpdateRoundFeePercentage(newFeePercentage uint32) (*types.Transaction, error) {
	return _Round.Contract.UpdateRoundFeePercentage(&_Round.TransactOpts, newFeePercentage)
}

// UpdateRoundMetaPtr is a paid mutator transaction binding the contract method 0xd7bf7840.
//
// Solidity: function updateRoundMetaPtr((uint256,string) newRoundMetaPtr) returns()
func (_Round *RoundTransactor) UpdateRoundMetaPtr(opts *bind.TransactOpts, newRoundMetaPtr MetaPtr) (*types.Transaction, error) {
	return _Round.contract.Transact(opts, "updateRoundMetaPtr", newRoundMetaPtr)
}

// UpdateRoundMetaPtr is a paid mutator transaction binding the contract method 0xd7bf7840.
//
// Solidity: function updateRoundMetaPtr((uint256,string) newRoundMetaPtr) returns()
func (_Round *RoundSession) UpdateRoundMetaPtr(newRoundMetaPtr MetaPtr) (*types.Transaction, error) {
	return _Round.Contract.UpdateRoundMetaPtr(&_Round.TransactOpts, newRoundMetaPtr)
}

// UpdateRoundMetaPtr is a paid mutator transaction binding the contract method 0xd7bf7840.
//
// Solidity: function updateRoundMetaPtr((uint256,string) newRoundMetaPtr) returns()
func (_Round *RoundTransactorSession) UpdateRoundMetaPtr(newRoundMetaPtr MetaPtr) (*types.Transaction, error) {
	return _Round.Contract.UpdateRoundMetaPtr(&_Round.TransactOpts, newRoundMetaPtr)
}

// UpdateStartAndEndTimes is a paid mutator transaction binding the contract method 0x8e421b85.
//
// Solidity: function updateStartAndEndTimes(uint256 newApplicationsStartTime, uint256 newApplicationsEndTime, uint256 newRoundStartTime, uint256 newRoundEndTime) returns()
func (_Round *RoundTransactor) UpdateStartAndEndTimes(opts *bind.TransactOpts, newApplicationsStartTime *big.Int, newApplicationsEndTime *big.Int, newRoundStartTime *big.Int, newRoundEndTime *big.Int) (*types.Transaction, error) {
	return _Round.contract.Transact(opts, "updateStartAndEndTimes", newApplicationsStartTime, newApplicationsEndTime, newRoundStartTime, newRoundEndTime)
}

// UpdateStartAndEndTimes is a paid mutator transaction binding the contract method 0x8e421b85.
//
// Solidity: function updateStartAndEndTimes(uint256 newApplicationsStartTime, uint256 newApplicationsEndTime, uint256 newRoundStartTime, uint256 newRoundEndTime) returns()
func (_Round *RoundSession) UpdateStartAndEndTimes(newApplicationsStartTime *big.Int, newApplicationsEndTime *big.Int, newRoundStartTime *big.Int, newRoundEndTime *big.Int) (*types.Transaction, error) {
	return _Round.Contract.UpdateStartAndEndTimes(&_Round.TransactOpts, newApplicationsStartTime, newApplicationsEndTime, newRoundStartTime, newRoundEndTime)
}

// UpdateStartAndEndTimes is a paid mutator transaction binding the contract method 0x8e421b85.
//
// Solidity: function updateStartAndEndTimes(uint256 newApplicationsStartTime, uint256 newApplicationsEndTime, uint256 newRoundStartTime, uint256 newRoundEndTime) returns()
func (_Round *RoundTransactorSession) UpdateStartAndEndTimes(newApplicationsStartTime *big.Int, newApplicationsEndTime *big.Int, newRoundStartTime *big.Int, newRoundEndTime *big.Int) (*types.Transaction, error) {
	return _Round.Contract.UpdateStartAndEndTimes(&_Round.TransactOpts, newApplicationsStartTime, newApplicationsEndTime, newRoundStartTime, newRoundEndTime)
}

// Vote is a paid mutator transaction binding the contract method 0x7aa54b68.
//
// Solidity: function vote(bytes[] encodedVotes) payable returns()
func (_Round *RoundTransactor) Vote(opts *bind.TransactOpts, encodedVotes [][]byte) (*types.Transaction, error) {
	return _Round.contract.Transact(opts, "vote", encodedVotes)
}

// Vote is a paid mutator transaction binding the contract method 0x7aa54b68.
//
// Solidity: function vote(bytes[] encodedVotes) payable returns()
func (_Round *RoundSession) Vote(encodedVotes [][]byte) (*types.Transaction, error) {
	return _Round.Contract.Vote(&_Round.TransactOpts, encodedVotes)
}

// Vote is a paid mutator transaction binding the contract method 0x7aa54b68.
//
// Solidity: function vote(bytes[] encodedVotes) payable returns()
func (_Round *RoundTransactorSession) Vote(encodedVotes [][]byte) (*types.Transaction, error) {
	return _Round.Contract.Vote(&_Round.TransactOpts, encodedVotes)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf940e385.
//
// Solidity: function withdraw(address tokenAddress, address recipent) returns()
func (_Round *RoundTransactor) Withdraw(opts *bind.TransactOpts, tokenAddress common.Address, recipent common.Address) (*types.Transaction, error) {
	return _Round.contract.Transact(opts, "withdraw", tokenAddress, recipent)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf940e385.
//
// Solidity: function withdraw(address tokenAddress, address recipent) returns()
func (_Round *RoundSession) Withdraw(tokenAddress common.Address, recipent common.Address) (*types.Transaction, error) {
	return _Round.Contract.Withdraw(&_Round.TransactOpts, tokenAddress, recipent)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf940e385.
//
// Solidity: function withdraw(address tokenAddress, address recipent) returns()
func (_Round *RoundTransactorSession) Withdraw(tokenAddress common.Address, recipent common.Address) (*types.Transaction, error) {
	return _Round.Contract.Withdraw(&_Round.TransactOpts, tokenAddress, recipent)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Round *RoundTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Round.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Round *RoundSession) Receive() (*types.Transaction, error) {
	return _Round.Contract.Receive(&_Round.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Round *RoundTransactorSession) Receive() (*types.Transaction, error) {
	return _Round.Contract.Receive(&_Round.TransactOpts)
}

// RoundApplicationMetaPtrUpdatedIterator is returned from FilterApplicationMetaPtrUpdated and is used to iterate over the raw logs and unpacked data for ApplicationMetaPtrUpdated events raised by the Round contract.
type RoundApplicationMetaPtrUpdatedIterator struct {
	Event *RoundApplicationMetaPtrUpdated // Event containing the contract specifics and raw log

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
func (it *RoundApplicationMetaPtrUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoundApplicationMetaPtrUpdated)
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
		it.Event = new(RoundApplicationMetaPtrUpdated)
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
func (it *RoundApplicationMetaPtrUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoundApplicationMetaPtrUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoundApplicationMetaPtrUpdated represents a ApplicationMetaPtrUpdated event raised by the Round contract.
type RoundApplicationMetaPtrUpdated struct {
	OldMetaPtr MetaPtr
	NewMetaPtr MetaPtr
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterApplicationMetaPtrUpdated is a free log retrieval operation binding the contract event 0xfd4633953debc7fbfababf529ffe33c4a6dd78c9ff13db552b571df088478641.
//
// Solidity: event ApplicationMetaPtrUpdated((uint256,string) oldMetaPtr, (uint256,string) newMetaPtr)
func (_Round *RoundFilterer) FilterApplicationMetaPtrUpdated(opts *bind.FilterOpts) (*RoundApplicationMetaPtrUpdatedIterator, error) {

	logs, sub, err := _Round.contract.FilterLogs(opts, "ApplicationMetaPtrUpdated")
	if err != nil {
		return nil, err
	}
	return &RoundApplicationMetaPtrUpdatedIterator{contract: _Round.contract, event: "ApplicationMetaPtrUpdated", logs: logs, sub: sub}, nil
}

// WatchApplicationMetaPtrUpdated is a free log subscription operation binding the contract event 0xfd4633953debc7fbfababf529ffe33c4a6dd78c9ff13db552b571df088478641.
//
// Solidity: event ApplicationMetaPtrUpdated((uint256,string) oldMetaPtr, (uint256,string) newMetaPtr)
func (_Round *RoundFilterer) WatchApplicationMetaPtrUpdated(opts *bind.WatchOpts, sink chan<- *RoundApplicationMetaPtrUpdated) (event.Subscription, error) {

	logs, sub, err := _Round.contract.WatchLogs(opts, "ApplicationMetaPtrUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoundApplicationMetaPtrUpdated)
				if err := _Round.contract.UnpackLog(event, "ApplicationMetaPtrUpdated", log); err != nil {
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

// ParseApplicationMetaPtrUpdated is a log parse operation binding the contract event 0xfd4633953debc7fbfababf529ffe33c4a6dd78c9ff13db552b571df088478641.
//
// Solidity: event ApplicationMetaPtrUpdated((uint256,string) oldMetaPtr, (uint256,string) newMetaPtr)
func (_Round *RoundFilterer) ParseApplicationMetaPtrUpdated(log types.Log) (*RoundApplicationMetaPtrUpdated, error) {
	event := new(RoundApplicationMetaPtrUpdated)
	if err := _Round.contract.UnpackLog(event, "ApplicationMetaPtrUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoundApplicationStatusesUpdatedIterator is returned from FilterApplicationStatusesUpdated and is used to iterate over the raw logs and unpacked data for ApplicationStatusesUpdated events raised by the Round contract.
type RoundApplicationStatusesUpdatedIterator struct {
	Event *RoundApplicationStatusesUpdated // Event containing the contract specifics and raw log

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
func (it *RoundApplicationStatusesUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoundApplicationStatusesUpdated)
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
		it.Event = new(RoundApplicationStatusesUpdated)
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
func (it *RoundApplicationStatusesUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoundApplicationStatusesUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoundApplicationStatusesUpdated represents a ApplicationStatusesUpdated event raised by the Round contract.
type RoundApplicationStatusesUpdated struct {
	Index  *big.Int
	Status *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterApplicationStatusesUpdated is a free log retrieval operation binding the contract event 0x5d498eb0ba08d4494201d7dd654e1f1c3e8a74dd0f860bb3c3e1b64eb79885c8.
//
// Solidity: event ApplicationStatusesUpdated(uint256 indexed index, uint256 indexed status)
func (_Round *RoundFilterer) FilterApplicationStatusesUpdated(opts *bind.FilterOpts, index []*big.Int, status []*big.Int) (*RoundApplicationStatusesUpdatedIterator, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}
	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _Round.contract.FilterLogs(opts, "ApplicationStatusesUpdated", indexRule, statusRule)
	if err != nil {
		return nil, err
	}
	return &RoundApplicationStatusesUpdatedIterator{contract: _Round.contract, event: "ApplicationStatusesUpdated", logs: logs, sub: sub}, nil
}

// WatchApplicationStatusesUpdated is a free log subscription operation binding the contract event 0x5d498eb0ba08d4494201d7dd654e1f1c3e8a74dd0f860bb3c3e1b64eb79885c8.
//
// Solidity: event ApplicationStatusesUpdated(uint256 indexed index, uint256 indexed status)
func (_Round *RoundFilterer) WatchApplicationStatusesUpdated(opts *bind.WatchOpts, sink chan<- *RoundApplicationStatusesUpdated, index []*big.Int, status []*big.Int) (event.Subscription, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}
	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _Round.contract.WatchLogs(opts, "ApplicationStatusesUpdated", indexRule, statusRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoundApplicationStatusesUpdated)
				if err := _Round.contract.UnpackLog(event, "ApplicationStatusesUpdated", log); err != nil {
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

// ParseApplicationStatusesUpdated is a log parse operation binding the contract event 0x5d498eb0ba08d4494201d7dd654e1f1c3e8a74dd0f860bb3c3e1b64eb79885c8.
//
// Solidity: event ApplicationStatusesUpdated(uint256 indexed index, uint256 indexed status)
func (_Round *RoundFilterer) ParseApplicationStatusesUpdated(log types.Log) (*RoundApplicationStatusesUpdated, error) {
	event := new(RoundApplicationStatusesUpdated)
	if err := _Round.contract.UnpackLog(event, "ApplicationStatusesUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoundApplicationsEndTimeUpdatedIterator is returned from FilterApplicationsEndTimeUpdated and is used to iterate over the raw logs and unpacked data for ApplicationsEndTimeUpdated events raised by the Round contract.
type RoundApplicationsEndTimeUpdatedIterator struct {
	Event *RoundApplicationsEndTimeUpdated // Event containing the contract specifics and raw log

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
func (it *RoundApplicationsEndTimeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoundApplicationsEndTimeUpdated)
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
		it.Event = new(RoundApplicationsEndTimeUpdated)
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
func (it *RoundApplicationsEndTimeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoundApplicationsEndTimeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoundApplicationsEndTimeUpdated represents a ApplicationsEndTimeUpdated event raised by the Round contract.
type RoundApplicationsEndTimeUpdated struct {
	OldTime *big.Int
	NewTime *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApplicationsEndTimeUpdated is a free log retrieval operation binding the contract event 0xa7fbc7103523f0d44281a3a3ba5aa119e9cab30ae02c09932eb2524c7d81f463.
//
// Solidity: event ApplicationsEndTimeUpdated(uint256 oldTime, uint256 newTime)
func (_Round *RoundFilterer) FilterApplicationsEndTimeUpdated(opts *bind.FilterOpts) (*RoundApplicationsEndTimeUpdatedIterator, error) {

	logs, sub, err := _Round.contract.FilterLogs(opts, "ApplicationsEndTimeUpdated")
	if err != nil {
		return nil, err
	}
	return &RoundApplicationsEndTimeUpdatedIterator{contract: _Round.contract, event: "ApplicationsEndTimeUpdated", logs: logs, sub: sub}, nil
}

// WatchApplicationsEndTimeUpdated is a free log subscription operation binding the contract event 0xa7fbc7103523f0d44281a3a3ba5aa119e9cab30ae02c09932eb2524c7d81f463.
//
// Solidity: event ApplicationsEndTimeUpdated(uint256 oldTime, uint256 newTime)
func (_Round *RoundFilterer) WatchApplicationsEndTimeUpdated(opts *bind.WatchOpts, sink chan<- *RoundApplicationsEndTimeUpdated) (event.Subscription, error) {

	logs, sub, err := _Round.contract.WatchLogs(opts, "ApplicationsEndTimeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoundApplicationsEndTimeUpdated)
				if err := _Round.contract.UnpackLog(event, "ApplicationsEndTimeUpdated", log); err != nil {
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

// ParseApplicationsEndTimeUpdated is a log parse operation binding the contract event 0xa7fbc7103523f0d44281a3a3ba5aa119e9cab30ae02c09932eb2524c7d81f463.
//
// Solidity: event ApplicationsEndTimeUpdated(uint256 oldTime, uint256 newTime)
func (_Round *RoundFilterer) ParseApplicationsEndTimeUpdated(log types.Log) (*RoundApplicationsEndTimeUpdated, error) {
	event := new(RoundApplicationsEndTimeUpdated)
	if err := _Round.contract.UnpackLog(event, "ApplicationsEndTimeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoundApplicationsStartTimeUpdatedIterator is returned from FilterApplicationsStartTimeUpdated and is used to iterate over the raw logs and unpacked data for ApplicationsStartTimeUpdated events raised by the Round contract.
type RoundApplicationsStartTimeUpdatedIterator struct {
	Event *RoundApplicationsStartTimeUpdated // Event containing the contract specifics and raw log

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
func (it *RoundApplicationsStartTimeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoundApplicationsStartTimeUpdated)
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
		it.Event = new(RoundApplicationsStartTimeUpdated)
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
func (it *RoundApplicationsStartTimeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoundApplicationsStartTimeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoundApplicationsStartTimeUpdated represents a ApplicationsStartTimeUpdated event raised by the Round contract.
type RoundApplicationsStartTimeUpdated struct {
	OldTime *big.Int
	NewTime *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApplicationsStartTimeUpdated is a free log retrieval operation binding the contract event 0x1fde68e630df12a4f2baa029dd016929d2c6bccc0fdd80dec2a787b01035517e.
//
// Solidity: event ApplicationsStartTimeUpdated(uint256 oldTime, uint256 newTime)
func (_Round *RoundFilterer) FilterApplicationsStartTimeUpdated(opts *bind.FilterOpts) (*RoundApplicationsStartTimeUpdatedIterator, error) {

	logs, sub, err := _Round.contract.FilterLogs(opts, "ApplicationsStartTimeUpdated")
	if err != nil {
		return nil, err
	}
	return &RoundApplicationsStartTimeUpdatedIterator{contract: _Round.contract, event: "ApplicationsStartTimeUpdated", logs: logs, sub: sub}, nil
}

// WatchApplicationsStartTimeUpdated is a free log subscription operation binding the contract event 0x1fde68e630df12a4f2baa029dd016929d2c6bccc0fdd80dec2a787b01035517e.
//
// Solidity: event ApplicationsStartTimeUpdated(uint256 oldTime, uint256 newTime)
func (_Round *RoundFilterer) WatchApplicationsStartTimeUpdated(opts *bind.WatchOpts, sink chan<- *RoundApplicationsStartTimeUpdated) (event.Subscription, error) {

	logs, sub, err := _Round.contract.WatchLogs(opts, "ApplicationsStartTimeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoundApplicationsStartTimeUpdated)
				if err := _Round.contract.UnpackLog(event, "ApplicationsStartTimeUpdated", log); err != nil {
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

// ParseApplicationsStartTimeUpdated is a log parse operation binding the contract event 0x1fde68e630df12a4f2baa029dd016929d2c6bccc0fdd80dec2a787b01035517e.
//
// Solidity: event ApplicationsStartTimeUpdated(uint256 oldTime, uint256 newTime)
func (_Round *RoundFilterer) ParseApplicationsStartTimeUpdated(log types.Log) (*RoundApplicationsStartTimeUpdated, error) {
	event := new(RoundApplicationsStartTimeUpdated)
	if err := _Round.contract.UnpackLog(event, "ApplicationsStartTimeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoundInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Round contract.
type RoundInitializedIterator struct {
	Event *RoundInitialized // Event containing the contract specifics and raw log

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
func (it *RoundInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoundInitialized)
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
		it.Event = new(RoundInitialized)
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
func (it *RoundInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoundInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoundInitialized represents a Initialized event raised by the Round contract.
type RoundInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Round *RoundFilterer) FilterInitialized(opts *bind.FilterOpts) (*RoundInitializedIterator, error) {

	logs, sub, err := _Round.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &RoundInitializedIterator{contract: _Round.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Round *RoundFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *RoundInitialized) (event.Subscription, error) {

	logs, sub, err := _Round.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoundInitialized)
				if err := _Round.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Round *RoundFilterer) ParseInitialized(log types.Log) (*RoundInitialized, error) {
	event := new(RoundInitialized)
	if err := _Round.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoundMatchAmountUpdatedIterator is returned from FilterMatchAmountUpdated and is used to iterate over the raw logs and unpacked data for MatchAmountUpdated events raised by the Round contract.
type RoundMatchAmountUpdatedIterator struct {
	Event *RoundMatchAmountUpdated // Event containing the contract specifics and raw log

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
func (it *RoundMatchAmountUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoundMatchAmountUpdated)
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
		it.Event = new(RoundMatchAmountUpdated)
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
func (it *RoundMatchAmountUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoundMatchAmountUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoundMatchAmountUpdated represents a MatchAmountUpdated event raised by the Round contract.
type RoundMatchAmountUpdated struct {
	NewAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMatchAmountUpdated is a free log retrieval operation binding the contract event 0x36407b7517f93b8dd6bf87384f27398dc6c7158776015e7492bbae639be60bdd.
//
// Solidity: event MatchAmountUpdated(uint256 newAmount)
func (_Round *RoundFilterer) FilterMatchAmountUpdated(opts *bind.FilterOpts) (*RoundMatchAmountUpdatedIterator, error) {

	logs, sub, err := _Round.contract.FilterLogs(opts, "MatchAmountUpdated")
	if err != nil {
		return nil, err
	}
	return &RoundMatchAmountUpdatedIterator{contract: _Round.contract, event: "MatchAmountUpdated", logs: logs, sub: sub}, nil
}

// WatchMatchAmountUpdated is a free log subscription operation binding the contract event 0x36407b7517f93b8dd6bf87384f27398dc6c7158776015e7492bbae639be60bdd.
//
// Solidity: event MatchAmountUpdated(uint256 newAmount)
func (_Round *RoundFilterer) WatchMatchAmountUpdated(opts *bind.WatchOpts, sink chan<- *RoundMatchAmountUpdated) (event.Subscription, error) {

	logs, sub, err := _Round.contract.WatchLogs(opts, "MatchAmountUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoundMatchAmountUpdated)
				if err := _Round.contract.UnpackLog(event, "MatchAmountUpdated", log); err != nil {
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

// ParseMatchAmountUpdated is a log parse operation binding the contract event 0x36407b7517f93b8dd6bf87384f27398dc6c7158776015e7492bbae639be60bdd.
//
// Solidity: event MatchAmountUpdated(uint256 newAmount)
func (_Round *RoundFilterer) ParseMatchAmountUpdated(log types.Log) (*RoundMatchAmountUpdated, error) {
	event := new(RoundMatchAmountUpdated)
	if err := _Round.contract.UnpackLog(event, "MatchAmountUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoundNewProjectApplicationIterator is returned from FilterNewProjectApplication and is used to iterate over the raw logs and unpacked data for NewProjectApplication events raised by the Round contract.
type RoundNewProjectApplicationIterator struct {
	Event *RoundNewProjectApplication // Event containing the contract specifics and raw log

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
func (it *RoundNewProjectApplicationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoundNewProjectApplication)
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
		it.Event = new(RoundNewProjectApplication)
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
func (it *RoundNewProjectApplicationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoundNewProjectApplicationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoundNewProjectApplication represents a NewProjectApplication event raised by the Round contract.
type RoundNewProjectApplication struct {
	ProjectID          [32]byte
	ApplicationIndex   *big.Int
	ApplicationMetaPtr MetaPtr
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterNewProjectApplication is a free log retrieval operation binding the contract event 0xca792622046325e9cd4e24b490cb000ef72acea3a15284efc14ee709307a5e00.
//
// Solidity: event NewProjectApplication(bytes32 indexed projectID, uint256 applicationIndex, (uint256,string) applicationMetaPtr)
func (_Round *RoundFilterer) FilterNewProjectApplication(opts *bind.FilterOpts, projectID [][32]byte) (*RoundNewProjectApplicationIterator, error) {

	var projectIDRule []interface{}
	for _, projectIDItem := range projectID {
		projectIDRule = append(projectIDRule, projectIDItem)
	}

	logs, sub, err := _Round.contract.FilterLogs(opts, "NewProjectApplication", projectIDRule)
	if err != nil {
		return nil, err
	}
	return &RoundNewProjectApplicationIterator{contract: _Round.contract, event: "NewProjectApplication", logs: logs, sub: sub}, nil
}

// WatchNewProjectApplication is a free log subscription operation binding the contract event 0xca792622046325e9cd4e24b490cb000ef72acea3a15284efc14ee709307a5e00.
//
// Solidity: event NewProjectApplication(bytes32 indexed projectID, uint256 applicationIndex, (uint256,string) applicationMetaPtr)
func (_Round *RoundFilterer) WatchNewProjectApplication(opts *bind.WatchOpts, sink chan<- *RoundNewProjectApplication, projectID [][32]byte) (event.Subscription, error) {

	var projectIDRule []interface{}
	for _, projectIDItem := range projectID {
		projectIDRule = append(projectIDRule, projectIDItem)
	}

	logs, sub, err := _Round.contract.WatchLogs(opts, "NewProjectApplication", projectIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoundNewProjectApplication)
				if err := _Round.contract.UnpackLog(event, "NewProjectApplication", log); err != nil {
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

// ParseNewProjectApplication is a log parse operation binding the contract event 0xca792622046325e9cd4e24b490cb000ef72acea3a15284efc14ee709307a5e00.
//
// Solidity: event NewProjectApplication(bytes32 indexed projectID, uint256 applicationIndex, (uint256,string) applicationMetaPtr)
func (_Round *RoundFilterer) ParseNewProjectApplication(log types.Log) (*RoundNewProjectApplication, error) {
	event := new(RoundNewProjectApplication)
	if err := _Round.contract.UnpackLog(event, "NewProjectApplication", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoundPayFeeAndEscrowFundsToPayoutContractIterator is returned from FilterPayFeeAndEscrowFundsToPayoutContract and is used to iterate over the raw logs and unpacked data for PayFeeAndEscrowFundsToPayoutContract events raised by the Round contract.
type RoundPayFeeAndEscrowFundsToPayoutContractIterator struct {
	Event *RoundPayFeeAndEscrowFundsToPayoutContract // Event containing the contract specifics and raw log

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
func (it *RoundPayFeeAndEscrowFundsToPayoutContractIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoundPayFeeAndEscrowFundsToPayoutContract)
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
		it.Event = new(RoundPayFeeAndEscrowFundsToPayoutContract)
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
func (it *RoundPayFeeAndEscrowFundsToPayoutContractIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoundPayFeeAndEscrowFundsToPayoutContractIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoundPayFeeAndEscrowFundsToPayoutContract represents a PayFeeAndEscrowFundsToPayoutContract event raised by the Round contract.
type RoundPayFeeAndEscrowFundsToPayoutContract struct {
	MatchAmountAfterFees *big.Int
	ProtocolFeeAmount    *big.Int
	RoundFeeAmount       *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterPayFeeAndEscrowFundsToPayoutContract is a free log retrieval operation binding the contract event 0xd3680093ec73f97fd84953765908e9ecbc0c567d97ee6e5465e742ace7b18bfc.
//
// Solidity: event PayFeeAndEscrowFundsToPayoutContract(uint256 matchAmountAfterFees, uint256 protocolFeeAmount, uint256 roundFeeAmount)
func (_Round *RoundFilterer) FilterPayFeeAndEscrowFundsToPayoutContract(opts *bind.FilterOpts) (*RoundPayFeeAndEscrowFundsToPayoutContractIterator, error) {

	logs, sub, err := _Round.contract.FilterLogs(opts, "PayFeeAndEscrowFundsToPayoutContract")
	if err != nil {
		return nil, err
	}
	return &RoundPayFeeAndEscrowFundsToPayoutContractIterator{contract: _Round.contract, event: "PayFeeAndEscrowFundsToPayoutContract", logs: logs, sub: sub}, nil
}

// WatchPayFeeAndEscrowFundsToPayoutContract is a free log subscription operation binding the contract event 0xd3680093ec73f97fd84953765908e9ecbc0c567d97ee6e5465e742ace7b18bfc.
//
// Solidity: event PayFeeAndEscrowFundsToPayoutContract(uint256 matchAmountAfterFees, uint256 protocolFeeAmount, uint256 roundFeeAmount)
func (_Round *RoundFilterer) WatchPayFeeAndEscrowFundsToPayoutContract(opts *bind.WatchOpts, sink chan<- *RoundPayFeeAndEscrowFundsToPayoutContract) (event.Subscription, error) {

	logs, sub, err := _Round.contract.WatchLogs(opts, "PayFeeAndEscrowFundsToPayoutContract")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoundPayFeeAndEscrowFundsToPayoutContract)
				if err := _Round.contract.UnpackLog(event, "PayFeeAndEscrowFundsToPayoutContract", log); err != nil {
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

// ParsePayFeeAndEscrowFundsToPayoutContract is a log parse operation binding the contract event 0xd3680093ec73f97fd84953765908e9ecbc0c567d97ee6e5465e742ace7b18bfc.
//
// Solidity: event PayFeeAndEscrowFundsToPayoutContract(uint256 matchAmountAfterFees, uint256 protocolFeeAmount, uint256 roundFeeAmount)
func (_Round *RoundFilterer) ParsePayFeeAndEscrowFundsToPayoutContract(log types.Log) (*RoundPayFeeAndEscrowFundsToPayoutContract, error) {
	event := new(RoundPayFeeAndEscrowFundsToPayoutContract)
	if err := _Round.contract.UnpackLog(event, "PayFeeAndEscrowFundsToPayoutContract", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoundProjectsMetaPtrUpdatedIterator is returned from FilterProjectsMetaPtrUpdated and is used to iterate over the raw logs and unpacked data for ProjectsMetaPtrUpdated events raised by the Round contract.
type RoundProjectsMetaPtrUpdatedIterator struct {
	Event *RoundProjectsMetaPtrUpdated // Event containing the contract specifics and raw log

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
func (it *RoundProjectsMetaPtrUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoundProjectsMetaPtrUpdated)
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
		it.Event = new(RoundProjectsMetaPtrUpdated)
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
func (it *RoundProjectsMetaPtrUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoundProjectsMetaPtrUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoundProjectsMetaPtrUpdated represents a ProjectsMetaPtrUpdated event raised by the Round contract.
type RoundProjectsMetaPtrUpdated struct {
	OldMetaPtr MetaPtr
	NewMetaPtr MetaPtr
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProjectsMetaPtrUpdated is a free log retrieval operation binding the contract event 0xd89ecd28a8dd4688401fd735d473fe15800c2551b32188cbefda0db6677dcc34.
//
// Solidity: event ProjectsMetaPtrUpdated((uint256,string) oldMetaPtr, (uint256,string) newMetaPtr)
func (_Round *RoundFilterer) FilterProjectsMetaPtrUpdated(opts *bind.FilterOpts) (*RoundProjectsMetaPtrUpdatedIterator, error) {

	logs, sub, err := _Round.contract.FilterLogs(opts, "ProjectsMetaPtrUpdated")
	if err != nil {
		return nil, err
	}
	return &RoundProjectsMetaPtrUpdatedIterator{contract: _Round.contract, event: "ProjectsMetaPtrUpdated", logs: logs, sub: sub}, nil
}

// WatchProjectsMetaPtrUpdated is a free log subscription operation binding the contract event 0xd89ecd28a8dd4688401fd735d473fe15800c2551b32188cbefda0db6677dcc34.
//
// Solidity: event ProjectsMetaPtrUpdated((uint256,string) oldMetaPtr, (uint256,string) newMetaPtr)
func (_Round *RoundFilterer) WatchProjectsMetaPtrUpdated(opts *bind.WatchOpts, sink chan<- *RoundProjectsMetaPtrUpdated) (event.Subscription, error) {

	logs, sub, err := _Round.contract.WatchLogs(opts, "ProjectsMetaPtrUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoundProjectsMetaPtrUpdated)
				if err := _Round.contract.UnpackLog(event, "ProjectsMetaPtrUpdated", log); err != nil {
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

// ParseProjectsMetaPtrUpdated is a log parse operation binding the contract event 0xd89ecd28a8dd4688401fd735d473fe15800c2551b32188cbefda0db6677dcc34.
//
// Solidity: event ProjectsMetaPtrUpdated((uint256,string) oldMetaPtr, (uint256,string) newMetaPtr)
func (_Round *RoundFilterer) ParseProjectsMetaPtrUpdated(log types.Log) (*RoundProjectsMetaPtrUpdated, error) {
	event := new(RoundProjectsMetaPtrUpdated)
	if err := _Round.contract.UnpackLog(event, "ProjectsMetaPtrUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoundRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Round contract.
type RoundRoleAdminChangedIterator struct {
	Event *RoundRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *RoundRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoundRoleAdminChanged)
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
		it.Event = new(RoundRoleAdminChanged)
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
func (it *RoundRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoundRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoundRoleAdminChanged represents a RoleAdminChanged event raised by the Round contract.
type RoundRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Round *RoundFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*RoundRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Round.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &RoundRoleAdminChangedIterator{contract: _Round.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Round *RoundFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *RoundRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Round.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoundRoleAdminChanged)
				if err := _Round.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Round *RoundFilterer) ParseRoleAdminChanged(log types.Log) (*RoundRoleAdminChanged, error) {
	event := new(RoundRoleAdminChanged)
	if err := _Round.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoundRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Round contract.
type RoundRoleGrantedIterator struct {
	Event *RoundRoleGranted // Event containing the contract specifics and raw log

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
func (it *RoundRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoundRoleGranted)
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
		it.Event = new(RoundRoleGranted)
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
func (it *RoundRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoundRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoundRoleGranted represents a RoleGranted event raised by the Round contract.
type RoundRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Round *RoundFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*RoundRoleGrantedIterator, error) {

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

	logs, sub, err := _Round.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &RoundRoleGrantedIterator{contract: _Round.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Round *RoundFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *RoundRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Round.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoundRoleGranted)
				if err := _Round.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Round *RoundFilterer) ParseRoleGranted(log types.Log) (*RoundRoleGranted, error) {
	event := new(RoundRoleGranted)
	if err := _Round.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoundRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Round contract.
type RoundRoleRevokedIterator struct {
	Event *RoundRoleRevoked // Event containing the contract specifics and raw log

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
func (it *RoundRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoundRoleRevoked)
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
		it.Event = new(RoundRoleRevoked)
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
func (it *RoundRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoundRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoundRoleRevoked represents a RoleRevoked event raised by the Round contract.
type RoundRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Round *RoundFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*RoundRoleRevokedIterator, error) {

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

	logs, sub, err := _Round.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &RoundRoleRevokedIterator{contract: _Round.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Round *RoundFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *RoundRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Round.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoundRoleRevoked)
				if err := _Round.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Round *RoundFilterer) ParseRoleRevoked(log types.Log) (*RoundRoleRevoked, error) {
	event := new(RoundRoleRevoked)
	if err := _Round.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoundRoundEndTimeUpdatedIterator is returned from FilterRoundEndTimeUpdated and is used to iterate over the raw logs and unpacked data for RoundEndTimeUpdated events raised by the Round contract.
type RoundRoundEndTimeUpdatedIterator struct {
	Event *RoundRoundEndTimeUpdated // Event containing the contract specifics and raw log

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
func (it *RoundRoundEndTimeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoundRoundEndTimeUpdated)
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
		it.Event = new(RoundRoundEndTimeUpdated)
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
func (it *RoundRoundEndTimeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoundRoundEndTimeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoundRoundEndTimeUpdated represents a RoundEndTimeUpdated event raised by the Round contract.
type RoundRoundEndTimeUpdated struct {
	OldTime *big.Int
	NewTime *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoundEndTimeUpdated is a free log retrieval operation binding the contract event 0xc0395d95a518149baa3521bb5b3c6988c7806a48ea42726ee3a9686bbd2a08cd.
//
// Solidity: event RoundEndTimeUpdated(uint256 oldTime, uint256 newTime)
func (_Round *RoundFilterer) FilterRoundEndTimeUpdated(opts *bind.FilterOpts) (*RoundRoundEndTimeUpdatedIterator, error) {

	logs, sub, err := _Round.contract.FilterLogs(opts, "RoundEndTimeUpdated")
	if err != nil {
		return nil, err
	}
	return &RoundRoundEndTimeUpdatedIterator{contract: _Round.contract, event: "RoundEndTimeUpdated", logs: logs, sub: sub}, nil
}

// WatchRoundEndTimeUpdated is a free log subscription operation binding the contract event 0xc0395d95a518149baa3521bb5b3c6988c7806a48ea42726ee3a9686bbd2a08cd.
//
// Solidity: event RoundEndTimeUpdated(uint256 oldTime, uint256 newTime)
func (_Round *RoundFilterer) WatchRoundEndTimeUpdated(opts *bind.WatchOpts, sink chan<- *RoundRoundEndTimeUpdated) (event.Subscription, error) {

	logs, sub, err := _Round.contract.WatchLogs(opts, "RoundEndTimeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoundRoundEndTimeUpdated)
				if err := _Round.contract.UnpackLog(event, "RoundEndTimeUpdated", log); err != nil {
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

// ParseRoundEndTimeUpdated is a log parse operation binding the contract event 0xc0395d95a518149baa3521bb5b3c6988c7806a48ea42726ee3a9686bbd2a08cd.
//
// Solidity: event RoundEndTimeUpdated(uint256 oldTime, uint256 newTime)
func (_Round *RoundFilterer) ParseRoundEndTimeUpdated(log types.Log) (*RoundRoundEndTimeUpdated, error) {
	event := new(RoundRoundEndTimeUpdated)
	if err := _Round.contract.UnpackLog(event, "RoundEndTimeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoundRoundFeeAddressUpdatedIterator is returned from FilterRoundFeeAddressUpdated and is used to iterate over the raw logs and unpacked data for RoundFeeAddressUpdated events raised by the Round contract.
type RoundRoundFeeAddressUpdatedIterator struct {
	Event *RoundRoundFeeAddressUpdated // Event containing the contract specifics and raw log

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
func (it *RoundRoundFeeAddressUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoundRoundFeeAddressUpdated)
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
		it.Event = new(RoundRoundFeeAddressUpdated)
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
func (it *RoundRoundFeeAddressUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoundRoundFeeAddressUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoundRoundFeeAddressUpdated represents a RoundFeeAddressUpdated event raised by the Round contract.
type RoundRoundFeeAddressUpdated struct {
	RoundFeeAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRoundFeeAddressUpdated is a free log retrieval operation binding the contract event 0x88ae92ef8eb8a634a01f4e404d706ac67bc00a88bdf46af83630c6ade6e97953.
//
// Solidity: event RoundFeeAddressUpdated(address roundFeeAddress)
func (_Round *RoundFilterer) FilterRoundFeeAddressUpdated(opts *bind.FilterOpts) (*RoundRoundFeeAddressUpdatedIterator, error) {

	logs, sub, err := _Round.contract.FilterLogs(opts, "RoundFeeAddressUpdated")
	if err != nil {
		return nil, err
	}
	return &RoundRoundFeeAddressUpdatedIterator{contract: _Round.contract, event: "RoundFeeAddressUpdated", logs: logs, sub: sub}, nil
}

// WatchRoundFeeAddressUpdated is a free log subscription operation binding the contract event 0x88ae92ef8eb8a634a01f4e404d706ac67bc00a88bdf46af83630c6ade6e97953.
//
// Solidity: event RoundFeeAddressUpdated(address roundFeeAddress)
func (_Round *RoundFilterer) WatchRoundFeeAddressUpdated(opts *bind.WatchOpts, sink chan<- *RoundRoundFeeAddressUpdated) (event.Subscription, error) {

	logs, sub, err := _Round.contract.WatchLogs(opts, "RoundFeeAddressUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoundRoundFeeAddressUpdated)
				if err := _Round.contract.UnpackLog(event, "RoundFeeAddressUpdated", log); err != nil {
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

// ParseRoundFeeAddressUpdated is a log parse operation binding the contract event 0x88ae92ef8eb8a634a01f4e404d706ac67bc00a88bdf46af83630c6ade6e97953.
//
// Solidity: event RoundFeeAddressUpdated(address roundFeeAddress)
func (_Round *RoundFilterer) ParseRoundFeeAddressUpdated(log types.Log) (*RoundRoundFeeAddressUpdated, error) {
	event := new(RoundRoundFeeAddressUpdated)
	if err := _Round.contract.UnpackLog(event, "RoundFeeAddressUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoundRoundFeePercentageUpdatedIterator is returned from FilterRoundFeePercentageUpdated and is used to iterate over the raw logs and unpacked data for RoundFeePercentageUpdated events raised by the Round contract.
type RoundRoundFeePercentageUpdatedIterator struct {
	Event *RoundRoundFeePercentageUpdated // Event containing the contract specifics and raw log

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
func (it *RoundRoundFeePercentageUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoundRoundFeePercentageUpdated)
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
		it.Event = new(RoundRoundFeePercentageUpdated)
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
func (it *RoundRoundFeePercentageUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoundRoundFeePercentageUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoundRoundFeePercentageUpdated represents a RoundFeePercentageUpdated event raised by the Round contract.
type RoundRoundFeePercentageUpdated struct {
	RoundFeePercentage uint32
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterRoundFeePercentageUpdated is a free log retrieval operation binding the contract event 0x94ef22170df1c9a0e250ce2eae5fd3abc6d1379de9f9855fee7e490f15ddda87.
//
// Solidity: event RoundFeePercentageUpdated(uint32 roundFeePercentage)
func (_Round *RoundFilterer) FilterRoundFeePercentageUpdated(opts *bind.FilterOpts) (*RoundRoundFeePercentageUpdatedIterator, error) {

	logs, sub, err := _Round.contract.FilterLogs(opts, "RoundFeePercentageUpdated")
	if err != nil {
		return nil, err
	}
	return &RoundRoundFeePercentageUpdatedIterator{contract: _Round.contract, event: "RoundFeePercentageUpdated", logs: logs, sub: sub}, nil
}

// WatchRoundFeePercentageUpdated is a free log subscription operation binding the contract event 0x94ef22170df1c9a0e250ce2eae5fd3abc6d1379de9f9855fee7e490f15ddda87.
//
// Solidity: event RoundFeePercentageUpdated(uint32 roundFeePercentage)
func (_Round *RoundFilterer) WatchRoundFeePercentageUpdated(opts *bind.WatchOpts, sink chan<- *RoundRoundFeePercentageUpdated) (event.Subscription, error) {

	logs, sub, err := _Round.contract.WatchLogs(opts, "RoundFeePercentageUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoundRoundFeePercentageUpdated)
				if err := _Round.contract.UnpackLog(event, "RoundFeePercentageUpdated", log); err != nil {
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

// ParseRoundFeePercentageUpdated is a log parse operation binding the contract event 0x94ef22170df1c9a0e250ce2eae5fd3abc6d1379de9f9855fee7e490f15ddda87.
//
// Solidity: event RoundFeePercentageUpdated(uint32 roundFeePercentage)
func (_Round *RoundFilterer) ParseRoundFeePercentageUpdated(log types.Log) (*RoundRoundFeePercentageUpdated, error) {
	event := new(RoundRoundFeePercentageUpdated)
	if err := _Round.contract.UnpackLog(event, "RoundFeePercentageUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoundRoundMetaPtrUpdatedIterator is returned from FilterRoundMetaPtrUpdated and is used to iterate over the raw logs and unpacked data for RoundMetaPtrUpdated events raised by the Round contract.
type RoundRoundMetaPtrUpdatedIterator struct {
	Event *RoundRoundMetaPtrUpdated // Event containing the contract specifics and raw log

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
func (it *RoundRoundMetaPtrUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoundRoundMetaPtrUpdated)
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
		it.Event = new(RoundRoundMetaPtrUpdated)
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
func (it *RoundRoundMetaPtrUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoundRoundMetaPtrUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoundRoundMetaPtrUpdated represents a RoundMetaPtrUpdated event raised by the Round contract.
type RoundRoundMetaPtrUpdated struct {
	OldMetaPtr MetaPtr
	NewMetaPtr MetaPtr
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRoundMetaPtrUpdated is a free log retrieval operation binding the contract event 0xb625c4faee7533bf09da813b9c0e169f98d36c22d9ce5181c9e4d7577556760a.
//
// Solidity: event RoundMetaPtrUpdated((uint256,string) oldMetaPtr, (uint256,string) newMetaPtr)
func (_Round *RoundFilterer) FilterRoundMetaPtrUpdated(opts *bind.FilterOpts) (*RoundRoundMetaPtrUpdatedIterator, error) {

	logs, sub, err := _Round.contract.FilterLogs(opts, "RoundMetaPtrUpdated")
	if err != nil {
		return nil, err
	}
	return &RoundRoundMetaPtrUpdatedIterator{contract: _Round.contract, event: "RoundMetaPtrUpdated", logs: logs, sub: sub}, nil
}

// WatchRoundMetaPtrUpdated is a free log subscription operation binding the contract event 0xb625c4faee7533bf09da813b9c0e169f98d36c22d9ce5181c9e4d7577556760a.
//
// Solidity: event RoundMetaPtrUpdated((uint256,string) oldMetaPtr, (uint256,string) newMetaPtr)
func (_Round *RoundFilterer) WatchRoundMetaPtrUpdated(opts *bind.WatchOpts, sink chan<- *RoundRoundMetaPtrUpdated) (event.Subscription, error) {

	logs, sub, err := _Round.contract.WatchLogs(opts, "RoundMetaPtrUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoundRoundMetaPtrUpdated)
				if err := _Round.contract.UnpackLog(event, "RoundMetaPtrUpdated", log); err != nil {
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

// ParseRoundMetaPtrUpdated is a log parse operation binding the contract event 0xb625c4faee7533bf09da813b9c0e169f98d36c22d9ce5181c9e4d7577556760a.
//
// Solidity: event RoundMetaPtrUpdated((uint256,string) oldMetaPtr, (uint256,string) newMetaPtr)
func (_Round *RoundFilterer) ParseRoundMetaPtrUpdated(log types.Log) (*RoundRoundMetaPtrUpdated, error) {
	event := new(RoundRoundMetaPtrUpdated)
	if err := _Round.contract.UnpackLog(event, "RoundMetaPtrUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoundRoundStartTimeUpdatedIterator is returned from FilterRoundStartTimeUpdated and is used to iterate over the raw logs and unpacked data for RoundStartTimeUpdated events raised by the Round contract.
type RoundRoundStartTimeUpdatedIterator struct {
	Event *RoundRoundStartTimeUpdated // Event containing the contract specifics and raw log

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
func (it *RoundRoundStartTimeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoundRoundStartTimeUpdated)
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
		it.Event = new(RoundRoundStartTimeUpdated)
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
func (it *RoundRoundStartTimeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoundRoundStartTimeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoundRoundStartTimeUpdated represents a RoundStartTimeUpdated event raised by the Round contract.
type RoundRoundStartTimeUpdated struct {
	OldTime *big.Int
	NewTime *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoundStartTimeUpdated is a free log retrieval operation binding the contract event 0xbd905b0cf9e380e3eaf972d7cb542694f66361043fcce248341d90626d467476.
//
// Solidity: event RoundStartTimeUpdated(uint256 oldTime, uint256 newTime)
func (_Round *RoundFilterer) FilterRoundStartTimeUpdated(opts *bind.FilterOpts) (*RoundRoundStartTimeUpdatedIterator, error) {

	logs, sub, err := _Round.contract.FilterLogs(opts, "RoundStartTimeUpdated")
	if err != nil {
		return nil, err
	}
	return &RoundRoundStartTimeUpdatedIterator{contract: _Round.contract, event: "RoundStartTimeUpdated", logs: logs, sub: sub}, nil
}

// WatchRoundStartTimeUpdated is a free log subscription operation binding the contract event 0xbd905b0cf9e380e3eaf972d7cb542694f66361043fcce248341d90626d467476.
//
// Solidity: event RoundStartTimeUpdated(uint256 oldTime, uint256 newTime)
func (_Round *RoundFilterer) WatchRoundStartTimeUpdated(opts *bind.WatchOpts, sink chan<- *RoundRoundStartTimeUpdated) (event.Subscription, error) {

	logs, sub, err := _Round.contract.WatchLogs(opts, "RoundStartTimeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoundRoundStartTimeUpdated)
				if err := _Round.contract.UnpackLog(event, "RoundStartTimeUpdated", log); err != nil {
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

// ParseRoundStartTimeUpdated is a log parse operation binding the contract event 0xbd905b0cf9e380e3eaf972d7cb542694f66361043fcce248341d90626d467476.
//
// Solidity: event RoundStartTimeUpdated(uint256 oldTime, uint256 newTime)
func (_Round *RoundFilterer) ParseRoundStartTimeUpdated(log types.Log) (*RoundRoundStartTimeUpdated, error) {
	event := new(RoundRoundStartTimeUpdated)
	if err := _Round.contract.UnpackLog(event, "RoundStartTimeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
