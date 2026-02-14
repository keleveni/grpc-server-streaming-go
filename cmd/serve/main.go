package main

import (
	"context"
	"flag"
	"fmt"
	"net"

	"github.com/cockroachdb/errors"
	"github.com/kodaikumatani/grpc-golang/internal/interceptor"
	"github.com/kodaikumatani/grpc-golang/internal/logger"
	"github.com/kodaikumatani/grpc-golang/internal/telemetry"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	port         = flag.Int("port", 50051, "The server port")
	otlpEndpoint = flag.String("otlp-endpoint", "localhost:4317", "OTLP gRPC endpoint")
	serviceName  = flag.String("service-name", "grpc-golang", "Service name for telemetry")
)

func run(ctx context.Context) error {
	flag.Parse()

	// Initialize OpenTelemetry
	cleanup, err := telemetry.InitTracer(ctx, *serviceName, *otlpEndpoint)
	if err != nil {
		return errors.Wrap(err, "failed to initialize telemetry")
	}
	defer func() {
		if err := cleanup(ctx); err != nil {
			log.Ctx(ctx).Error().Err(err).Msg("failed to cleanup telemetry")
		}
	}()

	log.Ctx(ctx).Info().Str("otlp_endpoint", *otlpEndpoint).Msg("OpenTelemetry initialized")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		return errors.Wrap(err, "failed to listen")
	}

	services, err := initializeServices(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to initialize services")
	}

	s := grpc.NewServer(
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
		grpc.ChainUnaryInterceptor(
			interceptor.RecoveryUnaryInterceptor(),
			interceptor.LoggingUnaryInterceptor(),
		),
		grpc.ChainStreamInterceptor(
			interceptor.RecoveryStreamInterceptor(),
			interceptor.LoggingStreamInterceptor(),
		),
	)
	s = services.Register(s)
	reflection.Register(s)

	log.Ctx(ctx).Info().Msgf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		return errors.Wrap(err, "failed to serve")
	}

	return nil
}

func main() {
	ctx := context.Background()
	ctx = logger.WithLogger(ctx)

	if err := run(ctx); err != nil {
		log.Ctx(ctx).Fatal().Err(err).Send()
	}
}
