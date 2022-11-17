// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package treaderjoe

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

// ILBRouterLiquidityParameters is an auto generated low-level Go binding around an user-defined struct.
type ILBRouterLiquidityParameters struct {
	TokenX          common.Address
	TokenY          common.Address
	BinStep         *big.Int
	AmountX         *big.Int
	AmountY         *big.Int
	AmountXMin      *big.Int
	AmountYMin      *big.Int
	ActiveIdDesired *big.Int
	IdSlippage      *big.Int
	DeltaIds        []*big.Int
	DistributionX   []*big.Int
	DistributionY   []*big.Int
	To              common.Address
	Deadline        *big.Int
}

// RouterMetaData contains all meta data concerning the Router contract.
var RouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractILBFactory\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"contractIJoeFactory\",\"name\":\"_oldFactory\",\"type\":\"address\"},{\"internalType\":\"contractIWAVAX\",\"name\":\"_wavax\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bp\",\"type\":\"uint256\"}],\"name\":\"BinHelper__BinStepOverflows\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BinHelper__IdOverflows\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"JoeLibrary__InsufficientAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"JoeLibrary__InsufficientLiquidity\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountXMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountYMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountY\",\"type\":\"uint256\"}],\"name\":\"LBRouter__AmountSlippageCaught\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"LBRouter__BinReserveOverflows\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LBRouter__BrokenSwapSafetyCheck\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentTimestamp\",\"type\":\"uint256\"}],\"name\":\"LBRouter__DeadlineExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LBRouter__FailedToSendAVAX\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"idSlippage\",\"type\":\"uint256\"}],\"name\":\"LBRouter__IdDesiredOverflows\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"id\",\"type\":\"int256\"}],\"name\":\"LBRouter__IdOverflows\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"activeIdDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"idSlippage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"activeId\",\"type\":\"uint256\"}],\"name\":\"LBRouter__IdSlippageCaught\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"LBRouter__InsufficientAmountOut\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wrongToken\",\"type\":\"address\"}],\"name\":\"LBRouter__InvalidTokenPath\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LBRouter__LengthsMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"name\":\"LBRouter__MaxAmountInExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LBRouter__NotFactoryOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenX\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenY\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"binStep\",\"type\":\"uint256\"}],\"name\":\"LBRouter__PairNotCreated\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LBRouter__SenderIsNotWAVAX\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"LBRouter__SwapOverflows\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"excess\",\"type\":\"uint256\"}],\"name\":\"LBRouter__TooMuchTokensIn\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserve\",\"type\":\"uint256\"}],\"name\":\"LBRouter__WrongAmounts\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenX\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenY\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountY\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"msgValue\",\"type\":\"uint256\"}],\"name\":\"LBRouter__WrongAvaxLiquidityParameters\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LBRouter__WrongTokenOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Math128x128__LogUnderflow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"y\",\"type\":\"int256\"}],\"name\":\"Math128x128__PowerUnderflow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"prod1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"denominator\",\"type\":\"uint256\"}],\"name\":\"Math512Bits__MulDivOverflow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"prod1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"}],\"name\":\"Math512Bits__MulShiftOverflow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"}],\"name\":\"Math512Bits__OffsetOverflows\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"}],\"name\":\"SafeCast__Exceeds128Bits\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"}],\"name\":\"SafeCast__Exceeds40Bits\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenHelper__CallFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenHelper__NonContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenHelper__TransferFailed\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"contractIERC20\",\"name\":\"tokenX\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenY\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"binStep\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountY\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountXMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountYMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"activeIdDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"idSlippage\",\"type\":\"uint256\"},{\"internalType\":\"int256[]\",\"name\":\"deltaIds\",\"type\":\"int256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"distributionX\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"distributionY\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structILBRouter.LiquidityParameters\",\"name\":\"_liquidityParameters\",\"type\":\"tuple\"}],\"name\":\"addLiquidity\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"depositIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"liquidityMinted\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"contractIERC20\",\"name\":\"tokenX\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenY\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"binStep\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountY\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountXMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountYMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"activeIdDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"idSlippage\",\"type\":\"uint256\"},{\"internalType\":\"int256[]\",\"name\":\"deltaIds\",\"type\":\"int256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"distributionX\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"distributionY\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structILBRouter.LiquidityParameters\",\"name\":\"_liquidityParameters\",\"type\":\"tuple\"}],\"name\":\"addLiquidityAVAX\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"depositIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"liquidityMinted\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_tokenX\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"_tokenY\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"_activeId\",\"type\":\"uint24\"},{\"internalType\":\"uint16\",\"name\":\"_binStep\",\"type\":\"uint16\"}],\"name\":\"createLBPair\",\"outputs\":[{\"internalType\":\"contractILBPair\",\"name\":\"pair\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"contractILBFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractILBPair\",\"name\":\"_LBPair\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"getIdFromPrice\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractILBPair\",\"name\":\"_LBPair\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"_id\",\"type\":\"uint24\"}],\"name\":\"getPriceFromId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractILBPair\",\"name\":\"_LBPair\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountOut\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_swapForY\",\"type\":\"bool\"}],\"name\":\"getSwapIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feesIn\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractILBPair\",\"name\":\"_LBPair\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountIn\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_swapForY\",\"type\":\"bool\"}],\"name\":\"getSwapOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feesIn\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oldFactory\",\"outputs\":[{\"internalType\":\"contractIJoeFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_tokenX\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"_tokenY\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_binStep\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_amountXMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amountYMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountY\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_binStep\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amountAVAXMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidityAVAX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAVAX\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_pairBinSteps\",\"type\":\"uint256[]\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"_tokenPath\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"}],\"name\":\"swapAVAXForExactTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amountsIn\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_pairBinSteps\",\"type\":\"uint256[]\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"_tokenPath\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactAVAXForTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_pairBinSteps\",\"type\":\"uint256[]\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"_tokenPath\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactAVAXForTokensSupportingFeeOnTransferTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amountOutMinAVAX\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_pairBinSteps\",\"type\":\"uint256[]\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"_tokenPath\",\"type\":\"address[]\"},{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForAVAX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amountOutMinAVAX\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_pairBinSteps\",\"type\":\"uint256[]\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"_tokenPath\",\"type\":\"address[]\"},{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForAVAXSupportingFeeOnTransferTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_pairBinSteps\",\"type\":\"uint256[]\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"_tokenPath\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_pairBinSteps\",\"type\":\"uint256[]\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"_tokenPath\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokensSupportingFeeOnTransferTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountAVAXOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_pairBinSteps\",\"type\":\"uint256[]\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"_tokenPath\",\"type\":\"address[]\"},{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"}],\"name\":\"swapTokensForExactAVAX\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amountsIn\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_pairBinSteps\",\"type\":\"uint256[]\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"_tokenPath\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"}],\"name\":\"swapTokensForExactTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amountsIn\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"sweep\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractILBToken\",\"name\":\"_lbToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"sweepLBToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wavax\",\"outputs\":[{\"internalType\":\"contractIWAVAX\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// RouterABI is the input ABI used to generate the binding from.
// Deprecated: Use RouterMetaData.ABI instead.
var RouterABI = RouterMetaData.ABI

// Router is an auto generated Go binding around an Ethereum contract.
type Router struct {
	RouterCaller     // Read-only binding to the contract
	RouterTransactor // Write-only binding to the contract
	RouterFilterer   // Log filterer for contract events
}

// RouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type RouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RouterSession struct {
	Contract     *Router           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RouterCallerSession struct {
	Contract *RouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RouterTransactorSession struct {
	Contract     *RouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type RouterRaw struct {
	Contract *Router // Generic contract binding to access the raw methods on
}

// RouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RouterCallerRaw struct {
	Contract *RouterCaller // Generic read-only contract binding to access the raw methods on
}

// RouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RouterTransactorRaw struct {
	Contract *RouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRouter creates a new instance of Router, bound to a specific deployed contract.
func NewRouter(address common.Address, backend bind.ContractBackend) (*Router, error) {
	contract, err := bindRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Router{RouterCaller: RouterCaller{contract: contract}, RouterTransactor: RouterTransactor{contract: contract}, RouterFilterer: RouterFilterer{contract: contract}}, nil
}

// NewRouterCaller creates a new read-only instance of Router, bound to a specific deployed contract.
func NewRouterCaller(address common.Address, caller bind.ContractCaller) (*RouterCaller, error) {
	contract, err := bindRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RouterCaller{contract: contract}, nil
}

// NewRouterTransactor creates a new write-only instance of Router, bound to a specific deployed contract.
func NewRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*RouterTransactor, error) {
	contract, err := bindRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RouterTransactor{contract: contract}, nil
}

// NewRouterFilterer creates a new log filterer instance of Router, bound to a specific deployed contract.
func NewRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*RouterFilterer, error) {
	contract, err := bindRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RouterFilterer{contract: contract}, nil
}

// bindRouter binds a generic wrapper to an already deployed contract.
func bindRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RouterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Router *RouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Router.Contract.RouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Router *RouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Router.Contract.RouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Router *RouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Router.Contract.RouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Router *RouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Router.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Router *RouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Router.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Router *RouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Router.Contract.contract.Transact(opts, method, params...)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Router *RouterCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Router *RouterSession) Factory() (common.Address, error) {
	return _Router.Contract.Factory(&_Router.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Router *RouterCallerSession) Factory() (common.Address, error) {
	return _Router.Contract.Factory(&_Router.CallOpts)
}

// GetIdFromPrice is a free data retrieval call binding the contract method 0xf96fe925.
//
// Solidity: function getIdFromPrice(address _LBPair, uint256 _price) view returns(uint24)
func (_Router *RouterCaller) GetIdFromPrice(opts *bind.CallOpts, _LBPair common.Address, _price *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "getIdFromPrice", _LBPair, _price)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetIdFromPrice is a free data retrieval call binding the contract method 0xf96fe925.
//
// Solidity: function getIdFromPrice(address _LBPair, uint256 _price) view returns(uint24)
func (_Router *RouterSession) GetIdFromPrice(_LBPair common.Address, _price *big.Int) (*big.Int, error) {
	return _Router.Contract.GetIdFromPrice(&_Router.CallOpts, _LBPair, _price)
}

// GetIdFromPrice is a free data retrieval call binding the contract method 0xf96fe925.
//
// Solidity: function getIdFromPrice(address _LBPair, uint256 _price) view returns(uint24)
func (_Router *RouterCallerSession) GetIdFromPrice(_LBPair common.Address, _price *big.Int) (*big.Int, error) {
	return _Router.Contract.GetIdFromPrice(&_Router.CallOpts, _LBPair, _price)
}

// GetPriceFromId is a free data retrieval call binding the contract method 0xd0e380f2.
//
// Solidity: function getPriceFromId(address _LBPair, uint24 _id) view returns(uint256)
func (_Router *RouterCaller) GetPriceFromId(opts *bind.CallOpts, _LBPair common.Address, _id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "getPriceFromId", _LBPair, _id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPriceFromId is a free data retrieval call binding the contract method 0xd0e380f2.
//
// Solidity: function getPriceFromId(address _LBPair, uint24 _id) view returns(uint256)
func (_Router *RouterSession) GetPriceFromId(_LBPair common.Address, _id *big.Int) (*big.Int, error) {
	return _Router.Contract.GetPriceFromId(&_Router.CallOpts, _LBPair, _id)
}

// GetPriceFromId is a free data retrieval call binding the contract method 0xd0e380f2.
//
// Solidity: function getPriceFromId(address _LBPair, uint24 _id) view returns(uint256)
func (_Router *RouterCallerSession) GetPriceFromId(_LBPair common.Address, _id *big.Int) (*big.Int, error) {
	return _Router.Contract.GetPriceFromId(&_Router.CallOpts, _LBPair, _id)
}

// GetSwapIn is a free data retrieval call binding the contract method 0x5bdd4b7c.
//
// Solidity: function getSwapIn(address _LBPair, uint256 _amountOut, bool _swapForY) view returns(uint256 amountIn, uint256 feesIn)
func (_Router *RouterCaller) GetSwapIn(opts *bind.CallOpts, _LBPair common.Address, _amountOut *big.Int, _swapForY bool) (struct {
	AmountIn *big.Int
	FeesIn   *big.Int
}, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "getSwapIn", _LBPair, _amountOut, _swapForY)

	outstruct := new(struct {
		AmountIn *big.Int
		FeesIn   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AmountIn = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.FeesIn = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetSwapIn is a free data retrieval call binding the contract method 0x5bdd4b7c.
//
// Solidity: function getSwapIn(address _LBPair, uint256 _amountOut, bool _swapForY) view returns(uint256 amountIn, uint256 feesIn)
func (_Router *RouterSession) GetSwapIn(_LBPair common.Address, _amountOut *big.Int, _swapForY bool) (struct {
	AmountIn *big.Int
	FeesIn   *big.Int
}, error) {
	return _Router.Contract.GetSwapIn(&_Router.CallOpts, _LBPair, _amountOut, _swapForY)
}

// GetSwapIn is a free data retrieval call binding the contract method 0x5bdd4b7c.
//
// Solidity: function getSwapIn(address _LBPair, uint256 _amountOut, bool _swapForY) view returns(uint256 amountIn, uint256 feesIn)
func (_Router *RouterCallerSession) GetSwapIn(_LBPair common.Address, _amountOut *big.Int, _swapForY bool) (struct {
	AmountIn *big.Int
	FeesIn   *big.Int
}, error) {
	return _Router.Contract.GetSwapIn(&_Router.CallOpts, _LBPair, _amountOut, _swapForY)
}

// GetSwapOut is a free data retrieval call binding the contract method 0x2004b724.
//
// Solidity: function getSwapOut(address _LBPair, uint256 _amountIn, bool _swapForY) view returns(uint256 amountOut, uint256 feesIn)
func (_Router *RouterCaller) GetSwapOut(opts *bind.CallOpts, _LBPair common.Address, _amountIn *big.Int, _swapForY bool) (struct {
	AmountOut *big.Int
	FeesIn    *big.Int
}, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "getSwapOut", _LBPair, _amountIn, _swapForY)

	outstruct := new(struct {
		AmountOut *big.Int
		FeesIn    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AmountOut = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.FeesIn = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetSwapOut is a free data retrieval call binding the contract method 0x2004b724.
//
// Solidity: function getSwapOut(address _LBPair, uint256 _amountIn, bool _swapForY) view returns(uint256 amountOut, uint256 feesIn)
func (_Router *RouterSession) GetSwapOut(_LBPair common.Address, _amountIn *big.Int, _swapForY bool) (struct {
	AmountOut *big.Int
	FeesIn    *big.Int
}, error) {
	return _Router.Contract.GetSwapOut(&_Router.CallOpts, _LBPair, _amountIn, _swapForY)
}

// GetSwapOut is a free data retrieval call binding the contract method 0x2004b724.
//
// Solidity: function getSwapOut(address _LBPair, uint256 _amountIn, bool _swapForY) view returns(uint256 amountOut, uint256 feesIn)
func (_Router *RouterCallerSession) GetSwapOut(_LBPair common.Address, _amountIn *big.Int, _swapForY bool) (struct {
	AmountOut *big.Int
	FeesIn    *big.Int
}, error) {
	return _Router.Contract.GetSwapOut(&_Router.CallOpts, _LBPair, _amountIn, _swapForY)
}

// OldFactory is a free data retrieval call binding the contract method 0x1bd6dfe1.
//
// Solidity: function oldFactory() view returns(address)
func (_Router *RouterCaller) OldFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "oldFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OldFactory is a free data retrieval call binding the contract method 0x1bd6dfe1.
//
// Solidity: function oldFactory() view returns(address)
func (_Router *RouterSession) OldFactory() (common.Address, error) {
	return _Router.Contract.OldFactory(&_Router.CallOpts)
}

// OldFactory is a free data retrieval call binding the contract method 0x1bd6dfe1.
//
// Solidity: function oldFactory() view returns(address)
func (_Router *RouterCallerSession) OldFactory() (common.Address, error) {
	return _Router.Contract.OldFactory(&_Router.CallOpts)
}

// Wavax is a free data retrieval call binding the contract method 0x117be4c2.
//
// Solidity: function wavax() view returns(address)
func (_Router *RouterCaller) Wavax(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "wavax")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Wavax is a free data retrieval call binding the contract method 0x117be4c2.
//
// Solidity: function wavax() view returns(address)
func (_Router *RouterSession) Wavax() (common.Address, error) {
	return _Router.Contract.Wavax(&_Router.CallOpts)
}

// Wavax is a free data retrieval call binding the contract method 0x117be4c2.
//
// Solidity: function wavax() view returns(address)
func (_Router *RouterCallerSession) Wavax() (common.Address, error) {
	return _Router.Contract.Wavax(&_Router.CallOpts)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xe324a3e4.
//
// Solidity: function addLiquidity((address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,int256[],uint256[],uint256[],address,uint256) _liquidityParameters) returns(uint256[] depositIds, uint256[] liquidityMinted)
func (_Router *RouterTransactor) AddLiquidity(opts *bind.TransactOpts, _liquidityParameters ILBRouterLiquidityParameters) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "addLiquidity", _liquidityParameters)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xe324a3e4.
//
// Solidity: function addLiquidity((address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,int256[],uint256[],uint256[],address,uint256) _liquidityParameters) returns(uint256[] depositIds, uint256[] liquidityMinted)
func (_Router *RouterSession) AddLiquidity(_liquidityParameters ILBRouterLiquidityParameters) (*types.Transaction, error) {
	return _Router.Contract.AddLiquidity(&_Router.TransactOpts, _liquidityParameters)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xe324a3e4.
//
// Solidity: function addLiquidity((address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,int256[],uint256[],uint256[],address,uint256) _liquidityParameters) returns(uint256[] depositIds, uint256[] liquidityMinted)
func (_Router *RouterTransactorSession) AddLiquidity(_liquidityParameters ILBRouterLiquidityParameters) (*types.Transaction, error) {
	return _Router.Contract.AddLiquidity(&_Router.TransactOpts, _liquidityParameters)
}

// AddLiquidityAVAX is a paid mutator transaction binding the contract method 0xea8f43d8.
//
// Solidity: function addLiquidityAVAX((address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,int256[],uint256[],uint256[],address,uint256) _liquidityParameters) payable returns(uint256[] depositIds, uint256[] liquidityMinted)
func (_Router *RouterTransactor) AddLiquidityAVAX(opts *bind.TransactOpts, _liquidityParameters ILBRouterLiquidityParameters) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "addLiquidityAVAX", _liquidityParameters)
}

// AddLiquidityAVAX is a paid mutator transaction binding the contract method 0xea8f43d8.
//
// Solidity: function addLiquidityAVAX((address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,int256[],uint256[],uint256[],address,uint256) _liquidityParameters) payable returns(uint256[] depositIds, uint256[] liquidityMinted)
func (_Router *RouterSession) AddLiquidityAVAX(_liquidityParameters ILBRouterLiquidityParameters) (*types.Transaction, error) {
	return _Router.Contract.AddLiquidityAVAX(&_Router.TransactOpts, _liquidityParameters)
}

// AddLiquidityAVAX is a paid mutator transaction binding the contract method 0xea8f43d8.
//
// Solidity: function addLiquidityAVAX((address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,int256[],uint256[],uint256[],address,uint256) _liquidityParameters) payable returns(uint256[] depositIds, uint256[] liquidityMinted)
func (_Router *RouterTransactorSession) AddLiquidityAVAX(_liquidityParameters ILBRouterLiquidityParameters) (*types.Transaction, error) {
	return _Router.Contract.AddLiquidityAVAX(&_Router.TransactOpts, _liquidityParameters)
}

// CreateLBPair is a paid mutator transaction binding the contract method 0x659ac74b.
//
// Solidity: function createLBPair(address _tokenX, address _tokenY, uint24 _activeId, uint16 _binStep) returns(address pair)
func (_Router *RouterTransactor) CreateLBPair(opts *bind.TransactOpts, _tokenX common.Address, _tokenY common.Address, _activeId *big.Int, _binStep uint16) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "createLBPair", _tokenX, _tokenY, _activeId, _binStep)
}

// CreateLBPair is a paid mutator transaction binding the contract method 0x659ac74b.
//
// Solidity: function createLBPair(address _tokenX, address _tokenY, uint24 _activeId, uint16 _binStep) returns(address pair)
func (_Router *RouterSession) CreateLBPair(_tokenX common.Address, _tokenY common.Address, _activeId *big.Int, _binStep uint16) (*types.Transaction, error) {
	return _Router.Contract.CreateLBPair(&_Router.TransactOpts, _tokenX, _tokenY, _activeId, _binStep)
}

// CreateLBPair is a paid mutator transaction binding the contract method 0x659ac74b.
//
// Solidity: function createLBPair(address _tokenX, address _tokenY, uint24 _activeId, uint16 _binStep) returns(address pair)
func (_Router *RouterTransactorSession) CreateLBPair(_tokenX common.Address, _tokenY common.Address, _activeId *big.Int, _binStep uint16) (*types.Transaction, error) {
	return _Router.Contract.CreateLBPair(&_Router.TransactOpts, _tokenX, _tokenY, _activeId, _binStep)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xc22159b6.
//
// Solidity: function removeLiquidity(address _tokenX, address _tokenY, uint16 _binStep, uint256 _amountXMin, uint256 _amountYMin, uint256[] _ids, uint256[] _amounts, address _to, uint256 _deadline) returns(uint256 amountX, uint256 amountY)
func (_Router *RouterTransactor) RemoveLiquidity(opts *bind.TransactOpts, _tokenX common.Address, _tokenY common.Address, _binStep uint16, _amountXMin *big.Int, _amountYMin *big.Int, _ids []*big.Int, _amounts []*big.Int, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "removeLiquidity", _tokenX, _tokenY, _binStep, _amountXMin, _amountYMin, _ids, _amounts, _to, _deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xc22159b6.
//
// Solidity: function removeLiquidity(address _tokenX, address _tokenY, uint16 _binStep, uint256 _amountXMin, uint256 _amountYMin, uint256[] _ids, uint256[] _amounts, address _to, uint256 _deadline) returns(uint256 amountX, uint256 amountY)
func (_Router *RouterSession) RemoveLiquidity(_tokenX common.Address, _tokenY common.Address, _binStep uint16, _amountXMin *big.Int, _amountYMin *big.Int, _ids []*big.Int, _amounts []*big.Int, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _Router.Contract.RemoveLiquidity(&_Router.TransactOpts, _tokenX, _tokenY, _binStep, _amountXMin, _amountYMin, _ids, _amounts, _to, _deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xc22159b6.
//
// Solidity: function removeLiquidity(address _tokenX, address _tokenY, uint16 _binStep, uint256 _amountXMin, uint256 _amountYMin, uint256[] _ids, uint256[] _amounts, address _to, uint256 _deadline) returns(uint256 amountX, uint256 amountY)
func (_Router *RouterTransactorSession) RemoveLiquidity(_tokenX common.Address, _tokenY common.Address, _binStep uint16, _amountXMin *big.Int, _amountYMin *big.Int, _ids []*big.Int, _amounts []*big.Int, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _Router.Contract.RemoveLiquidity(&_Router.TransactOpts, _tokenX, _tokenY, _binStep, _amountXMin, _amountYMin, _ids, _amounts, _to, _deadline)
}

// RemoveLiquidityAVAX is a paid mutator transaction binding the contract method 0xbcb1c957.
//
// Solidity: function removeLiquidityAVAX(address _token, uint16 _binStep, uint256 _amountTokenMin, uint256 _amountAVAXMin, uint256[] _ids, uint256[] _amounts, address _to, uint256 _deadline) returns(uint256 amountToken, uint256 amountAVAX)
func (_Router *RouterTransactor) RemoveLiquidityAVAX(opts *bind.TransactOpts, _token common.Address, _binStep uint16, _amountTokenMin *big.Int, _amountAVAXMin *big.Int, _ids []*big.Int, _amounts []*big.Int, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "removeLiquidityAVAX", _token, _binStep, _amountTokenMin, _amountAVAXMin, _ids, _amounts, _to, _deadline)
}

// RemoveLiquidityAVAX is a paid mutator transaction binding the contract method 0xbcb1c957.
//
// Solidity: function removeLiquidityAVAX(address _token, uint16 _binStep, uint256 _amountTokenMin, uint256 _amountAVAXMin, uint256[] _ids, uint256[] _amounts, address _to, uint256 _deadline) returns(uint256 amountToken, uint256 amountAVAX)
func (_Router *RouterSession) RemoveLiquidityAVAX(_token common.Address, _binStep uint16, _amountTokenMin *big.Int, _amountAVAXMin *big.Int, _ids []*big.Int, _amounts []*big.Int, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _Router.Contract.RemoveLiquidityAVAX(&_Router.TransactOpts, _token, _binStep, _amountTokenMin, _amountAVAXMin, _ids, _amounts, _to, _deadline)
}

// RemoveLiquidityAVAX is a paid mutator transaction binding the contract method 0xbcb1c957.
//
// Solidity: function removeLiquidityAVAX(address _token, uint16 _binStep, uint256 _amountTokenMin, uint256 _amountAVAXMin, uint256[] _ids, uint256[] _amounts, address _to, uint256 _deadline) returns(uint256 amountToken, uint256 amountAVAX)
func (_Router *RouterTransactorSession) RemoveLiquidityAVAX(_token common.Address, _binStep uint16, _amountTokenMin *big.Int, _amountAVAXMin *big.Int, _ids []*big.Int, _amounts []*big.Int, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _Router.Contract.RemoveLiquidityAVAX(&_Router.TransactOpts, _token, _binStep, _amountTokenMin, _amountAVAXMin, _ids, _amounts, _to, _deadline)
}

// SwapAVAXForExactTokens is a paid mutator transaction binding the contract method 0x42f564a0.
//
// Solidity: function swapAVAXForExactTokens(uint256 _amountOut, uint256[] _pairBinSteps, address[] _tokenPath, address _to, uint256 _deadline) payable returns(uint256[] amountsIn)
func (_Router *RouterTransactor) SwapAVAXForExactTokens(opts *bind.TransactOpts, _amountOut *big.Int, _pairBinSteps []*big.Int, _tokenPath []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "swapAVAXForExactTokens", _amountOut, _pairBinSteps, _tokenPath, _to, _deadline)
}

// SwapAVAXForExactTokens is a paid mutator transaction binding the contract method 0x42f564a0.
//
// Solidity: function swapAVAXForExactTokens(uint256 _amountOut, uint256[] _pairBinSteps, address[] _tokenPath, address _to, uint256 _deadline) payable returns(uint256[] amountsIn)
func (_Router *RouterSession) SwapAVAXForExactTokens(_amountOut *big.Int, _pairBinSteps []*big.Int, _tokenPath []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _Router.Contract.SwapAVAXForExactTokens(&_Router.TransactOpts, _amountOut, _pairBinSteps, _tokenPath, _to, _deadline)
}

// SwapAVAXForExactTokens is a paid mutator transaction binding the contract method 0x42f564a0.
//
// Solidity: function swapAVAXForExactTokens(uint256 _amountOut, uint256[] _pairBinSteps, address[] _tokenPath, address _to, uint256 _deadline) payable returns(uint256[] amountsIn)
func (_Router *RouterTransactorSession) SwapAVAXForExactTokens(_amountOut *big.Int, _pairBinSteps []*big.Int, _tokenPath []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _Router.Contract.SwapAVAXForExactTokens(&_Router.TransactOpts, _amountOut, _pairBinSteps, _tokenPath, _to, _deadline)
}

// SwapExactAVAXForTokens is a paid mutator transaction binding the contract method 0x264bb94e.
//
// Solidity: function swapExactAVAXForTokens(uint256 _amountOutMin, uint256[] _pairBinSteps, address[] _tokenPath, address _to, uint256 _deadline) payable returns(uint256 amountOut)
func (_Router *RouterTransactor) SwapExactAVAXForTokens(opts *bind.TransactOpts, _amountOutMin *big.Int, _pairBinSteps []*big.Int, _tokenPath []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "swapExactAVAXForTokens", _amountOutMin, _pairBinSteps, _tokenPath, _to, _deadline)
}

// SwapExactAVAXForTokens is a paid mutator transaction binding the contract method 0x264bb94e.
//
// Solidity: function swapExactAVAXForTokens(uint256 _amountOutMin, uint256[] _pairBinSteps, address[] _tokenPath, address _to, uint256 _deadline) payable returns(uint256 amountOut)
func (_Router *RouterSession) SwapExactAVAXForTokens(_amountOutMin *big.Int, _pairBinSteps []*big.Int, _tokenPath []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _Router.Contract.SwapExactAVAXForTokens(&_Router.TransactOpts, _amountOutMin, _pairBinSteps, _tokenPath, _to, _deadline)
}

// SwapExactAVAXForTokens is a paid mutator transaction binding the contract method 0x264bb94e.
//
// Solidity: function swapExactAVAXForTokens(uint256 _amountOutMin, uint256[] _pairBinSteps, address[] _tokenPath, address _to, uint256 _deadline) payable returns(uint256 amountOut)
func (_Router *RouterTransactorSession) SwapExactAVAXForTokens(_amountOutMin *big.Int, _pairBinSteps []*big.Int, _tokenPath []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _Router.Contract.SwapExactAVAXForTokens(&_Router.TransactOpts, _amountOutMin, _pairBinSteps, _tokenPath, _to, _deadline)
}

// SwapExactAVAXForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x440830bd.
//
// Solidity: function swapExactAVAXForTokensSupportingFeeOnTransferTokens(uint256 _amountOutMin, uint256[] _pairBinSteps, address[] _tokenPath, address _to, uint256 _deadline) payable returns(uint256 amountOut)
func (_Router *RouterTransactor) SwapExactAVAXForTokensSupportingFeeOnTransferTokens(opts *bind.TransactOpts, _amountOutMin *big.Int, _pairBinSteps []*big.Int, _tokenPath []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "swapExactAVAXForTokensSupportingFeeOnTransferTokens", _amountOutMin, _pairBinSteps, _tokenPath, _to, _deadline)
}

// SwapExactAVAXForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x440830bd.
//
// Solidity: function swapExactAVAXForTokensSupportingFeeOnTransferTokens(uint256 _amountOutMin, uint256[] _pairBinSteps, address[] _tokenPath, address _to, uint256 _deadline) payable returns(uint256 amountOut)
func (_Router *RouterSession) SwapExactAVAXForTokensSupportingFeeOnTransferTokens(_amountOutMin *big.Int, _pairBinSteps []*big.Int, _tokenPath []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _Router.Contract.SwapExactAVAXForTokensSupportingFeeOnTransferTokens(&_Router.TransactOpts, _amountOutMin, _pairBinSteps, _tokenPath, _to, _deadline)
}

// SwapExactAVAXForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x440830bd.
//
// Solidity: function swapExactAVAXForTokensSupportingFeeOnTransferTokens(uint256 _amountOutMin, uint256[] _pairBinSteps, address[] _tokenPath, address _to, uint256 _deadline) payable returns(uint256 amountOut)
func (_Router *RouterTransactorSession) SwapExactAVAXForTokensSupportingFeeOnTransferTokens(_amountOutMin *big.Int, _pairBinSteps []*big.Int, _tokenPath []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _Router.Contract.SwapExactAVAXForTokensSupportingFeeOnTransferTokens(&_Router.TransactOpts, _amountOutMin, _pairBinSteps, _tokenPath, _to, _deadline)
}

// SwapExactTokensForAVAX is a paid mutator transaction binding the contract method 0xfb321c70.
//
// Solidity: function swapExactTokensForAVAX(uint256 _amountIn, uint256 _amountOutMinAVAX, uint256[] _pairBinSteps, address[] _tokenPath, address _to, uint256 _deadline) returns(uint256 amountOut)
func (_Router *RouterTransactor) SwapExactTokensForAVAX(opts *bind.TransactOpts, _amountIn *big.Int, _amountOutMinAVAX *big.Int, _pairBinSteps []*big.Int, _tokenPath []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "swapExactTokensForAVAX", _amountIn, _amountOutMinAVAX, _pairBinSteps, _tokenPath, _to, _deadline)
}

// SwapExactTokensForAVAX is a paid mutator transaction binding the contract method 0xfb321c70.
//
// Solidity: function swapExactTokensForAVAX(uint256 _amountIn, uint256 _amountOutMinAVAX, uint256[] _pairBinSteps, address[] _tokenPath, address _to, uint256 _deadline) returns(uint256 amountOut)
func (_Router *RouterSession) SwapExactTokensForAVAX(_amountIn *big.Int, _amountOutMinAVAX *big.Int, _pairBinSteps []*big.Int, _tokenPath []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _Router.Contract.SwapExactTokensForAVAX(&_Router.TransactOpts, _amountIn, _amountOutMinAVAX, _pairBinSteps, _tokenPath, _to, _deadline)
}

// SwapExactTokensForAVAX is a paid mutator transaction binding the contract method 0xfb321c70.
//
// Solidity: function swapExactTokensForAVAX(uint256 _amountIn, uint256 _amountOutMinAVAX, uint256[] _pairBinSteps, address[] _tokenPath, address _to, uint256 _deadline) returns(uint256 amountOut)
func (_Router *RouterTransactorSession) SwapExactTokensForAVAX(_amountIn *big.Int, _amountOutMinAVAX *big.Int, _pairBinSteps []*big.Int, _tokenPath []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _Router.Contract.SwapExactTokensForAVAX(&_Router.TransactOpts, _amountIn, _amountOutMinAVAX, _pairBinSteps, _tokenPath, _to, _deadline)
}

// SwapExactTokensForAVAXSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x9a17e820.
//
// Solidity: function swapExactTokensForAVAXSupportingFeeOnTransferTokens(uint256 _amountIn, uint256 _amountOutMinAVAX, uint256[] _pairBinSteps, address[] _tokenPath, address _to, uint256 _deadline) returns(uint256 amountOut)
func (_Router *RouterTransactor) SwapExactTokensForAVAXSupportingFeeOnTransferTokens(opts *bind.TransactOpts, _amountIn *big.Int, _amountOutMinAVAX *big.Int, _pairBinSteps []*big.Int, _tokenPath []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "swapExactTokensForAVAXSupportingFeeOnTransferTokens", _amountIn, _amountOutMinAVAX, _pairBinSteps, _tokenPath, _to, _deadline)
}

// SwapExactTokensForAVAXSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x9a17e820.
//
// Solidity: function swapExactTokensForAVAXSupportingFeeOnTransferTokens(uint256 _amountIn, uint256 _amountOutMinAVAX, uint256[] _pairBinSteps, address[] _tokenPath, address _to, uint256 _deadline) returns(uint256 amountOut)
func (_Router *RouterSession) SwapExactTokensForAVAXSupportingFeeOnTransferTokens(_amountIn *big.Int, _amountOutMinAVAX *big.Int, _pairBinSteps []*big.Int, _tokenPath []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _Router.Contract.SwapExactTokensForAVAXSupportingFeeOnTransferTokens(&_Router.TransactOpts, _amountIn, _amountOutMinAVAX, _pairBinSteps, _tokenPath, _to, _deadline)
}

// SwapExactTokensForAVAXSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x9a17e820.
//
// Solidity: function swapExactTokensForAVAXSupportingFeeOnTransferTokens(uint256 _amountIn, uint256 _amountOutMinAVAX, uint256[] _pairBinSteps, address[] _tokenPath, address _to, uint256 _deadline) returns(uint256 amountOut)
func (_Router *RouterTransactorSession) SwapExactTokensForAVAXSupportingFeeOnTransferTokens(_amountIn *big.Int, _amountOutMinAVAX *big.Int, _pairBinSteps []*big.Int, _tokenPath []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _Router.Contract.SwapExactTokensForAVAXSupportingFeeOnTransferTokens(&_Router.TransactOpts, _amountIn, _amountOutMinAVAX, _pairBinSteps, _tokenPath, _to, _deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x6d0ff495.
//
// Solidity: function swapExactTokensForTokens(uint256 _amountIn, uint256 _amountOutMin, uint256[] _pairBinSteps, address[] _tokenPath, address _to, uint256 _deadline) returns(uint256 amountOut)
func (_Router *RouterTransactor) SwapExactTokensForTokens(opts *bind.TransactOpts, _amountIn *big.Int, _amountOutMin *big.Int, _pairBinSteps []*big.Int, _tokenPath []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "swapExactTokensForTokens", _amountIn, _amountOutMin, _pairBinSteps, _tokenPath, _to, _deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x6d0ff495.
//
// Solidity: function swapExactTokensForTokens(uint256 _amountIn, uint256 _amountOutMin, uint256[] _pairBinSteps, address[] _tokenPath, address _to, uint256 _deadline) returns(uint256 amountOut)
func (_Router *RouterSession) SwapExactTokensForTokens(_amountIn *big.Int, _amountOutMin *big.Int, _pairBinSteps []*big.Int, _tokenPath []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _Router.Contract.SwapExactTokensForTokens(&_Router.TransactOpts, _amountIn, _amountOutMin, _pairBinSteps, _tokenPath, _to, _deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x6d0ff495.
//
// Solidity: function swapExactTokensForTokens(uint256 _amountIn, uint256 _amountOutMin, uint256[] _pairBinSteps, address[] _tokenPath, address _to, uint256 _deadline) returns(uint256 amountOut)
func (_Router *RouterTransactorSession) SwapExactTokensForTokens(_amountIn *big.Int, _amountOutMin *big.Int, _pairBinSteps []*big.Int, _tokenPath []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _Router.Contract.SwapExactTokensForTokens(&_Router.TransactOpts, _amountIn, _amountOutMin, _pairBinSteps, _tokenPath, _to, _deadline)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x212a1d94.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 _amountIn, uint256 _amountOutMin, uint256[] _pairBinSteps, address[] _tokenPath, address _to, uint256 _deadline) returns(uint256 amountOut)
func (_Router *RouterTransactor) SwapExactTokensForTokensSupportingFeeOnTransferTokens(opts *bind.TransactOpts, _amountIn *big.Int, _amountOutMin *big.Int, _pairBinSteps []*big.Int, _tokenPath []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "swapExactTokensForTokensSupportingFeeOnTransferTokens", _amountIn, _amountOutMin, _pairBinSteps, _tokenPath, _to, _deadline)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x212a1d94.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 _amountIn, uint256 _amountOutMin, uint256[] _pairBinSteps, address[] _tokenPath, address _to, uint256 _deadline) returns(uint256 amountOut)
func (_Router *RouterSession) SwapExactTokensForTokensSupportingFeeOnTransferTokens(_amountIn *big.Int, _amountOutMin *big.Int, _pairBinSteps []*big.Int, _tokenPath []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _Router.Contract.SwapExactTokensForTokensSupportingFeeOnTransferTokens(&_Router.TransactOpts, _amountIn, _amountOutMin, _pairBinSteps, _tokenPath, _to, _deadline)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x212a1d94.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 _amountIn, uint256 _amountOutMin, uint256[] _pairBinSteps, address[] _tokenPath, address _to, uint256 _deadline) returns(uint256 amountOut)
func (_Router *RouterTransactorSession) SwapExactTokensForTokensSupportingFeeOnTransferTokens(_amountIn *big.Int, _amountOutMin *big.Int, _pairBinSteps []*big.Int, _tokenPath []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _Router.Contract.SwapExactTokensForTokensSupportingFeeOnTransferTokens(&_Router.TransactOpts, _amountIn, _amountOutMin, _pairBinSteps, _tokenPath, _to, _deadline)
}

// SwapTokensForExactAVAX is a paid mutator transaction binding the contract method 0x6d3420ed.
//
// Solidity: function swapTokensForExactAVAX(uint256 _amountAVAXOut, uint256 _amountInMax, uint256[] _pairBinSteps, address[] _tokenPath, address _to, uint256 _deadline) returns(uint256[] amountsIn)
func (_Router *RouterTransactor) SwapTokensForExactAVAX(opts *bind.TransactOpts, _amountAVAXOut *big.Int, _amountInMax *big.Int, _pairBinSteps []*big.Int, _tokenPath []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "swapTokensForExactAVAX", _amountAVAXOut, _amountInMax, _pairBinSteps, _tokenPath, _to, _deadline)
}

// SwapTokensForExactAVAX is a paid mutator transaction binding the contract method 0x6d3420ed.
//
// Solidity: function swapTokensForExactAVAX(uint256 _amountAVAXOut, uint256 _amountInMax, uint256[] _pairBinSteps, address[] _tokenPath, address _to, uint256 _deadline) returns(uint256[] amountsIn)
func (_Router *RouterSession) SwapTokensForExactAVAX(_amountAVAXOut *big.Int, _amountInMax *big.Int, _pairBinSteps []*big.Int, _tokenPath []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _Router.Contract.SwapTokensForExactAVAX(&_Router.TransactOpts, _amountAVAXOut, _amountInMax, _pairBinSteps, _tokenPath, _to, _deadline)
}

// SwapTokensForExactAVAX is a paid mutator transaction binding the contract method 0x6d3420ed.
//
// Solidity: function swapTokensForExactAVAX(uint256 _amountAVAXOut, uint256 _amountInMax, uint256[] _pairBinSteps, address[] _tokenPath, address _to, uint256 _deadline) returns(uint256[] amountsIn)
func (_Router *RouterTransactorSession) SwapTokensForExactAVAX(_amountAVAXOut *big.Int, _amountInMax *big.Int, _pairBinSteps []*big.Int, _tokenPath []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _Router.Contract.SwapTokensForExactAVAX(&_Router.TransactOpts, _amountAVAXOut, _amountInMax, _pairBinSteps, _tokenPath, _to, _deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0xa7b856d3.
//
// Solidity: function swapTokensForExactTokens(uint256 _amountOut, uint256 _amountInMax, uint256[] _pairBinSteps, address[] _tokenPath, address _to, uint256 _deadline) returns(uint256[] amountsIn)
func (_Router *RouterTransactor) SwapTokensForExactTokens(opts *bind.TransactOpts, _amountOut *big.Int, _amountInMax *big.Int, _pairBinSteps []*big.Int, _tokenPath []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "swapTokensForExactTokens", _amountOut, _amountInMax, _pairBinSteps, _tokenPath, _to, _deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0xa7b856d3.
//
// Solidity: function swapTokensForExactTokens(uint256 _amountOut, uint256 _amountInMax, uint256[] _pairBinSteps, address[] _tokenPath, address _to, uint256 _deadline) returns(uint256[] amountsIn)
func (_Router *RouterSession) SwapTokensForExactTokens(_amountOut *big.Int, _amountInMax *big.Int, _pairBinSteps []*big.Int, _tokenPath []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _Router.Contract.SwapTokensForExactTokens(&_Router.TransactOpts, _amountOut, _amountInMax, _pairBinSteps, _tokenPath, _to, _deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0xa7b856d3.
//
// Solidity: function swapTokensForExactTokens(uint256 _amountOut, uint256 _amountInMax, uint256[] _pairBinSteps, address[] _tokenPath, address _to, uint256 _deadline) returns(uint256[] amountsIn)
func (_Router *RouterTransactorSession) SwapTokensForExactTokens(_amountOut *big.Int, _amountInMax *big.Int, _pairBinSteps []*big.Int, _tokenPath []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _Router.Contract.SwapTokensForExactTokens(&_Router.TransactOpts, _amountOut, _amountInMax, _pairBinSteps, _tokenPath, _to, _deadline)
}

// Sweep is a paid mutator transaction binding the contract method 0x62c06767.
//
// Solidity: function sweep(address _token, address _to, uint256 _amount) returns()
func (_Router *RouterTransactor) Sweep(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "sweep", _token, _to, _amount)
}

// Sweep is a paid mutator transaction binding the contract method 0x62c06767.
//
// Solidity: function sweep(address _token, address _to, uint256 _amount) returns()
func (_Router *RouterSession) Sweep(_token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Router.Contract.Sweep(&_Router.TransactOpts, _token, _to, _amount)
}

// Sweep is a paid mutator transaction binding the contract method 0x62c06767.
//
// Solidity: function sweep(address _token, address _to, uint256 _amount) returns()
func (_Router *RouterTransactorSession) Sweep(_token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Router.Contract.Sweep(&_Router.TransactOpts, _token, _to, _amount)
}

// SweepLBToken is a paid mutator transaction binding the contract method 0xe9361c08.
//
// Solidity: function sweepLBToken(address _lbToken, address _to, uint256[] _ids, uint256[] _amounts) returns()
func (_Router *RouterTransactor) SweepLBToken(opts *bind.TransactOpts, _lbToken common.Address, _to common.Address, _ids []*big.Int, _amounts []*big.Int) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "sweepLBToken", _lbToken, _to, _ids, _amounts)
}

// SweepLBToken is a paid mutator transaction binding the contract method 0xe9361c08.
//
// Solidity: function sweepLBToken(address _lbToken, address _to, uint256[] _ids, uint256[] _amounts) returns()
func (_Router *RouterSession) SweepLBToken(_lbToken common.Address, _to common.Address, _ids []*big.Int, _amounts []*big.Int) (*types.Transaction, error) {
	return _Router.Contract.SweepLBToken(&_Router.TransactOpts, _lbToken, _to, _ids, _amounts)
}

// SweepLBToken is a paid mutator transaction binding the contract method 0xe9361c08.
//
// Solidity: function sweepLBToken(address _lbToken, address _to, uint256[] _ids, uint256[] _amounts) returns()
func (_Router *RouterTransactorSession) SweepLBToken(_lbToken common.Address, _to common.Address, _ids []*big.Int, _amounts []*big.Int) (*types.Transaction, error) {
	return _Router.Contract.SweepLBToken(&_Router.TransactOpts, _lbToken, _to, _ids, _amounts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Router *RouterTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Router.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Router *RouterSession) Receive() (*types.Transaction, error) {
	return _Router.Contract.Receive(&_Router.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Router *RouterTransactorSession) Receive() (*types.Transaction, error) {
	return _Router.Contract.Receive(&_Router.TransactOpts)
}
