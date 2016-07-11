package store

import "github.com/bobbytables/shipyard/data"

// Store defines an interface for interacting with all of the persisted
// bits of Shipyard
type Store interface {
	RetrieveDefinitions() ([]data.Definition, error)
	AddDefinition(d data.Definition) error
}
