package recipe

import (
	"github.com/kodaikumatani/grpc-golang/internal/app/recipe"
	pb "github.com/kodaikumatani/grpc-golang/internal/pb/recipe"
)

type handler struct {
	pb.UnimplementedRecipeServiceServer
	read   *recipe.ReadUseCase
	create *recipe.CreateUseCase
}

func NewHandler(
	read *recipe.ReadUseCase,
	create *recipe.CreateUseCase,
) pb.RecipeServiceServer {
	return &handler{
		read:   read,
		create: create,
	}
}
