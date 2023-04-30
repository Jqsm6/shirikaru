package anime

import (
	"context"

	"shirikaru/internal/model"
)

type Repository interface {
	Upload(context.Context, *model.DBAnime) (int, error)
	GetByID(ctx context.Context, id int) (*model.Anime, error)
	Delete(ctx context.Context, id int) error
}
