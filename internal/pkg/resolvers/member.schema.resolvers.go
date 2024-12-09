package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.56

import (
	"app/internal/core/graph/model"
	"context"
)

// AddMemberToNamespace is the resolver for the addMemberToNamespace field.
func (r *mutationResolver) AddMemberToNamespace(ctx context.Context, input model.MemberToNamespaceInput) (*model.Namespace, error) {
	return r.MemberProvider.API.AddMemberToNamespace(ctx, input)
}

// RemoveMemberFromNamespace is the resolver for the removeMemberFromNamespace field.
func (r *mutationResolver) RemoveMemberFromNamespace(ctx context.Context, input model.MemberToNamespaceInput) (*model.Namespace, error) {
	return r.MemberProvider.API.RemoveMemberFromNamespace(ctx, input)
}

// Members is the resolver for the members field.
func (r *queryResolver) Members(ctx context.Context, filter *model.MembersFilter) ([]*model.Member, error) {
	return r.MemberProvider.API.Members(ctx, filter)
}

// Member is the resolver for the member field.
func (r *queryResolver) Member(ctx context.Context, membershipID string) (*model.Member, error) {
	return r.MemberProvider.API.Member(ctx, membershipID)
}
