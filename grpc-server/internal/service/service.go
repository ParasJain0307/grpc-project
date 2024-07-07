package service

import (
	"context"
	"log"

	pb "github.com/ParasJain0307/grpc-project/grpc-server/api"
	"github.com/ParasJain0307/grpc-project/grpc-server/internal/database"
)

// UserService implements the UserServiceServer interface
type UserService struct {
	pb.UnimplementedUserServiceServer
	Database database.Database // Example simulated datastore
}

// mustEmbedUnimplementedUserServiceServer implements __.UserServiceServer.
// func (s *UserService) mustEmbedUnimplementedUserServiceServer() {
// 	panic("unimplemented")
// }

// GetUserByID implements the GetUserByID method from the protobuf definition
func (s *UserService) GetUserByID(ctx context.Context, req *pb.GetUserByIDRequest) (*pb.User, error) {
	// Example implementation fetching user from datastore
	log.Println("Request Received", req)
	user, err := s.Database.GetUserByID(req.UserId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// GetUsersByID implements the GetUsersByID method from the protobuf definition
func (s *UserService) GetUsersByID(ctx context.Context, req *pb.GetUsersByIDRequest) (*pb.UsersList, error) {
	// Example implementation fetching list of users from datastore
	log.Println("Request Received", req)
	users, err := s.Database.GetUsersByID(req.UserIds)
	if err != nil {
		return nil, err
	}
	return &pb.UsersList{Users: users}, nil
}

// SearchUsers implements the SearchUsers method from the protobuf definition
func (s *UserService) SearchUsers(ctx context.Context, req *pb.SearchUsersRequest) (*pb.UsersList, error) {
	// Example implementation searching users based on criteria
	log.Println("Request Received", req)
	criteria := req.GetCriterias()
	users, err := s.Database.SearchUsers(criteria)
	if err != nil {
		return nil, err
	}
	return &pb.UsersList{Users: users}, nil
}
