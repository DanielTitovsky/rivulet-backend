package users_transport_http

import (
	"fmt"
	"net/http"

	app_loger "github.com/DanielTitovsky/rivulet-backend.git/internal/app/loger"
	app_http_response "github.com/DanielTitovsky/rivulet-backend.git/internal/app/transport/http/response"
	app_http_utils "github.com/DanielTitovsky/rivulet-backend.git/internal/app/transport/http/utils"
	"github.com/gin-gonic/gin"
)

type DeleteUserResponse UserDTOResponse

func (uh *UsersHttpHandler) DeleteUser(c *gin.Context) {
	fmt.Println("DELETE USER HANDLER")
	ctx := c.Request.Context()
	log := app_loger.FromContext(ctx)

	responseHandler := app_http_response.NewHTTPResponseHandler(log, c.Writer)

	log.Debug("Invoke Create handler")

	userId, err := app_http_utils.GetQueryParamsId(c, "id")

	err = uh.userService.DeleteUser(ctx, *userId)

	if err != nil {
		responseHandler.ErrorResponse(err, "Failed to create user")
		return
	}

	responseHandler.JSONResponse(app_http_response.Response{Status: http.StatusOK, Data: true})
}
