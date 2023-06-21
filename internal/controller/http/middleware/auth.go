package middleware

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"t-board/internal/pkg/jwt"
)

const (
	header                  = "Authorization"
	excludedStringFromToken = "Bearer "
)

var claims *jwt.UserClaims

func AuthMiddleware(secret string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		jwtManager := jwt.CreateManager(secret)
		return func(c echo.Context) error {
			token := c.Request().Header.Get(header)
			if token == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "Not authorized")
			}
			accessToken := token[len(excludedStringFromToken):]
			verifiedClaims, err := jwtManager.Verify(accessToken)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
			}

			claims = verifiedClaims
			if err := next(c); err != nil {
				return err
			}

			return nil
		}
	}
}

func GetClaims() *jwt.UserClaims {
	// TODO provide claims must be with another way
	return claims
}
