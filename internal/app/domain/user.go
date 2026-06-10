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

var (
	passwordRegex = regexp.MustCompile(`([0-9].*[a-zA-Z].*[^a-zA-Z0-9]|[0-9].*[^a-zA-Z0-9].*[a-zA-Z]|[a-zA-Z].*[0-9].*[^a-zA-Z0-9]|[a-zA-Z].*[^a-zA-Z0-9].*[0-9]|[^a-zA-Z0-9].*[0-9].*[a-zA-Z]|[^a-zA-Z0-9].*[a-zA-Z].*[0-9])`)
	emailRegex    = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
)

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

type UserUpdate struct {
	Email    *string
	Password string
	Name     string
}

func (u *UserUpdate) Validate() error {
	if u.Name == "" {
		return fmt.Errorf("Name cant't be update to null")
	}
	return nil
}

func (u *User) ApplyUpdate(update UserUpdate) error {
	if err := update.Validate(); err != nil {
		return fmt.Errorf("validate user update: %w", err)
	}

	tmp := *u

	tmp.Email = update.Email
	tmp.Name = update.Name
	tmp.Password = &update.Password

	if err := tmp.Validate(); err != nil {
		return fmt.Errorf("validate updated user: %w", err)
	}

	*u = tmp

	return nil
}
