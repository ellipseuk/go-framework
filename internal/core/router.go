package core

import (
	"net/http"
	"regexp"
	"strings"
)

type route struct {
	method  string
	pattern *regexp.Regexp
	handler http.Handler
}

type Router struct {
	routes []route
}

func NewRouter() *Router {
	return &Router{
		routes: []route{},
	}
}

func (r *Router) Handle(method string, pattern string, handler http.Handler) {
	pattern = regexp.MustCompile(`:([a-zA-Z0-9_]+)`).ReplaceAllStringFunc(pattern, func(param string) string {
		paramName := strings.TrimPrefix(param, ":")
		return "(?P<" + paramName + ">[^/]+)"
	})

	regexPattern := regexp.MustCompile("^" + pattern + "$")
	r.routes = append(r.routes, route{method, regexPattern, handler})
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for _, route := range r.routes {
		if route.method == req.Method && route.pattern.MatchString(req.URL.Path) {
			route.handler.ServeHTTP(w, req)
			return
		}
	}
	http.NotFound(w, req)
}
