// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package partyco

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

// StructsAddressAndAmount is an auto generated low-level Go binding around an user-defined struct.
type StructsAddressAndAmount struct {
	Addr   common.Address
	Amount *big.Int
}

// CollectionPartyMetaData contains all meta data concerning the CollectionParty contract.
var CollectionPartyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_partyDAOMultisig\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenVaultFactory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_weth\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_allowList\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"triggeredBy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"targetAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ethSpent\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ethFeePaid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalContributed\",\"type\":\"uint256\"}],\"name\":\"Bought\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contributor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalContributed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"excessContribution\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"}],\"name\":\"Claimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contributor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousTotalContributedToParty\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalFromContributor\",\"type\":\"uint256\"}],\"name\":\"Contributed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"triggeredBy\",\"type\":\"address\"}],\"name\":\"Expired\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"PARTY_TYPE\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allowList\",\"outputs\":[{\"internalType\":\"contractIAllowList\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_targetContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"}],\"name\":\"buy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contributor\",\"type\":\"address\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"claimed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contribute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"contributions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"previousTotalContributedToParty\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"}],\"name\":\"emergencyCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"_returnData\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyForceLost\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"emergencyWithdrawEth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"expire\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"expiresAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gatedToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gatedTokenAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contributor\",\"type\":\"address\"}],\"name\":\"getClaimAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_ethAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaximumContributions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxContributions\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaximumSpend\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxSpend\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_maxPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_secondsToTimeout\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_deciders\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structStructs.AddressAndAmount\",\"name\":\"_split\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structStructs.AddressAndAmount\",\"name\":\"_tokenGate\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isDecider\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nftContract\",\"outputs\":[{\"internalType\":\"contractIERC721Metadata\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"partyDAOMultisig\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"partyFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"partyStatus\",\"outputs\":[{\"internalType\":\"enumParty.PartyStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"splitBasisPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"splitRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenVault\",\"outputs\":[{\"internalType\":\"contractITokenVault\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenVaultFactory\",\"outputs\":[{\"internalType\":\"contractIERC721VaultFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"totalContributed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalContributedToParty\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contributor\",\"type\":\"address\"}],\"name\":\"totalEthUsed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_total\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSpent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"valueToTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokens\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"weth\",\"outputs\":[{\"internalType\":\"contractIWETH\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// CollectionPartyABI is the input ABI used to generate the binding from.
// Deprecated: Use CollectionPartyMetaData.ABI instead.
var CollectionPartyABI = CollectionPartyMetaData.ABI

// CollectionParty is an auto generated Go binding around an Ethereum contract.
type CollectionParty struct {
	CollectionPartyCaller     // Read-only binding to the contract
	CollectionPartyTransactor // Write-only binding to the contract
	CollectionPartyFilterer   // Log filterer for contract events
}

// CollectionPartyCaller is an auto generated read-only Go binding around an Ethereum contract.
type CollectionPartyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CollectionPartyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CollectionPartyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CollectionPartyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CollectionPartyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CollectionPartySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CollectionPartySession struct {
	Contract     *CollectionParty  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CollectionPartyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CollectionPartyCallerSession struct {
	Contract *CollectionPartyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// CollectionPartyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CollectionPartyTransactorSession struct {
	Contract     *CollectionPartyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// CollectionPartyRaw is an auto generated low-level Go binding around an Ethereum contract.
type CollectionPartyRaw struct {
	Contract *CollectionParty // Generic contract binding to access the raw methods on
}

// CollectionPartyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CollectionPartyCallerRaw struct {
	Contract *CollectionPartyCaller // Generic read-only contract binding to access the raw methods on
}

// CollectionPartyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CollectionPartyTransactorRaw struct {
	Contract *CollectionPartyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCollectionParty creates a new instance of CollectionParty, bound to a specific deployed contract.
func NewCollectionParty(address common.Address, backend bind.ContractBackend) (*CollectionParty, error) {
	contract, err := bindCollectionParty(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CollectionParty{CollectionPartyCaller: CollectionPartyCaller{contract: contract}, CollectionPartyTransactor: CollectionPartyTransactor{contract: contract}, CollectionPartyFilterer: CollectionPartyFilterer{contract: contract}}, nil
}

// NewCollectionPartyCaller creates a new read-only instance of CollectionParty, bound to a specific deployed contract.
func NewCollectionPartyCaller(address common.Address, caller bind.ContractCaller) (*CollectionPartyCaller, error) {
	contract, err := bindCollectionParty(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CollectionPartyCaller{contract: contract}, nil
}

// NewCollectionPartyTransactor creates a new write-only instance of CollectionParty, bound to a specific deployed contract.
func NewCollectionPartyTransactor(address common.Address, transactor bind.ContractTransactor) (*CollectionPartyTransactor, error) {
	contract, err := bindCollectionParty(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CollectionPartyTransactor{contract: contract}, nil
}

// NewCollectionPartyFilterer creates a new log filterer instance of CollectionParty, bound to a specific deployed contract.
func NewCollectionPartyFilterer(address common.Address, filterer bind.ContractFilterer) (*CollectionPartyFilterer, error) {
	contract, err := bindCollectionParty(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CollectionPartyFilterer{contract: contract}, nil
}

// bindCollectionParty binds a generic wrapper to an already deployed contract.
func bindCollectionParty(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CollectionPartyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CollectionParty *CollectionPartyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CollectionParty.Contract.CollectionPartyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CollectionParty *CollectionPartyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CollectionParty.Contract.CollectionPartyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CollectionParty *CollectionPartyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CollectionParty.Contract.CollectionPartyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CollectionParty *CollectionPartyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CollectionParty.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CollectionParty *CollectionPartyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CollectionParty.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CollectionParty *CollectionPartyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CollectionParty.Contract.contract.Transact(opts, method, params...)
}

// PARTYTYPE is a free data retrieval call binding the contract method 0xe54329d4.
//
// Solidity: function PARTY_TYPE() view returns(string)
func (_CollectionParty *CollectionPartyCaller) PARTYTYPE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CollectionParty.contract.Call(opts, &out, "PARTY_TYPE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// PARTYTYPE is a free data retrieval call binding the contract method 0xe54329d4.
//
// Solidity: function PARTY_TYPE() view returns(string)
func (_CollectionParty *CollectionPartySession) PARTYTYPE() (string, error) {
	return _CollectionParty.Contract.PARTYTYPE(&_CollectionParty.CallOpts)
}

// PARTYTYPE is a free data retrieval call binding the contract method 0xe54329d4.
//
// Solidity: function PARTY_TYPE() view returns(string)
func (_CollectionParty *CollectionPartyCallerSession) PARTYTYPE() (string, error) {
	return _CollectionParty.Contract.PARTYTYPE(&_CollectionParty.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint16)
func (_CollectionParty *CollectionPartyCaller) VERSION(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _CollectionParty.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint16)
func (_CollectionParty *CollectionPartySession) VERSION() (uint16, error) {
	return _CollectionParty.Contract.VERSION(&_CollectionParty.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint16)
func (_CollectionParty *CollectionPartyCallerSession) VERSION() (uint16, error) {
	return _CollectionParty.Contract.VERSION(&_CollectionParty.CallOpts)
}

// AllowList is a free data retrieval call binding the contract method 0x87b9d25c.
//
// Solidity: function allowList() view returns(address)
func (_CollectionParty *CollectionPartyCaller) AllowList(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CollectionParty.contract.Call(opts, &out, "allowList")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllowList is a free data retrieval call binding the contract method 0x87b9d25c.
//
// Solidity: function allowList() view returns(address)
func (_CollectionParty *CollectionPartySession) AllowList() (common.Address, error) {
	return _CollectionParty.Contract.AllowList(&_CollectionParty.CallOpts)
}

// AllowList is a free data retrieval call binding the contract method 0x87b9d25c.
//
// Solidity: function allowList() view returns(address)
func (_CollectionParty *CollectionPartyCallerSession) AllowList() (common.Address, error) {
	return _CollectionParty.Contract.AllowList(&_CollectionParty.CallOpts)
}

// Claimed is a free data retrieval call binding the contract method 0xc884ef83.
//
// Solidity: function claimed(address ) view returns(bool)
func (_CollectionParty *CollectionPartyCaller) Claimed(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _CollectionParty.contract.Call(opts, &out, "claimed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Claimed is a free data retrieval call binding the contract method 0xc884ef83.
//
// Solidity: function claimed(address ) view returns(bool)
func (_CollectionParty *CollectionPartySession) Claimed(arg0 common.Address) (bool, error) {
	return _CollectionParty.Contract.Claimed(&_CollectionParty.CallOpts, arg0)
}

// Claimed is a free data retrieval call binding the contract method 0xc884ef83.
//
// Solidity: function claimed(address ) view returns(bool)
func (_CollectionParty *CollectionPartyCallerSession) Claimed(arg0 common.Address) (bool, error) {
	return _CollectionParty.Contract.Claimed(&_CollectionParty.CallOpts, arg0)
}

// Contributions is a free data retrieval call binding the contract method 0x550b521c.
//
// Solidity: function contributions(address , uint256 ) view returns(uint256 amount, uint256 previousTotalContributedToParty)
func (_CollectionParty *CollectionPartyCaller) Contributions(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Amount                          *big.Int
	PreviousTotalContributedToParty *big.Int
}, error) {
	var out []interface{}
	err := _CollectionParty.contract.Call(opts, &out, "contributions", arg0, arg1)

	outstruct := new(struct {
		Amount                          *big.Int
		PreviousTotalContributedToParty *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.PreviousTotalContributedToParty = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Contributions is a free data retrieval call binding the contract method 0x550b521c.
//
// Solidity: function contributions(address , uint256 ) view returns(uint256 amount, uint256 previousTotalContributedToParty)
func (_CollectionParty *CollectionPartySession) Contributions(arg0 common.Address, arg1 *big.Int) (struct {
	Amount                          *big.Int
	PreviousTotalContributedToParty *big.Int
}, error) {
	return _CollectionParty.Contract.Contributions(&_CollectionParty.CallOpts, arg0, arg1)
}

// Contributions is a free data retrieval call binding the contract method 0x550b521c.
//
// Solidity: function contributions(address , uint256 ) view returns(uint256 amount, uint256 previousTotalContributedToParty)
func (_CollectionParty *CollectionPartyCallerSession) Contributions(arg0 common.Address, arg1 *big.Int) (struct {
	Amount                          *big.Int
	PreviousTotalContributedToParty *big.Int
}, error) {
	return _CollectionParty.Contract.Contributions(&_CollectionParty.CallOpts, arg0, arg1)
}

// ExpiresAt is a free data retrieval call binding the contract method 0x8622a689.
//
// Solidity: function expiresAt() view returns(uint256)
func (_CollectionParty *CollectionPartyCaller) ExpiresAt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CollectionParty.contract.Call(opts, &out, "expiresAt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExpiresAt is a free data retrieval call binding the contract method 0x8622a689.
//
// Solidity: function expiresAt() view returns(uint256)
func (_CollectionParty *CollectionPartySession) ExpiresAt() (*big.Int, error) {
	return _CollectionParty.Contract.ExpiresAt(&_CollectionParty.CallOpts)
}

// ExpiresAt is a free data retrieval call binding the contract method 0x8622a689.
//
// Solidity: function expiresAt() view returns(uint256)
func (_CollectionParty *CollectionPartyCallerSession) ExpiresAt() (*big.Int, error) {
	return _CollectionParty.Contract.ExpiresAt(&_CollectionParty.CallOpts)
}

// GatedToken is a free data retrieval call binding the contract method 0x7a2ba9c4.
//
// Solidity: function gatedToken() view returns(address)
func (_CollectionParty *CollectionPartyCaller) GatedToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CollectionParty.contract.Call(opts, &out, "gatedToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GatedToken is a free data retrieval call binding the contract method 0x7a2ba9c4.
//
// Solidity: function gatedToken() view returns(address)
func (_CollectionParty *CollectionPartySession) GatedToken() (common.Address, error) {
	return _CollectionParty.Contract.GatedToken(&_CollectionParty.CallOpts)
}

// GatedToken is a free data retrieval call binding the contract method 0x7a2ba9c4.
//
// Solidity: function gatedToken() view returns(address)
func (_CollectionParty *CollectionPartyCallerSession) GatedToken() (common.Address, error) {
	return _CollectionParty.Contract.GatedToken(&_CollectionParty.CallOpts)
}

// GatedTokenAmount is a free data retrieval call binding the contract method 0x8d42ecd6.
//
// Solidity: function gatedTokenAmount() view returns(uint256)
func (_CollectionParty *CollectionPartyCaller) GatedTokenAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CollectionParty.contract.Call(opts, &out, "gatedTokenAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GatedTokenAmount is a free data retrieval call binding the contract method 0x8d42ecd6.
//
// Solidity: function gatedTokenAmount() view returns(uint256)
func (_CollectionParty *CollectionPartySession) GatedTokenAmount() (*big.Int, error) {
	return _CollectionParty.Contract.GatedTokenAmount(&_CollectionParty.CallOpts)
}

// GatedTokenAmount is a free data retrieval call binding the contract method 0x8d42ecd6.
//
// Solidity: function gatedTokenAmount() view returns(uint256)
func (_CollectionParty *CollectionPartyCallerSession) GatedTokenAmount() (*big.Int, error) {
	return _CollectionParty.Contract.GatedTokenAmount(&_CollectionParty.CallOpts)
}

// GetClaimAmounts is a free data retrieval call binding the contract method 0x6971524f.
//
// Solidity: function getClaimAmounts(address _contributor) view returns(uint256 _tokenAmount, uint256 _ethAmount)
func (_CollectionParty *CollectionPartyCaller) GetClaimAmounts(opts *bind.CallOpts, _contributor common.Address) (struct {
	TokenAmount *big.Int
	EthAmount   *big.Int
}, error) {
	var out []interface{}
	err := _CollectionParty.contract.Call(opts, &out, "getClaimAmounts", _contributor)

	outstruct := new(struct {
		TokenAmount *big.Int
		EthAmount   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokenAmount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.EthAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetClaimAmounts is a free data retrieval call binding the contract method 0x6971524f.
//
// Solidity: function getClaimAmounts(address _contributor) view returns(uint256 _tokenAmount, uint256 _ethAmount)
func (_CollectionParty *CollectionPartySession) GetClaimAmounts(_contributor common.Address) (struct {
	TokenAmount *big.Int
	EthAmount   *big.Int
}, error) {
	return _CollectionParty.Contract.GetClaimAmounts(&_CollectionParty.CallOpts, _contributor)
}

// GetClaimAmounts is a free data retrieval call binding the contract method 0x6971524f.
//
// Solidity: function getClaimAmounts(address _contributor) view returns(uint256 _tokenAmount, uint256 _ethAmount)
func (_CollectionParty *CollectionPartyCallerSession) GetClaimAmounts(_contributor common.Address) (struct {
	TokenAmount *big.Int
	EthAmount   *big.Int
}, error) {
	return _CollectionParty.Contract.GetClaimAmounts(&_CollectionParty.CallOpts, _contributor)
}

// GetMaximumContributions is a free data retrieval call binding the contract method 0x9559da3b.
//
// Solidity: function getMaximumContributions() view returns(uint256 _maxContributions)
func (_CollectionParty *CollectionPartyCaller) GetMaximumContributions(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CollectionParty.contract.Call(opts, &out, "getMaximumContributions")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaximumContributions is a free data retrieval call binding the contract method 0x9559da3b.
//
// Solidity: function getMaximumContributions() view returns(uint256 _maxContributions)
func (_CollectionParty *CollectionPartySession) GetMaximumContributions() (*big.Int, error) {
	return _CollectionParty.Contract.GetMaximumContributions(&_CollectionParty.CallOpts)
}

// GetMaximumContributions is a free data retrieval call binding the contract method 0x9559da3b.
//
// Solidity: function getMaximumContributions() view returns(uint256 _maxContributions)
func (_CollectionParty *CollectionPartyCallerSession) GetMaximumContributions() (*big.Int, error) {
	return _CollectionParty.Contract.GetMaximumContributions(&_CollectionParty.CallOpts)
}

// GetMaximumSpend is a free data retrieval call binding the contract method 0xacd13c59.
//
// Solidity: function getMaximumSpend() view returns(uint256 _maxSpend)
func (_CollectionParty *CollectionPartyCaller) GetMaximumSpend(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CollectionParty.contract.Call(opts, &out, "getMaximumSpend")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaximumSpend is a free data retrieval call binding the contract method 0xacd13c59.
//
// Solidity: function getMaximumSpend() view returns(uint256 _maxSpend)
func (_CollectionParty *CollectionPartySession) GetMaximumSpend() (*big.Int, error) {
	return _CollectionParty.Contract.GetMaximumSpend(&_CollectionParty.CallOpts)
}

// GetMaximumSpend is a free data retrieval call binding the contract method 0xacd13c59.
//
// Solidity: function getMaximumSpend() view returns(uint256 _maxSpend)
func (_CollectionParty *CollectionPartyCallerSession) GetMaximumSpend() (*big.Int, error) {
	return _CollectionParty.Contract.GetMaximumSpend(&_CollectionParty.CallOpts)
}

// IsDecider is a free data retrieval call binding the contract method 0xceacca87.
//
// Solidity: function isDecider(address ) view returns(bool)
func (_CollectionParty *CollectionPartyCaller) IsDecider(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _CollectionParty.contract.Call(opts, &out, "isDecider", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDecider is a free data retrieval call binding the contract method 0xceacca87.
//
// Solidity: function isDecider(address ) view returns(bool)
func (_CollectionParty *CollectionPartySession) IsDecider(arg0 common.Address) (bool, error) {
	return _CollectionParty.Contract.IsDecider(&_CollectionParty.CallOpts, arg0)
}

// IsDecider is a free data retrieval call binding the contract method 0xceacca87.
//
// Solidity: function isDecider(address ) view returns(bool)
func (_CollectionParty *CollectionPartyCallerSession) IsDecider(arg0 common.Address) (bool, error) {
	return _CollectionParty.Contract.IsDecider(&_CollectionParty.CallOpts, arg0)
}

// MaxPrice is a free data retrieval call binding the contract method 0xe38d6b5c.
//
// Solidity: function maxPrice() view returns(uint256)
func (_CollectionParty *CollectionPartyCaller) MaxPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CollectionParty.contract.Call(opts, &out, "maxPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxPrice is a free data retrieval call binding the contract method 0xe38d6b5c.
//
// Solidity: function maxPrice() view returns(uint256)
func (_CollectionParty *CollectionPartySession) MaxPrice() (*big.Int, error) {
	return _CollectionParty.Contract.MaxPrice(&_CollectionParty.CallOpts)
}

// MaxPrice is a free data retrieval call binding the contract method 0xe38d6b5c.
//
// Solidity: function maxPrice() view returns(uint256)
func (_CollectionParty *CollectionPartyCallerSession) MaxPrice() (*big.Int, error) {
	return _CollectionParty.Contract.MaxPrice(&_CollectionParty.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CollectionParty *CollectionPartyCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CollectionParty.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CollectionParty *CollectionPartySession) Name() (string, error) {
	return _CollectionParty.Contract.Name(&_CollectionParty.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CollectionParty *CollectionPartyCallerSession) Name() (string, error) {
	return _CollectionParty.Contract.Name(&_CollectionParty.CallOpts)
}

// NftContract is a free data retrieval call binding the contract method 0xd56d229d.
//
// Solidity: function nftContract() view returns(address)
func (_CollectionParty *CollectionPartyCaller) NftContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CollectionParty.contract.Call(opts, &out, "nftContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NftContract is a free data retrieval call binding the contract method 0xd56d229d.
//
// Solidity: function nftContract() view returns(address)
func (_CollectionParty *CollectionPartySession) NftContract() (common.Address, error) {
	return _CollectionParty.Contract.NftContract(&_CollectionParty.CallOpts)
}

// NftContract is a free data retrieval call binding the contract method 0xd56d229d.
//
// Solidity: function nftContract() view returns(address)
func (_CollectionParty *CollectionPartyCallerSession) NftContract() (common.Address, error) {
	return _CollectionParty.Contract.NftContract(&_CollectionParty.CallOpts)
}

// PartyDAOMultisig is a free data retrieval call binding the contract method 0x3c4d12d9.
//
// Solidity: function partyDAOMultisig() view returns(address)
func (_CollectionParty *CollectionPartyCaller) PartyDAOMultisig(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CollectionParty.contract.Call(opts, &out, "partyDAOMultisig")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PartyDAOMultisig is a free data retrieval call binding the contract method 0x3c4d12d9.
//
// Solidity: function partyDAOMultisig() view returns(address)
func (_CollectionParty *CollectionPartySession) PartyDAOMultisig() (common.Address, error) {
	return _CollectionParty.Contract.PartyDAOMultisig(&_CollectionParty.CallOpts)
}

// PartyDAOMultisig is a free data retrieval call binding the contract method 0x3c4d12d9.
//
// Solidity: function partyDAOMultisig() view returns(address)
func (_CollectionParty *CollectionPartyCallerSession) PartyDAOMultisig() (common.Address, error) {
	return _CollectionParty.Contract.PartyDAOMultisig(&_CollectionParty.CallOpts)
}

// PartyFactory is a free data retrieval call binding the contract method 0x7ca67b62.
//
// Solidity: function partyFactory() view returns(address)
func (_CollectionParty *CollectionPartyCaller) PartyFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CollectionParty.contract.Call(opts, &out, "partyFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PartyFactory is a free data retrieval call binding the contract method 0x7ca67b62.
//
// Solidity: function partyFactory() view returns(address)
func (_CollectionParty *CollectionPartySession) PartyFactory() (common.Address, error) {
	return _CollectionParty.Contract.PartyFactory(&_CollectionParty.CallOpts)
}

// PartyFactory is a free data retrieval call binding the contract method 0x7ca67b62.
//
// Solidity: function partyFactory() view returns(address)
func (_CollectionParty *CollectionPartyCallerSession) PartyFactory() (common.Address, error) {
	return _CollectionParty.Contract.PartyFactory(&_CollectionParty.CallOpts)
}

// PartyStatus is a free data retrieval call binding the contract method 0xdf51c07f.
//
// Solidity: function partyStatus() view returns(uint8)
func (_CollectionParty *CollectionPartyCaller) PartyStatus(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _CollectionParty.contract.Call(opts, &out, "partyStatus")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// PartyStatus is a free data retrieval call binding the contract method 0xdf51c07f.
//
// Solidity: function partyStatus() view returns(uint8)
func (_CollectionParty *CollectionPartySession) PartyStatus() (uint8, error) {
	return _CollectionParty.Contract.PartyStatus(&_CollectionParty.CallOpts)
}

// PartyStatus is a free data retrieval call binding the contract method 0xdf51c07f.
//
// Solidity: function partyStatus() view returns(uint8)
func (_CollectionParty *CollectionPartyCallerSession) PartyStatus() (uint8, error) {
	return _CollectionParty.Contract.PartyStatus(&_CollectionParty.CallOpts)
}

// SplitBasisPoints is a free data retrieval call binding the contract method 0x2bbce5e6.
//
// Solidity: function splitBasisPoints() view returns(uint256)
func (_CollectionParty *CollectionPartyCaller) SplitBasisPoints(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CollectionParty.contract.Call(opts, &out, "splitBasisPoints")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SplitBasisPoints is a free data retrieval call binding the contract method 0x2bbce5e6.
//
// Solidity: function splitBasisPoints() view returns(uint256)
func (_CollectionParty *CollectionPartySession) SplitBasisPoints() (*big.Int, error) {
	return _CollectionParty.Contract.SplitBasisPoints(&_CollectionParty.CallOpts)
}

// SplitBasisPoints is a free data retrieval call binding the contract method 0x2bbce5e6.
//
// Solidity: function splitBasisPoints() view returns(uint256)
func (_CollectionParty *CollectionPartyCallerSession) SplitBasisPoints() (*big.Int, error) {
	return _CollectionParty.Contract.SplitBasisPoints(&_CollectionParty.CallOpts)
}

// SplitRecipient is a free data retrieval call binding the contract method 0x4367a029.
//
// Solidity: function splitRecipient() view returns(address)
func (_CollectionParty *CollectionPartyCaller) SplitRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CollectionParty.contract.Call(opts, &out, "splitRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SplitRecipient is a free data retrieval call binding the contract method 0x4367a029.
//
// Solidity: function splitRecipient() view returns(address)
func (_CollectionParty *CollectionPartySession) SplitRecipient() (common.Address, error) {
	return _CollectionParty.Contract.SplitRecipient(&_CollectionParty.CallOpts)
}

// SplitRecipient is a free data retrieval call binding the contract method 0x4367a029.
//
// Solidity: function splitRecipient() view returns(address)
func (_CollectionParty *CollectionPartyCallerSession) SplitRecipient() (common.Address, error) {
	return _CollectionParty.Contract.SplitRecipient(&_CollectionParty.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CollectionParty *CollectionPartyCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CollectionParty.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CollectionParty *CollectionPartySession) Symbol() (string, error) {
	return _CollectionParty.Contract.Symbol(&_CollectionParty.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CollectionParty *CollectionPartyCallerSession) Symbol() (string, error) {
	return _CollectionParty.Contract.Symbol(&_CollectionParty.CallOpts)
}

// TokenId is a free data retrieval call binding the contract method 0x17d70f7c.
//
// Solidity: function tokenId() view returns(uint256)
func (_CollectionParty *CollectionPartyCaller) TokenId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CollectionParty.contract.Call(opts, &out, "tokenId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenId is a free data retrieval call binding the contract method 0x17d70f7c.
//
// Solidity: function tokenId() view returns(uint256)
func (_CollectionParty *CollectionPartySession) TokenId() (*big.Int, error) {
	return _CollectionParty.Contract.TokenId(&_CollectionParty.CallOpts)
}

// TokenId is a free data retrieval call binding the contract method 0x17d70f7c.
//
// Solidity: function tokenId() view returns(uint256)
func (_CollectionParty *CollectionPartyCallerSession) TokenId() (*big.Int, error) {
	return _CollectionParty.Contract.TokenId(&_CollectionParty.CallOpts)
}

// TokenVault is a free data retrieval call binding the contract method 0x5bc789d9.
//
// Solidity: function tokenVault() view returns(address)
func (_CollectionParty *CollectionPartyCaller) TokenVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CollectionParty.contract.Call(opts, &out, "tokenVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenVault is a free data retrieval call binding the contract method 0x5bc789d9.
//
// Solidity: function tokenVault() view returns(address)
func (_CollectionParty *CollectionPartySession) TokenVault() (common.Address, error) {
	return _CollectionParty.Contract.TokenVault(&_CollectionParty.CallOpts)
}

// TokenVault is a free data retrieval call binding the contract method 0x5bc789d9.
//
// Solidity: function tokenVault() view returns(address)
func (_CollectionParty *CollectionPartyCallerSession) TokenVault() (common.Address, error) {
	return _CollectionParty.Contract.TokenVault(&_CollectionParty.CallOpts)
}

// TokenVaultFactory is a free data retrieval call binding the contract method 0x0b203023.
//
// Solidity: function tokenVaultFactory() view returns(address)
func (_CollectionParty *CollectionPartyCaller) TokenVaultFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CollectionParty.contract.Call(opts, &out, "tokenVaultFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenVaultFactory is a free data retrieval call binding the contract method 0x0b203023.
//
// Solidity: function tokenVaultFactory() view returns(address)
func (_CollectionParty *CollectionPartySession) TokenVaultFactory() (common.Address, error) {
	return _CollectionParty.Contract.TokenVaultFactory(&_CollectionParty.CallOpts)
}

// TokenVaultFactory is a free data retrieval call binding the contract method 0x0b203023.
//
// Solidity: function tokenVaultFactory() view returns(address)
func (_CollectionParty *CollectionPartyCallerSession) TokenVaultFactory() (common.Address, error) {
	return _CollectionParty.Contract.TokenVaultFactory(&_CollectionParty.CallOpts)
}

// TotalContributed is a free data retrieval call binding the contract method 0xa0f243b8.
//
// Solidity: function totalContributed(address ) view returns(uint256)
func (_CollectionParty *CollectionPartyCaller) TotalContributed(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CollectionParty.contract.Call(opts, &out, "totalContributed", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalContributed is a free data retrieval call binding the contract method 0xa0f243b8.
//
// Solidity: function totalContributed(address ) view returns(uint256)
func (_CollectionParty *CollectionPartySession) TotalContributed(arg0 common.Address) (*big.Int, error) {
	return _CollectionParty.Contract.TotalContributed(&_CollectionParty.CallOpts, arg0)
}

// TotalContributed is a free data retrieval call binding the contract method 0xa0f243b8.
//
// Solidity: function totalContributed(address ) view returns(uint256)
func (_CollectionParty *CollectionPartyCallerSession) TotalContributed(arg0 common.Address) (*big.Int, error) {
	return _CollectionParty.Contract.TotalContributed(&_CollectionParty.CallOpts, arg0)
}

// TotalContributedToParty is a free data retrieval call binding the contract method 0x82a5c69a.
//
// Solidity: function totalContributedToParty() view returns(uint256)
func (_CollectionParty *CollectionPartyCaller) TotalContributedToParty(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CollectionParty.contract.Call(opts, &out, "totalContributedToParty")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalContributedToParty is a free data retrieval call binding the contract method 0x82a5c69a.
//
// Solidity: function totalContributedToParty() view returns(uint256)
func (_CollectionParty *CollectionPartySession) TotalContributedToParty() (*big.Int, error) {
	return _CollectionParty.Contract.TotalContributedToParty(&_CollectionParty.CallOpts)
}

// TotalContributedToParty is a free data retrieval call binding the contract method 0x82a5c69a.
//
// Solidity: function totalContributedToParty() view returns(uint256)
func (_CollectionParty *CollectionPartyCallerSession) TotalContributedToParty() (*big.Int, error) {
	return _CollectionParty.Contract.TotalContributedToParty(&_CollectionParty.CallOpts)
}

// TotalEthUsed is a free data retrieval call binding the contract method 0x25b42a26.
//
// Solidity: function totalEthUsed(address _contributor) view returns(uint256 _total)
func (_CollectionParty *CollectionPartyCaller) TotalEthUsed(opts *bind.CallOpts, _contributor common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CollectionParty.contract.Call(opts, &out, "totalEthUsed", _contributor)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalEthUsed is a free data retrieval call binding the contract method 0x25b42a26.
//
// Solidity: function totalEthUsed(address _contributor) view returns(uint256 _total)
func (_CollectionParty *CollectionPartySession) TotalEthUsed(_contributor common.Address) (*big.Int, error) {
	return _CollectionParty.Contract.TotalEthUsed(&_CollectionParty.CallOpts, _contributor)
}

// TotalEthUsed is a free data retrieval call binding the contract method 0x25b42a26.
//
// Solidity: function totalEthUsed(address _contributor) view returns(uint256 _total)
func (_CollectionParty *CollectionPartyCallerSession) TotalEthUsed(_contributor common.Address) (*big.Int, error) {
	return _CollectionParty.Contract.TotalEthUsed(&_CollectionParty.CallOpts, _contributor)
}

// TotalSpent is a free data retrieval call binding the contract method 0xfb346eab.
//
// Solidity: function totalSpent() view returns(uint256)
func (_CollectionParty *CollectionPartyCaller) TotalSpent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CollectionParty.contract.Call(opts, &out, "totalSpent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSpent is a free data retrieval call binding the contract method 0xfb346eab.
//
// Solidity: function totalSpent() view returns(uint256)
func (_CollectionParty *CollectionPartySession) TotalSpent() (*big.Int, error) {
	return _CollectionParty.Contract.TotalSpent(&_CollectionParty.CallOpts)
}

// TotalSpent is a free data retrieval call binding the contract method 0xfb346eab.
//
// Solidity: function totalSpent() view returns(uint256)
func (_CollectionParty *CollectionPartyCallerSession) TotalSpent() (*big.Int, error) {
	return _CollectionParty.Contract.TotalSpent(&_CollectionParty.CallOpts)
}

// ValueToTokens is a free data retrieval call binding the contract method 0x9744b8dc.
//
// Solidity: function valueToTokens(uint256 _value) pure returns(uint256 _tokens)
func (_CollectionParty *CollectionPartyCaller) ValueToTokens(opts *bind.CallOpts, _value *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CollectionParty.contract.Call(opts, &out, "valueToTokens", _value)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValueToTokens is a free data retrieval call binding the contract method 0x9744b8dc.
//
// Solidity: function valueToTokens(uint256 _value) pure returns(uint256 _tokens)
func (_CollectionParty *CollectionPartySession) ValueToTokens(_value *big.Int) (*big.Int, error) {
	return _CollectionParty.Contract.ValueToTokens(&_CollectionParty.CallOpts, _value)
}

// ValueToTokens is a free data retrieval call binding the contract method 0x9744b8dc.
//
// Solidity: function valueToTokens(uint256 _value) pure returns(uint256 _tokens)
func (_CollectionParty *CollectionPartyCallerSession) ValueToTokens(_value *big.Int) (*big.Int, error) {
	return _CollectionParty.Contract.ValueToTokens(&_CollectionParty.CallOpts, _value)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_CollectionParty *CollectionPartyCaller) Weth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CollectionParty.contract.Call(opts, &out, "weth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_CollectionParty *CollectionPartySession) Weth() (common.Address, error) {
	return _CollectionParty.Contract.Weth(&_CollectionParty.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_CollectionParty *CollectionPartyCallerSession) Weth() (common.Address, error) {
	return _CollectionParty.Contract.Weth(&_CollectionParty.CallOpts)
}

// Buy is a paid mutator transaction binding the contract method 0x9000ff09.
//
// Solidity: function buy(uint256 _tokenId, uint256 _value, address _targetContract, bytes _calldata) returns()
func (_CollectionParty *CollectionPartyTransactor) Buy(opts *bind.TransactOpts, _tokenId *big.Int, _value *big.Int, _targetContract common.Address, _calldata []byte) (*types.Transaction, error) {
	return _CollectionParty.contract.Transact(opts, "buy", _tokenId, _value, _targetContract, _calldata)
}

// Buy is a paid mutator transaction binding the contract method 0x9000ff09.
//
// Solidity: function buy(uint256 _tokenId, uint256 _value, address _targetContract, bytes _calldata) returns()
func (_CollectionParty *CollectionPartySession) Buy(_tokenId *big.Int, _value *big.Int, _targetContract common.Address, _calldata []byte) (*types.Transaction, error) {
	return _CollectionParty.Contract.Buy(&_CollectionParty.TransactOpts, _tokenId, _value, _targetContract, _calldata)
}

// Buy is a paid mutator transaction binding the contract method 0x9000ff09.
//
// Solidity: function buy(uint256 _tokenId, uint256 _value, address _targetContract, bytes _calldata) returns()
func (_CollectionParty *CollectionPartyTransactorSession) Buy(_tokenId *big.Int, _value *big.Int, _targetContract common.Address, _calldata []byte) (*types.Transaction, error) {
	return _CollectionParty.Contract.Buy(&_CollectionParty.TransactOpts, _tokenId, _value, _targetContract, _calldata)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address _contributor) returns()
func (_CollectionParty *CollectionPartyTransactor) Claim(opts *bind.TransactOpts, _contributor common.Address) (*types.Transaction, error) {
	return _CollectionParty.contract.Transact(opts, "claim", _contributor)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address _contributor) returns()
func (_CollectionParty *CollectionPartySession) Claim(_contributor common.Address) (*types.Transaction, error) {
	return _CollectionParty.Contract.Claim(&_CollectionParty.TransactOpts, _contributor)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address _contributor) returns()
func (_CollectionParty *CollectionPartyTransactorSession) Claim(_contributor common.Address) (*types.Transaction, error) {
	return _CollectionParty.Contract.Claim(&_CollectionParty.TransactOpts, _contributor)
}

// Contribute is a paid mutator transaction binding the contract method 0xd7bb99ba.
//
// Solidity: function contribute() payable returns()
func (_CollectionParty *CollectionPartyTransactor) Contribute(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CollectionParty.contract.Transact(opts, "contribute")
}

// Contribute is a paid mutator transaction binding the contract method 0xd7bb99ba.
//
// Solidity: function contribute() payable returns()
func (_CollectionParty *CollectionPartySession) Contribute() (*types.Transaction, error) {
	return _CollectionParty.Contract.Contribute(&_CollectionParty.TransactOpts)
}

// Contribute is a paid mutator transaction binding the contract method 0xd7bb99ba.
//
// Solidity: function contribute() payable returns()
func (_CollectionParty *CollectionPartyTransactorSession) Contribute() (*types.Transaction, error) {
	return _CollectionParty.Contract.Contribute(&_CollectionParty.TransactOpts)
}

// EmergencyCall is a paid mutator transaction binding the contract method 0xc4bf0220.
//
// Solidity: function emergencyCall(address _contract, bytes _calldata) returns(bool _success, bytes _returnData)
func (_CollectionParty *CollectionPartyTransactor) EmergencyCall(opts *bind.TransactOpts, _contract common.Address, _calldata []byte) (*types.Transaction, error) {
	return _CollectionParty.contract.Transact(opts, "emergencyCall", _contract, _calldata)
}

// EmergencyCall is a paid mutator transaction binding the contract method 0xc4bf0220.
//
// Solidity: function emergencyCall(address _contract, bytes _calldata) returns(bool _success, bytes _returnData)
func (_CollectionParty *CollectionPartySession) EmergencyCall(_contract common.Address, _calldata []byte) (*types.Transaction, error) {
	return _CollectionParty.Contract.EmergencyCall(&_CollectionParty.TransactOpts, _contract, _calldata)
}

// EmergencyCall is a paid mutator transaction binding the contract method 0xc4bf0220.
//
// Solidity: function emergencyCall(address _contract, bytes _calldata) returns(bool _success, bytes _returnData)
func (_CollectionParty *CollectionPartyTransactorSession) EmergencyCall(_contract common.Address, _calldata []byte) (*types.Transaction, error) {
	return _CollectionParty.Contract.EmergencyCall(&_CollectionParty.TransactOpts, _contract, _calldata)
}

// EmergencyForceLost is a paid mutator transaction binding the contract method 0x17821fdc.
//
// Solidity: function emergencyForceLost() returns()
func (_CollectionParty *CollectionPartyTransactor) EmergencyForceLost(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CollectionParty.contract.Transact(opts, "emergencyForceLost")
}

// EmergencyForceLost is a paid mutator transaction binding the contract method 0x17821fdc.
//
// Solidity: function emergencyForceLost() returns()
func (_CollectionParty *CollectionPartySession) EmergencyForceLost() (*types.Transaction, error) {
	return _CollectionParty.Contract.EmergencyForceLost(&_CollectionParty.TransactOpts)
}

// EmergencyForceLost is a paid mutator transaction binding the contract method 0x17821fdc.
//
// Solidity: function emergencyForceLost() returns()
func (_CollectionParty *CollectionPartyTransactorSession) EmergencyForceLost() (*types.Transaction, error) {
	return _CollectionParty.Contract.EmergencyForceLost(&_CollectionParty.TransactOpts)
}

// EmergencyWithdrawEth is a paid mutator transaction binding the contract method 0x429093cc.
//
// Solidity: function emergencyWithdrawEth(uint256 _value) returns()
func (_CollectionParty *CollectionPartyTransactor) EmergencyWithdrawEth(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _CollectionParty.contract.Transact(opts, "emergencyWithdrawEth", _value)
}

// EmergencyWithdrawEth is a paid mutator transaction binding the contract method 0x429093cc.
//
// Solidity: function emergencyWithdrawEth(uint256 _value) returns()
func (_CollectionParty *CollectionPartySession) EmergencyWithdrawEth(_value *big.Int) (*types.Transaction, error) {
	return _CollectionParty.Contract.EmergencyWithdrawEth(&_CollectionParty.TransactOpts, _value)
}

// EmergencyWithdrawEth is a paid mutator transaction binding the contract method 0x429093cc.
//
// Solidity: function emergencyWithdrawEth(uint256 _value) returns()
func (_CollectionParty *CollectionPartyTransactorSession) EmergencyWithdrawEth(_value *big.Int) (*types.Transaction, error) {
	return _CollectionParty.Contract.EmergencyWithdrawEth(&_CollectionParty.TransactOpts, _value)
}

// Expire is a paid mutator transaction binding the contract method 0x79599f96.
//
// Solidity: function expire() returns()
func (_CollectionParty *CollectionPartyTransactor) Expire(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CollectionParty.contract.Transact(opts, "expire")
}

// Expire is a paid mutator transaction binding the contract method 0x79599f96.
//
// Solidity: function expire() returns()
func (_CollectionParty *CollectionPartySession) Expire() (*types.Transaction, error) {
	return _CollectionParty.Contract.Expire(&_CollectionParty.TransactOpts)
}

// Expire is a paid mutator transaction binding the contract method 0x79599f96.
//
// Solidity: function expire() returns()
func (_CollectionParty *CollectionPartyTransactorSession) Expire() (*types.Transaction, error) {
	return _CollectionParty.Contract.Expire(&_CollectionParty.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x414bd4b2.
//
// Solidity: function initialize(address _nftContract, uint256 _maxPrice, uint256 _secondsToTimeout, address[] _deciders, (address,uint256) _split, (address,uint256) _tokenGate, string _name, string _symbol) returns()
func (_CollectionParty *CollectionPartyTransactor) Initialize(opts *bind.TransactOpts, _nftContract common.Address, _maxPrice *big.Int, _secondsToTimeout *big.Int, _deciders []common.Address, _split StructsAddressAndAmount, _tokenGate StructsAddressAndAmount, _name string, _symbol string) (*types.Transaction, error) {
	return _CollectionParty.contract.Transact(opts, "initialize", _nftContract, _maxPrice, _secondsToTimeout, _deciders, _split, _tokenGate, _name, _symbol)
}

// Initialize is a paid mutator transaction binding the contract method 0x414bd4b2.
//
// Solidity: function initialize(address _nftContract, uint256 _maxPrice, uint256 _secondsToTimeout, address[] _deciders, (address,uint256) _split, (address,uint256) _tokenGate, string _name, string _symbol) returns()
func (_CollectionParty *CollectionPartySession) Initialize(_nftContract common.Address, _maxPrice *big.Int, _secondsToTimeout *big.Int, _deciders []common.Address, _split StructsAddressAndAmount, _tokenGate StructsAddressAndAmount, _name string, _symbol string) (*types.Transaction, error) {
	return _CollectionParty.Contract.Initialize(&_CollectionParty.TransactOpts, _nftContract, _maxPrice, _secondsToTimeout, _deciders, _split, _tokenGate, _name, _symbol)
}

// Initialize is a paid mutator transaction binding the contract method 0x414bd4b2.
//
// Solidity: function initialize(address _nftContract, uint256 _maxPrice, uint256 _secondsToTimeout, address[] _deciders, (address,uint256) _split, (address,uint256) _tokenGate, string _name, string _symbol) returns()
func (_CollectionParty *CollectionPartyTransactorSession) Initialize(_nftContract common.Address, _maxPrice *big.Int, _secondsToTimeout *big.Int, _deciders []common.Address, _split StructsAddressAndAmount, _tokenGate StructsAddressAndAmount, _name string, _symbol string) (*types.Transaction, error) {
	return _CollectionParty.Contract.Initialize(&_CollectionParty.TransactOpts, _nftContract, _maxPrice, _secondsToTimeout, _deciders, _split, _tokenGate, _name, _symbol)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_CollectionParty *CollectionPartyTransactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _CollectionParty.contract.Transact(opts, "onERC721Received", arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_CollectionParty *CollectionPartySession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _CollectionParty.Contract.OnERC721Received(&_CollectionParty.TransactOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_CollectionParty *CollectionPartyTransactorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _CollectionParty.Contract.OnERC721Received(&_CollectionParty.TransactOpts, arg0, arg1, arg2, arg3)
}

// CollectionPartyBoughtIterator is returned from FilterBought and is used to iterate over the raw logs and unpacked data for Bought events raised by the CollectionParty contract.
type CollectionPartyBoughtIterator struct {
	Event *CollectionPartyBought // Event containing the contract specifics and raw log

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
func (it *CollectionPartyBoughtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollectionPartyBought)
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
		it.Event = new(CollectionPartyBought)
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
func (it *CollectionPartyBoughtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollectionPartyBoughtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollectionPartyBought represents a Bought event raised by the CollectionParty contract.
type CollectionPartyBought struct {
	TokenId          *big.Int
	TriggeredBy      common.Address
	TargetAddress    common.Address
	EthSpent         *big.Int
	EthFeePaid       *big.Int
	TotalContributed *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterBought is a free log retrieval operation binding the contract event 0xa101afbd2f2109dcfa253f66e003d0f14243f2f21b3d1c668ef016cd3648ca97.
//
// Solidity: event Bought(uint256 tokenId, address triggeredBy, address targetAddress, uint256 ethSpent, uint256 ethFeePaid, uint256 totalContributed)
func (_CollectionParty *CollectionPartyFilterer) FilterBought(opts *bind.FilterOpts) (*CollectionPartyBoughtIterator, error) {

	logs, sub, err := _CollectionParty.contract.FilterLogs(opts, "Bought")
	if err != nil {
		return nil, err
	}
	return &CollectionPartyBoughtIterator{contract: _CollectionParty.contract, event: "Bought", logs: logs, sub: sub}, nil
}

// WatchBought is a free log subscription operation binding the contract event 0xa101afbd2f2109dcfa253f66e003d0f14243f2f21b3d1c668ef016cd3648ca97.
//
// Solidity: event Bought(uint256 tokenId, address triggeredBy, address targetAddress, uint256 ethSpent, uint256 ethFeePaid, uint256 totalContributed)
func (_CollectionParty *CollectionPartyFilterer) WatchBought(opts *bind.WatchOpts, sink chan<- *CollectionPartyBought) (event.Subscription, error) {

	logs, sub, err := _CollectionParty.contract.WatchLogs(opts, "Bought")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollectionPartyBought)
				if err := _CollectionParty.contract.UnpackLog(event, "Bought", log); err != nil {
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

// ParseBought is a log parse operation binding the contract event 0xa101afbd2f2109dcfa253f66e003d0f14243f2f21b3d1c668ef016cd3648ca97.
//
// Solidity: event Bought(uint256 tokenId, address triggeredBy, address targetAddress, uint256 ethSpent, uint256 ethFeePaid, uint256 totalContributed)
func (_CollectionParty *CollectionPartyFilterer) ParseBought(log types.Log) (*CollectionPartyBought, error) {
	event := new(CollectionPartyBought)
	if err := _CollectionParty.contract.UnpackLog(event, "Bought", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CollectionPartyClaimedIterator is returned from FilterClaimed and is used to iterate over the raw logs and unpacked data for Claimed events raised by the CollectionParty contract.
type CollectionPartyClaimedIterator struct {
	Event *CollectionPartyClaimed // Event containing the contract specifics and raw log

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
func (it *CollectionPartyClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollectionPartyClaimed)
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
		it.Event = new(CollectionPartyClaimed)
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
func (it *CollectionPartyClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollectionPartyClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollectionPartyClaimed represents a Claimed event raised by the CollectionParty contract.
type CollectionPartyClaimed struct {
	Contributor        common.Address
	TotalContributed   *big.Int
	ExcessContribution *big.Int
	TokenAmount        *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterClaimed is a free log retrieval operation binding the contract event 0x9cdcf2f7714cca3508c7f0110b04a90a80a3a8dd0e35de99689db74d28c5383e.
//
// Solidity: event Claimed(address indexed contributor, uint256 totalContributed, uint256 excessContribution, uint256 tokenAmount)
func (_CollectionParty *CollectionPartyFilterer) FilterClaimed(opts *bind.FilterOpts, contributor []common.Address) (*CollectionPartyClaimedIterator, error) {

	var contributorRule []interface{}
	for _, contributorItem := range contributor {
		contributorRule = append(contributorRule, contributorItem)
	}

	logs, sub, err := _CollectionParty.contract.FilterLogs(opts, "Claimed", contributorRule)
	if err != nil {
		return nil, err
	}
	return &CollectionPartyClaimedIterator{contract: _CollectionParty.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0x9cdcf2f7714cca3508c7f0110b04a90a80a3a8dd0e35de99689db74d28c5383e.
//
// Solidity: event Claimed(address indexed contributor, uint256 totalContributed, uint256 excessContribution, uint256 tokenAmount)
func (_CollectionParty *CollectionPartyFilterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *CollectionPartyClaimed, contributor []common.Address) (event.Subscription, error) {

	var contributorRule []interface{}
	for _, contributorItem := range contributor {
		contributorRule = append(contributorRule, contributorItem)
	}

	logs, sub, err := _CollectionParty.contract.WatchLogs(opts, "Claimed", contributorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollectionPartyClaimed)
				if err := _CollectionParty.contract.UnpackLog(event, "Claimed", log); err != nil {
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

// ParseClaimed is a log parse operation binding the contract event 0x9cdcf2f7714cca3508c7f0110b04a90a80a3a8dd0e35de99689db74d28c5383e.
//
// Solidity: event Claimed(address indexed contributor, uint256 totalContributed, uint256 excessContribution, uint256 tokenAmount)
func (_CollectionParty *CollectionPartyFilterer) ParseClaimed(log types.Log) (*CollectionPartyClaimed, error) {
	event := new(CollectionPartyClaimed)
	if err := _CollectionParty.contract.UnpackLog(event, "Claimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CollectionPartyContributedIterator is returned from FilterContributed and is used to iterate over the raw logs and unpacked data for Contributed events raised by the CollectionParty contract.
type CollectionPartyContributedIterator struct {
	Event *CollectionPartyContributed // Event containing the contract specifics and raw log

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
func (it *CollectionPartyContributedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollectionPartyContributed)
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
		it.Event = new(CollectionPartyContributed)
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
func (it *CollectionPartyContributedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollectionPartyContributedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollectionPartyContributed represents a Contributed event raised by the CollectionParty contract.
type CollectionPartyContributed struct {
	Contributor                     common.Address
	Amount                          *big.Int
	PreviousTotalContributedToParty *big.Int
	TotalFromContributor            *big.Int
	Raw                             types.Log // Blockchain specific contextual infos
}

// FilterContributed is a free log retrieval operation binding the contract event 0xb2623081601722547aae8781994e01a1974d95b0ad9ce6a0cfbe17487556257f.
//
// Solidity: event Contributed(address indexed contributor, uint256 amount, uint256 previousTotalContributedToParty, uint256 totalFromContributor)
func (_CollectionParty *CollectionPartyFilterer) FilterContributed(opts *bind.FilterOpts, contributor []common.Address) (*CollectionPartyContributedIterator, error) {

	var contributorRule []interface{}
	for _, contributorItem := range contributor {
		contributorRule = append(contributorRule, contributorItem)
	}

	logs, sub, err := _CollectionParty.contract.FilterLogs(opts, "Contributed", contributorRule)
	if err != nil {
		return nil, err
	}
	return &CollectionPartyContributedIterator{contract: _CollectionParty.contract, event: "Contributed", logs: logs, sub: sub}, nil
}

// WatchContributed is a free log subscription operation binding the contract event 0xb2623081601722547aae8781994e01a1974d95b0ad9ce6a0cfbe17487556257f.
//
// Solidity: event Contributed(address indexed contributor, uint256 amount, uint256 previousTotalContributedToParty, uint256 totalFromContributor)
func (_CollectionParty *CollectionPartyFilterer) WatchContributed(opts *bind.WatchOpts, sink chan<- *CollectionPartyContributed, contributor []common.Address) (event.Subscription, error) {

	var contributorRule []interface{}
	for _, contributorItem := range contributor {
		contributorRule = append(contributorRule, contributorItem)
	}

	logs, sub, err := _CollectionParty.contract.WatchLogs(opts, "Contributed", contributorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollectionPartyContributed)
				if err := _CollectionParty.contract.UnpackLog(event, "Contributed", log); err != nil {
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

// ParseContributed is a log parse operation binding the contract event 0xb2623081601722547aae8781994e01a1974d95b0ad9ce6a0cfbe17487556257f.
//
// Solidity: event Contributed(address indexed contributor, uint256 amount, uint256 previousTotalContributedToParty, uint256 totalFromContributor)
func (_CollectionParty *CollectionPartyFilterer) ParseContributed(log types.Log) (*CollectionPartyContributed, error) {
	event := new(CollectionPartyContributed)
	if err := _CollectionParty.contract.UnpackLog(event, "Contributed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CollectionPartyExpiredIterator is returned from FilterExpired and is used to iterate over the raw logs and unpacked data for Expired events raised by the CollectionParty contract.
type CollectionPartyExpiredIterator struct {
	Event *CollectionPartyExpired // Event containing the contract specifics and raw log

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
func (it *CollectionPartyExpiredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollectionPartyExpired)
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
		it.Event = new(CollectionPartyExpired)
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
func (it *CollectionPartyExpiredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollectionPartyExpiredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollectionPartyExpired represents a Expired event raised by the CollectionParty contract.
type CollectionPartyExpired struct {
	TriggeredBy common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterExpired is a free log retrieval operation binding the contract event 0xd5669ebe8b90ed693033f1915dcea031e7a16a6f146a8326c02eec10eed77edd.
//
// Solidity: event Expired(address triggeredBy)
func (_CollectionParty *CollectionPartyFilterer) FilterExpired(opts *bind.FilterOpts) (*CollectionPartyExpiredIterator, error) {

	logs, sub, err := _CollectionParty.contract.FilterLogs(opts, "Expired")
	if err != nil {
		return nil, err
	}
	return &CollectionPartyExpiredIterator{contract: _CollectionParty.contract, event: "Expired", logs: logs, sub: sub}, nil
}

// WatchExpired is a free log subscription operation binding the contract event 0xd5669ebe8b90ed693033f1915dcea031e7a16a6f146a8326c02eec10eed77edd.
//
// Solidity: event Expired(address triggeredBy)
func (_CollectionParty *CollectionPartyFilterer) WatchExpired(opts *bind.WatchOpts, sink chan<- *CollectionPartyExpired) (event.Subscription, error) {

	logs, sub, err := _CollectionParty.contract.WatchLogs(opts, "Expired")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollectionPartyExpired)
				if err := _CollectionParty.contract.UnpackLog(event, "Expired", log); err != nil {
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

// ParseExpired is a log parse operation binding the contract event 0xd5669ebe8b90ed693033f1915dcea031e7a16a6f146a8326c02eec10eed77edd.
//
// Solidity: event Expired(address triggeredBy)
func (_CollectionParty *CollectionPartyFilterer) ParseExpired(log types.Log) (*CollectionPartyExpired, error) {
	event := new(CollectionPartyExpired)
	if err := _CollectionParty.contract.UnpackLog(event, "Expired", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
