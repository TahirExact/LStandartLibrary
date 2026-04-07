package departments

import (
	"context"
	"fmt"
	"my-project/internal/models"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreteDepartment(ctx context.Context, name string, id int) (*models.Department, error) {
	if name == "" {
		return nil, fmt.Errorf("name can't be empty")
	}

	if id == 0 {
		return nil, fmt.Errorf("id can't be empty")
	}

	newDepartment := &models.Department{
		ID:   id,
		Name: name,
	}

	if err := s.repo.Create(ctx, newDepartment); err != nil {
		return nil, err
	}
	return newDepartment, nil
}

func (s *Service) GetDepartment(ctx context.Context, id int) (*models.Department, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *Service) GetAll(ctx context.Context) ([]*models.Department, error) {
	return s.repo.GetAll(ctx)
}
