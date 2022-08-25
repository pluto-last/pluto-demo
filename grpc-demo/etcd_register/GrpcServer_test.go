package etcd_register

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
	lis, err := net.Listen("tcp", "127.0.0.1:8972")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()                  // 创建gRPC服务器
	pb.RegisterGreeterServer(s, &server{}) // 在gRPC服务端注册服务,并且注入我们自己的实现类

	reflection.Register(s) //在给定的gRPC服务器上注册服务器反射服务

	// 创建一个注册器
	etcdRegister, err := NewEtcdRegister()
	if err != nil {
		log.Println("创建etcd注册器失败", err)
		return
	}

	defer etcdRegister.Close()

	// 服务名字
	serviceName := "pluto-service-1"
	addr := "127.0.0.1:8972"

	// 注册服务
	err = etcdRegister.RegisterServer("/etcd/"+serviceName, addr, 5)
	if err != nil {
		log.Printf("register error %v \n", err)
		return
	}

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
