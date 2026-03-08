package db

import (
	"github.com/google/wire"
	"github.com/kodaikumatani/grpc-server-streaming-go/internal/db/follow"
	"github.com/kodaikumatani/grpc-server-streaming-go/internal/db/recipe"
)

var Set = wire.NewSet(
	recipe.Set,
	follow.Set,
)
