package main

import (
	"article_GraphQL_gRPC/repo/injector"
	"article_GraphQL_gRPC/repo/pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("failed to listen: %v\n", err)
	}
	defer lis.Close()

	service := injector.InjectArticleService()
	server := grpc.NewServer()

	pb.RegisterArticleServiceServer(server, service)

	log.Println("Listening on port 50051...")
	if err := server.Serve(lis); err != nil {
		log.Fatal("failed to serve: %v", err)
	}

}
