package notification

import (
	"github.com/kodaikumatani/grpc-golang/internal/app/notification"
	pb "github.com/kodaikumatani/grpc-golang/internal/pb/notification"
)

type Handler struct {
	pb.UnimplementedNotificationServiceServer
	subscribeFeed *notification.SubscribeFeedUseCase
}

func NewHandler(subscribeFeed *notification.SubscribeFeedUseCase) pb.NotificationServiceServer {
	return &Handler{
		subscribeFeed: subscribeFeed,
	}
}
