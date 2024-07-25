package pkg

import (
	"net/http"
)

const (
	GET    = "GET"
	POST   = "POST"
	DELETE = "DELETE"
	PUT    = "PUT"
)

type Router struct {
	routes map[string]map[string]http.HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		routes: make(map[string]map[string]http.HandlerFunc),
	}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	if handlers, ok := r.routes[req.URL.Path]; ok {
		if handler, methodExists := handlers[req.Method]; methodExists {
			handler(w, req)
			return
		}
	}

	http.NotFound(w, req)
}

func (r *Router) GET(path string, handler http.HandlerFunc) {
	r.setHandler(GET, path, handler)
}
func (r *Router) POST(path string, handler http.HandlerFunc) {
	r.setHandler(POST, path, handler)
}

func (r *Router) setHandler(method, path string, handler http.HandlerFunc) {

	if r.routes[path] == nil {
		r.routes[path] = make(map[string]http.HandlerFunc)
	}

	r.routes[path][method] = handler
}
