// Code generated by Kitex v0.10.1. DO NOT EDIT.

package two

import (
	server "github.com/cloudwego/kitex/server"
	echo "multiple/kitex_gen/echo"
)

// NewInvoker creates a server.Invoker with the given handler and options.
func NewInvoker(handler echo.Two, opts ...server.Option) server.Invoker {
	var options []server.Option

	options = append(options, opts...)

	s := server.NewInvoker(options...)
	if err := s.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	return s
}