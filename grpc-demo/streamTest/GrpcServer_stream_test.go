package streamTest

import (
	"context"
	"fmt"
	"io"
	"net"
	"strconv"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "grpc-demo/pb"
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
	pb.RegisterBaseServiceServer(s, &server{})                     // 在gRPC服务端注册服务

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
func LoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	fmt.Printf("gRPC method: %s, %v", info.FullMethod, req)
	resp, err := handler(ctx, req)
	fmt.Printf("gRPC method: %s, %v", info.FullMethod, resp)
	return resp, err
}

type wrappedStream struct {
	grpc.ServerStream
}

func newWrappedStream(s grpc.ServerStream) grpc.ServerStream {
	return &wrappedStream{s}
}

func (w *wrappedStream) RecvMsg(m interface{}) error {
	fmt.Printf("Receive a message (Type: %T) at %s", m, time.Now().Format(time.RFC3339))
	return w.ServerStream.RecvMsg(m)
}

func (w *wrappedStream) SendMsg(m interface{}) error {
	fmt.Printf("Send a message (Type: %T) at %v", m, time.Now().Format(time.RFC3339))
	return w.ServerStream.SendMsg(m)
}

//发消息前后流式调用拦截器
func SendMsgCheckStreamServerInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo,
	handler grpc.StreamHandler) error {
	fmt.Printf("gRPC method: %s,", info.FullMethod)
	err := handler(srv, newWrappedStream(ss))
	fmt.Printf("gRPC method: %s", info.FullMethod)
	return err
}

//普通前置流式调用拦截器
func StreamServerInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo,
	handler grpc.StreamHandler) error {
	fmt.Printf("gRPC method: %s,", info.FullMethod)
	err := handler(srv, ss)
	fmt.Printf("gRPC method: %s", info.FullMethod)
	return err
}

//sum案例--客户端流式处理
func (*server) Sum(req pb.BaseService_SumServer) (err error) {
	var sum int64 = 0
	for {
		reqObj, err := req.Recv()
		if err == io.EOF {
			fmt.Printf("Recv Sum err: %v", err)
			req.SendAndClose(&pb.SumResponse{Result: sum})
			return nil
		} else if err == nil {
			fmt.Printf("get client request param = %v", reqObj.Num)
			sum += reqObj.Num
		} else {
			return err
		}
	}
}

// 服务端流式处理
func (s *server) ServerStream(in *pb.StreamRequest, stream pb.BaseService_ServerStreamServer) error {
	input := in.Input
	for _, s := range input {
		stream.Send(&pb.StreamResponse{Output: string(s)})
	}
	return nil
}

// 客户端流式响应 - 服务端逻辑
func (s *server) ClientStream(stream pb.BaseService_ClientStreamServer) error {
	output := ""
	for {
		r, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.StreamResponse{Output: output})
		}
		if err != nil {
			fmt.Println(err)
		}
		output += r.Input
	}
}

// 双向流式处理
func (s *server) Streaming(stream pb.BaseService_StreamingServer) error {
	for n := 0; ; {
		res, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		v, _ := strconv.Atoi(res.Input)
		n += v
		stream.Send(&pb.StreamResponse{Output: strconv.Itoa(n)})
	}
}
