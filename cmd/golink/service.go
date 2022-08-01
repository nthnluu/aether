package golink

import (
	pb "aether/pb/out"
	"context"
	"fmt"
)

type Service interface {
	// Lookup looks up a go-link using its slug and (if found) returns the destination URL.
	Lookup(ctx context.Context, lookupRequest *pb.LookupRequest) (*pb.LookupResponse, error)
	CreateLink(ctx context.Context, createLinkRequest *pb.CreateLinkRequest) (*pb.CreateLinkResponse, error)
}

type service struct {
	goLinks Repository
}

func (s *service) Lookup(ctx context.Context, lookupRequest *pb.LookupRequest) (*pb.LookupResponse, error) {
	destinationUrl, err := s.goLinks.LookupBySlug(lookupRequest.GetSlug())
	if err != nil {
		return nil, err
	}

	return &pb.LookupResponse{DestinationUrl: destinationUrl}, nil
}

func (s *service) CreateLink(ctx context.Context, createLinkRequest *pb.CreateLinkRequest) (*pb.CreateLinkResponse, error) {
	err := s.goLinks.CreateLink(createLinkRequest.GetSlug(), createLinkRequest.GetDestinationUrl())
	if err != nil {
		return nil, err
	}
	return &pb.CreateLinkResponse{Url: fmt.Sprintf("go.fsab.io/%s", createLinkRequest.GetSlug())}, nil
}

func NewService() *service {
	return &service{goLinks: NewRepository()}
}
