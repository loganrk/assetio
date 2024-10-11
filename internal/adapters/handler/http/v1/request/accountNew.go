package request

import (
	"assetio/internal/port"
	"encoding/json"
	"errors"
	"net/http"
)

func NewAccount(r *http.Request) (port.AccountNewClientRequest, error) {
	var a accountNew
	if r.Method == http.MethodPost {
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(a)
		if err != nil {
			return &a, err
		}
	} else {
		a.Name = r.URL.Query().Get("name")
	}

	return &a, nil
}

func (a *accountNew) Validate() error {
	if a.Name == "" {
		return errors.New("invalid name")
	}

	if a.UserId == 0 {
		return errors.New("invalid user")
	}

	return nil
}

func (a *accountNew) GetUserId() int {
	return a.UserId
}

func (a *accountNew) GetName() string {
	return a.Name
}
