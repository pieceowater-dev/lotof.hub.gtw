package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.56

import (
	"app/internal/core/graph/model"
	"context"
)

// CreateNamespace is the resolver for the createNamespace field.
func (r *mutationResolver) CreateNamespace(ctx context.Context, input model.NamespaceInput) (*model.Namespace, error) {
	return r.NSProvider.API.CreateNamespace(ctx, input)
}

// UpdateNamespace is the resolver for the updateNamespace field.
func (r *mutationResolver) UpdateNamespace(ctx context.Context, id string, input model.NamespaceInput) (*model.Namespace, error) {
	return r.NSProvider.API.UpdateNamespace(ctx, input)
}

// Namespaces is the resolver for the namespaces field.
func (r *queryResolver) Namespaces(ctx context.Context, filter *model.DefaultFilterInput) (*model.PaginatedNamespaceList, error) {
	return r.NSProvider.API.Namespaces(ctx, filter)
}

// Namespace is the resolver for the namespace field.
func (r *queryResolver) Namespace(ctx context.Context, id string) (*model.Namespace, error) {
	return r.NSProvider.API.Namespace(ctx, id)
}
