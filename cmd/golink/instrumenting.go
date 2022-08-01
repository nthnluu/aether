package golink

import (
	pb "aether/pb/out"
	"context"
	"github.com/go-kit/log"
	"go.opencensus.io/trace"
	"time"
)

type instrumentingService struct {
	logger log.Logger
	Service
}

// NewInstrumentingService returns an instance of an instrumenting Service.
func NewInstrumentingService(logger log.Logger, s Service) Service {
	return &instrumentingService{
		logger:  logger,
		Service: s,
	}
}

func (s *instrumentingService) Lookup(ctx context.Context, req *pb.LookupRequest) (res *pb.LookupResponse, err error) {
	ctx, span := trace.StartSpan(ctx, "golink.grpc.Lookup")
	defer span.End()

	defer func(begin time.Time) {
		s.logger.Log(
			"method", "lookup",
			"slug", req.GetSlug(),
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())

	return s.Service.Lookup(ctx, req)
}

func (s *instrumentingService) CreateLink(ctx context.Context, req *pb.CreateLinkRequest) (res *pb.CreateLinkResponse, err error) {
	ctx, span := trace.StartSpan(ctx, "golink.grpc.CreateLink")
	defer span.End()

	defer func(begin time.Time) {
		s.logger.Log(
			"method", "createLink",
			"slug", req.GetSlug(),
			"destination_url", req.GetDestinationUrl(),
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())

	return s.Service.CreateLink(ctx, req)
}
