package request

import (
	"assetio/internal/port"
	"encoding/json"
	"errors"
	"net/http"
)

func NewSecuritySearch(r *http.Request) (port.SecuritySearchClientRequest, error) {
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

func (s *securitySearch) GetType() string {
	return s.Type
}
func (s *securitySearch) GetExchange() string {
	return s.Exchange
}

func (s *securitySearch) GetSearch() string {
	return s.Search
}
