package repository

import (
	"context"

	"github.com/jmoiron/sqlx"

	"shirikaru/internal/anime"
	"shirikaru/internal/models"
)

type animeRepo struct {
	db *sqlx.DB
}

func NewAnimeRepository(db *sqlx.DB) anime.Repository {
	return &animeRepo{db: db}
}

func (r *animeRepo) Upload(ctx context.Context, a *models.DBAnime) (int, error) {
	var id int
	err := r.db.QueryRowContext(ctx, uploadAnime, a.AnimeID, a.Title, a.AlternativeTitle, a.Description,
		a.ProductionStatus, a.Picture, a.Episode).Scan(&id)
	if err != nil {
		return id, err
	}

	return id, nil
}

func (r *animeRepo) GetAll(ctx context.Context) ([]*models.Anime, error) {
	var animeList []*models.Anime

	err := r.db.SelectContext(ctx, &animeList, getAnimeAll)
	if err != nil {
		return nil, err
	}

	return animeList, nil
}

func (r *animeRepo) GetByID(ctx context.Context, id int) (*models.Anime, error) {
	var a models.Anime
	err := r.db.QueryRowContext(ctx, getAnimeByID, id).Scan(&a.AnimeID, &a.Title, &a.AlternativeTitle, &a.Description, &a.ProductionStatus,
		&a.Picture, &a.Episode)
	if err != nil {
		return nil, err
	}

	return &a, nil
}

func (r *animeRepo) SearchByTitle(ctx context.Context, title string) ([]*models.Anime, error) {
	var animeList []*models.Anime

	err := r.db.SelectContext(ctx, &animeList, searchAnimeByTitle, "%"+title+"%")
	if err != nil {
		return nil, err
	}

	return animeList, nil
}

func (r *animeRepo) Delete(ctx context.Context, id int) error {
	_, err := r.db.ExecContext(ctx, deleteAnime, id)
	if err != nil {
		return err
	}

	return nil
}
