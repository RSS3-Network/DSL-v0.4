package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gabriel-vasile/mimetype"
	"github.com/mattn/go-mastodon"
	dbModel "github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/handler/maspool"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/model"
	lop "github.com/samber/lo/parallel"

	"go.uber.org/zap"
)

const (
	SourceName = "Mastodon"
	Limit      = 40
)

func (s *Service) GetMastodonContent(c context.Context, request model.GetRequest) ([]*dbModel.Transaction, error) {
	query := request.Address

	if len(query) == 0 {
		return nil, nil
	}

	var (
		contentList []*mastodon.Status
		locker      sync.Mutex
		err         error
		instance    *maspool.Instance
	)

	pagination := &mastodon.Pagination{
		Limit: Limit,
	}

	if isMastodonUsername(request.Address) {
		var accounts []*mastodon.Account

		instance = s.mastodonPool.GetAvailableInstance(2)
		if instance == nil {
			return nil, fmt.Errorf("mastodon api rate limit, please try again after %s", maspool.GetNext5MinuteTime().Format(time.RFC3339))
		}
		loggerx.Global().Info("mastodon instance", zap.String("name", instance.Client.Config.Server), zap.Int("rateLimit", instance.RateLimit), zap.Time("resetTime", instance.LastReset))
		accounts, err = instance.Client.AccountsSearch(context.Background(), query, 1)

		if err != nil {
			return nil, err
		}

		if len(accounts) == 0 {
			return nil, nil
		}

		// different id in different client
		contentList, err = instance.Client.GetAccountStatuses(c, accounts[0].ID, pagination)

		if err != nil {
			return nil, err
		}

	} else {
		var result *mastodon.Results

		instance = s.mastodonPool.GetAvailableInstance(1)
		if instance == nil {
			return nil, fmt.Errorf("mastodon api rate limit, please try this request again after %s", maspool.GetNext5MinuteTime().Format(time.RFC3339))
		}
		loggerx.Global().Info("mastodon instance", zap.String("name", instance.Client.Config.Server), zap.Int("rateLimit", instance.RateLimit), zap.Time("resetTime", instance.LastReset))
		result, err = instance.Client.Search(c, query, true)

		if err != nil {
			return nil, err
		}

		if len(result.Hashtags) == 0 {
			return nil, nil
		}

		instance = s.mastodonPool.GetAvailableInstance(len(result.Hashtags))
		if instance == nil {
			return nil, fmt.Errorf("mastodon api rate limit, please try this request again after %s", maspool.GetNext5MinuteTime().Format(time.RFC3339))
		}
		loggerx.Global().Info("mastodon instance", zap.String("name", instance.Client.Config.Server), zap.Int("rateLimit", instance.RateLimit), zap.Time("resetTime", instance.LastReset))

		opt := lop.NewOption().WithConcurrency(len(result.Hashtags))
		lop.ForEach(result.Hashtags, func(tag *mastodon.Tag, i int) {
			name := tag.Name

			var list []*mastodon.Status

			list, err = instance.Client.GetTimelineHashtag(c, name, false, pagination)

			if err != nil {
				zap.L().Error("get time line hashtag", zap.Error(err), zap.String("tag", name))

				return
			}

			locker.Lock()
			contentList = append(contentList, list...)
			locker.Unlock()
		}, opt)
	}

	var comments int
	for _, content := range contentList {
		if content.InReplyToID != nil {
			comments++
		}
	}

	if instance.RateLimit < comments {
		return nil, fmt.Errorf("mastodon api rate limit, please try this request again after %s", maspool.GetNext5MinuteTime().Format(time.RFC3339))
	}

	transactions, err := s.handleContentList(c, contentList, instance)
	if err != nil {
		return nil, err
	}

	// Sort, latest -> oldest
	sort.SliceStable(transactions, func(i, j int) bool {
		return transactions[i].BlockNumber > transactions[j].BlockNumber
	})

	return transactions, nil
}

func (s *Service) handleContentList(c context.Context, contentList []*mastodon.Status, instance *maspool.Instance) ([]*dbModel.Transaction, error) {
	var (
		transactions []*dbModel.Transaction
		tx           *dbModel.Transaction
		locker       sync.Mutex
	)

	opt := lop.NewOption().WithConcurrency(100)
	lop.ForEach(contentList, func(item *mastodon.Status, i int) {
		status := item

		switch {
		case status.InReplyToID != nil:
			tx = s.buildCommentTransaction(c, status, instance)
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

func (s *Service) buildPostTransaction(_ context.Context, status *mastodon.Status) *dbModel.Transaction {
	post := &metadata.Post{
		CreatedAt:      status.CreatedAt.Format(time.RFC3339),
		Author:         []string{status.Account.Username, status.Account.Acct, status.Account.DisplayName},
		Title:          status.SpoilerText,
		Body:           status.Content,
		Media:          []metadata.Media{{Address: status.Account.Avatar, MimeType: "image/png"}},
		TypeOnPlatform: []string{"post"},
	}

	s.buildPostAttachments(post, status)

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

func (s *Service) buildShareTransaction(_ context.Context, status *mastodon.Status) *dbModel.Transaction {
	target := &metadata.Post{
		CreatedAt:      status.Reblog.CreatedAt.Format(time.RFC3339),
		Author:         []string{status.Reblog.Account.Username, status.Reblog.Account.Acct, status.Reblog.Account.DisplayName},
		Title:          status.Reblog.SpoilerText,
		Body:           status.Reblog.Content,
		Media:          []metadata.Media{{Address: status.Reblog.Account.Avatar, MimeType: "image/png"}},
		TypeOnPlatform: []string{"post"},
	}

	s.buildPostAttachments(target, status.Reblog)

	post := &metadata.Post{
		CreatedAt:      status.CreatedAt.Format(time.RFC3339),
		Author:         []string{status.Account.Username, status.Account.Acct, status.Account.DisplayName},
		Target:         target,
		Media:          []metadata.Media{{Address: status.Account.Avatar, MimeType: "image/png"}},
		TypeOnPlatform: []string{"share"},
	}

	s.buildPostAttachments(post, status)

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

func (s *Service) buildCommentTransaction(c context.Context, status *mastodon.Status, instance *maspool.Instance) *dbModel.Transaction {
	var (
		target *mastodon.Status
		err    error
	)

	target, err = instance.Client.GetStatus(c, mastodon.ID(status.InReplyToID.(string)))

	if err != nil || target == nil {
		return nil
	}

	targetPost := &metadata.Post{
		CreatedAt:      target.CreatedAt.Format(time.RFC3339),
		Author:         []string{target.Account.Username, target.Account.Acct, target.Account.DisplayName},
		Title:          target.SpoilerText,
		Body:           target.Content,
		Media:          []metadata.Media{{Address: target.Account.Avatar, MimeType: "image/png"}},
		TypeOnPlatform: []string{"post"},
	}

	s.buildPostAttachments(targetPost, target)

	post := &metadata.Post{
		CreatedAt:      status.CreatedAt.Format(time.RFC3339),
		Author:         []string{status.Account.Username, status.Account.Acct, status.Account.DisplayName},
		Target:         targetPost,
		Body:           status.Content,
		Media:          []metadata.Media{{Address: status.Account.Avatar, MimeType: "image/png"}},
		TypeOnPlatform: []string{"comment"},
	}

	s.buildPostAttachments(post, status)

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

func (s *Service) buildPostAttachments(post *metadata.Post, status *mastodon.Status) {
	var locker sync.Mutex

	opt := lop.NewOption().WithConcurrency(10)

	lop.ForEach(status.MediaAttachments, func(item mastodon.Attachment, i int) {
		attachment := item
		mimeType, err := detectMimeType(attachment.URL)
		if err != nil {
			loggerx.Global().Warn("detect mime type error", zap.String("url", attachment.URL), zap.Error(err))
			return
		}

		locker.Lock()
		post.Media = append(post.Media, metadata.Media{Address: attachment.URL, MimeType: mimeType})
		locker.Unlock()
	}, opt)
}

func isMastodonUsername(username string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9_]{1,30}@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(username)
}

func detectMimeType(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("http request error %w", err)
	}
	defer resp.Body.Close()
	mimeType, err := mimetype.DetectReader(resp.Body)
	if err != nil {
		return "", fmt.Errorf("detect mime type error %w", err)
	}

	return mimeType.String(), nil
}
