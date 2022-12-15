// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package hop

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

// BridgeTransferRoot is an auto generated low-level Go binding around an user-defined struct.
type BridgeTransferRoot struct {
	Total           *big.Int
	AmountWithdrawn *big.Int
	CreatedAt       *big.Int
}

// L2BridgeMetaData contains all meta data concerning the L2Bridge contract.
var L2BridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractI_L2_PolygonMessengerProxy\",\"name\":\"_messengerProxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"l1Governance\",\"type\":\"address\"},{\"internalType\":\"contractHopBridgeToken\",\"name\":\"hToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"l1BridgeAddress\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"activeChainIds\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"bonders\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newBonder\",\"type\":\"address\"}],\"name\":\"BonderAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousBonder\",\"type\":\"address\"}],\"name\":\"BonderRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"L1_BridgeMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bonder\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"rootHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalBondsSettled\",\"type\":\"uint256\"}],\"name\":\"MultipleWithdrawalsSettled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Stake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"relayerFee\",\"type\":\"uint256\"}],\"name\":\"TransferFromL1Completed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"rootHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"}],\"name\":\"TransferRootSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transferId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"transferNonce\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bonderFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"TransferSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"rootHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rootCommittedAt\",\"type\":\"uint256\"}],\"name\":\"TransfersCommitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Unstake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bonder\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transferId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"rootHash\",\"type\":\"bytes32\"}],\"name\":\"WithdrawalBondSettled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transferId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawalBonded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transferId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"transferNonce\",\"type\":\"bytes32\"}],\"name\":\"Withdrew\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"activeChainIds\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"chainIds\",\"type\":\"uint256[]\"}],\"name\":\"addActiveChainIds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bonder\",\"type\":\"address\"}],\"name\":\"addBonder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ammWrapper\",\"outputs\":[{\"internalType\":\"contractI_L2_AmmWrapper\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"transferNonce\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"bonderFee\",\"type\":\"uint256\"}],\"name\":\"bondWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"transferNonce\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"bonderFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"bondWithdrawalAndDistribute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"}],\"name\":\"commitTransfers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"relayerFee\",\"type\":\"uint256\"}],\"name\":\"distribute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bonder\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"transferId\",\"type\":\"bytes32\"}],\"name\":\"getBondedWithdrawalAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bonder\",\"type\":\"address\"}],\"name\":\"getCredit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bonder\",\"type\":\"address\"}],\"name\":\"getDebitAndAdditionalDebit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"maybeBonder\",\"type\":\"address\"}],\"name\":\"getIsBonder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNextTransferNonce\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bonder\",\"type\":\"address\"}],\"name\":\"getRawDebit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"transferNonce\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"bonderFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"getTransferId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"rootHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"}],\"name\":\"getTransferRoot\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountWithdrawn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"}],\"internalType\":\"structBridge.TransferRoot\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"rootHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"}],\"name\":\"getTransferRootId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hToken\",\"outputs\":[{\"internalType\":\"contractHopBridgeToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transferId\",\"type\":\"bytes32\"}],\"name\":\"isTransferIdSpent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l1BridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l1BridgeCaller\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l1Governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lastCommitTimeForChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxPendingTransfers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messengerProxy\",\"outputs\":[{\"internalType\":\"contractI_L2_PolygonMessengerProxy\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minBonderBps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minBonderFeeAbsolute\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumForceCommitDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pendingAmountForChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pendingTransferIdsForChainId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"chainIds\",\"type\":\"uint256[]\"}],\"name\":\"removeActiveChainIds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bonder\",\"type\":\"address\"}],\"name\":\"removeBonder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"rootHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"originalAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"rescueTransferRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bonderFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractI_L2_AmmWrapper\",\"name\":\"_ammWrapper\",\"type\":\"address\"}],\"name\":\"setAmmWrapper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"setHopBridgeTokenOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1BridgeAddress\",\"type\":\"address\"}],\"name\":\"setL1BridgeAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1BridgeCaller\",\"type\":\"address\"}],\"name\":\"setL1BridgeCaller\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Governance\",\"type\":\"address\"}],\"name\":\"setL1Governance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxPendingTransfers\",\"type\":\"uint256\"}],\"name\":\"setMaxPendingTransfers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractI_L2_PolygonMessengerProxy\",\"name\":\"_messengerProxy\",\"type\":\"address\"}],\"name\":\"setMessengerProxy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minBonderBps\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minBonderFeeAbsolute\",\"type\":\"uint256\"}],\"name\":\"setMinimumBonderFeeRequirements\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minimumForceCommitDelay\",\"type\":\"uint256\"}],\"name\":\"setMinimumForceCommitDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"rootHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"}],\"name\":\"setTransferRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bonder\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"transferId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"rootHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"transferRootTotalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"transferIdTreeIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"siblings\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"totalLeaves\",\"type\":\"uint256\"}],\"name\":\"settleBondedWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bonder\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"transferIds\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"}],\"name\":\"settleBondedWithdrawals\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bonder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transferNonceIncrementer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"transferNonce\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"bonderFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"rootHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"transferRootTotalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"transferIdTreeIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"siblings\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"totalLeaves\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// L2BridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use L2BridgeMetaData.ABI instead.
var L2BridgeABI = L2BridgeMetaData.ABI

// L2Bridge is an auto generated Go binding around an Ethereum contract.
type L2Bridge struct {
	L2BridgeCaller     // Read-only binding to the contract
	L2BridgeTransactor // Write-only binding to the contract
	L2BridgeFilterer   // Log filterer for contract events
}

// L2BridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type L2BridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2BridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L2BridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2BridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L2BridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2BridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L2BridgeSession struct {
	Contract     *L2Bridge         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L2BridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L2BridgeCallerSession struct {
	Contract *L2BridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// L2BridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L2BridgeTransactorSession struct {
	Contract     *L2BridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// L2BridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type L2BridgeRaw struct {
	Contract *L2Bridge // Generic contract binding to access the raw methods on
}

// L2BridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L2BridgeCallerRaw struct {
	Contract *L2BridgeCaller // Generic read-only contract binding to access the raw methods on
}

// L2BridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L2BridgeTransactorRaw struct {
	Contract *L2BridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL2Bridge creates a new instance of L2Bridge, bound to a specific deployed contract.
func NewL2Bridge(address common.Address, backend bind.ContractBackend) (*L2Bridge, error) {
	contract, err := bindL2Bridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L2Bridge{L2BridgeCaller: L2BridgeCaller{contract: contract}, L2BridgeTransactor: L2BridgeTransactor{contract: contract}, L2BridgeFilterer: L2BridgeFilterer{contract: contract}}, nil
}

// NewL2BridgeCaller creates a new read-only instance of L2Bridge, bound to a specific deployed contract.
func NewL2BridgeCaller(address common.Address, caller bind.ContractCaller) (*L2BridgeCaller, error) {
	contract, err := bindL2Bridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L2BridgeCaller{contract: contract}, nil
}

// NewL2BridgeTransactor creates a new write-only instance of L2Bridge, bound to a specific deployed contract.
func NewL2BridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*L2BridgeTransactor, error) {
	contract, err := bindL2Bridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L2BridgeTransactor{contract: contract}, nil
}

// NewL2BridgeFilterer creates a new log filterer instance of L2Bridge, bound to a specific deployed contract.
func NewL2BridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*L2BridgeFilterer, error) {
	contract, err := bindL2Bridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L2BridgeFilterer{contract: contract}, nil
}

// bindL2Bridge binds a generic wrapper to an already deployed contract.
func bindL2Bridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(L2BridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2Bridge *L2BridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2Bridge.Contract.L2BridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2Bridge *L2BridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2Bridge.Contract.L2BridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2Bridge *L2BridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2Bridge.Contract.L2BridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2Bridge *L2BridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2Bridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2Bridge *L2BridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2Bridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2Bridge *L2BridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2Bridge.Contract.contract.Transact(opts, method, params...)
}

// ActiveChainIds is a free data retrieval call binding the contract method 0xc97d172e.
//
// Solidity: function activeChainIds(uint256 ) view returns(bool)
func (_L2Bridge *L2BridgeCaller) ActiveChainIds(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _L2Bridge.contract.Call(opts, &out, "activeChainIds", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ActiveChainIds is a free data retrieval call binding the contract method 0xc97d172e.
//
// Solidity: function activeChainIds(uint256 ) view returns(bool)
func (_L2Bridge *L2BridgeSession) ActiveChainIds(arg0 *big.Int) (bool, error) {
	return _L2Bridge.Contract.ActiveChainIds(&_L2Bridge.CallOpts, arg0)
}

// ActiveChainIds is a free data retrieval call binding the contract method 0xc97d172e.
//
// Solidity: function activeChainIds(uint256 ) view returns(bool)
func (_L2Bridge *L2BridgeCallerSession) ActiveChainIds(arg0 *big.Int) (bool, error) {
	return _L2Bridge.Contract.ActiveChainIds(&_L2Bridge.CallOpts, arg0)
}

// AmmWrapper is a free data retrieval call binding the contract method 0xe9cdfe51.
//
// Solidity: function ammWrapper() view returns(address)
func (_L2Bridge *L2BridgeCaller) AmmWrapper(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2Bridge.contract.Call(opts, &out, "ammWrapper")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AmmWrapper is a free data retrieval call binding the contract method 0xe9cdfe51.
//
// Solidity: function ammWrapper() view returns(address)
func (_L2Bridge *L2BridgeSession) AmmWrapper() (common.Address, error) {
	return _L2Bridge.Contract.AmmWrapper(&_L2Bridge.CallOpts)
}

// AmmWrapper is a free data retrieval call binding the contract method 0xe9cdfe51.
//
// Solidity: function ammWrapper() view returns(address)
func (_L2Bridge *L2BridgeCallerSession) AmmWrapper() (common.Address, error) {
	return _L2Bridge.Contract.AmmWrapper(&_L2Bridge.CallOpts)
}

// GetBondedWithdrawalAmount is a free data retrieval call binding the contract method 0x302830ab.
//
// Solidity: function getBondedWithdrawalAmount(address bonder, bytes32 transferId) view returns(uint256)
func (_L2Bridge *L2BridgeCaller) GetBondedWithdrawalAmount(opts *bind.CallOpts, bonder common.Address, transferId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _L2Bridge.contract.Call(opts, &out, "getBondedWithdrawalAmount", bonder, transferId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBondedWithdrawalAmount is a free data retrieval call binding the contract method 0x302830ab.
//
// Solidity: function getBondedWithdrawalAmount(address bonder, bytes32 transferId) view returns(uint256)
func (_L2Bridge *L2BridgeSession) GetBondedWithdrawalAmount(bonder common.Address, transferId [32]byte) (*big.Int, error) {
	return _L2Bridge.Contract.GetBondedWithdrawalAmount(&_L2Bridge.CallOpts, bonder, transferId)
}

// GetBondedWithdrawalAmount is a free data retrieval call binding the contract method 0x302830ab.
//
// Solidity: function getBondedWithdrawalAmount(address bonder, bytes32 transferId) view returns(uint256)
func (_L2Bridge *L2BridgeCallerSession) GetBondedWithdrawalAmount(bonder common.Address, transferId [32]byte) (*big.Int, error) {
	return _L2Bridge.Contract.GetBondedWithdrawalAmount(&_L2Bridge.CallOpts, bonder, transferId)
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256 chainId)
func (_L2Bridge *L2BridgeCaller) GetChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L2Bridge.contract.Call(opts, &out, "getChainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256 chainId)
func (_L2Bridge *L2BridgeSession) GetChainId() (*big.Int, error) {
	return _L2Bridge.Contract.GetChainId(&_L2Bridge.CallOpts)
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256 chainId)
func (_L2Bridge *L2BridgeCallerSession) GetChainId() (*big.Int, error) {
	return _L2Bridge.Contract.GetChainId(&_L2Bridge.CallOpts)
}

// GetCredit is a free data retrieval call binding the contract method 0x57344e6f.
//
// Solidity: function getCredit(address bonder) view returns(uint256)
func (_L2Bridge *L2BridgeCaller) GetCredit(opts *bind.CallOpts, bonder common.Address) (*big.Int, error) {
	var out []interface{}
	err := _L2Bridge.contract.Call(opts, &out, "getCredit", bonder)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCredit is a free data retrieval call binding the contract method 0x57344e6f.
//
// Solidity: function getCredit(address bonder) view returns(uint256)
func (_L2Bridge *L2BridgeSession) GetCredit(bonder common.Address) (*big.Int, error) {
	return _L2Bridge.Contract.GetCredit(&_L2Bridge.CallOpts, bonder)
}

// GetCredit is a free data retrieval call binding the contract method 0x57344e6f.
//
// Solidity: function getCredit(address bonder) view returns(uint256)
func (_L2Bridge *L2BridgeCallerSession) GetCredit(bonder common.Address) (*big.Int, error) {
	return _L2Bridge.Contract.GetCredit(&_L2Bridge.CallOpts, bonder)
}

// GetDebitAndAdditionalDebit is a free data retrieval call binding the contract method 0xffa9286c.
//
// Solidity: function getDebitAndAdditionalDebit(address bonder) view returns(uint256)
func (_L2Bridge *L2BridgeCaller) GetDebitAndAdditionalDebit(opts *bind.CallOpts, bonder common.Address) (*big.Int, error) {
	var out []interface{}
	err := _L2Bridge.contract.Call(opts, &out, "getDebitAndAdditionalDebit", bonder)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDebitAndAdditionalDebit is a free data retrieval call binding the contract method 0xffa9286c.
//
// Solidity: function getDebitAndAdditionalDebit(address bonder) view returns(uint256)
func (_L2Bridge *L2BridgeSession) GetDebitAndAdditionalDebit(bonder common.Address) (*big.Int, error) {
	return _L2Bridge.Contract.GetDebitAndAdditionalDebit(&_L2Bridge.CallOpts, bonder)
}

// GetDebitAndAdditionalDebit is a free data retrieval call binding the contract method 0xffa9286c.
//
// Solidity: function getDebitAndAdditionalDebit(address bonder) view returns(uint256)
func (_L2Bridge *L2BridgeCallerSession) GetDebitAndAdditionalDebit(bonder common.Address) (*big.Int, error) {
	return _L2Bridge.Contract.GetDebitAndAdditionalDebit(&_L2Bridge.CallOpts, bonder)
}

// GetIsBonder is a free data retrieval call binding the contract method 0xd5ef7551.
//
// Solidity: function getIsBonder(address maybeBonder) view returns(bool)
func (_L2Bridge *L2BridgeCaller) GetIsBonder(opts *bind.CallOpts, maybeBonder common.Address) (bool, error) {
	var out []interface{}
	err := _L2Bridge.contract.Call(opts, &out, "getIsBonder", maybeBonder)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetIsBonder is a free data retrieval call binding the contract method 0xd5ef7551.
//
// Solidity: function getIsBonder(address maybeBonder) view returns(bool)
func (_L2Bridge *L2BridgeSession) GetIsBonder(maybeBonder common.Address) (bool, error) {
	return _L2Bridge.Contract.GetIsBonder(&_L2Bridge.CallOpts, maybeBonder)
}

// GetIsBonder is a free data retrieval call binding the contract method 0xd5ef7551.
//
// Solidity: function getIsBonder(address maybeBonder) view returns(bool)
func (_L2Bridge *L2BridgeCallerSession) GetIsBonder(maybeBonder common.Address) (bool, error) {
	return _L2Bridge.Contract.GetIsBonder(&_L2Bridge.CallOpts, maybeBonder)
}

// GetNextTransferNonce is a free data retrieval call binding the contract method 0x051e7216.
//
// Solidity: function getNextTransferNonce() view returns(bytes32)
func (_L2Bridge *L2BridgeCaller) GetNextTransferNonce(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _L2Bridge.contract.Call(opts, &out, "getNextTransferNonce")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetNextTransferNonce is a free data retrieval call binding the contract method 0x051e7216.
//
// Solidity: function getNextTransferNonce() view returns(bytes32)
func (_L2Bridge *L2BridgeSession) GetNextTransferNonce() ([32]byte, error) {
	return _L2Bridge.Contract.GetNextTransferNonce(&_L2Bridge.CallOpts)
}

// GetNextTransferNonce is a free data retrieval call binding the contract method 0x051e7216.
//
// Solidity: function getNextTransferNonce() view returns(bytes32)
func (_L2Bridge *L2BridgeCallerSession) GetNextTransferNonce() ([32]byte, error) {
	return _L2Bridge.Contract.GetNextTransferNonce(&_L2Bridge.CallOpts)
}

// GetRawDebit is a free data retrieval call binding the contract method 0x13948c76.
//
// Solidity: function getRawDebit(address bonder) view returns(uint256)
func (_L2Bridge *L2BridgeCaller) GetRawDebit(opts *bind.CallOpts, bonder common.Address) (*big.Int, error) {
	var out []interface{}
	err := _L2Bridge.contract.Call(opts, &out, "getRawDebit", bonder)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRawDebit is a free data retrieval call binding the contract method 0x13948c76.
//
// Solidity: function getRawDebit(address bonder) view returns(uint256)
func (_L2Bridge *L2BridgeSession) GetRawDebit(bonder common.Address) (*big.Int, error) {
	return _L2Bridge.Contract.GetRawDebit(&_L2Bridge.CallOpts, bonder)
}

// GetRawDebit is a free data retrieval call binding the contract method 0x13948c76.
//
// Solidity: function getRawDebit(address bonder) view returns(uint256)
func (_L2Bridge *L2BridgeCallerSession) GetRawDebit(bonder common.Address) (*big.Int, error) {
	return _L2Bridge.Contract.GetRawDebit(&_L2Bridge.CallOpts, bonder)
}

// GetTransferId is a free data retrieval call binding the contract method 0xaf215f94.
//
// Solidity: function getTransferId(uint256 chainId, address recipient, uint256 amount, bytes32 transferNonce, uint256 bonderFee, uint256 amountOutMin, uint256 deadline) pure returns(bytes32)
func (_L2Bridge *L2BridgeCaller) GetTransferId(opts *bind.CallOpts, chainId *big.Int, recipient common.Address, amount *big.Int, transferNonce [32]byte, bonderFee *big.Int, amountOutMin *big.Int, deadline *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _L2Bridge.contract.Call(opts, &out, "getTransferId", chainId, recipient, amount, transferNonce, bonderFee, amountOutMin, deadline)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetTransferId is a free data retrieval call binding the contract method 0xaf215f94.
//
// Solidity: function getTransferId(uint256 chainId, address recipient, uint256 amount, bytes32 transferNonce, uint256 bonderFee, uint256 amountOutMin, uint256 deadline) pure returns(bytes32)
func (_L2Bridge *L2BridgeSession) GetTransferId(chainId *big.Int, recipient common.Address, amount *big.Int, transferNonce [32]byte, bonderFee *big.Int, amountOutMin *big.Int, deadline *big.Int) ([32]byte, error) {
	return _L2Bridge.Contract.GetTransferId(&_L2Bridge.CallOpts, chainId, recipient, amount, transferNonce, bonderFee, amountOutMin, deadline)
}

// GetTransferId is a free data retrieval call binding the contract method 0xaf215f94.
//
// Solidity: function getTransferId(uint256 chainId, address recipient, uint256 amount, bytes32 transferNonce, uint256 bonderFee, uint256 amountOutMin, uint256 deadline) pure returns(bytes32)
func (_L2Bridge *L2BridgeCallerSession) GetTransferId(chainId *big.Int, recipient common.Address, amount *big.Int, transferNonce [32]byte, bonderFee *big.Int, amountOutMin *big.Int, deadline *big.Int) ([32]byte, error) {
	return _L2Bridge.Contract.GetTransferId(&_L2Bridge.CallOpts, chainId, recipient, amount, transferNonce, bonderFee, amountOutMin, deadline)
}

// GetTransferRoot is a free data retrieval call binding the contract method 0xce803b4f.
//
// Solidity: function getTransferRoot(bytes32 rootHash, uint256 totalAmount) view returns((uint256,uint256,uint256))
func (_L2Bridge *L2BridgeCaller) GetTransferRoot(opts *bind.CallOpts, rootHash [32]byte, totalAmount *big.Int) (BridgeTransferRoot, error) {
	var out []interface{}
	err := _L2Bridge.contract.Call(opts, &out, "getTransferRoot", rootHash, totalAmount)

	if err != nil {
		return *new(BridgeTransferRoot), err
	}

	out0 := *abi.ConvertType(out[0], new(BridgeTransferRoot)).(*BridgeTransferRoot)

	return out0, err

}

// GetTransferRoot is a free data retrieval call binding the contract method 0xce803b4f.
//
// Solidity: function getTransferRoot(bytes32 rootHash, uint256 totalAmount) view returns((uint256,uint256,uint256))
func (_L2Bridge *L2BridgeSession) GetTransferRoot(rootHash [32]byte, totalAmount *big.Int) (BridgeTransferRoot, error) {
	return _L2Bridge.Contract.GetTransferRoot(&_L2Bridge.CallOpts, rootHash, totalAmount)
}

// GetTransferRoot is a free data retrieval call binding the contract method 0xce803b4f.
//
// Solidity: function getTransferRoot(bytes32 rootHash, uint256 totalAmount) view returns((uint256,uint256,uint256))
func (_L2Bridge *L2BridgeCallerSession) GetTransferRoot(rootHash [32]byte, totalAmount *big.Int) (BridgeTransferRoot, error) {
	return _L2Bridge.Contract.GetTransferRoot(&_L2Bridge.CallOpts, rootHash, totalAmount)
}

// GetTransferRootId is a free data retrieval call binding the contract method 0x960a7afa.
//
// Solidity: function getTransferRootId(bytes32 rootHash, uint256 totalAmount) pure returns(bytes32)
func (_L2Bridge *L2BridgeCaller) GetTransferRootId(opts *bind.CallOpts, rootHash [32]byte, totalAmount *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _L2Bridge.contract.Call(opts, &out, "getTransferRootId", rootHash, totalAmount)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetTransferRootId is a free data retrieval call binding the contract method 0x960a7afa.
//
// Solidity: function getTransferRootId(bytes32 rootHash, uint256 totalAmount) pure returns(bytes32)
func (_L2Bridge *L2BridgeSession) GetTransferRootId(rootHash [32]byte, totalAmount *big.Int) ([32]byte, error) {
	return _L2Bridge.Contract.GetTransferRootId(&_L2Bridge.CallOpts, rootHash, totalAmount)
}

// GetTransferRootId is a free data retrieval call binding the contract method 0x960a7afa.
//
// Solidity: function getTransferRootId(bytes32 rootHash, uint256 totalAmount) pure returns(bytes32)
func (_L2Bridge *L2BridgeCallerSession) GetTransferRootId(rootHash [32]byte, totalAmount *big.Int) ([32]byte, error) {
	return _L2Bridge.Contract.GetTransferRootId(&_L2Bridge.CallOpts, rootHash, totalAmount)
}

// HToken is a free data retrieval call binding the contract method 0xfc6e3b3b.
//
// Solidity: function hToken() view returns(address)
func (_L2Bridge *L2BridgeCaller) HToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2Bridge.contract.Call(opts, &out, "hToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HToken is a free data retrieval call binding the contract method 0xfc6e3b3b.
//
// Solidity: function hToken() view returns(address)
func (_L2Bridge *L2BridgeSession) HToken() (common.Address, error) {
	return _L2Bridge.Contract.HToken(&_L2Bridge.CallOpts)
}

// HToken is a free data retrieval call binding the contract method 0xfc6e3b3b.
//
// Solidity: function hToken() view returns(address)
func (_L2Bridge *L2BridgeCallerSession) HToken() (common.Address, error) {
	return _L2Bridge.Contract.HToken(&_L2Bridge.CallOpts)
}

// IsTransferIdSpent is a free data retrieval call binding the contract method 0x3a7af631.
//
// Solidity: function isTransferIdSpent(bytes32 transferId) view returns(bool)
func (_L2Bridge *L2BridgeCaller) IsTransferIdSpent(opts *bind.CallOpts, transferId [32]byte) (bool, error) {
	var out []interface{}
	err := _L2Bridge.contract.Call(opts, &out, "isTransferIdSpent", transferId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTransferIdSpent is a free data retrieval call binding the contract method 0x3a7af631.
//
// Solidity: function isTransferIdSpent(bytes32 transferId) view returns(bool)
func (_L2Bridge *L2BridgeSession) IsTransferIdSpent(transferId [32]byte) (bool, error) {
	return _L2Bridge.Contract.IsTransferIdSpent(&_L2Bridge.CallOpts, transferId)
}

// IsTransferIdSpent is a free data retrieval call binding the contract method 0x3a7af631.
//
// Solidity: function isTransferIdSpent(bytes32 transferId) view returns(bool)
func (_L2Bridge *L2BridgeCallerSession) IsTransferIdSpent(transferId [32]byte) (bool, error) {
	return _L2Bridge.Contract.IsTransferIdSpent(&_L2Bridge.CallOpts, transferId)
}

// L1BridgeAddress is a free data retrieval call binding the contract method 0x5ab2a558.
//
// Solidity: function l1BridgeAddress() view returns(address)
func (_L2Bridge *L2BridgeCaller) L1BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2Bridge.contract.Call(opts, &out, "l1BridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L1BridgeAddress is a free data retrieval call binding the contract method 0x5ab2a558.
//
// Solidity: function l1BridgeAddress() view returns(address)
func (_L2Bridge *L2BridgeSession) L1BridgeAddress() (common.Address, error) {
	return _L2Bridge.Contract.L1BridgeAddress(&_L2Bridge.CallOpts)
}

// L1BridgeAddress is a free data retrieval call binding the contract method 0x5ab2a558.
//
// Solidity: function l1BridgeAddress() view returns(address)
func (_L2Bridge *L2BridgeCallerSession) L1BridgeAddress() (common.Address, error) {
	return _L2Bridge.Contract.L1BridgeAddress(&_L2Bridge.CallOpts)
}

// L1BridgeCaller is a free data retrieval call binding the contract method 0xd2442783.
//
// Solidity: function l1BridgeCaller() view returns(address)
func (_L2Bridge *L2BridgeCaller) L1BridgeCaller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2Bridge.contract.Call(opts, &out, "l1BridgeCaller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L1BridgeCaller is a free data retrieval call binding the contract method 0xd2442783.
//
// Solidity: function l1BridgeCaller() view returns(address)
func (_L2Bridge *L2BridgeSession) L1BridgeCaller() (common.Address, error) {
	return _L2Bridge.Contract.L1BridgeCaller(&_L2Bridge.CallOpts)
}

// L1BridgeCaller is a free data retrieval call binding the contract method 0xd2442783.
//
// Solidity: function l1BridgeCaller() view returns(address)
func (_L2Bridge *L2BridgeCallerSession) L1BridgeCaller() (common.Address, error) {
	return _L2Bridge.Contract.L1BridgeCaller(&_L2Bridge.CallOpts)
}

// L1Governance is a free data retrieval call binding the contract method 0x3ef23f7f.
//
// Solidity: function l1Governance() view returns(address)
func (_L2Bridge *L2BridgeCaller) L1Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2Bridge.contract.Call(opts, &out, "l1Governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L1Governance is a free data retrieval call binding the contract method 0x3ef23f7f.
//
// Solidity: function l1Governance() view returns(address)
func (_L2Bridge *L2BridgeSession) L1Governance() (common.Address, error) {
	return _L2Bridge.Contract.L1Governance(&_L2Bridge.CallOpts)
}

// L1Governance is a free data retrieval call binding the contract method 0x3ef23f7f.
//
// Solidity: function l1Governance() view returns(address)
func (_L2Bridge *L2BridgeCallerSession) L1Governance() (common.Address, error) {
	return _L2Bridge.Contract.L1Governance(&_L2Bridge.CallOpts)
}

// LastCommitTimeForChainId is a free data retrieval call binding the contract method 0xd4e54c47.
//
// Solidity: function lastCommitTimeForChainId(uint256 ) view returns(uint256)
func (_L2Bridge *L2BridgeCaller) LastCommitTimeForChainId(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _L2Bridge.contract.Call(opts, &out, "lastCommitTimeForChainId", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastCommitTimeForChainId is a free data retrieval call binding the contract method 0xd4e54c47.
//
// Solidity: function lastCommitTimeForChainId(uint256 ) view returns(uint256)
func (_L2Bridge *L2BridgeSession) LastCommitTimeForChainId(arg0 *big.Int) (*big.Int, error) {
	return _L2Bridge.Contract.LastCommitTimeForChainId(&_L2Bridge.CallOpts, arg0)
}

// LastCommitTimeForChainId is a free data retrieval call binding the contract method 0xd4e54c47.
//
// Solidity: function lastCommitTimeForChainId(uint256 ) view returns(uint256)
func (_L2Bridge *L2BridgeCallerSession) LastCommitTimeForChainId(arg0 *big.Int) (*big.Int, error) {
	return _L2Bridge.Contract.LastCommitTimeForChainId(&_L2Bridge.CallOpts, arg0)
}

// MaxPendingTransfers is a free data retrieval call binding the contract method 0xbed93c84.
//
// Solidity: function maxPendingTransfers() view returns(uint256)
func (_L2Bridge *L2BridgeCaller) MaxPendingTransfers(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L2Bridge.contract.Call(opts, &out, "maxPendingTransfers")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxPendingTransfers is a free data retrieval call binding the contract method 0xbed93c84.
//
// Solidity: function maxPendingTransfers() view returns(uint256)
func (_L2Bridge *L2BridgeSession) MaxPendingTransfers() (*big.Int, error) {
	return _L2Bridge.Contract.MaxPendingTransfers(&_L2Bridge.CallOpts)
}

// MaxPendingTransfers is a free data retrieval call binding the contract method 0xbed93c84.
//
// Solidity: function maxPendingTransfers() view returns(uint256)
func (_L2Bridge *L2BridgeCallerSession) MaxPendingTransfers() (*big.Int, error) {
	return _L2Bridge.Contract.MaxPendingTransfers(&_L2Bridge.CallOpts)
}

// MessengerProxy is a free data retrieval call binding the contract method 0xce2d280e.
//
// Solidity: function messengerProxy() view returns(address)
func (_L2Bridge *L2BridgeCaller) MessengerProxy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2Bridge.contract.Call(opts, &out, "messengerProxy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MessengerProxy is a free data retrieval call binding the contract method 0xce2d280e.
//
// Solidity: function messengerProxy() view returns(address)
func (_L2Bridge *L2BridgeSession) MessengerProxy() (common.Address, error) {
	return _L2Bridge.Contract.MessengerProxy(&_L2Bridge.CallOpts)
}

// MessengerProxy is a free data retrieval call binding the contract method 0xce2d280e.
//
// Solidity: function messengerProxy() view returns(address)
func (_L2Bridge *L2BridgeCallerSession) MessengerProxy() (common.Address, error) {
	return _L2Bridge.Contract.MessengerProxy(&_L2Bridge.CallOpts)
}

// MinBonderBps is a free data retrieval call binding the contract method 0x35e2c4af.
//
// Solidity: function minBonderBps() view returns(uint256)
func (_L2Bridge *L2BridgeCaller) MinBonderBps(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L2Bridge.contract.Call(opts, &out, "minBonderBps")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinBonderBps is a free data retrieval call binding the contract method 0x35e2c4af.
//
// Solidity: function minBonderBps() view returns(uint256)
func (_L2Bridge *L2BridgeSession) MinBonderBps() (*big.Int, error) {
	return _L2Bridge.Contract.MinBonderBps(&_L2Bridge.CallOpts)
}

// MinBonderBps is a free data retrieval call binding the contract method 0x35e2c4af.
//
// Solidity: function minBonderBps() view returns(uint256)
func (_L2Bridge *L2BridgeCallerSession) MinBonderBps() (*big.Int, error) {
	return _L2Bridge.Contract.MinBonderBps(&_L2Bridge.CallOpts)
}

// MinBonderFeeAbsolute is a free data retrieval call binding the contract method 0xc3035261.
//
// Solidity: function minBonderFeeAbsolute() view returns(uint256)
func (_L2Bridge *L2BridgeCaller) MinBonderFeeAbsolute(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L2Bridge.contract.Call(opts, &out, "minBonderFeeAbsolute")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinBonderFeeAbsolute is a free data retrieval call binding the contract method 0xc3035261.
//
// Solidity: function minBonderFeeAbsolute() view returns(uint256)
func (_L2Bridge *L2BridgeSession) MinBonderFeeAbsolute() (*big.Int, error) {
	return _L2Bridge.Contract.MinBonderFeeAbsolute(&_L2Bridge.CallOpts)
}

// MinBonderFeeAbsolute is a free data retrieval call binding the contract method 0xc3035261.
//
// Solidity: function minBonderFeeAbsolute() view returns(uint256)
func (_L2Bridge *L2BridgeCallerSession) MinBonderFeeAbsolute() (*big.Int, error) {
	return _L2Bridge.Contract.MinBonderFeeAbsolute(&_L2Bridge.CallOpts)
}

// MinimumForceCommitDelay is a free data retrieval call binding the contract method 0x8f658198.
//
// Solidity: function minimumForceCommitDelay() view returns(uint256)
func (_L2Bridge *L2BridgeCaller) MinimumForceCommitDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L2Bridge.contract.Call(opts, &out, "minimumForceCommitDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumForceCommitDelay is a free data retrieval call binding the contract method 0x8f658198.
//
// Solidity: function minimumForceCommitDelay() view returns(uint256)
func (_L2Bridge *L2BridgeSession) MinimumForceCommitDelay() (*big.Int, error) {
	return _L2Bridge.Contract.MinimumForceCommitDelay(&_L2Bridge.CallOpts)
}

// MinimumForceCommitDelay is a free data retrieval call binding the contract method 0x8f658198.
//
// Solidity: function minimumForceCommitDelay() view returns(uint256)
func (_L2Bridge *L2BridgeCallerSession) MinimumForceCommitDelay() (*big.Int, error) {
	return _L2Bridge.Contract.MinimumForceCommitDelay(&_L2Bridge.CallOpts)
}

// PendingAmountForChainId is a free data retrieval call binding the contract method 0x0f5e09e7.
//
// Solidity: function pendingAmountForChainId(uint256 ) view returns(uint256)
func (_L2Bridge *L2BridgeCaller) PendingAmountForChainId(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _L2Bridge.contract.Call(opts, &out, "pendingAmountForChainId", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingAmountForChainId is a free data retrieval call binding the contract method 0x0f5e09e7.
//
// Solidity: function pendingAmountForChainId(uint256 ) view returns(uint256)
func (_L2Bridge *L2BridgeSession) PendingAmountForChainId(arg0 *big.Int) (*big.Int, error) {
	return _L2Bridge.Contract.PendingAmountForChainId(&_L2Bridge.CallOpts, arg0)
}

// PendingAmountForChainId is a free data retrieval call binding the contract method 0x0f5e09e7.
//
// Solidity: function pendingAmountForChainId(uint256 ) view returns(uint256)
func (_L2Bridge *L2BridgeCallerSession) PendingAmountForChainId(arg0 *big.Int) (*big.Int, error) {
	return _L2Bridge.Contract.PendingAmountForChainId(&_L2Bridge.CallOpts, arg0)
}

// PendingTransferIdsForChainId is a free data retrieval call binding the contract method 0x98445caf.
//
// Solidity: function pendingTransferIdsForChainId(uint256 , uint256 ) view returns(bytes32)
func (_L2Bridge *L2BridgeCaller) PendingTransferIdsForChainId(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _L2Bridge.contract.Call(opts, &out, "pendingTransferIdsForChainId", arg0, arg1)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PendingTransferIdsForChainId is a free data retrieval call binding the contract method 0x98445caf.
//
// Solidity: function pendingTransferIdsForChainId(uint256 , uint256 ) view returns(bytes32)
func (_L2Bridge *L2BridgeSession) PendingTransferIdsForChainId(arg0 *big.Int, arg1 *big.Int) ([32]byte, error) {
	return _L2Bridge.Contract.PendingTransferIdsForChainId(&_L2Bridge.CallOpts, arg0, arg1)
}

// PendingTransferIdsForChainId is a free data retrieval call binding the contract method 0x98445caf.
//
// Solidity: function pendingTransferIdsForChainId(uint256 , uint256 ) view returns(bytes32)
func (_L2Bridge *L2BridgeCallerSession) PendingTransferIdsForChainId(arg0 *big.Int, arg1 *big.Int) ([32]byte, error) {
	return _L2Bridge.Contract.PendingTransferIdsForChainId(&_L2Bridge.CallOpts, arg0, arg1)
}

// TransferNonceIncrementer is a free data retrieval call binding the contract method 0x82c69f9d.
//
// Solidity: function transferNonceIncrementer() view returns(uint256)
func (_L2Bridge *L2BridgeCaller) TransferNonceIncrementer(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L2Bridge.contract.Call(opts, &out, "transferNonceIncrementer")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TransferNonceIncrementer is a free data retrieval call binding the contract method 0x82c69f9d.
//
// Solidity: function transferNonceIncrementer() view returns(uint256)
func (_L2Bridge *L2BridgeSession) TransferNonceIncrementer() (*big.Int, error) {
	return _L2Bridge.Contract.TransferNonceIncrementer(&_L2Bridge.CallOpts)
}

// TransferNonceIncrementer is a free data retrieval call binding the contract method 0x82c69f9d.
//
// Solidity: function transferNonceIncrementer() view returns(uint256)
func (_L2Bridge *L2BridgeCallerSession) TransferNonceIncrementer() (*big.Int, error) {
	return _L2Bridge.Contract.TransferNonceIncrementer(&_L2Bridge.CallOpts)
}

// AddActiveChainIds is a paid mutator transaction binding the contract method 0xf8398fa4.
//
// Solidity: function addActiveChainIds(uint256[] chainIds) returns()
func (_L2Bridge *L2BridgeTransactor) AddActiveChainIds(opts *bind.TransactOpts, chainIds []*big.Int) (*types.Transaction, error) {
	return _L2Bridge.contract.Transact(opts, "addActiveChainIds", chainIds)
}

// AddActiveChainIds is a paid mutator transaction binding the contract method 0xf8398fa4.
//
// Solidity: function addActiveChainIds(uint256[] chainIds) returns()
func (_L2Bridge *L2BridgeSession) AddActiveChainIds(chainIds []*big.Int) (*types.Transaction, error) {
	return _L2Bridge.Contract.AddActiveChainIds(&_L2Bridge.TransactOpts, chainIds)
}

// AddActiveChainIds is a paid mutator transaction binding the contract method 0xf8398fa4.
//
// Solidity: function addActiveChainIds(uint256[] chainIds) returns()
func (_L2Bridge *L2BridgeTransactorSession) AddActiveChainIds(chainIds []*big.Int) (*types.Transaction, error) {
	return _L2Bridge.Contract.AddActiveChainIds(&_L2Bridge.TransactOpts, chainIds)
}

// AddBonder is a paid mutator transaction binding the contract method 0x5325937f.
//
// Solidity: function addBonder(address bonder) returns()
func (_L2Bridge *L2BridgeTransactor) AddBonder(opts *bind.TransactOpts, bonder common.Address) (*types.Transaction, error) {
	return _L2Bridge.contract.Transact(opts, "addBonder", bonder)
}

// AddBonder is a paid mutator transaction binding the contract method 0x5325937f.
//
// Solidity: function addBonder(address bonder) returns()
func (_L2Bridge *L2BridgeSession) AddBonder(bonder common.Address) (*types.Transaction, error) {
	return _L2Bridge.Contract.AddBonder(&_L2Bridge.TransactOpts, bonder)
}

// AddBonder is a paid mutator transaction binding the contract method 0x5325937f.
//
// Solidity: function addBonder(address bonder) returns()
func (_L2Bridge *L2BridgeTransactorSession) AddBonder(bonder common.Address) (*types.Transaction, error) {
	return _L2Bridge.Contract.AddBonder(&_L2Bridge.TransactOpts, bonder)
}

// BondWithdrawal is a paid mutator transaction binding the contract method 0x23c452cd.
//
// Solidity: function bondWithdrawal(address recipient, uint256 amount, bytes32 transferNonce, uint256 bonderFee) returns()
func (_L2Bridge *L2BridgeTransactor) BondWithdrawal(opts *bind.TransactOpts, recipient common.Address, amount *big.Int, transferNonce [32]byte, bonderFee *big.Int) (*types.Transaction, error) {
	return _L2Bridge.contract.Transact(opts, "bondWithdrawal", recipient, amount, transferNonce, bonderFee)
}

// BondWithdrawal is a paid mutator transaction binding the contract method 0x23c452cd.
//
// Solidity: function bondWithdrawal(address recipient, uint256 amount, bytes32 transferNonce, uint256 bonderFee) returns()
func (_L2Bridge *L2BridgeSession) BondWithdrawal(recipient common.Address, amount *big.Int, transferNonce [32]byte, bonderFee *big.Int) (*types.Transaction, error) {
	return _L2Bridge.Contract.BondWithdrawal(&_L2Bridge.TransactOpts, recipient, amount, transferNonce, bonderFee)
}

// BondWithdrawal is a paid mutator transaction binding the contract method 0x23c452cd.
//
// Solidity: function bondWithdrawal(address recipient, uint256 amount, bytes32 transferNonce, uint256 bonderFee) returns()
func (_L2Bridge *L2BridgeTransactorSession) BondWithdrawal(recipient common.Address, amount *big.Int, transferNonce [32]byte, bonderFee *big.Int) (*types.Transaction, error) {
	return _L2Bridge.Contract.BondWithdrawal(&_L2Bridge.TransactOpts, recipient, amount, transferNonce, bonderFee)
}

// BondWithdrawalAndDistribute is a paid mutator transaction binding the contract method 0x3d12a85a.
//
// Solidity: function bondWithdrawalAndDistribute(address recipient, uint256 amount, bytes32 transferNonce, uint256 bonderFee, uint256 amountOutMin, uint256 deadline) returns()
func (_L2Bridge *L2BridgeTransactor) BondWithdrawalAndDistribute(opts *bind.TransactOpts, recipient common.Address, amount *big.Int, transferNonce [32]byte, bonderFee *big.Int, amountOutMin *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _L2Bridge.contract.Transact(opts, "bondWithdrawalAndDistribute", recipient, amount, transferNonce, bonderFee, amountOutMin, deadline)
}

// BondWithdrawalAndDistribute is a paid mutator transaction binding the contract method 0x3d12a85a.
//
// Solidity: function bondWithdrawalAndDistribute(address recipient, uint256 amount, bytes32 transferNonce, uint256 bonderFee, uint256 amountOutMin, uint256 deadline) returns()
func (_L2Bridge *L2BridgeSession) BondWithdrawalAndDistribute(recipient common.Address, amount *big.Int, transferNonce [32]byte, bonderFee *big.Int, amountOutMin *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _L2Bridge.Contract.BondWithdrawalAndDistribute(&_L2Bridge.TransactOpts, recipient, amount, transferNonce, bonderFee, amountOutMin, deadline)
}

// BondWithdrawalAndDistribute is a paid mutator transaction binding the contract method 0x3d12a85a.
//
// Solidity: function bondWithdrawalAndDistribute(address recipient, uint256 amount, bytes32 transferNonce, uint256 bonderFee, uint256 amountOutMin, uint256 deadline) returns()
func (_L2Bridge *L2BridgeTransactorSession) BondWithdrawalAndDistribute(recipient common.Address, amount *big.Int, transferNonce [32]byte, bonderFee *big.Int, amountOutMin *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _L2Bridge.Contract.BondWithdrawalAndDistribute(&_L2Bridge.TransactOpts, recipient, amount, transferNonce, bonderFee, amountOutMin, deadline)
}

// CommitTransfers is a paid mutator transaction binding the contract method 0x32b949a2.
//
// Solidity: function commitTransfers(uint256 destinationChainId) returns()
func (_L2Bridge *L2BridgeTransactor) CommitTransfers(opts *bind.TransactOpts, destinationChainId *big.Int) (*types.Transaction, error) {
	return _L2Bridge.contract.Transact(opts, "commitTransfers", destinationChainId)
}

// CommitTransfers is a paid mutator transaction binding the contract method 0x32b949a2.
//
// Solidity: function commitTransfers(uint256 destinationChainId) returns()
func (_L2Bridge *L2BridgeSession) CommitTransfers(destinationChainId *big.Int) (*types.Transaction, error) {
	return _L2Bridge.Contract.CommitTransfers(&_L2Bridge.TransactOpts, destinationChainId)
}

// CommitTransfers is a paid mutator transaction binding the contract method 0x32b949a2.
//
// Solidity: function commitTransfers(uint256 destinationChainId) returns()
func (_L2Bridge *L2BridgeTransactorSession) CommitTransfers(destinationChainId *big.Int) (*types.Transaction, error) {
	return _L2Bridge.Contract.CommitTransfers(&_L2Bridge.TransactOpts, destinationChainId)
}

// Distribute is a paid mutator transaction binding the contract method 0xcc29a306.
//
// Solidity: function distribute(address recipient, uint256 amount, uint256 amountOutMin, uint256 deadline, address relayer, uint256 relayerFee) returns()
func (_L2Bridge *L2BridgeTransactor) Distribute(opts *bind.TransactOpts, recipient common.Address, amount *big.Int, amountOutMin *big.Int, deadline *big.Int, relayer common.Address, relayerFee *big.Int) (*types.Transaction, error) {
	return _L2Bridge.contract.Transact(opts, "distribute", recipient, amount, amountOutMin, deadline, relayer, relayerFee)
}

// Distribute is a paid mutator transaction binding the contract method 0xcc29a306.
//
// Solidity: function distribute(address recipient, uint256 amount, uint256 amountOutMin, uint256 deadline, address relayer, uint256 relayerFee) returns()
func (_L2Bridge *L2BridgeSession) Distribute(recipient common.Address, amount *big.Int, amountOutMin *big.Int, deadline *big.Int, relayer common.Address, relayerFee *big.Int) (*types.Transaction, error) {
	return _L2Bridge.Contract.Distribute(&_L2Bridge.TransactOpts, recipient, amount, amountOutMin, deadline, relayer, relayerFee)
}

// Distribute is a paid mutator transaction binding the contract method 0xcc29a306.
//
// Solidity: function distribute(address recipient, uint256 amount, uint256 amountOutMin, uint256 deadline, address relayer, uint256 relayerFee) returns()
func (_L2Bridge *L2BridgeTransactorSession) Distribute(recipient common.Address, amount *big.Int, amountOutMin *big.Int, deadline *big.Int, relayer common.Address, relayerFee *big.Int) (*types.Transaction, error) {
	return _L2Bridge.Contract.Distribute(&_L2Bridge.TransactOpts, recipient, amount, amountOutMin, deadline, relayer, relayerFee)
}

// RemoveActiveChainIds is a paid mutator transaction binding the contract method 0x9f600a0b.
//
// Solidity: function removeActiveChainIds(uint256[] chainIds) returns()
func (_L2Bridge *L2BridgeTransactor) RemoveActiveChainIds(opts *bind.TransactOpts, chainIds []*big.Int) (*types.Transaction, error) {
	return _L2Bridge.contract.Transact(opts, "removeActiveChainIds", chainIds)
}

// RemoveActiveChainIds is a paid mutator transaction binding the contract method 0x9f600a0b.
//
// Solidity: function removeActiveChainIds(uint256[] chainIds) returns()
func (_L2Bridge *L2BridgeSession) RemoveActiveChainIds(chainIds []*big.Int) (*types.Transaction, error) {
	return _L2Bridge.Contract.RemoveActiveChainIds(&_L2Bridge.TransactOpts, chainIds)
}

// RemoveActiveChainIds is a paid mutator transaction binding the contract method 0x9f600a0b.
//
// Solidity: function removeActiveChainIds(uint256[] chainIds) returns()
func (_L2Bridge *L2BridgeTransactorSession) RemoveActiveChainIds(chainIds []*big.Int) (*types.Transaction, error) {
	return _L2Bridge.Contract.RemoveActiveChainIds(&_L2Bridge.TransactOpts, chainIds)
}

// RemoveBonder is a paid mutator transaction binding the contract method 0x04e6c2c0.
//
// Solidity: function removeBonder(address bonder) returns()
func (_L2Bridge *L2BridgeTransactor) RemoveBonder(opts *bind.TransactOpts, bonder common.Address) (*types.Transaction, error) {
	return _L2Bridge.contract.Transact(opts, "removeBonder", bonder)
}

// RemoveBonder is a paid mutator transaction binding the contract method 0x04e6c2c0.
//
// Solidity: function removeBonder(address bonder) returns()
func (_L2Bridge *L2BridgeSession) RemoveBonder(bonder common.Address) (*types.Transaction, error) {
	return _L2Bridge.Contract.RemoveBonder(&_L2Bridge.TransactOpts, bonder)
}

// RemoveBonder is a paid mutator transaction binding the contract method 0x04e6c2c0.
//
// Solidity: function removeBonder(address bonder) returns()
func (_L2Bridge *L2BridgeTransactorSession) RemoveBonder(bonder common.Address) (*types.Transaction, error) {
	return _L2Bridge.Contract.RemoveBonder(&_L2Bridge.TransactOpts, bonder)
}

// RescueTransferRoot is a paid mutator transaction binding the contract method 0xcbd1642e.
//
// Solidity: function rescueTransferRoot(bytes32 rootHash, uint256 originalAmount, address recipient) returns()
func (_L2Bridge *L2BridgeTransactor) RescueTransferRoot(opts *bind.TransactOpts, rootHash [32]byte, originalAmount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _L2Bridge.contract.Transact(opts, "rescueTransferRoot", rootHash, originalAmount, recipient)
}

// RescueTransferRoot is a paid mutator transaction binding the contract method 0xcbd1642e.
//
// Solidity: function rescueTransferRoot(bytes32 rootHash, uint256 originalAmount, address recipient) returns()
func (_L2Bridge *L2BridgeSession) RescueTransferRoot(rootHash [32]byte, originalAmount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _L2Bridge.Contract.RescueTransferRoot(&_L2Bridge.TransactOpts, rootHash, originalAmount, recipient)
}

// RescueTransferRoot is a paid mutator transaction binding the contract method 0xcbd1642e.
//
// Solidity: function rescueTransferRoot(bytes32 rootHash, uint256 originalAmount, address recipient) returns()
func (_L2Bridge *L2BridgeTransactorSession) RescueTransferRoot(rootHash [32]byte, originalAmount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _L2Bridge.Contract.RescueTransferRoot(&_L2Bridge.TransactOpts, rootHash, originalAmount, recipient)
}

// Send is a paid mutator transaction binding the contract method 0xa6bd1b33.
//
// Solidity: function send(uint256 chainId, address recipient, uint256 amount, uint256 bonderFee, uint256 amountOutMin, uint256 deadline) returns()
func (_L2Bridge *L2BridgeTransactor) Send(opts *bind.TransactOpts, chainId *big.Int, recipient common.Address, amount *big.Int, bonderFee *big.Int, amountOutMin *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _L2Bridge.contract.Transact(opts, "send", chainId, recipient, amount, bonderFee, amountOutMin, deadline)
}

// Send is a paid mutator transaction binding the contract method 0xa6bd1b33.
//
// Solidity: function send(uint256 chainId, address recipient, uint256 amount, uint256 bonderFee, uint256 amountOutMin, uint256 deadline) returns()
func (_L2Bridge *L2BridgeSession) Send(chainId *big.Int, recipient common.Address, amount *big.Int, bonderFee *big.Int, amountOutMin *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _L2Bridge.Contract.Send(&_L2Bridge.TransactOpts, chainId, recipient, amount, bonderFee, amountOutMin, deadline)
}

// Send is a paid mutator transaction binding the contract method 0xa6bd1b33.
//
// Solidity: function send(uint256 chainId, address recipient, uint256 amount, uint256 bonderFee, uint256 amountOutMin, uint256 deadline) returns()
func (_L2Bridge *L2BridgeTransactorSession) Send(chainId *big.Int, recipient common.Address, amount *big.Int, bonderFee *big.Int, amountOutMin *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _L2Bridge.Contract.Send(&_L2Bridge.TransactOpts, chainId, recipient, amount, bonderFee, amountOutMin, deadline)
}

// SetAmmWrapper is a paid mutator transaction binding the contract method 0x64c6fdb4.
//
// Solidity: function setAmmWrapper(address _ammWrapper) returns()
func (_L2Bridge *L2BridgeTransactor) SetAmmWrapper(opts *bind.TransactOpts, _ammWrapper common.Address) (*types.Transaction, error) {
	return _L2Bridge.contract.Transact(opts, "setAmmWrapper", _ammWrapper)
}

// SetAmmWrapper is a paid mutator transaction binding the contract method 0x64c6fdb4.
//
// Solidity: function setAmmWrapper(address _ammWrapper) returns()
func (_L2Bridge *L2BridgeSession) SetAmmWrapper(_ammWrapper common.Address) (*types.Transaction, error) {
	return _L2Bridge.Contract.SetAmmWrapper(&_L2Bridge.TransactOpts, _ammWrapper)
}

// SetAmmWrapper is a paid mutator transaction binding the contract method 0x64c6fdb4.
//
// Solidity: function setAmmWrapper(address _ammWrapper) returns()
func (_L2Bridge *L2BridgeTransactorSession) SetAmmWrapper(_ammWrapper common.Address) (*types.Transaction, error) {
	return _L2Bridge.Contract.SetAmmWrapper(&_L2Bridge.TransactOpts, _ammWrapper)
}

// SetHopBridgeTokenOwner is a paid mutator transaction binding the contract method 0x8295f258.
//
// Solidity: function setHopBridgeTokenOwner(address newOwner) returns()
func (_L2Bridge *L2BridgeTransactor) SetHopBridgeTokenOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _L2Bridge.contract.Transact(opts, "setHopBridgeTokenOwner", newOwner)
}

// SetHopBridgeTokenOwner is a paid mutator transaction binding the contract method 0x8295f258.
//
// Solidity: function setHopBridgeTokenOwner(address newOwner) returns()
func (_L2Bridge *L2BridgeSession) SetHopBridgeTokenOwner(newOwner common.Address) (*types.Transaction, error) {
	return _L2Bridge.Contract.SetHopBridgeTokenOwner(&_L2Bridge.TransactOpts, newOwner)
}

// SetHopBridgeTokenOwner is a paid mutator transaction binding the contract method 0x8295f258.
//
// Solidity: function setHopBridgeTokenOwner(address newOwner) returns()
func (_L2Bridge *L2BridgeTransactorSession) SetHopBridgeTokenOwner(newOwner common.Address) (*types.Transaction, error) {
	return _L2Bridge.Contract.SetHopBridgeTokenOwner(&_L2Bridge.TransactOpts, newOwner)
}

// SetL1BridgeAddress is a paid mutator transaction binding the contract method 0xe1825d06.
//
// Solidity: function setL1BridgeAddress(address _l1BridgeAddress) returns()
func (_L2Bridge *L2BridgeTransactor) SetL1BridgeAddress(opts *bind.TransactOpts, _l1BridgeAddress common.Address) (*types.Transaction, error) {
	return _L2Bridge.contract.Transact(opts, "setL1BridgeAddress", _l1BridgeAddress)
}

// SetL1BridgeAddress is a paid mutator transaction binding the contract method 0xe1825d06.
//
// Solidity: function setL1BridgeAddress(address _l1BridgeAddress) returns()
func (_L2Bridge *L2BridgeSession) SetL1BridgeAddress(_l1BridgeAddress common.Address) (*types.Transaction, error) {
	return _L2Bridge.Contract.SetL1BridgeAddress(&_L2Bridge.TransactOpts, _l1BridgeAddress)
}

// SetL1BridgeAddress is a paid mutator transaction binding the contract method 0xe1825d06.
//
// Solidity: function setL1BridgeAddress(address _l1BridgeAddress) returns()
func (_L2Bridge *L2BridgeTransactorSession) SetL1BridgeAddress(_l1BridgeAddress common.Address) (*types.Transaction, error) {
	return _L2Bridge.Contract.SetL1BridgeAddress(&_L2Bridge.TransactOpts, _l1BridgeAddress)
}

// SetL1BridgeCaller is a paid mutator transaction binding the contract method 0xaf33ae69.
//
// Solidity: function setL1BridgeCaller(address _l1BridgeCaller) returns()
func (_L2Bridge *L2BridgeTransactor) SetL1BridgeCaller(opts *bind.TransactOpts, _l1BridgeCaller common.Address) (*types.Transaction, error) {
	return _L2Bridge.contract.Transact(opts, "setL1BridgeCaller", _l1BridgeCaller)
}

// SetL1BridgeCaller is a paid mutator transaction binding the contract method 0xaf33ae69.
//
// Solidity: function setL1BridgeCaller(address _l1BridgeCaller) returns()
func (_L2Bridge *L2BridgeSession) SetL1BridgeCaller(_l1BridgeCaller common.Address) (*types.Transaction, error) {
	return _L2Bridge.Contract.SetL1BridgeCaller(&_L2Bridge.TransactOpts, _l1BridgeCaller)
}

// SetL1BridgeCaller is a paid mutator transaction binding the contract method 0xaf33ae69.
//
// Solidity: function setL1BridgeCaller(address _l1BridgeCaller) returns()
func (_L2Bridge *L2BridgeTransactorSession) SetL1BridgeCaller(_l1BridgeCaller common.Address) (*types.Transaction, error) {
	return _L2Bridge.Contract.SetL1BridgeCaller(&_L2Bridge.TransactOpts, _l1BridgeCaller)
}

// SetL1Governance is a paid mutator transaction binding the contract method 0xe40272d7.
//
// Solidity: function setL1Governance(address _l1Governance) returns()
func (_L2Bridge *L2BridgeTransactor) SetL1Governance(opts *bind.TransactOpts, _l1Governance common.Address) (*types.Transaction, error) {
	return _L2Bridge.contract.Transact(opts, "setL1Governance", _l1Governance)
}

// SetL1Governance is a paid mutator transaction binding the contract method 0xe40272d7.
//
// Solidity: function setL1Governance(address _l1Governance) returns()
func (_L2Bridge *L2BridgeSession) SetL1Governance(_l1Governance common.Address) (*types.Transaction, error) {
	return _L2Bridge.Contract.SetL1Governance(&_L2Bridge.TransactOpts, _l1Governance)
}

// SetL1Governance is a paid mutator transaction binding the contract method 0xe40272d7.
//
// Solidity: function setL1Governance(address _l1Governance) returns()
func (_L2Bridge *L2BridgeTransactorSession) SetL1Governance(_l1Governance common.Address) (*types.Transaction, error) {
	return _L2Bridge.Contract.SetL1Governance(&_L2Bridge.TransactOpts, _l1Governance)
}

// SetMaxPendingTransfers is a paid mutator transaction binding the contract method 0x4742bbfb.
//
// Solidity: function setMaxPendingTransfers(uint256 _maxPendingTransfers) returns()
func (_L2Bridge *L2BridgeTransactor) SetMaxPendingTransfers(opts *bind.TransactOpts, _maxPendingTransfers *big.Int) (*types.Transaction, error) {
	return _L2Bridge.contract.Transact(opts, "setMaxPendingTransfers", _maxPendingTransfers)
}

// SetMaxPendingTransfers is a paid mutator transaction binding the contract method 0x4742bbfb.
//
// Solidity: function setMaxPendingTransfers(uint256 _maxPendingTransfers) returns()
func (_L2Bridge *L2BridgeSession) SetMaxPendingTransfers(_maxPendingTransfers *big.Int) (*types.Transaction, error) {
	return _L2Bridge.Contract.SetMaxPendingTransfers(&_L2Bridge.TransactOpts, _maxPendingTransfers)
}

// SetMaxPendingTransfers is a paid mutator transaction binding the contract method 0x4742bbfb.
//
// Solidity: function setMaxPendingTransfers(uint256 _maxPendingTransfers) returns()
func (_L2Bridge *L2BridgeTransactorSession) SetMaxPendingTransfers(_maxPendingTransfers *big.Int) (*types.Transaction, error) {
	return _L2Bridge.Contract.SetMaxPendingTransfers(&_L2Bridge.TransactOpts, _maxPendingTransfers)
}

// SetMessengerProxy is a paid mutator transaction binding the contract method 0x3b1c54fa.
//
// Solidity: function setMessengerProxy(address _messengerProxy) returns()
func (_L2Bridge *L2BridgeTransactor) SetMessengerProxy(opts *bind.TransactOpts, _messengerProxy common.Address) (*types.Transaction, error) {
	return _L2Bridge.contract.Transact(opts, "setMessengerProxy", _messengerProxy)
}

// SetMessengerProxy is a paid mutator transaction binding the contract method 0x3b1c54fa.
//
// Solidity: function setMessengerProxy(address _messengerProxy) returns()
func (_L2Bridge *L2BridgeSession) SetMessengerProxy(_messengerProxy common.Address) (*types.Transaction, error) {
	return _L2Bridge.Contract.SetMessengerProxy(&_L2Bridge.TransactOpts, _messengerProxy)
}

// SetMessengerProxy is a paid mutator transaction binding the contract method 0x3b1c54fa.
//
// Solidity: function setMessengerProxy(address _messengerProxy) returns()
func (_L2Bridge *L2BridgeTransactorSession) SetMessengerProxy(_messengerProxy common.Address) (*types.Transaction, error) {
	return _L2Bridge.Contract.SetMessengerProxy(&_L2Bridge.TransactOpts, _messengerProxy)
}

// SetMinimumBonderFeeRequirements is a paid mutator transaction binding the contract method 0xa9fa4ed5.
//
// Solidity: function setMinimumBonderFeeRequirements(uint256 _minBonderBps, uint256 _minBonderFeeAbsolute) returns()
func (_L2Bridge *L2BridgeTransactor) SetMinimumBonderFeeRequirements(opts *bind.TransactOpts, _minBonderBps *big.Int, _minBonderFeeAbsolute *big.Int) (*types.Transaction, error) {
	return _L2Bridge.contract.Transact(opts, "setMinimumBonderFeeRequirements", _minBonderBps, _minBonderFeeAbsolute)
}

// SetMinimumBonderFeeRequirements is a paid mutator transaction binding the contract method 0xa9fa4ed5.
//
// Solidity: function setMinimumBonderFeeRequirements(uint256 _minBonderBps, uint256 _minBonderFeeAbsolute) returns()
func (_L2Bridge *L2BridgeSession) SetMinimumBonderFeeRequirements(_minBonderBps *big.Int, _minBonderFeeAbsolute *big.Int) (*types.Transaction, error) {
	return _L2Bridge.Contract.SetMinimumBonderFeeRequirements(&_L2Bridge.TransactOpts, _minBonderBps, _minBonderFeeAbsolute)
}

// SetMinimumBonderFeeRequirements is a paid mutator transaction binding the contract method 0xa9fa4ed5.
//
// Solidity: function setMinimumBonderFeeRequirements(uint256 _minBonderBps, uint256 _minBonderFeeAbsolute) returns()
func (_L2Bridge *L2BridgeTransactorSession) SetMinimumBonderFeeRequirements(_minBonderBps *big.Int, _minBonderFeeAbsolute *big.Int) (*types.Transaction, error) {
	return _L2Bridge.Contract.SetMinimumBonderFeeRequirements(&_L2Bridge.TransactOpts, _minBonderBps, _minBonderFeeAbsolute)
}

// SetMinimumForceCommitDelay is a paid mutator transaction binding the contract method 0x9bf43028.
//
// Solidity: function setMinimumForceCommitDelay(uint256 _minimumForceCommitDelay) returns()
func (_L2Bridge *L2BridgeTransactor) SetMinimumForceCommitDelay(opts *bind.TransactOpts, _minimumForceCommitDelay *big.Int) (*types.Transaction, error) {
	return _L2Bridge.contract.Transact(opts, "setMinimumForceCommitDelay", _minimumForceCommitDelay)
}

// SetMinimumForceCommitDelay is a paid mutator transaction binding the contract method 0x9bf43028.
//
// Solidity: function setMinimumForceCommitDelay(uint256 _minimumForceCommitDelay) returns()
func (_L2Bridge *L2BridgeSession) SetMinimumForceCommitDelay(_minimumForceCommitDelay *big.Int) (*types.Transaction, error) {
	return _L2Bridge.Contract.SetMinimumForceCommitDelay(&_L2Bridge.TransactOpts, _minimumForceCommitDelay)
}

// SetMinimumForceCommitDelay is a paid mutator transaction binding the contract method 0x9bf43028.
//
// Solidity: function setMinimumForceCommitDelay(uint256 _minimumForceCommitDelay) returns()
func (_L2Bridge *L2BridgeTransactorSession) SetMinimumForceCommitDelay(_minimumForceCommitDelay *big.Int) (*types.Transaction, error) {
	return _L2Bridge.Contract.SetMinimumForceCommitDelay(&_L2Bridge.TransactOpts, _minimumForceCommitDelay)
}

// SetTransferRoot is a paid mutator transaction binding the contract method 0xfd31c5ba.
//
// Solidity: function setTransferRoot(bytes32 rootHash, uint256 totalAmount) returns()
func (_L2Bridge *L2BridgeTransactor) SetTransferRoot(opts *bind.TransactOpts, rootHash [32]byte, totalAmount *big.Int) (*types.Transaction, error) {
	return _L2Bridge.contract.Transact(opts, "setTransferRoot", rootHash, totalAmount)
}

// SetTransferRoot is a paid mutator transaction binding the contract method 0xfd31c5ba.
//
// Solidity: function setTransferRoot(bytes32 rootHash, uint256 totalAmount) returns()
func (_L2Bridge *L2BridgeSession) SetTransferRoot(rootHash [32]byte, totalAmount *big.Int) (*types.Transaction, error) {
	return _L2Bridge.Contract.SetTransferRoot(&_L2Bridge.TransactOpts, rootHash, totalAmount)
}

// SetTransferRoot is a paid mutator transaction binding the contract method 0xfd31c5ba.
//
// Solidity: function setTransferRoot(bytes32 rootHash, uint256 totalAmount) returns()
func (_L2Bridge *L2BridgeTransactorSession) SetTransferRoot(rootHash [32]byte, totalAmount *big.Int) (*types.Transaction, error) {
	return _L2Bridge.Contract.SetTransferRoot(&_L2Bridge.TransactOpts, rootHash, totalAmount)
}

// SettleBondedWithdrawal is a paid mutator transaction binding the contract method 0xc7525dd3.
//
// Solidity: function settleBondedWithdrawal(address bonder, bytes32 transferId, bytes32 rootHash, uint256 transferRootTotalAmount, uint256 transferIdTreeIndex, bytes32[] siblings, uint256 totalLeaves) returns()
func (_L2Bridge *L2BridgeTransactor) SettleBondedWithdrawal(opts *bind.TransactOpts, bonder common.Address, transferId [32]byte, rootHash [32]byte, transferRootTotalAmount *big.Int, transferIdTreeIndex *big.Int, siblings [][32]byte, totalLeaves *big.Int) (*types.Transaction, error) {
	return _L2Bridge.contract.Transact(opts, "settleBondedWithdrawal", bonder, transferId, rootHash, transferRootTotalAmount, transferIdTreeIndex, siblings, totalLeaves)
}

// SettleBondedWithdrawal is a paid mutator transaction binding the contract method 0xc7525dd3.
//
// Solidity: function settleBondedWithdrawal(address bonder, bytes32 transferId, bytes32 rootHash, uint256 transferRootTotalAmount, uint256 transferIdTreeIndex, bytes32[] siblings, uint256 totalLeaves) returns()
func (_L2Bridge *L2BridgeSession) SettleBondedWithdrawal(bonder common.Address, transferId [32]byte, rootHash [32]byte, transferRootTotalAmount *big.Int, transferIdTreeIndex *big.Int, siblings [][32]byte, totalLeaves *big.Int) (*types.Transaction, error) {
	return _L2Bridge.Contract.SettleBondedWithdrawal(&_L2Bridge.TransactOpts, bonder, transferId, rootHash, transferRootTotalAmount, transferIdTreeIndex, siblings, totalLeaves)
}

// SettleBondedWithdrawal is a paid mutator transaction binding the contract method 0xc7525dd3.
//
// Solidity: function settleBondedWithdrawal(address bonder, bytes32 transferId, bytes32 rootHash, uint256 transferRootTotalAmount, uint256 transferIdTreeIndex, bytes32[] siblings, uint256 totalLeaves) returns()
func (_L2Bridge *L2BridgeTransactorSession) SettleBondedWithdrawal(bonder common.Address, transferId [32]byte, rootHash [32]byte, transferRootTotalAmount *big.Int, transferIdTreeIndex *big.Int, siblings [][32]byte, totalLeaves *big.Int) (*types.Transaction, error) {
	return _L2Bridge.Contract.SettleBondedWithdrawal(&_L2Bridge.TransactOpts, bonder, transferId, rootHash, transferRootTotalAmount, transferIdTreeIndex, siblings, totalLeaves)
}

// SettleBondedWithdrawals is a paid mutator transaction binding the contract method 0xb162717e.
//
// Solidity: function settleBondedWithdrawals(address bonder, bytes32[] transferIds, uint256 totalAmount) returns()
func (_L2Bridge *L2BridgeTransactor) SettleBondedWithdrawals(opts *bind.TransactOpts, bonder common.Address, transferIds [][32]byte, totalAmount *big.Int) (*types.Transaction, error) {
	return _L2Bridge.contract.Transact(opts, "settleBondedWithdrawals", bonder, transferIds, totalAmount)
}

// SettleBondedWithdrawals is a paid mutator transaction binding the contract method 0xb162717e.
//
// Solidity: function settleBondedWithdrawals(address bonder, bytes32[] transferIds, uint256 totalAmount) returns()
func (_L2Bridge *L2BridgeSession) SettleBondedWithdrawals(bonder common.Address, transferIds [][32]byte, totalAmount *big.Int) (*types.Transaction, error) {
	return _L2Bridge.Contract.SettleBondedWithdrawals(&_L2Bridge.TransactOpts, bonder, transferIds, totalAmount)
}

// SettleBondedWithdrawals is a paid mutator transaction binding the contract method 0xb162717e.
//
// Solidity: function settleBondedWithdrawals(address bonder, bytes32[] transferIds, uint256 totalAmount) returns()
func (_L2Bridge *L2BridgeTransactorSession) SettleBondedWithdrawals(bonder common.Address, transferIds [][32]byte, totalAmount *big.Int) (*types.Transaction, error) {
	return _L2Bridge.Contract.SettleBondedWithdrawals(&_L2Bridge.TransactOpts, bonder, transferIds, totalAmount)
}

// Stake is a paid mutator transaction binding the contract method 0xadc9772e.
//
// Solidity: function stake(address bonder, uint256 amount) payable returns()
func (_L2Bridge *L2BridgeTransactor) Stake(opts *bind.TransactOpts, bonder common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L2Bridge.contract.Transact(opts, "stake", bonder, amount)
}

// Stake is a paid mutator transaction binding the contract method 0xadc9772e.
//
// Solidity: function stake(address bonder, uint256 amount) payable returns()
func (_L2Bridge *L2BridgeSession) Stake(bonder common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L2Bridge.Contract.Stake(&_L2Bridge.TransactOpts, bonder, amount)
}

// Stake is a paid mutator transaction binding the contract method 0xadc9772e.
//
// Solidity: function stake(address bonder, uint256 amount) payable returns()
func (_L2Bridge *L2BridgeTransactorSession) Stake(bonder common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L2Bridge.Contract.Stake(&_L2Bridge.TransactOpts, bonder, amount)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 amount) returns()
func (_L2Bridge *L2BridgeTransactor) Unstake(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _L2Bridge.contract.Transact(opts, "unstake", amount)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 amount) returns()
func (_L2Bridge *L2BridgeSession) Unstake(amount *big.Int) (*types.Transaction, error) {
	return _L2Bridge.Contract.Unstake(&_L2Bridge.TransactOpts, amount)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 amount) returns()
func (_L2Bridge *L2BridgeTransactorSession) Unstake(amount *big.Int) (*types.Transaction, error) {
	return _L2Bridge.Contract.Unstake(&_L2Bridge.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x0f7aadb7.
//
// Solidity: function withdraw(address recipient, uint256 amount, bytes32 transferNonce, uint256 bonderFee, uint256 amountOutMin, uint256 deadline, bytes32 rootHash, uint256 transferRootTotalAmount, uint256 transferIdTreeIndex, bytes32[] siblings, uint256 totalLeaves) returns()
func (_L2Bridge *L2BridgeTransactor) Withdraw(opts *bind.TransactOpts, recipient common.Address, amount *big.Int, transferNonce [32]byte, bonderFee *big.Int, amountOutMin *big.Int, deadline *big.Int, rootHash [32]byte, transferRootTotalAmount *big.Int, transferIdTreeIndex *big.Int, siblings [][32]byte, totalLeaves *big.Int) (*types.Transaction, error) {
	return _L2Bridge.contract.Transact(opts, "withdraw", recipient, amount, transferNonce, bonderFee, amountOutMin, deadline, rootHash, transferRootTotalAmount, transferIdTreeIndex, siblings, totalLeaves)
}

// Withdraw is a paid mutator transaction binding the contract method 0x0f7aadb7.
//
// Solidity: function withdraw(address recipient, uint256 amount, bytes32 transferNonce, uint256 bonderFee, uint256 amountOutMin, uint256 deadline, bytes32 rootHash, uint256 transferRootTotalAmount, uint256 transferIdTreeIndex, bytes32[] siblings, uint256 totalLeaves) returns()
func (_L2Bridge *L2BridgeSession) Withdraw(recipient common.Address, amount *big.Int, transferNonce [32]byte, bonderFee *big.Int, amountOutMin *big.Int, deadline *big.Int, rootHash [32]byte, transferRootTotalAmount *big.Int, transferIdTreeIndex *big.Int, siblings [][32]byte, totalLeaves *big.Int) (*types.Transaction, error) {
	return _L2Bridge.Contract.Withdraw(&_L2Bridge.TransactOpts, recipient, amount, transferNonce, bonderFee, amountOutMin, deadline, rootHash, transferRootTotalAmount, transferIdTreeIndex, siblings, totalLeaves)
}

// Withdraw is a paid mutator transaction binding the contract method 0x0f7aadb7.
//
// Solidity: function withdraw(address recipient, uint256 amount, bytes32 transferNonce, uint256 bonderFee, uint256 amountOutMin, uint256 deadline, bytes32 rootHash, uint256 transferRootTotalAmount, uint256 transferIdTreeIndex, bytes32[] siblings, uint256 totalLeaves) returns()
func (_L2Bridge *L2BridgeTransactorSession) Withdraw(recipient common.Address, amount *big.Int, transferNonce [32]byte, bonderFee *big.Int, amountOutMin *big.Int, deadline *big.Int, rootHash [32]byte, transferRootTotalAmount *big.Int, transferIdTreeIndex *big.Int, siblings [][32]byte, totalLeaves *big.Int) (*types.Transaction, error) {
	return _L2Bridge.Contract.Withdraw(&_L2Bridge.TransactOpts, recipient, amount, transferNonce, bonderFee, amountOutMin, deadline, rootHash, transferRootTotalAmount, transferIdTreeIndex, siblings, totalLeaves)
}

// L2BridgeBonderAddedIterator is returned from FilterBonderAdded and is used to iterate over the raw logs and unpacked data for BonderAdded events raised by the L2Bridge contract.
type L2BridgeBonderAddedIterator struct {
	Event *L2BridgeBonderAdded // Event containing the contract specifics and raw log

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
func (it *L2BridgeBonderAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2BridgeBonderAdded)
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
		it.Event = new(L2BridgeBonderAdded)
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
func (it *L2BridgeBonderAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2BridgeBonderAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2BridgeBonderAdded represents a BonderAdded event raised by the L2Bridge contract.
type L2BridgeBonderAdded struct {
	NewBonder common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBonderAdded is a free log retrieval operation binding the contract event 0x2cec73b7434d3b91198ad1a618f63e6a0761ce281af5ec9ec76606d948d03e23.
//
// Solidity: event BonderAdded(address indexed newBonder)
func (_L2Bridge *L2BridgeFilterer) FilterBonderAdded(opts *bind.FilterOpts, newBonder []common.Address) (*L2BridgeBonderAddedIterator, error) {

	var newBonderRule []interface{}
	for _, newBonderItem := range newBonder {
		newBonderRule = append(newBonderRule, newBonderItem)
	}

	logs, sub, err := _L2Bridge.contract.FilterLogs(opts, "BonderAdded", newBonderRule)
	if err != nil {
		return nil, err
	}
	return &L2BridgeBonderAddedIterator{contract: _L2Bridge.contract, event: "BonderAdded", logs: logs, sub: sub}, nil
}

// WatchBonderAdded is a free log subscription operation binding the contract event 0x2cec73b7434d3b91198ad1a618f63e6a0761ce281af5ec9ec76606d948d03e23.
//
// Solidity: event BonderAdded(address indexed newBonder)
func (_L2Bridge *L2BridgeFilterer) WatchBonderAdded(opts *bind.WatchOpts, sink chan<- *L2BridgeBonderAdded, newBonder []common.Address) (event.Subscription, error) {

	var newBonderRule []interface{}
	for _, newBonderItem := range newBonder {
		newBonderRule = append(newBonderRule, newBonderItem)
	}

	logs, sub, err := _L2Bridge.contract.WatchLogs(opts, "BonderAdded", newBonderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2BridgeBonderAdded)
				if err := _L2Bridge.contract.UnpackLog(event, "BonderAdded", log); err != nil {
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

// ParseBonderAdded is a log parse operation binding the contract event 0x2cec73b7434d3b91198ad1a618f63e6a0761ce281af5ec9ec76606d948d03e23.
//
// Solidity: event BonderAdded(address indexed newBonder)
func (_L2Bridge *L2BridgeFilterer) ParseBonderAdded(log types.Log) (*L2BridgeBonderAdded, error) {
	event := new(L2BridgeBonderAdded)
	if err := _L2Bridge.contract.UnpackLog(event, "BonderAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2BridgeBonderRemovedIterator is returned from FilterBonderRemoved and is used to iterate over the raw logs and unpacked data for BonderRemoved events raised by the L2Bridge contract.
type L2BridgeBonderRemovedIterator struct {
	Event *L2BridgeBonderRemoved // Event containing the contract specifics and raw log

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
func (it *L2BridgeBonderRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2BridgeBonderRemoved)
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
		it.Event = new(L2BridgeBonderRemoved)
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
func (it *L2BridgeBonderRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2BridgeBonderRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2BridgeBonderRemoved represents a BonderRemoved event raised by the L2Bridge contract.
type L2BridgeBonderRemoved struct {
	PreviousBonder common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBonderRemoved is a free log retrieval operation binding the contract event 0x4234ba611d325b3ba434c4e1b037967b955b1274d4185ee9847b7491111a48ff.
//
// Solidity: event BonderRemoved(address indexed previousBonder)
func (_L2Bridge *L2BridgeFilterer) FilterBonderRemoved(opts *bind.FilterOpts, previousBonder []common.Address) (*L2BridgeBonderRemovedIterator, error) {

	var previousBonderRule []interface{}
	for _, previousBonderItem := range previousBonder {
		previousBonderRule = append(previousBonderRule, previousBonderItem)
	}

	logs, sub, err := _L2Bridge.contract.FilterLogs(opts, "BonderRemoved", previousBonderRule)
	if err != nil {
		return nil, err
	}
	return &L2BridgeBonderRemovedIterator{contract: _L2Bridge.contract, event: "BonderRemoved", logs: logs, sub: sub}, nil
}

// WatchBonderRemoved is a free log subscription operation binding the contract event 0x4234ba611d325b3ba434c4e1b037967b955b1274d4185ee9847b7491111a48ff.
//
// Solidity: event BonderRemoved(address indexed previousBonder)
func (_L2Bridge *L2BridgeFilterer) WatchBonderRemoved(opts *bind.WatchOpts, sink chan<- *L2BridgeBonderRemoved, previousBonder []common.Address) (event.Subscription, error) {

	var previousBonderRule []interface{}
	for _, previousBonderItem := range previousBonder {
		previousBonderRule = append(previousBonderRule, previousBonderItem)
	}

	logs, sub, err := _L2Bridge.contract.WatchLogs(opts, "BonderRemoved", previousBonderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2BridgeBonderRemoved)
				if err := _L2Bridge.contract.UnpackLog(event, "BonderRemoved", log); err != nil {
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

// ParseBonderRemoved is a log parse operation binding the contract event 0x4234ba611d325b3ba434c4e1b037967b955b1274d4185ee9847b7491111a48ff.
//
// Solidity: event BonderRemoved(address indexed previousBonder)
func (_L2Bridge *L2BridgeFilterer) ParseBonderRemoved(log types.Log) (*L2BridgeBonderRemoved, error) {
	event := new(L2BridgeBonderRemoved)
	if err := _L2Bridge.contract.UnpackLog(event, "BonderRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2BridgeL1BridgeMessageIterator is returned from FilterL1BridgeMessage and is used to iterate over the raw logs and unpacked data for L1BridgeMessage events raised by the L2Bridge contract.
type L2BridgeL1BridgeMessageIterator struct {
	Event *L2BridgeL1BridgeMessage // Event containing the contract specifics and raw log

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
func (it *L2BridgeL1BridgeMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2BridgeL1BridgeMessage)
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
		it.Event = new(L2BridgeL1BridgeMessage)
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
func (it *L2BridgeL1BridgeMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2BridgeL1BridgeMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2BridgeL1BridgeMessage represents a L1BridgeMessage event raised by the L2Bridge contract.
type L2BridgeL1BridgeMessage struct {
	Data []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterL1BridgeMessage is a free log retrieval operation binding the contract event 0x1e5d8c05ae2f395d8a109cd90453147c849e4f85ba279e69d331d59f31356bcf.
//
// Solidity: event L1_BridgeMessage(bytes data)
func (_L2Bridge *L2BridgeFilterer) FilterL1BridgeMessage(opts *bind.FilterOpts) (*L2BridgeL1BridgeMessageIterator, error) {

	logs, sub, err := _L2Bridge.contract.FilterLogs(opts, "L1_BridgeMessage")
	if err != nil {
		return nil, err
	}
	return &L2BridgeL1BridgeMessageIterator{contract: _L2Bridge.contract, event: "L1_BridgeMessage", logs: logs, sub: sub}, nil
}

// WatchL1BridgeMessage is a free log subscription operation binding the contract event 0x1e5d8c05ae2f395d8a109cd90453147c849e4f85ba279e69d331d59f31356bcf.
//
// Solidity: event L1_BridgeMessage(bytes data)
func (_L2Bridge *L2BridgeFilterer) WatchL1BridgeMessage(opts *bind.WatchOpts, sink chan<- *L2BridgeL1BridgeMessage) (event.Subscription, error) {

	logs, sub, err := _L2Bridge.contract.WatchLogs(opts, "L1_BridgeMessage")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2BridgeL1BridgeMessage)
				if err := _L2Bridge.contract.UnpackLog(event, "L1_BridgeMessage", log); err != nil {
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

// ParseL1BridgeMessage is a log parse operation binding the contract event 0x1e5d8c05ae2f395d8a109cd90453147c849e4f85ba279e69d331d59f31356bcf.
//
// Solidity: event L1_BridgeMessage(bytes data)
func (_L2Bridge *L2BridgeFilterer) ParseL1BridgeMessage(log types.Log) (*L2BridgeL1BridgeMessage, error) {
	event := new(L2BridgeL1BridgeMessage)
	if err := _L2Bridge.contract.UnpackLog(event, "L1_BridgeMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2BridgeMultipleWithdrawalsSettledIterator is returned from FilterMultipleWithdrawalsSettled and is used to iterate over the raw logs and unpacked data for MultipleWithdrawalsSettled events raised by the L2Bridge contract.
type L2BridgeMultipleWithdrawalsSettledIterator struct {
	Event *L2BridgeMultipleWithdrawalsSettled // Event containing the contract specifics and raw log

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
func (it *L2BridgeMultipleWithdrawalsSettledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2BridgeMultipleWithdrawalsSettled)
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
		it.Event = new(L2BridgeMultipleWithdrawalsSettled)
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
func (it *L2BridgeMultipleWithdrawalsSettledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2BridgeMultipleWithdrawalsSettledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2BridgeMultipleWithdrawalsSettled represents a MultipleWithdrawalsSettled event raised by the L2Bridge contract.
type L2BridgeMultipleWithdrawalsSettled struct {
	Bonder            common.Address
	RootHash          [32]byte
	TotalBondsSettled *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterMultipleWithdrawalsSettled is a free log retrieval operation binding the contract event 0x78e830d08be9d5f957414c84d685c061ecbd8467be98b42ebb64f0118b57d2ff.
//
// Solidity: event MultipleWithdrawalsSettled(address indexed bonder, bytes32 indexed rootHash, uint256 totalBondsSettled)
func (_L2Bridge *L2BridgeFilterer) FilterMultipleWithdrawalsSettled(opts *bind.FilterOpts, bonder []common.Address, rootHash [][32]byte) (*L2BridgeMultipleWithdrawalsSettledIterator, error) {

	var bonderRule []interface{}
	for _, bonderItem := range bonder {
		bonderRule = append(bonderRule, bonderItem)
	}
	var rootHashRule []interface{}
	for _, rootHashItem := range rootHash {
		rootHashRule = append(rootHashRule, rootHashItem)
	}

	logs, sub, err := _L2Bridge.contract.FilterLogs(opts, "MultipleWithdrawalsSettled", bonderRule, rootHashRule)
	if err != nil {
		return nil, err
	}
	return &L2BridgeMultipleWithdrawalsSettledIterator{contract: _L2Bridge.contract, event: "MultipleWithdrawalsSettled", logs: logs, sub: sub}, nil
}

// WatchMultipleWithdrawalsSettled is a free log subscription operation binding the contract event 0x78e830d08be9d5f957414c84d685c061ecbd8467be98b42ebb64f0118b57d2ff.
//
// Solidity: event MultipleWithdrawalsSettled(address indexed bonder, bytes32 indexed rootHash, uint256 totalBondsSettled)
func (_L2Bridge *L2BridgeFilterer) WatchMultipleWithdrawalsSettled(opts *bind.WatchOpts, sink chan<- *L2BridgeMultipleWithdrawalsSettled, bonder []common.Address, rootHash [][32]byte) (event.Subscription, error) {

	var bonderRule []interface{}
	for _, bonderItem := range bonder {
		bonderRule = append(bonderRule, bonderItem)
	}
	var rootHashRule []interface{}
	for _, rootHashItem := range rootHash {
		rootHashRule = append(rootHashRule, rootHashItem)
	}

	logs, sub, err := _L2Bridge.contract.WatchLogs(opts, "MultipleWithdrawalsSettled", bonderRule, rootHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2BridgeMultipleWithdrawalsSettled)
				if err := _L2Bridge.contract.UnpackLog(event, "MultipleWithdrawalsSettled", log); err != nil {
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

// ParseMultipleWithdrawalsSettled is a log parse operation binding the contract event 0x78e830d08be9d5f957414c84d685c061ecbd8467be98b42ebb64f0118b57d2ff.
//
// Solidity: event MultipleWithdrawalsSettled(address indexed bonder, bytes32 indexed rootHash, uint256 totalBondsSettled)
func (_L2Bridge *L2BridgeFilterer) ParseMultipleWithdrawalsSettled(log types.Log) (*L2BridgeMultipleWithdrawalsSettled, error) {
	event := new(L2BridgeMultipleWithdrawalsSettled)
	if err := _L2Bridge.contract.UnpackLog(event, "MultipleWithdrawalsSettled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2BridgeStakeIterator is returned from FilterStake and is used to iterate over the raw logs and unpacked data for Stake events raised by the L2Bridge contract.
type L2BridgeStakeIterator struct {
	Event *L2BridgeStake // Event containing the contract specifics and raw log

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
func (it *L2BridgeStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2BridgeStake)
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
		it.Event = new(L2BridgeStake)
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
func (it *L2BridgeStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2BridgeStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2BridgeStake represents a Stake event raised by the L2Bridge contract.
type L2BridgeStake struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterStake is a free log retrieval operation binding the contract event 0xebedb8b3c678666e7f36970bc8f57abf6d8fa2e828c0da91ea5b75bf68ed101a.
//
// Solidity: event Stake(address indexed account, uint256 amount)
func (_L2Bridge *L2BridgeFilterer) FilterStake(opts *bind.FilterOpts, account []common.Address) (*L2BridgeStakeIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _L2Bridge.contract.FilterLogs(opts, "Stake", accountRule)
	if err != nil {
		return nil, err
	}
	return &L2BridgeStakeIterator{contract: _L2Bridge.contract, event: "Stake", logs: logs, sub: sub}, nil
}

// WatchStake is a free log subscription operation binding the contract event 0xebedb8b3c678666e7f36970bc8f57abf6d8fa2e828c0da91ea5b75bf68ed101a.
//
// Solidity: event Stake(address indexed account, uint256 amount)
func (_L2Bridge *L2BridgeFilterer) WatchStake(opts *bind.WatchOpts, sink chan<- *L2BridgeStake, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _L2Bridge.contract.WatchLogs(opts, "Stake", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2BridgeStake)
				if err := _L2Bridge.contract.UnpackLog(event, "Stake", log); err != nil {
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

// ParseStake is a log parse operation binding the contract event 0xebedb8b3c678666e7f36970bc8f57abf6d8fa2e828c0da91ea5b75bf68ed101a.
//
// Solidity: event Stake(address indexed account, uint256 amount)
func (_L2Bridge *L2BridgeFilterer) ParseStake(log types.Log) (*L2BridgeStake, error) {
	event := new(L2BridgeStake)
	if err := _L2Bridge.contract.UnpackLog(event, "Stake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2BridgeTransferFromL1CompletedIterator is returned from FilterTransferFromL1Completed and is used to iterate over the raw logs and unpacked data for TransferFromL1Completed events raised by the L2Bridge contract.
type L2BridgeTransferFromL1CompletedIterator struct {
	Event *L2BridgeTransferFromL1Completed // Event containing the contract specifics and raw log

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
func (it *L2BridgeTransferFromL1CompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2BridgeTransferFromL1Completed)
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
		it.Event = new(L2BridgeTransferFromL1Completed)
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
func (it *L2BridgeTransferFromL1CompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2BridgeTransferFromL1CompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2BridgeTransferFromL1Completed represents a TransferFromL1Completed event raised by the L2Bridge contract.
type L2BridgeTransferFromL1Completed struct {
	Recipient    common.Address
	Amount       *big.Int
	AmountOutMin *big.Int
	Deadline     *big.Int
	Relayer      common.Address
	RelayerFee   *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTransferFromL1Completed is a free log retrieval operation binding the contract event 0x320958176930804eb66c2343c7343fc0367dc16249590c0f195783bee199d094.
//
// Solidity: event TransferFromL1Completed(address indexed recipient, uint256 amount, uint256 amountOutMin, uint256 deadline, address indexed relayer, uint256 relayerFee)
func (_L2Bridge *L2BridgeFilterer) FilterTransferFromL1Completed(opts *bind.FilterOpts, recipient []common.Address, relayer []common.Address) (*L2BridgeTransferFromL1CompletedIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}

	logs, sub, err := _L2Bridge.contract.FilterLogs(opts, "TransferFromL1Completed", recipientRule, relayerRule)
	if err != nil {
		return nil, err
	}
	return &L2BridgeTransferFromL1CompletedIterator{contract: _L2Bridge.contract, event: "TransferFromL1Completed", logs: logs, sub: sub}, nil
}

// WatchTransferFromL1Completed is a free log subscription operation binding the contract event 0x320958176930804eb66c2343c7343fc0367dc16249590c0f195783bee199d094.
//
// Solidity: event TransferFromL1Completed(address indexed recipient, uint256 amount, uint256 amountOutMin, uint256 deadline, address indexed relayer, uint256 relayerFee)
func (_L2Bridge *L2BridgeFilterer) WatchTransferFromL1Completed(opts *bind.WatchOpts, sink chan<- *L2BridgeTransferFromL1Completed, recipient []common.Address, relayer []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}

	logs, sub, err := _L2Bridge.contract.WatchLogs(opts, "TransferFromL1Completed", recipientRule, relayerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2BridgeTransferFromL1Completed)
				if err := _L2Bridge.contract.UnpackLog(event, "TransferFromL1Completed", log); err != nil {
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

// ParseTransferFromL1Completed is a log parse operation binding the contract event 0x320958176930804eb66c2343c7343fc0367dc16249590c0f195783bee199d094.
//
// Solidity: event TransferFromL1Completed(address indexed recipient, uint256 amount, uint256 amountOutMin, uint256 deadline, address indexed relayer, uint256 relayerFee)
func (_L2Bridge *L2BridgeFilterer) ParseTransferFromL1Completed(log types.Log) (*L2BridgeTransferFromL1Completed, error) {
	event := new(L2BridgeTransferFromL1Completed)
	if err := _L2Bridge.contract.UnpackLog(event, "TransferFromL1Completed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2BridgeTransferRootSetIterator is returned from FilterTransferRootSet and is used to iterate over the raw logs and unpacked data for TransferRootSet events raised by the L2Bridge contract.
type L2BridgeTransferRootSetIterator struct {
	Event *L2BridgeTransferRootSet // Event containing the contract specifics and raw log

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
func (it *L2BridgeTransferRootSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2BridgeTransferRootSet)
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
		it.Event = new(L2BridgeTransferRootSet)
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
func (it *L2BridgeTransferRootSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2BridgeTransferRootSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2BridgeTransferRootSet represents a TransferRootSet event raised by the L2Bridge contract.
type L2BridgeTransferRootSet struct {
	RootHash    [32]byte
	TotalAmount *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTransferRootSet is a free log retrieval operation binding the contract event 0xb33d2162aead99dab59e77a7a67ea025b776bf8ca8079e132afdf9b23e03bd42.
//
// Solidity: event TransferRootSet(bytes32 indexed rootHash, uint256 totalAmount)
func (_L2Bridge *L2BridgeFilterer) FilterTransferRootSet(opts *bind.FilterOpts, rootHash [][32]byte) (*L2BridgeTransferRootSetIterator, error) {

	var rootHashRule []interface{}
	for _, rootHashItem := range rootHash {
		rootHashRule = append(rootHashRule, rootHashItem)
	}

	logs, sub, err := _L2Bridge.contract.FilterLogs(opts, "TransferRootSet", rootHashRule)
	if err != nil {
		return nil, err
	}
	return &L2BridgeTransferRootSetIterator{contract: _L2Bridge.contract, event: "TransferRootSet", logs: logs, sub: sub}, nil
}

// WatchTransferRootSet is a free log subscription operation binding the contract event 0xb33d2162aead99dab59e77a7a67ea025b776bf8ca8079e132afdf9b23e03bd42.
//
// Solidity: event TransferRootSet(bytes32 indexed rootHash, uint256 totalAmount)
func (_L2Bridge *L2BridgeFilterer) WatchTransferRootSet(opts *bind.WatchOpts, sink chan<- *L2BridgeTransferRootSet, rootHash [][32]byte) (event.Subscription, error) {

	var rootHashRule []interface{}
	for _, rootHashItem := range rootHash {
		rootHashRule = append(rootHashRule, rootHashItem)
	}

	logs, sub, err := _L2Bridge.contract.WatchLogs(opts, "TransferRootSet", rootHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2BridgeTransferRootSet)
				if err := _L2Bridge.contract.UnpackLog(event, "TransferRootSet", log); err != nil {
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

// ParseTransferRootSet is a log parse operation binding the contract event 0xb33d2162aead99dab59e77a7a67ea025b776bf8ca8079e132afdf9b23e03bd42.
//
// Solidity: event TransferRootSet(bytes32 indexed rootHash, uint256 totalAmount)
func (_L2Bridge *L2BridgeFilterer) ParseTransferRootSet(log types.Log) (*L2BridgeTransferRootSet, error) {
	event := new(L2BridgeTransferRootSet)
	if err := _L2Bridge.contract.UnpackLog(event, "TransferRootSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2BridgeTransferSentIterator is returned from FilterTransferSent and is used to iterate over the raw logs and unpacked data for TransferSent events raised by the L2Bridge contract.
type L2BridgeTransferSentIterator struct {
	Event *L2BridgeTransferSent // Event containing the contract specifics and raw log

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
func (it *L2BridgeTransferSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2BridgeTransferSent)
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
		it.Event = new(L2BridgeTransferSent)
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
func (it *L2BridgeTransferSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2BridgeTransferSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2BridgeTransferSent represents a TransferSent event raised by the L2Bridge contract.
type L2BridgeTransferSent struct {
	TransferId    [32]byte
	ChainId       *big.Int
	Recipient     common.Address
	Amount        *big.Int
	TransferNonce [32]byte
	BonderFee     *big.Int
	Index         *big.Int
	AmountOutMin  *big.Int
	Deadline      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTransferSent is a free log retrieval operation binding the contract event 0xe35dddd4ea75d7e9b3fe93af4f4e40e778c3da4074c9d93e7c6536f1e803c1eb.
//
// Solidity: event TransferSent(bytes32 indexed transferId, uint256 indexed chainId, address indexed recipient, uint256 amount, bytes32 transferNonce, uint256 bonderFee, uint256 index, uint256 amountOutMin, uint256 deadline)
func (_L2Bridge *L2BridgeFilterer) FilterTransferSent(opts *bind.FilterOpts, transferId [][32]byte, chainId []*big.Int, recipient []common.Address) (*L2BridgeTransferSentIterator, error) {

	var transferIdRule []interface{}
	for _, transferIdItem := range transferId {
		transferIdRule = append(transferIdRule, transferIdItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _L2Bridge.contract.FilterLogs(opts, "TransferSent", transferIdRule, chainIdRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &L2BridgeTransferSentIterator{contract: _L2Bridge.contract, event: "TransferSent", logs: logs, sub: sub}, nil
}

// WatchTransferSent is a free log subscription operation binding the contract event 0xe35dddd4ea75d7e9b3fe93af4f4e40e778c3da4074c9d93e7c6536f1e803c1eb.
//
// Solidity: event TransferSent(bytes32 indexed transferId, uint256 indexed chainId, address indexed recipient, uint256 amount, bytes32 transferNonce, uint256 bonderFee, uint256 index, uint256 amountOutMin, uint256 deadline)
func (_L2Bridge *L2BridgeFilterer) WatchTransferSent(opts *bind.WatchOpts, sink chan<- *L2BridgeTransferSent, transferId [][32]byte, chainId []*big.Int, recipient []common.Address) (event.Subscription, error) {

	var transferIdRule []interface{}
	for _, transferIdItem := range transferId {
		transferIdRule = append(transferIdRule, transferIdItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _L2Bridge.contract.WatchLogs(opts, "TransferSent", transferIdRule, chainIdRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2BridgeTransferSent)
				if err := _L2Bridge.contract.UnpackLog(event, "TransferSent", log); err != nil {
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

// ParseTransferSent is a log parse operation binding the contract event 0xe35dddd4ea75d7e9b3fe93af4f4e40e778c3da4074c9d93e7c6536f1e803c1eb.
//
// Solidity: event TransferSent(bytes32 indexed transferId, uint256 indexed chainId, address indexed recipient, uint256 amount, bytes32 transferNonce, uint256 bonderFee, uint256 index, uint256 amountOutMin, uint256 deadline)
func (_L2Bridge *L2BridgeFilterer) ParseTransferSent(log types.Log) (*L2BridgeTransferSent, error) {
	event := new(L2BridgeTransferSent)
	if err := _L2Bridge.contract.UnpackLog(event, "TransferSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2BridgeTransfersCommittedIterator is returned from FilterTransfersCommitted and is used to iterate over the raw logs and unpacked data for TransfersCommitted events raised by the L2Bridge contract.
type L2BridgeTransfersCommittedIterator struct {
	Event *L2BridgeTransfersCommitted // Event containing the contract specifics and raw log

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
func (it *L2BridgeTransfersCommittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2BridgeTransfersCommitted)
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
		it.Event = new(L2BridgeTransfersCommitted)
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
func (it *L2BridgeTransfersCommittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2BridgeTransfersCommittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2BridgeTransfersCommitted represents a TransfersCommitted event raised by the L2Bridge contract.
type L2BridgeTransfersCommitted struct {
	DestinationChainId *big.Int
	RootHash           [32]byte
	TotalAmount        *big.Int
	RootCommittedAt    *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterTransfersCommitted is a free log retrieval operation binding the contract event 0xf52ad20d3b4f50d1c40901dfb95a9ce5270b2fc32694e5c668354721cd87aa74.
//
// Solidity: event TransfersCommitted(uint256 indexed destinationChainId, bytes32 indexed rootHash, uint256 totalAmount, uint256 rootCommittedAt)
func (_L2Bridge *L2BridgeFilterer) FilterTransfersCommitted(opts *bind.FilterOpts, destinationChainId []*big.Int, rootHash [][32]byte) (*L2BridgeTransfersCommittedIterator, error) {

	var destinationChainIdRule []interface{}
	for _, destinationChainIdItem := range destinationChainId {
		destinationChainIdRule = append(destinationChainIdRule, destinationChainIdItem)
	}
	var rootHashRule []interface{}
	for _, rootHashItem := range rootHash {
		rootHashRule = append(rootHashRule, rootHashItem)
	}

	logs, sub, err := _L2Bridge.contract.FilterLogs(opts, "TransfersCommitted", destinationChainIdRule, rootHashRule)
	if err != nil {
		return nil, err
	}
	return &L2BridgeTransfersCommittedIterator{contract: _L2Bridge.contract, event: "TransfersCommitted", logs: logs, sub: sub}, nil
}

// WatchTransfersCommitted is a free log subscription operation binding the contract event 0xf52ad20d3b4f50d1c40901dfb95a9ce5270b2fc32694e5c668354721cd87aa74.
//
// Solidity: event TransfersCommitted(uint256 indexed destinationChainId, bytes32 indexed rootHash, uint256 totalAmount, uint256 rootCommittedAt)
func (_L2Bridge *L2BridgeFilterer) WatchTransfersCommitted(opts *bind.WatchOpts, sink chan<- *L2BridgeTransfersCommitted, destinationChainId []*big.Int, rootHash [][32]byte) (event.Subscription, error) {

	var destinationChainIdRule []interface{}
	for _, destinationChainIdItem := range destinationChainId {
		destinationChainIdRule = append(destinationChainIdRule, destinationChainIdItem)
	}
	var rootHashRule []interface{}
	for _, rootHashItem := range rootHash {
		rootHashRule = append(rootHashRule, rootHashItem)
	}

	logs, sub, err := _L2Bridge.contract.WatchLogs(opts, "TransfersCommitted", destinationChainIdRule, rootHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2BridgeTransfersCommitted)
				if err := _L2Bridge.contract.UnpackLog(event, "TransfersCommitted", log); err != nil {
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

// ParseTransfersCommitted is a log parse operation binding the contract event 0xf52ad20d3b4f50d1c40901dfb95a9ce5270b2fc32694e5c668354721cd87aa74.
//
// Solidity: event TransfersCommitted(uint256 indexed destinationChainId, bytes32 indexed rootHash, uint256 totalAmount, uint256 rootCommittedAt)
func (_L2Bridge *L2BridgeFilterer) ParseTransfersCommitted(log types.Log) (*L2BridgeTransfersCommitted, error) {
	event := new(L2BridgeTransfersCommitted)
	if err := _L2Bridge.contract.UnpackLog(event, "TransfersCommitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2BridgeUnstakeIterator is returned from FilterUnstake and is used to iterate over the raw logs and unpacked data for Unstake events raised by the L2Bridge contract.
type L2BridgeUnstakeIterator struct {
	Event *L2BridgeUnstake // Event containing the contract specifics and raw log

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
func (it *L2BridgeUnstakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2BridgeUnstake)
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
		it.Event = new(L2BridgeUnstake)
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
func (it *L2BridgeUnstakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2BridgeUnstakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2BridgeUnstake represents a Unstake event raised by the L2Bridge contract.
type L2BridgeUnstake struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnstake is a free log retrieval operation binding the contract event 0x85082129d87b2fe11527cb1b3b7a520aeb5aa6913f88a3d8757fe40d1db02fdd.
//
// Solidity: event Unstake(address indexed account, uint256 amount)
func (_L2Bridge *L2BridgeFilterer) FilterUnstake(opts *bind.FilterOpts, account []common.Address) (*L2BridgeUnstakeIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _L2Bridge.contract.FilterLogs(opts, "Unstake", accountRule)
	if err != nil {
		return nil, err
	}
	return &L2BridgeUnstakeIterator{contract: _L2Bridge.contract, event: "Unstake", logs: logs, sub: sub}, nil
}

// WatchUnstake is a free log subscription operation binding the contract event 0x85082129d87b2fe11527cb1b3b7a520aeb5aa6913f88a3d8757fe40d1db02fdd.
//
// Solidity: event Unstake(address indexed account, uint256 amount)
func (_L2Bridge *L2BridgeFilterer) WatchUnstake(opts *bind.WatchOpts, sink chan<- *L2BridgeUnstake, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _L2Bridge.contract.WatchLogs(opts, "Unstake", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2BridgeUnstake)
				if err := _L2Bridge.contract.UnpackLog(event, "Unstake", log); err != nil {
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

// ParseUnstake is a log parse operation binding the contract event 0x85082129d87b2fe11527cb1b3b7a520aeb5aa6913f88a3d8757fe40d1db02fdd.
//
// Solidity: event Unstake(address indexed account, uint256 amount)
func (_L2Bridge *L2BridgeFilterer) ParseUnstake(log types.Log) (*L2BridgeUnstake, error) {
	event := new(L2BridgeUnstake)
	if err := _L2Bridge.contract.UnpackLog(event, "Unstake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2BridgeWithdrawalBondSettledIterator is returned from FilterWithdrawalBondSettled and is used to iterate over the raw logs and unpacked data for WithdrawalBondSettled events raised by the L2Bridge contract.
type L2BridgeWithdrawalBondSettledIterator struct {
	Event *L2BridgeWithdrawalBondSettled // Event containing the contract specifics and raw log

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
func (it *L2BridgeWithdrawalBondSettledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2BridgeWithdrawalBondSettled)
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
		it.Event = new(L2BridgeWithdrawalBondSettled)
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
func (it *L2BridgeWithdrawalBondSettledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2BridgeWithdrawalBondSettledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2BridgeWithdrawalBondSettled represents a WithdrawalBondSettled event raised by the L2Bridge contract.
type L2BridgeWithdrawalBondSettled struct {
	Bonder     common.Address
	TransferId [32]byte
	RootHash   [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalBondSettled is a free log retrieval operation binding the contract event 0x84eb21b24c31b27a3bc67dde4a598aad06db6e9415cd66544492b9616996143c.
//
// Solidity: event WithdrawalBondSettled(address indexed bonder, bytes32 indexed transferId, bytes32 indexed rootHash)
func (_L2Bridge *L2BridgeFilterer) FilterWithdrawalBondSettled(opts *bind.FilterOpts, bonder []common.Address, transferId [][32]byte, rootHash [][32]byte) (*L2BridgeWithdrawalBondSettledIterator, error) {

	var bonderRule []interface{}
	for _, bonderItem := range bonder {
		bonderRule = append(bonderRule, bonderItem)
	}
	var transferIdRule []interface{}
	for _, transferIdItem := range transferId {
		transferIdRule = append(transferIdRule, transferIdItem)
	}
	var rootHashRule []interface{}
	for _, rootHashItem := range rootHash {
		rootHashRule = append(rootHashRule, rootHashItem)
	}

	logs, sub, err := _L2Bridge.contract.FilterLogs(opts, "WithdrawalBondSettled", bonderRule, transferIdRule, rootHashRule)
	if err != nil {
		return nil, err
	}
	return &L2BridgeWithdrawalBondSettledIterator{contract: _L2Bridge.contract, event: "WithdrawalBondSettled", logs: logs, sub: sub}, nil
}

// WatchWithdrawalBondSettled is a free log subscription operation binding the contract event 0x84eb21b24c31b27a3bc67dde4a598aad06db6e9415cd66544492b9616996143c.
//
// Solidity: event WithdrawalBondSettled(address indexed bonder, bytes32 indexed transferId, bytes32 indexed rootHash)
func (_L2Bridge *L2BridgeFilterer) WatchWithdrawalBondSettled(opts *bind.WatchOpts, sink chan<- *L2BridgeWithdrawalBondSettled, bonder []common.Address, transferId [][32]byte, rootHash [][32]byte) (event.Subscription, error) {

	var bonderRule []interface{}
	for _, bonderItem := range bonder {
		bonderRule = append(bonderRule, bonderItem)
	}
	var transferIdRule []interface{}
	for _, transferIdItem := range transferId {
		transferIdRule = append(transferIdRule, transferIdItem)
	}
	var rootHashRule []interface{}
	for _, rootHashItem := range rootHash {
		rootHashRule = append(rootHashRule, rootHashItem)
	}

	logs, sub, err := _L2Bridge.contract.WatchLogs(opts, "WithdrawalBondSettled", bonderRule, transferIdRule, rootHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2BridgeWithdrawalBondSettled)
				if err := _L2Bridge.contract.UnpackLog(event, "WithdrawalBondSettled", log); err != nil {
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

// ParseWithdrawalBondSettled is a log parse operation binding the contract event 0x84eb21b24c31b27a3bc67dde4a598aad06db6e9415cd66544492b9616996143c.
//
// Solidity: event WithdrawalBondSettled(address indexed bonder, bytes32 indexed transferId, bytes32 indexed rootHash)
func (_L2Bridge *L2BridgeFilterer) ParseWithdrawalBondSettled(log types.Log) (*L2BridgeWithdrawalBondSettled, error) {
	event := new(L2BridgeWithdrawalBondSettled)
	if err := _L2Bridge.contract.UnpackLog(event, "WithdrawalBondSettled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2BridgeWithdrawalBondedIterator is returned from FilterWithdrawalBonded and is used to iterate over the raw logs and unpacked data for WithdrawalBonded events raised by the L2Bridge contract.
type L2BridgeWithdrawalBondedIterator struct {
	Event *L2BridgeWithdrawalBonded // Event containing the contract specifics and raw log

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
func (it *L2BridgeWithdrawalBondedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2BridgeWithdrawalBonded)
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
		it.Event = new(L2BridgeWithdrawalBonded)
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
func (it *L2BridgeWithdrawalBondedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2BridgeWithdrawalBondedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2BridgeWithdrawalBonded represents a WithdrawalBonded event raised by the L2Bridge contract.
type L2BridgeWithdrawalBonded struct {
	TransferId [32]byte
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalBonded is a free log retrieval operation binding the contract event 0x0c3d250c7831051e78aa6a56679e590374c7c424415ffe4aa474491def2fe705.
//
// Solidity: event WithdrawalBonded(bytes32 indexed transferId, uint256 amount)
func (_L2Bridge *L2BridgeFilterer) FilterWithdrawalBonded(opts *bind.FilterOpts, transferId [][32]byte) (*L2BridgeWithdrawalBondedIterator, error) {

	var transferIdRule []interface{}
	for _, transferIdItem := range transferId {
		transferIdRule = append(transferIdRule, transferIdItem)
	}

	logs, sub, err := _L2Bridge.contract.FilterLogs(opts, "WithdrawalBonded", transferIdRule)
	if err != nil {
		return nil, err
	}
	return &L2BridgeWithdrawalBondedIterator{contract: _L2Bridge.contract, event: "WithdrawalBonded", logs: logs, sub: sub}, nil
}

// WatchWithdrawalBonded is a free log subscription operation binding the contract event 0x0c3d250c7831051e78aa6a56679e590374c7c424415ffe4aa474491def2fe705.
//
// Solidity: event WithdrawalBonded(bytes32 indexed transferId, uint256 amount)
func (_L2Bridge *L2BridgeFilterer) WatchWithdrawalBonded(opts *bind.WatchOpts, sink chan<- *L2BridgeWithdrawalBonded, transferId [][32]byte) (event.Subscription, error) {

	var transferIdRule []interface{}
	for _, transferIdItem := range transferId {
		transferIdRule = append(transferIdRule, transferIdItem)
	}

	logs, sub, err := _L2Bridge.contract.WatchLogs(opts, "WithdrawalBonded", transferIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2BridgeWithdrawalBonded)
				if err := _L2Bridge.contract.UnpackLog(event, "WithdrawalBonded", log); err != nil {
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

// ParseWithdrawalBonded is a log parse operation binding the contract event 0x0c3d250c7831051e78aa6a56679e590374c7c424415ffe4aa474491def2fe705.
//
// Solidity: event WithdrawalBonded(bytes32 indexed transferId, uint256 amount)
func (_L2Bridge *L2BridgeFilterer) ParseWithdrawalBonded(log types.Log) (*L2BridgeWithdrawalBonded, error) {
	event := new(L2BridgeWithdrawalBonded)
	if err := _L2Bridge.contract.UnpackLog(event, "WithdrawalBonded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2BridgeWithdrewIterator is returned from FilterWithdrew and is used to iterate over the raw logs and unpacked data for Withdrew events raised by the L2Bridge contract.
type L2BridgeWithdrewIterator struct {
	Event *L2BridgeWithdrew // Event containing the contract specifics and raw log

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
func (it *L2BridgeWithdrewIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2BridgeWithdrew)
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
		it.Event = new(L2BridgeWithdrew)
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
func (it *L2BridgeWithdrewIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2BridgeWithdrewIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2BridgeWithdrew represents a Withdrew event raised by the L2Bridge contract.
type L2BridgeWithdrew struct {
	TransferId    [32]byte
	Recipient     common.Address
	Amount        *big.Int
	TransferNonce [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterWithdrew is a free log retrieval operation binding the contract event 0x9475cdbde5fc71fe2ccd413c82878ee54d061b9f74f9e2e1a03ff1178821502c.
//
// Solidity: event Withdrew(bytes32 indexed transferId, address indexed recipient, uint256 amount, bytes32 transferNonce)
func (_L2Bridge *L2BridgeFilterer) FilterWithdrew(opts *bind.FilterOpts, transferId [][32]byte, recipient []common.Address) (*L2BridgeWithdrewIterator, error) {

	var transferIdRule []interface{}
	for _, transferIdItem := range transferId {
		transferIdRule = append(transferIdRule, transferIdItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _L2Bridge.contract.FilterLogs(opts, "Withdrew", transferIdRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &L2BridgeWithdrewIterator{contract: _L2Bridge.contract, event: "Withdrew", logs: logs, sub: sub}, nil
}

// WatchWithdrew is a free log subscription operation binding the contract event 0x9475cdbde5fc71fe2ccd413c82878ee54d061b9f74f9e2e1a03ff1178821502c.
//
// Solidity: event Withdrew(bytes32 indexed transferId, address indexed recipient, uint256 amount, bytes32 transferNonce)
func (_L2Bridge *L2BridgeFilterer) WatchWithdrew(opts *bind.WatchOpts, sink chan<- *L2BridgeWithdrew, transferId [][32]byte, recipient []common.Address) (event.Subscription, error) {

	var transferIdRule []interface{}
	for _, transferIdItem := range transferId {
		transferIdRule = append(transferIdRule, transferIdItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _L2Bridge.contract.WatchLogs(opts, "Withdrew", transferIdRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2BridgeWithdrew)
				if err := _L2Bridge.contract.UnpackLog(event, "Withdrew", log); err != nil {
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

// ParseWithdrew is a log parse operation binding the contract event 0x9475cdbde5fc71fe2ccd413c82878ee54d061b9f74f9e2e1a03ff1178821502c.
//
// Solidity: event Withdrew(bytes32 indexed transferId, address indexed recipient, uint256 amount, bytes32 transferNonce)
func (_L2Bridge *L2BridgeFilterer) ParseWithdrew(log types.Log) (*L2BridgeWithdrew, error) {
	event := new(L2BridgeWithdrew)
	if err := _L2Bridge.contract.UnpackLog(event, "Withdrew", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
