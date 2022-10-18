// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// UnstoppableDomainsMetaData contains all meta data concerning the UnstoppableDomains contract.
var UnstoppableDomainsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"RemoveReverse\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"SetReverse\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"removeReverse\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"reverseOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"setReverse\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// UnstoppableDomainsABI is the input ABI used to generate the binding from.
// Deprecated: Use UnstoppableDomainsMetaData.ABI instead.
var UnstoppableDomainsABI = UnstoppableDomainsMetaData.ABI

// UnstoppableDomains is an auto generated Go binding around an Ethereum contract.
type UnstoppableDomains struct {
	UnstoppableDomainsCaller     // Read-only binding to the contract
	UnstoppableDomainsTransactor // Write-only binding to the contract
	UnstoppableDomainsFilterer   // Log filterer for contract events
}

// UnstoppableDomainsCaller is an auto generated read-only Go binding around an Ethereum contract.
type UnstoppableDomainsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnstoppableDomainsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UnstoppableDomainsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnstoppableDomainsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UnstoppableDomainsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnstoppableDomainsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UnstoppableDomainsSession struct {
	Contract     *UnstoppableDomains // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// UnstoppableDomainsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UnstoppableDomainsCallerSession struct {
	Contract *UnstoppableDomainsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// UnstoppableDomainsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UnstoppableDomainsTransactorSession struct {
	Contract     *UnstoppableDomainsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// UnstoppableDomainsRaw is an auto generated low-level Go binding around an Ethereum contract.
type UnstoppableDomainsRaw struct {
	Contract *UnstoppableDomains // Generic contract binding to access the raw methods on
}

// UnstoppableDomainsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UnstoppableDomainsCallerRaw struct {
	Contract *UnstoppableDomainsCaller // Generic read-only contract binding to access the raw methods on
}

// UnstoppableDomainsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UnstoppableDomainsTransactorRaw struct {
	Contract *UnstoppableDomainsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUnstoppableDomains creates a new instance of UnstoppableDomains, bound to a specific deployed contract.
func NewUnstoppableDomains(address common.Address, backend bind.ContractBackend) (*UnstoppableDomains, error) {
	contract, err := bindUnstoppableDomains(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UnstoppableDomains{UnstoppableDomainsCaller: UnstoppableDomainsCaller{contract: contract}, UnstoppableDomainsTransactor: UnstoppableDomainsTransactor{contract: contract}, UnstoppableDomainsFilterer: UnstoppableDomainsFilterer{contract: contract}}, nil
}

// NewUnstoppableDomainsCaller creates a new read-only instance of UnstoppableDomains, bound to a specific deployed contract.
func NewUnstoppableDomainsCaller(address common.Address, caller bind.ContractCaller) (*UnstoppableDomainsCaller, error) {
	contract, err := bindUnstoppableDomains(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UnstoppableDomainsCaller{contract: contract}, nil
}

// NewUnstoppableDomainsTransactor creates a new write-only instance of UnstoppableDomains, bound to a specific deployed contract.
func NewUnstoppableDomainsTransactor(address common.Address, transactor bind.ContractTransactor) (*UnstoppableDomainsTransactor, error) {
	contract, err := bindUnstoppableDomains(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UnstoppableDomainsTransactor{contract: contract}, nil
}

// NewUnstoppableDomainsFilterer creates a new log filterer instance of UnstoppableDomains, bound to a specific deployed contract.
func NewUnstoppableDomainsFilterer(address common.Address, filterer bind.ContractFilterer) (*UnstoppableDomainsFilterer, error) {
	contract, err := bindUnstoppableDomains(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UnstoppableDomainsFilterer{contract: contract}, nil
}

// bindUnstoppableDomains binds a generic wrapper to an already deployed contract.
func bindUnstoppableDomains(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UnstoppableDomainsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UnstoppableDomains *UnstoppableDomainsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UnstoppableDomains.Contract.UnstoppableDomainsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UnstoppableDomains *UnstoppableDomainsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UnstoppableDomains.Contract.UnstoppableDomainsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UnstoppableDomains *UnstoppableDomainsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UnstoppableDomains.Contract.UnstoppableDomainsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UnstoppableDomains *UnstoppableDomainsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UnstoppableDomains.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UnstoppableDomains *UnstoppableDomainsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UnstoppableDomains.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UnstoppableDomains *UnstoppableDomainsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UnstoppableDomains.Contract.contract.Transact(opts, method, params...)
}

// ReverseOf is a free data retrieval call binding the contract method 0x7e37479e.
//
// Solidity: function reverseOf(address addr) view returns(uint256)
func (_UnstoppableDomains *UnstoppableDomainsCaller) ReverseOf(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _UnstoppableDomains.contract.Call(opts, &out, "reverseOf", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReverseOf is a free data retrieval call binding the contract method 0x7e37479e.
//
// Solidity: function reverseOf(address addr) view returns(uint256)
func (_UnstoppableDomains *UnstoppableDomainsSession) ReverseOf(addr common.Address) (*big.Int, error) {
	return _UnstoppableDomains.Contract.ReverseOf(&_UnstoppableDomains.CallOpts, addr)
}

// ReverseOf is a free data retrieval call binding the contract method 0x7e37479e.
//
// Solidity: function reverseOf(address addr) view returns(uint256)
func (_UnstoppableDomains *UnstoppableDomainsCallerSession) ReverseOf(addr common.Address) (*big.Int, error) {
	return _UnstoppableDomains.Contract.ReverseOf(&_UnstoppableDomains.CallOpts, addr)
}

// RemoveReverse is a paid mutator transaction binding the contract method 0xf25eb5c1.
//
// Solidity: function removeReverse() returns()
func (_UnstoppableDomains *UnstoppableDomainsTransactor) RemoveReverse(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UnstoppableDomains.contract.Transact(opts, "removeReverse")
}

// RemoveReverse is a paid mutator transaction binding the contract method 0xf25eb5c1.
//
// Solidity: function removeReverse() returns()
func (_UnstoppableDomains *UnstoppableDomainsSession) RemoveReverse() (*types.Transaction, error) {
	return _UnstoppableDomains.Contract.RemoveReverse(&_UnstoppableDomains.TransactOpts)
}

// RemoveReverse is a paid mutator transaction binding the contract method 0xf25eb5c1.
//
// Solidity: function removeReverse() returns()
func (_UnstoppableDomains *UnstoppableDomainsTransactorSession) RemoveReverse() (*types.Transaction, error) {
	return _UnstoppableDomains.Contract.RemoveReverse(&_UnstoppableDomains.TransactOpts)
}

// SetReverse is a paid mutator transaction binding the contract method 0x384e9a55.
//
// Solidity: function setReverse(uint256 tokenId) returns()
func (_UnstoppableDomains *UnstoppableDomainsTransactor) SetReverse(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _UnstoppableDomains.contract.Transact(opts, "setReverse", tokenId)
}

// SetReverse is a paid mutator transaction binding the contract method 0x384e9a55.
//
// Solidity: function setReverse(uint256 tokenId) returns()
func (_UnstoppableDomains *UnstoppableDomainsSession) SetReverse(tokenId *big.Int) (*types.Transaction, error) {
	return _UnstoppableDomains.Contract.SetReverse(&_UnstoppableDomains.TransactOpts, tokenId)
}

// SetReverse is a paid mutator transaction binding the contract method 0x384e9a55.
//
// Solidity: function setReverse(uint256 tokenId) returns()
func (_UnstoppableDomains *UnstoppableDomainsTransactorSession) SetReverse(tokenId *big.Int) (*types.Transaction, error) {
	return _UnstoppableDomains.Contract.SetReverse(&_UnstoppableDomains.TransactOpts, tokenId)
}

// UnstoppableDomainsRemoveReverseIterator is returned from FilterRemoveReverse and is used to iterate over the raw logs and unpacked data for RemoveReverse events raised by the UnstoppableDomains contract.
type UnstoppableDomainsRemoveReverseIterator struct {
	Event *UnstoppableDomainsRemoveReverse // Event containing the contract specifics and raw log

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
func (it *UnstoppableDomainsRemoveReverseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnstoppableDomainsRemoveReverse)
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
		it.Event = new(UnstoppableDomainsRemoveReverse)
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
func (it *UnstoppableDomainsRemoveReverseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnstoppableDomainsRemoveReverseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnstoppableDomainsRemoveReverse represents a RemoveReverse event raised by the UnstoppableDomains contract.
type UnstoppableDomainsRemoveReverse struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRemoveReverse is a free log retrieval operation binding the contract event 0xfcf5eec0cfa3e6332f5f0e63ec242d71f866a61d121d6cdf5c2eb3b668a26c4f.
//
// Solidity: event RemoveReverse(address indexed addr)
func (_UnstoppableDomains *UnstoppableDomainsFilterer) FilterRemoveReverse(opts *bind.FilterOpts, addr []common.Address) (*UnstoppableDomainsRemoveReverseIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _UnstoppableDomains.contract.FilterLogs(opts, "RemoveReverse", addrRule)
	if err != nil {
		return nil, err
	}
	return &UnstoppableDomainsRemoveReverseIterator{contract: _UnstoppableDomains.contract, event: "RemoveReverse", logs: logs, sub: sub}, nil
}

// WatchRemoveReverse is a free log subscription operation binding the contract event 0xfcf5eec0cfa3e6332f5f0e63ec242d71f866a61d121d6cdf5c2eb3b668a26c4f.
//
// Solidity: event RemoveReverse(address indexed addr)
func (_UnstoppableDomains *UnstoppableDomainsFilterer) WatchRemoveReverse(opts *bind.WatchOpts, sink chan<- *UnstoppableDomainsRemoveReverse, addr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _UnstoppableDomains.contract.WatchLogs(opts, "RemoveReverse", addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnstoppableDomainsRemoveReverse)
				if err := _UnstoppableDomains.contract.UnpackLog(event, "RemoveReverse", log); err != nil {
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

// ParseRemoveReverse is a log parse operation binding the contract event 0xfcf5eec0cfa3e6332f5f0e63ec242d71f866a61d121d6cdf5c2eb3b668a26c4f.
//
// Solidity: event RemoveReverse(address indexed addr)
func (_UnstoppableDomains *UnstoppableDomainsFilterer) ParseRemoveReverse(log types.Log) (*UnstoppableDomainsRemoveReverse, error) {
	event := new(UnstoppableDomainsRemoveReverse)
	if err := _UnstoppableDomains.contract.UnpackLog(event, "RemoveReverse", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnstoppableDomainsSetReverseIterator is returned from FilterSetReverse and is used to iterate over the raw logs and unpacked data for SetReverse events raised by the UnstoppableDomains contract.
type UnstoppableDomainsSetReverseIterator struct {
	Event *UnstoppableDomainsSetReverse // Event containing the contract specifics and raw log

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
func (it *UnstoppableDomainsSetReverseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnstoppableDomainsSetReverse)
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
		it.Event = new(UnstoppableDomainsSetReverse)
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
func (it *UnstoppableDomainsSetReverseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnstoppableDomainsSetReverseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnstoppableDomainsSetReverse represents a SetReverse event raised by the UnstoppableDomains contract.
type UnstoppableDomainsSetReverse struct {
	Addr    common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSetReverse is a free log retrieval operation binding the contract event 0xeb76a21470988c474a21f690cc28fee1ed511bd812dc3c21fd0f49c5e5d4708a.
//
// Solidity: event SetReverse(address indexed addr, uint256 indexed tokenId)
func (_UnstoppableDomains *UnstoppableDomainsFilterer) FilterSetReverse(opts *bind.FilterOpts, addr []common.Address, tokenId []*big.Int) (*UnstoppableDomainsSetReverseIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _UnstoppableDomains.contract.FilterLogs(opts, "SetReverse", addrRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &UnstoppableDomainsSetReverseIterator{contract: _UnstoppableDomains.contract, event: "SetReverse", logs: logs, sub: sub}, nil
}

// WatchSetReverse is a free log subscription operation binding the contract event 0xeb76a21470988c474a21f690cc28fee1ed511bd812dc3c21fd0f49c5e5d4708a.
//
// Solidity: event SetReverse(address indexed addr, uint256 indexed tokenId)
func (_UnstoppableDomains *UnstoppableDomainsFilterer) WatchSetReverse(opts *bind.WatchOpts, sink chan<- *UnstoppableDomainsSetReverse, addr []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _UnstoppableDomains.contract.WatchLogs(opts, "SetReverse", addrRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnstoppableDomainsSetReverse)
				if err := _UnstoppableDomains.contract.UnpackLog(event, "SetReverse", log); err != nil {
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

// ParseSetReverse is a log parse operation binding the contract event 0xeb76a21470988c474a21f690cc28fee1ed511bd812dc3c21fd0f49c5e5d4708a.
//
// Solidity: event SetReverse(address indexed addr, uint256 indexed tokenId)
func (_UnstoppableDomains *UnstoppableDomainsFilterer) ParseSetReverse(log types.Log) (*UnstoppableDomainsSetReverse, error) {
	event := new(UnstoppableDomainsSetReverse)
	if err := _UnstoppableDomains.contract.UnpackLog(event, "SetReverse", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
