package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	pb "github.com/ParasJain0307/grpc-project/grpc-server/api"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

const (
	grpcServerAddress = "localhost:50051" // Address of your gRPC server
)

func main() {
	mux := http.NewServeMux()
	grpcConn, err := grpc.Dial(grpcServerAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial gRPC server: %v", err)
	}
	defer grpcConn.Close()

	client := pb.NewUserServiceClient(grpcConn)

	mux.HandleFunc("/user/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Extract user_id from URL path
		user_idStr := r.URL.Path[len("/user/"):]
		user_id, err := strconv.Atoi(user_idStr)
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}

		// Call gRPC method GetUserByID with int32 user_id
		req := &pb.GetUserByIDRequest{UserId: int32(user_id)}
		user, err := client.GetUserByID(context.Background(), req)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to get user: %v", err), http.StatusInternalServerError)
			return
		}

		// Return user as JSON response (pretty-printed)
		w.Header().Set("Content-Type", "application/json")
		jsonData, err := json.MarshalIndent(user, "", "  ")
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to marshal JSON: %v", err), http.StatusInternalServerError)
			return
		}
		w.Write(jsonData)
	})

	mux.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Extract user_ids from URL path
		userIDsStr := r.URL.Path[len("/users/"):]
		userIDs := strings.Split(userIDsStr, ",")

		// Convert userIDs to []int32
		var userIDsInt32 []int32
		for _, idStr := range userIDs {
			id, err := strconv.Atoi(idStr)
			if err != nil {
				http.Error(w, fmt.Sprintf("Invalid user ID: %s", idStr), http.StatusBadRequest)
				return
			}
			userIDsInt32 = append(userIDsInt32, int32(id))
		}

		// Call gRPC method GetUsersByID with []int32 user_ids
		req := &pb.GetUsersByIDRequest{UserIds: userIDsInt32}
		usersList, err := client.GetUsersByID(context.Background(), req)
		if err != nil {
			st, _ := status.FromError(err)
			http.Error(w, fmt.Sprintf("Failed to get users: %v", st.Message()), http.StatusInternalServerError)
			return
		}

		// Return users as JSON response (pretty-printed)
		w.Header().Set("Content-Type", "application/json")
		jsonData, err := json.MarshalIndent(usersList, "", "  ")
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to marshal JSON: %v", err), http.StatusInternalServerError)
			return
		}
		w.Write(jsonData)
	})

	mux.HandleFunc("/users/search", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var criterias []*pb.SearchCriteria
		if err := json.NewDecoder(r.Body).Decode(&criterias); err != nil {
			http.Error(w, fmt.Sprintf("Failed to parse request body: %v", err), http.StatusBadRequest)
			return
		}

		// Call gRPC method SearchUsers with criterias
		req := &pb.SearchUsersRequest{Criterias: criterias}
		usersList, err := client.SearchUsers(context.Background(), req)
		if err != nil {
			st, _ := status.FromError(err)
			http.Error(w, fmt.Sprintf("Failed to search users: %v", st.Message()), http.StatusInternalServerError)
			return
		}

		// Return users as JSON response (pretty-printed)
		w.Header().Set("Content-Type", "application/json")
		jsonData, err := json.MarshalIndent(usersList, "", "  ")
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to marshal JSON: %v", err), http.StatusInternalServerError)
			return
		}
		w.Write(jsonData)
	})

	// Start HTTP server
	log.Println("Starting HTTP server on port :8082...")
	if err := http.ListenAndServe(":8082", mux); err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}
}
