package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.56

import (
	"app/internal/core/graph/model"
	"context"
	"fmt"
)

// AddMemberToNamespace is the resolver for the addMemberToNamespace field.
func (r *mutationResolver) AddMemberToNamespace(ctx context.Context, input model.AddMemberToNamespaceInput) (*model.Namespace, error) {
	panic(fmt.Errorf("not implemented: AddMemberToNamespace - addMemberToNamespace"))
}

// RemoveMemberFromNamespace is the resolver for the removeMemberFromNamespace field.
func (r *mutationResolver) RemoveMemberFromNamespace(ctx context.Context, input model.AddMemberToNamespaceInput) (*model.Namespace, error) {
	panic(fmt.Errorf("not implemented: RemoveMemberFromNamespace - removeMemberFromNamespace"))
}

// AddMemberToService is the resolver for the addMemberToService field.
func (r *mutationResolver) AddMemberToService(ctx context.Context, input model.AddMemberToServiceInput) (*model.Service, error) {
	panic(fmt.Errorf("not implemented: AddMemberToService - addMemberToService"))
}

// RemoveMemberFromService is the resolver for the removeMemberFromService field.
func (r *mutationResolver) RemoveMemberFromService(ctx context.Context, input model.AddMemberToServiceInput) (*model.Service, error) {
	panic(fmt.Errorf("not implemented: RemoveMemberFromService - removeMemberFromService"))
}

// Members is the resolver for the members field.
func (r *queryResolver) Members(ctx context.Context, filter *model.DefaultFilterInput) ([]*model.Member, error) {
	panic(fmt.Errorf("not implemented: Members - members"))
}

// Member is the resolver for the member field.
func (r *queryResolver) Member(ctx context.Context, id string) (*model.Member, error) {
	panic(fmt.Errorf("not implemented: Member - member"))
}
