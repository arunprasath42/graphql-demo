package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/arunprasath42/graphql-live/database"
	"github.com/arunprasath42/graphql-live/graph/model"
	pb "github.com/arunprasath42/graphql-live/grpc_stuff"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	db *database.DB
	pb.UnimplementedEmployeeServiceServer
}

func newServer(db database.DB) *server {
	return &server{
		db: &db,
	}
}

//CreateEmployee creates a new employee in the database
func (s *server) CreateEmployee(ctx context.Context, in *pb.NewEmployee) (*pb.Employee, error) {

	employee, err := s.db.CreateEmployee(model.NewEmployee{
		Name:       in.Name,
		IsTeamLead: in.IsTeamLead,
	})
	if err != nil {
		return nil, err
	}

	return &pb.Employee{
		Name:       employee.Name,
		IsTeamLead: employee.IsTeamLead,
	}, nil
}

// UpdateEmployee updates an existing employee in the database
func (s *server) UpdateEmployee(ctx context.Context, in *pb.UpdateEmployeeRequest) (*pb.Employee, error) {
	employee, err := s.db.UpdateEmployee(in.Id, model.NewEmployee{
		Name:       in.Input.Name,
		IsTeamLead: in.Input.IsTeamLead,
	})
	if err != nil {
		return nil, err
	}

	return &pb.Employee{
		Name:       employee.Name,
		IsTeamLead: employee.IsTeamLead,
	}, nil
}

// DeleteEmployee deletes an existing employee in the database
func (s *server) DeleteEmployee(ctx context.Context, in *pb.DeleteEmployeeByIdRequest) (*pb.Employee, error) {
	employee, err := s.db.DeleteEmployee(in.Id)
	if err != nil {
		return nil, err
	}

	return &pb.Employee{
		Name:       employee.Name,
		IsTeamLead: employee.IsTeamLead,
	}, nil
}

// GetEmployee gets an existing employee in the database
func (s *server) GetEmployee(ctx context.Context, in *pb.GetEmployeeByIdRequest) (*pb.Employee, error) {
	employee, err := s.db.GetEmployee(in.Id)
	if err != nil {
		return nil, err
	}

	return &pb.Employee{
		Name:       employee.Name,
		IsTeamLead: employee.IsTeamLead,
	}, nil
}

func main() {
	fmt.Println("Starting server...")
	flag.Parse()

	db := database.Connect()

	fmt.Println("Connected to DB", db)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	fmt.Println("Listening on port", *port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	fmt.Println("Registering server")

	pb.RegisterEmployeeServiceServer(s, newServer(*db))
	fmt.Println("Server registered")

	// Register reflection service on gRPC server.
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
