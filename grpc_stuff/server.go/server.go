package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/arunprasath42/graphql-live/grpc_stuff"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedEmployeeServiceServer
}

func (s *server) CreateEmployee(ctx context.Context, in *pb.NewEmployee) (*pb.Employee, error) {
	return &pb.Employee{Id: "1", Name: in.GetName(), IsTeamLead: in.GetIsTeamLead()}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50051))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterEmployeeServiceServer(s, &server{})
	log.Printf("Server listening on port %d", 50051)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
