package horoscope

import (
	"context"
	pb "github.com/nthnluu/aether/examples/horoscope/pb/out"
	"github.com/nthnluu/aether/pkg"
)

// service is a struct that implements the interface generated the protocol buffer service definition.
type service struct {
	horoscopes Repository
	*pb.UnimplementedHoroscopeServiceServer
}

// GetDailyHoroscope gets the daily horoscope for the given zodiac sign.
func (s *service) GetDailyHoroscope(ctx context.Context, getDailyHoroscopeRequest *pb.GetDailyHoroscopeRequest) (*pb.GetDailyHoroscopeResponse, error) {
	return nil, aether.NotYetImplementedError("Implement me!")
}

// GetHoroscope gets the horoscope for the given zodiac sign and date.
func (s *service) GetHoroscope(ctx context.Context, createLinkRequest *pb.GetHoroscopeRequest) (*pb.GetHoroscopeResponse, error) {
	return nil, aether.NotYetImplementedError("Implement me!")
}

// SuggestFortune suggests a fortune.
func (s *service) SuggestFortune(ctx context.Context, suggestFortuneRequest *pb.SuggestFortuneRequest) (*pb.SuggestFortuneResponse, error) {
	return nil, aether.NotYetImplementedError("Implement me!")
}

// module is a struct that implements Module. This module represents the Horoscope service.
type module struct {
	service *service
}

// Name is a method that returns a human-readable name for the module.
func (m *module) Name() string {
	return "Horoscope service"
}

// Configure is a function that is called with a `ServerConfig`. It can be used to install interceptors, register
// gRPC services, and more.
func (m *module) Configure(c *aether.ServerConfig) error {
	pb.RegisterHoroscopeServiceServer(c.GetGRPCServer(), m.service)
	return nil
}

// Module is a function that creates an instance of the `module` struct.
func Module(horoscopeRepository Repository) *module {
	return &module{
		service: &service{
			horoscopes: horoscopeRepository,
		},
	}
}
