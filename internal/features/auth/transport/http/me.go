package auth_transport_http

import (
	"net/http"

	app_loger "github.com/DanielTitovsky/rivulet-backend.git/internal/app/loger"
	app_http_response "github.com/DanielTitovsky/rivulet-backend.git/internal/app/transport/http/response"
	"github.com/gin-gonic/gin"
)

func (h *AuthHttpHandler) Me(c *gin.Context) {
	ctx := c.Request.Context()
	log := app_loger.FromContext(ctx)

	responseHandler := app_http_response.NewHTTPResponseHandler(log, c.Writer)

	cookie, err := c.Cookie("accessToken")
	if err != nil {
		responseHandler.ErrorResponse(err, "Access token not found")
		return
	}

	user, err := h.authService.Me(ctx, cookie)
	if err != nil {
		responseHandler.ErrorResponse(err, "Failed to get current user")
		return
	}

	responseHandler.JSONResponse(app_http_response.Response{
		Status: http.StatusOK,
		Data:   userDTOFromDomain(user),
	})
}
