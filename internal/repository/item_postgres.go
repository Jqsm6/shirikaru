package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"

	"shirikaru-rest-api/internal/logger"
	"shirikaru-rest-api/internal/model"
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
	query := "INSERT INTO anime_list (title, alternative_title, description, production_status, picture, episode) VALUES ($1, $2, $3, $4, $5, $6) RETURNING anime_id"

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
	query := "SELECT * FROM anime_list WHERE anime_id = $1"

	var a model.Anime
	err := repo.db.QueryRowContext(ctx, query, id).Scan(&a.ID, &a.Title, &a.AlternativeTitle, &a.Description, &a.ProductionStatus,
		&a.Picture, &a.Episode)
	if err != nil {
		log.Err(err).Msg("repository level")
		return nil, err
	}

	return &a, nil
}
