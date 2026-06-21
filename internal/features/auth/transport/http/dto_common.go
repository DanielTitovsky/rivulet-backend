package auth_transport_http

import (
	"github.com/DanielTitovsky/rivulet-backend.git/internal/app/domain"
	"github.com/google/uuid"
)

type UserDTOResponse struct {
	Id    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}

func userDTOFromDomain(user domain.User) UserDTOResponse {
	email := ""

	if user.Email != nil {
		email = *user.Email
	}

	return UserDTOResponse{
		Id:    user.Id,
		Name:  user.Name,
		Email: email,
	}
}
