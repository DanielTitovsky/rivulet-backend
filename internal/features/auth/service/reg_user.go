package auth_service

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

func (h *AuthHttpHandler) RegisterUser(ctx context.Context, user domain.User) (user domain.User, accessToken Token, refreshToken Token, error) {
	
	userDomain, err := h.userService.CreateUser(ctx, user)

	if err != nil {
		responseHandler.ErrorResponse(err, "Failed to create User")
		return
	}

	accessToken, refreshToken, err := h.tokenService.GenerateTokens(AccessTokenExpires, RefreshTokenExpiredExpires, userDomain)

	if err != nil {
		responseHandler.ErrorResponse(err, "Failed to generate tokens")
		return
	}

	_, err = h.tokenService.SaveToken(*refreshToken)

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