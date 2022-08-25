package etcd_register

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
	pb "grpc-demo/pb"
	"log"
	"testing"
	"time"
)

func TestGrpcClient(t *testing.T) {

	// 注册自定义ETCD解析器
	etcdResolverBuilder := NewEtcdResolverBuilder()
	resolver.Register(etcdResolverBuilder)

	// 使用自带的DNS解析器和负载均衡实现方式
	conn, err := grpc.Dial("etcd:///", grpc.WithBalancerName(roundrobin.Name), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	//// 执行RPC调用并打印收到的响应数据
	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()

	for {
		time.Sleep(2 * time.Second)
		// 发起普通RPC请求
		r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: "world!!!"})

		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Greeting: %s", r.Message)
	}

}
