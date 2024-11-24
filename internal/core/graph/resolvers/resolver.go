package graph

import (
	"app/internal/pkg/user"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UserProvider user.Module
}
