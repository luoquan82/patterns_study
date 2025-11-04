package chainOfResponsibility

import (
	"fmt"
	"net/http"
)

type Context struct {
	request  *http.Request
	w        http.ResponseWriter
	index    int
	handlers []HandlerFun
}

func (c *Context) Next() {
	c.index++
	if c.index < len(c.handlers) {
		c.handlers[c.index](c)
	}
}

func (c *Context) Abort() {
	c.index = len(c.handlers)
}

type HandlerFun func(*Context)
type Engine struct {
	handles []HandlerFun
}

func (e *Engine) Use(f HandlerFun) {
	e.handles = append(e.handles, f)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	context := &Context{
		request:  r,
		w:        w,
		index:    -1,
		handlers: e.handles,
	}
	context.Next()
}

func AuthMiddleware(c *Context) {
	fmt.Println("认证中间件")
}

func LogMiddleware(c *Context) {
	fmt.Println("日志中间件")
	c.Next()
}
