package adduser_usecase

import (
	"errors"

	user_repository "github.com/williamkoller/divine-beast/internal/user/repository"
)

var (
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrInvalidEmail      = errors.New("email is required")
	ErrInvalidAge        = errors.New("age must be at least 18")
)

type AddUserUseCase struct {
	repo user_repository.UserRepository
}

func NewAddUserUseCase(repo user_repository.UserRepository) *AddUserUseCase {
	return &AddUserUseCase{repo: repo}
}

func (uc *AddUserUseCase) Execute(email string, age int) error {
	if email == "" {
		return ErrInvalidEmail
	}
	if age < 18 {
		return ErrInvalidAge
	}

	if _, exists := uc.repo.GetUser(email); exists {
		return ErrUserAlreadyExists
	}

	user := user_repository.User{
		Email: email,
		Age:   age,
	}

	if !uc.repo.AddUser(user) {
		return errors.New("failed to add user")
	}

	return nil
}
