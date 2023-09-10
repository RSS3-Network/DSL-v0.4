// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package friendtech

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

// FriendTechMetaData contains all meta data concerning the FriendTech contract.
var FriendTechMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"subject\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isBuy\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shareAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ethAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"protocolEthAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"subjectEthAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"supply\",\"type\":\"uint256\"}],\"name\":\"Trade\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sharesSubject\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"buyShares\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sharesSubject\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"getBuyPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sharesSubject\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"getBuyPriceAfterFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"supply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sharesSubject\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"getSellPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sharesSubject\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"getSellPriceAfterFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolFeeDestination\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolFeePercent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sharesSubject\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"sellShares\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeDestination\",\"type\":\"address\"}],\"name\":\"setFeeDestination\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_feePercent\",\"type\":\"uint256\"}],\"name\":\"setProtocolFeePercent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_feePercent\",\"type\":\"uint256\"}],\"name\":\"setSubjectFeePercent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"sharesBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"sharesSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"subjectFeePercent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// FriendTechABI is the input ABI used to generate the binding from.
// Deprecated: Use FriendTechMetaData.ABI instead.
var FriendTechABI = FriendTechMetaData.ABI

// FriendTech is an auto generated Go binding around an Ethereum contract.
type FriendTech struct {
	FriendTechCaller     // Read-only binding to the contract
	FriendTechTransactor // Write-only binding to the contract
	FriendTechFilterer   // Log filterer for contract events
}

// FriendTechCaller is an auto generated read-only Go binding around an Ethereum contract.
type FriendTechCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FriendTechTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FriendTechTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FriendTechFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FriendTechFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FriendTechSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FriendTechSession struct {
	Contract     *FriendTech       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FriendTechCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FriendTechCallerSession struct {
	Contract *FriendTechCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// FriendTechTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FriendTechTransactorSession struct {
	Contract     *FriendTechTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// FriendTechRaw is an auto generated low-level Go binding around an Ethereum contract.
type FriendTechRaw struct {
	Contract *FriendTech // Generic contract binding to access the raw methods on
}

// FriendTechCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FriendTechCallerRaw struct {
	Contract *FriendTechCaller // Generic read-only contract binding to access the raw methods on
}

// FriendTechTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FriendTechTransactorRaw struct {
	Contract *FriendTechTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFriendTech creates a new instance of FriendTech, bound to a specific deployed contract.
func NewFriendTech(address common.Address, backend bind.ContractBackend) (*FriendTech, error) {
	contract, err := bindFriendTech(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FriendTech{FriendTechCaller: FriendTechCaller{contract: contract}, FriendTechTransactor: FriendTechTransactor{contract: contract}, FriendTechFilterer: FriendTechFilterer{contract: contract}}, nil
}

// NewFriendTechCaller creates a new read-only instance of FriendTech, bound to a specific deployed contract.
func NewFriendTechCaller(address common.Address, caller bind.ContractCaller) (*FriendTechCaller, error) {
	contract, err := bindFriendTech(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FriendTechCaller{contract: contract}, nil
}

// NewFriendTechTransactor creates a new write-only instance of FriendTech, bound to a specific deployed contract.
func NewFriendTechTransactor(address common.Address, transactor bind.ContractTransactor) (*FriendTechTransactor, error) {
	contract, err := bindFriendTech(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FriendTechTransactor{contract: contract}, nil
}

// NewFriendTechFilterer creates a new log filterer instance of FriendTech, bound to a specific deployed contract.
func NewFriendTechFilterer(address common.Address, filterer bind.ContractFilterer) (*FriendTechFilterer, error) {
	contract, err := bindFriendTech(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FriendTechFilterer{contract: contract}, nil
}

// bindFriendTech binds a generic wrapper to an already deployed contract.
func bindFriendTech(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FriendTechMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FriendTech *FriendTechRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FriendTech.Contract.FriendTechCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FriendTech *FriendTechRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FriendTech.Contract.FriendTechTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FriendTech *FriendTechRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FriendTech.Contract.FriendTechTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FriendTech *FriendTechCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FriendTech.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FriendTech *FriendTechTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FriendTech.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FriendTech *FriendTechTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FriendTech.Contract.contract.Transact(opts, method, params...)
}

// GetBuyPrice is a free data retrieval call binding the contract method 0x4635256e.
//
// Solidity: function getBuyPrice(address sharesSubject, uint256 amount) view returns(uint256)
func (_FriendTech *FriendTechCaller) GetBuyPrice(opts *bind.CallOpts, sharesSubject common.Address, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FriendTech.contract.Call(opts, &out, "getBuyPrice", sharesSubject, amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBuyPrice is a free data retrieval call binding the contract method 0x4635256e.
//
// Solidity: function getBuyPrice(address sharesSubject, uint256 amount) view returns(uint256)
func (_FriendTech *FriendTechSession) GetBuyPrice(sharesSubject common.Address, amount *big.Int) (*big.Int, error) {
	return _FriendTech.Contract.GetBuyPrice(&_FriendTech.CallOpts, sharesSubject, amount)
}

// GetBuyPrice is a free data retrieval call binding the contract method 0x4635256e.
//
// Solidity: function getBuyPrice(address sharesSubject, uint256 amount) view returns(uint256)
func (_FriendTech *FriendTechCallerSession) GetBuyPrice(sharesSubject common.Address, amount *big.Int) (*big.Int, error) {
	return _FriendTech.Contract.GetBuyPrice(&_FriendTech.CallOpts, sharesSubject, amount)
}

// GetBuyPriceAfterFee is a free data retrieval call binding the contract method 0x0f026f6d.
//
// Solidity: function getBuyPriceAfterFee(address sharesSubject, uint256 amount) view returns(uint256)
func (_FriendTech *FriendTechCaller) GetBuyPriceAfterFee(opts *bind.CallOpts, sharesSubject common.Address, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FriendTech.contract.Call(opts, &out, "getBuyPriceAfterFee", sharesSubject, amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBuyPriceAfterFee is a free data retrieval call binding the contract method 0x0f026f6d.
//
// Solidity: function getBuyPriceAfterFee(address sharesSubject, uint256 amount) view returns(uint256)
func (_FriendTech *FriendTechSession) GetBuyPriceAfterFee(sharesSubject common.Address, amount *big.Int) (*big.Int, error) {
	return _FriendTech.Contract.GetBuyPriceAfterFee(&_FriendTech.CallOpts, sharesSubject, amount)
}

// GetBuyPriceAfterFee is a free data retrieval call binding the contract method 0x0f026f6d.
//
// Solidity: function getBuyPriceAfterFee(address sharesSubject, uint256 amount) view returns(uint256)
func (_FriendTech *FriendTechCallerSession) GetBuyPriceAfterFee(sharesSubject common.Address, amount *big.Int) (*big.Int, error) {
	return _FriendTech.Contract.GetBuyPriceAfterFee(&_FriendTech.CallOpts, sharesSubject, amount)
}

// GetPrice is a free data retrieval call binding the contract method 0x5cf4ee91.
//
// Solidity: function getPrice(uint256 supply, uint256 amount) pure returns(uint256)
func (_FriendTech *FriendTechCaller) GetPrice(opts *bind.CallOpts, supply *big.Int, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FriendTech.contract.Call(opts, &out, "getPrice", supply, amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrice is a free data retrieval call binding the contract method 0x5cf4ee91.
//
// Solidity: function getPrice(uint256 supply, uint256 amount) pure returns(uint256)
func (_FriendTech *FriendTechSession) GetPrice(supply *big.Int, amount *big.Int) (*big.Int, error) {
	return _FriendTech.Contract.GetPrice(&_FriendTech.CallOpts, supply, amount)
}

// GetPrice is a free data retrieval call binding the contract method 0x5cf4ee91.
//
// Solidity: function getPrice(uint256 supply, uint256 amount) pure returns(uint256)
func (_FriendTech *FriendTechCallerSession) GetPrice(supply *big.Int, amount *big.Int) (*big.Int, error) {
	return _FriendTech.Contract.GetPrice(&_FriendTech.CallOpts, supply, amount)
}

// GetSellPrice is a free data retrieval call binding the contract method 0x9ae71781.
//
// Solidity: function getSellPrice(address sharesSubject, uint256 amount) view returns(uint256)
func (_FriendTech *FriendTechCaller) GetSellPrice(opts *bind.CallOpts, sharesSubject common.Address, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FriendTech.contract.Call(opts, &out, "getSellPrice", sharesSubject, amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSellPrice is a free data retrieval call binding the contract method 0x9ae71781.
//
// Solidity: function getSellPrice(address sharesSubject, uint256 amount) view returns(uint256)
func (_FriendTech *FriendTechSession) GetSellPrice(sharesSubject common.Address, amount *big.Int) (*big.Int, error) {
	return _FriendTech.Contract.GetSellPrice(&_FriendTech.CallOpts, sharesSubject, amount)
}

// GetSellPrice is a free data retrieval call binding the contract method 0x9ae71781.
//
// Solidity: function getSellPrice(address sharesSubject, uint256 amount) view returns(uint256)
func (_FriendTech *FriendTechCallerSession) GetSellPrice(sharesSubject common.Address, amount *big.Int) (*big.Int, error) {
	return _FriendTech.Contract.GetSellPrice(&_FriendTech.CallOpts, sharesSubject, amount)
}

// GetSellPriceAfterFee is a free data retrieval call binding the contract method 0x2267a89c.
//
// Solidity: function getSellPriceAfterFee(address sharesSubject, uint256 amount) view returns(uint256)
func (_FriendTech *FriendTechCaller) GetSellPriceAfterFee(opts *bind.CallOpts, sharesSubject common.Address, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FriendTech.contract.Call(opts, &out, "getSellPriceAfterFee", sharesSubject, amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSellPriceAfterFee is a free data retrieval call binding the contract method 0x2267a89c.
//
// Solidity: function getSellPriceAfterFee(address sharesSubject, uint256 amount) view returns(uint256)
func (_FriendTech *FriendTechSession) GetSellPriceAfterFee(sharesSubject common.Address, amount *big.Int) (*big.Int, error) {
	return _FriendTech.Contract.GetSellPriceAfterFee(&_FriendTech.CallOpts, sharesSubject, amount)
}

// GetSellPriceAfterFee is a free data retrieval call binding the contract method 0x2267a89c.
//
// Solidity: function getSellPriceAfterFee(address sharesSubject, uint256 amount) view returns(uint256)
func (_FriendTech *FriendTechCallerSession) GetSellPriceAfterFee(sharesSubject common.Address, amount *big.Int) (*big.Int, error) {
	return _FriendTech.Contract.GetSellPriceAfterFee(&_FriendTech.CallOpts, sharesSubject, amount)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FriendTech *FriendTechCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FriendTech.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FriendTech *FriendTechSession) Owner() (common.Address, error) {
	return _FriendTech.Contract.Owner(&_FriendTech.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FriendTech *FriendTechCallerSession) Owner() (common.Address, error) {
	return _FriendTech.Contract.Owner(&_FriendTech.CallOpts)
}

// ProtocolFeeDestination is a free data retrieval call binding the contract method 0x4ce7957c.
//
// Solidity: function protocolFeeDestination() view returns(address)
func (_FriendTech *FriendTechCaller) ProtocolFeeDestination(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FriendTech.contract.Call(opts, &out, "protocolFeeDestination")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProtocolFeeDestination is a free data retrieval call binding the contract method 0x4ce7957c.
//
// Solidity: function protocolFeeDestination() view returns(address)
func (_FriendTech *FriendTechSession) ProtocolFeeDestination() (common.Address, error) {
	return _FriendTech.Contract.ProtocolFeeDestination(&_FriendTech.CallOpts)
}

// ProtocolFeeDestination is a free data retrieval call binding the contract method 0x4ce7957c.
//
// Solidity: function protocolFeeDestination() view returns(address)
func (_FriendTech *FriendTechCallerSession) ProtocolFeeDestination() (common.Address, error) {
	return _FriendTech.Contract.ProtocolFeeDestination(&_FriendTech.CallOpts)
}

// ProtocolFeePercent is a free data retrieval call binding the contract method 0xd6e6eb9f.
//
// Solidity: function protocolFeePercent() view returns(uint256)
func (_FriendTech *FriendTechCaller) ProtocolFeePercent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FriendTech.contract.Call(opts, &out, "protocolFeePercent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProtocolFeePercent is a free data retrieval call binding the contract method 0xd6e6eb9f.
//
// Solidity: function protocolFeePercent() view returns(uint256)
func (_FriendTech *FriendTechSession) ProtocolFeePercent() (*big.Int, error) {
	return _FriendTech.Contract.ProtocolFeePercent(&_FriendTech.CallOpts)
}

// ProtocolFeePercent is a free data retrieval call binding the contract method 0xd6e6eb9f.
//
// Solidity: function protocolFeePercent() view returns(uint256)
func (_FriendTech *FriendTechCallerSession) ProtocolFeePercent() (*big.Int, error) {
	return _FriendTech.Contract.ProtocolFeePercent(&_FriendTech.CallOpts)
}

// SharesBalance is a free data retrieval call binding the contract method 0x020235ff.
//
// Solidity: function sharesBalance(address , address ) view returns(uint256)
func (_FriendTech *FriendTechCaller) SharesBalance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FriendTech.contract.Call(opts, &out, "sharesBalance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SharesBalance is a free data retrieval call binding the contract method 0x020235ff.
//
// Solidity: function sharesBalance(address , address ) view returns(uint256)
func (_FriendTech *FriendTechSession) SharesBalance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _FriendTech.Contract.SharesBalance(&_FriendTech.CallOpts, arg0, arg1)
}

// SharesBalance is a free data retrieval call binding the contract method 0x020235ff.
//
// Solidity: function sharesBalance(address , address ) view returns(uint256)
func (_FriendTech *FriendTechCallerSession) SharesBalance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _FriendTech.Contract.SharesBalance(&_FriendTech.CallOpts, arg0, arg1)
}

// SharesSupply is a free data retrieval call binding the contract method 0xf9931be0.
//
// Solidity: function sharesSupply(address ) view returns(uint256)
func (_FriendTech *FriendTechCaller) SharesSupply(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FriendTech.contract.Call(opts, &out, "sharesSupply", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SharesSupply is a free data retrieval call binding the contract method 0xf9931be0.
//
// Solidity: function sharesSupply(address ) view returns(uint256)
func (_FriendTech *FriendTechSession) SharesSupply(arg0 common.Address) (*big.Int, error) {
	return _FriendTech.Contract.SharesSupply(&_FriendTech.CallOpts, arg0)
}

// SharesSupply is a free data retrieval call binding the contract method 0xf9931be0.
//
// Solidity: function sharesSupply(address ) view returns(uint256)
func (_FriendTech *FriendTechCallerSession) SharesSupply(arg0 common.Address) (*big.Int, error) {
	return _FriendTech.Contract.SharesSupply(&_FriendTech.CallOpts, arg0)
}

// SubjectFeePercent is a free data retrieval call binding the contract method 0x24dc441d.
//
// Solidity: function subjectFeePercent() view returns(uint256)
func (_FriendTech *FriendTechCaller) SubjectFeePercent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FriendTech.contract.Call(opts, &out, "subjectFeePercent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SubjectFeePercent is a free data retrieval call binding the contract method 0x24dc441d.
//
// Solidity: function subjectFeePercent() view returns(uint256)
func (_FriendTech *FriendTechSession) SubjectFeePercent() (*big.Int, error) {
	return _FriendTech.Contract.SubjectFeePercent(&_FriendTech.CallOpts)
}

// SubjectFeePercent is a free data retrieval call binding the contract method 0x24dc441d.
//
// Solidity: function subjectFeePercent() view returns(uint256)
func (_FriendTech *FriendTechCallerSession) SubjectFeePercent() (*big.Int, error) {
	return _FriendTech.Contract.SubjectFeePercent(&_FriendTech.CallOpts)
}

// BuyShares is a paid mutator transaction binding the contract method 0x6945b123.
//
// Solidity: function buyShares(address sharesSubject, uint256 amount) payable returns()
func (_FriendTech *FriendTechTransactor) BuyShares(opts *bind.TransactOpts, sharesSubject common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FriendTech.contract.Transact(opts, "buyShares", sharesSubject, amount)
}

// BuyShares is a paid mutator transaction binding the contract method 0x6945b123.
//
// Solidity: function buyShares(address sharesSubject, uint256 amount) payable returns()
func (_FriendTech *FriendTechSession) BuyShares(sharesSubject common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FriendTech.Contract.BuyShares(&_FriendTech.TransactOpts, sharesSubject, amount)
}

// BuyShares is a paid mutator transaction binding the contract method 0x6945b123.
//
// Solidity: function buyShares(address sharesSubject, uint256 amount) payable returns()
func (_FriendTech *FriendTechTransactorSession) BuyShares(sharesSubject common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FriendTech.Contract.BuyShares(&_FriendTech.TransactOpts, sharesSubject, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FriendTech *FriendTechTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FriendTech.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FriendTech *FriendTechSession) RenounceOwnership() (*types.Transaction, error) {
	return _FriendTech.Contract.RenounceOwnership(&_FriendTech.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FriendTech *FriendTechTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _FriendTech.Contract.RenounceOwnership(&_FriendTech.TransactOpts)
}

// SellShares is a paid mutator transaction binding the contract method 0xb51d0534.
//
// Solidity: function sellShares(address sharesSubject, uint256 amount) payable returns()
func (_FriendTech *FriendTechTransactor) SellShares(opts *bind.TransactOpts, sharesSubject common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FriendTech.contract.Transact(opts, "sellShares", sharesSubject, amount)
}

// SellShares is a paid mutator transaction binding the contract method 0xb51d0534.
//
// Solidity: function sellShares(address sharesSubject, uint256 amount) payable returns()
func (_FriendTech *FriendTechSession) SellShares(sharesSubject common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FriendTech.Contract.SellShares(&_FriendTech.TransactOpts, sharesSubject, amount)
}

// SellShares is a paid mutator transaction binding the contract method 0xb51d0534.
//
// Solidity: function sellShares(address sharesSubject, uint256 amount) payable returns()
func (_FriendTech *FriendTechTransactorSession) SellShares(sharesSubject common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FriendTech.Contract.SellShares(&_FriendTech.TransactOpts, sharesSubject, amount)
}

// SetFeeDestination is a paid mutator transaction binding the contract method 0xfbe53234.
//
// Solidity: function setFeeDestination(address _feeDestination) returns()
func (_FriendTech *FriendTechTransactor) SetFeeDestination(opts *bind.TransactOpts, _feeDestination common.Address) (*types.Transaction, error) {
	return _FriendTech.contract.Transact(opts, "setFeeDestination", _feeDestination)
}

// SetFeeDestination is a paid mutator transaction binding the contract method 0xfbe53234.
//
// Solidity: function setFeeDestination(address _feeDestination) returns()
func (_FriendTech *FriendTechSession) SetFeeDestination(_feeDestination common.Address) (*types.Transaction, error) {
	return _FriendTech.Contract.SetFeeDestination(&_FriendTech.TransactOpts, _feeDestination)
}

// SetFeeDestination is a paid mutator transaction binding the contract method 0xfbe53234.
//
// Solidity: function setFeeDestination(address _feeDestination) returns()
func (_FriendTech *FriendTechTransactorSession) SetFeeDestination(_feeDestination common.Address) (*types.Transaction, error) {
	return _FriendTech.Contract.SetFeeDestination(&_FriendTech.TransactOpts, _feeDestination)
}

// SetProtocolFeePercent is a paid mutator transaction binding the contract method 0xa4983421.
//
// Solidity: function setProtocolFeePercent(uint256 _feePercent) returns()
func (_FriendTech *FriendTechTransactor) SetProtocolFeePercent(opts *bind.TransactOpts, _feePercent *big.Int) (*types.Transaction, error) {
	return _FriendTech.contract.Transact(opts, "setProtocolFeePercent", _feePercent)
}

// SetProtocolFeePercent is a paid mutator transaction binding the contract method 0xa4983421.
//
// Solidity: function setProtocolFeePercent(uint256 _feePercent) returns()
func (_FriendTech *FriendTechSession) SetProtocolFeePercent(_feePercent *big.Int) (*types.Transaction, error) {
	return _FriendTech.Contract.SetProtocolFeePercent(&_FriendTech.TransactOpts, _feePercent)
}

// SetProtocolFeePercent is a paid mutator transaction binding the contract method 0xa4983421.
//
// Solidity: function setProtocolFeePercent(uint256 _feePercent) returns()
func (_FriendTech *FriendTechTransactorSession) SetProtocolFeePercent(_feePercent *big.Int) (*types.Transaction, error) {
	return _FriendTech.Contract.SetProtocolFeePercent(&_FriendTech.TransactOpts, _feePercent)
}

// SetSubjectFeePercent is a paid mutator transaction binding the contract method 0x5a8a764e.
//
// Solidity: function setSubjectFeePercent(uint256 _feePercent) returns()
func (_FriendTech *FriendTechTransactor) SetSubjectFeePercent(opts *bind.TransactOpts, _feePercent *big.Int) (*types.Transaction, error) {
	return _FriendTech.contract.Transact(opts, "setSubjectFeePercent", _feePercent)
}

// SetSubjectFeePercent is a paid mutator transaction binding the contract method 0x5a8a764e.
//
// Solidity: function setSubjectFeePercent(uint256 _feePercent) returns()
func (_FriendTech *FriendTechSession) SetSubjectFeePercent(_feePercent *big.Int) (*types.Transaction, error) {
	return _FriendTech.Contract.SetSubjectFeePercent(&_FriendTech.TransactOpts, _feePercent)
}

// SetSubjectFeePercent is a paid mutator transaction binding the contract method 0x5a8a764e.
//
// Solidity: function setSubjectFeePercent(uint256 _feePercent) returns()
func (_FriendTech *FriendTechTransactorSession) SetSubjectFeePercent(_feePercent *big.Int) (*types.Transaction, error) {
	return _FriendTech.Contract.SetSubjectFeePercent(&_FriendTech.TransactOpts, _feePercent)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FriendTech *FriendTechTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _FriendTech.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FriendTech *FriendTechSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FriendTech.Contract.TransferOwnership(&_FriendTech.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FriendTech *FriendTechTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FriendTech.Contract.TransferOwnership(&_FriendTech.TransactOpts, newOwner)
}

// FriendTechOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the FriendTech contract.
type FriendTechOwnershipTransferredIterator struct {
	Event *FriendTechOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FriendTechOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FriendTechOwnershipTransferred)
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
		it.Event = new(FriendTechOwnershipTransferred)
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
func (it *FriendTechOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FriendTechOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FriendTechOwnershipTransferred represents a OwnershipTransferred event raised by the FriendTech contract.
type FriendTechOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FriendTech *FriendTechFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FriendTechOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FriendTech.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FriendTechOwnershipTransferredIterator{contract: _FriendTech.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FriendTech *FriendTechFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FriendTechOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FriendTech.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FriendTechOwnershipTransferred)
				if err := _FriendTech.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_FriendTech *FriendTechFilterer) ParseOwnershipTransferred(log types.Log) (*FriendTechOwnershipTransferred, error) {
	event := new(FriendTechOwnershipTransferred)
	if err := _FriendTech.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FriendTechTradeIterator is returned from FilterTrade and is used to iterate over the raw logs and unpacked data for Trade events raised by the FriendTech contract.
type FriendTechTradeIterator struct {
	Event *FriendTechTrade // Event containing the contract specifics and raw log

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
func (it *FriendTechTradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FriendTechTrade)
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
		it.Event = new(FriendTechTrade)
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
func (it *FriendTechTradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FriendTechTradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FriendTechTrade represents a Trade event raised by the FriendTech contract.
type FriendTechTrade struct {
	Trader            common.Address
	Subject           common.Address
	IsBuy             bool
	ShareAmount       *big.Int
	EthAmount         *big.Int
	ProtocolEthAmount *big.Int
	SubjectEthAmount  *big.Int
	Supply            *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterTrade is a free log retrieval operation binding the contract event 0x2c76e7a47fd53e2854856ac3f0a5f3ee40d15cfaa82266357ea9779c486ab9c3.
//
// Solidity: event Trade(address trader, address subject, bool isBuy, uint256 shareAmount, uint256 ethAmount, uint256 protocolEthAmount, uint256 subjectEthAmount, uint256 supply)
func (_FriendTech *FriendTechFilterer) FilterTrade(opts *bind.FilterOpts) (*FriendTechTradeIterator, error) {

	logs, sub, err := _FriendTech.contract.FilterLogs(opts, "Trade")
	if err != nil {
		return nil, err
	}
	return &FriendTechTradeIterator{contract: _FriendTech.contract, event: "Trade", logs: logs, sub: sub}, nil
}

// WatchTrade is a free log subscription operation binding the contract event 0x2c76e7a47fd53e2854856ac3f0a5f3ee40d15cfaa82266357ea9779c486ab9c3.
//
// Solidity: event Trade(address trader, address subject, bool isBuy, uint256 shareAmount, uint256 ethAmount, uint256 protocolEthAmount, uint256 subjectEthAmount, uint256 supply)
func (_FriendTech *FriendTechFilterer) WatchTrade(opts *bind.WatchOpts, sink chan<- *FriendTechTrade) (event.Subscription, error) {

	logs, sub, err := _FriendTech.contract.WatchLogs(opts, "Trade")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FriendTechTrade)
				if err := _FriendTech.contract.UnpackLog(event, "Trade", log); err != nil {
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

// ParseTrade is a log parse operation binding the contract event 0x2c76e7a47fd53e2854856ac3f0a5f3ee40d15cfaa82266357ea9779c486ab9c3.
//
// Solidity: event Trade(address trader, address subject, bool isBuy, uint256 shareAmount, uint256 ethAmount, uint256 protocolEthAmount, uint256 subjectEthAmount, uint256 supply)
func (_FriendTech *FriendTechFilterer) ParseTrade(log types.Log) (*FriendTechTrade, error) {
	event := new(FriendTechTrade)
	if err := _FriendTech.contract.UnpackLog(event, "Trade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
