package album_transport_http

import (
	"net/http"

	app_loger "github.com/DanielTitovsky/rivulet-backend.git/internal/app/loger"
	app_http_response "github.com/DanielTitovsky/rivulet-backend.git/internal/app/transport/http/response"
	app_http_utils "github.com/DanielTitovsky/rivulet-backend.git/internal/app/transport/http/utils"
	"github.com/gin-gonic/gin"
)

type AlbumResponse AlbumDTOResponse

func (h *AlbumHttpHandler) GetAlbum(c *gin.Context) {
	ctx := c.Request.Context()
	log := app_loger.FromContext(ctx)

	responceHandler := app_http_response.NewHTTPResponseHandler(log, c.Writer)

	albumId, err := app_http_utils.GetQueryParamsUUID(c, "id")

	if err != nil {
		responceHandler.ErrorResponse(err, "Invalid query params")
		return
	}

	albumDomain, err := h.AlbumService.GetAlbum(ctx, *albumId)

	if err != nil {
		responceHandler.ErrorResponse(err, "Failed to get Album")
		return
	}

	albumResponce := AlbumResponse(albumDomain)

	responceHandler.JSONResponse(app_http_response.Response{
		Status: http.StatusOK,
		Data:   albumResponce,
	})
}
