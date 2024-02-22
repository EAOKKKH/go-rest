package repository

import (
	"context"
	"go-rest/internal/models"

	"github.com/jmoiron/sqlx"
)

type IUrlRepository interface {
	SaveUrl(ctx context.Context, url *models.CreateUrl) (*models.UrlResponse, error)
	GetUrl(ctx context.Context, alias string) (string, error)
}

type urlRepository struct {
	db *sqlx.DB
}

func NewUrlRepository(db *sqlx.DB) IUrlRepository {
	return &urlRepository{db: db}
}

func (r *urlRepository) SaveUrl(ctx context.Context, url *models.CreateUrl) (*models.UrlResponse, error) {
	u := &models.UrlResponse{}
	if err := r.db.QueryRowxContext(ctx, `INSERT INTO urls (alias, url) VALUES ($1, $2) RETURNING *`,
		url.Alias, url.Url).StructScan(u); err != nil {
		return nil, err
	}
	return u, nil
}

func (r *urlRepository) GetUrl(ctx context.Context, alias string) (string, error) {
	var url string
	if err := r.db.GetContext(ctx, &url, `SELECT url from urls WHERE alias = $1`, alias); err != nil {
		return url, err
	}
	return url, nil
}
