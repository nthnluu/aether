package golink

import (
	pb "aether/pb/out"
	"aether/pkg/grpc"
	"context"
)

const (
	LinkReadPermission   = "golink.link.read"
	LinkCreatePermission = "golink.link.create"
)

type grpcServer struct {
	service Service
	pb.UnimplementedGoLinkServiceServer
}

func NewGRPCServer(s Service) pb.GoLinkServiceServer {
	return &grpcServer{
		service: s,
	}
}

func (s *grpcServer) Lookup(ctx context.Context, req *pb.LookupRequest) (*pb.LookupResponse, error) {
	endpoint := grpc.Endpoint(s.service.Lookup)         // Create a transport endpoint from a service method.
	endpoint.SetRequestValidator(ValidateLookupRequest) // Set ValidateLookupRequest as the request validation func.
	endpoint.SetDeadline(5)                             // Set the request deadline to 5 seconds.
	return endpoint.ServeGRPC(ctx, req)
}

func (s *grpcServer) CreateLink(ctx context.Context, req *pb.CreateLinkRequest) (*pb.CreateLinkResponse, error) {
	endpoint := grpc.Endpoint(s.service.CreateLink)
	endpoint.SetRequestValidator(ValidateCreateLinkRequest)
	endpoint.SetDeadline(180)
	return endpoint.ServeGRPC(ctx, req)
}
