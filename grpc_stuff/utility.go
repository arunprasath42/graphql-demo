package grpc_stuff

import (
	"context"
	"log"
	"time"

	"github.com/arunprasath42/graphql-live/database"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var db = database.Connect()

func NewClientConn(addr, name string) *grpc.ClientConn {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("did not connect: %v\n", err)
		panic("couldn't connect")
	}
	defer conn.Close()

	return conn
}

func ClientConn(addr, name string) string {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("did not connect: %v\n", err)
	}

	defer conn.Close()

	c := NewEmployeeServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.CreateEmployee(ctx, &NewEmployee{Name: name, IsTeamLead: true})
	if err != nil {
		log.Printf("could create employee: %v\n", err)
	}

	return r.Name
}
