package follow

import (
	"github.com/kodaikumatani/grpc-golang/internal/app/user"
	pb "github.com/kodaikumatani/grpc-golang/internal/pb/follow"
)

type handler struct {
	pb.UnimplementedFollowServiceServer
	follow *user.FollowUseCase
}

func NewHandler(
	follow *user.FollowUseCase,
) pb.FollowServiceServer {
	return &handler{
		follow: follow,
	}
}
