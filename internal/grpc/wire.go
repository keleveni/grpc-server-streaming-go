package grpc

import (
	"github.com/google/wire"
	"github.com/kodaikumatani/grpc-golang/internal/grpc/follow"
	"github.com/kodaikumatani/grpc-golang/internal/grpc/notification"
	"github.com/kodaikumatani/grpc-golang/internal/grpc/recipe"
)

var Set = wire.NewSet(
	recipe.NewHandler,
	notification.NewHandler,
	follow.NewHandler,
	NewRegistrar,
)
