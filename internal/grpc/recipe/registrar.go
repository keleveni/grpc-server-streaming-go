package recipe

import (
	pb "github.com/kodaikumatani/grpc-golang/internal/pb/recipe"
	"google.golang.org/grpc"
)

type Registrar struct {
	recipeHandler *RecipeHandler
}

func NewRegistrar(recipeHandler *RecipeHandler) *Registrar {
	return &Registrar{
		recipeHandler: recipeHandler,
	}
}

func (r *Registrar) Register(app *grpc.Server) *grpc.Server {
	pb.RegisterRecipeServiceServer(app, r.recipeHandler)

	return app
}
