package golink

import (
	"context"
	"fmt"

	pb "github.com/nthnluu/aether/pb/out"
)

type service struct {
	goLinks Repository
	*pb.UnimplementedGoLinkServiceServer
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

func CreateService(goLinksRepository Repository) *service {
	return &service{
		goLinks: goLinksRepository,
	}
}
