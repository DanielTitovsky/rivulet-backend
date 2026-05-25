package users_transport_http

import (
	"net/http"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
	app_loger "github.com/DanielTitovsky/rivulet-backend.git/internal/app/loger"
	app_http_request "github.com/DanielTitovsky/rivulet-backend.git/internal/app/transport/http/request"
	app_http_response "github.com/DanielTitovsky/rivulet-backend.git/internal/app/transport/http/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateUserRequest struct {
	Name     string `json:"name" validate:"required,min=3,max=25"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type CreateUserResponse struct {
	Id    uuid.UUID `json:"id" validate:"required"`
	Name  string    `json:"name" validate:"required"`
	Email string    `json:"email"`
}

func (uh *UsersHttpHandler) RegisterUser(c *gin.Context) {
	ctx := c.Request.Context()
	log := app_loger.FromContext(ctx)

	responseHandler := app_http_response.NewHTTPResponseHandler(log, c.Writer)

	log.Debug("Invoke Create handler")

	var requestUser CreateUserRequest

	if err := app_http_request.DecodeAndValidate(*c.Request, &requestUser); err != nil {
		responseHandler.ErrorResponse(err, "Invalid request")
		return
	}

	userDomain := domainFromDTO(requestUser)

	userDomain, err := uh.userService.CreateUser(ctx, userDomain)

	if err != nil {
		responseHandler.ErrorResponse(err, "Failed to create user")
		return
	}

	responseUser := dtoFromDomain(userDomain)

	responseHandler.JSONResponse(app_http_response.Response{Status: http.StatusCreated, Data: responseUser})
}

func domainFromDTO(dto CreateUserRequest) domain.User {
	return domain.NewUserUninitialized(&dto.Email, dto.Name, &dto.Password)
}

func dtoFromDomain(user domain.User) CreateUserResponse {
	return CreateUserResponse{
		Id:    user.Id,
		Name:  user.Name,
		Email: *user.Email,
	}
}
