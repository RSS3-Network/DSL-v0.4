// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tofunft

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

// MarketNGDetail is an auto generated low-level Go binding around an user-defined struct.
type MarketNGDetail struct {
	IntentionHash [32]byte
	Signer        common.Address
	TxDeadline    *big.Int
	Salt          [32]byte
	Id            *big.Int
	Opcode        uint8
	Caller        common.Address
	Currency      common.Address
	Price         *big.Int
	IncentiveRate *big.Int
	Settlement    MarketNGSettlement
	Bundle        []MarketNGTokenPair
	Deadline      *big.Int
}

// MarketNGIntention is an auto generated low-level Go binding around an user-defined struct.
type MarketNGIntention struct {
	User     common.Address
	Bundle   []MarketNGTokenPair
	Currency common.Address
	Price    *big.Int
	Deadline *big.Int
	Salt     [32]byte
	Kind     uint8
}

// MarketNGInventory is an auto generated low-level Go binding around an user-defined struct.
type MarketNGInventory struct {
	Seller   common.Address
	Buyer    common.Address
	Currency common.Address
	Price    *big.Int
	NetPrice *big.Int
	Deadline *big.Int
	Kind     uint8
	Status   uint8
}

// MarketNGPair721 is an auto generated low-level Go binding around an user-defined struct.
type MarketNGPair721 struct {
	Token   common.Address
	TokenId *big.Int
}

// MarketNGSettlement is an auto generated low-level Go binding around an user-defined struct.
type MarketNGSettlement struct {
	Coupons           []*big.Int
	FeeRate           *big.Int
	RoyaltyRate       *big.Int
	BuyerCashbackRate *big.Int
	FeeAddress        common.Address
	RoyaltyAddress    common.Address
}

// MarketNGSwap is an auto generated low-level Go binding around an user-defined struct.
type MarketNGSwap struct {
	Salt     [32]byte
	Creator  common.Address
	Deadline *big.Int
	Has      []MarketNGPair721
	Wants    []MarketNGPair721
}

// MarketNGTokenPair is an auto generated low-level Go binding around an user-defined struct.
type MarketNGTokenPair struct {
	Token    common.Address
	TokenId  *big.Int
	Amount   *big.Int
	Kind     uint8
	MintData []byte
}

// MarketplaceMetaData contains all meta data concerning the Marketplace contract.
var MarketplaceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIWETH\",\"name\":\"weth_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"refund\",\"type\":\"uint256\"}],\"name\":\"EvAuctionRefund\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"couponId\",\"type\":\"uint256\"}],\"name\":\"EvCouponSpent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"currency\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"}],\"indexed\":false,\"internalType\":\"structMarketNG.Inventory\",\"name\":\"inventory\",\"type\":\"tuple\"}],\"name\":\"EvInventoryUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isRemoval\",\"type\":\"bool\"}],\"name\":\"EvMarketSignerUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EvSettingsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractIERC721\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"internalType\":\"structMarketNG.Pair721[]\",\"name\":\"has\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"contractIERC721\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"internalType\":\"structMarketNG.Pair721[]\",\"name\":\"wants\",\"type\":\"tuple[]\"}],\"indexed\":false,\"internalType\":\"structMarketNG.Swap\",\"name\":\"req\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"swapper\",\"type\":\"address\"}],\"name\":\"EvSwapped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"KIND_AUCTION\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"KIND_BUY\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"KIND_SELL\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OP_ACCEPT_AUCTION\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OP_ACCEPT_BUY\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OP_BID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OP_BUY\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OP_CANCEL_BUY\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OP_COMPLETE_AUCTION\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OP_COMPLETE_BUY\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OP_COMPLETE_SELL\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OP_MAX\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OP_MIN\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OP_REJECT_BUY\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RATE_BASE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STATUS_CANCELLED\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STATUS_DONE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STATUS_OPEN\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN_1155\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN_721\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN_MINT\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"cancelBuys\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"couponSpent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"noBundle\",\"type\":\"bool\"}],\"name\":\"emergencyCancelAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"hasInv\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"op\",\"type\":\"uint8\"}],\"name\":\"hasSignedIntention\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"currency\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"inCaseMoneyGetsStuck\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"inventories\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"currency\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"inventoryTokenCounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"inventoryTokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"mintData\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"isAuction\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"isAuctionOpen\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"invId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"isBundleApproved\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"isBuy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"isBuyOpen\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"isExpired\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"isSell\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"isSignatureValid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"isStatusOpen\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"marketSigners\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minAuctionDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minAuctionIncrement\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"mintData\",\"type\":\"bytes\"}],\"internalType\":\"structMarketNG.TokenPair[]\",\"name\":\"bundle\",\"type\":\"tuple[]\"},{\"internalType\":\"contractIERC20\",\"name\":\"currency\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"}],\"internalType\":\"structMarketNG.Intention\",\"name\":\"intent\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"intentionHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"txDeadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"opcode\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"currency\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"incentiveRate\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"coupons\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"feeRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"royaltyRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"buyerCashbackRate\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"royaltyAddress\",\"type\":\"address\"}],\"internalType\":\"structMarketNG.Settlement\",\"name\":\"settlement\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"mintData\",\"type\":\"bytes\"}],\"internalType\":\"structMarketNG.TokenPair[]\",\"name\":\"bundle\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structMarketNG.Detail\",\"name\":\"detail\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"sigIntent\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sigDetail\",\"type\":\"bytes\"}],\"name\":\"run\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"contractIERC721\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"internalType\":\"structMarketNG.Pair721[]\",\"name\":\"tokens\",\"type\":\"tuple[]\"}],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractIERC721\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"internalType\":\"structMarketNG.Pair721[]\",\"name\":\"has\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"contractIERC721\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"internalType\":\"structMarketNG.Pair721[]\",\"name\":\"wants\",\"type\":\"tuple[]\"}],\"internalType\":\"structMarketNG.Swap\",\"name\":\"req\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"swap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minAuctionIncrement_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAuctionDuration_\",\"type\":\"uint256\"}],\"name\":\"updateSettings\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"remove\",\"type\":\"bool\"}],\"name\":\"updateSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"weth\",\"outputs\":[{\"internalType\":\"contractIWETH\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// MarketplaceABI is the input ABI used to generate the binding from.
// Deprecated: Use MarketplaceMetaData.ABI instead.
var MarketplaceABI = MarketplaceMetaData.ABI

// Marketplace is an auto generated Go binding around an Ethereum contract.
type Marketplace struct {
	MarketplaceCaller     // Read-only binding to the contract
	MarketplaceTransactor // Write-only binding to the contract
	MarketplaceFilterer   // Log filterer for contract events
}

// MarketplaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type MarketplaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketplaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MarketplaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketplaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MarketplaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketplaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MarketplaceSession struct {
	Contract     *Marketplace      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarketplaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MarketplaceCallerSession struct {
	Contract *MarketplaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// MarketplaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MarketplaceTransactorSession struct {
	Contract     *MarketplaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MarketplaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type MarketplaceRaw struct {
	Contract *Marketplace // Generic contract binding to access the raw methods on
}

// MarketplaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MarketplaceCallerRaw struct {
	Contract *MarketplaceCaller // Generic read-only contract binding to access the raw methods on
}

// MarketplaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MarketplaceTransactorRaw struct {
	Contract *MarketplaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMarketplace creates a new instance of Marketplace, bound to a specific deployed contract.
func NewMarketplace(address common.Address, backend bind.ContractBackend) (*Marketplace, error) {
	contract, err := bindMarketplace(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Marketplace{MarketplaceCaller: MarketplaceCaller{contract: contract}, MarketplaceTransactor: MarketplaceTransactor{contract: contract}, MarketplaceFilterer: MarketplaceFilterer{contract: contract}}, nil
}

// NewMarketplaceCaller creates a new read-only instance of Marketplace, bound to a specific deployed contract.
func NewMarketplaceCaller(address common.Address, caller bind.ContractCaller) (*MarketplaceCaller, error) {
	contract, err := bindMarketplace(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MarketplaceCaller{contract: contract}, nil
}

// NewMarketplaceTransactor creates a new write-only instance of Marketplace, bound to a specific deployed contract.
func NewMarketplaceTransactor(address common.Address, transactor bind.ContractTransactor) (*MarketplaceTransactor, error) {
	contract, err := bindMarketplace(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MarketplaceTransactor{contract: contract}, nil
}

// NewMarketplaceFilterer creates a new log filterer instance of Marketplace, bound to a specific deployed contract.
func NewMarketplaceFilterer(address common.Address, filterer bind.ContractFilterer) (*MarketplaceFilterer, error) {
	contract, err := bindMarketplace(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MarketplaceFilterer{contract: contract}, nil
}

// bindMarketplace binds a generic wrapper to an already deployed contract.
func bindMarketplace(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MarketplaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Marketplace *MarketplaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Marketplace.Contract.MarketplaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Marketplace *MarketplaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketplace.Contract.MarketplaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Marketplace *MarketplaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Marketplace.Contract.MarketplaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Marketplace *MarketplaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Marketplace.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Marketplace *MarketplaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketplace.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Marketplace *MarketplaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Marketplace.Contract.contract.Transact(opts, method, params...)
}

// KINDAUCTION is a free data retrieval call binding the contract method 0x7234d8f2.
//
// Solidity: function KIND_AUCTION() view returns(uint8)
func (_Marketplace *MarketplaceCaller) KINDAUCTION(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "KIND_AUCTION")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// KINDAUCTION is a free data retrieval call binding the contract method 0x7234d8f2.
//
// Solidity: function KIND_AUCTION() view returns(uint8)
func (_Marketplace *MarketplaceSession) KINDAUCTION() (uint8, error) {
	return _Marketplace.Contract.KINDAUCTION(&_Marketplace.CallOpts)
}

// KINDAUCTION is a free data retrieval call binding the contract method 0x7234d8f2.
//
// Solidity: function KIND_AUCTION() view returns(uint8)
func (_Marketplace *MarketplaceCallerSession) KINDAUCTION() (uint8, error) {
	return _Marketplace.Contract.KINDAUCTION(&_Marketplace.CallOpts)
}

// KINDBUY is a free data retrieval call binding the contract method 0xe1784a02.
//
// Solidity: function KIND_BUY() view returns(uint8)
func (_Marketplace *MarketplaceCaller) KINDBUY(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "KIND_BUY")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// KINDBUY is a free data retrieval call binding the contract method 0xe1784a02.
//
// Solidity: function KIND_BUY() view returns(uint8)
func (_Marketplace *MarketplaceSession) KINDBUY() (uint8, error) {
	return _Marketplace.Contract.KINDBUY(&_Marketplace.CallOpts)
}

// KINDBUY is a free data retrieval call binding the contract method 0xe1784a02.
//
// Solidity: function KIND_BUY() view returns(uint8)
func (_Marketplace *MarketplaceCallerSession) KINDBUY() (uint8, error) {
	return _Marketplace.Contract.KINDBUY(&_Marketplace.CallOpts)
}

// KINDSELL is a free data retrieval call binding the contract method 0x25593ac2.
//
// Solidity: function KIND_SELL() view returns(uint8)
func (_Marketplace *MarketplaceCaller) KINDSELL(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "KIND_SELL")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// KINDSELL is a free data retrieval call binding the contract method 0x25593ac2.
//
// Solidity: function KIND_SELL() view returns(uint8)
func (_Marketplace *MarketplaceSession) KINDSELL() (uint8, error) {
	return _Marketplace.Contract.KINDSELL(&_Marketplace.CallOpts)
}

// KINDSELL is a free data retrieval call binding the contract method 0x25593ac2.
//
// Solidity: function KIND_SELL() view returns(uint8)
func (_Marketplace *MarketplaceCallerSession) KINDSELL() (uint8, error) {
	return _Marketplace.Contract.KINDSELL(&_Marketplace.CallOpts)
}

// OPACCEPTAUCTION is a free data retrieval call binding the contract method 0x7ae1ace0.
//
// Solidity: function OP_ACCEPT_AUCTION() view returns(uint8)
func (_Marketplace *MarketplaceCaller) OPACCEPTAUCTION(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "OP_ACCEPT_AUCTION")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// OPACCEPTAUCTION is a free data retrieval call binding the contract method 0x7ae1ace0.
//
// Solidity: function OP_ACCEPT_AUCTION() view returns(uint8)
func (_Marketplace *MarketplaceSession) OPACCEPTAUCTION() (uint8, error) {
	return _Marketplace.Contract.OPACCEPTAUCTION(&_Marketplace.CallOpts)
}

// OPACCEPTAUCTION is a free data retrieval call binding the contract method 0x7ae1ace0.
//
// Solidity: function OP_ACCEPT_AUCTION() view returns(uint8)
func (_Marketplace *MarketplaceCallerSession) OPACCEPTAUCTION() (uint8, error) {
	return _Marketplace.Contract.OPACCEPTAUCTION(&_Marketplace.CallOpts)
}

// OPACCEPTBUY is a free data retrieval call binding the contract method 0x11f0794c.
//
// Solidity: function OP_ACCEPT_BUY() view returns(uint8)
func (_Marketplace *MarketplaceCaller) OPACCEPTBUY(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "OP_ACCEPT_BUY")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// OPACCEPTBUY is a free data retrieval call binding the contract method 0x11f0794c.
//
// Solidity: function OP_ACCEPT_BUY() view returns(uint8)
func (_Marketplace *MarketplaceSession) OPACCEPTBUY() (uint8, error) {
	return _Marketplace.Contract.OPACCEPTBUY(&_Marketplace.CallOpts)
}

// OPACCEPTBUY is a free data retrieval call binding the contract method 0x11f0794c.
//
// Solidity: function OP_ACCEPT_BUY() view returns(uint8)
func (_Marketplace *MarketplaceCallerSession) OPACCEPTBUY() (uint8, error) {
	return _Marketplace.Contract.OPACCEPTBUY(&_Marketplace.CallOpts)
}

// OPBID is a free data retrieval call binding the contract method 0x81787a85.
//
// Solidity: function OP_BID() view returns(uint8)
func (_Marketplace *MarketplaceCaller) OPBID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "OP_BID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// OPBID is a free data retrieval call binding the contract method 0x81787a85.
//
// Solidity: function OP_BID() view returns(uint8)
func (_Marketplace *MarketplaceSession) OPBID() (uint8, error) {
	return _Marketplace.Contract.OPBID(&_Marketplace.CallOpts)
}

// OPBID is a free data retrieval call binding the contract method 0x81787a85.
//
// Solidity: function OP_BID() view returns(uint8)
func (_Marketplace *MarketplaceCallerSession) OPBID() (uint8, error) {
	return _Marketplace.Contract.OPBID(&_Marketplace.CallOpts)
}

// OPBUY is a free data retrieval call binding the contract method 0xeb374261.
//
// Solidity: function OP_BUY() view returns(uint8)
func (_Marketplace *MarketplaceCaller) OPBUY(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "OP_BUY")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// OPBUY is a free data retrieval call binding the contract method 0xeb374261.
//
// Solidity: function OP_BUY() view returns(uint8)
func (_Marketplace *MarketplaceSession) OPBUY() (uint8, error) {
	return _Marketplace.Contract.OPBUY(&_Marketplace.CallOpts)
}

// OPBUY is a free data retrieval call binding the contract method 0xeb374261.
//
// Solidity: function OP_BUY() view returns(uint8)
func (_Marketplace *MarketplaceCallerSession) OPBUY() (uint8, error) {
	return _Marketplace.Contract.OPBUY(&_Marketplace.CallOpts)
}

// OPCANCELBUY is a free data retrieval call binding the contract method 0x9e57feb5.
//
// Solidity: function OP_CANCEL_BUY() view returns(uint8)
func (_Marketplace *MarketplaceCaller) OPCANCELBUY(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "OP_CANCEL_BUY")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// OPCANCELBUY is a free data retrieval call binding the contract method 0x9e57feb5.
//
// Solidity: function OP_CANCEL_BUY() view returns(uint8)
func (_Marketplace *MarketplaceSession) OPCANCELBUY() (uint8, error) {
	return _Marketplace.Contract.OPCANCELBUY(&_Marketplace.CallOpts)
}

// OPCANCELBUY is a free data retrieval call binding the contract method 0x9e57feb5.
//
// Solidity: function OP_CANCEL_BUY() view returns(uint8)
func (_Marketplace *MarketplaceCallerSession) OPCANCELBUY() (uint8, error) {
	return _Marketplace.Contract.OPCANCELBUY(&_Marketplace.CallOpts)
}

// OPCOMPLETEAUCTION is a free data retrieval call binding the contract method 0x6acc65db.
//
// Solidity: function OP_COMPLETE_AUCTION() view returns(uint8)
func (_Marketplace *MarketplaceCaller) OPCOMPLETEAUCTION(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "OP_COMPLETE_AUCTION")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// OPCOMPLETEAUCTION is a free data retrieval call binding the contract method 0x6acc65db.
//
// Solidity: function OP_COMPLETE_AUCTION() view returns(uint8)
func (_Marketplace *MarketplaceSession) OPCOMPLETEAUCTION() (uint8, error) {
	return _Marketplace.Contract.OPCOMPLETEAUCTION(&_Marketplace.CallOpts)
}

// OPCOMPLETEAUCTION is a free data retrieval call binding the contract method 0x6acc65db.
//
// Solidity: function OP_COMPLETE_AUCTION() view returns(uint8)
func (_Marketplace *MarketplaceCallerSession) OPCOMPLETEAUCTION() (uint8, error) {
	return _Marketplace.Contract.OPCOMPLETEAUCTION(&_Marketplace.CallOpts)
}

// OPCOMPLETEBUY is a free data retrieval call binding the contract method 0xb50a2a55.
//
// Solidity: function OP_COMPLETE_BUY() view returns(uint8)
func (_Marketplace *MarketplaceCaller) OPCOMPLETEBUY(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "OP_COMPLETE_BUY")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// OPCOMPLETEBUY is a free data retrieval call binding the contract method 0xb50a2a55.
//
// Solidity: function OP_COMPLETE_BUY() view returns(uint8)
func (_Marketplace *MarketplaceSession) OPCOMPLETEBUY() (uint8, error) {
	return _Marketplace.Contract.OPCOMPLETEBUY(&_Marketplace.CallOpts)
}

// OPCOMPLETEBUY is a free data retrieval call binding the contract method 0xb50a2a55.
//
// Solidity: function OP_COMPLETE_BUY() view returns(uint8)
func (_Marketplace *MarketplaceCallerSession) OPCOMPLETEBUY() (uint8, error) {
	return _Marketplace.Contract.OPCOMPLETEBUY(&_Marketplace.CallOpts)
}

// OPCOMPLETESELL is a free data retrieval call binding the contract method 0x8f18439e.
//
// Solidity: function OP_COMPLETE_SELL() view returns(uint8)
func (_Marketplace *MarketplaceCaller) OPCOMPLETESELL(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "OP_COMPLETE_SELL")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// OPCOMPLETESELL is a free data retrieval call binding the contract method 0x8f18439e.
//
// Solidity: function OP_COMPLETE_SELL() view returns(uint8)
func (_Marketplace *MarketplaceSession) OPCOMPLETESELL() (uint8, error) {
	return _Marketplace.Contract.OPCOMPLETESELL(&_Marketplace.CallOpts)
}

// OPCOMPLETESELL is a free data retrieval call binding the contract method 0x8f18439e.
//
// Solidity: function OP_COMPLETE_SELL() view returns(uint8)
func (_Marketplace *MarketplaceCallerSession) OPCOMPLETESELL() (uint8, error) {
	return _Marketplace.Contract.OPCOMPLETESELL(&_Marketplace.CallOpts)
}

// OPMAX is a free data retrieval call binding the contract method 0xf0954160.
//
// Solidity: function OP_MAX() view returns(uint8)
func (_Marketplace *MarketplaceCaller) OPMAX(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "OP_MAX")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// OPMAX is a free data retrieval call binding the contract method 0xf0954160.
//
// Solidity: function OP_MAX() view returns(uint8)
func (_Marketplace *MarketplaceSession) OPMAX() (uint8, error) {
	return _Marketplace.Contract.OPMAX(&_Marketplace.CallOpts)
}

// OPMAX is a free data retrieval call binding the contract method 0xf0954160.
//
// Solidity: function OP_MAX() view returns(uint8)
func (_Marketplace *MarketplaceCallerSession) OPMAX() (uint8, error) {
	return _Marketplace.Contract.OPMAX(&_Marketplace.CallOpts)
}

// OPMIN is a free data retrieval call binding the contract method 0x90c2b10e.
//
// Solidity: function OP_MIN() view returns(uint8)
func (_Marketplace *MarketplaceCaller) OPMIN(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "OP_MIN")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// OPMIN is a free data retrieval call binding the contract method 0x90c2b10e.
//
// Solidity: function OP_MIN() view returns(uint8)
func (_Marketplace *MarketplaceSession) OPMIN() (uint8, error) {
	return _Marketplace.Contract.OPMIN(&_Marketplace.CallOpts)
}

// OPMIN is a free data retrieval call binding the contract method 0x90c2b10e.
//
// Solidity: function OP_MIN() view returns(uint8)
func (_Marketplace *MarketplaceCallerSession) OPMIN() (uint8, error) {
	return _Marketplace.Contract.OPMIN(&_Marketplace.CallOpts)
}

// OPREJECTBUY is a free data retrieval call binding the contract method 0x1bb03ca9.
//
// Solidity: function OP_REJECT_BUY() view returns(uint8)
func (_Marketplace *MarketplaceCaller) OPREJECTBUY(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "OP_REJECT_BUY")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// OPREJECTBUY is a free data retrieval call binding the contract method 0x1bb03ca9.
//
// Solidity: function OP_REJECT_BUY() view returns(uint8)
func (_Marketplace *MarketplaceSession) OPREJECTBUY() (uint8, error) {
	return _Marketplace.Contract.OPREJECTBUY(&_Marketplace.CallOpts)
}

// OPREJECTBUY is a free data retrieval call binding the contract method 0x1bb03ca9.
//
// Solidity: function OP_REJECT_BUY() view returns(uint8)
func (_Marketplace *MarketplaceCallerSession) OPREJECTBUY() (uint8, error) {
	return _Marketplace.Contract.OPREJECTBUY(&_Marketplace.CallOpts)
}

// RATEBASE is a free data retrieval call binding the contract method 0x0873c6ec.
//
// Solidity: function RATE_BASE() view returns(uint256)
func (_Marketplace *MarketplaceCaller) RATEBASE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "RATE_BASE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RATEBASE is a free data retrieval call binding the contract method 0x0873c6ec.
//
// Solidity: function RATE_BASE() view returns(uint256)
func (_Marketplace *MarketplaceSession) RATEBASE() (*big.Int, error) {
	return _Marketplace.Contract.RATEBASE(&_Marketplace.CallOpts)
}

// RATEBASE is a free data retrieval call binding the contract method 0x0873c6ec.
//
// Solidity: function RATE_BASE() view returns(uint256)
func (_Marketplace *MarketplaceCallerSession) RATEBASE() (*big.Int, error) {
	return _Marketplace.Contract.RATEBASE(&_Marketplace.CallOpts)
}

// STATUSCANCELLED is a free data retrieval call binding the contract method 0x5a4e5a15.
//
// Solidity: function STATUS_CANCELLED() view returns(uint8)
func (_Marketplace *MarketplaceCaller) STATUSCANCELLED(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "STATUS_CANCELLED")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// STATUSCANCELLED is a free data retrieval call binding the contract method 0x5a4e5a15.
//
// Solidity: function STATUS_CANCELLED() view returns(uint8)
func (_Marketplace *MarketplaceSession) STATUSCANCELLED() (uint8, error) {
	return _Marketplace.Contract.STATUSCANCELLED(&_Marketplace.CallOpts)
}

// STATUSCANCELLED is a free data retrieval call binding the contract method 0x5a4e5a15.
//
// Solidity: function STATUS_CANCELLED() view returns(uint8)
func (_Marketplace *MarketplaceCallerSession) STATUSCANCELLED() (uint8, error) {
	return _Marketplace.Contract.STATUSCANCELLED(&_Marketplace.CallOpts)
}

// STATUSDONE is a free data retrieval call binding the contract method 0x740db280.
//
// Solidity: function STATUS_DONE() view returns(uint8)
func (_Marketplace *MarketplaceCaller) STATUSDONE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "STATUS_DONE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// STATUSDONE is a free data retrieval call binding the contract method 0x740db280.
//
// Solidity: function STATUS_DONE() view returns(uint8)
func (_Marketplace *MarketplaceSession) STATUSDONE() (uint8, error) {
	return _Marketplace.Contract.STATUSDONE(&_Marketplace.CallOpts)
}

// STATUSDONE is a free data retrieval call binding the contract method 0x740db280.
//
// Solidity: function STATUS_DONE() view returns(uint8)
func (_Marketplace *MarketplaceCallerSession) STATUSDONE() (uint8, error) {
	return _Marketplace.Contract.STATUSDONE(&_Marketplace.CallOpts)
}

// STATUSOPEN is a free data retrieval call binding the contract method 0x24f8515b.
//
// Solidity: function STATUS_OPEN() view returns(uint8)
func (_Marketplace *MarketplaceCaller) STATUSOPEN(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "STATUS_OPEN")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// STATUSOPEN is a free data retrieval call binding the contract method 0x24f8515b.
//
// Solidity: function STATUS_OPEN() view returns(uint8)
func (_Marketplace *MarketplaceSession) STATUSOPEN() (uint8, error) {
	return _Marketplace.Contract.STATUSOPEN(&_Marketplace.CallOpts)
}

// STATUSOPEN is a free data retrieval call binding the contract method 0x24f8515b.
//
// Solidity: function STATUS_OPEN() view returns(uint8)
func (_Marketplace *MarketplaceCallerSession) STATUSOPEN() (uint8, error) {
	return _Marketplace.Contract.STATUSOPEN(&_Marketplace.CallOpts)
}

// TOKEN1155 is a free data retrieval call binding the contract method 0xf0d250ba.
//
// Solidity: function TOKEN_1155() view returns(uint8)
func (_Marketplace *MarketplaceCaller) TOKEN1155(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "TOKEN_1155")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TOKEN1155 is a free data retrieval call binding the contract method 0xf0d250ba.
//
// Solidity: function TOKEN_1155() view returns(uint8)
func (_Marketplace *MarketplaceSession) TOKEN1155() (uint8, error) {
	return _Marketplace.Contract.TOKEN1155(&_Marketplace.CallOpts)
}

// TOKEN1155 is a free data retrieval call binding the contract method 0xf0d250ba.
//
// Solidity: function TOKEN_1155() view returns(uint8)
func (_Marketplace *MarketplaceCallerSession) TOKEN1155() (uint8, error) {
	return _Marketplace.Contract.TOKEN1155(&_Marketplace.CallOpts)
}

// TOKEN721 is a free data retrieval call binding the contract method 0xc477be20.
//
// Solidity: function TOKEN_721() view returns(uint8)
func (_Marketplace *MarketplaceCaller) TOKEN721(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "TOKEN_721")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TOKEN721 is a free data retrieval call binding the contract method 0xc477be20.
//
// Solidity: function TOKEN_721() view returns(uint8)
func (_Marketplace *MarketplaceSession) TOKEN721() (uint8, error) {
	return _Marketplace.Contract.TOKEN721(&_Marketplace.CallOpts)
}

// TOKEN721 is a free data retrieval call binding the contract method 0xc477be20.
//
// Solidity: function TOKEN_721() view returns(uint8)
func (_Marketplace *MarketplaceCallerSession) TOKEN721() (uint8, error) {
	return _Marketplace.Contract.TOKEN721(&_Marketplace.CallOpts)
}

// TOKENMINT is a free data retrieval call binding the contract method 0x853ca41a.
//
// Solidity: function TOKEN_MINT() view returns(uint8)
func (_Marketplace *MarketplaceCaller) TOKENMINT(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "TOKEN_MINT")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TOKENMINT is a free data retrieval call binding the contract method 0x853ca41a.
//
// Solidity: function TOKEN_MINT() view returns(uint8)
func (_Marketplace *MarketplaceSession) TOKENMINT() (uint8, error) {
	return _Marketplace.Contract.TOKENMINT(&_Marketplace.CallOpts)
}

// TOKENMINT is a free data retrieval call binding the contract method 0x853ca41a.
//
// Solidity: function TOKEN_MINT() view returns(uint8)
func (_Marketplace *MarketplaceCallerSession) TOKENMINT() (uint8, error) {
	return _Marketplace.Contract.TOKENMINT(&_Marketplace.CallOpts)
}

// CouponSpent is a free data retrieval call binding the contract method 0x3ed9ffb7.
//
// Solidity: function couponSpent(uint256 ) view returns(bool)
func (_Marketplace *MarketplaceCaller) CouponSpent(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "couponSpent", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CouponSpent is a free data retrieval call binding the contract method 0x3ed9ffb7.
//
// Solidity: function couponSpent(uint256 ) view returns(bool)
func (_Marketplace *MarketplaceSession) CouponSpent(arg0 *big.Int) (bool, error) {
	return _Marketplace.Contract.CouponSpent(&_Marketplace.CallOpts, arg0)
}

// CouponSpent is a free data retrieval call binding the contract method 0x3ed9ffb7.
//
// Solidity: function couponSpent(uint256 ) view returns(bool)
func (_Marketplace *MarketplaceCallerSession) CouponSpent(arg0 *big.Int) (bool, error) {
	return _Marketplace.Contract.CouponSpent(&_Marketplace.CallOpts, arg0)
}

// HasInv is a free data retrieval call binding the contract method 0xf5116bc9.
//
// Solidity: function hasInv(uint256 id) view returns(bool)
func (_Marketplace *MarketplaceCaller) HasInv(opts *bind.CallOpts, id *big.Int) (bool, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "hasInv", id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasInv is a free data retrieval call binding the contract method 0xf5116bc9.
//
// Solidity: function hasInv(uint256 id) view returns(bool)
func (_Marketplace *MarketplaceSession) HasInv(id *big.Int) (bool, error) {
	return _Marketplace.Contract.HasInv(&_Marketplace.CallOpts, id)
}

// HasInv is a free data retrieval call binding the contract method 0xf5116bc9.
//
// Solidity: function hasInv(uint256 id) view returns(bool)
func (_Marketplace *MarketplaceCallerSession) HasInv(id *big.Int) (bool, error) {
	return _Marketplace.Contract.HasInv(&_Marketplace.CallOpts, id)
}

// HasSignedIntention is a free data retrieval call binding the contract method 0xac5e2cb1.
//
// Solidity: function hasSignedIntention(uint8 op) pure returns(bool)
func (_Marketplace *MarketplaceCaller) HasSignedIntention(opts *bind.CallOpts, op uint8) (bool, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "hasSignedIntention", op)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasSignedIntention is a free data retrieval call binding the contract method 0xac5e2cb1.
//
// Solidity: function hasSignedIntention(uint8 op) pure returns(bool)
func (_Marketplace *MarketplaceSession) HasSignedIntention(op uint8) (bool, error) {
	return _Marketplace.Contract.HasSignedIntention(&_Marketplace.CallOpts, op)
}

// HasSignedIntention is a free data retrieval call binding the contract method 0xac5e2cb1.
//
// Solidity: function hasSignedIntention(uint8 op) pure returns(bool)
func (_Marketplace *MarketplaceCallerSession) HasSignedIntention(op uint8) (bool, error) {
	return _Marketplace.Contract.HasSignedIntention(&_Marketplace.CallOpts, op)
}

// Inventories is a free data retrieval call binding the contract method 0xcd78ba01.
//
// Solidity: function inventories(uint256 ) view returns(address seller, address buyer, address currency, uint256 price, uint256 netPrice, uint256 deadline, uint8 kind, uint8 status)
func (_Marketplace *MarketplaceCaller) Inventories(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Seller   common.Address
	Buyer    common.Address
	Currency common.Address
	Price    *big.Int
	NetPrice *big.Int
	Deadline *big.Int
	Kind     uint8
	Status   uint8
}, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "inventories", arg0)

	outstruct := new(struct {
		Seller   common.Address
		Buyer    common.Address
		Currency common.Address
		Price    *big.Int
		NetPrice *big.Int
		Deadline *big.Int
		Kind     uint8
		Status   uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Seller = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Buyer = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Currency = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Price = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.NetPrice = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Deadline = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Kind = *abi.ConvertType(out[6], new(uint8)).(*uint8)
	outstruct.Status = *abi.ConvertType(out[7], new(uint8)).(*uint8)

	return *outstruct, err

}

// Inventories is a free data retrieval call binding the contract method 0xcd78ba01.
//
// Solidity: function inventories(uint256 ) view returns(address seller, address buyer, address currency, uint256 price, uint256 netPrice, uint256 deadline, uint8 kind, uint8 status)
func (_Marketplace *MarketplaceSession) Inventories(arg0 *big.Int) (struct {
	Seller   common.Address
	Buyer    common.Address
	Currency common.Address
	Price    *big.Int
	NetPrice *big.Int
	Deadline *big.Int
	Kind     uint8
	Status   uint8
}, error) {
	return _Marketplace.Contract.Inventories(&_Marketplace.CallOpts, arg0)
}

// Inventories is a free data retrieval call binding the contract method 0xcd78ba01.
//
// Solidity: function inventories(uint256 ) view returns(address seller, address buyer, address currency, uint256 price, uint256 netPrice, uint256 deadline, uint8 kind, uint8 status)
func (_Marketplace *MarketplaceCallerSession) Inventories(arg0 *big.Int) (struct {
	Seller   common.Address
	Buyer    common.Address
	Currency common.Address
	Price    *big.Int
	NetPrice *big.Int
	Deadline *big.Int
	Kind     uint8
	Status   uint8
}, error) {
	return _Marketplace.Contract.Inventories(&_Marketplace.CallOpts, arg0)
}

// InventoryTokenCounts is a free data retrieval call binding the contract method 0x5fd34298.
//
// Solidity: function inventoryTokenCounts(uint256 ) view returns(uint256)
func (_Marketplace *MarketplaceCaller) InventoryTokenCounts(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "inventoryTokenCounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InventoryTokenCounts is a free data retrieval call binding the contract method 0x5fd34298.
//
// Solidity: function inventoryTokenCounts(uint256 ) view returns(uint256)
func (_Marketplace *MarketplaceSession) InventoryTokenCounts(arg0 *big.Int) (*big.Int, error) {
	return _Marketplace.Contract.InventoryTokenCounts(&_Marketplace.CallOpts, arg0)
}

// InventoryTokenCounts is a free data retrieval call binding the contract method 0x5fd34298.
//
// Solidity: function inventoryTokenCounts(uint256 ) view returns(uint256)
func (_Marketplace *MarketplaceCallerSession) InventoryTokenCounts(arg0 *big.Int) (*big.Int, error) {
	return _Marketplace.Contract.InventoryTokenCounts(&_Marketplace.CallOpts, arg0)
}

// InventoryTokens is a free data retrieval call binding the contract method 0xb4533aad.
//
// Solidity: function inventoryTokens(uint256 , uint256 ) view returns(address token, uint256 tokenId, uint256 amount, uint8 kind, bytes mintData)
func (_Marketplace *MarketplaceCaller) InventoryTokens(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (struct {
	Token    common.Address
	TokenId  *big.Int
	Amount   *big.Int
	Kind     uint8
	MintData []byte
}, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "inventoryTokens", arg0, arg1)

	outstruct := new(struct {
		Token    common.Address
		TokenId  *big.Int
		Amount   *big.Int
		Kind     uint8
		MintData []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Token = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.TokenId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Amount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Kind = *abi.ConvertType(out[3], new(uint8)).(*uint8)
	outstruct.MintData = *abi.ConvertType(out[4], new([]byte)).(*[]byte)

	return *outstruct, err

}

// InventoryTokens is a free data retrieval call binding the contract method 0xb4533aad.
//
// Solidity: function inventoryTokens(uint256 , uint256 ) view returns(address token, uint256 tokenId, uint256 amount, uint8 kind, bytes mintData)
func (_Marketplace *MarketplaceSession) InventoryTokens(arg0 *big.Int, arg1 *big.Int) (struct {
	Token    common.Address
	TokenId  *big.Int
	Amount   *big.Int
	Kind     uint8
	MintData []byte
}, error) {
	return _Marketplace.Contract.InventoryTokens(&_Marketplace.CallOpts, arg0, arg1)
}

// InventoryTokens is a free data retrieval call binding the contract method 0xb4533aad.
//
// Solidity: function inventoryTokens(uint256 , uint256 ) view returns(address token, uint256 tokenId, uint256 amount, uint8 kind, bytes mintData)
func (_Marketplace *MarketplaceCallerSession) InventoryTokens(arg0 *big.Int, arg1 *big.Int) (struct {
	Token    common.Address
	TokenId  *big.Int
	Amount   *big.Int
	Kind     uint8
	MintData []byte
}, error) {
	return _Marketplace.Contract.InventoryTokens(&_Marketplace.CallOpts, arg0, arg1)
}

// IsAuction is a free data retrieval call binding the contract method 0x8704f2a3.
//
// Solidity: function isAuction(uint256 id) view returns(bool)
func (_Marketplace *MarketplaceCaller) IsAuction(opts *bind.CallOpts, id *big.Int) (bool, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "isAuction", id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAuction is a free data retrieval call binding the contract method 0x8704f2a3.
//
// Solidity: function isAuction(uint256 id) view returns(bool)
func (_Marketplace *MarketplaceSession) IsAuction(id *big.Int) (bool, error) {
	return _Marketplace.Contract.IsAuction(&_Marketplace.CallOpts, id)
}

// IsAuction is a free data retrieval call binding the contract method 0x8704f2a3.
//
// Solidity: function isAuction(uint256 id) view returns(bool)
func (_Marketplace *MarketplaceCallerSession) IsAuction(id *big.Int) (bool, error) {
	return _Marketplace.Contract.IsAuction(&_Marketplace.CallOpts, id)
}

// IsAuctionOpen is a free data retrieval call binding the contract method 0x0ad48628.
//
// Solidity: function isAuctionOpen(uint256 id) view returns(bool)
func (_Marketplace *MarketplaceCaller) IsAuctionOpen(opts *bind.CallOpts, id *big.Int) (bool, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "isAuctionOpen", id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAuctionOpen is a free data retrieval call binding the contract method 0x0ad48628.
//
// Solidity: function isAuctionOpen(uint256 id) view returns(bool)
func (_Marketplace *MarketplaceSession) IsAuctionOpen(id *big.Int) (bool, error) {
	return _Marketplace.Contract.IsAuctionOpen(&_Marketplace.CallOpts, id)
}

// IsAuctionOpen is a free data retrieval call binding the contract method 0x0ad48628.
//
// Solidity: function isAuctionOpen(uint256 id) view returns(bool)
func (_Marketplace *MarketplaceCallerSession) IsAuctionOpen(id *big.Int) (bool, error) {
	return _Marketplace.Contract.IsAuctionOpen(&_Marketplace.CallOpts, id)
}

// IsBundleApproved is a free data retrieval call binding the contract method 0xf4a33e0d.
//
// Solidity: function isBundleApproved(uint256 invId, address owner) view returns(bool)
func (_Marketplace *MarketplaceCaller) IsBundleApproved(opts *bind.CallOpts, invId *big.Int, owner common.Address) (bool, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "isBundleApproved", invId, owner)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBundleApproved is a free data retrieval call binding the contract method 0xf4a33e0d.
//
// Solidity: function isBundleApproved(uint256 invId, address owner) view returns(bool)
func (_Marketplace *MarketplaceSession) IsBundleApproved(invId *big.Int, owner common.Address) (bool, error) {
	return _Marketplace.Contract.IsBundleApproved(&_Marketplace.CallOpts, invId, owner)
}

// IsBundleApproved is a free data retrieval call binding the contract method 0xf4a33e0d.
//
// Solidity: function isBundleApproved(uint256 invId, address owner) view returns(bool)
func (_Marketplace *MarketplaceCallerSession) IsBundleApproved(invId *big.Int, owner common.Address) (bool, error) {
	return _Marketplace.Contract.IsBundleApproved(&_Marketplace.CallOpts, invId, owner)
}

// IsBuy is a free data retrieval call binding the contract method 0xa80d33fb.
//
// Solidity: function isBuy(uint256 id) view returns(bool)
func (_Marketplace *MarketplaceCaller) IsBuy(opts *bind.CallOpts, id *big.Int) (bool, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "isBuy", id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBuy is a free data retrieval call binding the contract method 0xa80d33fb.
//
// Solidity: function isBuy(uint256 id) view returns(bool)
func (_Marketplace *MarketplaceSession) IsBuy(id *big.Int) (bool, error) {
	return _Marketplace.Contract.IsBuy(&_Marketplace.CallOpts, id)
}

// IsBuy is a free data retrieval call binding the contract method 0xa80d33fb.
//
// Solidity: function isBuy(uint256 id) view returns(bool)
func (_Marketplace *MarketplaceCallerSession) IsBuy(id *big.Int) (bool, error) {
	return _Marketplace.Contract.IsBuy(&_Marketplace.CallOpts, id)
}

// IsBuyOpen is a free data retrieval call binding the contract method 0xbdf52b45.
//
// Solidity: function isBuyOpen(uint256 id) view returns(bool)
func (_Marketplace *MarketplaceCaller) IsBuyOpen(opts *bind.CallOpts, id *big.Int) (bool, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "isBuyOpen", id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBuyOpen is a free data retrieval call binding the contract method 0xbdf52b45.
//
// Solidity: function isBuyOpen(uint256 id) view returns(bool)
func (_Marketplace *MarketplaceSession) IsBuyOpen(id *big.Int) (bool, error) {
	return _Marketplace.Contract.IsBuyOpen(&_Marketplace.CallOpts, id)
}

// IsBuyOpen is a free data retrieval call binding the contract method 0xbdf52b45.
//
// Solidity: function isBuyOpen(uint256 id) view returns(bool)
func (_Marketplace *MarketplaceCallerSession) IsBuyOpen(id *big.Int) (bool, error) {
	return _Marketplace.Contract.IsBuyOpen(&_Marketplace.CallOpts, id)
}

// IsExpired is a free data retrieval call binding the contract method 0xd9548e53.
//
// Solidity: function isExpired(uint256 id) view returns(bool)
func (_Marketplace *MarketplaceCaller) IsExpired(opts *bind.CallOpts, id *big.Int) (bool, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "isExpired", id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExpired is a free data retrieval call binding the contract method 0xd9548e53.
//
// Solidity: function isExpired(uint256 id) view returns(bool)
func (_Marketplace *MarketplaceSession) IsExpired(id *big.Int) (bool, error) {
	return _Marketplace.Contract.IsExpired(&_Marketplace.CallOpts, id)
}

// IsExpired is a free data retrieval call binding the contract method 0xd9548e53.
//
// Solidity: function isExpired(uint256 id) view returns(bool)
func (_Marketplace *MarketplaceCallerSession) IsExpired(id *big.Int) (bool, error) {
	return _Marketplace.Contract.IsExpired(&_Marketplace.CallOpts, id)
}

// IsSell is a free data retrieval call binding the contract method 0x1b01e72c.
//
// Solidity: function isSell(uint256 id) view returns(bool)
func (_Marketplace *MarketplaceCaller) IsSell(opts *bind.CallOpts, id *big.Int) (bool, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "isSell", id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSell is a free data retrieval call binding the contract method 0x1b01e72c.
//
// Solidity: function isSell(uint256 id) view returns(bool)
func (_Marketplace *MarketplaceSession) IsSell(id *big.Int) (bool, error) {
	return _Marketplace.Contract.IsSell(&_Marketplace.CallOpts, id)
}

// IsSell is a free data retrieval call binding the contract method 0x1b01e72c.
//
// Solidity: function isSell(uint256 id) view returns(bool)
func (_Marketplace *MarketplaceCallerSession) IsSell(id *big.Int) (bool, error) {
	return _Marketplace.Contract.IsSell(&_Marketplace.CallOpts, id)
}

// IsSignatureValid is a free data retrieval call binding the contract method 0x781dc70a.
//
// Solidity: function isSignatureValid(bytes signature, bytes32 hash, address signer) pure returns(bool)
func (_Marketplace *MarketplaceCaller) IsSignatureValid(opts *bind.CallOpts, signature []byte, hash [32]byte, signer common.Address) (bool, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "isSignatureValid", signature, hash, signer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSignatureValid is a free data retrieval call binding the contract method 0x781dc70a.
//
// Solidity: function isSignatureValid(bytes signature, bytes32 hash, address signer) pure returns(bool)
func (_Marketplace *MarketplaceSession) IsSignatureValid(signature []byte, hash [32]byte, signer common.Address) (bool, error) {
	return _Marketplace.Contract.IsSignatureValid(&_Marketplace.CallOpts, signature, hash, signer)
}

// IsSignatureValid is a free data retrieval call binding the contract method 0x781dc70a.
//
// Solidity: function isSignatureValid(bytes signature, bytes32 hash, address signer) pure returns(bool)
func (_Marketplace *MarketplaceCallerSession) IsSignatureValid(signature []byte, hash [32]byte, signer common.Address) (bool, error) {
	return _Marketplace.Contract.IsSignatureValid(&_Marketplace.CallOpts, signature, hash, signer)
}

// IsStatusOpen is a free data retrieval call binding the contract method 0xee98ce91.
//
// Solidity: function isStatusOpen(uint256 id) view returns(bool)
func (_Marketplace *MarketplaceCaller) IsStatusOpen(opts *bind.CallOpts, id *big.Int) (bool, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "isStatusOpen", id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStatusOpen is a free data retrieval call binding the contract method 0xee98ce91.
//
// Solidity: function isStatusOpen(uint256 id) view returns(bool)
func (_Marketplace *MarketplaceSession) IsStatusOpen(id *big.Int) (bool, error) {
	return _Marketplace.Contract.IsStatusOpen(&_Marketplace.CallOpts, id)
}

// IsStatusOpen is a free data retrieval call binding the contract method 0xee98ce91.
//
// Solidity: function isStatusOpen(uint256 id) view returns(bool)
func (_Marketplace *MarketplaceCallerSession) IsStatusOpen(id *big.Int) (bool, error) {
	return _Marketplace.Contract.IsStatusOpen(&_Marketplace.CallOpts, id)
}

// MarketSigners is a free data retrieval call binding the contract method 0x2bcd27df.
//
// Solidity: function marketSigners(address ) view returns(bool)
func (_Marketplace *MarketplaceCaller) MarketSigners(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "marketSigners", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MarketSigners is a free data retrieval call binding the contract method 0x2bcd27df.
//
// Solidity: function marketSigners(address ) view returns(bool)
func (_Marketplace *MarketplaceSession) MarketSigners(arg0 common.Address) (bool, error) {
	return _Marketplace.Contract.MarketSigners(&_Marketplace.CallOpts, arg0)
}

// MarketSigners is a free data retrieval call binding the contract method 0x2bcd27df.
//
// Solidity: function marketSigners(address ) view returns(bool)
func (_Marketplace *MarketplaceCallerSession) MarketSigners(arg0 common.Address) (bool, error) {
	return _Marketplace.Contract.MarketSigners(&_Marketplace.CallOpts, arg0)
}

// MinAuctionDuration is a free data retrieval call binding the contract method 0x54134876.
//
// Solidity: function minAuctionDuration() view returns(uint256)
func (_Marketplace *MarketplaceCaller) MinAuctionDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "minAuctionDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinAuctionDuration is a free data retrieval call binding the contract method 0x54134876.
//
// Solidity: function minAuctionDuration() view returns(uint256)
func (_Marketplace *MarketplaceSession) MinAuctionDuration() (*big.Int, error) {
	return _Marketplace.Contract.MinAuctionDuration(&_Marketplace.CallOpts)
}

// MinAuctionDuration is a free data retrieval call binding the contract method 0x54134876.
//
// Solidity: function minAuctionDuration() view returns(uint256)
func (_Marketplace *MarketplaceCallerSession) MinAuctionDuration() (*big.Int, error) {
	return _Marketplace.Contract.MinAuctionDuration(&_Marketplace.CallOpts)
}

// MinAuctionIncrement is a free data retrieval call binding the contract method 0x708d4d35.
//
// Solidity: function minAuctionIncrement() view returns(uint256)
func (_Marketplace *MarketplaceCaller) MinAuctionIncrement(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "minAuctionIncrement")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinAuctionIncrement is a free data retrieval call binding the contract method 0x708d4d35.
//
// Solidity: function minAuctionIncrement() view returns(uint256)
func (_Marketplace *MarketplaceSession) MinAuctionIncrement() (*big.Int, error) {
	return _Marketplace.Contract.MinAuctionIncrement(&_Marketplace.CallOpts)
}

// MinAuctionIncrement is a free data retrieval call binding the contract method 0x708d4d35.
//
// Solidity: function minAuctionIncrement() view returns(uint256)
func (_Marketplace *MarketplaceCallerSession) MinAuctionIncrement() (*big.Int, error) {
	return _Marketplace.Contract.MinAuctionIncrement(&_Marketplace.CallOpts)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address operator, address from, uint256[] ids, uint256[] values, bytes data) view returns(bytes4)
func (_Marketplace *MarketplaceCaller) OnERC1155BatchReceived(opts *bind.CallOpts, operator common.Address, from common.Address, ids []*big.Int, values []*big.Int, data []byte) ([4]byte, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "onERC1155BatchReceived", operator, from, ids, values, data)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address operator, address from, uint256[] ids, uint256[] values, bytes data) view returns(bytes4)
func (_Marketplace *MarketplaceSession) OnERC1155BatchReceived(operator common.Address, from common.Address, ids []*big.Int, values []*big.Int, data []byte) ([4]byte, error) {
	return _Marketplace.Contract.OnERC1155BatchReceived(&_Marketplace.CallOpts, operator, from, ids, values, data)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address operator, address from, uint256[] ids, uint256[] values, bytes data) view returns(bytes4)
func (_Marketplace *MarketplaceCallerSession) OnERC1155BatchReceived(operator common.Address, from common.Address, ids []*big.Int, values []*big.Int, data []byte) ([4]byte, error) {
	return _Marketplace.Contract.OnERC1155BatchReceived(&_Marketplace.CallOpts, operator, from, ids, values, data)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address operator, address from, uint256 id, uint256 value, bytes data) view returns(bytes4)
func (_Marketplace *MarketplaceCaller) OnERC1155Received(opts *bind.CallOpts, operator common.Address, from common.Address, id *big.Int, value *big.Int, data []byte) ([4]byte, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "onERC1155Received", operator, from, id, value, data)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address operator, address from, uint256 id, uint256 value, bytes data) view returns(bytes4)
func (_Marketplace *MarketplaceSession) OnERC1155Received(operator common.Address, from common.Address, id *big.Int, value *big.Int, data []byte) ([4]byte, error) {
	return _Marketplace.Contract.OnERC1155Received(&_Marketplace.CallOpts, operator, from, id, value, data)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address operator, address from, uint256 id, uint256 value, bytes data) view returns(bytes4)
func (_Marketplace *MarketplaceCallerSession) OnERC1155Received(operator common.Address, from common.Address, id *big.Int, value *big.Int, data []byte) ([4]byte, error) {
	return _Marketplace.Contract.OnERC1155Received(&_Marketplace.CallOpts, operator, from, id, value, data)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) view returns(bytes4)
func (_Marketplace *MarketplaceCaller) OnERC721Received(opts *bind.CallOpts, operator common.Address, from common.Address, tokenId *big.Int, data []byte) ([4]byte, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "onERC721Received", operator, from, tokenId, data)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) view returns(bytes4)
func (_Marketplace *MarketplaceSession) OnERC721Received(operator common.Address, from common.Address, tokenId *big.Int, data []byte) ([4]byte, error) {
	return _Marketplace.Contract.OnERC721Received(&_Marketplace.CallOpts, operator, from, tokenId, data)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) view returns(bytes4)
func (_Marketplace *MarketplaceCallerSession) OnERC721Received(operator common.Address, from common.Address, tokenId *big.Int, data []byte) ([4]byte, error) {
	return _Marketplace.Contract.OnERC721Received(&_Marketplace.CallOpts, operator, from, tokenId, data)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Marketplace *MarketplaceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Marketplace *MarketplaceSession) Owner() (common.Address, error) {
	return _Marketplace.Contract.Owner(&_Marketplace.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Marketplace *MarketplaceCallerSession) Owner() (common.Address, error) {
	return _Marketplace.Contract.Owner(&_Marketplace.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Marketplace *MarketplaceCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Marketplace *MarketplaceSession) Paused() (bool, error) {
	return _Marketplace.Contract.Paused(&_Marketplace.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Marketplace *MarketplaceCallerSession) Paused() (bool, error) {
	return _Marketplace.Contract.Paused(&_Marketplace.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_Marketplace *MarketplaceCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_Marketplace *MarketplaceSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Marketplace.Contract.SupportsInterface(&_Marketplace.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_Marketplace *MarketplaceCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Marketplace.Contract.SupportsInterface(&_Marketplace.CallOpts, interfaceId)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_Marketplace *MarketplaceCaller) Weth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "weth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_Marketplace *MarketplaceSession) Weth() (common.Address, error) {
	return _Marketplace.Contract.Weth(&_Marketplace.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_Marketplace *MarketplaceCallerSession) Weth() (common.Address, error) {
	return _Marketplace.Contract.Weth(&_Marketplace.CallOpts)
}

// CancelBuys is a paid mutator transaction binding the contract method 0xc1c30e80.
//
// Solidity: function cancelBuys(uint256[] ids) returns()
func (_Marketplace *MarketplaceTransactor) CancelBuys(opts *bind.TransactOpts, ids []*big.Int) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "cancelBuys", ids)
}

// CancelBuys is a paid mutator transaction binding the contract method 0xc1c30e80.
//
// Solidity: function cancelBuys(uint256[] ids) returns()
func (_Marketplace *MarketplaceSession) CancelBuys(ids []*big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.CancelBuys(&_Marketplace.TransactOpts, ids)
}

// CancelBuys is a paid mutator transaction binding the contract method 0xc1c30e80.
//
// Solidity: function cancelBuys(uint256[] ids) returns()
func (_Marketplace *MarketplaceTransactorSession) CancelBuys(ids []*big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.CancelBuys(&_Marketplace.TransactOpts, ids)
}

// EmergencyCancelAuction is a paid mutator transaction binding the contract method 0xe7d4a999.
//
// Solidity: function emergencyCancelAuction(uint256 id, bool noBundle) returns()
func (_Marketplace *MarketplaceTransactor) EmergencyCancelAuction(opts *bind.TransactOpts, id *big.Int, noBundle bool) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "emergencyCancelAuction", id, noBundle)
}

// EmergencyCancelAuction is a paid mutator transaction binding the contract method 0xe7d4a999.
//
// Solidity: function emergencyCancelAuction(uint256 id, bool noBundle) returns()
func (_Marketplace *MarketplaceSession) EmergencyCancelAuction(id *big.Int, noBundle bool) (*types.Transaction, error) {
	return _Marketplace.Contract.EmergencyCancelAuction(&_Marketplace.TransactOpts, id, noBundle)
}

// EmergencyCancelAuction is a paid mutator transaction binding the contract method 0xe7d4a999.
//
// Solidity: function emergencyCancelAuction(uint256 id, bool noBundle) returns()
func (_Marketplace *MarketplaceTransactorSession) EmergencyCancelAuction(id *big.Int, noBundle bool) (*types.Transaction, error) {
	return _Marketplace.Contract.EmergencyCancelAuction(&_Marketplace.TransactOpts, id, noBundle)
}

// InCaseMoneyGetsStuck is a paid mutator transaction binding the contract method 0x80bc688f.
//
// Solidity: function inCaseMoneyGetsStuck(address to, address currency, uint256 amount) returns()
func (_Marketplace *MarketplaceTransactor) InCaseMoneyGetsStuck(opts *bind.TransactOpts, to common.Address, currency common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "inCaseMoneyGetsStuck", to, currency, amount)
}

// InCaseMoneyGetsStuck is a paid mutator transaction binding the contract method 0x80bc688f.
//
// Solidity: function inCaseMoneyGetsStuck(address to, address currency, uint256 amount) returns()
func (_Marketplace *MarketplaceSession) InCaseMoneyGetsStuck(to common.Address, currency common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.InCaseMoneyGetsStuck(&_Marketplace.TransactOpts, to, currency, amount)
}

// InCaseMoneyGetsStuck is a paid mutator transaction binding the contract method 0x80bc688f.
//
// Solidity: function inCaseMoneyGetsStuck(address to, address currency, uint256 amount) returns()
func (_Marketplace *MarketplaceTransactorSession) InCaseMoneyGetsStuck(to common.Address, currency common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.InCaseMoneyGetsStuck(&_Marketplace.TransactOpts, to, currency, amount)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Marketplace *MarketplaceTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Marketplace *MarketplaceSession) Pause() (*types.Transaction, error) {
	return _Marketplace.Contract.Pause(&_Marketplace.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Marketplace *MarketplaceTransactorSession) Pause() (*types.Transaction, error) {
	return _Marketplace.Contract.Pause(&_Marketplace.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Marketplace *MarketplaceTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Marketplace *MarketplaceSession) RenounceOwnership() (*types.Transaction, error) {
	return _Marketplace.Contract.RenounceOwnership(&_Marketplace.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Marketplace *MarketplaceTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Marketplace.Contract.RenounceOwnership(&_Marketplace.TransactOpts)
}

// Run is a paid mutator transaction binding the contract method 0xba847759.
//
// Solidity: function run((address,(address,uint256,uint256,uint8,bytes)[],address,uint256,uint256,bytes32,uint8) intent, (bytes32,address,uint256,bytes32,uint256,uint8,address,address,uint256,uint256,(uint256[],uint256,uint256,uint256,address,address),(address,uint256,uint256,uint8,bytes)[],uint256) detail, bytes sigIntent, bytes sigDetail) payable returns()
func (_Marketplace *MarketplaceTransactor) Run(opts *bind.TransactOpts, intent MarketNGIntention, detail MarketNGDetail, sigIntent []byte, sigDetail []byte) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "run", intent, detail, sigIntent, sigDetail)
}

// Run is a paid mutator transaction binding the contract method 0xba847759.
//
// Solidity: function run((address,(address,uint256,uint256,uint8,bytes)[],address,uint256,uint256,bytes32,uint8) intent, (bytes32,address,uint256,bytes32,uint256,uint8,address,address,uint256,uint256,(uint256[],uint256,uint256,uint256,address,address),(address,uint256,uint256,uint8,bytes)[],uint256) detail, bytes sigIntent, bytes sigDetail) payable returns()
func (_Marketplace *MarketplaceSession) Run(intent MarketNGIntention, detail MarketNGDetail, sigIntent []byte, sigDetail []byte) (*types.Transaction, error) {
	return _Marketplace.Contract.Run(&_Marketplace.TransactOpts, intent, detail, sigIntent, sigDetail)
}

// Run is a paid mutator transaction binding the contract method 0xba847759.
//
// Solidity: function run((address,(address,uint256,uint256,uint8,bytes)[],address,uint256,uint256,bytes32,uint8) intent, (bytes32,address,uint256,bytes32,uint256,uint8,address,address,uint256,uint256,(uint256[],uint256,uint256,uint256,address,address),(address,uint256,uint256,uint8,bytes)[],uint256) detail, bytes sigIntent, bytes sigDetail) payable returns()
func (_Marketplace *MarketplaceTransactorSession) Run(intent MarketNGIntention, detail MarketNGDetail, sigIntent []byte, sigDetail []byte) (*types.Transaction, error) {
	return _Marketplace.Contract.Run(&_Marketplace.TransactOpts, intent, detail, sigIntent, sigDetail)
}

// Send is a paid mutator transaction binding the contract method 0xafd76a0b.
//
// Solidity: function send(address to, (address,uint256)[] tokens) returns()
func (_Marketplace *MarketplaceTransactor) Send(opts *bind.TransactOpts, to common.Address, tokens []MarketNGPair721) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "send", to, tokens)
}

// Send is a paid mutator transaction binding the contract method 0xafd76a0b.
//
// Solidity: function send(address to, (address,uint256)[] tokens) returns()
func (_Marketplace *MarketplaceSession) Send(to common.Address, tokens []MarketNGPair721) (*types.Transaction, error) {
	return _Marketplace.Contract.Send(&_Marketplace.TransactOpts, to, tokens)
}

// Send is a paid mutator transaction binding the contract method 0xafd76a0b.
//
// Solidity: function send(address to, (address,uint256)[] tokens) returns()
func (_Marketplace *MarketplaceTransactorSession) Send(to common.Address, tokens []MarketNGPair721) (*types.Transaction, error) {
	return _Marketplace.Contract.Send(&_Marketplace.TransactOpts, to, tokens)
}

// Swap is a paid mutator transaction binding the contract method 0xe91274f3.
//
// Solidity: function swap((bytes32,address,uint256,(address,uint256)[],(address,uint256)[]) req, bytes signature) returns()
func (_Marketplace *MarketplaceTransactor) Swap(opts *bind.TransactOpts, req MarketNGSwap, signature []byte) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "swap", req, signature)
}

// Swap is a paid mutator transaction binding the contract method 0xe91274f3.
//
// Solidity: function swap((bytes32,address,uint256,(address,uint256)[],(address,uint256)[]) req, bytes signature) returns()
func (_Marketplace *MarketplaceSession) Swap(req MarketNGSwap, signature []byte) (*types.Transaction, error) {
	return _Marketplace.Contract.Swap(&_Marketplace.TransactOpts, req, signature)
}

// Swap is a paid mutator transaction binding the contract method 0xe91274f3.
//
// Solidity: function swap((bytes32,address,uint256,(address,uint256)[],(address,uint256)[]) req, bytes signature) returns()
func (_Marketplace *MarketplaceTransactorSession) Swap(req MarketNGSwap, signature []byte) (*types.Transaction, error) {
	return _Marketplace.Contract.Swap(&_Marketplace.TransactOpts, req, signature)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Marketplace *MarketplaceTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Marketplace *MarketplaceSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.TransferOwnership(&_Marketplace.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Marketplace *MarketplaceTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.TransferOwnership(&_Marketplace.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Marketplace *MarketplaceTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Marketplace *MarketplaceSession) Unpause() (*types.Transaction, error) {
	return _Marketplace.Contract.Unpause(&_Marketplace.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Marketplace *MarketplaceTransactorSession) Unpause() (*types.Transaction, error) {
	return _Marketplace.Contract.Unpause(&_Marketplace.TransactOpts)
}

// UpdateSettings is a paid mutator transaction binding the contract method 0x015af8ee.
//
// Solidity: function updateSettings(uint256 minAuctionIncrement_, uint256 minAuctionDuration_) returns()
func (_Marketplace *MarketplaceTransactor) UpdateSettings(opts *bind.TransactOpts, minAuctionIncrement_ *big.Int, minAuctionDuration_ *big.Int) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "updateSettings", minAuctionIncrement_, minAuctionDuration_)
}

// UpdateSettings is a paid mutator transaction binding the contract method 0x015af8ee.
//
// Solidity: function updateSettings(uint256 minAuctionIncrement_, uint256 minAuctionDuration_) returns()
func (_Marketplace *MarketplaceSession) UpdateSettings(minAuctionIncrement_ *big.Int, minAuctionDuration_ *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.UpdateSettings(&_Marketplace.TransactOpts, minAuctionIncrement_, minAuctionDuration_)
}

// UpdateSettings is a paid mutator transaction binding the contract method 0x015af8ee.
//
// Solidity: function updateSettings(uint256 minAuctionIncrement_, uint256 minAuctionDuration_) returns()
func (_Marketplace *MarketplaceTransactorSession) UpdateSettings(minAuctionIncrement_ *big.Int, minAuctionDuration_ *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.UpdateSettings(&_Marketplace.TransactOpts, minAuctionIncrement_, minAuctionDuration_)
}

// UpdateSigner is a paid mutator transaction binding the contract method 0xf460590b.
//
// Solidity: function updateSigner(address addr, bool remove) returns()
func (_Marketplace *MarketplaceTransactor) UpdateSigner(opts *bind.TransactOpts, addr common.Address, remove bool) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "updateSigner", addr, remove)
}

// UpdateSigner is a paid mutator transaction binding the contract method 0xf460590b.
//
// Solidity: function updateSigner(address addr, bool remove) returns()
func (_Marketplace *MarketplaceSession) UpdateSigner(addr common.Address, remove bool) (*types.Transaction, error) {
	return _Marketplace.Contract.UpdateSigner(&_Marketplace.TransactOpts, addr, remove)
}

// UpdateSigner is a paid mutator transaction binding the contract method 0xf460590b.
//
// Solidity: function updateSigner(address addr, bool remove) returns()
func (_Marketplace *MarketplaceTransactorSession) UpdateSigner(addr common.Address, remove bool) (*types.Transaction, error) {
	return _Marketplace.Contract.UpdateSigner(&_Marketplace.TransactOpts, addr, remove)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Marketplace *MarketplaceTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketplace.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Marketplace *MarketplaceSession) Receive() (*types.Transaction, error) {
	return _Marketplace.Contract.Receive(&_Marketplace.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Marketplace *MarketplaceTransactorSession) Receive() (*types.Transaction, error) {
	return _Marketplace.Contract.Receive(&_Marketplace.TransactOpts)
}

// MarketplaceEvAuctionRefundIterator is returned from FilterEvAuctionRefund and is used to iterate over the raw logs and unpacked data for EvAuctionRefund events raised by the Marketplace contract.
type MarketplaceEvAuctionRefundIterator struct {
	Event *MarketplaceEvAuctionRefund // Event containing the contract specifics and raw log

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
func (it *MarketplaceEvAuctionRefundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceEvAuctionRefund)
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
		it.Event = new(MarketplaceEvAuctionRefund)
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
func (it *MarketplaceEvAuctionRefundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceEvAuctionRefundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceEvAuctionRefund represents a EvAuctionRefund event raised by the Marketplace contract.
type MarketplaceEvAuctionRefund struct {
	Id     *big.Int
	Bidder common.Address
	Refund *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEvAuctionRefund is a free log retrieval operation binding the contract event 0xa48bcf3362c21033397c03b92fb367d1962ba13b5bde0dfe491f9d88abb59e3f.
//
// Solidity: event EvAuctionRefund(uint256 indexed id, address bidder, uint256 refund)
func (_Marketplace *MarketplaceFilterer) FilterEvAuctionRefund(opts *bind.FilterOpts, id []*big.Int) (*MarketplaceEvAuctionRefundIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "EvAuctionRefund", idRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceEvAuctionRefundIterator{contract: _Marketplace.contract, event: "EvAuctionRefund", logs: logs, sub: sub}, nil
}

// WatchEvAuctionRefund is a free log subscription operation binding the contract event 0xa48bcf3362c21033397c03b92fb367d1962ba13b5bde0dfe491f9d88abb59e3f.
//
// Solidity: event EvAuctionRefund(uint256 indexed id, address bidder, uint256 refund)
func (_Marketplace *MarketplaceFilterer) WatchEvAuctionRefund(opts *bind.WatchOpts, sink chan<- *MarketplaceEvAuctionRefund, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "EvAuctionRefund", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceEvAuctionRefund)
				if err := _Marketplace.contract.UnpackLog(event, "EvAuctionRefund", log); err != nil {
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

// ParseEvAuctionRefund is a log parse operation binding the contract event 0xa48bcf3362c21033397c03b92fb367d1962ba13b5bde0dfe491f9d88abb59e3f.
//
// Solidity: event EvAuctionRefund(uint256 indexed id, address bidder, uint256 refund)
func (_Marketplace *MarketplaceFilterer) ParseEvAuctionRefund(log types.Log) (*MarketplaceEvAuctionRefund, error) {
	event := new(MarketplaceEvAuctionRefund)
	if err := _Marketplace.contract.UnpackLog(event, "EvAuctionRefund", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceEvCouponSpentIterator is returned from FilterEvCouponSpent and is used to iterate over the raw logs and unpacked data for EvCouponSpent events raised by the Marketplace contract.
type MarketplaceEvCouponSpentIterator struct {
	Event *MarketplaceEvCouponSpent // Event containing the contract specifics and raw log

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
func (it *MarketplaceEvCouponSpentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceEvCouponSpent)
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
		it.Event = new(MarketplaceEvCouponSpent)
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
func (it *MarketplaceEvCouponSpentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceEvCouponSpentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceEvCouponSpent represents a EvCouponSpent event raised by the Marketplace contract.
type MarketplaceEvCouponSpent struct {
	Id       *big.Int
	CouponId *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterEvCouponSpent is a free log retrieval operation binding the contract event 0x6aa71aa6b7aa6036ace4e4ceefbab7d89c4afb7fcfa1a3680499d7b37d32c82f.
//
// Solidity: event EvCouponSpent(uint256 indexed id, uint256 indexed couponId)
func (_Marketplace *MarketplaceFilterer) FilterEvCouponSpent(opts *bind.FilterOpts, id []*big.Int, couponId []*big.Int) (*MarketplaceEvCouponSpentIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var couponIdRule []interface{}
	for _, couponIdItem := range couponId {
		couponIdRule = append(couponIdRule, couponIdItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "EvCouponSpent", idRule, couponIdRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceEvCouponSpentIterator{contract: _Marketplace.contract, event: "EvCouponSpent", logs: logs, sub: sub}, nil
}

// WatchEvCouponSpent is a free log subscription operation binding the contract event 0x6aa71aa6b7aa6036ace4e4ceefbab7d89c4afb7fcfa1a3680499d7b37d32c82f.
//
// Solidity: event EvCouponSpent(uint256 indexed id, uint256 indexed couponId)
func (_Marketplace *MarketplaceFilterer) WatchEvCouponSpent(opts *bind.WatchOpts, sink chan<- *MarketplaceEvCouponSpent, id []*big.Int, couponId []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var couponIdRule []interface{}
	for _, couponIdItem := range couponId {
		couponIdRule = append(couponIdRule, couponIdItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "EvCouponSpent", idRule, couponIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceEvCouponSpent)
				if err := _Marketplace.contract.UnpackLog(event, "EvCouponSpent", log); err != nil {
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

// ParseEvCouponSpent is a log parse operation binding the contract event 0x6aa71aa6b7aa6036ace4e4ceefbab7d89c4afb7fcfa1a3680499d7b37d32c82f.
//
// Solidity: event EvCouponSpent(uint256 indexed id, uint256 indexed couponId)
func (_Marketplace *MarketplaceFilterer) ParseEvCouponSpent(log types.Log) (*MarketplaceEvCouponSpent, error) {
	event := new(MarketplaceEvCouponSpent)
	if err := _Marketplace.contract.UnpackLog(event, "EvCouponSpent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceEvInventoryUpdateIterator is returned from FilterEvInventoryUpdate and is used to iterate over the raw logs and unpacked data for EvInventoryUpdate events raised by the Marketplace contract.
type MarketplaceEvInventoryUpdateIterator struct {
	Event *MarketplaceEvInventoryUpdate // Event containing the contract specifics and raw log

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
func (it *MarketplaceEvInventoryUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceEvInventoryUpdate)
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
		it.Event = new(MarketplaceEvInventoryUpdate)
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
func (it *MarketplaceEvInventoryUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceEvInventoryUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceEvInventoryUpdate represents a EvInventoryUpdate event raised by the Marketplace contract.
type MarketplaceEvInventoryUpdate struct {
	Id        *big.Int
	Inventory MarketNGInventory
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEvInventoryUpdate is a free log retrieval operation binding the contract event 0x5beea7b3b87c573953fec05007114d17712e5775d364acc106d8da9e74849033.
//
// Solidity: event EvInventoryUpdate(uint256 indexed id, (address,address,address,uint256,uint256,uint256,uint8,uint8) inventory)
func (_Marketplace *MarketplaceFilterer) FilterEvInventoryUpdate(opts *bind.FilterOpts, id []*big.Int) (*MarketplaceEvInventoryUpdateIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "EvInventoryUpdate", idRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceEvInventoryUpdateIterator{contract: _Marketplace.contract, event: "EvInventoryUpdate", logs: logs, sub: sub}, nil
}

// WatchEvInventoryUpdate is a free log subscription operation binding the contract event 0x5beea7b3b87c573953fec05007114d17712e5775d364acc106d8da9e74849033.
//
// Solidity: event EvInventoryUpdate(uint256 indexed id, (address,address,address,uint256,uint256,uint256,uint8,uint8) inventory)
func (_Marketplace *MarketplaceFilterer) WatchEvInventoryUpdate(opts *bind.WatchOpts, sink chan<- *MarketplaceEvInventoryUpdate, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "EvInventoryUpdate", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceEvInventoryUpdate)
				if err := _Marketplace.contract.UnpackLog(event, "EvInventoryUpdate", log); err != nil {
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

// ParseEvInventoryUpdate is a log parse operation binding the contract event 0x5beea7b3b87c573953fec05007114d17712e5775d364acc106d8da9e74849033.
//
// Solidity: event EvInventoryUpdate(uint256 indexed id, (address,address,address,uint256,uint256,uint256,uint8,uint8) inventory)
func (_Marketplace *MarketplaceFilterer) ParseEvInventoryUpdate(log types.Log) (*MarketplaceEvInventoryUpdate, error) {
	event := new(MarketplaceEvInventoryUpdate)
	if err := _Marketplace.contract.UnpackLog(event, "EvInventoryUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceEvMarketSignerUpdateIterator is returned from FilterEvMarketSignerUpdate and is used to iterate over the raw logs and unpacked data for EvMarketSignerUpdate events raised by the Marketplace contract.
type MarketplaceEvMarketSignerUpdateIterator struct {
	Event *MarketplaceEvMarketSignerUpdate // Event containing the contract specifics and raw log

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
func (it *MarketplaceEvMarketSignerUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceEvMarketSignerUpdate)
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
		it.Event = new(MarketplaceEvMarketSignerUpdate)
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
func (it *MarketplaceEvMarketSignerUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceEvMarketSignerUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceEvMarketSignerUpdate represents a EvMarketSignerUpdate event raised by the Marketplace contract.
type MarketplaceEvMarketSignerUpdate struct {
	Addr      common.Address
	IsRemoval bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEvMarketSignerUpdate is a free log retrieval operation binding the contract event 0x90d56af4745c314d9b45054b55dc973378c558c1ad1554bccc70d39aa63a2cc5.
//
// Solidity: event EvMarketSignerUpdate(address addr, bool isRemoval)
func (_Marketplace *MarketplaceFilterer) FilterEvMarketSignerUpdate(opts *bind.FilterOpts) (*MarketplaceEvMarketSignerUpdateIterator, error) {

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "EvMarketSignerUpdate")
	if err != nil {
		return nil, err
	}
	return &MarketplaceEvMarketSignerUpdateIterator{contract: _Marketplace.contract, event: "EvMarketSignerUpdate", logs: logs, sub: sub}, nil
}

// WatchEvMarketSignerUpdate is a free log subscription operation binding the contract event 0x90d56af4745c314d9b45054b55dc973378c558c1ad1554bccc70d39aa63a2cc5.
//
// Solidity: event EvMarketSignerUpdate(address addr, bool isRemoval)
func (_Marketplace *MarketplaceFilterer) WatchEvMarketSignerUpdate(opts *bind.WatchOpts, sink chan<- *MarketplaceEvMarketSignerUpdate) (event.Subscription, error) {

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "EvMarketSignerUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceEvMarketSignerUpdate)
				if err := _Marketplace.contract.UnpackLog(event, "EvMarketSignerUpdate", log); err != nil {
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

// ParseEvMarketSignerUpdate is a log parse operation binding the contract event 0x90d56af4745c314d9b45054b55dc973378c558c1ad1554bccc70d39aa63a2cc5.
//
// Solidity: event EvMarketSignerUpdate(address addr, bool isRemoval)
func (_Marketplace *MarketplaceFilterer) ParseEvMarketSignerUpdate(log types.Log) (*MarketplaceEvMarketSignerUpdate, error) {
	event := new(MarketplaceEvMarketSignerUpdate)
	if err := _Marketplace.contract.UnpackLog(event, "EvMarketSignerUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceEvSettingsUpdatedIterator is returned from FilterEvSettingsUpdated and is used to iterate over the raw logs and unpacked data for EvSettingsUpdated events raised by the Marketplace contract.
type MarketplaceEvSettingsUpdatedIterator struct {
	Event *MarketplaceEvSettingsUpdated // Event containing the contract specifics and raw log

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
func (it *MarketplaceEvSettingsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceEvSettingsUpdated)
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
		it.Event = new(MarketplaceEvSettingsUpdated)
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
func (it *MarketplaceEvSettingsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceEvSettingsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceEvSettingsUpdated represents a EvSettingsUpdated event raised by the Marketplace contract.
type MarketplaceEvSettingsUpdated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEvSettingsUpdated is a free log retrieval operation binding the contract event 0x6c06ac894de6b71964f14d152b6674a4465a9b5d3f9cf9f216b8e7ea61467519.
//
// Solidity: event EvSettingsUpdated()
func (_Marketplace *MarketplaceFilterer) FilterEvSettingsUpdated(opts *bind.FilterOpts) (*MarketplaceEvSettingsUpdatedIterator, error) {

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "EvSettingsUpdated")
	if err != nil {
		return nil, err
	}
	return &MarketplaceEvSettingsUpdatedIterator{contract: _Marketplace.contract, event: "EvSettingsUpdated", logs: logs, sub: sub}, nil
}

// WatchEvSettingsUpdated is a free log subscription operation binding the contract event 0x6c06ac894de6b71964f14d152b6674a4465a9b5d3f9cf9f216b8e7ea61467519.
//
// Solidity: event EvSettingsUpdated()
func (_Marketplace *MarketplaceFilterer) WatchEvSettingsUpdated(opts *bind.WatchOpts, sink chan<- *MarketplaceEvSettingsUpdated) (event.Subscription, error) {

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "EvSettingsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceEvSettingsUpdated)
				if err := _Marketplace.contract.UnpackLog(event, "EvSettingsUpdated", log); err != nil {
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

// ParseEvSettingsUpdated is a log parse operation binding the contract event 0x6c06ac894de6b71964f14d152b6674a4465a9b5d3f9cf9f216b8e7ea61467519.
//
// Solidity: event EvSettingsUpdated()
func (_Marketplace *MarketplaceFilterer) ParseEvSettingsUpdated(log types.Log) (*MarketplaceEvSettingsUpdated, error) {
	event := new(MarketplaceEvSettingsUpdated)
	if err := _Marketplace.contract.UnpackLog(event, "EvSettingsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceEvSwappedIterator is returned from FilterEvSwapped and is used to iterate over the raw logs and unpacked data for EvSwapped events raised by the Marketplace contract.
type MarketplaceEvSwappedIterator struct {
	Event *MarketplaceEvSwapped // Event containing the contract specifics and raw log

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
func (it *MarketplaceEvSwappedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceEvSwapped)
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
		it.Event = new(MarketplaceEvSwapped)
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
func (it *MarketplaceEvSwappedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceEvSwappedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceEvSwapped represents a EvSwapped event raised by the Marketplace contract.
type MarketplaceEvSwapped struct {
	Req       MarketNGSwap
	Signature []byte
	Swapper   common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEvSwapped is a free log retrieval operation binding the contract event 0x92060d15ec9a14885865b744d2efb1fff3cab53411058a530f51d480288a864c.
//
// Solidity: event EvSwapped((bytes32,address,uint256,(address,uint256)[],(address,uint256)[]) req, bytes signature, address swapper)
func (_Marketplace *MarketplaceFilterer) FilterEvSwapped(opts *bind.FilterOpts) (*MarketplaceEvSwappedIterator, error) {

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "EvSwapped")
	if err != nil {
		return nil, err
	}
	return &MarketplaceEvSwappedIterator{contract: _Marketplace.contract, event: "EvSwapped", logs: logs, sub: sub}, nil
}

// WatchEvSwapped is a free log subscription operation binding the contract event 0x92060d15ec9a14885865b744d2efb1fff3cab53411058a530f51d480288a864c.
//
// Solidity: event EvSwapped((bytes32,address,uint256,(address,uint256)[],(address,uint256)[]) req, bytes signature, address swapper)
func (_Marketplace *MarketplaceFilterer) WatchEvSwapped(opts *bind.WatchOpts, sink chan<- *MarketplaceEvSwapped) (event.Subscription, error) {

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "EvSwapped")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceEvSwapped)
				if err := _Marketplace.contract.UnpackLog(event, "EvSwapped", log); err != nil {
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

// ParseEvSwapped is a log parse operation binding the contract event 0x92060d15ec9a14885865b744d2efb1fff3cab53411058a530f51d480288a864c.
//
// Solidity: event EvSwapped((bytes32,address,uint256,(address,uint256)[],(address,uint256)[]) req, bytes signature, address swapper)
func (_Marketplace *MarketplaceFilterer) ParseEvSwapped(log types.Log) (*MarketplaceEvSwapped, error) {
	event := new(MarketplaceEvSwapped)
	if err := _Marketplace.contract.UnpackLog(event, "EvSwapped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Marketplace contract.
type MarketplaceOwnershipTransferredIterator struct {
	Event *MarketplaceOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MarketplaceOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceOwnershipTransferred)
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
		it.Event = new(MarketplaceOwnershipTransferred)
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
func (it *MarketplaceOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceOwnershipTransferred represents a OwnershipTransferred event raised by the Marketplace contract.
type MarketplaceOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Marketplace *MarketplaceFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MarketplaceOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceOwnershipTransferredIterator{contract: _Marketplace.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Marketplace *MarketplaceFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MarketplaceOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceOwnershipTransferred)
				if err := _Marketplace.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Marketplace *MarketplaceFilterer) ParseOwnershipTransferred(log types.Log) (*MarketplaceOwnershipTransferred, error) {
	event := new(MarketplaceOwnershipTransferred)
	if err := _Marketplace.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplacePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Marketplace contract.
type MarketplacePausedIterator struct {
	Event *MarketplacePaused // Event containing the contract specifics and raw log

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
func (it *MarketplacePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplacePaused)
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
		it.Event = new(MarketplacePaused)
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
func (it *MarketplacePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplacePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplacePaused represents a Paused event raised by the Marketplace contract.
type MarketplacePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Marketplace *MarketplaceFilterer) FilterPaused(opts *bind.FilterOpts) (*MarketplacePausedIterator, error) {

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &MarketplacePausedIterator{contract: _Marketplace.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Marketplace *MarketplaceFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *MarketplacePaused) (event.Subscription, error) {

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplacePaused)
				if err := _Marketplace.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Marketplace *MarketplaceFilterer) ParsePaused(log types.Log) (*MarketplacePaused, error) {
	event := new(MarketplacePaused)
	if err := _Marketplace.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Marketplace contract.
type MarketplaceUnpausedIterator struct {
	Event *MarketplaceUnpaused // Event containing the contract specifics and raw log

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
func (it *MarketplaceUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceUnpaused)
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
		it.Event = new(MarketplaceUnpaused)
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
func (it *MarketplaceUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceUnpaused represents a Unpaused event raised by the Marketplace contract.
type MarketplaceUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Marketplace *MarketplaceFilterer) FilterUnpaused(opts *bind.FilterOpts) (*MarketplaceUnpausedIterator, error) {

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &MarketplaceUnpausedIterator{contract: _Marketplace.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Marketplace *MarketplaceFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *MarketplaceUnpaused) (event.Subscription, error) {

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceUnpaused)
				if err := _Marketplace.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Marketplace *MarketplaceFilterer) ParseUnpaused(log types.Log) (*MarketplaceUnpaused, error) {
	event := new(MarketplaceUnpaused)
	if err := _Marketplace.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
