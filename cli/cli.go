// Package cli contains the cli implementation to connect to our server.
package main

import (
	"context"
	"log"

	pb "com.github.anicolaspp/gftp/protos"
	"google.golang.org/grpc"
)

const (
	address     = ":50051"
	defaultName = "world"
)

func main() {
	log.Println("Running...")
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	log.Println("Connection open...")
	c := pb.NewGftpClient(conn)

	req := &pb.StatusRequest{
		Message: "nick",
	}

	log.Println("Making request...")
	res, err := c.Status(context.Background(), req)
	if err != nil {
		log.Fatalf("Error communicating with %s, %v", address, err)
	}

	log.Printf("Response from server: %s", res)
}
