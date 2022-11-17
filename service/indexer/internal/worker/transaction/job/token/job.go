package token

import (
	"context"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/naturalselectionlabs/pregod/common/database"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/datasource/coingecko"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/erc20"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/worker/zksync"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
	"github.com/robfig/cron/v3"
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
	crontab         *cron.Cron
}

func (j *Job) Name() string {
	return "token"
}

func (j *Job) Spec() string {
	return "@every 10s"
}

func (j *Job) Timeout() time.Duration {
	// It takes about ~5 (13264 / 50 / 60) hours in total,
	// but the lock may not be released properly when the pod is closed,
	// so the timeout is positioned as the value of a single token
	return 2 * time.Minute
}

func (j *Job) Run(renewal worker.RenewalFunc) error {
	errorGroup := errgroup.Group{}

	ctx := context.Background()

	_, _ = j.crontab.AddFunc("@every 1m", func() {
		_ = renewal(ctx, j.Timeout())
	})

	j.crontab.Start()
	defer j.crontab.Stop()

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

		zap.L().Debug("build token", zap.String("job", j.Name()), zap.String("token", coin.Name), zap.String("image", logo), zap.String("network", network), zap.String("contract_address", value))

		token, err := j.buildToken(ctx, value, network, ethereumClient)
		if err != nil {
			zap.L().Warn("build token", zap.String("job", j.Name()), zap.String("token", coin.Name), zap.String("image", logo), zap.String("network", network), zap.String("contract_address", value), zap.Error(err))

			continue
		}

		token.ID = coin.ID
		token.Logo = logo

		tokens = append(tokens, *token)
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
		// zkSync Native ETH
		if token.ID == 0 {
			tokens = append(tokens, model.Token{
				ID:       "eth",
				Name:     "ETH",
				Symbol:   "ETH",
				Logo:     "https://assets.coingecko.com/coins/images/279/large/ethereum.png",
				Standard: protocol.TokenStandardERC20,
				Network:  protocol.NetworkZkSync,
				Decimal:  18,
			})

			continue
		}

		coin, exists := lo.Find(coinList, func(coin coingecko.Coin) bool {
			contractAddress, exists := coin.Platforms[protocol.NetworkEthereum]
			if !exists {
				return exists
			}

			return strings.EqualFold(contractAddress, token.Address)
		})

		var logo, id string

		if exists {
			j.rateLimiter.Take()

			coingeckoCoin, err := j.coingeckoClient.Coin(ctx, coin.ID, coingecko.CoinParameter{})
			if err != nil {
				zap.L().Warn("fetch token from coingecko", zap.String("job", j.Name()), zap.String("token", token.Symbol), zap.String("contract_address", token.Address), zap.Error(err))

				continue
			}

			logo, _ = j.buildLogoURL(*coingeckoCoin)

			id = strings.ToLower(coin.ID)
		} else {
			id = strings.ToLower(token.Symbol)
		}

		ethereumClient, err := ethclientx.Global(protocol.NetworkEthereum)
		if err != nil {
			return nil, fmt.Errorf("load ethereum client: %w", err)
		}

		internalToken, err := j.buildToken(ctx, token.Address, protocol.NetworkZkSync, ethereumClient)
		if err != nil {
			zap.L().Warn("build token", zap.String("job", j.Name()), zap.String("token", coin.Name), zap.String("image", logo), zap.String("network", protocol.NetworkZkSync), zap.String("contract_address", token.Address), zap.Error(err))

			continue
		}

		internalToken.ID = id
		internalToken.Decimal = token.Decimals

		tokens = append(tokens, *internalToken)
	}

	return tokens, nil
}

func (j *Job) buildToken(ctx context.Context, address string, network string, ethereumClient *ethclient.Client) (*model.Token, error) {
	erc20Caller, err := erc20.NewERC20Caller(common.HexToAddress(address), ethereumClient)
	if err != nil {
		return nil, err
	}

	name, err := erc20Caller.Name(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}

	symbol, err := erc20Caller.Symbol(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}

	decimals, err := erc20Caller.Decimals(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}

	return &model.Token{
		Name:            name,
		Symbol:          symbol,
		Decimal:         decimals,
		Standard:        protocol.TokenStandardERC20,
		Network:         network,
		ContractAddress: strings.ToLower(address),
	}, nil
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
		crontab:     cron.New(),
	}
}
