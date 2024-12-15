package auth

import (
	ns "app/internal/pkg/msvc.namespaces/ns/svc"
	"app/internal/pkg/msvc.users/auth/ctrl"
	"app/internal/pkg/msvc.users/auth/svc"
)

type Module struct {
	API *ctrl.AuthController
}

func NewAuthModule() Module {
	service := svc.NewAuthService()
	namespaceService := ns.NewNSService()
	controller := ctrl.NewAuthController(service, namespaceService)
	return Module{
		API: controller,
	}
}
