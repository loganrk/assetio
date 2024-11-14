package v1

import (
	"assetio/internal/adapters/handler/response"
	"assetio/internal/domain"
	"encoding/json"

	"net/http"

	"github.com/gorilla/schema"
)

// SecurityCreate handles the creation of a new security.
func (h *handler) SecurityCreate(w http.ResponseWriter, r *http.Request) {
	var request domain.ClientSecurityCreateRequest
	res := response.New()

	// Decoding the request body or URL query parameters based on the HTTP method
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

	// Validate the request data
	err := h.validator.SecurityCreate(request)
	if err != nil {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
		res.Send(w)
		return
	}

	// Call use case to create security and send the response
	resData := h.usecases.Security.SecurityCreate(request)
	resData.Send(w)
}

// SecurityGet handles retrieving a specific security's details.
func (h *handler) SecurityGet(w http.ResponseWriter, r *http.Request) {
	var request domain.ClientSecurityGetRequest
	res := response.New()

	// Decoding the request body or URL query parameters based on the HTTP method
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

	// Validate the request data
	err := h.validator.SecurityGet(request)
	if err != nil {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
		res.Send(w)
		return
	}

	// Call use case to retrieve security details and send the response
	resData := h.usecases.Security.SecurityGet(request)
	resData.Send(w)
}

// SecurityAll handles retrieving a list of all securities.
func (h *handler) SecurityAll(w http.ResponseWriter, r *http.Request) {
	var request domain.ClientSecurityAllRequest
	res := response.New()

	// Decoding the request body or URL query parameters based on the HTTP method
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

	// Validate the request data
	err := h.validator.SecurityAll(request)
	if err != nil {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
		res.Send(w)
		return
	}

	// Call use case to retrieve all securities and send the response
	resData := h.usecases.Security.SecurityAll(request)
	resData.Send(w)
}

// SecuritySearch handles searching for securities based on specified criteria.
func (h *handler) SecuritySearch(w http.ResponseWriter, r *http.Request) {
	var request domain.ClientSecuritySearchRequest
	res := response.New()

	// Decoding the request body or URL query parameters based on the HTTP method
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

	// Validate the request data
	err := h.validator.SecuritySearch(request)
	if err != nil {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
		res.Send(w)
		return
	}

	// Call use case to perform security search and send the response
	resData := h.usecases.Security.SecuritySearch(request)
	resData.Send(w)
}

// SecurityUpdate handles updating an existing security.
func (h *handler) SecurityUpdate(w http.ResponseWriter, r *http.Request) {
	var request domain.ClientSecurityUpdateRequest
	res := response.New()

	// Decoding the request body or URL query parameters based on the HTTP method
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

	// Validate the request data
	err := h.validator.SecurityUpdate(request)
	if err != nil {
		res.SetStatus(http.StatusBadRequest)
		res.SetError(ERROR_CODE_REQUEST_INVALID, err.Error())
		res.Send(w)
		return
	}

	// Call use case to update the security and send the response
	resData := h.usecases.Security.SecurityUpdate(request)
	resData.Send(w)
}
