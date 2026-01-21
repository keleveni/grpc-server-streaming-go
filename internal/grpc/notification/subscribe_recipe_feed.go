package notification

import (
	"github.com/kodaikumatani/grpc-golang/internal/app/recipe"
	pb "github.com/kodaikumatani/grpc-golang/internal/pb/notification"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (h *Handler) SubscribeRecipeFeed(
	in *pb.SubscribeRecipeFeedRequest,
	stream grpc.ServerStreamingServer[pb.RecipeEvent],
) error {
	userID := in.GetUserId()
	if userID == "" {
		return status.Error(codes.InvalidArgument, "user_id is required")
	}

	return h.subscribeFeed.StreamRecipes(stream.Context(), userID, func(r recipe.Recipe) error {
		event := &pb.RecipeEvent{
			EventType: "created",
			Recipe: &pb.RecipeInfo{
				Id:          r.ID.String(),
				UserId:      r.UserID,
				Title:       r.Title,
				Description: r.Description,
				CreatedAt:   timestamppb.New(r.CreatedAt),
			},
			EventTime: timestamppb.New(r.CreatedAt),
		}

		if err := stream.Send(event); err != nil {
			return err
		}

		return nil
	})
}
