package domain

import (
	"time"

	"github.com/google/uuid"
)

type Album struct {
	Id           uuid.UUID
	Artists      []Artist
	Name         string
	CoverUrl     string
	Description  string
	Release_date time.Time
}
