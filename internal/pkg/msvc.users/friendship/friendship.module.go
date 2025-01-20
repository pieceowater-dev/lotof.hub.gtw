package friendship

import (
	"app/internal/pkg/msvc.users/friendship/ctrl"
	"app/internal/pkg/msvc.users/friendship/svc"
)

type Module struct {
	name    string
	version string
	API     *ctrl.FriendshipController
}

func New() Module {
	service := svc.NewFriendshipService()
	controller := ctrl.NewFriendshipController(service)
	return Module{
		name:    "FriendshipsMod",
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
