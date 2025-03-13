package user_cli

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	user "http-gateway/pb/user"
)

func CallSayHello() {
	// 连接到服务器
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("连接失败: %v", err)
	}
	defer conn.Close()

	// 创建客户端
	client := user.NewHelloServiceClient(conn)

	// 设置超时上下文
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 调用RPC方法
	resp, err := client.SayHello(ctx, &user.HelloRequest{Name: "世界"})
	if err != nil {
		log.Fatalf("调用失败: %v", err)
	}

	log.Printf("响应: %s", resp.GetMessage())
}
