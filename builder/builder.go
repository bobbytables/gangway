package builder

import "github.com/bobbytables/gangway/data"

// Builder is an interface for building containers from definitions
type Builder interface {
	Build(d data.Definition) *Result
}

// Result contains the result of a build
type Result struct {
	err error
}

// Err returns an error on a response (if any)
func (r Result) Err() error {
	return r.err
}
