package request

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

func NewAccountActivate(r *http.Request) (*accountActivate, error) {
	var a accountActivate

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

func NewAccountAll(r *http.Request) (*accountAll, error) {
	var a accountAll
	userid, _ := strconv.Atoi(r.URL.Query().Get("uid"))
	a.UserId = userid

	return &a, nil
}

func NewAccountCreate(r *http.Request) (*accountCreate, error) {
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

func NewAccountGet(r *http.Request) (*accountGet, error) {
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

func NewAccountInactivate(r *http.Request) (*accountInactivate, error) {
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

func NewAccountUpdate(r *http.Request) (*accountUpdate, error) {
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

func (a *accountInactivate) Validate() error {

	if a.AccountId == 0 {
		return errors.New("invalid account id ")
	}

	if a.UserId == 0 {
		return errors.New("invalid user id")
	}

	return nil
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

func (a *accountCreate) Validate() error {
	if a.Name == "" {
		return errors.New("invalid name")
	}

	if a.UserId == 0 {
		return errors.New("invalid user id")
	}

	return nil
}

func (a *accountAll) Validate() error {

	if a.UserId == 0 {
		return errors.New("invalid user id")
	}

	return nil
}

func (a *accountActivate) Validate() error {

	if a.AccountId == 0 {
		return errors.New("invalid account id ")
	}

	if a.UserId == 0 {
		return errors.New("invalid user id")
	}

	return nil
}
