// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zerox

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

// ZeroXMetaData contains all meta data concerning the ZeroX contract.
var ZeroXMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"inputToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"outputToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"inputTokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"outputTokenAmount\",\"type\":\"uint256\"}],\"name\":\"TransformedERC20\",\"type\":\"event\"}]",
}

// ZeroXABI is the input ABI used to generate the binding from.
// Deprecated: Use ZeroXMetaData.ABI instead.
var ZeroXABI = ZeroXMetaData.ABI

// ZeroX is an auto generated Go binding around an Ethereum contract.
type ZeroX struct {
	ZeroXCaller     // Read-only binding to the contract
	ZeroXTransactor // Write-only binding to the contract
	ZeroXFilterer   // Log filterer for contract events
}

// ZeroXCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZeroXCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZeroXTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZeroXTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZeroXFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZeroXFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZeroXSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZeroXSession struct {
	Contract     *ZeroX            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZeroXCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZeroXCallerSession struct {
	Contract *ZeroXCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ZeroXTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZeroXTransactorSession struct {
	Contract     *ZeroXTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZeroXRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZeroXRaw struct {
	Contract *ZeroX // Generic contract binding to access the raw methods on
}

// ZeroXCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZeroXCallerRaw struct {
	Contract *ZeroXCaller // Generic read-only contract binding to access the raw methods on
}

// ZeroXTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZeroXTransactorRaw struct {
	Contract *ZeroXTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZeroX creates a new instance of ZeroX, bound to a specific deployed contract.
func NewZeroX(address common.Address, backend bind.ContractBackend) (*ZeroX, error) {
	contract, err := bindZeroX(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZeroX{ZeroXCaller: ZeroXCaller{contract: contract}, ZeroXTransactor: ZeroXTransactor{contract: contract}, ZeroXFilterer: ZeroXFilterer{contract: contract}}, nil
}

// NewZeroXCaller creates a new read-only instance of ZeroX, bound to a specific deployed contract.
func NewZeroXCaller(address common.Address, caller bind.ContractCaller) (*ZeroXCaller, error) {
	contract, err := bindZeroX(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZeroXCaller{contract: contract}, nil
}

// NewZeroXTransactor creates a new write-only instance of ZeroX, bound to a specific deployed contract.
func NewZeroXTransactor(address common.Address, transactor bind.ContractTransactor) (*ZeroXTransactor, error) {
	contract, err := bindZeroX(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZeroXTransactor{contract: contract}, nil
}

// NewZeroXFilterer creates a new log filterer instance of ZeroX, bound to a specific deployed contract.
func NewZeroXFilterer(address common.Address, filterer bind.ContractFilterer) (*ZeroXFilterer, error) {
	contract, err := bindZeroX(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZeroXFilterer{contract: contract}, nil
}

// bindZeroX binds a generic wrapper to an already deployed contract.
func bindZeroX(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZeroXABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZeroX *ZeroXRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZeroX.Contract.ZeroXCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZeroX *ZeroXRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZeroX.Contract.ZeroXTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZeroX *ZeroXRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZeroX.Contract.ZeroXTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZeroX *ZeroXCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZeroX.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZeroX *ZeroXTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZeroX.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZeroX *ZeroXTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZeroX.Contract.contract.Transact(opts, method, params...)
}

// ZeroXTransformedERC20Iterator is returned from FilterTransformedERC20 and is used to iterate over the raw logs and unpacked data for TransformedERC20 events raised by the ZeroX contract.
type ZeroXTransformedERC20Iterator struct {
	Event *ZeroXTransformedERC20 // Event containing the contract specifics and raw log

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
func (it *ZeroXTransformedERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZeroXTransformedERC20)
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
		it.Event = new(ZeroXTransformedERC20)
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
func (it *ZeroXTransformedERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZeroXTransformedERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZeroXTransformedERC20 represents a TransformedERC20 event raised by the ZeroX contract.
type ZeroXTransformedERC20 struct {
	Taker             common.Address
	InputToken        common.Address
	OutputToken       common.Address
	InputTokenAmount  *big.Int
	OutputTokenAmount *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterTransformedERC20 is a free log retrieval operation binding the contract event 0x0f6672f78a59ba8e5e5b5d38df3ebc67f3c792e2c9259b8d97d7f00dd78ba1b3.
//
// Solidity: event TransformedERC20(address indexed taker, address inputToken, address outputToken, uint256 inputTokenAmount, uint256 outputTokenAmount)
func (_ZeroX *ZeroXFilterer) FilterTransformedERC20(opts *bind.FilterOpts, taker []common.Address) (*ZeroXTransformedERC20Iterator, error) {

	var takerRule []interface{}
	for _, takerItem := range taker {
		takerRule = append(takerRule, takerItem)
	}

	logs, sub, err := _ZeroX.contract.FilterLogs(opts, "TransformedERC20", takerRule)
	if err != nil {
		return nil, err
	}
	return &ZeroXTransformedERC20Iterator{contract: _ZeroX.contract, event: "TransformedERC20", logs: logs, sub: sub}, nil
}

// WatchTransformedERC20 is a free log subscription operation binding the contract event 0x0f6672f78a59ba8e5e5b5d38df3ebc67f3c792e2c9259b8d97d7f00dd78ba1b3.
//
// Solidity: event TransformedERC20(address indexed taker, address inputToken, address outputToken, uint256 inputTokenAmount, uint256 outputTokenAmount)
func (_ZeroX *ZeroXFilterer) WatchTransformedERC20(opts *bind.WatchOpts, sink chan<- *ZeroXTransformedERC20, taker []common.Address) (event.Subscription, error) {

	var takerRule []interface{}
	for _, takerItem := range taker {
		takerRule = append(takerRule, takerItem)
	}

	logs, sub, err := _ZeroX.contract.WatchLogs(opts, "TransformedERC20", takerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZeroXTransformedERC20)
				if err := _ZeroX.contract.UnpackLog(event, "TransformedERC20", log); err != nil {
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

// ParseTransformedERC20 is a log parse operation binding the contract event 0x0f6672f78a59ba8e5e5b5d38df3ebc67f3c792e2c9259b8d97d7f00dd78ba1b3.
//
// Solidity: event TransformedERC20(address indexed taker, address inputToken, address outputToken, uint256 inputTokenAmount, uint256 outputTokenAmount)
func (_ZeroX *ZeroXFilterer) ParseTransformedERC20(log types.Log) (*ZeroXTransformedERC20, error) {
	event := new(ZeroXTransformedERC20)
	if err := _ZeroX.contract.UnpackLog(event, "TransformedERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
