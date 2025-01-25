package main

import (
	"app/internal/core/cfg"
	"app/internal/core/generic/middleware"
	"app/internal/core/graph"
	"app/internal/pkg"
	res "app/internal/pkg/_resolvers"
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
	// Load application configuration.
	appCfg := cfg.Inst()

	// Initialize the application router.
	appRouter := pkg.NewRouter()

	// Initialize gRPC server and server manager.
	serverManager := gossiper.NewServerManager()
	grpcServ := grpc.NewServer()
	reflection.Register(grpcServ)
	serverManager.AddServer(gossiper.NewGRPCServ(appCfg.GrpcPort, grpcServ, appRouter.InitializeGRPCRoutes))

	var wg sync.WaitGroup
	wg.Add(1)
	// Start gRPC servers in a goroutine.
	go func() {
		defer wg.Done()
		defer serverManager.StopAll()
		serverManager.StartAll()
	}()

	// Initialize resolvers.
	resolvers, err := appRouter.InitializeRouter()
	if err != nil {
		log.Fatalf("Error initializing router: %v", err)
	}

	// Create GraphQL server.
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: resolvers.(graph.ResolverRoot),
		Directives: graph.DirectiveRoot{
			Auth: func(ctx context.Context, obj any, next graphql.Resolver) (any, error) {
				return middleware.AuthDirective(ctx, next, resolvers.(*res.Resolver).AuthMod.API.VerifyToken)
			},
		},
	}))

	// Set up the HTTP routes.
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	// Start the HTTP server.
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", appCfg.AppPort)
	log.Fatal(http.ListenAndServe(":"+appCfg.AppPort, nil))
}
