package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/ParasJain0307/grpc-project/grpc-server/api"
	"github.com/ParasJain0307/grpc-project/grpc-server/internal/database"
	"github.com/ParasJain0307/grpc-project/grpc-server/internal/service"
	"google.golang.org/grpc"
)

func main() {
	//Initialize the simulated Database
	db, _ := database.NewDatabase("/home/paras/grpc-project/grpc-project/grpc-server/internal/utils/simulated_entry.json")

	fmt.Println(db)
	grpcServer := grpc.NewServer()
	// Register service implementation with the server
	pb.RegisterUserServiceServer(grpcServer, &service.UserService{
		Database: *db,
	})
	port := ":50051"
	// Start gRPC server on port :50051
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Printf("Grpc server is listen in port %v", port)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
}
