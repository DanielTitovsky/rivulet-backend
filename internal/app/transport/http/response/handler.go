package app_http_response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	app_loger "github.com/DanielTitovsky/rivulet-backend.git/internal/app/loger"
	"go.uber.org/zap"
)

type HTTPResponseHandler struct {
	log *app_loger.Logger
	rw  http.ResponseWriter
}

type Response struct {
	Status  int
	Data    any
	Headers map[string]string
	Cookie  []*http.Cookie
}

func NewHTTPResponseHandler(logger *app_loger.Logger, rw http.ResponseWriter) *HTTPResponseHandler {
	return &HTTPResponseHandler{
		log: logger,
		rw:  rw,
	}
}

func (h *HTTPResponseHandler) JSONResponse(res Response) {
	h.rw.Header().Set("Content-Type", "application/json")

	for key, value := range res.Headers {
		h.rw.Header().Set(key, value)
	}

	for _, value := range res.Cookie {
		http.SetCookie(h.rw, value)
	}

	if res.Status == 0 {
		res.Status = http.StatusOK
	}

	h.rw.WriteHeader(res.Status)
	json.NewEncoder(h.rw).Encode(res.Data)

	if res.Status >= 400 {
		h.log.Error("writed HTTP response", zap.Any("error", res.Data))
	} else {
		h.log.Debug("writed HTTP response", zap.String("status", strconv.Itoa(res.Status)))
	}
}

func (h *HTTPResponseHandler) PanicResponse(p any, msg string) {
	statusCode := http.StatusInternalServerError

	err := fmt.Errorf("unexpected panic: %v", p)

	h.JSONResponse(Response{
		Status: statusCode,
		Data: struct {
			Err         error
			Description string
		}{
			Err:         err,
			Description: msg,
		},
	})
}

func (h *HTTPResponseHandler) ErrorResponse(err error, msg string) {
	statusCode := http.StatusBadRequest

	h.JSONResponse(Response{
		Status: statusCode,
		Data: struct {
			Error       error  `json:"error"`
			Description string `json:"description,omitempty"`
		}{
			Error:       err,
			Description: msg,
		},
	})
}
