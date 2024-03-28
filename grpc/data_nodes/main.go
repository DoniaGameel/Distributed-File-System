package main

// Import the generated package
import (
	pb "Lab1/services" // Import the generated package
	"context"
	"fmt"
	"net"
	"time"

	"google.golang.org/grpc"
)

type heartBeatServer struct {
	pb.UnimplementedHeartbeatServiceServer
}

func (s *heartBeatServer) HeartBeat(ctx context.Context, req *pb.HeartbeatRequest) (*pb.HeartbeatResponse, error) {
	return &pb.HeartbeatResponse{IsAlive: true}, nil
}
// Function to send "still alive" messages every 1 second
func sendStillAliveMessages(stream pb.HeartbeatService_PingServer, done <-chan struct{}) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			// Send "still alive" message
			err := stream.Send(&pb.HeartbeatResponse{IsAlive: true})
			if err != nil {
				fmt.Println("Error sending 'still alive' message:", err)
				return
			}
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("failed to listen:", err)
		return
	}
	s := grpc.NewServer()
	pb.RegisterHeartbeatServiceServer(s, &heartBeatServer{})
	fmt.Println("Server started. Listening on port 8080...")
	if err := s.Serve(lis); err != nil {
		fmt.Println("failed to serve:", err)
	}
}