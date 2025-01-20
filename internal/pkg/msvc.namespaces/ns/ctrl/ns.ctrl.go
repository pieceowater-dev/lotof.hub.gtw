package ctrl

import (
	"app/internal/core/generic/utils"
	"app/internal/core/graph/model"
	pbu "app/internal/core/grpc/generated/generic/utils"
	"app/internal/core/grpc/generated/lotof.hub.gtw/gtw"
	"app/internal/core/grpc/generated/lotof.hub.msvc.namespaces/ns"
	"app/internal/pkg/msvc.namespaces/ns/svc"
	"context"
	"log"
)

type NSController struct {
	nsService *svc.NSService
	gtw.UnimplementedGatewayServiceServer
}

func NewNSController(service *svc.NSService) *NSController {
	return &NSController{nsService: service}
}

func (c *NSController) CreateNamespace(
	ctx context.Context,
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
	res, err := c.nsService.CreateNamespace(ctx, request)
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
	ctx context.Context,
	input model.NamespaceInput,
) (*model.Namespace, error) {
	request := &ns.UpdateNamespaceRequest{
		Title:       input.Title,
		Slug:        input.Slug,
		Description: *input.Description,
	}

	res, err := c.nsService.UpdateNamespace(ctx, request)
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
	ctx context.Context,
	filter *model.DefaultFilterInput,
) (*model.PaginatedNamespaceList, error) {
	request := &ns.GetNamespacesRequest{
		Search:     "",
		Pagination: &pbu.Pagination{},
		Sort:       &pbu.Sort{},
	}

	if filter.Search != nil {
		request.Search = *filter.Search
	}
	if filter.Pagination != nil {
		request.Pagination = &pbu.Pagination{
			Page:   int32(*filter.Pagination.Page),
			Length: utils.PaginationLengthToInt(*filter.Pagination.Length),
		}
	}
	if filter.Sort != nil {
		request.Sort = &pbu.Sort{
			Field:     *filter.Sort.Field,
			Direction: utils.SortByEnumToString(filter.Sort.By),
		}
	}

	response, err := c.nsService.GetNamespaces(ctx, request)
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
	ctx context.Context,
	id string,
) (*model.Namespace, error) {
	request := &ns.GetNamespaceRequest{
		Id: id,
	}
	res, err := c.nsService.GetNamespace(ctx, request)
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
	ctx context.Context,
	namespaceID string,
	appBundle string,
) (*model.NamespaceApp, error) {
	request := &ns.AddAppToNamespaceRequest{
		NamespaceId: namespaceID,
		AppBundle:   appBundle,
	}
	res, err := c.nsService.AddAppToNamespace(ctx, request)
	if err != nil {
		return nil, err
	}

	return &model.NamespaceApp{
		ID:          res.Id,
		NamespaceID: res.NamespaceId,
		AppBundle:   res.AppBundle,
	}, nil
}

func (c *NSController) GatewayNamespacesByApp(
	ctx context.Context,
	request *gtw.GatewayNamespacesByAppRequest,
) (*gtw.GatewayNamespacesByAppResponse, error) {
	req := &ns.NamespacesByAppRequest{
		AppBundle: request.AppBundle,
	}
	res, err := c.nsService.NamespacesByApp(ctx, req)
	if err != nil {
		return nil, err
	}

	var tenants []*gtw.GatewayEncryptedTenant
	for _, tenant := range res.Tenants {
		tenants = append(tenants, &gtw.GatewayEncryptedTenant{
			Namespace:   tenant.Namespace,
			Credentials: tenant.Credentials,
		})
	}

	return &gtw.GatewayNamespacesByAppResponse{
		Tenants: tenants,
	}, nil
}
