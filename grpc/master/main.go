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
	portStatus  []bool // Add a slice to store busy status for each port
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

type fileNodesMap struct {
    sync.RWMutex
    data map[string][]string // Map to store file names and corresponding node IDs
}

// Initialize the map
func newFileNodesMap() *fileNodesMap {
    return &fileNodesMap{
        data: make(map[string][]string),
    }
}

// Add a file and its corresponding node ID to the map
func (m *fileNodesMap) addFileNode(fileName string, nodeID string) {
    m.Lock()
    defer m.Unlock()
    m.data[fileName] = append(m.data[fileName], nodeID)
}

// Get the list of node IDs associated with a file
func (m *fileNodesMap) getNodeIDs(fileName string) []string {
    m.RLock()
    defer m.RUnlock()
    return m.data[fileName]
}

// Initialize file-node map
var fileNodes = newFileNodesMap()

func (s *masterServer) getSourceMachine(fileName string) (string, error) {
    // Get the list of node IDs associated with the file from the fileNodes map
    nodeIDs := fileNodes.getNodeIDs(fileName)

    if len(nodeIDs) == 0 {
        return "", fmt.Errorf("file %s not found or not stored on any node", fileName)
    }

    // For simplicity, return the first node ID where the file is stored
    return nodeIDs[0], nil
}

func (s *masterServer) selectMachineToCopyTo(fileName string) (string, error) {
    // Iterate over available nodes and select one that does not contain the file
    for _, node := range s.nodes {
        if !containsFile(node.filenames, fileName) && node.alive {
            
            // Return the ID of the selected node
            return node.id, nil
        }
    }

    // If no suitable node is found, return an error
    return "", fmt.Errorf("no available machine found to copy the file")
}


func (s *masterServer) getNodeIPAndPorts(nodeID string) (string, []string, error) {
    // Search for the node with the provided ID
    node, ok := s.nodes[nodeID]
    if !ok {
        return "", nil, fmt.Errorf("node with ID %s not found", nodeID)
    }

    // Return the IP address and port numbers of the node
    return node.ipAddress, node.portNumbers, nil
}

func (s *masterServer) notifyMachineDataTransfer(sourceMachine string, destinationMachine string, fileName string) {
    // Get IP and port numbers for source and destination machines
    sourceIP, sourcePortNumbers, err := s.getNodeIPAndPorts(sourceMachine)
	fmt.Println("hello 34")
    if err != nil {
        fmt.Println("Error getting IP and port numbers for source machine:", err)
        return
    }

    destIP, destPortNumbers, err := s.getNodeIPAndPorts(destinationMachine)
    if err != nil {
        fmt.Println("Error getting IP and port numbers for destination machine:", err)
        return
    }

    // Find non-busy ports for both source and destination machines
    var sourcePort, destPort string
    for i, portBusy := range s.nodes[sourceMachine].portStatus {
        if !portBusy {
            sourcePort = sourcePortNumbers[i]
            break
        }
    }

    for i, portBusy := range s.nodes[destinationMachine].portStatus {
        if !portBusy {
            destPort = destPortNumbers[i]
            break
        }
    }

    // Check if both source and destination ports are found
    if sourcePort == "" || destPort == "" {
        fmt.Println("No available non-busy ports for source or destination machines")
        return
    }

    // Perform data transfer notification using selected ports
    fmt.Printf("Transfer initiated from %s:%s to %s:%s for file: %s\n", sourceIP, sourcePort, destIP, destPort, fileName)
}

// Function to check replication based on the specified algorithm
func (s *masterServer) checkReplication() {
    ticker := time.NewTicker(10 * time.Second)
    defer ticker.Stop()
    for range ticker.C {
        s.mutex.Lock()
        for fileName := range fileNodes.data {
            // Get list of node IDs associated with the file
            nodeIDs := fileNodes.getNodeIDs(fileName)
			fmt.Println("file name", fileName)
			fmt.Println("len(nodeIDs)", len(nodeIDs))
            if len(nodeIDs) < 3 {
                // Get source machine for replication
                sourceMachine, err := s.getSourceMachine(fileName)
				if(err != nil){
					fmt.Println("Error getting the source Machine ", err)
				}
                // Continue replication until the file is replicated on at least 3 alive data nodes
                for len(nodeIDs) < 3 {
                    // Select destination machine for replication
                    destinationMachine, err := s.selectMachineToCopyTo(fileName)
					if(err != nil){
						fmt.Println("Error getting a destination Machine ", err)
					}
                    // Notify data transfer
                    s.notifyMachineDataTransfer(sourceMachine, destinationMachine, fileName)
                    // Update fileNodes map with replicated node ID
                    fileNodes.addFileNode(fileName, destinationMachine)
                    // Get updated list of node IDs associated with the file
                    nodeIDs = fileNodes.getNodeIDs(fileName)
                }
            }
        }
        s.mutex.Unlock()
    }
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
		portStatus:  make([]bool, len(nodePorts)), // Initialize portStatus slice
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

    s.mutex.Lock()
    defer s.mutex.Unlock()

    // Choose a random node for upload, prioritizing non-busy ports
    var chosenNode *nodeStatus
    var chosenPortIndex int = -1

    for _, node := range s.nodes {
        if node.alive {
            for i, portBusy := range node.portStatus {
                if !portBusy { // Check for a non-busy port
                    chosenNode = node
                    chosenPortIndex = i
                    break  // Exit inner loop if a non-busy port is found
                }
            }
            if chosenNode != nil { // Exit outer loop if a non-busy node is found
                break
            }
        }
    }
	
    // Check if a live node with a non-busy port was found
    if chosenNode == nil || chosenPortIndex == -1 {
        fmt.Println("No available nodes or ports to upload to")
        return nil, fmt.Errorf("no available nodes or ports to upload to")
    }

    // Mark chosen port as busy
    chosenNode.portStatus[chosenPortIndex] = true

    // Prepare response with chosen IP and port
    return &pb.ClientToMasterUploadResponse{
        IpAddress: chosenNode.ipAddress,
        Port:      chosenNode.portNumbers[chosenPortIndex],
    }, nil
}



// handle datanode notification of receiving the file
func (s *masterServer) DataNodeNotifyMaster(ctx context.Context, req *pb.DataNodeNotificationRequest) (*pb.DataNodeNotificationResponse, error) {
	fileName := req.GetFileName()
	nodeId := req.GetNodeId()
	// Update fileNodes map with file-node association
    fileNodes.addFileNode(fileName, nodeId)
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

	// Choose two other nodes for replication, considering non-busy ports
	replicationNodes := make([]*nodeStatus, 0)
	for _, otherNode := range s.nodes {
		if otherNode != node && !containsFile(otherNode.filenames, fileName) {
			// Check for non-busy ports before adding the node
			for _, portBusy := range otherNode.portStatus {
				if !portBusy {
					replicationNodes = append(replicationNodes, otherNode)
					break // Break out of the inner loop if a non-busy port is found
				}
			}
			if len(replicationNodes) >= 2 {
				break
			}
		}
	}

	// Send replication requests to chosen nodes
	for _, replicaNode := range replicationNodes {
		// Find the index of a non-busy port in the chosen replica node
		var chosenPortIndex int = -1
		for i, portBusy := range replicaNode.portStatus {
			if !portBusy {
				chosenPortIndex = i
				break // Exit inner loop if a non-busy port is found
			}
		}

		// Check if a non-busy port was found
		if chosenPortIndex == -1 {
			fmt.Printf("No available non-busy ports on node %s for replication of %s\n", replicaNode.id, fileName)
			continue // Skip to the next node if no non-busy port is available
		}

		// Construct the connection address using replicaNode information
		replicaAddress := "localhost:" + replicaNode.portNumbers[chosenPortIndex]
		// Establish connection to the chosen replica node
		replicaConn, err := grpc.Dial(replicaAddress, grpc.WithInsecure())
		if err != nil {
			fmt.Println("Failed to connect to replica node:", err)
			// Handle connection error (e.g., retry or log)
			continue // Skip to the next node on connection failure
		}
		defer replicaConn.Close() // Ensure connection is closed after use

		// Create a gRPC client for the chosen replica node
		replicaClient := pb.NewServicesClient(replicaConn)

		// Prepare the replication request message
		request := &pb.MasterToDataKeeperReplicaRequest{
			FileName:  fileName,
			IpAddress: node.ipAddress, // Source node IP for replication
			Port:      port,           // Source node port for replication
		}

		// Send the replication request to the chosen replica node
		resp, err := replicaClient.MasterToDataKeeperReplica(context.Background(), request)
		if err != nil {
			fmt.Println("Error sending replica request:", err)
			// Handle notification error (e.g., retry or log)
		} else {
			fmt.Printf("Successfully sent replica request to node %s for file %s\n", replicaNode.id, fileName)
		}
		if resp.Success{
			fileNodes.addFileNode(fileName, replicaNode.id)
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
	// Start replication check routine
	//go server.checkReplication()

	if err := s.Serve(lis); err != nil {
		fmt.Println("failed to serve:", err)
	}
}
