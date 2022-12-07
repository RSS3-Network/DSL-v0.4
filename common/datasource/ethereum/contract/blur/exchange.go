// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package blur

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

// Execution is an auto generated low-level Go binding around an user-defined struct.
type Execution struct {
	Sell Input
	Buy  Input
}

// Fee is an auto generated low-level Go binding around an user-defined struct.
type Fee struct {
	Rate      uint16
	Recipient common.Address
}

// Input is an auto generated low-level Go binding around an user-defined struct.
type Input struct {
	Order            Order
	V                uint8
	R                [32]byte
	S                [32]byte
	ExtraSignature   []byte
	SignatureVersion uint8
	BlockNumber      *big.Int
}

// Order is an auto generated low-level Go binding around an user-defined struct.
type Order struct {
	Trader         common.Address
	Side           uint8
	MatchingPolicy common.Address
	Collection     common.Address
	TokenId        *big.Int
	Amount         *big.Int
	PaymentToken   common.Address
	Price          *big.Int
	ListingTime    *big.Int
	ExpirationTime *big.Int
	Fees           []Fee
	Salt           *big.Int
	ExtraParams    []byte
}

// ExchangeMetaData contains all meta data concerning the Exchange contract.
var ExchangeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Closed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockRange\",\"type\":\"uint256\"}],\"name\":\"NewBlockRange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIExecutionDelegate\",\"name\":\"executionDelegate\",\"type\":\"address\"}],\"name\":\"NewExecutionDelegate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"}],\"name\":\"NewOracle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIPolicyManager\",\"name\":\"policyManager\",\"type\":\"address\"}],\"name\":\"NewPolicyManager\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newNonce\",\"type\":\"uint256\"}],\"name\":\"NonceIncremented\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Opened\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"OrderCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"matchingPolicy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structFee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraParams\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structOrder\",\"name\":\"sell\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"sellHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"matchingPolicy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structFee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraParams\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structOrder\",\"name\":\"buy\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"buyHash\",\"type\":\"bytes32\"}],\"name\":\"OrdersMatched\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"FEE_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INVERSE_BASIS_POINT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NAME\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ORACLE_ORDER_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ORDER_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"POOL\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ROOT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"matchingPolicy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structFee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraParams\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"extraSignature\",\"type\":\"bytes\"},{\"internalType\":\"enumSignatureVersion\",\"name\":\"signatureVersion\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"internalType\":\"structInput\",\"name\":\"sell\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"matchingPolicy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structFee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraParams\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"extraSignature\",\"type\":\"bytes\"},{\"internalType\":\"enumSignatureVersion\",\"name\":\"signatureVersion\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"internalType\":\"structInput\",\"name\":\"buy\",\"type\":\"tuple\"}],\"name\":\"_execute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blockRange\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"matchingPolicy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structFee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraParams\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"extraSignature\",\"type\":\"bytes\"},{\"internalType\":\"enumSignatureVersion\",\"name\":\"signatureVersion\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"internalType\":\"structInput\",\"name\":\"sell\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"matchingPolicy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structFee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraParams\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"extraSignature\",\"type\":\"bytes\"},{\"internalType\":\"enumSignatureVersion\",\"name\":\"signatureVersion\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"internalType\":\"structInput\",\"name\":\"buy\",\"type\":\"tuple\"}],\"internalType\":\"structExecution[]\",\"name\":\"executions\",\"type\":\"tuple[]\"}],\"name\":\"bulkExecute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"matchingPolicy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structFee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraParams\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"cancelOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"matchingPolicy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structFee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraParams\",\"type\":\"bytes\"}],\"internalType\":\"structOrder[]\",\"name\":\"orders\",\"type\":\"tuple[]\"}],\"name\":\"cancelOrders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"cancelledOrFilled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"close\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"matchingPolicy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structFee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraParams\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"extraSignature\",\"type\":\"bytes\"},{\"internalType\":\"enumSignatureVersion\",\"name\":\"signatureVersion\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"internalType\":\"structInput\",\"name\":\"sell\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"matchingPolicy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structFee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraParams\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"extraSignature\",\"type\":\"bytes\"},{\"internalType\":\"enumSignatureVersion\",\"name\":\"signatureVersion\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"internalType\":\"structInput\",\"name\":\"buy\",\"type\":\"tuple\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"executionDelegate\",\"outputs\":[{\"internalType\":\"contractIExecutionDelegate\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"incrementNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIExecutionDelegate\",\"name\":\"_executionDelegate\",\"type\":\"address\"},{\"internalType\":\"contractIPolicyManager\",\"name\":\"_policyManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_blockRange\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isInternal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOpen\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"open\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"policyManager\",\"outputs\":[{\"internalType\":\"contractIPolicyManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"remainingETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockRange\",\"type\":\"uint256\"}],\"name\":\"setBlockRange\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIExecutionDelegate\",\"name\":\"_executionDelegate\",\"type\":\"address\"}],\"name\":\"setExecutionDelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"}],\"name\":\"setOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPolicyManager\",\"name\":\"_policyManager\",\"type\":\"address\"}],\"name\":\"setPolicyManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// ExchangeABI is the input ABI used to generate the binding from.
// Deprecated: Use ExchangeMetaData.ABI instead.
var ExchangeABI = ExchangeMetaData.ABI

// Exchange is an auto generated Go binding around an Ethereum contract.
type Exchange struct {
	ExchangeCaller     // Read-only binding to the contract
	ExchangeTransactor // Write-only binding to the contract
	ExchangeFilterer   // Log filterer for contract events
}

// ExchangeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExchangeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExchangeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExchangeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExchangeSession struct {
	Contract     *Exchange         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExchangeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExchangeCallerSession struct {
	Contract *ExchangeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ExchangeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExchangeTransactorSession struct {
	Contract     *ExchangeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ExchangeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExchangeRaw struct {
	Contract *Exchange // Generic contract binding to access the raw methods on
}

// ExchangeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExchangeCallerRaw struct {
	Contract *ExchangeCaller // Generic read-only contract binding to access the raw methods on
}

// ExchangeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExchangeTransactorRaw struct {
	Contract *ExchangeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExchange creates a new instance of Exchange, bound to a specific deployed contract.
func NewExchange(address common.Address, backend bind.ContractBackend) (*Exchange, error) {
	contract, err := bindExchange(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Exchange{ExchangeCaller: ExchangeCaller{contract: contract}, ExchangeTransactor: ExchangeTransactor{contract: contract}, ExchangeFilterer: ExchangeFilterer{contract: contract}}, nil
}

// NewExchangeCaller creates a new read-only instance of Exchange, bound to a specific deployed contract.
func NewExchangeCaller(address common.Address, caller bind.ContractCaller) (*ExchangeCaller, error) {
	contract, err := bindExchange(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExchangeCaller{contract: contract}, nil
}

// NewExchangeTransactor creates a new write-only instance of Exchange, bound to a specific deployed contract.
func NewExchangeTransactor(address common.Address, transactor bind.ContractTransactor) (*ExchangeTransactor, error) {
	contract, err := bindExchange(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExchangeTransactor{contract: contract}, nil
}

// NewExchangeFilterer creates a new log filterer instance of Exchange, bound to a specific deployed contract.
func NewExchangeFilterer(address common.Address, filterer bind.ContractFilterer) (*ExchangeFilterer, error) {
	contract, err := bindExchange(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExchangeFilterer{contract: contract}, nil
}

// bindExchange binds a generic wrapper to an already deployed contract.
func bindExchange(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ExchangeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Exchange *ExchangeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Exchange.Contract.ExchangeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Exchange *ExchangeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Exchange.Contract.ExchangeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Exchange *ExchangeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Exchange.Contract.ExchangeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Exchange *ExchangeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Exchange.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Exchange *ExchangeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Exchange.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Exchange *ExchangeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Exchange.Contract.contract.Transact(opts, method, params...)
}

// FEETYPEHASH is a free data retrieval call binding the contract method 0x4832ede1.
//
// Solidity: function FEE_TYPEHASH() view returns(bytes32)
func (_Exchange *ExchangeCaller) FEETYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "FEE_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// FEETYPEHASH is a free data retrieval call binding the contract method 0x4832ede1.
//
// Solidity: function FEE_TYPEHASH() view returns(bytes32)
func (_Exchange *ExchangeSession) FEETYPEHASH() ([32]byte, error) {
	return _Exchange.Contract.FEETYPEHASH(&_Exchange.CallOpts)
}

// FEETYPEHASH is a free data retrieval call binding the contract method 0x4832ede1.
//
// Solidity: function FEE_TYPEHASH() view returns(bytes32)
func (_Exchange *ExchangeCallerSession) FEETYPEHASH() ([32]byte, error) {
	return _Exchange.Contract.FEETYPEHASH(&_Exchange.CallOpts)
}

// INVERSEBASISPOINT is a free data retrieval call binding the contract method 0xcae6047f.
//
// Solidity: function INVERSE_BASIS_POINT() view returns(uint256)
func (_Exchange *ExchangeCaller) INVERSEBASISPOINT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "INVERSE_BASIS_POINT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// INVERSEBASISPOINT is a free data retrieval call binding the contract method 0xcae6047f.
//
// Solidity: function INVERSE_BASIS_POINT() view returns(uint256)
func (_Exchange *ExchangeSession) INVERSEBASISPOINT() (*big.Int, error) {
	return _Exchange.Contract.INVERSEBASISPOINT(&_Exchange.CallOpts)
}

// INVERSEBASISPOINT is a free data retrieval call binding the contract method 0xcae6047f.
//
// Solidity: function INVERSE_BASIS_POINT() view returns(uint256)
func (_Exchange *ExchangeCallerSession) INVERSEBASISPOINT() (*big.Int, error) {
	return _Exchange.Contract.INVERSEBASISPOINT(&_Exchange.CallOpts)
}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() view returns(string)
func (_Exchange *ExchangeCaller) NAME(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "NAME")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() view returns(string)
func (_Exchange *ExchangeSession) NAME() (string, error) {
	return _Exchange.Contract.NAME(&_Exchange.CallOpts)
}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() view returns(string)
func (_Exchange *ExchangeCallerSession) NAME() (string, error) {
	return _Exchange.Contract.NAME(&_Exchange.CallOpts)
}

// ORACLEORDERTYPEHASH is a free data retrieval call binding the contract method 0x1d97c9bb.
//
// Solidity: function ORACLE_ORDER_TYPEHASH() view returns(bytes32)
func (_Exchange *ExchangeCaller) ORACLEORDERTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "ORACLE_ORDER_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ORACLEORDERTYPEHASH is a free data retrieval call binding the contract method 0x1d97c9bb.
//
// Solidity: function ORACLE_ORDER_TYPEHASH() view returns(bytes32)
func (_Exchange *ExchangeSession) ORACLEORDERTYPEHASH() ([32]byte, error) {
	return _Exchange.Contract.ORACLEORDERTYPEHASH(&_Exchange.CallOpts)
}

// ORACLEORDERTYPEHASH is a free data retrieval call binding the contract method 0x1d97c9bb.
//
// Solidity: function ORACLE_ORDER_TYPEHASH() view returns(bytes32)
func (_Exchange *ExchangeCallerSession) ORACLEORDERTYPEHASH() ([32]byte, error) {
	return _Exchange.Contract.ORACLEORDERTYPEHASH(&_Exchange.CallOpts)
}

// ORDERTYPEHASH is a free data retrieval call binding the contract method 0xf973a209.
//
// Solidity: function ORDER_TYPEHASH() view returns(bytes32)
func (_Exchange *ExchangeCaller) ORDERTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "ORDER_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ORDERTYPEHASH is a free data retrieval call binding the contract method 0xf973a209.
//
// Solidity: function ORDER_TYPEHASH() view returns(bytes32)
func (_Exchange *ExchangeSession) ORDERTYPEHASH() ([32]byte, error) {
	return _Exchange.Contract.ORDERTYPEHASH(&_Exchange.CallOpts)
}

// ORDERTYPEHASH is a free data retrieval call binding the contract method 0xf973a209.
//
// Solidity: function ORDER_TYPEHASH() view returns(bytes32)
func (_Exchange *ExchangeCallerSession) ORDERTYPEHASH() ([32]byte, error) {
	return _Exchange.Contract.ORDERTYPEHASH(&_Exchange.CallOpts)
}

// POOL is a free data retrieval call binding the contract method 0x7535d246.
//
// Solidity: function POOL() view returns(address)
func (_Exchange *ExchangeCaller) POOL(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "POOL")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// POOL is a free data retrieval call binding the contract method 0x7535d246.
//
// Solidity: function POOL() view returns(address)
func (_Exchange *ExchangeSession) POOL() (common.Address, error) {
	return _Exchange.Contract.POOL(&_Exchange.CallOpts)
}

// POOL is a free data retrieval call binding the contract method 0x7535d246.
//
// Solidity: function POOL() view returns(address)
func (_Exchange *ExchangeCallerSession) POOL() (common.Address, error) {
	return _Exchange.Contract.POOL(&_Exchange.CallOpts)
}

// ROOTTYPEHASH is a free data retrieval call binding the contract method 0x31e6d0fe.
//
// Solidity: function ROOT_TYPEHASH() view returns(bytes32)
func (_Exchange *ExchangeCaller) ROOTTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "ROOT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ROOTTYPEHASH is a free data retrieval call binding the contract method 0x31e6d0fe.
//
// Solidity: function ROOT_TYPEHASH() view returns(bytes32)
func (_Exchange *ExchangeSession) ROOTTYPEHASH() ([32]byte, error) {
	return _Exchange.Contract.ROOTTYPEHASH(&_Exchange.CallOpts)
}

// ROOTTYPEHASH is a free data retrieval call binding the contract method 0x31e6d0fe.
//
// Solidity: function ROOT_TYPEHASH() view returns(bytes32)
func (_Exchange *ExchangeCallerSession) ROOTTYPEHASH() ([32]byte, error) {
	return _Exchange.Contract.ROOTTYPEHASH(&_Exchange.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Exchange *ExchangeCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Exchange *ExchangeSession) VERSION() (string, error) {
	return _Exchange.Contract.VERSION(&_Exchange.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Exchange *ExchangeCallerSession) VERSION() (string, error) {
	return _Exchange.Contract.VERSION(&_Exchange.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Exchange *ExchangeCaller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Exchange *ExchangeSession) WETH() (common.Address, error) {
	return _Exchange.Contract.WETH(&_Exchange.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Exchange *ExchangeCallerSession) WETH() (common.Address, error) {
	return _Exchange.Contract.WETH(&_Exchange.CallOpts)
}

// BlockRange is a free data retrieval call binding the contract method 0xa4b2c674.
//
// Solidity: function blockRange() view returns(uint256)
func (_Exchange *ExchangeCaller) BlockRange(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "blockRange")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BlockRange is a free data retrieval call binding the contract method 0xa4b2c674.
//
// Solidity: function blockRange() view returns(uint256)
func (_Exchange *ExchangeSession) BlockRange() (*big.Int, error) {
	return _Exchange.Contract.BlockRange(&_Exchange.CallOpts)
}

// BlockRange is a free data retrieval call binding the contract method 0xa4b2c674.
//
// Solidity: function blockRange() view returns(uint256)
func (_Exchange *ExchangeCallerSession) BlockRange() (*big.Int, error) {
	return _Exchange.Contract.BlockRange(&_Exchange.CallOpts)
}

// CancelledOrFilled is a free data retrieval call binding the contract method 0x5511f319.
//
// Solidity: function cancelledOrFilled(bytes32 ) view returns(bool)
func (_Exchange *ExchangeCaller) CancelledOrFilled(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "cancelledOrFilled", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CancelledOrFilled is a free data retrieval call binding the contract method 0x5511f319.
//
// Solidity: function cancelledOrFilled(bytes32 ) view returns(bool)
func (_Exchange *ExchangeSession) CancelledOrFilled(arg0 [32]byte) (bool, error) {
	return _Exchange.Contract.CancelledOrFilled(&_Exchange.CallOpts, arg0)
}

// CancelledOrFilled is a free data retrieval call binding the contract method 0x5511f319.
//
// Solidity: function cancelledOrFilled(bytes32 ) view returns(bool)
func (_Exchange *ExchangeCallerSession) CancelledOrFilled(arg0 [32]byte) (bool, error) {
	return _Exchange.Contract.CancelledOrFilled(&_Exchange.CallOpts, arg0)
}

// ExecutionDelegate is a free data retrieval call binding the contract method 0x986c9b20.
//
// Solidity: function executionDelegate() view returns(address)
func (_Exchange *ExchangeCaller) ExecutionDelegate(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "executionDelegate")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ExecutionDelegate is a free data retrieval call binding the contract method 0x986c9b20.
//
// Solidity: function executionDelegate() view returns(address)
func (_Exchange *ExchangeSession) ExecutionDelegate() (common.Address, error) {
	return _Exchange.Contract.ExecutionDelegate(&_Exchange.CallOpts)
}

// ExecutionDelegate is a free data retrieval call binding the contract method 0x986c9b20.
//
// Solidity: function executionDelegate() view returns(address)
func (_Exchange *ExchangeCallerSession) ExecutionDelegate() (common.Address, error) {
	return _Exchange.Contract.ExecutionDelegate(&_Exchange.CallOpts)
}

// IsInternal is a free data retrieval call binding the contract method 0x16e29d71.
//
// Solidity: function isInternal() view returns(bool)
func (_Exchange *ExchangeCaller) IsInternal(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "isInternal")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInternal is a free data retrieval call binding the contract method 0x16e29d71.
//
// Solidity: function isInternal() view returns(bool)
func (_Exchange *ExchangeSession) IsInternal() (bool, error) {
	return _Exchange.Contract.IsInternal(&_Exchange.CallOpts)
}

// IsInternal is a free data retrieval call binding the contract method 0x16e29d71.
//
// Solidity: function isInternal() view returns(bool)
func (_Exchange *ExchangeCallerSession) IsInternal() (bool, error) {
	return _Exchange.Contract.IsInternal(&_Exchange.CallOpts)
}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(uint256)
func (_Exchange *ExchangeCaller) IsOpen(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "isOpen")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(uint256)
func (_Exchange *ExchangeSession) IsOpen() (*big.Int, error) {
	return _Exchange.Contract.IsOpen(&_Exchange.CallOpts)
}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(uint256)
func (_Exchange *ExchangeCallerSession) IsOpen() (*big.Int, error) {
	return _Exchange.Contract.IsOpen(&_Exchange.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Exchange *ExchangeCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Exchange *ExchangeSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Exchange.Contract.Nonces(&_Exchange.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Exchange *ExchangeCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Exchange.Contract.Nonces(&_Exchange.CallOpts, arg0)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Exchange *ExchangeCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "oracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Exchange *ExchangeSession) Oracle() (common.Address, error) {
	return _Exchange.Contract.Oracle(&_Exchange.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Exchange *ExchangeCallerSession) Oracle() (common.Address, error) {
	return _Exchange.Contract.Oracle(&_Exchange.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Exchange *ExchangeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Exchange *ExchangeSession) Owner() (common.Address, error) {
	return _Exchange.Contract.Owner(&_Exchange.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Exchange *ExchangeCallerSession) Owner() (common.Address, error) {
	return _Exchange.Contract.Owner(&_Exchange.CallOpts)
}

// PolicyManager is a free data retrieval call binding the contract method 0xab3dbf3b.
//
// Solidity: function policyManager() view returns(address)
func (_Exchange *ExchangeCaller) PolicyManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "policyManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PolicyManager is a free data retrieval call binding the contract method 0xab3dbf3b.
//
// Solidity: function policyManager() view returns(address)
func (_Exchange *ExchangeSession) PolicyManager() (common.Address, error) {
	return _Exchange.Contract.PolicyManager(&_Exchange.CallOpts)
}

// PolicyManager is a free data retrieval call binding the contract method 0xab3dbf3b.
//
// Solidity: function policyManager() view returns(address)
func (_Exchange *ExchangeCallerSession) PolicyManager() (common.Address, error) {
	return _Exchange.Contract.PolicyManager(&_Exchange.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Exchange *ExchangeCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Exchange *ExchangeSession) ProxiableUUID() ([32]byte, error) {
	return _Exchange.Contract.ProxiableUUID(&_Exchange.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Exchange *ExchangeCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Exchange.Contract.ProxiableUUID(&_Exchange.CallOpts)
}

// RemainingETH is a free data retrieval call binding the contract method 0x2c7acf8c.
//
// Solidity: function remainingETH() view returns(uint256)
func (_Exchange *ExchangeCaller) RemainingETH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "remainingETH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RemainingETH is a free data retrieval call binding the contract method 0x2c7acf8c.
//
// Solidity: function remainingETH() view returns(uint256)
func (_Exchange *ExchangeSession) RemainingETH() (*big.Int, error) {
	return _Exchange.Contract.RemainingETH(&_Exchange.CallOpts)
}

// RemainingETH is a free data retrieval call binding the contract method 0x2c7acf8c.
//
// Solidity: function remainingETH() view returns(uint256)
func (_Exchange *ExchangeCallerSession) RemainingETH() (*big.Int, error) {
	return _Exchange.Contract.RemainingETH(&_Exchange.CallOpts)
}

// Execute is a paid mutator transaction binding the contract method 0xe04d94ae.
//
// Solidity: function _execute(((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256) sell, ((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256) buy) payable returns()
func (_Exchange *ExchangeTransactor) Execute(opts *bind.TransactOpts, sell Input, buy Input) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "_execute", sell, buy)
}

// Execute is a paid mutator transaction binding the contract method 0xe04d94ae.
//
// Solidity: function _execute(((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256) sell, ((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256) buy) payable returns()
func (_Exchange *ExchangeSession) Execute(sell Input, buy Input) (*types.Transaction, error) {
	return _Exchange.Contract.Execute(&_Exchange.TransactOpts, sell, buy)
}

// Execute is a paid mutator transaction binding the contract method 0xe04d94ae.
//
// Solidity: function _execute(((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256) sell, ((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256) buy) payable returns()
func (_Exchange *ExchangeTransactorSession) Execute(sell Input, buy Input) (*types.Transaction, error) {
	return _Exchange.Contract.Execute(&_Exchange.TransactOpts, sell, buy)
}

// BulkExecute is a paid mutator transaction binding the contract method 0xb3be57f8.
//
// Solidity: function bulkExecute((((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256),((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256))[] executions) payable returns()
func (_Exchange *ExchangeTransactor) BulkExecute(opts *bind.TransactOpts, executions []Execution) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "bulkExecute", executions)
}

// BulkExecute is a paid mutator transaction binding the contract method 0xb3be57f8.
//
// Solidity: function bulkExecute((((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256),((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256))[] executions) payable returns()
func (_Exchange *ExchangeSession) BulkExecute(executions []Execution) (*types.Transaction, error) {
	return _Exchange.Contract.BulkExecute(&_Exchange.TransactOpts, executions)
}

// BulkExecute is a paid mutator transaction binding the contract method 0xb3be57f8.
//
// Solidity: function bulkExecute((((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256),((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256))[] executions) payable returns()
func (_Exchange *ExchangeTransactorSession) BulkExecute(executions []Execution) (*types.Transaction, error) {
	return _Exchange.Contract.BulkExecute(&_Exchange.TransactOpts, executions)
}

// CancelOrder is a paid mutator transaction binding the contract method 0xf4acd740.
//
// Solidity: function cancelOrder((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes) order) returns()
func (_Exchange *ExchangeTransactor) CancelOrder(opts *bind.TransactOpts, order Order) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "cancelOrder", order)
}

// CancelOrder is a paid mutator transaction binding the contract method 0xf4acd740.
//
// Solidity: function cancelOrder((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes) order) returns()
func (_Exchange *ExchangeSession) CancelOrder(order Order) (*types.Transaction, error) {
	return _Exchange.Contract.CancelOrder(&_Exchange.TransactOpts, order)
}

// CancelOrder is a paid mutator transaction binding the contract method 0xf4acd740.
//
// Solidity: function cancelOrder((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes) order) returns()
func (_Exchange *ExchangeTransactorSession) CancelOrder(order Order) (*types.Transaction, error) {
	return _Exchange.Contract.CancelOrder(&_Exchange.TransactOpts, order)
}

// CancelOrders is a paid mutator transaction binding the contract method 0xab7e8cba.
//
// Solidity: function cancelOrders((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes)[] orders) returns()
func (_Exchange *ExchangeTransactor) CancelOrders(opts *bind.TransactOpts, orders []Order) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "cancelOrders", orders)
}

// CancelOrders is a paid mutator transaction binding the contract method 0xab7e8cba.
//
// Solidity: function cancelOrders((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes)[] orders) returns()
func (_Exchange *ExchangeSession) CancelOrders(orders []Order) (*types.Transaction, error) {
	return _Exchange.Contract.CancelOrders(&_Exchange.TransactOpts, orders)
}

// CancelOrders is a paid mutator transaction binding the contract method 0xab7e8cba.
//
// Solidity: function cancelOrders((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes)[] orders) returns()
func (_Exchange *ExchangeTransactorSession) CancelOrders(orders []Order) (*types.Transaction, error) {
	return _Exchange.Contract.CancelOrders(&_Exchange.TransactOpts, orders)
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_Exchange *ExchangeTransactor) Close(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "close")
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_Exchange *ExchangeSession) Close() (*types.Transaction, error) {
	return _Exchange.Contract.Close(&_Exchange.TransactOpts)
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_Exchange *ExchangeTransactorSession) Close() (*types.Transaction, error) {
	return _Exchange.Contract.Close(&_Exchange.TransactOpts)
}

// Execute2 is a paid mutator transaction binding the contract method 0x9a1fc3a7.
//
// Solidity: function execute(((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256) sell, ((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256) buy) payable returns()
func (_Exchange *ExchangeTransactor) Execute2(opts *bind.TransactOpts, sell Input, buy Input) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "execute", sell, buy)
}

// Execute2 is a paid mutator transaction binding the contract method 0x9a1fc3a7.
//
// Solidity: function execute(((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256) sell, ((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256) buy) payable returns()
func (_Exchange *ExchangeSession) Execute2(sell Input, buy Input) (*types.Transaction, error) {
	return _Exchange.Contract.Execute2(&_Exchange.TransactOpts, sell, buy)
}

// Execute2 is a paid mutator transaction binding the contract method 0x9a1fc3a7.
//
// Solidity: function execute(((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256) sell, ((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256) buy) payable returns()
func (_Exchange *ExchangeTransactorSession) Execute2(sell Input, buy Input) (*types.Transaction, error) {
	return _Exchange.Contract.Execute2(&_Exchange.TransactOpts, sell, buy)
}

// IncrementNonce is a paid mutator transaction binding the contract method 0x627cdcb9.
//
// Solidity: function incrementNonce() returns()
func (_Exchange *ExchangeTransactor) IncrementNonce(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "incrementNonce")
}

// IncrementNonce is a paid mutator transaction binding the contract method 0x627cdcb9.
//
// Solidity: function incrementNonce() returns()
func (_Exchange *ExchangeSession) IncrementNonce() (*types.Transaction, error) {
	return _Exchange.Contract.IncrementNonce(&_Exchange.TransactOpts)
}

// IncrementNonce is a paid mutator transaction binding the contract method 0x627cdcb9.
//
// Solidity: function incrementNonce() returns()
func (_Exchange *ExchangeTransactorSession) IncrementNonce() (*types.Transaction, error) {
	return _Exchange.Contract.IncrementNonce(&_Exchange.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xcf756fdf.
//
// Solidity: function initialize(address _executionDelegate, address _policyManager, address _oracle, uint256 _blockRange) returns()
func (_Exchange *ExchangeTransactor) Initialize(opts *bind.TransactOpts, _executionDelegate common.Address, _policyManager common.Address, _oracle common.Address, _blockRange *big.Int) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "initialize", _executionDelegate, _policyManager, _oracle, _blockRange)
}

// Initialize is a paid mutator transaction binding the contract method 0xcf756fdf.
//
// Solidity: function initialize(address _executionDelegate, address _policyManager, address _oracle, uint256 _blockRange) returns()
func (_Exchange *ExchangeSession) Initialize(_executionDelegate common.Address, _policyManager common.Address, _oracle common.Address, _blockRange *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.Initialize(&_Exchange.TransactOpts, _executionDelegate, _policyManager, _oracle, _blockRange)
}

// Initialize is a paid mutator transaction binding the contract method 0xcf756fdf.
//
// Solidity: function initialize(address _executionDelegate, address _policyManager, address _oracle, uint256 _blockRange) returns()
func (_Exchange *ExchangeTransactorSession) Initialize(_executionDelegate common.Address, _policyManager common.Address, _oracle common.Address, _blockRange *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.Initialize(&_Exchange.TransactOpts, _executionDelegate, _policyManager, _oracle, _blockRange)
}

// Open is a paid mutator transaction binding the contract method 0xfcfff16f.
//
// Solidity: function open() returns()
func (_Exchange *ExchangeTransactor) Open(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "open")
}

// Open is a paid mutator transaction binding the contract method 0xfcfff16f.
//
// Solidity: function open() returns()
func (_Exchange *ExchangeSession) Open() (*types.Transaction, error) {
	return _Exchange.Contract.Open(&_Exchange.TransactOpts)
}

// Open is a paid mutator transaction binding the contract method 0xfcfff16f.
//
// Solidity: function open() returns()
func (_Exchange *ExchangeTransactorSession) Open() (*types.Transaction, error) {
	return _Exchange.Contract.Open(&_Exchange.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Exchange *ExchangeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Exchange *ExchangeSession) RenounceOwnership() (*types.Transaction, error) {
	return _Exchange.Contract.RenounceOwnership(&_Exchange.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Exchange *ExchangeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Exchange.Contract.RenounceOwnership(&_Exchange.TransactOpts)
}

// SetBlockRange is a paid mutator transaction binding the contract method 0x6992aa36.
//
// Solidity: function setBlockRange(uint256 _blockRange) returns()
func (_Exchange *ExchangeTransactor) SetBlockRange(opts *bind.TransactOpts, _blockRange *big.Int) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "setBlockRange", _blockRange)
}

// SetBlockRange is a paid mutator transaction binding the contract method 0x6992aa36.
//
// Solidity: function setBlockRange(uint256 _blockRange) returns()
func (_Exchange *ExchangeSession) SetBlockRange(_blockRange *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.SetBlockRange(&_Exchange.TransactOpts, _blockRange)
}

// SetBlockRange is a paid mutator transaction binding the contract method 0x6992aa36.
//
// Solidity: function setBlockRange(uint256 _blockRange) returns()
func (_Exchange *ExchangeTransactorSession) SetBlockRange(_blockRange *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.SetBlockRange(&_Exchange.TransactOpts, _blockRange)
}

// SetExecutionDelegate is a paid mutator transaction binding the contract method 0x037c9be2.
//
// Solidity: function setExecutionDelegate(address _executionDelegate) returns()
func (_Exchange *ExchangeTransactor) SetExecutionDelegate(opts *bind.TransactOpts, _executionDelegate common.Address) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "setExecutionDelegate", _executionDelegate)
}

// SetExecutionDelegate is a paid mutator transaction binding the contract method 0x037c9be2.
//
// Solidity: function setExecutionDelegate(address _executionDelegate) returns()
func (_Exchange *ExchangeSession) SetExecutionDelegate(_executionDelegate common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.SetExecutionDelegate(&_Exchange.TransactOpts, _executionDelegate)
}

// SetExecutionDelegate is a paid mutator transaction binding the contract method 0x037c9be2.
//
// Solidity: function setExecutionDelegate(address _executionDelegate) returns()
func (_Exchange *ExchangeTransactorSession) SetExecutionDelegate(_executionDelegate common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.SetExecutionDelegate(&_Exchange.TransactOpts, _executionDelegate)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address _oracle) returns()
func (_Exchange *ExchangeTransactor) SetOracle(opts *bind.TransactOpts, _oracle common.Address) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "setOracle", _oracle)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address _oracle) returns()
func (_Exchange *ExchangeSession) SetOracle(_oracle common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.SetOracle(&_Exchange.TransactOpts, _oracle)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address _oracle) returns()
func (_Exchange *ExchangeTransactorSession) SetOracle(_oracle common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.SetOracle(&_Exchange.TransactOpts, _oracle)
}

// SetPolicyManager is a paid mutator transaction binding the contract method 0xadde41e1.
//
// Solidity: function setPolicyManager(address _policyManager) returns()
func (_Exchange *ExchangeTransactor) SetPolicyManager(opts *bind.TransactOpts, _policyManager common.Address) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "setPolicyManager", _policyManager)
}

// SetPolicyManager is a paid mutator transaction binding the contract method 0xadde41e1.
//
// Solidity: function setPolicyManager(address _policyManager) returns()
func (_Exchange *ExchangeSession) SetPolicyManager(_policyManager common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.SetPolicyManager(&_Exchange.TransactOpts, _policyManager)
}

// SetPolicyManager is a paid mutator transaction binding the contract method 0xadde41e1.
//
// Solidity: function setPolicyManager(address _policyManager) returns()
func (_Exchange *ExchangeTransactorSession) SetPolicyManager(_policyManager common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.SetPolicyManager(&_Exchange.TransactOpts, _policyManager)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Exchange *ExchangeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Exchange *ExchangeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.TransferOwnership(&_Exchange.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Exchange *ExchangeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.TransferOwnership(&_Exchange.TransactOpts, newOwner)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Exchange *ExchangeTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Exchange *ExchangeSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.UpgradeTo(&_Exchange.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Exchange *ExchangeTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.UpgradeTo(&_Exchange.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Exchange *ExchangeTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Exchange *ExchangeSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Exchange.Contract.UpgradeToAndCall(&_Exchange.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Exchange *ExchangeTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Exchange.Contract.UpgradeToAndCall(&_Exchange.TransactOpts, newImplementation, data)
}

// ExchangeAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Exchange contract.
type ExchangeAdminChangedIterator struct {
	Event *ExchangeAdminChanged // Event containing the contract specifics and raw log

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
func (it *ExchangeAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeAdminChanged)
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
		it.Event = new(ExchangeAdminChanged)
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
func (it *ExchangeAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeAdminChanged represents a AdminChanged event raised by the Exchange contract.
type ExchangeAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Exchange *ExchangeFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*ExchangeAdminChangedIterator, error) {

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &ExchangeAdminChangedIterator{contract: _Exchange.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Exchange *ExchangeFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *ExchangeAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeAdminChanged)
				if err := _Exchange.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Exchange *ExchangeFilterer) ParseAdminChanged(log types.Log) (*ExchangeAdminChanged, error) {
	event := new(ExchangeAdminChanged)
	if err := _Exchange.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Exchange contract.
type ExchangeBeaconUpgradedIterator struct {
	Event *ExchangeBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *ExchangeBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeBeaconUpgraded)
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
		it.Event = new(ExchangeBeaconUpgraded)
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
func (it *ExchangeBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeBeaconUpgraded represents a BeaconUpgraded event raised by the Exchange contract.
type ExchangeBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Exchange *ExchangeFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*ExchangeBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeBeaconUpgradedIterator{contract: _Exchange.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Exchange *ExchangeFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *ExchangeBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeBeaconUpgraded)
				if err := _Exchange.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Exchange *ExchangeFilterer) ParseBeaconUpgraded(log types.Log) (*ExchangeBeaconUpgraded, error) {
	event := new(ExchangeBeaconUpgraded)
	if err := _Exchange.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeClosedIterator is returned from FilterClosed and is used to iterate over the raw logs and unpacked data for Closed events raised by the Exchange contract.
type ExchangeClosedIterator struct {
	Event *ExchangeClosed // Event containing the contract specifics and raw log

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
func (it *ExchangeClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeClosed)
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
		it.Event = new(ExchangeClosed)
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
func (it *ExchangeClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeClosed represents a Closed event raised by the Exchange contract.
type ExchangeClosed struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterClosed is a free log retrieval operation binding the contract event 0x1cdde67b72a90f19919ac732a437ac2f7a10fc128d28c2a6e525d89ce5cd9d3a.
//
// Solidity: event Closed()
func (_Exchange *ExchangeFilterer) FilterClosed(opts *bind.FilterOpts) (*ExchangeClosedIterator, error) {

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "Closed")
	if err != nil {
		return nil, err
	}
	return &ExchangeClosedIterator{contract: _Exchange.contract, event: "Closed", logs: logs, sub: sub}, nil
}

// WatchClosed is a free log subscription operation binding the contract event 0x1cdde67b72a90f19919ac732a437ac2f7a10fc128d28c2a6e525d89ce5cd9d3a.
//
// Solidity: event Closed()
func (_Exchange *ExchangeFilterer) WatchClosed(opts *bind.WatchOpts, sink chan<- *ExchangeClosed) (event.Subscription, error) {

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "Closed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeClosed)
				if err := _Exchange.contract.UnpackLog(event, "Closed", log); err != nil {
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

// ParseClosed is a log parse operation binding the contract event 0x1cdde67b72a90f19919ac732a437ac2f7a10fc128d28c2a6e525d89ce5cd9d3a.
//
// Solidity: event Closed()
func (_Exchange *ExchangeFilterer) ParseClosed(log types.Log) (*ExchangeClosed, error) {
	event := new(ExchangeClosed)
	if err := _Exchange.contract.UnpackLog(event, "Closed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Exchange contract.
type ExchangeInitializedIterator struct {
	Event *ExchangeInitialized // Event containing the contract specifics and raw log

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
func (it *ExchangeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeInitialized)
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
		it.Event = new(ExchangeInitialized)
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
func (it *ExchangeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeInitialized represents a Initialized event raised by the Exchange contract.
type ExchangeInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Exchange *ExchangeFilterer) FilterInitialized(opts *bind.FilterOpts) (*ExchangeInitializedIterator, error) {

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ExchangeInitializedIterator{contract: _Exchange.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Exchange *ExchangeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ExchangeInitialized) (event.Subscription, error) {

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeInitialized)
				if err := _Exchange.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Exchange *ExchangeFilterer) ParseInitialized(log types.Log) (*ExchangeInitialized, error) {
	event := new(ExchangeInitialized)
	if err := _Exchange.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeNewBlockRangeIterator is returned from FilterNewBlockRange and is used to iterate over the raw logs and unpacked data for NewBlockRange events raised by the Exchange contract.
type ExchangeNewBlockRangeIterator struct {
	Event *ExchangeNewBlockRange // Event containing the contract specifics and raw log

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
func (it *ExchangeNewBlockRangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeNewBlockRange)
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
		it.Event = new(ExchangeNewBlockRange)
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
func (it *ExchangeNewBlockRangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeNewBlockRangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeNewBlockRange represents a NewBlockRange event raised by the Exchange contract.
type ExchangeNewBlockRange struct {
	BlockRange *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewBlockRange is a free log retrieval operation binding the contract event 0x7706177c541ba1b858371bfc568aa77450b4713bbdbba62c730d4484ab6c1251.
//
// Solidity: event NewBlockRange(uint256 blockRange)
func (_Exchange *ExchangeFilterer) FilterNewBlockRange(opts *bind.FilterOpts) (*ExchangeNewBlockRangeIterator, error) {

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "NewBlockRange")
	if err != nil {
		return nil, err
	}
	return &ExchangeNewBlockRangeIterator{contract: _Exchange.contract, event: "NewBlockRange", logs: logs, sub: sub}, nil
}

// WatchNewBlockRange is a free log subscription operation binding the contract event 0x7706177c541ba1b858371bfc568aa77450b4713bbdbba62c730d4484ab6c1251.
//
// Solidity: event NewBlockRange(uint256 blockRange)
func (_Exchange *ExchangeFilterer) WatchNewBlockRange(opts *bind.WatchOpts, sink chan<- *ExchangeNewBlockRange) (event.Subscription, error) {

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "NewBlockRange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeNewBlockRange)
				if err := _Exchange.contract.UnpackLog(event, "NewBlockRange", log); err != nil {
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

// ParseNewBlockRange is a log parse operation binding the contract event 0x7706177c541ba1b858371bfc568aa77450b4713bbdbba62c730d4484ab6c1251.
//
// Solidity: event NewBlockRange(uint256 blockRange)
func (_Exchange *ExchangeFilterer) ParseNewBlockRange(log types.Log) (*ExchangeNewBlockRange, error) {
	event := new(ExchangeNewBlockRange)
	if err := _Exchange.contract.UnpackLog(event, "NewBlockRange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeNewExecutionDelegateIterator is returned from FilterNewExecutionDelegate and is used to iterate over the raw logs and unpacked data for NewExecutionDelegate events raised by the Exchange contract.
type ExchangeNewExecutionDelegateIterator struct {
	Event *ExchangeNewExecutionDelegate // Event containing the contract specifics and raw log

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
func (it *ExchangeNewExecutionDelegateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeNewExecutionDelegate)
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
		it.Event = new(ExchangeNewExecutionDelegate)
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
func (it *ExchangeNewExecutionDelegateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeNewExecutionDelegateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeNewExecutionDelegate represents a NewExecutionDelegate event raised by the Exchange contract.
type ExchangeNewExecutionDelegate struct {
	ExecutionDelegate common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterNewExecutionDelegate is a free log retrieval operation binding the contract event 0xf9a0f356a7ef079355de09d32ce45cc3cfabc8f118681c19a17501f005a376ac.
//
// Solidity: event NewExecutionDelegate(address indexed executionDelegate)
func (_Exchange *ExchangeFilterer) FilterNewExecutionDelegate(opts *bind.FilterOpts, executionDelegate []common.Address) (*ExchangeNewExecutionDelegateIterator, error) {

	var executionDelegateRule []interface{}
	for _, executionDelegateItem := range executionDelegate {
		executionDelegateRule = append(executionDelegateRule, executionDelegateItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "NewExecutionDelegate", executionDelegateRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeNewExecutionDelegateIterator{contract: _Exchange.contract, event: "NewExecutionDelegate", logs: logs, sub: sub}, nil
}

// WatchNewExecutionDelegate is a free log subscription operation binding the contract event 0xf9a0f356a7ef079355de09d32ce45cc3cfabc8f118681c19a17501f005a376ac.
//
// Solidity: event NewExecutionDelegate(address indexed executionDelegate)
func (_Exchange *ExchangeFilterer) WatchNewExecutionDelegate(opts *bind.WatchOpts, sink chan<- *ExchangeNewExecutionDelegate, executionDelegate []common.Address) (event.Subscription, error) {

	var executionDelegateRule []interface{}
	for _, executionDelegateItem := range executionDelegate {
		executionDelegateRule = append(executionDelegateRule, executionDelegateItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "NewExecutionDelegate", executionDelegateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeNewExecutionDelegate)
				if err := _Exchange.contract.UnpackLog(event, "NewExecutionDelegate", log); err != nil {
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

// ParseNewExecutionDelegate is a log parse operation binding the contract event 0xf9a0f356a7ef079355de09d32ce45cc3cfabc8f118681c19a17501f005a376ac.
//
// Solidity: event NewExecutionDelegate(address indexed executionDelegate)
func (_Exchange *ExchangeFilterer) ParseNewExecutionDelegate(log types.Log) (*ExchangeNewExecutionDelegate, error) {
	event := new(ExchangeNewExecutionDelegate)
	if err := _Exchange.contract.UnpackLog(event, "NewExecutionDelegate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeNewOracleIterator is returned from FilterNewOracle and is used to iterate over the raw logs and unpacked data for NewOracle events raised by the Exchange contract.
type ExchangeNewOracleIterator struct {
	Event *ExchangeNewOracle // Event containing the contract specifics and raw log

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
func (it *ExchangeNewOracleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeNewOracle)
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
		it.Event = new(ExchangeNewOracle)
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
func (it *ExchangeNewOracleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeNewOracleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeNewOracle represents a NewOracle event raised by the Exchange contract.
type ExchangeNewOracle struct {
	Oracle common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNewOracle is a free log retrieval operation binding the contract event 0xb3eacd0e351fafdfefdec84e1cd19679b38dbcd63ea7c2c24da17fd2bc3b3c0e.
//
// Solidity: event NewOracle(address indexed oracle)
func (_Exchange *ExchangeFilterer) FilterNewOracle(opts *bind.FilterOpts, oracle []common.Address) (*ExchangeNewOracleIterator, error) {

	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "NewOracle", oracleRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeNewOracleIterator{contract: _Exchange.contract, event: "NewOracle", logs: logs, sub: sub}, nil
}

// WatchNewOracle is a free log subscription operation binding the contract event 0xb3eacd0e351fafdfefdec84e1cd19679b38dbcd63ea7c2c24da17fd2bc3b3c0e.
//
// Solidity: event NewOracle(address indexed oracle)
func (_Exchange *ExchangeFilterer) WatchNewOracle(opts *bind.WatchOpts, sink chan<- *ExchangeNewOracle, oracle []common.Address) (event.Subscription, error) {

	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "NewOracle", oracleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeNewOracle)
				if err := _Exchange.contract.UnpackLog(event, "NewOracle", log); err != nil {
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

// ParseNewOracle is a log parse operation binding the contract event 0xb3eacd0e351fafdfefdec84e1cd19679b38dbcd63ea7c2c24da17fd2bc3b3c0e.
//
// Solidity: event NewOracle(address indexed oracle)
func (_Exchange *ExchangeFilterer) ParseNewOracle(log types.Log) (*ExchangeNewOracle, error) {
	event := new(ExchangeNewOracle)
	if err := _Exchange.contract.UnpackLog(event, "NewOracle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeNewPolicyManagerIterator is returned from FilterNewPolicyManager and is used to iterate over the raw logs and unpacked data for NewPolicyManager events raised by the Exchange contract.
type ExchangeNewPolicyManagerIterator struct {
	Event *ExchangeNewPolicyManager // Event containing the contract specifics and raw log

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
func (it *ExchangeNewPolicyManagerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeNewPolicyManager)
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
		it.Event = new(ExchangeNewPolicyManager)
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
func (it *ExchangeNewPolicyManagerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeNewPolicyManagerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeNewPolicyManager represents a NewPolicyManager event raised by the Exchange contract.
type ExchangeNewPolicyManager struct {
	PolicyManager common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNewPolicyManager is a free log retrieval operation binding the contract event 0xdbe18f3fd927cc2aefe380ffd2abfdb8e13f0239c0258ccfc84c3d8fdd8c0418.
//
// Solidity: event NewPolicyManager(address indexed policyManager)
func (_Exchange *ExchangeFilterer) FilterNewPolicyManager(opts *bind.FilterOpts, policyManager []common.Address) (*ExchangeNewPolicyManagerIterator, error) {

	var policyManagerRule []interface{}
	for _, policyManagerItem := range policyManager {
		policyManagerRule = append(policyManagerRule, policyManagerItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "NewPolicyManager", policyManagerRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeNewPolicyManagerIterator{contract: _Exchange.contract, event: "NewPolicyManager", logs: logs, sub: sub}, nil
}

// WatchNewPolicyManager is a free log subscription operation binding the contract event 0xdbe18f3fd927cc2aefe380ffd2abfdb8e13f0239c0258ccfc84c3d8fdd8c0418.
//
// Solidity: event NewPolicyManager(address indexed policyManager)
func (_Exchange *ExchangeFilterer) WatchNewPolicyManager(opts *bind.WatchOpts, sink chan<- *ExchangeNewPolicyManager, policyManager []common.Address) (event.Subscription, error) {

	var policyManagerRule []interface{}
	for _, policyManagerItem := range policyManager {
		policyManagerRule = append(policyManagerRule, policyManagerItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "NewPolicyManager", policyManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeNewPolicyManager)
				if err := _Exchange.contract.UnpackLog(event, "NewPolicyManager", log); err != nil {
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

// ParseNewPolicyManager is a log parse operation binding the contract event 0xdbe18f3fd927cc2aefe380ffd2abfdb8e13f0239c0258ccfc84c3d8fdd8c0418.
//
// Solidity: event NewPolicyManager(address indexed policyManager)
func (_Exchange *ExchangeFilterer) ParseNewPolicyManager(log types.Log) (*ExchangeNewPolicyManager, error) {
	event := new(ExchangeNewPolicyManager)
	if err := _Exchange.contract.UnpackLog(event, "NewPolicyManager", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeNonceIncrementedIterator is returned from FilterNonceIncremented and is used to iterate over the raw logs and unpacked data for NonceIncremented events raised by the Exchange contract.
type ExchangeNonceIncrementedIterator struct {
	Event *ExchangeNonceIncremented // Event containing the contract specifics and raw log

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
func (it *ExchangeNonceIncrementedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeNonceIncremented)
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
		it.Event = new(ExchangeNonceIncremented)
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
func (it *ExchangeNonceIncrementedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeNonceIncrementedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeNonceIncremented represents a NonceIncremented event raised by the Exchange contract.
type ExchangeNonceIncremented struct {
	Trader   common.Address
	NewNonce *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNonceIncremented is a free log retrieval operation binding the contract event 0xa82a649bbd060c9099cd7b7326e2b0dc9e9af0836480e0f849dc9eaa79710b3b.
//
// Solidity: event NonceIncremented(address indexed trader, uint256 newNonce)
func (_Exchange *ExchangeFilterer) FilterNonceIncremented(opts *bind.FilterOpts, trader []common.Address) (*ExchangeNonceIncrementedIterator, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "NonceIncremented", traderRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeNonceIncrementedIterator{contract: _Exchange.contract, event: "NonceIncremented", logs: logs, sub: sub}, nil
}

// WatchNonceIncremented is a free log subscription operation binding the contract event 0xa82a649bbd060c9099cd7b7326e2b0dc9e9af0836480e0f849dc9eaa79710b3b.
//
// Solidity: event NonceIncremented(address indexed trader, uint256 newNonce)
func (_Exchange *ExchangeFilterer) WatchNonceIncremented(opts *bind.WatchOpts, sink chan<- *ExchangeNonceIncremented, trader []common.Address) (event.Subscription, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "NonceIncremented", traderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeNonceIncremented)
				if err := _Exchange.contract.UnpackLog(event, "NonceIncremented", log); err != nil {
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

// ParseNonceIncremented is a log parse operation binding the contract event 0xa82a649bbd060c9099cd7b7326e2b0dc9e9af0836480e0f849dc9eaa79710b3b.
//
// Solidity: event NonceIncremented(address indexed trader, uint256 newNonce)
func (_Exchange *ExchangeFilterer) ParseNonceIncremented(log types.Log) (*ExchangeNonceIncremented, error) {
	event := new(ExchangeNonceIncremented)
	if err := _Exchange.contract.UnpackLog(event, "NonceIncremented", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeOpenedIterator is returned from FilterOpened and is used to iterate over the raw logs and unpacked data for Opened events raised by the Exchange contract.
type ExchangeOpenedIterator struct {
	Event *ExchangeOpened // Event containing the contract specifics and raw log

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
func (it *ExchangeOpenedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeOpened)
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
		it.Event = new(ExchangeOpened)
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
func (it *ExchangeOpenedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeOpenedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeOpened represents a Opened event raised by the Exchange contract.
type ExchangeOpened struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterOpened is a free log retrieval operation binding the contract event 0xd1dcd00534373f20882b79e6ab6875a5c358c5bd576448757ed50e63069ab518.
//
// Solidity: event Opened()
func (_Exchange *ExchangeFilterer) FilterOpened(opts *bind.FilterOpts) (*ExchangeOpenedIterator, error) {

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "Opened")
	if err != nil {
		return nil, err
	}
	return &ExchangeOpenedIterator{contract: _Exchange.contract, event: "Opened", logs: logs, sub: sub}, nil
}

// WatchOpened is a free log subscription operation binding the contract event 0xd1dcd00534373f20882b79e6ab6875a5c358c5bd576448757ed50e63069ab518.
//
// Solidity: event Opened()
func (_Exchange *ExchangeFilterer) WatchOpened(opts *bind.WatchOpts, sink chan<- *ExchangeOpened) (event.Subscription, error) {

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "Opened")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeOpened)
				if err := _Exchange.contract.UnpackLog(event, "Opened", log); err != nil {
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

// ParseOpened is a log parse operation binding the contract event 0xd1dcd00534373f20882b79e6ab6875a5c358c5bd576448757ed50e63069ab518.
//
// Solidity: event Opened()
func (_Exchange *ExchangeFilterer) ParseOpened(log types.Log) (*ExchangeOpened, error) {
	event := new(ExchangeOpened)
	if err := _Exchange.contract.UnpackLog(event, "Opened", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeOrderCancelledIterator is returned from FilterOrderCancelled and is used to iterate over the raw logs and unpacked data for OrderCancelled events raised by the Exchange contract.
type ExchangeOrderCancelledIterator struct {
	Event *ExchangeOrderCancelled // Event containing the contract specifics and raw log

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
func (it *ExchangeOrderCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeOrderCancelled)
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
		it.Event = new(ExchangeOrderCancelled)
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
func (it *ExchangeOrderCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeOrderCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeOrderCancelled represents a OrderCancelled event raised by the Exchange contract.
type ExchangeOrderCancelled struct {
	Hash [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOrderCancelled is a free log retrieval operation binding the contract event 0x5152abf959f6564662358c2e52b702259b78bac5ee7842a0f01937e670efcc7d.
//
// Solidity: event OrderCancelled(bytes32 hash)
func (_Exchange *ExchangeFilterer) FilterOrderCancelled(opts *bind.FilterOpts) (*ExchangeOrderCancelledIterator, error) {

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "OrderCancelled")
	if err != nil {
		return nil, err
	}
	return &ExchangeOrderCancelledIterator{contract: _Exchange.contract, event: "OrderCancelled", logs: logs, sub: sub}, nil
}

// WatchOrderCancelled is a free log subscription operation binding the contract event 0x5152abf959f6564662358c2e52b702259b78bac5ee7842a0f01937e670efcc7d.
//
// Solidity: event OrderCancelled(bytes32 hash)
func (_Exchange *ExchangeFilterer) WatchOrderCancelled(opts *bind.WatchOpts, sink chan<- *ExchangeOrderCancelled) (event.Subscription, error) {

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "OrderCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeOrderCancelled)
				if err := _Exchange.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
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

// ParseOrderCancelled is a log parse operation binding the contract event 0x5152abf959f6564662358c2e52b702259b78bac5ee7842a0f01937e670efcc7d.
//
// Solidity: event OrderCancelled(bytes32 hash)
func (_Exchange *ExchangeFilterer) ParseOrderCancelled(log types.Log) (*ExchangeOrderCancelled, error) {
	event := new(ExchangeOrderCancelled)
	if err := _Exchange.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeOrdersMatchedIterator is returned from FilterOrdersMatched and is used to iterate over the raw logs and unpacked data for OrdersMatched events raised by the Exchange contract.
type ExchangeOrdersMatchedIterator struct {
	Event *ExchangeOrdersMatched // Event containing the contract specifics and raw log

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
func (it *ExchangeOrdersMatchedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeOrdersMatched)
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
		it.Event = new(ExchangeOrdersMatched)
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
func (it *ExchangeOrdersMatchedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeOrdersMatchedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeOrdersMatched represents a OrdersMatched event raised by the Exchange contract.
type ExchangeOrdersMatched struct {
	Maker    common.Address
	Taker    common.Address
	Sell     Order
	SellHash [32]byte
	Buy      Order
	BuyHash  [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOrdersMatched is a free log retrieval operation binding the contract event 0x61cbb2a3dee0b6064c2e681aadd61677fb4ef319f0b547508d495626f5a62f64.
//
// Solidity: event OrdersMatched(address indexed maker, address indexed taker, (address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes) sell, bytes32 sellHash, (address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes) buy, bytes32 buyHash)
func (_Exchange *ExchangeFilterer) FilterOrdersMatched(opts *bind.FilterOpts, maker []common.Address, taker []common.Address) (*ExchangeOrdersMatchedIterator, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}
	var takerRule []interface{}
	for _, takerItem := range taker {
		takerRule = append(takerRule, takerItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "OrdersMatched", makerRule, takerRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeOrdersMatchedIterator{contract: _Exchange.contract, event: "OrdersMatched", logs: logs, sub: sub}, nil
}

// WatchOrdersMatched is a free log subscription operation binding the contract event 0x61cbb2a3dee0b6064c2e681aadd61677fb4ef319f0b547508d495626f5a62f64.
//
// Solidity: event OrdersMatched(address indexed maker, address indexed taker, (address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes) sell, bytes32 sellHash, (address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes) buy, bytes32 buyHash)
func (_Exchange *ExchangeFilterer) WatchOrdersMatched(opts *bind.WatchOpts, sink chan<- *ExchangeOrdersMatched, maker []common.Address, taker []common.Address) (event.Subscription, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}
	var takerRule []interface{}
	for _, takerItem := range taker {
		takerRule = append(takerRule, takerItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "OrdersMatched", makerRule, takerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeOrdersMatched)
				if err := _Exchange.contract.UnpackLog(event, "OrdersMatched", log); err != nil {
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

// ParseOrdersMatched is a log parse operation binding the contract event 0x61cbb2a3dee0b6064c2e681aadd61677fb4ef319f0b547508d495626f5a62f64.
//
// Solidity: event OrdersMatched(address indexed maker, address indexed taker, (address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes) sell, bytes32 sellHash, (address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes) buy, bytes32 buyHash)
func (_Exchange *ExchangeFilterer) ParseOrdersMatched(log types.Log) (*ExchangeOrdersMatched, error) {
	event := new(ExchangeOrdersMatched)
	if err := _Exchange.contract.UnpackLog(event, "OrdersMatched", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Exchange contract.
type ExchangeOwnershipTransferredIterator struct {
	Event *ExchangeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ExchangeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeOwnershipTransferred)
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
		it.Event = new(ExchangeOwnershipTransferred)
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
func (it *ExchangeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeOwnershipTransferred represents a OwnershipTransferred event raised by the Exchange contract.
type ExchangeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Exchange *ExchangeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ExchangeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeOwnershipTransferredIterator{contract: _Exchange.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Exchange *ExchangeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ExchangeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeOwnershipTransferred)
				if err := _Exchange.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Exchange *ExchangeFilterer) ParseOwnershipTransferred(log types.Log) (*ExchangeOwnershipTransferred, error) {
	event := new(ExchangeOwnershipTransferred)
	if err := _Exchange.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Exchange contract.
type ExchangeUpgradedIterator struct {
	Event *ExchangeUpgraded // Event containing the contract specifics and raw log

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
func (it *ExchangeUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeUpgraded)
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
		it.Event = new(ExchangeUpgraded)
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
func (it *ExchangeUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeUpgraded represents a Upgraded event raised by the Exchange contract.
type ExchangeUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Exchange *ExchangeFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ExchangeUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeUpgradedIterator{contract: _Exchange.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Exchange *ExchangeFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ExchangeUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeUpgraded)
				if err := _Exchange.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Exchange *ExchangeFilterer) ParseUpgraded(log types.Log) (*ExchangeUpgraded, error) {
	event := new(ExchangeUpgraded)
	if err := _Exchange.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
