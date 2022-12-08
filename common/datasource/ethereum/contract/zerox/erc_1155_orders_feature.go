// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zerox

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

// LibNFTOrderERC1155Order is an auto generated low-level Go binding around an user-defined struct.
type LibNFTOrderERC1155Order struct {
	Direction              uint8
	Maker                  common.Address
	Taker                  common.Address
	Expiry                 *big.Int
	Nonce                  *big.Int
	Erc20Token             common.Address
	Erc20TokenAmount       *big.Int
	Fees                   []LibNFTOrderFee
	Erc1155Token           common.Address
	Erc1155TokenId         *big.Int
	Erc1155TokenProperties []LibNFTOrderProperty
	Erc1155TokenAmount     *big.Int
}

// LibNFTOrderOrderInfo is an auto generated low-level Go binding around an user-defined struct.
type LibNFTOrderOrderInfo struct {
	OrderHash       [32]byte
	Status          uint8
	OrderAmount     *big.Int
	RemainingAmount *big.Int
}

// NFTOrdersBuyParams is an auto generated low-level Go binding around an user-defined struct.
type NFTOrdersBuyParams struct {
	BuyAmount         *big.Int
	EthAvailable      *big.Int
	TakerCallbackData []byte
}

// ERC1155OrdersFeatureMetaData contains all meta data concerning the ERC1155OrdersFeature contract.
var ERC1155OrdersFeatureMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"zeroExAddress\",\"type\":\"address\"},{\"internalType\":\"contractIEtherTokenV06\",\"name\":\"weth\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"ERC1155OrderCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumLibNFTOrder.TradeDirection\",\"name\":\"direction\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIERC20TokenV06\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"erc20FillAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIERC1155Token\",\"name\":\"erc1155Token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"erc1155TokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"erc1155FillAmount\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"matcher\",\"type\":\"address\"}],\"name\":\"ERC1155OrderFilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumLibNFTOrder.TradeDirection\",\"name\":\"direction\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIERC20TokenV06\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"erc20TokenAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structLibNFTOrder.Fee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"contractIERC1155Token\",\"name\":\"erc1155Token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"erc1155TokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractIPropertyValidator\",\"name\":\"propertyValidator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"propertyData\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structLibNFTOrder.Property[]\",\"name\":\"erc1155TokenProperties\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"erc1155TokenAmount\",\"type\":\"uint128\"}],\"name\":\"ERC1155OrderPreSigned\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"EIP712_DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FEATURE_NAME\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FEATURE_VERSION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumLibNFTOrder.TradeDirection\",\"name\":\"direction\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20TokenV06\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc20TokenAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Fee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"contractIERC1155Token\",\"name\":\"erc1155Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc1155TokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractIPropertyValidator\",\"name\":\"propertyValidator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"propertyData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Property[]\",\"name\":\"erc1155TokenProperties\",\"type\":\"tuple[]\"},{\"internalType\":\"uint128\",\"name\":\"erc1155TokenAmount\",\"type\":\"uint128\"}],\"internalType\":\"structLibNFTOrder.ERC1155Order\",\"name\":\"sellOrder\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumLibSignature.SignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structLibSignature.Signature\",\"name\":\"signature\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"buyAmount\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"ethAvailable\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"takerCallbackData\",\"type\":\"bytes\"}],\"internalType\":\"structNFTOrders.BuyParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"_buyERC1155\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumLibNFTOrder.TradeDirection\",\"name\":\"direction\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20TokenV06\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc20TokenAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Fee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"contractIERC1155Token\",\"name\":\"erc1155Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc1155TokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractIPropertyValidator\",\"name\":\"propertyValidator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"propertyData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Property[]\",\"name\":\"erc1155TokenProperties\",\"type\":\"tuple[]\"},{\"internalType\":\"uint128\",\"name\":\"erc1155TokenAmount\",\"type\":\"uint128\"}],\"internalType\":\"structLibNFTOrder.ERC1155Order[]\",\"name\":\"sellOrders\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumLibSignature.SignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structLibSignature.Signature[]\",\"name\":\"signatures\",\"type\":\"tuple[]\"},{\"internalType\":\"uint128[]\",\"name\":\"erc1155FillAmounts\",\"type\":\"uint128[]\"},{\"internalType\":\"bytes[]\",\"name\":\"callbackData\",\"type\":\"bytes[]\"},{\"internalType\":\"bool\",\"name\":\"revertIfIncomplete\",\"type\":\"bool\"}],\"name\":\"batchBuyERC1155s\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"successes\",\"type\":\"bool[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"orderNonces\",\"type\":\"uint256[]\"}],\"name\":\"batchCancelERC1155Orders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumLibNFTOrder.TradeDirection\",\"name\":\"direction\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20TokenV06\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc20TokenAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Fee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"contractIERC1155Token\",\"name\":\"erc1155Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc1155TokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractIPropertyValidator\",\"name\":\"propertyValidator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"propertyData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Property[]\",\"name\":\"erc1155TokenProperties\",\"type\":\"tuple[]\"},{\"internalType\":\"uint128\",\"name\":\"erc1155TokenAmount\",\"type\":\"uint128\"}],\"internalType\":\"structLibNFTOrder.ERC1155Order\",\"name\":\"sellOrder\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumLibSignature.SignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structLibSignature.Signature\",\"name\":\"signature\",\"type\":\"tuple\"},{\"internalType\":\"uint128\",\"name\":\"erc1155BuyAmount\",\"type\":\"uint128\"},{\"internalType\":\"bytes\",\"name\":\"callbackData\",\"type\":\"bytes\"}],\"name\":\"buyERC1155\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderNonce\",\"type\":\"uint256\"}],\"name\":\"cancelERC1155Order\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumLibNFTOrder.TradeDirection\",\"name\":\"direction\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20TokenV06\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc20TokenAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Fee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"contractIERC1155Token\",\"name\":\"erc1155Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc1155TokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractIPropertyValidator\",\"name\":\"propertyValidator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"propertyData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Property[]\",\"name\":\"erc1155TokenProperties\",\"type\":\"tuple[]\"},{\"internalType\":\"uint128\",\"name\":\"erc1155TokenAmount\",\"type\":\"uint128\"}],\"internalType\":\"structLibNFTOrder.ERC1155Order\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"getERC1155OrderHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumLibNFTOrder.TradeDirection\",\"name\":\"direction\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20TokenV06\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc20TokenAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Fee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"contractIERC1155Token\",\"name\":\"erc1155Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc1155TokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractIPropertyValidator\",\"name\":\"propertyValidator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"propertyData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Property[]\",\"name\":\"erc1155TokenProperties\",\"type\":\"tuple[]\"},{\"internalType\":\"uint128\",\"name\":\"erc1155TokenAmount\",\"type\":\"uint128\"}],\"internalType\":\"structLibNFTOrder.ERC1155Order\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"getERC1155OrderInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"internalType\":\"enumLibNFTOrder.OrderStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint128\",\"name\":\"orderAmount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"remainingAmount\",\"type\":\"uint128\"}],\"internalType\":\"structLibNFTOrder.OrderInfo\",\"name\":\"orderInfo\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"migrate\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"success\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"success\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumLibNFTOrder.TradeDirection\",\"name\":\"direction\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20TokenV06\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc20TokenAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Fee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"contractIERC1155Token\",\"name\":\"erc1155Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc1155TokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractIPropertyValidator\",\"name\":\"propertyValidator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"propertyData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Property[]\",\"name\":\"erc1155TokenProperties\",\"type\":\"tuple[]\"},{\"internalType\":\"uint128\",\"name\":\"erc1155TokenAmount\",\"type\":\"uint128\"}],\"internalType\":\"structLibNFTOrder.ERC1155Order\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"preSignERC1155Order\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumLibNFTOrder.TradeDirection\",\"name\":\"direction\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20TokenV06\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc20TokenAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Fee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"contractIERC1155Token\",\"name\":\"erc1155Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc1155TokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractIPropertyValidator\",\"name\":\"propertyValidator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"propertyData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Property[]\",\"name\":\"erc1155TokenProperties\",\"type\":\"tuple[]\"},{\"internalType\":\"uint128\",\"name\":\"erc1155TokenAmount\",\"type\":\"uint128\"}],\"internalType\":\"structLibNFTOrder.ERC1155Order\",\"name\":\"buyOrder\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumLibSignature.SignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structLibSignature.Signature\",\"name\":\"signature\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"erc1155TokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"erc1155SellAmount\",\"type\":\"uint128\"},{\"internalType\":\"bool\",\"name\":\"unwrapNativeToken\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"callbackData\",\"type\":\"bytes\"}],\"name\":\"sellERC1155\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumLibNFTOrder.TradeDirection\",\"name\":\"direction\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20TokenV06\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc20TokenAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Fee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"contractIERC1155Token\",\"name\":\"erc1155Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc1155TokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractIPropertyValidator\",\"name\":\"propertyValidator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"propertyData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Property[]\",\"name\":\"erc1155TokenProperties\",\"type\":\"tuple[]\"},{\"internalType\":\"uint128\",\"name\":\"erc1155TokenAmount\",\"type\":\"uint128\"}],\"internalType\":\"structLibNFTOrder.ERC1155Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"erc1155TokenId\",\"type\":\"uint256\"}],\"name\":\"validateERC1155OrderProperties\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumLibNFTOrder.TradeDirection\",\"name\":\"direction\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20TokenV06\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc20TokenAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Fee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"contractIERC1155Token\",\"name\":\"erc1155Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc1155TokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractIPropertyValidator\",\"name\":\"propertyValidator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"propertyData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Property[]\",\"name\":\"erc1155TokenProperties\",\"type\":\"tuple[]\"},{\"internalType\":\"uint128\",\"name\":\"erc1155TokenAmount\",\"type\":\"uint128\"}],\"internalType\":\"structLibNFTOrder.ERC1155Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumLibSignature.SignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structLibSignature.Signature\",\"name\":\"signature\",\"type\":\"tuple\"}],\"name\":\"validateERC1155OrderSignature\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ERC1155OrdersFeatureABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC1155OrdersFeatureMetaData.ABI instead.
var ERC1155OrdersFeatureABI = ERC1155OrdersFeatureMetaData.ABI

// ERC1155OrdersFeature is an auto generated Go binding around an Ethereum contract.
type ERC1155OrdersFeature struct {
	ERC1155OrdersFeatureCaller     // Read-only binding to the contract
	ERC1155OrdersFeatureTransactor // Write-only binding to the contract
	ERC1155OrdersFeatureFilterer   // Log filterer for contract events
}

// ERC1155OrdersFeatureCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC1155OrdersFeatureCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1155OrdersFeatureTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC1155OrdersFeatureTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1155OrdersFeatureFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC1155OrdersFeatureFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1155OrdersFeatureSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC1155OrdersFeatureSession struct {
	Contract     *ERC1155OrdersFeature // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ERC1155OrdersFeatureCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC1155OrdersFeatureCallerSession struct {
	Contract *ERC1155OrdersFeatureCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// ERC1155OrdersFeatureTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC1155OrdersFeatureTransactorSession struct {
	Contract     *ERC1155OrdersFeatureTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// ERC1155OrdersFeatureRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC1155OrdersFeatureRaw struct {
	Contract *ERC1155OrdersFeature // Generic contract binding to access the raw methods on
}

// ERC1155OrdersFeatureCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC1155OrdersFeatureCallerRaw struct {
	Contract *ERC1155OrdersFeatureCaller // Generic read-only contract binding to access the raw methods on
}

// ERC1155OrdersFeatureTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC1155OrdersFeatureTransactorRaw struct {
	Contract *ERC1155OrdersFeatureTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC1155OrdersFeature creates a new instance of ERC1155OrdersFeature, bound to a specific deployed contract.
func NewERC1155OrdersFeature(address common.Address, backend bind.ContractBackend) (*ERC1155OrdersFeature, error) {
	contract, err := bindERC1155OrdersFeature(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC1155OrdersFeature{ERC1155OrdersFeatureCaller: ERC1155OrdersFeatureCaller{contract: contract}, ERC1155OrdersFeatureTransactor: ERC1155OrdersFeatureTransactor{contract: contract}, ERC1155OrdersFeatureFilterer: ERC1155OrdersFeatureFilterer{contract: contract}}, nil
}

// NewERC1155OrdersFeatureCaller creates a new read-only instance of ERC1155OrdersFeature, bound to a specific deployed contract.
func NewERC1155OrdersFeatureCaller(address common.Address, caller bind.ContractCaller) (*ERC1155OrdersFeatureCaller, error) {
	contract, err := bindERC1155OrdersFeature(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC1155OrdersFeatureCaller{contract: contract}, nil
}

// NewERC1155OrdersFeatureTransactor creates a new write-only instance of ERC1155OrdersFeature, bound to a specific deployed contract.
func NewERC1155OrdersFeatureTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC1155OrdersFeatureTransactor, error) {
	contract, err := bindERC1155OrdersFeature(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC1155OrdersFeatureTransactor{contract: contract}, nil
}

// NewERC1155OrdersFeatureFilterer creates a new log filterer instance of ERC1155OrdersFeature, bound to a specific deployed contract.
func NewERC1155OrdersFeatureFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC1155OrdersFeatureFilterer, error) {
	contract, err := bindERC1155OrdersFeature(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC1155OrdersFeatureFilterer{contract: contract}, nil
}

// bindERC1155OrdersFeature binds a generic wrapper to an already deployed contract.
func bindERC1155OrdersFeature(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC1155OrdersFeatureABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC1155OrdersFeature.Contract.ERC1155OrdersFeatureCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC1155OrdersFeature.Contract.ERC1155OrdersFeatureTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC1155OrdersFeature.Contract.ERC1155OrdersFeatureTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC1155OrdersFeature.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC1155OrdersFeature.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC1155OrdersFeature.Contract.contract.Transact(opts, method, params...)
}

// EIP712DOMAINSEPARATOR is a free data retrieval call binding the contract method 0xdab400f3.
//
// Solidity: function EIP712_DOMAIN_SEPARATOR() view returns(bytes32)
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureCaller) EIP712DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC1155OrdersFeature.contract.Call(opts, &out, "EIP712_DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EIP712DOMAINSEPARATOR is a free data retrieval call binding the contract method 0xdab400f3.
//
// Solidity: function EIP712_DOMAIN_SEPARATOR() view returns(bytes32)
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureSession) EIP712DOMAINSEPARATOR() ([32]byte, error) {
	return _ERC1155OrdersFeature.Contract.EIP712DOMAINSEPARATOR(&_ERC1155OrdersFeature.CallOpts)
}

// EIP712DOMAINSEPARATOR is a free data retrieval call binding the contract method 0xdab400f3.
//
// Solidity: function EIP712_DOMAIN_SEPARATOR() view returns(bytes32)
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureCallerSession) EIP712DOMAINSEPARATOR() ([32]byte, error) {
	return _ERC1155OrdersFeature.Contract.EIP712DOMAINSEPARATOR(&_ERC1155OrdersFeature.CallOpts)
}

// FEATURENAME is a free data retrieval call binding the contract method 0x6ae4b4f7.
//
// Solidity: function FEATURE_NAME() view returns(string)
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureCaller) FEATURENAME(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC1155OrdersFeature.contract.Call(opts, &out, "FEATURE_NAME")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// FEATURENAME is a free data retrieval call binding the contract method 0x6ae4b4f7.
//
// Solidity: function FEATURE_NAME() view returns(string)
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureSession) FEATURENAME() (string, error) {
	return _ERC1155OrdersFeature.Contract.FEATURENAME(&_ERC1155OrdersFeature.CallOpts)
}

// FEATURENAME is a free data retrieval call binding the contract method 0x6ae4b4f7.
//
// Solidity: function FEATURE_NAME() view returns(string)
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureCallerSession) FEATURENAME() (string, error) {
	return _ERC1155OrdersFeature.Contract.FEATURENAME(&_ERC1155OrdersFeature.CallOpts)
}

// FEATUREVERSION is a free data retrieval call binding the contract method 0x031b905c.
//
// Solidity: function FEATURE_VERSION() view returns(uint256)
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureCaller) FEATUREVERSION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC1155OrdersFeature.contract.Call(opts, &out, "FEATURE_VERSION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FEATUREVERSION is a free data retrieval call binding the contract method 0x031b905c.
//
// Solidity: function FEATURE_VERSION() view returns(uint256)
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureSession) FEATUREVERSION() (*big.Int, error) {
	return _ERC1155OrdersFeature.Contract.FEATUREVERSION(&_ERC1155OrdersFeature.CallOpts)
}

// FEATUREVERSION is a free data retrieval call binding the contract method 0x031b905c.
//
// Solidity: function FEATURE_VERSION() view returns(uint256)
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureCallerSession) FEATUREVERSION() (*big.Int, error) {
	return _ERC1155OrdersFeature.Contract.FEATUREVERSION(&_ERC1155OrdersFeature.CallOpts)
}

// GetERC1155OrderHash is a free data retrieval call binding the contract method 0x1de3a7ac.
//
// Solidity: function getERC1155OrderHash((uint8,address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[],uint128) order) view returns(bytes32 orderHash)
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureCaller) GetERC1155OrderHash(opts *bind.CallOpts, order LibNFTOrderERC1155Order) ([32]byte, error) {
	var out []interface{}
	err := _ERC1155OrdersFeature.contract.Call(opts, &out, "getERC1155OrderHash", order)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetERC1155OrderHash is a free data retrieval call binding the contract method 0x1de3a7ac.
//
// Solidity: function getERC1155OrderHash((uint8,address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[],uint128) order) view returns(bytes32 orderHash)
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureSession) GetERC1155OrderHash(order LibNFTOrderERC1155Order) ([32]byte, error) {
	return _ERC1155OrdersFeature.Contract.GetERC1155OrderHash(&_ERC1155OrdersFeature.CallOpts, order)
}

// GetERC1155OrderHash is a free data retrieval call binding the contract method 0x1de3a7ac.
//
// Solidity: function getERC1155OrderHash((uint8,address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[],uint128) order) view returns(bytes32 orderHash)
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureCallerSession) GetERC1155OrderHash(order LibNFTOrderERC1155Order) ([32]byte, error) {
	return _ERC1155OrdersFeature.Contract.GetERC1155OrderHash(&_ERC1155OrdersFeature.CallOpts, order)
}

// GetERC1155OrderInfo is a free data retrieval call binding the contract method 0x4991fd72.
//
// Solidity: function getERC1155OrderInfo((uint8,address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[],uint128) order) view returns((bytes32,uint8,uint128,uint128) orderInfo)
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureCaller) GetERC1155OrderInfo(opts *bind.CallOpts, order LibNFTOrderERC1155Order) (LibNFTOrderOrderInfo, error) {
	var out []interface{}
	err := _ERC1155OrdersFeature.contract.Call(opts, &out, "getERC1155OrderInfo", order)

	if err != nil {
		return *new(LibNFTOrderOrderInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(LibNFTOrderOrderInfo)).(*LibNFTOrderOrderInfo)

	return out0, err

}

// GetERC1155OrderInfo is a free data retrieval call binding the contract method 0x4991fd72.
//
// Solidity: function getERC1155OrderInfo((uint8,address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[],uint128) order) view returns((bytes32,uint8,uint128,uint128) orderInfo)
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureSession) GetERC1155OrderInfo(order LibNFTOrderERC1155Order) (LibNFTOrderOrderInfo, error) {
	return _ERC1155OrdersFeature.Contract.GetERC1155OrderInfo(&_ERC1155OrdersFeature.CallOpts, order)
}

// GetERC1155OrderInfo is a free data retrieval call binding the contract method 0x4991fd72.
//
// Solidity: function getERC1155OrderInfo((uint8,address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[],uint128) order) view returns((bytes32,uint8,uint128,uint128) orderInfo)
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureCallerSession) GetERC1155OrderInfo(order LibNFTOrderERC1155Order) (LibNFTOrderOrderInfo, error) {
	return _ERC1155OrdersFeature.Contract.GetERC1155OrderInfo(&_ERC1155OrdersFeature.CallOpts, order)
}

// ValidateERC1155OrderProperties is a free data retrieval call binding the contract method 0x2ac6f62a.
//
// Solidity: function validateERC1155OrderProperties((uint8,address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[],uint128) order, uint256 erc1155TokenId) view returns()
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureCaller) ValidateERC1155OrderProperties(opts *bind.CallOpts, order LibNFTOrderERC1155Order, erc1155TokenId *big.Int) error {
	var out []interface{}
	err := _ERC1155OrdersFeature.contract.Call(opts, &out, "validateERC1155OrderProperties", order, erc1155TokenId)

	if err != nil {
		return err
	}

	return err

}

// ValidateERC1155OrderProperties is a free data retrieval call binding the contract method 0x2ac6f62a.
//
// Solidity: function validateERC1155OrderProperties((uint8,address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[],uint128) order, uint256 erc1155TokenId) view returns()
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureSession) ValidateERC1155OrderProperties(order LibNFTOrderERC1155Order, erc1155TokenId *big.Int) error {
	return _ERC1155OrdersFeature.Contract.ValidateERC1155OrderProperties(&_ERC1155OrdersFeature.CallOpts, order, erc1155TokenId)
}

// ValidateERC1155OrderProperties is a free data retrieval call binding the contract method 0x2ac6f62a.
//
// Solidity: function validateERC1155OrderProperties((uint8,address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[],uint128) order, uint256 erc1155TokenId) view returns()
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureCallerSession) ValidateERC1155OrderProperties(order LibNFTOrderERC1155Order, erc1155TokenId *big.Int) error {
	return _ERC1155OrdersFeature.Contract.ValidateERC1155OrderProperties(&_ERC1155OrdersFeature.CallOpts, order, erc1155TokenId)
}

// ValidateERC1155OrderSignature is a free data retrieval call binding the contract method 0x0d32a531.
//
// Solidity: function validateERC1155OrderSignature((uint8,address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[],uint128) order, (uint8,uint8,bytes32,bytes32) signature) view returns()
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureCaller) ValidateERC1155OrderSignature(opts *bind.CallOpts, order LibNFTOrderERC1155Order, signature LibSignatureSignature) error {
	var out []interface{}
	err := _ERC1155OrdersFeature.contract.Call(opts, &out, "validateERC1155OrderSignature", order, signature)

	if err != nil {
		return err
	}

	return err

}

// ValidateERC1155OrderSignature is a free data retrieval call binding the contract method 0x0d32a531.
//
// Solidity: function validateERC1155OrderSignature((uint8,address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[],uint128) order, (uint8,uint8,bytes32,bytes32) signature) view returns()
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureSession) ValidateERC1155OrderSignature(order LibNFTOrderERC1155Order, signature LibSignatureSignature) error {
	return _ERC1155OrdersFeature.Contract.ValidateERC1155OrderSignature(&_ERC1155OrdersFeature.CallOpts, order, signature)
}

// ValidateERC1155OrderSignature is a free data retrieval call binding the contract method 0x0d32a531.
//
// Solidity: function validateERC1155OrderSignature((uint8,address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[],uint128) order, (uint8,uint8,bytes32,bytes32) signature) view returns()
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureCallerSession) ValidateERC1155OrderSignature(order LibNFTOrderERC1155Order, signature LibSignatureSignature) error {
	return _ERC1155OrdersFeature.Contract.ValidateERC1155OrderSignature(&_ERC1155OrdersFeature.CallOpts, order, signature)
}

// BuyERC1155 is a paid mutator transaction binding the contract method 0xa4f4d30d.
//
// Solidity: function _buyERC1155((uint8,address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[],uint128) sellOrder, (uint8,uint8,bytes32,bytes32) signature, (uint128,uint256,bytes) params) payable returns()
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureTransactor) BuyERC1155(opts *bind.TransactOpts, sellOrder LibNFTOrderERC1155Order, signature LibSignatureSignature, params NFTOrdersBuyParams) (*types.Transaction, error) {
	return _ERC1155OrdersFeature.contract.Transact(opts, "_buyERC1155", sellOrder, signature, params)
}

// BuyERC1155 is a paid mutator transaction binding the contract method 0xa4f4d30d.
//
// Solidity: function _buyERC1155((uint8,address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[],uint128) sellOrder, (uint8,uint8,bytes32,bytes32) signature, (uint128,uint256,bytes) params) payable returns()
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureSession) BuyERC1155(sellOrder LibNFTOrderERC1155Order, signature LibSignatureSignature, params NFTOrdersBuyParams) (*types.Transaction, error) {
	return _ERC1155OrdersFeature.Contract.BuyERC1155(&_ERC1155OrdersFeature.TransactOpts, sellOrder, signature, params)
}

// BuyERC1155 is a paid mutator transaction binding the contract method 0xa4f4d30d.
//
// Solidity: function _buyERC1155((uint8,address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[],uint128) sellOrder, (uint8,uint8,bytes32,bytes32) signature, (uint128,uint256,bytes) params) payable returns()
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureTransactorSession) BuyERC1155(sellOrder LibNFTOrderERC1155Order, signature LibSignatureSignature, params NFTOrdersBuyParams) (*types.Transaction, error) {
	return _ERC1155OrdersFeature.Contract.BuyERC1155(&_ERC1155OrdersFeature.TransactOpts, sellOrder, signature, params)
}

// BatchBuyERC1155s is a paid mutator transaction binding the contract method 0x84680615.
//
// Solidity: function batchBuyERC1155s((uint8,address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[],uint128)[] sellOrders, (uint8,uint8,bytes32,bytes32)[] signatures, uint128[] erc1155FillAmounts, bytes[] callbackData, bool revertIfIncomplete) payable returns(bool[] successes)
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureTransactor) BatchBuyERC1155s(opts *bind.TransactOpts, sellOrders []LibNFTOrderERC1155Order, signatures []LibSignatureSignature, erc1155FillAmounts []*big.Int, callbackData [][]byte, revertIfIncomplete bool) (*types.Transaction, error) {
	return _ERC1155OrdersFeature.contract.Transact(opts, "batchBuyERC1155s", sellOrders, signatures, erc1155FillAmounts, callbackData, revertIfIncomplete)
}

// BatchBuyERC1155s is a paid mutator transaction binding the contract method 0x84680615.
//
// Solidity: function batchBuyERC1155s((uint8,address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[],uint128)[] sellOrders, (uint8,uint8,bytes32,bytes32)[] signatures, uint128[] erc1155FillAmounts, bytes[] callbackData, bool revertIfIncomplete) payable returns(bool[] successes)
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureSession) BatchBuyERC1155s(sellOrders []LibNFTOrderERC1155Order, signatures []LibSignatureSignature, erc1155FillAmounts []*big.Int, callbackData [][]byte, revertIfIncomplete bool) (*types.Transaction, error) {
	return _ERC1155OrdersFeature.Contract.BatchBuyERC1155s(&_ERC1155OrdersFeature.TransactOpts, sellOrders, signatures, erc1155FillAmounts, callbackData, revertIfIncomplete)
}

// BatchBuyERC1155s is a paid mutator transaction binding the contract method 0x84680615.
//
// Solidity: function batchBuyERC1155s((uint8,address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[],uint128)[] sellOrders, (uint8,uint8,bytes32,bytes32)[] signatures, uint128[] erc1155FillAmounts, bytes[] callbackData, bool revertIfIncomplete) payable returns(bool[] successes)
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureTransactorSession) BatchBuyERC1155s(sellOrders []LibNFTOrderERC1155Order, signatures []LibSignatureSignature, erc1155FillAmounts []*big.Int, callbackData [][]byte, revertIfIncomplete bool) (*types.Transaction, error) {
	return _ERC1155OrdersFeature.Contract.BatchBuyERC1155s(&_ERC1155OrdersFeature.TransactOpts, sellOrders, signatures, erc1155FillAmounts, callbackData, revertIfIncomplete)
}

// BatchCancelERC1155Orders is a paid mutator transaction binding the contract method 0xa1865d6f.
//
// Solidity: function batchCancelERC1155Orders(uint256[] orderNonces) returns()
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureTransactor) BatchCancelERC1155Orders(opts *bind.TransactOpts, orderNonces []*big.Int) (*types.Transaction, error) {
	return _ERC1155OrdersFeature.contract.Transact(opts, "batchCancelERC1155Orders", orderNonces)
}

// BatchCancelERC1155Orders is a paid mutator transaction binding the contract method 0xa1865d6f.
//
// Solidity: function batchCancelERC1155Orders(uint256[] orderNonces) returns()
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureSession) BatchCancelERC1155Orders(orderNonces []*big.Int) (*types.Transaction, error) {
	return _ERC1155OrdersFeature.Contract.BatchCancelERC1155Orders(&_ERC1155OrdersFeature.TransactOpts, orderNonces)
}

// BatchCancelERC1155Orders is a paid mutator transaction binding the contract method 0xa1865d6f.
//
// Solidity: function batchCancelERC1155Orders(uint256[] orderNonces) returns()
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureTransactorSession) BatchCancelERC1155Orders(orderNonces []*big.Int) (*types.Transaction, error) {
	return _ERC1155OrdersFeature.Contract.BatchCancelERC1155Orders(&_ERC1155OrdersFeature.TransactOpts, orderNonces)
}

// BuyERC11552 is a paid mutator transaction binding the contract method 0x7cdb54d8.
//
// Solidity: function buyERC1155((uint8,address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[],uint128) sellOrder, (uint8,uint8,bytes32,bytes32) signature, uint128 erc1155BuyAmount, bytes callbackData) payable returns()
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureTransactor) BuyERC11552(opts *bind.TransactOpts, sellOrder LibNFTOrderERC1155Order, signature LibSignatureSignature, erc1155BuyAmount *big.Int, callbackData []byte) (*types.Transaction, error) {
	return _ERC1155OrdersFeature.contract.Transact(opts, "buyERC1155", sellOrder, signature, erc1155BuyAmount, callbackData)
}

// BuyERC11552 is a paid mutator transaction binding the contract method 0x7cdb54d8.
//
// Solidity: function buyERC1155((uint8,address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[],uint128) sellOrder, (uint8,uint8,bytes32,bytes32) signature, uint128 erc1155BuyAmount, bytes callbackData) payable returns()
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureSession) BuyERC11552(sellOrder LibNFTOrderERC1155Order, signature LibSignatureSignature, erc1155BuyAmount *big.Int, callbackData []byte) (*types.Transaction, error) {
	return _ERC1155OrdersFeature.Contract.BuyERC11552(&_ERC1155OrdersFeature.TransactOpts, sellOrder, signature, erc1155BuyAmount, callbackData)
}

// BuyERC11552 is a paid mutator transaction binding the contract method 0x7cdb54d8.
//
// Solidity: function buyERC1155((uint8,address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[],uint128) sellOrder, (uint8,uint8,bytes32,bytes32) signature, uint128 erc1155BuyAmount, bytes callbackData) payable returns()
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureTransactorSession) BuyERC11552(sellOrder LibNFTOrderERC1155Order, signature LibSignatureSignature, erc1155BuyAmount *big.Int, callbackData []byte) (*types.Transaction, error) {
	return _ERC1155OrdersFeature.Contract.BuyERC11552(&_ERC1155OrdersFeature.TransactOpts, sellOrder, signature, erc1155BuyAmount, callbackData)
}

// CancelERC1155Order is a paid mutator transaction binding the contract method 0x06d2596b.
//
// Solidity: function cancelERC1155Order(uint256 orderNonce) returns()
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureTransactor) CancelERC1155Order(opts *bind.TransactOpts, orderNonce *big.Int) (*types.Transaction, error) {
	return _ERC1155OrdersFeature.contract.Transact(opts, "cancelERC1155Order", orderNonce)
}

// CancelERC1155Order is a paid mutator transaction binding the contract method 0x06d2596b.
//
// Solidity: function cancelERC1155Order(uint256 orderNonce) returns()
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureSession) CancelERC1155Order(orderNonce *big.Int) (*types.Transaction, error) {
	return _ERC1155OrdersFeature.Contract.CancelERC1155Order(&_ERC1155OrdersFeature.TransactOpts, orderNonce)
}

// CancelERC1155Order is a paid mutator transaction binding the contract method 0x06d2596b.
//
// Solidity: function cancelERC1155Order(uint256 orderNonce) returns()
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureTransactorSession) CancelERC1155Order(orderNonce *big.Int) (*types.Transaction, error) {
	return _ERC1155OrdersFeature.Contract.CancelERC1155Order(&_ERC1155OrdersFeature.TransactOpts, orderNonce)
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns(bytes4 success)
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureTransactor) Migrate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC1155OrdersFeature.contract.Transact(opts, "migrate")
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns(bytes4 success)
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureSession) Migrate() (*types.Transaction, error) {
	return _ERC1155OrdersFeature.Contract.Migrate(&_ERC1155OrdersFeature.TransactOpts)
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns(bytes4 success)
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureTransactorSession) Migrate() (*types.Transaction, error) {
	return _ERC1155OrdersFeature.Contract.Migrate(&_ERC1155OrdersFeature.TransactOpts)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address operator, address , uint256 tokenId, uint256 value, bytes data) returns(bytes4 success)
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureTransactor) OnERC1155Received(opts *bind.TransactOpts, operator common.Address, arg1 common.Address, tokenId *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC1155OrdersFeature.contract.Transact(opts, "onERC1155Received", operator, arg1, tokenId, value, data)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address operator, address , uint256 tokenId, uint256 value, bytes data) returns(bytes4 success)
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureSession) OnERC1155Received(operator common.Address, arg1 common.Address, tokenId *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC1155OrdersFeature.Contract.OnERC1155Received(&_ERC1155OrdersFeature.TransactOpts, operator, arg1, tokenId, value, data)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address operator, address , uint256 tokenId, uint256 value, bytes data) returns(bytes4 success)
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureTransactorSession) OnERC1155Received(operator common.Address, arg1 common.Address, tokenId *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC1155OrdersFeature.Contract.OnERC1155Received(&_ERC1155OrdersFeature.TransactOpts, operator, arg1, tokenId, value, data)
}

// PreSignERC1155Order is a paid mutator transaction binding the contract method 0x7b757d97.
//
// Solidity: function preSignERC1155Order((uint8,address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[],uint128) order) returns()
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureTransactor) PreSignERC1155Order(opts *bind.TransactOpts, order LibNFTOrderERC1155Order) (*types.Transaction, error) {
	return _ERC1155OrdersFeature.contract.Transact(opts, "preSignERC1155Order", order)
}

// PreSignERC1155Order is a paid mutator transaction binding the contract method 0x7b757d97.
//
// Solidity: function preSignERC1155Order((uint8,address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[],uint128) order) returns()
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureSession) PreSignERC1155Order(order LibNFTOrderERC1155Order) (*types.Transaction, error) {
	return _ERC1155OrdersFeature.Contract.PreSignERC1155Order(&_ERC1155OrdersFeature.TransactOpts, order)
}

// PreSignERC1155Order is a paid mutator transaction binding the contract method 0x7b757d97.
//
// Solidity: function preSignERC1155Order((uint8,address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[],uint128) order) returns()
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureTransactorSession) PreSignERC1155Order(order LibNFTOrderERC1155Order) (*types.Transaction, error) {
	return _ERC1155OrdersFeature.Contract.PreSignERC1155Order(&_ERC1155OrdersFeature.TransactOpts, order)
}

// SellERC1155 is a paid mutator transaction binding the contract method 0x6e2eec9e.
//
// Solidity: function sellERC1155((uint8,address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[],uint128) buyOrder, (uint8,uint8,bytes32,bytes32) signature, uint256 erc1155TokenId, uint128 erc1155SellAmount, bool unwrapNativeToken, bytes callbackData) returns()
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureTransactor) SellERC1155(opts *bind.TransactOpts, buyOrder LibNFTOrderERC1155Order, signature LibSignatureSignature, erc1155TokenId *big.Int, erc1155SellAmount *big.Int, unwrapNativeToken bool, callbackData []byte) (*types.Transaction, error) {
	return _ERC1155OrdersFeature.contract.Transact(opts, "sellERC1155", buyOrder, signature, erc1155TokenId, erc1155SellAmount, unwrapNativeToken, callbackData)
}

// SellERC1155 is a paid mutator transaction binding the contract method 0x6e2eec9e.
//
// Solidity: function sellERC1155((uint8,address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[],uint128) buyOrder, (uint8,uint8,bytes32,bytes32) signature, uint256 erc1155TokenId, uint128 erc1155SellAmount, bool unwrapNativeToken, bytes callbackData) returns()
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureSession) SellERC1155(buyOrder LibNFTOrderERC1155Order, signature LibSignatureSignature, erc1155TokenId *big.Int, erc1155SellAmount *big.Int, unwrapNativeToken bool, callbackData []byte) (*types.Transaction, error) {
	return _ERC1155OrdersFeature.Contract.SellERC1155(&_ERC1155OrdersFeature.TransactOpts, buyOrder, signature, erc1155TokenId, erc1155SellAmount, unwrapNativeToken, callbackData)
}

// SellERC1155 is a paid mutator transaction binding the contract method 0x6e2eec9e.
//
// Solidity: function sellERC1155((uint8,address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[],uint128) buyOrder, (uint8,uint8,bytes32,bytes32) signature, uint256 erc1155TokenId, uint128 erc1155SellAmount, bool unwrapNativeToken, bytes callbackData) returns()
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureTransactorSession) SellERC1155(buyOrder LibNFTOrderERC1155Order, signature LibSignatureSignature, erc1155TokenId *big.Int, erc1155SellAmount *big.Int, unwrapNativeToken bool, callbackData []byte) (*types.Transaction, error) {
	return _ERC1155OrdersFeature.Contract.SellERC1155(&_ERC1155OrdersFeature.TransactOpts, buyOrder, signature, erc1155TokenId, erc1155SellAmount, unwrapNativeToken, callbackData)
}

// ERC1155OrdersFeatureERC1155OrderCancelledIterator is returned from FilterERC1155OrderCancelled and is used to iterate over the raw logs and unpacked data for ERC1155OrderCancelled events raised by the ERC1155OrdersFeature contract.
type ERC1155OrdersFeatureERC1155OrderCancelledIterator struct {
	Event *ERC1155OrdersFeatureERC1155OrderCancelled // Event containing the contract specifics and raw log

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
func (it *ERC1155OrdersFeatureERC1155OrderCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155OrdersFeatureERC1155OrderCancelled)
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
		it.Event = new(ERC1155OrdersFeatureERC1155OrderCancelled)
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
func (it *ERC1155OrdersFeatureERC1155OrderCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155OrdersFeatureERC1155OrderCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155OrdersFeatureERC1155OrderCancelled represents a ERC1155OrderCancelled event raised by the ERC1155OrdersFeature contract.
type ERC1155OrdersFeatureERC1155OrderCancelled struct {
	Maker common.Address
	Nonce *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterERC1155OrderCancelled is a free log retrieval operation binding the contract event 0x4d5ea7da64f50a4a329921b8d2cab52dff4ebcc58b61d10ff839e28e91445684.
//
// Solidity: event ERC1155OrderCancelled(address maker, uint256 nonce)
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureFilterer) FilterERC1155OrderCancelled(opts *bind.FilterOpts) (*ERC1155OrdersFeatureERC1155OrderCancelledIterator, error) {

	logs, sub, err := _ERC1155OrdersFeature.contract.FilterLogs(opts, "ERC1155OrderCancelled")
	if err != nil {
		return nil, err
	}
	return &ERC1155OrdersFeatureERC1155OrderCancelledIterator{contract: _ERC1155OrdersFeature.contract, event: "ERC1155OrderCancelled", logs: logs, sub: sub}, nil
}

// WatchERC1155OrderCancelled is a free log subscription operation binding the contract event 0x4d5ea7da64f50a4a329921b8d2cab52dff4ebcc58b61d10ff839e28e91445684.
//
// Solidity: event ERC1155OrderCancelled(address maker, uint256 nonce)
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureFilterer) WatchERC1155OrderCancelled(opts *bind.WatchOpts, sink chan<- *ERC1155OrdersFeatureERC1155OrderCancelled) (event.Subscription, error) {

	logs, sub, err := _ERC1155OrdersFeature.contract.WatchLogs(opts, "ERC1155OrderCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155OrdersFeatureERC1155OrderCancelled)
				if err := _ERC1155OrdersFeature.contract.UnpackLog(event, "ERC1155OrderCancelled", log); err != nil {
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

// ParseERC1155OrderCancelled is a log parse operation binding the contract event 0x4d5ea7da64f50a4a329921b8d2cab52dff4ebcc58b61d10ff839e28e91445684.
//
// Solidity: event ERC1155OrderCancelled(address maker, uint256 nonce)
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureFilterer) ParseERC1155OrderCancelled(log types.Log) (*ERC1155OrdersFeatureERC1155OrderCancelled, error) {
	event := new(ERC1155OrdersFeatureERC1155OrderCancelled)
	if err := _ERC1155OrdersFeature.contract.UnpackLog(event, "ERC1155OrderCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155OrdersFeatureERC1155OrderFilledIterator is returned from FilterERC1155OrderFilled and is used to iterate over the raw logs and unpacked data for ERC1155OrderFilled events raised by the ERC1155OrdersFeature contract.
type ERC1155OrdersFeatureERC1155OrderFilledIterator struct {
	Event *ERC1155OrdersFeatureERC1155OrderFilled // Event containing the contract specifics and raw log

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
func (it *ERC1155OrdersFeatureERC1155OrderFilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155OrdersFeatureERC1155OrderFilled)
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
		it.Event = new(ERC1155OrdersFeatureERC1155OrderFilled)
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
func (it *ERC1155OrdersFeatureERC1155OrderFilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155OrdersFeatureERC1155OrderFilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155OrdersFeatureERC1155OrderFilled represents a ERC1155OrderFilled event raised by the ERC1155OrdersFeature contract.
type ERC1155OrdersFeatureERC1155OrderFilled struct {
	Direction         uint8
	Maker             common.Address
	Taker             common.Address
	Nonce             *big.Int
	Erc20Token        common.Address
	Erc20FillAmount   *big.Int
	Erc1155Token      common.Address
	Erc1155TokenId    *big.Int
	Erc1155FillAmount *big.Int
	Matcher           common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterERC1155OrderFilled is a free log retrieval operation binding the contract event 0x20cca81b0e269b265b3229d6b537da91ef475ca0ef55caed7dd30731700ba98d.
//
// Solidity: event ERC1155OrderFilled(uint8 direction, address maker, address taker, uint256 nonce, address erc20Token, uint256 erc20FillAmount, address erc1155Token, uint256 erc1155TokenId, uint128 erc1155FillAmount, address matcher)
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureFilterer) FilterERC1155OrderFilled(opts *bind.FilterOpts) (*ERC1155OrdersFeatureERC1155OrderFilledIterator, error) {

	logs, sub, err := _ERC1155OrdersFeature.contract.FilterLogs(opts, "ERC1155OrderFilled")
	if err != nil {
		return nil, err
	}
	return &ERC1155OrdersFeatureERC1155OrderFilledIterator{contract: _ERC1155OrdersFeature.contract, event: "ERC1155OrderFilled", logs: logs, sub: sub}, nil
}

// WatchERC1155OrderFilled is a free log subscription operation binding the contract event 0x20cca81b0e269b265b3229d6b537da91ef475ca0ef55caed7dd30731700ba98d.
//
// Solidity: event ERC1155OrderFilled(uint8 direction, address maker, address taker, uint256 nonce, address erc20Token, uint256 erc20FillAmount, address erc1155Token, uint256 erc1155TokenId, uint128 erc1155FillAmount, address matcher)
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureFilterer) WatchERC1155OrderFilled(opts *bind.WatchOpts, sink chan<- *ERC1155OrdersFeatureERC1155OrderFilled) (event.Subscription, error) {

	logs, sub, err := _ERC1155OrdersFeature.contract.WatchLogs(opts, "ERC1155OrderFilled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155OrdersFeatureERC1155OrderFilled)
				if err := _ERC1155OrdersFeature.contract.UnpackLog(event, "ERC1155OrderFilled", log); err != nil {
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

// ParseERC1155OrderFilled is a log parse operation binding the contract event 0x20cca81b0e269b265b3229d6b537da91ef475ca0ef55caed7dd30731700ba98d.
//
// Solidity: event ERC1155OrderFilled(uint8 direction, address maker, address taker, uint256 nonce, address erc20Token, uint256 erc20FillAmount, address erc1155Token, uint256 erc1155TokenId, uint128 erc1155FillAmount, address matcher)
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureFilterer) ParseERC1155OrderFilled(log types.Log) (*ERC1155OrdersFeatureERC1155OrderFilled, error) {
	event := new(ERC1155OrdersFeatureERC1155OrderFilled)
	if err := _ERC1155OrdersFeature.contract.UnpackLog(event, "ERC1155OrderFilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155OrdersFeatureERC1155OrderPreSignedIterator is returned from FilterERC1155OrderPreSigned and is used to iterate over the raw logs and unpacked data for ERC1155OrderPreSigned events raised by the ERC1155OrdersFeature contract.
type ERC1155OrdersFeatureERC1155OrderPreSignedIterator struct {
	Event *ERC1155OrdersFeatureERC1155OrderPreSigned // Event containing the contract specifics and raw log

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
func (it *ERC1155OrdersFeatureERC1155OrderPreSignedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155OrdersFeatureERC1155OrderPreSigned)
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
		it.Event = new(ERC1155OrdersFeatureERC1155OrderPreSigned)
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
func (it *ERC1155OrdersFeatureERC1155OrderPreSignedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155OrdersFeatureERC1155OrderPreSignedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155OrdersFeatureERC1155OrderPreSigned represents a ERC1155OrderPreSigned event raised by the ERC1155OrdersFeature contract.
type ERC1155OrdersFeatureERC1155OrderPreSigned struct {
	Direction              uint8
	Maker                  common.Address
	Taker                  common.Address
	Expiry                 *big.Int
	Nonce                  *big.Int
	Erc20Token             common.Address
	Erc20TokenAmount       *big.Int
	Fees                   []LibNFTOrderFee
	Erc1155Token           common.Address
	Erc1155TokenId         *big.Int
	Erc1155TokenProperties []LibNFTOrderProperty
	Erc1155TokenAmount     *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterERC1155OrderPreSigned is a free log retrieval operation binding the contract event 0x5e91ddfeb7bf2e12f7e8ab017d2b63a9217f004a15a53346ad90353ec63d14e4.
//
// Solidity: event ERC1155OrderPreSigned(uint8 direction, address maker, address taker, uint256 expiry, uint256 nonce, address erc20Token, uint256 erc20TokenAmount, (address,uint256,bytes)[] fees, address erc1155Token, uint256 erc1155TokenId, (address,bytes)[] erc1155TokenProperties, uint128 erc1155TokenAmount)
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureFilterer) FilterERC1155OrderPreSigned(opts *bind.FilterOpts) (*ERC1155OrdersFeatureERC1155OrderPreSignedIterator, error) {

	logs, sub, err := _ERC1155OrdersFeature.contract.FilterLogs(opts, "ERC1155OrderPreSigned")
	if err != nil {
		return nil, err
	}
	return &ERC1155OrdersFeatureERC1155OrderPreSignedIterator{contract: _ERC1155OrdersFeature.contract, event: "ERC1155OrderPreSigned", logs: logs, sub: sub}, nil
}

// WatchERC1155OrderPreSigned is a free log subscription operation binding the contract event 0x5e91ddfeb7bf2e12f7e8ab017d2b63a9217f004a15a53346ad90353ec63d14e4.
//
// Solidity: event ERC1155OrderPreSigned(uint8 direction, address maker, address taker, uint256 expiry, uint256 nonce, address erc20Token, uint256 erc20TokenAmount, (address,uint256,bytes)[] fees, address erc1155Token, uint256 erc1155TokenId, (address,bytes)[] erc1155TokenProperties, uint128 erc1155TokenAmount)
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureFilterer) WatchERC1155OrderPreSigned(opts *bind.WatchOpts, sink chan<- *ERC1155OrdersFeatureERC1155OrderPreSigned) (event.Subscription, error) {

	logs, sub, err := _ERC1155OrdersFeature.contract.WatchLogs(opts, "ERC1155OrderPreSigned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155OrdersFeatureERC1155OrderPreSigned)
				if err := _ERC1155OrdersFeature.contract.UnpackLog(event, "ERC1155OrderPreSigned", log); err != nil {
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

// ParseERC1155OrderPreSigned is a log parse operation binding the contract event 0x5e91ddfeb7bf2e12f7e8ab017d2b63a9217f004a15a53346ad90353ec63d14e4.
//
// Solidity: event ERC1155OrderPreSigned(uint8 direction, address maker, address taker, uint256 expiry, uint256 nonce, address erc20Token, uint256 erc20TokenAmount, (address,uint256,bytes)[] fees, address erc1155Token, uint256 erc1155TokenId, (address,bytes)[] erc1155TokenProperties, uint128 erc1155TokenAmount)
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureFilterer) ParseERC1155OrderPreSigned(log types.Log) (*ERC1155OrdersFeatureERC1155OrderPreSigned, error) {
	event := new(ERC1155OrdersFeatureERC1155OrderPreSigned)
	if err := _ERC1155OrdersFeature.contract.UnpackLog(event, "ERC1155OrderPreSigned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
