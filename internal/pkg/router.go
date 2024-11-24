package pkg

import (
	resolvers "app/internal/core/graph/resolvers"
	"app/internal/pkg/user"
)

type Router struct{}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) Init() *resolvers.Resolver {
	return &resolvers.Resolver{
		UserProvider: user.NewUserModule(),
	}
}
