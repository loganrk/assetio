package request

import (
	"assetio/internal/port"
	"encoding/json"
	"errors"
	"net/http"
)

func NewSecurityCreate(r *http.Request) (port.SecurityCreateClientRequest, error) {
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

func (s *securityCreate) GetName() string {
	return s.Name
}
func (s *securityCreate) GetType() string {
	return s.Type
}

func (s *securityCreate) GetSymbol() string {
	return s.Symbol
}

func (s *securityCreate) GetExchange() string {
	return s.Exchange
}
