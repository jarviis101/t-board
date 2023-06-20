package router

import (
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
	"t-mail/internal/controller/http/types"
	"t-mail/internal/controller/http/validator"
	"t-mail/internal/usecase"
)

type UserRouteManager struct {
	router    *echo.Router
	validator *validator.Validator
	user      usecase.UserUseCase
}

func CreateUserRouteManager(
	r *echo.Router,
	v *validator.Validator,
	useCase usecase.UserUseCase,
) *UserRouteManager {
	return &UserRouteManager{r, v, useCase}
}

func (r *UserRouteManager) PopulateRoutes() {
	r.router.Add("POST", "/register", r.register)
}

func (r *UserRouteManager) register(c echo.Context) error {
	u := &types.User{}
	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := r.validator.Validate(u); err != nil {
		return err
	}
	if err := r.user.Register(context.Background(), u.Name, u.Email, u.Password); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return nil
}
