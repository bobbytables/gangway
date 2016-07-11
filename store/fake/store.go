package fake

import (
	"github.com/bobbytables/gangway/data"
	"github.com/bobbytables/gangway/store"
)

// Store implements store.Store
type Store struct {
	FakeRetrieveDefinitions (func() ([]data.Definition, error))
	FakeAddDefinition       (func(data.Definition) error)
}

var _ store.Store = &Store{}

// RetrieveDefinitions implements store.Store
func (s *Store) RetrieveDefinitions() ([]data.Definition, error) {
	return s.FakeRetrieveDefinitions()
}

// AddDefinition adds a definition to the store
func (s *Store) AddDefinition(d data.Definition) error {
	return s.FakeAddDefinition(d)
}
