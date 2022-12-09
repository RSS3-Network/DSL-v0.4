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

	// multiple search counts by 4, as we only started collecting the date in Q4
	wrappedResult.Search.Count *= 4

	wrappedResult.Gas, err = dao.CountGas(c, request)
	if err != nil {
		return err
	}

	wrappedResult.Transaction, err = dao.CountTransaction(c, request)
	if err != nil {
		return err
	}

	wrappedResult.NFT, err = dao.GetNFT(c, request)
	if err != nil {
		return err
	}

	wrappedResult.DApp, err = dao.GetDApp(c, request)
	if err != nil {
		return err
	}
	wrappedResult.DeFi, err = dao.GetDeFi(c, request)
	if err != nil {
		return err
	}

	return nil
}
