package gee

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

type Engine struct {
	route map[string]HandlerFunc
}

func New() *Engine {
	return &Engine{route: make(map[string]HandlerFunc)}
}

func (pe *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	pe.route[key] = handler
}

func (pe *Engine) GET(pattern string, handler HandlerFunc) {
	pe.addRoute("GET", pattern, handler)
}

func (pe *Engine) POST(pattern string, handler HandlerFunc) {
	pe.addRoute("POST", pattern, handler)
}

func (pe *Engine) Run(addr string) (err error) {
	fmt.Printf("server listen on %s\n", addr)
	return http.ListenAndServe(addr, pe)
}

func (pe *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := pe.route[key]; ok {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}
