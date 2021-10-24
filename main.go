package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	pb "com.github.anicolaspp/gftp/protos"
	"google.golang.org/grpc"
)

func main() {
	// http.HandleFunc("/", indexHandler)

	// // [START setting_port]
	// port := os.Getenv("PORT")
	// if port == "" {
	// 	port = "8080"
	// 	log.Printf("Defaulting to port %s", port)
	// }

	// log.Printf("Listening on port %s", port)
	// if err := http.ListenAndServe(":"+port, nil); err != nil {
	// 	log.Fatal(err)
	// }

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGftpServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// func indexHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/" {
// 		http.NotFound(w, r)
// 		return
// 	}

// 	t := time.Now()

// 	fmt.Fprint(w, t)
// }

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGftpServer
}

func (s *server) Status(ctx context.Context, in *pb.StatusRequest) (*pb.StatusResponse, error) {
	log.Printf("Received: %v", in.Message)

	return &pb.StatusResponse{
		Message: fmt.Sprintf("Hello back to you, %s", in.Message),
		Time:    time.Now().Unix(),
	}, nil
}
