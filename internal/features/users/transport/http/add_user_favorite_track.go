package users_transport_http

import (
	"errors"
	"net/http"

	app_loger "github.com/DanielTitovsky/rivulet-backend.git/internal/app/loger"
	app_http_response "github.com/DanielTitovsky/rivulet-backend.git/internal/app/transport/http/response"
	app_http_utils "github.com/DanielTitovsky/rivulet-backend.git/internal/app/transport/http/utils"
	"github.com/gin-gonic/gin"
)

func (uh *UsersHttpHandler) AddTrackToFavorite(c *gin.Context) {
	ctx := c.Request.Context()
	log := app_loger.FromContext(ctx)

	responseHandler := app_http_response.NewHTTPResponseHandler(log, c.Writer)

	userId, err := app_http_utils.GetQueryParamsUUID(c, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "Invalid user id")
		return
	}

	if userId == nil {
		responseHandler.ErrorResponse(errors.New("user id is empty"), "Invalid user id")
		return
	}

	trackId, err := app_http_utils.GetQueryParamsUUID(c, "trackId")
	if err != nil {
		responseHandler.ErrorResponse(err, "Invalid track id")
		return
	}

	if trackId == nil {
		responseHandler.ErrorResponse(errors.New("track id is empty"), "Invalid track id")
		return
	}

	err = uh.userService.AddTrackToFavorite(ctx, *userId, *trackId)
	if err != nil {
		responseHandler.ErrorResponse(err, "Failed to add track to favorite")
		return
	}

	responseHandler.JSONResponse(app_http_response.Response{
		Status: http.StatusCreated,
		Data:   "Track added to favorite",
	})
}
