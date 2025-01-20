package auth

import (
	ns "app/internal/pkg/msvc.namespaces/ns/svc"
	"app/internal/pkg/msvc.users/auth/ctrl"
	"app/internal/pkg/msvc.users/auth/svc"
)

type Module struct {
	name    string
	version string
	API     *ctrl.AuthController
}

func New() Module {
	service := svc.NewAuthService()
	namespaceService := ns.NewNSService()
	controller := ctrl.NewAuthController(service, namespaceService)
	return Module{
		name:    "AuthMod",
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
