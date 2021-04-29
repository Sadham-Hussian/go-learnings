package main

import (
	"fmt"
	"log"
	"net"

	"github.com/Sadham-Hussian/go-learnings/gofiber-grpc/proto"
	"github.com/Sadham-Hussian/go-learnings/gofiber-grpc/server/product"
	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":5050")
	if err != nil {
		log.Fatalf("Unable to listen on port :5050: %v", err)
	}

	fmt.Println("listening")

	grpcServer := grpc.NewServer()
	proto.RegisterProductServiceServer(grpcServer, &product.ProductServiceServer{})

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %s", err)
	}
}
