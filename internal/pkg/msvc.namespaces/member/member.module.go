package member

import (
	"app/internal/pkg/msvc.namespaces/member/ctrl"
	"app/internal/pkg/msvc.namespaces/member/svc"
)

type Module struct {
	name    string
	version string
	API     *ctrl.MemberController
}

func New() Module {
	service := svc.NewMemberService()
	controller := ctrl.NewMemberController(service)
	return Module{
		name:    "MembersMod",
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
