package builder

// Builder is an interface for building containers from definitions
type Builder interface {
	Build(bo BuildOpts) *Result
}

// Result contains the result of a build
type Result struct {
	Err error
}
