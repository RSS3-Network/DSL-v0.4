package service

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	kurora "github.com/naturalselectionlabs/kurora/client"
	"github.com/naturalselectionlabs/pregod/common/database/model/social"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	"github.com/naturalselectionlabs/pregod/common/worker/crossbell"
	"github.com/naturalselectionlabs/pregod/common/worker/lens"
	"github.com/naturalselectionlabs/pregod/common/worker/name_service"
	"github.com/naturalselectionlabs/pregod/common/worker/name_service/avvy"
	"github.com/naturalselectionlabs/pregod/common/worker/name_service/ens"
	"github.com/naturalselectionlabs/pregod/common/worker/name_service/spaceid"
	"github.com/naturalselectionlabs/pregod/common/worker/name_service/unstoppable"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/dao"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/model"
	"github.com/samber/lo"
	lop "github.com/samber/lo/parallel"
	"github.com/sirupsen/logrus"

	"go.uber.org/zap"
)

var ProfilePlatformList = []string{
	protocol.PlatformEns,
	protocol.PlatformLens,
	protocol.PlatformCrossbell,
	protocol.PlatformSpaceID,
	protocol.PlatformUnstoppableDomain,
	protocol.PlatformAvvy,
}

var ProfileLockKey = "profile:%v:%v"

func (s *Service) GetProfiles(c context.Context, request model.GetRequest) ([]*social.Profile, error) {
	m := make(map[string]*social.Profile)
	result := make([]*social.Profile, 0)

	profiles, _ := dao.GetProfiles(c, request)

	for _, profile := range profiles {
		m[profile.Platform] = profile
		result = append(result, profile)
	}

	lop.ForEach(ProfilePlatformList, func(platform string, i int) {
		if profile, ok := m[platform]; ok {
			go func() {
				if profile.UpdatedAt.After(time.Now().Add(-2 * time.Minute)) {
					return
				}

				if _, err := s.GetProfilesFromPlatform(context.Background(), platform, request.Address); err != nil {
					logrus.Errorf("[profile] getProfilesFromPlatform error, %v", err)
				}
			}()

			return
		}

		list, err := s.GetProfilesFromPlatform(c, platform, request.Address)
		if err == nil && len(list) > 0 {
			result = append(result, list...)
		}
	})

	return result, nil
}

func (s *Service) BatchGetProfiles(c context.Context, request model.BatchGetProfilesRequest) ([]*social.Profile, error) {
	profiles, err := dao.BatchGetProfiles(c, request)
	if err != nil {
		loggerx.Global().Error("batch get profiles error", zap.Error(err))

		return nil, err
	}

	go func() {
		lop.ForEach(request.Address, func(address string, i int) {
			for _, platform := range ProfilePlatformList {
				_, _ = s.GetProfilesFromPlatform(context.Background(), platform, address)
			}
		})
	}()

	return profiles, nil
}

func (s *Service) GetProfilesFromPlatform(c context.Context, platform, address string) ([]*social.Profile, error) {
	lockKey := fmt.Sprintf(ProfileLockKey, address, platform)
	if !s.employer.DoLock(lockKey, 2*time.Minute) {
		return nil, fmt.Errorf("%v lock", lockKey)
	}

	cctx, cancel := context.WithCancel(context.Background())
	go func(cctx context.Context) {
		for {
			time.Sleep(time.Second)
			if err := s.employer.Renewal(cctx, lockKey, time.Minute); err != nil {
				return
			}
		}
	}(cctx)

	defer func() {
		cancel()
		s.employer.UnLock(lockKey)
	}()

	var profile *social.Profile
	var profiles []*social.Profile
	var err error

	switch platform {
	case protocol.PlatformEns:
		ensClient := ens.New()
		profile, err = ensClient.GetProfile(address)
	case protocol.PlatformLens:
		var lensClient *lens.Client

		lensClient, err = lens.New(s.kuroraClient)
		if err != nil {
			return nil, err
		}

		profiles, err = lensClient.BatchGetProfiles(c, address)
	case protocol.PlatformCrossbell:
		var csbClient *crossbell.Client
		csbClient, err = crossbell.New()
		if err == nil {
			profiles, err = csbClient.GetProfile(address)
		}
	case protocol.PlatformSpaceID:
		spaceidClient := spaceid.New()
		profile, err = spaceidClient.GetProfile(address)
	case protocol.PlatformUnstoppableDomain:
		unstoppableClient := unstoppable.New()
		profile, err = unstoppableClient.GetProfile(address)
	case protocol.PlatformAvvy:
		avvyClient := avvy.New()
		profile, err = avvyClient.GetProfile(address)
	}

	if err != nil {
		loggerx.Global().Error("get profiles error", zap.Error(err))
		return nil, err
	}

	if profile != nil {
		profiles = append(profiles, profile)
	}

	if len(profiles) == 0 {
		return nil, nil
	}

	go dao.UpsertProfiles(c, profiles)

	return profiles, err
}

func (s *Service) GetKuroraProfiles(c context.Context, request model.GetRequest) ([]*social.Profile, error) {
	var err error
	result := make([]*social.Profile, 0)

	if name_service.IsEvmValidAddress(request.Address) && !strings.EqualFold(ethereum.AddressGenesis.String(), request.Address) {
		result, err = s.GetKuroraAddress(c, request)
		if err != nil {
			return nil, err
		}
	} else {
		domainsQuery := kurora.DatasetDomainQuery{
			Handle: lo.ToPtr(strings.ToLower(request.Address)),
		}

		profiles, err := s.kuroraClient.FetchDatasetDomains(c, domainsQuery)
		if err != nil {
			loggerx.Global().Error("get profiles error", zap.Error(err), zap.String("address", strings.ToLower(request.Address)))
			return nil, err
		}

		uniqueAddress := make(map[common.Address]bool)
		for _, profile := range profiles {
			if profile.ReverseAddress != profile.Address {
				var expiredAt *time.Time
				if !profile.ExpiredAt.IsZero() {
					expiredAt = &profile.ExpiredAt
				}
				result = append(result, &social.Profile{
					Address:     strings.ToLower(profile.Address.String()),
					Network:     profile.Network,
					Platform:    profile.Platform,
					Name:        profile.Name,
					Handle:      profile.Handle,
					Bio:         profile.Bio,
					URL:         profile.URL,
					ExpireAt:    expiredAt,
					ProfileUris: profile.ProfileUris,
					BannerUris:  profile.BannerUris,
					SocialUris:  profile.SocialUris,
					Source:      profile.Platform,
				})
			}

			if _, ok := uniqueAddress[profile.Address]; !ok {
				uniqueAddress[profile.Address] = true

				if profile.Address == ethereum.AddressGenesis {
					continue
				}
				request.Address = profile.Address.String()

				tempRes, err := s.GetKuroraAddress(c, request)
				if err != nil {
					return nil, err
				}

				result = append(result, tempRes...)
			}
		}
	}

	// kurora can not fetch the whole handle such as the avvy domain which sets enhanced privacy
	// https://avvy.domains/docs/privacy-features-registrations/
	if len(result) == 0 && !name_service.IsEvmValidAddress(request.Address) {
		res := name_service.ReverseResolveAll(c, request.Address, false)
		if len(res.Address) == 0 || !name_service.IsEvmValidAddress(res.Address) || strings.EqualFold(res.Address, ethereum.AddressGenesis.String()) {
			return result, nil
		}

		request.Address = res.Address

		result, err = s.GetKuroraAddress(c, request)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

func (s *Service) GetKuroraAddress(c context.Context, request model.GetRequest) ([]*social.Profile, error) {
	result := make([]*social.Profile, 0)

	domainsQuery := kurora.DatasetDomainQuery{
		ReverseAddress: lo.ToPtr(common.HexToAddress(request.Address)),
	}

	if len(request.Network) > 0 {
		for i, v := range request.Network {
			request.Network[i] = strings.ToLower(v)
		}

		domainsQuery.NetworkList = request.Network
	}

	if len(request.Platform) > 0 {
		domainsQuery.PlatformList = request.Platform
	}

	profiles, err := s.kuroraClient.FetchDatasetDomains(c, domainsQuery)
	if err != nil {
		loggerx.Global().Error("get profiles error", zap.Error(err), zap.String("address", strings.ToLower(request.Address)))

		return result, err
	}

	for _, profile := range profiles {
		var expiredAt *time.Time
		if !profile.ExpiredAt.IsZero() {
			expiredAt = &profile.ExpiredAt
		}
		result = append(result, &social.Profile{
			Address:     strings.ToLower(request.Address),
			Network:     profile.Network,
			Platform:    profile.Platform,
			Name:        profile.Name,
			Handle:      profile.Handle,
			Bio:         profile.Bio,
			URL:         profile.URL,
			ExpireAt:    expiredAt,
			ProfileUris: profile.ProfileUris,
			BannerUris:  profile.BannerUris,
			SocialUris:  profile.SocialUris,
			Source:      profile.Platform,
		})
	}

	return result, nil
}

func (s *Service) BatchKuroraProfiles(c context.Context, request model.BatchGetProfilesRequest) ([]*social.Profile, error) {
	result := make([]*social.Profile, 0)

	addresses := make([]common.Address, 0)
	handles := make([]string, 0)
	handleMap := make(map[string]struct{}, 0)

	query := kurora.DatasetDomainQuery{
		Limit: lo.ToPtr(10 * model.DefaultLimit),
	}

	for _, address := range request.Address {
		if !strings.EqualFold(ethereum.AddressGenesis.String(), address) && name_service.IsEvmValidAddress(address) {
			addresses = append(addresses, common.HexToAddress(address))
		} else if !strings.EqualFold("", address) && strings.Contains(address, ".") {
			handles = append(handles, address)
			handleMap[address] = struct{}{}
		}
	}

	if len(request.Network) > 0 {
		for i, v := range request.Network {
			request.Network[i] = strings.ToLower(v)
		}
		query.NetworkList = request.Network
	}

	if len(request.Platform) > 0 {
		query.PlatformList = request.Platform
	}

	// deal with handles
	if len(handles) > 0 {
		query.HandleList = handles

		handleProfiles, err := s.kuroraClient.FetchDatasetDomains(c, query)
		if err != nil {
			loggerx.Global().Error("batch get profiles by handles error", zap.Error(err))
		}

		for _, handleProfile := range handleProfiles {
			if !strings.EqualFold(ethereum.AddressGenesis.String(), handleProfile.Address.String()) {
				addresses = append(addresses, handleProfile.Address)

				var expiredAt *time.Time
				if !handleProfile.ExpiredAt.IsZero() {
					expiredAt = &handleProfile.ExpiredAt
				}
				result = append(result, &social.Profile{
					Address:     strings.ToLower(handleProfile.Address.String()),
					Network:     handleProfile.Network,
					Platform:    handleProfile.Platform,
					Name:        handleProfile.Name,
					Handle:      handleProfile.Handle,
					Bio:         handleProfile.Bio,
					URL:         handleProfile.URL,
					ExpireAt:    expiredAt,
					ProfileUris: handleProfile.ProfileUris,
					BannerUris:  handleProfile.BannerUris,
					SocialUris:  handleProfile.SocialUris,
					Source:      handleProfile.Platform,
				})
			}
		}
	}

	if len(addresses) > 0 {
		query.HandleList = []string{}
		query.ReverseAddressList = addresses

		profiles, err := s.kuroraClient.FetchDatasetDomains(c, query)
		if err != nil {
			loggerx.Global().Error("batch get profiles error", zap.Error(err))

			return result, err
		}

		for _, profile := range profiles {
			if _, ok := handleMap[profile.Handle]; ok {
				continue
			}
			var expiredAt *time.Time
			if !profile.ExpiredAt.IsZero() {
				expiredAt = &profile.ExpiredAt
			}
			result = append(result, &social.Profile{
				Address:     strings.ToLower(profile.ReverseAddress.String()),
				Network:     profile.Network,
				Platform:    profile.Platform,
				Name:        profile.Name,
				Handle:      profile.Handle,
				Bio:         profile.Bio,
				URL:         profile.URL,
				ExpireAt:    expiredAt,
				ProfileUris: profile.ProfileUris,
				BannerUris:  profile.BannerUris,
				SocialUris:  profile.SocialUris,
				Source:      profile.Platform,
			})
		}
	}

	return result, nil
}
