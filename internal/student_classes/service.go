package studentclasses

import (
	"context"
	"my-project/internal/classes"
	"my-project/internal/models"
	"my-project/internal/students"
)

type Service struct {
	repo        Repository
	studentRepo students.Repository
	classRepo   classes.Repository
}

func NewService(repo Repository, studentRepo students.Repository, classRepo classes.Repository) *Service {
	return &Service{
		repo:        repo,
		studentRepo: studentRepo,
		classRepo:   classRepo,
	}
}

func (s *Service) AssignClassToStudent(ctx context.Context, studentId, classId int) (*models.StudentClasses, error) {
	student, err := s.studentRepo.GetByID(ctx, studentId)
	if student == nil && err != nil {
		return nil, err
	}
	class, err := s.classRepo.GetByID(ctx, classId)
	if class == nil && err != nil {
		return nil, err
	}
	studentClass := &models.StudentClasses{
		StudentID: studentId,
		ClassID:   classId,
	}
	if err := s.repo.AssignClassToStudent(ctx, studentClass); err != nil {
		return nil, err
	}

	return studentClass, nil
}

func (s *Service) GetStudentWithClasses(ctx context.Context, studentId int) (*models.Student, error) {
	student, err := s.studentRepo.GetByID(ctx, studentId)
	if student == nil && err != nil {
		return nil, err
	}
	classId := make([]int, 0, 20)
	classes := make([]*models.Class, 0, 20)
	studentClasses, err := s.repo.GetAllStudentsWithClasses(ctx)
	if err != nil {
		return nil, err
	}
	for _, studentClass := range studentClasses {
		if studentClass.StudentID == studentId {
			classId = append(classId, studentClass.ClassID)
		}
	}

	for _, id := range classId {
		class, err := s.classRepo.GetByID(ctx, id)
		if err != nil {
			return nil, err
		}
		classes = append(classes, class)
	}

	student.Classes = classes
	return student, nil
}
