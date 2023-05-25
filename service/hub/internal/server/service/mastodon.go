package service

import (
	"context"
	"encoding/json"
	"regexp"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/mattn/go-mastodon"
	dbModel "github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/model"
	lop "github.com/samber/lo/parallel"

	"go.uber.org/zap"
)

const (
	SourceName = "Mastodon"
	ThreadSize = 10
	Limit      = 40
)

func (s *Service) GetMastodonContent(c context.Context, request model.GetRequest, client *mastodon.Client) ([]*dbModel.Transaction, error) {
	query := request.Address

	if len(query) == 0 {
		return nil, nil
	}

	var (
		contentList []*mastodon.Status
		locker      sync.Mutex
	)

	pagination := &mastodon.Pagination{
		Limit: Limit,
	}

	if isMastodonUsername(request.Address) {
		accounts, err := client.AccountsSearch(c, query, 1)
		if err != nil {
			return nil, err
		}

		if len(accounts) == 0 {
			return nil, nil
		}

		contentList, _ = client.GetAccountStatuses(c, accounts[0].ID, pagination)
	} else {
		result, err := client.Search(c, query, true)
		if err != nil {
			return nil, err
		}
		if len(result.Hashtags) == 0 {
			return nil, nil
		}

		opt := lop.NewOption().WithConcurrency(ThreadSize)
		lop.ForEach(result.Hashtags, func(tag *mastodon.Tag, i int) {
			name := tag.Name

			list, err := client.GetTimelineHashtag(c, name, false, pagination)
			if err != nil {
				zap.L().Error("get time line hashtag", zap.Error(err), zap.String("tag", name))

				return
			}

			locker.Lock()
			contentList = append(contentList, list...)
			locker.Unlock()
		}, opt)
	}

	transactions, err := s.handleContentList(c, contentList, client)
	if err != nil {
		return nil, err
	}

	// Sort, latest -> oldest
	sort.SliceStable(transactions, func(i, j int) bool {
		return transactions[i].BlockNumber > transactions[j].BlockNumber
	})

	return transactions, nil
}

func (s *Service) handleContentList(c context.Context, contentList []*mastodon.Status, client *mastodon.Client) ([]*dbModel.Transaction, error) {
	var (
		transactions []*dbModel.Transaction
		tx           *dbModel.Transaction
		locker       sync.Mutex
	)

	opt := lop.NewOption().WithConcurrency(ThreadSize)
	lop.ForEach(contentList, func(item *mastodon.Status, i int) {
		status := item

		switch {
		case status.InReplyToID != nil:
			tx = s.buildCommentTransaction(c, status, client)
		case status.Reblog != nil:
			tx = s.buildShareTransaction(c, status)
		default:
			tx = s.buildPostTransaction(c, status)
		}

		if tx == nil {
			return
		}

		locker.Lock()
		transactions = append(transactions, tx)
		locker.Unlock()
	}, opt)

	return transactions, nil
}

func (s *Service) buildPostTransaction(c context.Context, status *mastodon.Status) *dbModel.Transaction {
	post := &metadata.Post{
		CreatedAt:      status.CreatedAt.Format(time.RFC3339),
		Author:         []string{status.Account.Username, status.Account.Acct, status.Account.DisplayName},
		Title:          status.SpoilerText,
		Body:           status.Content,
		Media:          []metadata.Media{{Address: status.Account.Avatar, MimeType: "image/png"}},
		TypeOnPlatform: []string{"post"},
	}

	metadataPost, _ := json.Marshal(post)

	transaction := &dbModel.Transaction{
		// use timestamp as block number, as there is no actual block number on mastodon
		BlockNumber: status.CreatedAt.UnixMilli(),
		Timestamp:   status.CreatedAt,
		// use id as hash, as there is no actual hash on mastodon
		Hash:        strings.ToLower(common.BytesToHash([]byte(status.ID)).String()),
		Network:     SourceName,
		AddressFrom: status.Account.Acct,
		Owner:       status.Account.Acct,
		Tag:         filter.TagSocial,
		Type:        filter.SocialPost,
		Transfers: []dbModel.Transfer{
			{
				Index:       protocol.IndexVirtual,
				AddressFrom: status.Account.Acct,
				Tag:         filter.TagSocial,
				Type:        filter.SocialPost,
				Metadata:    metadataPost,
				RelatedUrls: []string{status.URI},
			},
		},
	}

	return transaction
}

func (s *Service) buildShareTransaction(c context.Context, status *mastodon.Status) *dbModel.Transaction {
	post := &metadata.Post{
		CreatedAt: status.CreatedAt.Format(time.RFC3339),
		Author:    []string{status.Account.Username, status.Account.Acct, status.Account.DisplayName},
		Target: &metadata.Post{
			CreatedAt:      status.Reblog.CreatedAt.Format(time.RFC3339),
			Author:         []string{status.Reblog.Account.Username, status.Reblog.Account.Acct, status.Reblog.Account.DisplayName},
			Title:          status.Reblog.SpoilerText,
			Body:           status.Reblog.Content,
			Media:          []metadata.Media{{Address: status.Reblog.Account.Avatar, MimeType: "image/png"}},
			TypeOnPlatform: []string{"post"},
		},
		Media:          []metadata.Media{{Address: status.Account.Avatar, MimeType: "image/png"}},
		TypeOnPlatform: []string{"share"},
	}

	metadataPost, _ := json.Marshal(post)

	transaction := &dbModel.Transaction{
		// use timestamp as block number, as there is no actual block number on mastodon
		BlockNumber: status.CreatedAt.UnixMilli(),
		Timestamp:   status.CreatedAt,
		// use id as hash, as there is no actual hash on mastodon
		Hash:        strings.ToLower(common.BytesToHash([]byte(status.ID)).String()),
		Network:     SourceName,
		AddressFrom: status.Account.Acct,
		Owner:       status.Account.Acct,
		Tag:         filter.TagSocial,
		Type:        filter.SocialShare,
		Transfers: []dbModel.Transfer{
			{
				Index:       protocol.IndexVirtual,
				AddressFrom: status.Account.Acct,
				Tag:         filter.TagSocial,
				Type:        filter.SocialShare,
				Metadata:    metadataPost,
				RelatedUrls: []string{status.URI},
			},
		},
	}

	return transaction
}

func (s *Service) buildCommentTransaction(c context.Context, status *mastodon.Status, client *mastodon.Client) *dbModel.Transaction {
	target, err := client.GetStatus(c, mastodon.ID(status.InReplyToID.(string)))

	if err != nil || target == nil {
		return nil
	}

	post := &metadata.Post{
		CreatedAt: status.CreatedAt.Format(time.RFC3339),
		Author:    []string{status.Account.Username, status.Account.Acct, status.Account.DisplayName},
		Target: &metadata.Post{
			CreatedAt:      target.CreatedAt.Format(time.RFC3339),
			Author:         []string{target.Account.Username, target.Account.Acct, target.Account.DisplayName},
			Title:          target.SpoilerText,
			Body:           target.Content,
			Media:          []metadata.Media{{Address: target.Account.Avatar, MimeType: "image/png"}},
			TypeOnPlatform: []string{"post"},
		},
		Body:           status.Content,
		Media:          []metadata.Media{{Address: status.Account.Avatar, MimeType: "image/png"}},
		TypeOnPlatform: []string{"comment"},
	}

	metadataPost, _ := json.Marshal(post)

	transaction := &dbModel.Transaction{
		// use timestamp as block number, as there is no actual block number on mastodon
		BlockNumber: status.CreatedAt.UnixMilli(),
		Timestamp:   status.CreatedAt,
		// use id as hash, as there is no actual hash on mastodon
		Hash:        strings.ToLower(common.BytesToHash([]byte(status.ID)).String()),
		Network:     SourceName,
		AddressFrom: status.Account.Acct,
		Owner:       status.Account.Acct,
		Tag:         filter.TagSocial,
		Type:        filter.SocialComment,
		Transfers: []dbModel.Transfer{
			{
				Index:       protocol.IndexVirtual,
				AddressFrom: status.Account.Acct,
				Tag:         filter.TagSocial,
				Type:        filter.SocialComment,
				Metadata:    metadataPost,
				RelatedUrls: []string{status.URI},
			},
		},
	}

	return transaction
}

func isMastodonUsername(username string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9_]{1,30}@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(username)
}
