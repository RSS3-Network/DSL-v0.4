// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package periphery

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

// DataTypesERC721Struct is an auto generated low-level Go binding around an user-defined struct.
type DataTypesERC721Struct struct {
	TokenAddress  common.Address
	Erc721TokenId *big.Int
}

// PeripheryMetaData contains all meta data concerning the Periphery contract.
var PeripheryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"linkKey\",\"type\":\"bytes32\"}],\"name\":\"getLinkingAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"}],\"name\":\"getLinkingAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"linkKey\",\"type\":\"bytes32\"}],\"name\":\"getLinkingAnyUri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"}],\"name\":\"getLinkingAnyUris\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"results\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"linkKey\",\"type\":\"bytes32\"}],\"name\":\"getLinkingCharacterId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"}],\"name\":\"getLinkingCharacterIds\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"results\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"linkKey\",\"type\":\"bytes32\"}],\"name\":\"getLinkingERC721\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc721TokenId\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.ERC721Struct\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"linklist\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"web3Entry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// PeripheryABI is the input ABI used to generate the binding from.
// Deprecated: Use PeripheryMetaData.ABI instead.
var PeripheryABI = PeripheryMetaData.ABI

// Periphery is an auto generated Go binding around an Ethereum contract.
type Periphery struct {
	PeripheryCaller     // Read-only binding to the contract
	PeripheryTransactor // Write-only binding to the contract
	PeripheryFilterer   // Log filterer for contract events
}

// PeripheryCaller is an auto generated read-only Go binding around an Ethereum contract.
type PeripheryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PeripheryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PeripheryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PeripheryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PeripheryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PeripherySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PeripherySession struct {
	Contract     *Periphery        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PeripheryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PeripheryCallerSession struct {
	Contract *PeripheryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// PeripheryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PeripheryTransactorSession struct {
	Contract     *PeripheryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// PeripheryRaw is an auto generated low-level Go binding around an Ethereum contract.
type PeripheryRaw struct {
	Contract *Periphery // Generic contract binding to access the raw methods on
}

// PeripheryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PeripheryCallerRaw struct {
	Contract *PeripheryCaller // Generic read-only contract binding to access the raw methods on
}

// PeripheryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PeripheryTransactorRaw struct {
	Contract *PeripheryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPeriphery creates a new instance of Periphery, bound to a specific deployed contract.
func NewPeriphery(address common.Address, backend bind.ContractBackend) (*Periphery, error) {
	contract, err := bindPeriphery(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Periphery{PeripheryCaller: PeripheryCaller{contract: contract}, PeripheryTransactor: PeripheryTransactor{contract: contract}, PeripheryFilterer: PeripheryFilterer{contract: contract}}, nil
}

// NewPeripheryCaller creates a new read-only instance of Periphery, bound to a specific deployed contract.
func NewPeripheryCaller(address common.Address, caller bind.ContractCaller) (*PeripheryCaller, error) {
	contract, err := bindPeriphery(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PeripheryCaller{contract: contract}, nil
}

// NewPeripheryTransactor creates a new write-only instance of Periphery, bound to a specific deployed contract.
func NewPeripheryTransactor(address common.Address, transactor bind.ContractTransactor) (*PeripheryTransactor, error) {
	contract, err := bindPeriphery(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PeripheryTransactor{contract: contract}, nil
}

// NewPeripheryFilterer creates a new log filterer instance of Periphery, bound to a specific deployed contract.
func NewPeripheryFilterer(address common.Address, filterer bind.ContractFilterer) (*PeripheryFilterer, error) {
	contract, err := bindPeriphery(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PeripheryFilterer{contract: contract}, nil
}

// bindPeriphery binds a generic wrapper to an already deployed contract.
func bindPeriphery(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PeripheryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Periphery *PeripheryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Periphery.Contract.PeripheryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Periphery *PeripheryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Periphery.Contract.PeripheryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Periphery *PeripheryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Periphery.Contract.PeripheryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Periphery *PeripheryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Periphery.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Periphery *PeripheryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Periphery.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Periphery *PeripheryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Periphery.Contract.contract.Transact(opts, method, params...)
}

// GetLinkingAddress is a free data retrieval call binding the contract method 0x3e7cb539.
//
// Solidity: function getLinkingAddress(bytes32 linkKey) view returns(address)
func (_Periphery *PeripheryCaller) GetLinkingAddress(opts *bind.CallOpts, linkKey [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Periphery.contract.Call(opts, &out, "getLinkingAddress", linkKey)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLinkingAddress is a free data retrieval call binding the contract method 0x3e7cb539.
//
// Solidity: function getLinkingAddress(bytes32 linkKey) view returns(address)
func (_Periphery *PeripherySession) GetLinkingAddress(linkKey [32]byte) (common.Address, error) {
	return _Periphery.Contract.GetLinkingAddress(&_Periphery.CallOpts, linkKey)
}

// GetLinkingAddress is a free data retrieval call binding the contract method 0x3e7cb539.
//
// Solidity: function getLinkingAddress(bytes32 linkKey) view returns(address)
func (_Periphery *PeripheryCallerSession) GetLinkingAddress(linkKey [32]byte) (common.Address, error) {
	return _Periphery.Contract.GetLinkingAddress(&_Periphery.CallOpts, linkKey)
}

// GetLinkingAddresses is a free data retrieval call binding the contract method 0x76ae90ef.
//
// Solidity: function getLinkingAddresses(uint256 fromCharacterId, bytes32 linkType) view returns(address[])
func (_Periphery *PeripheryCaller) GetLinkingAddresses(opts *bind.CallOpts, fromCharacterId *big.Int, linkType [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _Periphery.contract.Call(opts, &out, "getLinkingAddresses", fromCharacterId, linkType)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetLinkingAddresses is a free data retrieval call binding the contract method 0x76ae90ef.
//
// Solidity: function getLinkingAddresses(uint256 fromCharacterId, bytes32 linkType) view returns(address[])
func (_Periphery *PeripherySession) GetLinkingAddresses(fromCharacterId *big.Int, linkType [32]byte) ([]common.Address, error) {
	return _Periphery.Contract.GetLinkingAddresses(&_Periphery.CallOpts, fromCharacterId, linkType)
}

// GetLinkingAddresses is a free data retrieval call binding the contract method 0x76ae90ef.
//
// Solidity: function getLinkingAddresses(uint256 fromCharacterId, bytes32 linkType) view returns(address[])
func (_Periphery *PeripheryCallerSession) GetLinkingAddresses(fromCharacterId *big.Int, linkType [32]byte) ([]common.Address, error) {
	return _Periphery.Contract.GetLinkingAddresses(&_Periphery.CallOpts, fromCharacterId, linkType)
}

// GetLinkingAnyUri is a free data retrieval call binding the contract method 0x5f799cc6.
//
// Solidity: function getLinkingAnyUri(bytes32 linkKey) view returns(string)
func (_Periphery *PeripheryCaller) GetLinkingAnyUri(opts *bind.CallOpts, linkKey [32]byte) (string, error) {
	var out []interface{}
	err := _Periphery.contract.Call(opts, &out, "getLinkingAnyUri", linkKey)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetLinkingAnyUri is a free data retrieval call binding the contract method 0x5f799cc6.
//
// Solidity: function getLinkingAnyUri(bytes32 linkKey) view returns(string)
func (_Periphery *PeripherySession) GetLinkingAnyUri(linkKey [32]byte) (string, error) {
	return _Periphery.Contract.GetLinkingAnyUri(&_Periphery.CallOpts, linkKey)
}

// GetLinkingAnyUri is a free data retrieval call binding the contract method 0x5f799cc6.
//
// Solidity: function getLinkingAnyUri(bytes32 linkKey) view returns(string)
func (_Periphery *PeripheryCallerSession) GetLinkingAnyUri(linkKey [32]byte) (string, error) {
	return _Periphery.Contract.GetLinkingAnyUri(&_Periphery.CallOpts, linkKey)
}

// GetLinkingAnyUris is a free data retrieval call binding the contract method 0x1964a421.
//
// Solidity: function getLinkingAnyUris(uint256 fromCharacterId, bytes32 linkType) view returns(string[] results)
func (_Periphery *PeripheryCaller) GetLinkingAnyUris(opts *bind.CallOpts, fromCharacterId *big.Int, linkType [32]byte) ([]string, error) {
	var out []interface{}
	err := _Periphery.contract.Call(opts, &out, "getLinkingAnyUris", fromCharacterId, linkType)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetLinkingAnyUris is a free data retrieval call binding the contract method 0x1964a421.
//
// Solidity: function getLinkingAnyUris(uint256 fromCharacterId, bytes32 linkType) view returns(string[] results)
func (_Periphery *PeripherySession) GetLinkingAnyUris(fromCharacterId *big.Int, linkType [32]byte) ([]string, error) {
	return _Periphery.Contract.GetLinkingAnyUris(&_Periphery.CallOpts, fromCharacterId, linkType)
}

// GetLinkingAnyUris is a free data retrieval call binding the contract method 0x1964a421.
//
// Solidity: function getLinkingAnyUris(uint256 fromCharacterId, bytes32 linkType) view returns(string[] results)
func (_Periphery *PeripheryCallerSession) GetLinkingAnyUris(fromCharacterId *big.Int, linkType [32]byte) ([]string, error) {
	return _Periphery.Contract.GetLinkingAnyUris(&_Periphery.CallOpts, fromCharacterId, linkType)
}

// GetLinkingCharacterId is a free data retrieval call binding the contract method 0xd85d04f4.
//
// Solidity: function getLinkingCharacterId(bytes32 linkKey) view returns(uint256 characterId)
func (_Periphery *PeripheryCaller) GetLinkingCharacterId(opts *bind.CallOpts, linkKey [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Periphery.contract.Call(opts, &out, "getLinkingCharacterId", linkKey)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLinkingCharacterId is a free data retrieval call binding the contract method 0xd85d04f4.
//
// Solidity: function getLinkingCharacterId(bytes32 linkKey) view returns(uint256 characterId)
func (_Periphery *PeripherySession) GetLinkingCharacterId(linkKey [32]byte) (*big.Int, error) {
	return _Periphery.Contract.GetLinkingCharacterId(&_Periphery.CallOpts, linkKey)
}

// GetLinkingCharacterId is a free data retrieval call binding the contract method 0xd85d04f4.
//
// Solidity: function getLinkingCharacterId(bytes32 linkKey) view returns(uint256 characterId)
func (_Periphery *PeripheryCallerSession) GetLinkingCharacterId(linkKey [32]byte) (*big.Int, error) {
	return _Periphery.Contract.GetLinkingCharacterId(&_Periphery.CallOpts, linkKey)
}

// GetLinkingCharacterIds is a free data retrieval call binding the contract method 0xdfc1e2bf.
//
// Solidity: function getLinkingCharacterIds(uint256 fromCharacterId, bytes32 linkType) view returns(uint256[] results)
func (_Periphery *PeripheryCaller) GetLinkingCharacterIds(opts *bind.CallOpts, fromCharacterId *big.Int, linkType [32]byte) ([]*big.Int, error) {
	var out []interface{}
	err := _Periphery.contract.Call(opts, &out, "getLinkingCharacterIds", fromCharacterId, linkType)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetLinkingCharacterIds is a free data retrieval call binding the contract method 0xdfc1e2bf.
//
// Solidity: function getLinkingCharacterIds(uint256 fromCharacterId, bytes32 linkType) view returns(uint256[] results)
func (_Periphery *PeripherySession) GetLinkingCharacterIds(fromCharacterId *big.Int, linkType [32]byte) ([]*big.Int, error) {
	return _Periphery.Contract.GetLinkingCharacterIds(&_Periphery.CallOpts, fromCharacterId, linkType)
}

// GetLinkingCharacterIds is a free data retrieval call binding the contract method 0xdfc1e2bf.
//
// Solidity: function getLinkingCharacterIds(uint256 fromCharacterId, bytes32 linkType) view returns(uint256[] results)
func (_Periphery *PeripheryCallerSession) GetLinkingCharacterIds(fromCharacterId *big.Int, linkType [32]byte) ([]*big.Int, error) {
	return _Periphery.Contract.GetLinkingCharacterIds(&_Periphery.CallOpts, fromCharacterId, linkType)
}

// GetLinkingERC721 is a free data retrieval call binding the contract method 0x5440522b.
//
// Solidity: function getLinkingERC721(bytes32 linkKey) view returns((address,uint256))
func (_Periphery *PeripheryCaller) GetLinkingERC721(opts *bind.CallOpts, linkKey [32]byte) (DataTypesERC721Struct, error) {
	var out []interface{}
	err := _Periphery.contract.Call(opts, &out, "getLinkingERC721", linkKey)

	if err != nil {
		return *new(DataTypesERC721Struct), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesERC721Struct)).(*DataTypesERC721Struct)

	return out0, err

}

// GetLinkingERC721 is a free data retrieval call binding the contract method 0x5440522b.
//
// Solidity: function getLinkingERC721(bytes32 linkKey) view returns((address,uint256))
func (_Periphery *PeripherySession) GetLinkingERC721(linkKey [32]byte) (DataTypesERC721Struct, error) {
	return _Periphery.Contract.GetLinkingERC721(&_Periphery.CallOpts, linkKey)
}

// GetLinkingERC721 is a free data retrieval call binding the contract method 0x5440522b.
//
// Solidity: function getLinkingERC721(bytes32 linkKey) view returns((address,uint256))
func (_Periphery *PeripheryCallerSession) GetLinkingERC721(linkKey [32]byte) (DataTypesERC721Struct, error) {
	return _Periphery.Contract.GetLinkingERC721(&_Periphery.CallOpts, linkKey)
}

// Linklist is a free data retrieval call binding the contract method 0x06bee8d4.
//
// Solidity: function linklist() view returns(address)
func (_Periphery *PeripheryCaller) Linklist(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Periphery.contract.Call(opts, &out, "linklist")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Linklist is a free data retrieval call binding the contract method 0x06bee8d4.
//
// Solidity: function linklist() view returns(address)
func (_Periphery *PeripherySession) Linklist() (common.Address, error) {
	return _Periphery.Contract.Linklist(&_Periphery.CallOpts)
}

// Linklist is a free data retrieval call binding the contract method 0x06bee8d4.
//
// Solidity: function linklist() view returns(address)
func (_Periphery *PeripheryCallerSession) Linklist() (common.Address, error) {
	return _Periphery.Contract.Linklist(&_Periphery.CallOpts)
}

// Web3Entry is a free data retrieval call binding the contract method 0xe1332272.
//
// Solidity: function web3Entry() view returns(address)
func (_Periphery *PeripheryCaller) Web3Entry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Periphery.contract.Call(opts, &out, "web3Entry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Web3Entry is a free data retrieval call binding the contract method 0xe1332272.
//
// Solidity: function web3Entry() view returns(address)
func (_Periphery *PeripherySession) Web3Entry() (common.Address, error) {
	return _Periphery.Contract.Web3Entry(&_Periphery.CallOpts)
}

// Web3Entry is a free data retrieval call binding the contract method 0xe1332272.
//
// Solidity: function web3Entry() view returns(address)
func (_Periphery *PeripheryCallerSession) Web3Entry() (common.Address, error) {
	return _Periphery.Contract.Web3Entry(&_Periphery.CallOpts)
}
