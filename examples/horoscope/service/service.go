package horoscope

import (
	"context"
	pb "github.com/nthnluu/aether/examples/horoscope/pb/out"
	"github.com/nthnluu/aether/pkg"
)

type service struct {
	horoscopes Repository
	*pb.UnimplementedHoroscopeServiceServer
}

func (s *service) GetDailyHoroscope(ctx context.Context, getDailyHoroscopeRequest *pb.GetDailyHoroscopeRequest) (*pb.GetDailyHoroscopeResponse, error) {
	return nil, aether.NotYetImplementedError("Implement me!")
}

func (s *service) GetHoroscope(ctx context.Context, createLinkRequest *pb.GetHoroscopeRequest) (*pb.GetHoroscopeResponse, error) {
	return nil, aether.NotYetImplementedError("Implement me!")
}

func (s *service) SuggestFortune(ctx context.Context, suggestFortuneRequest *pb.SuggestFortuneRequest) (*pb.SuggestFortuneResponse, error) {
	return nil, aether.NotYetImplementedError("Implement me!")
}

func CreateService(horoscopeRepository Repository) *service {
	return &service{
		horoscopes: horoscopeRepository,
	}
}
