package utils

import (
	"fmt"
	"log"

	"github.com/Sadham-Hussian/go-learnings/gofiber-grpc/proto"
	"google.golang.org/grpc"
)

// Client connection gRPC server
var Client proto.ProductServiceClient

func init() {
	conn, err := grpc.Dial("localhost:5050", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Unable to dial port 5050: %v", err)
	}
	fmt.Printf("%T", conn)
	Client = proto.NewProductServiceClient(conn)
}

// ConnectServer function to connect grpc server
func ConnectServer() {
	fmt.Printf("%T", Client)
}
