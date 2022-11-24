package service

import (
	"context"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/dao"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/model"
)

func (s *Service) GetWrapped(c context.Context, request model.GetRequest, wrappedResult *model.WrappedResult) error {

	var err error
	wrappedResult.Social, err = dao.CountSocial(c, request)
	if err != nil {
		return err
	}

	wrappedResult.Search, err = dao.CountSearch(c, request)
	if err != nil {
		return err
	}

	// multiple search counts
	wrappedResult.Search.Count = wrappedResult.Search.Count * 4

	wrappedResult.Gas, err = dao.CountGas(c, request)
	if err != nil {
		return err
	}

	return nil
}
