package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"sync"

	pb "wireless_lab_1/grpc/services" // Import the generated package

	"google.golang.org/grpc"
)

type clientListener struct {
    pb.UnimplementedServicesServer
}

func (l *clientListener) MasterToClientSuccessNotify(ctx context.Context, req *pb.MasterToClientSuccessNotifyRequest) (*pb.MasterToClientSuccessNotifyResponse, error) {
    fmt.Println("Received success notification from master")
    return &pb.MasterToClientSuccessNotifyResponse{}, nil
}

func main() {

	// Dial the master server at localhost:8080
    conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
    if err != nil {
        fmt.Println("did not connect:", err)
        return
    }
    defer conn.Close()

    // Create a client instance
    c := pb.NewServicesClient(conn)

    var wg sync.WaitGroup
    wg.Add(1) // Indicate waiting for success notification

    // Establish a port for the client to listen on
    go func() {
        lis, err := net.Listen("tcp", ":50001")
        if err != nil {
            fmt.Println("failed to listen:", err)
            return
        }
        s := grpc.NewServer()
        pb.RegisterServicesServer(s, &clientListener{}) // Register clientListener without WaitGroup
        fmt.Println("Client is Listening on port 50001...")
        if err := s.Serve(lis); err != nil {
            fmt.Println("failed to serve:", err)
        }
    }()

	var file_name string
	// Read file name from the user
	fmt.Print("Enter File Name with extension: ")
	fmt.Scanln(&file_name)
	
	var req_type int
	fmt.Print("Enter 1 to upload or 2 to download: ")
	fmt.Scanln(&req_type)

	if req_type == 1{
		// Upload logic
		// Read file content
		fileContent, err := os.ReadFile(file_name)
		if err != nil {
			// Handle error
			fmt.Println("Error reading the file", err)
			return
		}
		// Call the RPC method
		resp, err := c.ClientToMasterUpload(context.Background(), &pb.ClientToMasterUploadRequest{})
		if err != nil {
			fmt.Println("Error calling Upload to master: ", err)
			return
		}
		fmt.Println("Received Upload Response from master:", resp)
		//IpAddress := resp.GetIpAddress()
		PortNumber := resp.GetPort()
	
		connection_port := "localhost:" + PortNumber
		conn_data, err := grpc.Dial(connection_port, grpc.WithInsecure())
		if err != nil {
			fmt.Println("did not connect:", err)
			return
		}
		defer conn_data.Close()
	
		// Create a client instance
		client_2 := pb.NewServicesClient(conn_data)
	
		// Call the RPC method
		resp_data, err := client_2.ClientToDataKeeperUpload(context.Background(), &pb.ClientToDataKeeperUploadRequest{
			FileName: file_name,
			FileContent: fileContent,
			Port: PortNumber,
		})
		if err != nil {
			fmt.Println("Error calling Upload to data node: ", err)
			return
		}
		fmt.Println("Successfully Uploaded to the datanode:", resp_data)
	
		// Signal that success notification is received
		wg.Done()
	
		// Wait for success notification
		wg.Wait()
		fmt.Println("Successfully uploaded file. Disconnecting...")
		
	} else if req_type ==2 {
		// download logic
		fmt.Println("You requested type 2.")
	} else {
		fmt.Println("Unknown request type:", req_type)
		return
	}
}

