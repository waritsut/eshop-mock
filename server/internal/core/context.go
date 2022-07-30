package core

import (
	"sync"

	"github.com/gin-gonic/gin"
)

type HandlerFunc func(c Context)

const (
	_UserId = "_user_id_"
)

var contextPool = &sync.Pool{
	New: func() interface{} {
		return new(MyContext)
	},
}

func newContext(ctx *gin.Context) Context {
	context := contextPool.Get().(*MyContext)
	context.ctx = ctx
	return context
}

func releaseContext(ctx Context) {
	c := ctx.(*MyContext)
	c.ctx = nil
	contextPool.Put(c)
}

type Context interface {
	Method() string
	Next()
	Bind(interface{}) error
	ShouldBindQuery(obj interface{}) error
	JSON(int, interface{})
	AbortWithStatusJSON(code int, jsonObj interface{})
	Param(key string) string
}

type MyContext struct {
	ctx *gin.Context
}

func (c *MyContext) Param(key string) string {
	return c.ctx.Param(key)
}

func (c *MyContext) Method() string {
	return c.ctx.Request.Method
}

func (c *MyContext) Next() {
	c.ctx.Next()
}

func (c *MyContext) Bind(v interface{}) error {
	return c.ctx.ShouldBindJSON(v)
}

func (c *MyContext) ShouldBindQuery(obj interface{}) error {
	return c.ctx.ShouldBindQuery(obj)
}

func (c *MyContext) JSON(statuscode int, v interface{}) {
	c.ctx.JSON(statuscode, v)
}

func (c *MyContext) AbortWithStatusJSON(code int, jsonObj interface{}) {
	c.ctx.AbortWithStatusJSON(code, jsonObj)
}

func NewMyContext(c *gin.Context) *MyContext {
	return &MyContext{ctx: c}
}

func NewGinHandler(handler func(Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler(NewMyContext(c))
	}
}
