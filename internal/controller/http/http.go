package http

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"t-mail/internal/controller/http/router"
	http_validator "t-mail/internal/controller/http/validator"
	"t-mail/internal/usecase"
)

func RunServer(useCase usecase.UserUseCase) {
	v := http_validator.CreateValidator(validator.New())
	e := echo.New()
	e.Validator = v
	e.Use(middleware.Logger())

	userRouter := router.CreateUserRouteManager(e.Router(), v, useCase)
	userRouter.PopulateRoutes()

	e.Logger.Fatal(e.Start(":8080"))
}
