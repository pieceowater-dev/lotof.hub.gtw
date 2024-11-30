package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.56

import (
	"app/internal/core/graph"
	"context"
	"fmt"
)

// Placeholder is the resolver for the _placeholder field.
func (r *mutationResolver) Placeholder(ctx context.Context) (*string, error) {
	panic(fmt.Errorf("not implemented: Placeholder - _placeholder"))
}

// Placeholder is the resolver for the _placeholder field.
func (r *queryResolver) Placeholder(ctx context.Context) (*string, error) {
	panic(fmt.Errorf("not implemented: Placeholder - _placeholder"))
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
