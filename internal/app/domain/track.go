package domain

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Track struct {
	Id              uuid.UUID
	Name            string
	Description     string
	AlbumId         *uuid.UUID
	ArtistIds       []uuid.UUID
	Artists         []Artist
	GenreIds        []uuid.UUID
	Genres          []string
	CoverStorageKey string
	DurationSecond  time.Duration
	ReleaseDate     time.Time
	IsExplicit      bool
	IsStreamable    bool
	IsDownloadable  bool
	StatusId        int
	Status          string
	AudioStorageKey string
}

type TrackUpdate struct {
	Name            string
	Description     string
	AlbumId         *uuid.UUID
	CoverStorageKey string
	DurationSecond  time.Duration
	ReleaseDate     time.Time
	IsExplicit      bool
	ArtistIds       []uuid.UUID
	GenreIds        []uuid.UUID
	IsStreamable    bool
	IsDownloadable  bool
	StatusId        int
	AudioStorageKey string
}

type TrackFilters struct {
	Limit  int `json:"limit" form:"limit"`
	Offset int `json:"offset" form:"offset"`
}

func (t *Track) Validate() error {
	return nil
}

func (t *TrackUpdate) Validate() error {
	return nil
}

func (t *Track) ApplyUpdate(update TrackUpdate) error {
	if err := update.Validate(); err != nil {
		return fmt.Errorf("validate user update: %w", err)
	}

	tmp := *t

	tmp.Name = update.Name
	tmp.Description = update.Description

	tmp.AlbumId = update.AlbumId
	tmp.ArtistIds = update.ArtistIds
	tmp.GenreIds = update.GenreIds
	tmp.StatusId = update.StatusId

	tmp.DurationSecond = update.DurationSecond
	tmp.ReleaseDate = update.ReleaseDate

	tmp.IsExplicit = update.IsExplicit
	tmp.IsStreamable = update.IsStreamable
	tmp.IsDownloadable = update.IsDownloadable

	tmp.CoverStorageKey = update.CoverStorageKey
	tmp.AudioStorageKey = update.AudioStorageKey

	if err := tmp.Validate(); err != nil {
		return fmt.Errorf("validate updated track: %w", err)
	}

	*t = tmp

	return nil
}

func NewTrackUninitialized(
	name string,
	description string,
	albumId *uuid.UUID,
	artistIds []uuid.UUID,
	genreIds []uuid.UUID,
	coverStorageKey string,
	durationSecond time.Duration,
	releaseDate time.Time,
	isExplicit bool,
	isStreamable bool,
	isDownloadable bool,
	statusId int,
	audioStorageKey string,
) Track {
	return Track{
		Id:              UninitializedId,
		Name:            name,
		Description:     description,
		AlbumId:         albumId,
		ArtistIds:       artistIds,
		GenreIds:        genreIds,
		CoverStorageKey: coverStorageKey,
		DurationSecond:  durationSecond,
		ReleaseDate:     releaseDate,
		IsExplicit:      isExplicit,
		IsStreamable:    isStreamable,
		IsDownloadable:  isDownloadable,
		StatusId:        statusId,
		AudioStorageKey: audioStorageKey,
	}
}
