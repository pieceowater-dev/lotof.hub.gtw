package svc

import (
	"app/internal/core/cfg"
	//service "app/internal/core/grpc/generated"
	gossiper "github.com/pieceowater-dev/lotof.lib.gossiper/v2"
)

type ServiceService struct {
	transport gossiper.Transport
	//client    service.ServiceServiceClient // gRPC client
}

func NewServiceService() *ServiceService {
	factory := gossiper.NewTransportFactory()
	grpcTransport := factory.CreateTransport(
		gossiper.GRPC,
		cfg.Inst().LotofHubMSvcNamespacesGrpcAddress,
	)

	// Create the client once and store it as a property
	//clientConstructor := service.NewServiceServiceClient
	//client, err := grpcTransport.CreateClient(clientConstructor)
	//if err != nil {
	//	log.Fatalf("Error creating client: %v", err)
	//}

	return &ServiceService{
		transport: grpcTransport,
		//client:    client.(service.ServiceServiceClient), // Cast to the correct type
	}
}
