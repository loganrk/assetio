package domain

type ClientAccountCreateRequest struct {
	Name   string `json:"name" schema:"name"`
	UserId int    `json:"uid" schema:"uid"`
}
type ClientAccountCreateResponse struct {
	Message string `json:"message" schema:"message"`
}

type ClientAccountAllRequest struct {
	UserId int `json:"uid" schema:"uid"`
}

type ClientAccountAllResponse struct {
	Id     int    `json:"id" schema:"id"`
	Name   string `json:"name" schema:"name"`
	Status string `json:"status" schema:"status"`
}

type ClientAccountGetRequest struct {
	UserId    int `json:"uid" schema:"uid"`
	AccountId int `json:"account_id" schema:"account_id"`
}
type ClientAccountGetResponse struct {
	Id     int    `json:"id" schema:"id"`
	Name   string `json:"name" schema:"name"`
	Status string `json:"status" schema:"status"`
}

type ClientAccountUpdateRequest struct {
	Name      string `json:"name" schema:"name"`
	UserId    int    `json:"uid" schema:"uid"`
	AccountId int    `json:"account_id" schema:"account_id"`
}
type ClientAccountUpdateResponse struct {
	Message string `json:"message" schema:"uid"`
}

type ClientAccountActivateRequest struct {
	UserId    int `json:"uid" schema:"uid"`
	AccountId int `json:"account_id" schema:"account_id"`
}
type ClientAccountActivateResponse struct {
	Message string `json:"message" schema:"message"`
}

type ClientAccountInactivateRequest struct {
	UserId    int `json:"uid" schema:"uid"`
	AccountId int `json:"account_id" schema:"account_id"`
}
type ClientAccountInActivateResponse struct {
	Message string `json:"message" schema:"message"`
}
