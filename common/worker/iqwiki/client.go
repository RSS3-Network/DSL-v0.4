package iqwiki

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"github.com/Khan/genqlient/graphql"
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-redis/redis/v8"
	"github.com/naturalselectionlabs/pregod/common/cache"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	graphqlx "github.com/naturalselectionlabs/pregod/common/worker/iqwiki/graphql"
	"go.uber.org/zap"
)

const (
	Scheme   = "https"
	Endpoint = "graph.everipedia.org"
	Path     = "graphql"
	Limit    = 30
)

var (
	IqwikiCacheMap map[string]int
	once           sync.Once
	globalLocker   sync.RWMutex
)

func ReplaceGlobal(address string, activityNum int) {
	globalLocker.Lock()

	defer globalLocker.Unlock()

	IqwikiCacheMap[address] = activityNum
}

type Client struct {
	GClient *graphql.Client
}

func (c *Client) GetIqwikiCacheMap() (map[string]int, error) {
	var err error
	once.Do(func() {
		IqwikiCacheMap, err = c.GetLastMapFromCache(context.Background())

		if len(IqwikiCacheMap) == 0 {
			IqwikiCacheMap = make(map[string]int)
			iqwikiUsers, err := c.GetUserList(context.Background())
			if err != nil {
				loggerx.Global().Named("GetIqwikiCacheMap").Warn("unable to get iqwiki data", zap.Error(err))
			}

			for _, user := range iqwikiUsers {
				if len(user.Id) > 0 {
					ReplaceGlobal(strings.ToLower(user.Id), 0)
				}
			}
			err = c.SetCurrentMap(context.Background(), IqwikiCacheMap)
			if err != nil && err != redis.Nil {
				loggerx.Global().Named("GetIqwikiCacheMap").Warn("unable to set cache data", zap.Error(err))
			}
		}
	})
	return IqwikiCacheMap, err
}

func (c *Client) UpdateIqwikiCacheMap() {
	iqwikiUsers, err := c.GetUserList(context.Background())
	if err != nil {
		loggerx.Global().Named("UpdateIqwikiCacheMap").Warn("unable to get iqwiki data", zap.Error(err))
	}

	for _, user := range iqwikiUsers {
		_, ok := IqwikiCacheMap[strings.ToLower(user.Id)]
		if !ok {
			if len(user.Id) > 0 {
				ReplaceGlobal(strings.ToLower(user.Id), 0)
			}
		}
	}
}

func (c *Client) GetLastMapFromCache(ctx context.Context) (iqwikiCacheMap map[string]int, err error) {
	if cache.Global() == nil {
		return nil, fmt.Errorf("redis worker is nil")
	}

	mapKey := fmt.Sprintf("crawler_%s", protocol.PlatformIQWiki)
	iqwikiCacheMap = make(map[string]int)

	data, err := cache.Global().Get(ctx, mapKey).Result()
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal([]byte(data), &iqwikiCacheMap); err != nil {
		return nil, fmt.Errorf("unmarshal %s from cache error:%+v", mapKey, err)
	}

	return iqwikiCacheMap, nil
}

func (c *Client) SetCurrentMap(ctx context.Context, iqwikiCacheMap map[string]int) error {
	if cache.Global() == nil {
		return fmt.Errorf("redis worker is nil")
	}

	if len(iqwikiCacheMap) == 0 {
		return fmt.Errorf("IqwikiCacheMap is less than 0")
	}

	data, err := json.Marshal(iqwikiCacheMap)
	if err != nil {
		return fmt.Errorf("marshal %+v to json error:%+v", iqwikiCacheMap, err)
	}

	mapKey := fmt.Sprintf("crawler_%s", protocol.PlatformIQWiki)
	cache.Global().Set(ctx, mapKey, data, 0)

	return nil
}

func (c *Client) GetUserActivityList(ctx context.Context, address string) ([]graphqlx.GetUserActivitiesActivitiesByUserActivity, error) {
	var activityList []graphqlx.GetUserActivitiesActivitiesByUserActivity
	userId := common.HexToAddress(address).String()
	offset := 0
	index := 0

	for {
		resp, err := graphqlx.GetUserActivities(ctx, *c.GClient, userId, Limit, offset)
		if err != nil {
			return nil, err
		}
		activityList = append(activityList, resp.ActivitiesByUser...)
		if len(resp.ActivitiesByUser) < Limit {
			break
		}
		index++
		offset = Limit * index
	}

	return activityList, nil
}

func (c *Client) GetUserList(ctx context.Context) ([]graphqlx.GetUsersUsersUser, error) {
	var userList []graphqlx.GetUsersUsersUser
	offset := 0
	index := 0

	for {
		resp, err := graphqlx.GetUsers(ctx, *c.GClient, Limit, offset)
		if err != nil {
			return nil, err
		}
		userList = append(userList, resp.Users...)
		if len(resp.Users) < Limit {
			break
		}
		index++
		offset = Limit * index
	}

	return userList, nil
}

func NewClient() *Client {
	requestURL := &url.URL{
		Scheme: Scheme,
		Host:   Endpoint,
		Path:   Path,
	}
	client := graphql.NewClient(requestURL.String(), http.DefaultClient)

	return &Client{
		GClient: &client,
	}
}
