package main

import (
	"context"
	"log"

	pb "github.com/ParasJain0307/grpc-project/grpc-client/api" // Update with your actual package path
	"github.com/ParasJain0307/grpc-project/grpc-client/internal/utils/logger"

	"google.golang.org/grpc"
)

// ServiceClient struct wraps the gRPC client and provides methods to interact with the server.
type ServiceClient struct {
	client pb.UserServiceClient
}

// NewServiceClient creates a new instance of ServiceClient.
func NewServiceClient(conn *grpc.ClientConn) *ServiceClient {
	return &ServiceClient{
		client: pb.NewUserServiceClient(conn),
	}
}

// GetUserByID fetches a user by their ID from the server.
func (sc *ServiceClient) GetUserByID(userID int32) (*pb.User, error) {
	resp, err := sc.client.GetUserByID(context.Background(), &pb.GetUserByIDRequest{UserId: userID})
	if err != nil {
		logger.Errorf("Error fetching user by ID: %v", err)
		return nil, err
	}
	return resp, nil
}

// GetUsersByID fetches users by their IDs from the server.
func (sc *ServiceClient) GetUsersByID(userIDs []int32) (*pb.UsersList, error) {
	resp, err := sc.client.GetUsersByID(context.Background(), &pb.GetUsersByIDRequest{UserIds: userIDs})
	if err != nil {
		logger.Errorf("Error fetching users by IDs: %v", err)
		return nil, err
	}
	return resp, nil
}

// SearchUsers searches users based on criteria from the server.
func (sc *ServiceClient) SearchUsers(criteria []*pb.SearchCriteria) (*pb.UsersList, error) {
	req := &pb.SearchUsersRequest{
		Criterias: criteria,
	}
	resp, err := sc.client.SearchUsers(context.Background(), req)
	if err != nil {
		logger.Errorf("Error searching users: %v", err)
		return nil, err
	}
	return resp, nil
}

func main() {
	// Initialize logger
	logger, err := logger.InitLogger()
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}

	// Set up a connection to the gRPC server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		logger.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	// Create a new instance of ServiceClient
	client := NewServiceClient(conn)

	// Example calls to methods using the client
	userID := int32(1)
	user, err := client.GetUserByID(userID)
	if err != nil {
		logger.Errorf("Error getting user by ID: %v", err)
	} else {
		logger.Infof("User retrieved: %+v", user)
	}

	userIDs := []int32{1, 2, 3}
	users, err := client.GetUsersByID(userIDs)
	if err != nil {
		logger.Errorf("Error getting users by IDs: %v", err)
	} else {
		logger.Infof("Users retrieved: %+v", users)
	}

	searchCriteria := []*pb.SearchCriteria{
		{FieldName: "fname", FieldValue: "John"},
		{FieldName: "city", FieldValue: "New York"},
	}
	searchedUsers, err := client.SearchUsers(searchCriteria)
	if err != nil {
		logger.Errorf("Error searching users: %v", err)
	} else {
		logger.Infof("Users found: %+v", searchedUsers)
	}
}
