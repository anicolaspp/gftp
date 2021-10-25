// Package ftp contains the FTP interface.
package ftp

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "com.github.anicolaspp/gftp/protos"
)

type Server struct {
	pb.UnimplementedGftpServer
}

func (s *Server) Status(ctx context.Context, in *pb.StatusRequest) (*pb.StatusResponse, error) {
	log.Printf("Received: %v", in.Message)

	return &pb.StatusResponse{
		Message: fmt.Sprintf("Hello back to you, %s", in.Message),
		Time:    time.Now().Unix(),
	}, nil
}
