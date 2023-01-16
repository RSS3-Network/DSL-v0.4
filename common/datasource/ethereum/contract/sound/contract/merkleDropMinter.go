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

// MintInfo is an auto generated low-level Go binding around an user-defined struct.
type MintInfo struct {
	StartTime             uint32
	EndTime               uint32
	AffiliateFeeBPS       uint16
	MintPaused            bool
	Price                 *big.Int
	MaxMintable           uint32
	MaxMintablePerAccount uint32
	TotalMinted           uint32
	MerkleRootHash        [32]byte
}

// MerkleDropMinterMetaData contains all meta data concerning the MerkleDropMinter contract.
var MerkleDropMinterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractISoundFeeRegistry\",\"name\":\"feeRegistry_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"available\",\"type\":\"uint32\"}],\"name\":\"ExceedsAvailableSupply\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExceedsMaxPerAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeRegistryIsZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAffiliateFeeBPS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMerkleProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTimeRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxMintablePerAccountIsZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MerkleRootHashIsEmpty\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"startTime\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"endTime\",\"type\":\"uint32\"}],\"name\":\"MintNotOpen\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintPaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"paid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"}],\"name\":\"Underpaid\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"edition\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"mintId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"bps\",\"type\":\"uint16\"}],\"name\":\"AffiliateFeeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"quantity\",\"type\":\"uint32\"}],\"name\":\"DropClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"edition\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"mintId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"maxMintablePerAccount\",\"type\":\"uint32\"}],\"name\":\"MaxMintablePerAccountSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"edition\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"mintId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"maxMintable\",\"type\":\"uint32\"}],\"name\":\"MaxMintableSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"edition\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"mintId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"merkleRootHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"price\",\"type\":\"uint96\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"startTime\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"endTime\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"affiliateFeeBPS\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"maxMintable\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"maxMintablePerAccount\",\"type\":\"uint32\"}],\"name\":\"MerkleDropMintCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"edition\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"mintId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"merkleRootHash\",\"type\":\"bytes32\"}],\"name\":\"MerkleRootHashSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"edition\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"mintId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"startTime\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"endTime\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"affiliateFeeBPS\",\"type\":\"uint16\"}],\"name\":\"MintConfigCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"edition\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"mintId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"}],\"name\":\"MintPausedSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"edition\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"mintId\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"fromTokenId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"quantity\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"requiredEtherValue\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"platformFee\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"affiliateFee\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"affiliate\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"affiliated\",\"type\":\"bool\"}],\"name\":\"Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"edition\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"mintId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"price\",\"type\":\"uint96\"}],\"name\":\"PriceSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"edition\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"mintId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"startTime\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"endTime\",\"type\":\"uint32\"}],\"name\":\"TimeRangeSet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_BPS\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"affiliate\",\"type\":\"address\"}],\"name\":\"affiliateFeesAccrued\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"edition\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRootHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint96\",\"name\":\"price\",\"type\":\"uint96\"},{\"internalType\":\"uint32\",\"name\":\"startTime\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"endTime\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"affiliateFeeBPS\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxMintable\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxMintablePerAccount\",\"type\":\"uint32\"}],\"name\":\"createEditionMint\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"mintId\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeRegistry\",\"outputs\":[{\"internalType\":\"contractISoundFeeRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"affiliate\",\"type\":\"address\"}],\"name\":\"isAffiliated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"edition\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"mintId\",\"type\":\"uint128\"},{\"internalType\":\"uint32\",\"name\":\"requestedQuantity\",\"type\":\"uint32\"},{\"internalType\":\"bytes32[]\",\"name\":\"merkleProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"address\",\"name\":\"affiliate\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"edition\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"mintId\",\"type\":\"uint128\"}],\"name\":\"mintInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"startTime\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"endTime\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"affiliateFeeBPS\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"mintPaused\",\"type\":\"bool\"},{\"internalType\":\"uint96\",\"name\":\"price\",\"type\":\"uint96\"},{\"internalType\":\"uint32\",\"name\":\"maxMintable\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxMintablePerAccount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"totalMinted\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRootHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMintInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"moduleInterfaceId\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextMintId\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"platformFeesAccrued\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"edition\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"mintId\",\"type\":\"uint128\"},{\"internalType\":\"uint16\",\"name\":\"feeBPS\",\"type\":\"uint16\"}],\"name\":\"setAffiliateFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"edition\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"mintId\",\"type\":\"uint128\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"}],\"name\":\"setEditionMintPaused\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"edition\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"mintId\",\"type\":\"uint128\"},{\"internalType\":\"uint32\",\"name\":\"maxMintable\",\"type\":\"uint32\"}],\"name\":\"setMaxMintable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"edition\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"mintId\",\"type\":\"uint128\"},{\"internalType\":\"uint32\",\"name\":\"maxMintablePerAccount\",\"type\":\"uint32\"}],\"name\":\"setMaxMintablePerAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"edition\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"mintId\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRootHash\",\"type\":\"bytes32\"}],\"name\":\"setMerkleRootHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"edition\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"mintId\",\"type\":\"uint128\"},{\"internalType\":\"uint96\",\"name\":\"price\",\"type\":\"uint96\"}],\"name\":\"setPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"edition\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"mintId\",\"type\":\"uint128\"},{\"internalType\":\"uint32\",\"name\":\"startTime\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"endTime\",\"type\":\"uint32\"}],\"name\":\"setTimeRange\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"edition\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"mintId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"quantity\",\"type\":\"uint32\"}],\"name\":\"totalPrice\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"affiliate\",\"type\":\"address\"}],\"name\":\"withdrawForAffiliate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawForPlatform\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MerkleDropMinterABI is the input ABI used to generate the binding from.
// Deprecated: Use MerkleDropMinterMetaData.ABI instead.
var MerkleDropMinterABI = MerkleDropMinterMetaData.ABI

// MerkleDropMinter is an auto generated Go binding around an Ethereum contract.
type MerkleDropMinter struct {
	MerkleDropMinterCaller     // Read-only binding to the contract
	MerkleDropMinterTransactor // Write-only binding to the contract
	MerkleDropMinterFilterer   // Log filterer for contract events
}

// MerkleDropMinterCaller is an auto generated read-only Go binding around an Ethereum contract.
type MerkleDropMinterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleDropMinterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MerkleDropMinterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleDropMinterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MerkleDropMinterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleDropMinterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MerkleDropMinterSession struct {
	Contract     *MerkleDropMinter // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MerkleDropMinterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MerkleDropMinterCallerSession struct {
	Contract *MerkleDropMinterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// MerkleDropMinterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MerkleDropMinterTransactorSession struct {
	Contract     *MerkleDropMinterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// MerkleDropMinterRaw is an auto generated low-level Go binding around an Ethereum contract.
type MerkleDropMinterRaw struct {
	Contract *MerkleDropMinter // Generic contract binding to access the raw methods on
}

// MerkleDropMinterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MerkleDropMinterCallerRaw struct {
	Contract *MerkleDropMinterCaller // Generic read-only contract binding to access the raw methods on
}

// MerkleDropMinterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MerkleDropMinterTransactorRaw struct {
	Contract *MerkleDropMinterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMerkleDropMinter creates a new instance of MerkleDropMinter, bound to a specific deployed contract.
func NewMerkleDropMinter(address common.Address, backend bind.ContractBackend) (*MerkleDropMinter, error) {
	contract, err := bindMerkleDropMinter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MerkleDropMinter{MerkleDropMinterCaller: MerkleDropMinterCaller{contract: contract}, MerkleDropMinterTransactor: MerkleDropMinterTransactor{contract: contract}, MerkleDropMinterFilterer: MerkleDropMinterFilterer{contract: contract}}, nil
}

// NewMerkleDropMinterCaller creates a new read-only instance of MerkleDropMinter, bound to a specific deployed contract.
func NewMerkleDropMinterCaller(address common.Address, caller bind.ContractCaller) (*MerkleDropMinterCaller, error) {
	contract, err := bindMerkleDropMinter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MerkleDropMinterCaller{contract: contract}, nil
}

// NewMerkleDropMinterTransactor creates a new write-only instance of MerkleDropMinter, bound to a specific deployed contract.
func NewMerkleDropMinterTransactor(address common.Address, transactor bind.ContractTransactor) (*MerkleDropMinterTransactor, error) {
	contract, err := bindMerkleDropMinter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MerkleDropMinterTransactor{contract: contract}, nil
}

// NewMerkleDropMinterFilterer creates a new log filterer instance of MerkleDropMinter, bound to a specific deployed contract.
func NewMerkleDropMinterFilterer(address common.Address, filterer bind.ContractFilterer) (*MerkleDropMinterFilterer, error) {
	contract, err := bindMerkleDropMinter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MerkleDropMinterFilterer{contract: contract}, nil
}

// bindMerkleDropMinter binds a generic wrapper to an already deployed contract.
func bindMerkleDropMinter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MerkleDropMinterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerkleDropMinter *MerkleDropMinterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MerkleDropMinter.Contract.MerkleDropMinterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerkleDropMinter *MerkleDropMinterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerkleDropMinter.Contract.MerkleDropMinterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerkleDropMinter *MerkleDropMinterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerkleDropMinter.Contract.MerkleDropMinterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerkleDropMinter *MerkleDropMinterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MerkleDropMinter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerkleDropMinter *MerkleDropMinterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerkleDropMinter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerkleDropMinter *MerkleDropMinterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerkleDropMinter.Contract.contract.Transact(opts, method, params...)
}

// MAXBPS is a free data retrieval call binding the contract method 0xfd967f47.
//
// Solidity: function MAX_BPS() pure returns(uint16)
func (_MerkleDropMinter *MerkleDropMinterCaller) MAXBPS(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _MerkleDropMinter.contract.Call(opts, &out, "MAX_BPS")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MAXBPS is a free data retrieval call binding the contract method 0xfd967f47.
//
// Solidity: function MAX_BPS() pure returns(uint16)
func (_MerkleDropMinter *MerkleDropMinterSession) MAXBPS() (uint16, error) {
	return _MerkleDropMinter.Contract.MAXBPS(&_MerkleDropMinter.CallOpts)
}

// MAXBPS is a free data retrieval call binding the contract method 0xfd967f47.
//
// Solidity: function MAX_BPS() pure returns(uint16)
func (_MerkleDropMinter *MerkleDropMinterCallerSession) MAXBPS() (uint16, error) {
	return _MerkleDropMinter.Contract.MAXBPS(&_MerkleDropMinter.CallOpts)
}

// AffiliateFeesAccrued is a free data retrieval call binding the contract method 0xac1fc22c.
//
// Solidity: function affiliateFeesAccrued(address affiliate) view returns(uint128)
func (_MerkleDropMinter *MerkleDropMinterCaller) AffiliateFeesAccrued(opts *bind.CallOpts, affiliate common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MerkleDropMinter.contract.Call(opts, &out, "affiliateFeesAccrued", affiliate)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AffiliateFeesAccrued is a free data retrieval call binding the contract method 0xac1fc22c.
//
// Solidity: function affiliateFeesAccrued(address affiliate) view returns(uint128)
func (_MerkleDropMinter *MerkleDropMinterSession) AffiliateFeesAccrued(affiliate common.Address) (*big.Int, error) {
	return _MerkleDropMinter.Contract.AffiliateFeesAccrued(&_MerkleDropMinter.CallOpts, affiliate)
}

// AffiliateFeesAccrued is a free data retrieval call binding the contract method 0xac1fc22c.
//
// Solidity: function affiliateFeesAccrued(address affiliate) view returns(uint128)
func (_MerkleDropMinter *MerkleDropMinterCallerSession) AffiliateFeesAccrued(affiliate common.Address) (*big.Int, error) {
	return _MerkleDropMinter.Contract.AffiliateFeesAccrued(&_MerkleDropMinter.CallOpts, affiliate)
}

// FeeRegistry is a free data retrieval call binding the contract method 0xb3a408b8.
//
// Solidity: function feeRegistry() view returns(address)
func (_MerkleDropMinter *MerkleDropMinterCaller) FeeRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MerkleDropMinter.contract.Call(opts, &out, "feeRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeRegistry is a free data retrieval call binding the contract method 0xb3a408b8.
//
// Solidity: function feeRegistry() view returns(address)
func (_MerkleDropMinter *MerkleDropMinterSession) FeeRegistry() (common.Address, error) {
	return _MerkleDropMinter.Contract.FeeRegistry(&_MerkleDropMinter.CallOpts)
}

// FeeRegistry is a free data retrieval call binding the contract method 0xb3a408b8.
//
// Solidity: function feeRegistry() view returns(address)
func (_MerkleDropMinter *MerkleDropMinterCallerSession) FeeRegistry() (common.Address, error) {
	return _MerkleDropMinter.Contract.FeeRegistry(&_MerkleDropMinter.CallOpts)
}

// IsAffiliated is a free data retrieval call binding the contract method 0xc94dcf8f.
//
// Solidity: function isAffiliated(address , uint128 , address affiliate) view returns(bool)
func (_MerkleDropMinter *MerkleDropMinterCaller) IsAffiliated(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int, affiliate common.Address) (bool, error) {
	var out []interface{}
	err := _MerkleDropMinter.contract.Call(opts, &out, "isAffiliated", arg0, arg1, affiliate)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAffiliated is a free data retrieval call binding the contract method 0xc94dcf8f.
//
// Solidity: function isAffiliated(address , uint128 , address affiliate) view returns(bool)
func (_MerkleDropMinter *MerkleDropMinterSession) IsAffiliated(arg0 common.Address, arg1 *big.Int, affiliate common.Address) (bool, error) {
	return _MerkleDropMinter.Contract.IsAffiliated(&_MerkleDropMinter.CallOpts, arg0, arg1, affiliate)
}

// IsAffiliated is a free data retrieval call binding the contract method 0xc94dcf8f.
//
// Solidity: function isAffiliated(address , uint128 , address affiliate) view returns(bool)
func (_MerkleDropMinter *MerkleDropMinterCallerSession) IsAffiliated(arg0 common.Address, arg1 *big.Int, affiliate common.Address) (bool, error) {
	return _MerkleDropMinter.Contract.IsAffiliated(&_MerkleDropMinter.CallOpts, arg0, arg1, affiliate)
}

// MintInfo is a free data retrieval call binding the contract method 0x03a6ccd0.
//
// Solidity: function mintInfo(address edition, uint128 mintId) view returns((uint32,uint32,uint16,bool,uint96,uint32,uint32,uint32,bytes32))
func (_MerkleDropMinter *MerkleDropMinterCaller) MintInfo(opts *bind.CallOpts, edition common.Address, mintId *big.Int) (MintInfo, error) {
	var out []interface{}
	err := _MerkleDropMinter.contract.Call(opts, &out, "mintInfo", edition, mintId)

	if err != nil {
		return *new(MintInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(MintInfo)).(*MintInfo)

	return out0, err

}

// MintInfo is a free data retrieval call binding the contract method 0x03a6ccd0.
//
// Solidity: function mintInfo(address edition, uint128 mintId) view returns((uint32,uint32,uint16,bool,uint96,uint32,uint32,uint32,bytes32))
func (_MerkleDropMinter *MerkleDropMinterSession) MintInfo(edition common.Address, mintId *big.Int) (MintInfo, error) {
	return _MerkleDropMinter.Contract.MintInfo(&_MerkleDropMinter.CallOpts, edition, mintId)
}

// MintInfo is a free data retrieval call binding the contract method 0x03a6ccd0.
//
// Solidity: function mintInfo(address edition, uint128 mintId) view returns((uint32,uint32,uint16,bool,uint96,uint32,uint32,uint32,bytes32))
func (_MerkleDropMinter *MerkleDropMinterCallerSession) MintInfo(edition common.Address, mintId *big.Int) (MintInfo, error) {
	return _MerkleDropMinter.Contract.MintInfo(&_MerkleDropMinter.CallOpts, edition, mintId)
}

// ModuleInterfaceId is a free data retrieval call binding the contract method 0xc7dd3228.
//
// Solidity: function moduleInterfaceId() pure returns(bytes4)
func (_MerkleDropMinter *MerkleDropMinterCaller) ModuleInterfaceId(opts *bind.CallOpts) ([4]byte, error) {
	var out []interface{}
	err := _MerkleDropMinter.contract.Call(opts, &out, "moduleInterfaceId")

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// ModuleInterfaceId is a free data retrieval call binding the contract method 0xc7dd3228.
//
// Solidity: function moduleInterfaceId() pure returns(bytes4)
func (_MerkleDropMinter *MerkleDropMinterSession) ModuleInterfaceId() ([4]byte, error) {
	return _MerkleDropMinter.Contract.ModuleInterfaceId(&_MerkleDropMinter.CallOpts)
}

// ModuleInterfaceId is a free data retrieval call binding the contract method 0xc7dd3228.
//
// Solidity: function moduleInterfaceId() pure returns(bytes4)
func (_MerkleDropMinter *MerkleDropMinterCallerSession) ModuleInterfaceId() ([4]byte, error) {
	return _MerkleDropMinter.Contract.ModuleInterfaceId(&_MerkleDropMinter.CallOpts)
}

// NextMintId is a free data retrieval call binding the contract method 0x6aa99da3.
//
// Solidity: function nextMintId() view returns(uint128)
func (_MerkleDropMinter *MerkleDropMinterCaller) NextMintId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MerkleDropMinter.contract.Call(opts, &out, "nextMintId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextMintId is a free data retrieval call binding the contract method 0x6aa99da3.
//
// Solidity: function nextMintId() view returns(uint128)
func (_MerkleDropMinter *MerkleDropMinterSession) NextMintId() (*big.Int, error) {
	return _MerkleDropMinter.Contract.NextMintId(&_MerkleDropMinter.CallOpts)
}

// NextMintId is a free data retrieval call binding the contract method 0x6aa99da3.
//
// Solidity: function nextMintId() view returns(uint128)
func (_MerkleDropMinter *MerkleDropMinterCallerSession) NextMintId() (*big.Int, error) {
	return _MerkleDropMinter.Contract.NextMintId(&_MerkleDropMinter.CallOpts)
}

// PlatformFeesAccrued is a free data retrieval call binding the contract method 0x22163b86.
//
// Solidity: function platformFeesAccrued() view returns(uint128)
func (_MerkleDropMinter *MerkleDropMinterCaller) PlatformFeesAccrued(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MerkleDropMinter.contract.Call(opts, &out, "platformFeesAccrued")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PlatformFeesAccrued is a free data retrieval call binding the contract method 0x22163b86.
//
// Solidity: function platformFeesAccrued() view returns(uint128)
func (_MerkleDropMinter *MerkleDropMinterSession) PlatformFeesAccrued() (*big.Int, error) {
	return _MerkleDropMinter.Contract.PlatformFeesAccrued(&_MerkleDropMinter.CallOpts)
}

// PlatformFeesAccrued is a free data retrieval call binding the contract method 0x22163b86.
//
// Solidity: function platformFeesAccrued() view returns(uint128)
func (_MerkleDropMinter *MerkleDropMinterCallerSession) PlatformFeesAccrued() (*big.Int, error) {
	return _MerkleDropMinter.Contract.PlatformFeesAccrued(&_MerkleDropMinter.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MerkleDropMinter *MerkleDropMinterCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _MerkleDropMinter.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MerkleDropMinter *MerkleDropMinterSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MerkleDropMinter.Contract.SupportsInterface(&_MerkleDropMinter.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MerkleDropMinter *MerkleDropMinterCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MerkleDropMinter.Contract.SupportsInterface(&_MerkleDropMinter.CallOpts, interfaceId)
}

// TotalPrice is a free data retrieval call binding the contract method 0xb84158a4.
//
// Solidity: function totalPrice(address edition, uint128 mintId, address , uint32 quantity) view returns(uint128)
func (_MerkleDropMinter *MerkleDropMinterCaller) TotalPrice(opts *bind.CallOpts, edition common.Address, mintId *big.Int, arg2 common.Address, quantity uint32) (*big.Int, error) {
	var out []interface{}
	err := _MerkleDropMinter.contract.Call(opts, &out, "totalPrice", edition, mintId, arg2, quantity)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalPrice is a free data retrieval call binding the contract method 0xb84158a4.
//
// Solidity: function totalPrice(address edition, uint128 mintId, address , uint32 quantity) view returns(uint128)
func (_MerkleDropMinter *MerkleDropMinterSession) TotalPrice(edition common.Address, mintId *big.Int, arg2 common.Address, quantity uint32) (*big.Int, error) {
	return _MerkleDropMinter.Contract.TotalPrice(&_MerkleDropMinter.CallOpts, edition, mintId, arg2, quantity)
}

// TotalPrice is a free data retrieval call binding the contract method 0xb84158a4.
//
// Solidity: function totalPrice(address edition, uint128 mintId, address , uint32 quantity) view returns(uint128)
func (_MerkleDropMinter *MerkleDropMinterCallerSession) TotalPrice(edition common.Address, mintId *big.Int, arg2 common.Address, quantity uint32) (*big.Int, error) {
	return _MerkleDropMinter.Contract.TotalPrice(&_MerkleDropMinter.CallOpts, edition, mintId, arg2, quantity)
}

// CreateEditionMint is a paid mutator transaction binding the contract method 0xc309c47d.
//
// Solidity: function createEditionMint(address edition, bytes32 merkleRootHash, uint96 price, uint32 startTime, uint32 endTime, uint16 affiliateFeeBPS, uint32 maxMintable, uint32 maxMintablePerAccount) returns(uint128 mintId)
func (_MerkleDropMinter *MerkleDropMinterTransactor) CreateEditionMint(opts *bind.TransactOpts, edition common.Address, merkleRootHash [32]byte, price *big.Int, startTime uint32, endTime uint32, affiliateFeeBPS uint16, maxMintable uint32, maxMintablePerAccount uint32) (*types.Transaction, error) {
	return _MerkleDropMinter.contract.Transact(opts, "createEditionMint", edition, merkleRootHash, price, startTime, endTime, affiliateFeeBPS, maxMintable, maxMintablePerAccount)
}

// CreateEditionMint is a paid mutator transaction binding the contract method 0xc309c47d.
//
// Solidity: function createEditionMint(address edition, bytes32 merkleRootHash, uint96 price, uint32 startTime, uint32 endTime, uint16 affiliateFeeBPS, uint32 maxMintable, uint32 maxMintablePerAccount) returns(uint128 mintId)
func (_MerkleDropMinter *MerkleDropMinterSession) CreateEditionMint(edition common.Address, merkleRootHash [32]byte, price *big.Int, startTime uint32, endTime uint32, affiliateFeeBPS uint16, maxMintable uint32, maxMintablePerAccount uint32) (*types.Transaction, error) {
	return _MerkleDropMinter.Contract.CreateEditionMint(&_MerkleDropMinter.TransactOpts, edition, merkleRootHash, price, startTime, endTime, affiliateFeeBPS, maxMintable, maxMintablePerAccount)
}

// CreateEditionMint is a paid mutator transaction binding the contract method 0xc309c47d.
//
// Solidity: function createEditionMint(address edition, bytes32 merkleRootHash, uint96 price, uint32 startTime, uint32 endTime, uint16 affiliateFeeBPS, uint32 maxMintable, uint32 maxMintablePerAccount) returns(uint128 mintId)
func (_MerkleDropMinter *MerkleDropMinterTransactorSession) CreateEditionMint(edition common.Address, merkleRootHash [32]byte, price *big.Int, startTime uint32, endTime uint32, affiliateFeeBPS uint16, maxMintable uint32, maxMintablePerAccount uint32) (*types.Transaction, error) {
	return _MerkleDropMinter.Contract.CreateEditionMint(&_MerkleDropMinter.TransactOpts, edition, merkleRootHash, price, startTime, endTime, affiliateFeeBPS, maxMintable, maxMintablePerAccount)
}

// Mint is a paid mutator transaction binding the contract method 0x159a76bd.
//
// Solidity: function mint(address edition, uint128 mintId, uint32 requestedQuantity, bytes32[] merkleProof, address affiliate) payable returns()
func (_MerkleDropMinter *MerkleDropMinterTransactor) Mint(opts *bind.TransactOpts, edition common.Address, mintId *big.Int, requestedQuantity uint32, merkleProof [][32]byte, affiliate common.Address) (*types.Transaction, error) {
	return _MerkleDropMinter.contract.Transact(opts, "mint", edition, mintId, requestedQuantity, merkleProof, affiliate)
}

// Mint is a paid mutator transaction binding the contract method 0x159a76bd.
//
// Solidity: function mint(address edition, uint128 mintId, uint32 requestedQuantity, bytes32[] merkleProof, address affiliate) payable returns()
func (_MerkleDropMinter *MerkleDropMinterSession) Mint(edition common.Address, mintId *big.Int, requestedQuantity uint32, merkleProof [][32]byte, affiliate common.Address) (*types.Transaction, error) {
	return _MerkleDropMinter.Contract.Mint(&_MerkleDropMinter.TransactOpts, edition, mintId, requestedQuantity, merkleProof, affiliate)
}

// Mint is a paid mutator transaction binding the contract method 0x159a76bd.
//
// Solidity: function mint(address edition, uint128 mintId, uint32 requestedQuantity, bytes32[] merkleProof, address affiliate) payable returns()
func (_MerkleDropMinter *MerkleDropMinterTransactorSession) Mint(edition common.Address, mintId *big.Int, requestedQuantity uint32, merkleProof [][32]byte, affiliate common.Address) (*types.Transaction, error) {
	return _MerkleDropMinter.Contract.Mint(&_MerkleDropMinter.TransactOpts, edition, mintId, requestedQuantity, merkleProof, affiliate)
}

// SetAffiliateFee is a paid mutator transaction binding the contract method 0xd73f3aab.
//
// Solidity: function setAffiliateFee(address edition, uint128 mintId, uint16 feeBPS) returns()
func (_MerkleDropMinter *MerkleDropMinterTransactor) SetAffiliateFee(opts *bind.TransactOpts, edition common.Address, mintId *big.Int, feeBPS uint16) (*types.Transaction, error) {
	return _MerkleDropMinter.contract.Transact(opts, "setAffiliateFee", edition, mintId, feeBPS)
}

// SetAffiliateFee is a paid mutator transaction binding the contract method 0xd73f3aab.
//
// Solidity: function setAffiliateFee(address edition, uint128 mintId, uint16 feeBPS) returns()
func (_MerkleDropMinter *MerkleDropMinterSession) SetAffiliateFee(edition common.Address, mintId *big.Int, feeBPS uint16) (*types.Transaction, error) {
	return _MerkleDropMinter.Contract.SetAffiliateFee(&_MerkleDropMinter.TransactOpts, edition, mintId, feeBPS)
}

// SetAffiliateFee is a paid mutator transaction binding the contract method 0xd73f3aab.
//
// Solidity: function setAffiliateFee(address edition, uint128 mintId, uint16 feeBPS) returns()
func (_MerkleDropMinter *MerkleDropMinterTransactorSession) SetAffiliateFee(edition common.Address, mintId *big.Int, feeBPS uint16) (*types.Transaction, error) {
	return _MerkleDropMinter.Contract.SetAffiliateFee(&_MerkleDropMinter.TransactOpts, edition, mintId, feeBPS)
}

// SetEditionMintPaused is a paid mutator transaction binding the contract method 0x522f7386.
//
// Solidity: function setEditionMintPaused(address edition, uint128 mintId, bool paused) returns()
func (_MerkleDropMinter *MerkleDropMinterTransactor) SetEditionMintPaused(opts *bind.TransactOpts, edition common.Address, mintId *big.Int, paused bool) (*types.Transaction, error) {
	return _MerkleDropMinter.contract.Transact(opts, "setEditionMintPaused", edition, mintId, paused)
}

// SetEditionMintPaused is a paid mutator transaction binding the contract method 0x522f7386.
//
// Solidity: function setEditionMintPaused(address edition, uint128 mintId, bool paused) returns()
func (_MerkleDropMinter *MerkleDropMinterSession) SetEditionMintPaused(edition common.Address, mintId *big.Int, paused bool) (*types.Transaction, error) {
	return _MerkleDropMinter.Contract.SetEditionMintPaused(&_MerkleDropMinter.TransactOpts, edition, mintId, paused)
}

// SetEditionMintPaused is a paid mutator transaction binding the contract method 0x522f7386.
//
// Solidity: function setEditionMintPaused(address edition, uint128 mintId, bool paused) returns()
func (_MerkleDropMinter *MerkleDropMinterTransactorSession) SetEditionMintPaused(edition common.Address, mintId *big.Int, paused bool) (*types.Transaction, error) {
	return _MerkleDropMinter.Contract.SetEditionMintPaused(&_MerkleDropMinter.TransactOpts, edition, mintId, paused)
}

// SetMaxMintable is a paid mutator transaction binding the contract method 0x153229bd.
//
// Solidity: function setMaxMintable(address edition, uint128 mintId, uint32 maxMintable) returns()
func (_MerkleDropMinter *MerkleDropMinterTransactor) SetMaxMintable(opts *bind.TransactOpts, edition common.Address, mintId *big.Int, maxMintable uint32) (*types.Transaction, error) {
	return _MerkleDropMinter.contract.Transact(opts, "setMaxMintable", edition, mintId, maxMintable)
}

// SetMaxMintable is a paid mutator transaction binding the contract method 0x153229bd.
//
// Solidity: function setMaxMintable(address edition, uint128 mintId, uint32 maxMintable) returns()
func (_MerkleDropMinter *MerkleDropMinterSession) SetMaxMintable(edition common.Address, mintId *big.Int, maxMintable uint32) (*types.Transaction, error) {
	return _MerkleDropMinter.Contract.SetMaxMintable(&_MerkleDropMinter.TransactOpts, edition, mintId, maxMintable)
}

// SetMaxMintable is a paid mutator transaction binding the contract method 0x153229bd.
//
// Solidity: function setMaxMintable(address edition, uint128 mintId, uint32 maxMintable) returns()
func (_MerkleDropMinter *MerkleDropMinterTransactorSession) SetMaxMintable(edition common.Address, mintId *big.Int, maxMintable uint32) (*types.Transaction, error) {
	return _MerkleDropMinter.Contract.SetMaxMintable(&_MerkleDropMinter.TransactOpts, edition, mintId, maxMintable)
}

// SetMaxMintablePerAccount is a paid mutator transaction binding the contract method 0x7614a751.
//
// Solidity: function setMaxMintablePerAccount(address edition, uint128 mintId, uint32 maxMintablePerAccount) returns()
func (_MerkleDropMinter *MerkleDropMinterTransactor) SetMaxMintablePerAccount(opts *bind.TransactOpts, edition common.Address, mintId *big.Int, maxMintablePerAccount uint32) (*types.Transaction, error) {
	return _MerkleDropMinter.contract.Transact(opts, "setMaxMintablePerAccount", edition, mintId, maxMintablePerAccount)
}

// SetMaxMintablePerAccount is a paid mutator transaction binding the contract method 0x7614a751.
//
// Solidity: function setMaxMintablePerAccount(address edition, uint128 mintId, uint32 maxMintablePerAccount) returns()
func (_MerkleDropMinter *MerkleDropMinterSession) SetMaxMintablePerAccount(edition common.Address, mintId *big.Int, maxMintablePerAccount uint32) (*types.Transaction, error) {
	return _MerkleDropMinter.Contract.SetMaxMintablePerAccount(&_MerkleDropMinter.TransactOpts, edition, mintId, maxMintablePerAccount)
}

// SetMaxMintablePerAccount is a paid mutator transaction binding the contract method 0x7614a751.
//
// Solidity: function setMaxMintablePerAccount(address edition, uint128 mintId, uint32 maxMintablePerAccount) returns()
func (_MerkleDropMinter *MerkleDropMinterTransactorSession) SetMaxMintablePerAccount(edition common.Address, mintId *big.Int, maxMintablePerAccount uint32) (*types.Transaction, error) {
	return _MerkleDropMinter.Contract.SetMaxMintablePerAccount(&_MerkleDropMinter.TransactOpts, edition, mintId, maxMintablePerAccount)
}

// SetMerkleRootHash is a paid mutator transaction binding the contract method 0xc3190a2e.
//
// Solidity: function setMerkleRootHash(address edition, uint128 mintId, bytes32 merkleRootHash) returns()
func (_MerkleDropMinter *MerkleDropMinterTransactor) SetMerkleRootHash(opts *bind.TransactOpts, edition common.Address, mintId *big.Int, merkleRootHash [32]byte) (*types.Transaction, error) {
	return _MerkleDropMinter.contract.Transact(opts, "setMerkleRootHash", edition, mintId, merkleRootHash)
}

// SetMerkleRootHash is a paid mutator transaction binding the contract method 0xc3190a2e.
//
// Solidity: function setMerkleRootHash(address edition, uint128 mintId, bytes32 merkleRootHash) returns()
func (_MerkleDropMinter *MerkleDropMinterSession) SetMerkleRootHash(edition common.Address, mintId *big.Int, merkleRootHash [32]byte) (*types.Transaction, error) {
	return _MerkleDropMinter.Contract.SetMerkleRootHash(&_MerkleDropMinter.TransactOpts, edition, mintId, merkleRootHash)
}

// SetMerkleRootHash is a paid mutator transaction binding the contract method 0xc3190a2e.
//
// Solidity: function setMerkleRootHash(address edition, uint128 mintId, bytes32 merkleRootHash) returns()
func (_MerkleDropMinter *MerkleDropMinterTransactorSession) SetMerkleRootHash(edition common.Address, mintId *big.Int, merkleRootHash [32]byte) (*types.Transaction, error) {
	return _MerkleDropMinter.Contract.SetMerkleRootHash(&_MerkleDropMinter.TransactOpts, edition, mintId, merkleRootHash)
}

// SetPrice is a paid mutator transaction binding the contract method 0xfc63e69e.
//
// Solidity: function setPrice(address edition, uint128 mintId, uint96 price) returns()
func (_MerkleDropMinter *MerkleDropMinterTransactor) SetPrice(opts *bind.TransactOpts, edition common.Address, mintId *big.Int, price *big.Int) (*types.Transaction, error) {
	return _MerkleDropMinter.contract.Transact(opts, "setPrice", edition, mintId, price)
}

// SetPrice is a paid mutator transaction binding the contract method 0xfc63e69e.
//
// Solidity: function setPrice(address edition, uint128 mintId, uint96 price) returns()
func (_MerkleDropMinter *MerkleDropMinterSession) SetPrice(edition common.Address, mintId *big.Int, price *big.Int) (*types.Transaction, error) {
	return _MerkleDropMinter.Contract.SetPrice(&_MerkleDropMinter.TransactOpts, edition, mintId, price)
}

// SetPrice is a paid mutator transaction binding the contract method 0xfc63e69e.
//
// Solidity: function setPrice(address edition, uint128 mintId, uint96 price) returns()
func (_MerkleDropMinter *MerkleDropMinterTransactorSession) SetPrice(edition common.Address, mintId *big.Int, price *big.Int) (*types.Transaction, error) {
	return _MerkleDropMinter.Contract.SetPrice(&_MerkleDropMinter.TransactOpts, edition, mintId, price)
}

// SetTimeRange is a paid mutator transaction binding the contract method 0xaec96b6e.
//
// Solidity: function setTimeRange(address edition, uint128 mintId, uint32 startTime, uint32 endTime) returns()
func (_MerkleDropMinter *MerkleDropMinterTransactor) SetTimeRange(opts *bind.TransactOpts, edition common.Address, mintId *big.Int, startTime uint32, endTime uint32) (*types.Transaction, error) {
	return _MerkleDropMinter.contract.Transact(opts, "setTimeRange", edition, mintId, startTime, endTime)
}

// SetTimeRange is a paid mutator transaction binding the contract method 0xaec96b6e.
//
// Solidity: function setTimeRange(address edition, uint128 mintId, uint32 startTime, uint32 endTime) returns()
func (_MerkleDropMinter *MerkleDropMinterSession) SetTimeRange(edition common.Address, mintId *big.Int, startTime uint32, endTime uint32) (*types.Transaction, error) {
	return _MerkleDropMinter.Contract.SetTimeRange(&_MerkleDropMinter.TransactOpts, edition, mintId, startTime, endTime)
}

// SetTimeRange is a paid mutator transaction binding the contract method 0xaec96b6e.
//
// Solidity: function setTimeRange(address edition, uint128 mintId, uint32 startTime, uint32 endTime) returns()
func (_MerkleDropMinter *MerkleDropMinterTransactorSession) SetTimeRange(edition common.Address, mintId *big.Int, startTime uint32, endTime uint32) (*types.Transaction, error) {
	return _MerkleDropMinter.Contract.SetTimeRange(&_MerkleDropMinter.TransactOpts, edition, mintId, startTime, endTime)
}

// WithdrawForAffiliate is a paid mutator transaction binding the contract method 0xf5740296.
//
// Solidity: function withdrawForAffiliate(address affiliate) returns()
func (_MerkleDropMinter *MerkleDropMinterTransactor) WithdrawForAffiliate(opts *bind.TransactOpts, affiliate common.Address) (*types.Transaction, error) {
	return _MerkleDropMinter.contract.Transact(opts, "withdrawForAffiliate", affiliate)
}

// WithdrawForAffiliate is a paid mutator transaction binding the contract method 0xf5740296.
//
// Solidity: function withdrawForAffiliate(address affiliate) returns()
func (_MerkleDropMinter *MerkleDropMinterSession) WithdrawForAffiliate(affiliate common.Address) (*types.Transaction, error) {
	return _MerkleDropMinter.Contract.WithdrawForAffiliate(&_MerkleDropMinter.TransactOpts, affiliate)
}

// WithdrawForAffiliate is a paid mutator transaction binding the contract method 0xf5740296.
//
// Solidity: function withdrawForAffiliate(address affiliate) returns()
func (_MerkleDropMinter *MerkleDropMinterTransactorSession) WithdrawForAffiliate(affiliate common.Address) (*types.Transaction, error) {
	return _MerkleDropMinter.Contract.WithdrawForAffiliate(&_MerkleDropMinter.TransactOpts, affiliate)
}

// WithdrawForPlatform is a paid mutator transaction binding the contract method 0x08bfa2bf.
//
// Solidity: function withdrawForPlatform() returns()
func (_MerkleDropMinter *MerkleDropMinterTransactor) WithdrawForPlatform(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerkleDropMinter.contract.Transact(opts, "withdrawForPlatform")
}

// WithdrawForPlatform is a paid mutator transaction binding the contract method 0x08bfa2bf.
//
// Solidity: function withdrawForPlatform() returns()
func (_MerkleDropMinter *MerkleDropMinterSession) WithdrawForPlatform() (*types.Transaction, error) {
	return _MerkleDropMinter.Contract.WithdrawForPlatform(&_MerkleDropMinter.TransactOpts)
}

// WithdrawForPlatform is a paid mutator transaction binding the contract method 0x08bfa2bf.
//
// Solidity: function withdrawForPlatform() returns()
func (_MerkleDropMinter *MerkleDropMinterTransactorSession) WithdrawForPlatform() (*types.Transaction, error) {
	return _MerkleDropMinter.Contract.WithdrawForPlatform(&_MerkleDropMinter.TransactOpts)
}

// MerkleDropMinterAffiliateFeeSetIterator is returned from FilterAffiliateFeeSet and is used to iterate over the raw logs and unpacked data for AffiliateFeeSet events raised by the MerkleDropMinter contract.
type MerkleDropMinterAffiliateFeeSetIterator struct {
	Event *MerkleDropMinterAffiliateFeeSet // Event containing the contract specifics and raw log

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
func (it *MerkleDropMinterAffiliateFeeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MerkleDropMinterAffiliateFeeSet)
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
		it.Event = new(MerkleDropMinterAffiliateFeeSet)
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
func (it *MerkleDropMinterAffiliateFeeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MerkleDropMinterAffiliateFeeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MerkleDropMinterAffiliateFeeSet represents a AffiliateFeeSet event raised by the MerkleDropMinter contract.
type MerkleDropMinterAffiliateFeeSet struct {
	Edition common.Address
	MintId  *big.Int
	Bps     uint16
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAffiliateFeeSet is a free log retrieval operation binding the contract event 0xae6d744348a49699fcb91e6a563f2224c0fb27c21053a9218ee827f9d6e698c7.
//
// Solidity: event AffiliateFeeSet(address indexed edition, uint128 indexed mintId, uint16 bps)
func (_MerkleDropMinter *MerkleDropMinterFilterer) FilterAffiliateFeeSet(opts *bind.FilterOpts, edition []common.Address, mintId []*big.Int) (*MerkleDropMinterAffiliateFeeSetIterator, error) {

	var editionRule []interface{}
	for _, editionItem := range edition {
		editionRule = append(editionRule, editionItem)
	}
	var mintIdRule []interface{}
	for _, mintIdItem := range mintId {
		mintIdRule = append(mintIdRule, mintIdItem)
	}

	logs, sub, err := _MerkleDropMinter.contract.FilterLogs(opts, "AffiliateFeeSet", editionRule, mintIdRule)
	if err != nil {
		return nil, err
	}
	return &MerkleDropMinterAffiliateFeeSetIterator{contract: _MerkleDropMinter.contract, event: "AffiliateFeeSet", logs: logs, sub: sub}, nil
}

// WatchAffiliateFeeSet is a free log subscription operation binding the contract event 0xae6d744348a49699fcb91e6a563f2224c0fb27c21053a9218ee827f9d6e698c7.
//
// Solidity: event AffiliateFeeSet(address indexed edition, uint128 indexed mintId, uint16 bps)
func (_MerkleDropMinter *MerkleDropMinterFilterer) WatchAffiliateFeeSet(opts *bind.WatchOpts, sink chan<- *MerkleDropMinterAffiliateFeeSet, edition []common.Address, mintId []*big.Int) (event.Subscription, error) {

	var editionRule []interface{}
	for _, editionItem := range edition {
		editionRule = append(editionRule, editionItem)
	}
	var mintIdRule []interface{}
	for _, mintIdItem := range mintId {
		mintIdRule = append(mintIdRule, mintIdItem)
	}

	logs, sub, err := _MerkleDropMinter.contract.WatchLogs(opts, "AffiliateFeeSet", editionRule, mintIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MerkleDropMinterAffiliateFeeSet)
				if err := _MerkleDropMinter.contract.UnpackLog(event, "AffiliateFeeSet", log); err != nil {
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

// ParseAffiliateFeeSet is a log parse operation binding the contract event 0xae6d744348a49699fcb91e6a563f2224c0fb27c21053a9218ee827f9d6e698c7.
//
// Solidity: event AffiliateFeeSet(address indexed edition, uint128 indexed mintId, uint16 bps)
func (_MerkleDropMinter *MerkleDropMinterFilterer) ParseAffiliateFeeSet(log types.Log) (*MerkleDropMinterAffiliateFeeSet, error) {
	event := new(MerkleDropMinterAffiliateFeeSet)
	if err := _MerkleDropMinter.contract.UnpackLog(event, "AffiliateFeeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MerkleDropMinterDropClaimedIterator is returned from FilterDropClaimed and is used to iterate over the raw logs and unpacked data for DropClaimed events raised by the MerkleDropMinter contract.
type MerkleDropMinterDropClaimedIterator struct {
	Event *MerkleDropMinterDropClaimed // Event containing the contract specifics and raw log

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
func (it *MerkleDropMinterDropClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MerkleDropMinterDropClaimed)
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
		it.Event = new(MerkleDropMinterDropClaimed)
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
func (it *MerkleDropMinterDropClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MerkleDropMinterDropClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MerkleDropMinterDropClaimed represents a DropClaimed event raised by the MerkleDropMinter contract.
type MerkleDropMinterDropClaimed struct {
	Recipient common.Address
	Quantity  uint32
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDropClaimed is a free log retrieval operation binding the contract event 0xbbce3c6646076fce6266a06486d114fd6f4fcdce3df843aadf9c56a2d89721a0.
//
// Solidity: event DropClaimed(address recipient, uint32 quantity)
func (_MerkleDropMinter *MerkleDropMinterFilterer) FilterDropClaimed(opts *bind.FilterOpts) (*MerkleDropMinterDropClaimedIterator, error) {

	logs, sub, err := _MerkleDropMinter.contract.FilterLogs(opts, "DropClaimed")
	if err != nil {
		return nil, err
	}
	return &MerkleDropMinterDropClaimedIterator{contract: _MerkleDropMinter.contract, event: "DropClaimed", logs: logs, sub: sub}, nil
}

// WatchDropClaimed is a free log subscription operation binding the contract event 0xbbce3c6646076fce6266a06486d114fd6f4fcdce3df843aadf9c56a2d89721a0.
//
// Solidity: event DropClaimed(address recipient, uint32 quantity)
func (_MerkleDropMinter *MerkleDropMinterFilterer) WatchDropClaimed(opts *bind.WatchOpts, sink chan<- *MerkleDropMinterDropClaimed) (event.Subscription, error) {

	logs, sub, err := _MerkleDropMinter.contract.WatchLogs(opts, "DropClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MerkleDropMinterDropClaimed)
				if err := _MerkleDropMinter.contract.UnpackLog(event, "DropClaimed", log); err != nil {
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

// ParseDropClaimed is a log parse operation binding the contract event 0xbbce3c6646076fce6266a06486d114fd6f4fcdce3df843aadf9c56a2d89721a0.
//
// Solidity: event DropClaimed(address recipient, uint32 quantity)
func (_MerkleDropMinter *MerkleDropMinterFilterer) ParseDropClaimed(log types.Log) (*MerkleDropMinterDropClaimed, error) {
	event := new(MerkleDropMinterDropClaimed)
	if err := _MerkleDropMinter.contract.UnpackLog(event, "DropClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MerkleDropMinterMaxMintablePerAccountSetIterator is returned from FilterMaxMintablePerAccountSet and is used to iterate over the raw logs and unpacked data for MaxMintablePerAccountSet events raised by the MerkleDropMinter contract.
type MerkleDropMinterMaxMintablePerAccountSetIterator struct {
	Event *MerkleDropMinterMaxMintablePerAccountSet // Event containing the contract specifics and raw log

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
func (it *MerkleDropMinterMaxMintablePerAccountSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MerkleDropMinterMaxMintablePerAccountSet)
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
		it.Event = new(MerkleDropMinterMaxMintablePerAccountSet)
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
func (it *MerkleDropMinterMaxMintablePerAccountSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MerkleDropMinterMaxMintablePerAccountSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MerkleDropMinterMaxMintablePerAccountSet represents a MaxMintablePerAccountSet event raised by the MerkleDropMinter contract.
type MerkleDropMinterMaxMintablePerAccountSet struct {
	Edition               common.Address
	MintId                *big.Int
	MaxMintablePerAccount uint32
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterMaxMintablePerAccountSet is a free log retrieval operation binding the contract event 0x021f3ca4d64a14574ed4fc7115d2c667604e54efb94b6187ed8117d49c6e4755.
//
// Solidity: event MaxMintablePerAccountSet(address indexed edition, uint128 indexed mintId, uint32 maxMintablePerAccount)
func (_MerkleDropMinter *MerkleDropMinterFilterer) FilterMaxMintablePerAccountSet(opts *bind.FilterOpts, edition []common.Address, mintId []*big.Int) (*MerkleDropMinterMaxMintablePerAccountSetIterator, error) {

	var editionRule []interface{}
	for _, editionItem := range edition {
		editionRule = append(editionRule, editionItem)
	}
	var mintIdRule []interface{}
	for _, mintIdItem := range mintId {
		mintIdRule = append(mintIdRule, mintIdItem)
	}

	logs, sub, err := _MerkleDropMinter.contract.FilterLogs(opts, "MaxMintablePerAccountSet", editionRule, mintIdRule)
	if err != nil {
		return nil, err
	}
	return &MerkleDropMinterMaxMintablePerAccountSetIterator{contract: _MerkleDropMinter.contract, event: "MaxMintablePerAccountSet", logs: logs, sub: sub}, nil
}

// WatchMaxMintablePerAccountSet is a free log subscription operation binding the contract event 0x021f3ca4d64a14574ed4fc7115d2c667604e54efb94b6187ed8117d49c6e4755.
//
// Solidity: event MaxMintablePerAccountSet(address indexed edition, uint128 indexed mintId, uint32 maxMintablePerAccount)
func (_MerkleDropMinter *MerkleDropMinterFilterer) WatchMaxMintablePerAccountSet(opts *bind.WatchOpts, sink chan<- *MerkleDropMinterMaxMintablePerAccountSet, edition []common.Address, mintId []*big.Int) (event.Subscription, error) {

	var editionRule []interface{}
	for _, editionItem := range edition {
		editionRule = append(editionRule, editionItem)
	}
	var mintIdRule []interface{}
	for _, mintIdItem := range mintId {
		mintIdRule = append(mintIdRule, mintIdItem)
	}

	logs, sub, err := _MerkleDropMinter.contract.WatchLogs(opts, "MaxMintablePerAccountSet", editionRule, mintIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MerkleDropMinterMaxMintablePerAccountSet)
				if err := _MerkleDropMinter.contract.UnpackLog(event, "MaxMintablePerAccountSet", log); err != nil {
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

// ParseMaxMintablePerAccountSet is a log parse operation binding the contract event 0x021f3ca4d64a14574ed4fc7115d2c667604e54efb94b6187ed8117d49c6e4755.
//
// Solidity: event MaxMintablePerAccountSet(address indexed edition, uint128 indexed mintId, uint32 maxMintablePerAccount)
func (_MerkleDropMinter *MerkleDropMinterFilterer) ParseMaxMintablePerAccountSet(log types.Log) (*MerkleDropMinterMaxMintablePerAccountSet, error) {
	event := new(MerkleDropMinterMaxMintablePerAccountSet)
	if err := _MerkleDropMinter.contract.UnpackLog(event, "MaxMintablePerAccountSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MerkleDropMinterMaxMintableSetIterator is returned from FilterMaxMintableSet and is used to iterate over the raw logs and unpacked data for MaxMintableSet events raised by the MerkleDropMinter contract.
type MerkleDropMinterMaxMintableSetIterator struct {
	Event *MerkleDropMinterMaxMintableSet // Event containing the contract specifics and raw log

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
func (it *MerkleDropMinterMaxMintableSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MerkleDropMinterMaxMintableSet)
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
		it.Event = new(MerkleDropMinterMaxMintableSet)
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
func (it *MerkleDropMinterMaxMintableSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MerkleDropMinterMaxMintableSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MerkleDropMinterMaxMintableSet represents a MaxMintableSet event raised by the MerkleDropMinter contract.
type MerkleDropMinterMaxMintableSet struct {
	Edition     common.Address
	MintId      *big.Int
	MaxMintable uint32
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMaxMintableSet is a free log retrieval operation binding the contract event 0x4c3c29668821111c3ddbcf87e1cd6047dc735ef691d46e09671629cb42cdd948.
//
// Solidity: event MaxMintableSet(address indexed edition, uint128 indexed mintId, uint32 maxMintable)
func (_MerkleDropMinter *MerkleDropMinterFilterer) FilterMaxMintableSet(opts *bind.FilterOpts, edition []common.Address, mintId []*big.Int) (*MerkleDropMinterMaxMintableSetIterator, error) {

	var editionRule []interface{}
	for _, editionItem := range edition {
		editionRule = append(editionRule, editionItem)
	}
	var mintIdRule []interface{}
	for _, mintIdItem := range mintId {
		mintIdRule = append(mintIdRule, mintIdItem)
	}

	logs, sub, err := _MerkleDropMinter.contract.FilterLogs(opts, "MaxMintableSet", editionRule, mintIdRule)
	if err != nil {
		return nil, err
	}
	return &MerkleDropMinterMaxMintableSetIterator{contract: _MerkleDropMinter.contract, event: "MaxMintableSet", logs: logs, sub: sub}, nil
}

// WatchMaxMintableSet is a free log subscription operation binding the contract event 0x4c3c29668821111c3ddbcf87e1cd6047dc735ef691d46e09671629cb42cdd948.
//
// Solidity: event MaxMintableSet(address indexed edition, uint128 indexed mintId, uint32 maxMintable)
func (_MerkleDropMinter *MerkleDropMinterFilterer) WatchMaxMintableSet(opts *bind.WatchOpts, sink chan<- *MerkleDropMinterMaxMintableSet, edition []common.Address, mintId []*big.Int) (event.Subscription, error) {

	var editionRule []interface{}
	for _, editionItem := range edition {
		editionRule = append(editionRule, editionItem)
	}
	var mintIdRule []interface{}
	for _, mintIdItem := range mintId {
		mintIdRule = append(mintIdRule, mintIdItem)
	}

	logs, sub, err := _MerkleDropMinter.contract.WatchLogs(opts, "MaxMintableSet", editionRule, mintIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MerkleDropMinterMaxMintableSet)
				if err := _MerkleDropMinter.contract.UnpackLog(event, "MaxMintableSet", log); err != nil {
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

// ParseMaxMintableSet is a log parse operation binding the contract event 0x4c3c29668821111c3ddbcf87e1cd6047dc735ef691d46e09671629cb42cdd948.
//
// Solidity: event MaxMintableSet(address indexed edition, uint128 indexed mintId, uint32 maxMintable)
func (_MerkleDropMinter *MerkleDropMinterFilterer) ParseMaxMintableSet(log types.Log) (*MerkleDropMinterMaxMintableSet, error) {
	event := new(MerkleDropMinterMaxMintableSet)
	if err := _MerkleDropMinter.contract.UnpackLog(event, "MaxMintableSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MerkleDropMinterMerkleDropMintCreatedIterator is returned from FilterMerkleDropMintCreated and is used to iterate over the raw logs and unpacked data for MerkleDropMintCreated events raised by the MerkleDropMinter contract.
type MerkleDropMinterMerkleDropMintCreatedIterator struct {
	Event *MerkleDropMinterMerkleDropMintCreated // Event containing the contract specifics and raw log

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
func (it *MerkleDropMinterMerkleDropMintCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MerkleDropMinterMerkleDropMintCreated)
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
		it.Event = new(MerkleDropMinterMerkleDropMintCreated)
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
func (it *MerkleDropMinterMerkleDropMintCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MerkleDropMinterMerkleDropMintCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MerkleDropMinterMerkleDropMintCreated represents a MerkleDropMintCreated event raised by the MerkleDropMinter contract.
type MerkleDropMinterMerkleDropMintCreated struct {
	Edition               common.Address
	MintId                *big.Int
	MerkleRootHash        [32]byte
	Price                 *big.Int
	StartTime             uint32
	EndTime               uint32
	AffiliateFeeBPS       uint16
	MaxMintable           uint32
	MaxMintablePerAccount uint32
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterMerkleDropMintCreated is a free log retrieval operation binding the contract event 0xd9faafd9b789bcd20399f1fafa1c6459996ac840e9177ee687c23cdbe3b7a9cb.
//
// Solidity: event MerkleDropMintCreated(address indexed edition, uint128 indexed mintId, bytes32 merkleRootHash, uint96 price, uint32 startTime, uint32 endTime, uint16 affiliateFeeBPS, uint32 maxMintable, uint32 maxMintablePerAccount)
func (_MerkleDropMinter *MerkleDropMinterFilterer) FilterMerkleDropMintCreated(opts *bind.FilterOpts, edition []common.Address, mintId []*big.Int) (*MerkleDropMinterMerkleDropMintCreatedIterator, error) {

	var editionRule []interface{}
	for _, editionItem := range edition {
		editionRule = append(editionRule, editionItem)
	}
	var mintIdRule []interface{}
	for _, mintIdItem := range mintId {
		mintIdRule = append(mintIdRule, mintIdItem)
	}

	logs, sub, err := _MerkleDropMinter.contract.FilterLogs(opts, "MerkleDropMintCreated", editionRule, mintIdRule)
	if err != nil {
		return nil, err
	}
	return &MerkleDropMinterMerkleDropMintCreatedIterator{contract: _MerkleDropMinter.contract, event: "MerkleDropMintCreated", logs: logs, sub: sub}, nil
}

// WatchMerkleDropMintCreated is a free log subscription operation binding the contract event 0xd9faafd9b789bcd20399f1fafa1c6459996ac840e9177ee687c23cdbe3b7a9cb.
//
// Solidity: event MerkleDropMintCreated(address indexed edition, uint128 indexed mintId, bytes32 merkleRootHash, uint96 price, uint32 startTime, uint32 endTime, uint16 affiliateFeeBPS, uint32 maxMintable, uint32 maxMintablePerAccount)
func (_MerkleDropMinter *MerkleDropMinterFilterer) WatchMerkleDropMintCreated(opts *bind.WatchOpts, sink chan<- *MerkleDropMinterMerkleDropMintCreated, edition []common.Address, mintId []*big.Int) (event.Subscription, error) {

	var editionRule []interface{}
	for _, editionItem := range edition {
		editionRule = append(editionRule, editionItem)
	}
	var mintIdRule []interface{}
	for _, mintIdItem := range mintId {
		mintIdRule = append(mintIdRule, mintIdItem)
	}

	logs, sub, err := _MerkleDropMinter.contract.WatchLogs(opts, "MerkleDropMintCreated", editionRule, mintIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MerkleDropMinterMerkleDropMintCreated)
				if err := _MerkleDropMinter.contract.UnpackLog(event, "MerkleDropMintCreated", log); err != nil {
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

// ParseMerkleDropMintCreated is a log parse operation binding the contract event 0xd9faafd9b789bcd20399f1fafa1c6459996ac840e9177ee687c23cdbe3b7a9cb.
//
// Solidity: event MerkleDropMintCreated(address indexed edition, uint128 indexed mintId, bytes32 merkleRootHash, uint96 price, uint32 startTime, uint32 endTime, uint16 affiliateFeeBPS, uint32 maxMintable, uint32 maxMintablePerAccount)
func (_MerkleDropMinter *MerkleDropMinterFilterer) ParseMerkleDropMintCreated(log types.Log) (*MerkleDropMinterMerkleDropMintCreated, error) {
	event := new(MerkleDropMinterMerkleDropMintCreated)
	if err := _MerkleDropMinter.contract.UnpackLog(event, "MerkleDropMintCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MerkleDropMinterMerkleRootHashSetIterator is returned from FilterMerkleRootHashSet and is used to iterate over the raw logs and unpacked data for MerkleRootHashSet events raised by the MerkleDropMinter contract.
type MerkleDropMinterMerkleRootHashSetIterator struct {
	Event *MerkleDropMinterMerkleRootHashSet // Event containing the contract specifics and raw log

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
func (it *MerkleDropMinterMerkleRootHashSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MerkleDropMinterMerkleRootHashSet)
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
		it.Event = new(MerkleDropMinterMerkleRootHashSet)
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
func (it *MerkleDropMinterMerkleRootHashSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MerkleDropMinterMerkleRootHashSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MerkleDropMinterMerkleRootHashSet represents a MerkleRootHashSet event raised by the MerkleDropMinter contract.
type MerkleDropMinterMerkleRootHashSet struct {
	Edition        common.Address
	MintId         *big.Int
	MerkleRootHash [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterMerkleRootHashSet is a free log retrieval operation binding the contract event 0xa863d0ab0f23accb1c42ee00c0587f1c94539e15956a678a54c1e710ba07066e.
//
// Solidity: event MerkleRootHashSet(address indexed edition, uint128 indexed mintId, bytes32 merkleRootHash)
func (_MerkleDropMinter *MerkleDropMinterFilterer) FilterMerkleRootHashSet(opts *bind.FilterOpts, edition []common.Address, mintId []*big.Int) (*MerkleDropMinterMerkleRootHashSetIterator, error) {

	var editionRule []interface{}
	for _, editionItem := range edition {
		editionRule = append(editionRule, editionItem)
	}
	var mintIdRule []interface{}
	for _, mintIdItem := range mintId {
		mintIdRule = append(mintIdRule, mintIdItem)
	}

	logs, sub, err := _MerkleDropMinter.contract.FilterLogs(opts, "MerkleRootHashSet", editionRule, mintIdRule)
	if err != nil {
		return nil, err
	}
	return &MerkleDropMinterMerkleRootHashSetIterator{contract: _MerkleDropMinter.contract, event: "MerkleRootHashSet", logs: logs, sub: sub}, nil
}

// WatchMerkleRootHashSet is a free log subscription operation binding the contract event 0xa863d0ab0f23accb1c42ee00c0587f1c94539e15956a678a54c1e710ba07066e.
//
// Solidity: event MerkleRootHashSet(address indexed edition, uint128 indexed mintId, bytes32 merkleRootHash)
func (_MerkleDropMinter *MerkleDropMinterFilterer) WatchMerkleRootHashSet(opts *bind.WatchOpts, sink chan<- *MerkleDropMinterMerkleRootHashSet, edition []common.Address, mintId []*big.Int) (event.Subscription, error) {

	var editionRule []interface{}
	for _, editionItem := range edition {
		editionRule = append(editionRule, editionItem)
	}
	var mintIdRule []interface{}
	for _, mintIdItem := range mintId {
		mintIdRule = append(mintIdRule, mintIdItem)
	}

	logs, sub, err := _MerkleDropMinter.contract.WatchLogs(opts, "MerkleRootHashSet", editionRule, mintIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MerkleDropMinterMerkleRootHashSet)
				if err := _MerkleDropMinter.contract.UnpackLog(event, "MerkleRootHashSet", log); err != nil {
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

// ParseMerkleRootHashSet is a log parse operation binding the contract event 0xa863d0ab0f23accb1c42ee00c0587f1c94539e15956a678a54c1e710ba07066e.
//
// Solidity: event MerkleRootHashSet(address indexed edition, uint128 indexed mintId, bytes32 merkleRootHash)
func (_MerkleDropMinter *MerkleDropMinterFilterer) ParseMerkleRootHashSet(log types.Log) (*MerkleDropMinterMerkleRootHashSet, error) {
	event := new(MerkleDropMinterMerkleRootHashSet)
	if err := _MerkleDropMinter.contract.UnpackLog(event, "MerkleRootHashSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MerkleDropMinterMintConfigCreatedIterator is returned from FilterMintConfigCreated and is used to iterate over the raw logs and unpacked data for MintConfigCreated events raised by the MerkleDropMinter contract.
type MerkleDropMinterMintConfigCreatedIterator struct {
	Event *MerkleDropMinterMintConfigCreated // Event containing the contract specifics and raw log

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
func (it *MerkleDropMinterMintConfigCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MerkleDropMinterMintConfigCreated)
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
		it.Event = new(MerkleDropMinterMintConfigCreated)
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
func (it *MerkleDropMinterMintConfigCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MerkleDropMinterMintConfigCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MerkleDropMinterMintConfigCreated represents a MintConfigCreated event raised by the MerkleDropMinter contract.
type MerkleDropMinterMintConfigCreated struct {
	Edition         common.Address
	Creator         common.Address
	MintId          *big.Int
	StartTime       uint32
	EndTime         uint32
	AffiliateFeeBPS uint16
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterMintConfigCreated is a free log retrieval operation binding the contract event 0xeeb5941fb79bbfe40edd76c2e086e8ccdb0b463fc8ac07416266100b4dfddccf.
//
// Solidity: event MintConfigCreated(address indexed edition, address indexed creator, uint128 mintId, uint32 startTime, uint32 endTime, uint16 affiliateFeeBPS)
func (_MerkleDropMinter *MerkleDropMinterFilterer) FilterMintConfigCreated(opts *bind.FilterOpts, edition []common.Address, creator []common.Address) (*MerkleDropMinterMintConfigCreatedIterator, error) {

	var editionRule []interface{}
	for _, editionItem := range edition {
		editionRule = append(editionRule, editionItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _MerkleDropMinter.contract.FilterLogs(opts, "MintConfigCreated", editionRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return &MerkleDropMinterMintConfigCreatedIterator{contract: _MerkleDropMinter.contract, event: "MintConfigCreated", logs: logs, sub: sub}, nil
}

// WatchMintConfigCreated is a free log subscription operation binding the contract event 0xeeb5941fb79bbfe40edd76c2e086e8ccdb0b463fc8ac07416266100b4dfddccf.
//
// Solidity: event MintConfigCreated(address indexed edition, address indexed creator, uint128 mintId, uint32 startTime, uint32 endTime, uint16 affiliateFeeBPS)
func (_MerkleDropMinter *MerkleDropMinterFilterer) WatchMintConfigCreated(opts *bind.WatchOpts, sink chan<- *MerkleDropMinterMintConfigCreated, edition []common.Address, creator []common.Address) (event.Subscription, error) {

	var editionRule []interface{}
	for _, editionItem := range edition {
		editionRule = append(editionRule, editionItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _MerkleDropMinter.contract.WatchLogs(opts, "MintConfigCreated", editionRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MerkleDropMinterMintConfigCreated)
				if err := _MerkleDropMinter.contract.UnpackLog(event, "MintConfigCreated", log); err != nil {
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

// ParseMintConfigCreated is a log parse operation binding the contract event 0xeeb5941fb79bbfe40edd76c2e086e8ccdb0b463fc8ac07416266100b4dfddccf.
//
// Solidity: event MintConfigCreated(address indexed edition, address indexed creator, uint128 mintId, uint32 startTime, uint32 endTime, uint16 affiliateFeeBPS)
func (_MerkleDropMinter *MerkleDropMinterFilterer) ParseMintConfigCreated(log types.Log) (*MerkleDropMinterMintConfigCreated, error) {
	event := new(MerkleDropMinterMintConfigCreated)
	if err := _MerkleDropMinter.contract.UnpackLog(event, "MintConfigCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MerkleDropMinterMintPausedSetIterator is returned from FilterMintPausedSet and is used to iterate over the raw logs and unpacked data for MintPausedSet events raised by the MerkleDropMinter contract.
type MerkleDropMinterMintPausedSetIterator struct {
	Event *MerkleDropMinterMintPausedSet // Event containing the contract specifics and raw log

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
func (it *MerkleDropMinterMintPausedSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MerkleDropMinterMintPausedSet)
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
		it.Event = new(MerkleDropMinterMintPausedSet)
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
func (it *MerkleDropMinterMintPausedSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MerkleDropMinterMintPausedSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MerkleDropMinterMintPausedSet represents a MintPausedSet event raised by the MerkleDropMinter contract.
type MerkleDropMinterMintPausedSet struct {
	Edition common.Address
	MintId  *big.Int
	Paused  bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMintPausedSet is a free log retrieval operation binding the contract event 0xe430910c8e4bfa8180b70a0b2aa923b82d6712a1d165e223053ef439a3f86147.
//
// Solidity: event MintPausedSet(address indexed edition, uint128 mintId, bool paused)
func (_MerkleDropMinter *MerkleDropMinterFilterer) FilterMintPausedSet(opts *bind.FilterOpts, edition []common.Address) (*MerkleDropMinterMintPausedSetIterator, error) {

	var editionRule []interface{}
	for _, editionItem := range edition {
		editionRule = append(editionRule, editionItem)
	}

	logs, sub, err := _MerkleDropMinter.contract.FilterLogs(opts, "MintPausedSet", editionRule)
	if err != nil {
		return nil, err
	}
	return &MerkleDropMinterMintPausedSetIterator{contract: _MerkleDropMinter.contract, event: "MintPausedSet", logs: logs, sub: sub}, nil
}

// WatchMintPausedSet is a free log subscription operation binding the contract event 0xe430910c8e4bfa8180b70a0b2aa923b82d6712a1d165e223053ef439a3f86147.
//
// Solidity: event MintPausedSet(address indexed edition, uint128 mintId, bool paused)
func (_MerkleDropMinter *MerkleDropMinterFilterer) WatchMintPausedSet(opts *bind.WatchOpts, sink chan<- *MerkleDropMinterMintPausedSet, edition []common.Address) (event.Subscription, error) {

	var editionRule []interface{}
	for _, editionItem := range edition {
		editionRule = append(editionRule, editionItem)
	}

	logs, sub, err := _MerkleDropMinter.contract.WatchLogs(opts, "MintPausedSet", editionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MerkleDropMinterMintPausedSet)
				if err := _MerkleDropMinter.contract.UnpackLog(event, "MintPausedSet", log); err != nil {
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

// ParseMintPausedSet is a log parse operation binding the contract event 0xe430910c8e4bfa8180b70a0b2aa923b82d6712a1d165e223053ef439a3f86147.
//
// Solidity: event MintPausedSet(address indexed edition, uint128 mintId, bool paused)
func (_MerkleDropMinter *MerkleDropMinterFilterer) ParseMintPausedSet(log types.Log) (*MerkleDropMinterMintPausedSet, error) {
	event := new(MerkleDropMinterMintPausedSet)
	if err := _MerkleDropMinter.contract.UnpackLog(event, "MintPausedSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MerkleDropMinterMintedIterator is returned from FilterMinted and is used to iterate over the raw logs and unpacked data for Minted events raised by the MerkleDropMinter contract.
type MerkleDropMinterMintedIterator struct {
	Event *MerkleDropMinterMinted // Event containing the contract specifics and raw log

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
func (it *MerkleDropMinterMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MerkleDropMinterMinted)
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
		it.Event = new(MerkleDropMinterMinted)
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
func (it *MerkleDropMinterMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MerkleDropMinterMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MerkleDropMinterMinted represents a Minted event raised by the MerkleDropMinter contract.
type MerkleDropMinterMinted struct {
	Edition            common.Address
	MintId             *big.Int
	Buyer              common.Address
	FromTokenId        uint32
	Quantity           uint32
	RequiredEtherValue *big.Int
	PlatformFee        *big.Int
	AffiliateFee       *big.Int
	Affiliate          common.Address
	Affiliated         bool
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterMinted is a free log retrieval operation binding the contract event 0x3d73a0206d94d61b7038abcd0eb766a5de22f9844b38a78449054d19a4f1b58a.
//
// Solidity: event Minted(address indexed edition, uint128 indexed mintId, address indexed buyer, uint32 fromTokenId, uint32 quantity, uint128 requiredEtherValue, uint128 platformFee, uint128 affiliateFee, address affiliate, bool affiliated)
func (_MerkleDropMinter *MerkleDropMinterFilterer) FilterMinted(opts *bind.FilterOpts, edition []common.Address, mintId []*big.Int, buyer []common.Address) (*MerkleDropMinterMintedIterator, error) {

	var editionRule []interface{}
	for _, editionItem := range edition {
		editionRule = append(editionRule, editionItem)
	}
	var mintIdRule []interface{}
	for _, mintIdItem := range mintId {
		mintIdRule = append(mintIdRule, mintIdItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _MerkleDropMinter.contract.FilterLogs(opts, "Minted", editionRule, mintIdRule, buyerRule)
	if err != nil {
		return nil, err
	}
	return &MerkleDropMinterMintedIterator{contract: _MerkleDropMinter.contract, event: "Minted", logs: logs, sub: sub}, nil
}

// WatchMinted is a free log subscription operation binding the contract event 0x3d73a0206d94d61b7038abcd0eb766a5de22f9844b38a78449054d19a4f1b58a.
//
// Solidity: event Minted(address indexed edition, uint128 indexed mintId, address indexed buyer, uint32 fromTokenId, uint32 quantity, uint128 requiredEtherValue, uint128 platformFee, uint128 affiliateFee, address affiliate, bool affiliated)
func (_MerkleDropMinter *MerkleDropMinterFilterer) WatchMinted(opts *bind.WatchOpts, sink chan<- *MerkleDropMinterMinted, edition []common.Address, mintId []*big.Int, buyer []common.Address) (event.Subscription, error) {

	var editionRule []interface{}
	for _, editionItem := range edition {
		editionRule = append(editionRule, editionItem)
	}
	var mintIdRule []interface{}
	for _, mintIdItem := range mintId {
		mintIdRule = append(mintIdRule, mintIdItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _MerkleDropMinter.contract.WatchLogs(opts, "Minted", editionRule, mintIdRule, buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MerkleDropMinterMinted)
				if err := _MerkleDropMinter.contract.UnpackLog(event, "Minted", log); err != nil {
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

// ParseMinted is a log parse operation binding the contract event 0x3d73a0206d94d61b7038abcd0eb766a5de22f9844b38a78449054d19a4f1b58a.
//
// Solidity: event Minted(address indexed edition, uint128 indexed mintId, address indexed buyer, uint32 fromTokenId, uint32 quantity, uint128 requiredEtherValue, uint128 platformFee, uint128 affiliateFee, address affiliate, bool affiliated)
func (_MerkleDropMinter *MerkleDropMinterFilterer) ParseMinted(log types.Log) (*MerkleDropMinterMinted, error) {
	event := new(MerkleDropMinterMinted)
	if err := _MerkleDropMinter.contract.UnpackLog(event, "Minted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MerkleDropMinterPriceSetIterator is returned from FilterPriceSet and is used to iterate over the raw logs and unpacked data for PriceSet events raised by the MerkleDropMinter contract.
type MerkleDropMinterPriceSetIterator struct {
	Event *MerkleDropMinterPriceSet // Event containing the contract specifics and raw log

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
func (it *MerkleDropMinterPriceSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MerkleDropMinterPriceSet)
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
		it.Event = new(MerkleDropMinterPriceSet)
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
func (it *MerkleDropMinterPriceSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MerkleDropMinterPriceSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MerkleDropMinterPriceSet represents a PriceSet event raised by the MerkleDropMinter contract.
type MerkleDropMinterPriceSet struct {
	Edition common.Address
	MintId  *big.Int
	Price   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPriceSet is a free log retrieval operation binding the contract event 0x7c77af090cbae157811da55b6d8ae1e307a85f5aa1dd2f7a13183279a8c2b4b2.
//
// Solidity: event PriceSet(address indexed edition, uint128 indexed mintId, uint96 price)
func (_MerkleDropMinter *MerkleDropMinterFilterer) FilterPriceSet(opts *bind.FilterOpts, edition []common.Address, mintId []*big.Int) (*MerkleDropMinterPriceSetIterator, error) {

	var editionRule []interface{}
	for _, editionItem := range edition {
		editionRule = append(editionRule, editionItem)
	}
	var mintIdRule []interface{}
	for _, mintIdItem := range mintId {
		mintIdRule = append(mintIdRule, mintIdItem)
	}

	logs, sub, err := _MerkleDropMinter.contract.FilterLogs(opts, "PriceSet", editionRule, mintIdRule)
	if err != nil {
		return nil, err
	}
	return &MerkleDropMinterPriceSetIterator{contract: _MerkleDropMinter.contract, event: "PriceSet", logs: logs, sub: sub}, nil
}

// WatchPriceSet is a free log subscription operation binding the contract event 0x7c77af090cbae157811da55b6d8ae1e307a85f5aa1dd2f7a13183279a8c2b4b2.
//
// Solidity: event PriceSet(address indexed edition, uint128 indexed mintId, uint96 price)
func (_MerkleDropMinter *MerkleDropMinterFilterer) WatchPriceSet(opts *bind.WatchOpts, sink chan<- *MerkleDropMinterPriceSet, edition []common.Address, mintId []*big.Int) (event.Subscription, error) {

	var editionRule []interface{}
	for _, editionItem := range edition {
		editionRule = append(editionRule, editionItem)
	}
	var mintIdRule []interface{}
	for _, mintIdItem := range mintId {
		mintIdRule = append(mintIdRule, mintIdItem)
	}

	logs, sub, err := _MerkleDropMinter.contract.WatchLogs(opts, "PriceSet", editionRule, mintIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MerkleDropMinterPriceSet)
				if err := _MerkleDropMinter.contract.UnpackLog(event, "PriceSet", log); err != nil {
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

// ParsePriceSet is a log parse operation binding the contract event 0x7c77af090cbae157811da55b6d8ae1e307a85f5aa1dd2f7a13183279a8c2b4b2.
//
// Solidity: event PriceSet(address indexed edition, uint128 indexed mintId, uint96 price)
func (_MerkleDropMinter *MerkleDropMinterFilterer) ParsePriceSet(log types.Log) (*MerkleDropMinterPriceSet, error) {
	event := new(MerkleDropMinterPriceSet)
	if err := _MerkleDropMinter.contract.UnpackLog(event, "PriceSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MerkleDropMinterTimeRangeSetIterator is returned from FilterTimeRangeSet and is used to iterate over the raw logs and unpacked data for TimeRangeSet events raised by the MerkleDropMinter contract.
type MerkleDropMinterTimeRangeSetIterator struct {
	Event *MerkleDropMinterTimeRangeSet // Event containing the contract specifics and raw log

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
func (it *MerkleDropMinterTimeRangeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MerkleDropMinterTimeRangeSet)
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
		it.Event = new(MerkleDropMinterTimeRangeSet)
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
func (it *MerkleDropMinterTimeRangeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MerkleDropMinterTimeRangeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MerkleDropMinterTimeRangeSet represents a TimeRangeSet event raised by the MerkleDropMinter contract.
type MerkleDropMinterTimeRangeSet struct {
	Edition   common.Address
	MintId    *big.Int
	StartTime uint32
	EndTime   uint32
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTimeRangeSet is a free log retrieval operation binding the contract event 0x13bc03d97cc4e2accb3b8290af069c19619d32a4e9c5219f8580108766fb18fd.
//
// Solidity: event TimeRangeSet(address indexed edition, uint128 indexed mintId, uint32 startTime, uint32 endTime)
func (_MerkleDropMinter *MerkleDropMinterFilterer) FilterTimeRangeSet(opts *bind.FilterOpts, edition []common.Address, mintId []*big.Int) (*MerkleDropMinterTimeRangeSetIterator, error) {

	var editionRule []interface{}
	for _, editionItem := range edition {
		editionRule = append(editionRule, editionItem)
	}
	var mintIdRule []interface{}
	for _, mintIdItem := range mintId {
		mintIdRule = append(mintIdRule, mintIdItem)
	}

	logs, sub, err := _MerkleDropMinter.contract.FilterLogs(opts, "TimeRangeSet", editionRule, mintIdRule)
	if err != nil {
		return nil, err
	}
	return &MerkleDropMinterTimeRangeSetIterator{contract: _MerkleDropMinter.contract, event: "TimeRangeSet", logs: logs, sub: sub}, nil
}

// WatchTimeRangeSet is a free log subscription operation binding the contract event 0x13bc03d97cc4e2accb3b8290af069c19619d32a4e9c5219f8580108766fb18fd.
//
// Solidity: event TimeRangeSet(address indexed edition, uint128 indexed mintId, uint32 startTime, uint32 endTime)
func (_MerkleDropMinter *MerkleDropMinterFilterer) WatchTimeRangeSet(opts *bind.WatchOpts, sink chan<- *MerkleDropMinterTimeRangeSet, edition []common.Address, mintId []*big.Int) (event.Subscription, error) {

	var editionRule []interface{}
	for _, editionItem := range edition {
		editionRule = append(editionRule, editionItem)
	}
	var mintIdRule []interface{}
	for _, mintIdItem := range mintId {
		mintIdRule = append(mintIdRule, mintIdItem)
	}

	logs, sub, err := _MerkleDropMinter.contract.WatchLogs(opts, "TimeRangeSet", editionRule, mintIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MerkleDropMinterTimeRangeSet)
				if err := _MerkleDropMinter.contract.UnpackLog(event, "TimeRangeSet", log); err != nil {
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

// ParseTimeRangeSet is a log parse operation binding the contract event 0x13bc03d97cc4e2accb3b8290af069c19619d32a4e9c5219f8580108766fb18fd.
//
// Solidity: event TimeRangeSet(address indexed edition, uint128 indexed mintId, uint32 startTime, uint32 endTime)
func (_MerkleDropMinter *MerkleDropMinterFilterer) ParseTimeRangeSet(log types.Log) (*MerkleDropMinterTimeRangeSet, error) {
	event := new(MerkleDropMinterTimeRangeSet)
	if err := _MerkleDropMinter.contract.UnpackLog(event, "TimeRangeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
