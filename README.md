# k8-go-backend

A simple Go backend API containerized with a Docker and deployed using Kubernetes (Minikube).

## Project Overview

This project demonstrates a basic Go HTTP server with JSON responses, the application provides two endpoints:
- `/` - Returns a JSON response with a message and version
- `/health` - Returns a simple "OK" text response for health checks

## Project Structure

```
k8-go-backend/
├── main.go            # Go application source code
├── Dockerfile         # Container build instructions
├── .gitignore         # Git ignore file
└── k8s/               # Kubernetes configuration
    ├── deployment.yaml  # Deployment specification
    └── service.yaml     # Service specification
```

## Application Details

The Go application uses the standard library to create a simple HTTP server with the following components:

- **Response Structure**: Defines the JSON response format
- **Responder Interface**: Demonstrates Go interfaces for flexible response creation
- **HTTP Handlers**: Functions for the root and health endpoints
- **Environment Configuration**: Uses environment variables with sensible defaults

## Docker Setup

The Dockerfile uses a multi-stage build process to create a small, efficient container:

1. **Build Stage**:
   - Uses `golang:1.21` as the base image
   - Builds the application with specific architecture settings
   - Creates a statically-linked binary

2. **Production Stage**:
   - Uses `alpine:latest` as a minimal base image
   - Includes only the compiled binary
   - Sets up compatibility libraries
   - Exposes port 8080

## Kubernetes Configuration

The project includes Kubernetes manifests for deployment to a cluster:

- **Deployment**: Creates a pod running the container
- **Service**: Exposes the application on NodePort 30001

## Running Locally

To build and run the application with Docker:

```bash
# Build the Docker image 
docker build -t k8-go-backend:latest .

# Run the container in docker from the image, use port 8080
docker run -p 8080:8080 k8-go-backend:latest
```

Then access:
- http://localhost:8080 - Main endpoint
- http://localhost:8080/health - Health endpoint

## Deploying to Minikube

Follow these steps to deploy to Minikube:

1. **Start Minikube**:
   ```bash
   minikube start
   ```

2. **Build and load the image into Minikube**:
   ```bash
   # Build the image locally
   docker build -t k8-go-backend:latest .
   
   # Load it into Minikube
   minikube image load k8-go-backend:latest
   ```

3. **Apply Kubernetes configurations**:
   ```bash
   kubectl apply -f k8s/deployment.yaml
   kubectl apply -f k8s/service.yaml
   ```

4. **Check the deployment**:
   ```bash
   kubectl get pods
   kubectl get services
   ```

5. **Access the service**:
   ```bash
   minikube service go-backend-service
   ```
   
   Or get the URL:
   ```bash
   minikube service go-backend-service --url
   ```

6. **For local development, set up port forwarding**:
   ```bash
   kubectl port-forward service/go-backend-service 8080:80
   ```
   
   Then access http://localhost:8080 and http://localhost:8080/health

## Technical Implementation Details

### Go Application

The Go application demonstrates several important concepts:

- **Interface Implementation**: The `Responder` interface is implemented by `CustomResponder`
- **JSON Marshaling**: Response data is automatically converted to JSON
- **HTTP Handler Pattern**: Using Go's standard HTTP library
- **Environment Configuration**: Using `os.Getenv()` with defaults

### Docker Optimizations

The Dockerfile includes several best practices:

- **Multi-stage builds**: Keeps the final image small
- **Cross-compilation**: Setting `CGO_ENABLED=0` and architecture flags
- **Alpine base**: Minimizes image size
- **Compatibility layer**: Includes `libc6-compat` for better compatibility

### Kubernetes Configuration

The Kubernetes configuration follows these practices:

- **Deployment**: Manages pod lifecycle and scaling
- **Service**: Exposes the application with a stable endpoint using NodePort
- **Labels and Selectors**: Properly connects the deployment to the service

## Troubleshooting

If pods show `ErrImageNeverPull` status:
1. Ensure the image is loaded into Minikube with `minikube image load`
2. Verify the image name in deployment.yaml matches what you built
3. Consider changing `imagePullPolicy` to `IfNotPresent`

If the service isn't accessible:
1. Check pod status with `kubectl get pods`
2. Check logs with `kubectl logs <pod-name>`
3. Use `minikube service go-backend-service` to get the correct URL