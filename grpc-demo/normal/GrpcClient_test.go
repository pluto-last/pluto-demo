package normal

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"log"
	"testing"
	"time"

	pb "grpc-demo/pb"
)

func TestGrpcClient(t *testing.T) {
	// 连接服务器
	// 连接到server端，此处禁用安全传输
	conn, err := grpc.Dial("127.0.0.1:8972", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// 执行RPC调用并打印收到的响应数据
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// hello world 的请求
	//r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "world!!!"})

	// 带元数据的请求

	// 创建带有metadata的context
	ctx = metadata.AppendToOutgoingContext(ctx, "k1", "v1", "k1", "v2", "k2", "v3")

	// 添加一些 metadata 到 context (e.g. in an interceptor)
	ctx = metadata.AppendToOutgoingContext(ctx, "k3", "v4")

	// 发起普通RPC请求
	r, err := c.SomeRPC(ctx, &pb.HelloRequest{Name: "world!!!"})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}
