package tracks_transport_http

import (
	"net/http"

	app_loger "github.com/DanielTitovsky/rivulet-backend.git/internal/app/loger"
	app_http_response "github.com/DanielTitovsky/rivulet-backend.git/internal/app/transport/http/response"
	"github.com/gin-gonic/gin"
)

type GetRandomTrackResponse TrackDTOResponse

func (h *TrackHttpHandler) GetRandomTrack(c *gin.Context) {
	ctx := c.Request.Context()
	log := app_loger.FromContext(ctx)

	responseHandler := app_http_response.NewHTTPResponseHandler(log, c.Writer)

	trackDomain, err := h.trackService.GetRandomTrack(ctx)

	if err != nil {
		responseHandler.ErrorResponse(err, "Failed to get track")
		return
	}

	responseTrack := GetRandomTrackResponse(trackDTOFromDomain(trackDomain))

	responseHandler.JSONResponse(app_http_response.Response{
		Status: http.StatusOK,
		Data:   responseTrack,
	})
}
