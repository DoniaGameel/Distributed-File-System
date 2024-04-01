package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"strings"

	pb "wireless_lab_1/grpc/services" // Import the generated package

	"google.golang.org/grpc"
)

type masterServer struct {
	pb.UnimplementedServicesServer
}

func (s *masterServer) Hello(ctx context.Context, req *pb.TextRequest) (*pb.TextResponse, error) {
	text := req.GetText()
	capitalizedText := strings.ToUpper(text)
	return &pb.TextResponse{CapitalizedText: capitalizedText}, nil
}

/*func (s *masterServer) TrackHeartbeat(ctx context.Context, req *pb.HeartbeatRequest) (*pb.HeartbeatResponse, error) {
	text := req.GetNodeId()
	fmt.Println("Node ID:", text)
	return &pb.HeartbeatResponse{}, nil
}*/

func (s *masterServer) TrackHeartbeat(stream pb.Services_TrackHeartbeatServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		text := req.GetNodeId()
		fmt.Println("Node ID:", text)

		// Send a HeartbeatResponse back to the Data Keeper node
        if err := stream.Send(&pb.HeartbeatResponse{}); err != nil {
            return err
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
	pb.RegisterServicesServer(s, &masterServer{})
	fmt.Println("Server started. Listening on port 8080...")
	if err := s.Serve(lis); err != nil {
		fmt.Println("failed to serve:", err)
	}
}
