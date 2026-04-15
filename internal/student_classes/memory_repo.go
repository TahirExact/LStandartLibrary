package studentclasses

import (
	"context"
	"my-project/internal/models"
	"sync"
)

type MemoryRepo struct {
	mu             sync.RWMutex
	studentClasses map[int]*models.StudentClasses
	nextID         int
}

func NewMemoryRepo() *MemoryRepo {
	return &MemoryRepo{
		mu:             sync.RWMutex{},
		studentClasses: make(map[int]*models.StudentClasses),
		nextID:         1,
	}
}

func (m *MemoryRepo) AssignClassToStudent(ctx context.Context, sc *models.StudentClasses) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	sc.ID = m.nextID
	m.studentClasses[sc.ID] = sc
	m.nextID++
	return nil
}

func (m *MemoryRepo) GetAllStudentsWithClasses(ctx context.Context) ([]*models.StudentClasses, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	result := make([]*models.StudentClasses, 0, len(m.studentClasses))

	for _, studentClass := range m.studentClasses {
		result = append(result, studentClass)
	}

	return result, nil
}
