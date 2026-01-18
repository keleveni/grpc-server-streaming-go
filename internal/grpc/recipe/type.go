package recipe

import (
	"github.com/kodaikumatani/grpc-golang/internal/app/recipe"
	pb "github.com/kodaikumatani/grpc-golang/internal/pb/recipe"
	"google.golang.org/grpc"
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

func (r *Handler) RegisterService(app *grpc.Server) {
	pb.RegisterRecipeServiceServer(app, r)
}
