package main

import (
	"context"
	"log"

	pb "github.com/ParasJain0307/grpc-project/grpc-client/api" // Update with your actual package path

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
		log.Printf("Error fetching user by ID: %v", err)
		return nil, err
	}
	return resp, nil
}

// GetUsersByID fetches users by their IDs from the server.
func (sc *ServiceClient) GetUsersByID(userIDs []int32) (*pb.UsersList, error) {
	resp, err := sc.client.GetUsersByID(context.Background(), &pb.GetUsersByIDRequest{UserIds: userIDs})
	if err != nil {
		log.Printf("Error fetching users by IDs: %v", err)
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
		log.Printf("Error searching users: %v", err)
		return nil, err
	}
	return resp, nil
}
