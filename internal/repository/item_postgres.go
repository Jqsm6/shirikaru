package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"shirikaru-rest-api/internal/model"
	"shirikaru-rest-api/pkg/logger"
)

type ItemPostgres struct {
	db  *sqlx.DB
	log *logger.Logger
}

func NewItemPostgres(db *sqlx.DB, log *logger.Logger) *ItemPostgres {
	return &ItemPostgres{
		db:  db,
		log: log,
	}
}

func (repo *ItemPostgres) Upload(ctx context.Context, a *model.DBAnime) (int, error) {
	query := "INSERT INTO anime (title, alternative_title, description, production_status, picture, episode) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id_"

	var id int
	err := repo.db.QueryRowContext(ctx, query, a.Title, a.AlternativeTitle, a.Description,
		a.ProductionStatus, a.Picture, a.Episode).Scan(&id)
	if err != nil {
		log.Err(err).Msg("repository level")
		return 0, err
	}

	return id, nil
}

func (repo *ItemPostgres) Get(ctx context.Context, id int) (*model.Anime, error) {
	query := "SELECT * FROM anime WHERE id_ = $1"

	var a model.Anime
	err := repo.db.QueryRowContext(ctx, query, id).Scan(&a.ID, &a.Title, &a.AlternativeTitle, &a.Description, &a.ProductionStatus,
		&a.Picture, &a.Episode)
	if err != nil {
		log.Err(err).Msg("repository level")
		return nil, err
	}

	return &a, nil
}
