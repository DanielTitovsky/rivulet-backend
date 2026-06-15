package tracks_transport_http

import (
	"net/http"
	"time"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
	app_loger "github.com/DanielTitovsky/rivulet-backend.git/internal/app/loger"
	app_http_request "github.com/DanielTitovsky/rivulet-backend.git/internal/app/transport/http/request"
	app_http_response "github.com/DanielTitovsky/rivulet-backend.git/internal/app/transport/http/response"
	app_http_utils "github.com/DanielTitovsky/rivulet-backend.git/internal/app/transport/http/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UpdateTrackRequest struct {
	AlbumId         uuid.UUID     `json:"album_id"`
	Name            string        `json:"name"`
	Description     string        `json:"description"`
	CoverStorageKey string        `json:"cover_storage_key"`
	DurationSeconds time.Duration `json:"duration_seconds"`
	ArtistIds       []uuid.UUID   `json:"artist_ids"`
	GenreIds        []uuid.UUID   `json:"genre_ids"`
	ReleaseDate     time.Time     `json:"release_date"`
	IsExplicit      bool          `json:"is_explicit"`
	IsStreamable    bool          `json:"is_streamable"`
	IsDownloadable  bool          `json:"is_downloadable"`
	AudioStorageKey string        `json:"audio_storage_key"`
	StatusId        int           `json:"status_id"`
}

type UpdateTrackResponse TrackDTOResponse

func (h *TrackHttpHandler) UpdateTrack(c *gin.Context) {
	ctx := c.Request.Context()
	log := app_loger.FromContext(ctx)

	responseHandler := app_http_response.NewHTTPResponseHandler(log, c.Writer)

	trackId, err := app_http_utils.GetQueryParamsUUID(c, "id")

	if err != nil {
		responseHandler.ErrorResponse(err, "Invalid parametrs")
		return
	}

	var reqiestTrack UpdateTrackRequest

	if err := app_http_request.DecodeAndValidate(*c.Request, &reqiestTrack); err != nil {
		responseHandler.ErrorResponse(err, "Invalid request")
		return
	}

	updatedTrack := trackUpdateFromRequest(reqiestTrack)

	trackDomain, err := h.trackService.UpdateTrack(ctx, *trackId, updatedTrack)

	if err != nil {
		responseHandler.ErrorResponse(err, "Invalid track Id")
		return
	}

	responseTrack := UpdateTrackResponse(trackDTOFromDomain(trackDomain))

	responseHandler.JSONResponse(app_http_response.Response{
		Status: http.StatusOK,
		Data:   responseTrack,
	})
}

func trackUpdateFromRequest(request UpdateTrackRequest) domain.TrackUpdate {
	return domain.TrackUpdate{
		Name:            request.Name,
		Description:     request.Description,
		AlbumId:         &request.AlbumId,
		CoverStorageKey: request.CoverStorageKey,
		DurationSecond:  request.DurationSeconds,
		ReleaseDate:     request.ReleaseDate,
		IsExplicit:      request.IsExplicit,
		ArtistIds:       request.ArtistIds,
		GenreIds:        request.GenreIds,
		IsStreamable:    request.IsStreamable,
		IsDownloadable:  request.IsDownloadable,
		StatusId:        request.StatusId,
		AudioStorageKey: request.AudioStorageKey,
	}
}
