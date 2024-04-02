package main

import (
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
