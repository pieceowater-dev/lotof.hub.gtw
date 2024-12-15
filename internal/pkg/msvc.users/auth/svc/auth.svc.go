package svc

import (
	"app/internal/core/cfg"
	auth "app/internal/core/grpc/generated"
	"context"
	"errors"
	"fmt"
	gossiper "github.com/pieceowater-dev/lotof.lib.gossiper/v2"
	"log"
)

type AuthService struct {
	transport gossiper.Transport
	client    auth.AuthServiceClient
}

func NewAuthService() *AuthService {
	factory := gossiper.NewTransportFactory()
	grpcTransport := factory.CreateTransport(
		gossiper.GRPC,
		cfg.Inst().LotofHubMSvcUsersGrpcAddress,
	)

	clientConstructor := auth.NewAuthServiceClient
	client, err := grpcTransport.CreateClient(clientConstructor)
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}

	return &AuthService{
		transport: grpcTransport,
		client:    client.(auth.AuthServiceClient),
	}
}

func (s *AuthService) VerifyToken(token string) (*auth.ValidateTokenResponse, error) {
	ctx := context.Background()
	req := &auth.ValidateTokenRequest{
		Token: token,
	}

	resp, err := s.client.ValidateToken(ctx, req)
	if err != nil {
		log.Printf("Error validating token: %v", err)
		return nil, err
	}

	if !resp.Valid {
		return nil, fmt.Errorf("token is invalid: %v", resp.Message)
	}

	return resp, nil
}

func (s *AuthService) Login(input *auth.LoginRequest) (*auth.AuthResponse, error) {
	ctx := context.Background()

	response, err := s.transport.Send(ctx, s.client, "Login", input)
	if err != nil {
		log.Printf("Error sending request: %v", err)
		return nil, err
	}

	res, ok := response.(*auth.AuthResponse)
	if !ok {
		return nil, errors.New("invalid response type from gRPC transport")
	}

	return res, nil
}

func (s *AuthService) Register(input *auth.RegisterRequest) (*auth.AuthResponse, error) {
	ctx := context.Background()

	response, err := s.transport.Send(ctx, s.client, "Register", input)
	if err != nil {
		log.Printf("Error sending request: %v", err)
		return nil, err
	}

	res, ok := response.(*auth.AuthResponse)
	if !ok {
		return nil, errors.New("invalid response type from gRPC transport")
	}

	return res, nil
}