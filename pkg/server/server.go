package server

import (
	"context"
	"google.golang.org/grpc"
)

type Service interface {
	Service() interface{}
	ServiceDesc() *grpc.ServiceDesc
}

type Server struct {
	service     interface{}
	serviceDesc *grpc.ServiceDesc
}

func (s *Server) RegisterService(service Service) {
	s.service = service.Service()
	s.serviceDesc = service.ServiceDesc()
}

func (s *Server) RunOrExit(ctx context.Context, port int) {

}
