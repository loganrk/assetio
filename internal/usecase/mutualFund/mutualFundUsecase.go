package mutualFund

import (
	"assetio/internal/adapters/handler/response"
	"assetio/internal/domain"
	"assetio/internal/port"
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

func (m *mutualFundUsecase) MutualFundBuy(request domain.ClientMutualFundBuyRequest) domain.Response {

	res := response.New()

	return res

}

func (m *mutualFundUsecase) MutualFundAdd(request domain.ClientMutualFundAddRequest) domain.Response {

	res := response.New()

	return res

}

func (m *mutualFundUsecase) MutualFundSell(request domain.ClientMutualFundSellRequest) domain.Response {

	res := response.New()

	return res

}

func (m *mutualFundUsecase) MutualFundSummary(request domain.ClientMutualFundSummaryRequest) domain.Response {

	res := response.New()

	return res

}

func (m *mutualFundUsecase) MutualFundInventory(request domain.ClientMutualFundInventoryRequest) domain.Response {

	res := response.New()

	return res

}
func (m *mutualFundUsecase) MutualFundInventoryLedgers(request domain.ClientMutualFundInventoryLedgersRequest) domain.Response {

	res := response.New()

	return res

}
