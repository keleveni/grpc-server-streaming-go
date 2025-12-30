package greeter

import (
	"github.com/google/wire"
)

var handlerSet = wire.NewSet(
	NewGreeterHandler,
)

var Set = wire.NewSet(
	handlerSet,
	NewRouter,
)

