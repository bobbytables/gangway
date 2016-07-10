package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Config stores configuration for a server instance
type Config struct{}

// Server handles incoming requests for shipyard
type Server struct {
	config Config
	m      *mux.Router
}

// NewServer initializes a server with the provided configuration
func NewServer(c Config) *Server {
	s := &Server{config: c}
	s.setupRouter()

	return s
}

// Listen starts the server on an address
func (s *Server) Listen(addr string) error {
	return http.ListenAndServe(addr, s.m)
}

func (s *Server) setupRouter() {
	s.m = mux.NewRouter()
	s.m.Handle("/recipes", NewEndpoint(getRecipes)).Methods("GET")
}
