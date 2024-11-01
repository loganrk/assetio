package usecases

import (
	"assetio/internal/adapters/handler/response"
	"assetio/internal/constant"
	"assetio/internal/domain"
	"assetio/internal/port"
	"context"
	"net/http"
)

type securityUsecase struct {
	logger port.Logger
	mysql  port.RepositoryStore
}

func New(loggerIns port.Logger, mysqlIns port.RepositoryStore) domain.SecuritySvr {
	return &securityUsecase{
		mysql:  mysqlIns,
		logger: loggerIns,
	}
}

func (s *securityUsecase) SecurityCreate(request domain.ClientSecurityCreateRequest) domain.Response {
	ctx := context.Background()
	res := response.New()

	securityType := s.getType(request.Type)
	if securityType == 0 {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(constant.ERROR_CODE_REQUEST_INVALID, "invalid type")
		return res
	}

	securityExchange := s.getExchange(request.Exchange)
	if securityExchange == 0 {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(constant.ERROR_CODE_REQUEST_INVALID, "invalid exchange")
		return res
	}

	securityData, err := s.mysql.GetSecuriryDataByTypeAndExchangeAndSymbol(ctx, securityType, securityExchange, request.Symbol)

	if err != nil {
		s.logger.Errorw(ctx, "GetSecuriryDataByTypeAndExchangeAndSymbol failed",
			constant.ERROR_TYPE, constant.ERROR_TYPE_DBEXECUTION,
			constant.ERROR_MESSAGE, err.Error(),
			constant.REQUEST, request,
		)

		res.SetStatus(http.StatusInternalServerError)
		res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
		return res
	}

	if securityData.Id != 0 {
		resData := domain.ClientSecurityCreateResponse{
			Message: "security symbol already available",
		}
		res.SetData(resData)
		return res
	}

	_, err = s.mysql.InsertSecurityData(ctx, domain.Securities{
		Type:     securityType,
		Exchange: securityExchange,
		Name:     request.Name,
		Symbol:   request.Symbol,
	})

	if err != nil {
		s.logger.Errorw(ctx, "InsertSecurityData failed",
			constant.ERROR_TYPE, constant.ERROR_TYPE_DBEXECUTION,
			constant.ERROR_MESSAGE, err.Error(),
			constant.REQUEST, request,
		)

		res.SetStatus(http.StatusInternalServerError)
		res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
		return res
	}

	resData := domain.ClientSecurityCreateResponse{
		Message: "security created successfully",
	}

	res.SetData(resData)
	return res

}

func (s *securityUsecase) SecurityAll(request domain.ClientSecurityAllRequest) domain.Response {
	ctx := context.Background()
	res := response.New()

	securityType := s.getType(request.Type)
	if securityType == 0 {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(constant.ERROR_CODE_REQUEST_INVALID, "invalid type")
		return res
	}

	securityExchange := s.getExchange(request.Exchange)
	if securityExchange == 0 {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(constant.ERROR_CODE_REQUEST_INVALID, "invalid exchange")
		return res
	}

	securitiesData, err := s.mysql.GetSecuritiesDataByExchange(ctx, securityType, securityExchange)

	if err != nil {
		s.logger.Errorw(ctx, "GetSecuritiesDataByExchange failed",
			constant.ERROR_TYPE, constant.ERROR_TYPE_DBEXECUTION,
			constant.ERROR_MESSAGE, err.Error(),
			constant.REQUEST, request,
		)

		res.SetStatus(http.StatusInternalServerError)
		res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
		return res
	}

	if len(securitiesData) == 0 {
		res.SetData(nil)
		return res
	}

	var resData []domain.ClientSecurityAllResponse
	for _, securityData := range securitiesData {
		resData = append(resData, domain.ClientSecurityAllResponse{
			Id:       securityData.Id,
			Type:     s.getTypeString(securityData.Type),
			Exchange: s.getExchangeString(securityData.Exchange),
			Symbol:   securityData.Symbol,
			Name:     securityData.Name,
		})
	}

	res.SetData(resData)
	return res
}

func (s *securityUsecase) SecurityGet(request domain.ClientSecurityGetRequest) domain.Response {
	ctx := context.Background()
	res := response.New()

	securityData, err := s.mysql.GetSecuriryDataById(ctx, request.SecurityId)
	if err != nil {
		s.logger.Errorw(ctx, "GetSecuriryDataById failed",
			constant.ERROR_TYPE, constant.ERROR_TYPE_DBEXECUTION,
			constant.ERROR_MESSAGE, err.Error(),
			constant.REQUEST, request,
		)

		res.SetStatus(http.StatusInternalServerError)
		res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
		return res
	}

	if securityData.Id == 0 {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(constant.ERROR_CODE_REQUEST_INVALID, "invalid security id")
		return res
	}

	resData := domain.ClientSecurityGetResponse{
		Id:       securityData.Id,
		Type:     s.getTypeString(securityData.Type),
		Exchange: s.getExchangeString(securityData.Exchange),
		Symbol:   securityData.Symbol,
		Name:     securityData.Name,
	}

	res.SetData(resData)
	return res

}

func (s *securityUsecase) SecuritySearch(request domain.ClientSecuritySearchRequest) domain.Response {
	ctx := context.Background()
	res := response.New()

	securityType := s.getType(request.Type)
	if securityType == 0 {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(constant.ERROR_CODE_REQUEST_INVALID, "invalid type")
		return res
	}

	securityExchange := s.getExchange(request.Exchange)
	if securityExchange == 0 {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(constant.ERROR_CODE_REQUEST_INVALID, "invalid exchange")
		return res
	}

	securitiesData, err := s.mysql.SearchSecuritiesDataByTypeAndExchange(ctx, securityType, securityExchange, request.Search)
	if err != nil {
		s.logger.Errorw(ctx, "SearchSecuritiesDataByTypeAndExchange failed",
			constant.ERROR_TYPE, constant.ERROR_TYPE_DBEXECUTION,
			constant.ERROR_MESSAGE, err.Error(),
			constant.REQUEST, request,
		)

		res.SetStatus(http.StatusInternalServerError)
		res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
		return res

	}

	if len(securitiesData) == 0 {
		res.SetData(nil)
		return res
	}

	var resData []domain.ClientSecuritySearchResponse
	for _, securityData := range securitiesData {
		resData = append(resData, domain.ClientSecuritySearchResponse{
			Id:       securityData.Id,
			Type:     s.getTypeString(securityData.Type),
			Exchange: s.getExchangeString(securityData.Exchange),
			Symbol:   securityData.Symbol,
			Name:     securityData.Name,
		})
	}

	res.SetData(resData)
	return res

}
func (s *securityUsecase) SecurityUpdate(request domain.ClientSecurityUpdateRequest) domain.Response {
	ctx := context.Background()
	res := response.New()

	securityType := s.getType(request.Type)
	if securityType == 0 {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(constant.ERROR_CODE_REQUEST_INVALID, "invalid type")
		return res
	}

	securityExchange := s.getExchange(request.Exchange)
	if securityExchange == 0 {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(constant.ERROR_CODE_REQUEST_INVALID, "invalid exchange")
		return res
	}

	securityData, err := s.mysql.GetSecuriryDataById(ctx, request.SecurityId)
	if err != nil {
		s.logger.Errorw(ctx, "GetSecuriryDataById failed",
			constant.ERROR_TYPE, constant.ERROR_TYPE_DBEXECUTION,
			constant.ERROR_MESSAGE, err.Error(),
			constant.REQUEST, request,
		)

		res.SetStatus(http.StatusInternalServerError)
		res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
		return res
	}

	if securityData.Id == 0 {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(constant.ERROR_CODE_REQUEST_INVALID, "invalid security id")
		return res
	}

	securityData, err = s.mysql.GetSecuriryDataByTypeAndExchangeAndSymbol(ctx, securityType, securityExchange, request.Symbol)
	if err != nil {
		s.logger.Errorw(ctx, "GetSecuriryDataByTypeAndExchangeAndSymbol failed",
			constant.ERROR_TYPE, constant.ERROR_TYPE_DBEXECUTION,
			constant.ERROR_MESSAGE, err.Error(),
			constant.REQUEST, request,
		)

		res.SetStatus(http.StatusInternalServerError)
		res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
		return res
	}

	if securityData.Id != 0 && securityData.Id != request.SecurityId {
		resData := domain.ClientSecurityUpdateResponse{
			Message: "security symbol already available",
		}
		res.SetData(resData)
		return res
	}

	securityData = domain.Securities{
		Type:     securityType,
		Exchange: securityExchange,
		Name:     request.Name,
		Symbol:   request.Symbol,
	}

	err = s.mysql.UpdateSecuriryData(ctx, request.SecurityId, securityData)

	if err != nil {
		s.logger.Errorw(ctx, "UpdateSecuriryData failed",
			constant.ERROR_TYPE, constant.ERROR_TYPE_DBEXECUTION,
			constant.ERROR_MESSAGE, err.Error(),
			constant.REQUEST, request,
		)

		res.SetStatus(http.StatusInternalServerError)
		res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
		return res
	}

	resData := domain.ClientSecurityUpdateResponse{
		Message: "security updated successfully",
	}

	res.SetData(resData)
	return res

}

func (s *securityUsecase) getType(typeData string) int {
	if typeData == constant.SECURITY_TYPE_STOCK_STRING {
		return constant.SECURITY_TYPE_STOCK
	}
	return 0
}

func (s *securityUsecase) getExchange(exchange string) int {
	if exchange == constant.EXCHANGE_TYPE_NSE_STRING {
		return constant.EXCHANGE_TYPE_NSE
	} else if exchange == constant.EXCHANGE_TYPE_BSE_STRING {
		return constant.EXCHANGE_TYPE_BSE
	}
	return 0
}

func (s *securityUsecase) getTypeString(typeData int) string {
	if typeData == constant.SECURITY_TYPE_STOCK {
		return constant.SECURITY_TYPE_STOCK_STRING
	}
	return ""
}

func (s *securityUsecase) getExchangeString(exchange int) string {
	if exchange == constant.EXCHANGE_TYPE_NSE {
		return constant.EXCHANGE_TYPE_NSE_STRING
	} else if exchange == constant.EXCHANGE_TYPE_BSE {
		return constant.EXCHANGE_TYPE_BSE_STRING
	}
	return ""
}
