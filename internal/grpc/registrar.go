package grpc

import (
	pb_follow "github.com/kodaikumatani/grpc-server-streaming-go/pkg/pb/follow"
	pb_notification "github.com/kodaikumatani/grpc-server-streaming-go/pkg/pb/notification"
	pb_recipe "github.com/kodaikumatani/grpc-server-streaming-go/pkg/pb/recipe"
	"google.golang.org/grpc"
)

type Registrar struct {
	recipeHandler       pb_recipe.RecipeServiceServer
	notificationHandler pb_notification.NotificationServiceServer
	followHandler       pb_follow.FollowServiceServer
}

func NewRegistrar(
	recipeHandler pb_recipe.RecipeServiceServer,
	notificationHandler pb_notification.NotificationServiceServer,
	followHandler pb_follow.FollowServiceServer,
) *Registrar {
	return &Registrar{
		recipeHandler:       recipeHandler,
		notificationHandler: notificationHandler,
		followHandler:       followHandler,
	}
}

func (r *Registrar) Register(app *grpc.Server) *grpc.Server {
	pb_recipe.RegisterRecipeServiceServer(app, r.recipeHandler)
	pb_notification.RegisterNotificationServiceServer(app, r.notificationHandler)
	pb_follow.RegisterFollowServiceServer(app, r.followHandler)

	return app
}
