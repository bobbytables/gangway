package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bobbytables/gangway/builder"
	"github.com/bobbytables/gangway/data"
	"github.com/bobbytables/gangway/store/fake"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestGetDefinitions(t *testing.T) {
	st := &fake.Store{}
	s := NewServer(st)
	server := httptest.NewServer(s.m)
	defer server.Close()

	defs := []data.Definition{
		data.Definition{Source: "githubIGuess"},
	}

	st.On("RetrieveDefinitions").Return(defs, nil)

	resp, err := http.Get(server.URL + "/definitions")
	require.Nil(t, err)

	var decodedResp getDefinitionsResp

	require.Nil(t, json.NewDecoder(resp.Body).Decode(&decodedResp))
	assert.Equal(t, resp.StatusCode, 200)
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
	s := NewServer(st)
	server := httptest.NewServer(s.m)
	defer server.Close()

	st.On("AddDefinition", newD).Return(nil)

	resp, err := http.Post(server.URL+"/definitions", "application/json", r)
	require.Nil(t, err)

	var decodedResp postDefinitionsResp

	require.Nil(t, json.NewDecoder(resp.Body).Decode(&decodedResp))
	assert.Equal(t, resp.StatusCode, 201, "status code is for a create")
	assert.Equal(t, newD, decodedResp.Definition, "response body matches JSON")
	st.AssertCalled(t, "AddDefinition", newD)
}

func TestPostWithBadDefinition(t *testing.T) {
	newD := data.Definition{}

	r := new(bytes.Buffer)
	require.Nil(t, json.NewEncoder(r).Encode(newD))

	st := &fake.Store{}
	s := NewServer(st)
	server := httptest.NewServer(s.m)
	defer server.Close()

	resp, err := http.Post(server.URL+"/definitions", "application/json", r)
	require.Nil(t, err)

	var decodedResp errorResponse

	require.Nil(t, json.NewDecoder(resp.Body).Decode(&decodedResp))
	assert.Equal(t, resp.StatusCode, 422, "status code is for a entity issue")
	assert.Equal(t, "label cannot be empty", decodedResp.Error, "response body matches JSON")
}

func TestPostDefinitionToCreateImage(t *testing.T) {
	label := "hello-kitty"
	st := &fake.Store{}
	bu := &builder.FakeBuilder{}
	s := NewServer(st, bu)
	server := httptest.NewServer(s.m)
	defer server.Close()

	d := data.Definition{
		Label:       label,
		Source:      "https://github.com/bobbytables/gangway.git",
		Dockerfile:  "Dockerfile",
		Environment: nil,
		Tag:         "lol-whut",
	}

	bu.On("Build", mock.Anything).Return(&builder.Result{}).Once()
	st.On("RetrieveDefinition", label).Return(d, nil)

	resp, err := http.Post(server.URL+"/definitions/"+label, "application/json", nil)
	require.Nil(t, err)
	assert.Equal(t, http.StatusAccepted, resp.StatusCode, "status is a 202")

	st.AssertCalled(t, "RetrieveDefinition", label)
	bu.AssertCalled(t, "Build", mock.Anything)
}
