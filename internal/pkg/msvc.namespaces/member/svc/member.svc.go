package svc

import (
	"app/internal/core/cfg"
	member "app/internal/core/grpc/generated"
	"context"
	"errors"
	gossiper "github.com/pieceowater-dev/lotof.lib.gossiper/v2"
	"log"
)

type MemberService struct {
	transport gossiper.Transport
	client    member.MemberServiceClient // gRPC client
}

func NewMemberService() *MemberService {
	factory := gossiper.NewTransportFactory()
	grpcTransport := factory.CreateTransport(
		gossiper.GRPC,
		cfg.Inst().LotofHubMSvcNamespacesGrpcAddress,
	)

	// Create the client once and store it as a property
	clientConstructor := member.NewMemberServiceClient
	client, err := grpcTransport.CreateClient(clientConstructor)
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}

	return &MemberService{
		transport: grpcTransport,
		client:    client.(member.MemberServiceClient), // Cast to the correct type
	}
}

func (s *MemberService) GetMembers(input *member.GetMembersRequest) (*member.GetMembersResponse, error) {
	ctx := context.Background()

	response, err := s.transport.Send(ctx, s.client, "GetMembers", input)
	if err != nil {
		log.Printf("Error sending request: %v", err)
		return nil, err
	}
	res, ok := response.(*member.GetMembersResponse)
	if !ok {
		return nil, errors.New("invalid response type from gRPC transport")
	}
	return res, nil
}

func (s *MemberService) GetMember(input *member.GetMemberRequest) (*member.GetMemberResponse, error) {
	ctx := context.Background()

	response, err := s.transport.Send(ctx, s.client, "GetMember", input)
	if err != nil {
		log.Printf("Error sending request: %v", err)
		return nil, err
	}
	res, ok := response.(*member.GetMemberResponse)
	if !ok {
		return nil, errors.New("invalid response type from gRPC transport")
	}
	return res, nil
}

func (s *MemberService) AddMemberToNamespace(input *member.MemberToNamespaceRequest) (*member.OK, error) {
	ctx := context.Background()

	response, err := s.transport.Send(ctx, s.client, "AddMemberToNamespace", input)
	if err != nil {
		log.Printf("Error sending request: %v", err)
		return &member.OK{Ok: false}, err
	}
	res, ok := response.(*member.OK)
	if !ok {
		return &member.OK{Ok: false}, errors.New("invalid response type from gRPC transport")
	}
	return res, nil
}

func (s *MemberService) RemoveMemberFromNamespace(input *member.MemberToNamespaceRequest) (*member.OK, error) {
	ctx := context.Background()

	response, err := s.transport.Send(ctx, s.client, "RemoveMemberFromNamespace", input)
	if err != nil {
		log.Printf("Error sending request: %v", err)
		return &member.OK{Ok: false}, err
	}
	res, ok := response.(*member.OK)
	if !ok {
		return &member.OK{Ok: false}, errors.New("invalid response type from gRPC transport")
	}
	return res, nil
}
