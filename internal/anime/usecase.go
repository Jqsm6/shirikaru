package anime

import (
	"context"

	"shirikaru/internal/model"
)

type UseCase interface {
	Upload(ctx context.Context, anime *model.Anime) (int, error)
	GetByID(ctx context.Context, id int) (*model.Anime, error)
	GetByTitle(ctx context.Context, title string) ([]*model.Anime, error)
	Delete(ctx context.Context, id int) error
}
