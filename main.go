package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/nthnluu/aether/cmd/golink"
	pb "github.com/nthnluu/aether/pb/out"
	"github.com/nthnluu/aether/pkg/server"
)

var (
	port = flag.Int("port", 9999, "Port for serving gRPC requests")
)

func main() {
	goLinkService := golink.CreateService(golink.NewRepository())
	grpcServer := server.CreateServer()
	pb.RegisterGoLinkServiceServer(grpcServer, goLinkService)

	server.RunOrExit(grpcServer, *port, func(b *server.ServerConfigurationBuilder) {
		b.AddMethodRequestInterceptor("/golink.GoLinkService/CreateLink", func(ctx context.Context, req interface{}) {
			request, ok := req.(*pb.CreateLinkRequest)
			if !ok {
				panic("ahahhahaha")
			}
			fmt.Printf(request.DestinationUrl)
		})

		b.AddMethodResponseInterceptor("/golink.GoLinkService/CreateLink", func(ctx context.Context, resp interface{}) {
			response, ok := resp.(*pb.CreateLinkResponse)
			if !ok {
				panic("ahahhahaha")
			}
			response.Url = response.GetUrl() + "HI NAGTHAN"
		})
	})
}
