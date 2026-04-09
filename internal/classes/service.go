package classes

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

func (s *Service) CreateClass(ctx context.Context, name string) (*models.Class, error) {
	if name == "" {
		return nil, fmt.Errorf("name cannot be empty")
	}
	newClass := &models.Class{
		Name: name,
	}

	if err := s.repo.Create(ctx, newClass); err != nil {
		return nil, err
	}
	return newClass, nil
}

func (s *Service) GetAllClasses(ctx context.Context) ([]*models.Class, error) {
	return s.repo.GetAll(ctx)
}

func (s *Service) GetByID(ctx context.Context, id int) (*models.Class, error) {
	return s.repo.GetByID(ctx, id)
}
