package pkg

import (
	resolvers "app/internal/pkg/resolvers"
	"app/internal/pkg/users/friendship"
	"app/internal/pkg/users/user"
)

type Router struct{}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) Init() *resolvers.Resolver {
	return &resolvers.Resolver{
		UserProvider: user.NewUserModule(),
		//AuthProvider:       auth.NewAuthModule(),
		FriendshipProvider: friendship.NewFriendshipModule(),
	}
}

// if this gateway serves as grpc server somehow uncomment below
//// InitGRPC initializes gRPC routes
//func (r *Router) InitGRPC(grpcServer *grpc.Server) {
//	// Register gRPC services
//	//pb.RegisterUserServiceServer(grpcServer, r.userModule.Controller)
//}
