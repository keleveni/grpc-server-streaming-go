package recipe

import (
	"context"

	"github.com/go-playground/validator/v10"
	pb "github.com/kodaikumatani/grpc-golang/internal/pb/recipe"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type createRecipeRequest struct {
	UserID      string `validate:"required,ulid"`
	Title       string `validate:"required"`
	Description string `validate:"required"`
}

func (h *Handler) CreateRecipe(
	ctx context.Context,
	in *pb.CreateRecipeRequest,
) (*pb.CreateRecipeResponse, error) {
	request := &createRecipeRequest{
		UserID:      in.GetUserId(),
		Title:       in.GetTitle(),
		Description: in.GetDescription(),
	}

	if err := validator.New().Struct(request); err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request parameters")
	}

	result, err := h.create.Invoke(ctx,
		request.UserID,
		request.Title,
		request.Description)
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
