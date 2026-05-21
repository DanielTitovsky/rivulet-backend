package domain

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id           uuid.UUID
	Email        *string
	Name         string
	HashPassword *string
	UpdatedAt    time.Time
	CreatedAt    time.Time
}

type ProviderType string

const (
	ProviderGoogle ProviderType = "google"
	ProviderYandex ProviderType = "yandex"
	ProviderGit    ProviderType = "git"
)

type ProvideUser struct {
	Provider          ProviderType
	ProviderUserId    string
	ProviderUserEmail string
	EmailVerified     bool
	Name              string
	GivenName         string
	FamilyName        string
}

func (u *User) CheckPassword(password string, hash string) bool {

	if hash == "" {
		hash = *u.HashPassword
	}

	fmt.Print(password)
	fmt.Print(hash)

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}

func HashUserPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(bytes)
}
