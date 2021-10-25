package main

import (
	"log"
	"net"

	"com.github.anicolaspp/gftp/ftp"
	"google.golang.org/grpc"

	pb "com.github.anicolaspp/gftp/protos"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGftpServer(s, &ftp.Server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
