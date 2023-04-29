package repository

import (
	"context"

	"shirikaru-rest-api/pkg/logger"

	"github.com/jmoiron/sqlx"
	"shirikaru-rest-api/internal/model"
)

type itemRepo interface {
	Upload(context.Context, *model.DBAnime) (int, error)
	Get(ctx context.Context, id int) (*model.Anime, error)
}

type Repository struct {
	itemRepo
}

func NewRepository(db *sqlx.DB, log *logger.Logger) *Repository {
	return &Repository{
		itemRepo: NewItemPostgres(db, log),
	}
}
