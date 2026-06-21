package playlist_transport_http

import (
	app_loger "github.com/DanielTitovsky/rivulet-backend.git/internal/app/loger"
	app_http_response "github.com/DanielTitovsky/rivulet-backend.git/internal/app/transport/http/response"
	app_http_utils "github.com/DanielTitovsky/rivulet-backend.git/internal/app/transport/http/utils"
	"github.com/gin-gonic/gin"
)

func (h *PlaylistHttpHandler) GetPlaylist(c *gin.Context) {
	ctx := c.Request.Context()
	log := app_loger.FromContext(ctx)

	responseHandler := app_http_response.NewHTTPResponseHandler(log, c.Writer)

	playlistId, err := app_http_utils.GetQueryParamsUUID(c, "id")

	if err != nil {
		responseHandler.ErrorResponse(err, "Invalid id")
		return
	}

	playlistDomain, err := h.PlaylistService.GetPlaylist(ctx, *playlistId)

	if err != nil {
		responseHandler.ErrorResponse(err, "Failet to get plaulist")
		return
	}

	responsePlaylist := 
}
