// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package nswap

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

// LibAssetAsset is an auto generated low-level Go binding around an user-defined struct.
type LibAssetAsset struct {
	AssetType LibAssetAssetType
	Value     *big.Int
}

// LibAssetAssetType is an auto generated low-level Go binding around an user-defined struct.
type LibAssetAssetType struct {
	AssetClass   [4]byte
	TokenAddress common.Address
	TokenID      *big.Int
}

// LibOrderOrder is an auto generated low-level Go binding around an user-defined struct.
type LibOrderOrder struct {
	Side     uint8
	SaleKind uint8
	Maker    common.Address
	Taker    common.Address
	Nft      LibAssetAsset
	Price    LibAssetAsset
	Salt     *big.Int
	Start    *big.Int
	End      *big.Int
	Version  [4]byte
	Extra    []byte
}

// LibPayInfoPayInfo is an auto generated low-level Go binding around an user-defined struct.
type LibPayInfoPayInfo struct {
	Receiver common.Address
	Share    *big.Int
}

// NSwapExchangeMatchDetails is an auto generated low-level Go binding around an user-defined struct.
type NSwapExchangeMatchDetails struct {
	SellOrder LibOrderOrder
	BuyOrder  LibOrderOrder
	Signature []byte
}

// NSwapExchangeMetaData contains all meta data concerning the NSwapExchange contract.
var NSwapExchangeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AssetMatcher_Asset_Class_Mismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AssetMatcher_Asset_Not_Supported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AssetMatcher_Token_Addr_Mismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AssetMatcher_Token_ID_Mismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CurrencyManager_Already_Approved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CurrencyManager_Currency_Zero_Address\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CurrencyManager_Not_Approved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NSwapExchange_Cannot_0_Salt\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NSwapExchange_Fill_Amount_Cannot_0\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NSwapExchange_Maker_Taker_Mismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NSwapExchange_Only_Maker_Can_Cancel\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NSwapExchange_Order_Side_Mismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NSwapExchange_Price_Mismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NSwapExchange_Prohibit_Self_Match\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NSwapExchange_Sale_Kind_Mismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NSwapExchange_Tx_Sender_Invalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NSwapExchange_Unsupported_Price_Asset\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NSwapExchange_Unsupported_Token\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NSwapExchange_Wrong_ETH_Value\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OrderData_Unsupported_Data_Version\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OrderValidator_Cannot_0_Salt\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OrderValidator_Contract_Sig_Verify_fail\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OrderValidator_Has_Canceled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OrderValidator_Has_Expired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OrderValidator_Has_Not_Started\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OrderValidator_Must_Have_Maker\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OrderValidator_Sig_Verify_fail\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OrderValidator_Unsupported_NFT_Asset\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OrderValidator_Unsupported_Price_Asset\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProtocolFeeManager_Fee_Exceed_Max_Protocol_Share\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProtocolFeeManager_Fee_Receiver_Zero_Address\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RoyaltiesRegistry_ERC2981_Royalties_Exceed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RoyaltiesRegistry_Get_Royalties_Error\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RoyaltiesRegistry_Invalid_Address\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RoyaltiesRegistry_Invalid_Royalty_Value\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RoyaltiesRegistry_Not_Owner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RoyaltiesRegistry_Royalties_Limit_Exceed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RoyaltiesRegistry_Sum_Royalties_Exceed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RoyaltiesRegistry_Wrong_RoyaltiesType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Signature_Invalid_Sig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Signature_Invalid_Sig_Length\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Signature_Invalid_Sig_S_Value\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Signature_Invalid_Sig_V_Value\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferExecutor_Currency_Not_Approved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferExecutor_ERC20_Failed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferExecutor_ERC721_Value_Err\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferExecutor_Unsupported_Token\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferManager_Royalty_Too_High\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferManager_SumOfShare_Mismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Transfer_Eth_Fail\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes4\",\"name\":\"assetClass\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structLibAsset.AssetType\",\"name\":\"priceAssetType\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes4\",\"name\":\"assetClass\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structLibAsset.AssetType\",\"name\":\"nftAssetType\",\"type\":\"tuple\"}],\"name\":\"EventCancel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"}],\"name\":\"EventCurrencyApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"}],\"name\":\"EventCurrencyRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"buyHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"sellHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"buyMaker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sellMaker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"priceFilled\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftFilled\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes4\",\"name\":\"assetClass\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structLibAsset.AssetType\",\"name\":\"priceAssetType\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes4\",\"name\":\"assetClass\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structLibAsset.AssetType\",\"name\":\"nftAssetType\",\"type\":\"tuple\"}],\"name\":\"EventMatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"customAccount\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"EventUpdatedCustomProtocolFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"defaultFee\",\"type\":\"uint256\"}],\"name\":\"EventUpdatedDefaultProtocolFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"defaultFeeReceiver\",\"type\":\"address\"}],\"name\":\"EventUpdatedDefaultProtocolFeeReceiver\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"addresspayable\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"share\",\"type\":\"uint96\"}],\"indexed\":false,\"internalType\":\"structLibPayInfo.PayInfo[]\",\"name\":\"royalties\",\"type\":\"tuple[]\"}],\"name\":\"EventUpdatedRoyaltiesByToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newRoyaltiesLimit\",\"type\":\"uint256\"}],\"name\":\"EventUpdatedRoyaltiesLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EventWithdrawETH\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"}],\"name\":\"approveCurrency\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumLibOrder.Side\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumLibOrder.SaleKind\",\"name\":\"saleKind\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes4\",\"name\":\"assetClass\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"internalType\":\"structLibAsset.AssetType\",\"name\":\"assetType\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structLibAsset.Asset\",\"name\":\"nft\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes4\",\"name\":\"assetClass\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"internalType\":\"structLibAsset.AssetType\",\"name\":\"assetType\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structLibAsset.Asset\",\"name\":\"price\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"version\",\"type\":\"bytes4\"},{\"internalType\":\"bytes\",\"name\":\"extra\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order[]\",\"name\":\"orders\",\"type\":\"tuple[]\"}],\"name\":\"batchCancel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"enumLibOrder.Side\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumLibOrder.SaleKind\",\"name\":\"saleKind\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes4\",\"name\":\"assetClass\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"internalType\":\"structLibAsset.AssetType\",\"name\":\"assetType\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structLibAsset.Asset\",\"name\":\"nft\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes4\",\"name\":\"assetClass\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"internalType\":\"structLibAsset.AssetType\",\"name\":\"assetType\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structLibAsset.Asset\",\"name\":\"price\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"version\",\"type\":\"bytes4\"},{\"internalType\":\"bytes\",\"name\":\"extra\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order\",\"name\":\"sellOrder\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumLibOrder.Side\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumLibOrder.SaleKind\",\"name\":\"saleKind\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes4\",\"name\":\"assetClass\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"internalType\":\"structLibAsset.AssetType\",\"name\":\"assetType\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structLibAsset.Asset\",\"name\":\"nft\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes4\",\"name\":\"assetClass\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"internalType\":\"structLibAsset.AssetType\",\"name\":\"assetType\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structLibAsset.Asset\",\"name\":\"price\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"version\",\"type\":\"bytes4\"},{\"internalType\":\"bytes\",\"name\":\"extra\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order\",\"name\":\"buyOrder\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structNSwapExchange.MatchDetails[]\",\"name\":\"matchDetails\",\"type\":\"tuple[]\"}],\"name\":\"batchMatch\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes4\",\"name\":\"assetClass\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"internalType\":\"structLibAsset.AssetType[]\",\"name\":\"assets\",\"type\":\"tuple[]\"}],\"name\":\"batchTransferERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumLibOrder.Side\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumLibOrder.SaleKind\",\"name\":\"saleKind\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes4\",\"name\":\"assetClass\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"internalType\":\"structLibAsset.AssetType\",\"name\":\"assetType\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structLibAsset.Asset\",\"name\":\"nft\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes4\",\"name\":\"assetClass\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"internalType\":\"structLibAsset.AssetType\",\"name\":\"assetType\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structLibAsset.Asset\",\"name\":\"price\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"version\",\"type\":\"bytes4\"},{\"internalType\":\"bytes\",\"name\":\"extra\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"cancel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"clearRoyaltiesType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultProtocolFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultProtocolFeeReceiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"fillsStat\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"getProtocolFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getRoyalties\",\"outputs\":[{\"components\":[{\"internalType\":\"addresspayable\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"share\",\"type\":\"uint96\"}],\"internalType\":\"structLibPayInfo.PayInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getRoyaltiesType\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getRoyaltiesView\",\"outputs\":[{\"components\":[{\"internalType\":\"addresspayable\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"share\",\"type\":\"uint96\"}],\"internalType\":\"structLibPayInfo.PayInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"currencies\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"defaultProtocolFee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"defaultFeeRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"newRoyaltiesLimit\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"EIP712Name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"EIP712Version\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"}],\"name\":\"isCurrencyApproved\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumLibOrder.Side\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumLibOrder.SaleKind\",\"name\":\"saleKind\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes4\",\"name\":\"assetClass\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"internalType\":\"structLibAsset.AssetType\",\"name\":\"assetType\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structLibAsset.Asset\",\"name\":\"nft\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes4\",\"name\":\"assetClass\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"internalType\":\"structLibAsset.AssetType\",\"name\":\"assetType\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structLibAsset.Asset\",\"name\":\"price\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"version\",\"type\":\"bytes4\"},{\"internalType\":\"bytes\",\"name\":\"extra\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order\",\"name\":\"sell\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumLibOrder.Side\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumLibOrder.SaleKind\",\"name\":\"saleKind\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes4\",\"name\":\"assetClass\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"internalType\":\"structLibAsset.AssetType\",\"name\":\"assetType\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structLibAsset.Asset\",\"name\":\"nft\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes4\",\"name\":\"assetClass\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"internalType\":\"structLibAsset.AssetType\",\"name\":\"assetType\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structLibAsset.Asset\",\"name\":\"price\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"version\",\"type\":\"bytes4\"},{\"internalType\":\"bytes\",\"name\":\"extra\",\"type\":\"bytes\"}],\"internalType\":\"structLibOrder.Order\",\"name\":\"buy\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"matchOrders\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"protocolFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"}],\"name\":\"revokeCurrency\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"royaltiesByToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isByToken\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"royaltiesLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"royaltiesType\",\"outputs\":[{\"internalType\":\"enumRoyaltiesManager.RoyaltyType\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newDefaultProtocolFee\",\"type\":\"uint256\"}],\"name\":\"setDefaultProtocolFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newDefaultProtocolFeeReceiver\",\"type\":\"address\"}],\"name\":\"setDefaultProtocolFeeReceiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"newProtocolFee\",\"type\":\"uint256\"}],\"name\":\"setProtocolFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"addresspayable\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"share\",\"type\":\"uint96\"}],\"internalType\":\"structLibPayInfo.PayInfo[]\",\"name\":\"royalties\",\"type\":\"tuple[]\"}],\"name\":\"setRoyaltiesByToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"newRoyaltiesType\",\"type\":\"uint256\"}],\"name\":\"setRoyaltiesType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newRoyaltiesLimit\",\"type\":\"uint256\"}],\"name\":\"updateRoyaltiesLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"viewApprovedCurrency\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"viewCountApprovedCurrencies\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// NSwapExchangeABI is the input ABI used to generate the binding from.
// Deprecated: Use NSwapExchangeMetaData.ABI instead.
var NSwapExchangeABI = NSwapExchangeMetaData.ABI

// NSwapExchange is an auto generated Go binding around an Ethereum contract.
type NSwapExchange struct {
	NSwapExchangeCaller     // Read-only binding to the contract
	NSwapExchangeTransactor // Write-only binding to the contract
	NSwapExchangeFilterer   // Log filterer for contract events
}

// NSwapExchangeCaller is an auto generated read-only Go binding around an Ethereum contract.
type NSwapExchangeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NSwapExchangeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NSwapExchangeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NSwapExchangeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NSwapExchangeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NSwapExchangeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NSwapExchangeSession struct {
	Contract     *NSwapExchange    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NSwapExchangeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NSwapExchangeCallerSession struct {
	Contract *NSwapExchangeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// NSwapExchangeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NSwapExchangeTransactorSession struct {
	Contract     *NSwapExchangeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// NSwapExchangeRaw is an auto generated low-level Go binding around an Ethereum contract.
type NSwapExchangeRaw struct {
	Contract *NSwapExchange // Generic contract binding to access the raw methods on
}

// NSwapExchangeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NSwapExchangeCallerRaw struct {
	Contract *NSwapExchangeCaller // Generic read-only contract binding to access the raw methods on
}

// NSwapExchangeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NSwapExchangeTransactorRaw struct {
	Contract *NSwapExchangeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNSwapExchange creates a new instance of NSwapExchange, bound to a specific deployed contract.
func NewNSwapExchange(address common.Address, backend bind.ContractBackend) (*NSwapExchange, error) {
	contract, err := bindNSwapExchange(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NSwapExchange{NSwapExchangeCaller: NSwapExchangeCaller{contract: contract}, NSwapExchangeTransactor: NSwapExchangeTransactor{contract: contract}, NSwapExchangeFilterer: NSwapExchangeFilterer{contract: contract}}, nil
}

// NewNSwapExchangeCaller creates a new read-only instance of NSwapExchange, bound to a specific deployed contract.
func NewNSwapExchangeCaller(address common.Address, caller bind.ContractCaller) (*NSwapExchangeCaller, error) {
	contract, err := bindNSwapExchange(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NSwapExchangeCaller{contract: contract}, nil
}

// NewNSwapExchangeTransactor creates a new write-only instance of NSwapExchange, bound to a specific deployed contract.
func NewNSwapExchangeTransactor(address common.Address, transactor bind.ContractTransactor) (*NSwapExchangeTransactor, error) {
	contract, err := bindNSwapExchange(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NSwapExchangeTransactor{contract: contract}, nil
}

// NewNSwapExchangeFilterer creates a new log filterer instance of NSwapExchange, bound to a specific deployed contract.
func NewNSwapExchangeFilterer(address common.Address, filterer bind.ContractFilterer) (*NSwapExchangeFilterer, error) {
	contract, err := bindNSwapExchange(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NSwapExchangeFilterer{contract: contract}, nil
}

// bindNSwapExchange binds a generic wrapper to an already deployed contract.
func bindNSwapExchange(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NSwapExchangeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NSwapExchange *NSwapExchangeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NSwapExchange.Contract.NSwapExchangeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NSwapExchange *NSwapExchangeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NSwapExchange.Contract.NSwapExchangeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NSwapExchange *NSwapExchangeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NSwapExchange.Contract.NSwapExchangeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NSwapExchange *NSwapExchangeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NSwapExchange.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NSwapExchange *NSwapExchangeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NSwapExchange.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NSwapExchange *NSwapExchangeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NSwapExchange.Contract.contract.Transact(opts, method, params...)
}

// DefaultProtocolFee is a free data retrieval call binding the contract method 0x421bcbdf.
//
// Solidity: function defaultProtocolFee() view returns(uint256)
func (_NSwapExchange *NSwapExchangeCaller) DefaultProtocolFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NSwapExchange.contract.Call(opts, &out, "defaultProtocolFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultProtocolFee is a free data retrieval call binding the contract method 0x421bcbdf.
//
// Solidity: function defaultProtocolFee() view returns(uint256)
func (_NSwapExchange *NSwapExchangeSession) DefaultProtocolFee() (*big.Int, error) {
	return _NSwapExchange.Contract.DefaultProtocolFee(&_NSwapExchange.CallOpts)
}

// DefaultProtocolFee is a free data retrieval call binding the contract method 0x421bcbdf.
//
// Solidity: function defaultProtocolFee() view returns(uint256)
func (_NSwapExchange *NSwapExchangeCallerSession) DefaultProtocolFee() (*big.Int, error) {
	return _NSwapExchange.Contract.DefaultProtocolFee(&_NSwapExchange.CallOpts)
}

// DefaultProtocolFeeReceiver is a free data retrieval call binding the contract method 0x1fbff894.
//
// Solidity: function defaultProtocolFeeReceiver() view returns(address)
func (_NSwapExchange *NSwapExchangeCaller) DefaultProtocolFeeReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NSwapExchange.contract.Call(opts, &out, "defaultProtocolFeeReceiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DefaultProtocolFeeReceiver is a free data retrieval call binding the contract method 0x1fbff894.
//
// Solidity: function defaultProtocolFeeReceiver() view returns(address)
func (_NSwapExchange *NSwapExchangeSession) DefaultProtocolFeeReceiver() (common.Address, error) {
	return _NSwapExchange.Contract.DefaultProtocolFeeReceiver(&_NSwapExchange.CallOpts)
}

// DefaultProtocolFeeReceiver is a free data retrieval call binding the contract method 0x1fbff894.
//
// Solidity: function defaultProtocolFeeReceiver() view returns(address)
func (_NSwapExchange *NSwapExchangeCallerSession) DefaultProtocolFeeReceiver() (common.Address, error) {
	return _NSwapExchange.Contract.DefaultProtocolFeeReceiver(&_NSwapExchange.CallOpts)
}

// FillsStat is a free data retrieval call binding the contract method 0x17db7e98.
//
// Solidity: function fillsStat(bytes32 ) view returns(uint256)
func (_NSwapExchange *NSwapExchangeCaller) FillsStat(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _NSwapExchange.contract.Call(opts, &out, "fillsStat", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FillsStat is a free data retrieval call binding the contract method 0x17db7e98.
//
// Solidity: function fillsStat(bytes32 ) view returns(uint256)
func (_NSwapExchange *NSwapExchangeSession) FillsStat(arg0 [32]byte) (*big.Int, error) {
	return _NSwapExchange.Contract.FillsStat(&_NSwapExchange.CallOpts, arg0)
}

// FillsStat is a free data retrieval call binding the contract method 0x17db7e98.
//
// Solidity: function fillsStat(bytes32 ) view returns(uint256)
func (_NSwapExchange *NSwapExchangeCallerSession) FillsStat(arg0 [32]byte) (*big.Int, error) {
	return _NSwapExchange.Contract.FillsStat(&_NSwapExchange.CallOpts, arg0)
}

// GetProtocolFee is a free data retrieval call binding the contract method 0x0a992e0c.
//
// Solidity: function getProtocolFee(address wallet) view returns(uint256)
func (_NSwapExchange *NSwapExchangeCaller) GetProtocolFee(opts *bind.CallOpts, wallet common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NSwapExchange.contract.Call(opts, &out, "getProtocolFee", wallet)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProtocolFee is a free data retrieval call binding the contract method 0x0a992e0c.
//
// Solidity: function getProtocolFee(address wallet) view returns(uint256)
func (_NSwapExchange *NSwapExchangeSession) GetProtocolFee(wallet common.Address) (*big.Int, error) {
	return _NSwapExchange.Contract.GetProtocolFee(&_NSwapExchange.CallOpts, wallet)
}

// GetProtocolFee is a free data retrieval call binding the contract method 0x0a992e0c.
//
// Solidity: function getProtocolFee(address wallet) view returns(uint256)
func (_NSwapExchange *NSwapExchangeCallerSession) GetProtocolFee(wallet common.Address) (*big.Int, error) {
	return _NSwapExchange.Contract.GetProtocolFee(&_NSwapExchange.CallOpts, wallet)
}

// GetRoyaltiesType is a free data retrieval call binding the contract method 0x82b19f12.
//
// Solidity: function getRoyaltiesType(address token) view returns(uint256)
func (_NSwapExchange *NSwapExchangeCaller) GetRoyaltiesType(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NSwapExchange.contract.Call(opts, &out, "getRoyaltiesType", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoyaltiesType is a free data retrieval call binding the contract method 0x82b19f12.
//
// Solidity: function getRoyaltiesType(address token) view returns(uint256)
func (_NSwapExchange *NSwapExchangeSession) GetRoyaltiesType(token common.Address) (*big.Int, error) {
	return _NSwapExchange.Contract.GetRoyaltiesType(&_NSwapExchange.CallOpts, token)
}

// GetRoyaltiesType is a free data retrieval call binding the contract method 0x82b19f12.
//
// Solidity: function getRoyaltiesType(address token) view returns(uint256)
func (_NSwapExchange *NSwapExchangeCallerSession) GetRoyaltiesType(token common.Address) (*big.Int, error) {
	return _NSwapExchange.Contract.GetRoyaltiesType(&_NSwapExchange.CallOpts, token)
}

// GetRoyaltiesView is a free data retrieval call binding the contract method 0x39f1930c.
//
// Solidity: function getRoyaltiesView(address token, uint256 tokenId) view returns((address,uint96)[])
func (_NSwapExchange *NSwapExchangeCaller) GetRoyaltiesView(opts *bind.CallOpts, token common.Address, tokenId *big.Int) ([]LibPayInfoPayInfo, error) {
	var out []interface{}
	err := _NSwapExchange.contract.Call(opts, &out, "getRoyaltiesView", token, tokenId)

	if err != nil {
		return *new([]LibPayInfoPayInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]LibPayInfoPayInfo)).(*[]LibPayInfoPayInfo)

	return out0, err

}

// GetRoyaltiesView is a free data retrieval call binding the contract method 0x39f1930c.
//
// Solidity: function getRoyaltiesView(address token, uint256 tokenId) view returns((address,uint96)[])
func (_NSwapExchange *NSwapExchangeSession) GetRoyaltiesView(token common.Address, tokenId *big.Int) ([]LibPayInfoPayInfo, error) {
	return _NSwapExchange.Contract.GetRoyaltiesView(&_NSwapExchange.CallOpts, token, tokenId)
}

// GetRoyaltiesView is a free data retrieval call binding the contract method 0x39f1930c.
//
// Solidity: function getRoyaltiesView(address token, uint256 tokenId) view returns((address,uint96)[])
func (_NSwapExchange *NSwapExchangeCallerSession) GetRoyaltiesView(token common.Address, tokenId *big.Int) ([]LibPayInfoPayInfo, error) {
	return _NSwapExchange.Contract.GetRoyaltiesView(&_NSwapExchange.CallOpts, token, tokenId)
}

// IsCurrencyApproved is a free data retrieval call binding the contract method 0x0d7b1acc.
//
// Solidity: function isCurrencyApproved(address currency) view returns(bool)
func (_NSwapExchange *NSwapExchangeCaller) IsCurrencyApproved(opts *bind.CallOpts, currency common.Address) (bool, error) {
	var out []interface{}
	err := _NSwapExchange.contract.Call(opts, &out, "isCurrencyApproved", currency)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCurrencyApproved is a free data retrieval call binding the contract method 0x0d7b1acc.
//
// Solidity: function isCurrencyApproved(address currency) view returns(bool)
func (_NSwapExchange *NSwapExchangeSession) IsCurrencyApproved(currency common.Address) (bool, error) {
	return _NSwapExchange.Contract.IsCurrencyApproved(&_NSwapExchange.CallOpts, currency)
}

// IsCurrencyApproved is a free data retrieval call binding the contract method 0x0d7b1acc.
//
// Solidity: function isCurrencyApproved(address currency) view returns(bool)
func (_NSwapExchange *NSwapExchangeCallerSession) IsCurrencyApproved(currency common.Address) (bool, error) {
	return _NSwapExchange.Contract.IsCurrencyApproved(&_NSwapExchange.CallOpts, currency)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NSwapExchange *NSwapExchangeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NSwapExchange.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NSwapExchange *NSwapExchangeSession) Owner() (common.Address, error) {
	return _NSwapExchange.Contract.Owner(&_NSwapExchange.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NSwapExchange *NSwapExchangeCallerSession) Owner() (common.Address, error) {
	return _NSwapExchange.Contract.Owner(&_NSwapExchange.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_NSwapExchange *NSwapExchangeCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _NSwapExchange.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_NSwapExchange *NSwapExchangeSession) Paused() (bool, error) {
	return _NSwapExchange.Contract.Paused(&_NSwapExchange.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_NSwapExchange *NSwapExchangeCallerSession) Paused() (bool, error) {
	return _NSwapExchange.Contract.Paused(&_NSwapExchange.CallOpts)
}

// ProtocolFee is a free data retrieval call binding the contract method 0xb62b31e4.
//
// Solidity: function protocolFee(address ) view returns(uint256)
func (_NSwapExchange *NSwapExchangeCaller) ProtocolFee(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NSwapExchange.contract.Call(opts, &out, "protocolFee", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProtocolFee is a free data retrieval call binding the contract method 0xb62b31e4.
//
// Solidity: function protocolFee(address ) view returns(uint256)
func (_NSwapExchange *NSwapExchangeSession) ProtocolFee(arg0 common.Address) (*big.Int, error) {
	return _NSwapExchange.Contract.ProtocolFee(&_NSwapExchange.CallOpts, arg0)
}

// ProtocolFee is a free data retrieval call binding the contract method 0xb62b31e4.
//
// Solidity: function protocolFee(address ) view returns(uint256)
func (_NSwapExchange *NSwapExchangeCallerSession) ProtocolFee(arg0 common.Address) (*big.Int, error) {
	return _NSwapExchange.Contract.ProtocolFee(&_NSwapExchange.CallOpts, arg0)
}

// RoyaltiesByToken is a free data retrieval call binding the contract method 0x05df952f.
//
// Solidity: function royaltiesByToken(address ) view returns(bool isByToken)
func (_NSwapExchange *NSwapExchangeCaller) RoyaltiesByToken(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _NSwapExchange.contract.Call(opts, &out, "royaltiesByToken", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RoyaltiesByToken is a free data retrieval call binding the contract method 0x05df952f.
//
// Solidity: function royaltiesByToken(address ) view returns(bool isByToken)
func (_NSwapExchange *NSwapExchangeSession) RoyaltiesByToken(arg0 common.Address) (bool, error) {
	return _NSwapExchange.Contract.RoyaltiesByToken(&_NSwapExchange.CallOpts, arg0)
}

// RoyaltiesByToken is a free data retrieval call binding the contract method 0x05df952f.
//
// Solidity: function royaltiesByToken(address ) view returns(bool isByToken)
func (_NSwapExchange *NSwapExchangeCallerSession) RoyaltiesByToken(arg0 common.Address) (bool, error) {
	return _NSwapExchange.Contract.RoyaltiesByToken(&_NSwapExchange.CallOpts, arg0)
}

// RoyaltiesLimit is a free data retrieval call binding the contract method 0x9499ef01.
//
// Solidity: function royaltiesLimit() view returns(uint256)
func (_NSwapExchange *NSwapExchangeCaller) RoyaltiesLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NSwapExchange.contract.Call(opts, &out, "royaltiesLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RoyaltiesLimit is a free data retrieval call binding the contract method 0x9499ef01.
//
// Solidity: function royaltiesLimit() view returns(uint256)
func (_NSwapExchange *NSwapExchangeSession) RoyaltiesLimit() (*big.Int, error) {
	return _NSwapExchange.Contract.RoyaltiesLimit(&_NSwapExchange.CallOpts)
}

// RoyaltiesLimit is a free data retrieval call binding the contract method 0x9499ef01.
//
// Solidity: function royaltiesLimit() view returns(uint256)
func (_NSwapExchange *NSwapExchangeCallerSession) RoyaltiesLimit() (*big.Int, error) {
	return _NSwapExchange.Contract.RoyaltiesLimit(&_NSwapExchange.CallOpts)
}

// RoyaltiesType is a free data retrieval call binding the contract method 0xb6d79755.
//
// Solidity: function royaltiesType(address ) view returns(uint8)
func (_NSwapExchange *NSwapExchangeCaller) RoyaltiesType(opts *bind.CallOpts, arg0 common.Address) (uint8, error) {
	var out []interface{}
	err := _NSwapExchange.contract.Call(opts, &out, "royaltiesType", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// RoyaltiesType is a free data retrieval call binding the contract method 0xb6d79755.
//
// Solidity: function royaltiesType(address ) view returns(uint8)
func (_NSwapExchange *NSwapExchangeSession) RoyaltiesType(arg0 common.Address) (uint8, error) {
	return _NSwapExchange.Contract.RoyaltiesType(&_NSwapExchange.CallOpts, arg0)
}

// RoyaltiesType is a free data retrieval call binding the contract method 0xb6d79755.
//
// Solidity: function royaltiesType(address ) view returns(uint8)
func (_NSwapExchange *NSwapExchangeCallerSession) RoyaltiesType(arg0 common.Address) (uint8, error) {
	return _NSwapExchange.Contract.RoyaltiesType(&_NSwapExchange.CallOpts, arg0)
}

// ViewApprovedCurrency is a free data retrieval call binding the contract method 0xa3aaf800.
//
// Solidity: function viewApprovedCurrency(uint256 index) view returns(address)
func (_NSwapExchange *NSwapExchangeCaller) ViewApprovedCurrency(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _NSwapExchange.contract.Call(opts, &out, "viewApprovedCurrency", index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ViewApprovedCurrency is a free data retrieval call binding the contract method 0xa3aaf800.
//
// Solidity: function viewApprovedCurrency(uint256 index) view returns(address)
func (_NSwapExchange *NSwapExchangeSession) ViewApprovedCurrency(index *big.Int) (common.Address, error) {
	return _NSwapExchange.Contract.ViewApprovedCurrency(&_NSwapExchange.CallOpts, index)
}

// ViewApprovedCurrency is a free data retrieval call binding the contract method 0xa3aaf800.
//
// Solidity: function viewApprovedCurrency(uint256 index) view returns(address)
func (_NSwapExchange *NSwapExchangeCallerSession) ViewApprovedCurrency(index *big.Int) (common.Address, error) {
	return _NSwapExchange.Contract.ViewApprovedCurrency(&_NSwapExchange.CallOpts, index)
}

// ViewCountApprovedCurrencies is a free data retrieval call binding the contract method 0x86f97e71.
//
// Solidity: function viewCountApprovedCurrencies() view returns(uint256)
func (_NSwapExchange *NSwapExchangeCaller) ViewCountApprovedCurrencies(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NSwapExchange.contract.Call(opts, &out, "viewCountApprovedCurrencies")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ViewCountApprovedCurrencies is a free data retrieval call binding the contract method 0x86f97e71.
//
// Solidity: function viewCountApprovedCurrencies() view returns(uint256)
func (_NSwapExchange *NSwapExchangeSession) ViewCountApprovedCurrencies() (*big.Int, error) {
	return _NSwapExchange.Contract.ViewCountApprovedCurrencies(&_NSwapExchange.CallOpts)
}

// ViewCountApprovedCurrencies is a free data retrieval call binding the contract method 0x86f97e71.
//
// Solidity: function viewCountApprovedCurrencies() view returns(uint256)
func (_NSwapExchange *NSwapExchangeCallerSession) ViewCountApprovedCurrencies() (*big.Int, error) {
	return _NSwapExchange.Contract.ViewCountApprovedCurrencies(&_NSwapExchange.CallOpts)
}

// ApproveCurrency is a paid mutator transaction binding the contract method 0xe20b5f6f.
//
// Solidity: function approveCurrency(address currency) returns()
func (_NSwapExchange *NSwapExchangeTransactor) ApproveCurrency(opts *bind.TransactOpts, currency common.Address) (*types.Transaction, error) {
	return _NSwapExchange.contract.Transact(opts, "approveCurrency", currency)
}

// ApproveCurrency is a paid mutator transaction binding the contract method 0xe20b5f6f.
//
// Solidity: function approveCurrency(address currency) returns()
func (_NSwapExchange *NSwapExchangeSession) ApproveCurrency(currency common.Address) (*types.Transaction, error) {
	return _NSwapExchange.Contract.ApproveCurrency(&_NSwapExchange.TransactOpts, currency)
}

// ApproveCurrency is a paid mutator transaction binding the contract method 0xe20b5f6f.
//
// Solidity: function approveCurrency(address currency) returns()
func (_NSwapExchange *NSwapExchangeTransactorSession) ApproveCurrency(currency common.Address) (*types.Transaction, error) {
	return _NSwapExchange.Contract.ApproveCurrency(&_NSwapExchange.TransactOpts, currency)
}

// BatchCancel is a paid mutator transaction binding the contract method 0x2d46942b.
//
// Solidity: function batchCancel((uint8,uint8,address,address,((bytes4,address,uint256),uint256),((bytes4,address,uint256),uint256),uint256,uint256,uint256,bytes4,bytes)[] orders) returns()
func (_NSwapExchange *NSwapExchangeTransactor) BatchCancel(opts *bind.TransactOpts, orders []LibOrderOrder) (*types.Transaction, error) {
	return _NSwapExchange.contract.Transact(opts, "batchCancel", orders)
}

// BatchCancel is a paid mutator transaction binding the contract method 0x2d46942b.
//
// Solidity: function batchCancel((uint8,uint8,address,address,((bytes4,address,uint256),uint256),((bytes4,address,uint256),uint256),uint256,uint256,uint256,bytes4,bytes)[] orders) returns()
func (_NSwapExchange *NSwapExchangeSession) BatchCancel(orders []LibOrderOrder) (*types.Transaction, error) {
	return _NSwapExchange.Contract.BatchCancel(&_NSwapExchange.TransactOpts, orders)
}

// BatchCancel is a paid mutator transaction binding the contract method 0x2d46942b.
//
// Solidity: function batchCancel((uint8,uint8,address,address,((bytes4,address,uint256),uint256),((bytes4,address,uint256),uint256),uint256,uint256,uint256,bytes4,bytes)[] orders) returns()
func (_NSwapExchange *NSwapExchangeTransactorSession) BatchCancel(orders []LibOrderOrder) (*types.Transaction, error) {
	return _NSwapExchange.Contract.BatchCancel(&_NSwapExchange.TransactOpts, orders)
}

// BatchMatch is a paid mutator transaction binding the contract method 0xf142878b.
//
// Solidity: function batchMatch(((uint8,uint8,address,address,((bytes4,address,uint256),uint256),((bytes4,address,uint256),uint256),uint256,uint256,uint256,bytes4,bytes),(uint8,uint8,address,address,((bytes4,address,uint256),uint256),((bytes4,address,uint256),uint256),uint256,uint256,uint256,bytes4,bytes),bytes)[] matchDetails) payable returns()
func (_NSwapExchange *NSwapExchangeTransactor) BatchMatch(opts *bind.TransactOpts, matchDetails []NSwapExchangeMatchDetails) (*types.Transaction, error) {
	return _NSwapExchange.contract.Transact(opts, "batchMatch", matchDetails)
}

// BatchMatch is a paid mutator transaction binding the contract method 0xf142878b.
//
// Solidity: function batchMatch(((uint8,uint8,address,address,((bytes4,address,uint256),uint256),((bytes4,address,uint256),uint256),uint256,uint256,uint256,bytes4,bytes),(uint8,uint8,address,address,((bytes4,address,uint256),uint256),((bytes4,address,uint256),uint256),uint256,uint256,uint256,bytes4,bytes),bytes)[] matchDetails) payable returns()
func (_NSwapExchange *NSwapExchangeSession) BatchMatch(matchDetails []NSwapExchangeMatchDetails) (*types.Transaction, error) {
	return _NSwapExchange.Contract.BatchMatch(&_NSwapExchange.TransactOpts, matchDetails)
}

// BatchMatch is a paid mutator transaction binding the contract method 0xf142878b.
//
// Solidity: function batchMatch(((uint8,uint8,address,address,((bytes4,address,uint256),uint256),((bytes4,address,uint256),uint256),uint256,uint256,uint256,bytes4,bytes),(uint8,uint8,address,address,((bytes4,address,uint256),uint256),((bytes4,address,uint256),uint256),uint256,uint256,uint256,bytes4,bytes),bytes)[] matchDetails) payable returns()
func (_NSwapExchange *NSwapExchangeTransactorSession) BatchMatch(matchDetails []NSwapExchangeMatchDetails) (*types.Transaction, error) {
	return _NSwapExchange.Contract.BatchMatch(&_NSwapExchange.TransactOpts, matchDetails)
}

// BatchTransferERC721 is a paid mutator transaction binding the contract method 0x3bb80e0e.
//
// Solidity: function batchTransferERC721(address to, (bytes4,address,uint256)[] assets) returns()
func (_NSwapExchange *NSwapExchangeTransactor) BatchTransferERC721(opts *bind.TransactOpts, to common.Address, assets []LibAssetAssetType) (*types.Transaction, error) {
	return _NSwapExchange.contract.Transact(opts, "batchTransferERC721", to, assets)
}

// BatchTransferERC721 is a paid mutator transaction binding the contract method 0x3bb80e0e.
//
// Solidity: function batchTransferERC721(address to, (bytes4,address,uint256)[] assets) returns()
func (_NSwapExchange *NSwapExchangeSession) BatchTransferERC721(to common.Address, assets []LibAssetAssetType) (*types.Transaction, error) {
	return _NSwapExchange.Contract.BatchTransferERC721(&_NSwapExchange.TransactOpts, to, assets)
}

// BatchTransferERC721 is a paid mutator transaction binding the contract method 0x3bb80e0e.
//
// Solidity: function batchTransferERC721(address to, (bytes4,address,uint256)[] assets) returns()
func (_NSwapExchange *NSwapExchangeTransactorSession) BatchTransferERC721(to common.Address, assets []LibAssetAssetType) (*types.Transaction, error) {
	return _NSwapExchange.Contract.BatchTransferERC721(&_NSwapExchange.TransactOpts, to, assets)
}

// Cancel is a paid mutator transaction binding the contract method 0xc118d526.
//
// Solidity: function cancel((uint8,uint8,address,address,((bytes4,address,uint256),uint256),((bytes4,address,uint256),uint256),uint256,uint256,uint256,bytes4,bytes) order) returns()
func (_NSwapExchange *NSwapExchangeTransactor) Cancel(opts *bind.TransactOpts, order LibOrderOrder) (*types.Transaction, error) {
	return _NSwapExchange.contract.Transact(opts, "cancel", order)
}

// Cancel is a paid mutator transaction binding the contract method 0xc118d526.
//
// Solidity: function cancel((uint8,uint8,address,address,((bytes4,address,uint256),uint256),((bytes4,address,uint256),uint256),uint256,uint256,uint256,bytes4,bytes) order) returns()
func (_NSwapExchange *NSwapExchangeSession) Cancel(order LibOrderOrder) (*types.Transaction, error) {
	return _NSwapExchange.Contract.Cancel(&_NSwapExchange.TransactOpts, order)
}

// Cancel is a paid mutator transaction binding the contract method 0xc118d526.
//
// Solidity: function cancel((uint8,uint8,address,address,((bytes4,address,uint256),uint256),((bytes4,address,uint256),uint256),uint256,uint256,uint256,bytes4,bytes) order) returns()
func (_NSwapExchange *NSwapExchangeTransactorSession) Cancel(order LibOrderOrder) (*types.Transaction, error) {
	return _NSwapExchange.Contract.Cancel(&_NSwapExchange.TransactOpts, order)
}

// ClearRoyaltiesType is a paid mutator transaction binding the contract method 0xfc73be00.
//
// Solidity: function clearRoyaltiesType(address token) returns()
func (_NSwapExchange *NSwapExchangeTransactor) ClearRoyaltiesType(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _NSwapExchange.contract.Transact(opts, "clearRoyaltiesType", token)
}

// ClearRoyaltiesType is a paid mutator transaction binding the contract method 0xfc73be00.
//
// Solidity: function clearRoyaltiesType(address token) returns()
func (_NSwapExchange *NSwapExchangeSession) ClearRoyaltiesType(token common.Address) (*types.Transaction, error) {
	return _NSwapExchange.Contract.ClearRoyaltiesType(&_NSwapExchange.TransactOpts, token)
}

// ClearRoyaltiesType is a paid mutator transaction binding the contract method 0xfc73be00.
//
// Solidity: function clearRoyaltiesType(address token) returns()
func (_NSwapExchange *NSwapExchangeTransactorSession) ClearRoyaltiesType(token common.Address) (*types.Transaction, error) {
	return _NSwapExchange.Contract.ClearRoyaltiesType(&_NSwapExchange.TransactOpts, token)
}

// GetRoyalties is a paid mutator transaction binding the contract method 0x9ca7dc7a.
//
// Solidity: function getRoyalties(address token, uint256 tokenId) returns((address,uint96)[])
func (_NSwapExchange *NSwapExchangeTransactor) GetRoyalties(opts *bind.TransactOpts, token common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NSwapExchange.contract.Transact(opts, "getRoyalties", token, tokenId)
}

// GetRoyalties is a paid mutator transaction binding the contract method 0x9ca7dc7a.
//
// Solidity: function getRoyalties(address token, uint256 tokenId) returns((address,uint96)[])
func (_NSwapExchange *NSwapExchangeSession) GetRoyalties(token common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NSwapExchange.Contract.GetRoyalties(&_NSwapExchange.TransactOpts, token, tokenId)
}

// GetRoyalties is a paid mutator transaction binding the contract method 0x9ca7dc7a.
//
// Solidity: function getRoyalties(address token, uint256 tokenId) returns((address,uint96)[])
func (_NSwapExchange *NSwapExchangeTransactorSession) GetRoyalties(token common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NSwapExchange.Contract.GetRoyalties(&_NSwapExchange.TransactOpts, token, tokenId)
}

// Initialize is a paid mutator transaction binding the contract method 0x10496465.
//
// Solidity: function initialize(address[] currencies, uint256 defaultProtocolFee, address defaultFeeRecipient, uint256 newRoyaltiesLimit, string EIP712Name, string EIP712Version) returns()
func (_NSwapExchange *NSwapExchangeTransactor) Initialize(opts *bind.TransactOpts, currencies []common.Address, defaultProtocolFee *big.Int, defaultFeeRecipient common.Address, newRoyaltiesLimit *big.Int, EIP712Name string, EIP712Version string) (*types.Transaction, error) {
	return _NSwapExchange.contract.Transact(opts, "initialize", currencies, defaultProtocolFee, defaultFeeRecipient, newRoyaltiesLimit, EIP712Name, EIP712Version)
}

// Initialize is a paid mutator transaction binding the contract method 0x10496465.
//
// Solidity: function initialize(address[] currencies, uint256 defaultProtocolFee, address defaultFeeRecipient, uint256 newRoyaltiesLimit, string EIP712Name, string EIP712Version) returns()
func (_NSwapExchange *NSwapExchangeSession) Initialize(currencies []common.Address, defaultProtocolFee *big.Int, defaultFeeRecipient common.Address, newRoyaltiesLimit *big.Int, EIP712Name string, EIP712Version string) (*types.Transaction, error) {
	return _NSwapExchange.Contract.Initialize(&_NSwapExchange.TransactOpts, currencies, defaultProtocolFee, defaultFeeRecipient, newRoyaltiesLimit, EIP712Name, EIP712Version)
}

// Initialize is a paid mutator transaction binding the contract method 0x10496465.
//
// Solidity: function initialize(address[] currencies, uint256 defaultProtocolFee, address defaultFeeRecipient, uint256 newRoyaltiesLimit, string EIP712Name, string EIP712Version) returns()
func (_NSwapExchange *NSwapExchangeTransactorSession) Initialize(currencies []common.Address, defaultProtocolFee *big.Int, defaultFeeRecipient common.Address, newRoyaltiesLimit *big.Int, EIP712Name string, EIP712Version string) (*types.Transaction, error) {
	return _NSwapExchange.Contract.Initialize(&_NSwapExchange.TransactOpts, currencies, defaultProtocolFee, defaultFeeRecipient, newRoyaltiesLimit, EIP712Name, EIP712Version)
}

// MatchOrders is a paid mutator transaction binding the contract method 0xa0cde1e8.
//
// Solidity: function matchOrders((uint8,uint8,address,address,((bytes4,address,uint256),uint256),((bytes4,address,uint256),uint256),uint256,uint256,uint256,bytes4,bytes) sell, (uint8,uint8,address,address,((bytes4,address,uint256),uint256),((bytes4,address,uint256),uint256),uint256,uint256,uint256,bytes4,bytes) buy, bytes signature) payable returns()
func (_NSwapExchange *NSwapExchangeTransactor) MatchOrders(opts *bind.TransactOpts, sell LibOrderOrder, buy LibOrderOrder, signature []byte) (*types.Transaction, error) {
	return _NSwapExchange.contract.Transact(opts, "matchOrders", sell, buy, signature)
}

// MatchOrders is a paid mutator transaction binding the contract method 0xa0cde1e8.
//
// Solidity: function matchOrders((uint8,uint8,address,address,((bytes4,address,uint256),uint256),((bytes4,address,uint256),uint256),uint256,uint256,uint256,bytes4,bytes) sell, (uint8,uint8,address,address,((bytes4,address,uint256),uint256),((bytes4,address,uint256),uint256),uint256,uint256,uint256,bytes4,bytes) buy, bytes signature) payable returns()
func (_NSwapExchange *NSwapExchangeSession) MatchOrders(sell LibOrderOrder, buy LibOrderOrder, signature []byte) (*types.Transaction, error) {
	return _NSwapExchange.Contract.MatchOrders(&_NSwapExchange.TransactOpts, sell, buy, signature)
}

// MatchOrders is a paid mutator transaction binding the contract method 0xa0cde1e8.
//
// Solidity: function matchOrders((uint8,uint8,address,address,((bytes4,address,uint256),uint256),((bytes4,address,uint256),uint256),uint256,uint256,uint256,bytes4,bytes) sell, (uint8,uint8,address,address,((bytes4,address,uint256),uint256),((bytes4,address,uint256),uint256),uint256,uint256,uint256,bytes4,bytes) buy, bytes signature) payable returns()
func (_NSwapExchange *NSwapExchangeTransactorSession) MatchOrders(sell LibOrderOrder, buy LibOrderOrder, signature []byte) (*types.Transaction, error) {
	return _NSwapExchange.Contract.MatchOrders(&_NSwapExchange.TransactOpts, sell, buy, signature)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_NSwapExchange *NSwapExchangeTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NSwapExchange.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_NSwapExchange *NSwapExchangeSession) Pause() (*types.Transaction, error) {
	return _NSwapExchange.Contract.Pause(&_NSwapExchange.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_NSwapExchange *NSwapExchangeTransactorSession) Pause() (*types.Transaction, error) {
	return _NSwapExchange.Contract.Pause(&_NSwapExchange.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NSwapExchange *NSwapExchangeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NSwapExchange.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NSwapExchange *NSwapExchangeSession) RenounceOwnership() (*types.Transaction, error) {
	return _NSwapExchange.Contract.RenounceOwnership(&_NSwapExchange.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NSwapExchange *NSwapExchangeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _NSwapExchange.Contract.RenounceOwnership(&_NSwapExchange.TransactOpts)
}

// RevokeCurrency is a paid mutator transaction binding the contract method 0x324a01fd.
//
// Solidity: function revokeCurrency(address currency) returns()
func (_NSwapExchange *NSwapExchangeTransactor) RevokeCurrency(opts *bind.TransactOpts, currency common.Address) (*types.Transaction, error) {
	return _NSwapExchange.contract.Transact(opts, "revokeCurrency", currency)
}

// RevokeCurrency is a paid mutator transaction binding the contract method 0x324a01fd.
//
// Solidity: function revokeCurrency(address currency) returns()
func (_NSwapExchange *NSwapExchangeSession) RevokeCurrency(currency common.Address) (*types.Transaction, error) {
	return _NSwapExchange.Contract.RevokeCurrency(&_NSwapExchange.TransactOpts, currency)
}

// RevokeCurrency is a paid mutator transaction binding the contract method 0x324a01fd.
//
// Solidity: function revokeCurrency(address currency) returns()
func (_NSwapExchange *NSwapExchangeTransactorSession) RevokeCurrency(currency common.Address) (*types.Transaction, error) {
	return _NSwapExchange.Contract.RevokeCurrency(&_NSwapExchange.TransactOpts, currency)
}

// SetDefaultProtocolFee is a paid mutator transaction binding the contract method 0x6034b510.
//
// Solidity: function setDefaultProtocolFee(uint256 newDefaultProtocolFee) returns()
func (_NSwapExchange *NSwapExchangeTransactor) SetDefaultProtocolFee(opts *bind.TransactOpts, newDefaultProtocolFee *big.Int) (*types.Transaction, error) {
	return _NSwapExchange.contract.Transact(opts, "setDefaultProtocolFee", newDefaultProtocolFee)
}

// SetDefaultProtocolFee is a paid mutator transaction binding the contract method 0x6034b510.
//
// Solidity: function setDefaultProtocolFee(uint256 newDefaultProtocolFee) returns()
func (_NSwapExchange *NSwapExchangeSession) SetDefaultProtocolFee(newDefaultProtocolFee *big.Int) (*types.Transaction, error) {
	return _NSwapExchange.Contract.SetDefaultProtocolFee(&_NSwapExchange.TransactOpts, newDefaultProtocolFee)
}

// SetDefaultProtocolFee is a paid mutator transaction binding the contract method 0x6034b510.
//
// Solidity: function setDefaultProtocolFee(uint256 newDefaultProtocolFee) returns()
func (_NSwapExchange *NSwapExchangeTransactorSession) SetDefaultProtocolFee(newDefaultProtocolFee *big.Int) (*types.Transaction, error) {
	return _NSwapExchange.Contract.SetDefaultProtocolFee(&_NSwapExchange.TransactOpts, newDefaultProtocolFee)
}

// SetDefaultProtocolFeeReceiver is a paid mutator transaction binding the contract method 0xf2ba9847.
//
// Solidity: function setDefaultProtocolFeeReceiver(address newDefaultProtocolFeeReceiver) returns()
func (_NSwapExchange *NSwapExchangeTransactor) SetDefaultProtocolFeeReceiver(opts *bind.TransactOpts, newDefaultProtocolFeeReceiver common.Address) (*types.Transaction, error) {
	return _NSwapExchange.contract.Transact(opts, "setDefaultProtocolFeeReceiver", newDefaultProtocolFeeReceiver)
}

// SetDefaultProtocolFeeReceiver is a paid mutator transaction binding the contract method 0xf2ba9847.
//
// Solidity: function setDefaultProtocolFeeReceiver(address newDefaultProtocolFeeReceiver) returns()
func (_NSwapExchange *NSwapExchangeSession) SetDefaultProtocolFeeReceiver(newDefaultProtocolFeeReceiver common.Address) (*types.Transaction, error) {
	return _NSwapExchange.Contract.SetDefaultProtocolFeeReceiver(&_NSwapExchange.TransactOpts, newDefaultProtocolFeeReceiver)
}

// SetDefaultProtocolFeeReceiver is a paid mutator transaction binding the contract method 0xf2ba9847.
//
// Solidity: function setDefaultProtocolFeeReceiver(address newDefaultProtocolFeeReceiver) returns()
func (_NSwapExchange *NSwapExchangeTransactorSession) SetDefaultProtocolFeeReceiver(newDefaultProtocolFeeReceiver common.Address) (*types.Transaction, error) {
	return _NSwapExchange.Contract.SetDefaultProtocolFeeReceiver(&_NSwapExchange.TransactOpts, newDefaultProtocolFeeReceiver)
}

// SetProtocolFee is a paid mutator transaction binding the contract method 0xb5b3ca2c.
//
// Solidity: function setProtocolFee(address wallet, uint256 newProtocolFee) returns()
func (_NSwapExchange *NSwapExchangeTransactor) SetProtocolFee(opts *bind.TransactOpts, wallet common.Address, newProtocolFee *big.Int) (*types.Transaction, error) {
	return _NSwapExchange.contract.Transact(opts, "setProtocolFee", wallet, newProtocolFee)
}

// SetProtocolFee is a paid mutator transaction binding the contract method 0xb5b3ca2c.
//
// Solidity: function setProtocolFee(address wallet, uint256 newProtocolFee) returns()
func (_NSwapExchange *NSwapExchangeSession) SetProtocolFee(wallet common.Address, newProtocolFee *big.Int) (*types.Transaction, error) {
	return _NSwapExchange.Contract.SetProtocolFee(&_NSwapExchange.TransactOpts, wallet, newProtocolFee)
}

// SetProtocolFee is a paid mutator transaction binding the contract method 0xb5b3ca2c.
//
// Solidity: function setProtocolFee(address wallet, uint256 newProtocolFee) returns()
func (_NSwapExchange *NSwapExchangeTransactorSession) SetProtocolFee(wallet common.Address, newProtocolFee *big.Int) (*types.Transaction, error) {
	return _NSwapExchange.Contract.SetProtocolFee(&_NSwapExchange.TransactOpts, wallet, newProtocolFee)
}

// SetRoyaltiesByToken is a paid mutator transaction binding the contract method 0xacf14efb.
//
// Solidity: function setRoyaltiesByToken(address token, (address,uint96)[] royalties) returns()
func (_NSwapExchange *NSwapExchangeTransactor) SetRoyaltiesByToken(opts *bind.TransactOpts, token common.Address, royalties []LibPayInfoPayInfo) (*types.Transaction, error) {
	return _NSwapExchange.contract.Transact(opts, "setRoyaltiesByToken", token, royalties)
}

// SetRoyaltiesByToken is a paid mutator transaction binding the contract method 0xacf14efb.
//
// Solidity: function setRoyaltiesByToken(address token, (address,uint96)[] royalties) returns()
func (_NSwapExchange *NSwapExchangeSession) SetRoyaltiesByToken(token common.Address, royalties []LibPayInfoPayInfo) (*types.Transaction, error) {
	return _NSwapExchange.Contract.SetRoyaltiesByToken(&_NSwapExchange.TransactOpts, token, royalties)
}

// SetRoyaltiesByToken is a paid mutator transaction binding the contract method 0xacf14efb.
//
// Solidity: function setRoyaltiesByToken(address token, (address,uint96)[] royalties) returns()
func (_NSwapExchange *NSwapExchangeTransactorSession) SetRoyaltiesByToken(token common.Address, royalties []LibPayInfoPayInfo) (*types.Transaction, error) {
	return _NSwapExchange.Contract.SetRoyaltiesByToken(&_NSwapExchange.TransactOpts, token, royalties)
}

// SetRoyaltiesType is a paid mutator transaction binding the contract method 0xb65e334d.
//
// Solidity: function setRoyaltiesType(address token, uint256 newRoyaltiesType) returns()
func (_NSwapExchange *NSwapExchangeTransactor) SetRoyaltiesType(opts *bind.TransactOpts, token common.Address, newRoyaltiesType *big.Int) (*types.Transaction, error) {
	return _NSwapExchange.contract.Transact(opts, "setRoyaltiesType", token, newRoyaltiesType)
}

// SetRoyaltiesType is a paid mutator transaction binding the contract method 0xb65e334d.
//
// Solidity: function setRoyaltiesType(address token, uint256 newRoyaltiesType) returns()
func (_NSwapExchange *NSwapExchangeSession) SetRoyaltiesType(token common.Address, newRoyaltiesType *big.Int) (*types.Transaction, error) {
	return _NSwapExchange.Contract.SetRoyaltiesType(&_NSwapExchange.TransactOpts, token, newRoyaltiesType)
}

// SetRoyaltiesType is a paid mutator transaction binding the contract method 0xb65e334d.
//
// Solidity: function setRoyaltiesType(address token, uint256 newRoyaltiesType) returns()
func (_NSwapExchange *NSwapExchangeTransactorSession) SetRoyaltiesType(token common.Address, newRoyaltiesType *big.Int) (*types.Transaction, error) {
	return _NSwapExchange.Contract.SetRoyaltiesType(&_NSwapExchange.TransactOpts, token, newRoyaltiesType)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NSwapExchange *NSwapExchangeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _NSwapExchange.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NSwapExchange *NSwapExchangeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NSwapExchange.Contract.TransferOwnership(&_NSwapExchange.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NSwapExchange *NSwapExchangeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NSwapExchange.Contract.TransferOwnership(&_NSwapExchange.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_NSwapExchange *NSwapExchangeTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NSwapExchange.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_NSwapExchange *NSwapExchangeSession) Unpause() (*types.Transaction, error) {
	return _NSwapExchange.Contract.Unpause(&_NSwapExchange.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_NSwapExchange *NSwapExchangeTransactorSession) Unpause() (*types.Transaction, error) {
	return _NSwapExchange.Contract.Unpause(&_NSwapExchange.TransactOpts)
}

// UpdateRoyaltiesLimit is a paid mutator transaction binding the contract method 0xf83e7999.
//
// Solidity: function updateRoyaltiesLimit(uint256 newRoyaltiesLimit) returns()
func (_NSwapExchange *NSwapExchangeTransactor) UpdateRoyaltiesLimit(opts *bind.TransactOpts, newRoyaltiesLimit *big.Int) (*types.Transaction, error) {
	return _NSwapExchange.contract.Transact(opts, "updateRoyaltiesLimit", newRoyaltiesLimit)
}

// UpdateRoyaltiesLimit is a paid mutator transaction binding the contract method 0xf83e7999.
//
// Solidity: function updateRoyaltiesLimit(uint256 newRoyaltiesLimit) returns()
func (_NSwapExchange *NSwapExchangeSession) UpdateRoyaltiesLimit(newRoyaltiesLimit *big.Int) (*types.Transaction, error) {
	return _NSwapExchange.Contract.UpdateRoyaltiesLimit(&_NSwapExchange.TransactOpts, newRoyaltiesLimit)
}

// UpdateRoyaltiesLimit is a paid mutator transaction binding the contract method 0xf83e7999.
//
// Solidity: function updateRoyaltiesLimit(uint256 newRoyaltiesLimit) returns()
func (_NSwapExchange *NSwapExchangeTransactorSession) UpdateRoyaltiesLimit(newRoyaltiesLimit *big.Int) (*types.Transaction, error) {
	return _NSwapExchange.Contract.UpdateRoyaltiesLimit(&_NSwapExchange.TransactOpts, newRoyaltiesLimit)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x4782f779.
//
// Solidity: function withdrawETH(address recipient, uint256 amount) returns()
func (_NSwapExchange *NSwapExchangeTransactor) WithdrawETH(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NSwapExchange.contract.Transact(opts, "withdrawETH", recipient, amount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x4782f779.
//
// Solidity: function withdrawETH(address recipient, uint256 amount) returns()
func (_NSwapExchange *NSwapExchangeSession) WithdrawETH(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NSwapExchange.Contract.WithdrawETH(&_NSwapExchange.TransactOpts, recipient, amount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x4782f779.
//
// Solidity: function withdrawETH(address recipient, uint256 amount) returns()
func (_NSwapExchange *NSwapExchangeTransactorSession) WithdrawETH(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NSwapExchange.Contract.WithdrawETH(&_NSwapExchange.TransactOpts, recipient, amount)
}

// NSwapExchangeEventCancelIterator is returned from FilterEventCancel and is used to iterate over the raw logs and unpacked data for EventCancel events raised by the NSwapExchange contract.
type NSwapExchangeEventCancelIterator struct {
	Event *NSwapExchangeEventCancel // Event containing the contract specifics and raw log

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
func (it *NSwapExchangeEventCancelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NSwapExchangeEventCancel)
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
		it.Event = new(NSwapExchangeEventCancel)
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
func (it *NSwapExchangeEventCancelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NSwapExchangeEventCancelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NSwapExchangeEventCancel represents a EventCancel event raised by the NSwapExchange contract.
type NSwapExchangeEventCancel struct {
	Hash           [32]byte
	Maker          common.Address
	PriceAssetType LibAssetAssetType
	NftAssetType   LibAssetAssetType
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterEventCancel is a free log retrieval operation binding the contract event 0xa6176bdddd641ae5eece4d50959123f875f97339d0f6a61260f2dc8f5316bff1.
//
// Solidity: event EventCancel(bytes32 hash, address maker, (bytes4,address,uint256) priceAssetType, (bytes4,address,uint256) nftAssetType)
func (_NSwapExchange *NSwapExchangeFilterer) FilterEventCancel(opts *bind.FilterOpts) (*NSwapExchangeEventCancelIterator, error) {

	logs, sub, err := _NSwapExchange.contract.FilterLogs(opts, "EventCancel")
	if err != nil {
		return nil, err
	}
	return &NSwapExchangeEventCancelIterator{contract: _NSwapExchange.contract, event: "EventCancel", logs: logs, sub: sub}, nil
}

// WatchEventCancel is a free log subscription operation binding the contract event 0xa6176bdddd641ae5eece4d50959123f875f97339d0f6a61260f2dc8f5316bff1.
//
// Solidity: event EventCancel(bytes32 hash, address maker, (bytes4,address,uint256) priceAssetType, (bytes4,address,uint256) nftAssetType)
func (_NSwapExchange *NSwapExchangeFilterer) WatchEventCancel(opts *bind.WatchOpts, sink chan<- *NSwapExchangeEventCancel) (event.Subscription, error) {

	logs, sub, err := _NSwapExchange.contract.WatchLogs(opts, "EventCancel")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NSwapExchangeEventCancel)
				if err := _NSwapExchange.contract.UnpackLog(event, "EventCancel", log); err != nil {
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

// ParseEventCancel is a log parse operation binding the contract event 0xa6176bdddd641ae5eece4d50959123f875f97339d0f6a61260f2dc8f5316bff1.
//
// Solidity: event EventCancel(bytes32 hash, address maker, (bytes4,address,uint256) priceAssetType, (bytes4,address,uint256) nftAssetType)
func (_NSwapExchange *NSwapExchangeFilterer) ParseEventCancel(log types.Log) (*NSwapExchangeEventCancel, error) {
	event := new(NSwapExchangeEventCancel)
	if err := _NSwapExchange.contract.UnpackLog(event, "EventCancel", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NSwapExchangeEventCurrencyApprovedIterator is returned from FilterEventCurrencyApproved and is used to iterate over the raw logs and unpacked data for EventCurrencyApproved events raised by the NSwapExchange contract.
type NSwapExchangeEventCurrencyApprovedIterator struct {
	Event *NSwapExchangeEventCurrencyApproved // Event containing the contract specifics and raw log

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
func (it *NSwapExchangeEventCurrencyApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NSwapExchangeEventCurrencyApproved)
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
		it.Event = new(NSwapExchangeEventCurrencyApproved)
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
func (it *NSwapExchangeEventCurrencyApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NSwapExchangeEventCurrencyApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NSwapExchangeEventCurrencyApproved represents a EventCurrencyApproved event raised by the NSwapExchange contract.
type NSwapExchangeEventCurrencyApproved struct {
	Currency common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterEventCurrencyApproved is a free log retrieval operation binding the contract event 0xac4fd32d89d9af9478e1baddbbe51b1860698884f84170d738ada493a2f4293f.
//
// Solidity: event EventCurrencyApproved(address indexed currency)
func (_NSwapExchange *NSwapExchangeFilterer) FilterEventCurrencyApproved(opts *bind.FilterOpts, currency []common.Address) (*NSwapExchangeEventCurrencyApprovedIterator, error) {

	var currencyRule []interface{}
	for _, currencyItem := range currency {
		currencyRule = append(currencyRule, currencyItem)
	}

	logs, sub, err := _NSwapExchange.contract.FilterLogs(opts, "EventCurrencyApproved", currencyRule)
	if err != nil {
		return nil, err
	}
	return &NSwapExchangeEventCurrencyApprovedIterator{contract: _NSwapExchange.contract, event: "EventCurrencyApproved", logs: logs, sub: sub}, nil
}

// WatchEventCurrencyApproved is a free log subscription operation binding the contract event 0xac4fd32d89d9af9478e1baddbbe51b1860698884f84170d738ada493a2f4293f.
//
// Solidity: event EventCurrencyApproved(address indexed currency)
func (_NSwapExchange *NSwapExchangeFilterer) WatchEventCurrencyApproved(opts *bind.WatchOpts, sink chan<- *NSwapExchangeEventCurrencyApproved, currency []common.Address) (event.Subscription, error) {

	var currencyRule []interface{}
	for _, currencyItem := range currency {
		currencyRule = append(currencyRule, currencyItem)
	}

	logs, sub, err := _NSwapExchange.contract.WatchLogs(opts, "EventCurrencyApproved", currencyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NSwapExchangeEventCurrencyApproved)
				if err := _NSwapExchange.contract.UnpackLog(event, "EventCurrencyApproved", log); err != nil {
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

// ParseEventCurrencyApproved is a log parse operation binding the contract event 0xac4fd32d89d9af9478e1baddbbe51b1860698884f84170d738ada493a2f4293f.
//
// Solidity: event EventCurrencyApproved(address indexed currency)
func (_NSwapExchange *NSwapExchangeFilterer) ParseEventCurrencyApproved(log types.Log) (*NSwapExchangeEventCurrencyApproved, error) {
	event := new(NSwapExchangeEventCurrencyApproved)
	if err := _NSwapExchange.contract.UnpackLog(event, "EventCurrencyApproved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NSwapExchangeEventCurrencyRevokedIterator is returned from FilterEventCurrencyRevoked and is used to iterate over the raw logs and unpacked data for EventCurrencyRevoked events raised by the NSwapExchange contract.
type NSwapExchangeEventCurrencyRevokedIterator struct {
	Event *NSwapExchangeEventCurrencyRevoked // Event containing the contract specifics and raw log

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
func (it *NSwapExchangeEventCurrencyRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NSwapExchangeEventCurrencyRevoked)
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
		it.Event = new(NSwapExchangeEventCurrencyRevoked)
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
func (it *NSwapExchangeEventCurrencyRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NSwapExchangeEventCurrencyRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NSwapExchangeEventCurrencyRevoked represents a EventCurrencyRevoked event raised by the NSwapExchange contract.
type NSwapExchangeEventCurrencyRevoked struct {
	Currency common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterEventCurrencyRevoked is a free log retrieval operation binding the contract event 0xb67d8eb3a61372d2181634e49962d41293df2708d75a7f40ee8155227819bb5b.
//
// Solidity: event EventCurrencyRevoked(address indexed currency)
func (_NSwapExchange *NSwapExchangeFilterer) FilterEventCurrencyRevoked(opts *bind.FilterOpts, currency []common.Address) (*NSwapExchangeEventCurrencyRevokedIterator, error) {

	var currencyRule []interface{}
	for _, currencyItem := range currency {
		currencyRule = append(currencyRule, currencyItem)
	}

	logs, sub, err := _NSwapExchange.contract.FilterLogs(opts, "EventCurrencyRevoked", currencyRule)
	if err != nil {
		return nil, err
	}
	return &NSwapExchangeEventCurrencyRevokedIterator{contract: _NSwapExchange.contract, event: "EventCurrencyRevoked", logs: logs, sub: sub}, nil
}

// WatchEventCurrencyRevoked is a free log subscription operation binding the contract event 0xb67d8eb3a61372d2181634e49962d41293df2708d75a7f40ee8155227819bb5b.
//
// Solidity: event EventCurrencyRevoked(address indexed currency)
func (_NSwapExchange *NSwapExchangeFilterer) WatchEventCurrencyRevoked(opts *bind.WatchOpts, sink chan<- *NSwapExchangeEventCurrencyRevoked, currency []common.Address) (event.Subscription, error) {

	var currencyRule []interface{}
	for _, currencyItem := range currency {
		currencyRule = append(currencyRule, currencyItem)
	}

	logs, sub, err := _NSwapExchange.contract.WatchLogs(opts, "EventCurrencyRevoked", currencyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NSwapExchangeEventCurrencyRevoked)
				if err := _NSwapExchange.contract.UnpackLog(event, "EventCurrencyRevoked", log); err != nil {
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

// ParseEventCurrencyRevoked is a log parse operation binding the contract event 0xb67d8eb3a61372d2181634e49962d41293df2708d75a7f40ee8155227819bb5b.
//
// Solidity: event EventCurrencyRevoked(address indexed currency)
func (_NSwapExchange *NSwapExchangeFilterer) ParseEventCurrencyRevoked(log types.Log) (*NSwapExchangeEventCurrencyRevoked, error) {
	event := new(NSwapExchangeEventCurrencyRevoked)
	if err := _NSwapExchange.contract.UnpackLog(event, "EventCurrencyRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NSwapExchangeEventMatchIterator is returned from FilterEventMatch and is used to iterate over the raw logs and unpacked data for EventMatch events raised by the NSwapExchange contract.
type NSwapExchangeEventMatchIterator struct {
	Event *NSwapExchangeEventMatch // Event containing the contract specifics and raw log

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
func (it *NSwapExchangeEventMatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NSwapExchangeEventMatch)
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
		it.Event = new(NSwapExchangeEventMatch)
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
func (it *NSwapExchangeEventMatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NSwapExchangeEventMatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NSwapExchangeEventMatch represents a EventMatch event raised by the NSwapExchange contract.
type NSwapExchangeEventMatch struct {
	BuyHash        [32]byte
	SellHash       [32]byte
	BuyMaker       common.Address
	SellMaker      common.Address
	PriceFilled    *big.Int
	NftFilled      *big.Int
	PriceAssetType LibAssetAssetType
	NftAssetType   LibAssetAssetType
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterEventMatch is a free log retrieval operation binding the contract event 0xeda59ad80e2778cfaad4edd371c6b5c9f3abdf97cc7651d9649b1fdb41af2b8d.
//
// Solidity: event EventMatch(bytes32 buyHash, bytes32 sellHash, address buyMaker, address sellMaker, uint256 priceFilled, uint256 nftFilled, (bytes4,address,uint256) priceAssetType, (bytes4,address,uint256) nftAssetType)
func (_NSwapExchange *NSwapExchangeFilterer) FilterEventMatch(opts *bind.FilterOpts) (*NSwapExchangeEventMatchIterator, error) {

	logs, sub, err := _NSwapExchange.contract.FilterLogs(opts, "EventMatch")
	if err != nil {
		return nil, err
	}
	return &NSwapExchangeEventMatchIterator{contract: _NSwapExchange.contract, event: "EventMatch", logs: logs, sub: sub}, nil
}

// WatchEventMatch is a free log subscription operation binding the contract event 0xeda59ad80e2778cfaad4edd371c6b5c9f3abdf97cc7651d9649b1fdb41af2b8d.
//
// Solidity: event EventMatch(bytes32 buyHash, bytes32 sellHash, address buyMaker, address sellMaker, uint256 priceFilled, uint256 nftFilled, (bytes4,address,uint256) priceAssetType, (bytes4,address,uint256) nftAssetType)
func (_NSwapExchange *NSwapExchangeFilterer) WatchEventMatch(opts *bind.WatchOpts, sink chan<- *NSwapExchangeEventMatch) (event.Subscription, error) {

	logs, sub, err := _NSwapExchange.contract.WatchLogs(opts, "EventMatch")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NSwapExchangeEventMatch)
				if err := _NSwapExchange.contract.UnpackLog(event, "EventMatch", log); err != nil {
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

// ParseEventMatch is a log parse operation binding the contract event 0xeda59ad80e2778cfaad4edd371c6b5c9f3abdf97cc7651d9649b1fdb41af2b8d.
//
// Solidity: event EventMatch(bytes32 buyHash, bytes32 sellHash, address buyMaker, address sellMaker, uint256 priceFilled, uint256 nftFilled, (bytes4,address,uint256) priceAssetType, (bytes4,address,uint256) nftAssetType)
func (_NSwapExchange *NSwapExchangeFilterer) ParseEventMatch(log types.Log) (*NSwapExchangeEventMatch, error) {
	event := new(NSwapExchangeEventMatch)
	if err := _NSwapExchange.contract.UnpackLog(event, "EventMatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NSwapExchangeEventUpdatedCustomProtocolFeeIterator is returned from FilterEventUpdatedCustomProtocolFee and is used to iterate over the raw logs and unpacked data for EventUpdatedCustomProtocolFee events raised by the NSwapExchange contract.
type NSwapExchangeEventUpdatedCustomProtocolFeeIterator struct {
	Event *NSwapExchangeEventUpdatedCustomProtocolFee // Event containing the contract specifics and raw log

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
func (it *NSwapExchangeEventUpdatedCustomProtocolFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NSwapExchangeEventUpdatedCustomProtocolFee)
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
		it.Event = new(NSwapExchangeEventUpdatedCustomProtocolFee)
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
func (it *NSwapExchangeEventUpdatedCustomProtocolFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NSwapExchangeEventUpdatedCustomProtocolFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NSwapExchangeEventUpdatedCustomProtocolFee represents a EventUpdatedCustomProtocolFee event raised by the NSwapExchange contract.
type NSwapExchangeEventUpdatedCustomProtocolFee struct {
	CustomAccount common.Address
	Fee           *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterEventUpdatedCustomProtocolFee is a free log retrieval operation binding the contract event 0x397347454c83dc973ba72a51102393bcf996b92d89a64abc335c10fc1cadbb12.
//
// Solidity: event EventUpdatedCustomProtocolFee(address indexed customAccount, uint256 indexed fee)
func (_NSwapExchange *NSwapExchangeFilterer) FilterEventUpdatedCustomProtocolFee(opts *bind.FilterOpts, customAccount []common.Address, fee []*big.Int) (*NSwapExchangeEventUpdatedCustomProtocolFeeIterator, error) {

	var customAccountRule []interface{}
	for _, customAccountItem := range customAccount {
		customAccountRule = append(customAccountRule, customAccountItem)
	}
	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}

	logs, sub, err := _NSwapExchange.contract.FilterLogs(opts, "EventUpdatedCustomProtocolFee", customAccountRule, feeRule)
	if err != nil {
		return nil, err
	}
	return &NSwapExchangeEventUpdatedCustomProtocolFeeIterator{contract: _NSwapExchange.contract, event: "EventUpdatedCustomProtocolFee", logs: logs, sub: sub}, nil
}

// WatchEventUpdatedCustomProtocolFee is a free log subscription operation binding the contract event 0x397347454c83dc973ba72a51102393bcf996b92d89a64abc335c10fc1cadbb12.
//
// Solidity: event EventUpdatedCustomProtocolFee(address indexed customAccount, uint256 indexed fee)
func (_NSwapExchange *NSwapExchangeFilterer) WatchEventUpdatedCustomProtocolFee(opts *bind.WatchOpts, sink chan<- *NSwapExchangeEventUpdatedCustomProtocolFee, customAccount []common.Address, fee []*big.Int) (event.Subscription, error) {

	var customAccountRule []interface{}
	for _, customAccountItem := range customAccount {
		customAccountRule = append(customAccountRule, customAccountItem)
	}
	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}

	logs, sub, err := _NSwapExchange.contract.WatchLogs(opts, "EventUpdatedCustomProtocolFee", customAccountRule, feeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NSwapExchangeEventUpdatedCustomProtocolFee)
				if err := _NSwapExchange.contract.UnpackLog(event, "EventUpdatedCustomProtocolFee", log); err != nil {
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

// ParseEventUpdatedCustomProtocolFee is a log parse operation binding the contract event 0x397347454c83dc973ba72a51102393bcf996b92d89a64abc335c10fc1cadbb12.
//
// Solidity: event EventUpdatedCustomProtocolFee(address indexed customAccount, uint256 indexed fee)
func (_NSwapExchange *NSwapExchangeFilterer) ParseEventUpdatedCustomProtocolFee(log types.Log) (*NSwapExchangeEventUpdatedCustomProtocolFee, error) {
	event := new(NSwapExchangeEventUpdatedCustomProtocolFee)
	if err := _NSwapExchange.contract.UnpackLog(event, "EventUpdatedCustomProtocolFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NSwapExchangeEventUpdatedDefaultProtocolFeeIterator is returned from FilterEventUpdatedDefaultProtocolFee and is used to iterate over the raw logs and unpacked data for EventUpdatedDefaultProtocolFee events raised by the NSwapExchange contract.
type NSwapExchangeEventUpdatedDefaultProtocolFeeIterator struct {
	Event *NSwapExchangeEventUpdatedDefaultProtocolFee // Event containing the contract specifics and raw log

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
func (it *NSwapExchangeEventUpdatedDefaultProtocolFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NSwapExchangeEventUpdatedDefaultProtocolFee)
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
		it.Event = new(NSwapExchangeEventUpdatedDefaultProtocolFee)
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
func (it *NSwapExchangeEventUpdatedDefaultProtocolFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NSwapExchangeEventUpdatedDefaultProtocolFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NSwapExchangeEventUpdatedDefaultProtocolFee represents a EventUpdatedDefaultProtocolFee event raised by the NSwapExchange contract.
type NSwapExchangeEventUpdatedDefaultProtocolFee struct {
	DefaultFee *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterEventUpdatedDefaultProtocolFee is a free log retrieval operation binding the contract event 0x8002d04443bb78a13249336dc6835c06ac183aad2721961625c0cfd37afc2261.
//
// Solidity: event EventUpdatedDefaultProtocolFee(uint256 indexed defaultFee)
func (_NSwapExchange *NSwapExchangeFilterer) FilterEventUpdatedDefaultProtocolFee(opts *bind.FilterOpts, defaultFee []*big.Int) (*NSwapExchangeEventUpdatedDefaultProtocolFeeIterator, error) {

	var defaultFeeRule []interface{}
	for _, defaultFeeItem := range defaultFee {
		defaultFeeRule = append(defaultFeeRule, defaultFeeItem)
	}

	logs, sub, err := _NSwapExchange.contract.FilterLogs(opts, "EventUpdatedDefaultProtocolFee", defaultFeeRule)
	if err != nil {
		return nil, err
	}
	return &NSwapExchangeEventUpdatedDefaultProtocolFeeIterator{contract: _NSwapExchange.contract, event: "EventUpdatedDefaultProtocolFee", logs: logs, sub: sub}, nil
}

// WatchEventUpdatedDefaultProtocolFee is a free log subscription operation binding the contract event 0x8002d04443bb78a13249336dc6835c06ac183aad2721961625c0cfd37afc2261.
//
// Solidity: event EventUpdatedDefaultProtocolFee(uint256 indexed defaultFee)
func (_NSwapExchange *NSwapExchangeFilterer) WatchEventUpdatedDefaultProtocolFee(opts *bind.WatchOpts, sink chan<- *NSwapExchangeEventUpdatedDefaultProtocolFee, defaultFee []*big.Int) (event.Subscription, error) {

	var defaultFeeRule []interface{}
	for _, defaultFeeItem := range defaultFee {
		defaultFeeRule = append(defaultFeeRule, defaultFeeItem)
	}

	logs, sub, err := _NSwapExchange.contract.WatchLogs(opts, "EventUpdatedDefaultProtocolFee", defaultFeeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NSwapExchangeEventUpdatedDefaultProtocolFee)
				if err := _NSwapExchange.contract.UnpackLog(event, "EventUpdatedDefaultProtocolFee", log); err != nil {
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

// ParseEventUpdatedDefaultProtocolFee is a log parse operation binding the contract event 0x8002d04443bb78a13249336dc6835c06ac183aad2721961625c0cfd37afc2261.
//
// Solidity: event EventUpdatedDefaultProtocolFee(uint256 indexed defaultFee)
func (_NSwapExchange *NSwapExchangeFilterer) ParseEventUpdatedDefaultProtocolFee(log types.Log) (*NSwapExchangeEventUpdatedDefaultProtocolFee, error) {
	event := new(NSwapExchangeEventUpdatedDefaultProtocolFee)
	if err := _NSwapExchange.contract.UnpackLog(event, "EventUpdatedDefaultProtocolFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NSwapExchangeEventUpdatedDefaultProtocolFeeReceiverIterator is returned from FilterEventUpdatedDefaultProtocolFeeReceiver and is used to iterate over the raw logs and unpacked data for EventUpdatedDefaultProtocolFeeReceiver events raised by the NSwapExchange contract.
type NSwapExchangeEventUpdatedDefaultProtocolFeeReceiverIterator struct {
	Event *NSwapExchangeEventUpdatedDefaultProtocolFeeReceiver // Event containing the contract specifics and raw log

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
func (it *NSwapExchangeEventUpdatedDefaultProtocolFeeReceiverIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NSwapExchangeEventUpdatedDefaultProtocolFeeReceiver)
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
		it.Event = new(NSwapExchangeEventUpdatedDefaultProtocolFeeReceiver)
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
func (it *NSwapExchangeEventUpdatedDefaultProtocolFeeReceiverIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NSwapExchangeEventUpdatedDefaultProtocolFeeReceiverIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NSwapExchangeEventUpdatedDefaultProtocolFeeReceiver represents a EventUpdatedDefaultProtocolFeeReceiver event raised by the NSwapExchange contract.
type NSwapExchangeEventUpdatedDefaultProtocolFeeReceiver struct {
	DefaultFeeReceiver common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterEventUpdatedDefaultProtocolFeeReceiver is a free log retrieval operation binding the contract event 0xf71cba06d4627dfe19c434c00bbe1d60d48bdfcfc27781ba977953777e7ee2ee.
//
// Solidity: event EventUpdatedDefaultProtocolFeeReceiver(address indexed defaultFeeReceiver)
func (_NSwapExchange *NSwapExchangeFilterer) FilterEventUpdatedDefaultProtocolFeeReceiver(opts *bind.FilterOpts, defaultFeeReceiver []common.Address) (*NSwapExchangeEventUpdatedDefaultProtocolFeeReceiverIterator, error) {

	var defaultFeeReceiverRule []interface{}
	for _, defaultFeeReceiverItem := range defaultFeeReceiver {
		defaultFeeReceiverRule = append(defaultFeeReceiverRule, defaultFeeReceiverItem)
	}

	logs, sub, err := _NSwapExchange.contract.FilterLogs(opts, "EventUpdatedDefaultProtocolFeeReceiver", defaultFeeReceiverRule)
	if err != nil {
		return nil, err
	}
	return &NSwapExchangeEventUpdatedDefaultProtocolFeeReceiverIterator{contract: _NSwapExchange.contract, event: "EventUpdatedDefaultProtocolFeeReceiver", logs: logs, sub: sub}, nil
}

// WatchEventUpdatedDefaultProtocolFeeReceiver is a free log subscription operation binding the contract event 0xf71cba06d4627dfe19c434c00bbe1d60d48bdfcfc27781ba977953777e7ee2ee.
//
// Solidity: event EventUpdatedDefaultProtocolFeeReceiver(address indexed defaultFeeReceiver)
func (_NSwapExchange *NSwapExchangeFilterer) WatchEventUpdatedDefaultProtocolFeeReceiver(opts *bind.WatchOpts, sink chan<- *NSwapExchangeEventUpdatedDefaultProtocolFeeReceiver, defaultFeeReceiver []common.Address) (event.Subscription, error) {

	var defaultFeeReceiverRule []interface{}
	for _, defaultFeeReceiverItem := range defaultFeeReceiver {
		defaultFeeReceiverRule = append(defaultFeeReceiverRule, defaultFeeReceiverItem)
	}

	logs, sub, err := _NSwapExchange.contract.WatchLogs(opts, "EventUpdatedDefaultProtocolFeeReceiver", defaultFeeReceiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NSwapExchangeEventUpdatedDefaultProtocolFeeReceiver)
				if err := _NSwapExchange.contract.UnpackLog(event, "EventUpdatedDefaultProtocolFeeReceiver", log); err != nil {
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

// ParseEventUpdatedDefaultProtocolFeeReceiver is a log parse operation binding the contract event 0xf71cba06d4627dfe19c434c00bbe1d60d48bdfcfc27781ba977953777e7ee2ee.
//
// Solidity: event EventUpdatedDefaultProtocolFeeReceiver(address indexed defaultFeeReceiver)
func (_NSwapExchange *NSwapExchangeFilterer) ParseEventUpdatedDefaultProtocolFeeReceiver(log types.Log) (*NSwapExchangeEventUpdatedDefaultProtocolFeeReceiver, error) {
	event := new(NSwapExchangeEventUpdatedDefaultProtocolFeeReceiver)
	if err := _NSwapExchange.contract.UnpackLog(event, "EventUpdatedDefaultProtocolFeeReceiver", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NSwapExchangeEventUpdatedRoyaltiesByTokenIterator is returned from FilterEventUpdatedRoyaltiesByToken and is used to iterate over the raw logs and unpacked data for EventUpdatedRoyaltiesByToken events raised by the NSwapExchange contract.
type NSwapExchangeEventUpdatedRoyaltiesByTokenIterator struct {
	Event *NSwapExchangeEventUpdatedRoyaltiesByToken // Event containing the contract specifics and raw log

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
func (it *NSwapExchangeEventUpdatedRoyaltiesByTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NSwapExchangeEventUpdatedRoyaltiesByToken)
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
		it.Event = new(NSwapExchangeEventUpdatedRoyaltiesByToken)
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
func (it *NSwapExchangeEventUpdatedRoyaltiesByTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NSwapExchangeEventUpdatedRoyaltiesByTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NSwapExchangeEventUpdatedRoyaltiesByToken represents a EventUpdatedRoyaltiesByToken event raised by the NSwapExchange contract.
type NSwapExchangeEventUpdatedRoyaltiesByToken struct {
	Token     common.Address
	Royalties []LibPayInfoPayInfo
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEventUpdatedRoyaltiesByToken is a free log retrieval operation binding the contract event 0x1052e0ab08ca295e8a29bcba484c82c15bd41df0bd193f13a7752b296da38db6.
//
// Solidity: event EventUpdatedRoyaltiesByToken(address indexed token, (address,uint96)[] royalties)
func (_NSwapExchange *NSwapExchangeFilterer) FilterEventUpdatedRoyaltiesByToken(opts *bind.FilterOpts, token []common.Address) (*NSwapExchangeEventUpdatedRoyaltiesByTokenIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _NSwapExchange.contract.FilterLogs(opts, "EventUpdatedRoyaltiesByToken", tokenRule)
	if err != nil {
		return nil, err
	}
	return &NSwapExchangeEventUpdatedRoyaltiesByTokenIterator{contract: _NSwapExchange.contract, event: "EventUpdatedRoyaltiesByToken", logs: logs, sub: sub}, nil
}

// WatchEventUpdatedRoyaltiesByToken is a free log subscription operation binding the contract event 0x1052e0ab08ca295e8a29bcba484c82c15bd41df0bd193f13a7752b296da38db6.
//
// Solidity: event EventUpdatedRoyaltiesByToken(address indexed token, (address,uint96)[] royalties)
func (_NSwapExchange *NSwapExchangeFilterer) WatchEventUpdatedRoyaltiesByToken(opts *bind.WatchOpts, sink chan<- *NSwapExchangeEventUpdatedRoyaltiesByToken, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _NSwapExchange.contract.WatchLogs(opts, "EventUpdatedRoyaltiesByToken", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NSwapExchangeEventUpdatedRoyaltiesByToken)
				if err := _NSwapExchange.contract.UnpackLog(event, "EventUpdatedRoyaltiesByToken", log); err != nil {
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

// ParseEventUpdatedRoyaltiesByToken is a log parse operation binding the contract event 0x1052e0ab08ca295e8a29bcba484c82c15bd41df0bd193f13a7752b296da38db6.
//
// Solidity: event EventUpdatedRoyaltiesByToken(address indexed token, (address,uint96)[] royalties)
func (_NSwapExchange *NSwapExchangeFilterer) ParseEventUpdatedRoyaltiesByToken(log types.Log) (*NSwapExchangeEventUpdatedRoyaltiesByToken, error) {
	event := new(NSwapExchangeEventUpdatedRoyaltiesByToken)
	if err := _NSwapExchange.contract.UnpackLog(event, "EventUpdatedRoyaltiesByToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NSwapExchangeEventUpdatedRoyaltiesLimitIterator is returned from FilterEventUpdatedRoyaltiesLimit and is used to iterate over the raw logs and unpacked data for EventUpdatedRoyaltiesLimit events raised by the NSwapExchange contract.
type NSwapExchangeEventUpdatedRoyaltiesLimitIterator struct {
	Event *NSwapExchangeEventUpdatedRoyaltiesLimit // Event containing the contract specifics and raw log

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
func (it *NSwapExchangeEventUpdatedRoyaltiesLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NSwapExchangeEventUpdatedRoyaltiesLimit)
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
		it.Event = new(NSwapExchangeEventUpdatedRoyaltiesLimit)
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
func (it *NSwapExchangeEventUpdatedRoyaltiesLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NSwapExchangeEventUpdatedRoyaltiesLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NSwapExchangeEventUpdatedRoyaltiesLimit represents a EventUpdatedRoyaltiesLimit event raised by the NSwapExchange contract.
type NSwapExchangeEventUpdatedRoyaltiesLimit struct {
	NewRoyaltiesLimit *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterEventUpdatedRoyaltiesLimit is a free log retrieval operation binding the contract event 0xb2c6d570e610ecd8ab1c03d181469e759af504adbf07b3b6b242828f63c55386.
//
// Solidity: event EventUpdatedRoyaltiesLimit(uint256 newRoyaltiesLimit)
func (_NSwapExchange *NSwapExchangeFilterer) FilterEventUpdatedRoyaltiesLimit(opts *bind.FilterOpts) (*NSwapExchangeEventUpdatedRoyaltiesLimitIterator, error) {

	logs, sub, err := _NSwapExchange.contract.FilterLogs(opts, "EventUpdatedRoyaltiesLimit")
	if err != nil {
		return nil, err
	}
	return &NSwapExchangeEventUpdatedRoyaltiesLimitIterator{contract: _NSwapExchange.contract, event: "EventUpdatedRoyaltiesLimit", logs: logs, sub: sub}, nil
}

// WatchEventUpdatedRoyaltiesLimit is a free log subscription operation binding the contract event 0xb2c6d570e610ecd8ab1c03d181469e759af504adbf07b3b6b242828f63c55386.
//
// Solidity: event EventUpdatedRoyaltiesLimit(uint256 newRoyaltiesLimit)
func (_NSwapExchange *NSwapExchangeFilterer) WatchEventUpdatedRoyaltiesLimit(opts *bind.WatchOpts, sink chan<- *NSwapExchangeEventUpdatedRoyaltiesLimit) (event.Subscription, error) {

	logs, sub, err := _NSwapExchange.contract.WatchLogs(opts, "EventUpdatedRoyaltiesLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NSwapExchangeEventUpdatedRoyaltiesLimit)
				if err := _NSwapExchange.contract.UnpackLog(event, "EventUpdatedRoyaltiesLimit", log); err != nil {
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

// ParseEventUpdatedRoyaltiesLimit is a log parse operation binding the contract event 0xb2c6d570e610ecd8ab1c03d181469e759af504adbf07b3b6b242828f63c55386.
//
// Solidity: event EventUpdatedRoyaltiesLimit(uint256 newRoyaltiesLimit)
func (_NSwapExchange *NSwapExchangeFilterer) ParseEventUpdatedRoyaltiesLimit(log types.Log) (*NSwapExchangeEventUpdatedRoyaltiesLimit, error) {
	event := new(NSwapExchangeEventUpdatedRoyaltiesLimit)
	if err := _NSwapExchange.contract.UnpackLog(event, "EventUpdatedRoyaltiesLimit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NSwapExchangeEventWithdrawETHIterator is returned from FilterEventWithdrawETH and is used to iterate over the raw logs and unpacked data for EventWithdrawETH events raised by the NSwapExchange contract.
type NSwapExchangeEventWithdrawETHIterator struct {
	Event *NSwapExchangeEventWithdrawETH // Event containing the contract specifics and raw log

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
func (it *NSwapExchangeEventWithdrawETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NSwapExchangeEventWithdrawETH)
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
		it.Event = new(NSwapExchangeEventWithdrawETH)
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
func (it *NSwapExchangeEventWithdrawETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NSwapExchangeEventWithdrawETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NSwapExchangeEventWithdrawETH represents a EventWithdrawETH event raised by the NSwapExchange contract.
type NSwapExchangeEventWithdrawETH struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEventWithdrawETH is a free log retrieval operation binding the contract event 0xce4bf35cf3bcfb887d48c146fb84feacfc51209bb32d40792099ce369b38e7ff.
//
// Solidity: event EventWithdrawETH(address recipient, uint256 amount)
func (_NSwapExchange *NSwapExchangeFilterer) FilterEventWithdrawETH(opts *bind.FilterOpts) (*NSwapExchangeEventWithdrawETHIterator, error) {

	logs, sub, err := _NSwapExchange.contract.FilterLogs(opts, "EventWithdrawETH")
	if err != nil {
		return nil, err
	}
	return &NSwapExchangeEventWithdrawETHIterator{contract: _NSwapExchange.contract, event: "EventWithdrawETH", logs: logs, sub: sub}, nil
}

// WatchEventWithdrawETH is a free log subscription operation binding the contract event 0xce4bf35cf3bcfb887d48c146fb84feacfc51209bb32d40792099ce369b38e7ff.
//
// Solidity: event EventWithdrawETH(address recipient, uint256 amount)
func (_NSwapExchange *NSwapExchangeFilterer) WatchEventWithdrawETH(opts *bind.WatchOpts, sink chan<- *NSwapExchangeEventWithdrawETH) (event.Subscription, error) {

	logs, sub, err := _NSwapExchange.contract.WatchLogs(opts, "EventWithdrawETH")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NSwapExchangeEventWithdrawETH)
				if err := _NSwapExchange.contract.UnpackLog(event, "EventWithdrawETH", log); err != nil {
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

// ParseEventWithdrawETH is a log parse operation binding the contract event 0xce4bf35cf3bcfb887d48c146fb84feacfc51209bb32d40792099ce369b38e7ff.
//
// Solidity: event EventWithdrawETH(address recipient, uint256 amount)
func (_NSwapExchange *NSwapExchangeFilterer) ParseEventWithdrawETH(log types.Log) (*NSwapExchangeEventWithdrawETH, error) {
	event := new(NSwapExchangeEventWithdrawETH)
	if err := _NSwapExchange.contract.UnpackLog(event, "EventWithdrawETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NSwapExchangeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the NSwapExchange contract.
type NSwapExchangeInitializedIterator struct {
	Event *NSwapExchangeInitialized // Event containing the contract specifics and raw log

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
func (it *NSwapExchangeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NSwapExchangeInitialized)
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
		it.Event = new(NSwapExchangeInitialized)
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
func (it *NSwapExchangeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NSwapExchangeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NSwapExchangeInitialized represents a Initialized event raised by the NSwapExchange contract.
type NSwapExchangeInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_NSwapExchange *NSwapExchangeFilterer) FilterInitialized(opts *bind.FilterOpts) (*NSwapExchangeInitializedIterator, error) {

	logs, sub, err := _NSwapExchange.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &NSwapExchangeInitializedIterator{contract: _NSwapExchange.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_NSwapExchange *NSwapExchangeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *NSwapExchangeInitialized) (event.Subscription, error) {

	logs, sub, err := _NSwapExchange.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NSwapExchangeInitialized)
				if err := _NSwapExchange.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_NSwapExchange *NSwapExchangeFilterer) ParseInitialized(log types.Log) (*NSwapExchangeInitialized, error) {
	event := new(NSwapExchangeInitialized)
	if err := _NSwapExchange.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NSwapExchangeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the NSwapExchange contract.
type NSwapExchangeOwnershipTransferredIterator struct {
	Event *NSwapExchangeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *NSwapExchangeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NSwapExchangeOwnershipTransferred)
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
		it.Event = new(NSwapExchangeOwnershipTransferred)
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
func (it *NSwapExchangeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NSwapExchangeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NSwapExchangeOwnershipTransferred represents a OwnershipTransferred event raised by the NSwapExchange contract.
type NSwapExchangeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NSwapExchange *NSwapExchangeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NSwapExchangeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NSwapExchange.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NSwapExchangeOwnershipTransferredIterator{contract: _NSwapExchange.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NSwapExchange *NSwapExchangeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NSwapExchangeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NSwapExchange.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NSwapExchangeOwnershipTransferred)
				if err := _NSwapExchange.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_NSwapExchange *NSwapExchangeFilterer) ParseOwnershipTransferred(log types.Log) (*NSwapExchangeOwnershipTransferred, error) {
	event := new(NSwapExchangeOwnershipTransferred)
	if err := _NSwapExchange.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NSwapExchangePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the NSwapExchange contract.
type NSwapExchangePausedIterator struct {
	Event *NSwapExchangePaused // Event containing the contract specifics and raw log

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
func (it *NSwapExchangePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NSwapExchangePaused)
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
		it.Event = new(NSwapExchangePaused)
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
func (it *NSwapExchangePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NSwapExchangePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NSwapExchangePaused represents a Paused event raised by the NSwapExchange contract.
type NSwapExchangePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_NSwapExchange *NSwapExchangeFilterer) FilterPaused(opts *bind.FilterOpts) (*NSwapExchangePausedIterator, error) {

	logs, sub, err := _NSwapExchange.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &NSwapExchangePausedIterator{contract: _NSwapExchange.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_NSwapExchange *NSwapExchangeFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *NSwapExchangePaused) (event.Subscription, error) {

	logs, sub, err := _NSwapExchange.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NSwapExchangePaused)
				if err := _NSwapExchange.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_NSwapExchange *NSwapExchangeFilterer) ParsePaused(log types.Log) (*NSwapExchangePaused, error) {
	event := new(NSwapExchangePaused)
	if err := _NSwapExchange.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NSwapExchangeUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the NSwapExchange contract.
type NSwapExchangeUnpausedIterator struct {
	Event *NSwapExchangeUnpaused // Event containing the contract specifics and raw log

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
func (it *NSwapExchangeUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NSwapExchangeUnpaused)
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
		it.Event = new(NSwapExchangeUnpaused)
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
func (it *NSwapExchangeUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NSwapExchangeUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NSwapExchangeUnpaused represents a Unpaused event raised by the NSwapExchange contract.
type NSwapExchangeUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_NSwapExchange *NSwapExchangeFilterer) FilterUnpaused(opts *bind.FilterOpts) (*NSwapExchangeUnpausedIterator, error) {

	logs, sub, err := _NSwapExchange.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &NSwapExchangeUnpausedIterator{contract: _NSwapExchange.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_NSwapExchange *NSwapExchangeFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *NSwapExchangeUnpaused) (event.Subscription, error) {

	logs, sub, err := _NSwapExchange.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NSwapExchangeUnpaused)
				if err := _NSwapExchange.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_NSwapExchange *NSwapExchangeFilterer) ParseUnpaused(log types.Log) (*NSwapExchangeUnpaused, error) {
	event := new(NSwapExchangeUnpaused)
	if err := _NSwapExchange.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
