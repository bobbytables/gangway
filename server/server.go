package server

import (
	"encoding/json"
	"net/http"

	"github.com/bobbytables/gangway/builder"
	"github.com/bobbytables/gangway/store"

	"github.com/gorilla/mux"
)

// Config stores configuration for a server instance
type Config struct{}

// Server handles incoming requests for gangway
type Server struct {
	config Config
	m      *mux.Router

	store   store.Store
	builder builder.Builder
}

// NewServer initializes a server with the provided configuration
func NewServer(opts ...interface{}) *Server {
	// config Config, store store.Store, builder builder.Builder
	s := &Server{}
	s.setupRouter()
	s.setupOptions(opts...)

	return s
}

// Listen starts the server on an address
func (s *Server) Listen(addr string) error {
	return http.ListenAndServe(addr, s.m)
}

func (s *Server) setupRouter() {
	s.m = mux.NewRouter()
	s.m.Handle("/definitions", NewEndpoint(s.getDefinitions)).Methods("GET")
	s.m.Handle("/definitions", NewEndpoint(s.postDefinitions)).Methods("POST")
	s.m.Handle("/definitions/{label}", NewEndpoint(s.postBuildDefinition)).Methods("POST")
}

func (s *Server) setupOptions(opts ...interface{}) {
	for _, v := range opts {
		switch v.(type) {
		case Config:
			s.config = v.(Config)
		case store.Store:
			s.store = v.(store.Store)
		case builder.Builder:
			s.builder = v.(builder.Builder)
		}
	}
}

func (s *Server) writeError(w http.ResponseWriter, err error, code int) {
	w.WriteHeader(code)
	errResp := errorResponse{Error: err.Error()}
	if err := json.NewEncoder(w).Encode(errResp); err != nil {
		return
	}
}
