package main

import (
	"flag"
	pb "github.com/nthnluu/aether/examples/horoscope/pb/out"
	horoscope "github.com/nthnluu/aether/examples/horoscope/service"
	aether "github.com/nthnluu/aether/pkg"
)

var (
	port = flag.Int("port", 9999, "Port for serving gRPC requests")
)

func configure(c *aether.ServerConfig) error {
	horoscopeService := horoscope.CreateService(horoscope.NewRepository())
	pb.RegisterHoroscopeServiceServer(c.GetGRPCServer(), horoscopeService)
	return nil
}

func main() {
	flag.Parse()

	// Run the server.
	aether.RunOrExit(configure, *port)
}
