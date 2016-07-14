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
func NewServer(config Config, store store.Store, builder builder.Builder) *Server {
	s := &Server{
		config:  config,
		store:   store,
		builder: builder,
	}

	s.setupRouter()

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
	s.m.Handle("/definitions/{label}", NewEndpoint(s.buildDefinition)).Methods("POST")
}

func (s *Server) writeError(w http.ResponseWriter, err error, code int) {
	w.WriteHeader(code)
	errResp := errorResponse{Error: err.Error()}
	if err := json.NewEncoder(w).Encode(errResp); err != nil {
		return
	}
}
