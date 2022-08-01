package main

import (
	"contrib.go.opencensus.io/exporter/stackdriver"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/log"
	"github.com/nthnluu/aether/cmd/golink"
	proto "github.com/nthnluu/aether/pb/out"
	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/trace"
	"google.golang.org/grpc"
	stdlog "log"
	"net"
	"os"
)

const (
	DevEnvironment        = "dev"
	StagingEnvironment    = "staging"
	ProductionEnvironment = "production"
)

const (
	defaultPort              = "8080"
	defaultRoutingServiceURL = "http://localhost:7878"
	defaultEnvironment       = DevEnvironment
)

func main() {
	var (
		grpcAddr    = "localhost:8082"
		environment = defaultEnvironment
	)

	// Set up the logger
	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	stdlog.SetOutput(log.NewStdlibAdapter(logger))

	if err := view.Register(ocgrpc.DefaultServerViews...); err != nil {
		stdlog.Fatalf("Failed to register ocgrpc server views: %v", err)
	}

	if environment != DevEnvironment {
		// Create and register the stackdriver exporter for exporting metrics to Google Cloud.
		sd, err := stackdriver.NewExporter(stackdriver.Options{
			ProjectID:    "census-demos", // Insert your projectID here
			MetricPrefix: "ocgrpctutorial",
		})
		if err != nil {
			stdlog.Fatalf("Failed to create Stackdriver exporter: %v", err)
		}
		defer sd.Flush()

		trace.RegisterExporter(sd)
		err = sd.StartMetricsExporter()
		if err != nil {
			stdlog.Fatalf("Failed to start metrics exporter: %v", err)
		}
		defer sd.StopMetricsExporter()
	}

	// Create services
	var goLinkService golink.Service
	goLinkService = golink.NewService()
	goLinkService = golink.NewInstrumentingService(log.With(logger, "component", "golink"), goLinkService)

	// Set up server
	grpcListener, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		logger.Log("transport", "gRPC", "during", "Listen", "err", err)
		os.Exit(1)
	}

	baseServer := grpc.NewServer(grpc.UnaryInterceptor(kitgrpc.Interceptor), grpc.StatsHandler(&ocgrpc.ServerHandler{}))
	proto.RegisterGoLinkServiceServer(baseServer, golink.NewGRPCServer(goLinkService))
	logger.Log("transport", "gRPC", "addr", grpcAddr, "msg", "listening")
	err = baseServer.Serve(grpcListener)

	if err != nil {
		logger.Log("transport", "gRPC", "during", "Listen", "err", err)
		os.Exit(1)
	}
}

// envString attempts to get env from the environment, and returns fallback if it's not defined.
func envString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}
