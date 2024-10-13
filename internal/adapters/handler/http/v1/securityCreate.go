package v1

import (
	"assetio/internal/adapters/handler/http/v1/request"
	"assetio/internal/adapters/handler/http/v1/response"
	"assetio/internal/port"
	"context"
	"net/http"
)

func (h *handler) SecurityCreate(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	res := response.New()

	req, err := request.NewSecurityCreate(r)
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

	securityType := h.usecases.Security.GetType(req.GetType())
	securityExchange := h.usecases.Security.GetExchange(req.GetExchange())

	security, err := h.usecases.Security.GetSecuriry(ctx, securityType, securityExchange, req.GetSymbol())
	if err != nil {
		res.SetStatus(http.StatusInternalServerError)
		res.SetError(ERROR_CODE_INTERNAL_SERVER, "internal server error")
		res.Send(w)
		return
	}

	if security.Id != 0 {
		resData := port.SecurityCreateClientResponse{
			Message: "security symbol already available",
		}
		res.SetData(resData)
		res.Send(w)
	}

	err = h.usecases.Security.CreateSecuriry(ctx, securityType, securityExchange, req.GetSymbol(), req.GetName())
	if err != nil {
		res.SetStatus(http.StatusInternalServerError)
		res.SetError(ERROR_CODE_INTERNAL_SERVER, "internal server error")
		res.Send(w)
		return
	}

	resData := port.SecurityCreateClientResponse{
		Message: "security created successfully",
	}
	res.SetData(resData)
	res.Send(w)
}
