package storage

import (
	"dance-instructor-api/internal/models"
	"sync"
)

type MemoryStore struct {
	instructors map[string]models.Instructor
	mu          sync.RWMutex
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		instructors: make(map[string]models.Instructor),
	}
}

func (s *MemoryStore) Create(id string, instructor models.Instructor) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.instructors[id] = instructor
}

func (s *MemoryStore) Get(id string) (models.Instructor, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	instructor, exists := s.instructors[id]
	return instructor, exists
}

func (s *MemoryStore) GetAll() []models.Instructor {
	s.mu.RLock()
	defer s.mu.RUnlock()
	instructors := make([]models.Instructor, 0, len(s.instructors))
	for _, instructor := range s.instructors {
		instructors = append(instructors, instructor)
	}
	return instructors
}

func (s *MemoryStore) Update(id string, instructor models.Instructor) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.instructors[id]; !exists {
		return false
	}
	s.instructors[id] = instructor
	return true
}

func (s *MemoryStore) Delete(id string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.instructors[id]; !exists {
		return false
	}
	delete(s.instructors, id)
	return true
}