package grpc

import (
	"github.com/kodaikumatani/grpc-golang/internal/grpc/recipe"
	pb "github.com/kodaikumatani/grpc-golang/internal/pb/recipe"
	"google.golang.org/grpc"
)

type Registrar struct {
	recipeHandler *recipe.Handler
}

func NewRegistrar(
	recipeHandler *recipe.Handler,
) *Registrar {
	return &Registrar{
		recipeHandler: recipeHandler,
	}
}

func (r *Registrar) Register(app *grpc.Server) *grpc.Server {
	pb.RegisterRecipeServiceServer(app, r.recipeHandler)

	return app
}
