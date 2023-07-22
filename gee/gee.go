package gee

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request)

type Engine struct {
	router map[string]HandlerFunc
}

func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

func (e *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	e.router[key] = handler
}

func (e *Engine) GET(pattern string, h HandlerFunc) {
	e.addRoute("GET", pattern, h)
}

func (e *Engine) POST(pattern string, h HandlerFunc) {
	e.addRoute("POST", pattern, h)
}

func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.Method + "-" + r.URL.Path
	if h, ok := e.router[key]; ok {
		h(w, r)
	} else {
		fmt.Fprintf(w, "404 NOT FOUNT: %s\n", r.URL)
	}
}
