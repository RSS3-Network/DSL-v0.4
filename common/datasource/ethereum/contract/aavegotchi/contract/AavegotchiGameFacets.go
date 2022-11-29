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

// AavegotchiGameFacetRevenueSharesIO is an auto generated low-level Go binding around an user-defined struct.
type AavegotchiGameFacetRevenueSharesIO struct {
	BurnAddress   common.Address
	DaoAddress    common.Address
	RarityFarming common.Address
	PixelCraft    common.Address
}

// AavegotchiGameFacetTokenIdsWithKinship is an auto generated low-level Go binding around an user-defined struct.
type AavegotchiGameFacetTokenIdsWithKinship struct {
	TokenId        *big.Int
	Kinship        *big.Int
	LastInteracted *big.Int
}

// Haunt is an auto generated low-level Go binding around an user-defined struct.
type Haunt struct {
	HauntMaxSize *big.Int
	PortalPrice  *big.Int
	BodyColor    [3]byte
	TotalCount   *big.Int
}

// PortalAavegotchiTraitsIO is an auto generated low-level Go binding around an user-defined struct.
type PortalAavegotchiTraitsIO struct {
	RandomNumber   *big.Int
	NumericTraits  [6]int16
	CollateralType common.Address
	MinimumStake   *big.Int
}

// GameMetaData contains all meta data concerning the Game contract.
var GameMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ClaimAavegotchi\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_time\",\"type\":\"uint256\"}],\"name\":\"LockAavegotchi\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_oldName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_newName\",\"type\":\"string\"}],\"name\":\"SetAavegotchiName\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_batchId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"SetBatchId\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int16[4]\",\"name\":\"_values\",\"type\":\"int16[4]\"}],\"name\":\"SpendSkillpoints\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_time\",\"type\":\"uint256\"}],\"name\":\"UnLockAavegotchi\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_experience\",\"type\":\"uint256\"}],\"name\":\"aavegotchiLevel\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"level_\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"aavegotchiNameAvailable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"available_\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"availableSkillPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int16[6]\",\"name\":\"_numericTraits\",\"type\":\"int16[6]\"}],\"name\":\"baseRarityScore\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rarityScore_\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_option\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_stakeAmount\",\"type\":\"uint256\"}],\"name\":\"claimAavegotchi\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentHaunt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"hauntId_\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"hauntMaxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"portalPrice\",\"type\":\"uint256\"},{\"internalType\":\"bytes3\",\"name\":\"bodyColor\",\"type\":\"bytes3\"},{\"internalType\":\"uint24\",\"name\":\"totalCount\",\"type\":\"uint24\"}],\"internalType\":\"structHaunt\",\"name\":\"haunt_\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getNumericTraits\",\"outputs\":[{\"internalType\":\"int16[6]\",\"name\":\"numericTraits_\",\"type\":\"int16[6]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ghstAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"contract_\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"interact\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"kinship\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"score_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"modifiedTraitsAndRarityScore\",\"outputs\":[{\"internalType\":\"int16[6]\",\"name\":\"numericTraits_\",\"type\":\"int16[6]\"},{\"internalType\":\"uint256\",\"name\":\"rarityScore_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"portalAavegotchiTraits\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"randomNumber\",\"type\":\"uint256\"},{\"internalType\":\"int16[6]\",\"name\":\"numericTraits\",\"type\":\"int16[6]\"},{\"internalType\":\"address\",\"name\":\"collateralType\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minimumStake\",\"type\":\"uint256\"}],\"internalType\":\"structPortalAavegotchiTraitsIO[10]\",\"name\":\"portalAavegotchiTraits_\",\"type\":\"tuple[10]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int16[6]\",\"name\":\"_numericTraits\",\"type\":\"int16[6]\"}],\"name\":\"rarityMultiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"multiplier_\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"realmInteract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revenueShares\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"burnAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"daoAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rarityFarming\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pixelCraft\",\"type\":\"address\"}],\"internalType\":\"structAavegotchiGameFacet.RevenueSharesIO\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"setAavegotchiName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_realm\",\"type\":\"address\"}],\"name\":\"setRealmAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"int16[4]\",\"name\":\"_values\",\"type\":\"int16[4]\"}],\"name\":\"spendSkillPoints\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_skip\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"all\",\"type\":\"bool\"}],\"name\":\"tokenIdsWithKinship\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"kinship\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastInteracted\",\"type\":\"uint256\"}],\"internalType\":\"structAavegotchiGameFacet.TokenIdsWithKinship[]\",\"name\":\"tokenIdsWithKinship_\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_experience\",\"type\":\"uint256\"}],\"name\":\"xpUntilNextLevel\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"requiredXp_\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// GameABI is the input ABI used to generate the binding from.
// Deprecated: Use GameMetaData.ABI instead.
var GameABI = GameMetaData.ABI

// Game is an auto generated Go binding around an Ethereum contract.
type Game struct {
	GameCaller     // Read-only binding to the contract
	GameTransactor // Write-only binding to the contract
	GameFilterer   // Log filterer for contract events
}

// GameCaller is an auto generated read-only Go binding around an Ethereum contract.
type GameCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GameTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GameTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GameFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GameFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GameSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GameSession struct {
	Contract     *Game             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GameCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GameCallerSession struct {
	Contract *GameCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// GameTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GameTransactorSession struct {
	Contract     *GameTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GameRaw is an auto generated low-level Go binding around an Ethereum contract.
type GameRaw struct {
	Contract *Game // Generic contract binding to access the raw methods on
}

// GameCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GameCallerRaw struct {
	Contract *GameCaller // Generic read-only contract binding to access the raw methods on
}

// GameTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GameTransactorRaw struct {
	Contract *GameTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGame creates a new instance of Game, bound to a specific deployed contract.
func NewGame(address common.Address, backend bind.ContractBackend) (*Game, error) {
	contract, err := bindGame(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Game{GameCaller: GameCaller{contract: contract}, GameTransactor: GameTransactor{contract: contract}, GameFilterer: GameFilterer{contract: contract}}, nil
}

// NewGameCaller creates a new read-only instance of Game, bound to a specific deployed contract.
func NewGameCaller(address common.Address, caller bind.ContractCaller) (*GameCaller, error) {
	contract, err := bindGame(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GameCaller{contract: contract}, nil
}

// NewGameTransactor creates a new write-only instance of Game, bound to a specific deployed contract.
func NewGameTransactor(address common.Address, transactor bind.ContractTransactor) (*GameTransactor, error) {
	contract, err := bindGame(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GameTransactor{contract: contract}, nil
}

// NewGameFilterer creates a new log filterer instance of Game, bound to a specific deployed contract.
func NewGameFilterer(address common.Address, filterer bind.ContractFilterer) (*GameFilterer, error) {
	contract, err := bindGame(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GameFilterer{contract: contract}, nil
}

// bindGame binds a generic wrapper to an already deployed contract.
func bindGame(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GameABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Game *GameRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Game.Contract.GameCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Game *GameRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Game.Contract.GameTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Game *GameRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Game.Contract.GameTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Game *GameCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Game.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Game *GameTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Game.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Game *GameTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Game.Contract.contract.Transact(opts, method, params...)
}

// AavegotchiLevel is a free data retrieval call binding the contract method 0xf98dc1e0.
//
// Solidity: function aavegotchiLevel(uint256 _experience) pure returns(uint256 level_)
func (_Game *GameCaller) AavegotchiLevel(opts *bind.CallOpts, _experience *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Game.contract.Call(opts, &out, "aavegotchiLevel", _experience)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AavegotchiLevel is a free data retrieval call binding the contract method 0xf98dc1e0.
//
// Solidity: function aavegotchiLevel(uint256 _experience) pure returns(uint256 level_)
func (_Game *GameSession) AavegotchiLevel(_experience *big.Int) (*big.Int, error) {
	return _Game.Contract.AavegotchiLevel(&_Game.CallOpts, _experience)
}

// AavegotchiLevel is a free data retrieval call binding the contract method 0xf98dc1e0.
//
// Solidity: function aavegotchiLevel(uint256 _experience) pure returns(uint256 level_)
func (_Game *GameCallerSession) AavegotchiLevel(_experience *big.Int) (*big.Int, error) {
	return _Game.Contract.AavegotchiLevel(&_Game.CallOpts, _experience)
}

// AavegotchiNameAvailable is a free data retrieval call binding the contract method 0xf1567e24.
//
// Solidity: function aavegotchiNameAvailable(string _name) view returns(bool available_)
func (_Game *GameCaller) AavegotchiNameAvailable(opts *bind.CallOpts, _name string) (bool, error) {
	var out []interface{}
	err := _Game.contract.Call(opts, &out, "aavegotchiNameAvailable", _name)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AavegotchiNameAvailable is a free data retrieval call binding the contract method 0xf1567e24.
//
// Solidity: function aavegotchiNameAvailable(string _name) view returns(bool available_)
func (_Game *GameSession) AavegotchiNameAvailable(_name string) (bool, error) {
	return _Game.Contract.AavegotchiNameAvailable(&_Game.CallOpts, _name)
}

// AavegotchiNameAvailable is a free data retrieval call binding the contract method 0xf1567e24.
//
// Solidity: function aavegotchiNameAvailable(string _name) view returns(bool available_)
func (_Game *GameCallerSession) AavegotchiNameAvailable(_name string) (bool, error) {
	return _Game.Contract.AavegotchiNameAvailable(&_Game.CallOpts, _name)
}

// AvailableSkillPoints is a free data retrieval call binding the contract method 0xa4c366fe.
//
// Solidity: function availableSkillPoints(uint256 _tokenId) view returns(uint256)
func (_Game *GameCaller) AvailableSkillPoints(opts *bind.CallOpts, _tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Game.contract.Call(opts, &out, "availableSkillPoints", _tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AvailableSkillPoints is a free data retrieval call binding the contract method 0xa4c366fe.
//
// Solidity: function availableSkillPoints(uint256 _tokenId) view returns(uint256)
func (_Game *GameSession) AvailableSkillPoints(_tokenId *big.Int) (*big.Int, error) {
	return _Game.Contract.AvailableSkillPoints(&_Game.CallOpts, _tokenId)
}

// AvailableSkillPoints is a free data retrieval call binding the contract method 0xa4c366fe.
//
// Solidity: function availableSkillPoints(uint256 _tokenId) view returns(uint256)
func (_Game *GameCallerSession) AvailableSkillPoints(_tokenId *big.Int) (*big.Int, error) {
	return _Game.Contract.AvailableSkillPoints(&_Game.CallOpts, _tokenId)
}

// BaseRarityScore is a free data retrieval call binding the contract method 0x9abf6c51.
//
// Solidity: function baseRarityScore(int16[6] _numericTraits) pure returns(uint256 rarityScore_)
func (_Game *GameCaller) BaseRarityScore(opts *bind.CallOpts, _numericTraits [6]int16) (*big.Int, error) {
	var out []interface{}
	err := _Game.contract.Call(opts, &out, "baseRarityScore", _numericTraits)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseRarityScore is a free data retrieval call binding the contract method 0x9abf6c51.
//
// Solidity: function baseRarityScore(int16[6] _numericTraits) pure returns(uint256 rarityScore_)
func (_Game *GameSession) BaseRarityScore(_numericTraits [6]int16) (*big.Int, error) {
	return _Game.Contract.BaseRarityScore(&_Game.CallOpts, _numericTraits)
}

// BaseRarityScore is a free data retrieval call binding the contract method 0x9abf6c51.
//
// Solidity: function baseRarityScore(int16[6] _numericTraits) pure returns(uint256 rarityScore_)
func (_Game *GameCallerSession) BaseRarityScore(_numericTraits [6]int16) (*big.Int, error) {
	return _Game.Contract.BaseRarityScore(&_Game.CallOpts, _numericTraits)
}

// CurrentHaunt is a free data retrieval call binding the contract method 0x1c2500ff.
//
// Solidity: function currentHaunt() view returns(uint256 hauntId_, (uint256,uint256,bytes3,uint24) haunt_)
func (_Game *GameCaller) CurrentHaunt(opts *bind.CallOpts) (struct {
	HauntId *big.Int
	Haunt   Haunt
}, error) {
	var out []interface{}
	err := _Game.contract.Call(opts, &out, "currentHaunt")

	outstruct := new(struct {
		HauntId *big.Int
		Haunt   Haunt
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.HauntId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Haunt = *abi.ConvertType(out[1], new(Haunt)).(*Haunt)

	return *outstruct, err

}

// CurrentHaunt is a free data retrieval call binding the contract method 0x1c2500ff.
//
// Solidity: function currentHaunt() view returns(uint256 hauntId_, (uint256,uint256,bytes3,uint24) haunt_)
func (_Game *GameSession) CurrentHaunt() (struct {
	HauntId *big.Int
	Haunt   Haunt
}, error) {
	return _Game.Contract.CurrentHaunt(&_Game.CallOpts)
}

// CurrentHaunt is a free data retrieval call binding the contract method 0x1c2500ff.
//
// Solidity: function currentHaunt() view returns(uint256 hauntId_, (uint256,uint256,bytes3,uint24) haunt_)
func (_Game *GameCallerSession) CurrentHaunt() (struct {
	HauntId *big.Int
	Haunt   Haunt
}, error) {
	return _Game.Contract.CurrentHaunt(&_Game.CallOpts)
}

// GetNumericTraits is a free data retrieval call binding the contract method 0xfa5f5295.
//
// Solidity: function getNumericTraits(uint256 _tokenId) view returns(int16[6] numericTraits_)
func (_Game *GameCaller) GetNumericTraits(opts *bind.CallOpts, _tokenId *big.Int) ([6]int16, error) {
	var out []interface{}
	err := _Game.contract.Call(opts, &out, "getNumericTraits", _tokenId)

	if err != nil {
		return *new([6]int16), err
	}

	out0 := *abi.ConvertType(out[0], new([6]int16)).(*[6]int16)

	return out0, err

}

// GetNumericTraits is a free data retrieval call binding the contract method 0xfa5f5295.
//
// Solidity: function getNumericTraits(uint256 _tokenId) view returns(int16[6] numericTraits_)
func (_Game *GameSession) GetNumericTraits(_tokenId *big.Int) ([6]int16, error) {
	return _Game.Contract.GetNumericTraits(&_Game.CallOpts, _tokenId)
}

// GetNumericTraits is a free data retrieval call binding the contract method 0xfa5f5295.
//
// Solidity: function getNumericTraits(uint256 _tokenId) view returns(int16[6] numericTraits_)
func (_Game *GameCallerSession) GetNumericTraits(_tokenId *big.Int) ([6]int16, error) {
	return _Game.Contract.GetNumericTraits(&_Game.CallOpts, _tokenId)
}

// GhstAddress is a free data retrieval call binding the contract method 0xa70135c2.
//
// Solidity: function ghstAddress() view returns(address contract_)
func (_Game *GameCaller) GhstAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Game.contract.Call(opts, &out, "ghstAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GhstAddress is a free data retrieval call binding the contract method 0xa70135c2.
//
// Solidity: function ghstAddress() view returns(address contract_)
func (_Game *GameSession) GhstAddress() (common.Address, error) {
	return _Game.Contract.GhstAddress(&_Game.CallOpts)
}

// GhstAddress is a free data retrieval call binding the contract method 0xa70135c2.
//
// Solidity: function ghstAddress() view returns(address contract_)
func (_Game *GameCallerSession) GhstAddress() (common.Address, error) {
	return _Game.Contract.GhstAddress(&_Game.CallOpts)
}

// Kinship is a free data retrieval call binding the contract method 0xf5b91852.
//
// Solidity: function kinship(uint256 _tokenId) view returns(uint256 score_)
func (_Game *GameCaller) Kinship(opts *bind.CallOpts, _tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Game.contract.Call(opts, &out, "kinship", _tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Kinship is a free data retrieval call binding the contract method 0xf5b91852.
//
// Solidity: function kinship(uint256 _tokenId) view returns(uint256 score_)
func (_Game *GameSession) Kinship(_tokenId *big.Int) (*big.Int, error) {
	return _Game.Contract.Kinship(&_Game.CallOpts, _tokenId)
}

// Kinship is a free data retrieval call binding the contract method 0xf5b91852.
//
// Solidity: function kinship(uint256 _tokenId) view returns(uint256 score_)
func (_Game *GameCallerSession) Kinship(_tokenId *big.Int) (*big.Int, error) {
	return _Game.Contract.Kinship(&_Game.CallOpts, _tokenId)
}

// ModifiedTraitsAndRarityScore is a free data retrieval call binding the contract method 0xfffebb41.
//
// Solidity: function modifiedTraitsAndRarityScore(uint256 _tokenId) view returns(int16[6] numericTraits_, uint256 rarityScore_)
func (_Game *GameCaller) ModifiedTraitsAndRarityScore(opts *bind.CallOpts, _tokenId *big.Int) (struct {
	NumericTraits [6]int16
	RarityScore   *big.Int
}, error) {
	var out []interface{}
	err := _Game.contract.Call(opts, &out, "modifiedTraitsAndRarityScore", _tokenId)

	outstruct := new(struct {
		NumericTraits [6]int16
		RarityScore   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NumericTraits = *abi.ConvertType(out[0], new([6]int16)).(*[6]int16)
	outstruct.RarityScore = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ModifiedTraitsAndRarityScore is a free data retrieval call binding the contract method 0xfffebb41.
//
// Solidity: function modifiedTraitsAndRarityScore(uint256 _tokenId) view returns(int16[6] numericTraits_, uint256 rarityScore_)
func (_Game *GameSession) ModifiedTraitsAndRarityScore(_tokenId *big.Int) (struct {
	NumericTraits [6]int16
	RarityScore   *big.Int
}, error) {
	return _Game.Contract.ModifiedTraitsAndRarityScore(&_Game.CallOpts, _tokenId)
}

// ModifiedTraitsAndRarityScore is a free data retrieval call binding the contract method 0xfffebb41.
//
// Solidity: function modifiedTraitsAndRarityScore(uint256 _tokenId) view returns(int16[6] numericTraits_, uint256 rarityScore_)
func (_Game *GameCallerSession) ModifiedTraitsAndRarityScore(_tokenId *big.Int) (struct {
	NumericTraits [6]int16
	RarityScore   *big.Int
}, error) {
	return _Game.Contract.ModifiedTraitsAndRarityScore(&_Game.CallOpts, _tokenId)
}

// PortalAavegotchiTraits is a free data retrieval call binding the contract method 0x28e8c193.
//
// Solidity: function portalAavegotchiTraits(uint256 _tokenId) view returns((uint256,int16[6],address,uint256)[10] portalAavegotchiTraits_)
func (_Game *GameCaller) PortalAavegotchiTraits(opts *bind.CallOpts, _tokenId *big.Int) ([10]PortalAavegotchiTraitsIO, error) {
	var out []interface{}
	err := _Game.contract.Call(opts, &out, "portalAavegotchiTraits", _tokenId)

	if err != nil {
		return *new([10]PortalAavegotchiTraitsIO), err
	}

	out0 := *abi.ConvertType(out[0], new([10]PortalAavegotchiTraitsIO)).(*[10]PortalAavegotchiTraitsIO)

	return out0, err

}

// PortalAavegotchiTraits is a free data retrieval call binding the contract method 0x28e8c193.
//
// Solidity: function portalAavegotchiTraits(uint256 _tokenId) view returns((uint256,int16[6],address,uint256)[10] portalAavegotchiTraits_)
func (_Game *GameSession) PortalAavegotchiTraits(_tokenId *big.Int) ([10]PortalAavegotchiTraitsIO, error) {
	return _Game.Contract.PortalAavegotchiTraits(&_Game.CallOpts, _tokenId)
}

// PortalAavegotchiTraits is a free data retrieval call binding the contract method 0x28e8c193.
//
// Solidity: function portalAavegotchiTraits(uint256 _tokenId) view returns((uint256,int16[6],address,uint256)[10] portalAavegotchiTraits_)
func (_Game *GameCallerSession) PortalAavegotchiTraits(_tokenId *big.Int) ([10]PortalAavegotchiTraitsIO, error) {
	return _Game.Contract.PortalAavegotchiTraits(&_Game.CallOpts, _tokenId)
}

// RarityMultiplier is a free data retrieval call binding the contract method 0x09cbaf06.
//
// Solidity: function rarityMultiplier(int16[6] _numericTraits) pure returns(uint256 multiplier_)
func (_Game *GameCaller) RarityMultiplier(opts *bind.CallOpts, _numericTraits [6]int16) (*big.Int, error) {
	var out []interface{}
	err := _Game.contract.Call(opts, &out, "rarityMultiplier", _numericTraits)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RarityMultiplier is a free data retrieval call binding the contract method 0x09cbaf06.
//
// Solidity: function rarityMultiplier(int16[6] _numericTraits) pure returns(uint256 multiplier_)
func (_Game *GameSession) RarityMultiplier(_numericTraits [6]int16) (*big.Int, error) {
	return _Game.Contract.RarityMultiplier(&_Game.CallOpts, _numericTraits)
}

// RarityMultiplier is a free data retrieval call binding the contract method 0x09cbaf06.
//
// Solidity: function rarityMultiplier(int16[6] _numericTraits) pure returns(uint256 multiplier_)
func (_Game *GameCallerSession) RarityMultiplier(_numericTraits [6]int16) (*big.Int, error) {
	return _Game.Contract.RarityMultiplier(&_Game.CallOpts, _numericTraits)
}

// RevenueShares is a free data retrieval call binding the contract method 0xf6f518c6.
//
// Solidity: function revenueShares() view returns((address,address,address,address))
func (_Game *GameCaller) RevenueShares(opts *bind.CallOpts) (AavegotchiGameFacetRevenueSharesIO, error) {
	var out []interface{}
	err := _Game.contract.Call(opts, &out, "revenueShares")

	if err != nil {
		return *new(AavegotchiGameFacetRevenueSharesIO), err
	}

	out0 := *abi.ConvertType(out[0], new(AavegotchiGameFacetRevenueSharesIO)).(*AavegotchiGameFacetRevenueSharesIO)

	return out0, err

}

// RevenueShares is a free data retrieval call binding the contract method 0xf6f518c6.
//
// Solidity: function revenueShares() view returns((address,address,address,address))
func (_Game *GameSession) RevenueShares() (AavegotchiGameFacetRevenueSharesIO, error) {
	return _Game.Contract.RevenueShares(&_Game.CallOpts)
}

// RevenueShares is a free data retrieval call binding the contract method 0xf6f518c6.
//
// Solidity: function revenueShares() view returns((address,address,address,address))
func (_Game *GameCallerSession) RevenueShares() (AavegotchiGameFacetRevenueSharesIO, error) {
	return _Game.Contract.RevenueShares(&_Game.CallOpts)
}

// TokenIdsWithKinship is a free data retrieval call binding the contract method 0x1436d4b1.
//
// Solidity: function tokenIdsWithKinship(address _owner, uint256 _count, uint256 _skip, bool all) view returns((uint256,uint256,uint256)[] tokenIdsWithKinship_)
func (_Game *GameCaller) TokenIdsWithKinship(opts *bind.CallOpts, _owner common.Address, _count *big.Int, _skip *big.Int, all bool) ([]AavegotchiGameFacetTokenIdsWithKinship, error) {
	var out []interface{}
	err := _Game.contract.Call(opts, &out, "tokenIdsWithKinship", _owner, _count, _skip, all)

	if err != nil {
		return *new([]AavegotchiGameFacetTokenIdsWithKinship), err
	}

	out0 := *abi.ConvertType(out[0], new([]AavegotchiGameFacetTokenIdsWithKinship)).(*[]AavegotchiGameFacetTokenIdsWithKinship)

	return out0, err

}

// TokenIdsWithKinship is a free data retrieval call binding the contract method 0x1436d4b1.
//
// Solidity: function tokenIdsWithKinship(address _owner, uint256 _count, uint256 _skip, bool all) view returns((uint256,uint256,uint256)[] tokenIdsWithKinship_)
func (_Game *GameSession) TokenIdsWithKinship(_owner common.Address, _count *big.Int, _skip *big.Int, all bool) ([]AavegotchiGameFacetTokenIdsWithKinship, error) {
	return _Game.Contract.TokenIdsWithKinship(&_Game.CallOpts, _owner, _count, _skip, all)
}

// TokenIdsWithKinship is a free data retrieval call binding the contract method 0x1436d4b1.
//
// Solidity: function tokenIdsWithKinship(address _owner, uint256 _count, uint256 _skip, bool all) view returns((uint256,uint256,uint256)[] tokenIdsWithKinship_)
func (_Game *GameCallerSession) TokenIdsWithKinship(_owner common.Address, _count *big.Int, _skip *big.Int, all bool) ([]AavegotchiGameFacetTokenIdsWithKinship, error) {
	return _Game.Contract.TokenIdsWithKinship(&_Game.CallOpts, _owner, _count, _skip, all)
}

// XpUntilNextLevel is a free data retrieval call binding the contract method 0x82401ccd.
//
// Solidity: function xpUntilNextLevel(uint256 _experience) pure returns(uint256 requiredXp_)
func (_Game *GameCaller) XpUntilNextLevel(opts *bind.CallOpts, _experience *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Game.contract.Call(opts, &out, "xpUntilNextLevel", _experience)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// XpUntilNextLevel is a free data retrieval call binding the contract method 0x82401ccd.
//
// Solidity: function xpUntilNextLevel(uint256 _experience) pure returns(uint256 requiredXp_)
func (_Game *GameSession) XpUntilNextLevel(_experience *big.Int) (*big.Int, error) {
	return _Game.Contract.XpUntilNextLevel(&_Game.CallOpts, _experience)
}

// XpUntilNextLevel is a free data retrieval call binding the contract method 0x82401ccd.
//
// Solidity: function xpUntilNextLevel(uint256 _experience) pure returns(uint256 requiredXp_)
func (_Game *GameCallerSession) XpUntilNextLevel(_experience *big.Int) (*big.Int, error) {
	return _Game.Contract.XpUntilNextLevel(&_Game.CallOpts, _experience)
}

// ClaimAavegotchi is a paid mutator transaction binding the contract method 0x4984b01c.
//
// Solidity: function claimAavegotchi(uint256 _tokenId, uint256 _option, uint256 _stakeAmount) returns()
func (_Game *GameTransactor) ClaimAavegotchi(opts *bind.TransactOpts, _tokenId *big.Int, _option *big.Int, _stakeAmount *big.Int) (*types.Transaction, error) {
	return _Game.contract.Transact(opts, "claimAavegotchi", _tokenId, _option, _stakeAmount)
}

// ClaimAavegotchi is a paid mutator transaction binding the contract method 0x4984b01c.
//
// Solidity: function claimAavegotchi(uint256 _tokenId, uint256 _option, uint256 _stakeAmount) returns()
func (_Game *GameSession) ClaimAavegotchi(_tokenId *big.Int, _option *big.Int, _stakeAmount *big.Int) (*types.Transaction, error) {
	return _Game.Contract.ClaimAavegotchi(&_Game.TransactOpts, _tokenId, _option, _stakeAmount)
}

// ClaimAavegotchi is a paid mutator transaction binding the contract method 0x4984b01c.
//
// Solidity: function claimAavegotchi(uint256 _tokenId, uint256 _option, uint256 _stakeAmount) returns()
func (_Game *GameTransactorSession) ClaimAavegotchi(_tokenId *big.Int, _option *big.Int, _stakeAmount *big.Int) (*types.Transaction, error) {
	return _Game.Contract.ClaimAavegotchi(&_Game.TransactOpts, _tokenId, _option, _stakeAmount)
}

// Interact is a paid mutator transaction binding the contract method 0x22c67519.
//
// Solidity: function interact(uint256[] _tokenIds) returns()
func (_Game *GameTransactor) Interact(opts *bind.TransactOpts, _tokenIds []*big.Int) (*types.Transaction, error) {
	return _Game.contract.Transact(opts, "interact", _tokenIds)
}

// Interact is a paid mutator transaction binding the contract method 0x22c67519.
//
// Solidity: function interact(uint256[] _tokenIds) returns()
func (_Game *GameSession) Interact(_tokenIds []*big.Int) (*types.Transaction, error) {
	return _Game.Contract.Interact(&_Game.TransactOpts, _tokenIds)
}

// Interact is a paid mutator transaction binding the contract method 0x22c67519.
//
// Solidity: function interact(uint256[] _tokenIds) returns()
func (_Game *GameTransactorSession) Interact(_tokenIds []*big.Int) (*types.Transaction, error) {
	return _Game.Contract.Interact(&_Game.TransactOpts, _tokenIds)
}

// RealmInteract is a paid mutator transaction binding the contract method 0x58f514aa.
//
// Solidity: function realmInteract(uint256 _tokenId) returns()
func (_Game *GameTransactor) RealmInteract(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _Game.contract.Transact(opts, "realmInteract", _tokenId)
}

// RealmInteract is a paid mutator transaction binding the contract method 0x58f514aa.
//
// Solidity: function realmInteract(uint256 _tokenId) returns()
func (_Game *GameSession) RealmInteract(_tokenId *big.Int) (*types.Transaction, error) {
	return _Game.Contract.RealmInteract(&_Game.TransactOpts, _tokenId)
}

// RealmInteract is a paid mutator transaction binding the contract method 0x58f514aa.
//
// Solidity: function realmInteract(uint256 _tokenId) returns()
func (_Game *GameTransactorSession) RealmInteract(_tokenId *big.Int) (*types.Transaction, error) {
	return _Game.Contract.RealmInteract(&_Game.TransactOpts, _tokenId)
}

// SetAavegotchiName is a paid mutator transaction binding the contract method 0xea7dac7f.
//
// Solidity: function setAavegotchiName(uint256 _tokenId, string _name) returns()
func (_Game *GameTransactor) SetAavegotchiName(opts *bind.TransactOpts, _tokenId *big.Int, _name string) (*types.Transaction, error) {
	return _Game.contract.Transact(opts, "setAavegotchiName", _tokenId, _name)
}

// SetAavegotchiName is a paid mutator transaction binding the contract method 0xea7dac7f.
//
// Solidity: function setAavegotchiName(uint256 _tokenId, string _name) returns()
func (_Game *GameSession) SetAavegotchiName(_tokenId *big.Int, _name string) (*types.Transaction, error) {
	return _Game.Contract.SetAavegotchiName(&_Game.TransactOpts, _tokenId, _name)
}

// SetAavegotchiName is a paid mutator transaction binding the contract method 0xea7dac7f.
//
// Solidity: function setAavegotchiName(uint256 _tokenId, string _name) returns()
func (_Game *GameTransactorSession) SetAavegotchiName(_tokenId *big.Int, _name string) (*types.Transaction, error) {
	return _Game.Contract.SetAavegotchiName(&_Game.TransactOpts, _tokenId, _name)
}

// SetRealmAddress is a paid mutator transaction binding the contract method 0x874ca042.
//
// Solidity: function setRealmAddress(address _realm) returns()
func (_Game *GameTransactor) SetRealmAddress(opts *bind.TransactOpts, _realm common.Address) (*types.Transaction, error) {
	return _Game.contract.Transact(opts, "setRealmAddress", _realm)
}

// SetRealmAddress is a paid mutator transaction binding the contract method 0x874ca042.
//
// Solidity: function setRealmAddress(address _realm) returns()
func (_Game *GameSession) SetRealmAddress(_realm common.Address) (*types.Transaction, error) {
	return _Game.Contract.SetRealmAddress(&_Game.TransactOpts, _realm)
}

// SetRealmAddress is a paid mutator transaction binding the contract method 0x874ca042.
//
// Solidity: function setRealmAddress(address _realm) returns()
func (_Game *GameTransactorSession) SetRealmAddress(_realm common.Address) (*types.Transaction, error) {
	return _Game.Contract.SetRealmAddress(&_Game.TransactOpts, _realm)
}

// SpendSkillPoints is a paid mutator transaction binding the contract method 0xc889dfc4.
//
// Solidity: function spendSkillPoints(uint256 _tokenId, int16[4] _values) returns()
func (_Game *GameTransactor) SpendSkillPoints(opts *bind.TransactOpts, _tokenId *big.Int, _values [4]int16) (*types.Transaction, error) {
	return _Game.contract.Transact(opts, "spendSkillPoints", _tokenId, _values)
}

// SpendSkillPoints is a paid mutator transaction binding the contract method 0xc889dfc4.
//
// Solidity: function spendSkillPoints(uint256 _tokenId, int16[4] _values) returns()
func (_Game *GameSession) SpendSkillPoints(_tokenId *big.Int, _values [4]int16) (*types.Transaction, error) {
	return _Game.Contract.SpendSkillPoints(&_Game.TransactOpts, _tokenId, _values)
}

// SpendSkillPoints is a paid mutator transaction binding the contract method 0xc889dfc4.
//
// Solidity: function spendSkillPoints(uint256 _tokenId, int16[4] _values) returns()
func (_Game *GameTransactorSession) SpendSkillPoints(_tokenId *big.Int, _values [4]int16) (*types.Transaction, error) {
	return _Game.Contract.SpendSkillPoints(&_Game.TransactOpts, _tokenId, _values)
}

// GameClaimAavegotchiIterator is returned from FilterClaimAavegotchi and is used to iterate over the raw logs and unpacked data for ClaimAavegotchi events raised by the Game contract.
type GameClaimAavegotchiIterator struct {
	Event *GameClaimAavegotchi // Event containing the contract specifics and raw log

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
func (it *GameClaimAavegotchiIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameClaimAavegotchi)
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
		it.Event = new(GameClaimAavegotchi)
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
func (it *GameClaimAavegotchiIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameClaimAavegotchiIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameClaimAavegotchi represents a ClaimAavegotchi event raised by the Game contract.
type GameClaimAavegotchi struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterClaimAavegotchi is a free log retrieval operation binding the contract event 0xadbf78ef322c0f286c9f0d0812f53f6db59fcefc4e41150758f43da97e2a80e3.
//
// Solidity: event ClaimAavegotchi(uint256 indexed _tokenId)
func (_Game *GameFilterer) FilterClaimAavegotchi(opts *bind.FilterOpts, _tokenId []*big.Int) (*GameClaimAavegotchiIterator, error) {

	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _Game.contract.FilterLogs(opts, "ClaimAavegotchi", _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &GameClaimAavegotchiIterator{contract: _Game.contract, event: "ClaimAavegotchi", logs: logs, sub: sub}, nil
}

// WatchClaimAavegotchi is a free log subscription operation binding the contract event 0xadbf78ef322c0f286c9f0d0812f53f6db59fcefc4e41150758f43da97e2a80e3.
//
// Solidity: event ClaimAavegotchi(uint256 indexed _tokenId)
func (_Game *GameFilterer) WatchClaimAavegotchi(opts *bind.WatchOpts, sink chan<- *GameClaimAavegotchi, _tokenId []*big.Int) (event.Subscription, error) {

	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _Game.contract.WatchLogs(opts, "ClaimAavegotchi", _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameClaimAavegotchi)
				if err := _Game.contract.UnpackLog(event, "ClaimAavegotchi", log); err != nil {
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

// ParseClaimAavegotchi is a log parse operation binding the contract event 0xadbf78ef322c0f286c9f0d0812f53f6db59fcefc4e41150758f43da97e2a80e3.
//
// Solidity: event ClaimAavegotchi(uint256 indexed _tokenId)
func (_Game *GameFilterer) ParseClaimAavegotchi(log types.Log) (*GameClaimAavegotchi, error) {
	event := new(GameClaimAavegotchi)
	if err := _Game.contract.UnpackLog(event, "ClaimAavegotchi", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GameLockAavegotchiIterator is returned from FilterLockAavegotchi and is used to iterate over the raw logs and unpacked data for LockAavegotchi events raised by the Game contract.
type GameLockAavegotchiIterator struct {
	Event *GameLockAavegotchi // Event containing the contract specifics and raw log

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
func (it *GameLockAavegotchiIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameLockAavegotchi)
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
		it.Event = new(GameLockAavegotchi)
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
func (it *GameLockAavegotchiIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameLockAavegotchiIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameLockAavegotchi represents a LockAavegotchi event raised by the Game contract.
type GameLockAavegotchi struct {
	TokenId *big.Int
	Time    *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLockAavegotchi is a free log retrieval operation binding the contract event 0x4cfe9c77ddb88998750256279a4a694b125d4aa88449f4654809e60ae6c01bac.
//
// Solidity: event LockAavegotchi(uint256 indexed _tokenId, uint256 _time)
func (_Game *GameFilterer) FilterLockAavegotchi(opts *bind.FilterOpts, _tokenId []*big.Int) (*GameLockAavegotchiIterator, error) {

	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _Game.contract.FilterLogs(opts, "LockAavegotchi", _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &GameLockAavegotchiIterator{contract: _Game.contract, event: "LockAavegotchi", logs: logs, sub: sub}, nil
}

// WatchLockAavegotchi is a free log subscription operation binding the contract event 0x4cfe9c77ddb88998750256279a4a694b125d4aa88449f4654809e60ae6c01bac.
//
// Solidity: event LockAavegotchi(uint256 indexed _tokenId, uint256 _time)
func (_Game *GameFilterer) WatchLockAavegotchi(opts *bind.WatchOpts, sink chan<- *GameLockAavegotchi, _tokenId []*big.Int) (event.Subscription, error) {

	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _Game.contract.WatchLogs(opts, "LockAavegotchi", _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameLockAavegotchi)
				if err := _Game.contract.UnpackLog(event, "LockAavegotchi", log); err != nil {
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

// ParseLockAavegotchi is a log parse operation binding the contract event 0x4cfe9c77ddb88998750256279a4a694b125d4aa88449f4654809e60ae6c01bac.
//
// Solidity: event LockAavegotchi(uint256 indexed _tokenId, uint256 _time)
func (_Game *GameFilterer) ParseLockAavegotchi(log types.Log) (*GameLockAavegotchi, error) {
	event := new(GameLockAavegotchi)
	if err := _Game.contract.UnpackLog(event, "LockAavegotchi", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GameSetAavegotchiNameIterator is returned from FilterSetAavegotchiName and is used to iterate over the raw logs and unpacked data for SetAavegotchiName events raised by the Game contract.
type GameSetAavegotchiNameIterator struct {
	Event *GameSetAavegotchiName // Event containing the contract specifics and raw log

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
func (it *GameSetAavegotchiNameIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameSetAavegotchiName)
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
		it.Event = new(GameSetAavegotchiName)
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
func (it *GameSetAavegotchiNameIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameSetAavegotchiNameIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameSetAavegotchiName represents a SetAavegotchiName event raised by the Game contract.
type GameSetAavegotchiName struct {
	TokenId *big.Int
	OldName string
	NewName string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSetAavegotchiName is a free log retrieval operation binding the contract event 0x9e94ae5a5d8c4dc3600d749ea1770ee468d1da35e7a5971eb3a889b0d0242504.
//
// Solidity: event SetAavegotchiName(uint256 indexed _tokenId, string _oldName, string _newName)
func (_Game *GameFilterer) FilterSetAavegotchiName(opts *bind.FilterOpts, _tokenId []*big.Int) (*GameSetAavegotchiNameIterator, error) {

	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _Game.contract.FilterLogs(opts, "SetAavegotchiName", _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &GameSetAavegotchiNameIterator{contract: _Game.contract, event: "SetAavegotchiName", logs: logs, sub: sub}, nil
}

// WatchSetAavegotchiName is a free log subscription operation binding the contract event 0x9e94ae5a5d8c4dc3600d749ea1770ee468d1da35e7a5971eb3a889b0d0242504.
//
// Solidity: event SetAavegotchiName(uint256 indexed _tokenId, string _oldName, string _newName)
func (_Game *GameFilterer) WatchSetAavegotchiName(opts *bind.WatchOpts, sink chan<- *GameSetAavegotchiName, _tokenId []*big.Int) (event.Subscription, error) {

	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _Game.contract.WatchLogs(opts, "SetAavegotchiName", _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameSetAavegotchiName)
				if err := _Game.contract.UnpackLog(event, "SetAavegotchiName", log); err != nil {
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

// ParseSetAavegotchiName is a log parse operation binding the contract event 0x9e94ae5a5d8c4dc3600d749ea1770ee468d1da35e7a5971eb3a889b0d0242504.
//
// Solidity: event SetAavegotchiName(uint256 indexed _tokenId, string _oldName, string _newName)
func (_Game *GameFilterer) ParseSetAavegotchiName(log types.Log) (*GameSetAavegotchiName, error) {
	event := new(GameSetAavegotchiName)
	if err := _Game.contract.UnpackLog(event, "SetAavegotchiName", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GameSetBatchIdIterator is returned from FilterSetBatchId and is used to iterate over the raw logs and unpacked data for SetBatchId events raised by the Game contract.
type GameSetBatchIdIterator struct {
	Event *GameSetBatchId // Event containing the contract specifics and raw log

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
func (it *GameSetBatchIdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameSetBatchId)
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
		it.Event = new(GameSetBatchId)
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
func (it *GameSetBatchIdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameSetBatchIdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameSetBatchId represents a SetBatchId event raised by the Game contract.
type GameSetBatchId struct {
	BatchId  *big.Int
	TokenIds []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetBatchId is a free log retrieval operation binding the contract event 0xb1e5db2cea945a3cba3caf23da4140976f1f3448b59fa0c009c4792d6c176524.
//
// Solidity: event SetBatchId(uint256 indexed _batchId, uint256[] tokenIds)
func (_Game *GameFilterer) FilterSetBatchId(opts *bind.FilterOpts, _batchId []*big.Int) (*GameSetBatchIdIterator, error) {

	var _batchIdRule []interface{}
	for _, _batchIdItem := range _batchId {
		_batchIdRule = append(_batchIdRule, _batchIdItem)
	}

	logs, sub, err := _Game.contract.FilterLogs(opts, "SetBatchId", _batchIdRule)
	if err != nil {
		return nil, err
	}
	return &GameSetBatchIdIterator{contract: _Game.contract, event: "SetBatchId", logs: logs, sub: sub}, nil
}

// WatchSetBatchId is a free log subscription operation binding the contract event 0xb1e5db2cea945a3cba3caf23da4140976f1f3448b59fa0c009c4792d6c176524.
//
// Solidity: event SetBatchId(uint256 indexed _batchId, uint256[] tokenIds)
func (_Game *GameFilterer) WatchSetBatchId(opts *bind.WatchOpts, sink chan<- *GameSetBatchId, _batchId []*big.Int) (event.Subscription, error) {

	var _batchIdRule []interface{}
	for _, _batchIdItem := range _batchId {
		_batchIdRule = append(_batchIdRule, _batchIdItem)
	}

	logs, sub, err := _Game.contract.WatchLogs(opts, "SetBatchId", _batchIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameSetBatchId)
				if err := _Game.contract.UnpackLog(event, "SetBatchId", log); err != nil {
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

// ParseSetBatchId is a log parse operation binding the contract event 0xb1e5db2cea945a3cba3caf23da4140976f1f3448b59fa0c009c4792d6c176524.
//
// Solidity: event SetBatchId(uint256 indexed _batchId, uint256[] tokenIds)
func (_Game *GameFilterer) ParseSetBatchId(log types.Log) (*GameSetBatchId, error) {
	event := new(GameSetBatchId)
	if err := _Game.contract.UnpackLog(event, "SetBatchId", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GameSpendSkillpointsIterator is returned from FilterSpendSkillpoints and is used to iterate over the raw logs and unpacked data for SpendSkillpoints events raised by the Game contract.
type GameSpendSkillpointsIterator struct {
	Event *GameSpendSkillpoints // Event containing the contract specifics and raw log

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
func (it *GameSpendSkillpointsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameSpendSkillpoints)
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
		it.Event = new(GameSpendSkillpoints)
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
func (it *GameSpendSkillpointsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameSpendSkillpointsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameSpendSkillpoints represents a SpendSkillpoints event raised by the Game contract.
type GameSpendSkillpoints struct {
	TokenId *big.Int
	Values  [4]int16
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSpendSkillpoints is a free log retrieval operation binding the contract event 0x4b92ed0db0758e7ed73674d6915bd532a55322aa1b13abeaad14acf3954b1a8c.
//
// Solidity: event SpendSkillpoints(uint256 indexed _tokenId, int16[4] _values)
func (_Game *GameFilterer) FilterSpendSkillpoints(opts *bind.FilterOpts, _tokenId []*big.Int) (*GameSpendSkillpointsIterator, error) {

	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _Game.contract.FilterLogs(opts, "SpendSkillpoints", _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &GameSpendSkillpointsIterator{contract: _Game.contract, event: "SpendSkillpoints", logs: logs, sub: sub}, nil
}

// WatchSpendSkillpoints is a free log subscription operation binding the contract event 0x4b92ed0db0758e7ed73674d6915bd532a55322aa1b13abeaad14acf3954b1a8c.
//
// Solidity: event SpendSkillpoints(uint256 indexed _tokenId, int16[4] _values)
func (_Game *GameFilterer) WatchSpendSkillpoints(opts *bind.WatchOpts, sink chan<- *GameSpendSkillpoints, _tokenId []*big.Int) (event.Subscription, error) {

	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _Game.contract.WatchLogs(opts, "SpendSkillpoints", _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameSpendSkillpoints)
				if err := _Game.contract.UnpackLog(event, "SpendSkillpoints", log); err != nil {
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

// ParseSpendSkillpoints is a log parse operation binding the contract event 0x4b92ed0db0758e7ed73674d6915bd532a55322aa1b13abeaad14acf3954b1a8c.
//
// Solidity: event SpendSkillpoints(uint256 indexed _tokenId, int16[4] _values)
func (_Game *GameFilterer) ParseSpendSkillpoints(log types.Log) (*GameSpendSkillpoints, error) {
	event := new(GameSpendSkillpoints)
	if err := _Game.contract.UnpackLog(event, "SpendSkillpoints", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GameUnLockAavegotchiIterator is returned from FilterUnLockAavegotchi and is used to iterate over the raw logs and unpacked data for UnLockAavegotchi events raised by the Game contract.
type GameUnLockAavegotchiIterator struct {
	Event *GameUnLockAavegotchi // Event containing the contract specifics and raw log

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
func (it *GameUnLockAavegotchiIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameUnLockAavegotchi)
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
		it.Event = new(GameUnLockAavegotchi)
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
func (it *GameUnLockAavegotchiIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameUnLockAavegotchiIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameUnLockAavegotchi represents a UnLockAavegotchi event raised by the Game contract.
type GameUnLockAavegotchi struct {
	TokenId *big.Int
	Time    *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnLockAavegotchi is a free log retrieval operation binding the contract event 0x5e71cac5a6405e3008ac4349af37d908bfd84e12e76f71e8ea270210fe2635a6.
//
// Solidity: event UnLockAavegotchi(uint256 indexed _tokenId, uint256 _time)
func (_Game *GameFilterer) FilterUnLockAavegotchi(opts *bind.FilterOpts, _tokenId []*big.Int) (*GameUnLockAavegotchiIterator, error) {

	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _Game.contract.FilterLogs(opts, "UnLockAavegotchi", _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &GameUnLockAavegotchiIterator{contract: _Game.contract, event: "UnLockAavegotchi", logs: logs, sub: sub}, nil
}

// WatchUnLockAavegotchi is a free log subscription operation binding the contract event 0x5e71cac5a6405e3008ac4349af37d908bfd84e12e76f71e8ea270210fe2635a6.
//
// Solidity: event UnLockAavegotchi(uint256 indexed _tokenId, uint256 _time)
func (_Game *GameFilterer) WatchUnLockAavegotchi(opts *bind.WatchOpts, sink chan<- *GameUnLockAavegotchi, _tokenId []*big.Int) (event.Subscription, error) {

	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _Game.contract.WatchLogs(opts, "UnLockAavegotchi", _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameUnLockAavegotchi)
				if err := _Game.contract.UnpackLog(event, "UnLockAavegotchi", log); err != nil {
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

// ParseUnLockAavegotchi is a log parse operation binding the contract event 0x5e71cac5a6405e3008ac4349af37d908bfd84e12e76f71e8ea270210fe2635a6.
//
// Solidity: event UnLockAavegotchi(uint256 indexed _tokenId, uint256 _time)
func (_Game *GameFilterer) ParseUnLockAavegotchi(log types.Log) (*GameUnLockAavegotchi, error) {
	event := new(GameUnLockAavegotchi)
	if err := _Game.contract.UnpackLog(event, "UnLockAavegotchi", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
