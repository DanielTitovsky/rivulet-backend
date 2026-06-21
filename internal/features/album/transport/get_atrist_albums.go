package album_transport_http

import (
	"net/http"

	app_loger "github.com/DanielTitovsky/rivulet-backend.git/internal/app/loger"
	app_http_response "github.com/DanielTitovsky/rivulet-backend.git/internal/app/transport/http/response"
	app_http_utils "github.com/DanielTitovsky/rivulet-backend.git/internal/app/transport/http/utils"
	"github.com/gin-gonic/gin"
)

type AlbomsArtistResponce []AlbumDTOResponse

func (h *AlbumHttpHandler) GetArtistAlbums(c *gin.Context) {
	ctx := c.Request.Context()
	log := app_loger.FromContext(ctx)

	responseHandler := app_http_response.NewHTTPResponseHandler(log, c.Writer)

	artistId, err := app_http_utils.GetQueryParamsUUID(c, "id")

	if err != nil {
		responseHandler.ErrorResponse(err, "Invalid filter params")
		return
	}

	albumsDomain, err := h.GetArtistAbums(ctx, *artistId)

	if err != nil {
		responseHandler.ErrorResponse(err, "Failed to get atrist albums")
		return
	}

	albumsResponce := AlbomsArtistResponce(albumsDTOFromDomain(albumsDomain))

	responseHandler.JSONResponse(app_http_response.Response{
		Status: http.StatusOK,
		Data:   albumsResponce,
	})
}
