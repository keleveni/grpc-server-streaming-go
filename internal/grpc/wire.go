package grpc

import (
	"github.com/google/wire"
	"github.com/kodaikumatani/grpc-golang/internal/grpc/recipe"
)

var Set = wire.NewSet(
	recipe.NewHandler,
	wire.Struct(new(Registrar), "*"),
)
