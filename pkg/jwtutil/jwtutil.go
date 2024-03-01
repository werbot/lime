package jwtutil

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// NewToken func for generate a new Access token.
func NewToken(id, expires string, credentials []string) (string, error) {
	duration, err := time.ParseDuration(expires)
	if err != nil {
		return "", err
	}

	iat := time.Now()
	exp := iat.Add(duration)

	claims := jwt.MapClaims{}
	claims["id"] = id
	claims["iat"] = jwt.NewNumericDate(iat)
	claims["exp"] = jwt.NewNumericDate(exp)

	for _, credential := range credentials {
		claims[credential] = true
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	accessToken, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
