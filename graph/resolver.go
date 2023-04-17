package graph

import (
	rpc_stuff "github.com/arunprasath42/graphql-live/grpc_stuff"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	EmployeeServiceClient rpc_stuff.EmployeeServiceClient
}
