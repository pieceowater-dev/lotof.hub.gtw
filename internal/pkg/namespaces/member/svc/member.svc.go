package svc

import (
	"app/internal/core/cfg"
	member "app/internal/core/grpc/generated"
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
