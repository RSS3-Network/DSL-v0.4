// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package socket

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

// RegistryBridgeRequest is an auto generated low-level Go binding around an user-defined struct.
type RegistryBridgeRequest struct {
	Id                   *big.Int
	OptionalNativeAmount *big.Int
	InputToken           common.Address
	Data                 []byte
}

// RegistryMiddlewareRequest is an auto generated low-level Go binding around an user-defined struct.
type RegistryMiddlewareRequest struct {
	Id                   *big.Int
	OptionalNativeAmount *big.Int
	InputToken           common.Address
	Data                 []byte
}

// RegistryRouteData is an auto generated low-level Go binding around an user-defined struct.
type RegistryRouteData struct {
	Route        common.Address
	IsEnabled    bool
	IsMiddleware bool
}

// RegistryUserRequest is an auto generated low-level Go binding around an user-defined struct.
type RegistryUserRequest struct {
	ReceiverAddress   common.Address
	ToChainId         *big.Int
	Amount            *big.Int
	MiddlewareRequest RegistryMiddlewareRequest
	BridgeRequest     RegistryBridgeRequest
}

// RegistryMetaData contains all meta data concerning the Registry contract.
var RegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"middlewareID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bridgeID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"inputAmount\",\"type\":\"uint256\"}],\"name\":\"ExecutionCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"routeID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"route\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isMiddleware\",\"type\":\"bool\"}],\"name\":\"NewRouteAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"routeID\",\"type\":\"uint256\"}],\"name\":\"RouteDisabled\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"route\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isMiddleware\",\"type\":\"bool\"}],\"internalType\":\"structRegistry.RouteData[]\",\"name\":\"_routes\",\"type\":\"tuple[]\"}],\"name\":\"addRoutes\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_routeId\",\"type\":\"uint256\"}],\"name\":\"disableRoute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"receiverAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"optionalNativeAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"inputToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structRegistry.MiddlewareRequest\",\"name\":\"middlewareRequest\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"optionalNativeAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"inputToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structRegistry.BridgeRequest\",\"name\":\"bridgeRequest\",\"type\":\"tuple\"}],\"internalType\":\"structRegistry.UserRequest\",\"name\":\"_userRequest\",\"type\":\"tuple\"}],\"name\":\"outboundTransferTo\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiverAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"rescueFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"routes\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"route\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isMiddleware\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// RegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use RegistryMetaData.ABI instead.
var RegistryABI = RegistryMetaData.ABI

// Registry is an auto generated Go binding around an Ethereum contract.
type Registry struct {
	RegistryCaller     // Read-only binding to the contract
	RegistryTransactor // Write-only binding to the contract
	RegistryFilterer   // Log filterer for contract events
}

// RegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type RegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RegistrySession struct {
	Contract     *Registry         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RegistryCallerSession struct {
	Contract *RegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// RegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RegistryTransactorSession struct {
	Contract     *RegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type RegistryRaw struct {
	Contract *Registry // Generic contract binding to access the raw methods on
}

// RegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RegistryCallerRaw struct {
	Contract *RegistryCaller // Generic read-only contract binding to access the raw methods on
}

// RegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RegistryTransactorRaw struct {
	Contract *RegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRegistry creates a new instance of Registry, bound to a specific deployed contract.
func NewRegistry(address common.Address, backend bind.ContractBackend) (*Registry, error) {
	contract, err := bindRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Registry{RegistryCaller: RegistryCaller{contract: contract}, RegistryTransactor: RegistryTransactor{contract: contract}, RegistryFilterer: RegistryFilterer{contract: contract}}, nil
}

// NewRegistryCaller creates a new read-only instance of Registry, bound to a specific deployed contract.
func NewRegistryCaller(address common.Address, caller bind.ContractCaller) (*RegistryCaller, error) {
	contract, err := bindRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryCaller{contract: contract}, nil
}

// NewRegistryTransactor creates a new write-only instance of Registry, bound to a specific deployed contract.
func NewRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*RegistryTransactor, error) {
	contract, err := bindRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryTransactor{contract: contract}, nil
}

// NewRegistryFilterer creates a new log filterer instance of Registry, bound to a specific deployed contract.
func NewRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*RegistryFilterer, error) {
	contract, err := bindRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RegistryFilterer{contract: contract}, nil
}

// bindRegistry binds a generic wrapper to an already deployed contract.
func bindRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.RegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Registry *RegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Registry *RegistrySession) Owner() (common.Address, error) {
	return _Registry.Contract.Owner(&_Registry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Registry *RegistryCallerSession) Owner() (common.Address, error) {
	return _Registry.Contract.Owner(&_Registry.CallOpts)
}

// Routes is a free data retrieval call binding the contract method 0x726f16d8.
//
// Solidity: function routes(uint256 ) view returns(address route, bool isEnabled, bool isMiddleware)
func (_Registry *RegistryCaller) Routes(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Route        common.Address
	IsEnabled    bool
	IsMiddleware bool
}, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "routes", arg0)

	outstruct := new(struct {
		Route        common.Address
		IsEnabled    bool
		IsMiddleware bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Route = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.IsEnabled = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.IsMiddleware = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// Routes is a free data retrieval call binding the contract method 0x726f16d8.
//
// Solidity: function routes(uint256 ) view returns(address route, bool isEnabled, bool isMiddleware)
func (_Registry *RegistrySession) Routes(arg0 *big.Int) (struct {
	Route        common.Address
	IsEnabled    bool
	IsMiddleware bool
}, error) {
	return _Registry.Contract.Routes(&_Registry.CallOpts, arg0)
}

// Routes is a free data retrieval call binding the contract method 0x726f16d8.
//
// Solidity: function routes(uint256 ) view returns(address route, bool isEnabled, bool isMiddleware)
func (_Registry *RegistryCallerSession) Routes(arg0 *big.Int) (struct {
	Route        common.Address
	IsEnabled    bool
	IsMiddleware bool
}, error) {
	return _Registry.Contract.Routes(&_Registry.CallOpts, arg0)
}

// AddRoutes is a paid mutator transaction binding the contract method 0x02a9c051.
//
// Solidity: function addRoutes((address,bool,bool)[] _routes) returns(uint256[])
func (_Registry *RegistryTransactor) AddRoutes(opts *bind.TransactOpts, _routes []RegistryRouteData) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "addRoutes", _routes)
}

// AddRoutes is a paid mutator transaction binding the contract method 0x02a9c051.
//
// Solidity: function addRoutes((address,bool,bool)[] _routes) returns(uint256[])
func (_Registry *RegistrySession) AddRoutes(_routes []RegistryRouteData) (*types.Transaction, error) {
	return _Registry.Contract.AddRoutes(&_Registry.TransactOpts, _routes)
}

// AddRoutes is a paid mutator transaction binding the contract method 0x02a9c051.
//
// Solidity: function addRoutes((address,bool,bool)[] _routes) returns(uint256[])
func (_Registry *RegistryTransactorSession) AddRoutes(_routes []RegistryRouteData) (*types.Transaction, error) {
	return _Registry.Contract.AddRoutes(&_Registry.TransactOpts, _routes)
}

// DisableRoute is a paid mutator transaction binding the contract method 0xffcdf4ed.
//
// Solidity: function disableRoute(uint256 _routeId) returns()
func (_Registry *RegistryTransactor) DisableRoute(opts *bind.TransactOpts, _routeId *big.Int) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "disableRoute", _routeId)
}

// DisableRoute is a paid mutator transaction binding the contract method 0xffcdf4ed.
//
// Solidity: function disableRoute(uint256 _routeId) returns()
func (_Registry *RegistrySession) DisableRoute(_routeId *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.DisableRoute(&_Registry.TransactOpts, _routeId)
}

// DisableRoute is a paid mutator transaction binding the contract method 0xffcdf4ed.
//
// Solidity: function disableRoute(uint256 _routeId) returns()
func (_Registry *RegistryTransactorSession) DisableRoute(_routeId *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.DisableRoute(&_Registry.TransactOpts, _routeId)
}

// OutboundTransferTo is a paid mutator transaction binding the contract method 0xa44bbb15.
//
// Solidity: function outboundTransferTo((address,uint256,uint256,(uint256,uint256,address,bytes),(uint256,uint256,address,bytes)) _userRequest) payable returns()
func (_Registry *RegistryTransactor) OutboundTransferTo(opts *bind.TransactOpts, _userRequest RegistryUserRequest) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "outboundTransferTo", _userRequest)
}

// OutboundTransferTo is a paid mutator transaction binding the contract method 0xa44bbb15.
//
// Solidity: function outboundTransferTo((address,uint256,uint256,(uint256,uint256,address,bytes),(uint256,uint256,address,bytes)) _userRequest) payable returns()
func (_Registry *RegistrySession) OutboundTransferTo(_userRequest RegistryUserRequest) (*types.Transaction, error) {
	return _Registry.Contract.OutboundTransferTo(&_Registry.TransactOpts, _userRequest)
}

// OutboundTransferTo is a paid mutator transaction binding the contract method 0xa44bbb15.
//
// Solidity: function outboundTransferTo((address,uint256,uint256,(uint256,uint256,address,bytes),(uint256,uint256,address,bytes)) _userRequest) payable returns()
func (_Registry *RegistryTransactorSession) OutboundTransferTo(_userRequest RegistryUserRequest) (*types.Transaction, error) {
	return _Registry.Contract.OutboundTransferTo(&_Registry.TransactOpts, _userRequest)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Registry *RegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Registry *RegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _Registry.Contract.RenounceOwnership(&_Registry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Registry *RegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Registry.Contract.RenounceOwnership(&_Registry.TransactOpts)
}

// RescueFunds is a paid mutator transaction binding the contract method 0x6ccae054.
//
// Solidity: function rescueFunds(address _token, address _receiverAddress, uint256 _amount) returns()
func (_Registry *RegistryTransactor) RescueFunds(opts *bind.TransactOpts, _token common.Address, _receiverAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "rescueFunds", _token, _receiverAddress, _amount)
}

// RescueFunds is a paid mutator transaction binding the contract method 0x6ccae054.
//
// Solidity: function rescueFunds(address _token, address _receiverAddress, uint256 _amount) returns()
func (_Registry *RegistrySession) RescueFunds(_token common.Address, _receiverAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.RescueFunds(&_Registry.TransactOpts, _token, _receiverAddress, _amount)
}

// RescueFunds is a paid mutator transaction binding the contract method 0x6ccae054.
//
// Solidity: function rescueFunds(address _token, address _receiverAddress, uint256 _amount) returns()
func (_Registry *RegistryTransactorSession) RescueFunds(_token common.Address, _receiverAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.RescueFunds(&_Registry.TransactOpts, _token, _receiverAddress, _amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Registry *RegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Registry *RegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Registry.Contract.TransferOwnership(&_Registry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Registry *RegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Registry.Contract.TransferOwnership(&_Registry.TransactOpts, newOwner)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Registry *RegistryTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Registry *RegistrySession) Receive() (*types.Transaction, error) {
	return _Registry.Contract.Receive(&_Registry.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Registry *RegistryTransactorSession) Receive() (*types.Transaction, error) {
	return _Registry.Contract.Receive(&_Registry.TransactOpts)
}

// RegistryExecutionCompletedIterator is returned from FilterExecutionCompleted and is used to iterate over the raw logs and unpacked data for ExecutionCompleted events raised by the Registry contract.
type RegistryExecutionCompletedIterator struct {
	Event *RegistryExecutionCompleted // Event containing the contract specifics and raw log

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
func (it *RegistryExecutionCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryExecutionCompleted)
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
		it.Event = new(RegistryExecutionCompleted)
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
func (it *RegistryExecutionCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryExecutionCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryExecutionCompleted represents a ExecutionCompleted event raised by the Registry contract.
type RegistryExecutionCompleted struct {
	MiddlewareID *big.Int
	BridgeID     *big.Int
	InputAmount  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterExecutionCompleted is a free log retrieval operation binding the contract event 0x28fd8a5dda29b4035905e0657f97244a0e0bef97951e248ed0f2c6878d6590c2.
//
// Solidity: event ExecutionCompleted(uint256 middlewareID, uint256 bridgeID, uint256 inputAmount)
func (_Registry *RegistryFilterer) FilterExecutionCompleted(opts *bind.FilterOpts) (*RegistryExecutionCompletedIterator, error) {

	logs, sub, err := _Registry.contract.FilterLogs(opts, "ExecutionCompleted")
	if err != nil {
		return nil, err
	}
	return &RegistryExecutionCompletedIterator{contract: _Registry.contract, event: "ExecutionCompleted", logs: logs, sub: sub}, nil
}

// WatchExecutionCompleted is a free log subscription operation binding the contract event 0x28fd8a5dda29b4035905e0657f97244a0e0bef97951e248ed0f2c6878d6590c2.
//
// Solidity: event ExecutionCompleted(uint256 middlewareID, uint256 bridgeID, uint256 inputAmount)
func (_Registry *RegistryFilterer) WatchExecutionCompleted(opts *bind.WatchOpts, sink chan<- *RegistryExecutionCompleted) (event.Subscription, error) {

	logs, sub, err := _Registry.contract.WatchLogs(opts, "ExecutionCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryExecutionCompleted)
				if err := _Registry.contract.UnpackLog(event, "ExecutionCompleted", log); err != nil {
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

// ParseExecutionCompleted is a log parse operation binding the contract event 0x28fd8a5dda29b4035905e0657f97244a0e0bef97951e248ed0f2c6878d6590c2.
//
// Solidity: event ExecutionCompleted(uint256 middlewareID, uint256 bridgeID, uint256 inputAmount)
func (_Registry *RegistryFilterer) ParseExecutionCompleted(log types.Log) (*RegistryExecutionCompleted, error) {
	event := new(RegistryExecutionCompleted)
	if err := _Registry.contract.UnpackLog(event, "ExecutionCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryNewRouteAddedIterator is returned from FilterNewRouteAdded and is used to iterate over the raw logs and unpacked data for NewRouteAdded events raised by the Registry contract.
type RegistryNewRouteAddedIterator struct {
	Event *RegistryNewRouteAdded // Event containing the contract specifics and raw log

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
func (it *RegistryNewRouteAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryNewRouteAdded)
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
		it.Event = new(RegistryNewRouteAdded)
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
func (it *RegistryNewRouteAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryNewRouteAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryNewRouteAdded represents a NewRouteAdded event raised by the Registry contract.
type RegistryNewRouteAdded struct {
	RouteID      *big.Int
	Route        common.Address
	IsEnabled    bool
	IsMiddleware bool
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNewRouteAdded is a free log retrieval operation binding the contract event 0xd7b1a492030c0a30b288b91bf46f20c49c7715b0dd6d42244c61c540111256b5.
//
// Solidity: event NewRouteAdded(uint256 routeID, address route, bool isEnabled, bool isMiddleware)
func (_Registry *RegistryFilterer) FilterNewRouteAdded(opts *bind.FilterOpts) (*RegistryNewRouteAddedIterator, error) {

	logs, sub, err := _Registry.contract.FilterLogs(opts, "NewRouteAdded")
	if err != nil {
		return nil, err
	}
	return &RegistryNewRouteAddedIterator{contract: _Registry.contract, event: "NewRouteAdded", logs: logs, sub: sub}, nil
}

// WatchNewRouteAdded is a free log subscription operation binding the contract event 0xd7b1a492030c0a30b288b91bf46f20c49c7715b0dd6d42244c61c540111256b5.
//
// Solidity: event NewRouteAdded(uint256 routeID, address route, bool isEnabled, bool isMiddleware)
func (_Registry *RegistryFilterer) WatchNewRouteAdded(opts *bind.WatchOpts, sink chan<- *RegistryNewRouteAdded) (event.Subscription, error) {

	logs, sub, err := _Registry.contract.WatchLogs(opts, "NewRouteAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryNewRouteAdded)
				if err := _Registry.contract.UnpackLog(event, "NewRouteAdded", log); err != nil {
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

// ParseNewRouteAdded is a log parse operation binding the contract event 0xd7b1a492030c0a30b288b91bf46f20c49c7715b0dd6d42244c61c540111256b5.
//
// Solidity: event NewRouteAdded(uint256 routeID, address route, bool isEnabled, bool isMiddleware)
func (_Registry *RegistryFilterer) ParseNewRouteAdded(log types.Log) (*RegistryNewRouteAdded, error) {
	event := new(RegistryNewRouteAdded)
	if err := _Registry.contract.UnpackLog(event, "NewRouteAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Registry contract.
type RegistryOwnershipTransferredIterator struct {
	Event *RegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryOwnershipTransferred)
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
		it.Event = new(RegistryOwnershipTransferred)
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
func (it *RegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryOwnershipTransferred represents a OwnershipTransferred event raised by the Registry contract.
type RegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Registry *RegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RegistryOwnershipTransferredIterator{contract: _Registry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Registry *RegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryOwnershipTransferred)
				if err := _Registry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Registry *RegistryFilterer) ParseOwnershipTransferred(log types.Log) (*RegistryOwnershipTransferred, error) {
	event := new(RegistryOwnershipTransferred)
	if err := _Registry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryRouteDisabledIterator is returned from FilterRouteDisabled and is used to iterate over the raw logs and unpacked data for RouteDisabled events raised by the Registry contract.
type RegistryRouteDisabledIterator struct {
	Event *RegistryRouteDisabled // Event containing the contract specifics and raw log

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
func (it *RegistryRouteDisabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryRouteDisabled)
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
		it.Event = new(RegistryRouteDisabled)
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
func (it *RegistryRouteDisabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryRouteDisabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryRouteDisabled represents a RouteDisabled event raised by the Registry contract.
type RegistryRouteDisabled struct {
	RouteID *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRouteDisabled is a free log retrieval operation binding the contract event 0x91a0168fe2af7d03fc4465ab611da7d997fe924f69c20e9d16a23e6fc7af88d4.
//
// Solidity: event RouteDisabled(uint256 routeID)
func (_Registry *RegistryFilterer) FilterRouteDisabled(opts *bind.FilterOpts) (*RegistryRouteDisabledIterator, error) {

	logs, sub, err := _Registry.contract.FilterLogs(opts, "RouteDisabled")
	if err != nil {
		return nil, err
	}
	return &RegistryRouteDisabledIterator{contract: _Registry.contract, event: "RouteDisabled", logs: logs, sub: sub}, nil
}

// WatchRouteDisabled is a free log subscription operation binding the contract event 0x91a0168fe2af7d03fc4465ab611da7d997fe924f69c20e9d16a23e6fc7af88d4.
//
// Solidity: event RouteDisabled(uint256 routeID)
func (_Registry *RegistryFilterer) WatchRouteDisabled(opts *bind.WatchOpts, sink chan<- *RegistryRouteDisabled) (event.Subscription, error) {

	logs, sub, err := _Registry.contract.WatchLogs(opts, "RouteDisabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryRouteDisabled)
				if err := _Registry.contract.UnpackLog(event, "RouteDisabled", log); err != nil {
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

// ParseRouteDisabled is a log parse operation binding the contract event 0x91a0168fe2af7d03fc4465ab611da7d997fe924f69c20e9d16a23e6fc7af88d4.
//
// Solidity: event RouteDisabled(uint256 routeID)
func (_Registry *RegistryFilterer) ParseRouteDisabled(log types.Log) (*RegistryRouteDisabled, error) {
	event := new(RegistryRouteDisabled)
	if err := _Registry.contract.UnpackLog(event, "RouteDisabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
