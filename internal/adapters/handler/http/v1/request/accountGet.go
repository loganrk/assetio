package request

import (
	"assetio/internal/port"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

func NewAccountGet(r *http.Request) (port.AccountGetClientRequest, error) {
	var a accountGet

	if r.Method == http.MethodPost {
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&a)
		if err != nil {
			return &a, err
		}
	} else {
		accountId, _ := strconv.Atoi(r.URL.Query().Get("account_id"))
		a.AccountId = accountId
	}

	userid, _ := strconv.Atoi(r.URL.Query().Get("uid"))
	a.UserId = userid

	return &a, nil
}

func (a *accountGet) Validate() error {

	if a.AccountId == 0 {
		return errors.New("invalid account id ")
	}

	if a.UserId == 0 {
		return errors.New("invalid user id")
	}

	return nil
}

func (a *accountGet) GetUserId() int {
	return a.UserId
}

func (a *accountGet) GetAccountId() int {
	return a.AccountId
}
