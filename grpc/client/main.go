package main

import (
	"context"
	"fmt"

	pb "wireless_lab_1/grpc/services" // Import the generated package

	"google.golang.org/grpc"
)

func main() {
	// Dial the master server at localhost:8080
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		fmt.Println("did not connect:", err)
		return
	}
	defer conn.Close()

	// Create a client instance
	c := pb.NewServicesClient(conn)

	// Call the RPC method
	resp, err := c.ClientToMasterUpload(context.Background(), &pb.ClientToMasterUploadRequest{})
	if err != nil {
		fmt.Println("Error calling Capitalize:", err)
		return
	}
	fmt.Println("Received Upload Response:", resp)
	
}
