// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dodo

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

// PMMPricingPMMState is an auto generated low-level Go binding around an user-defined struct.
type PMMPricingPMMState struct {
	I  *big.Int
	K  *big.Int
	B  *big.Int
	Q  *big.Int
	B0 *big.Int
	Q0 *big.Int
	R  uint8
}

// VendingMachineMetaData contains all meta data concerning the VendingMachine contract.
var VendingMachineMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"increaseShares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalShares\",\"type\":\"uint256\"}],\"name\":\"BuyShares\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"assetTo\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"baseAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quoteAmount\",\"type\":\"uint256\"}],\"name\":\"DODOFlashLoan\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"DODOSwap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"decreaseShares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalShares\",\"type\":\"uint256\"}],\"name\":\"SellShares\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMIT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_BASE_PRICE_CUMULATIVE_LAST_\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_BASE_RESERVE_\",\"outputs\":[{\"internalType\":\"uint112\",\"name\":\"\",\"type\":\"uint112\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_BASE_TOKEN_\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_BLOCK_TIMESTAMP_LAST_\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_IS_OPEN_TWAP_\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_I_\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_K_\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_LP_FEE_RATE_\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_MAINTAINER_\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_MT_FEE_RATE_MODEL_\",\"outputs\":[{\"internalType\":\"contractIFeeRateModel\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_QUOTE_RESERVE_\",\"outputs\":[{\"internalType\":\"uint112\",\"name\":\"\",\"type\":\"uint112\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_QUOTE_TOKEN_\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"addressToShortString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"buyShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseInput\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quoteInput\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"baseAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quoteAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"assetTo\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"flashLoan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBaseInput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"input\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMidPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"midPrice\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPMMState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"K\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"B\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Q\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"B0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Q0\",\"type\":\"uint256\"},{\"internalType\":\"enumPMMPricing.RState\",\"name\":\"R\",\"type\":\"uint8\"}],\"internalType\":\"structPMMPricing.PMMState\",\"name\":\"state\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPMMStateForCall\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"K\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"B\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Q\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"B0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Q0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"R\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getQuoteInput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"input\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserFeeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"lpFeeRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mtFeeRate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVaultReserve\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"baseReserve\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quoteReserve\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"maintainer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"baseTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quoteTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lpFeeRate\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"mtFeeRateModel\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"k\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isOpenTWAP\",\"type\":\"bool\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"payBaseAmount\",\"type\":\"uint256\"}],\"name\":\"querySellBase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"receiveQuoteAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mtFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"payQuoteAmount\",\"type\":\"uint256\"}],\"name\":\"querySellQuote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"receiveBaseAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mtFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"sellBase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"receiveQuoteAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"sellQuote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"receiveBaseAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shareAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"baseMinAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quoteMinAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"sellShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"baseAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quoteAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sync\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// VendingMachineABI is the input ABI used to generate the binding from.
// Deprecated: Use VendingMachineMetaData.ABI instead.
var VendingMachineABI = VendingMachineMetaData.ABI

// VendingMachine is an auto generated Go binding around an Ethereum contract.
type VendingMachine struct {
	VendingMachineCaller     // Read-only binding to the contract
	VendingMachineTransactor // Write-only binding to the contract
	VendingMachineFilterer   // Log filterer for contract events
}

// VendingMachineCaller is an auto generated read-only Go binding around an Ethereum contract.
type VendingMachineCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VendingMachineTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VendingMachineTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VendingMachineFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VendingMachineFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VendingMachineSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VendingMachineSession struct {
	Contract     *VendingMachine   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VendingMachineCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VendingMachineCallerSession struct {
	Contract *VendingMachineCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// VendingMachineTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VendingMachineTransactorSession struct {
	Contract     *VendingMachineTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// VendingMachineRaw is an auto generated low-level Go binding around an Ethereum contract.
type VendingMachineRaw struct {
	Contract *VendingMachine // Generic contract binding to access the raw methods on
}

// VendingMachineCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VendingMachineCallerRaw struct {
	Contract *VendingMachineCaller // Generic read-only contract binding to access the raw methods on
}

// VendingMachineTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VendingMachineTransactorRaw struct {
	Contract *VendingMachineTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVendingMachine creates a new instance of VendingMachine, bound to a specific deployed contract.
func NewVendingMachine(address common.Address, backend bind.ContractBackend) (*VendingMachine, error) {
	contract, err := bindVendingMachine(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VendingMachine{VendingMachineCaller: VendingMachineCaller{contract: contract}, VendingMachineTransactor: VendingMachineTransactor{contract: contract}, VendingMachineFilterer: VendingMachineFilterer{contract: contract}}, nil
}

// NewVendingMachineCaller creates a new read-only instance of VendingMachine, bound to a specific deployed contract.
func NewVendingMachineCaller(address common.Address, caller bind.ContractCaller) (*VendingMachineCaller, error) {
	contract, err := bindVendingMachine(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VendingMachineCaller{contract: contract}, nil
}

// NewVendingMachineTransactor creates a new write-only instance of VendingMachine, bound to a specific deployed contract.
func NewVendingMachineTransactor(address common.Address, transactor bind.ContractTransactor) (*VendingMachineTransactor, error) {
	contract, err := bindVendingMachine(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VendingMachineTransactor{contract: contract}, nil
}

// NewVendingMachineFilterer creates a new log filterer instance of VendingMachine, bound to a specific deployed contract.
func NewVendingMachineFilterer(address common.Address, filterer bind.ContractFilterer) (*VendingMachineFilterer, error) {
	contract, err := bindVendingMachine(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VendingMachineFilterer{contract: contract}, nil
}

// bindVendingMachine binds a generic wrapper to an already deployed contract.
func bindVendingMachine(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VendingMachineABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VendingMachine *VendingMachineRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VendingMachine.Contract.VendingMachineCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VendingMachine *VendingMachineRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VendingMachine.Contract.VendingMachineTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VendingMachine *VendingMachineRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VendingMachine.Contract.VendingMachineTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VendingMachine *VendingMachineCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VendingMachine.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VendingMachine *VendingMachineTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VendingMachine.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VendingMachine *VendingMachineTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VendingMachine.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_VendingMachine *VendingMachineCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _VendingMachine.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_VendingMachine *VendingMachineSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _VendingMachine.Contract.DOMAINSEPARATOR(&_VendingMachine.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_VendingMachine *VendingMachineCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _VendingMachine.Contract.DOMAINSEPARATOR(&_VendingMachine.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_VendingMachine *VendingMachineCaller) PERMITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _VendingMachine.contract.Call(opts, &out, "PERMIT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_VendingMachine *VendingMachineSession) PERMITTYPEHASH() ([32]byte, error) {
	return _VendingMachine.Contract.PERMITTYPEHASH(&_VendingMachine.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_VendingMachine *VendingMachineCallerSession) PERMITTYPEHASH() ([32]byte, error) {
	return _VendingMachine.Contract.PERMITTYPEHASH(&_VendingMachine.CallOpts)
}

// BASEPRICECUMULATIVELAST is a free data retrieval call binding the contract method 0xfe24cb7f.
//
// Solidity: function _BASE_PRICE_CUMULATIVE_LAST_() view returns(uint256)
func (_VendingMachine *VendingMachineCaller) BASEPRICECUMULATIVELAST(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VendingMachine.contract.Call(opts, &out, "_BASE_PRICE_CUMULATIVE_LAST_")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BASEPRICECUMULATIVELAST is a free data retrieval call binding the contract method 0xfe24cb7f.
//
// Solidity: function _BASE_PRICE_CUMULATIVE_LAST_() view returns(uint256)
func (_VendingMachine *VendingMachineSession) BASEPRICECUMULATIVELAST() (*big.Int, error) {
	return _VendingMachine.Contract.BASEPRICECUMULATIVELAST(&_VendingMachine.CallOpts)
}

// BASEPRICECUMULATIVELAST is a free data retrieval call binding the contract method 0xfe24cb7f.
//
// Solidity: function _BASE_PRICE_CUMULATIVE_LAST_() view returns(uint256)
func (_VendingMachine *VendingMachineCallerSession) BASEPRICECUMULATIVELAST() (*big.Int, error) {
	return _VendingMachine.Contract.BASEPRICECUMULATIVELAST(&_VendingMachine.CallOpts)
}

// BASERESERVE is a free data retrieval call binding the contract method 0x7d721504.
//
// Solidity: function _BASE_RESERVE_() view returns(uint112)
func (_VendingMachine *VendingMachineCaller) BASERESERVE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VendingMachine.contract.Call(opts, &out, "_BASE_RESERVE_")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BASERESERVE is a free data retrieval call binding the contract method 0x7d721504.
//
// Solidity: function _BASE_RESERVE_() view returns(uint112)
func (_VendingMachine *VendingMachineSession) BASERESERVE() (*big.Int, error) {
	return _VendingMachine.Contract.BASERESERVE(&_VendingMachine.CallOpts)
}

// BASERESERVE is a free data retrieval call binding the contract method 0x7d721504.
//
// Solidity: function _BASE_RESERVE_() view returns(uint112)
func (_VendingMachine *VendingMachineCallerSession) BASERESERVE() (*big.Int, error) {
	return _VendingMachine.Contract.BASERESERVE(&_VendingMachine.CallOpts)
}

// BASETOKEN is a free data retrieval call binding the contract method 0x4a248d2a.
//
// Solidity: function _BASE_TOKEN_() view returns(address)
func (_VendingMachine *VendingMachineCaller) BASETOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VendingMachine.contract.Call(opts, &out, "_BASE_TOKEN_")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BASETOKEN is a free data retrieval call binding the contract method 0x4a248d2a.
//
// Solidity: function _BASE_TOKEN_() view returns(address)
func (_VendingMachine *VendingMachineSession) BASETOKEN() (common.Address, error) {
	return _VendingMachine.Contract.BASETOKEN(&_VendingMachine.CallOpts)
}

// BASETOKEN is a free data retrieval call binding the contract method 0x4a248d2a.
//
// Solidity: function _BASE_TOKEN_() view returns(address)
func (_VendingMachine *VendingMachineCallerSession) BASETOKEN() (common.Address, error) {
	return _VendingMachine.Contract.BASETOKEN(&_VendingMachine.CallOpts)
}

// BLOCKTIMESTAMPLAST is a free data retrieval call binding the contract method 0x880a4d87.
//
// Solidity: function _BLOCK_TIMESTAMP_LAST_() view returns(uint32)
func (_VendingMachine *VendingMachineCaller) BLOCKTIMESTAMPLAST(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _VendingMachine.contract.Call(opts, &out, "_BLOCK_TIMESTAMP_LAST_")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// BLOCKTIMESTAMPLAST is a free data retrieval call binding the contract method 0x880a4d87.
//
// Solidity: function _BLOCK_TIMESTAMP_LAST_() view returns(uint32)
func (_VendingMachine *VendingMachineSession) BLOCKTIMESTAMPLAST() (uint32, error) {
	return _VendingMachine.Contract.BLOCKTIMESTAMPLAST(&_VendingMachine.CallOpts)
}

// BLOCKTIMESTAMPLAST is a free data retrieval call binding the contract method 0x880a4d87.
//
// Solidity: function _BLOCK_TIMESTAMP_LAST_() view returns(uint32)
func (_VendingMachine *VendingMachineCallerSession) BLOCKTIMESTAMPLAST() (uint32, error) {
	return _VendingMachine.Contract.BLOCKTIMESTAMPLAST(&_VendingMachine.CallOpts)
}

// ISOPENTWAP is a free data retrieval call binding the contract method 0x2df6cb48.
//
// Solidity: function _IS_OPEN_TWAP_() view returns(bool)
func (_VendingMachine *VendingMachineCaller) ISOPENTWAP(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _VendingMachine.contract.Call(opts, &out, "_IS_OPEN_TWAP_")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ISOPENTWAP is a free data retrieval call binding the contract method 0x2df6cb48.
//
// Solidity: function _IS_OPEN_TWAP_() view returns(bool)
func (_VendingMachine *VendingMachineSession) ISOPENTWAP() (bool, error) {
	return _VendingMachine.Contract.ISOPENTWAP(&_VendingMachine.CallOpts)
}

// ISOPENTWAP is a free data retrieval call binding the contract method 0x2df6cb48.
//
// Solidity: function _IS_OPEN_TWAP_() view returns(bool)
func (_VendingMachine *VendingMachineCallerSession) ISOPENTWAP() (bool, error) {
	return _VendingMachine.Contract.ISOPENTWAP(&_VendingMachine.CallOpts)
}

// I is a free data retrieval call binding the contract method 0xf811d692.
//
// Solidity: function _I_() view returns(uint256)
func (_VendingMachine *VendingMachineCaller) I(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VendingMachine.contract.Call(opts, &out, "_I_")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// I is a free data retrieval call binding the contract method 0xf811d692.
//
// Solidity: function _I_() view returns(uint256)
func (_VendingMachine *VendingMachineSession) I() (*big.Int, error) {
	return _VendingMachine.Contract.I(&_VendingMachine.CallOpts)
}

// I is a free data retrieval call binding the contract method 0xf811d692.
//
// Solidity: function _I_() view returns(uint256)
func (_VendingMachine *VendingMachineCallerSession) I() (*big.Int, error) {
	return _VendingMachine.Contract.I(&_VendingMachine.CallOpts)
}

// K is a free data retrieval call binding the contract method 0xec2fd46d.
//
// Solidity: function _K_() view returns(uint256)
func (_VendingMachine *VendingMachineCaller) K(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VendingMachine.contract.Call(opts, &out, "_K_")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// K is a free data retrieval call binding the contract method 0xec2fd46d.
//
// Solidity: function _K_() view returns(uint256)
func (_VendingMachine *VendingMachineSession) K() (*big.Int, error) {
	return _VendingMachine.Contract.K(&_VendingMachine.CallOpts)
}

// K is a free data retrieval call binding the contract method 0xec2fd46d.
//
// Solidity: function _K_() view returns(uint256)
func (_VendingMachine *VendingMachineCallerSession) K() (*big.Int, error) {
	return _VendingMachine.Contract.K(&_VendingMachine.CallOpts)
}

// LPFEERATE is a free data retrieval call binding the contract method 0xab44a7a3.
//
// Solidity: function _LP_FEE_RATE_() view returns(uint256)
func (_VendingMachine *VendingMachineCaller) LPFEERATE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VendingMachine.contract.Call(opts, &out, "_LP_FEE_RATE_")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LPFEERATE is a free data retrieval call binding the contract method 0xab44a7a3.
//
// Solidity: function _LP_FEE_RATE_() view returns(uint256)
func (_VendingMachine *VendingMachineSession) LPFEERATE() (*big.Int, error) {
	return _VendingMachine.Contract.LPFEERATE(&_VendingMachine.CallOpts)
}

// LPFEERATE is a free data retrieval call binding the contract method 0xab44a7a3.
//
// Solidity: function _LP_FEE_RATE_() view returns(uint256)
func (_VendingMachine *VendingMachineCallerSession) LPFEERATE() (*big.Int, error) {
	return _VendingMachine.Contract.LPFEERATE(&_VendingMachine.CallOpts)
}

// MAINTAINER is a free data retrieval call binding the contract method 0x4322ec83.
//
// Solidity: function _MAINTAINER_() view returns(address)
func (_VendingMachine *VendingMachineCaller) MAINTAINER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VendingMachine.contract.Call(opts, &out, "_MAINTAINER_")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MAINTAINER is a free data retrieval call binding the contract method 0x4322ec83.
//
// Solidity: function _MAINTAINER_() view returns(address)
func (_VendingMachine *VendingMachineSession) MAINTAINER() (common.Address, error) {
	return _VendingMachine.Contract.MAINTAINER(&_VendingMachine.CallOpts)
}

// MAINTAINER is a free data retrieval call binding the contract method 0x4322ec83.
//
// Solidity: function _MAINTAINER_() view returns(address)
func (_VendingMachine *VendingMachineCallerSession) MAINTAINER() (common.Address, error) {
	return _VendingMachine.Contract.MAINTAINER(&_VendingMachine.CallOpts)
}

// MTFEERATEMODEL is a free data retrieval call binding the contract method 0xf6b06e70.
//
// Solidity: function _MT_FEE_RATE_MODEL_() view returns(address)
func (_VendingMachine *VendingMachineCaller) MTFEERATEMODEL(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VendingMachine.contract.Call(opts, &out, "_MT_FEE_RATE_MODEL_")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MTFEERATEMODEL is a free data retrieval call binding the contract method 0xf6b06e70.
//
// Solidity: function _MT_FEE_RATE_MODEL_() view returns(address)
func (_VendingMachine *VendingMachineSession) MTFEERATEMODEL() (common.Address, error) {
	return _VendingMachine.Contract.MTFEERATEMODEL(&_VendingMachine.CallOpts)
}

// MTFEERATEMODEL is a free data retrieval call binding the contract method 0xf6b06e70.
//
// Solidity: function _MT_FEE_RATE_MODEL_() view returns(address)
func (_VendingMachine *VendingMachineCallerSession) MTFEERATEMODEL() (common.Address, error) {
	return _VendingMachine.Contract.MTFEERATEMODEL(&_VendingMachine.CallOpts)
}

// QUOTERESERVE is a free data retrieval call binding the contract method 0xbbf5ce78.
//
// Solidity: function _QUOTE_RESERVE_() view returns(uint112)
func (_VendingMachine *VendingMachineCaller) QUOTERESERVE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VendingMachine.contract.Call(opts, &out, "_QUOTE_RESERVE_")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QUOTERESERVE is a free data retrieval call binding the contract method 0xbbf5ce78.
//
// Solidity: function _QUOTE_RESERVE_() view returns(uint112)
func (_VendingMachine *VendingMachineSession) QUOTERESERVE() (*big.Int, error) {
	return _VendingMachine.Contract.QUOTERESERVE(&_VendingMachine.CallOpts)
}

// QUOTERESERVE is a free data retrieval call binding the contract method 0xbbf5ce78.
//
// Solidity: function _QUOTE_RESERVE_() view returns(uint112)
func (_VendingMachine *VendingMachineCallerSession) QUOTERESERVE() (*big.Int, error) {
	return _VendingMachine.Contract.QUOTERESERVE(&_VendingMachine.CallOpts)
}

// QUOTETOKEN is a free data retrieval call binding the contract method 0xd4b97046.
//
// Solidity: function _QUOTE_TOKEN_() view returns(address)
func (_VendingMachine *VendingMachineCaller) QUOTETOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VendingMachine.contract.Call(opts, &out, "_QUOTE_TOKEN_")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// QUOTETOKEN is a free data retrieval call binding the contract method 0xd4b97046.
//
// Solidity: function _QUOTE_TOKEN_() view returns(address)
func (_VendingMachine *VendingMachineSession) QUOTETOKEN() (common.Address, error) {
	return _VendingMachine.Contract.QUOTETOKEN(&_VendingMachine.CallOpts)
}

// QUOTETOKEN is a free data retrieval call binding the contract method 0xd4b97046.
//
// Solidity: function _QUOTE_TOKEN_() view returns(address)
func (_VendingMachine *VendingMachineCallerSession) QUOTETOKEN() (common.Address, error) {
	return _VendingMachine.Contract.QUOTETOKEN(&_VendingMachine.CallOpts)
}

// AddressToShortString is a free data retrieval call binding the contract method 0x17101940.
//
// Solidity: function addressToShortString(address _addr) pure returns(string)
func (_VendingMachine *VendingMachineCaller) AddressToShortString(opts *bind.CallOpts, _addr common.Address) (string, error) {
	var out []interface{}
	err := _VendingMachine.contract.Call(opts, &out, "addressToShortString", _addr)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// AddressToShortString is a free data retrieval call binding the contract method 0x17101940.
//
// Solidity: function addressToShortString(address _addr) pure returns(string)
func (_VendingMachine *VendingMachineSession) AddressToShortString(_addr common.Address) (string, error) {
	return _VendingMachine.Contract.AddressToShortString(&_VendingMachine.CallOpts, _addr)
}

// AddressToShortString is a free data retrieval call binding the contract method 0x17101940.
//
// Solidity: function addressToShortString(address _addr) pure returns(string)
func (_VendingMachine *VendingMachineCallerSession) AddressToShortString(_addr common.Address) (string, error) {
	return _VendingMachine.Contract.AddressToShortString(&_VendingMachine.CallOpts, _addr)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_VendingMachine *VendingMachineCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VendingMachine.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_VendingMachine *VendingMachineSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _VendingMachine.Contract.Allowance(&_VendingMachine.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_VendingMachine *VendingMachineCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _VendingMachine.Contract.Allowance(&_VendingMachine.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_VendingMachine *VendingMachineCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VendingMachine.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_VendingMachine *VendingMachineSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _VendingMachine.Contract.BalanceOf(&_VendingMachine.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_VendingMachine *VendingMachineCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _VendingMachine.Contract.BalanceOf(&_VendingMachine.CallOpts, owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_VendingMachine *VendingMachineCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _VendingMachine.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_VendingMachine *VendingMachineSession) Decimals() (uint8, error) {
	return _VendingMachine.Contract.Decimals(&_VendingMachine.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_VendingMachine *VendingMachineCallerSession) Decimals() (uint8, error) {
	return _VendingMachine.Contract.Decimals(&_VendingMachine.CallOpts)
}

// GetBaseInput is a free data retrieval call binding the contract method 0x65f6fcbb.
//
// Solidity: function getBaseInput() view returns(uint256 input)
func (_VendingMachine *VendingMachineCaller) GetBaseInput(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VendingMachine.contract.Call(opts, &out, "getBaseInput")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBaseInput is a free data retrieval call binding the contract method 0x65f6fcbb.
//
// Solidity: function getBaseInput() view returns(uint256 input)
func (_VendingMachine *VendingMachineSession) GetBaseInput() (*big.Int, error) {
	return _VendingMachine.Contract.GetBaseInput(&_VendingMachine.CallOpts)
}

// GetBaseInput is a free data retrieval call binding the contract method 0x65f6fcbb.
//
// Solidity: function getBaseInput() view returns(uint256 input)
func (_VendingMachine *VendingMachineCallerSession) GetBaseInput() (*big.Int, error) {
	return _VendingMachine.Contract.GetBaseInput(&_VendingMachine.CallOpts)
}

// GetMidPrice is a free data retrieval call binding the contract method 0xee27c689.
//
// Solidity: function getMidPrice() view returns(uint256 midPrice)
func (_VendingMachine *VendingMachineCaller) GetMidPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VendingMachine.contract.Call(opts, &out, "getMidPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMidPrice is a free data retrieval call binding the contract method 0xee27c689.
//
// Solidity: function getMidPrice() view returns(uint256 midPrice)
func (_VendingMachine *VendingMachineSession) GetMidPrice() (*big.Int, error) {
	return _VendingMachine.Contract.GetMidPrice(&_VendingMachine.CallOpts)
}

// GetMidPrice is a free data retrieval call binding the contract method 0xee27c689.
//
// Solidity: function getMidPrice() view returns(uint256 midPrice)
func (_VendingMachine *VendingMachineCallerSession) GetMidPrice() (*big.Int, error) {
	return _VendingMachine.Contract.GetMidPrice(&_VendingMachine.CallOpts)
}

// GetPMMState is a free data retrieval call binding the contract method 0xa382d1b9.
//
// Solidity: function getPMMState() view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint8) state)
func (_VendingMachine *VendingMachineCaller) GetPMMState(opts *bind.CallOpts) (PMMPricingPMMState, error) {
	var out []interface{}
	err := _VendingMachine.contract.Call(opts, &out, "getPMMState")

	if err != nil {
		return *new(PMMPricingPMMState), err
	}

	out0 := *abi.ConvertType(out[0], new(PMMPricingPMMState)).(*PMMPricingPMMState)

	return out0, err

}

// GetPMMState is a free data retrieval call binding the contract method 0xa382d1b9.
//
// Solidity: function getPMMState() view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint8) state)
func (_VendingMachine *VendingMachineSession) GetPMMState() (PMMPricingPMMState, error) {
	return _VendingMachine.Contract.GetPMMState(&_VendingMachine.CallOpts)
}

// GetPMMState is a free data retrieval call binding the contract method 0xa382d1b9.
//
// Solidity: function getPMMState() view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint8) state)
func (_VendingMachine *VendingMachineCallerSession) GetPMMState() (PMMPricingPMMState, error) {
	return _VendingMachine.Contract.GetPMMState(&_VendingMachine.CallOpts)
}

// GetPMMStateForCall is a free data retrieval call binding the contract method 0xfd1ed7e9.
//
// Solidity: function getPMMStateForCall() view returns(uint256 i, uint256 K, uint256 B, uint256 Q, uint256 B0, uint256 Q0, uint256 R)
func (_VendingMachine *VendingMachineCaller) GetPMMStateForCall(opts *bind.CallOpts) (struct {
	I  *big.Int
	K  *big.Int
	B  *big.Int
	Q  *big.Int
	B0 *big.Int
	Q0 *big.Int
	R  *big.Int
}, error) {
	var out []interface{}
	err := _VendingMachine.contract.Call(opts, &out, "getPMMStateForCall")

	outstruct := new(struct {
		I  *big.Int
		K  *big.Int
		B  *big.Int
		Q  *big.Int
		B0 *big.Int
		Q0 *big.Int
		R  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.I = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.K = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.B = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Q = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.B0 = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Q0 = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.R = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetPMMStateForCall is a free data retrieval call binding the contract method 0xfd1ed7e9.
//
// Solidity: function getPMMStateForCall() view returns(uint256 i, uint256 K, uint256 B, uint256 Q, uint256 B0, uint256 Q0, uint256 R)
func (_VendingMachine *VendingMachineSession) GetPMMStateForCall() (struct {
	I  *big.Int
	K  *big.Int
	B  *big.Int
	Q  *big.Int
	B0 *big.Int
	Q0 *big.Int
	R  *big.Int
}, error) {
	return _VendingMachine.Contract.GetPMMStateForCall(&_VendingMachine.CallOpts)
}

// GetPMMStateForCall is a free data retrieval call binding the contract method 0xfd1ed7e9.
//
// Solidity: function getPMMStateForCall() view returns(uint256 i, uint256 K, uint256 B, uint256 Q, uint256 B0, uint256 Q0, uint256 R)
func (_VendingMachine *VendingMachineCallerSession) GetPMMStateForCall() (struct {
	I  *big.Int
	K  *big.Int
	B  *big.Int
	Q  *big.Int
	B0 *big.Int
	Q0 *big.Int
	R  *big.Int
}, error) {
	return _VendingMachine.Contract.GetPMMStateForCall(&_VendingMachine.CallOpts)
}

// GetQuoteInput is a free data retrieval call binding the contract method 0x71f9100c.
//
// Solidity: function getQuoteInput() view returns(uint256 input)
func (_VendingMachine *VendingMachineCaller) GetQuoteInput(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VendingMachine.contract.Call(opts, &out, "getQuoteInput")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetQuoteInput is a free data retrieval call binding the contract method 0x71f9100c.
//
// Solidity: function getQuoteInput() view returns(uint256 input)
func (_VendingMachine *VendingMachineSession) GetQuoteInput() (*big.Int, error) {
	return _VendingMachine.Contract.GetQuoteInput(&_VendingMachine.CallOpts)
}

// GetQuoteInput is a free data retrieval call binding the contract method 0x71f9100c.
//
// Solidity: function getQuoteInput() view returns(uint256 input)
func (_VendingMachine *VendingMachineCallerSession) GetQuoteInput() (*big.Int, error) {
	return _VendingMachine.Contract.GetQuoteInput(&_VendingMachine.CallOpts)
}

// GetUserFeeRate is a free data retrieval call binding the contract method 0x44096609.
//
// Solidity: function getUserFeeRate(address user) view returns(uint256 lpFeeRate, uint256 mtFeeRate)
func (_VendingMachine *VendingMachineCaller) GetUserFeeRate(opts *bind.CallOpts, user common.Address) (struct {
	LpFeeRate *big.Int
	MtFeeRate *big.Int
}, error) {
	var out []interface{}
	err := _VendingMachine.contract.Call(opts, &out, "getUserFeeRate", user)

	outstruct := new(struct {
		LpFeeRate *big.Int
		MtFeeRate *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LpFeeRate = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.MtFeeRate = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetUserFeeRate is a free data retrieval call binding the contract method 0x44096609.
//
// Solidity: function getUserFeeRate(address user) view returns(uint256 lpFeeRate, uint256 mtFeeRate)
func (_VendingMachine *VendingMachineSession) GetUserFeeRate(user common.Address) (struct {
	LpFeeRate *big.Int
	MtFeeRate *big.Int
}, error) {
	return _VendingMachine.Contract.GetUserFeeRate(&_VendingMachine.CallOpts, user)
}

// GetUserFeeRate is a free data retrieval call binding the contract method 0x44096609.
//
// Solidity: function getUserFeeRate(address user) view returns(uint256 lpFeeRate, uint256 mtFeeRate)
func (_VendingMachine *VendingMachineCallerSession) GetUserFeeRate(user common.Address) (struct {
	LpFeeRate *big.Int
	MtFeeRate *big.Int
}, error) {
	return _VendingMachine.Contract.GetUserFeeRate(&_VendingMachine.CallOpts, user)
}

// GetVaultReserve is a free data retrieval call binding the contract method 0x36223ce9.
//
// Solidity: function getVaultReserve() view returns(uint256 baseReserve, uint256 quoteReserve)
func (_VendingMachine *VendingMachineCaller) GetVaultReserve(opts *bind.CallOpts) (struct {
	BaseReserve  *big.Int
	QuoteReserve *big.Int
}, error) {
	var out []interface{}
	err := _VendingMachine.contract.Call(opts, &out, "getVaultReserve")

	outstruct := new(struct {
		BaseReserve  *big.Int
		QuoteReserve *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BaseReserve = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.QuoteReserve = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetVaultReserve is a free data retrieval call binding the contract method 0x36223ce9.
//
// Solidity: function getVaultReserve() view returns(uint256 baseReserve, uint256 quoteReserve)
func (_VendingMachine *VendingMachineSession) GetVaultReserve() (struct {
	BaseReserve  *big.Int
	QuoteReserve *big.Int
}, error) {
	return _VendingMachine.Contract.GetVaultReserve(&_VendingMachine.CallOpts)
}

// GetVaultReserve is a free data retrieval call binding the contract method 0x36223ce9.
//
// Solidity: function getVaultReserve() view returns(uint256 baseReserve, uint256 quoteReserve)
func (_VendingMachine *VendingMachineCallerSession) GetVaultReserve() (struct {
	BaseReserve  *big.Int
	QuoteReserve *big.Int
}, error) {
	return _VendingMachine.Contract.GetVaultReserve(&_VendingMachine.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_VendingMachine *VendingMachineCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _VendingMachine.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_VendingMachine *VendingMachineSession) Name() (string, error) {
	return _VendingMachine.Contract.Name(&_VendingMachine.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_VendingMachine *VendingMachineCallerSession) Name() (string, error) {
	return _VendingMachine.Contract.Name(&_VendingMachine.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_VendingMachine *VendingMachineCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VendingMachine.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_VendingMachine *VendingMachineSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _VendingMachine.Contract.Nonces(&_VendingMachine.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_VendingMachine *VendingMachineCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _VendingMachine.Contract.Nonces(&_VendingMachine.CallOpts, arg0)
}

// QuerySellBase is a free data retrieval call binding the contract method 0x79a04876.
//
// Solidity: function querySellBase(address trader, uint256 payBaseAmount) view returns(uint256 receiveQuoteAmount, uint256 mtFee)
func (_VendingMachine *VendingMachineCaller) QuerySellBase(opts *bind.CallOpts, trader common.Address, payBaseAmount *big.Int) (struct {
	ReceiveQuoteAmount *big.Int
	MtFee              *big.Int
}, error) {
	var out []interface{}
	err := _VendingMachine.contract.Call(opts, &out, "querySellBase", trader, payBaseAmount)

	outstruct := new(struct {
		ReceiveQuoteAmount *big.Int
		MtFee              *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ReceiveQuoteAmount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.MtFee = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// QuerySellBase is a free data retrieval call binding the contract method 0x79a04876.
//
// Solidity: function querySellBase(address trader, uint256 payBaseAmount) view returns(uint256 receiveQuoteAmount, uint256 mtFee)
func (_VendingMachine *VendingMachineSession) QuerySellBase(trader common.Address, payBaseAmount *big.Int) (struct {
	ReceiveQuoteAmount *big.Int
	MtFee              *big.Int
}, error) {
	return _VendingMachine.Contract.QuerySellBase(&_VendingMachine.CallOpts, trader, payBaseAmount)
}

// QuerySellBase is a free data retrieval call binding the contract method 0x79a04876.
//
// Solidity: function querySellBase(address trader, uint256 payBaseAmount) view returns(uint256 receiveQuoteAmount, uint256 mtFee)
func (_VendingMachine *VendingMachineCallerSession) QuerySellBase(trader common.Address, payBaseAmount *big.Int) (struct {
	ReceiveQuoteAmount *big.Int
	MtFee              *big.Int
}, error) {
	return _VendingMachine.Contract.QuerySellBase(&_VendingMachine.CallOpts, trader, payBaseAmount)
}

// QuerySellQuote is a free data retrieval call binding the contract method 0x66410a21.
//
// Solidity: function querySellQuote(address trader, uint256 payQuoteAmount) view returns(uint256 receiveBaseAmount, uint256 mtFee)
func (_VendingMachine *VendingMachineCaller) QuerySellQuote(opts *bind.CallOpts, trader common.Address, payQuoteAmount *big.Int) (struct {
	ReceiveBaseAmount *big.Int
	MtFee             *big.Int
}, error) {
	var out []interface{}
	err := _VendingMachine.contract.Call(opts, &out, "querySellQuote", trader, payQuoteAmount)

	outstruct := new(struct {
		ReceiveBaseAmount *big.Int
		MtFee             *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ReceiveBaseAmount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.MtFee = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// QuerySellQuote is a free data retrieval call binding the contract method 0x66410a21.
//
// Solidity: function querySellQuote(address trader, uint256 payQuoteAmount) view returns(uint256 receiveBaseAmount, uint256 mtFee)
func (_VendingMachine *VendingMachineSession) QuerySellQuote(trader common.Address, payQuoteAmount *big.Int) (struct {
	ReceiveBaseAmount *big.Int
	MtFee             *big.Int
}, error) {
	return _VendingMachine.Contract.QuerySellQuote(&_VendingMachine.CallOpts, trader, payQuoteAmount)
}

// QuerySellQuote is a free data retrieval call binding the contract method 0x66410a21.
//
// Solidity: function querySellQuote(address trader, uint256 payQuoteAmount) view returns(uint256 receiveBaseAmount, uint256 mtFee)
func (_VendingMachine *VendingMachineCallerSession) QuerySellQuote(trader common.Address, payQuoteAmount *big.Int) (struct {
	ReceiveBaseAmount *big.Int
	MtFee             *big.Int
}, error) {
	return _VendingMachine.Contract.QuerySellQuote(&_VendingMachine.CallOpts, trader, payQuoteAmount)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_VendingMachine *VendingMachineCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _VendingMachine.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_VendingMachine *VendingMachineSession) Symbol() (string, error) {
	return _VendingMachine.Contract.Symbol(&_VendingMachine.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_VendingMachine *VendingMachineCallerSession) Symbol() (string, error) {
	return _VendingMachine.Contract.Symbol(&_VendingMachine.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_VendingMachine *VendingMachineCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VendingMachine.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_VendingMachine *VendingMachineSession) TotalSupply() (*big.Int, error) {
	return _VendingMachine.Contract.TotalSupply(&_VendingMachine.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_VendingMachine *VendingMachineCallerSession) TotalSupply() (*big.Int, error) {
	return _VendingMachine.Contract.TotalSupply(&_VendingMachine.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_VendingMachine *VendingMachineCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _VendingMachine.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_VendingMachine *VendingMachineSession) Version() (string, error) {
	return _VendingMachine.Contract.Version(&_VendingMachine.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_VendingMachine *VendingMachineCallerSession) Version() (string, error) {
	return _VendingMachine.Contract.Version(&_VendingMachine.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_VendingMachine *VendingMachineTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VendingMachine.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_VendingMachine *VendingMachineSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VendingMachine.Contract.Approve(&_VendingMachine.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_VendingMachine *VendingMachineTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VendingMachine.Contract.Approve(&_VendingMachine.TransactOpts, spender, amount)
}

// BuyShares is a paid mutator transaction binding the contract method 0x4c85b425.
//
// Solidity: function buyShares(address to) returns(uint256 shares, uint256 baseInput, uint256 quoteInput)
func (_VendingMachine *VendingMachineTransactor) BuyShares(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _VendingMachine.contract.Transact(opts, "buyShares", to)
}

// BuyShares is a paid mutator transaction binding the contract method 0x4c85b425.
//
// Solidity: function buyShares(address to) returns(uint256 shares, uint256 baseInput, uint256 quoteInput)
func (_VendingMachine *VendingMachineSession) BuyShares(to common.Address) (*types.Transaction, error) {
	return _VendingMachine.Contract.BuyShares(&_VendingMachine.TransactOpts, to)
}

// BuyShares is a paid mutator transaction binding the contract method 0x4c85b425.
//
// Solidity: function buyShares(address to) returns(uint256 shares, uint256 baseInput, uint256 quoteInput)
func (_VendingMachine *VendingMachineTransactorSession) BuyShares(to common.Address) (*types.Transaction, error) {
	return _VendingMachine.Contract.BuyShares(&_VendingMachine.TransactOpts, to)
}

// FlashLoan is a paid mutator transaction binding the contract method 0xd0a494e4.
//
// Solidity: function flashLoan(uint256 baseAmount, uint256 quoteAmount, address assetTo, bytes data) returns()
func (_VendingMachine *VendingMachineTransactor) FlashLoan(opts *bind.TransactOpts, baseAmount *big.Int, quoteAmount *big.Int, assetTo common.Address, data []byte) (*types.Transaction, error) {
	return _VendingMachine.contract.Transact(opts, "flashLoan", baseAmount, quoteAmount, assetTo, data)
}

// FlashLoan is a paid mutator transaction binding the contract method 0xd0a494e4.
//
// Solidity: function flashLoan(uint256 baseAmount, uint256 quoteAmount, address assetTo, bytes data) returns()
func (_VendingMachine *VendingMachineSession) FlashLoan(baseAmount *big.Int, quoteAmount *big.Int, assetTo common.Address, data []byte) (*types.Transaction, error) {
	return _VendingMachine.Contract.FlashLoan(&_VendingMachine.TransactOpts, baseAmount, quoteAmount, assetTo, data)
}

// FlashLoan is a paid mutator transaction binding the contract method 0xd0a494e4.
//
// Solidity: function flashLoan(uint256 baseAmount, uint256 quoteAmount, address assetTo, bytes data) returns()
func (_VendingMachine *VendingMachineTransactorSession) FlashLoan(baseAmount *big.Int, quoteAmount *big.Int, assetTo common.Address, data []byte) (*types.Transaction, error) {
	return _VendingMachine.Contract.FlashLoan(&_VendingMachine.TransactOpts, baseAmount, quoteAmount, assetTo, data)
}

// Init is a paid mutator transaction binding the contract method 0x5039972a.
//
// Solidity: function init(address maintainer, address baseTokenAddress, address quoteTokenAddress, uint256 lpFeeRate, address mtFeeRateModel, uint256 i, uint256 k, bool isOpenTWAP) returns()
func (_VendingMachine *VendingMachineTransactor) Init(opts *bind.TransactOpts, maintainer common.Address, baseTokenAddress common.Address, quoteTokenAddress common.Address, lpFeeRate *big.Int, mtFeeRateModel common.Address, i *big.Int, k *big.Int, isOpenTWAP bool) (*types.Transaction, error) {
	return _VendingMachine.contract.Transact(opts, "init", maintainer, baseTokenAddress, quoteTokenAddress, lpFeeRate, mtFeeRateModel, i, k, isOpenTWAP)
}

// Init is a paid mutator transaction binding the contract method 0x5039972a.
//
// Solidity: function init(address maintainer, address baseTokenAddress, address quoteTokenAddress, uint256 lpFeeRate, address mtFeeRateModel, uint256 i, uint256 k, bool isOpenTWAP) returns()
func (_VendingMachine *VendingMachineSession) Init(maintainer common.Address, baseTokenAddress common.Address, quoteTokenAddress common.Address, lpFeeRate *big.Int, mtFeeRateModel common.Address, i *big.Int, k *big.Int, isOpenTWAP bool) (*types.Transaction, error) {
	return _VendingMachine.Contract.Init(&_VendingMachine.TransactOpts, maintainer, baseTokenAddress, quoteTokenAddress, lpFeeRate, mtFeeRateModel, i, k, isOpenTWAP)
}

// Init is a paid mutator transaction binding the contract method 0x5039972a.
//
// Solidity: function init(address maintainer, address baseTokenAddress, address quoteTokenAddress, uint256 lpFeeRate, address mtFeeRateModel, uint256 i, uint256 k, bool isOpenTWAP) returns()
func (_VendingMachine *VendingMachineTransactorSession) Init(maintainer common.Address, baseTokenAddress common.Address, quoteTokenAddress common.Address, lpFeeRate *big.Int, mtFeeRateModel common.Address, i *big.Int, k *big.Int, isOpenTWAP bool) (*types.Transaction, error) {
	return _VendingMachine.Contract.Init(&_VendingMachine.TransactOpts, maintainer, baseTokenAddress, quoteTokenAddress, lpFeeRate, mtFeeRateModel, i, k, isOpenTWAP)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_VendingMachine *VendingMachineTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _VendingMachine.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_VendingMachine *VendingMachineSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _VendingMachine.Contract.Permit(&_VendingMachine.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_VendingMachine *VendingMachineTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _VendingMachine.Contract.Permit(&_VendingMachine.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// SellBase is a paid mutator transaction binding the contract method 0xbd6015b4.
//
// Solidity: function sellBase(address to) returns(uint256 receiveQuoteAmount)
func (_VendingMachine *VendingMachineTransactor) SellBase(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _VendingMachine.contract.Transact(opts, "sellBase", to)
}

// SellBase is a paid mutator transaction binding the contract method 0xbd6015b4.
//
// Solidity: function sellBase(address to) returns(uint256 receiveQuoteAmount)
func (_VendingMachine *VendingMachineSession) SellBase(to common.Address) (*types.Transaction, error) {
	return _VendingMachine.Contract.SellBase(&_VendingMachine.TransactOpts, to)
}

// SellBase is a paid mutator transaction binding the contract method 0xbd6015b4.
//
// Solidity: function sellBase(address to) returns(uint256 receiveQuoteAmount)
func (_VendingMachine *VendingMachineTransactorSession) SellBase(to common.Address) (*types.Transaction, error) {
	return _VendingMachine.Contract.SellBase(&_VendingMachine.TransactOpts, to)
}

// SellQuote is a paid mutator transaction binding the contract method 0xdd93f59a.
//
// Solidity: function sellQuote(address to) returns(uint256 receiveBaseAmount)
func (_VendingMachine *VendingMachineTransactor) SellQuote(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _VendingMachine.contract.Transact(opts, "sellQuote", to)
}

// SellQuote is a paid mutator transaction binding the contract method 0xdd93f59a.
//
// Solidity: function sellQuote(address to) returns(uint256 receiveBaseAmount)
func (_VendingMachine *VendingMachineSession) SellQuote(to common.Address) (*types.Transaction, error) {
	return _VendingMachine.Contract.SellQuote(&_VendingMachine.TransactOpts, to)
}

// SellQuote is a paid mutator transaction binding the contract method 0xdd93f59a.
//
// Solidity: function sellQuote(address to) returns(uint256 receiveBaseAmount)
func (_VendingMachine *VendingMachineTransactorSession) SellQuote(to common.Address) (*types.Transaction, error) {
	return _VendingMachine.Contract.SellQuote(&_VendingMachine.TransactOpts, to)
}

// SellShares is a paid mutator transaction binding the contract method 0xb56ceaa6.
//
// Solidity: function sellShares(uint256 shareAmount, address to, uint256 baseMinAmount, uint256 quoteMinAmount, bytes data, uint256 deadline) returns(uint256 baseAmount, uint256 quoteAmount)
func (_VendingMachine *VendingMachineTransactor) SellShares(opts *bind.TransactOpts, shareAmount *big.Int, to common.Address, baseMinAmount *big.Int, quoteMinAmount *big.Int, data []byte, deadline *big.Int) (*types.Transaction, error) {
	return _VendingMachine.contract.Transact(opts, "sellShares", shareAmount, to, baseMinAmount, quoteMinAmount, data, deadline)
}

// SellShares is a paid mutator transaction binding the contract method 0xb56ceaa6.
//
// Solidity: function sellShares(uint256 shareAmount, address to, uint256 baseMinAmount, uint256 quoteMinAmount, bytes data, uint256 deadline) returns(uint256 baseAmount, uint256 quoteAmount)
func (_VendingMachine *VendingMachineSession) SellShares(shareAmount *big.Int, to common.Address, baseMinAmount *big.Int, quoteMinAmount *big.Int, data []byte, deadline *big.Int) (*types.Transaction, error) {
	return _VendingMachine.Contract.SellShares(&_VendingMachine.TransactOpts, shareAmount, to, baseMinAmount, quoteMinAmount, data, deadline)
}

// SellShares is a paid mutator transaction binding the contract method 0xb56ceaa6.
//
// Solidity: function sellShares(uint256 shareAmount, address to, uint256 baseMinAmount, uint256 quoteMinAmount, bytes data, uint256 deadline) returns(uint256 baseAmount, uint256 quoteAmount)
func (_VendingMachine *VendingMachineTransactorSession) SellShares(shareAmount *big.Int, to common.Address, baseMinAmount *big.Int, quoteMinAmount *big.Int, data []byte, deadline *big.Int) (*types.Transaction, error) {
	return _VendingMachine.Contract.SellShares(&_VendingMachine.TransactOpts, shareAmount, to, baseMinAmount, quoteMinAmount, data, deadline)
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_VendingMachine *VendingMachineTransactor) Sync(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VendingMachine.contract.Transact(opts, "sync")
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_VendingMachine *VendingMachineSession) Sync() (*types.Transaction, error) {
	return _VendingMachine.Contract.Sync(&_VendingMachine.TransactOpts)
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_VendingMachine *VendingMachineTransactorSession) Sync() (*types.Transaction, error) {
	return _VendingMachine.Contract.Sync(&_VendingMachine.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_VendingMachine *VendingMachineTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VendingMachine.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_VendingMachine *VendingMachineSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VendingMachine.Contract.Transfer(&_VendingMachine.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_VendingMachine *VendingMachineTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VendingMachine.Contract.Transfer(&_VendingMachine.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_VendingMachine *VendingMachineTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VendingMachine.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_VendingMachine *VendingMachineSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VendingMachine.Contract.TransferFrom(&_VendingMachine.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_VendingMachine *VendingMachineTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VendingMachine.Contract.TransferFrom(&_VendingMachine.TransactOpts, from, to, amount)
}

// VendingMachineApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the VendingMachine contract.
type VendingMachineApprovalIterator struct {
	Event *VendingMachineApproval // Event containing the contract specifics and raw log

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
func (it *VendingMachineApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VendingMachineApproval)
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
		it.Event = new(VendingMachineApproval)
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
func (it *VendingMachineApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VendingMachineApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VendingMachineApproval represents a Approval event raised by the VendingMachine contract.
type VendingMachineApproval struct {
	Owner   common.Address
	Spender common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_VendingMachine *VendingMachineFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*VendingMachineApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _VendingMachine.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &VendingMachineApprovalIterator{contract: _VendingMachine.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_VendingMachine *VendingMachineFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *VendingMachineApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _VendingMachine.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VendingMachineApproval)
				if err := _VendingMachine.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_VendingMachine *VendingMachineFilterer) ParseApproval(log types.Log) (*VendingMachineApproval, error) {
	event := new(VendingMachineApproval)
	if err := _VendingMachine.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VendingMachineBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the VendingMachine contract.
type VendingMachineBurnIterator struct {
	Event *VendingMachineBurn // Event containing the contract specifics and raw log

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
func (it *VendingMachineBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VendingMachineBurn)
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
		it.Event = new(VendingMachineBurn)
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
func (it *VendingMachineBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VendingMachineBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VendingMachineBurn represents a Burn event raised by the VendingMachine contract.
type VendingMachineBurn struct {
	User  common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed user, uint256 value)
func (_VendingMachine *VendingMachineFilterer) FilterBurn(opts *bind.FilterOpts, user []common.Address) (*VendingMachineBurnIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _VendingMachine.contract.FilterLogs(opts, "Burn", userRule)
	if err != nil {
		return nil, err
	}
	return &VendingMachineBurnIterator{contract: _VendingMachine.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed user, uint256 value)
func (_VendingMachine *VendingMachineFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *VendingMachineBurn, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _VendingMachine.contract.WatchLogs(opts, "Burn", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VendingMachineBurn)
				if err := _VendingMachine.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed user, uint256 value)
func (_VendingMachine *VendingMachineFilterer) ParseBurn(log types.Log) (*VendingMachineBurn, error) {
	event := new(VendingMachineBurn)
	if err := _VendingMachine.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VendingMachineBuySharesIterator is returned from FilterBuyShares and is used to iterate over the raw logs and unpacked data for BuyShares events raised by the VendingMachine contract.
type VendingMachineBuySharesIterator struct {
	Event *VendingMachineBuyShares // Event containing the contract specifics and raw log

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
func (it *VendingMachineBuySharesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VendingMachineBuyShares)
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
		it.Event = new(VendingMachineBuyShares)
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
func (it *VendingMachineBuySharesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VendingMachineBuySharesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VendingMachineBuyShares represents a BuyShares event raised by the VendingMachine contract.
type VendingMachineBuyShares struct {
	To             common.Address
	IncreaseShares *big.Int
	TotalShares    *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBuyShares is a free log retrieval operation binding the contract event 0x1c172440bdebb59cd92a7f08f4227903a3305ab6f880cb25f93eddb66843a102.
//
// Solidity: event BuyShares(address to, uint256 increaseShares, uint256 totalShares)
func (_VendingMachine *VendingMachineFilterer) FilterBuyShares(opts *bind.FilterOpts) (*VendingMachineBuySharesIterator, error) {

	logs, sub, err := _VendingMachine.contract.FilterLogs(opts, "BuyShares")
	if err != nil {
		return nil, err
	}
	return &VendingMachineBuySharesIterator{contract: _VendingMachine.contract, event: "BuyShares", logs: logs, sub: sub}, nil
}

// WatchBuyShares is a free log subscription operation binding the contract event 0x1c172440bdebb59cd92a7f08f4227903a3305ab6f880cb25f93eddb66843a102.
//
// Solidity: event BuyShares(address to, uint256 increaseShares, uint256 totalShares)
func (_VendingMachine *VendingMachineFilterer) WatchBuyShares(opts *bind.WatchOpts, sink chan<- *VendingMachineBuyShares) (event.Subscription, error) {

	logs, sub, err := _VendingMachine.contract.WatchLogs(opts, "BuyShares")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VendingMachineBuyShares)
				if err := _VendingMachine.contract.UnpackLog(event, "BuyShares", log); err != nil {
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

// ParseBuyShares is a log parse operation binding the contract event 0x1c172440bdebb59cd92a7f08f4227903a3305ab6f880cb25f93eddb66843a102.
//
// Solidity: event BuyShares(address to, uint256 increaseShares, uint256 totalShares)
func (_VendingMachine *VendingMachineFilterer) ParseBuyShares(log types.Log) (*VendingMachineBuyShares, error) {
	event := new(VendingMachineBuyShares)
	if err := _VendingMachine.contract.UnpackLog(event, "BuyShares", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VendingMachineDODOFlashLoanIterator is returned from FilterDODOFlashLoan and is used to iterate over the raw logs and unpacked data for DODOFlashLoan events raised by the VendingMachine contract.
type VendingMachineDODOFlashLoanIterator struct {
	Event *VendingMachineDODOFlashLoan // Event containing the contract specifics and raw log

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
func (it *VendingMachineDODOFlashLoanIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VendingMachineDODOFlashLoan)
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
		it.Event = new(VendingMachineDODOFlashLoan)
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
func (it *VendingMachineDODOFlashLoanIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VendingMachineDODOFlashLoanIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VendingMachineDODOFlashLoan represents a DODOFlashLoan event raised by the VendingMachine contract.
type VendingMachineDODOFlashLoan struct {
	Borrower    common.Address
	AssetTo     common.Address
	BaseAmount  *big.Int
	QuoteAmount *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDODOFlashLoan is a free log retrieval operation binding the contract event 0x0b82e93068db15abd9fbb2682c65462ea8a0a10582dce93a5664818e296f54eb.
//
// Solidity: event DODOFlashLoan(address borrower, address assetTo, uint256 baseAmount, uint256 quoteAmount)
func (_VendingMachine *VendingMachineFilterer) FilterDODOFlashLoan(opts *bind.FilterOpts) (*VendingMachineDODOFlashLoanIterator, error) {

	logs, sub, err := _VendingMachine.contract.FilterLogs(opts, "DODOFlashLoan")
	if err != nil {
		return nil, err
	}
	return &VendingMachineDODOFlashLoanIterator{contract: _VendingMachine.contract, event: "DODOFlashLoan", logs: logs, sub: sub}, nil
}

// WatchDODOFlashLoan is a free log subscription operation binding the contract event 0x0b82e93068db15abd9fbb2682c65462ea8a0a10582dce93a5664818e296f54eb.
//
// Solidity: event DODOFlashLoan(address borrower, address assetTo, uint256 baseAmount, uint256 quoteAmount)
func (_VendingMachine *VendingMachineFilterer) WatchDODOFlashLoan(opts *bind.WatchOpts, sink chan<- *VendingMachineDODOFlashLoan) (event.Subscription, error) {

	logs, sub, err := _VendingMachine.contract.WatchLogs(opts, "DODOFlashLoan")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VendingMachineDODOFlashLoan)
				if err := _VendingMachine.contract.UnpackLog(event, "DODOFlashLoan", log); err != nil {
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

// ParseDODOFlashLoan is a log parse operation binding the contract event 0x0b82e93068db15abd9fbb2682c65462ea8a0a10582dce93a5664818e296f54eb.
//
// Solidity: event DODOFlashLoan(address borrower, address assetTo, uint256 baseAmount, uint256 quoteAmount)
func (_VendingMachine *VendingMachineFilterer) ParseDODOFlashLoan(log types.Log) (*VendingMachineDODOFlashLoan, error) {
	event := new(VendingMachineDODOFlashLoan)
	if err := _VendingMachine.contract.UnpackLog(event, "DODOFlashLoan", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VendingMachineDODOSwapIterator is returned from FilterDODOSwap and is used to iterate over the raw logs and unpacked data for DODOSwap events raised by the VendingMachine contract.
type VendingMachineDODOSwapIterator struct {
	Event *VendingMachineDODOSwap // Event containing the contract specifics and raw log

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
func (it *VendingMachineDODOSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VendingMachineDODOSwap)
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
		it.Event = new(VendingMachineDODOSwap)
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
func (it *VendingMachineDODOSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VendingMachineDODOSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VendingMachineDODOSwap represents a DODOSwap event raised by the VendingMachine contract.
type VendingMachineDODOSwap struct {
	FromToken  common.Address
	ToToken    common.Address
	FromAmount *big.Int
	ToAmount   *big.Int
	Trader     common.Address
	Receiver   common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDODOSwap is a free log retrieval operation binding the contract event 0xc2c0245e056d5fb095f04cd6373bc770802ebd1e6c918eb78fdef843cdb37b0f.
//
// Solidity: event DODOSwap(address fromToken, address toToken, uint256 fromAmount, uint256 toAmount, address trader, address receiver)
func (_VendingMachine *VendingMachineFilterer) FilterDODOSwap(opts *bind.FilterOpts) (*VendingMachineDODOSwapIterator, error) {

	logs, sub, err := _VendingMachine.contract.FilterLogs(opts, "DODOSwap")
	if err != nil {
		return nil, err
	}
	return &VendingMachineDODOSwapIterator{contract: _VendingMachine.contract, event: "DODOSwap", logs: logs, sub: sub}, nil
}

// WatchDODOSwap is a free log subscription operation binding the contract event 0xc2c0245e056d5fb095f04cd6373bc770802ebd1e6c918eb78fdef843cdb37b0f.
//
// Solidity: event DODOSwap(address fromToken, address toToken, uint256 fromAmount, uint256 toAmount, address trader, address receiver)
func (_VendingMachine *VendingMachineFilterer) WatchDODOSwap(opts *bind.WatchOpts, sink chan<- *VendingMachineDODOSwap) (event.Subscription, error) {

	logs, sub, err := _VendingMachine.contract.WatchLogs(opts, "DODOSwap")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VendingMachineDODOSwap)
				if err := _VendingMachine.contract.UnpackLog(event, "DODOSwap", log); err != nil {
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

// ParseDODOSwap is a log parse operation binding the contract event 0xc2c0245e056d5fb095f04cd6373bc770802ebd1e6c918eb78fdef843cdb37b0f.
//
// Solidity: event DODOSwap(address fromToken, address toToken, uint256 fromAmount, uint256 toAmount, address trader, address receiver)
func (_VendingMachine *VendingMachineFilterer) ParseDODOSwap(log types.Log) (*VendingMachineDODOSwap, error) {
	event := new(VendingMachineDODOSwap)
	if err := _VendingMachine.contract.UnpackLog(event, "DODOSwap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VendingMachineMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the VendingMachine contract.
type VendingMachineMintIterator struct {
	Event *VendingMachineMint // Event containing the contract specifics and raw log

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
func (it *VendingMachineMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VendingMachineMint)
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
		it.Event = new(VendingMachineMint)
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
func (it *VendingMachineMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VendingMachineMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VendingMachineMint represents a Mint event raised by the VendingMachine contract.
type VendingMachineMint struct {
	User  common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed user, uint256 value)
func (_VendingMachine *VendingMachineFilterer) FilterMint(opts *bind.FilterOpts, user []common.Address) (*VendingMachineMintIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _VendingMachine.contract.FilterLogs(opts, "Mint", userRule)
	if err != nil {
		return nil, err
	}
	return &VendingMachineMintIterator{contract: _VendingMachine.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed user, uint256 value)
func (_VendingMachine *VendingMachineFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *VendingMachineMint, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _VendingMachine.contract.WatchLogs(opts, "Mint", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VendingMachineMint)
				if err := _VendingMachine.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed user, uint256 value)
func (_VendingMachine *VendingMachineFilterer) ParseMint(log types.Log) (*VendingMachineMint, error) {
	event := new(VendingMachineMint)
	if err := _VendingMachine.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VendingMachineSellSharesIterator is returned from FilterSellShares and is used to iterate over the raw logs and unpacked data for SellShares events raised by the VendingMachine contract.
type VendingMachineSellSharesIterator struct {
	Event *VendingMachineSellShares // Event containing the contract specifics and raw log

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
func (it *VendingMachineSellSharesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VendingMachineSellShares)
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
		it.Event = new(VendingMachineSellShares)
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
func (it *VendingMachineSellSharesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VendingMachineSellSharesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VendingMachineSellShares represents a SellShares event raised by the VendingMachine contract.
type VendingMachineSellShares struct {
	Payer          common.Address
	To             common.Address
	DecreaseShares *big.Int
	TotalShares    *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSellShares is a free log retrieval operation binding the contract event 0x55caccde83781f39bfc1296eff45655b6496729443a7d48958b18b3b685600a5.
//
// Solidity: event SellShares(address payer, address to, uint256 decreaseShares, uint256 totalShares)
func (_VendingMachine *VendingMachineFilterer) FilterSellShares(opts *bind.FilterOpts) (*VendingMachineSellSharesIterator, error) {

	logs, sub, err := _VendingMachine.contract.FilterLogs(opts, "SellShares")
	if err != nil {
		return nil, err
	}
	return &VendingMachineSellSharesIterator{contract: _VendingMachine.contract, event: "SellShares", logs: logs, sub: sub}, nil
}

// WatchSellShares is a free log subscription operation binding the contract event 0x55caccde83781f39bfc1296eff45655b6496729443a7d48958b18b3b685600a5.
//
// Solidity: event SellShares(address payer, address to, uint256 decreaseShares, uint256 totalShares)
func (_VendingMachine *VendingMachineFilterer) WatchSellShares(opts *bind.WatchOpts, sink chan<- *VendingMachineSellShares) (event.Subscription, error) {

	logs, sub, err := _VendingMachine.contract.WatchLogs(opts, "SellShares")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VendingMachineSellShares)
				if err := _VendingMachine.contract.UnpackLog(event, "SellShares", log); err != nil {
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

// ParseSellShares is a log parse operation binding the contract event 0x55caccde83781f39bfc1296eff45655b6496729443a7d48958b18b3b685600a5.
//
// Solidity: event SellShares(address payer, address to, uint256 decreaseShares, uint256 totalShares)
func (_VendingMachine *VendingMachineFilterer) ParseSellShares(log types.Log) (*VendingMachineSellShares, error) {
	event := new(VendingMachineSellShares)
	if err := _VendingMachine.contract.UnpackLog(event, "SellShares", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VendingMachineTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the VendingMachine contract.
type VendingMachineTransferIterator struct {
	Event *VendingMachineTransfer // Event containing the contract specifics and raw log

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
func (it *VendingMachineTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VendingMachineTransfer)
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
		it.Event = new(VendingMachineTransfer)
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
func (it *VendingMachineTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VendingMachineTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VendingMachineTransfer represents a Transfer event raised by the VendingMachine contract.
type VendingMachineTransfer struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_VendingMachine *VendingMachineFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VendingMachineTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VendingMachine.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VendingMachineTransferIterator{contract: _VendingMachine.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_VendingMachine *VendingMachineFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *VendingMachineTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VendingMachine.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VendingMachineTransfer)
				if err := _VendingMachine.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_VendingMachine *VendingMachineFilterer) ParseTransfer(log types.Log) (*VendingMachineTransfer, error) {
	event := new(VendingMachineTransfer)
	if err := _VendingMachine.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
