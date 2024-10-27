package v1

import (
	"assetio/internal/adapters/handler/response"
	"assetio/internal/domain"
	"encoding/json"

	"net/http"

	"github.com/gorilla/schema"
)

func (h *handler) SecurityCreate(w http.ResponseWriter, r *http.Request) {
	var request domain.ClientSecurityCreateRequest
	res := response.New()
	if r.Method == http.MethodPost {
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&request)
		if err != nil {
			res.SetStatus(http.StatusBadRequest)
			res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
			res.Send(w)
			return
		}
	} else {
		var decoder = schema.NewDecoder()
		decoder.IgnoreUnknownKeys(true)
		if err := decoder.Decode(&request, r.URL.Query()); err != nil {
			res.SetStatus(http.StatusBadRequest)
			res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
			res.Send(w)
			return
		}
	}

	err := h.validator.SecurityCreate(request)
	if err != nil {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
		res.Send(w)
		return
	}

	resData := h.usecases.Security.SecurityCreate(request)
	resData.Send(w)

}

func (h *handler) SecurityGet(w http.ResponseWriter, r *http.Request) {
	var request domain.ClientSecurityGetRequest
	res := response.New()

	if r.Method == http.MethodPost {
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&request)
		if err != nil {
			res.SetStatus(http.StatusBadRequest)
			res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
			res.Send(w)
			return
		}
	} else {
		var decoder = schema.NewDecoder()
		decoder.IgnoreUnknownKeys(true)
		if err := decoder.Decode(&request, r.URL.Query()); err != nil {
			res.SetStatus(http.StatusBadRequest)
			res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
			res.Send(w)
			return
		}
	}

	err := h.validator.SecurityGet(request)
	if err != nil {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
		res.Send(w)
		return
	}

	resData := h.usecases.Security.SecurityGet(request)
	resData.Send(w)

}

func (h *handler) SecurityAll(w http.ResponseWriter, r *http.Request) {
	var request domain.ClientSecurityAllRequest
	res := response.New()

	if r.Method == http.MethodPost {
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&request)
		if err != nil {
			res.SetStatus(http.StatusBadRequest)
			res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
			res.Send(w)
			return
		}
	} else {
		var decoder = schema.NewDecoder()
		decoder.IgnoreUnknownKeys(true)
		if err := decoder.Decode(&request, r.URL.Query()); err != nil {
			res.SetStatus(http.StatusBadRequest)
			res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
			res.Send(w)
			return
		}
	}

	err := h.validator.SecurityAll(request)
	if err != nil {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
		res.Send(w)
		return
	}

	resData := h.usecases.Security.SecurityAll(request)
	resData.Send(w)

}

func (h *handler) SecuritySearch(w http.ResponseWriter, r *http.Request) {
	var request domain.ClientSecuritySearchRequest
	res := response.New()

	if r.Method == http.MethodPost {
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&request)
		if err != nil {
			res.SetStatus(http.StatusBadRequest)
			res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
			res.Send(w)
			return
		}
	} else {
		var decoder = schema.NewDecoder()
		decoder.IgnoreUnknownKeys(true)
		if err := decoder.Decode(&request, r.URL.Query()); err != nil {
			res.SetStatus(http.StatusBadRequest)
			res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
			res.Send(w)
			return
		}
	}

	err := h.validator.SecuritySearch(request)
	if err != nil {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
		res.Send(w)
		return
	}

	resData := h.usecases.Security.SecuritySearch(request)
	resData.Send(w)
}

func (h *handler) SecurityUpdate(w http.ResponseWriter, r *http.Request) {
	var request domain.ClientSecurityUpdateRequest
	res := response.New()

	if r.Method == http.MethodPost {
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&request)
		if err != nil {
			res.SetStatus(http.StatusBadRequest)
			res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
			res.Send(w)
			return
		}
	} else {
		var decoder = schema.NewDecoder()
		decoder.IgnoreUnknownKeys(true)
		if err := decoder.Decode(&request, r.URL.Query()); err != nil {
			res.SetStatus(http.StatusBadRequest)
			res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
			res.Send(w)
			return
		}
	}

	err := h.validator.SecurityUpdate(request)
	if err != nil {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
		res.Send(w)
		return
	}

	resData := h.usecases.Security.SecurityUpdate(request)
	resData.Send(w)
}
