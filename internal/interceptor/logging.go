package interceptor

import (
	"context"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
)

func LoggingUnaryInterceptor() grpc.UnaryServerInterceptor {
	opts := []logging.Option{
		withTraceID(),
	}

	return logging.UnaryServerInterceptor(
		loggingAdapter(),
		opts...,
	)
}

func LoggingStreamInterceptor() grpc.StreamServerInterceptor {
	return logging.StreamServerInterceptor(
		loggingAdapter(),
	)
}

func loggingAdapter() logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		logger := log.Ctx(ctx).With().Fields(fields).Logger()

		if span := trace.SpanContextFromContext(ctx); span.IsSampled() {
			logger = logger.With().
				Str("traceID", span.TraceID().String()).
				Str("spanID", span.SpanID().String()).
				Logger()
		}

		switch lvl {
		case logging.LevelDebug:
			logger.Debug().Msg(msg)
		case logging.LevelInfo:
			logger.Info().Msg(msg)
		case logging.LevelWarn:
			logger.Warn().Msg(msg)
		case logging.LevelError:
			logger.Error().Msg(msg)
		default:
			logger.Info().Msg(msg)
		}
	})
}

func withTraceID() logging.Option {
	return logging.WithFieldsFromContext(func(ctx context.Context) logging.Fields {
		if span := trace.SpanContextFromContext(ctx); span.IsSampled() {
			return logging.Fields{"traceID", span.TraceID().String()}
		}
		return nil
	})
}
