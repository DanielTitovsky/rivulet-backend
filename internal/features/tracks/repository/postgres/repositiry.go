package tracks_postgres_repository

import app_postgres_pool "github.com/DanielTitovsky/rivulet-backend.git/internal/app/repository/postgres/pool"

type TrackRepository struct {
	pool app_postgres_pool.Pool
}

func NewTrackRepository(pool app_postgres_pool.Pool) *TrackRepository {
	return &TrackRepository{
		pool: pool,
	}
}
