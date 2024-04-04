package main

import (
	"context"
	"fmt"
	"net"
	"time"

	pb "wireless_lab_1/grpc/services" // Import the generated package

	"google.golang.org/grpc"
)
type dataNodeServer struct {
	pb.UnimplementedServicesServer
}

func (s *dataNodeServer) ClientToDataKeeperUpload(_ context.Context, req *pb.ClientToDataKeeperUploadRequest) (*pb.ClientToDataKeeperUploadResponse, error) {
	fileName := req.GetFileName()
	fmt.Println("Received file:", fileName)
	masterConn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
    if err != nil {
        fmt.Println("Failed to connect to master:", err)
        // Handle connection error (e.g., retry or log)
        return nil, err
    }
    defer masterConn.Close() // Ensure master connection is closed

    masterClient := pb.NewServicesClient(masterConn)

    // Prepare notification message
    notification := &pb.DataNodeNotificationRequest{
        FileName: fileName,
		NodeId: nodeId,
    }

    // Send notification to the master
    _, err = masterClient.DataNodeNotifyMaster(context.Background(), notification)
    if err != nil {
        fmt.Println("Error sending notification to master:", err)
        // Handle notification error (e.g., retry or log)
    } else {
        fmt.Println("Successfully notified master about received file")
    }

    return &pb.ClientToDataKeeperUploadResponse{}, nil
}

var nodeId string

func main() {
	// establish the node as a client
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		fmt.Println("did not connect:", err)
		return
	}
	defer conn.Close()
	c := pb.NewServicesClient(conn)

	// establish the node as a server
	go func() {
        lis, err := net.Listen("tcp", ":8081")
        if err != nil {
            fmt.Println("failed to listen:", err)
            return
        }
        s := grpc.NewServer()
        pb.RegisterServicesServer(s, &dataNodeServer{})
        fmt.Println("Server started. Listening on port 8081...")
        if err := s.Serve(lis); err != nil {
            fmt.Println("failed to serve:", err)
        }
    }()
	// Read input from user
	fmt.Print("Enter Node ID: ")
	fmt.Scanln(&nodeId)

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
		err = stream.Send(&pb.HeartbeatRequest{NodeId: nodeId})
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