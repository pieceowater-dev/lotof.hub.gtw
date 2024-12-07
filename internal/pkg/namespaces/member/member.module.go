package member

import (
	"app/internal/pkg/namespaces/member/ctrl"
	"app/internal/pkg/namespaces/member/svc"
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
