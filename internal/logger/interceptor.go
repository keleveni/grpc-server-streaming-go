package logger

import (
	"context"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func PanicRecoveryHandler(ctx context.Context, p any) error {
	log.Ctx(ctx).Error().
		Interface("panic", p).
		Msg("panic recorvered in gRPC handler")

	return status.Error(codes.Internal, "internal server error")
}

func LoggingAdapter() logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		logger := log.Ctx(ctx)

		var event *zerolog.Event
		switch lvl {
		case logging.LevelDebug:
			event = logger.Debug()
		case logging.LevelInfo:
			event = logger.Info()
		case logging.LevelWarn:
			event = logger.Warn()
		case logging.LevelError:
			event = logger.Error()
		default:
			event = logger.Info()
		}

		for i := 0; i < len(fields); i += 2 {
			if i+1 < len(fields) {
				key, ok := fields[i].(string)
				if ok {
					event = event.Interface(key, fields[i+1])
				}
			}
		}

		event.Msg(msg)
	})
}
