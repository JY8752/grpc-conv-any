package main

import (
	"fmt"
	"log"
	"net"

	"github.com/JY8752/grpc-conv-any/example/service"
	"github.com/JY8752/grpc-conv-any/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	fmt.Println("Hello world")

	port := 6565
	listenPort, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	//サーバーにserviceを登録する
	server := grpc.NewServer()
	pb.RegisterHelloServer(server, &service.HelloService{})

	//リフレクション有効
	reflection.Register(server)

	//サーバー起動
	server.Serve(listenPort)
}
