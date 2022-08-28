package aether

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Service interface {
	Service() interface{}
	ServiceDesc() *grpc.ServiceDesc
}

// interceptorConfig represents the registered request and response interceptors
// for a specific method.
type interceptorConfig struct {
	req  []Interceptor
	resp []Interceptor
}

// ConfigureFunc is a function that is used to set up an Aether server.
// It is passed a pointer to a ServerConfig, which can be used to add interceptors,
// register services, and more.
type ConfigureFunc func(c *ServerConfig) error

func (i *interceptorConfig) init() {
	i.req = make([]Interceptor, 0)
	i.resp = make([]Interceptor, 0)
}

// methodInterceptorTable is a lookup table that maps GRPC method names to
// their registered interceptors.
var methodInterceptorTable map[string]interceptorConfig

var globalInterceptors interceptorConfig

func init() {
	methodInterceptorTable = make(map[string]interceptorConfig)
	globalInterceptors.init()
}

// RunOrExit runs the provided gRPC server on the specified port. `configure` is a
// function that can be used to set up resources and configure the server before running.
func RunOrExit(configure func(c *ServerConfig) error, port int) {
	serverConfig := &ServerConfig{grpcServer: grpc.NewServer(grpc.UnaryInterceptor(internalInterceptor))}
	err := configure(serverConfig)
	if err != nil {
		log.Fatal(err)
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Server is running on port %v\n", port)
	if err := serverConfig.grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

type ServerConfig struct {
	grpcServer *grpc.Server
}

// InstallModule installs an Aether Module on the current server.
func (s *ServerConfig) InstallModule(module Module) {
	fmt.Printf("Installing module %v\n", module.Name())
	err := module.Configure(s)
	if err != nil {
		log.Fatalf("failed to install module: %v\n", err)
	}
}

// GetGRPCServer gets a reference to the current gRPC server. Useful for registering services.
func (s *ServerConfig) GetGRPCServer() *grpc.Server {
	return s.grpcServer
}

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
