package stock

import (
	"assetio/internal/domain"
	"assetio/internal/port"
)

type stockUsecase struct {
	logger   port.Logger
	mysql    port.RepositoryStore
	marketer port.Marketer
}

func New(loggerIns port.Logger, mysqlIns port.RepositoryStore, marketerIns port.Marketer) domain.StockSvr {
	return &stockUsecase{
		mysql:    mysqlIns,
		logger:   loggerIns,
		marketer: marketerIns,
	}
}
