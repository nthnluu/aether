package main

import (
	"context"
	"flag"

	"github.com/nthnluu/aether/cmd/golink"
	pb "github.com/nthnluu/aether/pb/out"
	"github.com/nthnluu/aether/pkg/errors"
	"github.com/nthnluu/aether/pkg/server"
)

var (
	port = flag.Int("port", 9999, "Port for serving gRPC requests")
)

func main() {
	goLinkService := golink.CreateService(golink.NewRepository())
	grpcServer := server.CreateServer()
	pb.RegisterGoLinkServiceServer(grpcServer, goLinkService)

	server.RunOrExit(grpcServer, *port, func(b *server.ServerConfig) {
		b.AddMethodRequestInterceptor("/golink.GoLinkService/CreateLink", func(ctx context.Context, req interface{}) error {
			_, ok := req.(*pb.CreateLinkRequest)
			if !ok {
				panic("invalid req")
			}
			return errors.NotYetImplemented("Implement the create link request interceptor")
		})

		b.AddMethodResponseInterceptor("/golink.GoLinkService/CreateLink", func(ctx context.Context, resp interface{}) error {
			_, ok := resp.(*pb.CreateLinkResponse)
			if !ok {
				panic("invalid resp")
			}
			return errors.NotYetImplemented("Implement the create link request interceptor")
		})
	})
}
