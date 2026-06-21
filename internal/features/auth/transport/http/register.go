package auth_transport_http

import (
	"net/http"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
	app_loger "github.com/DanielTitovsky/rivulet-backend.git/internal/app/loger"
	app_http_response "github.com/DanielTitovsky/rivulet-backend.git/internal/app/transport/http/response"
	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (r RegisterRequest) ToDomain() domain.User {
	email := r.Email
	password := r.Password

	return domain.NewUserUninitialized(
		&email,
		r.Name,
		&password,
	)
}

func (h *AuthHttpHandler) Register(c *gin.Context) {
	ctx := c.Request.Context()
	log := app_loger.FromContext(ctx)

	responseHandler := app_http_response.NewHTTPResponseHandler(log, c.Writer)

	var request RegisterRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		responseHandler.ErrorResponse(err, "Invalid register body")
		return
	}

	user, accessToken, refreshToken, err := h.authService.Register(ctx, request.ToDomain())
	if err != nil {
		responseHandler.ErrorResponse(err, "Failed to register")
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "refreshToken",
		Value:    refreshToken.TokenString,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
		MaxAge:   int(RefreshExpired),
	})

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "accessToken",
		Value:    accessToken.TokenString,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
		MaxAge:   int(AccessExpired),
	})

	responseHandler.JSONResponse(app_http_response.Response{
		Status: http.StatusCreated,
		Data:   userDTOFromDomain(user),
	})
}
