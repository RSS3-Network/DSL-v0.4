// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gitcoin

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

// BulkCheckoutDonation is an auto generated low-level Go binding around an user-defined struct.
type BulkCheckoutDonation struct {
	Token  common.Address
	Amount *big.Int
	Dest   common.Address
}

// GitcoinMetaData contains all meta data concerning the Gitcoin contract.
var GitcoinMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"dest\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"donor\",\"type\":\"address\"}],\"name\":\"DonationSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dest\",\"type\":\"address\"}],\"name\":\"TokenWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"dest\",\"type\":\"address\"}],\"internalType\":\"structBulkCheckout.Donation[]\",\"name\":\"_donations\",\"type\":\"tuple[]\"}],\"name\":\"donate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_dest\",\"type\":\"address\"}],\"name\":\"withdrawEther\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_dest\",\"type\":\"address\"}],\"name\":\"withdrawToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// GitcoinABI is the input ABI used to generate the binding from.
// Deprecated: Use GitcoinMetaData.ABI instead.
var GitcoinABI = GitcoinMetaData.ABI

// Gitcoin is an auto generated Go binding around an Ethereum contract.
type Gitcoin struct {
	GitcoinCaller     // Read-only binding to the contract
	GitcoinTransactor // Write-only binding to the contract
	GitcoinFilterer   // Log filterer for contract events
}

// GitcoinCaller is an auto generated read-only Go binding around an Ethereum contract.
type GitcoinCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GitcoinTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GitcoinTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GitcoinFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GitcoinFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GitcoinSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GitcoinSession struct {
	Contract     *Gitcoin          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GitcoinCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GitcoinCallerSession struct {
	Contract *GitcoinCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// GitcoinTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GitcoinTransactorSession struct {
	Contract     *GitcoinTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// GitcoinRaw is an auto generated low-level Go binding around an Ethereum contract.
type GitcoinRaw struct {
	Contract *Gitcoin // Generic contract binding to access the raw methods on
}

// GitcoinCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GitcoinCallerRaw struct {
	Contract *GitcoinCaller // Generic read-only contract binding to access the raw methods on
}

// GitcoinTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GitcoinTransactorRaw struct {
	Contract *GitcoinTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGitcoin creates a new instance of Gitcoin, bound to a specific deployed contract.
func NewGitcoin(address common.Address, backend bind.ContractBackend) (*Gitcoin, error) {
	contract, err := bindGitcoin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Gitcoin{GitcoinCaller: GitcoinCaller{contract: contract}, GitcoinTransactor: GitcoinTransactor{contract: contract}, GitcoinFilterer: GitcoinFilterer{contract: contract}}, nil
}

// NewGitcoinCaller creates a new read-only instance of Gitcoin, bound to a specific deployed contract.
func NewGitcoinCaller(address common.Address, caller bind.ContractCaller) (*GitcoinCaller, error) {
	contract, err := bindGitcoin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GitcoinCaller{contract: contract}, nil
}

// NewGitcoinTransactor creates a new write-only instance of Gitcoin, bound to a specific deployed contract.
func NewGitcoinTransactor(address common.Address, transactor bind.ContractTransactor) (*GitcoinTransactor, error) {
	contract, err := bindGitcoin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GitcoinTransactor{contract: contract}, nil
}

// NewGitcoinFilterer creates a new log filterer instance of Gitcoin, bound to a specific deployed contract.
func NewGitcoinFilterer(address common.Address, filterer bind.ContractFilterer) (*GitcoinFilterer, error) {
	contract, err := bindGitcoin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GitcoinFilterer{contract: contract}, nil
}

// bindGitcoin binds a generic wrapper to an already deployed contract.
func bindGitcoin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GitcoinABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gitcoin *GitcoinRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Gitcoin.Contract.GitcoinCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gitcoin *GitcoinRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gitcoin.Contract.GitcoinTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gitcoin *GitcoinRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gitcoin.Contract.GitcoinTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gitcoin *GitcoinCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Gitcoin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gitcoin *GitcoinTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gitcoin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gitcoin *GitcoinTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gitcoin.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Gitcoin *GitcoinCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Gitcoin.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Gitcoin *GitcoinSession) Owner() (common.Address, error) {
	return _Gitcoin.Contract.Owner(&_Gitcoin.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Gitcoin *GitcoinCallerSession) Owner() (common.Address, error) {
	return _Gitcoin.Contract.Owner(&_Gitcoin.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Gitcoin *GitcoinCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Gitcoin.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Gitcoin *GitcoinSession) Paused() (bool, error) {
	return _Gitcoin.Contract.Paused(&_Gitcoin.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Gitcoin *GitcoinCallerSession) Paused() (bool, error) {
	return _Gitcoin.Contract.Paused(&_Gitcoin.CallOpts)
}

// Donate is a paid mutator transaction binding the contract method 0x9120491c.
//
// Solidity: function donate((address,uint256,address)[] _donations) payable returns()
func (_Gitcoin *GitcoinTransactor) Donate(opts *bind.TransactOpts, _donations []BulkCheckoutDonation) (*types.Transaction, error) {
	return _Gitcoin.contract.Transact(opts, "donate", _donations)
}

// Donate is a paid mutator transaction binding the contract method 0x9120491c.
//
// Solidity: function donate((address,uint256,address)[] _donations) payable returns()
func (_Gitcoin *GitcoinSession) Donate(_donations []BulkCheckoutDonation) (*types.Transaction, error) {
	return _Gitcoin.Contract.Donate(&_Gitcoin.TransactOpts, _donations)
}

// Donate is a paid mutator transaction binding the contract method 0x9120491c.
//
// Solidity: function donate((address,uint256,address)[] _donations) payable returns()
func (_Gitcoin *GitcoinTransactorSession) Donate(_donations []BulkCheckoutDonation) (*types.Transaction, error) {
	return _Gitcoin.Contract.Donate(&_Gitcoin.TransactOpts, _donations)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Gitcoin *GitcoinTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gitcoin.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Gitcoin *GitcoinSession) Pause() (*types.Transaction, error) {
	return _Gitcoin.Contract.Pause(&_Gitcoin.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Gitcoin *GitcoinTransactorSession) Pause() (*types.Transaction, error) {
	return _Gitcoin.Contract.Pause(&_Gitcoin.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Gitcoin *GitcoinTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gitcoin.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Gitcoin *GitcoinSession) RenounceOwnership() (*types.Transaction, error) {
	return _Gitcoin.Contract.RenounceOwnership(&_Gitcoin.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Gitcoin *GitcoinTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Gitcoin.Contract.RenounceOwnership(&_Gitcoin.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Gitcoin *GitcoinTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Gitcoin.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Gitcoin *GitcoinSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Gitcoin.Contract.TransferOwnership(&_Gitcoin.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Gitcoin *GitcoinTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Gitcoin.Contract.TransferOwnership(&_Gitcoin.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Gitcoin *GitcoinTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gitcoin.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Gitcoin *GitcoinSession) Unpause() (*types.Transaction, error) {
	return _Gitcoin.Contract.Unpause(&_Gitcoin.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Gitcoin *GitcoinTransactorSession) Unpause() (*types.Transaction, error) {
	return _Gitcoin.Contract.Unpause(&_Gitcoin.TransactOpts)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0xaf933b57.
//
// Solidity: function withdrawEther(address _dest) returns()
func (_Gitcoin *GitcoinTransactor) WithdrawEther(opts *bind.TransactOpts, _dest common.Address) (*types.Transaction, error) {
	return _Gitcoin.contract.Transact(opts, "withdrawEther", _dest)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0xaf933b57.
//
// Solidity: function withdrawEther(address _dest) returns()
func (_Gitcoin *GitcoinSession) WithdrawEther(_dest common.Address) (*types.Transaction, error) {
	return _Gitcoin.Contract.WithdrawEther(&_Gitcoin.TransactOpts, _dest)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0xaf933b57.
//
// Solidity: function withdrawEther(address _dest) returns()
func (_Gitcoin *GitcoinTransactorSession) WithdrawEther(_dest common.Address) (*types.Transaction, error) {
	return _Gitcoin.Contract.WithdrawEther(&_Gitcoin.TransactOpts, _dest)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x3aeac4e1.
//
// Solidity: function withdrawToken(address _tokenAddress, address _dest) returns()
func (_Gitcoin *GitcoinTransactor) WithdrawToken(opts *bind.TransactOpts, _tokenAddress common.Address, _dest common.Address) (*types.Transaction, error) {
	return _Gitcoin.contract.Transact(opts, "withdrawToken", _tokenAddress, _dest)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x3aeac4e1.
//
// Solidity: function withdrawToken(address _tokenAddress, address _dest) returns()
func (_Gitcoin *GitcoinSession) WithdrawToken(_tokenAddress common.Address, _dest common.Address) (*types.Transaction, error) {
	return _Gitcoin.Contract.WithdrawToken(&_Gitcoin.TransactOpts, _tokenAddress, _dest)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x3aeac4e1.
//
// Solidity: function withdrawToken(address _tokenAddress, address _dest) returns()
func (_Gitcoin *GitcoinTransactorSession) WithdrawToken(_tokenAddress common.Address, _dest common.Address) (*types.Transaction, error) {
	return _Gitcoin.Contract.WithdrawToken(&_Gitcoin.TransactOpts, _tokenAddress, _dest)
}

// GitcoinDonationSentIterator is returned from FilterDonationSent and is used to iterate over the raw logs and unpacked data for DonationSent events raised by the Gitcoin contract.
type GitcoinDonationSentIterator struct {
	Event *GitcoinDonationSent // Event containing the contract specifics and raw log

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
func (it *GitcoinDonationSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GitcoinDonationSent)
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
		it.Event = new(GitcoinDonationSent)
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
func (it *GitcoinDonationSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GitcoinDonationSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GitcoinDonationSent represents a DonationSent event raised by the Gitcoin contract.
type GitcoinDonationSent struct {
	Token  common.Address
	Amount *big.Int
	Dest   common.Address
	Donor  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDonationSent is a free log retrieval operation binding the contract event 0x3bb7428b25f9bdad9bd2faa4c6a7a9e5d5882657e96c1d24cc41c1d6c1910a98.
//
// Solidity: event DonationSent(address indexed token, uint256 indexed amount, address dest, address indexed donor)
func (_Gitcoin *GitcoinFilterer) FilterDonationSent(opts *bind.FilterOpts, token []common.Address, amount []*big.Int, donor []common.Address) (*GitcoinDonationSentIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	var donorRule []interface{}
	for _, donorItem := range donor {
		donorRule = append(donorRule, donorItem)
	}

	logs, sub, err := _Gitcoin.contract.FilterLogs(opts, "DonationSent", tokenRule, amountRule, donorRule)
	if err != nil {
		return nil, err
	}
	return &GitcoinDonationSentIterator{contract: _Gitcoin.contract, event: "DonationSent", logs: logs, sub: sub}, nil
}

// WatchDonationSent is a free log subscription operation binding the contract event 0x3bb7428b25f9bdad9bd2faa4c6a7a9e5d5882657e96c1d24cc41c1d6c1910a98.
//
// Solidity: event DonationSent(address indexed token, uint256 indexed amount, address dest, address indexed donor)
func (_Gitcoin *GitcoinFilterer) WatchDonationSent(opts *bind.WatchOpts, sink chan<- *GitcoinDonationSent, token []common.Address, amount []*big.Int, donor []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	var donorRule []interface{}
	for _, donorItem := range donor {
		donorRule = append(donorRule, donorItem)
	}

	logs, sub, err := _Gitcoin.contract.WatchLogs(opts, "DonationSent", tokenRule, amountRule, donorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GitcoinDonationSent)
				if err := _Gitcoin.contract.UnpackLog(event, "DonationSent", log); err != nil {
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

// ParseDonationSent is a log parse operation binding the contract event 0x3bb7428b25f9bdad9bd2faa4c6a7a9e5d5882657e96c1d24cc41c1d6c1910a98.
//
// Solidity: event DonationSent(address indexed token, uint256 indexed amount, address dest, address indexed donor)
func (_Gitcoin *GitcoinFilterer) ParseDonationSent(log types.Log) (*GitcoinDonationSent, error) {
	event := new(GitcoinDonationSent)
	if err := _Gitcoin.contract.UnpackLog(event, "DonationSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GitcoinOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Gitcoin contract.
type GitcoinOwnershipTransferredIterator struct {
	Event *GitcoinOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *GitcoinOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GitcoinOwnershipTransferred)
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
		it.Event = new(GitcoinOwnershipTransferred)
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
func (it *GitcoinOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GitcoinOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GitcoinOwnershipTransferred represents a OwnershipTransferred event raised by the Gitcoin contract.
type GitcoinOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Gitcoin *GitcoinFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*GitcoinOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Gitcoin.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &GitcoinOwnershipTransferredIterator{contract: _Gitcoin.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Gitcoin *GitcoinFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *GitcoinOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Gitcoin.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GitcoinOwnershipTransferred)
				if err := _Gitcoin.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Gitcoin *GitcoinFilterer) ParseOwnershipTransferred(log types.Log) (*GitcoinOwnershipTransferred, error) {
	event := new(GitcoinOwnershipTransferred)
	if err := _Gitcoin.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GitcoinPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Gitcoin contract.
type GitcoinPausedIterator struct {
	Event *GitcoinPaused // Event containing the contract specifics and raw log

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
func (it *GitcoinPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GitcoinPaused)
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
		it.Event = new(GitcoinPaused)
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
func (it *GitcoinPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GitcoinPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GitcoinPaused represents a Paused event raised by the Gitcoin contract.
type GitcoinPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Gitcoin *GitcoinFilterer) FilterPaused(opts *bind.FilterOpts) (*GitcoinPausedIterator, error) {

	logs, sub, err := _Gitcoin.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &GitcoinPausedIterator{contract: _Gitcoin.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Gitcoin *GitcoinFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *GitcoinPaused) (event.Subscription, error) {

	logs, sub, err := _Gitcoin.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GitcoinPaused)
				if err := _Gitcoin.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Gitcoin *GitcoinFilterer) ParsePaused(log types.Log) (*GitcoinPaused, error) {
	event := new(GitcoinPaused)
	if err := _Gitcoin.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GitcoinTokenWithdrawnIterator is returned from FilterTokenWithdrawn and is used to iterate over the raw logs and unpacked data for TokenWithdrawn events raised by the Gitcoin contract.
type GitcoinTokenWithdrawnIterator struct {
	Event *GitcoinTokenWithdrawn // Event containing the contract specifics and raw log

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
func (it *GitcoinTokenWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GitcoinTokenWithdrawn)
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
		it.Event = new(GitcoinTokenWithdrawn)
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
func (it *GitcoinTokenWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GitcoinTokenWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GitcoinTokenWithdrawn represents a TokenWithdrawn event raised by the Gitcoin contract.
type GitcoinTokenWithdrawn struct {
	Token  common.Address
	Amount *big.Int
	Dest   common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokenWithdrawn is a free log retrieval operation binding the contract event 0xa0524ee0fd8662d6c046d199da2a6d3dc49445182cec055873a5bb9c2843c8e0.
//
// Solidity: event TokenWithdrawn(address indexed token, uint256 indexed amount, address indexed dest)
func (_Gitcoin *GitcoinFilterer) FilterTokenWithdrawn(opts *bind.FilterOpts, token []common.Address, amount []*big.Int, dest []common.Address) (*GitcoinTokenWithdrawnIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}
	var destRule []interface{}
	for _, destItem := range dest {
		destRule = append(destRule, destItem)
	}

	logs, sub, err := _Gitcoin.contract.FilterLogs(opts, "TokenWithdrawn", tokenRule, amountRule, destRule)
	if err != nil {
		return nil, err
	}
	return &GitcoinTokenWithdrawnIterator{contract: _Gitcoin.contract, event: "TokenWithdrawn", logs: logs, sub: sub}, nil
}

// WatchTokenWithdrawn is a free log subscription operation binding the contract event 0xa0524ee0fd8662d6c046d199da2a6d3dc49445182cec055873a5bb9c2843c8e0.
//
// Solidity: event TokenWithdrawn(address indexed token, uint256 indexed amount, address indexed dest)
func (_Gitcoin *GitcoinFilterer) WatchTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *GitcoinTokenWithdrawn, token []common.Address, amount []*big.Int, dest []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}
	var destRule []interface{}
	for _, destItem := range dest {
		destRule = append(destRule, destItem)
	}

	logs, sub, err := _Gitcoin.contract.WatchLogs(opts, "TokenWithdrawn", tokenRule, amountRule, destRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GitcoinTokenWithdrawn)
				if err := _Gitcoin.contract.UnpackLog(event, "TokenWithdrawn", log); err != nil {
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

// ParseTokenWithdrawn is a log parse operation binding the contract event 0xa0524ee0fd8662d6c046d199da2a6d3dc49445182cec055873a5bb9c2843c8e0.
//
// Solidity: event TokenWithdrawn(address indexed token, uint256 indexed amount, address indexed dest)
func (_Gitcoin *GitcoinFilterer) ParseTokenWithdrawn(log types.Log) (*GitcoinTokenWithdrawn, error) {
	event := new(GitcoinTokenWithdrawn)
	if err := _Gitcoin.contract.UnpackLog(event, "TokenWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GitcoinUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Gitcoin contract.
type GitcoinUnpausedIterator struct {
	Event *GitcoinUnpaused // Event containing the contract specifics and raw log

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
func (it *GitcoinUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GitcoinUnpaused)
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
		it.Event = new(GitcoinUnpaused)
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
func (it *GitcoinUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GitcoinUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GitcoinUnpaused represents a Unpaused event raised by the Gitcoin contract.
type GitcoinUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Gitcoin *GitcoinFilterer) FilterUnpaused(opts *bind.FilterOpts) (*GitcoinUnpausedIterator, error) {

	logs, sub, err := _Gitcoin.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &GitcoinUnpausedIterator{contract: _Gitcoin.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Gitcoin *GitcoinFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *GitcoinUnpaused) (event.Subscription, error) {

	logs, sub, err := _Gitcoin.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GitcoinUnpaused)
				if err := _Gitcoin.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Gitcoin *GitcoinFilterer) ParseUnpaused(log types.Log) (*GitcoinUnpaused, error) {
	event := new(GitcoinUnpaused)
	if err := _Gitcoin.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
