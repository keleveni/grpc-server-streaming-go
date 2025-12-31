package recipe

import (
	"github.com/kodaikumatani/grpc-golang/internal/app/recipe"
	pb "github.com/kodaikumatani/grpc-golang/internal/pb/recipe"
)

type RecipeHandler struct {
	pb.UnimplementedRecipeServiceServer
	read *recipe.ReadUseCase
}

func NewRecipeHandler(read *recipe.ReadUseCase) *RecipeHandler {
	return &RecipeHandler{
		read: read,
	}
}
