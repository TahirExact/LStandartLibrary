package students

import (
	"context"
	"fmt"
	"my-project/internal/departments"
	"my-project/internal/models"
)

type Service struct {
	repo    Repository
	depRepo departments.Repository
}

func NewService(repo Repository, depRepo departments.Repository) *Service {
	return &Service{
		repo:    repo,
		depRepo: depRepo,
	}
}

func (s *Service) CreateStudent(ctx context.Context, name string, deptID int) (*models.Student, error) {
	if name == "" {
		return nil, fmt.Errorf("name cannot be empty")
	}
	if deptID == 0 {
		return nil, fmt.Errorf("department id cannot be empty")
	}

	department, err := s.depRepo.GetByID(ctx, deptID)
	if department == nil && err != nil {
		return nil, err
	}

	newStudent := &models.Student{
		Name:         name,
		DepartmentID: deptID,
	}

	if err := s.repo.Create(ctx, newStudent); err != nil {
		return nil, err
	}
	return newStudent, nil
}

func (s *Service) GetStudent(ctx context.Context, id int) (*models.Student, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *Service) GetStudents(ctx context.Context) ([]*models.Student, error) {
	return s.repo.GetAll(ctx)
}

