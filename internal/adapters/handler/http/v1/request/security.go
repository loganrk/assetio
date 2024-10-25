package request

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

func NewSecurityAll(r *http.Request) (*securityAll, error) {
	var s securityAll
	if r.Method == http.MethodPost {
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&s)
		if err != nil {
			return &s, err
		}
	} else {
		s.Type = r.URL.Query().Get("type")
		s.Exchange = r.URL.Query().Get("exchange")
	}

	return &s, nil
}

func NewSecuritySearch(r *http.Request) (*securitySearch, error) {
	var s securitySearch
	if r.Method == http.MethodPost {
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&s)
		if err != nil {
			return &s, err
		}
	} else {
		s.Type = r.URL.Query().Get("type")
		s.Exchange = r.URL.Query().Get("exchange")
		s.Search = r.URL.Query().Get("search")

	}

	return &s, nil
}

func NewSecurityUpdate(r *http.Request) (*securityUpdate, error) {
	var s securityUpdate
	if r.Method == http.MethodPost {
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&s)
		if err != nil {
			return &s, err
		}
	} else {
		securityId, _ := strconv.Atoi(r.URL.Query().Get("security_id"))
		s.SecurityId = securityId

		s.Name = r.URL.Query().Get("name")
		s.Type = r.URL.Query().Get("type")
		s.Exchange = r.URL.Query().Get("exchange")
		s.Symbol = r.URL.Query().Get("symbol")
	}

	return &s, nil
}

func NewSecurityCreate(r *http.Request) (*securityCreate, error) {
	var s securityCreate
	if r.Method == http.MethodPost {
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&s)
		if err != nil {
			return &s, err
		}
	} else {
		s.Name = r.URL.Query().Get("name")
		s.Type = r.URL.Query().Get("type")
		s.Exchange = r.URL.Query().Get("exchange")
		s.Symbol = r.URL.Query().Get("symbol")
	}

	return &s, nil
}

func NewSecurityGet(r *http.Request) (*securityGet, error) {
	var s securityGet
	if r.Method == http.MethodPost {
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&s)
		if err != nil {
			return &s, err
		}
	} else {
		securityId, _ := strconv.Atoi(r.URL.Query().Get("security_id"))
		s.SecurityId = securityId
	}

	return &s, nil
}

func (s *securityGet) Validate() error {

	if s.SecurityId == 0 {
		return errors.New("invalid security id")
	}

	return nil
}

func (s *securityCreate) Validate() error {

	if s.Name == "" {
		return errors.New("invalid name")
	}

	if s.Symbol == "" {
		return errors.New("invalid symbol")
	}

	if s.Exchange == "" {
		return errors.New("invalid exchange")
	}
	if s.Type == "" {
		return errors.New("invalid type")
	}

	return nil
}

func (s *securityUpdate) Validate() error {

	if s.SecurityId == 0 {
		return errors.New("invalid security id")
	}
	if s.Name == "" {
		return errors.New("invalid name")
	}

	if s.Symbol == "" {
		return errors.New("invalid symbol")
	}

	if s.Exchange == "" {
		return errors.New("invalid exchange")
	}
	if s.Type == "" {
		return errors.New("invalid type")
	}

	return nil
}
func (s *securitySearch) Validate() error {

	if s.Exchange == "" {
		return errors.New("invalid exchange")
	}
	if s.Type == "" {
		return errors.New("invalid type")
	}

	if s.Search == "" {
		return errors.New("invalid search keyword")
	}

	if len(s.Search) <= 1 {
		return errors.New("search keyword should be minimum 2 letters")
	}

	return nil
}

func (s *securityAll) Validate() error {

	if s.Exchange == "" {
		return errors.New("invalid exchange")
	}
	if s.Type == "" {
		return errors.New("invalid type")
	}

	return nil
}
