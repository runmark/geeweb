package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	W          http.ResponseWriter
	R          *http.Request
	Path       string
	Method     string
	StatusCode int
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		W:      w,
		R:      r,
		Path:   r.URL.Path,
		Method: r.Method,
	}
}

func (c *Context) Query(name string) string {
	return c.R.URL.Query().Get(name)
}

func (c *Context) PostForm(name string) string {
	return c.R.FormValue(name)
}

func (c *Context) SetHeader(k, v string) {
	c.W.Header().Set(k, v)
}

func (c *Context) Status(statuscode int) {
	c.StatusCode = statuscode
	c.W.WriteHeader(statuscode)
}

func (c *Context) Data(statuscode int, data []byte) {
	c.Status(statuscode)
	c.W.Write(data)
}

func (c *Context) HTML(statuscode int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(statuscode)
	c.W.Write([]byte(html))
}

func (c *Context) String(statuscode int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(statuscode)
	fmt.Fprintf(c.W, format, values...)
}

func (c *Context) JSON(statuscode int, h H) {

	c.SetHeader("Content-Type", "application/json")
	c.Status(statuscode)

	err := json.NewEncoder(c.W).Encode(h)
	if err != nil {
		http.Error(c.W, err.Error(), http.StatusBadRequest)
		return
	}
}
