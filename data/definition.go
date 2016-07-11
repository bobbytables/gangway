package data

// Definition represents a gangway definition.
// Definitions are used when building containers.
type Definition struct {
	Label       string
	Source      string
	Dockerfile  string
	Environment map[string]string
	Tag         string
}
