package port

import "net/http"

type Response interface {
	SetError(errCode string, errMsg string)
	SetStatus(status int)
	SetData(data any)
	Send(w http.ResponseWriter)
}
type AccountCreateClientResponse struct {
	Message string `json:"message"`
}

type AccountGetClientResponse struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

type AccountAllClientResponse struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

type AccountUpdateClientResponse struct {
	Message string `json:"message"`
}

type AccountActivateClientResponse struct {
	Message string `json:"message"`
}
type AccountInactivateClientResponse struct {
	Message string `json:"message"`
}

type SecurityCreateClientResponse struct {
	Message string `json:"message"`
}

type SecurityGetClientResponse struct {
	Type     string `json:"type"`
	Exchange string `json:"exchange"`
	Symbol   string `json:"symbol"`
	Name     string `json:"name"`
}

type SecurityAllClientResponse struct {
	Type     string `json:"type"`
	Exchange string `json:"exchange"`
	Symbol   string `json:"symbol"`
	Name     string `json:"name"`
}
