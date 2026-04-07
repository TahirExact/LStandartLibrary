package students

import (
	"context"
	"fmt"
	"my-project/internal/models"
	"sync"
)

type MemoryRepo struct {
	mu       sync.RWMutex
	students map[int]*models.Student
	nextID   int
}

func NewMemoryRepo() *MemoryRepo {
	return &MemoryRepo{
		mu:       sync.RWMutex{},
		students: make(map[int]*models.Student),
		nextID:   1,
	}
}

func (m *MemoryRepo) Create(ctx context.Context, s *models.Student) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	s.ID = m.nextID
	m.students[s.ID] = s

	m.nextID++
	return nil
}

func (m *MemoryRepo) GetByID(ctx context.Context, id int) (*models.Student, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	student, ok := m.students[id]

	if !ok {
		return nil, fmt.Errorf("there is no such student with id:%d", id)
	}

	return student, nil
}

func (m *MemoryRepo) GetAll(ctx context.Context) ([]*models.Student, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	result := make([]*models.Student, 0, len(m.students))

	for _, student := range m.students {
		result = append(result, student)
	}

	return result, nil
}
