package data

// Definition represents a gangway definition.
// Definitions are used when building containers.
type Definition struct {
	Label       string
	Source      string            `json:"source"`
	Dockerfile  string            `json:"dockerfile"`
	Environment map[string]string `json:"environment"`
	Tag         string            `json:"tag"`
}
