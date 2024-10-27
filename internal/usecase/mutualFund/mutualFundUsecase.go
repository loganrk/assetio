package mutualFund

import (
	"assetio/internal/domain"
	"assetio/internal/port"
	"context"
)

type mutualFundUsecase struct {
	logger port.Logger
	mysql  port.RepositoryStore
}

func New(loggerIns port.Logger, mysqlIns port.RepositoryStore) domain.MutualFundSvr {
	return &mutualFundUsecase{
		mysql:  mysqlIns,
		logger: loggerIns,
	}
}

func (m *mutualFundUsecase) BuyMutualFund(ctx context.Context, userId int, accountId int, inventoryId int, stockId int, quantity float64, amountPerQuantity float64, feeAmount float64) error {
	return nil

}

func (m *mutualFundUsecase) SellMutualFund(ctx context.Context, userId int, accountId int, inventoryId int, stockId int, quantity float64, amountPerQuantity float64, feeAmount float64) error {

	return nil
}
