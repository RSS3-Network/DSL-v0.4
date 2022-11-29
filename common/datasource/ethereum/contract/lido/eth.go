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

// ETHMetaData contains all meta data concerning the ETH contract.
var ETHMetaData = &bind.MetaData{
	ABI: "[{\"constant\":false,\"inputs\":[],\"name\":\"resume\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"stop\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"hasInitialized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"STAKING_CONTROL_ROLE\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_depositContract\",\"type\":\"address\"},{\"name\":\"_oracle\",\"type\":\"address\"},{\"name\":\"_operators\",\"type\":\"address\"},{\"name\":\"_treasury\",\"type\":\"address\"},{\"name\":\"_insuranceFund\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getInsuranceFund\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_ethAmount\",\"type\":\"uint256\"}],\"name\":\"getSharesByPooledEth\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isStakingPaused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_sender\",\"type\":\"address\"},{\"name\":\"_recipient\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOperators\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_script\",\"type\":\"bytes\"}],\"name\":\"getEVMScriptExecutor\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_maxStakeLimit\",\"type\":\"uint256\"},{\"name\":\"_stakeLimitIncreasePerBlock\",\"type\":\"uint256\"}],\"name\":\"setStakingLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"RESUME_ROLE\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRecoveryVault\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"DEPOSIT_ROLE\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"DEPOSIT_SIZE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTotalPooledEther\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PAUSE_ROLE\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTreasury\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isStopped\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MANAGE_WITHDRAWAL_KEY\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBufferedEther\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"receiveELRewards\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getELRewardsWithdrawalLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"SIGNATURE_LENGTH\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getWithdrawalCredentials\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurrentStakeLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_limitPoints\",\"type\":\"uint16\"}],\"name\":\"setELRewardsWithdrawalLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_beaconValidators\",\"type\":\"uint256\"},{\"name\":\"_beaconBalance\",\"type\":\"uint256\"}],\"name\":\"handleOracleReport\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStakeLimitFullInfo\",\"outputs\":[{\"name\":\"isStakingPaused\",\"type\":\"bool\"},{\"name\":\"isStakingLimitSet\",\"type\":\"bool\"},{\"name\":\"currentStakeLimit\",\"type\":\"uint256\"},{\"name\":\"maxStakeLimit\",\"type\":\"uint256\"},{\"name\":\"maxStakeLimitGrowthBlocks\",\"type\":\"uint256\"},{\"name\":\"prevStakeLimit\",\"type\":\"uint256\"},{\"name\":\"prevStakeBlockNumber\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"SET_EL_REWARDS_WITHDRAWAL_LIMIT_ROLE\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getELRewardsVault\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"resumeStaking\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getFeeDistribution\",\"outputs\":[{\"name\":\"treasuryFeeBasisPoints\",\"type\":\"uint16\"},{\"name\":\"insuranceFeeBasisPoints\",\"type\":\"uint16\"},{\"name\":\"operatorsFeeBasisPoints\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_sharesAmount\",\"type\":\"uint256\"}],\"name\":\"getPooledEthByShares\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_executionLayerRewardsVault\",\"type\":\"address\"}],\"name\":\"setELRewardsVault\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"allowRecoverability\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MANAGE_PROTOCOL_CONTRACTS_ROLE\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"appId\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOracle\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getInitializationBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_treasuryFeeBasisPoints\",\"type\":\"uint16\"},{\"name\":\"_insuranceFeeBasisPoints\",\"type\":\"uint16\"},{\"name\":\"_operatorsFeeBasisPoints\",\"type\":\"uint16\"}],\"name\":\"setFeeDistribution\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_feeBasisPoints\",\"type\":\"uint16\"}],\"name\":\"setFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_recipient\",\"type\":\"address\"},{\"name\":\"_sharesAmount\",\"type\":\"uint256\"}],\"name\":\"transferShares\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_maxDeposits\",\"type\":\"uint256\"}],\"name\":\"depositBufferedEther\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MANAGE_FEE\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"transferToVault\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_sender\",\"type\":\"address\"},{\"name\":\"_role\",\"type\":\"bytes32\"},{\"name\":\"_params\",\"type\":\"uint256[]\"}],\"name\":\"canPerform\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_referral\",\"type\":\"address\"}],\"name\":\"submit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"WITHDRAWAL_CREDENTIALS_LENGTH\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEVMScriptRegistry\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PUBKEY_LENGTH\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"SET_EL_REWARDS_VAULT_ROLE\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_recipient\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getDepositContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBeaconStat\",\"outputs\":[{\"name\":\"depositedValidators\",\"type\":\"uint256\"},{\"name\":\"beaconValidators\",\"type\":\"uint256\"},{\"name\":\"beaconBalance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"removeStakingLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BURN_ROLE\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getFee\",\"outputs\":[{\"name\":\"feeBasisPoints\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"kernel\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTotalShares\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isPetrified\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_oracle\",\"type\":\"address\"},{\"name\":\"_treasury\",\"type\":\"address\"},{\"name\":\"_insuranceFund\",\"type\":\"address\"}],\"name\":\"setProtocolContracts\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_withdrawalCredentials\",\"type\":\"bytes32\"}],\"name\":\"setWithdrawalCredentials\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"STAKING_PAUSE_ROLE\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"depositBufferedEther\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_sharesAmount\",\"type\":\"uint256\"}],\"name\":\"burnShares\",\"outputs\":[{\"name\":\"newTotalShares\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"sharesOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pauseStaking\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTotalELRewardsCollected\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"executor\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"script\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"input\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"returnData\",\"type\":\"bytes\"}],\"name\":\"ScriptResult\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"vault\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RecoverToVault\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"sharesValue\",\"type\":\"uint256\"}],\"name\":\"TransferShares\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"preRebaseTokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"postRebaseTokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"sharesAmount\",\"type\":\"uint256\"}],\"name\":\"SharesBurnt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Stopped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Resumed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"StakingPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"StakingResumed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"maxStakeLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"stakeLimitIncreasePerBlock\",\"type\":\"uint256\"}],\"name\":\"StakingLimitSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"StakingLimitRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"oracle\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"treasury\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"insuranceFund\",\"type\":\"address\"}],\"name\":\"ProtocolContactsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"feeBasisPoints\",\"type\":\"uint16\"}],\"name\":\"FeeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"treasuryFeeBasisPoints\",\"type\":\"uint16\"},{\"indexed\":false,\"name\":\"insuranceFeeBasisPoints\",\"type\":\"uint16\"},{\"indexed\":false,\"name\":\"operatorsFeeBasisPoints\",\"type\":\"uint16\"}],\"name\":\"FeeDistributionSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ELRewardsReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"limitPoints\",\"type\":\"uint256\"}],\"name\":\"ELRewardsWithdrawalLimitSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"withdrawalCredentials\",\"type\":\"bytes32\"}],\"name\":\"WithdrawalCredentialsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"executionLayerRewardsVault\",\"type\":\"address\"}],\"name\":\"ELRewardsVaultSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"referral\",\"type\":\"address\"}],\"name\":\"Submitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Unbuffered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"sentFromBuffer\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"pubkeyHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"etherAmount\",\"type\":\"uint256\"}],\"name\":\"Withdrawal\",\"type\":\"event\"}]",
}

// ETHABI is the input ABI used to generate the binding from.
// Deprecated: Use ETHMetaData.ABI instead.
var ETHABI = ETHMetaData.ABI

// ETH is an auto generated Go binding around an Ethereum contract.
type ETH struct {
	ETHCaller     // Read-only binding to the contract
	ETHTransactor // Write-only binding to the contract
	ETHFilterer   // Log filterer for contract events
}

// ETHCaller is an auto generated read-only Go binding around an Ethereum contract.
type ETHCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ETHTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ETHTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ETHFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ETHFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ETHSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ETHSession struct {
	Contract     *ETH              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ETHCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ETHCallerSession struct {
	Contract *ETHCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ETHTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ETHTransactorSession struct {
	Contract     *ETHTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ETHRaw is an auto generated low-level Go binding around an Ethereum contract.
type ETHRaw struct {
	Contract *ETH // Generic contract binding to access the raw methods on
}

// ETHCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ETHCallerRaw struct {
	Contract *ETHCaller // Generic read-only contract binding to access the raw methods on
}

// ETHTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ETHTransactorRaw struct {
	Contract *ETHTransactor // Generic write-only contract binding to access the raw methods on
}

// NewETH creates a new instance of ETH, bound to a specific deployed contract.
func NewETH(address common.Address, backend bind.ContractBackend) (*ETH, error) {
	contract, err := bindETH(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ETH{ETHCaller: ETHCaller{contract: contract}, ETHTransactor: ETHTransactor{contract: contract}, ETHFilterer: ETHFilterer{contract: contract}}, nil
}

// NewETHCaller creates a new read-only instance of ETH, bound to a specific deployed contract.
func NewETHCaller(address common.Address, caller bind.ContractCaller) (*ETHCaller, error) {
	contract, err := bindETH(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ETHCaller{contract: contract}, nil
}

// NewETHTransactor creates a new write-only instance of ETH, bound to a specific deployed contract.
func NewETHTransactor(address common.Address, transactor bind.ContractTransactor) (*ETHTransactor, error) {
	contract, err := bindETH(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ETHTransactor{contract: contract}, nil
}

// NewETHFilterer creates a new log filterer instance of ETH, bound to a specific deployed contract.
func NewETHFilterer(address common.Address, filterer bind.ContractFilterer) (*ETHFilterer, error) {
	contract, err := bindETH(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ETHFilterer{contract: contract}, nil
}

// bindETH binds a generic wrapper to an already deployed contract.
func bindETH(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ETHABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ETH *ETHRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ETH.Contract.ETHCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ETH *ETHRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETH.Contract.ETHTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ETH *ETHRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ETH.Contract.ETHTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ETH *ETHCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ETH.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ETH *ETHTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETH.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ETH *ETHTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ETH.Contract.contract.Transact(opts, method, params...)
}

// BURNROLE is a free data retrieval call binding the contract method 0xb930908f.
//
// Solidity: function BURN_ROLE() view returns(bytes32)
func (_ETH *ETHCaller) BURNROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ETH.contract.Call(opts, &out, "BURN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BURNROLE is a free data retrieval call binding the contract method 0xb930908f.
//
// Solidity: function BURN_ROLE() view returns(bytes32)
func (_ETH *ETHSession) BURNROLE() ([32]byte, error) {
	return _ETH.Contract.BURNROLE(&_ETH.CallOpts)
}

// BURNROLE is a free data retrieval call binding the contract method 0xb930908f.
//
// Solidity: function BURN_ROLE() view returns(bytes32)
func (_ETH *ETHCallerSession) BURNROLE() ([32]byte, error) {
	return _ETH.Contract.BURNROLE(&_ETH.CallOpts)
}

// DEPOSITROLE is a free data retrieval call binding the contract method 0x353efdcf.
//
// Solidity: function DEPOSIT_ROLE() view returns(bytes32)
func (_ETH *ETHCaller) DEPOSITROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ETH.contract.Call(opts, &out, "DEPOSIT_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEPOSITROLE is a free data retrieval call binding the contract method 0x353efdcf.
//
// Solidity: function DEPOSIT_ROLE() view returns(bytes32)
func (_ETH *ETHSession) DEPOSITROLE() ([32]byte, error) {
	return _ETH.Contract.DEPOSITROLE(&_ETH.CallOpts)
}

// DEPOSITROLE is a free data retrieval call binding the contract method 0x353efdcf.
//
// Solidity: function DEPOSIT_ROLE() view returns(bytes32)
func (_ETH *ETHCallerSession) DEPOSITROLE() ([32]byte, error) {
	return _ETH.Contract.DEPOSITROLE(&_ETH.CallOpts)
}

// DEPOSITSIZE is a free data retrieval call binding the contract method 0x36bf3325.
//
// Solidity: function DEPOSIT_SIZE() view returns(uint256)
func (_ETH *ETHCaller) DEPOSITSIZE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ETH.contract.Call(opts, &out, "DEPOSIT_SIZE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEPOSITSIZE is a free data retrieval call binding the contract method 0x36bf3325.
//
// Solidity: function DEPOSIT_SIZE() view returns(uint256)
func (_ETH *ETHSession) DEPOSITSIZE() (*big.Int, error) {
	return _ETH.Contract.DEPOSITSIZE(&_ETH.CallOpts)
}

// DEPOSITSIZE is a free data retrieval call binding the contract method 0x36bf3325.
//
// Solidity: function DEPOSIT_SIZE() view returns(uint256)
func (_ETH *ETHCallerSession) DEPOSITSIZE() (*big.Int, error) {
	return _ETH.Contract.DEPOSITSIZE(&_ETH.CallOpts)
}

// MANAGEFEE is a free data retrieval call binding the contract method 0x9aaa2d15.
//
// Solidity: function MANAGE_FEE() view returns(bytes32)
func (_ETH *ETHCaller) MANAGEFEE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ETH.contract.Call(opts, &out, "MANAGE_FEE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MANAGEFEE is a free data retrieval call binding the contract method 0x9aaa2d15.
//
// Solidity: function MANAGE_FEE() view returns(bytes32)
func (_ETH *ETHSession) MANAGEFEE() ([32]byte, error) {
	return _ETH.Contract.MANAGEFEE(&_ETH.CallOpts)
}

// MANAGEFEE is a free data retrieval call binding the contract method 0x9aaa2d15.
//
// Solidity: function MANAGE_FEE() view returns(bytes32)
func (_ETH *ETHCallerSession) MANAGEFEE() ([32]byte, error) {
	return _ETH.Contract.MANAGEFEE(&_ETH.CallOpts)
}

// MANAGEPROTOCOLCONTRACTSROLE is a free data retrieval call binding the contract method 0x7f6fdac7.
//
// Solidity: function MANAGE_PROTOCOL_CONTRACTS_ROLE() view returns(bytes32)
func (_ETH *ETHCaller) MANAGEPROTOCOLCONTRACTSROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ETH.contract.Call(opts, &out, "MANAGE_PROTOCOL_CONTRACTS_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MANAGEPROTOCOLCONTRACTSROLE is a free data retrieval call binding the contract method 0x7f6fdac7.
//
// Solidity: function MANAGE_PROTOCOL_CONTRACTS_ROLE() view returns(bytes32)
func (_ETH *ETHSession) MANAGEPROTOCOLCONTRACTSROLE() ([32]byte, error) {
	return _ETH.Contract.MANAGEPROTOCOLCONTRACTSROLE(&_ETH.CallOpts)
}

// MANAGEPROTOCOLCONTRACTSROLE is a free data retrieval call binding the contract method 0x7f6fdac7.
//
// Solidity: function MANAGE_PROTOCOL_CONTRACTS_ROLE() view returns(bytes32)
func (_ETH *ETHCallerSession) MANAGEPROTOCOLCONTRACTSROLE() ([32]byte, error) {
	return _ETH.Contract.MANAGEPROTOCOLCONTRACTSROLE(&_ETH.CallOpts)
}

// MANAGEWITHDRAWALKEY is a free data retrieval call binding the contract method 0x435721da.
//
// Solidity: function MANAGE_WITHDRAWAL_KEY() view returns(bytes32)
func (_ETH *ETHCaller) MANAGEWITHDRAWALKEY(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ETH.contract.Call(opts, &out, "MANAGE_WITHDRAWAL_KEY")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MANAGEWITHDRAWALKEY is a free data retrieval call binding the contract method 0x435721da.
//
// Solidity: function MANAGE_WITHDRAWAL_KEY() view returns(bytes32)
func (_ETH *ETHSession) MANAGEWITHDRAWALKEY() ([32]byte, error) {
	return _ETH.Contract.MANAGEWITHDRAWALKEY(&_ETH.CallOpts)
}

// MANAGEWITHDRAWALKEY is a free data retrieval call binding the contract method 0x435721da.
//
// Solidity: function MANAGE_WITHDRAWAL_KEY() view returns(bytes32)
func (_ETH *ETHCallerSession) MANAGEWITHDRAWALKEY() ([32]byte, error) {
	return _ETH.Contract.MANAGEWITHDRAWALKEY(&_ETH.CallOpts)
}

// PAUSEROLE is a free data retrieval call binding the contract method 0x389ed267.
//
// Solidity: function PAUSE_ROLE() view returns(bytes32)
func (_ETH *ETHCaller) PAUSEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ETH.contract.Call(opts, &out, "PAUSE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSEROLE is a free data retrieval call binding the contract method 0x389ed267.
//
// Solidity: function PAUSE_ROLE() view returns(bytes32)
func (_ETH *ETHSession) PAUSEROLE() ([32]byte, error) {
	return _ETH.Contract.PAUSEROLE(&_ETH.CallOpts)
}

// PAUSEROLE is a free data retrieval call binding the contract method 0x389ed267.
//
// Solidity: function PAUSE_ROLE() view returns(bytes32)
func (_ETH *ETHCallerSession) PAUSEROLE() ([32]byte, error) {
	return _ETH.Contract.PAUSEROLE(&_ETH.CallOpts)
}

// PUBKEYLENGTH is a free data retrieval call binding the contract method 0xa4d55d1d.
//
// Solidity: function PUBKEY_LENGTH() view returns(uint256)
func (_ETH *ETHCaller) PUBKEYLENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ETH.contract.Call(opts, &out, "PUBKEY_LENGTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PUBKEYLENGTH is a free data retrieval call binding the contract method 0xa4d55d1d.
//
// Solidity: function PUBKEY_LENGTH() view returns(uint256)
func (_ETH *ETHSession) PUBKEYLENGTH() (*big.Int, error) {
	return _ETH.Contract.PUBKEYLENGTH(&_ETH.CallOpts)
}

// PUBKEYLENGTH is a free data retrieval call binding the contract method 0xa4d55d1d.
//
// Solidity: function PUBKEY_LENGTH() view returns(uint256)
func (_ETH *ETHCallerSession) PUBKEYLENGTH() (*big.Int, error) {
	return _ETH.Contract.PUBKEYLENGTH(&_ETH.CallOpts)
}

// RESUMEROLE is a free data retrieval call binding the contract method 0x2de03aa1.
//
// Solidity: function RESUME_ROLE() view returns(bytes32)
func (_ETH *ETHCaller) RESUMEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ETH.contract.Call(opts, &out, "RESUME_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RESUMEROLE is a free data retrieval call binding the contract method 0x2de03aa1.
//
// Solidity: function RESUME_ROLE() view returns(bytes32)
func (_ETH *ETHSession) RESUMEROLE() ([32]byte, error) {
	return _ETH.Contract.RESUMEROLE(&_ETH.CallOpts)
}

// RESUMEROLE is a free data retrieval call binding the contract method 0x2de03aa1.
//
// Solidity: function RESUME_ROLE() view returns(bytes32)
func (_ETH *ETHCallerSession) RESUMEROLE() ([32]byte, error) {
	return _ETH.Contract.RESUMEROLE(&_ETH.CallOpts)
}

// SETELREWARDSVAULTROLE is a free data retrieval call binding the contract method 0xa6426f5f.
//
// Solidity: function SET_EL_REWARDS_VAULT_ROLE() view returns(bytes32)
func (_ETH *ETHCaller) SETELREWARDSVAULTROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ETH.contract.Call(opts, &out, "SET_EL_REWARDS_VAULT_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SETELREWARDSVAULTROLE is a free data retrieval call binding the contract method 0xa6426f5f.
//
// Solidity: function SET_EL_REWARDS_VAULT_ROLE() view returns(bytes32)
func (_ETH *ETHSession) SETELREWARDSVAULTROLE() ([32]byte, error) {
	return _ETH.Contract.SETELREWARDSVAULTROLE(&_ETH.CallOpts)
}

// SETELREWARDSVAULTROLE is a free data retrieval call binding the contract method 0xa6426f5f.
//
// Solidity: function SET_EL_REWARDS_VAULT_ROLE() view returns(bytes32)
func (_ETH *ETHCallerSession) SETELREWARDSVAULTROLE() ([32]byte, error) {
	return _ETH.Contract.SETELREWARDSVAULTROLE(&_ETH.CallOpts)
}

// SETELREWARDSWITHDRAWALLIMITROLE is a free data retrieval call binding the contract method 0x6bb98ad3.
//
// Solidity: function SET_EL_REWARDS_WITHDRAWAL_LIMIT_ROLE() view returns(bytes32)
func (_ETH *ETHCaller) SETELREWARDSWITHDRAWALLIMITROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ETH.contract.Call(opts, &out, "SET_EL_REWARDS_WITHDRAWAL_LIMIT_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SETELREWARDSWITHDRAWALLIMITROLE is a free data retrieval call binding the contract method 0x6bb98ad3.
//
// Solidity: function SET_EL_REWARDS_WITHDRAWAL_LIMIT_ROLE() view returns(bytes32)
func (_ETH *ETHSession) SETELREWARDSWITHDRAWALLIMITROLE() ([32]byte, error) {
	return _ETH.Contract.SETELREWARDSWITHDRAWALLIMITROLE(&_ETH.CallOpts)
}

// SETELREWARDSWITHDRAWALLIMITROLE is a free data retrieval call binding the contract method 0x6bb98ad3.
//
// Solidity: function SET_EL_REWARDS_WITHDRAWAL_LIMIT_ROLE() view returns(bytes32)
func (_ETH *ETHCallerSession) SETELREWARDSWITHDRAWALLIMITROLE() ([32]byte, error) {
	return _ETH.Contract.SETELREWARDSWITHDRAWALLIMITROLE(&_ETH.CallOpts)
}

// SIGNATURELENGTH is a free data retrieval call binding the contract method 0x540bc5ea.
//
// Solidity: function SIGNATURE_LENGTH() view returns(uint256)
func (_ETH *ETHCaller) SIGNATURELENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ETH.contract.Call(opts, &out, "SIGNATURE_LENGTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SIGNATURELENGTH is a free data retrieval call binding the contract method 0x540bc5ea.
//
// Solidity: function SIGNATURE_LENGTH() view returns(uint256)
func (_ETH *ETHSession) SIGNATURELENGTH() (*big.Int, error) {
	return _ETH.Contract.SIGNATURELENGTH(&_ETH.CallOpts)
}

// SIGNATURELENGTH is a free data retrieval call binding the contract method 0x540bc5ea.
//
// Solidity: function SIGNATURE_LENGTH() view returns(uint256)
func (_ETH *ETHCallerSession) SIGNATURELENGTH() (*big.Int, error) {
	return _ETH.Contract.SIGNATURELENGTH(&_ETH.CallOpts)
}

// STAKINGCONTROLROLE is a free data retrieval call binding the contract method 0x136dd43c.
//
// Solidity: function STAKING_CONTROL_ROLE() view returns(bytes32)
func (_ETH *ETHCaller) STAKINGCONTROLROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ETH.contract.Call(opts, &out, "STAKING_CONTROL_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// STAKINGCONTROLROLE is a free data retrieval call binding the contract method 0x136dd43c.
//
// Solidity: function STAKING_CONTROL_ROLE() view returns(bytes32)
func (_ETH *ETHSession) STAKINGCONTROLROLE() ([32]byte, error) {
	return _ETH.Contract.STAKINGCONTROLROLE(&_ETH.CallOpts)
}

// STAKINGCONTROLROLE is a free data retrieval call binding the contract method 0x136dd43c.
//
// Solidity: function STAKING_CONTROL_ROLE() view returns(bytes32)
func (_ETH *ETHCallerSession) STAKINGCONTROLROLE() ([32]byte, error) {
	return _ETH.Contract.STAKINGCONTROLROLE(&_ETH.CallOpts)
}

// STAKINGPAUSEROLE is a free data retrieval call binding the contract method 0xeb85262f.
//
// Solidity: function STAKING_PAUSE_ROLE() view returns(bytes32)
func (_ETH *ETHCaller) STAKINGPAUSEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ETH.contract.Call(opts, &out, "STAKING_PAUSE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// STAKINGPAUSEROLE is a free data retrieval call binding the contract method 0xeb85262f.
//
// Solidity: function STAKING_PAUSE_ROLE() view returns(bytes32)
func (_ETH *ETHSession) STAKINGPAUSEROLE() ([32]byte, error) {
	return _ETH.Contract.STAKINGPAUSEROLE(&_ETH.CallOpts)
}

// STAKINGPAUSEROLE is a free data retrieval call binding the contract method 0xeb85262f.
//
// Solidity: function STAKING_PAUSE_ROLE() view returns(bytes32)
func (_ETH *ETHCallerSession) STAKINGPAUSEROLE() ([32]byte, error) {
	return _ETH.Contract.STAKINGPAUSEROLE(&_ETH.CallOpts)
}

// WITHDRAWALCREDENTIALSLENGTH is a free data retrieval call binding the contract method 0xa30448c0.
//
// Solidity: function WITHDRAWAL_CREDENTIALS_LENGTH() view returns(uint256)
func (_ETH *ETHCaller) WITHDRAWALCREDENTIALSLENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ETH.contract.Call(opts, &out, "WITHDRAWAL_CREDENTIALS_LENGTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WITHDRAWALCREDENTIALSLENGTH is a free data retrieval call binding the contract method 0xa30448c0.
//
// Solidity: function WITHDRAWAL_CREDENTIALS_LENGTH() view returns(uint256)
func (_ETH *ETHSession) WITHDRAWALCREDENTIALSLENGTH() (*big.Int, error) {
	return _ETH.Contract.WITHDRAWALCREDENTIALSLENGTH(&_ETH.CallOpts)
}

// WITHDRAWALCREDENTIALSLENGTH is a free data retrieval call binding the contract method 0xa30448c0.
//
// Solidity: function WITHDRAWAL_CREDENTIALS_LENGTH() view returns(uint256)
func (_ETH *ETHCallerSession) WITHDRAWALCREDENTIALSLENGTH() (*big.Int, error) {
	return _ETH.Contract.WITHDRAWALCREDENTIALSLENGTH(&_ETH.CallOpts)
}

// AllowRecoverability is a free data retrieval call binding the contract method 0x7e7db6e1.
//
// Solidity: function allowRecoverability(address token) view returns(bool)
func (_ETH *ETHCaller) AllowRecoverability(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _ETH.contract.Call(opts, &out, "allowRecoverability", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowRecoverability is a free data retrieval call binding the contract method 0x7e7db6e1.
//
// Solidity: function allowRecoverability(address token) view returns(bool)
func (_ETH *ETHSession) AllowRecoverability(token common.Address) (bool, error) {
	return _ETH.Contract.AllowRecoverability(&_ETH.CallOpts, token)
}

// AllowRecoverability is a free data retrieval call binding the contract method 0x7e7db6e1.
//
// Solidity: function allowRecoverability(address token) view returns(bool)
func (_ETH *ETHCallerSession) AllowRecoverability(token common.Address) (bool, error) {
	return _ETH.Contract.AllowRecoverability(&_ETH.CallOpts, token)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_ETH *ETHCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ETH.contract.Call(opts, &out, "allowance", _owner, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_ETH *ETHSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _ETH.Contract.Allowance(&_ETH.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_ETH *ETHCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _ETH.Contract.Allowance(&_ETH.CallOpts, _owner, _spender)
}

// AppId is a free data retrieval call binding the contract method 0x80afdea8.
//
// Solidity: function appId() view returns(bytes32)
func (_ETH *ETHCaller) AppId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ETH.contract.Call(opts, &out, "appId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AppId is a free data retrieval call binding the contract method 0x80afdea8.
//
// Solidity: function appId() view returns(bytes32)
func (_ETH *ETHSession) AppId() ([32]byte, error) {
	return _ETH.Contract.AppId(&_ETH.CallOpts)
}

// AppId is a free data retrieval call binding the contract method 0x80afdea8.
//
// Solidity: function appId() view returns(bytes32)
func (_ETH *ETHCallerSession) AppId() ([32]byte, error) {
	return _ETH.Contract.AppId(&_ETH.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_ETH *ETHCaller) BalanceOf(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ETH.contract.Call(opts, &out, "balanceOf", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_ETH *ETHSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _ETH.Contract.BalanceOf(&_ETH.CallOpts, _account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_ETH *ETHCallerSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _ETH.Contract.BalanceOf(&_ETH.CallOpts, _account)
}

// CanPerform is a free data retrieval call binding the contract method 0xa1658fad.
//
// Solidity: function canPerform(address _sender, bytes32 _role, uint256[] _params) view returns(bool)
func (_ETH *ETHCaller) CanPerform(opts *bind.CallOpts, _sender common.Address, _role [32]byte, _params []*big.Int) (bool, error) {
	var out []interface{}
	err := _ETH.contract.Call(opts, &out, "canPerform", _sender, _role, _params)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanPerform is a free data retrieval call binding the contract method 0xa1658fad.
//
// Solidity: function canPerform(address _sender, bytes32 _role, uint256[] _params) view returns(bool)
func (_ETH *ETHSession) CanPerform(_sender common.Address, _role [32]byte, _params []*big.Int) (bool, error) {
	return _ETH.Contract.CanPerform(&_ETH.CallOpts, _sender, _role, _params)
}

// CanPerform is a free data retrieval call binding the contract method 0xa1658fad.
//
// Solidity: function canPerform(address _sender, bytes32 _role, uint256[] _params) view returns(bool)
func (_ETH *ETHCallerSession) CanPerform(_sender common.Address, _role [32]byte, _params []*big.Int) (bool, error) {
	return _ETH.Contract.CanPerform(&_ETH.CallOpts, _sender, _role, _params)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_ETH *ETHCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ETH.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_ETH *ETHSession) Decimals() (uint8, error) {
	return _ETH.Contract.Decimals(&_ETH.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_ETH *ETHCallerSession) Decimals() (uint8, error) {
	return _ETH.Contract.Decimals(&_ETH.CallOpts)
}

// GetBeaconStat is a free data retrieval call binding the contract method 0xae2e3538.
//
// Solidity: function getBeaconStat() view returns(uint256 depositedValidators, uint256 beaconValidators, uint256 beaconBalance)
func (_ETH *ETHCaller) GetBeaconStat(opts *bind.CallOpts) (struct {
	DepositedValidators *big.Int
	BeaconValidators    *big.Int
	BeaconBalance       *big.Int
}, error) {
	var out []interface{}
	err := _ETH.contract.Call(opts, &out, "getBeaconStat")

	outstruct := new(struct {
		DepositedValidators *big.Int
		BeaconValidators    *big.Int
		BeaconBalance       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.DepositedValidators = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BeaconValidators = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BeaconBalance = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetBeaconStat is a free data retrieval call binding the contract method 0xae2e3538.
//
// Solidity: function getBeaconStat() view returns(uint256 depositedValidators, uint256 beaconValidators, uint256 beaconBalance)
func (_ETH *ETHSession) GetBeaconStat() (struct {
	DepositedValidators *big.Int
	BeaconValidators    *big.Int
	BeaconBalance       *big.Int
}, error) {
	return _ETH.Contract.GetBeaconStat(&_ETH.CallOpts)
}

// GetBeaconStat is a free data retrieval call binding the contract method 0xae2e3538.
//
// Solidity: function getBeaconStat() view returns(uint256 depositedValidators, uint256 beaconValidators, uint256 beaconBalance)
func (_ETH *ETHCallerSession) GetBeaconStat() (struct {
	DepositedValidators *big.Int
	BeaconValidators    *big.Int
	BeaconBalance       *big.Int
}, error) {
	return _ETH.Contract.GetBeaconStat(&_ETH.CallOpts)
}

// GetBufferedEther is a free data retrieval call binding the contract method 0x47b714e0.
//
// Solidity: function getBufferedEther() view returns(uint256)
func (_ETH *ETHCaller) GetBufferedEther(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ETH.contract.Call(opts, &out, "getBufferedEther")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBufferedEther is a free data retrieval call binding the contract method 0x47b714e0.
//
// Solidity: function getBufferedEther() view returns(uint256)
func (_ETH *ETHSession) GetBufferedEther() (*big.Int, error) {
	return _ETH.Contract.GetBufferedEther(&_ETH.CallOpts)
}

// GetBufferedEther is a free data retrieval call binding the contract method 0x47b714e0.
//
// Solidity: function getBufferedEther() view returns(uint256)
func (_ETH *ETHCallerSession) GetBufferedEther() (*big.Int, error) {
	return _ETH.Contract.GetBufferedEther(&_ETH.CallOpts)
}

// GetCurrentStakeLimit is a free data retrieval call binding the contract method 0x609c4c6c.
//
// Solidity: function getCurrentStakeLimit() view returns(uint256)
func (_ETH *ETHCaller) GetCurrentStakeLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ETH.contract.Call(opts, &out, "getCurrentStakeLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentStakeLimit is a free data retrieval call binding the contract method 0x609c4c6c.
//
// Solidity: function getCurrentStakeLimit() view returns(uint256)
func (_ETH *ETHSession) GetCurrentStakeLimit() (*big.Int, error) {
	return _ETH.Contract.GetCurrentStakeLimit(&_ETH.CallOpts)
}

// GetCurrentStakeLimit is a free data retrieval call binding the contract method 0x609c4c6c.
//
// Solidity: function getCurrentStakeLimit() view returns(uint256)
func (_ETH *ETHCallerSession) GetCurrentStakeLimit() (*big.Int, error) {
	return _ETH.Contract.GetCurrentStakeLimit(&_ETH.CallOpts)
}

// GetDepositContract is a free data retrieval call binding the contract method 0xab94276a.
//
// Solidity: function getDepositContract() view returns(address)
func (_ETH *ETHCaller) GetDepositContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ETH.contract.Call(opts, &out, "getDepositContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetDepositContract is a free data retrieval call binding the contract method 0xab94276a.
//
// Solidity: function getDepositContract() view returns(address)
func (_ETH *ETHSession) GetDepositContract() (common.Address, error) {
	return _ETH.Contract.GetDepositContract(&_ETH.CallOpts)
}

// GetDepositContract is a free data retrieval call binding the contract method 0xab94276a.
//
// Solidity: function getDepositContract() view returns(address)
func (_ETH *ETHCallerSession) GetDepositContract() (common.Address, error) {
	return _ETH.Contract.GetDepositContract(&_ETH.CallOpts)
}

// GetELRewardsVault is a free data retrieval call binding the contract method 0x706aa30d.
//
// Solidity: function getELRewardsVault() view returns(address)
func (_ETH *ETHCaller) GetELRewardsVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ETH.contract.Call(opts, &out, "getELRewardsVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetELRewardsVault is a free data retrieval call binding the contract method 0x706aa30d.
//
// Solidity: function getELRewardsVault() view returns(address)
func (_ETH *ETHSession) GetELRewardsVault() (common.Address, error) {
	return _ETH.Contract.GetELRewardsVault(&_ETH.CallOpts)
}

// GetELRewardsVault is a free data retrieval call binding the contract method 0x706aa30d.
//
// Solidity: function getELRewardsVault() view returns(address)
func (_ETH *ETHCallerSession) GetELRewardsVault() (common.Address, error) {
	return _ETH.Contract.GetELRewardsVault(&_ETH.CallOpts)
}

// GetELRewardsWithdrawalLimit is a free data retrieval call binding the contract method 0x52b3af93.
//
// Solidity: function getELRewardsWithdrawalLimit() view returns(uint256)
func (_ETH *ETHCaller) GetELRewardsWithdrawalLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ETH.contract.Call(opts, &out, "getELRewardsWithdrawalLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetELRewardsWithdrawalLimit is a free data retrieval call binding the contract method 0x52b3af93.
//
// Solidity: function getELRewardsWithdrawalLimit() view returns(uint256)
func (_ETH *ETHSession) GetELRewardsWithdrawalLimit() (*big.Int, error) {
	return _ETH.Contract.GetELRewardsWithdrawalLimit(&_ETH.CallOpts)
}

// GetELRewardsWithdrawalLimit is a free data retrieval call binding the contract method 0x52b3af93.
//
// Solidity: function getELRewardsWithdrawalLimit() view returns(uint256)
func (_ETH *ETHCallerSession) GetELRewardsWithdrawalLimit() (*big.Int, error) {
	return _ETH.Contract.GetELRewardsWithdrawalLimit(&_ETH.CallOpts)
}

// GetEVMScriptExecutor is a free data retrieval call binding the contract method 0x2914b9bd.
//
// Solidity: function getEVMScriptExecutor(bytes _script) view returns(address)
func (_ETH *ETHCaller) GetEVMScriptExecutor(opts *bind.CallOpts, _script []byte) (common.Address, error) {
	var out []interface{}
	err := _ETH.contract.Call(opts, &out, "getEVMScriptExecutor", _script)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEVMScriptExecutor is a free data retrieval call binding the contract method 0x2914b9bd.
//
// Solidity: function getEVMScriptExecutor(bytes _script) view returns(address)
func (_ETH *ETHSession) GetEVMScriptExecutor(_script []byte) (common.Address, error) {
	return _ETH.Contract.GetEVMScriptExecutor(&_ETH.CallOpts, _script)
}

// GetEVMScriptExecutor is a free data retrieval call binding the contract method 0x2914b9bd.
//
// Solidity: function getEVMScriptExecutor(bytes _script) view returns(address)
func (_ETH *ETHCallerSession) GetEVMScriptExecutor(_script []byte) (common.Address, error) {
	return _ETH.Contract.GetEVMScriptExecutor(&_ETH.CallOpts, _script)
}

// GetEVMScriptRegistry is a free data retrieval call binding the contract method 0xa479e508.
//
// Solidity: function getEVMScriptRegistry() view returns(address)
func (_ETH *ETHCaller) GetEVMScriptRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ETH.contract.Call(opts, &out, "getEVMScriptRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEVMScriptRegistry is a free data retrieval call binding the contract method 0xa479e508.
//
// Solidity: function getEVMScriptRegistry() view returns(address)
func (_ETH *ETHSession) GetEVMScriptRegistry() (common.Address, error) {
	return _ETH.Contract.GetEVMScriptRegistry(&_ETH.CallOpts)
}

// GetEVMScriptRegistry is a free data retrieval call binding the contract method 0xa479e508.
//
// Solidity: function getEVMScriptRegistry() view returns(address)
func (_ETH *ETHCallerSession) GetEVMScriptRegistry() (common.Address, error) {
	return _ETH.Contract.GetEVMScriptRegistry(&_ETH.CallOpts)
}

// GetFee is a free data retrieval call binding the contract method 0xced72f87.
//
// Solidity: function getFee() view returns(uint16 feeBasisPoints)
func (_ETH *ETHCaller) GetFee(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _ETH.contract.Call(opts, &out, "getFee")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GetFee is a free data retrieval call binding the contract method 0xced72f87.
//
// Solidity: function getFee() view returns(uint16 feeBasisPoints)
func (_ETH *ETHSession) GetFee() (uint16, error) {
	return _ETH.Contract.GetFee(&_ETH.CallOpts)
}

// GetFee is a free data retrieval call binding the contract method 0xced72f87.
//
// Solidity: function getFee() view returns(uint16 feeBasisPoints)
func (_ETH *ETHCallerSession) GetFee() (uint16, error) {
	return _ETH.Contract.GetFee(&_ETH.CallOpts)
}

// GetFeeDistribution is a free data retrieval call binding the contract method 0x752f77f1.
//
// Solidity: function getFeeDistribution() view returns(uint16 treasuryFeeBasisPoints, uint16 insuranceFeeBasisPoints, uint16 operatorsFeeBasisPoints)
func (_ETH *ETHCaller) GetFeeDistribution(opts *bind.CallOpts) (struct {
	TreasuryFeeBasisPoints  uint16
	InsuranceFeeBasisPoints uint16
	OperatorsFeeBasisPoints uint16
}, error) {
	var out []interface{}
	err := _ETH.contract.Call(opts, &out, "getFeeDistribution")

	outstruct := new(struct {
		TreasuryFeeBasisPoints  uint16
		InsuranceFeeBasisPoints uint16
		OperatorsFeeBasisPoints uint16
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TreasuryFeeBasisPoints = *abi.ConvertType(out[0], new(uint16)).(*uint16)
	outstruct.InsuranceFeeBasisPoints = *abi.ConvertType(out[1], new(uint16)).(*uint16)
	outstruct.OperatorsFeeBasisPoints = *abi.ConvertType(out[2], new(uint16)).(*uint16)

	return *outstruct, err

}

// GetFeeDistribution is a free data retrieval call binding the contract method 0x752f77f1.
//
// Solidity: function getFeeDistribution() view returns(uint16 treasuryFeeBasisPoints, uint16 insuranceFeeBasisPoints, uint16 operatorsFeeBasisPoints)
func (_ETH *ETHSession) GetFeeDistribution() (struct {
	TreasuryFeeBasisPoints  uint16
	InsuranceFeeBasisPoints uint16
	OperatorsFeeBasisPoints uint16
}, error) {
	return _ETH.Contract.GetFeeDistribution(&_ETH.CallOpts)
}

// GetFeeDistribution is a free data retrieval call binding the contract method 0x752f77f1.
//
// Solidity: function getFeeDistribution() view returns(uint16 treasuryFeeBasisPoints, uint16 insuranceFeeBasisPoints, uint16 operatorsFeeBasisPoints)
func (_ETH *ETHCallerSession) GetFeeDistribution() (struct {
	TreasuryFeeBasisPoints  uint16
	InsuranceFeeBasisPoints uint16
	OperatorsFeeBasisPoints uint16
}, error) {
	return _ETH.Contract.GetFeeDistribution(&_ETH.CallOpts)
}

// GetInitializationBlock is a free data retrieval call binding the contract method 0x8b3dd749.
//
// Solidity: function getInitializationBlock() view returns(uint256)
func (_ETH *ETHCaller) GetInitializationBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ETH.contract.Call(opts, &out, "getInitializationBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInitializationBlock is a free data retrieval call binding the contract method 0x8b3dd749.
//
// Solidity: function getInitializationBlock() view returns(uint256)
func (_ETH *ETHSession) GetInitializationBlock() (*big.Int, error) {
	return _ETH.Contract.GetInitializationBlock(&_ETH.CallOpts)
}

// GetInitializationBlock is a free data retrieval call binding the contract method 0x8b3dd749.
//
// Solidity: function getInitializationBlock() view returns(uint256)
func (_ETH *ETHCallerSession) GetInitializationBlock() (*big.Int, error) {
	return _ETH.Contract.GetInitializationBlock(&_ETH.CallOpts)
}

// GetInsuranceFund is a free data retrieval call binding the contract method 0x158626f7.
//
// Solidity: function getInsuranceFund() view returns(address)
func (_ETH *ETHCaller) GetInsuranceFund(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ETH.contract.Call(opts, &out, "getInsuranceFund")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetInsuranceFund is a free data retrieval call binding the contract method 0x158626f7.
//
// Solidity: function getInsuranceFund() view returns(address)
func (_ETH *ETHSession) GetInsuranceFund() (common.Address, error) {
	return _ETH.Contract.GetInsuranceFund(&_ETH.CallOpts)
}

// GetInsuranceFund is a free data retrieval call binding the contract method 0x158626f7.
//
// Solidity: function getInsuranceFund() view returns(address)
func (_ETH *ETHCallerSession) GetInsuranceFund() (common.Address, error) {
	return _ETH.Contract.GetInsuranceFund(&_ETH.CallOpts)
}

// GetOperators is a free data retrieval call binding the contract method 0x27a099d8.
//
// Solidity: function getOperators() view returns(address)
func (_ETH *ETHCaller) GetOperators(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ETH.contract.Call(opts, &out, "getOperators")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOperators is a free data retrieval call binding the contract method 0x27a099d8.
//
// Solidity: function getOperators() view returns(address)
func (_ETH *ETHSession) GetOperators() (common.Address, error) {
	return _ETH.Contract.GetOperators(&_ETH.CallOpts)
}

// GetOperators is a free data retrieval call binding the contract method 0x27a099d8.
//
// Solidity: function getOperators() view returns(address)
func (_ETH *ETHCallerSession) GetOperators() (common.Address, error) {
	return _ETH.Contract.GetOperators(&_ETH.CallOpts)
}

// GetOracle is a free data retrieval call binding the contract method 0x833b1fce.
//
// Solidity: function getOracle() view returns(address)
func (_ETH *ETHCaller) GetOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ETH.contract.Call(opts, &out, "getOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOracle is a free data retrieval call binding the contract method 0x833b1fce.
//
// Solidity: function getOracle() view returns(address)
func (_ETH *ETHSession) GetOracle() (common.Address, error) {
	return _ETH.Contract.GetOracle(&_ETH.CallOpts)
}

// GetOracle is a free data retrieval call binding the contract method 0x833b1fce.
//
// Solidity: function getOracle() view returns(address)
func (_ETH *ETHCallerSession) GetOracle() (common.Address, error) {
	return _ETH.Contract.GetOracle(&_ETH.CallOpts)
}

// GetPooledEthByShares is a free data retrieval call binding the contract method 0x7a28fb88.
//
// Solidity: function getPooledEthByShares(uint256 _sharesAmount) view returns(uint256)
func (_ETH *ETHCaller) GetPooledEthByShares(opts *bind.CallOpts, _sharesAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ETH.contract.Call(opts, &out, "getPooledEthByShares", _sharesAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPooledEthByShares is a free data retrieval call binding the contract method 0x7a28fb88.
//
// Solidity: function getPooledEthByShares(uint256 _sharesAmount) view returns(uint256)
func (_ETH *ETHSession) GetPooledEthByShares(_sharesAmount *big.Int) (*big.Int, error) {
	return _ETH.Contract.GetPooledEthByShares(&_ETH.CallOpts, _sharesAmount)
}

// GetPooledEthByShares is a free data retrieval call binding the contract method 0x7a28fb88.
//
// Solidity: function getPooledEthByShares(uint256 _sharesAmount) view returns(uint256)
func (_ETH *ETHCallerSession) GetPooledEthByShares(_sharesAmount *big.Int) (*big.Int, error) {
	return _ETH.Contract.GetPooledEthByShares(&_ETH.CallOpts, _sharesAmount)
}

// GetRecoveryVault is a free data retrieval call binding the contract method 0x32f0a3b5.
//
// Solidity: function getRecoveryVault() view returns(address)
func (_ETH *ETHCaller) GetRecoveryVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ETH.contract.Call(opts, &out, "getRecoveryVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRecoveryVault is a free data retrieval call binding the contract method 0x32f0a3b5.
//
// Solidity: function getRecoveryVault() view returns(address)
func (_ETH *ETHSession) GetRecoveryVault() (common.Address, error) {
	return _ETH.Contract.GetRecoveryVault(&_ETH.CallOpts)
}

// GetRecoveryVault is a free data retrieval call binding the contract method 0x32f0a3b5.
//
// Solidity: function getRecoveryVault() view returns(address)
func (_ETH *ETHCallerSession) GetRecoveryVault() (common.Address, error) {
	return _ETH.Contract.GetRecoveryVault(&_ETH.CallOpts)
}

// GetSharesByPooledEth is a free data retrieval call binding the contract method 0x19208451.
//
// Solidity: function getSharesByPooledEth(uint256 _ethAmount) view returns(uint256)
func (_ETH *ETHCaller) GetSharesByPooledEth(opts *bind.CallOpts, _ethAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ETH.contract.Call(opts, &out, "getSharesByPooledEth", _ethAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSharesByPooledEth is a free data retrieval call binding the contract method 0x19208451.
//
// Solidity: function getSharesByPooledEth(uint256 _ethAmount) view returns(uint256)
func (_ETH *ETHSession) GetSharesByPooledEth(_ethAmount *big.Int) (*big.Int, error) {
	return _ETH.Contract.GetSharesByPooledEth(&_ETH.CallOpts, _ethAmount)
}

// GetSharesByPooledEth is a free data retrieval call binding the contract method 0x19208451.
//
// Solidity: function getSharesByPooledEth(uint256 _ethAmount) view returns(uint256)
func (_ETH *ETHCallerSession) GetSharesByPooledEth(_ethAmount *big.Int) (*big.Int, error) {
	return _ETH.Contract.GetSharesByPooledEth(&_ETH.CallOpts, _ethAmount)
}

// GetStakeLimitFullInfo is a free data retrieval call binding the contract method 0x665b4b0b.
//
// Solidity: function getStakeLimitFullInfo() view returns(bool isStakingPaused, bool isStakingLimitSet, uint256 currentStakeLimit, uint256 maxStakeLimit, uint256 maxStakeLimitGrowthBlocks, uint256 prevStakeLimit, uint256 prevStakeBlockNumber)
func (_ETH *ETHCaller) GetStakeLimitFullInfo(opts *bind.CallOpts) (struct {
	IsStakingPaused           bool
	IsStakingLimitSet         bool
	CurrentStakeLimit         *big.Int
	MaxStakeLimit             *big.Int
	MaxStakeLimitGrowthBlocks *big.Int
	PrevStakeLimit            *big.Int
	PrevStakeBlockNumber      *big.Int
}, error) {
	var out []interface{}
	err := _ETH.contract.Call(opts, &out, "getStakeLimitFullInfo")

	outstruct := new(struct {
		IsStakingPaused           bool
		IsStakingLimitSet         bool
		CurrentStakeLimit         *big.Int
		MaxStakeLimit             *big.Int
		MaxStakeLimitGrowthBlocks *big.Int
		PrevStakeLimit            *big.Int
		PrevStakeBlockNumber      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsStakingPaused = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.IsStakingLimitSet = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.CurrentStakeLimit = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.MaxStakeLimit = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.MaxStakeLimitGrowthBlocks = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.PrevStakeLimit = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.PrevStakeBlockNumber = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetStakeLimitFullInfo is a free data retrieval call binding the contract method 0x665b4b0b.
//
// Solidity: function getStakeLimitFullInfo() view returns(bool isStakingPaused, bool isStakingLimitSet, uint256 currentStakeLimit, uint256 maxStakeLimit, uint256 maxStakeLimitGrowthBlocks, uint256 prevStakeLimit, uint256 prevStakeBlockNumber)
func (_ETH *ETHSession) GetStakeLimitFullInfo() (struct {
	IsStakingPaused           bool
	IsStakingLimitSet         bool
	CurrentStakeLimit         *big.Int
	MaxStakeLimit             *big.Int
	MaxStakeLimitGrowthBlocks *big.Int
	PrevStakeLimit            *big.Int
	PrevStakeBlockNumber      *big.Int
}, error) {
	return _ETH.Contract.GetStakeLimitFullInfo(&_ETH.CallOpts)
}

// GetStakeLimitFullInfo is a free data retrieval call binding the contract method 0x665b4b0b.
//
// Solidity: function getStakeLimitFullInfo() view returns(bool isStakingPaused, bool isStakingLimitSet, uint256 currentStakeLimit, uint256 maxStakeLimit, uint256 maxStakeLimitGrowthBlocks, uint256 prevStakeLimit, uint256 prevStakeBlockNumber)
func (_ETH *ETHCallerSession) GetStakeLimitFullInfo() (struct {
	IsStakingPaused           bool
	IsStakingLimitSet         bool
	CurrentStakeLimit         *big.Int
	MaxStakeLimit             *big.Int
	MaxStakeLimitGrowthBlocks *big.Int
	PrevStakeLimit            *big.Int
	PrevStakeBlockNumber      *big.Int
}, error) {
	return _ETH.Contract.GetStakeLimitFullInfo(&_ETH.CallOpts)
}

// GetTotalELRewardsCollected is a free data retrieval call binding the contract method 0xfa64ebac.
//
// Solidity: function getTotalELRewardsCollected() view returns(uint256)
func (_ETH *ETHCaller) GetTotalELRewardsCollected(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ETH.contract.Call(opts, &out, "getTotalELRewardsCollected")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalELRewardsCollected is a free data retrieval call binding the contract method 0xfa64ebac.
//
// Solidity: function getTotalELRewardsCollected() view returns(uint256)
func (_ETH *ETHSession) GetTotalELRewardsCollected() (*big.Int, error) {
	return _ETH.Contract.GetTotalELRewardsCollected(&_ETH.CallOpts)
}

// GetTotalELRewardsCollected is a free data retrieval call binding the contract method 0xfa64ebac.
//
// Solidity: function getTotalELRewardsCollected() view returns(uint256)
func (_ETH *ETHCallerSession) GetTotalELRewardsCollected() (*big.Int, error) {
	return _ETH.Contract.GetTotalELRewardsCollected(&_ETH.CallOpts)
}

// GetTotalPooledEther is a free data retrieval call binding the contract method 0x37cfdaca.
//
// Solidity: function getTotalPooledEther() view returns(uint256)
func (_ETH *ETHCaller) GetTotalPooledEther(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ETH.contract.Call(opts, &out, "getTotalPooledEther")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalPooledEther is a free data retrieval call binding the contract method 0x37cfdaca.
//
// Solidity: function getTotalPooledEther() view returns(uint256)
func (_ETH *ETHSession) GetTotalPooledEther() (*big.Int, error) {
	return _ETH.Contract.GetTotalPooledEther(&_ETH.CallOpts)
}

// GetTotalPooledEther is a free data retrieval call binding the contract method 0x37cfdaca.
//
// Solidity: function getTotalPooledEther() view returns(uint256)
func (_ETH *ETHCallerSession) GetTotalPooledEther() (*big.Int, error) {
	return _ETH.Contract.GetTotalPooledEther(&_ETH.CallOpts)
}

// GetTotalShares is a free data retrieval call binding the contract method 0xd5002f2e.
//
// Solidity: function getTotalShares() view returns(uint256)
func (_ETH *ETHCaller) GetTotalShares(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ETH.contract.Call(opts, &out, "getTotalShares")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalShares is a free data retrieval call binding the contract method 0xd5002f2e.
//
// Solidity: function getTotalShares() view returns(uint256)
func (_ETH *ETHSession) GetTotalShares() (*big.Int, error) {
	return _ETH.Contract.GetTotalShares(&_ETH.CallOpts)
}

// GetTotalShares is a free data retrieval call binding the contract method 0xd5002f2e.
//
// Solidity: function getTotalShares() view returns(uint256)
func (_ETH *ETHCallerSession) GetTotalShares() (*big.Int, error) {
	return _ETH.Contract.GetTotalShares(&_ETH.CallOpts)
}

// GetTreasury is a free data retrieval call binding the contract method 0x3b19e84a.
//
// Solidity: function getTreasury() view returns(address)
func (_ETH *ETHCaller) GetTreasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ETH.contract.Call(opts, &out, "getTreasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTreasury is a free data retrieval call binding the contract method 0x3b19e84a.
//
// Solidity: function getTreasury() view returns(address)
func (_ETH *ETHSession) GetTreasury() (common.Address, error) {
	return _ETH.Contract.GetTreasury(&_ETH.CallOpts)
}

// GetTreasury is a free data retrieval call binding the contract method 0x3b19e84a.
//
// Solidity: function getTreasury() view returns(address)
func (_ETH *ETHCallerSession) GetTreasury() (common.Address, error) {
	return _ETH.Contract.GetTreasury(&_ETH.CallOpts)
}

// GetWithdrawalCredentials is a free data retrieval call binding the contract method 0x56396715.
//
// Solidity: function getWithdrawalCredentials() view returns(bytes32)
func (_ETH *ETHCaller) GetWithdrawalCredentials(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ETH.contract.Call(opts, &out, "getWithdrawalCredentials")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetWithdrawalCredentials is a free data retrieval call binding the contract method 0x56396715.
//
// Solidity: function getWithdrawalCredentials() view returns(bytes32)
func (_ETH *ETHSession) GetWithdrawalCredentials() ([32]byte, error) {
	return _ETH.Contract.GetWithdrawalCredentials(&_ETH.CallOpts)
}

// GetWithdrawalCredentials is a free data retrieval call binding the contract method 0x56396715.
//
// Solidity: function getWithdrawalCredentials() view returns(bytes32)
func (_ETH *ETHCallerSession) GetWithdrawalCredentials() ([32]byte, error) {
	return _ETH.Contract.GetWithdrawalCredentials(&_ETH.CallOpts)
}

// HasInitialized is a free data retrieval call binding the contract method 0x0803fac0.
//
// Solidity: function hasInitialized() view returns(bool)
func (_ETH *ETHCaller) HasInitialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ETH.contract.Call(opts, &out, "hasInitialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasInitialized is a free data retrieval call binding the contract method 0x0803fac0.
//
// Solidity: function hasInitialized() view returns(bool)
func (_ETH *ETHSession) HasInitialized() (bool, error) {
	return _ETH.Contract.HasInitialized(&_ETH.CallOpts)
}

// HasInitialized is a free data retrieval call binding the contract method 0x0803fac0.
//
// Solidity: function hasInitialized() view returns(bool)
func (_ETH *ETHCallerSession) HasInitialized() (bool, error) {
	return _ETH.Contract.HasInitialized(&_ETH.CallOpts)
}

// IsPetrified is a free data retrieval call binding the contract method 0xde4796ed.
//
// Solidity: function isPetrified() view returns(bool)
func (_ETH *ETHCaller) IsPetrified(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ETH.contract.Call(opts, &out, "isPetrified")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPetrified is a free data retrieval call binding the contract method 0xde4796ed.
//
// Solidity: function isPetrified() view returns(bool)
func (_ETH *ETHSession) IsPetrified() (bool, error) {
	return _ETH.Contract.IsPetrified(&_ETH.CallOpts)
}

// IsPetrified is a free data retrieval call binding the contract method 0xde4796ed.
//
// Solidity: function isPetrified() view returns(bool)
func (_ETH *ETHCallerSession) IsPetrified() (bool, error) {
	return _ETH.Contract.IsPetrified(&_ETH.CallOpts)
}

// IsStakingPaused is a free data retrieval call binding the contract method 0x1ea7ca89.
//
// Solidity: function isStakingPaused() view returns(bool)
func (_ETH *ETHCaller) IsStakingPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ETH.contract.Call(opts, &out, "isStakingPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStakingPaused is a free data retrieval call binding the contract method 0x1ea7ca89.
//
// Solidity: function isStakingPaused() view returns(bool)
func (_ETH *ETHSession) IsStakingPaused() (bool, error) {
	return _ETH.Contract.IsStakingPaused(&_ETH.CallOpts)
}

// IsStakingPaused is a free data retrieval call binding the contract method 0x1ea7ca89.
//
// Solidity: function isStakingPaused() view returns(bool)
func (_ETH *ETHCallerSession) IsStakingPaused() (bool, error) {
	return _ETH.Contract.IsStakingPaused(&_ETH.CallOpts)
}

// IsStopped is a free data retrieval call binding the contract method 0x3f683b6a.
//
// Solidity: function isStopped() view returns(bool)
func (_ETH *ETHCaller) IsStopped(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ETH.contract.Call(opts, &out, "isStopped")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStopped is a free data retrieval call binding the contract method 0x3f683b6a.
//
// Solidity: function isStopped() view returns(bool)
func (_ETH *ETHSession) IsStopped() (bool, error) {
	return _ETH.Contract.IsStopped(&_ETH.CallOpts)
}

// IsStopped is a free data retrieval call binding the contract method 0x3f683b6a.
//
// Solidity: function isStopped() view returns(bool)
func (_ETH *ETHCallerSession) IsStopped() (bool, error) {
	return _ETH.Contract.IsStopped(&_ETH.CallOpts)
}

// Kernel is a free data retrieval call binding the contract method 0xd4aae0c4.
//
// Solidity: function kernel() view returns(address)
func (_ETH *ETHCaller) Kernel(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ETH.contract.Call(opts, &out, "kernel")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Kernel is a free data retrieval call binding the contract method 0xd4aae0c4.
//
// Solidity: function kernel() view returns(address)
func (_ETH *ETHSession) Kernel() (common.Address, error) {
	return _ETH.Contract.Kernel(&_ETH.CallOpts)
}

// Kernel is a free data retrieval call binding the contract method 0xd4aae0c4.
//
// Solidity: function kernel() view returns(address)
func (_ETH *ETHCallerSession) Kernel() (common.Address, error) {
	return _ETH.Contract.Kernel(&_ETH.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_ETH *ETHCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ETH.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_ETH *ETHSession) Name() (string, error) {
	return _ETH.Contract.Name(&_ETH.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_ETH *ETHCallerSession) Name() (string, error) {
	return _ETH.Contract.Name(&_ETH.CallOpts)
}

// SharesOf is a free data retrieval call binding the contract method 0xf5eb42dc.
//
// Solidity: function sharesOf(address _account) view returns(uint256)
func (_ETH *ETHCaller) SharesOf(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ETH.contract.Call(opts, &out, "sharesOf", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SharesOf is a free data retrieval call binding the contract method 0xf5eb42dc.
//
// Solidity: function sharesOf(address _account) view returns(uint256)
func (_ETH *ETHSession) SharesOf(_account common.Address) (*big.Int, error) {
	return _ETH.Contract.SharesOf(&_ETH.CallOpts, _account)
}

// SharesOf is a free data retrieval call binding the contract method 0xf5eb42dc.
//
// Solidity: function sharesOf(address _account) view returns(uint256)
func (_ETH *ETHCallerSession) SharesOf(_account common.Address) (*big.Int, error) {
	return _ETH.Contract.SharesOf(&_ETH.CallOpts, _account)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_ETH *ETHCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ETH.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_ETH *ETHSession) Symbol() (string, error) {
	return _ETH.Contract.Symbol(&_ETH.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_ETH *ETHCallerSession) Symbol() (string, error) {
	return _ETH.Contract.Symbol(&_ETH.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ETH *ETHCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ETH.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ETH *ETHSession) TotalSupply() (*big.Int, error) {
	return _ETH.Contract.TotalSupply(&_ETH.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ETH *ETHCallerSession) TotalSupply() (*big.Int, error) {
	return _ETH.Contract.TotalSupply(&_ETH.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_ETH *ETHTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ETH.contract.Transact(opts, "approve", _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_ETH *ETHSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ETH.Contract.Approve(&_ETH.TransactOpts, _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_ETH *ETHTransactorSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ETH.Contract.Approve(&_ETH.TransactOpts, _spender, _amount)
}

// BurnShares is a paid mutator transaction binding the contract method 0xee7a7c04.
//
// Solidity: function burnShares(address _account, uint256 _sharesAmount) returns(uint256 newTotalShares)
func (_ETH *ETHTransactor) BurnShares(opts *bind.TransactOpts, _account common.Address, _sharesAmount *big.Int) (*types.Transaction, error) {
	return _ETH.contract.Transact(opts, "burnShares", _account, _sharesAmount)
}

// BurnShares is a paid mutator transaction binding the contract method 0xee7a7c04.
//
// Solidity: function burnShares(address _account, uint256 _sharesAmount) returns(uint256 newTotalShares)
func (_ETH *ETHSession) BurnShares(_account common.Address, _sharesAmount *big.Int) (*types.Transaction, error) {
	return _ETH.Contract.BurnShares(&_ETH.TransactOpts, _account, _sharesAmount)
}

// BurnShares is a paid mutator transaction binding the contract method 0xee7a7c04.
//
// Solidity: function burnShares(address _account, uint256 _sharesAmount) returns(uint256 newTotalShares)
func (_ETH *ETHTransactorSession) BurnShares(_account common.Address, _sharesAmount *big.Int) (*types.Transaction, error) {
	return _ETH.Contract.BurnShares(&_ETH.TransactOpts, _account, _sharesAmount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address _spender, uint256 _subtractedValue) returns(bool)
func (_ETH *ETHTransactor) DecreaseAllowance(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _ETH.contract.Transact(opts, "decreaseAllowance", _spender, _subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address _spender, uint256 _subtractedValue) returns(bool)
func (_ETH *ETHSession) DecreaseAllowance(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _ETH.Contract.DecreaseAllowance(&_ETH.TransactOpts, _spender, _subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address _spender, uint256 _subtractedValue) returns(bool)
func (_ETH *ETHTransactorSession) DecreaseAllowance(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _ETH.Contract.DecreaseAllowance(&_ETH.TransactOpts, _spender, _subtractedValue)
}

// DepositBufferedEther is a paid mutator transaction binding the contract method 0x90adc83b.
//
// Solidity: function depositBufferedEther(uint256 _maxDeposits) returns()
func (_ETH *ETHTransactor) DepositBufferedEther(opts *bind.TransactOpts, _maxDeposits *big.Int) (*types.Transaction, error) {
	return _ETH.contract.Transact(opts, "depositBufferedEther", _maxDeposits)
}

// DepositBufferedEther is a paid mutator transaction binding the contract method 0x90adc83b.
//
// Solidity: function depositBufferedEther(uint256 _maxDeposits) returns()
func (_ETH *ETHSession) DepositBufferedEther(_maxDeposits *big.Int) (*types.Transaction, error) {
	return _ETH.Contract.DepositBufferedEther(&_ETH.TransactOpts, _maxDeposits)
}

// DepositBufferedEther is a paid mutator transaction binding the contract method 0x90adc83b.
//
// Solidity: function depositBufferedEther(uint256 _maxDeposits) returns()
func (_ETH *ETHTransactorSession) DepositBufferedEther(_maxDeposits *big.Int) (*types.Transaction, error) {
	return _ETH.Contract.DepositBufferedEther(&_ETH.TransactOpts, _maxDeposits)
}

// DepositBufferedEther0 is a paid mutator transaction binding the contract method 0xecc1dcfb.
//
// Solidity: function depositBufferedEther() returns()
func (_ETH *ETHTransactor) DepositBufferedEther0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETH.contract.Transact(opts, "depositBufferedEther0")
}

// DepositBufferedEther0 is a paid mutator transaction binding the contract method 0xecc1dcfb.
//
// Solidity: function depositBufferedEther() returns()
func (_ETH *ETHSession) DepositBufferedEther0() (*types.Transaction, error) {
	return _ETH.Contract.DepositBufferedEther0(&_ETH.TransactOpts)
}

// DepositBufferedEther0 is a paid mutator transaction binding the contract method 0xecc1dcfb.
//
// Solidity: function depositBufferedEther() returns()
func (_ETH *ETHTransactorSession) DepositBufferedEther0() (*types.Transaction, error) {
	return _ETH.Contract.DepositBufferedEther0(&_ETH.TransactOpts)
}

// HandleOracleReport is a paid mutator transaction binding the contract method 0x64f9991a.
//
// Solidity: function handleOracleReport(uint256 _beaconValidators, uint256 _beaconBalance) returns()
func (_ETH *ETHTransactor) HandleOracleReport(opts *bind.TransactOpts, _beaconValidators *big.Int, _beaconBalance *big.Int) (*types.Transaction, error) {
	return _ETH.contract.Transact(opts, "handleOracleReport", _beaconValidators, _beaconBalance)
}

// HandleOracleReport is a paid mutator transaction binding the contract method 0x64f9991a.
//
// Solidity: function handleOracleReport(uint256 _beaconValidators, uint256 _beaconBalance) returns()
func (_ETH *ETHSession) HandleOracleReport(_beaconValidators *big.Int, _beaconBalance *big.Int) (*types.Transaction, error) {
	return _ETH.Contract.HandleOracleReport(&_ETH.TransactOpts, _beaconValidators, _beaconBalance)
}

// HandleOracleReport is a paid mutator transaction binding the contract method 0x64f9991a.
//
// Solidity: function handleOracleReport(uint256 _beaconValidators, uint256 _beaconBalance) returns()
func (_ETH *ETHTransactorSession) HandleOracleReport(_beaconValidators *big.Int, _beaconBalance *big.Int) (*types.Transaction, error) {
	return _ETH.Contract.HandleOracleReport(&_ETH.TransactOpts, _beaconValidators, _beaconBalance)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _spender, uint256 _addedValue) returns(bool)
func (_ETH *ETHTransactor) IncreaseAllowance(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _ETH.contract.Transact(opts, "increaseAllowance", _spender, _addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _spender, uint256 _addedValue) returns(bool)
func (_ETH *ETHSession) IncreaseAllowance(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _ETH.Contract.IncreaseAllowance(&_ETH.TransactOpts, _spender, _addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _spender, uint256 _addedValue) returns(bool)
func (_ETH *ETHTransactorSession) IncreaseAllowance(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _ETH.Contract.IncreaseAllowance(&_ETH.TransactOpts, _spender, _addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _depositContract, address _oracle, address _operators, address _treasury, address _insuranceFund) returns()
func (_ETH *ETHTransactor) Initialize(opts *bind.TransactOpts, _depositContract common.Address, _oracle common.Address, _operators common.Address, _treasury common.Address, _insuranceFund common.Address) (*types.Transaction, error) {
	return _ETH.contract.Transact(opts, "initialize", _depositContract, _oracle, _operators, _treasury, _insuranceFund)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _depositContract, address _oracle, address _operators, address _treasury, address _insuranceFund) returns()
func (_ETH *ETHSession) Initialize(_depositContract common.Address, _oracle common.Address, _operators common.Address, _treasury common.Address, _insuranceFund common.Address) (*types.Transaction, error) {
	return _ETH.Contract.Initialize(&_ETH.TransactOpts, _depositContract, _oracle, _operators, _treasury, _insuranceFund)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _depositContract, address _oracle, address _operators, address _treasury, address _insuranceFund) returns()
func (_ETH *ETHTransactorSession) Initialize(_depositContract common.Address, _oracle common.Address, _operators common.Address, _treasury common.Address, _insuranceFund common.Address) (*types.Transaction, error) {
	return _ETH.Contract.Initialize(&_ETH.TransactOpts, _depositContract, _oracle, _operators, _treasury, _insuranceFund)
}

// PauseStaking is a paid mutator transaction binding the contract method 0xf999c506.
//
// Solidity: function pauseStaking() returns()
func (_ETH *ETHTransactor) PauseStaking(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETH.contract.Transact(opts, "pauseStaking")
}

// PauseStaking is a paid mutator transaction binding the contract method 0xf999c506.
//
// Solidity: function pauseStaking() returns()
func (_ETH *ETHSession) PauseStaking() (*types.Transaction, error) {
	return _ETH.Contract.PauseStaking(&_ETH.TransactOpts)
}

// PauseStaking is a paid mutator transaction binding the contract method 0xf999c506.
//
// Solidity: function pauseStaking() returns()
func (_ETH *ETHTransactorSession) PauseStaking() (*types.Transaction, error) {
	return _ETH.Contract.PauseStaking(&_ETH.TransactOpts)
}

// ReceiveELRewards is a paid mutator transaction binding the contract method 0x4ad509b2.
//
// Solidity: function receiveELRewards() payable returns()
func (_ETH *ETHTransactor) ReceiveELRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETH.contract.Transact(opts, "receiveELRewards")
}

// ReceiveELRewards is a paid mutator transaction binding the contract method 0x4ad509b2.
//
// Solidity: function receiveELRewards() payable returns()
func (_ETH *ETHSession) ReceiveELRewards() (*types.Transaction, error) {
	return _ETH.Contract.ReceiveELRewards(&_ETH.TransactOpts)
}

// ReceiveELRewards is a paid mutator transaction binding the contract method 0x4ad509b2.
//
// Solidity: function receiveELRewards() payable returns()
func (_ETH *ETHTransactorSession) ReceiveELRewards() (*types.Transaction, error) {
	return _ETH.Contract.ReceiveELRewards(&_ETH.TransactOpts)
}

// RemoveStakingLimit is a paid mutator transaction binding the contract method 0xb3320d9a.
//
// Solidity: function removeStakingLimit() returns()
func (_ETH *ETHTransactor) RemoveStakingLimit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETH.contract.Transact(opts, "removeStakingLimit")
}

// RemoveStakingLimit is a paid mutator transaction binding the contract method 0xb3320d9a.
//
// Solidity: function removeStakingLimit() returns()
func (_ETH *ETHSession) RemoveStakingLimit() (*types.Transaction, error) {
	return _ETH.Contract.RemoveStakingLimit(&_ETH.TransactOpts)
}

// RemoveStakingLimit is a paid mutator transaction binding the contract method 0xb3320d9a.
//
// Solidity: function removeStakingLimit() returns()
func (_ETH *ETHTransactorSession) RemoveStakingLimit() (*types.Transaction, error) {
	return _ETH.Contract.RemoveStakingLimit(&_ETH.TransactOpts)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_ETH *ETHTransactor) Resume(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETH.contract.Transact(opts, "resume")
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_ETH *ETHSession) Resume() (*types.Transaction, error) {
	return _ETH.Contract.Resume(&_ETH.TransactOpts)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_ETH *ETHTransactorSession) Resume() (*types.Transaction, error) {
	return _ETH.Contract.Resume(&_ETH.TransactOpts)
}

// ResumeStaking is a paid mutator transaction binding the contract method 0x7475f913.
//
// Solidity: function resumeStaking() returns()
func (_ETH *ETHTransactor) ResumeStaking(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETH.contract.Transact(opts, "resumeStaking")
}

// ResumeStaking is a paid mutator transaction binding the contract method 0x7475f913.
//
// Solidity: function resumeStaking() returns()
func (_ETH *ETHSession) ResumeStaking() (*types.Transaction, error) {
	return _ETH.Contract.ResumeStaking(&_ETH.TransactOpts)
}

// ResumeStaking is a paid mutator transaction binding the contract method 0x7475f913.
//
// Solidity: function resumeStaking() returns()
func (_ETH *ETHTransactorSession) ResumeStaking() (*types.Transaction, error) {
	return _ETH.Contract.ResumeStaking(&_ETH.TransactOpts)
}

// SetELRewardsVault is a paid mutator transaction binding the contract method 0x7e4193c6.
//
// Solidity: function setELRewardsVault(address _executionLayerRewardsVault) returns()
func (_ETH *ETHTransactor) SetELRewardsVault(opts *bind.TransactOpts, _executionLayerRewardsVault common.Address) (*types.Transaction, error) {
	return _ETH.contract.Transact(opts, "setELRewardsVault", _executionLayerRewardsVault)
}

// SetELRewardsVault is a paid mutator transaction binding the contract method 0x7e4193c6.
//
// Solidity: function setELRewardsVault(address _executionLayerRewardsVault) returns()
func (_ETH *ETHSession) SetELRewardsVault(_executionLayerRewardsVault common.Address) (*types.Transaction, error) {
	return _ETH.Contract.SetELRewardsVault(&_ETH.TransactOpts, _executionLayerRewardsVault)
}

// SetELRewardsVault is a paid mutator transaction binding the contract method 0x7e4193c6.
//
// Solidity: function setELRewardsVault(address _executionLayerRewardsVault) returns()
func (_ETH *ETHTransactorSession) SetELRewardsVault(_executionLayerRewardsVault common.Address) (*types.Transaction, error) {
	return _ETH.Contract.SetELRewardsVault(&_ETH.TransactOpts, _executionLayerRewardsVault)
}

// SetELRewardsWithdrawalLimit is a paid mutator transaction binding the contract method 0x63c2eb53.
//
// Solidity: function setELRewardsWithdrawalLimit(uint16 _limitPoints) returns()
func (_ETH *ETHTransactor) SetELRewardsWithdrawalLimit(opts *bind.TransactOpts, _limitPoints uint16) (*types.Transaction, error) {
	return _ETH.contract.Transact(opts, "setELRewardsWithdrawalLimit", _limitPoints)
}

// SetELRewardsWithdrawalLimit is a paid mutator transaction binding the contract method 0x63c2eb53.
//
// Solidity: function setELRewardsWithdrawalLimit(uint16 _limitPoints) returns()
func (_ETH *ETHSession) SetELRewardsWithdrawalLimit(_limitPoints uint16) (*types.Transaction, error) {
	return _ETH.Contract.SetELRewardsWithdrawalLimit(&_ETH.TransactOpts, _limitPoints)
}

// SetELRewardsWithdrawalLimit is a paid mutator transaction binding the contract method 0x63c2eb53.
//
// Solidity: function setELRewardsWithdrawalLimit(uint16 _limitPoints) returns()
func (_ETH *ETHTransactorSession) SetELRewardsWithdrawalLimit(_limitPoints uint16) (*types.Transaction, error) {
	return _ETH.Contract.SetELRewardsWithdrawalLimit(&_ETH.TransactOpts, _limitPoints)
}

// SetFee is a paid mutator transaction binding the contract method 0x8e005553.
//
// Solidity: function setFee(uint16 _feeBasisPoints) returns()
func (_ETH *ETHTransactor) SetFee(opts *bind.TransactOpts, _feeBasisPoints uint16) (*types.Transaction, error) {
	return _ETH.contract.Transact(opts, "setFee", _feeBasisPoints)
}

// SetFee is a paid mutator transaction binding the contract method 0x8e005553.
//
// Solidity: function setFee(uint16 _feeBasisPoints) returns()
func (_ETH *ETHSession) SetFee(_feeBasisPoints uint16) (*types.Transaction, error) {
	return _ETH.Contract.SetFee(&_ETH.TransactOpts, _feeBasisPoints)
}

// SetFee is a paid mutator transaction binding the contract method 0x8e005553.
//
// Solidity: function setFee(uint16 _feeBasisPoints) returns()
func (_ETH *ETHTransactorSession) SetFee(_feeBasisPoints uint16) (*types.Transaction, error) {
	return _ETH.Contract.SetFee(&_ETH.TransactOpts, _feeBasisPoints)
}

// SetFeeDistribution is a paid mutator transaction binding the contract method 0x8cef3612.
//
// Solidity: function setFeeDistribution(uint16 _treasuryFeeBasisPoints, uint16 _insuranceFeeBasisPoints, uint16 _operatorsFeeBasisPoints) returns()
func (_ETH *ETHTransactor) SetFeeDistribution(opts *bind.TransactOpts, _treasuryFeeBasisPoints uint16, _insuranceFeeBasisPoints uint16, _operatorsFeeBasisPoints uint16) (*types.Transaction, error) {
	return _ETH.contract.Transact(opts, "setFeeDistribution", _treasuryFeeBasisPoints, _insuranceFeeBasisPoints, _operatorsFeeBasisPoints)
}

// SetFeeDistribution is a paid mutator transaction binding the contract method 0x8cef3612.
//
// Solidity: function setFeeDistribution(uint16 _treasuryFeeBasisPoints, uint16 _insuranceFeeBasisPoints, uint16 _operatorsFeeBasisPoints) returns()
func (_ETH *ETHSession) SetFeeDistribution(_treasuryFeeBasisPoints uint16, _insuranceFeeBasisPoints uint16, _operatorsFeeBasisPoints uint16) (*types.Transaction, error) {
	return _ETH.Contract.SetFeeDistribution(&_ETH.TransactOpts, _treasuryFeeBasisPoints, _insuranceFeeBasisPoints, _operatorsFeeBasisPoints)
}

// SetFeeDistribution is a paid mutator transaction binding the contract method 0x8cef3612.
//
// Solidity: function setFeeDistribution(uint16 _treasuryFeeBasisPoints, uint16 _insuranceFeeBasisPoints, uint16 _operatorsFeeBasisPoints) returns()
func (_ETH *ETHTransactorSession) SetFeeDistribution(_treasuryFeeBasisPoints uint16, _insuranceFeeBasisPoints uint16, _operatorsFeeBasisPoints uint16) (*types.Transaction, error) {
	return _ETH.Contract.SetFeeDistribution(&_ETH.TransactOpts, _treasuryFeeBasisPoints, _insuranceFeeBasisPoints, _operatorsFeeBasisPoints)
}

// SetProtocolContracts is a paid mutator transaction binding the contract method 0xe73f4529.
//
// Solidity: function setProtocolContracts(address _oracle, address _treasury, address _insuranceFund) returns()
func (_ETH *ETHTransactor) SetProtocolContracts(opts *bind.TransactOpts, _oracle common.Address, _treasury common.Address, _insuranceFund common.Address) (*types.Transaction, error) {
	return _ETH.contract.Transact(opts, "setProtocolContracts", _oracle, _treasury, _insuranceFund)
}

// SetProtocolContracts is a paid mutator transaction binding the contract method 0xe73f4529.
//
// Solidity: function setProtocolContracts(address _oracle, address _treasury, address _insuranceFund) returns()
func (_ETH *ETHSession) SetProtocolContracts(_oracle common.Address, _treasury common.Address, _insuranceFund common.Address) (*types.Transaction, error) {
	return _ETH.Contract.SetProtocolContracts(&_ETH.TransactOpts, _oracle, _treasury, _insuranceFund)
}

// SetProtocolContracts is a paid mutator transaction binding the contract method 0xe73f4529.
//
// Solidity: function setProtocolContracts(address _oracle, address _treasury, address _insuranceFund) returns()
func (_ETH *ETHTransactorSession) SetProtocolContracts(_oracle common.Address, _treasury common.Address, _insuranceFund common.Address) (*types.Transaction, error) {
	return _ETH.Contract.SetProtocolContracts(&_ETH.TransactOpts, _oracle, _treasury, _insuranceFund)
}

// SetStakingLimit is a paid mutator transaction binding the contract method 0x2cb5f784.
//
// Solidity: function setStakingLimit(uint256 _maxStakeLimit, uint256 _stakeLimitIncreasePerBlock) returns()
func (_ETH *ETHTransactor) SetStakingLimit(opts *bind.TransactOpts, _maxStakeLimit *big.Int, _stakeLimitIncreasePerBlock *big.Int) (*types.Transaction, error) {
	return _ETH.contract.Transact(opts, "setStakingLimit", _maxStakeLimit, _stakeLimitIncreasePerBlock)
}

// SetStakingLimit is a paid mutator transaction binding the contract method 0x2cb5f784.
//
// Solidity: function setStakingLimit(uint256 _maxStakeLimit, uint256 _stakeLimitIncreasePerBlock) returns()
func (_ETH *ETHSession) SetStakingLimit(_maxStakeLimit *big.Int, _stakeLimitIncreasePerBlock *big.Int) (*types.Transaction, error) {
	return _ETH.Contract.SetStakingLimit(&_ETH.TransactOpts, _maxStakeLimit, _stakeLimitIncreasePerBlock)
}

// SetStakingLimit is a paid mutator transaction binding the contract method 0x2cb5f784.
//
// Solidity: function setStakingLimit(uint256 _maxStakeLimit, uint256 _stakeLimitIncreasePerBlock) returns()
func (_ETH *ETHTransactorSession) SetStakingLimit(_maxStakeLimit *big.Int, _stakeLimitIncreasePerBlock *big.Int) (*types.Transaction, error) {
	return _ETH.Contract.SetStakingLimit(&_ETH.TransactOpts, _maxStakeLimit, _stakeLimitIncreasePerBlock)
}

// SetWithdrawalCredentials is a paid mutator transaction binding the contract method 0xe97ee8cc.
//
// Solidity: function setWithdrawalCredentials(bytes32 _withdrawalCredentials) returns()
func (_ETH *ETHTransactor) SetWithdrawalCredentials(opts *bind.TransactOpts, _withdrawalCredentials [32]byte) (*types.Transaction, error) {
	return _ETH.contract.Transact(opts, "setWithdrawalCredentials", _withdrawalCredentials)
}

// SetWithdrawalCredentials is a paid mutator transaction binding the contract method 0xe97ee8cc.
//
// Solidity: function setWithdrawalCredentials(bytes32 _withdrawalCredentials) returns()
func (_ETH *ETHSession) SetWithdrawalCredentials(_withdrawalCredentials [32]byte) (*types.Transaction, error) {
	return _ETH.Contract.SetWithdrawalCredentials(&_ETH.TransactOpts, _withdrawalCredentials)
}

// SetWithdrawalCredentials is a paid mutator transaction binding the contract method 0xe97ee8cc.
//
// Solidity: function setWithdrawalCredentials(bytes32 _withdrawalCredentials) returns()
func (_ETH *ETHTransactorSession) SetWithdrawalCredentials(_withdrawalCredentials [32]byte) (*types.Transaction, error) {
	return _ETH.Contract.SetWithdrawalCredentials(&_ETH.TransactOpts, _withdrawalCredentials)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_ETH *ETHTransactor) Stop(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETH.contract.Transact(opts, "stop")
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_ETH *ETHSession) Stop() (*types.Transaction, error) {
	return _ETH.Contract.Stop(&_ETH.TransactOpts)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_ETH *ETHTransactorSession) Stop() (*types.Transaction, error) {
	return _ETH.Contract.Stop(&_ETH.TransactOpts)
}

// Submit is a paid mutator transaction binding the contract method 0xa1903eab.
//
// Solidity: function submit(address _referral) payable returns(uint256)
func (_ETH *ETHTransactor) Submit(opts *bind.TransactOpts, _referral common.Address) (*types.Transaction, error) {
	return _ETH.contract.Transact(opts, "submit", _referral)
}

// Submit is a paid mutator transaction binding the contract method 0xa1903eab.
//
// Solidity: function submit(address _referral) payable returns(uint256)
func (_ETH *ETHSession) Submit(_referral common.Address) (*types.Transaction, error) {
	return _ETH.Contract.Submit(&_ETH.TransactOpts, _referral)
}

// Submit is a paid mutator transaction binding the contract method 0xa1903eab.
//
// Solidity: function submit(address _referral) payable returns(uint256)
func (_ETH *ETHTransactorSession) Submit(_referral common.Address) (*types.Transaction, error) {
	return _ETH.Contract.Submit(&_ETH.TransactOpts, _referral)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _recipient, uint256 _amount) returns(bool)
func (_ETH *ETHTransactor) Transfer(opts *bind.TransactOpts, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ETH.contract.Transact(opts, "transfer", _recipient, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _recipient, uint256 _amount) returns(bool)
func (_ETH *ETHSession) Transfer(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ETH.Contract.Transfer(&_ETH.TransactOpts, _recipient, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _recipient, uint256 _amount) returns(bool)
func (_ETH *ETHTransactorSession) Transfer(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ETH.Contract.Transfer(&_ETH.TransactOpts, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _sender, address _recipient, uint256 _amount) returns(bool)
func (_ETH *ETHTransactor) TransferFrom(opts *bind.TransactOpts, _sender common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ETH.contract.Transact(opts, "transferFrom", _sender, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _sender, address _recipient, uint256 _amount) returns(bool)
func (_ETH *ETHSession) TransferFrom(_sender common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ETH.Contract.TransferFrom(&_ETH.TransactOpts, _sender, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _sender, address _recipient, uint256 _amount) returns(bool)
func (_ETH *ETHTransactorSession) TransferFrom(_sender common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ETH.Contract.TransferFrom(&_ETH.TransactOpts, _sender, _recipient, _amount)
}

// TransferShares is a paid mutator transaction binding the contract method 0x8fcb4e5b.
//
// Solidity: function transferShares(address _recipient, uint256 _sharesAmount) returns(uint256)
func (_ETH *ETHTransactor) TransferShares(opts *bind.TransactOpts, _recipient common.Address, _sharesAmount *big.Int) (*types.Transaction, error) {
	return _ETH.contract.Transact(opts, "transferShares", _recipient, _sharesAmount)
}

// TransferShares is a paid mutator transaction binding the contract method 0x8fcb4e5b.
//
// Solidity: function transferShares(address _recipient, uint256 _sharesAmount) returns(uint256)
func (_ETH *ETHSession) TransferShares(_recipient common.Address, _sharesAmount *big.Int) (*types.Transaction, error) {
	return _ETH.Contract.TransferShares(&_ETH.TransactOpts, _recipient, _sharesAmount)
}

// TransferShares is a paid mutator transaction binding the contract method 0x8fcb4e5b.
//
// Solidity: function transferShares(address _recipient, uint256 _sharesAmount) returns(uint256)
func (_ETH *ETHTransactorSession) TransferShares(_recipient common.Address, _sharesAmount *big.Int) (*types.Transaction, error) {
	return _ETH.Contract.TransferShares(&_ETH.TransactOpts, _recipient, _sharesAmount)
}

// TransferToVault is a paid mutator transaction binding the contract method 0x9d4941d8.
//
// Solidity: function transferToVault(address _token) returns()
func (_ETH *ETHTransactor) TransferToVault(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _ETH.contract.Transact(opts, "transferToVault", _token)
}

// TransferToVault is a paid mutator transaction binding the contract method 0x9d4941d8.
//
// Solidity: function transferToVault(address _token) returns()
func (_ETH *ETHSession) TransferToVault(_token common.Address) (*types.Transaction, error) {
	return _ETH.Contract.TransferToVault(&_ETH.TransactOpts, _token)
}

// TransferToVault is a paid mutator transaction binding the contract method 0x9d4941d8.
//
// Solidity: function transferToVault(address _token) returns()
func (_ETH *ETHTransactorSession) TransferToVault(_token common.Address) (*types.Transaction, error) {
	return _ETH.Contract.TransferToVault(&_ETH.TransactOpts, _token)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ETH *ETHTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _ETH.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ETH *ETHSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _ETH.Contract.Fallback(&_ETH.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ETH *ETHTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _ETH.Contract.Fallback(&_ETH.TransactOpts, calldata)
}

// ETHApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ETH contract.
type ETHApprovalIterator struct {
	Event *ETHApproval // Event containing the contract specifics and raw log

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
func (it *ETHApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ETHApproval)
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
		it.Event = new(ETHApproval)
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
func (it *ETHApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ETHApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ETHApproval represents a Approval event raised by the ETH contract.
type ETHApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ETH *ETHFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ETHApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ETH.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ETHApprovalIterator{contract: _ETH.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ETH *ETHFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ETHApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ETH.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ETHApproval)
				if err := _ETH.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_ETH *ETHFilterer) ParseApproval(log types.Log) (*ETHApproval, error) {
	event := new(ETHApproval)
	if err := _ETH.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ETHELRewardsReceivedIterator is returned from FilterELRewardsReceived and is used to iterate over the raw logs and unpacked data for ELRewardsReceived events raised by the ETH contract.
type ETHELRewardsReceivedIterator struct {
	Event *ETHELRewardsReceived // Event containing the contract specifics and raw log

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
func (it *ETHELRewardsReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ETHELRewardsReceived)
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
		it.Event = new(ETHELRewardsReceived)
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
func (it *ETHELRewardsReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ETHELRewardsReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ETHELRewardsReceived represents a ELRewardsReceived event raised by the ETH contract.
type ETHELRewardsReceived struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterELRewardsReceived is a free log retrieval operation binding the contract event 0xd27f9b0c98bdee27044afa149eadcd2047d6399cb6613a45c5b87e6aca76e6b5.
//
// Solidity: event ELRewardsReceived(uint256 amount)
func (_ETH *ETHFilterer) FilterELRewardsReceived(opts *bind.FilterOpts) (*ETHELRewardsReceivedIterator, error) {

	logs, sub, err := _ETH.contract.FilterLogs(opts, "ELRewardsReceived")
	if err != nil {
		return nil, err
	}
	return &ETHELRewardsReceivedIterator{contract: _ETH.contract, event: "ELRewardsReceived", logs: logs, sub: sub}, nil
}

// WatchELRewardsReceived is a free log subscription operation binding the contract event 0xd27f9b0c98bdee27044afa149eadcd2047d6399cb6613a45c5b87e6aca76e6b5.
//
// Solidity: event ELRewardsReceived(uint256 amount)
func (_ETH *ETHFilterer) WatchELRewardsReceived(opts *bind.WatchOpts, sink chan<- *ETHELRewardsReceived) (event.Subscription, error) {

	logs, sub, err := _ETH.contract.WatchLogs(opts, "ELRewardsReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ETHELRewardsReceived)
				if err := _ETH.contract.UnpackLog(event, "ELRewardsReceived", log); err != nil {
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

// ParseELRewardsReceived is a log parse operation binding the contract event 0xd27f9b0c98bdee27044afa149eadcd2047d6399cb6613a45c5b87e6aca76e6b5.
//
// Solidity: event ELRewardsReceived(uint256 amount)
func (_ETH *ETHFilterer) ParseELRewardsReceived(log types.Log) (*ETHELRewardsReceived, error) {
	event := new(ETHELRewardsReceived)
	if err := _ETH.contract.UnpackLog(event, "ELRewardsReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ETHELRewardsVaultSetIterator is returned from FilterELRewardsVaultSet and is used to iterate over the raw logs and unpacked data for ELRewardsVaultSet events raised by the ETH contract.
type ETHELRewardsVaultSetIterator struct {
	Event *ETHELRewardsVaultSet // Event containing the contract specifics and raw log

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
func (it *ETHELRewardsVaultSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ETHELRewardsVaultSet)
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
		it.Event = new(ETHELRewardsVaultSet)
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
func (it *ETHELRewardsVaultSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ETHELRewardsVaultSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ETHELRewardsVaultSet represents a ELRewardsVaultSet event raised by the ETH contract.
type ETHELRewardsVaultSet struct {
	ExecutionLayerRewardsVault common.Address
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterELRewardsVaultSet is a free log retrieval operation binding the contract event 0x8e2d01c4cfaa88fa4d772d37e4d068deda4342bf4ef6dc4b0cf3e868be5ebb40.
//
// Solidity: event ELRewardsVaultSet(address executionLayerRewardsVault)
func (_ETH *ETHFilterer) FilterELRewardsVaultSet(opts *bind.FilterOpts) (*ETHELRewardsVaultSetIterator, error) {

	logs, sub, err := _ETH.contract.FilterLogs(opts, "ELRewardsVaultSet")
	if err != nil {
		return nil, err
	}
	return &ETHELRewardsVaultSetIterator{contract: _ETH.contract, event: "ELRewardsVaultSet", logs: logs, sub: sub}, nil
}

// WatchELRewardsVaultSet is a free log subscription operation binding the contract event 0x8e2d01c4cfaa88fa4d772d37e4d068deda4342bf4ef6dc4b0cf3e868be5ebb40.
//
// Solidity: event ELRewardsVaultSet(address executionLayerRewardsVault)
func (_ETH *ETHFilterer) WatchELRewardsVaultSet(opts *bind.WatchOpts, sink chan<- *ETHELRewardsVaultSet) (event.Subscription, error) {

	logs, sub, err := _ETH.contract.WatchLogs(opts, "ELRewardsVaultSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ETHELRewardsVaultSet)
				if err := _ETH.contract.UnpackLog(event, "ELRewardsVaultSet", log); err != nil {
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

// ParseELRewardsVaultSet is a log parse operation binding the contract event 0x8e2d01c4cfaa88fa4d772d37e4d068deda4342bf4ef6dc4b0cf3e868be5ebb40.
//
// Solidity: event ELRewardsVaultSet(address executionLayerRewardsVault)
func (_ETH *ETHFilterer) ParseELRewardsVaultSet(log types.Log) (*ETHELRewardsVaultSet, error) {
	event := new(ETHELRewardsVaultSet)
	if err := _ETH.contract.UnpackLog(event, "ELRewardsVaultSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ETHELRewardsWithdrawalLimitSetIterator is returned from FilterELRewardsWithdrawalLimitSet and is used to iterate over the raw logs and unpacked data for ELRewardsWithdrawalLimitSet events raised by the ETH contract.
type ETHELRewardsWithdrawalLimitSetIterator struct {
	Event *ETHELRewardsWithdrawalLimitSet // Event containing the contract specifics and raw log

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
func (it *ETHELRewardsWithdrawalLimitSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ETHELRewardsWithdrawalLimitSet)
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
		it.Event = new(ETHELRewardsWithdrawalLimitSet)
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
func (it *ETHELRewardsWithdrawalLimitSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ETHELRewardsWithdrawalLimitSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ETHELRewardsWithdrawalLimitSet represents a ELRewardsWithdrawalLimitSet event raised by the ETH contract.
type ETHELRewardsWithdrawalLimitSet struct {
	LimitPoints *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterELRewardsWithdrawalLimitSet is a free log retrieval operation binding the contract event 0x166eb213129ab51688433b859b5a206403ee174774a1430f8ffb83af316161f6.
//
// Solidity: event ELRewardsWithdrawalLimitSet(uint256 limitPoints)
func (_ETH *ETHFilterer) FilterELRewardsWithdrawalLimitSet(opts *bind.FilterOpts) (*ETHELRewardsWithdrawalLimitSetIterator, error) {

	logs, sub, err := _ETH.contract.FilterLogs(opts, "ELRewardsWithdrawalLimitSet")
	if err != nil {
		return nil, err
	}
	return &ETHELRewardsWithdrawalLimitSetIterator{contract: _ETH.contract, event: "ELRewardsWithdrawalLimitSet", logs: logs, sub: sub}, nil
}

// WatchELRewardsWithdrawalLimitSet is a free log subscription operation binding the contract event 0x166eb213129ab51688433b859b5a206403ee174774a1430f8ffb83af316161f6.
//
// Solidity: event ELRewardsWithdrawalLimitSet(uint256 limitPoints)
func (_ETH *ETHFilterer) WatchELRewardsWithdrawalLimitSet(opts *bind.WatchOpts, sink chan<- *ETHELRewardsWithdrawalLimitSet) (event.Subscription, error) {

	logs, sub, err := _ETH.contract.WatchLogs(opts, "ELRewardsWithdrawalLimitSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ETHELRewardsWithdrawalLimitSet)
				if err := _ETH.contract.UnpackLog(event, "ELRewardsWithdrawalLimitSet", log); err != nil {
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

// ParseELRewardsWithdrawalLimitSet is a log parse operation binding the contract event 0x166eb213129ab51688433b859b5a206403ee174774a1430f8ffb83af316161f6.
//
// Solidity: event ELRewardsWithdrawalLimitSet(uint256 limitPoints)
func (_ETH *ETHFilterer) ParseELRewardsWithdrawalLimitSet(log types.Log) (*ETHELRewardsWithdrawalLimitSet, error) {
	event := new(ETHELRewardsWithdrawalLimitSet)
	if err := _ETH.contract.UnpackLog(event, "ELRewardsWithdrawalLimitSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ETHFeeDistributionSetIterator is returned from FilterFeeDistributionSet and is used to iterate over the raw logs and unpacked data for FeeDistributionSet events raised by the ETH contract.
type ETHFeeDistributionSetIterator struct {
	Event *ETHFeeDistributionSet // Event containing the contract specifics and raw log

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
func (it *ETHFeeDistributionSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ETHFeeDistributionSet)
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
		it.Event = new(ETHFeeDistributionSet)
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
func (it *ETHFeeDistributionSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ETHFeeDistributionSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ETHFeeDistributionSet represents a FeeDistributionSet event raised by the ETH contract.
type ETHFeeDistributionSet struct {
	TreasuryFeeBasisPoints  uint16
	InsuranceFeeBasisPoints uint16
	OperatorsFeeBasisPoints uint16
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterFeeDistributionSet is a free log retrieval operation binding the contract event 0x034529db1bba3830b8877e116871f19c5b96ef86c739f2a05668c860c8466898.
//
// Solidity: event FeeDistributionSet(uint16 treasuryFeeBasisPoints, uint16 insuranceFeeBasisPoints, uint16 operatorsFeeBasisPoints)
func (_ETH *ETHFilterer) FilterFeeDistributionSet(opts *bind.FilterOpts) (*ETHFeeDistributionSetIterator, error) {

	logs, sub, err := _ETH.contract.FilterLogs(opts, "FeeDistributionSet")
	if err != nil {
		return nil, err
	}
	return &ETHFeeDistributionSetIterator{contract: _ETH.contract, event: "FeeDistributionSet", logs: logs, sub: sub}, nil
}

// WatchFeeDistributionSet is a free log subscription operation binding the contract event 0x034529db1bba3830b8877e116871f19c5b96ef86c739f2a05668c860c8466898.
//
// Solidity: event FeeDistributionSet(uint16 treasuryFeeBasisPoints, uint16 insuranceFeeBasisPoints, uint16 operatorsFeeBasisPoints)
func (_ETH *ETHFilterer) WatchFeeDistributionSet(opts *bind.WatchOpts, sink chan<- *ETHFeeDistributionSet) (event.Subscription, error) {

	logs, sub, err := _ETH.contract.WatchLogs(opts, "FeeDistributionSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ETHFeeDistributionSet)
				if err := _ETH.contract.UnpackLog(event, "FeeDistributionSet", log); err != nil {
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

// ParseFeeDistributionSet is a log parse operation binding the contract event 0x034529db1bba3830b8877e116871f19c5b96ef86c739f2a05668c860c8466898.
//
// Solidity: event FeeDistributionSet(uint16 treasuryFeeBasisPoints, uint16 insuranceFeeBasisPoints, uint16 operatorsFeeBasisPoints)
func (_ETH *ETHFilterer) ParseFeeDistributionSet(log types.Log) (*ETHFeeDistributionSet, error) {
	event := new(ETHFeeDistributionSet)
	if err := _ETH.contract.UnpackLog(event, "FeeDistributionSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ETHFeeSetIterator is returned from FilterFeeSet and is used to iterate over the raw logs and unpacked data for FeeSet events raised by the ETH contract.
type ETHFeeSetIterator struct {
	Event *ETHFeeSet // Event containing the contract specifics and raw log

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
func (it *ETHFeeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ETHFeeSet)
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
		it.Event = new(ETHFeeSet)
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
func (it *ETHFeeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ETHFeeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ETHFeeSet represents a FeeSet event raised by the ETH contract.
type ETHFeeSet struct {
	FeeBasisPoints uint16
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterFeeSet is a free log retrieval operation binding the contract event 0xaab062e3faf62b6c9a0f8e62af66e0310e27127a8c871a67be7dd4d93de6da53.
//
// Solidity: event FeeSet(uint16 feeBasisPoints)
func (_ETH *ETHFilterer) FilterFeeSet(opts *bind.FilterOpts) (*ETHFeeSetIterator, error) {

	logs, sub, err := _ETH.contract.FilterLogs(opts, "FeeSet")
	if err != nil {
		return nil, err
	}
	return &ETHFeeSetIterator{contract: _ETH.contract, event: "FeeSet", logs: logs, sub: sub}, nil
}

// WatchFeeSet is a free log subscription operation binding the contract event 0xaab062e3faf62b6c9a0f8e62af66e0310e27127a8c871a67be7dd4d93de6da53.
//
// Solidity: event FeeSet(uint16 feeBasisPoints)
func (_ETH *ETHFilterer) WatchFeeSet(opts *bind.WatchOpts, sink chan<- *ETHFeeSet) (event.Subscription, error) {

	logs, sub, err := _ETH.contract.WatchLogs(opts, "FeeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ETHFeeSet)
				if err := _ETH.contract.UnpackLog(event, "FeeSet", log); err != nil {
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

// ParseFeeSet is a log parse operation binding the contract event 0xaab062e3faf62b6c9a0f8e62af66e0310e27127a8c871a67be7dd4d93de6da53.
//
// Solidity: event FeeSet(uint16 feeBasisPoints)
func (_ETH *ETHFilterer) ParseFeeSet(log types.Log) (*ETHFeeSet, error) {
	event := new(ETHFeeSet)
	if err := _ETH.contract.UnpackLog(event, "FeeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ETHProtocolContactsSetIterator is returned from FilterProtocolContactsSet and is used to iterate over the raw logs and unpacked data for ProtocolContactsSet events raised by the ETH contract.
type ETHProtocolContactsSetIterator struct {
	Event *ETHProtocolContactsSet // Event containing the contract specifics and raw log

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
func (it *ETHProtocolContactsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ETHProtocolContactsSet)
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
		it.Event = new(ETHProtocolContactsSet)
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
func (it *ETHProtocolContactsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ETHProtocolContactsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ETHProtocolContactsSet represents a ProtocolContactsSet event raised by the ETH contract.
type ETHProtocolContactsSet struct {
	Oracle        common.Address
	Treasury      common.Address
	InsuranceFund common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterProtocolContactsSet is a free log retrieval operation binding the contract event 0x7df55cbe17c0cf85c9c23753c046f686eeb4c6b2ce13182943d215e92afc565a.
//
// Solidity: event ProtocolContactsSet(address oracle, address treasury, address insuranceFund)
func (_ETH *ETHFilterer) FilterProtocolContactsSet(opts *bind.FilterOpts) (*ETHProtocolContactsSetIterator, error) {

	logs, sub, err := _ETH.contract.FilterLogs(opts, "ProtocolContactsSet")
	if err != nil {
		return nil, err
	}
	return &ETHProtocolContactsSetIterator{contract: _ETH.contract, event: "ProtocolContactsSet", logs: logs, sub: sub}, nil
}

// WatchProtocolContactsSet is a free log subscription operation binding the contract event 0x7df55cbe17c0cf85c9c23753c046f686eeb4c6b2ce13182943d215e92afc565a.
//
// Solidity: event ProtocolContactsSet(address oracle, address treasury, address insuranceFund)
func (_ETH *ETHFilterer) WatchProtocolContactsSet(opts *bind.WatchOpts, sink chan<- *ETHProtocolContactsSet) (event.Subscription, error) {

	logs, sub, err := _ETH.contract.WatchLogs(opts, "ProtocolContactsSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ETHProtocolContactsSet)
				if err := _ETH.contract.UnpackLog(event, "ProtocolContactsSet", log); err != nil {
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

// ParseProtocolContactsSet is a log parse operation binding the contract event 0x7df55cbe17c0cf85c9c23753c046f686eeb4c6b2ce13182943d215e92afc565a.
//
// Solidity: event ProtocolContactsSet(address oracle, address treasury, address insuranceFund)
func (_ETH *ETHFilterer) ParseProtocolContactsSet(log types.Log) (*ETHProtocolContactsSet, error) {
	event := new(ETHProtocolContactsSet)
	if err := _ETH.contract.UnpackLog(event, "ProtocolContactsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ETHRecoverToVaultIterator is returned from FilterRecoverToVault and is used to iterate over the raw logs and unpacked data for RecoverToVault events raised by the ETH contract.
type ETHRecoverToVaultIterator struct {
	Event *ETHRecoverToVault // Event containing the contract specifics and raw log

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
func (it *ETHRecoverToVaultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ETHRecoverToVault)
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
		it.Event = new(ETHRecoverToVault)
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
func (it *ETHRecoverToVaultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ETHRecoverToVaultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ETHRecoverToVault represents a RecoverToVault event raised by the ETH contract.
type ETHRecoverToVault struct {
	Vault  common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRecoverToVault is a free log retrieval operation binding the contract event 0x596caf56044b55fb8c4ca640089bbc2b63cae3e978b851f5745cbb7c5b288e02.
//
// Solidity: event RecoverToVault(address indexed vault, address indexed token, uint256 amount)
func (_ETH *ETHFilterer) FilterRecoverToVault(opts *bind.FilterOpts, vault []common.Address, token []common.Address) (*ETHRecoverToVaultIterator, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ETH.contract.FilterLogs(opts, "RecoverToVault", vaultRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &ETHRecoverToVaultIterator{contract: _ETH.contract, event: "RecoverToVault", logs: logs, sub: sub}, nil
}

// WatchRecoverToVault is a free log subscription operation binding the contract event 0x596caf56044b55fb8c4ca640089bbc2b63cae3e978b851f5745cbb7c5b288e02.
//
// Solidity: event RecoverToVault(address indexed vault, address indexed token, uint256 amount)
func (_ETH *ETHFilterer) WatchRecoverToVault(opts *bind.WatchOpts, sink chan<- *ETHRecoverToVault, vault []common.Address, token []common.Address) (event.Subscription, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ETH.contract.WatchLogs(opts, "RecoverToVault", vaultRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ETHRecoverToVault)
				if err := _ETH.contract.UnpackLog(event, "RecoverToVault", log); err != nil {
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

// ParseRecoverToVault is a log parse operation binding the contract event 0x596caf56044b55fb8c4ca640089bbc2b63cae3e978b851f5745cbb7c5b288e02.
//
// Solidity: event RecoverToVault(address indexed vault, address indexed token, uint256 amount)
func (_ETH *ETHFilterer) ParseRecoverToVault(log types.Log) (*ETHRecoverToVault, error) {
	event := new(ETHRecoverToVault)
	if err := _ETH.contract.UnpackLog(event, "RecoverToVault", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ETHResumedIterator is returned from FilterResumed and is used to iterate over the raw logs and unpacked data for Resumed events raised by the ETH contract.
type ETHResumedIterator struct {
	Event *ETHResumed // Event containing the contract specifics and raw log

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
func (it *ETHResumedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ETHResumed)
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
		it.Event = new(ETHResumed)
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
func (it *ETHResumedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ETHResumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ETHResumed represents a Resumed event raised by the ETH contract.
type ETHResumed struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterResumed is a free log retrieval operation binding the contract event 0x62451d457bc659158be6e6247f56ec1df424a5c7597f71c20c2bc44e0965c8f9.
//
// Solidity: event Resumed()
func (_ETH *ETHFilterer) FilterResumed(opts *bind.FilterOpts) (*ETHResumedIterator, error) {

	logs, sub, err := _ETH.contract.FilterLogs(opts, "Resumed")
	if err != nil {
		return nil, err
	}
	return &ETHResumedIterator{contract: _ETH.contract, event: "Resumed", logs: logs, sub: sub}, nil
}

// WatchResumed is a free log subscription operation binding the contract event 0x62451d457bc659158be6e6247f56ec1df424a5c7597f71c20c2bc44e0965c8f9.
//
// Solidity: event Resumed()
func (_ETH *ETHFilterer) WatchResumed(opts *bind.WatchOpts, sink chan<- *ETHResumed) (event.Subscription, error) {

	logs, sub, err := _ETH.contract.WatchLogs(opts, "Resumed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ETHResumed)
				if err := _ETH.contract.UnpackLog(event, "Resumed", log); err != nil {
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

// ParseResumed is a log parse operation binding the contract event 0x62451d457bc659158be6e6247f56ec1df424a5c7597f71c20c2bc44e0965c8f9.
//
// Solidity: event Resumed()
func (_ETH *ETHFilterer) ParseResumed(log types.Log) (*ETHResumed, error) {
	event := new(ETHResumed)
	if err := _ETH.contract.UnpackLog(event, "Resumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ETHScriptResultIterator is returned from FilterScriptResult and is used to iterate over the raw logs and unpacked data for ScriptResult events raised by the ETH contract.
type ETHScriptResultIterator struct {
	Event *ETHScriptResult // Event containing the contract specifics and raw log

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
func (it *ETHScriptResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ETHScriptResult)
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
		it.Event = new(ETHScriptResult)
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
func (it *ETHScriptResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ETHScriptResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ETHScriptResult represents a ScriptResult event raised by the ETH contract.
type ETHScriptResult struct {
	Executor   common.Address
	Script     []byte
	Input      []byte
	ReturnData []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterScriptResult is a free log retrieval operation binding the contract event 0x5229a5dba83a54ae8cb5b51bdd6de9474cacbe9dd332f5185f3a4f4f2e3f4ad9.
//
// Solidity: event ScriptResult(address indexed executor, bytes script, bytes input, bytes returnData)
func (_ETH *ETHFilterer) FilterScriptResult(opts *bind.FilterOpts, executor []common.Address) (*ETHScriptResultIterator, error) {

	var executorRule []interface{}
	for _, executorItem := range executor {
		executorRule = append(executorRule, executorItem)
	}

	logs, sub, err := _ETH.contract.FilterLogs(opts, "ScriptResult", executorRule)
	if err != nil {
		return nil, err
	}
	return &ETHScriptResultIterator{contract: _ETH.contract, event: "ScriptResult", logs: logs, sub: sub}, nil
}

// WatchScriptResult is a free log subscription operation binding the contract event 0x5229a5dba83a54ae8cb5b51bdd6de9474cacbe9dd332f5185f3a4f4f2e3f4ad9.
//
// Solidity: event ScriptResult(address indexed executor, bytes script, bytes input, bytes returnData)
func (_ETH *ETHFilterer) WatchScriptResult(opts *bind.WatchOpts, sink chan<- *ETHScriptResult, executor []common.Address) (event.Subscription, error) {

	var executorRule []interface{}
	for _, executorItem := range executor {
		executorRule = append(executorRule, executorItem)
	}

	logs, sub, err := _ETH.contract.WatchLogs(opts, "ScriptResult", executorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ETHScriptResult)
				if err := _ETH.contract.UnpackLog(event, "ScriptResult", log); err != nil {
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

// ParseScriptResult is a log parse operation binding the contract event 0x5229a5dba83a54ae8cb5b51bdd6de9474cacbe9dd332f5185f3a4f4f2e3f4ad9.
//
// Solidity: event ScriptResult(address indexed executor, bytes script, bytes input, bytes returnData)
func (_ETH *ETHFilterer) ParseScriptResult(log types.Log) (*ETHScriptResult, error) {
	event := new(ETHScriptResult)
	if err := _ETH.contract.UnpackLog(event, "ScriptResult", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ETHSharesBurntIterator is returned from FilterSharesBurnt and is used to iterate over the raw logs and unpacked data for SharesBurnt events raised by the ETH contract.
type ETHSharesBurntIterator struct {
	Event *ETHSharesBurnt // Event containing the contract specifics and raw log

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
func (it *ETHSharesBurntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ETHSharesBurnt)
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
		it.Event = new(ETHSharesBurnt)
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
func (it *ETHSharesBurntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ETHSharesBurntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ETHSharesBurnt represents a SharesBurnt event raised by the ETH contract.
type ETHSharesBurnt struct {
	Account               common.Address
	PreRebaseTokenAmount  *big.Int
	PostRebaseTokenAmount *big.Int
	SharesAmount          *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterSharesBurnt is a free log retrieval operation binding the contract event 0x8b2a1e1ad5e0578c3dd82494156e985dade827a87c573b5c1c7716a32162ad64.
//
// Solidity: event SharesBurnt(address indexed account, uint256 preRebaseTokenAmount, uint256 postRebaseTokenAmount, uint256 sharesAmount)
func (_ETH *ETHFilterer) FilterSharesBurnt(opts *bind.FilterOpts, account []common.Address) (*ETHSharesBurntIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ETH.contract.FilterLogs(opts, "SharesBurnt", accountRule)
	if err != nil {
		return nil, err
	}
	return &ETHSharesBurntIterator{contract: _ETH.contract, event: "SharesBurnt", logs: logs, sub: sub}, nil
}

// WatchSharesBurnt is a free log subscription operation binding the contract event 0x8b2a1e1ad5e0578c3dd82494156e985dade827a87c573b5c1c7716a32162ad64.
//
// Solidity: event SharesBurnt(address indexed account, uint256 preRebaseTokenAmount, uint256 postRebaseTokenAmount, uint256 sharesAmount)
func (_ETH *ETHFilterer) WatchSharesBurnt(opts *bind.WatchOpts, sink chan<- *ETHSharesBurnt, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ETH.contract.WatchLogs(opts, "SharesBurnt", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ETHSharesBurnt)
				if err := _ETH.contract.UnpackLog(event, "SharesBurnt", log); err != nil {
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

// ParseSharesBurnt is a log parse operation binding the contract event 0x8b2a1e1ad5e0578c3dd82494156e985dade827a87c573b5c1c7716a32162ad64.
//
// Solidity: event SharesBurnt(address indexed account, uint256 preRebaseTokenAmount, uint256 postRebaseTokenAmount, uint256 sharesAmount)
func (_ETH *ETHFilterer) ParseSharesBurnt(log types.Log) (*ETHSharesBurnt, error) {
	event := new(ETHSharesBurnt)
	if err := _ETH.contract.UnpackLog(event, "SharesBurnt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ETHStakingLimitRemovedIterator is returned from FilterStakingLimitRemoved and is used to iterate over the raw logs and unpacked data for StakingLimitRemoved events raised by the ETH contract.
type ETHStakingLimitRemovedIterator struct {
	Event *ETHStakingLimitRemoved // Event containing the contract specifics and raw log

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
func (it *ETHStakingLimitRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ETHStakingLimitRemoved)
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
		it.Event = new(ETHStakingLimitRemoved)
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
func (it *ETHStakingLimitRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ETHStakingLimitRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ETHStakingLimitRemoved represents a StakingLimitRemoved event raised by the ETH contract.
type ETHStakingLimitRemoved struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterStakingLimitRemoved is a free log retrieval operation binding the contract event 0x9b2a687c198898fcc32a33bbc610d478f177a73ab7352023e6cc1de5bf12a3df.
//
// Solidity: event StakingLimitRemoved()
func (_ETH *ETHFilterer) FilterStakingLimitRemoved(opts *bind.FilterOpts) (*ETHStakingLimitRemovedIterator, error) {

	logs, sub, err := _ETH.contract.FilterLogs(opts, "StakingLimitRemoved")
	if err != nil {
		return nil, err
	}
	return &ETHStakingLimitRemovedIterator{contract: _ETH.contract, event: "StakingLimitRemoved", logs: logs, sub: sub}, nil
}

// WatchStakingLimitRemoved is a free log subscription operation binding the contract event 0x9b2a687c198898fcc32a33bbc610d478f177a73ab7352023e6cc1de5bf12a3df.
//
// Solidity: event StakingLimitRemoved()
func (_ETH *ETHFilterer) WatchStakingLimitRemoved(opts *bind.WatchOpts, sink chan<- *ETHStakingLimitRemoved) (event.Subscription, error) {

	logs, sub, err := _ETH.contract.WatchLogs(opts, "StakingLimitRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ETHStakingLimitRemoved)
				if err := _ETH.contract.UnpackLog(event, "StakingLimitRemoved", log); err != nil {
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

// ParseStakingLimitRemoved is a log parse operation binding the contract event 0x9b2a687c198898fcc32a33bbc610d478f177a73ab7352023e6cc1de5bf12a3df.
//
// Solidity: event StakingLimitRemoved()
func (_ETH *ETHFilterer) ParseStakingLimitRemoved(log types.Log) (*ETHStakingLimitRemoved, error) {
	event := new(ETHStakingLimitRemoved)
	if err := _ETH.contract.UnpackLog(event, "StakingLimitRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ETHStakingLimitSetIterator is returned from FilterStakingLimitSet and is used to iterate over the raw logs and unpacked data for StakingLimitSet events raised by the ETH contract.
type ETHStakingLimitSetIterator struct {
	Event *ETHStakingLimitSet // Event containing the contract specifics and raw log

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
func (it *ETHStakingLimitSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ETHStakingLimitSet)
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
		it.Event = new(ETHStakingLimitSet)
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
func (it *ETHStakingLimitSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ETHStakingLimitSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ETHStakingLimitSet represents a StakingLimitSet event raised by the ETH contract.
type ETHStakingLimitSet struct {
	MaxStakeLimit              *big.Int
	StakeLimitIncreasePerBlock *big.Int
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterStakingLimitSet is a free log retrieval operation binding the contract event 0xce9fddf6179affa1ea7bf36d80a6bf0284e0f3b91f4b2fa6eea2af923e7fac2d.
//
// Solidity: event StakingLimitSet(uint256 maxStakeLimit, uint256 stakeLimitIncreasePerBlock)
func (_ETH *ETHFilterer) FilterStakingLimitSet(opts *bind.FilterOpts) (*ETHStakingLimitSetIterator, error) {

	logs, sub, err := _ETH.contract.FilterLogs(opts, "StakingLimitSet")
	if err != nil {
		return nil, err
	}
	return &ETHStakingLimitSetIterator{contract: _ETH.contract, event: "StakingLimitSet", logs: logs, sub: sub}, nil
}

// WatchStakingLimitSet is a free log subscription operation binding the contract event 0xce9fddf6179affa1ea7bf36d80a6bf0284e0f3b91f4b2fa6eea2af923e7fac2d.
//
// Solidity: event StakingLimitSet(uint256 maxStakeLimit, uint256 stakeLimitIncreasePerBlock)
func (_ETH *ETHFilterer) WatchStakingLimitSet(opts *bind.WatchOpts, sink chan<- *ETHStakingLimitSet) (event.Subscription, error) {

	logs, sub, err := _ETH.contract.WatchLogs(opts, "StakingLimitSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ETHStakingLimitSet)
				if err := _ETH.contract.UnpackLog(event, "StakingLimitSet", log); err != nil {
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

// ParseStakingLimitSet is a log parse operation binding the contract event 0xce9fddf6179affa1ea7bf36d80a6bf0284e0f3b91f4b2fa6eea2af923e7fac2d.
//
// Solidity: event StakingLimitSet(uint256 maxStakeLimit, uint256 stakeLimitIncreasePerBlock)
func (_ETH *ETHFilterer) ParseStakingLimitSet(log types.Log) (*ETHStakingLimitSet, error) {
	event := new(ETHStakingLimitSet)
	if err := _ETH.contract.UnpackLog(event, "StakingLimitSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ETHStakingPausedIterator is returned from FilterStakingPaused and is used to iterate over the raw logs and unpacked data for StakingPaused events raised by the ETH contract.
type ETHStakingPausedIterator struct {
	Event *ETHStakingPaused // Event containing the contract specifics and raw log

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
func (it *ETHStakingPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ETHStakingPaused)
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
		it.Event = new(ETHStakingPaused)
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
func (it *ETHStakingPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ETHStakingPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ETHStakingPaused represents a StakingPaused event raised by the ETH contract.
type ETHStakingPaused struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterStakingPaused is a free log retrieval operation binding the contract event 0x26d1807b479eaba249c1214b82e4b65bbb0cc73ee8a17901324b1ef1b5904e49.
//
// Solidity: event StakingPaused()
func (_ETH *ETHFilterer) FilterStakingPaused(opts *bind.FilterOpts) (*ETHStakingPausedIterator, error) {

	logs, sub, err := _ETH.contract.FilterLogs(opts, "StakingPaused")
	if err != nil {
		return nil, err
	}
	return &ETHStakingPausedIterator{contract: _ETH.contract, event: "StakingPaused", logs: logs, sub: sub}, nil
}

// WatchStakingPaused is a free log subscription operation binding the contract event 0x26d1807b479eaba249c1214b82e4b65bbb0cc73ee8a17901324b1ef1b5904e49.
//
// Solidity: event StakingPaused()
func (_ETH *ETHFilterer) WatchStakingPaused(opts *bind.WatchOpts, sink chan<- *ETHStakingPaused) (event.Subscription, error) {

	logs, sub, err := _ETH.contract.WatchLogs(opts, "StakingPaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ETHStakingPaused)
				if err := _ETH.contract.UnpackLog(event, "StakingPaused", log); err != nil {
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

// ParseStakingPaused is a log parse operation binding the contract event 0x26d1807b479eaba249c1214b82e4b65bbb0cc73ee8a17901324b1ef1b5904e49.
//
// Solidity: event StakingPaused()
func (_ETH *ETHFilterer) ParseStakingPaused(log types.Log) (*ETHStakingPaused, error) {
	event := new(ETHStakingPaused)
	if err := _ETH.contract.UnpackLog(event, "StakingPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ETHStakingResumedIterator is returned from FilterStakingResumed and is used to iterate over the raw logs and unpacked data for StakingResumed events raised by the ETH contract.
type ETHStakingResumedIterator struct {
	Event *ETHStakingResumed // Event containing the contract specifics and raw log

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
func (it *ETHStakingResumedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ETHStakingResumed)
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
		it.Event = new(ETHStakingResumed)
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
func (it *ETHStakingResumedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ETHStakingResumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ETHStakingResumed represents a StakingResumed event raised by the ETH contract.
type ETHStakingResumed struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterStakingResumed is a free log retrieval operation binding the contract event 0xedaeeae9aed70c4545d3ab0065713261c9cee8d6cf5c8b07f52f0a65fd91efda.
//
// Solidity: event StakingResumed()
func (_ETH *ETHFilterer) FilterStakingResumed(opts *bind.FilterOpts) (*ETHStakingResumedIterator, error) {

	logs, sub, err := _ETH.contract.FilterLogs(opts, "StakingResumed")
	if err != nil {
		return nil, err
	}
	return &ETHStakingResumedIterator{contract: _ETH.contract, event: "StakingResumed", logs: logs, sub: sub}, nil
}

// WatchStakingResumed is a free log subscription operation binding the contract event 0xedaeeae9aed70c4545d3ab0065713261c9cee8d6cf5c8b07f52f0a65fd91efda.
//
// Solidity: event StakingResumed()
func (_ETH *ETHFilterer) WatchStakingResumed(opts *bind.WatchOpts, sink chan<- *ETHStakingResumed) (event.Subscription, error) {

	logs, sub, err := _ETH.contract.WatchLogs(opts, "StakingResumed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ETHStakingResumed)
				if err := _ETH.contract.UnpackLog(event, "StakingResumed", log); err != nil {
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

// ParseStakingResumed is a log parse operation binding the contract event 0xedaeeae9aed70c4545d3ab0065713261c9cee8d6cf5c8b07f52f0a65fd91efda.
//
// Solidity: event StakingResumed()
func (_ETH *ETHFilterer) ParseStakingResumed(log types.Log) (*ETHStakingResumed, error) {
	event := new(ETHStakingResumed)
	if err := _ETH.contract.UnpackLog(event, "StakingResumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ETHStoppedIterator is returned from FilterStopped and is used to iterate over the raw logs and unpacked data for Stopped events raised by the ETH contract.
type ETHStoppedIterator struct {
	Event *ETHStopped // Event containing the contract specifics and raw log

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
func (it *ETHStoppedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ETHStopped)
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
		it.Event = new(ETHStopped)
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
func (it *ETHStoppedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ETHStoppedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ETHStopped represents a Stopped event raised by the ETH contract.
type ETHStopped struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterStopped is a free log retrieval operation binding the contract event 0x7acc84e34091ae817647a4c49116f5cc07f319078ba80f8f5fde37ea7e25cbd6.
//
// Solidity: event Stopped()
func (_ETH *ETHFilterer) FilterStopped(opts *bind.FilterOpts) (*ETHStoppedIterator, error) {

	logs, sub, err := _ETH.contract.FilterLogs(opts, "Stopped")
	if err != nil {
		return nil, err
	}
	return &ETHStoppedIterator{contract: _ETH.contract, event: "Stopped", logs: logs, sub: sub}, nil
}

// WatchStopped is a free log subscription operation binding the contract event 0x7acc84e34091ae817647a4c49116f5cc07f319078ba80f8f5fde37ea7e25cbd6.
//
// Solidity: event Stopped()
func (_ETH *ETHFilterer) WatchStopped(opts *bind.WatchOpts, sink chan<- *ETHStopped) (event.Subscription, error) {

	logs, sub, err := _ETH.contract.WatchLogs(opts, "Stopped")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ETHStopped)
				if err := _ETH.contract.UnpackLog(event, "Stopped", log); err != nil {
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

// ParseStopped is a log parse operation binding the contract event 0x7acc84e34091ae817647a4c49116f5cc07f319078ba80f8f5fde37ea7e25cbd6.
//
// Solidity: event Stopped()
func (_ETH *ETHFilterer) ParseStopped(log types.Log) (*ETHStopped, error) {
	event := new(ETHStopped)
	if err := _ETH.contract.UnpackLog(event, "Stopped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ETHSubmittedIterator is returned from FilterSubmitted and is used to iterate over the raw logs and unpacked data for Submitted events raised by the ETH contract.
type ETHSubmittedIterator struct {
	Event *ETHSubmitted // Event containing the contract specifics and raw log

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
func (it *ETHSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ETHSubmitted)
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
		it.Event = new(ETHSubmitted)
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
func (it *ETHSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ETHSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ETHSubmitted represents a Submitted event raised by the ETH contract.
type ETHSubmitted struct {
	Sender   common.Address
	Amount   *big.Int
	Referral common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSubmitted is a free log retrieval operation binding the contract event 0x96a25c8ce0baabc1fdefd93e9ed25d8e092a3332f3aa9a41722b5697231d1d1a.
//
// Solidity: event Submitted(address indexed sender, uint256 amount, address referral)
func (_ETH *ETHFilterer) FilterSubmitted(opts *bind.FilterOpts, sender []common.Address) (*ETHSubmittedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ETH.contract.FilterLogs(opts, "Submitted", senderRule)
	if err != nil {
		return nil, err
	}
	return &ETHSubmittedIterator{contract: _ETH.contract, event: "Submitted", logs: logs, sub: sub}, nil
}

// WatchSubmitted is a free log subscription operation binding the contract event 0x96a25c8ce0baabc1fdefd93e9ed25d8e092a3332f3aa9a41722b5697231d1d1a.
//
// Solidity: event Submitted(address indexed sender, uint256 amount, address referral)
func (_ETH *ETHFilterer) WatchSubmitted(opts *bind.WatchOpts, sink chan<- *ETHSubmitted, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ETH.contract.WatchLogs(opts, "Submitted", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ETHSubmitted)
				if err := _ETH.contract.UnpackLog(event, "Submitted", log); err != nil {
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

// ParseSubmitted is a log parse operation binding the contract event 0x96a25c8ce0baabc1fdefd93e9ed25d8e092a3332f3aa9a41722b5697231d1d1a.
//
// Solidity: event Submitted(address indexed sender, uint256 amount, address referral)
func (_ETH *ETHFilterer) ParseSubmitted(log types.Log) (*ETHSubmitted, error) {
	event := new(ETHSubmitted)
	if err := _ETH.contract.UnpackLog(event, "Submitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ETHTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ETH contract.
type ETHTransferIterator struct {
	Event *ETHTransfer // Event containing the contract specifics and raw log

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
func (it *ETHTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ETHTransfer)
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
		it.Event = new(ETHTransfer)
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
func (it *ETHTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ETHTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ETHTransfer represents a Transfer event raised by the ETH contract.
type ETHTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ETH *ETHFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ETHTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ETH.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ETHTransferIterator{contract: _ETH.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ETH *ETHFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ETHTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ETH.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ETHTransfer)
				if err := _ETH.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_ETH *ETHFilterer) ParseTransfer(log types.Log) (*ETHTransfer, error) {
	event := new(ETHTransfer)
	if err := _ETH.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ETHTransferSharesIterator is returned from FilterTransferShares and is used to iterate over the raw logs and unpacked data for TransferShares events raised by the ETH contract.
type ETHTransferSharesIterator struct {
	Event *ETHTransferShares // Event containing the contract specifics and raw log

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
func (it *ETHTransferSharesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ETHTransferShares)
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
		it.Event = new(ETHTransferShares)
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
func (it *ETHTransferSharesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ETHTransferSharesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ETHTransferShares represents a TransferShares event raised by the ETH contract.
type ETHTransferShares struct {
	From        common.Address
	To          common.Address
	SharesValue *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTransferShares is a free log retrieval operation binding the contract event 0x9d9c909296d9c674451c0c24f02cb64981eb3b727f99865939192f880a755dcb.
//
// Solidity: event TransferShares(address indexed from, address indexed to, uint256 sharesValue)
func (_ETH *ETHFilterer) FilterTransferShares(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ETHTransferSharesIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ETH.contract.FilterLogs(opts, "TransferShares", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ETHTransferSharesIterator{contract: _ETH.contract, event: "TransferShares", logs: logs, sub: sub}, nil
}

// WatchTransferShares is a free log subscription operation binding the contract event 0x9d9c909296d9c674451c0c24f02cb64981eb3b727f99865939192f880a755dcb.
//
// Solidity: event TransferShares(address indexed from, address indexed to, uint256 sharesValue)
func (_ETH *ETHFilterer) WatchTransferShares(opts *bind.WatchOpts, sink chan<- *ETHTransferShares, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ETH.contract.WatchLogs(opts, "TransferShares", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ETHTransferShares)
				if err := _ETH.contract.UnpackLog(event, "TransferShares", log); err != nil {
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

// ParseTransferShares is a log parse operation binding the contract event 0x9d9c909296d9c674451c0c24f02cb64981eb3b727f99865939192f880a755dcb.
//
// Solidity: event TransferShares(address indexed from, address indexed to, uint256 sharesValue)
func (_ETH *ETHFilterer) ParseTransferShares(log types.Log) (*ETHTransferShares, error) {
	event := new(ETHTransferShares)
	if err := _ETH.contract.UnpackLog(event, "TransferShares", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ETHUnbufferedIterator is returned from FilterUnbuffered and is used to iterate over the raw logs and unpacked data for Unbuffered events raised by the ETH contract.
type ETHUnbufferedIterator struct {
	Event *ETHUnbuffered // Event containing the contract specifics and raw log

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
func (it *ETHUnbufferedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ETHUnbuffered)
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
		it.Event = new(ETHUnbuffered)
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
func (it *ETHUnbufferedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ETHUnbufferedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ETHUnbuffered represents a Unbuffered event raised by the ETH contract.
type ETHUnbuffered struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUnbuffered is a free log retrieval operation binding the contract event 0x76a397bea5768d4fca97ef47792796e35f98dc81b16c1de84e28a818e1f97108.
//
// Solidity: event Unbuffered(uint256 amount)
func (_ETH *ETHFilterer) FilterUnbuffered(opts *bind.FilterOpts) (*ETHUnbufferedIterator, error) {

	logs, sub, err := _ETH.contract.FilterLogs(opts, "Unbuffered")
	if err != nil {
		return nil, err
	}
	return &ETHUnbufferedIterator{contract: _ETH.contract, event: "Unbuffered", logs: logs, sub: sub}, nil
}

// WatchUnbuffered is a free log subscription operation binding the contract event 0x76a397bea5768d4fca97ef47792796e35f98dc81b16c1de84e28a818e1f97108.
//
// Solidity: event Unbuffered(uint256 amount)
func (_ETH *ETHFilterer) WatchUnbuffered(opts *bind.WatchOpts, sink chan<- *ETHUnbuffered) (event.Subscription, error) {

	logs, sub, err := _ETH.contract.WatchLogs(opts, "Unbuffered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ETHUnbuffered)
				if err := _ETH.contract.UnpackLog(event, "Unbuffered", log); err != nil {
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

// ParseUnbuffered is a log parse operation binding the contract event 0x76a397bea5768d4fca97ef47792796e35f98dc81b16c1de84e28a818e1f97108.
//
// Solidity: event Unbuffered(uint256 amount)
func (_ETH *ETHFilterer) ParseUnbuffered(log types.Log) (*ETHUnbuffered, error) {
	event := new(ETHUnbuffered)
	if err := _ETH.contract.UnpackLog(event, "Unbuffered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ETHWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the ETH contract.
type ETHWithdrawalIterator struct {
	Event *ETHWithdrawal // Event containing the contract specifics and raw log

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
func (it *ETHWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ETHWithdrawal)
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
		it.Event = new(ETHWithdrawal)
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
func (it *ETHWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ETHWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ETHWithdrawal represents a Withdrawal event raised by the ETH contract.
type ETHWithdrawal struct {
	Sender         common.Address
	TokenAmount    *big.Int
	SentFromBuffer *big.Int
	PubkeyHash     [32]byte
	EtherAmount    *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0xcf8f72073b13b07fe51690fd7c43414d1a0ef6f21c9896ba1814a08be9bdab3a.
//
// Solidity: event Withdrawal(address indexed sender, uint256 tokenAmount, uint256 sentFromBuffer, bytes32 indexed pubkeyHash, uint256 etherAmount)
func (_ETH *ETHFilterer) FilterWithdrawal(opts *bind.FilterOpts, sender []common.Address, pubkeyHash [][32]byte) (*ETHWithdrawalIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var pubkeyHashRule []interface{}
	for _, pubkeyHashItem := range pubkeyHash {
		pubkeyHashRule = append(pubkeyHashRule, pubkeyHashItem)
	}

	logs, sub, err := _ETH.contract.FilterLogs(opts, "Withdrawal", senderRule, pubkeyHashRule)
	if err != nil {
		return nil, err
	}
	return &ETHWithdrawalIterator{contract: _ETH.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0xcf8f72073b13b07fe51690fd7c43414d1a0ef6f21c9896ba1814a08be9bdab3a.
//
// Solidity: event Withdrawal(address indexed sender, uint256 tokenAmount, uint256 sentFromBuffer, bytes32 indexed pubkeyHash, uint256 etherAmount)
func (_ETH *ETHFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *ETHWithdrawal, sender []common.Address, pubkeyHash [][32]byte) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var pubkeyHashRule []interface{}
	for _, pubkeyHashItem := range pubkeyHash {
		pubkeyHashRule = append(pubkeyHashRule, pubkeyHashItem)
	}

	logs, sub, err := _ETH.contract.WatchLogs(opts, "Withdrawal", senderRule, pubkeyHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ETHWithdrawal)
				if err := _ETH.contract.UnpackLog(event, "Withdrawal", log); err != nil {
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

// ParseWithdrawal is a log parse operation binding the contract event 0xcf8f72073b13b07fe51690fd7c43414d1a0ef6f21c9896ba1814a08be9bdab3a.
//
// Solidity: event Withdrawal(address indexed sender, uint256 tokenAmount, uint256 sentFromBuffer, bytes32 indexed pubkeyHash, uint256 etherAmount)
func (_ETH *ETHFilterer) ParseWithdrawal(log types.Log) (*ETHWithdrawal, error) {
	event := new(ETHWithdrawal)
	if err := _ETH.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ETHWithdrawalCredentialsSetIterator is returned from FilterWithdrawalCredentialsSet and is used to iterate over the raw logs and unpacked data for WithdrawalCredentialsSet events raised by the ETH contract.
type ETHWithdrawalCredentialsSetIterator struct {
	Event *ETHWithdrawalCredentialsSet // Event containing the contract specifics and raw log

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
func (it *ETHWithdrawalCredentialsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ETHWithdrawalCredentialsSet)
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
		it.Event = new(ETHWithdrawalCredentialsSet)
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
func (it *ETHWithdrawalCredentialsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ETHWithdrawalCredentialsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ETHWithdrawalCredentialsSet represents a WithdrawalCredentialsSet event raised by the ETH contract.
type ETHWithdrawalCredentialsSet struct {
	WithdrawalCredentials [32]byte
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalCredentialsSet is a free log retrieval operation binding the contract event 0x13eb80e900aa05a2696d50d5de33ef631c73493c4921da233b17335ff6b7b114.
//
// Solidity: event WithdrawalCredentialsSet(bytes32 withdrawalCredentials)
func (_ETH *ETHFilterer) FilterWithdrawalCredentialsSet(opts *bind.FilterOpts) (*ETHWithdrawalCredentialsSetIterator, error) {

	logs, sub, err := _ETH.contract.FilterLogs(opts, "WithdrawalCredentialsSet")
	if err != nil {
		return nil, err
	}
	return &ETHWithdrawalCredentialsSetIterator{contract: _ETH.contract, event: "WithdrawalCredentialsSet", logs: logs, sub: sub}, nil
}

// WatchWithdrawalCredentialsSet is a free log subscription operation binding the contract event 0x13eb80e900aa05a2696d50d5de33ef631c73493c4921da233b17335ff6b7b114.
//
// Solidity: event WithdrawalCredentialsSet(bytes32 withdrawalCredentials)
func (_ETH *ETHFilterer) WatchWithdrawalCredentialsSet(opts *bind.WatchOpts, sink chan<- *ETHWithdrawalCredentialsSet) (event.Subscription, error) {

	logs, sub, err := _ETH.contract.WatchLogs(opts, "WithdrawalCredentialsSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ETHWithdrawalCredentialsSet)
				if err := _ETH.contract.UnpackLog(event, "WithdrawalCredentialsSet", log); err != nil {
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

// ParseWithdrawalCredentialsSet is a log parse operation binding the contract event 0x13eb80e900aa05a2696d50d5de33ef631c73493c4921da233b17335ff6b7b114.
//
// Solidity: event WithdrawalCredentialsSet(bytes32 withdrawalCredentials)
func (_ETH *ETHFilterer) ParseWithdrawalCredentialsSet(log types.Log) (*ETHWithdrawalCredentialsSet, error) {
	event := new(ETHWithdrawalCredentialsSet)
	if err := _ETH.contract.UnpackLog(event, "WithdrawalCredentialsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
