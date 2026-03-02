package notification

import (
	"github.com/kodaikumatani/grpc-golang/internal/app/notification"
	pb "github.com/kodaikumatani/grpc-golang/internal/pb/notification"
)

type handler struct {
	pb.UnimplementedNotificationServiceServer
	subscribeFeed *notification.SubscribeFeedUseCase
}

func NewHandler(subscribeFeed *notification.SubscribeFeedUseCase) pb.NotificationServiceServer {
	return &handler{
		subscribeFeed: subscribeFeed,
	}
}
