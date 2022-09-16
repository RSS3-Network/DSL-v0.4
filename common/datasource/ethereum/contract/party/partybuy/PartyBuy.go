// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package partybuy

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

// PartyBuyMetaData contains all meta data concerning the PartyBuy contract.
var PartyBuyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_partyDAOMultisig\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenVaultFactory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_weth\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_allowList\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"triggeredBy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"targetAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ethSpent\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ethFeePaid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalContributed\",\"type\":\"uint256\"}],\"name\":\"Bought\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contributor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalContributed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"excessContribution\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"}],\"name\":\"Claimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contributor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousTotalContributedToParty\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalFromContributor\",\"type\":\"uint256\"}],\"name\":\"Contributed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"triggeredBy\",\"type\":\"address\"}],\"name\":\"Expired\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allowList\",\"outputs\":[{\"internalType\":\"contractIAllowList\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_targetContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"}],\"name\":\"buy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contributor\",\"type\":\"address\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"claimed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contribute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"contributions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"previousTotalContributedToParty\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"}],\"name\":\"emergencyCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"_returnData\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyForceLost\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"emergencyWithdrawEth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"expire\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"expiresAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gatedToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gatedTokenAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contributor\",\"type\":\"address\"}],\"name\":\"getClaimAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_ethAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaximumContributions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxContributions\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaximumSpend\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxSpend\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_secondsToTimeout\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structStructs.AddressAndAmount\",\"name\":\"_split\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structStructs.AddressAndAmount\",\"name\":\"_tokenGate\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nftContract\",\"outputs\":[{\"internalType\":\"contractIERC721Metadata\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"partyDAOMultisig\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"partyStatus\",\"outputs\":[{\"internalType\":\"enumParty.PartyStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"splitBasisPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"splitRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenVault\",\"outputs\":[{\"internalType\":\"contractITokenVault\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenVaultFactory\",\"outputs\":[{\"internalType\":\"contractIERC721VaultFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"totalContributed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalContributedToParty\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contributor\",\"type\":\"address\"}],\"name\":\"totalEthUsed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_total\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSpent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"valueToTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokens\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"weth\",\"outputs\":[{\"internalType\":\"contractIWETH\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// PartyBuyABI is the input ABI used to generate the binding from.
// Deprecated: Use PartyBuyMetaData.ABI instead.
var PartyBuyABI = PartyBuyMetaData.ABI

// PartyBuy is an auto generated Go binding around an Ethereum contract.
type PartyBuy struct {
	PartyBuyCaller     // Read-only binding to the contract
	PartyBuyTransactor // Write-only binding to the contract
	PartyBuyFilterer   // Log filterer for contract events
}

// PartyBuyCaller is an auto generated read-only Go binding around an Ethereum contract.
type PartyBuyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PartyBuyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PartyBuyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PartyBuyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PartyBuyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PartyBuySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PartyBuySession struct {
	Contract     *PartyBuy         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PartyBuyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PartyBuyCallerSession struct {
	Contract *PartyBuyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// PartyBuyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PartyBuyTransactorSession struct {
	Contract     *PartyBuyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PartyBuyRaw is an auto generated low-level Go binding around an Ethereum contract.
type PartyBuyRaw struct {
	Contract *PartyBuy // Generic contract binding to access the raw methods on
}

// PartyBuyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PartyBuyCallerRaw struct {
	Contract *PartyBuyCaller // Generic read-only contract binding to access the raw methods on
}

// PartyBuyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PartyBuyTransactorRaw struct {
	Contract *PartyBuyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPartyBuy creates a new instance of PartyBuy, bound to a specific deployed contract.
func NewPartyBuy(address common.Address, backend bind.ContractBackend) (*PartyBuy, error) {
	contract, err := bindPartyBuy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PartyBuy{PartyBuyCaller: PartyBuyCaller{contract: contract}, PartyBuyTransactor: PartyBuyTransactor{contract: contract}, PartyBuyFilterer: PartyBuyFilterer{contract: contract}}, nil
}

// NewPartyBuyCaller creates a new read-only instance of PartyBuy, bound to a specific deployed contract.
func NewPartyBuyCaller(address common.Address, caller bind.ContractCaller) (*PartyBuyCaller, error) {
	contract, err := bindPartyBuy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PartyBuyCaller{contract: contract}, nil
}

// NewPartyBuyTransactor creates a new write-only instance of PartyBuy, bound to a specific deployed contract.
func NewPartyBuyTransactor(address common.Address, transactor bind.ContractTransactor) (*PartyBuyTransactor, error) {
	contract, err := bindPartyBuy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PartyBuyTransactor{contract: contract}, nil
}

// NewPartyBuyFilterer creates a new log filterer instance of PartyBuy, bound to a specific deployed contract.
func NewPartyBuyFilterer(address common.Address, filterer bind.ContractFilterer) (*PartyBuyFilterer, error) {
	contract, err := bindPartyBuy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PartyBuyFilterer{contract: contract}, nil
}

// bindPartyBuy binds a generic wrapper to an already deployed contract.
func bindPartyBuy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PartyBuyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PartyBuy *PartyBuyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PartyBuy.Contract.PartyBuyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PartyBuy *PartyBuyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PartyBuy.Contract.PartyBuyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PartyBuy *PartyBuyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PartyBuy.Contract.PartyBuyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PartyBuy *PartyBuyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PartyBuy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PartyBuy *PartyBuyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PartyBuy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PartyBuy *PartyBuyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PartyBuy.Contract.contract.Transact(opts, method, params...)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint16)
func (_PartyBuy *PartyBuyCaller) VERSION(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _PartyBuy.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint16)
func (_PartyBuy *PartyBuySession) VERSION() (uint16, error) {
	return _PartyBuy.Contract.VERSION(&_PartyBuy.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint16)
func (_PartyBuy *PartyBuyCallerSession) VERSION() (uint16, error) {
	return _PartyBuy.Contract.VERSION(&_PartyBuy.CallOpts)
}

// AllowList is a free data retrieval call binding the contract method 0x87b9d25c.
//
// Solidity: function allowList() view returns(address)
func (_PartyBuy *PartyBuyCaller) AllowList(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PartyBuy.contract.Call(opts, &out, "allowList")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllowList is a free data retrieval call binding the contract method 0x87b9d25c.
//
// Solidity: function allowList() view returns(address)
func (_PartyBuy *PartyBuySession) AllowList() (common.Address, error) {
	return _PartyBuy.Contract.AllowList(&_PartyBuy.CallOpts)
}

// AllowList is a free data retrieval call binding the contract method 0x87b9d25c.
//
// Solidity: function allowList() view returns(address)
func (_PartyBuy *PartyBuyCallerSession) AllowList() (common.Address, error) {
	return _PartyBuy.Contract.AllowList(&_PartyBuy.CallOpts)
}

// Claimed is a free data retrieval call binding the contract method 0xc884ef83.
//
// Solidity: function claimed(address ) view returns(bool)
func (_PartyBuy *PartyBuyCaller) Claimed(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _PartyBuy.contract.Call(opts, &out, "claimed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Claimed is a free data retrieval call binding the contract method 0xc884ef83.
//
// Solidity: function claimed(address ) view returns(bool)
func (_PartyBuy *PartyBuySession) Claimed(arg0 common.Address) (bool, error) {
	return _PartyBuy.Contract.Claimed(&_PartyBuy.CallOpts, arg0)
}

// Claimed is a free data retrieval call binding the contract method 0xc884ef83.
//
// Solidity: function claimed(address ) view returns(bool)
func (_PartyBuy *PartyBuyCallerSession) Claimed(arg0 common.Address) (bool, error) {
	return _PartyBuy.Contract.Claimed(&_PartyBuy.CallOpts, arg0)
}

// Contributions is a free data retrieval call binding the contract method 0x550b521c.
//
// Solidity: function contributions(address , uint256 ) view returns(uint256 amount, uint256 previousTotalContributedToParty)
func (_PartyBuy *PartyBuyCaller) Contributions(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Amount                          *big.Int
	PreviousTotalContributedToParty *big.Int
}, error) {
	var out []interface{}
	err := _PartyBuy.contract.Call(opts, &out, "contributions", arg0, arg1)

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
func (_PartyBuy *PartyBuySession) Contributions(arg0 common.Address, arg1 *big.Int) (struct {
	Amount                          *big.Int
	PreviousTotalContributedToParty *big.Int
}, error) {
	return _PartyBuy.Contract.Contributions(&_PartyBuy.CallOpts, arg0, arg1)
}

// Contributions is a free data retrieval call binding the contract method 0x550b521c.
//
// Solidity: function contributions(address , uint256 ) view returns(uint256 amount, uint256 previousTotalContributedToParty)
func (_PartyBuy *PartyBuyCallerSession) Contributions(arg0 common.Address, arg1 *big.Int) (struct {
	Amount                          *big.Int
	PreviousTotalContributedToParty *big.Int
}, error) {
	return _PartyBuy.Contract.Contributions(&_PartyBuy.CallOpts, arg0, arg1)
}

// ExpiresAt is a free data retrieval call binding the contract method 0x8622a689.
//
// Solidity: function expiresAt() view returns(uint256)
func (_PartyBuy *PartyBuyCaller) ExpiresAt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PartyBuy.contract.Call(opts, &out, "expiresAt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExpiresAt is a free data retrieval call binding the contract method 0x8622a689.
//
// Solidity: function expiresAt() view returns(uint256)
func (_PartyBuy *PartyBuySession) ExpiresAt() (*big.Int, error) {
	return _PartyBuy.Contract.ExpiresAt(&_PartyBuy.CallOpts)
}

// ExpiresAt is a free data retrieval call binding the contract method 0x8622a689.
//
// Solidity: function expiresAt() view returns(uint256)
func (_PartyBuy *PartyBuyCallerSession) ExpiresAt() (*big.Int, error) {
	return _PartyBuy.Contract.ExpiresAt(&_PartyBuy.CallOpts)
}

// GatedToken is a free data retrieval call binding the contract method 0x7a2ba9c4.
//
// Solidity: function gatedToken() view returns(address)
func (_PartyBuy *PartyBuyCaller) GatedToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PartyBuy.contract.Call(opts, &out, "gatedToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GatedToken is a free data retrieval call binding the contract method 0x7a2ba9c4.
//
// Solidity: function gatedToken() view returns(address)
func (_PartyBuy *PartyBuySession) GatedToken() (common.Address, error) {
	return _PartyBuy.Contract.GatedToken(&_PartyBuy.CallOpts)
}

// GatedToken is a free data retrieval call binding the contract method 0x7a2ba9c4.
//
// Solidity: function gatedToken() view returns(address)
func (_PartyBuy *PartyBuyCallerSession) GatedToken() (common.Address, error) {
	return _PartyBuy.Contract.GatedToken(&_PartyBuy.CallOpts)
}

// GatedTokenAmount is a free data retrieval call binding the contract method 0x8d42ecd6.
//
// Solidity: function gatedTokenAmount() view returns(uint256)
func (_PartyBuy *PartyBuyCaller) GatedTokenAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PartyBuy.contract.Call(opts, &out, "gatedTokenAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GatedTokenAmount is a free data retrieval call binding the contract method 0x8d42ecd6.
//
// Solidity: function gatedTokenAmount() view returns(uint256)
func (_PartyBuy *PartyBuySession) GatedTokenAmount() (*big.Int, error) {
	return _PartyBuy.Contract.GatedTokenAmount(&_PartyBuy.CallOpts)
}

// GatedTokenAmount is a free data retrieval call binding the contract method 0x8d42ecd6.
//
// Solidity: function gatedTokenAmount() view returns(uint256)
func (_PartyBuy *PartyBuyCallerSession) GatedTokenAmount() (*big.Int, error) {
	return _PartyBuy.Contract.GatedTokenAmount(&_PartyBuy.CallOpts)
}

// GetClaimAmounts is a free data retrieval call binding the contract method 0x6971524f.
//
// Solidity: function getClaimAmounts(address _contributor) view returns(uint256 _tokenAmount, uint256 _ethAmount)
func (_PartyBuy *PartyBuyCaller) GetClaimAmounts(opts *bind.CallOpts, _contributor common.Address) (struct {
	TokenAmount *big.Int
	EthAmount   *big.Int
}, error) {
	var out []interface{}
	err := _PartyBuy.contract.Call(opts, &out, "getClaimAmounts", _contributor)

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
func (_PartyBuy *PartyBuySession) GetClaimAmounts(_contributor common.Address) (struct {
	TokenAmount *big.Int
	EthAmount   *big.Int
}, error) {
	return _PartyBuy.Contract.GetClaimAmounts(&_PartyBuy.CallOpts, _contributor)
}

// GetClaimAmounts is a free data retrieval call binding the contract method 0x6971524f.
//
// Solidity: function getClaimAmounts(address _contributor) view returns(uint256 _tokenAmount, uint256 _ethAmount)
func (_PartyBuy *PartyBuyCallerSession) GetClaimAmounts(_contributor common.Address) (struct {
	TokenAmount *big.Int
	EthAmount   *big.Int
}, error) {
	return _PartyBuy.Contract.GetClaimAmounts(&_PartyBuy.CallOpts, _contributor)
}

// GetMaximumContributions is a free data retrieval call binding the contract method 0x9559da3b.
//
// Solidity: function getMaximumContributions() view returns(uint256 _maxContributions)
func (_PartyBuy *PartyBuyCaller) GetMaximumContributions(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PartyBuy.contract.Call(opts, &out, "getMaximumContributions")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaximumContributions is a free data retrieval call binding the contract method 0x9559da3b.
//
// Solidity: function getMaximumContributions() view returns(uint256 _maxContributions)
func (_PartyBuy *PartyBuySession) GetMaximumContributions() (*big.Int, error) {
	return _PartyBuy.Contract.GetMaximumContributions(&_PartyBuy.CallOpts)
}

// GetMaximumContributions is a free data retrieval call binding the contract method 0x9559da3b.
//
// Solidity: function getMaximumContributions() view returns(uint256 _maxContributions)
func (_PartyBuy *PartyBuyCallerSession) GetMaximumContributions() (*big.Int, error) {
	return _PartyBuy.Contract.GetMaximumContributions(&_PartyBuy.CallOpts)
}

// GetMaximumSpend is a free data retrieval call binding the contract method 0xacd13c59.
//
// Solidity: function getMaximumSpend() view returns(uint256 _maxSpend)
func (_PartyBuy *PartyBuyCaller) GetMaximumSpend(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PartyBuy.contract.Call(opts, &out, "getMaximumSpend")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaximumSpend is a free data retrieval call binding the contract method 0xacd13c59.
//
// Solidity: function getMaximumSpend() view returns(uint256 _maxSpend)
func (_PartyBuy *PartyBuySession) GetMaximumSpend() (*big.Int, error) {
	return _PartyBuy.Contract.GetMaximumSpend(&_PartyBuy.CallOpts)
}

// GetMaximumSpend is a free data retrieval call binding the contract method 0xacd13c59.
//
// Solidity: function getMaximumSpend() view returns(uint256 _maxSpend)
func (_PartyBuy *PartyBuyCallerSession) GetMaximumSpend() (*big.Int, error) {
	return _PartyBuy.Contract.GetMaximumSpend(&_PartyBuy.CallOpts)
}

// MaxPrice is a free data retrieval call binding the contract method 0xe38d6b5c.
//
// Solidity: function maxPrice() view returns(uint256)
func (_PartyBuy *PartyBuyCaller) MaxPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PartyBuy.contract.Call(opts, &out, "maxPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxPrice is a free data retrieval call binding the contract method 0xe38d6b5c.
//
// Solidity: function maxPrice() view returns(uint256)
func (_PartyBuy *PartyBuySession) MaxPrice() (*big.Int, error) {
	return _PartyBuy.Contract.MaxPrice(&_PartyBuy.CallOpts)
}

// MaxPrice is a free data retrieval call binding the contract method 0xe38d6b5c.
//
// Solidity: function maxPrice() view returns(uint256)
func (_PartyBuy *PartyBuyCallerSession) MaxPrice() (*big.Int, error) {
	return _PartyBuy.Contract.MaxPrice(&_PartyBuy.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PartyBuy *PartyBuyCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PartyBuy.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PartyBuy *PartyBuySession) Name() (string, error) {
	return _PartyBuy.Contract.Name(&_PartyBuy.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PartyBuy *PartyBuyCallerSession) Name() (string, error) {
	return _PartyBuy.Contract.Name(&_PartyBuy.CallOpts)
}

// NftContract is a free data retrieval call binding the contract method 0xd56d229d.
//
// Solidity: function nftContract() view returns(address)
func (_PartyBuy *PartyBuyCaller) NftContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PartyBuy.contract.Call(opts, &out, "nftContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NftContract is a free data retrieval call binding the contract method 0xd56d229d.
//
// Solidity: function nftContract() view returns(address)
func (_PartyBuy *PartyBuySession) NftContract() (common.Address, error) {
	return _PartyBuy.Contract.NftContract(&_PartyBuy.CallOpts)
}

// NftContract is a free data retrieval call binding the contract method 0xd56d229d.
//
// Solidity: function nftContract() view returns(address)
func (_PartyBuy *PartyBuyCallerSession) NftContract() (common.Address, error) {
	return _PartyBuy.Contract.NftContract(&_PartyBuy.CallOpts)
}

// PartyDAOMultisig is a free data retrieval call binding the contract method 0x3c4d12d9.
//
// Solidity: function partyDAOMultisig() view returns(address)
func (_PartyBuy *PartyBuyCaller) PartyDAOMultisig(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PartyBuy.contract.Call(opts, &out, "partyDAOMultisig")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PartyDAOMultisig is a free data retrieval call binding the contract method 0x3c4d12d9.
//
// Solidity: function partyDAOMultisig() view returns(address)
func (_PartyBuy *PartyBuySession) PartyDAOMultisig() (common.Address, error) {
	return _PartyBuy.Contract.PartyDAOMultisig(&_PartyBuy.CallOpts)
}

// PartyDAOMultisig is a free data retrieval call binding the contract method 0x3c4d12d9.
//
// Solidity: function partyDAOMultisig() view returns(address)
func (_PartyBuy *PartyBuyCallerSession) PartyDAOMultisig() (common.Address, error) {
	return _PartyBuy.Contract.PartyDAOMultisig(&_PartyBuy.CallOpts)
}

// PartyStatus is a free data retrieval call binding the contract method 0xdf51c07f.
//
// Solidity: function partyStatus() view returns(uint8)
func (_PartyBuy *PartyBuyCaller) PartyStatus(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _PartyBuy.contract.Call(opts, &out, "partyStatus")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// PartyStatus is a free data retrieval call binding the contract method 0xdf51c07f.
//
// Solidity: function partyStatus() view returns(uint8)
func (_PartyBuy *PartyBuySession) PartyStatus() (uint8, error) {
	return _PartyBuy.Contract.PartyStatus(&_PartyBuy.CallOpts)
}

// PartyStatus is a free data retrieval call binding the contract method 0xdf51c07f.
//
// Solidity: function partyStatus() view returns(uint8)
func (_PartyBuy *PartyBuyCallerSession) PartyStatus() (uint8, error) {
	return _PartyBuy.Contract.PartyStatus(&_PartyBuy.CallOpts)
}

// SplitBasisPoints is a free data retrieval call binding the contract method 0x2bbce5e6.
//
// Solidity: function splitBasisPoints() view returns(uint256)
func (_PartyBuy *PartyBuyCaller) SplitBasisPoints(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PartyBuy.contract.Call(opts, &out, "splitBasisPoints")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SplitBasisPoints is a free data retrieval call binding the contract method 0x2bbce5e6.
//
// Solidity: function splitBasisPoints() view returns(uint256)
func (_PartyBuy *PartyBuySession) SplitBasisPoints() (*big.Int, error) {
	return _PartyBuy.Contract.SplitBasisPoints(&_PartyBuy.CallOpts)
}

// SplitBasisPoints is a free data retrieval call binding the contract method 0x2bbce5e6.
//
// Solidity: function splitBasisPoints() view returns(uint256)
func (_PartyBuy *PartyBuyCallerSession) SplitBasisPoints() (*big.Int, error) {
	return _PartyBuy.Contract.SplitBasisPoints(&_PartyBuy.CallOpts)
}

// SplitRecipient is a free data retrieval call binding the contract method 0x4367a029.
//
// Solidity: function splitRecipient() view returns(address)
func (_PartyBuy *PartyBuyCaller) SplitRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PartyBuy.contract.Call(opts, &out, "splitRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SplitRecipient is a free data retrieval call binding the contract method 0x4367a029.
//
// Solidity: function splitRecipient() view returns(address)
func (_PartyBuy *PartyBuySession) SplitRecipient() (common.Address, error) {
	return _PartyBuy.Contract.SplitRecipient(&_PartyBuy.CallOpts)
}

// SplitRecipient is a free data retrieval call binding the contract method 0x4367a029.
//
// Solidity: function splitRecipient() view returns(address)
func (_PartyBuy *PartyBuyCallerSession) SplitRecipient() (common.Address, error) {
	return _PartyBuy.Contract.SplitRecipient(&_PartyBuy.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PartyBuy *PartyBuyCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PartyBuy.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PartyBuy *PartyBuySession) Symbol() (string, error) {
	return _PartyBuy.Contract.Symbol(&_PartyBuy.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PartyBuy *PartyBuyCallerSession) Symbol() (string, error) {
	return _PartyBuy.Contract.Symbol(&_PartyBuy.CallOpts)
}

// TokenId is a free data retrieval call binding the contract method 0x17d70f7c.
//
// Solidity: function tokenId() view returns(uint256)
func (_PartyBuy *PartyBuyCaller) TokenId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PartyBuy.contract.Call(opts, &out, "tokenId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenId is a free data retrieval call binding the contract method 0x17d70f7c.
//
// Solidity: function tokenId() view returns(uint256)
func (_PartyBuy *PartyBuySession) TokenId() (*big.Int, error) {
	return _PartyBuy.Contract.TokenId(&_PartyBuy.CallOpts)
}

// TokenId is a free data retrieval call binding the contract method 0x17d70f7c.
//
// Solidity: function tokenId() view returns(uint256)
func (_PartyBuy *PartyBuyCallerSession) TokenId() (*big.Int, error) {
	return _PartyBuy.Contract.TokenId(&_PartyBuy.CallOpts)
}

// TokenVault is a free data retrieval call binding the contract method 0x5bc789d9.
//
// Solidity: function tokenVault() view returns(address)
func (_PartyBuy *PartyBuyCaller) TokenVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PartyBuy.contract.Call(opts, &out, "tokenVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenVault is a free data retrieval call binding the contract method 0x5bc789d9.
//
// Solidity: function tokenVault() view returns(address)
func (_PartyBuy *PartyBuySession) TokenVault() (common.Address, error) {
	return _PartyBuy.Contract.TokenVault(&_PartyBuy.CallOpts)
}

// TokenVault is a free data retrieval call binding the contract method 0x5bc789d9.
//
// Solidity: function tokenVault() view returns(address)
func (_PartyBuy *PartyBuyCallerSession) TokenVault() (common.Address, error) {
	return _PartyBuy.Contract.TokenVault(&_PartyBuy.CallOpts)
}

// TokenVaultFactory is a free data retrieval call binding the contract method 0x0b203023.
//
// Solidity: function tokenVaultFactory() view returns(address)
func (_PartyBuy *PartyBuyCaller) TokenVaultFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PartyBuy.contract.Call(opts, &out, "tokenVaultFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenVaultFactory is a free data retrieval call binding the contract method 0x0b203023.
//
// Solidity: function tokenVaultFactory() view returns(address)
func (_PartyBuy *PartyBuySession) TokenVaultFactory() (common.Address, error) {
	return _PartyBuy.Contract.TokenVaultFactory(&_PartyBuy.CallOpts)
}

// TokenVaultFactory is a free data retrieval call binding the contract method 0x0b203023.
//
// Solidity: function tokenVaultFactory() view returns(address)
func (_PartyBuy *PartyBuyCallerSession) TokenVaultFactory() (common.Address, error) {
	return _PartyBuy.Contract.TokenVaultFactory(&_PartyBuy.CallOpts)
}

// TotalContributed is a free data retrieval call binding the contract method 0xa0f243b8.
//
// Solidity: function totalContributed(address ) view returns(uint256)
func (_PartyBuy *PartyBuyCaller) TotalContributed(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PartyBuy.contract.Call(opts, &out, "totalContributed", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalContributed is a free data retrieval call binding the contract method 0xa0f243b8.
//
// Solidity: function totalContributed(address ) view returns(uint256)
func (_PartyBuy *PartyBuySession) TotalContributed(arg0 common.Address) (*big.Int, error) {
	return _PartyBuy.Contract.TotalContributed(&_PartyBuy.CallOpts, arg0)
}

// TotalContributed is a free data retrieval call binding the contract method 0xa0f243b8.
//
// Solidity: function totalContributed(address ) view returns(uint256)
func (_PartyBuy *PartyBuyCallerSession) TotalContributed(arg0 common.Address) (*big.Int, error) {
	return _PartyBuy.Contract.TotalContributed(&_PartyBuy.CallOpts, arg0)
}

// TotalContributedToParty is a free data retrieval call binding the contract method 0x82a5c69a.
//
// Solidity: function totalContributedToParty() view returns(uint256)
func (_PartyBuy *PartyBuyCaller) TotalContributedToParty(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PartyBuy.contract.Call(opts, &out, "totalContributedToParty")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalContributedToParty is a free data retrieval call binding the contract method 0x82a5c69a.
//
// Solidity: function totalContributedToParty() view returns(uint256)
func (_PartyBuy *PartyBuySession) TotalContributedToParty() (*big.Int, error) {
	return _PartyBuy.Contract.TotalContributedToParty(&_PartyBuy.CallOpts)
}

// TotalContributedToParty is a free data retrieval call binding the contract method 0x82a5c69a.
//
// Solidity: function totalContributedToParty() view returns(uint256)
func (_PartyBuy *PartyBuyCallerSession) TotalContributedToParty() (*big.Int, error) {
	return _PartyBuy.Contract.TotalContributedToParty(&_PartyBuy.CallOpts)
}

// TotalEthUsed is a free data retrieval call binding the contract method 0x25b42a26.
//
// Solidity: function totalEthUsed(address _contributor) view returns(uint256 _total)
func (_PartyBuy *PartyBuyCaller) TotalEthUsed(opts *bind.CallOpts, _contributor common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PartyBuy.contract.Call(opts, &out, "totalEthUsed", _contributor)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalEthUsed is a free data retrieval call binding the contract method 0x25b42a26.
//
// Solidity: function totalEthUsed(address _contributor) view returns(uint256 _total)
func (_PartyBuy *PartyBuySession) TotalEthUsed(_contributor common.Address) (*big.Int, error) {
	return _PartyBuy.Contract.TotalEthUsed(&_PartyBuy.CallOpts, _contributor)
}

// TotalEthUsed is a free data retrieval call binding the contract method 0x25b42a26.
//
// Solidity: function totalEthUsed(address _contributor) view returns(uint256 _total)
func (_PartyBuy *PartyBuyCallerSession) TotalEthUsed(_contributor common.Address) (*big.Int, error) {
	return _PartyBuy.Contract.TotalEthUsed(&_PartyBuy.CallOpts, _contributor)
}

// TotalSpent is a free data retrieval call binding the contract method 0xfb346eab.
//
// Solidity: function totalSpent() view returns(uint256)
func (_PartyBuy *PartyBuyCaller) TotalSpent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PartyBuy.contract.Call(opts, &out, "totalSpent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSpent is a free data retrieval call binding the contract method 0xfb346eab.
//
// Solidity: function totalSpent() view returns(uint256)
func (_PartyBuy *PartyBuySession) TotalSpent() (*big.Int, error) {
	return _PartyBuy.Contract.TotalSpent(&_PartyBuy.CallOpts)
}

// TotalSpent is a free data retrieval call binding the contract method 0xfb346eab.
//
// Solidity: function totalSpent() view returns(uint256)
func (_PartyBuy *PartyBuyCallerSession) TotalSpent() (*big.Int, error) {
	return _PartyBuy.Contract.TotalSpent(&_PartyBuy.CallOpts)
}

// ValueToTokens is a free data retrieval call binding the contract method 0x9744b8dc.
//
// Solidity: function valueToTokens(uint256 _value) pure returns(uint256 _tokens)
func (_PartyBuy *PartyBuyCaller) ValueToTokens(opts *bind.CallOpts, _value *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PartyBuy.contract.Call(opts, &out, "valueToTokens", _value)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValueToTokens is a free data retrieval call binding the contract method 0x9744b8dc.
//
// Solidity: function valueToTokens(uint256 _value) pure returns(uint256 _tokens)
func (_PartyBuy *PartyBuySession) ValueToTokens(_value *big.Int) (*big.Int, error) {
	return _PartyBuy.Contract.ValueToTokens(&_PartyBuy.CallOpts, _value)
}

// ValueToTokens is a free data retrieval call binding the contract method 0x9744b8dc.
//
// Solidity: function valueToTokens(uint256 _value) pure returns(uint256 _tokens)
func (_PartyBuy *PartyBuyCallerSession) ValueToTokens(_value *big.Int) (*big.Int, error) {
	return _PartyBuy.Contract.ValueToTokens(&_PartyBuy.CallOpts, _value)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_PartyBuy *PartyBuyCaller) Weth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PartyBuy.contract.Call(opts, &out, "weth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_PartyBuy *PartyBuySession) Weth() (common.Address, error) {
	return _PartyBuy.Contract.Weth(&_PartyBuy.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_PartyBuy *PartyBuyCallerSession) Weth() (common.Address, error) {
	return _PartyBuy.Contract.Weth(&_PartyBuy.CallOpts)
}

// Buy is a paid mutator transaction binding the contract method 0xcb36ea38.
//
// Solidity: function buy(uint256 _value, address _targetContract, bytes _calldata) returns()
func (_PartyBuy *PartyBuyTransactor) Buy(opts *bind.TransactOpts, _value *big.Int, _targetContract common.Address, _calldata []byte) (*types.Transaction, error) {
	return _PartyBuy.contract.Transact(opts, "buy", _value, _targetContract, _calldata)
}

// Buy is a paid mutator transaction binding the contract method 0xcb36ea38.
//
// Solidity: function buy(uint256 _value, address _targetContract, bytes _calldata) returns()
func (_PartyBuy *PartyBuySession) Buy(_value *big.Int, _targetContract common.Address, _calldata []byte) (*types.Transaction, error) {
	return _PartyBuy.Contract.Buy(&_PartyBuy.TransactOpts, _value, _targetContract, _calldata)
}

// Buy is a paid mutator transaction binding the contract method 0xcb36ea38.
//
// Solidity: function buy(uint256 _value, address _targetContract, bytes _calldata) returns()
func (_PartyBuy *PartyBuyTransactorSession) Buy(_value *big.Int, _targetContract common.Address, _calldata []byte) (*types.Transaction, error) {
	return _PartyBuy.Contract.Buy(&_PartyBuy.TransactOpts, _value, _targetContract, _calldata)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address _contributor) returns()
func (_PartyBuy *PartyBuyTransactor) Claim(opts *bind.TransactOpts, _contributor common.Address) (*types.Transaction, error) {
	return _PartyBuy.contract.Transact(opts, "claim", _contributor)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address _contributor) returns()
func (_PartyBuy *PartyBuySession) Claim(_contributor common.Address) (*types.Transaction, error) {
	return _PartyBuy.Contract.Claim(&_PartyBuy.TransactOpts, _contributor)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address _contributor) returns()
func (_PartyBuy *PartyBuyTransactorSession) Claim(_contributor common.Address) (*types.Transaction, error) {
	return _PartyBuy.Contract.Claim(&_PartyBuy.TransactOpts, _contributor)
}

// Contribute is a paid mutator transaction binding the contract method 0xd7bb99ba.
//
// Solidity: function contribute() payable returns()
func (_PartyBuy *PartyBuyTransactor) Contribute(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PartyBuy.contract.Transact(opts, "contribute")
}

// Contribute is a paid mutator transaction binding the contract method 0xd7bb99ba.
//
// Solidity: function contribute() payable returns()
func (_PartyBuy *PartyBuySession) Contribute() (*types.Transaction, error) {
	return _PartyBuy.Contract.Contribute(&_PartyBuy.TransactOpts)
}

// Contribute is a paid mutator transaction binding the contract method 0xd7bb99ba.
//
// Solidity: function contribute() payable returns()
func (_PartyBuy *PartyBuyTransactorSession) Contribute() (*types.Transaction, error) {
	return _PartyBuy.Contract.Contribute(&_PartyBuy.TransactOpts)
}

// EmergencyCall is a paid mutator transaction binding the contract method 0xc4bf0220.
//
// Solidity: function emergencyCall(address _contract, bytes _calldata) returns(bool _success, bytes _returnData)
func (_PartyBuy *PartyBuyTransactor) EmergencyCall(opts *bind.TransactOpts, _contract common.Address, _calldata []byte) (*types.Transaction, error) {
	return _PartyBuy.contract.Transact(opts, "emergencyCall", _contract, _calldata)
}

// EmergencyCall is a paid mutator transaction binding the contract method 0xc4bf0220.
//
// Solidity: function emergencyCall(address _contract, bytes _calldata) returns(bool _success, bytes _returnData)
func (_PartyBuy *PartyBuySession) EmergencyCall(_contract common.Address, _calldata []byte) (*types.Transaction, error) {
	return _PartyBuy.Contract.EmergencyCall(&_PartyBuy.TransactOpts, _contract, _calldata)
}

// EmergencyCall is a paid mutator transaction binding the contract method 0xc4bf0220.
//
// Solidity: function emergencyCall(address _contract, bytes _calldata) returns(bool _success, bytes _returnData)
func (_PartyBuy *PartyBuyTransactorSession) EmergencyCall(_contract common.Address, _calldata []byte) (*types.Transaction, error) {
	return _PartyBuy.Contract.EmergencyCall(&_PartyBuy.TransactOpts, _contract, _calldata)
}

// EmergencyForceLost is a paid mutator transaction binding the contract method 0x17821fdc.
//
// Solidity: function emergencyForceLost() returns()
func (_PartyBuy *PartyBuyTransactor) EmergencyForceLost(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PartyBuy.contract.Transact(opts, "emergencyForceLost")
}

// EmergencyForceLost is a paid mutator transaction binding the contract method 0x17821fdc.
//
// Solidity: function emergencyForceLost() returns()
func (_PartyBuy *PartyBuySession) EmergencyForceLost() (*types.Transaction, error) {
	return _PartyBuy.Contract.EmergencyForceLost(&_PartyBuy.TransactOpts)
}

// EmergencyForceLost is a paid mutator transaction binding the contract method 0x17821fdc.
//
// Solidity: function emergencyForceLost() returns()
func (_PartyBuy *PartyBuyTransactorSession) EmergencyForceLost() (*types.Transaction, error) {
	return _PartyBuy.Contract.EmergencyForceLost(&_PartyBuy.TransactOpts)
}

// EmergencyWithdrawEth is a paid mutator transaction binding the contract method 0x429093cc.
//
// Solidity: function emergencyWithdrawEth(uint256 _value) returns()
func (_PartyBuy *PartyBuyTransactor) EmergencyWithdrawEth(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _PartyBuy.contract.Transact(opts, "emergencyWithdrawEth", _value)
}

// EmergencyWithdrawEth is a paid mutator transaction binding the contract method 0x429093cc.
//
// Solidity: function emergencyWithdrawEth(uint256 _value) returns()
func (_PartyBuy *PartyBuySession) EmergencyWithdrawEth(_value *big.Int) (*types.Transaction, error) {
	return _PartyBuy.Contract.EmergencyWithdrawEth(&_PartyBuy.TransactOpts, _value)
}

// EmergencyWithdrawEth is a paid mutator transaction binding the contract method 0x429093cc.
//
// Solidity: function emergencyWithdrawEth(uint256 _value) returns()
func (_PartyBuy *PartyBuyTransactorSession) EmergencyWithdrawEth(_value *big.Int) (*types.Transaction, error) {
	return _PartyBuy.Contract.EmergencyWithdrawEth(&_PartyBuy.TransactOpts, _value)
}

// Expire is a paid mutator transaction binding the contract method 0x79599f96.
//
// Solidity: function expire() returns()
func (_PartyBuy *PartyBuyTransactor) Expire(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PartyBuy.contract.Transact(opts, "expire")
}

// Expire is a paid mutator transaction binding the contract method 0x79599f96.
//
// Solidity: function expire() returns()
func (_PartyBuy *PartyBuySession) Expire() (*types.Transaction, error) {
	return _PartyBuy.Contract.Expire(&_PartyBuy.TransactOpts)
}

// Expire is a paid mutator transaction binding the contract method 0x79599f96.
//
// Solidity: function expire() returns()
func (_PartyBuy *PartyBuyTransactorSession) Expire() (*types.Transaction, error) {
	return _PartyBuy.Contract.Expire(&_PartyBuy.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x9a3596a3.
//
// Solidity: function initialize(address _nftContract, uint256 _tokenId, uint256 _maxPrice, uint256 _secondsToTimeout, (address,uint256) _split, (address,uint256) _tokenGate, string _name, string _symbol) returns()
func (_PartyBuy *PartyBuyTransactor) Initialize(opts *bind.TransactOpts, _nftContract common.Address, _tokenId *big.Int, _maxPrice *big.Int, _secondsToTimeout *big.Int, _split StructsAddressAndAmount, _tokenGate StructsAddressAndAmount, _name string, _symbol string) (*types.Transaction, error) {
	return _PartyBuy.contract.Transact(opts, "initialize", _nftContract, _tokenId, _maxPrice, _secondsToTimeout, _split, _tokenGate, _name, _symbol)
}

// Initialize is a paid mutator transaction binding the contract method 0x9a3596a3.
//
// Solidity: function initialize(address _nftContract, uint256 _tokenId, uint256 _maxPrice, uint256 _secondsToTimeout, (address,uint256) _split, (address,uint256) _tokenGate, string _name, string _symbol) returns()
func (_PartyBuy *PartyBuySession) Initialize(_nftContract common.Address, _tokenId *big.Int, _maxPrice *big.Int, _secondsToTimeout *big.Int, _split StructsAddressAndAmount, _tokenGate StructsAddressAndAmount, _name string, _symbol string) (*types.Transaction, error) {
	return _PartyBuy.Contract.Initialize(&_PartyBuy.TransactOpts, _nftContract, _tokenId, _maxPrice, _secondsToTimeout, _split, _tokenGate, _name, _symbol)
}

// Initialize is a paid mutator transaction binding the contract method 0x9a3596a3.
//
// Solidity: function initialize(address _nftContract, uint256 _tokenId, uint256 _maxPrice, uint256 _secondsToTimeout, (address,uint256) _split, (address,uint256) _tokenGate, string _name, string _symbol) returns()
func (_PartyBuy *PartyBuyTransactorSession) Initialize(_nftContract common.Address, _tokenId *big.Int, _maxPrice *big.Int, _secondsToTimeout *big.Int, _split StructsAddressAndAmount, _tokenGate StructsAddressAndAmount, _name string, _symbol string) (*types.Transaction, error) {
	return _PartyBuy.Contract.Initialize(&_PartyBuy.TransactOpts, _nftContract, _tokenId, _maxPrice, _secondsToTimeout, _split, _tokenGate, _name, _symbol)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_PartyBuy *PartyBuyTransactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _PartyBuy.contract.Transact(opts, "onERC721Received", arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_PartyBuy *PartyBuySession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _PartyBuy.Contract.OnERC721Received(&_PartyBuy.TransactOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_PartyBuy *PartyBuyTransactorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _PartyBuy.Contract.OnERC721Received(&_PartyBuy.TransactOpts, arg0, arg1, arg2, arg3)
}

// PartyBuyBoughtIterator is returned from FilterBought and is used to iterate over the raw logs and unpacked data for Bought events raised by the PartyBuy contract.
type PartyBuyBoughtIterator struct {
	Event *PartyBuyBought // Event containing the contract specifics and raw log

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
func (it *PartyBuyBoughtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PartyBuyBought)
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
		it.Event = new(PartyBuyBought)
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
func (it *PartyBuyBoughtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PartyBuyBoughtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PartyBuyBought represents a Bought event raised by the PartyBuy contract.
type PartyBuyBought struct {
	TriggeredBy      common.Address
	TargetAddress    common.Address
	EthSpent         *big.Int
	EthFeePaid       *big.Int
	TotalContributed *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterBought is a free log retrieval operation binding the contract event 0x7ce543d1780f3bdc3dac42da06c95da802653cd1b212b8d74ec3e3c33ad7095c.
//
// Solidity: event Bought(address triggeredBy, address targetAddress, uint256 ethSpent, uint256 ethFeePaid, uint256 totalContributed)
func (_PartyBuy *PartyBuyFilterer) FilterBought(opts *bind.FilterOpts) (*PartyBuyBoughtIterator, error) {

	logs, sub, err := _PartyBuy.contract.FilterLogs(opts, "Bought")
	if err != nil {
		return nil, err
	}
	return &PartyBuyBoughtIterator{contract: _PartyBuy.contract, event: "Bought", logs: logs, sub: sub}, nil
}

// WatchBought is a free log subscription operation binding the contract event 0x7ce543d1780f3bdc3dac42da06c95da802653cd1b212b8d74ec3e3c33ad7095c.
//
// Solidity: event Bought(address triggeredBy, address targetAddress, uint256 ethSpent, uint256 ethFeePaid, uint256 totalContributed)
func (_PartyBuy *PartyBuyFilterer) WatchBought(opts *bind.WatchOpts, sink chan<- *PartyBuyBought) (event.Subscription, error) {

	logs, sub, err := _PartyBuy.contract.WatchLogs(opts, "Bought")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PartyBuyBought)
				if err := _PartyBuy.contract.UnpackLog(event, "Bought", log); err != nil {
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

// ParseBought is a log parse operation binding the contract event 0x7ce543d1780f3bdc3dac42da06c95da802653cd1b212b8d74ec3e3c33ad7095c.
//
// Solidity: event Bought(address triggeredBy, address targetAddress, uint256 ethSpent, uint256 ethFeePaid, uint256 totalContributed)
func (_PartyBuy *PartyBuyFilterer) ParseBought(log types.Log) (*PartyBuyBought, error) {
	event := new(PartyBuyBought)
	if err := _PartyBuy.contract.UnpackLog(event, "Bought", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PartyBuyClaimedIterator is returned from FilterClaimed and is used to iterate over the raw logs and unpacked data for Claimed events raised by the PartyBuy contract.
type PartyBuyClaimedIterator struct {
	Event *PartyBuyClaimed // Event containing the contract specifics and raw log

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
func (it *PartyBuyClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PartyBuyClaimed)
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
		it.Event = new(PartyBuyClaimed)
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
func (it *PartyBuyClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PartyBuyClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PartyBuyClaimed represents a Claimed event raised by the PartyBuy contract.
type PartyBuyClaimed struct {
	Contributor        common.Address
	TotalContributed   *big.Int
	ExcessContribution *big.Int
	TokenAmount        *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterClaimed is a free log retrieval operation binding the contract event 0x9cdcf2f7714cca3508c7f0110b04a90a80a3a8dd0e35de99689db74d28c5383e.
//
// Solidity: event Claimed(address indexed contributor, uint256 totalContributed, uint256 excessContribution, uint256 tokenAmount)
func (_PartyBuy *PartyBuyFilterer) FilterClaimed(opts *bind.FilterOpts, contributor []common.Address) (*PartyBuyClaimedIterator, error) {

	var contributorRule []interface{}
	for _, contributorItem := range contributor {
		contributorRule = append(contributorRule, contributorItem)
	}

	logs, sub, err := _PartyBuy.contract.FilterLogs(opts, "Claimed", contributorRule)
	if err != nil {
		return nil, err
	}
	return &PartyBuyClaimedIterator{contract: _PartyBuy.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0x9cdcf2f7714cca3508c7f0110b04a90a80a3a8dd0e35de99689db74d28c5383e.
//
// Solidity: event Claimed(address indexed contributor, uint256 totalContributed, uint256 excessContribution, uint256 tokenAmount)
func (_PartyBuy *PartyBuyFilterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *PartyBuyClaimed, contributor []common.Address) (event.Subscription, error) {

	var contributorRule []interface{}
	for _, contributorItem := range contributor {
		contributorRule = append(contributorRule, contributorItem)
	}

	logs, sub, err := _PartyBuy.contract.WatchLogs(opts, "Claimed", contributorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PartyBuyClaimed)
				if err := _PartyBuy.contract.UnpackLog(event, "Claimed", log); err != nil {
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
func (_PartyBuy *PartyBuyFilterer) ParseClaimed(log types.Log) (*PartyBuyClaimed, error) {
	event := new(PartyBuyClaimed)
	if err := _PartyBuy.contract.UnpackLog(event, "Claimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PartyBuyContributedIterator is returned from FilterContributed and is used to iterate over the raw logs and unpacked data for Contributed events raised by the PartyBuy contract.
type PartyBuyContributedIterator struct {
	Event *PartyBuyContributed // Event containing the contract specifics and raw log

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
func (it *PartyBuyContributedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PartyBuyContributed)
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
		it.Event = new(PartyBuyContributed)
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
func (it *PartyBuyContributedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PartyBuyContributedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PartyBuyContributed represents a Contributed event raised by the PartyBuy contract.
type PartyBuyContributed struct {
	Contributor                     common.Address
	Amount                          *big.Int
	PreviousTotalContributedToParty *big.Int
	TotalFromContributor            *big.Int
	Raw                             types.Log // Blockchain specific contextual infos
}

// FilterContributed is a free log retrieval operation binding the contract event 0xb2623081601722547aae8781994e01a1974d95b0ad9ce6a0cfbe17487556257f.
//
// Solidity: event Contributed(address indexed contributor, uint256 amount, uint256 previousTotalContributedToParty, uint256 totalFromContributor)
func (_PartyBuy *PartyBuyFilterer) FilterContributed(opts *bind.FilterOpts, contributor []common.Address) (*PartyBuyContributedIterator, error) {

	var contributorRule []interface{}
	for _, contributorItem := range contributor {
		contributorRule = append(contributorRule, contributorItem)
	}

	logs, sub, err := _PartyBuy.contract.FilterLogs(opts, "Contributed", contributorRule)
	if err != nil {
		return nil, err
	}
	return &PartyBuyContributedIterator{contract: _PartyBuy.contract, event: "Contributed", logs: logs, sub: sub}, nil
}

// WatchContributed is a free log subscription operation binding the contract event 0xb2623081601722547aae8781994e01a1974d95b0ad9ce6a0cfbe17487556257f.
//
// Solidity: event Contributed(address indexed contributor, uint256 amount, uint256 previousTotalContributedToParty, uint256 totalFromContributor)
func (_PartyBuy *PartyBuyFilterer) WatchContributed(opts *bind.WatchOpts, sink chan<- *PartyBuyContributed, contributor []common.Address) (event.Subscription, error) {

	var contributorRule []interface{}
	for _, contributorItem := range contributor {
		contributorRule = append(contributorRule, contributorItem)
	}

	logs, sub, err := _PartyBuy.contract.WatchLogs(opts, "Contributed", contributorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PartyBuyContributed)
				if err := _PartyBuy.contract.UnpackLog(event, "Contributed", log); err != nil {
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
func (_PartyBuy *PartyBuyFilterer) ParseContributed(log types.Log) (*PartyBuyContributed, error) {
	event := new(PartyBuyContributed)
	if err := _PartyBuy.contract.UnpackLog(event, "Contributed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PartyBuyExpiredIterator is returned from FilterExpired and is used to iterate over the raw logs and unpacked data for Expired events raised by the PartyBuy contract.
type PartyBuyExpiredIterator struct {
	Event *PartyBuyExpired // Event containing the contract specifics and raw log

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
func (it *PartyBuyExpiredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PartyBuyExpired)
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
		it.Event = new(PartyBuyExpired)
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
func (it *PartyBuyExpiredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PartyBuyExpiredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PartyBuyExpired represents a Expired event raised by the PartyBuy contract.
type PartyBuyExpired struct {
	TriggeredBy common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterExpired is a free log retrieval operation binding the contract event 0xd5669ebe8b90ed693033f1915dcea031e7a16a6f146a8326c02eec10eed77edd.
//
// Solidity: event Expired(address triggeredBy)
func (_PartyBuy *PartyBuyFilterer) FilterExpired(opts *bind.FilterOpts) (*PartyBuyExpiredIterator, error) {

	logs, sub, err := _PartyBuy.contract.FilterLogs(opts, "Expired")
	if err != nil {
		return nil, err
	}
	return &PartyBuyExpiredIterator{contract: _PartyBuy.contract, event: "Expired", logs: logs, sub: sub}, nil
}

// WatchExpired is a free log subscription operation binding the contract event 0xd5669ebe8b90ed693033f1915dcea031e7a16a6f146a8326c02eec10eed77edd.
//
// Solidity: event Expired(address triggeredBy)
func (_PartyBuy *PartyBuyFilterer) WatchExpired(opts *bind.WatchOpts, sink chan<- *PartyBuyExpired) (event.Subscription, error) {

	logs, sub, err := _PartyBuy.contract.WatchLogs(opts, "Expired")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PartyBuyExpired)
				if err := _PartyBuy.contract.UnpackLog(event, "Expired", log); err != nil {
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
func (_PartyBuy *PartyBuyFilterer) ParseExpired(log types.Log) (*PartyBuyExpired, error) {
	event := new(PartyBuyExpired)
	if err := _PartyBuy.contract.UnpackLog(event, "Expired", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
