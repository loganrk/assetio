package v1

import (
	"assetio/internal/adapters/handler/http/v1/request"
	"assetio/internal/adapters/handler/http/v1/response"
	"assetio/internal/port"
	"context"
	"net/http"
)

func (h *handler) SecurityGet(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	res := response.New()

	req, err := request.NewSecurityGet(r)
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

	security, err := h.usecases.Security.GetSecuriryById(ctx, req.GetSecuriryId())
	if err != nil {
		res.SetStatus(http.StatusInternalServerError)
		res.SetError(ERROR_CODE_INTERNAL_SERVER, "internal server error")
		res.Send(w)
		return
	}

	if security.Id == 0 {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(ERROR_CODE_REQUEST_INVALID, "invalid security id")
		res.Send(w)
		return
	}

	resData := port.SecurityGetClientResponse{
		Id:       security.Id,
		Type:     h.usecases.Security.GetTypeString(security.Type),
		Exchange: h.usecases.Security.GetExchangeString(security.Exchange),
		Symbol:   security.Symbol,
		Name:     security.Name,
	}

	res.SetData(resData)
	res.Send(w)

}
