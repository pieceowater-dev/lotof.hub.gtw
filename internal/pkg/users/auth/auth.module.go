package auth

import (
	"app/internal/pkg/users/auth/ctrl"
	"app/internal/pkg/users/auth/svc"
)

type Module struct {
	API *ctrl.AuthController
}

func NewAuthModule() Module {
	service := svc.NewAuthService()
	controller := ctrl.NewAuthController(service)
	return Module{
		API: controller,
	}
}
