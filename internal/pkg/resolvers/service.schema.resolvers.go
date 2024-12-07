package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.56

import (
	"app/internal/core/graph/model"
	"context"
	"fmt"
)

// CreateService is the resolver for the createService field.
func (r *mutationResolver) CreateService(ctx context.Context, input model.ServiceInput) (*model.Service, error) {
	return r.ServiceProvider.API.CreateService(ctx, input)
}

// UpdateService is the resolver for the updateService field.
func (r *mutationResolver) UpdateService(ctx context.Context, id string, input model.ServiceInput) (*model.Service, error) {
	return r.ServiceProvider.API.UpdateService(ctx, id, input)
}

// DeleteService is the resolver for the deleteService field.
func (r *mutationResolver) DeleteService(ctx context.Context, id string) (bool, error) {
	return r.ServiceProvider.API.DeleteService(ctx, id)
}

// Services is the resolver for the services field.
func (r *queryResolver) Services(ctx context.Context, filter *model.DefaultFilterInput) (*model.PaginatedServiceList, error) {
	return r.ServiceProvider.API.Services(ctx, filter)
}

// ServicesByNamespace is the resolver for the servicesByNamespace field.
func (r *queryResolver) ServicesByNamespace(ctx context.Context, nsID string, filter *model.DefaultFilterInput) (*model.PaginatedServiceList, error) {
	return r.ServiceProvider.API.ServicesByNamespace(ctx, nsID, filter)
}

// Service is the resolver for the service field.
func (r *queryResolver) Service(ctx context.Context, id string) (*model.Service, error) {
	panic(fmt.Errorf("not implemented: Service - service"))
}
