package service

import (
	"context"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/dao"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/model"
)

func (s *Service) GetWrapped(c context.Context, request model.GetRequest, wrappedResult *model.WrappedResult) error {
	result := model.SocialResult{}

	var err error

	result, err = dao.CountSocial(c, request)
	if err != nil {
		return err
	}

	wrappedResult.Social = result
	return nil
}
