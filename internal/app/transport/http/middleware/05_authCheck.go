package app_http_middleware

import (
	"context"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
	app_loger "github.com/DanielTitovsky/rivulet-backend.git/internal/app/loger"
	app_http_response "github.com/DanielTitovsky/rivulet-backend.git/internal/app/transport/http/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const AuthUserIdContextKey = "authUserId"

type TokenService interface {
	ValidateToken(tokenString string, tokenType string) (*domain.TokenClaims, error)
}

func Auth(tokenService TokenService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		log := app_loger.FromContext(ctx)
		responseHandler := app_http_response.NewHTTPResponseHandler(log, c.Writer)

		accessToken, err := c.Cookie("accessToken")
		if err != nil {
			responseHandler.ErrorResponse(err, "Access token not found")
			c.Abort()
			return
		}

		claims, err := tokenService.ValidateToken(accessToken, "access")
		if err != nil {
			responseHandler.ErrorResponse(err, "Invalid access token")
			c.Abort()
			return
		}

		ctx = context.WithValue(ctx, AuthUserIdContextKey, claims.UserId)
		c.Request = c.Request.WithContext(ctx)

		c.Set(AuthUserIdContextKey, claims.UserId)

		c.Next()
	}
}

func GetAuthUserId(c *gin.Context) (*uuid.UUID, bool) {
	userIdValue, exists := c.Get(AuthUserIdContextKey)
	if !exists {
		return nil, false
	}

	userId, ok := userIdValue.(uuid.UUID)
	if !ok {
		return nil, false
	}

	return &userId, true
}
