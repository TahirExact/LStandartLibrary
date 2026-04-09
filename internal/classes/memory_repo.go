package classes

import (
	"context"
	"fmt"
	"my-project/internal/models"
	"sync"
)

type MemoryRepo struct {
	mu      sync.RWMutex
	classes map[int]*models.Class
	nextID  int
}

func NewMemoryRepo() *MemoryRepo {
	return &MemoryRepo{
		mu:      sync.RWMutex{},
		classes: make(map[int]*models.Class),
		nextID:  1,
	}
}

func (m *MemoryRepo) Create(ctx context.Context, c *models.Class) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	c.ID = m.nextID
	m.classes[c.ID] = c
	m.nextID++
	return nil
}

func (m *MemoryRepo) GetAll(ctx context.Context) ([]*models.Class, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	result := make([]*models.Class, 0, len(m.classes))

	for _, class := range m.classes {
		result = append(result, class)
	}
	return result, nil
}

func (m *MemoryRepo) GetByID(ctx context.Context, id int) (*models.Class, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	class, ok := m.classes[id]
	if !ok {
		return nil, fmt.Errorf("there is no such class with id:%d", id)
	}
	return class, nil
}
