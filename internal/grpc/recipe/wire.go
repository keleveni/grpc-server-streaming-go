package recipe

import (
	"github.com/google/wire"
	apprecipe "github.com/kodaikumatani/grpc-golang/internal/app/recipe"
)

var handlerSet = wire.NewSet(
	NewRecipeHandler,
)

var Set = wire.NewSet(
	handlerSet,
	NewRegistrar,
	apprecipe.Set,
)
