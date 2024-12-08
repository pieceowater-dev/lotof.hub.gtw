package svc

import (
	"app/internal/core/cfg"
	ns "app/internal/core/grpc/generated"
	"context"
	"errors"
	gossiper "github.com/pieceowater-dev/lotof.lib.gossiper/v2"
	"log"
)

type NSService struct {
	transport gossiper.Transport
	client    ns.NamespaceServiceClient
}

func NewNSService() *NSService {
	factory := gossiper.NewTransportFactory()
	grpcTransport := factory.CreateTransport(
		gossiper.GRPC,
		cfg.Inst().LotofHubMSvcNamespacesGrpcAddress,
	)

	// Create the client once and store it as a property
	clientConstructor := ns.NewNamespaceServiceClient
	client, err := grpcTransport.CreateClient(clientConstructor)
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}

	return &NSService{
		transport: grpcTransport,
		client:    client.(ns.NamespaceServiceClient),
	}
}

func (s *NSService) GetNamespaces(input *ns.NamespacesRequest) (*ns.NamespacesResponse, error) {
	ctx := context.Background()
	response, err := s.transport.Send(ctx, s.client, "GetNamespaces", input)
	if err != nil {
		log.Printf("Error sending request: %v", err)
		return nil, err
	}
	res, ok := response.(*ns.NamespacesResponse)
	if !ok {
		return nil, errors.New("invalid response type from gRPC transport")
	}
	return res, nil
}

func (s *NSService) GetNamespace(input *ns.NamespaceRequest) (*ns.NamespaceResponse, error) {
	ctx := context.Background()
	response, err := s.transport.Send(ctx, s.client, "GetNamespace", input)
	if err != nil {
		log.Printf("Error sending request: %v", err)
		return nil, err
	}
	res, ok := response.(*ns.NamespaceResponse)
	if !ok {
		return nil, errors.New("invalid response type from gRPC transport")
	}
	return res, nil
}

func (s *NSService) CreateNamespace(input *ns.NamespaceInput) (*ns.NamespaceMutationResponse, error) {
	ctx := context.Background()

	response, err := s.transport.Send(ctx, s.client, "CreateNamespace", input)
	if err != nil {
		log.Printf("Error sending request: %v", err)
		return nil, err
	}

	res, ok := response.(*ns.NamespaceMutationResponse)
	if !ok {
		return nil, errors.New("invalid response type from gRPC transport")
	}

	return res, nil
}

func (s *NSService) UpdateNamespace(input *ns.NamespaceMutationRequest) (*ns.NamespaceMutationResponse, error) {
	ctx := context.Background()
	response, err := s.transport.Send(ctx, s.client, "UpdateNamespace", input)
	if err != nil {
		log.Printf("Error sending request: %v", err)
		return nil, err
	}
	res, ok := response.(*ns.NamespaceMutationResponse)
	if !ok {
		return nil, errors.New("invalid response type from gRPC transport")
	}
	return res, nil
}

func (s *NSService) DeleteNamespace(input *ns.DeleteNamespaceRequest) (*ns.DeleteNamespaceResponse, error) {
	ctx := context.Background()
	response, err := s.transport.Send(ctx, s.client, "DeleteNamespace", input)
	if err != nil {
		log.Printf("Error sending request: %v", err)
		return nil, err
	}
	res, ok := response.(*ns.DeleteNamespaceResponse)
	if !ok {
		return nil, errors.New("invalid response type from gRPC transport")
	}
	return res, nil
}
