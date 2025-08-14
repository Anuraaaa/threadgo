package usecase

import (
	"errors"
	"time"

	"github.com/anuraaaa/threadgo/domain"
	"github.com/anuraaaa/threadgo/utils"
)

type AuthUseCase struct{ userRepo domain.UserRepository }

func NewAuthUseCase(r domain.UserRepository) *AuthUseCase { return &AuthUseCase{userRepo: r} }

type AuthRegisterInput struct {
	Name, Email, Password string
}

type AuthLoginInput struct {
	Email, Password string
}

func (uc *AuthUseCase) Register(in AuthRegisterInput) error {
	if in.Email == "" || in.Password == "" {
		return errors.New("email/password required")
	}
	hash, err := utils.HashPassword(in.Password)
	if err != nil {
		return err
	}
	u := &domain.User{Name: in.Name, Email: in.Email, Password: hash}
	return uc.userRepo.Create(u)
}

func (uc *AuthUseCase) Login(in AuthLoginInput) (string, *domain.User, error) {
	u, err := uc.userRepo.GetByEmail(in.Email)
	if err != nil {
		return "", nil, errors.New("user not found")
	}
	if !utils.CheckPassword(u.Password, in.Password) {
		return "", nil, errors.New("invalid credentials")
	}
	token, err := utils.GenerateToken(u.ID, 24*time.Hour)
	return token, u, err
}
