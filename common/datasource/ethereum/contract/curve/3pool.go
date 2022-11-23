// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package curve

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

// ThreePoolMetaData contains all meta data concerning the ThreePool contract.
var ThreePoolMetaData = &bind.MetaData{
	ABI: "[{\"name\":\"TokenExchange\",\"inputs\":[{\"type\":\"address\",\"name\":\"buyer\",\"indexed\":true},{\"type\":\"int128\",\"name\":\"sold_id\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"tokens_sold\",\"indexed\":false},{\"type\":\"int128\",\"name\":\"bought_id\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"tokens_bought\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"AddLiquidity\",\"inputs\":[{\"type\":\"address\",\"name\":\"provider\",\"indexed\":true},{\"type\":\"uint256[3]\",\"name\":\"token_amounts\",\"indexed\":false},{\"type\":\"uint256[3]\",\"name\":\"fees\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"invariant\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"token_supply\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RemoveLiquidity\",\"inputs\":[{\"type\":\"address\",\"name\":\"provider\",\"indexed\":true},{\"type\":\"uint256[3]\",\"name\":\"token_amounts\",\"indexed\":false},{\"type\":\"uint256[3]\",\"name\":\"fees\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"token_supply\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RemoveLiquidityOne\",\"inputs\":[{\"type\":\"address\",\"name\":\"provider\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"token_amount\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"coin_amount\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RemoveLiquidityImbalance\",\"inputs\":[{\"type\":\"address\",\"name\":\"provider\",\"indexed\":true},{\"type\":\"uint256[3]\",\"name\":\"token_amounts\",\"indexed\":false},{\"type\":\"uint256[3]\",\"name\":\"fees\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"invariant\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"token_supply\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"CommitNewAdmin\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"deadline\",\"indexed\":true},{\"type\":\"address\",\"name\":\"admin\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"NewAdmin\",\"inputs\":[{\"type\":\"address\",\"name\":\"admin\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"CommitNewFee\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"deadline\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"fee\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"admin_fee\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"NewFee\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"fee\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"admin_fee\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RampA\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"old_A\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"new_A\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"initial_time\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"future_time\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"StopRampA\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"A\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"t\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_owner\"},{\"type\":\"address[3]\",\"name\":\"_coins\"},{\"type\":\"address\",\"name\":\"_pool_token\"},{\"type\":\"uint256\",\"name\":\"_A\"},{\"type\":\"uint256\",\"name\":\"_fee\"},{\"type\":\"uint256\",\"name\":\"_admin_fee\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"name\":\"A\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":5227},{\"name\":\"get_virtual_price\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1133537},{\"name\":\"calc_token_amount\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256[3]\",\"name\":\"amounts\"},{\"type\":\"bool\",\"name\":\"deposit\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":4508776},{\"name\":\"add_liquidity\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256[3]\",\"name\":\"amounts\"},{\"type\":\"uint256\",\"name\":\"min_mint_amount\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":6954858},{\"name\":\"get_dy\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"int128\",\"name\":\"i\"},{\"type\":\"int128\",\"name\":\"j\"},{\"type\":\"uint256\",\"name\":\"dx\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2673791},{\"name\":\"get_dy_underlying\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"int128\",\"name\":\"i\"},{\"type\":\"int128\",\"name\":\"j\"},{\"type\":\"uint256\",\"name\":\"dx\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2673474},{\"name\":\"exchange\",\"outputs\":[],\"inputs\":[{\"type\":\"int128\",\"name\":\"i\"},{\"type\":\"int128\",\"name\":\"j\"},{\"type\":\"uint256\",\"name\":\"dx\"},{\"type\":\"uint256\",\"name\":\"min_dy\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":2818066},{\"name\":\"remove_liquidity\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"name\":\"_amount\"},{\"type\":\"uint256[3]\",\"name\":\"min_amounts\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":192846},{\"name\":\"remove_liquidity_imbalance\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256[3]\",\"name\":\"amounts\"},{\"type\":\"uint256\",\"name\":\"max_burn_amount\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":6951851},{\"name\":\"calc_withdraw_one_coin\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"_token_amount\"},{\"type\":\"int128\",\"name\":\"i\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1102},{\"name\":\"remove_liquidity_one_coin\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"name\":\"_token_amount\"},{\"type\":\"int128\",\"name\":\"i\"},{\"type\":\"uint256\",\"name\":\"min_amount\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":4025523},{\"name\":\"ramp_A\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"name\":\"_future_A\"},{\"type\":\"uint256\",\"name\":\"_future_time\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":151919},{\"name\":\"stop_ramp_A\",\"outputs\":[],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":148637},{\"name\":\"commit_new_fee\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"name\":\"new_fee\"},{\"type\":\"uint256\",\"name\":\"new_admin_fee\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":110461},{\"name\":\"apply_new_fee\",\"outputs\":[],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":97242},{\"name\":\"revert_new_parameters\",\"outputs\":[],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":21895},{\"name\":\"commit_transfer_ownership\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_owner\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":74572},{\"name\":\"apply_transfer_ownership\",\"outputs\":[],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":60710},{\"name\":\"revert_transfer_ownership\",\"outputs\":[],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":21985},{\"name\":\"admin_balances\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"i\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":3481},{\"name\":\"withdraw_admin_fees\",\"outputs\":[],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":21502},{\"name\":\"donate_admin_fees\",\"outputs\":[],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":111389},{\"name\":\"kill_me\",\"outputs\":[],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":37998},{\"name\":\"unkill_me\",\"outputs\":[],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":22135},{\"name\":\"coins\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2220},{\"name\":\"balances\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2250},{\"name\":\"fee\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2171},{\"name\":\"admin_fee\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2201},{\"name\":\"owner\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2231},{\"name\":\"initial_A\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2261},{\"name\":\"future_A\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2291},{\"name\":\"initial_A_time\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2321},{\"name\":\"future_A_time\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2351},{\"name\":\"admin_actions_deadline\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2381},{\"name\":\"transfer_ownership_deadline\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2411},{\"name\":\"future_fee\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2441},{\"name\":\"future_admin_fee\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2471},{\"name\":\"future_owner\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2501}]",
}

// ThreePoolABI is the input ABI used to generate the binding from.
// Deprecated: Use ThreePoolMetaData.ABI instead.
var ThreePoolABI = ThreePoolMetaData.ABI

// ThreePool is an auto generated Go binding around an Ethereum contract.
type ThreePool struct {
	ThreePoolCaller     // Read-only binding to the contract
	ThreePoolTransactor // Write-only binding to the contract
	ThreePoolFilterer   // Log filterer for contract events
}

// ThreePoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type ThreePoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ThreePoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ThreePoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ThreePoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ThreePoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ThreePoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ThreePoolSession struct {
	Contract     *ThreePool        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ThreePoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ThreePoolCallerSession struct {
	Contract *ThreePoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ThreePoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ThreePoolTransactorSession struct {
	Contract     *ThreePoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ThreePoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type ThreePoolRaw struct {
	Contract *ThreePool // Generic contract binding to access the raw methods on
}

// ThreePoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ThreePoolCallerRaw struct {
	Contract *ThreePoolCaller // Generic read-only contract binding to access the raw methods on
}

// ThreePoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ThreePoolTransactorRaw struct {
	Contract *ThreePoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewThreePool creates a new instance of ThreePool, bound to a specific deployed contract.
func NewThreePool(address common.Address, backend bind.ContractBackend) (*ThreePool, error) {
	contract, err := bindThreePool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ThreePool{ThreePoolCaller: ThreePoolCaller{contract: contract}, ThreePoolTransactor: ThreePoolTransactor{contract: contract}, ThreePoolFilterer: ThreePoolFilterer{contract: contract}}, nil
}

// NewThreePoolCaller creates a new read-only instance of ThreePool, bound to a specific deployed contract.
func NewThreePoolCaller(address common.Address, caller bind.ContractCaller) (*ThreePoolCaller, error) {
	contract, err := bindThreePool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ThreePoolCaller{contract: contract}, nil
}

// NewThreePoolTransactor creates a new write-only instance of ThreePool, bound to a specific deployed contract.
func NewThreePoolTransactor(address common.Address, transactor bind.ContractTransactor) (*ThreePoolTransactor, error) {
	contract, err := bindThreePool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ThreePoolTransactor{contract: contract}, nil
}

// NewThreePoolFilterer creates a new log filterer instance of ThreePool, bound to a specific deployed contract.
func NewThreePoolFilterer(address common.Address, filterer bind.ContractFilterer) (*ThreePoolFilterer, error) {
	contract, err := bindThreePool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ThreePoolFilterer{contract: contract}, nil
}

// bindThreePool binds a generic wrapper to an already deployed contract.
func bindThreePool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ThreePoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ThreePool *ThreePoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ThreePool.Contract.ThreePoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ThreePool *ThreePoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ThreePool.Contract.ThreePoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ThreePool *ThreePoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ThreePool.Contract.ThreePoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ThreePool *ThreePoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ThreePool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ThreePool *ThreePoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ThreePool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ThreePool *ThreePoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ThreePool.Contract.contract.Transact(opts, method, params...)
}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_ThreePool *ThreePoolCaller) A(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ThreePool.contract.Call(opts, &out, "A")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_ThreePool *ThreePoolSession) A() (*big.Int, error) {
	return _ThreePool.Contract.A(&_ThreePool.CallOpts)
}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_ThreePool *ThreePoolCallerSession) A() (*big.Int, error) {
	return _ThreePool.Contract.A(&_ThreePool.CallOpts)
}

// AdminActionsDeadline is a free data retrieval call binding the contract method 0x405e28f8.
//
// Solidity: function admin_actions_deadline() view returns(uint256)
func (_ThreePool *ThreePoolCaller) AdminActionsDeadline(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ThreePool.contract.Call(opts, &out, "admin_actions_deadline")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdminActionsDeadline is a free data retrieval call binding the contract method 0x405e28f8.
//
// Solidity: function admin_actions_deadline() view returns(uint256)
func (_ThreePool *ThreePoolSession) AdminActionsDeadline() (*big.Int, error) {
	return _ThreePool.Contract.AdminActionsDeadline(&_ThreePool.CallOpts)
}

// AdminActionsDeadline is a free data retrieval call binding the contract method 0x405e28f8.
//
// Solidity: function admin_actions_deadline() view returns(uint256)
func (_ThreePool *ThreePoolCallerSession) AdminActionsDeadline() (*big.Int, error) {
	return _ThreePool.Contract.AdminActionsDeadline(&_ThreePool.CallOpts)
}

// AdminBalances is a free data retrieval call binding the contract method 0xe2e7d264.
//
// Solidity: function admin_balances(uint256 i) view returns(uint256)
func (_ThreePool *ThreePoolCaller) AdminBalances(opts *bind.CallOpts, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ThreePool.contract.Call(opts, &out, "admin_balances", i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdminBalances is a free data retrieval call binding the contract method 0xe2e7d264.
//
// Solidity: function admin_balances(uint256 i) view returns(uint256)
func (_ThreePool *ThreePoolSession) AdminBalances(i *big.Int) (*big.Int, error) {
	return _ThreePool.Contract.AdminBalances(&_ThreePool.CallOpts, i)
}

// AdminBalances is a free data retrieval call binding the contract method 0xe2e7d264.
//
// Solidity: function admin_balances(uint256 i) view returns(uint256)
func (_ThreePool *ThreePoolCallerSession) AdminBalances(i *big.Int) (*big.Int, error) {
	return _ThreePool.Contract.AdminBalances(&_ThreePool.CallOpts, i)
}

// AdminFee is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() view returns(uint256)
func (_ThreePool *ThreePoolCaller) AdminFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ThreePool.contract.Call(opts, &out, "admin_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdminFee is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() view returns(uint256)
func (_ThreePool *ThreePoolSession) AdminFee() (*big.Int, error) {
	return _ThreePool.Contract.AdminFee(&_ThreePool.CallOpts)
}

// AdminFee is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() view returns(uint256)
func (_ThreePool *ThreePoolCallerSession) AdminFee() (*big.Int, error) {
	return _ThreePool.Contract.AdminFee(&_ThreePool.CallOpts)
}

// Balances is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 arg0) view returns(uint256)
func (_ThreePool *ThreePoolCaller) Balances(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ThreePool.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 arg0) view returns(uint256)
func (_ThreePool *ThreePoolSession) Balances(arg0 *big.Int) (*big.Int, error) {
	return _ThreePool.Contract.Balances(&_ThreePool.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 arg0) view returns(uint256)
func (_ThreePool *ThreePoolCallerSession) Balances(arg0 *big.Int) (*big.Int, error) {
	return _ThreePool.Contract.Balances(&_ThreePool.CallOpts, arg0)
}

// CalcTokenAmount is a free data retrieval call binding the contract method 0x3883e119.
//
// Solidity: function calc_token_amount(uint256[3] amounts, bool deposit) view returns(uint256)
func (_ThreePool *ThreePoolCaller) CalcTokenAmount(opts *bind.CallOpts, amounts [3]*big.Int, deposit bool) (*big.Int, error) {
	var out []interface{}
	err := _ThreePool.contract.Call(opts, &out, "calc_token_amount", amounts, deposit)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcTokenAmount is a free data retrieval call binding the contract method 0x3883e119.
//
// Solidity: function calc_token_amount(uint256[3] amounts, bool deposit) view returns(uint256)
func (_ThreePool *ThreePoolSession) CalcTokenAmount(amounts [3]*big.Int, deposit bool) (*big.Int, error) {
	return _ThreePool.Contract.CalcTokenAmount(&_ThreePool.CallOpts, amounts, deposit)
}

// CalcTokenAmount is a free data retrieval call binding the contract method 0x3883e119.
//
// Solidity: function calc_token_amount(uint256[3] amounts, bool deposit) view returns(uint256)
func (_ThreePool *ThreePoolCallerSession) CalcTokenAmount(amounts [3]*big.Int, deposit bool) (*big.Int, error) {
	return _ThreePool.Contract.CalcTokenAmount(&_ThreePool.CallOpts, amounts, deposit)
}

// CalcWithdrawOneCoin is a free data retrieval call binding the contract method 0xcc2b27d7.
//
// Solidity: function calc_withdraw_one_coin(uint256 _token_amount, int128 i) view returns(uint256)
func (_ThreePool *ThreePoolCaller) CalcWithdrawOneCoin(opts *bind.CallOpts, _token_amount *big.Int, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ThreePool.contract.Call(opts, &out, "calc_withdraw_one_coin", _token_amount, i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcWithdrawOneCoin is a free data retrieval call binding the contract method 0xcc2b27d7.
//
// Solidity: function calc_withdraw_one_coin(uint256 _token_amount, int128 i) view returns(uint256)
func (_ThreePool *ThreePoolSession) CalcWithdrawOneCoin(_token_amount *big.Int, i *big.Int) (*big.Int, error) {
	return _ThreePool.Contract.CalcWithdrawOneCoin(&_ThreePool.CallOpts, _token_amount, i)
}

// CalcWithdrawOneCoin is a free data retrieval call binding the contract method 0xcc2b27d7.
//
// Solidity: function calc_withdraw_one_coin(uint256 _token_amount, int128 i) view returns(uint256)
func (_ThreePool *ThreePoolCallerSession) CalcWithdrawOneCoin(_token_amount *big.Int, i *big.Int) (*big.Int, error) {
	return _ThreePool.Contract.CalcWithdrawOneCoin(&_ThreePool.CallOpts, _token_amount, i)
}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 arg0) view returns(address)
func (_ThreePool *ThreePoolCaller) Coins(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ThreePool.contract.Call(opts, &out, "coins", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 arg0) view returns(address)
func (_ThreePool *ThreePoolSession) Coins(arg0 *big.Int) (common.Address, error) {
	return _ThreePool.Contract.Coins(&_ThreePool.CallOpts, arg0)
}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 arg0) view returns(address)
func (_ThreePool *ThreePoolCallerSession) Coins(arg0 *big.Int) (common.Address, error) {
	return _ThreePool.Contract.Coins(&_ThreePool.CallOpts, arg0)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_ThreePool *ThreePoolCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ThreePool.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_ThreePool *ThreePoolSession) Fee() (*big.Int, error) {
	return _ThreePool.Contract.Fee(&_ThreePool.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_ThreePool *ThreePoolCallerSession) Fee() (*big.Int, error) {
	return _ThreePool.Contract.Fee(&_ThreePool.CallOpts)
}

// FutureA is a free data retrieval call binding the contract method 0xb4b577ad.
//
// Solidity: function future_A() view returns(uint256)
func (_ThreePool *ThreePoolCaller) FutureA(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ThreePool.contract.Call(opts, &out, "future_A")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureA is a free data retrieval call binding the contract method 0xb4b577ad.
//
// Solidity: function future_A() view returns(uint256)
func (_ThreePool *ThreePoolSession) FutureA() (*big.Int, error) {
	return _ThreePool.Contract.FutureA(&_ThreePool.CallOpts)
}

// FutureA is a free data retrieval call binding the contract method 0xb4b577ad.
//
// Solidity: function future_A() view returns(uint256)
func (_ThreePool *ThreePoolCallerSession) FutureA() (*big.Int, error) {
	return _ThreePool.Contract.FutureA(&_ThreePool.CallOpts)
}

// FutureATime is a free data retrieval call binding the contract method 0x14052288.
//
// Solidity: function future_A_time() view returns(uint256)
func (_ThreePool *ThreePoolCaller) FutureATime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ThreePool.contract.Call(opts, &out, "future_A_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureATime is a free data retrieval call binding the contract method 0x14052288.
//
// Solidity: function future_A_time() view returns(uint256)
func (_ThreePool *ThreePoolSession) FutureATime() (*big.Int, error) {
	return _ThreePool.Contract.FutureATime(&_ThreePool.CallOpts)
}

// FutureATime is a free data retrieval call binding the contract method 0x14052288.
//
// Solidity: function future_A_time() view returns(uint256)
func (_ThreePool *ThreePoolCallerSession) FutureATime() (*big.Int, error) {
	return _ThreePool.Contract.FutureATime(&_ThreePool.CallOpts)
}

// FutureAdminFee is a free data retrieval call binding the contract method 0xe3824462.
//
// Solidity: function future_admin_fee() view returns(uint256)
func (_ThreePool *ThreePoolCaller) FutureAdminFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ThreePool.contract.Call(opts, &out, "future_admin_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureAdminFee is a free data retrieval call binding the contract method 0xe3824462.
//
// Solidity: function future_admin_fee() view returns(uint256)
func (_ThreePool *ThreePoolSession) FutureAdminFee() (*big.Int, error) {
	return _ThreePool.Contract.FutureAdminFee(&_ThreePool.CallOpts)
}

// FutureAdminFee is a free data retrieval call binding the contract method 0xe3824462.
//
// Solidity: function future_admin_fee() view returns(uint256)
func (_ThreePool *ThreePoolCallerSession) FutureAdminFee() (*big.Int, error) {
	return _ThreePool.Contract.FutureAdminFee(&_ThreePool.CallOpts)
}

// FutureFee is a free data retrieval call binding the contract method 0x58680d0b.
//
// Solidity: function future_fee() view returns(uint256)
func (_ThreePool *ThreePoolCaller) FutureFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ThreePool.contract.Call(opts, &out, "future_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureFee is a free data retrieval call binding the contract method 0x58680d0b.
//
// Solidity: function future_fee() view returns(uint256)
func (_ThreePool *ThreePoolSession) FutureFee() (*big.Int, error) {
	return _ThreePool.Contract.FutureFee(&_ThreePool.CallOpts)
}

// FutureFee is a free data retrieval call binding the contract method 0x58680d0b.
//
// Solidity: function future_fee() view returns(uint256)
func (_ThreePool *ThreePoolCallerSession) FutureFee() (*big.Int, error) {
	return _ThreePool.Contract.FutureFee(&_ThreePool.CallOpts)
}

// FutureOwner is a free data retrieval call binding the contract method 0x1ec0cdc1.
//
// Solidity: function future_owner() view returns(address)
func (_ThreePool *ThreePoolCaller) FutureOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ThreePool.contract.Call(opts, &out, "future_owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FutureOwner is a free data retrieval call binding the contract method 0x1ec0cdc1.
//
// Solidity: function future_owner() view returns(address)
func (_ThreePool *ThreePoolSession) FutureOwner() (common.Address, error) {
	return _ThreePool.Contract.FutureOwner(&_ThreePool.CallOpts)
}

// FutureOwner is a free data retrieval call binding the contract method 0x1ec0cdc1.
//
// Solidity: function future_owner() view returns(address)
func (_ThreePool *ThreePoolCallerSession) FutureOwner() (common.Address, error) {
	return _ThreePool.Contract.FutureOwner(&_ThreePool.CallOpts)
}

// GetDy is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_ThreePool *ThreePoolCaller) GetDy(opts *bind.CallOpts, i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ThreePool.contract.Call(opts, &out, "get_dy", i, j, dx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDy is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_ThreePool *ThreePoolSession) GetDy(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _ThreePool.Contract.GetDy(&_ThreePool.CallOpts, i, j, dx)
}

// GetDy is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_ThreePool *ThreePoolCallerSession) GetDy(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _ThreePool.Contract.GetDy(&_ThreePool.CallOpts, i, j, dx)
}

// GetDyUnderlying is a free data retrieval call binding the contract method 0x07211ef7.
//
// Solidity: function get_dy_underlying(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_ThreePool *ThreePoolCaller) GetDyUnderlying(opts *bind.CallOpts, i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ThreePool.contract.Call(opts, &out, "get_dy_underlying", i, j, dx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDyUnderlying is a free data retrieval call binding the contract method 0x07211ef7.
//
// Solidity: function get_dy_underlying(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_ThreePool *ThreePoolSession) GetDyUnderlying(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _ThreePool.Contract.GetDyUnderlying(&_ThreePool.CallOpts, i, j, dx)
}

// GetDyUnderlying is a free data retrieval call binding the contract method 0x07211ef7.
//
// Solidity: function get_dy_underlying(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_ThreePool *ThreePoolCallerSession) GetDyUnderlying(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _ThreePool.Contract.GetDyUnderlying(&_ThreePool.CallOpts, i, j, dx)
}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_ThreePool *ThreePoolCaller) GetVirtualPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ThreePool.contract.Call(opts, &out, "get_virtual_price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_ThreePool *ThreePoolSession) GetVirtualPrice() (*big.Int, error) {
	return _ThreePool.Contract.GetVirtualPrice(&_ThreePool.CallOpts)
}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_ThreePool *ThreePoolCallerSession) GetVirtualPrice() (*big.Int, error) {
	return _ThreePool.Contract.GetVirtualPrice(&_ThreePool.CallOpts)
}

// InitialA is a free data retrieval call binding the contract method 0x5409491a.
//
// Solidity: function initial_A() view returns(uint256)
func (_ThreePool *ThreePoolCaller) InitialA(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ThreePool.contract.Call(opts, &out, "initial_A")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialA is a free data retrieval call binding the contract method 0x5409491a.
//
// Solidity: function initial_A() view returns(uint256)
func (_ThreePool *ThreePoolSession) InitialA() (*big.Int, error) {
	return _ThreePool.Contract.InitialA(&_ThreePool.CallOpts)
}

// InitialA is a free data retrieval call binding the contract method 0x5409491a.
//
// Solidity: function initial_A() view returns(uint256)
func (_ThreePool *ThreePoolCallerSession) InitialA() (*big.Int, error) {
	return _ThreePool.Contract.InitialA(&_ThreePool.CallOpts)
}

// InitialATime is a free data retrieval call binding the contract method 0x2081066c.
//
// Solidity: function initial_A_time() view returns(uint256)
func (_ThreePool *ThreePoolCaller) InitialATime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ThreePool.contract.Call(opts, &out, "initial_A_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialATime is a free data retrieval call binding the contract method 0x2081066c.
//
// Solidity: function initial_A_time() view returns(uint256)
func (_ThreePool *ThreePoolSession) InitialATime() (*big.Int, error) {
	return _ThreePool.Contract.InitialATime(&_ThreePool.CallOpts)
}

// InitialATime is a free data retrieval call binding the contract method 0x2081066c.
//
// Solidity: function initial_A_time() view returns(uint256)
func (_ThreePool *ThreePoolCallerSession) InitialATime() (*big.Int, error) {
	return _ThreePool.Contract.InitialATime(&_ThreePool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ThreePool *ThreePoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ThreePool.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ThreePool *ThreePoolSession) Owner() (common.Address, error) {
	return _ThreePool.Contract.Owner(&_ThreePool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ThreePool *ThreePoolCallerSession) Owner() (common.Address, error) {
	return _ThreePool.Contract.Owner(&_ThreePool.CallOpts)
}

// TransferOwnershipDeadline is a free data retrieval call binding the contract method 0xe0a0b586.
//
// Solidity: function transfer_ownership_deadline() view returns(uint256)
func (_ThreePool *ThreePoolCaller) TransferOwnershipDeadline(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ThreePool.contract.Call(opts, &out, "transfer_ownership_deadline")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TransferOwnershipDeadline is a free data retrieval call binding the contract method 0xe0a0b586.
//
// Solidity: function transfer_ownership_deadline() view returns(uint256)
func (_ThreePool *ThreePoolSession) TransferOwnershipDeadline() (*big.Int, error) {
	return _ThreePool.Contract.TransferOwnershipDeadline(&_ThreePool.CallOpts)
}

// TransferOwnershipDeadline is a free data retrieval call binding the contract method 0xe0a0b586.
//
// Solidity: function transfer_ownership_deadline() view returns(uint256)
func (_ThreePool *ThreePoolCallerSession) TransferOwnershipDeadline() (*big.Int, error) {
	return _ThreePool.Contract.TransferOwnershipDeadline(&_ThreePool.CallOpts)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x4515cef3.
//
// Solidity: function add_liquidity(uint256[3] amounts, uint256 min_mint_amount) returns()
func (_ThreePool *ThreePoolTransactor) AddLiquidity(opts *bind.TransactOpts, amounts [3]*big.Int, min_mint_amount *big.Int) (*types.Transaction, error) {
	return _ThreePool.contract.Transact(opts, "add_liquidity", amounts, min_mint_amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x4515cef3.
//
// Solidity: function add_liquidity(uint256[3] amounts, uint256 min_mint_amount) returns()
func (_ThreePool *ThreePoolSession) AddLiquidity(amounts [3]*big.Int, min_mint_amount *big.Int) (*types.Transaction, error) {
	return _ThreePool.Contract.AddLiquidity(&_ThreePool.TransactOpts, amounts, min_mint_amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x4515cef3.
//
// Solidity: function add_liquidity(uint256[3] amounts, uint256 min_mint_amount) returns()
func (_ThreePool *ThreePoolTransactorSession) AddLiquidity(amounts [3]*big.Int, min_mint_amount *big.Int) (*types.Transaction, error) {
	return _ThreePool.Contract.AddLiquidity(&_ThreePool.TransactOpts, amounts, min_mint_amount)
}

// ApplyNewFee is a paid mutator transaction binding the contract method 0x4f12fe97.
//
// Solidity: function apply_new_fee() returns()
func (_ThreePool *ThreePoolTransactor) ApplyNewFee(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ThreePool.contract.Transact(opts, "apply_new_fee")
}

// ApplyNewFee is a paid mutator transaction binding the contract method 0x4f12fe97.
//
// Solidity: function apply_new_fee() returns()
func (_ThreePool *ThreePoolSession) ApplyNewFee() (*types.Transaction, error) {
	return _ThreePool.Contract.ApplyNewFee(&_ThreePool.TransactOpts)
}

// ApplyNewFee is a paid mutator transaction binding the contract method 0x4f12fe97.
//
// Solidity: function apply_new_fee() returns()
func (_ThreePool *ThreePoolTransactorSession) ApplyNewFee() (*types.Transaction, error) {
	return _ThreePool.Contract.ApplyNewFee(&_ThreePool.TransactOpts)
}

// ApplyTransferOwnership is a paid mutator transaction binding the contract method 0x6a1c05ae.
//
// Solidity: function apply_transfer_ownership() returns()
func (_ThreePool *ThreePoolTransactor) ApplyTransferOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ThreePool.contract.Transact(opts, "apply_transfer_ownership")
}

// ApplyTransferOwnership is a paid mutator transaction binding the contract method 0x6a1c05ae.
//
// Solidity: function apply_transfer_ownership() returns()
func (_ThreePool *ThreePoolSession) ApplyTransferOwnership() (*types.Transaction, error) {
	return _ThreePool.Contract.ApplyTransferOwnership(&_ThreePool.TransactOpts)
}

// ApplyTransferOwnership is a paid mutator transaction binding the contract method 0x6a1c05ae.
//
// Solidity: function apply_transfer_ownership() returns()
func (_ThreePool *ThreePoolTransactorSession) ApplyTransferOwnership() (*types.Transaction, error) {
	return _ThreePool.Contract.ApplyTransferOwnership(&_ThreePool.TransactOpts)
}

// CommitNewFee is a paid mutator transaction binding the contract method 0x5b5a1467.
//
// Solidity: function commit_new_fee(uint256 new_fee, uint256 new_admin_fee) returns()
func (_ThreePool *ThreePoolTransactor) CommitNewFee(opts *bind.TransactOpts, new_fee *big.Int, new_admin_fee *big.Int) (*types.Transaction, error) {
	return _ThreePool.contract.Transact(opts, "commit_new_fee", new_fee, new_admin_fee)
}

// CommitNewFee is a paid mutator transaction binding the contract method 0x5b5a1467.
//
// Solidity: function commit_new_fee(uint256 new_fee, uint256 new_admin_fee) returns()
func (_ThreePool *ThreePoolSession) CommitNewFee(new_fee *big.Int, new_admin_fee *big.Int) (*types.Transaction, error) {
	return _ThreePool.Contract.CommitNewFee(&_ThreePool.TransactOpts, new_fee, new_admin_fee)
}

// CommitNewFee is a paid mutator transaction binding the contract method 0x5b5a1467.
//
// Solidity: function commit_new_fee(uint256 new_fee, uint256 new_admin_fee) returns()
func (_ThreePool *ThreePoolTransactorSession) CommitNewFee(new_fee *big.Int, new_admin_fee *big.Int) (*types.Transaction, error) {
	return _ThreePool.Contract.CommitNewFee(&_ThreePool.TransactOpts, new_fee, new_admin_fee)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address _owner) returns()
func (_ThreePool *ThreePoolTransactor) CommitTransferOwnership(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _ThreePool.contract.Transact(opts, "commit_transfer_ownership", _owner)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address _owner) returns()
func (_ThreePool *ThreePoolSession) CommitTransferOwnership(_owner common.Address) (*types.Transaction, error) {
	return _ThreePool.Contract.CommitTransferOwnership(&_ThreePool.TransactOpts, _owner)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address _owner) returns()
func (_ThreePool *ThreePoolTransactorSession) CommitTransferOwnership(_owner common.Address) (*types.Transaction, error) {
	return _ThreePool.Contract.CommitTransferOwnership(&_ThreePool.TransactOpts, _owner)
}

// DonateAdminFees is a paid mutator transaction binding the contract method 0x524c3901.
//
// Solidity: function donate_admin_fees() returns()
func (_ThreePool *ThreePoolTransactor) DonateAdminFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ThreePool.contract.Transact(opts, "donate_admin_fees")
}

// DonateAdminFees is a paid mutator transaction binding the contract method 0x524c3901.
//
// Solidity: function donate_admin_fees() returns()
func (_ThreePool *ThreePoolSession) DonateAdminFees() (*types.Transaction, error) {
	return _ThreePool.Contract.DonateAdminFees(&_ThreePool.TransactOpts)
}

// DonateAdminFees is a paid mutator transaction binding the contract method 0x524c3901.
//
// Solidity: function donate_admin_fees() returns()
func (_ThreePool *ThreePoolTransactorSession) DonateAdminFees() (*types.Transaction, error) {
	return _ThreePool.Contract.DonateAdminFees(&_ThreePool.TransactOpts)
}

// Exchange is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 dx, uint256 min_dy) returns()
func (_ThreePool *ThreePoolTransactor) Exchange(opts *bind.TransactOpts, i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _ThreePool.contract.Transact(opts, "exchange", i, j, dx, min_dy)
}

// Exchange is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 dx, uint256 min_dy) returns()
func (_ThreePool *ThreePoolSession) Exchange(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _ThreePool.Contract.Exchange(&_ThreePool.TransactOpts, i, j, dx, min_dy)
}

// Exchange is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 dx, uint256 min_dy) returns()
func (_ThreePool *ThreePoolTransactorSession) Exchange(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _ThreePool.Contract.Exchange(&_ThreePool.TransactOpts, i, j, dx, min_dy)
}

// KillMe is a paid mutator transaction binding the contract method 0xe3698853.
//
// Solidity: function kill_me() returns()
func (_ThreePool *ThreePoolTransactor) KillMe(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ThreePool.contract.Transact(opts, "kill_me")
}

// KillMe is a paid mutator transaction binding the contract method 0xe3698853.
//
// Solidity: function kill_me() returns()
func (_ThreePool *ThreePoolSession) KillMe() (*types.Transaction, error) {
	return _ThreePool.Contract.KillMe(&_ThreePool.TransactOpts)
}

// KillMe is a paid mutator transaction binding the contract method 0xe3698853.
//
// Solidity: function kill_me() returns()
func (_ThreePool *ThreePoolTransactorSession) KillMe() (*types.Transaction, error) {
	return _ThreePool.Contract.KillMe(&_ThreePool.TransactOpts)
}

// RampA is a paid mutator transaction binding the contract method 0x3c157e64.
//
// Solidity: function ramp_A(uint256 _future_A, uint256 _future_time) returns()
func (_ThreePool *ThreePoolTransactor) RampA(opts *bind.TransactOpts, _future_A *big.Int, _future_time *big.Int) (*types.Transaction, error) {
	return _ThreePool.contract.Transact(opts, "ramp_A", _future_A, _future_time)
}

// RampA is a paid mutator transaction binding the contract method 0x3c157e64.
//
// Solidity: function ramp_A(uint256 _future_A, uint256 _future_time) returns()
func (_ThreePool *ThreePoolSession) RampA(_future_A *big.Int, _future_time *big.Int) (*types.Transaction, error) {
	return _ThreePool.Contract.RampA(&_ThreePool.TransactOpts, _future_A, _future_time)
}

// RampA is a paid mutator transaction binding the contract method 0x3c157e64.
//
// Solidity: function ramp_A(uint256 _future_A, uint256 _future_time) returns()
func (_ThreePool *ThreePoolTransactorSession) RampA(_future_A *big.Int, _future_time *big.Int) (*types.Transaction, error) {
	return _ThreePool.Contract.RampA(&_ThreePool.TransactOpts, _future_A, _future_time)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xecb586a5.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[3] min_amounts) returns()
func (_ThreePool *ThreePoolTransactor) RemoveLiquidity(opts *bind.TransactOpts, _amount *big.Int, min_amounts [3]*big.Int) (*types.Transaction, error) {
	return _ThreePool.contract.Transact(opts, "remove_liquidity", _amount, min_amounts)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xecb586a5.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[3] min_amounts) returns()
func (_ThreePool *ThreePoolSession) RemoveLiquidity(_amount *big.Int, min_amounts [3]*big.Int) (*types.Transaction, error) {
	return _ThreePool.Contract.RemoveLiquidity(&_ThreePool.TransactOpts, _amount, min_amounts)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xecb586a5.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[3] min_amounts) returns()
func (_ThreePool *ThreePoolTransactorSession) RemoveLiquidity(_amount *big.Int, min_amounts [3]*big.Int) (*types.Transaction, error) {
	return _ThreePool.Contract.RemoveLiquidity(&_ThreePool.TransactOpts, _amount, min_amounts)
}

// RemoveLiquidityImbalance is a paid mutator transaction binding the contract method 0x9fdaea0c.
//
// Solidity: function remove_liquidity_imbalance(uint256[3] amounts, uint256 max_burn_amount) returns()
func (_ThreePool *ThreePoolTransactor) RemoveLiquidityImbalance(opts *bind.TransactOpts, amounts [3]*big.Int, max_burn_amount *big.Int) (*types.Transaction, error) {
	return _ThreePool.contract.Transact(opts, "remove_liquidity_imbalance", amounts, max_burn_amount)
}

// RemoveLiquidityImbalance is a paid mutator transaction binding the contract method 0x9fdaea0c.
//
// Solidity: function remove_liquidity_imbalance(uint256[3] amounts, uint256 max_burn_amount) returns()
func (_ThreePool *ThreePoolSession) RemoveLiquidityImbalance(amounts [3]*big.Int, max_burn_amount *big.Int) (*types.Transaction, error) {
	return _ThreePool.Contract.RemoveLiquidityImbalance(&_ThreePool.TransactOpts, amounts, max_burn_amount)
}

// RemoveLiquidityImbalance is a paid mutator transaction binding the contract method 0x9fdaea0c.
//
// Solidity: function remove_liquidity_imbalance(uint256[3] amounts, uint256 max_burn_amount) returns()
func (_ThreePool *ThreePoolTransactorSession) RemoveLiquidityImbalance(amounts [3]*big.Int, max_burn_amount *big.Int) (*types.Transaction, error) {
	return _ThreePool.Contract.RemoveLiquidityImbalance(&_ThreePool.TransactOpts, amounts, max_burn_amount)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0x1a4d01d2.
//
// Solidity: function remove_liquidity_one_coin(uint256 _token_amount, int128 i, uint256 min_amount) returns()
func (_ThreePool *ThreePoolTransactor) RemoveLiquidityOneCoin(opts *bind.TransactOpts, _token_amount *big.Int, i *big.Int, min_amount *big.Int) (*types.Transaction, error) {
	return _ThreePool.contract.Transact(opts, "remove_liquidity_one_coin", _token_amount, i, min_amount)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0x1a4d01d2.
//
// Solidity: function remove_liquidity_one_coin(uint256 _token_amount, int128 i, uint256 min_amount) returns()
func (_ThreePool *ThreePoolSession) RemoveLiquidityOneCoin(_token_amount *big.Int, i *big.Int, min_amount *big.Int) (*types.Transaction, error) {
	return _ThreePool.Contract.RemoveLiquidityOneCoin(&_ThreePool.TransactOpts, _token_amount, i, min_amount)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0x1a4d01d2.
//
// Solidity: function remove_liquidity_one_coin(uint256 _token_amount, int128 i, uint256 min_amount) returns()
func (_ThreePool *ThreePoolTransactorSession) RemoveLiquidityOneCoin(_token_amount *big.Int, i *big.Int, min_amount *big.Int) (*types.Transaction, error) {
	return _ThreePool.Contract.RemoveLiquidityOneCoin(&_ThreePool.TransactOpts, _token_amount, i, min_amount)
}

// RevertNewParameters is a paid mutator transaction binding the contract method 0x226840fb.
//
// Solidity: function revert_new_parameters() returns()
func (_ThreePool *ThreePoolTransactor) RevertNewParameters(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ThreePool.contract.Transact(opts, "revert_new_parameters")
}

// RevertNewParameters is a paid mutator transaction binding the contract method 0x226840fb.
//
// Solidity: function revert_new_parameters() returns()
func (_ThreePool *ThreePoolSession) RevertNewParameters() (*types.Transaction, error) {
	return _ThreePool.Contract.RevertNewParameters(&_ThreePool.TransactOpts)
}

// RevertNewParameters is a paid mutator transaction binding the contract method 0x226840fb.
//
// Solidity: function revert_new_parameters() returns()
func (_ThreePool *ThreePoolTransactorSession) RevertNewParameters() (*types.Transaction, error) {
	return _ThreePool.Contract.RevertNewParameters(&_ThreePool.TransactOpts)
}

// RevertTransferOwnership is a paid mutator transaction binding the contract method 0x86fbf193.
//
// Solidity: function revert_transfer_ownership() returns()
func (_ThreePool *ThreePoolTransactor) RevertTransferOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ThreePool.contract.Transact(opts, "revert_transfer_ownership")
}

// RevertTransferOwnership is a paid mutator transaction binding the contract method 0x86fbf193.
//
// Solidity: function revert_transfer_ownership() returns()
func (_ThreePool *ThreePoolSession) RevertTransferOwnership() (*types.Transaction, error) {
	return _ThreePool.Contract.RevertTransferOwnership(&_ThreePool.TransactOpts)
}

// RevertTransferOwnership is a paid mutator transaction binding the contract method 0x86fbf193.
//
// Solidity: function revert_transfer_ownership() returns()
func (_ThreePool *ThreePoolTransactorSession) RevertTransferOwnership() (*types.Transaction, error) {
	return _ThreePool.Contract.RevertTransferOwnership(&_ThreePool.TransactOpts)
}

// StopRampA is a paid mutator transaction binding the contract method 0x551a6588.
//
// Solidity: function stop_ramp_A() returns()
func (_ThreePool *ThreePoolTransactor) StopRampA(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ThreePool.contract.Transact(opts, "stop_ramp_A")
}

// StopRampA is a paid mutator transaction binding the contract method 0x551a6588.
//
// Solidity: function stop_ramp_A() returns()
func (_ThreePool *ThreePoolSession) StopRampA() (*types.Transaction, error) {
	return _ThreePool.Contract.StopRampA(&_ThreePool.TransactOpts)
}

// StopRampA is a paid mutator transaction binding the contract method 0x551a6588.
//
// Solidity: function stop_ramp_A() returns()
func (_ThreePool *ThreePoolTransactorSession) StopRampA() (*types.Transaction, error) {
	return _ThreePool.Contract.StopRampA(&_ThreePool.TransactOpts)
}

// UnkillMe is a paid mutator transaction binding the contract method 0x3046f972.
//
// Solidity: function unkill_me() returns()
func (_ThreePool *ThreePoolTransactor) UnkillMe(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ThreePool.contract.Transact(opts, "unkill_me")
}

// UnkillMe is a paid mutator transaction binding the contract method 0x3046f972.
//
// Solidity: function unkill_me() returns()
func (_ThreePool *ThreePoolSession) UnkillMe() (*types.Transaction, error) {
	return _ThreePool.Contract.UnkillMe(&_ThreePool.TransactOpts)
}

// UnkillMe is a paid mutator transaction binding the contract method 0x3046f972.
//
// Solidity: function unkill_me() returns()
func (_ThreePool *ThreePoolTransactorSession) UnkillMe() (*types.Transaction, error) {
	return _ThreePool.Contract.UnkillMe(&_ThreePool.TransactOpts)
}

// WithdrawAdminFees is a paid mutator transaction binding the contract method 0x30c54085.
//
// Solidity: function withdraw_admin_fees() returns()
func (_ThreePool *ThreePoolTransactor) WithdrawAdminFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ThreePool.contract.Transact(opts, "withdraw_admin_fees")
}

// WithdrawAdminFees is a paid mutator transaction binding the contract method 0x30c54085.
//
// Solidity: function withdraw_admin_fees() returns()
func (_ThreePool *ThreePoolSession) WithdrawAdminFees() (*types.Transaction, error) {
	return _ThreePool.Contract.WithdrawAdminFees(&_ThreePool.TransactOpts)
}

// WithdrawAdminFees is a paid mutator transaction binding the contract method 0x30c54085.
//
// Solidity: function withdraw_admin_fees() returns()
func (_ThreePool *ThreePoolTransactorSession) WithdrawAdminFees() (*types.Transaction, error) {
	return _ThreePool.Contract.WithdrawAdminFees(&_ThreePool.TransactOpts)
}

// ThreePoolAddLiquidityIterator is returned from FilterAddLiquidity and is used to iterate over the raw logs and unpacked data for AddLiquidity events raised by the ThreePool contract.
type ThreePoolAddLiquidityIterator struct {
	Event *ThreePoolAddLiquidity // Event containing the contract specifics and raw log

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
func (it *ThreePoolAddLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ThreePoolAddLiquidity)
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
		it.Event = new(ThreePoolAddLiquidity)
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
func (it *ThreePoolAddLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ThreePoolAddLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ThreePoolAddLiquidity represents a AddLiquidity event raised by the ThreePool contract.
type ThreePoolAddLiquidity struct {
	Provider     common.Address
	TokenAmounts [3]*big.Int
	Fees         [3]*big.Int
	Invariant    *big.Int
	TokenSupply  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAddLiquidity is a free log retrieval operation binding the contract event 0x423f6495a08fc652425cf4ed0d1f9e37e571d9b9529b1c1c23cce780b2e7df0d.
//
// Solidity: event AddLiquidity(address indexed provider, uint256[3] token_amounts, uint256[3] fees, uint256 invariant, uint256 token_supply)
func (_ThreePool *ThreePoolFilterer) FilterAddLiquidity(opts *bind.FilterOpts, provider []common.Address) (*ThreePoolAddLiquidityIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _ThreePool.contract.FilterLogs(opts, "AddLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return &ThreePoolAddLiquidityIterator{contract: _ThreePool.contract, event: "AddLiquidity", logs: logs, sub: sub}, nil
}

// WatchAddLiquidity is a free log subscription operation binding the contract event 0x423f6495a08fc652425cf4ed0d1f9e37e571d9b9529b1c1c23cce780b2e7df0d.
//
// Solidity: event AddLiquidity(address indexed provider, uint256[3] token_amounts, uint256[3] fees, uint256 invariant, uint256 token_supply)
func (_ThreePool *ThreePoolFilterer) WatchAddLiquidity(opts *bind.WatchOpts, sink chan<- *ThreePoolAddLiquidity, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _ThreePool.contract.WatchLogs(opts, "AddLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ThreePoolAddLiquidity)
				if err := _ThreePool.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
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

// ParseAddLiquidity is a log parse operation binding the contract event 0x423f6495a08fc652425cf4ed0d1f9e37e571d9b9529b1c1c23cce780b2e7df0d.
//
// Solidity: event AddLiquidity(address indexed provider, uint256[3] token_amounts, uint256[3] fees, uint256 invariant, uint256 token_supply)
func (_ThreePool *ThreePoolFilterer) ParseAddLiquidity(log types.Log) (*ThreePoolAddLiquidity, error) {
	event := new(ThreePoolAddLiquidity)
	if err := _ThreePool.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ThreePoolCommitNewAdminIterator is returned from FilterCommitNewAdmin and is used to iterate over the raw logs and unpacked data for CommitNewAdmin events raised by the ThreePool contract.
type ThreePoolCommitNewAdminIterator struct {
	Event *ThreePoolCommitNewAdmin // Event containing the contract specifics and raw log

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
func (it *ThreePoolCommitNewAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ThreePoolCommitNewAdmin)
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
		it.Event = new(ThreePoolCommitNewAdmin)
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
func (it *ThreePoolCommitNewAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ThreePoolCommitNewAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ThreePoolCommitNewAdmin represents a CommitNewAdmin event raised by the ThreePool contract.
type ThreePoolCommitNewAdmin struct {
	Deadline *big.Int
	Admin    common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCommitNewAdmin is a free log retrieval operation binding the contract event 0x181aa3aa17d4cbf99265dd4443eba009433d3cde79d60164fde1d1a192beb935.
//
// Solidity: event CommitNewAdmin(uint256 indexed deadline, address indexed admin)
func (_ThreePool *ThreePoolFilterer) FilterCommitNewAdmin(opts *bind.FilterOpts, deadline []*big.Int, admin []common.Address) (*ThreePoolCommitNewAdminIterator, error) {

	var deadlineRule []interface{}
	for _, deadlineItem := range deadline {
		deadlineRule = append(deadlineRule, deadlineItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _ThreePool.contract.FilterLogs(opts, "CommitNewAdmin", deadlineRule, adminRule)
	if err != nil {
		return nil, err
	}
	return &ThreePoolCommitNewAdminIterator{contract: _ThreePool.contract, event: "CommitNewAdmin", logs: logs, sub: sub}, nil
}

// WatchCommitNewAdmin is a free log subscription operation binding the contract event 0x181aa3aa17d4cbf99265dd4443eba009433d3cde79d60164fde1d1a192beb935.
//
// Solidity: event CommitNewAdmin(uint256 indexed deadline, address indexed admin)
func (_ThreePool *ThreePoolFilterer) WatchCommitNewAdmin(opts *bind.WatchOpts, sink chan<- *ThreePoolCommitNewAdmin, deadline []*big.Int, admin []common.Address) (event.Subscription, error) {

	var deadlineRule []interface{}
	for _, deadlineItem := range deadline {
		deadlineRule = append(deadlineRule, deadlineItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _ThreePool.contract.WatchLogs(opts, "CommitNewAdmin", deadlineRule, adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ThreePoolCommitNewAdmin)
				if err := _ThreePool.contract.UnpackLog(event, "CommitNewAdmin", log); err != nil {
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

// ParseCommitNewAdmin is a log parse operation binding the contract event 0x181aa3aa17d4cbf99265dd4443eba009433d3cde79d60164fde1d1a192beb935.
//
// Solidity: event CommitNewAdmin(uint256 indexed deadline, address indexed admin)
func (_ThreePool *ThreePoolFilterer) ParseCommitNewAdmin(log types.Log) (*ThreePoolCommitNewAdmin, error) {
	event := new(ThreePoolCommitNewAdmin)
	if err := _ThreePool.contract.UnpackLog(event, "CommitNewAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ThreePoolCommitNewFeeIterator is returned from FilterCommitNewFee and is used to iterate over the raw logs and unpacked data for CommitNewFee events raised by the ThreePool contract.
type ThreePoolCommitNewFeeIterator struct {
	Event *ThreePoolCommitNewFee // Event containing the contract specifics and raw log

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
func (it *ThreePoolCommitNewFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ThreePoolCommitNewFee)
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
		it.Event = new(ThreePoolCommitNewFee)
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
func (it *ThreePoolCommitNewFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ThreePoolCommitNewFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ThreePoolCommitNewFee represents a CommitNewFee event raised by the ThreePool contract.
type ThreePoolCommitNewFee struct {
	Deadline *big.Int
	Fee      *big.Int
	AdminFee *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCommitNewFee is a free log retrieval operation binding the contract event 0x351fc5da2fbf480f2225debf3664a4bc90fa9923743aad58b4603f648e931fe0.
//
// Solidity: event CommitNewFee(uint256 indexed deadline, uint256 fee, uint256 admin_fee)
func (_ThreePool *ThreePoolFilterer) FilterCommitNewFee(opts *bind.FilterOpts, deadline []*big.Int) (*ThreePoolCommitNewFeeIterator, error) {

	var deadlineRule []interface{}
	for _, deadlineItem := range deadline {
		deadlineRule = append(deadlineRule, deadlineItem)
	}

	logs, sub, err := _ThreePool.contract.FilterLogs(opts, "CommitNewFee", deadlineRule)
	if err != nil {
		return nil, err
	}
	return &ThreePoolCommitNewFeeIterator{contract: _ThreePool.contract, event: "CommitNewFee", logs: logs, sub: sub}, nil
}

// WatchCommitNewFee is a free log subscription operation binding the contract event 0x351fc5da2fbf480f2225debf3664a4bc90fa9923743aad58b4603f648e931fe0.
//
// Solidity: event CommitNewFee(uint256 indexed deadline, uint256 fee, uint256 admin_fee)
func (_ThreePool *ThreePoolFilterer) WatchCommitNewFee(opts *bind.WatchOpts, sink chan<- *ThreePoolCommitNewFee, deadline []*big.Int) (event.Subscription, error) {

	var deadlineRule []interface{}
	for _, deadlineItem := range deadline {
		deadlineRule = append(deadlineRule, deadlineItem)
	}

	logs, sub, err := _ThreePool.contract.WatchLogs(opts, "CommitNewFee", deadlineRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ThreePoolCommitNewFee)
				if err := _ThreePool.contract.UnpackLog(event, "CommitNewFee", log); err != nil {
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

// ParseCommitNewFee is a log parse operation binding the contract event 0x351fc5da2fbf480f2225debf3664a4bc90fa9923743aad58b4603f648e931fe0.
//
// Solidity: event CommitNewFee(uint256 indexed deadline, uint256 fee, uint256 admin_fee)
func (_ThreePool *ThreePoolFilterer) ParseCommitNewFee(log types.Log) (*ThreePoolCommitNewFee, error) {
	event := new(ThreePoolCommitNewFee)
	if err := _ThreePool.contract.UnpackLog(event, "CommitNewFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ThreePoolNewAdminIterator is returned from FilterNewAdmin and is used to iterate over the raw logs and unpacked data for NewAdmin events raised by the ThreePool contract.
type ThreePoolNewAdminIterator struct {
	Event *ThreePoolNewAdmin // Event containing the contract specifics and raw log

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
func (it *ThreePoolNewAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ThreePoolNewAdmin)
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
		it.Event = new(ThreePoolNewAdmin)
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
func (it *ThreePoolNewAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ThreePoolNewAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ThreePoolNewAdmin represents a NewAdmin event raised by the ThreePool contract.
type ThreePoolNewAdmin struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNewAdmin is a free log retrieval operation binding the contract event 0x71614071b88dee5e0b2ae578a9dd7b2ebbe9ae832ba419dc0242cd065a290b6c.
//
// Solidity: event NewAdmin(address indexed admin)
func (_ThreePool *ThreePoolFilterer) FilterNewAdmin(opts *bind.FilterOpts, admin []common.Address) (*ThreePoolNewAdminIterator, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _ThreePool.contract.FilterLogs(opts, "NewAdmin", adminRule)
	if err != nil {
		return nil, err
	}
	return &ThreePoolNewAdminIterator{contract: _ThreePool.contract, event: "NewAdmin", logs: logs, sub: sub}, nil
}

// WatchNewAdmin is a free log subscription operation binding the contract event 0x71614071b88dee5e0b2ae578a9dd7b2ebbe9ae832ba419dc0242cd065a290b6c.
//
// Solidity: event NewAdmin(address indexed admin)
func (_ThreePool *ThreePoolFilterer) WatchNewAdmin(opts *bind.WatchOpts, sink chan<- *ThreePoolNewAdmin, admin []common.Address) (event.Subscription, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _ThreePool.contract.WatchLogs(opts, "NewAdmin", adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ThreePoolNewAdmin)
				if err := _ThreePool.contract.UnpackLog(event, "NewAdmin", log); err != nil {
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

// ParseNewAdmin is a log parse operation binding the contract event 0x71614071b88dee5e0b2ae578a9dd7b2ebbe9ae832ba419dc0242cd065a290b6c.
//
// Solidity: event NewAdmin(address indexed admin)
func (_ThreePool *ThreePoolFilterer) ParseNewAdmin(log types.Log) (*ThreePoolNewAdmin, error) {
	event := new(ThreePoolNewAdmin)
	if err := _ThreePool.contract.UnpackLog(event, "NewAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ThreePoolNewFeeIterator is returned from FilterNewFee and is used to iterate over the raw logs and unpacked data for NewFee events raised by the ThreePool contract.
type ThreePoolNewFeeIterator struct {
	Event *ThreePoolNewFee // Event containing the contract specifics and raw log

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
func (it *ThreePoolNewFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ThreePoolNewFee)
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
		it.Event = new(ThreePoolNewFee)
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
func (it *ThreePoolNewFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ThreePoolNewFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ThreePoolNewFee represents a NewFee event raised by the ThreePool contract.
type ThreePoolNewFee struct {
	Fee      *big.Int
	AdminFee *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewFee is a free log retrieval operation binding the contract event 0xbe12859b636aed607d5230b2cc2711f68d70e51060e6cca1f575ef5d2fcc95d1.
//
// Solidity: event NewFee(uint256 fee, uint256 admin_fee)
func (_ThreePool *ThreePoolFilterer) FilterNewFee(opts *bind.FilterOpts) (*ThreePoolNewFeeIterator, error) {

	logs, sub, err := _ThreePool.contract.FilterLogs(opts, "NewFee")
	if err != nil {
		return nil, err
	}
	return &ThreePoolNewFeeIterator{contract: _ThreePool.contract, event: "NewFee", logs: logs, sub: sub}, nil
}

// WatchNewFee is a free log subscription operation binding the contract event 0xbe12859b636aed607d5230b2cc2711f68d70e51060e6cca1f575ef5d2fcc95d1.
//
// Solidity: event NewFee(uint256 fee, uint256 admin_fee)
func (_ThreePool *ThreePoolFilterer) WatchNewFee(opts *bind.WatchOpts, sink chan<- *ThreePoolNewFee) (event.Subscription, error) {

	logs, sub, err := _ThreePool.contract.WatchLogs(opts, "NewFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ThreePoolNewFee)
				if err := _ThreePool.contract.UnpackLog(event, "NewFee", log); err != nil {
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

// ParseNewFee is a log parse operation binding the contract event 0xbe12859b636aed607d5230b2cc2711f68d70e51060e6cca1f575ef5d2fcc95d1.
//
// Solidity: event NewFee(uint256 fee, uint256 admin_fee)
func (_ThreePool *ThreePoolFilterer) ParseNewFee(log types.Log) (*ThreePoolNewFee, error) {
	event := new(ThreePoolNewFee)
	if err := _ThreePool.contract.UnpackLog(event, "NewFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ThreePoolRampAIterator is returned from FilterRampA and is used to iterate over the raw logs and unpacked data for RampA events raised by the ThreePool contract.
type ThreePoolRampAIterator struct {
	Event *ThreePoolRampA // Event containing the contract specifics and raw log

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
func (it *ThreePoolRampAIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ThreePoolRampA)
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
		it.Event = new(ThreePoolRampA)
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
func (it *ThreePoolRampAIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ThreePoolRampAIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ThreePoolRampA represents a RampA event raised by the ThreePool contract.
type ThreePoolRampA struct {
	OldA        *big.Int
	NewA        *big.Int
	InitialTime *big.Int
	FutureTime  *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRampA is a free log retrieval operation binding the contract event 0xa2b71ec6df949300b59aab36b55e189697b750119dd349fcfa8c0f779e83c254.
//
// Solidity: event RampA(uint256 old_A, uint256 new_A, uint256 initial_time, uint256 future_time)
func (_ThreePool *ThreePoolFilterer) FilterRampA(opts *bind.FilterOpts) (*ThreePoolRampAIterator, error) {

	logs, sub, err := _ThreePool.contract.FilterLogs(opts, "RampA")
	if err != nil {
		return nil, err
	}
	return &ThreePoolRampAIterator{contract: _ThreePool.contract, event: "RampA", logs: logs, sub: sub}, nil
}

// WatchRampA is a free log subscription operation binding the contract event 0xa2b71ec6df949300b59aab36b55e189697b750119dd349fcfa8c0f779e83c254.
//
// Solidity: event RampA(uint256 old_A, uint256 new_A, uint256 initial_time, uint256 future_time)
func (_ThreePool *ThreePoolFilterer) WatchRampA(opts *bind.WatchOpts, sink chan<- *ThreePoolRampA) (event.Subscription, error) {

	logs, sub, err := _ThreePool.contract.WatchLogs(opts, "RampA")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ThreePoolRampA)
				if err := _ThreePool.contract.UnpackLog(event, "RampA", log); err != nil {
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

// ParseRampA is a log parse operation binding the contract event 0xa2b71ec6df949300b59aab36b55e189697b750119dd349fcfa8c0f779e83c254.
//
// Solidity: event RampA(uint256 old_A, uint256 new_A, uint256 initial_time, uint256 future_time)
func (_ThreePool *ThreePoolFilterer) ParseRampA(log types.Log) (*ThreePoolRampA, error) {
	event := new(ThreePoolRampA)
	if err := _ThreePool.contract.UnpackLog(event, "RampA", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ThreePoolRemoveLiquidityIterator is returned from FilterRemoveLiquidity and is used to iterate over the raw logs and unpacked data for RemoveLiquidity events raised by the ThreePool contract.
type ThreePoolRemoveLiquidityIterator struct {
	Event *ThreePoolRemoveLiquidity // Event containing the contract specifics and raw log

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
func (it *ThreePoolRemoveLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ThreePoolRemoveLiquidity)
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
		it.Event = new(ThreePoolRemoveLiquidity)
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
func (it *ThreePoolRemoveLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ThreePoolRemoveLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ThreePoolRemoveLiquidity represents a RemoveLiquidity event raised by the ThreePool contract.
type ThreePoolRemoveLiquidity struct {
	Provider     common.Address
	TokenAmounts [3]*big.Int
	Fees         [3]*big.Int
	TokenSupply  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidity is a free log retrieval operation binding the contract event 0xa49d4cf02656aebf8c771f5a8585638a2a15ee6c97cf7205d4208ed7c1df252d.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256[3] token_amounts, uint256[3] fees, uint256 token_supply)
func (_ThreePool *ThreePoolFilterer) FilterRemoveLiquidity(opts *bind.FilterOpts, provider []common.Address) (*ThreePoolRemoveLiquidityIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _ThreePool.contract.FilterLogs(opts, "RemoveLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return &ThreePoolRemoveLiquidityIterator{contract: _ThreePool.contract, event: "RemoveLiquidity", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidity is a free log subscription operation binding the contract event 0xa49d4cf02656aebf8c771f5a8585638a2a15ee6c97cf7205d4208ed7c1df252d.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256[3] token_amounts, uint256[3] fees, uint256 token_supply)
func (_ThreePool *ThreePoolFilterer) WatchRemoveLiquidity(opts *bind.WatchOpts, sink chan<- *ThreePoolRemoveLiquidity, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _ThreePool.contract.WatchLogs(opts, "RemoveLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ThreePoolRemoveLiquidity)
				if err := _ThreePool.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
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

// ParseRemoveLiquidity is a log parse operation binding the contract event 0xa49d4cf02656aebf8c771f5a8585638a2a15ee6c97cf7205d4208ed7c1df252d.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256[3] token_amounts, uint256[3] fees, uint256 token_supply)
func (_ThreePool *ThreePoolFilterer) ParseRemoveLiquidity(log types.Log) (*ThreePoolRemoveLiquidity, error) {
	event := new(ThreePoolRemoveLiquidity)
	if err := _ThreePool.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ThreePoolRemoveLiquidityImbalanceIterator is returned from FilterRemoveLiquidityImbalance and is used to iterate over the raw logs and unpacked data for RemoveLiquidityImbalance events raised by the ThreePool contract.
type ThreePoolRemoveLiquidityImbalanceIterator struct {
	Event *ThreePoolRemoveLiquidityImbalance // Event containing the contract specifics and raw log

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
func (it *ThreePoolRemoveLiquidityImbalanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ThreePoolRemoveLiquidityImbalance)
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
		it.Event = new(ThreePoolRemoveLiquidityImbalance)
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
func (it *ThreePoolRemoveLiquidityImbalanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ThreePoolRemoveLiquidityImbalanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ThreePoolRemoveLiquidityImbalance represents a RemoveLiquidityImbalance event raised by the ThreePool contract.
type ThreePoolRemoveLiquidityImbalance struct {
	Provider     common.Address
	TokenAmounts [3]*big.Int
	Fees         [3]*big.Int
	Invariant    *big.Int
	TokenSupply  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidityImbalance is a free log retrieval operation binding the contract event 0x173599dbf9c6ca6f7c3b590df07ae98a45d74ff54065505141e7de6c46a624c2.
//
// Solidity: event RemoveLiquidityImbalance(address indexed provider, uint256[3] token_amounts, uint256[3] fees, uint256 invariant, uint256 token_supply)
func (_ThreePool *ThreePoolFilterer) FilterRemoveLiquidityImbalance(opts *bind.FilterOpts, provider []common.Address) (*ThreePoolRemoveLiquidityImbalanceIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _ThreePool.contract.FilterLogs(opts, "RemoveLiquidityImbalance", providerRule)
	if err != nil {
		return nil, err
	}
	return &ThreePoolRemoveLiquidityImbalanceIterator{contract: _ThreePool.contract, event: "RemoveLiquidityImbalance", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidityImbalance is a free log subscription operation binding the contract event 0x173599dbf9c6ca6f7c3b590df07ae98a45d74ff54065505141e7de6c46a624c2.
//
// Solidity: event RemoveLiquidityImbalance(address indexed provider, uint256[3] token_amounts, uint256[3] fees, uint256 invariant, uint256 token_supply)
func (_ThreePool *ThreePoolFilterer) WatchRemoveLiquidityImbalance(opts *bind.WatchOpts, sink chan<- *ThreePoolRemoveLiquidityImbalance, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _ThreePool.contract.WatchLogs(opts, "RemoveLiquidityImbalance", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ThreePoolRemoveLiquidityImbalance)
				if err := _ThreePool.contract.UnpackLog(event, "RemoveLiquidityImbalance", log); err != nil {
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

// ParseRemoveLiquidityImbalance is a log parse operation binding the contract event 0x173599dbf9c6ca6f7c3b590df07ae98a45d74ff54065505141e7de6c46a624c2.
//
// Solidity: event RemoveLiquidityImbalance(address indexed provider, uint256[3] token_amounts, uint256[3] fees, uint256 invariant, uint256 token_supply)
func (_ThreePool *ThreePoolFilterer) ParseRemoveLiquidityImbalance(log types.Log) (*ThreePoolRemoveLiquidityImbalance, error) {
	event := new(ThreePoolRemoveLiquidityImbalance)
	if err := _ThreePool.contract.UnpackLog(event, "RemoveLiquidityImbalance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ThreePoolRemoveLiquidityOneIterator is returned from FilterRemoveLiquidityOne and is used to iterate over the raw logs and unpacked data for RemoveLiquidityOne events raised by the ThreePool contract.
type ThreePoolRemoveLiquidityOneIterator struct {
	Event *ThreePoolRemoveLiquidityOne // Event containing the contract specifics and raw log

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
func (it *ThreePoolRemoveLiquidityOneIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ThreePoolRemoveLiquidityOne)
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
		it.Event = new(ThreePoolRemoveLiquidityOne)
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
func (it *ThreePoolRemoveLiquidityOneIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ThreePoolRemoveLiquidityOneIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ThreePoolRemoveLiquidityOne represents a RemoveLiquidityOne event raised by the ThreePool contract.
type ThreePoolRemoveLiquidityOne struct {
	Provider    common.Address
	TokenAmount *big.Int
	CoinAmount  *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidityOne is a free log retrieval operation binding the contract event 0x9e96dd3b997a2a257eec4df9bb6eaf626e206df5f543bd963682d143300be310.
//
// Solidity: event RemoveLiquidityOne(address indexed provider, uint256 token_amount, uint256 coin_amount)
func (_ThreePool *ThreePoolFilterer) FilterRemoveLiquidityOne(opts *bind.FilterOpts, provider []common.Address) (*ThreePoolRemoveLiquidityOneIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _ThreePool.contract.FilterLogs(opts, "RemoveLiquidityOne", providerRule)
	if err != nil {
		return nil, err
	}
	return &ThreePoolRemoveLiquidityOneIterator{contract: _ThreePool.contract, event: "RemoveLiquidityOne", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidityOne is a free log subscription operation binding the contract event 0x9e96dd3b997a2a257eec4df9bb6eaf626e206df5f543bd963682d143300be310.
//
// Solidity: event RemoveLiquidityOne(address indexed provider, uint256 token_amount, uint256 coin_amount)
func (_ThreePool *ThreePoolFilterer) WatchRemoveLiquidityOne(opts *bind.WatchOpts, sink chan<- *ThreePoolRemoveLiquidityOne, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _ThreePool.contract.WatchLogs(opts, "RemoveLiquidityOne", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ThreePoolRemoveLiquidityOne)
				if err := _ThreePool.contract.UnpackLog(event, "RemoveLiquidityOne", log); err != nil {
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

// ParseRemoveLiquidityOne is a log parse operation binding the contract event 0x9e96dd3b997a2a257eec4df9bb6eaf626e206df5f543bd963682d143300be310.
//
// Solidity: event RemoveLiquidityOne(address indexed provider, uint256 token_amount, uint256 coin_amount)
func (_ThreePool *ThreePoolFilterer) ParseRemoveLiquidityOne(log types.Log) (*ThreePoolRemoveLiquidityOne, error) {
	event := new(ThreePoolRemoveLiquidityOne)
	if err := _ThreePool.contract.UnpackLog(event, "RemoveLiquidityOne", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ThreePoolStopRampAIterator is returned from FilterStopRampA and is used to iterate over the raw logs and unpacked data for StopRampA events raised by the ThreePool contract.
type ThreePoolStopRampAIterator struct {
	Event *ThreePoolStopRampA // Event containing the contract specifics and raw log

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
func (it *ThreePoolStopRampAIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ThreePoolStopRampA)
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
		it.Event = new(ThreePoolStopRampA)
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
func (it *ThreePoolStopRampAIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ThreePoolStopRampAIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ThreePoolStopRampA represents a StopRampA event raised by the ThreePool contract.
type ThreePoolStopRampA struct {
	A   *big.Int
	T   *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterStopRampA is a free log retrieval operation binding the contract event 0x46e22fb3709ad289f62ce63d469248536dbc78d82b84a3d7e74ad606dc201938.
//
// Solidity: event StopRampA(uint256 A, uint256 t)
func (_ThreePool *ThreePoolFilterer) FilterStopRampA(opts *bind.FilterOpts) (*ThreePoolStopRampAIterator, error) {

	logs, sub, err := _ThreePool.contract.FilterLogs(opts, "StopRampA")
	if err != nil {
		return nil, err
	}
	return &ThreePoolStopRampAIterator{contract: _ThreePool.contract, event: "StopRampA", logs: logs, sub: sub}, nil
}

// WatchStopRampA is a free log subscription operation binding the contract event 0x46e22fb3709ad289f62ce63d469248536dbc78d82b84a3d7e74ad606dc201938.
//
// Solidity: event StopRampA(uint256 A, uint256 t)
func (_ThreePool *ThreePoolFilterer) WatchStopRampA(opts *bind.WatchOpts, sink chan<- *ThreePoolStopRampA) (event.Subscription, error) {

	logs, sub, err := _ThreePool.contract.WatchLogs(opts, "StopRampA")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ThreePoolStopRampA)
				if err := _ThreePool.contract.UnpackLog(event, "StopRampA", log); err != nil {
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

// ParseStopRampA is a log parse operation binding the contract event 0x46e22fb3709ad289f62ce63d469248536dbc78d82b84a3d7e74ad606dc201938.
//
// Solidity: event StopRampA(uint256 A, uint256 t)
func (_ThreePool *ThreePoolFilterer) ParseStopRampA(log types.Log) (*ThreePoolStopRampA, error) {
	event := new(ThreePoolStopRampA)
	if err := _ThreePool.contract.UnpackLog(event, "StopRampA", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ThreePoolTokenExchangeIterator is returned from FilterTokenExchange and is used to iterate over the raw logs and unpacked data for TokenExchange events raised by the ThreePool contract.
type ThreePoolTokenExchangeIterator struct {
	Event *ThreePoolTokenExchange // Event containing the contract specifics and raw log

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
func (it *ThreePoolTokenExchangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ThreePoolTokenExchange)
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
		it.Event = new(ThreePoolTokenExchange)
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
func (it *ThreePoolTokenExchangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ThreePoolTokenExchangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ThreePoolTokenExchange represents a TokenExchange event raised by the ThreePool contract.
type ThreePoolTokenExchange struct {
	Buyer        common.Address
	SoldId       *big.Int
	TokensSold   *big.Int
	BoughtId     *big.Int
	TokensBought *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTokenExchange is a free log retrieval operation binding the contract event 0x8b3e96f2b889fa771c53c981b40daf005f63f637f1869f707052d15a3dd97140.
//
// Solidity: event TokenExchange(address indexed buyer, int128 sold_id, uint256 tokens_sold, int128 bought_id, uint256 tokens_bought)
func (_ThreePool *ThreePoolFilterer) FilterTokenExchange(opts *bind.FilterOpts, buyer []common.Address) (*ThreePoolTokenExchangeIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _ThreePool.contract.FilterLogs(opts, "TokenExchange", buyerRule)
	if err != nil {
		return nil, err
	}
	return &ThreePoolTokenExchangeIterator{contract: _ThreePool.contract, event: "TokenExchange", logs: logs, sub: sub}, nil
}

// WatchTokenExchange is a free log subscription operation binding the contract event 0x8b3e96f2b889fa771c53c981b40daf005f63f637f1869f707052d15a3dd97140.
//
// Solidity: event TokenExchange(address indexed buyer, int128 sold_id, uint256 tokens_sold, int128 bought_id, uint256 tokens_bought)
func (_ThreePool *ThreePoolFilterer) WatchTokenExchange(opts *bind.WatchOpts, sink chan<- *ThreePoolTokenExchange, buyer []common.Address) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _ThreePool.contract.WatchLogs(opts, "TokenExchange", buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ThreePoolTokenExchange)
				if err := _ThreePool.contract.UnpackLog(event, "TokenExchange", log); err != nil {
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

// ParseTokenExchange is a log parse operation binding the contract event 0x8b3e96f2b889fa771c53c981b40daf005f63f637f1869f707052d15a3dd97140.
//
// Solidity: event TokenExchange(address indexed buyer, int128 sold_id, uint256 tokens_sold, int128 bought_id, uint256 tokens_bought)
func (_ThreePool *ThreePoolFilterer) ParseTokenExchange(log types.Log) (*ThreePoolTokenExchange, error) {
	event := new(ThreePoolTokenExchange)
	if err := _ThreePool.contract.UnpackLog(event, "TokenExchange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
