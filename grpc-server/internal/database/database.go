package database

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/ParasJain0307/grpc-project/grpc-server/internal/logger"
	"github.com/ParasJain0307/grpc-project/grpc-server/internal/utils"
	"go.uber.org/zap"

	pb "github.com/ParasJain0307/grpc-project/grpc-server/api"
)

type Database struct {
	users map[int32]*pb.User
}

// NewDatabase initializes a new database instance
func NewDatabase(jsonPath string) (*Database, error) {
	logger.Info("Initializing new database from JSON file")

	// Read JSON file
	jsonData, err := os.ReadFile(jsonPath)
	if err != nil {
		logger.Error("Failed to read JSON file", zap.Error(err))
		return nil, fmt.Errorf("error reading JSON file: %v", err)
	}

	// Unmarshal JSON directly into a slice of map[string]interface{}
	var users []map[string]interface{}
	err = json.Unmarshal(jsonData, &users)
	if err != nil {
		logger.Error("Failed to unmarshal JSON data", zap.Error(err))
		return nil, fmt.Errorf("error unmarshalling JSON data: %v", err)
	}

	// Initialize map with user IDs as keys
	userMap := make(map[int32]*pb.User)
	for _, u := range users {
		id, ok := u["id"].(float64) // assuming "id" field is float64 in JSON
		if !ok {
			logger.Warn("Failed to parse user ID")
			continue
		}
		user := &pb.User{
			Id:      int32(id),
			Fname:   u["fname"].(string),
			City:    u["city"].(string),
			Phone:   int64(u["phone"].(float64)),
			Height:  float32(u["height"].(float64)),
			Married: u["married"].(bool),
		}
		userMap[user.Id] = user
	}
	logger.Info("Database initialization complete")
	return &Database{
		users: userMap,
	}, nil
}

// GetUserByID retrieves a user by ID from the datastore
func (d *Database) GetUserByID(id int32) (*pb.User, error) {
	logger.Debugf("Fetching user with ID %v", id)
	user, ok := d.users[id]
	if !ok {
		logger.Warnf("User with ID %v not found", id)
		return nil, errors.New("user not found")
	}
	logger.Infof("User with ID %v found", id)
	return user, nil
}

// GetUsersByID retrieves a list of users by IDs from the datastore
func (d *Database) GetUsersByID(ids []int32) ([]*pb.User, error) {
	var users []*pb.User
	for _, id := range ids {
		user, ok := d.users[id]
		if ok {
			users = append(users, user)
			logger.Debugf("User with ID %v added to result", id)
		} else {
			logger.Warnf("User with ID %v not found", id)
		}
	}
	logger.Infof("Retrieved %v users by IDs", len(users))
	return users, nil
}

// SearchUsers searches users based on criteria in the datastore
func (d *Database) SearchUsers(criteria []*pb.SearchCriteria) ([]*pb.User, error) {
	var users []*pb.User
	for _, user := range d.users {
		// Example search criteria (can be customized)
		if matchCriteria(user, criteria) {
			users = append(users, user)
		}

	}
	if len(users) == 0 {
		logger.Warn("No users found matching the criteria")
		return nil, errors.New("no users found")
	}
	logger.Infof("Found %v users matching the criteria", len(users))
	return users, nil
}

// Function to check if a user matches the search criteria
func matchCriteria(user *pb.User, criterias []*pb.SearchCriteria) bool {
	// Iterate over criterias and check each one against user's attributes
	for _, criteria := range criterias {
		if !checkSingleCriteria(user, criteria) {
			return false
		}
	}
	return true
}

// checkSingleCriteria checks if a user matches a single criteria
func checkSingleCriteria(user *pb.User, criteria *pb.SearchCriteria) bool {
	switch criteria.FieldName {
	case utils.FIRSTNAME:
		if user.Fname != criteria.FieldValue {
			return false
		}
	case utils.CITY:
		if user.City != criteria.FieldValue {
			return false
		}
	case utils.PHONE:
		phone, err := strconv.ParseInt(criteria.FieldValue, 10, 64)
		if err != nil || user.Phone != phone {
			return false
		}
	case utils.HEIGHT:
		height, err := strconv.ParseFloat(criteria.FieldValue, 32)
		if err != nil || float32(height) != user.Height {
			return false
		}
	case utils.MARRIED:
		married, err := strconv.ParseBool(criteria.FieldValue)
		if err != nil || user.Married != married {
			return false
		}
	default:
		return false
	}

	return true
}
