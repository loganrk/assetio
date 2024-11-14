package middleware

import (
	"assetio/internal/port"
	"net/http"
	"strconv"
	"strings"
	"time"

	"golang.org/x/exp/slices"
)

// middleware struct holds API keys and a Token instance for handling API authentication.
type middleware struct {
	apiKeys  []string   // List of valid API keys
	tokenIns port.Token // Token interface for handling token-related operations
}

// New creates a new middleware instance with the given API keys and Token instance.
func New(apiKeys []string, tokenIns port.Token) port.Auth {
	return &middleware{
		apiKeys:  apiKeys,  // Initialize with the provided API keys
		tokenIns: tokenIns, // Initialize with the provided Token instance
	}
}

// ValidateApiKey returns an HTTP handler that checks if the request contains a valid API key.
func (m *middleware) ValidateApiKey() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract the API key from the query parameters
		reqApiKey := r.URL.Query().Get("key")
		if reqApiKey == "" {
			// If no API key is provided, respond with Unauthorized error.
			http.Error(w, "api key is required", http.StatusUnauthorized)
			return
		}

		// Check if the provided API key is valid
		if !slices.Contains(m.apiKeys, reqApiKey) {
			// If the API key is invalid, respond with Unauthorized error.
			http.Error(w, "api key is invalid", http.StatusUnauthorized)
			return
		}

		// Set content type to JSON for successful validation.
		w.Header().Set("Content-Type", "application/json")
	})
}

// ValidateAccessToken returns an HTTP handler that validates the access token from the Authorization header.
func (m *middleware) ValidateAccessToken() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract the access token from the Authorization header
		accessToken := r.Header.Get("Authorization")
		token := m.exactToken(accessToken)
		if token == "" {
			// If no token is provided or token is malformed, respond with Unauthorized error.
			http.Error(w, "authorization header required", http.StatusUnauthorized)
			return
		}

		// Get user ID and expiration time from the token
		userid, expiresAt, err := m.tokenIns.GetAccessTokenData(token)
		if err != nil {
			// If there’s an internal error while fetching token data, respond with internal server error.
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		// If the token is invalid (no user ID found), respond with Bad Request error.
		if userid == 0 {
			http.Error(w, "incorrect token", http.StatusBadRequest)
			return
		}

		// If the token is expired, respond with Bad Request error.
		if expiresAt.Before(time.Now()) {
			http.Error(w, "token is expired", http.StatusBadRequest)
			return
		}

		// Add user ID to the query parameters for further processing in the request lifecycle.
		queryParams := r.URL.Query()
		queryParams.Set("uid", strconv.Itoa(userid))
		r.URL.RawQuery = queryParams.Encode()
	})
}

// exactToken extracts the token from the Authorization header by splitting it into "Bearer <token>" format.
func (m *middleware) exactToken(token string) string {
	// Split the Authorization header to extract the token part.
	parts := strings.SplitN(token, " ", 2)
	if len(parts) == 2 && parts[0] == "Bearer" {
		return parts[1] // Return the token part of the "Bearer <token>" format.
	}
	return "" // Return an empty string if the token is not in the correct format.
}
