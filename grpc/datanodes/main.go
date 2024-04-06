package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"time"

	pb "wireless_lab_1/grpc/services" // Import the generated package

	"google.golang.org/grpc"
)
type dataNodeServer struct {
	pb.UnimplementedServicesServer
}

func (s *dataNodeServer) ClientToDataKeeperUpload(_ context.Context, req *pb.ClientToDataKeeperUploadRequest) (*pb.ClientToDataKeeperUploadResponse, error) {
	fileName := req.GetFileName()

    // Get the current working directory of the project
    cwd, err := os.Getwd()
    if err != nil {
        // Handle error
        return &pb.ClientToDataKeeperUploadResponse{Success: false}, err
    }

    // Specify the relative directory path (change this as needed)
    relativeDir := "copied_" + nodeId
    
    // Join the current working directory with the relative directory path
    directory := filepath.Join(cwd, relativeDir)

    // Create the directory if it doesn't exist
    if _, err := os.Stat(directory); os.IsNotExist(err) {
        if err := os.MkdirAll(directory, 0755); err != nil {
            // Handle error
            return &pb.ClientToDataKeeperUploadResponse{Success: false}, err
        }
    }

    // Join the directory with the file name to get the full file path
    filePath := filepath.Join(directory, fileName)
    
    // Write the file content to the specified file path
    err = ioutil.WriteFile(filePath, req.GetFileContent(), 0644)
    if err != nil {
        // Handle error
        fmt.Println("Error receiving the file:", fileName)
        return &pb.ClientToDataKeeperUploadResponse{Success: false}, err
    }
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
        PortNumber: req.GetPort(),
    }
    fmt.Println("nodeId ", nodeId)
    // Send notification to the master
    _, err = masterClient.DataNodeNotifyMaster(context.Background(), notification)
    if err != nil {
        fmt.Println("Error sending notification to master:", err)
        // Handle notification error (e.g., retry or log)
    } else {
        fmt.Println("Successfully notified master about received file")
    }

    return &pb.ClientToDataKeeperUploadResponse{Success: true}, nil
}

// handle replica request
func (s *dataNodeServer) MasterToDataKeeperReplica(ctx context.Context, req *pb.MasterToDataKeeperReplicaRequest) (*pb.MasterToDataKeeperReplicaResponse, error) {
	fmt.Println("Received replica request .. IpAddress: ", req.IpAddress, " file_name: ", req.FileName, " port: ", req.Port)
	
    node_port := req.IpAddress + ":" + req.Port
    replicaConn, err := grpc.Dial(node_port, grpc.WithInsecure())
    if err != nil {
        fmt.Println("Failed to connect to node:", err)
        // Handle connection error (e.g., retry or log)
        return nil, err
    }
    defer replicaConn.Close()

    replicaClient := pb.NewServicesClient(replicaConn)

    // Prepare notification message
    request := &pb.NodeToNodeReplicaRequest{
        FileName: req.FileName,
    }

    // Send notification to the master
    resp, err := replicaClient.NodeToNodeReplica(context.Background(), request)
    if err != nil {
        fmt.Println("Error sending replica request:", err)
        // Handle notification error (e.g., retry or log)
    } else {
        fmt.Println("Successfully sent replica request")
    }
    fmt.Println("Response status: ", resp.Success)

    // Get the current working directory of the project
    cwd, err := os.Getwd()
    if err != nil {
        // Handle error
        return &pb.MasterToDataKeeperReplicaResponse{}, nil
    }

    // Specify the relative directory path (change this as needed)
    relativeDir := "copied_" + nodeId
    
    // Join the current working directory with the relative directory path
    directory := filepath.Join(cwd, relativeDir)

    // Create the directory if it doesn't exist
    if _, err := os.Stat(directory); os.IsNotExist(err) {
        if err := os.MkdirAll(directory, 0755); err != nil {
            // Handle error
            return &pb.MasterToDataKeeperReplicaResponse{}, nil
        }
    }

    // Join the directory with the file name to get the full file path
    filePath := filepath.Join(directory, req.FileName)
    err = ioutil.WriteFile(filePath,resp.GetFileContent(), 0644)
    if err != nil {
        // Handle error
        fmt.Println("Error saving the file")
    }
    return &pb.MasterToDataKeeperReplicaResponse{}, nil
}

// handle replication request
func (s *dataNodeServer) NodeToNodeReplica(ctx context.Context, req *pb.NodeToNodeReplicaRequest) (*pb.NodeToNodeReplicaResponse, error) {
	fmt.Println("Received replica request")
    file_name := req.GetFileName()
    file_content, err := os.ReadFile(file_name)
    if err != nil {
        // Handle error
        fmt.Println("Error reading the file", err)
        return &pb.NodeToNodeReplicaResponse{Success: false}, nil
    }
	return &pb.NodeToNodeReplicaResponse{FileContent: file_content, Success: true}, nil
}

var nodeId string
var receivedPorts []string // Store received port numbers

func main() {
    
	// establish the node as a client
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		fmt.Println("did not connect:", err)
		return
	}
	defer conn.Close()
	c := pb.NewServicesClient(conn)

    // Call the register RPC method
    resp, err := c.Register(context.Background(), &pb.RegisterRequest{IpAddress: "localhost"})
    if err != nil {
        fmt.Println("Error calling Register to master: ", err)
        return
    }

    fmt.Println("Received ID: ", resp.GetNodeId())
    nodeId = resp.GetNodeId()
    receivedPorts = resp.GetPortNumbers() // Store port numbers

    // establish the node as a server on the received ports
    for _, port := range receivedPorts {
        go connectToNode(port) // Start a goroutine for each connection
    }

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

func connectToNode(port_number string) {
    port_number = ":" + port_number
    lis, err := net.Listen("tcp", port_number)
        if err != nil {
            fmt.Println("failed to listen:", err)
            return
        }
        s := grpc.NewServer()
        pb.RegisterServicesServer(s, &dataNodeServer{})
        fmt.Println("Server started. Listening on port ", port_number, "...")
        if err := s.Serve(lis); err != nil {
            fmt.Println("failed to serve:", err)
        }
}