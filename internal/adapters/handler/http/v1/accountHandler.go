package v1

import (
	"assetio/internal/adapters/handler/response"
	"assetio/internal/domain"
	"encoding/json"
	"strconv"

	"net/http"

	"github.com/gorilla/schema"
)

// AccountCreate handles the creation of a new account. It reads the request body or query parameters,
// decodes the data into a ClientAccountCreateRequest struct, validates it, and then calls the corresponding use case.
func (h *handler) AccountCreate(w http.ResponseWriter, r *http.Request) {
	var request domain.ClientAccountCreateRequest
	res := response.New()

	// Check if the method is POST, and decode the request body
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
		// If not POST, decode query parameters
		var decoder = schema.NewDecoder()
		decoder.IgnoreUnknownKeys(true)
		if err := decoder.Decode(&request, r.URL.Query()); err != nil {
			res.SetStatus(http.StatusBadRequest)
			res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
			res.Send(w)
			return
		}
	}

	// Extract user ID from query parameters and assign it to the request
	userid, _ := strconv.Atoi(r.URL.Query().Get("uid"))
	request.UserId = userid

	// Validate the request data
	err := h.validator.AccountCreate(request)
	if err != nil {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
		res.Send(w)
		return
	}

	// Call the use case to create the account and send the response
	resData := h.usecases.Account.AccountCreate(request)
	resData.Send(w)
}

// AccountAll handles fetching all accounts for a user. It decodes the request data, validates it, and calls the use case.
func (h *handler) AccountAll(w http.ResponseWriter, r *http.Request) {
	var request domain.ClientAccountAllRequest
	res := response.New()

	// Check if the method is POST, and decode the request body
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
		// If not POST, decode query parameters
		var decoder = schema.NewDecoder()
		decoder.IgnoreUnknownKeys(true)
		if err := decoder.Decode(&request, r.URL.Query()); err != nil {
			res.SetStatus(http.StatusBadRequest)
			res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
			res.Send(w)
			return
		}
	}

	// Extract user ID from query parameters and assign it to the request
	userid, _ := strconv.Atoi(r.URL.Query().Get("uid"))
	request.UserId = userid

	// Validate the request data
	err := h.validator.AccountAll(request)
	if err != nil {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
		res.Send(w)
		return
	}

	// Call the use case to fetch all accounts and send the response
	resData := h.usecases.Account.AccountAll(request)
	resData.Send(w)
}

// AccountGet handles fetching a specific account based on the provided account ID.
// It decodes the request, validates the data, and calls the corresponding use case to get the account.
func (h *handler) AccountGet(w http.ResponseWriter, r *http.Request) {
	var request domain.ClientAccountGetRequest
	res := response.New()

	// Check if the method is POST, and decode the request body
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
		// If not POST, decode query parameters
		var decoder = schema.NewDecoder()
		decoder.IgnoreUnknownKeys(true)
		if err := decoder.Decode(&request, r.URL.Query()); err != nil {
			res.SetStatus(http.StatusBadRequest)
			res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
			res.Send(w)
			return
		}
	}

	// Extract user ID from query parameters and assign it to the request
	userid, _ := strconv.Atoi(r.URL.Query().Get("uid"))
	request.UserId = userid

	// Validate the request data
	err := h.validator.AccountGet(request)
	if err != nil {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
		res.Send(w)
		return
	}

	// Call the use case to fetch the specific account and send the response
	resData := h.usecases.Account.AccountGet(request)
	resData.Send(w)
}

// AccountActivate handles activating an account. It decodes the request, validates it, and calls the use case to activate the account.
func (h *handler) AccountActivate(w http.ResponseWriter, r *http.Request) {
	var request domain.ClientAccountActivateRequest
	res := response.New()

	// Check if the method is POST, and decode the request body
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
		// If not POST, decode query parameters
		var decoder = schema.NewDecoder()
		decoder.IgnoreUnknownKeys(true)
		if err := decoder.Decode(&request, r.URL.Query()); err != nil {
			res.SetStatus(http.StatusBadRequest)
			res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
			res.Send(w)
			return
		}
	}

	// Extract user ID from query parameters and assign it to the request
	userid, _ := strconv.Atoi(r.URL.Query().Get("uid"))
	request.UserId = userid

	// Validate the request data
	err := h.validator.AccountActivate(request)
	if err != nil {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
		res.Send(w)
		return
	}

	// Call the use case to activate the account and send the response
	resData := h.usecases.Account.AccountActivate(request)
	resData.Send(w)
}

// AccountInactivate handles inactivating an account. It decodes the request, validates it, and calls the use case to inactivate the account.
func (h *handler) AccountInactivate(w http.ResponseWriter, r *http.Request) {
	var request domain.ClientAccountInactivateRequest
	res := response.New()

	// Check if the method is POST, and decode the request body
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
		// If not POST, decode query parameters
		var decoder = schema.NewDecoder()
		decoder.IgnoreUnknownKeys(true)
		if err := decoder.Decode(&request, r.URL.Query()); err != nil {
			res.SetStatus(http.StatusBadRequest)
			res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
			res.Send(w)
			return
		}
	}

	// Extract user ID from query parameters and assign it to the request
	userid, _ := strconv.Atoi(r.URL.Query().Get("uid"))
	request.UserId = userid

	// Validate the request data
	err := h.validator.AccountInactivate(request)
	if err != nil {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
		res.Send(w)
		return
	}

	// Call the use case to inactivate the account and send the response
	resData := h.usecases.Account.AccountInactivate(request)
	resData.Send(w)
}

// AccountUpdate handles updating an account. It decodes the request, validates it, and calls the use case to update the account.
func (h *handler) AccountUpdate(w http.ResponseWriter, r *http.Request) {
	var request domain.ClientAccountUpdateRequest
	res := response.New()

	// Check if the method is POST, and decode the request body
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
		// If not POST, decode query parameters
		var decoder = schema.NewDecoder()
		decoder.IgnoreUnknownKeys(true)
		if err := decoder.Decode(&request, r.URL.Query()); err != nil {
			res.SetStatus(http.StatusBadRequest)
			res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
			res.Send(w)
			return
		}
	}

	// Extract user ID from query parameters and assign it to the request
	userid, _ := strconv.Atoi(r.URL.Query().Get("uid"))
	request.UserId = userid

	// Validate the request data
	err := h.validator.AccountUpdate(request)
	if err != nil {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
		res.Send(w)
		return
	}

	// Call the use case to update the account and send the response
	resData := h.usecases.Account.AccountUpdate(request)
	resData.Send(w)
}
