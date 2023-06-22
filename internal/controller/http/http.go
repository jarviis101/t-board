package http

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"t-board/internal/controller/http/graphql/directives"
	"t-board/internal/controller/http/graphql/graph"
	"t-board/internal/controller/http/graphql/transformers"
	"t-board/internal/controller/http/router"
	"t-board/internal/controller/http/validator"
	"t-board/internal/usecase"
	"t-board/pkg"
)

func RunServer(v *validator.Validator, sc pkg.Server, u usecase.UserUseCase, b usecase.BoardUseCase) {
	e := echo.New()
	e.Validator = v
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	collectRESTRoutes(e, v, u)
	collectGraphQLRoutes(e, sc, u, b)

	host := fmt.Sprintf(":%s", sc.Port)
	e.Logger.Fatal(e.Start(host))
}

func collectRESTRoutes(e *echo.Echo, v *validator.Validator, u usecase.UserUseCase) {
	authGroup := e.Group("/api/auth")
	userRouter := router.CreateUserRouteManager(authGroup, v, u)
	userRouter.PopulateRoutes()
}

func collectGraphQLRoutes(e *echo.Echo, sc pkg.Server, u usecase.UserUseCase, b usecase.BoardUseCase) {
	userTransformer := transformers.CreateUserTransformer()
	boardTransformer := transformers.CreateBoardTransformer()

	resolver := graph.CreateResolver(u, b, userTransformer, boardTransformer)
	c := graph.Config{Resolvers: resolver}
	c.Directives.Auth = directives.Auth
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(c))
	pg := playground.Handler("GraphQL playground", "/query")

	graphqlGroup := e.Group("")
	graphqlRouter := router.CreateGraphqlRouterManager(graphqlGroup, srv, pg, sc)
	graphqlRouter.PopulateRoutes()
}
