package recipe

import (
	"context"

	pb "github.com/kodaikumatani/grpc-golang/internal/pb/recipe"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (h *Handler) CreateRecipe(
	ctx context.Context,
	in *pb.CreateRecipeRequest,
) (*pb.CreateRecipeResponse, error) {
	result, err := h.create.Invoke(ctx,
		in.GetUserId(),
		in.GetTitle(),
		in.GetDescription())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.CreateRecipeResponse{
		Recipe: &pb.Recipe{
			Id:          result.ID.String(),
			UserId:      result.UserID,
			Title:       result.Title,
			Description: result.Description,
			CreatedAt:   timestamppb.New(result.CreatedAt),
			UpdatedAt:   timestamppb.New(result.UpdatedAt),
		},
	}, nil
}
