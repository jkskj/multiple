package main

import (
	"context"

	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"multiple/kitex_gen/echo"
	"multiple/kitex_gen/echo/one"
)

func main() {
	cli, err := one.NewClient("one", client.WithHostPorts("127.0.0.1:10001"), client.WithTransportProtocol(transport.TTHeader), client.WithMetaHandler(transmeta.ClientTTHeaderHandler))
	if err != nil {
		klog.Warnf("failed to new client: %s", err)
		return
	}
	req := &echo.Request{}
	resp, err := cli.One(context.Background(), req)
	if err != nil {
		klog.Warnf("failed to call: %s", err)
		return
	}
	klog.Infof("resp: %s", resp.Message)
}
