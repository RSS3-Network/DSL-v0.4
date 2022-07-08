package coinmarketcap

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-resty/resty/v2"
	"github.com/naturalselectionlabs/pregod/common/cache"
	"github.com/naturalselectionlabs/pregod/common/database"
	"github.com/naturalselectionlabs/pregod/common/database/model/transaction"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/transaction/contract"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	Prefix = "https://pro-api.coinmarketcap.com"
)

var (
	client      = resty.New()
	Network2RPC = map[string]string{
		"ethereum":            "https://eth.rss3.dev",
		"binance_smart_chain": "https://bsc-dataseed.binance.org/",
		"polygon":             "https://polygon-mainnet.infura.io/v3/56ff3a0f0ad14ec19dc933038c69bfbe",
		"zksync":              "https://zksync2-testnet.zksync.dev",
		// "xdai":                "https://rpc.ankr.com/gnosis",
		"xdai":      "https://rpc.gnosischain.com/",
		"arbitrum":  "https://rpc.ankr.com/arbitrum",
		"avalanche": "https://rpc.ankr.com/avalanche",
		"celo":      "https://rpc.ankr.com/celo",
		"fantom":    "https://rpc.ankr.com/fantom",
		"harmony":   "https://rpc.ankr.com/harmony",
		"iotex":     "https://rpc.ankr.com/iotex",
		"moonbeam":  "https://rpc.ankr.com/moonbeam",
		"near":      "https://rpc.ankr.com/near",
		"nervos":    "https://rpc.ankr.com/nervos_gw",
		"syscoin":   "https://rpc.ankr.com/syscoin",
	}
	Network2Token = map[string]struct {
		Symbol  string
		Address string
	}{
		"ethereum":            {"ETH", ""},
		"binance_smart_chain": {"WBNB", "0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c"},  // WBNB address (on BSC)
		"polygon":             {"MATIC", "0x0000000000000000000000000000000000001010"}, // MATIC address (on polygon)
		"xdai":                {"GNO", "0x9c58bacc331c9aa871afd802db6379a98e80cedb"},   // Gnosis address (on xDAI)
	}
)

func Init(apikey string) {
	client.SetHeaders(map[string]string{
		"Content-Type":      "application/json",
		"X-CMC_PRO_API_KEY": apikey,
	})
}

type CoinInfos struct {
	Data map[string]*transaction.CoinMarketCapCoinInfo
}

func (i *CoinInfos) List() []*transaction.CoinMarketCapCoinInfo {
	result := []*transaction.CoinMarketCapCoinInfo{}
	for _, v := range i.Data {
		result = append(result, v)
	}
	return result
}

// Not the same as the documentation
type CoinInfosBySymbol struct {
	Data map[string][]*transaction.CoinMarketCapCoinInfo
}

func (i *CoinInfosBySymbol) List() []*transaction.CoinMarketCapCoinInfo {
	result := []*transaction.CoinMarketCapCoinInfo{}
	for _, v := range i.Data {
		result = append(result, v...)
	}
	return result
}

// https://coinmarketcap.com/api/documentation/v1/#operation/getV2CryptocurrencyInfo
// Will get a empty result when resp code is not OK
func getCoinInfo(ctx context.Context, network, address string) (*transaction.CoinMarketCapCoinInfo, error) {
	path := "/v2/cryptocurrency/info"

	result := CoinInfos{}
	resp, err := client.R().SetResult(&result).SetQueryParams(map[string]string{
		"address": address,
	}).Get(Prefix + path)
	if err != nil {
		return nil, err
	}
	if resp.IsError() { // skip this one and return a empty result
		return new(transaction.CoinMarketCapCoinInfo), nil
		// return nil, fmt.Errorf("bad resp code: %v", resp.StatusCode())
	}
	info := result.List()[0]

	// get decimals
	decimals, err := getDecimals(ctx, network, address)
	if err != nil {
		return nil, err
	}
	info.Decimals = decimals

	// addresses
	info.FillFields()

	return info, nil
}

func getDecimals(ctx context.Context, network, addressHex string) (uint8, error) {
	rpcUrl := Network2RPC[network]
	if rpcUrl == "" {
		rpcUrl = Network2RPC["ethereum"]
	}

	client, err := ethclient.DialContext(ctx, rpcUrl)
	if err != nil {
		return 0, err
	}

	contractCaller, err := contract.NewERC20(common.HexToAddress(addressHex), client)
	if err != nil {
		return 0, err
	}

	return contractCaller.Decimals(&bind.CallOpts{})
}

func CachedGetCoinInfo(ctx context.Context, network, address string) (*transaction.CoinMarketCapCoinInfo, error) {
	result := &transaction.CoinMarketCapCoinInfo{}

	// get from redis cache
	key := fmt.Sprintf("get_coin_info.%s.%s", network, address)
	exists, err := cache.GetMsgPack(ctx, key, result)
	if err != nil || exists {
		return result, err
	}

	// get from db
	if err := database.Client.Where("? = ANY(addresses)", address).First(result).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			logrus.Error(err)
			return nil, err
		}
	} else {
		return result, nil
	}

	result, err = getCoinInfo(ctx, network, address)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	// save
	if err := database.Client.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(result).Error; err != nil {
		logrus.Error(err)
		return nil, err
	}
	err = cache.SetMsgPack(ctx, key, result, 24*time.Hour)

	return result, err
}

// for transaction native transfer
func getCoinInfoByNetwork(ctx context.Context, network string) (*transaction.CoinMarketCapCoinInfo, error) {
	token, exists := Network2Token[network]
	if !exists {
		return nil, errors.New(network + ": symbol for the network not exists")
	}

	path := "/v2/cryptocurrency/info"

	result := CoinInfosBySymbol{}
	resp, err := client.R().SetResult(&result).SetQueryParams(map[string]string{
		"symbol": token.Symbol,
	}).Get(Prefix + path)
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, fmt.Errorf("bad resp code: %v", resp.StatusCode())
	}
	info := result.List()[0]

	// get decimals
	switch {
	case token.Symbol == "ETH":
		info.Decimals = 18
	default:
		decimals, err := getDecimals(ctx, network, token.Address)
		if err != nil {
			return nil, err
		}
		info.Decimals = decimals
	}

	// addresses
	info.FillFields()

	return info, nil
}

// for transaction native transfer
func CachedGetCoinInfoByNetwork(ctx context.Context, network string, logIndex int64) (*transaction.CoinMarketCapCoinInfo, error) {
	// hard code xDAI / BNB info
	switch {
	case logIndex == -1 && network == protocol.NetworkXDAI:
		return &transaction.CoinMarketCapCoinInfo{
			Name:        "xDAI",
			Symbol:      "xDAI",
			Logo:        "https://s2.coinmarketcap.com/static/img/coins/64x64/8635.png",
			Slug:        "xdaistable",
			Category:    "transaction",
			Description: "xDAI (xDAI) is a cryptocurrency . xDAI has a current supply of 0. The last known price of xDAI is 0.99581946 USD and is up 0.16 over the last 24 hours. It is currently trading on 2 active market(s) with $108,496.32 traded over the last 24 hours. More information can be found at http://xdaichain.com/.",
			Decimals:    18,
		}, nil
	case logIndex == -1 && network == protocol.NetworkBinanceSmartChain:
		return &transaction.CoinMarketCapCoinInfo{
			Name:        "BNB",
			Symbol:      "BNB",
			Logo:        "https://s2.coinmarketcap.com/static/img/coins/64x64/1839.png",
			Slug:        "bnb",
			Category:    "coin",
			Description: "BNB (BNB) is a cryptocurrency . BNB has a current supply of 163,276,974.63. The last known price of BNB is 220.69016174 USD and is down -5.85 over the last 24 hours. It is currently trading on 981 active market(s) with $1,183,715,416.39 traded over the last 24 hours. More information can be found at https://www.binance.com/.",
			Decimals:    18,
		}, nil
	}
	// end: hard code xDAI / BNB info

	token, exists := Network2Token[network]
	if !exists {
		return nil, errors.New(network + ": symbol for the network not exists")
	}

	result := &transaction.CoinMarketCapCoinInfo{}

	// get from redis cache
	key := fmt.Sprintf("get_coin_info_by_network.%s", network)
	exists, err := cache.GetMsgPack(ctx, key, result)
	if err != nil || exists {
		return result, err
	}

	// get from db
	if err := database.Client.Where("symbol = ?", token.Symbol).First(result).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, err
		}
	} else {
		return result, nil
	}

	result, err = getCoinInfoByNetwork(ctx, network)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	// save
	if err := database.Client.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(result).Error; err != nil {
		logrus.Error(err)
		return nil, err
	}
	err = cache.SetMsgPack(ctx, key, result, 24*time.Hour)

	return result, err
}
