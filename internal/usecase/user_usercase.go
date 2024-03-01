package usecase

import (
	"context"
	"go-rest/internal/config"
	"go-rest/internal/models"
	"go-rest/internal/repository"

	"github.com/sirupsen/logrus"
)

type IAuthUsecase interface {
	Register(ctx context.Context, user *models.User) (*models.User, error)
}

type authUsecase struct {
	logger         *logrus.Logger
	cfg            *config.Config
	authRepository repository.IAuthRepository
}

// Register implements IAuthUsecase.
func (u *authUsecase) Register(ctx context.Context, user *models.User) (*models.User, error) {
	if err := user.HashPassword(); err != nil {
		return nil, err
	}
	user, err := u.authRepository.SaveUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func NewAuthUsecase(logger *logrus.Logger, cfg *config.Config, authRepository repository.IAuthRepository) IAuthUsecase {
	return &authUsecase{logger: logger, cfg: cfg, authRepository: authRepository}
}
