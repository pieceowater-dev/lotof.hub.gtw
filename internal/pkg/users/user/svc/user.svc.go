package svc

import (
	"app/internal/core/cfg"
	user "app/internal/core/grpc/generated"
	"context"
	"errors"
	gossiper "github.com/pieceowater-dev/lotof.lib.gossiper/v2"
	"log"
)

type UserService struct {
	transport gossiper.Transport
	client    user.UserServiceClient // gRPC client
}

func NewUserService() *UserService {
	factory := gossiper.NewTransportFactory()
	grpcTransport := factory.CreateTransport(
		gossiper.GRPC,
		cfg.Inst().LotofHubMSvcUsersGrpcAddress,
	)

	// Create the client once and store it as a property
	clientConstructor := user.NewUserServiceClient
	client, err := grpcTransport.CreateClient(clientConstructor)
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}

	return &UserService{
		transport: grpcTransport,
		client:    client.(user.UserServiceClient), // Cast to the correct type
	}
}

// CreateUser handles creating a new user and returns raw data (gRPC response).
func (s *UserService) CreateUser(input *user.CreateUserRequest) (*user.User, error) {
	ctx := context.Background()

	response, err := s.transport.Send(ctx, s.client, "CreateUser", input)
	if err != nil {
		log.Printf("Error sending request: %v", err)
		return nil, err
	}

	res, ok := response.(*user.User)
	if !ok {
		return nil, errors.New("invalid response type from gRPC transport")
	}

	return res, nil
}

// FindAllUsers retrieves all users (raw gRPC response).
func (s *UserService) FindAllUsers(request *user.GetUsersRequest) (*user.GetUsersResponse, error) {
	ctx := context.Background()

	response, err := s.transport.Send(ctx, s.client, "GetUsers", request)
	if err != nil {
		log.Printf("Error sending request: %v", err)
		return nil, err
	}

	res, ok := response.(*user.GetUsersResponse)
	if !ok {
		return nil, errors.New("invalid response type from gRPC transport")
	}

	return res, nil
}

// FindOneUser retrieves a single user by ID (raw gRPC response).
func (s *UserService) FindOneUser(id string) (*user.User, error) {
	ctx := context.Background()

	request := &user.GetUserRequest{
		Id: id,
	}

	response, err := s.transport.Send(ctx, s.client, "GetUser", request)
	if err != nil {
		log.Printf("Error sending request: %v", err)
		return nil, err
	}

	res, ok := response.(*user.User)
	if !ok {
		return nil, errors.New("invalid response type from gRPC transport")
	}

	return res, nil
}

// UpdateUser updates an existing user's details and returns raw gRPC response.
func (s *UserService) UpdateUser(input *user.UpdateUserRequest) (*user.User, error) {
	ctx := context.Background()

	response, err := s.transport.Send(ctx, s.client, "UpdateUser", input)
	if err != nil {
		log.Printf("Error sending request: %v", err)
		return nil, err
	}

	res, ok := response.(*user.User)
	if !ok {
		return nil, errors.New("invalid response type from gRPC transport")
	}

	return res, nil
}
