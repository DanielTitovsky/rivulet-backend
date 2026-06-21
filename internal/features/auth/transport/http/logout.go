package auth_transport_http

import (
	"net/http"

	app_loger "github.com/DanielTitovsky/rivulet-backend.git/internal/app/loger"
	app_http_response "github.com/DanielTitovsky/rivulet-backend.git/internal/app/transport/http/response"
	"github.com/gin-gonic/gin"
)

func (h *AuthHttpHandler) Logout(c *gin.Context) {
	ctx := c.Request.Context()
	log := app_loger.FromContext(ctx)

	responseHandler := app_http_response.NewHTTPResponseHandler(log, c.Writer)

	cookie, err := c.Cookie("refreshToken")
	if err != nil {
		responseHandler.ErrorResponse(err, "Refresh token not found")
		return
	}

	err = h.authService.Logout(ctx, cookie)
	if err != nil {
		responseHandler.ErrorResponse(err, "Failed to logout")
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "refreshToken",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
		MaxAge:   -1,
	})

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "accessToken",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
		MaxAge:   -1,
	})

	responseHandler.JSONResponse(app_http_response.Response{
		Status: http.StatusOK,
		Data:   "logout success",
	})
}
