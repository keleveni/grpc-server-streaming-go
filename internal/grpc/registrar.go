package grpc

import (
	"github.com/kodaikumatani/grpc-golang/internal/grpc/notification"
	"github.com/kodaikumatani/grpc-golang/internal/grpc/recipe"
	"google.golang.org/grpc"
)

type Registrar struct {
	recipeHandler       *recipe.Handler
	notificationHandler *notification.Handler
}

func NewRegistrar(
	recipeHandler *recipe.Handler,
	notificationHandler *notification.Handler,
) *Registrar {
	return &Registrar{
		recipeHandler:       recipeHandler,
		notificationHandler: notificationHandler,
	}
}

func (r *Registrar) Register(app *grpc.Server) *grpc.Server {
	r.recipeHandler.RegisterService(app)
	r.notificationHandler.RegisterService(app)

	return app
}
