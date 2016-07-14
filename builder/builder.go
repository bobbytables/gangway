package builder

import "github.com/bobbytables/gangway/data"

// Builder is an interface for building containers from definitions
type Builder interface {
	Build(d data.Definition) *Result
}

// Result contains the result of a build
type Result struct {
	Err error
}
