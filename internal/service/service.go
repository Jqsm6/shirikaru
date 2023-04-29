package service

import (
	"context"

	"shirikaru-rest-api/internal/model"
	"shirikaru-rest-api/internal/repository"
)

type item interface {
	Upload(ctx context.Context, anime *model.Anime) (int, error)
	Get(ctx context.Context, id int) (*model.Anime, error)
}

type Service struct {
	item
}

func NewServices(repo *repository.Repository) *Service {
	return &Service{
		item: NewItemService(repo),
	}
}
