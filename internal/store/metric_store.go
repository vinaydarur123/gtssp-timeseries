package store

import (
	"sync"

	"github.com/your-org/gtssp/internal/model"
	"github.com/your-org/gtssp/internal/util"
)

// MetricStore defines the contract for temporary metric storage
type MetricStore interface {
	Add(metric model.Metric) error
	GetAll() []model.Metric
	Clear()
}

// InMemoryMetricStore is a concurrency-safe in-memory store
type InMemoryMetricStore struct {
	mu      sync.RWMutex
	metrics []model.Metric
}

// NewInMemoryMetricStore creates a new in-memory metric store
func NewInMemoryMetricStore() *InMemoryMetricStore {
	return &InMemoryMetricStore{
		metrics: make([]model.Metric, 0),
	}
}

// Add stores a metric in the temporary cache
func (s *InMemoryMetricStore) Add(metric model.Metric) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Trace what is being stored
	util.PrintStage("STORED IN METRICSTORE (Temporary In-Memory Cache)", metric)

	s.metrics = append(s.metrics, metric)
	return nil
}

// GetAll returns all stored metrics
func (s *InMemoryMetricStore) GetAll() []model.Metric {
	s.mu.RLock()
	defer s.mu.RUnlock()

	// Return a copy to prevent external modification
	result := make([]model.Metric, len(s.metrics))
	copy(result, s.metrics)
	return result
}

// Clear removes all metrics from the store
func (s *InMemoryMetricStore) Clear() {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.metrics = make([]model.Metric, 0)
}
