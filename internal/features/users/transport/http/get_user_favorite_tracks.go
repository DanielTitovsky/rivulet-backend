package users_transport_http

import (
	"net/http"

	app_loger "github.com/DanielTitovsky/rivulet-backend.git/internal/app/loger"
	app_http_response "github.com/DanielTitovsky/rivulet-backend.git/internal/app/transport/http/response"
	app_http_utils "github.com/DanielTitovsky/rivulet-backend.git/internal/app/transport/http/utils"
	"github.com/gin-gonic/gin"
)

type GetUserFavoriteTracksResponse []UserDTOResponse

func (h *UsersHttpHandler) GetUserFavoriteTracks(c *gin.Context) {
	ctx := c.Request.Context()
	log := app_loger.FromContext(ctx)

	responseHandler := app_http_response.NewHTTPResponseHandler(log, c.Writer)

	userId, err := app_http_utils.GetQueryParamsUUID(c, "id")

	if err != nil {
		responseHandler.ErrorResponse(err, "Invalid user id")
		return
	}

	if userId == nil {
		responseHandler.ErrorResponse(err, "Invalid user id")
		return
	}

	tracksDomain, err := h.userService.GetUserFavoriteTracks(ctx, *userId)

	if err != nil {
		responseHandler.ErrorResponse(err, "Failed to get user favorite tracks")
		return
	}

	responseTracks := tracksDTOFromDomain(tracksDomain)

	responseHandler.JSONResponse(app_http_response.Response{
		Status: http.StatusOK,
		Data:   responseTracks,
	})
}
