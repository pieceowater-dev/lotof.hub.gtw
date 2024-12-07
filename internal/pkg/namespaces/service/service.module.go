package service

import (
	"app/internal/pkg/namespaces/service/ctrl"
	"app/internal/pkg/namespaces/service/svc"
)

type Module struct {
	API *ctrl.ServiceController
}

func NewServiceModule() Module {
	service := svc.NewServiceService()
	controller := ctrl.NewServiceController(service)
	return Module{
		API: controller,
	}
}
