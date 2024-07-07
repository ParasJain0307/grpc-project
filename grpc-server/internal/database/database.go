package database

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	pb "github.com/ParasJain0307/grpc-project/grpc-server/api" // Replace with your actual package path
)

type Database struct {
	users map[int32]*pb.User
}

// NewDatabase initializes a new database instance
func NewDatabase(jsonPath string) (*Database, error) {
	// Read JSON file
	jsonData, err := os.ReadFile(jsonPath)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("error reading JSON file: %v", err)
	}
	fmt.Println(jsonData)

	// Unmarshal JSON directly into a slice of map[string]interface{}
	var users []map[string]interface{}
	err = json.Unmarshal(jsonData, &users)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON data: %v", err)
	}

	// Initialize map with user IDs as keys
	userMap := make(map[int32]*pb.User)
	for _, u := range users {
		id, ok := u["id"].(float64) // assuming "id" field is float64 in JSON
		if !ok {
			return nil, fmt.Errorf("error parsing user ID")
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
	return &Database{
		users: userMap,
	}, nil
}

// GetUserByID retrieves a user by ID from the datastore
func (d *Database) GetUserByID(id int32) (*pb.User, error) {
	user, ok := d.users[id]
	if !ok {
		return nil, errors.New("user not found")
	}
	return user, nil
}

// GetUsersByID retrieves a list of users by IDs from the datastore
func (d *Database) GetUsersByID(ids []int32) ([]*pb.User, error) {
	var users []*pb.User
	for _, id := range ids {
		user, ok := d.users[id]
		if ok {
			users = append(users, user)
		}
	}
	return users, nil
}

// SearchUsers searches users based on criteria in the datastore
func (d *Database) SearchUsers(criteria string) ([]*pb.User, error) {
	var users []*pb.User
	for _, user := range d.users {
		// Example search criteria (can be customized)
		if user.City == criteria {
			users = append(users, user)
		}
	}
	if len(users) == 0 {
		return nil, errors.New("Request User Info does not Exist")
	}
	return users, nil
}
