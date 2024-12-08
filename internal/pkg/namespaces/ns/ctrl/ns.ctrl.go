package ctrl

import (
	"app/internal/core/generic"
	"app/internal/core/graph/model"
	ns "app/internal/core/grpc/generated"
	"app/internal/pkg/namespaces/ns/svc"
	"context"
	"log"
)

type NSController struct {
	nsService *svc.NSService
}

func NewNSController(service *svc.NSService) *NSController {
	return &NSController{nsService: service}
}

func (c *NSController) CreateNamespace(
	_ context.Context,
	input model.NamespaceInput,
) (*model.Namespace, error) {
	request := &ns.NamespaceInput{
		Title:       input.Title,
		Slug:        input.Slug,
		Description: *input.Description,
	}
	res, err := c.nsService.CreateNamespace(request)
	if err != nil {
		return nil, err
	}

	return &model.Namespace{
		ID:          res.Namespace.Id,
		Title:       res.Namespace.Title,
		Slug:        res.Namespace.Slug,
		Description: &res.Namespace.Description,
		Owner:       res.Namespace.Owner,
	}, nil
}

func (c *NSController) UpdateNamespace(
	_ context.Context,
	id string,
	input model.NamespaceInput,
) (*model.Namespace, error) {
	request := &ns.NamespaceMutationRequest{
		Id: id,
		Input: &ns.NamespaceInput{
			Title:       input.Title,
			Slug:        input.Slug,
			Description: *input.Description,
		},
	}

	res, err := c.nsService.UpdateNamespace(request)
	if err != nil {
		return nil, err
	}

	return &model.Namespace{
		ID:          res.Namespace.Id,
		Title:       res.Namespace.Title,
		Slug:        res.Namespace.Slug,
		Description: &res.Namespace.Description,
		Owner:       res.Namespace.Owner,
	}, nil
}

func (c *NSController) DeleteNamespace(
	_ context.Context,
	id string,
) (bool, error) {
	request := &ns.DeleteNamespaceRequest{
		Id: id,
	}

	res, err := c.nsService.DeleteNamespace(request)
	if err != nil {
		return false, err
	}

	return res.Success, nil
}

func (c *NSController) Namespaces(
	_ context.Context,
	filter *model.DefaultFilterInput,
) (*model.PaginatedNamespaceList, error) {
	request := &ns.NamespacesRequest{
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
	request := &ns.NamespaceRequest{
		Id: id,
	}
	res, err := c.nsService.GetNamespace(request)
	if err != nil {
		return nil, err
	}

	return &model.Namespace{
		ID:          res.Namespace.Id,
		Title:       res.Namespace.Title,
		Slug:        res.Namespace.Slug,
		Description: &res.Namespace.Description,
		Owner:       res.Namespace.Owner,
	}, nil
}
