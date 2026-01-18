package recipe

import (
	"github.com/google/wire"
	"github.com/kodaikumatani/grpc-golang/internal/app/recipe"
)

var Set = wire.NewSet(
	NewStorage,

	wire.Bind(new(recipe.Storage), new(*Storage)),
)
