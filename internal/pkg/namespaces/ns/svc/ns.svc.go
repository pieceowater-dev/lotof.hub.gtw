package svc

import (
	"app/internal/core/cfg"
	//ns "app/internal/core/grpc/generated"
	gossiper "github.com/pieceowater-dev/lotof.lib.gossiper/v2"
)

type NSService struct {
	transport gossiper.Transport
	//client    ns.NamespaceServiceClient // gRPC client
}

func NewNSService() *NSService {
	factory := gossiper.NewTransportFactory()
	grpcTransport := factory.CreateTransport(
		gossiper.GRPC,
		cfg.Inst().LotofHubMSvcNamespacesGrpcAddress,
	)

	// Create the client once and store it as a property
	//clientConstructor := ns.NewNamespaceServiceClient
	//client, err := grpcTransport.CreateClient(clientConstructor)
	//if err != nil {
	//	log.Fatalf("Error creating client: %v", err)
	//}

	return &NSService{
		transport: grpcTransport,
		//client:    client.(ns.NamespaceServiceClient), // Cast to the correct type
	}
}
