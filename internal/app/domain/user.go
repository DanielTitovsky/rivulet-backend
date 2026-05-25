package domain

import (
	"fmt"
	"regexp"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        uuid.UUID
	Email     *string
	Name      string
	Password  *string
	UpdatedAt time.Time
	CreatedAt time.Time
}

type ProviderType string

const (
	ProviderGoogle ProviderType = "google"
	ProviderYandex ProviderType = "yandex"
	ProviderGit    ProviderType = "git"
)

var (
	passwordRegex = regexp.MustCompile(`([0-9].*[a-zA-Z].*[^a-zA-Z0-9]|[0-9].*[^a-zA-Z0-9].*[a-zA-Z]|[a-zA-Z].*[0-9].*[^a-zA-Z0-9]|[a-zA-Z].*[^a-zA-Z0-9].*[0-9]|[^a-zA-Z0-9].*[0-9].*[a-zA-Z]|[^a-zA-Z0-9].*[a-zA-Z].*[0-9])`)
	emailRegex    = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
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

func (u *User) PasswordComparison(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}

func (u *User) HashUserPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(bytes)
}

func NewUserUninitialized(email *string, name string, password *string) User {
	return User{
		Id:        UninitializedId,
		Email:     email,
		Name:      name,
		Password:  password,
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}
}

func (u *User) Validate() error {

	passwordLenght := len([]rune(*u.Password))

	if passwordLenght < 10 {
		return fmt.Errorf("invalid `Pasword` len: %d", passwordLenght)
	}

	if match := passwordRegex.MatchString(*u.Password); !match {
		return fmt.Errorf("invalid `Pasword` len")
	}

	nameLenght := len([]rune(u.Name))

	if nameLenght < 3 || nameLenght > 25 {
		return fmt.Errorf("invalid `Name` len: %d", nameLenght)
	}

	if match := emailRegex.MatchString(*u.Email); !match {
		return fmt.Errorf("invalid `Email`")
	}

	return nil
}
