# Start with a base image that includes Go
FROM golang:1.22.3 AS builder

# Set the current working directory inside the container
WORKDIR /app

# Copy the Go module files for dependency resolution
COPY go.mod go.sum ./

# Create necessary directories
RUN mkdir -p internal/utils

# Copy simulated_entry.json from local filesystem to internal/utils directory in the Docker image
COPY internal/utils/simulated_entry.json internal/utils/

# Download Go module dependencies
RUN go mod download

# Copy the entire source code into the container
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/service/main.go

# Start a new stage for a minimal deployment image
FROM alpine:latest

# Set environment variables if necessary
ENV PORT=50051

# Create necessary directories
RUN mkdir -p /app/internal/utils

# Copy the built executable from the builder stage
COPY --from=builder /app/server /app/server

# Copy the simulated_entry.json from builder stage
COPY --from=builder /app/internal/utils/simulated_entry.json /app/internal/utils/simulated_entry.json

# Expose the port on which the service will run
EXPOSE $PORT

# Command to run the executable
CMD ["/app/server"]
