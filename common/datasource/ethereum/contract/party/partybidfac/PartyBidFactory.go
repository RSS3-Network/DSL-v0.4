// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package partybidfac

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

// PartyBidFactoryMetaData contains all meta data concerning the PartyBidFactory contract.
var PartyBidFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_partyDAOMultisig\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenVaultFactory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_weth\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_logicMarketWrapper\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_logicNftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_logicTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_logicAuctionId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"partyBidProxy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"marketWrapper\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"splitRecipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"splitBasisPoints\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"gatedToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gatedTokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"name\":\"PartyBidDeployed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"deployedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"logic\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"partyDAOMultisig\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_marketWrapper\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_auctionId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structStructs.AddressAndAmount\",\"name\":\"_split\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structStructs.AddressAndAmount\",\"name\":\"_tokenGate\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"name\":\"startParty\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"partyBidProxy\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenVaultFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"weth\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// PartyBidFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use PartyBidFactoryMetaData.ABI instead.
var PartyBidFactoryABI = PartyBidFactoryMetaData.ABI

// PartyBidFactory is an auto generated Go binding around an Ethereum contract.
type PartyBidFactory struct {
	PartyBidFactoryCaller     // Read-only binding to the contract
	PartyBidFactoryTransactor // Write-only binding to the contract
	PartyBidFactoryFilterer   // Log filterer for contract events
}

// PartyBidFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type PartyBidFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PartyBidFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PartyBidFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PartyBidFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PartyBidFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PartyBidFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PartyBidFactorySession struct {
	Contract     *PartyBidFactory  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PartyBidFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PartyBidFactoryCallerSession struct {
	Contract *PartyBidFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// PartyBidFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PartyBidFactoryTransactorSession struct {
	Contract     *PartyBidFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// PartyBidFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type PartyBidFactoryRaw struct {
	Contract *PartyBidFactory // Generic contract binding to access the raw methods on
}

// PartyBidFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PartyBidFactoryCallerRaw struct {
	Contract *PartyBidFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// PartyBidFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PartyBidFactoryTransactorRaw struct {
	Contract *PartyBidFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPartyBidFactory creates a new instance of PartyBidFactory, bound to a specific deployed contract.
func NewPartyBidFactory(address common.Address, backend bind.ContractBackend) (*PartyBidFactory, error) {
	contract, err := bindPartyBidFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PartyBidFactory{PartyBidFactoryCaller: PartyBidFactoryCaller{contract: contract}, PartyBidFactoryTransactor: PartyBidFactoryTransactor{contract: contract}, PartyBidFactoryFilterer: PartyBidFactoryFilterer{contract: contract}}, nil
}

// NewPartyBidFactoryCaller creates a new read-only instance of PartyBidFactory, bound to a specific deployed contract.
func NewPartyBidFactoryCaller(address common.Address, caller bind.ContractCaller) (*PartyBidFactoryCaller, error) {
	contract, err := bindPartyBidFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PartyBidFactoryCaller{contract: contract}, nil
}

// NewPartyBidFactoryTransactor creates a new write-only instance of PartyBidFactory, bound to a specific deployed contract.
func NewPartyBidFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*PartyBidFactoryTransactor, error) {
	contract, err := bindPartyBidFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PartyBidFactoryTransactor{contract: contract}, nil
}

// NewPartyBidFactoryFilterer creates a new log filterer instance of PartyBidFactory, bound to a specific deployed contract.
func NewPartyBidFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*PartyBidFactoryFilterer, error) {
	contract, err := bindPartyBidFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PartyBidFactoryFilterer{contract: contract}, nil
}

// bindPartyBidFactory binds a generic wrapper to an already deployed contract.
func bindPartyBidFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PartyBidFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PartyBidFactory *PartyBidFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PartyBidFactory.Contract.PartyBidFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PartyBidFactory *PartyBidFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PartyBidFactory.Contract.PartyBidFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PartyBidFactory *PartyBidFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PartyBidFactory.Contract.PartyBidFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PartyBidFactory *PartyBidFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PartyBidFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PartyBidFactory *PartyBidFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PartyBidFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PartyBidFactory *PartyBidFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PartyBidFactory.Contract.contract.Transact(opts, method, params...)
}

// DeployedAt is a free data retrieval call binding the contract method 0x57dbdf54.
//
// Solidity: function deployedAt(address ) view returns(uint256)
func (_PartyBidFactory *PartyBidFactoryCaller) DeployedAt(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PartyBidFactory.contract.Call(opts, &out, "deployedAt", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DeployedAt is a free data retrieval call binding the contract method 0x57dbdf54.
//
// Solidity: function deployedAt(address ) view returns(uint256)
func (_PartyBidFactory *PartyBidFactorySession) DeployedAt(arg0 common.Address) (*big.Int, error) {
	return _PartyBidFactory.Contract.DeployedAt(&_PartyBidFactory.CallOpts, arg0)
}

// DeployedAt is a free data retrieval call binding the contract method 0x57dbdf54.
//
// Solidity: function deployedAt(address ) view returns(uint256)
func (_PartyBidFactory *PartyBidFactoryCallerSession) DeployedAt(arg0 common.Address) (*big.Int, error) {
	return _PartyBidFactory.Contract.DeployedAt(&_PartyBidFactory.CallOpts, arg0)
}

// Logic is a free data retrieval call binding the contract method 0xd7dfa0dd.
//
// Solidity: function logic() view returns(address)
func (_PartyBidFactory *PartyBidFactoryCaller) Logic(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PartyBidFactory.contract.Call(opts, &out, "logic")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Logic is a free data retrieval call binding the contract method 0xd7dfa0dd.
//
// Solidity: function logic() view returns(address)
func (_PartyBidFactory *PartyBidFactorySession) Logic() (common.Address, error) {
	return _PartyBidFactory.Contract.Logic(&_PartyBidFactory.CallOpts)
}

// Logic is a free data retrieval call binding the contract method 0xd7dfa0dd.
//
// Solidity: function logic() view returns(address)
func (_PartyBidFactory *PartyBidFactoryCallerSession) Logic() (common.Address, error) {
	return _PartyBidFactory.Contract.Logic(&_PartyBidFactory.CallOpts)
}

// PartyDAOMultisig is a free data retrieval call binding the contract method 0x3c4d12d9.
//
// Solidity: function partyDAOMultisig() view returns(address)
func (_PartyBidFactory *PartyBidFactoryCaller) PartyDAOMultisig(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PartyBidFactory.contract.Call(opts, &out, "partyDAOMultisig")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PartyDAOMultisig is a free data retrieval call binding the contract method 0x3c4d12d9.
//
// Solidity: function partyDAOMultisig() view returns(address)
func (_PartyBidFactory *PartyBidFactorySession) PartyDAOMultisig() (common.Address, error) {
	return _PartyBidFactory.Contract.PartyDAOMultisig(&_PartyBidFactory.CallOpts)
}

// PartyDAOMultisig is a free data retrieval call binding the contract method 0x3c4d12d9.
//
// Solidity: function partyDAOMultisig() view returns(address)
func (_PartyBidFactory *PartyBidFactoryCallerSession) PartyDAOMultisig() (common.Address, error) {
	return _PartyBidFactory.Contract.PartyDAOMultisig(&_PartyBidFactory.CallOpts)
}

// TokenVaultFactory is a free data retrieval call binding the contract method 0x0b203023.
//
// Solidity: function tokenVaultFactory() view returns(address)
func (_PartyBidFactory *PartyBidFactoryCaller) TokenVaultFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PartyBidFactory.contract.Call(opts, &out, "tokenVaultFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenVaultFactory is a free data retrieval call binding the contract method 0x0b203023.
//
// Solidity: function tokenVaultFactory() view returns(address)
func (_PartyBidFactory *PartyBidFactorySession) TokenVaultFactory() (common.Address, error) {
	return _PartyBidFactory.Contract.TokenVaultFactory(&_PartyBidFactory.CallOpts)
}

// TokenVaultFactory is a free data retrieval call binding the contract method 0x0b203023.
//
// Solidity: function tokenVaultFactory() view returns(address)
func (_PartyBidFactory *PartyBidFactoryCallerSession) TokenVaultFactory() (common.Address, error) {
	return _PartyBidFactory.Contract.TokenVaultFactory(&_PartyBidFactory.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_PartyBidFactory *PartyBidFactoryCaller) Weth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PartyBidFactory.contract.Call(opts, &out, "weth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_PartyBidFactory *PartyBidFactorySession) Weth() (common.Address, error) {
	return _PartyBidFactory.Contract.Weth(&_PartyBidFactory.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_PartyBidFactory *PartyBidFactoryCallerSession) Weth() (common.Address, error) {
	return _PartyBidFactory.Contract.Weth(&_PartyBidFactory.CallOpts)
}

// StartParty is a paid mutator transaction binding the contract method 0xc82816cb.
//
// Solidity: function startParty(address _marketWrapper, address _nftContract, uint256 _tokenId, uint256 _auctionId, (address,uint256) _split, (address,uint256) _tokenGate, string _name, string _symbol) returns(address partyBidProxy)
func (_PartyBidFactory *PartyBidFactoryTransactor) StartParty(opts *bind.TransactOpts, _marketWrapper common.Address, _nftContract common.Address, _tokenId *big.Int, _auctionId *big.Int, _split StructsAddressAndAmount, _tokenGate StructsAddressAndAmount, _name string, _symbol string) (*types.Transaction, error) {
	return _PartyBidFactory.contract.Transact(opts, "startParty", _marketWrapper, _nftContract, _tokenId, _auctionId, _split, _tokenGate, _name, _symbol)
}

// StartParty is a paid mutator transaction binding the contract method 0xc82816cb.
//
// Solidity: function startParty(address _marketWrapper, address _nftContract, uint256 _tokenId, uint256 _auctionId, (address,uint256) _split, (address,uint256) _tokenGate, string _name, string _symbol) returns(address partyBidProxy)
func (_PartyBidFactory *PartyBidFactorySession) StartParty(_marketWrapper common.Address, _nftContract common.Address, _tokenId *big.Int, _auctionId *big.Int, _split StructsAddressAndAmount, _tokenGate StructsAddressAndAmount, _name string, _symbol string) (*types.Transaction, error) {
	return _PartyBidFactory.Contract.StartParty(&_PartyBidFactory.TransactOpts, _marketWrapper, _nftContract, _tokenId, _auctionId, _split, _tokenGate, _name, _symbol)
}

// StartParty is a paid mutator transaction binding the contract method 0xc82816cb.
//
// Solidity: function startParty(address _marketWrapper, address _nftContract, uint256 _tokenId, uint256 _auctionId, (address,uint256) _split, (address,uint256) _tokenGate, string _name, string _symbol) returns(address partyBidProxy)
func (_PartyBidFactory *PartyBidFactoryTransactorSession) StartParty(_marketWrapper common.Address, _nftContract common.Address, _tokenId *big.Int, _auctionId *big.Int, _split StructsAddressAndAmount, _tokenGate StructsAddressAndAmount, _name string, _symbol string) (*types.Transaction, error) {
	return _PartyBidFactory.Contract.StartParty(&_PartyBidFactory.TransactOpts, _marketWrapper, _nftContract, _tokenId, _auctionId, _split, _tokenGate, _name, _symbol)
}

// PartyBidFactoryPartyBidDeployedIterator is returned from FilterPartyBidDeployed and is used to iterate over the raw logs and unpacked data for PartyBidDeployed events raised by the PartyBidFactory contract.
type PartyBidFactoryPartyBidDeployedIterator struct {
	Event *PartyBidFactoryPartyBidDeployed // Event containing the contract specifics and raw log

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
func (it *PartyBidFactoryPartyBidDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PartyBidFactoryPartyBidDeployed)
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
		it.Event = new(PartyBidFactoryPartyBidDeployed)
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
func (it *PartyBidFactoryPartyBidDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PartyBidFactoryPartyBidDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PartyBidFactoryPartyBidDeployed represents a PartyBidDeployed event raised by the PartyBidFactory contract.
type PartyBidFactoryPartyBidDeployed struct {
	PartyBidProxy    common.Address
	Creator          common.Address
	NftContract      common.Address
	TokenId          *big.Int
	MarketWrapper    common.Address
	AuctionId        *big.Int
	SplitRecipient   common.Address
	SplitBasisPoints *big.Int
	GatedToken       common.Address
	GatedTokenAmount *big.Int
	Name             string
	Symbol           string
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterPartyBidDeployed is a free log retrieval operation binding the contract event 0x2e13a1db67858d419104348669d9bb778dcbefae90d4c5b743825c5ddba7b09b.
//
// Solidity: event PartyBidDeployed(address partyBidProxy, address creator, address nftContract, uint256 tokenId, address marketWrapper, uint256 auctionId, address splitRecipient, uint256 splitBasisPoints, address gatedToken, uint256 gatedTokenAmount, string name, string symbol)
func (_PartyBidFactory *PartyBidFactoryFilterer) FilterPartyBidDeployed(opts *bind.FilterOpts) (*PartyBidFactoryPartyBidDeployedIterator, error) {

	logs, sub, err := _PartyBidFactory.contract.FilterLogs(opts, "PartyBidDeployed")
	if err != nil {
		return nil, err
	}
	return &PartyBidFactoryPartyBidDeployedIterator{contract: _PartyBidFactory.contract, event: "PartyBidDeployed", logs: logs, sub: sub}, nil
}

// WatchPartyBidDeployed is a free log subscription operation binding the contract event 0x2e13a1db67858d419104348669d9bb778dcbefae90d4c5b743825c5ddba7b09b.
//
// Solidity: event PartyBidDeployed(address partyBidProxy, address creator, address nftContract, uint256 tokenId, address marketWrapper, uint256 auctionId, address splitRecipient, uint256 splitBasisPoints, address gatedToken, uint256 gatedTokenAmount, string name, string symbol)
func (_PartyBidFactory *PartyBidFactoryFilterer) WatchPartyBidDeployed(opts *bind.WatchOpts, sink chan<- *PartyBidFactoryPartyBidDeployed) (event.Subscription, error) {

	logs, sub, err := _PartyBidFactory.contract.WatchLogs(opts, "PartyBidDeployed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PartyBidFactoryPartyBidDeployed)
				if err := _PartyBidFactory.contract.UnpackLog(event, "PartyBidDeployed", log); err != nil {
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

// ParsePartyBidDeployed is a log parse operation binding the contract event 0x2e13a1db67858d419104348669d9bb778dcbefae90d4c5b743825c5ddba7b09b.
//
// Solidity: event PartyBidDeployed(address partyBidProxy, address creator, address nftContract, uint256 tokenId, address marketWrapper, uint256 auctionId, address splitRecipient, uint256 splitBasisPoints, address gatedToken, uint256 gatedTokenAmount, string name, string symbol)
func (_PartyBidFactory *PartyBidFactoryFilterer) ParsePartyBidDeployed(log types.Log) (*PartyBidFactoryPartyBidDeployed, error) {
	event := new(PartyBidFactoryPartyBidDeployed)
	if err := _PartyBidFactory.contract.UnpackLog(event, "PartyBidDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
