package request

import (
	"assetio/internal/port"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

func NewSecurityUpdate(r *http.Request) (port.SecurityUpdateClientRequest, error) {
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

func (s *securityUpdate) GetSecuriryId() int {
	return s.SecurityId
}

func (s *securityUpdate) GetName() string {
	return s.Name
}
func (s *securityUpdate) GetType() string {
	return s.Type
}

func (s *securityUpdate) GetSymbol() string {
	return s.Symbol
}

func (s *securityUpdate) GetExchange() string {
	return s.Exchange
}
