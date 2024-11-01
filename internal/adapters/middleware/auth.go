package middleware

import (
	"assetio/internal/port"
	"net/http"
	"strconv"
	"strings"
	"time"

	"golang.org/x/exp/slices"
)

type middleware struct {
	apiKeys  []string
	tokenIns port.Token
}

func New(apiKeys []string, tokenIns port.Token) port.Auth {
	return &middleware{
		apiKeys:  apiKeys,
		tokenIns: tokenIns,
	}
}

func (m *middleware) ValidateApiKey() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqApiKey := r.URL.Query().Get("key")
		if reqApiKey == "" {
			http.Error(w, "api key is required", http.StatusUnauthorized)
			return
		}

		if !slices.Contains(m.apiKeys, reqApiKey) {
			http.Error(w, "api key is invalid", http.StatusUnauthorized)
			return
		}

		w.Header().Set("Content-Type", "application/json")

	})
}

func (m *middleware) ValidateAccessToken() http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessToken := r.Header.Get("Authorization")
		token := m.exactToken(accessToken)
		if token == "" {
			http.Error(w, "authorization header required", http.StatusUnauthorized)
			return
		}

		userid, expiresAt, err := m.tokenIns.GetAccessTokenData(token)
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		if userid == 0 {
			http.Error(w, "incorrect token", http.StatusBadRequest)
			return
		}

		if expiresAt.Before(time.Now()) {
			http.Error(w, "token is expired", http.StatusBadRequest)
			return
		}

		queryParams := r.URL.Query()
		queryParams.Set("uid", strconv.Itoa(userid))
		r.URL.RawQuery = queryParams.Encode()
	})

}

func (m *middleware) exactToken(token string) string {
	parts := strings.SplitN(token, " ", 2)
	if len(parts) == 2 && parts[0] == "Bearer" {
		return parts[1]
	}
	return ""
}
