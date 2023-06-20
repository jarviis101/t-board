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
	group     *echo.Group
	validator *validator.Validator
	useCase   usecase.UserUseCase
}

func CreateUserRouteManager(
	g *echo.Group,
	v *validator.Validator,
	useCase usecase.UserUseCase,
) RouteManager {
	return &userRouteManager{g, v, useCase}
}

func (r *userRouteManager) PopulateRoutes() {
	r.group.Add("POST", "/register", r.register)
	r.group.Add("POST", "/login", r.login)
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
