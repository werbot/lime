package jwtutil

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

// TokenMetadata struct to describe metadata in JWT.
type TokenMetadata struct {
	ID        string
	IssuedAt  float64
	ExpiresAt float64
}

var (
	ErrUnexpectedSigningMethod = errors.New("unexpected signing method")
	ErrVerificationError       = errors.New("verification error")
	ErrInvalidToken            = errors.New("the token is invalid")
)

// ExtractMetadata func to extract metadata from JWT.
func ExtractMetadata(tokenKey string) (*TokenMetadata, error) {
	token, err := jwt.Parse(tokenKey, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, ErrUnexpectedSigningMethod
		}
		return publicKey, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return &TokenMetadata{
			ID:        claims["id"].(string),
			IssuedAt:  claims["iat"].(float64),
			ExpiresAt: claims["exp"].(float64),
		}, nil
	}

	return nil, ErrInvalidToken
}
