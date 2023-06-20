package http

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"t-mail/internal/controller/http/graphql/directives"
	"t-mail/internal/controller/http/graphql/graph"
	"t-mail/internal/controller/http/router"
	"t-mail/internal/controller/http/validator"
	"t-mail/internal/usecase"
)

func RunServer(useCase usecase.UserUseCase, v *validator.Validator) {
	e := echo.New()
	e.Validator = v
	e.Use(middleware.Logger())

	collectRESTRoutes(e, v, useCase)
	collectGraphQLRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}

func collectRESTRoutes(e *echo.Echo, v *validator.Validator, useCase usecase.UserUseCase) {
	authGroup := e.Group("/api/auth")
	userRouter := router.CreateUserRouteManager(authGroup, v, useCase)
	userRouter.PopulateRoutes()
}

func collectGraphQLRoutes(e *echo.Echo) {
	c := graph.Config{Resolvers: &graph.Resolver{}}
	c.Directives.Auth = directives.Auth
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(c))
	pg := playground.Handler("GraphQL playground", "/query")
	graphqlGroup := e.Group("")
	graphqlRouter := router.CreateGraphqlRouterManager(graphqlGroup, srv, pg)
	graphqlRouter.PopulateRoutes()
}
