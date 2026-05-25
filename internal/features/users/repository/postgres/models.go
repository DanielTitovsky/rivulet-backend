package users_postgres_repository

import (
	"time"

	"github.com/google/uuid"
)

type UserModel struct {
	Id        uuid.UUID
	Email     *string
	Name      string
	Password  *string
	UpdatedAt time.Time
	CreatedAt time.Time
}
