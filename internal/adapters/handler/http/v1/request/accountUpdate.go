package request

import (
	"assetio/internal/port"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

func NewAccountUpdate(r *http.Request) (port.AccountUpdateClientRequest, error) {
	var a accountUpdate
	if r.Method == http.MethodPost {
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&a)
		if err != nil {
			return &a, err
		}
	} else {

		accountId, _ := strconv.Atoi(r.URL.Query().Get("account_id"))
		a.AccountId = accountId
		a.Name = r.URL.Query().Get("name")
	}

	userid, _ := strconv.Atoi(r.URL.Query().Get("uid"))
	a.UserId = userid
	return &a, nil
}

func (a *accountUpdate) Validate() error {
	if a.AccountId == 0 {
		return errors.New("invalid account id")
	}
	if a.UserId == 0 {
		return errors.New("invalid user id")
	}
	return nil
}

func (a *accountUpdate) GetUserId() int {
	return a.UserId
}

func (a *accountUpdate) GetAccountId() int {
	return a.AccountId
}
func (a *accountUpdate) GetName() string {
	return a.Name
}
