package auth_service

import (
	"context"

	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
	token_service "github.com/DanielTitovsky/rivulet-backend.git/internal/features/tokens/service"
	user_service "github.com/DanielTitovsky/rivulet-backend.git/internal/features/users/service"
)

type AuthServise struct {
	AuthRepository AuthRepository
	TokenService   token_service.TokenService
	UserService    user_service.UsersServise
}

type AuthRepository interface {
	RegisterUser(ctx context.Context, user domain.User) (domain.User, error)
}

// FindUserById(ctx context.Context, userId uuid.UUID) (*domain.User, error)
// 	FindUserByEmail(ctx context.Context, userEmail string) (*domain.User, error)
// 	SaveOAthAccaounts(ctx context.Context, provideUser domain.ProvideUser, userId uuid.UUID) error
// 	UpdateUser(ctx context.Context, userId domain.User) (*domain.User, error)
// 	DeleteUser(ctx context.Context, userId uuid.UUID) error

func NewUserServise(rep AuthRepository, tokenService token_service.TokenService, userService user_service.UsersServise) *AuthServise {
	return &AuthServise{
		AuthRepository: rep,
		TokenService:   tokenService,
		UserService:    userService,
	}
}
