package server

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Service interface {
	Service() interface{}
	ServiceDesc() *grpc.ServiceDesc
}

// Interceptor is a function that is run before a request or after a response.
type Interceptor = func(context.Context, interface{}) error

// interceptorConfig represents the registered request and response interceptors
// for a specific method.
type interceptorConfig struct {
	req  []Interceptor
	resp []Interceptor
}

func (i *interceptorConfig) init() {
	i.req = make([]Interceptor, 0)
	i.resp = make([]Interceptor, 0)
}

// methodInterceptorTable is a lookup table that maps GRPC method names to
// their registered interceptors.
var methodInterceptorTable map[string]interceptorConfig

// globalInterceptors contains interceptors that will be run on every RPC call.
var globalInterceptors interceptorConfig

func init() {
	methodInterceptorTable = make(map[string]interceptorConfig)
	globalInterceptors.init()
}

// RunOrExit runs the provided gRPC server on the specified port. `configure` is a
// function that can be used to set up resources and configure the server before running.
func RunOrExit(grpcServer *grpc.Server, port int, configure func(*ServerConfig)) {
	serverConfig := &ServerConfig{}
	configure(serverConfig)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Server is running on port %v\n", port)
	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}

// internalInterceptor runs the registered interceptors for requests then responses
func internalInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	fmt.Println(info.FullMethod)

	// Run globally registered req interceptors.
	for _, interceptor := range globalInterceptors.req {
		if err := interceptor(ctx, req); err != nil {
			return nil, err
		}
	}

	// Run req interceptors registered for the specific method.
	methodInterceptors, ok := methodInterceptorTable[info.FullMethod]
	if ok {
		// Run all registered interceptors for the current method.
		for _, interceptor := range methodInterceptors.req {
			if err := interceptor(ctx, req); err != nil {
				return nil, err
			}
		}
	}

	resp, err := handler(ctx, req)

	// Run globally registered resp interceptors.
	for _, interceptor := range globalInterceptors.resp {
		if err := interceptor(ctx, resp); err != nil {
			return nil, err
		}
	}

	// Run interceptors registered for the specific method.
	if ok {
		// Run all registered interceptors for the current method.
		for _, interceptor := range methodInterceptors.resp {
			if err := interceptor(ctx, resp); err != nil {
				return nil, err
			}
		}
	}
	return resp, err
}

func CreateServer() *grpc.Server {
	return grpc.NewServer(grpc.UnaryInterceptor(internalInterceptor))
}

type ServerConfig struct {
}

type RequestInterceptor = func()

// AddMethodRequestInterceptor adds a function that will be called with an incoming request before the specified
// method is called. Interceptors are called in FIFO order: interceptors added first will be called first.
// The `methodName` is the full RPC method string, i.e., /package.service/method.
func (*ServerConfig) AddMethodRequestInterceptor(methodName string, interceptor Interceptor) {
	config, ok := methodInterceptorTable[methodName]
	if ok {
		config.req = append(config.req, interceptor)
		methodInterceptorTable[methodName] = config
	} else {
		newConfig := &interceptorConfig{}
		newConfig.init()
		newConfig.req = append(newConfig.req, interceptor)
		methodInterceptorTable[methodName] = *newConfig
	}
}

// AddMethodResponseInterceptor adds a function that will be called with an outgoing response from a specified method
// before it's sent to the caller. Interceptors are called in FIFO order: interceptors added first will be called first.
func (*ServerConfig) AddMethodResponseInterceptor(methodName string, interceptor Interceptor) {
	config, ok := methodInterceptorTable[methodName]
	if ok {
		config.resp = append(config.resp, interceptor)
		methodInterceptorTable[methodName] = config
	} else {
		newConfig := &interceptorConfig{}
		newConfig.init()
		newConfig.req = append(newConfig.resp, interceptor)
		methodInterceptorTable[methodName] = *newConfig
	}
}

// AddRequestInterceptor adds a function that will be called with an incoming request before the handler is called.
// Interceptors are called in FIFO order: interceptors added first will be called first.
func (*ServerConfig) AddRequestInterceptor(interceptor Interceptor) {
	globalInterceptors.req = append(globalInterceptors.req, interceptor)
}

// AddResponseInterceptor adds a function that will be called with an outgoing response before it's sent to the caller.
// Interceptors are called in FIFO order: interceptors added first will be called first.
func (*ServerConfig) AddResponseInterceptor(interceptor Interceptor) {
	globalInterceptors.resp = append(globalInterceptors.resp, interceptor)
}
