package ctrl

import (
	"app/internal/core/graph/model"
	"app/internal/pkg/msvc.namespaces/member/svc"
	"context"
)

type MemberController struct {
	memberService *svc.MemberService
}

func NewMemberController(service *svc.MemberService) *MemberController {
	return &MemberController{memberService: service}
}

func (c *MemberController) Members(
	_ context.Context,
	filter *model.MembersFilter,
) ([]*model.Member, error) {
	return []*model.Member{}, nil
}

func (c *MemberController) Member(
	_ context.Context,
	membershipID string,
) (*model.Member, error) {
	return &model.Member{}, nil
}

func (c *MemberController) AddMemberToNamespace(
	_ context.Context,
	input model.MemberToNamespaceInput,
) (*model.Namespace, error) {
	return &model.Namespace{}, nil
}

func (c *MemberController) RemoveMemberFromNamespace(
	_ context.Context,
	input model.MemberToNamespaceInput,
) (*model.Namespace, error) {
	return &model.Namespace{}, nil
}
