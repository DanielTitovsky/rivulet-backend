package tracks_transport_http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
	app_loger "github.com/DanielTitovsky/rivulet-backend.git/internal/app/loger"
	app_http_request "github.com/DanielTitovsky/rivulet-backend.git/internal/app/transport/http/request"
	app_http_response "github.com/DanielTitovsky/rivulet-backend.git/internal/app/transport/http/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateTrackRequest struct {
	AlbumId         uuid.UUID     `json:"album_id" validate:"required"`
	Name            string        `json:"name" validate:"required,min=1,max=255"`
	Description     string        `json:"description" validate:"max=2000"`
	CoverStorageKey string        `json:"cover_storage_key" validate:"required"`
	DurationSeconds time.Duration `json:"duration_seconds" validate:"required"`
	ArtistIds       []uuid.UUID   `json:"artist_ids" validate:"required,min=1"`
	GenreIds        []uuid.UUID   `json:"genre_ids" validate:"required,min=1"`
	ReleaseDate     time.Time     `json:"release_date" validate:"required"`
	IsExplicit      bool          `json:"is_explicit"`
	IsStreamable    bool          `json:"is_streamable"`
	IsDownloadable  bool          `json:"is_downloadable"`
	AudioStorageKey string        `json:"audio_storage_key" validate:"required"`
	StatusId        int           `json:"status_id" validate:"required"`
}

type CreateTrackResponse TrackDTOResponse

func (h *TrackHttpHandler) CreateTrack(c *gin.Context) {
	ctx := c.Request.Context()
	log := app_loger.FromContext(ctx)

	responseHandler := app_http_response.NewHTTPResponseHandler(log, c.Writer)

	var requestTrack CreateTrackRequest

	if err := app_http_request.DecodeAndValidate(*c.Request, &requestTrack); err != nil {
		fmt.Print("\n")
		fmt.Print(err)
		fmt.Print("\n")
		responseHandler.ErrorResponse(err, "Invalid request")
		return
	}

	trackDomain := domainFromDTO(requestTrack)

	trackDomain, err := h.trackService.CreateTrack(ctx, trackDomain)

	if err != nil {
		responseHandler.ErrorResponse(err, "Failed to create track")
		return
	}

	responseTrack := CreateTrackResponse(trackDTOFromDomain(trackDomain))

	responseHandler.JSONResponse(app_http_response.Response{Status: http.StatusCreated, Data: responseTrack})
}

func domainFromDTO(dto CreateTrackRequest) domain.Track {
	return domain.NewTrackUninitialized(
		dto.Name,
		dto.Description,
		&dto.AlbumId,
		dto.ArtistIds,
		dto.GenreIds,
		dto.CoverStorageKey,
		dto.DurationSeconds,
		dto.ReleaseDate,
		dto.IsExplicit,
		dto.IsStreamable,
		dto.IsDownloadable,
		dto.StatusId,
		dto.AudioStorageKey,
	)
}
