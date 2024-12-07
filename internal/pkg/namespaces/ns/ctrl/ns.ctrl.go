package ctrl

import (
	"app/internal/core/graph/model"
	"app/internal/pkg/namespaces/ns/svc"
	"context"
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
	return &model.Namespace{
		ID:          "",
		Title:       "",
		Slug:        "",
		Description: nil,
		Owner:       "",
		Services:    nil,
		Members:     nil,
	}, nil
}

func (c *NSController) UpdateNamespace(
	_ context.Context,
	id string,
	input model.NamespaceInput,
) (*model.Namespace, error) {
	return &model.Namespace{
		ID:          "",
		Title:       "",
		Slug:        "",
		Description: nil,
		Owner:       "",
		Services:    nil,
		Members:     nil,
	}, nil
}

func (c *NSController) DeleteNamespace(
	_ context.Context,
	id string,
) (bool, error) {
	return false, nil
}

func (c *NSController) Namespaces(
	_ context.Context,
	filter *model.DefaultFilterInput,
) (*model.PaginatedNamespaceList, error) {
	return &model.PaginatedNamespaceList{
		Rows: []*model.Namespace{},
		Info: nil,
	}, nil
}

func (c *NSController) Namespace(
	_ context.Context,
	id string,
) (*model.Namespace, error) {
	return nil, nil
}
