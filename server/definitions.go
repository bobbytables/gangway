package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/bobbytables/gangway/builder"
	"github.com/bobbytables/gangway/data"
	"github.com/bobbytables/gangway/source"
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

func (s *Server) buildDefinition(w http.ResponseWriter, r *http.Request) {
	ds, err := s.store.RetrieveDefinitions()
	if err != nil {
		s.writeError(w, err, http.StatusInternalServerError)
		return
	}

	var d data.Definition
	for _, dd := range ds {
		if dd.Label == mux.Vars(r)["label"] {
			d = dd
			break
		}
	}

	src := source.NewSource(d.Source)
	if err := src.Pull(); err != nil {
		s.writeError(w, err, http.StatusInternalServerError)
		return
	}

	bo := builder.BuildOpts{
		ContextDir:   src.Directory(),
		OutputStream: ioutil.Discard,
		Dockerfile:   d.Dockerfile,
		Tag:          d.Tag,
	}
	res := s.builder.Build(bo)
	if res.Err != nil {
		s.writeError(w, err, http.StatusInternalServerError)
		return
	}
}
