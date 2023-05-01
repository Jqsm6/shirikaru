package repository

import (
	"context"

	"github.com/jmoiron/sqlx"

	"shirikaru/internal/anime"
	"shirikaru/internal/model"
)

type animeRepo struct {
	db *sqlx.DB
}

func NewAnimeRepository(db *sqlx.DB) anime.Repository {
	return &animeRepo{db: db}
}

func (r *animeRepo) Upload(ctx context.Context, a *model.DBAnime) (int, error) {
	var id int
	err := r.db.QueryRowContext(ctx, uploadAnime, a.Title, a.AlternativeTitle, a.Description,
		a.ProductionStatus, a.Picture, a.Episode).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *animeRepo) GetAll(ctx context.Context) ([]*model.Anime, error) {
	var animeList []*model.Anime

	err := r.db.SelectContext(ctx, &animeList, getAnimeAll)
	if err != nil {
		return nil, err
	}

	return animeList, nil
}

func (r *animeRepo) GetByID(ctx context.Context, id int) (*model.Anime, error) {
	var a model.Anime
	err := r.db.QueryRowContext(ctx, getAnimeByID, id).Scan(&a.ID, &a.Title, &a.AlternativeTitle, &a.Description, &a.ProductionStatus,
		&a.Picture, &a.Episode)
	if err != nil {
		return nil, err
	}

	return &a, nil
}

func (r *animeRepo) SearchByTitle(ctx context.Context, title string) ([]*model.Anime, error) {
	var animeList []*model.Anime

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
