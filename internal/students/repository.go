package students

import (
	"context"
	"my-project/internal/models"
)

type Repository interface {
	Create(ctx context.Context, s *models.Student) error
	GetByID(ctx context.Context, id int) (*models.Student, error)
	GetAll(ctx context.Context) ([]*models.Student, error)
}
