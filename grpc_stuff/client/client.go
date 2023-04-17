package client

import (
	"log"

	rpc_stuff "github.com/arunprasath42/graphql-live/grpc_stuff"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ClientConn() rpc_stuff.EmployeeServiceClient {
	conn, err := grpc.Dial(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := rpc_stuff.NewEmployeeServiceClient(conn)
	return c
}
