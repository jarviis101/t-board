package router

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/labstack/echo/v4"
	"net/http"
	"t-mail/internal/controller/http/middleware"
)

type graphqlRouterManager struct {
	group  *echo.Group
	server *handler.Server
	pg     http.HandlerFunc
}

func CreateGraphqlRouterManager(group *echo.Group, server *handler.Server, pg http.HandlerFunc) RouteManager {
	return &graphqlRouterManager{group, server, pg}
}

func (r *graphqlRouterManager) PopulateRoutes() {
	r.group.Add("GET", "/graphql", r.playground)
	r.group.Add("POST", "/query", r.query, middleware.AuthMiddleware)
}

func (r *graphqlRouterManager) query(c echo.Context) error {
	r.server.ServeHTTP(c.Response(), c.Request())
	return nil
}

func (r *graphqlRouterManager) playground(c echo.Context) error {
	r.pg.ServeHTTP(c.Response(), c.Request())
	return nil
}
