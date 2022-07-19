package lens

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/hasura/go-graphql-client"
	"github.com/lib/pq"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/common/utils/opentelemetry"
	lensClient "github.com/naturalselectionlabs/pregod/common/worker/lens"
	graphqlx "github.com/naturalselectionlabs/pregod/common/worker/lens/graphql"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"

	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	Name = protocol.PlatfromLens
)

var _ worker.Worker = (*service)(nil)

type service struct {
	databaseClient *gorm.DB
	lensClient     *lensClient.Client
}

func New(databaseClient *gorm.DB) worker.Worker {
	return &service{
		databaseClient: databaseClient,
		lensClient:     lensClient.NewClient(),
	}
}

func (s *service) Name() string {
	return Name
}

func (s *service) Networks() []string {
	return []string{
		protocol.NetworkPolygon,
	}
}

func (s *service) Initialize(ctx context.Context) error {
	return nil
}

func (s *service) Jobs() []worker.Job {
	return nil
}

func (s *service) Handle(ctx context.Context, message *protocol.Message, transactions []model.Transaction) (data []model.Transaction, err error) {
	tracer := otel.Tracer("lens_worker")
	_, trace := tracer.Start(ctx, "lens_worker:Handle")

	defer func() { opentelemetry.Log(trace, transactions, data, err) }()

	// handle posts
	transactions, err = s.handlePost(ctx, message, transactions)
	if err != nil {
		logrus.Error(err)
	}

	// handle profiles
	transactions, err = s.handleProfile(ctx, message, transactions)
	if err != nil {
		logrus.Error(err)
	}

	// handle follows
	transactions, err = s.handleFollow(ctx, message, transactions)
	if err != nil {
		logrus.Error(err)
	}

	return transactions, nil
}

func (s *service) handlePost(ctx context.Context, message *protocol.Message, transactions []model.Transaction) (data []model.Transaction, err error) {
	tracer := otel.Tracer("lens_worker")
	_, trace := tracer.Start(ctx, "lens_worker:handlePost")

	defer func() { opentelemetry.Log(trace, transactions, data, err) }()

	// read the last cursor value from the database
	var lastLens model.Transaction
	var cursor string

	if err := s.databaseClient.Model((*model.Transaction)(nil)).
		Where("address_from = ?", message.Address).
		Where("source = ?", s.Name()).
		Where("type = ANY(?)", pq.StringArray{filter.SocialPost, filter.SocialComment, filter.SocialShare}).
		Order("timestamp DESC").First(&lastLens).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return transactions, err
	}

	// if it's a new address, the cursor will be empty
	if len(lastLens.Hash) == 0 {
		cursor = "{}"
	} else {
		cursor = fmt.Sprintf(`{"entityIdentifier":"%s","timestamp":%d,"cursorDirection":"BEFORE"}`, lastLens.Hash, lastLens.Timestamp.Unix())
	}

	options := lensClient.Options{
		Address: graphql.String(message.Address),
		Cursor:  graphql.String(cursor),
	}

	// handle publications
	result, err := s.lensClient.GetAllPublicationsByAddress(ctx, &options)
	if err != nil {
		return transactions, err
	}

	// returns when no results are found
	if len(result) == 0 {
		return transactions, nil
	}

	for _, publication := range result {
		sourceData, err := json.Marshal(publication)
		if err != nil {
			continue
		}

		post := &metadata.Post{
			TypeOnPlatform: []string{string(publication.Type)},
			Body:           string(publication.Metadata.Description),
			Title:          string(publication.Metadata.Description),
		}

		formatPostMedia(publication.Metadata.Media, post)

		if publication.Type != "Post" {

			targetContent := &metadata.Post{
				TypeOnPlatform: []string{string(publication.Target.Type)},
				Body:           string(publication.Target.Metadata.Description),
				Title:          string(publication.Target.Metadata.Description),
			}

			formatPostMedia(publication.Metadata.Media, targetContent)

			post.Target = targetContent
		}

		rawMetadata, err := json.Marshal(post)
		if err != nil {
			continue
		}

		transactions = append(transactions, model.Transaction{
			Timestamp:   publication.CreatedAt,
			Hash:        string(publication.ID),
			AddressFrom: message.Address,
			AddressTo:   "",
			Tag:         filter.TagSocial,
			Type:        getType(string(publication.Type)),
			Network:     message.Network,
			Platform:    string(publication.Platform),
			Source:      s.Name(),
			Transfers: []model.Transfer{
				{
					Tag:             filter.TagSocial,
					Type:            getType(string(publication.Type)),
					TransactionHash: string(publication.ID),
					Timestamp:       publication.CreatedAt,
					AddressFrom:     message.Address,
					AddressTo:       "",
					Metadata:        rawMetadata,
					Network:         protocol.NetworkPolygon,
					Platform:        string(publication.Platform),
					Source:          s.Name(),
					RelatedUrls:     []string{string(publication.RelatedURL)},
					SourceData:      sourceData,
				},
			},
		})
	}

	return transactions, nil
}

func (s *service) handleProfile(ctx context.Context, message *protocol.Message, transactions []model.Transaction) (data []model.Transaction, err error) {
	tracer := otel.Tracer("lens_worker")
	_, trace := tracer.Start(ctx, "lens_worker:handleProfile")

	defer func() { opentelemetry.Log(trace, transactions, data, err) }()

	options := lensClient.Options{
		Address: graphql.String(message.Address),
	}

	profiles, err := s.lensClient.GetProfiles(ctx, &options)
	if err != nil {
		return transactions, err
	}

	// returns when no results are found
	if len(profiles) == 0 {
		return transactions, nil
	}

	for _, profile := range profiles {
		// ignore non-default profiles
		if !profile.IsDefault {
			continue
		}

		profileTemp := &model.Profile{
			Address:  message.Address,
			Network:  message.Network,
			Platform: s.Name(),
			Source:   s.Name(),
			Name:     string(profile.Name),
			Handle:   string(profile.Handle),
			Bio:      string(profile.Bio),
		}

		formatProfileMedia(&profile, profileTemp)

		// use created_at from the previous profile
		var existingProfile model.Profile
		if err := s.databaseClient.Model(&model.Profile{}).
			Where("address = ?", message.Address).
			Where("handle = ?", string(profile.Handle)).
			Where("platform = ?", s.Name()).First(&existingProfile).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
			profileTemp.CreatedAt = existingProfile.CreatedAt
		}

		// this will also update the profile in the database
		// if the default profile has changed
		s.databaseClient.Model(&model.Profile{}).Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(&profileTemp)

		rawMetadata, err := json.Marshal(profileTemp)
		if err != nil {
			continue
		}

		transactions = append(transactions, model.Transaction{
			Hash:        fmt.Sprintf("%s-%s", string(profile.ID), string(profile.Handle)),
			AddressFrom: message.Address,
			AddressTo:   "",
			Timestamp:   profileTemp.CreatedAt,
			Tag:         filter.TagSocial,
			Type:        filter.SocialProfile,
			Network:     message.Network,
			Platform:    s.Name(),
			Source:      s.Name(),
			Transfers: []model.Transfer{
				{
					Tag:             filter.TagSocial,
					Type:            filter.SocialProfile,
					TransactionHash: fmt.Sprintf("%s-%s", string(profile.ID), string(profile.Handle)),
					AddressFrom:     message.Address,
					AddressTo:       "",
					Metadata:        rawMetadata,
					Network:         protocol.NetworkPolygon,
					Platform:        s.Name(),
					Source:          s.Name(),
					RelatedUrls:     []string{string(profile.Metadata), fmt.Sprintf("https://www.lensfrens.xyz/%s", string(profile.Handle))},
				},
			},
		})
	}

	return transactions, nil
}

func (s *service) handleFollow(ctx context.Context, message *protocol.Message, transactions []model.Transaction) (data []model.Transaction, err error) {
	tracer := otel.Tracer("lens_worker")
	_, trace := tracer.Start(ctx, "lens_worker:handlePost")

	defer func() { opentelemetry.Log(trace, transactions, data, err) }()

	// read the last cursor value from the database
	var followCount int64
	var cursor string

	s.databaseClient.Model((*model.Transaction)(nil)).
		Where("address_from = ?", message.Address).
		Where("source = ?", s.Name()).
		Where("type = ?", filter.SocialFollow).
		Count(&followCount)

	cursor = fmt.Sprintf(`{"offset":"%d"}`, followCount)

	options := lensClient.Options{
		Address: graphql.String(message.Address),
		Cursor:  graphql.String(cursor),
	}

	// handle publications
	result, err := s.lensClient.GetFollowings(ctx, &options)
	if err != nil {
		return transactions, err
	}

	// returns when no results are found
	if len(result) == 0 {
		return transactions, nil
	}

	for _, profile := range result {
		// follow is essentially targeting a profile
		profileTemp := &model.Profile{
			Address:  string(profile.OwnedBy),
			Network:  message.Network,
			Platform: s.Name(),
			Source:   s.Name(),
			Name:     string(profile.Name),
			Handle:   string(profile.Handle),
			Bio:      string(profile.Bio),
		}

		formatProfileMedia(&profile, profileTemp)

		rawMetadata, err := json.Marshal(profileTemp)
		if err != nil {
			continue
		}

		transactions = append(transactions, model.Transaction{
			Hash:        fmt.Sprintf("%s-%s", message.Address, string(profile.Handle)),
			AddressFrom: message.Address,
			AddressTo:   "",
			Timestamp:   time.Now(),
			Tag:         filter.TagSocial,
			Type:        filter.SocialFollow,
			Network:     message.Network,
			Platform:    s.Name(),
			Source:      s.Name(),
			Transfers: []model.Transfer{
				{
					Tag:             filter.TagSocial,
					Type:            filter.SocialFollow,
					TransactionHash: fmt.Sprintf("%s-%s", message.Address, string(profile.Handle)),
					AddressFrom:     message.Address,
					AddressTo:       "",
					Metadata:        rawMetadata,
					Network:         protocol.NetworkPolygon,
					Platform:        s.Name(),
					Source:          s.Name(),
				},
			},
		})
	}

	return transactions, nil
}

func getType(pubType string) string {
	switch pubType {
	case "Post":
		return filter.SocialPost
	case "Comment":
		return filter.SocialComment
	case "Mirror":
		return filter.SocialShare
	}

	return filter.SocialPost
}

func formatPostMedia(mediaList []graphqlx.Media, post *metadata.Post) {
	for _, media := range mediaList {
		post.Media = append(post.Media, metadata.Media{
			Address:  string(media.Original.URL),
			MimeType: string(media.Original.MimeType),
		})
	}
}

func formatProfileMedia(profile *graphqlx.Profile, result *model.Profile) {
	// either `NftImage` or `MediaSet`
	switch profile.Picture.Type {
	case "NftImage":
		result.ProfileUris = append(result.ProfileUris, string(profile.Picture.NFTImage.Uri))
	case "MediaSet":
		result.ProfileUris = append(result.ProfileUris, string(profile.Picture.MediaSet.Original.URL))
	}

	switch profile.CoverPicture.Type {
	case "NftImage":
		result.BannerUris = append(result.BannerUris, string(profile.CoverPicture.NFTImage.Uri))
	case "MediaSet":
		result.BannerUris = append(result.BannerUris, string(profile.CoverPicture.MediaSet.Original.URL))
	}
}
