package departments

import (
	"context"
	"fmt"
	"my-project/internal/models"
	"sync"
)

type MemoryRepo struct {
	mu          sync.RWMutex
	departments map[int]*models.Department
	nextID      int
}

func NewMemoryRepo() *MemoryRepo {
	return &MemoryRepo{
		mu:          sync.RWMutex{},
		departments: make(map[int]*models.Department),
		nextID:      1,
	}
}

func (m *MemoryRepo) Create(ctx context.Context, d *models.Department) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	d.ID = m.nextID

	m.departments[d.ID] = d
	m.nextID++
	return nil
}

func (m *MemoryRepo) GetByID(ctx context.Context, id int) (*models.Department, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	department, ok := m.departments[id]
	if !ok {
		return nil, fmt.Errorf("no department with id:%d", id)
	}

	return department, nil
}

func (m *MemoryRepo) GetAll(ctx context.Context) ([]*models.Department, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	result := make([]*models.Department, 0, len(m.departments))

	for _, department := range m.departments {
		result = append(result, department)
	}
	return result, nil
}
