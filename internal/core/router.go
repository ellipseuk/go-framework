package core

import (
	"net/http"
)

type Router struct {
	mux *http.ServeMux
}

// NewRouter creates a new Router
func NewRouter() *Router {
	return &Router{
		mux: http.NewServeMux(),
	}
}

// Handle registers a new route with a matcher for the URL path
func (r *Router) Handle(pattern string, handler http.Handler) {
	r.mux.Handle(pattern, handler)
}

// ServeHTTP implements the http.Handler interface
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.mux.ServeHTTP(w, req)
}
