package main

import (
	"context"
	"fmt"
	"time"

	pb "wireless_lab_1/grpc/services" // Import the generated package

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		fmt.Println("did not connect:", err)
		return
	}
	defer conn.Close()
	c := pb.NewServicesClient(conn)

	// Start a ticker that triggers sending heartbeat every 1 second
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		// Call the RPC HeartBeat method
		stream, err := c.TrackHeartbeat(context.Background())
		if err != nil {
			fmt.Println("Error calling TrackHeartbeat:", err)
			continue // Retry sending heartbeat on the next tick
		}

		// Send the request
		err = stream.Send(&pb.HeartbeatRequest{NodeId: "1000"})
		if err != nil {
			fmt.Println("Error sending HeartbeatRequest:", err)
			continue // Retry sending heartbeat on the next tick
		}

		// Receive and process responses from the stream
		resp, err := stream.Recv()
		if err != nil {
			fmt.Println("Error receiving HeartbeatResponse:", err)
			continue // Retry sending heartbeat on the next tick
		}
		fmt.Println("Received HeartbeatResponse:", resp)
	}
}