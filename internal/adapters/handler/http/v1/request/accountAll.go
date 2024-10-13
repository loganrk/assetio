package request

import (
	"assetio/internal/port"
	"errors"
	"net/http"
	"strconv"
)

func NewAccountAll(r *http.Request) (port.AccountAllClientRequest, error) {
	var a accountAll
	userid, _ := strconv.Atoi(r.URL.Query().Get("uid"))
	a.UserId = userid

	return &a, nil
}

func (a *accountAll) Validate() error {

	if a.UserId == 0 {
		return errors.New("invalid user id")
	}

	return nil
}

func (a *accountAll) GetUserId() int {
	return a.UserId
}
