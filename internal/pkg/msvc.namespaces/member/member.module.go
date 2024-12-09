package member

import (
	"app/internal/pkg/msvc.namespaces/member/ctrl"
	"app/internal/pkg/msvc.namespaces/member/svc"
)

type Module struct {
	API *ctrl.MemberController
}

func NewMemberModule() Module {
	service := svc.NewMemberService()
	controller := ctrl.NewMemberController(service)
	return Module{
		API: controller,
	}
}
