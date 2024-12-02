package graph

import (
	"app/internal/pkg/users/friendship"
	"app/internal/pkg/users/user"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UserProvider user.Module
	//AuthProvider       auth.Module
	FriendshipProvider friendship.Module
}
