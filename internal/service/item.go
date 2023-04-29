package service

import (
	"context"

	"shirikaru-rest-api/internal/model"
	"shirikaru-rest-api/internal/repository"
)

type ItemService struct {
	repo *repository.Repository
}

func NewItemService(repo *repository.Repository) *ItemService {
	return &ItemService{
		repo: repo,
	}
}

func (is *ItemService) Upload(ctx context.Context, anime *model.Anime) (int, error) {
	id, err := is.repo.Upload(ctx, anime.ToDB())
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (is *ItemService) Get(ctx context.Context, id int) (*model.Anime, error) {
	anime, err := is.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return anime, nil
}
