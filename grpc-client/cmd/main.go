package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	pb "github.com/ParasJain0307/grpc-project/grpc-client/api" // Update with your actual package path
	"github.com/ParasJain0307/grpc-project/grpc-client/internal/utils/logger"
	"github.com/ParasJain0307/grpc-project/grpc-client/internal/validation"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

const (
	serverAddress = "localhost:50052" // Address of the gRPC server
)

var loggerv1 *zap.SugaredLogger

func main() {
	// Initialize the logger
	var err error
	loggerv1, err = logger.InitLogger()
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer loggerv1.Sync() // Ensure any buffered log entries are flushed before the program exits

	// Set up a connection to the server
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		loggerv1.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	// Create a client instance using the connection.
	client := pb.NewUserServiceClient(conn)

	// Create a reader to read user input.
	reader := bufio.NewReader(os.Stdin)

	// Menu loop
	for {
		printMenu()

		// Read user input
		option, err := reader.ReadString('\n')
		if err != nil {
			loggerv1.Errorf("Failed to read input: %v", err)
			continue
		}

		// Remove newline character from input
		option = strings.TrimSpace(option)

		// Process user choice
		switch option {
		case "1":
			// Example: Fetch user by ID
			loggerv1.Info("Enter user ID: ")
			userIDInput, err := reader.ReadString('\n')
			if err != nil {
				loggerv1.Errorf("Error reading user ID input: %v", err)
				break
			}
			if err := validation.ValidateUserID(userIDInput); err != nil {
				loggerv1.Errorf("Validation error: %v", err)
				break
			}
			userIDInput = strings.TrimSpace(userIDInput)
			loggerv1.Info("After trimm", userIDInput)
			userID, err := strconv.Atoi(userIDInput)
			if err != nil {
				loggerv1.Errorf("Invalid user ID: %v", err)
				break
			}
			getUserByID(client, int32(userID))

		case "2":
			// Example: Fetch users by IDs
			loggerv1.Info("Enter user IDs (comma-separated): ")
			userIDsInput, err := reader.ReadString('\n')
			if err != nil {
				loggerv1.Errorf("Error reading user IDs input: %v", err)
				break
			}
			userIDsInput = strings.TrimSpace(userIDsInput)
			userIDsStr := strings.Split(userIDsInput, ",")
			var userIDs []int32
			for _, idStr := range userIDsStr {
				id, err := strconv.Atoi(strings.TrimSpace(idStr))
				if err != nil {
					loggerv1.Errorf("Invalid user ID: %v", err)
					continue
				}
				if err := validation.ValidateUserID(idStr); err != nil {
					loggerv1.Errorf("Validation error: %v", err)
					break
				}
				userIDs = append(userIDs, int32(id))
			}
			getUsersByID(client, userIDs)

		case "3":
			// Example: Search users by criteria
			loggerv1.Info("Enter search criteria:")
			criterias := readSearchCriteria(reader)
			loggerv1.Infof("Search criteria: %+v", criterias)
			if err := validation.ValidateSearchCriteria(criterias); err != nil {
				loggerv1.Errorf("Validation error: %v", err)
				break
			}
			searchUsers(client, criterias)

		case "q":
			// Quit
			loggerv1.Info("Exiting...")
			return

		default:
			loggerv1.Warn("Invalid option selected.")
		}
	}
}

func readSearchCriteria(reader *bufio.Reader) []*pb.SearchCriteria {
	var criterias []*pb.SearchCriteria

	for {
		fmt.Println("Enter search criteria (leave empty to finish):")
		fmt.Print("Field Name (e.g., fname, city, phone, height, married): ")
		fieldNameInput, err := reader.ReadString('\n')
		if err != nil {
			loggerv1.Fatalf("Error reading field name input: %v", err)
		}
		fieldNameInput = strings.TrimSpace(fieldNameInput)
		if fieldNameInput == "" {
			break
		}

		// Prompt user for criteria value
		fmt.Print("Enter value for " + fieldNameInput + ": ")
		valueInput, err := reader.ReadString('\n')
		if err != nil {
			loggerv1.Fatalf("Error reading value input: %v", err)
		}
		valueInput = strings.TrimSpace(valueInput)

		// Create SearchCriteria object and add to slice
		criteria := &pb.SearchCriteria{
			FieldName:  fieldNameInput,
			FieldValue: valueInput,
		}
		criterias = append(criterias, criteria)
	}

	return criterias
}

func searchUsers(client pb.UserServiceClient, criterias []*pb.SearchCriteria) {
	// Call the SearchUsers RPC method with multiple criterias
	req := &pb.SearchUsersRequest{
		Criterias: criterias,
	}
	resp, err := client.SearchUsers(context.Background(), req)
	if err != nil {
		loggerv1.Errorf("Error searching users: %v", err)
		return
	}
	loggerv1.Info(formatUsersListResponse(resp))
}

func printMenu() {
	fmt.Println("===== Menu =====")
	fmt.Println("1. Fetch User by ID")
	fmt.Println("2. Fetch Users by IDs")
	fmt.Println("3. Search Users by Criteria")
	fmt.Println("q. Quit")
	fmt.Print("Enter your choice: ")
}

func getUserByID(client pb.UserServiceClient, userID int32) {
	// Call the GetUserByID RPC method
	resp, err := client.GetUserByID(context.Background(), &pb.GetUserByIDRequest{UserId: userID})
	if err != nil {
		loggerv1.Errorf("Error fetching user by ID: %v", err)
		return
	}
	loggerv1.Info(formatUserResponse(resp))
}

func getUsersByID(client pb.UserServiceClient, userIDs []int32) {
	// Call the GetUsersByID RPC method
	resp, err := client.GetUsersByID(context.Background(), &pb.GetUsersByIDRequest{UserIds: userIDs})
	if err != nil {
		loggerv1.Errorf("Error fetching users by IDs: %v", err)
		return
	}
	loggerv1.Info(formatUsersListResponse(resp))
}

func formatUserResponse(user *pb.User) string {
	// Marshal user struct into JSON
	userJSON, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		loggerv1.Errorf("Error marshaling user to JSON: %v", err)
		return ""
	}

	// Create a strings.Builder for formatting
	var builder strings.Builder
	builder.WriteString(string(userJSON))
	builder.WriteString("\n")
	builder.WriteString(strings.Repeat("-", 30)) // Adding partition line
	builder.WriteString("\n")

	return builder.String()
}

func formatUsersListResponse(users *pb.UsersList) string {
	var builder strings.Builder
	for _, user := range users.Users {
		builder.WriteString(formatUserResponse(user))
		builder.WriteString("\n")
	}
	return builder.String()
}
