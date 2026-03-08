package internal

import (
	"github.com/google/wire"
	"github.com/kodaikumatani/grpc-server-streaming-go/internal/app"
	"github.com/kodaikumatani/grpc-server-streaming-go/internal/db"
	"github.com/kodaikumatani/grpc-server-streaming-go/internal/grpc"
	"github.com/kodaikumatani/grpc-server-streaming-go/internal/orm"
)

var Set = wire.NewSet(
	grpc.Set,
	app.Set,
	orm.Set,
	db.Set,
)
