package notification

import (
	"github.com/go-playground/validator/v10"
	"github.com/kodaikumatani/grpc-server-streaming-go/internal/app/recipe"
	pb "github.com/kodaikumatani/grpc-server-streaming-go/pkg/pb/notification"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type subscribeRecipeFeedRequest struct {
	UserID string `validate:"required,ulid"`
}

func (h *handler) SubscribeRecipeFeed(
	in *pb.SubscribeRecipeFeedRequest,
	stream grpc.ServerStreamingServer[pb.RecipeEvent],
) error {
	request := &subscribeRecipeFeedRequest{
		UserID: in.GetUserId(),
	}

	if err := validator.New().Struct(request); err != nil {
		return status.Error(codes.InvalidArgument, "invalid request parameters")
	}

	return h.subscribeFeed.StreamRecipes(stream.Context(), request.UserID,
		func(r recipe.Recipe) error {
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

			return stream.Send(event)
		},
	)
}
