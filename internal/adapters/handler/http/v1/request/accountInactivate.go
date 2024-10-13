package request

import (
	"assetio/internal/port"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

func NewAccountInactivate(r *http.Request) (port.AccountInactivateClientRequest, error) {
	var a accountInactivate

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

func (a *accountInactivate) Validate() error {

	if a.AccountId == 0 {
		return errors.New("invalid account id ")
	}

	if a.UserId == 0 {
		return errors.New("invalid user id")
	}

	return nil
}

func (a *accountInactivate) GetUserId() int {
	return a.UserId
}

func (a *accountInactivate) GetAccountId() int {
	return a.AccountId
}
