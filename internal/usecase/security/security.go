package usecases

import (
	"assetio/internal/constant"
	"assetio/internal/domain"
	"assetio/internal/port"
	"context"
)

type securityUsecase struct {
	logger port.Logger
	mysql  port.RepositoryMySQL
}

func New(loggerIns port.Logger, mysqlIns port.RepositoryMySQL) domain.SecuritySvr {
	return &securityUsecase{
		mysql:  mysqlIns,
		logger: loggerIns,
	}
}

func (s *securityUsecase) GetType(typeData string) int {
	if typeData == constant.SECURITY_TYPE_STOCK_STRING {
		return constant.SECURITY_TYPE_STOCK
	}
	return 0
}

func (s *securityUsecase) GetExchange(exchange string) int {
	if exchange == constant.EXCHANGE_TYPE_NSE_STRING {
		return constant.EXCHANGE_TYPE_NSE
	} else if exchange == constant.EXCHANGE_TYPE_BSE_STRING {
		return constant.EXCHANGE_TYPE_BSE
	}
	return 0
}

func (s *securityUsecase) GetTypeString(typeData int) string {
	if typeData == constant.SECURITY_TYPE_STOCK {
		return constant.SECURITY_TYPE_STOCK_STRING
	}
	return ""
}

func (s *securityUsecase) GetExchangeString(exchange int) string {
	if exchange == constant.EXCHANGE_TYPE_NSE {
		return constant.EXCHANGE_TYPE_NSE_STRING
	} else if exchange == constant.EXCHANGE_TYPE_BSE {
		return constant.EXCHANGE_TYPE_BSE_STRING
	}
	return ""
}

func (s *securityUsecase) CreateSecuriry(ctx context.Context, types, exchange int, symbol, name string) error {

	_, err := s.mysql.CreateSecuriry(ctx, domain.Security{
		Type:     types,
		Exchange: exchange,
		Name:     name,
		Symbol:   symbol,
	})
	return err
}

func (s *securityUsecase) GetSecuriry(ctx context.Context, types, exchange int, symbol string) (domain.Security, error) {
	securityData, err := s.mysql.GetSecuriry(ctx, types, exchange, symbol)

	return securityData, err
}

func (s *securityUsecase) GetSecuriryById(ctx context.Context, securityId int) (domain.Security, error) {
	securityData, err := s.mysql.GetSecuriryById(ctx, securityId)

	return securityData, err
}

func (s *securityUsecase) UpdateSecuriry(ctx context.Context, securityId, types, exchange int, symbol string, name string) error {
	securityData := domain.Security{
		Type:     types,
		Exchange: exchange,
		Name:     name,
		Symbol:   symbol,
	}
	err := s.mysql.UpdateSecuriry(ctx, securityId, securityData)

	return err
}

func (s *securityUsecase) GetSecurities(ctx context.Context, types, exchange int) ([]domain.Security, error) {
	securityData, err := s.mysql.GetSecurities(ctx, types, exchange)

	return securityData, err
}

func (s *securityUsecase) SearchSecurities(ctx context.Context, types, exchange int, search string) ([]domain.Security, error) {
	securityData, err := s.mysql.SearchSecurities(ctx, types, exchange, search)

	return securityData, err
}
