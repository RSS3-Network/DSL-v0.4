// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// IPIXPIXInfo is an auto generated low-level Go binding around an user-defined struct.
type IPIXPIXInfo struct {
	PixId    *big.Int
	Category uint8
	Size     uint8
}

// MarketplaceMetaData contains all meta data concerning the Marketplace contract.
var MarketplaceMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"saleId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"Purchased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"PurchasedWithSignature\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"saleId\",\"type\":\"uint256\"}],\"name\":\"SaleCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"saleId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"SaleRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"saleId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPrice\",\"type\":\"uint256\"}],\"name\":\"SaleUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"treasury\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"burnFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"mode\",\"type\":\"bool\"}],\"name\":\"TreasuryUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"burnHolder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_saleId\",\"type\":\"uint256\"}],\"name\":\"cancelSale\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"expirationTimes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_pixt\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pix\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"landTreasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"treasury\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"burnFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastSaleId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"noncesForSale\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pixMerkleMinter\",\"outputs\":[{\"internalType\":\"contractIPIXMerkleMinter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pixNFT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pixToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pixtTreasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"treasury\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"burnFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_saleId\",\"type\":\"uint256\"}],\"name\":\"purchaseNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftToken\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"requestSale\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pixId\",\"type\":\"uint256\"},{\"internalType\":\"enumIPIX.PIXCategory\",\"name\":\"category\",\"type\":\"uint8\"},{\"internalType\":\"enumIPIX.PIXSize\",\"name\":\"size\",\"type\":\"uint8\"}],\"internalType\":\"structIPIX.PIXInfo[]\",\"name\":\"info\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"merkleRoot\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"merkleProofs\",\"type\":\"bytes32[][]\"}],\"name\":\"requestSaleWithHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"saleInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nftToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"nftToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"sellNFTWithSignature\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pixId\",\"type\":\"uint256\"},{\"internalType\":\"enumIPIX.PIXCategory\",\"name\":\"category\",\"type\":\"uint8\"},{\"internalType\":\"enumIPIX.PIXSize\",\"name\":\"size\",\"type\":\"uint8\"}],\"internalType\":\"structIPIX.PIXInfo\",\"name\":\"info\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"merkleProofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"sellNFTWithSignatureWithHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"saleId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"sellSaleWithSignature\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"}],\"name\":\"setBurnHolder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_pixMerkleMinter\",\"type\":\"address\"}],\"name\":\"setPixMerkleMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_treasury\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_burnFee\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_mode\",\"type\":\"bool\"}],\"name\":\"setTreasury\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_whitelist\",\"type\":\"bool\"}],\"name\":\"setWhitelistedNFTs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_saleId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"updateSale\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelistedNFTs\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// MarketplaceABI is the input ABI used to generate the binding from.
// Deprecated: Use MarketplaceMetaData.ABI instead.
var MarketplaceABI = MarketplaceMetaData.ABI

// Marketplace is an auto generated Go binding around an Ethereum contract.
type Marketplace struct {
	MarketplaceCaller     // Read-only binding to the contract
	MarketplaceTransactor // Write-only binding to the contract
	MarketplaceFilterer   // Log filterer for contract events
}

// MarketplaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type MarketplaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketplaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MarketplaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketplaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MarketplaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketplaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MarketplaceSession struct {
	Contract     *Marketplace      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarketplaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MarketplaceCallerSession struct {
	Contract *MarketplaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// MarketplaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MarketplaceTransactorSession struct {
	Contract     *MarketplaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MarketplaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type MarketplaceRaw struct {
	Contract *Marketplace // Generic contract binding to access the raw methods on
}

// MarketplaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MarketplaceCallerRaw struct {
	Contract *MarketplaceCaller // Generic read-only contract binding to access the raw methods on
}

// MarketplaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MarketplaceTransactorRaw struct {
	Contract *MarketplaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMarketplace creates a new instance of Marketplace, bound to a specific deployed contract.
func NewMarketplace(address common.Address, backend bind.ContractBackend) (*Marketplace, error) {
	contract, err := bindMarketplace(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Marketplace{MarketplaceCaller: MarketplaceCaller{contract: contract}, MarketplaceTransactor: MarketplaceTransactor{contract: contract}, MarketplaceFilterer: MarketplaceFilterer{contract: contract}}, nil
}

// NewMarketplaceCaller creates a new read-only instance of Marketplace, bound to a specific deployed contract.
func NewMarketplaceCaller(address common.Address, caller bind.ContractCaller) (*MarketplaceCaller, error) {
	contract, err := bindMarketplace(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MarketplaceCaller{contract: contract}, nil
}

// NewMarketplaceTransactor creates a new write-only instance of Marketplace, bound to a specific deployed contract.
func NewMarketplaceTransactor(address common.Address, transactor bind.ContractTransactor) (*MarketplaceTransactor, error) {
	contract, err := bindMarketplace(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MarketplaceTransactor{contract: contract}, nil
}

// NewMarketplaceFilterer creates a new log filterer instance of Marketplace, bound to a specific deployed contract.
func NewMarketplaceFilterer(address common.Address, filterer bind.ContractFilterer) (*MarketplaceFilterer, error) {
	contract, err := bindMarketplace(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MarketplaceFilterer{contract: contract}, nil
}

// bindMarketplace binds a generic wrapper to an already deployed contract.
func bindMarketplace(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MarketplaceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Marketplace *MarketplaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Marketplace.Contract.MarketplaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Marketplace *MarketplaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketplace.Contract.MarketplaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Marketplace *MarketplaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Marketplace.Contract.MarketplaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Marketplace *MarketplaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Marketplace.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Marketplace *MarketplaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketplace.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Marketplace *MarketplaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Marketplace.Contract.contract.Transact(opts, method, params...)
}

// BurnHolder is a free data retrieval call binding the contract method 0x0599243c.
//
// Solidity: function burnHolder() view returns(address)
func (_Marketplace *MarketplaceCaller) BurnHolder(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "burnHolder")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BurnHolder is a free data retrieval call binding the contract method 0x0599243c.
//
// Solidity: function burnHolder() view returns(address)
func (_Marketplace *MarketplaceSession) BurnHolder() (common.Address, error) {
	return _Marketplace.Contract.BurnHolder(&_Marketplace.CallOpts)
}

// BurnHolder is a free data retrieval call binding the contract method 0x0599243c.
//
// Solidity: function burnHolder() view returns(address)
func (_Marketplace *MarketplaceCallerSession) BurnHolder() (common.Address, error) {
	return _Marketplace.Contract.BurnHolder(&_Marketplace.CallOpts)
}

// ExpirationTimes is a free data retrieval call binding the contract method 0xe5e7a7b3.
//
// Solidity: function expirationTimes(uint256 ) view returns(uint256)
func (_Marketplace *MarketplaceCaller) ExpirationTimes(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "expirationTimes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExpirationTimes is a free data retrieval call binding the contract method 0xe5e7a7b3.
//
// Solidity: function expirationTimes(uint256 ) view returns(uint256)
func (_Marketplace *MarketplaceSession) ExpirationTimes(arg0 *big.Int) (*big.Int, error) {
	return _Marketplace.Contract.ExpirationTimes(&_Marketplace.CallOpts, arg0)
}

// ExpirationTimes is a free data retrieval call binding the contract method 0xe5e7a7b3.
//
// Solidity: function expirationTimes(uint256 ) view returns(uint256)
func (_Marketplace *MarketplaceCallerSession) ExpirationTimes(arg0 *big.Int) (*big.Int, error) {
	return _Marketplace.Contract.ExpirationTimes(&_Marketplace.CallOpts, arg0)
}

// LandTreasury is a free data retrieval call binding the contract method 0x32f4ad3d.
//
// Solidity: function landTreasury() view returns(address treasury, uint256 fee, uint256 burnFee)
func (_Marketplace *MarketplaceCaller) LandTreasury(opts *bind.CallOpts) (struct {
	Treasury common.Address
	Fee      *big.Int
	BurnFee  *big.Int
}, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "landTreasury")

	outstruct := new(struct {
		Treasury common.Address
		Fee      *big.Int
		BurnFee  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Treasury = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Fee = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BurnFee = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// LandTreasury is a free data retrieval call binding the contract method 0x32f4ad3d.
//
// Solidity: function landTreasury() view returns(address treasury, uint256 fee, uint256 burnFee)
func (_Marketplace *MarketplaceSession) LandTreasury() (struct {
	Treasury common.Address
	Fee      *big.Int
	BurnFee  *big.Int
}, error) {
	return _Marketplace.Contract.LandTreasury(&_Marketplace.CallOpts)
}

// LandTreasury is a free data retrieval call binding the contract method 0x32f4ad3d.
//
// Solidity: function landTreasury() view returns(address treasury, uint256 fee, uint256 burnFee)
func (_Marketplace *MarketplaceCallerSession) LandTreasury() (struct {
	Treasury common.Address
	Fee      *big.Int
	BurnFee  *big.Int
}, error) {
	return _Marketplace.Contract.LandTreasury(&_Marketplace.CallOpts)
}

// LastSaleId is a free data retrieval call binding the contract method 0x2a2f9721.
//
// Solidity: function lastSaleId() view returns(uint256)
func (_Marketplace *MarketplaceCaller) LastSaleId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "lastSaleId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastSaleId is a free data retrieval call binding the contract method 0x2a2f9721.
//
// Solidity: function lastSaleId() view returns(uint256)
func (_Marketplace *MarketplaceSession) LastSaleId() (*big.Int, error) {
	return _Marketplace.Contract.LastSaleId(&_Marketplace.CallOpts)
}

// LastSaleId is a free data retrieval call binding the contract method 0x2a2f9721.
//
// Solidity: function lastSaleId() view returns(uint256)
func (_Marketplace *MarketplaceCallerSession) LastSaleId() (*big.Int, error) {
	return _Marketplace.Contract.LastSaleId(&_Marketplace.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0xfbf58dbb.
//
// Solidity: function nonces(address , address , uint256 ) view returns(uint256)
func (_Marketplace *MarketplaceCaller) Nonces(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "nonces", arg0, arg1, arg2)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0xfbf58dbb.
//
// Solidity: function nonces(address , address , uint256 ) view returns(uint256)
func (_Marketplace *MarketplaceSession) Nonces(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*big.Int, error) {
	return _Marketplace.Contract.Nonces(&_Marketplace.CallOpts, arg0, arg1, arg2)
}

// Nonces is a free data retrieval call binding the contract method 0xfbf58dbb.
//
// Solidity: function nonces(address , address , uint256 ) view returns(uint256)
func (_Marketplace *MarketplaceCallerSession) Nonces(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*big.Int, error) {
	return _Marketplace.Contract.Nonces(&_Marketplace.CallOpts, arg0, arg1, arg2)
}

// NoncesForSale is a free data retrieval call binding the contract method 0x44e3a27d.
//
// Solidity: function noncesForSale(address , uint256 ) view returns(uint256)
func (_Marketplace *MarketplaceCaller) NoncesForSale(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "noncesForSale", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NoncesForSale is a free data retrieval call binding the contract method 0x44e3a27d.
//
// Solidity: function noncesForSale(address , uint256 ) view returns(uint256)
func (_Marketplace *MarketplaceSession) NoncesForSale(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Marketplace.Contract.NoncesForSale(&_Marketplace.CallOpts, arg0, arg1)
}

// NoncesForSale is a free data retrieval call binding the contract method 0x44e3a27d.
//
// Solidity: function noncesForSale(address , uint256 ) view returns(uint256)
func (_Marketplace *MarketplaceCallerSession) NoncesForSale(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Marketplace.Contract.NoncesForSale(&_Marketplace.CallOpts, arg0, arg1)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Marketplace *MarketplaceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Marketplace *MarketplaceSession) Owner() (common.Address, error) {
	return _Marketplace.Contract.Owner(&_Marketplace.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Marketplace *MarketplaceCallerSession) Owner() (common.Address, error) {
	return _Marketplace.Contract.Owner(&_Marketplace.CallOpts)
}

// PixMerkleMinter is a free data retrieval call binding the contract method 0x376d0154.
//
// Solidity: function pixMerkleMinter() view returns(address)
func (_Marketplace *MarketplaceCaller) PixMerkleMinter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "pixMerkleMinter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PixMerkleMinter is a free data retrieval call binding the contract method 0x376d0154.
//
// Solidity: function pixMerkleMinter() view returns(address)
func (_Marketplace *MarketplaceSession) PixMerkleMinter() (common.Address, error) {
	return _Marketplace.Contract.PixMerkleMinter(&_Marketplace.CallOpts)
}

// PixMerkleMinter is a free data retrieval call binding the contract method 0x376d0154.
//
// Solidity: function pixMerkleMinter() view returns(address)
func (_Marketplace *MarketplaceCallerSession) PixMerkleMinter() (common.Address, error) {
	return _Marketplace.Contract.PixMerkleMinter(&_Marketplace.CallOpts)
}

// PixNFT is a free data retrieval call binding the contract method 0x6bd488e2.
//
// Solidity: function pixNFT() view returns(address)
func (_Marketplace *MarketplaceCaller) PixNFT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "pixNFT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PixNFT is a free data retrieval call binding the contract method 0x6bd488e2.
//
// Solidity: function pixNFT() view returns(address)
func (_Marketplace *MarketplaceSession) PixNFT() (common.Address, error) {
	return _Marketplace.Contract.PixNFT(&_Marketplace.CallOpts)
}

// PixNFT is a free data retrieval call binding the contract method 0x6bd488e2.
//
// Solidity: function pixNFT() view returns(address)
func (_Marketplace *MarketplaceCallerSession) PixNFT() (common.Address, error) {
	return _Marketplace.Contract.PixNFT(&_Marketplace.CallOpts)
}

// PixToken is a free data retrieval call binding the contract method 0x36325eff.
//
// Solidity: function pixToken() view returns(address)
func (_Marketplace *MarketplaceCaller) PixToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "pixToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PixToken is a free data retrieval call binding the contract method 0x36325eff.
//
// Solidity: function pixToken() view returns(address)
func (_Marketplace *MarketplaceSession) PixToken() (common.Address, error) {
	return _Marketplace.Contract.PixToken(&_Marketplace.CallOpts)
}

// PixToken is a free data retrieval call binding the contract method 0x36325eff.
//
// Solidity: function pixToken() view returns(address)
func (_Marketplace *MarketplaceCallerSession) PixToken() (common.Address, error) {
	return _Marketplace.Contract.PixToken(&_Marketplace.CallOpts)
}

// PixtTreasury is a free data retrieval call binding the contract method 0xf1d92982.
//
// Solidity: function pixtTreasury() view returns(address treasury, uint256 fee, uint256 burnFee)
func (_Marketplace *MarketplaceCaller) PixtTreasury(opts *bind.CallOpts) (struct {
	Treasury common.Address
	Fee      *big.Int
	BurnFee  *big.Int
}, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "pixtTreasury")

	outstruct := new(struct {
		Treasury common.Address
		Fee      *big.Int
		BurnFee  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Treasury = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Fee = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BurnFee = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PixtTreasury is a free data retrieval call binding the contract method 0xf1d92982.
//
// Solidity: function pixtTreasury() view returns(address treasury, uint256 fee, uint256 burnFee)
func (_Marketplace *MarketplaceSession) PixtTreasury() (struct {
	Treasury common.Address
	Fee      *big.Int
	BurnFee  *big.Int
}, error) {
	return _Marketplace.Contract.PixtTreasury(&_Marketplace.CallOpts)
}

// PixtTreasury is a free data retrieval call binding the contract method 0xf1d92982.
//
// Solidity: function pixtTreasury() view returns(address treasury, uint256 fee, uint256 burnFee)
func (_Marketplace *MarketplaceCallerSession) PixtTreasury() (struct {
	Treasury common.Address
	Fee      *big.Int
	BurnFee  *big.Int
}, error) {
	return _Marketplace.Contract.PixtTreasury(&_Marketplace.CallOpts)
}

// SaleInfo is a free data retrieval call binding the contract method 0x02eb4b88.
//
// Solidity: function saleInfo(uint256 ) view returns(address seller, address nftToken, uint256 price)
func (_Marketplace *MarketplaceCaller) SaleInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Seller   common.Address
	NftToken common.Address
	Price    *big.Int
}, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "saleInfo", arg0)

	outstruct := new(struct {
		Seller   common.Address
		NftToken common.Address
		Price    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Seller = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.NftToken = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Price = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// SaleInfo is a free data retrieval call binding the contract method 0x02eb4b88.
//
// Solidity: function saleInfo(uint256 ) view returns(address seller, address nftToken, uint256 price)
func (_Marketplace *MarketplaceSession) SaleInfo(arg0 *big.Int) (struct {
	Seller   common.Address
	NftToken common.Address
	Price    *big.Int
}, error) {
	return _Marketplace.Contract.SaleInfo(&_Marketplace.CallOpts, arg0)
}

// SaleInfo is a free data retrieval call binding the contract method 0x02eb4b88.
//
// Solidity: function saleInfo(uint256 ) view returns(address seller, address nftToken, uint256 price)
func (_Marketplace *MarketplaceCallerSession) SaleInfo(arg0 *big.Int) (struct {
	Seller   common.Address
	NftToken common.Address
	Price    *big.Int
}, error) {
	return _Marketplace.Contract.SaleInfo(&_Marketplace.CallOpts, arg0)
}

// WhitelistedNFTs is a free data retrieval call binding the contract method 0x3f0c188f.
//
// Solidity: function whitelistedNFTs(address ) view returns(bool)
func (_Marketplace *MarketplaceCaller) WhitelistedNFTs(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "whitelistedNFTs", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WhitelistedNFTs is a free data retrieval call binding the contract method 0x3f0c188f.
//
// Solidity: function whitelistedNFTs(address ) view returns(bool)
func (_Marketplace *MarketplaceSession) WhitelistedNFTs(arg0 common.Address) (bool, error) {
	return _Marketplace.Contract.WhitelistedNFTs(&_Marketplace.CallOpts, arg0)
}

// WhitelistedNFTs is a free data retrieval call binding the contract method 0x3f0c188f.
//
// Solidity: function whitelistedNFTs(address ) view returns(bool)
func (_Marketplace *MarketplaceCallerSession) WhitelistedNFTs(arg0 common.Address) (bool, error) {
	return _Marketplace.Contract.WhitelistedNFTs(&_Marketplace.CallOpts, arg0)
}

// CancelSale is a paid mutator transaction binding the contract method 0xbd94b005.
//
// Solidity: function cancelSale(uint256 _saleId) returns()
func (_Marketplace *MarketplaceTransactor) CancelSale(opts *bind.TransactOpts, _saleId *big.Int) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "cancelSale", _saleId)
}

// CancelSale is a paid mutator transaction binding the contract method 0xbd94b005.
//
// Solidity: function cancelSale(uint256 _saleId) returns()
func (_Marketplace *MarketplaceSession) CancelSale(_saleId *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.CancelSale(&_Marketplace.TransactOpts, _saleId)
}

// CancelSale is a paid mutator transaction binding the contract method 0xbd94b005.
//
// Solidity: function cancelSale(uint256 _saleId) returns()
func (_Marketplace *MarketplaceTransactorSession) CancelSale(_saleId *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.CancelSale(&_Marketplace.TransactOpts, _saleId)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _pixt, address _pix) returns()
func (_Marketplace *MarketplaceTransactor) Initialize(opts *bind.TransactOpts, _pixt common.Address, _pix common.Address) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "initialize", _pixt, _pix)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _pixt, address _pix) returns()
func (_Marketplace *MarketplaceSession) Initialize(_pixt common.Address, _pix common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.Initialize(&_Marketplace.TransactOpts, _pixt, _pix)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _pixt, address _pix) returns()
func (_Marketplace *MarketplaceTransactorSession) Initialize(_pixt common.Address, _pix common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.Initialize(&_Marketplace.TransactOpts, _pixt, _pix)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Marketplace *MarketplaceTransactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "onERC721Received", arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Marketplace *MarketplaceSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Marketplace.Contract.OnERC721Received(&_Marketplace.TransactOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Marketplace *MarketplaceTransactorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Marketplace.Contract.OnERC721Received(&_Marketplace.TransactOpts, arg0, arg1, arg2, arg3)
}

// PurchaseNFT is a paid mutator transaction binding the contract method 0x150bde03.
//
// Solidity: function purchaseNFT(uint256 _saleId) returns()
func (_Marketplace *MarketplaceTransactor) PurchaseNFT(opts *bind.TransactOpts, _saleId *big.Int) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "purchaseNFT", _saleId)
}

// PurchaseNFT is a paid mutator transaction binding the contract method 0x150bde03.
//
// Solidity: function purchaseNFT(uint256 _saleId) returns()
func (_Marketplace *MarketplaceSession) PurchaseNFT(_saleId *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.PurchaseNFT(&_Marketplace.TransactOpts, _saleId)
}

// PurchaseNFT is a paid mutator transaction binding the contract method 0x150bde03.
//
// Solidity: function purchaseNFT(uint256 _saleId) returns()
func (_Marketplace *MarketplaceTransactorSession) PurchaseNFT(_saleId *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.PurchaseNFT(&_Marketplace.TransactOpts, _saleId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Marketplace *MarketplaceTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Marketplace *MarketplaceSession) RenounceOwnership() (*types.Transaction, error) {
	return _Marketplace.Contract.RenounceOwnership(&_Marketplace.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Marketplace *MarketplaceTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Marketplace.Contract.RenounceOwnership(&_Marketplace.TransactOpts)
}

// RequestSale is a paid mutator transaction binding the contract method 0x5895de93.
//
// Solidity: function requestSale(address _nftToken, uint256[] _tokenIds, uint256 _price) returns()
func (_Marketplace *MarketplaceTransactor) RequestSale(opts *bind.TransactOpts, _nftToken common.Address, _tokenIds []*big.Int, _price *big.Int) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "requestSale", _nftToken, _tokenIds, _price)
}

// RequestSale is a paid mutator transaction binding the contract method 0x5895de93.
//
// Solidity: function requestSale(address _nftToken, uint256[] _tokenIds, uint256 _price) returns()
func (_Marketplace *MarketplaceSession) RequestSale(_nftToken common.Address, _tokenIds []*big.Int, _price *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.RequestSale(&_Marketplace.TransactOpts, _nftToken, _tokenIds, _price)
}

// RequestSale is a paid mutator transaction binding the contract method 0x5895de93.
//
// Solidity: function requestSale(address _nftToken, uint256[] _tokenIds, uint256 _price) returns()
func (_Marketplace *MarketplaceTransactorSession) RequestSale(_nftToken common.Address, _tokenIds []*big.Int, _price *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.RequestSale(&_Marketplace.TransactOpts, _nftToken, _tokenIds, _price)
}

// RequestSaleWithHash is a paid mutator transaction binding the contract method 0x973dcb09.
//
// Solidity: function requestSaleWithHash(uint256[] _tokenIds, uint256 _price, (uint256,uint8,uint8)[] info, bytes32[] merkleRoot, bytes32[][] merkleProofs) returns()
func (_Marketplace *MarketplaceTransactor) RequestSaleWithHash(opts *bind.TransactOpts, _tokenIds []*big.Int, _price *big.Int, info []IPIXPIXInfo, merkleRoot [][32]byte, merkleProofs [][][32]byte) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "requestSaleWithHash", _tokenIds, _price, info, merkleRoot, merkleProofs)
}

// RequestSaleWithHash is a paid mutator transaction binding the contract method 0x973dcb09.
//
// Solidity: function requestSaleWithHash(uint256[] _tokenIds, uint256 _price, (uint256,uint8,uint8)[] info, bytes32[] merkleRoot, bytes32[][] merkleProofs) returns()
func (_Marketplace *MarketplaceSession) RequestSaleWithHash(_tokenIds []*big.Int, _price *big.Int, info []IPIXPIXInfo, merkleRoot [][32]byte, merkleProofs [][][32]byte) (*types.Transaction, error) {
	return _Marketplace.Contract.RequestSaleWithHash(&_Marketplace.TransactOpts, _tokenIds, _price, info, merkleRoot, merkleProofs)
}

// RequestSaleWithHash is a paid mutator transaction binding the contract method 0x973dcb09.
//
// Solidity: function requestSaleWithHash(uint256[] _tokenIds, uint256 _price, (uint256,uint8,uint8)[] info, bytes32[] merkleRoot, bytes32[][] merkleProofs) returns()
func (_Marketplace *MarketplaceTransactorSession) RequestSaleWithHash(_tokenIds []*big.Int, _price *big.Int, info []IPIXPIXInfo, merkleRoot [][32]byte, merkleProofs [][][32]byte) (*types.Transaction, error) {
	return _Marketplace.Contract.RequestSaleWithHash(&_Marketplace.TransactOpts, _tokenIds, _price, info, merkleRoot, merkleProofs)
}

// SellNFTWithSignature is a paid mutator transaction binding the contract method 0x540fdb37.
//
// Solidity: function sellNFTWithSignature(address buyer, uint256 price, address nftToken, uint256 tokenId, uint8 v, bytes32 r, bytes32 s) returns()
func (_Marketplace *MarketplaceTransactor) SellNFTWithSignature(opts *bind.TransactOpts, buyer common.Address, price *big.Int, nftToken common.Address, tokenId *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "sellNFTWithSignature", buyer, price, nftToken, tokenId, v, r, s)
}

// SellNFTWithSignature is a paid mutator transaction binding the contract method 0x540fdb37.
//
// Solidity: function sellNFTWithSignature(address buyer, uint256 price, address nftToken, uint256 tokenId, uint8 v, bytes32 r, bytes32 s) returns()
func (_Marketplace *MarketplaceSession) SellNFTWithSignature(buyer common.Address, price *big.Int, nftToken common.Address, tokenId *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Marketplace.Contract.SellNFTWithSignature(&_Marketplace.TransactOpts, buyer, price, nftToken, tokenId, v, r, s)
}

// SellNFTWithSignature is a paid mutator transaction binding the contract method 0x540fdb37.
//
// Solidity: function sellNFTWithSignature(address buyer, uint256 price, address nftToken, uint256 tokenId, uint8 v, bytes32 r, bytes32 s) returns()
func (_Marketplace *MarketplaceTransactorSession) SellNFTWithSignature(buyer common.Address, price *big.Int, nftToken common.Address, tokenId *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Marketplace.Contract.SellNFTWithSignature(&_Marketplace.TransactOpts, buyer, price, nftToken, tokenId, v, r, s)
}

// SellNFTWithSignatureWithHash is a paid mutator transaction binding the contract method 0xc4605394.
//
// Solidity: function sellNFTWithSignatureWithHash(address buyer, uint256 price, (uint256,uint8,uint8) info, bytes32 merkleRoot, bytes32[] merkleProofs, uint8 v, bytes32 r, bytes32 s) returns()
func (_Marketplace *MarketplaceTransactor) SellNFTWithSignatureWithHash(opts *bind.TransactOpts, buyer common.Address, price *big.Int, info IPIXPIXInfo, merkleRoot [32]byte, merkleProofs [][32]byte, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "sellNFTWithSignatureWithHash", buyer, price, info, merkleRoot, merkleProofs, v, r, s)
}

// SellNFTWithSignatureWithHash is a paid mutator transaction binding the contract method 0xc4605394.
//
// Solidity: function sellNFTWithSignatureWithHash(address buyer, uint256 price, (uint256,uint8,uint8) info, bytes32 merkleRoot, bytes32[] merkleProofs, uint8 v, bytes32 r, bytes32 s) returns()
func (_Marketplace *MarketplaceSession) SellNFTWithSignatureWithHash(buyer common.Address, price *big.Int, info IPIXPIXInfo, merkleRoot [32]byte, merkleProofs [][32]byte, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Marketplace.Contract.SellNFTWithSignatureWithHash(&_Marketplace.TransactOpts, buyer, price, info, merkleRoot, merkleProofs, v, r, s)
}

// SellNFTWithSignatureWithHash is a paid mutator transaction binding the contract method 0xc4605394.
//
// Solidity: function sellNFTWithSignatureWithHash(address buyer, uint256 price, (uint256,uint8,uint8) info, bytes32 merkleRoot, bytes32[] merkleProofs, uint8 v, bytes32 r, bytes32 s) returns()
func (_Marketplace *MarketplaceTransactorSession) SellNFTWithSignatureWithHash(buyer common.Address, price *big.Int, info IPIXPIXInfo, merkleRoot [32]byte, merkleProofs [][32]byte, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Marketplace.Contract.SellNFTWithSignatureWithHash(&_Marketplace.TransactOpts, buyer, price, info, merkleRoot, merkleProofs, v, r, s)
}

// SellSaleWithSignature is a paid mutator transaction binding the contract method 0x00d26c0c.
//
// Solidity: function sellSaleWithSignature(address buyer, uint256 price, uint256 saleId, uint8 v, bytes32 r, bytes32 s) returns()
func (_Marketplace *MarketplaceTransactor) SellSaleWithSignature(opts *bind.TransactOpts, buyer common.Address, price *big.Int, saleId *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "sellSaleWithSignature", buyer, price, saleId, v, r, s)
}

// SellSaleWithSignature is a paid mutator transaction binding the contract method 0x00d26c0c.
//
// Solidity: function sellSaleWithSignature(address buyer, uint256 price, uint256 saleId, uint8 v, bytes32 r, bytes32 s) returns()
func (_Marketplace *MarketplaceSession) SellSaleWithSignature(buyer common.Address, price *big.Int, saleId *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Marketplace.Contract.SellSaleWithSignature(&_Marketplace.TransactOpts, buyer, price, saleId, v, r, s)
}

// SellSaleWithSignature is a paid mutator transaction binding the contract method 0x00d26c0c.
//
// Solidity: function sellSaleWithSignature(address buyer, uint256 price, uint256 saleId, uint8 v, bytes32 r, bytes32 s) returns()
func (_Marketplace *MarketplaceTransactorSession) SellSaleWithSignature(buyer common.Address, price *big.Int, saleId *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Marketplace.Contract.SellSaleWithSignature(&_Marketplace.TransactOpts, buyer, price, saleId, v, r, s)
}

// SetBurnHolder is a paid mutator transaction binding the contract method 0x06f0bbe3.
//
// Solidity: function setBurnHolder(address holder) returns()
func (_Marketplace *MarketplaceTransactor) SetBurnHolder(opts *bind.TransactOpts, holder common.Address) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "setBurnHolder", holder)
}

// SetBurnHolder is a paid mutator transaction binding the contract method 0x06f0bbe3.
//
// Solidity: function setBurnHolder(address holder) returns()
func (_Marketplace *MarketplaceSession) SetBurnHolder(holder common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.SetBurnHolder(&_Marketplace.TransactOpts, holder)
}

// SetBurnHolder is a paid mutator transaction binding the contract method 0x06f0bbe3.
//
// Solidity: function setBurnHolder(address holder) returns()
func (_Marketplace *MarketplaceTransactorSession) SetBurnHolder(holder common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.SetBurnHolder(&_Marketplace.TransactOpts, holder)
}

// SetPixMerkleMinter is a paid mutator transaction binding the contract method 0x20c1d3c2.
//
// Solidity: function setPixMerkleMinter(address _pixMerkleMinter) returns()
func (_Marketplace *MarketplaceTransactor) SetPixMerkleMinter(opts *bind.TransactOpts, _pixMerkleMinter common.Address) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "setPixMerkleMinter", _pixMerkleMinter)
}

// SetPixMerkleMinter is a paid mutator transaction binding the contract method 0x20c1d3c2.
//
// Solidity: function setPixMerkleMinter(address _pixMerkleMinter) returns()
func (_Marketplace *MarketplaceSession) SetPixMerkleMinter(_pixMerkleMinter common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.SetPixMerkleMinter(&_Marketplace.TransactOpts, _pixMerkleMinter)
}

// SetPixMerkleMinter is a paid mutator transaction binding the contract method 0x20c1d3c2.
//
// Solidity: function setPixMerkleMinter(address _pixMerkleMinter) returns()
func (_Marketplace *MarketplaceTransactorSession) SetPixMerkleMinter(_pixMerkleMinter common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.SetPixMerkleMinter(&_Marketplace.TransactOpts, _pixMerkleMinter)
}

// SetTreasury is a paid mutator transaction binding the contract method 0x5e9d003b.
//
// Solidity: function setTreasury(address _treasury, uint256 _fee, uint256 _burnFee, bool _mode) returns()
func (_Marketplace *MarketplaceTransactor) SetTreasury(opts *bind.TransactOpts, _treasury common.Address, _fee *big.Int, _burnFee *big.Int, _mode bool) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "setTreasury", _treasury, _fee, _burnFee, _mode)
}

// SetTreasury is a paid mutator transaction binding the contract method 0x5e9d003b.
//
// Solidity: function setTreasury(address _treasury, uint256 _fee, uint256 _burnFee, bool _mode) returns()
func (_Marketplace *MarketplaceSession) SetTreasury(_treasury common.Address, _fee *big.Int, _burnFee *big.Int, _mode bool) (*types.Transaction, error) {
	return _Marketplace.Contract.SetTreasury(&_Marketplace.TransactOpts, _treasury, _fee, _burnFee, _mode)
}

// SetTreasury is a paid mutator transaction binding the contract method 0x5e9d003b.
//
// Solidity: function setTreasury(address _treasury, uint256 _fee, uint256 _burnFee, bool _mode) returns()
func (_Marketplace *MarketplaceTransactorSession) SetTreasury(_treasury common.Address, _fee *big.Int, _burnFee *big.Int, _mode bool) (*types.Transaction, error) {
	return _Marketplace.Contract.SetTreasury(&_Marketplace.TransactOpts, _treasury, _fee, _burnFee, _mode)
}

// SetWhitelistedNFTs is a paid mutator transaction binding the contract method 0x9014ec21.
//
// Solidity: function setWhitelistedNFTs(address _token, bool _whitelist) returns()
func (_Marketplace *MarketplaceTransactor) SetWhitelistedNFTs(opts *bind.TransactOpts, _token common.Address, _whitelist bool) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "setWhitelistedNFTs", _token, _whitelist)
}

// SetWhitelistedNFTs is a paid mutator transaction binding the contract method 0x9014ec21.
//
// Solidity: function setWhitelistedNFTs(address _token, bool _whitelist) returns()
func (_Marketplace *MarketplaceSession) SetWhitelistedNFTs(_token common.Address, _whitelist bool) (*types.Transaction, error) {
	return _Marketplace.Contract.SetWhitelistedNFTs(&_Marketplace.TransactOpts, _token, _whitelist)
}

// SetWhitelistedNFTs is a paid mutator transaction binding the contract method 0x9014ec21.
//
// Solidity: function setWhitelistedNFTs(address _token, bool _whitelist) returns()
func (_Marketplace *MarketplaceTransactorSession) SetWhitelistedNFTs(_token common.Address, _whitelist bool) (*types.Transaction, error) {
	return _Marketplace.Contract.SetWhitelistedNFTs(&_Marketplace.TransactOpts, _token, _whitelist)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Marketplace *MarketplaceTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Marketplace *MarketplaceSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.TransferOwnership(&_Marketplace.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Marketplace *MarketplaceTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.TransferOwnership(&_Marketplace.TransactOpts, newOwner)
}

// UpdateSale is a paid mutator transaction binding the contract method 0xbdde7897.
//
// Solidity: function updateSale(uint256 _saleId, uint256 _price) returns()
func (_Marketplace *MarketplaceTransactor) UpdateSale(opts *bind.TransactOpts, _saleId *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "updateSale", _saleId, _price)
}

// UpdateSale is a paid mutator transaction binding the contract method 0xbdde7897.
//
// Solidity: function updateSale(uint256 _saleId, uint256 _price) returns()
func (_Marketplace *MarketplaceSession) UpdateSale(_saleId *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.UpdateSale(&_Marketplace.TransactOpts, _saleId, _price)
}

// UpdateSale is a paid mutator transaction binding the contract method 0xbdde7897.
//
// Solidity: function updateSale(uint256 _saleId, uint256 _price) returns()
func (_Marketplace *MarketplaceTransactorSession) UpdateSale(_saleId *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.UpdateSale(&_Marketplace.TransactOpts, _saleId, _price)
}

// MarketplaceOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Marketplace contract.
type MarketplaceOwnershipTransferredIterator struct {
	Event *MarketplaceOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MarketplaceOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceOwnershipTransferred)
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
		it.Event = new(MarketplaceOwnershipTransferred)
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
func (it *MarketplaceOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceOwnershipTransferred represents a OwnershipTransferred event raised by the Marketplace contract.
type MarketplaceOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Marketplace *MarketplaceFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MarketplaceOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceOwnershipTransferredIterator{contract: _Marketplace.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Marketplace *MarketplaceFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MarketplaceOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceOwnershipTransferred)
				if err := _Marketplace.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Marketplace *MarketplaceFilterer) ParseOwnershipTransferred(log types.Log) (*MarketplaceOwnershipTransferred, error) {
	event := new(MarketplaceOwnershipTransferred)
	if err := _Marketplace.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplacePurchasedIterator is returned from FilterPurchased and is used to iterate over the raw logs and unpacked data for Purchased events raised by the Marketplace contract.
type MarketplacePurchasedIterator struct {
	Event *MarketplacePurchased // Event containing the contract specifics and raw log

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
func (it *MarketplacePurchasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplacePurchased)
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
		it.Event = new(MarketplacePurchased)
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
func (it *MarketplacePurchasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplacePurchasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplacePurchased represents a Purchased event raised by the Marketplace contract.
type MarketplacePurchased struct {
	Seller common.Address
	Buyer  common.Address
	SaleId *big.Int
	Price  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPurchased is a free log retrieval operation binding the contract event 0xa326259ec721617acd3cb2a00bcbeac91eefe409880e49aa2bbf473ed648da49.
//
// Solidity: event Purchased(address indexed seller, address indexed buyer, uint256 indexed saleId, uint256 price)
func (_Marketplace *MarketplaceFilterer) FilterPurchased(opts *bind.FilterOpts, seller []common.Address, buyer []common.Address, saleId []*big.Int) (*MarketplacePurchasedIterator, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var saleIdRule []interface{}
	for _, saleIdItem := range saleId {
		saleIdRule = append(saleIdRule, saleIdItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "Purchased", sellerRule, buyerRule, saleIdRule)
	if err != nil {
		return nil, err
	}
	return &MarketplacePurchasedIterator{contract: _Marketplace.contract, event: "Purchased", logs: logs, sub: sub}, nil
}

// WatchPurchased is a free log subscription operation binding the contract event 0xa326259ec721617acd3cb2a00bcbeac91eefe409880e49aa2bbf473ed648da49.
//
// Solidity: event Purchased(address indexed seller, address indexed buyer, uint256 indexed saleId, uint256 price)
func (_Marketplace *MarketplaceFilterer) WatchPurchased(opts *bind.WatchOpts, sink chan<- *MarketplacePurchased, seller []common.Address, buyer []common.Address, saleId []*big.Int) (event.Subscription, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var saleIdRule []interface{}
	for _, saleIdItem := range saleId {
		saleIdRule = append(saleIdRule, saleIdItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "Purchased", sellerRule, buyerRule, saleIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplacePurchased)
				if err := _Marketplace.contract.UnpackLog(event, "Purchased", log); err != nil {
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

// ParsePurchased is a log parse operation binding the contract event 0xa326259ec721617acd3cb2a00bcbeac91eefe409880e49aa2bbf473ed648da49.
//
// Solidity: event Purchased(address indexed seller, address indexed buyer, uint256 indexed saleId, uint256 price)
func (_Marketplace *MarketplaceFilterer) ParsePurchased(log types.Log) (*MarketplacePurchased, error) {
	event := new(MarketplacePurchased)
	if err := _Marketplace.contract.UnpackLog(event, "Purchased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplacePurchasedWithSignatureIterator is returned from FilterPurchasedWithSignature and is used to iterate over the raw logs and unpacked data for PurchasedWithSignature events raised by the Marketplace contract.
type MarketplacePurchasedWithSignatureIterator struct {
	Event *MarketplacePurchasedWithSignature // Event containing the contract specifics and raw log

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
func (it *MarketplacePurchasedWithSignatureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplacePurchasedWithSignature)
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
		it.Event = new(MarketplacePurchasedWithSignature)
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
func (it *MarketplacePurchasedWithSignatureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplacePurchasedWithSignatureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplacePurchasedWithSignature represents a PurchasedWithSignature event raised by the Marketplace contract.
type MarketplacePurchasedWithSignature struct {
	Seller   common.Address
	Buyer    common.Address
	NftToken common.Address
	TokenId  *big.Int
	Price    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPurchasedWithSignature is a free log retrieval operation binding the contract event 0x9c7cafb13f1d4f1fb7524d90bce003abf99d7c2d560a00ae2b68747ecd94f4be.
//
// Solidity: event PurchasedWithSignature(address indexed seller, address indexed buyer, address nftToken, uint256 tokenId, uint256 price)
func (_Marketplace *MarketplaceFilterer) FilterPurchasedWithSignature(opts *bind.FilterOpts, seller []common.Address, buyer []common.Address) (*MarketplacePurchasedWithSignatureIterator, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "PurchasedWithSignature", sellerRule, buyerRule)
	if err != nil {
		return nil, err
	}
	return &MarketplacePurchasedWithSignatureIterator{contract: _Marketplace.contract, event: "PurchasedWithSignature", logs: logs, sub: sub}, nil
}

// WatchPurchasedWithSignature is a free log subscription operation binding the contract event 0x9c7cafb13f1d4f1fb7524d90bce003abf99d7c2d560a00ae2b68747ecd94f4be.
//
// Solidity: event PurchasedWithSignature(address indexed seller, address indexed buyer, address nftToken, uint256 tokenId, uint256 price)
func (_Marketplace *MarketplaceFilterer) WatchPurchasedWithSignature(opts *bind.WatchOpts, sink chan<- *MarketplacePurchasedWithSignature, seller []common.Address, buyer []common.Address) (event.Subscription, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "PurchasedWithSignature", sellerRule, buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplacePurchasedWithSignature)
				if err := _Marketplace.contract.UnpackLog(event, "PurchasedWithSignature", log); err != nil {
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

// ParsePurchasedWithSignature is a log parse operation binding the contract event 0x9c7cafb13f1d4f1fb7524d90bce003abf99d7c2d560a00ae2b68747ecd94f4be.
//
// Solidity: event PurchasedWithSignature(address indexed seller, address indexed buyer, address nftToken, uint256 tokenId, uint256 price)
func (_Marketplace *MarketplaceFilterer) ParsePurchasedWithSignature(log types.Log) (*MarketplacePurchasedWithSignature, error) {
	event := new(MarketplacePurchasedWithSignature)
	if err := _Marketplace.contract.UnpackLog(event, "PurchasedWithSignature", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceSaleCancelledIterator is returned from FilterSaleCancelled and is used to iterate over the raw logs and unpacked data for SaleCancelled events raised by the Marketplace contract.
type MarketplaceSaleCancelledIterator struct {
	Event *MarketplaceSaleCancelled // Event containing the contract specifics and raw log

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
func (it *MarketplaceSaleCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceSaleCancelled)
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
		it.Event = new(MarketplaceSaleCancelled)
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
func (it *MarketplaceSaleCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceSaleCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceSaleCancelled represents a SaleCancelled event raised by the Marketplace contract.
type MarketplaceSaleCancelled struct {
	SaleId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSaleCancelled is a free log retrieval operation binding the contract event 0x2c56893f6f6026d19bd17b7d05c9f15c522de1ae2b1c3a825f91a73c799321f2.
//
// Solidity: event SaleCancelled(uint256 indexed saleId)
func (_Marketplace *MarketplaceFilterer) FilterSaleCancelled(opts *bind.FilterOpts, saleId []*big.Int) (*MarketplaceSaleCancelledIterator, error) {

	var saleIdRule []interface{}
	for _, saleIdItem := range saleId {
		saleIdRule = append(saleIdRule, saleIdItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "SaleCancelled", saleIdRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceSaleCancelledIterator{contract: _Marketplace.contract, event: "SaleCancelled", logs: logs, sub: sub}, nil
}

// WatchSaleCancelled is a free log subscription operation binding the contract event 0x2c56893f6f6026d19bd17b7d05c9f15c522de1ae2b1c3a825f91a73c799321f2.
//
// Solidity: event SaleCancelled(uint256 indexed saleId)
func (_Marketplace *MarketplaceFilterer) WatchSaleCancelled(opts *bind.WatchOpts, sink chan<- *MarketplaceSaleCancelled, saleId []*big.Int) (event.Subscription, error) {

	var saleIdRule []interface{}
	for _, saleIdItem := range saleId {
		saleIdRule = append(saleIdRule, saleIdItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "SaleCancelled", saleIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceSaleCancelled)
				if err := _Marketplace.contract.UnpackLog(event, "SaleCancelled", log); err != nil {
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

// ParseSaleCancelled is a log parse operation binding the contract event 0x2c56893f6f6026d19bd17b7d05c9f15c522de1ae2b1c3a825f91a73c799321f2.
//
// Solidity: event SaleCancelled(uint256 indexed saleId)
func (_Marketplace *MarketplaceFilterer) ParseSaleCancelled(log types.Log) (*MarketplaceSaleCancelled, error) {
	event := new(MarketplaceSaleCancelled)
	if err := _Marketplace.contract.UnpackLog(event, "SaleCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceSaleRequestedIterator is returned from FilterSaleRequested and is used to iterate over the raw logs and unpacked data for SaleRequested events raised by the Marketplace contract.
type MarketplaceSaleRequestedIterator struct {
	Event *MarketplaceSaleRequested // Event containing the contract specifics and raw log

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
func (it *MarketplaceSaleRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceSaleRequested)
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
		it.Event = new(MarketplaceSaleRequested)
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
func (it *MarketplaceSaleRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceSaleRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceSaleRequested represents a SaleRequested event raised by the Marketplace contract.
type MarketplaceSaleRequested struct {
	Seller   common.Address
	SaleId   *big.Int
	NftToken common.Address
	TokenIds []*big.Int
	Price    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSaleRequested is a free log retrieval operation binding the contract event 0x54c64a8e9baf339dcc5553d21c85217d7950f11149e756c2580b0ee1cd4f555d.
//
// Solidity: event SaleRequested(address indexed seller, uint256 indexed saleId, address nftToken, uint256[] tokenIds, uint256 price)
func (_Marketplace *MarketplaceFilterer) FilterSaleRequested(opts *bind.FilterOpts, seller []common.Address, saleId []*big.Int) (*MarketplaceSaleRequestedIterator, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var saleIdRule []interface{}
	for _, saleIdItem := range saleId {
		saleIdRule = append(saleIdRule, saleIdItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "SaleRequested", sellerRule, saleIdRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceSaleRequestedIterator{contract: _Marketplace.contract, event: "SaleRequested", logs: logs, sub: sub}, nil
}

// WatchSaleRequested is a free log subscription operation binding the contract event 0x54c64a8e9baf339dcc5553d21c85217d7950f11149e756c2580b0ee1cd4f555d.
//
// Solidity: event SaleRequested(address indexed seller, uint256 indexed saleId, address nftToken, uint256[] tokenIds, uint256 price)
func (_Marketplace *MarketplaceFilterer) WatchSaleRequested(opts *bind.WatchOpts, sink chan<- *MarketplaceSaleRequested, seller []common.Address, saleId []*big.Int) (event.Subscription, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var saleIdRule []interface{}
	for _, saleIdItem := range saleId {
		saleIdRule = append(saleIdRule, saleIdItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "SaleRequested", sellerRule, saleIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceSaleRequested)
				if err := _Marketplace.contract.UnpackLog(event, "SaleRequested", log); err != nil {
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

// ParseSaleRequested is a log parse operation binding the contract event 0x54c64a8e9baf339dcc5553d21c85217d7950f11149e756c2580b0ee1cd4f555d.
//
// Solidity: event SaleRequested(address indexed seller, uint256 indexed saleId, address nftToken, uint256[] tokenIds, uint256 price)
func (_Marketplace *MarketplaceFilterer) ParseSaleRequested(log types.Log) (*MarketplaceSaleRequested, error) {
	event := new(MarketplaceSaleRequested)
	if err := _Marketplace.contract.UnpackLog(event, "SaleRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceSaleUpdatedIterator is returned from FilterSaleUpdated and is used to iterate over the raw logs and unpacked data for SaleUpdated events raised by the Marketplace contract.
type MarketplaceSaleUpdatedIterator struct {
	Event *MarketplaceSaleUpdated // Event containing the contract specifics and raw log

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
func (it *MarketplaceSaleUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceSaleUpdated)
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
		it.Event = new(MarketplaceSaleUpdated)
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
func (it *MarketplaceSaleUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceSaleUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceSaleUpdated represents a SaleUpdated event raised by the Marketplace contract.
type MarketplaceSaleUpdated struct {
	SaleId   *big.Int
	NewPrice *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSaleUpdated is a free log retrieval operation binding the contract event 0xeefe35a464293e437181813c2025f351aa0efedbb4ada5e147b74bb431c4fd96.
//
// Solidity: event SaleUpdated(uint256 indexed saleId, uint256 newPrice)
func (_Marketplace *MarketplaceFilterer) FilterSaleUpdated(opts *bind.FilterOpts, saleId []*big.Int) (*MarketplaceSaleUpdatedIterator, error) {

	var saleIdRule []interface{}
	for _, saleIdItem := range saleId {
		saleIdRule = append(saleIdRule, saleIdItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "SaleUpdated", saleIdRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceSaleUpdatedIterator{contract: _Marketplace.contract, event: "SaleUpdated", logs: logs, sub: sub}, nil
}

// WatchSaleUpdated is a free log subscription operation binding the contract event 0xeefe35a464293e437181813c2025f351aa0efedbb4ada5e147b74bb431c4fd96.
//
// Solidity: event SaleUpdated(uint256 indexed saleId, uint256 newPrice)
func (_Marketplace *MarketplaceFilterer) WatchSaleUpdated(opts *bind.WatchOpts, sink chan<- *MarketplaceSaleUpdated, saleId []*big.Int) (event.Subscription, error) {

	var saleIdRule []interface{}
	for _, saleIdItem := range saleId {
		saleIdRule = append(saleIdRule, saleIdItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "SaleUpdated", saleIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceSaleUpdated)
				if err := _Marketplace.contract.UnpackLog(event, "SaleUpdated", log); err != nil {
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

// ParseSaleUpdated is a log parse operation binding the contract event 0xeefe35a464293e437181813c2025f351aa0efedbb4ada5e147b74bb431c4fd96.
//
// Solidity: event SaleUpdated(uint256 indexed saleId, uint256 newPrice)
func (_Marketplace *MarketplaceFilterer) ParseSaleUpdated(log types.Log) (*MarketplaceSaleUpdated, error) {
	event := new(MarketplaceSaleUpdated)
	if err := _Marketplace.contract.UnpackLog(event, "SaleUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceTreasuryUpdatedIterator is returned from FilterTreasuryUpdated and is used to iterate over the raw logs and unpacked data for TreasuryUpdated events raised by the Marketplace contract.
type MarketplaceTreasuryUpdatedIterator struct {
	Event *MarketplaceTreasuryUpdated // Event containing the contract specifics and raw log

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
func (it *MarketplaceTreasuryUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceTreasuryUpdated)
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
		it.Event = new(MarketplaceTreasuryUpdated)
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
func (it *MarketplaceTreasuryUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceTreasuryUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceTreasuryUpdated represents a TreasuryUpdated event raised by the Marketplace contract.
type MarketplaceTreasuryUpdated struct {
	Treasury common.Address
	Fee      *big.Int
	BurnFee  *big.Int
	Mode     bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTreasuryUpdated is a free log retrieval operation binding the contract event 0x05efa3de75132e86aa39590bbe333c2ca7b072ffd34e2684dfee84406396cfbf.
//
// Solidity: event TreasuryUpdated(address treasury, uint256 fee, uint256 burnFee, bool mode)
func (_Marketplace *MarketplaceFilterer) FilterTreasuryUpdated(opts *bind.FilterOpts) (*MarketplaceTreasuryUpdatedIterator, error) {

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "TreasuryUpdated")
	if err != nil {
		return nil, err
	}
	return &MarketplaceTreasuryUpdatedIterator{contract: _Marketplace.contract, event: "TreasuryUpdated", logs: logs, sub: sub}, nil
}

// WatchTreasuryUpdated is a free log subscription operation binding the contract event 0x05efa3de75132e86aa39590bbe333c2ca7b072ffd34e2684dfee84406396cfbf.
//
// Solidity: event TreasuryUpdated(address treasury, uint256 fee, uint256 burnFee, bool mode)
func (_Marketplace *MarketplaceFilterer) WatchTreasuryUpdated(opts *bind.WatchOpts, sink chan<- *MarketplaceTreasuryUpdated) (event.Subscription, error) {

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "TreasuryUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceTreasuryUpdated)
				if err := _Marketplace.contract.UnpackLog(event, "TreasuryUpdated", log); err != nil {
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

// ParseTreasuryUpdated is a log parse operation binding the contract event 0x05efa3de75132e86aa39590bbe333c2ca7b072ffd34e2684dfee84406396cfbf.
//
// Solidity: event TreasuryUpdated(address treasury, uint256 fee, uint256 burnFee, bool mode)
func (_Marketplace *MarketplaceFilterer) ParseTreasuryUpdated(log types.Log) (*MarketplaceTreasuryUpdated, error) {
	event := new(MarketplaceTreasuryUpdated)
	if err := _Marketplace.contract.UnpackLog(event, "TreasuryUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
