package user

import (
	"context"
	"errors"
	"t-mail/internal/infrastructure/repository"
	"t-mail/internal/pkg/hasher"
	"t-mail/internal/pkg/jwt"
)

type AuthService interface {
	Authenticate(ctx context.Context, email, password string) (string, error)
}

type authService struct {
	repository repository.UserRepository
	hasher     hasher.Manager
	jwtManager jwt.Manager
}

func CreateAuthService(
	repository repository.UserRepository,
	hasher hasher.Manager,
	jwtManager jwt.Manager,
) AuthService {
	return &authService{repository, hasher, jwtManager}
}

func (s *authService) Authenticate(ctx context.Context, email, password string) (string, error) {
	user, err := s.repository.GetByEmail(ctx, email)
	if err != nil {
		return "", errors.New("Bad credentials")
	}
	if !s.hasher.ComparePassword(password, user.Password) {
		return "", errors.New("Bad credentials")
	}

	token, err := s.jwtManager.Generate(user)
	if err != nil {
		return "", errors.New("Bad credentials")
	}

	return token, nil
}
