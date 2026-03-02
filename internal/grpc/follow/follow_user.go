package follow

import (
	"context"

	"github.com/go-playground/validator/v10"
	pb "github.com/kodaikumatani/grpc-golang/internal/pb/follow"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type followUserRequest struct {
	FolloweeID string `validate:"required,ulid,nefield=FollowerID"`
	FollowerID string `validate:"required,ulid"`
}

func (h *handler) FollowUser(
	ctx context.Context,
	in *pb.FollowUserRequest,
) (*pb.FollowUserResponse, error) {
	request := &followUserRequest{
		FolloweeID: in.GetFolloweeId(),
		FollowerID: in.GetFollowerId(),
	}

	if err := validator.New().Struct(request); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if err := h.follow.Follow(ctx, request.FollowerID, request.FolloweeID); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.FollowUserResponse{
		Follow: &pb.Follow{
			FollowerId: request.FollowerID,
			FolloweeId: request.FolloweeID,
			CreatedAt:  timestamppb.Now(),
		},
	}, nil
}
