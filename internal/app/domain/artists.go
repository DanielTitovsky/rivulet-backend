package domain

import (
	"github.com/google/uuid"
)

type Artist struct {
	Id          uuid.UUID
	Name        string
	AvatarUrl   string
	Description string
}
