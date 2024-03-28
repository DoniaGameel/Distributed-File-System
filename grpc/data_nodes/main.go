package main

import (
    "context"
    "log"
	"fmt"
    "google.golang.org/grpc"
    pb "wireless_lab_1/grpc/services" // Import the generated code
)

func main() {
    // Set up a connection to the gRPC server.
    conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
    if err != nil {
		fmt.Println("did not connect:", err)
		return
	}
    defer conn.Close()
    // Create a gRPC client
    client := pb.NewServicesClient(conn)

    // Call the Ping function
    response, err := client.Ping(context.Background(), &pb.HeartbeatRequest{DataNodeId: "1"})
    if err != nil {
        log.Fatalf("Error calling Ping: %v", err)
    }

    // Process the response
    if response.IsAlive {
        log.Println("Server is alive")
    } else {
        log.Println("Server is not alive")
    }
}
