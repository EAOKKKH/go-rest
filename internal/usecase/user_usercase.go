package usecase

import (
	"context"
	"go-rest/internal/config"
	"go-rest/internal/models"
	"go-rest/internal/repository"
	"go-rest/internal/utils"

	"github.com/sirupsen/logrus"
)

type IAuthUsecase interface {
	Register(ctx context.Context, user *models.User) (*models.User, error)
	Login(ctx context.Context, user *models.User) (*models.Token, error)
}

type authUsecase struct {
	logger         *logrus.Logger
	cfg            *config.Config
	authRepository repository.IAuthRepository
}

// Login implements IAuthUsecase.
func (u *authUsecase) Login(ctx context.Context, user *models.User) (*models.Token, error) {
	foundUser, err := u.authRepository.FindUserByLogin(ctx, user.Login)
	if err != nil {
		return nil, err
	}
	if err = foundUser.ComparePasswords(user.Password); err != nil {
		return nil, err
	}
	token, err := utils.GenerateJwtToken(foundUser.Login, u.cfg)
	if err != nil {
		return nil, err
	}
	return &models.Token{
		Token: token,
	}, nil
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
	user.SanitizePassword()
	return user, nil
}

func NewAuthUsecase(logger *logrus.Logger, cfg *config.Config, authRepository repository.IAuthRepository) IAuthUsecase {
	return &authUsecase{logger: logger, cfg: cfg, authRepository: authRepository}
}
