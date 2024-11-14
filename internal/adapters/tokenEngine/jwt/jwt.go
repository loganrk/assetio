package jwt

import (
	"assetio/internal/port"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type token struct {
	cipher port.Cipher
}

func New(cipherIns port.Cipher) port.Token {
	return &token{
		cipher: cipherIns,
	}
}

// GetAccessTokenData retrieves the user ID (uid) and expiration time (exp) from an encrypted access token.
func (t *token) GetAccessTokenData(encryptedToken string) (int, time.Time, error) {
	// Decrypt the encrypted token using the cipher instance
	tokenString, err := t.cipher.Decrypt(encryptedToken)
	if err != nil {
		// Return error if decryption fails
		return 0, time.Time{}, err
	}

	// Parse the decrypted token (without verification) to extract claims
	token, _, err := jwt.NewParser().ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		// Return error if parsing fails
		return 0, time.Time{}, err
	}

	// Extract claims from the parsed token if it is of type jwt.MapClaims
	if claims, ok := token.Claims.(jwt.MapClaims); ok {

		// Validate that the token type is "access"
		if tokenType, ok := claims["type"].(string); !ok || tokenType != "access" {
			// Return error if token type is not found or does not match "access"
			return 0, time.Time{}, errors.New("token type (`type`) not found or mismatch in token")
		}

		// Extract the user ID (uid) from the claims
		if uid, ok := claims["uid"].(float64); ok {
			// Extract the expiration time (exp) from the claims
			if exp, ok := claims["exp"].(float64); ok {
				// Convert the expiration time to a time.Time object
				expirationTime := time.Unix(int64(exp), 0)
				// Return the user ID and expiration time
				return int(uid), expirationTime, nil
			}
			// Return error if expiration time is not found in token
			return 0, time.Time{}, errors.New("expiration time (`exp`) not found in token")
		}

		// Return error if user ID is not found in token
		return 0, time.Time{}, errors.New("user id (`uid`) not found in token")
	}

	// Return error if claims are invalid or missing
	return 0, time.Time{}, errors.New("invalid token claims")
}
