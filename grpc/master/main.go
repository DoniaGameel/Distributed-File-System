package main

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"strconv"
	"sync"
	"time"

	pb "wireless_lab_1/grpc/services" // Import the generated package

	"google.golang.org/grpc"
	"gopkg.in/yaml.v2"
)

// NodeConfig represents the structure of the configuration file
type NodeConfig struct {
	DataNodes []struct {
		ID     string `yaml:"id"`
		Port1  int    `yaml:"port_1"`
		Port2  int    `yaml:"port_2"`
		Port3  int    `yaml:"port_3"`
	} `yaml:"data_nodes"`
}

type nodeStatus struct {
	id          string
	alive       bool
	lastSeen    time.Time
	filenames   []string
	mutex       sync.Mutex
	ipAddress   string
	portNumbers []string
}

type masterServer struct {
	nodes      map[string]*nodeStatus
	nodeConfig NodeConfig // Store parsed node configuration
	nodeIDSeq  int        // Sequence for assigning node IDs
	mutex      sync.Mutex
	pb.UnimplementedServicesServer
}

// mustEmbedUnimplementedServicesServer implements services.ServicesServer.
func (s *masterServer) mustEmbedUnimplementedServicesServer() {
	panic("unimplemented")
}

// LoadNodeConfig loads node configuration from the given YAML file
func (s *masterServer) LoadNodeConfig(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(data, &s.nodeConfig)
	if err != nil {
		return err
	}
	// Initialize the nodes map here
    s.nodes = make(map[string]*nodeStatus)

    return nil
}

// Modify the Register method to save IP and assign ID
func (s *masterServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// Increment the nodeIDSeq to get the next available ID
	s.nodeIDSeq++

	// Generate a unique ID for the data node
	nodeID := strconv.Itoa(s.nodeIDSeq)

	// Find the node configuration based on node ID
	var nodePorts []int

	for _, node := range s.nodeConfig.DataNodes {
		if node.ID == nodeID {
			nodePorts = []int{node.Port1, node.Port2, node.Port3}
			break
		}
	}
	// Create a new node status with ID, IP, and other default values
	newNode := &nodeStatus{
		id:          nodeID,
		alive:       true,
		lastSeen:    time.Now(),
		filenames:   []string{},
		ipAddress:   req.GetIpAddress(),
		portNumbers: make([]string, len(nodePorts)),
	}

	// Convert port numbers to string and store them in newNode
	for i, port := range nodePorts {
		newNode.portNumbers[i] = strconv.Itoa(port)
	}

	// Add the new node to the nodes map
	s.nodes[nodeID] = newNode

	return &pb.RegisterResponse{NodeId: nodeID, PortNumbers: newNode.portNumbers}, nil
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
			if time.Since(node.lastSeen) > 10*time.Second {
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
	return &pb.ClientToMasterUploadResponse{IpAddress: "ip", Port: "50002"}, nil
}

// handle datanode notification of receiving the file
func (s *masterServer) DataNodeNotifyMaster(ctx context.Context, req *pb.DataNodeNotificationRequest) (*pb.DataNodeNotificationResponse, error) {
	fileName := req.GetFileName()
	nodeId := req.GetNodeId()
	port := req.GetPortNumber()
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

	// Choose two other nodes for replication
	replicationNodes := make([]*nodeStatus, 0)
	for _, otherNode := range s.nodes {
		if otherNode != node && !containsFile(otherNode.filenames, fileName) {
			replicationNodes = append(replicationNodes, otherNode)
			if len(replicationNodes) >= 2 {
				break
			}
		}
	}

	// Send replication requests to chosen nodes
	for _, replicaNode := range replicationNodes {
		// Implement replication request logic here
		fmt.Printf("Sending replication request to node %s for file %s\n", replicaNode.id, fileName)
		fmt.Println("Enter the port number of it ")
		var node_port string
		fmt.Scanln(&node_port)

		node_port = "localhost:" + node_port
		replicaConn, err := grpc.Dial(node_port, grpc.WithInsecure())
		if err != nil {
			fmt.Println("Failed to connect to client:", err)
			// Handle connection error (e.g., retry or log)
			return nil, err
		}
		defer replicaConn.Close() // Ensure master connection is closed

		replicaClient := pb.NewServicesClient(replicaConn)

		// Prepare notification message
		request := &pb.MasterToDataKeeperReplicaRequest{
			FileName:  fileName,
			IpAddress: "localhost",
			Port:      port,
		}

		// Send notification to the master
		_, err = replicaClient.MasterToDataKeeperReplica(context.Background(), request)
		if err != nil {
			fmt.Println("Error sending replica request:", err)
			// Handle notification error (e.g., retry or log)
		} else {
			fmt.Println("Successfully sent replica request")
		}
	}

	return &pb.DataNodeNotificationResponse{}, nil
}

func containsFile(files []string, fileName string) bool {
	for _, file := range files {
		if file == fileName {
			return true
		}
	}
	return false
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("failed to listen:", err)
		return
	}

	// Load node configuration from config.yaml
	var server masterServer
	err = server.LoadNodeConfig("config.yaml")
	if err != nil {
		fmt.Println("failed to load node configuration:", err)
		return
	}
	s := grpc.NewServer()
	pb.RegisterServicesServer(s, &server) // Register the existing server instance
	fmt.Println("Server started. Listening on port 8080...")

	go server.monitorLiveness()

	if err := s.Serve(lis); err != nil {
		fmt.Println("failed to serve:", err)
	}
}
