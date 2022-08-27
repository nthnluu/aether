package main

import (
	"context"
	"flag"
	"log"

	"github.com/nthnluu/aether/cmd/golink"
	pb "github.com/nthnluu/aether/pb/out"
	"github.com/nthnluu/aether/pkg/errors"
	"github.com/nthnluu/aether/pkg/server"
)

var (
	port = flag.Int("port", 9999, "Port for serving gRPC requests")
)

func main() {
	flag.Parse()

	// Create the service
	goLinkService := golink.CreateService(golink.NewRepository())

	// Create the gRPC server and register your service.
	grpcServer := server.CreateServer()
	pb.RegisterGoLinkServiceServer(grpcServer, goLinkService)

	// Run the server.
	server.RunOrExit(grpcServer, *port, func(b *server.ServerConfig) {
		b.AddMethodRequestInterceptor("/golink.GoLinkService/CreateLink", func(ctx context.Context, req interface{}) error {
			_, ok := req.(*pb.CreateLinkRequest)
			if !ok {
				log.Fatalf("invalid req")
			}
			return errors.NotYetImplementedError("Implement the create link request interceptor")
		})

		b.AddMethodResponseInterceptor("/golink.GoLinkService/CreateLink", func(ctx context.Context, resp interface{}) error {
			_, ok := resp.(*pb.CreateLinkResponse)
			if !ok {
				log.Fatalf("invalid resp")
			}
			return errors.NotYetImplementedError("Implement the create link request interceptor")
		})
	})
}
