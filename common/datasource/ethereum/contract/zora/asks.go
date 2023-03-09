// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zora

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

// AsksV11Ask is an auto generated low-level Go binding around an user-defined struct.
type AsksV11Ask struct {
	Seller               common.Address
	SellerFundsRecipient common.Address
	AskCurrency          common.Address
	FindersFeeBps        uint16
	AskPrice             *big.Int
}

// UniversalExchangeEventV1ExchangeDetails is an auto generated low-level Go binding around an user-defined struct.
type UniversalExchangeEventV1ExchangeDetails struct {
	TokenContract common.Address
	TokenId       *big.Int
	Amount        *big.Int
}

// ZoraMetaData contains all meta data concerning the Zora contract.
var ZoraMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_erc20TransferHelper\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_erc721TransferHelper\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_royaltyEngine\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_protocolFeeSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_wethAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sellerFundsRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"askCurrency\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"findersFeeBps\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"askPrice\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structAsksV1_1.Ask\",\"name\":\"ask\",\"type\":\"tuple\"}],\"name\":\"AskCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sellerFundsRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"askCurrency\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"findersFeeBps\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"askPrice\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structAsksV1_1.Ask\",\"name\":\"ask\",\"type\":\"tuple\"}],\"name\":\"AskCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"finder\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sellerFundsRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"askCurrency\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"findersFeeBps\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"askPrice\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structAsksV1_1.Ask\",\"name\":\"ask\",\"type\":\"tuple\"}],\"name\":\"AskFilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sellerFundsRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"askCurrency\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"findersFeeBps\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"askPrice\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structAsksV1_1.Ask\",\"name\":\"ask\",\"type\":\"tuple\"}],\"name\":\"AskPriceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"userA\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"userB\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structUniversalExchangeEventV1.ExchangeDetails\",\"name\":\"a\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structUniversalExchangeEventV1.ExchangeDetails\",\"name\":\"b\",\"type\":\"tuple\"}],\"name\":\"ExchangeExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RoyaltyPayout\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_payoutCurrency\",\"type\":\"address\"}],\"name\":\"_handleRoyaltyEnginePayout\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"askForNFT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sellerFundsRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"askCurrency\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"findersFeeBps\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"askPrice\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"cancelAsk\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_askPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_askCurrency\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_sellerFundsRecipient\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_findersFeeBps\",\"type\":\"uint16\"}],\"name\":\"createAsk\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"erc20TransferHelper\",\"outputs\":[{\"internalType\":\"contractERC20TransferHelper\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"erc721TransferHelper\",\"outputs\":[{\"internalType\":\"contractERC721TransferHelper\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_fillCurrency\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_fillAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_finder\",\"type\":\"address\"}],\"name\":\"fillAsk\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registrar\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_askPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_askCurrency\",\"type\":\"address\"}],\"name\":\"setAskPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_royaltyEngine\",\"type\":\"address\"}],\"name\":\"setRoyaltyEngineAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ZoraABI is the input ABI used to generate the binding from.
// Deprecated: Use ZoraMetaData.ABI instead.
var ZoraABI = ZoraMetaData.ABI

// Zora is an auto generated Go binding around an Ethereum contract.
type Zora struct {
	ZoraCaller     // Read-only binding to the contract
	ZoraTransactor // Write-only binding to the contract
	ZoraFilterer   // Log filterer for contract events
}

// ZoraCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZoraCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZoraTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZoraTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZoraFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZoraFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZoraSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZoraSession struct {
	Contract     *Zora             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZoraCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZoraCallerSession struct {
	Contract *ZoraCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ZoraTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZoraTransactorSession struct {
	Contract     *ZoraTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZoraRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZoraRaw struct {
	Contract *Zora // Generic contract binding to access the raw methods on
}

// ZoraCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZoraCallerRaw struct {
	Contract *ZoraCaller // Generic read-only contract binding to access the raw methods on
}

// ZoraTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZoraTransactorRaw struct {
	Contract *ZoraTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZora creates a new instance of Zora, bound to a specific deployed contract.
func NewZora(address common.Address, backend bind.ContractBackend) (*Zora, error) {
	contract, err := bindZora(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Zora{ZoraCaller: ZoraCaller{contract: contract}, ZoraTransactor: ZoraTransactor{contract: contract}, ZoraFilterer: ZoraFilterer{contract: contract}}, nil
}

// NewZoraCaller creates a new read-only instance of Zora, bound to a specific deployed contract.
func NewZoraCaller(address common.Address, caller bind.ContractCaller) (*ZoraCaller, error) {
	contract, err := bindZora(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZoraCaller{contract: contract}, nil
}

// NewZoraTransactor creates a new write-only instance of Zora, bound to a specific deployed contract.
func NewZoraTransactor(address common.Address, transactor bind.ContractTransactor) (*ZoraTransactor, error) {
	contract, err := bindZora(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZoraTransactor{contract: contract}, nil
}

// NewZoraFilterer creates a new log filterer instance of Zora, bound to a specific deployed contract.
func NewZoraFilterer(address common.Address, filterer bind.ContractFilterer) (*ZoraFilterer, error) {
	contract, err := bindZora(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZoraFilterer{contract: contract}, nil
}

// bindZora binds a generic wrapper to an already deployed contract.
func bindZora(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZoraABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Zora *ZoraRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Zora.Contract.ZoraCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Zora *ZoraRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Zora.Contract.ZoraTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Zora *ZoraRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Zora.Contract.ZoraTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Zora *ZoraCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Zora.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Zora *ZoraTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Zora.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Zora *ZoraTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Zora.Contract.contract.Transact(opts, method, params...)
}

// AskForNFT is a free data retrieval call binding the contract method 0x418f0656.
//
// Solidity: function askForNFT(address , uint256 ) view returns(address seller, address sellerFundsRecipient, address askCurrency, uint16 findersFeeBps, uint256 askPrice)
func (_Zora *ZoraCaller) AskForNFT(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Seller               common.Address
	SellerFundsRecipient common.Address
	AskCurrency          common.Address
	FindersFeeBps        uint16
	AskPrice             *big.Int
}, error) {
	var out []interface{}
	err := _Zora.contract.Call(opts, &out, "askForNFT", arg0, arg1)

	outstruct := new(struct {
		Seller               common.Address
		SellerFundsRecipient common.Address
		AskCurrency          common.Address
		FindersFeeBps        uint16
		AskPrice             *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Seller = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.SellerFundsRecipient = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.AskCurrency = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.FindersFeeBps = *abi.ConvertType(out[3], new(uint16)).(*uint16)
	outstruct.AskPrice = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// AskForNFT is a free data retrieval call binding the contract method 0x418f0656.
//
// Solidity: function askForNFT(address , uint256 ) view returns(address seller, address sellerFundsRecipient, address askCurrency, uint16 findersFeeBps, uint256 askPrice)
func (_Zora *ZoraSession) AskForNFT(arg0 common.Address, arg1 *big.Int) (struct {
	Seller               common.Address
	SellerFundsRecipient common.Address
	AskCurrency          common.Address
	FindersFeeBps        uint16
	AskPrice             *big.Int
}, error) {
	return _Zora.Contract.AskForNFT(&_Zora.CallOpts, arg0, arg1)
}

// AskForNFT is a free data retrieval call binding the contract method 0x418f0656.
//
// Solidity: function askForNFT(address , uint256 ) view returns(address seller, address sellerFundsRecipient, address askCurrency, uint16 findersFeeBps, uint256 askPrice)
func (_Zora *ZoraCallerSession) AskForNFT(arg0 common.Address, arg1 *big.Int) (struct {
	Seller               common.Address
	SellerFundsRecipient common.Address
	AskCurrency          common.Address
	FindersFeeBps        uint16
	AskPrice             *big.Int
}, error) {
	return _Zora.Contract.AskForNFT(&_Zora.CallOpts, arg0, arg1)
}

// Erc20TransferHelper is a free data retrieval call binding the contract method 0x8f9d3251.
//
// Solidity: function erc20TransferHelper() view returns(address)
func (_Zora *ZoraCaller) Erc20TransferHelper(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Zora.contract.Call(opts, &out, "erc20TransferHelper")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Erc20TransferHelper is a free data retrieval call binding the contract method 0x8f9d3251.
//
// Solidity: function erc20TransferHelper() view returns(address)
func (_Zora *ZoraSession) Erc20TransferHelper() (common.Address, error) {
	return _Zora.Contract.Erc20TransferHelper(&_Zora.CallOpts)
}

// Erc20TransferHelper is a free data retrieval call binding the contract method 0x8f9d3251.
//
// Solidity: function erc20TransferHelper() view returns(address)
func (_Zora *ZoraCallerSession) Erc20TransferHelper() (common.Address, error) {
	return _Zora.Contract.Erc20TransferHelper(&_Zora.CallOpts)
}

// Erc721TransferHelper is a free data retrieval call binding the contract method 0xf7cd1d9b.
//
// Solidity: function erc721TransferHelper() view returns(address)
func (_Zora *ZoraCaller) Erc721TransferHelper(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Zora.contract.Call(opts, &out, "erc721TransferHelper")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Erc721TransferHelper is a free data retrieval call binding the contract method 0xf7cd1d9b.
//
// Solidity: function erc721TransferHelper() view returns(address)
func (_Zora *ZoraSession) Erc721TransferHelper() (common.Address, error) {
	return _Zora.Contract.Erc721TransferHelper(&_Zora.CallOpts)
}

// Erc721TransferHelper is a free data retrieval call binding the contract method 0xf7cd1d9b.
//
// Solidity: function erc721TransferHelper() view returns(address)
func (_Zora *ZoraCallerSession) Erc721TransferHelper() (common.Address, error) {
	return _Zora.Contract.Erc721TransferHelper(&_Zora.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Zora *ZoraCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Zora.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Zora *ZoraSession) Name() (string, error) {
	return _Zora.Contract.Name(&_Zora.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Zora *ZoraCallerSession) Name() (string, error) {
	return _Zora.Contract.Name(&_Zora.CallOpts)
}

// Registrar is a free data retrieval call binding the contract method 0x2b20e397.
//
// Solidity: function registrar() view returns(address)
func (_Zora *ZoraCaller) Registrar(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Zora.contract.Call(opts, &out, "registrar")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registrar is a free data retrieval call binding the contract method 0x2b20e397.
//
// Solidity: function registrar() view returns(address)
func (_Zora *ZoraSession) Registrar() (common.Address, error) {
	return _Zora.Contract.Registrar(&_Zora.CallOpts)
}

// Registrar is a free data retrieval call binding the contract method 0x2b20e397.
//
// Solidity: function registrar() view returns(address)
func (_Zora *ZoraCallerSession) Registrar() (common.Address, error) {
	return _Zora.Contract.Registrar(&_Zora.CallOpts)
}

// HandleRoyaltyEnginePayout is a paid mutator transaction binding the contract method 0x9128c22c.
//
// Solidity: function _handleRoyaltyEnginePayout(address _tokenContract, uint256 _tokenId, uint256 _amount, address _payoutCurrency) payable returns(uint256)
func (_Zora *ZoraTransactor) HandleRoyaltyEnginePayout(opts *bind.TransactOpts, _tokenContract common.Address, _tokenId *big.Int, _amount *big.Int, _payoutCurrency common.Address) (*types.Transaction, error) {
	return _Zora.contract.Transact(opts, "_handleRoyaltyEnginePayout", _tokenContract, _tokenId, _amount, _payoutCurrency)
}

// HandleRoyaltyEnginePayout is a paid mutator transaction binding the contract method 0x9128c22c.
//
// Solidity: function _handleRoyaltyEnginePayout(address _tokenContract, uint256 _tokenId, uint256 _amount, address _payoutCurrency) payable returns(uint256)
func (_Zora *ZoraSession) HandleRoyaltyEnginePayout(_tokenContract common.Address, _tokenId *big.Int, _amount *big.Int, _payoutCurrency common.Address) (*types.Transaction, error) {
	return _Zora.Contract.HandleRoyaltyEnginePayout(&_Zora.TransactOpts, _tokenContract, _tokenId, _amount, _payoutCurrency)
}

// HandleRoyaltyEnginePayout is a paid mutator transaction binding the contract method 0x9128c22c.
//
// Solidity: function _handleRoyaltyEnginePayout(address _tokenContract, uint256 _tokenId, uint256 _amount, address _payoutCurrency) payable returns(uint256)
func (_Zora *ZoraTransactorSession) HandleRoyaltyEnginePayout(_tokenContract common.Address, _tokenId *big.Int, _amount *big.Int, _payoutCurrency common.Address) (*types.Transaction, error) {
	return _Zora.Contract.HandleRoyaltyEnginePayout(&_Zora.TransactOpts, _tokenContract, _tokenId, _amount, _payoutCurrency)
}

// CancelAsk is a paid mutator transaction binding the contract method 0x40b80746.
//
// Solidity: function cancelAsk(address _tokenContract, uint256 _tokenId) returns()
func (_Zora *ZoraTransactor) CancelAsk(opts *bind.TransactOpts, _tokenContract common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Zora.contract.Transact(opts, "cancelAsk", _tokenContract, _tokenId)
}

// CancelAsk is a paid mutator transaction binding the contract method 0x40b80746.
//
// Solidity: function cancelAsk(address _tokenContract, uint256 _tokenId) returns()
func (_Zora *ZoraSession) CancelAsk(_tokenContract common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Zora.Contract.CancelAsk(&_Zora.TransactOpts, _tokenContract, _tokenId)
}

// CancelAsk is a paid mutator transaction binding the contract method 0x40b80746.
//
// Solidity: function cancelAsk(address _tokenContract, uint256 _tokenId) returns()
func (_Zora *ZoraTransactorSession) CancelAsk(_tokenContract common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Zora.Contract.CancelAsk(&_Zora.TransactOpts, _tokenContract, _tokenId)
}

// CreateAsk is a paid mutator transaction binding the contract method 0x9e847108.
//
// Solidity: function createAsk(address _tokenContract, uint256 _tokenId, uint256 _askPrice, address _askCurrency, address _sellerFundsRecipient, uint16 _findersFeeBps) returns()
func (_Zora *ZoraTransactor) CreateAsk(opts *bind.TransactOpts, _tokenContract common.Address, _tokenId *big.Int, _askPrice *big.Int, _askCurrency common.Address, _sellerFundsRecipient common.Address, _findersFeeBps uint16) (*types.Transaction, error) {
	return _Zora.contract.Transact(opts, "createAsk", _tokenContract, _tokenId, _askPrice, _askCurrency, _sellerFundsRecipient, _findersFeeBps)
}

// CreateAsk is a paid mutator transaction binding the contract method 0x9e847108.
//
// Solidity: function createAsk(address _tokenContract, uint256 _tokenId, uint256 _askPrice, address _askCurrency, address _sellerFundsRecipient, uint16 _findersFeeBps) returns()
func (_Zora *ZoraSession) CreateAsk(_tokenContract common.Address, _tokenId *big.Int, _askPrice *big.Int, _askCurrency common.Address, _sellerFundsRecipient common.Address, _findersFeeBps uint16) (*types.Transaction, error) {
	return _Zora.Contract.CreateAsk(&_Zora.TransactOpts, _tokenContract, _tokenId, _askPrice, _askCurrency, _sellerFundsRecipient, _findersFeeBps)
}

// CreateAsk is a paid mutator transaction binding the contract method 0x9e847108.
//
// Solidity: function createAsk(address _tokenContract, uint256 _tokenId, uint256 _askPrice, address _askCurrency, address _sellerFundsRecipient, uint16 _findersFeeBps) returns()
func (_Zora *ZoraTransactorSession) CreateAsk(_tokenContract common.Address, _tokenId *big.Int, _askPrice *big.Int, _askCurrency common.Address, _sellerFundsRecipient common.Address, _findersFeeBps uint16) (*types.Transaction, error) {
	return _Zora.Contract.CreateAsk(&_Zora.TransactOpts, _tokenContract, _tokenId, _askPrice, _askCurrency, _sellerFundsRecipient, _findersFeeBps)
}

// FillAsk is a paid mutator transaction binding the contract method 0x622dcbd7.
//
// Solidity: function fillAsk(address _tokenContract, uint256 _tokenId, address _fillCurrency, uint256 _fillAmount, address _finder) payable returns()
func (_Zora *ZoraTransactor) FillAsk(opts *bind.TransactOpts, _tokenContract common.Address, _tokenId *big.Int, _fillCurrency common.Address, _fillAmount *big.Int, _finder common.Address) (*types.Transaction, error) {
	return _Zora.contract.Transact(opts, "fillAsk", _tokenContract, _tokenId, _fillCurrency, _fillAmount, _finder)
}

// FillAsk is a paid mutator transaction binding the contract method 0x622dcbd7.
//
// Solidity: function fillAsk(address _tokenContract, uint256 _tokenId, address _fillCurrency, uint256 _fillAmount, address _finder) payable returns()
func (_Zora *ZoraSession) FillAsk(_tokenContract common.Address, _tokenId *big.Int, _fillCurrency common.Address, _fillAmount *big.Int, _finder common.Address) (*types.Transaction, error) {
	return _Zora.Contract.FillAsk(&_Zora.TransactOpts, _tokenContract, _tokenId, _fillCurrency, _fillAmount, _finder)
}

// FillAsk is a paid mutator transaction binding the contract method 0x622dcbd7.
//
// Solidity: function fillAsk(address _tokenContract, uint256 _tokenId, address _fillCurrency, uint256 _fillAmount, address _finder) payable returns()
func (_Zora *ZoraTransactorSession) FillAsk(_tokenContract common.Address, _tokenId *big.Int, _fillCurrency common.Address, _fillAmount *big.Int, _finder common.Address) (*types.Transaction, error) {
	return _Zora.Contract.FillAsk(&_Zora.TransactOpts, _tokenContract, _tokenId, _fillCurrency, _fillAmount, _finder)
}

// SetAskPrice is a paid mutator transaction binding the contract method 0xb2007533.
//
// Solidity: function setAskPrice(address _tokenContract, uint256 _tokenId, uint256 _askPrice, address _askCurrency) returns()
func (_Zora *ZoraTransactor) SetAskPrice(opts *bind.TransactOpts, _tokenContract common.Address, _tokenId *big.Int, _askPrice *big.Int, _askCurrency common.Address) (*types.Transaction, error) {
	return _Zora.contract.Transact(opts, "setAskPrice", _tokenContract, _tokenId, _askPrice, _askCurrency)
}

// SetAskPrice is a paid mutator transaction binding the contract method 0xb2007533.
//
// Solidity: function setAskPrice(address _tokenContract, uint256 _tokenId, uint256 _askPrice, address _askCurrency) returns()
func (_Zora *ZoraSession) SetAskPrice(_tokenContract common.Address, _tokenId *big.Int, _askPrice *big.Int, _askCurrency common.Address) (*types.Transaction, error) {
	return _Zora.Contract.SetAskPrice(&_Zora.TransactOpts, _tokenContract, _tokenId, _askPrice, _askCurrency)
}

// SetAskPrice is a paid mutator transaction binding the contract method 0xb2007533.
//
// Solidity: function setAskPrice(address _tokenContract, uint256 _tokenId, uint256 _askPrice, address _askCurrency) returns()
func (_Zora *ZoraTransactorSession) SetAskPrice(_tokenContract common.Address, _tokenId *big.Int, _askPrice *big.Int, _askCurrency common.Address) (*types.Transaction, error) {
	return _Zora.Contract.SetAskPrice(&_Zora.TransactOpts, _tokenContract, _tokenId, _askPrice, _askCurrency)
}

// SetRoyaltyEngineAddress is a paid mutator transaction binding the contract method 0xb249bb30.
//
// Solidity: function setRoyaltyEngineAddress(address _royaltyEngine) returns()
func (_Zora *ZoraTransactor) SetRoyaltyEngineAddress(opts *bind.TransactOpts, _royaltyEngine common.Address) (*types.Transaction, error) {
	return _Zora.contract.Transact(opts, "setRoyaltyEngineAddress", _royaltyEngine)
}

// SetRoyaltyEngineAddress is a paid mutator transaction binding the contract method 0xb249bb30.
//
// Solidity: function setRoyaltyEngineAddress(address _royaltyEngine) returns()
func (_Zora *ZoraSession) SetRoyaltyEngineAddress(_royaltyEngine common.Address) (*types.Transaction, error) {
	return _Zora.Contract.SetRoyaltyEngineAddress(&_Zora.TransactOpts, _royaltyEngine)
}

// SetRoyaltyEngineAddress is a paid mutator transaction binding the contract method 0xb249bb30.
//
// Solidity: function setRoyaltyEngineAddress(address _royaltyEngine) returns()
func (_Zora *ZoraTransactorSession) SetRoyaltyEngineAddress(_royaltyEngine common.Address) (*types.Transaction, error) {
	return _Zora.Contract.SetRoyaltyEngineAddress(&_Zora.TransactOpts, _royaltyEngine)
}

// ZoraAskCanceledIterator is returned from FilterAskCanceled and is used to iterate over the raw logs and unpacked data for AskCanceled events raised by the Zora contract.
type ZoraAskCanceledIterator struct {
	Event *ZoraAskCanceled // Event containing the contract specifics and raw log

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
func (it *ZoraAskCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZoraAskCanceled)
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
		it.Event = new(ZoraAskCanceled)
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
func (it *ZoraAskCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZoraAskCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZoraAskCanceled represents a AskCanceled event raised by the Zora contract.
type ZoraAskCanceled struct {
	TokenContract common.Address
	TokenId       *big.Int
	Ask           AsksV11Ask
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAskCanceled is a free log retrieval operation binding the contract event 0x871956abf85befb7c955eacd40fcabe7e01b1702d75764bf7f54bf481933fd35.
//
// Solidity: event AskCanceled(address indexed tokenContract, uint256 indexed tokenId, (address,address,address,uint16,uint256) ask)
func (_Zora *ZoraFilterer) FilterAskCanceled(opts *bind.FilterOpts, tokenContract []common.Address, tokenId []*big.Int) (*ZoraAskCanceledIterator, error) {

	var tokenContractRule []interface{}
	for _, tokenContractItem := range tokenContract {
		tokenContractRule = append(tokenContractRule, tokenContractItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Zora.contract.FilterLogs(opts, "AskCanceled", tokenContractRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ZoraAskCanceledIterator{contract: _Zora.contract, event: "AskCanceled", logs: logs, sub: sub}, nil
}

// WatchAskCanceled is a free log subscription operation binding the contract event 0x871956abf85befb7c955eacd40fcabe7e01b1702d75764bf7f54bf481933fd35.
//
// Solidity: event AskCanceled(address indexed tokenContract, uint256 indexed tokenId, (address,address,address,uint16,uint256) ask)
func (_Zora *ZoraFilterer) WatchAskCanceled(opts *bind.WatchOpts, sink chan<- *ZoraAskCanceled, tokenContract []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var tokenContractRule []interface{}
	for _, tokenContractItem := range tokenContract {
		tokenContractRule = append(tokenContractRule, tokenContractItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Zora.contract.WatchLogs(opts, "AskCanceled", tokenContractRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZoraAskCanceled)
				if err := _Zora.contract.UnpackLog(event, "AskCanceled", log); err != nil {
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

// ParseAskCanceled is a log parse operation binding the contract event 0x871956abf85befb7c955eacd40fcabe7e01b1702d75764bf7f54bf481933fd35.
//
// Solidity: event AskCanceled(address indexed tokenContract, uint256 indexed tokenId, (address,address,address,uint16,uint256) ask)
func (_Zora *ZoraFilterer) ParseAskCanceled(log types.Log) (*ZoraAskCanceled, error) {
	event := new(ZoraAskCanceled)
	if err := _Zora.contract.UnpackLog(event, "AskCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZoraAskCreatedIterator is returned from FilterAskCreated and is used to iterate over the raw logs and unpacked data for AskCreated events raised by the Zora contract.
type ZoraAskCreatedIterator struct {
	Event *ZoraAskCreated // Event containing the contract specifics and raw log

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
func (it *ZoraAskCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZoraAskCreated)
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
		it.Event = new(ZoraAskCreated)
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
func (it *ZoraAskCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZoraAskCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZoraAskCreated represents a AskCreated event raised by the Zora contract.
type ZoraAskCreated struct {
	TokenContract common.Address
	TokenId       *big.Int
	Ask           AsksV11Ask
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAskCreated is a free log retrieval operation binding the contract event 0x5b65b398e1d736436510f4da442eaec71466d2abee0816567088c892c4bcee70.
//
// Solidity: event AskCreated(address indexed tokenContract, uint256 indexed tokenId, (address,address,address,uint16,uint256) ask)
func (_Zora *ZoraFilterer) FilterAskCreated(opts *bind.FilterOpts, tokenContract []common.Address, tokenId []*big.Int) (*ZoraAskCreatedIterator, error) {

	var tokenContractRule []interface{}
	for _, tokenContractItem := range tokenContract {
		tokenContractRule = append(tokenContractRule, tokenContractItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Zora.contract.FilterLogs(opts, "AskCreated", tokenContractRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ZoraAskCreatedIterator{contract: _Zora.contract, event: "AskCreated", logs: logs, sub: sub}, nil
}

// WatchAskCreated is a free log subscription operation binding the contract event 0x5b65b398e1d736436510f4da442eaec71466d2abee0816567088c892c4bcee70.
//
// Solidity: event AskCreated(address indexed tokenContract, uint256 indexed tokenId, (address,address,address,uint16,uint256) ask)
func (_Zora *ZoraFilterer) WatchAskCreated(opts *bind.WatchOpts, sink chan<- *ZoraAskCreated, tokenContract []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var tokenContractRule []interface{}
	for _, tokenContractItem := range tokenContract {
		tokenContractRule = append(tokenContractRule, tokenContractItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Zora.contract.WatchLogs(opts, "AskCreated", tokenContractRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZoraAskCreated)
				if err := _Zora.contract.UnpackLog(event, "AskCreated", log); err != nil {
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

// ParseAskCreated is a log parse operation binding the contract event 0x5b65b398e1d736436510f4da442eaec71466d2abee0816567088c892c4bcee70.
//
// Solidity: event AskCreated(address indexed tokenContract, uint256 indexed tokenId, (address,address,address,uint16,uint256) ask)
func (_Zora *ZoraFilterer) ParseAskCreated(log types.Log) (*ZoraAskCreated, error) {
	event := new(ZoraAskCreated)
	if err := _Zora.contract.UnpackLog(event, "AskCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZoraAskFilledIterator is returned from FilterAskFilled and is used to iterate over the raw logs and unpacked data for AskFilled events raised by the Zora contract.
type ZoraAskFilledIterator struct {
	Event *ZoraAskFilled // Event containing the contract specifics and raw log

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
func (it *ZoraAskFilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZoraAskFilled)
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
		it.Event = new(ZoraAskFilled)
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
func (it *ZoraAskFilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZoraAskFilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZoraAskFilled represents a AskFilled event raised by the Zora contract.
type ZoraAskFilled struct {
	TokenContract common.Address
	TokenId       *big.Int
	Buyer         common.Address
	Finder        common.Address
	Ask           AsksV11Ask
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAskFilled is a free log retrieval operation binding the contract event 0x21a9d8e221211780696258a05c6225b1a24f428e2fd4d51708f1ab2be4224d39.
//
// Solidity: event AskFilled(address indexed tokenContract, uint256 indexed tokenId, address indexed buyer, address finder, (address,address,address,uint16,uint256) ask)
func (_Zora *ZoraFilterer) FilterAskFilled(opts *bind.FilterOpts, tokenContract []common.Address, tokenId []*big.Int, buyer []common.Address) (*ZoraAskFilledIterator, error) {

	var tokenContractRule []interface{}
	for _, tokenContractItem := range tokenContract {
		tokenContractRule = append(tokenContractRule, tokenContractItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Zora.contract.FilterLogs(opts, "AskFilled", tokenContractRule, tokenIdRule, buyerRule)
	if err != nil {
		return nil, err
	}
	return &ZoraAskFilledIterator{contract: _Zora.contract, event: "AskFilled", logs: logs, sub: sub}, nil
}

// WatchAskFilled is a free log subscription operation binding the contract event 0x21a9d8e221211780696258a05c6225b1a24f428e2fd4d51708f1ab2be4224d39.
//
// Solidity: event AskFilled(address indexed tokenContract, uint256 indexed tokenId, address indexed buyer, address finder, (address,address,address,uint16,uint256) ask)
func (_Zora *ZoraFilterer) WatchAskFilled(opts *bind.WatchOpts, sink chan<- *ZoraAskFilled, tokenContract []common.Address, tokenId []*big.Int, buyer []common.Address) (event.Subscription, error) {

	var tokenContractRule []interface{}
	for _, tokenContractItem := range tokenContract {
		tokenContractRule = append(tokenContractRule, tokenContractItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Zora.contract.WatchLogs(opts, "AskFilled", tokenContractRule, tokenIdRule, buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZoraAskFilled)
				if err := _Zora.contract.UnpackLog(event, "AskFilled", log); err != nil {
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

// ParseAskFilled is a log parse operation binding the contract event 0x21a9d8e221211780696258a05c6225b1a24f428e2fd4d51708f1ab2be4224d39.
//
// Solidity: event AskFilled(address indexed tokenContract, uint256 indexed tokenId, address indexed buyer, address finder, (address,address,address,uint16,uint256) ask)
func (_Zora *ZoraFilterer) ParseAskFilled(log types.Log) (*ZoraAskFilled, error) {
	event := new(ZoraAskFilled)
	if err := _Zora.contract.UnpackLog(event, "AskFilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZoraAskPriceUpdatedIterator is returned from FilterAskPriceUpdated and is used to iterate over the raw logs and unpacked data for AskPriceUpdated events raised by the Zora contract.
type ZoraAskPriceUpdatedIterator struct {
	Event *ZoraAskPriceUpdated // Event containing the contract specifics and raw log

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
func (it *ZoraAskPriceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZoraAskPriceUpdated)
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
		it.Event = new(ZoraAskPriceUpdated)
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
func (it *ZoraAskPriceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZoraAskPriceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZoraAskPriceUpdated represents a AskPriceUpdated event raised by the Zora contract.
type ZoraAskPriceUpdated struct {
	TokenContract common.Address
	TokenId       *big.Int
	Ask           AsksV11Ask
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAskPriceUpdated is a free log retrieval operation binding the contract event 0x1a24bcf5290feab70f35cfb4870c294ebf497e608d4262b0ec0debe045008140.
//
// Solidity: event AskPriceUpdated(address indexed tokenContract, uint256 indexed tokenId, (address,address,address,uint16,uint256) ask)
func (_Zora *ZoraFilterer) FilterAskPriceUpdated(opts *bind.FilterOpts, tokenContract []common.Address, tokenId []*big.Int) (*ZoraAskPriceUpdatedIterator, error) {

	var tokenContractRule []interface{}
	for _, tokenContractItem := range tokenContract {
		tokenContractRule = append(tokenContractRule, tokenContractItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Zora.contract.FilterLogs(opts, "AskPriceUpdated", tokenContractRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ZoraAskPriceUpdatedIterator{contract: _Zora.contract, event: "AskPriceUpdated", logs: logs, sub: sub}, nil
}

// WatchAskPriceUpdated is a free log subscription operation binding the contract event 0x1a24bcf5290feab70f35cfb4870c294ebf497e608d4262b0ec0debe045008140.
//
// Solidity: event AskPriceUpdated(address indexed tokenContract, uint256 indexed tokenId, (address,address,address,uint16,uint256) ask)
func (_Zora *ZoraFilterer) WatchAskPriceUpdated(opts *bind.WatchOpts, sink chan<- *ZoraAskPriceUpdated, tokenContract []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var tokenContractRule []interface{}
	for _, tokenContractItem := range tokenContract {
		tokenContractRule = append(tokenContractRule, tokenContractItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Zora.contract.WatchLogs(opts, "AskPriceUpdated", tokenContractRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZoraAskPriceUpdated)
				if err := _Zora.contract.UnpackLog(event, "AskPriceUpdated", log); err != nil {
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

// ParseAskPriceUpdated is a log parse operation binding the contract event 0x1a24bcf5290feab70f35cfb4870c294ebf497e608d4262b0ec0debe045008140.
//
// Solidity: event AskPriceUpdated(address indexed tokenContract, uint256 indexed tokenId, (address,address,address,uint16,uint256) ask)
func (_Zora *ZoraFilterer) ParseAskPriceUpdated(log types.Log) (*ZoraAskPriceUpdated, error) {
	event := new(ZoraAskPriceUpdated)
	if err := _Zora.contract.UnpackLog(event, "AskPriceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZoraExchangeExecutedIterator is returned from FilterExchangeExecuted and is used to iterate over the raw logs and unpacked data for ExchangeExecuted events raised by the Zora contract.
type ZoraExchangeExecutedIterator struct {
	Event *ZoraExchangeExecuted // Event containing the contract specifics and raw log

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
func (it *ZoraExchangeExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZoraExchangeExecuted)
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
		it.Event = new(ZoraExchangeExecuted)
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
func (it *ZoraExchangeExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZoraExchangeExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZoraExchangeExecuted represents a ExchangeExecuted event raised by the Zora contract.
type ZoraExchangeExecuted struct {
	UserA common.Address
	UserB common.Address
	A     UniversalExchangeEventV1ExchangeDetails
	B     UniversalExchangeEventV1ExchangeDetails
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterExchangeExecuted is a free log retrieval operation binding the contract event 0x1f432c9454edd444f55492be01e3fa82aa88bfa28e120a039be204253c10c36e.
//
// Solidity: event ExchangeExecuted(address indexed userA, address indexed userB, (address,uint256,uint256) a, (address,uint256,uint256) b)
func (_Zora *ZoraFilterer) FilterExchangeExecuted(opts *bind.FilterOpts, userA []common.Address, userB []common.Address) (*ZoraExchangeExecutedIterator, error) {

	var userARule []interface{}
	for _, userAItem := range userA {
		userARule = append(userARule, userAItem)
	}
	var userBRule []interface{}
	for _, userBItem := range userB {
		userBRule = append(userBRule, userBItem)
	}

	logs, sub, err := _Zora.contract.FilterLogs(opts, "ExchangeExecuted", userARule, userBRule)
	if err != nil {
		return nil, err
	}
	return &ZoraExchangeExecutedIterator{contract: _Zora.contract, event: "ExchangeExecuted", logs: logs, sub: sub}, nil
}

// WatchExchangeExecuted is a free log subscription operation binding the contract event 0x1f432c9454edd444f55492be01e3fa82aa88bfa28e120a039be204253c10c36e.
//
// Solidity: event ExchangeExecuted(address indexed userA, address indexed userB, (address,uint256,uint256) a, (address,uint256,uint256) b)
func (_Zora *ZoraFilterer) WatchExchangeExecuted(opts *bind.WatchOpts, sink chan<- *ZoraExchangeExecuted, userA []common.Address, userB []common.Address) (event.Subscription, error) {

	var userARule []interface{}
	for _, userAItem := range userA {
		userARule = append(userARule, userAItem)
	}
	var userBRule []interface{}
	for _, userBItem := range userB {
		userBRule = append(userBRule, userBItem)
	}

	logs, sub, err := _Zora.contract.WatchLogs(opts, "ExchangeExecuted", userARule, userBRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZoraExchangeExecuted)
				if err := _Zora.contract.UnpackLog(event, "ExchangeExecuted", log); err != nil {
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

// ParseExchangeExecuted is a log parse operation binding the contract event 0x1f432c9454edd444f55492be01e3fa82aa88bfa28e120a039be204253c10c36e.
//
// Solidity: event ExchangeExecuted(address indexed userA, address indexed userB, (address,uint256,uint256) a, (address,uint256,uint256) b)
func (_Zora *ZoraFilterer) ParseExchangeExecuted(log types.Log) (*ZoraExchangeExecuted, error) {
	event := new(ZoraExchangeExecuted)
	if err := _Zora.contract.UnpackLog(event, "ExchangeExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZoraRoyaltyPayoutIterator is returned from FilterRoyaltyPayout and is used to iterate over the raw logs and unpacked data for RoyaltyPayout events raised by the Zora contract.
type ZoraRoyaltyPayoutIterator struct {
	Event *ZoraRoyaltyPayout // Event containing the contract specifics and raw log

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
func (it *ZoraRoyaltyPayoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZoraRoyaltyPayout)
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
		it.Event = new(ZoraRoyaltyPayout)
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
func (it *ZoraRoyaltyPayoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZoraRoyaltyPayoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZoraRoyaltyPayout represents a RoyaltyPayout event raised by the Zora contract.
type ZoraRoyaltyPayout struct {
	TokenContract common.Address
	TokenId       *big.Int
	Recipient     common.Address
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRoyaltyPayout is a free log retrieval operation binding the contract event 0x866e6ef8682ddf5f1025e64dfdb45527077f7be70fa9ef680b7ffd8cf4ab9c50.
//
// Solidity: event RoyaltyPayout(address indexed tokenContract, uint256 indexed tokenId, address indexed recipient, uint256 amount)
func (_Zora *ZoraFilterer) FilterRoyaltyPayout(opts *bind.FilterOpts, tokenContract []common.Address, tokenId []*big.Int, recipient []common.Address) (*ZoraRoyaltyPayoutIterator, error) {

	var tokenContractRule []interface{}
	for _, tokenContractItem := range tokenContract {
		tokenContractRule = append(tokenContractRule, tokenContractItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Zora.contract.FilterLogs(opts, "RoyaltyPayout", tokenContractRule, tokenIdRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &ZoraRoyaltyPayoutIterator{contract: _Zora.contract, event: "RoyaltyPayout", logs: logs, sub: sub}, nil
}

// WatchRoyaltyPayout is a free log subscription operation binding the contract event 0x866e6ef8682ddf5f1025e64dfdb45527077f7be70fa9ef680b7ffd8cf4ab9c50.
//
// Solidity: event RoyaltyPayout(address indexed tokenContract, uint256 indexed tokenId, address indexed recipient, uint256 amount)
func (_Zora *ZoraFilterer) WatchRoyaltyPayout(opts *bind.WatchOpts, sink chan<- *ZoraRoyaltyPayout, tokenContract []common.Address, tokenId []*big.Int, recipient []common.Address) (event.Subscription, error) {

	var tokenContractRule []interface{}
	for _, tokenContractItem := range tokenContract {
		tokenContractRule = append(tokenContractRule, tokenContractItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Zora.contract.WatchLogs(opts, "RoyaltyPayout", tokenContractRule, tokenIdRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZoraRoyaltyPayout)
				if err := _Zora.contract.UnpackLog(event, "RoyaltyPayout", log); err != nil {
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

// ParseRoyaltyPayout is a log parse operation binding the contract event 0x866e6ef8682ddf5f1025e64dfdb45527077f7be70fa9ef680b7ffd8cf4ab9c50.
//
// Solidity: event RoyaltyPayout(address indexed tokenContract, uint256 indexed tokenId, address indexed recipient, uint256 amount)
func (_Zora *ZoraFilterer) ParseRoyaltyPayout(log types.Log) (*ZoraRoyaltyPayout, error) {
	event := new(ZoraRoyaltyPayout)
	if err := _Zora.contract.UnpackLog(event, "RoyaltyPayout", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
