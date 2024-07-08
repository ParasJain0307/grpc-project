package service

import (
	"context"

	pb "github.com/ParasJain0307/grpc-project/grpc-server/api"
	"github.com/ParasJain0307/grpc-project/grpc-server/internal/database"
	"github.com/ParasJain0307/grpc-project/grpc-server/internal/logger"
)

// UserService implements the UserServiceServer interface
type UserService struct {
	pb.UnimplementedUserServiceServer
	Database database.Database // Example simulated datastore
}

// NewService creates a new UserService instance
func NewService() *UserService {
	return &UserService{}
}

// GetUserByID implements the GetUserByID method from the protobuf definition
func (s *UserService) GetUserByID(ctx context.Context, req *pb.GetUserByIDRequest) (*pb.User, error) {
	logger.Info("GetUserByID called", "user_id", req.UserId)
	user, err := s.Database.GetUserByID(req.UserId)
	if err != nil {
		logger.Error("Failed to get user by ID", "user_id", req.UserId, "error", err)
		return nil, err
	}
	logger.Info("User retrieved by ID", "user_id", req.UserId)
	return user, nil
}

// GetUsersByID implements the GetUsersByID method from the protobuf definition
func (s *UserService) GetUsersByID(ctx context.Context, req *pb.GetUsersByIDRequest) (*pb.UsersList, error) {
	logger.Info("GetUsersByID called", "user_ids", req.UserIds)
	users, err := s.Database.GetUsersByID(req.UserIds)
	if err != nil {
		logger.Error("Failed to get users by IDs", "user_ids", req.UserIds, "error", err)
		return nil, err
	}
	logger.Info("Users retrieved by IDs", "num_users", len(users))
	return &pb.UsersList{Users: users}, nil
}

// SearchUsers implements the SearchUsers method from the protobuf definition
func (s *UserService) SearchUsers(ctx context.Context, req *pb.SearchUsersRequest) (*pb.UsersList, error) {
	logger.Info("SearchUsers called", "request", req)
	criteria := req.GetCriterias()
	users, err := s.Database.SearchUsers(criteria)
	if err != nil {
		logger.Error("Failed to search users", "request", req, "error", err)
		return nil, err
	}
	logger.Info("Users found matching criteria", "num_users", len(users))
	return &pb.UsersList{Users: users}, nil
}
