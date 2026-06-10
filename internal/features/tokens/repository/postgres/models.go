package tokens_postgres_repository

import (
	"time"

	"github.com/google/uuid"
)

type TokenModel struct {
	Id          uuid.UUID  `json:"id"`
	UserId      uuid.UUID  `json:"userId"`
	TokenString string     `json:"tokenString"`
	ExpiresAt   time.Time  `json:"expiresAt"`
	RevokedAt   *time.Time `json:"revokedAt"`
	CreatedAt   time.Time  `json:"createdAt"`
}
