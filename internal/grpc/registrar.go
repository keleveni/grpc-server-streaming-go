package grpc

import (
	"github.com/kodaikumatani/grpc-golang/internal/grpc/follow"
	"github.com/kodaikumatani/grpc-golang/internal/grpc/notification"
	"github.com/kodaikumatani/grpc-golang/internal/grpc/recipe"
	"google.golang.org/grpc"
)

type Registrar struct {
	recipeHandler       *recipe.Handler
	notificationHandler *notification.Handler
	followHandler       *follow.Handler
}

func NewRegistrar(
	recipeHandler *recipe.Handler,
	notificationHandler *notification.Handler,
	followHandler *follow.Handler,
) *Registrar {
	return &Registrar{
		recipeHandler:       recipeHandler,
		notificationHandler: notificationHandler,
		followHandler:       followHandler,
	}
}

func (r *Registrar) Register(app *grpc.Server) *grpc.Server {
	r.recipeHandler.RegisterService(app)
	r.notificationHandler.RegisterService(app)
	r.followHandler.RegisterService(app)

	return app
}
