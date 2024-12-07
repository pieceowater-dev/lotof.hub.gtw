package ctrl

import (
	"app/internal/pkg/namespaces/ns/svc"
)

type NSController struct {
	nsService *svc.NSService
}

func NewNSController(service *svc.NSService) *NSController {
	return &NSController{nsService: service}
}
