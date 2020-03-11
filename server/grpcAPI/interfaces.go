package grpcAPI

import (
	"context"
	pb "ggrpc/helloworld"
	"io"
	"log"
)

// server is used to implement helloworld.GreeterServer.
type Server struct {
}

func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v\n", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func (s *Server) SendGifts(stream pb.Greeter_SendGiftsServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		in.Gifts = "for vela:" + in.Gifts
		log.Println(in.Gifts)
		err = stream.Send(&pb.GiftsReply{Quantity: int64(len(in.Gifts)), Gifts: in.Gifts})
		if err != nil {
			return err
		}
	}

}
