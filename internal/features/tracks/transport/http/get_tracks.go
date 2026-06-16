package tracks_transport_http

import (
	"net/http"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
	app_loger "github.com/DanielTitovsky/rivulet-backend.git/internal/app/loger"
	app_http_response "github.com/DanielTitovsky/rivulet-backend.git/internal/app/transport/http/response"
	app_http_utils "github.com/DanielTitovsky/rivulet-backend.git/internal/app/transport/http/utils"
	"github.com/gin-gonic/gin"
)

type GetTracksResponse []TrackDTOResponse

func (h *TrackHttpHandler) GetTracks(c *gin.Context) {
	ctx := c.Request.Context()
	log := app_loger.FromContext(ctx)

	responseHandler := app_http_response.NewHTTPResponseHandler(log, c.Writer)

	trackFilters, err := app_http_utils.GetQueryаFilter[domain.TrackFilters](c)

	if err != nil {
		responseHandler.ErrorResponse(err, "Invalid filter params")
		return
	}

	trackDomains, err := h.trackService.GetTracks(ctx, trackFilters)

	if err != nil {
		responseHandler.ErrorResponse(err, "Failed to get tracks")
		return
	}

	responseTracks := GetTracksResponse(tracksDTOFromDomain(trackDomains))

	responseHandler.JSONResponse(app_http_response.Response{
		Status: http.StatusOK,
		Data:   responseTracks,
	})
}
