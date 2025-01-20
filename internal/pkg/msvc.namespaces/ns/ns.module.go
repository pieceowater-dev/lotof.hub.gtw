package ns

import (
	"app/internal/pkg/msvc.namespaces/ns/ctrl"
	"app/internal/pkg/msvc.namespaces/ns/svc"
)

type Module struct {
	name    string
	version string
	API     *ctrl.NSController
}

func New() Module {
	service := svc.NewNSService()
	controller := ctrl.NewNSController(service)
	return Module{
		name:    "NamespacesMod",
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
