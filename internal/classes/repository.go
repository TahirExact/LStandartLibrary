package classes

import (
	"context"
	"my-project/internal/models"
)

type Repository interface {
	Create(ctx context.Context, c *models.Class) error
	GetAll(ctx context.Context) ([]*models.Class, error)
	GetByID(ctx context.Context, id int) (*models.Class, error)
}
