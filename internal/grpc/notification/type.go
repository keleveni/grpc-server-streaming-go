package notification

import (
	"github.com/kodaikumatani/grpc-golang/internal/app/notification"
	pb "github.com/kodaikumatani/grpc-golang/internal/pb/notification"
	"google.golang.org/grpc"
)

type Handler struct {
	pb.UnimplementedNotificationServiceServer
	subscribeFeed *notification.SubscribeFeedUseCase
}

func NewHandler(subscribeFeed *notification.SubscribeFeedUseCase) *Handler {
	return &Handler{
		subscribeFeed: subscribeFeed,
	}
}

func (r *Handler) RegisterService(app *grpc.Server) {
	pb.RegisterNotificationServiceServer(app, r)
}
