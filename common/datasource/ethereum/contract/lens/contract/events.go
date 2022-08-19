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
)

// EventsMetaData contains all meta data concerning the Events contract.
var EventsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"BaseInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collectModule\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"whitelisted\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"CollectModuleWhitelisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pubId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collectNFT\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"CollectNFTDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pubId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"CollectNFTInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pubId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"collectNFTId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"CollectNFTTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collector\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pubId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rootProfileId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rootPubId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"collectModuleData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Collected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pubId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"contentURI\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"profileIdPointed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pubIdPointed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"referenceModuleData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collectModule\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"collectModuleReturnData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"referenceModule\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"referenceModuleReturnData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"CommentCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"DefaultProfileSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dispatcher\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"DispatcherSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldEmergencyAdmin\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newEmergencyAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"EmergencyAdminSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"moduleGlobals\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"FeeModuleBaseConstructed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"followModule\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"followModuleReturnData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"FollowModuleSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"followModule\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"whitelisted\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"FollowModuleWhitelisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newPower\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"FollowNFTDelegatedPowerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"followNFT\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"FollowNFTDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"FollowNFTInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"followNFTId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"FollowNFTTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"followNFTURI\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"FollowNFTURISet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"follower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"profileIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"followModuleDatas\",\"type\":\"bytes[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Followed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"bool[]\",\"name\":\"approved\",\"type\":\"bool[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"FollowsApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"profileIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"bool[]\",\"name\":\"enabled\",\"type\":\"bool[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"FollowsToggled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"prevGovernance\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newGovernance\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"GovernanceSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pubId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"profileIdPointed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pubIdPointed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"referenceModuleData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"referenceModule\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"referenceModuleReturnData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"MirrorCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"hub\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ModuleBaseConstructed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"prevWhitelisted\",\"type\":\"bool\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"whitelisted\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ModuleGlobalsCurrencyWhitelisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"prevGovernance\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newGovernance\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ModuleGlobalsGovernanceSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"prevTreasuryFee\",\"type\":\"uint16\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"newTreasuryFee\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ModuleGlobalsTreasuryFeeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"prevTreasury\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newTreasury\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ModuleGlobalsTreasurySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pubId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"contentURI\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collectModule\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"collectModuleReturnData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"referenceModule\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"referenceModuleReturnData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"PostCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"handle\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"imageURI\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"followModule\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"followModuleReturnData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"followNFTURI\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ProfileCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"profileCreator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"whitelisted\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ProfileCreatorWhitelisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"imageURI\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ProfileImageURISet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ProfileMetadataSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"referenceModule\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"whitelisted\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ReferenceModuleWhitelisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumDataTypes.ProtocolState\",\"name\":\"prevState\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"enumDataTypes.ProtocolState\",\"name\":\"newState\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"StateSet\",\"type\":\"event\"}]",
}

// EventsABI is the input ABI used to generate the binding from.
// Deprecated: Use EventsMetaData.ABI instead.
var EventsABI = EventsMetaData.ABI

// Events is an auto generated Go binding around an Ethereum contract.
type Events struct {
	EventsCaller     // Read-only binding to the contract
	EventsTransactor // Write-only binding to the contract
	EventsFilterer   // Log filterer for contract events
}

// EventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type EventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EventsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EventsSession struct {
	Contract     *Events           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EventsCallerSession struct {
	Contract *EventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// EventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EventsTransactorSession struct {
	Contract     *EventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type EventsRaw struct {
	Contract *Events // Generic contract binding to access the raw methods on
}

// EventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EventsCallerRaw struct {
	Contract *EventsCaller // Generic read-only contract binding to access the raw methods on
}

// EventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EventsTransactorRaw struct {
	Contract *EventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEvents creates a new instance of Events, bound to a specific deployed contract.
func NewEvents(address common.Address, backend bind.ContractBackend) (*Events, error) {
	contract, err := bindEvents(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Events{EventsCaller: EventsCaller{contract: contract}, EventsTransactor: EventsTransactor{contract: contract}, EventsFilterer: EventsFilterer{contract: contract}}, nil
}

// NewEventsCaller creates a new read-only instance of Events, bound to a specific deployed contract.
func NewEventsCaller(address common.Address, caller bind.ContractCaller) (*EventsCaller, error) {
	contract, err := bindEvents(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EventsCaller{contract: contract}, nil
}

// NewEventsTransactor creates a new write-only instance of Events, bound to a specific deployed contract.
func NewEventsTransactor(address common.Address, transactor bind.ContractTransactor) (*EventsTransactor, error) {
	contract, err := bindEvents(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EventsTransactor{contract: contract}, nil
}

// NewEventsFilterer creates a new log filterer instance of Events, bound to a specific deployed contract.
func NewEventsFilterer(address common.Address, filterer bind.ContractFilterer) (*EventsFilterer, error) {
	contract, err := bindEvents(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EventsFilterer{contract: contract}, nil
}

// bindEvents binds a generic wrapper to an already deployed contract.
func bindEvents(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EventsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Events *EventsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Events.Contract.EventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Events *EventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Events.Contract.EventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Events *EventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Events.Contract.EventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Events *EventsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Events.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Events *EventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Events.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Events *EventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Events.Contract.contract.Transact(opts, method, params...)
}

// EventsBaseInitializedIterator is returned from FilterBaseInitialized and is used to iterate over the raw logs and unpacked data for BaseInitialized events raised by the Events contract.
type EventsBaseInitializedIterator struct {
	Event *EventsBaseInitialized // Event containing the contract specifics and raw log

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
func (it *EventsBaseInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventsBaseInitialized)
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
		it.Event = new(EventsBaseInitialized)
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
func (it *EventsBaseInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventsBaseInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventsBaseInitialized represents a BaseInitialized event raised by the Events contract.
type EventsBaseInitialized struct {
	Name      string
	Symbol    string
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBaseInitialized is a free log retrieval operation binding the contract event 0x414cd0b34676984f09a5f76ce9718d4062e50283abe0e7e274a9a5b4e0c99c30.
//
// Solidity: event BaseInitialized(string name, string symbol, uint256 timestamp)
func (_Events *EventsFilterer) FilterBaseInitialized(opts *bind.FilterOpts) (*EventsBaseInitializedIterator, error) {

	logs, sub, err := _Events.contract.FilterLogs(opts, "BaseInitialized")
	if err != nil {
		return nil, err
	}
	return &EventsBaseInitializedIterator{contract: _Events.contract, event: "BaseInitialized", logs: logs, sub: sub}, nil
}

// WatchBaseInitialized is a free log subscription operation binding the contract event 0x414cd0b34676984f09a5f76ce9718d4062e50283abe0e7e274a9a5b4e0c99c30.
//
// Solidity: event BaseInitialized(string name, string symbol, uint256 timestamp)
func (_Events *EventsFilterer) WatchBaseInitialized(opts *bind.WatchOpts, sink chan<- *EventsBaseInitialized) (event.Subscription, error) {

	logs, sub, err := _Events.contract.WatchLogs(opts, "BaseInitialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventsBaseInitialized)
				if err := _Events.contract.UnpackLog(event, "BaseInitialized", log); err != nil {
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

// ParseBaseInitialized is a log parse operation binding the contract event 0x414cd0b34676984f09a5f76ce9718d4062e50283abe0e7e274a9a5b4e0c99c30.
//
// Solidity: event BaseInitialized(string name, string symbol, uint256 timestamp)
func (_Events *EventsFilterer) ParseBaseInitialized(log types.Log) (*EventsBaseInitialized, error) {
	event := new(EventsBaseInitialized)
	if err := _Events.contract.UnpackLog(event, "BaseInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventsCollectModuleWhitelistedIterator is returned from FilterCollectModuleWhitelisted and is used to iterate over the raw logs and unpacked data for CollectModuleWhitelisted events raised by the Events contract.
type EventsCollectModuleWhitelistedIterator struct {
	Event *EventsCollectModuleWhitelisted // Event containing the contract specifics and raw log

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
func (it *EventsCollectModuleWhitelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventsCollectModuleWhitelisted)
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
		it.Event = new(EventsCollectModuleWhitelisted)
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
func (it *EventsCollectModuleWhitelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventsCollectModuleWhitelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventsCollectModuleWhitelisted represents a CollectModuleWhitelisted event raised by the Events contract.
type EventsCollectModuleWhitelisted struct {
	CollectModule common.Address
	Whitelisted   bool
	Timestamp     *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCollectModuleWhitelisted is a free log retrieval operation binding the contract event 0x6cc19a794d6a439023150cd58748eed4353190c0bb060d2e6250e2df4a68b673.
//
// Solidity: event CollectModuleWhitelisted(address indexed collectModule, bool indexed whitelisted, uint256 timestamp)
func (_Events *EventsFilterer) FilterCollectModuleWhitelisted(opts *bind.FilterOpts, collectModule []common.Address, whitelisted []bool) (*EventsCollectModuleWhitelistedIterator, error) {

	var collectModuleRule []interface{}
	for _, collectModuleItem := range collectModule {
		collectModuleRule = append(collectModuleRule, collectModuleItem)
	}
	var whitelistedRule []interface{}
	for _, whitelistedItem := range whitelisted {
		whitelistedRule = append(whitelistedRule, whitelistedItem)
	}

	logs, sub, err := _Events.contract.FilterLogs(opts, "CollectModuleWhitelisted", collectModuleRule, whitelistedRule)
	if err != nil {
		return nil, err
	}
	return &EventsCollectModuleWhitelistedIterator{contract: _Events.contract, event: "CollectModuleWhitelisted", logs: logs, sub: sub}, nil
}

// WatchCollectModuleWhitelisted is a free log subscription operation binding the contract event 0x6cc19a794d6a439023150cd58748eed4353190c0bb060d2e6250e2df4a68b673.
//
// Solidity: event CollectModuleWhitelisted(address indexed collectModule, bool indexed whitelisted, uint256 timestamp)
func (_Events *EventsFilterer) WatchCollectModuleWhitelisted(opts *bind.WatchOpts, sink chan<- *EventsCollectModuleWhitelisted, collectModule []common.Address, whitelisted []bool) (event.Subscription, error) {

	var collectModuleRule []interface{}
	for _, collectModuleItem := range collectModule {
		collectModuleRule = append(collectModuleRule, collectModuleItem)
	}
	var whitelistedRule []interface{}
	for _, whitelistedItem := range whitelisted {
		whitelistedRule = append(whitelistedRule, whitelistedItem)
	}

	logs, sub, err := _Events.contract.WatchLogs(opts, "CollectModuleWhitelisted", collectModuleRule, whitelistedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventsCollectModuleWhitelisted)
				if err := _Events.contract.UnpackLog(event, "CollectModuleWhitelisted", log); err != nil {
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

// ParseCollectModuleWhitelisted is a log parse operation binding the contract event 0x6cc19a794d6a439023150cd58748eed4353190c0bb060d2e6250e2df4a68b673.
//
// Solidity: event CollectModuleWhitelisted(address indexed collectModule, bool indexed whitelisted, uint256 timestamp)
func (_Events *EventsFilterer) ParseCollectModuleWhitelisted(log types.Log) (*EventsCollectModuleWhitelisted, error) {
	event := new(EventsCollectModuleWhitelisted)
	if err := _Events.contract.UnpackLog(event, "CollectModuleWhitelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventsCollectNFTDeployedIterator is returned from FilterCollectNFTDeployed and is used to iterate over the raw logs and unpacked data for CollectNFTDeployed events raised by the Events contract.
type EventsCollectNFTDeployedIterator struct {
	Event *EventsCollectNFTDeployed // Event containing the contract specifics and raw log

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
func (it *EventsCollectNFTDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventsCollectNFTDeployed)
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
		it.Event = new(EventsCollectNFTDeployed)
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
func (it *EventsCollectNFTDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventsCollectNFTDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventsCollectNFTDeployed represents a CollectNFTDeployed event raised by the Events contract.
type EventsCollectNFTDeployed struct {
	ProfileId  *big.Int
	PubId      *big.Int
	CollectNFT common.Address
	Timestamp  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCollectNFTDeployed is a free log retrieval operation binding the contract event 0x0b227b550ffed48af813b32e246f787e99581ee13206ba8f9d90d63615269b3f.
//
// Solidity: event CollectNFTDeployed(uint256 indexed profileId, uint256 indexed pubId, address indexed collectNFT, uint256 timestamp)
func (_Events *EventsFilterer) FilterCollectNFTDeployed(opts *bind.FilterOpts, profileId []*big.Int, pubId []*big.Int, collectNFT []common.Address) (*EventsCollectNFTDeployedIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var pubIdRule []interface{}
	for _, pubIdItem := range pubId {
		pubIdRule = append(pubIdRule, pubIdItem)
	}
	var collectNFTRule []interface{}
	for _, collectNFTItem := range collectNFT {
		collectNFTRule = append(collectNFTRule, collectNFTItem)
	}

	logs, sub, err := _Events.contract.FilterLogs(opts, "CollectNFTDeployed", profileIdRule, pubIdRule, collectNFTRule)
	if err != nil {
		return nil, err
	}
	return &EventsCollectNFTDeployedIterator{contract: _Events.contract, event: "CollectNFTDeployed", logs: logs, sub: sub}, nil
}

// WatchCollectNFTDeployed is a free log subscription operation binding the contract event 0x0b227b550ffed48af813b32e246f787e99581ee13206ba8f9d90d63615269b3f.
//
// Solidity: event CollectNFTDeployed(uint256 indexed profileId, uint256 indexed pubId, address indexed collectNFT, uint256 timestamp)
func (_Events *EventsFilterer) WatchCollectNFTDeployed(opts *bind.WatchOpts, sink chan<- *EventsCollectNFTDeployed, profileId []*big.Int, pubId []*big.Int, collectNFT []common.Address) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var pubIdRule []interface{}
	for _, pubIdItem := range pubId {
		pubIdRule = append(pubIdRule, pubIdItem)
	}
	var collectNFTRule []interface{}
	for _, collectNFTItem := range collectNFT {
		collectNFTRule = append(collectNFTRule, collectNFTItem)
	}

	logs, sub, err := _Events.contract.WatchLogs(opts, "CollectNFTDeployed", profileIdRule, pubIdRule, collectNFTRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventsCollectNFTDeployed)
				if err := _Events.contract.UnpackLog(event, "CollectNFTDeployed", log); err != nil {
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

// ParseCollectNFTDeployed is a log parse operation binding the contract event 0x0b227b550ffed48af813b32e246f787e99581ee13206ba8f9d90d63615269b3f.
//
// Solidity: event CollectNFTDeployed(uint256 indexed profileId, uint256 indexed pubId, address indexed collectNFT, uint256 timestamp)
func (_Events *EventsFilterer) ParseCollectNFTDeployed(log types.Log) (*EventsCollectNFTDeployed, error) {
	event := new(EventsCollectNFTDeployed)
	if err := _Events.contract.UnpackLog(event, "CollectNFTDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventsCollectNFTInitializedIterator is returned from FilterCollectNFTInitialized and is used to iterate over the raw logs and unpacked data for CollectNFTInitialized events raised by the Events contract.
type EventsCollectNFTInitializedIterator struct {
	Event *EventsCollectNFTInitialized // Event containing the contract specifics and raw log

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
func (it *EventsCollectNFTInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventsCollectNFTInitialized)
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
		it.Event = new(EventsCollectNFTInitialized)
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
func (it *EventsCollectNFTInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventsCollectNFTInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventsCollectNFTInitialized represents a CollectNFTInitialized event raised by the Events contract.
type EventsCollectNFTInitialized struct {
	ProfileId *big.Int
	PubId     *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCollectNFTInitialized is a free log retrieval operation binding the contract event 0x898a2dec95856255977a0fb48cebc30051d50c0d8d33f93dea1e3ddb2e342442.
//
// Solidity: event CollectNFTInitialized(uint256 indexed profileId, uint256 indexed pubId, uint256 timestamp)
func (_Events *EventsFilterer) FilterCollectNFTInitialized(opts *bind.FilterOpts, profileId []*big.Int, pubId []*big.Int) (*EventsCollectNFTInitializedIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var pubIdRule []interface{}
	for _, pubIdItem := range pubId {
		pubIdRule = append(pubIdRule, pubIdItem)
	}

	logs, sub, err := _Events.contract.FilterLogs(opts, "CollectNFTInitialized", profileIdRule, pubIdRule)
	if err != nil {
		return nil, err
	}
	return &EventsCollectNFTInitializedIterator{contract: _Events.contract, event: "CollectNFTInitialized", logs: logs, sub: sub}, nil
}

// WatchCollectNFTInitialized is a free log subscription operation binding the contract event 0x898a2dec95856255977a0fb48cebc30051d50c0d8d33f93dea1e3ddb2e342442.
//
// Solidity: event CollectNFTInitialized(uint256 indexed profileId, uint256 indexed pubId, uint256 timestamp)
func (_Events *EventsFilterer) WatchCollectNFTInitialized(opts *bind.WatchOpts, sink chan<- *EventsCollectNFTInitialized, profileId []*big.Int, pubId []*big.Int) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var pubIdRule []interface{}
	for _, pubIdItem := range pubId {
		pubIdRule = append(pubIdRule, pubIdItem)
	}

	logs, sub, err := _Events.contract.WatchLogs(opts, "CollectNFTInitialized", profileIdRule, pubIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventsCollectNFTInitialized)
				if err := _Events.contract.UnpackLog(event, "CollectNFTInitialized", log); err != nil {
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

// ParseCollectNFTInitialized is a log parse operation binding the contract event 0x898a2dec95856255977a0fb48cebc30051d50c0d8d33f93dea1e3ddb2e342442.
//
// Solidity: event CollectNFTInitialized(uint256 indexed profileId, uint256 indexed pubId, uint256 timestamp)
func (_Events *EventsFilterer) ParseCollectNFTInitialized(log types.Log) (*EventsCollectNFTInitialized, error) {
	event := new(EventsCollectNFTInitialized)
	if err := _Events.contract.UnpackLog(event, "CollectNFTInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventsCollectNFTTransferredIterator is returned from FilterCollectNFTTransferred and is used to iterate over the raw logs and unpacked data for CollectNFTTransferred events raised by the Events contract.
type EventsCollectNFTTransferredIterator struct {
	Event *EventsCollectNFTTransferred // Event containing the contract specifics and raw log

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
func (it *EventsCollectNFTTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventsCollectNFTTransferred)
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
		it.Event = new(EventsCollectNFTTransferred)
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
func (it *EventsCollectNFTTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventsCollectNFTTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventsCollectNFTTransferred represents a CollectNFTTransferred event raised by the Events contract.
type EventsCollectNFTTransferred struct {
	ProfileId    *big.Int
	PubId        *big.Int
	CollectNFTId *big.Int
	From         common.Address
	To           common.Address
	Timestamp    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCollectNFTTransferred is a free log retrieval operation binding the contract event 0x68edb7ec2c37d21b3b72233960b487f2966f4ac82b7430d39f24d1f8d6f99106.
//
// Solidity: event CollectNFTTransferred(uint256 indexed profileId, uint256 indexed pubId, uint256 indexed collectNFTId, address from, address to, uint256 timestamp)
func (_Events *EventsFilterer) FilterCollectNFTTransferred(opts *bind.FilterOpts, profileId []*big.Int, pubId []*big.Int, collectNFTId []*big.Int) (*EventsCollectNFTTransferredIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var pubIdRule []interface{}
	for _, pubIdItem := range pubId {
		pubIdRule = append(pubIdRule, pubIdItem)
	}
	var collectNFTIdRule []interface{}
	for _, collectNFTIdItem := range collectNFTId {
		collectNFTIdRule = append(collectNFTIdRule, collectNFTIdItem)
	}

	logs, sub, err := _Events.contract.FilterLogs(opts, "CollectNFTTransferred", profileIdRule, pubIdRule, collectNFTIdRule)
	if err != nil {
		return nil, err
	}
	return &EventsCollectNFTTransferredIterator{contract: _Events.contract, event: "CollectNFTTransferred", logs: logs, sub: sub}, nil
}

// WatchCollectNFTTransferred is a free log subscription operation binding the contract event 0x68edb7ec2c37d21b3b72233960b487f2966f4ac82b7430d39f24d1f8d6f99106.
//
// Solidity: event CollectNFTTransferred(uint256 indexed profileId, uint256 indexed pubId, uint256 indexed collectNFTId, address from, address to, uint256 timestamp)
func (_Events *EventsFilterer) WatchCollectNFTTransferred(opts *bind.WatchOpts, sink chan<- *EventsCollectNFTTransferred, profileId []*big.Int, pubId []*big.Int, collectNFTId []*big.Int) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var pubIdRule []interface{}
	for _, pubIdItem := range pubId {
		pubIdRule = append(pubIdRule, pubIdItem)
	}
	var collectNFTIdRule []interface{}
	for _, collectNFTIdItem := range collectNFTId {
		collectNFTIdRule = append(collectNFTIdRule, collectNFTIdItem)
	}

	logs, sub, err := _Events.contract.WatchLogs(opts, "CollectNFTTransferred", profileIdRule, pubIdRule, collectNFTIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventsCollectNFTTransferred)
				if err := _Events.contract.UnpackLog(event, "CollectNFTTransferred", log); err != nil {
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

// ParseCollectNFTTransferred is a log parse operation binding the contract event 0x68edb7ec2c37d21b3b72233960b487f2966f4ac82b7430d39f24d1f8d6f99106.
//
// Solidity: event CollectNFTTransferred(uint256 indexed profileId, uint256 indexed pubId, uint256 indexed collectNFTId, address from, address to, uint256 timestamp)
func (_Events *EventsFilterer) ParseCollectNFTTransferred(log types.Log) (*EventsCollectNFTTransferred, error) {
	event := new(EventsCollectNFTTransferred)
	if err := _Events.contract.UnpackLog(event, "CollectNFTTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventsCollectedIterator is returned from FilterCollected and is used to iterate over the raw logs and unpacked data for Collected events raised by the Events contract.
type EventsCollectedIterator struct {
	Event *EventsCollected // Event containing the contract specifics and raw log

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
func (it *EventsCollectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventsCollected)
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
		it.Event = new(EventsCollected)
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
func (it *EventsCollectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventsCollectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventsCollected represents a Collected event raised by the Events contract.
type EventsCollected struct {
	Collector         common.Address
	ProfileId         *big.Int
	PubId             *big.Int
	RootProfileId     *big.Int
	RootPubId         *big.Int
	CollectModuleData []byte
	Timestamp         *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterCollected is a free log retrieval operation binding the contract event 0xed39bf0d9afa849610b901c9d7f4d00751ba20de2db023428065bec153833218.
//
// Solidity: event Collected(address indexed collector, uint256 indexed profileId, uint256 indexed pubId, uint256 rootProfileId, uint256 rootPubId, bytes collectModuleData, uint256 timestamp)
func (_Events *EventsFilterer) FilterCollected(opts *bind.FilterOpts, collector []common.Address, profileId []*big.Int, pubId []*big.Int) (*EventsCollectedIterator, error) {

	var collectorRule []interface{}
	for _, collectorItem := range collector {
		collectorRule = append(collectorRule, collectorItem)
	}
	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var pubIdRule []interface{}
	for _, pubIdItem := range pubId {
		pubIdRule = append(pubIdRule, pubIdItem)
	}

	logs, sub, err := _Events.contract.FilterLogs(opts, "Collected", collectorRule, profileIdRule, pubIdRule)
	if err != nil {
		return nil, err
	}
	return &EventsCollectedIterator{contract: _Events.contract, event: "Collected", logs: logs, sub: sub}, nil
}

// WatchCollected is a free log subscription operation binding the contract event 0xed39bf0d9afa849610b901c9d7f4d00751ba20de2db023428065bec153833218.
//
// Solidity: event Collected(address indexed collector, uint256 indexed profileId, uint256 indexed pubId, uint256 rootProfileId, uint256 rootPubId, bytes collectModuleData, uint256 timestamp)
func (_Events *EventsFilterer) WatchCollected(opts *bind.WatchOpts, sink chan<- *EventsCollected, collector []common.Address, profileId []*big.Int, pubId []*big.Int) (event.Subscription, error) {

	var collectorRule []interface{}
	for _, collectorItem := range collector {
		collectorRule = append(collectorRule, collectorItem)
	}
	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var pubIdRule []interface{}
	for _, pubIdItem := range pubId {
		pubIdRule = append(pubIdRule, pubIdItem)
	}

	logs, sub, err := _Events.contract.WatchLogs(opts, "Collected", collectorRule, profileIdRule, pubIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventsCollected)
				if err := _Events.contract.UnpackLog(event, "Collected", log); err != nil {
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

// ParseCollected is a log parse operation binding the contract event 0xed39bf0d9afa849610b901c9d7f4d00751ba20de2db023428065bec153833218.
//
// Solidity: event Collected(address indexed collector, uint256 indexed profileId, uint256 indexed pubId, uint256 rootProfileId, uint256 rootPubId, bytes collectModuleData, uint256 timestamp)
func (_Events *EventsFilterer) ParseCollected(log types.Log) (*EventsCollected, error) {
	event := new(EventsCollected)
	if err := _Events.contract.UnpackLog(event, "Collected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventsCommentCreatedIterator is returned from FilterCommentCreated and is used to iterate over the raw logs and unpacked data for CommentCreated events raised by the Events contract.
type EventsCommentCreatedIterator struct {
	Event *EventsCommentCreated // Event containing the contract specifics and raw log

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
func (it *EventsCommentCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventsCommentCreated)
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
		it.Event = new(EventsCommentCreated)
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
func (it *EventsCommentCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventsCommentCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventsCommentCreated represents a CommentCreated event raised by the Events contract.
type EventsCommentCreated struct {
	ProfileId                 *big.Int
	PubId                     *big.Int
	ContentURI                string
	ProfileIdPointed          *big.Int
	PubIdPointed              *big.Int
	ReferenceModuleData       []byte
	CollectModule             common.Address
	CollectModuleReturnData   []byte
	ReferenceModule           common.Address
	ReferenceModuleReturnData []byte
	Timestamp                 *big.Int
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterCommentCreated is a free log retrieval operation binding the contract event 0x7b4d1aa33773161799847429e4fbf29f56dbf1a3fe815f5070231cbfba402c37.
//
// Solidity: event CommentCreated(uint256 indexed profileId, uint256 indexed pubId, string contentURI, uint256 profileIdPointed, uint256 pubIdPointed, bytes referenceModuleData, address collectModule, bytes collectModuleReturnData, address referenceModule, bytes referenceModuleReturnData, uint256 timestamp)
func (_Events *EventsFilterer) FilterCommentCreated(opts *bind.FilterOpts, profileId []*big.Int, pubId []*big.Int) (*EventsCommentCreatedIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var pubIdRule []interface{}
	for _, pubIdItem := range pubId {
		pubIdRule = append(pubIdRule, pubIdItem)
	}

	logs, sub, err := _Events.contract.FilterLogs(opts, "CommentCreated", profileIdRule, pubIdRule)
	if err != nil {
		return nil, err
	}
	return &EventsCommentCreatedIterator{contract: _Events.contract, event: "CommentCreated", logs: logs, sub: sub}, nil
}

// WatchCommentCreated is a free log subscription operation binding the contract event 0x7b4d1aa33773161799847429e4fbf29f56dbf1a3fe815f5070231cbfba402c37.
//
// Solidity: event CommentCreated(uint256 indexed profileId, uint256 indexed pubId, string contentURI, uint256 profileIdPointed, uint256 pubIdPointed, bytes referenceModuleData, address collectModule, bytes collectModuleReturnData, address referenceModule, bytes referenceModuleReturnData, uint256 timestamp)
func (_Events *EventsFilterer) WatchCommentCreated(opts *bind.WatchOpts, sink chan<- *EventsCommentCreated, profileId []*big.Int, pubId []*big.Int) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var pubIdRule []interface{}
	for _, pubIdItem := range pubId {
		pubIdRule = append(pubIdRule, pubIdItem)
	}

	logs, sub, err := _Events.contract.WatchLogs(opts, "CommentCreated", profileIdRule, pubIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventsCommentCreated)
				if err := _Events.contract.UnpackLog(event, "CommentCreated", log); err != nil {
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

// ParseCommentCreated is a log parse operation binding the contract event 0x7b4d1aa33773161799847429e4fbf29f56dbf1a3fe815f5070231cbfba402c37.
//
// Solidity: event CommentCreated(uint256 indexed profileId, uint256 indexed pubId, string contentURI, uint256 profileIdPointed, uint256 pubIdPointed, bytes referenceModuleData, address collectModule, bytes collectModuleReturnData, address referenceModule, bytes referenceModuleReturnData, uint256 timestamp)
func (_Events *EventsFilterer) ParseCommentCreated(log types.Log) (*EventsCommentCreated, error) {
	event := new(EventsCommentCreated)
	if err := _Events.contract.UnpackLog(event, "CommentCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventsDefaultProfileSetIterator is returned from FilterDefaultProfileSet and is used to iterate over the raw logs and unpacked data for DefaultProfileSet events raised by the Events contract.
type EventsDefaultProfileSetIterator struct {
	Event *EventsDefaultProfileSet // Event containing the contract specifics and raw log

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
func (it *EventsDefaultProfileSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventsDefaultProfileSet)
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
		it.Event = new(EventsDefaultProfileSet)
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
func (it *EventsDefaultProfileSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventsDefaultProfileSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventsDefaultProfileSet represents a DefaultProfileSet event raised by the Events contract.
type EventsDefaultProfileSet struct {
	Wallet    common.Address
	ProfileId *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDefaultProfileSet is a free log retrieval operation binding the contract event 0x0afd7c479e8bc7dcdb856b3cc27d2332dfe1f018fde574ea124919ddcae8a933.
//
// Solidity: event DefaultProfileSet(address indexed wallet, uint256 indexed profileId, uint256 timestamp)
func (_Events *EventsFilterer) FilterDefaultProfileSet(opts *bind.FilterOpts, wallet []common.Address, profileId []*big.Int) (*EventsDefaultProfileSetIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _Events.contract.FilterLogs(opts, "DefaultProfileSet", walletRule, profileIdRule)
	if err != nil {
		return nil, err
	}
	return &EventsDefaultProfileSetIterator{contract: _Events.contract, event: "DefaultProfileSet", logs: logs, sub: sub}, nil
}

// WatchDefaultProfileSet is a free log subscription operation binding the contract event 0x0afd7c479e8bc7dcdb856b3cc27d2332dfe1f018fde574ea124919ddcae8a933.
//
// Solidity: event DefaultProfileSet(address indexed wallet, uint256 indexed profileId, uint256 timestamp)
func (_Events *EventsFilterer) WatchDefaultProfileSet(opts *bind.WatchOpts, sink chan<- *EventsDefaultProfileSet, wallet []common.Address, profileId []*big.Int) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _Events.contract.WatchLogs(opts, "DefaultProfileSet", walletRule, profileIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventsDefaultProfileSet)
				if err := _Events.contract.UnpackLog(event, "DefaultProfileSet", log); err != nil {
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

// ParseDefaultProfileSet is a log parse operation binding the contract event 0x0afd7c479e8bc7dcdb856b3cc27d2332dfe1f018fde574ea124919ddcae8a933.
//
// Solidity: event DefaultProfileSet(address indexed wallet, uint256 indexed profileId, uint256 timestamp)
func (_Events *EventsFilterer) ParseDefaultProfileSet(log types.Log) (*EventsDefaultProfileSet, error) {
	event := new(EventsDefaultProfileSet)
	if err := _Events.contract.UnpackLog(event, "DefaultProfileSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventsDispatcherSetIterator is returned from FilterDispatcherSet and is used to iterate over the raw logs and unpacked data for DispatcherSet events raised by the Events contract.
type EventsDispatcherSetIterator struct {
	Event *EventsDispatcherSet // Event containing the contract specifics and raw log

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
func (it *EventsDispatcherSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventsDispatcherSet)
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
		it.Event = new(EventsDispatcherSet)
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
func (it *EventsDispatcherSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventsDispatcherSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventsDispatcherSet represents a DispatcherSet event raised by the Events contract.
type EventsDispatcherSet struct {
	ProfileId  *big.Int
	Dispatcher common.Address
	Timestamp  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDispatcherSet is a free log retrieval operation binding the contract event 0x22baaec4952f35f59e45bd2ddb287e1ccc6d319375770c09428eb8f8d604e065.
//
// Solidity: event DispatcherSet(uint256 indexed profileId, address indexed dispatcher, uint256 timestamp)
func (_Events *EventsFilterer) FilterDispatcherSet(opts *bind.FilterOpts, profileId []*big.Int, dispatcher []common.Address) (*EventsDispatcherSetIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var dispatcherRule []interface{}
	for _, dispatcherItem := range dispatcher {
		dispatcherRule = append(dispatcherRule, dispatcherItem)
	}

	logs, sub, err := _Events.contract.FilterLogs(opts, "DispatcherSet", profileIdRule, dispatcherRule)
	if err != nil {
		return nil, err
	}
	return &EventsDispatcherSetIterator{contract: _Events.contract, event: "DispatcherSet", logs: logs, sub: sub}, nil
}

// WatchDispatcherSet is a free log subscription operation binding the contract event 0x22baaec4952f35f59e45bd2ddb287e1ccc6d319375770c09428eb8f8d604e065.
//
// Solidity: event DispatcherSet(uint256 indexed profileId, address indexed dispatcher, uint256 timestamp)
func (_Events *EventsFilterer) WatchDispatcherSet(opts *bind.WatchOpts, sink chan<- *EventsDispatcherSet, profileId []*big.Int, dispatcher []common.Address) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var dispatcherRule []interface{}
	for _, dispatcherItem := range dispatcher {
		dispatcherRule = append(dispatcherRule, dispatcherItem)
	}

	logs, sub, err := _Events.contract.WatchLogs(opts, "DispatcherSet", profileIdRule, dispatcherRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventsDispatcherSet)
				if err := _Events.contract.UnpackLog(event, "DispatcherSet", log); err != nil {
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

// ParseDispatcherSet is a log parse operation binding the contract event 0x22baaec4952f35f59e45bd2ddb287e1ccc6d319375770c09428eb8f8d604e065.
//
// Solidity: event DispatcherSet(uint256 indexed profileId, address indexed dispatcher, uint256 timestamp)
func (_Events *EventsFilterer) ParseDispatcherSet(log types.Log) (*EventsDispatcherSet, error) {
	event := new(EventsDispatcherSet)
	if err := _Events.contract.UnpackLog(event, "DispatcherSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventsEmergencyAdminSetIterator is returned from FilterEmergencyAdminSet and is used to iterate over the raw logs and unpacked data for EmergencyAdminSet events raised by the Events contract.
type EventsEmergencyAdminSetIterator struct {
	Event *EventsEmergencyAdminSet // Event containing the contract specifics and raw log

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
func (it *EventsEmergencyAdminSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventsEmergencyAdminSet)
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
		it.Event = new(EventsEmergencyAdminSet)
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
func (it *EventsEmergencyAdminSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventsEmergencyAdminSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventsEmergencyAdminSet represents a EmergencyAdminSet event raised by the Events contract.
type EventsEmergencyAdminSet struct {
	Caller            common.Address
	OldEmergencyAdmin common.Address
	NewEmergencyAdmin common.Address
	Timestamp         *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterEmergencyAdminSet is a free log retrieval operation binding the contract event 0x676c0801b0f400762e958ee31cfbb10870e70786f6761f57c8647e766b0db3d9.
//
// Solidity: event EmergencyAdminSet(address indexed caller, address indexed oldEmergencyAdmin, address indexed newEmergencyAdmin, uint256 timestamp)
func (_Events *EventsFilterer) FilterEmergencyAdminSet(opts *bind.FilterOpts, caller []common.Address, oldEmergencyAdmin []common.Address, newEmergencyAdmin []common.Address) (*EventsEmergencyAdminSetIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var oldEmergencyAdminRule []interface{}
	for _, oldEmergencyAdminItem := range oldEmergencyAdmin {
		oldEmergencyAdminRule = append(oldEmergencyAdminRule, oldEmergencyAdminItem)
	}
	var newEmergencyAdminRule []interface{}
	for _, newEmergencyAdminItem := range newEmergencyAdmin {
		newEmergencyAdminRule = append(newEmergencyAdminRule, newEmergencyAdminItem)
	}

	logs, sub, err := _Events.contract.FilterLogs(opts, "EmergencyAdminSet", callerRule, oldEmergencyAdminRule, newEmergencyAdminRule)
	if err != nil {
		return nil, err
	}
	return &EventsEmergencyAdminSetIterator{contract: _Events.contract, event: "EmergencyAdminSet", logs: logs, sub: sub}, nil
}

// WatchEmergencyAdminSet is a free log subscription operation binding the contract event 0x676c0801b0f400762e958ee31cfbb10870e70786f6761f57c8647e766b0db3d9.
//
// Solidity: event EmergencyAdminSet(address indexed caller, address indexed oldEmergencyAdmin, address indexed newEmergencyAdmin, uint256 timestamp)
func (_Events *EventsFilterer) WatchEmergencyAdminSet(opts *bind.WatchOpts, sink chan<- *EventsEmergencyAdminSet, caller []common.Address, oldEmergencyAdmin []common.Address, newEmergencyAdmin []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var oldEmergencyAdminRule []interface{}
	for _, oldEmergencyAdminItem := range oldEmergencyAdmin {
		oldEmergencyAdminRule = append(oldEmergencyAdminRule, oldEmergencyAdminItem)
	}
	var newEmergencyAdminRule []interface{}
	for _, newEmergencyAdminItem := range newEmergencyAdmin {
		newEmergencyAdminRule = append(newEmergencyAdminRule, newEmergencyAdminItem)
	}

	logs, sub, err := _Events.contract.WatchLogs(opts, "EmergencyAdminSet", callerRule, oldEmergencyAdminRule, newEmergencyAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventsEmergencyAdminSet)
				if err := _Events.contract.UnpackLog(event, "EmergencyAdminSet", log); err != nil {
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

// ParseEmergencyAdminSet is a log parse operation binding the contract event 0x676c0801b0f400762e958ee31cfbb10870e70786f6761f57c8647e766b0db3d9.
//
// Solidity: event EmergencyAdminSet(address indexed caller, address indexed oldEmergencyAdmin, address indexed newEmergencyAdmin, uint256 timestamp)
func (_Events *EventsFilterer) ParseEmergencyAdminSet(log types.Log) (*EventsEmergencyAdminSet, error) {
	event := new(EventsEmergencyAdminSet)
	if err := _Events.contract.UnpackLog(event, "EmergencyAdminSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventsFeeModuleBaseConstructedIterator is returned from FilterFeeModuleBaseConstructed and is used to iterate over the raw logs and unpacked data for FeeModuleBaseConstructed events raised by the Events contract.
type EventsFeeModuleBaseConstructedIterator struct {
	Event *EventsFeeModuleBaseConstructed // Event containing the contract specifics and raw log

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
func (it *EventsFeeModuleBaseConstructedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventsFeeModuleBaseConstructed)
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
		it.Event = new(EventsFeeModuleBaseConstructed)
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
func (it *EventsFeeModuleBaseConstructedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventsFeeModuleBaseConstructedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventsFeeModuleBaseConstructed represents a FeeModuleBaseConstructed event raised by the Events contract.
type EventsFeeModuleBaseConstructed struct {
	ModuleGlobals common.Address
	Timestamp     *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFeeModuleBaseConstructed is a free log retrieval operation binding the contract event 0x4e84a529f4c627b5e787037d117873af1018768804cca3c7f0d47041fe2c89ed.
//
// Solidity: event FeeModuleBaseConstructed(address indexed moduleGlobals, uint256 timestamp)
func (_Events *EventsFilterer) FilterFeeModuleBaseConstructed(opts *bind.FilterOpts, moduleGlobals []common.Address) (*EventsFeeModuleBaseConstructedIterator, error) {

	var moduleGlobalsRule []interface{}
	for _, moduleGlobalsItem := range moduleGlobals {
		moduleGlobalsRule = append(moduleGlobalsRule, moduleGlobalsItem)
	}

	logs, sub, err := _Events.contract.FilterLogs(opts, "FeeModuleBaseConstructed", moduleGlobalsRule)
	if err != nil {
		return nil, err
	}
	return &EventsFeeModuleBaseConstructedIterator{contract: _Events.contract, event: "FeeModuleBaseConstructed", logs: logs, sub: sub}, nil
}

// WatchFeeModuleBaseConstructed is a free log subscription operation binding the contract event 0x4e84a529f4c627b5e787037d117873af1018768804cca3c7f0d47041fe2c89ed.
//
// Solidity: event FeeModuleBaseConstructed(address indexed moduleGlobals, uint256 timestamp)
func (_Events *EventsFilterer) WatchFeeModuleBaseConstructed(opts *bind.WatchOpts, sink chan<- *EventsFeeModuleBaseConstructed, moduleGlobals []common.Address) (event.Subscription, error) {

	var moduleGlobalsRule []interface{}
	for _, moduleGlobalsItem := range moduleGlobals {
		moduleGlobalsRule = append(moduleGlobalsRule, moduleGlobalsItem)
	}

	logs, sub, err := _Events.contract.WatchLogs(opts, "FeeModuleBaseConstructed", moduleGlobalsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventsFeeModuleBaseConstructed)
				if err := _Events.contract.UnpackLog(event, "FeeModuleBaseConstructed", log); err != nil {
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

// ParseFeeModuleBaseConstructed is a log parse operation binding the contract event 0x4e84a529f4c627b5e787037d117873af1018768804cca3c7f0d47041fe2c89ed.
//
// Solidity: event FeeModuleBaseConstructed(address indexed moduleGlobals, uint256 timestamp)
func (_Events *EventsFilterer) ParseFeeModuleBaseConstructed(log types.Log) (*EventsFeeModuleBaseConstructed, error) {
	event := new(EventsFeeModuleBaseConstructed)
	if err := _Events.contract.UnpackLog(event, "FeeModuleBaseConstructed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventsFollowModuleSetIterator is returned from FilterFollowModuleSet and is used to iterate over the raw logs and unpacked data for FollowModuleSet events raised by the Events contract.
type EventsFollowModuleSetIterator struct {
	Event *EventsFollowModuleSet // Event containing the contract specifics and raw log

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
func (it *EventsFollowModuleSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventsFollowModuleSet)
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
		it.Event = new(EventsFollowModuleSet)
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
func (it *EventsFollowModuleSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventsFollowModuleSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventsFollowModuleSet represents a FollowModuleSet event raised by the Events contract.
type EventsFollowModuleSet struct {
	ProfileId              *big.Int
	FollowModule           common.Address
	FollowModuleReturnData []byte
	Timestamp              *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterFollowModuleSet is a free log retrieval operation binding the contract event 0x92d95e400932d129885e627b38b169cbb28443ffaaa282d0fba0cf8797721359.
//
// Solidity: event FollowModuleSet(uint256 indexed profileId, address followModule, bytes followModuleReturnData, uint256 timestamp)
func (_Events *EventsFilterer) FilterFollowModuleSet(opts *bind.FilterOpts, profileId []*big.Int) (*EventsFollowModuleSetIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _Events.contract.FilterLogs(opts, "FollowModuleSet", profileIdRule)
	if err != nil {
		return nil, err
	}
	return &EventsFollowModuleSetIterator{contract: _Events.contract, event: "FollowModuleSet", logs: logs, sub: sub}, nil
}

// WatchFollowModuleSet is a free log subscription operation binding the contract event 0x92d95e400932d129885e627b38b169cbb28443ffaaa282d0fba0cf8797721359.
//
// Solidity: event FollowModuleSet(uint256 indexed profileId, address followModule, bytes followModuleReturnData, uint256 timestamp)
func (_Events *EventsFilterer) WatchFollowModuleSet(opts *bind.WatchOpts, sink chan<- *EventsFollowModuleSet, profileId []*big.Int) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _Events.contract.WatchLogs(opts, "FollowModuleSet", profileIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventsFollowModuleSet)
				if err := _Events.contract.UnpackLog(event, "FollowModuleSet", log); err != nil {
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

// ParseFollowModuleSet is a log parse operation binding the contract event 0x92d95e400932d129885e627b38b169cbb28443ffaaa282d0fba0cf8797721359.
//
// Solidity: event FollowModuleSet(uint256 indexed profileId, address followModule, bytes followModuleReturnData, uint256 timestamp)
func (_Events *EventsFilterer) ParseFollowModuleSet(log types.Log) (*EventsFollowModuleSet, error) {
	event := new(EventsFollowModuleSet)
	if err := _Events.contract.UnpackLog(event, "FollowModuleSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventsFollowModuleWhitelistedIterator is returned from FilterFollowModuleWhitelisted and is used to iterate over the raw logs and unpacked data for FollowModuleWhitelisted events raised by the Events contract.
type EventsFollowModuleWhitelistedIterator struct {
	Event *EventsFollowModuleWhitelisted // Event containing the contract specifics and raw log

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
func (it *EventsFollowModuleWhitelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventsFollowModuleWhitelisted)
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
		it.Event = new(EventsFollowModuleWhitelisted)
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
func (it *EventsFollowModuleWhitelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventsFollowModuleWhitelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventsFollowModuleWhitelisted represents a FollowModuleWhitelisted event raised by the Events contract.
type EventsFollowModuleWhitelisted struct {
	FollowModule common.Address
	Whitelisted  bool
	Timestamp    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFollowModuleWhitelisted is a free log retrieval operation binding the contract event 0x52c5b7889df9f12f84ec3da051e854e5876678370d8357959c23ef59dd6486df.
//
// Solidity: event FollowModuleWhitelisted(address indexed followModule, bool indexed whitelisted, uint256 timestamp)
func (_Events *EventsFilterer) FilterFollowModuleWhitelisted(opts *bind.FilterOpts, followModule []common.Address, whitelisted []bool) (*EventsFollowModuleWhitelistedIterator, error) {

	var followModuleRule []interface{}
	for _, followModuleItem := range followModule {
		followModuleRule = append(followModuleRule, followModuleItem)
	}
	var whitelistedRule []interface{}
	for _, whitelistedItem := range whitelisted {
		whitelistedRule = append(whitelistedRule, whitelistedItem)
	}

	logs, sub, err := _Events.contract.FilterLogs(opts, "FollowModuleWhitelisted", followModuleRule, whitelistedRule)
	if err != nil {
		return nil, err
	}
	return &EventsFollowModuleWhitelistedIterator{contract: _Events.contract, event: "FollowModuleWhitelisted", logs: logs, sub: sub}, nil
}

// WatchFollowModuleWhitelisted is a free log subscription operation binding the contract event 0x52c5b7889df9f12f84ec3da051e854e5876678370d8357959c23ef59dd6486df.
//
// Solidity: event FollowModuleWhitelisted(address indexed followModule, bool indexed whitelisted, uint256 timestamp)
func (_Events *EventsFilterer) WatchFollowModuleWhitelisted(opts *bind.WatchOpts, sink chan<- *EventsFollowModuleWhitelisted, followModule []common.Address, whitelisted []bool) (event.Subscription, error) {

	var followModuleRule []interface{}
	for _, followModuleItem := range followModule {
		followModuleRule = append(followModuleRule, followModuleItem)
	}
	var whitelistedRule []interface{}
	for _, whitelistedItem := range whitelisted {
		whitelistedRule = append(whitelistedRule, whitelistedItem)
	}

	logs, sub, err := _Events.contract.WatchLogs(opts, "FollowModuleWhitelisted", followModuleRule, whitelistedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventsFollowModuleWhitelisted)
				if err := _Events.contract.UnpackLog(event, "FollowModuleWhitelisted", log); err != nil {
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

// ParseFollowModuleWhitelisted is a log parse operation binding the contract event 0x52c5b7889df9f12f84ec3da051e854e5876678370d8357959c23ef59dd6486df.
//
// Solidity: event FollowModuleWhitelisted(address indexed followModule, bool indexed whitelisted, uint256 timestamp)
func (_Events *EventsFilterer) ParseFollowModuleWhitelisted(log types.Log) (*EventsFollowModuleWhitelisted, error) {
	event := new(EventsFollowModuleWhitelisted)
	if err := _Events.contract.UnpackLog(event, "FollowModuleWhitelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventsFollowNFTDelegatedPowerChangedIterator is returned from FilterFollowNFTDelegatedPowerChanged and is used to iterate over the raw logs and unpacked data for FollowNFTDelegatedPowerChanged events raised by the Events contract.
type EventsFollowNFTDelegatedPowerChangedIterator struct {
	Event *EventsFollowNFTDelegatedPowerChanged // Event containing the contract specifics and raw log

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
func (it *EventsFollowNFTDelegatedPowerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventsFollowNFTDelegatedPowerChanged)
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
		it.Event = new(EventsFollowNFTDelegatedPowerChanged)
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
func (it *EventsFollowNFTDelegatedPowerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventsFollowNFTDelegatedPowerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventsFollowNFTDelegatedPowerChanged represents a FollowNFTDelegatedPowerChanged event raised by the Events contract.
type EventsFollowNFTDelegatedPowerChanged struct {
	Delegate  common.Address
	NewPower  *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFollowNFTDelegatedPowerChanged is a free log retrieval operation binding the contract event 0xd9a6070174f4ccca76ed4896432e9a090b16e07e8fe27f275f50b33500b98e52.
//
// Solidity: event FollowNFTDelegatedPowerChanged(address indexed delegate, uint256 indexed newPower, uint256 timestamp)
func (_Events *EventsFilterer) FilterFollowNFTDelegatedPowerChanged(opts *bind.FilterOpts, delegate []common.Address, newPower []*big.Int) (*EventsFollowNFTDelegatedPowerChangedIterator, error) {

	var delegateRule []interface{}
	for _, delegateItem := range delegate {
		delegateRule = append(delegateRule, delegateItem)
	}
	var newPowerRule []interface{}
	for _, newPowerItem := range newPower {
		newPowerRule = append(newPowerRule, newPowerItem)
	}

	logs, sub, err := _Events.contract.FilterLogs(opts, "FollowNFTDelegatedPowerChanged", delegateRule, newPowerRule)
	if err != nil {
		return nil, err
	}
	return &EventsFollowNFTDelegatedPowerChangedIterator{contract: _Events.contract, event: "FollowNFTDelegatedPowerChanged", logs: logs, sub: sub}, nil
}

// WatchFollowNFTDelegatedPowerChanged is a free log subscription operation binding the contract event 0xd9a6070174f4ccca76ed4896432e9a090b16e07e8fe27f275f50b33500b98e52.
//
// Solidity: event FollowNFTDelegatedPowerChanged(address indexed delegate, uint256 indexed newPower, uint256 timestamp)
func (_Events *EventsFilterer) WatchFollowNFTDelegatedPowerChanged(opts *bind.WatchOpts, sink chan<- *EventsFollowNFTDelegatedPowerChanged, delegate []common.Address, newPower []*big.Int) (event.Subscription, error) {

	var delegateRule []interface{}
	for _, delegateItem := range delegate {
		delegateRule = append(delegateRule, delegateItem)
	}
	var newPowerRule []interface{}
	for _, newPowerItem := range newPower {
		newPowerRule = append(newPowerRule, newPowerItem)
	}

	logs, sub, err := _Events.contract.WatchLogs(opts, "FollowNFTDelegatedPowerChanged", delegateRule, newPowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventsFollowNFTDelegatedPowerChanged)
				if err := _Events.contract.UnpackLog(event, "FollowNFTDelegatedPowerChanged", log); err != nil {
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

// ParseFollowNFTDelegatedPowerChanged is a log parse operation binding the contract event 0xd9a6070174f4ccca76ed4896432e9a090b16e07e8fe27f275f50b33500b98e52.
//
// Solidity: event FollowNFTDelegatedPowerChanged(address indexed delegate, uint256 indexed newPower, uint256 timestamp)
func (_Events *EventsFilterer) ParseFollowNFTDelegatedPowerChanged(log types.Log) (*EventsFollowNFTDelegatedPowerChanged, error) {
	event := new(EventsFollowNFTDelegatedPowerChanged)
	if err := _Events.contract.UnpackLog(event, "FollowNFTDelegatedPowerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventsFollowNFTDeployedIterator is returned from FilterFollowNFTDeployed and is used to iterate over the raw logs and unpacked data for FollowNFTDeployed events raised by the Events contract.
type EventsFollowNFTDeployedIterator struct {
	Event *EventsFollowNFTDeployed // Event containing the contract specifics and raw log

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
func (it *EventsFollowNFTDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventsFollowNFTDeployed)
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
		it.Event = new(EventsFollowNFTDeployed)
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
func (it *EventsFollowNFTDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventsFollowNFTDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventsFollowNFTDeployed represents a FollowNFTDeployed event raised by the Events contract.
type EventsFollowNFTDeployed struct {
	ProfileId *big.Int
	FollowNFT common.Address
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFollowNFTDeployed is a free log retrieval operation binding the contract event 0x44403e38baed5e40df7f64ff8708b076c75a0dfda8380e75df5c36f11a476743.
//
// Solidity: event FollowNFTDeployed(uint256 indexed profileId, address indexed followNFT, uint256 timestamp)
func (_Events *EventsFilterer) FilterFollowNFTDeployed(opts *bind.FilterOpts, profileId []*big.Int, followNFT []common.Address) (*EventsFollowNFTDeployedIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var followNFTRule []interface{}
	for _, followNFTItem := range followNFT {
		followNFTRule = append(followNFTRule, followNFTItem)
	}

	logs, sub, err := _Events.contract.FilterLogs(opts, "FollowNFTDeployed", profileIdRule, followNFTRule)
	if err != nil {
		return nil, err
	}
	return &EventsFollowNFTDeployedIterator{contract: _Events.contract, event: "FollowNFTDeployed", logs: logs, sub: sub}, nil
}

// WatchFollowNFTDeployed is a free log subscription operation binding the contract event 0x44403e38baed5e40df7f64ff8708b076c75a0dfda8380e75df5c36f11a476743.
//
// Solidity: event FollowNFTDeployed(uint256 indexed profileId, address indexed followNFT, uint256 timestamp)
func (_Events *EventsFilterer) WatchFollowNFTDeployed(opts *bind.WatchOpts, sink chan<- *EventsFollowNFTDeployed, profileId []*big.Int, followNFT []common.Address) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var followNFTRule []interface{}
	for _, followNFTItem := range followNFT {
		followNFTRule = append(followNFTRule, followNFTItem)
	}

	logs, sub, err := _Events.contract.WatchLogs(opts, "FollowNFTDeployed", profileIdRule, followNFTRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventsFollowNFTDeployed)
				if err := _Events.contract.UnpackLog(event, "FollowNFTDeployed", log); err != nil {
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

// ParseFollowNFTDeployed is a log parse operation binding the contract event 0x44403e38baed5e40df7f64ff8708b076c75a0dfda8380e75df5c36f11a476743.
//
// Solidity: event FollowNFTDeployed(uint256 indexed profileId, address indexed followNFT, uint256 timestamp)
func (_Events *EventsFilterer) ParseFollowNFTDeployed(log types.Log) (*EventsFollowNFTDeployed, error) {
	event := new(EventsFollowNFTDeployed)
	if err := _Events.contract.UnpackLog(event, "FollowNFTDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventsFollowNFTInitializedIterator is returned from FilterFollowNFTInitialized and is used to iterate over the raw logs and unpacked data for FollowNFTInitialized events raised by the Events contract.
type EventsFollowNFTInitializedIterator struct {
	Event *EventsFollowNFTInitialized // Event containing the contract specifics and raw log

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
func (it *EventsFollowNFTInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventsFollowNFTInitialized)
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
		it.Event = new(EventsFollowNFTInitialized)
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
func (it *EventsFollowNFTInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventsFollowNFTInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventsFollowNFTInitialized represents a FollowNFTInitialized event raised by the Events contract.
type EventsFollowNFTInitialized struct {
	ProfileId *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFollowNFTInitialized is a free log retrieval operation binding the contract event 0xaec15127df11a6b562c87d31bcb8f4cd2f0cf57fb9b663d6334abf41fea94d95.
//
// Solidity: event FollowNFTInitialized(uint256 indexed profileId, uint256 timestamp)
func (_Events *EventsFilterer) FilterFollowNFTInitialized(opts *bind.FilterOpts, profileId []*big.Int) (*EventsFollowNFTInitializedIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _Events.contract.FilterLogs(opts, "FollowNFTInitialized", profileIdRule)
	if err != nil {
		return nil, err
	}
	return &EventsFollowNFTInitializedIterator{contract: _Events.contract, event: "FollowNFTInitialized", logs: logs, sub: sub}, nil
}

// WatchFollowNFTInitialized is a free log subscription operation binding the contract event 0xaec15127df11a6b562c87d31bcb8f4cd2f0cf57fb9b663d6334abf41fea94d95.
//
// Solidity: event FollowNFTInitialized(uint256 indexed profileId, uint256 timestamp)
func (_Events *EventsFilterer) WatchFollowNFTInitialized(opts *bind.WatchOpts, sink chan<- *EventsFollowNFTInitialized, profileId []*big.Int) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _Events.contract.WatchLogs(opts, "FollowNFTInitialized", profileIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventsFollowNFTInitialized)
				if err := _Events.contract.UnpackLog(event, "FollowNFTInitialized", log); err != nil {
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

// ParseFollowNFTInitialized is a log parse operation binding the contract event 0xaec15127df11a6b562c87d31bcb8f4cd2f0cf57fb9b663d6334abf41fea94d95.
//
// Solidity: event FollowNFTInitialized(uint256 indexed profileId, uint256 timestamp)
func (_Events *EventsFilterer) ParseFollowNFTInitialized(log types.Log) (*EventsFollowNFTInitialized, error) {
	event := new(EventsFollowNFTInitialized)
	if err := _Events.contract.UnpackLog(event, "FollowNFTInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventsFollowNFTTransferredIterator is returned from FilterFollowNFTTransferred and is used to iterate over the raw logs and unpacked data for FollowNFTTransferred events raised by the Events contract.
type EventsFollowNFTTransferredIterator struct {
	Event *EventsFollowNFTTransferred // Event containing the contract specifics and raw log

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
func (it *EventsFollowNFTTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventsFollowNFTTransferred)
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
		it.Event = new(EventsFollowNFTTransferred)
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
func (it *EventsFollowNFTTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventsFollowNFTTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventsFollowNFTTransferred represents a FollowNFTTransferred event raised by the Events contract.
type EventsFollowNFTTransferred struct {
	ProfileId   *big.Int
	FollowNFTId *big.Int
	From        common.Address
	To          common.Address
	Timestamp   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFollowNFTTransferred is a free log retrieval operation binding the contract event 0x4996ad2257e7db44908136c43128cc10ca988096f67dc6bb0bcee11d151368fb.
//
// Solidity: event FollowNFTTransferred(uint256 indexed profileId, uint256 indexed followNFTId, address from, address to, uint256 timestamp)
func (_Events *EventsFilterer) FilterFollowNFTTransferred(opts *bind.FilterOpts, profileId []*big.Int, followNFTId []*big.Int) (*EventsFollowNFTTransferredIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var followNFTIdRule []interface{}
	for _, followNFTIdItem := range followNFTId {
		followNFTIdRule = append(followNFTIdRule, followNFTIdItem)
	}

	logs, sub, err := _Events.contract.FilterLogs(opts, "FollowNFTTransferred", profileIdRule, followNFTIdRule)
	if err != nil {
		return nil, err
	}
	return &EventsFollowNFTTransferredIterator{contract: _Events.contract, event: "FollowNFTTransferred", logs: logs, sub: sub}, nil
}

// WatchFollowNFTTransferred is a free log subscription operation binding the contract event 0x4996ad2257e7db44908136c43128cc10ca988096f67dc6bb0bcee11d151368fb.
//
// Solidity: event FollowNFTTransferred(uint256 indexed profileId, uint256 indexed followNFTId, address from, address to, uint256 timestamp)
func (_Events *EventsFilterer) WatchFollowNFTTransferred(opts *bind.WatchOpts, sink chan<- *EventsFollowNFTTransferred, profileId []*big.Int, followNFTId []*big.Int) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var followNFTIdRule []interface{}
	for _, followNFTIdItem := range followNFTId {
		followNFTIdRule = append(followNFTIdRule, followNFTIdItem)
	}

	logs, sub, err := _Events.contract.WatchLogs(opts, "FollowNFTTransferred", profileIdRule, followNFTIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventsFollowNFTTransferred)
				if err := _Events.contract.UnpackLog(event, "FollowNFTTransferred", log); err != nil {
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

// ParseFollowNFTTransferred is a log parse operation binding the contract event 0x4996ad2257e7db44908136c43128cc10ca988096f67dc6bb0bcee11d151368fb.
//
// Solidity: event FollowNFTTransferred(uint256 indexed profileId, uint256 indexed followNFTId, address from, address to, uint256 timestamp)
func (_Events *EventsFilterer) ParseFollowNFTTransferred(log types.Log) (*EventsFollowNFTTransferred, error) {
	event := new(EventsFollowNFTTransferred)
	if err := _Events.contract.UnpackLog(event, "FollowNFTTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventsFollowNFTURISetIterator is returned from FilterFollowNFTURISet and is used to iterate over the raw logs and unpacked data for FollowNFTURISet events raised by the Events contract.
type EventsFollowNFTURISetIterator struct {
	Event *EventsFollowNFTURISet // Event containing the contract specifics and raw log

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
func (it *EventsFollowNFTURISetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventsFollowNFTURISet)
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
		it.Event = new(EventsFollowNFTURISet)
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
func (it *EventsFollowNFTURISetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventsFollowNFTURISetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventsFollowNFTURISet represents a FollowNFTURISet event raised by the Events contract.
type EventsFollowNFTURISet struct {
	ProfileId    *big.Int
	FollowNFTURI string
	Timestamp    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFollowNFTURISet is a free log retrieval operation binding the contract event 0xe82886e1af6fcab5caef13815b22f51384e970c367a785f265d13860a7d6966d.
//
// Solidity: event FollowNFTURISet(uint256 indexed profileId, string followNFTURI, uint256 timestamp)
func (_Events *EventsFilterer) FilterFollowNFTURISet(opts *bind.FilterOpts, profileId []*big.Int) (*EventsFollowNFTURISetIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _Events.contract.FilterLogs(opts, "FollowNFTURISet", profileIdRule)
	if err != nil {
		return nil, err
	}
	return &EventsFollowNFTURISetIterator{contract: _Events.contract, event: "FollowNFTURISet", logs: logs, sub: sub}, nil
}

// WatchFollowNFTURISet is a free log subscription operation binding the contract event 0xe82886e1af6fcab5caef13815b22f51384e970c367a785f265d13860a7d6966d.
//
// Solidity: event FollowNFTURISet(uint256 indexed profileId, string followNFTURI, uint256 timestamp)
func (_Events *EventsFilterer) WatchFollowNFTURISet(opts *bind.WatchOpts, sink chan<- *EventsFollowNFTURISet, profileId []*big.Int) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _Events.contract.WatchLogs(opts, "FollowNFTURISet", profileIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventsFollowNFTURISet)
				if err := _Events.contract.UnpackLog(event, "FollowNFTURISet", log); err != nil {
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

// ParseFollowNFTURISet is a log parse operation binding the contract event 0xe82886e1af6fcab5caef13815b22f51384e970c367a785f265d13860a7d6966d.
//
// Solidity: event FollowNFTURISet(uint256 indexed profileId, string followNFTURI, uint256 timestamp)
func (_Events *EventsFilterer) ParseFollowNFTURISet(log types.Log) (*EventsFollowNFTURISet, error) {
	event := new(EventsFollowNFTURISet)
	if err := _Events.contract.UnpackLog(event, "FollowNFTURISet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventsFollowedIterator is returned from FilterFollowed and is used to iterate over the raw logs and unpacked data for Followed events raised by the Events contract.
type EventsFollowedIterator struct {
	Event *EventsFollowed // Event containing the contract specifics and raw log

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
func (it *EventsFollowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventsFollowed)
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
		it.Event = new(EventsFollowed)
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
func (it *EventsFollowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventsFollowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventsFollowed represents a Followed event raised by the Events contract.
type EventsFollowed struct {
	Follower          common.Address
	ProfileIds        []*big.Int
	FollowModuleDatas [][]byte
	Timestamp         *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterFollowed is a free log retrieval operation binding the contract event 0x40487072dc56f384287d26fbe090f404143c2737d54632177451d1f74bd82c76.
//
// Solidity: event Followed(address indexed follower, uint256[] profileIds, bytes[] followModuleDatas, uint256 timestamp)
func (_Events *EventsFilterer) FilterFollowed(opts *bind.FilterOpts, follower []common.Address) (*EventsFollowedIterator, error) {

	var followerRule []interface{}
	for _, followerItem := range follower {
		followerRule = append(followerRule, followerItem)
	}

	logs, sub, err := _Events.contract.FilterLogs(opts, "Followed", followerRule)
	if err != nil {
		return nil, err
	}
	return &EventsFollowedIterator{contract: _Events.contract, event: "Followed", logs: logs, sub: sub}, nil
}

// WatchFollowed is a free log subscription operation binding the contract event 0x40487072dc56f384287d26fbe090f404143c2737d54632177451d1f74bd82c76.
//
// Solidity: event Followed(address indexed follower, uint256[] profileIds, bytes[] followModuleDatas, uint256 timestamp)
func (_Events *EventsFilterer) WatchFollowed(opts *bind.WatchOpts, sink chan<- *EventsFollowed, follower []common.Address) (event.Subscription, error) {

	var followerRule []interface{}
	for _, followerItem := range follower {
		followerRule = append(followerRule, followerItem)
	}

	logs, sub, err := _Events.contract.WatchLogs(opts, "Followed", followerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventsFollowed)
				if err := _Events.contract.UnpackLog(event, "Followed", log); err != nil {
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

// ParseFollowed is a log parse operation binding the contract event 0x40487072dc56f384287d26fbe090f404143c2737d54632177451d1f74bd82c76.
//
// Solidity: event Followed(address indexed follower, uint256[] profileIds, bytes[] followModuleDatas, uint256 timestamp)
func (_Events *EventsFilterer) ParseFollowed(log types.Log) (*EventsFollowed, error) {
	event := new(EventsFollowed)
	if err := _Events.contract.UnpackLog(event, "Followed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventsFollowsApprovedIterator is returned from FilterFollowsApproved and is used to iterate over the raw logs and unpacked data for FollowsApproved events raised by the Events contract.
type EventsFollowsApprovedIterator struct {
	Event *EventsFollowsApproved // Event containing the contract specifics and raw log

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
func (it *EventsFollowsApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventsFollowsApproved)
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
		it.Event = new(EventsFollowsApproved)
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
func (it *EventsFollowsApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventsFollowsApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventsFollowsApproved represents a FollowsApproved event raised by the Events contract.
type EventsFollowsApproved struct {
	Owner     common.Address
	ProfileId *big.Int
	Addresses []common.Address
	Approved  []bool
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFollowsApproved is a free log retrieval operation binding the contract event 0xc67fc3972da5d6434ab7b796ba133c240d40ee4e69129963c5aa0f2a6f7c3ad6.
//
// Solidity: event FollowsApproved(address indexed owner, uint256 indexed profileId, address[] addresses, bool[] approved, uint256 timestamp)
func (_Events *EventsFilterer) FilterFollowsApproved(opts *bind.FilterOpts, owner []common.Address, profileId []*big.Int) (*EventsFollowsApprovedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _Events.contract.FilterLogs(opts, "FollowsApproved", ownerRule, profileIdRule)
	if err != nil {
		return nil, err
	}
	return &EventsFollowsApprovedIterator{contract: _Events.contract, event: "FollowsApproved", logs: logs, sub: sub}, nil
}

// WatchFollowsApproved is a free log subscription operation binding the contract event 0xc67fc3972da5d6434ab7b796ba133c240d40ee4e69129963c5aa0f2a6f7c3ad6.
//
// Solidity: event FollowsApproved(address indexed owner, uint256 indexed profileId, address[] addresses, bool[] approved, uint256 timestamp)
func (_Events *EventsFilterer) WatchFollowsApproved(opts *bind.WatchOpts, sink chan<- *EventsFollowsApproved, owner []common.Address, profileId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _Events.contract.WatchLogs(opts, "FollowsApproved", ownerRule, profileIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventsFollowsApproved)
				if err := _Events.contract.UnpackLog(event, "FollowsApproved", log); err != nil {
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

// ParseFollowsApproved is a log parse operation binding the contract event 0xc67fc3972da5d6434ab7b796ba133c240d40ee4e69129963c5aa0f2a6f7c3ad6.
//
// Solidity: event FollowsApproved(address indexed owner, uint256 indexed profileId, address[] addresses, bool[] approved, uint256 timestamp)
func (_Events *EventsFilterer) ParseFollowsApproved(log types.Log) (*EventsFollowsApproved, error) {
	event := new(EventsFollowsApproved)
	if err := _Events.contract.UnpackLog(event, "FollowsApproved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventsFollowsToggledIterator is returned from FilterFollowsToggled and is used to iterate over the raw logs and unpacked data for FollowsToggled events raised by the Events contract.
type EventsFollowsToggledIterator struct {
	Event *EventsFollowsToggled // Event containing the contract specifics and raw log

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
func (it *EventsFollowsToggledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventsFollowsToggled)
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
		it.Event = new(EventsFollowsToggled)
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
func (it *EventsFollowsToggledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventsFollowsToggledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventsFollowsToggled represents a FollowsToggled event raised by the Events contract.
type EventsFollowsToggled struct {
	Owner      common.Address
	ProfileIds []*big.Int
	Enabled    []bool
	Timestamp  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFollowsToggled is a free log retrieval operation binding the contract event 0x5538c80c8d3bee397d87a7d153f7f085bb12adf2fe25a026c7cc4e83d8c5f1d7.
//
// Solidity: event FollowsToggled(address indexed owner, uint256[] profileIds, bool[] enabled, uint256 timestamp)
func (_Events *EventsFilterer) FilterFollowsToggled(opts *bind.FilterOpts, owner []common.Address) (*EventsFollowsToggledIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Events.contract.FilterLogs(opts, "FollowsToggled", ownerRule)
	if err != nil {
		return nil, err
	}
	return &EventsFollowsToggledIterator{contract: _Events.contract, event: "FollowsToggled", logs: logs, sub: sub}, nil
}

// WatchFollowsToggled is a free log subscription operation binding the contract event 0x5538c80c8d3bee397d87a7d153f7f085bb12adf2fe25a026c7cc4e83d8c5f1d7.
//
// Solidity: event FollowsToggled(address indexed owner, uint256[] profileIds, bool[] enabled, uint256 timestamp)
func (_Events *EventsFilterer) WatchFollowsToggled(opts *bind.WatchOpts, sink chan<- *EventsFollowsToggled, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Events.contract.WatchLogs(opts, "FollowsToggled", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventsFollowsToggled)
				if err := _Events.contract.UnpackLog(event, "FollowsToggled", log); err != nil {
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

// ParseFollowsToggled is a log parse operation binding the contract event 0x5538c80c8d3bee397d87a7d153f7f085bb12adf2fe25a026c7cc4e83d8c5f1d7.
//
// Solidity: event FollowsToggled(address indexed owner, uint256[] profileIds, bool[] enabled, uint256 timestamp)
func (_Events *EventsFilterer) ParseFollowsToggled(log types.Log) (*EventsFollowsToggled, error) {
	event := new(EventsFollowsToggled)
	if err := _Events.contract.UnpackLog(event, "FollowsToggled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventsGovernanceSetIterator is returned from FilterGovernanceSet and is used to iterate over the raw logs and unpacked data for GovernanceSet events raised by the Events contract.
type EventsGovernanceSetIterator struct {
	Event *EventsGovernanceSet // Event containing the contract specifics and raw log

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
func (it *EventsGovernanceSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventsGovernanceSet)
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
		it.Event = new(EventsGovernanceSet)
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
func (it *EventsGovernanceSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventsGovernanceSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventsGovernanceSet represents a GovernanceSet event raised by the Events contract.
type EventsGovernanceSet struct {
	Caller         common.Address
	PrevGovernance common.Address
	NewGovernance  common.Address
	Timestamp      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterGovernanceSet is a free log retrieval operation binding the contract event 0xe552a55455b740845a5c07ed445d1724142fc997b389835495a29b30cddc1ccd.
//
// Solidity: event GovernanceSet(address indexed caller, address indexed prevGovernance, address indexed newGovernance, uint256 timestamp)
func (_Events *EventsFilterer) FilterGovernanceSet(opts *bind.FilterOpts, caller []common.Address, prevGovernance []common.Address, newGovernance []common.Address) (*EventsGovernanceSetIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var prevGovernanceRule []interface{}
	for _, prevGovernanceItem := range prevGovernance {
		prevGovernanceRule = append(prevGovernanceRule, prevGovernanceItem)
	}
	var newGovernanceRule []interface{}
	for _, newGovernanceItem := range newGovernance {
		newGovernanceRule = append(newGovernanceRule, newGovernanceItem)
	}

	logs, sub, err := _Events.contract.FilterLogs(opts, "GovernanceSet", callerRule, prevGovernanceRule, newGovernanceRule)
	if err != nil {
		return nil, err
	}
	return &EventsGovernanceSetIterator{contract: _Events.contract, event: "GovernanceSet", logs: logs, sub: sub}, nil
}

// WatchGovernanceSet is a free log subscription operation binding the contract event 0xe552a55455b740845a5c07ed445d1724142fc997b389835495a29b30cddc1ccd.
//
// Solidity: event GovernanceSet(address indexed caller, address indexed prevGovernance, address indexed newGovernance, uint256 timestamp)
func (_Events *EventsFilterer) WatchGovernanceSet(opts *bind.WatchOpts, sink chan<- *EventsGovernanceSet, caller []common.Address, prevGovernance []common.Address, newGovernance []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var prevGovernanceRule []interface{}
	for _, prevGovernanceItem := range prevGovernance {
		prevGovernanceRule = append(prevGovernanceRule, prevGovernanceItem)
	}
	var newGovernanceRule []interface{}
	for _, newGovernanceItem := range newGovernance {
		newGovernanceRule = append(newGovernanceRule, newGovernanceItem)
	}

	logs, sub, err := _Events.contract.WatchLogs(opts, "GovernanceSet", callerRule, prevGovernanceRule, newGovernanceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventsGovernanceSet)
				if err := _Events.contract.UnpackLog(event, "GovernanceSet", log); err != nil {
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

// ParseGovernanceSet is a log parse operation binding the contract event 0xe552a55455b740845a5c07ed445d1724142fc997b389835495a29b30cddc1ccd.
//
// Solidity: event GovernanceSet(address indexed caller, address indexed prevGovernance, address indexed newGovernance, uint256 timestamp)
func (_Events *EventsFilterer) ParseGovernanceSet(log types.Log) (*EventsGovernanceSet, error) {
	event := new(EventsGovernanceSet)
	if err := _Events.contract.UnpackLog(event, "GovernanceSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventsMirrorCreatedIterator is returned from FilterMirrorCreated and is used to iterate over the raw logs and unpacked data for MirrorCreated events raised by the Events contract.
type EventsMirrorCreatedIterator struct {
	Event *EventsMirrorCreated // Event containing the contract specifics and raw log

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
func (it *EventsMirrorCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventsMirrorCreated)
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
		it.Event = new(EventsMirrorCreated)
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
func (it *EventsMirrorCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventsMirrorCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventsMirrorCreated represents a MirrorCreated event raised by the Events contract.
type EventsMirrorCreated struct {
	ProfileId                 *big.Int
	PubId                     *big.Int
	ProfileIdPointed          *big.Int
	PubIdPointed              *big.Int
	ReferenceModuleData       []byte
	ReferenceModule           common.Address
	ReferenceModuleReturnData []byte
	Timestamp                 *big.Int
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterMirrorCreated is a free log retrieval operation binding the contract event 0x9ea5dedb85bd9da4e264ee5a39b7ba0982e5d4d035d55edfa98a36b00e770b5a.
//
// Solidity: event MirrorCreated(uint256 indexed profileId, uint256 indexed pubId, uint256 profileIdPointed, uint256 pubIdPointed, bytes referenceModuleData, address referenceModule, bytes referenceModuleReturnData, uint256 timestamp)
func (_Events *EventsFilterer) FilterMirrorCreated(opts *bind.FilterOpts, profileId []*big.Int, pubId []*big.Int) (*EventsMirrorCreatedIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var pubIdRule []interface{}
	for _, pubIdItem := range pubId {
		pubIdRule = append(pubIdRule, pubIdItem)
	}

	logs, sub, err := _Events.contract.FilterLogs(opts, "MirrorCreated", profileIdRule, pubIdRule)
	if err != nil {
		return nil, err
	}
	return &EventsMirrorCreatedIterator{contract: _Events.contract, event: "MirrorCreated", logs: logs, sub: sub}, nil
}

// WatchMirrorCreated is a free log subscription operation binding the contract event 0x9ea5dedb85bd9da4e264ee5a39b7ba0982e5d4d035d55edfa98a36b00e770b5a.
//
// Solidity: event MirrorCreated(uint256 indexed profileId, uint256 indexed pubId, uint256 profileIdPointed, uint256 pubIdPointed, bytes referenceModuleData, address referenceModule, bytes referenceModuleReturnData, uint256 timestamp)
func (_Events *EventsFilterer) WatchMirrorCreated(opts *bind.WatchOpts, sink chan<- *EventsMirrorCreated, profileId []*big.Int, pubId []*big.Int) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var pubIdRule []interface{}
	for _, pubIdItem := range pubId {
		pubIdRule = append(pubIdRule, pubIdItem)
	}

	logs, sub, err := _Events.contract.WatchLogs(opts, "MirrorCreated", profileIdRule, pubIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventsMirrorCreated)
				if err := _Events.contract.UnpackLog(event, "MirrorCreated", log); err != nil {
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

// ParseMirrorCreated is a log parse operation binding the contract event 0x9ea5dedb85bd9da4e264ee5a39b7ba0982e5d4d035d55edfa98a36b00e770b5a.
//
// Solidity: event MirrorCreated(uint256 indexed profileId, uint256 indexed pubId, uint256 profileIdPointed, uint256 pubIdPointed, bytes referenceModuleData, address referenceModule, bytes referenceModuleReturnData, uint256 timestamp)
func (_Events *EventsFilterer) ParseMirrorCreated(log types.Log) (*EventsMirrorCreated, error) {
	event := new(EventsMirrorCreated)
	if err := _Events.contract.UnpackLog(event, "MirrorCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventsModuleBaseConstructedIterator is returned from FilterModuleBaseConstructed and is used to iterate over the raw logs and unpacked data for ModuleBaseConstructed events raised by the Events contract.
type EventsModuleBaseConstructedIterator struct {
	Event *EventsModuleBaseConstructed // Event containing the contract specifics and raw log

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
func (it *EventsModuleBaseConstructedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventsModuleBaseConstructed)
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
		it.Event = new(EventsModuleBaseConstructed)
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
func (it *EventsModuleBaseConstructedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventsModuleBaseConstructedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventsModuleBaseConstructed represents a ModuleBaseConstructed event raised by the Events contract.
type EventsModuleBaseConstructed struct {
	Hub       common.Address
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterModuleBaseConstructed is a free log retrieval operation binding the contract event 0xf1a1fa6b64aa95186f5a1285e76198d0da80d9c5a88062641d447f1d7c54e56c.
//
// Solidity: event ModuleBaseConstructed(address indexed hub, uint256 timestamp)
func (_Events *EventsFilterer) FilterModuleBaseConstructed(opts *bind.FilterOpts, hub []common.Address) (*EventsModuleBaseConstructedIterator, error) {

	var hubRule []interface{}
	for _, hubItem := range hub {
		hubRule = append(hubRule, hubItem)
	}

	logs, sub, err := _Events.contract.FilterLogs(opts, "ModuleBaseConstructed", hubRule)
	if err != nil {
		return nil, err
	}
	return &EventsModuleBaseConstructedIterator{contract: _Events.contract, event: "ModuleBaseConstructed", logs: logs, sub: sub}, nil
}

// WatchModuleBaseConstructed is a free log subscription operation binding the contract event 0xf1a1fa6b64aa95186f5a1285e76198d0da80d9c5a88062641d447f1d7c54e56c.
//
// Solidity: event ModuleBaseConstructed(address indexed hub, uint256 timestamp)
func (_Events *EventsFilterer) WatchModuleBaseConstructed(opts *bind.WatchOpts, sink chan<- *EventsModuleBaseConstructed, hub []common.Address) (event.Subscription, error) {

	var hubRule []interface{}
	for _, hubItem := range hub {
		hubRule = append(hubRule, hubItem)
	}

	logs, sub, err := _Events.contract.WatchLogs(opts, "ModuleBaseConstructed", hubRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventsModuleBaseConstructed)
				if err := _Events.contract.UnpackLog(event, "ModuleBaseConstructed", log); err != nil {
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

// ParseModuleBaseConstructed is a log parse operation binding the contract event 0xf1a1fa6b64aa95186f5a1285e76198d0da80d9c5a88062641d447f1d7c54e56c.
//
// Solidity: event ModuleBaseConstructed(address indexed hub, uint256 timestamp)
func (_Events *EventsFilterer) ParseModuleBaseConstructed(log types.Log) (*EventsModuleBaseConstructed, error) {
	event := new(EventsModuleBaseConstructed)
	if err := _Events.contract.UnpackLog(event, "ModuleBaseConstructed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventsModuleGlobalsCurrencyWhitelistedIterator is returned from FilterModuleGlobalsCurrencyWhitelisted and is used to iterate over the raw logs and unpacked data for ModuleGlobalsCurrencyWhitelisted events raised by the Events contract.
type EventsModuleGlobalsCurrencyWhitelistedIterator struct {
	Event *EventsModuleGlobalsCurrencyWhitelisted // Event containing the contract specifics and raw log

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
func (it *EventsModuleGlobalsCurrencyWhitelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventsModuleGlobalsCurrencyWhitelisted)
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
		it.Event = new(EventsModuleGlobalsCurrencyWhitelisted)
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
func (it *EventsModuleGlobalsCurrencyWhitelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventsModuleGlobalsCurrencyWhitelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventsModuleGlobalsCurrencyWhitelisted represents a ModuleGlobalsCurrencyWhitelisted event raised by the Events contract.
type EventsModuleGlobalsCurrencyWhitelisted struct {
	Currency        common.Address
	PrevWhitelisted bool
	Whitelisted     bool
	Timestamp       *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterModuleGlobalsCurrencyWhitelisted is a free log retrieval operation binding the contract event 0x79c3cefc851fd6040f06af202c542818d9fb39bcddcb7a7e3f563b15300a2743.
//
// Solidity: event ModuleGlobalsCurrencyWhitelisted(address indexed currency, bool indexed prevWhitelisted, bool indexed whitelisted, uint256 timestamp)
func (_Events *EventsFilterer) FilterModuleGlobalsCurrencyWhitelisted(opts *bind.FilterOpts, currency []common.Address, prevWhitelisted []bool, whitelisted []bool) (*EventsModuleGlobalsCurrencyWhitelistedIterator, error) {

	var currencyRule []interface{}
	for _, currencyItem := range currency {
		currencyRule = append(currencyRule, currencyItem)
	}
	var prevWhitelistedRule []interface{}
	for _, prevWhitelistedItem := range prevWhitelisted {
		prevWhitelistedRule = append(prevWhitelistedRule, prevWhitelistedItem)
	}
	var whitelistedRule []interface{}
	for _, whitelistedItem := range whitelisted {
		whitelistedRule = append(whitelistedRule, whitelistedItem)
	}

	logs, sub, err := _Events.contract.FilterLogs(opts, "ModuleGlobalsCurrencyWhitelisted", currencyRule, prevWhitelistedRule, whitelistedRule)
	if err != nil {
		return nil, err
	}
	return &EventsModuleGlobalsCurrencyWhitelistedIterator{contract: _Events.contract, event: "ModuleGlobalsCurrencyWhitelisted", logs: logs, sub: sub}, nil
}

// WatchModuleGlobalsCurrencyWhitelisted is a free log subscription operation binding the contract event 0x79c3cefc851fd6040f06af202c542818d9fb39bcddcb7a7e3f563b15300a2743.
//
// Solidity: event ModuleGlobalsCurrencyWhitelisted(address indexed currency, bool indexed prevWhitelisted, bool indexed whitelisted, uint256 timestamp)
func (_Events *EventsFilterer) WatchModuleGlobalsCurrencyWhitelisted(opts *bind.WatchOpts, sink chan<- *EventsModuleGlobalsCurrencyWhitelisted, currency []common.Address, prevWhitelisted []bool, whitelisted []bool) (event.Subscription, error) {

	var currencyRule []interface{}
	for _, currencyItem := range currency {
		currencyRule = append(currencyRule, currencyItem)
	}
	var prevWhitelistedRule []interface{}
	for _, prevWhitelistedItem := range prevWhitelisted {
		prevWhitelistedRule = append(prevWhitelistedRule, prevWhitelistedItem)
	}
	var whitelistedRule []interface{}
	for _, whitelistedItem := range whitelisted {
		whitelistedRule = append(whitelistedRule, whitelistedItem)
	}

	logs, sub, err := _Events.contract.WatchLogs(opts, "ModuleGlobalsCurrencyWhitelisted", currencyRule, prevWhitelistedRule, whitelistedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventsModuleGlobalsCurrencyWhitelisted)
				if err := _Events.contract.UnpackLog(event, "ModuleGlobalsCurrencyWhitelisted", log); err != nil {
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

// ParseModuleGlobalsCurrencyWhitelisted is a log parse operation binding the contract event 0x79c3cefc851fd6040f06af202c542818d9fb39bcddcb7a7e3f563b15300a2743.
//
// Solidity: event ModuleGlobalsCurrencyWhitelisted(address indexed currency, bool indexed prevWhitelisted, bool indexed whitelisted, uint256 timestamp)
func (_Events *EventsFilterer) ParseModuleGlobalsCurrencyWhitelisted(log types.Log) (*EventsModuleGlobalsCurrencyWhitelisted, error) {
	event := new(EventsModuleGlobalsCurrencyWhitelisted)
	if err := _Events.contract.UnpackLog(event, "ModuleGlobalsCurrencyWhitelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventsModuleGlobalsGovernanceSetIterator is returned from FilterModuleGlobalsGovernanceSet and is used to iterate over the raw logs and unpacked data for ModuleGlobalsGovernanceSet events raised by the Events contract.
type EventsModuleGlobalsGovernanceSetIterator struct {
	Event *EventsModuleGlobalsGovernanceSet // Event containing the contract specifics and raw log

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
func (it *EventsModuleGlobalsGovernanceSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventsModuleGlobalsGovernanceSet)
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
		it.Event = new(EventsModuleGlobalsGovernanceSet)
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
func (it *EventsModuleGlobalsGovernanceSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventsModuleGlobalsGovernanceSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventsModuleGlobalsGovernanceSet represents a ModuleGlobalsGovernanceSet event raised by the Events contract.
type EventsModuleGlobalsGovernanceSet struct {
	PrevGovernance common.Address
	NewGovernance  common.Address
	Timestamp      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterModuleGlobalsGovernanceSet is a free log retrieval operation binding the contract event 0xbf538a2c0db3d440906b8179dd0394f68a65b0b1481da70ffee24e19dccee84c.
//
// Solidity: event ModuleGlobalsGovernanceSet(address indexed prevGovernance, address indexed newGovernance, uint256 timestamp)
func (_Events *EventsFilterer) FilterModuleGlobalsGovernanceSet(opts *bind.FilterOpts, prevGovernance []common.Address, newGovernance []common.Address) (*EventsModuleGlobalsGovernanceSetIterator, error) {

	var prevGovernanceRule []interface{}
	for _, prevGovernanceItem := range prevGovernance {
		prevGovernanceRule = append(prevGovernanceRule, prevGovernanceItem)
	}
	var newGovernanceRule []interface{}
	for _, newGovernanceItem := range newGovernance {
		newGovernanceRule = append(newGovernanceRule, newGovernanceItem)
	}

	logs, sub, err := _Events.contract.FilterLogs(opts, "ModuleGlobalsGovernanceSet", prevGovernanceRule, newGovernanceRule)
	if err != nil {
		return nil, err
	}
	return &EventsModuleGlobalsGovernanceSetIterator{contract: _Events.contract, event: "ModuleGlobalsGovernanceSet", logs: logs, sub: sub}, nil
}

// WatchModuleGlobalsGovernanceSet is a free log subscription operation binding the contract event 0xbf538a2c0db3d440906b8179dd0394f68a65b0b1481da70ffee24e19dccee84c.
//
// Solidity: event ModuleGlobalsGovernanceSet(address indexed prevGovernance, address indexed newGovernance, uint256 timestamp)
func (_Events *EventsFilterer) WatchModuleGlobalsGovernanceSet(opts *bind.WatchOpts, sink chan<- *EventsModuleGlobalsGovernanceSet, prevGovernance []common.Address, newGovernance []common.Address) (event.Subscription, error) {

	var prevGovernanceRule []interface{}
	for _, prevGovernanceItem := range prevGovernance {
		prevGovernanceRule = append(prevGovernanceRule, prevGovernanceItem)
	}
	var newGovernanceRule []interface{}
	for _, newGovernanceItem := range newGovernance {
		newGovernanceRule = append(newGovernanceRule, newGovernanceItem)
	}

	logs, sub, err := _Events.contract.WatchLogs(opts, "ModuleGlobalsGovernanceSet", prevGovernanceRule, newGovernanceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventsModuleGlobalsGovernanceSet)
				if err := _Events.contract.UnpackLog(event, "ModuleGlobalsGovernanceSet", log); err != nil {
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

// ParseModuleGlobalsGovernanceSet is a log parse operation binding the contract event 0xbf538a2c0db3d440906b8179dd0394f68a65b0b1481da70ffee24e19dccee84c.
//
// Solidity: event ModuleGlobalsGovernanceSet(address indexed prevGovernance, address indexed newGovernance, uint256 timestamp)
func (_Events *EventsFilterer) ParseModuleGlobalsGovernanceSet(log types.Log) (*EventsModuleGlobalsGovernanceSet, error) {
	event := new(EventsModuleGlobalsGovernanceSet)
	if err := _Events.contract.UnpackLog(event, "ModuleGlobalsGovernanceSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventsModuleGlobalsTreasuryFeeSetIterator is returned from FilterModuleGlobalsTreasuryFeeSet and is used to iterate over the raw logs and unpacked data for ModuleGlobalsTreasuryFeeSet events raised by the Events contract.
type EventsModuleGlobalsTreasuryFeeSetIterator struct {
	Event *EventsModuleGlobalsTreasuryFeeSet // Event containing the contract specifics and raw log

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
func (it *EventsModuleGlobalsTreasuryFeeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventsModuleGlobalsTreasuryFeeSet)
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
		it.Event = new(EventsModuleGlobalsTreasuryFeeSet)
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
func (it *EventsModuleGlobalsTreasuryFeeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventsModuleGlobalsTreasuryFeeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventsModuleGlobalsTreasuryFeeSet represents a ModuleGlobalsTreasuryFeeSet event raised by the Events contract.
type EventsModuleGlobalsTreasuryFeeSet struct {
	PrevTreasuryFee uint16
	NewTreasuryFee  uint16
	Timestamp       *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterModuleGlobalsTreasuryFeeSet is a free log retrieval operation binding the contract event 0xec936862e6bb897cd711a5f31825057583128c2a482f0a4c9a4e6c3fd7c023f4.
//
// Solidity: event ModuleGlobalsTreasuryFeeSet(uint16 indexed prevTreasuryFee, uint16 indexed newTreasuryFee, uint256 timestamp)
func (_Events *EventsFilterer) FilterModuleGlobalsTreasuryFeeSet(opts *bind.FilterOpts, prevTreasuryFee []uint16, newTreasuryFee []uint16) (*EventsModuleGlobalsTreasuryFeeSetIterator, error) {

	var prevTreasuryFeeRule []interface{}
	for _, prevTreasuryFeeItem := range prevTreasuryFee {
		prevTreasuryFeeRule = append(prevTreasuryFeeRule, prevTreasuryFeeItem)
	}
	var newTreasuryFeeRule []interface{}
	for _, newTreasuryFeeItem := range newTreasuryFee {
		newTreasuryFeeRule = append(newTreasuryFeeRule, newTreasuryFeeItem)
	}

	logs, sub, err := _Events.contract.FilterLogs(opts, "ModuleGlobalsTreasuryFeeSet", prevTreasuryFeeRule, newTreasuryFeeRule)
	if err != nil {
		return nil, err
	}
	return &EventsModuleGlobalsTreasuryFeeSetIterator{contract: _Events.contract, event: "ModuleGlobalsTreasuryFeeSet", logs: logs, sub: sub}, nil
}

// WatchModuleGlobalsTreasuryFeeSet is a free log subscription operation binding the contract event 0xec936862e6bb897cd711a5f31825057583128c2a482f0a4c9a4e6c3fd7c023f4.
//
// Solidity: event ModuleGlobalsTreasuryFeeSet(uint16 indexed prevTreasuryFee, uint16 indexed newTreasuryFee, uint256 timestamp)
func (_Events *EventsFilterer) WatchModuleGlobalsTreasuryFeeSet(opts *bind.WatchOpts, sink chan<- *EventsModuleGlobalsTreasuryFeeSet, prevTreasuryFee []uint16, newTreasuryFee []uint16) (event.Subscription, error) {

	var prevTreasuryFeeRule []interface{}
	for _, prevTreasuryFeeItem := range prevTreasuryFee {
		prevTreasuryFeeRule = append(prevTreasuryFeeRule, prevTreasuryFeeItem)
	}
	var newTreasuryFeeRule []interface{}
	for _, newTreasuryFeeItem := range newTreasuryFee {
		newTreasuryFeeRule = append(newTreasuryFeeRule, newTreasuryFeeItem)
	}

	logs, sub, err := _Events.contract.WatchLogs(opts, "ModuleGlobalsTreasuryFeeSet", prevTreasuryFeeRule, newTreasuryFeeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventsModuleGlobalsTreasuryFeeSet)
				if err := _Events.contract.UnpackLog(event, "ModuleGlobalsTreasuryFeeSet", log); err != nil {
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

// ParseModuleGlobalsTreasuryFeeSet is a log parse operation binding the contract event 0xec936862e6bb897cd711a5f31825057583128c2a482f0a4c9a4e6c3fd7c023f4.
//
// Solidity: event ModuleGlobalsTreasuryFeeSet(uint16 indexed prevTreasuryFee, uint16 indexed newTreasuryFee, uint256 timestamp)
func (_Events *EventsFilterer) ParseModuleGlobalsTreasuryFeeSet(log types.Log) (*EventsModuleGlobalsTreasuryFeeSet, error) {
	event := new(EventsModuleGlobalsTreasuryFeeSet)
	if err := _Events.contract.UnpackLog(event, "ModuleGlobalsTreasuryFeeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventsModuleGlobalsTreasurySetIterator is returned from FilterModuleGlobalsTreasurySet and is used to iterate over the raw logs and unpacked data for ModuleGlobalsTreasurySet events raised by the Events contract.
type EventsModuleGlobalsTreasurySetIterator struct {
	Event *EventsModuleGlobalsTreasurySet // Event containing the contract specifics and raw log

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
func (it *EventsModuleGlobalsTreasurySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventsModuleGlobalsTreasurySet)
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
		it.Event = new(EventsModuleGlobalsTreasurySet)
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
func (it *EventsModuleGlobalsTreasurySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventsModuleGlobalsTreasurySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventsModuleGlobalsTreasurySet represents a ModuleGlobalsTreasurySet event raised by the Events contract.
type EventsModuleGlobalsTreasurySet struct {
	PrevTreasury common.Address
	NewTreasury  common.Address
	Timestamp    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterModuleGlobalsTreasurySet is a free log retrieval operation binding the contract event 0x3dfc53d6b49bfbc932b215ba515f0d0ab0e17aac17726fba48075f0c16c7ffe3.
//
// Solidity: event ModuleGlobalsTreasurySet(address indexed prevTreasury, address indexed newTreasury, uint256 timestamp)
func (_Events *EventsFilterer) FilterModuleGlobalsTreasurySet(opts *bind.FilterOpts, prevTreasury []common.Address, newTreasury []common.Address) (*EventsModuleGlobalsTreasurySetIterator, error) {

	var prevTreasuryRule []interface{}
	for _, prevTreasuryItem := range prevTreasury {
		prevTreasuryRule = append(prevTreasuryRule, prevTreasuryItem)
	}
	var newTreasuryRule []interface{}
	for _, newTreasuryItem := range newTreasury {
		newTreasuryRule = append(newTreasuryRule, newTreasuryItem)
	}

	logs, sub, err := _Events.contract.FilterLogs(opts, "ModuleGlobalsTreasurySet", prevTreasuryRule, newTreasuryRule)
	if err != nil {
		return nil, err
	}
	return &EventsModuleGlobalsTreasurySetIterator{contract: _Events.contract, event: "ModuleGlobalsTreasurySet", logs: logs, sub: sub}, nil
}

// WatchModuleGlobalsTreasurySet is a free log subscription operation binding the contract event 0x3dfc53d6b49bfbc932b215ba515f0d0ab0e17aac17726fba48075f0c16c7ffe3.
//
// Solidity: event ModuleGlobalsTreasurySet(address indexed prevTreasury, address indexed newTreasury, uint256 timestamp)
func (_Events *EventsFilterer) WatchModuleGlobalsTreasurySet(opts *bind.WatchOpts, sink chan<- *EventsModuleGlobalsTreasurySet, prevTreasury []common.Address, newTreasury []common.Address) (event.Subscription, error) {

	var prevTreasuryRule []interface{}
	for _, prevTreasuryItem := range prevTreasury {
		prevTreasuryRule = append(prevTreasuryRule, prevTreasuryItem)
	}
	var newTreasuryRule []interface{}
	for _, newTreasuryItem := range newTreasury {
		newTreasuryRule = append(newTreasuryRule, newTreasuryItem)
	}

	logs, sub, err := _Events.contract.WatchLogs(opts, "ModuleGlobalsTreasurySet", prevTreasuryRule, newTreasuryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventsModuleGlobalsTreasurySet)
				if err := _Events.contract.UnpackLog(event, "ModuleGlobalsTreasurySet", log); err != nil {
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

// ParseModuleGlobalsTreasurySet is a log parse operation binding the contract event 0x3dfc53d6b49bfbc932b215ba515f0d0ab0e17aac17726fba48075f0c16c7ffe3.
//
// Solidity: event ModuleGlobalsTreasurySet(address indexed prevTreasury, address indexed newTreasury, uint256 timestamp)
func (_Events *EventsFilterer) ParseModuleGlobalsTreasurySet(log types.Log) (*EventsModuleGlobalsTreasurySet, error) {
	event := new(EventsModuleGlobalsTreasurySet)
	if err := _Events.contract.UnpackLog(event, "ModuleGlobalsTreasurySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventsPostCreatedIterator is returned from FilterPostCreated and is used to iterate over the raw logs and unpacked data for PostCreated events raised by the Events contract.
type EventsPostCreatedIterator struct {
	Event *EventsPostCreated // Event containing the contract specifics and raw log

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
func (it *EventsPostCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventsPostCreated)
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
		it.Event = new(EventsPostCreated)
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
func (it *EventsPostCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventsPostCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventsPostCreated represents a PostCreated event raised by the Events contract.
type EventsPostCreated struct {
	ProfileId                 *big.Int
	PubId                     *big.Int
	ContentURI                string
	CollectModule             common.Address
	CollectModuleReturnData   []byte
	ReferenceModule           common.Address
	ReferenceModuleReturnData []byte
	Timestamp                 *big.Int
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterPostCreated is a free log retrieval operation binding the contract event 0xc672c38b4d26c3c978228e99164105280410b144af24dd3ed8e4f9d211d96a50.
//
// Solidity: event PostCreated(uint256 indexed profileId, uint256 indexed pubId, string contentURI, address collectModule, bytes collectModuleReturnData, address referenceModule, bytes referenceModuleReturnData, uint256 timestamp)
func (_Events *EventsFilterer) FilterPostCreated(opts *bind.FilterOpts, profileId []*big.Int, pubId []*big.Int) (*EventsPostCreatedIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var pubIdRule []interface{}
	for _, pubIdItem := range pubId {
		pubIdRule = append(pubIdRule, pubIdItem)
	}

	logs, sub, err := _Events.contract.FilterLogs(opts, "PostCreated", profileIdRule, pubIdRule)
	if err != nil {
		return nil, err
	}
	return &EventsPostCreatedIterator{contract: _Events.contract, event: "PostCreated", logs: logs, sub: sub}, nil
}

// WatchPostCreated is a free log subscription operation binding the contract event 0xc672c38b4d26c3c978228e99164105280410b144af24dd3ed8e4f9d211d96a50.
//
// Solidity: event PostCreated(uint256 indexed profileId, uint256 indexed pubId, string contentURI, address collectModule, bytes collectModuleReturnData, address referenceModule, bytes referenceModuleReturnData, uint256 timestamp)
func (_Events *EventsFilterer) WatchPostCreated(opts *bind.WatchOpts, sink chan<- *EventsPostCreated, profileId []*big.Int, pubId []*big.Int) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var pubIdRule []interface{}
	for _, pubIdItem := range pubId {
		pubIdRule = append(pubIdRule, pubIdItem)
	}

	logs, sub, err := _Events.contract.WatchLogs(opts, "PostCreated", profileIdRule, pubIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventsPostCreated)
				if err := _Events.contract.UnpackLog(event, "PostCreated", log); err != nil {
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

// ParsePostCreated is a log parse operation binding the contract event 0xc672c38b4d26c3c978228e99164105280410b144af24dd3ed8e4f9d211d96a50.
//
// Solidity: event PostCreated(uint256 indexed profileId, uint256 indexed pubId, string contentURI, address collectModule, bytes collectModuleReturnData, address referenceModule, bytes referenceModuleReturnData, uint256 timestamp)
func (_Events *EventsFilterer) ParsePostCreated(log types.Log) (*EventsPostCreated, error) {
	event := new(EventsPostCreated)
	if err := _Events.contract.UnpackLog(event, "PostCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventsProfileCreatedIterator is returned from FilterProfileCreated and is used to iterate over the raw logs and unpacked data for ProfileCreated events raised by the Events contract.
type EventsProfileCreatedIterator struct {
	Event *EventsProfileCreated // Event containing the contract specifics and raw log

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
func (it *EventsProfileCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventsProfileCreated)
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
		it.Event = new(EventsProfileCreated)
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
func (it *EventsProfileCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventsProfileCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventsProfileCreated represents a ProfileCreated event raised by the Events contract.
type EventsProfileCreated struct {
	ProfileId              *big.Int
	Creator                common.Address
	To                     common.Address
	Handle                 string
	ImageURI               string
	FollowModule           common.Address
	FollowModuleReturnData []byte
	FollowNFTURI           string
	Timestamp              *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterProfileCreated is a free log retrieval operation binding the contract event 0x4e14f57cff7910416f2ef43cf05019b5a97a313de71fec9344be11b9b88fed12.
//
// Solidity: event ProfileCreated(uint256 indexed profileId, address indexed creator, address indexed to, string handle, string imageURI, address followModule, bytes followModuleReturnData, string followNFTURI, uint256 timestamp)
func (_Events *EventsFilterer) FilterProfileCreated(opts *bind.FilterOpts, profileId []*big.Int, creator []common.Address, to []common.Address) (*EventsProfileCreatedIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Events.contract.FilterLogs(opts, "ProfileCreated", profileIdRule, creatorRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EventsProfileCreatedIterator{contract: _Events.contract, event: "ProfileCreated", logs: logs, sub: sub}, nil
}

// WatchProfileCreated is a free log subscription operation binding the contract event 0x4e14f57cff7910416f2ef43cf05019b5a97a313de71fec9344be11b9b88fed12.
//
// Solidity: event ProfileCreated(uint256 indexed profileId, address indexed creator, address indexed to, string handle, string imageURI, address followModule, bytes followModuleReturnData, string followNFTURI, uint256 timestamp)
func (_Events *EventsFilterer) WatchProfileCreated(opts *bind.WatchOpts, sink chan<- *EventsProfileCreated, profileId []*big.Int, creator []common.Address, to []common.Address) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Events.contract.WatchLogs(opts, "ProfileCreated", profileIdRule, creatorRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventsProfileCreated)
				if err := _Events.contract.UnpackLog(event, "ProfileCreated", log); err != nil {
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

// ParseProfileCreated is a log parse operation binding the contract event 0x4e14f57cff7910416f2ef43cf05019b5a97a313de71fec9344be11b9b88fed12.
//
// Solidity: event ProfileCreated(uint256 indexed profileId, address indexed creator, address indexed to, string handle, string imageURI, address followModule, bytes followModuleReturnData, string followNFTURI, uint256 timestamp)
func (_Events *EventsFilterer) ParseProfileCreated(log types.Log) (*EventsProfileCreated, error) {
	event := new(EventsProfileCreated)
	if err := _Events.contract.UnpackLog(event, "ProfileCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventsProfileCreatorWhitelistedIterator is returned from FilterProfileCreatorWhitelisted and is used to iterate over the raw logs and unpacked data for ProfileCreatorWhitelisted events raised by the Events contract.
type EventsProfileCreatorWhitelistedIterator struct {
	Event *EventsProfileCreatorWhitelisted // Event containing the contract specifics and raw log

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
func (it *EventsProfileCreatorWhitelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventsProfileCreatorWhitelisted)
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
		it.Event = new(EventsProfileCreatorWhitelisted)
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
func (it *EventsProfileCreatorWhitelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventsProfileCreatorWhitelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventsProfileCreatorWhitelisted represents a ProfileCreatorWhitelisted event raised by the Events contract.
type EventsProfileCreatorWhitelisted struct {
	ProfileCreator common.Address
	Whitelisted    bool
	Timestamp      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterProfileCreatorWhitelisted is a free log retrieval operation binding the contract event 0x8f617843889b94892bd44852d36ca6a7f49ecf4350a01e7b68e22d80f4ed95bc.
//
// Solidity: event ProfileCreatorWhitelisted(address indexed profileCreator, bool indexed whitelisted, uint256 timestamp)
func (_Events *EventsFilterer) FilterProfileCreatorWhitelisted(opts *bind.FilterOpts, profileCreator []common.Address, whitelisted []bool) (*EventsProfileCreatorWhitelistedIterator, error) {

	var profileCreatorRule []interface{}
	for _, profileCreatorItem := range profileCreator {
		profileCreatorRule = append(profileCreatorRule, profileCreatorItem)
	}
	var whitelistedRule []interface{}
	for _, whitelistedItem := range whitelisted {
		whitelistedRule = append(whitelistedRule, whitelistedItem)
	}

	logs, sub, err := _Events.contract.FilterLogs(opts, "ProfileCreatorWhitelisted", profileCreatorRule, whitelistedRule)
	if err != nil {
		return nil, err
	}
	return &EventsProfileCreatorWhitelistedIterator{contract: _Events.contract, event: "ProfileCreatorWhitelisted", logs: logs, sub: sub}, nil
}

// WatchProfileCreatorWhitelisted is a free log subscription operation binding the contract event 0x8f617843889b94892bd44852d36ca6a7f49ecf4350a01e7b68e22d80f4ed95bc.
//
// Solidity: event ProfileCreatorWhitelisted(address indexed profileCreator, bool indexed whitelisted, uint256 timestamp)
func (_Events *EventsFilterer) WatchProfileCreatorWhitelisted(opts *bind.WatchOpts, sink chan<- *EventsProfileCreatorWhitelisted, profileCreator []common.Address, whitelisted []bool) (event.Subscription, error) {

	var profileCreatorRule []interface{}
	for _, profileCreatorItem := range profileCreator {
		profileCreatorRule = append(profileCreatorRule, profileCreatorItem)
	}
	var whitelistedRule []interface{}
	for _, whitelistedItem := range whitelisted {
		whitelistedRule = append(whitelistedRule, whitelistedItem)
	}

	logs, sub, err := _Events.contract.WatchLogs(opts, "ProfileCreatorWhitelisted", profileCreatorRule, whitelistedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventsProfileCreatorWhitelisted)
				if err := _Events.contract.UnpackLog(event, "ProfileCreatorWhitelisted", log); err != nil {
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

// ParseProfileCreatorWhitelisted is a log parse operation binding the contract event 0x8f617843889b94892bd44852d36ca6a7f49ecf4350a01e7b68e22d80f4ed95bc.
//
// Solidity: event ProfileCreatorWhitelisted(address indexed profileCreator, bool indexed whitelisted, uint256 timestamp)
func (_Events *EventsFilterer) ParseProfileCreatorWhitelisted(log types.Log) (*EventsProfileCreatorWhitelisted, error) {
	event := new(EventsProfileCreatorWhitelisted)
	if err := _Events.contract.UnpackLog(event, "ProfileCreatorWhitelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventsProfileImageURISetIterator is returned from FilterProfileImageURISet and is used to iterate over the raw logs and unpacked data for ProfileImageURISet events raised by the Events contract.
type EventsProfileImageURISetIterator struct {
	Event *EventsProfileImageURISet // Event containing the contract specifics and raw log

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
func (it *EventsProfileImageURISetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventsProfileImageURISet)
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
		it.Event = new(EventsProfileImageURISet)
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
func (it *EventsProfileImageURISetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventsProfileImageURISetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventsProfileImageURISet represents a ProfileImageURISet event raised by the Events contract.
type EventsProfileImageURISet struct {
	ProfileId *big.Int
	ImageURI  string
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProfileImageURISet is a free log retrieval operation binding the contract event 0xd5a5879cad33c830cc1432c1850107029a09c80c60e9bce1ecd08d24880bd46c.
//
// Solidity: event ProfileImageURISet(uint256 indexed profileId, string imageURI, uint256 timestamp)
func (_Events *EventsFilterer) FilterProfileImageURISet(opts *bind.FilterOpts, profileId []*big.Int) (*EventsProfileImageURISetIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _Events.contract.FilterLogs(opts, "ProfileImageURISet", profileIdRule)
	if err != nil {
		return nil, err
	}
	return &EventsProfileImageURISetIterator{contract: _Events.contract, event: "ProfileImageURISet", logs: logs, sub: sub}, nil
}

// WatchProfileImageURISet is a free log subscription operation binding the contract event 0xd5a5879cad33c830cc1432c1850107029a09c80c60e9bce1ecd08d24880bd46c.
//
// Solidity: event ProfileImageURISet(uint256 indexed profileId, string imageURI, uint256 timestamp)
func (_Events *EventsFilterer) WatchProfileImageURISet(opts *bind.WatchOpts, sink chan<- *EventsProfileImageURISet, profileId []*big.Int) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _Events.contract.WatchLogs(opts, "ProfileImageURISet", profileIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventsProfileImageURISet)
				if err := _Events.contract.UnpackLog(event, "ProfileImageURISet", log); err != nil {
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

// ParseProfileImageURISet is a log parse operation binding the contract event 0xd5a5879cad33c830cc1432c1850107029a09c80c60e9bce1ecd08d24880bd46c.
//
// Solidity: event ProfileImageURISet(uint256 indexed profileId, string imageURI, uint256 timestamp)
func (_Events *EventsFilterer) ParseProfileImageURISet(log types.Log) (*EventsProfileImageURISet, error) {
	event := new(EventsProfileImageURISet)
	if err := _Events.contract.UnpackLog(event, "ProfileImageURISet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventsProfileMetadataSetIterator is returned from FilterProfileMetadataSet and is used to iterate over the raw logs and unpacked data for ProfileMetadataSet events raised by the Events contract.
type EventsProfileMetadataSetIterator struct {
	Event *EventsProfileMetadataSet // Event containing the contract specifics and raw log

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
func (it *EventsProfileMetadataSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventsProfileMetadataSet)
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
		it.Event = new(EventsProfileMetadataSet)
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
func (it *EventsProfileMetadataSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventsProfileMetadataSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventsProfileMetadataSet represents a ProfileMetadataSet event raised by the Events contract.
type EventsProfileMetadataSet struct {
	ProfileId *big.Int
	Metadata  string
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProfileMetadataSet is a free log retrieval operation binding the contract event 0xf901a8b3832914a45999dd4c425fbe42eb4182724d394100401e633d9f0b286a.
//
// Solidity: event ProfileMetadataSet(uint256 indexed profileId, string metadata, uint256 timestamp)
func (_Events *EventsFilterer) FilterProfileMetadataSet(opts *bind.FilterOpts, profileId []*big.Int) (*EventsProfileMetadataSetIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _Events.contract.FilterLogs(opts, "ProfileMetadataSet", profileIdRule)
	if err != nil {
		return nil, err
	}
	return &EventsProfileMetadataSetIterator{contract: _Events.contract, event: "ProfileMetadataSet", logs: logs, sub: sub}, nil
}

// WatchProfileMetadataSet is a free log subscription operation binding the contract event 0xf901a8b3832914a45999dd4c425fbe42eb4182724d394100401e633d9f0b286a.
//
// Solidity: event ProfileMetadataSet(uint256 indexed profileId, string metadata, uint256 timestamp)
func (_Events *EventsFilterer) WatchProfileMetadataSet(opts *bind.WatchOpts, sink chan<- *EventsProfileMetadataSet, profileId []*big.Int) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _Events.contract.WatchLogs(opts, "ProfileMetadataSet", profileIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventsProfileMetadataSet)
				if err := _Events.contract.UnpackLog(event, "ProfileMetadataSet", log); err != nil {
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

// ParseProfileMetadataSet is a log parse operation binding the contract event 0xf901a8b3832914a45999dd4c425fbe42eb4182724d394100401e633d9f0b286a.
//
// Solidity: event ProfileMetadataSet(uint256 indexed profileId, string metadata, uint256 timestamp)
func (_Events *EventsFilterer) ParseProfileMetadataSet(log types.Log) (*EventsProfileMetadataSet, error) {
	event := new(EventsProfileMetadataSet)
	if err := _Events.contract.UnpackLog(event, "ProfileMetadataSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventsReferenceModuleWhitelistedIterator is returned from FilterReferenceModuleWhitelisted and is used to iterate over the raw logs and unpacked data for ReferenceModuleWhitelisted events raised by the Events contract.
type EventsReferenceModuleWhitelistedIterator struct {
	Event *EventsReferenceModuleWhitelisted // Event containing the contract specifics and raw log

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
func (it *EventsReferenceModuleWhitelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventsReferenceModuleWhitelisted)
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
		it.Event = new(EventsReferenceModuleWhitelisted)
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
func (it *EventsReferenceModuleWhitelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventsReferenceModuleWhitelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventsReferenceModuleWhitelisted represents a ReferenceModuleWhitelisted event raised by the Events contract.
type EventsReferenceModuleWhitelisted struct {
	ReferenceModule common.Address
	Whitelisted     bool
	Timestamp       *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterReferenceModuleWhitelisted is a free log retrieval operation binding the contract event 0x37872a053ef20cb52defb7c9ec20e1a87cb8dd5098ac9e76a144be263dfef572.
//
// Solidity: event ReferenceModuleWhitelisted(address indexed referenceModule, bool indexed whitelisted, uint256 timestamp)
func (_Events *EventsFilterer) FilterReferenceModuleWhitelisted(opts *bind.FilterOpts, referenceModule []common.Address, whitelisted []bool) (*EventsReferenceModuleWhitelistedIterator, error) {

	var referenceModuleRule []interface{}
	for _, referenceModuleItem := range referenceModule {
		referenceModuleRule = append(referenceModuleRule, referenceModuleItem)
	}
	var whitelistedRule []interface{}
	for _, whitelistedItem := range whitelisted {
		whitelistedRule = append(whitelistedRule, whitelistedItem)
	}

	logs, sub, err := _Events.contract.FilterLogs(opts, "ReferenceModuleWhitelisted", referenceModuleRule, whitelistedRule)
	if err != nil {
		return nil, err
	}
	return &EventsReferenceModuleWhitelistedIterator{contract: _Events.contract, event: "ReferenceModuleWhitelisted", logs: logs, sub: sub}, nil
}

// WatchReferenceModuleWhitelisted is a free log subscription operation binding the contract event 0x37872a053ef20cb52defb7c9ec20e1a87cb8dd5098ac9e76a144be263dfef572.
//
// Solidity: event ReferenceModuleWhitelisted(address indexed referenceModule, bool indexed whitelisted, uint256 timestamp)
func (_Events *EventsFilterer) WatchReferenceModuleWhitelisted(opts *bind.WatchOpts, sink chan<- *EventsReferenceModuleWhitelisted, referenceModule []common.Address, whitelisted []bool) (event.Subscription, error) {

	var referenceModuleRule []interface{}
	for _, referenceModuleItem := range referenceModule {
		referenceModuleRule = append(referenceModuleRule, referenceModuleItem)
	}
	var whitelistedRule []interface{}
	for _, whitelistedItem := range whitelisted {
		whitelistedRule = append(whitelistedRule, whitelistedItem)
	}

	logs, sub, err := _Events.contract.WatchLogs(opts, "ReferenceModuleWhitelisted", referenceModuleRule, whitelistedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventsReferenceModuleWhitelisted)
				if err := _Events.contract.UnpackLog(event, "ReferenceModuleWhitelisted", log); err != nil {
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

// ParseReferenceModuleWhitelisted is a log parse operation binding the contract event 0x37872a053ef20cb52defb7c9ec20e1a87cb8dd5098ac9e76a144be263dfef572.
//
// Solidity: event ReferenceModuleWhitelisted(address indexed referenceModule, bool indexed whitelisted, uint256 timestamp)
func (_Events *EventsFilterer) ParseReferenceModuleWhitelisted(log types.Log) (*EventsReferenceModuleWhitelisted, error) {
	event := new(EventsReferenceModuleWhitelisted)
	if err := _Events.contract.UnpackLog(event, "ReferenceModuleWhitelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventsStateSetIterator is returned from FilterStateSet and is used to iterate over the raw logs and unpacked data for StateSet events raised by the Events contract.
type EventsStateSetIterator struct {
	Event *EventsStateSet // Event containing the contract specifics and raw log

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
func (it *EventsStateSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventsStateSet)
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
		it.Event = new(EventsStateSet)
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
func (it *EventsStateSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventsStateSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventsStateSet represents a StateSet event raised by the Events contract.
type EventsStateSet struct {
	Caller    common.Address
	PrevState uint8
	NewState  uint8
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStateSet is a free log retrieval operation binding the contract event 0xa2f9a1499fc1f9b7796d21fe5761290ccb7e0ef6ccf35fa58b668f304a62a1ca.
//
// Solidity: event StateSet(address indexed caller, uint8 indexed prevState, uint8 indexed newState, uint256 timestamp)
func (_Events *EventsFilterer) FilterStateSet(opts *bind.FilterOpts, caller []common.Address, prevState []uint8, newState []uint8) (*EventsStateSetIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var prevStateRule []interface{}
	for _, prevStateItem := range prevState {
		prevStateRule = append(prevStateRule, prevStateItem)
	}
	var newStateRule []interface{}
	for _, newStateItem := range newState {
		newStateRule = append(newStateRule, newStateItem)
	}

	logs, sub, err := _Events.contract.FilterLogs(opts, "StateSet", callerRule, prevStateRule, newStateRule)
	if err != nil {
		return nil, err
	}
	return &EventsStateSetIterator{contract: _Events.contract, event: "StateSet", logs: logs, sub: sub}, nil
}

// WatchStateSet is a free log subscription operation binding the contract event 0xa2f9a1499fc1f9b7796d21fe5761290ccb7e0ef6ccf35fa58b668f304a62a1ca.
//
// Solidity: event StateSet(address indexed caller, uint8 indexed prevState, uint8 indexed newState, uint256 timestamp)
func (_Events *EventsFilterer) WatchStateSet(opts *bind.WatchOpts, sink chan<- *EventsStateSet, caller []common.Address, prevState []uint8, newState []uint8) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var prevStateRule []interface{}
	for _, prevStateItem := range prevState {
		prevStateRule = append(prevStateRule, prevStateItem)
	}
	var newStateRule []interface{}
	for _, newStateItem := range newState {
		newStateRule = append(newStateRule, newStateItem)
	}

	logs, sub, err := _Events.contract.WatchLogs(opts, "StateSet", callerRule, prevStateRule, newStateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventsStateSet)
				if err := _Events.contract.UnpackLog(event, "StateSet", log); err != nil {
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

// ParseStateSet is a log parse operation binding the contract event 0xa2f9a1499fc1f9b7796d21fe5761290ccb7e0ef6ccf35fa58b668f304a62a1ca.
//
// Solidity: event StateSet(address indexed caller, uint8 indexed prevState, uint8 indexed newState, uint256 timestamp)
func (_Events *EventsFilterer) ParseStateSet(log types.Log) (*EventsStateSet, error) {
	event := new(EventsStateSet)
	if err := _Events.contract.UnpackLog(event, "StateSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
