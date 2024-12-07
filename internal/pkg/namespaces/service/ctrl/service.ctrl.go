package ctrl

import (
	"app/internal/core/graph/model"
	"app/internal/pkg/namespaces/service/svc"
	"context"
)

type ServiceController struct {
	serviceService *svc.ServiceService
}

func NewServiceController(service *svc.ServiceService) *ServiceController {
	return &ServiceController{serviceService: service}
}

func (c *ServiceController) CreateService(
	_ context.Context,
	input model.ServiceInput,
) (*model.Service, error) {
	return &model.Service{}, nil
}

func (c *ServiceController) UpdateService(
	_ context.Context,
	id string,
	input model.ServiceInput,
) (*model.Service, error) {
	return &model.Service{}, nil
}

func (c *ServiceController) DeleteService(
	_ context.Context,
	id string,
) (bool, error) {
	return false, nil
}

func (c *ServiceController) Services(
	_ context.Context,
	filter *model.DefaultFilterInput,
) (*model.PaginatedServiceList, error) {
	return &model.PaginatedServiceList{}, nil
}

func (c *ServiceController) ServicesByNamespace(
	_ context.Context,
	nsID string,
	filter *model.DefaultFilterInput,
) (*model.PaginatedServiceList, error) {
	return &model.PaginatedServiceList{}, nil
}

func (c *ServiceController) Service(
	_ context.Context,
	id string,
) (*model.Service, error) {
	return &model.Service{}, nil
}
