package playlist_transport_http

import (
	"net/http"

	app_loger "github.com/DanielTitovsky/rivulet-backend.git/internal/app/loger"
	app_http_response "github.com/DanielTitovsky/rivulet-backend.git/internal/app/transport/http/response"
	"github.com/gin-gonic/gin"
)

type GetPlaylistsResponse []PlaylistDTOResponce

func (h *PlaylistHttpHandler) GetPlaylists(c *gin.Context) {
	ctx := c.Request.Context()
	log := app_loger.FromContext(ctx)

	responseHandler := app_http_response.NewHTTPResponseHandler(log, c.Writer)

	playlistsDomain, err := h.PlaylistService.GetPlaylists(ctx)

	if err != nil {
		responseHandler.ErrorResponse(err, "Failed to get playlists")
		return
	}

	responsePlaylists := make(GetPlaylistsResponse, 0, len(playlistsDomain))

	for _, playlist := range playlistsDomain {
		responsePlaylists = append(responsePlaylists, playlistDTOFromDomain(playlist))
	}

	responseHandler.JSONResponse(app_http_response.Response{
		Status: http.StatusOK,
		Data:   responsePlaylists,
	})
}
