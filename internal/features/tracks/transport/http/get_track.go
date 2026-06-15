package tracks_transport_http

import (
	"net/http"

	app_loger "github.com/DanielTitovsky/rivulet-backend.git/internal/app/loger"
	app_http_response "github.com/DanielTitovsky/rivulet-backend.git/internal/app/transport/http/response"
	app_http_utils "github.com/DanielTitovsky/rivulet-backend.git/internal/app/transport/http/utils"
	"github.com/gin-gonic/gin"
)

type GetTrackResponse TrackDTOResponse

func (h *TrackHttpHandler) GetTrack(c *gin.Context) {
	ctx := c.Request.Context()
	log := app_loger.FromContext(ctx)

	responseHandler := app_http_response.NewHTTPResponseHandler(log, c.Writer)

	trackId, err := app_http_utils.GetQueryParamsUUID(c, "id")

	if err != nil {
		responseHandler.ErrorResponse(err, "Invalid id")
		return
	}

	trackDomain, err := h.trackService.GetTrack(ctx, *trackId)

	if err != nil {
		responseHandler.ErrorResponse(err, "Failed to get track")
		return
	}

	responseTrack := GetTrackResponse(trackDTOFromDomain(trackDomain))

	responseHandler.JSONResponse(app_http_response.Response{
		Status: http.StatusOK,
		Data:   responseTrack,
	})
}
