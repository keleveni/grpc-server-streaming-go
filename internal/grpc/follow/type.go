package follow

import (
	"github.com/kodaikumatani/grpc-golang/internal/app/user"
	pb "github.com/kodaikumatani/grpc-golang/internal/pb/follow"
	"google.golang.org/grpc"
)

type Handler struct {
	pb.UnimplementedFollowServiceServer
	follow *user.FollowUseCase
}

func NewHandler(
	follow *user.FollowUseCase,
) *Handler {
	return &Handler{
		follow: follow,
	}
}

func (h *Handler) RegisterService(app *grpc.Server) {
	pb.RegisterFollowServiceServer(app, h)
}
