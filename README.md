# gRPC Service for User Management and Search

This project implements a gRPC service in Go for managing user details and includes a search capability.


Project Structure
project-root/
│
├── api/
│ ├── user.proto # gRPC protocol buffer definition
│ └── user.pb.go # Generated Go code from the proto file
│
├── cmd/
│ └── service/
│ └── main.go # Main application entry point
│
├── internal/
│ ├── service/
│ │ ├── service.go # Implementation of gRPC service methods
│ │ └── user.go # User management logic
│ │
│ ├── datastore/
│ │ └── datastore.go # Simulated database logic
│ │
│ ├── search/
│ │ └── search.go # Implementation of search functionality
│ │
│ ├── validation/
│ │ └── validation.go # Input validation logic
│ │
│ └── utils/
│ └── utils.go # Utility functions
│
├── pkg/
│ └── grpc/
│ └── middleware.go # Middleware for gRPC server
│
├── tests/
│ ├── unit/ # Unit tests
│ └── integration/ # Integration tests
│
├── configs/
│ └── config.yaml # Configuration file
│
├── Dockerfile # Dockerfile for containerization
├── README.md # Project documentation (this file)
└── go.mod # Go module file


Requirements
 Go version 1.16 or higher.
 Protocol Buffers Compiler (protoc) and Go plugin (protoc-gen-go)


Installation

1. Clone the repository:
    git clone <repository_url>
    cd grpc-project

2. Install dependencies:
    go mod tidy


Building and Running the Application
   To build and run the gRPC server:
    go run cmd/service/main.go

The server will start listening on port 50051.


gRPC Client Usage

Client Setup
Make sure to update the gRPC client (grpc-client/main.go) with the correct server address (serverAddress variable).

const (
    serverAddress = "localhost:50051"  // Update with your server address
)


Running the Client
    go run grpc-client/main.go

Available Options
    Fetch User by ID: Fetches user details by ID.
    Fetch Users by IDs: Fetches details for multiple users by their IDs.
    Search Users by Criteria: Searches for users based on specific criteria (e.g., city, phone number).

API Documentation
gRPC Endpoints
  GetUserByID
    Fetches user details by ID.
  
   GetUsersByID
    Fetches details for multiple users by their IDs.

   SearchUsers
    Searches for users based on specified criteria.

Configuration
    Configuration details can be found in configs/config.yaml.


Testing
    Running Tests
        Run unit tests:
            go test ./tests/unit/...

Run integration tests:
    go test ./tests/integration/...


Docker Support
    To containerize the application using Docker:

        docker build -t grpc-project .
        docker run -p 50051:50051 grpc-project

    Replace grpc-project with your preferred image name.


Additional Notes
    Ensure all dependencies are installed and configured correctly before running the application.
    Handle errors gracefully and provide meaningful responses.
    Maintain code quality by adhering to Go best practices and conventions.




