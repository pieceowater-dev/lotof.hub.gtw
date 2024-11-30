package svc

import (
	"app/internal/core/cfg"
	"app/internal/core/grpc/generated"
	"context"
	gossiper "github.com/pieceowater-dev/lotof.lib.gossiper/v2"
	"log"
)

type AuthService struct {
	client generated.AuthServiceClient
}

func NewAuthService() *AuthService {
	factory := gossiper.NewTransportFactory()
	grpcTransport := factory.CreateTransport(
		gossiper.GRPC,
		cfg.Inst().LotofHubMSvcUsersGrpcAddress,
	)

	clientConstructor := generated.NewAuthServiceClient
	client, err := grpcTransport.CreateClient(clientConstructor)
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}

	return &AuthService{
		client: client.(generated.AuthServiceClient),
	}
}

func (s *AuthService) Login(ctx context.Context, input *generated.LoginRequest) (*generated.AuthResponse, error) {
	return s.client.Login(ctx, input)
}

func (s *AuthService) Register(ctx context.Context, input *generated.RegisterRequest) (*generated.AuthResponse, error) {
	return s.client.Register(ctx, input)
}
