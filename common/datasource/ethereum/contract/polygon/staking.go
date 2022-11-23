// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package polygon

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

// StakingMetaData contains all meta data concerning the Staking contract.
var StakingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"ClaimFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"}],\"name\":\"ClaimRewards\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newValidatorId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"oldValidatorId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ConfirmAuction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"}],\"name\":\"DelegatorClaimedRewards\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"totalStaked\",\"type\":\"uint256\"}],\"name\":\"DelegatorRestaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DelegatorUnstaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newDynasty\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldDynasty\",\"type\":\"uint256\"}],\"name\":\"DynastyValueChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"exitEpoch\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"Jailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newProposerBonus\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldProposerBonus\",\"type\":\"uint256\"}],\"name\":\"ProposerBonusChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"}],\"name\":\"Restaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldReward\",\"type\":\"uint256\"}],\"name\":\"RewardUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"ShareBurned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"ShareMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldSigner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newSigner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"signerPubkey\",\"type\":\"bytes\"}],\"name\":\"SignerChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Slashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newAmount\",\"type\":\"uint256\"}],\"name\":\"StakeUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"activationEpoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"signerPubkey\",\"type\":\"bytes\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"auctionAmount\",\"type\":\"uint256\"}],\"name\":\"StartAuction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newThreshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldThreshold\",\"type\":\"uint256\"}],\"name\":\"ThresholdChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"TopUpFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"UnJailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deactivationEpoch\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"UnstakeInit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"}],\"name\":\"Unstaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newCommissionRate\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"oldCommissionRate\",\"type\":\"uint256\"}],\"name\":\"UpdateCommissionRate\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAccountStateRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"accountStateRoot\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"}],\"name\":\"getStakerDetails\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"activationEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deactivationEpoch\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_status\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"}],\"name\":\"getValidatorContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"ValidatorContract\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"logClaimFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"}],\"name\":\"logClaimRewards\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValidatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"oldValidatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"logConfirmAuction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"}],\"name\":\"logDelegatorClaimRewards\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalStaked\",\"type\":\"uint256\"}],\"name\":\"logDelegatorRestaked\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"logDelegatorUnstaked\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newDynasty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"oldDynasty\",\"type\":\"uint256\"}],\"name\":\"logDynastyValueChange\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exitEpoch\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"logJailed\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newProposerBonus\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"oldProposerBonus\",\"type\":\"uint256\"}],\"name\":\"logProposerBonusChange\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"}],\"name\":\"logRestaked\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"oldReward\",\"type\":\"uint256\"}],\"name\":\"logRewardUpdate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"logShareBurned\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"logShareMinted\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"oldSigner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newSigner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signerPubkey\",\"type\":\"bytes\"}],\"name\":\"logSignerChange\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"logSlashed\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"}],\"name\":\"logStakeUpdate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signerPubkey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"activationEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"}],\"name\":\"logStaked\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"auctionAmount\",\"type\":\"uint256\"}],\"name\":\"logStartAuction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"oldThreshold\",\"type\":\"uint256\"}],\"name\":\"logThresholdChange\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"logTopUpFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"logUnjailed\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deactivationEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"logUnstakeInit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"}],\"name\":\"logUnstaked\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newCommissionRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"oldCommissionRate\",\"type\":\"uint256\"}],\"name\":\"logUpdateCommissionRate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"contractRegistry\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"}],\"name\":\"totalValidatorStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorStake\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"validatorIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"nonces\",\"type\":\"uint256[]\"}],\"name\":\"updateNonce\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"validatorNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// StakingABI is the input ABI used to generate the binding from.
// Deprecated: Use StakingMetaData.ABI instead.
var StakingABI = StakingMetaData.ABI

// Staking is an auto generated Go binding around an Ethereum contract.
type Staking struct {
	StakingCaller     // Read-only binding to the contract
	StakingTransactor // Write-only binding to the contract
	StakingFilterer   // Log filterer for contract events
}

// StakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingSession struct {
	Contract     *Staking          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingCallerSession struct {
	Contract *StakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingTransactorSession struct {
	Contract     *StakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingRaw struct {
	Contract *Staking // Generic contract binding to access the raw methods on
}

// StakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingCallerRaw struct {
	Contract *StakingCaller // Generic read-only contract binding to access the raw methods on
}

// StakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingTransactorRaw struct {
	Contract *StakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStaking creates a new instance of Staking, bound to a specific deployed contract.
func NewStaking(address common.Address, backend bind.ContractBackend) (*Staking, error) {
	contract, err := bindStaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Staking{StakingCaller: StakingCaller{contract: contract}, StakingTransactor: StakingTransactor{contract: contract}, StakingFilterer: StakingFilterer{contract: contract}}, nil
}

// NewStakingCaller creates a new read-only instance of Staking, bound to a specific deployed contract.
func NewStakingCaller(address common.Address, caller bind.ContractCaller) (*StakingCaller, error) {
	contract, err := bindStaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingCaller{contract: contract}, nil
}

// NewStakingTransactor creates a new write-only instance of Staking, bound to a specific deployed contract.
func NewStakingTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingTransactor, error) {
	contract, err := bindStaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingTransactor{contract: contract}, nil
}

// NewStakingFilterer creates a new log filterer instance of Staking, bound to a specific deployed contract.
func NewStakingFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingFilterer, error) {
	contract, err := bindStaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingFilterer{contract: contract}, nil
}

// bindStaking binds a generic wrapper to an already deployed contract.
func bindStaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StakingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Staking *StakingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Staking.Contract.StakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Staking *StakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.Contract.StakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Staking *StakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Staking.Contract.StakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Staking *StakingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Staking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Staking *StakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Staking *StakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Staking.Contract.contract.Transact(opts, method, params...)
}

// GetAccountStateRoot is a free data retrieval call binding the contract method 0x4b6b87ce.
//
// Solidity: function getAccountStateRoot() view returns(bytes32 accountStateRoot)
func (_Staking *StakingCaller) GetAccountStateRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getAccountStateRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetAccountStateRoot is a free data retrieval call binding the contract method 0x4b6b87ce.
//
// Solidity: function getAccountStateRoot() view returns(bytes32 accountStateRoot)
func (_Staking *StakingSession) GetAccountStateRoot() ([32]byte, error) {
	return _Staking.Contract.GetAccountStateRoot(&_Staking.CallOpts)
}

// GetAccountStateRoot is a free data retrieval call binding the contract method 0x4b6b87ce.
//
// Solidity: function getAccountStateRoot() view returns(bytes32 accountStateRoot)
func (_Staking *StakingCallerSession) GetAccountStateRoot() ([32]byte, error) {
	return _Staking.Contract.GetAccountStateRoot(&_Staking.CallOpts)
}

// GetStakerDetails is a free data retrieval call binding the contract method 0x78daaf69.
//
// Solidity: function getStakerDetails(uint256 validatorId) view returns(uint256 amount, uint256 reward, uint256 activationEpoch, uint256 deactivationEpoch, address signer, uint256 _status)
func (_Staking *StakingCaller) GetStakerDetails(opts *bind.CallOpts, validatorId *big.Int) (struct {
	Amount            *big.Int
	Reward            *big.Int
	ActivationEpoch   *big.Int
	DeactivationEpoch *big.Int
	Signer            common.Address
	Status            *big.Int
}, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getStakerDetails", validatorId)

	outstruct := new(struct {
		Amount            *big.Int
		Reward            *big.Int
		ActivationEpoch   *big.Int
		DeactivationEpoch *big.Int
		Signer            common.Address
		Status            *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Reward = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ActivationEpoch = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.DeactivationEpoch = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Signer = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Status = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetStakerDetails is a free data retrieval call binding the contract method 0x78daaf69.
//
// Solidity: function getStakerDetails(uint256 validatorId) view returns(uint256 amount, uint256 reward, uint256 activationEpoch, uint256 deactivationEpoch, address signer, uint256 _status)
func (_Staking *StakingSession) GetStakerDetails(validatorId *big.Int) (struct {
	Amount            *big.Int
	Reward            *big.Int
	ActivationEpoch   *big.Int
	DeactivationEpoch *big.Int
	Signer            common.Address
	Status            *big.Int
}, error) {
	return _Staking.Contract.GetStakerDetails(&_Staking.CallOpts, validatorId)
}

// GetStakerDetails is a free data retrieval call binding the contract method 0x78daaf69.
//
// Solidity: function getStakerDetails(uint256 validatorId) view returns(uint256 amount, uint256 reward, uint256 activationEpoch, uint256 deactivationEpoch, address signer, uint256 _status)
func (_Staking *StakingCallerSession) GetStakerDetails(validatorId *big.Int) (struct {
	Amount            *big.Int
	Reward            *big.Int
	ActivationEpoch   *big.Int
	DeactivationEpoch *big.Int
	Signer            common.Address
	Status            *big.Int
}, error) {
	return _Staking.Contract.GetStakerDetails(&_Staking.CallOpts, validatorId)
}

// GetValidatorContractAddress is a free data retrieval call binding the contract method 0x178d46aa.
//
// Solidity: function getValidatorContractAddress(uint256 validatorId) view returns(address ValidatorContract)
func (_Staking *StakingCaller) GetValidatorContractAddress(opts *bind.CallOpts, validatorId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getValidatorContractAddress", validatorId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetValidatorContractAddress is a free data retrieval call binding the contract method 0x178d46aa.
//
// Solidity: function getValidatorContractAddress(uint256 validatorId) view returns(address ValidatorContract)
func (_Staking *StakingSession) GetValidatorContractAddress(validatorId *big.Int) (common.Address, error) {
	return _Staking.Contract.GetValidatorContractAddress(&_Staking.CallOpts, validatorId)
}

// GetValidatorContractAddress is a free data retrieval call binding the contract method 0x178d46aa.
//
// Solidity: function getValidatorContractAddress(uint256 validatorId) view returns(address ValidatorContract)
func (_Staking *StakingCallerSession) GetValidatorContractAddress(validatorId *big.Int) (common.Address, error) {
	return _Staking.Contract.GetValidatorContractAddress(&_Staking.CallOpts, validatorId)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Staking *StakingCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "isOwner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Staking *StakingSession) IsOwner() (bool, error) {
	return _Staking.Contract.IsOwner(&_Staking.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Staking *StakingCallerSession) IsOwner() (bool, error) {
	return _Staking.Contract.IsOwner(&_Staking.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Staking *StakingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Staking *StakingSession) Owner() (common.Address, error) {
	return _Staking.Contract.Owner(&_Staking.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Staking *StakingCallerSession) Owner() (common.Address, error) {
	return _Staking.Contract.Owner(&_Staking.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Staking *StakingCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "registry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Staking *StakingSession) Registry() (common.Address, error) {
	return _Staking.Contract.Registry(&_Staking.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Staking *StakingCallerSession) Registry() (common.Address, error) {
	return _Staking.Contract.Registry(&_Staking.CallOpts)
}

// TotalValidatorStake is a free data retrieval call binding the contract method 0xca7d34b6.
//
// Solidity: function totalValidatorStake(uint256 validatorId) view returns(uint256 validatorStake)
func (_Staking *StakingCaller) TotalValidatorStake(opts *bind.CallOpts, validatorId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "totalValidatorStake", validatorId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalValidatorStake is a free data retrieval call binding the contract method 0xca7d34b6.
//
// Solidity: function totalValidatorStake(uint256 validatorId) view returns(uint256 validatorStake)
func (_Staking *StakingSession) TotalValidatorStake(validatorId *big.Int) (*big.Int, error) {
	return _Staking.Contract.TotalValidatorStake(&_Staking.CallOpts, validatorId)
}

// TotalValidatorStake is a free data retrieval call binding the contract method 0xca7d34b6.
//
// Solidity: function totalValidatorStake(uint256 validatorId) view returns(uint256 validatorStake)
func (_Staking *StakingCallerSession) TotalValidatorStake(validatorId *big.Int) (*big.Int, error) {
	return _Staking.Contract.TotalValidatorStake(&_Staking.CallOpts, validatorId)
}

// ValidatorNonce is a free data retrieval call binding the contract method 0xebde9f93.
//
// Solidity: function validatorNonce(uint256 ) view returns(uint256)
func (_Staking *StakingCaller) ValidatorNonce(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "validatorNonce", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorNonce is a free data retrieval call binding the contract method 0xebde9f93.
//
// Solidity: function validatorNonce(uint256 ) view returns(uint256)
func (_Staking *StakingSession) ValidatorNonce(arg0 *big.Int) (*big.Int, error) {
	return _Staking.Contract.ValidatorNonce(&_Staking.CallOpts, arg0)
}

// ValidatorNonce is a free data retrieval call binding the contract method 0xebde9f93.
//
// Solidity: function validatorNonce(uint256 ) view returns(uint256)
func (_Staking *StakingCallerSession) ValidatorNonce(arg0 *big.Int) (*big.Int, error) {
	return _Staking.Contract.ValidatorNonce(&_Staking.CallOpts, arg0)
}

// LogClaimFee is a paid mutator transaction binding the contract method 0x122b6481.
//
// Solidity: function logClaimFee(address user, uint256 fee) returns()
func (_Staking *StakingTransactor) LogClaimFee(opts *bind.TransactOpts, user common.Address, fee *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "logClaimFee", user, fee)
}

// LogClaimFee is a paid mutator transaction binding the contract method 0x122b6481.
//
// Solidity: function logClaimFee(address user, uint256 fee) returns()
func (_Staking *StakingSession) LogClaimFee(user common.Address, fee *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.LogClaimFee(&_Staking.TransactOpts, user, fee)
}

// LogClaimFee is a paid mutator transaction binding the contract method 0x122b6481.
//
// Solidity: function logClaimFee(address user, uint256 fee) returns()
func (_Staking *StakingTransactorSession) LogClaimFee(user common.Address, fee *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.LogClaimFee(&_Staking.TransactOpts, user, fee)
}

// LogClaimRewards is a paid mutator transaction binding the contract method 0xb685b26a.
//
// Solidity: function logClaimRewards(uint256 validatorId, uint256 amount, uint256 totalAmount) returns()
func (_Staking *StakingTransactor) LogClaimRewards(opts *bind.TransactOpts, validatorId *big.Int, amount *big.Int, totalAmount *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "logClaimRewards", validatorId, amount, totalAmount)
}

// LogClaimRewards is a paid mutator transaction binding the contract method 0xb685b26a.
//
// Solidity: function logClaimRewards(uint256 validatorId, uint256 amount, uint256 totalAmount) returns()
func (_Staking *StakingSession) LogClaimRewards(validatorId *big.Int, amount *big.Int, totalAmount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.LogClaimRewards(&_Staking.TransactOpts, validatorId, amount, totalAmount)
}

// LogClaimRewards is a paid mutator transaction binding the contract method 0xb685b26a.
//
// Solidity: function logClaimRewards(uint256 validatorId, uint256 amount, uint256 totalAmount) returns()
func (_Staking *StakingTransactorSession) LogClaimRewards(validatorId *big.Int, amount *big.Int, totalAmount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.LogClaimRewards(&_Staking.TransactOpts, validatorId, amount, totalAmount)
}

// LogConfirmAuction is a paid mutator transaction binding the contract method 0xe12ab1af.
//
// Solidity: function logConfirmAuction(uint256 newValidatorId, uint256 oldValidatorId, uint256 amount) returns()
func (_Staking *StakingTransactor) LogConfirmAuction(opts *bind.TransactOpts, newValidatorId *big.Int, oldValidatorId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "logConfirmAuction", newValidatorId, oldValidatorId, amount)
}

// LogConfirmAuction is a paid mutator transaction binding the contract method 0xe12ab1af.
//
// Solidity: function logConfirmAuction(uint256 newValidatorId, uint256 oldValidatorId, uint256 amount) returns()
func (_Staking *StakingSession) LogConfirmAuction(newValidatorId *big.Int, oldValidatorId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.LogConfirmAuction(&_Staking.TransactOpts, newValidatorId, oldValidatorId, amount)
}

// LogConfirmAuction is a paid mutator transaction binding the contract method 0xe12ab1af.
//
// Solidity: function logConfirmAuction(uint256 newValidatorId, uint256 oldValidatorId, uint256 amount) returns()
func (_Staking *StakingTransactorSession) LogConfirmAuction(newValidatorId *big.Int, oldValidatorId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.LogConfirmAuction(&_Staking.TransactOpts, newValidatorId, oldValidatorId, amount)
}

// LogDelegatorClaimRewards is a paid mutator transaction binding the contract method 0xb7721d2d.
//
// Solidity: function logDelegatorClaimRewards(uint256 validatorId, address user, uint256 rewards) returns()
func (_Staking *StakingTransactor) LogDelegatorClaimRewards(opts *bind.TransactOpts, validatorId *big.Int, user common.Address, rewards *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "logDelegatorClaimRewards", validatorId, user, rewards)
}

// LogDelegatorClaimRewards is a paid mutator transaction binding the contract method 0xb7721d2d.
//
// Solidity: function logDelegatorClaimRewards(uint256 validatorId, address user, uint256 rewards) returns()
func (_Staking *StakingSession) LogDelegatorClaimRewards(validatorId *big.Int, user common.Address, rewards *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.LogDelegatorClaimRewards(&_Staking.TransactOpts, validatorId, user, rewards)
}

// LogDelegatorClaimRewards is a paid mutator transaction binding the contract method 0xb7721d2d.
//
// Solidity: function logDelegatorClaimRewards(uint256 validatorId, address user, uint256 rewards) returns()
func (_Staking *StakingTransactorSession) LogDelegatorClaimRewards(validatorId *big.Int, user common.Address, rewards *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.LogDelegatorClaimRewards(&_Staking.TransactOpts, validatorId, user, rewards)
}

// LogDelegatorRestaked is a paid mutator transaction binding the contract method 0x7f88a957.
//
// Solidity: function logDelegatorRestaked(uint256 validatorId, address user, uint256 totalStaked) returns()
func (_Staking *StakingTransactor) LogDelegatorRestaked(opts *bind.TransactOpts, validatorId *big.Int, user common.Address, totalStaked *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "logDelegatorRestaked", validatorId, user, totalStaked)
}

// LogDelegatorRestaked is a paid mutator transaction binding the contract method 0x7f88a957.
//
// Solidity: function logDelegatorRestaked(uint256 validatorId, address user, uint256 totalStaked) returns()
func (_Staking *StakingSession) LogDelegatorRestaked(validatorId *big.Int, user common.Address, totalStaked *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.LogDelegatorRestaked(&_Staking.TransactOpts, validatorId, user, totalStaked)
}

// LogDelegatorRestaked is a paid mutator transaction binding the contract method 0x7f88a957.
//
// Solidity: function logDelegatorRestaked(uint256 validatorId, address user, uint256 totalStaked) returns()
func (_Staking *StakingTransactorSession) LogDelegatorRestaked(validatorId *big.Int, user common.Address, totalStaked *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.LogDelegatorRestaked(&_Staking.TransactOpts, validatorId, user, totalStaked)
}

// LogDelegatorUnstaked is a paid mutator transaction binding the contract method 0x605be9be.
//
// Solidity: function logDelegatorUnstaked(uint256 validatorId, address user, uint256 amount) returns()
func (_Staking *StakingTransactor) LogDelegatorUnstaked(opts *bind.TransactOpts, validatorId *big.Int, user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "logDelegatorUnstaked", validatorId, user, amount)
}

// LogDelegatorUnstaked is a paid mutator transaction binding the contract method 0x605be9be.
//
// Solidity: function logDelegatorUnstaked(uint256 validatorId, address user, uint256 amount) returns()
func (_Staking *StakingSession) LogDelegatorUnstaked(validatorId *big.Int, user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.LogDelegatorUnstaked(&_Staking.TransactOpts, validatorId, user, amount)
}

// LogDelegatorUnstaked is a paid mutator transaction binding the contract method 0x605be9be.
//
// Solidity: function logDelegatorUnstaked(uint256 validatorId, address user, uint256 amount) returns()
func (_Staking *StakingTransactorSession) LogDelegatorUnstaked(validatorId *big.Int, user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.LogDelegatorUnstaked(&_Staking.TransactOpts, validatorId, user, amount)
}

// LogDynastyValueChange is a paid mutator transaction binding the contract method 0xa0e300a6.
//
// Solidity: function logDynastyValueChange(uint256 newDynasty, uint256 oldDynasty) returns()
func (_Staking *StakingTransactor) LogDynastyValueChange(opts *bind.TransactOpts, newDynasty *big.Int, oldDynasty *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "logDynastyValueChange", newDynasty, oldDynasty)
}

// LogDynastyValueChange is a paid mutator transaction binding the contract method 0xa0e300a6.
//
// Solidity: function logDynastyValueChange(uint256 newDynasty, uint256 oldDynasty) returns()
func (_Staking *StakingSession) LogDynastyValueChange(newDynasty *big.Int, oldDynasty *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.LogDynastyValueChange(&_Staking.TransactOpts, newDynasty, oldDynasty)
}

// LogDynastyValueChange is a paid mutator transaction binding the contract method 0xa0e300a6.
//
// Solidity: function logDynastyValueChange(uint256 newDynasty, uint256 oldDynasty) returns()
func (_Staking *StakingTransactorSession) LogDynastyValueChange(newDynasty *big.Int, oldDynasty *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.LogDynastyValueChange(&_Staking.TransactOpts, newDynasty, oldDynasty)
}

// LogJailed is a paid mutator transaction binding the contract method 0x81dc101b.
//
// Solidity: function logJailed(uint256 validatorId, uint256 exitEpoch, address signer) returns()
func (_Staking *StakingTransactor) LogJailed(opts *bind.TransactOpts, validatorId *big.Int, exitEpoch *big.Int, signer common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "logJailed", validatorId, exitEpoch, signer)
}

// LogJailed is a paid mutator transaction binding the contract method 0x81dc101b.
//
// Solidity: function logJailed(uint256 validatorId, uint256 exitEpoch, address signer) returns()
func (_Staking *StakingSession) LogJailed(validatorId *big.Int, exitEpoch *big.Int, signer common.Address) (*types.Transaction, error) {
	return _Staking.Contract.LogJailed(&_Staking.TransactOpts, validatorId, exitEpoch, signer)
}

// LogJailed is a paid mutator transaction binding the contract method 0x81dc101b.
//
// Solidity: function logJailed(uint256 validatorId, uint256 exitEpoch, address signer) returns()
func (_Staking *StakingTransactorSession) LogJailed(validatorId *big.Int, exitEpoch *big.Int, signer common.Address) (*types.Transaction, error) {
	return _Staking.Contract.LogJailed(&_Staking.TransactOpts, validatorId, exitEpoch, signer)
}

// LogProposerBonusChange is a paid mutator transaction binding the contract method 0xa3b1d8cb.
//
// Solidity: function logProposerBonusChange(uint256 newProposerBonus, uint256 oldProposerBonus) returns()
func (_Staking *StakingTransactor) LogProposerBonusChange(opts *bind.TransactOpts, newProposerBonus *big.Int, oldProposerBonus *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "logProposerBonusChange", newProposerBonus, oldProposerBonus)
}

// LogProposerBonusChange is a paid mutator transaction binding the contract method 0xa3b1d8cb.
//
// Solidity: function logProposerBonusChange(uint256 newProposerBonus, uint256 oldProposerBonus) returns()
func (_Staking *StakingSession) LogProposerBonusChange(newProposerBonus *big.Int, oldProposerBonus *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.LogProposerBonusChange(&_Staking.TransactOpts, newProposerBonus, oldProposerBonus)
}

// LogProposerBonusChange is a paid mutator transaction binding the contract method 0xa3b1d8cb.
//
// Solidity: function logProposerBonusChange(uint256 newProposerBonus, uint256 oldProposerBonus) returns()
func (_Staking *StakingTransactorSession) LogProposerBonusChange(newProposerBonus *big.Int, oldProposerBonus *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.LogProposerBonusChange(&_Staking.TransactOpts, newProposerBonus, oldProposerBonus)
}

// LogRestaked is a paid mutator transaction binding the contract method 0x5616a7cc.
//
// Solidity: function logRestaked(uint256 validatorId, uint256 amount, uint256 total) returns()
func (_Staking *StakingTransactor) LogRestaked(opts *bind.TransactOpts, validatorId *big.Int, amount *big.Int, total *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "logRestaked", validatorId, amount, total)
}

// LogRestaked is a paid mutator transaction binding the contract method 0x5616a7cc.
//
// Solidity: function logRestaked(uint256 validatorId, uint256 amount, uint256 total) returns()
func (_Staking *StakingSession) LogRestaked(validatorId *big.Int, amount *big.Int, total *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.LogRestaked(&_Staking.TransactOpts, validatorId, amount, total)
}

// LogRestaked is a paid mutator transaction binding the contract method 0x5616a7cc.
//
// Solidity: function logRestaked(uint256 validatorId, uint256 amount, uint256 total) returns()
func (_Staking *StakingTransactorSession) LogRestaked(validatorId *big.Int, amount *big.Int, total *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.LogRestaked(&_Staking.TransactOpts, validatorId, amount, total)
}

// LogRewardUpdate is a paid mutator transaction binding the contract method 0xb6fa74c4.
//
// Solidity: function logRewardUpdate(uint256 newReward, uint256 oldReward) returns()
func (_Staking *StakingTransactor) LogRewardUpdate(opts *bind.TransactOpts, newReward *big.Int, oldReward *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "logRewardUpdate", newReward, oldReward)
}

// LogRewardUpdate is a paid mutator transaction binding the contract method 0xb6fa74c4.
//
// Solidity: function logRewardUpdate(uint256 newReward, uint256 oldReward) returns()
func (_Staking *StakingSession) LogRewardUpdate(newReward *big.Int, oldReward *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.LogRewardUpdate(&_Staking.TransactOpts, newReward, oldReward)
}

// LogRewardUpdate is a paid mutator transaction binding the contract method 0xb6fa74c4.
//
// Solidity: function logRewardUpdate(uint256 newReward, uint256 oldReward) returns()
func (_Staking *StakingTransactorSession) LogRewardUpdate(newReward *big.Int, oldReward *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.LogRewardUpdate(&_Staking.TransactOpts, newReward, oldReward)
}

// LogShareBurned is a paid mutator transaction binding the contract method 0xf1382b53.
//
// Solidity: function logShareBurned(uint256 validatorId, address user, uint256 amount, uint256 tokens) returns()
func (_Staking *StakingTransactor) LogShareBurned(opts *bind.TransactOpts, validatorId *big.Int, user common.Address, amount *big.Int, tokens *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "logShareBurned", validatorId, user, amount, tokens)
}

// LogShareBurned is a paid mutator transaction binding the contract method 0xf1382b53.
//
// Solidity: function logShareBurned(uint256 validatorId, address user, uint256 amount, uint256 tokens) returns()
func (_Staking *StakingSession) LogShareBurned(validatorId *big.Int, user common.Address, amount *big.Int, tokens *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.LogShareBurned(&_Staking.TransactOpts, validatorId, user, amount, tokens)
}

// LogShareBurned is a paid mutator transaction binding the contract method 0xf1382b53.
//
// Solidity: function logShareBurned(uint256 validatorId, address user, uint256 amount, uint256 tokens) returns()
func (_Staking *StakingTransactorSession) LogShareBurned(validatorId *big.Int, user common.Address, amount *big.Int, tokens *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.LogShareBurned(&_Staking.TransactOpts, validatorId, user, amount, tokens)
}

// LogShareMinted is a paid mutator transaction binding the contract method 0xc69d0573.
//
// Solidity: function logShareMinted(uint256 validatorId, address user, uint256 amount, uint256 tokens) returns()
func (_Staking *StakingTransactor) LogShareMinted(opts *bind.TransactOpts, validatorId *big.Int, user common.Address, amount *big.Int, tokens *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "logShareMinted", validatorId, user, amount, tokens)
}

// LogShareMinted is a paid mutator transaction binding the contract method 0xc69d0573.
//
// Solidity: function logShareMinted(uint256 validatorId, address user, uint256 amount, uint256 tokens) returns()
func (_Staking *StakingSession) LogShareMinted(validatorId *big.Int, user common.Address, amount *big.Int, tokens *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.LogShareMinted(&_Staking.TransactOpts, validatorId, user, amount, tokens)
}

// LogShareMinted is a paid mutator transaction binding the contract method 0xc69d0573.
//
// Solidity: function logShareMinted(uint256 validatorId, address user, uint256 amount, uint256 tokens) returns()
func (_Staking *StakingTransactorSession) LogShareMinted(validatorId *big.Int, user common.Address, amount *big.Int, tokens *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.LogShareMinted(&_Staking.TransactOpts, validatorId, user, amount, tokens)
}

// LogSignerChange is a paid mutator transaction binding the contract method 0xb80fbce5.
//
// Solidity: function logSignerChange(uint256 validatorId, address oldSigner, address newSigner, bytes signerPubkey) returns()
func (_Staking *StakingTransactor) LogSignerChange(opts *bind.TransactOpts, validatorId *big.Int, oldSigner common.Address, newSigner common.Address, signerPubkey []byte) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "logSignerChange", validatorId, oldSigner, newSigner, signerPubkey)
}

// LogSignerChange is a paid mutator transaction binding the contract method 0xb80fbce5.
//
// Solidity: function logSignerChange(uint256 validatorId, address oldSigner, address newSigner, bytes signerPubkey) returns()
func (_Staking *StakingSession) LogSignerChange(validatorId *big.Int, oldSigner common.Address, newSigner common.Address, signerPubkey []byte) (*types.Transaction, error) {
	return _Staking.Contract.LogSignerChange(&_Staking.TransactOpts, validatorId, oldSigner, newSigner, signerPubkey)
}

// LogSignerChange is a paid mutator transaction binding the contract method 0xb80fbce5.
//
// Solidity: function logSignerChange(uint256 validatorId, address oldSigner, address newSigner, bytes signerPubkey) returns()
func (_Staking *StakingTransactorSession) LogSignerChange(validatorId *big.Int, oldSigner common.Address, newSigner common.Address, signerPubkey []byte) (*types.Transaction, error) {
	return _Staking.Contract.LogSignerChange(&_Staking.TransactOpts, validatorId, oldSigner, newSigner, signerPubkey)
}

// LogSlashed is a paid mutator transaction binding the contract method 0xfb77c94e.
//
// Solidity: function logSlashed(uint256 nonce, uint256 amount) returns()
func (_Staking *StakingTransactor) LogSlashed(opts *bind.TransactOpts, nonce *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "logSlashed", nonce, amount)
}

// LogSlashed is a paid mutator transaction binding the contract method 0xfb77c94e.
//
// Solidity: function logSlashed(uint256 nonce, uint256 amount) returns()
func (_Staking *StakingSession) LogSlashed(nonce *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.LogSlashed(&_Staking.TransactOpts, nonce, amount)
}

// LogSlashed is a paid mutator transaction binding the contract method 0xfb77c94e.
//
// Solidity: function logSlashed(uint256 nonce, uint256 amount) returns()
func (_Staking *StakingTransactorSession) LogSlashed(nonce *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.LogSlashed(&_Staking.TransactOpts, nonce, amount)
}

// LogStakeUpdate is a paid mutator transaction binding the contract method 0x532e19a9.
//
// Solidity: function logStakeUpdate(uint256 validatorId) returns()
func (_Staking *StakingTransactor) LogStakeUpdate(opts *bind.TransactOpts, validatorId *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "logStakeUpdate", validatorId)
}

// LogStakeUpdate is a paid mutator transaction binding the contract method 0x532e19a9.
//
// Solidity: function logStakeUpdate(uint256 validatorId) returns()
func (_Staking *StakingSession) LogStakeUpdate(validatorId *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.LogStakeUpdate(&_Staking.TransactOpts, validatorId)
}

// LogStakeUpdate is a paid mutator transaction binding the contract method 0x532e19a9.
//
// Solidity: function logStakeUpdate(uint256 validatorId) returns()
func (_Staking *StakingTransactorSession) LogStakeUpdate(validatorId *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.LogStakeUpdate(&_Staking.TransactOpts, validatorId)
}

// LogStaked is a paid mutator transaction binding the contract method 0x33a8383c.
//
// Solidity: function logStaked(address signer, bytes signerPubkey, uint256 validatorId, uint256 activationEpoch, uint256 amount, uint256 total) returns()
func (_Staking *StakingTransactor) LogStaked(opts *bind.TransactOpts, signer common.Address, signerPubkey []byte, validatorId *big.Int, activationEpoch *big.Int, amount *big.Int, total *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "logStaked", signer, signerPubkey, validatorId, activationEpoch, amount, total)
}

// LogStaked is a paid mutator transaction binding the contract method 0x33a8383c.
//
// Solidity: function logStaked(address signer, bytes signerPubkey, uint256 validatorId, uint256 activationEpoch, uint256 amount, uint256 total) returns()
func (_Staking *StakingSession) LogStaked(signer common.Address, signerPubkey []byte, validatorId *big.Int, activationEpoch *big.Int, amount *big.Int, total *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.LogStaked(&_Staking.TransactOpts, signer, signerPubkey, validatorId, activationEpoch, amount, total)
}

// LogStaked is a paid mutator transaction binding the contract method 0x33a8383c.
//
// Solidity: function logStaked(address signer, bytes signerPubkey, uint256 validatorId, uint256 activationEpoch, uint256 amount, uint256 total) returns()
func (_Staking *StakingTransactorSession) LogStaked(signer common.Address, signerPubkey []byte, validatorId *big.Int, activationEpoch *big.Int, amount *big.Int, total *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.LogStaked(&_Staking.TransactOpts, signer, signerPubkey, validatorId, activationEpoch, amount, total)
}

// LogStartAuction is a paid mutator transaction binding the contract method 0x0934a6df.
//
// Solidity: function logStartAuction(uint256 validatorId, uint256 amount, uint256 auctionAmount) returns()
func (_Staking *StakingTransactor) LogStartAuction(opts *bind.TransactOpts, validatorId *big.Int, amount *big.Int, auctionAmount *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "logStartAuction", validatorId, amount, auctionAmount)
}

// LogStartAuction is a paid mutator transaction binding the contract method 0x0934a6df.
//
// Solidity: function logStartAuction(uint256 validatorId, uint256 amount, uint256 auctionAmount) returns()
func (_Staking *StakingSession) LogStartAuction(validatorId *big.Int, amount *big.Int, auctionAmount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.LogStartAuction(&_Staking.TransactOpts, validatorId, amount, auctionAmount)
}

// LogStartAuction is a paid mutator transaction binding the contract method 0x0934a6df.
//
// Solidity: function logStartAuction(uint256 validatorId, uint256 amount, uint256 auctionAmount) returns()
func (_Staking *StakingTransactorSession) LogStartAuction(validatorId *big.Int, amount *big.Int, auctionAmount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.LogStartAuction(&_Staking.TransactOpts, validatorId, amount, auctionAmount)
}

// LogThresholdChange is a paid mutator transaction binding the contract method 0xf1980a50.
//
// Solidity: function logThresholdChange(uint256 newThreshold, uint256 oldThreshold) returns()
func (_Staking *StakingTransactor) LogThresholdChange(opts *bind.TransactOpts, newThreshold *big.Int, oldThreshold *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "logThresholdChange", newThreshold, oldThreshold)
}

// LogThresholdChange is a paid mutator transaction binding the contract method 0xf1980a50.
//
// Solidity: function logThresholdChange(uint256 newThreshold, uint256 oldThreshold) returns()
func (_Staking *StakingSession) LogThresholdChange(newThreshold *big.Int, oldThreshold *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.LogThresholdChange(&_Staking.TransactOpts, newThreshold, oldThreshold)
}

// LogThresholdChange is a paid mutator transaction binding the contract method 0xf1980a50.
//
// Solidity: function logThresholdChange(uint256 newThreshold, uint256 oldThreshold) returns()
func (_Staking *StakingTransactorSession) LogThresholdChange(newThreshold *big.Int, oldThreshold *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.LogThresholdChange(&_Staking.TransactOpts, newThreshold, oldThreshold)
}

// LogTopUpFee is a paid mutator transaction binding the contract method 0xa449d795.
//
// Solidity: function logTopUpFee(address user, uint256 fee) returns()
func (_Staking *StakingTransactor) LogTopUpFee(opts *bind.TransactOpts, user common.Address, fee *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "logTopUpFee", user, fee)
}

// LogTopUpFee is a paid mutator transaction binding the contract method 0xa449d795.
//
// Solidity: function logTopUpFee(address user, uint256 fee) returns()
func (_Staking *StakingSession) LogTopUpFee(user common.Address, fee *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.LogTopUpFee(&_Staking.TransactOpts, user, fee)
}

// LogTopUpFee is a paid mutator transaction binding the contract method 0xa449d795.
//
// Solidity: function logTopUpFee(address user, uint256 fee) returns()
func (_Staking *StakingTransactorSession) LogTopUpFee(user common.Address, fee *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.LogTopUpFee(&_Staking.TransactOpts, user, fee)
}

// LogUnjailed is a paid mutator transaction binding the contract method 0xf92ec5af.
//
// Solidity: function logUnjailed(uint256 validatorId, address signer) returns()
func (_Staking *StakingTransactor) LogUnjailed(opts *bind.TransactOpts, validatorId *big.Int, signer common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "logUnjailed", validatorId, signer)
}

// LogUnjailed is a paid mutator transaction binding the contract method 0xf92ec5af.
//
// Solidity: function logUnjailed(uint256 validatorId, address signer) returns()
func (_Staking *StakingSession) LogUnjailed(validatorId *big.Int, signer common.Address) (*types.Transaction, error) {
	return _Staking.Contract.LogUnjailed(&_Staking.TransactOpts, validatorId, signer)
}

// LogUnjailed is a paid mutator transaction binding the contract method 0xf92ec5af.
//
// Solidity: function logUnjailed(uint256 validatorId, address signer) returns()
func (_Staking *StakingTransactorSession) LogUnjailed(validatorId *big.Int, signer common.Address) (*types.Transaction, error) {
	return _Staking.Contract.LogUnjailed(&_Staking.TransactOpts, validatorId, signer)
}

// LogUnstakeInit is a paid mutator transaction binding the contract method 0x5e04d483.
//
// Solidity: function logUnstakeInit(address user, uint256 validatorId, uint256 deactivationEpoch, uint256 amount) returns()
func (_Staking *StakingTransactor) LogUnstakeInit(opts *bind.TransactOpts, user common.Address, validatorId *big.Int, deactivationEpoch *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "logUnstakeInit", user, validatorId, deactivationEpoch, amount)
}

// LogUnstakeInit is a paid mutator transaction binding the contract method 0x5e04d483.
//
// Solidity: function logUnstakeInit(address user, uint256 validatorId, uint256 deactivationEpoch, uint256 amount) returns()
func (_Staking *StakingSession) LogUnstakeInit(user common.Address, validatorId *big.Int, deactivationEpoch *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.LogUnstakeInit(&_Staking.TransactOpts, user, validatorId, deactivationEpoch, amount)
}

// LogUnstakeInit is a paid mutator transaction binding the contract method 0x5e04d483.
//
// Solidity: function logUnstakeInit(address user, uint256 validatorId, uint256 deactivationEpoch, uint256 amount) returns()
func (_Staking *StakingTransactorSession) LogUnstakeInit(user common.Address, validatorId *big.Int, deactivationEpoch *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.LogUnstakeInit(&_Staking.TransactOpts, user, validatorId, deactivationEpoch, amount)
}

// LogUnstaked is a paid mutator transaction binding the contract method 0xae2e26b1.
//
// Solidity: function logUnstaked(address user, uint256 validatorId, uint256 amount, uint256 total) returns()
func (_Staking *StakingTransactor) LogUnstaked(opts *bind.TransactOpts, user common.Address, validatorId *big.Int, amount *big.Int, total *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "logUnstaked", user, validatorId, amount, total)
}

// LogUnstaked is a paid mutator transaction binding the contract method 0xae2e26b1.
//
// Solidity: function logUnstaked(address user, uint256 validatorId, uint256 amount, uint256 total) returns()
func (_Staking *StakingSession) LogUnstaked(user common.Address, validatorId *big.Int, amount *big.Int, total *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.LogUnstaked(&_Staking.TransactOpts, user, validatorId, amount, total)
}

// LogUnstaked is a paid mutator transaction binding the contract method 0xae2e26b1.
//
// Solidity: function logUnstaked(address user, uint256 validatorId, uint256 amount, uint256 total) returns()
func (_Staking *StakingTransactorSession) LogUnstaked(user common.Address, validatorId *big.Int, amount *big.Int, total *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.LogUnstaked(&_Staking.TransactOpts, user, validatorId, amount, total)
}

// LogUpdateCommissionRate is a paid mutator transaction binding the contract method 0xc98cc002.
//
// Solidity: function logUpdateCommissionRate(uint256 validatorId, uint256 newCommissionRate, uint256 oldCommissionRate) returns()
func (_Staking *StakingTransactor) LogUpdateCommissionRate(opts *bind.TransactOpts, validatorId *big.Int, newCommissionRate *big.Int, oldCommissionRate *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "logUpdateCommissionRate", validatorId, newCommissionRate, oldCommissionRate)
}

// LogUpdateCommissionRate is a paid mutator transaction binding the contract method 0xc98cc002.
//
// Solidity: function logUpdateCommissionRate(uint256 validatorId, uint256 newCommissionRate, uint256 oldCommissionRate) returns()
func (_Staking *StakingSession) LogUpdateCommissionRate(validatorId *big.Int, newCommissionRate *big.Int, oldCommissionRate *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.LogUpdateCommissionRate(&_Staking.TransactOpts, validatorId, newCommissionRate, oldCommissionRate)
}

// LogUpdateCommissionRate is a paid mutator transaction binding the contract method 0xc98cc002.
//
// Solidity: function logUpdateCommissionRate(uint256 validatorId, uint256 newCommissionRate, uint256 oldCommissionRate) returns()
func (_Staking *StakingTransactorSession) LogUpdateCommissionRate(validatorId *big.Int, newCommissionRate *big.Int, oldCommissionRate *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.LogUpdateCommissionRate(&_Staking.TransactOpts, validatorId, newCommissionRate, oldCommissionRate)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Staking *StakingTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Staking *StakingSession) RenounceOwnership() (*types.Transaction, error) {
	return _Staking.Contract.RenounceOwnership(&_Staking.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Staking *StakingTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Staking.Contract.RenounceOwnership(&_Staking.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Staking *StakingTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Staking *StakingSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Staking.Contract.TransferOwnership(&_Staking.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Staking *StakingTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Staking.Contract.TransferOwnership(&_Staking.TransactOpts, newOwner)
}

// UpdateNonce is a paid mutator transaction binding the contract method 0xeae3f749.
//
// Solidity: function updateNonce(uint256[] validatorIds, uint256[] nonces) returns()
func (_Staking *StakingTransactor) UpdateNonce(opts *bind.TransactOpts, validatorIds []*big.Int, nonces []*big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "updateNonce", validatorIds, nonces)
}

// UpdateNonce is a paid mutator transaction binding the contract method 0xeae3f749.
//
// Solidity: function updateNonce(uint256[] validatorIds, uint256[] nonces) returns()
func (_Staking *StakingSession) UpdateNonce(validatorIds []*big.Int, nonces []*big.Int) (*types.Transaction, error) {
	return _Staking.Contract.UpdateNonce(&_Staking.TransactOpts, validatorIds, nonces)
}

// UpdateNonce is a paid mutator transaction binding the contract method 0xeae3f749.
//
// Solidity: function updateNonce(uint256[] validatorIds, uint256[] nonces) returns()
func (_Staking *StakingTransactorSession) UpdateNonce(validatorIds []*big.Int, nonces []*big.Int) (*types.Transaction, error) {
	return _Staking.Contract.UpdateNonce(&_Staking.TransactOpts, validatorIds, nonces)
}

// StakingClaimFeeIterator is returned from FilterClaimFee and is used to iterate over the raw logs and unpacked data for ClaimFee events raised by the Staking contract.
type StakingClaimFeeIterator struct {
	Event *StakingClaimFee // Event containing the contract specifics and raw log

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
func (it *StakingClaimFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingClaimFee)
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
		it.Event = new(StakingClaimFee)
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
func (it *StakingClaimFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingClaimFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingClaimFee represents a ClaimFee event raised by the Staking contract.
type StakingClaimFee struct {
	User common.Address
	Fee  *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterClaimFee is a free log retrieval operation binding the contract event 0xf40b9ca28516abde647ef8ed0e7b155e16347eb4d8dd6eb29989ed2c0c3d27e8.
//
// Solidity: event ClaimFee(address indexed user, uint256 indexed fee)
func (_Staking *StakingFilterer) FilterClaimFee(opts *bind.FilterOpts, user []common.Address, fee []*big.Int) (*StakingClaimFeeIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "ClaimFee", userRule, feeRule)
	if err != nil {
		return nil, err
	}
	return &StakingClaimFeeIterator{contract: _Staking.contract, event: "ClaimFee", logs: logs, sub: sub}, nil
}

// WatchClaimFee is a free log subscription operation binding the contract event 0xf40b9ca28516abde647ef8ed0e7b155e16347eb4d8dd6eb29989ed2c0c3d27e8.
//
// Solidity: event ClaimFee(address indexed user, uint256 indexed fee)
func (_Staking *StakingFilterer) WatchClaimFee(opts *bind.WatchOpts, sink chan<- *StakingClaimFee, user []common.Address, fee []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "ClaimFee", userRule, feeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingClaimFee)
				if err := _Staking.contract.UnpackLog(event, "ClaimFee", log); err != nil {
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

// ParseClaimFee is a log parse operation binding the contract event 0xf40b9ca28516abde647ef8ed0e7b155e16347eb4d8dd6eb29989ed2c0c3d27e8.
//
// Solidity: event ClaimFee(address indexed user, uint256 indexed fee)
func (_Staking *StakingFilterer) ParseClaimFee(log types.Log) (*StakingClaimFee, error) {
	event := new(StakingClaimFee)
	if err := _Staking.contract.UnpackLog(event, "ClaimFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingClaimRewardsIterator is returned from FilterClaimRewards and is used to iterate over the raw logs and unpacked data for ClaimRewards events raised by the Staking contract.
type StakingClaimRewardsIterator struct {
	Event *StakingClaimRewards // Event containing the contract specifics and raw log

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
func (it *StakingClaimRewardsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingClaimRewards)
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
		it.Event = new(StakingClaimRewards)
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
func (it *StakingClaimRewardsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingClaimRewardsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingClaimRewards represents a ClaimRewards event raised by the Staking contract.
type StakingClaimRewards struct {
	ValidatorId *big.Int
	Amount      *big.Int
	TotalAmount *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterClaimRewards is a free log retrieval operation binding the contract event 0x41e5e4590cfcde2f03ee9281c54d03acad8adffb83f8310d66b894532470ba35.
//
// Solidity: event ClaimRewards(uint256 indexed validatorId, uint256 indexed amount, uint256 indexed totalAmount)
func (_Staking *StakingFilterer) FilterClaimRewards(opts *bind.FilterOpts, validatorId []*big.Int, amount []*big.Int, totalAmount []*big.Int) (*StakingClaimRewardsIterator, error) {

	var validatorIdRule []interface{}
	for _, validatorIdItem := range validatorId {
		validatorIdRule = append(validatorIdRule, validatorIdItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}
	var totalAmountRule []interface{}
	for _, totalAmountItem := range totalAmount {
		totalAmountRule = append(totalAmountRule, totalAmountItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "ClaimRewards", validatorIdRule, amountRule, totalAmountRule)
	if err != nil {
		return nil, err
	}
	return &StakingClaimRewardsIterator{contract: _Staking.contract, event: "ClaimRewards", logs: logs, sub: sub}, nil
}

// WatchClaimRewards is a free log subscription operation binding the contract event 0x41e5e4590cfcde2f03ee9281c54d03acad8adffb83f8310d66b894532470ba35.
//
// Solidity: event ClaimRewards(uint256 indexed validatorId, uint256 indexed amount, uint256 indexed totalAmount)
func (_Staking *StakingFilterer) WatchClaimRewards(opts *bind.WatchOpts, sink chan<- *StakingClaimRewards, validatorId []*big.Int, amount []*big.Int, totalAmount []*big.Int) (event.Subscription, error) {

	var validatorIdRule []interface{}
	for _, validatorIdItem := range validatorId {
		validatorIdRule = append(validatorIdRule, validatorIdItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}
	var totalAmountRule []interface{}
	for _, totalAmountItem := range totalAmount {
		totalAmountRule = append(totalAmountRule, totalAmountItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "ClaimRewards", validatorIdRule, amountRule, totalAmountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingClaimRewards)
				if err := _Staking.contract.UnpackLog(event, "ClaimRewards", log); err != nil {
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

// ParseClaimRewards is a log parse operation binding the contract event 0x41e5e4590cfcde2f03ee9281c54d03acad8adffb83f8310d66b894532470ba35.
//
// Solidity: event ClaimRewards(uint256 indexed validatorId, uint256 indexed amount, uint256 indexed totalAmount)
func (_Staking *StakingFilterer) ParseClaimRewards(log types.Log) (*StakingClaimRewards, error) {
	event := new(StakingClaimRewards)
	if err := _Staking.contract.UnpackLog(event, "ClaimRewards", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingConfirmAuctionIterator is returned from FilterConfirmAuction and is used to iterate over the raw logs and unpacked data for ConfirmAuction events raised by the Staking contract.
type StakingConfirmAuctionIterator struct {
	Event *StakingConfirmAuction // Event containing the contract specifics and raw log

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
func (it *StakingConfirmAuctionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingConfirmAuction)
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
		it.Event = new(StakingConfirmAuction)
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
func (it *StakingConfirmAuctionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingConfirmAuctionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingConfirmAuction represents a ConfirmAuction event raised by the Staking contract.
type StakingConfirmAuction struct {
	NewValidatorId *big.Int
	OldValidatorId *big.Int
	Amount         *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterConfirmAuction is a free log retrieval operation binding the contract event 0x1002381ecf76700f6f0ab4c90b9f523e39df7b0482b71ec63cf62cf854120470.
//
// Solidity: event ConfirmAuction(uint256 indexed newValidatorId, uint256 indexed oldValidatorId, uint256 indexed amount)
func (_Staking *StakingFilterer) FilterConfirmAuction(opts *bind.FilterOpts, newValidatorId []*big.Int, oldValidatorId []*big.Int, amount []*big.Int) (*StakingConfirmAuctionIterator, error) {

	var newValidatorIdRule []interface{}
	for _, newValidatorIdItem := range newValidatorId {
		newValidatorIdRule = append(newValidatorIdRule, newValidatorIdItem)
	}
	var oldValidatorIdRule []interface{}
	for _, oldValidatorIdItem := range oldValidatorId {
		oldValidatorIdRule = append(oldValidatorIdRule, oldValidatorIdItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "ConfirmAuction", newValidatorIdRule, oldValidatorIdRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &StakingConfirmAuctionIterator{contract: _Staking.contract, event: "ConfirmAuction", logs: logs, sub: sub}, nil
}

// WatchConfirmAuction is a free log subscription operation binding the contract event 0x1002381ecf76700f6f0ab4c90b9f523e39df7b0482b71ec63cf62cf854120470.
//
// Solidity: event ConfirmAuction(uint256 indexed newValidatorId, uint256 indexed oldValidatorId, uint256 indexed amount)
func (_Staking *StakingFilterer) WatchConfirmAuction(opts *bind.WatchOpts, sink chan<- *StakingConfirmAuction, newValidatorId []*big.Int, oldValidatorId []*big.Int, amount []*big.Int) (event.Subscription, error) {

	var newValidatorIdRule []interface{}
	for _, newValidatorIdItem := range newValidatorId {
		newValidatorIdRule = append(newValidatorIdRule, newValidatorIdItem)
	}
	var oldValidatorIdRule []interface{}
	for _, oldValidatorIdItem := range oldValidatorId {
		oldValidatorIdRule = append(oldValidatorIdRule, oldValidatorIdItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "ConfirmAuction", newValidatorIdRule, oldValidatorIdRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingConfirmAuction)
				if err := _Staking.contract.UnpackLog(event, "ConfirmAuction", log); err != nil {
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

// ParseConfirmAuction is a log parse operation binding the contract event 0x1002381ecf76700f6f0ab4c90b9f523e39df7b0482b71ec63cf62cf854120470.
//
// Solidity: event ConfirmAuction(uint256 indexed newValidatorId, uint256 indexed oldValidatorId, uint256 indexed amount)
func (_Staking *StakingFilterer) ParseConfirmAuction(log types.Log) (*StakingConfirmAuction, error) {
	event := new(StakingConfirmAuction)
	if err := _Staking.contract.UnpackLog(event, "ConfirmAuction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingDelegatorClaimedRewardsIterator is returned from FilterDelegatorClaimedRewards and is used to iterate over the raw logs and unpacked data for DelegatorClaimedRewards events raised by the Staking contract.
type StakingDelegatorClaimedRewardsIterator struct {
	Event *StakingDelegatorClaimedRewards // Event containing the contract specifics and raw log

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
func (it *StakingDelegatorClaimedRewardsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingDelegatorClaimedRewards)
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
		it.Event = new(StakingDelegatorClaimedRewards)
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
func (it *StakingDelegatorClaimedRewardsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingDelegatorClaimedRewardsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingDelegatorClaimedRewards represents a DelegatorClaimedRewards event raised by the Staking contract.
type StakingDelegatorClaimedRewards struct {
	ValidatorId *big.Int
	User        common.Address
	Rewards     *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDelegatorClaimedRewards is a free log retrieval operation binding the contract event 0x31d1715032654fde9867c0f095aecce1113049e30b9f4ecbaa6954ed6c63b8df.
//
// Solidity: event DelegatorClaimedRewards(uint256 indexed validatorId, address indexed user, uint256 indexed rewards)
func (_Staking *StakingFilterer) FilterDelegatorClaimedRewards(opts *bind.FilterOpts, validatorId []*big.Int, user []common.Address, rewards []*big.Int) (*StakingDelegatorClaimedRewardsIterator, error) {

	var validatorIdRule []interface{}
	for _, validatorIdItem := range validatorId {
		validatorIdRule = append(validatorIdRule, validatorIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var rewardsRule []interface{}
	for _, rewardsItem := range rewards {
		rewardsRule = append(rewardsRule, rewardsItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "DelegatorClaimedRewards", validatorIdRule, userRule, rewardsRule)
	if err != nil {
		return nil, err
	}
	return &StakingDelegatorClaimedRewardsIterator{contract: _Staking.contract, event: "DelegatorClaimedRewards", logs: logs, sub: sub}, nil
}

// WatchDelegatorClaimedRewards is a free log subscription operation binding the contract event 0x31d1715032654fde9867c0f095aecce1113049e30b9f4ecbaa6954ed6c63b8df.
//
// Solidity: event DelegatorClaimedRewards(uint256 indexed validatorId, address indexed user, uint256 indexed rewards)
func (_Staking *StakingFilterer) WatchDelegatorClaimedRewards(opts *bind.WatchOpts, sink chan<- *StakingDelegatorClaimedRewards, validatorId []*big.Int, user []common.Address, rewards []*big.Int) (event.Subscription, error) {

	var validatorIdRule []interface{}
	for _, validatorIdItem := range validatorId {
		validatorIdRule = append(validatorIdRule, validatorIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var rewardsRule []interface{}
	for _, rewardsItem := range rewards {
		rewardsRule = append(rewardsRule, rewardsItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "DelegatorClaimedRewards", validatorIdRule, userRule, rewardsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingDelegatorClaimedRewards)
				if err := _Staking.contract.UnpackLog(event, "DelegatorClaimedRewards", log); err != nil {
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

// ParseDelegatorClaimedRewards is a log parse operation binding the contract event 0x31d1715032654fde9867c0f095aecce1113049e30b9f4ecbaa6954ed6c63b8df.
//
// Solidity: event DelegatorClaimedRewards(uint256 indexed validatorId, address indexed user, uint256 indexed rewards)
func (_Staking *StakingFilterer) ParseDelegatorClaimedRewards(log types.Log) (*StakingDelegatorClaimedRewards, error) {
	event := new(StakingDelegatorClaimedRewards)
	if err := _Staking.contract.UnpackLog(event, "DelegatorClaimedRewards", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingDelegatorRestakedIterator is returned from FilterDelegatorRestaked and is used to iterate over the raw logs and unpacked data for DelegatorRestaked events raised by the Staking contract.
type StakingDelegatorRestakedIterator struct {
	Event *StakingDelegatorRestaked // Event containing the contract specifics and raw log

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
func (it *StakingDelegatorRestakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingDelegatorRestaked)
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
		it.Event = new(StakingDelegatorRestaked)
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
func (it *StakingDelegatorRestakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingDelegatorRestakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingDelegatorRestaked represents a DelegatorRestaked event raised by the Staking contract.
type StakingDelegatorRestaked struct {
	ValidatorId *big.Int
	User        common.Address
	TotalStaked *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDelegatorRestaked is a free log retrieval operation binding the contract event 0x0f9ccdda16b467e719059c85ebd8383fcb7f8ffa5576629fe3b842836e04dad1.
//
// Solidity: event DelegatorRestaked(uint256 indexed validatorId, address indexed user, uint256 indexed totalStaked)
func (_Staking *StakingFilterer) FilterDelegatorRestaked(opts *bind.FilterOpts, validatorId []*big.Int, user []common.Address, totalStaked []*big.Int) (*StakingDelegatorRestakedIterator, error) {

	var validatorIdRule []interface{}
	for _, validatorIdItem := range validatorId {
		validatorIdRule = append(validatorIdRule, validatorIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var totalStakedRule []interface{}
	for _, totalStakedItem := range totalStaked {
		totalStakedRule = append(totalStakedRule, totalStakedItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "DelegatorRestaked", validatorIdRule, userRule, totalStakedRule)
	if err != nil {
		return nil, err
	}
	return &StakingDelegatorRestakedIterator{contract: _Staking.contract, event: "DelegatorRestaked", logs: logs, sub: sub}, nil
}

// WatchDelegatorRestaked is a free log subscription operation binding the contract event 0x0f9ccdda16b467e719059c85ebd8383fcb7f8ffa5576629fe3b842836e04dad1.
//
// Solidity: event DelegatorRestaked(uint256 indexed validatorId, address indexed user, uint256 indexed totalStaked)
func (_Staking *StakingFilterer) WatchDelegatorRestaked(opts *bind.WatchOpts, sink chan<- *StakingDelegatorRestaked, validatorId []*big.Int, user []common.Address, totalStaked []*big.Int) (event.Subscription, error) {

	var validatorIdRule []interface{}
	for _, validatorIdItem := range validatorId {
		validatorIdRule = append(validatorIdRule, validatorIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var totalStakedRule []interface{}
	for _, totalStakedItem := range totalStaked {
		totalStakedRule = append(totalStakedRule, totalStakedItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "DelegatorRestaked", validatorIdRule, userRule, totalStakedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingDelegatorRestaked)
				if err := _Staking.contract.UnpackLog(event, "DelegatorRestaked", log); err != nil {
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

// ParseDelegatorRestaked is a log parse operation binding the contract event 0x0f9ccdda16b467e719059c85ebd8383fcb7f8ffa5576629fe3b842836e04dad1.
//
// Solidity: event DelegatorRestaked(uint256 indexed validatorId, address indexed user, uint256 indexed totalStaked)
func (_Staking *StakingFilterer) ParseDelegatorRestaked(log types.Log) (*StakingDelegatorRestaked, error) {
	event := new(StakingDelegatorRestaked)
	if err := _Staking.contract.UnpackLog(event, "DelegatorRestaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingDelegatorUnstakedIterator is returned from FilterDelegatorUnstaked and is used to iterate over the raw logs and unpacked data for DelegatorUnstaked events raised by the Staking contract.
type StakingDelegatorUnstakedIterator struct {
	Event *StakingDelegatorUnstaked // Event containing the contract specifics and raw log

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
func (it *StakingDelegatorUnstakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingDelegatorUnstaked)
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
		it.Event = new(StakingDelegatorUnstaked)
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
func (it *StakingDelegatorUnstakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingDelegatorUnstakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingDelegatorUnstaked represents a DelegatorUnstaked event raised by the Staking contract.
type StakingDelegatorUnstaked struct {
	ValidatorId *big.Int
	User        common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDelegatorUnstaked is a free log retrieval operation binding the contract event 0x770c7c7d8e20347e5080e2ac70e8519793bedaff621f044396fd8d6d052c4aa8.
//
// Solidity: event DelegatorUnstaked(uint256 indexed validatorId, address indexed user, uint256 amount)
func (_Staking *StakingFilterer) FilterDelegatorUnstaked(opts *bind.FilterOpts, validatorId []*big.Int, user []common.Address) (*StakingDelegatorUnstakedIterator, error) {

	var validatorIdRule []interface{}
	for _, validatorIdItem := range validatorId {
		validatorIdRule = append(validatorIdRule, validatorIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "DelegatorUnstaked", validatorIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return &StakingDelegatorUnstakedIterator{contract: _Staking.contract, event: "DelegatorUnstaked", logs: logs, sub: sub}, nil
}

// WatchDelegatorUnstaked is a free log subscription operation binding the contract event 0x770c7c7d8e20347e5080e2ac70e8519793bedaff621f044396fd8d6d052c4aa8.
//
// Solidity: event DelegatorUnstaked(uint256 indexed validatorId, address indexed user, uint256 amount)
func (_Staking *StakingFilterer) WatchDelegatorUnstaked(opts *bind.WatchOpts, sink chan<- *StakingDelegatorUnstaked, validatorId []*big.Int, user []common.Address) (event.Subscription, error) {

	var validatorIdRule []interface{}
	for _, validatorIdItem := range validatorId {
		validatorIdRule = append(validatorIdRule, validatorIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "DelegatorUnstaked", validatorIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingDelegatorUnstaked)
				if err := _Staking.contract.UnpackLog(event, "DelegatorUnstaked", log); err != nil {
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

// ParseDelegatorUnstaked is a log parse operation binding the contract event 0x770c7c7d8e20347e5080e2ac70e8519793bedaff621f044396fd8d6d052c4aa8.
//
// Solidity: event DelegatorUnstaked(uint256 indexed validatorId, address indexed user, uint256 amount)
func (_Staking *StakingFilterer) ParseDelegatorUnstaked(log types.Log) (*StakingDelegatorUnstaked, error) {
	event := new(StakingDelegatorUnstaked)
	if err := _Staking.contract.UnpackLog(event, "DelegatorUnstaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingDynastyValueChangeIterator is returned from FilterDynastyValueChange and is used to iterate over the raw logs and unpacked data for DynastyValueChange events raised by the Staking contract.
type StakingDynastyValueChangeIterator struct {
	Event *StakingDynastyValueChange // Event containing the contract specifics and raw log

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
func (it *StakingDynastyValueChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingDynastyValueChange)
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
		it.Event = new(StakingDynastyValueChange)
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
func (it *StakingDynastyValueChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingDynastyValueChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingDynastyValueChange represents a DynastyValueChange event raised by the Staking contract.
type StakingDynastyValueChange struct {
	NewDynasty *big.Int
	OldDynasty *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDynastyValueChange is a free log retrieval operation binding the contract event 0x9444bfcfa6aed72a15da73de1220dcc07d7864119c44abfec0037bbcacefda98.
//
// Solidity: event DynastyValueChange(uint256 newDynasty, uint256 oldDynasty)
func (_Staking *StakingFilterer) FilterDynastyValueChange(opts *bind.FilterOpts) (*StakingDynastyValueChangeIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "DynastyValueChange")
	if err != nil {
		return nil, err
	}
	return &StakingDynastyValueChangeIterator{contract: _Staking.contract, event: "DynastyValueChange", logs: logs, sub: sub}, nil
}

// WatchDynastyValueChange is a free log subscription operation binding the contract event 0x9444bfcfa6aed72a15da73de1220dcc07d7864119c44abfec0037bbcacefda98.
//
// Solidity: event DynastyValueChange(uint256 newDynasty, uint256 oldDynasty)
func (_Staking *StakingFilterer) WatchDynastyValueChange(opts *bind.WatchOpts, sink chan<- *StakingDynastyValueChange) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "DynastyValueChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingDynastyValueChange)
				if err := _Staking.contract.UnpackLog(event, "DynastyValueChange", log); err != nil {
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

// ParseDynastyValueChange is a log parse operation binding the contract event 0x9444bfcfa6aed72a15da73de1220dcc07d7864119c44abfec0037bbcacefda98.
//
// Solidity: event DynastyValueChange(uint256 newDynasty, uint256 oldDynasty)
func (_Staking *StakingFilterer) ParseDynastyValueChange(log types.Log) (*StakingDynastyValueChange, error) {
	event := new(StakingDynastyValueChange)
	if err := _Staking.contract.UnpackLog(event, "DynastyValueChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingJailedIterator is returned from FilterJailed and is used to iterate over the raw logs and unpacked data for Jailed events raised by the Staking contract.
type StakingJailedIterator struct {
	Event *StakingJailed // Event containing the contract specifics and raw log

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
func (it *StakingJailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingJailed)
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
		it.Event = new(StakingJailed)
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
func (it *StakingJailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingJailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingJailed represents a Jailed event raised by the Staking contract.
type StakingJailed struct {
	ValidatorId *big.Int
	ExitEpoch   *big.Int
	Signer      common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterJailed is a free log retrieval operation binding the contract event 0xf6566d8fbe8f23227826ba3da2ecc1ec48698c5be051a829965e3358fd5b9658.
//
// Solidity: event Jailed(uint256 indexed validatorId, uint256 indexed exitEpoch, address indexed signer)
func (_Staking *StakingFilterer) FilterJailed(opts *bind.FilterOpts, validatorId []*big.Int, exitEpoch []*big.Int, signer []common.Address) (*StakingJailedIterator, error) {

	var validatorIdRule []interface{}
	for _, validatorIdItem := range validatorId {
		validatorIdRule = append(validatorIdRule, validatorIdItem)
	}
	var exitEpochRule []interface{}
	for _, exitEpochItem := range exitEpoch {
		exitEpochRule = append(exitEpochRule, exitEpochItem)
	}
	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Jailed", validatorIdRule, exitEpochRule, signerRule)
	if err != nil {
		return nil, err
	}
	return &StakingJailedIterator{contract: _Staking.contract, event: "Jailed", logs: logs, sub: sub}, nil
}

// WatchJailed is a free log subscription operation binding the contract event 0xf6566d8fbe8f23227826ba3da2ecc1ec48698c5be051a829965e3358fd5b9658.
//
// Solidity: event Jailed(uint256 indexed validatorId, uint256 indexed exitEpoch, address indexed signer)
func (_Staking *StakingFilterer) WatchJailed(opts *bind.WatchOpts, sink chan<- *StakingJailed, validatorId []*big.Int, exitEpoch []*big.Int, signer []common.Address) (event.Subscription, error) {

	var validatorIdRule []interface{}
	for _, validatorIdItem := range validatorId {
		validatorIdRule = append(validatorIdRule, validatorIdItem)
	}
	var exitEpochRule []interface{}
	for _, exitEpochItem := range exitEpoch {
		exitEpochRule = append(exitEpochRule, exitEpochItem)
	}
	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Jailed", validatorIdRule, exitEpochRule, signerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingJailed)
				if err := _Staking.contract.UnpackLog(event, "Jailed", log); err != nil {
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

// ParseJailed is a log parse operation binding the contract event 0xf6566d8fbe8f23227826ba3da2ecc1ec48698c5be051a829965e3358fd5b9658.
//
// Solidity: event Jailed(uint256 indexed validatorId, uint256 indexed exitEpoch, address indexed signer)
func (_Staking *StakingFilterer) ParseJailed(log types.Log) (*StakingJailed, error) {
	event := new(StakingJailed)
	if err := _Staking.contract.UnpackLog(event, "Jailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Staking contract.
type StakingOwnershipTransferredIterator struct {
	Event *StakingOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StakingOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingOwnershipTransferred)
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
		it.Event = new(StakingOwnershipTransferred)
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
func (it *StakingOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingOwnershipTransferred represents a OwnershipTransferred event raised by the Staking contract.
type StakingOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Staking *StakingFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StakingOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StakingOwnershipTransferredIterator{contract: _Staking.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Staking *StakingFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StakingOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingOwnershipTransferred)
				if err := _Staking.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Staking *StakingFilterer) ParseOwnershipTransferred(log types.Log) (*StakingOwnershipTransferred, error) {
	event := new(StakingOwnershipTransferred)
	if err := _Staking.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingProposerBonusChangeIterator is returned from FilterProposerBonusChange and is used to iterate over the raw logs and unpacked data for ProposerBonusChange events raised by the Staking contract.
type StakingProposerBonusChangeIterator struct {
	Event *StakingProposerBonusChange // Event containing the contract specifics and raw log

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
func (it *StakingProposerBonusChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingProposerBonusChange)
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
		it.Event = new(StakingProposerBonusChange)
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
func (it *StakingProposerBonusChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingProposerBonusChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingProposerBonusChange represents a ProposerBonusChange event raised by the Staking contract.
type StakingProposerBonusChange struct {
	NewProposerBonus *big.Int
	OldProposerBonus *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterProposerBonusChange is a free log retrieval operation binding the contract event 0x4a501a9c4d5cce5c32415945bbc8973764f31b844e3e8fd4c15f51f315ac8792.
//
// Solidity: event ProposerBonusChange(uint256 newProposerBonus, uint256 oldProposerBonus)
func (_Staking *StakingFilterer) FilterProposerBonusChange(opts *bind.FilterOpts) (*StakingProposerBonusChangeIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "ProposerBonusChange")
	if err != nil {
		return nil, err
	}
	return &StakingProposerBonusChangeIterator{contract: _Staking.contract, event: "ProposerBonusChange", logs: logs, sub: sub}, nil
}

// WatchProposerBonusChange is a free log subscription operation binding the contract event 0x4a501a9c4d5cce5c32415945bbc8973764f31b844e3e8fd4c15f51f315ac8792.
//
// Solidity: event ProposerBonusChange(uint256 newProposerBonus, uint256 oldProposerBonus)
func (_Staking *StakingFilterer) WatchProposerBonusChange(opts *bind.WatchOpts, sink chan<- *StakingProposerBonusChange) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "ProposerBonusChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingProposerBonusChange)
				if err := _Staking.contract.UnpackLog(event, "ProposerBonusChange", log); err != nil {
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

// ParseProposerBonusChange is a log parse operation binding the contract event 0x4a501a9c4d5cce5c32415945bbc8973764f31b844e3e8fd4c15f51f315ac8792.
//
// Solidity: event ProposerBonusChange(uint256 newProposerBonus, uint256 oldProposerBonus)
func (_Staking *StakingFilterer) ParseProposerBonusChange(log types.Log) (*StakingProposerBonusChange, error) {
	event := new(StakingProposerBonusChange)
	if err := _Staking.contract.UnpackLog(event, "ProposerBonusChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingRestakedIterator is returned from FilterRestaked and is used to iterate over the raw logs and unpacked data for Restaked events raised by the Staking contract.
type StakingRestakedIterator struct {
	Event *StakingRestaked // Event containing the contract specifics and raw log

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
func (it *StakingRestakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingRestaked)
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
		it.Event = new(StakingRestaked)
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
func (it *StakingRestakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingRestakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingRestaked represents a Restaked event raised by the Staking contract.
type StakingRestaked struct {
	ValidatorId *big.Int
	Amount      *big.Int
	Total       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRestaked is a free log retrieval operation binding the contract event 0x09b24121f82c610c13909ec63bd0843468819a45f6eda5838c3a80568c2046a8.
//
// Solidity: event Restaked(uint256 indexed validatorId, uint256 amount, uint256 total)
func (_Staking *StakingFilterer) FilterRestaked(opts *bind.FilterOpts, validatorId []*big.Int) (*StakingRestakedIterator, error) {

	var validatorIdRule []interface{}
	for _, validatorIdItem := range validatorId {
		validatorIdRule = append(validatorIdRule, validatorIdItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Restaked", validatorIdRule)
	if err != nil {
		return nil, err
	}
	return &StakingRestakedIterator{contract: _Staking.contract, event: "Restaked", logs: logs, sub: sub}, nil
}

// WatchRestaked is a free log subscription operation binding the contract event 0x09b24121f82c610c13909ec63bd0843468819a45f6eda5838c3a80568c2046a8.
//
// Solidity: event Restaked(uint256 indexed validatorId, uint256 amount, uint256 total)
func (_Staking *StakingFilterer) WatchRestaked(opts *bind.WatchOpts, sink chan<- *StakingRestaked, validatorId []*big.Int) (event.Subscription, error) {

	var validatorIdRule []interface{}
	for _, validatorIdItem := range validatorId {
		validatorIdRule = append(validatorIdRule, validatorIdItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Restaked", validatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingRestaked)
				if err := _Staking.contract.UnpackLog(event, "Restaked", log); err != nil {
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

// ParseRestaked is a log parse operation binding the contract event 0x09b24121f82c610c13909ec63bd0843468819a45f6eda5838c3a80568c2046a8.
//
// Solidity: event Restaked(uint256 indexed validatorId, uint256 amount, uint256 total)
func (_Staking *StakingFilterer) ParseRestaked(log types.Log) (*StakingRestaked, error) {
	event := new(StakingRestaked)
	if err := _Staking.contract.UnpackLog(event, "Restaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingRewardUpdateIterator is returned from FilterRewardUpdate and is used to iterate over the raw logs and unpacked data for RewardUpdate events raised by the Staking contract.
type StakingRewardUpdateIterator struct {
	Event *StakingRewardUpdate // Event containing the contract specifics and raw log

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
func (it *StakingRewardUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingRewardUpdate)
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
		it.Event = new(StakingRewardUpdate)
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
func (it *StakingRewardUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingRewardUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingRewardUpdate represents a RewardUpdate event raised by the Staking contract.
type StakingRewardUpdate struct {
	NewReward *big.Int
	OldReward *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRewardUpdate is a free log retrieval operation binding the contract event 0xf67f33e8589d3ea0356303c0f9a8e764873692159f777ff79e4fc523d389dfcd.
//
// Solidity: event RewardUpdate(uint256 newReward, uint256 oldReward)
func (_Staking *StakingFilterer) FilterRewardUpdate(opts *bind.FilterOpts) (*StakingRewardUpdateIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "RewardUpdate")
	if err != nil {
		return nil, err
	}
	return &StakingRewardUpdateIterator{contract: _Staking.contract, event: "RewardUpdate", logs: logs, sub: sub}, nil
}

// WatchRewardUpdate is a free log subscription operation binding the contract event 0xf67f33e8589d3ea0356303c0f9a8e764873692159f777ff79e4fc523d389dfcd.
//
// Solidity: event RewardUpdate(uint256 newReward, uint256 oldReward)
func (_Staking *StakingFilterer) WatchRewardUpdate(opts *bind.WatchOpts, sink chan<- *StakingRewardUpdate) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "RewardUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingRewardUpdate)
				if err := _Staking.contract.UnpackLog(event, "RewardUpdate", log); err != nil {
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

// ParseRewardUpdate is a log parse operation binding the contract event 0xf67f33e8589d3ea0356303c0f9a8e764873692159f777ff79e4fc523d389dfcd.
//
// Solidity: event RewardUpdate(uint256 newReward, uint256 oldReward)
func (_Staking *StakingFilterer) ParseRewardUpdate(log types.Log) (*StakingRewardUpdate, error) {
	event := new(StakingRewardUpdate)
	if err := _Staking.contract.UnpackLog(event, "RewardUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingShareBurnedIterator is returned from FilterShareBurned and is used to iterate over the raw logs and unpacked data for ShareBurned events raised by the Staking contract.
type StakingShareBurnedIterator struct {
	Event *StakingShareBurned // Event containing the contract specifics and raw log

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
func (it *StakingShareBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingShareBurned)
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
		it.Event = new(StakingShareBurned)
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
func (it *StakingShareBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingShareBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingShareBurned represents a ShareBurned event raised by the Staking contract.
type StakingShareBurned struct {
	ValidatorId *big.Int
	User        common.Address
	Amount      *big.Int
	Tokens      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterShareBurned is a free log retrieval operation binding the contract event 0x7e86625aa6e668407f095af342e0cc237809c4c5086b4d665a0067de122980a9.
//
// Solidity: event ShareBurned(uint256 indexed validatorId, address indexed user, uint256 indexed amount, uint256 tokens)
func (_Staking *StakingFilterer) FilterShareBurned(opts *bind.FilterOpts, validatorId []*big.Int, user []common.Address, amount []*big.Int) (*StakingShareBurnedIterator, error) {

	var validatorIdRule []interface{}
	for _, validatorIdItem := range validatorId {
		validatorIdRule = append(validatorIdRule, validatorIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "ShareBurned", validatorIdRule, userRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &StakingShareBurnedIterator{contract: _Staking.contract, event: "ShareBurned", logs: logs, sub: sub}, nil
}

// WatchShareBurned is a free log subscription operation binding the contract event 0x7e86625aa6e668407f095af342e0cc237809c4c5086b4d665a0067de122980a9.
//
// Solidity: event ShareBurned(uint256 indexed validatorId, address indexed user, uint256 indexed amount, uint256 tokens)
func (_Staking *StakingFilterer) WatchShareBurned(opts *bind.WatchOpts, sink chan<- *StakingShareBurned, validatorId []*big.Int, user []common.Address, amount []*big.Int) (event.Subscription, error) {

	var validatorIdRule []interface{}
	for _, validatorIdItem := range validatorId {
		validatorIdRule = append(validatorIdRule, validatorIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "ShareBurned", validatorIdRule, userRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingShareBurned)
				if err := _Staking.contract.UnpackLog(event, "ShareBurned", log); err != nil {
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

// ParseShareBurned is a log parse operation binding the contract event 0x7e86625aa6e668407f095af342e0cc237809c4c5086b4d665a0067de122980a9.
//
// Solidity: event ShareBurned(uint256 indexed validatorId, address indexed user, uint256 indexed amount, uint256 tokens)
func (_Staking *StakingFilterer) ParseShareBurned(log types.Log) (*StakingShareBurned, error) {
	event := new(StakingShareBurned)
	if err := _Staking.contract.UnpackLog(event, "ShareBurned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingShareMintedIterator is returned from FilterShareMinted and is used to iterate over the raw logs and unpacked data for ShareMinted events raised by the Staking contract.
type StakingShareMintedIterator struct {
	Event *StakingShareMinted // Event containing the contract specifics and raw log

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
func (it *StakingShareMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingShareMinted)
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
		it.Event = new(StakingShareMinted)
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
func (it *StakingShareMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingShareMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingShareMinted represents a ShareMinted event raised by the Staking contract.
type StakingShareMinted struct {
	ValidatorId *big.Int
	User        common.Address
	Amount      *big.Int
	Tokens      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterShareMinted is a free log retrieval operation binding the contract event 0xc9afff0972d33d68c8d330fe0ebd0e9f54491ad8c59ae17330a9206f280f0865.
//
// Solidity: event ShareMinted(uint256 indexed validatorId, address indexed user, uint256 indexed amount, uint256 tokens)
func (_Staking *StakingFilterer) FilterShareMinted(opts *bind.FilterOpts, validatorId []*big.Int, user []common.Address, amount []*big.Int) (*StakingShareMintedIterator, error) {

	var validatorIdRule []interface{}
	for _, validatorIdItem := range validatorId {
		validatorIdRule = append(validatorIdRule, validatorIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "ShareMinted", validatorIdRule, userRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &StakingShareMintedIterator{contract: _Staking.contract, event: "ShareMinted", logs: logs, sub: sub}, nil
}

// WatchShareMinted is a free log subscription operation binding the contract event 0xc9afff0972d33d68c8d330fe0ebd0e9f54491ad8c59ae17330a9206f280f0865.
//
// Solidity: event ShareMinted(uint256 indexed validatorId, address indexed user, uint256 indexed amount, uint256 tokens)
func (_Staking *StakingFilterer) WatchShareMinted(opts *bind.WatchOpts, sink chan<- *StakingShareMinted, validatorId []*big.Int, user []common.Address, amount []*big.Int) (event.Subscription, error) {

	var validatorIdRule []interface{}
	for _, validatorIdItem := range validatorId {
		validatorIdRule = append(validatorIdRule, validatorIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "ShareMinted", validatorIdRule, userRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingShareMinted)
				if err := _Staking.contract.UnpackLog(event, "ShareMinted", log); err != nil {
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

// ParseShareMinted is a log parse operation binding the contract event 0xc9afff0972d33d68c8d330fe0ebd0e9f54491ad8c59ae17330a9206f280f0865.
//
// Solidity: event ShareMinted(uint256 indexed validatorId, address indexed user, uint256 indexed amount, uint256 tokens)
func (_Staking *StakingFilterer) ParseShareMinted(log types.Log) (*StakingShareMinted, error) {
	event := new(StakingShareMinted)
	if err := _Staking.contract.UnpackLog(event, "ShareMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingSignerChangeIterator is returned from FilterSignerChange and is used to iterate over the raw logs and unpacked data for SignerChange events raised by the Staking contract.
type StakingSignerChangeIterator struct {
	Event *StakingSignerChange // Event containing the contract specifics and raw log

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
func (it *StakingSignerChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingSignerChange)
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
		it.Event = new(StakingSignerChange)
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
func (it *StakingSignerChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingSignerChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingSignerChange represents a SignerChange event raised by the Staking contract.
type StakingSignerChange struct {
	ValidatorId  *big.Int
	Nonce        *big.Int
	OldSigner    common.Address
	NewSigner    common.Address
	SignerPubkey []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSignerChange is a free log retrieval operation binding the contract event 0x086044c0612a8c965d4cccd907f0d588e40ad68438bd4c1274cac60f4c3a9d1f.
//
// Solidity: event SignerChange(uint256 indexed validatorId, uint256 nonce, address indexed oldSigner, address indexed newSigner, bytes signerPubkey)
func (_Staking *StakingFilterer) FilterSignerChange(opts *bind.FilterOpts, validatorId []*big.Int, oldSigner []common.Address, newSigner []common.Address) (*StakingSignerChangeIterator, error) {

	var validatorIdRule []interface{}
	for _, validatorIdItem := range validatorId {
		validatorIdRule = append(validatorIdRule, validatorIdItem)
	}

	var oldSignerRule []interface{}
	for _, oldSignerItem := range oldSigner {
		oldSignerRule = append(oldSignerRule, oldSignerItem)
	}
	var newSignerRule []interface{}
	for _, newSignerItem := range newSigner {
		newSignerRule = append(newSignerRule, newSignerItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "SignerChange", validatorIdRule, oldSignerRule, newSignerRule)
	if err != nil {
		return nil, err
	}
	return &StakingSignerChangeIterator{contract: _Staking.contract, event: "SignerChange", logs: logs, sub: sub}, nil
}

// WatchSignerChange is a free log subscription operation binding the contract event 0x086044c0612a8c965d4cccd907f0d588e40ad68438bd4c1274cac60f4c3a9d1f.
//
// Solidity: event SignerChange(uint256 indexed validatorId, uint256 nonce, address indexed oldSigner, address indexed newSigner, bytes signerPubkey)
func (_Staking *StakingFilterer) WatchSignerChange(opts *bind.WatchOpts, sink chan<- *StakingSignerChange, validatorId []*big.Int, oldSigner []common.Address, newSigner []common.Address) (event.Subscription, error) {

	var validatorIdRule []interface{}
	for _, validatorIdItem := range validatorId {
		validatorIdRule = append(validatorIdRule, validatorIdItem)
	}

	var oldSignerRule []interface{}
	for _, oldSignerItem := range oldSigner {
		oldSignerRule = append(oldSignerRule, oldSignerItem)
	}
	var newSignerRule []interface{}
	for _, newSignerItem := range newSigner {
		newSignerRule = append(newSignerRule, newSignerItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "SignerChange", validatorIdRule, oldSignerRule, newSignerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingSignerChange)
				if err := _Staking.contract.UnpackLog(event, "SignerChange", log); err != nil {
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

// ParseSignerChange is a log parse operation binding the contract event 0x086044c0612a8c965d4cccd907f0d588e40ad68438bd4c1274cac60f4c3a9d1f.
//
// Solidity: event SignerChange(uint256 indexed validatorId, uint256 nonce, address indexed oldSigner, address indexed newSigner, bytes signerPubkey)
func (_Staking *StakingFilterer) ParseSignerChange(log types.Log) (*StakingSignerChange, error) {
	event := new(StakingSignerChange)
	if err := _Staking.contract.UnpackLog(event, "SignerChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingSlashedIterator is returned from FilterSlashed and is used to iterate over the raw logs and unpacked data for Slashed events raised by the Staking contract.
type StakingSlashedIterator struct {
	Event *StakingSlashed // Event containing the contract specifics and raw log

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
func (it *StakingSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingSlashed)
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
		it.Event = new(StakingSlashed)
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
func (it *StakingSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingSlashed represents a Slashed event raised by the Staking contract.
type StakingSlashed struct {
	Nonce  *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSlashed is a free log retrieval operation binding the contract event 0x4f5f38ee30b01a960b4dfdcd520a3ca59c1a664a32dcfe5418ca79b0de6b7236.
//
// Solidity: event Slashed(uint256 indexed nonce, uint256 indexed amount)
func (_Staking *StakingFilterer) FilterSlashed(opts *bind.FilterOpts, nonce []*big.Int, amount []*big.Int) (*StakingSlashedIterator, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Slashed", nonceRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &StakingSlashedIterator{contract: _Staking.contract, event: "Slashed", logs: logs, sub: sub}, nil
}

// WatchSlashed is a free log subscription operation binding the contract event 0x4f5f38ee30b01a960b4dfdcd520a3ca59c1a664a32dcfe5418ca79b0de6b7236.
//
// Solidity: event Slashed(uint256 indexed nonce, uint256 indexed amount)
func (_Staking *StakingFilterer) WatchSlashed(opts *bind.WatchOpts, sink chan<- *StakingSlashed, nonce []*big.Int, amount []*big.Int) (event.Subscription, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Slashed", nonceRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingSlashed)
				if err := _Staking.contract.UnpackLog(event, "Slashed", log); err != nil {
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

// ParseSlashed is a log parse operation binding the contract event 0x4f5f38ee30b01a960b4dfdcd520a3ca59c1a664a32dcfe5418ca79b0de6b7236.
//
// Solidity: event Slashed(uint256 indexed nonce, uint256 indexed amount)
func (_Staking *StakingFilterer) ParseSlashed(log types.Log) (*StakingSlashed, error) {
	event := new(StakingSlashed)
	if err := _Staking.contract.UnpackLog(event, "Slashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingStakeUpdateIterator is returned from FilterStakeUpdate and is used to iterate over the raw logs and unpacked data for StakeUpdate events raised by the Staking contract.
type StakingStakeUpdateIterator struct {
	Event *StakingStakeUpdate // Event containing the contract specifics and raw log

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
func (it *StakingStakeUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingStakeUpdate)
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
		it.Event = new(StakingStakeUpdate)
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
func (it *StakingStakeUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingStakeUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingStakeUpdate represents a StakeUpdate event raised by the Staking contract.
type StakingStakeUpdate struct {
	ValidatorId *big.Int
	Nonce       *big.Int
	NewAmount   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterStakeUpdate is a free log retrieval operation binding the contract event 0x35af9eea1f0e7b300b0a14fae90139a072470e44daa3f14b5069bebbc1265bda.
//
// Solidity: event StakeUpdate(uint256 indexed validatorId, uint256 indexed nonce, uint256 indexed newAmount)
func (_Staking *StakingFilterer) FilterStakeUpdate(opts *bind.FilterOpts, validatorId []*big.Int, nonce []*big.Int, newAmount []*big.Int) (*StakingStakeUpdateIterator, error) {

	var validatorIdRule []interface{}
	for _, validatorIdItem := range validatorId {
		validatorIdRule = append(validatorIdRule, validatorIdItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var newAmountRule []interface{}
	for _, newAmountItem := range newAmount {
		newAmountRule = append(newAmountRule, newAmountItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "StakeUpdate", validatorIdRule, nonceRule, newAmountRule)
	if err != nil {
		return nil, err
	}
	return &StakingStakeUpdateIterator{contract: _Staking.contract, event: "StakeUpdate", logs: logs, sub: sub}, nil
}

// WatchStakeUpdate is a free log subscription operation binding the contract event 0x35af9eea1f0e7b300b0a14fae90139a072470e44daa3f14b5069bebbc1265bda.
//
// Solidity: event StakeUpdate(uint256 indexed validatorId, uint256 indexed nonce, uint256 indexed newAmount)
func (_Staking *StakingFilterer) WatchStakeUpdate(opts *bind.WatchOpts, sink chan<- *StakingStakeUpdate, validatorId []*big.Int, nonce []*big.Int, newAmount []*big.Int) (event.Subscription, error) {

	var validatorIdRule []interface{}
	for _, validatorIdItem := range validatorId {
		validatorIdRule = append(validatorIdRule, validatorIdItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var newAmountRule []interface{}
	for _, newAmountItem := range newAmount {
		newAmountRule = append(newAmountRule, newAmountItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "StakeUpdate", validatorIdRule, nonceRule, newAmountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingStakeUpdate)
				if err := _Staking.contract.UnpackLog(event, "StakeUpdate", log); err != nil {
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

// ParseStakeUpdate is a log parse operation binding the contract event 0x35af9eea1f0e7b300b0a14fae90139a072470e44daa3f14b5069bebbc1265bda.
//
// Solidity: event StakeUpdate(uint256 indexed validatorId, uint256 indexed nonce, uint256 indexed newAmount)
func (_Staking *StakingFilterer) ParseStakeUpdate(log types.Log) (*StakingStakeUpdate, error) {
	event := new(StakingStakeUpdate)
	if err := _Staking.contract.UnpackLog(event, "StakeUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the Staking contract.
type StakingStakedIterator struct {
	Event *StakingStaked // Event containing the contract specifics and raw log

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
func (it *StakingStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingStaked)
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
		it.Event = new(StakingStaked)
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
func (it *StakingStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingStaked represents a Staked event raised by the Staking contract.
type StakingStaked struct {
	Signer          common.Address
	ValidatorId     *big.Int
	Nonce           *big.Int
	ActivationEpoch *big.Int
	Amount          *big.Int
	Total           *big.Int
	SignerPubkey    []byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x68c13e4125b983d7e2d6114246f443e567ec6c4ee5b4d4a7ef6100b1402bfd84.
//
// Solidity: event Staked(address indexed signer, uint256 indexed validatorId, uint256 nonce, uint256 indexed activationEpoch, uint256 amount, uint256 total, bytes signerPubkey)
func (_Staking *StakingFilterer) FilterStaked(opts *bind.FilterOpts, signer []common.Address, validatorId []*big.Int, activationEpoch []*big.Int) (*StakingStakedIterator, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}
	var validatorIdRule []interface{}
	for _, validatorIdItem := range validatorId {
		validatorIdRule = append(validatorIdRule, validatorIdItem)
	}

	var activationEpochRule []interface{}
	for _, activationEpochItem := range activationEpoch {
		activationEpochRule = append(activationEpochRule, activationEpochItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Staked", signerRule, validatorIdRule, activationEpochRule)
	if err != nil {
		return nil, err
	}
	return &StakingStakedIterator{contract: _Staking.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x68c13e4125b983d7e2d6114246f443e567ec6c4ee5b4d4a7ef6100b1402bfd84.
//
// Solidity: event Staked(address indexed signer, uint256 indexed validatorId, uint256 nonce, uint256 indexed activationEpoch, uint256 amount, uint256 total, bytes signerPubkey)
func (_Staking *StakingFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *StakingStaked, signer []common.Address, validatorId []*big.Int, activationEpoch []*big.Int) (event.Subscription, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}
	var validatorIdRule []interface{}
	for _, validatorIdItem := range validatorId {
		validatorIdRule = append(validatorIdRule, validatorIdItem)
	}

	var activationEpochRule []interface{}
	for _, activationEpochItem := range activationEpoch {
		activationEpochRule = append(activationEpochRule, activationEpochItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Staked", signerRule, validatorIdRule, activationEpochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingStaked)
				if err := _Staking.contract.UnpackLog(event, "Staked", log); err != nil {
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

// ParseStaked is a log parse operation binding the contract event 0x68c13e4125b983d7e2d6114246f443e567ec6c4ee5b4d4a7ef6100b1402bfd84.
//
// Solidity: event Staked(address indexed signer, uint256 indexed validatorId, uint256 nonce, uint256 indexed activationEpoch, uint256 amount, uint256 total, bytes signerPubkey)
func (_Staking *StakingFilterer) ParseStaked(log types.Log) (*StakingStaked, error) {
	event := new(StakingStaked)
	if err := _Staking.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingStartAuctionIterator is returned from FilterStartAuction and is used to iterate over the raw logs and unpacked data for StartAuction events raised by the Staking contract.
type StakingStartAuctionIterator struct {
	Event *StakingStartAuction // Event containing the contract specifics and raw log

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
func (it *StakingStartAuctionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingStartAuction)
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
		it.Event = new(StakingStartAuction)
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
func (it *StakingStartAuctionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingStartAuctionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingStartAuction represents a StartAuction event raised by the Staking contract.
type StakingStartAuction struct {
	ValidatorId   *big.Int
	Amount        *big.Int
	AuctionAmount *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterStartAuction is a free log retrieval operation binding the contract event 0x683d0f47c7fa11331f4e9563b3f5a7fdc3d3c5b75c600357a91d991f5a13a437.
//
// Solidity: event StartAuction(uint256 indexed validatorId, uint256 indexed amount, uint256 indexed auctionAmount)
func (_Staking *StakingFilterer) FilterStartAuction(opts *bind.FilterOpts, validatorId []*big.Int, amount []*big.Int, auctionAmount []*big.Int) (*StakingStartAuctionIterator, error) {

	var validatorIdRule []interface{}
	for _, validatorIdItem := range validatorId {
		validatorIdRule = append(validatorIdRule, validatorIdItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}
	var auctionAmountRule []interface{}
	for _, auctionAmountItem := range auctionAmount {
		auctionAmountRule = append(auctionAmountRule, auctionAmountItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "StartAuction", validatorIdRule, amountRule, auctionAmountRule)
	if err != nil {
		return nil, err
	}
	return &StakingStartAuctionIterator{contract: _Staking.contract, event: "StartAuction", logs: logs, sub: sub}, nil
}

// WatchStartAuction is a free log subscription operation binding the contract event 0x683d0f47c7fa11331f4e9563b3f5a7fdc3d3c5b75c600357a91d991f5a13a437.
//
// Solidity: event StartAuction(uint256 indexed validatorId, uint256 indexed amount, uint256 indexed auctionAmount)
func (_Staking *StakingFilterer) WatchStartAuction(opts *bind.WatchOpts, sink chan<- *StakingStartAuction, validatorId []*big.Int, amount []*big.Int, auctionAmount []*big.Int) (event.Subscription, error) {

	var validatorIdRule []interface{}
	for _, validatorIdItem := range validatorId {
		validatorIdRule = append(validatorIdRule, validatorIdItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}
	var auctionAmountRule []interface{}
	for _, auctionAmountItem := range auctionAmount {
		auctionAmountRule = append(auctionAmountRule, auctionAmountItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "StartAuction", validatorIdRule, amountRule, auctionAmountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingStartAuction)
				if err := _Staking.contract.UnpackLog(event, "StartAuction", log); err != nil {
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

// ParseStartAuction is a log parse operation binding the contract event 0x683d0f47c7fa11331f4e9563b3f5a7fdc3d3c5b75c600357a91d991f5a13a437.
//
// Solidity: event StartAuction(uint256 indexed validatorId, uint256 indexed amount, uint256 indexed auctionAmount)
func (_Staking *StakingFilterer) ParseStartAuction(log types.Log) (*StakingStartAuction, error) {
	event := new(StakingStartAuction)
	if err := _Staking.contract.UnpackLog(event, "StartAuction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingThresholdChangeIterator is returned from FilterThresholdChange and is used to iterate over the raw logs and unpacked data for ThresholdChange events raised by the Staking contract.
type StakingThresholdChangeIterator struct {
	Event *StakingThresholdChange // Event containing the contract specifics and raw log

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
func (it *StakingThresholdChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingThresholdChange)
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
		it.Event = new(StakingThresholdChange)
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
func (it *StakingThresholdChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingThresholdChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingThresholdChange represents a ThresholdChange event raised by the Staking contract.
type StakingThresholdChange struct {
	NewThreshold *big.Int
	OldThreshold *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterThresholdChange is a free log retrieval operation binding the contract event 0x5d16a900896e1160c2033bc940e6b072d3dc3b6a996fefb9b3b9b9678841824c.
//
// Solidity: event ThresholdChange(uint256 newThreshold, uint256 oldThreshold)
func (_Staking *StakingFilterer) FilterThresholdChange(opts *bind.FilterOpts) (*StakingThresholdChangeIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "ThresholdChange")
	if err != nil {
		return nil, err
	}
	return &StakingThresholdChangeIterator{contract: _Staking.contract, event: "ThresholdChange", logs: logs, sub: sub}, nil
}

// WatchThresholdChange is a free log subscription operation binding the contract event 0x5d16a900896e1160c2033bc940e6b072d3dc3b6a996fefb9b3b9b9678841824c.
//
// Solidity: event ThresholdChange(uint256 newThreshold, uint256 oldThreshold)
func (_Staking *StakingFilterer) WatchThresholdChange(opts *bind.WatchOpts, sink chan<- *StakingThresholdChange) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "ThresholdChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingThresholdChange)
				if err := _Staking.contract.UnpackLog(event, "ThresholdChange", log); err != nil {
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

// ParseThresholdChange is a log parse operation binding the contract event 0x5d16a900896e1160c2033bc940e6b072d3dc3b6a996fefb9b3b9b9678841824c.
//
// Solidity: event ThresholdChange(uint256 newThreshold, uint256 oldThreshold)
func (_Staking *StakingFilterer) ParseThresholdChange(log types.Log) (*StakingThresholdChange, error) {
	event := new(StakingThresholdChange)
	if err := _Staking.contract.UnpackLog(event, "ThresholdChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingTopUpFeeIterator is returned from FilterTopUpFee and is used to iterate over the raw logs and unpacked data for TopUpFee events raised by the Staking contract.
type StakingTopUpFeeIterator struct {
	Event *StakingTopUpFee // Event containing the contract specifics and raw log

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
func (it *StakingTopUpFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingTopUpFee)
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
		it.Event = new(StakingTopUpFee)
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
func (it *StakingTopUpFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingTopUpFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingTopUpFee represents a TopUpFee event raised by the Staking contract.
type StakingTopUpFee struct {
	User common.Address
	Fee  *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterTopUpFee is a free log retrieval operation binding the contract event 0x2c3bb5458e3dd671c31974c4ca8e8ebc2cdd892ae8602374d9a6f789b00c6b94.
//
// Solidity: event TopUpFee(address indexed user, uint256 indexed fee)
func (_Staking *StakingFilterer) FilterTopUpFee(opts *bind.FilterOpts, user []common.Address, fee []*big.Int) (*StakingTopUpFeeIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "TopUpFee", userRule, feeRule)
	if err != nil {
		return nil, err
	}
	return &StakingTopUpFeeIterator{contract: _Staking.contract, event: "TopUpFee", logs: logs, sub: sub}, nil
}

// WatchTopUpFee is a free log subscription operation binding the contract event 0x2c3bb5458e3dd671c31974c4ca8e8ebc2cdd892ae8602374d9a6f789b00c6b94.
//
// Solidity: event TopUpFee(address indexed user, uint256 indexed fee)
func (_Staking *StakingFilterer) WatchTopUpFee(opts *bind.WatchOpts, sink chan<- *StakingTopUpFee, user []common.Address, fee []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "TopUpFee", userRule, feeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingTopUpFee)
				if err := _Staking.contract.UnpackLog(event, "TopUpFee", log); err != nil {
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

// ParseTopUpFee is a log parse operation binding the contract event 0x2c3bb5458e3dd671c31974c4ca8e8ebc2cdd892ae8602374d9a6f789b00c6b94.
//
// Solidity: event TopUpFee(address indexed user, uint256 indexed fee)
func (_Staking *StakingFilterer) ParseTopUpFee(log types.Log) (*StakingTopUpFee, error) {
	event := new(StakingTopUpFee)
	if err := _Staking.contract.UnpackLog(event, "TopUpFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingUnJailedIterator is returned from FilterUnJailed and is used to iterate over the raw logs and unpacked data for UnJailed events raised by the Staking contract.
type StakingUnJailedIterator struct {
	Event *StakingUnJailed // Event containing the contract specifics and raw log

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
func (it *StakingUnJailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingUnJailed)
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
		it.Event = new(StakingUnJailed)
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
func (it *StakingUnJailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingUnJailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingUnJailed represents a UnJailed event raised by the Staking contract.
type StakingUnJailed struct {
	ValidatorId *big.Int
	Signer      common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUnJailed is a free log retrieval operation binding the contract event 0xd3cb87a9c75a0d21336afc0f79f7e398f06748db5ce1815af01d315c7c135c0b.
//
// Solidity: event UnJailed(uint256 indexed validatorId, address indexed signer)
func (_Staking *StakingFilterer) FilterUnJailed(opts *bind.FilterOpts, validatorId []*big.Int, signer []common.Address) (*StakingUnJailedIterator, error) {

	var validatorIdRule []interface{}
	for _, validatorIdItem := range validatorId {
		validatorIdRule = append(validatorIdRule, validatorIdItem)
	}
	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "UnJailed", validatorIdRule, signerRule)
	if err != nil {
		return nil, err
	}
	return &StakingUnJailedIterator{contract: _Staking.contract, event: "UnJailed", logs: logs, sub: sub}, nil
}

// WatchUnJailed is a free log subscription operation binding the contract event 0xd3cb87a9c75a0d21336afc0f79f7e398f06748db5ce1815af01d315c7c135c0b.
//
// Solidity: event UnJailed(uint256 indexed validatorId, address indexed signer)
func (_Staking *StakingFilterer) WatchUnJailed(opts *bind.WatchOpts, sink chan<- *StakingUnJailed, validatorId []*big.Int, signer []common.Address) (event.Subscription, error) {

	var validatorIdRule []interface{}
	for _, validatorIdItem := range validatorId {
		validatorIdRule = append(validatorIdRule, validatorIdItem)
	}
	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "UnJailed", validatorIdRule, signerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingUnJailed)
				if err := _Staking.contract.UnpackLog(event, "UnJailed", log); err != nil {
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

// ParseUnJailed is a log parse operation binding the contract event 0xd3cb87a9c75a0d21336afc0f79f7e398f06748db5ce1815af01d315c7c135c0b.
//
// Solidity: event UnJailed(uint256 indexed validatorId, address indexed signer)
func (_Staking *StakingFilterer) ParseUnJailed(log types.Log) (*StakingUnJailed, error) {
	event := new(StakingUnJailed)
	if err := _Staking.contract.UnpackLog(event, "UnJailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingUnstakeInitIterator is returned from FilterUnstakeInit and is used to iterate over the raw logs and unpacked data for UnstakeInit events raised by the Staking contract.
type StakingUnstakeInitIterator struct {
	Event *StakingUnstakeInit // Event containing the contract specifics and raw log

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
func (it *StakingUnstakeInitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingUnstakeInit)
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
		it.Event = new(StakingUnstakeInit)
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
func (it *StakingUnstakeInitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingUnstakeInitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingUnstakeInit represents a UnstakeInit event raised by the Staking contract.
type StakingUnstakeInit struct {
	User              common.Address
	ValidatorId       *big.Int
	Nonce             *big.Int
	DeactivationEpoch *big.Int
	Amount            *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterUnstakeInit is a free log retrieval operation binding the contract event 0x69b288bb79cd5386c9fe0af060f650e823bcdfa96a44cdc07f862db060f57120.
//
// Solidity: event UnstakeInit(address indexed user, uint256 indexed validatorId, uint256 nonce, uint256 deactivationEpoch, uint256 indexed amount)
func (_Staking *StakingFilterer) FilterUnstakeInit(opts *bind.FilterOpts, user []common.Address, validatorId []*big.Int, amount []*big.Int) (*StakingUnstakeInitIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var validatorIdRule []interface{}
	for _, validatorIdItem := range validatorId {
		validatorIdRule = append(validatorIdRule, validatorIdItem)
	}

	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "UnstakeInit", userRule, validatorIdRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &StakingUnstakeInitIterator{contract: _Staking.contract, event: "UnstakeInit", logs: logs, sub: sub}, nil
}

// WatchUnstakeInit is a free log subscription operation binding the contract event 0x69b288bb79cd5386c9fe0af060f650e823bcdfa96a44cdc07f862db060f57120.
//
// Solidity: event UnstakeInit(address indexed user, uint256 indexed validatorId, uint256 nonce, uint256 deactivationEpoch, uint256 indexed amount)
func (_Staking *StakingFilterer) WatchUnstakeInit(opts *bind.WatchOpts, sink chan<- *StakingUnstakeInit, user []common.Address, validatorId []*big.Int, amount []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var validatorIdRule []interface{}
	for _, validatorIdItem := range validatorId {
		validatorIdRule = append(validatorIdRule, validatorIdItem)
	}

	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "UnstakeInit", userRule, validatorIdRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingUnstakeInit)
				if err := _Staking.contract.UnpackLog(event, "UnstakeInit", log); err != nil {
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

// ParseUnstakeInit is a log parse operation binding the contract event 0x69b288bb79cd5386c9fe0af060f650e823bcdfa96a44cdc07f862db060f57120.
//
// Solidity: event UnstakeInit(address indexed user, uint256 indexed validatorId, uint256 nonce, uint256 deactivationEpoch, uint256 indexed amount)
func (_Staking *StakingFilterer) ParseUnstakeInit(log types.Log) (*StakingUnstakeInit, error) {
	event := new(StakingUnstakeInit)
	if err := _Staking.contract.UnpackLog(event, "UnstakeInit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingUnstakedIterator is returned from FilterUnstaked and is used to iterate over the raw logs and unpacked data for Unstaked events raised by the Staking contract.
type StakingUnstakedIterator struct {
	Event *StakingUnstaked // Event containing the contract specifics and raw log

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
func (it *StakingUnstakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingUnstaked)
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
		it.Event = new(StakingUnstaked)
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
func (it *StakingUnstakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingUnstakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingUnstaked represents a Unstaked event raised by the Staking contract.
type StakingUnstaked struct {
	User        common.Address
	ValidatorId *big.Int
	Amount      *big.Int
	Total       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUnstaked is a free log retrieval operation binding the contract event 0x204fccf0d92ed8d48f204adb39b2e81e92bad0dedb93f5716ca9478cfb57de00.
//
// Solidity: event Unstaked(address indexed user, uint256 indexed validatorId, uint256 amount, uint256 total)
func (_Staking *StakingFilterer) FilterUnstaked(opts *bind.FilterOpts, user []common.Address, validatorId []*big.Int) (*StakingUnstakedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var validatorIdRule []interface{}
	for _, validatorIdItem := range validatorId {
		validatorIdRule = append(validatorIdRule, validatorIdItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Unstaked", userRule, validatorIdRule)
	if err != nil {
		return nil, err
	}
	return &StakingUnstakedIterator{contract: _Staking.contract, event: "Unstaked", logs: logs, sub: sub}, nil
}

// WatchUnstaked is a free log subscription operation binding the contract event 0x204fccf0d92ed8d48f204adb39b2e81e92bad0dedb93f5716ca9478cfb57de00.
//
// Solidity: event Unstaked(address indexed user, uint256 indexed validatorId, uint256 amount, uint256 total)
func (_Staking *StakingFilterer) WatchUnstaked(opts *bind.WatchOpts, sink chan<- *StakingUnstaked, user []common.Address, validatorId []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var validatorIdRule []interface{}
	for _, validatorIdItem := range validatorId {
		validatorIdRule = append(validatorIdRule, validatorIdItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Unstaked", userRule, validatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingUnstaked)
				if err := _Staking.contract.UnpackLog(event, "Unstaked", log); err != nil {
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

// ParseUnstaked is a log parse operation binding the contract event 0x204fccf0d92ed8d48f204adb39b2e81e92bad0dedb93f5716ca9478cfb57de00.
//
// Solidity: event Unstaked(address indexed user, uint256 indexed validatorId, uint256 amount, uint256 total)
func (_Staking *StakingFilterer) ParseUnstaked(log types.Log) (*StakingUnstaked, error) {
	event := new(StakingUnstaked)
	if err := _Staking.contract.UnpackLog(event, "Unstaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingUpdateCommissionRateIterator is returned from FilterUpdateCommissionRate and is used to iterate over the raw logs and unpacked data for UpdateCommissionRate events raised by the Staking contract.
type StakingUpdateCommissionRateIterator struct {
	Event *StakingUpdateCommissionRate // Event containing the contract specifics and raw log

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
func (it *StakingUpdateCommissionRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingUpdateCommissionRate)
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
		it.Event = new(StakingUpdateCommissionRate)
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
func (it *StakingUpdateCommissionRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingUpdateCommissionRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingUpdateCommissionRate represents a UpdateCommissionRate event raised by the Staking contract.
type StakingUpdateCommissionRate struct {
	ValidatorId       *big.Int
	NewCommissionRate *big.Int
	OldCommissionRate *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterUpdateCommissionRate is a free log retrieval operation binding the contract event 0x7d5da5ece9d43013d62ab966f4704ca376b92be29ca6fbb958154baf1c0dc17e.
//
// Solidity: event UpdateCommissionRate(uint256 indexed validatorId, uint256 indexed newCommissionRate, uint256 indexed oldCommissionRate)
func (_Staking *StakingFilterer) FilterUpdateCommissionRate(opts *bind.FilterOpts, validatorId []*big.Int, newCommissionRate []*big.Int, oldCommissionRate []*big.Int) (*StakingUpdateCommissionRateIterator, error) {

	var validatorIdRule []interface{}
	for _, validatorIdItem := range validatorId {
		validatorIdRule = append(validatorIdRule, validatorIdItem)
	}
	var newCommissionRateRule []interface{}
	for _, newCommissionRateItem := range newCommissionRate {
		newCommissionRateRule = append(newCommissionRateRule, newCommissionRateItem)
	}
	var oldCommissionRateRule []interface{}
	for _, oldCommissionRateItem := range oldCommissionRate {
		oldCommissionRateRule = append(oldCommissionRateRule, oldCommissionRateItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "UpdateCommissionRate", validatorIdRule, newCommissionRateRule, oldCommissionRateRule)
	if err != nil {
		return nil, err
	}
	return &StakingUpdateCommissionRateIterator{contract: _Staking.contract, event: "UpdateCommissionRate", logs: logs, sub: sub}, nil
}

// WatchUpdateCommissionRate is a free log subscription operation binding the contract event 0x7d5da5ece9d43013d62ab966f4704ca376b92be29ca6fbb958154baf1c0dc17e.
//
// Solidity: event UpdateCommissionRate(uint256 indexed validatorId, uint256 indexed newCommissionRate, uint256 indexed oldCommissionRate)
func (_Staking *StakingFilterer) WatchUpdateCommissionRate(opts *bind.WatchOpts, sink chan<- *StakingUpdateCommissionRate, validatorId []*big.Int, newCommissionRate []*big.Int, oldCommissionRate []*big.Int) (event.Subscription, error) {

	var validatorIdRule []interface{}
	for _, validatorIdItem := range validatorId {
		validatorIdRule = append(validatorIdRule, validatorIdItem)
	}
	var newCommissionRateRule []interface{}
	for _, newCommissionRateItem := range newCommissionRate {
		newCommissionRateRule = append(newCommissionRateRule, newCommissionRateItem)
	}
	var oldCommissionRateRule []interface{}
	for _, oldCommissionRateItem := range oldCommissionRate {
		oldCommissionRateRule = append(oldCommissionRateRule, oldCommissionRateItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "UpdateCommissionRate", validatorIdRule, newCommissionRateRule, oldCommissionRateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingUpdateCommissionRate)
				if err := _Staking.contract.UnpackLog(event, "UpdateCommissionRate", log); err != nil {
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

// ParseUpdateCommissionRate is a log parse operation binding the contract event 0x7d5da5ece9d43013d62ab966f4704ca376b92be29ca6fbb958154baf1c0dc17e.
//
// Solidity: event UpdateCommissionRate(uint256 indexed validatorId, uint256 indexed newCommissionRate, uint256 indexed oldCommissionRate)
func (_Staking *StakingFilterer) ParseUpdateCommissionRate(log types.Log) (*StakingUpdateCommissionRate, error) {
	event := new(StakingUpdateCommissionRate)
	if err := _Staking.contract.UnpackLog(event, "UpdateCommissionRate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
