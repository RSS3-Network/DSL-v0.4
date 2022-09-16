// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package partybuyfac

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

// PartyBuyFactoryMetaData contains all meta data concerning the PartyBuyFactory contract.
var PartyBuyFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_partyDAOMultisig\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenVaultFactory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_weth\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_allowList\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_logicNftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_logicTokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"partyProxy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"secondsToTimeout\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"splitRecipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"splitBasisPoints\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"gatedToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gatedTokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"name\":\"PartyBuyDeployed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"deployedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"logic\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"partyDAOMultisig\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_secondsToTimeout\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structStructs.AddressAndAmount\",\"name\":\"_split\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structStructs.AddressAndAmount\",\"name\":\"_tokenGate\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"name\":\"startParty\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"partyBuyProxy\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenVaultFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"weth\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// PartyBuyFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use PartyBuyFactoryMetaData.ABI instead.
var PartyBuyFactoryABI = PartyBuyFactoryMetaData.ABI

// PartyBuyFactory is an auto generated Go binding around an Ethereum contract.
type PartyBuyFactory struct {
	PartyBuyFactoryCaller     // Read-only binding to the contract
	PartyBuyFactoryTransactor // Write-only binding to the contract
	PartyBuyFactoryFilterer   // Log filterer for contract events
}

// PartyBuyFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type PartyBuyFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PartyBuyFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PartyBuyFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PartyBuyFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PartyBuyFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PartyBuyFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PartyBuyFactorySession struct {
	Contract     *PartyBuyFactory  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PartyBuyFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PartyBuyFactoryCallerSession struct {
	Contract *PartyBuyFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// PartyBuyFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PartyBuyFactoryTransactorSession struct {
	Contract     *PartyBuyFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// PartyBuyFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type PartyBuyFactoryRaw struct {
	Contract *PartyBuyFactory // Generic contract binding to access the raw methods on
}

// PartyBuyFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PartyBuyFactoryCallerRaw struct {
	Contract *PartyBuyFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// PartyBuyFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PartyBuyFactoryTransactorRaw struct {
	Contract *PartyBuyFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPartyBuyFactory creates a new instance of PartyBuyFactory, bound to a specific deployed contract.
func NewPartyBuyFactory(address common.Address, backend bind.ContractBackend) (*PartyBuyFactory, error) {
	contract, err := bindPartyBuyFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PartyBuyFactory{PartyBuyFactoryCaller: PartyBuyFactoryCaller{contract: contract}, PartyBuyFactoryTransactor: PartyBuyFactoryTransactor{contract: contract}, PartyBuyFactoryFilterer: PartyBuyFactoryFilterer{contract: contract}}, nil
}

// NewPartyBuyFactoryCaller creates a new read-only instance of PartyBuyFactory, bound to a specific deployed contract.
func NewPartyBuyFactoryCaller(address common.Address, caller bind.ContractCaller) (*PartyBuyFactoryCaller, error) {
	contract, err := bindPartyBuyFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PartyBuyFactoryCaller{contract: contract}, nil
}

// NewPartyBuyFactoryTransactor creates a new write-only instance of PartyBuyFactory, bound to a specific deployed contract.
func NewPartyBuyFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*PartyBuyFactoryTransactor, error) {
	contract, err := bindPartyBuyFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PartyBuyFactoryTransactor{contract: contract}, nil
}

// NewPartyBuyFactoryFilterer creates a new log filterer instance of PartyBuyFactory, bound to a specific deployed contract.
func NewPartyBuyFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*PartyBuyFactoryFilterer, error) {
	contract, err := bindPartyBuyFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PartyBuyFactoryFilterer{contract: contract}, nil
}

// bindPartyBuyFactory binds a generic wrapper to an already deployed contract.
func bindPartyBuyFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PartyBuyFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PartyBuyFactory *PartyBuyFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PartyBuyFactory.Contract.PartyBuyFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PartyBuyFactory *PartyBuyFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PartyBuyFactory.Contract.PartyBuyFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PartyBuyFactory *PartyBuyFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PartyBuyFactory.Contract.PartyBuyFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PartyBuyFactory *PartyBuyFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PartyBuyFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PartyBuyFactory *PartyBuyFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PartyBuyFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PartyBuyFactory *PartyBuyFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PartyBuyFactory.Contract.contract.Transact(opts, method, params...)
}

// DeployedAt is a free data retrieval call binding the contract method 0x57dbdf54.
//
// Solidity: function deployedAt(address ) view returns(uint256)
func (_PartyBuyFactory *PartyBuyFactoryCaller) DeployedAt(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PartyBuyFactory.contract.Call(opts, &out, "deployedAt", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DeployedAt is a free data retrieval call binding the contract method 0x57dbdf54.
//
// Solidity: function deployedAt(address ) view returns(uint256)
func (_PartyBuyFactory *PartyBuyFactorySession) DeployedAt(arg0 common.Address) (*big.Int, error) {
	return _PartyBuyFactory.Contract.DeployedAt(&_PartyBuyFactory.CallOpts, arg0)
}

// DeployedAt is a free data retrieval call binding the contract method 0x57dbdf54.
//
// Solidity: function deployedAt(address ) view returns(uint256)
func (_PartyBuyFactory *PartyBuyFactoryCallerSession) DeployedAt(arg0 common.Address) (*big.Int, error) {
	return _PartyBuyFactory.Contract.DeployedAt(&_PartyBuyFactory.CallOpts, arg0)
}

// Logic is a free data retrieval call binding the contract method 0xd7dfa0dd.
//
// Solidity: function logic() view returns(address)
func (_PartyBuyFactory *PartyBuyFactoryCaller) Logic(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PartyBuyFactory.contract.Call(opts, &out, "logic")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Logic is a free data retrieval call binding the contract method 0xd7dfa0dd.
//
// Solidity: function logic() view returns(address)
func (_PartyBuyFactory *PartyBuyFactorySession) Logic() (common.Address, error) {
	return _PartyBuyFactory.Contract.Logic(&_PartyBuyFactory.CallOpts)
}

// Logic is a free data retrieval call binding the contract method 0xd7dfa0dd.
//
// Solidity: function logic() view returns(address)
func (_PartyBuyFactory *PartyBuyFactoryCallerSession) Logic() (common.Address, error) {
	return _PartyBuyFactory.Contract.Logic(&_PartyBuyFactory.CallOpts)
}

// PartyDAOMultisig is a free data retrieval call binding the contract method 0x3c4d12d9.
//
// Solidity: function partyDAOMultisig() view returns(address)
func (_PartyBuyFactory *PartyBuyFactoryCaller) PartyDAOMultisig(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PartyBuyFactory.contract.Call(opts, &out, "partyDAOMultisig")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PartyDAOMultisig is a free data retrieval call binding the contract method 0x3c4d12d9.
//
// Solidity: function partyDAOMultisig() view returns(address)
func (_PartyBuyFactory *PartyBuyFactorySession) PartyDAOMultisig() (common.Address, error) {
	return _PartyBuyFactory.Contract.PartyDAOMultisig(&_PartyBuyFactory.CallOpts)
}

// PartyDAOMultisig is a free data retrieval call binding the contract method 0x3c4d12d9.
//
// Solidity: function partyDAOMultisig() view returns(address)
func (_PartyBuyFactory *PartyBuyFactoryCallerSession) PartyDAOMultisig() (common.Address, error) {
	return _PartyBuyFactory.Contract.PartyDAOMultisig(&_PartyBuyFactory.CallOpts)
}

// TokenVaultFactory is a free data retrieval call binding the contract method 0x0b203023.
//
// Solidity: function tokenVaultFactory() view returns(address)
func (_PartyBuyFactory *PartyBuyFactoryCaller) TokenVaultFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PartyBuyFactory.contract.Call(opts, &out, "tokenVaultFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenVaultFactory is a free data retrieval call binding the contract method 0x0b203023.
//
// Solidity: function tokenVaultFactory() view returns(address)
func (_PartyBuyFactory *PartyBuyFactorySession) TokenVaultFactory() (common.Address, error) {
	return _PartyBuyFactory.Contract.TokenVaultFactory(&_PartyBuyFactory.CallOpts)
}

// TokenVaultFactory is a free data retrieval call binding the contract method 0x0b203023.
//
// Solidity: function tokenVaultFactory() view returns(address)
func (_PartyBuyFactory *PartyBuyFactoryCallerSession) TokenVaultFactory() (common.Address, error) {
	return _PartyBuyFactory.Contract.TokenVaultFactory(&_PartyBuyFactory.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_PartyBuyFactory *PartyBuyFactoryCaller) Weth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PartyBuyFactory.contract.Call(opts, &out, "weth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_PartyBuyFactory *PartyBuyFactorySession) Weth() (common.Address, error) {
	return _PartyBuyFactory.Contract.Weth(&_PartyBuyFactory.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_PartyBuyFactory *PartyBuyFactoryCallerSession) Weth() (common.Address, error) {
	return _PartyBuyFactory.Contract.Weth(&_PartyBuyFactory.CallOpts)
}

// StartParty is a paid mutator transaction binding the contract method 0x7da82d7d.
//
// Solidity: function startParty(address _nftContract, uint256 _tokenId, uint256 _maxPrice, uint256 _secondsToTimeout, (address,uint256) _split, (address,uint256) _tokenGate, string _name, string _symbol) returns(address partyBuyProxy)
func (_PartyBuyFactory *PartyBuyFactoryTransactor) StartParty(opts *bind.TransactOpts, _nftContract common.Address, _tokenId *big.Int, _maxPrice *big.Int, _secondsToTimeout *big.Int, _split StructsAddressAndAmount, _tokenGate StructsAddressAndAmount, _name string, _symbol string) (*types.Transaction, error) {
	return _PartyBuyFactory.contract.Transact(opts, "startParty", _nftContract, _tokenId, _maxPrice, _secondsToTimeout, _split, _tokenGate, _name, _symbol)
}

// StartParty is a paid mutator transaction binding the contract method 0x7da82d7d.
//
// Solidity: function startParty(address _nftContract, uint256 _tokenId, uint256 _maxPrice, uint256 _secondsToTimeout, (address,uint256) _split, (address,uint256) _tokenGate, string _name, string _symbol) returns(address partyBuyProxy)
func (_PartyBuyFactory *PartyBuyFactorySession) StartParty(_nftContract common.Address, _tokenId *big.Int, _maxPrice *big.Int, _secondsToTimeout *big.Int, _split StructsAddressAndAmount, _tokenGate StructsAddressAndAmount, _name string, _symbol string) (*types.Transaction, error) {
	return _PartyBuyFactory.Contract.StartParty(&_PartyBuyFactory.TransactOpts, _nftContract, _tokenId, _maxPrice, _secondsToTimeout, _split, _tokenGate, _name, _symbol)
}

// StartParty is a paid mutator transaction binding the contract method 0x7da82d7d.
//
// Solidity: function startParty(address _nftContract, uint256 _tokenId, uint256 _maxPrice, uint256 _secondsToTimeout, (address,uint256) _split, (address,uint256) _tokenGate, string _name, string _symbol) returns(address partyBuyProxy)
func (_PartyBuyFactory *PartyBuyFactoryTransactorSession) StartParty(_nftContract common.Address, _tokenId *big.Int, _maxPrice *big.Int, _secondsToTimeout *big.Int, _split StructsAddressAndAmount, _tokenGate StructsAddressAndAmount, _name string, _symbol string) (*types.Transaction, error) {
	return _PartyBuyFactory.Contract.StartParty(&_PartyBuyFactory.TransactOpts, _nftContract, _tokenId, _maxPrice, _secondsToTimeout, _split, _tokenGate, _name, _symbol)
}

// PartyBuyFactoryPartyBuyDeployedIterator is returned from FilterPartyBuyDeployed and is used to iterate over the raw logs and unpacked data for PartyBuyDeployed events raised by the PartyBuyFactory contract.
type PartyBuyFactoryPartyBuyDeployedIterator struct {
	Event *PartyBuyFactoryPartyBuyDeployed // Event containing the contract specifics and raw log

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
func (it *PartyBuyFactoryPartyBuyDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PartyBuyFactoryPartyBuyDeployed)
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
		it.Event = new(PartyBuyFactoryPartyBuyDeployed)
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
func (it *PartyBuyFactoryPartyBuyDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PartyBuyFactoryPartyBuyDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PartyBuyFactoryPartyBuyDeployed represents a PartyBuyDeployed event raised by the PartyBuyFactory contract.
type PartyBuyFactoryPartyBuyDeployed struct {
	PartyProxy       common.Address
	Creator          common.Address
	NftContract      common.Address
	TokenId          *big.Int
	MaxPrice         *big.Int
	SecondsToTimeout *big.Int
	SplitRecipient   common.Address
	SplitBasisPoints *big.Int
	GatedToken       common.Address
	GatedTokenAmount *big.Int
	Name             string
	Symbol           string
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterPartyBuyDeployed is a free log retrieval operation binding the contract event 0xa9072e473c4d01d531fe9f31078a1b2566e08692f30ef95e9afad42b8d25ecd1.
//
// Solidity: event PartyBuyDeployed(address partyProxy, address creator, address nftContract, uint256 tokenId, uint256 maxPrice, uint256 secondsToTimeout, address splitRecipient, uint256 splitBasisPoints, address gatedToken, uint256 gatedTokenAmount, string name, string symbol)
func (_PartyBuyFactory *PartyBuyFactoryFilterer) FilterPartyBuyDeployed(opts *bind.FilterOpts) (*PartyBuyFactoryPartyBuyDeployedIterator, error) {

	logs, sub, err := _PartyBuyFactory.contract.FilterLogs(opts, "PartyBuyDeployed")
	if err != nil {
		return nil, err
	}
	return &PartyBuyFactoryPartyBuyDeployedIterator{contract: _PartyBuyFactory.contract, event: "PartyBuyDeployed", logs: logs, sub: sub}, nil
}

// WatchPartyBuyDeployed is a free log subscription operation binding the contract event 0xa9072e473c4d01d531fe9f31078a1b2566e08692f30ef95e9afad42b8d25ecd1.
//
// Solidity: event PartyBuyDeployed(address partyProxy, address creator, address nftContract, uint256 tokenId, uint256 maxPrice, uint256 secondsToTimeout, address splitRecipient, uint256 splitBasisPoints, address gatedToken, uint256 gatedTokenAmount, string name, string symbol)
func (_PartyBuyFactory *PartyBuyFactoryFilterer) WatchPartyBuyDeployed(opts *bind.WatchOpts, sink chan<- *PartyBuyFactoryPartyBuyDeployed) (event.Subscription, error) {

	logs, sub, err := _PartyBuyFactory.contract.WatchLogs(opts, "PartyBuyDeployed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PartyBuyFactoryPartyBuyDeployed)
				if err := _PartyBuyFactory.contract.UnpackLog(event, "PartyBuyDeployed", log); err != nil {
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

// ParsePartyBuyDeployed is a log parse operation binding the contract event 0xa9072e473c4d01d531fe9f31078a1b2566e08692f30ef95e9afad42b8d25ecd1.
//
// Solidity: event PartyBuyDeployed(address partyProxy, address creator, address nftContract, uint256 tokenId, uint256 maxPrice, uint256 secondsToTimeout, address splitRecipient, uint256 splitBasisPoints, address gatedToken, uint256 gatedTokenAmount, string name, string symbol)
func (_PartyBuyFactory *PartyBuyFactoryFilterer) ParsePartyBuyDeployed(log types.Log) (*PartyBuyFactoryPartyBuyDeployed, error) {
	event := new(PartyBuyFactoryPartyBuyDeployed)
	if err := _PartyBuyFactory.contract.UnpackLog(event, "PartyBuyDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
