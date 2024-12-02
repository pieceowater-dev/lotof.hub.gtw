package svc

import (
	"app/internal/core/cfg"
	"app/internal/core/graph/model"
	fr "app/internal/core/grpc/generated"
	"context"
	"errors"
	gossiper "github.com/pieceowater-dev/lotof.lib.gossiper/v2"
	"log"
)

type FriendshipService struct {
	transport gossiper.Transport
	client    fr.FriendshipServiceClient
}

func NewFriendshipService() *FriendshipService {
	factory := gossiper.NewTransportFactory()
	grpcTransport := factory.CreateTransport(
		gossiper.GRPC,
		cfg.Inst().LotofHubMSvcUsersGrpcAddress,
	)

	clientConstructor := fr.NewFriendshipServiceClient
	client, err := grpcTransport.CreateClient(clientConstructor)
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}

	return &FriendshipService{
		transport: grpcTransport,
		client:    client.(fr.FriendshipServiceClient),
	}
}

func (s *FriendshipService) CreateFriendship(input *fr.CreateFriendshipInput) (*fr.Friendship, error) {
	ctx := context.Background()

	response, err := s.transport.Send(ctx, s.client, "CreateFriendship", input)
	if err != nil {
		log.Printf("Error sending request: %v", err)
		return nil, err
	}

	res, ok := response.(*fr.Friendship)
	if !ok {
		return nil, errors.New("invalid response type from gRPC transport")
	}

	return res, nil
}

func (s *FriendshipService) AcceptFriendshipRequest(input *model.AcceptFriendshipInput) (*model.Friendship, error) {
	return nil, nil
}

func (s *FriendshipService) RemoveFriendship(input *model.RemoveFriendshipInput) error {
	return nil
}

func (s *FriendshipService) FriendshipList(filter *model.FriendshipFilter) (*model.PaginatedFriendshipList, error) {
	return nil, nil
}
