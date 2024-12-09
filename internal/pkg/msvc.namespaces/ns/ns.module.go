package ns

import (
	"app/internal/pkg/msvc.namespaces/ns/ctrl"
	"app/internal/pkg/msvc.namespaces/ns/svc"
)

type Module struct {
	API *ctrl.NSController
}

func NewNSModule() Module {
	service := svc.NewNSService()
	controller := ctrl.NewNSController(service)
	return Module{
		API: controller,
	}
}
