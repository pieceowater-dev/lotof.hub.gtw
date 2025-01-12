package main

import (
	"app/internal/core/cfg"
	"app/internal/core/graph"
	"app/internal/pkg"
	"app/internal/pkg/msvc.users/auth/middleware"
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	gossiper "github.com/pieceowater-dev/lotof.lib.gossiper/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net/http"
	"sync"
)

func main() {
	//todo: clean up later
	appCfg := cfg.Inst()
	appRouter := pkg.NewRouter()

	serverManager := gossiper.NewServerManager()
	grpcServ := grpc.NewServer()
	reflection.Register(grpcServ)
	serverManager.AddServer(gossiper.NewGRPCServ(appCfg.GrpcPort, grpcServ, appRouter.InitGRPC))
	var wg sync.WaitGroup
	wg.Add(1)
	// Start gRPC servers in a goroutine
	go func() {
		defer wg.Done()
		serverManager.StartAll()
	}()

	// Initialize resolvers
	resolvers := appRouter.Init()

	// Create GraphQL server
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: resolvers,
		Directives: graph.DirectiveRoot{
			Auth: func(ctx context.Context, obj any, next graphql.Resolver) (any, error) {
				return middleware.AuthDirective(ctx, next, resolvers.AuthProvider.API.VerifyToken)
			},
		},
	}))

	// Apply middleware to the GraphQL server
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", appCfg.AppPort)
	log.Fatal(http.ListenAndServe(":"+appCfg.AppPort, nil))
}
