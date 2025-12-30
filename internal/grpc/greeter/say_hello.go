package greeter

import (
	"context"
	"log"

	pb "github.com/kodaikumatani/grpc-golang/internal/pb/greeter"
)

func (h *GreeterHandler) SayHello(
	ctx context.Context,
	in *pb.HelloRequest,
) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}
