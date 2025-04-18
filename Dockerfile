# Use Go official image as build environment
FROM golang:1.21 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy all files from the current directory to the working directory
COPY . .

# Initialize a Go module and build the application
# - CGO_ENABLED=0: Disable C Go to create a statically linked binary
# - GOOS=linux: Target Linux operating system
# - GOARCH=amd64: Target 64-bit x86 architecture
RUN go mod init myapp && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server main.go

# Start a new build stage with a minimal Alpine Linux image
FROM alpine:latest

# Set the working directory in the new image
WORKDIR /root/

# Copy only the compiled binary from the builder stage
# This reduces the final image size by excluding unnecessary files
# and dependencies that were only needed for building the application
COPY --from=builder /app/server .

# Install compatibility libraries needed for Go binaries on Alpine
RUN apk add --no-cache libc6-compat

# Expose port 8080 for the application
# This allows the application to accept incoming connections on this port
EXPOSE 8080

# Define the command to run when the container starts
CMD ["./server"]