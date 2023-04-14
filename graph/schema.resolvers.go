package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.29

import (
	"context"

	"github.com/arunprasath42/graphql-live/database"
	"github.com/arunprasath42/graphql-live/graph/model"
)

var db = database.Connect()

// CreateEmployee is the resolver for the createEmployee field.
func (r *mutationResolver) CreateEmployee(ctx context.Context, input model.NewEmployee) (*model.Employee, error) {
	return db.CreateEmployee(input)
}

// UpdateEmployee is the resolver for the updateEmployee field.
func (r *mutationResolver) UpdateEmployee(ctx context.Context, id string, input model.NewEmployee) (*model.Employee, error) {
	return db.UpdateEmployee(id, input)
}

// DeleteEmployee is the resolver for the deleteEmployee field.
func (r *mutationResolver) DeleteEmployee(ctx context.Context, id string) (*model.Employee, error) {
	return db.DeleteEmployee(id)
}

// GetEmployee is the resolver for the getEmployee field.
func (r *queryResolver) GetEmployee(ctx context.Context, id string) (*model.Employee, error) {
	return db.GetEmployee(id)
}

// GetEmployees is the resolver for the getEmployees field.
func (r *queryResolver) GetEmployees(ctx context.Context) ([]*model.Employee, error) {
	return db.GetEmployees()
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
