package graph

import (
	"app/internal/pkg/msvc.namespaces/member"
	"app/internal/pkg/msvc.namespaces/ns"
	"app/internal/pkg/msvc.users/auth"
	"app/internal/pkg/msvc.users/friendship"
	"app/internal/pkg/msvc.users/user"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	AuthMod        auth.Module
	UsersMod       user.Module
	FriendshipsMod friendship.Module

	NamespacesMod ns.Module
	MembersMod    member.Module
}
