package server

import (
	"net/http"

	"github.com/Sirupsen/logrus"
)

// Endpoint contains the handler for a particular route
type Endpoint struct {
	handler http.HandlerFunc
}

// NewEndpoint creates an enpoint
func NewEndpoint(h http.HandlerFunc) Endpoint {
	return Endpoint{handler: h}
}

// ServeHTTP delegates an incoming request to the appointed handler
func (e Endpoint) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logrus.WithField("path", r.URL.Path).Info()
	e.handler(w, r)
}
