package recipe

import (
	"github.com/kodaikumatani/grpc-golang/internal/app/recipe"
	pb "github.com/kodaikumatani/grpc-golang/internal/pb/recipe"
)

type RecipeHandler struct {
	pb.UnimplementedRecipeServiceServer
	read   *recipe.ReadUseCase
	create *recipe.CreateUseCase
}

func NewRecipeHandler(
	read *recipe.ReadUseCase,
	create *recipe.CreateUseCase,
) *RecipeHandler {
	return &RecipeHandler{
		read:   read,
		create: create,
	}
}
