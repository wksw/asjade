package rest

import (
	"asjade/server"
	"errors"
)

const (
	// NAME server name
	NAME = "rest"
)

// Context http请求上下文
type Context struct {
	options *server.Options
}

// ReadEntry 参数解析
func (c *Context) ReadEntry(req interface{}) error {
	return nil
}

// ReadQueryEntry 参数解析，不解析body参数
func (c *Context) ReadQueryEntry(req interface{}) error {
	return nil
}

func init() {
	server.RegisteServer(NAME, New)
}

// New .
func New(options ...server.Option) (server.Server, error) {
	return &Context{}, nil
}

// Registe .
func (c *Context) Registe(handle interface{}, options ...server.RegisteOption) error {
	if _, ok := handle.(Handler); !ok {
		return errors.New("unimplement handler")
	}
	return nil
}

// Start start the rest server
func (c *Context) Start(option ...server.RunOption) error {
	return nil
}

// Stop stop the rest server
func (c *Context) Stop() error {
	return nil
}
