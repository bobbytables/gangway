package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefinitionValidation(t *testing.T) {
	testcases := []struct {
		scenario   string
		definition Definition
		err        error
	}{
		{
			scenario:   "without a label is invalid",
			definition: Definition{Label: "", Source: "git@github.com:bobbytables/gangway.git", Dockerfile: "Dockerfile"},
			err:        definitionError{Failure: "label cannot be empty"},
		},
		{
			scenario:   "a label with spaces is invalid",
			definition: Definition{Label: "hello world", Source: "git@github.com:bobbytables/gangway.git", Dockerfile: "Dockerfile"},
			err:        definitionError{Failure: "label must be only letters, numbers, and dashes"},
		},
		{
			scenario:   "source must be present",
			definition: Definition{Label: "hello-world", Source: "", Dockerfile: "Dockerfile"},
			err:        definitionError{Failure: "source cannot be empty"},
		},
		{
			scenario:   "dockerfile must be present",
			definition: Definition{Label: "hello-world", Source: "git", Dockerfile: ""},
			err:        definitionError{Failure: "dockerfile cannot be empty"},
		},
	}

	for _, tc := range testcases {
		err := IsValidDefinition(tc.definition)
		assert.Equal(t, tc.err, err)
	}
}
