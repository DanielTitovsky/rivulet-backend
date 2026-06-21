package tracks_service

import (
	"context"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
	app_postgres_transaction "github.com/DanielTitovsky/rivulet-backend.git/internal/app/repository/postgres/transaction"
	"github.com/google/uuid"
)

type TrackServise struct {
	TrackRepository        TrackRepository
	TrackStorageRepository TrackStorageRepository
	TransactionManager     app_postgres_transaction.TransactionManager
}

type TrackRepository interface {
	CreateTrack(ctx context.Context, track domain.Track) (uuid.UUID, error)
	GetTrack(ctx context.Context, trackId uuid.UUID) (domain.Track, error)
	GetRandomTrack(ctx context.Context) (domain.Track, error)
	GetTracks(ctx context.Context, trackId domain.TrackFilters) ([]domain.Track, error)
	GetTracksByArtistId(ctx context.Context, artistId uuid.UUID) ([]domain.Track, error)
	GetTracksByPlaylistId(ctx context.Context, playlistId uuid.UUID) ([]domain.Track, error)
	UpdateTrack(ctx context.Context, trackId uuid.UUID, track domain.Track) error
	UpdateTrackArtists(ctx context.Context, trackId uuid.UUID, artistId []uuid.UUID) error
	UpdateTrackGenres(ctx context.Context, trackId uuid.UUID, genreId []uuid.UUID) error
	CreateTrackArtists(ctx context.Context, trackId uuid.UUID, artistIds []uuid.UUID) error
	CreateTrackGenres(ctx context.Context, trackId uuid.UUID, genreIds []uuid.UUID) error
	DeleteTrack(ctx context.Context, trackId uuid.UUID) error
}

type TrackStorageRepository interface {
	GetTrackAudioLink(ctx context.Context, link string) (string, error)
	GetTrackCoverLink(ctx context.Context, link string) (string, error)
}

func NewTrackServise(
	rep TrackRepository,
	transactionManager app_postgres_transaction.TransactionManager,
	TrackStorageRepository TrackStorageRepository,
) *TrackServise {
	return &TrackServise{
		TrackRepository:        rep,
		TransactionManager:     transactionManager,
		TrackStorageRepository: TrackStorageRepository,
	}
}
