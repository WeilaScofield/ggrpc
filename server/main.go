package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	etcdnaming "github.com/coreos/etcd/clientv3/naming"
	"google.golang.org/grpc/naming"
	"log"
	"net"
	"os"

	"ggrpc/common"
	"ggrpc/etcd"
	pb "ggrpc/helloworld"
	api "ggrpc/server/grpcAPI"
	"google.golang.org/grpc"
)

var (
	cliv3   *clientv3.Client
	resolver *etcdnaming.GRPCResolver
	config  = common.NewConfig()
	port    = ":50053"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile) //选择日志格式

	if len(os.Args) > 1 {
		port = ":" + os.Args[1]
	}
	addr := fmt.Sprintf("localhost%v", port)

	cliv3 = etcd.NewClientV3(config.Etcd.Addrs)
	resolver = etcd.NewResolver(cliv3)
	if err := resolver.Update(context.TODO(), config.Etcd.Resolver,
		naming.Update{Op: naming.Add, Addr: addr, Metadata: "..."}); err != nil {
		log.Fatalf("register to etcd failed: %v", err)
	}

}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &api.Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
