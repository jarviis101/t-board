package router

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/labstack/echo/v4"
	"net/http"
	"t-board/internal/controller/http/middleware"
	"t-board/pkg"
)

type graphqlRouterManager struct {
	group  *echo.Group
	server *handler.Server
	pg     http.HandlerFunc
	sc     pkg.Server
}

func CreateGraphqlRouterManager(
	group *echo.Group,
	server *handler.Server,
	pg http.HandlerFunc,
	s pkg.Server,
) RouteManager {
	return &graphqlRouterManager{group, server, pg, s}
}

func (r *graphqlRouterManager) PopulateRoutes() {
	r.group.Add("GET", "/graphql", r.playground)
	r.group.Add("POST", "/query", r.query, middleware.AuthMiddleware(r.sc.Secret))
}

func (r *graphqlRouterManager) query(c echo.Context) error {
	r.server.ServeHTTP(c.Response(), c.Request())
	return nil
}

func (r *graphqlRouterManager) playground(c echo.Context) error {
	r.pg.ServeHTTP(c.Response(), c.Request())
	return nil
}
