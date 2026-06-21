package app_http_middleware

import (
	"fmt"

	app_loger "github.com/DanielTitovsky/rivulet-backend.git/internal/app/loger"
	app_http_response "github.com/DanielTitovsky/rivulet-backend.git/internal/app/transport/http/response"
	"github.com/gin-gonic/gin"
)

func Panic() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			ctx := c.Request.Context()

			log := app_loger.FromContext(ctx)
			responseHendler := app_http_response.NewHTTPResponseHandler(log, c.Writer)

			if p := recover(); p != nil {
				fmt.Print("\n")
				fmt.Print("\n")
				fmt.Print(p)
				fmt.Print("\n")
				fmt.Print("\n")
				responseHendler.PanicResponse(p, "during handle HTTP request got unexpected panic")
			}
		}()

		c.Next()
	}
}
