package token

import (
	"context"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/pregod/common/database"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/datasource/coingecko"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/erc20"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/worker/zksync"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
	"github.com/samber/lo"

	"golang.org/x/sync/errgroup"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"go.uber.org/ratelimit"
	"go.uber.org/zap"
)

var _ worker.Job = (*Job)(nil)

type Job struct {
	coingeckoClient *coingecko.Client
	zksyncClient    *zksync.Client
	databaseClient  *gorm.DB
	rateLimiter     ratelimit.Limiter
}

func (j *Job) Name() string {
	return "token"
}

func (j *Job) Spec() string {
	return "@every 1m"
}

func (j *Job) Timeout() time.Duration {
	return 5 * time.Hour // ~ (13264 / 50 / 60)
}

func (j *Job) Run(_ worker.RenewalFunc) error {
	errorGroup := errgroup.Group{}

	ctx := context.Background()

	errorGroup.Go(func() error {
		return j.refreshTokenListFromCoinGecko(ctx)
	})

	errorGroup.Go(func() error {
		return j.refreshTokenListFromZkSync(ctx)
	})

	return errorGroup.Wait()
}

func (j *Job) refreshTokenListFromCoinGecko(ctx context.Context) error {
	j.rateLimiter.Take()

	coinList, err := j.coingeckoClient.CoinList(ctx, coingecko.CoinListParameter{
		IncludePlatform: true,
	})
	if err != nil {
		return fmt.Errorf("fetch token list from coingecko: %w", err)
	}

	return j.buildTokenListFromCoinGecko(ctx, coinList)
}

func (j *Job) buildTokenListFromCoinGecko(ctx context.Context, coinList []coingecko.Coin) error {
	// The Token List provided by CoinGecko has some tokens with empty information,
	// we should filter them.
	coinList = lo.Filter(coinList, func(coin coingecko.Coin, _ int) bool {
		for _, value := range coin.Platforms {
			if value != "" {
				return true
			}
		}

		return false
	})

	for _, coin := range coinList {
		tokens, err := j.buildTokenFromCoinGecko(ctx, coin)
		if err != nil {
			return fmt.Errorf("build token from coingecko: %w", err)
		}

		if err := j.storeTokens(ctx, tokens); err != nil {
			return fmt.Errorf("store token: %w", err)
		}
	}

	return nil
}

func (j *Job) buildTokenFromCoinGecko(ctx context.Context, coin coingecko.Coin) ([]model.Token, error) {
	j.rateLimiter.Take()

	internalCoin, err := j.coingeckoClient.Coin(ctx, coin.ID, coingecko.CoinParameter{})
	if err != nil {
		return nil, fmt.Errorf("coin from coingecko: %w", err)
	}

	logo, _ := j.buildLogoURL(*internalCoin)

	tokens := make([]model.Token, 0, len(coin.Platforms))

	for platform, value := range internalCoin.Platforms {
		// Filter unsupported networks to reduce unnecessary third-party API and RPC requests
		network, ok := j.platformToNetwork(platform)
		if !ok {
			continue
		}

		ethereumClient, err := ethclientx.Global(network)
		if err != nil {
			zap.L().Debug("load ethereum client", zap.String("job", j.Name()), zap.String("network", network), zap.Error(err))

			// Skip networks that may be supported but not enabled
			continue
		}

		zap.L().Debug("build token", zap.String("job", j.Name()), zap.String("token", coin.Name), zap.String("image", logo), zap.String("platform", platform), zap.String("value", value))

		erc20Caller, err := erc20.NewERC20Caller(common.HexToAddress(value), ethereumClient)
		if err != nil {
			zap.L().Warn("create erc20 caller", zap.String("job", j.Name()), zap.String("network", network), zap.String("contract_address", value), zap.Error(err))

			continue
		}

		name, err := erc20Caller.Name(&bind.CallOpts{})
		if err != nil {
			zap.L().Warn("call erc20 name", zap.String("job", j.Name()), zap.String("network", network), zap.String("contract_address", value), zap.Error(err))

			continue
		}

		symbol, err := erc20Caller.Symbol(&bind.CallOpts{})
		if err != nil {
			zap.L().Warn("call erc20 symbol", zap.String("job", j.Name()), zap.String("network", network), zap.String("contract_address", value), zap.Error(err))

			continue
		}

		decimals, err := erc20Caller.Decimals(&bind.CallOpts{})
		if err != nil {
			zap.L().Warn("call erc20 decimals", zap.String("job", j.Name()), zap.String("network", network), zap.String("contract_address", value), zap.Error(err))

			continue
		}

		tokens = append(tokens, model.Token{
			ID:              coin.ID,
			Name:            name,
			Symbol:          symbol,
			Decimal:         decimals,
			Logo:            logo,
			Standard:        protocol.TokenStandardERC20,
			Network:         network,
			ContractAddress: strings.ToLower(value),
		})
	}

	return tokens, nil
}

func (j *Job) refreshTokenListFromZkSync(ctx context.Context) error {
	j.rateLimiter.Take()

	coinList, err := j.coingeckoClient.CoinList(ctx, coingecko.CoinListParameter{
		IncludePlatform: true,
	})
	if err != nil {
		return fmt.Errorf("fetch token list from coingecko: %w", err)
	}

	tokenList, err := j.zksyncClient.GetTokenList()
	if err != nil {
		return fmt.Errorf("fetch token list from zksync: %w", err)
	}

	tokens, err := j.buildTokenListFromZkSync(ctx, tokenList, coinList)
	if err != nil {
		return fmt.Errorf("update token list from zksync: %w", err)
	}

	return j.storeTokens(ctx, tokens)
}

func (j *Job) buildTokenListFromZkSync(ctx context.Context, tokenList zksync.GetTokenList, coinList []coingecko.Coin) ([]model.Token, error) {
	tokens := make([]model.Token, 0, len(tokenList))

	for _, token := range tokenList {
		coin, exists := lo.Find(coinList, func(coin coingecko.Coin) bool {
			contractAddress, exists := coin.Platforms["ethereum"]
			if !exists {
				return exists
			}

			return strings.EqualFold(contractAddress, token.Address)
		})

		var logo string

		if exists {
			logo, _ = j.buildLogoURL(coin)
		}

		tokens = append(tokens, model.Token{
			ID:              coin.ID,
			Name:            coin.Name,
			Symbol:          coin.Symbol,
			Decimal:         token.Decimals,
			Logo:            logo,
			Standard:        protocol.TokenStandardERC20,
			Network:         protocol.NetworkZkSync,
			ContractAddress: strings.ToLower(token.Address),
		})
	}

	return tokens, nil
}

func (j *Job) storeTokens(ctx context.Context, tokens []model.Token) error {
	if len(tokens) > 0 {
		onConflictClause := clause.OnConflict{
			UpdateAll: true,
		}

		if err := j.databaseClient.Model([]model.Token{}).Clauses(onConflictClause).Create(tokens).Error; err != nil {
			return fmt.Errorf("create token in batches: %w", err)
		}
	}

	return nil
}

func (j *Job) buildLogoURL(coin coingecko.Coin) (string, error) {
	imageURL := ""

	switch {
	case coin.Image == nil:
	case coin.Image.Large != "":
		imageURL = coin.Image.Large
	case coin.Image.Thumb != "":
		imageURL = coin.Image.Thumb
	case coin.Image.Small != "":
		imageURL = coin.Image.Small
	}

	logoURL, err := url.Parse(imageURL)
	if err != nil {
		return "", fmt.Errorf("parse image url: %w", err)
	}

	// Purge query parameters used to bypass CDN cache.
	logoURL.RawQuery = ""

	return logoURL.String(), nil
}

func (j *Job) platformToNetwork(platform string) (string, bool) {
	var network string

	switch platform {
	case "ethereum":
		network = protocol.NetworkEthereum
	case "binance-smart-chain":
		network = protocol.NetworkBinanceSmartChain
	case "polygon-pos":
		network = protocol.NetworkPolygon
	case "avalanche":
		network = protocol.NetworkAvalanche
	case "fantom":
		network = protocol.NetworkFantom
	case "arbitrum-one":
		network = protocol.NetworkArbitrum
	case "optimistic-ethereum":
		network = protocol.NetworkOptimism
	case "xdai":
		network = protocol.NetworkXDAI
	}

	return network, network != ""
}

func New() worker.Job {
	return &Job{
		coingeckoClient: coingecko.New(),
		zksyncClient:    zksync.New(),
		databaseClient:  database.Global(),
		// https://www.coingecko.com/en/api/documentation
		// `Our Free API* has a rate limit of 10-50 calls/minute, and doesn't require API key.`
		rateLimiter: ratelimit.New(50, ratelimit.Per(time.Minute)),
	}
}
