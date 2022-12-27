package service

import (
	"context"
	"fmt"
	"time"

	"github.com/naturalselectionlabs/pregod/common/database/model/social"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	"github.com/naturalselectionlabs/pregod/common/worker/crossbell"
	"github.com/naturalselectionlabs/pregod/common/worker/lens"
	"github.com/naturalselectionlabs/pregod/common/worker/name_service/avvy"
	"github.com/naturalselectionlabs/pregod/common/worker/name_service/ens"
	"github.com/naturalselectionlabs/pregod/common/worker/name_service/spaceid"
	"github.com/naturalselectionlabs/pregod/common/worker/name_service/unstoppable"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/dao"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/model"
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
		lensClient := lens.New()
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
