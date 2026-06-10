package auth_transport_http

import (
	"net/http"
	"time"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
	app_loger "github.com/DanielTitovsky/rivulet-backend.git/internal/app/loger"
	app_http_request "github.com/DanielTitovsky/rivulet-backend.git/internal/app/transport/http/request"
	app_http_response "github.com/DanielTitovsky/rivulet-backend.git/internal/app/transport/http/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//TODO:Добавить RequestWithJwt чтобы сдлеать функцию по респонсу меньше если мы туда добавляем токены jwt

type RegisterAuthRequest struct {
	Name     string `json:"name" validate:"required,min=3,max=25"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RegisterAuthResponse struct {
	Id    uuid.UUID `json:"id" validate:"required"`
	Name  string    `json:"name" validate:"required"`
	Email string    `json:"email"`
}

var (
	AccessTokenExpires         = 15 * time.Minute
	RefreshTokenExpiredExpires = 30 * 24 * time.Hour
)

func (h *AuthHttpHandler) RegisterUser(c *gin.Context) {
	ctx := c.Request.Context()
	log := app_loger.FromContext(ctx)

	responseHandler := app_http_response.NewHTTPResponseHandler(log, c.Writer)

	log.Debug("Invoke Create Auth handler")

	var requestRegUser RegisterAuthRequest

	if err := app_http_request.DecodeAndValidate(*c.Request, requestRegUser); err != nil {
		responseHandler.ErrorResponse(err, "Invalid request")
		return
	}

	userDomain := domainFromDTO(requestRegUser)
	userDomain, err := h.authService.RegisterUser(ctx, userDomain)

	if err != nil {
		responseHandler.ErrorResponse(err, "Failed to create User")
		return
	}

	accessToken, refreshToken, err := h.authService.GenerateTokens(ctx, AccessTokenExpires, RefreshTokenExpiredExpires, userDomain)

	if err != nil {
		responseHandler.ErrorResponse(err, "Failed to generate tokens")
		return
	}

	_, err = h.authService.SaveToken(ctx, refreshToken)

	if err != nil {
		responseHandler.ErrorResponse(err, "Failed to save token")
	}

	responseUser := dtoFromDomain(userDomain)

	res := app_http_response.Response{
		Status: http.StatusCreated,
		Cookie: []*http.Cookie{
			{
				Name:     "refreshToken",
				Value:    string(refreshToken.TokenString),
				Path:     "/",
				HttpOnly: true,
				Secure:   false,
				SameSite: http.SameSiteStrictMode,
				MaxAge:   int(RefreshTokenExpiredExpires.Seconds()),
			},
			{
				Name:     "accessToken",
				Value:    string(accessToken.TokenString),
				Path:     "/",
				HttpOnly: true,
				Secure:   false,
				SameSite: http.SameSiteStrictMode,
				MaxAge:   int(AccessTokenExpires.Seconds()),
			},
		},
		Data: struct {
			User RegisterAuthResponse
		}{
			User: responseUser,
		},
	}

	responseHandler.JSONResponse(res)
}

func domainFromDTO(dto RegisterAuthRequest) domain.User {
	return domain.NewUserUninitialized(&dto.Email, dto.Name, &dto.Password)
}

func dtoFromDomain(user domain.User) RegisterAuthResponse {
	return RegisterAuthResponse{
		Id:    user.Id,
		Name:  user.Name,
		Email: *user.Email,
	}
}
