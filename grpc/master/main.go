package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"sync"
	"time"

	pb "wireless_lab_1/grpc/services" // Import the generated package

	"google.golang.org/grpc"
)

type nodeStatus struct {
	id       string
	alive    bool
	lastSeen time.Time
	filenames []string
	mutex     sync.Mutex
}

type masterServer struct {
	nodes map[string]*nodeStatus
	mutex sync.Mutex
	pb.UnimplementedServicesServer
}

// mustEmbedUnimplementedServicesServer implements services.ServicesServer.
func (s *masterServer) mustEmbedUnimplementedServicesServer() {
	panic("unimplemented")
}

func (s *masterServer) TrackHeartbeat(stream pb.Services_TrackHeartbeatServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		nodeId := req.GetNodeId()

		s.mutex.Lock()
		node, ok := s.nodes[nodeId]
		if !ok {
			node = &nodeStatus{
				id:       nodeId,
				alive:    true,
				lastSeen: time.Now(),
				filenames: []string{},
			}
			s.nodes[nodeId] = node
		}
		node.lastSeen = time.Now()
		node.alive = true
		s.mutex.Unlock()

		// Send a HeartbeatResponse back to the Data Keeper node
		if err := stream.Send(&pb.HeartbeatResponse{}); err != nil {
			return err
		}
	}
}

func (s *masterServer) monitorLiveness() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		s.mutex.Lock()
		for nodeId, node := range s.nodes {
			if time.Since(node.lastSeen) > 1*time.Second {
				node.alive = false
				fmt.Printf("Node with ID %s interrupted and marked down as non-responsive\n", nodeId)
			}
		}
		s.mutex.Unlock()
	}
}
// handle client upload request
func (s *masterServer) ClientToMasterUpload(ctx context.Context, req *pb.ClientToMasterUploadRequest) (*pb.ClientToMasterUploadResponse, error) {
	fmt.Println("Received client request")
	return &pb.ClientToMasterUploadResponse{IpAddress: "ip", Port: "8081"}, nil
}
// handle datanode notification of receiving the file
func (s *masterServer) DataNodeNotifyMaster(ctx context.Context, req *pb.DataNodeNotificationRequest) (*pb.DataNodeNotificationResponse, error) {
	fileName := req.GetFileName()
	nodeId := req.GetNodeId()
	fmt.Println("")
	node, ok := s.nodes[nodeId]
    if !ok {
        fmt.Printf("Node with ID %s not found\n", nodeId)
        return nil, fmt.Errorf("node with ID %s not found", nodeId)
    }

    // Acquire the node's mutex for thread safety
    node.mutex.Lock()
    defer node.mutex.Unlock()

    node.filenames = append(node.filenames, fileName)
    fmt.Printf("Data node %s notified the master about receiving file: %s\n", nodeId, fileName)

	clientConn, err := grpc.Dial("localhost:50001", grpc.WithInsecure())
    if err != nil {
        fmt.Println("Failed to connect to client:", err)
        // Handle connection error (e.g., retry or log)
        return nil, err
    }
    defer clientConn.Close() // Ensure master connection is closed

    masterClient := pb.NewServicesClient(clientConn)

    // Prepare notification message
    notification := &pb.MasterToClientSuccessNotifyRequest{}

    // Send notification to the master
    _, err = masterClient.MasterToClientSuccessNotify(context.Background(), notification)
    if err != nil {
        fmt.Println("Error sending notification to the client:", err)
        // Handle notification error (e.g., retry or log)
    } else {
        fmt.Println("Successfully notified client about success")
    }

    return &pb.DataNodeNotificationResponse{}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("failed to listen:", err)
		return
	}
	s := grpc.NewServer()
	masterServer := &masterServer{
		nodes: make(map[string]*nodeStatus),
	}
	pb.RegisterServicesServer(s, masterServer)
	fmt.Println("Server started. Listening on port 8080...")

	go masterServer.monitorLiveness()

	if err := s.Serve(lis); err != nil {
		fmt.Println("failed to serve:", err)
	}
}
