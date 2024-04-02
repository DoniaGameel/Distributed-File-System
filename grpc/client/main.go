package main

import (
	"context"
	"fmt"

	pb "wireless_lab_1/grpc/services" // Import the generated package

	"google.golang.org/grpc"
)

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

	// Call the RPC method
	resp, err := c.ClientToMasterUpload(context.Background(), &pb.ClientToMasterUploadRequest{})
	if err != nil {
		fmt.Println("Error calling Upload to master: ", err)
		return
	}
	fmt.Println("Received Upload Response:", resp)
	//IpAdress := resp.GetIpAddress()
	PortNumber := resp.GetPort()

	connection_port := "localhost:" + PortNumber ;
	conn_data, err := grpc.Dial(connection_port, grpc.WithInsecure())
	if err != nil {
		fmt.Println("did not connect:", err)
		return
	}
	defer conn_data.Close()

	// Create a client instance
	client_2 := pb.NewServicesClient(conn_data)

	// Call the RPC method
	resp_data, err := client_2.ClientToDataKeeperUpload(context.Background(), &pb.ClientToDataKeeperUploadRequest{FileName: "file"})
	if err != nil {
		fmt.Println("Error calling Upload to data node: ", err)
		return
	}
	fmt.Println("Received Upload Response:", resp_data)
	
}
