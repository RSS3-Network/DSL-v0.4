// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package foundation

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

// NFTMarketReserveAuctionReserveAuction is an auto generated low-level Go binding around an user-defined struct.
type NFTMarketReserveAuctionReserveAuction struct {
	NftContract       common.Address
	TokenId           *big.Int
	Seller            common.Address
	Duration          *big.Int
	ExtensionDuration *big.Int
	EndTime           *big.Int
	Bidder            common.Address
	Amount            *big.Int
}

// FoundationMetaData contains all meta data concerning the Foundation contract.
var FoundationMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"treasury\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feth\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"royaltyRegistry\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"FETHNode_FETH_Address_Is_Not_A_Contract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FETHNode_Only_FETH_Can_Transfer_ETH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FoundationTreasuryNode_Address_Is_Not_A_Contract\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"buyPrice\",\"type\":\"uint256\"}],\"name\":\"NFTMarketBuyPrice_Cannot_Buy_At_Lower_Price\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NFTMarketBuyPrice_Cannot_Buy_Unset_Price\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NFTMarketBuyPrice_Cannot_Cancel_Unset_Price\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"NFTMarketBuyPrice_Only_Owner_Can_Cancel_Price\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"NFTMarketBuyPrice_Only_Owner_Can_Set_Price\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NFTMarketBuyPrice_Price_Already_Set\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NFTMarketBuyPrice_Price_Too_High\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"}],\"name\":\"NFTMarketBuyPrice_Seller_Mismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NFTMarketCore_Seller_Not_Found\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"curator\",\"type\":\"address\"}],\"name\":\"NFTMarketExhibition_Caller_Is_Not_Curator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NFTMarketExhibition_Can_Not_Add_Dupe_Seller\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NFTMarketExhibition_Curator_Automatically_Allowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NFTMarketExhibition_Exhibition_Does_Not_Exist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NFTMarketExhibition_Seller_Not_Allowed_In_Exhibition\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NFTMarketExhibition_Sellers_Required\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NFTMarketExhibition_Take_Rate_Too_High\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NFTMarketFees_Invalid_Protocol_Fee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NFTMarketFees_Royalty_Registry_Is_Not_A_Contract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NFTMarketOffer_Cannot_Be_Made_While_In_Auction\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"currentOfferAmount\",\"type\":\"uint256\"}],\"name\":\"NFTMarketOffer_Offer_Below_Min_Amount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"name\":\"NFTMarketOffer_Offer_Expired\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"currentOfferFrom\",\"type\":\"address\"}],\"name\":\"NFTMarketOffer_Offer_From_Does_Not_Match\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minOfferAmount\",\"type\":\"uint256\"}],\"name\":\"NFTMarketOffer_Offer_Must_Be_At_Least_Min_Amount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"}],\"name\":\"NFTMarketReserveAuction_Already_Listed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"}],\"name\":\"NFTMarketReserveAuction_Bid_Must_Be_At_Least_Min_Amount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"reservePrice\",\"type\":\"uint256\"}],\"name\":\"NFTMarketReserveAuction_Cannot_Bid_Lower_Than_Reserve_Price\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"name\":\"NFTMarketReserveAuction_Cannot_Bid_On_Ended_Auction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NFTMarketReserveAuction_Cannot_Bid_On_Nonexistent_Auction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NFTMarketReserveAuction_Cannot_Finalize_Already_Settled_Auction\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"name\":\"NFTMarketReserveAuction_Cannot_Finalize_Auction_In_Progress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NFTMarketReserveAuction_Cannot_Rebid_Over_Outstanding_Bid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NFTMarketReserveAuction_Cannot_Update_Auction_In_Progress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxDuration\",\"type\":\"uint256\"}],\"name\":\"NFTMarketReserveAuction_Exceeds_Max_Duration\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"extensionDuration\",\"type\":\"uint256\"}],\"name\":\"NFTMarketReserveAuction_Less_Than_Extension_Duration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NFTMarketReserveAuction_Must_Set_Non_Zero_Reserve_Price\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"}],\"name\":\"NFTMarketReserveAuction_Not_Matching_Seller\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"NFTMarketReserveAuction_Only_Owner_Can_Update_Auction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NFTMarketReserveAuction_Price_Already_Set\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NFTMarketReserveAuction_Too_Much_Value_Provided\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RouterContext_Not_A_Contract\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalFees\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"creatorRev\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sellerRev\",\"type\":\"uint256\"}],\"name\":\"BuyPriceAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"BuyPriceCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"BuyPriceInvalidated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"BuyPriceSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"buyReferrer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"buyReferrerFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"buyReferrerSellerFee\",\"type\":\"uint256\"}],\"name\":\"BuyReferralPaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"exhibitionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"curator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"takeRateInBasisPoints\",\"type\":\"uint16\"}],\"name\":\"ExhibitionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"exhibitionId\",\"type\":\"uint256\"}],\"name\":\"ExhibitionDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"exhibitionId\",\"type\":\"uint256\"}],\"name\":\"NftAddedToExhibition\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"exhibitionId\",\"type\":\"uint256\"}],\"name\":\"NftRemovedFromExhibition\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalFees\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"creatorRev\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sellerRev\",\"type\":\"uint256\"}],\"name\":\"OfferAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"OfferInvalidated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"}],\"name\":\"OfferMade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"name\":\"ReserveAuctionBidPlaced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"}],\"name\":\"ReserveAuctionCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"extensionDuration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reservePrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"}],\"name\":\"ReserveAuctionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalFees\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"creatorRev\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sellerRev\",\"type\":\"uint256\"}],\"name\":\"ReserveAuctionFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"}],\"name\":\"ReserveAuctionInvalidated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reservePrice\",\"type\":\"uint256\"}],\"name\":\"ReserveAuctionUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sellerReferrer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sellerReferrerFee\",\"type\":\"uint256\"}],\"name\":\"SellerReferralPaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"exhibitionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"sellers\",\"type\":\"address[]\"}],\"name\":\"SellersAddedToExhibition\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawalToFETH\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"offerFrom\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"}],\"name\":\"acceptOffer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"exhibitionId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"sellers\",\"type\":\"address[]\"}],\"name\":\"addSellersToExhibition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPrice\",\"type\":\"uint256\"}],\"name\":\"buy\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPrice\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"referrer\",\"type\":\"address\"}],\"name\":\"buyV2\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"cancelBuyPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"}],\"name\":\"cancelReserveAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint16\",\"name\":\"takeRateInBasisPoints\",\"type\":\"uint16\"},{\"internalType\":\"address[]\",\"name\":\"sellers\",\"type\":\"address[]\"}],\"name\":\"createExhibition\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"exhibitionId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reservePrice\",\"type\":\"uint256\"}],\"name\":\"createReserveAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reservePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exhibitionId\",\"type\":\"uint256\"}],\"name\":\"createReserveAuctionV2\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"exhibitionId\",\"type\":\"uint256\"}],\"name\":\"deleteExhibition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"}],\"name\":\"finalizeReserveAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getApprovedRouterAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getBuyPrice\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"exhibitionId\",\"type\":\"uint256\"}],\"name\":\"getExhibition\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"addresspayable\",\"name\":\"curator\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"takeRateInBasisPoints\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getExhibitionIdForNft\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"exhibitionId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"getFeesAndRecipients\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"creatorRev\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable[]\",\"name\":\"creatorRecipients\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"creatorShares\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"sellerRev\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"seller\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFethAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"fethAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFoundationTreasury\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"treasuryAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"}],\"name\":\"getMinBidAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"minimum\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getMinOfferAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"minimum\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getOffer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getOfferReferrer\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"referrer\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"}],\"name\":\"getReserveAuction\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"extensionDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structNFTMarketReserveAuction.ReserveAuction\",\"name\":\"auction\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"}],\"name\":\"getReserveAuctionBidReferrer\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"referrer\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getReserveAuctionIdFor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRoyaltyRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"registry\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getSellerOf\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"seller\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"internalGetImmutableRoyalties\",\"outputs\":[{\"internalType\":\"addresspayable[]\",\"name\":\"recipients\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"splitPerRecipientInBasisPoints\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"internalGetMutableRoyalties\",\"outputs\":[{\"internalType\":\"addresspayable[]\",\"name\":\"recipients\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"splitPerRecipientInBasisPoints\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"internalGetTokenCreator\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"creator\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"exhibitionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"}],\"name\":\"isAllowedSellerForExhibition\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"allowedSeller\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"referrer\",\"type\":\"address\"}],\"name\":\"makeOfferV2\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"}],\"name\":\"placeBid\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"referrer\",\"type\":\"address\"}],\"name\":\"placeBidV2\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"setBuyPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reservePrice\",\"type\":\"uint256\"}],\"name\":\"updateReserveAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// FoundationABI is the input ABI used to generate the binding from.
// Deprecated: Use FoundationMetaData.ABI instead.
var FoundationABI = FoundationMetaData.ABI

// Foundation is an auto generated Go binding around an Ethereum contract.
type Foundation struct {
	FoundationCaller     // Read-only binding to the contract
	FoundationTransactor // Write-only binding to the contract
	FoundationFilterer   // Log filterer for contract events
}

// FoundationCaller is an auto generated read-only Go binding around an Ethereum contract.
type FoundationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FoundationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FoundationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FoundationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FoundationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FoundationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FoundationSession struct {
	Contract     *Foundation       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FoundationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FoundationCallerSession struct {
	Contract *FoundationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// FoundationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FoundationTransactorSession struct {
	Contract     *FoundationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// FoundationRaw is an auto generated low-level Go binding around an Ethereum contract.
type FoundationRaw struct {
	Contract *Foundation // Generic contract binding to access the raw methods on
}

// FoundationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FoundationCallerRaw struct {
	Contract *FoundationCaller // Generic read-only contract binding to access the raw methods on
}

// FoundationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FoundationTransactorRaw struct {
	Contract *FoundationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFoundation creates a new instance of Foundation, bound to a specific deployed contract.
func NewFoundation(address common.Address, backend bind.ContractBackend) (*Foundation, error) {
	contract, err := bindFoundation(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Foundation{FoundationCaller: FoundationCaller{contract: contract}, FoundationTransactor: FoundationTransactor{contract: contract}, FoundationFilterer: FoundationFilterer{contract: contract}}, nil
}

// NewFoundationCaller creates a new read-only instance of Foundation, bound to a specific deployed contract.
func NewFoundationCaller(address common.Address, caller bind.ContractCaller) (*FoundationCaller, error) {
	contract, err := bindFoundation(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FoundationCaller{contract: contract}, nil
}

// NewFoundationTransactor creates a new write-only instance of Foundation, bound to a specific deployed contract.
func NewFoundationTransactor(address common.Address, transactor bind.ContractTransactor) (*FoundationTransactor, error) {
	contract, err := bindFoundation(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FoundationTransactor{contract: contract}, nil
}

// NewFoundationFilterer creates a new log filterer instance of Foundation, bound to a specific deployed contract.
func NewFoundationFilterer(address common.Address, filterer bind.ContractFilterer) (*FoundationFilterer, error) {
	contract, err := bindFoundation(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FoundationFilterer{contract: contract}, nil
}

// bindFoundation binds a generic wrapper to an already deployed contract.
func bindFoundation(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FoundationMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Foundation *FoundationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Foundation.Contract.FoundationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Foundation *FoundationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Foundation.Contract.FoundationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Foundation *FoundationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Foundation.Contract.FoundationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Foundation *FoundationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Foundation.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Foundation *FoundationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Foundation.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Foundation *FoundationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Foundation.Contract.contract.Transact(opts, method, params...)
}

// GetApprovedRouterAddress is a free data retrieval call binding the contract method 0x6a90a827.
//
// Solidity: function getApprovedRouterAddress() view returns(address router)
func (_Foundation *FoundationCaller) GetApprovedRouterAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Foundation.contract.Call(opts, &out, "getApprovedRouterAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApprovedRouterAddress is a free data retrieval call binding the contract method 0x6a90a827.
//
// Solidity: function getApprovedRouterAddress() view returns(address router)
func (_Foundation *FoundationSession) GetApprovedRouterAddress() (common.Address, error) {
	return _Foundation.Contract.GetApprovedRouterAddress(&_Foundation.CallOpts)
}

// GetApprovedRouterAddress is a free data retrieval call binding the contract method 0x6a90a827.
//
// Solidity: function getApprovedRouterAddress() view returns(address router)
func (_Foundation *FoundationCallerSession) GetApprovedRouterAddress() (common.Address, error) {
	return _Foundation.Contract.GetApprovedRouterAddress(&_Foundation.CallOpts)
}

// GetBuyPrice is a free data retrieval call binding the contract method 0x4635256e.
//
// Solidity: function getBuyPrice(address nftContract, uint256 tokenId) view returns(address seller, uint256 price)
func (_Foundation *FoundationCaller) GetBuyPrice(opts *bind.CallOpts, nftContract common.Address, tokenId *big.Int) (struct {
	Seller common.Address
	Price  *big.Int
}, error) {
	var out []interface{}
	err := _Foundation.contract.Call(opts, &out, "getBuyPrice", nftContract, tokenId)

	outstruct := new(struct {
		Seller common.Address
		Price  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Seller = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Price = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetBuyPrice is a free data retrieval call binding the contract method 0x4635256e.
//
// Solidity: function getBuyPrice(address nftContract, uint256 tokenId) view returns(address seller, uint256 price)
func (_Foundation *FoundationSession) GetBuyPrice(nftContract common.Address, tokenId *big.Int) (struct {
	Seller common.Address
	Price  *big.Int
}, error) {
	return _Foundation.Contract.GetBuyPrice(&_Foundation.CallOpts, nftContract, tokenId)
}

// GetBuyPrice is a free data retrieval call binding the contract method 0x4635256e.
//
// Solidity: function getBuyPrice(address nftContract, uint256 tokenId) view returns(address seller, uint256 price)
func (_Foundation *FoundationCallerSession) GetBuyPrice(nftContract common.Address, tokenId *big.Int) (struct {
	Seller common.Address
	Price  *big.Int
}, error) {
	return _Foundation.Contract.GetBuyPrice(&_Foundation.CallOpts, nftContract, tokenId)
}

// GetExhibition is a free data retrieval call binding the contract method 0x3c58e54d.
//
// Solidity: function getExhibition(uint256 exhibitionId) view returns(string name, address curator, uint16 takeRateInBasisPoints)
func (_Foundation *FoundationCaller) GetExhibition(opts *bind.CallOpts, exhibitionId *big.Int) (struct {
	Name                  string
	Curator               common.Address
	TakeRateInBasisPoints uint16
}, error) {
	var out []interface{}
	err := _Foundation.contract.Call(opts, &out, "getExhibition", exhibitionId)

	outstruct := new(struct {
		Name                  string
		Curator               common.Address
		TakeRateInBasisPoints uint16
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Curator = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.TakeRateInBasisPoints = *abi.ConvertType(out[2], new(uint16)).(*uint16)

	return *outstruct, err

}

// GetExhibition is a free data retrieval call binding the contract method 0x3c58e54d.
//
// Solidity: function getExhibition(uint256 exhibitionId) view returns(string name, address curator, uint16 takeRateInBasisPoints)
func (_Foundation *FoundationSession) GetExhibition(exhibitionId *big.Int) (struct {
	Name                  string
	Curator               common.Address
	TakeRateInBasisPoints uint16
}, error) {
	return _Foundation.Contract.GetExhibition(&_Foundation.CallOpts, exhibitionId)
}

// GetExhibition is a free data retrieval call binding the contract method 0x3c58e54d.
//
// Solidity: function getExhibition(uint256 exhibitionId) view returns(string name, address curator, uint16 takeRateInBasisPoints)
func (_Foundation *FoundationCallerSession) GetExhibition(exhibitionId *big.Int) (struct {
	Name                  string
	Curator               common.Address
	TakeRateInBasisPoints uint16
}, error) {
	return _Foundation.Contract.GetExhibition(&_Foundation.CallOpts, exhibitionId)
}

// GetExhibitionIdForNft is a free data retrieval call binding the contract method 0x442559a2.
//
// Solidity: function getExhibitionIdForNft(address nftContract, uint256 tokenId) view returns(uint256 exhibitionId)
func (_Foundation *FoundationCaller) GetExhibitionIdForNft(opts *bind.CallOpts, nftContract common.Address, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Foundation.contract.Call(opts, &out, "getExhibitionIdForNft", nftContract, tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetExhibitionIdForNft is a free data retrieval call binding the contract method 0x442559a2.
//
// Solidity: function getExhibitionIdForNft(address nftContract, uint256 tokenId) view returns(uint256 exhibitionId)
func (_Foundation *FoundationSession) GetExhibitionIdForNft(nftContract common.Address, tokenId *big.Int) (*big.Int, error) {
	return _Foundation.Contract.GetExhibitionIdForNft(&_Foundation.CallOpts, nftContract, tokenId)
}

// GetExhibitionIdForNft is a free data retrieval call binding the contract method 0x442559a2.
//
// Solidity: function getExhibitionIdForNft(address nftContract, uint256 tokenId) view returns(uint256 exhibitionId)
func (_Foundation *FoundationCallerSession) GetExhibitionIdForNft(nftContract common.Address, tokenId *big.Int) (*big.Int, error) {
	return _Foundation.Contract.GetExhibitionIdForNft(&_Foundation.CallOpts, nftContract, tokenId)
}

// GetFeesAndRecipients is a free data retrieval call binding the contract method 0xaf1e1de3.
//
// Solidity: function getFeesAndRecipients(address nftContract, uint256 tokenId, uint256 price) view returns(uint256 totalFees, uint256 creatorRev, address[] creatorRecipients, uint256[] creatorShares, uint256 sellerRev, address seller)
func (_Foundation *FoundationCaller) GetFeesAndRecipients(opts *bind.CallOpts, nftContract common.Address, tokenId *big.Int, price *big.Int) (struct {
	TotalFees         *big.Int
	CreatorRev        *big.Int
	CreatorRecipients []common.Address
	CreatorShares     []*big.Int
	SellerRev         *big.Int
	Seller            common.Address
}, error) {
	var out []interface{}
	err := _Foundation.contract.Call(opts, &out, "getFeesAndRecipients", nftContract, tokenId, price)

	outstruct := new(struct {
		TotalFees         *big.Int
		CreatorRev        *big.Int
		CreatorRecipients []common.Address
		CreatorShares     []*big.Int
		SellerRev         *big.Int
		Seller            common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalFees = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.CreatorRev = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.CreatorRecipients = *abi.ConvertType(out[2], new([]common.Address)).(*[]common.Address)
	outstruct.CreatorShares = *abi.ConvertType(out[3], new([]*big.Int)).(*[]*big.Int)
	outstruct.SellerRev = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Seller = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// GetFeesAndRecipients is a free data retrieval call binding the contract method 0xaf1e1de3.
//
// Solidity: function getFeesAndRecipients(address nftContract, uint256 tokenId, uint256 price) view returns(uint256 totalFees, uint256 creatorRev, address[] creatorRecipients, uint256[] creatorShares, uint256 sellerRev, address seller)
func (_Foundation *FoundationSession) GetFeesAndRecipients(nftContract common.Address, tokenId *big.Int, price *big.Int) (struct {
	TotalFees         *big.Int
	CreatorRev        *big.Int
	CreatorRecipients []common.Address
	CreatorShares     []*big.Int
	SellerRev         *big.Int
	Seller            common.Address
}, error) {
	return _Foundation.Contract.GetFeesAndRecipients(&_Foundation.CallOpts, nftContract, tokenId, price)
}

// GetFeesAndRecipients is a free data retrieval call binding the contract method 0xaf1e1de3.
//
// Solidity: function getFeesAndRecipients(address nftContract, uint256 tokenId, uint256 price) view returns(uint256 totalFees, uint256 creatorRev, address[] creatorRecipients, uint256[] creatorShares, uint256 sellerRev, address seller)
func (_Foundation *FoundationCallerSession) GetFeesAndRecipients(nftContract common.Address, tokenId *big.Int, price *big.Int) (struct {
	TotalFees         *big.Int
	CreatorRev        *big.Int
	CreatorRecipients []common.Address
	CreatorShares     []*big.Int
	SellerRev         *big.Int
	Seller            common.Address
}, error) {
	return _Foundation.Contract.GetFeesAndRecipients(&_Foundation.CallOpts, nftContract, tokenId, price)
}

// GetFethAddress is a free data retrieval call binding the contract method 0x895633ba.
//
// Solidity: function getFethAddress() view returns(address fethAddress)
func (_Foundation *FoundationCaller) GetFethAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Foundation.contract.Call(opts, &out, "getFethAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFethAddress is a free data retrieval call binding the contract method 0x895633ba.
//
// Solidity: function getFethAddress() view returns(address fethAddress)
func (_Foundation *FoundationSession) GetFethAddress() (common.Address, error) {
	return _Foundation.Contract.GetFethAddress(&_Foundation.CallOpts)
}

// GetFethAddress is a free data retrieval call binding the contract method 0x895633ba.
//
// Solidity: function getFethAddress() view returns(address fethAddress)
func (_Foundation *FoundationCallerSession) GetFethAddress() (common.Address, error) {
	return _Foundation.Contract.GetFethAddress(&_Foundation.CallOpts)
}

// GetFoundationTreasury is a free data retrieval call binding the contract method 0xf7a2da23.
//
// Solidity: function getFoundationTreasury() view returns(address treasuryAddress)
func (_Foundation *FoundationCaller) GetFoundationTreasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Foundation.contract.Call(opts, &out, "getFoundationTreasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFoundationTreasury is a free data retrieval call binding the contract method 0xf7a2da23.
//
// Solidity: function getFoundationTreasury() view returns(address treasuryAddress)
func (_Foundation *FoundationSession) GetFoundationTreasury() (common.Address, error) {
	return _Foundation.Contract.GetFoundationTreasury(&_Foundation.CallOpts)
}

// GetFoundationTreasury is a free data retrieval call binding the contract method 0xf7a2da23.
//
// Solidity: function getFoundationTreasury() view returns(address treasuryAddress)
func (_Foundation *FoundationCallerSession) GetFoundationTreasury() (common.Address, error) {
	return _Foundation.Contract.GetFoundationTreasury(&_Foundation.CallOpts)
}

// GetMinBidAmount is a free data retrieval call binding the contract method 0x47e35740.
//
// Solidity: function getMinBidAmount(uint256 auctionId) view returns(uint256 minimum)
func (_Foundation *FoundationCaller) GetMinBidAmount(opts *bind.CallOpts, auctionId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Foundation.contract.Call(opts, &out, "getMinBidAmount", auctionId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinBidAmount is a free data retrieval call binding the contract method 0x47e35740.
//
// Solidity: function getMinBidAmount(uint256 auctionId) view returns(uint256 minimum)
func (_Foundation *FoundationSession) GetMinBidAmount(auctionId *big.Int) (*big.Int, error) {
	return _Foundation.Contract.GetMinBidAmount(&_Foundation.CallOpts, auctionId)
}

// GetMinBidAmount is a free data retrieval call binding the contract method 0x47e35740.
//
// Solidity: function getMinBidAmount(uint256 auctionId) view returns(uint256 minimum)
func (_Foundation *FoundationCallerSession) GetMinBidAmount(auctionId *big.Int) (*big.Int, error) {
	return _Foundation.Contract.GetMinBidAmount(&_Foundation.CallOpts, auctionId)
}

// GetMinOfferAmount is a free data retrieval call binding the contract method 0xe5d1e723.
//
// Solidity: function getMinOfferAmount(address nftContract, uint256 tokenId) view returns(uint256 minimum)
func (_Foundation *FoundationCaller) GetMinOfferAmount(opts *bind.CallOpts, nftContract common.Address, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Foundation.contract.Call(opts, &out, "getMinOfferAmount", nftContract, tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinOfferAmount is a free data retrieval call binding the contract method 0xe5d1e723.
//
// Solidity: function getMinOfferAmount(address nftContract, uint256 tokenId) view returns(uint256 minimum)
func (_Foundation *FoundationSession) GetMinOfferAmount(nftContract common.Address, tokenId *big.Int) (*big.Int, error) {
	return _Foundation.Contract.GetMinOfferAmount(&_Foundation.CallOpts, nftContract, tokenId)
}

// GetMinOfferAmount is a free data retrieval call binding the contract method 0xe5d1e723.
//
// Solidity: function getMinOfferAmount(address nftContract, uint256 tokenId) view returns(uint256 minimum)
func (_Foundation *FoundationCallerSession) GetMinOfferAmount(nftContract common.Address, tokenId *big.Int) (*big.Int, error) {
	return _Foundation.Contract.GetMinOfferAmount(&_Foundation.CallOpts, nftContract, tokenId)
}

// GetOffer is a free data retrieval call binding the contract method 0xac71045e.
//
// Solidity: function getOffer(address nftContract, uint256 tokenId) view returns(address buyer, uint256 expiration, uint256 amount)
func (_Foundation *FoundationCaller) GetOffer(opts *bind.CallOpts, nftContract common.Address, tokenId *big.Int) (struct {
	Buyer      common.Address
	Expiration *big.Int
	Amount     *big.Int
}, error) {
	var out []interface{}
	err := _Foundation.contract.Call(opts, &out, "getOffer", nftContract, tokenId)

	outstruct := new(struct {
		Buyer      common.Address
		Expiration *big.Int
		Amount     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Buyer = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Expiration = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Amount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetOffer is a free data retrieval call binding the contract method 0xac71045e.
//
// Solidity: function getOffer(address nftContract, uint256 tokenId) view returns(address buyer, uint256 expiration, uint256 amount)
func (_Foundation *FoundationSession) GetOffer(nftContract common.Address, tokenId *big.Int) (struct {
	Buyer      common.Address
	Expiration *big.Int
	Amount     *big.Int
}, error) {
	return _Foundation.Contract.GetOffer(&_Foundation.CallOpts, nftContract, tokenId)
}

// GetOffer is a free data retrieval call binding the contract method 0xac71045e.
//
// Solidity: function getOffer(address nftContract, uint256 tokenId) view returns(address buyer, uint256 expiration, uint256 amount)
func (_Foundation *FoundationCallerSession) GetOffer(nftContract common.Address, tokenId *big.Int) (struct {
	Buyer      common.Address
	Expiration *big.Int
	Amount     *big.Int
}, error) {
	return _Foundation.Contract.GetOffer(&_Foundation.CallOpts, nftContract, tokenId)
}

// GetOfferReferrer is a free data retrieval call binding the contract method 0x262907c5.
//
// Solidity: function getOfferReferrer(address nftContract, uint256 tokenId) view returns(address referrer)
func (_Foundation *FoundationCaller) GetOfferReferrer(opts *bind.CallOpts, nftContract common.Address, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Foundation.contract.Call(opts, &out, "getOfferReferrer", nftContract, tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOfferReferrer is a free data retrieval call binding the contract method 0x262907c5.
//
// Solidity: function getOfferReferrer(address nftContract, uint256 tokenId) view returns(address referrer)
func (_Foundation *FoundationSession) GetOfferReferrer(nftContract common.Address, tokenId *big.Int) (common.Address, error) {
	return _Foundation.Contract.GetOfferReferrer(&_Foundation.CallOpts, nftContract, tokenId)
}

// GetOfferReferrer is a free data retrieval call binding the contract method 0x262907c5.
//
// Solidity: function getOfferReferrer(address nftContract, uint256 tokenId) view returns(address referrer)
func (_Foundation *FoundationCallerSession) GetOfferReferrer(nftContract common.Address, tokenId *big.Int) (common.Address, error) {
	return _Foundation.Contract.GetOfferReferrer(&_Foundation.CallOpts, nftContract, tokenId)
}

// GetReserveAuction is a free data retrieval call binding the contract method 0x9e79b41f.
//
// Solidity: function getReserveAuction(uint256 auctionId) view returns((address,uint256,address,uint256,uint256,uint256,address,uint256) auction)
func (_Foundation *FoundationCaller) GetReserveAuction(opts *bind.CallOpts, auctionId *big.Int) (NFTMarketReserveAuctionReserveAuction, error) {
	var out []interface{}
	err := _Foundation.contract.Call(opts, &out, "getReserveAuction", auctionId)

	if err != nil {
		return *new(NFTMarketReserveAuctionReserveAuction), err
	}

	out0 := *abi.ConvertType(out[0], new(NFTMarketReserveAuctionReserveAuction)).(*NFTMarketReserveAuctionReserveAuction)

	return out0, err

}

// GetReserveAuction is a free data retrieval call binding the contract method 0x9e79b41f.
//
// Solidity: function getReserveAuction(uint256 auctionId) view returns((address,uint256,address,uint256,uint256,uint256,address,uint256) auction)
func (_Foundation *FoundationSession) GetReserveAuction(auctionId *big.Int) (NFTMarketReserveAuctionReserveAuction, error) {
	return _Foundation.Contract.GetReserveAuction(&_Foundation.CallOpts, auctionId)
}

// GetReserveAuction is a free data retrieval call binding the contract method 0x9e79b41f.
//
// Solidity: function getReserveAuction(uint256 auctionId) view returns((address,uint256,address,uint256,uint256,uint256,address,uint256) auction)
func (_Foundation *FoundationCallerSession) GetReserveAuction(auctionId *big.Int) (NFTMarketReserveAuctionReserveAuction, error) {
	return _Foundation.Contract.GetReserveAuction(&_Foundation.CallOpts, auctionId)
}

// GetReserveAuctionBidReferrer is a free data retrieval call binding the contract method 0x9e64ba6c.
//
// Solidity: function getReserveAuctionBidReferrer(uint256 auctionId) view returns(address referrer)
func (_Foundation *FoundationCaller) GetReserveAuctionBidReferrer(opts *bind.CallOpts, auctionId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Foundation.contract.Call(opts, &out, "getReserveAuctionBidReferrer", auctionId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetReserveAuctionBidReferrer is a free data retrieval call binding the contract method 0x9e64ba6c.
//
// Solidity: function getReserveAuctionBidReferrer(uint256 auctionId) view returns(address referrer)
func (_Foundation *FoundationSession) GetReserveAuctionBidReferrer(auctionId *big.Int) (common.Address, error) {
	return _Foundation.Contract.GetReserveAuctionBidReferrer(&_Foundation.CallOpts, auctionId)
}

// GetReserveAuctionBidReferrer is a free data retrieval call binding the contract method 0x9e64ba6c.
//
// Solidity: function getReserveAuctionBidReferrer(uint256 auctionId) view returns(address referrer)
func (_Foundation *FoundationCallerSession) GetReserveAuctionBidReferrer(auctionId *big.Int) (common.Address, error) {
	return _Foundation.Contract.GetReserveAuctionBidReferrer(&_Foundation.CallOpts, auctionId)
}

// GetReserveAuctionIdFor is a free data retrieval call binding the contract method 0x2ab2b52b.
//
// Solidity: function getReserveAuctionIdFor(address nftContract, uint256 tokenId) view returns(uint256 auctionId)
func (_Foundation *FoundationCaller) GetReserveAuctionIdFor(opts *bind.CallOpts, nftContract common.Address, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Foundation.contract.Call(opts, &out, "getReserveAuctionIdFor", nftContract, tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReserveAuctionIdFor is a free data retrieval call binding the contract method 0x2ab2b52b.
//
// Solidity: function getReserveAuctionIdFor(address nftContract, uint256 tokenId) view returns(uint256 auctionId)
func (_Foundation *FoundationSession) GetReserveAuctionIdFor(nftContract common.Address, tokenId *big.Int) (*big.Int, error) {
	return _Foundation.Contract.GetReserveAuctionIdFor(&_Foundation.CallOpts, nftContract, tokenId)
}

// GetReserveAuctionIdFor is a free data retrieval call binding the contract method 0x2ab2b52b.
//
// Solidity: function getReserveAuctionIdFor(address nftContract, uint256 tokenId) view returns(uint256 auctionId)
func (_Foundation *FoundationCallerSession) GetReserveAuctionIdFor(nftContract common.Address, tokenId *big.Int) (*big.Int, error) {
	return _Foundation.Contract.GetReserveAuctionIdFor(&_Foundation.CallOpts, nftContract, tokenId)
}

// GetRoyaltyRegistry is a free data retrieval call binding the contract method 0xdaa351d4.
//
// Solidity: function getRoyaltyRegistry() view returns(address registry)
func (_Foundation *FoundationCaller) GetRoyaltyRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Foundation.contract.Call(opts, &out, "getRoyaltyRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoyaltyRegistry is a free data retrieval call binding the contract method 0xdaa351d4.
//
// Solidity: function getRoyaltyRegistry() view returns(address registry)
func (_Foundation *FoundationSession) GetRoyaltyRegistry() (common.Address, error) {
	return _Foundation.Contract.GetRoyaltyRegistry(&_Foundation.CallOpts)
}

// GetRoyaltyRegistry is a free data retrieval call binding the contract method 0xdaa351d4.
//
// Solidity: function getRoyaltyRegistry() view returns(address registry)
func (_Foundation *FoundationCallerSession) GetRoyaltyRegistry() (common.Address, error) {
	return _Foundation.Contract.GetRoyaltyRegistry(&_Foundation.CallOpts)
}

// GetSellerOf is a free data retrieval call binding the contract method 0x4fca06c6.
//
// Solidity: function getSellerOf(address nftContract, uint256 tokenId) view returns(address seller)
func (_Foundation *FoundationCaller) GetSellerOf(opts *bind.CallOpts, nftContract common.Address, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Foundation.contract.Call(opts, &out, "getSellerOf", nftContract, tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSellerOf is a free data retrieval call binding the contract method 0x4fca06c6.
//
// Solidity: function getSellerOf(address nftContract, uint256 tokenId) view returns(address seller)
func (_Foundation *FoundationSession) GetSellerOf(nftContract common.Address, tokenId *big.Int) (common.Address, error) {
	return _Foundation.Contract.GetSellerOf(&_Foundation.CallOpts, nftContract, tokenId)
}

// GetSellerOf is a free data retrieval call binding the contract method 0x4fca06c6.
//
// Solidity: function getSellerOf(address nftContract, uint256 tokenId) view returns(address seller)
func (_Foundation *FoundationCallerSession) GetSellerOf(nftContract common.Address, tokenId *big.Int) (common.Address, error) {
	return _Foundation.Contract.GetSellerOf(&_Foundation.CallOpts, nftContract, tokenId)
}

// InternalGetImmutableRoyalties is a free data retrieval call binding the contract method 0x0d7daf3e.
//
// Solidity: function internalGetImmutableRoyalties(address nftContract, uint256 tokenId) view returns(address[] recipients, uint256[] splitPerRecipientInBasisPoints)
func (_Foundation *FoundationCaller) InternalGetImmutableRoyalties(opts *bind.CallOpts, nftContract common.Address, tokenId *big.Int) (struct {
	Recipients                     []common.Address
	SplitPerRecipientInBasisPoints []*big.Int
}, error) {
	var out []interface{}
	err := _Foundation.contract.Call(opts, &out, "internalGetImmutableRoyalties", nftContract, tokenId)

	outstruct := new(struct {
		Recipients                     []common.Address
		SplitPerRecipientInBasisPoints []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Recipients = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.SplitPerRecipientInBasisPoints = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// InternalGetImmutableRoyalties is a free data retrieval call binding the contract method 0x0d7daf3e.
//
// Solidity: function internalGetImmutableRoyalties(address nftContract, uint256 tokenId) view returns(address[] recipients, uint256[] splitPerRecipientInBasisPoints)
func (_Foundation *FoundationSession) InternalGetImmutableRoyalties(nftContract common.Address, tokenId *big.Int) (struct {
	Recipients                     []common.Address
	SplitPerRecipientInBasisPoints []*big.Int
}, error) {
	return _Foundation.Contract.InternalGetImmutableRoyalties(&_Foundation.CallOpts, nftContract, tokenId)
}

// InternalGetImmutableRoyalties is a free data retrieval call binding the contract method 0x0d7daf3e.
//
// Solidity: function internalGetImmutableRoyalties(address nftContract, uint256 tokenId) view returns(address[] recipients, uint256[] splitPerRecipientInBasisPoints)
func (_Foundation *FoundationCallerSession) InternalGetImmutableRoyalties(nftContract common.Address, tokenId *big.Int) (struct {
	Recipients                     []common.Address
	SplitPerRecipientInBasisPoints []*big.Int
}, error) {
	return _Foundation.Contract.InternalGetImmutableRoyalties(&_Foundation.CallOpts, nftContract, tokenId)
}

// InternalGetMutableRoyalties is a free data retrieval call binding the contract method 0xefef76f8.
//
// Solidity: function internalGetMutableRoyalties(address nftContract, uint256 tokenId, address creator) view returns(address[] recipients, uint256[] splitPerRecipientInBasisPoints)
func (_Foundation *FoundationCaller) InternalGetMutableRoyalties(opts *bind.CallOpts, nftContract common.Address, tokenId *big.Int, creator common.Address) (struct {
	Recipients                     []common.Address
	SplitPerRecipientInBasisPoints []*big.Int
}, error) {
	var out []interface{}
	err := _Foundation.contract.Call(opts, &out, "internalGetMutableRoyalties", nftContract, tokenId, creator)

	outstruct := new(struct {
		Recipients                     []common.Address
		SplitPerRecipientInBasisPoints []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Recipients = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.SplitPerRecipientInBasisPoints = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// InternalGetMutableRoyalties is a free data retrieval call binding the contract method 0xefef76f8.
//
// Solidity: function internalGetMutableRoyalties(address nftContract, uint256 tokenId, address creator) view returns(address[] recipients, uint256[] splitPerRecipientInBasisPoints)
func (_Foundation *FoundationSession) InternalGetMutableRoyalties(nftContract common.Address, tokenId *big.Int, creator common.Address) (struct {
	Recipients                     []common.Address
	SplitPerRecipientInBasisPoints []*big.Int
}, error) {
	return _Foundation.Contract.InternalGetMutableRoyalties(&_Foundation.CallOpts, nftContract, tokenId, creator)
}

// InternalGetMutableRoyalties is a free data retrieval call binding the contract method 0xefef76f8.
//
// Solidity: function internalGetMutableRoyalties(address nftContract, uint256 tokenId, address creator) view returns(address[] recipients, uint256[] splitPerRecipientInBasisPoints)
func (_Foundation *FoundationCallerSession) InternalGetMutableRoyalties(nftContract common.Address, tokenId *big.Int, creator common.Address) (struct {
	Recipients                     []common.Address
	SplitPerRecipientInBasisPoints []*big.Int
}, error) {
	return _Foundation.Contract.InternalGetMutableRoyalties(&_Foundation.CallOpts, nftContract, tokenId, creator)
}

// InternalGetTokenCreator is a free data retrieval call binding the contract method 0x4c542f77.
//
// Solidity: function internalGetTokenCreator(address nftContract, uint256 tokenId) view returns(address creator)
func (_Foundation *FoundationCaller) InternalGetTokenCreator(opts *bind.CallOpts, nftContract common.Address, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Foundation.contract.Call(opts, &out, "internalGetTokenCreator", nftContract, tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InternalGetTokenCreator is a free data retrieval call binding the contract method 0x4c542f77.
//
// Solidity: function internalGetTokenCreator(address nftContract, uint256 tokenId) view returns(address creator)
func (_Foundation *FoundationSession) InternalGetTokenCreator(nftContract common.Address, tokenId *big.Int) (common.Address, error) {
	return _Foundation.Contract.InternalGetTokenCreator(&_Foundation.CallOpts, nftContract, tokenId)
}

// InternalGetTokenCreator is a free data retrieval call binding the contract method 0x4c542f77.
//
// Solidity: function internalGetTokenCreator(address nftContract, uint256 tokenId) view returns(address creator)
func (_Foundation *FoundationCallerSession) InternalGetTokenCreator(nftContract common.Address, tokenId *big.Int) (common.Address, error) {
	return _Foundation.Contract.InternalGetTokenCreator(&_Foundation.CallOpts, nftContract, tokenId)
}

// IsAllowedSellerForExhibition is a free data retrieval call binding the contract method 0x2e06db96.
//
// Solidity: function isAllowedSellerForExhibition(uint256 exhibitionId, address seller) view returns(bool allowedSeller)
func (_Foundation *FoundationCaller) IsAllowedSellerForExhibition(opts *bind.CallOpts, exhibitionId *big.Int, seller common.Address) (bool, error) {
	var out []interface{}
	err := _Foundation.contract.Call(opts, &out, "isAllowedSellerForExhibition", exhibitionId, seller)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAllowedSellerForExhibition is a free data retrieval call binding the contract method 0x2e06db96.
//
// Solidity: function isAllowedSellerForExhibition(uint256 exhibitionId, address seller) view returns(bool allowedSeller)
func (_Foundation *FoundationSession) IsAllowedSellerForExhibition(exhibitionId *big.Int, seller common.Address) (bool, error) {
	return _Foundation.Contract.IsAllowedSellerForExhibition(&_Foundation.CallOpts, exhibitionId, seller)
}

// IsAllowedSellerForExhibition is a free data retrieval call binding the contract method 0x2e06db96.
//
// Solidity: function isAllowedSellerForExhibition(uint256 exhibitionId, address seller) view returns(bool allowedSeller)
func (_Foundation *FoundationCallerSession) IsAllowedSellerForExhibition(exhibitionId *big.Int, seller common.Address) (bool, error) {
	return _Foundation.Contract.IsAllowedSellerForExhibition(&_Foundation.CallOpts, exhibitionId, seller)
}

// AcceptOffer is a paid mutator transaction binding the contract method 0x29e0e160.
//
// Solidity: function acceptOffer(address nftContract, uint256 tokenId, address offerFrom, uint256 minAmount) returns()
func (_Foundation *FoundationTransactor) AcceptOffer(opts *bind.TransactOpts, nftContract common.Address, tokenId *big.Int, offerFrom common.Address, minAmount *big.Int) (*types.Transaction, error) {
	return _Foundation.contract.Transact(opts, "acceptOffer", nftContract, tokenId, offerFrom, minAmount)
}

// AcceptOffer is a paid mutator transaction binding the contract method 0x29e0e160.
//
// Solidity: function acceptOffer(address nftContract, uint256 tokenId, address offerFrom, uint256 minAmount) returns()
func (_Foundation *FoundationSession) AcceptOffer(nftContract common.Address, tokenId *big.Int, offerFrom common.Address, minAmount *big.Int) (*types.Transaction, error) {
	return _Foundation.Contract.AcceptOffer(&_Foundation.TransactOpts, nftContract, tokenId, offerFrom, minAmount)
}

// AcceptOffer is a paid mutator transaction binding the contract method 0x29e0e160.
//
// Solidity: function acceptOffer(address nftContract, uint256 tokenId, address offerFrom, uint256 minAmount) returns()
func (_Foundation *FoundationTransactorSession) AcceptOffer(nftContract common.Address, tokenId *big.Int, offerFrom common.Address, minAmount *big.Int) (*types.Transaction, error) {
	return _Foundation.Contract.AcceptOffer(&_Foundation.TransactOpts, nftContract, tokenId, offerFrom, minAmount)
}

// AddSellersToExhibition is a paid mutator transaction binding the contract method 0x6512ed2d.
//
// Solidity: function addSellersToExhibition(uint256 exhibitionId, address[] sellers) returns()
func (_Foundation *FoundationTransactor) AddSellersToExhibition(opts *bind.TransactOpts, exhibitionId *big.Int, sellers []common.Address) (*types.Transaction, error) {
	return _Foundation.contract.Transact(opts, "addSellersToExhibition", exhibitionId, sellers)
}

// AddSellersToExhibition is a paid mutator transaction binding the contract method 0x6512ed2d.
//
// Solidity: function addSellersToExhibition(uint256 exhibitionId, address[] sellers) returns()
func (_Foundation *FoundationSession) AddSellersToExhibition(exhibitionId *big.Int, sellers []common.Address) (*types.Transaction, error) {
	return _Foundation.Contract.AddSellersToExhibition(&_Foundation.TransactOpts, exhibitionId, sellers)
}

// AddSellersToExhibition is a paid mutator transaction binding the contract method 0x6512ed2d.
//
// Solidity: function addSellersToExhibition(uint256 exhibitionId, address[] sellers) returns()
func (_Foundation *FoundationTransactorSession) AddSellersToExhibition(exhibitionId *big.Int, sellers []common.Address) (*types.Transaction, error) {
	return _Foundation.Contract.AddSellersToExhibition(&_Foundation.TransactOpts, exhibitionId, sellers)
}

// Buy is a paid mutator transaction binding the contract method 0xa59ac6dd.
//
// Solidity: function buy(address nftContract, uint256 tokenId, uint256 maxPrice) payable returns()
func (_Foundation *FoundationTransactor) Buy(opts *bind.TransactOpts, nftContract common.Address, tokenId *big.Int, maxPrice *big.Int) (*types.Transaction, error) {
	return _Foundation.contract.Transact(opts, "buy", nftContract, tokenId, maxPrice)
}

// Buy is a paid mutator transaction binding the contract method 0xa59ac6dd.
//
// Solidity: function buy(address nftContract, uint256 tokenId, uint256 maxPrice) payable returns()
func (_Foundation *FoundationSession) Buy(nftContract common.Address, tokenId *big.Int, maxPrice *big.Int) (*types.Transaction, error) {
	return _Foundation.Contract.Buy(&_Foundation.TransactOpts, nftContract, tokenId, maxPrice)
}

// Buy is a paid mutator transaction binding the contract method 0xa59ac6dd.
//
// Solidity: function buy(address nftContract, uint256 tokenId, uint256 maxPrice) payable returns()
func (_Foundation *FoundationTransactorSession) Buy(nftContract common.Address, tokenId *big.Int, maxPrice *big.Int) (*types.Transaction, error) {
	return _Foundation.Contract.Buy(&_Foundation.TransactOpts, nftContract, tokenId, maxPrice)
}

// BuyV2 is a paid mutator transaction binding the contract method 0xb01ef608.
//
// Solidity: function buyV2(address nftContract, uint256 tokenId, uint256 maxPrice, address referrer) payable returns()
func (_Foundation *FoundationTransactor) BuyV2(opts *bind.TransactOpts, nftContract common.Address, tokenId *big.Int, maxPrice *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _Foundation.contract.Transact(opts, "buyV2", nftContract, tokenId, maxPrice, referrer)
}

// BuyV2 is a paid mutator transaction binding the contract method 0xb01ef608.
//
// Solidity: function buyV2(address nftContract, uint256 tokenId, uint256 maxPrice, address referrer) payable returns()
func (_Foundation *FoundationSession) BuyV2(nftContract common.Address, tokenId *big.Int, maxPrice *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _Foundation.Contract.BuyV2(&_Foundation.TransactOpts, nftContract, tokenId, maxPrice, referrer)
}

// BuyV2 is a paid mutator transaction binding the contract method 0xb01ef608.
//
// Solidity: function buyV2(address nftContract, uint256 tokenId, uint256 maxPrice, address referrer) payable returns()
func (_Foundation *FoundationTransactorSession) BuyV2(nftContract common.Address, tokenId *big.Int, maxPrice *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _Foundation.Contract.BuyV2(&_Foundation.TransactOpts, nftContract, tokenId, maxPrice, referrer)
}

// CancelBuyPrice is a paid mutator transaction binding the contract method 0x21561935.
//
// Solidity: function cancelBuyPrice(address nftContract, uint256 tokenId) returns()
func (_Foundation *FoundationTransactor) CancelBuyPrice(opts *bind.TransactOpts, nftContract common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Foundation.contract.Transact(opts, "cancelBuyPrice", nftContract, tokenId)
}

// CancelBuyPrice is a paid mutator transaction binding the contract method 0x21561935.
//
// Solidity: function cancelBuyPrice(address nftContract, uint256 tokenId) returns()
func (_Foundation *FoundationSession) CancelBuyPrice(nftContract common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Foundation.Contract.CancelBuyPrice(&_Foundation.TransactOpts, nftContract, tokenId)
}

// CancelBuyPrice is a paid mutator transaction binding the contract method 0x21561935.
//
// Solidity: function cancelBuyPrice(address nftContract, uint256 tokenId) returns()
func (_Foundation *FoundationTransactorSession) CancelBuyPrice(nftContract common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Foundation.Contract.CancelBuyPrice(&_Foundation.TransactOpts, nftContract, tokenId)
}

// CancelReserveAuction is a paid mutator transaction binding the contract method 0x21506fff.
//
// Solidity: function cancelReserveAuction(uint256 auctionId) returns()
func (_Foundation *FoundationTransactor) CancelReserveAuction(opts *bind.TransactOpts, auctionId *big.Int) (*types.Transaction, error) {
	return _Foundation.contract.Transact(opts, "cancelReserveAuction", auctionId)
}

// CancelReserveAuction is a paid mutator transaction binding the contract method 0x21506fff.
//
// Solidity: function cancelReserveAuction(uint256 auctionId) returns()
func (_Foundation *FoundationSession) CancelReserveAuction(auctionId *big.Int) (*types.Transaction, error) {
	return _Foundation.Contract.CancelReserveAuction(&_Foundation.TransactOpts, auctionId)
}

// CancelReserveAuction is a paid mutator transaction binding the contract method 0x21506fff.
//
// Solidity: function cancelReserveAuction(uint256 auctionId) returns()
func (_Foundation *FoundationTransactorSession) CancelReserveAuction(auctionId *big.Int) (*types.Transaction, error) {
	return _Foundation.Contract.CancelReserveAuction(&_Foundation.TransactOpts, auctionId)
}

// CreateExhibition is a paid mutator transaction binding the contract method 0x445738d8.
//
// Solidity: function createExhibition(string name, uint16 takeRateInBasisPoints, address[] sellers) returns(uint256 exhibitionId)
func (_Foundation *FoundationTransactor) CreateExhibition(opts *bind.TransactOpts, name string, takeRateInBasisPoints uint16, sellers []common.Address) (*types.Transaction, error) {
	return _Foundation.contract.Transact(opts, "createExhibition", name, takeRateInBasisPoints, sellers)
}

// CreateExhibition is a paid mutator transaction binding the contract method 0x445738d8.
//
// Solidity: function createExhibition(string name, uint16 takeRateInBasisPoints, address[] sellers) returns(uint256 exhibitionId)
func (_Foundation *FoundationSession) CreateExhibition(name string, takeRateInBasisPoints uint16, sellers []common.Address) (*types.Transaction, error) {
	return _Foundation.Contract.CreateExhibition(&_Foundation.TransactOpts, name, takeRateInBasisPoints, sellers)
}

// CreateExhibition is a paid mutator transaction binding the contract method 0x445738d8.
//
// Solidity: function createExhibition(string name, uint16 takeRateInBasisPoints, address[] sellers) returns(uint256 exhibitionId)
func (_Foundation *FoundationTransactorSession) CreateExhibition(name string, takeRateInBasisPoints uint16, sellers []common.Address) (*types.Transaction, error) {
	return _Foundation.Contract.CreateExhibition(&_Foundation.TransactOpts, name, takeRateInBasisPoints, sellers)
}

// CreateReserveAuction is a paid mutator transaction binding the contract method 0x4ce6931a.
//
// Solidity: function createReserveAuction(address nftContract, uint256 tokenId, uint256 reservePrice) returns()
func (_Foundation *FoundationTransactor) CreateReserveAuction(opts *bind.TransactOpts, nftContract common.Address, tokenId *big.Int, reservePrice *big.Int) (*types.Transaction, error) {
	return _Foundation.contract.Transact(opts, "createReserveAuction", nftContract, tokenId, reservePrice)
}

// CreateReserveAuction is a paid mutator transaction binding the contract method 0x4ce6931a.
//
// Solidity: function createReserveAuction(address nftContract, uint256 tokenId, uint256 reservePrice) returns()
func (_Foundation *FoundationSession) CreateReserveAuction(nftContract common.Address, tokenId *big.Int, reservePrice *big.Int) (*types.Transaction, error) {
	return _Foundation.Contract.CreateReserveAuction(&_Foundation.TransactOpts, nftContract, tokenId, reservePrice)
}

// CreateReserveAuction is a paid mutator transaction binding the contract method 0x4ce6931a.
//
// Solidity: function createReserveAuction(address nftContract, uint256 tokenId, uint256 reservePrice) returns()
func (_Foundation *FoundationTransactorSession) CreateReserveAuction(nftContract common.Address, tokenId *big.Int, reservePrice *big.Int) (*types.Transaction, error) {
	return _Foundation.Contract.CreateReserveAuction(&_Foundation.TransactOpts, nftContract, tokenId, reservePrice)
}

// CreateReserveAuctionV2 is a paid mutator transaction binding the contract method 0xbeb5127c.
//
// Solidity: function createReserveAuctionV2(address nftContract, uint256 tokenId, uint256 reservePrice, uint256 exhibitionId) returns(uint256 auctionId)
func (_Foundation *FoundationTransactor) CreateReserveAuctionV2(opts *bind.TransactOpts, nftContract common.Address, tokenId *big.Int, reservePrice *big.Int, exhibitionId *big.Int) (*types.Transaction, error) {
	return _Foundation.contract.Transact(opts, "createReserveAuctionV2", nftContract, tokenId, reservePrice, exhibitionId)
}

// CreateReserveAuctionV2 is a paid mutator transaction binding the contract method 0xbeb5127c.
//
// Solidity: function createReserveAuctionV2(address nftContract, uint256 tokenId, uint256 reservePrice, uint256 exhibitionId) returns(uint256 auctionId)
func (_Foundation *FoundationSession) CreateReserveAuctionV2(nftContract common.Address, tokenId *big.Int, reservePrice *big.Int, exhibitionId *big.Int) (*types.Transaction, error) {
	return _Foundation.Contract.CreateReserveAuctionV2(&_Foundation.TransactOpts, nftContract, tokenId, reservePrice, exhibitionId)
}

// CreateReserveAuctionV2 is a paid mutator transaction binding the contract method 0xbeb5127c.
//
// Solidity: function createReserveAuctionV2(address nftContract, uint256 tokenId, uint256 reservePrice, uint256 exhibitionId) returns(uint256 auctionId)
func (_Foundation *FoundationTransactorSession) CreateReserveAuctionV2(nftContract common.Address, tokenId *big.Int, reservePrice *big.Int, exhibitionId *big.Int) (*types.Transaction, error) {
	return _Foundation.Contract.CreateReserveAuctionV2(&_Foundation.TransactOpts, nftContract, tokenId, reservePrice, exhibitionId)
}

// DeleteExhibition is a paid mutator transaction binding the contract method 0x7b3a5884.
//
// Solidity: function deleteExhibition(uint256 exhibitionId) returns()
func (_Foundation *FoundationTransactor) DeleteExhibition(opts *bind.TransactOpts, exhibitionId *big.Int) (*types.Transaction, error) {
	return _Foundation.contract.Transact(opts, "deleteExhibition", exhibitionId)
}

// DeleteExhibition is a paid mutator transaction binding the contract method 0x7b3a5884.
//
// Solidity: function deleteExhibition(uint256 exhibitionId) returns()
func (_Foundation *FoundationSession) DeleteExhibition(exhibitionId *big.Int) (*types.Transaction, error) {
	return _Foundation.Contract.DeleteExhibition(&_Foundation.TransactOpts, exhibitionId)
}

// DeleteExhibition is a paid mutator transaction binding the contract method 0x7b3a5884.
//
// Solidity: function deleteExhibition(uint256 exhibitionId) returns()
func (_Foundation *FoundationTransactorSession) DeleteExhibition(exhibitionId *big.Int) (*types.Transaction, error) {
	return _Foundation.Contract.DeleteExhibition(&_Foundation.TransactOpts, exhibitionId)
}

// FinalizeReserveAuction is a paid mutator transaction binding the contract method 0x7430e0c6.
//
// Solidity: function finalizeReserveAuction(uint256 auctionId) returns()
func (_Foundation *FoundationTransactor) FinalizeReserveAuction(opts *bind.TransactOpts, auctionId *big.Int) (*types.Transaction, error) {
	return _Foundation.contract.Transact(opts, "finalizeReserveAuction", auctionId)
}

// FinalizeReserveAuction is a paid mutator transaction binding the contract method 0x7430e0c6.
//
// Solidity: function finalizeReserveAuction(uint256 auctionId) returns()
func (_Foundation *FoundationSession) FinalizeReserveAuction(auctionId *big.Int) (*types.Transaction, error) {
	return _Foundation.Contract.FinalizeReserveAuction(&_Foundation.TransactOpts, auctionId)
}

// FinalizeReserveAuction is a paid mutator transaction binding the contract method 0x7430e0c6.
//
// Solidity: function finalizeReserveAuction(uint256 auctionId) returns()
func (_Foundation *FoundationTransactorSession) FinalizeReserveAuction(auctionId *big.Int) (*types.Transaction, error) {
	return _Foundation.Contract.FinalizeReserveAuction(&_Foundation.TransactOpts, auctionId)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Foundation *FoundationTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Foundation.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Foundation *FoundationSession) Initialize() (*types.Transaction, error) {
	return _Foundation.Contract.Initialize(&_Foundation.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Foundation *FoundationTransactorSession) Initialize() (*types.Transaction, error) {
	return _Foundation.Contract.Initialize(&_Foundation.TransactOpts)
}

// MakeOfferV2 is a paid mutator transaction binding the contract method 0x614b151c.
//
// Solidity: function makeOfferV2(address nftContract, uint256 tokenId, uint256 amount, address referrer) payable returns(uint256 expiration)
func (_Foundation *FoundationTransactor) MakeOfferV2(opts *bind.TransactOpts, nftContract common.Address, tokenId *big.Int, amount *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _Foundation.contract.Transact(opts, "makeOfferV2", nftContract, tokenId, amount, referrer)
}

// MakeOfferV2 is a paid mutator transaction binding the contract method 0x614b151c.
//
// Solidity: function makeOfferV2(address nftContract, uint256 tokenId, uint256 amount, address referrer) payable returns(uint256 expiration)
func (_Foundation *FoundationSession) MakeOfferV2(nftContract common.Address, tokenId *big.Int, amount *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _Foundation.Contract.MakeOfferV2(&_Foundation.TransactOpts, nftContract, tokenId, amount, referrer)
}

// MakeOfferV2 is a paid mutator transaction binding the contract method 0x614b151c.
//
// Solidity: function makeOfferV2(address nftContract, uint256 tokenId, uint256 amount, address referrer) payable returns(uint256 expiration)
func (_Foundation *FoundationTransactorSession) MakeOfferV2(nftContract common.Address, tokenId *big.Int, amount *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _Foundation.Contract.MakeOfferV2(&_Foundation.TransactOpts, nftContract, tokenId, amount, referrer)
}

// PlaceBid is a paid mutator transaction binding the contract method 0x9979ef45.
//
// Solidity: function placeBid(uint256 auctionId) payable returns()
func (_Foundation *FoundationTransactor) PlaceBid(opts *bind.TransactOpts, auctionId *big.Int) (*types.Transaction, error) {
	return _Foundation.contract.Transact(opts, "placeBid", auctionId)
}

// PlaceBid is a paid mutator transaction binding the contract method 0x9979ef45.
//
// Solidity: function placeBid(uint256 auctionId) payable returns()
func (_Foundation *FoundationSession) PlaceBid(auctionId *big.Int) (*types.Transaction, error) {
	return _Foundation.Contract.PlaceBid(&_Foundation.TransactOpts, auctionId)
}

// PlaceBid is a paid mutator transaction binding the contract method 0x9979ef45.
//
// Solidity: function placeBid(uint256 auctionId) payable returns()
func (_Foundation *FoundationTransactorSession) PlaceBid(auctionId *big.Int) (*types.Transaction, error) {
	return _Foundation.Contract.PlaceBid(&_Foundation.TransactOpts, auctionId)
}

// PlaceBidV2 is a paid mutator transaction binding the contract method 0xb6aff8c1.
//
// Solidity: function placeBidV2(uint256 auctionId, uint256 amount, address referrer) payable returns()
func (_Foundation *FoundationTransactor) PlaceBidV2(opts *bind.TransactOpts, auctionId *big.Int, amount *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _Foundation.contract.Transact(opts, "placeBidV2", auctionId, amount, referrer)
}

// PlaceBidV2 is a paid mutator transaction binding the contract method 0xb6aff8c1.
//
// Solidity: function placeBidV2(uint256 auctionId, uint256 amount, address referrer) payable returns()
func (_Foundation *FoundationSession) PlaceBidV2(auctionId *big.Int, amount *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _Foundation.Contract.PlaceBidV2(&_Foundation.TransactOpts, auctionId, amount, referrer)
}

// PlaceBidV2 is a paid mutator transaction binding the contract method 0xb6aff8c1.
//
// Solidity: function placeBidV2(uint256 auctionId, uint256 amount, address referrer) payable returns()
func (_Foundation *FoundationTransactorSession) PlaceBidV2(auctionId *big.Int, amount *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _Foundation.Contract.PlaceBidV2(&_Foundation.TransactOpts, auctionId, amount, referrer)
}

// SetBuyPrice is a paid mutator transaction binding the contract method 0x798bac8d.
//
// Solidity: function setBuyPrice(address nftContract, uint256 tokenId, uint256 price) returns()
func (_Foundation *FoundationTransactor) SetBuyPrice(opts *bind.TransactOpts, nftContract common.Address, tokenId *big.Int, price *big.Int) (*types.Transaction, error) {
	return _Foundation.contract.Transact(opts, "setBuyPrice", nftContract, tokenId, price)
}

// SetBuyPrice is a paid mutator transaction binding the contract method 0x798bac8d.
//
// Solidity: function setBuyPrice(address nftContract, uint256 tokenId, uint256 price) returns()
func (_Foundation *FoundationSession) SetBuyPrice(nftContract common.Address, tokenId *big.Int, price *big.Int) (*types.Transaction, error) {
	return _Foundation.Contract.SetBuyPrice(&_Foundation.TransactOpts, nftContract, tokenId, price)
}

// SetBuyPrice is a paid mutator transaction binding the contract method 0x798bac8d.
//
// Solidity: function setBuyPrice(address nftContract, uint256 tokenId, uint256 price) returns()
func (_Foundation *FoundationTransactorSession) SetBuyPrice(nftContract common.Address, tokenId *big.Int, price *big.Int) (*types.Transaction, error) {
	return _Foundation.Contract.SetBuyPrice(&_Foundation.TransactOpts, nftContract, tokenId, price)
}

// UpdateReserveAuction is a paid mutator transaction binding the contract method 0x03ec16d7.
//
// Solidity: function updateReserveAuction(uint256 auctionId, uint256 reservePrice) returns()
func (_Foundation *FoundationTransactor) UpdateReserveAuction(opts *bind.TransactOpts, auctionId *big.Int, reservePrice *big.Int) (*types.Transaction, error) {
	return _Foundation.contract.Transact(opts, "updateReserveAuction", auctionId, reservePrice)
}

// UpdateReserveAuction is a paid mutator transaction binding the contract method 0x03ec16d7.
//
// Solidity: function updateReserveAuction(uint256 auctionId, uint256 reservePrice) returns()
func (_Foundation *FoundationSession) UpdateReserveAuction(auctionId *big.Int, reservePrice *big.Int) (*types.Transaction, error) {
	return _Foundation.Contract.UpdateReserveAuction(&_Foundation.TransactOpts, auctionId, reservePrice)
}

// UpdateReserveAuction is a paid mutator transaction binding the contract method 0x03ec16d7.
//
// Solidity: function updateReserveAuction(uint256 auctionId, uint256 reservePrice) returns()
func (_Foundation *FoundationTransactorSession) UpdateReserveAuction(auctionId *big.Int, reservePrice *big.Int) (*types.Transaction, error) {
	return _Foundation.Contract.UpdateReserveAuction(&_Foundation.TransactOpts, auctionId, reservePrice)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Foundation *FoundationTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Foundation.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Foundation *FoundationSession) Receive() (*types.Transaction, error) {
	return _Foundation.Contract.Receive(&_Foundation.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Foundation *FoundationTransactorSession) Receive() (*types.Transaction, error) {
	return _Foundation.Contract.Receive(&_Foundation.TransactOpts)
}

// FoundationBuyPriceAcceptedIterator is returned from FilterBuyPriceAccepted and is used to iterate over the raw logs and unpacked data for BuyPriceAccepted events raised by the Foundation contract.
type FoundationBuyPriceAcceptedIterator struct {
	Event *FoundationBuyPriceAccepted // Event containing the contract specifics and raw log

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
func (it *FoundationBuyPriceAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FoundationBuyPriceAccepted)
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
		it.Event = new(FoundationBuyPriceAccepted)
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
func (it *FoundationBuyPriceAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FoundationBuyPriceAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FoundationBuyPriceAccepted represents a BuyPriceAccepted event raised by the Foundation contract.
type FoundationBuyPriceAccepted struct {
	NftContract common.Address
	TokenId     *big.Int
	Seller      common.Address
	Buyer       common.Address
	TotalFees   *big.Int
	CreatorRev  *big.Int
	SellerRev   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBuyPriceAccepted is a free log retrieval operation binding the contract event 0xd28c0a7dd63bc853a4e36306655da9f8c0b29ff9d0605bb976ae420e46a99930.
//
// Solidity: event BuyPriceAccepted(address indexed nftContract, uint256 indexed tokenId, address indexed seller, address buyer, uint256 totalFees, uint256 creatorRev, uint256 sellerRev)
func (_Foundation *FoundationFilterer) FilterBuyPriceAccepted(opts *bind.FilterOpts, nftContract []common.Address, tokenId []*big.Int, seller []common.Address) (*FoundationBuyPriceAcceptedIterator, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _Foundation.contract.FilterLogs(opts, "BuyPriceAccepted", nftContractRule, tokenIdRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return &FoundationBuyPriceAcceptedIterator{contract: _Foundation.contract, event: "BuyPriceAccepted", logs: logs, sub: sub}, nil
}

// WatchBuyPriceAccepted is a free log subscription operation binding the contract event 0xd28c0a7dd63bc853a4e36306655da9f8c0b29ff9d0605bb976ae420e46a99930.
//
// Solidity: event BuyPriceAccepted(address indexed nftContract, uint256 indexed tokenId, address indexed seller, address buyer, uint256 totalFees, uint256 creatorRev, uint256 sellerRev)
func (_Foundation *FoundationFilterer) WatchBuyPriceAccepted(opts *bind.WatchOpts, sink chan<- *FoundationBuyPriceAccepted, nftContract []common.Address, tokenId []*big.Int, seller []common.Address) (event.Subscription, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _Foundation.contract.WatchLogs(opts, "BuyPriceAccepted", nftContractRule, tokenIdRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FoundationBuyPriceAccepted)
				if err := _Foundation.contract.UnpackLog(event, "BuyPriceAccepted", log); err != nil {
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

// ParseBuyPriceAccepted is a log parse operation binding the contract event 0xd28c0a7dd63bc853a4e36306655da9f8c0b29ff9d0605bb976ae420e46a99930.
//
// Solidity: event BuyPriceAccepted(address indexed nftContract, uint256 indexed tokenId, address indexed seller, address buyer, uint256 totalFees, uint256 creatorRev, uint256 sellerRev)
func (_Foundation *FoundationFilterer) ParseBuyPriceAccepted(log types.Log) (*FoundationBuyPriceAccepted, error) {
	event := new(FoundationBuyPriceAccepted)
	if err := _Foundation.contract.UnpackLog(event, "BuyPriceAccepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FoundationBuyPriceCanceledIterator is returned from FilterBuyPriceCanceled and is used to iterate over the raw logs and unpacked data for BuyPriceCanceled events raised by the Foundation contract.
type FoundationBuyPriceCanceledIterator struct {
	Event *FoundationBuyPriceCanceled // Event containing the contract specifics and raw log

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
func (it *FoundationBuyPriceCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FoundationBuyPriceCanceled)
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
		it.Event = new(FoundationBuyPriceCanceled)
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
func (it *FoundationBuyPriceCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FoundationBuyPriceCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FoundationBuyPriceCanceled represents a BuyPriceCanceled event raised by the Foundation contract.
type FoundationBuyPriceCanceled struct {
	NftContract common.Address
	TokenId     *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBuyPriceCanceled is a free log retrieval operation binding the contract event 0x70c7877531c04c7d9caa8a7eca127384f04e8a6ee58b63f778ce5401d8bcae41.
//
// Solidity: event BuyPriceCanceled(address indexed nftContract, uint256 indexed tokenId)
func (_Foundation *FoundationFilterer) FilterBuyPriceCanceled(opts *bind.FilterOpts, nftContract []common.Address, tokenId []*big.Int) (*FoundationBuyPriceCanceledIterator, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Foundation.contract.FilterLogs(opts, "BuyPriceCanceled", nftContractRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &FoundationBuyPriceCanceledIterator{contract: _Foundation.contract, event: "BuyPriceCanceled", logs: logs, sub: sub}, nil
}

// WatchBuyPriceCanceled is a free log subscription operation binding the contract event 0x70c7877531c04c7d9caa8a7eca127384f04e8a6ee58b63f778ce5401d8bcae41.
//
// Solidity: event BuyPriceCanceled(address indexed nftContract, uint256 indexed tokenId)
func (_Foundation *FoundationFilterer) WatchBuyPriceCanceled(opts *bind.WatchOpts, sink chan<- *FoundationBuyPriceCanceled, nftContract []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Foundation.contract.WatchLogs(opts, "BuyPriceCanceled", nftContractRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FoundationBuyPriceCanceled)
				if err := _Foundation.contract.UnpackLog(event, "BuyPriceCanceled", log); err != nil {
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

// ParseBuyPriceCanceled is a log parse operation binding the contract event 0x70c7877531c04c7d9caa8a7eca127384f04e8a6ee58b63f778ce5401d8bcae41.
//
// Solidity: event BuyPriceCanceled(address indexed nftContract, uint256 indexed tokenId)
func (_Foundation *FoundationFilterer) ParseBuyPriceCanceled(log types.Log) (*FoundationBuyPriceCanceled, error) {
	event := new(FoundationBuyPriceCanceled)
	if err := _Foundation.contract.UnpackLog(event, "BuyPriceCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FoundationBuyPriceInvalidatedIterator is returned from FilterBuyPriceInvalidated and is used to iterate over the raw logs and unpacked data for BuyPriceInvalidated events raised by the Foundation contract.
type FoundationBuyPriceInvalidatedIterator struct {
	Event *FoundationBuyPriceInvalidated // Event containing the contract specifics and raw log

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
func (it *FoundationBuyPriceInvalidatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FoundationBuyPriceInvalidated)
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
		it.Event = new(FoundationBuyPriceInvalidated)
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
func (it *FoundationBuyPriceInvalidatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FoundationBuyPriceInvalidatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FoundationBuyPriceInvalidated represents a BuyPriceInvalidated event raised by the Foundation contract.
type FoundationBuyPriceInvalidated struct {
	NftContract common.Address
	TokenId     *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBuyPriceInvalidated is a free log retrieval operation binding the contract event 0xaa6271d89a385571e237d3e7254ccc7c09f68055e6e9b410ed08233a8b9a05cf.
//
// Solidity: event BuyPriceInvalidated(address indexed nftContract, uint256 indexed tokenId)
func (_Foundation *FoundationFilterer) FilterBuyPriceInvalidated(opts *bind.FilterOpts, nftContract []common.Address, tokenId []*big.Int) (*FoundationBuyPriceInvalidatedIterator, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Foundation.contract.FilterLogs(opts, "BuyPriceInvalidated", nftContractRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &FoundationBuyPriceInvalidatedIterator{contract: _Foundation.contract, event: "BuyPriceInvalidated", logs: logs, sub: sub}, nil
}

// WatchBuyPriceInvalidated is a free log subscription operation binding the contract event 0xaa6271d89a385571e237d3e7254ccc7c09f68055e6e9b410ed08233a8b9a05cf.
//
// Solidity: event BuyPriceInvalidated(address indexed nftContract, uint256 indexed tokenId)
func (_Foundation *FoundationFilterer) WatchBuyPriceInvalidated(opts *bind.WatchOpts, sink chan<- *FoundationBuyPriceInvalidated, nftContract []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Foundation.contract.WatchLogs(opts, "BuyPriceInvalidated", nftContractRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FoundationBuyPriceInvalidated)
				if err := _Foundation.contract.UnpackLog(event, "BuyPriceInvalidated", log); err != nil {
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

// ParseBuyPriceInvalidated is a log parse operation binding the contract event 0xaa6271d89a385571e237d3e7254ccc7c09f68055e6e9b410ed08233a8b9a05cf.
//
// Solidity: event BuyPriceInvalidated(address indexed nftContract, uint256 indexed tokenId)
func (_Foundation *FoundationFilterer) ParseBuyPriceInvalidated(log types.Log) (*FoundationBuyPriceInvalidated, error) {
	event := new(FoundationBuyPriceInvalidated)
	if err := _Foundation.contract.UnpackLog(event, "BuyPriceInvalidated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FoundationBuyPriceSetIterator is returned from FilterBuyPriceSet and is used to iterate over the raw logs and unpacked data for BuyPriceSet events raised by the Foundation contract.
type FoundationBuyPriceSetIterator struct {
	Event *FoundationBuyPriceSet // Event containing the contract specifics and raw log

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
func (it *FoundationBuyPriceSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FoundationBuyPriceSet)
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
		it.Event = new(FoundationBuyPriceSet)
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
func (it *FoundationBuyPriceSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FoundationBuyPriceSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FoundationBuyPriceSet represents a BuyPriceSet event raised by the Foundation contract.
type FoundationBuyPriceSet struct {
	NftContract common.Address
	TokenId     *big.Int
	Seller      common.Address
	Price       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBuyPriceSet is a free log retrieval operation binding the contract event 0xfcc77ea8bdcce862f43b7fb00fe6b0eb90d6aeead27d3800d9257cf7a05f9d96.
//
// Solidity: event BuyPriceSet(address indexed nftContract, uint256 indexed tokenId, address indexed seller, uint256 price)
func (_Foundation *FoundationFilterer) FilterBuyPriceSet(opts *bind.FilterOpts, nftContract []common.Address, tokenId []*big.Int, seller []common.Address) (*FoundationBuyPriceSetIterator, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _Foundation.contract.FilterLogs(opts, "BuyPriceSet", nftContractRule, tokenIdRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return &FoundationBuyPriceSetIterator{contract: _Foundation.contract, event: "BuyPriceSet", logs: logs, sub: sub}, nil
}

// WatchBuyPriceSet is a free log subscription operation binding the contract event 0xfcc77ea8bdcce862f43b7fb00fe6b0eb90d6aeead27d3800d9257cf7a05f9d96.
//
// Solidity: event BuyPriceSet(address indexed nftContract, uint256 indexed tokenId, address indexed seller, uint256 price)
func (_Foundation *FoundationFilterer) WatchBuyPriceSet(opts *bind.WatchOpts, sink chan<- *FoundationBuyPriceSet, nftContract []common.Address, tokenId []*big.Int, seller []common.Address) (event.Subscription, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _Foundation.contract.WatchLogs(opts, "BuyPriceSet", nftContractRule, tokenIdRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FoundationBuyPriceSet)
				if err := _Foundation.contract.UnpackLog(event, "BuyPriceSet", log); err != nil {
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

// ParseBuyPriceSet is a log parse operation binding the contract event 0xfcc77ea8bdcce862f43b7fb00fe6b0eb90d6aeead27d3800d9257cf7a05f9d96.
//
// Solidity: event BuyPriceSet(address indexed nftContract, uint256 indexed tokenId, address indexed seller, uint256 price)
func (_Foundation *FoundationFilterer) ParseBuyPriceSet(log types.Log) (*FoundationBuyPriceSet, error) {
	event := new(FoundationBuyPriceSet)
	if err := _Foundation.contract.UnpackLog(event, "BuyPriceSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FoundationBuyReferralPaidIterator is returned from FilterBuyReferralPaid and is used to iterate over the raw logs and unpacked data for BuyReferralPaid events raised by the Foundation contract.
type FoundationBuyReferralPaidIterator struct {
	Event *FoundationBuyReferralPaid // Event containing the contract specifics and raw log

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
func (it *FoundationBuyReferralPaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FoundationBuyReferralPaid)
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
		it.Event = new(FoundationBuyReferralPaid)
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
func (it *FoundationBuyReferralPaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FoundationBuyReferralPaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FoundationBuyReferralPaid represents a BuyReferralPaid event raised by the Foundation contract.
type FoundationBuyReferralPaid struct {
	NftContract          common.Address
	TokenId              *big.Int
	BuyReferrer          common.Address
	BuyReferrerFee       *big.Int
	BuyReferrerSellerFee *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterBuyReferralPaid is a free log retrieval operation binding the contract event 0x141b92fd9766c80ab120598ea2f6be9802470ec59b5446dd9bf46214ead8d08e.
//
// Solidity: event BuyReferralPaid(address indexed nftContract, uint256 indexed tokenId, address buyReferrer, uint256 buyReferrerFee, uint256 buyReferrerSellerFee)
func (_Foundation *FoundationFilterer) FilterBuyReferralPaid(opts *bind.FilterOpts, nftContract []common.Address, tokenId []*big.Int) (*FoundationBuyReferralPaidIterator, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Foundation.contract.FilterLogs(opts, "BuyReferralPaid", nftContractRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &FoundationBuyReferralPaidIterator{contract: _Foundation.contract, event: "BuyReferralPaid", logs: logs, sub: sub}, nil
}

// WatchBuyReferralPaid is a free log subscription operation binding the contract event 0x141b92fd9766c80ab120598ea2f6be9802470ec59b5446dd9bf46214ead8d08e.
//
// Solidity: event BuyReferralPaid(address indexed nftContract, uint256 indexed tokenId, address buyReferrer, uint256 buyReferrerFee, uint256 buyReferrerSellerFee)
func (_Foundation *FoundationFilterer) WatchBuyReferralPaid(opts *bind.WatchOpts, sink chan<- *FoundationBuyReferralPaid, nftContract []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Foundation.contract.WatchLogs(opts, "BuyReferralPaid", nftContractRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FoundationBuyReferralPaid)
				if err := _Foundation.contract.UnpackLog(event, "BuyReferralPaid", log); err != nil {
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

// ParseBuyReferralPaid is a log parse operation binding the contract event 0x141b92fd9766c80ab120598ea2f6be9802470ec59b5446dd9bf46214ead8d08e.
//
// Solidity: event BuyReferralPaid(address indexed nftContract, uint256 indexed tokenId, address buyReferrer, uint256 buyReferrerFee, uint256 buyReferrerSellerFee)
func (_Foundation *FoundationFilterer) ParseBuyReferralPaid(log types.Log) (*FoundationBuyReferralPaid, error) {
	event := new(FoundationBuyReferralPaid)
	if err := _Foundation.contract.UnpackLog(event, "BuyReferralPaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FoundationExhibitionCreatedIterator is returned from FilterExhibitionCreated and is used to iterate over the raw logs and unpacked data for ExhibitionCreated events raised by the Foundation contract.
type FoundationExhibitionCreatedIterator struct {
	Event *FoundationExhibitionCreated // Event containing the contract specifics and raw log

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
func (it *FoundationExhibitionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FoundationExhibitionCreated)
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
		it.Event = new(FoundationExhibitionCreated)
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
func (it *FoundationExhibitionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FoundationExhibitionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FoundationExhibitionCreated represents a ExhibitionCreated event raised by the Foundation contract.
type FoundationExhibitionCreated struct {
	ExhibitionId          *big.Int
	Curator               common.Address
	Name                  string
	TakeRateInBasisPoints uint16
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterExhibitionCreated is a free log retrieval operation binding the contract event 0x9eee3ce0e6f7eeabd69ecf363898e9f490dbfda9ad953e1019a2c6aeceb4a7ef.
//
// Solidity: event ExhibitionCreated(uint256 indexed exhibitionId, address indexed curator, string name, uint16 takeRateInBasisPoints)
func (_Foundation *FoundationFilterer) FilterExhibitionCreated(opts *bind.FilterOpts, exhibitionId []*big.Int, curator []common.Address) (*FoundationExhibitionCreatedIterator, error) {

	var exhibitionIdRule []interface{}
	for _, exhibitionIdItem := range exhibitionId {
		exhibitionIdRule = append(exhibitionIdRule, exhibitionIdItem)
	}
	var curatorRule []interface{}
	for _, curatorItem := range curator {
		curatorRule = append(curatorRule, curatorItem)
	}

	logs, sub, err := _Foundation.contract.FilterLogs(opts, "ExhibitionCreated", exhibitionIdRule, curatorRule)
	if err != nil {
		return nil, err
	}
	return &FoundationExhibitionCreatedIterator{contract: _Foundation.contract, event: "ExhibitionCreated", logs: logs, sub: sub}, nil
}

// WatchExhibitionCreated is a free log subscription operation binding the contract event 0x9eee3ce0e6f7eeabd69ecf363898e9f490dbfda9ad953e1019a2c6aeceb4a7ef.
//
// Solidity: event ExhibitionCreated(uint256 indexed exhibitionId, address indexed curator, string name, uint16 takeRateInBasisPoints)
func (_Foundation *FoundationFilterer) WatchExhibitionCreated(opts *bind.WatchOpts, sink chan<- *FoundationExhibitionCreated, exhibitionId []*big.Int, curator []common.Address) (event.Subscription, error) {

	var exhibitionIdRule []interface{}
	for _, exhibitionIdItem := range exhibitionId {
		exhibitionIdRule = append(exhibitionIdRule, exhibitionIdItem)
	}
	var curatorRule []interface{}
	for _, curatorItem := range curator {
		curatorRule = append(curatorRule, curatorItem)
	}

	logs, sub, err := _Foundation.contract.WatchLogs(opts, "ExhibitionCreated", exhibitionIdRule, curatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FoundationExhibitionCreated)
				if err := _Foundation.contract.UnpackLog(event, "ExhibitionCreated", log); err != nil {
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

// ParseExhibitionCreated is a log parse operation binding the contract event 0x9eee3ce0e6f7eeabd69ecf363898e9f490dbfda9ad953e1019a2c6aeceb4a7ef.
//
// Solidity: event ExhibitionCreated(uint256 indexed exhibitionId, address indexed curator, string name, uint16 takeRateInBasisPoints)
func (_Foundation *FoundationFilterer) ParseExhibitionCreated(log types.Log) (*FoundationExhibitionCreated, error) {
	event := new(FoundationExhibitionCreated)
	if err := _Foundation.contract.UnpackLog(event, "ExhibitionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FoundationExhibitionDeletedIterator is returned from FilterExhibitionDeleted and is used to iterate over the raw logs and unpacked data for ExhibitionDeleted events raised by the Foundation contract.
type FoundationExhibitionDeletedIterator struct {
	Event *FoundationExhibitionDeleted // Event containing the contract specifics and raw log

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
func (it *FoundationExhibitionDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FoundationExhibitionDeleted)
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
		it.Event = new(FoundationExhibitionDeleted)
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
func (it *FoundationExhibitionDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FoundationExhibitionDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FoundationExhibitionDeleted represents a ExhibitionDeleted event raised by the Foundation contract.
type FoundationExhibitionDeleted struct {
	ExhibitionId *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterExhibitionDeleted is a free log retrieval operation binding the contract event 0x2a9aeaf340ca0da469c1f7e3d513c0e6c9cd287016f29d257a4ef70e13dc441c.
//
// Solidity: event ExhibitionDeleted(uint256 indexed exhibitionId)
func (_Foundation *FoundationFilterer) FilterExhibitionDeleted(opts *bind.FilterOpts, exhibitionId []*big.Int) (*FoundationExhibitionDeletedIterator, error) {

	var exhibitionIdRule []interface{}
	for _, exhibitionIdItem := range exhibitionId {
		exhibitionIdRule = append(exhibitionIdRule, exhibitionIdItem)
	}

	logs, sub, err := _Foundation.contract.FilterLogs(opts, "ExhibitionDeleted", exhibitionIdRule)
	if err != nil {
		return nil, err
	}
	return &FoundationExhibitionDeletedIterator{contract: _Foundation.contract, event: "ExhibitionDeleted", logs: logs, sub: sub}, nil
}

// WatchExhibitionDeleted is a free log subscription operation binding the contract event 0x2a9aeaf340ca0da469c1f7e3d513c0e6c9cd287016f29d257a4ef70e13dc441c.
//
// Solidity: event ExhibitionDeleted(uint256 indexed exhibitionId)
func (_Foundation *FoundationFilterer) WatchExhibitionDeleted(opts *bind.WatchOpts, sink chan<- *FoundationExhibitionDeleted, exhibitionId []*big.Int) (event.Subscription, error) {

	var exhibitionIdRule []interface{}
	for _, exhibitionIdItem := range exhibitionId {
		exhibitionIdRule = append(exhibitionIdRule, exhibitionIdItem)
	}

	logs, sub, err := _Foundation.contract.WatchLogs(opts, "ExhibitionDeleted", exhibitionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FoundationExhibitionDeleted)
				if err := _Foundation.contract.UnpackLog(event, "ExhibitionDeleted", log); err != nil {
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

// ParseExhibitionDeleted is a log parse operation binding the contract event 0x2a9aeaf340ca0da469c1f7e3d513c0e6c9cd287016f29d257a4ef70e13dc441c.
//
// Solidity: event ExhibitionDeleted(uint256 indexed exhibitionId)
func (_Foundation *FoundationFilterer) ParseExhibitionDeleted(log types.Log) (*FoundationExhibitionDeleted, error) {
	event := new(FoundationExhibitionDeleted)
	if err := _Foundation.contract.UnpackLog(event, "ExhibitionDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FoundationInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Foundation contract.
type FoundationInitializedIterator struct {
	Event *FoundationInitialized // Event containing the contract specifics and raw log

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
func (it *FoundationInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FoundationInitialized)
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
		it.Event = new(FoundationInitialized)
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
func (it *FoundationInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FoundationInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FoundationInitialized represents a Initialized event raised by the Foundation contract.
type FoundationInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Foundation *FoundationFilterer) FilterInitialized(opts *bind.FilterOpts) (*FoundationInitializedIterator, error) {

	logs, sub, err := _Foundation.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &FoundationInitializedIterator{contract: _Foundation.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Foundation *FoundationFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *FoundationInitialized) (event.Subscription, error) {

	logs, sub, err := _Foundation.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FoundationInitialized)
				if err := _Foundation.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Foundation *FoundationFilterer) ParseInitialized(log types.Log) (*FoundationInitialized, error) {
	event := new(FoundationInitialized)
	if err := _Foundation.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FoundationNftAddedToExhibitionIterator is returned from FilterNftAddedToExhibition and is used to iterate over the raw logs and unpacked data for NftAddedToExhibition events raised by the Foundation contract.
type FoundationNftAddedToExhibitionIterator struct {
	Event *FoundationNftAddedToExhibition // Event containing the contract specifics and raw log

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
func (it *FoundationNftAddedToExhibitionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FoundationNftAddedToExhibition)
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
		it.Event = new(FoundationNftAddedToExhibition)
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
func (it *FoundationNftAddedToExhibitionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FoundationNftAddedToExhibitionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FoundationNftAddedToExhibition represents a NftAddedToExhibition event raised by the Foundation contract.
type FoundationNftAddedToExhibition struct {
	NftContract  common.Address
	TokenId      *big.Int
	ExhibitionId *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNftAddedToExhibition is a free log retrieval operation binding the contract event 0xb17e0c916df75a12480835f00b3927cb871bbe00bacf819f81a1d92f9ff7f38d.
//
// Solidity: event NftAddedToExhibition(address indexed nftContract, uint256 indexed tokenId, uint256 indexed exhibitionId)
func (_Foundation *FoundationFilterer) FilterNftAddedToExhibition(opts *bind.FilterOpts, nftContract []common.Address, tokenId []*big.Int, exhibitionId []*big.Int) (*FoundationNftAddedToExhibitionIterator, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var exhibitionIdRule []interface{}
	for _, exhibitionIdItem := range exhibitionId {
		exhibitionIdRule = append(exhibitionIdRule, exhibitionIdItem)
	}

	logs, sub, err := _Foundation.contract.FilterLogs(opts, "NftAddedToExhibition", nftContractRule, tokenIdRule, exhibitionIdRule)
	if err != nil {
		return nil, err
	}
	return &FoundationNftAddedToExhibitionIterator{contract: _Foundation.contract, event: "NftAddedToExhibition", logs: logs, sub: sub}, nil
}

// WatchNftAddedToExhibition is a free log subscription operation binding the contract event 0xb17e0c916df75a12480835f00b3927cb871bbe00bacf819f81a1d92f9ff7f38d.
//
// Solidity: event NftAddedToExhibition(address indexed nftContract, uint256 indexed tokenId, uint256 indexed exhibitionId)
func (_Foundation *FoundationFilterer) WatchNftAddedToExhibition(opts *bind.WatchOpts, sink chan<- *FoundationNftAddedToExhibition, nftContract []common.Address, tokenId []*big.Int, exhibitionId []*big.Int) (event.Subscription, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var exhibitionIdRule []interface{}
	for _, exhibitionIdItem := range exhibitionId {
		exhibitionIdRule = append(exhibitionIdRule, exhibitionIdItem)
	}

	logs, sub, err := _Foundation.contract.WatchLogs(opts, "NftAddedToExhibition", nftContractRule, tokenIdRule, exhibitionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FoundationNftAddedToExhibition)
				if err := _Foundation.contract.UnpackLog(event, "NftAddedToExhibition", log); err != nil {
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

// ParseNftAddedToExhibition is a log parse operation binding the contract event 0xb17e0c916df75a12480835f00b3927cb871bbe00bacf819f81a1d92f9ff7f38d.
//
// Solidity: event NftAddedToExhibition(address indexed nftContract, uint256 indexed tokenId, uint256 indexed exhibitionId)
func (_Foundation *FoundationFilterer) ParseNftAddedToExhibition(log types.Log) (*FoundationNftAddedToExhibition, error) {
	event := new(FoundationNftAddedToExhibition)
	if err := _Foundation.contract.UnpackLog(event, "NftAddedToExhibition", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FoundationNftRemovedFromExhibitionIterator is returned from FilterNftRemovedFromExhibition and is used to iterate over the raw logs and unpacked data for NftRemovedFromExhibition events raised by the Foundation contract.
type FoundationNftRemovedFromExhibitionIterator struct {
	Event *FoundationNftRemovedFromExhibition // Event containing the contract specifics and raw log

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
func (it *FoundationNftRemovedFromExhibitionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FoundationNftRemovedFromExhibition)
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
		it.Event = new(FoundationNftRemovedFromExhibition)
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
func (it *FoundationNftRemovedFromExhibitionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FoundationNftRemovedFromExhibitionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FoundationNftRemovedFromExhibition represents a NftRemovedFromExhibition event raised by the Foundation contract.
type FoundationNftRemovedFromExhibition struct {
	NftContract  common.Address
	TokenId      *big.Int
	ExhibitionId *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNftRemovedFromExhibition is a free log retrieval operation binding the contract event 0x2ea2946ee16c4a1d0ec58464194022e54432a6d7db359835ddf283555f2c8eee.
//
// Solidity: event NftRemovedFromExhibition(address indexed nftContract, uint256 indexed tokenId, uint256 indexed exhibitionId)
func (_Foundation *FoundationFilterer) FilterNftRemovedFromExhibition(opts *bind.FilterOpts, nftContract []common.Address, tokenId []*big.Int, exhibitionId []*big.Int) (*FoundationNftRemovedFromExhibitionIterator, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var exhibitionIdRule []interface{}
	for _, exhibitionIdItem := range exhibitionId {
		exhibitionIdRule = append(exhibitionIdRule, exhibitionIdItem)
	}

	logs, sub, err := _Foundation.contract.FilterLogs(opts, "NftRemovedFromExhibition", nftContractRule, tokenIdRule, exhibitionIdRule)
	if err != nil {
		return nil, err
	}
	return &FoundationNftRemovedFromExhibitionIterator{contract: _Foundation.contract, event: "NftRemovedFromExhibition", logs: logs, sub: sub}, nil
}

// WatchNftRemovedFromExhibition is a free log subscription operation binding the contract event 0x2ea2946ee16c4a1d0ec58464194022e54432a6d7db359835ddf283555f2c8eee.
//
// Solidity: event NftRemovedFromExhibition(address indexed nftContract, uint256 indexed tokenId, uint256 indexed exhibitionId)
func (_Foundation *FoundationFilterer) WatchNftRemovedFromExhibition(opts *bind.WatchOpts, sink chan<- *FoundationNftRemovedFromExhibition, nftContract []common.Address, tokenId []*big.Int, exhibitionId []*big.Int) (event.Subscription, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var exhibitionIdRule []interface{}
	for _, exhibitionIdItem := range exhibitionId {
		exhibitionIdRule = append(exhibitionIdRule, exhibitionIdItem)
	}

	logs, sub, err := _Foundation.contract.WatchLogs(opts, "NftRemovedFromExhibition", nftContractRule, tokenIdRule, exhibitionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FoundationNftRemovedFromExhibition)
				if err := _Foundation.contract.UnpackLog(event, "NftRemovedFromExhibition", log); err != nil {
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

// ParseNftRemovedFromExhibition is a log parse operation binding the contract event 0x2ea2946ee16c4a1d0ec58464194022e54432a6d7db359835ddf283555f2c8eee.
//
// Solidity: event NftRemovedFromExhibition(address indexed nftContract, uint256 indexed tokenId, uint256 indexed exhibitionId)
func (_Foundation *FoundationFilterer) ParseNftRemovedFromExhibition(log types.Log) (*FoundationNftRemovedFromExhibition, error) {
	event := new(FoundationNftRemovedFromExhibition)
	if err := _Foundation.contract.UnpackLog(event, "NftRemovedFromExhibition", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FoundationOfferAcceptedIterator is returned from FilterOfferAccepted and is used to iterate over the raw logs and unpacked data for OfferAccepted events raised by the Foundation contract.
type FoundationOfferAcceptedIterator struct {
	Event *FoundationOfferAccepted // Event containing the contract specifics and raw log

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
func (it *FoundationOfferAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FoundationOfferAccepted)
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
		it.Event = new(FoundationOfferAccepted)
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
func (it *FoundationOfferAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FoundationOfferAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FoundationOfferAccepted represents a OfferAccepted event raised by the Foundation contract.
type FoundationOfferAccepted struct {
	NftContract common.Address
	TokenId     *big.Int
	Buyer       common.Address
	Seller      common.Address
	TotalFees   *big.Int
	CreatorRev  *big.Int
	SellerRev   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOfferAccepted is a free log retrieval operation binding the contract event 0x1cb8adb37d6d35e94cd0695ca39895b84371864713f5ca7eada52af9ff23744b.
//
// Solidity: event OfferAccepted(address indexed nftContract, uint256 indexed tokenId, address indexed buyer, address seller, uint256 totalFees, uint256 creatorRev, uint256 sellerRev)
func (_Foundation *FoundationFilterer) FilterOfferAccepted(opts *bind.FilterOpts, nftContract []common.Address, tokenId []*big.Int, buyer []common.Address) (*FoundationOfferAcceptedIterator, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Foundation.contract.FilterLogs(opts, "OfferAccepted", nftContractRule, tokenIdRule, buyerRule)
	if err != nil {
		return nil, err
	}
	return &FoundationOfferAcceptedIterator{contract: _Foundation.contract, event: "OfferAccepted", logs: logs, sub: sub}, nil
}

// WatchOfferAccepted is a free log subscription operation binding the contract event 0x1cb8adb37d6d35e94cd0695ca39895b84371864713f5ca7eada52af9ff23744b.
//
// Solidity: event OfferAccepted(address indexed nftContract, uint256 indexed tokenId, address indexed buyer, address seller, uint256 totalFees, uint256 creatorRev, uint256 sellerRev)
func (_Foundation *FoundationFilterer) WatchOfferAccepted(opts *bind.WatchOpts, sink chan<- *FoundationOfferAccepted, nftContract []common.Address, tokenId []*big.Int, buyer []common.Address) (event.Subscription, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Foundation.contract.WatchLogs(opts, "OfferAccepted", nftContractRule, tokenIdRule, buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FoundationOfferAccepted)
				if err := _Foundation.contract.UnpackLog(event, "OfferAccepted", log); err != nil {
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

// ParseOfferAccepted is a log parse operation binding the contract event 0x1cb8adb37d6d35e94cd0695ca39895b84371864713f5ca7eada52af9ff23744b.
//
// Solidity: event OfferAccepted(address indexed nftContract, uint256 indexed tokenId, address indexed buyer, address seller, uint256 totalFees, uint256 creatorRev, uint256 sellerRev)
func (_Foundation *FoundationFilterer) ParseOfferAccepted(log types.Log) (*FoundationOfferAccepted, error) {
	event := new(FoundationOfferAccepted)
	if err := _Foundation.contract.UnpackLog(event, "OfferAccepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FoundationOfferInvalidatedIterator is returned from FilterOfferInvalidated and is used to iterate over the raw logs and unpacked data for OfferInvalidated events raised by the Foundation contract.
type FoundationOfferInvalidatedIterator struct {
	Event *FoundationOfferInvalidated // Event containing the contract specifics and raw log

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
func (it *FoundationOfferInvalidatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FoundationOfferInvalidated)
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
		it.Event = new(FoundationOfferInvalidated)
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
func (it *FoundationOfferInvalidatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FoundationOfferInvalidatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FoundationOfferInvalidated represents a OfferInvalidated event raised by the Foundation contract.
type FoundationOfferInvalidated struct {
	NftContract common.Address
	TokenId     *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOfferInvalidated is a free log retrieval operation binding the contract event 0x30c264456cbd17f5f67d7534654161414f34c0e6cc1b7500e169b7a7aea4afc0.
//
// Solidity: event OfferInvalidated(address indexed nftContract, uint256 indexed tokenId)
func (_Foundation *FoundationFilterer) FilterOfferInvalidated(opts *bind.FilterOpts, nftContract []common.Address, tokenId []*big.Int) (*FoundationOfferInvalidatedIterator, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Foundation.contract.FilterLogs(opts, "OfferInvalidated", nftContractRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &FoundationOfferInvalidatedIterator{contract: _Foundation.contract, event: "OfferInvalidated", logs: logs, sub: sub}, nil
}

// WatchOfferInvalidated is a free log subscription operation binding the contract event 0x30c264456cbd17f5f67d7534654161414f34c0e6cc1b7500e169b7a7aea4afc0.
//
// Solidity: event OfferInvalidated(address indexed nftContract, uint256 indexed tokenId)
func (_Foundation *FoundationFilterer) WatchOfferInvalidated(opts *bind.WatchOpts, sink chan<- *FoundationOfferInvalidated, nftContract []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Foundation.contract.WatchLogs(opts, "OfferInvalidated", nftContractRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FoundationOfferInvalidated)
				if err := _Foundation.contract.UnpackLog(event, "OfferInvalidated", log); err != nil {
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

// ParseOfferInvalidated is a log parse operation binding the contract event 0x30c264456cbd17f5f67d7534654161414f34c0e6cc1b7500e169b7a7aea4afc0.
//
// Solidity: event OfferInvalidated(address indexed nftContract, uint256 indexed tokenId)
func (_Foundation *FoundationFilterer) ParseOfferInvalidated(log types.Log) (*FoundationOfferInvalidated, error) {
	event := new(FoundationOfferInvalidated)
	if err := _Foundation.contract.UnpackLog(event, "OfferInvalidated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FoundationOfferMadeIterator is returned from FilterOfferMade and is used to iterate over the raw logs and unpacked data for OfferMade events raised by the Foundation contract.
type FoundationOfferMadeIterator struct {
	Event *FoundationOfferMade // Event containing the contract specifics and raw log

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
func (it *FoundationOfferMadeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FoundationOfferMade)
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
		it.Event = new(FoundationOfferMade)
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
func (it *FoundationOfferMadeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FoundationOfferMadeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FoundationOfferMade represents a OfferMade event raised by the Foundation contract.
type FoundationOfferMade struct {
	NftContract common.Address
	TokenId     *big.Int
	Buyer       common.Address
	Amount      *big.Int
	Expiration  *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOfferMade is a free log retrieval operation binding the contract event 0x00ce0a712e4e277ac7b34942865f0de7a5629dffe0539b70423ad5ff1ed6ab42.
//
// Solidity: event OfferMade(address indexed nftContract, uint256 indexed tokenId, address indexed buyer, uint256 amount, uint256 expiration)
func (_Foundation *FoundationFilterer) FilterOfferMade(opts *bind.FilterOpts, nftContract []common.Address, tokenId []*big.Int, buyer []common.Address) (*FoundationOfferMadeIterator, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Foundation.contract.FilterLogs(opts, "OfferMade", nftContractRule, tokenIdRule, buyerRule)
	if err != nil {
		return nil, err
	}
	return &FoundationOfferMadeIterator{contract: _Foundation.contract, event: "OfferMade", logs: logs, sub: sub}, nil
}

// WatchOfferMade is a free log subscription operation binding the contract event 0x00ce0a712e4e277ac7b34942865f0de7a5629dffe0539b70423ad5ff1ed6ab42.
//
// Solidity: event OfferMade(address indexed nftContract, uint256 indexed tokenId, address indexed buyer, uint256 amount, uint256 expiration)
func (_Foundation *FoundationFilterer) WatchOfferMade(opts *bind.WatchOpts, sink chan<- *FoundationOfferMade, nftContract []common.Address, tokenId []*big.Int, buyer []common.Address) (event.Subscription, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Foundation.contract.WatchLogs(opts, "OfferMade", nftContractRule, tokenIdRule, buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FoundationOfferMade)
				if err := _Foundation.contract.UnpackLog(event, "OfferMade", log); err != nil {
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

// ParseOfferMade is a log parse operation binding the contract event 0x00ce0a712e4e277ac7b34942865f0de7a5629dffe0539b70423ad5ff1ed6ab42.
//
// Solidity: event OfferMade(address indexed nftContract, uint256 indexed tokenId, address indexed buyer, uint256 amount, uint256 expiration)
func (_Foundation *FoundationFilterer) ParseOfferMade(log types.Log) (*FoundationOfferMade, error) {
	event := new(FoundationOfferMade)
	if err := _Foundation.contract.UnpackLog(event, "OfferMade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FoundationReserveAuctionBidPlacedIterator is returned from FilterReserveAuctionBidPlaced and is used to iterate over the raw logs and unpacked data for ReserveAuctionBidPlaced events raised by the Foundation contract.
type FoundationReserveAuctionBidPlacedIterator struct {
	Event *FoundationReserveAuctionBidPlaced // Event containing the contract specifics and raw log

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
func (it *FoundationReserveAuctionBidPlacedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FoundationReserveAuctionBidPlaced)
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
		it.Event = new(FoundationReserveAuctionBidPlaced)
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
func (it *FoundationReserveAuctionBidPlacedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FoundationReserveAuctionBidPlacedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FoundationReserveAuctionBidPlaced represents a ReserveAuctionBidPlaced event raised by the Foundation contract.
type FoundationReserveAuctionBidPlaced struct {
	AuctionId *big.Int
	Bidder    common.Address
	Amount    *big.Int
	EndTime   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterReserveAuctionBidPlaced is a free log retrieval operation binding the contract event 0x26ea3ebbda62eb1baef13e1c237dddd956c87f80b2801f2616d806d52557b121.
//
// Solidity: event ReserveAuctionBidPlaced(uint256 indexed auctionId, address indexed bidder, uint256 amount, uint256 endTime)
func (_Foundation *FoundationFilterer) FilterReserveAuctionBidPlaced(opts *bind.FilterOpts, auctionId []*big.Int, bidder []common.Address) (*FoundationReserveAuctionBidPlacedIterator, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}
	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _Foundation.contract.FilterLogs(opts, "ReserveAuctionBidPlaced", auctionIdRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return &FoundationReserveAuctionBidPlacedIterator{contract: _Foundation.contract, event: "ReserveAuctionBidPlaced", logs: logs, sub: sub}, nil
}

// WatchReserveAuctionBidPlaced is a free log subscription operation binding the contract event 0x26ea3ebbda62eb1baef13e1c237dddd956c87f80b2801f2616d806d52557b121.
//
// Solidity: event ReserveAuctionBidPlaced(uint256 indexed auctionId, address indexed bidder, uint256 amount, uint256 endTime)
func (_Foundation *FoundationFilterer) WatchReserveAuctionBidPlaced(opts *bind.WatchOpts, sink chan<- *FoundationReserveAuctionBidPlaced, auctionId []*big.Int, bidder []common.Address) (event.Subscription, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}
	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _Foundation.contract.WatchLogs(opts, "ReserveAuctionBidPlaced", auctionIdRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FoundationReserveAuctionBidPlaced)
				if err := _Foundation.contract.UnpackLog(event, "ReserveAuctionBidPlaced", log); err != nil {
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

// ParseReserveAuctionBidPlaced is a log parse operation binding the contract event 0x26ea3ebbda62eb1baef13e1c237dddd956c87f80b2801f2616d806d52557b121.
//
// Solidity: event ReserveAuctionBidPlaced(uint256 indexed auctionId, address indexed bidder, uint256 amount, uint256 endTime)
func (_Foundation *FoundationFilterer) ParseReserveAuctionBidPlaced(log types.Log) (*FoundationReserveAuctionBidPlaced, error) {
	event := new(FoundationReserveAuctionBidPlaced)
	if err := _Foundation.contract.UnpackLog(event, "ReserveAuctionBidPlaced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FoundationReserveAuctionCanceledIterator is returned from FilterReserveAuctionCanceled and is used to iterate over the raw logs and unpacked data for ReserveAuctionCanceled events raised by the Foundation contract.
type FoundationReserveAuctionCanceledIterator struct {
	Event *FoundationReserveAuctionCanceled // Event containing the contract specifics and raw log

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
func (it *FoundationReserveAuctionCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FoundationReserveAuctionCanceled)
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
		it.Event = new(FoundationReserveAuctionCanceled)
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
func (it *FoundationReserveAuctionCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FoundationReserveAuctionCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FoundationReserveAuctionCanceled represents a ReserveAuctionCanceled event raised by the Foundation contract.
type FoundationReserveAuctionCanceled struct {
	AuctionId *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterReserveAuctionCanceled is a free log retrieval operation binding the contract event 0x14b9c40404d5b41deb481f9a40b8aeb2bf4b47679b38cf757075a66ed510f7f1.
//
// Solidity: event ReserveAuctionCanceled(uint256 indexed auctionId)
func (_Foundation *FoundationFilterer) FilterReserveAuctionCanceled(opts *bind.FilterOpts, auctionId []*big.Int) (*FoundationReserveAuctionCanceledIterator, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}

	logs, sub, err := _Foundation.contract.FilterLogs(opts, "ReserveAuctionCanceled", auctionIdRule)
	if err != nil {
		return nil, err
	}
	return &FoundationReserveAuctionCanceledIterator{contract: _Foundation.contract, event: "ReserveAuctionCanceled", logs: logs, sub: sub}, nil
}

// WatchReserveAuctionCanceled is a free log subscription operation binding the contract event 0x14b9c40404d5b41deb481f9a40b8aeb2bf4b47679b38cf757075a66ed510f7f1.
//
// Solidity: event ReserveAuctionCanceled(uint256 indexed auctionId)
func (_Foundation *FoundationFilterer) WatchReserveAuctionCanceled(opts *bind.WatchOpts, sink chan<- *FoundationReserveAuctionCanceled, auctionId []*big.Int) (event.Subscription, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}

	logs, sub, err := _Foundation.contract.WatchLogs(opts, "ReserveAuctionCanceled", auctionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FoundationReserveAuctionCanceled)
				if err := _Foundation.contract.UnpackLog(event, "ReserveAuctionCanceled", log); err != nil {
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

// ParseReserveAuctionCanceled is a log parse operation binding the contract event 0x14b9c40404d5b41deb481f9a40b8aeb2bf4b47679b38cf757075a66ed510f7f1.
//
// Solidity: event ReserveAuctionCanceled(uint256 indexed auctionId)
func (_Foundation *FoundationFilterer) ParseReserveAuctionCanceled(log types.Log) (*FoundationReserveAuctionCanceled, error) {
	event := new(FoundationReserveAuctionCanceled)
	if err := _Foundation.contract.UnpackLog(event, "ReserveAuctionCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FoundationReserveAuctionCreatedIterator is returned from FilterReserveAuctionCreated and is used to iterate over the raw logs and unpacked data for ReserveAuctionCreated events raised by the Foundation contract.
type FoundationReserveAuctionCreatedIterator struct {
	Event *FoundationReserveAuctionCreated // Event containing the contract specifics and raw log

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
func (it *FoundationReserveAuctionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FoundationReserveAuctionCreated)
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
		it.Event = new(FoundationReserveAuctionCreated)
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
func (it *FoundationReserveAuctionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FoundationReserveAuctionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FoundationReserveAuctionCreated represents a ReserveAuctionCreated event raised by the Foundation contract.
type FoundationReserveAuctionCreated struct {
	Seller            common.Address
	NftContract       common.Address
	TokenId           *big.Int
	Duration          *big.Int
	ExtensionDuration *big.Int
	ReservePrice      *big.Int
	AuctionId         *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterReserveAuctionCreated is a free log retrieval operation binding the contract event 0x1062dd3b35f12b4064331244d00f40c1d4831965e4285654157a2409c6217cff.
//
// Solidity: event ReserveAuctionCreated(address indexed seller, address indexed nftContract, uint256 indexed tokenId, uint256 duration, uint256 extensionDuration, uint256 reservePrice, uint256 auctionId)
func (_Foundation *FoundationFilterer) FilterReserveAuctionCreated(opts *bind.FilterOpts, seller []common.Address, nftContract []common.Address, tokenId []*big.Int) (*FoundationReserveAuctionCreatedIterator, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Foundation.contract.FilterLogs(opts, "ReserveAuctionCreated", sellerRule, nftContractRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &FoundationReserveAuctionCreatedIterator{contract: _Foundation.contract, event: "ReserveAuctionCreated", logs: logs, sub: sub}, nil
}

// WatchReserveAuctionCreated is a free log subscription operation binding the contract event 0x1062dd3b35f12b4064331244d00f40c1d4831965e4285654157a2409c6217cff.
//
// Solidity: event ReserveAuctionCreated(address indexed seller, address indexed nftContract, uint256 indexed tokenId, uint256 duration, uint256 extensionDuration, uint256 reservePrice, uint256 auctionId)
func (_Foundation *FoundationFilterer) WatchReserveAuctionCreated(opts *bind.WatchOpts, sink chan<- *FoundationReserveAuctionCreated, seller []common.Address, nftContract []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Foundation.contract.WatchLogs(opts, "ReserveAuctionCreated", sellerRule, nftContractRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FoundationReserveAuctionCreated)
				if err := _Foundation.contract.UnpackLog(event, "ReserveAuctionCreated", log); err != nil {
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

// ParseReserveAuctionCreated is a log parse operation binding the contract event 0x1062dd3b35f12b4064331244d00f40c1d4831965e4285654157a2409c6217cff.
//
// Solidity: event ReserveAuctionCreated(address indexed seller, address indexed nftContract, uint256 indexed tokenId, uint256 duration, uint256 extensionDuration, uint256 reservePrice, uint256 auctionId)
func (_Foundation *FoundationFilterer) ParseReserveAuctionCreated(log types.Log) (*FoundationReserveAuctionCreated, error) {
	event := new(FoundationReserveAuctionCreated)
	if err := _Foundation.contract.UnpackLog(event, "ReserveAuctionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FoundationReserveAuctionFinalizedIterator is returned from FilterReserveAuctionFinalized and is used to iterate over the raw logs and unpacked data for ReserveAuctionFinalized events raised by the Foundation contract.
type FoundationReserveAuctionFinalizedIterator struct {
	Event *FoundationReserveAuctionFinalized // Event containing the contract specifics and raw log

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
func (it *FoundationReserveAuctionFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FoundationReserveAuctionFinalized)
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
		it.Event = new(FoundationReserveAuctionFinalized)
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
func (it *FoundationReserveAuctionFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FoundationReserveAuctionFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FoundationReserveAuctionFinalized represents a ReserveAuctionFinalized event raised by the Foundation contract.
type FoundationReserveAuctionFinalized struct {
	AuctionId  *big.Int
	Seller     common.Address
	Bidder     common.Address
	TotalFees  *big.Int
	CreatorRev *big.Int
	SellerRev  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterReserveAuctionFinalized is a free log retrieval operation binding the contract event 0x2edb0e99c6ac35be6731dab554c1d1fa1b7beb675090dbb09fb14e615aca1c4a.
//
// Solidity: event ReserveAuctionFinalized(uint256 indexed auctionId, address indexed seller, address indexed bidder, uint256 totalFees, uint256 creatorRev, uint256 sellerRev)
func (_Foundation *FoundationFilterer) FilterReserveAuctionFinalized(opts *bind.FilterOpts, auctionId []*big.Int, seller []common.Address, bidder []common.Address) (*FoundationReserveAuctionFinalizedIterator, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _Foundation.contract.FilterLogs(opts, "ReserveAuctionFinalized", auctionIdRule, sellerRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return &FoundationReserveAuctionFinalizedIterator{contract: _Foundation.contract, event: "ReserveAuctionFinalized", logs: logs, sub: sub}, nil
}

// WatchReserveAuctionFinalized is a free log subscription operation binding the contract event 0x2edb0e99c6ac35be6731dab554c1d1fa1b7beb675090dbb09fb14e615aca1c4a.
//
// Solidity: event ReserveAuctionFinalized(uint256 indexed auctionId, address indexed seller, address indexed bidder, uint256 totalFees, uint256 creatorRev, uint256 sellerRev)
func (_Foundation *FoundationFilterer) WatchReserveAuctionFinalized(opts *bind.WatchOpts, sink chan<- *FoundationReserveAuctionFinalized, auctionId []*big.Int, seller []common.Address, bidder []common.Address) (event.Subscription, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _Foundation.contract.WatchLogs(opts, "ReserveAuctionFinalized", auctionIdRule, sellerRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FoundationReserveAuctionFinalized)
				if err := _Foundation.contract.UnpackLog(event, "ReserveAuctionFinalized", log); err != nil {
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

// ParseReserveAuctionFinalized is a log parse operation binding the contract event 0x2edb0e99c6ac35be6731dab554c1d1fa1b7beb675090dbb09fb14e615aca1c4a.
//
// Solidity: event ReserveAuctionFinalized(uint256 indexed auctionId, address indexed seller, address indexed bidder, uint256 totalFees, uint256 creatorRev, uint256 sellerRev)
func (_Foundation *FoundationFilterer) ParseReserveAuctionFinalized(log types.Log) (*FoundationReserveAuctionFinalized, error) {
	event := new(FoundationReserveAuctionFinalized)
	if err := _Foundation.contract.UnpackLog(event, "ReserveAuctionFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FoundationReserveAuctionInvalidatedIterator is returned from FilterReserveAuctionInvalidated and is used to iterate over the raw logs and unpacked data for ReserveAuctionInvalidated events raised by the Foundation contract.
type FoundationReserveAuctionInvalidatedIterator struct {
	Event *FoundationReserveAuctionInvalidated // Event containing the contract specifics and raw log

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
func (it *FoundationReserveAuctionInvalidatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FoundationReserveAuctionInvalidated)
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
		it.Event = new(FoundationReserveAuctionInvalidated)
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
func (it *FoundationReserveAuctionInvalidatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FoundationReserveAuctionInvalidatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FoundationReserveAuctionInvalidated represents a ReserveAuctionInvalidated event raised by the Foundation contract.
type FoundationReserveAuctionInvalidated struct {
	AuctionId *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterReserveAuctionInvalidated is a free log retrieval operation binding the contract event 0x5603897cc9b1e866f3f7395ffc6638776041f21c094d0b4e748ff44c407fa362.
//
// Solidity: event ReserveAuctionInvalidated(uint256 indexed auctionId)
func (_Foundation *FoundationFilterer) FilterReserveAuctionInvalidated(opts *bind.FilterOpts, auctionId []*big.Int) (*FoundationReserveAuctionInvalidatedIterator, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}

	logs, sub, err := _Foundation.contract.FilterLogs(opts, "ReserveAuctionInvalidated", auctionIdRule)
	if err != nil {
		return nil, err
	}
	return &FoundationReserveAuctionInvalidatedIterator{contract: _Foundation.contract, event: "ReserveAuctionInvalidated", logs: logs, sub: sub}, nil
}

// WatchReserveAuctionInvalidated is a free log subscription operation binding the contract event 0x5603897cc9b1e866f3f7395ffc6638776041f21c094d0b4e748ff44c407fa362.
//
// Solidity: event ReserveAuctionInvalidated(uint256 indexed auctionId)
func (_Foundation *FoundationFilterer) WatchReserveAuctionInvalidated(opts *bind.WatchOpts, sink chan<- *FoundationReserveAuctionInvalidated, auctionId []*big.Int) (event.Subscription, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}

	logs, sub, err := _Foundation.contract.WatchLogs(opts, "ReserveAuctionInvalidated", auctionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FoundationReserveAuctionInvalidated)
				if err := _Foundation.contract.UnpackLog(event, "ReserveAuctionInvalidated", log); err != nil {
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

// ParseReserveAuctionInvalidated is a log parse operation binding the contract event 0x5603897cc9b1e866f3f7395ffc6638776041f21c094d0b4e748ff44c407fa362.
//
// Solidity: event ReserveAuctionInvalidated(uint256 indexed auctionId)
func (_Foundation *FoundationFilterer) ParseReserveAuctionInvalidated(log types.Log) (*FoundationReserveAuctionInvalidated, error) {
	event := new(FoundationReserveAuctionInvalidated)
	if err := _Foundation.contract.UnpackLog(event, "ReserveAuctionInvalidated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FoundationReserveAuctionUpdatedIterator is returned from FilterReserveAuctionUpdated and is used to iterate over the raw logs and unpacked data for ReserveAuctionUpdated events raised by the Foundation contract.
type FoundationReserveAuctionUpdatedIterator struct {
	Event *FoundationReserveAuctionUpdated // Event containing the contract specifics and raw log

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
func (it *FoundationReserveAuctionUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FoundationReserveAuctionUpdated)
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
		it.Event = new(FoundationReserveAuctionUpdated)
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
func (it *FoundationReserveAuctionUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FoundationReserveAuctionUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FoundationReserveAuctionUpdated represents a ReserveAuctionUpdated event raised by the Foundation contract.
type FoundationReserveAuctionUpdated struct {
	AuctionId    *big.Int
	ReservePrice *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterReserveAuctionUpdated is a free log retrieval operation binding the contract event 0x0c0f2662914f0cd1e952db2aa425901cb00e7c1f507687d22cb04e836d55d9c7.
//
// Solidity: event ReserveAuctionUpdated(uint256 indexed auctionId, uint256 reservePrice)
func (_Foundation *FoundationFilterer) FilterReserveAuctionUpdated(opts *bind.FilterOpts, auctionId []*big.Int) (*FoundationReserveAuctionUpdatedIterator, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}

	logs, sub, err := _Foundation.contract.FilterLogs(opts, "ReserveAuctionUpdated", auctionIdRule)
	if err != nil {
		return nil, err
	}
	return &FoundationReserveAuctionUpdatedIterator{contract: _Foundation.contract, event: "ReserveAuctionUpdated", logs: logs, sub: sub}, nil
}

// WatchReserveAuctionUpdated is a free log subscription operation binding the contract event 0x0c0f2662914f0cd1e952db2aa425901cb00e7c1f507687d22cb04e836d55d9c7.
//
// Solidity: event ReserveAuctionUpdated(uint256 indexed auctionId, uint256 reservePrice)
func (_Foundation *FoundationFilterer) WatchReserveAuctionUpdated(opts *bind.WatchOpts, sink chan<- *FoundationReserveAuctionUpdated, auctionId []*big.Int) (event.Subscription, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}

	logs, sub, err := _Foundation.contract.WatchLogs(opts, "ReserveAuctionUpdated", auctionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FoundationReserveAuctionUpdated)
				if err := _Foundation.contract.UnpackLog(event, "ReserveAuctionUpdated", log); err != nil {
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

// ParseReserveAuctionUpdated is a log parse operation binding the contract event 0x0c0f2662914f0cd1e952db2aa425901cb00e7c1f507687d22cb04e836d55d9c7.
//
// Solidity: event ReserveAuctionUpdated(uint256 indexed auctionId, uint256 reservePrice)
func (_Foundation *FoundationFilterer) ParseReserveAuctionUpdated(log types.Log) (*FoundationReserveAuctionUpdated, error) {
	event := new(FoundationReserveAuctionUpdated)
	if err := _Foundation.contract.UnpackLog(event, "ReserveAuctionUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FoundationSellerReferralPaidIterator is returned from FilterSellerReferralPaid and is used to iterate over the raw logs and unpacked data for SellerReferralPaid events raised by the Foundation contract.
type FoundationSellerReferralPaidIterator struct {
	Event *FoundationSellerReferralPaid // Event containing the contract specifics and raw log

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
func (it *FoundationSellerReferralPaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FoundationSellerReferralPaid)
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
		it.Event = new(FoundationSellerReferralPaid)
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
func (it *FoundationSellerReferralPaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FoundationSellerReferralPaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FoundationSellerReferralPaid represents a SellerReferralPaid event raised by the Foundation contract.
type FoundationSellerReferralPaid struct {
	NftContract       common.Address
	TokenId           *big.Int
	SellerReferrer    common.Address
	SellerReferrerFee *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSellerReferralPaid is a free log retrieval operation binding the contract event 0x27a4dd4ff659a9e6354fb079b2208365e5b83f55c22a4150eee2bca89501cb98.
//
// Solidity: event SellerReferralPaid(address indexed nftContract, uint256 indexed tokenId, address sellerReferrer, uint256 sellerReferrerFee)
func (_Foundation *FoundationFilterer) FilterSellerReferralPaid(opts *bind.FilterOpts, nftContract []common.Address, tokenId []*big.Int) (*FoundationSellerReferralPaidIterator, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Foundation.contract.FilterLogs(opts, "SellerReferralPaid", nftContractRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &FoundationSellerReferralPaidIterator{contract: _Foundation.contract, event: "SellerReferralPaid", logs: logs, sub: sub}, nil
}

// WatchSellerReferralPaid is a free log subscription operation binding the contract event 0x27a4dd4ff659a9e6354fb079b2208365e5b83f55c22a4150eee2bca89501cb98.
//
// Solidity: event SellerReferralPaid(address indexed nftContract, uint256 indexed tokenId, address sellerReferrer, uint256 sellerReferrerFee)
func (_Foundation *FoundationFilterer) WatchSellerReferralPaid(opts *bind.WatchOpts, sink chan<- *FoundationSellerReferralPaid, nftContract []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Foundation.contract.WatchLogs(opts, "SellerReferralPaid", nftContractRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FoundationSellerReferralPaid)
				if err := _Foundation.contract.UnpackLog(event, "SellerReferralPaid", log); err != nil {
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

// ParseSellerReferralPaid is a log parse operation binding the contract event 0x27a4dd4ff659a9e6354fb079b2208365e5b83f55c22a4150eee2bca89501cb98.
//
// Solidity: event SellerReferralPaid(address indexed nftContract, uint256 indexed tokenId, address sellerReferrer, uint256 sellerReferrerFee)
func (_Foundation *FoundationFilterer) ParseSellerReferralPaid(log types.Log) (*FoundationSellerReferralPaid, error) {
	event := new(FoundationSellerReferralPaid)
	if err := _Foundation.contract.UnpackLog(event, "SellerReferralPaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FoundationSellersAddedToExhibitionIterator is returned from FilterSellersAddedToExhibition and is used to iterate over the raw logs and unpacked data for SellersAddedToExhibition events raised by the Foundation contract.
type FoundationSellersAddedToExhibitionIterator struct {
	Event *FoundationSellersAddedToExhibition // Event containing the contract specifics and raw log

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
func (it *FoundationSellersAddedToExhibitionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FoundationSellersAddedToExhibition)
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
		it.Event = new(FoundationSellersAddedToExhibition)
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
func (it *FoundationSellersAddedToExhibitionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FoundationSellersAddedToExhibitionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FoundationSellersAddedToExhibition represents a SellersAddedToExhibition event raised by the Foundation contract.
type FoundationSellersAddedToExhibition struct {
	ExhibitionId *big.Int
	Sellers      []common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSellersAddedToExhibition is a free log retrieval operation binding the contract event 0xd5a31bd2d34d303249ac7f54bfc7578390f90f5d39cb39813f67190fa36f5c17.
//
// Solidity: event SellersAddedToExhibition(uint256 indexed exhibitionId, address[] sellers)
func (_Foundation *FoundationFilterer) FilterSellersAddedToExhibition(opts *bind.FilterOpts, exhibitionId []*big.Int) (*FoundationSellersAddedToExhibitionIterator, error) {

	var exhibitionIdRule []interface{}
	for _, exhibitionIdItem := range exhibitionId {
		exhibitionIdRule = append(exhibitionIdRule, exhibitionIdItem)
	}

	logs, sub, err := _Foundation.contract.FilterLogs(opts, "SellersAddedToExhibition", exhibitionIdRule)
	if err != nil {
		return nil, err
	}
	return &FoundationSellersAddedToExhibitionIterator{contract: _Foundation.contract, event: "SellersAddedToExhibition", logs: logs, sub: sub}, nil
}

// WatchSellersAddedToExhibition is a free log subscription operation binding the contract event 0xd5a31bd2d34d303249ac7f54bfc7578390f90f5d39cb39813f67190fa36f5c17.
//
// Solidity: event SellersAddedToExhibition(uint256 indexed exhibitionId, address[] sellers)
func (_Foundation *FoundationFilterer) WatchSellersAddedToExhibition(opts *bind.WatchOpts, sink chan<- *FoundationSellersAddedToExhibition, exhibitionId []*big.Int) (event.Subscription, error) {

	var exhibitionIdRule []interface{}
	for _, exhibitionIdItem := range exhibitionId {
		exhibitionIdRule = append(exhibitionIdRule, exhibitionIdItem)
	}

	logs, sub, err := _Foundation.contract.WatchLogs(opts, "SellersAddedToExhibition", exhibitionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FoundationSellersAddedToExhibition)
				if err := _Foundation.contract.UnpackLog(event, "SellersAddedToExhibition", log); err != nil {
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

// ParseSellersAddedToExhibition is a log parse operation binding the contract event 0xd5a31bd2d34d303249ac7f54bfc7578390f90f5d39cb39813f67190fa36f5c17.
//
// Solidity: event SellersAddedToExhibition(uint256 indexed exhibitionId, address[] sellers)
func (_Foundation *FoundationFilterer) ParseSellersAddedToExhibition(log types.Log) (*FoundationSellersAddedToExhibition, error) {
	event := new(FoundationSellersAddedToExhibition)
	if err := _Foundation.contract.UnpackLog(event, "SellersAddedToExhibition", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FoundationWithdrawalToFETHIterator is returned from FilterWithdrawalToFETH and is used to iterate over the raw logs and unpacked data for WithdrawalToFETH events raised by the Foundation contract.
type FoundationWithdrawalToFETHIterator struct {
	Event *FoundationWithdrawalToFETH // Event containing the contract specifics and raw log

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
func (it *FoundationWithdrawalToFETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FoundationWithdrawalToFETH)
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
		it.Event = new(FoundationWithdrawalToFETH)
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
func (it *FoundationWithdrawalToFETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FoundationWithdrawalToFETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FoundationWithdrawalToFETH represents a WithdrawalToFETH event raised by the Foundation contract.
type FoundationWithdrawalToFETH struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalToFETH is a free log retrieval operation binding the contract event 0xa2201512569adb2d513531dfd69b66df50bd5cffb8c1bbe65a4611f9e1eadbd1.
//
// Solidity: event WithdrawalToFETH(address indexed user, uint256 amount)
func (_Foundation *FoundationFilterer) FilterWithdrawalToFETH(opts *bind.FilterOpts, user []common.Address) (*FoundationWithdrawalToFETHIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Foundation.contract.FilterLogs(opts, "WithdrawalToFETH", userRule)
	if err != nil {
		return nil, err
	}
	return &FoundationWithdrawalToFETHIterator{contract: _Foundation.contract, event: "WithdrawalToFETH", logs: logs, sub: sub}, nil
}

// WatchWithdrawalToFETH is a free log subscription operation binding the contract event 0xa2201512569adb2d513531dfd69b66df50bd5cffb8c1bbe65a4611f9e1eadbd1.
//
// Solidity: event WithdrawalToFETH(address indexed user, uint256 amount)
func (_Foundation *FoundationFilterer) WatchWithdrawalToFETH(opts *bind.WatchOpts, sink chan<- *FoundationWithdrawalToFETH, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Foundation.contract.WatchLogs(opts, "WithdrawalToFETH", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FoundationWithdrawalToFETH)
				if err := _Foundation.contract.UnpackLog(event, "WithdrawalToFETH", log); err != nil {
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

// ParseWithdrawalToFETH is a log parse operation binding the contract event 0xa2201512569adb2d513531dfd69b66df50bd5cffb8c1bbe65a4611f9e1eadbd1.
//
// Solidity: event WithdrawalToFETH(address indexed user, uint256 amount)
func (_Foundation *FoundationFilterer) ParseWithdrawalToFETH(log types.Log) (*FoundationWithdrawalToFETH, error) {
	event := new(FoundationWithdrawalToFETH)
	if err := _Foundation.contract.UnpackLog(event, "WithdrawalToFETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
