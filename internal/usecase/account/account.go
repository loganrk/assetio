package account

import (
	"assetio/internal/constant"
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

func (a *accountUsecase) CreateAccount(ctx context.Context, userId int, name string) error {
	_, err := a.mysql.CreateAccount(ctx, domain.Account{
		Name:   name,
		UserId: userId,
		Status: constant.ACCOUNT_STATUS_ACTIVE,
	})

	return err
}

func (a *accountUsecase) GetAccounts(ctx context.Context, userId int) ([]domain.Account, error) {
	accounts, err := a.mysql.GetAccounts(ctx, userId)

	return accounts, err
}

func (a *accountUsecase) GetAccount(ctx context.Context, accountId, userId int) (domain.Account, error) {
	account, err := a.mysql.GetAccount(ctx, accountId, userId)

	return account, err
}

func (a *accountUsecase) UpdateAccount(ctx context.Context, accountId, userId int, name string) error {
	accountData := domain.Account{
		Name: name,
	}
	err := a.mysql.UpdateAccount(ctx, accountId, userId, accountData)

	return err
}

func (a *accountUsecase) AccountActivate(ctx context.Context, accountId, userId int) error {
	accountData := domain.Account{
		Status: constant.ACCOUNT_STATUS_ACTIVE,
	}
	err := a.mysql.UpdateAccount(ctx, accountId, userId, accountData)

	return err
}

func (a *accountUsecase) AccountInactivate(ctx context.Context, accountId, userId int) error {
	accountData := domain.Account{
		Status: constant.ACCOUNT_STATUS_INACTIVE,
	}
	err := a.mysql.UpdateAccount(ctx, accountId, userId, accountData)

	return err
}

func (a *accountUsecase) GetStatusString(status int) string {
	if status == constant.ACCOUNT_STATUS_ACTIVE {
		return "active"
	} else if status == constant.ACCOUNT_STATUS_INACTIVE {
		return "inactive"
	}
	return "unkown"
}
