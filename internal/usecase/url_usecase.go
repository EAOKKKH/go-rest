package usecase

import (
	"context"
	"go-rest/internal/config"
	"go-rest/internal/models"
	"go-rest/internal/repository"
	"go-rest/internal/utils"
	"strings"

	"github.com/sirupsen/logrus"
)

const (
	PREFIX = "https://"
	WWW    = "www"
)

type IUrlUsecase interface {
	CreateUrl(ctx context.Context, url *models.Url) (*models.UrlResponse, error)
	GetUrl(ctx context.Context, alias string) (string, error)
}

type urlUsecase struct {
	logger        *logrus.Logger
	cfg           *config.Config
	urlRepository repository.IUrlRepository
}

func NewUrlUsecase(logger *logrus.Logger, cfg *config.Config, urlRepository repository.IUrlRepository) IUrlUsecase {
	return &urlUsecase{
		logger:        logger,
		cfg:           cfg,
		urlRepository: urlRepository,
	}
}

func (u *urlUsecase) CreateUrl(ctx context.Context, url *models.Url) (*models.UrlResponse, error) {
	url_parse := url.Url
	url_parse = strings.Replace(url_parse, WWW, "", 1)
	if !strings.Contains(url_parse, PREFIX) {
		url_parse = PREFIX + url_parse
	}
	createUrl := &models.CreateUrl{
		Url:   url_parse,
		Alias: utils.GetAlias(7),
	}
	createdUrl, err := u.urlRepository.SaveUrl(ctx, createUrl)
	if err != nil {
		u.logger.Error(err)
		return nil, err
	}
	return createdUrl, nil
}

func (u *urlUsecase) GetUrl(ctx context.Context, alias string) (string, error) {
	alias, err := u.urlRepository.GetUrl(ctx, alias)
	if err != nil {
		return "", err
	}
	return alias, nil
}
