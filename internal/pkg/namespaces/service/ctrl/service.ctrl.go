package ctrl

import (
	"app/internal/pkg/namespaces/service/svc"
)

type ServiceController struct {
	serviceService *svc.ServiceService
}

func NewServiceController(service *svc.ServiceService) *ServiceController {
	return &ServiceController{serviceService: service}
}
