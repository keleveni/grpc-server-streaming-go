package recipe

import (
	"github.com/kodaikumatani/grpc-golang/internal/app/recipe"
	pb "github.com/kodaikumatani/grpc-golang/internal/pb/recipe"
)

type Handler struct {
	pb.UnimplementedRecipeServiceServer
	read   *recipe.ReadUseCase
	create *recipe.CreateUseCase
}

func NewHandler(
	read *recipe.ReadUseCase,
	create *recipe.CreateUseCase,
) *Handler {
	return &Handler{
		read:   read,
		create: create,
	}
}
