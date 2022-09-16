// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package partybid

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

// PartyBidMetaData contains all meta data concerning the PartyBid contract.
var PartyBidMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_partyDAOMultisig\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenVaultFactory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_weth\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Bid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contributor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalContributed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"excessContribution\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"}],\"name\":\"Claimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contributor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousTotalContributedToParty\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalFromContributor\",\"type\":\"uint256\"}],\"name\":\"Contributed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumParty.PartyStatus\",\"name\":\"result\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalSpent\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalContributed\",\"type\":\"uint256\"}],\"name\":\"Finalized\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"auctionId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contributor\",\"type\":\"address\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"claimed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contribute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"contributions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"previousTotalContributedToParty\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"}],\"name\":\"emergencyCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"_returnData\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyForceLost\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"emergencyWithdrawEth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finalize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gatedToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gatedTokenAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contributor\",\"type\":\"address\"}],\"name\":\"getClaimAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_ethAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaximumBid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxBid\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaximumSpend\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxSpend\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"highestBid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_marketWrapper\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_auctionId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structStructs.AddressAndAmount\",\"name\":\"_split\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structStructs.AddressAndAmount\",\"name\":\"_tokenGate\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketWrapper\",\"outputs\":[{\"internalType\":\"contractIMarketWrapper\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nftContract\",\"outputs\":[{\"internalType\":\"contractIERC721Metadata\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"partyDAOMultisig\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"partyStatus\",\"outputs\":[{\"internalType\":\"enumParty.PartyStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"splitBasisPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"splitRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenVault\",\"outputs\":[{\"internalType\":\"contractITokenVault\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenVaultFactory\",\"outputs\":[{\"internalType\":\"contractIERC721VaultFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"totalContributed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalContributedToParty\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contributor\",\"type\":\"address\"}],\"name\":\"totalEthUsed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_total\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSpent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"valueToTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokens\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"weth\",\"outputs\":[{\"internalType\":\"contractIWETH\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// PartyBidABI is the input ABI used to generate the binding from.
// Deprecated: Use PartyBidMetaData.ABI instead.
var PartyBidABI = PartyBidMetaData.ABI

// PartyBid is an auto generated Go binding around an Ethereum contract.
type PartyBid struct {
	PartyBidCaller     // Read-only binding to the contract
	PartyBidTransactor // Write-only binding to the contract
	PartyBidFilterer   // Log filterer for contract events
}

// PartyBidCaller is an auto generated read-only Go binding around an Ethereum contract.
type PartyBidCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PartyBidTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PartyBidTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PartyBidFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PartyBidFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PartyBidSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PartyBidSession struct {
	Contract     *PartyBid         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PartyBidCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PartyBidCallerSession struct {
	Contract *PartyBidCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// PartyBidTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PartyBidTransactorSession struct {
	Contract     *PartyBidTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PartyBidRaw is an auto generated low-level Go binding around an Ethereum contract.
type PartyBidRaw struct {
	Contract *PartyBid // Generic contract binding to access the raw methods on
}

// PartyBidCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PartyBidCallerRaw struct {
	Contract *PartyBidCaller // Generic read-only contract binding to access the raw methods on
}

// PartyBidTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PartyBidTransactorRaw struct {
	Contract *PartyBidTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPartyBid creates a new instance of PartyBid, bound to a specific deployed contract.
func NewPartyBid(address common.Address, backend bind.ContractBackend) (*PartyBid, error) {
	contract, err := bindPartyBid(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PartyBid{PartyBidCaller: PartyBidCaller{contract: contract}, PartyBidTransactor: PartyBidTransactor{contract: contract}, PartyBidFilterer: PartyBidFilterer{contract: contract}}, nil
}

// NewPartyBidCaller creates a new read-only instance of PartyBid, bound to a specific deployed contract.
func NewPartyBidCaller(address common.Address, caller bind.ContractCaller) (*PartyBidCaller, error) {
	contract, err := bindPartyBid(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PartyBidCaller{contract: contract}, nil
}

// NewPartyBidTransactor creates a new write-only instance of PartyBid, bound to a specific deployed contract.
func NewPartyBidTransactor(address common.Address, transactor bind.ContractTransactor) (*PartyBidTransactor, error) {
	contract, err := bindPartyBid(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PartyBidTransactor{contract: contract}, nil
}

// NewPartyBidFilterer creates a new log filterer instance of PartyBid, bound to a specific deployed contract.
func NewPartyBidFilterer(address common.Address, filterer bind.ContractFilterer) (*PartyBidFilterer, error) {
	contract, err := bindPartyBid(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PartyBidFilterer{contract: contract}, nil
}

// bindPartyBid binds a generic wrapper to an already deployed contract.
func bindPartyBid(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PartyBidABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PartyBid *PartyBidRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PartyBid.Contract.PartyBidCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PartyBid *PartyBidRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PartyBid.Contract.PartyBidTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PartyBid *PartyBidRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PartyBid.Contract.PartyBidTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PartyBid *PartyBidCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PartyBid.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PartyBid *PartyBidTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PartyBid.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PartyBid *PartyBidTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PartyBid.Contract.contract.Transact(opts, method, params...)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint16)
func (_PartyBid *PartyBidCaller) VERSION(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _PartyBid.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint16)
func (_PartyBid *PartyBidSession) VERSION() (uint16, error) {
	return _PartyBid.Contract.VERSION(&_PartyBid.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint16)
func (_PartyBid *PartyBidCallerSession) VERSION() (uint16, error) {
	return _PartyBid.Contract.VERSION(&_PartyBid.CallOpts)
}

// AuctionId is a free data retrieval call binding the contract method 0x10782f8f.
//
// Solidity: function auctionId() view returns(uint256)
func (_PartyBid *PartyBidCaller) AuctionId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PartyBid.contract.Call(opts, &out, "auctionId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AuctionId is a free data retrieval call binding the contract method 0x10782f8f.
//
// Solidity: function auctionId() view returns(uint256)
func (_PartyBid *PartyBidSession) AuctionId() (*big.Int, error) {
	return _PartyBid.Contract.AuctionId(&_PartyBid.CallOpts)
}

// AuctionId is a free data retrieval call binding the contract method 0x10782f8f.
//
// Solidity: function auctionId() view returns(uint256)
func (_PartyBid *PartyBidCallerSession) AuctionId() (*big.Int, error) {
	return _PartyBid.Contract.AuctionId(&_PartyBid.CallOpts)
}

// Claimed is a free data retrieval call binding the contract method 0xc884ef83.
//
// Solidity: function claimed(address ) view returns(bool)
func (_PartyBid *PartyBidCaller) Claimed(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _PartyBid.contract.Call(opts, &out, "claimed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Claimed is a free data retrieval call binding the contract method 0xc884ef83.
//
// Solidity: function claimed(address ) view returns(bool)
func (_PartyBid *PartyBidSession) Claimed(arg0 common.Address) (bool, error) {
	return _PartyBid.Contract.Claimed(&_PartyBid.CallOpts, arg0)
}

// Claimed is a free data retrieval call binding the contract method 0xc884ef83.
//
// Solidity: function claimed(address ) view returns(bool)
func (_PartyBid *PartyBidCallerSession) Claimed(arg0 common.Address) (bool, error) {
	return _PartyBid.Contract.Claimed(&_PartyBid.CallOpts, arg0)
}

// Contributions is a free data retrieval call binding the contract method 0x550b521c.
//
// Solidity: function contributions(address , uint256 ) view returns(uint256 amount, uint256 previousTotalContributedToParty)
func (_PartyBid *PartyBidCaller) Contributions(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Amount                          *big.Int
	PreviousTotalContributedToParty *big.Int
}, error) {
	var out []interface{}
	err := _PartyBid.contract.Call(opts, &out, "contributions", arg0, arg1)

	outstruct := new(struct {
		Amount                          *big.Int
		PreviousTotalContributedToParty *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.PreviousTotalContributedToParty = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Contributions is a free data retrieval call binding the contract method 0x550b521c.
//
// Solidity: function contributions(address , uint256 ) view returns(uint256 amount, uint256 previousTotalContributedToParty)
func (_PartyBid *PartyBidSession) Contributions(arg0 common.Address, arg1 *big.Int) (struct {
	Amount                          *big.Int
	PreviousTotalContributedToParty *big.Int
}, error) {
	return _PartyBid.Contract.Contributions(&_PartyBid.CallOpts, arg0, arg1)
}

// Contributions is a free data retrieval call binding the contract method 0x550b521c.
//
// Solidity: function contributions(address , uint256 ) view returns(uint256 amount, uint256 previousTotalContributedToParty)
func (_PartyBid *PartyBidCallerSession) Contributions(arg0 common.Address, arg1 *big.Int) (struct {
	Amount                          *big.Int
	PreviousTotalContributedToParty *big.Int
}, error) {
	return _PartyBid.Contract.Contributions(&_PartyBid.CallOpts, arg0, arg1)
}

// GatedToken is a free data retrieval call binding the contract method 0x7a2ba9c4.
//
// Solidity: function gatedToken() view returns(address)
func (_PartyBid *PartyBidCaller) GatedToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PartyBid.contract.Call(opts, &out, "gatedToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GatedToken is a free data retrieval call binding the contract method 0x7a2ba9c4.
//
// Solidity: function gatedToken() view returns(address)
func (_PartyBid *PartyBidSession) GatedToken() (common.Address, error) {
	return _PartyBid.Contract.GatedToken(&_PartyBid.CallOpts)
}

// GatedToken is a free data retrieval call binding the contract method 0x7a2ba9c4.
//
// Solidity: function gatedToken() view returns(address)
func (_PartyBid *PartyBidCallerSession) GatedToken() (common.Address, error) {
	return _PartyBid.Contract.GatedToken(&_PartyBid.CallOpts)
}

// GatedTokenAmount is a free data retrieval call binding the contract method 0x8d42ecd6.
//
// Solidity: function gatedTokenAmount() view returns(uint256)
func (_PartyBid *PartyBidCaller) GatedTokenAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PartyBid.contract.Call(opts, &out, "gatedTokenAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GatedTokenAmount is a free data retrieval call binding the contract method 0x8d42ecd6.
//
// Solidity: function gatedTokenAmount() view returns(uint256)
func (_PartyBid *PartyBidSession) GatedTokenAmount() (*big.Int, error) {
	return _PartyBid.Contract.GatedTokenAmount(&_PartyBid.CallOpts)
}

// GatedTokenAmount is a free data retrieval call binding the contract method 0x8d42ecd6.
//
// Solidity: function gatedTokenAmount() view returns(uint256)
func (_PartyBid *PartyBidCallerSession) GatedTokenAmount() (*big.Int, error) {
	return _PartyBid.Contract.GatedTokenAmount(&_PartyBid.CallOpts)
}

// GetClaimAmounts is a free data retrieval call binding the contract method 0x6971524f.
//
// Solidity: function getClaimAmounts(address _contributor) view returns(uint256 _tokenAmount, uint256 _ethAmount)
func (_PartyBid *PartyBidCaller) GetClaimAmounts(opts *bind.CallOpts, _contributor common.Address) (struct {
	TokenAmount *big.Int
	EthAmount   *big.Int
}, error) {
	var out []interface{}
	err := _PartyBid.contract.Call(opts, &out, "getClaimAmounts", _contributor)

	outstruct := new(struct {
		TokenAmount *big.Int
		EthAmount   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokenAmount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.EthAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetClaimAmounts is a free data retrieval call binding the contract method 0x6971524f.
//
// Solidity: function getClaimAmounts(address _contributor) view returns(uint256 _tokenAmount, uint256 _ethAmount)
func (_PartyBid *PartyBidSession) GetClaimAmounts(_contributor common.Address) (struct {
	TokenAmount *big.Int
	EthAmount   *big.Int
}, error) {
	return _PartyBid.Contract.GetClaimAmounts(&_PartyBid.CallOpts, _contributor)
}

// GetClaimAmounts is a free data retrieval call binding the contract method 0x6971524f.
//
// Solidity: function getClaimAmounts(address _contributor) view returns(uint256 _tokenAmount, uint256 _ethAmount)
func (_PartyBid *PartyBidCallerSession) GetClaimAmounts(_contributor common.Address) (struct {
	TokenAmount *big.Int
	EthAmount   *big.Int
}, error) {
	return _PartyBid.Contract.GetClaimAmounts(&_PartyBid.CallOpts, _contributor)
}

// GetMaximumBid is a free data retrieval call binding the contract method 0xdd06ba1a.
//
// Solidity: function getMaximumBid() view returns(uint256 _maxBid)
func (_PartyBid *PartyBidCaller) GetMaximumBid(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PartyBid.contract.Call(opts, &out, "getMaximumBid")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaximumBid is a free data retrieval call binding the contract method 0xdd06ba1a.
//
// Solidity: function getMaximumBid() view returns(uint256 _maxBid)
func (_PartyBid *PartyBidSession) GetMaximumBid() (*big.Int, error) {
	return _PartyBid.Contract.GetMaximumBid(&_PartyBid.CallOpts)
}

// GetMaximumBid is a free data retrieval call binding the contract method 0xdd06ba1a.
//
// Solidity: function getMaximumBid() view returns(uint256 _maxBid)
func (_PartyBid *PartyBidCallerSession) GetMaximumBid() (*big.Int, error) {
	return _PartyBid.Contract.GetMaximumBid(&_PartyBid.CallOpts)
}

// GetMaximumSpend is a free data retrieval call binding the contract method 0xacd13c59.
//
// Solidity: function getMaximumSpend() view returns(uint256 _maxSpend)
func (_PartyBid *PartyBidCaller) GetMaximumSpend(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PartyBid.contract.Call(opts, &out, "getMaximumSpend")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaximumSpend is a free data retrieval call binding the contract method 0xacd13c59.
//
// Solidity: function getMaximumSpend() view returns(uint256 _maxSpend)
func (_PartyBid *PartyBidSession) GetMaximumSpend() (*big.Int, error) {
	return _PartyBid.Contract.GetMaximumSpend(&_PartyBid.CallOpts)
}

// GetMaximumSpend is a free data retrieval call binding the contract method 0xacd13c59.
//
// Solidity: function getMaximumSpend() view returns(uint256 _maxSpend)
func (_PartyBid *PartyBidCallerSession) GetMaximumSpend() (*big.Int, error) {
	return _PartyBid.Contract.GetMaximumSpend(&_PartyBid.CallOpts)
}

// HighestBid is a free data retrieval call binding the contract method 0xd57bde79.
//
// Solidity: function highestBid() view returns(uint256)
func (_PartyBid *PartyBidCaller) HighestBid(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PartyBid.contract.Call(opts, &out, "highestBid")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HighestBid is a free data retrieval call binding the contract method 0xd57bde79.
//
// Solidity: function highestBid() view returns(uint256)
func (_PartyBid *PartyBidSession) HighestBid() (*big.Int, error) {
	return _PartyBid.Contract.HighestBid(&_PartyBid.CallOpts)
}

// HighestBid is a free data retrieval call binding the contract method 0xd57bde79.
//
// Solidity: function highestBid() view returns(uint256)
func (_PartyBid *PartyBidCallerSession) HighestBid() (*big.Int, error) {
	return _PartyBid.Contract.HighestBid(&_PartyBid.CallOpts)
}

// MarketWrapper is a free data retrieval call binding the contract method 0xef38bf01.
//
// Solidity: function marketWrapper() view returns(address)
func (_PartyBid *PartyBidCaller) MarketWrapper(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PartyBid.contract.Call(opts, &out, "marketWrapper")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MarketWrapper is a free data retrieval call binding the contract method 0xef38bf01.
//
// Solidity: function marketWrapper() view returns(address)
func (_PartyBid *PartyBidSession) MarketWrapper() (common.Address, error) {
	return _PartyBid.Contract.MarketWrapper(&_PartyBid.CallOpts)
}

// MarketWrapper is a free data retrieval call binding the contract method 0xef38bf01.
//
// Solidity: function marketWrapper() view returns(address)
func (_PartyBid *PartyBidCallerSession) MarketWrapper() (common.Address, error) {
	return _PartyBid.Contract.MarketWrapper(&_PartyBid.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PartyBid *PartyBidCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PartyBid.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PartyBid *PartyBidSession) Name() (string, error) {
	return _PartyBid.Contract.Name(&_PartyBid.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PartyBid *PartyBidCallerSession) Name() (string, error) {
	return _PartyBid.Contract.Name(&_PartyBid.CallOpts)
}

// NftContract is a free data retrieval call binding the contract method 0xd56d229d.
//
// Solidity: function nftContract() view returns(address)
func (_PartyBid *PartyBidCaller) NftContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PartyBid.contract.Call(opts, &out, "nftContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NftContract is a free data retrieval call binding the contract method 0xd56d229d.
//
// Solidity: function nftContract() view returns(address)
func (_PartyBid *PartyBidSession) NftContract() (common.Address, error) {
	return _PartyBid.Contract.NftContract(&_PartyBid.CallOpts)
}

// NftContract is a free data retrieval call binding the contract method 0xd56d229d.
//
// Solidity: function nftContract() view returns(address)
func (_PartyBid *PartyBidCallerSession) NftContract() (common.Address, error) {
	return _PartyBid.Contract.NftContract(&_PartyBid.CallOpts)
}

// PartyDAOMultisig is a free data retrieval call binding the contract method 0x3c4d12d9.
//
// Solidity: function partyDAOMultisig() view returns(address)
func (_PartyBid *PartyBidCaller) PartyDAOMultisig(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PartyBid.contract.Call(opts, &out, "partyDAOMultisig")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PartyDAOMultisig is a free data retrieval call binding the contract method 0x3c4d12d9.
//
// Solidity: function partyDAOMultisig() view returns(address)
func (_PartyBid *PartyBidSession) PartyDAOMultisig() (common.Address, error) {
	return _PartyBid.Contract.PartyDAOMultisig(&_PartyBid.CallOpts)
}

// PartyDAOMultisig is a free data retrieval call binding the contract method 0x3c4d12d9.
//
// Solidity: function partyDAOMultisig() view returns(address)
func (_PartyBid *PartyBidCallerSession) PartyDAOMultisig() (common.Address, error) {
	return _PartyBid.Contract.PartyDAOMultisig(&_PartyBid.CallOpts)
}

// PartyStatus is a free data retrieval call binding the contract method 0xdf51c07f.
//
// Solidity: function partyStatus() view returns(uint8)
func (_PartyBid *PartyBidCaller) PartyStatus(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _PartyBid.contract.Call(opts, &out, "partyStatus")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// PartyStatus is a free data retrieval call binding the contract method 0xdf51c07f.
//
// Solidity: function partyStatus() view returns(uint8)
func (_PartyBid *PartyBidSession) PartyStatus() (uint8, error) {
	return _PartyBid.Contract.PartyStatus(&_PartyBid.CallOpts)
}

// PartyStatus is a free data retrieval call binding the contract method 0xdf51c07f.
//
// Solidity: function partyStatus() view returns(uint8)
func (_PartyBid *PartyBidCallerSession) PartyStatus() (uint8, error) {
	return _PartyBid.Contract.PartyStatus(&_PartyBid.CallOpts)
}

// SplitBasisPoints is a free data retrieval call binding the contract method 0x2bbce5e6.
//
// Solidity: function splitBasisPoints() view returns(uint256)
func (_PartyBid *PartyBidCaller) SplitBasisPoints(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PartyBid.contract.Call(opts, &out, "splitBasisPoints")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SplitBasisPoints is a free data retrieval call binding the contract method 0x2bbce5e6.
//
// Solidity: function splitBasisPoints() view returns(uint256)
func (_PartyBid *PartyBidSession) SplitBasisPoints() (*big.Int, error) {
	return _PartyBid.Contract.SplitBasisPoints(&_PartyBid.CallOpts)
}

// SplitBasisPoints is a free data retrieval call binding the contract method 0x2bbce5e6.
//
// Solidity: function splitBasisPoints() view returns(uint256)
func (_PartyBid *PartyBidCallerSession) SplitBasisPoints() (*big.Int, error) {
	return _PartyBid.Contract.SplitBasisPoints(&_PartyBid.CallOpts)
}

// SplitRecipient is a free data retrieval call binding the contract method 0x4367a029.
//
// Solidity: function splitRecipient() view returns(address)
func (_PartyBid *PartyBidCaller) SplitRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PartyBid.contract.Call(opts, &out, "splitRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SplitRecipient is a free data retrieval call binding the contract method 0x4367a029.
//
// Solidity: function splitRecipient() view returns(address)
func (_PartyBid *PartyBidSession) SplitRecipient() (common.Address, error) {
	return _PartyBid.Contract.SplitRecipient(&_PartyBid.CallOpts)
}

// SplitRecipient is a free data retrieval call binding the contract method 0x4367a029.
//
// Solidity: function splitRecipient() view returns(address)
func (_PartyBid *PartyBidCallerSession) SplitRecipient() (common.Address, error) {
	return _PartyBid.Contract.SplitRecipient(&_PartyBid.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PartyBid *PartyBidCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PartyBid.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PartyBid *PartyBidSession) Symbol() (string, error) {
	return _PartyBid.Contract.Symbol(&_PartyBid.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PartyBid *PartyBidCallerSession) Symbol() (string, error) {
	return _PartyBid.Contract.Symbol(&_PartyBid.CallOpts)
}

// TokenId is a free data retrieval call binding the contract method 0x17d70f7c.
//
// Solidity: function tokenId() view returns(uint256)
func (_PartyBid *PartyBidCaller) TokenId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PartyBid.contract.Call(opts, &out, "tokenId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenId is a free data retrieval call binding the contract method 0x17d70f7c.
//
// Solidity: function tokenId() view returns(uint256)
func (_PartyBid *PartyBidSession) TokenId() (*big.Int, error) {
	return _PartyBid.Contract.TokenId(&_PartyBid.CallOpts)
}

// TokenId is a free data retrieval call binding the contract method 0x17d70f7c.
//
// Solidity: function tokenId() view returns(uint256)
func (_PartyBid *PartyBidCallerSession) TokenId() (*big.Int, error) {
	return _PartyBid.Contract.TokenId(&_PartyBid.CallOpts)
}

// TokenVault is a free data retrieval call binding the contract method 0x5bc789d9.
//
// Solidity: function tokenVault() view returns(address)
func (_PartyBid *PartyBidCaller) TokenVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PartyBid.contract.Call(opts, &out, "tokenVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenVault is a free data retrieval call binding the contract method 0x5bc789d9.
//
// Solidity: function tokenVault() view returns(address)
func (_PartyBid *PartyBidSession) TokenVault() (common.Address, error) {
	return _PartyBid.Contract.TokenVault(&_PartyBid.CallOpts)
}

// TokenVault is a free data retrieval call binding the contract method 0x5bc789d9.
//
// Solidity: function tokenVault() view returns(address)
func (_PartyBid *PartyBidCallerSession) TokenVault() (common.Address, error) {
	return _PartyBid.Contract.TokenVault(&_PartyBid.CallOpts)
}

// TokenVaultFactory is a free data retrieval call binding the contract method 0x0b203023.
//
// Solidity: function tokenVaultFactory() view returns(address)
func (_PartyBid *PartyBidCaller) TokenVaultFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PartyBid.contract.Call(opts, &out, "tokenVaultFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenVaultFactory is a free data retrieval call binding the contract method 0x0b203023.
//
// Solidity: function tokenVaultFactory() view returns(address)
func (_PartyBid *PartyBidSession) TokenVaultFactory() (common.Address, error) {
	return _PartyBid.Contract.TokenVaultFactory(&_PartyBid.CallOpts)
}

// TokenVaultFactory is a free data retrieval call binding the contract method 0x0b203023.
//
// Solidity: function tokenVaultFactory() view returns(address)
func (_PartyBid *PartyBidCallerSession) TokenVaultFactory() (common.Address, error) {
	return _PartyBid.Contract.TokenVaultFactory(&_PartyBid.CallOpts)
}

// TotalContributed is a free data retrieval call binding the contract method 0xa0f243b8.
//
// Solidity: function totalContributed(address ) view returns(uint256)
func (_PartyBid *PartyBidCaller) TotalContributed(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PartyBid.contract.Call(opts, &out, "totalContributed", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalContributed is a free data retrieval call binding the contract method 0xa0f243b8.
//
// Solidity: function totalContributed(address ) view returns(uint256)
func (_PartyBid *PartyBidSession) TotalContributed(arg0 common.Address) (*big.Int, error) {
	return _PartyBid.Contract.TotalContributed(&_PartyBid.CallOpts, arg0)
}

// TotalContributed is a free data retrieval call binding the contract method 0xa0f243b8.
//
// Solidity: function totalContributed(address ) view returns(uint256)
func (_PartyBid *PartyBidCallerSession) TotalContributed(arg0 common.Address) (*big.Int, error) {
	return _PartyBid.Contract.TotalContributed(&_PartyBid.CallOpts, arg0)
}

// TotalContributedToParty is a free data retrieval call binding the contract method 0x82a5c69a.
//
// Solidity: function totalContributedToParty() view returns(uint256)
func (_PartyBid *PartyBidCaller) TotalContributedToParty(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PartyBid.contract.Call(opts, &out, "totalContributedToParty")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalContributedToParty is a free data retrieval call binding the contract method 0x82a5c69a.
//
// Solidity: function totalContributedToParty() view returns(uint256)
func (_PartyBid *PartyBidSession) TotalContributedToParty() (*big.Int, error) {
	return _PartyBid.Contract.TotalContributedToParty(&_PartyBid.CallOpts)
}

// TotalContributedToParty is a free data retrieval call binding the contract method 0x82a5c69a.
//
// Solidity: function totalContributedToParty() view returns(uint256)
func (_PartyBid *PartyBidCallerSession) TotalContributedToParty() (*big.Int, error) {
	return _PartyBid.Contract.TotalContributedToParty(&_PartyBid.CallOpts)
}

// TotalEthUsed is a free data retrieval call binding the contract method 0x25b42a26.
//
// Solidity: function totalEthUsed(address _contributor) view returns(uint256 _total)
func (_PartyBid *PartyBidCaller) TotalEthUsed(opts *bind.CallOpts, _contributor common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PartyBid.contract.Call(opts, &out, "totalEthUsed", _contributor)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalEthUsed is a free data retrieval call binding the contract method 0x25b42a26.
//
// Solidity: function totalEthUsed(address _contributor) view returns(uint256 _total)
func (_PartyBid *PartyBidSession) TotalEthUsed(_contributor common.Address) (*big.Int, error) {
	return _PartyBid.Contract.TotalEthUsed(&_PartyBid.CallOpts, _contributor)
}

// TotalEthUsed is a free data retrieval call binding the contract method 0x25b42a26.
//
// Solidity: function totalEthUsed(address _contributor) view returns(uint256 _total)
func (_PartyBid *PartyBidCallerSession) TotalEthUsed(_contributor common.Address) (*big.Int, error) {
	return _PartyBid.Contract.TotalEthUsed(&_PartyBid.CallOpts, _contributor)
}

// TotalSpent is a free data retrieval call binding the contract method 0xfb346eab.
//
// Solidity: function totalSpent() view returns(uint256)
func (_PartyBid *PartyBidCaller) TotalSpent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PartyBid.contract.Call(opts, &out, "totalSpent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSpent is a free data retrieval call binding the contract method 0xfb346eab.
//
// Solidity: function totalSpent() view returns(uint256)
func (_PartyBid *PartyBidSession) TotalSpent() (*big.Int, error) {
	return _PartyBid.Contract.TotalSpent(&_PartyBid.CallOpts)
}

// TotalSpent is a free data retrieval call binding the contract method 0xfb346eab.
//
// Solidity: function totalSpent() view returns(uint256)
func (_PartyBid *PartyBidCallerSession) TotalSpent() (*big.Int, error) {
	return _PartyBid.Contract.TotalSpent(&_PartyBid.CallOpts)
}

// ValueToTokens is a free data retrieval call binding the contract method 0x9744b8dc.
//
// Solidity: function valueToTokens(uint256 _value) pure returns(uint256 _tokens)
func (_PartyBid *PartyBidCaller) ValueToTokens(opts *bind.CallOpts, _value *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PartyBid.contract.Call(opts, &out, "valueToTokens", _value)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValueToTokens is a free data retrieval call binding the contract method 0x9744b8dc.
//
// Solidity: function valueToTokens(uint256 _value) pure returns(uint256 _tokens)
func (_PartyBid *PartyBidSession) ValueToTokens(_value *big.Int) (*big.Int, error) {
	return _PartyBid.Contract.ValueToTokens(&_PartyBid.CallOpts, _value)
}

// ValueToTokens is a free data retrieval call binding the contract method 0x9744b8dc.
//
// Solidity: function valueToTokens(uint256 _value) pure returns(uint256 _tokens)
func (_PartyBid *PartyBidCallerSession) ValueToTokens(_value *big.Int) (*big.Int, error) {
	return _PartyBid.Contract.ValueToTokens(&_PartyBid.CallOpts, _value)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_PartyBid *PartyBidCaller) Weth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PartyBid.contract.Call(opts, &out, "weth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_PartyBid *PartyBidSession) Weth() (common.Address, error) {
	return _PartyBid.Contract.Weth(&_PartyBid.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_PartyBid *PartyBidCallerSession) Weth() (common.Address, error) {
	return _PartyBid.Contract.Weth(&_PartyBid.CallOpts)
}

// Bid is a paid mutator transaction binding the contract method 0x1998aeef.
//
// Solidity: function bid() returns()
func (_PartyBid *PartyBidTransactor) Bid(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PartyBid.contract.Transact(opts, "bid")
}

// Bid is a paid mutator transaction binding the contract method 0x1998aeef.
//
// Solidity: function bid() returns()
func (_PartyBid *PartyBidSession) Bid() (*types.Transaction, error) {
	return _PartyBid.Contract.Bid(&_PartyBid.TransactOpts)
}

// Bid is a paid mutator transaction binding the contract method 0x1998aeef.
//
// Solidity: function bid() returns()
func (_PartyBid *PartyBidTransactorSession) Bid() (*types.Transaction, error) {
	return _PartyBid.Contract.Bid(&_PartyBid.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address _contributor) returns()
func (_PartyBid *PartyBidTransactor) Claim(opts *bind.TransactOpts, _contributor common.Address) (*types.Transaction, error) {
	return _PartyBid.contract.Transact(opts, "claim", _contributor)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address _contributor) returns()
func (_PartyBid *PartyBidSession) Claim(_contributor common.Address) (*types.Transaction, error) {
	return _PartyBid.Contract.Claim(&_PartyBid.TransactOpts, _contributor)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address _contributor) returns()
func (_PartyBid *PartyBidTransactorSession) Claim(_contributor common.Address) (*types.Transaction, error) {
	return _PartyBid.Contract.Claim(&_PartyBid.TransactOpts, _contributor)
}

// Contribute is a paid mutator transaction binding the contract method 0xd7bb99ba.
//
// Solidity: function contribute() payable returns()
func (_PartyBid *PartyBidTransactor) Contribute(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PartyBid.contract.Transact(opts, "contribute")
}

// Contribute is a paid mutator transaction binding the contract method 0xd7bb99ba.
//
// Solidity: function contribute() payable returns()
func (_PartyBid *PartyBidSession) Contribute() (*types.Transaction, error) {
	return _PartyBid.Contract.Contribute(&_PartyBid.TransactOpts)
}

// Contribute is a paid mutator transaction binding the contract method 0xd7bb99ba.
//
// Solidity: function contribute() payable returns()
func (_PartyBid *PartyBidTransactorSession) Contribute() (*types.Transaction, error) {
	return _PartyBid.Contract.Contribute(&_PartyBid.TransactOpts)
}

// EmergencyCall is a paid mutator transaction binding the contract method 0xc4bf0220.
//
// Solidity: function emergencyCall(address _contract, bytes _calldata) returns(bool _success, bytes _returnData)
func (_PartyBid *PartyBidTransactor) EmergencyCall(opts *bind.TransactOpts, _contract common.Address, _calldata []byte) (*types.Transaction, error) {
	return _PartyBid.contract.Transact(opts, "emergencyCall", _contract, _calldata)
}

// EmergencyCall is a paid mutator transaction binding the contract method 0xc4bf0220.
//
// Solidity: function emergencyCall(address _contract, bytes _calldata) returns(bool _success, bytes _returnData)
func (_PartyBid *PartyBidSession) EmergencyCall(_contract common.Address, _calldata []byte) (*types.Transaction, error) {
	return _PartyBid.Contract.EmergencyCall(&_PartyBid.TransactOpts, _contract, _calldata)
}

// EmergencyCall is a paid mutator transaction binding the contract method 0xc4bf0220.
//
// Solidity: function emergencyCall(address _contract, bytes _calldata) returns(bool _success, bytes _returnData)
func (_PartyBid *PartyBidTransactorSession) EmergencyCall(_contract common.Address, _calldata []byte) (*types.Transaction, error) {
	return _PartyBid.Contract.EmergencyCall(&_PartyBid.TransactOpts, _contract, _calldata)
}

// EmergencyForceLost is a paid mutator transaction binding the contract method 0x17821fdc.
//
// Solidity: function emergencyForceLost() returns()
func (_PartyBid *PartyBidTransactor) EmergencyForceLost(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PartyBid.contract.Transact(opts, "emergencyForceLost")
}

// EmergencyForceLost is a paid mutator transaction binding the contract method 0x17821fdc.
//
// Solidity: function emergencyForceLost() returns()
func (_PartyBid *PartyBidSession) EmergencyForceLost() (*types.Transaction, error) {
	return _PartyBid.Contract.EmergencyForceLost(&_PartyBid.TransactOpts)
}

// EmergencyForceLost is a paid mutator transaction binding the contract method 0x17821fdc.
//
// Solidity: function emergencyForceLost() returns()
func (_PartyBid *PartyBidTransactorSession) EmergencyForceLost() (*types.Transaction, error) {
	return _PartyBid.Contract.EmergencyForceLost(&_PartyBid.TransactOpts)
}

// EmergencyWithdrawEth is a paid mutator transaction binding the contract method 0x429093cc.
//
// Solidity: function emergencyWithdrawEth(uint256 _value) returns()
func (_PartyBid *PartyBidTransactor) EmergencyWithdrawEth(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _PartyBid.contract.Transact(opts, "emergencyWithdrawEth", _value)
}

// EmergencyWithdrawEth is a paid mutator transaction binding the contract method 0x429093cc.
//
// Solidity: function emergencyWithdrawEth(uint256 _value) returns()
func (_PartyBid *PartyBidSession) EmergencyWithdrawEth(_value *big.Int) (*types.Transaction, error) {
	return _PartyBid.Contract.EmergencyWithdrawEth(&_PartyBid.TransactOpts, _value)
}

// EmergencyWithdrawEth is a paid mutator transaction binding the contract method 0x429093cc.
//
// Solidity: function emergencyWithdrawEth(uint256 _value) returns()
func (_PartyBid *PartyBidTransactorSession) EmergencyWithdrawEth(_value *big.Int) (*types.Transaction, error) {
	return _PartyBid.Contract.EmergencyWithdrawEth(&_PartyBid.TransactOpts, _value)
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_PartyBid *PartyBidTransactor) Finalize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PartyBid.contract.Transact(opts, "finalize")
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_PartyBid *PartyBidSession) Finalize() (*types.Transaction, error) {
	return _PartyBid.Contract.Finalize(&_PartyBid.TransactOpts)
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_PartyBid *PartyBidTransactorSession) Finalize() (*types.Transaction, error) {
	return _PartyBid.Contract.Finalize(&_PartyBid.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x1d4667fa.
//
// Solidity: function initialize(address _marketWrapper, address _nftContract, uint256 _tokenId, uint256 _auctionId, (address,uint256) _split, (address,uint256) _tokenGate, string _name, string _symbol) returns()
func (_PartyBid *PartyBidTransactor) Initialize(opts *bind.TransactOpts, _marketWrapper common.Address, _nftContract common.Address, _tokenId *big.Int, _auctionId *big.Int, _split StructsAddressAndAmount, _tokenGate StructsAddressAndAmount, _name string, _symbol string) (*types.Transaction, error) {
	return _PartyBid.contract.Transact(opts, "initialize", _marketWrapper, _nftContract, _tokenId, _auctionId, _split, _tokenGate, _name, _symbol)
}

// Initialize is a paid mutator transaction binding the contract method 0x1d4667fa.
//
// Solidity: function initialize(address _marketWrapper, address _nftContract, uint256 _tokenId, uint256 _auctionId, (address,uint256) _split, (address,uint256) _tokenGate, string _name, string _symbol) returns()
func (_PartyBid *PartyBidSession) Initialize(_marketWrapper common.Address, _nftContract common.Address, _tokenId *big.Int, _auctionId *big.Int, _split StructsAddressAndAmount, _tokenGate StructsAddressAndAmount, _name string, _symbol string) (*types.Transaction, error) {
	return _PartyBid.Contract.Initialize(&_PartyBid.TransactOpts, _marketWrapper, _nftContract, _tokenId, _auctionId, _split, _tokenGate, _name, _symbol)
}

// Initialize is a paid mutator transaction binding the contract method 0x1d4667fa.
//
// Solidity: function initialize(address _marketWrapper, address _nftContract, uint256 _tokenId, uint256 _auctionId, (address,uint256) _split, (address,uint256) _tokenGate, string _name, string _symbol) returns()
func (_PartyBid *PartyBidTransactorSession) Initialize(_marketWrapper common.Address, _nftContract common.Address, _tokenId *big.Int, _auctionId *big.Int, _split StructsAddressAndAmount, _tokenGate StructsAddressAndAmount, _name string, _symbol string) (*types.Transaction, error) {
	return _PartyBid.Contract.Initialize(&_PartyBid.TransactOpts, _marketWrapper, _nftContract, _tokenId, _auctionId, _split, _tokenGate, _name, _symbol)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_PartyBid *PartyBidTransactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _PartyBid.contract.Transact(opts, "onERC721Received", arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_PartyBid *PartyBidSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _PartyBid.Contract.OnERC721Received(&_PartyBid.TransactOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_PartyBid *PartyBidTransactorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _PartyBid.Contract.OnERC721Received(&_PartyBid.TransactOpts, arg0, arg1, arg2, arg3)
}

// PartyBidBidIterator is returned from FilterBid and is used to iterate over the raw logs and unpacked data for Bid events raised by the PartyBid contract.
type PartyBidBidIterator struct {
	Event *PartyBidBid // Event containing the contract specifics and raw log

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
func (it *PartyBidBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PartyBidBid)
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
		it.Event = new(PartyBidBid)
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
func (it *PartyBidBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PartyBidBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PartyBidBid represents a Bid event raised by the PartyBid contract.
type PartyBidBid struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBid is a free log retrieval operation binding the contract event 0x7a183e84509e3fe5b0b3aac15347fd1c7d71fd1503001f1a1d7c9658077eb35f.
//
// Solidity: event Bid(uint256 amount)
func (_PartyBid *PartyBidFilterer) FilterBid(opts *bind.FilterOpts) (*PartyBidBidIterator, error) {

	logs, sub, err := _PartyBid.contract.FilterLogs(opts, "Bid")
	if err != nil {
		return nil, err
	}
	return &PartyBidBidIterator{contract: _PartyBid.contract, event: "Bid", logs: logs, sub: sub}, nil
}

// WatchBid is a free log subscription operation binding the contract event 0x7a183e84509e3fe5b0b3aac15347fd1c7d71fd1503001f1a1d7c9658077eb35f.
//
// Solidity: event Bid(uint256 amount)
func (_PartyBid *PartyBidFilterer) WatchBid(opts *bind.WatchOpts, sink chan<- *PartyBidBid) (event.Subscription, error) {

	logs, sub, err := _PartyBid.contract.WatchLogs(opts, "Bid")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PartyBidBid)
				if err := _PartyBid.contract.UnpackLog(event, "Bid", log); err != nil {
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

// ParseBid is a log parse operation binding the contract event 0x7a183e84509e3fe5b0b3aac15347fd1c7d71fd1503001f1a1d7c9658077eb35f.
//
// Solidity: event Bid(uint256 amount)
func (_PartyBid *PartyBidFilterer) ParseBid(log types.Log) (*PartyBidBid, error) {
	event := new(PartyBidBid)
	if err := _PartyBid.contract.UnpackLog(event, "Bid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PartyBidClaimedIterator is returned from FilterClaimed and is used to iterate over the raw logs and unpacked data for Claimed events raised by the PartyBid contract.
type PartyBidClaimedIterator struct {
	Event *PartyBidClaimed // Event containing the contract specifics and raw log

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
func (it *PartyBidClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PartyBidClaimed)
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
		it.Event = new(PartyBidClaimed)
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
func (it *PartyBidClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PartyBidClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PartyBidClaimed represents a Claimed event raised by the PartyBid contract.
type PartyBidClaimed struct {
	Contributor        common.Address
	TotalContributed   *big.Int
	ExcessContribution *big.Int
	TokenAmount        *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterClaimed is a free log retrieval operation binding the contract event 0x9cdcf2f7714cca3508c7f0110b04a90a80a3a8dd0e35de99689db74d28c5383e.
//
// Solidity: event Claimed(address indexed contributor, uint256 totalContributed, uint256 excessContribution, uint256 tokenAmount)
func (_PartyBid *PartyBidFilterer) FilterClaimed(opts *bind.FilterOpts, contributor []common.Address) (*PartyBidClaimedIterator, error) {

	var contributorRule []interface{}
	for _, contributorItem := range contributor {
		contributorRule = append(contributorRule, contributorItem)
	}

	logs, sub, err := _PartyBid.contract.FilterLogs(opts, "Claimed", contributorRule)
	if err != nil {
		return nil, err
	}
	return &PartyBidClaimedIterator{contract: _PartyBid.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0x9cdcf2f7714cca3508c7f0110b04a90a80a3a8dd0e35de99689db74d28c5383e.
//
// Solidity: event Claimed(address indexed contributor, uint256 totalContributed, uint256 excessContribution, uint256 tokenAmount)
func (_PartyBid *PartyBidFilterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *PartyBidClaimed, contributor []common.Address) (event.Subscription, error) {

	var contributorRule []interface{}
	for _, contributorItem := range contributor {
		contributorRule = append(contributorRule, contributorItem)
	}

	logs, sub, err := _PartyBid.contract.WatchLogs(opts, "Claimed", contributorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PartyBidClaimed)
				if err := _PartyBid.contract.UnpackLog(event, "Claimed", log); err != nil {
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

// ParseClaimed is a log parse operation binding the contract event 0x9cdcf2f7714cca3508c7f0110b04a90a80a3a8dd0e35de99689db74d28c5383e.
//
// Solidity: event Claimed(address indexed contributor, uint256 totalContributed, uint256 excessContribution, uint256 tokenAmount)
func (_PartyBid *PartyBidFilterer) ParseClaimed(log types.Log) (*PartyBidClaimed, error) {
	event := new(PartyBidClaimed)
	if err := _PartyBid.contract.UnpackLog(event, "Claimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PartyBidContributedIterator is returned from FilterContributed and is used to iterate over the raw logs and unpacked data for Contributed events raised by the PartyBid contract.
type PartyBidContributedIterator struct {
	Event *PartyBidContributed // Event containing the contract specifics and raw log

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
func (it *PartyBidContributedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PartyBidContributed)
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
		it.Event = new(PartyBidContributed)
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
func (it *PartyBidContributedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PartyBidContributedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PartyBidContributed represents a Contributed event raised by the PartyBid contract.
type PartyBidContributed struct {
	Contributor                     common.Address
	Amount                          *big.Int
	PreviousTotalContributedToParty *big.Int
	TotalFromContributor            *big.Int
	Raw                             types.Log // Blockchain specific contextual infos
}

// FilterContributed is a free log retrieval operation binding the contract event 0xb2623081601722547aae8781994e01a1974d95b0ad9ce6a0cfbe17487556257f.
//
// Solidity: event Contributed(address indexed contributor, uint256 amount, uint256 previousTotalContributedToParty, uint256 totalFromContributor)
func (_PartyBid *PartyBidFilterer) FilterContributed(opts *bind.FilterOpts, contributor []common.Address) (*PartyBidContributedIterator, error) {

	var contributorRule []interface{}
	for _, contributorItem := range contributor {
		contributorRule = append(contributorRule, contributorItem)
	}

	logs, sub, err := _PartyBid.contract.FilterLogs(opts, "Contributed", contributorRule)
	if err != nil {
		return nil, err
	}
	return &PartyBidContributedIterator{contract: _PartyBid.contract, event: "Contributed", logs: logs, sub: sub}, nil
}

// WatchContributed is a free log subscription operation binding the contract event 0xb2623081601722547aae8781994e01a1974d95b0ad9ce6a0cfbe17487556257f.
//
// Solidity: event Contributed(address indexed contributor, uint256 amount, uint256 previousTotalContributedToParty, uint256 totalFromContributor)
func (_PartyBid *PartyBidFilterer) WatchContributed(opts *bind.WatchOpts, sink chan<- *PartyBidContributed, contributor []common.Address) (event.Subscription, error) {

	var contributorRule []interface{}
	for _, contributorItem := range contributor {
		contributorRule = append(contributorRule, contributorItem)
	}

	logs, sub, err := _PartyBid.contract.WatchLogs(opts, "Contributed", contributorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PartyBidContributed)
				if err := _PartyBid.contract.UnpackLog(event, "Contributed", log); err != nil {
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

// ParseContributed is a log parse operation binding the contract event 0xb2623081601722547aae8781994e01a1974d95b0ad9ce6a0cfbe17487556257f.
//
// Solidity: event Contributed(address indexed contributor, uint256 amount, uint256 previousTotalContributedToParty, uint256 totalFromContributor)
func (_PartyBid *PartyBidFilterer) ParseContributed(log types.Log) (*PartyBidContributed, error) {
	event := new(PartyBidContributed)
	if err := _PartyBid.contract.UnpackLog(event, "Contributed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PartyBidFinalizedIterator is returned from FilterFinalized and is used to iterate over the raw logs and unpacked data for Finalized events raised by the PartyBid contract.
type PartyBidFinalizedIterator struct {
	Event *PartyBidFinalized // Event containing the contract specifics and raw log

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
func (it *PartyBidFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PartyBidFinalized)
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
		it.Event = new(PartyBidFinalized)
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
func (it *PartyBidFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PartyBidFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PartyBidFinalized represents a Finalized event raised by the PartyBid contract.
type PartyBidFinalized struct {
	Result           uint8
	TotalSpent       *big.Int
	Fee              *big.Int
	TotalContributed *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterFinalized is a free log retrieval operation binding the contract event 0x9a2087478f16b801ecd568a6676f5db758bda2a01b954b2c754257d11eb3770b.
//
// Solidity: event Finalized(uint8 result, uint256 totalSpent, uint256 fee, uint256 totalContributed)
func (_PartyBid *PartyBidFilterer) FilterFinalized(opts *bind.FilterOpts) (*PartyBidFinalizedIterator, error) {

	logs, sub, err := _PartyBid.contract.FilterLogs(opts, "Finalized")
	if err != nil {
		return nil, err
	}
	return &PartyBidFinalizedIterator{contract: _PartyBid.contract, event: "Finalized", logs: logs, sub: sub}, nil
}

// WatchFinalized is a free log subscription operation binding the contract event 0x9a2087478f16b801ecd568a6676f5db758bda2a01b954b2c754257d11eb3770b.
//
// Solidity: event Finalized(uint8 result, uint256 totalSpent, uint256 fee, uint256 totalContributed)
func (_PartyBid *PartyBidFilterer) WatchFinalized(opts *bind.WatchOpts, sink chan<- *PartyBidFinalized) (event.Subscription, error) {

	logs, sub, err := _PartyBid.contract.WatchLogs(opts, "Finalized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PartyBidFinalized)
				if err := _PartyBid.contract.UnpackLog(event, "Finalized", log); err != nil {
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

// ParseFinalized is a log parse operation binding the contract event 0x9a2087478f16b801ecd568a6676f5db758bda2a01b954b2c754257d11eb3770b.
//
// Solidity: event Finalized(uint8 result, uint256 totalSpent, uint256 fee, uint256 totalContributed)
func (_PartyBid *PartyBidFilterer) ParseFinalized(log types.Log) (*PartyBidFinalized, error) {
	event := new(PartyBidFinalized)
	if err := _PartyBid.contract.UnpackLog(event, "Finalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
