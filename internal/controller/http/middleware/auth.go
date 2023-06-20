package middleware

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"t-mail/internal/pkg/jwt"
)

const (
	header                  = "Authorization"
	excludedStringFromToken = "Bearer "
)

var claims *jwt.UserClaims

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	jwtManager := jwt.CreateManager()
	return func(ctx echo.Context) error {
		token := ctx.Request().Header.Get(header)
		if token == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Not authorized")
		}
		accessToken := token[len(excludedStringFromToken):]
		c, err := jwtManager.Verify(accessToken)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}

		claims = c
		if err := next(ctx); err != nil {
			return err
		}

		return nil
	}
}

func GetClaims() *jwt.UserClaims {
	// TODO resolve
	return claims
}
