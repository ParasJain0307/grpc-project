# gRPC Service for User Management and Search

This project implements a gRPC service in Go for managing user details and includes a search capability.

Requirements
- Go version 1.16 or higher.
- Protocol Buffers Compiler (protoc) and Go plugin (protoc-gen-go).


Installation

1. Clone the repository:
    git clone <repository_url>
    cd grpc-project

2. Install dependencies for both grpc-server and grpc-client:
    go mod tidy


Building and Running the Application
    
    To build and run the gRPC server:
        make clean
        make run

The gRPC server will start listening on port 50051.


gRPC Client Usage

Client Setup

Ensure to update the gRPC client (grpc-client/main.go) with the correct server address (serverAddress variable).

const (
    serverAddress = "localhost:50051"  // Update with your server address
)

Running the Client
    make run
    Menu-Driven will be open like 
        1. Fetch User by ID
        2. Fetch Users by IDs <this should be comma separated>
        3. Search Users by Criteria <this should key value form refer Ex below>
        q. Quit
    User can choose the option provide the info and get the desired data.
        For Search Users Input like 
            Enter search criteria (leave empty to finish):
                Field Name (e.g., fname, city, phone, height, married): married

Available Options
- Fetch User by ID: Fetches user details by ID.
- Fetch Users by IDs: Fetches details for multiple users by their IDs.
- Search Users by Criteria: Searches for users based on specific criteria (e.g., city, phone number).

There is another way to access user info without running the server. 
    Http server is running asynchronous while grpc-server is up and it will expose the Api endpoint through which user can get the data
    - Fetch User by ID: Fetches user details by ID.
      GET <localhost:port>/user/{userid}
    - Fetch Users by IDs: Fetches details for multiple users by their IDs.
      GET  <localhost:port>/users/{userid1,userid2 ...}
    - Search Users by Criteria: Searches for users based on specific criteria (e.g., city, phone number).
      POST <localhost:port>/users/search and in body provide like [{"field_name":"married", "field_value":"true"} ...]


API Documentation

gRPC Endpoints
- GetUserByID: Fetches user details by ID.
- GetUsersByID: Fetches details for multiple users by their IDs.
- SearchUsers: Searches for users based on specified criteria.


Docker Support

To containerize the application using Docker:
    cd grpc-project/grpc-server
    docker build -t grpc-project .
    docker run -p 50051:50051 -p 8082:8082 grpc-project

Replace grpc-project with your preferred image name.


Kubernetes Engine Support
    
To deploy your application over a Kubernetes cluster:

1. Ensure you have a Kubernetes environment like EKS, Minikube, etc.
2. Use deployment.yaml for deploying the application. A template is available at config/deployments/deployment.yaml. Modify the image   repository details as needed.
3. Create a Service for communication through the host. A template is available at config/deployments/service.yaml.
4. Create Secrets for storing Docker Hub credentials. A template is available at config/deployments/secrets.yaml.


Additional Notes
- Ensure all dependencies are correctly installed and configured before running the application.
- Handle errors gracefully and provide meaningful responses.
- Maintain code quality by adhering to Go best practices and conventions.




