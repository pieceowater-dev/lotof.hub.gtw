package pkg

import (
	"app/internal/core/generic/interfaces"
	pb "app/internal/core/grpc/generated/lotof.hub.gtw/gtw"
	resolvers "app/internal/pkg/_resolvers"
	"app/internal/pkg/msvc.namespaces/member"
	"app/internal/pkg/msvc.namespaces/ns"
	"app/internal/pkg/msvc.users/auth"
	"app/internal/pkg/msvc.users/friendship"
	"app/internal/pkg/msvc.users/user"
	"google.golang.org/grpc"
	"reflect"
)

type Router struct {
	modules map[string]interfaces.IModule // Map of module names to their instances.
}

func NewRouter() *Router {
	userModule := user.New()
	authModule := auth.New()
	friendshipModule := friendship.New()
	nsModule := ns.New()
	memberModule := member.New()
	return &Router{
		modules: map[string]interfaces.IModule{
			userModule.Name():       userModule,
			authModule.Name():       authModule,
			friendshipModule.Name(): friendshipModule,
			nsModule.Name():         nsModule,
			memberModule.Name():     memberModule,
		},
	}
}

// InitializeRouter initializes all modules and returns the GraphQL resolver.
func (r *Router) InitializeRouter() (any, error) {
	// Initialize all modules
	resolver := r.initializeGQLResolvers()
	// r.initializeGRPCRoutes(r.server)
	return resolver, nil
}

// initializeGQLResolvers initializes the GraphQL resolvers for all modules.
func (r *Router) initializeGQLResolvers() *resolvers.Resolver {
	resolver := &resolvers.Resolver{}
	resolverValue := reflect.ValueOf(resolver).Elem()

	for i := 0; i < resolverValue.NumField(); i++ {
		field := resolverValue.Type().Field(i)
		moduleName := field.Name
		if module, ok := r.modules[moduleName]; ok {
			resolverValue.Field(i).Set(reflect.ValueOf(module))
		}
	}

	return resolver
}

// InitializeGRPCRoutes initializes the gRPC routes for all modules.
func (r *Router) InitializeGRPCRoutes(server *grpc.Server) {
	pb.RegisterGatewayServiceServer(server, r.modules["NamespacesMod"].(ns.Module).API)
}

// GetModules returns the map of modules.
func (r *Router) GetModules() map[string]interfaces.IModule {
	return r.modules
}
