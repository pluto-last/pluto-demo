module grpc-demo

go 1.16

require (
	go.etcd.io/etcd/client/v3 v3.5.4
	go.uber.org/zap v1.23.0 // indirect
	golang.org/x/net v0.0.0-20220822230855-b0a4917ee28c
	google.golang.org/grpc v1.49.0
	google.golang.org/protobuf v1.27.1

)

replace google.golang.org/grpc v1.49.0 => google.golang.org/grpc v1.45.0
