package grpc

import (
	"context"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"time"
)

type (
	request  = any
	response = any
)

// makeTransportEndpoint converts a service endpoint into a transport endpoint.
func makeTransportEndpoint[ReqT request, ResT response](serviceEndpoint func(ctx context.Context, req *ReqT) (res *ResT, err error)) *grpctransport.Server {
	return grpctransport.NewServer(
		func(ctx context.Context, request interface{}) (interface{}, error) {
			req := request.(*ReqT)
			res, err := serviceEndpoint(ctx, req)
			if err != nil {
				return nil, err
			}

			return res, nil
		},
		func(_ context.Context, grpcReq interface{}) (interface{}, error) {
			req := grpcReq.(*ReqT)
			return req, nil
		},
		func(_ context.Context, response interface{}) (interface{}, error) {
			return response.(*ResT), nil
		},
	)
}

// Endpoint creates a new `endpoint` that responds to requests (of type ReqT) with a service endpoint.
func Endpoint[ReqT request, ResT response](serviceEndpoint func(ctx context.Context, req *ReqT) (res *ResT, err error)) *endpoint[ReqT, ResT] {
	return &endpoint[ReqT, ResT]{
		serviceEndpoint: serviceEndpoint,
	}
}

type endpoint[ReqT request, ResT response] struct {
	// serviceEndpoint is a function that
	serviceEndpoint interface{}
	// deadline represents the maximum amount of time (in seconds) to wait for the endpoint to complete before
	// cancelling the request.
	deadline uint64
	// requestValidator is a function that validates the incoming request.
	requestValidator interface{}
	// aclPolicy represents a Bouncer ACL rules that will be used to determine access to this endpoint.
	aclPolicy interface{}
	// requiredPermission defines a required Gatekeeper permission.
	requiredPermission string
	// allowUnauthenticatedRequests determines whether
	allowUnauthenticatedRequests bool
}

// SetRequestValidator adds a validator function to the endpoint that will be used to validate incoming requests.
func (e *endpoint[ReqT, ResT]) SetRequestValidator(validator func(*ReqT) error) {
	e.requestValidator = validator
}

// SetDeadline sets the maximum amount of time (in seconds) to wait for the endpoint to complete before
// cancelling the request.
func (e *endpoint[ReqT, ResT]) SetDeadline(seconds uint64) {
	e.deadline = seconds
}

// RequirePermission sets the maximum amount of time (in seconds) to wait for the endpoint to complete before
// cancelling the request.
func (e *endpoint[ReqT, ResT]) RequirePermission(permission string) {
	e.requiredPermission = permission
}

// ServeGRPC serves an incoming request using the endpoint.
func (e *endpoint[ReqT, ResT]) ServeGRPC(ctx context.Context, req *ReqT) (r *ResT, err error) {
	// If provided, validate the request with the requestValidator function.
	if e.requestValidator != nil {
		validatorFunc := e.requestValidator.(func(req *ReqT) (err error))
		if err = validatorFunc(req); err != nil {
			return nil, err
		}
	}

	if e.requiredPermission != "" {
		// TODO: Add permission check
	}

	serviceEndpoint := e.serviceEndpoint.(func(ctx context.Context, req *ReqT) (res *ResT, err error))
	grpcEndpoint := makeTransportEndpoint[ReqT, ResT](serviceEndpoint)

	clientDeadline := time.Now().Add(time.Duration(10) * time.Millisecond)
	ctx.Deadline()
	ctx, cancel := context.WithDeadline(ctx, clientDeadline)
	defer cancel()

	_, res, err := grpcEndpoint.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*ResT), nil
}
