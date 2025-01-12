package pkg

import (
	pb "app/internal/core/grpc/generated"
	"app/internal/pkg/msvc.namespaces/member"
	"app/internal/pkg/msvc.namespaces/ns"
	"app/internal/pkg/msvc.users/auth"
	"app/internal/pkg/msvc.users/friendship"
	"app/internal/pkg/msvc.users/user"
	resolvers "app/internal/pkg/resolvers"
	"google.golang.org/grpc"
)

type Router struct {
	userModule       user.Module
	authModule       auth.Module
	friendshipModule friendship.Module
	nsModule         ns.Module
	memberModule     member.Module
}

func NewRouter() *Router {
	return &Router{
		userModule:       user.NewUserModule(),
		authModule:       auth.NewAuthModule(),
		friendshipModule: friendship.NewFriendshipModule(),
		nsModule:         ns.NewNSModule(),
		memberModule:     member.NewMemberModule(),
	}
}

func (r *Router) Init() *resolvers.Resolver {
	return &resolvers.Resolver{
		UserProvider:       r.userModule,
		AuthProvider:       r.authModule,
		FriendshipProvider: r.friendshipModule,

		NSProvider:     r.nsModule,
		MemberProvider: r.memberModule,
	}
}

// InitGRPC initializes gRPC routes
func (r *Router) InitGRPC(grpcServer *grpc.Server) {
	// Register gRPC services
	pb.RegisterGatewayServiceServer(grpcServer, r.nsModule.API)
}
