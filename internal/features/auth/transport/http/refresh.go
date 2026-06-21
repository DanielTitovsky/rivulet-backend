package auth_transport_http

import (
	"net/http"

	app_loger "github.com/DanielTitovsky/rivulet-backend.git/internal/app/loger"
	app_http_response "github.com/DanielTitovsky/rivulet-backend.git/internal/app/transport/http/response"
	"github.com/gin-gonic/gin"
)

func (h *AuthHttpHandler) Refresh(c *gin.Context) {
	ctx := c.Request.Context()
	log := app_loger.FromContext(ctx)

	responseHandler := app_http_response.NewHTTPResponseHandler(log, c.Writer)

	cookie, err := c.Cookie("refreshToken")
	if err != nil {
		responseHandler.ErrorResponse(err, "Refresh token not found")
		return
	}

	user, accessToken, refreshToken, err := h.authService.Refresh(ctx, cookie)
	if err != nil {
		responseHandler.ErrorResponse(err, "Failed to refresh")
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "refreshToken",
		Value:    refreshToken.TokenString,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
		MaxAge:   tokenCookieMaxAge(refreshToken.ExpiresAt),
	})

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "accessToken",
		Value:    accessToken.TokenString,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
		MaxAge:   tokenCookieMaxAge(accessToken.ExpiresAt),
	})

	responseHandler.JSONResponse(app_http_response.Response{
		Status: http.StatusOK,
		Data:   userDTOFromDomain(user),
	})
}
