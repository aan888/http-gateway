package user

import (
	"context"
	"log"
	"net"

	user "http-gateway/pb/user"

	"google.golang.org/grpc"
)

// 服务实现
type server struct {
	user.UnimplementedHelloServiceServer
}

// 实现SayHello方法
func (s *server) SayHello(ctx context.Context, req *user.HelloRequest) (*user.HelloResponse, error) {
	log.Printf("收到请求: %v", req.GetName())
	return &user.HelloResponse{Message: "你好, " + req.GetName()}, nil
}

func StartUserServer() {
	// 监听指定端口
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("监听失败: %v", err)
	}

	// 创建gRPC服务器
	s := grpc.NewServer()

	// 注册服务
	user.RegisterHelloServiceServer(s, &server{})

	log.Println("服务器启动，监听端口: 50051")

	// 启动服务
	if err := s.Serve(lis); err != nil {
		log.Fatalf("服务启动失败: %v", err)
	}
}
