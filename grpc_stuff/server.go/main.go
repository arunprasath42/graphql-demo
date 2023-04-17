package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/arunprasath42/graphql-live/graph"
	"github.com/arunprasath42/graphql-live/graph/model"
	pb "github.com/arunprasath42/graphql-live/grpc_stuff"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedEmployeeServiceServer
}

func (s *server) CreateEmployee(ctx context.Context, in *pb.NewEmployee) (*pb.Employee, error) {

	r := graph.Resolver{}
	employee, err := r.Mutation().CreateEmployee(ctx, model.NewEmployee{
		Name:       in.Name,
		IsTeamLead: in.IsTeamLead,
	})
	if err != nil {
		log.Fatal(err)
	}
	return &pb.Employee{
		Id:         employee.ID,
		Name:       employee.Name,
		IsTeamLead: employee.IsTeamLead,
	}, nil

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
