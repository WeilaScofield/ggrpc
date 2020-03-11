package etcd

import (
	"github.com/coreos/etcd/clientv3"
	etcdnaming "github.com/coreos/etcd/clientv3/naming"
	"google.golang.org/grpc"
	"log"
	"time"
)

func NewClientV3(url string) *clientv3.Client {
	c, err := clientv3.New(clientv3.Config{Endpoints: []string{url}, DialTimeout: 2 * time.Second})
	if err != nil {
		log.Fatalf("etcd bootstrap failed:%v", err)
	}
	return c
}

func NewResolver(cli *clientv3.Client) *etcdnaming.GRPCResolver {
	return &etcdnaming.GRPCResolver{Client: cli}
}

func NewLB(cli *clientv3.Client) grpc.Balancer {
	r := &etcdnaming.GRPCResolver{Client: cli}
	return grpc.RoundRobin(r)
}
