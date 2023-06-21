package http

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"t-mail/internal/controller/http/graphql/directives"
	"t-mail/internal/controller/http/graphql/graph"
	"t-mail/internal/controller/http/router"
	"t-mail/internal/controller/http/validator"
	"t-mail/internal/usecase"
	"t-mail/pkg"
)

func RunServer(useCase usecase.UserUseCase, v *validator.Validator, serverConfig pkg.Server) {
	e := echo.New()
	e.Validator = v
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	collectRESTRoutes(e, v, useCase)
	collectGraphQLRoutes(e, useCase, serverConfig)

	host := fmt.Sprintf(":%s", serverConfig.Port)
	e.Logger.Fatal(e.Start(host))
}

func collectRESTRoutes(e *echo.Echo, v *validator.Validator, useCase usecase.UserUseCase) {
	authGroup := e.Group("/api/auth")
	userRouter := router.CreateUserRouteManager(authGroup, v, useCase)
	userRouter.PopulateRoutes()
}

func collectGraphQLRoutes(e *echo.Echo, useCase usecase.UserUseCase, serverConfig pkg.Server) {
	resolver := graph.CreateResolver(useCase)
	c := graph.Config{Resolvers: resolver}
	c.Directives.Auth = directives.Auth
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(c))
	pg := playground.Handler("GraphQL playground", "/query")
	graphqlGroup := e.Group("")
	graphqlRouter := router.CreateGraphqlRouterManager(graphqlGroup, srv, pg, serverConfig)
	graphqlRouter.PopulateRoutes()
}