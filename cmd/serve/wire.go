//go:build wireinject
// +build wireinject

package main

import (
	"context"

	"github.com/google/wire"
	"github.com/kodaikumatani/grpc-golang/internal"
	"github.com/kodaikumatani/grpc-golang/internal/grpc"
)

type services struct {
	*grpc.Registrar
}

var set = wire.NewSet(
	internal.Set,
	wire.Struct(new(services), "*"),
)

func initializeServices(context.Context) (*services, error) {
	panic(wire.Build(set))
}
