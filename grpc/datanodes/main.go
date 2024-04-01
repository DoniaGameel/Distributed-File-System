package main

import (
	"context"
	"fmt"

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

	// Call the RPC HeartBeat method
	stream, err := c.TrackHeartbeat(context.Background())
	if err != nil {
		fmt.Println("Error calling TrackHeartbeat:", err)
		return
	}

	// Send the request
	err = stream.Send(&pb.HeartbeatRequest{NodeId: "1000"})
	if err != nil {
		fmt.Println("Error sending HeartbeatRequest:", err)
		return
	}

	// Receive and process responses from the stream
	for {
		resp, err := stream.Recv()
		if err != nil {
			fmt.Println("Error receiving HeartbeatResponse:", err)
			return
		}
		fmt.Println("Received HeartbeatResponse:", resp)
	}
}
