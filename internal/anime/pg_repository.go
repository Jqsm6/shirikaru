package anime

import (
	"context"

	"shirikaru/internal/models"
)

type Repository interface {
	Upload(context.Context, *models.DBAnime) (int, error)
	GetAll(ctx context.Context) ([]*models.Anime, error)
	GetByID(ctx context.Context, id int) (*models.Anime, error)
	SearchByTitle(ctx context.Context, title string) ([]*models.Anime, error)
	Delete(ctx context.Context, id int) error
}
