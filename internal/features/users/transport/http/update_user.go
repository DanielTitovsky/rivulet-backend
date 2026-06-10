package users_transport_http

import (
	"net/http"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
	app_loger "github.com/DanielTitovsky/rivulet-backend.git/internal/app/loger"
	app_http_request "github.com/DanielTitovsky/rivulet-backend.git/internal/app/transport/http/request"
	app_http_response "github.com/DanielTitovsky/rivulet-backend.git/internal/app/transport/http/response"
	app_http_utils "github.com/DanielTitovsky/rivulet-backend.git/internal/app/transport/http/utils"
	"github.com/gin-gonic/gin"
)

type UpdateUserRequest struct {
	Name     string `json:"name" validate:"required,min=3,max=25"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" `
}

type UpdateUserResponse UserDTOResponse

func (uh *UsersHttpHandler) UpdateUser(c *gin.Context) {
	ctx := c.Request.Context()
	log := app_loger.FromContext(ctx)

	responseHandler := app_http_response.NewHTTPResponseHandler(log, c.Writer)

	userId, err := app_http_utils.GetQueryParamsId(c, "id")

	if err != nil {
		responseHandler.ErrorResponse(err, "Invaled parametrs")
		return
	}

	log.Debug("Invoke Create handler")

	var requestUser UpdateUserRequest

	if err := app_http_request.DecodeAndValidate(*c.Request, &requestUser); err != nil {
		responseHandler.ErrorResponse(err, "Invalid request")
		return
	}

	userUpdate := userUpdateFromRequest(requestUser)

	userDomain, err := uh.userService.UpdateUser(ctx, *userId, userUpdate)

	if err != nil {
		responseHandler.ErrorResponse(err, "Failed to create user")
		return
	}

	responseUser := userDTOFromDomain(userDomain)

	responseHandler.JSONResponse(app_http_response.Response{Status: http.StatusOK, Data: responseUser})
}

func userUpdateFromRequest(request UpdateUserRequest) domain.UserUpdate {
	return domain.UserUpdate{
		Name:     request.Name,
		Email:    &request.Email,
		Password: request.Password,
	}
}
