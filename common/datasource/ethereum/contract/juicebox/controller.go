// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package juicebox

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

// JBFundAccessConstraints is an auto generated low-level Go binding around an user-defined struct.
type JBFundAccessConstraints struct {
	Terminal                  common.Address
	Token                     common.Address
	DistributionLimit         *big.Int
	DistributionLimitCurrency *big.Int
	OverflowAllowance         *big.Int
	OverflowAllowanceCurrency *big.Int
}

// JBFundingCycle is an auto generated low-level Go binding around an user-defined struct.
type JBFundingCycle struct {
	Number        *big.Int
	Configuration *big.Int
	BasedOn       *big.Int
	Start         *big.Int
	Duration      *big.Int
	Weight        *big.Int
	DiscountRate  *big.Int
	Ballot        common.Address
	Metadata      *big.Int
}

// JBFundingCycleData is an auto generated low-level Go binding around an user-defined struct.
type JBFundingCycleData struct {
	Duration     *big.Int
	Weight       *big.Int
	DiscountRate *big.Int
	Ballot       common.Address
}

// JBFundingCycleMetadata is an auto generated low-level Go binding around an user-defined struct.
type JBFundingCycleMetadata struct {
	Global                         JBGlobalFundingCycleMetadata
	ReservedRate                   *big.Int
	RedemptionRate                 *big.Int
	BallotRedemptionRate           *big.Int
	PausePay                       bool
	PauseDistributions             bool
	PauseRedeem                    bool
	PauseBurn                      bool
	AllowMinting                   bool
	AllowTerminalMigration         bool
	AllowControllerMigration       bool
	HoldFees                       bool
	PreferClaimedTokenOverride     bool
	UseTotalOverflowForRedemptions bool
	UseDataSourceForPay            bool
	UseDataSourceForRedeem         bool
	DataSource                     common.Address
	Metadata                       *big.Int
}

// JBGlobalFundingCycleMetadata is an auto generated low-level Go binding around an user-defined struct.
type JBGlobalFundingCycleMetadata struct {
	AllowSetTerminals  bool
	AllowSetController bool
	PauseTransfers     bool
}

// JBGroupedSplits is an auto generated low-level Go binding around an user-defined struct.
type JBGroupedSplits struct {
	Group  *big.Int
	Splits []JBSplit
}

// ControllerMetaData contains all meta data concerning the Controller contract.
var ControllerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIJBOperatorStore\",\"name\":\"_operatorStore\",\"type\":\"address\"},{\"internalType\":\"contractIJBProjects\",\"name\":\"_projects\",\"type\":\"address\"},{\"internalType\":\"contractIJBDirectory\",\"name\":\"_directory\",\"type\":\"address\"},{\"internalType\":\"contractIJBFundingCycleStore\",\"name\":\"_fundingCycleStore\",\"type\":\"address\"},{\"internalType\":\"contractIJBTokenStore\",\"name\":\"_tokenStore\",\"type\":\"address\"},{\"internalType\":\"contractIJBSplitsStore\",\"name\":\"_splitsStore\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BURN_PAUSED_AND_SENDER_NOT_VALID_TERMINAL_DELEGATE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CANT_MIGRATE_TO_CURRENT_CONTROLLER\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FUNDING_CYCLE_ALREADY_LAUNCHED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"INVALID_BALLOT_REDEMPTION_RATE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"INVALID_DISTRIBUTION_LIMIT\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"INVALID_DISTRIBUTION_LIMIT_CURRENCY\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"INVALID_OVERFLOW_ALLOWANCE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"INVALID_OVERFLOW_ALLOWANCE_CURRENCY\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"INVALID_REDEMPTION_RATE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"INVALID_RESERVED_RATE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MIGRATION_NOT_ALLOWED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MINT_NOT_ALLOWED_AND_NOT_TERMINAL_DELEGATE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NOT_CURRENT_CONTROLLER\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NO_BURNABLE_TOKENS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OVERFLOW_ALERT\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"prod1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"denominator\",\"type\":\"uint256\"}],\"name\":\"PRBMath__MulDivOverflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UNAUTHORIZED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZERO_TOKENS_TO_MINT\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"projectId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"memo\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"BurnTokens\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fundingCycleConfiguration\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fundingCycleNumber\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"projectId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"beneficiaryTokenCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"memo\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"DistributeReservedTokens\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"projectId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"domain\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"group\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"preferClaimed\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"preferAddToBalance\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"percent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"projectId\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lockedUntil\",\"type\":\"uint256\"},{\"internalType\":\"contractIJBSplitAllocator\",\"name\":\"allocator\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structJBSplit\",\"name\":\"split\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"DistributeToReservedTokenSplit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"configuration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"projectId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"memo\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"LaunchFundingCycles\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"configuration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"projectId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"memo\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"LaunchProject\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"projectId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIJBMigratable\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"Migrate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"projectId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"beneficiaryTokenCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"memo\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reservedRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"MintTokens\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"projectId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"PrepMigration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"configuration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"projectId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"memo\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"ReconfigureFundingCycles\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fundingCycleConfiguration\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fundingCycleNumber\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"projectId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractIJBPaymentTerminal\",\"name\":\"terminal\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"distributionLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"distributionLimitCurrency\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"overflowAllowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"overflowAllowanceCurrency\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structJBFundAccessConstraints\",\"name\":\"constraints\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"SetFundAccessConstraints\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_holder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tokenCount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_memo\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"_preferClaimedTokens\",\"type\":\"bool\"}],\"name\":\"burnTokensOf\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"}],\"name\":\"currentFundingCycleOf\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"configuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"basedOn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"discountRate\",\"type\":\"uint256\"},{\"internalType\":\"contractIJBFundingCycleBallot\",\"name\":\"ballot\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"metadata\",\"type\":\"uint256\"}],\"internalType\":\"structJBFundingCycle\",\"name\":\"fundingCycle\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"allowSetTerminals\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"allowSetController\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"pauseTransfers\",\"type\":\"bool\"}],\"internalType\":\"structJBGlobalFundingCycleMetadata\",\"name\":\"global\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"reservedRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"redemptionRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ballotRedemptionRate\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"pausePay\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"pauseDistributions\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"pauseRedeem\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"pauseBurn\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"allowMinting\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"allowTerminalMigration\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"allowControllerMigration\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"holdFees\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"preferClaimedTokenOverride\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"useTotalOverflowForRedemptions\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"useDataSourceForPay\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"useDataSourceForRedeem\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"dataSource\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"metadata\",\"type\":\"uint256\"}],\"internalType\":\"structJBFundingCycleMetadata\",\"name\":\"metadata\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"directory\",\"outputs\":[{\"internalType\":\"contractIJBDirectory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_memo\",\"type\":\"string\"}],\"name\":\"distributeReservedTokensOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_configuration\",\"type\":\"uint256\"},{\"internalType\":\"contractIJBPaymentTerminal\",\"name\":\"_terminal\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"distributionLimitOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fundingCycleStore\",\"outputs\":[{\"internalType\":\"contractIJBFundingCycleStore\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_configuration\",\"type\":\"uint256\"}],\"name\":\"getFundingCycleOf\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"configuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"basedOn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"discountRate\",\"type\":\"uint256\"},{\"internalType\":\"contractIJBFundingCycleBallot\",\"name\":\"ballot\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"metadata\",\"type\":\"uint256\"}],\"internalType\":\"structJBFundingCycle\",\"name\":\"fundingCycle\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"allowSetTerminals\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"allowSetController\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"pauseTransfers\",\"type\":\"bool\"}],\"internalType\":\"structJBGlobalFundingCycleMetadata\",\"name\":\"global\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"reservedRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"redemptionRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ballotRedemptionRate\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"pausePay\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"pauseDistributions\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"pauseRedeem\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"pauseBurn\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"allowMinting\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"allowTerminalMigration\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"allowControllerMigration\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"holdFees\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"preferClaimedTokenOverride\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"useTotalOverflowForRedemptions\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"useDataSourceForPay\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"useDataSourceForRedeem\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"dataSource\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"metadata\",\"type\":\"uint256\"}],\"internalType\":\"structJBFundingCycleMetadata\",\"name\":\"metadata\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"}],\"name\":\"latestConfiguredFundingCycleOf\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"configuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"basedOn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"discountRate\",\"type\":\"uint256\"},{\"internalType\":\"contractIJBFundingCycleBallot\",\"name\":\"ballot\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"metadata\",\"type\":\"uint256\"}],\"internalType\":\"structJBFundingCycle\",\"name\":\"fundingCycle\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"allowSetTerminals\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"allowSetController\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"pauseTransfers\",\"type\":\"bool\"}],\"internalType\":\"structJBGlobalFundingCycleMetadata\",\"name\":\"global\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"reservedRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"redemptionRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ballotRedemptionRate\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"pausePay\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"pauseDistributions\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"pauseRedeem\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"pauseBurn\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"allowMinting\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"allowTerminalMigration\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"allowControllerMigration\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"holdFees\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"preferClaimedTokenOverride\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"useTotalOverflowForRedemptions\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"useDataSourceForPay\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"useDataSourceForRedeem\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"dataSource\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"metadata\",\"type\":\"uint256\"}],\"internalType\":\"structJBFundingCycleMetadata\",\"name\":\"metadata\",\"type\":\"tuple\"},{\"internalType\":\"enumJBBallotState\",\"name\":\"ballotState\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"discountRate\",\"type\":\"uint256\"},{\"internalType\":\"contractIJBFundingCycleBallot\",\"name\":\"ballot\",\"type\":\"address\"}],\"internalType\":\"structJBFundingCycleData\",\"name\":\"_data\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"allowSetTerminals\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"allowSetController\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"pauseTransfers\",\"type\":\"bool\"}],\"internalType\":\"structJBGlobalFundingCycleMetadata\",\"name\":\"global\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"reservedRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"redemptionRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ballotRedemptionRate\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"pausePay\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"pauseDistributions\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"pauseRedeem\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"pauseBurn\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"allowMinting\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"allowTerminalMigration\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"allowControllerMigration\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"holdFees\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"preferClaimedTokenOverride\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"useTotalOverflowForRedemptions\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"useDataSourceForPay\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"useDataSourceForRedeem\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"dataSource\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"metadata\",\"type\":\"uint256\"}],\"internalType\":\"structJBFundingCycleMetadata\",\"name\":\"_metadata\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_mustStartAtOrAfter\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"group\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"preferClaimed\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"preferAddToBalance\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"percent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"projectId\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lockedUntil\",\"type\":\"uint256\"},{\"internalType\":\"contractIJBSplitAllocator\",\"name\":\"allocator\",\"type\":\"address\"}],\"internalType\":\"structJBSplit[]\",\"name\":\"splits\",\"type\":\"tuple[]\"}],\"internalType\":\"structJBGroupedSplits[]\",\"name\":\"_groupedSplits\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"contractIJBPaymentTerminal\",\"name\":\"terminal\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"distributionLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"distributionLimitCurrency\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"overflowAllowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"overflowAllowanceCurrency\",\"type\":\"uint256\"}],\"internalType\":\"structJBFundAccessConstraints[]\",\"name\":\"_fundAccessConstraints\",\"type\":\"tuple[]\"},{\"internalType\":\"contractIJBPaymentTerminal[]\",\"name\":\"_terminals\",\"type\":\"address[]\"},{\"internalType\":\"string\",\"name\":\"_memo\",\"type\":\"string\"}],\"name\":\"launchFundingCyclesFor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"configuration\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"domain\",\"type\":\"uint256\"}],\"internalType\":\"structJBProjectMetadata\",\"name\":\"_projectMetadata\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"discountRate\",\"type\":\"uint256\"},{\"internalType\":\"contractIJBFundingCycleBallot\",\"name\":\"ballot\",\"type\":\"address\"}],\"internalType\":\"structJBFundingCycleData\",\"name\":\"_data\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"allowSetTerminals\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"allowSetController\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"pauseTransfers\",\"type\":\"bool\"}],\"internalType\":\"structJBGlobalFundingCycleMetadata\",\"name\":\"global\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"reservedRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"redemptionRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ballotRedemptionRate\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"pausePay\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"pauseDistributions\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"pauseRedeem\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"pauseBurn\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"allowMinting\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"allowTerminalMigration\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"allowControllerMigration\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"holdFees\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"preferClaimedTokenOverride\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"useTotalOverflowForRedemptions\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"useDataSourceForPay\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"useDataSourceForRedeem\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"dataSource\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"metadata\",\"type\":\"uint256\"}],\"internalType\":\"structJBFundingCycleMetadata\",\"name\":\"_metadata\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_mustStartAtOrAfter\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"group\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"preferClaimed\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"preferAddToBalance\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"percent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"projectId\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lockedUntil\",\"type\":\"uint256\"},{\"internalType\":\"contractIJBSplitAllocator\",\"name\":\"allocator\",\"type\":\"address\"}],\"internalType\":\"structJBSplit[]\",\"name\":\"splits\",\"type\":\"tuple[]\"}],\"internalType\":\"structJBGroupedSplits[]\",\"name\":\"_groupedSplits\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"contractIJBPaymentTerminal\",\"name\":\"terminal\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"distributionLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"distributionLimitCurrency\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"overflowAllowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"overflowAllowanceCurrency\",\"type\":\"uint256\"}],\"internalType\":\"structJBFundAccessConstraints[]\",\"name\":\"_fundAccessConstraints\",\"type\":\"tuple[]\"},{\"internalType\":\"contractIJBPaymentTerminal[]\",\"name\":\"_terminals\",\"type\":\"address[]\"},{\"internalType\":\"string\",\"name\":\"_memo\",\"type\":\"string\"}],\"name\":\"launchProjectFor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"projectId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"contractIJBMigratable\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"migrate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tokenCount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_beneficiary\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_memo\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"_preferClaimedTokens\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_useReservedRate\",\"type\":\"bool\"}],\"name\":\"mintTokensOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"beneficiaryTokenCount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operatorStore\",\"outputs\":[{\"internalType\":\"contractIJBOperatorStore\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_configuration\",\"type\":\"uint256\"},{\"internalType\":\"contractIJBPaymentTerminal\",\"name\":\"_terminal\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"overflowAllowanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"}],\"name\":\"prepForMigrationOf\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"projects\",\"outputs\":[{\"internalType\":\"contractIJBProjects\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"}],\"name\":\"queuedFundingCycleOf\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"configuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"basedOn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"discountRate\",\"type\":\"uint256\"},{\"internalType\":\"contractIJBFundingCycleBallot\",\"name\":\"ballot\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"metadata\",\"type\":\"uint256\"}],\"internalType\":\"structJBFundingCycle\",\"name\":\"fundingCycle\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"allowSetTerminals\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"allowSetController\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"pauseTransfers\",\"type\":\"bool\"}],\"internalType\":\"structJBGlobalFundingCycleMetadata\",\"name\":\"global\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"reservedRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"redemptionRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ballotRedemptionRate\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"pausePay\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"pauseDistributions\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"pauseRedeem\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"pauseBurn\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"allowMinting\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"allowTerminalMigration\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"allowControllerMigration\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"holdFees\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"preferClaimedTokenOverride\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"useTotalOverflowForRedemptions\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"useDataSourceForPay\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"useDataSourceForRedeem\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"dataSource\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"metadata\",\"type\":\"uint256\"}],\"internalType\":\"structJBFundingCycleMetadata\",\"name\":\"metadata\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"discountRate\",\"type\":\"uint256\"},{\"internalType\":\"contractIJBFundingCycleBallot\",\"name\":\"ballot\",\"type\":\"address\"}],\"internalType\":\"structJBFundingCycleData\",\"name\":\"_data\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"allowSetTerminals\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"allowSetController\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"pauseTransfers\",\"type\":\"bool\"}],\"internalType\":\"structJBGlobalFundingCycleMetadata\",\"name\":\"global\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"reservedRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"redemptionRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ballotRedemptionRate\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"pausePay\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"pauseDistributions\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"pauseRedeem\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"pauseBurn\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"allowMinting\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"allowTerminalMigration\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"allowControllerMigration\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"holdFees\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"preferClaimedTokenOverride\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"useTotalOverflowForRedemptions\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"useDataSourceForPay\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"useDataSourceForRedeem\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"dataSource\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"metadata\",\"type\":\"uint256\"}],\"internalType\":\"structJBFundingCycleMetadata\",\"name\":\"_metadata\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_mustStartAtOrAfter\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"group\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"preferClaimed\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"preferAddToBalance\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"percent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"projectId\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lockedUntil\",\"type\":\"uint256\"},{\"internalType\":\"contractIJBSplitAllocator\",\"name\":\"allocator\",\"type\":\"address\"}],\"internalType\":\"structJBSplit[]\",\"name\":\"splits\",\"type\":\"tuple[]\"}],\"internalType\":\"structJBGroupedSplits[]\",\"name\":\"_groupedSplits\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"contractIJBPaymentTerminal\",\"name\":\"terminal\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"distributionLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"distributionLimitCurrency\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"overflowAllowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"overflowAllowanceCurrency\",\"type\":\"uint256\"}],\"internalType\":\"structJBFundAccessConstraints[]\",\"name\":\"_fundAccessConstraints\",\"type\":\"tuple[]\"},{\"internalType\":\"string\",\"name\":\"_memo\",\"type\":\"string\"}],\"name\":\"reconfigureFundingCyclesOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"configuration\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_reservedRate\",\"type\":\"uint256\"}],\"name\":\"reservedTokenBalanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"splitsStore\",\"outputs\":[{\"internalType\":\"contractIJBSplitsStore\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenStore\",\"outputs\":[{\"internalType\":\"contractIJBTokenStore\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_reservedRate\",\"type\":\"uint256\"}],\"name\":\"totalOutstandingTokensOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ControllerABI is the input ABI used to generate the binding from.
// Deprecated: Use ControllerMetaData.ABI instead.
var ControllerABI = ControllerMetaData.ABI

// Controller is an auto generated Go binding around an Ethereum contract.
type Controller struct {
	ControllerCaller     // Read-only binding to the contract
	ControllerTransactor // Write-only binding to the contract
	ControllerFilterer   // Log filterer for contract events
}

// ControllerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ControllerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ControllerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ControllerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ControllerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ControllerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ControllerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ControllerSession struct {
	Contract     *Controller       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ControllerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ControllerCallerSession struct {
	Contract *ControllerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ControllerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ControllerTransactorSession struct {
	Contract     *ControllerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ControllerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ControllerRaw struct {
	Contract *Controller // Generic contract binding to access the raw methods on
}

// ControllerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ControllerCallerRaw struct {
	Contract *ControllerCaller // Generic read-only contract binding to access the raw methods on
}

// ControllerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ControllerTransactorRaw struct {
	Contract *ControllerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewController creates a new instance of Controller, bound to a specific deployed contract.
func NewController(address common.Address, backend bind.ContractBackend) (*Controller, error) {
	contract, err := bindController(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Controller{ControllerCaller: ControllerCaller{contract: contract}, ControllerTransactor: ControllerTransactor{contract: contract}, ControllerFilterer: ControllerFilterer{contract: contract}}, nil
}

// NewControllerCaller creates a new read-only instance of Controller, bound to a specific deployed contract.
func NewControllerCaller(address common.Address, caller bind.ContractCaller) (*ControllerCaller, error) {
	contract, err := bindController(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ControllerCaller{contract: contract}, nil
}

// NewControllerTransactor creates a new write-only instance of Controller, bound to a specific deployed contract.
func NewControllerTransactor(address common.Address, transactor bind.ContractTransactor) (*ControllerTransactor, error) {
	contract, err := bindController(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ControllerTransactor{contract: contract}, nil
}

// NewControllerFilterer creates a new log filterer instance of Controller, bound to a specific deployed contract.
func NewControllerFilterer(address common.Address, filterer bind.ContractFilterer) (*ControllerFilterer, error) {
	contract, err := bindController(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ControllerFilterer{contract: contract}, nil
}

// bindController binds a generic wrapper to an already deployed contract.
func bindController(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ControllerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Controller *ControllerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Controller.Contract.ControllerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Controller *ControllerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Controller.Contract.ControllerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Controller *ControllerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Controller.Contract.ControllerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Controller *ControllerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Controller.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Controller *ControllerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Controller.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Controller *ControllerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Controller.Contract.contract.Transact(opts, method, params...)
}

// CurrentFundingCycleOf is a free data retrieval call binding the contract method 0x8776c499.
//
// Solidity: function currentFundingCycleOf(uint256 _projectId) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,uint256) fundingCycle, ((bool,bool,bool),uint256,uint256,uint256,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,address,uint256) metadata)
func (_Controller *ControllerCaller) CurrentFundingCycleOf(opts *bind.CallOpts, _projectId *big.Int) (struct {
	FundingCycle JBFundingCycle
	Metadata     JBFundingCycleMetadata
}, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "currentFundingCycleOf", _projectId)

	outstruct := new(struct {
		FundingCycle JBFundingCycle
		Metadata     JBFundingCycleMetadata
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FundingCycle = *abi.ConvertType(out[0], new(JBFundingCycle)).(*JBFundingCycle)
	outstruct.Metadata = *abi.ConvertType(out[1], new(JBFundingCycleMetadata)).(*JBFundingCycleMetadata)

	return *outstruct, err

}

// CurrentFundingCycleOf is a free data retrieval call binding the contract method 0x8776c499.
//
// Solidity: function currentFundingCycleOf(uint256 _projectId) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,uint256) fundingCycle, ((bool,bool,bool),uint256,uint256,uint256,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,address,uint256) metadata)
func (_Controller *ControllerSession) CurrentFundingCycleOf(_projectId *big.Int) (struct {
	FundingCycle JBFundingCycle
	Metadata     JBFundingCycleMetadata
}, error) {
	return _Controller.Contract.CurrentFundingCycleOf(&_Controller.CallOpts, _projectId)
}

// CurrentFundingCycleOf is a free data retrieval call binding the contract method 0x8776c499.
//
// Solidity: function currentFundingCycleOf(uint256 _projectId) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,uint256) fundingCycle, ((bool,bool,bool),uint256,uint256,uint256,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,address,uint256) metadata)
func (_Controller *ControllerCallerSession) CurrentFundingCycleOf(_projectId *big.Int) (struct {
	FundingCycle JBFundingCycle
	Metadata     JBFundingCycleMetadata
}, error) {
	return _Controller.Contract.CurrentFundingCycleOf(&_Controller.CallOpts, _projectId)
}

// Directory is a free data retrieval call binding the contract method 0xc41c2f24.
//
// Solidity: function directory() view returns(address)
func (_Controller *ControllerCaller) Directory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "directory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Directory is a free data retrieval call binding the contract method 0xc41c2f24.
//
// Solidity: function directory() view returns(address)
func (_Controller *ControllerSession) Directory() (common.Address, error) {
	return _Controller.Contract.Directory(&_Controller.CallOpts)
}

// Directory is a free data retrieval call binding the contract method 0xc41c2f24.
//
// Solidity: function directory() view returns(address)
func (_Controller *ControllerCallerSession) Directory() (common.Address, error) {
	return _Controller.Contract.Directory(&_Controller.CallOpts)
}

// DistributionLimitOf is a free data retrieval call binding the contract method 0xe8db2130.
//
// Solidity: function distributionLimitOf(uint256 _projectId, uint256 _configuration, address _terminal, address _token) view returns(uint256, uint256)
func (_Controller *ControllerCaller) DistributionLimitOf(opts *bind.CallOpts, _projectId *big.Int, _configuration *big.Int, _terminal common.Address, _token common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "distributionLimitOf", _projectId, _configuration, _terminal, _token)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// DistributionLimitOf is a free data retrieval call binding the contract method 0xe8db2130.
//
// Solidity: function distributionLimitOf(uint256 _projectId, uint256 _configuration, address _terminal, address _token) view returns(uint256, uint256)
func (_Controller *ControllerSession) DistributionLimitOf(_projectId *big.Int, _configuration *big.Int, _terminal common.Address, _token common.Address) (*big.Int, *big.Int, error) {
	return _Controller.Contract.DistributionLimitOf(&_Controller.CallOpts, _projectId, _configuration, _terminal, _token)
}

// DistributionLimitOf is a free data retrieval call binding the contract method 0xe8db2130.
//
// Solidity: function distributionLimitOf(uint256 _projectId, uint256 _configuration, address _terminal, address _token) view returns(uint256, uint256)
func (_Controller *ControllerCallerSession) DistributionLimitOf(_projectId *big.Int, _configuration *big.Int, _terminal common.Address, _token common.Address) (*big.Int, *big.Int, error) {
	return _Controller.Contract.DistributionLimitOf(&_Controller.CallOpts, _projectId, _configuration, _terminal, _token)
}

// FundingCycleStore is a free data retrieval call binding the contract method 0x557e7155.
//
// Solidity: function fundingCycleStore() view returns(address)
func (_Controller *ControllerCaller) FundingCycleStore(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "fundingCycleStore")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FundingCycleStore is a free data retrieval call binding the contract method 0x557e7155.
//
// Solidity: function fundingCycleStore() view returns(address)
func (_Controller *ControllerSession) FundingCycleStore() (common.Address, error) {
	return _Controller.Contract.FundingCycleStore(&_Controller.CallOpts)
}

// FundingCycleStore is a free data retrieval call binding the contract method 0x557e7155.
//
// Solidity: function fundingCycleStore() view returns(address)
func (_Controller *ControllerCallerSession) FundingCycleStore() (common.Address, error) {
	return _Controller.Contract.FundingCycleStore(&_Controller.CallOpts)
}

// GetFundingCycleOf is a free data retrieval call binding the contract method 0xa40bb9c7.
//
// Solidity: function getFundingCycleOf(uint256 _projectId, uint256 _configuration) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,uint256) fundingCycle, ((bool,bool,bool),uint256,uint256,uint256,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,address,uint256) metadata)
func (_Controller *ControllerCaller) GetFundingCycleOf(opts *bind.CallOpts, _projectId *big.Int, _configuration *big.Int) (struct {
	FundingCycle JBFundingCycle
	Metadata     JBFundingCycleMetadata
}, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "getFundingCycleOf", _projectId, _configuration)

	outstruct := new(struct {
		FundingCycle JBFundingCycle
		Metadata     JBFundingCycleMetadata
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FundingCycle = *abi.ConvertType(out[0], new(JBFundingCycle)).(*JBFundingCycle)
	outstruct.Metadata = *abi.ConvertType(out[1], new(JBFundingCycleMetadata)).(*JBFundingCycleMetadata)

	return *outstruct, err

}

// GetFundingCycleOf is a free data retrieval call binding the contract method 0xa40bb9c7.
//
// Solidity: function getFundingCycleOf(uint256 _projectId, uint256 _configuration) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,uint256) fundingCycle, ((bool,bool,bool),uint256,uint256,uint256,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,address,uint256) metadata)
func (_Controller *ControllerSession) GetFundingCycleOf(_projectId *big.Int, _configuration *big.Int) (struct {
	FundingCycle JBFundingCycle
	Metadata     JBFundingCycleMetadata
}, error) {
	return _Controller.Contract.GetFundingCycleOf(&_Controller.CallOpts, _projectId, _configuration)
}

// GetFundingCycleOf is a free data retrieval call binding the contract method 0xa40bb9c7.
//
// Solidity: function getFundingCycleOf(uint256 _projectId, uint256 _configuration) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,uint256) fundingCycle, ((bool,bool,bool),uint256,uint256,uint256,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,address,uint256) metadata)
func (_Controller *ControllerCallerSession) GetFundingCycleOf(_projectId *big.Int, _configuration *big.Int) (struct {
	FundingCycle JBFundingCycle
	Metadata     JBFundingCycleMetadata
}, error) {
	return _Controller.Contract.GetFundingCycleOf(&_Controller.CallOpts, _projectId, _configuration)
}

// LatestConfiguredFundingCycleOf is a free data retrieval call binding the contract method 0x1f510453.
//
// Solidity: function latestConfiguredFundingCycleOf(uint256 _projectId) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,uint256) fundingCycle, ((bool,bool,bool),uint256,uint256,uint256,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,address,uint256) metadata, uint8 ballotState)
func (_Controller *ControllerCaller) LatestConfiguredFundingCycleOf(opts *bind.CallOpts, _projectId *big.Int) (struct {
	FundingCycle JBFundingCycle
	Metadata     JBFundingCycleMetadata
	BallotState  uint8
}, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "latestConfiguredFundingCycleOf", _projectId)

	outstruct := new(struct {
		FundingCycle JBFundingCycle
		Metadata     JBFundingCycleMetadata
		BallotState  uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FundingCycle = *abi.ConvertType(out[0], new(JBFundingCycle)).(*JBFundingCycle)
	outstruct.Metadata = *abi.ConvertType(out[1], new(JBFundingCycleMetadata)).(*JBFundingCycleMetadata)
	outstruct.BallotState = *abi.ConvertType(out[2], new(uint8)).(*uint8)

	return *outstruct, err

}

// LatestConfiguredFundingCycleOf is a free data retrieval call binding the contract method 0x1f510453.
//
// Solidity: function latestConfiguredFundingCycleOf(uint256 _projectId) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,uint256) fundingCycle, ((bool,bool,bool),uint256,uint256,uint256,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,address,uint256) metadata, uint8 ballotState)
func (_Controller *ControllerSession) LatestConfiguredFundingCycleOf(_projectId *big.Int) (struct {
	FundingCycle JBFundingCycle
	Metadata     JBFundingCycleMetadata
	BallotState  uint8
}, error) {
	return _Controller.Contract.LatestConfiguredFundingCycleOf(&_Controller.CallOpts, _projectId)
}

// LatestConfiguredFundingCycleOf is a free data retrieval call binding the contract method 0x1f510453.
//
// Solidity: function latestConfiguredFundingCycleOf(uint256 _projectId) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,uint256) fundingCycle, ((bool,bool,bool),uint256,uint256,uint256,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,address,uint256) metadata, uint8 ballotState)
func (_Controller *ControllerCallerSession) LatestConfiguredFundingCycleOf(_projectId *big.Int) (struct {
	FundingCycle JBFundingCycle
	Metadata     JBFundingCycleMetadata
	BallotState  uint8
}, error) {
	return _Controller.Contract.LatestConfiguredFundingCycleOf(&_Controller.CallOpts, _projectId)
}

// OperatorStore is a free data retrieval call binding the contract method 0xad007d63.
//
// Solidity: function operatorStore() view returns(address)
func (_Controller *ControllerCaller) OperatorStore(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "operatorStore")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OperatorStore is a free data retrieval call binding the contract method 0xad007d63.
//
// Solidity: function operatorStore() view returns(address)
func (_Controller *ControllerSession) OperatorStore() (common.Address, error) {
	return _Controller.Contract.OperatorStore(&_Controller.CallOpts)
}

// OperatorStore is a free data retrieval call binding the contract method 0xad007d63.
//
// Solidity: function operatorStore() view returns(address)
func (_Controller *ControllerCallerSession) OperatorStore() (common.Address, error) {
	return _Controller.Contract.OperatorStore(&_Controller.CallOpts)
}

// OverflowAllowanceOf is a free data retrieval call binding the contract method 0x7a81b562.
//
// Solidity: function overflowAllowanceOf(uint256 _projectId, uint256 _configuration, address _terminal, address _token) view returns(uint256, uint256)
func (_Controller *ControllerCaller) OverflowAllowanceOf(opts *bind.CallOpts, _projectId *big.Int, _configuration *big.Int, _terminal common.Address, _token common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "overflowAllowanceOf", _projectId, _configuration, _terminal, _token)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// OverflowAllowanceOf is a free data retrieval call binding the contract method 0x7a81b562.
//
// Solidity: function overflowAllowanceOf(uint256 _projectId, uint256 _configuration, address _terminal, address _token) view returns(uint256, uint256)
func (_Controller *ControllerSession) OverflowAllowanceOf(_projectId *big.Int, _configuration *big.Int, _terminal common.Address, _token common.Address) (*big.Int, *big.Int, error) {
	return _Controller.Contract.OverflowAllowanceOf(&_Controller.CallOpts, _projectId, _configuration, _terminal, _token)
}

// OverflowAllowanceOf is a free data retrieval call binding the contract method 0x7a81b562.
//
// Solidity: function overflowAllowanceOf(uint256 _projectId, uint256 _configuration, address _terminal, address _token) view returns(uint256, uint256)
func (_Controller *ControllerCallerSession) OverflowAllowanceOf(_projectId *big.Int, _configuration *big.Int, _terminal common.Address, _token common.Address) (*big.Int, *big.Int, error) {
	return _Controller.Contract.OverflowAllowanceOf(&_Controller.CallOpts, _projectId, _configuration, _terminal, _token)
}

// Projects is a free data retrieval call binding the contract method 0x8b79543c.
//
// Solidity: function projects() view returns(address)
func (_Controller *ControllerCaller) Projects(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "projects")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Projects is a free data retrieval call binding the contract method 0x8b79543c.
//
// Solidity: function projects() view returns(address)
func (_Controller *ControllerSession) Projects() (common.Address, error) {
	return _Controller.Contract.Projects(&_Controller.CallOpts)
}

// Projects is a free data retrieval call binding the contract method 0x8b79543c.
//
// Solidity: function projects() view returns(address)
func (_Controller *ControllerCallerSession) Projects() (common.Address, error) {
	return _Controller.Contract.Projects(&_Controller.CallOpts)
}

// QueuedFundingCycleOf is a free data retrieval call binding the contract method 0x12b37b14.
//
// Solidity: function queuedFundingCycleOf(uint256 _projectId) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,uint256) fundingCycle, ((bool,bool,bool),uint256,uint256,uint256,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,address,uint256) metadata)
func (_Controller *ControllerCaller) QueuedFundingCycleOf(opts *bind.CallOpts, _projectId *big.Int) (struct {
	FundingCycle JBFundingCycle
	Metadata     JBFundingCycleMetadata
}, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "queuedFundingCycleOf", _projectId)

	outstruct := new(struct {
		FundingCycle JBFundingCycle
		Metadata     JBFundingCycleMetadata
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FundingCycle = *abi.ConvertType(out[0], new(JBFundingCycle)).(*JBFundingCycle)
	outstruct.Metadata = *abi.ConvertType(out[1], new(JBFundingCycleMetadata)).(*JBFundingCycleMetadata)

	return *outstruct, err

}

// QueuedFundingCycleOf is a free data retrieval call binding the contract method 0x12b37b14.
//
// Solidity: function queuedFundingCycleOf(uint256 _projectId) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,uint256) fundingCycle, ((bool,bool,bool),uint256,uint256,uint256,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,address,uint256) metadata)
func (_Controller *ControllerSession) QueuedFundingCycleOf(_projectId *big.Int) (struct {
	FundingCycle JBFundingCycle
	Metadata     JBFundingCycleMetadata
}, error) {
	return _Controller.Contract.QueuedFundingCycleOf(&_Controller.CallOpts, _projectId)
}

// QueuedFundingCycleOf is a free data retrieval call binding the contract method 0x12b37b14.
//
// Solidity: function queuedFundingCycleOf(uint256 _projectId) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,uint256) fundingCycle, ((bool,bool,bool),uint256,uint256,uint256,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,address,uint256) metadata)
func (_Controller *ControllerCallerSession) QueuedFundingCycleOf(_projectId *big.Int) (struct {
	FundingCycle JBFundingCycle
	Metadata     JBFundingCycleMetadata
}, error) {
	return _Controller.Contract.QueuedFundingCycleOf(&_Controller.CallOpts, _projectId)
}

// ReservedTokenBalanceOf is a free data retrieval call binding the contract method 0xf033b3b4.
//
// Solidity: function reservedTokenBalanceOf(uint256 _projectId, uint256 _reservedRate) view returns(uint256)
func (_Controller *ControllerCaller) ReservedTokenBalanceOf(opts *bind.CallOpts, _projectId *big.Int, _reservedRate *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "reservedTokenBalanceOf", _projectId, _reservedRate)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReservedTokenBalanceOf is a free data retrieval call binding the contract method 0xf033b3b4.
//
// Solidity: function reservedTokenBalanceOf(uint256 _projectId, uint256 _reservedRate) view returns(uint256)
func (_Controller *ControllerSession) ReservedTokenBalanceOf(_projectId *big.Int, _reservedRate *big.Int) (*big.Int, error) {
	return _Controller.Contract.ReservedTokenBalanceOf(&_Controller.CallOpts, _projectId, _reservedRate)
}

// ReservedTokenBalanceOf is a free data retrieval call binding the contract method 0xf033b3b4.
//
// Solidity: function reservedTokenBalanceOf(uint256 _projectId, uint256 _reservedRate) view returns(uint256)
func (_Controller *ControllerCallerSession) ReservedTokenBalanceOf(_projectId *big.Int, _reservedRate *big.Int) (*big.Int, error) {
	return _Controller.Contract.ReservedTokenBalanceOf(&_Controller.CallOpts, _projectId, _reservedRate)
}

// SplitsStore is a free data retrieval call binding the contract method 0x2bdfe004.
//
// Solidity: function splitsStore() view returns(address)
func (_Controller *ControllerCaller) SplitsStore(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "splitsStore")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SplitsStore is a free data retrieval call binding the contract method 0x2bdfe004.
//
// Solidity: function splitsStore() view returns(address)
func (_Controller *ControllerSession) SplitsStore() (common.Address, error) {
	return _Controller.Contract.SplitsStore(&_Controller.CallOpts)
}

// SplitsStore is a free data retrieval call binding the contract method 0x2bdfe004.
//
// Solidity: function splitsStore() view returns(address)
func (_Controller *ControllerCallerSession) SplitsStore() (common.Address, error) {
	return _Controller.Contract.SplitsStore(&_Controller.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (_Controller *ControllerCaller) SupportsInterface(opts *bind.CallOpts, _interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "supportsInterface", _interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (_Controller *ControllerSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _Controller.Contract.SupportsInterface(&_Controller.CallOpts, _interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (_Controller *ControllerCallerSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _Controller.Contract.SupportsInterface(&_Controller.CallOpts, _interfaceId)
}

// TokenStore is a free data retrieval call binding the contract method 0x61930630.
//
// Solidity: function tokenStore() view returns(address)
func (_Controller *ControllerCaller) TokenStore(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "tokenStore")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenStore is a free data retrieval call binding the contract method 0x61930630.
//
// Solidity: function tokenStore() view returns(address)
func (_Controller *ControllerSession) TokenStore() (common.Address, error) {
	return _Controller.Contract.TokenStore(&_Controller.CallOpts)
}

// TokenStore is a free data retrieval call binding the contract method 0x61930630.
//
// Solidity: function tokenStore() view returns(address)
func (_Controller *ControllerCallerSession) TokenStore() (common.Address, error) {
	return _Controller.Contract.TokenStore(&_Controller.CallOpts)
}

// TotalOutstandingTokensOf is a free data retrieval call binding the contract method 0xb5f1e34d.
//
// Solidity: function totalOutstandingTokensOf(uint256 _projectId, uint256 _reservedRate) view returns(uint256)
func (_Controller *ControllerCaller) TotalOutstandingTokensOf(opts *bind.CallOpts, _projectId *big.Int, _reservedRate *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "totalOutstandingTokensOf", _projectId, _reservedRate)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalOutstandingTokensOf is a free data retrieval call binding the contract method 0xb5f1e34d.
//
// Solidity: function totalOutstandingTokensOf(uint256 _projectId, uint256 _reservedRate) view returns(uint256)
func (_Controller *ControllerSession) TotalOutstandingTokensOf(_projectId *big.Int, _reservedRate *big.Int) (*big.Int, error) {
	return _Controller.Contract.TotalOutstandingTokensOf(&_Controller.CallOpts, _projectId, _reservedRate)
}

// TotalOutstandingTokensOf is a free data retrieval call binding the contract method 0xb5f1e34d.
//
// Solidity: function totalOutstandingTokensOf(uint256 _projectId, uint256 _reservedRate) view returns(uint256)
func (_Controller *ControllerCallerSession) TotalOutstandingTokensOf(_projectId *big.Int, _reservedRate *big.Int) (*big.Int, error) {
	return _Controller.Contract.TotalOutstandingTokensOf(&_Controller.CallOpts, _projectId, _reservedRate)
}

// BurnTokensOf is a paid mutator transaction binding the contract method 0x1665bc0f.
//
// Solidity: function burnTokensOf(address _holder, uint256 _projectId, uint256 _tokenCount, string _memo, bool _preferClaimedTokens) returns()
func (_Controller *ControllerTransactor) BurnTokensOf(opts *bind.TransactOpts, _holder common.Address, _projectId *big.Int, _tokenCount *big.Int, _memo string, _preferClaimedTokens bool) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "burnTokensOf", _holder, _projectId, _tokenCount, _memo, _preferClaimedTokens)
}

// BurnTokensOf is a paid mutator transaction binding the contract method 0x1665bc0f.
//
// Solidity: function burnTokensOf(address _holder, uint256 _projectId, uint256 _tokenCount, string _memo, bool _preferClaimedTokens) returns()
func (_Controller *ControllerSession) BurnTokensOf(_holder common.Address, _projectId *big.Int, _tokenCount *big.Int, _memo string, _preferClaimedTokens bool) (*types.Transaction, error) {
	return _Controller.Contract.BurnTokensOf(&_Controller.TransactOpts, _holder, _projectId, _tokenCount, _memo, _preferClaimedTokens)
}

// BurnTokensOf is a paid mutator transaction binding the contract method 0x1665bc0f.
//
// Solidity: function burnTokensOf(address _holder, uint256 _projectId, uint256 _tokenCount, string _memo, bool _preferClaimedTokens) returns()
func (_Controller *ControllerTransactorSession) BurnTokensOf(_holder common.Address, _projectId *big.Int, _tokenCount *big.Int, _memo string, _preferClaimedTokens bool) (*types.Transaction, error) {
	return _Controller.Contract.BurnTokensOf(&_Controller.TransactOpts, _holder, _projectId, _tokenCount, _memo, _preferClaimedTokens)
}

// DistributeReservedTokensOf is a paid mutator transaction binding the contract method 0xe867c59c.
//
// Solidity: function distributeReservedTokensOf(uint256 _projectId, string _memo) returns(uint256)
func (_Controller *ControllerTransactor) DistributeReservedTokensOf(opts *bind.TransactOpts, _projectId *big.Int, _memo string) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "distributeReservedTokensOf", _projectId, _memo)
}

// DistributeReservedTokensOf is a paid mutator transaction binding the contract method 0xe867c59c.
//
// Solidity: function distributeReservedTokensOf(uint256 _projectId, string _memo) returns(uint256)
func (_Controller *ControllerSession) DistributeReservedTokensOf(_projectId *big.Int, _memo string) (*types.Transaction, error) {
	return _Controller.Contract.DistributeReservedTokensOf(&_Controller.TransactOpts, _projectId, _memo)
}

// DistributeReservedTokensOf is a paid mutator transaction binding the contract method 0xe867c59c.
//
// Solidity: function distributeReservedTokensOf(uint256 _projectId, string _memo) returns(uint256)
func (_Controller *ControllerTransactorSession) DistributeReservedTokensOf(_projectId *big.Int, _memo string) (*types.Transaction, error) {
	return _Controller.Contract.DistributeReservedTokensOf(&_Controller.TransactOpts, _projectId, _memo)
}

// LaunchFundingCyclesFor is a paid mutator transaction binding the contract method 0x111d861b.
//
// Solidity: function launchFundingCyclesFor(uint256 _projectId, (uint256,uint256,uint256,address) _data, ((bool,bool,bool),uint256,uint256,uint256,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,address,uint256) _metadata, uint256 _mustStartAtOrAfter, (uint256,(bool,bool,uint256,uint256,address,uint256,address)[])[] _groupedSplits, (address,address,uint256,uint256,uint256,uint256)[] _fundAccessConstraints, address[] _terminals, string _memo) returns(uint256 configuration)
func (_Controller *ControllerTransactor) LaunchFundingCyclesFor(opts *bind.TransactOpts, _projectId *big.Int, _data JBFundingCycleData, _metadata JBFundingCycleMetadata, _mustStartAtOrAfter *big.Int, _groupedSplits []JBGroupedSplits, _fundAccessConstraints []JBFundAccessConstraints, _terminals []common.Address, _memo string) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "launchFundingCyclesFor", _projectId, _data, _metadata, _mustStartAtOrAfter, _groupedSplits, _fundAccessConstraints, _terminals, _memo)
}

// LaunchFundingCyclesFor is a paid mutator transaction binding the contract method 0x111d861b.
//
// Solidity: function launchFundingCyclesFor(uint256 _projectId, (uint256,uint256,uint256,address) _data, ((bool,bool,bool),uint256,uint256,uint256,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,address,uint256) _metadata, uint256 _mustStartAtOrAfter, (uint256,(bool,bool,uint256,uint256,address,uint256,address)[])[] _groupedSplits, (address,address,uint256,uint256,uint256,uint256)[] _fundAccessConstraints, address[] _terminals, string _memo) returns(uint256 configuration)
func (_Controller *ControllerSession) LaunchFundingCyclesFor(_projectId *big.Int, _data JBFundingCycleData, _metadata JBFundingCycleMetadata, _mustStartAtOrAfter *big.Int, _groupedSplits []JBGroupedSplits, _fundAccessConstraints []JBFundAccessConstraints, _terminals []common.Address, _memo string) (*types.Transaction, error) {
	return _Controller.Contract.LaunchFundingCyclesFor(&_Controller.TransactOpts, _projectId, _data, _metadata, _mustStartAtOrAfter, _groupedSplits, _fundAccessConstraints, _terminals, _memo)
}

// LaunchFundingCyclesFor is a paid mutator transaction binding the contract method 0x111d861b.
//
// Solidity: function launchFundingCyclesFor(uint256 _projectId, (uint256,uint256,uint256,address) _data, ((bool,bool,bool),uint256,uint256,uint256,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,address,uint256) _metadata, uint256 _mustStartAtOrAfter, (uint256,(bool,bool,uint256,uint256,address,uint256,address)[])[] _groupedSplits, (address,address,uint256,uint256,uint256,uint256)[] _fundAccessConstraints, address[] _terminals, string _memo) returns(uint256 configuration)
func (_Controller *ControllerTransactorSession) LaunchFundingCyclesFor(_projectId *big.Int, _data JBFundingCycleData, _metadata JBFundingCycleMetadata, _mustStartAtOrAfter *big.Int, _groupedSplits []JBGroupedSplits, _fundAccessConstraints []JBFundAccessConstraints, _terminals []common.Address, _memo string) (*types.Transaction, error) {
	return _Controller.Contract.LaunchFundingCyclesFor(&_Controller.TransactOpts, _projectId, _data, _metadata, _mustStartAtOrAfter, _groupedSplits, _fundAccessConstraints, _terminals, _memo)
}

// LaunchProjectFor is a paid mutator transaction binding the contract method 0xb3c52673.
//
// Solidity: function launchProjectFor(address _owner, (string,uint256) _projectMetadata, (uint256,uint256,uint256,address) _data, ((bool,bool,bool),uint256,uint256,uint256,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,address,uint256) _metadata, uint256 _mustStartAtOrAfter, (uint256,(bool,bool,uint256,uint256,address,uint256,address)[])[] _groupedSplits, (address,address,uint256,uint256,uint256,uint256)[] _fundAccessConstraints, address[] _terminals, string _memo) returns(uint256 projectId)
func (_Controller *ControllerTransactor) LaunchProjectFor(opts *bind.TransactOpts, _owner common.Address, _projectMetadata JBProjectMetadata, _data JBFundingCycleData, _metadata JBFundingCycleMetadata, _mustStartAtOrAfter *big.Int, _groupedSplits []JBGroupedSplits, _fundAccessConstraints []JBFundAccessConstraints, _terminals []common.Address, _memo string) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "launchProjectFor", _owner, _projectMetadata, _data, _metadata, _mustStartAtOrAfter, _groupedSplits, _fundAccessConstraints, _terminals, _memo)
}

// LaunchProjectFor is a paid mutator transaction binding the contract method 0xb3c52673.
//
// Solidity: function launchProjectFor(address _owner, (string,uint256) _projectMetadata, (uint256,uint256,uint256,address) _data, ((bool,bool,bool),uint256,uint256,uint256,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,address,uint256) _metadata, uint256 _mustStartAtOrAfter, (uint256,(bool,bool,uint256,uint256,address,uint256,address)[])[] _groupedSplits, (address,address,uint256,uint256,uint256,uint256)[] _fundAccessConstraints, address[] _terminals, string _memo) returns(uint256 projectId)
func (_Controller *ControllerSession) LaunchProjectFor(_owner common.Address, _projectMetadata JBProjectMetadata, _data JBFundingCycleData, _metadata JBFundingCycleMetadata, _mustStartAtOrAfter *big.Int, _groupedSplits []JBGroupedSplits, _fundAccessConstraints []JBFundAccessConstraints, _terminals []common.Address, _memo string) (*types.Transaction, error) {
	return _Controller.Contract.LaunchProjectFor(&_Controller.TransactOpts, _owner, _projectMetadata, _data, _metadata, _mustStartAtOrAfter, _groupedSplits, _fundAccessConstraints, _terminals, _memo)
}

// LaunchProjectFor is a paid mutator transaction binding the contract method 0xb3c52673.
//
// Solidity: function launchProjectFor(address _owner, (string,uint256) _projectMetadata, (uint256,uint256,uint256,address) _data, ((bool,bool,bool),uint256,uint256,uint256,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,address,uint256) _metadata, uint256 _mustStartAtOrAfter, (uint256,(bool,bool,uint256,uint256,address,uint256,address)[])[] _groupedSplits, (address,address,uint256,uint256,uint256,uint256)[] _fundAccessConstraints, address[] _terminals, string _memo) returns(uint256 projectId)
func (_Controller *ControllerTransactorSession) LaunchProjectFor(_owner common.Address, _projectMetadata JBProjectMetadata, _data JBFundingCycleData, _metadata JBFundingCycleMetadata, _mustStartAtOrAfter *big.Int, _groupedSplits []JBGroupedSplits, _fundAccessConstraints []JBFundAccessConstraints, _terminals []common.Address, _memo string) (*types.Transaction, error) {
	return _Controller.Contract.LaunchProjectFor(&_Controller.TransactOpts, _owner, _projectMetadata, _data, _metadata, _mustStartAtOrAfter, _groupedSplits, _fundAccessConstraints, _terminals, _memo)
}

// Migrate is a paid mutator transaction binding the contract method 0x405b84fa.
//
// Solidity: function migrate(uint256 _projectId, address _to) returns()
func (_Controller *ControllerTransactor) Migrate(opts *bind.TransactOpts, _projectId *big.Int, _to common.Address) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "migrate", _projectId, _to)
}

// Migrate is a paid mutator transaction binding the contract method 0x405b84fa.
//
// Solidity: function migrate(uint256 _projectId, address _to) returns()
func (_Controller *ControllerSession) Migrate(_projectId *big.Int, _to common.Address) (*types.Transaction, error) {
	return _Controller.Contract.Migrate(&_Controller.TransactOpts, _projectId, _to)
}

// Migrate is a paid mutator transaction binding the contract method 0x405b84fa.
//
// Solidity: function migrate(uint256 _projectId, address _to) returns()
func (_Controller *ControllerTransactorSession) Migrate(_projectId *big.Int, _to common.Address) (*types.Transaction, error) {
	return _Controller.Contract.Migrate(&_Controller.TransactOpts, _projectId, _to)
}

// MintTokensOf is a paid mutator transaction binding the contract method 0x8ae9c07b.
//
// Solidity: function mintTokensOf(uint256 _projectId, uint256 _tokenCount, address _beneficiary, string _memo, bool _preferClaimedTokens, bool _useReservedRate) returns(uint256 beneficiaryTokenCount)
func (_Controller *ControllerTransactor) MintTokensOf(opts *bind.TransactOpts, _projectId *big.Int, _tokenCount *big.Int, _beneficiary common.Address, _memo string, _preferClaimedTokens bool, _useReservedRate bool) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "mintTokensOf", _projectId, _tokenCount, _beneficiary, _memo, _preferClaimedTokens, _useReservedRate)
}

// MintTokensOf is a paid mutator transaction binding the contract method 0x8ae9c07b.
//
// Solidity: function mintTokensOf(uint256 _projectId, uint256 _tokenCount, address _beneficiary, string _memo, bool _preferClaimedTokens, bool _useReservedRate) returns(uint256 beneficiaryTokenCount)
func (_Controller *ControllerSession) MintTokensOf(_projectId *big.Int, _tokenCount *big.Int, _beneficiary common.Address, _memo string, _preferClaimedTokens bool, _useReservedRate bool) (*types.Transaction, error) {
	return _Controller.Contract.MintTokensOf(&_Controller.TransactOpts, _projectId, _tokenCount, _beneficiary, _memo, _preferClaimedTokens, _useReservedRate)
}

// MintTokensOf is a paid mutator transaction binding the contract method 0x8ae9c07b.
//
// Solidity: function mintTokensOf(uint256 _projectId, uint256 _tokenCount, address _beneficiary, string _memo, bool _preferClaimedTokens, bool _useReservedRate) returns(uint256 beneficiaryTokenCount)
func (_Controller *ControllerTransactorSession) MintTokensOf(_projectId *big.Int, _tokenCount *big.Int, _beneficiary common.Address, _memo string, _preferClaimedTokens bool, _useReservedRate bool) (*types.Transaction, error) {
	return _Controller.Contract.MintTokensOf(&_Controller.TransactOpts, _projectId, _tokenCount, _beneficiary, _memo, _preferClaimedTokens, _useReservedRate)
}

// PrepForMigrationOf is a paid mutator transaction binding the contract method 0x3e8c615b.
//
// Solidity: function prepForMigrationOf(uint256 _projectId, address _from) returns()
func (_Controller *ControllerTransactor) PrepForMigrationOf(opts *bind.TransactOpts, _projectId *big.Int, _from common.Address) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "prepForMigrationOf", _projectId, _from)
}

// PrepForMigrationOf is a paid mutator transaction binding the contract method 0x3e8c615b.
//
// Solidity: function prepForMigrationOf(uint256 _projectId, address _from) returns()
func (_Controller *ControllerSession) PrepForMigrationOf(_projectId *big.Int, _from common.Address) (*types.Transaction, error) {
	return _Controller.Contract.PrepForMigrationOf(&_Controller.TransactOpts, _projectId, _from)
}

// PrepForMigrationOf is a paid mutator transaction binding the contract method 0x3e8c615b.
//
// Solidity: function prepForMigrationOf(uint256 _projectId, address _from) returns()
func (_Controller *ControllerTransactorSession) PrepForMigrationOf(_projectId *big.Int, _from common.Address) (*types.Transaction, error) {
	return _Controller.Contract.PrepForMigrationOf(&_Controller.TransactOpts, _projectId, _from)
}

// ReconfigureFundingCyclesOf is a paid mutator transaction binding the contract method 0xbadbdf28.
//
// Solidity: function reconfigureFundingCyclesOf(uint256 _projectId, (uint256,uint256,uint256,address) _data, ((bool,bool,bool),uint256,uint256,uint256,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,address,uint256) _metadata, uint256 _mustStartAtOrAfter, (uint256,(bool,bool,uint256,uint256,address,uint256,address)[])[] _groupedSplits, (address,address,uint256,uint256,uint256,uint256)[] _fundAccessConstraints, string _memo) returns(uint256 configuration)
func (_Controller *ControllerTransactor) ReconfigureFundingCyclesOf(opts *bind.TransactOpts, _projectId *big.Int, _data JBFundingCycleData, _metadata JBFundingCycleMetadata, _mustStartAtOrAfter *big.Int, _groupedSplits []JBGroupedSplits, _fundAccessConstraints []JBFundAccessConstraints, _memo string) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "reconfigureFundingCyclesOf", _projectId, _data, _metadata, _mustStartAtOrAfter, _groupedSplits, _fundAccessConstraints, _memo)
}

// ReconfigureFundingCyclesOf is a paid mutator transaction binding the contract method 0xbadbdf28.
//
// Solidity: function reconfigureFundingCyclesOf(uint256 _projectId, (uint256,uint256,uint256,address) _data, ((bool,bool,bool),uint256,uint256,uint256,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,address,uint256) _metadata, uint256 _mustStartAtOrAfter, (uint256,(bool,bool,uint256,uint256,address,uint256,address)[])[] _groupedSplits, (address,address,uint256,uint256,uint256,uint256)[] _fundAccessConstraints, string _memo) returns(uint256 configuration)
func (_Controller *ControllerSession) ReconfigureFundingCyclesOf(_projectId *big.Int, _data JBFundingCycleData, _metadata JBFundingCycleMetadata, _mustStartAtOrAfter *big.Int, _groupedSplits []JBGroupedSplits, _fundAccessConstraints []JBFundAccessConstraints, _memo string) (*types.Transaction, error) {
	return _Controller.Contract.ReconfigureFundingCyclesOf(&_Controller.TransactOpts, _projectId, _data, _metadata, _mustStartAtOrAfter, _groupedSplits, _fundAccessConstraints, _memo)
}

// ReconfigureFundingCyclesOf is a paid mutator transaction binding the contract method 0xbadbdf28.
//
// Solidity: function reconfigureFundingCyclesOf(uint256 _projectId, (uint256,uint256,uint256,address) _data, ((bool,bool,bool),uint256,uint256,uint256,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,address,uint256) _metadata, uint256 _mustStartAtOrAfter, (uint256,(bool,bool,uint256,uint256,address,uint256,address)[])[] _groupedSplits, (address,address,uint256,uint256,uint256,uint256)[] _fundAccessConstraints, string _memo) returns(uint256 configuration)
func (_Controller *ControllerTransactorSession) ReconfigureFundingCyclesOf(_projectId *big.Int, _data JBFundingCycleData, _metadata JBFundingCycleMetadata, _mustStartAtOrAfter *big.Int, _groupedSplits []JBGroupedSplits, _fundAccessConstraints []JBFundAccessConstraints, _memo string) (*types.Transaction, error) {
	return _Controller.Contract.ReconfigureFundingCyclesOf(&_Controller.TransactOpts, _projectId, _data, _metadata, _mustStartAtOrAfter, _groupedSplits, _fundAccessConstraints, _memo)
}

// ControllerBurnTokensIterator is returned from FilterBurnTokens and is used to iterate over the raw logs and unpacked data for BurnTokens events raised by the Controller contract.
type ControllerBurnTokensIterator struct {
	Event *ControllerBurnTokens // Event containing the contract specifics and raw log

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
func (it *ControllerBurnTokensIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerBurnTokens)
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
		it.Event = new(ControllerBurnTokens)
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
func (it *ControllerBurnTokensIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerBurnTokensIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerBurnTokens represents a BurnTokens event raised by the Controller contract.
type ControllerBurnTokens struct {
	Holder     common.Address
	ProjectId  *big.Int
	TokenCount *big.Int
	Memo       string
	Caller     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBurnTokens is a free log retrieval operation binding the contract event 0xdf04e13ee4fcd48a81ab2fd114757093740a3efa9b6475d86e05878b4c59d079.
//
// Solidity: event BurnTokens(address indexed holder, uint256 indexed projectId, uint256 tokenCount, string memo, address caller)
func (_Controller *ControllerFilterer) FilterBurnTokens(opts *bind.FilterOpts, holder []common.Address, projectId []*big.Int) (*ControllerBurnTokensIterator, error) {

	var holderRule []interface{}
	for _, holderItem := range holder {
		holderRule = append(holderRule, holderItem)
	}
	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}

	logs, sub, err := _Controller.contract.FilterLogs(opts, "BurnTokens", holderRule, projectIdRule)
	if err != nil {
		return nil, err
	}
	return &ControllerBurnTokensIterator{contract: _Controller.contract, event: "BurnTokens", logs: logs, sub: sub}, nil
}

// WatchBurnTokens is a free log subscription operation binding the contract event 0xdf04e13ee4fcd48a81ab2fd114757093740a3efa9b6475d86e05878b4c59d079.
//
// Solidity: event BurnTokens(address indexed holder, uint256 indexed projectId, uint256 tokenCount, string memo, address caller)
func (_Controller *ControllerFilterer) WatchBurnTokens(opts *bind.WatchOpts, sink chan<- *ControllerBurnTokens, holder []common.Address, projectId []*big.Int) (event.Subscription, error) {

	var holderRule []interface{}
	for _, holderItem := range holder {
		holderRule = append(holderRule, holderItem)
	}
	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}

	logs, sub, err := _Controller.contract.WatchLogs(opts, "BurnTokens", holderRule, projectIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerBurnTokens)
				if err := _Controller.contract.UnpackLog(event, "BurnTokens", log); err != nil {
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

// ParseBurnTokens is a log parse operation binding the contract event 0xdf04e13ee4fcd48a81ab2fd114757093740a3efa9b6475d86e05878b4c59d079.
//
// Solidity: event BurnTokens(address indexed holder, uint256 indexed projectId, uint256 tokenCount, string memo, address caller)
func (_Controller *ControllerFilterer) ParseBurnTokens(log types.Log) (*ControllerBurnTokens, error) {
	event := new(ControllerBurnTokens)
	if err := _Controller.contract.UnpackLog(event, "BurnTokens", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ControllerDistributeReservedTokensIterator is returned from FilterDistributeReservedTokens and is used to iterate over the raw logs and unpacked data for DistributeReservedTokens events raised by the Controller contract.
type ControllerDistributeReservedTokensIterator struct {
	Event *ControllerDistributeReservedTokens // Event containing the contract specifics and raw log

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
func (it *ControllerDistributeReservedTokensIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerDistributeReservedTokens)
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
		it.Event = new(ControllerDistributeReservedTokens)
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
func (it *ControllerDistributeReservedTokensIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerDistributeReservedTokensIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerDistributeReservedTokens represents a DistributeReservedTokens event raised by the Controller contract.
type ControllerDistributeReservedTokens struct {
	FundingCycleConfiguration *big.Int
	FundingCycleNumber        *big.Int
	ProjectId                 *big.Int
	Beneficiary               common.Address
	TokenCount                *big.Int
	BeneficiaryTokenCount     *big.Int
	Memo                      string
	Caller                    common.Address
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterDistributeReservedTokens is a free log retrieval operation binding the contract event 0xb12d7a78048433f69fe6d30145bf08aad8e82985b96e4db6d5c6a7e94d57086e.
//
// Solidity: event DistributeReservedTokens(uint256 indexed fundingCycleConfiguration, uint256 indexed fundingCycleNumber, uint256 indexed projectId, address beneficiary, uint256 tokenCount, uint256 beneficiaryTokenCount, string memo, address caller)
func (_Controller *ControllerFilterer) FilterDistributeReservedTokens(opts *bind.FilterOpts, fundingCycleConfiguration []*big.Int, fundingCycleNumber []*big.Int, projectId []*big.Int) (*ControllerDistributeReservedTokensIterator, error) {

	var fundingCycleConfigurationRule []interface{}
	for _, fundingCycleConfigurationItem := range fundingCycleConfiguration {
		fundingCycleConfigurationRule = append(fundingCycleConfigurationRule, fundingCycleConfigurationItem)
	}
	var fundingCycleNumberRule []interface{}
	for _, fundingCycleNumberItem := range fundingCycleNumber {
		fundingCycleNumberRule = append(fundingCycleNumberRule, fundingCycleNumberItem)
	}
	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}

	logs, sub, err := _Controller.contract.FilterLogs(opts, "DistributeReservedTokens", fundingCycleConfigurationRule, fundingCycleNumberRule, projectIdRule)
	if err != nil {
		return nil, err
	}
	return &ControllerDistributeReservedTokensIterator{contract: _Controller.contract, event: "DistributeReservedTokens", logs: logs, sub: sub}, nil
}

// WatchDistributeReservedTokens is a free log subscription operation binding the contract event 0xb12d7a78048433f69fe6d30145bf08aad8e82985b96e4db6d5c6a7e94d57086e.
//
// Solidity: event DistributeReservedTokens(uint256 indexed fundingCycleConfiguration, uint256 indexed fundingCycleNumber, uint256 indexed projectId, address beneficiary, uint256 tokenCount, uint256 beneficiaryTokenCount, string memo, address caller)
func (_Controller *ControllerFilterer) WatchDistributeReservedTokens(opts *bind.WatchOpts, sink chan<- *ControllerDistributeReservedTokens, fundingCycleConfiguration []*big.Int, fundingCycleNumber []*big.Int, projectId []*big.Int) (event.Subscription, error) {

	var fundingCycleConfigurationRule []interface{}
	for _, fundingCycleConfigurationItem := range fundingCycleConfiguration {
		fundingCycleConfigurationRule = append(fundingCycleConfigurationRule, fundingCycleConfigurationItem)
	}
	var fundingCycleNumberRule []interface{}
	for _, fundingCycleNumberItem := range fundingCycleNumber {
		fundingCycleNumberRule = append(fundingCycleNumberRule, fundingCycleNumberItem)
	}
	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}

	logs, sub, err := _Controller.contract.WatchLogs(opts, "DistributeReservedTokens", fundingCycleConfigurationRule, fundingCycleNumberRule, projectIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerDistributeReservedTokens)
				if err := _Controller.contract.UnpackLog(event, "DistributeReservedTokens", log); err != nil {
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

// ParseDistributeReservedTokens is a log parse operation binding the contract event 0xb12d7a78048433f69fe6d30145bf08aad8e82985b96e4db6d5c6a7e94d57086e.
//
// Solidity: event DistributeReservedTokens(uint256 indexed fundingCycleConfiguration, uint256 indexed fundingCycleNumber, uint256 indexed projectId, address beneficiary, uint256 tokenCount, uint256 beneficiaryTokenCount, string memo, address caller)
func (_Controller *ControllerFilterer) ParseDistributeReservedTokens(log types.Log) (*ControllerDistributeReservedTokens, error) {
	event := new(ControllerDistributeReservedTokens)
	if err := _Controller.contract.UnpackLog(event, "DistributeReservedTokens", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ControllerDistributeToReservedTokenSplitIterator is returned from FilterDistributeToReservedTokenSplit and is used to iterate over the raw logs and unpacked data for DistributeToReservedTokenSplit events raised by the Controller contract.
type ControllerDistributeToReservedTokenSplitIterator struct {
	Event *ControllerDistributeToReservedTokenSplit // Event containing the contract specifics and raw log

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
func (it *ControllerDistributeToReservedTokenSplitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerDistributeToReservedTokenSplit)
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
		it.Event = new(ControllerDistributeToReservedTokenSplit)
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
func (it *ControllerDistributeToReservedTokenSplitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerDistributeToReservedTokenSplitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerDistributeToReservedTokenSplit represents a DistributeToReservedTokenSplit event raised by the Controller contract.
type ControllerDistributeToReservedTokenSplit struct {
	ProjectId  *big.Int
	Domain     *big.Int
	Group      *big.Int
	Split      JBSplit
	TokenCount *big.Int
	Caller     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDistributeToReservedTokenSplit is a free log retrieval operation binding the contract event 0x990c1da5d280602f972ae3c429b6dc66a66569875205cc6e67f14e0b3fc4d2ac.
//
// Solidity: event DistributeToReservedTokenSplit(uint256 indexed projectId, uint256 indexed domain, uint256 indexed group, (bool,bool,uint256,uint256,address,uint256,address) split, uint256 tokenCount, address caller)
func (_Controller *ControllerFilterer) FilterDistributeToReservedTokenSplit(opts *bind.FilterOpts, projectId []*big.Int, domain []*big.Int, group []*big.Int) (*ControllerDistributeToReservedTokenSplitIterator, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var groupRule []interface{}
	for _, groupItem := range group {
		groupRule = append(groupRule, groupItem)
	}

	logs, sub, err := _Controller.contract.FilterLogs(opts, "DistributeToReservedTokenSplit", projectIdRule, domainRule, groupRule)
	if err != nil {
		return nil, err
	}
	return &ControllerDistributeToReservedTokenSplitIterator{contract: _Controller.contract, event: "DistributeToReservedTokenSplit", logs: logs, sub: sub}, nil
}

// WatchDistributeToReservedTokenSplit is a free log subscription operation binding the contract event 0x990c1da5d280602f972ae3c429b6dc66a66569875205cc6e67f14e0b3fc4d2ac.
//
// Solidity: event DistributeToReservedTokenSplit(uint256 indexed projectId, uint256 indexed domain, uint256 indexed group, (bool,bool,uint256,uint256,address,uint256,address) split, uint256 tokenCount, address caller)
func (_Controller *ControllerFilterer) WatchDistributeToReservedTokenSplit(opts *bind.WatchOpts, sink chan<- *ControllerDistributeToReservedTokenSplit, projectId []*big.Int, domain []*big.Int, group []*big.Int) (event.Subscription, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var groupRule []interface{}
	for _, groupItem := range group {
		groupRule = append(groupRule, groupItem)
	}

	logs, sub, err := _Controller.contract.WatchLogs(opts, "DistributeToReservedTokenSplit", projectIdRule, domainRule, groupRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerDistributeToReservedTokenSplit)
				if err := _Controller.contract.UnpackLog(event, "DistributeToReservedTokenSplit", log); err != nil {
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

// ParseDistributeToReservedTokenSplit is a log parse operation binding the contract event 0x990c1da5d280602f972ae3c429b6dc66a66569875205cc6e67f14e0b3fc4d2ac.
//
// Solidity: event DistributeToReservedTokenSplit(uint256 indexed projectId, uint256 indexed domain, uint256 indexed group, (bool,bool,uint256,uint256,address,uint256,address) split, uint256 tokenCount, address caller)
func (_Controller *ControllerFilterer) ParseDistributeToReservedTokenSplit(log types.Log) (*ControllerDistributeToReservedTokenSplit, error) {
	event := new(ControllerDistributeToReservedTokenSplit)
	if err := _Controller.contract.UnpackLog(event, "DistributeToReservedTokenSplit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ControllerLaunchFundingCyclesIterator is returned from FilterLaunchFundingCycles and is used to iterate over the raw logs and unpacked data for LaunchFundingCycles events raised by the Controller contract.
type ControllerLaunchFundingCyclesIterator struct {
	Event *ControllerLaunchFundingCycles // Event containing the contract specifics and raw log

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
func (it *ControllerLaunchFundingCyclesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerLaunchFundingCycles)
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
		it.Event = new(ControllerLaunchFundingCycles)
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
func (it *ControllerLaunchFundingCyclesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerLaunchFundingCyclesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerLaunchFundingCycles represents a LaunchFundingCycles event raised by the Controller contract.
type ControllerLaunchFundingCycles struct {
	Configuration *big.Int
	ProjectId     *big.Int
	Memo          string
	Caller        common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterLaunchFundingCycles is a free log retrieval operation binding the contract event 0x8621c3692c18d8ddd4cd26457d6a3e1d916e7001c4243a6c8dfa862a7e10a823.
//
// Solidity: event LaunchFundingCycles(uint256 configuration, uint256 projectId, string memo, address caller)
func (_Controller *ControllerFilterer) FilterLaunchFundingCycles(opts *bind.FilterOpts) (*ControllerLaunchFundingCyclesIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "LaunchFundingCycles")
	if err != nil {
		return nil, err
	}
	return &ControllerLaunchFundingCyclesIterator{contract: _Controller.contract, event: "LaunchFundingCycles", logs: logs, sub: sub}, nil
}

// WatchLaunchFundingCycles is a free log subscription operation binding the contract event 0x8621c3692c18d8ddd4cd26457d6a3e1d916e7001c4243a6c8dfa862a7e10a823.
//
// Solidity: event LaunchFundingCycles(uint256 configuration, uint256 projectId, string memo, address caller)
func (_Controller *ControllerFilterer) WatchLaunchFundingCycles(opts *bind.WatchOpts, sink chan<- *ControllerLaunchFundingCycles) (event.Subscription, error) {

	logs, sub, err := _Controller.contract.WatchLogs(opts, "LaunchFundingCycles")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerLaunchFundingCycles)
				if err := _Controller.contract.UnpackLog(event, "LaunchFundingCycles", log); err != nil {
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

// ParseLaunchFundingCycles is a log parse operation binding the contract event 0x8621c3692c18d8ddd4cd26457d6a3e1d916e7001c4243a6c8dfa862a7e10a823.
//
// Solidity: event LaunchFundingCycles(uint256 configuration, uint256 projectId, string memo, address caller)
func (_Controller *ControllerFilterer) ParseLaunchFundingCycles(log types.Log) (*ControllerLaunchFundingCycles, error) {
	event := new(ControllerLaunchFundingCycles)
	if err := _Controller.contract.UnpackLog(event, "LaunchFundingCycles", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ControllerLaunchProjectIterator is returned from FilterLaunchProject and is used to iterate over the raw logs and unpacked data for LaunchProject events raised by the Controller contract.
type ControllerLaunchProjectIterator struct {
	Event *ControllerLaunchProject // Event containing the contract specifics and raw log

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
func (it *ControllerLaunchProjectIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerLaunchProject)
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
		it.Event = new(ControllerLaunchProject)
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
func (it *ControllerLaunchProjectIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerLaunchProjectIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerLaunchProject represents a LaunchProject event raised by the Controller contract.
type ControllerLaunchProject struct {
	Configuration *big.Int
	ProjectId     *big.Int
	Memo          string
	Caller        common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterLaunchProject is a free log retrieval operation binding the contract event 0xf3e6948ba8b32d557363ea08470121c47c0127659aed09320812174d373afef2.
//
// Solidity: event LaunchProject(uint256 configuration, uint256 projectId, string memo, address caller)
func (_Controller *ControllerFilterer) FilterLaunchProject(opts *bind.FilterOpts) (*ControllerLaunchProjectIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "LaunchProject")
	if err != nil {
		return nil, err
	}
	return &ControllerLaunchProjectIterator{contract: _Controller.contract, event: "LaunchProject", logs: logs, sub: sub}, nil
}

// WatchLaunchProject is a free log subscription operation binding the contract event 0xf3e6948ba8b32d557363ea08470121c47c0127659aed09320812174d373afef2.
//
// Solidity: event LaunchProject(uint256 configuration, uint256 projectId, string memo, address caller)
func (_Controller *ControllerFilterer) WatchLaunchProject(opts *bind.WatchOpts, sink chan<- *ControllerLaunchProject) (event.Subscription, error) {

	logs, sub, err := _Controller.contract.WatchLogs(opts, "LaunchProject")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerLaunchProject)
				if err := _Controller.contract.UnpackLog(event, "LaunchProject", log); err != nil {
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

// ParseLaunchProject is a log parse operation binding the contract event 0xf3e6948ba8b32d557363ea08470121c47c0127659aed09320812174d373afef2.
//
// Solidity: event LaunchProject(uint256 configuration, uint256 projectId, string memo, address caller)
func (_Controller *ControllerFilterer) ParseLaunchProject(log types.Log) (*ControllerLaunchProject, error) {
	event := new(ControllerLaunchProject)
	if err := _Controller.contract.UnpackLog(event, "LaunchProject", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ControllerMigrateIterator is returned from FilterMigrate and is used to iterate over the raw logs and unpacked data for Migrate events raised by the Controller contract.
type ControllerMigrateIterator struct {
	Event *ControllerMigrate // Event containing the contract specifics and raw log

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
func (it *ControllerMigrateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerMigrate)
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
		it.Event = new(ControllerMigrate)
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
func (it *ControllerMigrateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerMigrateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerMigrate represents a Migrate event raised by the Controller contract.
type ControllerMigrate struct {
	ProjectId *big.Int
	To        common.Address
	Caller    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMigrate is a free log retrieval operation binding the contract event 0x01f954abace731a88ab86e71186040cc2be49fe517ea06bc0d24f25b82b83456.
//
// Solidity: event Migrate(uint256 indexed projectId, address to, address caller)
func (_Controller *ControllerFilterer) FilterMigrate(opts *bind.FilterOpts, projectId []*big.Int) (*ControllerMigrateIterator, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}

	logs, sub, err := _Controller.contract.FilterLogs(opts, "Migrate", projectIdRule)
	if err != nil {
		return nil, err
	}
	return &ControllerMigrateIterator{contract: _Controller.contract, event: "Migrate", logs: logs, sub: sub}, nil
}

// WatchMigrate is a free log subscription operation binding the contract event 0x01f954abace731a88ab86e71186040cc2be49fe517ea06bc0d24f25b82b83456.
//
// Solidity: event Migrate(uint256 indexed projectId, address to, address caller)
func (_Controller *ControllerFilterer) WatchMigrate(opts *bind.WatchOpts, sink chan<- *ControllerMigrate, projectId []*big.Int) (event.Subscription, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}

	logs, sub, err := _Controller.contract.WatchLogs(opts, "Migrate", projectIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerMigrate)
				if err := _Controller.contract.UnpackLog(event, "Migrate", log); err != nil {
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

// ParseMigrate is a log parse operation binding the contract event 0x01f954abace731a88ab86e71186040cc2be49fe517ea06bc0d24f25b82b83456.
//
// Solidity: event Migrate(uint256 indexed projectId, address to, address caller)
func (_Controller *ControllerFilterer) ParseMigrate(log types.Log) (*ControllerMigrate, error) {
	event := new(ControllerMigrate)
	if err := _Controller.contract.UnpackLog(event, "Migrate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ControllerMintTokensIterator is returned from FilterMintTokens and is used to iterate over the raw logs and unpacked data for MintTokens events raised by the Controller contract.
type ControllerMintTokensIterator struct {
	Event *ControllerMintTokens // Event containing the contract specifics and raw log

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
func (it *ControllerMintTokensIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerMintTokens)
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
		it.Event = new(ControllerMintTokens)
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
func (it *ControllerMintTokensIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerMintTokensIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerMintTokens represents a MintTokens event raised by the Controller contract.
type ControllerMintTokens struct {
	Beneficiary           common.Address
	ProjectId             *big.Int
	TokenCount            *big.Int
	BeneficiaryTokenCount *big.Int
	Memo                  string
	ReservedRate          *big.Int
	Caller                common.Address
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterMintTokens is a free log retrieval operation binding the contract event 0xe6fee9c572244c0c2238c3112ac12d411750a7ee00eeebd32521c3e5a666c14b.
//
// Solidity: event MintTokens(address indexed beneficiary, uint256 indexed projectId, uint256 tokenCount, uint256 beneficiaryTokenCount, string memo, uint256 reservedRate, address caller)
func (_Controller *ControllerFilterer) FilterMintTokens(opts *bind.FilterOpts, beneficiary []common.Address, projectId []*big.Int) (*ControllerMintTokensIterator, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}
	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}

	logs, sub, err := _Controller.contract.FilterLogs(opts, "MintTokens", beneficiaryRule, projectIdRule)
	if err != nil {
		return nil, err
	}
	return &ControllerMintTokensIterator{contract: _Controller.contract, event: "MintTokens", logs: logs, sub: sub}, nil
}

// WatchMintTokens is a free log subscription operation binding the contract event 0xe6fee9c572244c0c2238c3112ac12d411750a7ee00eeebd32521c3e5a666c14b.
//
// Solidity: event MintTokens(address indexed beneficiary, uint256 indexed projectId, uint256 tokenCount, uint256 beneficiaryTokenCount, string memo, uint256 reservedRate, address caller)
func (_Controller *ControllerFilterer) WatchMintTokens(opts *bind.WatchOpts, sink chan<- *ControllerMintTokens, beneficiary []common.Address, projectId []*big.Int) (event.Subscription, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}
	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}

	logs, sub, err := _Controller.contract.WatchLogs(opts, "MintTokens", beneficiaryRule, projectIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerMintTokens)
				if err := _Controller.contract.UnpackLog(event, "MintTokens", log); err != nil {
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

// ParseMintTokens is a log parse operation binding the contract event 0xe6fee9c572244c0c2238c3112ac12d411750a7ee00eeebd32521c3e5a666c14b.
//
// Solidity: event MintTokens(address indexed beneficiary, uint256 indexed projectId, uint256 tokenCount, uint256 beneficiaryTokenCount, string memo, uint256 reservedRate, address caller)
func (_Controller *ControllerFilterer) ParseMintTokens(log types.Log) (*ControllerMintTokens, error) {
	event := new(ControllerMintTokens)
	if err := _Controller.contract.UnpackLog(event, "MintTokens", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ControllerPrepMigrationIterator is returned from FilterPrepMigration and is used to iterate over the raw logs and unpacked data for PrepMigration events raised by the Controller contract.
type ControllerPrepMigrationIterator struct {
	Event *ControllerPrepMigration // Event containing the contract specifics and raw log

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
func (it *ControllerPrepMigrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerPrepMigration)
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
		it.Event = new(ControllerPrepMigration)
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
func (it *ControllerPrepMigrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerPrepMigrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerPrepMigration represents a PrepMigration event raised by the Controller contract.
type ControllerPrepMigration struct {
	ProjectId *big.Int
	From      common.Address
	Caller    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPrepMigration is a free log retrieval operation binding the contract event 0xf389f4f5d01fe4903d6a7a63b8790b7bf80d374b6afed808c03795c3b323d4d3.
//
// Solidity: event PrepMigration(uint256 indexed projectId, address from, address caller)
func (_Controller *ControllerFilterer) FilterPrepMigration(opts *bind.FilterOpts, projectId []*big.Int) (*ControllerPrepMigrationIterator, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}

	logs, sub, err := _Controller.contract.FilterLogs(opts, "PrepMigration", projectIdRule)
	if err != nil {
		return nil, err
	}
	return &ControllerPrepMigrationIterator{contract: _Controller.contract, event: "PrepMigration", logs: logs, sub: sub}, nil
}

// WatchPrepMigration is a free log subscription operation binding the contract event 0xf389f4f5d01fe4903d6a7a63b8790b7bf80d374b6afed808c03795c3b323d4d3.
//
// Solidity: event PrepMigration(uint256 indexed projectId, address from, address caller)
func (_Controller *ControllerFilterer) WatchPrepMigration(opts *bind.WatchOpts, sink chan<- *ControllerPrepMigration, projectId []*big.Int) (event.Subscription, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}

	logs, sub, err := _Controller.contract.WatchLogs(opts, "PrepMigration", projectIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerPrepMigration)
				if err := _Controller.contract.UnpackLog(event, "PrepMigration", log); err != nil {
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

// ParsePrepMigration is a log parse operation binding the contract event 0xf389f4f5d01fe4903d6a7a63b8790b7bf80d374b6afed808c03795c3b323d4d3.
//
// Solidity: event PrepMigration(uint256 indexed projectId, address from, address caller)
func (_Controller *ControllerFilterer) ParsePrepMigration(log types.Log) (*ControllerPrepMigration, error) {
	event := new(ControllerPrepMigration)
	if err := _Controller.contract.UnpackLog(event, "PrepMigration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ControllerReconfigureFundingCyclesIterator is returned from FilterReconfigureFundingCycles and is used to iterate over the raw logs and unpacked data for ReconfigureFundingCycles events raised by the Controller contract.
type ControllerReconfigureFundingCyclesIterator struct {
	Event *ControllerReconfigureFundingCycles // Event containing the contract specifics and raw log

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
func (it *ControllerReconfigureFundingCyclesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerReconfigureFundingCycles)
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
		it.Event = new(ControllerReconfigureFundingCycles)
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
func (it *ControllerReconfigureFundingCyclesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerReconfigureFundingCyclesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerReconfigureFundingCycles represents a ReconfigureFundingCycles event raised by the Controller contract.
type ControllerReconfigureFundingCycles struct {
	Configuration *big.Int
	ProjectId     *big.Int
	Memo          string
	Caller        common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterReconfigureFundingCycles is a free log retrieval operation binding the contract event 0x1a08ad37c93969a586fd0605d240c17113b20a9c9ac3847595d0ab425225f6af.
//
// Solidity: event ReconfigureFundingCycles(uint256 configuration, uint256 projectId, string memo, address caller)
func (_Controller *ControllerFilterer) FilterReconfigureFundingCycles(opts *bind.FilterOpts) (*ControllerReconfigureFundingCyclesIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "ReconfigureFundingCycles")
	if err != nil {
		return nil, err
	}
	return &ControllerReconfigureFundingCyclesIterator{contract: _Controller.contract, event: "ReconfigureFundingCycles", logs: logs, sub: sub}, nil
}

// WatchReconfigureFundingCycles is a free log subscription operation binding the contract event 0x1a08ad37c93969a586fd0605d240c17113b20a9c9ac3847595d0ab425225f6af.
//
// Solidity: event ReconfigureFundingCycles(uint256 configuration, uint256 projectId, string memo, address caller)
func (_Controller *ControllerFilterer) WatchReconfigureFundingCycles(opts *bind.WatchOpts, sink chan<- *ControllerReconfigureFundingCycles) (event.Subscription, error) {

	logs, sub, err := _Controller.contract.WatchLogs(opts, "ReconfigureFundingCycles")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerReconfigureFundingCycles)
				if err := _Controller.contract.UnpackLog(event, "ReconfigureFundingCycles", log); err != nil {
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

// ParseReconfigureFundingCycles is a log parse operation binding the contract event 0x1a08ad37c93969a586fd0605d240c17113b20a9c9ac3847595d0ab425225f6af.
//
// Solidity: event ReconfigureFundingCycles(uint256 configuration, uint256 projectId, string memo, address caller)
func (_Controller *ControllerFilterer) ParseReconfigureFundingCycles(log types.Log) (*ControllerReconfigureFundingCycles, error) {
	event := new(ControllerReconfigureFundingCycles)
	if err := _Controller.contract.UnpackLog(event, "ReconfigureFundingCycles", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ControllerSetFundAccessConstraintsIterator is returned from FilterSetFundAccessConstraints and is used to iterate over the raw logs and unpacked data for SetFundAccessConstraints events raised by the Controller contract.
type ControllerSetFundAccessConstraintsIterator struct {
	Event *ControllerSetFundAccessConstraints // Event containing the contract specifics and raw log

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
func (it *ControllerSetFundAccessConstraintsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerSetFundAccessConstraints)
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
		it.Event = new(ControllerSetFundAccessConstraints)
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
func (it *ControllerSetFundAccessConstraintsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerSetFundAccessConstraintsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerSetFundAccessConstraints represents a SetFundAccessConstraints event raised by the Controller contract.
type ControllerSetFundAccessConstraints struct {
	FundingCycleConfiguration *big.Int
	FundingCycleNumber        *big.Int
	ProjectId                 *big.Int
	Constraints               JBFundAccessConstraints
	Caller                    common.Address
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterSetFundAccessConstraints is a free log retrieval operation binding the contract event 0x8a1c80ab9ab05d6bf02096dd94aa2fa05a9de212e793835a50dd5ffe8d8cbab8.
//
// Solidity: event SetFundAccessConstraints(uint256 indexed fundingCycleConfiguration, uint256 indexed fundingCycleNumber, uint256 indexed projectId, (address,address,uint256,uint256,uint256,uint256) constraints, address caller)
func (_Controller *ControllerFilterer) FilterSetFundAccessConstraints(opts *bind.FilterOpts, fundingCycleConfiguration []*big.Int, fundingCycleNumber []*big.Int, projectId []*big.Int) (*ControllerSetFundAccessConstraintsIterator, error) {

	var fundingCycleConfigurationRule []interface{}
	for _, fundingCycleConfigurationItem := range fundingCycleConfiguration {
		fundingCycleConfigurationRule = append(fundingCycleConfigurationRule, fundingCycleConfigurationItem)
	}
	var fundingCycleNumberRule []interface{}
	for _, fundingCycleNumberItem := range fundingCycleNumber {
		fundingCycleNumberRule = append(fundingCycleNumberRule, fundingCycleNumberItem)
	}
	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}

	logs, sub, err := _Controller.contract.FilterLogs(opts, "SetFundAccessConstraints", fundingCycleConfigurationRule, fundingCycleNumberRule, projectIdRule)
	if err != nil {
		return nil, err
	}
	return &ControllerSetFundAccessConstraintsIterator{contract: _Controller.contract, event: "SetFundAccessConstraints", logs: logs, sub: sub}, nil
}

// WatchSetFundAccessConstraints is a free log subscription operation binding the contract event 0x8a1c80ab9ab05d6bf02096dd94aa2fa05a9de212e793835a50dd5ffe8d8cbab8.
//
// Solidity: event SetFundAccessConstraints(uint256 indexed fundingCycleConfiguration, uint256 indexed fundingCycleNumber, uint256 indexed projectId, (address,address,uint256,uint256,uint256,uint256) constraints, address caller)
func (_Controller *ControllerFilterer) WatchSetFundAccessConstraints(opts *bind.WatchOpts, sink chan<- *ControllerSetFundAccessConstraints, fundingCycleConfiguration []*big.Int, fundingCycleNumber []*big.Int, projectId []*big.Int) (event.Subscription, error) {

	var fundingCycleConfigurationRule []interface{}
	for _, fundingCycleConfigurationItem := range fundingCycleConfiguration {
		fundingCycleConfigurationRule = append(fundingCycleConfigurationRule, fundingCycleConfigurationItem)
	}
	var fundingCycleNumberRule []interface{}
	for _, fundingCycleNumberItem := range fundingCycleNumber {
		fundingCycleNumberRule = append(fundingCycleNumberRule, fundingCycleNumberItem)
	}
	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}

	logs, sub, err := _Controller.contract.WatchLogs(opts, "SetFundAccessConstraints", fundingCycleConfigurationRule, fundingCycleNumberRule, projectIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerSetFundAccessConstraints)
				if err := _Controller.contract.UnpackLog(event, "SetFundAccessConstraints", log); err != nil {
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

// ParseSetFundAccessConstraints is a log parse operation binding the contract event 0x8a1c80ab9ab05d6bf02096dd94aa2fa05a9de212e793835a50dd5ffe8d8cbab8.
//
// Solidity: event SetFundAccessConstraints(uint256 indexed fundingCycleConfiguration, uint256 indexed fundingCycleNumber, uint256 indexed projectId, (address,address,uint256,uint256,uint256,uint256) constraints, address caller)
func (_Controller *ControllerFilterer) ParseSetFundAccessConstraints(log types.Log) (*ControllerSetFundAccessConstraints, error) {
	event := new(ControllerSetFundAccessConstraints)
	if err := _Controller.contract.UnpackLog(event, "SetFundAccessConstraints", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
