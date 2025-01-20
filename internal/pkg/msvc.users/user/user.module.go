package user

import (
	"app/internal/pkg/msvc.users/user/ctrl"
	"app/internal/pkg/msvc.users/user/svc"
)

type Module struct {
	name    string
	version string
	API     *ctrl.UserController
}

func New() Module {
	service := svc.NewUserService()
	controller := ctrl.NewUserController(service)
	return Module{
		name:    "UsersMod",
		version: "v1",
		API:     controller,
	}
}

// Initialize initializes the module. Currently not implemented.
func (m Module) Initialize() error {
	panic("Not implemented")
}

// Version returns the version of the module.
func (m Module) Version() string {
	return m.version
}

// Name returns the name of the module.
func (m Module) Name() string {
	return m.name
}
