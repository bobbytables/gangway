package server

import (
	"encoding/json"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/bobbytables/gangway/data"
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
		logrus.WithError(err).Error()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp := getDefinitionsResp{Definitions: defs}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		logrus.WithError(err).Error()
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
