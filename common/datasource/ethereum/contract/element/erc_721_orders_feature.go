// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package element

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

// INFTOrdersFeatureFee is an auto generated low-level Go binding around an user-defined struct.
type INFTOrdersFeatureFee struct {
	Recipient common.Address
	Amount    *big.Int
}

// LibNFTOrderFee is an auto generated low-level Go binding around an user-defined struct.
type LibNFTOrderFee struct {
	Recipient common.Address
	Amount    *big.Int
	FeeData   []byte
}

// LibNFTOrderNFTBuyOrder is an auto generated low-level Go binding around an user-defined struct.
type LibNFTOrderNFTBuyOrder struct {
	Maker            common.Address
	Taker            common.Address
	Expiry           *big.Int
	Nonce            *big.Int
	Erc20Token       common.Address
	Erc20TokenAmount *big.Int
	Fees             []LibNFTOrderFee
	Nft              common.Address
	NftId            *big.Int
	NftProperties    []LibNFTOrderProperty
}

// LibNFTOrderNFTSellOrder is an auto generated low-level Go binding around an user-defined struct.
type LibNFTOrderNFTSellOrder struct {
	Maker            common.Address
	Taker            common.Address
	Expiry           *big.Int
	Nonce            *big.Int
	Erc20Token       common.Address
	Erc20TokenAmount *big.Int
	Fees             []LibNFTOrderFee
	Nft              common.Address
	NftId            *big.Int
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
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIEtherToken\",\"name\":\"weth\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"erc20TokenAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structINFTOrdersFeature.Fee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"erc721Token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"erc721TokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721BuyOrderFilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"erc20TokenAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structLibNFTOrder.Fee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"erc721Token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"erc721TokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractIPropertyValidator\",\"name\":\"propertyValidator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"propertyData\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structLibNFTOrder.Property[]\",\"name\":\"nftProperties\",\"type\":\"tuple[]\"}],\"name\":\"ERC721BuyOrderPreSigned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"ERC721OrderCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"erc20TokenAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structINFTOrdersFeature.Fee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"erc721Token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"erc721TokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721SellOrderFilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"erc20TokenAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structLibNFTOrder.Fee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"erc721Token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"erc721TokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721SellOrderPreSigned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newHashNonce\",\"type\":\"uint256\"}],\"name\":\"HashNonceIncremented\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"takerData\",\"type\":\"bytes\"}],\"name\":\"TakerDataEmitted\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc20TokenAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Fee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"}],\"internalType\":\"structLibNFTOrder.NFTSellOrder[]\",\"name\":\"sellOrders\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumLibSignature.SignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structLibSignature.Signature[]\",\"name\":\"signatures\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"revertIfIncomplete\",\"type\":\"bool\"}],\"name\":\"batchBuyERC721s\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"successes\",\"type\":\"bool[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc20TokenAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Fee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"}],\"internalType\":\"structLibNFTOrder.NFTSellOrder[]\",\"name\":\"sellOrders\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumLibSignature.SignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structLibSignature.Signature[]\",\"name\":\"signatures\",\"type\":\"tuple[]\"},{\"internalType\":\"address[]\",\"name\":\"takers\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"takerDatas\",\"type\":\"bytes[]\"},{\"internalType\":\"bool\",\"name\":\"revertIfIncomplete\",\"type\":\"bool\"}],\"name\":\"batchBuyERC721sEx\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"successes\",\"type\":\"bool[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"orderNonces\",\"type\":\"uint256[]\"}],\"name\":\"batchCancelERC721Orders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc20TokenAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Fee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"}],\"internalType\":\"structLibNFTOrder.NFTSellOrder[]\",\"name\":\"sellOrders\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc20TokenAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Fee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractIPropertyValidator\",\"name\":\"propertyValidator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"propertyData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Property[]\",\"name\":\"nftProperties\",\"type\":\"tuple[]\"}],\"internalType\":\"structLibNFTOrder.NFTBuyOrder[]\",\"name\":\"buyOrders\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumLibSignature.SignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structLibSignature.Signature[]\",\"name\":\"sellOrderSignatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumLibSignature.SignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structLibSignature.Signature[]\",\"name\":\"buyOrderSignatures\",\"type\":\"tuple[]\"}],\"name\":\"batchMatchERC721Orders\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"profits\",\"type\":\"uint256[]\"},{\"internalType\":\"bool[]\",\"name\":\"successes\",\"type\":\"bool[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc20TokenAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Fee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"}],\"internalType\":\"structLibNFTOrder.NFTSellOrder\",\"name\":\"sellOrder\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumLibSignature.SignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structLibSignature.Signature\",\"name\":\"signature\",\"type\":\"tuple\"}],\"name\":\"buyERC721\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc20TokenAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Fee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"}],\"internalType\":\"structLibNFTOrder.NFTSellOrder\",\"name\":\"sellOrder\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumLibSignature.SignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structLibSignature.Signature\",\"name\":\"signature\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"takerData\",\"type\":\"bytes\"}],\"name\":\"buyERC721Ex\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc20TokenAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Fee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"}],\"internalType\":\"structLibNFTOrder.NFTSellOrder\",\"name\":\"sellOrder\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumLibSignature.SignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structLibSignature.Signature\",\"name\":\"signature\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"ethAvailable\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"takerData\",\"type\":\"bytes\"}],\"name\":\"buyERC721ExFromProxy\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc20TokenAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Fee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"}],\"internalType\":\"structLibNFTOrder.NFTSellOrder\",\"name\":\"sellOrder\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumLibSignature.SignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structLibSignature.Signature\",\"name\":\"signature\",\"type\":\"tuple\"}],\"name\":\"buyERC721FromProxy\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderNonce\",\"type\":\"uint256\"}],\"name\":\"cancelERC721Order\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc20TokenAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Fee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractIPropertyValidator\",\"name\":\"propertyValidator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"propertyData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Property[]\",\"name\":\"nftProperties\",\"type\":\"tuple[]\"}],\"internalType\":\"structLibNFTOrder.NFTBuyOrder\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"getERC721BuyOrderHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc20TokenAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Fee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractIPropertyValidator\",\"name\":\"propertyValidator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"propertyData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Property[]\",\"name\":\"nftProperties\",\"type\":\"tuple[]\"}],\"internalType\":\"structLibNFTOrder.NFTBuyOrder\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"getERC721BuyOrderStatus\",\"outputs\":[{\"internalType\":\"enumLibNFTOrder.OrderStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"uint248\",\"name\":\"nonceRange\",\"type\":\"uint248\"}],\"name\":\"getERC721OrderStatusBitVector\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc20TokenAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Fee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"}],\"internalType\":\"structLibNFTOrder.NFTSellOrder\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"getERC721SellOrderHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc20TokenAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Fee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"}],\"internalType\":\"structLibNFTOrder.NFTSellOrder\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"getERC721SellOrderStatus\",\"outputs\":[{\"internalType\":\"enumLibNFTOrder.OrderStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"}],\"name\":\"getHashNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"incrementHashNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc20TokenAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Fee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"}],\"internalType\":\"structLibNFTOrder.NFTSellOrder\",\"name\":\"sellOrder\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc20TokenAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Fee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractIPropertyValidator\",\"name\":\"propertyValidator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"propertyData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Property[]\",\"name\":\"nftProperties\",\"type\":\"tuple[]\"}],\"internalType\":\"structLibNFTOrder.NFTBuyOrder\",\"name\":\"buyOrder\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumLibSignature.SignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structLibSignature.Signature\",\"name\":\"sellOrderSignature\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumLibSignature.SignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structLibSignature.Signature\",\"name\":\"buyOrderSignature\",\"type\":\"tuple\"}],\"name\":\"matchERC721Orders\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"profit\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"success\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc20TokenAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Fee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractIPropertyValidator\",\"name\":\"propertyValidator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"propertyData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Property[]\",\"name\":\"nftProperties\",\"type\":\"tuple[]\"}],\"internalType\":\"structLibNFTOrder.NFTBuyOrder\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"preSignERC721BuyOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc20TokenAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Fee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"}],\"internalType\":\"structLibNFTOrder.NFTSellOrder\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"preSignERC721SellOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc20TokenAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Fee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractIPropertyValidator\",\"name\":\"propertyValidator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"propertyData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Property[]\",\"name\":\"nftProperties\",\"type\":\"tuple[]\"}],\"internalType\":\"structLibNFTOrder.NFTBuyOrder\",\"name\":\"buyOrder\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumLibSignature.SignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structLibSignature.Signature\",\"name\":\"signature\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"erc721TokenId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"unwrapNativeToken\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"takerData\",\"type\":\"bytes\"}],\"name\":\"sellERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc20TokenAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Fee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractIPropertyValidator\",\"name\":\"propertyValidator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"propertyData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Property[]\",\"name\":\"nftProperties\",\"type\":\"tuple[]\"}],\"internalType\":\"structLibNFTOrder.NFTBuyOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumLibSignature.SignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structLibSignature.Signature\",\"name\":\"signature\",\"type\":\"tuple\"}],\"name\":\"validateERC721BuyOrderSignature\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc20TokenAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Fee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"}],\"internalType\":\"structLibNFTOrder.NFTSellOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumLibSignature.SignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structLibSignature.Signature\",\"name\":\"signature\",\"type\":\"tuple\"}],\"name\":\"validateERC721SellOrderSignature\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

// GetERC721BuyOrderHash is a free data retrieval call binding the contract method 0x078e6b33.
//
// Solidity: function getERC721BuyOrderHash((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[]) order) view returns(bytes32)
func (_ERC721OrdersFeature *ERC721OrdersFeatureCaller) GetERC721BuyOrderHash(opts *bind.CallOpts, order LibNFTOrderNFTBuyOrder) ([32]byte, error) {
	var out []interface{}
	err := _ERC721OrdersFeature.contract.Call(opts, &out, "getERC721BuyOrderHash", order)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetERC721BuyOrderHash is a free data retrieval call binding the contract method 0x078e6b33.
//
// Solidity: function getERC721BuyOrderHash((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[]) order) view returns(bytes32)
func (_ERC721OrdersFeature *ERC721OrdersFeatureSession) GetERC721BuyOrderHash(order LibNFTOrderNFTBuyOrder) ([32]byte, error) {
	return _ERC721OrdersFeature.Contract.GetERC721BuyOrderHash(&_ERC721OrdersFeature.CallOpts, order)
}

// GetERC721BuyOrderHash is a free data retrieval call binding the contract method 0x078e6b33.
//
// Solidity: function getERC721BuyOrderHash((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[]) order) view returns(bytes32)
func (_ERC721OrdersFeature *ERC721OrdersFeatureCallerSession) GetERC721BuyOrderHash(order LibNFTOrderNFTBuyOrder) ([32]byte, error) {
	return _ERC721OrdersFeature.Contract.GetERC721BuyOrderHash(&_ERC721OrdersFeature.CallOpts, order)
}

// GetERC721BuyOrderStatus is a free data retrieval call binding the contract method 0xaf3de155.
//
// Solidity: function getERC721BuyOrderStatus((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[]) order) view returns(uint8)
func (_ERC721OrdersFeature *ERC721OrdersFeatureCaller) GetERC721BuyOrderStatus(opts *bind.CallOpts, order LibNFTOrderNFTBuyOrder) (uint8, error) {
	var out []interface{}
	err := _ERC721OrdersFeature.contract.Call(opts, &out, "getERC721BuyOrderStatus", order)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetERC721BuyOrderStatus is a free data retrieval call binding the contract method 0xaf3de155.
//
// Solidity: function getERC721BuyOrderStatus((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[]) order) view returns(uint8)
func (_ERC721OrdersFeature *ERC721OrdersFeatureSession) GetERC721BuyOrderStatus(order LibNFTOrderNFTBuyOrder) (uint8, error) {
	return _ERC721OrdersFeature.Contract.GetERC721BuyOrderStatus(&_ERC721OrdersFeature.CallOpts, order)
}

// GetERC721BuyOrderStatus is a free data retrieval call binding the contract method 0xaf3de155.
//
// Solidity: function getERC721BuyOrderStatus((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[]) order) view returns(uint8)
func (_ERC721OrdersFeature *ERC721OrdersFeatureCallerSession) GetERC721BuyOrderStatus(order LibNFTOrderNFTBuyOrder) (uint8, error) {
	return _ERC721OrdersFeature.Contract.GetERC721BuyOrderStatus(&_ERC721OrdersFeature.CallOpts, order)
}

// GetERC721OrderStatusBitVector is a free data retrieval call binding the contract method 0x030b2730.
//
// Solidity: function getERC721OrderStatusBitVector(address maker, uint248 nonceRange) view returns(uint256)
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
// Solidity: function getERC721OrderStatusBitVector(address maker, uint248 nonceRange) view returns(uint256)
func (_ERC721OrdersFeature *ERC721OrdersFeatureSession) GetERC721OrderStatusBitVector(maker common.Address, nonceRange *big.Int) (*big.Int, error) {
	return _ERC721OrdersFeature.Contract.GetERC721OrderStatusBitVector(&_ERC721OrdersFeature.CallOpts, maker, nonceRange)
}

// GetERC721OrderStatusBitVector is a free data retrieval call binding the contract method 0x030b2730.
//
// Solidity: function getERC721OrderStatusBitVector(address maker, uint248 nonceRange) view returns(uint256)
func (_ERC721OrdersFeature *ERC721OrdersFeatureCallerSession) GetERC721OrderStatusBitVector(maker common.Address, nonceRange *big.Int) (*big.Int, error) {
	return _ERC721OrdersFeature.Contract.GetERC721OrderStatusBitVector(&_ERC721OrdersFeature.CallOpts, maker, nonceRange)
}

// GetERC721SellOrderHash is a free data retrieval call binding the contract method 0x28784668.
//
// Solidity: function getERC721SellOrderHash((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256) order) view returns(bytes32)
func (_ERC721OrdersFeature *ERC721OrdersFeatureCaller) GetERC721SellOrderHash(opts *bind.CallOpts, order LibNFTOrderNFTSellOrder) ([32]byte, error) {
	var out []interface{}
	err := _ERC721OrdersFeature.contract.Call(opts, &out, "getERC721SellOrderHash", order)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetERC721SellOrderHash is a free data retrieval call binding the contract method 0x28784668.
//
// Solidity: function getERC721SellOrderHash((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256) order) view returns(bytes32)
func (_ERC721OrdersFeature *ERC721OrdersFeatureSession) GetERC721SellOrderHash(order LibNFTOrderNFTSellOrder) ([32]byte, error) {
	return _ERC721OrdersFeature.Contract.GetERC721SellOrderHash(&_ERC721OrdersFeature.CallOpts, order)
}

// GetERC721SellOrderHash is a free data retrieval call binding the contract method 0x28784668.
//
// Solidity: function getERC721SellOrderHash((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256) order) view returns(bytes32)
func (_ERC721OrdersFeature *ERC721OrdersFeatureCallerSession) GetERC721SellOrderHash(order LibNFTOrderNFTSellOrder) ([32]byte, error) {
	return _ERC721OrdersFeature.Contract.GetERC721SellOrderHash(&_ERC721OrdersFeature.CallOpts, order)
}

// GetERC721SellOrderStatus is a free data retrieval call binding the contract method 0x10a1ea2b.
//
// Solidity: function getERC721SellOrderStatus((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256) order) view returns(uint8)
func (_ERC721OrdersFeature *ERC721OrdersFeatureCaller) GetERC721SellOrderStatus(opts *bind.CallOpts, order LibNFTOrderNFTSellOrder) (uint8, error) {
	var out []interface{}
	err := _ERC721OrdersFeature.contract.Call(opts, &out, "getERC721SellOrderStatus", order)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetERC721SellOrderStatus is a free data retrieval call binding the contract method 0x10a1ea2b.
//
// Solidity: function getERC721SellOrderStatus((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256) order) view returns(uint8)
func (_ERC721OrdersFeature *ERC721OrdersFeatureSession) GetERC721SellOrderStatus(order LibNFTOrderNFTSellOrder) (uint8, error) {
	return _ERC721OrdersFeature.Contract.GetERC721SellOrderStatus(&_ERC721OrdersFeature.CallOpts, order)
}

// GetERC721SellOrderStatus is a free data retrieval call binding the contract method 0x10a1ea2b.
//
// Solidity: function getERC721SellOrderStatus((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256) order) view returns(uint8)
func (_ERC721OrdersFeature *ERC721OrdersFeatureCallerSession) GetERC721SellOrderStatus(order LibNFTOrderNFTSellOrder) (uint8, error) {
	return _ERC721OrdersFeature.Contract.GetERC721SellOrderStatus(&_ERC721OrdersFeature.CallOpts, order)
}

// GetHashNonce is a free data retrieval call binding the contract method 0x5e725186.
//
// Solidity: function getHashNonce(address maker) view returns(uint256)
func (_ERC721OrdersFeature *ERC721OrdersFeatureCaller) GetHashNonce(opts *bind.CallOpts, maker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC721OrdersFeature.contract.Call(opts, &out, "getHashNonce", maker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetHashNonce is a free data retrieval call binding the contract method 0x5e725186.
//
// Solidity: function getHashNonce(address maker) view returns(uint256)
func (_ERC721OrdersFeature *ERC721OrdersFeatureSession) GetHashNonce(maker common.Address) (*big.Int, error) {
	return _ERC721OrdersFeature.Contract.GetHashNonce(&_ERC721OrdersFeature.CallOpts, maker)
}

// GetHashNonce is a free data retrieval call binding the contract method 0x5e725186.
//
// Solidity: function getHashNonce(address maker) view returns(uint256)
func (_ERC721OrdersFeature *ERC721OrdersFeatureCallerSession) GetHashNonce(maker common.Address) (*big.Int, error) {
	return _ERC721OrdersFeature.Contract.GetHashNonce(&_ERC721OrdersFeature.CallOpts, maker)
}

// ValidateERC721BuyOrderSignature is a free data retrieval call binding the contract method 0xc67a8911.
//
// Solidity: function validateERC721BuyOrderSignature((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[]) order, (uint8,uint8,bytes32,bytes32) signature) view returns()
func (_ERC721OrdersFeature *ERC721OrdersFeatureCaller) ValidateERC721BuyOrderSignature(opts *bind.CallOpts, order LibNFTOrderNFTBuyOrder, signature LibSignatureSignature) error {
	var out []interface{}
	err := _ERC721OrdersFeature.contract.Call(opts, &out, "validateERC721BuyOrderSignature", order, signature)

	if err != nil {
		return err
	}

	return err

}

// ValidateERC721BuyOrderSignature is a free data retrieval call binding the contract method 0xc67a8911.
//
// Solidity: function validateERC721BuyOrderSignature((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[]) order, (uint8,uint8,bytes32,bytes32) signature) view returns()
func (_ERC721OrdersFeature *ERC721OrdersFeatureSession) ValidateERC721BuyOrderSignature(order LibNFTOrderNFTBuyOrder, signature LibSignatureSignature) error {
	return _ERC721OrdersFeature.Contract.ValidateERC721BuyOrderSignature(&_ERC721OrdersFeature.CallOpts, order, signature)
}

// ValidateERC721BuyOrderSignature is a free data retrieval call binding the contract method 0xc67a8911.
//
// Solidity: function validateERC721BuyOrderSignature((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[]) order, (uint8,uint8,bytes32,bytes32) signature) view returns()
func (_ERC721OrdersFeature *ERC721OrdersFeatureCallerSession) ValidateERC721BuyOrderSignature(order LibNFTOrderNFTBuyOrder, signature LibSignatureSignature) error {
	return _ERC721OrdersFeature.Contract.ValidateERC721BuyOrderSignature(&_ERC721OrdersFeature.CallOpts, order, signature)
}

// ValidateERC721SellOrderSignature is a free data retrieval call binding the contract method 0x6e74f68a.
//
// Solidity: function validateERC721SellOrderSignature((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256) order, (uint8,uint8,bytes32,bytes32) signature) view returns()
func (_ERC721OrdersFeature *ERC721OrdersFeatureCaller) ValidateERC721SellOrderSignature(opts *bind.CallOpts, order LibNFTOrderNFTSellOrder, signature LibSignatureSignature) error {
	var out []interface{}
	err := _ERC721OrdersFeature.contract.Call(opts, &out, "validateERC721SellOrderSignature", order, signature)

	if err != nil {
		return err
	}

	return err

}

// ValidateERC721SellOrderSignature is a free data retrieval call binding the contract method 0x6e74f68a.
//
// Solidity: function validateERC721SellOrderSignature((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256) order, (uint8,uint8,bytes32,bytes32) signature) view returns()
func (_ERC721OrdersFeature *ERC721OrdersFeatureSession) ValidateERC721SellOrderSignature(order LibNFTOrderNFTSellOrder, signature LibSignatureSignature) error {
	return _ERC721OrdersFeature.Contract.ValidateERC721SellOrderSignature(&_ERC721OrdersFeature.CallOpts, order, signature)
}

// ValidateERC721SellOrderSignature is a free data retrieval call binding the contract method 0x6e74f68a.
//
// Solidity: function validateERC721SellOrderSignature((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256) order, (uint8,uint8,bytes32,bytes32) signature) view returns()
func (_ERC721OrdersFeature *ERC721OrdersFeatureCallerSession) ValidateERC721SellOrderSignature(order LibNFTOrderNFTSellOrder, signature LibSignatureSignature) error {
	return _ERC721OrdersFeature.Contract.ValidateERC721SellOrderSignature(&_ERC721OrdersFeature.CallOpts, order, signature)
}

// BatchBuyERC721s is a paid mutator transaction binding the contract method 0xd8d2ae77.
//
// Solidity: function batchBuyERC721s((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256)[] sellOrders, (uint8,uint8,bytes32,bytes32)[] signatures, bool revertIfIncomplete) payable returns(bool[] successes)
func (_ERC721OrdersFeature *ERC721OrdersFeatureTransactor) BatchBuyERC721s(opts *bind.TransactOpts, sellOrders []LibNFTOrderNFTSellOrder, signatures []LibSignatureSignature, revertIfIncomplete bool) (*types.Transaction, error) {
	return _ERC721OrdersFeature.contract.Transact(opts, "batchBuyERC721s", sellOrders, signatures, revertIfIncomplete)
}

// BatchBuyERC721s is a paid mutator transaction binding the contract method 0xd8d2ae77.
//
// Solidity: function batchBuyERC721s((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256)[] sellOrders, (uint8,uint8,bytes32,bytes32)[] signatures, bool revertIfIncomplete) payable returns(bool[] successes)
func (_ERC721OrdersFeature *ERC721OrdersFeatureSession) BatchBuyERC721s(sellOrders []LibNFTOrderNFTSellOrder, signatures []LibSignatureSignature, revertIfIncomplete bool) (*types.Transaction, error) {
	return _ERC721OrdersFeature.Contract.BatchBuyERC721s(&_ERC721OrdersFeature.TransactOpts, sellOrders, signatures, revertIfIncomplete)
}

// BatchBuyERC721s is a paid mutator transaction binding the contract method 0xd8d2ae77.
//
// Solidity: function batchBuyERC721s((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256)[] sellOrders, (uint8,uint8,bytes32,bytes32)[] signatures, bool revertIfIncomplete) payable returns(bool[] successes)
func (_ERC721OrdersFeature *ERC721OrdersFeatureTransactorSession) BatchBuyERC721s(sellOrders []LibNFTOrderNFTSellOrder, signatures []LibSignatureSignature, revertIfIncomplete bool) (*types.Transaction, error) {
	return _ERC721OrdersFeature.Contract.BatchBuyERC721s(&_ERC721OrdersFeature.TransactOpts, sellOrders, signatures, revertIfIncomplete)
}

// BatchBuyERC721sEx is a paid mutator transaction binding the contract method 0x053c23f1.
//
// Solidity: function batchBuyERC721sEx((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256)[] sellOrders, (uint8,uint8,bytes32,bytes32)[] signatures, address[] takers, bytes[] takerDatas, bool revertIfIncomplete) payable returns(bool[] successes)
func (_ERC721OrdersFeature *ERC721OrdersFeatureTransactor) BatchBuyERC721sEx(opts *bind.TransactOpts, sellOrders []LibNFTOrderNFTSellOrder, signatures []LibSignatureSignature, takers []common.Address, takerDatas [][]byte, revertIfIncomplete bool) (*types.Transaction, error) {
	return _ERC721OrdersFeature.contract.Transact(opts, "batchBuyERC721sEx", sellOrders, signatures, takers, takerDatas, revertIfIncomplete)
}

// BatchBuyERC721sEx is a paid mutator transaction binding the contract method 0x053c23f1.
//
// Solidity: function batchBuyERC721sEx((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256)[] sellOrders, (uint8,uint8,bytes32,bytes32)[] signatures, address[] takers, bytes[] takerDatas, bool revertIfIncomplete) payable returns(bool[] successes)
func (_ERC721OrdersFeature *ERC721OrdersFeatureSession) BatchBuyERC721sEx(sellOrders []LibNFTOrderNFTSellOrder, signatures []LibSignatureSignature, takers []common.Address, takerDatas [][]byte, revertIfIncomplete bool) (*types.Transaction, error) {
	return _ERC721OrdersFeature.Contract.BatchBuyERC721sEx(&_ERC721OrdersFeature.TransactOpts, sellOrders, signatures, takers, takerDatas, revertIfIncomplete)
}

// BatchBuyERC721sEx is a paid mutator transaction binding the contract method 0x053c23f1.
//
// Solidity: function batchBuyERC721sEx((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256)[] sellOrders, (uint8,uint8,bytes32,bytes32)[] signatures, address[] takers, bytes[] takerDatas, bool revertIfIncomplete) payable returns(bool[] successes)
func (_ERC721OrdersFeature *ERC721OrdersFeatureTransactorSession) BatchBuyERC721sEx(sellOrders []LibNFTOrderNFTSellOrder, signatures []LibSignatureSignature, takers []common.Address, takerDatas [][]byte, revertIfIncomplete bool) (*types.Transaction, error) {
	return _ERC721OrdersFeature.Contract.BatchBuyERC721sEx(&_ERC721OrdersFeature.TransactOpts, sellOrders, signatures, takers, takerDatas, revertIfIncomplete)
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

// BatchMatchERC721Orders is a paid mutator transaction binding the contract method 0x8e3efd59.
//
// Solidity: function batchMatchERC721Orders((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256)[] sellOrders, (address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[])[] buyOrders, (uint8,uint8,bytes32,bytes32)[] sellOrderSignatures, (uint8,uint8,bytes32,bytes32)[] buyOrderSignatures) returns(uint256[] profits, bool[] successes)
func (_ERC721OrdersFeature *ERC721OrdersFeatureTransactor) BatchMatchERC721Orders(opts *bind.TransactOpts, sellOrders []LibNFTOrderNFTSellOrder, buyOrders []LibNFTOrderNFTBuyOrder, sellOrderSignatures []LibSignatureSignature, buyOrderSignatures []LibSignatureSignature) (*types.Transaction, error) {
	return _ERC721OrdersFeature.contract.Transact(opts, "batchMatchERC721Orders", sellOrders, buyOrders, sellOrderSignatures, buyOrderSignatures)
}

// BatchMatchERC721Orders is a paid mutator transaction binding the contract method 0x8e3efd59.
//
// Solidity: function batchMatchERC721Orders((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256)[] sellOrders, (address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[])[] buyOrders, (uint8,uint8,bytes32,bytes32)[] sellOrderSignatures, (uint8,uint8,bytes32,bytes32)[] buyOrderSignatures) returns(uint256[] profits, bool[] successes)
func (_ERC721OrdersFeature *ERC721OrdersFeatureSession) BatchMatchERC721Orders(sellOrders []LibNFTOrderNFTSellOrder, buyOrders []LibNFTOrderNFTBuyOrder, sellOrderSignatures []LibSignatureSignature, buyOrderSignatures []LibSignatureSignature) (*types.Transaction, error) {
	return _ERC721OrdersFeature.Contract.BatchMatchERC721Orders(&_ERC721OrdersFeature.TransactOpts, sellOrders, buyOrders, sellOrderSignatures, buyOrderSignatures)
}

// BatchMatchERC721Orders is a paid mutator transaction binding the contract method 0x8e3efd59.
//
// Solidity: function batchMatchERC721Orders((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256)[] sellOrders, (address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[])[] buyOrders, (uint8,uint8,bytes32,bytes32)[] sellOrderSignatures, (uint8,uint8,bytes32,bytes32)[] buyOrderSignatures) returns(uint256[] profits, bool[] successes)
func (_ERC721OrdersFeature *ERC721OrdersFeatureTransactorSession) BatchMatchERC721Orders(sellOrders []LibNFTOrderNFTSellOrder, buyOrders []LibNFTOrderNFTBuyOrder, sellOrderSignatures []LibSignatureSignature, buyOrderSignatures []LibSignatureSignature) (*types.Transaction, error) {
	return _ERC721OrdersFeature.Contract.BatchMatchERC721Orders(&_ERC721OrdersFeature.TransactOpts, sellOrders, buyOrders, sellOrderSignatures, buyOrderSignatures)
}

// BuyERC721 is a paid mutator transaction binding the contract method 0x5f57685e.
//
// Solidity: function buyERC721((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256) sellOrder, (uint8,uint8,bytes32,bytes32) signature) payable returns()
func (_ERC721OrdersFeature *ERC721OrdersFeatureTransactor) BuyERC721(opts *bind.TransactOpts, sellOrder LibNFTOrderNFTSellOrder, signature LibSignatureSignature) (*types.Transaction, error) {
	return _ERC721OrdersFeature.contract.Transact(opts, "buyERC721", sellOrder, signature)
}

// BuyERC721 is a paid mutator transaction binding the contract method 0x5f57685e.
//
// Solidity: function buyERC721((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256) sellOrder, (uint8,uint8,bytes32,bytes32) signature) payable returns()
func (_ERC721OrdersFeature *ERC721OrdersFeatureSession) BuyERC721(sellOrder LibNFTOrderNFTSellOrder, signature LibSignatureSignature) (*types.Transaction, error) {
	return _ERC721OrdersFeature.Contract.BuyERC721(&_ERC721OrdersFeature.TransactOpts, sellOrder, signature)
}

// BuyERC721 is a paid mutator transaction binding the contract method 0x5f57685e.
//
// Solidity: function buyERC721((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256) sellOrder, (uint8,uint8,bytes32,bytes32) signature) payable returns()
func (_ERC721OrdersFeature *ERC721OrdersFeatureTransactorSession) BuyERC721(sellOrder LibNFTOrderNFTSellOrder, signature LibSignatureSignature) (*types.Transaction, error) {
	return _ERC721OrdersFeature.Contract.BuyERC721(&_ERC721OrdersFeature.TransactOpts, sellOrder, signature)
}

// BuyERC721Ex is a paid mutator transaction binding the contract method 0xb18d619f.
//
// Solidity: function buyERC721Ex((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256) sellOrder, (uint8,uint8,bytes32,bytes32) signature, address taker, bytes takerData) payable returns()
func (_ERC721OrdersFeature *ERC721OrdersFeatureTransactor) BuyERC721Ex(opts *bind.TransactOpts, sellOrder LibNFTOrderNFTSellOrder, signature LibSignatureSignature, taker common.Address, takerData []byte) (*types.Transaction, error) {
	return _ERC721OrdersFeature.contract.Transact(opts, "buyERC721Ex", sellOrder, signature, taker, takerData)
}

// BuyERC721Ex is a paid mutator transaction binding the contract method 0xb18d619f.
//
// Solidity: function buyERC721Ex((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256) sellOrder, (uint8,uint8,bytes32,bytes32) signature, address taker, bytes takerData) payable returns()
func (_ERC721OrdersFeature *ERC721OrdersFeatureSession) BuyERC721Ex(sellOrder LibNFTOrderNFTSellOrder, signature LibSignatureSignature, taker common.Address, takerData []byte) (*types.Transaction, error) {
	return _ERC721OrdersFeature.Contract.BuyERC721Ex(&_ERC721OrdersFeature.TransactOpts, sellOrder, signature, taker, takerData)
}

// BuyERC721Ex is a paid mutator transaction binding the contract method 0xb18d619f.
//
// Solidity: function buyERC721Ex((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256) sellOrder, (uint8,uint8,bytes32,bytes32) signature, address taker, bytes takerData) payable returns()
func (_ERC721OrdersFeature *ERC721OrdersFeatureTransactorSession) BuyERC721Ex(sellOrder LibNFTOrderNFTSellOrder, signature LibSignatureSignature, taker common.Address, takerData []byte) (*types.Transaction, error) {
	return _ERC721OrdersFeature.Contract.BuyERC721Ex(&_ERC721OrdersFeature.TransactOpts, sellOrder, signature, taker, takerData)
}

// BuyERC721ExFromProxy is a paid mutator transaction binding the contract method 0x32766750.
//
// Solidity: function buyERC721ExFromProxy((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256) sellOrder, (uint8,uint8,bytes32,bytes32) signature, address taker, uint256 ethAvailable, bytes takerData) payable returns()
func (_ERC721OrdersFeature *ERC721OrdersFeatureTransactor) BuyERC721ExFromProxy(opts *bind.TransactOpts, sellOrder LibNFTOrderNFTSellOrder, signature LibSignatureSignature, taker common.Address, ethAvailable *big.Int, takerData []byte) (*types.Transaction, error) {
	return _ERC721OrdersFeature.contract.Transact(opts, "buyERC721ExFromProxy", sellOrder, signature, taker, ethAvailable, takerData)
}

// BuyERC721ExFromProxy is a paid mutator transaction binding the contract method 0x32766750.
//
// Solidity: function buyERC721ExFromProxy((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256) sellOrder, (uint8,uint8,bytes32,bytes32) signature, address taker, uint256 ethAvailable, bytes takerData) payable returns()
func (_ERC721OrdersFeature *ERC721OrdersFeatureSession) BuyERC721ExFromProxy(sellOrder LibNFTOrderNFTSellOrder, signature LibSignatureSignature, taker common.Address, ethAvailable *big.Int, takerData []byte) (*types.Transaction, error) {
	return _ERC721OrdersFeature.Contract.BuyERC721ExFromProxy(&_ERC721OrdersFeature.TransactOpts, sellOrder, signature, taker, ethAvailable, takerData)
}

// BuyERC721ExFromProxy is a paid mutator transaction binding the contract method 0x32766750.
//
// Solidity: function buyERC721ExFromProxy((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256) sellOrder, (uint8,uint8,bytes32,bytes32) signature, address taker, uint256 ethAvailable, bytes takerData) payable returns()
func (_ERC721OrdersFeature *ERC721OrdersFeatureTransactorSession) BuyERC721ExFromProxy(sellOrder LibNFTOrderNFTSellOrder, signature LibSignatureSignature, taker common.Address, ethAvailable *big.Int, takerData []byte) (*types.Transaction, error) {
	return _ERC721OrdersFeature.Contract.BuyERC721ExFromProxy(&_ERC721OrdersFeature.TransactOpts, sellOrder, signature, taker, ethAvailable, takerData)
}

// BuyERC721FromProxy is a paid mutator transaction binding the contract method 0x33ba623b.
//
// Solidity: function buyERC721FromProxy((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256) sellOrder, (uint8,uint8,bytes32,bytes32) signature) payable returns()
func (_ERC721OrdersFeature *ERC721OrdersFeatureTransactor) BuyERC721FromProxy(opts *bind.TransactOpts, sellOrder LibNFTOrderNFTSellOrder, signature LibSignatureSignature) (*types.Transaction, error) {
	return _ERC721OrdersFeature.contract.Transact(opts, "buyERC721FromProxy", sellOrder, signature)
}

// BuyERC721FromProxy is a paid mutator transaction binding the contract method 0x33ba623b.
//
// Solidity: function buyERC721FromProxy((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256) sellOrder, (uint8,uint8,bytes32,bytes32) signature) payable returns()
func (_ERC721OrdersFeature *ERC721OrdersFeatureSession) BuyERC721FromProxy(sellOrder LibNFTOrderNFTSellOrder, signature LibSignatureSignature) (*types.Transaction, error) {
	return _ERC721OrdersFeature.Contract.BuyERC721FromProxy(&_ERC721OrdersFeature.TransactOpts, sellOrder, signature)
}

// BuyERC721FromProxy is a paid mutator transaction binding the contract method 0x33ba623b.
//
// Solidity: function buyERC721FromProxy((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256) sellOrder, (uint8,uint8,bytes32,bytes32) signature) payable returns()
func (_ERC721OrdersFeature *ERC721OrdersFeatureTransactorSession) BuyERC721FromProxy(sellOrder LibNFTOrderNFTSellOrder, signature LibSignatureSignature) (*types.Transaction, error) {
	return _ERC721OrdersFeature.Contract.BuyERC721FromProxy(&_ERC721OrdersFeature.TransactOpts, sellOrder, signature)
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

// IncrementHashNonce is a paid mutator transaction binding the contract method 0x050505d6.
//
// Solidity: function incrementHashNonce() returns()
func (_ERC721OrdersFeature *ERC721OrdersFeatureTransactor) IncrementHashNonce(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721OrdersFeature.contract.Transact(opts, "incrementHashNonce")
}

// IncrementHashNonce is a paid mutator transaction binding the contract method 0x050505d6.
//
// Solidity: function incrementHashNonce() returns()
func (_ERC721OrdersFeature *ERC721OrdersFeatureSession) IncrementHashNonce() (*types.Transaction, error) {
	return _ERC721OrdersFeature.Contract.IncrementHashNonce(&_ERC721OrdersFeature.TransactOpts)
}

// IncrementHashNonce is a paid mutator transaction binding the contract method 0x050505d6.
//
// Solidity: function incrementHashNonce() returns()
func (_ERC721OrdersFeature *ERC721OrdersFeatureTransactorSession) IncrementHashNonce() (*types.Transaction, error) {
	return _ERC721OrdersFeature.Contract.IncrementHashNonce(&_ERC721OrdersFeature.TransactOpts)
}

// MatchERC721Orders is a paid mutator transaction binding the contract method 0x1fc34ffe.
//
// Solidity: function matchERC721Orders((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256) sellOrder, (address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[]) buyOrder, (uint8,uint8,bytes32,bytes32) sellOrderSignature, (uint8,uint8,bytes32,bytes32) buyOrderSignature) returns(uint256 profit)
func (_ERC721OrdersFeature *ERC721OrdersFeatureTransactor) MatchERC721Orders(opts *bind.TransactOpts, sellOrder LibNFTOrderNFTSellOrder, buyOrder LibNFTOrderNFTBuyOrder, sellOrderSignature LibSignatureSignature, buyOrderSignature LibSignatureSignature) (*types.Transaction, error) {
	return _ERC721OrdersFeature.contract.Transact(opts, "matchERC721Orders", sellOrder, buyOrder, sellOrderSignature, buyOrderSignature)
}

// MatchERC721Orders is a paid mutator transaction binding the contract method 0x1fc34ffe.
//
// Solidity: function matchERC721Orders((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256) sellOrder, (address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[]) buyOrder, (uint8,uint8,bytes32,bytes32) sellOrderSignature, (uint8,uint8,bytes32,bytes32) buyOrderSignature) returns(uint256 profit)
func (_ERC721OrdersFeature *ERC721OrdersFeatureSession) MatchERC721Orders(sellOrder LibNFTOrderNFTSellOrder, buyOrder LibNFTOrderNFTBuyOrder, sellOrderSignature LibSignatureSignature, buyOrderSignature LibSignatureSignature) (*types.Transaction, error) {
	return _ERC721OrdersFeature.Contract.MatchERC721Orders(&_ERC721OrdersFeature.TransactOpts, sellOrder, buyOrder, sellOrderSignature, buyOrderSignature)
}

// MatchERC721Orders is a paid mutator transaction binding the contract method 0x1fc34ffe.
//
// Solidity: function matchERC721Orders((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256) sellOrder, (address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[]) buyOrder, (uint8,uint8,bytes32,bytes32) sellOrderSignature, (uint8,uint8,bytes32,bytes32) buyOrderSignature) returns(uint256 profit)
func (_ERC721OrdersFeature *ERC721OrdersFeatureTransactorSession) MatchERC721Orders(sellOrder LibNFTOrderNFTSellOrder, buyOrder LibNFTOrderNFTBuyOrder, sellOrderSignature LibSignatureSignature, buyOrderSignature LibSignatureSignature) (*types.Transaction, error) {
	return _ERC721OrdersFeature.Contract.MatchERC721Orders(&_ERC721OrdersFeature.TransactOpts, sellOrder, buyOrder, sellOrderSignature, buyOrderSignature)
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

// PreSignERC721BuyOrder is a paid mutator transaction binding the contract method 0x332b024f.
//
// Solidity: function preSignERC721BuyOrder((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[]) order) returns()
func (_ERC721OrdersFeature *ERC721OrdersFeatureTransactor) PreSignERC721BuyOrder(opts *bind.TransactOpts, order LibNFTOrderNFTBuyOrder) (*types.Transaction, error) {
	return _ERC721OrdersFeature.contract.Transact(opts, "preSignERC721BuyOrder", order)
}

// PreSignERC721BuyOrder is a paid mutator transaction binding the contract method 0x332b024f.
//
// Solidity: function preSignERC721BuyOrder((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[]) order) returns()
func (_ERC721OrdersFeature *ERC721OrdersFeatureSession) PreSignERC721BuyOrder(order LibNFTOrderNFTBuyOrder) (*types.Transaction, error) {
	return _ERC721OrdersFeature.Contract.PreSignERC721BuyOrder(&_ERC721OrdersFeature.TransactOpts, order)
}

// PreSignERC721BuyOrder is a paid mutator transaction binding the contract method 0x332b024f.
//
// Solidity: function preSignERC721BuyOrder((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[]) order) returns()
func (_ERC721OrdersFeature *ERC721OrdersFeatureTransactorSession) PreSignERC721BuyOrder(order LibNFTOrderNFTBuyOrder) (*types.Transaction, error) {
	return _ERC721OrdersFeature.Contract.PreSignERC721BuyOrder(&_ERC721OrdersFeature.TransactOpts, order)
}

// PreSignERC721SellOrder is a paid mutator transaction binding the contract method 0x28a96208.
//
// Solidity: function preSignERC721SellOrder((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256) order) returns()
func (_ERC721OrdersFeature *ERC721OrdersFeatureTransactor) PreSignERC721SellOrder(opts *bind.TransactOpts, order LibNFTOrderNFTSellOrder) (*types.Transaction, error) {
	return _ERC721OrdersFeature.contract.Transact(opts, "preSignERC721SellOrder", order)
}

// PreSignERC721SellOrder is a paid mutator transaction binding the contract method 0x28a96208.
//
// Solidity: function preSignERC721SellOrder((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256) order) returns()
func (_ERC721OrdersFeature *ERC721OrdersFeatureSession) PreSignERC721SellOrder(order LibNFTOrderNFTSellOrder) (*types.Transaction, error) {
	return _ERC721OrdersFeature.Contract.PreSignERC721SellOrder(&_ERC721OrdersFeature.TransactOpts, order)
}

// PreSignERC721SellOrder is a paid mutator transaction binding the contract method 0x28a96208.
//
// Solidity: function preSignERC721SellOrder((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256) order) returns()
func (_ERC721OrdersFeature *ERC721OrdersFeatureTransactorSession) PreSignERC721SellOrder(order LibNFTOrderNFTSellOrder) (*types.Transaction, error) {
	return _ERC721OrdersFeature.Contract.PreSignERC721SellOrder(&_ERC721OrdersFeature.TransactOpts, order)
}

// SellERC721 is a paid mutator transaction binding the contract method 0xa8809485.
//
// Solidity: function sellERC721((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[]) buyOrder, (uint8,uint8,bytes32,bytes32) signature, uint256 erc721TokenId, bool unwrapNativeToken, bytes takerData) returns()
func (_ERC721OrdersFeature *ERC721OrdersFeatureTransactor) SellERC721(opts *bind.TransactOpts, buyOrder LibNFTOrderNFTBuyOrder, signature LibSignatureSignature, erc721TokenId *big.Int, unwrapNativeToken bool, takerData []byte) (*types.Transaction, error) {
	return _ERC721OrdersFeature.contract.Transact(opts, "sellERC721", buyOrder, signature, erc721TokenId, unwrapNativeToken, takerData)
}

// SellERC721 is a paid mutator transaction binding the contract method 0xa8809485.
//
// Solidity: function sellERC721((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[]) buyOrder, (uint8,uint8,bytes32,bytes32) signature, uint256 erc721TokenId, bool unwrapNativeToken, bytes takerData) returns()
func (_ERC721OrdersFeature *ERC721OrdersFeatureSession) SellERC721(buyOrder LibNFTOrderNFTBuyOrder, signature LibSignatureSignature, erc721TokenId *big.Int, unwrapNativeToken bool, takerData []byte) (*types.Transaction, error) {
	return _ERC721OrdersFeature.Contract.SellERC721(&_ERC721OrdersFeature.TransactOpts, buyOrder, signature, erc721TokenId, unwrapNativeToken, takerData)
}

// SellERC721 is a paid mutator transaction binding the contract method 0xa8809485.
//
// Solidity: function sellERC721((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[]) buyOrder, (uint8,uint8,bytes32,bytes32) signature, uint256 erc721TokenId, bool unwrapNativeToken, bytes takerData) returns()
func (_ERC721OrdersFeature *ERC721OrdersFeatureTransactorSession) SellERC721(buyOrder LibNFTOrderNFTBuyOrder, signature LibSignatureSignature, erc721TokenId *big.Int, unwrapNativeToken bool, takerData []byte) (*types.Transaction, error) {
	return _ERC721OrdersFeature.Contract.SellERC721(&_ERC721OrdersFeature.TransactOpts, buyOrder, signature, erc721TokenId, unwrapNativeToken, takerData)
}

// ERC721OrdersFeatureERC721BuyOrderFilledIterator is returned from FilterERC721BuyOrderFilled and is used to iterate over the raw logs and unpacked data for ERC721BuyOrderFilled events raised by the ERC721OrdersFeature contract.
type ERC721OrdersFeatureERC721BuyOrderFilledIterator struct {
	Event *ERC721OrdersFeatureERC721BuyOrderFilled // Event containing the contract specifics and raw log

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
func (it *ERC721OrdersFeatureERC721BuyOrderFilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721OrdersFeatureERC721BuyOrderFilled)
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
		it.Event = new(ERC721OrdersFeatureERC721BuyOrderFilled)
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
func (it *ERC721OrdersFeatureERC721BuyOrderFilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721OrdersFeatureERC721BuyOrderFilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721OrdersFeatureERC721BuyOrderFilled represents a ERC721BuyOrderFilled event raised by the ERC721OrdersFeature contract.
type ERC721OrdersFeatureERC721BuyOrderFilled struct {
	OrderHash        [32]byte
	Maker            common.Address
	Taker            common.Address
	Nonce            *big.Int
	Erc20Token       common.Address
	Erc20TokenAmount *big.Int
	Fees             []INFTOrdersFeatureFee
	Erc721Token      common.Address
	Erc721TokenId    *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterERC721BuyOrderFilled is a free log retrieval operation binding the contract event 0xd90a5c60975c6ff8eafcf02088e7b50ae5d9e156a79206ba553df1c4fb4594c2.
//
// Solidity: event ERC721BuyOrderFilled(bytes32 orderHash, address maker, address taker, uint256 nonce, address erc20Token, uint256 erc20TokenAmount, (address,uint256)[] fees, address erc721Token, uint256 erc721TokenId)
func (_ERC721OrdersFeature *ERC721OrdersFeatureFilterer) FilterERC721BuyOrderFilled(opts *bind.FilterOpts) (*ERC721OrdersFeatureERC721BuyOrderFilledIterator, error) {

	logs, sub, err := _ERC721OrdersFeature.contract.FilterLogs(opts, "ERC721BuyOrderFilled")
	if err != nil {
		return nil, err
	}
	return &ERC721OrdersFeatureERC721BuyOrderFilledIterator{contract: _ERC721OrdersFeature.contract, event: "ERC721BuyOrderFilled", logs: logs, sub: sub}, nil
}

// WatchERC721BuyOrderFilled is a free log subscription operation binding the contract event 0xd90a5c60975c6ff8eafcf02088e7b50ae5d9e156a79206ba553df1c4fb4594c2.
//
// Solidity: event ERC721BuyOrderFilled(bytes32 orderHash, address maker, address taker, uint256 nonce, address erc20Token, uint256 erc20TokenAmount, (address,uint256)[] fees, address erc721Token, uint256 erc721TokenId)
func (_ERC721OrdersFeature *ERC721OrdersFeatureFilterer) WatchERC721BuyOrderFilled(opts *bind.WatchOpts, sink chan<- *ERC721OrdersFeatureERC721BuyOrderFilled) (event.Subscription, error) {

	logs, sub, err := _ERC721OrdersFeature.contract.WatchLogs(opts, "ERC721BuyOrderFilled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721OrdersFeatureERC721BuyOrderFilled)
				if err := _ERC721OrdersFeature.contract.UnpackLog(event, "ERC721BuyOrderFilled", log); err != nil {
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

// ParseERC721BuyOrderFilled is a log parse operation binding the contract event 0xd90a5c60975c6ff8eafcf02088e7b50ae5d9e156a79206ba553df1c4fb4594c2.
//
// Solidity: event ERC721BuyOrderFilled(bytes32 orderHash, address maker, address taker, uint256 nonce, address erc20Token, uint256 erc20TokenAmount, (address,uint256)[] fees, address erc721Token, uint256 erc721TokenId)
func (_ERC721OrdersFeature *ERC721OrdersFeatureFilterer) ParseERC721BuyOrderFilled(log types.Log) (*ERC721OrdersFeatureERC721BuyOrderFilled, error) {
	event := new(ERC721OrdersFeatureERC721BuyOrderFilled)
	if err := _ERC721OrdersFeature.contract.UnpackLog(event, "ERC721BuyOrderFilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC721OrdersFeatureERC721BuyOrderPreSignedIterator is returned from FilterERC721BuyOrderPreSigned and is used to iterate over the raw logs and unpacked data for ERC721BuyOrderPreSigned events raised by the ERC721OrdersFeature contract.
type ERC721OrdersFeatureERC721BuyOrderPreSignedIterator struct {
	Event *ERC721OrdersFeatureERC721BuyOrderPreSigned // Event containing the contract specifics and raw log

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
func (it *ERC721OrdersFeatureERC721BuyOrderPreSignedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721OrdersFeatureERC721BuyOrderPreSigned)
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
		it.Event = new(ERC721OrdersFeatureERC721BuyOrderPreSigned)
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
func (it *ERC721OrdersFeatureERC721BuyOrderPreSignedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721OrdersFeatureERC721BuyOrderPreSignedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721OrdersFeatureERC721BuyOrderPreSigned represents a ERC721BuyOrderPreSigned event raised by the ERC721OrdersFeature contract.
type ERC721OrdersFeatureERC721BuyOrderPreSigned struct {
	Maker            common.Address
	Taker            common.Address
	Expiry           *big.Int
	Nonce            *big.Int
	Erc20Token       common.Address
	Erc20TokenAmount *big.Int
	Fees             []LibNFTOrderFee
	Erc721Token      common.Address
	Erc721TokenId    *big.Int
	NftProperties    []LibNFTOrderProperty
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterERC721BuyOrderPreSigned is a free log retrieval operation binding the contract event 0x4c2669b38ff3018c301fbc65423ac87447906bcf66b95a5fe0d3c5bbd6bcb297.
//
// Solidity: event ERC721BuyOrderPreSigned(address maker, address taker, uint256 expiry, uint256 nonce, address erc20Token, uint256 erc20TokenAmount, (address,uint256,bytes)[] fees, address erc721Token, uint256 erc721TokenId, (address,bytes)[] nftProperties)
func (_ERC721OrdersFeature *ERC721OrdersFeatureFilterer) FilterERC721BuyOrderPreSigned(opts *bind.FilterOpts) (*ERC721OrdersFeatureERC721BuyOrderPreSignedIterator, error) {

	logs, sub, err := _ERC721OrdersFeature.contract.FilterLogs(opts, "ERC721BuyOrderPreSigned")
	if err != nil {
		return nil, err
	}
	return &ERC721OrdersFeatureERC721BuyOrderPreSignedIterator{contract: _ERC721OrdersFeature.contract, event: "ERC721BuyOrderPreSigned", logs: logs, sub: sub}, nil
}

// WatchERC721BuyOrderPreSigned is a free log subscription operation binding the contract event 0x4c2669b38ff3018c301fbc65423ac87447906bcf66b95a5fe0d3c5bbd6bcb297.
//
// Solidity: event ERC721BuyOrderPreSigned(address maker, address taker, uint256 expiry, uint256 nonce, address erc20Token, uint256 erc20TokenAmount, (address,uint256,bytes)[] fees, address erc721Token, uint256 erc721TokenId, (address,bytes)[] nftProperties)
func (_ERC721OrdersFeature *ERC721OrdersFeatureFilterer) WatchERC721BuyOrderPreSigned(opts *bind.WatchOpts, sink chan<- *ERC721OrdersFeatureERC721BuyOrderPreSigned) (event.Subscription, error) {

	logs, sub, err := _ERC721OrdersFeature.contract.WatchLogs(opts, "ERC721BuyOrderPreSigned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721OrdersFeatureERC721BuyOrderPreSigned)
				if err := _ERC721OrdersFeature.contract.UnpackLog(event, "ERC721BuyOrderPreSigned", log); err != nil {
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

// ParseERC721BuyOrderPreSigned is a log parse operation binding the contract event 0x4c2669b38ff3018c301fbc65423ac87447906bcf66b95a5fe0d3c5bbd6bcb297.
//
// Solidity: event ERC721BuyOrderPreSigned(address maker, address taker, uint256 expiry, uint256 nonce, address erc20Token, uint256 erc20TokenAmount, (address,uint256,bytes)[] fees, address erc721Token, uint256 erc721TokenId, (address,bytes)[] nftProperties)
func (_ERC721OrdersFeature *ERC721OrdersFeatureFilterer) ParseERC721BuyOrderPreSigned(log types.Log) (*ERC721OrdersFeatureERC721BuyOrderPreSigned, error) {
	event := new(ERC721OrdersFeatureERC721BuyOrderPreSigned)
	if err := _ERC721OrdersFeature.contract.UnpackLog(event, "ERC721BuyOrderPreSigned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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

// ERC721OrdersFeatureERC721SellOrderFilledIterator is returned from FilterERC721SellOrderFilled and is used to iterate over the raw logs and unpacked data for ERC721SellOrderFilled events raised by the ERC721OrdersFeature contract.
type ERC721OrdersFeatureERC721SellOrderFilledIterator struct {
	Event *ERC721OrdersFeatureERC721SellOrderFilled // Event containing the contract specifics and raw log

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
func (it *ERC721OrdersFeatureERC721SellOrderFilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721OrdersFeatureERC721SellOrderFilled)
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
		it.Event = new(ERC721OrdersFeatureERC721SellOrderFilled)
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
func (it *ERC721OrdersFeatureERC721SellOrderFilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721OrdersFeatureERC721SellOrderFilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721OrdersFeatureERC721SellOrderFilled represents a ERC721SellOrderFilled event raised by the ERC721OrdersFeature contract.
type ERC721OrdersFeatureERC721SellOrderFilled struct {
	OrderHash        [32]byte
	Maker            common.Address
	Taker            common.Address
	Nonce            *big.Int
	Erc20Token       common.Address
	Erc20TokenAmount *big.Int
	Fees             []INFTOrdersFeatureFee
	Erc721Token      common.Address
	Erc721TokenId    *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterERC721SellOrderFilled is a free log retrieval operation binding the contract event 0x9c248aa1a265aa616f707b979d57f4529bb63a4fc34dc7fc61fdddc18410f74e.
//
// Solidity: event ERC721SellOrderFilled(bytes32 orderHash, address maker, address taker, uint256 nonce, address erc20Token, uint256 erc20TokenAmount, (address,uint256)[] fees, address erc721Token, uint256 erc721TokenId)
func (_ERC721OrdersFeature *ERC721OrdersFeatureFilterer) FilterERC721SellOrderFilled(opts *bind.FilterOpts) (*ERC721OrdersFeatureERC721SellOrderFilledIterator, error) {

	logs, sub, err := _ERC721OrdersFeature.contract.FilterLogs(opts, "ERC721SellOrderFilled")
	if err != nil {
		return nil, err
	}
	return &ERC721OrdersFeatureERC721SellOrderFilledIterator{contract: _ERC721OrdersFeature.contract, event: "ERC721SellOrderFilled", logs: logs, sub: sub}, nil
}

// WatchERC721SellOrderFilled is a free log subscription operation binding the contract event 0x9c248aa1a265aa616f707b979d57f4529bb63a4fc34dc7fc61fdddc18410f74e.
//
// Solidity: event ERC721SellOrderFilled(bytes32 orderHash, address maker, address taker, uint256 nonce, address erc20Token, uint256 erc20TokenAmount, (address,uint256)[] fees, address erc721Token, uint256 erc721TokenId)
func (_ERC721OrdersFeature *ERC721OrdersFeatureFilterer) WatchERC721SellOrderFilled(opts *bind.WatchOpts, sink chan<- *ERC721OrdersFeatureERC721SellOrderFilled) (event.Subscription, error) {

	logs, sub, err := _ERC721OrdersFeature.contract.WatchLogs(opts, "ERC721SellOrderFilled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721OrdersFeatureERC721SellOrderFilled)
				if err := _ERC721OrdersFeature.contract.UnpackLog(event, "ERC721SellOrderFilled", log); err != nil {
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

// ParseERC721SellOrderFilled is a log parse operation binding the contract event 0x9c248aa1a265aa616f707b979d57f4529bb63a4fc34dc7fc61fdddc18410f74e.
//
// Solidity: event ERC721SellOrderFilled(bytes32 orderHash, address maker, address taker, uint256 nonce, address erc20Token, uint256 erc20TokenAmount, (address,uint256)[] fees, address erc721Token, uint256 erc721TokenId)
func (_ERC721OrdersFeature *ERC721OrdersFeatureFilterer) ParseERC721SellOrderFilled(log types.Log) (*ERC721OrdersFeatureERC721SellOrderFilled, error) {
	event := new(ERC721OrdersFeatureERC721SellOrderFilled)
	if err := _ERC721OrdersFeature.contract.UnpackLog(event, "ERC721SellOrderFilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC721OrdersFeatureERC721SellOrderPreSignedIterator is returned from FilterERC721SellOrderPreSigned and is used to iterate over the raw logs and unpacked data for ERC721SellOrderPreSigned events raised by the ERC721OrdersFeature contract.
type ERC721OrdersFeatureERC721SellOrderPreSignedIterator struct {
	Event *ERC721OrdersFeatureERC721SellOrderPreSigned // Event containing the contract specifics and raw log

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
func (it *ERC721OrdersFeatureERC721SellOrderPreSignedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721OrdersFeatureERC721SellOrderPreSigned)
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
		it.Event = new(ERC721OrdersFeatureERC721SellOrderPreSigned)
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
func (it *ERC721OrdersFeatureERC721SellOrderPreSignedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721OrdersFeatureERC721SellOrderPreSignedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721OrdersFeatureERC721SellOrderPreSigned represents a ERC721SellOrderPreSigned event raised by the ERC721OrdersFeature contract.
type ERC721OrdersFeatureERC721SellOrderPreSigned struct {
	Maker            common.Address
	Taker            common.Address
	Expiry           *big.Int
	Nonce            *big.Int
	Erc20Token       common.Address
	Erc20TokenAmount *big.Int
	Fees             []LibNFTOrderFee
	Erc721Token      common.Address
	Erc721TokenId    *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterERC721SellOrderPreSigned is a free log retrieval operation binding the contract event 0x29806076879d6116f3a8b8f81980ee6273d4ae8cb3ede88be4bb96f88787c26c.
//
// Solidity: event ERC721SellOrderPreSigned(address maker, address taker, uint256 expiry, uint256 nonce, address erc20Token, uint256 erc20TokenAmount, (address,uint256,bytes)[] fees, address erc721Token, uint256 erc721TokenId)
func (_ERC721OrdersFeature *ERC721OrdersFeatureFilterer) FilterERC721SellOrderPreSigned(opts *bind.FilterOpts) (*ERC721OrdersFeatureERC721SellOrderPreSignedIterator, error) {

	logs, sub, err := _ERC721OrdersFeature.contract.FilterLogs(opts, "ERC721SellOrderPreSigned")
	if err != nil {
		return nil, err
	}
	return &ERC721OrdersFeatureERC721SellOrderPreSignedIterator{contract: _ERC721OrdersFeature.contract, event: "ERC721SellOrderPreSigned", logs: logs, sub: sub}, nil
}

// WatchERC721SellOrderPreSigned is a free log subscription operation binding the contract event 0x29806076879d6116f3a8b8f81980ee6273d4ae8cb3ede88be4bb96f88787c26c.
//
// Solidity: event ERC721SellOrderPreSigned(address maker, address taker, uint256 expiry, uint256 nonce, address erc20Token, uint256 erc20TokenAmount, (address,uint256,bytes)[] fees, address erc721Token, uint256 erc721TokenId)
func (_ERC721OrdersFeature *ERC721OrdersFeatureFilterer) WatchERC721SellOrderPreSigned(opts *bind.WatchOpts, sink chan<- *ERC721OrdersFeatureERC721SellOrderPreSigned) (event.Subscription, error) {

	logs, sub, err := _ERC721OrdersFeature.contract.WatchLogs(opts, "ERC721SellOrderPreSigned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721OrdersFeatureERC721SellOrderPreSigned)
				if err := _ERC721OrdersFeature.contract.UnpackLog(event, "ERC721SellOrderPreSigned", log); err != nil {
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

// ParseERC721SellOrderPreSigned is a log parse operation binding the contract event 0x29806076879d6116f3a8b8f81980ee6273d4ae8cb3ede88be4bb96f88787c26c.
//
// Solidity: event ERC721SellOrderPreSigned(address maker, address taker, uint256 expiry, uint256 nonce, address erc20Token, uint256 erc20TokenAmount, (address,uint256,bytes)[] fees, address erc721Token, uint256 erc721TokenId)
func (_ERC721OrdersFeature *ERC721OrdersFeatureFilterer) ParseERC721SellOrderPreSigned(log types.Log) (*ERC721OrdersFeatureERC721SellOrderPreSigned, error) {
	event := new(ERC721OrdersFeatureERC721SellOrderPreSigned)
	if err := _ERC721OrdersFeature.contract.UnpackLog(event, "ERC721SellOrderPreSigned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC721OrdersFeatureHashNonceIncrementedIterator is returned from FilterHashNonceIncremented and is used to iterate over the raw logs and unpacked data for HashNonceIncremented events raised by the ERC721OrdersFeature contract.
type ERC721OrdersFeatureHashNonceIncrementedIterator struct {
	Event *ERC721OrdersFeatureHashNonceIncremented // Event containing the contract specifics and raw log

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
func (it *ERC721OrdersFeatureHashNonceIncrementedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721OrdersFeatureHashNonceIncremented)
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
		it.Event = new(ERC721OrdersFeatureHashNonceIncremented)
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
func (it *ERC721OrdersFeatureHashNonceIncrementedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721OrdersFeatureHashNonceIncrementedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721OrdersFeatureHashNonceIncremented represents a HashNonceIncremented event raised by the ERC721OrdersFeature contract.
type ERC721OrdersFeatureHashNonceIncremented struct {
	Maker        common.Address
	NewHashNonce *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterHashNonceIncremented is a free log retrieval operation binding the contract event 0x4cf3e8a83c6bf8a510613208458629675b4ae99b8029e3ab6cb6a86e5f01fd31.
//
// Solidity: event HashNonceIncremented(address maker, uint256 newHashNonce)
func (_ERC721OrdersFeature *ERC721OrdersFeatureFilterer) FilterHashNonceIncremented(opts *bind.FilterOpts) (*ERC721OrdersFeatureHashNonceIncrementedIterator, error) {

	logs, sub, err := _ERC721OrdersFeature.contract.FilterLogs(opts, "HashNonceIncremented")
	if err != nil {
		return nil, err
	}
	return &ERC721OrdersFeatureHashNonceIncrementedIterator{contract: _ERC721OrdersFeature.contract, event: "HashNonceIncremented", logs: logs, sub: sub}, nil
}

// WatchHashNonceIncremented is a free log subscription operation binding the contract event 0x4cf3e8a83c6bf8a510613208458629675b4ae99b8029e3ab6cb6a86e5f01fd31.
//
// Solidity: event HashNonceIncremented(address maker, uint256 newHashNonce)
func (_ERC721OrdersFeature *ERC721OrdersFeatureFilterer) WatchHashNonceIncremented(opts *bind.WatchOpts, sink chan<- *ERC721OrdersFeatureHashNonceIncremented) (event.Subscription, error) {

	logs, sub, err := _ERC721OrdersFeature.contract.WatchLogs(opts, "HashNonceIncremented")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721OrdersFeatureHashNonceIncremented)
				if err := _ERC721OrdersFeature.contract.UnpackLog(event, "HashNonceIncremented", log); err != nil {
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

// ParseHashNonceIncremented is a log parse operation binding the contract event 0x4cf3e8a83c6bf8a510613208458629675b4ae99b8029e3ab6cb6a86e5f01fd31.
//
// Solidity: event HashNonceIncremented(address maker, uint256 newHashNonce)
func (_ERC721OrdersFeature *ERC721OrdersFeatureFilterer) ParseHashNonceIncremented(log types.Log) (*ERC721OrdersFeatureHashNonceIncremented, error) {
	event := new(ERC721OrdersFeatureHashNonceIncremented)
	if err := _ERC721OrdersFeature.contract.UnpackLog(event, "HashNonceIncremented", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC721OrdersFeatureTakerDataEmittedIterator is returned from FilterTakerDataEmitted and is used to iterate over the raw logs and unpacked data for TakerDataEmitted events raised by the ERC721OrdersFeature contract.
type ERC721OrdersFeatureTakerDataEmittedIterator struct {
	Event *ERC721OrdersFeatureTakerDataEmitted // Event containing the contract specifics and raw log

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
func (it *ERC721OrdersFeatureTakerDataEmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721OrdersFeatureTakerDataEmitted)
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
		it.Event = new(ERC721OrdersFeatureTakerDataEmitted)
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
func (it *ERC721OrdersFeatureTakerDataEmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721OrdersFeatureTakerDataEmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721OrdersFeatureTakerDataEmitted represents a TakerDataEmitted event raised by the ERC721OrdersFeature contract.
type ERC721OrdersFeatureTakerDataEmitted struct {
	OrderHash [32]byte
	TakerData []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTakerDataEmitted is a free log retrieval operation binding the contract event 0xf61c2baa48a3c53b619a8e6c2bb6d677f82466c61970ee6afd4a157b0fbf8756.
//
// Solidity: event TakerDataEmitted(bytes32 orderHash, bytes takerData)
func (_ERC721OrdersFeature *ERC721OrdersFeatureFilterer) FilterTakerDataEmitted(opts *bind.FilterOpts) (*ERC721OrdersFeatureTakerDataEmittedIterator, error) {

	logs, sub, err := _ERC721OrdersFeature.contract.FilterLogs(opts, "TakerDataEmitted")
	if err != nil {
		return nil, err
	}
	return &ERC721OrdersFeatureTakerDataEmittedIterator{contract: _ERC721OrdersFeature.contract, event: "TakerDataEmitted", logs: logs, sub: sub}, nil
}

// WatchTakerDataEmitted is a free log subscription operation binding the contract event 0xf61c2baa48a3c53b619a8e6c2bb6d677f82466c61970ee6afd4a157b0fbf8756.
//
// Solidity: event TakerDataEmitted(bytes32 orderHash, bytes takerData)
func (_ERC721OrdersFeature *ERC721OrdersFeatureFilterer) WatchTakerDataEmitted(opts *bind.WatchOpts, sink chan<- *ERC721OrdersFeatureTakerDataEmitted) (event.Subscription, error) {

	logs, sub, err := _ERC721OrdersFeature.contract.WatchLogs(opts, "TakerDataEmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721OrdersFeatureTakerDataEmitted)
				if err := _ERC721OrdersFeature.contract.UnpackLog(event, "TakerDataEmitted", log); err != nil {
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

// ParseTakerDataEmitted is a log parse operation binding the contract event 0xf61c2baa48a3c53b619a8e6c2bb6d677f82466c61970ee6afd4a157b0fbf8756.
//
// Solidity: event TakerDataEmitted(bytes32 orderHash, bytes takerData)
func (_ERC721OrdersFeature *ERC721OrdersFeatureFilterer) ParseTakerDataEmitted(log types.Log) (*ERC721OrdersFeatureTakerDataEmitted, error) {
	event := new(ERC721OrdersFeatureTakerDataEmitted)
	if err := _ERC721OrdersFeature.contract.UnpackLog(event, "TakerDataEmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
