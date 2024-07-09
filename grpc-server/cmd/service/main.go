package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	pb "github.com/ParasJain0307/grpc-project/grpc-server/api"
	"github.com/ParasJain0307/grpc-project/grpc-server/internal/database"
	"github.com/ParasJain0307/grpc-project/grpc-server/internal/logger"
	"github.com/ParasJain0307/grpc-project/grpc-server/internal/service"
	"google.golang.org/grpc"
)

func main() {
	// Initialize the logger
	loggerv1, err := logger.InitLogger()
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}

	// Initialize the simulated Database
	jsonFilePath := "/home/paras/grpc-project/backup/grpc-project/grpc-server/internal/utils/simulated_entry.json"

	// Read the contents of simulated.json
	//absJSONFilePath := filepath.Join("/app", jsonFilePath)
	//loggerv1.Infof("Using JSON file path: %s", absJSONFilePath)
	db, err := database.NewDatabase(jsonFilePath)
	if err != nil {
		loggerv1.Errorf("Error while initializing database: %v", err)
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Create a new gRPC server instance
	grpcServer := grpc.NewServer()

	// Register your service implementation with the gRPC server
	pb.RegisterUserServiceServer(grpcServer, &service.UserService{
		Database: *db,
	})

	// Define the port on which the server will listen
	port := ":50051"

	// Start listening for incoming connections on port :50051
	listener, err := net.Listen("tcp", port)
	if err != nil {
		loggerv1.Errorf("Failed to listen: %v", err)
		log.Fatalf("Failed to listen: %v", err)
	}

	// Log the server start
	loggerv1.Infof("gRPC server is listening on port %s", port)

	// Handle OS signals for graceful shutdown
	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
		s := <-sig
		loggerv1.Infof("Received signal %v. Gracefully shutting down...", s)
		grpcServer.GracefulStop()
	}()

	// Serve blocks until the server is stopped with `grpcServer.Stop()` or `grpcServer.GracefulStop()`
	if err := grpcServer.Serve(listener); err != nil {
		loggerv1.Errorf("Failed to serve gRPC server: %v", err)
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
