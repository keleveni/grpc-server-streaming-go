package notification

import (
	"github.com/kodaikumatani/grpc-server-streaming-go/internal/app/notification"
	pb "github.com/kodaikumatani/grpc-server-streaming-go/internal/pb/notification"
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
