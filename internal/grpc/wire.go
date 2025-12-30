package grpc

import (
	"github.com/google/wire"
	"github.com/kodaikumatani/grpc-golang/internal/grpc/greeter"
)

var Set = wire.NewSet(
	greeter.Set,
	NewRouter,
)
