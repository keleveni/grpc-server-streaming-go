package follow

import (
	"github.com/kodaikumatani/grpc-golang/internal/app/user"
	pb "github.com/kodaikumatani/grpc-golang/internal/pb/follow"
)

type Handler struct {
	pb.UnimplementedFollowServiceServer
	follow *user.FollowUseCase
}

func NewHandler(
	follow *user.FollowUseCase,
) pb.FollowServiceServer {
	return &Handler{
		follow: follow,
	}
}
