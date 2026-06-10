package users_transport_http

import (
	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
	"github.com/google/uuid"
)

type UserDTOResponse struct {
	Id    uuid.UUID `json:"id" validate:"required"`
	Name  string    `json:"name" validate:"required"`
	Email string    `json:"email"`
}

func userDTOFromDomain(user domain.User) UserDTOResponse {
	return UserDTOResponse{
		Id:    user.Id,
		Name:  user.Name,
		Email: *user.Email,
	}
}
