package data

import "regexp"

// Definition represents a gangway definition.
// Definitions are used when building containers.
type Definition struct {
	Label       string            `json:"label"`
	Source      string            `json:"source"`
	Dockerfile  string            `json:"dockerfile"`
	Environment map[string]string `json:"environment"`
	Tag         string            `json:"tag"`
}

type definitionError struct {
	Failure string
}

func (d definitionError) Error() string {
	return d.Failure
}

var validationRegex = regexp.MustCompile("^[A-Za-z0-9-]+$")

// IsValidDefinition checks if a definition is valid to be stored
func IsValidDefinition(d Definition) error {
	if d.Label == "" {
		return definitionError{Failure: "label cannot be empty"}
	}

	if d.Source == "" {
		return definitionError{Failure: "source cannot be empty"}
	}

	if d.Dockerfile == "" {
		return definitionError{Failure: "dockerfile cannot be empty"}
	}

	if !validationRegex.MatchString(d.Label) {
		return definitionError{Failure: "label must be only letters, numbers, and dashes"}
	}

	return nil
}
