package app_http_middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const requestIdHeader = "X-Request-ID"

func RequestId() gin.HandlerFunc {

	return func(c *gin.Context) {
		requestId := c.Request.Header.Get(requestIdHeader)
		if requestId == "" {
			requestId = uuid.NewString()
		}

		c.Request.Header.Set(requestIdHeader, requestId)
		c.Writer.Header().Set(requestIdHeader, requestId)

		c.Next()
	}
}
