package greeter

import (
	pb "github.com/kodaikumatani/grpc-golang/internal/pb/greeter"
	"google.golang.org/grpc"
)

type Router struct {
	greeterHandler *GreeterHandler
}

func NewRouter(greeterHandler *GreeterHandler) *Router {
	return &Router{
		greeterHandler: greeterHandler,
	}
}

func (r *Router) Routes(app *grpc.Server) *grpc.Server {
	pb.RegisterGreeterServer(app, &GreeterHandler{})

	return app
}
