package usecase

import (
	"context"

	"github.com/rs/zerolog/log"

	"shirikaru/internal/anime"
	"shirikaru/internal/model"
	"shirikaru/pkg/logger"
)

type animeUseCase struct {
	repo anime.Repository
	log  *logger.Logger
}

func NewAnimeUseCase(repo anime.Repository, log *logger.Logger) anime.UseCase {
	return &animeUseCase{repo: repo, log: log}
}

func (aus *animeUseCase) Upload(ctx context.Context, anime *model.Anime) (int, error) {
	id, err := aus.repo.Upload(ctx, anime.ToDB())
	if err != nil {
		log.Err(err).Msg("")
		return 0, err
	}

	return id, nil
}

func (aus *animeUseCase) GetByID(ctx context.Context, id int) (*model.Anime, error) {
	model, err := aus.repo.GetByID(ctx, id)
	if err != nil {
		log.Err(err).Msg("")
		return nil, err
	}

	return model, nil
}

func (aus *animeUseCase) GetByTitle(ctx context.Context, title string) ([]*model.Anime, error) {
	modelList, err := aus.repo.GetByTitle(ctx, title)
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
