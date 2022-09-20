// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zora

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

// DecimalD256 is an auto generated low-level Go binding around an user-defined struct.
type DecimalD256 struct {
	Value *big.Int
}

// IMarketAsk is an auto generated low-level Go binding around an user-defined struct.
type IMarketAsk struct {
	Amount   *big.Int
	Currency common.Address
}

// IMarketBid is an auto generated low-level Go binding around an user-defined struct.
type IMarketBid struct {
	Amount      *big.Int
	Currency    common.Address
	Bidder      common.Address
	Recipient   common.Address
	SellOnShare DecimalD256
}

// IMarketBidShares is an auto generated low-level Go binding around an user-defined struct.
type IMarketBidShares struct {
	PrevOwner DecimalD256
	Creator   DecimalD256
	Owner     DecimalD256
}

// IMediaEIP712Signature is an auto generated low-level Go binding around an user-defined struct.
type IMediaEIP712Signature struct {
	Deadline *big.Int
	V        uint8
	R        [32]byte
	S        [32]byte
}

// IMediaMediaData is an auto generated low-level Go binding around an user-defined struct.
type IMediaMediaData struct {
	TokenURI     string
	MetadataURI  string
	ContentHash  [32]byte
	MetadataHash [32]byte
}

// ZoraMetaData contains all meta data concerning the Zora contract.
var ZoraMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"marketContractAddr\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_uri\",\"type\":\"string\"}],\"name\":\"TokenMetadataURIUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_uri\",\"type\":\"string\"}],\"name\":\"TokenURIUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MINT_WITH_SIG_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMIT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structDecimal.D256\",\"name\":\"sellOnShare\",\"type\":\"tuple\"}],\"internalType\":\"structIMarket.Bid\",\"name\":\"bid\",\"type\":\"tuple\"}],\"name\":\"acceptBid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"auctionTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"contentHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"metadataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structIMedia.MediaData\",\"name\":\"data\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structDecimal.D256\",\"name\":\"prevOwner\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structDecimal.D256\",\"name\":\"creator\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structDecimal.D256\",\"name\":\"owner\",\"type\":\"tuple\"}],\"internalType\":\"structIMarket.BidShares\",\"name\":\"bidShares\",\"type\":\"tuple\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"contentHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"metadataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structIMedia.MediaData\",\"name\":\"data\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structDecimal.D256\",\"name\":\"prevOwner\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structDecimal.D256\",\"name\":\"creator\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structDecimal.D256\",\"name\":\"owner\",\"type\":\"tuple\"}],\"internalType\":\"structIMarket.BidShares\",\"name\":\"bidShares\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIMedia.EIP712Signature\",\"name\":\"sig\",\"type\":\"tuple\"}],\"name\":\"mintWithSig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"mintWithSigNonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIMedia.EIP712Signature\",\"name\":\"sig\",\"type\":\"tuple\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"permitNonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"previousTokenOwners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"removeAsk\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"removeBid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"revokeApproval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"}],\"internalType\":\"structIMarket.Ask\",\"name\":\"ask\",\"type\":\"tuple\"}],\"name\":\"setAsk\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structDecimal.D256\",\"name\":\"sellOnShare\",\"type\":\"tuple\"}],\"internalType\":\"structIMarket.Bid\",\"name\":\"bid\",\"type\":\"tuple\"}],\"name\":\"setBid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenContentHashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenCreators\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenMetadataHashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenMetadataURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"}],\"name\":\"updateTokenMetadataURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"name\":\"updateTokenURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ZoraABI is the input ABI used to generate the binding from.
// Deprecated: Use ZoraMetaData.ABI instead.
var ZoraABI = ZoraMetaData.ABI

// Zora is an auto generated Go binding around an Ethereum contract.
type Zora struct {
	ZoraCaller     // Read-only binding to the contract
	ZoraTransactor // Write-only binding to the contract
	ZoraFilterer   // Log filterer for contract events
}

// ZoraCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZoraCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZoraTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZoraTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZoraFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZoraFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZoraSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZoraSession struct {
	Contract     *Zora             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZoraCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZoraCallerSession struct {
	Contract *ZoraCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ZoraTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZoraTransactorSession struct {
	Contract     *ZoraTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZoraRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZoraRaw struct {
	Contract *Zora // Generic contract binding to access the raw methods on
}

// ZoraCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZoraCallerRaw struct {
	Contract *ZoraCaller // Generic read-only contract binding to access the raw methods on
}

// ZoraTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZoraTransactorRaw struct {
	Contract *ZoraTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZora creates a new instance of Zora, bound to a specific deployed contract.
func NewZora(address common.Address, backend bind.ContractBackend) (*Zora, error) {
	contract, err := bindZora(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Zora{ZoraCaller: ZoraCaller{contract: contract}, ZoraTransactor: ZoraTransactor{contract: contract}, ZoraFilterer: ZoraFilterer{contract: contract}}, nil
}

// NewZoraCaller creates a new read-only instance of Zora, bound to a specific deployed contract.
func NewZoraCaller(address common.Address, caller bind.ContractCaller) (*ZoraCaller, error) {
	contract, err := bindZora(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZoraCaller{contract: contract}, nil
}

// NewZoraTransactor creates a new write-only instance of Zora, bound to a specific deployed contract.
func NewZoraTransactor(address common.Address, transactor bind.ContractTransactor) (*ZoraTransactor, error) {
	contract, err := bindZora(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZoraTransactor{contract: contract}, nil
}

// NewZoraFilterer creates a new log filterer instance of Zora, bound to a specific deployed contract.
func NewZoraFilterer(address common.Address, filterer bind.ContractFilterer) (*ZoraFilterer, error) {
	contract, err := bindZora(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZoraFilterer{contract: contract}, nil
}

// bindZora binds a generic wrapper to an already deployed contract.
func bindZora(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZoraABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Zora *ZoraRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Zora.Contract.ZoraCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Zora *ZoraRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Zora.Contract.ZoraTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Zora *ZoraRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Zora.Contract.ZoraTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Zora *ZoraCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Zora.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Zora *ZoraTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Zora.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Zora *ZoraTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Zora.Contract.contract.Transact(opts, method, params...)
}

// MINTWITHSIGTYPEHASH is a free data retrieval call binding the contract method 0xde5236fb.
//
// Solidity: function MINT_WITH_SIG_TYPEHASH() view returns(bytes32)
func (_Zora *ZoraCaller) MINTWITHSIGTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Zora.contract.Call(opts, &out, "MINT_WITH_SIG_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MINTWITHSIGTYPEHASH is a free data retrieval call binding the contract method 0xde5236fb.
//
// Solidity: function MINT_WITH_SIG_TYPEHASH() view returns(bytes32)
func (_Zora *ZoraSession) MINTWITHSIGTYPEHASH() ([32]byte, error) {
	return _Zora.Contract.MINTWITHSIGTYPEHASH(&_Zora.CallOpts)
}

// MINTWITHSIGTYPEHASH is a free data retrieval call binding the contract method 0xde5236fb.
//
// Solidity: function MINT_WITH_SIG_TYPEHASH() view returns(bytes32)
func (_Zora *ZoraCallerSession) MINTWITHSIGTYPEHASH() ([32]byte, error) {
	return _Zora.Contract.MINTWITHSIGTYPEHASH(&_Zora.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_Zora *ZoraCaller) PERMITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Zora.contract.Call(opts, &out, "PERMIT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_Zora *ZoraSession) PERMITTYPEHASH() ([32]byte, error) {
	return _Zora.Contract.PERMITTYPEHASH(&_Zora.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_Zora *ZoraCallerSession) PERMITTYPEHASH() ([32]byte, error) {
	return _Zora.Contract.PERMITTYPEHASH(&_Zora.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Zora *ZoraCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Zora.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Zora *ZoraSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Zora.Contract.BalanceOf(&_Zora.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Zora *ZoraCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Zora.Contract.BalanceOf(&_Zora.CallOpts, owner)
}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_Zora *ZoraCaller) BaseURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Zora.contract.Call(opts, &out, "baseURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_Zora *ZoraSession) BaseURI() (string, error) {
	return _Zora.Contract.BaseURI(&_Zora.CallOpts)
}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_Zora *ZoraCallerSession) BaseURI() (string, error) {
	return _Zora.Contract.BaseURI(&_Zora.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Zora *ZoraCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Zora.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Zora *ZoraSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Zora.Contract.GetApproved(&_Zora.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Zora *ZoraCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Zora.Contract.GetApproved(&_Zora.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Zora *ZoraCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Zora.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Zora *ZoraSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Zora.Contract.IsApprovedForAll(&_Zora.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Zora *ZoraCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Zora.Contract.IsApprovedForAll(&_Zora.CallOpts, owner, operator)
}

// MarketContract is a free data retrieval call binding the contract method 0xa1794bcd.
//
// Solidity: function marketContract() view returns(address)
func (_Zora *ZoraCaller) MarketContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Zora.contract.Call(opts, &out, "marketContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MarketContract is a free data retrieval call binding the contract method 0xa1794bcd.
//
// Solidity: function marketContract() view returns(address)
func (_Zora *ZoraSession) MarketContract() (common.Address, error) {
	return _Zora.Contract.MarketContract(&_Zora.CallOpts)
}

// MarketContract is a free data retrieval call binding the contract method 0xa1794bcd.
//
// Solidity: function marketContract() view returns(address)
func (_Zora *ZoraCallerSession) MarketContract() (common.Address, error) {
	return _Zora.Contract.MarketContract(&_Zora.CallOpts)
}

// MintWithSigNonces is a free data retrieval call binding the contract method 0x0bcd899b.
//
// Solidity: function mintWithSigNonces(address ) view returns(uint256)
func (_Zora *ZoraCaller) MintWithSigNonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Zora.contract.Call(opts, &out, "mintWithSigNonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintWithSigNonces is a free data retrieval call binding the contract method 0x0bcd899b.
//
// Solidity: function mintWithSigNonces(address ) view returns(uint256)
func (_Zora *ZoraSession) MintWithSigNonces(arg0 common.Address) (*big.Int, error) {
	return _Zora.Contract.MintWithSigNonces(&_Zora.CallOpts, arg0)
}

// MintWithSigNonces is a free data retrieval call binding the contract method 0x0bcd899b.
//
// Solidity: function mintWithSigNonces(address ) view returns(uint256)
func (_Zora *ZoraCallerSession) MintWithSigNonces(arg0 common.Address) (*big.Int, error) {
	return _Zora.Contract.MintWithSigNonces(&_Zora.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Zora *ZoraCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Zora.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Zora *ZoraSession) Name() (string, error) {
	return _Zora.Contract.Name(&_Zora.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Zora *ZoraCallerSession) Name() (string, error) {
	return _Zora.Contract.Name(&_Zora.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Zora *ZoraCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Zora.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Zora *ZoraSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Zora.Contract.OwnerOf(&_Zora.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Zora *ZoraCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Zora.Contract.OwnerOf(&_Zora.CallOpts, tokenId)
}

// PermitNonces is a free data retrieval call binding the contract method 0xf8ccd5de.
//
// Solidity: function permitNonces(address , uint256 ) view returns(uint256)
func (_Zora *ZoraCaller) PermitNonces(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Zora.contract.Call(opts, &out, "permitNonces", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PermitNonces is a free data retrieval call binding the contract method 0xf8ccd5de.
//
// Solidity: function permitNonces(address , uint256 ) view returns(uint256)
func (_Zora *ZoraSession) PermitNonces(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Zora.Contract.PermitNonces(&_Zora.CallOpts, arg0, arg1)
}

// PermitNonces is a free data retrieval call binding the contract method 0xf8ccd5de.
//
// Solidity: function permitNonces(address , uint256 ) view returns(uint256)
func (_Zora *ZoraCallerSession) PermitNonces(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Zora.Contract.PermitNonces(&_Zora.CallOpts, arg0, arg1)
}

// PreviousTokenOwners is a free data retrieval call binding the contract method 0x9d8e7260.
//
// Solidity: function previousTokenOwners(uint256 ) view returns(address)
func (_Zora *ZoraCaller) PreviousTokenOwners(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Zora.contract.Call(opts, &out, "previousTokenOwners", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PreviousTokenOwners is a free data retrieval call binding the contract method 0x9d8e7260.
//
// Solidity: function previousTokenOwners(uint256 ) view returns(address)
func (_Zora *ZoraSession) PreviousTokenOwners(arg0 *big.Int) (common.Address, error) {
	return _Zora.Contract.PreviousTokenOwners(&_Zora.CallOpts, arg0)
}

// PreviousTokenOwners is a free data retrieval call binding the contract method 0x9d8e7260.
//
// Solidity: function previousTokenOwners(uint256 ) view returns(address)
func (_Zora *ZoraCallerSession) PreviousTokenOwners(arg0 *big.Int) (common.Address, error) {
	return _Zora.Contract.PreviousTokenOwners(&_Zora.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Zora *ZoraCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Zora.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Zora *ZoraSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Zora.Contract.SupportsInterface(&_Zora.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Zora *ZoraCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Zora.Contract.SupportsInterface(&_Zora.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Zora *ZoraCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Zora.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Zora *ZoraSession) Symbol() (string, error) {
	return _Zora.Contract.Symbol(&_Zora.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Zora *ZoraCallerSession) Symbol() (string, error) {
	return _Zora.Contract.Symbol(&_Zora.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Zora *ZoraCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Zora.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Zora *ZoraSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Zora.Contract.TokenByIndex(&_Zora.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Zora *ZoraCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Zora.Contract.TokenByIndex(&_Zora.CallOpts, index)
}

// TokenContentHashes is a free data retrieval call binding the contract method 0xfad32197.
//
// Solidity: function tokenContentHashes(uint256 ) view returns(bytes32)
func (_Zora *ZoraCaller) TokenContentHashes(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Zora.contract.Call(opts, &out, "tokenContentHashes", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TokenContentHashes is a free data retrieval call binding the contract method 0xfad32197.
//
// Solidity: function tokenContentHashes(uint256 ) view returns(bytes32)
func (_Zora *ZoraSession) TokenContentHashes(arg0 *big.Int) ([32]byte, error) {
	return _Zora.Contract.TokenContentHashes(&_Zora.CallOpts, arg0)
}

// TokenContentHashes is a free data retrieval call binding the contract method 0xfad32197.
//
// Solidity: function tokenContentHashes(uint256 ) view returns(bytes32)
func (_Zora *ZoraCallerSession) TokenContentHashes(arg0 *big.Int) ([32]byte, error) {
	return _Zora.Contract.TokenContentHashes(&_Zora.CallOpts, arg0)
}

// TokenCreators is a free data retrieval call binding the contract method 0xe0fd045f.
//
// Solidity: function tokenCreators(uint256 ) view returns(address)
func (_Zora *ZoraCaller) TokenCreators(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Zora.contract.Call(opts, &out, "tokenCreators", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenCreators is a free data retrieval call binding the contract method 0xe0fd045f.
//
// Solidity: function tokenCreators(uint256 ) view returns(address)
func (_Zora *ZoraSession) TokenCreators(arg0 *big.Int) (common.Address, error) {
	return _Zora.Contract.TokenCreators(&_Zora.CallOpts, arg0)
}

// TokenCreators is a free data retrieval call binding the contract method 0xe0fd045f.
//
// Solidity: function tokenCreators(uint256 ) view returns(address)
func (_Zora *ZoraCallerSession) TokenCreators(arg0 *big.Int) (common.Address, error) {
	return _Zora.Contract.TokenCreators(&_Zora.CallOpts, arg0)
}

// TokenMetadataHashes is a free data retrieval call binding the contract method 0x01ddc3b5.
//
// Solidity: function tokenMetadataHashes(uint256 ) view returns(bytes32)
func (_Zora *ZoraCaller) TokenMetadataHashes(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Zora.contract.Call(opts, &out, "tokenMetadataHashes", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TokenMetadataHashes is a free data retrieval call binding the contract method 0x01ddc3b5.
//
// Solidity: function tokenMetadataHashes(uint256 ) view returns(bytes32)
func (_Zora *ZoraSession) TokenMetadataHashes(arg0 *big.Int) ([32]byte, error) {
	return _Zora.Contract.TokenMetadataHashes(&_Zora.CallOpts, arg0)
}

// TokenMetadataHashes is a free data retrieval call binding the contract method 0x01ddc3b5.
//
// Solidity: function tokenMetadataHashes(uint256 ) view returns(bytes32)
func (_Zora *ZoraCallerSession) TokenMetadataHashes(arg0 *big.Int) ([32]byte, error) {
	return _Zora.Contract.TokenMetadataHashes(&_Zora.CallOpts, arg0)
}

// TokenMetadataURI is a free data retrieval call binding the contract method 0x157c3df9.
//
// Solidity: function tokenMetadataURI(uint256 tokenId) view returns(string)
func (_Zora *ZoraCaller) TokenMetadataURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Zora.contract.Call(opts, &out, "tokenMetadataURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenMetadataURI is a free data retrieval call binding the contract method 0x157c3df9.
//
// Solidity: function tokenMetadataURI(uint256 tokenId) view returns(string)
func (_Zora *ZoraSession) TokenMetadataURI(tokenId *big.Int) (string, error) {
	return _Zora.Contract.TokenMetadataURI(&_Zora.CallOpts, tokenId)
}

// TokenMetadataURI is a free data retrieval call binding the contract method 0x157c3df9.
//
// Solidity: function tokenMetadataURI(uint256 tokenId) view returns(string)
func (_Zora *ZoraCallerSession) TokenMetadataURI(tokenId *big.Int) (string, error) {
	return _Zora.Contract.TokenMetadataURI(&_Zora.CallOpts, tokenId)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Zora *ZoraCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Zora.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Zora *ZoraSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Zora.Contract.TokenOfOwnerByIndex(&_Zora.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Zora *ZoraCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Zora.Contract.TokenOfOwnerByIndex(&_Zora.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Zora *ZoraCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Zora.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Zora *ZoraSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Zora.Contract.TokenURI(&_Zora.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Zora *ZoraCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Zora.Contract.TokenURI(&_Zora.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Zora *ZoraCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Zora.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Zora *ZoraSession) TotalSupply() (*big.Int, error) {
	return _Zora.Contract.TotalSupply(&_Zora.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Zora *ZoraCallerSession) TotalSupply() (*big.Int, error) {
	return _Zora.Contract.TotalSupply(&_Zora.CallOpts)
}

// AcceptBid is a paid mutator transaction binding the contract method 0xba339399.
//
// Solidity: function acceptBid(uint256 tokenId, (uint256,address,address,address,(uint256)) bid) returns()
func (_Zora *ZoraTransactor) AcceptBid(opts *bind.TransactOpts, tokenId *big.Int, bid IMarketBid) (*types.Transaction, error) {
	return _Zora.contract.Transact(opts, "acceptBid", tokenId, bid)
}

// AcceptBid is a paid mutator transaction binding the contract method 0xba339399.
//
// Solidity: function acceptBid(uint256 tokenId, (uint256,address,address,address,(uint256)) bid) returns()
func (_Zora *ZoraSession) AcceptBid(tokenId *big.Int, bid IMarketBid) (*types.Transaction, error) {
	return _Zora.Contract.AcceptBid(&_Zora.TransactOpts, tokenId, bid)
}

// AcceptBid is a paid mutator transaction binding the contract method 0xba339399.
//
// Solidity: function acceptBid(uint256 tokenId, (uint256,address,address,address,(uint256)) bid) returns()
func (_Zora *ZoraTransactorSession) AcceptBid(tokenId *big.Int, bid IMarketBid) (*types.Transaction, error) {
	return _Zora.Contract.AcceptBid(&_Zora.TransactOpts, tokenId, bid)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Zora *ZoraTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Zora.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Zora *ZoraSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Zora.Contract.Approve(&_Zora.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Zora *ZoraTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Zora.Contract.Approve(&_Zora.TransactOpts, to, tokenId)
}

// AuctionTransfer is a paid mutator transaction binding the contract method 0xf6b630f0.
//
// Solidity: function auctionTransfer(uint256 tokenId, address recipient) returns()
func (_Zora *ZoraTransactor) AuctionTransfer(opts *bind.TransactOpts, tokenId *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Zora.contract.Transact(opts, "auctionTransfer", tokenId, recipient)
}

// AuctionTransfer is a paid mutator transaction binding the contract method 0xf6b630f0.
//
// Solidity: function auctionTransfer(uint256 tokenId, address recipient) returns()
func (_Zora *ZoraSession) AuctionTransfer(tokenId *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Zora.Contract.AuctionTransfer(&_Zora.TransactOpts, tokenId, recipient)
}

// AuctionTransfer is a paid mutator transaction binding the contract method 0xf6b630f0.
//
// Solidity: function auctionTransfer(uint256 tokenId, address recipient) returns()
func (_Zora *ZoraTransactorSession) AuctionTransfer(tokenId *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Zora.Contract.AuctionTransfer(&_Zora.TransactOpts, tokenId, recipient)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_Zora *ZoraTransactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Zora.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_Zora *ZoraSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _Zora.Contract.Burn(&_Zora.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_Zora *ZoraTransactorSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _Zora.Contract.Burn(&_Zora.TransactOpts, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x2cca3237.
//
// Solidity: function mint((string,string,bytes32,bytes32) data, ((uint256),(uint256),(uint256)) bidShares) returns()
func (_Zora *ZoraTransactor) Mint(opts *bind.TransactOpts, data IMediaMediaData, bidShares IMarketBidShares) (*types.Transaction, error) {
	return _Zora.contract.Transact(opts, "mint", data, bidShares)
}

// Mint is a paid mutator transaction binding the contract method 0x2cca3237.
//
// Solidity: function mint((string,string,bytes32,bytes32) data, ((uint256),(uint256),(uint256)) bidShares) returns()
func (_Zora *ZoraSession) Mint(data IMediaMediaData, bidShares IMarketBidShares) (*types.Transaction, error) {
	return _Zora.Contract.Mint(&_Zora.TransactOpts, data, bidShares)
}

// Mint is a paid mutator transaction binding the contract method 0x2cca3237.
//
// Solidity: function mint((string,string,bytes32,bytes32) data, ((uint256),(uint256),(uint256)) bidShares) returns()
func (_Zora *ZoraTransactorSession) Mint(data IMediaMediaData, bidShares IMarketBidShares) (*types.Transaction, error) {
	return _Zora.Contract.Mint(&_Zora.TransactOpts, data, bidShares)
}

// MintWithSig is a paid mutator transaction binding the contract method 0x7a7a1202.
//
// Solidity: function mintWithSig(address creator, (string,string,bytes32,bytes32) data, ((uint256),(uint256),(uint256)) bidShares, (uint256,uint8,bytes32,bytes32) sig) returns()
func (_Zora *ZoraTransactor) MintWithSig(opts *bind.TransactOpts, creator common.Address, data IMediaMediaData, bidShares IMarketBidShares, sig IMediaEIP712Signature) (*types.Transaction, error) {
	return _Zora.contract.Transact(opts, "mintWithSig", creator, data, bidShares, sig)
}

// MintWithSig is a paid mutator transaction binding the contract method 0x7a7a1202.
//
// Solidity: function mintWithSig(address creator, (string,string,bytes32,bytes32) data, ((uint256),(uint256),(uint256)) bidShares, (uint256,uint8,bytes32,bytes32) sig) returns()
func (_Zora *ZoraSession) MintWithSig(creator common.Address, data IMediaMediaData, bidShares IMarketBidShares, sig IMediaEIP712Signature) (*types.Transaction, error) {
	return _Zora.Contract.MintWithSig(&_Zora.TransactOpts, creator, data, bidShares, sig)
}

// MintWithSig is a paid mutator transaction binding the contract method 0x7a7a1202.
//
// Solidity: function mintWithSig(address creator, (string,string,bytes32,bytes32) data, ((uint256),(uint256),(uint256)) bidShares, (uint256,uint8,bytes32,bytes32) sig) returns()
func (_Zora *ZoraTransactorSession) MintWithSig(creator common.Address, data IMediaMediaData, bidShares IMarketBidShares, sig IMediaEIP712Signature) (*types.Transaction, error) {
	return _Zora.Contract.MintWithSig(&_Zora.TransactOpts, creator, data, bidShares, sig)
}

// Permit is a paid mutator transaction binding the contract method 0x0e2a1778.
//
// Solidity: function permit(address spender, uint256 tokenId, (uint256,uint8,bytes32,bytes32) sig) returns()
func (_Zora *ZoraTransactor) Permit(opts *bind.TransactOpts, spender common.Address, tokenId *big.Int, sig IMediaEIP712Signature) (*types.Transaction, error) {
	return _Zora.contract.Transact(opts, "permit", spender, tokenId, sig)
}

// Permit is a paid mutator transaction binding the contract method 0x0e2a1778.
//
// Solidity: function permit(address spender, uint256 tokenId, (uint256,uint8,bytes32,bytes32) sig) returns()
func (_Zora *ZoraSession) Permit(spender common.Address, tokenId *big.Int, sig IMediaEIP712Signature) (*types.Transaction, error) {
	return _Zora.Contract.Permit(&_Zora.TransactOpts, spender, tokenId, sig)
}

// Permit is a paid mutator transaction binding the contract method 0x0e2a1778.
//
// Solidity: function permit(address spender, uint256 tokenId, (uint256,uint8,bytes32,bytes32) sig) returns()
func (_Zora *ZoraTransactorSession) Permit(spender common.Address, tokenId *big.Int, sig IMediaEIP712Signature) (*types.Transaction, error) {
	return _Zora.Contract.Permit(&_Zora.TransactOpts, spender, tokenId, sig)
}

// RemoveAsk is a paid mutator transaction binding the contract method 0x28220f35.
//
// Solidity: function removeAsk(uint256 tokenId) returns()
func (_Zora *ZoraTransactor) RemoveAsk(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Zora.contract.Transact(opts, "removeAsk", tokenId)
}

// RemoveAsk is a paid mutator transaction binding the contract method 0x28220f35.
//
// Solidity: function removeAsk(uint256 tokenId) returns()
func (_Zora *ZoraSession) RemoveAsk(tokenId *big.Int) (*types.Transaction, error) {
	return _Zora.Contract.RemoveAsk(&_Zora.TransactOpts, tokenId)
}

// RemoveAsk is a paid mutator transaction binding the contract method 0x28220f35.
//
// Solidity: function removeAsk(uint256 tokenId) returns()
func (_Zora *ZoraTransactorSession) RemoveAsk(tokenId *big.Int) (*types.Transaction, error) {
	return _Zora.Contract.RemoveAsk(&_Zora.TransactOpts, tokenId)
}

// RemoveBid is a paid mutator transaction binding the contract method 0xb320f459.
//
// Solidity: function removeBid(uint256 tokenId) returns()
func (_Zora *ZoraTransactor) RemoveBid(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Zora.contract.Transact(opts, "removeBid", tokenId)
}

// RemoveBid is a paid mutator transaction binding the contract method 0xb320f459.
//
// Solidity: function removeBid(uint256 tokenId) returns()
func (_Zora *ZoraSession) RemoveBid(tokenId *big.Int) (*types.Transaction, error) {
	return _Zora.Contract.RemoveBid(&_Zora.TransactOpts, tokenId)
}

// RemoveBid is a paid mutator transaction binding the contract method 0xb320f459.
//
// Solidity: function removeBid(uint256 tokenId) returns()
func (_Zora *ZoraTransactorSession) RemoveBid(tokenId *big.Int) (*types.Transaction, error) {
	return _Zora.Contract.RemoveBid(&_Zora.TransactOpts, tokenId)
}

// RevokeApproval is a paid mutator transaction binding the contract method 0xb1e130fc.
//
// Solidity: function revokeApproval(uint256 tokenId) returns()
func (_Zora *ZoraTransactor) RevokeApproval(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Zora.contract.Transact(opts, "revokeApproval", tokenId)
}

// RevokeApproval is a paid mutator transaction binding the contract method 0xb1e130fc.
//
// Solidity: function revokeApproval(uint256 tokenId) returns()
func (_Zora *ZoraSession) RevokeApproval(tokenId *big.Int) (*types.Transaction, error) {
	return _Zora.Contract.RevokeApproval(&_Zora.TransactOpts, tokenId)
}

// RevokeApproval is a paid mutator transaction binding the contract method 0xb1e130fc.
//
// Solidity: function revokeApproval(uint256 tokenId) returns()
func (_Zora *ZoraTransactorSession) RevokeApproval(tokenId *big.Int) (*types.Transaction, error) {
	return _Zora.Contract.RevokeApproval(&_Zora.TransactOpts, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Zora *ZoraTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Zora.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Zora *ZoraSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Zora.Contract.SafeTransferFrom(&_Zora.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Zora *ZoraTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Zora.Contract.SafeTransferFrom(&_Zora.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Zora *ZoraTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Zora.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Zora *ZoraSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Zora.Contract.SafeTransferFrom0(&_Zora.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Zora *ZoraTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Zora.Contract.SafeTransferFrom0(&_Zora.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Zora *ZoraTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Zora.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Zora *ZoraSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Zora.Contract.SetApprovalForAll(&_Zora.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Zora *ZoraTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Zora.Contract.SetApprovalForAll(&_Zora.TransactOpts, operator, approved)
}

// SetAsk is a paid mutator transaction binding the contract method 0x62f24b70.
//
// Solidity: function setAsk(uint256 tokenId, (uint256,address) ask) returns()
func (_Zora *ZoraTransactor) SetAsk(opts *bind.TransactOpts, tokenId *big.Int, ask IMarketAsk) (*types.Transaction, error) {
	return _Zora.contract.Transact(opts, "setAsk", tokenId, ask)
}

// SetAsk is a paid mutator transaction binding the contract method 0x62f24b70.
//
// Solidity: function setAsk(uint256 tokenId, (uint256,address) ask) returns()
func (_Zora *ZoraSession) SetAsk(tokenId *big.Int, ask IMarketAsk) (*types.Transaction, error) {
	return _Zora.Contract.SetAsk(&_Zora.TransactOpts, tokenId, ask)
}

// SetAsk is a paid mutator transaction binding the contract method 0x62f24b70.
//
// Solidity: function setAsk(uint256 tokenId, (uint256,address) ask) returns()
func (_Zora *ZoraTransactorSession) SetAsk(tokenId *big.Int, ask IMarketAsk) (*types.Transaction, error) {
	return _Zora.Contract.SetAsk(&_Zora.TransactOpts, tokenId, ask)
}

// SetBid is a paid mutator transaction binding the contract method 0x5bf62422.
//
// Solidity: function setBid(uint256 tokenId, (uint256,address,address,address,(uint256)) bid) returns()
func (_Zora *ZoraTransactor) SetBid(opts *bind.TransactOpts, tokenId *big.Int, bid IMarketBid) (*types.Transaction, error) {
	return _Zora.contract.Transact(opts, "setBid", tokenId, bid)
}

// SetBid is a paid mutator transaction binding the contract method 0x5bf62422.
//
// Solidity: function setBid(uint256 tokenId, (uint256,address,address,address,(uint256)) bid) returns()
func (_Zora *ZoraSession) SetBid(tokenId *big.Int, bid IMarketBid) (*types.Transaction, error) {
	return _Zora.Contract.SetBid(&_Zora.TransactOpts, tokenId, bid)
}

// SetBid is a paid mutator transaction binding the contract method 0x5bf62422.
//
// Solidity: function setBid(uint256 tokenId, (uint256,address,address,address,(uint256)) bid) returns()
func (_Zora *ZoraTransactorSession) SetBid(tokenId *big.Int, bid IMarketBid) (*types.Transaction, error) {
	return _Zora.Contract.SetBid(&_Zora.TransactOpts, tokenId, bid)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Zora *ZoraTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Zora.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Zora *ZoraSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Zora.Contract.TransferFrom(&_Zora.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Zora *ZoraTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Zora.Contract.TransferFrom(&_Zora.TransactOpts, from, to, tokenId)
}

// UpdateTokenMetadataURI is a paid mutator transaction binding the contract method 0x75682e79.
//
// Solidity: function updateTokenMetadataURI(uint256 tokenId, string metadataURI) returns()
func (_Zora *ZoraTransactor) UpdateTokenMetadataURI(opts *bind.TransactOpts, tokenId *big.Int, metadataURI string) (*types.Transaction, error) {
	return _Zora.contract.Transact(opts, "updateTokenMetadataURI", tokenId, metadataURI)
}

// UpdateTokenMetadataURI is a paid mutator transaction binding the contract method 0x75682e79.
//
// Solidity: function updateTokenMetadataURI(uint256 tokenId, string metadataURI) returns()
func (_Zora *ZoraSession) UpdateTokenMetadataURI(tokenId *big.Int, metadataURI string) (*types.Transaction, error) {
	return _Zora.Contract.UpdateTokenMetadataURI(&_Zora.TransactOpts, tokenId, metadataURI)
}

// UpdateTokenMetadataURI is a paid mutator transaction binding the contract method 0x75682e79.
//
// Solidity: function updateTokenMetadataURI(uint256 tokenId, string metadataURI) returns()
func (_Zora *ZoraTransactorSession) UpdateTokenMetadataURI(tokenId *big.Int, metadataURI string) (*types.Transaction, error) {
	return _Zora.Contract.UpdateTokenMetadataURI(&_Zora.TransactOpts, tokenId, metadataURI)
}

// UpdateTokenURI is a paid mutator transaction binding the contract method 0x18e97fd1.
//
// Solidity: function updateTokenURI(uint256 tokenId, string tokenURI) returns()
func (_Zora *ZoraTransactor) UpdateTokenURI(opts *bind.TransactOpts, tokenId *big.Int, tokenURI string) (*types.Transaction, error) {
	return _Zora.contract.Transact(opts, "updateTokenURI", tokenId, tokenURI)
}

// UpdateTokenURI is a paid mutator transaction binding the contract method 0x18e97fd1.
//
// Solidity: function updateTokenURI(uint256 tokenId, string tokenURI) returns()
func (_Zora *ZoraSession) UpdateTokenURI(tokenId *big.Int, tokenURI string) (*types.Transaction, error) {
	return _Zora.Contract.UpdateTokenURI(&_Zora.TransactOpts, tokenId, tokenURI)
}

// UpdateTokenURI is a paid mutator transaction binding the contract method 0x18e97fd1.
//
// Solidity: function updateTokenURI(uint256 tokenId, string tokenURI) returns()
func (_Zora *ZoraTransactorSession) UpdateTokenURI(tokenId *big.Int, tokenURI string) (*types.Transaction, error) {
	return _Zora.Contract.UpdateTokenURI(&_Zora.TransactOpts, tokenId, tokenURI)
}

// ZoraApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Zora contract.
type ZoraApprovalIterator struct {
	Event *ZoraApproval // Event containing the contract specifics and raw log

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
func (it *ZoraApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZoraApproval)
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
		it.Event = new(ZoraApproval)
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
func (it *ZoraApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZoraApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZoraApproval represents a Approval event raised by the Zora contract.
type ZoraApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Zora *ZoraFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*ZoraApprovalIterator, error) {

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

	logs, sub, err := _Zora.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ZoraApprovalIterator{contract: _Zora.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Zora *ZoraFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ZoraApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Zora.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZoraApproval)
				if err := _Zora.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Zora *ZoraFilterer) ParseApproval(log types.Log) (*ZoraApproval, error) {
	event := new(ZoraApproval)
	if err := _Zora.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZoraApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Zora contract.
type ZoraApprovalForAllIterator struct {
	Event *ZoraApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ZoraApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZoraApprovalForAll)
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
		it.Event = new(ZoraApprovalForAll)
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
func (it *ZoraApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZoraApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZoraApprovalForAll represents a ApprovalForAll event raised by the Zora contract.
type ZoraApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Zora *ZoraFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ZoraApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Zora.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ZoraApprovalForAllIterator{contract: _Zora.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Zora *ZoraFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ZoraApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Zora.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZoraApprovalForAll)
				if err := _Zora.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_Zora *ZoraFilterer) ParseApprovalForAll(log types.Log) (*ZoraApprovalForAll, error) {
	event := new(ZoraApprovalForAll)
	if err := _Zora.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZoraTokenMetadataURIUpdatedIterator is returned from FilterTokenMetadataURIUpdated and is used to iterate over the raw logs and unpacked data for TokenMetadataURIUpdated events raised by the Zora contract.
type ZoraTokenMetadataURIUpdatedIterator struct {
	Event *ZoraTokenMetadataURIUpdated // Event containing the contract specifics and raw log

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
func (it *ZoraTokenMetadataURIUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZoraTokenMetadataURIUpdated)
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
		it.Event = new(ZoraTokenMetadataURIUpdated)
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
func (it *ZoraTokenMetadataURIUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZoraTokenMetadataURIUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZoraTokenMetadataURIUpdated represents a TokenMetadataURIUpdated event raised by the Zora contract.
type ZoraTokenMetadataURIUpdated struct {
	TokenId *big.Int
	Owner   common.Address
	Uri     string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTokenMetadataURIUpdated is a free log retrieval operation binding the contract event 0xe3df41127db820c79e5b8d541a63e40e3e97b9af96f7a50bded13091b70df9ae.
//
// Solidity: event TokenMetadataURIUpdated(uint256 indexed _tokenId, address owner, string _uri)
func (_Zora *ZoraFilterer) FilterTokenMetadataURIUpdated(opts *bind.FilterOpts, _tokenId []*big.Int) (*ZoraTokenMetadataURIUpdatedIterator, error) {

	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _Zora.contract.FilterLogs(opts, "TokenMetadataURIUpdated", _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ZoraTokenMetadataURIUpdatedIterator{contract: _Zora.contract, event: "TokenMetadataURIUpdated", logs: logs, sub: sub}, nil
}

// WatchTokenMetadataURIUpdated is a free log subscription operation binding the contract event 0xe3df41127db820c79e5b8d541a63e40e3e97b9af96f7a50bded13091b70df9ae.
//
// Solidity: event TokenMetadataURIUpdated(uint256 indexed _tokenId, address owner, string _uri)
func (_Zora *ZoraFilterer) WatchTokenMetadataURIUpdated(opts *bind.WatchOpts, sink chan<- *ZoraTokenMetadataURIUpdated, _tokenId []*big.Int) (event.Subscription, error) {

	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _Zora.contract.WatchLogs(opts, "TokenMetadataURIUpdated", _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZoraTokenMetadataURIUpdated)
				if err := _Zora.contract.UnpackLog(event, "TokenMetadataURIUpdated", log); err != nil {
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

// ParseTokenMetadataURIUpdated is a log parse operation binding the contract event 0xe3df41127db820c79e5b8d541a63e40e3e97b9af96f7a50bded13091b70df9ae.
//
// Solidity: event TokenMetadataURIUpdated(uint256 indexed _tokenId, address owner, string _uri)
func (_Zora *ZoraFilterer) ParseTokenMetadataURIUpdated(log types.Log) (*ZoraTokenMetadataURIUpdated, error) {
	event := new(ZoraTokenMetadataURIUpdated)
	if err := _Zora.contract.UnpackLog(event, "TokenMetadataURIUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZoraTokenURIUpdatedIterator is returned from FilterTokenURIUpdated and is used to iterate over the raw logs and unpacked data for TokenURIUpdated events raised by the Zora contract.
type ZoraTokenURIUpdatedIterator struct {
	Event *ZoraTokenURIUpdated // Event containing the contract specifics and raw log

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
func (it *ZoraTokenURIUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZoraTokenURIUpdated)
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
		it.Event = new(ZoraTokenURIUpdated)
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
func (it *ZoraTokenURIUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZoraTokenURIUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZoraTokenURIUpdated represents a TokenURIUpdated event raised by the Zora contract.
type ZoraTokenURIUpdated struct {
	TokenId *big.Int
	Owner   common.Address
	Uri     string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTokenURIUpdated is a free log retrieval operation binding the contract event 0x702fe2dc2dc0f68023540aa4a1e11811c0f29112f6ebf01e61b90538e4f29810.
//
// Solidity: event TokenURIUpdated(uint256 indexed _tokenId, address owner, string _uri)
func (_Zora *ZoraFilterer) FilterTokenURIUpdated(opts *bind.FilterOpts, _tokenId []*big.Int) (*ZoraTokenURIUpdatedIterator, error) {

	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _Zora.contract.FilterLogs(opts, "TokenURIUpdated", _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ZoraTokenURIUpdatedIterator{contract: _Zora.contract, event: "TokenURIUpdated", logs: logs, sub: sub}, nil
}

// WatchTokenURIUpdated is a free log subscription operation binding the contract event 0x702fe2dc2dc0f68023540aa4a1e11811c0f29112f6ebf01e61b90538e4f29810.
//
// Solidity: event TokenURIUpdated(uint256 indexed _tokenId, address owner, string _uri)
func (_Zora *ZoraFilterer) WatchTokenURIUpdated(opts *bind.WatchOpts, sink chan<- *ZoraTokenURIUpdated, _tokenId []*big.Int) (event.Subscription, error) {

	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _Zora.contract.WatchLogs(opts, "TokenURIUpdated", _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZoraTokenURIUpdated)
				if err := _Zora.contract.UnpackLog(event, "TokenURIUpdated", log); err != nil {
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

// ParseTokenURIUpdated is a log parse operation binding the contract event 0x702fe2dc2dc0f68023540aa4a1e11811c0f29112f6ebf01e61b90538e4f29810.
//
// Solidity: event TokenURIUpdated(uint256 indexed _tokenId, address owner, string _uri)
func (_Zora *ZoraFilterer) ParseTokenURIUpdated(log types.Log) (*ZoraTokenURIUpdated, error) {
	event := new(ZoraTokenURIUpdated)
	if err := _Zora.contract.UnpackLog(event, "TokenURIUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZoraTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Zora contract.
type ZoraTransferIterator struct {
	Event *ZoraTransfer // Event containing the contract specifics and raw log

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
func (it *ZoraTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZoraTransfer)
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
		it.Event = new(ZoraTransfer)
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
func (it *ZoraTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZoraTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZoraTransfer represents a Transfer event raised by the Zora contract.
type ZoraTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Zora *ZoraFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*ZoraTransferIterator, error) {

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

	logs, sub, err := _Zora.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ZoraTransferIterator{contract: _Zora.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Zora *ZoraFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ZoraTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Zora.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZoraTransfer)
				if err := _Zora.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Zora *ZoraFilterer) ParseTransfer(log types.Log) (*ZoraTransfer, error) {
	event := new(ZoraTransfer)
	if err := _Zora.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
