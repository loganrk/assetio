package request

import (
	"assetio/internal/port"
	"encoding/json"
	"errors"
	"net/http"
)

func NewSecurityAll(r *http.Request) (port.SecurityAllClientRequest, error) {
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

func (s *securityAll) Validate() error {

	if s.Exchange == "" {
		return errors.New("invalid exchange")
	}
	if s.Type == "" {
		return errors.New("invalid type")
	}

	return nil
}

func (s *securityAll) GetType() string {
	return s.Type
}
func (s *securityAll) GetExchange() string {
	return s.Exchange
}
