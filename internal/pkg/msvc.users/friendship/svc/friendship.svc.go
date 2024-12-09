package svc

import (
	"app/internal/core/cfg"
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

func (s *FriendshipService) AcceptFriendshipRequest(input *fr.AcceptFriendshipInput) (*fr.Friendship, error) {
	ctx := context.Background()

	response, err := s.transport.Send(ctx, s.client, "AcceptFriendshipRequest", input)
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

func (s *FriendshipService) RemoveFriend(friendshipID string) error {
	ctx := context.Background()
	request := &fr.RemoveFriendshipInput{
		FriendshipId: friendshipID,
	}

	// Make gRPC call
	_, err := s.transport.Send(ctx, s.client, "RemoveFriendshipRequest", request)
	if err != nil {
		log.Printf("Error in RemoveFriendshipRequest gRPC call: %v", err)
		return err
	}

	return nil
}

func (s *FriendshipService) FriendshipList(filter *fr.FriendshipFilter) (*fr.PaginatedFriendshipList, error) {
	ctx := context.Background()

	// Sending the request to list friendships based on filter
	response, err := s.transport.Send(ctx, s.client, "FriendshipRequestList", filter)
	if err != nil {
		log.Printf("Error sending request: %v", err)
		return nil, err
	}

	// Check if the response is of the expected type
	friendshipList, ok := response.(*fr.PaginatedFriendshipList)
	if !ok {
		return nil, errors.New("invalid response type from gRPC transport")
	}

	return friendshipList, nil
}
