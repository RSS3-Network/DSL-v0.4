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

// JBDidPayData is an auto generated low-level Go binding around an user-defined struct.
type JBDidPayData struct {
	Payer                            common.Address
	ProjectId                        *big.Int
	CurrentFundingCycleConfiguration *big.Int
	Amount                           JBTokenAmount
	ForwardedAmount                  JBTokenAmount
	ProjectTokenCount                *big.Int
	Beneficiary                      common.Address
	PreferClaimedTokens              bool
	Memo                             string
	Metadata                         []byte
}

// JBDidRedeemData is an auto generated low-level Go binding around an user-defined struct.
type JBDidRedeemData struct {
	Holder                           common.Address
	ProjectId                        *big.Int
	CurrentFundingCycleConfiguration *big.Int
	ProjectTokenCount                *big.Int
	ReclaimedAmount                  JBTokenAmount
	ForwardedAmount                  JBTokenAmount
	Beneficiary                      common.Address
	Memo                             string
	Metadata                         []byte
}

// JBFee is an auto generated low-level Go binding around an user-defined struct.
type JBFee struct {
	Amount      *big.Int
	Fee         uint32
	FeeDiscount uint32
	Beneficiary common.Address
}

// JBSplit is an auto generated low-level Go binding around an user-defined struct.
type JBSplit struct {
	PreferClaimed      bool
	PreferAddToBalance bool
	Percent            *big.Int
	ProjectId          *big.Int
	Beneficiary        common.Address
	LockedUntil        *big.Int
	Allocator          common.Address
}

// JBTokenAmount is an auto generated low-level Go binding around an user-defined struct.
type JBTokenAmount struct {
	Token    common.Address
	Value    *big.Int
	Decimals *big.Int
	Currency *big.Int
}

// ETHPaymentTerminalMetaData contains all meta data concerning the ETHPaymentTerminal contract.
var ETHPaymentTerminalMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_baseWeightCurrency\",\"type\":\"uint256\"},{\"internalType\":\"contractIJBOperatorStore\",\"name\":\"_operatorStore\",\"type\":\"address\"},{\"internalType\":\"contractIJBProjects\",\"name\":\"_projects\",\"type\":\"address\"},{\"internalType\":\"contractIJBDirectory\",\"name\":\"_directory\",\"type\":\"address\"},{\"internalType\":\"contractIJBSplitsStore\",\"name\":\"_splitsStore\",\"type\":\"address\"},{\"internalType\":\"contractIJBPrices\",\"name\":\"_prices\",\"type\":\"address\"},{\"internalType\":\"contractIJBSingleTokenPaymentTerminalStore\",\"name\":\"_store\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"FEE_TOO_HIGH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"INADEQUATE_DISTRIBUTION_AMOUNT\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"INADEQUATE_RECLAIM_AMOUNT\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"INADEQUATE_TOKEN_COUNT\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NO_MSG_VALUE_ALLOWED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PAY_TO_ZERO_ADDRESS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"prod1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"denominator\",\"type\":\"uint256\"}],\"name\":\"PRBMath__MulDivOverflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PROJECT_TERMINAL_MISMATCH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"REDEEM_TO_ZERO_ADDRESS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TERMINAL_IN_SPLIT_ZERO_ADDRESS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TERMINAL_TOKENS_INCOMPATIBLE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UNAUTHORIZED\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"projectId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"refundedFees\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"memo\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"AddToBalance\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIJBPayDelegate\",\"name\":\"delegate\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"projectId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentFundingCycleConfiguration\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currency\",\"type\":\"uint256\"}],\"internalType\":\"structJBTokenAmount\",\"name\":\"amount\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currency\",\"type\":\"uint256\"}],\"internalType\":\"structJBTokenAmount\",\"name\":\"forwardedAmount\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"projectTokenCount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"preferClaimedTokens\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"memo\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structJBDidPayData\",\"name\":\"data\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delegatedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"DelegateDidPay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIJBRedemptionDelegate\",\"name\":\"delegate\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"projectId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentFundingCycleConfiguration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"projectTokenCount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currency\",\"type\":\"uint256\"}],\"internalType\":\"structJBTokenAmount\",\"name\":\"reclaimedAmount\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currency\",\"type\":\"uint256\"}],\"internalType\":\"structJBTokenAmount\",\"name\":\"forwardedAmount\",\"type\":\"tuple\"},{\"internalType\":\"addresspayable\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"memo\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structJBDidRedeemData\",\"name\":\"data\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delegatedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"DelegateDidRedeem\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fundingCycleConfiguration\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fundingCycleNumber\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"projectId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"distributedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"beneficiaryDistributionAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"memo\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"DistributePayouts\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"projectId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"domain\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"group\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"preferClaimed\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"preferAddToBalance\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"percent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"projectId\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lockedUntil\",\"type\":\"uint256\"},{\"internalType\":\"contractIJBSplitAllocator\",\"name\":\"allocator\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structJBSplit\",\"name\":\"split\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"DistributeToPayoutSplit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"projectId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeDiscount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"HoldFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"projectId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIJBPaymentTerminal\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"Migrate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fundingCycleConfiguration\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fundingCycleNumber\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"projectId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"beneficiaryTokenCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"memo\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"Pay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"projectId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"wasHeld\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"ProcessFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fundingCycleConfiguration\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fundingCycleNumber\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"projectId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reclaimedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"memo\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"RedeemTokens\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"projectId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"refundedFees\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"leftoverAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"RefundHeldFees\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"SetFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIJBFeeGauge\",\"name\":\"feeGauge\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"SetFeeGauge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addrs\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"flag\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"SetFeelessAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fundingCycleConfiguration\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fundingCycleNumber\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"projectId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"distributedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netDistributedamount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"memo\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"UseAllowance\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"}],\"name\":\"acceptsToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_memo\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"_metadata\",\"type\":\"bytes\"}],\"name\":\"addToBalanceOf\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseWeightCurrency\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currency\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"currencyForToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"}],\"name\":\"currentEthOverflowOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"decimalsForToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"directory\",\"outputs\":[{\"internalType\":\"contractIJBDirectory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_currency\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_minReturnedTokens\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_memo\",\"type\":\"string\"}],\"name\":\"distributePayoutsOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"netLeftoverDistributionAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeGauge\",\"outputs\":[{\"internalType\":\"contractIJBFeeGauge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"}],\"name\":\"heldFeesOf\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"fee\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"feeDiscount\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"internalType\":\"structJBFee[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isFeelessAddress\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"contractIJBPaymentTerminal\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"migrate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operatorStore\",\"outputs\":[{\"internalType\":\"contractIJBOperatorStore\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_beneficiary\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_minReturnedTokens\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_preferClaimedTokens\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"_memo\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"_metadata\",\"type\":\"bytes\"}],\"name\":\"pay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"payoutSplitsGroup\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"prices\",\"outputs\":[{\"internalType\":\"contractIJBPrices\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"}],\"name\":\"processFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"projects\",\"outputs\":[{\"internalType\":\"contractIJBProjects\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_holder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tokenCount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_minReturnedTokens\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"_beneficiary\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_memo\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"_metadata\",\"type\":\"bytes\"}],\"name\":\"redeemTokensOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"reclaimAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"name\":\"setFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIJBFeeGauge\",\"name\":\"_feeGauge\",\"type\":\"address\"}],\"name\":\"setFeeGauge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_flag\",\"type\":\"bool\"}],\"name\":\"setFeelessAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"splitsStore\",\"outputs\":[{\"internalType\":\"contractIJBSplitsStore\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"store\",\"outputs\":[{\"internalType\":\"contractIJBSingleTokenPaymentTerminalStore\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_currency\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_minReturnedTokens\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"_beneficiary\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_memo\",\"type\":\"string\"}],\"name\":\"useAllowanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"netDistributedAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ETHPaymentTerminalABI is the input ABI used to generate the binding from.
// Deprecated: Use ETHPaymentTerminalMetaData.ABI instead.
var ETHPaymentTerminalABI = ETHPaymentTerminalMetaData.ABI

// ETHPaymentTerminal is an auto generated Go binding around an Ethereum contract.
type ETHPaymentTerminal struct {
	ETHPaymentTerminalCaller     // Read-only binding to the contract
	ETHPaymentTerminalTransactor // Write-only binding to the contract
	ETHPaymentTerminalFilterer   // Log filterer for contract events
}

// ETHPaymentTerminalCaller is an auto generated read-only Go binding around an Ethereum contract.
type ETHPaymentTerminalCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ETHPaymentTerminalTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ETHPaymentTerminalTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ETHPaymentTerminalFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ETHPaymentTerminalFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ETHPaymentTerminalSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ETHPaymentTerminalSession struct {
	Contract     *ETHPaymentTerminal // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ETHPaymentTerminalCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ETHPaymentTerminalCallerSession struct {
	Contract *ETHPaymentTerminalCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ETHPaymentTerminalTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ETHPaymentTerminalTransactorSession struct {
	Contract     *ETHPaymentTerminalTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ETHPaymentTerminalRaw is an auto generated low-level Go binding around an Ethereum contract.
type ETHPaymentTerminalRaw struct {
	Contract *ETHPaymentTerminal // Generic contract binding to access the raw methods on
}

// ETHPaymentTerminalCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ETHPaymentTerminalCallerRaw struct {
	Contract *ETHPaymentTerminalCaller // Generic read-only contract binding to access the raw methods on
}

// ETHPaymentTerminalTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ETHPaymentTerminalTransactorRaw struct {
	Contract *ETHPaymentTerminalTransactor // Generic write-only contract binding to access the raw methods on
}

// NewETHPaymentTerminal creates a new instance of ETHPaymentTerminal, bound to a specific deployed contract.
func NewETHPaymentTerminal(address common.Address, backend bind.ContractBackend) (*ETHPaymentTerminal, error) {
	contract, err := bindETHPaymentTerminal(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ETHPaymentTerminal{ETHPaymentTerminalCaller: ETHPaymentTerminalCaller{contract: contract}, ETHPaymentTerminalTransactor: ETHPaymentTerminalTransactor{contract: contract}, ETHPaymentTerminalFilterer: ETHPaymentTerminalFilterer{contract: contract}}, nil
}

// NewETHPaymentTerminalCaller creates a new read-only instance of ETHPaymentTerminal, bound to a specific deployed contract.
func NewETHPaymentTerminalCaller(address common.Address, caller bind.ContractCaller) (*ETHPaymentTerminalCaller, error) {
	contract, err := bindETHPaymentTerminal(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ETHPaymentTerminalCaller{contract: contract}, nil
}

// NewETHPaymentTerminalTransactor creates a new write-only instance of ETHPaymentTerminal, bound to a specific deployed contract.
func NewETHPaymentTerminalTransactor(address common.Address, transactor bind.ContractTransactor) (*ETHPaymentTerminalTransactor, error) {
	contract, err := bindETHPaymentTerminal(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ETHPaymentTerminalTransactor{contract: contract}, nil
}

// NewETHPaymentTerminalFilterer creates a new log filterer instance of ETHPaymentTerminal, bound to a specific deployed contract.
func NewETHPaymentTerminalFilterer(address common.Address, filterer bind.ContractFilterer) (*ETHPaymentTerminalFilterer, error) {
	contract, err := bindETHPaymentTerminal(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ETHPaymentTerminalFilterer{contract: contract}, nil
}

// bindETHPaymentTerminal binds a generic wrapper to an already deployed contract.
func bindETHPaymentTerminal(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ETHPaymentTerminalABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ETHPaymentTerminal *ETHPaymentTerminalRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ETHPaymentTerminal.Contract.ETHPaymentTerminalCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ETHPaymentTerminal *ETHPaymentTerminalRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETHPaymentTerminal.Contract.ETHPaymentTerminalTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ETHPaymentTerminal *ETHPaymentTerminalRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ETHPaymentTerminal.Contract.ETHPaymentTerminalTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ETHPaymentTerminal *ETHPaymentTerminalCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ETHPaymentTerminal.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ETHPaymentTerminal *ETHPaymentTerminalTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETHPaymentTerminal.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ETHPaymentTerminal *ETHPaymentTerminalTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ETHPaymentTerminal.Contract.contract.Transact(opts, method, params...)
}

// AcceptsToken is a free data retrieval call binding the contract method 0xdf21a7dd.
//
// Solidity: function acceptsToken(address _token, uint256 _projectId) view returns(bool)
func (_ETHPaymentTerminal *ETHPaymentTerminalCaller) AcceptsToken(opts *bind.CallOpts, _token common.Address, _projectId *big.Int) (bool, error) {
	var out []interface{}
	err := _ETHPaymentTerminal.contract.Call(opts, &out, "acceptsToken", _token, _projectId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AcceptsToken is a free data retrieval call binding the contract method 0xdf21a7dd.
//
// Solidity: function acceptsToken(address _token, uint256 _projectId) view returns(bool)
func (_ETHPaymentTerminal *ETHPaymentTerminalSession) AcceptsToken(_token common.Address, _projectId *big.Int) (bool, error) {
	return _ETHPaymentTerminal.Contract.AcceptsToken(&_ETHPaymentTerminal.CallOpts, _token, _projectId)
}

// AcceptsToken is a free data retrieval call binding the contract method 0xdf21a7dd.
//
// Solidity: function acceptsToken(address _token, uint256 _projectId) view returns(bool)
func (_ETHPaymentTerminal *ETHPaymentTerminalCallerSession) AcceptsToken(_token common.Address, _projectId *big.Int) (bool, error) {
	return _ETHPaymentTerminal.Contract.AcceptsToken(&_ETHPaymentTerminal.CallOpts, _token, _projectId)
}

// BaseWeightCurrency is a free data retrieval call binding the contract method 0x2d1a5903.
//
// Solidity: function baseWeightCurrency() view returns(uint256)
func (_ETHPaymentTerminal *ETHPaymentTerminalCaller) BaseWeightCurrency(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ETHPaymentTerminal.contract.Call(opts, &out, "baseWeightCurrency")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseWeightCurrency is a free data retrieval call binding the contract method 0x2d1a5903.
//
// Solidity: function baseWeightCurrency() view returns(uint256)
func (_ETHPaymentTerminal *ETHPaymentTerminalSession) BaseWeightCurrency() (*big.Int, error) {
	return _ETHPaymentTerminal.Contract.BaseWeightCurrency(&_ETHPaymentTerminal.CallOpts)
}

// BaseWeightCurrency is a free data retrieval call binding the contract method 0x2d1a5903.
//
// Solidity: function baseWeightCurrency() view returns(uint256)
func (_ETHPaymentTerminal *ETHPaymentTerminalCallerSession) BaseWeightCurrency() (*big.Int, error) {
	return _ETHPaymentTerminal.Contract.BaseWeightCurrency(&_ETHPaymentTerminal.CallOpts)
}

// Currency is a free data retrieval call binding the contract method 0xe5a6b10f.
//
// Solidity: function currency() view returns(uint256)
func (_ETHPaymentTerminal *ETHPaymentTerminalCaller) Currency(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ETHPaymentTerminal.contract.Call(opts, &out, "currency")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Currency is a free data retrieval call binding the contract method 0xe5a6b10f.
//
// Solidity: function currency() view returns(uint256)
func (_ETHPaymentTerminal *ETHPaymentTerminalSession) Currency() (*big.Int, error) {
	return _ETHPaymentTerminal.Contract.Currency(&_ETHPaymentTerminal.CallOpts)
}

// Currency is a free data retrieval call binding the contract method 0xe5a6b10f.
//
// Solidity: function currency() view returns(uint256)
func (_ETHPaymentTerminal *ETHPaymentTerminalCallerSession) Currency() (*big.Int, error) {
	return _ETHPaymentTerminal.Contract.Currency(&_ETHPaymentTerminal.CallOpts)
}

// CurrencyForToken is a free data retrieval call binding the contract method 0x1982d679.
//
// Solidity: function currencyForToken(address _token) view returns(uint256)
func (_ETHPaymentTerminal *ETHPaymentTerminalCaller) CurrencyForToken(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ETHPaymentTerminal.contract.Call(opts, &out, "currencyForToken", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrencyForToken is a free data retrieval call binding the contract method 0x1982d679.
//
// Solidity: function currencyForToken(address _token) view returns(uint256)
func (_ETHPaymentTerminal *ETHPaymentTerminalSession) CurrencyForToken(_token common.Address) (*big.Int, error) {
	return _ETHPaymentTerminal.Contract.CurrencyForToken(&_ETHPaymentTerminal.CallOpts, _token)
}

// CurrencyForToken is a free data retrieval call binding the contract method 0x1982d679.
//
// Solidity: function currencyForToken(address _token) view returns(uint256)
func (_ETHPaymentTerminal *ETHPaymentTerminalCallerSession) CurrencyForToken(_token common.Address) (*big.Int, error) {
	return _ETHPaymentTerminal.Contract.CurrencyForToken(&_ETHPaymentTerminal.CallOpts, _token)
}

// CurrentEthOverflowOf is a free data retrieval call binding the contract method 0xa32e1e96.
//
// Solidity: function currentEthOverflowOf(uint256 _projectId) view returns(uint256)
func (_ETHPaymentTerminal *ETHPaymentTerminalCaller) CurrentEthOverflowOf(opts *bind.CallOpts, _projectId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ETHPaymentTerminal.contract.Call(opts, &out, "currentEthOverflowOf", _projectId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentEthOverflowOf is a free data retrieval call binding the contract method 0xa32e1e96.
//
// Solidity: function currentEthOverflowOf(uint256 _projectId) view returns(uint256)
func (_ETHPaymentTerminal *ETHPaymentTerminalSession) CurrentEthOverflowOf(_projectId *big.Int) (*big.Int, error) {
	return _ETHPaymentTerminal.Contract.CurrentEthOverflowOf(&_ETHPaymentTerminal.CallOpts, _projectId)
}

// CurrentEthOverflowOf is a free data retrieval call binding the contract method 0xa32e1e96.
//
// Solidity: function currentEthOverflowOf(uint256 _projectId) view returns(uint256)
func (_ETHPaymentTerminal *ETHPaymentTerminalCallerSession) CurrentEthOverflowOf(_projectId *big.Int) (*big.Int, error) {
	return _ETHPaymentTerminal.Contract.CurrentEthOverflowOf(&_ETHPaymentTerminal.CallOpts, _projectId)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_ETHPaymentTerminal *ETHPaymentTerminalCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ETHPaymentTerminal.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_ETHPaymentTerminal *ETHPaymentTerminalSession) Decimals() (*big.Int, error) {
	return _ETHPaymentTerminal.Contract.Decimals(&_ETHPaymentTerminal.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_ETHPaymentTerminal *ETHPaymentTerminalCallerSession) Decimals() (*big.Int, error) {
	return _ETHPaymentTerminal.Contract.Decimals(&_ETHPaymentTerminal.CallOpts)
}

// DecimalsForToken is a free data retrieval call binding the contract method 0xb7bad1b1.
//
// Solidity: function decimalsForToken(address _token) view returns(uint256)
func (_ETHPaymentTerminal *ETHPaymentTerminalCaller) DecimalsForToken(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ETHPaymentTerminal.contract.Call(opts, &out, "decimalsForToken", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DecimalsForToken is a free data retrieval call binding the contract method 0xb7bad1b1.
//
// Solidity: function decimalsForToken(address _token) view returns(uint256)
func (_ETHPaymentTerminal *ETHPaymentTerminalSession) DecimalsForToken(_token common.Address) (*big.Int, error) {
	return _ETHPaymentTerminal.Contract.DecimalsForToken(&_ETHPaymentTerminal.CallOpts, _token)
}

// DecimalsForToken is a free data retrieval call binding the contract method 0xb7bad1b1.
//
// Solidity: function decimalsForToken(address _token) view returns(uint256)
func (_ETHPaymentTerminal *ETHPaymentTerminalCallerSession) DecimalsForToken(_token common.Address) (*big.Int, error) {
	return _ETHPaymentTerminal.Contract.DecimalsForToken(&_ETHPaymentTerminal.CallOpts, _token)
}

// Directory is a free data retrieval call binding the contract method 0xc41c2f24.
//
// Solidity: function directory() view returns(address)
func (_ETHPaymentTerminal *ETHPaymentTerminalCaller) Directory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ETHPaymentTerminal.contract.Call(opts, &out, "directory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Directory is a free data retrieval call binding the contract method 0xc41c2f24.
//
// Solidity: function directory() view returns(address)
func (_ETHPaymentTerminal *ETHPaymentTerminalSession) Directory() (common.Address, error) {
	return _ETHPaymentTerminal.Contract.Directory(&_ETHPaymentTerminal.CallOpts)
}

// Directory is a free data retrieval call binding the contract method 0xc41c2f24.
//
// Solidity: function directory() view returns(address)
func (_ETHPaymentTerminal *ETHPaymentTerminalCallerSession) Directory() (common.Address, error) {
	return _ETHPaymentTerminal.Contract.Directory(&_ETHPaymentTerminal.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_ETHPaymentTerminal *ETHPaymentTerminalCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ETHPaymentTerminal.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_ETHPaymentTerminal *ETHPaymentTerminalSession) Fee() (*big.Int, error) {
	return _ETHPaymentTerminal.Contract.Fee(&_ETHPaymentTerminal.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_ETHPaymentTerminal *ETHPaymentTerminalCallerSession) Fee() (*big.Int, error) {
	return _ETHPaymentTerminal.Contract.Fee(&_ETHPaymentTerminal.CallOpts)
}

// FeeGauge is a free data retrieval call binding the contract method 0xd6dacc53.
//
// Solidity: function feeGauge() view returns(address)
func (_ETHPaymentTerminal *ETHPaymentTerminalCaller) FeeGauge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ETHPaymentTerminal.contract.Call(opts, &out, "feeGauge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeGauge is a free data retrieval call binding the contract method 0xd6dacc53.
//
// Solidity: function feeGauge() view returns(address)
func (_ETHPaymentTerminal *ETHPaymentTerminalSession) FeeGauge() (common.Address, error) {
	return _ETHPaymentTerminal.Contract.FeeGauge(&_ETHPaymentTerminal.CallOpts)
}

// FeeGauge is a free data retrieval call binding the contract method 0xd6dacc53.
//
// Solidity: function feeGauge() view returns(address)
func (_ETHPaymentTerminal *ETHPaymentTerminalCallerSession) FeeGauge() (common.Address, error) {
	return _ETHPaymentTerminal.Contract.FeeGauge(&_ETHPaymentTerminal.CallOpts)
}

// HeldFeesOf is a free data retrieval call binding the contract method 0x8af56094.
//
// Solidity: function heldFeesOf(uint256 _projectId) view returns((uint256,uint32,uint32,address)[])
func (_ETHPaymentTerminal *ETHPaymentTerminalCaller) HeldFeesOf(opts *bind.CallOpts, _projectId *big.Int) ([]JBFee, error) {
	var out []interface{}
	err := _ETHPaymentTerminal.contract.Call(opts, &out, "heldFeesOf", _projectId)

	if err != nil {
		return *new([]JBFee), err
	}

	out0 := *abi.ConvertType(out[0], new([]JBFee)).(*[]JBFee)

	return out0, err

}

// HeldFeesOf is a free data retrieval call binding the contract method 0x8af56094.
//
// Solidity: function heldFeesOf(uint256 _projectId) view returns((uint256,uint32,uint32,address)[])
func (_ETHPaymentTerminal *ETHPaymentTerminalSession) HeldFeesOf(_projectId *big.Int) ([]JBFee, error) {
	return _ETHPaymentTerminal.Contract.HeldFeesOf(&_ETHPaymentTerminal.CallOpts, _projectId)
}

// HeldFeesOf is a free data retrieval call binding the contract method 0x8af56094.
//
// Solidity: function heldFeesOf(uint256 _projectId) view returns((uint256,uint32,uint32,address)[])
func (_ETHPaymentTerminal *ETHPaymentTerminalCallerSession) HeldFeesOf(_projectId *big.Int) ([]JBFee, error) {
	return _ETHPaymentTerminal.Contract.HeldFeesOf(&_ETHPaymentTerminal.CallOpts, _projectId)
}

// IsFeelessAddress is a free data retrieval call binding the contract method 0xb631b500.
//
// Solidity: function isFeelessAddress(address ) view returns(bool)
func (_ETHPaymentTerminal *ETHPaymentTerminalCaller) IsFeelessAddress(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ETHPaymentTerminal.contract.Call(opts, &out, "isFeelessAddress", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFeelessAddress is a free data retrieval call binding the contract method 0xb631b500.
//
// Solidity: function isFeelessAddress(address ) view returns(bool)
func (_ETHPaymentTerminal *ETHPaymentTerminalSession) IsFeelessAddress(arg0 common.Address) (bool, error) {
	return _ETHPaymentTerminal.Contract.IsFeelessAddress(&_ETHPaymentTerminal.CallOpts, arg0)
}

// IsFeelessAddress is a free data retrieval call binding the contract method 0xb631b500.
//
// Solidity: function isFeelessAddress(address ) view returns(bool)
func (_ETHPaymentTerminal *ETHPaymentTerminalCallerSession) IsFeelessAddress(arg0 common.Address) (bool, error) {
	return _ETHPaymentTerminal.Contract.IsFeelessAddress(&_ETHPaymentTerminal.CallOpts, arg0)
}

// OperatorStore is a free data retrieval call binding the contract method 0xad007d63.
//
// Solidity: function operatorStore() view returns(address)
func (_ETHPaymentTerminal *ETHPaymentTerminalCaller) OperatorStore(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ETHPaymentTerminal.contract.Call(opts, &out, "operatorStore")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OperatorStore is a free data retrieval call binding the contract method 0xad007d63.
//
// Solidity: function operatorStore() view returns(address)
func (_ETHPaymentTerminal *ETHPaymentTerminalSession) OperatorStore() (common.Address, error) {
	return _ETHPaymentTerminal.Contract.OperatorStore(&_ETHPaymentTerminal.CallOpts)
}

// OperatorStore is a free data retrieval call binding the contract method 0xad007d63.
//
// Solidity: function operatorStore() view returns(address)
func (_ETHPaymentTerminal *ETHPaymentTerminalCallerSession) OperatorStore() (common.Address, error) {
	return _ETHPaymentTerminal.Contract.OperatorStore(&_ETHPaymentTerminal.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ETHPaymentTerminal *ETHPaymentTerminalCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ETHPaymentTerminal.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ETHPaymentTerminal *ETHPaymentTerminalSession) Owner() (common.Address, error) {
	return _ETHPaymentTerminal.Contract.Owner(&_ETHPaymentTerminal.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ETHPaymentTerminal *ETHPaymentTerminalCallerSession) Owner() (common.Address, error) {
	return _ETHPaymentTerminal.Contract.Owner(&_ETHPaymentTerminal.CallOpts)
}

// PayoutSplitsGroup is a free data retrieval call binding the contract method 0x66248b86.
//
// Solidity: function payoutSplitsGroup() view returns(uint256)
func (_ETHPaymentTerminal *ETHPaymentTerminalCaller) PayoutSplitsGroup(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ETHPaymentTerminal.contract.Call(opts, &out, "payoutSplitsGroup")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PayoutSplitsGroup is a free data retrieval call binding the contract method 0x66248b86.
//
// Solidity: function payoutSplitsGroup() view returns(uint256)
func (_ETHPaymentTerminal *ETHPaymentTerminalSession) PayoutSplitsGroup() (*big.Int, error) {
	return _ETHPaymentTerminal.Contract.PayoutSplitsGroup(&_ETHPaymentTerminal.CallOpts)
}

// PayoutSplitsGroup is a free data retrieval call binding the contract method 0x66248b86.
//
// Solidity: function payoutSplitsGroup() view returns(uint256)
func (_ETHPaymentTerminal *ETHPaymentTerminalCallerSession) PayoutSplitsGroup() (*big.Int, error) {
	return _ETHPaymentTerminal.Contract.PayoutSplitsGroup(&_ETHPaymentTerminal.CallOpts)
}

// Prices is a free data retrieval call binding the contract method 0xd3419bf3.
//
// Solidity: function prices() view returns(address)
func (_ETHPaymentTerminal *ETHPaymentTerminalCaller) Prices(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ETHPaymentTerminal.contract.Call(opts, &out, "prices")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Prices is a free data retrieval call binding the contract method 0xd3419bf3.
//
// Solidity: function prices() view returns(address)
func (_ETHPaymentTerminal *ETHPaymentTerminalSession) Prices() (common.Address, error) {
	return _ETHPaymentTerminal.Contract.Prices(&_ETHPaymentTerminal.CallOpts)
}

// Prices is a free data retrieval call binding the contract method 0xd3419bf3.
//
// Solidity: function prices() view returns(address)
func (_ETHPaymentTerminal *ETHPaymentTerminalCallerSession) Prices() (common.Address, error) {
	return _ETHPaymentTerminal.Contract.Prices(&_ETHPaymentTerminal.CallOpts)
}

// Projects is a free data retrieval call binding the contract method 0x8b79543c.
//
// Solidity: function projects() view returns(address)
func (_ETHPaymentTerminal *ETHPaymentTerminalCaller) Projects(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ETHPaymentTerminal.contract.Call(opts, &out, "projects")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Projects is a free data retrieval call binding the contract method 0x8b79543c.
//
// Solidity: function projects() view returns(address)
func (_ETHPaymentTerminal *ETHPaymentTerminalSession) Projects() (common.Address, error) {
	return _ETHPaymentTerminal.Contract.Projects(&_ETHPaymentTerminal.CallOpts)
}

// Projects is a free data retrieval call binding the contract method 0x8b79543c.
//
// Solidity: function projects() view returns(address)
func (_ETHPaymentTerminal *ETHPaymentTerminalCallerSession) Projects() (common.Address, error) {
	return _ETHPaymentTerminal.Contract.Projects(&_ETHPaymentTerminal.CallOpts)
}

// SplitsStore is a free data retrieval call binding the contract method 0x2bdfe004.
//
// Solidity: function splitsStore() view returns(address)
func (_ETHPaymentTerminal *ETHPaymentTerminalCaller) SplitsStore(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ETHPaymentTerminal.contract.Call(opts, &out, "splitsStore")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SplitsStore is a free data retrieval call binding the contract method 0x2bdfe004.
//
// Solidity: function splitsStore() view returns(address)
func (_ETHPaymentTerminal *ETHPaymentTerminalSession) SplitsStore() (common.Address, error) {
	return _ETHPaymentTerminal.Contract.SplitsStore(&_ETHPaymentTerminal.CallOpts)
}

// SplitsStore is a free data retrieval call binding the contract method 0x2bdfe004.
//
// Solidity: function splitsStore() view returns(address)
func (_ETHPaymentTerminal *ETHPaymentTerminalCallerSession) SplitsStore() (common.Address, error) {
	return _ETHPaymentTerminal.Contract.SplitsStore(&_ETHPaymentTerminal.CallOpts)
}

// Store is a free data retrieval call binding the contract method 0x975057e7.
//
// Solidity: function store() view returns(address)
func (_ETHPaymentTerminal *ETHPaymentTerminalCaller) Store(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ETHPaymentTerminal.contract.Call(opts, &out, "store")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Store is a free data retrieval call binding the contract method 0x975057e7.
//
// Solidity: function store() view returns(address)
func (_ETHPaymentTerminal *ETHPaymentTerminalSession) Store() (common.Address, error) {
	return _ETHPaymentTerminal.Contract.Store(&_ETHPaymentTerminal.CallOpts)
}

// Store is a free data retrieval call binding the contract method 0x975057e7.
//
// Solidity: function store() view returns(address)
func (_ETHPaymentTerminal *ETHPaymentTerminalCallerSession) Store() (common.Address, error) {
	return _ETHPaymentTerminal.Contract.Store(&_ETHPaymentTerminal.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (_ETHPaymentTerminal *ETHPaymentTerminalCaller) SupportsInterface(opts *bind.CallOpts, _interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ETHPaymentTerminal.contract.Call(opts, &out, "supportsInterface", _interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (_ETHPaymentTerminal *ETHPaymentTerminalSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _ETHPaymentTerminal.Contract.SupportsInterface(&_ETHPaymentTerminal.CallOpts, _interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (_ETHPaymentTerminal *ETHPaymentTerminalCallerSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _ETHPaymentTerminal.Contract.SupportsInterface(&_ETHPaymentTerminal.CallOpts, _interfaceId)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_ETHPaymentTerminal *ETHPaymentTerminalCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ETHPaymentTerminal.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_ETHPaymentTerminal *ETHPaymentTerminalSession) Token() (common.Address, error) {
	return _ETHPaymentTerminal.Contract.Token(&_ETHPaymentTerminal.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_ETHPaymentTerminal *ETHPaymentTerminalCallerSession) Token() (common.Address, error) {
	return _ETHPaymentTerminal.Contract.Token(&_ETHPaymentTerminal.CallOpts)
}

// AddToBalanceOf is a paid mutator transaction binding the contract method 0x0cf8e858.
//
// Solidity: function addToBalanceOf(uint256 _projectId, uint256 _amount, address _token, string _memo, bytes _metadata) payable returns()
func (_ETHPaymentTerminal *ETHPaymentTerminalTransactor) AddToBalanceOf(opts *bind.TransactOpts, _projectId *big.Int, _amount *big.Int, _token common.Address, _memo string, _metadata []byte) (*types.Transaction, error) {
	return _ETHPaymentTerminal.contract.Transact(opts, "addToBalanceOf", _projectId, _amount, _token, _memo, _metadata)
}

// AddToBalanceOf is a paid mutator transaction binding the contract method 0x0cf8e858.
//
// Solidity: function addToBalanceOf(uint256 _projectId, uint256 _amount, address _token, string _memo, bytes _metadata) payable returns()
func (_ETHPaymentTerminal *ETHPaymentTerminalSession) AddToBalanceOf(_projectId *big.Int, _amount *big.Int, _token common.Address, _memo string, _metadata []byte) (*types.Transaction, error) {
	return _ETHPaymentTerminal.Contract.AddToBalanceOf(&_ETHPaymentTerminal.TransactOpts, _projectId, _amount, _token, _memo, _metadata)
}

// AddToBalanceOf is a paid mutator transaction binding the contract method 0x0cf8e858.
//
// Solidity: function addToBalanceOf(uint256 _projectId, uint256 _amount, address _token, string _memo, bytes _metadata) payable returns()
func (_ETHPaymentTerminal *ETHPaymentTerminalTransactorSession) AddToBalanceOf(_projectId *big.Int, _amount *big.Int, _token common.Address, _memo string, _metadata []byte) (*types.Transaction, error) {
	return _ETHPaymentTerminal.Contract.AddToBalanceOf(&_ETHPaymentTerminal.TransactOpts, _projectId, _amount, _token, _memo, _metadata)
}

// DistributePayoutsOf is a paid mutator transaction binding the contract method 0x2b267b4e.
//
// Solidity: function distributePayoutsOf(uint256 _projectId, uint256 _amount, uint256 _currency, address _token, uint256 _minReturnedTokens, string _memo) returns(uint256 netLeftoverDistributionAmount)
func (_ETHPaymentTerminal *ETHPaymentTerminalTransactor) DistributePayoutsOf(opts *bind.TransactOpts, _projectId *big.Int, _amount *big.Int, _currency *big.Int, _token common.Address, _minReturnedTokens *big.Int, _memo string) (*types.Transaction, error) {
	return _ETHPaymentTerminal.contract.Transact(opts, "distributePayoutsOf", _projectId, _amount, _currency, _token, _minReturnedTokens, _memo)
}

// DistributePayoutsOf is a paid mutator transaction binding the contract method 0x2b267b4e.
//
// Solidity: function distributePayoutsOf(uint256 _projectId, uint256 _amount, uint256 _currency, address _token, uint256 _minReturnedTokens, string _memo) returns(uint256 netLeftoverDistributionAmount)
func (_ETHPaymentTerminal *ETHPaymentTerminalSession) DistributePayoutsOf(_projectId *big.Int, _amount *big.Int, _currency *big.Int, _token common.Address, _minReturnedTokens *big.Int, _memo string) (*types.Transaction, error) {
	return _ETHPaymentTerminal.Contract.DistributePayoutsOf(&_ETHPaymentTerminal.TransactOpts, _projectId, _amount, _currency, _token, _minReturnedTokens, _memo)
}

// DistributePayoutsOf is a paid mutator transaction binding the contract method 0x2b267b4e.
//
// Solidity: function distributePayoutsOf(uint256 _projectId, uint256 _amount, uint256 _currency, address _token, uint256 _minReturnedTokens, string _memo) returns(uint256 netLeftoverDistributionAmount)
func (_ETHPaymentTerminal *ETHPaymentTerminalTransactorSession) DistributePayoutsOf(_projectId *big.Int, _amount *big.Int, _currency *big.Int, _token common.Address, _minReturnedTokens *big.Int, _memo string) (*types.Transaction, error) {
	return _ETHPaymentTerminal.Contract.DistributePayoutsOf(&_ETHPaymentTerminal.TransactOpts, _projectId, _amount, _currency, _token, _minReturnedTokens, _memo)
}

// Migrate is a paid mutator transaction binding the contract method 0x405b84fa.
//
// Solidity: function migrate(uint256 _projectId, address _to) returns(uint256 balance)
func (_ETHPaymentTerminal *ETHPaymentTerminalTransactor) Migrate(opts *bind.TransactOpts, _projectId *big.Int, _to common.Address) (*types.Transaction, error) {
	return _ETHPaymentTerminal.contract.Transact(opts, "migrate", _projectId, _to)
}

// Migrate is a paid mutator transaction binding the contract method 0x405b84fa.
//
// Solidity: function migrate(uint256 _projectId, address _to) returns(uint256 balance)
func (_ETHPaymentTerminal *ETHPaymentTerminalSession) Migrate(_projectId *big.Int, _to common.Address) (*types.Transaction, error) {
	return _ETHPaymentTerminal.Contract.Migrate(&_ETHPaymentTerminal.TransactOpts, _projectId, _to)
}

// Migrate is a paid mutator transaction binding the contract method 0x405b84fa.
//
// Solidity: function migrate(uint256 _projectId, address _to) returns(uint256 balance)
func (_ETHPaymentTerminal *ETHPaymentTerminalTransactorSession) Migrate(_projectId *big.Int, _to common.Address) (*types.Transaction, error) {
	return _ETHPaymentTerminal.Contract.Migrate(&_ETHPaymentTerminal.TransactOpts, _projectId, _to)
}

// Pay is a paid mutator transaction binding the contract method 0x1ebc263f.
//
// Solidity: function pay(uint256 _projectId, uint256 _amount, address _token, address _beneficiary, uint256 _minReturnedTokens, bool _preferClaimedTokens, string _memo, bytes _metadata) payable returns(uint256)
func (_ETHPaymentTerminal *ETHPaymentTerminalTransactor) Pay(opts *bind.TransactOpts, _projectId *big.Int, _amount *big.Int, _token common.Address, _beneficiary common.Address, _minReturnedTokens *big.Int, _preferClaimedTokens bool, _memo string, _metadata []byte) (*types.Transaction, error) {
	return _ETHPaymentTerminal.contract.Transact(opts, "pay", _projectId, _amount, _token, _beneficiary, _minReturnedTokens, _preferClaimedTokens, _memo, _metadata)
}

// Pay is a paid mutator transaction binding the contract method 0x1ebc263f.
//
// Solidity: function pay(uint256 _projectId, uint256 _amount, address _token, address _beneficiary, uint256 _minReturnedTokens, bool _preferClaimedTokens, string _memo, bytes _metadata) payable returns(uint256)
func (_ETHPaymentTerminal *ETHPaymentTerminalSession) Pay(_projectId *big.Int, _amount *big.Int, _token common.Address, _beneficiary common.Address, _minReturnedTokens *big.Int, _preferClaimedTokens bool, _memo string, _metadata []byte) (*types.Transaction, error) {
	return _ETHPaymentTerminal.Contract.Pay(&_ETHPaymentTerminal.TransactOpts, _projectId, _amount, _token, _beneficiary, _minReturnedTokens, _preferClaimedTokens, _memo, _metadata)
}

// Pay is a paid mutator transaction binding the contract method 0x1ebc263f.
//
// Solidity: function pay(uint256 _projectId, uint256 _amount, address _token, address _beneficiary, uint256 _minReturnedTokens, bool _preferClaimedTokens, string _memo, bytes _metadata) payable returns(uint256)
func (_ETHPaymentTerminal *ETHPaymentTerminalTransactorSession) Pay(_projectId *big.Int, _amount *big.Int, _token common.Address, _beneficiary common.Address, _minReturnedTokens *big.Int, _preferClaimedTokens bool, _memo string, _metadata []byte) (*types.Transaction, error) {
	return _ETHPaymentTerminal.Contract.Pay(&_ETHPaymentTerminal.TransactOpts, _projectId, _amount, _token, _beneficiary, _minReturnedTokens, _preferClaimedTokens, _memo, _metadata)
}

// ProcessFees is a paid mutator transaction binding the contract method 0x89701db5.
//
// Solidity: function processFees(uint256 _projectId) returns()
func (_ETHPaymentTerminal *ETHPaymentTerminalTransactor) ProcessFees(opts *bind.TransactOpts, _projectId *big.Int) (*types.Transaction, error) {
	return _ETHPaymentTerminal.contract.Transact(opts, "processFees", _projectId)
}

// ProcessFees is a paid mutator transaction binding the contract method 0x89701db5.
//
// Solidity: function processFees(uint256 _projectId) returns()
func (_ETHPaymentTerminal *ETHPaymentTerminalSession) ProcessFees(_projectId *big.Int) (*types.Transaction, error) {
	return _ETHPaymentTerminal.Contract.ProcessFees(&_ETHPaymentTerminal.TransactOpts, _projectId)
}

// ProcessFees is a paid mutator transaction binding the contract method 0x89701db5.
//
// Solidity: function processFees(uint256 _projectId) returns()
func (_ETHPaymentTerminal *ETHPaymentTerminalTransactorSession) ProcessFees(_projectId *big.Int) (*types.Transaction, error) {
	return _ETHPaymentTerminal.Contract.ProcessFees(&_ETHPaymentTerminal.TransactOpts, _projectId)
}

// RedeemTokensOf is a paid mutator transaction binding the contract method 0xfe663f0f.
//
// Solidity: function redeemTokensOf(address _holder, uint256 _projectId, uint256 _tokenCount, address _token, uint256 _minReturnedTokens, address _beneficiary, string _memo, bytes _metadata) returns(uint256 reclaimAmount)
func (_ETHPaymentTerminal *ETHPaymentTerminalTransactor) RedeemTokensOf(opts *bind.TransactOpts, _holder common.Address, _projectId *big.Int, _tokenCount *big.Int, _token common.Address, _minReturnedTokens *big.Int, _beneficiary common.Address, _memo string, _metadata []byte) (*types.Transaction, error) {
	return _ETHPaymentTerminal.contract.Transact(opts, "redeemTokensOf", _holder, _projectId, _tokenCount, _token, _minReturnedTokens, _beneficiary, _memo, _metadata)
}

// RedeemTokensOf is a paid mutator transaction binding the contract method 0xfe663f0f.
//
// Solidity: function redeemTokensOf(address _holder, uint256 _projectId, uint256 _tokenCount, address _token, uint256 _minReturnedTokens, address _beneficiary, string _memo, bytes _metadata) returns(uint256 reclaimAmount)
func (_ETHPaymentTerminal *ETHPaymentTerminalSession) RedeemTokensOf(_holder common.Address, _projectId *big.Int, _tokenCount *big.Int, _token common.Address, _minReturnedTokens *big.Int, _beneficiary common.Address, _memo string, _metadata []byte) (*types.Transaction, error) {
	return _ETHPaymentTerminal.Contract.RedeemTokensOf(&_ETHPaymentTerminal.TransactOpts, _holder, _projectId, _tokenCount, _token, _minReturnedTokens, _beneficiary, _memo, _metadata)
}

// RedeemTokensOf is a paid mutator transaction binding the contract method 0xfe663f0f.
//
// Solidity: function redeemTokensOf(address _holder, uint256 _projectId, uint256 _tokenCount, address _token, uint256 _minReturnedTokens, address _beneficiary, string _memo, bytes _metadata) returns(uint256 reclaimAmount)
func (_ETHPaymentTerminal *ETHPaymentTerminalTransactorSession) RedeemTokensOf(_holder common.Address, _projectId *big.Int, _tokenCount *big.Int, _token common.Address, _minReturnedTokens *big.Int, _beneficiary common.Address, _memo string, _metadata []byte) (*types.Transaction, error) {
	return _ETHPaymentTerminal.Contract.RedeemTokensOf(&_ETHPaymentTerminal.TransactOpts, _holder, _projectId, _tokenCount, _token, _minReturnedTokens, _beneficiary, _memo, _metadata)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ETHPaymentTerminal *ETHPaymentTerminalTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETHPaymentTerminal.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ETHPaymentTerminal *ETHPaymentTerminalSession) RenounceOwnership() (*types.Transaction, error) {
	return _ETHPaymentTerminal.Contract.RenounceOwnership(&_ETHPaymentTerminal.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ETHPaymentTerminal *ETHPaymentTerminalTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ETHPaymentTerminal.Contract.RenounceOwnership(&_ETHPaymentTerminal.TransactOpts)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 _fee) returns()
func (_ETHPaymentTerminal *ETHPaymentTerminalTransactor) SetFee(opts *bind.TransactOpts, _fee *big.Int) (*types.Transaction, error) {
	return _ETHPaymentTerminal.contract.Transact(opts, "setFee", _fee)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 _fee) returns()
func (_ETHPaymentTerminal *ETHPaymentTerminalSession) SetFee(_fee *big.Int) (*types.Transaction, error) {
	return _ETHPaymentTerminal.Contract.SetFee(&_ETHPaymentTerminal.TransactOpts, _fee)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 _fee) returns()
func (_ETHPaymentTerminal *ETHPaymentTerminalTransactorSession) SetFee(_fee *big.Int) (*types.Transaction, error) {
	return _ETHPaymentTerminal.Contract.SetFee(&_ETHPaymentTerminal.TransactOpts, _fee)
}

// SetFeeGauge is a paid mutator transaction binding the contract method 0x637913ac.
//
// Solidity: function setFeeGauge(address _feeGauge) returns()
func (_ETHPaymentTerminal *ETHPaymentTerminalTransactor) SetFeeGauge(opts *bind.TransactOpts, _feeGauge common.Address) (*types.Transaction, error) {
	return _ETHPaymentTerminal.contract.Transact(opts, "setFeeGauge", _feeGauge)
}

// SetFeeGauge is a paid mutator transaction binding the contract method 0x637913ac.
//
// Solidity: function setFeeGauge(address _feeGauge) returns()
func (_ETHPaymentTerminal *ETHPaymentTerminalSession) SetFeeGauge(_feeGauge common.Address) (*types.Transaction, error) {
	return _ETHPaymentTerminal.Contract.SetFeeGauge(&_ETHPaymentTerminal.TransactOpts, _feeGauge)
}

// SetFeeGauge is a paid mutator transaction binding the contract method 0x637913ac.
//
// Solidity: function setFeeGauge(address _feeGauge) returns()
func (_ETHPaymentTerminal *ETHPaymentTerminalTransactorSession) SetFeeGauge(_feeGauge common.Address) (*types.Transaction, error) {
	return _ETHPaymentTerminal.Contract.SetFeeGauge(&_ETHPaymentTerminal.TransactOpts, _feeGauge)
}

// SetFeelessAddress is a paid mutator transaction binding the contract method 0x7258002c.
//
// Solidity: function setFeelessAddress(address _address, bool _flag) returns()
func (_ETHPaymentTerminal *ETHPaymentTerminalTransactor) SetFeelessAddress(opts *bind.TransactOpts, _address common.Address, _flag bool) (*types.Transaction, error) {
	return _ETHPaymentTerminal.contract.Transact(opts, "setFeelessAddress", _address, _flag)
}

// SetFeelessAddress is a paid mutator transaction binding the contract method 0x7258002c.
//
// Solidity: function setFeelessAddress(address _address, bool _flag) returns()
func (_ETHPaymentTerminal *ETHPaymentTerminalSession) SetFeelessAddress(_address common.Address, _flag bool) (*types.Transaction, error) {
	return _ETHPaymentTerminal.Contract.SetFeelessAddress(&_ETHPaymentTerminal.TransactOpts, _address, _flag)
}

// SetFeelessAddress is a paid mutator transaction binding the contract method 0x7258002c.
//
// Solidity: function setFeelessAddress(address _address, bool _flag) returns()
func (_ETHPaymentTerminal *ETHPaymentTerminalTransactorSession) SetFeelessAddress(_address common.Address, _flag bool) (*types.Transaction, error) {
	return _ETHPaymentTerminal.Contract.SetFeelessAddress(&_ETHPaymentTerminal.TransactOpts, _address, _flag)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ETHPaymentTerminal *ETHPaymentTerminalTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ETHPaymentTerminal.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ETHPaymentTerminal *ETHPaymentTerminalSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ETHPaymentTerminal.Contract.TransferOwnership(&_ETHPaymentTerminal.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ETHPaymentTerminal *ETHPaymentTerminalTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ETHPaymentTerminal.Contract.TransferOwnership(&_ETHPaymentTerminal.TransactOpts, newOwner)
}

// UseAllowanceOf is a paid mutator transaction binding the contract method 0xbc8926e9.
//
// Solidity: function useAllowanceOf(uint256 _projectId, uint256 _amount, uint256 _currency, address _token, uint256 _minReturnedTokens, address _beneficiary, string _memo) returns(uint256 netDistributedAmount)
func (_ETHPaymentTerminal *ETHPaymentTerminalTransactor) UseAllowanceOf(opts *bind.TransactOpts, _projectId *big.Int, _amount *big.Int, _currency *big.Int, _token common.Address, _minReturnedTokens *big.Int, _beneficiary common.Address, _memo string) (*types.Transaction, error) {
	return _ETHPaymentTerminal.contract.Transact(opts, "useAllowanceOf", _projectId, _amount, _currency, _token, _minReturnedTokens, _beneficiary, _memo)
}

// UseAllowanceOf is a paid mutator transaction binding the contract method 0xbc8926e9.
//
// Solidity: function useAllowanceOf(uint256 _projectId, uint256 _amount, uint256 _currency, address _token, uint256 _minReturnedTokens, address _beneficiary, string _memo) returns(uint256 netDistributedAmount)
func (_ETHPaymentTerminal *ETHPaymentTerminalSession) UseAllowanceOf(_projectId *big.Int, _amount *big.Int, _currency *big.Int, _token common.Address, _minReturnedTokens *big.Int, _beneficiary common.Address, _memo string) (*types.Transaction, error) {
	return _ETHPaymentTerminal.Contract.UseAllowanceOf(&_ETHPaymentTerminal.TransactOpts, _projectId, _amount, _currency, _token, _minReturnedTokens, _beneficiary, _memo)
}

// UseAllowanceOf is a paid mutator transaction binding the contract method 0xbc8926e9.
//
// Solidity: function useAllowanceOf(uint256 _projectId, uint256 _amount, uint256 _currency, address _token, uint256 _minReturnedTokens, address _beneficiary, string _memo) returns(uint256 netDistributedAmount)
func (_ETHPaymentTerminal *ETHPaymentTerminalTransactorSession) UseAllowanceOf(_projectId *big.Int, _amount *big.Int, _currency *big.Int, _token common.Address, _minReturnedTokens *big.Int, _beneficiary common.Address, _memo string) (*types.Transaction, error) {
	return _ETHPaymentTerminal.Contract.UseAllowanceOf(&_ETHPaymentTerminal.TransactOpts, _projectId, _amount, _currency, _token, _minReturnedTokens, _beneficiary, _memo)
}

// ETHPaymentTerminalAddToBalanceIterator is returned from FilterAddToBalance and is used to iterate over the raw logs and unpacked data for AddToBalance events raised by the ETHPaymentTerminal contract.
type ETHPaymentTerminalAddToBalanceIterator struct {
	Event *ETHPaymentTerminalAddToBalance // Event containing the contract specifics and raw log

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
func (it *ETHPaymentTerminalAddToBalanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ETHPaymentTerminalAddToBalance)
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
		it.Event = new(ETHPaymentTerminalAddToBalance)
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
func (it *ETHPaymentTerminalAddToBalanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ETHPaymentTerminalAddToBalanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ETHPaymentTerminalAddToBalance represents a AddToBalance event raised by the ETHPaymentTerminal contract.
type ETHPaymentTerminalAddToBalance struct {
	ProjectId    *big.Int
	Amount       *big.Int
	RefundedFees *big.Int
	Memo         string
	Metadata     []byte
	Caller       common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAddToBalance is a free log retrieval operation binding the contract event 0x9ecaf7fc3dfffd6867c175d6e684b1f1e3aef019398ba8db2c1ffab4a09db253.
//
// Solidity: event AddToBalance(uint256 indexed projectId, uint256 amount, uint256 refundedFees, string memo, bytes metadata, address caller)
func (_ETHPaymentTerminal *ETHPaymentTerminalFilterer) FilterAddToBalance(opts *bind.FilterOpts, projectId []*big.Int) (*ETHPaymentTerminalAddToBalanceIterator, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}

	logs, sub, err := _ETHPaymentTerminal.contract.FilterLogs(opts, "AddToBalance", projectIdRule)
	if err != nil {
		return nil, err
	}
	return &ETHPaymentTerminalAddToBalanceIterator{contract: _ETHPaymentTerminal.contract, event: "AddToBalance", logs: logs, sub: sub}, nil
}

// WatchAddToBalance is a free log subscription operation binding the contract event 0x9ecaf7fc3dfffd6867c175d6e684b1f1e3aef019398ba8db2c1ffab4a09db253.
//
// Solidity: event AddToBalance(uint256 indexed projectId, uint256 amount, uint256 refundedFees, string memo, bytes metadata, address caller)
func (_ETHPaymentTerminal *ETHPaymentTerminalFilterer) WatchAddToBalance(opts *bind.WatchOpts, sink chan<- *ETHPaymentTerminalAddToBalance, projectId []*big.Int) (event.Subscription, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}

	logs, sub, err := _ETHPaymentTerminal.contract.WatchLogs(opts, "AddToBalance", projectIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ETHPaymentTerminalAddToBalance)
				if err := _ETHPaymentTerminal.contract.UnpackLog(event, "AddToBalance", log); err != nil {
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

// ParseAddToBalance is a log parse operation binding the contract event 0x9ecaf7fc3dfffd6867c175d6e684b1f1e3aef019398ba8db2c1ffab4a09db253.
//
// Solidity: event AddToBalance(uint256 indexed projectId, uint256 amount, uint256 refundedFees, string memo, bytes metadata, address caller)
func (_ETHPaymentTerminal *ETHPaymentTerminalFilterer) ParseAddToBalance(log types.Log) (*ETHPaymentTerminalAddToBalance, error) {
	event := new(ETHPaymentTerminalAddToBalance)
	if err := _ETHPaymentTerminal.contract.UnpackLog(event, "AddToBalance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ETHPaymentTerminalDelegateDidPayIterator is returned from FilterDelegateDidPay and is used to iterate over the raw logs and unpacked data for DelegateDidPay events raised by the ETHPaymentTerminal contract.
type ETHPaymentTerminalDelegateDidPayIterator struct {
	Event *ETHPaymentTerminalDelegateDidPay // Event containing the contract specifics and raw log

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
func (it *ETHPaymentTerminalDelegateDidPayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ETHPaymentTerminalDelegateDidPay)
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
		it.Event = new(ETHPaymentTerminalDelegateDidPay)
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
func (it *ETHPaymentTerminalDelegateDidPayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ETHPaymentTerminalDelegateDidPayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ETHPaymentTerminalDelegateDidPay represents a DelegateDidPay event raised by the ETHPaymentTerminal contract.
type ETHPaymentTerminalDelegateDidPay struct {
	Delegate        common.Address
	Data            JBDidPayData
	DelegatedAmount *big.Int
	Caller          common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDelegateDidPay is a free log retrieval operation binding the contract event 0x16112c26e14efc4be6c690149aa5a1ba75160de245f60d2273e28adb277b9e12.
//
// Solidity: event DelegateDidPay(address indexed delegate, (address,uint256,uint256,(address,uint256,uint256,uint256),(address,uint256,uint256,uint256),uint256,address,bool,string,bytes) data, uint256 delegatedAmount, address caller)
func (_ETHPaymentTerminal *ETHPaymentTerminalFilterer) FilterDelegateDidPay(opts *bind.FilterOpts, delegate []common.Address) (*ETHPaymentTerminalDelegateDidPayIterator, error) {

	var delegateRule []interface{}
	for _, delegateItem := range delegate {
		delegateRule = append(delegateRule, delegateItem)
	}

	logs, sub, err := _ETHPaymentTerminal.contract.FilterLogs(opts, "DelegateDidPay", delegateRule)
	if err != nil {
		return nil, err
	}
	return &ETHPaymentTerminalDelegateDidPayIterator{contract: _ETHPaymentTerminal.contract, event: "DelegateDidPay", logs: logs, sub: sub}, nil
}

// WatchDelegateDidPay is a free log subscription operation binding the contract event 0x16112c26e14efc4be6c690149aa5a1ba75160de245f60d2273e28adb277b9e12.
//
// Solidity: event DelegateDidPay(address indexed delegate, (address,uint256,uint256,(address,uint256,uint256,uint256),(address,uint256,uint256,uint256),uint256,address,bool,string,bytes) data, uint256 delegatedAmount, address caller)
func (_ETHPaymentTerminal *ETHPaymentTerminalFilterer) WatchDelegateDidPay(opts *bind.WatchOpts, sink chan<- *ETHPaymentTerminalDelegateDidPay, delegate []common.Address) (event.Subscription, error) {

	var delegateRule []interface{}
	for _, delegateItem := range delegate {
		delegateRule = append(delegateRule, delegateItem)
	}

	logs, sub, err := _ETHPaymentTerminal.contract.WatchLogs(opts, "DelegateDidPay", delegateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ETHPaymentTerminalDelegateDidPay)
				if err := _ETHPaymentTerminal.contract.UnpackLog(event, "DelegateDidPay", log); err != nil {
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

// ParseDelegateDidPay is a log parse operation binding the contract event 0x16112c26e14efc4be6c690149aa5a1ba75160de245f60d2273e28adb277b9e12.
//
// Solidity: event DelegateDidPay(address indexed delegate, (address,uint256,uint256,(address,uint256,uint256,uint256),(address,uint256,uint256,uint256),uint256,address,bool,string,bytes) data, uint256 delegatedAmount, address caller)
func (_ETHPaymentTerminal *ETHPaymentTerminalFilterer) ParseDelegateDidPay(log types.Log) (*ETHPaymentTerminalDelegateDidPay, error) {
	event := new(ETHPaymentTerminalDelegateDidPay)
	if err := _ETHPaymentTerminal.contract.UnpackLog(event, "DelegateDidPay", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ETHPaymentTerminalDelegateDidRedeemIterator is returned from FilterDelegateDidRedeem and is used to iterate over the raw logs and unpacked data for DelegateDidRedeem events raised by the ETHPaymentTerminal contract.
type ETHPaymentTerminalDelegateDidRedeemIterator struct {
	Event *ETHPaymentTerminalDelegateDidRedeem // Event containing the contract specifics and raw log

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
func (it *ETHPaymentTerminalDelegateDidRedeemIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ETHPaymentTerminalDelegateDidRedeem)
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
		it.Event = new(ETHPaymentTerminalDelegateDidRedeem)
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
func (it *ETHPaymentTerminalDelegateDidRedeemIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ETHPaymentTerminalDelegateDidRedeemIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ETHPaymentTerminalDelegateDidRedeem represents a DelegateDidRedeem event raised by the ETHPaymentTerminal contract.
type ETHPaymentTerminalDelegateDidRedeem struct {
	Delegate        common.Address
	Data            JBDidRedeemData
	DelegatedAmount *big.Int
	Caller          common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDelegateDidRedeem is a free log retrieval operation binding the contract event 0x54b3744c489f40987dd2726ca12131243334e8292f567389f761c5a432d813e4.
//
// Solidity: event DelegateDidRedeem(address indexed delegate, (address,uint256,uint256,uint256,(address,uint256,uint256,uint256),(address,uint256,uint256,uint256),address,string,bytes) data, uint256 delegatedAmount, address caller)
func (_ETHPaymentTerminal *ETHPaymentTerminalFilterer) FilterDelegateDidRedeem(opts *bind.FilterOpts, delegate []common.Address) (*ETHPaymentTerminalDelegateDidRedeemIterator, error) {

	var delegateRule []interface{}
	for _, delegateItem := range delegate {
		delegateRule = append(delegateRule, delegateItem)
	}

	logs, sub, err := _ETHPaymentTerminal.contract.FilterLogs(opts, "DelegateDidRedeem", delegateRule)
	if err != nil {
		return nil, err
	}
	return &ETHPaymentTerminalDelegateDidRedeemIterator{contract: _ETHPaymentTerminal.contract, event: "DelegateDidRedeem", logs: logs, sub: sub}, nil
}

// WatchDelegateDidRedeem is a free log subscription operation binding the contract event 0x54b3744c489f40987dd2726ca12131243334e8292f567389f761c5a432d813e4.
//
// Solidity: event DelegateDidRedeem(address indexed delegate, (address,uint256,uint256,uint256,(address,uint256,uint256,uint256),(address,uint256,uint256,uint256),address,string,bytes) data, uint256 delegatedAmount, address caller)
func (_ETHPaymentTerminal *ETHPaymentTerminalFilterer) WatchDelegateDidRedeem(opts *bind.WatchOpts, sink chan<- *ETHPaymentTerminalDelegateDidRedeem, delegate []common.Address) (event.Subscription, error) {

	var delegateRule []interface{}
	for _, delegateItem := range delegate {
		delegateRule = append(delegateRule, delegateItem)
	}

	logs, sub, err := _ETHPaymentTerminal.contract.WatchLogs(opts, "DelegateDidRedeem", delegateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ETHPaymentTerminalDelegateDidRedeem)
				if err := _ETHPaymentTerminal.contract.UnpackLog(event, "DelegateDidRedeem", log); err != nil {
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

// ParseDelegateDidRedeem is a log parse operation binding the contract event 0x54b3744c489f40987dd2726ca12131243334e8292f567389f761c5a432d813e4.
//
// Solidity: event DelegateDidRedeem(address indexed delegate, (address,uint256,uint256,uint256,(address,uint256,uint256,uint256),(address,uint256,uint256,uint256),address,string,bytes) data, uint256 delegatedAmount, address caller)
func (_ETHPaymentTerminal *ETHPaymentTerminalFilterer) ParseDelegateDidRedeem(log types.Log) (*ETHPaymentTerminalDelegateDidRedeem, error) {
	event := new(ETHPaymentTerminalDelegateDidRedeem)
	if err := _ETHPaymentTerminal.contract.UnpackLog(event, "DelegateDidRedeem", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ETHPaymentTerminalDistributePayoutsIterator is returned from FilterDistributePayouts and is used to iterate over the raw logs and unpacked data for DistributePayouts events raised by the ETHPaymentTerminal contract.
type ETHPaymentTerminalDistributePayoutsIterator struct {
	Event *ETHPaymentTerminalDistributePayouts // Event containing the contract specifics and raw log

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
func (it *ETHPaymentTerminalDistributePayoutsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ETHPaymentTerminalDistributePayouts)
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
		it.Event = new(ETHPaymentTerminalDistributePayouts)
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
func (it *ETHPaymentTerminalDistributePayoutsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ETHPaymentTerminalDistributePayoutsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ETHPaymentTerminalDistributePayouts represents a DistributePayouts event raised by the ETHPaymentTerminal contract.
type ETHPaymentTerminalDistributePayouts struct {
	FundingCycleConfiguration     *big.Int
	FundingCycleNumber            *big.Int
	ProjectId                     *big.Int
	Beneficiary                   common.Address
	Amount                        *big.Int
	DistributedAmount             *big.Int
	Fee                           *big.Int
	BeneficiaryDistributionAmount *big.Int
	Memo                          string
	Caller                        common.Address
	Raw                           types.Log // Blockchain specific contextual infos
}

// FilterDistributePayouts is a free log retrieval operation binding the contract event 0x24352f49df447b14e0e08a323625c663d865ce20c343c4638af12e1dc48aa760.
//
// Solidity: event DistributePayouts(uint256 indexed fundingCycleConfiguration, uint256 indexed fundingCycleNumber, uint256 indexed projectId, address beneficiary, uint256 amount, uint256 distributedAmount, uint256 fee, uint256 beneficiaryDistributionAmount, string memo, address caller)
func (_ETHPaymentTerminal *ETHPaymentTerminalFilterer) FilterDistributePayouts(opts *bind.FilterOpts, fundingCycleConfiguration []*big.Int, fundingCycleNumber []*big.Int, projectId []*big.Int) (*ETHPaymentTerminalDistributePayoutsIterator, error) {

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

	logs, sub, err := _ETHPaymentTerminal.contract.FilterLogs(opts, "DistributePayouts", fundingCycleConfigurationRule, fundingCycleNumberRule, projectIdRule)
	if err != nil {
		return nil, err
	}
	return &ETHPaymentTerminalDistributePayoutsIterator{contract: _ETHPaymentTerminal.contract, event: "DistributePayouts", logs: logs, sub: sub}, nil
}

// WatchDistributePayouts is a free log subscription operation binding the contract event 0x24352f49df447b14e0e08a323625c663d865ce20c343c4638af12e1dc48aa760.
//
// Solidity: event DistributePayouts(uint256 indexed fundingCycleConfiguration, uint256 indexed fundingCycleNumber, uint256 indexed projectId, address beneficiary, uint256 amount, uint256 distributedAmount, uint256 fee, uint256 beneficiaryDistributionAmount, string memo, address caller)
func (_ETHPaymentTerminal *ETHPaymentTerminalFilterer) WatchDistributePayouts(opts *bind.WatchOpts, sink chan<- *ETHPaymentTerminalDistributePayouts, fundingCycleConfiguration []*big.Int, fundingCycleNumber []*big.Int, projectId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _ETHPaymentTerminal.contract.WatchLogs(opts, "DistributePayouts", fundingCycleConfigurationRule, fundingCycleNumberRule, projectIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ETHPaymentTerminalDistributePayouts)
				if err := _ETHPaymentTerminal.contract.UnpackLog(event, "DistributePayouts", log); err != nil {
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

// ParseDistributePayouts is a log parse operation binding the contract event 0x24352f49df447b14e0e08a323625c663d865ce20c343c4638af12e1dc48aa760.
//
// Solidity: event DistributePayouts(uint256 indexed fundingCycleConfiguration, uint256 indexed fundingCycleNumber, uint256 indexed projectId, address beneficiary, uint256 amount, uint256 distributedAmount, uint256 fee, uint256 beneficiaryDistributionAmount, string memo, address caller)
func (_ETHPaymentTerminal *ETHPaymentTerminalFilterer) ParseDistributePayouts(log types.Log) (*ETHPaymentTerminalDistributePayouts, error) {
	event := new(ETHPaymentTerminalDistributePayouts)
	if err := _ETHPaymentTerminal.contract.UnpackLog(event, "DistributePayouts", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ETHPaymentTerminalDistributeToPayoutSplitIterator is returned from FilterDistributeToPayoutSplit and is used to iterate over the raw logs and unpacked data for DistributeToPayoutSplit events raised by the ETHPaymentTerminal contract.
type ETHPaymentTerminalDistributeToPayoutSplitIterator struct {
	Event *ETHPaymentTerminalDistributeToPayoutSplit // Event containing the contract specifics and raw log

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
func (it *ETHPaymentTerminalDistributeToPayoutSplitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ETHPaymentTerminalDistributeToPayoutSplit)
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
		it.Event = new(ETHPaymentTerminalDistributeToPayoutSplit)
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
func (it *ETHPaymentTerminalDistributeToPayoutSplitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ETHPaymentTerminalDistributeToPayoutSplitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ETHPaymentTerminalDistributeToPayoutSplit represents a DistributeToPayoutSplit event raised by the ETHPaymentTerminal contract.
type ETHPaymentTerminalDistributeToPayoutSplit struct {
	ProjectId *big.Int
	Domain    *big.Int
	Group     *big.Int
	Split     JBSplit
	Amount    *big.Int
	Caller    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDistributeToPayoutSplit is a free log retrieval operation binding the contract event 0x2a1f2df21da49f011c6165709ae4b279f8d6d7cffe9043c582352882d8c9698b.
//
// Solidity: event DistributeToPayoutSplit(uint256 indexed projectId, uint256 indexed domain, uint256 indexed group, (bool,bool,uint256,uint256,address,uint256,address) split, uint256 amount, address caller)
func (_ETHPaymentTerminal *ETHPaymentTerminalFilterer) FilterDistributeToPayoutSplit(opts *bind.FilterOpts, projectId []*big.Int, domain []*big.Int, group []*big.Int) (*ETHPaymentTerminalDistributeToPayoutSplitIterator, error) {

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

	logs, sub, err := _ETHPaymentTerminal.contract.FilterLogs(opts, "DistributeToPayoutSplit", projectIdRule, domainRule, groupRule)
	if err != nil {
		return nil, err
	}
	return &ETHPaymentTerminalDistributeToPayoutSplitIterator{contract: _ETHPaymentTerminal.contract, event: "DistributeToPayoutSplit", logs: logs, sub: sub}, nil
}

// WatchDistributeToPayoutSplit is a free log subscription operation binding the contract event 0x2a1f2df21da49f011c6165709ae4b279f8d6d7cffe9043c582352882d8c9698b.
//
// Solidity: event DistributeToPayoutSplit(uint256 indexed projectId, uint256 indexed domain, uint256 indexed group, (bool,bool,uint256,uint256,address,uint256,address) split, uint256 amount, address caller)
func (_ETHPaymentTerminal *ETHPaymentTerminalFilterer) WatchDistributeToPayoutSplit(opts *bind.WatchOpts, sink chan<- *ETHPaymentTerminalDistributeToPayoutSplit, projectId []*big.Int, domain []*big.Int, group []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _ETHPaymentTerminal.contract.WatchLogs(opts, "DistributeToPayoutSplit", projectIdRule, domainRule, groupRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ETHPaymentTerminalDistributeToPayoutSplit)
				if err := _ETHPaymentTerminal.contract.UnpackLog(event, "DistributeToPayoutSplit", log); err != nil {
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

// ParseDistributeToPayoutSplit is a log parse operation binding the contract event 0x2a1f2df21da49f011c6165709ae4b279f8d6d7cffe9043c582352882d8c9698b.
//
// Solidity: event DistributeToPayoutSplit(uint256 indexed projectId, uint256 indexed domain, uint256 indexed group, (bool,bool,uint256,uint256,address,uint256,address) split, uint256 amount, address caller)
func (_ETHPaymentTerminal *ETHPaymentTerminalFilterer) ParseDistributeToPayoutSplit(log types.Log) (*ETHPaymentTerminalDistributeToPayoutSplit, error) {
	event := new(ETHPaymentTerminalDistributeToPayoutSplit)
	if err := _ETHPaymentTerminal.contract.UnpackLog(event, "DistributeToPayoutSplit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ETHPaymentTerminalHoldFeeIterator is returned from FilterHoldFee and is used to iterate over the raw logs and unpacked data for HoldFee events raised by the ETHPaymentTerminal contract.
type ETHPaymentTerminalHoldFeeIterator struct {
	Event *ETHPaymentTerminalHoldFee // Event containing the contract specifics and raw log

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
func (it *ETHPaymentTerminalHoldFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ETHPaymentTerminalHoldFee)
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
		it.Event = new(ETHPaymentTerminalHoldFee)
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
func (it *ETHPaymentTerminalHoldFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ETHPaymentTerminalHoldFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ETHPaymentTerminalHoldFee represents a HoldFee event raised by the ETHPaymentTerminal contract.
type ETHPaymentTerminalHoldFee struct {
	ProjectId   *big.Int
	Amount      *big.Int
	Fee         *big.Int
	FeeDiscount *big.Int
	Beneficiary common.Address
	Caller      common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterHoldFee is a free log retrieval operation binding the contract event 0x77813be0661650ddc1a5193ff2837df4162b251cb432651e2c060c3fc39756be.
//
// Solidity: event HoldFee(uint256 indexed projectId, uint256 indexed amount, uint256 indexed fee, uint256 feeDiscount, address beneficiary, address caller)
func (_ETHPaymentTerminal *ETHPaymentTerminalFilterer) FilterHoldFee(opts *bind.FilterOpts, projectId []*big.Int, amount []*big.Int, fee []*big.Int) (*ETHPaymentTerminalHoldFeeIterator, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}
	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}

	logs, sub, err := _ETHPaymentTerminal.contract.FilterLogs(opts, "HoldFee", projectIdRule, amountRule, feeRule)
	if err != nil {
		return nil, err
	}
	return &ETHPaymentTerminalHoldFeeIterator{contract: _ETHPaymentTerminal.contract, event: "HoldFee", logs: logs, sub: sub}, nil
}

// WatchHoldFee is a free log subscription operation binding the contract event 0x77813be0661650ddc1a5193ff2837df4162b251cb432651e2c060c3fc39756be.
//
// Solidity: event HoldFee(uint256 indexed projectId, uint256 indexed amount, uint256 indexed fee, uint256 feeDiscount, address beneficiary, address caller)
func (_ETHPaymentTerminal *ETHPaymentTerminalFilterer) WatchHoldFee(opts *bind.WatchOpts, sink chan<- *ETHPaymentTerminalHoldFee, projectId []*big.Int, amount []*big.Int, fee []*big.Int) (event.Subscription, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}
	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}

	logs, sub, err := _ETHPaymentTerminal.contract.WatchLogs(opts, "HoldFee", projectIdRule, amountRule, feeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ETHPaymentTerminalHoldFee)
				if err := _ETHPaymentTerminal.contract.UnpackLog(event, "HoldFee", log); err != nil {
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

// ParseHoldFee is a log parse operation binding the contract event 0x77813be0661650ddc1a5193ff2837df4162b251cb432651e2c060c3fc39756be.
//
// Solidity: event HoldFee(uint256 indexed projectId, uint256 indexed amount, uint256 indexed fee, uint256 feeDiscount, address beneficiary, address caller)
func (_ETHPaymentTerminal *ETHPaymentTerminalFilterer) ParseHoldFee(log types.Log) (*ETHPaymentTerminalHoldFee, error) {
	event := new(ETHPaymentTerminalHoldFee)
	if err := _ETHPaymentTerminal.contract.UnpackLog(event, "HoldFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ETHPaymentTerminalMigrateIterator is returned from FilterMigrate and is used to iterate over the raw logs and unpacked data for Migrate events raised by the ETHPaymentTerminal contract.
type ETHPaymentTerminalMigrateIterator struct {
	Event *ETHPaymentTerminalMigrate // Event containing the contract specifics and raw log

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
func (it *ETHPaymentTerminalMigrateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ETHPaymentTerminalMigrate)
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
		it.Event = new(ETHPaymentTerminalMigrate)
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
func (it *ETHPaymentTerminalMigrateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ETHPaymentTerminalMigrateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ETHPaymentTerminalMigrate represents a Migrate event raised by the ETHPaymentTerminal contract.
type ETHPaymentTerminalMigrate struct {
	ProjectId *big.Int
	To        common.Address
	Amount    *big.Int
	Caller    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMigrate is a free log retrieval operation binding the contract event 0xa7519e5f94697b7f53e97c5eb46a0c730a296ab686ab8fd333835c5f735784eb.
//
// Solidity: event Migrate(uint256 indexed projectId, address indexed to, uint256 amount, address caller)
func (_ETHPaymentTerminal *ETHPaymentTerminalFilterer) FilterMigrate(opts *bind.FilterOpts, projectId []*big.Int, to []common.Address) (*ETHPaymentTerminalMigrateIterator, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ETHPaymentTerminal.contract.FilterLogs(opts, "Migrate", projectIdRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ETHPaymentTerminalMigrateIterator{contract: _ETHPaymentTerminal.contract, event: "Migrate", logs: logs, sub: sub}, nil
}

// WatchMigrate is a free log subscription operation binding the contract event 0xa7519e5f94697b7f53e97c5eb46a0c730a296ab686ab8fd333835c5f735784eb.
//
// Solidity: event Migrate(uint256 indexed projectId, address indexed to, uint256 amount, address caller)
func (_ETHPaymentTerminal *ETHPaymentTerminalFilterer) WatchMigrate(opts *bind.WatchOpts, sink chan<- *ETHPaymentTerminalMigrate, projectId []*big.Int, to []common.Address) (event.Subscription, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ETHPaymentTerminal.contract.WatchLogs(opts, "Migrate", projectIdRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ETHPaymentTerminalMigrate)
				if err := _ETHPaymentTerminal.contract.UnpackLog(event, "Migrate", log); err != nil {
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

// ParseMigrate is a log parse operation binding the contract event 0xa7519e5f94697b7f53e97c5eb46a0c730a296ab686ab8fd333835c5f735784eb.
//
// Solidity: event Migrate(uint256 indexed projectId, address indexed to, uint256 amount, address caller)
func (_ETHPaymentTerminal *ETHPaymentTerminalFilterer) ParseMigrate(log types.Log) (*ETHPaymentTerminalMigrate, error) {
	event := new(ETHPaymentTerminalMigrate)
	if err := _ETHPaymentTerminal.contract.UnpackLog(event, "Migrate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ETHPaymentTerminalOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ETHPaymentTerminal contract.
type ETHPaymentTerminalOwnershipTransferredIterator struct {
	Event *ETHPaymentTerminalOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ETHPaymentTerminalOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ETHPaymentTerminalOwnershipTransferred)
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
		it.Event = new(ETHPaymentTerminalOwnershipTransferred)
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
func (it *ETHPaymentTerminalOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ETHPaymentTerminalOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ETHPaymentTerminalOwnershipTransferred represents a OwnershipTransferred event raised by the ETHPaymentTerminal contract.
type ETHPaymentTerminalOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ETHPaymentTerminal *ETHPaymentTerminalFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ETHPaymentTerminalOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ETHPaymentTerminal.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ETHPaymentTerminalOwnershipTransferredIterator{contract: _ETHPaymentTerminal.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ETHPaymentTerminal *ETHPaymentTerminalFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ETHPaymentTerminalOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ETHPaymentTerminal.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ETHPaymentTerminalOwnershipTransferred)
				if err := _ETHPaymentTerminal.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ETHPaymentTerminal *ETHPaymentTerminalFilterer) ParseOwnershipTransferred(log types.Log) (*ETHPaymentTerminalOwnershipTransferred, error) {
	event := new(ETHPaymentTerminalOwnershipTransferred)
	if err := _ETHPaymentTerminal.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ETHPaymentTerminalPayIterator is returned from FilterPay and is used to iterate over the raw logs and unpacked data for Pay events raised by the ETHPaymentTerminal contract.
type ETHPaymentTerminalPayIterator struct {
	Event *ETHPaymentTerminalPay // Event containing the contract specifics and raw log

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
func (it *ETHPaymentTerminalPayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ETHPaymentTerminalPay)
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
		it.Event = new(ETHPaymentTerminalPay)
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
func (it *ETHPaymentTerminalPayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ETHPaymentTerminalPayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ETHPaymentTerminalPay represents a Pay event raised by the ETHPaymentTerminal contract.
type ETHPaymentTerminalPay struct {
	FundingCycleConfiguration *big.Int
	FundingCycleNumber        *big.Int
	ProjectId                 *big.Int
	Payer                     common.Address
	Beneficiary               common.Address
	Amount                    *big.Int
	BeneficiaryTokenCount     *big.Int
	Memo                      string
	Metadata                  []byte
	Caller                    common.Address
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterPay is a free log retrieval operation binding the contract event 0x133161f1c9161488f777ab9a26aae91d47c0d9a3fafb398960f138db02c73797.
//
// Solidity: event Pay(uint256 indexed fundingCycleConfiguration, uint256 indexed fundingCycleNumber, uint256 indexed projectId, address payer, address beneficiary, uint256 amount, uint256 beneficiaryTokenCount, string memo, bytes metadata, address caller)
func (_ETHPaymentTerminal *ETHPaymentTerminalFilterer) FilterPay(opts *bind.FilterOpts, fundingCycleConfiguration []*big.Int, fundingCycleNumber []*big.Int, projectId []*big.Int) (*ETHPaymentTerminalPayIterator, error) {

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

	logs, sub, err := _ETHPaymentTerminal.contract.FilterLogs(opts, "Pay", fundingCycleConfigurationRule, fundingCycleNumberRule, projectIdRule)
	if err != nil {
		return nil, err
	}
	return &ETHPaymentTerminalPayIterator{contract: _ETHPaymentTerminal.contract, event: "Pay", logs: logs, sub: sub}, nil
}

// WatchPay is a free log subscription operation binding the contract event 0x133161f1c9161488f777ab9a26aae91d47c0d9a3fafb398960f138db02c73797.
//
// Solidity: event Pay(uint256 indexed fundingCycleConfiguration, uint256 indexed fundingCycleNumber, uint256 indexed projectId, address payer, address beneficiary, uint256 amount, uint256 beneficiaryTokenCount, string memo, bytes metadata, address caller)
func (_ETHPaymentTerminal *ETHPaymentTerminalFilterer) WatchPay(opts *bind.WatchOpts, sink chan<- *ETHPaymentTerminalPay, fundingCycleConfiguration []*big.Int, fundingCycleNumber []*big.Int, projectId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _ETHPaymentTerminal.contract.WatchLogs(opts, "Pay", fundingCycleConfigurationRule, fundingCycleNumberRule, projectIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ETHPaymentTerminalPay)
				if err := _ETHPaymentTerminal.contract.UnpackLog(event, "Pay", log); err != nil {
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

// ParsePay is a log parse operation binding the contract event 0x133161f1c9161488f777ab9a26aae91d47c0d9a3fafb398960f138db02c73797.
//
// Solidity: event Pay(uint256 indexed fundingCycleConfiguration, uint256 indexed fundingCycleNumber, uint256 indexed projectId, address payer, address beneficiary, uint256 amount, uint256 beneficiaryTokenCount, string memo, bytes metadata, address caller)
func (_ETHPaymentTerminal *ETHPaymentTerminalFilterer) ParsePay(log types.Log) (*ETHPaymentTerminalPay, error) {
	event := new(ETHPaymentTerminalPay)
	if err := _ETHPaymentTerminal.contract.UnpackLog(event, "Pay", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ETHPaymentTerminalProcessFeeIterator is returned from FilterProcessFee and is used to iterate over the raw logs and unpacked data for ProcessFee events raised by the ETHPaymentTerminal contract.
type ETHPaymentTerminalProcessFeeIterator struct {
	Event *ETHPaymentTerminalProcessFee // Event containing the contract specifics and raw log

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
func (it *ETHPaymentTerminalProcessFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ETHPaymentTerminalProcessFee)
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
		it.Event = new(ETHPaymentTerminalProcessFee)
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
func (it *ETHPaymentTerminalProcessFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ETHPaymentTerminalProcessFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ETHPaymentTerminalProcessFee represents a ProcessFee event raised by the ETHPaymentTerminal contract.
type ETHPaymentTerminalProcessFee struct {
	ProjectId   *big.Int
	Amount      *big.Int
	WasHeld     bool
	Beneficiary common.Address
	Caller      common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterProcessFee is a free log retrieval operation binding the contract event 0xcf0c92a2c6d7c42f488326b0cb900104b99984b6b218db81cd29371364a35251.
//
// Solidity: event ProcessFee(uint256 indexed projectId, uint256 indexed amount, bool indexed wasHeld, address beneficiary, address caller)
func (_ETHPaymentTerminal *ETHPaymentTerminalFilterer) FilterProcessFee(opts *bind.FilterOpts, projectId []*big.Int, amount []*big.Int, wasHeld []bool) (*ETHPaymentTerminalProcessFeeIterator, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}
	var wasHeldRule []interface{}
	for _, wasHeldItem := range wasHeld {
		wasHeldRule = append(wasHeldRule, wasHeldItem)
	}

	logs, sub, err := _ETHPaymentTerminal.contract.FilterLogs(opts, "ProcessFee", projectIdRule, amountRule, wasHeldRule)
	if err != nil {
		return nil, err
	}
	return &ETHPaymentTerminalProcessFeeIterator{contract: _ETHPaymentTerminal.contract, event: "ProcessFee", logs: logs, sub: sub}, nil
}

// WatchProcessFee is a free log subscription operation binding the contract event 0xcf0c92a2c6d7c42f488326b0cb900104b99984b6b218db81cd29371364a35251.
//
// Solidity: event ProcessFee(uint256 indexed projectId, uint256 indexed amount, bool indexed wasHeld, address beneficiary, address caller)
func (_ETHPaymentTerminal *ETHPaymentTerminalFilterer) WatchProcessFee(opts *bind.WatchOpts, sink chan<- *ETHPaymentTerminalProcessFee, projectId []*big.Int, amount []*big.Int, wasHeld []bool) (event.Subscription, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}
	var wasHeldRule []interface{}
	for _, wasHeldItem := range wasHeld {
		wasHeldRule = append(wasHeldRule, wasHeldItem)
	}

	logs, sub, err := _ETHPaymentTerminal.contract.WatchLogs(opts, "ProcessFee", projectIdRule, amountRule, wasHeldRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ETHPaymentTerminalProcessFee)
				if err := _ETHPaymentTerminal.contract.UnpackLog(event, "ProcessFee", log); err != nil {
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

// ParseProcessFee is a log parse operation binding the contract event 0xcf0c92a2c6d7c42f488326b0cb900104b99984b6b218db81cd29371364a35251.
//
// Solidity: event ProcessFee(uint256 indexed projectId, uint256 indexed amount, bool indexed wasHeld, address beneficiary, address caller)
func (_ETHPaymentTerminal *ETHPaymentTerminalFilterer) ParseProcessFee(log types.Log) (*ETHPaymentTerminalProcessFee, error) {
	event := new(ETHPaymentTerminalProcessFee)
	if err := _ETHPaymentTerminal.contract.UnpackLog(event, "ProcessFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ETHPaymentTerminalRedeemTokensIterator is returned from FilterRedeemTokens and is used to iterate over the raw logs and unpacked data for RedeemTokens events raised by the ETHPaymentTerminal contract.
type ETHPaymentTerminalRedeemTokensIterator struct {
	Event *ETHPaymentTerminalRedeemTokens // Event containing the contract specifics and raw log

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
func (it *ETHPaymentTerminalRedeemTokensIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ETHPaymentTerminalRedeemTokens)
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
		it.Event = new(ETHPaymentTerminalRedeemTokens)
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
func (it *ETHPaymentTerminalRedeemTokensIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ETHPaymentTerminalRedeemTokensIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ETHPaymentTerminalRedeemTokens represents a RedeemTokens event raised by the ETHPaymentTerminal contract.
type ETHPaymentTerminalRedeemTokens struct {
	FundingCycleConfiguration *big.Int
	FundingCycleNumber        *big.Int
	ProjectId                 *big.Int
	Holder                    common.Address
	Beneficiary               common.Address
	TokenCount                *big.Int
	ReclaimedAmount           *big.Int
	Memo                      string
	Metadata                  []byte
	Caller                    common.Address
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterRedeemTokens is a free log retrieval operation binding the contract event 0x2be10f2a0203c77d0fcaa9fd6484a8a1d6904de31cd820587f60c1c8c338c814.
//
// Solidity: event RedeemTokens(uint256 indexed fundingCycleConfiguration, uint256 indexed fundingCycleNumber, uint256 indexed projectId, address holder, address beneficiary, uint256 tokenCount, uint256 reclaimedAmount, string memo, bytes metadata, address caller)
func (_ETHPaymentTerminal *ETHPaymentTerminalFilterer) FilterRedeemTokens(opts *bind.FilterOpts, fundingCycleConfiguration []*big.Int, fundingCycleNumber []*big.Int, projectId []*big.Int) (*ETHPaymentTerminalRedeemTokensIterator, error) {

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

	logs, sub, err := _ETHPaymentTerminal.contract.FilterLogs(opts, "RedeemTokens", fundingCycleConfigurationRule, fundingCycleNumberRule, projectIdRule)
	if err != nil {
		return nil, err
	}
	return &ETHPaymentTerminalRedeemTokensIterator{contract: _ETHPaymentTerminal.contract, event: "RedeemTokens", logs: logs, sub: sub}, nil
}

// WatchRedeemTokens is a free log subscription operation binding the contract event 0x2be10f2a0203c77d0fcaa9fd6484a8a1d6904de31cd820587f60c1c8c338c814.
//
// Solidity: event RedeemTokens(uint256 indexed fundingCycleConfiguration, uint256 indexed fundingCycleNumber, uint256 indexed projectId, address holder, address beneficiary, uint256 tokenCount, uint256 reclaimedAmount, string memo, bytes metadata, address caller)
func (_ETHPaymentTerminal *ETHPaymentTerminalFilterer) WatchRedeemTokens(opts *bind.WatchOpts, sink chan<- *ETHPaymentTerminalRedeemTokens, fundingCycleConfiguration []*big.Int, fundingCycleNumber []*big.Int, projectId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _ETHPaymentTerminal.contract.WatchLogs(opts, "RedeemTokens", fundingCycleConfigurationRule, fundingCycleNumberRule, projectIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ETHPaymentTerminalRedeemTokens)
				if err := _ETHPaymentTerminal.contract.UnpackLog(event, "RedeemTokens", log); err != nil {
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

// ParseRedeemTokens is a log parse operation binding the contract event 0x2be10f2a0203c77d0fcaa9fd6484a8a1d6904de31cd820587f60c1c8c338c814.
//
// Solidity: event RedeemTokens(uint256 indexed fundingCycleConfiguration, uint256 indexed fundingCycleNumber, uint256 indexed projectId, address holder, address beneficiary, uint256 tokenCount, uint256 reclaimedAmount, string memo, bytes metadata, address caller)
func (_ETHPaymentTerminal *ETHPaymentTerminalFilterer) ParseRedeemTokens(log types.Log) (*ETHPaymentTerminalRedeemTokens, error) {
	event := new(ETHPaymentTerminalRedeemTokens)
	if err := _ETHPaymentTerminal.contract.UnpackLog(event, "RedeemTokens", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ETHPaymentTerminalRefundHeldFeesIterator is returned from FilterRefundHeldFees and is used to iterate over the raw logs and unpacked data for RefundHeldFees events raised by the ETHPaymentTerminal contract.
type ETHPaymentTerminalRefundHeldFeesIterator struct {
	Event *ETHPaymentTerminalRefundHeldFees // Event containing the contract specifics and raw log

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
func (it *ETHPaymentTerminalRefundHeldFeesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ETHPaymentTerminalRefundHeldFees)
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
		it.Event = new(ETHPaymentTerminalRefundHeldFees)
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
func (it *ETHPaymentTerminalRefundHeldFeesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ETHPaymentTerminalRefundHeldFeesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ETHPaymentTerminalRefundHeldFees represents a RefundHeldFees event raised by the ETHPaymentTerminal contract.
type ETHPaymentTerminalRefundHeldFees struct {
	ProjectId      *big.Int
	Amount         *big.Int
	RefundedFees   *big.Int
	LeftoverAmount *big.Int
	Caller         common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRefundHeldFees is a free log retrieval operation binding the contract event 0x59860d79d97c1fce2be7f987915c631471f4b08f671200463cc40a3380194ffb.
//
// Solidity: event RefundHeldFees(uint256 indexed projectId, uint256 indexed amount, uint256 indexed refundedFees, uint256 leftoverAmount, address caller)
func (_ETHPaymentTerminal *ETHPaymentTerminalFilterer) FilterRefundHeldFees(opts *bind.FilterOpts, projectId []*big.Int, amount []*big.Int, refundedFees []*big.Int) (*ETHPaymentTerminalRefundHeldFeesIterator, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}
	var refundedFeesRule []interface{}
	for _, refundedFeesItem := range refundedFees {
		refundedFeesRule = append(refundedFeesRule, refundedFeesItem)
	}

	logs, sub, err := _ETHPaymentTerminal.contract.FilterLogs(opts, "RefundHeldFees", projectIdRule, amountRule, refundedFeesRule)
	if err != nil {
		return nil, err
	}
	return &ETHPaymentTerminalRefundHeldFeesIterator{contract: _ETHPaymentTerminal.contract, event: "RefundHeldFees", logs: logs, sub: sub}, nil
}

// WatchRefundHeldFees is a free log subscription operation binding the contract event 0x59860d79d97c1fce2be7f987915c631471f4b08f671200463cc40a3380194ffb.
//
// Solidity: event RefundHeldFees(uint256 indexed projectId, uint256 indexed amount, uint256 indexed refundedFees, uint256 leftoverAmount, address caller)
func (_ETHPaymentTerminal *ETHPaymentTerminalFilterer) WatchRefundHeldFees(opts *bind.WatchOpts, sink chan<- *ETHPaymentTerminalRefundHeldFees, projectId []*big.Int, amount []*big.Int, refundedFees []*big.Int) (event.Subscription, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}
	var refundedFeesRule []interface{}
	for _, refundedFeesItem := range refundedFees {
		refundedFeesRule = append(refundedFeesRule, refundedFeesItem)
	}

	logs, sub, err := _ETHPaymentTerminal.contract.WatchLogs(opts, "RefundHeldFees", projectIdRule, amountRule, refundedFeesRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ETHPaymentTerminalRefundHeldFees)
				if err := _ETHPaymentTerminal.contract.UnpackLog(event, "RefundHeldFees", log); err != nil {
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

// ParseRefundHeldFees is a log parse operation binding the contract event 0x59860d79d97c1fce2be7f987915c631471f4b08f671200463cc40a3380194ffb.
//
// Solidity: event RefundHeldFees(uint256 indexed projectId, uint256 indexed amount, uint256 indexed refundedFees, uint256 leftoverAmount, address caller)
func (_ETHPaymentTerminal *ETHPaymentTerminalFilterer) ParseRefundHeldFees(log types.Log) (*ETHPaymentTerminalRefundHeldFees, error) {
	event := new(ETHPaymentTerminalRefundHeldFees)
	if err := _ETHPaymentTerminal.contract.UnpackLog(event, "RefundHeldFees", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ETHPaymentTerminalSetFeeIterator is returned from FilterSetFee and is used to iterate over the raw logs and unpacked data for SetFee events raised by the ETHPaymentTerminal contract.
type ETHPaymentTerminalSetFeeIterator struct {
	Event *ETHPaymentTerminalSetFee // Event containing the contract specifics and raw log

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
func (it *ETHPaymentTerminalSetFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ETHPaymentTerminalSetFee)
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
		it.Event = new(ETHPaymentTerminalSetFee)
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
func (it *ETHPaymentTerminalSetFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ETHPaymentTerminalSetFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ETHPaymentTerminalSetFee represents a SetFee event raised by the ETHPaymentTerminal contract.
type ETHPaymentTerminalSetFee struct {
	Fee    *big.Int
	Caller common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetFee is a free log retrieval operation binding the contract event 0xd7414e590e1cb532989ab2a34c8f4c2c17f7ab6f006efeeaef2e87cd5008c202.
//
// Solidity: event SetFee(uint256 fee, address caller)
func (_ETHPaymentTerminal *ETHPaymentTerminalFilterer) FilterSetFee(opts *bind.FilterOpts) (*ETHPaymentTerminalSetFeeIterator, error) {

	logs, sub, err := _ETHPaymentTerminal.contract.FilterLogs(opts, "SetFee")
	if err != nil {
		return nil, err
	}
	return &ETHPaymentTerminalSetFeeIterator{contract: _ETHPaymentTerminal.contract, event: "SetFee", logs: logs, sub: sub}, nil
}

// WatchSetFee is a free log subscription operation binding the contract event 0xd7414e590e1cb532989ab2a34c8f4c2c17f7ab6f006efeeaef2e87cd5008c202.
//
// Solidity: event SetFee(uint256 fee, address caller)
func (_ETHPaymentTerminal *ETHPaymentTerminalFilterer) WatchSetFee(opts *bind.WatchOpts, sink chan<- *ETHPaymentTerminalSetFee) (event.Subscription, error) {

	logs, sub, err := _ETHPaymentTerminal.contract.WatchLogs(opts, "SetFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ETHPaymentTerminalSetFee)
				if err := _ETHPaymentTerminal.contract.UnpackLog(event, "SetFee", log); err != nil {
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

// ParseSetFee is a log parse operation binding the contract event 0xd7414e590e1cb532989ab2a34c8f4c2c17f7ab6f006efeeaef2e87cd5008c202.
//
// Solidity: event SetFee(uint256 fee, address caller)
func (_ETHPaymentTerminal *ETHPaymentTerminalFilterer) ParseSetFee(log types.Log) (*ETHPaymentTerminalSetFee, error) {
	event := new(ETHPaymentTerminalSetFee)
	if err := _ETHPaymentTerminal.contract.UnpackLog(event, "SetFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ETHPaymentTerminalSetFeeGaugeIterator is returned from FilterSetFeeGauge and is used to iterate over the raw logs and unpacked data for SetFeeGauge events raised by the ETHPaymentTerminal contract.
type ETHPaymentTerminalSetFeeGaugeIterator struct {
	Event *ETHPaymentTerminalSetFeeGauge // Event containing the contract specifics and raw log

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
func (it *ETHPaymentTerminalSetFeeGaugeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ETHPaymentTerminalSetFeeGauge)
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
		it.Event = new(ETHPaymentTerminalSetFeeGauge)
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
func (it *ETHPaymentTerminalSetFeeGaugeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ETHPaymentTerminalSetFeeGaugeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ETHPaymentTerminalSetFeeGauge represents a SetFeeGauge event raised by the ETHPaymentTerminal contract.
type ETHPaymentTerminalSetFeeGauge struct {
	FeeGauge common.Address
	Caller   common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetFeeGauge is a free log retrieval operation binding the contract event 0x0a9a80fe9716605b3e52abb3d792d6a4e7816d6afc02a5a4ef023081feaf9f60.
//
// Solidity: event SetFeeGauge(address indexed feeGauge, address caller)
func (_ETHPaymentTerminal *ETHPaymentTerminalFilterer) FilterSetFeeGauge(opts *bind.FilterOpts, feeGauge []common.Address) (*ETHPaymentTerminalSetFeeGaugeIterator, error) {

	var feeGaugeRule []interface{}
	for _, feeGaugeItem := range feeGauge {
		feeGaugeRule = append(feeGaugeRule, feeGaugeItem)
	}

	logs, sub, err := _ETHPaymentTerminal.contract.FilterLogs(opts, "SetFeeGauge", feeGaugeRule)
	if err != nil {
		return nil, err
	}
	return &ETHPaymentTerminalSetFeeGaugeIterator{contract: _ETHPaymentTerminal.contract, event: "SetFeeGauge", logs: logs, sub: sub}, nil
}

// WatchSetFeeGauge is a free log subscription operation binding the contract event 0x0a9a80fe9716605b3e52abb3d792d6a4e7816d6afc02a5a4ef023081feaf9f60.
//
// Solidity: event SetFeeGauge(address indexed feeGauge, address caller)
func (_ETHPaymentTerminal *ETHPaymentTerminalFilterer) WatchSetFeeGauge(opts *bind.WatchOpts, sink chan<- *ETHPaymentTerminalSetFeeGauge, feeGauge []common.Address) (event.Subscription, error) {

	var feeGaugeRule []interface{}
	for _, feeGaugeItem := range feeGauge {
		feeGaugeRule = append(feeGaugeRule, feeGaugeItem)
	}

	logs, sub, err := _ETHPaymentTerminal.contract.WatchLogs(opts, "SetFeeGauge", feeGaugeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ETHPaymentTerminalSetFeeGauge)
				if err := _ETHPaymentTerminal.contract.UnpackLog(event, "SetFeeGauge", log); err != nil {
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

// ParseSetFeeGauge is a log parse operation binding the contract event 0x0a9a80fe9716605b3e52abb3d792d6a4e7816d6afc02a5a4ef023081feaf9f60.
//
// Solidity: event SetFeeGauge(address indexed feeGauge, address caller)
func (_ETHPaymentTerminal *ETHPaymentTerminalFilterer) ParseSetFeeGauge(log types.Log) (*ETHPaymentTerminalSetFeeGauge, error) {
	event := new(ETHPaymentTerminalSetFeeGauge)
	if err := _ETHPaymentTerminal.contract.UnpackLog(event, "SetFeeGauge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ETHPaymentTerminalSetFeelessAddressIterator is returned from FilterSetFeelessAddress and is used to iterate over the raw logs and unpacked data for SetFeelessAddress events raised by the ETHPaymentTerminal contract.
type ETHPaymentTerminalSetFeelessAddressIterator struct {
	Event *ETHPaymentTerminalSetFeelessAddress // Event containing the contract specifics and raw log

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
func (it *ETHPaymentTerminalSetFeelessAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ETHPaymentTerminalSetFeelessAddress)
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
		it.Event = new(ETHPaymentTerminalSetFeelessAddress)
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
func (it *ETHPaymentTerminalSetFeelessAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ETHPaymentTerminalSetFeelessAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ETHPaymentTerminalSetFeelessAddress represents a SetFeelessAddress event raised by the ETHPaymentTerminal contract.
type ETHPaymentTerminalSetFeelessAddress struct {
	Addrs  common.Address
	Flag   bool
	Caller common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetFeelessAddress is a free log retrieval operation binding the contract event 0xa2653e25a502c023a5830d0de847ef6f458387865b1f4f575d7594f9f2c0d71e.
//
// Solidity: event SetFeelessAddress(address indexed addrs, bool indexed flag, address caller)
func (_ETHPaymentTerminal *ETHPaymentTerminalFilterer) FilterSetFeelessAddress(opts *bind.FilterOpts, addrs []common.Address, flag []bool) (*ETHPaymentTerminalSetFeelessAddressIterator, error) {

	var addrsRule []interface{}
	for _, addrsItem := range addrs {
		addrsRule = append(addrsRule, addrsItem)
	}
	var flagRule []interface{}
	for _, flagItem := range flag {
		flagRule = append(flagRule, flagItem)
	}

	logs, sub, err := _ETHPaymentTerminal.contract.FilterLogs(opts, "SetFeelessAddress", addrsRule, flagRule)
	if err != nil {
		return nil, err
	}
	return &ETHPaymentTerminalSetFeelessAddressIterator{contract: _ETHPaymentTerminal.contract, event: "SetFeelessAddress", logs: logs, sub: sub}, nil
}

// WatchSetFeelessAddress is a free log subscription operation binding the contract event 0xa2653e25a502c023a5830d0de847ef6f458387865b1f4f575d7594f9f2c0d71e.
//
// Solidity: event SetFeelessAddress(address indexed addrs, bool indexed flag, address caller)
func (_ETHPaymentTerminal *ETHPaymentTerminalFilterer) WatchSetFeelessAddress(opts *bind.WatchOpts, sink chan<- *ETHPaymentTerminalSetFeelessAddress, addrs []common.Address, flag []bool) (event.Subscription, error) {

	var addrsRule []interface{}
	for _, addrsItem := range addrs {
		addrsRule = append(addrsRule, addrsItem)
	}
	var flagRule []interface{}
	for _, flagItem := range flag {
		flagRule = append(flagRule, flagItem)
	}

	logs, sub, err := _ETHPaymentTerminal.contract.WatchLogs(opts, "SetFeelessAddress", addrsRule, flagRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ETHPaymentTerminalSetFeelessAddress)
				if err := _ETHPaymentTerminal.contract.UnpackLog(event, "SetFeelessAddress", log); err != nil {
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

// ParseSetFeelessAddress is a log parse operation binding the contract event 0xa2653e25a502c023a5830d0de847ef6f458387865b1f4f575d7594f9f2c0d71e.
//
// Solidity: event SetFeelessAddress(address indexed addrs, bool indexed flag, address caller)
func (_ETHPaymentTerminal *ETHPaymentTerminalFilterer) ParseSetFeelessAddress(log types.Log) (*ETHPaymentTerminalSetFeelessAddress, error) {
	event := new(ETHPaymentTerminalSetFeelessAddress)
	if err := _ETHPaymentTerminal.contract.UnpackLog(event, "SetFeelessAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ETHPaymentTerminalUseAllowanceIterator is returned from FilterUseAllowance and is used to iterate over the raw logs and unpacked data for UseAllowance events raised by the ETHPaymentTerminal contract.
type ETHPaymentTerminalUseAllowanceIterator struct {
	Event *ETHPaymentTerminalUseAllowance // Event containing the contract specifics and raw log

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
func (it *ETHPaymentTerminalUseAllowanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ETHPaymentTerminalUseAllowance)
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
		it.Event = new(ETHPaymentTerminalUseAllowance)
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
func (it *ETHPaymentTerminalUseAllowanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ETHPaymentTerminalUseAllowanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ETHPaymentTerminalUseAllowance represents a UseAllowance event raised by the ETHPaymentTerminal contract.
type ETHPaymentTerminalUseAllowance struct {
	FundingCycleConfiguration *big.Int
	FundingCycleNumber        *big.Int
	ProjectId                 *big.Int
	Beneficiary               common.Address
	Amount                    *big.Int
	DistributedAmount         *big.Int
	NetDistributedamount      *big.Int
	Memo                      string
	Caller                    common.Address
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterUseAllowance is a free log retrieval operation binding the contract event 0x8657a0c05a68a912c23c1bd00124afaa8c669063b046bd9bfd22b21d573c5e6d.
//
// Solidity: event UseAllowance(uint256 indexed fundingCycleConfiguration, uint256 indexed fundingCycleNumber, uint256 indexed projectId, address beneficiary, uint256 amount, uint256 distributedAmount, uint256 netDistributedamount, string memo, address caller)
func (_ETHPaymentTerminal *ETHPaymentTerminalFilterer) FilterUseAllowance(opts *bind.FilterOpts, fundingCycleConfiguration []*big.Int, fundingCycleNumber []*big.Int, projectId []*big.Int) (*ETHPaymentTerminalUseAllowanceIterator, error) {

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

	logs, sub, err := _ETHPaymentTerminal.contract.FilterLogs(opts, "UseAllowance", fundingCycleConfigurationRule, fundingCycleNumberRule, projectIdRule)
	if err != nil {
		return nil, err
	}
	return &ETHPaymentTerminalUseAllowanceIterator{contract: _ETHPaymentTerminal.contract, event: "UseAllowance", logs: logs, sub: sub}, nil
}

// WatchUseAllowance is a free log subscription operation binding the contract event 0x8657a0c05a68a912c23c1bd00124afaa8c669063b046bd9bfd22b21d573c5e6d.
//
// Solidity: event UseAllowance(uint256 indexed fundingCycleConfiguration, uint256 indexed fundingCycleNumber, uint256 indexed projectId, address beneficiary, uint256 amount, uint256 distributedAmount, uint256 netDistributedamount, string memo, address caller)
func (_ETHPaymentTerminal *ETHPaymentTerminalFilterer) WatchUseAllowance(opts *bind.WatchOpts, sink chan<- *ETHPaymentTerminalUseAllowance, fundingCycleConfiguration []*big.Int, fundingCycleNumber []*big.Int, projectId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _ETHPaymentTerminal.contract.WatchLogs(opts, "UseAllowance", fundingCycleConfigurationRule, fundingCycleNumberRule, projectIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ETHPaymentTerminalUseAllowance)
				if err := _ETHPaymentTerminal.contract.UnpackLog(event, "UseAllowance", log); err != nil {
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

// ParseUseAllowance is a log parse operation binding the contract event 0x8657a0c05a68a912c23c1bd00124afaa8c669063b046bd9bfd22b21d573c5e6d.
//
// Solidity: event UseAllowance(uint256 indexed fundingCycleConfiguration, uint256 indexed fundingCycleNumber, uint256 indexed projectId, address beneficiary, uint256 amount, uint256 distributedAmount, uint256 netDistributedamount, string memo, address caller)
func (_ETHPaymentTerminal *ETHPaymentTerminalFilterer) ParseUseAllowance(log types.Log) (*ETHPaymentTerminalUseAllowance, error) {
	event := new(ETHPaymentTerminalUseAllowance)
	if err := _ETHPaymentTerminal.contract.UnpackLog(event, "UseAllowance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
