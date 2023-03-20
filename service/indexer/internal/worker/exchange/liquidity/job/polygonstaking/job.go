package polygonstaking

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
)

const (
	Key      = "polygon_staking_validators"
	Endpoint = "https://staking-api.polygon.technology"
)

var _ worker.Job = (*Job)(nil)

type Job struct {
	redisClient *redis.Client
}

func (j *Job) Name() string {
	return "polygon staking validators"
}

func (j *Job) Spec() string {
	return "@every 10s"
}

func (j *Job) Timeout() time.Duration {
	return 5 * time.Second
}

func (j *Job) Run(_ worker.RenewalFunc) error {
	var (
		ctx                  = context.Background()
		total, offset, limit = 0, 0, 100
	)

	for first := true; first || total > offset+limit; first = false {
		requestURL := fmt.Sprintf("%s/api/v2/validators?limit=%d&offset=%d&sortBy=delegatedStake", Endpoint, limit, offset)

		httpRequest, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL, nil)
		if err != nil {
			return fmt.Errorf("new request: %w", err)
		}

		httpResponse, err := http.DefaultClient.Do(httpRequest)
		if err != nil {
			return fmt.Errorf("http client do request: %w", err)
		}

		var response Response
		if err := json.NewDecoder(httpResponse.Body).Decode(&response); err != nil {
			return err
		}

		if !response.Success {
			return fmt.Errorf("response return an error: %w", errors.New(response.Status))
		}

		offset += len(response.Result)

		for _, result := range response.Result {
			data, err := json.Marshal(result)
			if err != nil {
				return fmt.Errorf("marshal result: %w", err)
			}

			j.redisClient.HSet(ctx, Key, strings.ToLower(result.ContractAddress.String()), data)
		}
	}

	return nil
}

func New(redisClient *redis.Client) worker.Job {
	return &Job{
		redisClient: redisClient,
	}
}
