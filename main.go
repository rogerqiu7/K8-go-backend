// Declare the main package - this is the entry point of a standalone Go application
package main

// Import necessary standard libraries
import (
	"encoding/json" // for encoding data into JSON format
	"log"           // for logging messages to the console
	"net/http"      // for building an HTTP server and handling HTTP requests
	"os"            // for accessing environment variables like PORT
)

// Define a struct to shape the JSON response
type Response struct {
	Message string `json:"message"` // Message field will appear as "message" in JSON
	Version string `json:"version"` // Version field will appear as "version" in JSON
}

// Responder interface defines a method that returns a Response
type Responder interface {
	GetResponse() Response
}

// CustomResponder is a type that implements the Responder interface
type CustomResponder struct {
	Name    string
	Version string
}

// GetResponse is a method on CustomResponder that returns a Response
func (c CustomResponder) GetResponse() Response {
	return Response{
		Message: "Hello from " + c.Name + "!", // Create dynamic message
		Version: c.Version,
	}
}

// main is the entry point for the application
func main() {
	// Get the port from the environment variable "PORT"
	port := os.Getenv("PORT")

	// If PORT is not set, default to "8080"
	if port == "" {
		port = "8080"
	}

	// Register HTTP handlers:
	// "/" will be handled by handleRoot function
	// "/health" will be handled by handleHealth function
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/health", handleHealth)

	// Log the port the server is starting on
	log.Printf("Server starting on port %s", port)

	// Start the HTTP server on the given port
	// If server fails to start, log the error and exit
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

// handleRoot handles HTTP requests to the "/" endpoint
// w is the ResponseWriter used to construct the HTTP response
// r is the incoming HTTP request
func handleRoot(w http.ResponseWriter, r *http.Request) {

	// Create a CustomResponder and use it via the Responder interface
	// We create a CustomResponder instance and assign it to a Responder interface variable.
	var responder Responder = CustomResponder{
		Name:    "Go backend",
		Version: "1.0.0",
	}

	// Get the response using the interface method
	// The GetResponse method of the CustomResponder struct is called to get the response data.
	response := responder.GetResponse()

	// Set the HTTP response header "Content-Type" to "application/json"
	// This tells the client that the response body will be in JSON format.
	w.Header().Set("Content-Type", "application/json")

	// Encode the response struct to JSON and write it to the response
	json.NewEncoder(w).Encode(response)
}

// handleHealth handles HTTP requests to the "/health" endpoint
// w is the ResponseWriter used to construct the HTTP response
// r is the incoming HTTP request
func handleHealth(w http.ResponseWriter, r *http.Request) {
	// Set HTTP status code to 200 OK
	w.WriteHeader(http.StatusOK)

	// Write a plain-text "OK" response body
	w.Write([]byte("OK"))
}
