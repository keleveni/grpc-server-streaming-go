package grpc

import (
	"github.com/kodaikumatani/grpc-golang/internal/grpc/recipe"
	"google.golang.org/grpc"
)

type Registrar struct {
	recipeRegistrar *recipe.Registrar
}

func NewRegistrar(
	recipeRegistrar *recipe.Registrar,
) *Registrar {
	return &Registrar{
		recipeRegistrar: recipeRegistrar,
	}
}

func (r *Registrar) Register(app *grpc.Server) *grpc.Server {
	r.recipeRegistrar.Register(app)

	return app
}
