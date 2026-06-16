package app_http_utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetQueryParamsUUID(c *gin.Context, key string) (*uuid.UUID, error) {
	param := c.Param(key)

	if param == "" {
		return nil, nil
	}

	val, err := uuid.Parse(param)

	if err != nil {
		return nil, fmt.Errorf(
			"param='%s' by key='%s' not a valid uuid: %v: %w",
			param, key, val, err,
		)
	}

	return &val, nil
}

func GetQueryаFilter[T any](c *gin.Context) (T, error) {
	var filter T

	if err := c.ShouldBindQuery(&filter); err != nil {
		var nilValue T
		return nilValue, fmt.Errorf("Filter not a valid: %v: %w", filter, err)
	}

	return filter, nil
}
