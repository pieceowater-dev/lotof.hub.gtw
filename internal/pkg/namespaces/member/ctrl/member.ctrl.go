package ctrl

import (
	"app/internal/core/graph/model"
	"app/internal/pkg/namespaces/member/svc"
	"context"
)

type MemberController struct {
	memberService *svc.MemberService
}

func NewMemberController(service *svc.MemberService) *MemberController {
	return &MemberController{memberService: service}
}

func (c *MemberController) AddMemberToNamespace(
	_ context.Context,
	input model.AddMemberToNamespaceInput,
) (*model.Namespace, error) {
	return &model.Namespace{}, nil
}

func (c *MemberController) RemoveMemberFromNamespace(
	_ context.Context,
	input model.AddMemberToNamespaceInput,
) (*model.Namespace, error) {
	return &model.Namespace{}, nil
}

func (c *MemberController) AddMemberToService(
	_ context.Context,
	input model.AddMemberToServiceInput,
) (*model.Service, error) {
	return &model.Service{}, nil
}

func (c *MemberController) RemoveMemberFromService(
	_ context.Context,
	input model.AddMemberToServiceInput,
) (*model.Service, error) {
	return &model.Service{}, nil
}

func (c *MemberController) Members(
	_ context.Context,
	filter *model.DefaultFilterInput,
) ([]*model.Member, error) {
	return []*model.Member{}, nil
}

func (c *MemberController) Member(
	_ context.Context,
	membershipID string,
) (*model.Member, error) {
	return &model.Member{}, nil
}
