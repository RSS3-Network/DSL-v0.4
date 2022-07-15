// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package oneinch

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

// AggregationRouterV3SwapDescription is an auto generated low-level Go binding around an user-defined struct.
type AggregationRouterV3SwapDescription struct {
	SrcToken        common.Address
	DstToken        common.Address
	SrcReceiver     common.Address
	DstReceiver     common.Address
	Amount          *big.Int
	MinReturnAmount *big.Int
	Flags           *big.Int
	Permit          []byte
}

// OneinchMetaData contains all meta data concerning the Oneinch contract.
var OneinchMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"Error\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"dstToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"dstReceiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"spentAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"name\":\"Swapped\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"destroy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIAggregationExecutor\",\"name\":\"caller\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"srcReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dstReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"internalType\":\"structAggregationRouterV3.SwapDescription\",\"name\":\"desc\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"discountedSwap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLeft\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chiSpent\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"rescueFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIAggregationExecutor\",\"name\":\"caller\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"srcReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dstReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"internalType\":\"structAggregationRouterV3.SwapDescription\",\"name\":\"desc\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLeft\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"name\":\"unoswap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"pools\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"name\":\"unoswapWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// OneinchABI is the input ABI used to generate the binding from.
// Deprecated: Use OneinchMetaData.ABI instead.
var OneinchABI = OneinchMetaData.ABI

// Oneinch is an auto generated Go binding around an Ethereum contract.
type Oneinch struct {
	OneinchCaller     // Read-only binding to the contract
	OneinchTransactor // Write-only binding to the contract
	OneinchFilterer   // Log filterer for contract events
}

// OneinchCaller is an auto generated read-only Go binding around an Ethereum contract.
type OneinchCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneinchTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OneinchTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneinchFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OneinchFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneinchSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OneinchSession struct {
	Contract     *Oneinch          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OneinchCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OneinchCallerSession struct {
	Contract *OneinchCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OneinchTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OneinchTransactorSession struct {
	Contract     *OneinchTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OneinchRaw is an auto generated low-level Go binding around an Ethereum contract.
type OneinchRaw struct {
	Contract *Oneinch // Generic contract binding to access the raw methods on
}

// OneinchCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OneinchCallerRaw struct {
	Contract *OneinchCaller // Generic read-only contract binding to access the raw methods on
}

// OneinchTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OneinchTransactorRaw struct {
	Contract *OneinchTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOneinch creates a new instance of Oneinch, bound to a specific deployed contract.
func NewOneinch(address common.Address, backend bind.ContractBackend) (*Oneinch, error) {
	contract, err := bindOneinch(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Oneinch{OneinchCaller: OneinchCaller{contract: contract}, OneinchTransactor: OneinchTransactor{contract: contract}, OneinchFilterer: OneinchFilterer{contract: contract}}, nil
}

// NewOneinchCaller creates a new read-only instance of Oneinch, bound to a specific deployed contract.
func NewOneinchCaller(address common.Address, caller bind.ContractCaller) (*OneinchCaller, error) {
	contract, err := bindOneinch(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OneinchCaller{contract: contract}, nil
}

// NewOneinchTransactor creates a new write-only instance of Oneinch, bound to a specific deployed contract.
func NewOneinchTransactor(address common.Address, transactor bind.ContractTransactor) (*OneinchTransactor, error) {
	contract, err := bindOneinch(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OneinchTransactor{contract: contract}, nil
}

// NewOneinchFilterer creates a new log filterer instance of Oneinch, bound to a specific deployed contract.
func NewOneinchFilterer(address common.Address, filterer bind.ContractFilterer) (*OneinchFilterer, error) {
	contract, err := bindOneinch(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OneinchFilterer{contract: contract}, nil
}

// bindOneinch binds a generic wrapper to an already deployed contract.
func bindOneinch(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OneinchABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oneinch *OneinchRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Oneinch.Contract.OneinchCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oneinch *OneinchRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oneinch.Contract.OneinchTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oneinch *OneinchRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oneinch.Contract.OneinchTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oneinch *OneinchCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Oneinch.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oneinch *OneinchTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oneinch.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oneinch *OneinchTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oneinch.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Oneinch *OneinchCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Oneinch.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Oneinch *OneinchSession) Owner() (common.Address, error) {
	return _Oneinch.Contract.Owner(&_Oneinch.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Oneinch *OneinchCallerSession) Owner() (common.Address, error) {
	return _Oneinch.Contract.Owner(&_Oneinch.CallOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_Oneinch *OneinchTransactor) Destroy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oneinch.contract.Transact(opts, "destroy")
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_Oneinch *OneinchSession) Destroy() (*types.Transaction, error) {
	return _Oneinch.Contract.Destroy(&_Oneinch.TransactOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_Oneinch *OneinchTransactorSession) Destroy() (*types.Transaction, error) {
	return _Oneinch.Contract.Destroy(&_Oneinch.TransactOpts)
}

// DiscountedSwap is a paid mutator transaction binding the contract method 0x6c4a483e.
//
// Solidity: function discountedSwap(address caller, (address,address,address,address,uint256,uint256,uint256,bytes) desc, bytes data) payable returns(uint256 returnAmount, uint256 gasLeft, uint256 chiSpent)
func (_Oneinch *OneinchTransactor) DiscountedSwap(opts *bind.TransactOpts, caller common.Address, desc AggregationRouterV3SwapDescription, data []byte) (*types.Transaction, error) {
	return _Oneinch.contract.Transact(opts, "discountedSwap", caller, desc, data)
}

// DiscountedSwap is a paid mutator transaction binding the contract method 0x6c4a483e.
//
// Solidity: function discountedSwap(address caller, (address,address,address,address,uint256,uint256,uint256,bytes) desc, bytes data) payable returns(uint256 returnAmount, uint256 gasLeft, uint256 chiSpent)
func (_Oneinch *OneinchSession) DiscountedSwap(caller common.Address, desc AggregationRouterV3SwapDescription, data []byte) (*types.Transaction, error) {
	return _Oneinch.Contract.DiscountedSwap(&_Oneinch.TransactOpts, caller, desc, data)
}

// DiscountedSwap is a paid mutator transaction binding the contract method 0x6c4a483e.
//
// Solidity: function discountedSwap(address caller, (address,address,address,address,uint256,uint256,uint256,bytes) desc, bytes data) payable returns(uint256 returnAmount, uint256 gasLeft, uint256 chiSpent)
func (_Oneinch *OneinchTransactorSession) DiscountedSwap(caller common.Address, desc AggregationRouterV3SwapDescription, data []byte) (*types.Transaction, error) {
	return _Oneinch.Contract.DiscountedSwap(&_Oneinch.TransactOpts, caller, desc, data)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Oneinch *OneinchTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oneinch.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Oneinch *OneinchSession) RenounceOwnership() (*types.Transaction, error) {
	return _Oneinch.Contract.RenounceOwnership(&_Oneinch.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Oneinch *OneinchTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Oneinch.Contract.RenounceOwnership(&_Oneinch.TransactOpts)
}

// RescueFunds is a paid mutator transaction binding the contract method 0x78e3214f.
//
// Solidity: function rescueFunds(address token, uint256 amount) returns()
func (_Oneinch *OneinchTransactor) RescueFunds(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Oneinch.contract.Transact(opts, "rescueFunds", token, amount)
}

// RescueFunds is a paid mutator transaction binding the contract method 0x78e3214f.
//
// Solidity: function rescueFunds(address token, uint256 amount) returns()
func (_Oneinch *OneinchSession) RescueFunds(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Oneinch.Contract.RescueFunds(&_Oneinch.TransactOpts, token, amount)
}

// RescueFunds is a paid mutator transaction binding the contract method 0x78e3214f.
//
// Solidity: function rescueFunds(address token, uint256 amount) returns()
func (_Oneinch *OneinchTransactorSession) RescueFunds(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Oneinch.Contract.RescueFunds(&_Oneinch.TransactOpts, token, amount)
}

// Swap is a paid mutator transaction binding the contract method 0x7c025200.
//
// Solidity: function swap(address caller, (address,address,address,address,uint256,uint256,uint256,bytes) desc, bytes data) payable returns(uint256 returnAmount, uint256 gasLeft)
func (_Oneinch *OneinchTransactor) Swap(opts *bind.TransactOpts, caller common.Address, desc AggregationRouterV3SwapDescription, data []byte) (*types.Transaction, error) {
	return _Oneinch.contract.Transact(opts, "swap", caller, desc, data)
}

// Swap is a paid mutator transaction binding the contract method 0x7c025200.
//
// Solidity: function swap(address caller, (address,address,address,address,uint256,uint256,uint256,bytes) desc, bytes data) payable returns(uint256 returnAmount, uint256 gasLeft)
func (_Oneinch *OneinchSession) Swap(caller common.Address, desc AggregationRouterV3SwapDescription, data []byte) (*types.Transaction, error) {
	return _Oneinch.Contract.Swap(&_Oneinch.TransactOpts, caller, desc, data)
}

// Swap is a paid mutator transaction binding the contract method 0x7c025200.
//
// Solidity: function swap(address caller, (address,address,address,address,uint256,uint256,uint256,bytes) desc, bytes data) payable returns(uint256 returnAmount, uint256 gasLeft)
func (_Oneinch *OneinchTransactorSession) Swap(caller common.Address, desc AggregationRouterV3SwapDescription, data []byte) (*types.Transaction, error) {
	return _Oneinch.Contract.Swap(&_Oneinch.TransactOpts, caller, desc, data)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Oneinch *OneinchTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Oneinch.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Oneinch *OneinchSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Oneinch.Contract.TransferOwnership(&_Oneinch.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Oneinch *OneinchTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Oneinch.Contract.TransferOwnership(&_Oneinch.TransactOpts, newOwner)
}

// Unoswap is a paid mutator transaction binding the contract method 0x2e95b6c8.
//
// Solidity: function unoswap(address srcToken, uint256 amount, uint256 minReturn, bytes32[] ) payable returns(uint256 returnAmount)
func (_Oneinch *OneinchTransactor) Unoswap(opts *bind.TransactOpts, srcToken common.Address, amount *big.Int, minReturn *big.Int, arg3 [][32]byte) (*types.Transaction, error) {
	return _Oneinch.contract.Transact(opts, "unoswap", srcToken, amount, minReturn, arg3)
}

// Unoswap is a paid mutator transaction binding the contract method 0x2e95b6c8.
//
// Solidity: function unoswap(address srcToken, uint256 amount, uint256 minReturn, bytes32[] ) payable returns(uint256 returnAmount)
func (_Oneinch *OneinchSession) Unoswap(srcToken common.Address, amount *big.Int, minReturn *big.Int, arg3 [][32]byte) (*types.Transaction, error) {
	return _Oneinch.Contract.Unoswap(&_Oneinch.TransactOpts, srcToken, amount, minReturn, arg3)
}

// Unoswap is a paid mutator transaction binding the contract method 0x2e95b6c8.
//
// Solidity: function unoswap(address srcToken, uint256 amount, uint256 minReturn, bytes32[] ) payable returns(uint256 returnAmount)
func (_Oneinch *OneinchTransactorSession) Unoswap(srcToken common.Address, amount *big.Int, minReturn *big.Int, arg3 [][32]byte) (*types.Transaction, error) {
	return _Oneinch.Contract.Unoswap(&_Oneinch.TransactOpts, srcToken, amount, minReturn, arg3)
}

// UnoswapWithPermit is a paid mutator transaction binding the contract method 0xa1251d75.
//
// Solidity: function unoswapWithPermit(address srcToken, uint256 amount, uint256 minReturn, bytes32[] pools, bytes permit) payable returns(uint256 returnAmount)
func (_Oneinch *OneinchTransactor) UnoswapWithPermit(opts *bind.TransactOpts, srcToken common.Address, amount *big.Int, minReturn *big.Int, pools [][32]byte, permit []byte) (*types.Transaction, error) {
	return _Oneinch.contract.Transact(opts, "unoswapWithPermit", srcToken, amount, minReturn, pools, permit)
}

// UnoswapWithPermit is a paid mutator transaction binding the contract method 0xa1251d75.
//
// Solidity: function unoswapWithPermit(address srcToken, uint256 amount, uint256 minReturn, bytes32[] pools, bytes permit) payable returns(uint256 returnAmount)
func (_Oneinch *OneinchSession) UnoswapWithPermit(srcToken common.Address, amount *big.Int, minReturn *big.Int, pools [][32]byte, permit []byte) (*types.Transaction, error) {
	return _Oneinch.Contract.UnoswapWithPermit(&_Oneinch.TransactOpts, srcToken, amount, minReturn, pools, permit)
}

// UnoswapWithPermit is a paid mutator transaction binding the contract method 0xa1251d75.
//
// Solidity: function unoswapWithPermit(address srcToken, uint256 amount, uint256 minReturn, bytes32[] pools, bytes permit) payable returns(uint256 returnAmount)
func (_Oneinch *OneinchTransactorSession) UnoswapWithPermit(srcToken common.Address, amount *big.Int, minReturn *big.Int, pools [][32]byte, permit []byte) (*types.Transaction, error) {
	return _Oneinch.Contract.UnoswapWithPermit(&_Oneinch.TransactOpts, srcToken, amount, minReturn, pools, permit)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Oneinch *OneinchTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oneinch.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Oneinch *OneinchSession) Receive() (*types.Transaction, error) {
	return _Oneinch.Contract.Receive(&_Oneinch.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Oneinch *OneinchTransactorSession) Receive() (*types.Transaction, error) {
	return _Oneinch.Contract.Receive(&_Oneinch.TransactOpts)
}

// OneinchErrorIterator is returned from FilterError and is used to iterate over the raw logs and unpacked data for Error events raised by the Oneinch contract.
type OneinchErrorIterator struct {
	Event *OneinchError // Event containing the contract specifics and raw log

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
func (it *OneinchErrorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OneinchError)
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
		it.Event = new(OneinchError)
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
func (it *OneinchErrorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OneinchErrorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OneinchError represents a Error event raised by the Oneinch contract.
type OneinchError struct {
	Reason string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterError is a free log retrieval operation binding the contract event 0x08c379a0afcc32b1a39302f7cb8073359698411ab5fd6e3edb2c02c0b5fba8aa.
//
// Solidity: event Error(string reason)
func (_Oneinch *OneinchFilterer) FilterError(opts *bind.FilterOpts) (*OneinchErrorIterator, error) {

	logs, sub, err := _Oneinch.contract.FilterLogs(opts, "Error")
	if err != nil {
		return nil, err
	}
	return &OneinchErrorIterator{contract: _Oneinch.contract, event: "Error", logs: logs, sub: sub}, nil
}

// WatchError is a free log subscription operation binding the contract event 0x08c379a0afcc32b1a39302f7cb8073359698411ab5fd6e3edb2c02c0b5fba8aa.
//
// Solidity: event Error(string reason)
func (_Oneinch *OneinchFilterer) WatchError(opts *bind.WatchOpts, sink chan<- *OneinchError) (event.Subscription, error) {

	logs, sub, err := _Oneinch.contract.WatchLogs(opts, "Error")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OneinchError)
				if err := _Oneinch.contract.UnpackLog(event, "Error", log); err != nil {
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

// ParseError is a log parse operation binding the contract event 0x08c379a0afcc32b1a39302f7cb8073359698411ab5fd6e3edb2c02c0b5fba8aa.
//
// Solidity: event Error(string reason)
func (_Oneinch *OneinchFilterer) ParseError(log types.Log) (*OneinchError, error) {
	event := new(OneinchError)
	if err := _Oneinch.contract.UnpackLog(event, "Error", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OneinchOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Oneinch contract.
type OneinchOwnershipTransferredIterator struct {
	Event *OneinchOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OneinchOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OneinchOwnershipTransferred)
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
		it.Event = new(OneinchOwnershipTransferred)
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
func (it *OneinchOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OneinchOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OneinchOwnershipTransferred represents a OwnershipTransferred event raised by the Oneinch contract.
type OneinchOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Oneinch *OneinchFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OneinchOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Oneinch.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OneinchOwnershipTransferredIterator{contract: _Oneinch.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Oneinch *OneinchFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OneinchOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Oneinch.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OneinchOwnershipTransferred)
				if err := _Oneinch.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Oneinch *OneinchFilterer) ParseOwnershipTransferred(log types.Log) (*OneinchOwnershipTransferred, error) {
	event := new(OneinchOwnershipTransferred)
	if err := _Oneinch.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OneinchSwappedIterator is returned from FilterSwapped and is used to iterate over the raw logs and unpacked data for Swapped events raised by the Oneinch contract.
type OneinchSwappedIterator struct {
	Event *OneinchSwapped // Event containing the contract specifics and raw log

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
func (it *OneinchSwappedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OneinchSwapped)
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
		it.Event = new(OneinchSwapped)
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
func (it *OneinchSwappedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OneinchSwappedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OneinchSwapped represents a Swapped event raised by the Oneinch contract.
type OneinchSwapped struct {
	Sender       common.Address
	SrcToken     common.Address
	DstToken     common.Address
	DstReceiver  common.Address
	SpentAmount  *big.Int
	ReturnAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSwapped is a free log retrieval operation binding the contract event 0xd6d4f5681c246c9f42c203e287975af1601f8df8035a9251f79aab5c8f09e2f8.
//
// Solidity: event Swapped(address sender, address srcToken, address dstToken, address dstReceiver, uint256 spentAmount, uint256 returnAmount)
func (_Oneinch *OneinchFilterer) FilterSwapped(opts *bind.FilterOpts) (*OneinchSwappedIterator, error) {

	logs, sub, err := _Oneinch.contract.FilterLogs(opts, "Swapped")
	if err != nil {
		return nil, err
	}
	return &OneinchSwappedIterator{contract: _Oneinch.contract, event: "Swapped", logs: logs, sub: sub}, nil
}

// WatchSwapped is a free log subscription operation binding the contract event 0xd6d4f5681c246c9f42c203e287975af1601f8df8035a9251f79aab5c8f09e2f8.
//
// Solidity: event Swapped(address sender, address srcToken, address dstToken, address dstReceiver, uint256 spentAmount, uint256 returnAmount)
func (_Oneinch *OneinchFilterer) WatchSwapped(opts *bind.WatchOpts, sink chan<- *OneinchSwapped) (event.Subscription, error) {

	logs, sub, err := _Oneinch.contract.WatchLogs(opts, "Swapped")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OneinchSwapped)
				if err := _Oneinch.contract.UnpackLog(event, "Swapped", log); err != nil {
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

// ParseSwapped is a log parse operation binding the contract event 0xd6d4f5681c246c9f42c203e287975af1601f8df8035a9251f79aab5c8f09e2f8.
//
// Solidity: event Swapped(address sender, address srcToken, address dstToken, address dstReceiver, uint256 spentAmount, uint256 returnAmount)
func (_Oneinch *OneinchFilterer) ParseSwapped(log types.Log) (*OneinchSwapped, error) {
	event := new(OneinchSwapped)
	if err := _Oneinch.contract.UnpackLog(event, "Swapped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
