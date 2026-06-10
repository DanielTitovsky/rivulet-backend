package users_transport_http

import (
	"net/http"

	app_loger "github.com/DanielTitovsky/rivulet-backend.git/internal/app/loger"
	app_http_response "github.com/DanielTitovsky/rivulet-backend.git/internal/app/transport/http/response"
	app_http_utils "github.com/DanielTitovsky/rivulet-backend.git/internal/app/transport/http/utils"
	"github.com/gin-gonic/gin"
)

type GetUserResponse UserDTOResponse

func (uh *UsersHttpHandler) GetUser(c *gin.Context) {
	ctx := c.Request.Context()
	log := app_loger.FromContext(ctx)

	responseHandler := app_http_response.NewHTTPResponseHandler(log, c.Writer)

	log.Debug("Invoke Create handler")

	userId, err := app_http_utils.GetQueryParamsId(c, "id")

	if err != nil {
		responseHandler.ErrorResponse(err, "Invaled parametrs")
		return
	}

	userDomain, err := uh.userService.GetUser(ctx, *userId)

	if err != nil {
		responseHandler.ErrorResponse(err, "Failed to get user")
		return
	}

	responseUser := userDTOFromDomain(userDomain)

	responseHandler.JSONResponse(app_http_response.Response{Status: http.StatusOK, Data: responseUser})
}
