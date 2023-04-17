package main

import (
	"flag"
	"log"
	"net"

	pb "github.com/arunprasath42/graphql-live/grpc_stuff"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedEmployeeServiceServer
}

//the grpc server is running on port 50051 and it will be triggered when a mutation is called

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterEmployeeServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
