package user

import (
	"app/internal/pkg/users/user/ctrl"
	"app/internal/pkg/users/user/svc"
)

type Module struct {
	API *ctrl.UserController
}

func NewUserModule() Module {
	service := svc.NewUserService()
	controller := ctrl.NewUserController(service)
	return Module{
		API: controller,
	}
}
