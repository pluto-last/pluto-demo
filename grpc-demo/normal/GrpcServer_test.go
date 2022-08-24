package normal

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	pb "grpc-demo/pb"
	"log"
	"net"
	"testing"
)

type server struct{}

func TestGrpcServer(t *testing.T) {

	// 监听本地的8972端口
	lis, err := net.Listen("tcp", ":8972")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}
	s := grpc.NewServer(grpc.UnaryInterceptor(LoggingInterceptor)) // 创建gRPC服务器
	pb.RegisterGreeterServer(s, &server{})                         // 在gRPC服务端注册服务,并且注入我们自己的实现类

	reflection.Register(s) //在给定的gRPC服务器上注册服务器反射服务
	// Serve方法在lis上接受传入连接，为每个连接创建一个ServerTransport和server的goroutine。
	// 该goroutine读取gRPC请求，然后调用已注册的处理程序来响应它们。
	err = s.Serve(lis)
	if err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}

}

//拦截器 - 打印日志
//拦截器 - 打印日志
func LoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	fmt.Printf("gRPC method: %s, %v", info.FullMethod, req)
	resp, err := handler(ctx, req)
	fmt.Printf("gRPC method: %s, %v", info.FullMethod, resp)
	return resp, err
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func (s *server) SomeRPC(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	log.Println("拿到的元数据是", md)

	return &pb.HelloReply{Message: "测试拿元数据"}, nil
}
