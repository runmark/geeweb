package gee

import (
	"net/http"
)

//type HandlerFunc func(w http.ResponseWriter, r *http.Request)

type HandlerFunc func(*Context)


// 类比 SpringMVC，Engine 就是 dispatchservlet
type Engine struct {
	router map[string]HandlerFunc
}

func New() *Engine {
	return &Engine{
		router: make(map[string]HandlerFunc),
	}
}

func (e *Engine) addRoute(method, path string, handlerFunc HandlerFunc) {
	key := method + "-" + path
	e.router[key] = handlerFunc
}

func (e *Engine) GET(path string, handlerFunc HandlerFunc) {
	e.addRoute("GET", path, handlerFunc)
}

func (e *Engine) POST(path string, handlerFunc HandlerFunc) {
	e.addRoute("POST", path, handlerFunc)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.Method + "-" + r.URL.Path

	c := NewContext(w, r)
	f, ok := e.router[key]
	if ok {
		f(c)
		return
	}

	w.WriteHeader(http.StatusNotFound)
}

func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}
