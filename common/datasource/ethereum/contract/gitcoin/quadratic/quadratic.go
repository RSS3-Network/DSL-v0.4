// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package quadratic

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

// QuadraticMetaData contains all meta data concerning the Quadratic contract.
var QuadraticMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"grantAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"applicationIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"roundAddress\",\"type\":\"address\"}],\"name\":\"Voted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"roundAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"encodedVotes\",\"type\":\"bytes[]\"},{\"internalType\":\"address\",\"name\":\"voterAddress\",\"type\":\"address\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// QuadraticABI is the input ABI used to generate the binding from.
// Deprecated: Use QuadraticMetaData.ABI instead.
var QuadraticABI = QuadraticMetaData.ABI

// Quadratic is an auto generated Go binding around an Ethereum contract.
type Quadratic struct {
	QuadraticCaller     // Read-only binding to the contract
	QuadraticTransactor // Write-only binding to the contract
	QuadraticFilterer   // Log filterer for contract events
}

// QuadraticCaller is an auto generated read-only Go binding around an Ethereum contract.
type QuadraticCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuadraticTransactor is an auto generated write-only Go binding around an Ethereum contract.
type QuadraticTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuadraticFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type QuadraticFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuadraticSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type QuadraticSession struct {
	Contract     *Quadratic        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// QuadraticCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type QuadraticCallerSession struct {
	Contract *QuadraticCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// QuadraticTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type QuadraticTransactorSession struct {
	Contract     *QuadraticTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// QuadraticRaw is an auto generated low-level Go binding around an Ethereum contract.
type QuadraticRaw struct {
	Contract *Quadratic // Generic contract binding to access the raw methods on
}

// QuadraticCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type QuadraticCallerRaw struct {
	Contract *QuadraticCaller // Generic read-only contract binding to access the raw methods on
}

// QuadraticTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type QuadraticTransactorRaw struct {
	Contract *QuadraticTransactor // Generic write-only contract binding to access the raw methods on
}

// NewQuadratic creates a new instance of Quadratic, bound to a specific deployed contract.
func NewQuadratic(address common.Address, backend bind.ContractBackend) (*Quadratic, error) {
	contract, err := bindQuadratic(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Quadratic{QuadraticCaller: QuadraticCaller{contract: contract}, QuadraticTransactor: QuadraticTransactor{contract: contract}, QuadraticFilterer: QuadraticFilterer{contract: contract}}, nil
}

// NewQuadraticCaller creates a new read-only instance of Quadratic, bound to a specific deployed contract.
func NewQuadraticCaller(address common.Address, caller bind.ContractCaller) (*QuadraticCaller, error) {
	contract, err := bindQuadratic(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &QuadraticCaller{contract: contract}, nil
}

// NewQuadraticTransactor creates a new write-only instance of Quadratic, bound to a specific deployed contract.
func NewQuadraticTransactor(address common.Address, transactor bind.ContractTransactor) (*QuadraticTransactor, error) {
	contract, err := bindQuadratic(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &QuadraticTransactor{contract: contract}, nil
}

// NewQuadraticFilterer creates a new log filterer instance of Quadratic, bound to a specific deployed contract.
func NewQuadraticFilterer(address common.Address, filterer bind.ContractFilterer) (*QuadraticFilterer, error) {
	contract, err := bindQuadratic(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &QuadraticFilterer{contract: contract}, nil
}

// bindQuadratic binds a generic wrapper to an already deployed contract.
func bindQuadratic(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := QuadraticMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Quadratic *QuadraticRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Quadratic.Contract.QuadraticCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Quadratic *QuadraticRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Quadratic.Contract.QuadraticTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Quadratic *QuadraticRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Quadratic.Contract.QuadraticTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Quadratic *QuadraticCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Quadratic.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Quadratic *QuadraticTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Quadratic.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Quadratic *QuadraticTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Quadratic.Contract.contract.Transact(opts, method, params...)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Quadratic *QuadraticCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Quadratic.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Quadratic *QuadraticSession) VERSION() (string, error) {
	return _Quadratic.Contract.VERSION(&_Quadratic.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Quadratic *QuadraticCallerSession) VERSION() (string, error) {
	return _Quadratic.Contract.VERSION(&_Quadratic.CallOpts)
}

// RoundAddress is a free data retrieval call binding the contract method 0x0b67d925.
//
// Solidity: function roundAddress() view returns(address)
func (_Quadratic *QuadraticCaller) RoundAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Quadratic.contract.Call(opts, &out, "roundAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RoundAddress is a free data retrieval call binding the contract method 0x0b67d925.
//
// Solidity: function roundAddress() view returns(address)
func (_Quadratic *QuadraticSession) RoundAddress() (common.Address, error) {
	return _Quadratic.Contract.RoundAddress(&_Quadratic.CallOpts)
}

// RoundAddress is a free data retrieval call binding the contract method 0x0b67d925.
//
// Solidity: function roundAddress() view returns(address)
func (_Quadratic *QuadraticCallerSession) RoundAddress() (common.Address, error) {
	return _Quadratic.Contract.RoundAddress(&_Quadratic.CallOpts)
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_Quadratic *QuadraticTransactor) Init(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Quadratic.contract.Transact(opts, "init")
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_Quadratic *QuadraticSession) Init() (*types.Transaction, error) {
	return _Quadratic.Contract.Init(&_Quadratic.TransactOpts)
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_Quadratic *QuadraticTransactorSession) Init() (*types.Transaction, error) {
	return _Quadratic.Contract.Init(&_Quadratic.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Quadratic *QuadraticTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Quadratic.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Quadratic *QuadraticSession) Initialize() (*types.Transaction, error) {
	return _Quadratic.Contract.Initialize(&_Quadratic.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Quadratic *QuadraticTransactorSession) Initialize() (*types.Transaction, error) {
	return _Quadratic.Contract.Initialize(&_Quadratic.TransactOpts)
}

// Vote is a paid mutator transaction binding the contract method 0xfc6d4e39.
//
// Solidity: function vote(bytes[] encodedVotes, address voterAddress) payable returns()
func (_Quadratic *QuadraticTransactor) Vote(opts *bind.TransactOpts, encodedVotes [][]byte, voterAddress common.Address) (*types.Transaction, error) {
	return _Quadratic.contract.Transact(opts, "vote", encodedVotes, voterAddress)
}

// Vote is a paid mutator transaction binding the contract method 0xfc6d4e39.
//
// Solidity: function vote(bytes[] encodedVotes, address voterAddress) payable returns()
func (_Quadratic *QuadraticSession) Vote(encodedVotes [][]byte, voterAddress common.Address) (*types.Transaction, error) {
	return _Quadratic.Contract.Vote(&_Quadratic.TransactOpts, encodedVotes, voterAddress)
}

// Vote is a paid mutator transaction binding the contract method 0xfc6d4e39.
//
// Solidity: function vote(bytes[] encodedVotes, address voterAddress) payable returns()
func (_Quadratic *QuadraticTransactorSession) Vote(encodedVotes [][]byte, voterAddress common.Address) (*types.Transaction, error) {
	return _Quadratic.Contract.Vote(&_Quadratic.TransactOpts, encodedVotes, voterAddress)
}

// QuadraticInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Quadratic contract.
type QuadraticInitializedIterator struct {
	Event *QuadraticInitialized // Event containing the contract specifics and raw log

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
func (it *QuadraticInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuadraticInitialized)
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
		it.Event = new(QuadraticInitialized)
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
func (it *QuadraticInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuadraticInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuadraticInitialized represents a Initialized event raised by the Quadratic contract.
type QuadraticInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Quadratic *QuadraticFilterer) FilterInitialized(opts *bind.FilterOpts) (*QuadraticInitializedIterator, error) {

	logs, sub, err := _Quadratic.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &QuadraticInitializedIterator{contract: _Quadratic.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Quadratic *QuadraticFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *QuadraticInitialized) (event.Subscription, error) {

	logs, sub, err := _Quadratic.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuadraticInitialized)
				if err := _Quadratic.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Quadratic *QuadraticFilterer) ParseInitialized(log types.Log) (*QuadraticInitialized, error) {
	event := new(QuadraticInitialized)
	if err := _Quadratic.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QuadraticVotedIterator is returned from FilterVoted and is used to iterate over the raw logs and unpacked data for Voted events raised by the Quadratic contract.
type QuadraticVotedIterator struct {
	Event *QuadraticVoted // Event containing the contract specifics and raw log

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
func (it *QuadraticVotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuadraticVoted)
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
		it.Event = new(QuadraticVoted)
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
func (it *QuadraticVotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuadraticVotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuadraticVoted represents a Voted event raised by the Quadratic contract.
type QuadraticVoted struct {
	Token            common.Address
	Amount           *big.Int
	Voter            common.Address
	GrantAddress     common.Address
	ProjectId        [32]byte
	ApplicationIndex *big.Int
	RoundAddress     common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterVoted is a free log retrieval operation binding the contract event 0x0064caa73f1d59b69adbeb65654b0f0953597994e4241ee2460b560b8d65aaa2.
//
// Solidity: event Voted(address token, uint256 amount, address indexed voter, address grantAddress, bytes32 indexed projectId, uint256 applicationIndex, address indexed roundAddress)
func (_Quadratic *QuadraticFilterer) FilterVoted(opts *bind.FilterOpts, voter []common.Address, projectId [][32]byte, roundAddress []common.Address) (*QuadraticVotedIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}

	var roundAddressRule []interface{}
	for _, roundAddressItem := range roundAddress {
		roundAddressRule = append(roundAddressRule, roundAddressItem)
	}

	logs, sub, err := _Quadratic.contract.FilterLogs(opts, "Voted", voterRule, projectIdRule, roundAddressRule)
	if err != nil {
		return nil, err
	}
	return &QuadraticVotedIterator{contract: _Quadratic.contract, event: "Voted", logs: logs, sub: sub}, nil
}

// WatchVoted is a free log subscription operation binding the contract event 0x0064caa73f1d59b69adbeb65654b0f0953597994e4241ee2460b560b8d65aaa2.
//
// Solidity: event Voted(address token, uint256 amount, address indexed voter, address grantAddress, bytes32 indexed projectId, uint256 applicationIndex, address indexed roundAddress)
func (_Quadratic *QuadraticFilterer) WatchVoted(opts *bind.WatchOpts, sink chan<- *QuadraticVoted, voter []common.Address, projectId [][32]byte, roundAddress []common.Address) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}

	var roundAddressRule []interface{}
	for _, roundAddressItem := range roundAddress {
		roundAddressRule = append(roundAddressRule, roundAddressItem)
	}

	logs, sub, err := _Quadratic.contract.WatchLogs(opts, "Voted", voterRule, projectIdRule, roundAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuadraticVoted)
				if err := _Quadratic.contract.UnpackLog(event, "Voted", log); err != nil {
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

// ParseVoted is a log parse operation binding the contract event 0x0064caa73f1d59b69adbeb65654b0f0953597994e4241ee2460b560b8d65aaa2.
//
// Solidity: event Voted(address token, uint256 amount, address indexed voter, address grantAddress, bytes32 indexed projectId, uint256 applicationIndex, address indexed roundAddress)
func (_Quadratic *QuadraticFilterer) ParseVoted(log types.Log) (*QuadraticVoted, error) {
	event := new(QuadraticVoted)
	if err := _Quadratic.contract.UnpackLog(event, "Voted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
