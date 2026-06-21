package auth_transport_http

import (
	"net/http"
	"time"

	app_loger "github.com/DanielTitovsky/rivulet-backend.git/internal/app/loger"
	app_http_response "github.com/DanielTitovsky/rivulet-backend.git/internal/app/transport/http/response"
	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func tokenCookieMaxAge(expiresAt time.Time) int {
	maxAge := int(time.Until(expiresAt).Seconds())

	if maxAge < 0 {
		return 0
	}

	return maxAge
}

func (h *AuthHttpHandler) Login(c *gin.Context) {
	ctx := c.Request.Context()
	log := app_loger.FromContext(ctx)

	responseHandler := app_http_response.NewHTTPResponseHandler(log, c.Writer)

	var request LoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		responseHandler.ErrorResponse(err, "Invalid login body")
		return
	}

	user, accessToken, refreshToken, err := h.authService.Login(ctx, request.Email, request.Password)
	if err != nil {
		responseHandler.ErrorResponse(err, "Failed to login")
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
