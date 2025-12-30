package internal

import (
	"github.com/google/wire"
	"github.com/kodaikumatani/grpc-golang/internal/grpc"
)

var Set = wire.NewSet(
	grpc.Set,
)
