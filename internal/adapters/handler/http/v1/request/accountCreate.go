package request

import (
	"assetio/internal/port"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

func NewAccountCreate(r *http.Request) (port.AccountCreateClientRequest, error) {
	var a accountCreate
	if r.Method == http.MethodPost {
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&a)
		if err != nil {
			return &a, err
		}
	} else {
		a.Name = r.URL.Query().Get("name")
	}

	userid, _ := strconv.Atoi(r.URL.Query().Get("uid"))
	a.UserId = userid

	return &a, nil
}

func (a *accountCreate) Validate() error {
	if a.Name == "" {
		return errors.New("invalid name")
	}

	if a.UserId == 0 {
		return errors.New("invalid user id")
	}

	return nil
}

func (a *accountCreate) GetUserId() int {
	return a.UserId
}

func (a *accountCreate) GetName() string {
	return a.Name
}
