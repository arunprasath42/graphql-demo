package client

import (
	"log"

	rpc_stuff "github.com/arunprasath42/graphql-live/grpc_stuff"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ClientConn() rpc_stuff.EmployeeServiceClient {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := rpc_stuff.NewEmployeeServiceClient(conn)

	// Contact the server and print out its response.
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// defer cancel()

	// r, err := c.CreateEmployee(ctx, &rpc_stuff.NewEmployee{Name: "Nikil Joshi", IsTeamLead: true})
	// if err != nil {
	// 	log.Fatalf("could not greet: %v", err)
	// }
	// log.Printf("Greeting: %s", r.GetId())
	return c
}
