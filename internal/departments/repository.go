package departments

import (
	"context"
	"my-project/internal/models"
)

type Repository interface {
	Create(ctx context.Context, d *models.Department) error
	GetByID(ctx context.Context, id int) (*models.Department, error)
	GetAll(ctx context.Context) ([]*models.Department, error)
}
