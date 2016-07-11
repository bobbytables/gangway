package store

import "github.com/bobbytables/gangway/data"

// Store defines an interface for interacting with all of the persisted
// bits of gangway
type Store interface {
	RetrieveDefinitions() ([]data.Definition, error)
	AddDefinition(d data.Definition) error
}
