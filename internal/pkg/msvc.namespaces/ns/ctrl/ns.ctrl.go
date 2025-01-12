package ctrl

import (
	"app/internal/core/generic"
	"app/internal/core/graph/model"
	ns "app/internal/core/grpc/generated"
	pb "app/internal/core/grpc/generated"
	"app/internal/pkg/msvc.namespaces/ns/svc"
	"context"
	"log"
)

type NSController struct {
	nsService *svc.NSService
	pb.UnimplementedGatewayServiceServer
}

func NewNSController(service *svc.NSService) *NSController {
	return &NSController{nsService: service}
}

func (c *NSController) CreateNamespace(
	_ context.Context,
	input model.NamespaceInput,
) (*model.Namespace, error) {
	nsDesc := ""
	if input.Description != nil {
		nsDesc = *input.Description
	}
	request := &ns.NamespaceRequest{
		Title:       input.Title,
		Slug:        input.Slug,
		Description: nsDesc,
	}
	res, err := c.nsService.CreateNamespace(request)
	if err != nil {
		return nil, err
	}

	return &model.Namespace{
		ID:          res.Id,
		Title:       res.Title,
		Slug:        res.Slug,
		Description: &res.Description,
		Owner:       res.Owner,
	}, err
}

func (c *NSController) UpdateNamespace(
	_ context.Context,
	input model.NamespaceInput,
) (*model.Namespace, error) {
	request := &ns.UpdateNamespaceRequest{
		Title:       input.Title,
		Slug:        input.Slug,
		Description: *input.Description,
	}

	res, err := c.nsService.UpdateNamespace(request)
	if err != nil {
		return nil, err
	}

	return &model.Namespace{
		ID:          res.Id,
		Title:       res.Title,
		Slug:        res.Slug,
		Description: &res.Description,
		Owner:       res.Owner,
	}, nil
}

func (c *NSController) Namespaces(
	_ context.Context,
	filter *model.DefaultFilterInput,
) (*model.PaginatedNamespaceList, error) {
	request := &ns.GetNamespacesRequest{
		Search:     "",
		Pagination: &ns.Pagination{},
		Sort:       &ns.Sort{},
	}

	if filter.Search != nil {
		request.Search = *filter.Search
	}
	if filter.Pagination != nil {
		request.Pagination = &ns.Pagination{
			Page:   int32(*filter.Pagination.Page),
			Length: generic.PaginationLengthToInt(*filter.Pagination.Length),
		}
	}
	if filter.Sort != nil {
		request.Sort = &ns.Sort{
			Field:     *filter.Sort.Field,
			Direction: generic.SortByEnumToString(filter.Sort.By),
		}
	}

	response, err := c.nsService.GetNamespaces(request)
	if err != nil {
		log.Printf("Error fetching: %v", err)
		return nil, err
	}

	var result []*model.Namespace
	for _, n := range response.Namespaces.Rows {
		result = append(result, &model.Namespace{
			ID:          n.Id,
			Title:       n.Title,
			Slug:        n.Slug,
			Description: &n.Description,
			Owner:       n.Owner,
		})
	}

	return &model.PaginatedNamespaceList{
		Rows: result,
		Info: &model.PaginationInfo{
			Count: int(response.Namespaces.Info.Count),
		},
	}, nil
}

func (c *NSController) Namespace(
	_ context.Context,
	id string,
) (*model.Namespace, error) {
	request := &ns.GetNamespaceRequest{
		Id: id,
	}
	res, err := c.nsService.GetNamespace(request)
	if err != nil {
		return nil, err
	}

	return &model.Namespace{
		ID:          res.Id,
		Title:       res.Title,
		Slug:        res.Slug,
		Description: &res.Description,
		Owner:       res.Owner,
	}, nil
}

func (c *NSController) AddAppToNamespace(
	_ context.Context,
	namespaceID string,
	appBundle string,
) (*model.NamespaceApp, error) {
	request := &ns.AddAppToNamespaceRequest{
		NamespaceId: namespaceID,
		AppBundle:   appBundle,
	}
	res, err := c.nsService.AddAppToNamespace(request)
	if err != nil {
		return nil, err
	}

	return &model.NamespaceApp{
		ID:          res.Id,
		NamespaceID: res.NamespaceId,
		AppBundle:   res.AppBundle,
	}, nil
}

func (c *NSController) GatewayNamespacesByApp(_ context.Context, request *pb.GatewayNamespacesByAppRequest) (*pb.GatewayNamespacesByAppResponse, error) {
	req := &ns.NamespacesByAppRequest{
		AppBundle: request.AppBundle,
	}
	res, err := c.nsService.NamespacesByApp(req)
	if err != nil {
		return nil, err
	}

	var tenants []*pb.GatewayEncryptedTenant
	for _, tenant := range res.Tenants {
		tenants = append(tenants, &pb.GatewayEncryptedTenant{
			Namespace:   tenant.Namespace,
			Credentials: tenant.Credentials,
		})
	}

	return &pb.GatewayNamespacesByAppResponse{
		Tenants: tenants,
	}, nil
}
