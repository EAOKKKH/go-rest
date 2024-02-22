package postgres

import (
	"go-rest/internal/config"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func NewPgConn(cfg *config.Config) (*sqlx.DB, error) {
	dsn := cfg.GetDbDsn()
	dbConn, err := sqlx.Connect(cfg.Postgres.PgDriver, dsn)
	if err != nil {
		return nil, err
	}
	if err = dbConn.Ping(); err != nil {
		return nil, err
	}
	return dbConn, nil
}
