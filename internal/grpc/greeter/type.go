package greeter

import (
	pb "github.com/kodaikumatani/grpc-golang/internal/pb/greeter"
)

type GreeterHandler struct {
	pb.UnimplementedGreeterServer
}

func NewGreeterHandler() *GreeterHandler {
	return &GreeterHandler{}
}
