package grpc

import (
	"github.com/kodaikumatani/grpc-golang/internal/grpc/greeter"
	"google.golang.org/grpc"
)

type Router struct {
	greeterRouter *greeter.Router
}

func NewRouter(greeterRouter *greeter.Router) *Router {
	return &Router{
		greeterRouter: greeterRouter,
	}
}

func (r *Router) Routes(app *grpc.Server) *grpc.Server {
	r.greeterRouter.Routes(app)

	return app
}
