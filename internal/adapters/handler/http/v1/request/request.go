package request

type accountCreate struct {
	Name   string `json:"name"`
	UserId int    `json:"uid"`
}

type accountAll struct {
	UserId int `json:"uid"`
}

type accountGet struct {
	UserId    int `json:"uid"`
	AccountId int `json:"account_id"`
}

type accountUpdate struct {
	Name      string `json:"name"`
	UserId    int    `json:"uid"`
	AccountId int    `json:"account_id"`
}
type accountActivate struct {
	UserId    int `json:"uid"`
	AccountId int `json:"account_id"`
}

type accountInactivate struct {
	UserId    int `json:"uid"`
	AccountId int `json:"account_id"`
}

type securityCreate struct {
	Type     string `json:"type"`
	Exchange string `json:"exchange"`
	Symbol   string `json:"symbol"`
	Name     string `json:"name"`
}

type securityUpdate struct {
	SecurityId int    `json:"security_id"`
	Type       string `json:"type"`
	Exchange   string `json:"exchange"`
	Symbol     string `json:"symbol"`
	Name       string `json:"name"`
}

type securityGet struct {
	SecurityId int `json:"security_id"`
}

type securityAll struct {
	Type     string `json:"type"`
	Exchange string `json:"exchange"`
}
