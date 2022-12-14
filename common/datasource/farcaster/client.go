package farcaster

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"sync"

	"github.com/go-redis/redis/v8"
	"github.com/go-resty/resty/v2"
	"github.com/naturalselectionlabs/pregod/common/cache"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	"go.uber.org/zap"
)

const (
	// https://www.notion.so/Migration-Guide-Phase-I-f34bf0b713924ba6b9652b6bb4fe44ef
	BaseUrl  = "https://api.farcaster.xyz"
	PageSize = 1000
)

var (
	farcasterCacheMap map[string]*CacheAddress
	globalLocker      sync.RWMutex
)

type Client struct {
	client *resty.Client
}

func Global() map[string]*CacheAddress {
	globalLocker.RLock()

	defer globalLocker.RUnlock()

	item := make(map[string]*CacheAddress)
	for key, value := range farcasterCacheMap {
		item[key] = value
	}

	return item
}

func ReplaceGlobal(address string, cacheAddress *CacheAddress) {
	globalLocker.Lock()

	defer globalLocker.Unlock()

	farcasterCacheMap[address] = cacheAddress
}

func (c *Client) SetFarcasterCacheMap() error {
	var err error
	farcasterCacheMap, err = c.GetLastMapFromCache(context.Background())

	if len(farcasterCacheMap) == 0 {
		farcasterCacheMap = make(map[string]*CacheAddress)
		farAddresses, err := c.GetUserList(context.Background())
		if err != nil {
			loggerx.Global().Named("GetFarcasterCacheMap").Warn("unable to get farcaster data", zap.Error(err))
		}

		for _, cast := range farAddresses {
			evmAddress, err := c.ConvertFarcasterAdderess(context.Background(), cast.Address)
			if err != nil {
				loggerx.Global().Named("GetFarcasterCacheMap").Warn("unable to convert farcaster address", zap.Error(err))
			}
			if len(evmAddress) > 0 {
				ReplaceGlobal(strings.ToLower(cast.Address), &CacheAddress{
					EvmAddress: strings.ToLower(evmAddress),
				})
			}
		}
		err = c.SetCurrentMap(context.Background(), farcasterCacheMap)
		if err != nil && err != redis.Nil {
			loggerx.Global().Named("GetFarcasterCacheMap").Warn("unable to set cache data", zap.Error(err))
		}
	}

	return err
}

func (c *Client) UpdateFarcasterCacheMap() {
	farAddresses, err := c.GetUserList(context.Background())
	if err != nil {
		loggerx.Global().Named("UpdateFarcasterCacheMap").Warn("unable to get farcaster data", zap.Error(err))
	}

	for _, cast := range farAddresses {
		_, ok := farcasterCacheMap[strings.ToLower(cast.Address)]
		if !ok {
			evmAddress, err := c.ConvertFarcasterAdderess(context.Background(), cast.Address)
			if err != nil {
				loggerx.Global().Named("UpdateFarcasterCacheMap").Warn("unable to convert farcaster address", zap.Error(err))
			}
			if len(evmAddress) > 0 {
				ReplaceGlobal(strings.ToLower(cast.Address), &CacheAddress{
					EvmAddress: strings.ToLower(evmAddress),
				})
			}
		}
	}
}

func (c *Client) GetLastMapFromCache(ctx context.Context) (farcasterCacheMap map[string]*CacheAddress, err error) {
	if cache.Global() == nil {
		return nil, fmt.Errorf("redis worker is nil")
	}

	mapKey := fmt.Sprintf("crawler_%s", protocol.PlatformFarcaster)
	farcasterCacheMap = make(map[string]*CacheAddress)

	data, err := cache.Global().Get(ctx, mapKey).Result()
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal([]byte(data), &farcasterCacheMap); err != nil {
		return nil, fmt.Errorf("unmarshal %s from cache error:%+v", mapKey, err)
	}

	return farcasterCacheMap, nil
}

func (c *Client) SetCurrentMap(ctx context.Context, farcasterCacheMap map[string]*CacheAddress) error {
	if cache.Global() == nil {
		return fmt.Errorf("redis worker is nil")
	}

	if len(farcasterCacheMap) == 0 {
		return fmt.Errorf("farcasterCacheMap is less than 0")
	}

	data, err := json.Marshal(farcasterCacheMap)
	if err != nil {
		return fmt.Errorf("marshal %+v to json error:%+v", farcasterCacheMap, err)
	}

	mapKey := fmt.Sprintf("crawler_%s", protocol.PlatformFarcaster)
	cache.Global().Set(ctx, mapKey, data, 0)

	return nil
}

func (c *Client) GetUserList(ctx context.Context) ([]FarcasterUser, error) {
	var res []FarcasterUser
	param := map[string]string{
		"filter":   "recent",
		"per_page": fmt.Sprint(PageSize),
	}

	page := 1

	for {
		var list []FarcasterUser
		param["page"] = fmt.Sprint(page)
		_, err := c.client.R().
			SetResult(&list).
			SetQueryParams(param).
			SetContext(ctx).
			Get("/indexer/users")
		if err != nil {
			loggerx.Global().Named("GetUserList").Warn("unable to get farcaster data", zap.Error(err))
			return nil, err
		}

		res = append(res, list...)
		if len(list) < PageSize {
			break
		}

		page++

	}

	return res, nil
}

func (c *Client) ConvertFarcasterAdderess(ctx context.Context, address string) (string, error) {
	var evmAddress ConvertAddress

	param := map[string]string{
		"address": address,
	}

	_, err := c.client.R().
		SetResult(&evmAddress).
		SetPathParams(param).
		SetContext(ctx).
		Get("/origin/proof/{address}")
	if err != nil {
		loggerx.Global().Named("ConvertFarcasterAdderess").Warn("unable to convert farcaster address", zap.Error(err))
		return "", err
	}

	return strings.ToLower(evmAddress.SignerAddress), nil
}

func (c *Client) GetActivityList(ctx context.Context, address string) ([]Cast, error) {
	var activityList []Cast

	param := map[string]string{
		"address": address,
	}

	_, err := c.client.R().
		SetResult(&activityList).
		SetPathParams(param).
		SetContext(ctx).
		Get("/origin/address_activity/{address}")
	if err != nil {
		loggerx.Global().Named("GetActivityList").Warn("unable to get farcaster data", zap.Error(err))
		return nil, err
	}

	return activityList, nil
}

func NewClient() *Client {
	client := resty.New()
	client.SetBaseURL(BaseUrl)
	return &Client{
		client: client,
	}
}
