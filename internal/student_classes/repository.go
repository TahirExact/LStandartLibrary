package studentclasses

import (
	"context"
	"my-project/internal/models"
)

type Repository interface {
	AssignClassToStudent(ctx context.Context, sc *models.StudentClasses) error
	GetAllStudentsWithClasses(ctx context.Context) ([]*models.StudentClasses, error)
}
