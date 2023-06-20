package router

import (
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
	"t-mail/internal/controller/http/types"
	"t-mail/internal/controller/http/validator"
	"t-mail/internal/usecase"
)

type RouteManager interface {
	PopulateRoutes()
}

type userRouteManager struct {
	router    *echo.Router
	validator *validator.Validator
	useCase   usecase.UserUseCase
}

func CreateUserRouteManager(
	r *echo.Router,
	v *validator.Validator,
	useCase usecase.UserUseCase,
) RouteManager {
	return &userRouteManager{r, v, useCase}
}

func (r *userRouteManager) PopulateRoutes() {
	r.router.Add("POST", "/register", r.register)
	r.router.Add("POST", "/login", r.login)
	r.router.Add("GET", "/me", r.me)
}

func (r *userRouteManager) register(c echo.Context) error {
	u := &types.CreateUser{}
	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := r.validator.Validate(u); err != nil {
		return err
	}
	if err := r.useCase.Register(context.Background(), u.Name, u.Email, u.Password); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return nil
}

func (r *userRouteManager) login(c echo.Context) error {
	u := &types.LoginUserRequest{}
	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := r.validator.Validate(u); err != nil {
		return err
	}

	token, err := r.useCase.Login(context.Background(), u.Email, u.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	response := &types.LoginUserResponse{Token: token}
	return c.JSON(http.StatusOK, response)
}

func (r *userRouteManager) me(c echo.Context) error {
	return nil
}
