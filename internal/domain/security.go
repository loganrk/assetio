package domain

type ClientSecurityCreateRequest struct {
	Type     string `json:"type" schema:"type"`
	Exchange string `json:"exchange" schema:"exchange"`
	Symbol   string `json:"symbol" schema:"symbol"`
	Name     string `json:"name" schema:"name"`
}

type ClientSecurityCreateResponse struct {
	Message string `json:"message" schema:"message"`
}

type ClientSecurityUpdateRequest struct {
	SecurityId int    `json:"security_id" schema:"security_id"`
	Type       string `json:"type" schema:"type"`
	Exchange   string `json:"exchange" schema:"exchange"`
	Symbol     string `json:"symbol" schema:"symbol"`
	Name       string `json:"name" schema:"name"`
}

type ClientSecurityUpdateResponse struct {
	Message string `json:"message" schema:"message"`
}

type ClientSecurityAllRequest struct {
	Type     string `json:"type" schema:"type"`
	Exchange string `json:"exchange" schema:"exchange"`
}

type ClientSecurityAllResponse struct {
	Id       int    `json:"id" schema:"id"`
	Type     string `json:"type" schema:"type"`
	Exchange string `json:"exchange" schema:"exchange"`
	Symbol   string `json:"symbol" schema:"symbol"`
	Name     string `json:"name" schema:"name"`
}

type ClientSecurityGetRequest struct {
	SecurityId int `json:"security_id" schema:"security_id"`
}

type ClientSecurityGetResponse struct {
	Id       int    `json:"id" schema:"id"`
	Type     string `json:"type" schema:"type"`
	Exchange string `json:"exchange" schema:"exchange"`
	Symbol   string `json:"symbol" schema:"symbol"`
	Name     string `json:"name" schema:"name"`
}

type ClientSecuritySearchRequest struct {
	Type     string `json:"type" schema:"type"`
	Exchange string `json:"exchange" schema:"exchange"`
	Search   string `json:"search" schema:"search"`
}

type ClientSecuritySearchResponse struct {
	Id       int    `json:"id" schema:"id"`
	Type     string `json:"type" schema:"type"`
	Exchange string `json:"exchange" schema:"exchange"`
	Symbol   string `json:"symbol" schema:"symbol"`
	Name     string `json:"name" schema:"name"`
}
