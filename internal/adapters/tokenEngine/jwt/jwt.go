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

func (t *token) GetAccessTokenData(encryptedToken string) (int, time.Time, error) {
	tokenString, err := t.cipher.Decrypt(encryptedToken)
	if err != nil {
		return 0, time.Time{}, err
	}

	token, _, err := jwt.NewParser().ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return 0, time.Time{}, err

	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {

		if tokenType, ok := claims["type"].(string); !ok || tokenType != "access" {
			return 0, time.Time{}, errors.New("token type (`type`) not found or mismatch in token")
		}
		if uid, ok := claims["uid"].(float64); ok {
			if exp, ok := claims["exp"].(float64); ok {
				expirationTime := time.Unix(int64(exp), 0)
				return int(uid), expirationTime, nil
			}
			return 0, time.Time{}, errors.New("expiration time (`exp`) not found in token")
		}
		return 0, time.Time{}, errors.New("user id (`uid`) not found in token")
	}
	return 0, time.Time{}, errors.New("invalid token claims")
}
