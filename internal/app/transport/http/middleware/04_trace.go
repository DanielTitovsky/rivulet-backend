package app_http_middleware

import (
	"time"

	app_loger "github.com/DanielTitovsky/rivulet-backend.git/internal/app/loger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Trace() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		log := app_loger.FromContext(ctx)

		before := time.Now()
		log.Debug(
			">>> incoming HTTP request",
			zap.Time("time", before.UTC()),
		)

		c.Next()

		log.Debug(
			">>> done HTTP request",
			zap.Int("status_code", c.Writer.Status()),
			zap.Duration("latency", time.Now().Sub(before)),
		)
	}
}
