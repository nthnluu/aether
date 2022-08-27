package aether

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
)

// FullMethodName is the key of the full method name in the context passed into interceptor functions.
// This allows interceptor functions to access the full method name from `ctx`.
const FullMethodName = "fullMethodName"

// Interceptor is a function that is run before a request or after a response.
type Interceptor = func(context.Context, interface{}) error

// internalInterceptor runs the registered interceptors for requests then responses
func internalInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	fmt.Println(info.FullMethod)

	// Run globally registered req interceptors.
	for _, interceptor := range globalInterceptors.req {
		if err := interceptor(context.WithValue(ctx, FullMethodName, info.FullMethod), req); err != nil {
			return nil, err
		}
	}

	// Run req interceptors registered for the specific method.
	methodInterceptors, ok := methodInterceptorTable[info.FullMethod]
	if ok {
		// Run all registered interceptors for the current method.
		for _, interceptor := range methodInterceptors.req {
			if err := interceptor(context.WithValue(ctx, FullMethodName, info.FullMethod), req); err != nil {
				return nil, err
			}
		}
	}

	resp, err := handler(ctx, req)

	// Run globally registered resp interceptors.
	for _, interceptor := range globalInterceptors.resp {
		if err := interceptor(context.WithValue(ctx, FullMethodName, info.FullMethod), resp); err != nil {
			return nil, err
		}
	}

	// Run interceptors registered for the specific method.
	if ok {
		// Run all registered interceptors for the current method.
		for _, interceptor := range methodInterceptors.resp {
			if err := interceptor(context.WithValue(ctx, FullMethodName, info.FullMethod), resp); err != nil {
				return nil, err
			}
		}
	}
	return resp, err
}
