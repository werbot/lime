package middleware

import (
	"strings"

	jwtMiddleware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"

	"github.com/werbot/lime/pkg/jwtutil"
	"github.com/werbot/lime/pkg/webutil"
)

// JWTProtected is ...
func JWTProtected(section string) func(*fiber.Ctx) error {
	return jwtMiddleware.New(jwtMiddleware.Config{
		ErrorHandler: jwtError,
		ContextKey:   "jwt",
		TokenLookup:  "cookie:" + section,
		SigningKey: jwtMiddleware.SigningKey{
			JWTAlg: jwtMiddleware.RS256,
			Key:    jwtutil.PrivateKey().Public(),
		},
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if strings.HasPrefix(c.Path(), "/api") || strings.HasPrefix(c.Path(), "/_/api") {
		if err.Error() == "Missing or malformed token" {
			return webutil.Response(c, fiber.StatusBadRequest, "Bad request", err.Error())
		}
		return webutil.Response(c, fiber.StatusUnauthorized, "Unauthorized", err.Error())
	}

	return webutil.Response(c, fiber.StatusUnauthorized, "Unauthorized", err.Error())
}
