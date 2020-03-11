package main

import (
	api "ggrpc/client/invokeAPI"
	"ggrpc/common"
	"ggrpc/etcd"
	pb "ggrpc/helloworld"
	"log"
	"time"

	"google.golang.org/grpc"
)

var config common.Config

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	config = common.NewConfig()
}

func main() {
	// Set up a connection to the server.
	cliv3 := etcd.NewClientV3(config.Etcd.Addrs)
	defer cliv3.Close()

	conn, err := grpc.Dial(config.Etcd.Resolver, grpc.WithBalancer(etcd.NewLB(cliv3)),
		grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(2*time.Second))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	// invoker.Run() is the only entry to call grpc service
	var invoker api.HelloInvoker = api.NewInvoker(c)
	invoker.ToSayHello()
	invoker.ToSendGifts()

	select {}
}
