package main

import (
	"flag"
	horoscope "github.com/nthnluu/aether/examples/horoscope/service"
	aether "github.com/nthnluu/aether/pkg"
)

var (
	port = flag.Int("port", 9999, "Port for serving gRPC requests")
)

func configure(c *aether.ServerConfig) error {
	c.InstallModule(horoscope.Module(horoscope.NewRepository()))
	return nil
}

func main() {
	flag.Parse()

	// Run the server.
	aether.RunOrExit(configure, *port)
}
