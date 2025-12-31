package internal

import (
	"github.com/google/wire"
	"github.com/kodaikumatani/grpc-golang/internal/grpc"
	"github.com/kodaikumatani/grpc-golang/internal/orm"
)

var Set = wire.NewSet(
	grpc.Set,
	orm.Set,
)
