// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package balancer

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

// IVaultBatchSwapStep is an auto generated low-level Go binding around an user-defined struct.
type IVaultBatchSwapStep struct {
	PoolId        [32]byte
	AssetInIndex  *big.Int
	AssetOutIndex *big.Int
	Amount        *big.Int
	UserData      []byte
}

// IVaultExitPoolRequest is an auto generated low-level Go binding around an user-defined struct.
type IVaultExitPoolRequest struct {
	Assets            []common.Address
	MinAmountsOut     []*big.Int
	UserData          []byte
	ToInternalBalance bool
}

// IVaultFundManagement is an auto generated low-level Go binding around an user-defined struct.
type IVaultFundManagement struct {
	Sender              common.Address
	FromInternalBalance bool
	Recipient           common.Address
	ToInternalBalance   bool
}

// IVaultJoinPoolRequest is an auto generated low-level Go binding around an user-defined struct.
type IVaultJoinPoolRequest struct {
	Assets              []common.Address
	MaxAmountsIn        []*big.Int
	UserData            []byte
	FromInternalBalance bool
}

// IVaultPoolBalanceOp is an auto generated low-level Go binding around an user-defined struct.
type IVaultPoolBalanceOp struct {
	Kind   uint8
	PoolId [32]byte
	Token  common.Address
	Amount *big.Int
}

// IVaultSingleSwap is an auto generated low-level Go binding around an user-defined struct.
type IVaultSingleSwap struct {
	PoolId   [32]byte
	Kind     uint8
	AssetIn  common.Address
	AssetOut common.Address
	Amount   *big.Int
	UserData []byte
}

// IVaultUserBalanceOp is an auto generated low-level Go binding around an user-defined struct.
type IVaultUserBalanceOp struct {
	Kind      uint8
	Asset     common.Address
	Amount    *big.Int
	Sender    common.Address
	Recipient common.Address
}

// VaultMetaData contains all meta data concerning the Vault contract.
var VaultMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIAuthorizer\",\"name\":\"authorizer\",\"type\":\"address\"},{\"internalType\":\"contractIWETH\",\"name\":\"weth\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pauseWindowDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bufferPeriodDuration\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIAuthorizer\",\"name\":\"newAuthorizer\",\"type\":\"address\"}],\"name\":\"AuthorizerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ExternalBalanceTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIFlashLoanRecipient\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"FlashLoan\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"delta\",\"type\":\"int256\"}],\"name\":\"InternalBalanceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"}],\"name\":\"PausedStateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"liquidityProvider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"int256[]\",\"name\":\"deltas\",\"type\":\"int256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"protocolFeeAmounts\",\"type\":\"uint256[]\"}],\"name\":\"PoolBalanceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"assetManager\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"cashDelta\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"managedDelta\",\"type\":\"int256\"}],\"name\":\"PoolBalanceManaged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"poolAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumIVault.PoolSpecialization\",\"name\":\"specialization\",\"type\":\"uint8\"}],\"name\":\"PoolRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"RelayerApprovalChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"TokensDeregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"assetManagers\",\"type\":\"address[]\"}],\"name\":\"TokensRegistered\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"contractIWETH\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumIVault.SwapKind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"assetInIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetOutIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"internalType\":\"structIVault.BatchSwapStep[]\",\"name\":\"swaps\",\"type\":\"tuple[]\"},{\"internalType\":\"contractIAsset[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"fromInternalBalance\",\"type\":\"bool\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"toInternalBalance\",\"type\":\"bool\"}],\"internalType\":\"structIVault.FundManagement\",\"name\":\"funds\",\"type\":\"tuple\"},{\"internalType\":\"int256[]\",\"name\":\"limits\",\"type\":\"int256[]\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"batchSwap\",\"outputs\":[{\"internalType\":\"int256[]\",\"name\":\"assetDeltas\",\"type\":\"int256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"deregisterTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"contractIAsset[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"minAmountsOut\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"toInternalBalance\",\"type\":\"bool\"}],\"internalType\":\"structIVault.ExitPoolRequest\",\"name\":\"request\",\"type\":\"tuple\"}],\"name\":\"exitPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIFlashLoanRecipient\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"name\":\"flashLoan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"getActionId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAuthorizer\",\"outputs\":[{\"internalType\":\"contractIAuthorizer\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDomainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getInternalBalance\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getNextNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPausedState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"pauseWindowEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bufferPeriodEndTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"}],\"name\":\"getPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"enumIVault.PoolSpecialization\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getPoolTokenInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"cash\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"managed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastChangeBlock\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"assetManager\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"}],\"name\":\"getPoolTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"lastChangeBlock\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProtocolFeesCollector\",\"outputs\":[{\"internalType\":\"contractProtocolFeesCollector\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"}],\"name\":\"hasApprovedRelayer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"contractIAsset[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"maxAmountsIn\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"fromInternalBalance\",\"type\":\"bool\"}],\"internalType\":\"structIVault.JoinPoolRequest\",\"name\":\"request\",\"type\":\"tuple\"}],\"name\":\"joinPool\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumIVault.PoolBalanceOpKind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIVault.PoolBalanceOp[]\",\"name\":\"ops\",\"type\":\"tuple[]\"}],\"name\":\"managePoolBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumIVault.UserBalanceOpKind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"contractIAsset\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structIVault.UserBalanceOp[]\",\"name\":\"ops\",\"type\":\"tuple[]\"}],\"name\":\"manageUserBalance\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumIVault.SwapKind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"assetInIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetOutIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"internalType\":\"structIVault.BatchSwapStep[]\",\"name\":\"swaps\",\"type\":\"tuple[]\"},{\"internalType\":\"contractIAsset[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"fromInternalBalance\",\"type\":\"bool\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"toInternalBalance\",\"type\":\"bool\"}],\"internalType\":\"structIVault.FundManagement\",\"name\":\"funds\",\"type\":\"tuple\"}],\"name\":\"queryBatchSwap\",\"outputs\":[{\"internalType\":\"int256[]\",\"name\":\"\",\"type\":\"int256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumIVault.PoolSpecialization\",\"name\":\"specialization\",\"type\":\"uint8\"}],\"name\":\"registerPool\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"assetManagers\",\"type\":\"address[]\"}],\"name\":\"registerTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIAuthorizer\",\"name\":\"newAuthorizer\",\"type\":\"address\"}],\"name\":\"setAuthorizer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"}],\"name\":\"setPaused\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setRelayerApproval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"enumIVault.SwapKind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"contractIAsset\",\"name\":\"assetIn\",\"type\":\"address\"},{\"internalType\":\"contractIAsset\",\"name\":\"assetOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"internalType\":\"structIVault.SingleSwap\",\"name\":\"singleSwap\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"fromInternalBalance\",\"type\":\"bool\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"toInternalBalance\",\"type\":\"bool\"}],\"internalType\":\"structIVault.FundManagement\",\"name\":\"funds\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountCalculated\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// VaultABI is the input ABI used to generate the binding from.
// Deprecated: Use VaultMetaData.ABI instead.
var VaultABI = VaultMetaData.ABI

// Vault is an auto generated Go binding around an Ethereum contract.
type Vault struct {
	VaultCaller     // Read-only binding to the contract
	VaultTransactor // Write-only binding to the contract
	VaultFilterer   // Log filterer for contract events
}

// VaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type VaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VaultSession struct {
	Contract     *Vault            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VaultCallerSession struct {
	Contract *VaultCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VaultTransactorSession struct {
	Contract     *VaultTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type VaultRaw struct {
	Contract *Vault // Generic contract binding to access the raw methods on
}

// VaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VaultCallerRaw struct {
	Contract *VaultCaller // Generic read-only contract binding to access the raw methods on
}

// VaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VaultTransactorRaw struct {
	Contract *VaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVault creates a new instance of Vault, bound to a specific deployed contract.
func NewVault(address common.Address, backend bind.ContractBackend) (*Vault, error) {
	contract, err := bindVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vault{VaultCaller: VaultCaller{contract: contract}, VaultTransactor: VaultTransactor{contract: contract}, VaultFilterer: VaultFilterer{contract: contract}}, nil
}

// NewVaultCaller creates a new read-only instance of Vault, bound to a specific deployed contract.
func NewVaultCaller(address common.Address, caller bind.ContractCaller) (*VaultCaller, error) {
	contract, err := bindVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VaultCaller{contract: contract}, nil
}

// NewVaultTransactor creates a new write-only instance of Vault, bound to a specific deployed contract.
func NewVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*VaultTransactor, error) {
	contract, err := bindVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VaultTransactor{contract: contract}, nil
}

// NewVaultFilterer creates a new log filterer instance of Vault, bound to a specific deployed contract.
func NewVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*VaultFilterer, error) {
	contract, err := bindVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VaultFilterer{contract: contract}, nil
}

// bindVault binds a generic wrapper to an already deployed contract.
func bindVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VaultABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vault *VaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vault.Contract.VaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vault *VaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vault.Contract.VaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vault *VaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vault.Contract.VaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vault *VaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vault *VaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vault *VaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vault.Contract.contract.Transact(opts, method, params...)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Vault *VaultCaller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Vault *VaultSession) WETH() (common.Address, error) {
	return _Vault.Contract.WETH(&_Vault.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Vault *VaultCallerSession) WETH() (common.Address, error) {
	return _Vault.Contract.WETH(&_Vault.CallOpts)
}

// GetActionId is a free data retrieval call binding the contract method 0x851c1bb3.
//
// Solidity: function getActionId(bytes4 selector) view returns(bytes32)
func (_Vault *VaultCaller) GetActionId(opts *bind.CallOpts, selector [4]byte) ([32]byte, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "getActionId", selector)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetActionId is a free data retrieval call binding the contract method 0x851c1bb3.
//
// Solidity: function getActionId(bytes4 selector) view returns(bytes32)
func (_Vault *VaultSession) GetActionId(selector [4]byte) ([32]byte, error) {
	return _Vault.Contract.GetActionId(&_Vault.CallOpts, selector)
}

// GetActionId is a free data retrieval call binding the contract method 0x851c1bb3.
//
// Solidity: function getActionId(bytes4 selector) view returns(bytes32)
func (_Vault *VaultCallerSession) GetActionId(selector [4]byte) ([32]byte, error) {
	return _Vault.Contract.GetActionId(&_Vault.CallOpts, selector)
}

// GetAuthorizer is a free data retrieval call binding the contract method 0xaaabadc5.
//
// Solidity: function getAuthorizer() view returns(address)
func (_Vault *VaultCaller) GetAuthorizer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "getAuthorizer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAuthorizer is a free data retrieval call binding the contract method 0xaaabadc5.
//
// Solidity: function getAuthorizer() view returns(address)
func (_Vault *VaultSession) GetAuthorizer() (common.Address, error) {
	return _Vault.Contract.GetAuthorizer(&_Vault.CallOpts)
}

// GetAuthorizer is a free data retrieval call binding the contract method 0xaaabadc5.
//
// Solidity: function getAuthorizer() view returns(address)
func (_Vault *VaultCallerSession) GetAuthorizer() (common.Address, error) {
	return _Vault.Contract.GetAuthorizer(&_Vault.CallOpts)
}

// GetDomainSeparator is a free data retrieval call binding the contract method 0xed24911d.
//
// Solidity: function getDomainSeparator() view returns(bytes32)
func (_Vault *VaultCaller) GetDomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "getDomainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetDomainSeparator is a free data retrieval call binding the contract method 0xed24911d.
//
// Solidity: function getDomainSeparator() view returns(bytes32)
func (_Vault *VaultSession) GetDomainSeparator() ([32]byte, error) {
	return _Vault.Contract.GetDomainSeparator(&_Vault.CallOpts)
}

// GetDomainSeparator is a free data retrieval call binding the contract method 0xed24911d.
//
// Solidity: function getDomainSeparator() view returns(bytes32)
func (_Vault *VaultCallerSession) GetDomainSeparator() ([32]byte, error) {
	return _Vault.Contract.GetDomainSeparator(&_Vault.CallOpts)
}

// GetInternalBalance is a free data retrieval call binding the contract method 0x0f5a6efa.
//
// Solidity: function getInternalBalance(address user, address[] tokens) view returns(uint256[] balances)
func (_Vault *VaultCaller) GetInternalBalance(opts *bind.CallOpts, user common.Address, tokens []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "getInternalBalance", user, tokens)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetInternalBalance is a free data retrieval call binding the contract method 0x0f5a6efa.
//
// Solidity: function getInternalBalance(address user, address[] tokens) view returns(uint256[] balances)
func (_Vault *VaultSession) GetInternalBalance(user common.Address, tokens []common.Address) ([]*big.Int, error) {
	return _Vault.Contract.GetInternalBalance(&_Vault.CallOpts, user, tokens)
}

// GetInternalBalance is a free data retrieval call binding the contract method 0x0f5a6efa.
//
// Solidity: function getInternalBalance(address user, address[] tokens) view returns(uint256[] balances)
func (_Vault *VaultCallerSession) GetInternalBalance(user common.Address, tokens []common.Address) ([]*big.Int, error) {
	return _Vault.Contract.GetInternalBalance(&_Vault.CallOpts, user, tokens)
}

// GetNextNonce is a free data retrieval call binding the contract method 0x90193b7c.
//
// Solidity: function getNextNonce(address user) view returns(uint256)
func (_Vault *VaultCaller) GetNextNonce(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "getNextNonce", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextNonce is a free data retrieval call binding the contract method 0x90193b7c.
//
// Solidity: function getNextNonce(address user) view returns(uint256)
func (_Vault *VaultSession) GetNextNonce(user common.Address) (*big.Int, error) {
	return _Vault.Contract.GetNextNonce(&_Vault.CallOpts, user)
}

// GetNextNonce is a free data retrieval call binding the contract method 0x90193b7c.
//
// Solidity: function getNextNonce(address user) view returns(uint256)
func (_Vault *VaultCallerSession) GetNextNonce(user common.Address) (*big.Int, error) {
	return _Vault.Contract.GetNextNonce(&_Vault.CallOpts, user)
}

// GetPausedState is a free data retrieval call binding the contract method 0x1c0de051.
//
// Solidity: function getPausedState() view returns(bool paused, uint256 pauseWindowEndTime, uint256 bufferPeriodEndTime)
func (_Vault *VaultCaller) GetPausedState(opts *bind.CallOpts) (struct {
	Paused              bool
	PauseWindowEndTime  *big.Int
	BufferPeriodEndTime *big.Int
}, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "getPausedState")

	outstruct := new(struct {
		Paused              bool
		PauseWindowEndTime  *big.Int
		BufferPeriodEndTime *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Paused = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.PauseWindowEndTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BufferPeriodEndTime = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetPausedState is a free data retrieval call binding the contract method 0x1c0de051.
//
// Solidity: function getPausedState() view returns(bool paused, uint256 pauseWindowEndTime, uint256 bufferPeriodEndTime)
func (_Vault *VaultSession) GetPausedState() (struct {
	Paused              bool
	PauseWindowEndTime  *big.Int
	BufferPeriodEndTime *big.Int
}, error) {
	return _Vault.Contract.GetPausedState(&_Vault.CallOpts)
}

// GetPausedState is a free data retrieval call binding the contract method 0x1c0de051.
//
// Solidity: function getPausedState() view returns(bool paused, uint256 pauseWindowEndTime, uint256 bufferPeriodEndTime)
func (_Vault *VaultCallerSession) GetPausedState() (struct {
	Paused              bool
	PauseWindowEndTime  *big.Int
	BufferPeriodEndTime *big.Int
}, error) {
	return _Vault.Contract.GetPausedState(&_Vault.CallOpts)
}

// GetPool is a free data retrieval call binding the contract method 0xf6c00927.
//
// Solidity: function getPool(bytes32 poolId) view returns(address, uint8)
func (_Vault *VaultCaller) GetPool(opts *bind.CallOpts, poolId [32]byte) (common.Address, uint8, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "getPool", poolId)

	if err != nil {
		return *new(common.Address), *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(uint8)).(*uint8)

	return out0, out1, err

}

// GetPool is a free data retrieval call binding the contract method 0xf6c00927.
//
// Solidity: function getPool(bytes32 poolId) view returns(address, uint8)
func (_Vault *VaultSession) GetPool(poolId [32]byte) (common.Address, uint8, error) {
	return _Vault.Contract.GetPool(&_Vault.CallOpts, poolId)
}

// GetPool is a free data retrieval call binding the contract method 0xf6c00927.
//
// Solidity: function getPool(bytes32 poolId) view returns(address, uint8)
func (_Vault *VaultCallerSession) GetPool(poolId [32]byte) (common.Address, uint8, error) {
	return _Vault.Contract.GetPool(&_Vault.CallOpts, poolId)
}

// GetPoolTokenInfo is a free data retrieval call binding the contract method 0xb05f8e48.
//
// Solidity: function getPoolTokenInfo(bytes32 poolId, address token) view returns(uint256 cash, uint256 managed, uint256 lastChangeBlock, address assetManager)
func (_Vault *VaultCaller) GetPoolTokenInfo(opts *bind.CallOpts, poolId [32]byte, token common.Address) (struct {
	Cash            *big.Int
	Managed         *big.Int
	LastChangeBlock *big.Int
	AssetManager    common.Address
}, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "getPoolTokenInfo", poolId, token)

	outstruct := new(struct {
		Cash            *big.Int
		Managed         *big.Int
		LastChangeBlock *big.Int
		AssetManager    common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Cash = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Managed = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LastChangeBlock = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.AssetManager = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// GetPoolTokenInfo is a free data retrieval call binding the contract method 0xb05f8e48.
//
// Solidity: function getPoolTokenInfo(bytes32 poolId, address token) view returns(uint256 cash, uint256 managed, uint256 lastChangeBlock, address assetManager)
func (_Vault *VaultSession) GetPoolTokenInfo(poolId [32]byte, token common.Address) (struct {
	Cash            *big.Int
	Managed         *big.Int
	LastChangeBlock *big.Int
	AssetManager    common.Address
}, error) {
	return _Vault.Contract.GetPoolTokenInfo(&_Vault.CallOpts, poolId, token)
}

// GetPoolTokenInfo is a free data retrieval call binding the contract method 0xb05f8e48.
//
// Solidity: function getPoolTokenInfo(bytes32 poolId, address token) view returns(uint256 cash, uint256 managed, uint256 lastChangeBlock, address assetManager)
func (_Vault *VaultCallerSession) GetPoolTokenInfo(poolId [32]byte, token common.Address) (struct {
	Cash            *big.Int
	Managed         *big.Int
	LastChangeBlock *big.Int
	AssetManager    common.Address
}, error) {
	return _Vault.Contract.GetPoolTokenInfo(&_Vault.CallOpts, poolId, token)
}

// GetPoolTokens is a free data retrieval call binding the contract method 0xf94d4668.
//
// Solidity: function getPoolTokens(bytes32 poolId) view returns(address[] tokens, uint256[] balances, uint256 lastChangeBlock)
func (_Vault *VaultCaller) GetPoolTokens(opts *bind.CallOpts, poolId [32]byte) (struct {
	Tokens          []common.Address
	Balances        []*big.Int
	LastChangeBlock *big.Int
}, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "getPoolTokens", poolId)

	outstruct := new(struct {
		Tokens          []common.Address
		Balances        []*big.Int
		LastChangeBlock *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Tokens = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Balances = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	outstruct.LastChangeBlock = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetPoolTokens is a free data retrieval call binding the contract method 0xf94d4668.
//
// Solidity: function getPoolTokens(bytes32 poolId) view returns(address[] tokens, uint256[] balances, uint256 lastChangeBlock)
func (_Vault *VaultSession) GetPoolTokens(poolId [32]byte) (struct {
	Tokens          []common.Address
	Balances        []*big.Int
	LastChangeBlock *big.Int
}, error) {
	return _Vault.Contract.GetPoolTokens(&_Vault.CallOpts, poolId)
}

// GetPoolTokens is a free data retrieval call binding the contract method 0xf94d4668.
//
// Solidity: function getPoolTokens(bytes32 poolId) view returns(address[] tokens, uint256[] balances, uint256 lastChangeBlock)
func (_Vault *VaultCallerSession) GetPoolTokens(poolId [32]byte) (struct {
	Tokens          []common.Address
	Balances        []*big.Int
	LastChangeBlock *big.Int
}, error) {
	return _Vault.Contract.GetPoolTokens(&_Vault.CallOpts, poolId)
}

// GetProtocolFeesCollector is a free data retrieval call binding the contract method 0xd2946c2b.
//
// Solidity: function getProtocolFeesCollector() view returns(address)
func (_Vault *VaultCaller) GetProtocolFeesCollector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "getProtocolFeesCollector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetProtocolFeesCollector is a free data retrieval call binding the contract method 0xd2946c2b.
//
// Solidity: function getProtocolFeesCollector() view returns(address)
func (_Vault *VaultSession) GetProtocolFeesCollector() (common.Address, error) {
	return _Vault.Contract.GetProtocolFeesCollector(&_Vault.CallOpts)
}

// GetProtocolFeesCollector is a free data retrieval call binding the contract method 0xd2946c2b.
//
// Solidity: function getProtocolFeesCollector() view returns(address)
func (_Vault *VaultCallerSession) GetProtocolFeesCollector() (common.Address, error) {
	return _Vault.Contract.GetProtocolFeesCollector(&_Vault.CallOpts)
}

// HasApprovedRelayer is a free data retrieval call binding the contract method 0xfec90d72.
//
// Solidity: function hasApprovedRelayer(address user, address relayer) view returns(bool)
func (_Vault *VaultCaller) HasApprovedRelayer(opts *bind.CallOpts, user common.Address, relayer common.Address) (bool, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "hasApprovedRelayer", user, relayer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasApprovedRelayer is a free data retrieval call binding the contract method 0xfec90d72.
//
// Solidity: function hasApprovedRelayer(address user, address relayer) view returns(bool)
func (_Vault *VaultSession) HasApprovedRelayer(user common.Address, relayer common.Address) (bool, error) {
	return _Vault.Contract.HasApprovedRelayer(&_Vault.CallOpts, user, relayer)
}

// HasApprovedRelayer is a free data retrieval call binding the contract method 0xfec90d72.
//
// Solidity: function hasApprovedRelayer(address user, address relayer) view returns(bool)
func (_Vault *VaultCallerSession) HasApprovedRelayer(user common.Address, relayer common.Address) (bool, error) {
	return _Vault.Contract.HasApprovedRelayer(&_Vault.CallOpts, user, relayer)
}

// BatchSwap is a paid mutator transaction binding the contract method 0x945bcec9.
//
// Solidity: function batchSwap(uint8 kind, (bytes32,uint256,uint256,uint256,bytes)[] swaps, address[] assets, (address,bool,address,bool) funds, int256[] limits, uint256 deadline) payable returns(int256[] assetDeltas)
func (_Vault *VaultTransactor) BatchSwap(opts *bind.TransactOpts, kind uint8, swaps []IVaultBatchSwapStep, assets []common.Address, funds IVaultFundManagement, limits []*big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "batchSwap", kind, swaps, assets, funds, limits, deadline)
}

// BatchSwap is a paid mutator transaction binding the contract method 0x945bcec9.
//
// Solidity: function batchSwap(uint8 kind, (bytes32,uint256,uint256,uint256,bytes)[] swaps, address[] assets, (address,bool,address,bool) funds, int256[] limits, uint256 deadline) payable returns(int256[] assetDeltas)
func (_Vault *VaultSession) BatchSwap(kind uint8, swaps []IVaultBatchSwapStep, assets []common.Address, funds IVaultFundManagement, limits []*big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.BatchSwap(&_Vault.TransactOpts, kind, swaps, assets, funds, limits, deadline)
}

// BatchSwap is a paid mutator transaction binding the contract method 0x945bcec9.
//
// Solidity: function batchSwap(uint8 kind, (bytes32,uint256,uint256,uint256,bytes)[] swaps, address[] assets, (address,bool,address,bool) funds, int256[] limits, uint256 deadline) payable returns(int256[] assetDeltas)
func (_Vault *VaultTransactorSession) BatchSwap(kind uint8, swaps []IVaultBatchSwapStep, assets []common.Address, funds IVaultFundManagement, limits []*big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.BatchSwap(&_Vault.TransactOpts, kind, swaps, assets, funds, limits, deadline)
}

// DeregisterTokens is a paid mutator transaction binding the contract method 0x7d3aeb96.
//
// Solidity: function deregisterTokens(bytes32 poolId, address[] tokens) returns()
func (_Vault *VaultTransactor) DeregisterTokens(opts *bind.TransactOpts, poolId [32]byte, tokens []common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "deregisterTokens", poolId, tokens)
}

// DeregisterTokens is a paid mutator transaction binding the contract method 0x7d3aeb96.
//
// Solidity: function deregisterTokens(bytes32 poolId, address[] tokens) returns()
func (_Vault *VaultSession) DeregisterTokens(poolId [32]byte, tokens []common.Address) (*types.Transaction, error) {
	return _Vault.Contract.DeregisterTokens(&_Vault.TransactOpts, poolId, tokens)
}

// DeregisterTokens is a paid mutator transaction binding the contract method 0x7d3aeb96.
//
// Solidity: function deregisterTokens(bytes32 poolId, address[] tokens) returns()
func (_Vault *VaultTransactorSession) DeregisterTokens(poolId [32]byte, tokens []common.Address) (*types.Transaction, error) {
	return _Vault.Contract.DeregisterTokens(&_Vault.TransactOpts, poolId, tokens)
}

// ExitPool is a paid mutator transaction binding the contract method 0x8bdb3913.
//
// Solidity: function exitPool(bytes32 poolId, address sender, address recipient, (address[],uint256[],bytes,bool) request) returns()
func (_Vault *VaultTransactor) ExitPool(opts *bind.TransactOpts, poolId [32]byte, sender common.Address, recipient common.Address, request IVaultExitPoolRequest) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "exitPool", poolId, sender, recipient, request)
}

// ExitPool is a paid mutator transaction binding the contract method 0x8bdb3913.
//
// Solidity: function exitPool(bytes32 poolId, address sender, address recipient, (address[],uint256[],bytes,bool) request) returns()
func (_Vault *VaultSession) ExitPool(poolId [32]byte, sender common.Address, recipient common.Address, request IVaultExitPoolRequest) (*types.Transaction, error) {
	return _Vault.Contract.ExitPool(&_Vault.TransactOpts, poolId, sender, recipient, request)
}

// ExitPool is a paid mutator transaction binding the contract method 0x8bdb3913.
//
// Solidity: function exitPool(bytes32 poolId, address sender, address recipient, (address[],uint256[],bytes,bool) request) returns()
func (_Vault *VaultTransactorSession) ExitPool(poolId [32]byte, sender common.Address, recipient common.Address, request IVaultExitPoolRequest) (*types.Transaction, error) {
	return _Vault.Contract.ExitPool(&_Vault.TransactOpts, poolId, sender, recipient, request)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5c38449e.
//
// Solidity: function flashLoan(address recipient, address[] tokens, uint256[] amounts, bytes userData) returns()
func (_Vault *VaultTransactor) FlashLoan(opts *bind.TransactOpts, recipient common.Address, tokens []common.Address, amounts []*big.Int, userData []byte) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "flashLoan", recipient, tokens, amounts, userData)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5c38449e.
//
// Solidity: function flashLoan(address recipient, address[] tokens, uint256[] amounts, bytes userData) returns()
func (_Vault *VaultSession) FlashLoan(recipient common.Address, tokens []common.Address, amounts []*big.Int, userData []byte) (*types.Transaction, error) {
	return _Vault.Contract.FlashLoan(&_Vault.TransactOpts, recipient, tokens, amounts, userData)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5c38449e.
//
// Solidity: function flashLoan(address recipient, address[] tokens, uint256[] amounts, bytes userData) returns()
func (_Vault *VaultTransactorSession) FlashLoan(recipient common.Address, tokens []common.Address, amounts []*big.Int, userData []byte) (*types.Transaction, error) {
	return _Vault.Contract.FlashLoan(&_Vault.TransactOpts, recipient, tokens, amounts, userData)
}

// JoinPool is a paid mutator transaction binding the contract method 0xb95cac28.
//
// Solidity: function joinPool(bytes32 poolId, address sender, address recipient, (address[],uint256[],bytes,bool) request) payable returns()
func (_Vault *VaultTransactor) JoinPool(opts *bind.TransactOpts, poolId [32]byte, sender common.Address, recipient common.Address, request IVaultJoinPoolRequest) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "joinPool", poolId, sender, recipient, request)
}

// JoinPool is a paid mutator transaction binding the contract method 0xb95cac28.
//
// Solidity: function joinPool(bytes32 poolId, address sender, address recipient, (address[],uint256[],bytes,bool) request) payable returns()
func (_Vault *VaultSession) JoinPool(poolId [32]byte, sender common.Address, recipient common.Address, request IVaultJoinPoolRequest) (*types.Transaction, error) {
	return _Vault.Contract.JoinPool(&_Vault.TransactOpts, poolId, sender, recipient, request)
}

// JoinPool is a paid mutator transaction binding the contract method 0xb95cac28.
//
// Solidity: function joinPool(bytes32 poolId, address sender, address recipient, (address[],uint256[],bytes,bool) request) payable returns()
func (_Vault *VaultTransactorSession) JoinPool(poolId [32]byte, sender common.Address, recipient common.Address, request IVaultJoinPoolRequest) (*types.Transaction, error) {
	return _Vault.Contract.JoinPool(&_Vault.TransactOpts, poolId, sender, recipient, request)
}

// ManagePoolBalance is a paid mutator transaction binding the contract method 0xe6c46092.
//
// Solidity: function managePoolBalance((uint8,bytes32,address,uint256)[] ops) returns()
func (_Vault *VaultTransactor) ManagePoolBalance(opts *bind.TransactOpts, ops []IVaultPoolBalanceOp) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "managePoolBalance", ops)
}

// ManagePoolBalance is a paid mutator transaction binding the contract method 0xe6c46092.
//
// Solidity: function managePoolBalance((uint8,bytes32,address,uint256)[] ops) returns()
func (_Vault *VaultSession) ManagePoolBalance(ops []IVaultPoolBalanceOp) (*types.Transaction, error) {
	return _Vault.Contract.ManagePoolBalance(&_Vault.TransactOpts, ops)
}

// ManagePoolBalance is a paid mutator transaction binding the contract method 0xe6c46092.
//
// Solidity: function managePoolBalance((uint8,bytes32,address,uint256)[] ops) returns()
func (_Vault *VaultTransactorSession) ManagePoolBalance(ops []IVaultPoolBalanceOp) (*types.Transaction, error) {
	return _Vault.Contract.ManagePoolBalance(&_Vault.TransactOpts, ops)
}

// ManageUserBalance is a paid mutator transaction binding the contract method 0x0e8e3e84.
//
// Solidity: function manageUserBalance((uint8,address,uint256,address,address)[] ops) payable returns()
func (_Vault *VaultTransactor) ManageUserBalance(opts *bind.TransactOpts, ops []IVaultUserBalanceOp) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "manageUserBalance", ops)
}

// ManageUserBalance is a paid mutator transaction binding the contract method 0x0e8e3e84.
//
// Solidity: function manageUserBalance((uint8,address,uint256,address,address)[] ops) payable returns()
func (_Vault *VaultSession) ManageUserBalance(ops []IVaultUserBalanceOp) (*types.Transaction, error) {
	return _Vault.Contract.ManageUserBalance(&_Vault.TransactOpts, ops)
}

// ManageUserBalance is a paid mutator transaction binding the contract method 0x0e8e3e84.
//
// Solidity: function manageUserBalance((uint8,address,uint256,address,address)[] ops) payable returns()
func (_Vault *VaultTransactorSession) ManageUserBalance(ops []IVaultUserBalanceOp) (*types.Transaction, error) {
	return _Vault.Contract.ManageUserBalance(&_Vault.TransactOpts, ops)
}

// QueryBatchSwap is a paid mutator transaction binding the contract method 0xf84d066e.
//
// Solidity: function queryBatchSwap(uint8 kind, (bytes32,uint256,uint256,uint256,bytes)[] swaps, address[] assets, (address,bool,address,bool) funds) returns(int256[])
func (_Vault *VaultTransactor) QueryBatchSwap(opts *bind.TransactOpts, kind uint8, swaps []IVaultBatchSwapStep, assets []common.Address, funds IVaultFundManagement) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "queryBatchSwap", kind, swaps, assets, funds)
}

// QueryBatchSwap is a paid mutator transaction binding the contract method 0xf84d066e.
//
// Solidity: function queryBatchSwap(uint8 kind, (bytes32,uint256,uint256,uint256,bytes)[] swaps, address[] assets, (address,bool,address,bool) funds) returns(int256[])
func (_Vault *VaultSession) QueryBatchSwap(kind uint8, swaps []IVaultBatchSwapStep, assets []common.Address, funds IVaultFundManagement) (*types.Transaction, error) {
	return _Vault.Contract.QueryBatchSwap(&_Vault.TransactOpts, kind, swaps, assets, funds)
}

// QueryBatchSwap is a paid mutator transaction binding the contract method 0xf84d066e.
//
// Solidity: function queryBatchSwap(uint8 kind, (bytes32,uint256,uint256,uint256,bytes)[] swaps, address[] assets, (address,bool,address,bool) funds) returns(int256[])
func (_Vault *VaultTransactorSession) QueryBatchSwap(kind uint8, swaps []IVaultBatchSwapStep, assets []common.Address, funds IVaultFundManagement) (*types.Transaction, error) {
	return _Vault.Contract.QueryBatchSwap(&_Vault.TransactOpts, kind, swaps, assets, funds)
}

// RegisterPool is a paid mutator transaction binding the contract method 0x09b2760f.
//
// Solidity: function registerPool(uint8 specialization) returns(bytes32)
func (_Vault *VaultTransactor) RegisterPool(opts *bind.TransactOpts, specialization uint8) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "registerPool", specialization)
}

// RegisterPool is a paid mutator transaction binding the contract method 0x09b2760f.
//
// Solidity: function registerPool(uint8 specialization) returns(bytes32)
func (_Vault *VaultSession) RegisterPool(specialization uint8) (*types.Transaction, error) {
	return _Vault.Contract.RegisterPool(&_Vault.TransactOpts, specialization)
}

// RegisterPool is a paid mutator transaction binding the contract method 0x09b2760f.
//
// Solidity: function registerPool(uint8 specialization) returns(bytes32)
func (_Vault *VaultTransactorSession) RegisterPool(specialization uint8) (*types.Transaction, error) {
	return _Vault.Contract.RegisterPool(&_Vault.TransactOpts, specialization)
}

// RegisterTokens is a paid mutator transaction binding the contract method 0x66a9c7d2.
//
// Solidity: function registerTokens(bytes32 poolId, address[] tokens, address[] assetManagers) returns()
func (_Vault *VaultTransactor) RegisterTokens(opts *bind.TransactOpts, poolId [32]byte, tokens []common.Address, assetManagers []common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "registerTokens", poolId, tokens, assetManagers)
}

// RegisterTokens is a paid mutator transaction binding the contract method 0x66a9c7d2.
//
// Solidity: function registerTokens(bytes32 poolId, address[] tokens, address[] assetManagers) returns()
func (_Vault *VaultSession) RegisterTokens(poolId [32]byte, tokens []common.Address, assetManagers []common.Address) (*types.Transaction, error) {
	return _Vault.Contract.RegisterTokens(&_Vault.TransactOpts, poolId, tokens, assetManagers)
}

// RegisterTokens is a paid mutator transaction binding the contract method 0x66a9c7d2.
//
// Solidity: function registerTokens(bytes32 poolId, address[] tokens, address[] assetManagers) returns()
func (_Vault *VaultTransactorSession) RegisterTokens(poolId [32]byte, tokens []common.Address, assetManagers []common.Address) (*types.Transaction, error) {
	return _Vault.Contract.RegisterTokens(&_Vault.TransactOpts, poolId, tokens, assetManagers)
}

// SetAuthorizer is a paid mutator transaction binding the contract method 0x058a628f.
//
// Solidity: function setAuthorizer(address newAuthorizer) returns()
func (_Vault *VaultTransactor) SetAuthorizer(opts *bind.TransactOpts, newAuthorizer common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "setAuthorizer", newAuthorizer)
}

// SetAuthorizer is a paid mutator transaction binding the contract method 0x058a628f.
//
// Solidity: function setAuthorizer(address newAuthorizer) returns()
func (_Vault *VaultSession) SetAuthorizer(newAuthorizer common.Address) (*types.Transaction, error) {
	return _Vault.Contract.SetAuthorizer(&_Vault.TransactOpts, newAuthorizer)
}

// SetAuthorizer is a paid mutator transaction binding the contract method 0x058a628f.
//
// Solidity: function setAuthorizer(address newAuthorizer) returns()
func (_Vault *VaultTransactorSession) SetAuthorizer(newAuthorizer common.Address) (*types.Transaction, error) {
	return _Vault.Contract.SetAuthorizer(&_Vault.TransactOpts, newAuthorizer)
}

// SetPaused is a paid mutator transaction binding the contract method 0x16c38b3c.
//
// Solidity: function setPaused(bool paused) returns()
func (_Vault *VaultTransactor) SetPaused(opts *bind.TransactOpts, paused bool) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "setPaused", paused)
}

// SetPaused is a paid mutator transaction binding the contract method 0x16c38b3c.
//
// Solidity: function setPaused(bool paused) returns()
func (_Vault *VaultSession) SetPaused(paused bool) (*types.Transaction, error) {
	return _Vault.Contract.SetPaused(&_Vault.TransactOpts, paused)
}

// SetPaused is a paid mutator transaction binding the contract method 0x16c38b3c.
//
// Solidity: function setPaused(bool paused) returns()
func (_Vault *VaultTransactorSession) SetPaused(paused bool) (*types.Transaction, error) {
	return _Vault.Contract.SetPaused(&_Vault.TransactOpts, paused)
}

// SetRelayerApproval is a paid mutator transaction binding the contract method 0xfa6e671d.
//
// Solidity: function setRelayerApproval(address sender, address relayer, bool approved) returns()
func (_Vault *VaultTransactor) SetRelayerApproval(opts *bind.TransactOpts, sender common.Address, relayer common.Address, approved bool) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "setRelayerApproval", sender, relayer, approved)
}

// SetRelayerApproval is a paid mutator transaction binding the contract method 0xfa6e671d.
//
// Solidity: function setRelayerApproval(address sender, address relayer, bool approved) returns()
func (_Vault *VaultSession) SetRelayerApproval(sender common.Address, relayer common.Address, approved bool) (*types.Transaction, error) {
	return _Vault.Contract.SetRelayerApproval(&_Vault.TransactOpts, sender, relayer, approved)
}

// SetRelayerApproval is a paid mutator transaction binding the contract method 0xfa6e671d.
//
// Solidity: function setRelayerApproval(address sender, address relayer, bool approved) returns()
func (_Vault *VaultTransactorSession) SetRelayerApproval(sender common.Address, relayer common.Address, approved bool) (*types.Transaction, error) {
	return _Vault.Contract.SetRelayerApproval(&_Vault.TransactOpts, sender, relayer, approved)
}

// Swap is a paid mutator transaction binding the contract method 0x52bbbe29.
//
// Solidity: function swap((bytes32,uint8,address,address,uint256,bytes) singleSwap, (address,bool,address,bool) funds, uint256 limit, uint256 deadline) payable returns(uint256 amountCalculated)
func (_Vault *VaultTransactor) Swap(opts *bind.TransactOpts, singleSwap IVaultSingleSwap, funds IVaultFundManagement, limit *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "swap", singleSwap, funds, limit, deadline)
}

// Swap is a paid mutator transaction binding the contract method 0x52bbbe29.
//
// Solidity: function swap((bytes32,uint8,address,address,uint256,bytes) singleSwap, (address,bool,address,bool) funds, uint256 limit, uint256 deadline) payable returns(uint256 amountCalculated)
func (_Vault *VaultSession) Swap(singleSwap IVaultSingleSwap, funds IVaultFundManagement, limit *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.Swap(&_Vault.TransactOpts, singleSwap, funds, limit, deadline)
}

// Swap is a paid mutator transaction binding the contract method 0x52bbbe29.
//
// Solidity: function swap((bytes32,uint8,address,address,uint256,bytes) singleSwap, (address,bool,address,bool) funds, uint256 limit, uint256 deadline) payable returns(uint256 amountCalculated)
func (_Vault *VaultTransactorSession) Swap(singleSwap IVaultSingleSwap, funds IVaultFundManagement, limit *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.Swap(&_Vault.TransactOpts, singleSwap, funds, limit, deadline)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Vault *VaultTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vault.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Vault *VaultSession) Receive() (*types.Transaction, error) {
	return _Vault.Contract.Receive(&_Vault.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Vault *VaultTransactorSession) Receive() (*types.Transaction, error) {
	return _Vault.Contract.Receive(&_Vault.TransactOpts)
}

// VaultAuthorizerChangedIterator is returned from FilterAuthorizerChanged and is used to iterate over the raw logs and unpacked data for AuthorizerChanged events raised by the Vault contract.
type VaultAuthorizerChangedIterator struct {
	Event *VaultAuthorizerChanged // Event containing the contract specifics and raw log

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
func (it *VaultAuthorizerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultAuthorizerChanged)
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
		it.Event = new(VaultAuthorizerChanged)
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
func (it *VaultAuthorizerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultAuthorizerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultAuthorizerChanged represents a AuthorizerChanged event raised by the Vault contract.
type VaultAuthorizerChanged struct {
	NewAuthorizer common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAuthorizerChanged is a free log retrieval operation binding the contract event 0x94b979b6831a51293e2641426f97747feed46f17779fed9cd18d1ecefcfe92ef.
//
// Solidity: event AuthorizerChanged(address indexed newAuthorizer)
func (_Vault *VaultFilterer) FilterAuthorizerChanged(opts *bind.FilterOpts, newAuthorizer []common.Address) (*VaultAuthorizerChangedIterator, error) {

	var newAuthorizerRule []interface{}
	for _, newAuthorizerItem := range newAuthorizer {
		newAuthorizerRule = append(newAuthorizerRule, newAuthorizerItem)
	}

	logs, sub, err := _Vault.contract.FilterLogs(opts, "AuthorizerChanged", newAuthorizerRule)
	if err != nil {
		return nil, err
	}
	return &VaultAuthorizerChangedIterator{contract: _Vault.contract, event: "AuthorizerChanged", logs: logs, sub: sub}, nil
}

// WatchAuthorizerChanged is a free log subscription operation binding the contract event 0x94b979b6831a51293e2641426f97747feed46f17779fed9cd18d1ecefcfe92ef.
//
// Solidity: event AuthorizerChanged(address indexed newAuthorizer)
func (_Vault *VaultFilterer) WatchAuthorizerChanged(opts *bind.WatchOpts, sink chan<- *VaultAuthorizerChanged, newAuthorizer []common.Address) (event.Subscription, error) {

	var newAuthorizerRule []interface{}
	for _, newAuthorizerItem := range newAuthorizer {
		newAuthorizerRule = append(newAuthorizerRule, newAuthorizerItem)
	}

	logs, sub, err := _Vault.contract.WatchLogs(opts, "AuthorizerChanged", newAuthorizerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultAuthorizerChanged)
				if err := _Vault.contract.UnpackLog(event, "AuthorizerChanged", log); err != nil {
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

// ParseAuthorizerChanged is a log parse operation binding the contract event 0x94b979b6831a51293e2641426f97747feed46f17779fed9cd18d1ecefcfe92ef.
//
// Solidity: event AuthorizerChanged(address indexed newAuthorizer)
func (_Vault *VaultFilterer) ParseAuthorizerChanged(log types.Log) (*VaultAuthorizerChanged, error) {
	event := new(VaultAuthorizerChanged)
	if err := _Vault.contract.UnpackLog(event, "AuthorizerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultExternalBalanceTransferIterator is returned from FilterExternalBalanceTransfer and is used to iterate over the raw logs and unpacked data for ExternalBalanceTransfer events raised by the Vault contract.
type VaultExternalBalanceTransferIterator struct {
	Event *VaultExternalBalanceTransfer // Event containing the contract specifics and raw log

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
func (it *VaultExternalBalanceTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultExternalBalanceTransfer)
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
		it.Event = new(VaultExternalBalanceTransfer)
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
func (it *VaultExternalBalanceTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultExternalBalanceTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultExternalBalanceTransfer represents a ExternalBalanceTransfer event raised by the Vault contract.
type VaultExternalBalanceTransfer struct {
	Token     common.Address
	Sender    common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterExternalBalanceTransfer is a free log retrieval operation binding the contract event 0x540a1a3f28340caec336c81d8d7b3df139ee5cdc1839a4f283d7ebb7eaae2d5c.
//
// Solidity: event ExternalBalanceTransfer(address indexed token, address indexed sender, address recipient, uint256 amount)
func (_Vault *VaultFilterer) FilterExternalBalanceTransfer(opts *bind.FilterOpts, token []common.Address, sender []common.Address) (*VaultExternalBalanceTransferIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Vault.contract.FilterLogs(opts, "ExternalBalanceTransfer", tokenRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &VaultExternalBalanceTransferIterator{contract: _Vault.contract, event: "ExternalBalanceTransfer", logs: logs, sub: sub}, nil
}

// WatchExternalBalanceTransfer is a free log subscription operation binding the contract event 0x540a1a3f28340caec336c81d8d7b3df139ee5cdc1839a4f283d7ebb7eaae2d5c.
//
// Solidity: event ExternalBalanceTransfer(address indexed token, address indexed sender, address recipient, uint256 amount)
func (_Vault *VaultFilterer) WatchExternalBalanceTransfer(opts *bind.WatchOpts, sink chan<- *VaultExternalBalanceTransfer, token []common.Address, sender []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Vault.contract.WatchLogs(opts, "ExternalBalanceTransfer", tokenRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultExternalBalanceTransfer)
				if err := _Vault.contract.UnpackLog(event, "ExternalBalanceTransfer", log); err != nil {
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

// ParseExternalBalanceTransfer is a log parse operation binding the contract event 0x540a1a3f28340caec336c81d8d7b3df139ee5cdc1839a4f283d7ebb7eaae2d5c.
//
// Solidity: event ExternalBalanceTransfer(address indexed token, address indexed sender, address recipient, uint256 amount)
func (_Vault *VaultFilterer) ParseExternalBalanceTransfer(log types.Log) (*VaultExternalBalanceTransfer, error) {
	event := new(VaultExternalBalanceTransfer)
	if err := _Vault.contract.UnpackLog(event, "ExternalBalanceTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultFlashLoanIterator is returned from FilterFlashLoan and is used to iterate over the raw logs and unpacked data for FlashLoan events raised by the Vault contract.
type VaultFlashLoanIterator struct {
	Event *VaultFlashLoan // Event containing the contract specifics and raw log

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
func (it *VaultFlashLoanIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultFlashLoan)
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
		it.Event = new(VaultFlashLoan)
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
func (it *VaultFlashLoanIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultFlashLoanIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultFlashLoan represents a FlashLoan event raised by the Vault contract.
type VaultFlashLoan struct {
	Recipient common.Address
	Token     common.Address
	Amount    *big.Int
	FeeAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFlashLoan is a free log retrieval operation binding the contract event 0x0d7d75e01ab95780d3cd1c8ec0dd6c2ce19e3a20427eec8bf53283b6fb8e95f0.
//
// Solidity: event FlashLoan(address indexed recipient, address indexed token, uint256 amount, uint256 feeAmount)
func (_Vault *VaultFilterer) FilterFlashLoan(opts *bind.FilterOpts, recipient []common.Address, token []common.Address) (*VaultFlashLoanIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Vault.contract.FilterLogs(opts, "FlashLoan", recipientRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &VaultFlashLoanIterator{contract: _Vault.contract, event: "FlashLoan", logs: logs, sub: sub}, nil
}

// WatchFlashLoan is a free log subscription operation binding the contract event 0x0d7d75e01ab95780d3cd1c8ec0dd6c2ce19e3a20427eec8bf53283b6fb8e95f0.
//
// Solidity: event FlashLoan(address indexed recipient, address indexed token, uint256 amount, uint256 feeAmount)
func (_Vault *VaultFilterer) WatchFlashLoan(opts *bind.WatchOpts, sink chan<- *VaultFlashLoan, recipient []common.Address, token []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Vault.contract.WatchLogs(opts, "FlashLoan", recipientRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultFlashLoan)
				if err := _Vault.contract.UnpackLog(event, "FlashLoan", log); err != nil {
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

// ParseFlashLoan is a log parse operation binding the contract event 0x0d7d75e01ab95780d3cd1c8ec0dd6c2ce19e3a20427eec8bf53283b6fb8e95f0.
//
// Solidity: event FlashLoan(address indexed recipient, address indexed token, uint256 amount, uint256 feeAmount)
func (_Vault *VaultFilterer) ParseFlashLoan(log types.Log) (*VaultFlashLoan, error) {
	event := new(VaultFlashLoan)
	if err := _Vault.contract.UnpackLog(event, "FlashLoan", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultInternalBalanceChangedIterator is returned from FilterInternalBalanceChanged and is used to iterate over the raw logs and unpacked data for InternalBalanceChanged events raised by the Vault contract.
type VaultInternalBalanceChangedIterator struct {
	Event *VaultInternalBalanceChanged // Event containing the contract specifics and raw log

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
func (it *VaultInternalBalanceChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultInternalBalanceChanged)
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
		it.Event = new(VaultInternalBalanceChanged)
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
func (it *VaultInternalBalanceChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultInternalBalanceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultInternalBalanceChanged represents a InternalBalanceChanged event raised by the Vault contract.
type VaultInternalBalanceChanged struct {
	User  common.Address
	Token common.Address
	Delta *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterInternalBalanceChanged is a free log retrieval operation binding the contract event 0x18e1ea4139e68413d7d08aa752e71568e36b2c5bf940893314c2c5b01eaa0c42.
//
// Solidity: event InternalBalanceChanged(address indexed user, address indexed token, int256 delta)
func (_Vault *VaultFilterer) FilterInternalBalanceChanged(opts *bind.FilterOpts, user []common.Address, token []common.Address) (*VaultInternalBalanceChangedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Vault.contract.FilterLogs(opts, "InternalBalanceChanged", userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &VaultInternalBalanceChangedIterator{contract: _Vault.contract, event: "InternalBalanceChanged", logs: logs, sub: sub}, nil
}

// WatchInternalBalanceChanged is a free log subscription operation binding the contract event 0x18e1ea4139e68413d7d08aa752e71568e36b2c5bf940893314c2c5b01eaa0c42.
//
// Solidity: event InternalBalanceChanged(address indexed user, address indexed token, int256 delta)
func (_Vault *VaultFilterer) WatchInternalBalanceChanged(opts *bind.WatchOpts, sink chan<- *VaultInternalBalanceChanged, user []common.Address, token []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Vault.contract.WatchLogs(opts, "InternalBalanceChanged", userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultInternalBalanceChanged)
				if err := _Vault.contract.UnpackLog(event, "InternalBalanceChanged", log); err != nil {
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

// ParseInternalBalanceChanged is a log parse operation binding the contract event 0x18e1ea4139e68413d7d08aa752e71568e36b2c5bf940893314c2c5b01eaa0c42.
//
// Solidity: event InternalBalanceChanged(address indexed user, address indexed token, int256 delta)
func (_Vault *VaultFilterer) ParseInternalBalanceChanged(log types.Log) (*VaultInternalBalanceChanged, error) {
	event := new(VaultInternalBalanceChanged)
	if err := _Vault.contract.UnpackLog(event, "InternalBalanceChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultPausedStateChangedIterator is returned from FilterPausedStateChanged and is used to iterate over the raw logs and unpacked data for PausedStateChanged events raised by the Vault contract.
type VaultPausedStateChangedIterator struct {
	Event *VaultPausedStateChanged // Event containing the contract specifics and raw log

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
func (it *VaultPausedStateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultPausedStateChanged)
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
		it.Event = new(VaultPausedStateChanged)
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
func (it *VaultPausedStateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultPausedStateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultPausedStateChanged represents a PausedStateChanged event raised by the Vault contract.
type VaultPausedStateChanged struct {
	Paused bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPausedStateChanged is a free log retrieval operation binding the contract event 0x9e3a5e37224532dea67b89face185703738a228a6e8a23dee546960180d3be64.
//
// Solidity: event PausedStateChanged(bool paused)
func (_Vault *VaultFilterer) FilterPausedStateChanged(opts *bind.FilterOpts) (*VaultPausedStateChangedIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "PausedStateChanged")
	if err != nil {
		return nil, err
	}
	return &VaultPausedStateChangedIterator{contract: _Vault.contract, event: "PausedStateChanged", logs: logs, sub: sub}, nil
}

// WatchPausedStateChanged is a free log subscription operation binding the contract event 0x9e3a5e37224532dea67b89face185703738a228a6e8a23dee546960180d3be64.
//
// Solidity: event PausedStateChanged(bool paused)
func (_Vault *VaultFilterer) WatchPausedStateChanged(opts *bind.WatchOpts, sink chan<- *VaultPausedStateChanged) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "PausedStateChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultPausedStateChanged)
				if err := _Vault.contract.UnpackLog(event, "PausedStateChanged", log); err != nil {
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

// ParsePausedStateChanged is a log parse operation binding the contract event 0x9e3a5e37224532dea67b89face185703738a228a6e8a23dee546960180d3be64.
//
// Solidity: event PausedStateChanged(bool paused)
func (_Vault *VaultFilterer) ParsePausedStateChanged(log types.Log) (*VaultPausedStateChanged, error) {
	event := new(VaultPausedStateChanged)
	if err := _Vault.contract.UnpackLog(event, "PausedStateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultPoolBalanceChangedIterator is returned from FilterPoolBalanceChanged and is used to iterate over the raw logs and unpacked data for PoolBalanceChanged events raised by the Vault contract.
type VaultPoolBalanceChangedIterator struct {
	Event *VaultPoolBalanceChanged // Event containing the contract specifics and raw log

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
func (it *VaultPoolBalanceChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultPoolBalanceChanged)
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
		it.Event = new(VaultPoolBalanceChanged)
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
func (it *VaultPoolBalanceChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultPoolBalanceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultPoolBalanceChanged represents a PoolBalanceChanged event raised by the Vault contract.
type VaultPoolBalanceChanged struct {
	PoolId             [32]byte
	LiquidityProvider  common.Address
	Tokens             []common.Address
	Deltas             []*big.Int
	ProtocolFeeAmounts []*big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterPoolBalanceChanged is a free log retrieval operation binding the contract event 0xe5ce249087ce04f05a957192435400fd97868dba0e6a4b4c049abf8af80dae78.
//
// Solidity: event PoolBalanceChanged(bytes32 indexed poolId, address indexed liquidityProvider, address[] tokens, int256[] deltas, uint256[] protocolFeeAmounts)
func (_Vault *VaultFilterer) FilterPoolBalanceChanged(opts *bind.FilterOpts, poolId [][32]byte, liquidityProvider []common.Address) (*VaultPoolBalanceChangedIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var liquidityProviderRule []interface{}
	for _, liquidityProviderItem := range liquidityProvider {
		liquidityProviderRule = append(liquidityProviderRule, liquidityProviderItem)
	}

	logs, sub, err := _Vault.contract.FilterLogs(opts, "PoolBalanceChanged", poolIdRule, liquidityProviderRule)
	if err != nil {
		return nil, err
	}
	return &VaultPoolBalanceChangedIterator{contract: _Vault.contract, event: "PoolBalanceChanged", logs: logs, sub: sub}, nil
}

// WatchPoolBalanceChanged is a free log subscription operation binding the contract event 0xe5ce249087ce04f05a957192435400fd97868dba0e6a4b4c049abf8af80dae78.
//
// Solidity: event PoolBalanceChanged(bytes32 indexed poolId, address indexed liquidityProvider, address[] tokens, int256[] deltas, uint256[] protocolFeeAmounts)
func (_Vault *VaultFilterer) WatchPoolBalanceChanged(opts *bind.WatchOpts, sink chan<- *VaultPoolBalanceChanged, poolId [][32]byte, liquidityProvider []common.Address) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var liquidityProviderRule []interface{}
	for _, liquidityProviderItem := range liquidityProvider {
		liquidityProviderRule = append(liquidityProviderRule, liquidityProviderItem)
	}

	logs, sub, err := _Vault.contract.WatchLogs(opts, "PoolBalanceChanged", poolIdRule, liquidityProviderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultPoolBalanceChanged)
				if err := _Vault.contract.UnpackLog(event, "PoolBalanceChanged", log); err != nil {
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

// ParsePoolBalanceChanged is a log parse operation binding the contract event 0xe5ce249087ce04f05a957192435400fd97868dba0e6a4b4c049abf8af80dae78.
//
// Solidity: event PoolBalanceChanged(bytes32 indexed poolId, address indexed liquidityProvider, address[] tokens, int256[] deltas, uint256[] protocolFeeAmounts)
func (_Vault *VaultFilterer) ParsePoolBalanceChanged(log types.Log) (*VaultPoolBalanceChanged, error) {
	event := new(VaultPoolBalanceChanged)
	if err := _Vault.contract.UnpackLog(event, "PoolBalanceChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultPoolBalanceManagedIterator is returned from FilterPoolBalanceManaged and is used to iterate over the raw logs and unpacked data for PoolBalanceManaged events raised by the Vault contract.
type VaultPoolBalanceManagedIterator struct {
	Event *VaultPoolBalanceManaged // Event containing the contract specifics and raw log

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
func (it *VaultPoolBalanceManagedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultPoolBalanceManaged)
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
		it.Event = new(VaultPoolBalanceManaged)
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
func (it *VaultPoolBalanceManagedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultPoolBalanceManagedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultPoolBalanceManaged represents a PoolBalanceManaged event raised by the Vault contract.
type VaultPoolBalanceManaged struct {
	PoolId       [32]byte
	AssetManager common.Address
	Token        common.Address
	CashDelta    *big.Int
	ManagedDelta *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterPoolBalanceManaged is a free log retrieval operation binding the contract event 0x6edcaf6241105b4c94c2efdbf3a6b12458eb3d07be3a0e81d24b13c44045fe7a.
//
// Solidity: event PoolBalanceManaged(bytes32 indexed poolId, address indexed assetManager, address indexed token, int256 cashDelta, int256 managedDelta)
func (_Vault *VaultFilterer) FilterPoolBalanceManaged(opts *bind.FilterOpts, poolId [][32]byte, assetManager []common.Address, token []common.Address) (*VaultPoolBalanceManagedIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var assetManagerRule []interface{}
	for _, assetManagerItem := range assetManager {
		assetManagerRule = append(assetManagerRule, assetManagerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Vault.contract.FilterLogs(opts, "PoolBalanceManaged", poolIdRule, assetManagerRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &VaultPoolBalanceManagedIterator{contract: _Vault.contract, event: "PoolBalanceManaged", logs: logs, sub: sub}, nil
}

// WatchPoolBalanceManaged is a free log subscription operation binding the contract event 0x6edcaf6241105b4c94c2efdbf3a6b12458eb3d07be3a0e81d24b13c44045fe7a.
//
// Solidity: event PoolBalanceManaged(bytes32 indexed poolId, address indexed assetManager, address indexed token, int256 cashDelta, int256 managedDelta)
func (_Vault *VaultFilterer) WatchPoolBalanceManaged(opts *bind.WatchOpts, sink chan<- *VaultPoolBalanceManaged, poolId [][32]byte, assetManager []common.Address, token []common.Address) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var assetManagerRule []interface{}
	for _, assetManagerItem := range assetManager {
		assetManagerRule = append(assetManagerRule, assetManagerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Vault.contract.WatchLogs(opts, "PoolBalanceManaged", poolIdRule, assetManagerRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultPoolBalanceManaged)
				if err := _Vault.contract.UnpackLog(event, "PoolBalanceManaged", log); err != nil {
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

// ParsePoolBalanceManaged is a log parse operation binding the contract event 0x6edcaf6241105b4c94c2efdbf3a6b12458eb3d07be3a0e81d24b13c44045fe7a.
//
// Solidity: event PoolBalanceManaged(bytes32 indexed poolId, address indexed assetManager, address indexed token, int256 cashDelta, int256 managedDelta)
func (_Vault *VaultFilterer) ParsePoolBalanceManaged(log types.Log) (*VaultPoolBalanceManaged, error) {
	event := new(VaultPoolBalanceManaged)
	if err := _Vault.contract.UnpackLog(event, "PoolBalanceManaged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultPoolRegisteredIterator is returned from FilterPoolRegistered and is used to iterate over the raw logs and unpacked data for PoolRegistered events raised by the Vault contract.
type VaultPoolRegisteredIterator struct {
	Event *VaultPoolRegistered // Event containing the contract specifics and raw log

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
func (it *VaultPoolRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultPoolRegistered)
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
		it.Event = new(VaultPoolRegistered)
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
func (it *VaultPoolRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultPoolRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultPoolRegistered represents a PoolRegistered event raised by the Vault contract.
type VaultPoolRegistered struct {
	PoolId         [32]byte
	PoolAddress    common.Address
	Specialization uint8
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterPoolRegistered is a free log retrieval operation binding the contract event 0x3c13bc30b8e878c53fd2a36b679409c073afd75950be43d8858768e956fbc20e.
//
// Solidity: event PoolRegistered(bytes32 indexed poolId, address indexed poolAddress, uint8 specialization)
func (_Vault *VaultFilterer) FilterPoolRegistered(opts *bind.FilterOpts, poolId [][32]byte, poolAddress []common.Address) (*VaultPoolRegisteredIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var poolAddressRule []interface{}
	for _, poolAddressItem := range poolAddress {
		poolAddressRule = append(poolAddressRule, poolAddressItem)
	}

	logs, sub, err := _Vault.contract.FilterLogs(opts, "PoolRegistered", poolIdRule, poolAddressRule)
	if err != nil {
		return nil, err
	}
	return &VaultPoolRegisteredIterator{contract: _Vault.contract, event: "PoolRegistered", logs: logs, sub: sub}, nil
}

// WatchPoolRegistered is a free log subscription operation binding the contract event 0x3c13bc30b8e878c53fd2a36b679409c073afd75950be43d8858768e956fbc20e.
//
// Solidity: event PoolRegistered(bytes32 indexed poolId, address indexed poolAddress, uint8 specialization)
func (_Vault *VaultFilterer) WatchPoolRegistered(opts *bind.WatchOpts, sink chan<- *VaultPoolRegistered, poolId [][32]byte, poolAddress []common.Address) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var poolAddressRule []interface{}
	for _, poolAddressItem := range poolAddress {
		poolAddressRule = append(poolAddressRule, poolAddressItem)
	}

	logs, sub, err := _Vault.contract.WatchLogs(opts, "PoolRegistered", poolIdRule, poolAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultPoolRegistered)
				if err := _Vault.contract.UnpackLog(event, "PoolRegistered", log); err != nil {
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

// ParsePoolRegistered is a log parse operation binding the contract event 0x3c13bc30b8e878c53fd2a36b679409c073afd75950be43d8858768e956fbc20e.
//
// Solidity: event PoolRegistered(bytes32 indexed poolId, address indexed poolAddress, uint8 specialization)
func (_Vault *VaultFilterer) ParsePoolRegistered(log types.Log) (*VaultPoolRegistered, error) {
	event := new(VaultPoolRegistered)
	if err := _Vault.contract.UnpackLog(event, "PoolRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultRelayerApprovalChangedIterator is returned from FilterRelayerApprovalChanged and is used to iterate over the raw logs and unpacked data for RelayerApprovalChanged events raised by the Vault contract.
type VaultRelayerApprovalChangedIterator struct {
	Event *VaultRelayerApprovalChanged // Event containing the contract specifics and raw log

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
func (it *VaultRelayerApprovalChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultRelayerApprovalChanged)
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
		it.Event = new(VaultRelayerApprovalChanged)
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
func (it *VaultRelayerApprovalChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultRelayerApprovalChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultRelayerApprovalChanged represents a RelayerApprovalChanged event raised by the Vault contract.
type VaultRelayerApprovalChanged struct {
	Relayer  common.Address
	Sender   common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRelayerApprovalChanged is a free log retrieval operation binding the contract event 0x46961fdb4502b646d5095fba7600486a8ac05041d55cdf0f16ed677180b5cad8.
//
// Solidity: event RelayerApprovalChanged(address indexed relayer, address indexed sender, bool approved)
func (_Vault *VaultFilterer) FilterRelayerApprovalChanged(opts *bind.FilterOpts, relayer []common.Address, sender []common.Address) (*VaultRelayerApprovalChangedIterator, error) {

	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Vault.contract.FilterLogs(opts, "RelayerApprovalChanged", relayerRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &VaultRelayerApprovalChangedIterator{contract: _Vault.contract, event: "RelayerApprovalChanged", logs: logs, sub: sub}, nil
}

// WatchRelayerApprovalChanged is a free log subscription operation binding the contract event 0x46961fdb4502b646d5095fba7600486a8ac05041d55cdf0f16ed677180b5cad8.
//
// Solidity: event RelayerApprovalChanged(address indexed relayer, address indexed sender, bool approved)
func (_Vault *VaultFilterer) WatchRelayerApprovalChanged(opts *bind.WatchOpts, sink chan<- *VaultRelayerApprovalChanged, relayer []common.Address, sender []common.Address) (event.Subscription, error) {

	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Vault.contract.WatchLogs(opts, "RelayerApprovalChanged", relayerRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultRelayerApprovalChanged)
				if err := _Vault.contract.UnpackLog(event, "RelayerApprovalChanged", log); err != nil {
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

// ParseRelayerApprovalChanged is a log parse operation binding the contract event 0x46961fdb4502b646d5095fba7600486a8ac05041d55cdf0f16ed677180b5cad8.
//
// Solidity: event RelayerApprovalChanged(address indexed relayer, address indexed sender, bool approved)
func (_Vault *VaultFilterer) ParseRelayerApprovalChanged(log types.Log) (*VaultRelayerApprovalChanged, error) {
	event := new(VaultRelayerApprovalChanged)
	if err := _Vault.contract.UnpackLog(event, "RelayerApprovalChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the Vault contract.
type VaultSwapIterator struct {
	Event *VaultSwap // Event containing the contract specifics and raw log

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
func (it *VaultSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultSwap)
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
		it.Event = new(VaultSwap)
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
func (it *VaultSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultSwap represents a Swap event raised by the Vault contract.
type VaultSwap struct {
	PoolId    [32]byte
	TokenIn   common.Address
	TokenOut  common.Address
	AmountIn  *big.Int
	AmountOut *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0x2170c741c41531aec20e7c107c24eecfdd15e69c9bb0a8dd37b1840b9e0b207b.
//
// Solidity: event Swap(bytes32 indexed poolId, address indexed tokenIn, address indexed tokenOut, uint256 amountIn, uint256 amountOut)
func (_Vault *VaultFilterer) FilterSwap(opts *bind.FilterOpts, poolId [][32]byte, tokenIn []common.Address, tokenOut []common.Address) (*VaultSwapIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var tokenInRule []interface{}
	for _, tokenInItem := range tokenIn {
		tokenInRule = append(tokenInRule, tokenInItem)
	}
	var tokenOutRule []interface{}
	for _, tokenOutItem := range tokenOut {
		tokenOutRule = append(tokenOutRule, tokenOutItem)
	}

	logs, sub, err := _Vault.contract.FilterLogs(opts, "Swap", poolIdRule, tokenInRule, tokenOutRule)
	if err != nil {
		return nil, err
	}
	return &VaultSwapIterator{contract: _Vault.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0x2170c741c41531aec20e7c107c24eecfdd15e69c9bb0a8dd37b1840b9e0b207b.
//
// Solidity: event Swap(bytes32 indexed poolId, address indexed tokenIn, address indexed tokenOut, uint256 amountIn, uint256 amountOut)
func (_Vault *VaultFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *VaultSwap, poolId [][32]byte, tokenIn []common.Address, tokenOut []common.Address) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var tokenInRule []interface{}
	for _, tokenInItem := range tokenIn {
		tokenInRule = append(tokenInRule, tokenInItem)
	}
	var tokenOutRule []interface{}
	for _, tokenOutItem := range tokenOut {
		tokenOutRule = append(tokenOutRule, tokenOutItem)
	}

	logs, sub, err := _Vault.contract.WatchLogs(opts, "Swap", poolIdRule, tokenInRule, tokenOutRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultSwap)
				if err := _Vault.contract.UnpackLog(event, "Swap", log); err != nil {
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

// ParseSwap is a log parse operation binding the contract event 0x2170c741c41531aec20e7c107c24eecfdd15e69c9bb0a8dd37b1840b9e0b207b.
//
// Solidity: event Swap(bytes32 indexed poolId, address indexed tokenIn, address indexed tokenOut, uint256 amountIn, uint256 amountOut)
func (_Vault *VaultFilterer) ParseSwap(log types.Log) (*VaultSwap, error) {
	event := new(VaultSwap)
	if err := _Vault.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultTokensDeregisteredIterator is returned from FilterTokensDeregistered and is used to iterate over the raw logs and unpacked data for TokensDeregistered events raised by the Vault contract.
type VaultTokensDeregisteredIterator struct {
	Event *VaultTokensDeregistered // Event containing the contract specifics and raw log

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
func (it *VaultTokensDeregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultTokensDeregistered)
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
		it.Event = new(VaultTokensDeregistered)
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
func (it *VaultTokensDeregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultTokensDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultTokensDeregistered represents a TokensDeregistered event raised by the Vault contract.
type VaultTokensDeregistered struct {
	PoolId [32]byte
	Tokens []common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokensDeregistered is a free log retrieval operation binding the contract event 0x7dcdc6d02ef40c7c1a7046a011b058bd7f988fa14e20a66344f9d4e60657d610.
//
// Solidity: event TokensDeregistered(bytes32 indexed poolId, address[] tokens)
func (_Vault *VaultFilterer) FilterTokensDeregistered(opts *bind.FilterOpts, poolId [][32]byte) (*VaultTokensDeregisteredIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _Vault.contract.FilterLogs(opts, "TokensDeregistered", poolIdRule)
	if err != nil {
		return nil, err
	}
	return &VaultTokensDeregisteredIterator{contract: _Vault.contract, event: "TokensDeregistered", logs: logs, sub: sub}, nil
}

// WatchTokensDeregistered is a free log subscription operation binding the contract event 0x7dcdc6d02ef40c7c1a7046a011b058bd7f988fa14e20a66344f9d4e60657d610.
//
// Solidity: event TokensDeregistered(bytes32 indexed poolId, address[] tokens)
func (_Vault *VaultFilterer) WatchTokensDeregistered(opts *bind.WatchOpts, sink chan<- *VaultTokensDeregistered, poolId [][32]byte) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _Vault.contract.WatchLogs(opts, "TokensDeregistered", poolIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultTokensDeregistered)
				if err := _Vault.contract.UnpackLog(event, "TokensDeregistered", log); err != nil {
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

// ParseTokensDeregistered is a log parse operation binding the contract event 0x7dcdc6d02ef40c7c1a7046a011b058bd7f988fa14e20a66344f9d4e60657d610.
//
// Solidity: event TokensDeregistered(bytes32 indexed poolId, address[] tokens)
func (_Vault *VaultFilterer) ParseTokensDeregistered(log types.Log) (*VaultTokensDeregistered, error) {
	event := new(VaultTokensDeregistered)
	if err := _Vault.contract.UnpackLog(event, "TokensDeregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultTokensRegisteredIterator is returned from FilterTokensRegistered and is used to iterate over the raw logs and unpacked data for TokensRegistered events raised by the Vault contract.
type VaultTokensRegisteredIterator struct {
	Event *VaultTokensRegistered // Event containing the contract specifics and raw log

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
func (it *VaultTokensRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultTokensRegistered)
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
		it.Event = new(VaultTokensRegistered)
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
func (it *VaultTokensRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultTokensRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultTokensRegistered represents a TokensRegistered event raised by the Vault contract.
type VaultTokensRegistered struct {
	PoolId        [32]byte
	Tokens        []common.Address
	AssetManagers []common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokensRegistered is a free log retrieval operation binding the contract event 0xf5847d3f2197b16cdcd2098ec95d0905cd1abdaf415f07bb7cef2bba8ac5dec4.
//
// Solidity: event TokensRegistered(bytes32 indexed poolId, address[] tokens, address[] assetManagers)
func (_Vault *VaultFilterer) FilterTokensRegistered(opts *bind.FilterOpts, poolId [][32]byte) (*VaultTokensRegisteredIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _Vault.contract.FilterLogs(opts, "TokensRegistered", poolIdRule)
	if err != nil {
		return nil, err
	}
	return &VaultTokensRegisteredIterator{contract: _Vault.contract, event: "TokensRegistered", logs: logs, sub: sub}, nil
}

// WatchTokensRegistered is a free log subscription operation binding the contract event 0xf5847d3f2197b16cdcd2098ec95d0905cd1abdaf415f07bb7cef2bba8ac5dec4.
//
// Solidity: event TokensRegistered(bytes32 indexed poolId, address[] tokens, address[] assetManagers)
func (_Vault *VaultFilterer) WatchTokensRegistered(opts *bind.WatchOpts, sink chan<- *VaultTokensRegistered, poolId [][32]byte) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _Vault.contract.WatchLogs(opts, "TokensRegistered", poolIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultTokensRegistered)
				if err := _Vault.contract.UnpackLog(event, "TokensRegistered", log); err != nil {
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

// ParseTokensRegistered is a log parse operation binding the contract event 0xf5847d3f2197b16cdcd2098ec95d0905cd1abdaf415f07bb7cef2bba8ac5dec4.
//
// Solidity: event TokensRegistered(bytes32 indexed poolId, address[] tokens, address[] assetManagers)
func (_Vault *VaultFilterer) ParseTokensRegistered(log types.Log) (*VaultTokensRegistered, error) {
	event := new(VaultTokensRegistered)
	if err := _Vault.contract.UnpackLog(event, "TokensRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
