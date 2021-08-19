package gee

import (
	"net/http"
)

//type HandlerFunc func(w http.ResponseWriter, r *http.Request)
type HandlerFunc func(*Context)

// 类比 SpringMVC，Engine 就是 dispatchservlet
type Engine struct {
	router *router
}

func New() *Engine {
	return &Engine{
		router: newRouter(),
	}
}

func (e *Engine) GET(path string, handlerFunc HandlerFunc) {
	e.router.addRoute("GET", path, handlerFunc)
}

func (e *Engine) POST(path string, handlerFunc HandlerFunc) {
	e.router.addRoute("POST", path, handlerFunc)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := NewContext(w, r)
	e.router.handle(c)
}

func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}
