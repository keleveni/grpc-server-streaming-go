package db

import (
	"github.com/google/wire"
	"github.com/kodaikumatani/grpc-golang/internal/db/follow"
	"github.com/kodaikumatani/grpc-golang/internal/db/recipe"
)

var Set = wire.NewSet(
	recipe.Set,
	follow.Set,
)
