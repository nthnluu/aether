package main

import (
	"context"
	"flag"
	"github.com/nthnluu/aether/cmd/golink"
	"github.com/nthnluu/aether/pkg/server"
)

var (
	port = flag.Int("port", 9999, "Port for serving gRPC requests")
)

func main() {
	goLinkService := golink.CreateService(golink.NewRepository())
	ctx := context.Background()

	grpcServer := &server.Server{}
	grpcServer.RegisterService(goLinkService)

	grpcServer.RunOrExit(ctx, *port)
}
