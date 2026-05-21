package app_http_middleware

import (
	"context"

	app_loger "github.com/DanielTitovsky/rivulet-backend.git/internal/app/loger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Logger(log *app_loger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestId := c.Request.Header.Get("requestIdHeader")

		l := log.With(
			zap.String("request_id", requestId),
			zap.String("url", c.Request.RequestURI),
		)

		ctx := context.WithValue(c.Request.Context(), "logger", l)

		c.Request.WithContext(ctx)

		c.Next()
	}
}
