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

// LibNFTOrderERC1155BuyOrder is an auto generated low-level Go binding around an user-defined struct.
type LibNFTOrderERC1155BuyOrder struct {
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

// LibNFTOrderERC1155SellOrder is an auto generated low-level Go binding around an user-defined struct.
type LibNFTOrderERC1155SellOrder struct {
	Maker              common.Address
	Taker              common.Address
	Expiry             *big.Int
	Nonce              *big.Int
	Erc20Token         common.Address
	Erc20TokenAmount   *big.Int
	Fees               []LibNFTOrderFee
	Erc1155Token       common.Address
	Erc1155TokenId     *big.Int
	Erc1155TokenAmount *big.Int
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
	BuyAmount    *big.Int
	EthAvailable *big.Int
	Taker        common.Address
	TakerData    []byte
}

// ERC1155OrdersFeatureMetaData contains all meta data concerning the ERC1155OrdersFeature contract.
var ERC1155OrdersFeatureMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIEtherToken\",\"name\":\"weth\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"erc20FillAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structINFTOrdersFeature.Fee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"erc1155Token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"erc1155TokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"erc1155FillAmount\",\"type\":\"uint128\"}],\"name\":\"ERC1155BuyOrderFilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"erc20TokenAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structLibNFTOrder.Fee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"erc1155Token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"erc1155TokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractIPropertyValidator\",\"name\":\"propertyValidator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"propertyData\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structLibNFTOrder.Property[]\",\"name\":\"erc1155TokenProperties\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"erc1155TokenAmount\",\"type\":\"uint128\"}],\"name\":\"ERC1155BuyOrderPreSigned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"ERC1155OrderCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"erc20FillAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structINFTOrdersFeature.Fee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"erc1155Token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"erc1155TokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"erc1155FillAmount\",\"type\":\"uint128\"}],\"name\":\"ERC1155SellOrderFilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"erc20TokenAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structLibNFTOrder.Fee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"erc1155Token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"erc1155TokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"erc1155TokenAmount\",\"type\":\"uint128\"}],\"name\":\"ERC1155SellOrderPreSigned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"takerData\",\"type\":\"bytes\"}],\"name\":\"TakerDataEmitted\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc20TokenAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Fee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"erc1155Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc1155TokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"erc1155TokenAmount\",\"type\":\"uint128\"}],\"internalType\":\"structLibNFTOrder.ERC1155SellOrder[]\",\"name\":\"sellOrders\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumLibSignature.SignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structLibSignature.Signature[]\",\"name\":\"signatures\",\"type\":\"tuple[]\"},{\"internalType\":\"uint128[]\",\"name\":\"erc1155FillAmounts\",\"type\":\"uint128[]\"},{\"internalType\":\"bool\",\"name\":\"revertIfIncomplete\",\"type\":\"bool\"}],\"name\":\"batchBuyERC1155s\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"successes\",\"type\":\"bool[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc20TokenAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Fee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"erc1155Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc1155TokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"erc1155TokenAmount\",\"type\":\"uint128\"}],\"internalType\":\"structLibNFTOrder.ERC1155SellOrder[]\",\"name\":\"sellOrders\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumLibSignature.SignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structLibSignature.Signature[]\",\"name\":\"signatures\",\"type\":\"tuple[]\"},{\"internalType\":\"address[]\",\"name\":\"takers\",\"type\":\"address[]\"},{\"internalType\":\"uint128[]\",\"name\":\"erc1155FillAmounts\",\"type\":\"uint128[]\"},{\"internalType\":\"bytes[]\",\"name\":\"takerDatas\",\"type\":\"bytes[]\"},{\"internalType\":\"bool\",\"name\":\"revertIfIncomplete\",\"type\":\"bool\"}],\"name\":\"batchBuyERC1155sEx\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"successes\",\"type\":\"bool[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"orderNonces\",\"type\":\"uint256[]\"}],\"name\":\"batchCancelERC1155Orders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc20TokenAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Fee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"erc1155Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc1155TokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"erc1155TokenAmount\",\"type\":\"uint128\"}],\"internalType\":\"structLibNFTOrder.ERC1155SellOrder[]\",\"name\":\"sellOrders\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc20TokenAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Fee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"erc1155Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc1155TokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractIPropertyValidator\",\"name\":\"propertyValidator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"propertyData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Property[]\",\"name\":\"erc1155TokenProperties\",\"type\":\"tuple[]\"},{\"internalType\":\"uint128\",\"name\":\"erc1155TokenAmount\",\"type\":\"uint128\"}],\"internalType\":\"structLibNFTOrder.ERC1155BuyOrder[]\",\"name\":\"buyOrders\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumLibSignature.SignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structLibSignature.Signature[]\",\"name\":\"sellOrderSignatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumLibSignature.SignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structLibSignature.Signature[]\",\"name\":\"buyOrderSignatures\",\"type\":\"tuple[]\"}],\"name\":\"batchMatchERC1155Orders\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"profits\",\"type\":\"uint256[]\"},{\"internalType\":\"bool[]\",\"name\":\"successes\",\"type\":\"bool[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc20TokenAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Fee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"erc1155Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc1155TokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"erc1155TokenAmount\",\"type\":\"uint128\"}],\"internalType\":\"structLibNFTOrder.ERC1155SellOrder\",\"name\":\"sellOrder\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumLibSignature.SignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structLibSignature.Signature\",\"name\":\"signature\",\"type\":\"tuple\"},{\"internalType\":\"uint128\",\"name\":\"erc1155BuyAmount\",\"type\":\"uint128\"}],\"name\":\"buyERC1155\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc20TokenAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Fee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"erc1155Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc1155TokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"erc1155TokenAmount\",\"type\":\"uint128\"}],\"internalType\":\"structLibNFTOrder.ERC1155SellOrder\",\"name\":\"sellOrder\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumLibSignature.SignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structLibSignature.Signature\",\"name\":\"signature\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"erc1155BuyAmount\",\"type\":\"uint128\"},{\"internalType\":\"bytes\",\"name\":\"takerData\",\"type\":\"bytes\"}],\"name\":\"buyERC1155Ex\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc20TokenAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Fee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"erc1155Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc1155TokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"erc1155TokenAmount\",\"type\":\"uint128\"}],\"internalType\":\"structLibNFTOrder.ERC1155SellOrder\",\"name\":\"sellOrder\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumLibSignature.SignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structLibSignature.Signature\",\"name\":\"signature\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"buyAmount\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"ethAvailable\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"takerData\",\"type\":\"bytes\"}],\"internalType\":\"structNFTOrders.BuyParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"buyERC1155ExFromProxy\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc20TokenAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Fee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"erc1155Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc1155TokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"erc1155TokenAmount\",\"type\":\"uint128\"}],\"internalType\":\"structLibNFTOrder.ERC1155SellOrder\",\"name\":\"sellOrder\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumLibSignature.SignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structLibSignature.Signature\",\"name\":\"signature\",\"type\":\"tuple\"},{\"internalType\":\"uint128\",\"name\":\"buyAmount\",\"type\":\"uint128\"}],\"name\":\"buyERC1155FromProxy\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderNonce\",\"type\":\"uint256\"}],\"name\":\"cancelERC1155Order\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc20TokenAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Fee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"erc1155Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc1155TokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractIPropertyValidator\",\"name\":\"propertyValidator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"propertyData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Property[]\",\"name\":\"erc1155TokenProperties\",\"type\":\"tuple[]\"},{\"internalType\":\"uint128\",\"name\":\"erc1155TokenAmount\",\"type\":\"uint128\"}],\"internalType\":\"structLibNFTOrder.ERC1155BuyOrder\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"getERC1155BuyOrderHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc20TokenAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Fee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"erc1155Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc1155TokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractIPropertyValidator\",\"name\":\"propertyValidator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"propertyData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Property[]\",\"name\":\"erc1155TokenProperties\",\"type\":\"tuple[]\"},{\"internalType\":\"uint128\",\"name\":\"erc1155TokenAmount\",\"type\":\"uint128\"}],\"internalType\":\"structLibNFTOrder.ERC1155BuyOrder\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"getERC1155BuyOrderInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"internalType\":\"enumLibNFTOrder.OrderStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint128\",\"name\":\"orderAmount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"remainingAmount\",\"type\":\"uint128\"}],\"internalType\":\"structLibNFTOrder.OrderInfo\",\"name\":\"orderInfo\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"uint248\",\"name\":\"nonceRange\",\"type\":\"uint248\"}],\"name\":\"getERC1155OrderNonceStatusBitVector\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc20TokenAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Fee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"erc1155Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc1155TokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"erc1155TokenAmount\",\"type\":\"uint128\"}],\"internalType\":\"structLibNFTOrder.ERC1155SellOrder\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"getERC1155SellOrderHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc20TokenAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Fee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"erc1155Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc1155TokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"erc1155TokenAmount\",\"type\":\"uint128\"}],\"internalType\":\"structLibNFTOrder.ERC1155SellOrder\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"getERC1155SellOrderInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"internalType\":\"enumLibNFTOrder.OrderStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint128\",\"name\":\"orderAmount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"remainingAmount\",\"type\":\"uint128\"}],\"internalType\":\"structLibNFTOrder.OrderInfo\",\"name\":\"orderInfo\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc20TokenAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Fee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"erc1155Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc1155TokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"erc1155TokenAmount\",\"type\":\"uint128\"}],\"internalType\":\"structLibNFTOrder.ERC1155SellOrder\",\"name\":\"sellOrder\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc20TokenAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Fee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"erc1155Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc1155TokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractIPropertyValidator\",\"name\":\"propertyValidator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"propertyData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Property[]\",\"name\":\"erc1155TokenProperties\",\"type\":\"tuple[]\"},{\"internalType\":\"uint128\",\"name\":\"erc1155TokenAmount\",\"type\":\"uint128\"}],\"internalType\":\"structLibNFTOrder.ERC1155BuyOrder\",\"name\":\"buyOrder\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumLibSignature.SignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structLibSignature.Signature\",\"name\":\"sellOrderSignature\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumLibSignature.SignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structLibSignature.Signature\",\"name\":\"buyOrderSignature\",\"type\":\"tuple\"}],\"name\":\"matchERC1155Orders\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"profit\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"success\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc20TokenAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Fee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"erc1155Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc1155TokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractIPropertyValidator\",\"name\":\"propertyValidator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"propertyData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Property[]\",\"name\":\"erc1155TokenProperties\",\"type\":\"tuple[]\"},{\"internalType\":\"uint128\",\"name\":\"erc1155TokenAmount\",\"type\":\"uint128\"}],\"internalType\":\"structLibNFTOrder.ERC1155BuyOrder\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"preSignERC1155BuyOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc20TokenAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Fee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"erc1155Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc1155TokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"erc1155TokenAmount\",\"type\":\"uint128\"}],\"internalType\":\"structLibNFTOrder.ERC1155SellOrder\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"preSignERC1155SellOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc20TokenAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Fee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"erc1155Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc1155TokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractIPropertyValidator\",\"name\":\"propertyValidator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"propertyData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Property[]\",\"name\":\"erc1155TokenProperties\",\"type\":\"tuple[]\"},{\"internalType\":\"uint128\",\"name\":\"erc1155TokenAmount\",\"type\":\"uint128\"}],\"internalType\":\"structLibNFTOrder.ERC1155BuyOrder\",\"name\":\"buyOrder\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumLibSignature.SignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structLibSignature.Signature\",\"name\":\"signature\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"erc1155TokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"erc1155SellAmount\",\"type\":\"uint128\"},{\"internalType\":\"bool\",\"name\":\"unwrapNativeToken\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"takerData\",\"type\":\"bytes\"}],\"name\":\"sellERC1155\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc20TokenAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Fee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"erc1155Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc1155TokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractIPropertyValidator\",\"name\":\"propertyValidator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"propertyData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Property[]\",\"name\":\"erc1155TokenProperties\",\"type\":\"tuple[]\"},{\"internalType\":\"uint128\",\"name\":\"erc1155TokenAmount\",\"type\":\"uint128\"}],\"internalType\":\"structLibNFTOrder.ERC1155BuyOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumLibSignature.SignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structLibSignature.Signature\",\"name\":\"signature\",\"type\":\"tuple\"}],\"name\":\"validateERC1155BuyOrderSignature\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc20TokenAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"internalType\":\"structLibNFTOrder.Fee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"erc1155Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc1155TokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"erc1155TokenAmount\",\"type\":\"uint128\"}],\"internalType\":\"structLibNFTOrder.ERC1155SellOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumLibSignature.SignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structLibSignature.Signature\",\"name\":\"signature\",\"type\":\"tuple\"}],\"name\":\"validateERC1155SellOrderSignature\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

// GetERC1155BuyOrderHash is a free data retrieval call binding the contract method 0x274ed32e.
//
// Solidity: function getERC1155BuyOrderHash((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[],uint128) order) view returns(bytes32 orderHash)
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureCaller) GetERC1155BuyOrderHash(opts *bind.CallOpts, order LibNFTOrderERC1155BuyOrder) ([32]byte, error) {
	var out []interface{}
	err := _ERC1155OrdersFeature.contract.Call(opts, &out, "getERC1155BuyOrderHash", order)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetERC1155BuyOrderHash is a free data retrieval call binding the contract method 0x274ed32e.
//
// Solidity: function getERC1155BuyOrderHash((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[],uint128) order) view returns(bytes32 orderHash)
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureSession) GetERC1155BuyOrderHash(order LibNFTOrderERC1155BuyOrder) ([32]byte, error) {
	return _ERC1155OrdersFeature.Contract.GetERC1155BuyOrderHash(&_ERC1155OrdersFeature.CallOpts, order)
}

// GetERC1155BuyOrderHash is a free data retrieval call binding the contract method 0x274ed32e.
//
// Solidity: function getERC1155BuyOrderHash((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[],uint128) order) view returns(bytes32 orderHash)
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureCallerSession) GetERC1155BuyOrderHash(order LibNFTOrderERC1155BuyOrder) ([32]byte, error) {
	return _ERC1155OrdersFeature.Contract.GetERC1155BuyOrderHash(&_ERC1155OrdersFeature.CallOpts, order)
}

// GetERC1155BuyOrderInfo is a free data retrieval call binding the contract method 0x894e8ad9.
//
// Solidity: function getERC1155BuyOrderInfo((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[],uint128) order) view returns((bytes32,uint8,uint128,uint128) orderInfo)
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureCaller) GetERC1155BuyOrderInfo(opts *bind.CallOpts, order LibNFTOrderERC1155BuyOrder) (LibNFTOrderOrderInfo, error) {
	var out []interface{}
	err := _ERC1155OrdersFeature.contract.Call(opts, &out, "getERC1155BuyOrderInfo", order)

	if err != nil {
		return *new(LibNFTOrderOrderInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(LibNFTOrderOrderInfo)).(*LibNFTOrderOrderInfo)

	return out0, err

}

// GetERC1155BuyOrderInfo is a free data retrieval call binding the contract method 0x894e8ad9.
//
// Solidity: function getERC1155BuyOrderInfo((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[],uint128) order) view returns((bytes32,uint8,uint128,uint128) orderInfo)
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureSession) GetERC1155BuyOrderInfo(order LibNFTOrderERC1155BuyOrder) (LibNFTOrderOrderInfo, error) {
	return _ERC1155OrdersFeature.Contract.GetERC1155BuyOrderInfo(&_ERC1155OrdersFeature.CallOpts, order)
}

// GetERC1155BuyOrderInfo is a free data retrieval call binding the contract method 0x894e8ad9.
//
// Solidity: function getERC1155BuyOrderInfo((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[],uint128) order) view returns((bytes32,uint8,uint128,uint128) orderInfo)
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureCallerSession) GetERC1155BuyOrderInfo(order LibNFTOrderERC1155BuyOrder) (LibNFTOrderOrderInfo, error) {
	return _ERC1155OrdersFeature.Contract.GetERC1155BuyOrderInfo(&_ERC1155OrdersFeature.CallOpts, order)
}

// GetERC1155OrderNonceStatusBitVector is a free data retrieval call binding the contract method 0x2011b06a.
//
// Solidity: function getERC1155OrderNonceStatusBitVector(address maker, uint248 nonceRange) view returns(uint256)
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureCaller) GetERC1155OrderNonceStatusBitVector(opts *bind.CallOpts, maker common.Address, nonceRange *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ERC1155OrdersFeature.contract.Call(opts, &out, "getERC1155OrderNonceStatusBitVector", maker, nonceRange)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetERC1155OrderNonceStatusBitVector is a free data retrieval call binding the contract method 0x2011b06a.
//
// Solidity: function getERC1155OrderNonceStatusBitVector(address maker, uint248 nonceRange) view returns(uint256)
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureSession) GetERC1155OrderNonceStatusBitVector(maker common.Address, nonceRange *big.Int) (*big.Int, error) {
	return _ERC1155OrdersFeature.Contract.GetERC1155OrderNonceStatusBitVector(&_ERC1155OrdersFeature.CallOpts, maker, nonceRange)
}

// GetERC1155OrderNonceStatusBitVector is a free data retrieval call binding the contract method 0x2011b06a.
//
// Solidity: function getERC1155OrderNonceStatusBitVector(address maker, uint248 nonceRange) view returns(uint256)
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureCallerSession) GetERC1155OrderNonceStatusBitVector(maker common.Address, nonceRange *big.Int) (*big.Int, error) {
	return _ERC1155OrdersFeature.Contract.GetERC1155OrderNonceStatusBitVector(&_ERC1155OrdersFeature.CallOpts, maker, nonceRange)
}

// GetERC1155SellOrderHash is a free data retrieval call binding the contract method 0x5e0ccb6b.
//
// Solidity: function getERC1155SellOrderHash((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,uint128) order) view returns(bytes32 orderHash)
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureCaller) GetERC1155SellOrderHash(opts *bind.CallOpts, order LibNFTOrderERC1155SellOrder) ([32]byte, error) {
	var out []interface{}
	err := _ERC1155OrdersFeature.contract.Call(opts, &out, "getERC1155SellOrderHash", order)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetERC1155SellOrderHash is a free data retrieval call binding the contract method 0x5e0ccb6b.
//
// Solidity: function getERC1155SellOrderHash((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,uint128) order) view returns(bytes32 orderHash)
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureSession) GetERC1155SellOrderHash(order LibNFTOrderERC1155SellOrder) ([32]byte, error) {
	return _ERC1155OrdersFeature.Contract.GetERC1155SellOrderHash(&_ERC1155OrdersFeature.CallOpts, order)
}

// GetERC1155SellOrderHash is a free data retrieval call binding the contract method 0x5e0ccb6b.
//
// Solidity: function getERC1155SellOrderHash((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,uint128) order) view returns(bytes32 orderHash)
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureCallerSession) GetERC1155SellOrderHash(order LibNFTOrderERC1155SellOrder) ([32]byte, error) {
	return _ERC1155OrdersFeature.Contract.GetERC1155SellOrderHash(&_ERC1155OrdersFeature.CallOpts, order)
}

// GetERC1155SellOrderInfo is a free data retrieval call binding the contract method 0xbd7ad3e5.
//
// Solidity: function getERC1155SellOrderInfo((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,uint128) order) view returns((bytes32,uint8,uint128,uint128) orderInfo)
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureCaller) GetERC1155SellOrderInfo(opts *bind.CallOpts, order LibNFTOrderERC1155SellOrder) (LibNFTOrderOrderInfo, error) {
	var out []interface{}
	err := _ERC1155OrdersFeature.contract.Call(opts, &out, "getERC1155SellOrderInfo", order)

	if err != nil {
		return *new(LibNFTOrderOrderInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(LibNFTOrderOrderInfo)).(*LibNFTOrderOrderInfo)

	return out0, err

}

// GetERC1155SellOrderInfo is a free data retrieval call binding the contract method 0xbd7ad3e5.
//
// Solidity: function getERC1155SellOrderInfo((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,uint128) order) view returns((bytes32,uint8,uint128,uint128) orderInfo)
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureSession) GetERC1155SellOrderInfo(order LibNFTOrderERC1155SellOrder) (LibNFTOrderOrderInfo, error) {
	return _ERC1155OrdersFeature.Contract.GetERC1155SellOrderInfo(&_ERC1155OrdersFeature.CallOpts, order)
}

// GetERC1155SellOrderInfo is a free data retrieval call binding the contract method 0xbd7ad3e5.
//
// Solidity: function getERC1155SellOrderInfo((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,uint128) order) view returns((bytes32,uint8,uint128,uint128) orderInfo)
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureCallerSession) GetERC1155SellOrderInfo(order LibNFTOrderERC1155SellOrder) (LibNFTOrderOrderInfo, error) {
	return _ERC1155OrdersFeature.Contract.GetERC1155SellOrderInfo(&_ERC1155OrdersFeature.CallOpts, order)
}

// ValidateERC1155BuyOrderSignature is a free data retrieval call binding the contract method 0x28cdc923.
//
// Solidity: function validateERC1155BuyOrderSignature((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[],uint128) order, (uint8,uint8,bytes32,bytes32) signature) view returns()
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureCaller) ValidateERC1155BuyOrderSignature(opts *bind.CallOpts, order LibNFTOrderERC1155BuyOrder, signature LibSignatureSignature) error {
	var out []interface{}
	err := _ERC1155OrdersFeature.contract.Call(opts, &out, "validateERC1155BuyOrderSignature", order, signature)

	if err != nil {
		return err
	}

	return err

}

// ValidateERC1155BuyOrderSignature is a free data retrieval call binding the contract method 0x28cdc923.
//
// Solidity: function validateERC1155BuyOrderSignature((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[],uint128) order, (uint8,uint8,bytes32,bytes32) signature) view returns()
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureSession) ValidateERC1155BuyOrderSignature(order LibNFTOrderERC1155BuyOrder, signature LibSignatureSignature) error {
	return _ERC1155OrdersFeature.Contract.ValidateERC1155BuyOrderSignature(&_ERC1155OrdersFeature.CallOpts, order, signature)
}

// ValidateERC1155BuyOrderSignature is a free data retrieval call binding the contract method 0x28cdc923.
//
// Solidity: function validateERC1155BuyOrderSignature((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[],uint128) order, (uint8,uint8,bytes32,bytes32) signature) view returns()
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureCallerSession) ValidateERC1155BuyOrderSignature(order LibNFTOrderERC1155BuyOrder, signature LibSignatureSignature) error {
	return _ERC1155OrdersFeature.Contract.ValidateERC1155BuyOrderSignature(&_ERC1155OrdersFeature.CallOpts, order, signature)
}

// ValidateERC1155SellOrderSignature is a free data retrieval call binding the contract method 0x88df8e9e.
//
// Solidity: function validateERC1155SellOrderSignature((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,uint128) order, (uint8,uint8,bytes32,bytes32) signature) view returns()
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureCaller) ValidateERC1155SellOrderSignature(opts *bind.CallOpts, order LibNFTOrderERC1155SellOrder, signature LibSignatureSignature) error {
	var out []interface{}
	err := _ERC1155OrdersFeature.contract.Call(opts, &out, "validateERC1155SellOrderSignature", order, signature)

	if err != nil {
		return err
	}

	return err

}

// ValidateERC1155SellOrderSignature is a free data retrieval call binding the contract method 0x88df8e9e.
//
// Solidity: function validateERC1155SellOrderSignature((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,uint128) order, (uint8,uint8,bytes32,bytes32) signature) view returns()
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureSession) ValidateERC1155SellOrderSignature(order LibNFTOrderERC1155SellOrder, signature LibSignatureSignature) error {
	return _ERC1155OrdersFeature.Contract.ValidateERC1155SellOrderSignature(&_ERC1155OrdersFeature.CallOpts, order, signature)
}

// ValidateERC1155SellOrderSignature is a free data retrieval call binding the contract method 0x88df8e9e.
//
// Solidity: function validateERC1155SellOrderSignature((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,uint128) order, (uint8,uint8,bytes32,bytes32) signature) view returns()
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureCallerSession) ValidateERC1155SellOrderSignature(order LibNFTOrderERC1155SellOrder, signature LibSignatureSignature) error {
	return _ERC1155OrdersFeature.Contract.ValidateERC1155SellOrderSignature(&_ERC1155OrdersFeature.CallOpts, order, signature)
}

// BatchBuyERC1155s is a paid mutator transaction binding the contract method 0x1ff0bd99.
//
// Solidity: function batchBuyERC1155s((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,uint128)[] sellOrders, (uint8,uint8,bytes32,bytes32)[] signatures, uint128[] erc1155FillAmounts, bool revertIfIncomplete) payable returns(bool[] successes)
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureTransactor) BatchBuyERC1155s(opts *bind.TransactOpts, sellOrders []LibNFTOrderERC1155SellOrder, signatures []LibSignatureSignature, erc1155FillAmounts []*big.Int, revertIfIncomplete bool) (*types.Transaction, error) {
	return _ERC1155OrdersFeature.contract.Transact(opts, "batchBuyERC1155s", sellOrders, signatures, erc1155FillAmounts, revertIfIncomplete)
}

// BatchBuyERC1155s is a paid mutator transaction binding the contract method 0x1ff0bd99.
//
// Solidity: function batchBuyERC1155s((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,uint128)[] sellOrders, (uint8,uint8,bytes32,bytes32)[] signatures, uint128[] erc1155FillAmounts, bool revertIfIncomplete) payable returns(bool[] successes)
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureSession) BatchBuyERC1155s(sellOrders []LibNFTOrderERC1155SellOrder, signatures []LibSignatureSignature, erc1155FillAmounts []*big.Int, revertIfIncomplete bool) (*types.Transaction, error) {
	return _ERC1155OrdersFeature.Contract.BatchBuyERC1155s(&_ERC1155OrdersFeature.TransactOpts, sellOrders, signatures, erc1155FillAmounts, revertIfIncomplete)
}

// BatchBuyERC1155s is a paid mutator transaction binding the contract method 0x1ff0bd99.
//
// Solidity: function batchBuyERC1155s((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,uint128)[] sellOrders, (uint8,uint8,bytes32,bytes32)[] signatures, uint128[] erc1155FillAmounts, bool revertIfIncomplete) payable returns(bool[] successes)
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureTransactorSession) BatchBuyERC1155s(sellOrders []LibNFTOrderERC1155SellOrder, signatures []LibSignatureSignature, erc1155FillAmounts []*big.Int, revertIfIncomplete bool) (*types.Transaction, error) {
	return _ERC1155OrdersFeature.Contract.BatchBuyERC1155s(&_ERC1155OrdersFeature.TransactOpts, sellOrders, signatures, erc1155FillAmounts, revertIfIncomplete)
}

// BatchBuyERC1155sEx is a paid mutator transaction binding the contract method 0x6623a502.
//
// Solidity: function batchBuyERC1155sEx((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,uint128)[] sellOrders, (uint8,uint8,bytes32,bytes32)[] signatures, address[] takers, uint128[] erc1155FillAmounts, bytes[] takerDatas, bool revertIfIncomplete) payable returns(bool[] successes)
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureTransactor) BatchBuyERC1155sEx(opts *bind.TransactOpts, sellOrders []LibNFTOrderERC1155SellOrder, signatures []LibSignatureSignature, takers []common.Address, erc1155FillAmounts []*big.Int, takerDatas [][]byte, revertIfIncomplete bool) (*types.Transaction, error) {
	return _ERC1155OrdersFeature.contract.Transact(opts, "batchBuyERC1155sEx", sellOrders, signatures, takers, erc1155FillAmounts, takerDatas, revertIfIncomplete)
}

// BatchBuyERC1155sEx is a paid mutator transaction binding the contract method 0x6623a502.
//
// Solidity: function batchBuyERC1155sEx((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,uint128)[] sellOrders, (uint8,uint8,bytes32,bytes32)[] signatures, address[] takers, uint128[] erc1155FillAmounts, bytes[] takerDatas, bool revertIfIncomplete) payable returns(bool[] successes)
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureSession) BatchBuyERC1155sEx(sellOrders []LibNFTOrderERC1155SellOrder, signatures []LibSignatureSignature, takers []common.Address, erc1155FillAmounts []*big.Int, takerDatas [][]byte, revertIfIncomplete bool) (*types.Transaction, error) {
	return _ERC1155OrdersFeature.Contract.BatchBuyERC1155sEx(&_ERC1155OrdersFeature.TransactOpts, sellOrders, signatures, takers, erc1155FillAmounts, takerDatas, revertIfIncomplete)
}

// BatchBuyERC1155sEx is a paid mutator transaction binding the contract method 0x6623a502.
//
// Solidity: function batchBuyERC1155sEx((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,uint128)[] sellOrders, (uint8,uint8,bytes32,bytes32)[] signatures, address[] takers, uint128[] erc1155FillAmounts, bytes[] takerDatas, bool revertIfIncomplete) payable returns(bool[] successes)
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureTransactorSession) BatchBuyERC1155sEx(sellOrders []LibNFTOrderERC1155SellOrder, signatures []LibSignatureSignature, takers []common.Address, erc1155FillAmounts []*big.Int, takerDatas [][]byte, revertIfIncomplete bool) (*types.Transaction, error) {
	return _ERC1155OrdersFeature.Contract.BatchBuyERC1155sEx(&_ERC1155OrdersFeature.TransactOpts, sellOrders, signatures, takers, erc1155FillAmounts, takerDatas, revertIfIncomplete)
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

// BatchMatchERC1155Orders is a paid mutator transaction binding the contract method 0xbb437393.
//
// Solidity: function batchMatchERC1155Orders((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,uint128)[] sellOrders, (address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[],uint128)[] buyOrders, (uint8,uint8,bytes32,bytes32)[] sellOrderSignatures, (uint8,uint8,bytes32,bytes32)[] buyOrderSignatures) returns(uint256[] profits, bool[] successes)
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureTransactor) BatchMatchERC1155Orders(opts *bind.TransactOpts, sellOrders []LibNFTOrderERC1155SellOrder, buyOrders []LibNFTOrderERC1155BuyOrder, sellOrderSignatures []LibSignatureSignature, buyOrderSignatures []LibSignatureSignature) (*types.Transaction, error) {
	return _ERC1155OrdersFeature.contract.Transact(opts, "batchMatchERC1155Orders", sellOrders, buyOrders, sellOrderSignatures, buyOrderSignatures)
}

// BatchMatchERC1155Orders is a paid mutator transaction binding the contract method 0xbb437393.
//
// Solidity: function batchMatchERC1155Orders((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,uint128)[] sellOrders, (address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[],uint128)[] buyOrders, (uint8,uint8,bytes32,bytes32)[] sellOrderSignatures, (uint8,uint8,bytes32,bytes32)[] buyOrderSignatures) returns(uint256[] profits, bool[] successes)
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureSession) BatchMatchERC1155Orders(sellOrders []LibNFTOrderERC1155SellOrder, buyOrders []LibNFTOrderERC1155BuyOrder, sellOrderSignatures []LibSignatureSignature, buyOrderSignatures []LibSignatureSignature) (*types.Transaction, error) {
	return _ERC1155OrdersFeature.Contract.BatchMatchERC1155Orders(&_ERC1155OrdersFeature.TransactOpts, sellOrders, buyOrders, sellOrderSignatures, buyOrderSignatures)
}

// BatchMatchERC1155Orders is a paid mutator transaction binding the contract method 0xbb437393.
//
// Solidity: function batchMatchERC1155Orders((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,uint128)[] sellOrders, (address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[],uint128)[] buyOrders, (uint8,uint8,bytes32,bytes32)[] sellOrderSignatures, (uint8,uint8,bytes32,bytes32)[] buyOrderSignatures) returns(uint256[] profits, bool[] successes)
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureTransactorSession) BatchMatchERC1155Orders(sellOrders []LibNFTOrderERC1155SellOrder, buyOrders []LibNFTOrderERC1155BuyOrder, sellOrderSignatures []LibSignatureSignature, buyOrderSignatures []LibSignatureSignature) (*types.Transaction, error) {
	return _ERC1155OrdersFeature.Contract.BatchMatchERC1155Orders(&_ERC1155OrdersFeature.TransactOpts, sellOrders, buyOrders, sellOrderSignatures, buyOrderSignatures)
}

// BuyERC1155 is a paid mutator transaction binding the contract method 0x418795ab.
//
// Solidity: function buyERC1155((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,uint128) sellOrder, (uint8,uint8,bytes32,bytes32) signature, uint128 erc1155BuyAmount) payable returns()
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureTransactor) BuyERC1155(opts *bind.TransactOpts, sellOrder LibNFTOrderERC1155SellOrder, signature LibSignatureSignature, erc1155BuyAmount *big.Int) (*types.Transaction, error) {
	return _ERC1155OrdersFeature.contract.Transact(opts, "buyERC1155", sellOrder, signature, erc1155BuyAmount)
}

// BuyERC1155 is a paid mutator transaction binding the contract method 0x418795ab.
//
// Solidity: function buyERC1155((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,uint128) sellOrder, (uint8,uint8,bytes32,bytes32) signature, uint128 erc1155BuyAmount) payable returns()
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureSession) BuyERC1155(sellOrder LibNFTOrderERC1155SellOrder, signature LibSignatureSignature, erc1155BuyAmount *big.Int) (*types.Transaction, error) {
	return _ERC1155OrdersFeature.Contract.BuyERC1155(&_ERC1155OrdersFeature.TransactOpts, sellOrder, signature, erc1155BuyAmount)
}

// BuyERC1155 is a paid mutator transaction binding the contract method 0x418795ab.
//
// Solidity: function buyERC1155((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,uint128) sellOrder, (uint8,uint8,bytes32,bytes32) signature, uint128 erc1155BuyAmount) payable returns()
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureTransactorSession) BuyERC1155(sellOrder LibNFTOrderERC1155SellOrder, signature LibSignatureSignature, erc1155BuyAmount *big.Int) (*types.Transaction, error) {
	return _ERC1155OrdersFeature.Contract.BuyERC1155(&_ERC1155OrdersFeature.TransactOpts, sellOrder, signature, erc1155BuyAmount)
}

// BuyERC1155Ex is a paid mutator transaction binding the contract method 0x744773f1.
//
// Solidity: function buyERC1155Ex((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,uint128) sellOrder, (uint8,uint8,bytes32,bytes32) signature, address taker, uint128 erc1155BuyAmount, bytes takerData) payable returns()
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureTransactor) BuyERC1155Ex(opts *bind.TransactOpts, sellOrder LibNFTOrderERC1155SellOrder, signature LibSignatureSignature, taker common.Address, erc1155BuyAmount *big.Int, takerData []byte) (*types.Transaction, error) {
	return _ERC1155OrdersFeature.contract.Transact(opts, "buyERC1155Ex", sellOrder, signature, taker, erc1155BuyAmount, takerData)
}

// BuyERC1155Ex is a paid mutator transaction binding the contract method 0x744773f1.
//
// Solidity: function buyERC1155Ex((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,uint128) sellOrder, (uint8,uint8,bytes32,bytes32) signature, address taker, uint128 erc1155BuyAmount, bytes takerData) payable returns()
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureSession) BuyERC1155Ex(sellOrder LibNFTOrderERC1155SellOrder, signature LibSignatureSignature, taker common.Address, erc1155BuyAmount *big.Int, takerData []byte) (*types.Transaction, error) {
	return _ERC1155OrdersFeature.Contract.BuyERC1155Ex(&_ERC1155OrdersFeature.TransactOpts, sellOrder, signature, taker, erc1155BuyAmount, takerData)
}

// BuyERC1155Ex is a paid mutator transaction binding the contract method 0x744773f1.
//
// Solidity: function buyERC1155Ex((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,uint128) sellOrder, (uint8,uint8,bytes32,bytes32) signature, address taker, uint128 erc1155BuyAmount, bytes takerData) payable returns()
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureTransactorSession) BuyERC1155Ex(sellOrder LibNFTOrderERC1155SellOrder, signature LibSignatureSignature, taker common.Address, erc1155BuyAmount *big.Int, takerData []byte) (*types.Transaction, error) {
	return _ERC1155OrdersFeature.Contract.BuyERC1155Ex(&_ERC1155OrdersFeature.TransactOpts, sellOrder, signature, taker, erc1155BuyAmount, takerData)
}

// BuyERC1155ExFromProxy is a paid mutator transaction binding the contract method 0x013d3a75.
//
// Solidity: function buyERC1155ExFromProxy((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,uint128) sellOrder, (uint8,uint8,bytes32,bytes32) signature, (uint128,uint256,address,bytes) params) payable returns()
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureTransactor) BuyERC1155ExFromProxy(opts *bind.TransactOpts, sellOrder LibNFTOrderERC1155SellOrder, signature LibSignatureSignature, params NFTOrdersBuyParams) (*types.Transaction, error) {
	return _ERC1155OrdersFeature.contract.Transact(opts, "buyERC1155ExFromProxy", sellOrder, signature, params)
}

// BuyERC1155ExFromProxy is a paid mutator transaction binding the contract method 0x013d3a75.
//
// Solidity: function buyERC1155ExFromProxy((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,uint128) sellOrder, (uint8,uint8,bytes32,bytes32) signature, (uint128,uint256,address,bytes) params) payable returns()
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureSession) BuyERC1155ExFromProxy(sellOrder LibNFTOrderERC1155SellOrder, signature LibSignatureSignature, params NFTOrdersBuyParams) (*types.Transaction, error) {
	return _ERC1155OrdersFeature.Contract.BuyERC1155ExFromProxy(&_ERC1155OrdersFeature.TransactOpts, sellOrder, signature, params)
}

// BuyERC1155ExFromProxy is a paid mutator transaction binding the contract method 0x013d3a75.
//
// Solidity: function buyERC1155ExFromProxy((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,uint128) sellOrder, (uint8,uint8,bytes32,bytes32) signature, (uint128,uint256,address,bytes) params) payable returns()
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureTransactorSession) BuyERC1155ExFromProxy(sellOrder LibNFTOrderERC1155SellOrder, signature LibSignatureSignature, params NFTOrdersBuyParams) (*types.Transaction, error) {
	return _ERC1155OrdersFeature.Contract.BuyERC1155ExFromProxy(&_ERC1155OrdersFeature.TransactOpts, sellOrder, signature, params)
}

// BuyERC1155FromProxy is a paid mutator transaction binding the contract method 0xc62b893b.
//
// Solidity: function buyERC1155FromProxy((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,uint128) sellOrder, (uint8,uint8,bytes32,bytes32) signature, uint128 buyAmount) payable returns()
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureTransactor) BuyERC1155FromProxy(opts *bind.TransactOpts, sellOrder LibNFTOrderERC1155SellOrder, signature LibSignatureSignature, buyAmount *big.Int) (*types.Transaction, error) {
	return _ERC1155OrdersFeature.contract.Transact(opts, "buyERC1155FromProxy", sellOrder, signature, buyAmount)
}

// BuyERC1155FromProxy is a paid mutator transaction binding the contract method 0xc62b893b.
//
// Solidity: function buyERC1155FromProxy((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,uint128) sellOrder, (uint8,uint8,bytes32,bytes32) signature, uint128 buyAmount) payable returns()
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureSession) BuyERC1155FromProxy(sellOrder LibNFTOrderERC1155SellOrder, signature LibSignatureSignature, buyAmount *big.Int) (*types.Transaction, error) {
	return _ERC1155OrdersFeature.Contract.BuyERC1155FromProxy(&_ERC1155OrdersFeature.TransactOpts, sellOrder, signature, buyAmount)
}

// BuyERC1155FromProxy is a paid mutator transaction binding the contract method 0xc62b893b.
//
// Solidity: function buyERC1155FromProxy((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,uint128) sellOrder, (uint8,uint8,bytes32,bytes32) signature, uint128 buyAmount) payable returns()
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureTransactorSession) BuyERC1155FromProxy(sellOrder LibNFTOrderERC1155SellOrder, signature LibSignatureSignature, buyAmount *big.Int) (*types.Transaction, error) {
	return _ERC1155OrdersFeature.Contract.BuyERC1155FromProxy(&_ERC1155OrdersFeature.TransactOpts, sellOrder, signature, buyAmount)
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

// MatchERC1155Orders is a paid mutator transaction binding the contract method 0x36a2c310.
//
// Solidity: function matchERC1155Orders((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,uint128) sellOrder, (address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[],uint128) buyOrder, (uint8,uint8,bytes32,bytes32) sellOrderSignature, (uint8,uint8,bytes32,bytes32) buyOrderSignature) returns(uint256 profit)
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureTransactor) MatchERC1155Orders(opts *bind.TransactOpts, sellOrder LibNFTOrderERC1155SellOrder, buyOrder LibNFTOrderERC1155BuyOrder, sellOrderSignature LibSignatureSignature, buyOrderSignature LibSignatureSignature) (*types.Transaction, error) {
	return _ERC1155OrdersFeature.contract.Transact(opts, "matchERC1155Orders", sellOrder, buyOrder, sellOrderSignature, buyOrderSignature)
}

// MatchERC1155Orders is a paid mutator transaction binding the contract method 0x36a2c310.
//
// Solidity: function matchERC1155Orders((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,uint128) sellOrder, (address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[],uint128) buyOrder, (uint8,uint8,bytes32,bytes32) sellOrderSignature, (uint8,uint8,bytes32,bytes32) buyOrderSignature) returns(uint256 profit)
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureSession) MatchERC1155Orders(sellOrder LibNFTOrderERC1155SellOrder, buyOrder LibNFTOrderERC1155BuyOrder, sellOrderSignature LibSignatureSignature, buyOrderSignature LibSignatureSignature) (*types.Transaction, error) {
	return _ERC1155OrdersFeature.Contract.MatchERC1155Orders(&_ERC1155OrdersFeature.TransactOpts, sellOrder, buyOrder, sellOrderSignature, buyOrderSignature)
}

// MatchERC1155Orders is a paid mutator transaction binding the contract method 0x36a2c310.
//
// Solidity: function matchERC1155Orders((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,uint128) sellOrder, (address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[],uint128) buyOrder, (uint8,uint8,bytes32,bytes32) sellOrderSignature, (uint8,uint8,bytes32,bytes32) buyOrderSignature) returns(uint256 profit)
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureTransactorSession) MatchERC1155Orders(sellOrder LibNFTOrderERC1155SellOrder, buyOrder LibNFTOrderERC1155BuyOrder, sellOrderSignature LibSignatureSignature, buyOrderSignature LibSignatureSignature) (*types.Transaction, error) {
	return _ERC1155OrdersFeature.Contract.MatchERC1155Orders(&_ERC1155OrdersFeature.TransactOpts, sellOrder, buyOrder, sellOrderSignature, buyOrderSignature)
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

// PreSignERC1155BuyOrder is a paid mutator transaction binding the contract method 0x4a309223.
//
// Solidity: function preSignERC1155BuyOrder((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[],uint128) order) returns()
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureTransactor) PreSignERC1155BuyOrder(opts *bind.TransactOpts, order LibNFTOrderERC1155BuyOrder) (*types.Transaction, error) {
	return _ERC1155OrdersFeature.contract.Transact(opts, "preSignERC1155BuyOrder", order)
}

// PreSignERC1155BuyOrder is a paid mutator transaction binding the contract method 0x4a309223.
//
// Solidity: function preSignERC1155BuyOrder((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[],uint128) order) returns()
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureSession) PreSignERC1155BuyOrder(order LibNFTOrderERC1155BuyOrder) (*types.Transaction, error) {
	return _ERC1155OrdersFeature.Contract.PreSignERC1155BuyOrder(&_ERC1155OrdersFeature.TransactOpts, order)
}

// PreSignERC1155BuyOrder is a paid mutator transaction binding the contract method 0x4a309223.
//
// Solidity: function preSignERC1155BuyOrder((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[],uint128) order) returns()
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureTransactorSession) PreSignERC1155BuyOrder(order LibNFTOrderERC1155BuyOrder) (*types.Transaction, error) {
	return _ERC1155OrdersFeature.Contract.PreSignERC1155BuyOrder(&_ERC1155OrdersFeature.TransactOpts, order)
}

// PreSignERC1155SellOrder is a paid mutator transaction binding the contract method 0xd3206638.
//
// Solidity: function preSignERC1155SellOrder((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,uint128) order) returns()
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureTransactor) PreSignERC1155SellOrder(opts *bind.TransactOpts, order LibNFTOrderERC1155SellOrder) (*types.Transaction, error) {
	return _ERC1155OrdersFeature.contract.Transact(opts, "preSignERC1155SellOrder", order)
}

// PreSignERC1155SellOrder is a paid mutator transaction binding the contract method 0xd3206638.
//
// Solidity: function preSignERC1155SellOrder((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,uint128) order) returns()
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureSession) PreSignERC1155SellOrder(order LibNFTOrderERC1155SellOrder) (*types.Transaction, error) {
	return _ERC1155OrdersFeature.Contract.PreSignERC1155SellOrder(&_ERC1155OrdersFeature.TransactOpts, order)
}

// PreSignERC1155SellOrder is a paid mutator transaction binding the contract method 0xd3206638.
//
// Solidity: function preSignERC1155SellOrder((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,uint128) order) returns()
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureTransactorSession) PreSignERC1155SellOrder(order LibNFTOrderERC1155SellOrder) (*types.Transaction, error) {
	return _ERC1155OrdersFeature.Contract.PreSignERC1155SellOrder(&_ERC1155OrdersFeature.TransactOpts, order)
}

// SellERC1155 is a paid mutator transaction binding the contract method 0x496c5a55.
//
// Solidity: function sellERC1155((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[],uint128) buyOrder, (uint8,uint8,bytes32,bytes32) signature, uint256 erc1155TokenId, uint128 erc1155SellAmount, bool unwrapNativeToken, bytes takerData) returns()
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureTransactor) SellERC1155(opts *bind.TransactOpts, buyOrder LibNFTOrderERC1155BuyOrder, signature LibSignatureSignature, erc1155TokenId *big.Int, erc1155SellAmount *big.Int, unwrapNativeToken bool, takerData []byte) (*types.Transaction, error) {
	return _ERC1155OrdersFeature.contract.Transact(opts, "sellERC1155", buyOrder, signature, erc1155TokenId, erc1155SellAmount, unwrapNativeToken, takerData)
}

// SellERC1155 is a paid mutator transaction binding the contract method 0x496c5a55.
//
// Solidity: function sellERC1155((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[],uint128) buyOrder, (uint8,uint8,bytes32,bytes32) signature, uint256 erc1155TokenId, uint128 erc1155SellAmount, bool unwrapNativeToken, bytes takerData) returns()
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureSession) SellERC1155(buyOrder LibNFTOrderERC1155BuyOrder, signature LibSignatureSignature, erc1155TokenId *big.Int, erc1155SellAmount *big.Int, unwrapNativeToken bool, takerData []byte) (*types.Transaction, error) {
	return _ERC1155OrdersFeature.Contract.SellERC1155(&_ERC1155OrdersFeature.TransactOpts, buyOrder, signature, erc1155TokenId, erc1155SellAmount, unwrapNativeToken, takerData)
}

// SellERC1155 is a paid mutator transaction binding the contract method 0x496c5a55.
//
// Solidity: function sellERC1155((address,address,uint256,uint256,address,uint256,(address,uint256,bytes)[],address,uint256,(address,bytes)[],uint128) buyOrder, (uint8,uint8,bytes32,bytes32) signature, uint256 erc1155TokenId, uint128 erc1155SellAmount, bool unwrapNativeToken, bytes takerData) returns()
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureTransactorSession) SellERC1155(buyOrder LibNFTOrderERC1155BuyOrder, signature LibSignatureSignature, erc1155TokenId *big.Int, erc1155SellAmount *big.Int, unwrapNativeToken bool, takerData []byte) (*types.Transaction, error) {
	return _ERC1155OrdersFeature.Contract.SellERC1155(&_ERC1155OrdersFeature.TransactOpts, buyOrder, signature, erc1155TokenId, erc1155SellAmount, unwrapNativeToken, takerData)
}

// ERC1155OrdersFeatureERC1155BuyOrderFilledIterator is returned from FilterERC1155BuyOrderFilled and is used to iterate over the raw logs and unpacked data for ERC1155BuyOrderFilled events raised by the ERC1155OrdersFeature contract.
type ERC1155OrdersFeatureERC1155BuyOrderFilledIterator struct {
	Event *ERC1155OrdersFeatureERC1155BuyOrderFilled // Event containing the contract specifics and raw log

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
func (it *ERC1155OrdersFeatureERC1155BuyOrderFilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155OrdersFeatureERC1155BuyOrderFilled)
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
		it.Event = new(ERC1155OrdersFeatureERC1155BuyOrderFilled)
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
func (it *ERC1155OrdersFeatureERC1155BuyOrderFilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155OrdersFeatureERC1155BuyOrderFilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155OrdersFeatureERC1155BuyOrderFilled represents a ERC1155BuyOrderFilled event raised by the ERC1155OrdersFeature contract.
type ERC1155OrdersFeatureERC1155BuyOrderFilled struct {
	OrderHash         [32]byte
	Maker             common.Address
	Taker             common.Address
	Nonce             *big.Int
	Erc20Token        common.Address
	Erc20FillAmount   *big.Int
	Fees              []INFTOrdersFeatureFee
	Erc1155Token      common.Address
	Erc1155TokenId    *big.Int
	Erc1155FillAmount *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterERC1155BuyOrderFilled is a free log retrieval operation binding the contract event 0x105616901449a64554ca9246a5bbcaca973b40b3c0055e5070c6fa191618d9f3.
//
// Solidity: event ERC1155BuyOrderFilled(bytes32 orderHash, address maker, address taker, uint256 nonce, address erc20Token, uint256 erc20FillAmount, (address,uint256)[] fees, address erc1155Token, uint256 erc1155TokenId, uint128 erc1155FillAmount)
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureFilterer) FilterERC1155BuyOrderFilled(opts *bind.FilterOpts) (*ERC1155OrdersFeatureERC1155BuyOrderFilledIterator, error) {

	logs, sub, err := _ERC1155OrdersFeature.contract.FilterLogs(opts, "ERC1155BuyOrderFilled")
	if err != nil {
		return nil, err
	}
	return &ERC1155OrdersFeatureERC1155BuyOrderFilledIterator{contract: _ERC1155OrdersFeature.contract, event: "ERC1155BuyOrderFilled", logs: logs, sub: sub}, nil
}

// WatchERC1155BuyOrderFilled is a free log subscription operation binding the contract event 0x105616901449a64554ca9246a5bbcaca973b40b3c0055e5070c6fa191618d9f3.
//
// Solidity: event ERC1155BuyOrderFilled(bytes32 orderHash, address maker, address taker, uint256 nonce, address erc20Token, uint256 erc20FillAmount, (address,uint256)[] fees, address erc1155Token, uint256 erc1155TokenId, uint128 erc1155FillAmount)
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureFilterer) WatchERC1155BuyOrderFilled(opts *bind.WatchOpts, sink chan<- *ERC1155OrdersFeatureERC1155BuyOrderFilled) (event.Subscription, error) {

	logs, sub, err := _ERC1155OrdersFeature.contract.WatchLogs(opts, "ERC1155BuyOrderFilled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155OrdersFeatureERC1155BuyOrderFilled)
				if err := _ERC1155OrdersFeature.contract.UnpackLog(event, "ERC1155BuyOrderFilled", log); err != nil {
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

// ParseERC1155BuyOrderFilled is a log parse operation binding the contract event 0x105616901449a64554ca9246a5bbcaca973b40b3c0055e5070c6fa191618d9f3.
//
// Solidity: event ERC1155BuyOrderFilled(bytes32 orderHash, address maker, address taker, uint256 nonce, address erc20Token, uint256 erc20FillAmount, (address,uint256)[] fees, address erc1155Token, uint256 erc1155TokenId, uint128 erc1155FillAmount)
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureFilterer) ParseERC1155BuyOrderFilled(log types.Log) (*ERC1155OrdersFeatureERC1155BuyOrderFilled, error) {
	event := new(ERC1155OrdersFeatureERC1155BuyOrderFilled)
	if err := _ERC1155OrdersFeature.contract.UnpackLog(event, "ERC1155BuyOrderFilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155OrdersFeatureERC1155BuyOrderPreSignedIterator is returned from FilterERC1155BuyOrderPreSigned and is used to iterate over the raw logs and unpacked data for ERC1155BuyOrderPreSigned events raised by the ERC1155OrdersFeature contract.
type ERC1155OrdersFeatureERC1155BuyOrderPreSignedIterator struct {
	Event *ERC1155OrdersFeatureERC1155BuyOrderPreSigned // Event containing the contract specifics and raw log

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
func (it *ERC1155OrdersFeatureERC1155BuyOrderPreSignedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155OrdersFeatureERC1155BuyOrderPreSigned)
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
		it.Event = new(ERC1155OrdersFeatureERC1155BuyOrderPreSigned)
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
func (it *ERC1155OrdersFeatureERC1155BuyOrderPreSignedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155OrdersFeatureERC1155BuyOrderPreSignedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155OrdersFeatureERC1155BuyOrderPreSigned represents a ERC1155BuyOrderPreSigned event raised by the ERC1155OrdersFeature contract.
type ERC1155OrdersFeatureERC1155BuyOrderPreSigned struct {
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

// FilterERC1155BuyOrderPreSigned is a free log retrieval operation binding the contract event 0x46c048795fb749a72307b741e6796bb17ce9fa1c489ff16f4df4792aa7581850.
//
// Solidity: event ERC1155BuyOrderPreSigned(address maker, address taker, uint256 expiry, uint256 nonce, address erc20Token, uint256 erc20TokenAmount, (address,uint256,bytes)[] fees, address erc1155Token, uint256 erc1155TokenId, (address,bytes)[] erc1155TokenProperties, uint128 erc1155TokenAmount)
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureFilterer) FilterERC1155BuyOrderPreSigned(opts *bind.FilterOpts) (*ERC1155OrdersFeatureERC1155BuyOrderPreSignedIterator, error) {

	logs, sub, err := _ERC1155OrdersFeature.contract.FilterLogs(opts, "ERC1155BuyOrderPreSigned")
	if err != nil {
		return nil, err
	}
	return &ERC1155OrdersFeatureERC1155BuyOrderPreSignedIterator{contract: _ERC1155OrdersFeature.contract, event: "ERC1155BuyOrderPreSigned", logs: logs, sub: sub}, nil
}

// WatchERC1155BuyOrderPreSigned is a free log subscription operation binding the contract event 0x46c048795fb749a72307b741e6796bb17ce9fa1c489ff16f4df4792aa7581850.
//
// Solidity: event ERC1155BuyOrderPreSigned(address maker, address taker, uint256 expiry, uint256 nonce, address erc20Token, uint256 erc20TokenAmount, (address,uint256,bytes)[] fees, address erc1155Token, uint256 erc1155TokenId, (address,bytes)[] erc1155TokenProperties, uint128 erc1155TokenAmount)
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureFilterer) WatchERC1155BuyOrderPreSigned(opts *bind.WatchOpts, sink chan<- *ERC1155OrdersFeatureERC1155BuyOrderPreSigned) (event.Subscription, error) {

	logs, sub, err := _ERC1155OrdersFeature.contract.WatchLogs(opts, "ERC1155BuyOrderPreSigned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155OrdersFeatureERC1155BuyOrderPreSigned)
				if err := _ERC1155OrdersFeature.contract.UnpackLog(event, "ERC1155BuyOrderPreSigned", log); err != nil {
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

// ParseERC1155BuyOrderPreSigned is a log parse operation binding the contract event 0x46c048795fb749a72307b741e6796bb17ce9fa1c489ff16f4df4792aa7581850.
//
// Solidity: event ERC1155BuyOrderPreSigned(address maker, address taker, uint256 expiry, uint256 nonce, address erc20Token, uint256 erc20TokenAmount, (address,uint256,bytes)[] fees, address erc1155Token, uint256 erc1155TokenId, (address,bytes)[] erc1155TokenProperties, uint128 erc1155TokenAmount)
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureFilterer) ParseERC1155BuyOrderPreSigned(log types.Log) (*ERC1155OrdersFeatureERC1155BuyOrderPreSigned, error) {
	event := new(ERC1155OrdersFeatureERC1155BuyOrderPreSigned)
	if err := _ERC1155OrdersFeature.contract.UnpackLog(event, "ERC1155BuyOrderPreSigned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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

// ERC1155OrdersFeatureERC1155SellOrderFilledIterator is returned from FilterERC1155SellOrderFilled and is used to iterate over the raw logs and unpacked data for ERC1155SellOrderFilled events raised by the ERC1155OrdersFeature contract.
type ERC1155OrdersFeatureERC1155SellOrderFilledIterator struct {
	Event *ERC1155OrdersFeatureERC1155SellOrderFilled // Event containing the contract specifics and raw log

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
func (it *ERC1155OrdersFeatureERC1155SellOrderFilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155OrdersFeatureERC1155SellOrderFilled)
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
		it.Event = new(ERC1155OrdersFeatureERC1155SellOrderFilled)
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
func (it *ERC1155OrdersFeatureERC1155SellOrderFilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155OrdersFeatureERC1155SellOrderFilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155OrdersFeatureERC1155SellOrderFilled represents a ERC1155SellOrderFilled event raised by the ERC1155OrdersFeature contract.
type ERC1155OrdersFeatureERC1155SellOrderFilled struct {
	OrderHash         [32]byte
	Maker             common.Address
	Taker             common.Address
	Nonce             *big.Int
	Erc20Token        common.Address
	Erc20FillAmount   *big.Int
	Fees              []INFTOrdersFeatureFee
	Erc1155Token      common.Address
	Erc1155TokenId    *big.Int
	Erc1155FillAmount *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterERC1155SellOrderFilled is a free log retrieval operation binding the contract event 0xfcde121a3f6a9b14a3ce266d61fc00940de86c4d8c1d733fe62d503ae5d99ff9.
//
// Solidity: event ERC1155SellOrderFilled(bytes32 orderHash, address maker, address taker, uint256 nonce, address erc20Token, uint256 erc20FillAmount, (address,uint256)[] fees, address erc1155Token, uint256 erc1155TokenId, uint128 erc1155FillAmount)
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureFilterer) FilterERC1155SellOrderFilled(opts *bind.FilterOpts) (*ERC1155OrdersFeatureERC1155SellOrderFilledIterator, error) {

	logs, sub, err := _ERC1155OrdersFeature.contract.FilterLogs(opts, "ERC1155SellOrderFilled")
	if err != nil {
		return nil, err
	}
	return &ERC1155OrdersFeatureERC1155SellOrderFilledIterator{contract: _ERC1155OrdersFeature.contract, event: "ERC1155SellOrderFilled", logs: logs, sub: sub}, nil
}

// WatchERC1155SellOrderFilled is a free log subscription operation binding the contract event 0xfcde121a3f6a9b14a3ce266d61fc00940de86c4d8c1d733fe62d503ae5d99ff9.
//
// Solidity: event ERC1155SellOrderFilled(bytes32 orderHash, address maker, address taker, uint256 nonce, address erc20Token, uint256 erc20FillAmount, (address,uint256)[] fees, address erc1155Token, uint256 erc1155TokenId, uint128 erc1155FillAmount)
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureFilterer) WatchERC1155SellOrderFilled(opts *bind.WatchOpts, sink chan<- *ERC1155OrdersFeatureERC1155SellOrderFilled) (event.Subscription, error) {

	logs, sub, err := _ERC1155OrdersFeature.contract.WatchLogs(opts, "ERC1155SellOrderFilled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155OrdersFeatureERC1155SellOrderFilled)
				if err := _ERC1155OrdersFeature.contract.UnpackLog(event, "ERC1155SellOrderFilled", log); err != nil {
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

// ParseERC1155SellOrderFilled is a log parse operation binding the contract event 0xfcde121a3f6a9b14a3ce266d61fc00940de86c4d8c1d733fe62d503ae5d99ff9.
//
// Solidity: event ERC1155SellOrderFilled(bytes32 orderHash, address maker, address taker, uint256 nonce, address erc20Token, uint256 erc20FillAmount, (address,uint256)[] fees, address erc1155Token, uint256 erc1155TokenId, uint128 erc1155FillAmount)
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureFilterer) ParseERC1155SellOrderFilled(log types.Log) (*ERC1155OrdersFeatureERC1155SellOrderFilled, error) {
	event := new(ERC1155OrdersFeatureERC1155SellOrderFilled)
	if err := _ERC1155OrdersFeature.contract.UnpackLog(event, "ERC1155SellOrderFilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155OrdersFeatureERC1155SellOrderPreSignedIterator is returned from FilterERC1155SellOrderPreSigned and is used to iterate over the raw logs and unpacked data for ERC1155SellOrderPreSigned events raised by the ERC1155OrdersFeature contract.
type ERC1155OrdersFeatureERC1155SellOrderPreSignedIterator struct {
	Event *ERC1155OrdersFeatureERC1155SellOrderPreSigned // Event containing the contract specifics and raw log

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
func (it *ERC1155OrdersFeatureERC1155SellOrderPreSignedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155OrdersFeatureERC1155SellOrderPreSigned)
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
		it.Event = new(ERC1155OrdersFeatureERC1155SellOrderPreSigned)
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
func (it *ERC1155OrdersFeatureERC1155SellOrderPreSignedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155OrdersFeatureERC1155SellOrderPreSignedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155OrdersFeatureERC1155SellOrderPreSigned represents a ERC1155SellOrderPreSigned event raised by the ERC1155OrdersFeature contract.
type ERC1155OrdersFeatureERC1155SellOrderPreSigned struct {
	Maker              common.Address
	Taker              common.Address
	Expiry             *big.Int
	Nonce              *big.Int
	Erc20Token         common.Address
	Erc20TokenAmount   *big.Int
	Fees               []LibNFTOrderFee
	Erc1155Token       common.Address
	Erc1155TokenId     *big.Int
	Erc1155TokenAmount *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterERC1155SellOrderPreSigned is a free log retrieval operation binding the contract event 0xebbc2d593c5767af9512abf631b124dbec053e62355184ce6a40e4acff1fc3b0.
//
// Solidity: event ERC1155SellOrderPreSigned(address maker, address taker, uint256 expiry, uint256 nonce, address erc20Token, uint256 erc20TokenAmount, (address,uint256,bytes)[] fees, address erc1155Token, uint256 erc1155TokenId, uint128 erc1155TokenAmount)
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureFilterer) FilterERC1155SellOrderPreSigned(opts *bind.FilterOpts) (*ERC1155OrdersFeatureERC1155SellOrderPreSignedIterator, error) {

	logs, sub, err := _ERC1155OrdersFeature.contract.FilterLogs(opts, "ERC1155SellOrderPreSigned")
	if err != nil {
		return nil, err
	}
	return &ERC1155OrdersFeatureERC1155SellOrderPreSignedIterator{contract: _ERC1155OrdersFeature.contract, event: "ERC1155SellOrderPreSigned", logs: logs, sub: sub}, nil
}

// WatchERC1155SellOrderPreSigned is a free log subscription operation binding the contract event 0xebbc2d593c5767af9512abf631b124dbec053e62355184ce6a40e4acff1fc3b0.
//
// Solidity: event ERC1155SellOrderPreSigned(address maker, address taker, uint256 expiry, uint256 nonce, address erc20Token, uint256 erc20TokenAmount, (address,uint256,bytes)[] fees, address erc1155Token, uint256 erc1155TokenId, uint128 erc1155TokenAmount)
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureFilterer) WatchERC1155SellOrderPreSigned(opts *bind.WatchOpts, sink chan<- *ERC1155OrdersFeatureERC1155SellOrderPreSigned) (event.Subscription, error) {

	logs, sub, err := _ERC1155OrdersFeature.contract.WatchLogs(opts, "ERC1155SellOrderPreSigned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155OrdersFeatureERC1155SellOrderPreSigned)
				if err := _ERC1155OrdersFeature.contract.UnpackLog(event, "ERC1155SellOrderPreSigned", log); err != nil {
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

// ParseERC1155SellOrderPreSigned is a log parse operation binding the contract event 0xebbc2d593c5767af9512abf631b124dbec053e62355184ce6a40e4acff1fc3b0.
//
// Solidity: event ERC1155SellOrderPreSigned(address maker, address taker, uint256 expiry, uint256 nonce, address erc20Token, uint256 erc20TokenAmount, (address,uint256,bytes)[] fees, address erc1155Token, uint256 erc1155TokenId, uint128 erc1155TokenAmount)
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureFilterer) ParseERC1155SellOrderPreSigned(log types.Log) (*ERC1155OrdersFeatureERC1155SellOrderPreSigned, error) {
	event := new(ERC1155OrdersFeatureERC1155SellOrderPreSigned)
	if err := _ERC1155OrdersFeature.contract.UnpackLog(event, "ERC1155SellOrderPreSigned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155OrdersFeatureTakerDataEmittedIterator is returned from FilterTakerDataEmitted and is used to iterate over the raw logs and unpacked data for TakerDataEmitted events raised by the ERC1155OrdersFeature contract.
type ERC1155OrdersFeatureTakerDataEmittedIterator struct {
	Event *ERC1155OrdersFeatureTakerDataEmitted // Event containing the contract specifics and raw log

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
func (it *ERC1155OrdersFeatureTakerDataEmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155OrdersFeatureTakerDataEmitted)
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
		it.Event = new(ERC1155OrdersFeatureTakerDataEmitted)
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
func (it *ERC1155OrdersFeatureTakerDataEmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155OrdersFeatureTakerDataEmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155OrdersFeatureTakerDataEmitted represents a TakerDataEmitted event raised by the ERC1155OrdersFeature contract.
type ERC1155OrdersFeatureTakerDataEmitted struct {
	OrderHash [32]byte
	TakerData []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTakerDataEmitted is a free log retrieval operation binding the contract event 0xf61c2baa48a3c53b619a8e6c2bb6d677f82466c61970ee6afd4a157b0fbf8756.
//
// Solidity: event TakerDataEmitted(bytes32 orderHash, bytes takerData)
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureFilterer) FilterTakerDataEmitted(opts *bind.FilterOpts) (*ERC1155OrdersFeatureTakerDataEmittedIterator, error) {

	logs, sub, err := _ERC1155OrdersFeature.contract.FilterLogs(opts, "TakerDataEmitted")
	if err != nil {
		return nil, err
	}
	return &ERC1155OrdersFeatureTakerDataEmittedIterator{contract: _ERC1155OrdersFeature.contract, event: "TakerDataEmitted", logs: logs, sub: sub}, nil
}

// WatchTakerDataEmitted is a free log subscription operation binding the contract event 0xf61c2baa48a3c53b619a8e6c2bb6d677f82466c61970ee6afd4a157b0fbf8756.
//
// Solidity: event TakerDataEmitted(bytes32 orderHash, bytes takerData)
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureFilterer) WatchTakerDataEmitted(opts *bind.WatchOpts, sink chan<- *ERC1155OrdersFeatureTakerDataEmitted) (event.Subscription, error) {

	logs, sub, err := _ERC1155OrdersFeature.contract.WatchLogs(opts, "TakerDataEmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155OrdersFeatureTakerDataEmitted)
				if err := _ERC1155OrdersFeature.contract.UnpackLog(event, "TakerDataEmitted", log); err != nil {
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
func (_ERC1155OrdersFeature *ERC1155OrdersFeatureFilterer) ParseTakerDataEmitted(log types.Log) (*ERC1155OrdersFeatureTakerDataEmitted, error) {
	event := new(ERC1155OrdersFeatureTakerDataEmitted)
	if err := _ERC1155OrdersFeature.contract.UnpackLog(event, "TakerDataEmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
