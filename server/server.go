package server

import (
	"net/http"

	"github.com/bobbytables/gangway/store"
	"github.com/gorilla/mux"
)

// Config stores configuration for a server instance
type Config struct{}

// Server handles incoming requests for gangway
type Server struct {
	config Config
	store  store.Store
	m      *mux.Router
}

// NewServer initializes a server with the provided configuration
func NewServer(config Config, store store.Store) *Server {
	s := &Server{config: config, store: store}
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
}
