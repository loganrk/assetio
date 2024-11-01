package v1

import (
	"assetio/internal/adapters/handler/response"
	"assetio/internal/domain"
	"encoding/json"
	"fmt"
	"strconv"

	"net/http"

	"github.com/gorilla/schema"
)

func (h *handler) AccountCreate(w http.ResponseWriter, r *http.Request) {
	var request domain.ClientAccountCreateRequest
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

	userid, _ := strconv.Atoi(r.URL.Query().Get("uid"))
	request.UserId = userid

	err := h.validator.AccountCreate(request)
	if err != nil {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
		res.Send(w)
		return
	}

	resData := h.usecases.Account.AccountCreate(request)
	resData.Send(w)

}

func (h *handler) AccountAll(w http.ResponseWriter, r *http.Request) {
	var request domain.ClientAccountAllRequest
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

	userid, _ := strconv.Atoi(r.URL.Query().Get("uid"))
	request.UserId = userid

	err := h.validator.AccountAll(request)
	if err != nil {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
		res.Send(w)
		return
	}

	resData := h.usecases.Account.AccountAll(request)
	resData.Send(w)
}

func (h *handler) AccountGet(w http.ResponseWriter, r *http.Request) {

	fmt.Println(w)
	var request domain.ClientAccountGetRequest
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

	userid, _ := strconv.Atoi(r.URL.Query().Get("uid"))
	request.UserId = userid

	err := h.validator.AccountGet(request)
	if err != nil {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
		res.Send(w)
		return
	}

	resData := h.usecases.Account.AccountGet(request)
	resData.Send(w)
}

func (h *handler) AccountActivate(w http.ResponseWriter, r *http.Request) {
	var request domain.ClientAccountActivateRequest
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

	userid, _ := strconv.Atoi(r.URL.Query().Get("uid"))
	request.UserId = userid

	err := h.validator.AccountActivate(request)
	if err != nil {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
		res.Send(w)
		return
	}

	resData := h.usecases.Account.AccountActivate(request)
	resData.Send(w)
}

func (h *handler) AccountInactivate(w http.ResponseWriter, r *http.Request) {

	var request domain.ClientAccountInactivateRequest
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

	userid, _ := strconv.Atoi(r.URL.Query().Get("uid"))
	request.UserId = userid

	err := h.validator.AccountInactivate(request)
	if err != nil {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
		res.Send(w)
		return
	}

	resData := h.usecases.Account.AccountInactivate(request)
	resData.Send(w)
}

func (h *handler) AccountUpdate(w http.ResponseWriter, r *http.Request) {

	var request domain.ClientAccountUpdateRequest
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

	userid, _ := strconv.Atoi(r.URL.Query().Get("uid"))
	request.UserId = userid

	err := h.validator.AccountUpdate(request)
	if err != nil {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
		res.Send(w)
		return
	}

	resData := h.usecases.Account.AccountUpdate(request)
	resData.Send(w)
}
