// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lido

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

// IStMATICRequestWithdraw is an auto generated low-level Go binding around an user-defined struct.
type IStMATICRequestWithdraw struct {
	Amount2WithdrawFromStMATIC *big.Int
	ValidatorNonce             *big.Int
	RequestEpoch               *big.Int
	ValidatorAddress           common.Address
}

// MaticMetaData contains all meta data concerning the Matic contract.
var MaticMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_amountClaimed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amountBurned\",\"type\":\"uint256\"}],\"name\":\"ClaimTokensEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_amountClaimed\",\"type\":\"uint256\"}],\"name\":\"ClaimTotalDelegatedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_amountDelegated\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_remainder\",\"type\":\"uint256\"}],\"name\":\"DelegateEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"DistributeRewardsEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"RequestWithdrawEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"SubmitEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawTotalDelegatedEvent\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DAO\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount2WithdrawFromStMATIC\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requestEpoch\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"}],\"internalType\":\"structIStMATIC.RequestWithdraw\",\"name\":\"requestData\",\"type\":\"tuple\"}],\"name\":\"_getMaticFromRequestWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculatePendingBufferedTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"pendingBufferedTokens\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"claimTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"claimTotalDelegated2StMatic\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_balance\",\"type\":\"uint256\"}],\"name\":\"convertMaticToStMatic\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_balance\",\"type\":\"uint256\"}],\"name\":\"convertStMaticToMatic\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dao\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delegationLowerBound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"distributeRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"entityFees\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"dao\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"operators\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"insurance\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flipSubmitHandler\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fxStateRootTunnel\",\"outputs\":[{\"internalType\":\"contractIFxStateRootTunnel\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIValidatorShare\",\"name\":\"_validatorShare\",\"type\":\"address\"}],\"name\":\"getLiquidRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getMaticFromTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinValidatorBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalPooledMatic\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIValidatorShare\",\"name\":\"_validatorShare\",\"type\":\"address\"}],\"name\":\"getTotalStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalStakeAcrossAllValidators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalWithdrawRequest\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount2WithdrawFromStMATIC\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requestEpoch\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"}],\"internalType\":\"structIStMATIC.RequestWithdraw[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeOperatorRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_dao\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_insurance\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stakeManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_poLidoNFT\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_fxStateRootTunnel\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_submitThreshold\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"insurance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastWithdrawnValidatorId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nodeOperatorRegistry\",\"outputs\":[{\"internalType\":\"contractINodeOperatorRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poLidoNFT\",\"outputs\":[{\"internalType\":\"contractIPoLidoNFT\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"requestWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reservedFunds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardDistributionLowerBound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setDaoAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_delegationLowerBound\",\"type\":\"uint256\"}],\"name\":\"setDelegationLowerBound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_daoFee\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_operatorsFee\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_insuranceFee\",\"type\":\"uint8\"}],\"name\":\"setFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_fxStateRootTunnel\",\"type\":\"address\"}],\"name\":\"setFxStateRootTunnel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setInsuranceAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setNodeOperatorRegistryAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_poLidoNFT\",\"type\":\"address\"}],\"name\":\"setPoLidoNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardDistributionLowerBound\",\"type\":\"uint256\"}],\"name\":\"setRewardDistributionLowerBound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_submitThreshold\",\"type\":\"uint256\"}],\"name\":\"setSubmitThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_version\",\"type\":\"string\"}],\"name\":\"setVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stMaticWithdrawRequest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount2WithdrawFromStMATIC\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requestEpoch\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeManager\",\"outputs\":[{\"internalType\":\"contractIStakeManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"submit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"submitHandler\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"submitThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"togglePause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"token2WithdrawRequest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount2WithdrawFromStMATIC\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requestEpoch\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBuffered\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validatorShare\",\"type\":\"address\"}],\"name\":\"withdrawTotalDelegated\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MaticABI is the input ABI used to generate the binding from.
// Deprecated: Use MaticMetaData.ABI instead.
var MaticABI = MaticMetaData.ABI

// Matic is an auto generated Go binding around an Ethereum contract.
type Matic struct {
	MaticCaller     // Read-only binding to the contract
	MaticTransactor // Write-only binding to the contract
	MaticFilterer   // Log filterer for contract events
}

// MaticCaller is an auto generated read-only Go binding around an Ethereum contract.
type MaticCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MaticTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MaticTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MaticFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MaticFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MaticSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MaticSession struct {
	Contract     *Matic            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MaticCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MaticCallerSession struct {
	Contract *MaticCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MaticTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MaticTransactorSession struct {
	Contract     *MaticTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MaticRaw is an auto generated low-level Go binding around an Ethereum contract.
type MaticRaw struct {
	Contract *Matic // Generic contract binding to access the raw methods on
}

// MaticCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MaticCallerRaw struct {
	Contract *MaticCaller // Generic read-only contract binding to access the raw methods on
}

// MaticTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MaticTransactorRaw struct {
	Contract *MaticTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMatic creates a new instance of Matic, bound to a specific deployed contract.
func NewMatic(address common.Address, backend bind.ContractBackend) (*Matic, error) {
	contract, err := bindMatic(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Matic{MaticCaller: MaticCaller{contract: contract}, MaticTransactor: MaticTransactor{contract: contract}, MaticFilterer: MaticFilterer{contract: contract}}, nil
}

// NewMaticCaller creates a new read-only instance of Matic, bound to a specific deployed contract.
func NewMaticCaller(address common.Address, caller bind.ContractCaller) (*MaticCaller, error) {
	contract, err := bindMatic(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MaticCaller{contract: contract}, nil
}

// NewMaticTransactor creates a new write-only instance of Matic, bound to a specific deployed contract.
func NewMaticTransactor(address common.Address, transactor bind.ContractTransactor) (*MaticTransactor, error) {
	contract, err := bindMatic(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MaticTransactor{contract: contract}, nil
}

// NewMaticFilterer creates a new log filterer instance of Matic, bound to a specific deployed contract.
func NewMaticFilterer(address common.Address, filterer bind.ContractFilterer) (*MaticFilterer, error) {
	contract, err := bindMatic(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MaticFilterer{contract: contract}, nil
}

// bindMatic binds a generic wrapper to an already deployed contract.
func bindMatic(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MaticABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Matic *MaticRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Matic.Contract.MaticCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Matic *MaticRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Matic.Contract.MaticTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Matic *MaticRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Matic.Contract.MaticTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Matic *MaticCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Matic.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Matic *MaticTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Matic.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Matic *MaticTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Matic.Contract.contract.Transact(opts, method, params...)
}

// DAO is a free data retrieval call binding the contract method 0x98fabd3a.
//
// Solidity: function DAO() view returns(bytes32)
func (_Matic *MaticCaller) DAO(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Matic.contract.Call(opts, &out, "DAO")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DAO is a free data retrieval call binding the contract method 0x98fabd3a.
//
// Solidity: function DAO() view returns(bytes32)
func (_Matic *MaticSession) DAO() ([32]byte, error) {
	return _Matic.Contract.DAO(&_Matic.CallOpts)
}

// DAO is a free data retrieval call binding the contract method 0x98fabd3a.
//
// Solidity: function DAO() view returns(bytes32)
func (_Matic *MaticCallerSession) DAO() ([32]byte, error) {
	return _Matic.Contract.DAO(&_Matic.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Matic *MaticCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Matic.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Matic *MaticSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Matic.Contract.DEFAULTADMINROLE(&_Matic.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Matic *MaticCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Matic.Contract.DEFAULTADMINROLE(&_Matic.CallOpts)
}

// GetMaticFromRequestWithdraw is a free data retrieval call binding the contract method 0x65ebbeed.
//
// Solidity: function _getMaticFromRequestWithdraw((uint256,uint256,uint256,address) requestData) view returns(uint256)
func (_Matic *MaticCaller) GetMaticFromRequestWithdraw(opts *bind.CallOpts, requestData IStMATICRequestWithdraw) (*big.Int, error) {
	var out []interface{}
	err := _Matic.contract.Call(opts, &out, "_getMaticFromRequestWithdraw", requestData)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaticFromRequestWithdraw is a free data retrieval call binding the contract method 0x65ebbeed.
//
// Solidity: function _getMaticFromRequestWithdraw((uint256,uint256,uint256,address) requestData) view returns(uint256)
func (_Matic *MaticSession) GetMaticFromRequestWithdraw(requestData IStMATICRequestWithdraw) (*big.Int, error) {
	return _Matic.Contract.GetMaticFromRequestWithdraw(&_Matic.CallOpts, requestData)
}

// GetMaticFromRequestWithdraw is a free data retrieval call binding the contract method 0x65ebbeed.
//
// Solidity: function _getMaticFromRequestWithdraw((uint256,uint256,uint256,address) requestData) view returns(uint256)
func (_Matic *MaticCallerSession) GetMaticFromRequestWithdraw(requestData IStMATICRequestWithdraw) (*big.Int, error) {
	return _Matic.Contract.GetMaticFromRequestWithdraw(&_Matic.CallOpts, requestData)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Matic *MaticCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Matic.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Matic *MaticSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Matic.Contract.Allowance(&_Matic.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Matic *MaticCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Matic.Contract.Allowance(&_Matic.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Matic *MaticCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Matic.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Matic *MaticSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Matic.Contract.BalanceOf(&_Matic.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Matic *MaticCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Matic.Contract.BalanceOf(&_Matic.CallOpts, account)
}

// CalculatePendingBufferedTokens is a free data retrieval call binding the contract method 0xafd290a7.
//
// Solidity: function calculatePendingBufferedTokens() view returns(uint256 pendingBufferedTokens)
func (_Matic *MaticCaller) CalculatePendingBufferedTokens(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Matic.contract.Call(opts, &out, "calculatePendingBufferedTokens")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculatePendingBufferedTokens is a free data retrieval call binding the contract method 0xafd290a7.
//
// Solidity: function calculatePendingBufferedTokens() view returns(uint256 pendingBufferedTokens)
func (_Matic *MaticSession) CalculatePendingBufferedTokens() (*big.Int, error) {
	return _Matic.Contract.CalculatePendingBufferedTokens(&_Matic.CallOpts)
}

// CalculatePendingBufferedTokens is a free data retrieval call binding the contract method 0xafd290a7.
//
// Solidity: function calculatePendingBufferedTokens() view returns(uint256 pendingBufferedTokens)
func (_Matic *MaticCallerSession) CalculatePendingBufferedTokens() (*big.Int, error) {
	return _Matic.Contract.CalculatePendingBufferedTokens(&_Matic.CallOpts)
}

// ConvertMaticToStMatic is a free data retrieval call binding the contract method 0x917a52f5.
//
// Solidity: function convertMaticToStMatic(uint256 _balance) view returns(uint256, uint256, uint256)
func (_Matic *MaticCaller) ConvertMaticToStMatic(opts *bind.CallOpts, _balance *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Matic.contract.Call(opts, &out, "convertMaticToStMatic", _balance)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// ConvertMaticToStMatic is a free data retrieval call binding the contract method 0x917a52f5.
//
// Solidity: function convertMaticToStMatic(uint256 _balance) view returns(uint256, uint256, uint256)
func (_Matic *MaticSession) ConvertMaticToStMatic(_balance *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _Matic.Contract.ConvertMaticToStMatic(&_Matic.CallOpts, _balance)
}

// ConvertMaticToStMatic is a free data retrieval call binding the contract method 0x917a52f5.
//
// Solidity: function convertMaticToStMatic(uint256 _balance) view returns(uint256, uint256, uint256)
func (_Matic *MaticCallerSession) ConvertMaticToStMatic(_balance *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _Matic.Contract.ConvertMaticToStMatic(&_Matic.CallOpts, _balance)
}

// ConvertStMaticToMatic is a free data retrieval call binding the contract method 0xd968447c.
//
// Solidity: function convertStMaticToMatic(uint256 _balance) view returns(uint256, uint256, uint256)
func (_Matic *MaticCaller) ConvertStMaticToMatic(opts *bind.CallOpts, _balance *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Matic.contract.Call(opts, &out, "convertStMaticToMatic", _balance)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// ConvertStMaticToMatic is a free data retrieval call binding the contract method 0xd968447c.
//
// Solidity: function convertStMaticToMatic(uint256 _balance) view returns(uint256, uint256, uint256)
func (_Matic *MaticSession) ConvertStMaticToMatic(_balance *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _Matic.Contract.ConvertStMaticToMatic(&_Matic.CallOpts, _balance)
}

// ConvertStMaticToMatic is a free data retrieval call binding the contract method 0xd968447c.
//
// Solidity: function convertStMaticToMatic(uint256 _balance) view returns(uint256, uint256, uint256)
func (_Matic *MaticCallerSession) ConvertStMaticToMatic(_balance *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _Matic.Contract.ConvertStMaticToMatic(&_Matic.CallOpts, _balance)
}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(address)
func (_Matic *MaticCaller) Dao(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Matic.contract.Call(opts, &out, "dao")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(address)
func (_Matic *MaticSession) Dao() (common.Address, error) {
	return _Matic.Contract.Dao(&_Matic.CallOpts)
}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(address)
func (_Matic *MaticCallerSession) Dao() (common.Address, error) {
	return _Matic.Contract.Dao(&_Matic.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Matic *MaticCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Matic.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Matic *MaticSession) Decimals() (uint8, error) {
	return _Matic.Contract.Decimals(&_Matic.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Matic *MaticCallerSession) Decimals() (uint8, error) {
	return _Matic.Contract.Decimals(&_Matic.CallOpts)
}

// DelegationLowerBound is a free data retrieval call binding the contract method 0x0d7abc33.
//
// Solidity: function delegationLowerBound() view returns(uint256)
func (_Matic *MaticCaller) DelegationLowerBound(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Matic.contract.Call(opts, &out, "delegationLowerBound")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelegationLowerBound is a free data retrieval call binding the contract method 0x0d7abc33.
//
// Solidity: function delegationLowerBound() view returns(uint256)
func (_Matic *MaticSession) DelegationLowerBound() (*big.Int, error) {
	return _Matic.Contract.DelegationLowerBound(&_Matic.CallOpts)
}

// DelegationLowerBound is a free data retrieval call binding the contract method 0x0d7abc33.
//
// Solidity: function delegationLowerBound() view returns(uint256)
func (_Matic *MaticCallerSession) DelegationLowerBound() (*big.Int, error) {
	return _Matic.Contract.DelegationLowerBound(&_Matic.CallOpts)
}

// EntityFees is a free data retrieval call binding the contract method 0x964a7596.
//
// Solidity: function entityFees() view returns(uint8 dao, uint8 operators, uint8 insurance)
func (_Matic *MaticCaller) EntityFees(opts *bind.CallOpts) (struct {
	Dao       uint8
	Operators uint8
	Insurance uint8
}, error) {
	var out []interface{}
	err := _Matic.contract.Call(opts, &out, "entityFees")

	outstruct := new(struct {
		Dao       uint8
		Operators uint8
		Insurance uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Dao = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.Operators = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.Insurance = *abi.ConvertType(out[2], new(uint8)).(*uint8)

	return *outstruct, err

}

// EntityFees is a free data retrieval call binding the contract method 0x964a7596.
//
// Solidity: function entityFees() view returns(uint8 dao, uint8 operators, uint8 insurance)
func (_Matic *MaticSession) EntityFees() (struct {
	Dao       uint8
	Operators uint8
	Insurance uint8
}, error) {
	return _Matic.Contract.EntityFees(&_Matic.CallOpts)
}

// EntityFees is a free data retrieval call binding the contract method 0x964a7596.
//
// Solidity: function entityFees() view returns(uint8 dao, uint8 operators, uint8 insurance)
func (_Matic *MaticCallerSession) EntityFees() (struct {
	Dao       uint8
	Operators uint8
	Insurance uint8
}, error) {
	return _Matic.Contract.EntityFees(&_Matic.CallOpts)
}

// FxStateRootTunnel is a free data retrieval call binding the contract method 0xe062b10b.
//
// Solidity: function fxStateRootTunnel() view returns(address)
func (_Matic *MaticCaller) FxStateRootTunnel(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Matic.contract.Call(opts, &out, "fxStateRootTunnel")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FxStateRootTunnel is a free data retrieval call binding the contract method 0xe062b10b.
//
// Solidity: function fxStateRootTunnel() view returns(address)
func (_Matic *MaticSession) FxStateRootTunnel() (common.Address, error) {
	return _Matic.Contract.FxStateRootTunnel(&_Matic.CallOpts)
}

// FxStateRootTunnel is a free data retrieval call binding the contract method 0xe062b10b.
//
// Solidity: function fxStateRootTunnel() view returns(address)
func (_Matic *MaticCallerSession) FxStateRootTunnel() (common.Address, error) {
	return _Matic.Contract.FxStateRootTunnel(&_Matic.CallOpts)
}

// GetLiquidRewards is a free data retrieval call binding the contract method 0x676e5550.
//
// Solidity: function getLiquidRewards(address _validatorShare) view returns(uint256)
func (_Matic *MaticCaller) GetLiquidRewards(opts *bind.CallOpts, _validatorShare common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Matic.contract.Call(opts, &out, "getLiquidRewards", _validatorShare)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLiquidRewards is a free data retrieval call binding the contract method 0x676e5550.
//
// Solidity: function getLiquidRewards(address _validatorShare) view returns(uint256)
func (_Matic *MaticSession) GetLiquidRewards(_validatorShare common.Address) (*big.Int, error) {
	return _Matic.Contract.GetLiquidRewards(&_Matic.CallOpts, _validatorShare)
}

// GetLiquidRewards is a free data retrieval call binding the contract method 0x676e5550.
//
// Solidity: function getLiquidRewards(address _validatorShare) view returns(uint256)
func (_Matic *MaticCallerSession) GetLiquidRewards(_validatorShare common.Address) (*big.Int, error) {
	return _Matic.Contract.GetLiquidRewards(&_Matic.CallOpts, _validatorShare)
}

// GetMaticFromTokenId is a free data retrieval call binding the contract method 0x720bcf1d.
//
// Solidity: function getMaticFromTokenId(uint256 _tokenId) view returns(uint256)
func (_Matic *MaticCaller) GetMaticFromTokenId(opts *bind.CallOpts, _tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Matic.contract.Call(opts, &out, "getMaticFromTokenId", _tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaticFromTokenId is a free data retrieval call binding the contract method 0x720bcf1d.
//
// Solidity: function getMaticFromTokenId(uint256 _tokenId) view returns(uint256)
func (_Matic *MaticSession) GetMaticFromTokenId(_tokenId *big.Int) (*big.Int, error) {
	return _Matic.Contract.GetMaticFromTokenId(&_Matic.CallOpts, _tokenId)
}

// GetMaticFromTokenId is a free data retrieval call binding the contract method 0x720bcf1d.
//
// Solidity: function getMaticFromTokenId(uint256 _tokenId) view returns(uint256)
func (_Matic *MaticCallerSession) GetMaticFromTokenId(_tokenId *big.Int) (*big.Int, error) {
	return _Matic.Contract.GetMaticFromTokenId(&_Matic.CallOpts, _tokenId)
}

// GetMinValidatorBalance is a free data retrieval call binding the contract method 0x0d946b71.
//
// Solidity: function getMinValidatorBalance() view returns(uint256)
func (_Matic *MaticCaller) GetMinValidatorBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Matic.contract.Call(opts, &out, "getMinValidatorBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinValidatorBalance is a free data retrieval call binding the contract method 0x0d946b71.
//
// Solidity: function getMinValidatorBalance() view returns(uint256)
func (_Matic *MaticSession) GetMinValidatorBalance() (*big.Int, error) {
	return _Matic.Contract.GetMinValidatorBalance(&_Matic.CallOpts)
}

// GetMinValidatorBalance is a free data retrieval call binding the contract method 0x0d946b71.
//
// Solidity: function getMinValidatorBalance() view returns(uint256)
func (_Matic *MaticCallerSession) GetMinValidatorBalance() (*big.Int, error) {
	return _Matic.Contract.GetMinValidatorBalance(&_Matic.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Matic *MaticCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Matic.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Matic *MaticSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Matic.Contract.GetRoleAdmin(&_Matic.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Matic *MaticCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Matic.Contract.GetRoleAdmin(&_Matic.CallOpts, role)
}

// GetTotalPooledMatic is a free data retrieval call binding the contract method 0xe00222a0.
//
// Solidity: function getTotalPooledMatic() view returns(uint256)
func (_Matic *MaticCaller) GetTotalPooledMatic(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Matic.contract.Call(opts, &out, "getTotalPooledMatic")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalPooledMatic is a free data retrieval call binding the contract method 0xe00222a0.
//
// Solidity: function getTotalPooledMatic() view returns(uint256)
func (_Matic *MaticSession) GetTotalPooledMatic() (*big.Int, error) {
	return _Matic.Contract.GetTotalPooledMatic(&_Matic.CallOpts)
}

// GetTotalPooledMatic is a free data retrieval call binding the contract method 0xe00222a0.
//
// Solidity: function getTotalPooledMatic() view returns(uint256)
func (_Matic *MaticCallerSession) GetTotalPooledMatic() (*big.Int, error) {
	return _Matic.Contract.GetTotalPooledMatic(&_Matic.CallOpts)
}

// GetTotalStake is a free data retrieval call binding the contract method 0x1e7ff8f6.
//
// Solidity: function getTotalStake(address _validatorShare) view returns(uint256, uint256)
func (_Matic *MaticCaller) GetTotalStake(opts *bind.CallOpts, _validatorShare common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Matic.contract.Call(opts, &out, "getTotalStake", _validatorShare)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetTotalStake is a free data retrieval call binding the contract method 0x1e7ff8f6.
//
// Solidity: function getTotalStake(address _validatorShare) view returns(uint256, uint256)
func (_Matic *MaticSession) GetTotalStake(_validatorShare common.Address) (*big.Int, *big.Int, error) {
	return _Matic.Contract.GetTotalStake(&_Matic.CallOpts, _validatorShare)
}

// GetTotalStake is a free data retrieval call binding the contract method 0x1e7ff8f6.
//
// Solidity: function getTotalStake(address _validatorShare) view returns(uint256, uint256)
func (_Matic *MaticCallerSession) GetTotalStake(_validatorShare common.Address) (*big.Int, *big.Int, error) {
	return _Matic.Contract.GetTotalStake(&_Matic.CallOpts, _validatorShare)
}

// GetTotalStakeAcrossAllValidators is a free data retrieval call binding the contract method 0x7e978af8.
//
// Solidity: function getTotalStakeAcrossAllValidators() view returns(uint256)
func (_Matic *MaticCaller) GetTotalStakeAcrossAllValidators(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Matic.contract.Call(opts, &out, "getTotalStakeAcrossAllValidators")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalStakeAcrossAllValidators is a free data retrieval call binding the contract method 0x7e978af8.
//
// Solidity: function getTotalStakeAcrossAllValidators() view returns(uint256)
func (_Matic *MaticSession) GetTotalStakeAcrossAllValidators() (*big.Int, error) {
	return _Matic.Contract.GetTotalStakeAcrossAllValidators(&_Matic.CallOpts)
}

// GetTotalStakeAcrossAllValidators is a free data retrieval call binding the contract method 0x7e978af8.
//
// Solidity: function getTotalStakeAcrossAllValidators() view returns(uint256)
func (_Matic *MaticCallerSession) GetTotalStakeAcrossAllValidators() (*big.Int, error) {
	return _Matic.Contract.GetTotalStakeAcrossAllValidators(&_Matic.CallOpts)
}

// GetTotalWithdrawRequest is a free data retrieval call binding the contract method 0x916b9eba.
//
// Solidity: function getTotalWithdrawRequest() view returns((uint256,uint256,uint256,address)[])
func (_Matic *MaticCaller) GetTotalWithdrawRequest(opts *bind.CallOpts) ([]IStMATICRequestWithdraw, error) {
	var out []interface{}
	err := _Matic.contract.Call(opts, &out, "getTotalWithdrawRequest")

	if err != nil {
		return *new([]IStMATICRequestWithdraw), err
	}

	out0 := *abi.ConvertType(out[0], new([]IStMATICRequestWithdraw)).(*[]IStMATICRequestWithdraw)

	return out0, err

}

// GetTotalWithdrawRequest is a free data retrieval call binding the contract method 0x916b9eba.
//
// Solidity: function getTotalWithdrawRequest() view returns((uint256,uint256,uint256,address)[])
func (_Matic *MaticSession) GetTotalWithdrawRequest() ([]IStMATICRequestWithdraw, error) {
	return _Matic.Contract.GetTotalWithdrawRequest(&_Matic.CallOpts)
}

// GetTotalWithdrawRequest is a free data retrieval call binding the contract method 0x916b9eba.
//
// Solidity: function getTotalWithdrawRequest() view returns((uint256,uint256,uint256,address)[])
func (_Matic *MaticCallerSession) GetTotalWithdrawRequest() ([]IStMATICRequestWithdraw, error) {
	return _Matic.Contract.GetTotalWithdrawRequest(&_Matic.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Matic *MaticCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Matic.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Matic *MaticSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Matic.Contract.HasRole(&_Matic.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Matic *MaticCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Matic.Contract.HasRole(&_Matic.CallOpts, role, account)
}

// Insurance is a free data retrieval call binding the contract method 0x89cf3204.
//
// Solidity: function insurance() view returns(address)
func (_Matic *MaticCaller) Insurance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Matic.contract.Call(opts, &out, "insurance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Insurance is a free data retrieval call binding the contract method 0x89cf3204.
//
// Solidity: function insurance() view returns(address)
func (_Matic *MaticSession) Insurance() (common.Address, error) {
	return _Matic.Contract.Insurance(&_Matic.CallOpts)
}

// Insurance is a free data retrieval call binding the contract method 0x89cf3204.
//
// Solidity: function insurance() view returns(address)
func (_Matic *MaticCallerSession) Insurance() (common.Address, error) {
	return _Matic.Contract.Insurance(&_Matic.CallOpts)
}

// LastWithdrawnValidatorId is a free data retrieval call binding the contract method 0x71975a3e.
//
// Solidity: function lastWithdrawnValidatorId() view returns(uint256)
func (_Matic *MaticCaller) LastWithdrawnValidatorId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Matic.contract.Call(opts, &out, "lastWithdrawnValidatorId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastWithdrawnValidatorId is a free data retrieval call binding the contract method 0x71975a3e.
//
// Solidity: function lastWithdrawnValidatorId() view returns(uint256)
func (_Matic *MaticSession) LastWithdrawnValidatorId() (*big.Int, error) {
	return _Matic.Contract.LastWithdrawnValidatorId(&_Matic.CallOpts)
}

// LastWithdrawnValidatorId is a free data retrieval call binding the contract method 0x71975a3e.
//
// Solidity: function lastWithdrawnValidatorId() view returns(uint256)
func (_Matic *MaticCallerSession) LastWithdrawnValidatorId() (*big.Int, error) {
	return _Matic.Contract.LastWithdrawnValidatorId(&_Matic.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Matic *MaticCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Matic.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Matic *MaticSession) Name() (string, error) {
	return _Matic.Contract.Name(&_Matic.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Matic *MaticCallerSession) Name() (string, error) {
	return _Matic.Contract.Name(&_Matic.CallOpts)
}

// NodeOperatorRegistry is a free data retrieval call binding the contract method 0xe8f8708f.
//
// Solidity: function nodeOperatorRegistry() view returns(address)
func (_Matic *MaticCaller) NodeOperatorRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Matic.contract.Call(opts, &out, "nodeOperatorRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NodeOperatorRegistry is a free data retrieval call binding the contract method 0xe8f8708f.
//
// Solidity: function nodeOperatorRegistry() view returns(address)
func (_Matic *MaticSession) NodeOperatorRegistry() (common.Address, error) {
	return _Matic.Contract.NodeOperatorRegistry(&_Matic.CallOpts)
}

// NodeOperatorRegistry is a free data retrieval call binding the contract method 0xe8f8708f.
//
// Solidity: function nodeOperatorRegistry() view returns(address)
func (_Matic *MaticCallerSession) NodeOperatorRegistry() (common.Address, error) {
	return _Matic.Contract.NodeOperatorRegistry(&_Matic.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Matic *MaticCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Matic.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Matic *MaticSession) Paused() (bool, error) {
	return _Matic.Contract.Paused(&_Matic.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Matic *MaticCallerSession) Paused() (bool, error) {
	return _Matic.Contract.Paused(&_Matic.CallOpts)
}

// PoLidoNFT is a free data retrieval call binding the contract method 0x7029c90e.
//
// Solidity: function poLidoNFT() view returns(address)
func (_Matic *MaticCaller) PoLidoNFT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Matic.contract.Call(opts, &out, "poLidoNFT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoLidoNFT is a free data retrieval call binding the contract method 0x7029c90e.
//
// Solidity: function poLidoNFT() view returns(address)
func (_Matic *MaticSession) PoLidoNFT() (common.Address, error) {
	return _Matic.Contract.PoLidoNFT(&_Matic.CallOpts)
}

// PoLidoNFT is a free data retrieval call binding the contract method 0x7029c90e.
//
// Solidity: function poLidoNFT() view returns(address)
func (_Matic *MaticCallerSession) PoLidoNFT() (common.Address, error) {
	return _Matic.Contract.PoLidoNFT(&_Matic.CallOpts)
}

// ReservedFunds is a free data retrieval call binding the contract method 0x509c5df6.
//
// Solidity: function reservedFunds() view returns(uint256)
func (_Matic *MaticCaller) ReservedFunds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Matic.contract.Call(opts, &out, "reservedFunds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReservedFunds is a free data retrieval call binding the contract method 0x509c5df6.
//
// Solidity: function reservedFunds() view returns(uint256)
func (_Matic *MaticSession) ReservedFunds() (*big.Int, error) {
	return _Matic.Contract.ReservedFunds(&_Matic.CallOpts)
}

// ReservedFunds is a free data retrieval call binding the contract method 0x509c5df6.
//
// Solidity: function reservedFunds() view returns(uint256)
func (_Matic *MaticCallerSession) ReservedFunds() (*big.Int, error) {
	return _Matic.Contract.ReservedFunds(&_Matic.CallOpts)
}

// RewardDistributionLowerBound is a free data retrieval call binding the contract method 0xa2452947.
//
// Solidity: function rewardDistributionLowerBound() view returns(uint256)
func (_Matic *MaticCaller) RewardDistributionLowerBound(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Matic.contract.Call(opts, &out, "rewardDistributionLowerBound")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardDistributionLowerBound is a free data retrieval call binding the contract method 0xa2452947.
//
// Solidity: function rewardDistributionLowerBound() view returns(uint256)
func (_Matic *MaticSession) RewardDistributionLowerBound() (*big.Int, error) {
	return _Matic.Contract.RewardDistributionLowerBound(&_Matic.CallOpts)
}

// RewardDistributionLowerBound is a free data retrieval call binding the contract method 0xa2452947.
//
// Solidity: function rewardDistributionLowerBound() view returns(uint256)
func (_Matic *MaticCallerSession) RewardDistributionLowerBound() (*big.Int, error) {
	return _Matic.Contract.RewardDistributionLowerBound(&_Matic.CallOpts)
}

// StMaticWithdrawRequest is a free data retrieval call binding the contract method 0xf1a13fce.
//
// Solidity: function stMaticWithdrawRequest(uint256 ) view returns(uint256 amount2WithdrawFromStMATIC, uint256 validatorNonce, uint256 requestEpoch, address validatorAddress)
func (_Matic *MaticCaller) StMaticWithdrawRequest(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Amount2WithdrawFromStMATIC *big.Int
	ValidatorNonce             *big.Int
	RequestEpoch               *big.Int
	ValidatorAddress           common.Address
}, error) {
	var out []interface{}
	err := _Matic.contract.Call(opts, &out, "stMaticWithdrawRequest", arg0)

	outstruct := new(struct {
		Amount2WithdrawFromStMATIC *big.Int
		ValidatorNonce             *big.Int
		RequestEpoch               *big.Int
		ValidatorAddress           common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount2WithdrawFromStMATIC = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ValidatorNonce = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.RequestEpoch = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.ValidatorAddress = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// StMaticWithdrawRequest is a free data retrieval call binding the contract method 0xf1a13fce.
//
// Solidity: function stMaticWithdrawRequest(uint256 ) view returns(uint256 amount2WithdrawFromStMATIC, uint256 validatorNonce, uint256 requestEpoch, address validatorAddress)
func (_Matic *MaticSession) StMaticWithdrawRequest(arg0 *big.Int) (struct {
	Amount2WithdrawFromStMATIC *big.Int
	ValidatorNonce             *big.Int
	RequestEpoch               *big.Int
	ValidatorAddress           common.Address
}, error) {
	return _Matic.Contract.StMaticWithdrawRequest(&_Matic.CallOpts, arg0)
}

// StMaticWithdrawRequest is a free data retrieval call binding the contract method 0xf1a13fce.
//
// Solidity: function stMaticWithdrawRequest(uint256 ) view returns(uint256 amount2WithdrawFromStMATIC, uint256 validatorNonce, uint256 requestEpoch, address validatorAddress)
func (_Matic *MaticCallerSession) StMaticWithdrawRequest(arg0 *big.Int) (struct {
	Amount2WithdrawFromStMATIC *big.Int
	ValidatorNonce             *big.Int
	RequestEpoch               *big.Int
	ValidatorAddress           common.Address
}, error) {
	return _Matic.Contract.StMaticWithdrawRequest(&_Matic.CallOpts, arg0)
}

// StakeManager is a free data retrieval call binding the contract method 0x7542ff95.
//
// Solidity: function stakeManager() view returns(address)
func (_Matic *MaticCaller) StakeManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Matic.contract.Call(opts, &out, "stakeManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeManager is a free data retrieval call binding the contract method 0x7542ff95.
//
// Solidity: function stakeManager() view returns(address)
func (_Matic *MaticSession) StakeManager() (common.Address, error) {
	return _Matic.Contract.StakeManager(&_Matic.CallOpts)
}

// StakeManager is a free data retrieval call binding the contract method 0x7542ff95.
//
// Solidity: function stakeManager() view returns(address)
func (_Matic *MaticCallerSession) StakeManager() (common.Address, error) {
	return _Matic.Contract.StakeManager(&_Matic.CallOpts)
}

// SubmitHandler is a free data retrieval call binding the contract method 0xe259faf7.
//
// Solidity: function submitHandler() view returns(bool)
func (_Matic *MaticCaller) SubmitHandler(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Matic.contract.Call(opts, &out, "submitHandler")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SubmitHandler is a free data retrieval call binding the contract method 0xe259faf7.
//
// Solidity: function submitHandler() view returns(bool)
func (_Matic *MaticSession) SubmitHandler() (bool, error) {
	return _Matic.Contract.SubmitHandler(&_Matic.CallOpts)
}

// SubmitHandler is a free data retrieval call binding the contract method 0xe259faf7.
//
// Solidity: function submitHandler() view returns(bool)
func (_Matic *MaticCallerSession) SubmitHandler() (bool, error) {
	return _Matic.Contract.SubmitHandler(&_Matic.CallOpts)
}

// SubmitThreshold is a free data retrieval call binding the contract method 0x893818a3.
//
// Solidity: function submitThreshold() view returns(uint256)
func (_Matic *MaticCaller) SubmitThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Matic.contract.Call(opts, &out, "submitThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SubmitThreshold is a free data retrieval call binding the contract method 0x893818a3.
//
// Solidity: function submitThreshold() view returns(uint256)
func (_Matic *MaticSession) SubmitThreshold() (*big.Int, error) {
	return _Matic.Contract.SubmitThreshold(&_Matic.CallOpts)
}

// SubmitThreshold is a free data retrieval call binding the contract method 0x893818a3.
//
// Solidity: function submitThreshold() view returns(uint256)
func (_Matic *MaticCallerSession) SubmitThreshold() (*big.Int, error) {
	return _Matic.Contract.SubmitThreshold(&_Matic.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Matic *MaticCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Matic.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Matic *MaticSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Matic.Contract.SupportsInterface(&_Matic.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Matic *MaticCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Matic.Contract.SupportsInterface(&_Matic.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Matic *MaticCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Matic.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Matic *MaticSession) Symbol() (string, error) {
	return _Matic.Contract.Symbol(&_Matic.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Matic *MaticCallerSession) Symbol() (string, error) {
	return _Matic.Contract.Symbol(&_Matic.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Matic *MaticCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Matic.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Matic *MaticSession) Token() (common.Address, error) {
	return _Matic.Contract.Token(&_Matic.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Matic *MaticCallerSession) Token() (common.Address, error) {
	return _Matic.Contract.Token(&_Matic.CallOpts)
}

// Token2WithdrawRequest is a free data retrieval call binding the contract method 0xf08711fe.
//
// Solidity: function token2WithdrawRequest(uint256 ) view returns(uint256 amount2WithdrawFromStMATIC, uint256 validatorNonce, uint256 requestEpoch, address validatorAddress)
func (_Matic *MaticCaller) Token2WithdrawRequest(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Amount2WithdrawFromStMATIC *big.Int
	ValidatorNonce             *big.Int
	RequestEpoch               *big.Int
	ValidatorAddress           common.Address
}, error) {
	var out []interface{}
	err := _Matic.contract.Call(opts, &out, "token2WithdrawRequest", arg0)

	outstruct := new(struct {
		Amount2WithdrawFromStMATIC *big.Int
		ValidatorNonce             *big.Int
		RequestEpoch               *big.Int
		ValidatorAddress           common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount2WithdrawFromStMATIC = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ValidatorNonce = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.RequestEpoch = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.ValidatorAddress = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Token2WithdrawRequest is a free data retrieval call binding the contract method 0xf08711fe.
//
// Solidity: function token2WithdrawRequest(uint256 ) view returns(uint256 amount2WithdrawFromStMATIC, uint256 validatorNonce, uint256 requestEpoch, address validatorAddress)
func (_Matic *MaticSession) Token2WithdrawRequest(arg0 *big.Int) (struct {
	Amount2WithdrawFromStMATIC *big.Int
	ValidatorNonce             *big.Int
	RequestEpoch               *big.Int
	ValidatorAddress           common.Address
}, error) {
	return _Matic.Contract.Token2WithdrawRequest(&_Matic.CallOpts, arg0)
}

// Token2WithdrawRequest is a free data retrieval call binding the contract method 0xf08711fe.
//
// Solidity: function token2WithdrawRequest(uint256 ) view returns(uint256 amount2WithdrawFromStMATIC, uint256 validatorNonce, uint256 requestEpoch, address validatorAddress)
func (_Matic *MaticCallerSession) Token2WithdrawRequest(arg0 *big.Int) (struct {
	Amount2WithdrawFromStMATIC *big.Int
	ValidatorNonce             *big.Int
	RequestEpoch               *big.Int
	ValidatorAddress           common.Address
}, error) {
	return _Matic.Contract.Token2WithdrawRequest(&_Matic.CallOpts, arg0)
}

// TotalBuffered is a free data retrieval call binding the contract method 0x52349b17.
//
// Solidity: function totalBuffered() view returns(uint256)
func (_Matic *MaticCaller) TotalBuffered(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Matic.contract.Call(opts, &out, "totalBuffered")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBuffered is a free data retrieval call binding the contract method 0x52349b17.
//
// Solidity: function totalBuffered() view returns(uint256)
func (_Matic *MaticSession) TotalBuffered() (*big.Int, error) {
	return _Matic.Contract.TotalBuffered(&_Matic.CallOpts)
}

// TotalBuffered is a free data retrieval call binding the contract method 0x52349b17.
//
// Solidity: function totalBuffered() view returns(uint256)
func (_Matic *MaticCallerSession) TotalBuffered() (*big.Int, error) {
	return _Matic.Contract.TotalBuffered(&_Matic.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Matic *MaticCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Matic.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Matic *MaticSession) TotalSupply() (*big.Int, error) {
	return _Matic.Contract.TotalSupply(&_Matic.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Matic *MaticCallerSession) TotalSupply() (*big.Int, error) {
	return _Matic.Contract.TotalSupply(&_Matic.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Matic *MaticCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Matic.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Matic *MaticSession) Version() (string, error) {
	return _Matic.Contract.Version(&_Matic.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Matic *MaticCallerSession) Version() (string, error) {
	return _Matic.Contract.Version(&_Matic.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Matic *MaticTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Matic.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Matic *MaticSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Matic.Contract.Approve(&_Matic.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Matic *MaticTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Matic.Contract.Approve(&_Matic.TransactOpts, spender, amount)
}

// ClaimTokens is a paid mutator transaction binding the contract method 0x46e04a2f.
//
// Solidity: function claimTokens(uint256 _tokenId) returns()
func (_Matic *MaticTransactor) ClaimTokens(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _Matic.contract.Transact(opts, "claimTokens", _tokenId)
}

// ClaimTokens is a paid mutator transaction binding the contract method 0x46e04a2f.
//
// Solidity: function claimTokens(uint256 _tokenId) returns()
func (_Matic *MaticSession) ClaimTokens(_tokenId *big.Int) (*types.Transaction, error) {
	return _Matic.Contract.ClaimTokens(&_Matic.TransactOpts, _tokenId)
}

// ClaimTokens is a paid mutator transaction binding the contract method 0x46e04a2f.
//
// Solidity: function claimTokens(uint256 _tokenId) returns()
func (_Matic *MaticTransactorSession) ClaimTokens(_tokenId *big.Int) (*types.Transaction, error) {
	return _Matic.Contract.ClaimTokens(&_Matic.TransactOpts, _tokenId)
}

// ClaimTotalDelegated2StMatic is a paid mutator transaction binding the contract method 0x70af1d13.
//
// Solidity: function claimTotalDelegated2StMatic(uint256 _index) returns()
func (_Matic *MaticTransactor) ClaimTotalDelegated2StMatic(opts *bind.TransactOpts, _index *big.Int) (*types.Transaction, error) {
	return _Matic.contract.Transact(opts, "claimTotalDelegated2StMatic", _index)
}

// ClaimTotalDelegated2StMatic is a paid mutator transaction binding the contract method 0x70af1d13.
//
// Solidity: function claimTotalDelegated2StMatic(uint256 _index) returns()
func (_Matic *MaticSession) ClaimTotalDelegated2StMatic(_index *big.Int) (*types.Transaction, error) {
	return _Matic.Contract.ClaimTotalDelegated2StMatic(&_Matic.TransactOpts, _index)
}

// ClaimTotalDelegated2StMatic is a paid mutator transaction binding the contract method 0x70af1d13.
//
// Solidity: function claimTotalDelegated2StMatic(uint256 _index) returns()
func (_Matic *MaticTransactorSession) ClaimTotalDelegated2StMatic(_index *big.Int) (*types.Transaction, error) {
	return _Matic.Contract.ClaimTotalDelegated2StMatic(&_Matic.TransactOpts, _index)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Matic *MaticTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Matic.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Matic *MaticSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Matic.Contract.DecreaseAllowance(&_Matic.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Matic *MaticTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Matic.Contract.DecreaseAllowance(&_Matic.TransactOpts, spender, subtractedValue)
}

// Delegate is a paid mutator transaction binding the contract method 0xc89e4361.
//
// Solidity: function delegate() returns()
func (_Matic *MaticTransactor) Delegate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Matic.contract.Transact(opts, "delegate")
}

// Delegate is a paid mutator transaction binding the contract method 0xc89e4361.
//
// Solidity: function delegate() returns()
func (_Matic *MaticSession) Delegate() (*types.Transaction, error) {
	return _Matic.Contract.Delegate(&_Matic.TransactOpts)
}

// Delegate is a paid mutator transaction binding the contract method 0xc89e4361.
//
// Solidity: function delegate() returns()
func (_Matic *MaticTransactorSession) Delegate() (*types.Transaction, error) {
	return _Matic.Contract.Delegate(&_Matic.TransactOpts)
}

// DistributeRewards is a paid mutator transaction binding the contract method 0x6f4a2cd0.
//
// Solidity: function distributeRewards() returns()
func (_Matic *MaticTransactor) DistributeRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Matic.contract.Transact(opts, "distributeRewards")
}

// DistributeRewards is a paid mutator transaction binding the contract method 0x6f4a2cd0.
//
// Solidity: function distributeRewards() returns()
func (_Matic *MaticSession) DistributeRewards() (*types.Transaction, error) {
	return _Matic.Contract.DistributeRewards(&_Matic.TransactOpts)
}

// DistributeRewards is a paid mutator transaction binding the contract method 0x6f4a2cd0.
//
// Solidity: function distributeRewards() returns()
func (_Matic *MaticTransactorSession) DistributeRewards() (*types.Transaction, error) {
	return _Matic.Contract.DistributeRewards(&_Matic.TransactOpts)
}

// FlipSubmitHandler is a paid mutator transaction binding the contract method 0xc07c030e.
//
// Solidity: function flipSubmitHandler() returns()
func (_Matic *MaticTransactor) FlipSubmitHandler(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Matic.contract.Transact(opts, "flipSubmitHandler")
}

// FlipSubmitHandler is a paid mutator transaction binding the contract method 0xc07c030e.
//
// Solidity: function flipSubmitHandler() returns()
func (_Matic *MaticSession) FlipSubmitHandler() (*types.Transaction, error) {
	return _Matic.Contract.FlipSubmitHandler(&_Matic.TransactOpts)
}

// FlipSubmitHandler is a paid mutator transaction binding the contract method 0xc07c030e.
//
// Solidity: function flipSubmitHandler() returns()
func (_Matic *MaticTransactorSession) FlipSubmitHandler() (*types.Transaction, error) {
	return _Matic.Contract.FlipSubmitHandler(&_Matic.TransactOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Matic *MaticTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Matic.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Matic *MaticSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Matic.Contract.GrantRole(&_Matic.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Matic *MaticTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Matic.Contract.GrantRole(&_Matic.TransactOpts, role, account)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Matic *MaticTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Matic.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Matic *MaticSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Matic.Contract.IncreaseAllowance(&_Matic.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Matic *MaticTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Matic.Contract.IncreaseAllowance(&_Matic.TransactOpts, spender, addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0x88301911.
//
// Solidity: function initialize(address _nodeOperatorRegistry, address _token, address _dao, address _insurance, address _stakeManager, address _poLidoNFT, address _fxStateRootTunnel, uint256 _submitThreshold) returns()
func (_Matic *MaticTransactor) Initialize(opts *bind.TransactOpts, _nodeOperatorRegistry common.Address, _token common.Address, _dao common.Address, _insurance common.Address, _stakeManager common.Address, _poLidoNFT common.Address, _fxStateRootTunnel common.Address, _submitThreshold *big.Int) (*types.Transaction, error) {
	return _Matic.contract.Transact(opts, "initialize", _nodeOperatorRegistry, _token, _dao, _insurance, _stakeManager, _poLidoNFT, _fxStateRootTunnel, _submitThreshold)
}

// Initialize is a paid mutator transaction binding the contract method 0x88301911.
//
// Solidity: function initialize(address _nodeOperatorRegistry, address _token, address _dao, address _insurance, address _stakeManager, address _poLidoNFT, address _fxStateRootTunnel, uint256 _submitThreshold) returns()
func (_Matic *MaticSession) Initialize(_nodeOperatorRegistry common.Address, _token common.Address, _dao common.Address, _insurance common.Address, _stakeManager common.Address, _poLidoNFT common.Address, _fxStateRootTunnel common.Address, _submitThreshold *big.Int) (*types.Transaction, error) {
	return _Matic.Contract.Initialize(&_Matic.TransactOpts, _nodeOperatorRegistry, _token, _dao, _insurance, _stakeManager, _poLidoNFT, _fxStateRootTunnel, _submitThreshold)
}

// Initialize is a paid mutator transaction binding the contract method 0x88301911.
//
// Solidity: function initialize(address _nodeOperatorRegistry, address _token, address _dao, address _insurance, address _stakeManager, address _poLidoNFT, address _fxStateRootTunnel, uint256 _submitThreshold) returns()
func (_Matic *MaticTransactorSession) Initialize(_nodeOperatorRegistry common.Address, _token common.Address, _dao common.Address, _insurance common.Address, _stakeManager common.Address, _poLidoNFT common.Address, _fxStateRootTunnel common.Address, _submitThreshold *big.Int) (*types.Transaction, error) {
	return _Matic.Contract.Initialize(&_Matic.TransactOpts, _nodeOperatorRegistry, _token, _dao, _insurance, _stakeManager, _poLidoNFT, _fxStateRootTunnel, _submitThreshold)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Matic *MaticTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Matic.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Matic *MaticSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Matic.Contract.RenounceRole(&_Matic.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Matic *MaticTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Matic.Contract.RenounceRole(&_Matic.TransactOpts, role, account)
}

// RequestWithdraw is a paid mutator transaction binding the contract method 0x745400c9.
//
// Solidity: function requestWithdraw(uint256 _amount) returns()
func (_Matic *MaticTransactor) RequestWithdraw(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Matic.contract.Transact(opts, "requestWithdraw", _amount)
}

// RequestWithdraw is a paid mutator transaction binding the contract method 0x745400c9.
//
// Solidity: function requestWithdraw(uint256 _amount) returns()
func (_Matic *MaticSession) RequestWithdraw(_amount *big.Int) (*types.Transaction, error) {
	return _Matic.Contract.RequestWithdraw(&_Matic.TransactOpts, _amount)
}

// RequestWithdraw is a paid mutator transaction binding the contract method 0x745400c9.
//
// Solidity: function requestWithdraw(uint256 _amount) returns()
func (_Matic *MaticTransactorSession) RequestWithdraw(_amount *big.Int) (*types.Transaction, error) {
	return _Matic.Contract.RequestWithdraw(&_Matic.TransactOpts, _amount)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Matic *MaticTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Matic.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Matic *MaticSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Matic.Contract.RevokeRole(&_Matic.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Matic *MaticTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Matic.Contract.RevokeRole(&_Matic.TransactOpts, role, account)
}

// SetDaoAddress is a paid mutator transaction binding the contract method 0x9a3cac6a.
//
// Solidity: function setDaoAddress(address _address) returns()
func (_Matic *MaticTransactor) SetDaoAddress(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Matic.contract.Transact(opts, "setDaoAddress", _address)
}

// SetDaoAddress is a paid mutator transaction binding the contract method 0x9a3cac6a.
//
// Solidity: function setDaoAddress(address _address) returns()
func (_Matic *MaticSession) SetDaoAddress(_address common.Address) (*types.Transaction, error) {
	return _Matic.Contract.SetDaoAddress(&_Matic.TransactOpts, _address)
}

// SetDaoAddress is a paid mutator transaction binding the contract method 0x9a3cac6a.
//
// Solidity: function setDaoAddress(address _address) returns()
func (_Matic *MaticTransactorSession) SetDaoAddress(_address common.Address) (*types.Transaction, error) {
	return _Matic.Contract.SetDaoAddress(&_Matic.TransactOpts, _address)
}

// SetDelegationLowerBound is a paid mutator transaction binding the contract method 0x7682c902.
//
// Solidity: function setDelegationLowerBound(uint256 _delegationLowerBound) returns()
func (_Matic *MaticTransactor) SetDelegationLowerBound(opts *bind.TransactOpts, _delegationLowerBound *big.Int) (*types.Transaction, error) {
	return _Matic.contract.Transact(opts, "setDelegationLowerBound", _delegationLowerBound)
}

// SetDelegationLowerBound is a paid mutator transaction binding the contract method 0x7682c902.
//
// Solidity: function setDelegationLowerBound(uint256 _delegationLowerBound) returns()
func (_Matic *MaticSession) SetDelegationLowerBound(_delegationLowerBound *big.Int) (*types.Transaction, error) {
	return _Matic.Contract.SetDelegationLowerBound(&_Matic.TransactOpts, _delegationLowerBound)
}

// SetDelegationLowerBound is a paid mutator transaction binding the contract method 0x7682c902.
//
// Solidity: function setDelegationLowerBound(uint256 _delegationLowerBound) returns()
func (_Matic *MaticTransactorSession) SetDelegationLowerBound(_delegationLowerBound *big.Int) (*types.Transaction, error) {
	return _Matic.Contract.SetDelegationLowerBound(&_Matic.TransactOpts, _delegationLowerBound)
}

// SetFees is a paid mutator transaction binding the contract method 0xf6794fdb.
//
// Solidity: function setFees(uint8 _daoFee, uint8 _operatorsFee, uint8 _insuranceFee) returns()
func (_Matic *MaticTransactor) SetFees(opts *bind.TransactOpts, _daoFee uint8, _operatorsFee uint8, _insuranceFee uint8) (*types.Transaction, error) {
	return _Matic.contract.Transact(opts, "setFees", _daoFee, _operatorsFee, _insuranceFee)
}

// SetFees is a paid mutator transaction binding the contract method 0xf6794fdb.
//
// Solidity: function setFees(uint8 _daoFee, uint8 _operatorsFee, uint8 _insuranceFee) returns()
func (_Matic *MaticSession) SetFees(_daoFee uint8, _operatorsFee uint8, _insuranceFee uint8) (*types.Transaction, error) {
	return _Matic.Contract.SetFees(&_Matic.TransactOpts, _daoFee, _operatorsFee, _insuranceFee)
}

// SetFees is a paid mutator transaction binding the contract method 0xf6794fdb.
//
// Solidity: function setFees(uint8 _daoFee, uint8 _operatorsFee, uint8 _insuranceFee) returns()
func (_Matic *MaticTransactorSession) SetFees(_daoFee uint8, _operatorsFee uint8, _insuranceFee uint8) (*types.Transaction, error) {
	return _Matic.Contract.SetFees(&_Matic.TransactOpts, _daoFee, _operatorsFee, _insuranceFee)
}

// SetFxStateRootTunnel is a paid mutator transaction binding the contract method 0x70bf9fe9.
//
// Solidity: function setFxStateRootTunnel(address _fxStateRootTunnel) returns()
func (_Matic *MaticTransactor) SetFxStateRootTunnel(opts *bind.TransactOpts, _fxStateRootTunnel common.Address) (*types.Transaction, error) {
	return _Matic.contract.Transact(opts, "setFxStateRootTunnel", _fxStateRootTunnel)
}

// SetFxStateRootTunnel is a paid mutator transaction binding the contract method 0x70bf9fe9.
//
// Solidity: function setFxStateRootTunnel(address _fxStateRootTunnel) returns()
func (_Matic *MaticSession) SetFxStateRootTunnel(_fxStateRootTunnel common.Address) (*types.Transaction, error) {
	return _Matic.Contract.SetFxStateRootTunnel(&_Matic.TransactOpts, _fxStateRootTunnel)
}

// SetFxStateRootTunnel is a paid mutator transaction binding the contract method 0x70bf9fe9.
//
// Solidity: function setFxStateRootTunnel(address _fxStateRootTunnel) returns()
func (_Matic *MaticTransactorSession) SetFxStateRootTunnel(_fxStateRootTunnel common.Address) (*types.Transaction, error) {
	return _Matic.Contract.SetFxStateRootTunnel(&_Matic.TransactOpts, _fxStateRootTunnel)
}

// SetInsuranceAddress is a paid mutator transaction binding the contract method 0xbb208f55.
//
// Solidity: function setInsuranceAddress(address _address) returns()
func (_Matic *MaticTransactor) SetInsuranceAddress(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Matic.contract.Transact(opts, "setInsuranceAddress", _address)
}

// SetInsuranceAddress is a paid mutator transaction binding the contract method 0xbb208f55.
//
// Solidity: function setInsuranceAddress(address _address) returns()
func (_Matic *MaticSession) SetInsuranceAddress(_address common.Address) (*types.Transaction, error) {
	return _Matic.Contract.SetInsuranceAddress(&_Matic.TransactOpts, _address)
}

// SetInsuranceAddress is a paid mutator transaction binding the contract method 0xbb208f55.
//
// Solidity: function setInsuranceAddress(address _address) returns()
func (_Matic *MaticTransactorSession) SetInsuranceAddress(_address common.Address) (*types.Transaction, error) {
	return _Matic.Contract.SetInsuranceAddress(&_Matic.TransactOpts, _address)
}

// SetNodeOperatorRegistryAddress is a paid mutator transaction binding the contract method 0x0f2b2639.
//
// Solidity: function setNodeOperatorRegistryAddress(address _address) returns()
func (_Matic *MaticTransactor) SetNodeOperatorRegistryAddress(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Matic.contract.Transact(opts, "setNodeOperatorRegistryAddress", _address)
}

// SetNodeOperatorRegistryAddress is a paid mutator transaction binding the contract method 0x0f2b2639.
//
// Solidity: function setNodeOperatorRegistryAddress(address _address) returns()
func (_Matic *MaticSession) SetNodeOperatorRegistryAddress(_address common.Address) (*types.Transaction, error) {
	return _Matic.Contract.SetNodeOperatorRegistryAddress(&_Matic.TransactOpts, _address)
}

// SetNodeOperatorRegistryAddress is a paid mutator transaction binding the contract method 0x0f2b2639.
//
// Solidity: function setNodeOperatorRegistryAddress(address _address) returns()
func (_Matic *MaticTransactorSession) SetNodeOperatorRegistryAddress(_address common.Address) (*types.Transaction, error) {
	return _Matic.Contract.SetNodeOperatorRegistryAddress(&_Matic.TransactOpts, _address)
}

// SetPoLidoNFT is a paid mutator transaction binding the contract method 0x15539d3f.
//
// Solidity: function setPoLidoNFT(address _poLidoNFT) returns()
func (_Matic *MaticTransactor) SetPoLidoNFT(opts *bind.TransactOpts, _poLidoNFT common.Address) (*types.Transaction, error) {
	return _Matic.contract.Transact(opts, "setPoLidoNFT", _poLidoNFT)
}

// SetPoLidoNFT is a paid mutator transaction binding the contract method 0x15539d3f.
//
// Solidity: function setPoLidoNFT(address _poLidoNFT) returns()
func (_Matic *MaticSession) SetPoLidoNFT(_poLidoNFT common.Address) (*types.Transaction, error) {
	return _Matic.Contract.SetPoLidoNFT(&_Matic.TransactOpts, _poLidoNFT)
}

// SetPoLidoNFT is a paid mutator transaction binding the contract method 0x15539d3f.
//
// Solidity: function setPoLidoNFT(address _poLidoNFT) returns()
func (_Matic *MaticTransactorSession) SetPoLidoNFT(_poLidoNFT common.Address) (*types.Transaction, error) {
	return _Matic.Contract.SetPoLidoNFT(&_Matic.TransactOpts, _poLidoNFT)
}

// SetRewardDistributionLowerBound is a paid mutator transaction binding the contract method 0x3b573c4a.
//
// Solidity: function setRewardDistributionLowerBound(uint256 _rewardDistributionLowerBound) returns()
func (_Matic *MaticTransactor) SetRewardDistributionLowerBound(opts *bind.TransactOpts, _rewardDistributionLowerBound *big.Int) (*types.Transaction, error) {
	return _Matic.contract.Transact(opts, "setRewardDistributionLowerBound", _rewardDistributionLowerBound)
}

// SetRewardDistributionLowerBound is a paid mutator transaction binding the contract method 0x3b573c4a.
//
// Solidity: function setRewardDistributionLowerBound(uint256 _rewardDistributionLowerBound) returns()
func (_Matic *MaticSession) SetRewardDistributionLowerBound(_rewardDistributionLowerBound *big.Int) (*types.Transaction, error) {
	return _Matic.Contract.SetRewardDistributionLowerBound(&_Matic.TransactOpts, _rewardDistributionLowerBound)
}

// SetRewardDistributionLowerBound is a paid mutator transaction binding the contract method 0x3b573c4a.
//
// Solidity: function setRewardDistributionLowerBound(uint256 _rewardDistributionLowerBound) returns()
func (_Matic *MaticTransactorSession) SetRewardDistributionLowerBound(_rewardDistributionLowerBound *big.Int) (*types.Transaction, error) {
	return _Matic.Contract.SetRewardDistributionLowerBound(&_Matic.TransactOpts, _rewardDistributionLowerBound)
}

// SetSubmitThreshold is a paid mutator transaction binding the contract method 0xee319c21.
//
// Solidity: function setSubmitThreshold(uint256 _submitThreshold) returns()
func (_Matic *MaticTransactor) SetSubmitThreshold(opts *bind.TransactOpts, _submitThreshold *big.Int) (*types.Transaction, error) {
	return _Matic.contract.Transact(opts, "setSubmitThreshold", _submitThreshold)
}

// SetSubmitThreshold is a paid mutator transaction binding the contract method 0xee319c21.
//
// Solidity: function setSubmitThreshold(uint256 _submitThreshold) returns()
func (_Matic *MaticSession) SetSubmitThreshold(_submitThreshold *big.Int) (*types.Transaction, error) {
	return _Matic.Contract.SetSubmitThreshold(&_Matic.TransactOpts, _submitThreshold)
}

// SetSubmitThreshold is a paid mutator transaction binding the contract method 0xee319c21.
//
// Solidity: function setSubmitThreshold(uint256 _submitThreshold) returns()
func (_Matic *MaticTransactorSession) SetSubmitThreshold(_submitThreshold *big.Int) (*types.Transaction, error) {
	return _Matic.Contract.SetSubmitThreshold(&_Matic.TransactOpts, _submitThreshold)
}

// SetVersion is a paid mutator transaction binding the contract method 0x788bc78c.
//
// Solidity: function setVersion(string _version) returns()
func (_Matic *MaticTransactor) SetVersion(opts *bind.TransactOpts, _version string) (*types.Transaction, error) {
	return _Matic.contract.Transact(opts, "setVersion", _version)
}

// SetVersion is a paid mutator transaction binding the contract method 0x788bc78c.
//
// Solidity: function setVersion(string _version) returns()
func (_Matic *MaticSession) SetVersion(_version string) (*types.Transaction, error) {
	return _Matic.Contract.SetVersion(&_Matic.TransactOpts, _version)
}

// SetVersion is a paid mutator transaction binding the contract method 0x788bc78c.
//
// Solidity: function setVersion(string _version) returns()
func (_Matic *MaticTransactorSession) SetVersion(_version string) (*types.Transaction, error) {
	return _Matic.Contract.SetVersion(&_Matic.TransactOpts, _version)
}

// Submit is a paid mutator transaction binding the contract method 0xea99c2a6.
//
// Solidity: function submit(uint256 _amount) returns(uint256)
func (_Matic *MaticTransactor) Submit(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Matic.contract.Transact(opts, "submit", _amount)
}

// Submit is a paid mutator transaction binding the contract method 0xea99c2a6.
//
// Solidity: function submit(uint256 _amount) returns(uint256)
func (_Matic *MaticSession) Submit(_amount *big.Int) (*types.Transaction, error) {
	return _Matic.Contract.Submit(&_Matic.TransactOpts, _amount)
}

// Submit is a paid mutator transaction binding the contract method 0xea99c2a6.
//
// Solidity: function submit(uint256 _amount) returns(uint256)
func (_Matic *MaticTransactorSession) Submit(_amount *big.Int) (*types.Transaction, error) {
	return _Matic.Contract.Submit(&_Matic.TransactOpts, _amount)
}

// TogglePause is a paid mutator transaction binding the contract method 0xc4ae3168.
//
// Solidity: function togglePause() returns()
func (_Matic *MaticTransactor) TogglePause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Matic.contract.Transact(opts, "togglePause")
}

// TogglePause is a paid mutator transaction binding the contract method 0xc4ae3168.
//
// Solidity: function togglePause() returns()
func (_Matic *MaticSession) TogglePause() (*types.Transaction, error) {
	return _Matic.Contract.TogglePause(&_Matic.TransactOpts)
}

// TogglePause is a paid mutator transaction binding the contract method 0xc4ae3168.
//
// Solidity: function togglePause() returns()
func (_Matic *MaticTransactorSession) TogglePause() (*types.Transaction, error) {
	return _Matic.Contract.TogglePause(&_Matic.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Matic *MaticTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Matic.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Matic *MaticSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Matic.Contract.Transfer(&_Matic.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Matic *MaticTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Matic.Contract.Transfer(&_Matic.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Matic *MaticTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Matic.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Matic *MaticSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Matic.Contract.TransferFrom(&_Matic.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Matic *MaticTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Matic.Contract.TransferFrom(&_Matic.TransactOpts, sender, recipient, amount)
}

// WithdrawTotalDelegated is a paid mutator transaction binding the contract method 0xc75e7832.
//
// Solidity: function withdrawTotalDelegated(address _validatorShare) returns()
func (_Matic *MaticTransactor) WithdrawTotalDelegated(opts *bind.TransactOpts, _validatorShare common.Address) (*types.Transaction, error) {
	return _Matic.contract.Transact(opts, "withdrawTotalDelegated", _validatorShare)
}

// WithdrawTotalDelegated is a paid mutator transaction binding the contract method 0xc75e7832.
//
// Solidity: function withdrawTotalDelegated(address _validatorShare) returns()
func (_Matic *MaticSession) WithdrawTotalDelegated(_validatorShare common.Address) (*types.Transaction, error) {
	return _Matic.Contract.WithdrawTotalDelegated(&_Matic.TransactOpts, _validatorShare)
}

// WithdrawTotalDelegated is a paid mutator transaction binding the contract method 0xc75e7832.
//
// Solidity: function withdrawTotalDelegated(address _validatorShare) returns()
func (_Matic *MaticTransactorSession) WithdrawTotalDelegated(_validatorShare common.Address) (*types.Transaction, error) {
	return _Matic.Contract.WithdrawTotalDelegated(&_Matic.TransactOpts, _validatorShare)
}

// MaticApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Matic contract.
type MaticApprovalIterator struct {
	Event *MaticApproval // Event containing the contract specifics and raw log

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
func (it *MaticApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MaticApproval)
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
		it.Event = new(MaticApproval)
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
func (it *MaticApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MaticApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MaticApproval represents a Approval event raised by the Matic contract.
type MaticApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Matic *MaticFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*MaticApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Matic.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &MaticApprovalIterator{contract: _Matic.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Matic *MaticFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MaticApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Matic.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MaticApproval)
				if err := _Matic.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Matic *MaticFilterer) ParseApproval(log types.Log) (*MaticApproval, error) {
	event := new(MaticApproval)
	if err := _Matic.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MaticClaimTokensEventIterator is returned from FilterClaimTokensEvent and is used to iterate over the raw logs and unpacked data for ClaimTokensEvent events raised by the Matic contract.
type MaticClaimTokensEventIterator struct {
	Event *MaticClaimTokensEvent // Event containing the contract specifics and raw log

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
func (it *MaticClaimTokensEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MaticClaimTokensEvent)
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
		it.Event = new(MaticClaimTokensEvent)
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
func (it *MaticClaimTokensEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MaticClaimTokensEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MaticClaimTokensEvent represents a ClaimTokensEvent event raised by the Matic contract.
type MaticClaimTokensEvent struct {
	From          common.Address
	Id            *big.Int
	AmountClaimed *big.Int
	AmountBurned  *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterClaimTokensEvent is a free log retrieval operation binding the contract event 0xaca94a3466fab333b79851ab29b0715612740e4ae0d891ef8e9bd2a1bf5e24dd.
//
// Solidity: event ClaimTokensEvent(address indexed _from, uint256 indexed _id, uint256 indexed _amountClaimed, uint256 _amountBurned)
func (_Matic *MaticFilterer) FilterClaimTokensEvent(opts *bind.FilterOpts, _from []common.Address, _id []*big.Int, _amountClaimed []*big.Int) (*MaticClaimTokensEventIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}
	var _amountClaimedRule []interface{}
	for _, _amountClaimedItem := range _amountClaimed {
		_amountClaimedRule = append(_amountClaimedRule, _amountClaimedItem)
	}

	logs, sub, err := _Matic.contract.FilterLogs(opts, "ClaimTokensEvent", _fromRule, _idRule, _amountClaimedRule)
	if err != nil {
		return nil, err
	}
	return &MaticClaimTokensEventIterator{contract: _Matic.contract, event: "ClaimTokensEvent", logs: logs, sub: sub}, nil
}

// WatchClaimTokensEvent is a free log subscription operation binding the contract event 0xaca94a3466fab333b79851ab29b0715612740e4ae0d891ef8e9bd2a1bf5e24dd.
//
// Solidity: event ClaimTokensEvent(address indexed _from, uint256 indexed _id, uint256 indexed _amountClaimed, uint256 _amountBurned)
func (_Matic *MaticFilterer) WatchClaimTokensEvent(opts *bind.WatchOpts, sink chan<- *MaticClaimTokensEvent, _from []common.Address, _id []*big.Int, _amountClaimed []*big.Int) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}
	var _amountClaimedRule []interface{}
	for _, _amountClaimedItem := range _amountClaimed {
		_amountClaimedRule = append(_amountClaimedRule, _amountClaimedItem)
	}

	logs, sub, err := _Matic.contract.WatchLogs(opts, "ClaimTokensEvent", _fromRule, _idRule, _amountClaimedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MaticClaimTokensEvent)
				if err := _Matic.contract.UnpackLog(event, "ClaimTokensEvent", log); err != nil {
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

// ParseClaimTokensEvent is a log parse operation binding the contract event 0xaca94a3466fab333b79851ab29b0715612740e4ae0d891ef8e9bd2a1bf5e24dd.
//
// Solidity: event ClaimTokensEvent(address indexed _from, uint256 indexed _id, uint256 indexed _amountClaimed, uint256 _amountBurned)
func (_Matic *MaticFilterer) ParseClaimTokensEvent(log types.Log) (*MaticClaimTokensEvent, error) {
	event := new(MaticClaimTokensEvent)
	if err := _Matic.contract.UnpackLog(event, "ClaimTokensEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MaticClaimTotalDelegatedEventIterator is returned from FilterClaimTotalDelegatedEvent and is used to iterate over the raw logs and unpacked data for ClaimTotalDelegatedEvent events raised by the Matic contract.
type MaticClaimTotalDelegatedEventIterator struct {
	Event *MaticClaimTotalDelegatedEvent // Event containing the contract specifics and raw log

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
func (it *MaticClaimTotalDelegatedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MaticClaimTotalDelegatedEvent)
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
		it.Event = new(MaticClaimTotalDelegatedEvent)
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
func (it *MaticClaimTotalDelegatedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MaticClaimTotalDelegatedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MaticClaimTotalDelegatedEvent represents a ClaimTotalDelegatedEvent event raised by the Matic contract.
type MaticClaimTotalDelegatedEvent struct {
	From          common.Address
	AmountClaimed *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterClaimTotalDelegatedEvent is a free log retrieval operation binding the contract event 0x4c42a3bec298a4d82d41b7a540d8ebc22d91ee8a61459bce23849ff470d31dea.
//
// Solidity: event ClaimTotalDelegatedEvent(address indexed _from, uint256 indexed _amountClaimed)
func (_Matic *MaticFilterer) FilterClaimTotalDelegatedEvent(opts *bind.FilterOpts, _from []common.Address, _amountClaimed []*big.Int) (*MaticClaimTotalDelegatedEventIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _amountClaimedRule []interface{}
	for _, _amountClaimedItem := range _amountClaimed {
		_amountClaimedRule = append(_amountClaimedRule, _amountClaimedItem)
	}

	logs, sub, err := _Matic.contract.FilterLogs(opts, "ClaimTotalDelegatedEvent", _fromRule, _amountClaimedRule)
	if err != nil {
		return nil, err
	}
	return &MaticClaimTotalDelegatedEventIterator{contract: _Matic.contract, event: "ClaimTotalDelegatedEvent", logs: logs, sub: sub}, nil
}

// WatchClaimTotalDelegatedEvent is a free log subscription operation binding the contract event 0x4c42a3bec298a4d82d41b7a540d8ebc22d91ee8a61459bce23849ff470d31dea.
//
// Solidity: event ClaimTotalDelegatedEvent(address indexed _from, uint256 indexed _amountClaimed)
func (_Matic *MaticFilterer) WatchClaimTotalDelegatedEvent(opts *bind.WatchOpts, sink chan<- *MaticClaimTotalDelegatedEvent, _from []common.Address, _amountClaimed []*big.Int) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _amountClaimedRule []interface{}
	for _, _amountClaimedItem := range _amountClaimed {
		_amountClaimedRule = append(_amountClaimedRule, _amountClaimedItem)
	}

	logs, sub, err := _Matic.contract.WatchLogs(opts, "ClaimTotalDelegatedEvent", _fromRule, _amountClaimedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MaticClaimTotalDelegatedEvent)
				if err := _Matic.contract.UnpackLog(event, "ClaimTotalDelegatedEvent", log); err != nil {
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

// ParseClaimTotalDelegatedEvent is a log parse operation binding the contract event 0x4c42a3bec298a4d82d41b7a540d8ebc22d91ee8a61459bce23849ff470d31dea.
//
// Solidity: event ClaimTotalDelegatedEvent(address indexed _from, uint256 indexed _amountClaimed)
func (_Matic *MaticFilterer) ParseClaimTotalDelegatedEvent(log types.Log) (*MaticClaimTotalDelegatedEvent, error) {
	event := new(MaticClaimTotalDelegatedEvent)
	if err := _Matic.contract.UnpackLog(event, "ClaimTotalDelegatedEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MaticDelegateEventIterator is returned from FilterDelegateEvent and is used to iterate over the raw logs and unpacked data for DelegateEvent events raised by the Matic contract.
type MaticDelegateEventIterator struct {
	Event *MaticDelegateEvent // Event containing the contract specifics and raw log

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
func (it *MaticDelegateEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MaticDelegateEvent)
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
		it.Event = new(MaticDelegateEvent)
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
func (it *MaticDelegateEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MaticDelegateEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MaticDelegateEvent represents a DelegateEvent event raised by the Matic contract.
type MaticDelegateEvent struct {
	AmountDelegated *big.Int
	Remainder       *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDelegateEvent is a free log retrieval operation binding the contract event 0x421adba60af7a6b11679e2ac133b1bc91d3de91d56866ec19703d9d60cf950c8.
//
// Solidity: event DelegateEvent(uint256 indexed _amountDelegated, uint256 indexed _remainder)
func (_Matic *MaticFilterer) FilterDelegateEvent(opts *bind.FilterOpts, _amountDelegated []*big.Int, _remainder []*big.Int) (*MaticDelegateEventIterator, error) {

	var _amountDelegatedRule []interface{}
	for _, _amountDelegatedItem := range _amountDelegated {
		_amountDelegatedRule = append(_amountDelegatedRule, _amountDelegatedItem)
	}
	var _remainderRule []interface{}
	for _, _remainderItem := range _remainder {
		_remainderRule = append(_remainderRule, _remainderItem)
	}

	logs, sub, err := _Matic.contract.FilterLogs(opts, "DelegateEvent", _amountDelegatedRule, _remainderRule)
	if err != nil {
		return nil, err
	}
	return &MaticDelegateEventIterator{contract: _Matic.contract, event: "DelegateEvent", logs: logs, sub: sub}, nil
}

// WatchDelegateEvent is a free log subscription operation binding the contract event 0x421adba60af7a6b11679e2ac133b1bc91d3de91d56866ec19703d9d60cf950c8.
//
// Solidity: event DelegateEvent(uint256 indexed _amountDelegated, uint256 indexed _remainder)
func (_Matic *MaticFilterer) WatchDelegateEvent(opts *bind.WatchOpts, sink chan<- *MaticDelegateEvent, _amountDelegated []*big.Int, _remainder []*big.Int) (event.Subscription, error) {

	var _amountDelegatedRule []interface{}
	for _, _amountDelegatedItem := range _amountDelegated {
		_amountDelegatedRule = append(_amountDelegatedRule, _amountDelegatedItem)
	}
	var _remainderRule []interface{}
	for _, _remainderItem := range _remainder {
		_remainderRule = append(_remainderRule, _remainderItem)
	}

	logs, sub, err := _Matic.contract.WatchLogs(opts, "DelegateEvent", _amountDelegatedRule, _remainderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MaticDelegateEvent)
				if err := _Matic.contract.UnpackLog(event, "DelegateEvent", log); err != nil {
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

// ParseDelegateEvent is a log parse operation binding the contract event 0x421adba60af7a6b11679e2ac133b1bc91d3de91d56866ec19703d9d60cf950c8.
//
// Solidity: event DelegateEvent(uint256 indexed _amountDelegated, uint256 indexed _remainder)
func (_Matic *MaticFilterer) ParseDelegateEvent(log types.Log) (*MaticDelegateEvent, error) {
	event := new(MaticDelegateEvent)
	if err := _Matic.contract.UnpackLog(event, "DelegateEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MaticDistributeRewardsEventIterator is returned from FilterDistributeRewardsEvent and is used to iterate over the raw logs and unpacked data for DistributeRewardsEvent events raised by the Matic contract.
type MaticDistributeRewardsEventIterator struct {
	Event *MaticDistributeRewardsEvent // Event containing the contract specifics and raw log

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
func (it *MaticDistributeRewardsEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MaticDistributeRewardsEvent)
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
		it.Event = new(MaticDistributeRewardsEvent)
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
func (it *MaticDistributeRewardsEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MaticDistributeRewardsEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MaticDistributeRewardsEvent represents a DistributeRewardsEvent event raised by the Matic contract.
type MaticDistributeRewardsEvent struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDistributeRewardsEvent is a free log retrieval operation binding the contract event 0x4e3c6a1e602996ae70905ac6165ed2434753246e3bfa52b6ca6852b40e2d4408.
//
// Solidity: event DistributeRewardsEvent(uint256 indexed _amount)
func (_Matic *MaticFilterer) FilterDistributeRewardsEvent(opts *bind.FilterOpts, _amount []*big.Int) (*MaticDistributeRewardsEventIterator, error) {

	var _amountRule []interface{}
	for _, _amountItem := range _amount {
		_amountRule = append(_amountRule, _amountItem)
	}

	logs, sub, err := _Matic.contract.FilterLogs(opts, "DistributeRewardsEvent", _amountRule)
	if err != nil {
		return nil, err
	}
	return &MaticDistributeRewardsEventIterator{contract: _Matic.contract, event: "DistributeRewardsEvent", logs: logs, sub: sub}, nil
}

// WatchDistributeRewardsEvent is a free log subscription operation binding the contract event 0x4e3c6a1e602996ae70905ac6165ed2434753246e3bfa52b6ca6852b40e2d4408.
//
// Solidity: event DistributeRewardsEvent(uint256 indexed _amount)
func (_Matic *MaticFilterer) WatchDistributeRewardsEvent(opts *bind.WatchOpts, sink chan<- *MaticDistributeRewardsEvent, _amount []*big.Int) (event.Subscription, error) {

	var _amountRule []interface{}
	for _, _amountItem := range _amount {
		_amountRule = append(_amountRule, _amountItem)
	}

	logs, sub, err := _Matic.contract.WatchLogs(opts, "DistributeRewardsEvent", _amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MaticDistributeRewardsEvent)
				if err := _Matic.contract.UnpackLog(event, "DistributeRewardsEvent", log); err != nil {
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

// ParseDistributeRewardsEvent is a log parse operation binding the contract event 0x4e3c6a1e602996ae70905ac6165ed2434753246e3bfa52b6ca6852b40e2d4408.
//
// Solidity: event DistributeRewardsEvent(uint256 indexed _amount)
func (_Matic *MaticFilterer) ParseDistributeRewardsEvent(log types.Log) (*MaticDistributeRewardsEvent, error) {
	event := new(MaticDistributeRewardsEvent)
	if err := _Matic.contract.UnpackLog(event, "DistributeRewardsEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MaticPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Matic contract.
type MaticPausedIterator struct {
	Event *MaticPaused // Event containing the contract specifics and raw log

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
func (it *MaticPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MaticPaused)
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
		it.Event = new(MaticPaused)
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
func (it *MaticPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MaticPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MaticPaused represents a Paused event raised by the Matic contract.
type MaticPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Matic *MaticFilterer) FilterPaused(opts *bind.FilterOpts) (*MaticPausedIterator, error) {

	logs, sub, err := _Matic.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &MaticPausedIterator{contract: _Matic.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Matic *MaticFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *MaticPaused) (event.Subscription, error) {

	logs, sub, err := _Matic.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MaticPaused)
				if err := _Matic.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Matic *MaticFilterer) ParsePaused(log types.Log) (*MaticPaused, error) {
	event := new(MaticPaused)
	if err := _Matic.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MaticRequestWithdrawEventIterator is returned from FilterRequestWithdrawEvent and is used to iterate over the raw logs and unpacked data for RequestWithdrawEvent events raised by the Matic contract.
type MaticRequestWithdrawEventIterator struct {
	Event *MaticRequestWithdrawEvent // Event containing the contract specifics and raw log

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
func (it *MaticRequestWithdrawEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MaticRequestWithdrawEvent)
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
		it.Event = new(MaticRequestWithdrawEvent)
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
func (it *MaticRequestWithdrawEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MaticRequestWithdrawEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MaticRequestWithdrawEvent represents a RequestWithdrawEvent event raised by the Matic contract.
type MaticRequestWithdrawEvent struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRequestWithdrawEvent is a free log retrieval operation binding the contract event 0x8a8169c8a646f81d6d6ad8ed0cf560361c75cb37a74656f2487d0fa9bfcb0844.
//
// Solidity: event RequestWithdrawEvent(address indexed _from, uint256 indexed _amount)
func (_Matic *MaticFilterer) FilterRequestWithdrawEvent(opts *bind.FilterOpts, _from []common.Address, _amount []*big.Int) (*MaticRequestWithdrawEventIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _amountRule []interface{}
	for _, _amountItem := range _amount {
		_amountRule = append(_amountRule, _amountItem)
	}

	logs, sub, err := _Matic.contract.FilterLogs(opts, "RequestWithdrawEvent", _fromRule, _amountRule)
	if err != nil {
		return nil, err
	}
	return &MaticRequestWithdrawEventIterator{contract: _Matic.contract, event: "RequestWithdrawEvent", logs: logs, sub: sub}, nil
}

// WatchRequestWithdrawEvent is a free log subscription operation binding the contract event 0x8a8169c8a646f81d6d6ad8ed0cf560361c75cb37a74656f2487d0fa9bfcb0844.
//
// Solidity: event RequestWithdrawEvent(address indexed _from, uint256 indexed _amount)
func (_Matic *MaticFilterer) WatchRequestWithdrawEvent(opts *bind.WatchOpts, sink chan<- *MaticRequestWithdrawEvent, _from []common.Address, _amount []*big.Int) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _amountRule []interface{}
	for _, _amountItem := range _amount {
		_amountRule = append(_amountRule, _amountItem)
	}

	logs, sub, err := _Matic.contract.WatchLogs(opts, "RequestWithdrawEvent", _fromRule, _amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MaticRequestWithdrawEvent)
				if err := _Matic.contract.UnpackLog(event, "RequestWithdrawEvent", log); err != nil {
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

// ParseRequestWithdrawEvent is a log parse operation binding the contract event 0x8a8169c8a646f81d6d6ad8ed0cf560361c75cb37a74656f2487d0fa9bfcb0844.
//
// Solidity: event RequestWithdrawEvent(address indexed _from, uint256 indexed _amount)
func (_Matic *MaticFilterer) ParseRequestWithdrawEvent(log types.Log) (*MaticRequestWithdrawEvent, error) {
	event := new(MaticRequestWithdrawEvent)
	if err := _Matic.contract.UnpackLog(event, "RequestWithdrawEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MaticRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Matic contract.
type MaticRoleAdminChangedIterator struct {
	Event *MaticRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *MaticRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MaticRoleAdminChanged)
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
		it.Event = new(MaticRoleAdminChanged)
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
func (it *MaticRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MaticRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MaticRoleAdminChanged represents a RoleAdminChanged event raised by the Matic contract.
type MaticRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Matic *MaticFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*MaticRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Matic.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &MaticRoleAdminChangedIterator{contract: _Matic.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Matic *MaticFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *MaticRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Matic.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MaticRoleAdminChanged)
				if err := _Matic.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Matic *MaticFilterer) ParseRoleAdminChanged(log types.Log) (*MaticRoleAdminChanged, error) {
	event := new(MaticRoleAdminChanged)
	if err := _Matic.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MaticRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Matic contract.
type MaticRoleGrantedIterator struct {
	Event *MaticRoleGranted // Event containing the contract specifics and raw log

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
func (it *MaticRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MaticRoleGranted)
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
		it.Event = new(MaticRoleGranted)
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
func (it *MaticRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MaticRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MaticRoleGranted represents a RoleGranted event raised by the Matic contract.
type MaticRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Matic *MaticFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*MaticRoleGrantedIterator, error) {

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

	logs, sub, err := _Matic.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &MaticRoleGrantedIterator{contract: _Matic.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Matic *MaticFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *MaticRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Matic.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MaticRoleGranted)
				if err := _Matic.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Matic *MaticFilterer) ParseRoleGranted(log types.Log) (*MaticRoleGranted, error) {
	event := new(MaticRoleGranted)
	if err := _Matic.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MaticRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Matic contract.
type MaticRoleRevokedIterator struct {
	Event *MaticRoleRevoked // Event containing the contract specifics and raw log

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
func (it *MaticRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MaticRoleRevoked)
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
		it.Event = new(MaticRoleRevoked)
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
func (it *MaticRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MaticRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MaticRoleRevoked represents a RoleRevoked event raised by the Matic contract.
type MaticRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Matic *MaticFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*MaticRoleRevokedIterator, error) {

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

	logs, sub, err := _Matic.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &MaticRoleRevokedIterator{contract: _Matic.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Matic *MaticFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *MaticRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Matic.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MaticRoleRevoked)
				if err := _Matic.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Matic *MaticFilterer) ParseRoleRevoked(log types.Log) (*MaticRoleRevoked, error) {
	event := new(MaticRoleRevoked)
	if err := _Matic.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MaticSubmitEventIterator is returned from FilterSubmitEvent and is used to iterate over the raw logs and unpacked data for SubmitEvent events raised by the Matic contract.
type MaticSubmitEventIterator struct {
	Event *MaticSubmitEvent // Event containing the contract specifics and raw log

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
func (it *MaticSubmitEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MaticSubmitEvent)
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
		it.Event = new(MaticSubmitEvent)
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
func (it *MaticSubmitEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MaticSubmitEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MaticSubmitEvent represents a SubmitEvent event raised by the Matic contract.
type MaticSubmitEvent struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSubmitEvent is a free log retrieval operation binding the contract event 0x8cab5a17f7d817d11abfe3fb3f8dd67646d2643cb4222e5354bde1f65ef6c44c.
//
// Solidity: event SubmitEvent(address indexed _from, uint256 indexed _amount)
func (_Matic *MaticFilterer) FilterSubmitEvent(opts *bind.FilterOpts, _from []common.Address, _amount []*big.Int) (*MaticSubmitEventIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _amountRule []interface{}
	for _, _amountItem := range _amount {
		_amountRule = append(_amountRule, _amountItem)
	}

	logs, sub, err := _Matic.contract.FilterLogs(opts, "SubmitEvent", _fromRule, _amountRule)
	if err != nil {
		return nil, err
	}
	return &MaticSubmitEventIterator{contract: _Matic.contract, event: "SubmitEvent", logs: logs, sub: sub}, nil
}

// WatchSubmitEvent is a free log subscription operation binding the contract event 0x8cab5a17f7d817d11abfe3fb3f8dd67646d2643cb4222e5354bde1f65ef6c44c.
//
// Solidity: event SubmitEvent(address indexed _from, uint256 indexed _amount)
func (_Matic *MaticFilterer) WatchSubmitEvent(opts *bind.WatchOpts, sink chan<- *MaticSubmitEvent, _from []common.Address, _amount []*big.Int) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _amountRule []interface{}
	for _, _amountItem := range _amount {
		_amountRule = append(_amountRule, _amountItem)
	}

	logs, sub, err := _Matic.contract.WatchLogs(opts, "SubmitEvent", _fromRule, _amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MaticSubmitEvent)
				if err := _Matic.contract.UnpackLog(event, "SubmitEvent", log); err != nil {
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

// ParseSubmitEvent is a log parse operation binding the contract event 0x8cab5a17f7d817d11abfe3fb3f8dd67646d2643cb4222e5354bde1f65ef6c44c.
//
// Solidity: event SubmitEvent(address indexed _from, uint256 indexed _amount)
func (_Matic *MaticFilterer) ParseSubmitEvent(log types.Log) (*MaticSubmitEvent, error) {
	event := new(MaticSubmitEvent)
	if err := _Matic.contract.UnpackLog(event, "SubmitEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MaticTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Matic contract.
type MaticTransferIterator struct {
	Event *MaticTransfer // Event containing the contract specifics and raw log

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
func (it *MaticTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MaticTransfer)
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
		it.Event = new(MaticTransfer)
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
func (it *MaticTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MaticTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MaticTransfer represents a Transfer event raised by the Matic contract.
type MaticTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Matic *MaticFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MaticTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Matic.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MaticTransferIterator{contract: _Matic.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Matic *MaticFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MaticTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Matic.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MaticTransfer)
				if err := _Matic.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Matic *MaticFilterer) ParseTransfer(log types.Log) (*MaticTransfer, error) {
	event := new(MaticTransfer)
	if err := _Matic.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MaticUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Matic contract.
type MaticUnpausedIterator struct {
	Event *MaticUnpaused // Event containing the contract specifics and raw log

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
func (it *MaticUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MaticUnpaused)
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
		it.Event = new(MaticUnpaused)
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
func (it *MaticUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MaticUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MaticUnpaused represents a Unpaused event raised by the Matic contract.
type MaticUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Matic *MaticFilterer) FilterUnpaused(opts *bind.FilterOpts) (*MaticUnpausedIterator, error) {

	logs, sub, err := _Matic.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &MaticUnpausedIterator{contract: _Matic.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Matic *MaticFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *MaticUnpaused) (event.Subscription, error) {

	logs, sub, err := _Matic.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MaticUnpaused)
				if err := _Matic.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Matic *MaticFilterer) ParseUnpaused(log types.Log) (*MaticUnpaused, error) {
	event := new(MaticUnpaused)
	if err := _Matic.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MaticWithdrawTotalDelegatedEventIterator is returned from FilterWithdrawTotalDelegatedEvent and is used to iterate over the raw logs and unpacked data for WithdrawTotalDelegatedEvent events raised by the Matic contract.
type MaticWithdrawTotalDelegatedEventIterator struct {
	Event *MaticWithdrawTotalDelegatedEvent // Event containing the contract specifics and raw log

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
func (it *MaticWithdrawTotalDelegatedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MaticWithdrawTotalDelegatedEvent)
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
		it.Event = new(MaticWithdrawTotalDelegatedEvent)
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
func (it *MaticWithdrawTotalDelegatedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MaticWithdrawTotalDelegatedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MaticWithdrawTotalDelegatedEvent represents a WithdrawTotalDelegatedEvent event raised by the Matic contract.
type MaticWithdrawTotalDelegatedEvent struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawTotalDelegatedEvent is a free log retrieval operation binding the contract event 0x65fcdf1cdc99352d178d6d953d52e01307cde7a592027b09c9e1d9ac8eb09ab7.
//
// Solidity: event WithdrawTotalDelegatedEvent(address indexed _from, uint256 indexed _amount)
func (_Matic *MaticFilterer) FilterWithdrawTotalDelegatedEvent(opts *bind.FilterOpts, _from []common.Address, _amount []*big.Int) (*MaticWithdrawTotalDelegatedEventIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _amountRule []interface{}
	for _, _amountItem := range _amount {
		_amountRule = append(_amountRule, _amountItem)
	}

	logs, sub, err := _Matic.contract.FilterLogs(opts, "WithdrawTotalDelegatedEvent", _fromRule, _amountRule)
	if err != nil {
		return nil, err
	}
	return &MaticWithdrawTotalDelegatedEventIterator{contract: _Matic.contract, event: "WithdrawTotalDelegatedEvent", logs: logs, sub: sub}, nil
}

// WatchWithdrawTotalDelegatedEvent is a free log subscription operation binding the contract event 0x65fcdf1cdc99352d178d6d953d52e01307cde7a592027b09c9e1d9ac8eb09ab7.
//
// Solidity: event WithdrawTotalDelegatedEvent(address indexed _from, uint256 indexed _amount)
func (_Matic *MaticFilterer) WatchWithdrawTotalDelegatedEvent(opts *bind.WatchOpts, sink chan<- *MaticWithdrawTotalDelegatedEvent, _from []common.Address, _amount []*big.Int) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _amountRule []interface{}
	for _, _amountItem := range _amount {
		_amountRule = append(_amountRule, _amountItem)
	}

	logs, sub, err := _Matic.contract.WatchLogs(opts, "WithdrawTotalDelegatedEvent", _fromRule, _amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MaticWithdrawTotalDelegatedEvent)
				if err := _Matic.contract.UnpackLog(event, "WithdrawTotalDelegatedEvent", log); err != nil {
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

// ParseWithdrawTotalDelegatedEvent is a log parse operation binding the contract event 0x65fcdf1cdc99352d178d6d953d52e01307cde7a592027b09c9e1d9ac8eb09ab7.
//
// Solidity: event WithdrawTotalDelegatedEvent(address indexed _from, uint256 indexed _amount)
func (_Matic *MaticFilterer) ParseWithdrawTotalDelegatedEvent(log types.Log) (*MaticWithdrawTotalDelegatedEvent, error) {
	event := new(MaticWithdrawTotalDelegatedEvent)
	if err := _Matic.contract.UnpackLog(event, "WithdrawTotalDelegatedEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
