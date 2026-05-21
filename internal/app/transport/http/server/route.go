package app_http_server

import (
	"github.com/gin-gonic/gin"
)

type Route struct {
	Method      string
	Path        string
	Handler     gin.HandlerFunc
	Middlewares []gin.HandlerFunc
}

func NewRoute(method string, path string, handler gin.HandlerFunc) *Route {
	return &Route{
		Method:  method,
		Path:    path,
		Handler: handler,
	}
}
