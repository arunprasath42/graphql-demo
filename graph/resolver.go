package graph

import (
	"context"
	"log"

	"github.com/arunprasath42/graphql-live/graph/model"
	pb "github.com/arunprasath42/graphql-live/grpc_stuff"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Resolvers implements the GraphQL resolver functions.
type Resolver struct{}

// sendToGRPCServer sends the given employee to the gRPC server.
func sendToGRPCServer(employee *model.Employee) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	c := pb.NewEmployeeServiceClient(conn)
	_, err = c.CreateEmployee(context.Background(), &pb.NewEmployee{
		Name:       employee.Name,
		IsTeamLead: employee.IsTeamLead,
	})
	if err != nil {
		log.Fatalf("failed to send to gRPC server: %v", err)
	}
}
