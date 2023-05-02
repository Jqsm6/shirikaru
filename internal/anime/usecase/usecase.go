package usecase

import (
	"context"

	"github.com/rs/zerolog/log"

	"shirikaru/internal/anime"
	"shirikaru/internal/models"
	"shirikaru/pkg/logger"
)

type animeUseCase struct {
	repo anime.Repository
	log  *logger.Logger
}

func NewAnimeUseCase(repo anime.Repository, log *logger.Logger) anime.UseCase {
	return &animeUseCase{repo: repo, log: log}
}

func (aus *animeUseCase) Upload(ctx context.Context, anime *models.Anime) (int, error) {
	id, err := aus.repo.Upload(ctx, anime.ToDB())
	if err != nil {
		log.Err(err).Msg("")
		return id, err
	}

	return id, nil
}

func (aus *animeUseCase) GetAll(ctx context.Context) ([]*models.Anime, error) {
	modelList, err := aus.repo.GetAll(ctx)
	if err != nil {
		log.Err(err).Msg("")
		return nil, err
	}

	return modelList, nil
}

func (aus *animeUseCase) GetByID(ctx context.Context, id int) (*models.Anime, error) {
	model, err := aus.repo.GetByID(ctx, id)
	if err != nil {
		log.Err(err).Msg("")
		return nil, err
	}

	return model, nil
}

func (aus *animeUseCase) SearchByTitle(ctx context.Context, title string) ([]*models.Anime, error) {
	modelList, err := aus.repo.SearchByTitle(ctx, title)
	if err != nil {
		log.Err(err).Msg("")
		return nil, err
	}

	return modelList, nil
}

func (aus *animeUseCase) Delete(ctx context.Context, id int) error {
	err := aus.repo.Delete(ctx, id)
	if err != nil {
		log.Err(err).Msg("")
		return err
	}

	return nil
}
