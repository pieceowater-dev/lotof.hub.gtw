package ctrl

import (
	"app/internal/pkg/namespaces/member/svc"
)

type MemberController struct {
	memberService *svc.MemberService
}

func NewMemberController(service *svc.MemberService) *MemberController {
	return &MemberController{memberService: service}
}
