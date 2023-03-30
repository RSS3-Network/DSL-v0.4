// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package benddao

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

// DataTypesNftConfigurationMap is an auto generated low-level Go binding around an user-defined struct.
type DataTypesNftConfigurationMap struct {
	Data *big.Int
}

// DataTypesNftData is an auto generated low-level Go binding around an user-defined struct.
type DataTypesNftData struct {
	Configuration DataTypesNftConfigurationMap
	BNftAddress   common.Address
	Id            uint8
	MaxSupply     *big.Int
	MaxTokenId    *big.Int
}

// DataTypesReserveConfigurationMap is an auto generated low-level Go binding around an user-defined struct.
type DataTypesReserveConfigurationMap struct {
	Data *big.Int
}

// DataTypesReserveData is an auto generated low-level Go binding around an user-defined struct.
type DataTypesReserveData struct {
	Configuration             DataTypesReserveConfigurationMap
	LiquidityIndex            *big.Int
	VariableBorrowIndex       *big.Int
	CurrentLiquidityRate      *big.Int
	CurrentVariableBorrowRate *big.Int
	LastUpdateTimestamp       *big.Int
	BTokenAddress             common.Address
	DebtTokenAddress          common.Address
	InterestRateAddress       common.Address
	Id                        uint8
}

// BendDAOMetaData contains all meta data concerning the BendDAO contract.
var BendDAOMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bidPrice\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftAsset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"loanId\",\"type\":\"uint256\"}],\"name\":\"Auction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftAsset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftTokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"borrowRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"loanId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"referral\",\"type\":\"uint16\"}],\"name\":\"Borrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"referral\",\"type\":\"uint16\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"remainAmount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftAsset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftTokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"loanId\",\"type\":\"uint256\"}],\"name\":\"Liquidate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"durationTime\",\"type\":\"uint256\"}],\"name\":\"PausedTimeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fineAmount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftAsset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftTokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"loanId\",\"type\":\"uint256\"}],\"name\":\"Redeem\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftAsset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftTokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"loanId\",\"type\":\"uint256\"}],\"name\":\"Repay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidityRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"variableBorrowRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidityIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"variableBorrowIndex\",\"type\":\"uint256\"}],\"name\":\"ReserveDataUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftAsset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bidPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"}],\"name\":\"auction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"nftAssets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"nftTokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"referralCode\",\"type\":\"uint16\"}],\"name\":\"batchBorrow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"nftAssets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"nftTokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"batchRepay\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bool[]\",\"name\":\"\",\"type\":\"bool[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"nftAsset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftTokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"referralCode\",\"type\":\"uint16\"}],\"name\":\"borrow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"referralCode\",\"type\":\"uint16\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balanceFromBefore\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balanceToBefore\",\"type\":\"uint256\"}],\"name\":\"finalizeTransfer\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAddressesProvider\",\"outputs\":[{\"internalType\":\"contractILendPoolAddressesProvider\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxNumberOfNfts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxNumberOfReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftAsset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftTokenId\",\"type\":\"uint256\"}],\"name\":\"getNftAuctionData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"loanId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"bidderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bidPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bidBorrowAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bidFine\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftAsset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftTokenId\",\"type\":\"uint256\"}],\"name\":\"getNftAuctionEndTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"loanId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bidStartTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bidEndTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"redeemEndTimestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"reserveAsset\",\"type\":\"address\"}],\"name\":\"getNftCollateralData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalCollateralInETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalCollateralInReserve\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableBorrowsInETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableBorrowsInReserve\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ltv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationBonus\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getNftConfiguration\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.NftConfigurationMap\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getNftData\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.NftConfigurationMap\",\"name\":\"configuration\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"bNftAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"maxSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTokenId\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.NftData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftAsset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftTokenId\",\"type\":\"uint256\"}],\"name\":\"getNftDebtData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"loanId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"reserveAsset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableBorrows\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"healthFactor\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftAsset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftTokenId\",\"type\":\"uint256\"}],\"name\":\"getNftLiquidatePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"liquidatePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paybackAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNftsList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPausedTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getReserveConfiguration\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.ReserveConfigurationMap\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getReserveData\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.ReserveConfigurationMap\",\"name\":\"configuration\",\"type\":\"tuple\"},{\"internalType\":\"uint128\",\"name\":\"liquidityIndex\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"variableBorrowIndex\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"currentLiquidityRate\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"currentVariableBorrowRate\",\"type\":\"uint128\"},{\"internalType\":\"uint40\",\"name\":\"lastUpdateTimestamp\",\"type\":\"uint40\"},{\"internalType\":\"address\",\"name\":\"bTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"debtTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"interestRateAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"}],\"internalType\":\"structDataTypes.ReserveData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getReserveNormalizedIncome\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getReserveNormalizedVariableDebt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReservesList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bNftAddress\",\"type\":\"address\"}],\"name\":\"initNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"debtTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"interestRateAddress\",\"type\":\"address\"}],\"name\":\"initReserve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractILendPoolAddressesProvider\",\"name\":\"provider\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftAsset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"liquidate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftAsset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bidFine\",\"type\":\"uint256\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftAsset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"repay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"setMaxNumberOfNfts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"setMaxNumberOfReserves\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"configuration\",\"type\":\"uint256\"}],\"name\":\"setNftConfiguration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTokenId\",\"type\":\"uint256\"}],\"name\":\"setNftMaxSupplyAndTokenId\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"val\",\"type\":\"bool\"}],\"name\":\"setPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"durationTime\",\"type\":\"uint256\"}],\"name\":\"setPausedTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"configuration\",\"type\":\"uint256\"}],\"name\":\"setReserveConfiguration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rateAddress\",\"type\":\"address\"}],\"name\":\"setReserveInterestRateAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BendDAOABI is the input ABI used to generate the binding from.
// Deprecated: Use BendDAOMetaData.ABI instead.
var BendDAOABI = BendDAOMetaData.ABI

// BendDAO is an auto generated Go binding around an Ethereum contract.
type BendDAO struct {
	BendDAOCaller     // Read-only binding to the contract
	BendDAOTransactor // Write-only binding to the contract
	BendDAOFilterer   // Log filterer for contract events
}

// BendDAOCaller is an auto generated read-only Go binding around an Ethereum contract.
type BendDAOCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BendDAOTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BendDAOTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BendDAOFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BendDAOFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BendDAOSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BendDAOSession struct {
	Contract     *BendDAO          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BendDAOCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BendDAOCallerSession struct {
	Contract *BendDAOCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// BendDAOTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BendDAOTransactorSession struct {
	Contract     *BendDAOTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// BendDAORaw is an auto generated low-level Go binding around an Ethereum contract.
type BendDAORaw struct {
	Contract *BendDAO // Generic contract binding to access the raw methods on
}

// BendDAOCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BendDAOCallerRaw struct {
	Contract *BendDAOCaller // Generic read-only contract binding to access the raw methods on
}

// BendDAOTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BendDAOTransactorRaw struct {
	Contract *BendDAOTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBendDAO creates a new instance of BendDAO, bound to a specific deployed contract.
func NewBendDAO(address common.Address, backend bind.ContractBackend) (*BendDAO, error) {
	contract, err := bindBendDAO(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BendDAO{BendDAOCaller: BendDAOCaller{contract: contract}, BendDAOTransactor: BendDAOTransactor{contract: contract}, BendDAOFilterer: BendDAOFilterer{contract: contract}}, nil
}

// NewBendDAOCaller creates a new read-only instance of BendDAO, bound to a specific deployed contract.
func NewBendDAOCaller(address common.Address, caller bind.ContractCaller) (*BendDAOCaller, error) {
	contract, err := bindBendDAO(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BendDAOCaller{contract: contract}, nil
}

// NewBendDAOTransactor creates a new write-only instance of BendDAO, bound to a specific deployed contract.
func NewBendDAOTransactor(address common.Address, transactor bind.ContractTransactor) (*BendDAOTransactor, error) {
	contract, err := bindBendDAO(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BendDAOTransactor{contract: contract}, nil
}

// NewBendDAOFilterer creates a new log filterer instance of BendDAO, bound to a specific deployed contract.
func NewBendDAOFilterer(address common.Address, filterer bind.ContractFilterer) (*BendDAOFilterer, error) {
	contract, err := bindBendDAO(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BendDAOFilterer{contract: contract}, nil
}

// bindBendDAO binds a generic wrapper to an already deployed contract.
func bindBendDAO(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BendDAOMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BendDAO *BendDAORaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BendDAO.Contract.BendDAOCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BendDAO *BendDAORaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BendDAO.Contract.BendDAOTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BendDAO *BendDAORaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BendDAO.Contract.BendDAOTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BendDAO *BendDAOCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BendDAO.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BendDAO *BendDAOTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BendDAO.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BendDAO *BendDAOTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BendDAO.Contract.contract.Transact(opts, method, params...)
}

// FinalizeTransfer is a free data retrieval call binding the contract method 0xd5ed3933.
//
// Solidity: function finalizeTransfer(address asset, address from, address to, uint256 amount, uint256 balanceFromBefore, uint256 balanceToBefore) view returns()
func (_BendDAO *BendDAOCaller) FinalizeTransfer(opts *bind.CallOpts, asset common.Address, from common.Address, to common.Address, amount *big.Int, balanceFromBefore *big.Int, balanceToBefore *big.Int) error {
	var out []interface{}
	err := _BendDAO.contract.Call(opts, &out, "finalizeTransfer", asset, from, to, amount, balanceFromBefore, balanceToBefore)

	if err != nil {
		return err
	}

	return err

}

// FinalizeTransfer is a free data retrieval call binding the contract method 0xd5ed3933.
//
// Solidity: function finalizeTransfer(address asset, address from, address to, uint256 amount, uint256 balanceFromBefore, uint256 balanceToBefore) view returns()
func (_BendDAO *BendDAOSession) FinalizeTransfer(asset common.Address, from common.Address, to common.Address, amount *big.Int, balanceFromBefore *big.Int, balanceToBefore *big.Int) error {
	return _BendDAO.Contract.FinalizeTransfer(&_BendDAO.CallOpts, asset, from, to, amount, balanceFromBefore, balanceToBefore)
}

// FinalizeTransfer is a free data retrieval call binding the contract method 0xd5ed3933.
//
// Solidity: function finalizeTransfer(address asset, address from, address to, uint256 amount, uint256 balanceFromBefore, uint256 balanceToBefore) view returns()
func (_BendDAO *BendDAOCallerSession) FinalizeTransfer(asset common.Address, from common.Address, to common.Address, amount *big.Int, balanceFromBefore *big.Int, balanceToBefore *big.Int) error {
	return _BendDAO.Contract.FinalizeTransfer(&_BendDAO.CallOpts, asset, from, to, amount, balanceFromBefore, balanceToBefore)
}

// GetAddressesProvider is a free data retrieval call binding the contract method 0xfe65acfe.
//
// Solidity: function getAddressesProvider() view returns(address)
func (_BendDAO *BendDAOCaller) GetAddressesProvider(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BendDAO.contract.Call(opts, &out, "getAddressesProvider")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressesProvider is a free data retrieval call binding the contract method 0xfe65acfe.
//
// Solidity: function getAddressesProvider() view returns(address)
func (_BendDAO *BendDAOSession) GetAddressesProvider() (common.Address, error) {
	return _BendDAO.Contract.GetAddressesProvider(&_BendDAO.CallOpts)
}

// GetAddressesProvider is a free data retrieval call binding the contract method 0xfe65acfe.
//
// Solidity: function getAddressesProvider() view returns(address)
func (_BendDAO *BendDAOCallerSession) GetAddressesProvider() (common.Address, error) {
	return _BendDAO.Contract.GetAddressesProvider(&_BendDAO.CallOpts)
}

// GetMaxNumberOfNfts is a free data retrieval call binding the contract method 0xdd90ff38.
//
// Solidity: function getMaxNumberOfNfts() view returns(uint256)
func (_BendDAO *BendDAOCaller) GetMaxNumberOfNfts(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BendDAO.contract.Call(opts, &out, "getMaxNumberOfNfts")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxNumberOfNfts is a free data retrieval call binding the contract method 0xdd90ff38.
//
// Solidity: function getMaxNumberOfNfts() view returns(uint256)
func (_BendDAO *BendDAOSession) GetMaxNumberOfNfts() (*big.Int, error) {
	return _BendDAO.Contract.GetMaxNumberOfNfts(&_BendDAO.CallOpts)
}

// GetMaxNumberOfNfts is a free data retrieval call binding the contract method 0xdd90ff38.
//
// Solidity: function getMaxNumberOfNfts() view returns(uint256)
func (_BendDAO *BendDAOCallerSession) GetMaxNumberOfNfts() (*big.Int, error) {
	return _BendDAO.Contract.GetMaxNumberOfNfts(&_BendDAO.CallOpts)
}

// GetMaxNumberOfReserves is a free data retrieval call binding the contract method 0x08ac08b9.
//
// Solidity: function getMaxNumberOfReserves() view returns(uint256)
func (_BendDAO *BendDAOCaller) GetMaxNumberOfReserves(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BendDAO.contract.Call(opts, &out, "getMaxNumberOfReserves")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxNumberOfReserves is a free data retrieval call binding the contract method 0x08ac08b9.
//
// Solidity: function getMaxNumberOfReserves() view returns(uint256)
func (_BendDAO *BendDAOSession) GetMaxNumberOfReserves() (*big.Int, error) {
	return _BendDAO.Contract.GetMaxNumberOfReserves(&_BendDAO.CallOpts)
}

// GetMaxNumberOfReserves is a free data retrieval call binding the contract method 0x08ac08b9.
//
// Solidity: function getMaxNumberOfReserves() view returns(uint256)
func (_BendDAO *BendDAOCallerSession) GetMaxNumberOfReserves() (*big.Int, error) {
	return _BendDAO.Contract.GetMaxNumberOfReserves(&_BendDAO.CallOpts)
}

// GetNftAuctionData is a free data retrieval call binding the contract method 0xe5bceca5.
//
// Solidity: function getNftAuctionData(address nftAsset, uint256 nftTokenId) view returns(uint256 loanId, address bidderAddress, uint256 bidPrice, uint256 bidBorrowAmount, uint256 bidFine)
func (_BendDAO *BendDAOCaller) GetNftAuctionData(opts *bind.CallOpts, nftAsset common.Address, nftTokenId *big.Int) (struct {
	LoanId          *big.Int
	BidderAddress   common.Address
	BidPrice        *big.Int
	BidBorrowAmount *big.Int
	BidFine         *big.Int
}, error) {
	var out []interface{}
	err := _BendDAO.contract.Call(opts, &out, "getNftAuctionData", nftAsset, nftTokenId)

	outstruct := new(struct {
		LoanId          *big.Int
		BidderAddress   common.Address
		BidPrice        *big.Int
		BidBorrowAmount *big.Int
		BidFine         *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LoanId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BidderAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.BidPrice = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.BidBorrowAmount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.BidFine = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetNftAuctionData is a free data retrieval call binding the contract method 0xe5bceca5.
//
// Solidity: function getNftAuctionData(address nftAsset, uint256 nftTokenId) view returns(uint256 loanId, address bidderAddress, uint256 bidPrice, uint256 bidBorrowAmount, uint256 bidFine)
func (_BendDAO *BendDAOSession) GetNftAuctionData(nftAsset common.Address, nftTokenId *big.Int) (struct {
	LoanId          *big.Int
	BidderAddress   common.Address
	BidPrice        *big.Int
	BidBorrowAmount *big.Int
	BidFine         *big.Int
}, error) {
	return _BendDAO.Contract.GetNftAuctionData(&_BendDAO.CallOpts, nftAsset, nftTokenId)
}

// GetNftAuctionData is a free data retrieval call binding the contract method 0xe5bceca5.
//
// Solidity: function getNftAuctionData(address nftAsset, uint256 nftTokenId) view returns(uint256 loanId, address bidderAddress, uint256 bidPrice, uint256 bidBorrowAmount, uint256 bidFine)
func (_BendDAO *BendDAOCallerSession) GetNftAuctionData(nftAsset common.Address, nftTokenId *big.Int) (struct {
	LoanId          *big.Int
	BidderAddress   common.Address
	BidPrice        *big.Int
	BidBorrowAmount *big.Int
	BidFine         *big.Int
}, error) {
	return _BendDAO.Contract.GetNftAuctionData(&_BendDAO.CallOpts, nftAsset, nftTokenId)
}

// GetNftAuctionEndTime is a free data retrieval call binding the contract method 0x17c8a456.
//
// Solidity: function getNftAuctionEndTime(address nftAsset, uint256 nftTokenId) view returns(uint256 loanId, uint256 bidStartTimestamp, uint256 bidEndTimestamp, uint256 redeemEndTimestamp)
func (_BendDAO *BendDAOCaller) GetNftAuctionEndTime(opts *bind.CallOpts, nftAsset common.Address, nftTokenId *big.Int) (struct {
	LoanId             *big.Int
	BidStartTimestamp  *big.Int
	BidEndTimestamp    *big.Int
	RedeemEndTimestamp *big.Int
}, error) {
	var out []interface{}
	err := _BendDAO.contract.Call(opts, &out, "getNftAuctionEndTime", nftAsset, nftTokenId)

	outstruct := new(struct {
		LoanId             *big.Int
		BidStartTimestamp  *big.Int
		BidEndTimestamp    *big.Int
		RedeemEndTimestamp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LoanId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BidStartTimestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BidEndTimestamp = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.RedeemEndTimestamp = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetNftAuctionEndTime is a free data retrieval call binding the contract method 0x17c8a456.
//
// Solidity: function getNftAuctionEndTime(address nftAsset, uint256 nftTokenId) view returns(uint256 loanId, uint256 bidStartTimestamp, uint256 bidEndTimestamp, uint256 redeemEndTimestamp)
func (_BendDAO *BendDAOSession) GetNftAuctionEndTime(nftAsset common.Address, nftTokenId *big.Int) (struct {
	LoanId             *big.Int
	BidStartTimestamp  *big.Int
	BidEndTimestamp    *big.Int
	RedeemEndTimestamp *big.Int
}, error) {
	return _BendDAO.Contract.GetNftAuctionEndTime(&_BendDAO.CallOpts, nftAsset, nftTokenId)
}

// GetNftAuctionEndTime is a free data retrieval call binding the contract method 0x17c8a456.
//
// Solidity: function getNftAuctionEndTime(address nftAsset, uint256 nftTokenId) view returns(uint256 loanId, uint256 bidStartTimestamp, uint256 bidEndTimestamp, uint256 redeemEndTimestamp)
func (_BendDAO *BendDAOCallerSession) GetNftAuctionEndTime(nftAsset common.Address, nftTokenId *big.Int) (struct {
	LoanId             *big.Int
	BidStartTimestamp  *big.Int
	BidEndTimestamp    *big.Int
	RedeemEndTimestamp *big.Int
}, error) {
	return _BendDAO.Contract.GetNftAuctionEndTime(&_BendDAO.CallOpts, nftAsset, nftTokenId)
}

// GetNftCollateralData is a free data retrieval call binding the contract method 0xcc8ccdf2.
//
// Solidity: function getNftCollateralData(address nftAsset, address reserveAsset) view returns(uint256 totalCollateralInETH, uint256 totalCollateralInReserve, uint256 availableBorrowsInETH, uint256 availableBorrowsInReserve, uint256 ltv, uint256 liquidationThreshold, uint256 liquidationBonus)
func (_BendDAO *BendDAOCaller) GetNftCollateralData(opts *bind.CallOpts, nftAsset common.Address, reserveAsset common.Address) (struct {
	TotalCollateralInETH      *big.Int
	TotalCollateralInReserve  *big.Int
	AvailableBorrowsInETH     *big.Int
	AvailableBorrowsInReserve *big.Int
	Ltv                       *big.Int
	LiquidationThreshold      *big.Int
	LiquidationBonus          *big.Int
}, error) {
	var out []interface{}
	err := _BendDAO.contract.Call(opts, &out, "getNftCollateralData", nftAsset, reserveAsset)

	outstruct := new(struct {
		TotalCollateralInETH      *big.Int
		TotalCollateralInReserve  *big.Int
		AvailableBorrowsInETH     *big.Int
		AvailableBorrowsInReserve *big.Int
		Ltv                       *big.Int
		LiquidationThreshold      *big.Int
		LiquidationBonus          *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalCollateralInETH = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TotalCollateralInReserve = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.AvailableBorrowsInETH = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.AvailableBorrowsInReserve = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Ltv = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.LiquidationThreshold = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.LiquidationBonus = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetNftCollateralData is a free data retrieval call binding the contract method 0xcc8ccdf2.
//
// Solidity: function getNftCollateralData(address nftAsset, address reserveAsset) view returns(uint256 totalCollateralInETH, uint256 totalCollateralInReserve, uint256 availableBorrowsInETH, uint256 availableBorrowsInReserve, uint256 ltv, uint256 liquidationThreshold, uint256 liquidationBonus)
func (_BendDAO *BendDAOSession) GetNftCollateralData(nftAsset common.Address, reserveAsset common.Address) (struct {
	TotalCollateralInETH      *big.Int
	TotalCollateralInReserve  *big.Int
	AvailableBorrowsInETH     *big.Int
	AvailableBorrowsInReserve *big.Int
	Ltv                       *big.Int
	LiquidationThreshold      *big.Int
	LiquidationBonus          *big.Int
}, error) {
	return _BendDAO.Contract.GetNftCollateralData(&_BendDAO.CallOpts, nftAsset, reserveAsset)
}

// GetNftCollateralData is a free data retrieval call binding the contract method 0xcc8ccdf2.
//
// Solidity: function getNftCollateralData(address nftAsset, address reserveAsset) view returns(uint256 totalCollateralInETH, uint256 totalCollateralInReserve, uint256 availableBorrowsInETH, uint256 availableBorrowsInReserve, uint256 ltv, uint256 liquidationThreshold, uint256 liquidationBonus)
func (_BendDAO *BendDAOCallerSession) GetNftCollateralData(nftAsset common.Address, reserveAsset common.Address) (struct {
	TotalCollateralInETH      *big.Int
	TotalCollateralInReserve  *big.Int
	AvailableBorrowsInETH     *big.Int
	AvailableBorrowsInReserve *big.Int
	Ltv                       *big.Int
	LiquidationThreshold      *big.Int
	LiquidationBonus          *big.Int
}, error) {
	return _BendDAO.Contract.GetNftCollateralData(&_BendDAO.CallOpts, nftAsset, reserveAsset)
}

// GetNftConfiguration is a free data retrieval call binding the contract method 0x87c32dec.
//
// Solidity: function getNftConfiguration(address asset) view returns((uint256))
func (_BendDAO *BendDAOCaller) GetNftConfiguration(opts *bind.CallOpts, asset common.Address) (DataTypesNftConfigurationMap, error) {
	var out []interface{}
	err := _BendDAO.contract.Call(opts, &out, "getNftConfiguration", asset)

	if err != nil {
		return *new(DataTypesNftConfigurationMap), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesNftConfigurationMap)).(*DataTypesNftConfigurationMap)

	return out0, err

}

// GetNftConfiguration is a free data retrieval call binding the contract method 0x87c32dec.
//
// Solidity: function getNftConfiguration(address asset) view returns((uint256))
func (_BendDAO *BendDAOSession) GetNftConfiguration(asset common.Address) (DataTypesNftConfigurationMap, error) {
	return _BendDAO.Contract.GetNftConfiguration(&_BendDAO.CallOpts, asset)
}

// GetNftConfiguration is a free data retrieval call binding the contract method 0x87c32dec.
//
// Solidity: function getNftConfiguration(address asset) view returns((uint256))
func (_BendDAO *BendDAOCallerSession) GetNftConfiguration(asset common.Address) (DataTypesNftConfigurationMap, error) {
	return _BendDAO.Contract.GetNftConfiguration(&_BendDAO.CallOpts, asset)
}

// GetNftData is a free data retrieval call binding the contract method 0x77bdc0c3.
//
// Solidity: function getNftData(address asset) view returns(((uint256),address,uint8,uint256,uint256))
func (_BendDAO *BendDAOCaller) GetNftData(opts *bind.CallOpts, asset common.Address) (DataTypesNftData, error) {
	var out []interface{}
	err := _BendDAO.contract.Call(opts, &out, "getNftData", asset)

	if err != nil {
		return *new(DataTypesNftData), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesNftData)).(*DataTypesNftData)

	return out0, err

}

// GetNftData is a free data retrieval call binding the contract method 0x77bdc0c3.
//
// Solidity: function getNftData(address asset) view returns(((uint256),address,uint8,uint256,uint256))
func (_BendDAO *BendDAOSession) GetNftData(asset common.Address) (DataTypesNftData, error) {
	return _BendDAO.Contract.GetNftData(&_BendDAO.CallOpts, asset)
}

// GetNftData is a free data retrieval call binding the contract method 0x77bdc0c3.
//
// Solidity: function getNftData(address asset) view returns(((uint256),address,uint8,uint256,uint256))
func (_BendDAO *BendDAOCallerSession) GetNftData(asset common.Address) (DataTypesNftData, error) {
	return _BendDAO.Contract.GetNftData(&_BendDAO.CallOpts, asset)
}

// GetNftDebtData is a free data retrieval call binding the contract method 0xec765d3d.
//
// Solidity: function getNftDebtData(address nftAsset, uint256 nftTokenId) view returns(uint256 loanId, address reserveAsset, uint256 totalCollateral, uint256 totalDebt, uint256 availableBorrows, uint256 healthFactor)
func (_BendDAO *BendDAOCaller) GetNftDebtData(opts *bind.CallOpts, nftAsset common.Address, nftTokenId *big.Int) (struct {
	LoanId           *big.Int
	ReserveAsset     common.Address
	TotalCollateral  *big.Int
	TotalDebt        *big.Int
	AvailableBorrows *big.Int
	HealthFactor     *big.Int
}, error) {
	var out []interface{}
	err := _BendDAO.contract.Call(opts, &out, "getNftDebtData", nftAsset, nftTokenId)

	outstruct := new(struct {
		LoanId           *big.Int
		ReserveAsset     common.Address
		TotalCollateral  *big.Int
		TotalDebt        *big.Int
		AvailableBorrows *big.Int
		HealthFactor     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LoanId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ReserveAsset = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.TotalCollateral = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TotalDebt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.AvailableBorrows = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.HealthFactor = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetNftDebtData is a free data retrieval call binding the contract method 0xec765d3d.
//
// Solidity: function getNftDebtData(address nftAsset, uint256 nftTokenId) view returns(uint256 loanId, address reserveAsset, uint256 totalCollateral, uint256 totalDebt, uint256 availableBorrows, uint256 healthFactor)
func (_BendDAO *BendDAOSession) GetNftDebtData(nftAsset common.Address, nftTokenId *big.Int) (struct {
	LoanId           *big.Int
	ReserveAsset     common.Address
	TotalCollateral  *big.Int
	TotalDebt        *big.Int
	AvailableBorrows *big.Int
	HealthFactor     *big.Int
}, error) {
	return _BendDAO.Contract.GetNftDebtData(&_BendDAO.CallOpts, nftAsset, nftTokenId)
}

// GetNftDebtData is a free data retrieval call binding the contract method 0xec765d3d.
//
// Solidity: function getNftDebtData(address nftAsset, uint256 nftTokenId) view returns(uint256 loanId, address reserveAsset, uint256 totalCollateral, uint256 totalDebt, uint256 availableBorrows, uint256 healthFactor)
func (_BendDAO *BendDAOCallerSession) GetNftDebtData(nftAsset common.Address, nftTokenId *big.Int) (struct {
	LoanId           *big.Int
	ReserveAsset     common.Address
	TotalCollateral  *big.Int
	TotalDebt        *big.Int
	AvailableBorrows *big.Int
	HealthFactor     *big.Int
}, error) {
	return _BendDAO.Contract.GetNftDebtData(&_BendDAO.CallOpts, nftAsset, nftTokenId)
}

// GetNftLiquidatePrice is a free data retrieval call binding the contract method 0x798b9e3d.
//
// Solidity: function getNftLiquidatePrice(address nftAsset, uint256 nftTokenId) view returns(uint256 liquidatePrice, uint256 paybackAmount)
func (_BendDAO *BendDAOCaller) GetNftLiquidatePrice(opts *bind.CallOpts, nftAsset common.Address, nftTokenId *big.Int) (struct {
	LiquidatePrice *big.Int
	PaybackAmount  *big.Int
}, error) {
	var out []interface{}
	err := _BendDAO.contract.Call(opts, &out, "getNftLiquidatePrice", nftAsset, nftTokenId)

	outstruct := new(struct {
		LiquidatePrice *big.Int
		PaybackAmount  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LiquidatePrice = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.PaybackAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetNftLiquidatePrice is a free data retrieval call binding the contract method 0x798b9e3d.
//
// Solidity: function getNftLiquidatePrice(address nftAsset, uint256 nftTokenId) view returns(uint256 liquidatePrice, uint256 paybackAmount)
func (_BendDAO *BendDAOSession) GetNftLiquidatePrice(nftAsset common.Address, nftTokenId *big.Int) (struct {
	LiquidatePrice *big.Int
	PaybackAmount  *big.Int
}, error) {
	return _BendDAO.Contract.GetNftLiquidatePrice(&_BendDAO.CallOpts, nftAsset, nftTokenId)
}

// GetNftLiquidatePrice is a free data retrieval call binding the contract method 0x798b9e3d.
//
// Solidity: function getNftLiquidatePrice(address nftAsset, uint256 nftTokenId) view returns(uint256 liquidatePrice, uint256 paybackAmount)
func (_BendDAO *BendDAOCallerSession) GetNftLiquidatePrice(nftAsset common.Address, nftTokenId *big.Int) (struct {
	LiquidatePrice *big.Int
	PaybackAmount  *big.Int
}, error) {
	return _BendDAO.Contract.GetNftLiquidatePrice(&_BendDAO.CallOpts, nftAsset, nftTokenId)
}

// GetNftsList is a free data retrieval call binding the contract method 0x6b25c835.
//
// Solidity: function getNftsList() view returns(address[])
func (_BendDAO *BendDAOCaller) GetNftsList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _BendDAO.contract.Call(opts, &out, "getNftsList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetNftsList is a free data retrieval call binding the contract method 0x6b25c835.
//
// Solidity: function getNftsList() view returns(address[])
func (_BendDAO *BendDAOSession) GetNftsList() ([]common.Address, error) {
	return _BendDAO.Contract.GetNftsList(&_BendDAO.CallOpts)
}

// GetNftsList is a free data retrieval call binding the contract method 0x6b25c835.
//
// Solidity: function getNftsList() view returns(address[])
func (_BendDAO *BendDAOCallerSession) GetNftsList() ([]common.Address, error) {
	return _BendDAO.Contract.GetNftsList(&_BendDAO.CallOpts)
}

// GetPausedTime is a free data retrieval call binding the contract method 0x8fc42188.
//
// Solidity: function getPausedTime() view returns(uint256, uint256)
func (_BendDAO *BendDAOCaller) GetPausedTime(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _BendDAO.contract.Call(opts, &out, "getPausedTime")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetPausedTime is a free data retrieval call binding the contract method 0x8fc42188.
//
// Solidity: function getPausedTime() view returns(uint256, uint256)
func (_BendDAO *BendDAOSession) GetPausedTime() (*big.Int, *big.Int, error) {
	return _BendDAO.Contract.GetPausedTime(&_BendDAO.CallOpts)
}

// GetPausedTime is a free data retrieval call binding the contract method 0x8fc42188.
//
// Solidity: function getPausedTime() view returns(uint256, uint256)
func (_BendDAO *BendDAOCallerSession) GetPausedTime() (*big.Int, *big.Int, error) {
	return _BendDAO.Contract.GetPausedTime(&_BendDAO.CallOpts)
}

// GetReserveConfiguration is a free data retrieval call binding the contract method 0x5fc526ff.
//
// Solidity: function getReserveConfiguration(address asset) view returns((uint256))
func (_BendDAO *BendDAOCaller) GetReserveConfiguration(opts *bind.CallOpts, asset common.Address) (DataTypesReserveConfigurationMap, error) {
	var out []interface{}
	err := _BendDAO.contract.Call(opts, &out, "getReserveConfiguration", asset)

	if err != nil {
		return *new(DataTypesReserveConfigurationMap), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesReserveConfigurationMap)).(*DataTypesReserveConfigurationMap)

	return out0, err

}

// GetReserveConfiguration is a free data retrieval call binding the contract method 0x5fc526ff.
//
// Solidity: function getReserveConfiguration(address asset) view returns((uint256))
func (_BendDAO *BendDAOSession) GetReserveConfiguration(asset common.Address) (DataTypesReserveConfigurationMap, error) {
	return _BendDAO.Contract.GetReserveConfiguration(&_BendDAO.CallOpts, asset)
}

// GetReserveConfiguration is a free data retrieval call binding the contract method 0x5fc526ff.
//
// Solidity: function getReserveConfiguration(address asset) view returns((uint256))
func (_BendDAO *BendDAOCallerSession) GetReserveConfiguration(asset common.Address) (DataTypesReserveConfigurationMap, error) {
	return _BendDAO.Contract.GetReserveConfiguration(&_BendDAO.CallOpts, asset)
}

// GetReserveData is a free data retrieval call binding the contract method 0x35ea6a75.
//
// Solidity: function getReserveData(address asset) view returns(((uint256),uint128,uint128,uint128,uint128,uint40,address,address,address,uint8))
func (_BendDAO *BendDAOCaller) GetReserveData(opts *bind.CallOpts, asset common.Address) (DataTypesReserveData, error) {
	var out []interface{}
	err := _BendDAO.contract.Call(opts, &out, "getReserveData", asset)

	if err != nil {
		return *new(DataTypesReserveData), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesReserveData)).(*DataTypesReserveData)

	return out0, err

}

// GetReserveData is a free data retrieval call binding the contract method 0x35ea6a75.
//
// Solidity: function getReserveData(address asset) view returns(((uint256),uint128,uint128,uint128,uint128,uint40,address,address,address,uint8))
func (_BendDAO *BendDAOSession) GetReserveData(asset common.Address) (DataTypesReserveData, error) {
	return _BendDAO.Contract.GetReserveData(&_BendDAO.CallOpts, asset)
}

// GetReserveData is a free data retrieval call binding the contract method 0x35ea6a75.
//
// Solidity: function getReserveData(address asset) view returns(((uint256),uint128,uint128,uint128,uint128,uint40,address,address,address,uint8))
func (_BendDAO *BendDAOCallerSession) GetReserveData(asset common.Address) (DataTypesReserveData, error) {
	return _BendDAO.Contract.GetReserveData(&_BendDAO.CallOpts, asset)
}

// GetReserveNormalizedIncome is a free data retrieval call binding the contract method 0xd15e0053.
//
// Solidity: function getReserveNormalizedIncome(address asset) view returns(uint256)
func (_BendDAO *BendDAOCaller) GetReserveNormalizedIncome(opts *bind.CallOpts, asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BendDAO.contract.Call(opts, &out, "getReserveNormalizedIncome", asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReserveNormalizedIncome is a free data retrieval call binding the contract method 0xd15e0053.
//
// Solidity: function getReserveNormalizedIncome(address asset) view returns(uint256)
func (_BendDAO *BendDAOSession) GetReserveNormalizedIncome(asset common.Address) (*big.Int, error) {
	return _BendDAO.Contract.GetReserveNormalizedIncome(&_BendDAO.CallOpts, asset)
}

// GetReserveNormalizedIncome is a free data retrieval call binding the contract method 0xd15e0053.
//
// Solidity: function getReserveNormalizedIncome(address asset) view returns(uint256)
func (_BendDAO *BendDAOCallerSession) GetReserveNormalizedIncome(asset common.Address) (*big.Int, error) {
	return _BendDAO.Contract.GetReserveNormalizedIncome(&_BendDAO.CallOpts, asset)
}

// GetReserveNormalizedVariableDebt is a free data retrieval call binding the contract method 0x386497fd.
//
// Solidity: function getReserveNormalizedVariableDebt(address asset) view returns(uint256)
func (_BendDAO *BendDAOCaller) GetReserveNormalizedVariableDebt(opts *bind.CallOpts, asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BendDAO.contract.Call(opts, &out, "getReserveNormalizedVariableDebt", asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReserveNormalizedVariableDebt is a free data retrieval call binding the contract method 0x386497fd.
//
// Solidity: function getReserveNormalizedVariableDebt(address asset) view returns(uint256)
func (_BendDAO *BendDAOSession) GetReserveNormalizedVariableDebt(asset common.Address) (*big.Int, error) {
	return _BendDAO.Contract.GetReserveNormalizedVariableDebt(&_BendDAO.CallOpts, asset)
}

// GetReserveNormalizedVariableDebt is a free data retrieval call binding the contract method 0x386497fd.
//
// Solidity: function getReserveNormalizedVariableDebt(address asset) view returns(uint256)
func (_BendDAO *BendDAOCallerSession) GetReserveNormalizedVariableDebt(asset common.Address) (*big.Int, error) {
	return _BendDAO.Contract.GetReserveNormalizedVariableDebt(&_BendDAO.CallOpts, asset)
}

// GetReservesList is a free data retrieval call binding the contract method 0xd1946dbc.
//
// Solidity: function getReservesList() view returns(address[])
func (_BendDAO *BendDAOCaller) GetReservesList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _BendDAO.contract.Call(opts, &out, "getReservesList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetReservesList is a free data retrieval call binding the contract method 0xd1946dbc.
//
// Solidity: function getReservesList() view returns(address[])
func (_BendDAO *BendDAOSession) GetReservesList() ([]common.Address, error) {
	return _BendDAO.Contract.GetReservesList(&_BendDAO.CallOpts)
}

// GetReservesList is a free data retrieval call binding the contract method 0xd1946dbc.
//
// Solidity: function getReservesList() view returns(address[])
func (_BendDAO *BendDAOCallerSession) GetReservesList() ([]common.Address, error) {
	return _BendDAO.Contract.GetReservesList(&_BendDAO.CallOpts)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) pure returns(bytes4)
func (_BendDAO *BendDAOCaller) OnERC721Received(opts *bind.CallOpts, operator common.Address, from common.Address, tokenId *big.Int, data []byte) ([4]byte, error) {
	var out []interface{}
	err := _BendDAO.contract.Call(opts, &out, "onERC721Received", operator, from, tokenId, data)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) pure returns(bytes4)
func (_BendDAO *BendDAOSession) OnERC721Received(operator common.Address, from common.Address, tokenId *big.Int, data []byte) ([4]byte, error) {
	return _BendDAO.Contract.OnERC721Received(&_BendDAO.CallOpts, operator, from, tokenId, data)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) pure returns(bytes4)
func (_BendDAO *BendDAOCallerSession) OnERC721Received(operator common.Address, from common.Address, tokenId *big.Int, data []byte) ([4]byte, error) {
	return _BendDAO.Contract.OnERC721Received(&_BendDAO.CallOpts, operator, from, tokenId, data)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BendDAO *BendDAOCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BendDAO.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BendDAO *BendDAOSession) Paused() (bool, error) {
	return _BendDAO.Contract.Paused(&_BendDAO.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BendDAO *BendDAOCallerSession) Paused() (bool, error) {
	return _BendDAO.Contract.Paused(&_BendDAO.CallOpts)
}

// Auction is a paid mutator transaction binding the contract method 0xa4c0166b.
//
// Solidity: function auction(address nftAsset, uint256 nftTokenId, uint256 bidPrice, address onBehalfOf) returns()
func (_BendDAO *BendDAOTransactor) Auction(opts *bind.TransactOpts, nftAsset common.Address, nftTokenId *big.Int, bidPrice *big.Int, onBehalfOf common.Address) (*types.Transaction, error) {
	return _BendDAO.contract.Transact(opts, "auction", nftAsset, nftTokenId, bidPrice, onBehalfOf)
}

// Auction is a paid mutator transaction binding the contract method 0xa4c0166b.
//
// Solidity: function auction(address nftAsset, uint256 nftTokenId, uint256 bidPrice, address onBehalfOf) returns()
func (_BendDAO *BendDAOSession) Auction(nftAsset common.Address, nftTokenId *big.Int, bidPrice *big.Int, onBehalfOf common.Address) (*types.Transaction, error) {
	return _BendDAO.Contract.Auction(&_BendDAO.TransactOpts, nftAsset, nftTokenId, bidPrice, onBehalfOf)
}

// Auction is a paid mutator transaction binding the contract method 0xa4c0166b.
//
// Solidity: function auction(address nftAsset, uint256 nftTokenId, uint256 bidPrice, address onBehalfOf) returns()
func (_BendDAO *BendDAOTransactorSession) Auction(nftAsset common.Address, nftTokenId *big.Int, bidPrice *big.Int, onBehalfOf common.Address) (*types.Transaction, error) {
	return _BendDAO.Contract.Auction(&_BendDAO.TransactOpts, nftAsset, nftTokenId, bidPrice, onBehalfOf)
}

// BatchBorrow is a paid mutator transaction binding the contract method 0xc9fef2fe.
//
// Solidity: function batchBorrow(address[] assets, uint256[] amounts, address[] nftAssets, uint256[] nftTokenIds, address onBehalfOf, uint16 referralCode) returns()
func (_BendDAO *BendDAOTransactor) BatchBorrow(opts *bind.TransactOpts, assets []common.Address, amounts []*big.Int, nftAssets []common.Address, nftTokenIds []*big.Int, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _BendDAO.contract.Transact(opts, "batchBorrow", assets, amounts, nftAssets, nftTokenIds, onBehalfOf, referralCode)
}

// BatchBorrow is a paid mutator transaction binding the contract method 0xc9fef2fe.
//
// Solidity: function batchBorrow(address[] assets, uint256[] amounts, address[] nftAssets, uint256[] nftTokenIds, address onBehalfOf, uint16 referralCode) returns()
func (_BendDAO *BendDAOSession) BatchBorrow(assets []common.Address, amounts []*big.Int, nftAssets []common.Address, nftTokenIds []*big.Int, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _BendDAO.Contract.BatchBorrow(&_BendDAO.TransactOpts, assets, amounts, nftAssets, nftTokenIds, onBehalfOf, referralCode)
}

// BatchBorrow is a paid mutator transaction binding the contract method 0xc9fef2fe.
//
// Solidity: function batchBorrow(address[] assets, uint256[] amounts, address[] nftAssets, uint256[] nftTokenIds, address onBehalfOf, uint16 referralCode) returns()
func (_BendDAO *BendDAOTransactorSession) BatchBorrow(assets []common.Address, amounts []*big.Int, nftAssets []common.Address, nftTokenIds []*big.Int, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _BendDAO.Contract.BatchBorrow(&_BendDAO.TransactOpts, assets, amounts, nftAssets, nftTokenIds, onBehalfOf, referralCode)
}

// BatchRepay is a paid mutator transaction binding the contract method 0x5bc5bbf1.
//
// Solidity: function batchRepay(address[] nftAssets, uint256[] nftTokenIds, uint256[] amounts) returns(uint256[], bool[])
func (_BendDAO *BendDAOTransactor) BatchRepay(opts *bind.TransactOpts, nftAssets []common.Address, nftTokenIds []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	return _BendDAO.contract.Transact(opts, "batchRepay", nftAssets, nftTokenIds, amounts)
}

// BatchRepay is a paid mutator transaction binding the contract method 0x5bc5bbf1.
//
// Solidity: function batchRepay(address[] nftAssets, uint256[] nftTokenIds, uint256[] amounts) returns(uint256[], bool[])
func (_BendDAO *BendDAOSession) BatchRepay(nftAssets []common.Address, nftTokenIds []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	return _BendDAO.Contract.BatchRepay(&_BendDAO.TransactOpts, nftAssets, nftTokenIds, amounts)
}

// BatchRepay is a paid mutator transaction binding the contract method 0x5bc5bbf1.
//
// Solidity: function batchRepay(address[] nftAssets, uint256[] nftTokenIds, uint256[] amounts) returns(uint256[], bool[])
func (_BendDAO *BendDAOTransactorSession) BatchRepay(nftAssets []common.Address, nftTokenIds []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	return _BendDAO.Contract.BatchRepay(&_BendDAO.TransactOpts, nftAssets, nftTokenIds, amounts)
}

// Borrow is a paid mutator transaction binding the contract method 0xb6529aee.
//
// Solidity: function borrow(address asset, uint256 amount, address nftAsset, uint256 nftTokenId, address onBehalfOf, uint16 referralCode) returns()
func (_BendDAO *BendDAOTransactor) Borrow(opts *bind.TransactOpts, asset common.Address, amount *big.Int, nftAsset common.Address, nftTokenId *big.Int, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _BendDAO.contract.Transact(opts, "borrow", asset, amount, nftAsset, nftTokenId, onBehalfOf, referralCode)
}

// Borrow is a paid mutator transaction binding the contract method 0xb6529aee.
//
// Solidity: function borrow(address asset, uint256 amount, address nftAsset, uint256 nftTokenId, address onBehalfOf, uint16 referralCode) returns()
func (_BendDAO *BendDAOSession) Borrow(asset common.Address, amount *big.Int, nftAsset common.Address, nftTokenId *big.Int, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _BendDAO.Contract.Borrow(&_BendDAO.TransactOpts, asset, amount, nftAsset, nftTokenId, onBehalfOf, referralCode)
}

// Borrow is a paid mutator transaction binding the contract method 0xb6529aee.
//
// Solidity: function borrow(address asset, uint256 amount, address nftAsset, uint256 nftTokenId, address onBehalfOf, uint16 referralCode) returns()
func (_BendDAO *BendDAOTransactorSession) Borrow(asset common.Address, amount *big.Int, nftAsset common.Address, nftTokenId *big.Int, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _BendDAO.Contract.Borrow(&_BendDAO.TransactOpts, asset, amount, nftAsset, nftTokenId, onBehalfOf, referralCode)
}

// Deposit is a paid mutator transaction binding the contract method 0xe8eda9df.
//
// Solidity: function deposit(address asset, uint256 amount, address onBehalfOf, uint16 referralCode) returns()
func (_BendDAO *BendDAOTransactor) Deposit(opts *bind.TransactOpts, asset common.Address, amount *big.Int, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _BendDAO.contract.Transact(opts, "deposit", asset, amount, onBehalfOf, referralCode)
}

// Deposit is a paid mutator transaction binding the contract method 0xe8eda9df.
//
// Solidity: function deposit(address asset, uint256 amount, address onBehalfOf, uint16 referralCode) returns()
func (_BendDAO *BendDAOSession) Deposit(asset common.Address, amount *big.Int, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _BendDAO.Contract.Deposit(&_BendDAO.TransactOpts, asset, amount, onBehalfOf, referralCode)
}

// Deposit is a paid mutator transaction binding the contract method 0xe8eda9df.
//
// Solidity: function deposit(address asset, uint256 amount, address onBehalfOf, uint16 referralCode) returns()
func (_BendDAO *BendDAOTransactorSession) Deposit(asset common.Address, amount *big.Int, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _BendDAO.Contract.Deposit(&_BendDAO.TransactOpts, asset, amount, onBehalfOf, referralCode)
}

// InitNft is a paid mutator transaction binding the contract method 0x873e4dab.
//
// Solidity: function initNft(address asset, address bNftAddress) returns()
func (_BendDAO *BendDAOTransactor) InitNft(opts *bind.TransactOpts, asset common.Address, bNftAddress common.Address) (*types.Transaction, error) {
	return _BendDAO.contract.Transact(opts, "initNft", asset, bNftAddress)
}

// InitNft is a paid mutator transaction binding the contract method 0x873e4dab.
//
// Solidity: function initNft(address asset, address bNftAddress) returns()
func (_BendDAO *BendDAOSession) InitNft(asset common.Address, bNftAddress common.Address) (*types.Transaction, error) {
	return _BendDAO.Contract.InitNft(&_BendDAO.TransactOpts, asset, bNftAddress)
}

// InitNft is a paid mutator transaction binding the contract method 0x873e4dab.
//
// Solidity: function initNft(address asset, address bNftAddress) returns()
func (_BendDAO *BendDAOTransactorSession) InitNft(asset common.Address, bNftAddress common.Address) (*types.Transaction, error) {
	return _BendDAO.Contract.InitNft(&_BendDAO.TransactOpts, asset, bNftAddress)
}

// InitReserve is a paid mutator transaction binding the contract method 0x8bd25677.
//
// Solidity: function initReserve(address asset, address bTokenAddress, address debtTokenAddress, address interestRateAddress) returns()
func (_BendDAO *BendDAOTransactor) InitReserve(opts *bind.TransactOpts, asset common.Address, bTokenAddress common.Address, debtTokenAddress common.Address, interestRateAddress common.Address) (*types.Transaction, error) {
	return _BendDAO.contract.Transact(opts, "initReserve", asset, bTokenAddress, debtTokenAddress, interestRateAddress)
}

// InitReserve is a paid mutator transaction binding the contract method 0x8bd25677.
//
// Solidity: function initReserve(address asset, address bTokenAddress, address debtTokenAddress, address interestRateAddress) returns()
func (_BendDAO *BendDAOSession) InitReserve(asset common.Address, bTokenAddress common.Address, debtTokenAddress common.Address, interestRateAddress common.Address) (*types.Transaction, error) {
	return _BendDAO.Contract.InitReserve(&_BendDAO.TransactOpts, asset, bTokenAddress, debtTokenAddress, interestRateAddress)
}

// InitReserve is a paid mutator transaction binding the contract method 0x8bd25677.
//
// Solidity: function initReserve(address asset, address bTokenAddress, address debtTokenAddress, address interestRateAddress) returns()
func (_BendDAO *BendDAOTransactorSession) InitReserve(asset common.Address, bTokenAddress common.Address, debtTokenAddress common.Address, interestRateAddress common.Address) (*types.Transaction, error) {
	return _BendDAO.Contract.InitReserve(&_BendDAO.TransactOpts, asset, bTokenAddress, debtTokenAddress, interestRateAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address provider) returns()
func (_BendDAO *BendDAOTransactor) Initialize(opts *bind.TransactOpts, provider common.Address) (*types.Transaction, error) {
	return _BendDAO.contract.Transact(opts, "initialize", provider)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address provider) returns()
func (_BendDAO *BendDAOSession) Initialize(provider common.Address) (*types.Transaction, error) {
	return _BendDAO.Contract.Initialize(&_BendDAO.TransactOpts, provider)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address provider) returns()
func (_BendDAO *BendDAOTransactorSession) Initialize(provider common.Address) (*types.Transaction, error) {
	return _BendDAO.Contract.Initialize(&_BendDAO.TransactOpts, provider)
}

// Liquidate is a paid mutator transaction binding the contract method 0x0710285c.
//
// Solidity: function liquidate(address nftAsset, uint256 nftTokenId, uint256 amount) returns(uint256)
func (_BendDAO *BendDAOTransactor) Liquidate(opts *bind.TransactOpts, nftAsset common.Address, nftTokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _BendDAO.contract.Transact(opts, "liquidate", nftAsset, nftTokenId, amount)
}

// Liquidate is a paid mutator transaction binding the contract method 0x0710285c.
//
// Solidity: function liquidate(address nftAsset, uint256 nftTokenId, uint256 amount) returns(uint256)
func (_BendDAO *BendDAOSession) Liquidate(nftAsset common.Address, nftTokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _BendDAO.Contract.Liquidate(&_BendDAO.TransactOpts, nftAsset, nftTokenId, amount)
}

// Liquidate is a paid mutator transaction binding the contract method 0x0710285c.
//
// Solidity: function liquidate(address nftAsset, uint256 nftTokenId, uint256 amount) returns(uint256)
func (_BendDAO *BendDAOTransactorSession) Liquidate(nftAsset common.Address, nftTokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _BendDAO.Contract.Liquidate(&_BendDAO.TransactOpts, nftAsset, nftTokenId, amount)
}

// Redeem is a paid mutator transaction binding the contract method 0xea2092f3.
//
// Solidity: function redeem(address nftAsset, uint256 nftTokenId, uint256 amount, uint256 bidFine) returns(uint256)
func (_BendDAO *BendDAOTransactor) Redeem(opts *bind.TransactOpts, nftAsset common.Address, nftTokenId *big.Int, amount *big.Int, bidFine *big.Int) (*types.Transaction, error) {
	return _BendDAO.contract.Transact(opts, "redeem", nftAsset, nftTokenId, amount, bidFine)
}

// Redeem is a paid mutator transaction binding the contract method 0xea2092f3.
//
// Solidity: function redeem(address nftAsset, uint256 nftTokenId, uint256 amount, uint256 bidFine) returns(uint256)
func (_BendDAO *BendDAOSession) Redeem(nftAsset common.Address, nftTokenId *big.Int, amount *big.Int, bidFine *big.Int) (*types.Transaction, error) {
	return _BendDAO.Contract.Redeem(&_BendDAO.TransactOpts, nftAsset, nftTokenId, amount, bidFine)
}

// Redeem is a paid mutator transaction binding the contract method 0xea2092f3.
//
// Solidity: function redeem(address nftAsset, uint256 nftTokenId, uint256 amount, uint256 bidFine) returns(uint256)
func (_BendDAO *BendDAOTransactorSession) Redeem(nftAsset common.Address, nftTokenId *big.Int, amount *big.Int, bidFine *big.Int) (*types.Transaction, error) {
	return _BendDAO.Contract.Redeem(&_BendDAO.TransactOpts, nftAsset, nftTokenId, amount, bidFine)
}

// Repay is a paid mutator transaction binding the contract method 0x8cd2e0c7.
//
// Solidity: function repay(address nftAsset, uint256 nftTokenId, uint256 amount) returns(uint256, bool)
func (_BendDAO *BendDAOTransactor) Repay(opts *bind.TransactOpts, nftAsset common.Address, nftTokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _BendDAO.contract.Transact(opts, "repay", nftAsset, nftTokenId, amount)
}

// Repay is a paid mutator transaction binding the contract method 0x8cd2e0c7.
//
// Solidity: function repay(address nftAsset, uint256 nftTokenId, uint256 amount) returns(uint256, bool)
func (_BendDAO *BendDAOSession) Repay(nftAsset common.Address, nftTokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _BendDAO.Contract.Repay(&_BendDAO.TransactOpts, nftAsset, nftTokenId, amount)
}

// Repay is a paid mutator transaction binding the contract method 0x8cd2e0c7.
//
// Solidity: function repay(address nftAsset, uint256 nftTokenId, uint256 amount) returns(uint256, bool)
func (_BendDAO *BendDAOTransactorSession) Repay(nftAsset common.Address, nftTokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _BendDAO.Contract.Repay(&_BendDAO.TransactOpts, nftAsset, nftTokenId, amount)
}

// SetMaxNumberOfNfts is a paid mutator transaction binding the contract method 0xfdff6f26.
//
// Solidity: function setMaxNumberOfNfts(uint256 val) returns()
func (_BendDAO *BendDAOTransactor) SetMaxNumberOfNfts(opts *bind.TransactOpts, val *big.Int) (*types.Transaction, error) {
	return _BendDAO.contract.Transact(opts, "setMaxNumberOfNfts", val)
}

// SetMaxNumberOfNfts is a paid mutator transaction binding the contract method 0xfdff6f26.
//
// Solidity: function setMaxNumberOfNfts(uint256 val) returns()
func (_BendDAO *BendDAOSession) SetMaxNumberOfNfts(val *big.Int) (*types.Transaction, error) {
	return _BendDAO.Contract.SetMaxNumberOfNfts(&_BendDAO.TransactOpts, val)
}

// SetMaxNumberOfNfts is a paid mutator transaction binding the contract method 0xfdff6f26.
//
// Solidity: function setMaxNumberOfNfts(uint256 val) returns()
func (_BendDAO *BendDAOTransactorSession) SetMaxNumberOfNfts(val *big.Int) (*types.Transaction, error) {
	return _BendDAO.Contract.SetMaxNumberOfNfts(&_BendDAO.TransactOpts, val)
}

// SetMaxNumberOfReserves is a paid mutator transaction binding the contract method 0x746c35a2.
//
// Solidity: function setMaxNumberOfReserves(uint256 val) returns()
func (_BendDAO *BendDAOTransactor) SetMaxNumberOfReserves(opts *bind.TransactOpts, val *big.Int) (*types.Transaction, error) {
	return _BendDAO.contract.Transact(opts, "setMaxNumberOfReserves", val)
}

// SetMaxNumberOfReserves is a paid mutator transaction binding the contract method 0x746c35a2.
//
// Solidity: function setMaxNumberOfReserves(uint256 val) returns()
func (_BendDAO *BendDAOSession) SetMaxNumberOfReserves(val *big.Int) (*types.Transaction, error) {
	return _BendDAO.Contract.SetMaxNumberOfReserves(&_BendDAO.TransactOpts, val)
}

// SetMaxNumberOfReserves is a paid mutator transaction binding the contract method 0x746c35a2.
//
// Solidity: function setMaxNumberOfReserves(uint256 val) returns()
func (_BendDAO *BendDAOTransactorSession) SetMaxNumberOfReserves(val *big.Int) (*types.Transaction, error) {
	return _BendDAO.Contract.SetMaxNumberOfReserves(&_BendDAO.TransactOpts, val)
}

// SetNftConfiguration is a paid mutator transaction binding the contract method 0x83c8afd7.
//
// Solidity: function setNftConfiguration(address asset, uint256 configuration) returns()
func (_BendDAO *BendDAOTransactor) SetNftConfiguration(opts *bind.TransactOpts, asset common.Address, configuration *big.Int) (*types.Transaction, error) {
	return _BendDAO.contract.Transact(opts, "setNftConfiguration", asset, configuration)
}

// SetNftConfiguration is a paid mutator transaction binding the contract method 0x83c8afd7.
//
// Solidity: function setNftConfiguration(address asset, uint256 configuration) returns()
func (_BendDAO *BendDAOSession) SetNftConfiguration(asset common.Address, configuration *big.Int) (*types.Transaction, error) {
	return _BendDAO.Contract.SetNftConfiguration(&_BendDAO.TransactOpts, asset, configuration)
}

// SetNftConfiguration is a paid mutator transaction binding the contract method 0x83c8afd7.
//
// Solidity: function setNftConfiguration(address asset, uint256 configuration) returns()
func (_BendDAO *BendDAOTransactorSession) SetNftConfiguration(asset common.Address, configuration *big.Int) (*types.Transaction, error) {
	return _BendDAO.Contract.SetNftConfiguration(&_BendDAO.TransactOpts, asset, configuration)
}

// SetNftMaxSupplyAndTokenId is a paid mutator transaction binding the contract method 0xdb78f216.
//
// Solidity: function setNftMaxSupplyAndTokenId(address asset, uint256 maxSupply, uint256 maxTokenId) returns()
func (_BendDAO *BendDAOTransactor) SetNftMaxSupplyAndTokenId(opts *bind.TransactOpts, asset common.Address, maxSupply *big.Int, maxTokenId *big.Int) (*types.Transaction, error) {
	return _BendDAO.contract.Transact(opts, "setNftMaxSupplyAndTokenId", asset, maxSupply, maxTokenId)
}

// SetNftMaxSupplyAndTokenId is a paid mutator transaction binding the contract method 0xdb78f216.
//
// Solidity: function setNftMaxSupplyAndTokenId(address asset, uint256 maxSupply, uint256 maxTokenId) returns()
func (_BendDAO *BendDAOSession) SetNftMaxSupplyAndTokenId(asset common.Address, maxSupply *big.Int, maxTokenId *big.Int) (*types.Transaction, error) {
	return _BendDAO.Contract.SetNftMaxSupplyAndTokenId(&_BendDAO.TransactOpts, asset, maxSupply, maxTokenId)
}

// SetNftMaxSupplyAndTokenId is a paid mutator transaction binding the contract method 0xdb78f216.
//
// Solidity: function setNftMaxSupplyAndTokenId(address asset, uint256 maxSupply, uint256 maxTokenId) returns()
func (_BendDAO *BendDAOTransactorSession) SetNftMaxSupplyAndTokenId(asset common.Address, maxSupply *big.Int, maxTokenId *big.Int) (*types.Transaction, error) {
	return _BendDAO.Contract.SetNftMaxSupplyAndTokenId(&_BendDAO.TransactOpts, asset, maxSupply, maxTokenId)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool val) returns()
func (_BendDAO *BendDAOTransactor) SetPause(opts *bind.TransactOpts, val bool) (*types.Transaction, error) {
	return _BendDAO.contract.Transact(opts, "setPause", val)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool val) returns()
func (_BendDAO *BendDAOSession) SetPause(val bool) (*types.Transaction, error) {
	return _BendDAO.Contract.SetPause(&_BendDAO.TransactOpts, val)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool val) returns()
func (_BendDAO *BendDAOTransactorSession) SetPause(val bool) (*types.Transaction, error) {
	return _BendDAO.Contract.SetPause(&_BendDAO.TransactOpts, val)
}

// SetPausedTime is a paid mutator transaction binding the contract method 0x2f923ff7.
//
// Solidity: function setPausedTime(uint256 startTime, uint256 durationTime) returns()
func (_BendDAO *BendDAOTransactor) SetPausedTime(opts *bind.TransactOpts, startTime *big.Int, durationTime *big.Int) (*types.Transaction, error) {
	return _BendDAO.contract.Transact(opts, "setPausedTime", startTime, durationTime)
}

// SetPausedTime is a paid mutator transaction binding the contract method 0x2f923ff7.
//
// Solidity: function setPausedTime(uint256 startTime, uint256 durationTime) returns()
func (_BendDAO *BendDAOSession) SetPausedTime(startTime *big.Int, durationTime *big.Int) (*types.Transaction, error) {
	return _BendDAO.Contract.SetPausedTime(&_BendDAO.TransactOpts, startTime, durationTime)
}

// SetPausedTime is a paid mutator transaction binding the contract method 0x2f923ff7.
//
// Solidity: function setPausedTime(uint256 startTime, uint256 durationTime) returns()
func (_BendDAO *BendDAOTransactorSession) SetPausedTime(startTime *big.Int, durationTime *big.Int) (*types.Transaction, error) {
	return _BendDAO.Contract.SetPausedTime(&_BendDAO.TransactOpts, startTime, durationTime)
}

// SetReserveConfiguration is a paid mutator transaction binding the contract method 0x43f0f733.
//
// Solidity: function setReserveConfiguration(address asset, uint256 configuration) returns()
func (_BendDAO *BendDAOTransactor) SetReserveConfiguration(opts *bind.TransactOpts, asset common.Address, configuration *big.Int) (*types.Transaction, error) {
	return _BendDAO.contract.Transact(opts, "setReserveConfiguration", asset, configuration)
}

// SetReserveConfiguration is a paid mutator transaction binding the contract method 0x43f0f733.
//
// Solidity: function setReserveConfiguration(address asset, uint256 configuration) returns()
func (_BendDAO *BendDAOSession) SetReserveConfiguration(asset common.Address, configuration *big.Int) (*types.Transaction, error) {
	return _BendDAO.Contract.SetReserveConfiguration(&_BendDAO.TransactOpts, asset, configuration)
}

// SetReserveConfiguration is a paid mutator transaction binding the contract method 0x43f0f733.
//
// Solidity: function setReserveConfiguration(address asset, uint256 configuration) returns()
func (_BendDAO *BendDAOTransactorSession) SetReserveConfiguration(asset common.Address, configuration *big.Int) (*types.Transaction, error) {
	return _BendDAO.Contract.SetReserveConfiguration(&_BendDAO.TransactOpts, asset, configuration)
}

// SetReserveInterestRateAddress is a paid mutator transaction binding the contract method 0x83b1555f.
//
// Solidity: function setReserveInterestRateAddress(address asset, address rateAddress) returns()
func (_BendDAO *BendDAOTransactor) SetReserveInterestRateAddress(opts *bind.TransactOpts, asset common.Address, rateAddress common.Address) (*types.Transaction, error) {
	return _BendDAO.contract.Transact(opts, "setReserveInterestRateAddress", asset, rateAddress)
}

// SetReserveInterestRateAddress is a paid mutator transaction binding the contract method 0x83b1555f.
//
// Solidity: function setReserveInterestRateAddress(address asset, address rateAddress) returns()
func (_BendDAO *BendDAOSession) SetReserveInterestRateAddress(asset common.Address, rateAddress common.Address) (*types.Transaction, error) {
	return _BendDAO.Contract.SetReserveInterestRateAddress(&_BendDAO.TransactOpts, asset, rateAddress)
}

// SetReserveInterestRateAddress is a paid mutator transaction binding the contract method 0x83b1555f.
//
// Solidity: function setReserveInterestRateAddress(address asset, address rateAddress) returns()
func (_BendDAO *BendDAOTransactorSession) SetReserveInterestRateAddress(asset common.Address, rateAddress common.Address) (*types.Transaction, error) {
	return _BendDAO.Contract.SetReserveInterestRateAddress(&_BendDAO.TransactOpts, asset, rateAddress)
}

// Withdraw is a paid mutator transaction binding the contract method 0x69328dec.
//
// Solidity: function withdraw(address asset, uint256 amount, address to) returns(uint256)
func (_BendDAO *BendDAOTransactor) Withdraw(opts *bind.TransactOpts, asset common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _BendDAO.contract.Transact(opts, "withdraw", asset, amount, to)
}

// Withdraw is a paid mutator transaction binding the contract method 0x69328dec.
//
// Solidity: function withdraw(address asset, uint256 amount, address to) returns(uint256)
func (_BendDAO *BendDAOSession) Withdraw(asset common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _BendDAO.Contract.Withdraw(&_BendDAO.TransactOpts, asset, amount, to)
}

// Withdraw is a paid mutator transaction binding the contract method 0x69328dec.
//
// Solidity: function withdraw(address asset, uint256 amount, address to) returns(uint256)
func (_BendDAO *BendDAOTransactorSession) Withdraw(asset common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _BendDAO.Contract.Withdraw(&_BendDAO.TransactOpts, asset, amount, to)
}

// BendDAOAuctionIterator is returned from FilterAuction and is used to iterate over the raw logs and unpacked data for Auction events raised by the BendDAO contract.
type BendDAOAuctionIterator struct {
	Event *BendDAOAuction // Event containing the contract specifics and raw log

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
func (it *BendDAOAuctionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BendDAOAuction)
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
		it.Event = new(BendDAOAuction)
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
func (it *BendDAOAuctionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BendDAOAuctionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BendDAOAuction represents a Auction event raised by the BendDAO contract.
type BendDAOAuction struct {
	User       common.Address
	Reserve    common.Address
	BidPrice   *big.Int
	NftAsset   common.Address
	NftTokenId *big.Int
	OnBehalfOf common.Address
	Borrower   common.Address
	LoanId     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAuction is a free log retrieval operation binding the contract event 0xd4c7449fa0ea241233dd0a9e78a940879918f95e0caa34e0399a7d2813c8efba.
//
// Solidity: event Auction(address user, address indexed reserve, uint256 bidPrice, address indexed nftAsset, uint256 nftTokenId, address onBehalfOf, address indexed borrower, uint256 loanId)
func (_BendDAO *BendDAOFilterer) FilterAuction(opts *bind.FilterOpts, reserve []common.Address, nftAsset []common.Address, borrower []common.Address) (*BendDAOAuctionIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	var nftAssetRule []interface{}
	for _, nftAssetItem := range nftAsset {
		nftAssetRule = append(nftAssetRule, nftAssetItem)
	}

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _BendDAO.contract.FilterLogs(opts, "Auction", reserveRule, nftAssetRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return &BendDAOAuctionIterator{contract: _BendDAO.contract, event: "Auction", logs: logs, sub: sub}, nil
}

// WatchAuction is a free log subscription operation binding the contract event 0xd4c7449fa0ea241233dd0a9e78a940879918f95e0caa34e0399a7d2813c8efba.
//
// Solidity: event Auction(address user, address indexed reserve, uint256 bidPrice, address indexed nftAsset, uint256 nftTokenId, address onBehalfOf, address indexed borrower, uint256 loanId)
func (_BendDAO *BendDAOFilterer) WatchAuction(opts *bind.WatchOpts, sink chan<- *BendDAOAuction, reserve []common.Address, nftAsset []common.Address, borrower []common.Address) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	var nftAssetRule []interface{}
	for _, nftAssetItem := range nftAsset {
		nftAssetRule = append(nftAssetRule, nftAssetItem)
	}

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _BendDAO.contract.WatchLogs(opts, "Auction", reserveRule, nftAssetRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BendDAOAuction)
				if err := _BendDAO.contract.UnpackLog(event, "Auction", log); err != nil {
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

// ParseAuction is a log parse operation binding the contract event 0xd4c7449fa0ea241233dd0a9e78a940879918f95e0caa34e0399a7d2813c8efba.
//
// Solidity: event Auction(address user, address indexed reserve, uint256 bidPrice, address indexed nftAsset, uint256 nftTokenId, address onBehalfOf, address indexed borrower, uint256 loanId)
func (_BendDAO *BendDAOFilterer) ParseAuction(log types.Log) (*BendDAOAuction, error) {
	event := new(BendDAOAuction)
	if err := _BendDAO.contract.UnpackLog(event, "Auction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BendDAOBorrowIterator is returned from FilterBorrow and is used to iterate over the raw logs and unpacked data for Borrow events raised by the BendDAO contract.
type BendDAOBorrowIterator struct {
	Event *BendDAOBorrow // Event containing the contract specifics and raw log

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
func (it *BendDAOBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BendDAOBorrow)
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
		it.Event = new(BendDAOBorrow)
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
func (it *BendDAOBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BendDAOBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BendDAOBorrow represents a Borrow event raised by the BendDAO contract.
type BendDAOBorrow struct {
	User       common.Address
	Reserve    common.Address
	Amount     *big.Int
	NftAsset   common.Address
	NftTokenId *big.Int
	OnBehalfOf common.Address
	BorrowRate *big.Int
	LoanId     *big.Int
	Referral   uint16
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBorrow is a free log retrieval operation binding the contract event 0xcfb3a669117d9dc90f0d3521228dc9fe67c5102dde205ef16fe9a1f81be698d5.
//
// Solidity: event Borrow(address user, address indexed reserve, uint256 amount, address nftAsset, uint256 nftTokenId, address indexed onBehalfOf, uint256 borrowRate, uint256 loanId, uint16 indexed referral)
func (_BendDAO *BendDAOFilterer) FilterBorrow(opts *bind.FilterOpts, reserve []common.Address, onBehalfOf []common.Address, referral []uint16) (*BendDAOBorrowIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	var referralRule []interface{}
	for _, referralItem := range referral {
		referralRule = append(referralRule, referralItem)
	}

	logs, sub, err := _BendDAO.contract.FilterLogs(opts, "Borrow", reserveRule, onBehalfOfRule, referralRule)
	if err != nil {
		return nil, err
	}
	return &BendDAOBorrowIterator{contract: _BendDAO.contract, event: "Borrow", logs: logs, sub: sub}, nil
}

// WatchBorrow is a free log subscription operation binding the contract event 0xcfb3a669117d9dc90f0d3521228dc9fe67c5102dde205ef16fe9a1f81be698d5.
//
// Solidity: event Borrow(address user, address indexed reserve, uint256 amount, address nftAsset, uint256 nftTokenId, address indexed onBehalfOf, uint256 borrowRate, uint256 loanId, uint16 indexed referral)
func (_BendDAO *BendDAOFilterer) WatchBorrow(opts *bind.WatchOpts, sink chan<- *BendDAOBorrow, reserve []common.Address, onBehalfOf []common.Address, referral []uint16) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	var referralRule []interface{}
	for _, referralItem := range referral {
		referralRule = append(referralRule, referralItem)
	}

	logs, sub, err := _BendDAO.contract.WatchLogs(opts, "Borrow", reserveRule, onBehalfOfRule, referralRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BendDAOBorrow)
				if err := _BendDAO.contract.UnpackLog(event, "Borrow", log); err != nil {
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

// ParseBorrow is a log parse operation binding the contract event 0xcfb3a669117d9dc90f0d3521228dc9fe67c5102dde205ef16fe9a1f81be698d5.
//
// Solidity: event Borrow(address user, address indexed reserve, uint256 amount, address nftAsset, uint256 nftTokenId, address indexed onBehalfOf, uint256 borrowRate, uint256 loanId, uint16 indexed referral)
func (_BendDAO *BendDAOFilterer) ParseBorrow(log types.Log) (*BendDAOBorrow, error) {
	event := new(BendDAOBorrow)
	if err := _BendDAO.contract.UnpackLog(event, "Borrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BendDAODepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the BendDAO contract.
type BendDAODepositIterator struct {
	Event *BendDAODeposit // Event containing the contract specifics and raw log

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
func (it *BendDAODepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BendDAODeposit)
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
		it.Event = new(BendDAODeposit)
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
func (it *BendDAODepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BendDAODepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BendDAODeposit represents a Deposit event raised by the BendDAO contract.
type BendDAODeposit struct {
	User       common.Address
	Reserve    common.Address
	Amount     *big.Int
	OnBehalfOf common.Address
	Referral   uint16
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x443ff2d25883a4800d36062db52ca3dd7ced05bd8627c8a6a37f8699715b5431.
//
// Solidity: event Deposit(address user, address indexed reserve, uint256 amount, address indexed onBehalfOf, uint16 indexed referral)
func (_BendDAO *BendDAOFilterer) FilterDeposit(opts *bind.FilterOpts, reserve []common.Address, onBehalfOf []common.Address, referral []uint16) (*BendDAODepositIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}
	var referralRule []interface{}
	for _, referralItem := range referral {
		referralRule = append(referralRule, referralItem)
	}

	logs, sub, err := _BendDAO.contract.FilterLogs(opts, "Deposit", reserveRule, onBehalfOfRule, referralRule)
	if err != nil {
		return nil, err
	}
	return &BendDAODepositIterator{contract: _BendDAO.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x443ff2d25883a4800d36062db52ca3dd7ced05bd8627c8a6a37f8699715b5431.
//
// Solidity: event Deposit(address user, address indexed reserve, uint256 amount, address indexed onBehalfOf, uint16 indexed referral)
func (_BendDAO *BendDAOFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *BendDAODeposit, reserve []common.Address, onBehalfOf []common.Address, referral []uint16) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}
	var referralRule []interface{}
	for _, referralItem := range referral {
		referralRule = append(referralRule, referralItem)
	}

	logs, sub, err := _BendDAO.contract.WatchLogs(opts, "Deposit", reserveRule, onBehalfOfRule, referralRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BendDAODeposit)
				if err := _BendDAO.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x443ff2d25883a4800d36062db52ca3dd7ced05bd8627c8a6a37f8699715b5431.
//
// Solidity: event Deposit(address user, address indexed reserve, uint256 amount, address indexed onBehalfOf, uint16 indexed referral)
func (_BendDAO *BendDAOFilterer) ParseDeposit(log types.Log) (*BendDAODeposit, error) {
	event := new(BendDAODeposit)
	if err := _BendDAO.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BendDAOLiquidateIterator is returned from FilterLiquidate and is used to iterate over the raw logs and unpacked data for Liquidate events raised by the BendDAO contract.
type BendDAOLiquidateIterator struct {
	Event *BendDAOLiquidate // Event containing the contract specifics and raw log

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
func (it *BendDAOLiquidateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BendDAOLiquidate)
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
		it.Event = new(BendDAOLiquidate)
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
func (it *BendDAOLiquidateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BendDAOLiquidateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BendDAOLiquidate represents a Liquidate event raised by the BendDAO contract.
type BendDAOLiquidate struct {
	User         common.Address
	Reserve      common.Address
	RepayAmount  *big.Int
	RemainAmount *big.Int
	NftAsset     common.Address
	NftTokenId   *big.Int
	Borrower     common.Address
	LoanId       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterLiquidate is a free log retrieval operation binding the contract event 0xf028795898a18c6fc88094dc5671c6a79d5dc3458c44015e9299fbc6c6268cf8.
//
// Solidity: event Liquidate(address user, address indexed reserve, uint256 repayAmount, uint256 remainAmount, address indexed nftAsset, uint256 nftTokenId, address indexed borrower, uint256 loanId)
func (_BendDAO *BendDAOFilterer) FilterLiquidate(opts *bind.FilterOpts, reserve []common.Address, nftAsset []common.Address, borrower []common.Address) (*BendDAOLiquidateIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	var nftAssetRule []interface{}
	for _, nftAssetItem := range nftAsset {
		nftAssetRule = append(nftAssetRule, nftAssetItem)
	}

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _BendDAO.contract.FilterLogs(opts, "Liquidate", reserveRule, nftAssetRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return &BendDAOLiquidateIterator{contract: _BendDAO.contract, event: "Liquidate", logs: logs, sub: sub}, nil
}

// WatchLiquidate is a free log subscription operation binding the contract event 0xf028795898a18c6fc88094dc5671c6a79d5dc3458c44015e9299fbc6c6268cf8.
//
// Solidity: event Liquidate(address user, address indexed reserve, uint256 repayAmount, uint256 remainAmount, address indexed nftAsset, uint256 nftTokenId, address indexed borrower, uint256 loanId)
func (_BendDAO *BendDAOFilterer) WatchLiquidate(opts *bind.WatchOpts, sink chan<- *BendDAOLiquidate, reserve []common.Address, nftAsset []common.Address, borrower []common.Address) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	var nftAssetRule []interface{}
	for _, nftAssetItem := range nftAsset {
		nftAssetRule = append(nftAssetRule, nftAssetItem)
	}

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _BendDAO.contract.WatchLogs(opts, "Liquidate", reserveRule, nftAssetRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BendDAOLiquidate)
				if err := _BendDAO.contract.UnpackLog(event, "Liquidate", log); err != nil {
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

// ParseLiquidate is a log parse operation binding the contract event 0xf028795898a18c6fc88094dc5671c6a79d5dc3458c44015e9299fbc6c6268cf8.
//
// Solidity: event Liquidate(address user, address indexed reserve, uint256 repayAmount, uint256 remainAmount, address indexed nftAsset, uint256 nftTokenId, address indexed borrower, uint256 loanId)
func (_BendDAO *BendDAOFilterer) ParseLiquidate(log types.Log) (*BendDAOLiquidate, error) {
	event := new(BendDAOLiquidate)
	if err := _BendDAO.contract.UnpackLog(event, "Liquidate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BendDAOPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the BendDAO contract.
type BendDAOPausedIterator struct {
	Event *BendDAOPaused // Event containing the contract specifics and raw log

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
func (it *BendDAOPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BendDAOPaused)
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
		it.Event = new(BendDAOPaused)
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
func (it *BendDAOPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BendDAOPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BendDAOPaused represents a Paused event raised by the BendDAO contract.
type BendDAOPaused struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x9e87fac88ff661f02d44f95383c817fece4bce600a3dab7a54406878b965e752.
//
// Solidity: event Paused()
func (_BendDAO *BendDAOFilterer) FilterPaused(opts *bind.FilterOpts) (*BendDAOPausedIterator, error) {

	logs, sub, err := _BendDAO.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &BendDAOPausedIterator{contract: _BendDAO.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x9e87fac88ff661f02d44f95383c817fece4bce600a3dab7a54406878b965e752.
//
// Solidity: event Paused()
func (_BendDAO *BendDAOFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *BendDAOPaused) (event.Subscription, error) {

	logs, sub, err := _BendDAO.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BendDAOPaused)
				if err := _BendDAO.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x9e87fac88ff661f02d44f95383c817fece4bce600a3dab7a54406878b965e752.
//
// Solidity: event Paused()
func (_BendDAO *BendDAOFilterer) ParsePaused(log types.Log) (*BendDAOPaused, error) {
	event := new(BendDAOPaused)
	if err := _BendDAO.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BendDAOPausedTimeUpdatedIterator is returned from FilterPausedTimeUpdated and is used to iterate over the raw logs and unpacked data for PausedTimeUpdated events raised by the BendDAO contract.
type BendDAOPausedTimeUpdatedIterator struct {
	Event *BendDAOPausedTimeUpdated // Event containing the contract specifics and raw log

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
func (it *BendDAOPausedTimeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BendDAOPausedTimeUpdated)
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
		it.Event = new(BendDAOPausedTimeUpdated)
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
func (it *BendDAOPausedTimeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BendDAOPausedTimeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BendDAOPausedTimeUpdated represents a PausedTimeUpdated event raised by the BendDAO contract.
type BendDAOPausedTimeUpdated struct {
	StartTime    *big.Int
	DurationTime *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterPausedTimeUpdated is a free log retrieval operation binding the contract event 0xd897a722b1c0a957941f99a13c0ea24d7d4ffafe0953658f68f49e13ccba5c5a.
//
// Solidity: event PausedTimeUpdated(uint256 startTime, uint256 durationTime)
func (_BendDAO *BendDAOFilterer) FilterPausedTimeUpdated(opts *bind.FilterOpts) (*BendDAOPausedTimeUpdatedIterator, error) {

	logs, sub, err := _BendDAO.contract.FilterLogs(opts, "PausedTimeUpdated")
	if err != nil {
		return nil, err
	}
	return &BendDAOPausedTimeUpdatedIterator{contract: _BendDAO.contract, event: "PausedTimeUpdated", logs: logs, sub: sub}, nil
}

// WatchPausedTimeUpdated is a free log subscription operation binding the contract event 0xd897a722b1c0a957941f99a13c0ea24d7d4ffafe0953658f68f49e13ccba5c5a.
//
// Solidity: event PausedTimeUpdated(uint256 startTime, uint256 durationTime)
func (_BendDAO *BendDAOFilterer) WatchPausedTimeUpdated(opts *bind.WatchOpts, sink chan<- *BendDAOPausedTimeUpdated) (event.Subscription, error) {

	logs, sub, err := _BendDAO.contract.WatchLogs(opts, "PausedTimeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BendDAOPausedTimeUpdated)
				if err := _BendDAO.contract.UnpackLog(event, "PausedTimeUpdated", log); err != nil {
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

// ParsePausedTimeUpdated is a log parse operation binding the contract event 0xd897a722b1c0a957941f99a13c0ea24d7d4ffafe0953658f68f49e13ccba5c5a.
//
// Solidity: event PausedTimeUpdated(uint256 startTime, uint256 durationTime)
func (_BendDAO *BendDAOFilterer) ParsePausedTimeUpdated(log types.Log) (*BendDAOPausedTimeUpdated, error) {
	event := new(BendDAOPausedTimeUpdated)
	if err := _BendDAO.contract.UnpackLog(event, "PausedTimeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BendDAORedeemIterator is returned from FilterRedeem and is used to iterate over the raw logs and unpacked data for Redeem events raised by the BendDAO contract.
type BendDAORedeemIterator struct {
	Event *BendDAORedeem // Event containing the contract specifics and raw log

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
func (it *BendDAORedeemIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BendDAORedeem)
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
		it.Event = new(BendDAORedeem)
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
func (it *BendDAORedeemIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BendDAORedeemIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BendDAORedeem represents a Redeem event raised by the BendDAO contract.
type BendDAORedeem struct {
	User         common.Address
	Reserve      common.Address
	BorrowAmount *big.Int
	FineAmount   *big.Int
	NftAsset     common.Address
	NftTokenId   *big.Int
	Borrower     common.Address
	LoanId       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRedeem is a free log retrieval operation binding the contract event 0x0fcfe1a3f2afab13e32fa3c091795159ed5dfe66dc078e21c7f521f42e163afc.
//
// Solidity: event Redeem(address user, address indexed reserve, uint256 borrowAmount, uint256 fineAmount, address indexed nftAsset, uint256 nftTokenId, address indexed borrower, uint256 loanId)
func (_BendDAO *BendDAOFilterer) FilterRedeem(opts *bind.FilterOpts, reserve []common.Address, nftAsset []common.Address, borrower []common.Address) (*BendDAORedeemIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	var nftAssetRule []interface{}
	for _, nftAssetItem := range nftAsset {
		nftAssetRule = append(nftAssetRule, nftAssetItem)
	}

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _BendDAO.contract.FilterLogs(opts, "Redeem", reserveRule, nftAssetRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return &BendDAORedeemIterator{contract: _BendDAO.contract, event: "Redeem", logs: logs, sub: sub}, nil
}

// WatchRedeem is a free log subscription operation binding the contract event 0x0fcfe1a3f2afab13e32fa3c091795159ed5dfe66dc078e21c7f521f42e163afc.
//
// Solidity: event Redeem(address user, address indexed reserve, uint256 borrowAmount, uint256 fineAmount, address indexed nftAsset, uint256 nftTokenId, address indexed borrower, uint256 loanId)
func (_BendDAO *BendDAOFilterer) WatchRedeem(opts *bind.WatchOpts, sink chan<- *BendDAORedeem, reserve []common.Address, nftAsset []common.Address, borrower []common.Address) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	var nftAssetRule []interface{}
	for _, nftAssetItem := range nftAsset {
		nftAssetRule = append(nftAssetRule, nftAssetItem)
	}

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _BendDAO.contract.WatchLogs(opts, "Redeem", reserveRule, nftAssetRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BendDAORedeem)
				if err := _BendDAO.contract.UnpackLog(event, "Redeem", log); err != nil {
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

// ParseRedeem is a log parse operation binding the contract event 0x0fcfe1a3f2afab13e32fa3c091795159ed5dfe66dc078e21c7f521f42e163afc.
//
// Solidity: event Redeem(address user, address indexed reserve, uint256 borrowAmount, uint256 fineAmount, address indexed nftAsset, uint256 nftTokenId, address indexed borrower, uint256 loanId)
func (_BendDAO *BendDAOFilterer) ParseRedeem(log types.Log) (*BendDAORedeem, error) {
	event := new(BendDAORedeem)
	if err := _BendDAO.contract.UnpackLog(event, "Redeem", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BendDAORepayIterator is returned from FilterRepay and is used to iterate over the raw logs and unpacked data for Repay events raised by the BendDAO contract.
type BendDAORepayIterator struct {
	Event *BendDAORepay // Event containing the contract specifics and raw log

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
func (it *BendDAORepayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BendDAORepay)
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
		it.Event = new(BendDAORepay)
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
func (it *BendDAORepayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BendDAORepayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BendDAORepay represents a Repay event raised by the BendDAO contract.
type BendDAORepay struct {
	User       common.Address
	Reserve    common.Address
	Amount     *big.Int
	NftAsset   common.Address
	NftTokenId *big.Int
	Borrower   common.Address
	LoanId     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRepay is a free log retrieval operation binding the contract event 0x50e03867c1178391f204f7bf0eb2f52d5167dc65a99a9650a95abe55d17be17e.
//
// Solidity: event Repay(address user, address indexed reserve, uint256 amount, address indexed nftAsset, uint256 nftTokenId, address indexed borrower, uint256 loanId)
func (_BendDAO *BendDAOFilterer) FilterRepay(opts *bind.FilterOpts, reserve []common.Address, nftAsset []common.Address, borrower []common.Address) (*BendDAORepayIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	var nftAssetRule []interface{}
	for _, nftAssetItem := range nftAsset {
		nftAssetRule = append(nftAssetRule, nftAssetItem)
	}

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _BendDAO.contract.FilterLogs(opts, "Repay", reserveRule, nftAssetRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return &BendDAORepayIterator{contract: _BendDAO.contract, event: "Repay", logs: logs, sub: sub}, nil
}

// WatchRepay is a free log subscription operation binding the contract event 0x50e03867c1178391f204f7bf0eb2f52d5167dc65a99a9650a95abe55d17be17e.
//
// Solidity: event Repay(address user, address indexed reserve, uint256 amount, address indexed nftAsset, uint256 nftTokenId, address indexed borrower, uint256 loanId)
func (_BendDAO *BendDAOFilterer) WatchRepay(opts *bind.WatchOpts, sink chan<- *BendDAORepay, reserve []common.Address, nftAsset []common.Address, borrower []common.Address) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	var nftAssetRule []interface{}
	for _, nftAssetItem := range nftAsset {
		nftAssetRule = append(nftAssetRule, nftAssetItem)
	}

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _BendDAO.contract.WatchLogs(opts, "Repay", reserveRule, nftAssetRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BendDAORepay)
				if err := _BendDAO.contract.UnpackLog(event, "Repay", log); err != nil {
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

// ParseRepay is a log parse operation binding the contract event 0x50e03867c1178391f204f7bf0eb2f52d5167dc65a99a9650a95abe55d17be17e.
//
// Solidity: event Repay(address user, address indexed reserve, uint256 amount, address indexed nftAsset, uint256 nftTokenId, address indexed borrower, uint256 loanId)
func (_BendDAO *BendDAOFilterer) ParseRepay(log types.Log) (*BendDAORepay, error) {
	event := new(BendDAORepay)
	if err := _BendDAO.contract.UnpackLog(event, "Repay", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BendDAOReserveDataUpdatedIterator is returned from FilterReserveDataUpdated and is used to iterate over the raw logs and unpacked data for ReserveDataUpdated events raised by the BendDAO contract.
type BendDAOReserveDataUpdatedIterator struct {
	Event *BendDAOReserveDataUpdated // Event containing the contract specifics and raw log

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
func (it *BendDAOReserveDataUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BendDAOReserveDataUpdated)
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
		it.Event = new(BendDAOReserveDataUpdated)
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
func (it *BendDAOReserveDataUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BendDAOReserveDataUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BendDAOReserveDataUpdated represents a ReserveDataUpdated event raised by the BendDAO contract.
type BendDAOReserveDataUpdated struct {
	Reserve             common.Address
	LiquidityRate       *big.Int
	VariableBorrowRate  *big.Int
	LiquidityIndex      *big.Int
	VariableBorrowIndex *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterReserveDataUpdated is a free log retrieval operation binding the contract event 0x4063a2df84b66bb796eb32622851d833e57b2c4292900c18f963af8808b13e35.
//
// Solidity: event ReserveDataUpdated(address indexed reserve, uint256 liquidityRate, uint256 variableBorrowRate, uint256 liquidityIndex, uint256 variableBorrowIndex)
func (_BendDAO *BendDAOFilterer) FilterReserveDataUpdated(opts *bind.FilterOpts, reserve []common.Address) (*BendDAOReserveDataUpdatedIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	logs, sub, err := _BendDAO.contract.FilterLogs(opts, "ReserveDataUpdated", reserveRule)
	if err != nil {
		return nil, err
	}
	return &BendDAOReserveDataUpdatedIterator{contract: _BendDAO.contract, event: "ReserveDataUpdated", logs: logs, sub: sub}, nil
}

// WatchReserveDataUpdated is a free log subscription operation binding the contract event 0x4063a2df84b66bb796eb32622851d833e57b2c4292900c18f963af8808b13e35.
//
// Solidity: event ReserveDataUpdated(address indexed reserve, uint256 liquidityRate, uint256 variableBorrowRate, uint256 liquidityIndex, uint256 variableBorrowIndex)
func (_BendDAO *BendDAOFilterer) WatchReserveDataUpdated(opts *bind.WatchOpts, sink chan<- *BendDAOReserveDataUpdated, reserve []common.Address) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	logs, sub, err := _BendDAO.contract.WatchLogs(opts, "ReserveDataUpdated", reserveRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BendDAOReserveDataUpdated)
				if err := _BendDAO.contract.UnpackLog(event, "ReserveDataUpdated", log); err != nil {
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

// ParseReserveDataUpdated is a log parse operation binding the contract event 0x4063a2df84b66bb796eb32622851d833e57b2c4292900c18f963af8808b13e35.
//
// Solidity: event ReserveDataUpdated(address indexed reserve, uint256 liquidityRate, uint256 variableBorrowRate, uint256 liquidityIndex, uint256 variableBorrowIndex)
func (_BendDAO *BendDAOFilterer) ParseReserveDataUpdated(log types.Log) (*BendDAOReserveDataUpdated, error) {
	event := new(BendDAOReserveDataUpdated)
	if err := _BendDAO.contract.UnpackLog(event, "ReserveDataUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BendDAOUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the BendDAO contract.
type BendDAOUnpausedIterator struct {
	Event *BendDAOUnpaused // Event containing the contract specifics and raw log

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
func (it *BendDAOUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BendDAOUnpaused)
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
		it.Event = new(BendDAOUnpaused)
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
func (it *BendDAOUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BendDAOUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BendDAOUnpaused represents a Unpaused event raised by the BendDAO contract.
type BendDAOUnpaused struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0xa45f47fdea8a1efdd9029a5691c7f759c32b7c698632b563573e155625d16933.
//
// Solidity: event Unpaused()
func (_BendDAO *BendDAOFilterer) FilterUnpaused(opts *bind.FilterOpts) (*BendDAOUnpausedIterator, error) {

	logs, sub, err := _BendDAO.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &BendDAOUnpausedIterator{contract: _BendDAO.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0xa45f47fdea8a1efdd9029a5691c7f759c32b7c698632b563573e155625d16933.
//
// Solidity: event Unpaused()
func (_BendDAO *BendDAOFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *BendDAOUnpaused) (event.Subscription, error) {

	logs, sub, err := _BendDAO.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BendDAOUnpaused)
				if err := _BendDAO.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0xa45f47fdea8a1efdd9029a5691c7f759c32b7c698632b563573e155625d16933.
//
// Solidity: event Unpaused()
func (_BendDAO *BendDAOFilterer) ParseUnpaused(log types.Log) (*BendDAOUnpaused, error) {
	event := new(BendDAOUnpaused)
	if err := _BendDAO.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BendDAOWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the BendDAO contract.
type BendDAOWithdrawIterator struct {
	Event *BendDAOWithdraw // Event containing the contract specifics and raw log

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
func (it *BendDAOWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BendDAOWithdraw)
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
		it.Event = new(BendDAOWithdraw)
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
func (it *BendDAOWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BendDAOWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BendDAOWithdraw represents a Withdraw event raised by the BendDAO contract.
type BendDAOWithdraw struct {
	User    common.Address
	Reserve common.Address
	Amount  *big.Int
	To      common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x3ed4ee04a905a278b050a856bbe7ddaaf327a30514373e65aa6103beeae488c3.
//
// Solidity: event Withdraw(address indexed user, address indexed reserve, uint256 amount, address indexed to)
func (_BendDAO *BendDAOFilterer) FilterWithdraw(opts *bind.FilterOpts, user []common.Address, reserve []common.Address, to []common.Address) (*BendDAOWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BendDAO.contract.FilterLogs(opts, "Withdraw", userRule, reserveRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BendDAOWithdrawIterator{contract: _BendDAO.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x3ed4ee04a905a278b050a856bbe7ddaaf327a30514373e65aa6103beeae488c3.
//
// Solidity: event Withdraw(address indexed user, address indexed reserve, uint256 amount, address indexed to)
func (_BendDAO *BendDAOFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *BendDAOWithdraw, user []common.Address, reserve []common.Address, to []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BendDAO.contract.WatchLogs(opts, "Withdraw", userRule, reserveRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BendDAOWithdraw)
				if err := _BendDAO.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x3ed4ee04a905a278b050a856bbe7ddaaf327a30514373e65aa6103beeae488c3.
//
// Solidity: event Withdraw(address indexed user, address indexed reserve, uint256 amount, address indexed to)
func (_BendDAO *BendDAOFilterer) ParseWithdraw(log types.Log) (*BendDAOWithdraw, error) {
	event := new(BendDAOWithdraw)
	if err := _BendDAO.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
