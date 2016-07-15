package fake

import (
	"github.com/bobbytables/gangway/data"
	"github.com/bobbytables/gangway/store"

	"github.com/stretchr/testify/mock"
)

// Store implements store.Store
type Store struct {
	mock.Mock
}

var _ store.Store = &Store{}

// RetrieveDefinitions implements store.Store
func (s *Store) RetrieveDefinitions() ([]data.Definition, error) {
	o := s.Called()
	return o.Get(0).([]data.Definition), o.Error(1)
}

// AddDefinition adds a definition to the store
func (s *Store) AddDefinition(d data.Definition) error {
	o := s.Called(d)
	return o.Error(0)
}

// RetrieveDefinition retrieves a definition
func (s *Store) RetrieveDefinition(label string) (data.Definition, error) {
	o := s.Called(label)
	return o.Get(0).(data.Definition), o.Error(1)
}
