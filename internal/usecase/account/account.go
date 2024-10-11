package account

import (
	"assetio/internal/domain"
	"assetio/internal/port"
	"context"
)

type accountUsecase struct {
	logger port.Logger
	mysql  port.RepositoryMySQL
}

func New(loggerIns port.Logger, mysqlIns port.RepositoryMySQL) domain.AccountSvr {
	return &accountUsecase{
		mysql:  mysqlIns,
		logger: loggerIns,
	}
}

func (a *accountUsecase) CreateAccount(ctx context.Context, userId int, name string) (int, error) {
	accountId, err := a.mysql.CreateAccount(ctx, domain.Account{
		Name:   name,
		UserId: userId,
	})

	return accountId, err
}
