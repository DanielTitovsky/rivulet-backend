package app_http_utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetQueryParamsId(r *gin.Context, key string) (*uuid.UUID, error) {
	param := r.Param(key)

	if param == "" {
		return nil, nil
	}

	val, err := uuid.Parse(param)

	if err != nil {
		return nil, fmt.Errorf(
			"param='%s' by key='%s' not a valid integer: %v: %w",
			param, key, val, err,
		)
	}

	return &val, nil
}
