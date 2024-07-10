package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	pb "github.com/ParasJain0307/grpc-project/grpc-server/api"
	httpServer "github.com/ParasJain0307/grpc-project/grpc-server/httpserver"
	"github.com/ParasJain0307/grpc-project/grpc-server/internal/database"
	"github.com/ParasJain0307/grpc-project/grpc-server/internal/logger"
	"github.com/ParasJain0307/grpc-project/grpc-server/internal/service"
	"github.com/ParasJain0307/grpc-project/grpc-server/internal/utils"
	"google.golang.org/grpc"
)

func main() {
	// Initialize the  custom logger
	loggerv1, err := logger.InitLogger()
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}

	// Retrieve JSON file path for simulated entries from environment variable
	jsonFilePath := os.Getenv("JSON_FILE_PATH")
	if jsonFilePath == "" {
		jsonFilePath = "internal/utils/simulated_entry.json"
	}

	// Initialize the simulated database
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

	// Start listening for incoming connections on port :50051
	listener, err := net.Listen("tcp", utils.GRPCSERVERPORT)
	if err != nil {
		loggerv1.Errorf("Failed to listen: %v", err)
		log.Fatalf("Failed to listen: %v", err)
	}

	// Log the gRPC server start
	loggerv1.Infof("gRPC server is listening on port %s", utils.GRPCSERVERPORT)

	// Handle OS signals for graceful shutdown of gRPC server
	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
		s := <-sig
		loggerv1.Infof("Received signal %v. Gracefully shutting down gRPC server...", s)
		grpcServer.GracefulStop()
	}()

	// Start the gRPC server in a separate goroutine
	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			loggerv1.Errorf("Failed to serve gRPC server: %v", err)
			log.Fatalf("Failed to serve gRPC server: %v", err)
		}
	}()
	// Calling HttpServer for exposing endpoint to the server asynchronise
	go httpServer.HttpServer()

	waitForSignal()
}

// waitForSignal blocks until SIGINT or SIGTERM signal is received
func waitForSignal() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
}
