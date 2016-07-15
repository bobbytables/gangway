package server

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/bobbytables/gangway/builder"
	"github.com/bobbytables/gangway/data"
	"github.com/gorilla/mux"
)

// getDefinitionsResp contains the structure of the response
// that gangway will return when asked about all definitions
type getDefinitionsResp struct {
	Definitions []data.Definition `json:"definitions"`
}

// postDefinitionsResp contains the structure of the response
// that gangway returns upon a successful add of a definition
type postDefinitionsResp struct {
	Definition data.Definition `json:"definition"`
}

func (s *Server) getDefinitions(w http.ResponseWriter, r *http.Request) {
	defs, err := s.store.RetrieveDefinitions()
	if err != nil {
		logrus.WithError(err).Error("could not retrieve definitions")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp := getDefinitionsResp{Definitions: defs}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		logrus.WithError(err).Error("could not encode definitions")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// postDefinitions handles adding definitions to the underlying store
func (s *Server) postDefinitions(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var d data.Definition
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if vErr := data.IsValidDefinition(d); vErr != nil {
		s.writeError(w, vErr, 422)
		return
	}

	if err := s.store.AddDefinition(d); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	resp := postDefinitionsResp{Definition: d}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (s *Server) postBuildDefinition(w http.ResponseWriter, r *http.Request) {
	d, err := s.store.RetrieveDefinition(mux.Vars(r)["label"])
	if err != nil {
		s.writeError(w, err, http.StatusInternalServerError)
		return
	}

	bo := builder.BuildOpts{
		OutputStream: os.Stdout,
		Dockerfile:   d.Dockerfile,
		Tag:          d.Tag,
		Source:       d.Source,
	}

	res := s.builder.Build(bo)
	if res.Err != nil {
		s.writeError(w, res.Err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
