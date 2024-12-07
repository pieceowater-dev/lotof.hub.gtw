package graph

import (
	"app/internal/pkg/namespaces/member"
	"app/internal/pkg/namespaces/ns"
	"app/internal/pkg/namespaces/service"
	"app/internal/pkg/users/auth"
	"app/internal/pkg/users/friendship"
	"app/internal/pkg/users/user"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UserProvider       user.Module
	AuthProvider       auth.Module
	FriendshipProvider friendship.Module

	NSProvider      ns.Module
	MemberProvider  member.Module
	ServiceProvider service.Module
}
