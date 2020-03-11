package invokeAPI

import (
	"context"
	"fmt"
	pb "ggrpc/helloworld"
	"io"
	"log"
	"time"
)

// HelloInvoker provide methods to call grpc api
type HelloInvoker interface {
	ToSayHello()
	ToSendGifts()
}

type Invoker struct {
	Client pb.GreeterClient
}

func NewInvoker(client pb.GreeterClient) *Invoker {
	return &Invoker{Client: client}
}

func (i *Invoker) ToSayHello() {
	name := "vela"
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := i.Client.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}

func (i *Invoker) ToSendGifts() {
	giftsList := []string{"mac", "iphone", "girlfriend", "house"}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := i.Client.SendGifts(ctx)
	if err != nil {
		log.Fatalf("%v.SendGifts(_) = _, %v", i.Client, err)
	}

	waitc := make(chan struct{})
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("Failed to receive a note : %v", err)
			}
			fmt.Printf("Got gift:%v quantity:%v\n", in.Gifts, in.Quantity)
		}
	}()

	for _, v := range giftsList {
		if err := stream.Send(&pb.GiftsRequest{Quantity: int64(len(giftsList)), Gifts: v}); err != nil {
			log.Fatalf("Failed to send gift:%v quantity:%v", v, len(giftsList))
		}
	}

	stream.CloseSend()
	<-waitc
}
