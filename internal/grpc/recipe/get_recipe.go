package recipe

import (
	"context"
	"errors"

	"github.com/kodaikumatani/grpc-golang/internal/app/recipe"
	pb "github.com/kodaikumatani/grpc-golang/internal/pb/recipe"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (h *Handler) GetRecipe(
	ctx context.Context,
	in *pb.GetRecipeRequest,
) (*pb.GetRecipeResponse, error) {
	result, err := h.read.Get(ctx, in.GetId())
	if err != nil {
		if errors.Is(err, recipe.ErrRecipeNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.GetRecipeResponse{Recipe: &pb.Recipe{
		Id:          result.ID.String(),
		UserId:      result.UserID,
		Title:       result.Title,
		Description: result.Description,
		CreatedAt:   timestamppb.New(result.CreatedAt),
		UpdatedAt:   timestamppb.New(result.UpdatedAt),
	}}, nil
}
