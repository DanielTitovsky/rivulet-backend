package auth_service

import (
	"context"
	"time"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
	app_postgres_transaction "github.com/DanielTitovsky/rivulet-backend.git/internal/app/repository/postgres/transaction"
	"github.com/google/uuid"
)

type UserService interface {
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
	GetUser(ctx context.Context, userId uuid.UUID) (domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (domain.User, error)
	GetOrCreateOAuthUser(ctx context.Context, email string, name string) (domain.User, error)
}

type TokenService interface {
	GenerateTokens(
		accessExpires time.Duration,
		refreshExpires time.Duration,
		user domain.User,
	) (*domain.Token, *domain.Token, error)
	SaveRefreshToken(ctx context.Context, token domain.Token) (domain.Token, error)
	ValidateToken(tokenString string, tokenType string) (*domain.TokenClaims, error)
	GetRefreshToken(ctx context.Context, rawToken string) (domain.Token, error)
	RemoveToken(ctx context.Context, tokenId uuid.UUID) error
}

type AuthServise struct {
	UserService        UserService
	TokenService       TokenService
	AuthRepository     AuthRepository
	TransactionManager app_postgres_transaction.TransactionManager
}

type AuthRepository interface {
	GetOAuthAccountUserId(ctx context.Context, provider domain.ProviderType, providerUserId string) (uuid.UUID, error)
	CreateOAuthAccount(ctx context.Context, userId uuid.UUID, oauthUser domain.OAuthUser) error
}

func NewAuthServise(
	userService UserService,
	tokenService TokenService,
	transactionManager app_postgres_transaction.TransactionManager,
	authRepository AuthRepository,
) *AuthServise {
	return &AuthServise{
		UserService:        userService,
		TokenService:       tokenService,
		AuthRepository:     authRepository,
		TransactionManager: transactionManager,
	}
}
