package v1

import (
	"assetio/internal/adapters/handler/http/v1/request"
	"assetio/internal/adapters/handler/http/v1/response"
	"assetio/internal/constant"
	"assetio/internal/port"
	"context"
	"net/http"
)

func (h *handler) AccountActivate(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	res := response.New()

	req, err := request.NewAccountActivate(r)
	if err != nil {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(ERROR_CODE_REQUEST_INVALID, "invalid request parameters")
		res.Send(w)
		return
	}

	err = req.Validate()
	if err != nil {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
		res.Send(w)
		return
	}

	account, err := h.usecases.Account.GetAccount(ctx, req.GetAccountId(), req.GetUserId())
	if err != nil {
		res.SetStatus(http.StatusInternalServerError)
		res.SetError(ERROR_CODE_REQUEST_INVALID, "internal server error")
		res.Send(w)
		return
	}

	if account.Id == 0 {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(ERROR_CODE_REQUEST_INVALID, "incorrect account id")
		res.Send(w)
		return
	}

	if account.Status == constant.ACCOUNT_STATUS_ACTIVE {
		resData := port.AccountActivateClientResponse{
			Message: "account already active",
		}
		res.SetData(resData)
		res.Send(w)
	}

	err = h.usecases.Account.AccountActivate(ctx, req.GetAccountId(), req.GetUserId())
	if err != nil {
		res.SetStatus(http.StatusInternalServerError)
		res.SetError(ERROR_CODE_INTERNAL_SERVER, "internal server error")
		res.Send(w)
		return
	}

	resData := port.AccountActivateClientResponse{
		Message: "account activated successfully",
	}
	res.SetData(resData)
	res.Send(w)
}
