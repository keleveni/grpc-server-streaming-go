package grpc

import (
	"github.com/google/wire"
	"github.com/kodaikumatani/grpc-server-streaming-go/internal/grpc/follow"
	"github.com/kodaikumatani/grpc-server-streaming-go/internal/grpc/notification"
	"github.com/kodaikumatani/grpc-server-streaming-go/internal/grpc/recipe"
)

var Set = wire.NewSet(
	recipe.NewHandler,
	notification.NewHandler,
	follow.NewHandler,
	NewRegistrar,
)
