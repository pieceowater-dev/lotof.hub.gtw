package svc

import (
	"app/internal/core/cfg"
	"app/internal/core/graph/model"
	"app/internal/core/grpc/generated"
	"context"
	gossiper "github.com/pieceowater-dev/lotof.lib.gossiper/v2"
	"log"
)

type FriendshipService struct {
	client generated.FriendshipServiceClient
}

func NewFriendshipService() *FriendshipService {
	factory := gossiper.NewTransportFactory()
	grpcTransport := factory.CreateTransport(
		gossiper.GRPC,
		cfg.Inst().LotofHubMSvcUsersGrpcAddress,
	)

	clientConstructor := generated.NewFriendshipServiceClient
	client, err := grpcTransport.CreateClient(clientConstructor)
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}

	return &FriendshipService{
		client: client.(generated.FriendshipServiceClient),
	}
}

func (s *FriendshipService) CreateFriendship(ctx context.Context, input *model.CreateFriendshipInput) (*model.Friendship, error) {
	//return s.client.CreateFriendship(ctx, input)
	return nil, nil
}

func (s *FriendshipService) AcceptFriendshipRequest(ctx context.Context, input *model.AcceptFriendshipInput) (*model.Friendship, error) {
	//return s.client.AcceptFriendshipRequest(ctx, input)
	return nil, nil
}

func (s *FriendshipService) RemoveFriendship(ctx context.Context, input *model.RemoveFriendshipInput) error {
	//_, err := s.client.RemoveFriendshipRequest(ctx, input)
	//return err
	return nil
}

func (s *FriendshipService) FriendshipList(ctx context.Context, filter *model.FriendshipFilter) (*model.PaginatedFriendshipList, error) {
	//return s.client.FriendshipRequestList(ctx, filter)
	return nil, nil
}
