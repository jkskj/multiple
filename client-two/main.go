package main

import (
	"context"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"multiple/kitex_gen/echo"
	"multiple/kitex_gen/echo/two"
)

func main() {
	cli, err := two.NewClient("two", client.WithHostPorts("127.0.0.1:10001"))
	if err != nil {
		klog.Warnf("failed to new client: %s", err)
		return
	}
	req := &echo.Request{}
	resp, err := cli.Two(context.Background(), req)
	if err != nil {
		klog.Warnf("failed to call: %s", err)
		return
	}
	klog.Infof("resp: %s", resp.Message)
}
