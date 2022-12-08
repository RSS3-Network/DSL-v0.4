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

// LibNFTOrderERC721Order is an auto generated low-level Go binding around an user-defined struct.
type LibNFTOrderERC721Order struct {
	Direction             uint8
	Maker                 common.Address
	Taker                 common.Address
	Expiry                *big.Int
	Nonce                 *big.Int
	Erc20Token            common.Address
	Erc20TokenAmount      *big.Int
	Fees                  []LibNFTOrderFee
	Erc721Token           common.Address
	Erc721TokenId         *big.Int
	Erc721TokenProperties []LibNFTOrderProperty
}

// LibNFTOrderFee is an auto generated low-level Go binding around an user-defined struct.
type LibNFTOrderFee struct {
	Recipient common.Address
	Amount    *big.Int
	FeeData   []byte
}

// LibNFTOrderProperty is an auto generated low-level Go binding around an user-defined struct.
type LibNFTOrderProperty struct {
	PropertyValidator common.Address
	PropertyData      []byte
}

// LibSignatureSignature is an auto generated low-level Go binding around an user-defined struct.
type LibSignatureSignature struct {
	SignatureType uint8
	V             uint8
	R             [32]byte
	S             [32]byte
}

// ERC721OrdersFeatureMetaData contains all meta data concerning the ERC721OrdersFeature contract.
var ERC721OrdersFeatureMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"zeroExAddress\",\"type\":\"address\"},{\"internalType\":\"contractIEtherTokenV06\",\"name\":\"weth\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"ERC721OrderCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumLibNFTOrder.TradeDirection\",\"name\":\"direction\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIERC20TokenV06\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"erc20TokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIERC721Token\",\"name\":\"erc721Token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"erc721TokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"matcher\",\"type\":\"address\"}],\"name\":\"ERC721OrderFilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumLibNFTOrder.TradeDirection\",\"name\":\"direction\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIERC20TokenV06\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"erc20TokenAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structLibNFTOrder.Fee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"contractIERC721Token\",\"name\":\"erc721Token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"erc721TokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractIPropertyValidator\",\"name\":\"propertyValidator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"propertyData\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structLibNFTOrder.Property[]\",\"name\":\"erc721TokenProperties\",\"type\":\"tuple[]\"}],\"name\":\"ERC721OrderPreSigned\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"EIP712_DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FEATURE_NAME\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FEATURE_VERSION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumLibNFTOrder.TradeDirection\",\"name\":\"direction\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20TokenV06\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc20TokenAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Fee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"contractIERC721Token\",\"name\":\"erc721Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc721TokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractIPropertyValidator\",\"name\":\"propertyValidator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"propertyData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Property[]\",\"name\":\"erc721TokenProperties\",\"type\":\"tuple[]\"}],\"internalType\":\"structLibNFTOrder.ERC721Order\",\"name\":\"sellOrder\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumLibSignature.SignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structLibSignature.Signature\",\"name\":\"signature\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"ethAvailable\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"takerCallbackData\",\"type\":\"bytes\"}],\"name\":\"_buyERC721\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumLibNFTOrder.TradeDirection\",\"name\":\"direction\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20TokenV06\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc20TokenAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Fee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"contractIERC721Token\",\"name\":\"erc721Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc721TokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractIPropertyValidator\",\"name\":\"propertyValidator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"propertyData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Property[]\",\"name\":\"erc721TokenProperties\",\"type\":\"tuple[]\"}],\"internalType\":\"structLibNFTOrder.ERC721Order[]\",\"name\":\"sellOrders\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumLibSignature.SignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structLibSignature.Signature[]\",\"name\":\"signatures\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"callbackData\",\"type\":\"bytes[]\"},{\"internalType\":\"bool\",\"name\":\"revertIfIncomplete\",\"type\":\"bool\"}],\"name\":\"batchBuyERC721s\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"successes\",\"type\":\"bool[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"orderNonces\",\"type\":\"uint256[]\"}],\"name\":\"batchCancelERC721Orders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumLibNFTOrder.TradeDirection\",\"name\":\"direction\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20TokenV06\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc20TokenAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Fee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"contractIERC721Token\",\"name\":\"erc721Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc721TokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractIPropertyValidator\",\"name\":\"propertyValidator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"propertyData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Property[]\",\"name\":\"erc721TokenProperties\",\"type\":\"tuple[]\"}],\"internalType\":\"structLibNFTOrder.ERC721Order[]\",\"name\":\"sellOrders\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumLibNFTOrder.TradeDirection\",\"name\":\"direction\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20TokenV06\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc20TokenAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Fee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"contractIERC721Token\",\"name\":\"erc721Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc721TokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractIPropertyValidator\",\"name\":\"propertyValidator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"propertyData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Property[]\",\"name\":\"erc721TokenProperties\",\"type\":\"tuple[]\"}],\"internalType\":\"structLibNFTOrder.ERC721Order[]\",\"name\":\"buyOrders\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumLibSignature.SignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structLibSignature.Signature[]\",\"name\":\"sellOrderSignatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumLibSignature.SignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structLibSignature.Signature[]\",\"name\":\"buyOrderSignatures\",\"type\":\"tuple[]\"}],\"name\":\"batchMatchERC721Orders\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"profits\",\"type\":\"uint256[]\"},{\"internalType\":\"bool[]\",\"name\":\"successes\",\"type\":\"bool[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumLibNFTOrder.TradeDirection\",\"name\":\"direction\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20TokenV06\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc20TokenAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Fee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"contractIERC721Token\",\"name\":\"erc721Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc721TokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractIPropertyValidator\",\"name\":\"propertyValidator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"propertyData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Property[]\",\"name\":\"erc721TokenProperties\",\"type\":\"tuple[]\"}],\"internalType\":\"structLibNFTOrder.ERC721Order\",\"name\":\"sellOrder\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumLibSignature.SignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structLibSignature.Signature\",\"name\":\"signature\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"callbackData\",\"type\":\"bytes\"}],\"name\":\"buyERC721\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderNonce\",\"type\":\"uint256\"}],\"name\":\"cancelERC721Order\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumLibNFTOrder.TradeDirection\",\"name\":\"direction\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20TokenV06\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc20TokenAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Fee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"contractIERC721Token\",\"name\":\"erc721Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc721TokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractIPropertyValidator\",\"name\":\"propertyValidator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"propertyData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Property[]\",\"name\":\"erc721TokenProperties\",\"type\":\"tuple[]\"}],\"internalType\":\"structLibNFTOrder.ERC721Order\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"getERC721OrderHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumLibNFTOrder.TradeDirection\",\"name\":\"direction\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20TokenV06\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc20TokenAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Fee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"contractIERC721Token\",\"name\":\"erc721Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc721TokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractIPropertyValidator\",\"name\":\"propertyValidator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"propertyData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Property[]\",\"name\":\"erc721TokenProperties\",\"type\":\"tuple[]\"}],\"internalType\":\"structLibNFTOrder.ERC721Order\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"getERC721OrderStatus\",\"outputs\":[{\"internalType\":\"enumLibNFTOrder.OrderStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"uint248\",\"name\":\"nonceRange\",\"type\":\"uint248\"}],\"name\":\"getERC721OrderStatusBitVector\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"bitVector\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumLibNFTOrder.TradeDirection\",\"name\":\"direction\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20TokenV06\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc20TokenAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Fee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"contractIERC721Token\",\"name\":\"erc721Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc721TokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractIPropertyValidator\",\"name\":\"propertyValidator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"propertyData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Property[]\",\"name\":\"erc721TokenProperties\",\"type\":\"tuple[]\"}],\"internalType\":\"structLibNFTOrder.ERC721Order\",\"name\":\"sellOrder\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumLibNFTOrder.TradeDirection\",\"name\":\"direction\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20TokenV06\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc20TokenAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Fee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"contractIERC721Token\",\"name\":\"erc721Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc721TokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractIPropertyValidator\",\"name\":\"propertyValidator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"propertyData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Property[]\",\"name\":\"erc721TokenProperties\",\"type\":\"tuple[]\"}],\"internalType\":\"structLibNFTOrder.ERC721Order\",\"name\":\"buyOrder\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumLibSignature.SignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structLibSignature.Signature\",\"name\":\"sellOrderSignature\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumLibSignature.SignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structLibSignature.Signature\",\"name\":\"buyOrderSignature\",\"type\":\"tuple\"}],\"name\":\"matchERC721Orders\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"profit\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"migrate\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"success\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"success\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumLibNFTOrder.TradeDirection\",\"name\":\"direction\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20TokenV06\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc20TokenAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Fee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"contractIERC721Token\",\"name\":\"erc721Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc721TokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractIPropertyValidator\",\"name\":\"propertyValidator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"propertyData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Property[]\",\"name\":\"erc721TokenProperties\",\"type\":\"tuple[]\"}],\"internalType\":\"structLibNFTOrder.ERC721Order\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"preSignERC721Order\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumLibNFTOrder.TradeDirection\",\"name\":\"direction\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20TokenV06\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc20TokenAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Fee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"contractIERC721Token\",\"name\":\"erc721Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc721TokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractIPropertyValidator\",\"name\":\"propertyValidator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"propertyData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Property[]\",\"name\":\"erc721TokenProperties\",\"type\":\"tuple[]\"}],\"internalType\":\"structLibNFTOrder.ERC721Order\",\"name\":\"buyOrder\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumLibSignature.SignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structLibSignature.Signature\",\"name\":\"signature\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"erc721TokenId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"unwrapNativeToken\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"callbackData\",\"type\":\"bytes\"}],\"name\":\"sellERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumLibNFTOrder.TradeDirection\",\"name\":\"direction\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20TokenV06\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc20TokenAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Fee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"contractIERC721Token\",\"name\":\"erc721Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc721TokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractIPropertyValidator\",\"name\":\"propertyValidator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"propertyData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Property[]\",\"name\":\"erc721TokenProperties\",\"type\":\"tuple[]\"}],\"internalType\":\"structLibNFTOrder.ERC721Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"erc721TokenId\",\"type\":\"uint256\"}],\"name\":\"validateERC721OrderProperties\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumLibNFTOrder.TradeDirection\",\"name\":\"direction\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20TokenV06\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc20TokenAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Fee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"contractIERC721Token\",\"name\":\"erc721Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc721TokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractIPropertyValidator\",\"name\":\"propertyValidator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"propertyData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Property[]\",\"name\":\"erc721TokenProperties\",\"type\":\"tuple[]\"}],\"internalType\":\"structLibNFTOrder.ERC721Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumLibSignature.SignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structLibSignature.Signature\",\"name\":\"signature\",\"type\":\"tuple\"}],\"name\":\"validateERC721OrderSignature\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ERC721OrdersFeatureABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC721OrdersFeatureMetaData.ABI instead.
var ERC721OrdersFeatureABI = ERC721OrdersFeatureMetaData.ABI

// ERC721OrdersFeature is an auto generated Go binding around an Ethereum contract.
type ERC721OrdersFeature struct {
	ERC721OrdersFeatureCaller     // Read-only binding to the contract
	ERC721OrdersFeatureTransactor // Write-only binding to the contract
	ERC721OrdersFeatureFilterer   // Log filterer for contract events
}

// ERC721OrdersFeatureCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC721OrdersFeatureCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721OrdersFeatureTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC721OrdersFeatureTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721OrdersFeatureFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC721OrdersFeatureFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721OrdersFeatureSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC721OrdersFeatureSession struct {
	Contract     *ERC721OrdersFeature // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ERC721OrdersFeatureCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC721OrdersFeatureCallerSession struct {
	Contract *ERC721OrdersFeatureCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// ERC721OrdersFeatureTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC721OrdersFeatureTransactorSession struct {
	Contract     *ERC721OrdersFeatureTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// ERC721OrdersFeatureRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC721OrdersFeatureRaw struct {
	Contract *ERC721OrdersFeature // Generic contract binding to access the raw methods on
}

// ERC721OrdersFeatureCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC721OrdersFeatureCallerRaw struct {
	Contract *ERC721OrdersFeatureCaller // Generic read-only contract binding to access the raw methods on
}

// ERC721OrdersFeatureTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC721OrdersFeatureTransactorRaw struct {
	Contract *ERC721OrdersFeatureTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC721OrdersFeature creates a new instance of ERC721OrdersFeature, bound to a specific deployed contract.
func NewERC721OrdersFeature(address common.Address, backend bind.ContractBackend) (*ERC721OrdersFeature, error) {
	contract, err := bindERC721OrdersFeature(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC721OrdersFeature{ERC721OrdersFeatureCaller: ERC721OrdersFeatureCaller{contract: contract}, ERC721OrdersFeatureTransactor: ERC721OrdersFeatureTransactor{contract: contract}, ERC721OrdersFeatureFilterer: ERC721OrdersFeatureFilterer{contract: contract}}, nil
}

// NewERC721OrdersFeatureCaller creates a new read-only instance of ERC721OrdersFeature, bound to a specific deployed contract.
func NewERC721OrdersFeatureCaller(address common.Address, caller bind.ContractCaller) (*ERC721OrdersFeatureCaller, error) {
	contract, err := bindERC721OrdersFeature(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721OrdersFeatureCaller{contract: contract}, nil
}

// NewERC721OrdersFeatureTransactor creates a new write-only instance of ERC721OrdersFeature, bound to a specific deployed contract.
func NewERC721OrdersFeatureTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC721OrdersFeatureTransactor, error) {
	contract, err := bindERC721OrdersFeature(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721OrdersFeatureTransactor{contract: contract}, nil
}

// NewERC721OrdersFeatureFilterer creates a new log filterer instance of ERC721OrdersFeature, bound to a specific deployed contract.
func NewERC721OrdersFeatureFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC721OrdersFeatureFilterer, error) {
	contract, err := bindERC721OrdersFeature(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC721OrdersFeatureFilterer{contract: contract}, nil
}

// bindERC721OrdersFeature binds a generic wrapper to an already deployed contract.
func bindERC721OrdersFeature(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721OrdersFeatureABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721OrdersFeature *ERC721OrdersFeatureRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC721OrdersFeature.Contract.ERC721OrdersFeatureCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721OrdersFeature *ERC721OrdersFeatureRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721OrdersFeature.Contract.ERC721OrdersFeatureTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721OrdersFeature *ERC721OrdersFeatureRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721OrdersFeature.Contract.ERC721OrdersFeatureTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721OrdersFeature *ERC721OrdersFeatureCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC721OrdersFeature.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721OrdersFeature *ERC721OrdersFeatureTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721OrdersFeature.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721OrdersFeature *ERC721OrdersFeatureTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721OrdersFeature.Contract.contract.Transact(opts, method, params...)
}

// EIP712DOMAINSEPARATOR is a free data retrieval call binding the contract method 0xdab400f3.
//
// Solidity: function EIP712_DOMAIN_SEPARATOR() view returns(bytes32)
func (_ERC721OrdersFeature *ERC721OrdersFeatureCaller) EIP712DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC721OrdersFeature.contract.Call(opts, &out, "EIP712_DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EIP712DOMAINSEPARATOR is a free data retrieval call binding the contract method 0xdab400f3.
//
// Solidity: function EIP712_DOMAIN_SEPARATOR() view returns(bytes32)
func (_ERC721OrdersFeature *ERC721OrdersFeatureSession) EIP712DOMAINSEPARATOR() ([32]byte, error) {
	return _ERC721OrdersFeature.Contract.EIP712DOMAINSEPARATOR(&_ERC721OrdersFeature.CallOpts)
}

// EIP712DOMAINSEPARATOR is a free data retrieval call binding the contract method 0xdab400f3.
//
// Solidity: function EIP712_DOMAIN_SEPARATOR() view returns(bytes32)
func (_ERC721OrdersFeature *ERC721OrdersFeatureCallerSession) EIP712DOMAINSEPARATOR() ([32]byte, error) {
	return _ERC721OrdersFeature.Contract.EIP712DOMAINSEPARATOR(&_ERC721OrdersFeature.CallOpts)
}

// FEATURENAME is a free data retrieval call binding the contract method 0x6ae4b4f7.
//
// Solidity: function FEATURE_NAME() view returns(string)
func (_ERC721OrdersFeature *ERC721OrdersFeatureCaller) FEATURENAME(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC721OrdersFeature.contract.Call(opts, &out, "FEATURE_NAME")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// FEATURENAME is a free data retrieval call binding the contract method 0x6ae4b4f7.
//
// Solidity: function FEATURE_NAME() view returns(string)
func (_ERC721OrdersFeature *ERC721OrdersFeatureSession) FEATURENAME() (string, error) {
	return _ERC721OrdersFeature.Contract.FEATURENAME(&_ERC721OrdersFeature.CallOpts)
}

// FEATURENAME is a free data retrieval call binding the contract method 0x6ae4b4f7.
//
// Solidity: function FEATURE_NAME() view returns(string)
func (_ERC721OrdersFeature *ERC721OrdersFeatureCallerSession) FEATURENAME() (string, error) {
	return _ERC721OrdersFeature.Contract.FEATURENAME(&_ERC721OrdersFeature.CallOpts)
}

// FEATUREVERSION is a free data retrieval call binding the contract method 0x031b905c.
//
// Solidity: function FEATURE_VERSION() view returns(uint256)
func (_ERC721OrdersFeature *ERC721OrdersFeatureCaller) FEATUREVERSION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC721OrdersFeature.contract.Call(opts, &out, "FEATURE_VERSION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FEATUREVERSION is a free data retrieval call binding the contract method 0x031b905c.
//
// Solidity: function FEATURE_VERSION() view returns(uint256)
func (_ERC721OrdersFeature *ERC721OrdersFeatureSession) FEATUREVERSION() (*big.Int, error) {
	return _ERC721OrdersFeature.Contract.FEATUREVERSION(&_ERC721OrdersFeature.CallOpts)
}

// FEATUREVERSION is a free data retrieval call binding the contract method 0x031b905c.
//
// Solidity: function FEATURE_VERSION() view returns(uint256)
func (_ERC721OrdersFeature *ERC721OrdersFeatureCallerSession) FEATUREVERSION() (*big.Int, error) {
	return _ERC721OrdersFeature.Contract.FEATUREVERSION(&_ERC721OrdersFeature.CallOpts)
}

// GetERC721OrderHash is a free data retrieval call binding the contract method 0xb73a6027.
//
// Solidity: function getERC721OrderHash((uint8,address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[]) order) view returns(bytes32 orderHash)
func (_ERC721OrdersFeature *ERC721OrdersFeatureCaller) GetERC721OrderHash(opts *bind.CallOpts, order LibNFTOrderERC721Order) ([32]byte, error) {
	var out []interface{}
	err := _ERC721OrdersFeature.contract.Call(opts, &out, "getERC721OrderHash", order)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetERC721OrderHash is a free data retrieval call binding the contract method 0xb73a6027.
//
// Solidity: function getERC721OrderHash((uint8,address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[]) order) view returns(bytes32 orderHash)
func (_ERC721OrdersFeature *ERC721OrdersFeatureSession) GetERC721OrderHash(order LibNFTOrderERC721Order) ([32]byte, error) {
	return _ERC721OrdersFeature.Contract.GetERC721OrderHash(&_ERC721OrdersFeature.CallOpts, order)
}

// GetERC721OrderHash is a free data retrieval call binding the contract method 0xb73a6027.
//
// Solidity: function getERC721OrderHash((uint8,address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[]) order) view returns(bytes32 orderHash)
func (_ERC721OrdersFeature *ERC721OrdersFeatureCallerSession) GetERC721OrderHash(order LibNFTOrderERC721Order) ([32]byte, error) {
	return _ERC721OrdersFeature.Contract.GetERC721OrderHash(&_ERC721OrdersFeature.CallOpts, order)
}

// GetERC721OrderStatus is a free data retrieval call binding the contract method 0xfbc4a518.
//
// Solidity: function getERC721OrderStatus((uint8,address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[]) order) view returns(uint8 status)
func (_ERC721OrdersFeature *ERC721OrdersFeatureCaller) GetERC721OrderStatus(opts *bind.CallOpts, order LibNFTOrderERC721Order) (uint8, error) {
	var out []interface{}
	err := _ERC721OrdersFeature.contract.Call(opts, &out, "getERC721OrderStatus", order)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetERC721OrderStatus is a free data retrieval call binding the contract method 0xfbc4a518.
//
// Solidity: function getERC721OrderStatus((uint8,address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[]) order) view returns(uint8 status)
func (_ERC721OrdersFeature *ERC721OrdersFeatureSession) GetERC721OrderStatus(order LibNFTOrderERC721Order) (uint8, error) {
	return _ERC721OrdersFeature.Contract.GetERC721OrderStatus(&_ERC721OrdersFeature.CallOpts, order)
}

// GetERC721OrderStatus is a free data retrieval call binding the contract method 0xfbc4a518.
//
// Solidity: function getERC721OrderStatus((uint8,address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[]) order) view returns(uint8 status)
func (_ERC721OrdersFeature *ERC721OrdersFeatureCallerSession) GetERC721OrderStatus(order LibNFTOrderERC721Order) (uint8, error) {
	return _ERC721OrdersFeature.Contract.GetERC721OrderStatus(&_ERC721OrdersFeature.CallOpts, order)
}

// GetERC721OrderStatusBitVector is a free data retrieval call binding the contract method 0x030b2730.
//
// Solidity: function getERC721OrderStatusBitVector(address maker, uint248 nonceRange) view returns(uint256 bitVector)
func (_ERC721OrdersFeature *ERC721OrdersFeatureCaller) GetERC721OrderStatusBitVector(opts *bind.CallOpts, maker common.Address, nonceRange *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ERC721OrdersFeature.contract.Call(opts, &out, "getERC721OrderStatusBitVector", maker, nonceRange)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetERC721OrderStatusBitVector is a free data retrieval call binding the contract method 0x030b2730.
//
// Solidity: function getERC721OrderStatusBitVector(address maker, uint248 nonceRange) view returns(uint256 bitVector)
func (_ERC721OrdersFeature *ERC721OrdersFeatureSession) GetERC721OrderStatusBitVector(maker common.Address, nonceRange *big.Int) (*big.Int, error) {
	return _ERC721OrdersFeature.Contract.GetERC721OrderStatusBitVector(&_ERC721OrdersFeature.CallOpts, maker, nonceRange)
}

// GetERC721OrderStatusBitVector is a free data retrieval call binding the contract method 0x030b2730.
//
// Solidity: function getERC721OrderStatusBitVector(address maker, uint248 nonceRange) view returns(uint256 bitVector)
func (_ERC721OrdersFeature *ERC721OrdersFeatureCallerSession) GetERC721OrderStatusBitVector(maker common.Address, nonceRange *big.Int) (*big.Int, error) {
	return _ERC721OrdersFeature.Contract.GetERC721OrderStatusBitVector(&_ERC721OrdersFeature.CallOpts, maker, nonceRange)
}

// ValidateERC721OrderProperties is a free data retrieval call binding the contract method 0x4a13d797.
//
// Solidity: function validateERC721OrderProperties((uint8,address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[]) order, uint256 erc721TokenId) view returns()
func (_ERC721OrdersFeature *ERC721OrdersFeatureCaller) ValidateERC721OrderProperties(opts *bind.CallOpts, order LibNFTOrderERC721Order, erc721TokenId *big.Int) error {
	var out []interface{}
	err := _ERC721OrdersFeature.contract.Call(opts, &out, "validateERC721OrderProperties", order, erc721TokenId)

	if err != nil {
		return err
	}

	return err

}

// ValidateERC721OrderProperties is a free data retrieval call binding the contract method 0x4a13d797.
//
// Solidity: function validateERC721OrderProperties((uint8,address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[]) order, uint256 erc721TokenId) view returns()
func (_ERC721OrdersFeature *ERC721OrdersFeatureSession) ValidateERC721OrderProperties(order LibNFTOrderERC721Order, erc721TokenId *big.Int) error {
	return _ERC721OrdersFeature.Contract.ValidateERC721OrderProperties(&_ERC721OrdersFeature.CallOpts, order, erc721TokenId)
}

// ValidateERC721OrderProperties is a free data retrieval call binding the contract method 0x4a13d797.
//
// Solidity: function validateERC721OrderProperties((uint8,address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[]) order, uint256 erc721TokenId) view returns()
func (_ERC721OrdersFeature *ERC721OrdersFeatureCallerSession) ValidateERC721OrderProperties(order LibNFTOrderERC721Order, erc721TokenId *big.Int) error {
	return _ERC721OrdersFeature.Contract.ValidateERC721OrderProperties(&_ERC721OrdersFeature.CallOpts, order, erc721TokenId)
}

// ValidateERC721OrderSignature is a free data retrieval call binding the contract method 0xd1ca183b.
//
// Solidity: function validateERC721OrderSignature((uint8,address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[]) order, (uint8,uint8,bytes32,bytes32) signature) view returns()
func (_ERC721OrdersFeature *ERC721OrdersFeatureCaller) ValidateERC721OrderSignature(opts *bind.CallOpts, order LibNFTOrderERC721Order, signature LibSignatureSignature) error {
	var out []interface{}
	err := _ERC721OrdersFeature.contract.Call(opts, &out, "validateERC721OrderSignature", order, signature)

	if err != nil {
		return err
	}

	return err

}

// ValidateERC721OrderSignature is a free data retrieval call binding the contract method 0xd1ca183b.
//
// Solidity: function validateERC721OrderSignature((uint8,address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[]) order, (uint8,uint8,bytes32,bytes32) signature) view returns()
func (_ERC721OrdersFeature *ERC721OrdersFeatureSession) ValidateERC721OrderSignature(order LibNFTOrderERC721Order, signature LibSignatureSignature) error {
	return _ERC721OrdersFeature.Contract.ValidateERC721OrderSignature(&_ERC721OrdersFeature.CallOpts, order, signature)
}

// ValidateERC721OrderSignature is a free data retrieval call binding the contract method 0xd1ca183b.
//
// Solidity: function validateERC721OrderSignature((uint8,address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[]) order, (uint8,uint8,bytes32,bytes32) signature) view returns()
func (_ERC721OrdersFeature *ERC721OrdersFeatureCallerSession) ValidateERC721OrderSignature(order LibNFTOrderERC721Order, signature LibSignatureSignature) error {
	return _ERC721OrdersFeature.Contract.ValidateERC721OrderSignature(&_ERC721OrdersFeature.CallOpts, order, signature)
}

// BuyERC721 is a paid mutator transaction binding the contract method 0x60018b41.
//
// Solidity: function _buyERC721((uint8,address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[]) sellOrder, (uint8,uint8,bytes32,bytes32) signature, uint256 ethAvailable, bytes takerCallbackData) payable returns()
func (_ERC721OrdersFeature *ERC721OrdersFeatureTransactor) BuyERC721(opts *bind.TransactOpts, sellOrder LibNFTOrderERC721Order, signature LibSignatureSignature, ethAvailable *big.Int, takerCallbackData []byte) (*types.Transaction, error) {
	return _ERC721OrdersFeature.contract.Transact(opts, "_buyERC721", sellOrder, signature, ethAvailable, takerCallbackData)
}

// BuyERC721 is a paid mutator transaction binding the contract method 0x60018b41.
//
// Solidity: function _buyERC721((uint8,address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[]) sellOrder, (uint8,uint8,bytes32,bytes32) signature, uint256 ethAvailable, bytes takerCallbackData) payable returns()
func (_ERC721OrdersFeature *ERC721OrdersFeatureSession) BuyERC721(sellOrder LibNFTOrderERC721Order, signature LibSignatureSignature, ethAvailable *big.Int, takerCallbackData []byte) (*types.Transaction, error) {
	return _ERC721OrdersFeature.Contract.BuyERC721(&_ERC721OrdersFeature.TransactOpts, sellOrder, signature, ethAvailable, takerCallbackData)
}

// BuyERC721 is a paid mutator transaction binding the contract method 0x60018b41.
//
// Solidity: function _buyERC721((uint8,address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[]) sellOrder, (uint8,uint8,bytes32,bytes32) signature, uint256 ethAvailable, bytes takerCallbackData) payable returns()
func (_ERC721OrdersFeature *ERC721OrdersFeatureTransactorSession) BuyERC721(sellOrder LibNFTOrderERC721Order, signature LibSignatureSignature, ethAvailable *big.Int, takerCallbackData []byte) (*types.Transaction, error) {
	return _ERC721OrdersFeature.Contract.BuyERC721(&_ERC721OrdersFeature.TransactOpts, sellOrder, signature, ethAvailable, takerCallbackData)
}

// BatchBuyERC721s is a paid mutator transaction binding the contract method 0xeae93ee7.
//
// Solidity: function batchBuyERC721s((uint8,address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[])[] sellOrders, (uint8,uint8,bytes32,bytes32)[] signatures, bytes[] callbackData, bool revertIfIncomplete) payable returns(bool[] successes)
func (_ERC721OrdersFeature *ERC721OrdersFeatureTransactor) BatchBuyERC721s(opts *bind.TransactOpts, sellOrders []LibNFTOrderERC721Order, signatures []LibSignatureSignature, callbackData [][]byte, revertIfIncomplete bool) (*types.Transaction, error) {
	return _ERC721OrdersFeature.contract.Transact(opts, "batchBuyERC721s", sellOrders, signatures, callbackData, revertIfIncomplete)
}

// BatchBuyERC721s is a paid mutator transaction binding the contract method 0xeae93ee7.
//
// Solidity: function batchBuyERC721s((uint8,address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[])[] sellOrders, (uint8,uint8,bytes32,bytes32)[] signatures, bytes[] callbackData, bool revertIfIncomplete) payable returns(bool[] successes)
func (_ERC721OrdersFeature *ERC721OrdersFeatureSession) BatchBuyERC721s(sellOrders []LibNFTOrderERC721Order, signatures []LibSignatureSignature, callbackData [][]byte, revertIfIncomplete bool) (*types.Transaction, error) {
	return _ERC721OrdersFeature.Contract.BatchBuyERC721s(&_ERC721OrdersFeature.TransactOpts, sellOrders, signatures, callbackData, revertIfIncomplete)
}

// BatchBuyERC721s is a paid mutator transaction binding the contract method 0xeae93ee7.
//
// Solidity: function batchBuyERC721s((uint8,address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[])[] sellOrders, (uint8,uint8,bytes32,bytes32)[] signatures, bytes[] callbackData, bool revertIfIncomplete) payable returns(bool[] successes)
func (_ERC721OrdersFeature *ERC721OrdersFeatureTransactorSession) BatchBuyERC721s(sellOrders []LibNFTOrderERC721Order, signatures []LibSignatureSignature, callbackData [][]byte, revertIfIncomplete bool) (*types.Transaction, error) {
	return _ERC721OrdersFeature.Contract.BatchBuyERC721s(&_ERC721OrdersFeature.TransactOpts, sellOrders, signatures, callbackData, revertIfIncomplete)
}

// BatchCancelERC721Orders is a paid mutator transaction binding the contract method 0x86219940.
//
// Solidity: function batchCancelERC721Orders(uint256[] orderNonces) returns()
func (_ERC721OrdersFeature *ERC721OrdersFeatureTransactor) BatchCancelERC721Orders(opts *bind.TransactOpts, orderNonces []*big.Int) (*types.Transaction, error) {
	return _ERC721OrdersFeature.contract.Transact(opts, "batchCancelERC721Orders", orderNonces)
}

// BatchCancelERC721Orders is a paid mutator transaction binding the contract method 0x86219940.
//
// Solidity: function batchCancelERC721Orders(uint256[] orderNonces) returns()
func (_ERC721OrdersFeature *ERC721OrdersFeatureSession) BatchCancelERC721Orders(orderNonces []*big.Int) (*types.Transaction, error) {
	return _ERC721OrdersFeature.Contract.BatchCancelERC721Orders(&_ERC721OrdersFeature.TransactOpts, orderNonces)
}

// BatchCancelERC721Orders is a paid mutator transaction binding the contract method 0x86219940.
//
// Solidity: function batchCancelERC721Orders(uint256[] orderNonces) returns()
func (_ERC721OrdersFeature *ERC721OrdersFeatureTransactorSession) BatchCancelERC721Orders(orderNonces []*big.Int) (*types.Transaction, error) {
	return _ERC721OrdersFeature.Contract.BatchCancelERC721Orders(&_ERC721OrdersFeature.TransactOpts, orderNonces)
}

// BatchMatchERC721Orders is a paid mutator transaction binding the contract method 0x7da9e2cf.
//
// Solidity: function batchMatchERC721Orders((uint8,address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[])[] sellOrders, (uint8,address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[])[] buyOrders, (uint8,uint8,bytes32,bytes32)[] sellOrderSignatures, (uint8,uint8,bytes32,bytes32)[] buyOrderSignatures) returns(uint256[] profits, bool[] successes)
func (_ERC721OrdersFeature *ERC721OrdersFeatureTransactor) BatchMatchERC721Orders(opts *bind.TransactOpts, sellOrders []LibNFTOrderERC721Order, buyOrders []LibNFTOrderERC721Order, sellOrderSignatures []LibSignatureSignature, buyOrderSignatures []LibSignatureSignature) (*types.Transaction, error) {
	return _ERC721OrdersFeature.contract.Transact(opts, "batchMatchERC721Orders", sellOrders, buyOrders, sellOrderSignatures, buyOrderSignatures)
}

// BatchMatchERC721Orders is a paid mutator transaction binding the contract method 0x7da9e2cf.
//
// Solidity: function batchMatchERC721Orders((uint8,address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[])[] sellOrders, (uint8,address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[])[] buyOrders, (uint8,uint8,bytes32,bytes32)[] sellOrderSignatures, (uint8,uint8,bytes32,bytes32)[] buyOrderSignatures) returns(uint256[] profits, bool[] successes)
func (_ERC721OrdersFeature *ERC721OrdersFeatureSession) BatchMatchERC721Orders(sellOrders []LibNFTOrderERC721Order, buyOrders []LibNFTOrderERC721Order, sellOrderSignatures []LibSignatureSignature, buyOrderSignatures []LibSignatureSignature) (*types.Transaction, error) {
	return _ERC721OrdersFeature.Contract.BatchMatchERC721Orders(&_ERC721OrdersFeature.TransactOpts, sellOrders, buyOrders, sellOrderSignatures, buyOrderSignatures)
}

// BatchMatchERC721Orders is a paid mutator transaction binding the contract method 0x7da9e2cf.
//
// Solidity: function batchMatchERC721Orders((uint8,address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[])[] sellOrders, (uint8,address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[])[] buyOrders, (uint8,uint8,bytes32,bytes32)[] sellOrderSignatures, (uint8,uint8,bytes32,bytes32)[] buyOrderSignatures) returns(uint256[] profits, bool[] successes)
func (_ERC721OrdersFeature *ERC721OrdersFeatureTransactorSession) BatchMatchERC721Orders(sellOrders []LibNFTOrderERC721Order, buyOrders []LibNFTOrderERC721Order, sellOrderSignatures []LibSignatureSignature, buyOrderSignatures []LibSignatureSignature) (*types.Transaction, error) {
	return _ERC721OrdersFeature.Contract.BatchMatchERC721Orders(&_ERC721OrdersFeature.TransactOpts, sellOrders, buyOrders, sellOrderSignatures, buyOrderSignatures)
}

// BuyERC7212 is a paid mutator transaction binding the contract method 0xfbee349d.
//
// Solidity: function buyERC721((uint8,address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[]) sellOrder, (uint8,uint8,bytes32,bytes32) signature, bytes callbackData) payable returns()
func (_ERC721OrdersFeature *ERC721OrdersFeatureTransactor) BuyERC7212(opts *bind.TransactOpts, sellOrder LibNFTOrderERC721Order, signature LibSignatureSignature, callbackData []byte) (*types.Transaction, error) {
	return _ERC721OrdersFeature.contract.Transact(opts, "buyERC721", sellOrder, signature, callbackData)
}

// BuyERC7212 is a paid mutator transaction binding the contract method 0xfbee349d.
//
// Solidity: function buyERC721((uint8,address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[]) sellOrder, (uint8,uint8,bytes32,bytes32) signature, bytes callbackData) payable returns()
func (_ERC721OrdersFeature *ERC721OrdersFeatureSession) BuyERC7212(sellOrder LibNFTOrderERC721Order, signature LibSignatureSignature, callbackData []byte) (*types.Transaction, error) {
	return _ERC721OrdersFeature.Contract.BuyERC7212(&_ERC721OrdersFeature.TransactOpts, sellOrder, signature, callbackData)
}

// BuyERC7212 is a paid mutator transaction binding the contract method 0xfbee349d.
//
// Solidity: function buyERC721((uint8,address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[]) sellOrder, (uint8,uint8,bytes32,bytes32) signature, bytes callbackData) payable returns()
func (_ERC721OrdersFeature *ERC721OrdersFeatureTransactorSession) BuyERC7212(sellOrder LibNFTOrderERC721Order, signature LibSignatureSignature, callbackData []byte) (*types.Transaction, error) {
	return _ERC721OrdersFeature.Contract.BuyERC7212(&_ERC721OrdersFeature.TransactOpts, sellOrder, signature, callbackData)
}

// CancelERC721Order is a paid mutator transaction binding the contract method 0xbe167b9d.
//
// Solidity: function cancelERC721Order(uint256 orderNonce) returns()
func (_ERC721OrdersFeature *ERC721OrdersFeatureTransactor) CancelERC721Order(opts *bind.TransactOpts, orderNonce *big.Int) (*types.Transaction, error) {
	return _ERC721OrdersFeature.contract.Transact(opts, "cancelERC721Order", orderNonce)
}

// CancelERC721Order is a paid mutator transaction binding the contract method 0xbe167b9d.
//
// Solidity: function cancelERC721Order(uint256 orderNonce) returns()
func (_ERC721OrdersFeature *ERC721OrdersFeatureSession) CancelERC721Order(orderNonce *big.Int) (*types.Transaction, error) {
	return _ERC721OrdersFeature.Contract.CancelERC721Order(&_ERC721OrdersFeature.TransactOpts, orderNonce)
}

// CancelERC721Order is a paid mutator transaction binding the contract method 0xbe167b9d.
//
// Solidity: function cancelERC721Order(uint256 orderNonce) returns()
func (_ERC721OrdersFeature *ERC721OrdersFeatureTransactorSession) CancelERC721Order(orderNonce *big.Int) (*types.Transaction, error) {
	return _ERC721OrdersFeature.Contract.CancelERC721Order(&_ERC721OrdersFeature.TransactOpts, orderNonce)
}

// MatchERC721Orders is a paid mutator transaction binding the contract method 0x0d8261eb.
//
// Solidity: function matchERC721Orders((uint8,address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[]) sellOrder, (uint8,address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[]) buyOrder, (uint8,uint8,bytes32,bytes32) sellOrderSignature, (uint8,uint8,bytes32,bytes32) buyOrderSignature) returns(uint256 profit)
func (_ERC721OrdersFeature *ERC721OrdersFeatureTransactor) MatchERC721Orders(opts *bind.TransactOpts, sellOrder LibNFTOrderERC721Order, buyOrder LibNFTOrderERC721Order, sellOrderSignature LibSignatureSignature, buyOrderSignature LibSignatureSignature) (*types.Transaction, error) {
	return _ERC721OrdersFeature.contract.Transact(opts, "matchERC721Orders", sellOrder, buyOrder, sellOrderSignature, buyOrderSignature)
}

// MatchERC721Orders is a paid mutator transaction binding the contract method 0x0d8261eb.
//
// Solidity: function matchERC721Orders((uint8,address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[]) sellOrder, (uint8,address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[]) buyOrder, (uint8,uint8,bytes32,bytes32) sellOrderSignature, (uint8,uint8,bytes32,bytes32) buyOrderSignature) returns(uint256 profit)
func (_ERC721OrdersFeature *ERC721OrdersFeatureSession) MatchERC721Orders(sellOrder LibNFTOrderERC721Order, buyOrder LibNFTOrderERC721Order, sellOrderSignature LibSignatureSignature, buyOrderSignature LibSignatureSignature) (*types.Transaction, error) {
	return _ERC721OrdersFeature.Contract.MatchERC721Orders(&_ERC721OrdersFeature.TransactOpts, sellOrder, buyOrder, sellOrderSignature, buyOrderSignature)
}

// MatchERC721Orders is a paid mutator transaction binding the contract method 0x0d8261eb.
//
// Solidity: function matchERC721Orders((uint8,address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[]) sellOrder, (uint8,address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[]) buyOrder, (uint8,uint8,bytes32,bytes32) sellOrderSignature, (uint8,uint8,bytes32,bytes32) buyOrderSignature) returns(uint256 profit)
func (_ERC721OrdersFeature *ERC721OrdersFeatureTransactorSession) MatchERC721Orders(sellOrder LibNFTOrderERC721Order, buyOrder LibNFTOrderERC721Order, sellOrderSignature LibSignatureSignature, buyOrderSignature LibSignatureSignature) (*types.Transaction, error) {
	return _ERC721OrdersFeature.Contract.MatchERC721Orders(&_ERC721OrdersFeature.TransactOpts, sellOrder, buyOrder, sellOrderSignature, buyOrderSignature)
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns(bytes4 success)
func (_ERC721OrdersFeature *ERC721OrdersFeatureTransactor) Migrate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721OrdersFeature.contract.Transact(opts, "migrate")
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns(bytes4 success)
func (_ERC721OrdersFeature *ERC721OrdersFeatureSession) Migrate() (*types.Transaction, error) {
	return _ERC721OrdersFeature.Contract.Migrate(&_ERC721OrdersFeature.TransactOpts)
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns(bytes4 success)
func (_ERC721OrdersFeature *ERC721OrdersFeatureTransactorSession) Migrate() (*types.Transaction, error) {
	return _ERC721OrdersFeature.Contract.Migrate(&_ERC721OrdersFeature.TransactOpts)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address , uint256 tokenId, bytes data) returns(bytes4 success)
func (_ERC721OrdersFeature *ERC721OrdersFeatureTransactor) OnERC721Received(opts *bind.TransactOpts, operator common.Address, arg1 common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC721OrdersFeature.contract.Transact(opts, "onERC721Received", operator, arg1, tokenId, data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address , uint256 tokenId, bytes data) returns(bytes4 success)
func (_ERC721OrdersFeature *ERC721OrdersFeatureSession) OnERC721Received(operator common.Address, arg1 common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC721OrdersFeature.Contract.OnERC721Received(&_ERC721OrdersFeature.TransactOpts, operator, arg1, tokenId, data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address , uint256 tokenId, bytes data) returns(bytes4 success)
func (_ERC721OrdersFeature *ERC721OrdersFeatureTransactorSession) OnERC721Received(operator common.Address, arg1 common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC721OrdersFeature.Contract.OnERC721Received(&_ERC721OrdersFeature.TransactOpts, operator, arg1, tokenId, data)
}

// PreSignERC721Order is a paid mutator transaction binding the contract method 0x462103af.
//
// Solidity: function preSignERC721Order((uint8,address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[]) order) returns()
func (_ERC721OrdersFeature *ERC721OrdersFeatureTransactor) PreSignERC721Order(opts *bind.TransactOpts, order LibNFTOrderERC721Order) (*types.Transaction, error) {
	return _ERC721OrdersFeature.contract.Transact(opts, "preSignERC721Order", order)
}

// PreSignERC721Order is a paid mutator transaction binding the contract method 0x462103af.
//
// Solidity: function preSignERC721Order((uint8,address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[]) order) returns()
func (_ERC721OrdersFeature *ERC721OrdersFeatureSession) PreSignERC721Order(order LibNFTOrderERC721Order) (*types.Transaction, error) {
	return _ERC721OrdersFeature.Contract.PreSignERC721Order(&_ERC721OrdersFeature.TransactOpts, order)
}

// PreSignERC721Order is a paid mutator transaction binding the contract method 0x462103af.
//
// Solidity: function preSignERC721Order((uint8,address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[]) order) returns()
func (_ERC721OrdersFeature *ERC721OrdersFeatureTransactorSession) PreSignERC721Order(order LibNFTOrderERC721Order) (*types.Transaction, error) {
	return _ERC721OrdersFeature.Contract.PreSignERC721Order(&_ERC721OrdersFeature.TransactOpts, order)
}

// SellERC721 is a paid mutator transaction binding the contract method 0xafde1b3c.
//
// Solidity: function sellERC721((uint8,address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[]) buyOrder, (uint8,uint8,bytes32,bytes32) signature, uint256 erc721TokenId, bool unwrapNativeToken, bytes callbackData) returns()
func (_ERC721OrdersFeature *ERC721OrdersFeatureTransactor) SellERC721(opts *bind.TransactOpts, buyOrder LibNFTOrderERC721Order, signature LibSignatureSignature, erc721TokenId *big.Int, unwrapNativeToken bool, callbackData []byte) (*types.Transaction, error) {
	return _ERC721OrdersFeature.contract.Transact(opts, "sellERC721", buyOrder, signature, erc721TokenId, unwrapNativeToken, callbackData)
}

// SellERC721 is a paid mutator transaction binding the contract method 0xafde1b3c.
//
// Solidity: function sellERC721((uint8,address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[]) buyOrder, (uint8,uint8,bytes32,bytes32) signature, uint256 erc721TokenId, bool unwrapNativeToken, bytes callbackData) returns()
func (_ERC721OrdersFeature *ERC721OrdersFeatureSession) SellERC721(buyOrder LibNFTOrderERC721Order, signature LibSignatureSignature, erc721TokenId *big.Int, unwrapNativeToken bool, callbackData []byte) (*types.Transaction, error) {
	return _ERC721OrdersFeature.Contract.SellERC721(&_ERC721OrdersFeature.TransactOpts, buyOrder, signature, erc721TokenId, unwrapNativeToken, callbackData)
}

// SellERC721 is a paid mutator transaction binding the contract method 0xafde1b3c.
//
// Solidity: function sellERC721((uint8,address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[]) buyOrder, (uint8,uint8,bytes32,bytes32) signature, uint256 erc721TokenId, bool unwrapNativeToken, bytes callbackData) returns()
func (_ERC721OrdersFeature *ERC721OrdersFeatureTransactorSession) SellERC721(buyOrder LibNFTOrderERC721Order, signature LibSignatureSignature, erc721TokenId *big.Int, unwrapNativeToken bool, callbackData []byte) (*types.Transaction, error) {
	return _ERC721OrdersFeature.Contract.SellERC721(&_ERC721OrdersFeature.TransactOpts, buyOrder, signature, erc721TokenId, unwrapNativeToken, callbackData)
}

// ERC721OrdersFeatureERC721OrderCancelledIterator is returned from FilterERC721OrderCancelled and is used to iterate over the raw logs and unpacked data for ERC721OrderCancelled events raised by the ERC721OrdersFeature contract.
type ERC721OrdersFeatureERC721OrderCancelledIterator struct {
	Event *ERC721OrdersFeatureERC721OrderCancelled // Event containing the contract specifics and raw log

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
func (it *ERC721OrdersFeatureERC721OrderCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721OrdersFeatureERC721OrderCancelled)
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
		it.Event = new(ERC721OrdersFeatureERC721OrderCancelled)
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
func (it *ERC721OrdersFeatureERC721OrderCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721OrdersFeatureERC721OrderCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721OrdersFeatureERC721OrderCancelled represents a ERC721OrderCancelled event raised by the ERC721OrdersFeature contract.
type ERC721OrdersFeatureERC721OrderCancelled struct {
	Maker common.Address
	Nonce *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterERC721OrderCancelled is a free log retrieval operation binding the contract event 0xa015ad2dc32f266993958a0fd9884c746b971b254206f3478bc43e2f125c7b9e.
//
// Solidity: event ERC721OrderCancelled(address maker, uint256 nonce)
func (_ERC721OrdersFeature *ERC721OrdersFeatureFilterer) FilterERC721OrderCancelled(opts *bind.FilterOpts) (*ERC721OrdersFeatureERC721OrderCancelledIterator, error) {

	logs, sub, err := _ERC721OrdersFeature.contract.FilterLogs(opts, "ERC721OrderCancelled")
	if err != nil {
		return nil, err
	}
	return &ERC721OrdersFeatureERC721OrderCancelledIterator{contract: _ERC721OrdersFeature.contract, event: "ERC721OrderCancelled", logs: logs, sub: sub}, nil
}

// WatchERC721OrderCancelled is a free log subscription operation binding the contract event 0xa015ad2dc32f266993958a0fd9884c746b971b254206f3478bc43e2f125c7b9e.
//
// Solidity: event ERC721OrderCancelled(address maker, uint256 nonce)
func (_ERC721OrdersFeature *ERC721OrdersFeatureFilterer) WatchERC721OrderCancelled(opts *bind.WatchOpts, sink chan<- *ERC721OrdersFeatureERC721OrderCancelled) (event.Subscription, error) {

	logs, sub, err := _ERC721OrdersFeature.contract.WatchLogs(opts, "ERC721OrderCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721OrdersFeatureERC721OrderCancelled)
				if err := _ERC721OrdersFeature.contract.UnpackLog(event, "ERC721OrderCancelled", log); err != nil {
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

// ParseERC721OrderCancelled is a log parse operation binding the contract event 0xa015ad2dc32f266993958a0fd9884c746b971b254206f3478bc43e2f125c7b9e.
//
// Solidity: event ERC721OrderCancelled(address maker, uint256 nonce)
func (_ERC721OrdersFeature *ERC721OrdersFeatureFilterer) ParseERC721OrderCancelled(log types.Log) (*ERC721OrdersFeatureERC721OrderCancelled, error) {
	event := new(ERC721OrdersFeatureERC721OrderCancelled)
	if err := _ERC721OrdersFeature.contract.UnpackLog(event, "ERC721OrderCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC721OrdersFeatureERC721OrderFilledIterator is returned from FilterERC721OrderFilled and is used to iterate over the raw logs and unpacked data for ERC721OrderFilled events raised by the ERC721OrdersFeature contract.
type ERC721OrdersFeatureERC721OrderFilledIterator struct {
	Event *ERC721OrdersFeatureERC721OrderFilled // Event containing the contract specifics and raw log

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
func (it *ERC721OrdersFeatureERC721OrderFilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721OrdersFeatureERC721OrderFilled)
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
		it.Event = new(ERC721OrdersFeatureERC721OrderFilled)
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
func (it *ERC721OrdersFeatureERC721OrderFilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721OrdersFeatureERC721OrderFilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721OrdersFeatureERC721OrderFilled represents a ERC721OrderFilled event raised by the ERC721OrdersFeature contract.
type ERC721OrdersFeatureERC721OrderFilled struct {
	Direction        uint8
	Maker            common.Address
	Taker            common.Address
	Nonce            *big.Int
	Erc20Token       common.Address
	Erc20TokenAmount *big.Int
	Erc721Token      common.Address
	Erc721TokenId    *big.Int
	Matcher          common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterERC721OrderFilled is a free log retrieval operation binding the contract event 0x50273fa02273cceea9cf085b42de5c8af60624140168bd71357db833535877af.
//
// Solidity: event ERC721OrderFilled(uint8 direction, address maker, address taker, uint256 nonce, address erc20Token, uint256 erc20TokenAmount, address erc721Token, uint256 erc721TokenId, address matcher)
func (_ERC721OrdersFeature *ERC721OrdersFeatureFilterer) FilterERC721OrderFilled(opts *bind.FilterOpts) (*ERC721OrdersFeatureERC721OrderFilledIterator, error) {

	logs, sub, err := _ERC721OrdersFeature.contract.FilterLogs(opts, "ERC721OrderFilled")
	if err != nil {
		return nil, err
	}
	return &ERC721OrdersFeatureERC721OrderFilledIterator{contract: _ERC721OrdersFeature.contract, event: "ERC721OrderFilled", logs: logs, sub: sub}, nil
}

// WatchERC721OrderFilled is a free log subscription operation binding the contract event 0x50273fa02273cceea9cf085b42de5c8af60624140168bd71357db833535877af.
//
// Solidity: event ERC721OrderFilled(uint8 direction, address maker, address taker, uint256 nonce, address erc20Token, uint256 erc20TokenAmount, address erc721Token, uint256 erc721TokenId, address matcher)
func (_ERC721OrdersFeature *ERC721OrdersFeatureFilterer) WatchERC721OrderFilled(opts *bind.WatchOpts, sink chan<- *ERC721OrdersFeatureERC721OrderFilled) (event.Subscription, error) {

	logs, sub, err := _ERC721OrdersFeature.contract.WatchLogs(opts, "ERC721OrderFilled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721OrdersFeatureERC721OrderFilled)
				if err := _ERC721OrdersFeature.contract.UnpackLog(event, "ERC721OrderFilled", log); err != nil {
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

// ParseERC721OrderFilled is a log parse operation binding the contract event 0x50273fa02273cceea9cf085b42de5c8af60624140168bd71357db833535877af.
//
// Solidity: event ERC721OrderFilled(uint8 direction, address maker, address taker, uint256 nonce, address erc20Token, uint256 erc20TokenAmount, address erc721Token, uint256 erc721TokenId, address matcher)
func (_ERC721OrdersFeature *ERC721OrdersFeatureFilterer) ParseERC721OrderFilled(log types.Log) (*ERC721OrdersFeatureERC721OrderFilled, error) {
	event := new(ERC721OrdersFeatureERC721OrderFilled)
	if err := _ERC721OrdersFeature.contract.UnpackLog(event, "ERC721OrderFilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC721OrdersFeatureERC721OrderPreSignedIterator is returned from FilterERC721OrderPreSigned and is used to iterate over the raw logs and unpacked data for ERC721OrderPreSigned events raised by the ERC721OrdersFeature contract.
type ERC721OrdersFeatureERC721OrderPreSignedIterator struct {
	Event *ERC721OrdersFeatureERC721OrderPreSigned // Event containing the contract specifics and raw log

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
func (it *ERC721OrdersFeatureERC721OrderPreSignedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721OrdersFeatureERC721OrderPreSigned)
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
		it.Event = new(ERC721OrdersFeatureERC721OrderPreSigned)
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
func (it *ERC721OrdersFeatureERC721OrderPreSignedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721OrdersFeatureERC721OrderPreSignedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721OrdersFeatureERC721OrderPreSigned represents a ERC721OrderPreSigned event raised by the ERC721OrdersFeature contract.
type ERC721OrdersFeatureERC721OrderPreSigned struct {
	Direction             uint8
	Maker                 common.Address
	Taker                 common.Address
	Expiry                *big.Int
	Nonce                 *big.Int
	Erc20Token            common.Address
	Erc20TokenAmount      *big.Int
	Fees                  []LibNFTOrderFee
	Erc721Token           common.Address
	Erc721TokenId         *big.Int
	Erc721TokenProperties []LibNFTOrderProperty
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterERC721OrderPreSigned is a free log retrieval operation binding the contract event 0x8c5d0c41fb16a7317a6c55ff7ba93d9d74f79e434fefa694e50d6028afbfa3f0.
//
// Solidity: event ERC721OrderPreSigned(uint8 direction, address maker, address taker, uint256 expiry, uint256 nonce, address erc20Token, uint256 erc20TokenAmount, (address,uint256,bytes)[] fees, address erc721Token, uint256 erc721TokenId, (address,bytes)[] erc721TokenProperties)
func (_ERC721OrdersFeature *ERC721OrdersFeatureFilterer) FilterERC721OrderPreSigned(opts *bind.FilterOpts) (*ERC721OrdersFeatureERC721OrderPreSignedIterator, error) {

	logs, sub, err := _ERC721OrdersFeature.contract.FilterLogs(opts, "ERC721OrderPreSigned")
	if err != nil {
		return nil, err
	}
	return &ERC721OrdersFeatureERC721OrderPreSignedIterator{contract: _ERC721OrdersFeature.contract, event: "ERC721OrderPreSigned", logs: logs, sub: sub}, nil
}

// WatchERC721OrderPreSigned is a free log subscription operation binding the contract event 0x8c5d0c41fb16a7317a6c55ff7ba93d9d74f79e434fefa694e50d6028afbfa3f0.
//
// Solidity: event ERC721OrderPreSigned(uint8 direction, address maker, address taker, uint256 expiry, uint256 nonce, address erc20Token, uint256 erc20TokenAmount, (address,uint256,bytes)[] fees, address erc721Token, uint256 erc721TokenId, (address,bytes)[] erc721TokenProperties)
func (_ERC721OrdersFeature *ERC721OrdersFeatureFilterer) WatchERC721OrderPreSigned(opts *bind.WatchOpts, sink chan<- *ERC721OrdersFeatureERC721OrderPreSigned) (event.Subscription, error) {

	logs, sub, err := _ERC721OrdersFeature.contract.WatchLogs(opts, "ERC721OrderPreSigned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721OrdersFeatureERC721OrderPreSigned)
				if err := _ERC721OrdersFeature.contract.UnpackLog(event, "ERC721OrderPreSigned", log); err != nil {
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

// ParseERC721OrderPreSigned is a log parse operation binding the contract event 0x8c5d0c41fb16a7317a6c55ff7ba93d9d74f79e434fefa694e50d6028afbfa3f0.
//
// Solidity: event ERC721OrderPreSigned(uint8 direction, address maker, address taker, uint256 expiry, uint256 nonce, address erc20Token, uint256 erc20TokenAmount, (address,uint256,bytes)[] fees, address erc721Token, uint256 erc721TokenId, (address,bytes)[] erc721TokenProperties)
func (_ERC721OrdersFeature *ERC721OrdersFeatureFilterer) ParseERC721OrderPreSigned(log types.Log) (*ERC721OrdersFeatureERC721OrderPreSigned, error) {
	event := new(ERC721OrdersFeatureERC721OrderPreSigned)
	if err := _ERC721OrdersFeature.contract.UnpackLog(event, "ERC721OrderPreSigned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
