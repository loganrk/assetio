package request

import (
	"assetio/internal/port"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

func NewSecurityGet(r *http.Request) (port.SecurityGetClientRequest, error) {
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

func (s *securityGet) GetSecuriryId() int {
	return s.SecurityId
}
