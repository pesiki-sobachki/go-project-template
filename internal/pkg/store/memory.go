package store

import (
	"context"
	"sync"

	"github.com/shanth1/gotools/errs"
)

var ErrNotFound = errs.ErrNotFound

type MemoryStore[K comparable, V any] struct {
	mu   sync.RWMutex
	data map[K]V
}

func NewMemoryStore[K comparable, V any]() *MemoryStore[K, V] {
	return &MemoryStore[K, V]{
		data: make(map[K]V),
	}
}

func (s *MemoryStore[K, V]) Set(_ context.Context, key K, value V) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = value
	return nil
}

func (s *MemoryStore[K, V]) Get(_ context.Context, key K) (V, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	val, ok := s.data[key]
	if !ok {
		var empty V
		return empty, ErrNotFound
	}
	return val, nil
}

func (s *MemoryStore[K, V]) Delete(_ context.Context, key K) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.data, key)
	return nil
}
