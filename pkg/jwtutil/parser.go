package jwtutil

import (
	"errors"

	"github.com/gofiber/fiber/v2"
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

func extractClaimsMetadata(claims jwt.MapClaims) (*TokenMetadata, error) {
	id, ok := claims["id"].(string)
	if !ok {
		return nil, ErrInvalidToken
	}
	iat, ok := claims["iat"].(float64)
	if !ok {
		return nil, ErrInvalidToken
	}
	exp, ok := claims["exp"].(float64)
	if !ok {
		return nil, ErrInvalidToken
	}

	return &TokenMetadata{
		ID:        id,
		IssuedAt:  iat,
		ExpiresAt: exp,
	}, nil
}

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
		return extractClaimsMetadata(claims)
	}

	return nil, ErrInvalidToken
}

// ExtractMetadataFiber extracts the metadata from a JWT token within a Fiber context.
func ExtractMetadataFiber(c *fiber.Ctx) (*TokenMetadata, error) {
	token := c.Locals("jwt").(*jwt.Token)
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return extractClaimsMetadata(claims)
	}
	return nil, ErrInvalidToken
}
