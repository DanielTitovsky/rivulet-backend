package domain

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

// TODO: добавить type чтобы понимать где какой токен
type Token struct {
	Id          uuid.UUID  `json:"id"`
	UserId      uuid.UUID  `json:"userId"`
	TokenString string     `json:"tokenString"`
	ExpiresAt   time.Time  `json:"expiresAt"`
	RevokedAt   *time.Time `json:"revokedAt"`
	CreatedAt   time.Time  `json:"createdAt"`
}

type TokenClaims struct {
	Id        uuid.UUID `json:"id"`
	UserId    uuid.UUID `json:"userId"`
	UserEmail string    `json:"userEmail"`
	jwt.StandardClaims
}

func NewTokenUninitialized(id uuid.UUID, userId uuid.UUID, expiresAt time.Time, createdAt time.Time) *Token {
	return &Token{
		Id:          id,
		UserId:      userId,
		TokenString: UninitializedTokenString,
		ExpiresAt:   expiresAt,
		RevokedAt:   nil,
		CreatedAt:   createdAt,
	}
}
