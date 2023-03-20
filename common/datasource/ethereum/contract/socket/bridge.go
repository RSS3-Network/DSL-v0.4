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

// BridgeMetaData contains all meta data concerning the Bridge contract.
var BridgeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"DelayPeriodUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"DelayThresholdUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"DelayedTransferAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DelayedTransferExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"EpochLengthUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cap\",\"type\":\"uint256\"}],\"name\":\"EpochVolumeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"GovernorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"GovernorRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"seqnum\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LiquidityAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MaxSendUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MinAddUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MinSendUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"transferId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"srcTransferId\",\"type\":\"bytes32\"}],\"name\":\"Relay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"resetTime\",\"type\":\"uint256\"}],\"name\":\"ResetNotification\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"transferId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"maxSlippage\",\"type\":\"uint32\"}],\"name\":\"Send\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"SignersUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"withdrawId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"seqnum\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"refid\",\"type\":\"bytes32\"}],\"name\":\"WithdrawDone\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"addGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"addNativeLiquidity\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addPauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addseq\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delayPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"delayThresholds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"delayedTransfers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"epochVolumeCaps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"epochVolumes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"executeDelayedTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"governors\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"increaseNoticePeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"isGovernor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isPauser\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lastOpTimestamps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxSend\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"minAdd\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"minSend\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimalMaxSlippage\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nativeWrap\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"noticePeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"notifyResetSigners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pausers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_relayRequest\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"relay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"removeGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removePauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renouncePauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"resetSigners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resetTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"_maxSlippage\",\"type\":\"uint32\"}],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"_maxSlippage\",\"type\":\"uint32\"}],\"name\":\"sendNative\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_period\",\"type\":\"uint256\"}],\"name\":\"setDelayPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_thresholds\",\"type\":\"uint256[]\"}],\"name\":\"setDelayThresholds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_length\",\"type\":\"uint256\"}],\"name\":\"setEpochLength\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_caps\",\"type\":\"uint256[]\"}],\"name\":\"setEpochVolumeCaps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"setMaxSend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"setMinAdd\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"setMinSend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_minimalMaxSlippage\",\"type\":\"uint32\"}],\"name\":\"setMinimalMaxSlippage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_weth\",\"type\":\"address\"}],\"name\":\"setWrap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ssHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"transfers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"triggerTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_triggerTime\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_newSigners\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_newPowers\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"_curSigners\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_curPowers\",\"type\":\"uint256[]\"}],\"name\":\"updateSigners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_msg\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"verifySigs\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_wdmsg\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"withdraws\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// BridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeMetaData.ABI instead.
var BridgeABI = BridgeMetaData.ABI

// Bridge is an auto generated Go binding around an Ethereum contract.
type Bridge struct {
	BridgeCaller     // Read-only binding to the contract
	BridgeTransactor // Write-only binding to the contract
	BridgeFilterer   // Log filterer for contract events
}

// BridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeSession struct {
	Contract     *Bridge           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeCallerSession struct {
	Contract *BridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeTransactorSession struct {
	Contract     *BridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeRaw struct {
	Contract *Bridge // Generic contract binding to access the raw methods on
}

// BridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeCallerRaw struct {
	Contract *BridgeCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeTransactorRaw struct {
	Contract *BridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridge creates a new instance of Bridge, bound to a specific deployed contract.
func NewBridge(address common.Address, backend bind.ContractBackend) (*Bridge, error) {
	contract, err := bindBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bridge{BridgeCaller: BridgeCaller{contract: contract}, BridgeTransactor: BridgeTransactor{contract: contract}, BridgeFilterer: BridgeFilterer{contract: contract}}, nil
}

// NewBridgeCaller creates a new read-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeCaller(address common.Address, caller bind.ContractCaller) (*BridgeCaller, error) {
	contract, err := bindBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeCaller{contract: contract}, nil
}

// NewBridgeTransactor creates a new write-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeTransactor, error) {
	contract, err := bindBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeTransactor{contract: contract}, nil
}

// NewBridgeFilterer creates a new log filterer instance of Bridge, bound to a specific deployed contract.
func NewBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeFilterer, error) {
	contract, err := bindBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeFilterer{contract: contract}, nil
}

// bindBridge binds a generic wrapper to an already deployed contract.
func bindBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.BridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transact(opts, method, params...)
}

// Addseq is a free data retrieval call binding the contract method 0x89e39127.
//
// Solidity: function addseq() view returns(uint64)
func (_Bridge *BridgeCaller) Addseq(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "addseq")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Addseq is a free data retrieval call binding the contract method 0x89e39127.
//
// Solidity: function addseq() view returns(uint64)
func (_Bridge *BridgeSession) Addseq() (uint64, error) {
	return _Bridge.Contract.Addseq(&_Bridge.CallOpts)
}

// Addseq is a free data retrieval call binding the contract method 0x89e39127.
//
// Solidity: function addseq() view returns(uint64)
func (_Bridge *BridgeCallerSession) Addseq() (uint64, error) {
	return _Bridge.Contract.Addseq(&_Bridge.CallOpts)
}

// DelayPeriod is a free data retrieval call binding the contract method 0xb1c94d94.
//
// Solidity: function delayPeriod() view returns(uint256)
func (_Bridge *BridgeCaller) DelayPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "delayPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelayPeriod is a free data retrieval call binding the contract method 0xb1c94d94.
//
// Solidity: function delayPeriod() view returns(uint256)
func (_Bridge *BridgeSession) DelayPeriod() (*big.Int, error) {
	return _Bridge.Contract.DelayPeriod(&_Bridge.CallOpts)
}

// DelayPeriod is a free data retrieval call binding the contract method 0xb1c94d94.
//
// Solidity: function delayPeriod() view returns(uint256)
func (_Bridge *BridgeCallerSession) DelayPeriod() (*big.Int, error) {
	return _Bridge.Contract.DelayPeriod(&_Bridge.CallOpts)
}

// DelayThresholds is a free data retrieval call binding the contract method 0x52532faa.
//
// Solidity: function delayThresholds(address ) view returns(uint256)
func (_Bridge *BridgeCaller) DelayThresholds(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "delayThresholds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelayThresholds is a free data retrieval call binding the contract method 0x52532faa.
//
// Solidity: function delayThresholds(address ) view returns(uint256)
func (_Bridge *BridgeSession) DelayThresholds(arg0 common.Address) (*big.Int, error) {
	return _Bridge.Contract.DelayThresholds(&_Bridge.CallOpts, arg0)
}

// DelayThresholds is a free data retrieval call binding the contract method 0x52532faa.
//
// Solidity: function delayThresholds(address ) view returns(uint256)
func (_Bridge *BridgeCallerSession) DelayThresholds(arg0 common.Address) (*big.Int, error) {
	return _Bridge.Contract.DelayThresholds(&_Bridge.CallOpts, arg0)
}

// DelayedTransfers is a free data retrieval call binding the contract method 0xadc0d57f.
//
// Solidity: function delayedTransfers(bytes32 ) view returns(address receiver, address token, uint256 amount, uint256 timestamp)
func (_Bridge *BridgeCaller) DelayedTransfers(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Receiver  common.Address
	Token     common.Address
	Amount    *big.Int
	Timestamp *big.Int
}, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "delayedTransfers", arg0)

	outstruct := new(struct {
		Receiver  common.Address
		Token     common.Address
		Amount    *big.Int
		Timestamp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Receiver = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Token = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// DelayedTransfers is a free data retrieval call binding the contract method 0xadc0d57f.
//
// Solidity: function delayedTransfers(bytes32 ) view returns(address receiver, address token, uint256 amount, uint256 timestamp)
func (_Bridge *BridgeSession) DelayedTransfers(arg0 [32]byte) (struct {
	Receiver  common.Address
	Token     common.Address
	Amount    *big.Int
	Timestamp *big.Int
}, error) {
	return _Bridge.Contract.DelayedTransfers(&_Bridge.CallOpts, arg0)
}

// DelayedTransfers is a free data retrieval call binding the contract method 0xadc0d57f.
//
// Solidity: function delayedTransfers(bytes32 ) view returns(address receiver, address token, uint256 amount, uint256 timestamp)
func (_Bridge *BridgeCallerSession) DelayedTransfers(arg0 [32]byte) (struct {
	Receiver  common.Address
	Token     common.Address
	Amount    *big.Int
	Timestamp *big.Int
}, error) {
	return _Bridge.Contract.DelayedTransfers(&_Bridge.CallOpts, arg0)
}

// EpochLength is a free data retrieval call binding the contract method 0x57d775f8.
//
// Solidity: function epochLength() view returns(uint256)
func (_Bridge *BridgeCaller) EpochLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "epochLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochLength is a free data retrieval call binding the contract method 0x57d775f8.
//
// Solidity: function epochLength() view returns(uint256)
func (_Bridge *BridgeSession) EpochLength() (*big.Int, error) {
	return _Bridge.Contract.EpochLength(&_Bridge.CallOpts)
}

// EpochLength is a free data retrieval call binding the contract method 0x57d775f8.
//
// Solidity: function epochLength() view returns(uint256)
func (_Bridge *BridgeCallerSession) EpochLength() (*big.Int, error) {
	return _Bridge.Contract.EpochLength(&_Bridge.CallOpts)
}

// EpochVolumeCaps is a free data retrieval call binding the contract method 0xb5f2bc47.
//
// Solidity: function epochVolumeCaps(address ) view returns(uint256)
func (_Bridge *BridgeCaller) EpochVolumeCaps(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "epochVolumeCaps", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochVolumeCaps is a free data retrieval call binding the contract method 0xb5f2bc47.
//
// Solidity: function epochVolumeCaps(address ) view returns(uint256)
func (_Bridge *BridgeSession) EpochVolumeCaps(arg0 common.Address) (*big.Int, error) {
	return _Bridge.Contract.EpochVolumeCaps(&_Bridge.CallOpts, arg0)
}

// EpochVolumeCaps is a free data retrieval call binding the contract method 0xb5f2bc47.
//
// Solidity: function epochVolumeCaps(address ) view returns(uint256)
func (_Bridge *BridgeCallerSession) EpochVolumeCaps(arg0 common.Address) (*big.Int, error) {
	return _Bridge.Contract.EpochVolumeCaps(&_Bridge.CallOpts, arg0)
}

// EpochVolumes is a free data retrieval call binding the contract method 0x60216b00.
//
// Solidity: function epochVolumes(address ) view returns(uint256)
func (_Bridge *BridgeCaller) EpochVolumes(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "epochVolumes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochVolumes is a free data retrieval call binding the contract method 0x60216b00.
//
// Solidity: function epochVolumes(address ) view returns(uint256)
func (_Bridge *BridgeSession) EpochVolumes(arg0 common.Address) (*big.Int, error) {
	return _Bridge.Contract.EpochVolumes(&_Bridge.CallOpts, arg0)
}

// EpochVolumes is a free data retrieval call binding the contract method 0x60216b00.
//
// Solidity: function epochVolumes(address ) view returns(uint256)
func (_Bridge *BridgeCallerSession) EpochVolumes(arg0 common.Address) (*big.Int, error) {
	return _Bridge.Contract.EpochVolumes(&_Bridge.CallOpts, arg0)
}

// Governors is a free data retrieval call binding the contract method 0xe3eece26.
//
// Solidity: function governors(address ) view returns(bool)
func (_Bridge *BridgeCaller) Governors(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "governors", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Governors is a free data retrieval call binding the contract method 0xe3eece26.
//
// Solidity: function governors(address ) view returns(bool)
func (_Bridge *BridgeSession) Governors(arg0 common.Address) (bool, error) {
	return _Bridge.Contract.Governors(&_Bridge.CallOpts, arg0)
}

// Governors is a free data retrieval call binding the contract method 0xe3eece26.
//
// Solidity: function governors(address ) view returns(bool)
func (_Bridge *BridgeCallerSession) Governors(arg0 common.Address) (bool, error) {
	return _Bridge.Contract.Governors(&_Bridge.CallOpts, arg0)
}

// IsGovernor is a free data retrieval call binding the contract method 0xe43581b8.
//
// Solidity: function isGovernor(address _account) view returns(bool)
func (_Bridge *BridgeCaller) IsGovernor(opts *bind.CallOpts, _account common.Address) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "isGovernor", _account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsGovernor is a free data retrieval call binding the contract method 0xe43581b8.
//
// Solidity: function isGovernor(address _account) view returns(bool)
func (_Bridge *BridgeSession) IsGovernor(_account common.Address) (bool, error) {
	return _Bridge.Contract.IsGovernor(&_Bridge.CallOpts, _account)
}

// IsGovernor is a free data retrieval call binding the contract method 0xe43581b8.
//
// Solidity: function isGovernor(address _account) view returns(bool)
func (_Bridge *BridgeCallerSession) IsGovernor(_account common.Address) (bool, error) {
	return _Bridge.Contract.IsGovernor(&_Bridge.CallOpts, _account)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) view returns(bool)
func (_Bridge *BridgeCaller) IsPauser(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "isPauser", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) view returns(bool)
func (_Bridge *BridgeSession) IsPauser(account common.Address) (bool, error) {
	return _Bridge.Contract.IsPauser(&_Bridge.CallOpts, account)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) view returns(bool)
func (_Bridge *BridgeCallerSession) IsPauser(account common.Address) (bool, error) {
	return _Bridge.Contract.IsPauser(&_Bridge.CallOpts, account)
}

// LastOpTimestamps is a free data retrieval call binding the contract method 0xf8321383.
//
// Solidity: function lastOpTimestamps(address ) view returns(uint256)
func (_Bridge *BridgeCaller) LastOpTimestamps(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "lastOpTimestamps", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastOpTimestamps is a free data retrieval call binding the contract method 0xf8321383.
//
// Solidity: function lastOpTimestamps(address ) view returns(uint256)
func (_Bridge *BridgeSession) LastOpTimestamps(arg0 common.Address) (*big.Int, error) {
	return _Bridge.Contract.LastOpTimestamps(&_Bridge.CallOpts, arg0)
}

// LastOpTimestamps is a free data retrieval call binding the contract method 0xf8321383.
//
// Solidity: function lastOpTimestamps(address ) view returns(uint256)
func (_Bridge *BridgeCallerSession) LastOpTimestamps(arg0 common.Address) (*big.Int, error) {
	return _Bridge.Contract.LastOpTimestamps(&_Bridge.CallOpts, arg0)
}

// MaxSend is a free data retrieval call binding the contract method 0x618ee055.
//
// Solidity: function maxSend(address ) view returns(uint256)
func (_Bridge *BridgeCaller) MaxSend(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "maxSend", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxSend is a free data retrieval call binding the contract method 0x618ee055.
//
// Solidity: function maxSend(address ) view returns(uint256)
func (_Bridge *BridgeSession) MaxSend(arg0 common.Address) (*big.Int, error) {
	return _Bridge.Contract.MaxSend(&_Bridge.CallOpts, arg0)
}

// MaxSend is a free data retrieval call binding the contract method 0x618ee055.
//
// Solidity: function maxSend(address ) view returns(uint256)
func (_Bridge *BridgeCallerSession) MaxSend(arg0 common.Address) (*big.Int, error) {
	return _Bridge.Contract.MaxSend(&_Bridge.CallOpts, arg0)
}

// MinAdd is a free data retrieval call binding the contract method 0xccde517a.
//
// Solidity: function minAdd(address ) view returns(uint256)
func (_Bridge *BridgeCaller) MinAdd(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "minAdd", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinAdd is a free data retrieval call binding the contract method 0xccde517a.
//
// Solidity: function minAdd(address ) view returns(uint256)
func (_Bridge *BridgeSession) MinAdd(arg0 common.Address) (*big.Int, error) {
	return _Bridge.Contract.MinAdd(&_Bridge.CallOpts, arg0)
}

// MinAdd is a free data retrieval call binding the contract method 0xccde517a.
//
// Solidity: function minAdd(address ) view returns(uint256)
func (_Bridge *BridgeCallerSession) MinAdd(arg0 common.Address) (*big.Int, error) {
	return _Bridge.Contract.MinAdd(&_Bridge.CallOpts, arg0)
}

// MinSend is a free data retrieval call binding the contract method 0xf8b30d7d.
//
// Solidity: function minSend(address ) view returns(uint256)
func (_Bridge *BridgeCaller) MinSend(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "minSend", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinSend is a free data retrieval call binding the contract method 0xf8b30d7d.
//
// Solidity: function minSend(address ) view returns(uint256)
func (_Bridge *BridgeSession) MinSend(arg0 common.Address) (*big.Int, error) {
	return _Bridge.Contract.MinSend(&_Bridge.CallOpts, arg0)
}

// MinSend is a free data retrieval call binding the contract method 0xf8b30d7d.
//
// Solidity: function minSend(address ) view returns(uint256)
func (_Bridge *BridgeCallerSession) MinSend(arg0 common.Address) (*big.Int, error) {
	return _Bridge.Contract.MinSend(&_Bridge.CallOpts, arg0)
}

// MinimalMaxSlippage is a free data retrieval call binding the contract method 0x2fd1b0a4.
//
// Solidity: function minimalMaxSlippage() view returns(uint32)
func (_Bridge *BridgeCaller) MinimalMaxSlippage(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "minimalMaxSlippage")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MinimalMaxSlippage is a free data retrieval call binding the contract method 0x2fd1b0a4.
//
// Solidity: function minimalMaxSlippage() view returns(uint32)
func (_Bridge *BridgeSession) MinimalMaxSlippage() (uint32, error) {
	return _Bridge.Contract.MinimalMaxSlippage(&_Bridge.CallOpts)
}

// MinimalMaxSlippage is a free data retrieval call binding the contract method 0x2fd1b0a4.
//
// Solidity: function minimalMaxSlippage() view returns(uint32)
func (_Bridge *BridgeCallerSession) MinimalMaxSlippage() (uint32, error) {
	return _Bridge.Contract.MinimalMaxSlippage(&_Bridge.CallOpts)
}

// NativeWrap is a free data retrieval call binding the contract method 0x457bfa2f.
//
// Solidity: function nativeWrap() view returns(address)
func (_Bridge *BridgeCaller) NativeWrap(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "nativeWrap")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NativeWrap is a free data retrieval call binding the contract method 0x457bfa2f.
//
// Solidity: function nativeWrap() view returns(address)
func (_Bridge *BridgeSession) NativeWrap() (common.Address, error) {
	return _Bridge.Contract.NativeWrap(&_Bridge.CallOpts)
}

// NativeWrap is a free data retrieval call binding the contract method 0x457bfa2f.
//
// Solidity: function nativeWrap() view returns(address)
func (_Bridge *BridgeCallerSession) NativeWrap() (common.Address, error) {
	return _Bridge.Contract.NativeWrap(&_Bridge.CallOpts)
}

// NoticePeriod is a free data retrieval call binding the contract method 0x9b14d4c6.
//
// Solidity: function noticePeriod() view returns(uint256)
func (_Bridge *BridgeCaller) NoticePeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "noticePeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NoticePeriod is a free data retrieval call binding the contract method 0x9b14d4c6.
//
// Solidity: function noticePeriod() view returns(uint256)
func (_Bridge *BridgeSession) NoticePeriod() (*big.Int, error) {
	return _Bridge.Contract.NoticePeriod(&_Bridge.CallOpts)
}

// NoticePeriod is a free data retrieval call binding the contract method 0x9b14d4c6.
//
// Solidity: function noticePeriod() view returns(uint256)
func (_Bridge *BridgeCallerSession) NoticePeriod() (*big.Int, error) {
	return _Bridge.Contract.NoticePeriod(&_Bridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bridge *BridgeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bridge *BridgeSession) Owner() (common.Address, error) {
	return _Bridge.Contract.Owner(&_Bridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bridge *BridgeCallerSession) Owner() (common.Address, error) {
	return _Bridge.Contract.Owner(&_Bridge.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Bridge *BridgeCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Bridge *BridgeSession) Paused() (bool, error) {
	return _Bridge.Contract.Paused(&_Bridge.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Bridge *BridgeCallerSession) Paused() (bool, error) {
	return _Bridge.Contract.Paused(&_Bridge.CallOpts)
}

// Pausers is a free data retrieval call binding the contract method 0x80f51c12.
//
// Solidity: function pausers(address ) view returns(bool)
func (_Bridge *BridgeCaller) Pausers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "pausers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Pausers is a free data retrieval call binding the contract method 0x80f51c12.
//
// Solidity: function pausers(address ) view returns(bool)
func (_Bridge *BridgeSession) Pausers(arg0 common.Address) (bool, error) {
	return _Bridge.Contract.Pausers(&_Bridge.CallOpts, arg0)
}

// Pausers is a free data retrieval call binding the contract method 0x80f51c12.
//
// Solidity: function pausers(address ) view returns(bool)
func (_Bridge *BridgeCallerSession) Pausers(arg0 common.Address) (bool, error) {
	return _Bridge.Contract.Pausers(&_Bridge.CallOpts, arg0)
}

// ResetTime is a free data retrieval call binding the contract method 0x65a114f1.
//
// Solidity: function resetTime() view returns(uint256)
func (_Bridge *BridgeCaller) ResetTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "resetTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ResetTime is a free data retrieval call binding the contract method 0x65a114f1.
//
// Solidity: function resetTime() view returns(uint256)
func (_Bridge *BridgeSession) ResetTime() (*big.Int, error) {
	return _Bridge.Contract.ResetTime(&_Bridge.CallOpts)
}

// ResetTime is a free data retrieval call binding the contract method 0x65a114f1.
//
// Solidity: function resetTime() view returns(uint256)
func (_Bridge *BridgeCallerSession) ResetTime() (*big.Int, error) {
	return _Bridge.Contract.ResetTime(&_Bridge.CallOpts)
}

// SsHash is a free data retrieval call binding the contract method 0xd0790da9.
//
// Solidity: function ssHash() view returns(bytes32)
func (_Bridge *BridgeCaller) SsHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "ssHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SsHash is a free data retrieval call binding the contract method 0xd0790da9.
//
// Solidity: function ssHash() view returns(bytes32)
func (_Bridge *BridgeSession) SsHash() ([32]byte, error) {
	return _Bridge.Contract.SsHash(&_Bridge.CallOpts)
}

// SsHash is a free data retrieval call binding the contract method 0xd0790da9.
//
// Solidity: function ssHash() view returns(bytes32)
func (_Bridge *BridgeCallerSession) SsHash() ([32]byte, error) {
	return _Bridge.Contract.SsHash(&_Bridge.CallOpts)
}

// Transfers is a free data retrieval call binding the contract method 0x3c64f04b.
//
// Solidity: function transfers(bytes32 ) view returns(bool)
func (_Bridge *BridgeCaller) Transfers(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "transfers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Transfers is a free data retrieval call binding the contract method 0x3c64f04b.
//
// Solidity: function transfers(bytes32 ) view returns(bool)
func (_Bridge *BridgeSession) Transfers(arg0 [32]byte) (bool, error) {
	return _Bridge.Contract.Transfers(&_Bridge.CallOpts, arg0)
}

// Transfers is a free data retrieval call binding the contract method 0x3c64f04b.
//
// Solidity: function transfers(bytes32 ) view returns(bool)
func (_Bridge *BridgeCallerSession) Transfers(arg0 [32]byte) (bool, error) {
	return _Bridge.Contract.Transfers(&_Bridge.CallOpts, arg0)
}

// TriggerTime is a free data retrieval call binding the contract method 0x370fb47b.
//
// Solidity: function triggerTime() view returns(uint256)
func (_Bridge *BridgeCaller) TriggerTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "triggerTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TriggerTime is a free data retrieval call binding the contract method 0x370fb47b.
//
// Solidity: function triggerTime() view returns(uint256)
func (_Bridge *BridgeSession) TriggerTime() (*big.Int, error) {
	return _Bridge.Contract.TriggerTime(&_Bridge.CallOpts)
}

// TriggerTime is a free data retrieval call binding the contract method 0x370fb47b.
//
// Solidity: function triggerTime() view returns(uint256)
func (_Bridge *BridgeCallerSession) TriggerTime() (*big.Int, error) {
	return _Bridge.Contract.TriggerTime(&_Bridge.CallOpts)
}

// VerifySigs is a free data retrieval call binding the contract method 0x682dbc22.
//
// Solidity: function verifySigs(bytes _msg, bytes[] _sigs, address[] _signers, uint256[] _powers) view returns()
func (_Bridge *BridgeCaller) VerifySigs(opts *bind.CallOpts, _msg []byte, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) error {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "verifySigs", _msg, _sigs, _signers, _powers)

	if err != nil {
		return err
	}

	return err

}

// VerifySigs is a free data retrieval call binding the contract method 0x682dbc22.
//
// Solidity: function verifySigs(bytes _msg, bytes[] _sigs, address[] _signers, uint256[] _powers) view returns()
func (_Bridge *BridgeSession) VerifySigs(_msg []byte, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) error {
	return _Bridge.Contract.VerifySigs(&_Bridge.CallOpts, _msg, _sigs, _signers, _powers)
}

// VerifySigs is a free data retrieval call binding the contract method 0x682dbc22.
//
// Solidity: function verifySigs(bytes _msg, bytes[] _sigs, address[] _signers, uint256[] _powers) view returns()
func (_Bridge *BridgeCallerSession) VerifySigs(_msg []byte, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) error {
	return _Bridge.Contract.VerifySigs(&_Bridge.CallOpts, _msg, _sigs, _signers, _powers)
}

// Withdraws is a free data retrieval call binding the contract method 0xe09ab428.
//
// Solidity: function withdraws(bytes32 ) view returns(bool)
func (_Bridge *BridgeCaller) Withdraws(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "withdraws", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Withdraws is a free data retrieval call binding the contract method 0xe09ab428.
//
// Solidity: function withdraws(bytes32 ) view returns(bool)
func (_Bridge *BridgeSession) Withdraws(arg0 [32]byte) (bool, error) {
	return _Bridge.Contract.Withdraws(&_Bridge.CallOpts, arg0)
}

// Withdraws is a free data retrieval call binding the contract method 0xe09ab428.
//
// Solidity: function withdraws(bytes32 ) view returns(bool)
func (_Bridge *BridgeCallerSession) Withdraws(arg0 [32]byte) (bool, error) {
	return _Bridge.Contract.Withdraws(&_Bridge.CallOpts, arg0)
}

// AddGovernor is a paid mutator transaction binding the contract method 0x3c4a25d0.
//
// Solidity: function addGovernor(address _account) returns()
func (_Bridge *BridgeTransactor) AddGovernor(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "addGovernor", _account)
}

// AddGovernor is a paid mutator transaction binding the contract method 0x3c4a25d0.
//
// Solidity: function addGovernor(address _account) returns()
func (_Bridge *BridgeSession) AddGovernor(_account common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.AddGovernor(&_Bridge.TransactOpts, _account)
}

// AddGovernor is a paid mutator transaction binding the contract method 0x3c4a25d0.
//
// Solidity: function addGovernor(address _account) returns()
func (_Bridge *BridgeTransactorSession) AddGovernor(_account common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.AddGovernor(&_Bridge.TransactOpts, _account)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x56688700.
//
// Solidity: function addLiquidity(address _token, uint256 _amount) returns()
func (_Bridge *BridgeTransactor) AddLiquidity(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "addLiquidity", _token, _amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x56688700.
//
// Solidity: function addLiquidity(address _token, uint256 _amount) returns()
func (_Bridge *BridgeSession) AddLiquidity(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.AddLiquidity(&_Bridge.TransactOpts, _token, _amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x56688700.
//
// Solidity: function addLiquidity(address _token, uint256 _amount) returns()
func (_Bridge *BridgeTransactorSession) AddLiquidity(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.AddLiquidity(&_Bridge.TransactOpts, _token, _amount)
}

// AddNativeLiquidity is a paid mutator transaction binding the contract method 0x7044c89e.
//
// Solidity: function addNativeLiquidity(uint256 _amount) payable returns()
func (_Bridge *BridgeTransactor) AddNativeLiquidity(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "addNativeLiquidity", _amount)
}

// AddNativeLiquidity is a paid mutator transaction binding the contract method 0x7044c89e.
//
// Solidity: function addNativeLiquidity(uint256 _amount) payable returns()
func (_Bridge *BridgeSession) AddNativeLiquidity(_amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.AddNativeLiquidity(&_Bridge.TransactOpts, _amount)
}

// AddNativeLiquidity is a paid mutator transaction binding the contract method 0x7044c89e.
//
// Solidity: function addNativeLiquidity(uint256 _amount) payable returns()
func (_Bridge *BridgeTransactorSession) AddNativeLiquidity(_amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.AddNativeLiquidity(&_Bridge.TransactOpts, _amount)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_Bridge *BridgeTransactor) AddPauser(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "addPauser", account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_Bridge *BridgeSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.AddPauser(&_Bridge.TransactOpts, account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_Bridge *BridgeTransactorSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.AddPauser(&_Bridge.TransactOpts, account)
}

// ExecuteDelayedTransfer is a paid mutator transaction binding the contract method 0x9e25fc5c.
//
// Solidity: function executeDelayedTransfer(bytes32 id) returns()
func (_Bridge *BridgeTransactor) ExecuteDelayedTransfer(opts *bind.TransactOpts, id [32]byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "executeDelayedTransfer", id)
}

// ExecuteDelayedTransfer is a paid mutator transaction binding the contract method 0x9e25fc5c.
//
// Solidity: function executeDelayedTransfer(bytes32 id) returns()
func (_Bridge *BridgeSession) ExecuteDelayedTransfer(id [32]byte) (*types.Transaction, error) {
	return _Bridge.Contract.ExecuteDelayedTransfer(&_Bridge.TransactOpts, id)
}

// ExecuteDelayedTransfer is a paid mutator transaction binding the contract method 0x9e25fc5c.
//
// Solidity: function executeDelayedTransfer(bytes32 id) returns()
func (_Bridge *BridgeTransactorSession) ExecuteDelayedTransfer(id [32]byte) (*types.Transaction, error) {
	return _Bridge.Contract.ExecuteDelayedTransfer(&_Bridge.TransactOpts, id)
}

// IncreaseNoticePeriod is a paid mutator transaction binding the contract method 0xf20c922a.
//
// Solidity: function increaseNoticePeriod(uint256 period) returns()
func (_Bridge *BridgeTransactor) IncreaseNoticePeriod(opts *bind.TransactOpts, period *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "increaseNoticePeriod", period)
}

// IncreaseNoticePeriod is a paid mutator transaction binding the contract method 0xf20c922a.
//
// Solidity: function increaseNoticePeriod(uint256 period) returns()
func (_Bridge *BridgeSession) IncreaseNoticePeriod(period *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.IncreaseNoticePeriod(&_Bridge.TransactOpts, period)
}

// IncreaseNoticePeriod is a paid mutator transaction binding the contract method 0xf20c922a.
//
// Solidity: function increaseNoticePeriod(uint256 period) returns()
func (_Bridge *BridgeTransactorSession) IncreaseNoticePeriod(period *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.IncreaseNoticePeriod(&_Bridge.TransactOpts, period)
}

// NotifyResetSigners is a paid mutator transaction binding the contract method 0x25c38b9f.
//
// Solidity: function notifyResetSigners() returns()
func (_Bridge *BridgeTransactor) NotifyResetSigners(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "notifyResetSigners")
}

// NotifyResetSigners is a paid mutator transaction binding the contract method 0x25c38b9f.
//
// Solidity: function notifyResetSigners() returns()
func (_Bridge *BridgeSession) NotifyResetSigners() (*types.Transaction, error) {
	return _Bridge.Contract.NotifyResetSigners(&_Bridge.TransactOpts)
}

// NotifyResetSigners is a paid mutator transaction binding the contract method 0x25c38b9f.
//
// Solidity: function notifyResetSigners() returns()
func (_Bridge *BridgeTransactorSession) NotifyResetSigners() (*types.Transaction, error) {
	return _Bridge.Contract.NotifyResetSigners(&_Bridge.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Bridge *BridgeTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Bridge *BridgeSession) Pause() (*types.Transaction, error) {
	return _Bridge.Contract.Pause(&_Bridge.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Bridge *BridgeTransactorSession) Pause() (*types.Transaction, error) {
	return _Bridge.Contract.Pause(&_Bridge.TransactOpts)
}

// Relay is a paid mutator transaction binding the contract method 0xcdd1b25d.
//
// Solidity: function relay(bytes _relayRequest, bytes[] _sigs, address[] _signers, uint256[] _powers) returns()
func (_Bridge *BridgeTransactor) Relay(opts *bind.TransactOpts, _relayRequest []byte, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "relay", _relayRequest, _sigs, _signers, _powers)
}

// Relay is a paid mutator transaction binding the contract method 0xcdd1b25d.
//
// Solidity: function relay(bytes _relayRequest, bytes[] _sigs, address[] _signers, uint256[] _powers) returns()
func (_Bridge *BridgeSession) Relay(_relayRequest []byte, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.Relay(&_Bridge.TransactOpts, _relayRequest, _sigs, _signers, _powers)
}

// Relay is a paid mutator transaction binding the contract method 0xcdd1b25d.
//
// Solidity: function relay(bytes _relayRequest, bytes[] _sigs, address[] _signers, uint256[] _powers) returns()
func (_Bridge *BridgeTransactorSession) Relay(_relayRequest []byte, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.Relay(&_Bridge.TransactOpts, _relayRequest, _sigs, _signers, _powers)
}

// RemoveGovernor is a paid mutator transaction binding the contract method 0xeecdac88.
//
// Solidity: function removeGovernor(address _account) returns()
func (_Bridge *BridgeTransactor) RemoveGovernor(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "removeGovernor", _account)
}

// RemoveGovernor is a paid mutator transaction binding the contract method 0xeecdac88.
//
// Solidity: function removeGovernor(address _account) returns()
func (_Bridge *BridgeSession) RemoveGovernor(_account common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.RemoveGovernor(&_Bridge.TransactOpts, _account)
}

// RemoveGovernor is a paid mutator transaction binding the contract method 0xeecdac88.
//
// Solidity: function removeGovernor(address _account) returns()
func (_Bridge *BridgeTransactorSession) RemoveGovernor(_account common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.RemoveGovernor(&_Bridge.TransactOpts, _account)
}

// RemovePauser is a paid mutator transaction binding the contract method 0x6b2c0f55.
//
// Solidity: function removePauser(address account) returns()
func (_Bridge *BridgeTransactor) RemovePauser(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "removePauser", account)
}

// RemovePauser is a paid mutator transaction binding the contract method 0x6b2c0f55.
//
// Solidity: function removePauser(address account) returns()
func (_Bridge *BridgeSession) RemovePauser(account common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.RemovePauser(&_Bridge.TransactOpts, account)
}

// RemovePauser is a paid mutator transaction binding the contract method 0x6b2c0f55.
//
// Solidity: function removePauser(address account) returns()
func (_Bridge *BridgeTransactorSession) RemovePauser(account common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.RemovePauser(&_Bridge.TransactOpts, account)
}

// RenounceGovernor is a paid mutator transaction binding the contract method 0xe026049c.
//
// Solidity: function renounceGovernor() returns()
func (_Bridge *BridgeTransactor) RenounceGovernor(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "renounceGovernor")
}

// RenounceGovernor is a paid mutator transaction binding the contract method 0xe026049c.
//
// Solidity: function renounceGovernor() returns()
func (_Bridge *BridgeSession) RenounceGovernor() (*types.Transaction, error) {
	return _Bridge.Contract.RenounceGovernor(&_Bridge.TransactOpts)
}

// RenounceGovernor is a paid mutator transaction binding the contract method 0xe026049c.
//
// Solidity: function renounceGovernor() returns()
func (_Bridge *BridgeTransactorSession) RenounceGovernor() (*types.Transaction, error) {
	return _Bridge.Contract.RenounceGovernor(&_Bridge.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bridge *BridgeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bridge *BridgeSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bridge.Contract.RenounceOwnership(&_Bridge.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bridge *BridgeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bridge.Contract.RenounceOwnership(&_Bridge.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_Bridge *BridgeTransactor) RenouncePauser(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "renouncePauser")
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_Bridge *BridgeSession) RenouncePauser() (*types.Transaction, error) {
	return _Bridge.Contract.RenouncePauser(&_Bridge.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_Bridge *BridgeTransactorSession) RenouncePauser() (*types.Transaction, error) {
	return _Bridge.Contract.RenouncePauser(&_Bridge.TransactOpts)
}

// ResetSigners is a paid mutator transaction binding the contract method 0xa7bdf45a.
//
// Solidity: function resetSigners(address[] _signers, uint256[] _powers) returns()
func (_Bridge *BridgeTransactor) ResetSigners(opts *bind.TransactOpts, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "resetSigners", _signers, _powers)
}

// ResetSigners is a paid mutator transaction binding the contract method 0xa7bdf45a.
//
// Solidity: function resetSigners(address[] _signers, uint256[] _powers) returns()
func (_Bridge *BridgeSession) ResetSigners(_signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.ResetSigners(&_Bridge.TransactOpts, _signers, _powers)
}

// ResetSigners is a paid mutator transaction binding the contract method 0xa7bdf45a.
//
// Solidity: function resetSigners(address[] _signers, uint256[] _powers) returns()
func (_Bridge *BridgeTransactorSession) ResetSigners(_signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.ResetSigners(&_Bridge.TransactOpts, _signers, _powers)
}

// Send is a paid mutator transaction binding the contract method 0xa5977fbb.
//
// Solidity: function send(address _receiver, address _token, uint256 _amount, uint64 _dstChainId, uint64 _nonce, uint32 _maxSlippage) returns()
func (_Bridge *BridgeTransactor) Send(opts *bind.TransactOpts, _receiver common.Address, _token common.Address, _amount *big.Int, _dstChainId uint64, _nonce uint64, _maxSlippage uint32) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "send", _receiver, _token, _amount, _dstChainId, _nonce, _maxSlippage)
}

// Send is a paid mutator transaction binding the contract method 0xa5977fbb.
//
// Solidity: function send(address _receiver, address _token, uint256 _amount, uint64 _dstChainId, uint64 _nonce, uint32 _maxSlippage) returns()
func (_Bridge *BridgeSession) Send(_receiver common.Address, _token common.Address, _amount *big.Int, _dstChainId uint64, _nonce uint64, _maxSlippage uint32) (*types.Transaction, error) {
	return _Bridge.Contract.Send(&_Bridge.TransactOpts, _receiver, _token, _amount, _dstChainId, _nonce, _maxSlippage)
}

// Send is a paid mutator transaction binding the contract method 0xa5977fbb.
//
// Solidity: function send(address _receiver, address _token, uint256 _amount, uint64 _dstChainId, uint64 _nonce, uint32 _maxSlippage) returns()
func (_Bridge *BridgeTransactorSession) Send(_receiver common.Address, _token common.Address, _amount *big.Int, _dstChainId uint64, _nonce uint64, _maxSlippage uint32) (*types.Transaction, error) {
	return _Bridge.Contract.Send(&_Bridge.TransactOpts, _receiver, _token, _amount, _dstChainId, _nonce, _maxSlippage)
}

// SendNative is a paid mutator transaction binding the contract method 0x3f2e5fc3.
//
// Solidity: function sendNative(address _receiver, uint256 _amount, uint64 _dstChainId, uint64 _nonce, uint32 _maxSlippage) payable returns()
func (_Bridge *BridgeTransactor) SendNative(opts *bind.TransactOpts, _receiver common.Address, _amount *big.Int, _dstChainId uint64, _nonce uint64, _maxSlippage uint32) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "sendNative", _receiver, _amount, _dstChainId, _nonce, _maxSlippage)
}

// SendNative is a paid mutator transaction binding the contract method 0x3f2e5fc3.
//
// Solidity: function sendNative(address _receiver, uint256 _amount, uint64 _dstChainId, uint64 _nonce, uint32 _maxSlippage) payable returns()
func (_Bridge *BridgeSession) SendNative(_receiver common.Address, _amount *big.Int, _dstChainId uint64, _nonce uint64, _maxSlippage uint32) (*types.Transaction, error) {
	return _Bridge.Contract.SendNative(&_Bridge.TransactOpts, _receiver, _amount, _dstChainId, _nonce, _maxSlippage)
}

// SendNative is a paid mutator transaction binding the contract method 0x3f2e5fc3.
//
// Solidity: function sendNative(address _receiver, uint256 _amount, uint64 _dstChainId, uint64 _nonce, uint32 _maxSlippage) payable returns()
func (_Bridge *BridgeTransactorSession) SendNative(_receiver common.Address, _amount *big.Int, _dstChainId uint64, _nonce uint64, _maxSlippage uint32) (*types.Transaction, error) {
	return _Bridge.Contract.SendNative(&_Bridge.TransactOpts, _receiver, _amount, _dstChainId, _nonce, _maxSlippage)
}

// SetDelayPeriod is a paid mutator transaction binding the contract method 0x3d572107.
//
// Solidity: function setDelayPeriod(uint256 _period) returns()
func (_Bridge *BridgeTransactor) SetDelayPeriod(opts *bind.TransactOpts, _period *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setDelayPeriod", _period)
}

// SetDelayPeriod is a paid mutator transaction binding the contract method 0x3d572107.
//
// Solidity: function setDelayPeriod(uint256 _period) returns()
func (_Bridge *BridgeSession) SetDelayPeriod(_period *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.SetDelayPeriod(&_Bridge.TransactOpts, _period)
}

// SetDelayPeriod is a paid mutator transaction binding the contract method 0x3d572107.
//
// Solidity: function setDelayPeriod(uint256 _period) returns()
func (_Bridge *BridgeTransactorSession) SetDelayPeriod(_period *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.SetDelayPeriod(&_Bridge.TransactOpts, _period)
}

// SetDelayThresholds is a paid mutator transaction binding the contract method 0x17bdbae5.
//
// Solidity: function setDelayThresholds(address[] _tokens, uint256[] _thresholds) returns()
func (_Bridge *BridgeTransactor) SetDelayThresholds(opts *bind.TransactOpts, _tokens []common.Address, _thresholds []*big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setDelayThresholds", _tokens, _thresholds)
}

// SetDelayThresholds is a paid mutator transaction binding the contract method 0x17bdbae5.
//
// Solidity: function setDelayThresholds(address[] _tokens, uint256[] _thresholds) returns()
func (_Bridge *BridgeSession) SetDelayThresholds(_tokens []common.Address, _thresholds []*big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.SetDelayThresholds(&_Bridge.TransactOpts, _tokens, _thresholds)
}

// SetDelayThresholds is a paid mutator transaction binding the contract method 0x17bdbae5.
//
// Solidity: function setDelayThresholds(address[] _tokens, uint256[] _thresholds) returns()
func (_Bridge *BridgeTransactorSession) SetDelayThresholds(_tokens []common.Address, _thresholds []*big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.SetDelayThresholds(&_Bridge.TransactOpts, _tokens, _thresholds)
}

// SetEpochLength is a paid mutator transaction binding the contract method 0x54eea796.
//
// Solidity: function setEpochLength(uint256 _length) returns()
func (_Bridge *BridgeTransactor) SetEpochLength(opts *bind.TransactOpts, _length *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setEpochLength", _length)
}

// SetEpochLength is a paid mutator transaction binding the contract method 0x54eea796.
//
// Solidity: function setEpochLength(uint256 _length) returns()
func (_Bridge *BridgeSession) SetEpochLength(_length *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.SetEpochLength(&_Bridge.TransactOpts, _length)
}

// SetEpochLength is a paid mutator transaction binding the contract method 0x54eea796.
//
// Solidity: function setEpochLength(uint256 _length) returns()
func (_Bridge *BridgeTransactorSession) SetEpochLength(_length *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.SetEpochLength(&_Bridge.TransactOpts, _length)
}

// SetEpochVolumeCaps is a paid mutator transaction binding the contract method 0x47b16c6c.
//
// Solidity: function setEpochVolumeCaps(address[] _tokens, uint256[] _caps) returns()
func (_Bridge *BridgeTransactor) SetEpochVolumeCaps(opts *bind.TransactOpts, _tokens []common.Address, _caps []*big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setEpochVolumeCaps", _tokens, _caps)
}

// SetEpochVolumeCaps is a paid mutator transaction binding the contract method 0x47b16c6c.
//
// Solidity: function setEpochVolumeCaps(address[] _tokens, uint256[] _caps) returns()
func (_Bridge *BridgeSession) SetEpochVolumeCaps(_tokens []common.Address, _caps []*big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.SetEpochVolumeCaps(&_Bridge.TransactOpts, _tokens, _caps)
}

// SetEpochVolumeCaps is a paid mutator transaction binding the contract method 0x47b16c6c.
//
// Solidity: function setEpochVolumeCaps(address[] _tokens, uint256[] _caps) returns()
func (_Bridge *BridgeTransactorSession) SetEpochVolumeCaps(_tokens []common.Address, _caps []*big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.SetEpochVolumeCaps(&_Bridge.TransactOpts, _tokens, _caps)
}

// SetMaxSend is a paid mutator transaction binding the contract method 0x878fe1ce.
//
// Solidity: function setMaxSend(address[] _tokens, uint256[] _amounts) returns()
func (_Bridge *BridgeTransactor) SetMaxSend(opts *bind.TransactOpts, _tokens []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setMaxSend", _tokens, _amounts)
}

// SetMaxSend is a paid mutator transaction binding the contract method 0x878fe1ce.
//
// Solidity: function setMaxSend(address[] _tokens, uint256[] _amounts) returns()
func (_Bridge *BridgeSession) SetMaxSend(_tokens []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.SetMaxSend(&_Bridge.TransactOpts, _tokens, _amounts)
}

// SetMaxSend is a paid mutator transaction binding the contract method 0x878fe1ce.
//
// Solidity: function setMaxSend(address[] _tokens, uint256[] _amounts) returns()
func (_Bridge *BridgeTransactorSession) SetMaxSend(_tokens []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.SetMaxSend(&_Bridge.TransactOpts, _tokens, _amounts)
}

// SetMinAdd is a paid mutator transaction binding the contract method 0xe999e5f4.
//
// Solidity: function setMinAdd(address[] _tokens, uint256[] _amounts) returns()
func (_Bridge *BridgeTransactor) SetMinAdd(opts *bind.TransactOpts, _tokens []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setMinAdd", _tokens, _amounts)
}

// SetMinAdd is a paid mutator transaction binding the contract method 0xe999e5f4.
//
// Solidity: function setMinAdd(address[] _tokens, uint256[] _amounts) returns()
func (_Bridge *BridgeSession) SetMinAdd(_tokens []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.SetMinAdd(&_Bridge.TransactOpts, _tokens, _amounts)
}

// SetMinAdd is a paid mutator transaction binding the contract method 0xe999e5f4.
//
// Solidity: function setMinAdd(address[] _tokens, uint256[] _amounts) returns()
func (_Bridge *BridgeTransactorSession) SetMinAdd(_tokens []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.SetMinAdd(&_Bridge.TransactOpts, _tokens, _amounts)
}

// SetMinSend is a paid mutator transaction binding the contract method 0x08992741.
//
// Solidity: function setMinSend(address[] _tokens, uint256[] _amounts) returns()
func (_Bridge *BridgeTransactor) SetMinSend(opts *bind.TransactOpts, _tokens []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setMinSend", _tokens, _amounts)
}

// SetMinSend is a paid mutator transaction binding the contract method 0x08992741.
//
// Solidity: function setMinSend(address[] _tokens, uint256[] _amounts) returns()
func (_Bridge *BridgeSession) SetMinSend(_tokens []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.SetMinSend(&_Bridge.TransactOpts, _tokens, _amounts)
}

// SetMinSend is a paid mutator transaction binding the contract method 0x08992741.
//
// Solidity: function setMinSend(address[] _tokens, uint256[] _amounts) returns()
func (_Bridge *BridgeTransactorSession) SetMinSend(_tokens []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.SetMinSend(&_Bridge.TransactOpts, _tokens, _amounts)
}

// SetMinimalMaxSlippage is a paid mutator transaction binding the contract method 0x48234126.
//
// Solidity: function setMinimalMaxSlippage(uint32 _minimalMaxSlippage) returns()
func (_Bridge *BridgeTransactor) SetMinimalMaxSlippage(opts *bind.TransactOpts, _minimalMaxSlippage uint32) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setMinimalMaxSlippage", _minimalMaxSlippage)
}

// SetMinimalMaxSlippage is a paid mutator transaction binding the contract method 0x48234126.
//
// Solidity: function setMinimalMaxSlippage(uint32 _minimalMaxSlippage) returns()
func (_Bridge *BridgeSession) SetMinimalMaxSlippage(_minimalMaxSlippage uint32) (*types.Transaction, error) {
	return _Bridge.Contract.SetMinimalMaxSlippage(&_Bridge.TransactOpts, _minimalMaxSlippage)
}

// SetMinimalMaxSlippage is a paid mutator transaction binding the contract method 0x48234126.
//
// Solidity: function setMinimalMaxSlippage(uint32 _minimalMaxSlippage) returns()
func (_Bridge *BridgeTransactorSession) SetMinimalMaxSlippage(_minimalMaxSlippage uint32) (*types.Transaction, error) {
	return _Bridge.Contract.SetMinimalMaxSlippage(&_Bridge.TransactOpts, _minimalMaxSlippage)
}

// SetWrap is a paid mutator transaction binding the contract method 0x9ff9001a.
//
// Solidity: function setWrap(address _weth) returns()
func (_Bridge *BridgeTransactor) SetWrap(opts *bind.TransactOpts, _weth common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setWrap", _weth)
}

// SetWrap is a paid mutator transaction binding the contract method 0x9ff9001a.
//
// Solidity: function setWrap(address _weth) returns()
func (_Bridge *BridgeSession) SetWrap(_weth common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.SetWrap(&_Bridge.TransactOpts, _weth)
}

// SetWrap is a paid mutator transaction binding the contract method 0x9ff9001a.
//
// Solidity: function setWrap(address _weth) returns()
func (_Bridge *BridgeTransactorSession) SetWrap(_weth common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.SetWrap(&_Bridge.TransactOpts, _weth)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bridge *BridgeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bridge *BridgeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.TransferOwnership(&_Bridge.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bridge *BridgeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.TransferOwnership(&_Bridge.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Bridge *BridgeTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Bridge *BridgeSession) Unpause() (*types.Transaction, error) {
	return _Bridge.Contract.Unpause(&_Bridge.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Bridge *BridgeTransactorSession) Unpause() (*types.Transaction, error) {
	return _Bridge.Contract.Unpause(&_Bridge.TransactOpts)
}

// UpdateSigners is a paid mutator transaction binding the contract method 0xba2cb25c.
//
// Solidity: function updateSigners(uint256 _triggerTime, address[] _newSigners, uint256[] _newPowers, bytes[] _sigs, address[] _curSigners, uint256[] _curPowers) returns()
func (_Bridge *BridgeTransactor) UpdateSigners(opts *bind.TransactOpts, _triggerTime *big.Int, _newSigners []common.Address, _newPowers []*big.Int, _sigs [][]byte, _curSigners []common.Address, _curPowers []*big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "updateSigners", _triggerTime, _newSigners, _newPowers, _sigs, _curSigners, _curPowers)
}

// UpdateSigners is a paid mutator transaction binding the contract method 0xba2cb25c.
//
// Solidity: function updateSigners(uint256 _triggerTime, address[] _newSigners, uint256[] _newPowers, bytes[] _sigs, address[] _curSigners, uint256[] _curPowers) returns()
func (_Bridge *BridgeSession) UpdateSigners(_triggerTime *big.Int, _newSigners []common.Address, _newPowers []*big.Int, _sigs [][]byte, _curSigners []common.Address, _curPowers []*big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.UpdateSigners(&_Bridge.TransactOpts, _triggerTime, _newSigners, _newPowers, _sigs, _curSigners, _curPowers)
}

// UpdateSigners is a paid mutator transaction binding the contract method 0xba2cb25c.
//
// Solidity: function updateSigners(uint256 _triggerTime, address[] _newSigners, uint256[] _newPowers, bytes[] _sigs, address[] _curSigners, uint256[] _curPowers) returns()
func (_Bridge *BridgeTransactorSession) UpdateSigners(_triggerTime *big.Int, _newSigners []common.Address, _newPowers []*big.Int, _sigs [][]byte, _curSigners []common.Address, _curPowers []*big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.UpdateSigners(&_Bridge.TransactOpts, _triggerTime, _newSigners, _newPowers, _sigs, _curSigners, _curPowers)
}

// Withdraw is a paid mutator transaction binding the contract method 0xa21a9280.
//
// Solidity: function withdraw(bytes _wdmsg, bytes[] _sigs, address[] _signers, uint256[] _powers) returns()
func (_Bridge *BridgeTransactor) Withdraw(opts *bind.TransactOpts, _wdmsg []byte, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "withdraw", _wdmsg, _sigs, _signers, _powers)
}

// Withdraw is a paid mutator transaction binding the contract method 0xa21a9280.
//
// Solidity: function withdraw(bytes _wdmsg, bytes[] _sigs, address[] _signers, uint256[] _powers) returns()
func (_Bridge *BridgeSession) Withdraw(_wdmsg []byte, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.Withdraw(&_Bridge.TransactOpts, _wdmsg, _sigs, _signers, _powers)
}

// Withdraw is a paid mutator transaction binding the contract method 0xa21a9280.
//
// Solidity: function withdraw(bytes _wdmsg, bytes[] _sigs, address[] _signers, uint256[] _powers) returns()
func (_Bridge *BridgeTransactorSession) Withdraw(_wdmsg []byte, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.Withdraw(&_Bridge.TransactOpts, _wdmsg, _sigs, _signers, _powers)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Bridge *BridgeTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Bridge *BridgeSession) Receive() (*types.Transaction, error) {
	return _Bridge.Contract.Receive(&_Bridge.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Bridge *BridgeTransactorSession) Receive() (*types.Transaction, error) {
	return _Bridge.Contract.Receive(&_Bridge.TransactOpts)
}

// BridgeDelayPeriodUpdatedIterator is returned from FilterDelayPeriodUpdated and is used to iterate over the raw logs and unpacked data for DelayPeriodUpdated events raised by the Bridge contract.
type BridgeDelayPeriodUpdatedIterator struct {
	Event *BridgeDelayPeriodUpdated // Event containing the contract specifics and raw log

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
func (it *BridgeDelayPeriodUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeDelayPeriodUpdated)
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
		it.Event = new(BridgeDelayPeriodUpdated)
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
func (it *BridgeDelayPeriodUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeDelayPeriodUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeDelayPeriodUpdated represents a DelayPeriodUpdated event raised by the Bridge contract.
type BridgeDelayPeriodUpdated struct {
	Period *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDelayPeriodUpdated is a free log retrieval operation binding the contract event 0xc0a39f234199b125fb93713c4d067bdcebbf691087f87b79c0feb92b156ba8b6.
//
// Solidity: event DelayPeriodUpdated(uint256 period)
func (_Bridge *BridgeFilterer) FilterDelayPeriodUpdated(opts *bind.FilterOpts) (*BridgeDelayPeriodUpdatedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "DelayPeriodUpdated")
	if err != nil {
		return nil, err
	}
	return &BridgeDelayPeriodUpdatedIterator{contract: _Bridge.contract, event: "DelayPeriodUpdated", logs: logs, sub: sub}, nil
}

// WatchDelayPeriodUpdated is a free log subscription operation binding the contract event 0xc0a39f234199b125fb93713c4d067bdcebbf691087f87b79c0feb92b156ba8b6.
//
// Solidity: event DelayPeriodUpdated(uint256 period)
func (_Bridge *BridgeFilterer) WatchDelayPeriodUpdated(opts *bind.WatchOpts, sink chan<- *BridgeDelayPeriodUpdated) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "DelayPeriodUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeDelayPeriodUpdated)
				if err := _Bridge.contract.UnpackLog(event, "DelayPeriodUpdated", log); err != nil {
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

// ParseDelayPeriodUpdated is a log parse operation binding the contract event 0xc0a39f234199b125fb93713c4d067bdcebbf691087f87b79c0feb92b156ba8b6.
//
// Solidity: event DelayPeriodUpdated(uint256 period)
func (_Bridge *BridgeFilterer) ParseDelayPeriodUpdated(log types.Log) (*BridgeDelayPeriodUpdated, error) {
	event := new(BridgeDelayPeriodUpdated)
	if err := _Bridge.contract.UnpackLog(event, "DelayPeriodUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeDelayThresholdUpdatedIterator is returned from FilterDelayThresholdUpdated and is used to iterate over the raw logs and unpacked data for DelayThresholdUpdated events raised by the Bridge contract.
type BridgeDelayThresholdUpdatedIterator struct {
	Event *BridgeDelayThresholdUpdated // Event containing the contract specifics and raw log

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
func (it *BridgeDelayThresholdUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeDelayThresholdUpdated)
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
		it.Event = new(BridgeDelayThresholdUpdated)
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
func (it *BridgeDelayThresholdUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeDelayThresholdUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeDelayThresholdUpdated represents a DelayThresholdUpdated event raised by the Bridge contract.
type BridgeDelayThresholdUpdated struct {
	Token     common.Address
	Threshold *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDelayThresholdUpdated is a free log retrieval operation binding the contract event 0xceaad6533bfb481492fb3e08ef19297f46611b8fa9de5ef4cf8dc23a56ad09ce.
//
// Solidity: event DelayThresholdUpdated(address token, uint256 threshold)
func (_Bridge *BridgeFilterer) FilterDelayThresholdUpdated(opts *bind.FilterOpts) (*BridgeDelayThresholdUpdatedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "DelayThresholdUpdated")
	if err != nil {
		return nil, err
	}
	return &BridgeDelayThresholdUpdatedIterator{contract: _Bridge.contract, event: "DelayThresholdUpdated", logs: logs, sub: sub}, nil
}

// WatchDelayThresholdUpdated is a free log subscription operation binding the contract event 0xceaad6533bfb481492fb3e08ef19297f46611b8fa9de5ef4cf8dc23a56ad09ce.
//
// Solidity: event DelayThresholdUpdated(address token, uint256 threshold)
func (_Bridge *BridgeFilterer) WatchDelayThresholdUpdated(opts *bind.WatchOpts, sink chan<- *BridgeDelayThresholdUpdated) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "DelayThresholdUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeDelayThresholdUpdated)
				if err := _Bridge.contract.UnpackLog(event, "DelayThresholdUpdated", log); err != nil {
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

// ParseDelayThresholdUpdated is a log parse operation binding the contract event 0xceaad6533bfb481492fb3e08ef19297f46611b8fa9de5ef4cf8dc23a56ad09ce.
//
// Solidity: event DelayThresholdUpdated(address token, uint256 threshold)
func (_Bridge *BridgeFilterer) ParseDelayThresholdUpdated(log types.Log) (*BridgeDelayThresholdUpdated, error) {
	event := new(BridgeDelayThresholdUpdated)
	if err := _Bridge.contract.UnpackLog(event, "DelayThresholdUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeDelayedTransferAddedIterator is returned from FilterDelayedTransferAdded and is used to iterate over the raw logs and unpacked data for DelayedTransferAdded events raised by the Bridge contract.
type BridgeDelayedTransferAddedIterator struct {
	Event *BridgeDelayedTransferAdded // Event containing the contract specifics and raw log

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
func (it *BridgeDelayedTransferAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeDelayedTransferAdded)
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
		it.Event = new(BridgeDelayedTransferAdded)
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
func (it *BridgeDelayedTransferAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeDelayedTransferAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeDelayedTransferAdded represents a DelayedTransferAdded event raised by the Bridge contract.
type BridgeDelayedTransferAdded struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDelayedTransferAdded is a free log retrieval operation binding the contract event 0xcbcfffe5102114216a85d3aceb14ad4b81a3935b1b5c468fadf3889eb9c5dce6.
//
// Solidity: event DelayedTransferAdded(bytes32 id)
func (_Bridge *BridgeFilterer) FilterDelayedTransferAdded(opts *bind.FilterOpts) (*BridgeDelayedTransferAddedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "DelayedTransferAdded")
	if err != nil {
		return nil, err
	}
	return &BridgeDelayedTransferAddedIterator{contract: _Bridge.contract, event: "DelayedTransferAdded", logs: logs, sub: sub}, nil
}

// WatchDelayedTransferAdded is a free log subscription operation binding the contract event 0xcbcfffe5102114216a85d3aceb14ad4b81a3935b1b5c468fadf3889eb9c5dce6.
//
// Solidity: event DelayedTransferAdded(bytes32 id)
func (_Bridge *BridgeFilterer) WatchDelayedTransferAdded(opts *bind.WatchOpts, sink chan<- *BridgeDelayedTransferAdded) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "DelayedTransferAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeDelayedTransferAdded)
				if err := _Bridge.contract.UnpackLog(event, "DelayedTransferAdded", log); err != nil {
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

// ParseDelayedTransferAdded is a log parse operation binding the contract event 0xcbcfffe5102114216a85d3aceb14ad4b81a3935b1b5c468fadf3889eb9c5dce6.
//
// Solidity: event DelayedTransferAdded(bytes32 id)
func (_Bridge *BridgeFilterer) ParseDelayedTransferAdded(log types.Log) (*BridgeDelayedTransferAdded, error) {
	event := new(BridgeDelayedTransferAdded)
	if err := _Bridge.contract.UnpackLog(event, "DelayedTransferAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeDelayedTransferExecutedIterator is returned from FilterDelayedTransferExecuted and is used to iterate over the raw logs and unpacked data for DelayedTransferExecuted events raised by the Bridge contract.
type BridgeDelayedTransferExecutedIterator struct {
	Event *BridgeDelayedTransferExecuted // Event containing the contract specifics and raw log

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
func (it *BridgeDelayedTransferExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeDelayedTransferExecuted)
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
		it.Event = new(BridgeDelayedTransferExecuted)
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
func (it *BridgeDelayedTransferExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeDelayedTransferExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeDelayedTransferExecuted represents a DelayedTransferExecuted event raised by the Bridge contract.
type BridgeDelayedTransferExecuted struct {
	Id       [32]byte
	Receiver common.Address
	Token    common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDelayedTransferExecuted is a free log retrieval operation binding the contract event 0x3b40e5089937425d14cdd96947e5661868357e224af59bd8b24a4b8a330d4426.
//
// Solidity: event DelayedTransferExecuted(bytes32 id, address receiver, address token, uint256 amount)
func (_Bridge *BridgeFilterer) FilterDelayedTransferExecuted(opts *bind.FilterOpts) (*BridgeDelayedTransferExecutedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "DelayedTransferExecuted")
	if err != nil {
		return nil, err
	}
	return &BridgeDelayedTransferExecutedIterator{contract: _Bridge.contract, event: "DelayedTransferExecuted", logs: logs, sub: sub}, nil
}

// WatchDelayedTransferExecuted is a free log subscription operation binding the contract event 0x3b40e5089937425d14cdd96947e5661868357e224af59bd8b24a4b8a330d4426.
//
// Solidity: event DelayedTransferExecuted(bytes32 id, address receiver, address token, uint256 amount)
func (_Bridge *BridgeFilterer) WatchDelayedTransferExecuted(opts *bind.WatchOpts, sink chan<- *BridgeDelayedTransferExecuted) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "DelayedTransferExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeDelayedTransferExecuted)
				if err := _Bridge.contract.UnpackLog(event, "DelayedTransferExecuted", log); err != nil {
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

// ParseDelayedTransferExecuted is a log parse operation binding the contract event 0x3b40e5089937425d14cdd96947e5661868357e224af59bd8b24a4b8a330d4426.
//
// Solidity: event DelayedTransferExecuted(bytes32 id, address receiver, address token, uint256 amount)
func (_Bridge *BridgeFilterer) ParseDelayedTransferExecuted(log types.Log) (*BridgeDelayedTransferExecuted, error) {
	event := new(BridgeDelayedTransferExecuted)
	if err := _Bridge.contract.UnpackLog(event, "DelayedTransferExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeEpochLengthUpdatedIterator is returned from FilterEpochLengthUpdated and is used to iterate over the raw logs and unpacked data for EpochLengthUpdated events raised by the Bridge contract.
type BridgeEpochLengthUpdatedIterator struct {
	Event *BridgeEpochLengthUpdated // Event containing the contract specifics and raw log

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
func (it *BridgeEpochLengthUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeEpochLengthUpdated)
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
		it.Event = new(BridgeEpochLengthUpdated)
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
func (it *BridgeEpochLengthUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeEpochLengthUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeEpochLengthUpdated represents a EpochLengthUpdated event raised by the Bridge contract.
type BridgeEpochLengthUpdated struct {
	Length *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEpochLengthUpdated is a free log retrieval operation binding the contract event 0x2664fec2ff76486ac58ed087310855b648b15b9d19f3de8529e95f7c46b7d6b3.
//
// Solidity: event EpochLengthUpdated(uint256 length)
func (_Bridge *BridgeFilterer) FilterEpochLengthUpdated(opts *bind.FilterOpts) (*BridgeEpochLengthUpdatedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "EpochLengthUpdated")
	if err != nil {
		return nil, err
	}
	return &BridgeEpochLengthUpdatedIterator{contract: _Bridge.contract, event: "EpochLengthUpdated", logs: logs, sub: sub}, nil
}

// WatchEpochLengthUpdated is a free log subscription operation binding the contract event 0x2664fec2ff76486ac58ed087310855b648b15b9d19f3de8529e95f7c46b7d6b3.
//
// Solidity: event EpochLengthUpdated(uint256 length)
func (_Bridge *BridgeFilterer) WatchEpochLengthUpdated(opts *bind.WatchOpts, sink chan<- *BridgeEpochLengthUpdated) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "EpochLengthUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeEpochLengthUpdated)
				if err := _Bridge.contract.UnpackLog(event, "EpochLengthUpdated", log); err != nil {
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

// ParseEpochLengthUpdated is a log parse operation binding the contract event 0x2664fec2ff76486ac58ed087310855b648b15b9d19f3de8529e95f7c46b7d6b3.
//
// Solidity: event EpochLengthUpdated(uint256 length)
func (_Bridge *BridgeFilterer) ParseEpochLengthUpdated(log types.Log) (*BridgeEpochLengthUpdated, error) {
	event := new(BridgeEpochLengthUpdated)
	if err := _Bridge.contract.UnpackLog(event, "EpochLengthUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeEpochVolumeUpdatedIterator is returned from FilterEpochVolumeUpdated and is used to iterate over the raw logs and unpacked data for EpochVolumeUpdated events raised by the Bridge contract.
type BridgeEpochVolumeUpdatedIterator struct {
	Event *BridgeEpochVolumeUpdated // Event containing the contract specifics and raw log

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
func (it *BridgeEpochVolumeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeEpochVolumeUpdated)
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
		it.Event = new(BridgeEpochVolumeUpdated)
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
func (it *BridgeEpochVolumeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeEpochVolumeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeEpochVolumeUpdated represents a EpochVolumeUpdated event raised by the Bridge contract.
type BridgeEpochVolumeUpdated struct {
	Token common.Address
	Cap   *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterEpochVolumeUpdated is a free log retrieval operation binding the contract event 0x608e49c22994f20b5d3496dca088b88dfd81b4a3e8cc3809ea1e10a320107e89.
//
// Solidity: event EpochVolumeUpdated(address token, uint256 cap)
func (_Bridge *BridgeFilterer) FilterEpochVolumeUpdated(opts *bind.FilterOpts) (*BridgeEpochVolumeUpdatedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "EpochVolumeUpdated")
	if err != nil {
		return nil, err
	}
	return &BridgeEpochVolumeUpdatedIterator{contract: _Bridge.contract, event: "EpochVolumeUpdated", logs: logs, sub: sub}, nil
}

// WatchEpochVolumeUpdated is a free log subscription operation binding the contract event 0x608e49c22994f20b5d3496dca088b88dfd81b4a3e8cc3809ea1e10a320107e89.
//
// Solidity: event EpochVolumeUpdated(address token, uint256 cap)
func (_Bridge *BridgeFilterer) WatchEpochVolumeUpdated(opts *bind.WatchOpts, sink chan<- *BridgeEpochVolumeUpdated) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "EpochVolumeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeEpochVolumeUpdated)
				if err := _Bridge.contract.UnpackLog(event, "EpochVolumeUpdated", log); err != nil {
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

// ParseEpochVolumeUpdated is a log parse operation binding the contract event 0x608e49c22994f20b5d3496dca088b88dfd81b4a3e8cc3809ea1e10a320107e89.
//
// Solidity: event EpochVolumeUpdated(address token, uint256 cap)
func (_Bridge *BridgeFilterer) ParseEpochVolumeUpdated(log types.Log) (*BridgeEpochVolumeUpdated, error) {
	event := new(BridgeEpochVolumeUpdated)
	if err := _Bridge.contract.UnpackLog(event, "EpochVolumeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeGovernorAddedIterator is returned from FilterGovernorAdded and is used to iterate over the raw logs and unpacked data for GovernorAdded events raised by the Bridge contract.
type BridgeGovernorAddedIterator struct {
	Event *BridgeGovernorAdded // Event containing the contract specifics and raw log

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
func (it *BridgeGovernorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeGovernorAdded)
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
		it.Event = new(BridgeGovernorAdded)
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
func (it *BridgeGovernorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeGovernorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeGovernorAdded represents a GovernorAdded event raised by the Bridge contract.
type BridgeGovernorAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterGovernorAdded is a free log retrieval operation binding the contract event 0xdc5a48d79e2e147530ff63ecdbed5a5a66adb9d5cf339384d5d076da197c40b5.
//
// Solidity: event GovernorAdded(address account)
func (_Bridge *BridgeFilterer) FilterGovernorAdded(opts *bind.FilterOpts) (*BridgeGovernorAddedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "GovernorAdded")
	if err != nil {
		return nil, err
	}
	return &BridgeGovernorAddedIterator{contract: _Bridge.contract, event: "GovernorAdded", logs: logs, sub: sub}, nil
}

// WatchGovernorAdded is a free log subscription operation binding the contract event 0xdc5a48d79e2e147530ff63ecdbed5a5a66adb9d5cf339384d5d076da197c40b5.
//
// Solidity: event GovernorAdded(address account)
func (_Bridge *BridgeFilterer) WatchGovernorAdded(opts *bind.WatchOpts, sink chan<- *BridgeGovernorAdded) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "GovernorAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeGovernorAdded)
				if err := _Bridge.contract.UnpackLog(event, "GovernorAdded", log); err != nil {
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

// ParseGovernorAdded is a log parse operation binding the contract event 0xdc5a48d79e2e147530ff63ecdbed5a5a66adb9d5cf339384d5d076da197c40b5.
//
// Solidity: event GovernorAdded(address account)
func (_Bridge *BridgeFilterer) ParseGovernorAdded(log types.Log) (*BridgeGovernorAdded, error) {
	event := new(BridgeGovernorAdded)
	if err := _Bridge.contract.UnpackLog(event, "GovernorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeGovernorRemovedIterator is returned from FilterGovernorRemoved and is used to iterate over the raw logs and unpacked data for GovernorRemoved events raised by the Bridge contract.
type BridgeGovernorRemovedIterator struct {
	Event *BridgeGovernorRemoved // Event containing the contract specifics and raw log

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
func (it *BridgeGovernorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeGovernorRemoved)
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
		it.Event = new(BridgeGovernorRemoved)
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
func (it *BridgeGovernorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeGovernorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeGovernorRemoved represents a GovernorRemoved event raised by the Bridge contract.
type BridgeGovernorRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterGovernorRemoved is a free log retrieval operation binding the contract event 0x1ebe834e73d60a5fec822c1e1727d34bc79f2ad977ed504581cc1822fe20fb5b.
//
// Solidity: event GovernorRemoved(address account)
func (_Bridge *BridgeFilterer) FilterGovernorRemoved(opts *bind.FilterOpts) (*BridgeGovernorRemovedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "GovernorRemoved")
	if err != nil {
		return nil, err
	}
	return &BridgeGovernorRemovedIterator{contract: _Bridge.contract, event: "GovernorRemoved", logs: logs, sub: sub}, nil
}

// WatchGovernorRemoved is a free log subscription operation binding the contract event 0x1ebe834e73d60a5fec822c1e1727d34bc79f2ad977ed504581cc1822fe20fb5b.
//
// Solidity: event GovernorRemoved(address account)
func (_Bridge *BridgeFilterer) WatchGovernorRemoved(opts *bind.WatchOpts, sink chan<- *BridgeGovernorRemoved) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "GovernorRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeGovernorRemoved)
				if err := _Bridge.contract.UnpackLog(event, "GovernorRemoved", log); err != nil {
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

// ParseGovernorRemoved is a log parse operation binding the contract event 0x1ebe834e73d60a5fec822c1e1727d34bc79f2ad977ed504581cc1822fe20fb5b.
//
// Solidity: event GovernorRemoved(address account)
func (_Bridge *BridgeFilterer) ParseGovernorRemoved(log types.Log) (*BridgeGovernorRemoved, error) {
	event := new(BridgeGovernorRemoved)
	if err := _Bridge.contract.UnpackLog(event, "GovernorRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeLiquidityAddedIterator is returned from FilterLiquidityAdded and is used to iterate over the raw logs and unpacked data for LiquidityAdded events raised by the Bridge contract.
type BridgeLiquidityAddedIterator struct {
	Event *BridgeLiquidityAdded // Event containing the contract specifics and raw log

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
func (it *BridgeLiquidityAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeLiquidityAdded)
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
		it.Event = new(BridgeLiquidityAdded)
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
func (it *BridgeLiquidityAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeLiquidityAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeLiquidityAdded represents a LiquidityAdded event raised by the Bridge contract.
type BridgeLiquidityAdded struct {
	Seqnum   uint64
	Provider common.Address
	Token    common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLiquidityAdded is a free log retrieval operation binding the contract event 0xd5d28426c3248963b1719df49aa4c665120372e02c8249bbea03d019c39ce764.
//
// Solidity: event LiquidityAdded(uint64 seqnum, address provider, address token, uint256 amount)
func (_Bridge *BridgeFilterer) FilterLiquidityAdded(opts *bind.FilterOpts) (*BridgeLiquidityAddedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "LiquidityAdded")
	if err != nil {
		return nil, err
	}
	return &BridgeLiquidityAddedIterator{contract: _Bridge.contract, event: "LiquidityAdded", logs: logs, sub: sub}, nil
}

// WatchLiquidityAdded is a free log subscription operation binding the contract event 0xd5d28426c3248963b1719df49aa4c665120372e02c8249bbea03d019c39ce764.
//
// Solidity: event LiquidityAdded(uint64 seqnum, address provider, address token, uint256 amount)
func (_Bridge *BridgeFilterer) WatchLiquidityAdded(opts *bind.WatchOpts, sink chan<- *BridgeLiquidityAdded) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "LiquidityAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeLiquidityAdded)
				if err := _Bridge.contract.UnpackLog(event, "LiquidityAdded", log); err != nil {
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

// ParseLiquidityAdded is a log parse operation binding the contract event 0xd5d28426c3248963b1719df49aa4c665120372e02c8249bbea03d019c39ce764.
//
// Solidity: event LiquidityAdded(uint64 seqnum, address provider, address token, uint256 amount)
func (_Bridge *BridgeFilterer) ParseLiquidityAdded(log types.Log) (*BridgeLiquidityAdded, error) {
	event := new(BridgeLiquidityAdded)
	if err := _Bridge.contract.UnpackLog(event, "LiquidityAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeMaxSendUpdatedIterator is returned from FilterMaxSendUpdated and is used to iterate over the raw logs and unpacked data for MaxSendUpdated events raised by the Bridge contract.
type BridgeMaxSendUpdatedIterator struct {
	Event *BridgeMaxSendUpdated // Event containing the contract specifics and raw log

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
func (it *BridgeMaxSendUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeMaxSendUpdated)
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
		it.Event = new(BridgeMaxSendUpdated)
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
func (it *BridgeMaxSendUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeMaxSendUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeMaxSendUpdated represents a MaxSendUpdated event raised by the Bridge contract.
type BridgeMaxSendUpdated struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMaxSendUpdated is a free log retrieval operation binding the contract event 0x4f12d1a5bfb3ccd3719255d4d299d808d50cdca9a0a5c2b3a5aaa7edde73052c.
//
// Solidity: event MaxSendUpdated(address token, uint256 amount)
func (_Bridge *BridgeFilterer) FilterMaxSendUpdated(opts *bind.FilterOpts) (*BridgeMaxSendUpdatedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "MaxSendUpdated")
	if err != nil {
		return nil, err
	}
	return &BridgeMaxSendUpdatedIterator{contract: _Bridge.contract, event: "MaxSendUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxSendUpdated is a free log subscription operation binding the contract event 0x4f12d1a5bfb3ccd3719255d4d299d808d50cdca9a0a5c2b3a5aaa7edde73052c.
//
// Solidity: event MaxSendUpdated(address token, uint256 amount)
func (_Bridge *BridgeFilterer) WatchMaxSendUpdated(opts *bind.WatchOpts, sink chan<- *BridgeMaxSendUpdated) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "MaxSendUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeMaxSendUpdated)
				if err := _Bridge.contract.UnpackLog(event, "MaxSendUpdated", log); err != nil {
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

// ParseMaxSendUpdated is a log parse operation binding the contract event 0x4f12d1a5bfb3ccd3719255d4d299d808d50cdca9a0a5c2b3a5aaa7edde73052c.
//
// Solidity: event MaxSendUpdated(address token, uint256 amount)
func (_Bridge *BridgeFilterer) ParseMaxSendUpdated(log types.Log) (*BridgeMaxSendUpdated, error) {
	event := new(BridgeMaxSendUpdated)
	if err := _Bridge.contract.UnpackLog(event, "MaxSendUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeMinAddUpdatedIterator is returned from FilterMinAddUpdated and is used to iterate over the raw logs and unpacked data for MinAddUpdated events raised by the Bridge contract.
type BridgeMinAddUpdatedIterator struct {
	Event *BridgeMinAddUpdated // Event containing the contract specifics and raw log

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
func (it *BridgeMinAddUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeMinAddUpdated)
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
		it.Event = new(BridgeMinAddUpdated)
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
func (it *BridgeMinAddUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeMinAddUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeMinAddUpdated represents a MinAddUpdated event raised by the Bridge contract.
type BridgeMinAddUpdated struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMinAddUpdated is a free log retrieval operation binding the contract event 0xc56b0d14c4940515800d94ebbd0f3f5d8cc58ba1109c12536bd993b72e466e4f.
//
// Solidity: event MinAddUpdated(address token, uint256 amount)
func (_Bridge *BridgeFilterer) FilterMinAddUpdated(opts *bind.FilterOpts) (*BridgeMinAddUpdatedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "MinAddUpdated")
	if err != nil {
		return nil, err
	}
	return &BridgeMinAddUpdatedIterator{contract: _Bridge.contract, event: "MinAddUpdated", logs: logs, sub: sub}, nil
}

// WatchMinAddUpdated is a free log subscription operation binding the contract event 0xc56b0d14c4940515800d94ebbd0f3f5d8cc58ba1109c12536bd993b72e466e4f.
//
// Solidity: event MinAddUpdated(address token, uint256 amount)
func (_Bridge *BridgeFilterer) WatchMinAddUpdated(opts *bind.WatchOpts, sink chan<- *BridgeMinAddUpdated) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "MinAddUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeMinAddUpdated)
				if err := _Bridge.contract.UnpackLog(event, "MinAddUpdated", log); err != nil {
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

// ParseMinAddUpdated is a log parse operation binding the contract event 0xc56b0d14c4940515800d94ebbd0f3f5d8cc58ba1109c12536bd993b72e466e4f.
//
// Solidity: event MinAddUpdated(address token, uint256 amount)
func (_Bridge *BridgeFilterer) ParseMinAddUpdated(log types.Log) (*BridgeMinAddUpdated, error) {
	event := new(BridgeMinAddUpdated)
	if err := _Bridge.contract.UnpackLog(event, "MinAddUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeMinSendUpdatedIterator is returned from FilterMinSendUpdated and is used to iterate over the raw logs and unpacked data for MinSendUpdated events raised by the Bridge contract.
type BridgeMinSendUpdatedIterator struct {
	Event *BridgeMinSendUpdated // Event containing the contract specifics and raw log

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
func (it *BridgeMinSendUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeMinSendUpdated)
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
		it.Event = new(BridgeMinSendUpdated)
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
func (it *BridgeMinSendUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeMinSendUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeMinSendUpdated represents a MinSendUpdated event raised by the Bridge contract.
type BridgeMinSendUpdated struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMinSendUpdated is a free log retrieval operation binding the contract event 0x8b59d386e660418a48d742213ad5ce7c4dd51ae81f30e4e2c387f17d907010c9.
//
// Solidity: event MinSendUpdated(address token, uint256 amount)
func (_Bridge *BridgeFilterer) FilterMinSendUpdated(opts *bind.FilterOpts) (*BridgeMinSendUpdatedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "MinSendUpdated")
	if err != nil {
		return nil, err
	}
	return &BridgeMinSendUpdatedIterator{contract: _Bridge.contract, event: "MinSendUpdated", logs: logs, sub: sub}, nil
}

// WatchMinSendUpdated is a free log subscription operation binding the contract event 0x8b59d386e660418a48d742213ad5ce7c4dd51ae81f30e4e2c387f17d907010c9.
//
// Solidity: event MinSendUpdated(address token, uint256 amount)
func (_Bridge *BridgeFilterer) WatchMinSendUpdated(opts *bind.WatchOpts, sink chan<- *BridgeMinSendUpdated) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "MinSendUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeMinSendUpdated)
				if err := _Bridge.contract.UnpackLog(event, "MinSendUpdated", log); err != nil {
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

// ParseMinSendUpdated is a log parse operation binding the contract event 0x8b59d386e660418a48d742213ad5ce7c4dd51ae81f30e4e2c387f17d907010c9.
//
// Solidity: event MinSendUpdated(address token, uint256 amount)
func (_Bridge *BridgeFilterer) ParseMinSendUpdated(log types.Log) (*BridgeMinSendUpdated, error) {
	event := new(BridgeMinSendUpdated)
	if err := _Bridge.contract.UnpackLog(event, "MinSendUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Bridge contract.
type BridgeOwnershipTransferredIterator struct {
	Event *BridgeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BridgeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeOwnershipTransferred)
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
		it.Event = new(BridgeOwnershipTransferred)
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
func (it *BridgeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeOwnershipTransferred represents a OwnershipTransferred event raised by the Bridge contract.
type BridgeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bridge *BridgeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BridgeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BridgeOwnershipTransferredIterator{contract: _Bridge.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bridge *BridgeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BridgeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeOwnershipTransferred)
				if err := _Bridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Bridge *BridgeFilterer) ParseOwnershipTransferred(log types.Log) (*BridgeOwnershipTransferred, error) {
	event := new(BridgeOwnershipTransferred)
	if err := _Bridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Bridge contract.
type BridgePausedIterator struct {
	Event *BridgePaused // Event containing the contract specifics and raw log

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
func (it *BridgePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgePaused)
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
		it.Event = new(BridgePaused)
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
func (it *BridgePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgePaused represents a Paused event raised by the Bridge contract.
type BridgePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Bridge *BridgeFilterer) FilterPaused(opts *bind.FilterOpts) (*BridgePausedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &BridgePausedIterator{contract: _Bridge.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Bridge *BridgeFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *BridgePaused) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgePaused)
				if err := _Bridge.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Bridge *BridgeFilterer) ParsePaused(log types.Log) (*BridgePaused, error) {
	event := new(BridgePaused)
	if err := _Bridge.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgePauserAddedIterator is returned from FilterPauserAdded and is used to iterate over the raw logs and unpacked data for PauserAdded events raised by the Bridge contract.
type BridgePauserAddedIterator struct {
	Event *BridgePauserAdded // Event containing the contract specifics and raw log

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
func (it *BridgePauserAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgePauserAdded)
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
		it.Event = new(BridgePauserAdded)
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
func (it *BridgePauserAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgePauserAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgePauserAdded represents a PauserAdded event raised by the Bridge contract.
type BridgePauserAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserAdded is a free log retrieval operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address account)
func (_Bridge *BridgeFilterer) FilterPauserAdded(opts *bind.FilterOpts) (*BridgePauserAddedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "PauserAdded")
	if err != nil {
		return nil, err
	}
	return &BridgePauserAddedIterator{contract: _Bridge.contract, event: "PauserAdded", logs: logs, sub: sub}, nil
}

// WatchPauserAdded is a free log subscription operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address account)
func (_Bridge *BridgeFilterer) WatchPauserAdded(opts *bind.WatchOpts, sink chan<- *BridgePauserAdded) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "PauserAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgePauserAdded)
				if err := _Bridge.contract.UnpackLog(event, "PauserAdded", log); err != nil {
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

// ParsePauserAdded is a log parse operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address account)
func (_Bridge *BridgeFilterer) ParsePauserAdded(log types.Log) (*BridgePauserAdded, error) {
	event := new(BridgePauserAdded)
	if err := _Bridge.contract.UnpackLog(event, "PauserAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgePauserRemovedIterator is returned from FilterPauserRemoved and is used to iterate over the raw logs and unpacked data for PauserRemoved events raised by the Bridge contract.
type BridgePauserRemovedIterator struct {
	Event *BridgePauserRemoved // Event containing the contract specifics and raw log

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
func (it *BridgePauserRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgePauserRemoved)
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
		it.Event = new(BridgePauserRemoved)
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
func (it *BridgePauserRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgePauserRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgePauserRemoved represents a PauserRemoved event raised by the Bridge contract.
type BridgePauserRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserRemoved is a free log retrieval operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address account)
func (_Bridge *BridgeFilterer) FilterPauserRemoved(opts *bind.FilterOpts) (*BridgePauserRemovedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "PauserRemoved")
	if err != nil {
		return nil, err
	}
	return &BridgePauserRemovedIterator{contract: _Bridge.contract, event: "PauserRemoved", logs: logs, sub: sub}, nil
}

// WatchPauserRemoved is a free log subscription operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address account)
func (_Bridge *BridgeFilterer) WatchPauserRemoved(opts *bind.WatchOpts, sink chan<- *BridgePauserRemoved) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "PauserRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgePauserRemoved)
				if err := _Bridge.contract.UnpackLog(event, "PauserRemoved", log); err != nil {
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

// ParsePauserRemoved is a log parse operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address account)
func (_Bridge *BridgeFilterer) ParsePauserRemoved(log types.Log) (*BridgePauserRemoved, error) {
	event := new(BridgePauserRemoved)
	if err := _Bridge.contract.UnpackLog(event, "PauserRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeRelayIterator is returned from FilterRelay and is used to iterate over the raw logs and unpacked data for Relay events raised by the Bridge contract.
type BridgeRelayIterator struct {
	Event *BridgeRelay // Event containing the contract specifics and raw log

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
func (it *BridgeRelayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeRelay)
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
		it.Event = new(BridgeRelay)
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
func (it *BridgeRelayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRelayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeRelay represents a Relay event raised by the Bridge contract.
type BridgeRelay struct {
	TransferId    [32]byte
	Sender        common.Address
	Receiver      common.Address
	Token         common.Address
	Amount        *big.Int
	SrcChainId    uint64
	SrcTransferId [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRelay is a free log retrieval operation binding the contract event 0x79fa08de5149d912dce8e5e8da7a7c17ccdf23dd5d3bfe196802e6eb86347c7c.
//
// Solidity: event Relay(bytes32 transferId, address sender, address receiver, address token, uint256 amount, uint64 srcChainId, bytes32 srcTransferId)
func (_Bridge *BridgeFilterer) FilterRelay(opts *bind.FilterOpts) (*BridgeRelayIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Relay")
	if err != nil {
		return nil, err
	}
	return &BridgeRelayIterator{contract: _Bridge.contract, event: "Relay", logs: logs, sub: sub}, nil
}

// WatchRelay is a free log subscription operation binding the contract event 0x79fa08de5149d912dce8e5e8da7a7c17ccdf23dd5d3bfe196802e6eb86347c7c.
//
// Solidity: event Relay(bytes32 transferId, address sender, address receiver, address token, uint256 amount, uint64 srcChainId, bytes32 srcTransferId)
func (_Bridge *BridgeFilterer) WatchRelay(opts *bind.WatchOpts, sink chan<- *BridgeRelay) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "Relay")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeRelay)
				if err := _Bridge.contract.UnpackLog(event, "Relay", log); err != nil {
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

// ParseRelay is a log parse operation binding the contract event 0x79fa08de5149d912dce8e5e8da7a7c17ccdf23dd5d3bfe196802e6eb86347c7c.
//
// Solidity: event Relay(bytes32 transferId, address sender, address receiver, address token, uint256 amount, uint64 srcChainId, bytes32 srcTransferId)
func (_Bridge *BridgeFilterer) ParseRelay(log types.Log) (*BridgeRelay, error) {
	event := new(BridgeRelay)
	if err := _Bridge.contract.UnpackLog(event, "Relay", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeResetNotificationIterator is returned from FilterResetNotification and is used to iterate over the raw logs and unpacked data for ResetNotification events raised by the Bridge contract.
type BridgeResetNotificationIterator struct {
	Event *BridgeResetNotification // Event containing the contract specifics and raw log

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
func (it *BridgeResetNotificationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeResetNotification)
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
		it.Event = new(BridgeResetNotification)
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
func (it *BridgeResetNotificationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeResetNotificationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeResetNotification represents a ResetNotification event raised by the Bridge contract.
type BridgeResetNotification struct {
	ResetTime *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterResetNotification is a free log retrieval operation binding the contract event 0x68e825132f7d4bc837dea2d64ac9fc19912bf0224b67f9317d8f1a917f5304a1.
//
// Solidity: event ResetNotification(uint256 resetTime)
func (_Bridge *BridgeFilterer) FilterResetNotification(opts *bind.FilterOpts) (*BridgeResetNotificationIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "ResetNotification")
	if err != nil {
		return nil, err
	}
	return &BridgeResetNotificationIterator{contract: _Bridge.contract, event: "ResetNotification", logs: logs, sub: sub}, nil
}

// WatchResetNotification is a free log subscription operation binding the contract event 0x68e825132f7d4bc837dea2d64ac9fc19912bf0224b67f9317d8f1a917f5304a1.
//
// Solidity: event ResetNotification(uint256 resetTime)
func (_Bridge *BridgeFilterer) WatchResetNotification(opts *bind.WatchOpts, sink chan<- *BridgeResetNotification) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "ResetNotification")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeResetNotification)
				if err := _Bridge.contract.UnpackLog(event, "ResetNotification", log); err != nil {
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

// ParseResetNotification is a log parse operation binding the contract event 0x68e825132f7d4bc837dea2d64ac9fc19912bf0224b67f9317d8f1a917f5304a1.
//
// Solidity: event ResetNotification(uint256 resetTime)
func (_Bridge *BridgeFilterer) ParseResetNotification(log types.Log) (*BridgeResetNotification, error) {
	event := new(BridgeResetNotification)
	if err := _Bridge.contract.UnpackLog(event, "ResetNotification", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeSendIterator is returned from FilterSend and is used to iterate over the raw logs and unpacked data for Send events raised by the Bridge contract.
type BridgeSendIterator struct {
	Event *BridgeSend // Event containing the contract specifics and raw log

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
func (it *BridgeSendIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeSend)
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
		it.Event = new(BridgeSend)
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
func (it *BridgeSendIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeSendIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeSend represents a Send event raised by the Bridge contract.
type BridgeSend struct {
	TransferId  [32]byte
	Sender      common.Address
	Receiver    common.Address
	Token       common.Address
	Amount      *big.Int
	DstChainId  uint64
	Nonce       uint64
	MaxSlippage uint32
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSend is a free log retrieval operation binding the contract event 0x89d8051e597ab4178a863a5190407b98abfeff406aa8db90c59af76612e58f01.
//
// Solidity: event Send(bytes32 transferId, address sender, address receiver, address token, uint256 amount, uint64 dstChainId, uint64 nonce, uint32 maxSlippage)
func (_Bridge *BridgeFilterer) FilterSend(opts *bind.FilterOpts) (*BridgeSendIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Send")
	if err != nil {
		return nil, err
	}
	return &BridgeSendIterator{contract: _Bridge.contract, event: "Send", logs: logs, sub: sub}, nil
}

// WatchSend is a free log subscription operation binding the contract event 0x89d8051e597ab4178a863a5190407b98abfeff406aa8db90c59af76612e58f01.
//
// Solidity: event Send(bytes32 transferId, address sender, address receiver, address token, uint256 amount, uint64 dstChainId, uint64 nonce, uint32 maxSlippage)
func (_Bridge *BridgeFilterer) WatchSend(opts *bind.WatchOpts, sink chan<- *BridgeSend) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "Send")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeSend)
				if err := _Bridge.contract.UnpackLog(event, "Send", log); err != nil {
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

// ParseSend is a log parse operation binding the contract event 0x89d8051e597ab4178a863a5190407b98abfeff406aa8db90c59af76612e58f01.
//
// Solidity: event Send(bytes32 transferId, address sender, address receiver, address token, uint256 amount, uint64 dstChainId, uint64 nonce, uint32 maxSlippage)
func (_Bridge *BridgeFilterer) ParseSend(log types.Log) (*BridgeSend, error) {
	event := new(BridgeSend)
	if err := _Bridge.contract.UnpackLog(event, "Send", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeSignersUpdatedIterator is returned from FilterSignersUpdated and is used to iterate over the raw logs and unpacked data for SignersUpdated events raised by the Bridge contract.
type BridgeSignersUpdatedIterator struct {
	Event *BridgeSignersUpdated // Event containing the contract specifics and raw log

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
func (it *BridgeSignersUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeSignersUpdated)
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
		it.Event = new(BridgeSignersUpdated)
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
func (it *BridgeSignersUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeSignersUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeSignersUpdated represents a SignersUpdated event raised by the Bridge contract.
type BridgeSignersUpdated struct {
	Signers []common.Address
	Powers  []*big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSignersUpdated is a free log retrieval operation binding the contract event 0xf126123539a68393c55697f617e7d1148e371988daed246c2f41da99965a23f8.
//
// Solidity: event SignersUpdated(address[] _signers, uint256[] _powers)
func (_Bridge *BridgeFilterer) FilterSignersUpdated(opts *bind.FilterOpts) (*BridgeSignersUpdatedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "SignersUpdated")
	if err != nil {
		return nil, err
	}
	return &BridgeSignersUpdatedIterator{contract: _Bridge.contract, event: "SignersUpdated", logs: logs, sub: sub}, nil
}

// WatchSignersUpdated is a free log subscription operation binding the contract event 0xf126123539a68393c55697f617e7d1148e371988daed246c2f41da99965a23f8.
//
// Solidity: event SignersUpdated(address[] _signers, uint256[] _powers)
func (_Bridge *BridgeFilterer) WatchSignersUpdated(opts *bind.WatchOpts, sink chan<- *BridgeSignersUpdated) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "SignersUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeSignersUpdated)
				if err := _Bridge.contract.UnpackLog(event, "SignersUpdated", log); err != nil {
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

// ParseSignersUpdated is a log parse operation binding the contract event 0xf126123539a68393c55697f617e7d1148e371988daed246c2f41da99965a23f8.
//
// Solidity: event SignersUpdated(address[] _signers, uint256[] _powers)
func (_Bridge *BridgeFilterer) ParseSignersUpdated(log types.Log) (*BridgeSignersUpdated, error) {
	event := new(BridgeSignersUpdated)
	if err := _Bridge.contract.UnpackLog(event, "SignersUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Bridge contract.
type BridgeUnpausedIterator struct {
	Event *BridgeUnpaused // Event containing the contract specifics and raw log

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
func (it *BridgeUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeUnpaused)
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
		it.Event = new(BridgeUnpaused)
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
func (it *BridgeUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeUnpaused represents a Unpaused event raised by the Bridge contract.
type BridgeUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Bridge *BridgeFilterer) FilterUnpaused(opts *bind.FilterOpts) (*BridgeUnpausedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &BridgeUnpausedIterator{contract: _Bridge.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Bridge *BridgeFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *BridgeUnpaused) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeUnpaused)
				if err := _Bridge.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Bridge *BridgeFilterer) ParseUnpaused(log types.Log) (*BridgeUnpaused, error) {
	event := new(BridgeUnpaused)
	if err := _Bridge.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeWithdrawDoneIterator is returned from FilterWithdrawDone and is used to iterate over the raw logs and unpacked data for WithdrawDone events raised by the Bridge contract.
type BridgeWithdrawDoneIterator struct {
	Event *BridgeWithdrawDone // Event containing the contract specifics and raw log

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
func (it *BridgeWithdrawDoneIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeWithdrawDone)
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
		it.Event = new(BridgeWithdrawDone)
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
func (it *BridgeWithdrawDoneIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeWithdrawDoneIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeWithdrawDone represents a WithdrawDone event raised by the Bridge contract.
type BridgeWithdrawDone struct {
	WithdrawId [32]byte
	Seqnum     uint64
	Receiver   common.Address
	Token      common.Address
	Amount     *big.Int
	Refid      [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWithdrawDone is a free log retrieval operation binding the contract event 0x48a1ab26f3aa7b62bb6b6e8eed182f292b84eb7b006c0254386b268af20774be.
//
// Solidity: event WithdrawDone(bytes32 withdrawId, uint64 seqnum, address receiver, address token, uint256 amount, bytes32 refid)
func (_Bridge *BridgeFilterer) FilterWithdrawDone(opts *bind.FilterOpts) (*BridgeWithdrawDoneIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "WithdrawDone")
	if err != nil {
		return nil, err
	}
	return &BridgeWithdrawDoneIterator{contract: _Bridge.contract, event: "WithdrawDone", logs: logs, sub: sub}, nil
}

// WatchWithdrawDone is a free log subscription operation binding the contract event 0x48a1ab26f3aa7b62bb6b6e8eed182f292b84eb7b006c0254386b268af20774be.
//
// Solidity: event WithdrawDone(bytes32 withdrawId, uint64 seqnum, address receiver, address token, uint256 amount, bytes32 refid)
func (_Bridge *BridgeFilterer) WatchWithdrawDone(opts *bind.WatchOpts, sink chan<- *BridgeWithdrawDone) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "WithdrawDone")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeWithdrawDone)
				if err := _Bridge.contract.UnpackLog(event, "WithdrawDone", log); err != nil {
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

// ParseWithdrawDone is a log parse operation binding the contract event 0x48a1ab26f3aa7b62bb6b6e8eed182f292b84eb7b006c0254386b268af20774be.
//
// Solidity: event WithdrawDone(bytes32 withdrawId, uint64 seqnum, address receiver, address token, uint256 amount, bytes32 refid)
func (_Bridge *BridgeFilterer) ParseWithdrawDone(log types.Log) (*BridgeWithdrawDone, error) {
	event := new(BridgeWithdrawDone)
	if err := _Bridge.contract.UnpackLog(event, "WithdrawDone", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
