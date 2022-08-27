package horoscope

import (
	"flag"
	pb "github.com/nthnluu/aether/examples/horoscope/pb/out"
	horoscope "github.com/nthnluu/aether/examples/horoscope/service"
	aether "github.com/nthnluu/aether/pkg"
)

var (
	port = flag.Int("port", 9999, "Port for serving gRPC requests")
)

func main() {
	flag.Parse()

	// Create the service
	horoscopeService := horoscope.CreateService(horoscope.NewRepository())

	// Create the gRPC server and register your service.
	grpcServer := aether.CreateServer()
	pb.RegisterHoroscopeServiceServer(grpcServer, horoscopeService)

	// Run the server.
	aether.RunOrExit(grpcServer, *port, func(c *aether.ServerConfig) {})
}
