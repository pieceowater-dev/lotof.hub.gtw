package pkg

import (
	"app/internal/pkg/msvc.namespaces/member"
	"app/internal/pkg/msvc.namespaces/ns"
	"app/internal/pkg/msvc.users/auth"
	"app/internal/pkg/msvc.users/friendship"
	"app/internal/pkg/msvc.users/user"
	resolvers "app/internal/pkg/resolvers"
)

type Router struct{}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) Init() *resolvers.Resolver {
	return &resolvers.Resolver{
		UserProvider:       user.NewUserModule(),
		AuthProvider:       auth.NewAuthModule(),
		FriendshipProvider: friendship.NewFriendshipModule(),

		NSProvider:     ns.NewNSModule(),
		MemberProvider: member.NewMemberModule(),
	}
}

// if this gateway serves as grpc server somehow uncomment below
//// InitGRPC initializes gRPC routes
//func (r *Router) InitGRPC(grpcServer *grpc.Server) {
//	// Register gRPC services
//	//pb.RegisterUserServiceServer(grpcServer, r.userModule.Controller)
//}
