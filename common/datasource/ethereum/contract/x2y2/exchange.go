// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package x2y2

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

// MarketFee is an auto generated low-level Go binding around an user-defined struct.
type MarketFee struct {
	Percentage *big.Int
	To         common.Address
}

// MarketOrder is an auto generated low-level Go binding around an user-defined struct.
type MarketOrder struct {
	Salt         *big.Int
	User         common.Address
	Network      *big.Int
	Intent       *big.Int
	DelegateType *big.Int
	Deadline     *big.Int
	Currency     common.Address
	DataMask     []byte
	Items        []MarketOrderItem
	R            [32]byte
	S            [32]byte
	V            uint8
	SignVersion  uint8
}

// MarketOrderItem is an auto generated low-level Go binding around an user-defined struct.
type MarketOrderItem struct {
	Price *big.Int
	Data  []byte
}

// MarketRunInput is an auto generated low-level Go binding around an user-defined struct.
type MarketRunInput struct {
	Orders  []MarketOrder
	Details []MarketSettleDetail
	Shared  MarketSettleShared
	R       [32]byte
	S       [32]byte
	V       uint8
}

// MarketSettleDetail is an auto generated low-level Go binding around an user-defined struct.
type MarketSettleDetail struct {
	Op                 uint8
	OrderIdx           *big.Int
	ItemIdx            *big.Int
	Price              *big.Int
	ItemHash           [32]byte
	ExecutionDelegate  common.Address
	DataReplacement    []byte
	BidIncentivePct    *big.Int
	AucMinIncrementPct *big.Int
	AucIncDurationSecs *big.Int
	Fees               []MarketFee
}

// MarketSettleShared is an auto generated low-level Go binding around an user-defined struct.
type MarketSettleShared struct {
	Salt         *big.Int
	Deadline     *big.Int
	AmountToEth  *big.Int
	AmountToWeth *big.Int
	User         common.Address
	CanFail      bool
}

// ExchangeMetaData contains all meta data concerning the Exchange contract.
var ExchangeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"itemHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"incentive\",\"type\":\"uint256\"}],\"name\":\"EvAuctionRefund\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"itemHash\",\"type\":\"bytes32\"}],\"name\":\"EvCancel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isRemoval\",\"type\":\"bool\"}],\"name\":\"EvDelegate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"error\",\"type\":\"bytes\"}],\"name\":\"EvFailure\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"EvFeeCapUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"itemHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"orderSalt\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"settleSalt\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"intent\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delegateType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"currency\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"dataMask\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structMarket.OrderItem\",\"name\":\"item\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumMarket.Op\",\"name\":\"op\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"orderIdx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"itemIdx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"itemHash\",\"type\":\"bytes32\"},{\"internalType\":\"contractIDelegate\",\"name\":\"executionDelegate\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"dataReplacement\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"bidIncentivePct\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"aucMinIncrementPct\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"aucIncDurationSecs\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"percentage\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"internalType\":\"structMarket.Fee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"}],\"indexed\":false,\"internalType\":\"structMarket.SettleDetail\",\"name\":\"detail\",\"type\":\"tuple\"}],\"name\":\"EvInventory\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"itemHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EvProfit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isRemoval\",\"type\":\"bool\"}],\"name\":\"EvSigner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"RATE_BASE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"itemHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"cancel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"delegates\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeCapPct\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"feeCapPct_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"weth_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"inventoryStatus\",\"outputs\":[{\"internalType\":\"enumMarket.InvStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"ongoingAuctions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"network\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"intent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delegateType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"currency\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"dataMask\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structMarket.OrderItem[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"signVersion\",\"type\":\"uint8\"}],\"internalType\":\"structMarket.Order[]\",\"name\":\"orders\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumMarket.Op\",\"name\":\"op\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"orderIdx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"itemIdx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"itemHash\",\"type\":\"bytes32\"},{\"internalType\":\"contractIDelegate\",\"name\":\"executionDelegate\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"dataReplacement\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"bidIncentivePct\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"aucMinIncrementPct\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"aucIncDurationSecs\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"percentage\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"internalType\":\"structMarket.Fee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"}],\"internalType\":\"structMarket.SettleDetail[]\",\"name\":\"details\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountToEth\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountToWeth\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"canFail\",\"type\":\"bool\"}],\"internalType\":\"structMarket.SettleShared\",\"name\":\"shared\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"internalType\":\"structMarket.RunInput\",\"name\":\"input\",\"type\":\"tuple\"}],\"name\":\"run\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"network\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"intent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delegateType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"currency\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"dataMask\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structMarket.OrderItem[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"signVersion\",\"type\":\"uint8\"}],\"internalType\":\"structMarket.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountToEth\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountToWeth\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"canFail\",\"type\":\"bool\"}],\"internalType\":\"structMarket.SettleShared\",\"name\":\"shared\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumMarket.Op\",\"name\":\"op\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"orderIdx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"itemIdx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"itemHash\",\"type\":\"bytes32\"},{\"internalType\":\"contractIDelegate\",\"name\":\"executionDelegate\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"dataReplacement\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"bidIncentivePct\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"aucMinIncrementPct\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"aucIncDurationSecs\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"percentage\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"internalType\":\"structMarket.Fee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"}],\"internalType\":\"structMarket.SettleDetail\",\"name\":\"detail\",\"type\":\"tuple\"}],\"name\":\"run1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"signers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"toAdd\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"toRemove\",\"type\":\"address[]\"}],\"name\":\"updateDelegates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"updateFeeCap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"toAdd\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"toRemove\",\"type\":\"address[]\"}],\"name\":\"updateSigners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"weth\",\"outputs\":[{\"internalType\":\"contractIWETHUpgradable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// ExchangeABI is the input ABI used to generate the binding from.
// Deprecated: Use ExchangeMetaData.ABI instead.
var ExchangeABI = ExchangeMetaData.ABI

// Exchange is an auto generated Go binding around an Ethereum contract.
type Exchange struct {
	ExchangeCaller     // Read-only binding to the contract
	ExchangeTransactor // Write-only binding to the contract
	ExchangeFilterer   // Log filterer for contract events
}

// ExchangeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExchangeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExchangeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExchangeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExchangeSession struct {
	Contract     *Exchange         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExchangeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExchangeCallerSession struct {
	Contract *ExchangeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ExchangeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExchangeTransactorSession struct {
	Contract     *ExchangeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ExchangeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExchangeRaw struct {
	Contract *Exchange // Generic contract binding to access the raw methods on
}

// ExchangeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExchangeCallerRaw struct {
	Contract *ExchangeCaller // Generic read-only contract binding to access the raw methods on
}

// ExchangeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExchangeTransactorRaw struct {
	Contract *ExchangeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExchange creates a new instance of Exchange, bound to a specific deployed contract.
func NewExchange(address common.Address, backend bind.ContractBackend) (*Exchange, error) {
	contract, err := bindExchange(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Exchange{ExchangeCaller: ExchangeCaller{contract: contract}, ExchangeTransactor: ExchangeTransactor{contract: contract}, ExchangeFilterer: ExchangeFilterer{contract: contract}}, nil
}

// NewExchangeCaller creates a new read-only instance of Exchange, bound to a specific deployed contract.
func NewExchangeCaller(address common.Address, caller bind.ContractCaller) (*ExchangeCaller, error) {
	contract, err := bindExchange(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExchangeCaller{contract: contract}, nil
}

// NewExchangeTransactor creates a new write-only instance of Exchange, bound to a specific deployed contract.
func NewExchangeTransactor(address common.Address, transactor bind.ContractTransactor) (*ExchangeTransactor, error) {
	contract, err := bindExchange(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExchangeTransactor{contract: contract}, nil
}

// NewExchangeFilterer creates a new log filterer instance of Exchange, bound to a specific deployed contract.
func NewExchangeFilterer(address common.Address, filterer bind.ContractFilterer) (*ExchangeFilterer, error) {
	contract, err := bindExchange(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExchangeFilterer{contract: contract}, nil
}

// bindExchange binds a generic wrapper to an already deployed contract.
func bindExchange(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ExchangeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Exchange *ExchangeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Exchange.Contract.ExchangeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Exchange *ExchangeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Exchange.Contract.ExchangeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Exchange *ExchangeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Exchange.Contract.ExchangeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Exchange *ExchangeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Exchange.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Exchange *ExchangeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Exchange.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Exchange *ExchangeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Exchange.Contract.contract.Transact(opts, method, params...)
}

// RATEBASE is a free data retrieval call binding the contract method 0x0873c6ec.
//
// Solidity: function RATE_BASE() view returns(uint256)
func (_Exchange *ExchangeCaller) RATEBASE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "RATE_BASE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RATEBASE is a free data retrieval call binding the contract method 0x0873c6ec.
//
// Solidity: function RATE_BASE() view returns(uint256)
func (_Exchange *ExchangeSession) RATEBASE() (*big.Int, error) {
	return _Exchange.Contract.RATEBASE(&_Exchange.CallOpts)
}

// RATEBASE is a free data retrieval call binding the contract method 0x0873c6ec.
//
// Solidity: function RATE_BASE() view returns(uint256)
func (_Exchange *ExchangeCallerSession) RATEBASE() (*big.Int, error) {
	return _Exchange.Contract.RATEBASE(&_Exchange.CallOpts)
}

// Delegates is a free data retrieval call binding the contract method 0x587cde1e.
//
// Solidity: function delegates(address ) view returns(bool)
func (_Exchange *ExchangeCaller) Delegates(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "delegates", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Delegates is a free data retrieval call binding the contract method 0x587cde1e.
//
// Solidity: function delegates(address ) view returns(bool)
func (_Exchange *ExchangeSession) Delegates(arg0 common.Address) (bool, error) {
	return _Exchange.Contract.Delegates(&_Exchange.CallOpts, arg0)
}

// Delegates is a free data retrieval call binding the contract method 0x587cde1e.
//
// Solidity: function delegates(address ) view returns(bool)
func (_Exchange *ExchangeCallerSession) Delegates(arg0 common.Address) (bool, error) {
	return _Exchange.Contract.Delegates(&_Exchange.CallOpts, arg0)
}

// FeeCapPct is a free data retrieval call binding the contract method 0x9fb51467.
//
// Solidity: function feeCapPct() view returns(uint256)
func (_Exchange *ExchangeCaller) FeeCapPct(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "feeCapPct")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeCapPct is a free data retrieval call binding the contract method 0x9fb51467.
//
// Solidity: function feeCapPct() view returns(uint256)
func (_Exchange *ExchangeSession) FeeCapPct() (*big.Int, error) {
	return _Exchange.Contract.FeeCapPct(&_Exchange.CallOpts)
}

// FeeCapPct is a free data retrieval call binding the contract method 0x9fb51467.
//
// Solidity: function feeCapPct() view returns(uint256)
func (_Exchange *ExchangeCallerSession) FeeCapPct() (*big.Int, error) {
	return _Exchange.Contract.FeeCapPct(&_Exchange.CallOpts)
}

// InventoryStatus is a free data retrieval call binding the contract method 0x912c860c.
//
// Solidity: function inventoryStatus(bytes32 ) view returns(uint8)
func (_Exchange *ExchangeCaller) InventoryStatus(opts *bind.CallOpts, arg0 [32]byte) (uint8, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "inventoryStatus", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// InventoryStatus is a free data retrieval call binding the contract method 0x912c860c.
//
// Solidity: function inventoryStatus(bytes32 ) view returns(uint8)
func (_Exchange *ExchangeSession) InventoryStatus(arg0 [32]byte) (uint8, error) {
	return _Exchange.Contract.InventoryStatus(&_Exchange.CallOpts, arg0)
}

// InventoryStatus is a free data retrieval call binding the contract method 0x912c860c.
//
// Solidity: function inventoryStatus(bytes32 ) view returns(uint8)
func (_Exchange *ExchangeCallerSession) InventoryStatus(arg0 [32]byte) (uint8, error) {
	return _Exchange.Contract.InventoryStatus(&_Exchange.CallOpts, arg0)
}

// OngoingAuctions is a free data retrieval call binding the contract method 0xea805917.
//
// Solidity: function ongoingAuctions(bytes32 ) view returns(uint256 price, uint256 netPrice, uint256 endAt, address bidder)
func (_Exchange *ExchangeCaller) OngoingAuctions(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Price    *big.Int
	NetPrice *big.Int
	EndAt    *big.Int
	Bidder   common.Address
}, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "ongoingAuctions", arg0)

	outstruct := new(struct {
		Price    *big.Int
		NetPrice *big.Int
		EndAt    *big.Int
		Bidder   common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Price = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.NetPrice = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.EndAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Bidder = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// OngoingAuctions is a free data retrieval call binding the contract method 0xea805917.
//
// Solidity: function ongoingAuctions(bytes32 ) view returns(uint256 price, uint256 netPrice, uint256 endAt, address bidder)
func (_Exchange *ExchangeSession) OngoingAuctions(arg0 [32]byte) (struct {
	Price    *big.Int
	NetPrice *big.Int
	EndAt    *big.Int
	Bidder   common.Address
}, error) {
	return _Exchange.Contract.OngoingAuctions(&_Exchange.CallOpts, arg0)
}

// OngoingAuctions is a free data retrieval call binding the contract method 0xea805917.
//
// Solidity: function ongoingAuctions(bytes32 ) view returns(uint256 price, uint256 netPrice, uint256 endAt, address bidder)
func (_Exchange *ExchangeCallerSession) OngoingAuctions(arg0 [32]byte) (struct {
	Price    *big.Int
	NetPrice *big.Int
	EndAt    *big.Int
	Bidder   common.Address
}, error) {
	return _Exchange.Contract.OngoingAuctions(&_Exchange.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Exchange *ExchangeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Exchange *ExchangeSession) Owner() (common.Address, error) {
	return _Exchange.Contract.Owner(&_Exchange.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Exchange *ExchangeCallerSession) Owner() (common.Address, error) {
	return _Exchange.Contract.Owner(&_Exchange.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Exchange *ExchangeCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Exchange *ExchangeSession) Paused() (bool, error) {
	return _Exchange.Contract.Paused(&_Exchange.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Exchange *ExchangeCallerSession) Paused() (bool, error) {
	return _Exchange.Contract.Paused(&_Exchange.CallOpts)
}

// Signers is a free data retrieval call binding the contract method 0x736c0d5b.
//
// Solidity: function signers(address ) view returns(bool)
func (_Exchange *ExchangeCaller) Signers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "signers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Signers is a free data retrieval call binding the contract method 0x736c0d5b.
//
// Solidity: function signers(address ) view returns(bool)
func (_Exchange *ExchangeSession) Signers(arg0 common.Address) (bool, error) {
	return _Exchange.Contract.Signers(&_Exchange.CallOpts, arg0)
}

// Signers is a free data retrieval call binding the contract method 0x736c0d5b.
//
// Solidity: function signers(address ) view returns(bool)
func (_Exchange *ExchangeCallerSession) Signers(arg0 common.Address) (bool, error) {
	return _Exchange.Contract.Signers(&_Exchange.CallOpts, arg0)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_Exchange *ExchangeCaller) Weth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "weth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_Exchange *ExchangeSession) Weth() (common.Address, error) {
	return _Exchange.Contract.Weth(&_Exchange.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_Exchange *ExchangeCallerSession) Weth() (common.Address, error) {
	return _Exchange.Contract.Weth(&_Exchange.CallOpts)
}

// Cancel is a paid mutator transaction binding the contract method 0x2295f9bf.
//
// Solidity: function cancel(bytes32[] itemHashes, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Exchange *ExchangeTransactor) Cancel(opts *bind.TransactOpts, itemHashes [][32]byte, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "cancel", itemHashes, deadline, v, r, s)
}

// Cancel is a paid mutator transaction binding the contract method 0x2295f9bf.
//
// Solidity: function cancel(bytes32[] itemHashes, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Exchange *ExchangeSession) Cancel(itemHashes [][32]byte, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Exchange.Contract.Cancel(&_Exchange.TransactOpts, itemHashes, deadline, v, r, s)
}

// Cancel is a paid mutator transaction binding the contract method 0x2295f9bf.
//
// Solidity: function cancel(bytes32[] itemHashes, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Exchange *ExchangeTransactorSession) Cancel(itemHashes [][32]byte, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Exchange.Contract.Cancel(&_Exchange.TransactOpts, itemHashes, deadline, v, r, s)
}

// Initialize is a paid mutator transaction binding the contract method 0xda35a26f.
//
// Solidity: function initialize(uint256 feeCapPct_, address weth_) returns()
func (_Exchange *ExchangeTransactor) Initialize(opts *bind.TransactOpts, feeCapPct_ *big.Int, weth_ common.Address) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "initialize", feeCapPct_, weth_)
}

// Initialize is a paid mutator transaction binding the contract method 0xda35a26f.
//
// Solidity: function initialize(uint256 feeCapPct_, address weth_) returns()
func (_Exchange *ExchangeSession) Initialize(feeCapPct_ *big.Int, weth_ common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.Initialize(&_Exchange.TransactOpts, feeCapPct_, weth_)
}

// Initialize is a paid mutator transaction binding the contract method 0xda35a26f.
//
// Solidity: function initialize(uint256 feeCapPct_, address weth_) returns()
func (_Exchange *ExchangeTransactorSession) Initialize(feeCapPct_ *big.Int, weth_ common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.Initialize(&_Exchange.TransactOpts, feeCapPct_, weth_)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Exchange *ExchangeTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Exchange *ExchangeSession) Pause() (*types.Transaction, error) {
	return _Exchange.Contract.Pause(&_Exchange.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Exchange *ExchangeTransactorSession) Pause() (*types.Transaction, error) {
	return _Exchange.Contract.Pause(&_Exchange.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Exchange *ExchangeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Exchange *ExchangeSession) RenounceOwnership() (*types.Transaction, error) {
	return _Exchange.Contract.RenounceOwnership(&_Exchange.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Exchange *ExchangeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Exchange.Contract.RenounceOwnership(&_Exchange.TransactOpts)
}

// Run is a paid mutator transaction binding the contract method 0x357a150b.
//
// Solidity: function run(((uint256,address,uint256,uint256,uint256,uint256,address,bytes,(uint256,bytes)[],bytes32,bytes32,uint8,uint8)[],(uint8,uint256,uint256,uint256,bytes32,address,bytes,uint256,uint256,uint256,(uint256,address)[])[],(uint256,uint256,uint256,uint256,address,bool),bytes32,bytes32,uint8) input) payable returns()
func (_Exchange *ExchangeTransactor) Run(opts *bind.TransactOpts, input MarketRunInput) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "run", input)
}

// Run is a paid mutator transaction binding the contract method 0x357a150b.
//
// Solidity: function run(((uint256,address,uint256,uint256,uint256,uint256,address,bytes,(uint256,bytes)[],bytes32,bytes32,uint8,uint8)[],(uint8,uint256,uint256,uint256,bytes32,address,bytes,uint256,uint256,uint256,(uint256,address)[])[],(uint256,uint256,uint256,uint256,address,bool),bytes32,bytes32,uint8) input) payable returns()
func (_Exchange *ExchangeSession) Run(input MarketRunInput) (*types.Transaction, error) {
	return _Exchange.Contract.Run(&_Exchange.TransactOpts, input)
}

// Run is a paid mutator transaction binding the contract method 0x357a150b.
//
// Solidity: function run(((uint256,address,uint256,uint256,uint256,uint256,address,bytes,(uint256,bytes)[],bytes32,bytes32,uint8,uint8)[],(uint8,uint256,uint256,uint256,bytes32,address,bytes,uint256,uint256,uint256,(uint256,address)[])[],(uint256,uint256,uint256,uint256,address,bool),bytes32,bytes32,uint8) input) payable returns()
func (_Exchange *ExchangeTransactorSession) Run(input MarketRunInput) (*types.Transaction, error) {
	return _Exchange.Contract.Run(&_Exchange.TransactOpts, input)
}

// Run1 is a paid mutator transaction binding the contract method 0xd95e3c54.
//
// Solidity: function run1((uint256,address,uint256,uint256,uint256,uint256,address,bytes,(uint256,bytes)[],bytes32,bytes32,uint8,uint8) order, (uint256,uint256,uint256,uint256,address,bool) shared, (uint8,uint256,uint256,uint256,bytes32,address,bytes,uint256,uint256,uint256,(uint256,address)[]) detail) returns(uint256)
func (_Exchange *ExchangeTransactor) Run1(opts *bind.TransactOpts, order MarketOrder, shared MarketSettleShared, detail MarketSettleDetail) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "run1", order, shared, detail)
}

// Run1 is a paid mutator transaction binding the contract method 0xd95e3c54.
//
// Solidity: function run1((uint256,address,uint256,uint256,uint256,uint256,address,bytes,(uint256,bytes)[],bytes32,bytes32,uint8,uint8) order, (uint256,uint256,uint256,uint256,address,bool) shared, (uint8,uint256,uint256,uint256,bytes32,address,bytes,uint256,uint256,uint256,(uint256,address)[]) detail) returns(uint256)
func (_Exchange *ExchangeSession) Run1(order MarketOrder, shared MarketSettleShared, detail MarketSettleDetail) (*types.Transaction, error) {
	return _Exchange.Contract.Run1(&_Exchange.TransactOpts, order, shared, detail)
}

// Run1 is a paid mutator transaction binding the contract method 0xd95e3c54.
//
// Solidity: function run1((uint256,address,uint256,uint256,uint256,uint256,address,bytes,(uint256,bytes)[],bytes32,bytes32,uint8,uint8) order, (uint256,uint256,uint256,uint256,address,bool) shared, (uint8,uint256,uint256,uint256,bytes32,address,bytes,uint256,uint256,uint256,(uint256,address)[]) detail) returns(uint256)
func (_Exchange *ExchangeTransactorSession) Run1(order MarketOrder, shared MarketSettleShared, detail MarketSettleDetail) (*types.Transaction, error) {
	return _Exchange.Contract.Run1(&_Exchange.TransactOpts, order, shared, detail)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Exchange *ExchangeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Exchange *ExchangeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.TransferOwnership(&_Exchange.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Exchange *ExchangeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.TransferOwnership(&_Exchange.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Exchange *ExchangeTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Exchange *ExchangeSession) Unpause() (*types.Transaction, error) {
	return _Exchange.Contract.Unpause(&_Exchange.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Exchange *ExchangeTransactorSession) Unpause() (*types.Transaction, error) {
	return _Exchange.Contract.Unpause(&_Exchange.TransactOpts)
}

// UpdateDelegates is a paid mutator transaction binding the contract method 0xe59f739a.
//
// Solidity: function updateDelegates(address[] toAdd, address[] toRemove) returns()
func (_Exchange *ExchangeTransactor) UpdateDelegates(opts *bind.TransactOpts, toAdd []common.Address, toRemove []common.Address) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "updateDelegates", toAdd, toRemove)
}

// UpdateDelegates is a paid mutator transaction binding the contract method 0xe59f739a.
//
// Solidity: function updateDelegates(address[] toAdd, address[] toRemove) returns()
func (_Exchange *ExchangeSession) UpdateDelegates(toAdd []common.Address, toRemove []common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.UpdateDelegates(&_Exchange.TransactOpts, toAdd, toRemove)
}

// UpdateDelegates is a paid mutator transaction binding the contract method 0xe59f739a.
//
// Solidity: function updateDelegates(address[] toAdd, address[] toRemove) returns()
func (_Exchange *ExchangeTransactorSession) UpdateDelegates(toAdd []common.Address, toRemove []common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.UpdateDelegates(&_Exchange.TransactOpts, toAdd, toRemove)
}

// UpdateFeeCap is a paid mutator transaction binding the contract method 0x95835fea.
//
// Solidity: function updateFeeCap(uint256 val) returns()
func (_Exchange *ExchangeTransactor) UpdateFeeCap(opts *bind.TransactOpts, val *big.Int) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "updateFeeCap", val)
}

// UpdateFeeCap is a paid mutator transaction binding the contract method 0x95835fea.
//
// Solidity: function updateFeeCap(uint256 val) returns()
func (_Exchange *ExchangeSession) UpdateFeeCap(val *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.UpdateFeeCap(&_Exchange.TransactOpts, val)
}

// UpdateFeeCap is a paid mutator transaction binding the contract method 0x95835fea.
//
// Solidity: function updateFeeCap(uint256 val) returns()
func (_Exchange *ExchangeTransactorSession) UpdateFeeCap(val *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.UpdateFeeCap(&_Exchange.TransactOpts, val)
}

// UpdateSigners is a paid mutator transaction binding the contract method 0x350b2369.
//
// Solidity: function updateSigners(address[] toAdd, address[] toRemove) returns()
func (_Exchange *ExchangeTransactor) UpdateSigners(opts *bind.TransactOpts, toAdd []common.Address, toRemove []common.Address) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "updateSigners", toAdd, toRemove)
}

// UpdateSigners is a paid mutator transaction binding the contract method 0x350b2369.
//
// Solidity: function updateSigners(address[] toAdd, address[] toRemove) returns()
func (_Exchange *ExchangeSession) UpdateSigners(toAdd []common.Address, toRemove []common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.UpdateSigners(&_Exchange.TransactOpts, toAdd, toRemove)
}

// UpdateSigners is a paid mutator transaction binding the contract method 0x350b2369.
//
// Solidity: function updateSigners(address[] toAdd, address[] toRemove) returns()
func (_Exchange *ExchangeTransactorSession) UpdateSigners(toAdd []common.Address, toRemove []common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.UpdateSigners(&_Exchange.TransactOpts, toAdd, toRemove)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Exchange *ExchangeTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Exchange.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Exchange *ExchangeSession) Receive() (*types.Transaction, error) {
	return _Exchange.Contract.Receive(&_Exchange.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Exchange *ExchangeTransactorSession) Receive() (*types.Transaction, error) {
	return _Exchange.Contract.Receive(&_Exchange.TransactOpts)
}

// ExchangeEvAuctionRefundIterator is returned from FilterEvAuctionRefund and is used to iterate over the raw logs and unpacked data for EvAuctionRefund events raised by the Exchange contract.
type ExchangeEvAuctionRefundIterator struct {
	Event *ExchangeEvAuctionRefund // Event containing the contract specifics and raw log

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
func (it *ExchangeEvAuctionRefundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeEvAuctionRefund)
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
		it.Event = new(ExchangeEvAuctionRefund)
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
func (it *ExchangeEvAuctionRefundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeEvAuctionRefundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeEvAuctionRefund represents a EvAuctionRefund event raised by the Exchange contract.
type ExchangeEvAuctionRefund struct {
	ItemHash  [32]byte
	Currency  common.Address
	To        common.Address
	Amount    *big.Int
	Incentive *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEvAuctionRefund is a free log retrieval operation binding the contract event 0x681e2055b67e23ce693a446bd0567fb9df559ce6f82da4397482bad968551ac2.
//
// Solidity: event EvAuctionRefund(bytes32 indexed itemHash, address currency, address to, uint256 amount, uint256 incentive)
func (_Exchange *ExchangeFilterer) FilterEvAuctionRefund(opts *bind.FilterOpts, itemHash [][32]byte) (*ExchangeEvAuctionRefundIterator, error) {

	var itemHashRule []interface{}
	for _, itemHashItem := range itemHash {
		itemHashRule = append(itemHashRule, itemHashItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "EvAuctionRefund", itemHashRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeEvAuctionRefundIterator{contract: _Exchange.contract, event: "EvAuctionRefund", logs: logs, sub: sub}, nil
}

// WatchEvAuctionRefund is a free log subscription operation binding the contract event 0x681e2055b67e23ce693a446bd0567fb9df559ce6f82da4397482bad968551ac2.
//
// Solidity: event EvAuctionRefund(bytes32 indexed itemHash, address currency, address to, uint256 amount, uint256 incentive)
func (_Exchange *ExchangeFilterer) WatchEvAuctionRefund(opts *bind.WatchOpts, sink chan<- *ExchangeEvAuctionRefund, itemHash [][32]byte) (event.Subscription, error) {

	var itemHashRule []interface{}
	for _, itemHashItem := range itemHash {
		itemHashRule = append(itemHashRule, itemHashItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "EvAuctionRefund", itemHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeEvAuctionRefund)
				if err := _Exchange.contract.UnpackLog(event, "EvAuctionRefund", log); err != nil {
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

// ParseEvAuctionRefund is a log parse operation binding the contract event 0x681e2055b67e23ce693a446bd0567fb9df559ce6f82da4397482bad968551ac2.
//
// Solidity: event EvAuctionRefund(bytes32 indexed itemHash, address currency, address to, uint256 amount, uint256 incentive)
func (_Exchange *ExchangeFilterer) ParseEvAuctionRefund(log types.Log) (*ExchangeEvAuctionRefund, error) {
	event := new(ExchangeEvAuctionRefund)
	if err := _Exchange.contract.UnpackLog(event, "EvAuctionRefund", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeEvCancelIterator is returned from FilterEvCancel and is used to iterate over the raw logs and unpacked data for EvCancel events raised by the Exchange contract.
type ExchangeEvCancelIterator struct {
	Event *ExchangeEvCancel // Event containing the contract specifics and raw log

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
func (it *ExchangeEvCancelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeEvCancel)
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
		it.Event = new(ExchangeEvCancel)
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
func (it *ExchangeEvCancelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeEvCancelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeEvCancel represents a EvCancel event raised by the Exchange contract.
type ExchangeEvCancel struct {
	ItemHash [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterEvCancel is a free log retrieval operation binding the contract event 0x5b0b06d07e20243724d90e17a20034972f339eb28bd1c9437a71999bd15a1e7a.
//
// Solidity: event EvCancel(bytes32 indexed itemHash)
func (_Exchange *ExchangeFilterer) FilterEvCancel(opts *bind.FilterOpts, itemHash [][32]byte) (*ExchangeEvCancelIterator, error) {

	var itemHashRule []interface{}
	for _, itemHashItem := range itemHash {
		itemHashRule = append(itemHashRule, itemHashItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "EvCancel", itemHashRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeEvCancelIterator{contract: _Exchange.contract, event: "EvCancel", logs: logs, sub: sub}, nil
}

// WatchEvCancel is a free log subscription operation binding the contract event 0x5b0b06d07e20243724d90e17a20034972f339eb28bd1c9437a71999bd15a1e7a.
//
// Solidity: event EvCancel(bytes32 indexed itemHash)
func (_Exchange *ExchangeFilterer) WatchEvCancel(opts *bind.WatchOpts, sink chan<- *ExchangeEvCancel, itemHash [][32]byte) (event.Subscription, error) {

	var itemHashRule []interface{}
	for _, itemHashItem := range itemHash {
		itemHashRule = append(itemHashRule, itemHashItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "EvCancel", itemHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeEvCancel)
				if err := _Exchange.contract.UnpackLog(event, "EvCancel", log); err != nil {
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

// ParseEvCancel is a log parse operation binding the contract event 0x5b0b06d07e20243724d90e17a20034972f339eb28bd1c9437a71999bd15a1e7a.
//
// Solidity: event EvCancel(bytes32 indexed itemHash)
func (_Exchange *ExchangeFilterer) ParseEvCancel(log types.Log) (*ExchangeEvCancel, error) {
	event := new(ExchangeEvCancel)
	if err := _Exchange.contract.UnpackLog(event, "EvCancel", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeEvDelegateIterator is returned from FilterEvDelegate and is used to iterate over the raw logs and unpacked data for EvDelegate events raised by the Exchange contract.
type ExchangeEvDelegateIterator struct {
	Event *ExchangeEvDelegate // Event containing the contract specifics and raw log

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
func (it *ExchangeEvDelegateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeEvDelegate)
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
		it.Event = new(ExchangeEvDelegate)
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
func (it *ExchangeEvDelegateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeEvDelegateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeEvDelegate represents a EvDelegate event raised by the Exchange contract.
type ExchangeEvDelegate struct {
	Delegate  common.Address
	IsRemoval bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEvDelegate is a free log retrieval operation binding the contract event 0x4a31a64b928a0e8aff42ef84d144ffe82d08cb41c8027060593135e2026899b2.
//
// Solidity: event EvDelegate(address delegate, bool isRemoval)
func (_Exchange *ExchangeFilterer) FilterEvDelegate(opts *bind.FilterOpts) (*ExchangeEvDelegateIterator, error) {

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "EvDelegate")
	if err != nil {
		return nil, err
	}
	return &ExchangeEvDelegateIterator{contract: _Exchange.contract, event: "EvDelegate", logs: logs, sub: sub}, nil
}

// WatchEvDelegate is a free log subscription operation binding the contract event 0x4a31a64b928a0e8aff42ef84d144ffe82d08cb41c8027060593135e2026899b2.
//
// Solidity: event EvDelegate(address delegate, bool isRemoval)
func (_Exchange *ExchangeFilterer) WatchEvDelegate(opts *bind.WatchOpts, sink chan<- *ExchangeEvDelegate) (event.Subscription, error) {

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "EvDelegate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeEvDelegate)
				if err := _Exchange.contract.UnpackLog(event, "EvDelegate", log); err != nil {
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

// ParseEvDelegate is a log parse operation binding the contract event 0x4a31a64b928a0e8aff42ef84d144ffe82d08cb41c8027060593135e2026899b2.
//
// Solidity: event EvDelegate(address delegate, bool isRemoval)
func (_Exchange *ExchangeFilterer) ParseEvDelegate(log types.Log) (*ExchangeEvDelegate, error) {
	event := new(ExchangeEvDelegate)
	if err := _Exchange.contract.UnpackLog(event, "EvDelegate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeEvFailureIterator is returned from FilterEvFailure and is used to iterate over the raw logs and unpacked data for EvFailure events raised by the Exchange contract.
type ExchangeEvFailureIterator struct {
	Event *ExchangeEvFailure // Event containing the contract specifics and raw log

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
func (it *ExchangeEvFailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeEvFailure)
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
		it.Event = new(ExchangeEvFailure)
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
func (it *ExchangeEvFailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeEvFailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeEvFailure represents a EvFailure event raised by the Exchange contract.
type ExchangeEvFailure struct {
	Index *big.Int
	Error []byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterEvFailure is a free log retrieval operation binding the contract event 0x97c789f43a3e7ac27906b5fbdac832f54441771021fba06f71207d9be6d4b623.
//
// Solidity: event EvFailure(uint256 index, bytes error)
func (_Exchange *ExchangeFilterer) FilterEvFailure(opts *bind.FilterOpts) (*ExchangeEvFailureIterator, error) {

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "EvFailure")
	if err != nil {
		return nil, err
	}
	return &ExchangeEvFailureIterator{contract: _Exchange.contract, event: "EvFailure", logs: logs, sub: sub}, nil
}

// WatchEvFailure is a free log subscription operation binding the contract event 0x97c789f43a3e7ac27906b5fbdac832f54441771021fba06f71207d9be6d4b623.
//
// Solidity: event EvFailure(uint256 index, bytes error)
func (_Exchange *ExchangeFilterer) WatchEvFailure(opts *bind.WatchOpts, sink chan<- *ExchangeEvFailure) (event.Subscription, error) {

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "EvFailure")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeEvFailure)
				if err := _Exchange.contract.UnpackLog(event, "EvFailure", log); err != nil {
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

// ParseEvFailure is a log parse operation binding the contract event 0x97c789f43a3e7ac27906b5fbdac832f54441771021fba06f71207d9be6d4b623.
//
// Solidity: event EvFailure(uint256 index, bytes error)
func (_Exchange *ExchangeFilterer) ParseEvFailure(log types.Log) (*ExchangeEvFailure, error) {
	event := new(ExchangeEvFailure)
	if err := _Exchange.contract.UnpackLog(event, "EvFailure", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeEvFeeCapUpdateIterator is returned from FilterEvFeeCapUpdate and is used to iterate over the raw logs and unpacked data for EvFeeCapUpdate events raised by the Exchange contract.
type ExchangeEvFeeCapUpdateIterator struct {
	Event *ExchangeEvFeeCapUpdate // Event containing the contract specifics and raw log

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
func (it *ExchangeEvFeeCapUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeEvFeeCapUpdate)
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
		it.Event = new(ExchangeEvFeeCapUpdate)
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
func (it *ExchangeEvFeeCapUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeEvFeeCapUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeEvFeeCapUpdate represents a EvFeeCapUpdate event raised by the Exchange contract.
type ExchangeEvFeeCapUpdate struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterEvFeeCapUpdate is a free log retrieval operation binding the contract event 0x19fc3beddeea399f0966d5f8664ad94006f16a10fb28c4e2fe6fae62626b7128.
//
// Solidity: event EvFeeCapUpdate(uint256 newValue)
func (_Exchange *ExchangeFilterer) FilterEvFeeCapUpdate(opts *bind.FilterOpts) (*ExchangeEvFeeCapUpdateIterator, error) {

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "EvFeeCapUpdate")
	if err != nil {
		return nil, err
	}
	return &ExchangeEvFeeCapUpdateIterator{contract: _Exchange.contract, event: "EvFeeCapUpdate", logs: logs, sub: sub}, nil
}

// WatchEvFeeCapUpdate is a free log subscription operation binding the contract event 0x19fc3beddeea399f0966d5f8664ad94006f16a10fb28c4e2fe6fae62626b7128.
//
// Solidity: event EvFeeCapUpdate(uint256 newValue)
func (_Exchange *ExchangeFilterer) WatchEvFeeCapUpdate(opts *bind.WatchOpts, sink chan<- *ExchangeEvFeeCapUpdate) (event.Subscription, error) {

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "EvFeeCapUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeEvFeeCapUpdate)
				if err := _Exchange.contract.UnpackLog(event, "EvFeeCapUpdate", log); err != nil {
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

// ParseEvFeeCapUpdate is a log parse operation binding the contract event 0x19fc3beddeea399f0966d5f8664ad94006f16a10fb28c4e2fe6fae62626b7128.
//
// Solidity: event EvFeeCapUpdate(uint256 newValue)
func (_Exchange *ExchangeFilterer) ParseEvFeeCapUpdate(log types.Log) (*ExchangeEvFeeCapUpdate, error) {
	event := new(ExchangeEvFeeCapUpdate)
	if err := _Exchange.contract.UnpackLog(event, "EvFeeCapUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeEvInventoryIterator is returned from FilterEvInventory and is used to iterate over the raw logs and unpacked data for EvInventory events raised by the Exchange contract.
type ExchangeEvInventoryIterator struct {
	Event *ExchangeEvInventory // Event containing the contract specifics and raw log

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
func (it *ExchangeEvInventoryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeEvInventory)
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
		it.Event = new(ExchangeEvInventory)
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
func (it *ExchangeEvInventoryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeEvInventoryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeEvInventory represents a EvInventory event raised by the Exchange contract.
type ExchangeEvInventory struct {
	ItemHash     [32]byte
	Maker        common.Address
	Taker        common.Address
	OrderSalt    *big.Int
	SettleSalt   *big.Int
	Intent       *big.Int
	DelegateType *big.Int
	Deadline     *big.Int
	Currency     common.Address
	DataMask     []byte
	Item         MarketOrderItem
	Detail       MarketSettleDetail
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterEvInventory is a free log retrieval operation binding the contract event 0x3cbb63f144840e5b1b0a38a7c19211d2e89de4d7c5faf8b2d3c1776c302d1d33.
//
// Solidity: event EvInventory(bytes32 indexed itemHash, address maker, address taker, uint256 orderSalt, uint256 settleSalt, uint256 intent, uint256 delegateType, uint256 deadline, address currency, bytes dataMask, (uint256,bytes) item, (uint8,uint256,uint256,uint256,bytes32,address,bytes,uint256,uint256,uint256,(uint256,address)[]) detail)
func (_Exchange *ExchangeFilterer) FilterEvInventory(opts *bind.FilterOpts, itemHash [][32]byte) (*ExchangeEvInventoryIterator, error) {

	var itemHashRule []interface{}
	for _, itemHashItem := range itemHash {
		itemHashRule = append(itemHashRule, itemHashItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "EvInventory", itemHashRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeEvInventoryIterator{contract: _Exchange.contract, event: "EvInventory", logs: logs, sub: sub}, nil
}

// WatchEvInventory is a free log subscription operation binding the contract event 0x3cbb63f144840e5b1b0a38a7c19211d2e89de4d7c5faf8b2d3c1776c302d1d33.
//
// Solidity: event EvInventory(bytes32 indexed itemHash, address maker, address taker, uint256 orderSalt, uint256 settleSalt, uint256 intent, uint256 delegateType, uint256 deadline, address currency, bytes dataMask, (uint256,bytes) item, (uint8,uint256,uint256,uint256,bytes32,address,bytes,uint256,uint256,uint256,(uint256,address)[]) detail)
func (_Exchange *ExchangeFilterer) WatchEvInventory(opts *bind.WatchOpts, sink chan<- *ExchangeEvInventory, itemHash [][32]byte) (event.Subscription, error) {

	var itemHashRule []interface{}
	for _, itemHashItem := range itemHash {
		itemHashRule = append(itemHashRule, itemHashItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "EvInventory", itemHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeEvInventory)
				if err := _Exchange.contract.UnpackLog(event, "EvInventory", log); err != nil {
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

// ParseEvInventory is a log parse operation binding the contract event 0x3cbb63f144840e5b1b0a38a7c19211d2e89de4d7c5faf8b2d3c1776c302d1d33.
//
// Solidity: event EvInventory(bytes32 indexed itemHash, address maker, address taker, uint256 orderSalt, uint256 settleSalt, uint256 intent, uint256 delegateType, uint256 deadline, address currency, bytes dataMask, (uint256,bytes) item, (uint8,uint256,uint256,uint256,bytes32,address,bytes,uint256,uint256,uint256,(uint256,address)[]) detail)
func (_Exchange *ExchangeFilterer) ParseEvInventory(log types.Log) (*ExchangeEvInventory, error) {
	event := new(ExchangeEvInventory)
	if err := _Exchange.contract.UnpackLog(event, "EvInventory", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeEvProfitIterator is returned from FilterEvProfit and is used to iterate over the raw logs and unpacked data for EvProfit events raised by the Exchange contract.
type ExchangeEvProfitIterator struct {
	Event *ExchangeEvProfit // Event containing the contract specifics and raw log

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
func (it *ExchangeEvProfitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeEvProfit)
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
		it.Event = new(ExchangeEvProfit)
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
func (it *ExchangeEvProfitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeEvProfitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeEvProfit represents a EvProfit event raised by the Exchange contract.
type ExchangeEvProfit struct {
	ItemHash [32]byte
	Currency common.Address
	To       common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterEvProfit is a free log retrieval operation binding the contract event 0xe2c49856b032c255ae7e325d18109bc4e22a2804e2e49a017ec0f59f19cd447b.
//
// Solidity: event EvProfit(bytes32 itemHash, address currency, address to, uint256 amount)
func (_Exchange *ExchangeFilterer) FilterEvProfit(opts *bind.FilterOpts) (*ExchangeEvProfitIterator, error) {

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "EvProfit")
	if err != nil {
		return nil, err
	}
	return &ExchangeEvProfitIterator{contract: _Exchange.contract, event: "EvProfit", logs: logs, sub: sub}, nil
}

// WatchEvProfit is a free log subscription operation binding the contract event 0xe2c49856b032c255ae7e325d18109bc4e22a2804e2e49a017ec0f59f19cd447b.
//
// Solidity: event EvProfit(bytes32 itemHash, address currency, address to, uint256 amount)
func (_Exchange *ExchangeFilterer) WatchEvProfit(opts *bind.WatchOpts, sink chan<- *ExchangeEvProfit) (event.Subscription, error) {

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "EvProfit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeEvProfit)
				if err := _Exchange.contract.UnpackLog(event, "EvProfit", log); err != nil {
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

// ParseEvProfit is a log parse operation binding the contract event 0xe2c49856b032c255ae7e325d18109bc4e22a2804e2e49a017ec0f59f19cd447b.
//
// Solidity: event EvProfit(bytes32 itemHash, address currency, address to, uint256 amount)
func (_Exchange *ExchangeFilterer) ParseEvProfit(log types.Log) (*ExchangeEvProfit, error) {
	event := new(ExchangeEvProfit)
	if err := _Exchange.contract.UnpackLog(event, "EvProfit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeEvSignerIterator is returned from FilterEvSigner and is used to iterate over the raw logs and unpacked data for EvSigner events raised by the Exchange contract.
type ExchangeEvSignerIterator struct {
	Event *ExchangeEvSigner // Event containing the contract specifics and raw log

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
func (it *ExchangeEvSignerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeEvSigner)
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
		it.Event = new(ExchangeEvSigner)
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
func (it *ExchangeEvSignerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeEvSignerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeEvSigner represents a EvSigner event raised by the Exchange contract.
type ExchangeEvSigner struct {
	Signer    common.Address
	IsRemoval bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEvSigner is a free log retrieval operation binding the contract event 0x0127aee741cbb6bc48b5475b8eb3eb2e5d053809d551dedd517a0b5b52b80fd5.
//
// Solidity: event EvSigner(address signer, bool isRemoval)
func (_Exchange *ExchangeFilterer) FilterEvSigner(opts *bind.FilterOpts) (*ExchangeEvSignerIterator, error) {

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "EvSigner")
	if err != nil {
		return nil, err
	}
	return &ExchangeEvSignerIterator{contract: _Exchange.contract, event: "EvSigner", logs: logs, sub: sub}, nil
}

// WatchEvSigner is a free log subscription operation binding the contract event 0x0127aee741cbb6bc48b5475b8eb3eb2e5d053809d551dedd517a0b5b52b80fd5.
//
// Solidity: event EvSigner(address signer, bool isRemoval)
func (_Exchange *ExchangeFilterer) WatchEvSigner(opts *bind.WatchOpts, sink chan<- *ExchangeEvSigner) (event.Subscription, error) {

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "EvSigner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeEvSigner)
				if err := _Exchange.contract.UnpackLog(event, "EvSigner", log); err != nil {
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

// ParseEvSigner is a log parse operation binding the contract event 0x0127aee741cbb6bc48b5475b8eb3eb2e5d053809d551dedd517a0b5b52b80fd5.
//
// Solidity: event EvSigner(address signer, bool isRemoval)
func (_Exchange *ExchangeFilterer) ParseEvSigner(log types.Log) (*ExchangeEvSigner, error) {
	event := new(ExchangeEvSigner)
	if err := _Exchange.contract.UnpackLog(event, "EvSigner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Exchange contract.
type ExchangeOwnershipTransferredIterator struct {
	Event *ExchangeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ExchangeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeOwnershipTransferred)
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
		it.Event = new(ExchangeOwnershipTransferred)
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
func (it *ExchangeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeOwnershipTransferred represents a OwnershipTransferred event raised by the Exchange contract.
type ExchangeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Exchange *ExchangeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ExchangeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeOwnershipTransferredIterator{contract: _Exchange.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Exchange *ExchangeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ExchangeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeOwnershipTransferred)
				if err := _Exchange.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Exchange *ExchangeFilterer) ParseOwnershipTransferred(log types.Log) (*ExchangeOwnershipTransferred, error) {
	event := new(ExchangeOwnershipTransferred)
	if err := _Exchange.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Exchange contract.
type ExchangePausedIterator struct {
	Event *ExchangePaused // Event containing the contract specifics and raw log

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
func (it *ExchangePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangePaused)
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
		it.Event = new(ExchangePaused)
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
func (it *ExchangePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangePaused represents a Paused event raised by the Exchange contract.
type ExchangePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Exchange *ExchangeFilterer) FilterPaused(opts *bind.FilterOpts) (*ExchangePausedIterator, error) {

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &ExchangePausedIterator{contract: _Exchange.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Exchange *ExchangeFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ExchangePaused) (event.Subscription, error) {

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangePaused)
				if err := _Exchange.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Exchange *ExchangeFilterer) ParsePaused(log types.Log) (*ExchangePaused, error) {
	event := new(ExchangePaused)
	if err := _Exchange.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Exchange contract.
type ExchangeUnpausedIterator struct {
	Event *ExchangeUnpaused // Event containing the contract specifics and raw log

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
func (it *ExchangeUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeUnpaused)
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
		it.Event = new(ExchangeUnpaused)
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
func (it *ExchangeUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeUnpaused represents a Unpaused event raised by the Exchange contract.
type ExchangeUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Exchange *ExchangeFilterer) FilterUnpaused(opts *bind.FilterOpts) (*ExchangeUnpausedIterator, error) {

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &ExchangeUnpausedIterator{contract: _Exchange.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Exchange *ExchangeFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ExchangeUnpaused) (event.Subscription, error) {

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeUnpaused)
				if err := _Exchange.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Exchange *ExchangeFilterer) ParseUnpaused(log types.Log) (*ExchangeUnpaused, error) {
	event := new(ExchangeUnpaused)
	if err := _Exchange.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
