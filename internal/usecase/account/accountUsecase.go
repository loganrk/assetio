package account

import (
	"assetio/internal/adapters/handler/response"
	"assetio/internal/constant"
	"assetio/internal/domain"
	"assetio/internal/port"
	"context"
	"net/http"
)

type accountUsecase struct {
	logger port.Logger
	mysql  port.RepositoryStore
}

func New(loggerIns port.Logger, mysqlIns port.RepositoryStore) domain.AccountSvr {
	return &accountUsecase{
		mysql:  mysqlIns,
		logger: loggerIns,
	}
}
func (a *accountUsecase) AccountCreate(request domain.ClientAccountCreateRequest) domain.Response {
	ctx := context.Background()
	res := response.New()

	_, err := a.mysql.InsertAccountData(ctx, domain.Accounts{
		Name:   request.Name,
		UserId: request.UserId,
		Status: constant.ACCOUNT_STATUS_ACTIVE,
	})

	if err != nil {
		a.logger.Errorw(ctx, "InsertAccountData failed",
			constant.ERROR_TYPE, constant.ERROR_TYPE_DBEXECUTION,
			constant.ERROR_MESSAGE, err.Error(),
			constant.REQUEST, request,
		)
		res.SetStatus(http.StatusInternalServerError)
		res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
		return res
	}

	resData := domain.ClientAccountCreateResponse{
		Message: "account created successfully",
	}
	res.SetData(resData)
	return res
}

func (a *accountUsecase) AccountAll(request domain.ClientAccountAllRequest) domain.Response {

	ctx := context.Background()
	res := response.New()

	accounts, err := a.mysql.GetAccountsData(ctx, request.UserId)

	if err != nil {
		a.logger.Errorw(ctx, "GetAccountsData failed",
			constant.ERROR_TYPE, constant.ERROR_TYPE_DBEXECUTION,
			constant.ERROR_MESSAGE, err.Error(),
			constant.REQUEST, request,
		)

		res.SetStatus(http.StatusInternalServerError)
		res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
		return res
	}

	if len(accounts) == 0 {
		res.SetData(nil)
		return res
	}

	var resData []domain.ClientAccountAllResponse

	for _, account := range accounts {
		resData = append(resData, domain.ClientAccountAllResponse{
			Id:     account.Id,
			Name:   account.Name,
			Status: a.getStatusString(account.Status),
		})
	}

	res.SetData(resData)
	return res

}

func (a *accountUsecase) AccountGet(request domain.ClientAccountGetRequest) domain.Response {
	ctx := context.Background()
	res := response.New()

	account, err := a.mysql.GetAccountDataByIdAndUserId(ctx, request.AccountId, request.UserId)

	if err != nil {
		a.logger.Errorw(ctx, "GetAccountDataByIdAndUserId failed",
			constant.ERROR_TYPE, constant.ERROR_TYPE_DBEXECUTION,
			constant.ERROR_MESSAGE, err.Error(),
			constant.REQUEST, request,
		)

		res.SetStatus(http.StatusInternalServerError)
		res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
		return res
	}

	if account.Id == 0 {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(constant.ERROR_CODE_REQUEST_INVALID, "incorrect account id")
		return res
	}

	resData := domain.ClientAccountGetResponse{
		Id:     account.Id,
		Name:   account.Name,
		Status: a.getStatusString(account.Status),
	}

	res.SetData(resData)
	return res

}

func (a *accountUsecase) AccountActivate(request domain.ClientAccountActivateRequest) domain.Response {
	ctx := context.Background()
	res := response.New()

	account, err := a.mysql.GetAccountDataByIdAndUserId(ctx, request.AccountId, request.UserId)

	if err != nil {
		a.logger.Errorw(ctx, "GetAccountDataByIdAndUserId failed",
			constant.ERROR_TYPE, constant.ERROR_TYPE_DBEXECUTION,
			constant.ERROR_MESSAGE, err.Error(),
			constant.REQUEST, request,
		)
		res.SetStatus(http.StatusInternalServerError)
		res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
		return res
	}

	if account.Id == 0 {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(constant.ERROR_CODE_REQUEST_INVALID, "incorrect account id")
		return res
	}

	if account.Status == constant.ACCOUNT_STATUS_ACTIVE {
		resData := domain.ClientAccountActivateResponse{
			Message: "account already active",
		}
		res.SetData(resData)
		return res
	}

	accountData := domain.Accounts{
		Status: constant.ACCOUNT_STATUS_ACTIVE,
	}

	err = a.mysql.UpdateAccountData(ctx, request.AccountId, request.UserId, accountData)
	if err != nil {
		a.logger.Errorw(ctx, "UpdateAccountData failed",
			constant.ERROR_TYPE, constant.ERROR_TYPE_DBEXECUTION,
			constant.ERROR_MESSAGE, err.Error(),
			constant.REQUEST, request,
		)

		res.SetStatus(http.StatusInternalServerError)
		res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
		return res
	}

	resData := domain.ClientAccountActivateResponse{
		Message: "account activated successfully",
	}
	res.SetData(resData)
	return res

}

func (a *accountUsecase) AccountInactivate(request domain.ClientAccountInactivateRequest) domain.Response {
	ctx := context.Background()
	res := response.New()

	account, err := a.mysql.GetAccountDataByIdAndUserId(ctx, request.AccountId, request.UserId)

	if err != nil {
		a.logger.Errorw(ctx, "GetAccountDataByIdAndUserId failed",
			constant.ERROR_TYPE, constant.ERROR_TYPE_DBEXECUTION,
			constant.ERROR_MESSAGE, err.Error(),
			constant.REQUEST, request,
		)

		res.SetStatus(http.StatusInternalServerError)
		res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
		return res
	}

	if account.Id == 0 {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(constant.ERROR_CODE_REQUEST_INVALID, "incorrect account id")
		return res
	}

	if account.Status == constant.ACCOUNT_STATUS_INACTIVE {
		resData := domain.ClientAccountInActivateResponse{
			Message: "account already inactive",
		}
		res.SetData(resData)
		return res
	}

	accountData := domain.Accounts{
		Status: constant.ACCOUNT_STATUS_INACTIVE,
	}

	err = a.mysql.UpdateAccountData(ctx, request.AccountId, request.UserId, accountData)
	if err != nil {
		a.logger.Errorw(ctx, "GetAccountDataByIdAndUserId failed",
			constant.ERROR_TYPE, constant.ERROR_TYPE_DBEXECUTION,
			constant.ERROR_MESSAGE, err.Error(),
			constant.REQUEST, request,
		)
		res.SetStatus(http.StatusInternalServerError)
		res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
		return res
	}

	resData := domain.ClientAccountInActivateResponse{
		Message: "account inactivated successfully",
	}
	res.SetData(resData)
	return res

}

func (a *accountUsecase) AccountUpdate(request domain.ClientAccountUpdateRequest) domain.Response {
	ctx := context.Background()
	res := response.New()

	accountData := domain.Accounts{
		Name: request.Name,
	}
	err := a.mysql.UpdateAccountData(ctx, request.AccountId, request.UserId, accountData)
	if err != nil {
		a.logger.Errorw(ctx, "UpdateAccountData failed",
			constant.ERROR_TYPE, constant.ERROR_TYPE_DBEXECUTION,
			constant.ERROR_MESSAGE, err.Error(),
			constant.REQUEST, request,
		)
		res.SetStatus(http.StatusInternalServerError)
		res.SetError(constant.ERROR_CODE_INTERNAL_SERVER, "internal server error")
		return res
	}

	resData := domain.ClientAccountUpdateResponse{
		Message: "account updated successfully",
	}
	res.SetData(resData)
	return res

}

func (a *accountUsecase) getStatusString(status int) string {
	if status == constant.ACCOUNT_STATUS_ACTIVE {
		return "active"
	} else if status == constant.ACCOUNT_STATUS_INACTIVE {
		return "inactive"
	}
	return "unkown"
}
