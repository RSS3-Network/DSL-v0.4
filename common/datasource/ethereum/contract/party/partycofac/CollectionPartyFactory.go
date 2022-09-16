// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package partycofac

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

// CollectionPartyFactoryMetaData contains all meta data concerning the CollectionPartyFactory contract.
var CollectionPartyFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_partyDAOMultisig\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenVaultFactory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_weth\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_allowList\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"partyProxy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"secondsToTimeout\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"deciders\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"splitRecipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"splitBasisPoints\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"gatedToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gatedTokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"name\":\"CollectionPartyDeployed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"deployedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"logic\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"partyDAOMultisig\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_maxPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_secondsToTimeout\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_deciders\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structStructs.AddressAndAmount\",\"name\":\"_split\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structStructs.AddressAndAmount\",\"name\":\"_tokenGate\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"name\":\"startParty\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"partyProxy\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenVaultFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"weth\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// CollectionPartyFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use CollectionPartyFactoryMetaData.ABI instead.
var CollectionPartyFactoryABI = CollectionPartyFactoryMetaData.ABI

// CollectionPartyFactory is an auto generated Go binding around an Ethereum contract.
type CollectionPartyFactory struct {
	CollectionPartyFactoryCaller     // Read-only binding to the contract
	CollectionPartyFactoryTransactor // Write-only binding to the contract
	CollectionPartyFactoryFilterer   // Log filterer for contract events
}

// CollectionPartyFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type CollectionPartyFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CollectionPartyFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CollectionPartyFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CollectionPartyFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CollectionPartyFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CollectionPartyFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CollectionPartyFactorySession struct {
	Contract     *CollectionPartyFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// CollectionPartyFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CollectionPartyFactoryCallerSession struct {
	Contract *CollectionPartyFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// CollectionPartyFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CollectionPartyFactoryTransactorSession struct {
	Contract     *CollectionPartyFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// CollectionPartyFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type CollectionPartyFactoryRaw struct {
	Contract *CollectionPartyFactory // Generic contract binding to access the raw methods on
}

// CollectionPartyFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CollectionPartyFactoryCallerRaw struct {
	Contract *CollectionPartyFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// CollectionPartyFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CollectionPartyFactoryTransactorRaw struct {
	Contract *CollectionPartyFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCollectionPartyFactory creates a new instance of CollectionPartyFactory, bound to a specific deployed contract.
func NewCollectionPartyFactory(address common.Address, backend bind.ContractBackend) (*CollectionPartyFactory, error) {
	contract, err := bindCollectionPartyFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CollectionPartyFactory{CollectionPartyFactoryCaller: CollectionPartyFactoryCaller{contract: contract}, CollectionPartyFactoryTransactor: CollectionPartyFactoryTransactor{contract: contract}, CollectionPartyFactoryFilterer: CollectionPartyFactoryFilterer{contract: contract}}, nil
}

// NewCollectionPartyFactoryCaller creates a new read-only instance of CollectionPartyFactory, bound to a specific deployed contract.
func NewCollectionPartyFactoryCaller(address common.Address, caller bind.ContractCaller) (*CollectionPartyFactoryCaller, error) {
	contract, err := bindCollectionPartyFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CollectionPartyFactoryCaller{contract: contract}, nil
}

// NewCollectionPartyFactoryTransactor creates a new write-only instance of CollectionPartyFactory, bound to a specific deployed contract.
func NewCollectionPartyFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*CollectionPartyFactoryTransactor, error) {
	contract, err := bindCollectionPartyFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CollectionPartyFactoryTransactor{contract: contract}, nil
}

// NewCollectionPartyFactoryFilterer creates a new log filterer instance of CollectionPartyFactory, bound to a specific deployed contract.
func NewCollectionPartyFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*CollectionPartyFactoryFilterer, error) {
	contract, err := bindCollectionPartyFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CollectionPartyFactoryFilterer{contract: contract}, nil
}

// bindCollectionPartyFactory binds a generic wrapper to an already deployed contract.
func bindCollectionPartyFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CollectionPartyFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CollectionPartyFactory *CollectionPartyFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CollectionPartyFactory.Contract.CollectionPartyFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CollectionPartyFactory *CollectionPartyFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CollectionPartyFactory.Contract.CollectionPartyFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CollectionPartyFactory *CollectionPartyFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CollectionPartyFactory.Contract.CollectionPartyFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CollectionPartyFactory *CollectionPartyFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CollectionPartyFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CollectionPartyFactory *CollectionPartyFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CollectionPartyFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CollectionPartyFactory *CollectionPartyFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CollectionPartyFactory.Contract.contract.Transact(opts, method, params...)
}

// DeployedAt is a free data retrieval call binding the contract method 0x57dbdf54.
//
// Solidity: function deployedAt(address ) view returns(uint256)
func (_CollectionPartyFactory *CollectionPartyFactoryCaller) DeployedAt(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CollectionPartyFactory.contract.Call(opts, &out, "deployedAt", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DeployedAt is a free data retrieval call binding the contract method 0x57dbdf54.
//
// Solidity: function deployedAt(address ) view returns(uint256)
func (_CollectionPartyFactory *CollectionPartyFactorySession) DeployedAt(arg0 common.Address) (*big.Int, error) {
	return _CollectionPartyFactory.Contract.DeployedAt(&_CollectionPartyFactory.CallOpts, arg0)
}

// DeployedAt is a free data retrieval call binding the contract method 0x57dbdf54.
//
// Solidity: function deployedAt(address ) view returns(uint256)
func (_CollectionPartyFactory *CollectionPartyFactoryCallerSession) DeployedAt(arg0 common.Address) (*big.Int, error) {
	return _CollectionPartyFactory.Contract.DeployedAt(&_CollectionPartyFactory.CallOpts, arg0)
}

// Logic is a free data retrieval call binding the contract method 0xd7dfa0dd.
//
// Solidity: function logic() view returns(address)
func (_CollectionPartyFactory *CollectionPartyFactoryCaller) Logic(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CollectionPartyFactory.contract.Call(opts, &out, "logic")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Logic is a free data retrieval call binding the contract method 0xd7dfa0dd.
//
// Solidity: function logic() view returns(address)
func (_CollectionPartyFactory *CollectionPartyFactorySession) Logic() (common.Address, error) {
	return _CollectionPartyFactory.Contract.Logic(&_CollectionPartyFactory.CallOpts)
}

// Logic is a free data retrieval call binding the contract method 0xd7dfa0dd.
//
// Solidity: function logic() view returns(address)
func (_CollectionPartyFactory *CollectionPartyFactoryCallerSession) Logic() (common.Address, error) {
	return _CollectionPartyFactory.Contract.Logic(&_CollectionPartyFactory.CallOpts)
}

// PartyDAOMultisig is a free data retrieval call binding the contract method 0x3c4d12d9.
//
// Solidity: function partyDAOMultisig() view returns(address)
func (_CollectionPartyFactory *CollectionPartyFactoryCaller) PartyDAOMultisig(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CollectionPartyFactory.contract.Call(opts, &out, "partyDAOMultisig")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PartyDAOMultisig is a free data retrieval call binding the contract method 0x3c4d12d9.
//
// Solidity: function partyDAOMultisig() view returns(address)
func (_CollectionPartyFactory *CollectionPartyFactorySession) PartyDAOMultisig() (common.Address, error) {
	return _CollectionPartyFactory.Contract.PartyDAOMultisig(&_CollectionPartyFactory.CallOpts)
}

// PartyDAOMultisig is a free data retrieval call binding the contract method 0x3c4d12d9.
//
// Solidity: function partyDAOMultisig() view returns(address)
func (_CollectionPartyFactory *CollectionPartyFactoryCallerSession) PartyDAOMultisig() (common.Address, error) {
	return _CollectionPartyFactory.Contract.PartyDAOMultisig(&_CollectionPartyFactory.CallOpts)
}

// TokenVaultFactory is a free data retrieval call binding the contract method 0x0b203023.
//
// Solidity: function tokenVaultFactory() view returns(address)
func (_CollectionPartyFactory *CollectionPartyFactoryCaller) TokenVaultFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CollectionPartyFactory.contract.Call(opts, &out, "tokenVaultFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenVaultFactory is a free data retrieval call binding the contract method 0x0b203023.
//
// Solidity: function tokenVaultFactory() view returns(address)
func (_CollectionPartyFactory *CollectionPartyFactorySession) TokenVaultFactory() (common.Address, error) {
	return _CollectionPartyFactory.Contract.TokenVaultFactory(&_CollectionPartyFactory.CallOpts)
}

// TokenVaultFactory is a free data retrieval call binding the contract method 0x0b203023.
//
// Solidity: function tokenVaultFactory() view returns(address)
func (_CollectionPartyFactory *CollectionPartyFactoryCallerSession) TokenVaultFactory() (common.Address, error) {
	return _CollectionPartyFactory.Contract.TokenVaultFactory(&_CollectionPartyFactory.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_CollectionPartyFactory *CollectionPartyFactoryCaller) Weth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CollectionPartyFactory.contract.Call(opts, &out, "weth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_CollectionPartyFactory *CollectionPartyFactorySession) Weth() (common.Address, error) {
	return _CollectionPartyFactory.Contract.Weth(&_CollectionPartyFactory.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_CollectionPartyFactory *CollectionPartyFactoryCallerSession) Weth() (common.Address, error) {
	return _CollectionPartyFactory.Contract.Weth(&_CollectionPartyFactory.CallOpts)
}

// StartParty is a paid mutator transaction binding the contract method 0x35a6c767.
//
// Solidity: function startParty(address _nftContract, uint256 _maxPrice, uint256 _secondsToTimeout, address[] _deciders, (address,uint256) _split, (address,uint256) _tokenGate, string _name, string _symbol) returns(address partyProxy)
func (_CollectionPartyFactory *CollectionPartyFactoryTransactor) StartParty(opts *bind.TransactOpts, _nftContract common.Address, _maxPrice *big.Int, _secondsToTimeout *big.Int, _deciders []common.Address, _split StructsAddressAndAmount, _tokenGate StructsAddressAndAmount, _name string, _symbol string) (*types.Transaction, error) {
	return _CollectionPartyFactory.contract.Transact(opts, "startParty", _nftContract, _maxPrice, _secondsToTimeout, _deciders, _split, _tokenGate, _name, _symbol)
}

// StartParty is a paid mutator transaction binding the contract method 0x35a6c767.
//
// Solidity: function startParty(address _nftContract, uint256 _maxPrice, uint256 _secondsToTimeout, address[] _deciders, (address,uint256) _split, (address,uint256) _tokenGate, string _name, string _symbol) returns(address partyProxy)
func (_CollectionPartyFactory *CollectionPartyFactorySession) StartParty(_nftContract common.Address, _maxPrice *big.Int, _secondsToTimeout *big.Int, _deciders []common.Address, _split StructsAddressAndAmount, _tokenGate StructsAddressAndAmount, _name string, _symbol string) (*types.Transaction, error) {
	return _CollectionPartyFactory.Contract.StartParty(&_CollectionPartyFactory.TransactOpts, _nftContract, _maxPrice, _secondsToTimeout, _deciders, _split, _tokenGate, _name, _symbol)
}

// StartParty is a paid mutator transaction binding the contract method 0x35a6c767.
//
// Solidity: function startParty(address _nftContract, uint256 _maxPrice, uint256 _secondsToTimeout, address[] _deciders, (address,uint256) _split, (address,uint256) _tokenGate, string _name, string _symbol) returns(address partyProxy)
func (_CollectionPartyFactory *CollectionPartyFactoryTransactorSession) StartParty(_nftContract common.Address, _maxPrice *big.Int, _secondsToTimeout *big.Int, _deciders []common.Address, _split StructsAddressAndAmount, _tokenGate StructsAddressAndAmount, _name string, _symbol string) (*types.Transaction, error) {
	return _CollectionPartyFactory.Contract.StartParty(&_CollectionPartyFactory.TransactOpts, _nftContract, _maxPrice, _secondsToTimeout, _deciders, _split, _tokenGate, _name, _symbol)
}

// CollectionPartyFactoryCollectionPartyDeployedIterator is returned from FilterCollectionPartyDeployed and is used to iterate over the raw logs and unpacked data for CollectionPartyDeployed events raised by the CollectionPartyFactory contract.
type CollectionPartyFactoryCollectionPartyDeployedIterator struct {
	Event *CollectionPartyFactoryCollectionPartyDeployed // Event containing the contract specifics and raw log

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
func (it *CollectionPartyFactoryCollectionPartyDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollectionPartyFactoryCollectionPartyDeployed)
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
		it.Event = new(CollectionPartyFactoryCollectionPartyDeployed)
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
func (it *CollectionPartyFactoryCollectionPartyDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollectionPartyFactoryCollectionPartyDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollectionPartyFactoryCollectionPartyDeployed represents a CollectionPartyDeployed event raised by the CollectionPartyFactory contract.
type CollectionPartyFactoryCollectionPartyDeployed struct {
	PartyProxy       common.Address
	Creator          common.Address
	NftContract      common.Address
	MaxPrice         *big.Int
	SecondsToTimeout *big.Int
	Deciders         []common.Address
	SplitRecipient   common.Address
	SplitBasisPoints *big.Int
	GatedToken       common.Address
	GatedTokenAmount *big.Int
	Name             string
	Symbol           string
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterCollectionPartyDeployed is a free log retrieval operation binding the contract event 0xd3e9bcbc4c2d9dbcd75b567511ec7f1fd89a91f05e37ae87fa8f2b9f88e4b6d2.
//
// Solidity: event CollectionPartyDeployed(address partyProxy, address creator, address nftContract, uint256 maxPrice, uint256 secondsToTimeout, address[] deciders, address splitRecipient, uint256 splitBasisPoints, address gatedToken, uint256 gatedTokenAmount, string name, string symbol)
func (_CollectionPartyFactory *CollectionPartyFactoryFilterer) FilterCollectionPartyDeployed(opts *bind.FilterOpts) (*CollectionPartyFactoryCollectionPartyDeployedIterator, error) {

	logs, sub, err := _CollectionPartyFactory.contract.FilterLogs(opts, "CollectionPartyDeployed")
	if err != nil {
		return nil, err
	}
	return &CollectionPartyFactoryCollectionPartyDeployedIterator{contract: _CollectionPartyFactory.contract, event: "CollectionPartyDeployed", logs: logs, sub: sub}, nil
}

// WatchCollectionPartyDeployed is a free log subscription operation binding the contract event 0xd3e9bcbc4c2d9dbcd75b567511ec7f1fd89a91f05e37ae87fa8f2b9f88e4b6d2.
//
// Solidity: event CollectionPartyDeployed(address partyProxy, address creator, address nftContract, uint256 maxPrice, uint256 secondsToTimeout, address[] deciders, address splitRecipient, uint256 splitBasisPoints, address gatedToken, uint256 gatedTokenAmount, string name, string symbol)
func (_CollectionPartyFactory *CollectionPartyFactoryFilterer) WatchCollectionPartyDeployed(opts *bind.WatchOpts, sink chan<- *CollectionPartyFactoryCollectionPartyDeployed) (event.Subscription, error) {

	logs, sub, err := _CollectionPartyFactory.contract.WatchLogs(opts, "CollectionPartyDeployed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollectionPartyFactoryCollectionPartyDeployed)
				if err := _CollectionPartyFactory.contract.UnpackLog(event, "CollectionPartyDeployed", log); err != nil {
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

// ParseCollectionPartyDeployed is a log parse operation binding the contract event 0xd3e9bcbc4c2d9dbcd75b567511ec7f1fd89a91f05e37ae87fa8f2b9f88e4b6d2.
//
// Solidity: event CollectionPartyDeployed(address partyProxy, address creator, address nftContract, uint256 maxPrice, uint256 secondsToTimeout, address[] deciders, address splitRecipient, uint256 splitBasisPoints, address gatedToken, uint256 gatedTokenAmount, string name, string symbol)
func (_CollectionPartyFactory *CollectionPartyFactoryFilterer) ParseCollectionPartyDeployed(log types.Log) (*CollectionPartyFactoryCollectionPartyDeployed, error) {
	event := new(CollectionPartyFactoryCollectionPartyDeployed)
	if err := _CollectionPartyFactory.contract.UnpackLog(event, "CollectionPartyDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
