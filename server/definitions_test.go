package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bobbytables/gangway/data"
	"github.com/bobbytables/gangway/store/fake"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetDefinitions(t *testing.T) {
	st := &fake.Store{}
	s := NewServer(Config{}, st, nil)
	rec := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/definitions", nil)
	require.Nil(t, err)

	defs := []data.Definition{
		data.Definition{Source: "githubIGuess"},
	}

	st.FakeRetrieveDefinitions = func() ([]data.Definition, error) {
		return defs, nil
	}

	s.getDefinitions(rec, req)

	var decodedResp getDefinitionsResp

	require.Nil(t, json.NewDecoder(rec.Body).Decode(&decodedResp))
	assert.Equal(t, rec.Code, 200)
	assert.Equal(t, len(defs), len(decodedResp.Definitions))
	assert.Equal(t, defs, decodedResp.Definitions)
}

func TestPostDefinitions(t *testing.T) {
	newD := data.Definition{
		Label:      "hcm-rails",
		Source:     "git@github.com:bobbytables/gangway.git",
		Dockerfile: "Dockerfile-build",
		Tag:        "gangway/gangway:latest",
	}

	r := new(bytes.Buffer)
	require.Nil(t, json.NewEncoder(r).Encode(newD))

	st := &fake.Store{}
	s := NewServer(Config{}, st, nil)
	rec := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/definitions", r)
	require.Nil(t, err)

	st.FakeAddDefinition = func(d data.Definition) error {
		assert.Equal(t, newD, d, "definition matching")
		return nil
	}

	s.postDefinitions(rec, req)

	var decodedResp postDefinitionsResp

	require.Nil(t, json.NewDecoder(rec.Body).Decode(&decodedResp))
	assert.Equal(t, rec.Code, 201, "status code is for a create")
	assert.Equal(t, newD, decodedResp.Definition, "response body matches JSON")
}

func TestPostWithBadDefinition(t *testing.T) {
	newD := data.Definition{}

	r := new(bytes.Buffer)
	require.Nil(t, json.NewEncoder(r).Encode(newD))

	st := &fake.Store{}
	s := NewServer(Config{}, st, nil)
	rec := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/definitions", r)
	require.Nil(t, err)

	st.FakeAddDefinition = func(d data.Definition) error {
		return nil
	}

	s.postDefinitions(rec, req)

	var decodedResp errorResponse

	require.Nil(t, json.NewDecoder(rec.Body).Decode(&decodedResp))
	assert.Equal(t, rec.Code, 422, "status code is for a entity issue")
	assert.Equal(t, "label cannot be empty", decodedResp.Error, "response body matches JSON")
}
