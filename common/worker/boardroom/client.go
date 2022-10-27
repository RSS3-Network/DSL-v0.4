package boardroom

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	"github.com/naturalselectionlabs/pregod/common/worker/boardroom/model"
	"go.uber.org/zap"
)

const (
	BaseUrl = "https://api.boardroom.info/v1"
	Key     = "e44e095b44cde53fe00f79000c554c72"
)

type Client struct {
	boardroomClient *resty.Client
}

type GetProposalsByRefIdResponse struct {
	Data model.Proposal `json:"data"`
}

func (c *Client) GetProposalsByRefId(ctx context.Context, refId string) (*GetProposalsByRefIdResponse, error) {
	var response *GetProposalsByRefIdResponse
	_, err := c.boardroomClient.R().
		SetResult(&response).
		SetPathParams(map[string]string{
			"refId": refId,
		}).
		SetContext(ctx).
		Get("/proposals/{refId}")
	if err != nil {
		loggerx.Global().Named("GetProposalsByRefId").Warn("unable to get boardroom data", zap.Error(err))
		return nil, err
	}
	return response, nil
}

type GetProposalsByAddressResponse struct {
	Data   []model.Proposal `json:"data"`
	Cursor string           `json:"nextCursor"`
}

func (c *Client) GetProposalsByAddress(ctx context.Context, address string) ([]model.Proposal, error) {
	proposals := make([]model.Proposal, 0)
	param := map[string]string{
		"proposer": address,
	}

	for {
		var response *GetProposalsByAddressResponse
		_, err := c.boardroomClient.R().
			SetResult(&response).
			SetQueryParams(param).
			SetContext(ctx).
			Get("/proposals")
		if err != nil {
			loggerx.Global().Named("GetProposalsByAddress").Warn("unable to get boardroom data", zap.Error(err))
			return nil, err
		}

		if response == nil {
			loggerx.Global().Named("GetProposalsByAddress").Warn("response is nil")
			break
		}

		proposals = append(proposals, response.Data...)

		if len(response.Cursor) == 0 {
			break
		}
		param["cursor"] = response.Cursor
	}

	return proposals, nil
}

type GetVotesByAddressResponse struct {
	Data   []model.Vote `json:"data"`
	Cursor string       `json:"nextCursor"`
}

func (c *Client) GetVotesByAddress(ctx context.Context, address string) ([]model.Vote, error) {
	votes := make([]model.Vote, 0)
	param := map[string]string{
		"addresses": address,
	}

	for {
		var response *GetVotesByAddressResponse
		_, err := c.boardroomClient.R().
			SetResult(&response).
			SetQueryParams(param).
			SetContext(ctx).
			Get("/getVotesByAddresses")
		if err != nil {

			loggerx.Global().Named("GetVotesByAddress").Warn("unable to get boardroom data", zap.Error(err))
			return nil, err
		}
		if response == nil {
			loggerx.Global().Named("GetVotesByAddress").Warn("response is nil")
			break
		}
		votes = append(votes, response.Data...)

		if len(response.Cursor) == 0 {
			break
		}
		param["cursor"] = response.Cursor
	}

	return votes, nil
}

func NewClient() *Client {
	client := resty.New()
	client.SetBaseURL(BaseUrl)
	client.SetQueryParams(map[string]string{
		"key": Key,
	})
	return &Client{
		boardroomClient: client,
	}
}
