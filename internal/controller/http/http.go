package http

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"t-board/internal/controller"
	"t-board/internal/controller/http/graphql/directives"
	"t-board/internal/controller/http/graphql/graph"
	"t-board/internal/controller/http/graphql/transformers"
	"t-board/internal/controller/http/router"
	"t-board/internal/controller/http/validator"
	"t-board/internal/usecase"
	"t-board/pkg"
)

type http struct {
	serverConfig     pkg.Server
	validator        *validator.Validator
	userUseCase      usecase.UserUseCase
	boardUseCase     usecase.BoardUseCase
	userTransformer  transformers.UserTransformer
	boardTransformer transformers.BoardTransformer
	echo             *echo.Echo
}

func CreateServer(
	sc pkg.Server,
	v *validator.Validator,
	u usecase.UserUseCase,
	b usecase.BoardUseCase,
) controller.Server {
	e := echo.New()
	e.Validator = v
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	baseTransformer := transformers.CreateBaseTransformer()
	ut := transformers.CreateUserTransformer(baseTransformer)
	bt := transformers.CreateBoardTransformer(baseTransformer)

	return &http{
		serverConfig:     sc,
		validator:        v,
		userUseCase:      u,
		boardUseCase:     b,
		userTransformer:  ut,
		boardTransformer: bt,
	}
}

func (h *http) RunServer() error {
	e := echo.New()

	h.appendRestRoutes(e)
	h.appendGraphqlRoutes(e)

	host := fmt.Sprintf(":%s", h.serverConfig.Port)
	return h.echo.Start(host)
}

func (h *http) appendRestRoutes(e *echo.Echo) {
	authGroup := e.Group("/api/auth")
	userRouter := router.CreateUserRouteManager(authGroup, h.validator, h.userUseCase)

	userRouter.PopulateRoutes()
}

func (h *http) appendGraphqlRoutes(e *echo.Echo) {
	resolver := graph.CreateResolver(h.userUseCase, h.boardUseCase, h.userTransformer, h.boardTransformer)
	c := graph.Config{Resolvers: resolver}
	c.Directives.Auth = directives.Auth
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(c))
	pg := playground.Handler("GraphQL playground", "/query")

	graphqlRouter := router.CreateGraphqlRouterManager(e.Group(""), srv, pg, h.serverConfig)

	graphqlRouter.PopulateRoutes()
}
