package app_http_server

import (
	"github.com/gin-gonic/gin"
)

type ApiVersion string

var (
	ApiVersion1 = ApiVersion("v1")
	ApiVersion2 = ApiVersion("v2")
	ApiVersion3 = ApiVersion("v3")
)

type ApiVersinRouter struct {
	ApiGroup *gin.RouterGroup
	ApiVersion
}

func NewApiVersinRouter(apiVersion ApiVersion, apiGroup *gin.RouterGroup) *ApiVersinRouter {
	return &ApiVersinRouter{
		ApiGroup:   apiGroup,
		ApiVersion: apiVersion,
	}
}

func (r *ApiVersinRouter) RegisterRouters(group *gin.RouterGroup, routes ...Route) {
	r.ApiGroup = group

	for _, route := range routes {
		handlers := make([]gin.HandlerFunc, 0, len(route.Middlewares)+1)

		handlers = append(handlers, route.Middlewares...)
		handlers = append(handlers, route.Handler)

		r.ApiGroup.Handle(route.Method, route.Path, handlers...)
	}
}
