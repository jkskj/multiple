package main

import (
	"context"
	"net"

	"github.com/cloudwego/kitex/pkg/klog"

	"github.com/cloudwego/kitex/server"

	"multiple/kitex_gen/echo"
	"multiple/kitex_gen/echo/one"
	"multiple/kitex_gen/echo/two"
)

var _ echo.One = &server1{}

type server1 struct{}

func (o server1) One(context.Context, *echo.Request) (res *echo.Response, err error) {
	klog.Info("one")
	return &echo.Response{Message: "one"}, nil
}

var _ echo.Two = &server2{}

type server2 struct{}

func (o server2) Two(context.Context, *echo.Request) (res *echo.Response, err error) {
	klog.Info("two")
	return &echo.Response{Message: "two"}, nil
}

func main() {
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:10001")
	if err != nil {
		panic(err)
	}
	svr := server.NewServer(
		server.WithServiceAddr(addr),
	)
	if err := one.RegisterService(svr, new(server1)); err != nil {
		panic(err)
	}
	if err := two.RegisterService(svr, new(server2)); err != nil {
		panic(err)
	}

	if err = svr.Run(); err != nil {
		klog.Info(err.Error())
	}
}

// kitex -module multiple ./idl/echo.proto

//	kitexcall -idl-path  idl/echo.proto -m Two/Two -d '{}' -e 127.0.0.1:10001
//	[Status]: Success
//	{
//    	"message": "two"
//	}

//	kitexcall -idl-path  idl/echo.proto -m Two/Two -d '{}' -e 127.0.0.1:10001 -transport TTHeader
//	[Status]: Success
//	{
//    	"message": "two"
//	}

//	kitexcall -idl-path  idl/echo.proto -m One/One -d '{}' -e 127.0.0.1:10001
//	[Status]: Failed
//	[ServerError]: RPC call failed: missing method: One in service

//	kitexcall -idl-path  idl/echo.proto -m One/Two -d '{}' -e 127.0.0.1:10001
//	[Status]: Success
//	{
//    	"message": "two"
//	}

//	kitexcall -idl-path  idl/echo.proto -m One/One -d '{}' -e 127.0.0.1:10001 -transport TTHeader
//	[Status]: Failed
//	[ServerError]: RPC call failed: missing method: One in service

//	kitexcall -idl-path  idl/echo.proto -m One/Two -d '{}' -e 127.0.0.1:10001 -transport TTHeader
//	[Status]: Failed
//	[ServerError]: RPC call failed: remote or network error: default codec read failed: EOF peer close
