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

// ArtistV5MetaData contains all meta data concerning the ArtistV5 contract.
var ArtistV5MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumArtistV5.TimeType\",\"name\":\"timeType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"editionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"newTime\",\"type\":\"uint32\"}],\"name\":\"AuctionTimeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"editionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"baseURI\",\"type\":\"string\"}],\"name\":\"BaseURISet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"editionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"fundingRecipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"quantity\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"royaltyBPS\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"startTime\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"endTime\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"permissionedQuantity\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signerAddress\",\"type\":\"address\"}],\"name\":\"EditionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"editionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"numSold\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ticketNumber\",\"type\":\"uint256\"}],\"name\":\"EditionPurchased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"editionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"permissionedQuantity\",\"type\":\"uint32\"}],\"name\":\"PermissionedQuantitySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"editionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signerAddress\",\"type\":\"address\"}],\"name\":\"SignerAddressSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMISSIONED_SALE_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"atEditionId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"atTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_editionId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_ticketNumber\",\"type\":\"uint256\"}],\"name\":\"buyEdition\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_editionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_ticketNumbers\",\"type\":\"uint256[]\"}],\"name\":\"checkTicketNumbers\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_fundingRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_quantity\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_royaltyBPS\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_startTime\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_endTime\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_permissionedQuantity\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_signerAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_editionId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_baseURI\",\"type\":\"string\"}],\"name\":\"createEdition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"depositedForEdition\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"editionCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"editions\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"fundingRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"numSold\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"quantity\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"royaltyBPS\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"startTime\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"endTime\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"permissionedQuantity\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"signerAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"baseURI\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_baseURI\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"ownersOfTokenIds\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_salePrice\",\"type\":\"uint256\"}],\"name\":\"royaltyInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"fundingRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"royaltyAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_editionId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_baseURI\",\"type\":\"string\"}],\"name\":\"setEditionBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_editionId\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_endTime\",\"type\":\"uint32\"}],\"name\":\"setEndTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"setOwnerOverride\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_editionId\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_permissionedQuantity\",\"type\":\"uint32\"}],\"name\":\"setPermissionedQuantity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_editionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_newSignerAddress\",\"type\":\"address\"}],\"name\":\"setSignerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_editionId\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_startTime\",\"type\":\"uint32\"}],\"name\":\"setStartTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"soundRecoveryAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenToEdition\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_editionId\",\"type\":\"uint256\"}],\"name\":\"withdrawFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"withdrawnForEdition\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ArtistV5ABI is the input ABI used to generate the binding from.
// Deprecated: Use ArtistV5MetaData.ABI instead.
var ArtistV5ABI = ArtistV5MetaData.ABI

// ArtistV5 is an auto generated Go binding around an Ethereum contract.
type ArtistV5 struct {
	ArtistV5Caller     // Read-only binding to the contract
	ArtistV5Transactor // Write-only binding to the contract
	ArtistV5Filterer   // Log filterer for contract events
}

// ArtistV5Caller is an auto generated read-only Go binding around an Ethereum contract.
type ArtistV5Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArtistV5Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ArtistV5Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArtistV5Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArtistV5Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArtistV5Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArtistV5Session struct {
	Contract     *ArtistV5         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArtistV5CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArtistV5CallerSession struct {
	Contract *ArtistV5Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ArtistV5TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArtistV5TransactorSession struct {
	Contract     *ArtistV5Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ArtistV5Raw is an auto generated low-level Go binding around an Ethereum contract.
type ArtistV5Raw struct {
	Contract *ArtistV5 // Generic contract binding to access the raw methods on
}

// ArtistV5CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArtistV5CallerRaw struct {
	Contract *ArtistV5Caller // Generic read-only contract binding to access the raw methods on
}

// ArtistV5TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArtistV5TransactorRaw struct {
	Contract *ArtistV5Transactor // Generic write-only contract binding to access the raw methods on
}

// NewArtistV5 creates a new instance of ArtistV5, bound to a specific deployed contract.
func NewArtistV5(address common.Address, backend bind.ContractBackend) (*ArtistV5, error) {
	contract, err := bindArtistV5(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ArtistV5{ArtistV5Caller: ArtistV5Caller{contract: contract}, ArtistV5Transactor: ArtistV5Transactor{contract: contract}, ArtistV5Filterer: ArtistV5Filterer{contract: contract}}, nil
}

// NewArtistV5Caller creates a new read-only instance of ArtistV5, bound to a specific deployed contract.
func NewArtistV5Caller(address common.Address, caller bind.ContractCaller) (*ArtistV5Caller, error) {
	contract, err := bindArtistV5(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArtistV5Caller{contract: contract}, nil
}

// NewArtistV5Transactor creates a new write-only instance of ArtistV5, bound to a specific deployed contract.
func NewArtistV5Transactor(address common.Address, transactor bind.ContractTransactor) (*ArtistV5Transactor, error) {
	contract, err := bindArtistV5(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArtistV5Transactor{contract: contract}, nil
}

// NewArtistV5Filterer creates a new log filterer instance of ArtistV5, bound to a specific deployed contract.
func NewArtistV5Filterer(address common.Address, filterer bind.ContractFilterer) (*ArtistV5Filterer, error) {
	contract, err := bindArtistV5(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArtistV5Filterer{contract: contract}, nil
}

// bindArtistV5 binds a generic wrapper to an already deployed contract.
func bindArtistV5(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ArtistV5ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArtistV5 *ArtistV5Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArtistV5.Contract.ArtistV5Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArtistV5 *ArtistV5Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArtistV5.Contract.ArtistV5Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArtistV5 *ArtistV5Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArtistV5.Contract.ArtistV5Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArtistV5 *ArtistV5CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArtistV5.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArtistV5 *ArtistV5TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArtistV5.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArtistV5 *ArtistV5TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArtistV5.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_ArtistV5 *ArtistV5Caller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ArtistV5.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_ArtistV5 *ArtistV5Session) ADMINROLE() ([32]byte, error) {
	return _ArtistV5.Contract.ADMINROLE(&_ArtistV5.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_ArtistV5 *ArtistV5CallerSession) ADMINROLE() ([32]byte, error) {
	return _ArtistV5.Contract.ADMINROLE(&_ArtistV5.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_ArtistV5 *ArtistV5Caller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ArtistV5.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_ArtistV5 *ArtistV5Session) DOMAINSEPARATOR() ([32]byte, error) {
	return _ArtistV5.Contract.DOMAINSEPARATOR(&_ArtistV5.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_ArtistV5 *ArtistV5CallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _ArtistV5.Contract.DOMAINSEPARATOR(&_ArtistV5.CallOpts)
}

// PERMISSIONEDSALETYPEHASH is a free data retrieval call binding the contract method 0x27399d36.
//
// Solidity: function PERMISSIONED_SALE_TYPEHASH() view returns(bytes32)
func (_ArtistV5 *ArtistV5Caller) PERMISSIONEDSALETYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ArtistV5.contract.Call(opts, &out, "PERMISSIONED_SALE_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PERMISSIONEDSALETYPEHASH is a free data retrieval call binding the contract method 0x27399d36.
//
// Solidity: function PERMISSIONED_SALE_TYPEHASH() view returns(bytes32)
func (_ArtistV5 *ArtistV5Session) PERMISSIONEDSALETYPEHASH() ([32]byte, error) {
	return _ArtistV5.Contract.PERMISSIONEDSALETYPEHASH(&_ArtistV5.CallOpts)
}

// PERMISSIONEDSALETYPEHASH is a free data retrieval call binding the contract method 0x27399d36.
//
// Solidity: function PERMISSIONED_SALE_TYPEHASH() view returns(bytes32)
func (_ArtistV5 *ArtistV5CallerSession) PERMISSIONEDSALETYPEHASH() ([32]byte, error) {
	return _ArtistV5.Contract.PERMISSIONEDSALETYPEHASH(&_ArtistV5.CallOpts)
}

// AtEditionId is a free data retrieval call binding the contract method 0x9725d92e.
//
// Solidity: function atEditionId() view returns(uint256 _value)
func (_ArtistV5 *ArtistV5Caller) AtEditionId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ArtistV5.contract.Call(opts, &out, "atEditionId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AtEditionId is a free data retrieval call binding the contract method 0x9725d92e.
//
// Solidity: function atEditionId() view returns(uint256 _value)
func (_ArtistV5 *ArtistV5Session) AtEditionId() (*big.Int, error) {
	return _ArtistV5.Contract.AtEditionId(&_ArtistV5.CallOpts)
}

// AtEditionId is a free data retrieval call binding the contract method 0x9725d92e.
//
// Solidity: function atEditionId() view returns(uint256 _value)
func (_ArtistV5 *ArtistV5CallerSession) AtEditionId() (*big.Int, error) {
	return _ArtistV5.Contract.AtEditionId(&_ArtistV5.CallOpts)
}

// AtTokenId is a free data retrieval call binding the contract method 0x3ef2dbc2.
//
// Solidity: function atTokenId() view returns(uint256 _value)
func (_ArtistV5 *ArtistV5Caller) AtTokenId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ArtistV5.contract.Call(opts, &out, "atTokenId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AtTokenId is a free data retrieval call binding the contract method 0x3ef2dbc2.
//
// Solidity: function atTokenId() view returns(uint256 _value)
func (_ArtistV5 *ArtistV5Session) AtTokenId() (*big.Int, error) {
	return _ArtistV5.Contract.AtTokenId(&_ArtistV5.CallOpts)
}

// AtTokenId is a free data retrieval call binding the contract method 0x3ef2dbc2.
//
// Solidity: function atTokenId() view returns(uint256 _value)
func (_ArtistV5 *ArtistV5CallerSession) AtTokenId() (*big.Int, error) {
	return _ArtistV5.Contract.AtTokenId(&_ArtistV5.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ArtistV5 *ArtistV5Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ArtistV5.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ArtistV5 *ArtistV5Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ArtistV5.Contract.BalanceOf(&_ArtistV5.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ArtistV5 *ArtistV5CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ArtistV5.Contract.BalanceOf(&_ArtistV5.CallOpts, owner)
}

// CheckTicketNumbers is a free data retrieval call binding the contract method 0x065d5b85.
//
// Solidity: function checkTicketNumbers(uint256 _editionId, uint256[] _ticketNumbers) view returns(bool[])
func (_ArtistV5 *ArtistV5Caller) CheckTicketNumbers(opts *bind.CallOpts, _editionId *big.Int, _ticketNumbers []*big.Int) ([]bool, error) {
	var out []interface{}
	err := _ArtistV5.contract.Call(opts, &out, "checkTicketNumbers", _editionId, _ticketNumbers)

	if err != nil {
		return *new([]bool), err
	}

	out0 := *abi.ConvertType(out[0], new([]bool)).(*[]bool)

	return out0, err

}

// CheckTicketNumbers is a free data retrieval call binding the contract method 0x065d5b85.
//
// Solidity: function checkTicketNumbers(uint256 _editionId, uint256[] _ticketNumbers) view returns(bool[])
func (_ArtistV5 *ArtistV5Session) CheckTicketNumbers(_editionId *big.Int, _ticketNumbers []*big.Int) ([]bool, error) {
	return _ArtistV5.Contract.CheckTicketNumbers(&_ArtistV5.CallOpts, _editionId, _ticketNumbers)
}

// CheckTicketNumbers is a free data retrieval call binding the contract method 0x065d5b85.
//
// Solidity: function checkTicketNumbers(uint256 _editionId, uint256[] _ticketNumbers) view returns(bool[])
func (_ArtistV5 *ArtistV5CallerSession) CheckTicketNumbers(_editionId *big.Int, _ticketNumbers []*big.Int) ([]bool, error) {
	return _ArtistV5.Contract.CheckTicketNumbers(&_ArtistV5.CallOpts, _editionId, _ticketNumbers)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_ArtistV5 *ArtistV5Caller) ContractURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ArtistV5.contract.Call(opts, &out, "contractURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_ArtistV5 *ArtistV5Session) ContractURI() (string, error) {
	return _ArtistV5.Contract.ContractURI(&_ArtistV5.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_ArtistV5 *ArtistV5CallerSession) ContractURI() (string, error) {
	return _ArtistV5.Contract.ContractURI(&_ArtistV5.CallOpts)
}

// DepositedForEdition is a free data retrieval call binding the contract method 0xe1a3d573.
//
// Solidity: function depositedForEdition(uint256 ) view returns(uint256)
func (_ArtistV5 *ArtistV5Caller) DepositedForEdition(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ArtistV5.contract.Call(opts, &out, "depositedForEdition", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositedForEdition is a free data retrieval call binding the contract method 0xe1a3d573.
//
// Solidity: function depositedForEdition(uint256 ) view returns(uint256)
func (_ArtistV5 *ArtistV5Session) DepositedForEdition(arg0 *big.Int) (*big.Int, error) {
	return _ArtistV5.Contract.DepositedForEdition(&_ArtistV5.CallOpts, arg0)
}

// DepositedForEdition is a free data retrieval call binding the contract method 0xe1a3d573.
//
// Solidity: function depositedForEdition(uint256 ) view returns(uint256)
func (_ArtistV5 *ArtistV5CallerSession) DepositedForEdition(arg0 *big.Int) (*big.Int, error) {
	return _ArtistV5.Contract.DepositedForEdition(&_ArtistV5.CallOpts, arg0)
}

// EditionCount is a free data retrieval call binding the contract method 0x4bf44026.
//
// Solidity: function editionCount() view returns(uint256)
func (_ArtistV5 *ArtistV5Caller) EditionCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ArtistV5.contract.Call(opts, &out, "editionCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EditionCount is a free data retrieval call binding the contract method 0x4bf44026.
//
// Solidity: function editionCount() view returns(uint256)
func (_ArtistV5 *ArtistV5Session) EditionCount() (*big.Int, error) {
	return _ArtistV5.Contract.EditionCount(&_ArtistV5.CallOpts)
}

// EditionCount is a free data retrieval call binding the contract method 0x4bf44026.
//
// Solidity: function editionCount() view returns(uint256)
func (_ArtistV5 *ArtistV5CallerSession) EditionCount() (*big.Int, error) {
	return _ArtistV5.Contract.EditionCount(&_ArtistV5.CallOpts)
}

// Editions is a free data retrieval call binding the contract method 0x279c806e.
//
// Solidity: function editions(uint256 ) view returns(address fundingRecipient, uint256 price, uint32 numSold, uint32 quantity, uint32 royaltyBPS, uint32 startTime, uint32 endTime, uint32 permissionedQuantity, address signerAddress, string baseURI)
func (_ArtistV5 *ArtistV5Caller) Editions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	FundingRecipient     common.Address
	Price                *big.Int
	NumSold              uint32
	Quantity             uint32
	RoyaltyBPS           uint32
	StartTime            uint32
	EndTime              uint32
	PermissionedQuantity uint32
	SignerAddress        common.Address
	BaseURI              string
}, error) {
	var out []interface{}
	err := _ArtistV5.contract.Call(opts, &out, "editions", arg0)

	outstruct := new(struct {
		FundingRecipient     common.Address
		Price                *big.Int
		NumSold              uint32
		Quantity             uint32
		RoyaltyBPS           uint32
		StartTime            uint32
		EndTime              uint32
		PermissionedQuantity uint32
		SignerAddress        common.Address
		BaseURI              string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FundingRecipient = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Price = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.NumSold = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.Quantity = *abi.ConvertType(out[3], new(uint32)).(*uint32)
	outstruct.RoyaltyBPS = *abi.ConvertType(out[4], new(uint32)).(*uint32)
	outstruct.StartTime = *abi.ConvertType(out[5], new(uint32)).(*uint32)
	outstruct.EndTime = *abi.ConvertType(out[6], new(uint32)).(*uint32)
	outstruct.PermissionedQuantity = *abi.ConvertType(out[7], new(uint32)).(*uint32)
	outstruct.SignerAddress = *abi.ConvertType(out[8], new(common.Address)).(*common.Address)
	outstruct.BaseURI = *abi.ConvertType(out[9], new(string)).(*string)

	return *outstruct, err

}

// Editions is a free data retrieval call binding the contract method 0x279c806e.
//
// Solidity: function editions(uint256 ) view returns(address fundingRecipient, uint256 price, uint32 numSold, uint32 quantity, uint32 royaltyBPS, uint32 startTime, uint32 endTime, uint32 permissionedQuantity, address signerAddress, string baseURI)
func (_ArtistV5 *ArtistV5Session) Editions(arg0 *big.Int) (struct {
	FundingRecipient     common.Address
	Price                *big.Int
	NumSold              uint32
	Quantity             uint32
	RoyaltyBPS           uint32
	StartTime            uint32
	EndTime              uint32
	PermissionedQuantity uint32
	SignerAddress        common.Address
	BaseURI              string
}, error) {
	return _ArtistV5.Contract.Editions(&_ArtistV5.CallOpts, arg0)
}

// Editions is a free data retrieval call binding the contract method 0x279c806e.
//
// Solidity: function editions(uint256 ) view returns(address fundingRecipient, uint256 price, uint32 numSold, uint32 quantity, uint32 royaltyBPS, uint32 startTime, uint32 endTime, uint32 permissionedQuantity, address signerAddress, string baseURI)
func (_ArtistV5 *ArtistV5CallerSession) Editions(arg0 *big.Int) (struct {
	FundingRecipient     common.Address
	Price                *big.Int
	NumSold              uint32
	Quantity             uint32
	RoyaltyBPS           uint32
	StartTime            uint32
	EndTime              uint32
	PermissionedQuantity uint32
	SignerAddress        common.Address
	BaseURI              string
}, error) {
	return _ArtistV5.Contract.Editions(&_ArtistV5.CallOpts, arg0)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ArtistV5 *ArtistV5Caller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ArtistV5.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ArtistV5 *ArtistV5Session) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ArtistV5.Contract.GetApproved(&_ArtistV5.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ArtistV5 *ArtistV5CallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ArtistV5.Contract.GetApproved(&_ArtistV5.CallOpts, tokenId)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ArtistV5 *ArtistV5Caller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _ArtistV5.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ArtistV5 *ArtistV5Session) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ArtistV5.Contract.HasRole(&_ArtistV5.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ArtistV5 *ArtistV5CallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ArtistV5.Contract.HasRole(&_ArtistV5.CallOpts, role, account)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ArtistV5 *ArtistV5Caller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _ArtistV5.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ArtistV5 *ArtistV5Session) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ArtistV5.Contract.IsApprovedForAll(&_ArtistV5.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ArtistV5 *ArtistV5CallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ArtistV5.Contract.IsApprovedForAll(&_ArtistV5.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ArtistV5 *ArtistV5Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ArtistV5.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ArtistV5 *ArtistV5Session) Name() (string, error) {
	return _ArtistV5.Contract.Name(&_ArtistV5.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ArtistV5 *ArtistV5CallerSession) Name() (string, error) {
	return _ArtistV5.Contract.Name(&_ArtistV5.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ArtistV5 *ArtistV5Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ArtistV5.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ArtistV5 *ArtistV5Session) Owner() (common.Address, error) {
	return _ArtistV5.Contract.Owner(&_ArtistV5.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ArtistV5 *ArtistV5CallerSession) Owner() (common.Address, error) {
	return _ArtistV5.Contract.Owner(&_ArtistV5.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ArtistV5 *ArtistV5Caller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ArtistV5.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ArtistV5 *ArtistV5Session) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ArtistV5.Contract.OwnerOf(&_ArtistV5.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ArtistV5 *ArtistV5CallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ArtistV5.Contract.OwnerOf(&_ArtistV5.CallOpts, tokenId)
}

// OwnersOfTokenIds is a free data retrieval call binding the contract method 0x52f5c2e4.
//
// Solidity: function ownersOfTokenIds(uint256[] _tokenIds) view returns(address[])
func (_ArtistV5 *ArtistV5Caller) OwnersOfTokenIds(opts *bind.CallOpts, _tokenIds []*big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _ArtistV5.contract.Call(opts, &out, "ownersOfTokenIds", _tokenIds)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// OwnersOfTokenIds is a free data retrieval call binding the contract method 0x52f5c2e4.
//
// Solidity: function ownersOfTokenIds(uint256[] _tokenIds) view returns(address[])
func (_ArtistV5 *ArtistV5Session) OwnersOfTokenIds(_tokenIds []*big.Int) ([]common.Address, error) {
	return _ArtistV5.Contract.OwnersOfTokenIds(&_ArtistV5.CallOpts, _tokenIds)
}

// OwnersOfTokenIds is a free data retrieval call binding the contract method 0x52f5c2e4.
//
// Solidity: function ownersOfTokenIds(uint256[] _tokenIds) view returns(address[])
func (_ArtistV5 *ArtistV5CallerSession) OwnersOfTokenIds(_tokenIds []*big.Int) ([]common.Address, error) {
	return _ArtistV5.Contract.OwnersOfTokenIds(&_ArtistV5.CallOpts, _tokenIds)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 _tokenId, uint256 _salePrice) view returns(address fundingRecipient, uint256 royaltyAmount)
func (_ArtistV5 *ArtistV5Caller) RoyaltyInfo(opts *bind.CallOpts, _tokenId *big.Int, _salePrice *big.Int) (struct {
	FundingRecipient common.Address
	RoyaltyAmount    *big.Int
}, error) {
	var out []interface{}
	err := _ArtistV5.contract.Call(opts, &out, "royaltyInfo", _tokenId, _salePrice)

	outstruct := new(struct {
		FundingRecipient common.Address
		RoyaltyAmount    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FundingRecipient = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.RoyaltyAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 _tokenId, uint256 _salePrice) view returns(address fundingRecipient, uint256 royaltyAmount)
func (_ArtistV5 *ArtistV5Session) RoyaltyInfo(_tokenId *big.Int, _salePrice *big.Int) (struct {
	FundingRecipient common.Address
	RoyaltyAmount    *big.Int
}, error) {
	return _ArtistV5.Contract.RoyaltyInfo(&_ArtistV5.CallOpts, _tokenId, _salePrice)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 _tokenId, uint256 _salePrice) view returns(address fundingRecipient, uint256 royaltyAmount)
func (_ArtistV5 *ArtistV5CallerSession) RoyaltyInfo(_tokenId *big.Int, _salePrice *big.Int) (struct {
	FundingRecipient common.Address
	RoyaltyAmount    *big.Int
}, error) {
	return _ArtistV5.Contract.RoyaltyInfo(&_ArtistV5.CallOpts, _tokenId, _salePrice)
}

// SoundRecoveryAddress is a free data retrieval call binding the contract method 0x0bcca831.
//
// Solidity: function soundRecoveryAddress() view returns(address)
func (_ArtistV5 *ArtistV5Caller) SoundRecoveryAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ArtistV5.contract.Call(opts, &out, "soundRecoveryAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SoundRecoveryAddress is a free data retrieval call binding the contract method 0x0bcca831.
//
// Solidity: function soundRecoveryAddress() view returns(address)
func (_ArtistV5 *ArtistV5Session) SoundRecoveryAddress() (common.Address, error) {
	return _ArtistV5.Contract.SoundRecoveryAddress(&_ArtistV5.CallOpts)
}

// SoundRecoveryAddress is a free data retrieval call binding the contract method 0x0bcca831.
//
// Solidity: function soundRecoveryAddress() view returns(address)
func (_ArtistV5 *ArtistV5CallerSession) SoundRecoveryAddress() (common.Address, error) {
	return _ArtistV5.Contract.SoundRecoveryAddress(&_ArtistV5.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (_ArtistV5 *ArtistV5Caller) SupportsInterface(opts *bind.CallOpts, _interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ArtistV5.contract.Call(opts, &out, "supportsInterface", _interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (_ArtistV5 *ArtistV5Session) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _ArtistV5.Contract.SupportsInterface(&_ArtistV5.CallOpts, _interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (_ArtistV5 *ArtistV5CallerSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _ArtistV5.Contract.SupportsInterface(&_ArtistV5.CallOpts, _interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ArtistV5 *ArtistV5Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ArtistV5.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ArtistV5 *ArtistV5Session) Symbol() (string, error) {
	return _ArtistV5.Contract.Symbol(&_ArtistV5.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ArtistV5 *ArtistV5CallerSession) Symbol() (string, error) {
	return _ArtistV5.Contract.Symbol(&_ArtistV5.CallOpts)
}

// TokenToEdition is a free data retrieval call binding the contract method 0x602787ed.
//
// Solidity: function tokenToEdition(uint256 _tokenId) view returns(uint256)
func (_ArtistV5 *ArtistV5Caller) TokenToEdition(opts *bind.CallOpts, _tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ArtistV5.contract.Call(opts, &out, "tokenToEdition", _tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenToEdition is a free data retrieval call binding the contract method 0x602787ed.
//
// Solidity: function tokenToEdition(uint256 _tokenId) view returns(uint256)
func (_ArtistV5 *ArtistV5Session) TokenToEdition(_tokenId *big.Int) (*big.Int, error) {
	return _ArtistV5.Contract.TokenToEdition(&_ArtistV5.CallOpts, _tokenId)
}

// TokenToEdition is a free data retrieval call binding the contract method 0x602787ed.
//
// Solidity: function tokenToEdition(uint256 _tokenId) view returns(uint256)
func (_ArtistV5 *ArtistV5CallerSession) TokenToEdition(_tokenId *big.Int) (*big.Int, error) {
	return _ArtistV5.Contract.TokenToEdition(&_ArtistV5.CallOpts, _tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_ArtistV5 *ArtistV5Caller) TokenURI(opts *bind.CallOpts, _tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _ArtistV5.contract.Call(opts, &out, "tokenURI", _tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_ArtistV5 *ArtistV5Session) TokenURI(_tokenId *big.Int) (string, error) {
	return _ArtistV5.Contract.TokenURI(&_ArtistV5.CallOpts, _tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_ArtistV5 *ArtistV5CallerSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _ArtistV5.Contract.TokenURI(&_ArtistV5.CallOpts, _tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ArtistV5 *ArtistV5Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ArtistV5.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ArtistV5 *ArtistV5Session) TotalSupply() (*big.Int, error) {
	return _ArtistV5.Contract.TotalSupply(&_ArtistV5.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ArtistV5 *ArtistV5CallerSession) TotalSupply() (*big.Int, error) {
	return _ArtistV5.Contract.TotalSupply(&_ArtistV5.CallOpts)
}

// WithdrawnForEdition is a free data retrieval call binding the contract method 0xd3bb0528.
//
// Solidity: function withdrawnForEdition(uint256 ) view returns(uint256)
func (_ArtistV5 *ArtistV5Caller) WithdrawnForEdition(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ArtistV5.contract.Call(opts, &out, "withdrawnForEdition", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawnForEdition is a free data retrieval call binding the contract method 0xd3bb0528.
//
// Solidity: function withdrawnForEdition(uint256 ) view returns(uint256)
func (_ArtistV5 *ArtistV5Session) WithdrawnForEdition(arg0 *big.Int) (*big.Int, error) {
	return _ArtistV5.Contract.WithdrawnForEdition(&_ArtistV5.CallOpts, arg0)
}

// WithdrawnForEdition is a free data retrieval call binding the contract method 0xd3bb0528.
//
// Solidity: function withdrawnForEdition(uint256 ) view returns(uint256)
func (_ArtistV5 *ArtistV5CallerSession) WithdrawnForEdition(arg0 *big.Int) (*big.Int, error) {
	return _ArtistV5.Contract.WithdrawnForEdition(&_ArtistV5.CallOpts, arg0)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ArtistV5 *ArtistV5Transactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ArtistV5.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ArtistV5 *ArtistV5Session) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ArtistV5.Contract.Approve(&_ArtistV5.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ArtistV5 *ArtistV5TransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ArtistV5.Contract.Approve(&_ArtistV5.TransactOpts, to, tokenId)
}

// BuyEdition is a paid mutator transaction binding the contract method 0xf71e54fb.
//
// Solidity: function buyEdition(uint256 _editionId, bytes _signature, uint256 _ticketNumber) payable returns()
func (_ArtistV5 *ArtistV5Transactor) BuyEdition(opts *bind.TransactOpts, _editionId *big.Int, _signature []byte, _ticketNumber *big.Int) (*types.Transaction, error) {
	return _ArtistV5.contract.Transact(opts, "buyEdition", _editionId, _signature, _ticketNumber)
}

// BuyEdition is a paid mutator transaction binding the contract method 0xf71e54fb.
//
// Solidity: function buyEdition(uint256 _editionId, bytes _signature, uint256 _ticketNumber) payable returns()
func (_ArtistV5 *ArtistV5Session) BuyEdition(_editionId *big.Int, _signature []byte, _ticketNumber *big.Int) (*types.Transaction, error) {
	return _ArtistV5.Contract.BuyEdition(&_ArtistV5.TransactOpts, _editionId, _signature, _ticketNumber)
}

// BuyEdition is a paid mutator transaction binding the contract method 0xf71e54fb.
//
// Solidity: function buyEdition(uint256 _editionId, bytes _signature, uint256 _ticketNumber) payable returns()
func (_ArtistV5 *ArtistV5TransactorSession) BuyEdition(_editionId *big.Int, _signature []byte, _ticketNumber *big.Int) (*types.Transaction, error) {
	return _ArtistV5.Contract.BuyEdition(&_ArtistV5.TransactOpts, _editionId, _signature, _ticketNumber)
}

// CreateEdition is a paid mutator transaction binding the contract method 0x8e116aea.
//
// Solidity: function createEdition(address _fundingRecipient, uint256 _price, uint32 _quantity, uint32 _royaltyBPS, uint32 _startTime, uint32 _endTime, uint32 _permissionedQuantity, address _signerAddress, uint256 _editionId, string _baseURI) returns()
func (_ArtistV5 *ArtistV5Transactor) CreateEdition(opts *bind.TransactOpts, _fundingRecipient common.Address, _price *big.Int, _quantity uint32, _royaltyBPS uint32, _startTime uint32, _endTime uint32, _permissionedQuantity uint32, _signerAddress common.Address, _editionId *big.Int, _baseURI string) (*types.Transaction, error) {
	return _ArtistV5.contract.Transact(opts, "createEdition", _fundingRecipient, _price, _quantity, _royaltyBPS, _startTime, _endTime, _permissionedQuantity, _signerAddress, _editionId, _baseURI)
}

// CreateEdition is a paid mutator transaction binding the contract method 0x8e116aea.
//
// Solidity: function createEdition(address _fundingRecipient, uint256 _price, uint32 _quantity, uint32 _royaltyBPS, uint32 _startTime, uint32 _endTime, uint32 _permissionedQuantity, address _signerAddress, uint256 _editionId, string _baseURI) returns()
func (_ArtistV5 *ArtistV5Session) CreateEdition(_fundingRecipient common.Address, _price *big.Int, _quantity uint32, _royaltyBPS uint32, _startTime uint32, _endTime uint32, _permissionedQuantity uint32, _signerAddress common.Address, _editionId *big.Int, _baseURI string) (*types.Transaction, error) {
	return _ArtistV5.Contract.CreateEdition(&_ArtistV5.TransactOpts, _fundingRecipient, _price, _quantity, _royaltyBPS, _startTime, _endTime, _permissionedQuantity, _signerAddress, _editionId, _baseURI)
}

// CreateEdition is a paid mutator transaction binding the contract method 0x8e116aea.
//
// Solidity: function createEdition(address _fundingRecipient, uint256 _price, uint32 _quantity, uint32 _royaltyBPS, uint32 _startTime, uint32 _endTime, uint32 _permissionedQuantity, address _signerAddress, uint256 _editionId, string _baseURI) returns()
func (_ArtistV5 *ArtistV5TransactorSession) CreateEdition(_fundingRecipient common.Address, _price *big.Int, _quantity uint32, _royaltyBPS uint32, _startTime uint32, _endTime uint32, _permissionedQuantity uint32, _signerAddress common.Address, _editionId *big.Int, _baseURI string) (*types.Transaction, error) {
	return _ArtistV5.Contract.CreateEdition(&_ArtistV5.TransactOpts, _fundingRecipient, _price, _quantity, _royaltyBPS, _startTime, _endTime, _permissionedQuantity, _signerAddress, _editionId, _baseURI)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ArtistV5 *ArtistV5Transactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ArtistV5.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ArtistV5 *ArtistV5Session) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ArtistV5.Contract.GrantRole(&_ArtistV5.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ArtistV5 *ArtistV5TransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ArtistV5.Contract.GrantRole(&_ArtistV5.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x5f1e6f6d.
//
// Solidity: function initialize(address _owner, string _name, string _symbol, string _baseURI) returns()
func (_ArtistV5 *ArtistV5Transactor) Initialize(opts *bind.TransactOpts, _owner common.Address, _name string, _symbol string, _baseURI string) (*types.Transaction, error) {
	return _ArtistV5.contract.Transact(opts, "initialize", _owner, _name, _symbol, _baseURI)
}

// Initialize is a paid mutator transaction binding the contract method 0x5f1e6f6d.
//
// Solidity: function initialize(address _owner, string _name, string _symbol, string _baseURI) returns()
func (_ArtistV5 *ArtistV5Session) Initialize(_owner common.Address, _name string, _symbol string, _baseURI string) (*types.Transaction, error) {
	return _ArtistV5.Contract.Initialize(&_ArtistV5.TransactOpts, _owner, _name, _symbol, _baseURI)
}

// Initialize is a paid mutator transaction binding the contract method 0x5f1e6f6d.
//
// Solidity: function initialize(address _owner, string _name, string _symbol, string _baseURI) returns()
func (_ArtistV5 *ArtistV5TransactorSession) Initialize(_owner common.Address, _name string, _symbol string, _baseURI string) (*types.Transaction, error) {
	return _ArtistV5.Contract.Initialize(&_ArtistV5.TransactOpts, _owner, _name, _symbol, _baseURI)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ArtistV5 *ArtistV5Transactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ArtistV5.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ArtistV5 *ArtistV5Session) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ArtistV5.Contract.RevokeRole(&_ArtistV5.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ArtistV5 *ArtistV5TransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ArtistV5.Contract.RevokeRole(&_ArtistV5.TransactOpts, role, account)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ArtistV5 *ArtistV5Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ArtistV5.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ArtistV5 *ArtistV5Session) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ArtistV5.Contract.SafeTransferFrom(&_ArtistV5.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ArtistV5 *ArtistV5TransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ArtistV5.Contract.SafeTransferFrom(&_ArtistV5.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_ArtistV5 *ArtistV5Transactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ArtistV5.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_ArtistV5 *ArtistV5Session) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ArtistV5.Contract.SafeTransferFrom0(&_ArtistV5.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_ArtistV5 *ArtistV5TransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ArtistV5.Contract.SafeTransferFrom0(&_ArtistV5.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ArtistV5 *ArtistV5Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _ArtistV5.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ArtistV5 *ArtistV5Session) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ArtistV5.Contract.SetApprovalForAll(&_ArtistV5.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ArtistV5 *ArtistV5TransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ArtistV5.Contract.SetApprovalForAll(&_ArtistV5.TransactOpts, operator, approved)
}

// SetEditionBaseURI is a paid mutator transaction binding the contract method 0x79672692.
//
// Solidity: function setEditionBaseURI(uint256 _editionId, string _baseURI) returns()
func (_ArtistV5 *ArtistV5Transactor) SetEditionBaseURI(opts *bind.TransactOpts, _editionId *big.Int, _baseURI string) (*types.Transaction, error) {
	return _ArtistV5.contract.Transact(opts, "setEditionBaseURI", _editionId, _baseURI)
}

// SetEditionBaseURI is a paid mutator transaction binding the contract method 0x79672692.
//
// Solidity: function setEditionBaseURI(uint256 _editionId, string _baseURI) returns()
func (_ArtistV5 *ArtistV5Session) SetEditionBaseURI(_editionId *big.Int, _baseURI string) (*types.Transaction, error) {
	return _ArtistV5.Contract.SetEditionBaseURI(&_ArtistV5.TransactOpts, _editionId, _baseURI)
}

// SetEditionBaseURI is a paid mutator transaction binding the contract method 0x79672692.
//
// Solidity: function setEditionBaseURI(uint256 _editionId, string _baseURI) returns()
func (_ArtistV5 *ArtistV5TransactorSession) SetEditionBaseURI(_editionId *big.Int, _baseURI string) (*types.Transaction, error) {
	return _ArtistV5.Contract.SetEditionBaseURI(&_ArtistV5.TransactOpts, _editionId, _baseURI)
}

// SetEndTime is a paid mutator transaction binding the contract method 0xbb314ca1.
//
// Solidity: function setEndTime(uint256 _editionId, uint32 _endTime) returns()
func (_ArtistV5 *ArtistV5Transactor) SetEndTime(opts *bind.TransactOpts, _editionId *big.Int, _endTime uint32) (*types.Transaction, error) {
	return _ArtistV5.contract.Transact(opts, "setEndTime", _editionId, _endTime)
}

// SetEndTime is a paid mutator transaction binding the contract method 0xbb314ca1.
//
// Solidity: function setEndTime(uint256 _editionId, uint32 _endTime) returns()
func (_ArtistV5 *ArtistV5Session) SetEndTime(_editionId *big.Int, _endTime uint32) (*types.Transaction, error) {
	return _ArtistV5.Contract.SetEndTime(&_ArtistV5.TransactOpts, _editionId, _endTime)
}

// SetEndTime is a paid mutator transaction binding the contract method 0xbb314ca1.
//
// Solidity: function setEndTime(uint256 _editionId, uint32 _endTime) returns()
func (_ArtistV5 *ArtistV5TransactorSession) SetEndTime(_editionId *big.Int, _endTime uint32) (*types.Transaction, error) {
	return _ArtistV5.Contract.SetEndTime(&_ArtistV5.TransactOpts, _editionId, _endTime)
}

// SetOwnerOverride is a paid mutator transaction binding the contract method 0x75a8f08f.
//
// Solidity: function setOwnerOverride(address _newOwner) returns()
func (_ArtistV5 *ArtistV5Transactor) SetOwnerOverride(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _ArtistV5.contract.Transact(opts, "setOwnerOverride", _newOwner)
}

// SetOwnerOverride is a paid mutator transaction binding the contract method 0x75a8f08f.
//
// Solidity: function setOwnerOverride(address _newOwner) returns()
func (_ArtistV5 *ArtistV5Session) SetOwnerOverride(_newOwner common.Address) (*types.Transaction, error) {
	return _ArtistV5.Contract.SetOwnerOverride(&_ArtistV5.TransactOpts, _newOwner)
}

// SetOwnerOverride is a paid mutator transaction binding the contract method 0x75a8f08f.
//
// Solidity: function setOwnerOverride(address _newOwner) returns()
func (_ArtistV5 *ArtistV5TransactorSession) SetOwnerOverride(_newOwner common.Address) (*types.Transaction, error) {
	return _ArtistV5.Contract.SetOwnerOverride(&_ArtistV5.TransactOpts, _newOwner)
}

// SetPermissionedQuantity is a paid mutator transaction binding the contract method 0x52e25bf2.
//
// Solidity: function setPermissionedQuantity(uint256 _editionId, uint32 _permissionedQuantity) returns()
func (_ArtistV5 *ArtistV5Transactor) SetPermissionedQuantity(opts *bind.TransactOpts, _editionId *big.Int, _permissionedQuantity uint32) (*types.Transaction, error) {
	return _ArtistV5.contract.Transact(opts, "setPermissionedQuantity", _editionId, _permissionedQuantity)
}

// SetPermissionedQuantity is a paid mutator transaction binding the contract method 0x52e25bf2.
//
// Solidity: function setPermissionedQuantity(uint256 _editionId, uint32 _permissionedQuantity) returns()
func (_ArtistV5 *ArtistV5Session) SetPermissionedQuantity(_editionId *big.Int, _permissionedQuantity uint32) (*types.Transaction, error) {
	return _ArtistV5.Contract.SetPermissionedQuantity(&_ArtistV5.TransactOpts, _editionId, _permissionedQuantity)
}

// SetPermissionedQuantity is a paid mutator transaction binding the contract method 0x52e25bf2.
//
// Solidity: function setPermissionedQuantity(uint256 _editionId, uint32 _permissionedQuantity) returns()
func (_ArtistV5 *ArtistV5TransactorSession) SetPermissionedQuantity(_editionId *big.Int, _permissionedQuantity uint32) (*types.Transaction, error) {
	return _ArtistV5.Contract.SetPermissionedQuantity(&_ArtistV5.TransactOpts, _editionId, _permissionedQuantity)
}

// SetSignerAddress is a paid mutator transaction binding the contract method 0x56dee996.
//
// Solidity: function setSignerAddress(uint256 _editionId, address _newSignerAddress) returns()
func (_ArtistV5 *ArtistV5Transactor) SetSignerAddress(opts *bind.TransactOpts, _editionId *big.Int, _newSignerAddress common.Address) (*types.Transaction, error) {
	return _ArtistV5.contract.Transact(opts, "setSignerAddress", _editionId, _newSignerAddress)
}

// SetSignerAddress is a paid mutator transaction binding the contract method 0x56dee996.
//
// Solidity: function setSignerAddress(uint256 _editionId, address _newSignerAddress) returns()
func (_ArtistV5 *ArtistV5Session) SetSignerAddress(_editionId *big.Int, _newSignerAddress common.Address) (*types.Transaction, error) {
	return _ArtistV5.Contract.SetSignerAddress(&_ArtistV5.TransactOpts, _editionId, _newSignerAddress)
}

// SetSignerAddress is a paid mutator transaction binding the contract method 0x56dee996.
//
// Solidity: function setSignerAddress(uint256 _editionId, address _newSignerAddress) returns()
func (_ArtistV5 *ArtistV5TransactorSession) SetSignerAddress(_editionId *big.Int, _newSignerAddress common.Address) (*types.Transaction, error) {
	return _ArtistV5.Contract.SetSignerAddress(&_ArtistV5.TransactOpts, _editionId, _newSignerAddress)
}

// SetStartTime is a paid mutator transaction binding the contract method 0xfbab9e04.
//
// Solidity: function setStartTime(uint256 _editionId, uint32 _startTime) returns()
func (_ArtistV5 *ArtistV5Transactor) SetStartTime(opts *bind.TransactOpts, _editionId *big.Int, _startTime uint32) (*types.Transaction, error) {
	return _ArtistV5.contract.Transact(opts, "setStartTime", _editionId, _startTime)
}

// SetStartTime is a paid mutator transaction binding the contract method 0xfbab9e04.
//
// Solidity: function setStartTime(uint256 _editionId, uint32 _startTime) returns()
func (_ArtistV5 *ArtistV5Session) SetStartTime(_editionId *big.Int, _startTime uint32) (*types.Transaction, error) {
	return _ArtistV5.Contract.SetStartTime(&_ArtistV5.TransactOpts, _editionId, _startTime)
}

// SetStartTime is a paid mutator transaction binding the contract method 0xfbab9e04.
//
// Solidity: function setStartTime(uint256 _editionId, uint32 _startTime) returns()
func (_ArtistV5 *ArtistV5TransactorSession) SetStartTime(_editionId *big.Int, _startTime uint32) (*types.Transaction, error) {
	return _ArtistV5.Contract.SetStartTime(&_ArtistV5.TransactOpts, _editionId, _startTime)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ArtistV5 *ArtistV5Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ArtistV5.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ArtistV5 *ArtistV5Session) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ArtistV5.Contract.TransferFrom(&_ArtistV5.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ArtistV5 *ArtistV5TransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ArtistV5.Contract.TransferFrom(&_ArtistV5.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ArtistV5 *ArtistV5Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ArtistV5.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ArtistV5 *ArtistV5Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ArtistV5.Contract.TransferOwnership(&_ArtistV5.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ArtistV5 *ArtistV5TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ArtistV5.Contract.TransferOwnership(&_ArtistV5.TransactOpts, newOwner)
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0x155dd5ee.
//
// Solidity: function withdrawFunds(uint256 _editionId) returns()
func (_ArtistV5 *ArtistV5Transactor) WithdrawFunds(opts *bind.TransactOpts, _editionId *big.Int) (*types.Transaction, error) {
	return _ArtistV5.contract.Transact(opts, "withdrawFunds", _editionId)
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0x155dd5ee.
//
// Solidity: function withdrawFunds(uint256 _editionId) returns()
func (_ArtistV5 *ArtistV5Session) WithdrawFunds(_editionId *big.Int) (*types.Transaction, error) {
	return _ArtistV5.Contract.WithdrawFunds(&_ArtistV5.TransactOpts, _editionId)
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0x155dd5ee.
//
// Solidity: function withdrawFunds(uint256 _editionId) returns()
func (_ArtistV5 *ArtistV5TransactorSession) WithdrawFunds(_editionId *big.Int) (*types.Transaction, error) {
	return _ArtistV5.Contract.WithdrawFunds(&_ArtistV5.TransactOpts, _editionId)
}

// ArtistV5ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ArtistV5 contract.
type ArtistV5ApprovalIterator struct {
	Event *ArtistV5Approval // Event containing the contract specifics and raw log

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
func (it *ArtistV5ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArtistV5Approval)
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
		it.Event = new(ArtistV5Approval)
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
func (it *ArtistV5ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArtistV5ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArtistV5Approval represents a Approval event raised by the ArtistV5 contract.
type ArtistV5Approval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ArtistV5 *ArtistV5Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*ArtistV5ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ArtistV5.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ArtistV5ApprovalIterator{contract: _ArtistV5.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ArtistV5 *ArtistV5Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ArtistV5Approval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ArtistV5.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArtistV5Approval)
				if err := _ArtistV5.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ArtistV5 *ArtistV5Filterer) ParseApproval(log types.Log) (*ArtistV5Approval, error) {
	event := new(ArtistV5Approval)
	if err := _ArtistV5.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArtistV5ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ArtistV5 contract.
type ArtistV5ApprovalForAllIterator struct {
	Event *ArtistV5ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ArtistV5ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArtistV5ApprovalForAll)
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
		it.Event = new(ArtistV5ApprovalForAll)
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
func (it *ArtistV5ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArtistV5ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArtistV5ApprovalForAll represents a ApprovalForAll event raised by the ArtistV5 contract.
type ArtistV5ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ArtistV5 *ArtistV5Filterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ArtistV5ApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ArtistV5.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ArtistV5ApprovalForAllIterator{contract: _ArtistV5.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ArtistV5 *ArtistV5Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ArtistV5ApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ArtistV5.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArtistV5ApprovalForAll)
				if err := _ArtistV5.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ArtistV5 *ArtistV5Filterer) ParseApprovalForAll(log types.Log) (*ArtistV5ApprovalForAll, error) {
	event := new(ArtistV5ApprovalForAll)
	if err := _ArtistV5.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArtistV5AuctionTimeSetIterator is returned from FilterAuctionTimeSet and is used to iterate over the raw logs and unpacked data for AuctionTimeSet events raised by the ArtistV5 contract.
type ArtistV5AuctionTimeSetIterator struct {
	Event *ArtistV5AuctionTimeSet // Event containing the contract specifics and raw log

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
func (it *ArtistV5AuctionTimeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArtistV5AuctionTimeSet)
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
		it.Event = new(ArtistV5AuctionTimeSet)
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
func (it *ArtistV5AuctionTimeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArtistV5AuctionTimeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArtistV5AuctionTimeSet represents a AuctionTimeSet event raised by the ArtistV5 contract.
type ArtistV5AuctionTimeSet struct {
	TimeType  uint8
	EditionId *big.Int
	NewTime   uint32
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAuctionTimeSet is a free log retrieval operation binding the contract event 0x494264d1227744d2d86690c5355813b007a13955626b261c2e02901a73f6f90c.
//
// Solidity: event AuctionTimeSet(uint8 timeType, uint256 editionId, uint32 indexed newTime)
func (_ArtistV5 *ArtistV5Filterer) FilterAuctionTimeSet(opts *bind.FilterOpts, newTime []uint32) (*ArtistV5AuctionTimeSetIterator, error) {

	var newTimeRule []interface{}
	for _, newTimeItem := range newTime {
		newTimeRule = append(newTimeRule, newTimeItem)
	}

	logs, sub, err := _ArtistV5.contract.FilterLogs(opts, "AuctionTimeSet", newTimeRule)
	if err != nil {
		return nil, err
	}
	return &ArtistV5AuctionTimeSetIterator{contract: _ArtistV5.contract, event: "AuctionTimeSet", logs: logs, sub: sub}, nil
}

// WatchAuctionTimeSet is a free log subscription operation binding the contract event 0x494264d1227744d2d86690c5355813b007a13955626b261c2e02901a73f6f90c.
//
// Solidity: event AuctionTimeSet(uint8 timeType, uint256 editionId, uint32 indexed newTime)
func (_ArtistV5 *ArtistV5Filterer) WatchAuctionTimeSet(opts *bind.WatchOpts, sink chan<- *ArtistV5AuctionTimeSet, newTime []uint32) (event.Subscription, error) {

	var newTimeRule []interface{}
	for _, newTimeItem := range newTime {
		newTimeRule = append(newTimeRule, newTimeItem)
	}

	logs, sub, err := _ArtistV5.contract.WatchLogs(opts, "AuctionTimeSet", newTimeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArtistV5AuctionTimeSet)
				if err := _ArtistV5.contract.UnpackLog(event, "AuctionTimeSet", log); err != nil {
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

// ParseAuctionTimeSet is a log parse operation binding the contract event 0x494264d1227744d2d86690c5355813b007a13955626b261c2e02901a73f6f90c.
//
// Solidity: event AuctionTimeSet(uint8 timeType, uint256 editionId, uint32 indexed newTime)
func (_ArtistV5 *ArtistV5Filterer) ParseAuctionTimeSet(log types.Log) (*ArtistV5AuctionTimeSet, error) {
	event := new(ArtistV5AuctionTimeSet)
	if err := _ArtistV5.contract.UnpackLog(event, "AuctionTimeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArtistV5BaseURISetIterator is returned from FilterBaseURISet and is used to iterate over the raw logs and unpacked data for BaseURISet events raised by the ArtistV5 contract.
type ArtistV5BaseURISetIterator struct {
	Event *ArtistV5BaseURISet // Event containing the contract specifics and raw log

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
func (it *ArtistV5BaseURISetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArtistV5BaseURISet)
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
		it.Event = new(ArtistV5BaseURISet)
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
func (it *ArtistV5BaseURISetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArtistV5BaseURISetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArtistV5BaseURISet represents a BaseURISet event raised by the ArtistV5 contract.
type ArtistV5BaseURISet struct {
	EditionId *big.Int
	BaseURI   string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBaseURISet is a free log retrieval operation binding the contract event 0x1a7fce8987747a468a9fd9d320891f1519376dab1baeb8e72e8d4447fd412589.
//
// Solidity: event BaseURISet(uint256 editionId, string baseURI)
func (_ArtistV5 *ArtistV5Filterer) FilterBaseURISet(opts *bind.FilterOpts) (*ArtistV5BaseURISetIterator, error) {

	logs, sub, err := _ArtistV5.contract.FilterLogs(opts, "BaseURISet")
	if err != nil {
		return nil, err
	}
	return &ArtistV5BaseURISetIterator{contract: _ArtistV5.contract, event: "BaseURISet", logs: logs, sub: sub}, nil
}

// WatchBaseURISet is a free log subscription operation binding the contract event 0x1a7fce8987747a468a9fd9d320891f1519376dab1baeb8e72e8d4447fd412589.
//
// Solidity: event BaseURISet(uint256 editionId, string baseURI)
func (_ArtistV5 *ArtistV5Filterer) WatchBaseURISet(opts *bind.WatchOpts, sink chan<- *ArtistV5BaseURISet) (event.Subscription, error) {

	logs, sub, err := _ArtistV5.contract.WatchLogs(opts, "BaseURISet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArtistV5BaseURISet)
				if err := _ArtistV5.contract.UnpackLog(event, "BaseURISet", log); err != nil {
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

// ParseBaseURISet is a log parse operation binding the contract event 0x1a7fce8987747a468a9fd9d320891f1519376dab1baeb8e72e8d4447fd412589.
//
// Solidity: event BaseURISet(uint256 editionId, string baseURI)
func (_ArtistV5 *ArtistV5Filterer) ParseBaseURISet(log types.Log) (*ArtistV5BaseURISet, error) {
	event := new(ArtistV5BaseURISet)
	if err := _ArtistV5.contract.UnpackLog(event, "BaseURISet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArtistV5EditionCreatedIterator is returned from FilterEditionCreated and is used to iterate over the raw logs and unpacked data for EditionCreated events raised by the ArtistV5 contract.
type ArtistV5EditionCreatedIterator struct {
	Event *ArtistV5EditionCreated // Event containing the contract specifics and raw log

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
func (it *ArtistV5EditionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArtistV5EditionCreated)
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
		it.Event = new(ArtistV5EditionCreated)
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
func (it *ArtistV5EditionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArtistV5EditionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArtistV5EditionCreated represents a EditionCreated event raised by the ArtistV5 contract.
type ArtistV5EditionCreated struct {
	EditionId            *big.Int
	FundingRecipient     common.Address
	Price                *big.Int
	Quantity             uint32
	RoyaltyBPS           uint32
	StartTime            uint32
	EndTime              uint32
	PermissionedQuantity uint32
	SignerAddress        common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterEditionCreated is a free log retrieval operation binding the contract event 0xb56f9ba6a8af17a142f8ad372c6fc283e81d8c6193b86afe7f6ca901acd698f3.
//
// Solidity: event EditionCreated(uint256 indexed editionId, address fundingRecipient, uint256 price, uint32 quantity, uint32 royaltyBPS, uint32 startTime, uint32 endTime, uint32 permissionedQuantity, address signerAddress)
func (_ArtistV5 *ArtistV5Filterer) FilterEditionCreated(opts *bind.FilterOpts, editionId []*big.Int) (*ArtistV5EditionCreatedIterator, error) {

	var editionIdRule []interface{}
	for _, editionIdItem := range editionId {
		editionIdRule = append(editionIdRule, editionIdItem)
	}

	logs, sub, err := _ArtistV5.contract.FilterLogs(opts, "EditionCreated", editionIdRule)
	if err != nil {
		return nil, err
	}
	return &ArtistV5EditionCreatedIterator{contract: _ArtistV5.contract, event: "EditionCreated", logs: logs, sub: sub}, nil
}

// WatchEditionCreated is a free log subscription operation binding the contract event 0xb56f9ba6a8af17a142f8ad372c6fc283e81d8c6193b86afe7f6ca901acd698f3.
//
// Solidity: event EditionCreated(uint256 indexed editionId, address fundingRecipient, uint256 price, uint32 quantity, uint32 royaltyBPS, uint32 startTime, uint32 endTime, uint32 permissionedQuantity, address signerAddress)
func (_ArtistV5 *ArtistV5Filterer) WatchEditionCreated(opts *bind.WatchOpts, sink chan<- *ArtistV5EditionCreated, editionId []*big.Int) (event.Subscription, error) {

	var editionIdRule []interface{}
	for _, editionIdItem := range editionId {
		editionIdRule = append(editionIdRule, editionIdItem)
	}

	logs, sub, err := _ArtistV5.contract.WatchLogs(opts, "EditionCreated", editionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArtistV5EditionCreated)
				if err := _ArtistV5.contract.UnpackLog(event, "EditionCreated", log); err != nil {
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

// ParseEditionCreated is a log parse operation binding the contract event 0xb56f9ba6a8af17a142f8ad372c6fc283e81d8c6193b86afe7f6ca901acd698f3.
//
// Solidity: event EditionCreated(uint256 indexed editionId, address fundingRecipient, uint256 price, uint32 quantity, uint32 royaltyBPS, uint32 startTime, uint32 endTime, uint32 permissionedQuantity, address signerAddress)
func (_ArtistV5 *ArtistV5Filterer) ParseEditionCreated(log types.Log) (*ArtistV5EditionCreated, error) {
	event := new(ArtistV5EditionCreated)
	if err := _ArtistV5.contract.UnpackLog(event, "EditionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArtistV5EditionPurchasedIterator is returned from FilterEditionPurchased and is used to iterate over the raw logs and unpacked data for EditionPurchased events raised by the ArtistV5 contract.
type ArtistV5EditionPurchasedIterator struct {
	Event *ArtistV5EditionPurchased // Event containing the contract specifics and raw log

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
func (it *ArtistV5EditionPurchasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArtistV5EditionPurchased)
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
		it.Event = new(ArtistV5EditionPurchased)
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
func (it *ArtistV5EditionPurchasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArtistV5EditionPurchasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArtistV5EditionPurchased represents a EditionPurchased event raised by the ArtistV5 contract.
type ArtistV5EditionPurchased struct {
	EditionId    *big.Int
	TokenId      *big.Int
	NumSold      uint32
	Buyer        common.Address
	TicketNumber *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterEditionPurchased is a free log retrieval operation binding the contract event 0xc3e4ad784e3889561b308df24921cf08a47410a8ed63b8afe80c50a002efb251.
//
// Solidity: event EditionPurchased(uint256 indexed editionId, uint256 indexed tokenId, uint32 numSold, address indexed buyer, uint256 ticketNumber)
func (_ArtistV5 *ArtistV5Filterer) FilterEditionPurchased(opts *bind.FilterOpts, editionId []*big.Int, tokenId []*big.Int, buyer []common.Address) (*ArtistV5EditionPurchasedIterator, error) {

	var editionIdRule []interface{}
	for _, editionIdItem := range editionId {
		editionIdRule = append(editionIdRule, editionIdItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _ArtistV5.contract.FilterLogs(opts, "EditionPurchased", editionIdRule, tokenIdRule, buyerRule)
	if err != nil {
		return nil, err
	}
	return &ArtistV5EditionPurchasedIterator{contract: _ArtistV5.contract, event: "EditionPurchased", logs: logs, sub: sub}, nil
}

// WatchEditionPurchased is a free log subscription operation binding the contract event 0xc3e4ad784e3889561b308df24921cf08a47410a8ed63b8afe80c50a002efb251.
//
// Solidity: event EditionPurchased(uint256 indexed editionId, uint256 indexed tokenId, uint32 numSold, address indexed buyer, uint256 ticketNumber)
func (_ArtistV5 *ArtistV5Filterer) WatchEditionPurchased(opts *bind.WatchOpts, sink chan<- *ArtistV5EditionPurchased, editionId []*big.Int, tokenId []*big.Int, buyer []common.Address) (event.Subscription, error) {

	var editionIdRule []interface{}
	for _, editionIdItem := range editionId {
		editionIdRule = append(editionIdRule, editionIdItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _ArtistV5.contract.WatchLogs(opts, "EditionPurchased", editionIdRule, tokenIdRule, buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArtistV5EditionPurchased)
				if err := _ArtistV5.contract.UnpackLog(event, "EditionPurchased", log); err != nil {
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

// ParseEditionPurchased is a log parse operation binding the contract event 0xc3e4ad784e3889561b308df24921cf08a47410a8ed63b8afe80c50a002efb251.
//
// Solidity: event EditionPurchased(uint256 indexed editionId, uint256 indexed tokenId, uint32 numSold, address indexed buyer, uint256 ticketNumber)
func (_ArtistV5 *ArtistV5Filterer) ParseEditionPurchased(log types.Log) (*ArtistV5EditionPurchased, error) {
	event := new(ArtistV5EditionPurchased)
	if err := _ArtistV5.contract.UnpackLog(event, "EditionPurchased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArtistV5OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ArtistV5 contract.
type ArtistV5OwnershipTransferredIterator struct {
	Event *ArtistV5OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ArtistV5OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArtistV5OwnershipTransferred)
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
		it.Event = new(ArtistV5OwnershipTransferred)
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
func (it *ArtistV5OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArtistV5OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArtistV5OwnershipTransferred represents a OwnershipTransferred event raised by the ArtistV5 contract.
type ArtistV5OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ArtistV5 *ArtistV5Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ArtistV5OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ArtistV5.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ArtistV5OwnershipTransferredIterator{contract: _ArtistV5.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ArtistV5 *ArtistV5Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ArtistV5OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ArtistV5.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArtistV5OwnershipTransferred)
				if err := _ArtistV5.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ArtistV5 *ArtistV5Filterer) ParseOwnershipTransferred(log types.Log) (*ArtistV5OwnershipTransferred, error) {
	event := new(ArtistV5OwnershipTransferred)
	if err := _ArtistV5.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArtistV5PermissionedQuantitySetIterator is returned from FilterPermissionedQuantitySet and is used to iterate over the raw logs and unpacked data for PermissionedQuantitySet events raised by the ArtistV5 contract.
type ArtistV5PermissionedQuantitySetIterator struct {
	Event *ArtistV5PermissionedQuantitySet // Event containing the contract specifics and raw log

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
func (it *ArtistV5PermissionedQuantitySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArtistV5PermissionedQuantitySet)
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
		it.Event = new(ArtistV5PermissionedQuantitySet)
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
func (it *ArtistV5PermissionedQuantitySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArtistV5PermissionedQuantitySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArtistV5PermissionedQuantitySet represents a PermissionedQuantitySet event raised by the ArtistV5 contract.
type ArtistV5PermissionedQuantitySet struct {
	EditionId            *big.Int
	PermissionedQuantity uint32
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterPermissionedQuantitySet is a free log retrieval operation binding the contract event 0x1de856603e1e748b6f6726e7a51520fe7c63e9b801b8e4b2f8de1a02ae0a8dae.
//
// Solidity: event PermissionedQuantitySet(uint256 editionId, uint32 permissionedQuantity)
func (_ArtistV5 *ArtistV5Filterer) FilterPermissionedQuantitySet(opts *bind.FilterOpts) (*ArtistV5PermissionedQuantitySetIterator, error) {

	logs, sub, err := _ArtistV5.contract.FilterLogs(opts, "PermissionedQuantitySet")
	if err != nil {
		return nil, err
	}
	return &ArtistV5PermissionedQuantitySetIterator{contract: _ArtistV5.contract, event: "PermissionedQuantitySet", logs: logs, sub: sub}, nil
}

// WatchPermissionedQuantitySet is a free log subscription operation binding the contract event 0x1de856603e1e748b6f6726e7a51520fe7c63e9b801b8e4b2f8de1a02ae0a8dae.
//
// Solidity: event PermissionedQuantitySet(uint256 editionId, uint32 permissionedQuantity)
func (_ArtistV5 *ArtistV5Filterer) WatchPermissionedQuantitySet(opts *bind.WatchOpts, sink chan<- *ArtistV5PermissionedQuantitySet) (event.Subscription, error) {

	logs, sub, err := _ArtistV5.contract.WatchLogs(opts, "PermissionedQuantitySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArtistV5PermissionedQuantitySet)
				if err := _ArtistV5.contract.UnpackLog(event, "PermissionedQuantitySet", log); err != nil {
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

// ParsePermissionedQuantitySet is a log parse operation binding the contract event 0x1de856603e1e748b6f6726e7a51520fe7c63e9b801b8e4b2f8de1a02ae0a8dae.
//
// Solidity: event PermissionedQuantitySet(uint256 editionId, uint32 permissionedQuantity)
func (_ArtistV5 *ArtistV5Filterer) ParsePermissionedQuantitySet(log types.Log) (*ArtistV5PermissionedQuantitySet, error) {
	event := new(ArtistV5PermissionedQuantitySet)
	if err := _ArtistV5.contract.UnpackLog(event, "PermissionedQuantitySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArtistV5RoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the ArtistV5 contract.
type ArtistV5RoleGrantedIterator struct {
	Event *ArtistV5RoleGranted // Event containing the contract specifics and raw log

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
func (it *ArtistV5RoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArtistV5RoleGranted)
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
		it.Event = new(ArtistV5RoleGranted)
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
func (it *ArtistV5RoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArtistV5RoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArtistV5RoleGranted represents a RoleGranted event raised by the ArtistV5 contract.
type ArtistV5RoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ArtistV5 *ArtistV5Filterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ArtistV5RoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ArtistV5.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ArtistV5RoleGrantedIterator{contract: _ArtistV5.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ArtistV5 *ArtistV5Filterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ArtistV5RoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ArtistV5.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArtistV5RoleGranted)
				if err := _ArtistV5.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ArtistV5 *ArtistV5Filterer) ParseRoleGranted(log types.Log) (*ArtistV5RoleGranted, error) {
	event := new(ArtistV5RoleGranted)
	if err := _ArtistV5.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArtistV5RoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the ArtistV5 contract.
type ArtistV5RoleRevokedIterator struct {
	Event *ArtistV5RoleRevoked // Event containing the contract specifics and raw log

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
func (it *ArtistV5RoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArtistV5RoleRevoked)
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
		it.Event = new(ArtistV5RoleRevoked)
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
func (it *ArtistV5RoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArtistV5RoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArtistV5RoleRevoked represents a RoleRevoked event raised by the ArtistV5 contract.
type ArtistV5RoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ArtistV5 *ArtistV5Filterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ArtistV5RoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ArtistV5.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ArtistV5RoleRevokedIterator{contract: _ArtistV5.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ArtistV5 *ArtistV5Filterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ArtistV5RoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ArtistV5.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArtistV5RoleRevoked)
				if err := _ArtistV5.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ArtistV5 *ArtistV5Filterer) ParseRoleRevoked(log types.Log) (*ArtistV5RoleRevoked, error) {
	event := new(ArtistV5RoleRevoked)
	if err := _ArtistV5.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArtistV5SignerAddressSetIterator is returned from FilterSignerAddressSet and is used to iterate over the raw logs and unpacked data for SignerAddressSet events raised by the ArtistV5 contract.
type ArtistV5SignerAddressSetIterator struct {
	Event *ArtistV5SignerAddressSet // Event containing the contract specifics and raw log

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
func (it *ArtistV5SignerAddressSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArtistV5SignerAddressSet)
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
		it.Event = new(ArtistV5SignerAddressSet)
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
func (it *ArtistV5SignerAddressSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArtistV5SignerAddressSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArtistV5SignerAddressSet represents a SignerAddressSet event raised by the ArtistV5 contract.
type ArtistV5SignerAddressSet struct {
	EditionId     *big.Int
	SignerAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSignerAddressSet is a free log retrieval operation binding the contract event 0x73b6ab337f7463563f035a0b9f885872e16ad1b5043b679cdf802cfcbaa3a534.
//
// Solidity: event SignerAddressSet(uint256 editionId, address indexed signerAddress)
func (_ArtistV5 *ArtistV5Filterer) FilterSignerAddressSet(opts *bind.FilterOpts, signerAddress []common.Address) (*ArtistV5SignerAddressSetIterator, error) {

	var signerAddressRule []interface{}
	for _, signerAddressItem := range signerAddress {
		signerAddressRule = append(signerAddressRule, signerAddressItem)
	}

	logs, sub, err := _ArtistV5.contract.FilterLogs(opts, "SignerAddressSet", signerAddressRule)
	if err != nil {
		return nil, err
	}
	return &ArtistV5SignerAddressSetIterator{contract: _ArtistV5.contract, event: "SignerAddressSet", logs: logs, sub: sub}, nil
}

// WatchSignerAddressSet is a free log subscription operation binding the contract event 0x73b6ab337f7463563f035a0b9f885872e16ad1b5043b679cdf802cfcbaa3a534.
//
// Solidity: event SignerAddressSet(uint256 editionId, address indexed signerAddress)
func (_ArtistV5 *ArtistV5Filterer) WatchSignerAddressSet(opts *bind.WatchOpts, sink chan<- *ArtistV5SignerAddressSet, signerAddress []common.Address) (event.Subscription, error) {

	var signerAddressRule []interface{}
	for _, signerAddressItem := range signerAddress {
		signerAddressRule = append(signerAddressRule, signerAddressItem)
	}

	logs, sub, err := _ArtistV5.contract.WatchLogs(opts, "SignerAddressSet", signerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArtistV5SignerAddressSet)
				if err := _ArtistV5.contract.UnpackLog(event, "SignerAddressSet", log); err != nil {
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

// ParseSignerAddressSet is a log parse operation binding the contract event 0x73b6ab337f7463563f035a0b9f885872e16ad1b5043b679cdf802cfcbaa3a534.
//
// Solidity: event SignerAddressSet(uint256 editionId, address indexed signerAddress)
func (_ArtistV5 *ArtistV5Filterer) ParseSignerAddressSet(log types.Log) (*ArtistV5SignerAddressSet, error) {
	event := new(ArtistV5SignerAddressSet)
	if err := _ArtistV5.contract.UnpackLog(event, "SignerAddressSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArtistV5TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ArtistV5 contract.
type ArtistV5TransferIterator struct {
	Event *ArtistV5Transfer // Event containing the contract specifics and raw log

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
func (it *ArtistV5TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArtistV5Transfer)
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
		it.Event = new(ArtistV5Transfer)
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
func (it *ArtistV5TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArtistV5TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArtistV5Transfer represents a Transfer event raised by the ArtistV5 contract.
type ArtistV5Transfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ArtistV5 *ArtistV5Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*ArtistV5TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ArtistV5.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ArtistV5TransferIterator{contract: _ArtistV5.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ArtistV5 *ArtistV5Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ArtistV5Transfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ArtistV5.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArtistV5Transfer)
				if err := _ArtistV5.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ArtistV5 *ArtistV5Filterer) ParseTransfer(log types.Log) (*ArtistV5Transfer, error) {
	event := new(ArtistV5Transfer)
	if err := _ArtistV5.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
