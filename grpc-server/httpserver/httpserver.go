package httpserver

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	pb "github.com/ParasJain0307/grpc-project/grpc-server/api"
	"github.com/ParasJain0307/grpc-project/grpc-server/internal/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

const (
	grpcServerAddress = utils.GRPCSERVERADDR + utils.GRPCSERVERPORT // Address of your gRPC server
	httpServerPort    = utils.HTTPSERVERPORT                        // Port on which the HTTP server will listen
)

func HttpServer() {
	mux := http.NewServeMux()

	// Create a gRPC client connection
	grpcConn, err := grpc.Dial(grpcServerAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial gRPC server: %v", err)
	}
	defer grpcConn.Close()

	client := pb.NewUserServiceClient(grpcConn)

	// Handler for /user/:id endpoint
	mux.HandleFunc("/user/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		userIDStr := r.URL.Path[len("/user/"):]
		userID, err := strconv.Atoi(userIDStr)
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}

		req := &pb.GetUserByIDRequest{UserId: int32(userID)}
		user, err := client.GetUserByID(context.Background(), req)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to get user: %v", err), http.StatusInternalServerError)
			return
		}

		writeJSONResponse(w, user)
	})

	// Handler for /users/:ids endpoint
	mux.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		userIDsStr := r.URL.Path[len("/users/"):]
		userIDs := strings.Split(userIDsStr, ",")

		var userIDsInt32 []int32
		for _, idStr := range userIDs {
			id, err := strconv.Atoi(idStr)
			if err != nil {
				http.Error(w, fmt.Sprintf("Invalid user ID: %s", idStr), http.StatusBadRequest)
				return
			}
			userIDsInt32 = append(userIDsInt32, int32(id))
		}

		req := &pb.GetUsersByIDRequest{UserIds: userIDsInt32}
		usersList, err := client.GetUsersByID(context.Background(), req)
		if err != nil {
			handleGRPCError(w, err)
			return
		}

		writeJSONResponse(w, usersList)
	})

	// Handler for /users/search endpoint
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

		req := &pb.SearchUsersRequest{Criterias: criterias}
		usersList, err := client.SearchUsers(context.Background(), req)
		if err != nil {
			handleGRPCError(w, err)
			return
		}

		writeJSONResponse(w, usersList)
	})

	// Start HTTP server
	server := &http.Server{
		Addr:    httpServerPort,
		Handler: mux,
	}

	log.Printf("Starting HTTP server on port %s...\n", httpServerPort)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}
}

func writeJSONResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to marshal JSON: %v", err), http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}

func handleGRPCError(w http.ResponseWriter, err error) {
	st, _ := status.FromError(err)
	http.Error(w, fmt.Sprintf("Failed to execute gRPC request: %v", st.Message()), http.StatusInternalServerError)
}
