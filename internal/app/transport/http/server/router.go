package app_http_server

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type ApiVersion string

var (
	ApiVersion1 = ApiVersion("v1")
	ApiVersion2 = ApiVersion("v2")
	ApiVersion3 = ApiVersion("v3")
)

type ApiVersinRouter struct {
	ServeGin *gin.Engine
	ApiVersion
}

func NewApiVersinRouter(apiVersion ApiVersion) *ApiVersinRouter {
	return &ApiVersinRouter{
		ServeGin:   gin.New(),
		ApiVersion: apiVersion,
	}
}

func (r *ApiVersinRouter) RegisterRouters(routers ...Route) {
	for _, route := range routers {
		pattern := fmt.Sprintf("%s %s", route.Method, route.Path)

		r.ServeGin.Handle(route.Method, pattern, route.Handler)
	}
}
