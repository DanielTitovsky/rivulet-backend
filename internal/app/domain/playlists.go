package domain

import (
	"github.com/google/uuid"
)

type Playlsit struct {
	Id          uuid.UUID
	UserId      *uuid.UUID
	Name        string
	Description string
	CoverUrl    string
	Tracks      []Track
	Visibility  bool
}
