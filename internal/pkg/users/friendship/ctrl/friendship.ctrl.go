package ctrl

import (
	"app/internal/core/graph/model"
	"app/internal/pkg/users/friendship/svc"
)

type FriendshipController struct {
	friendshipService *svc.FriendshipService
}

func NewFriendshipController(service *svc.FriendshipService) *FriendshipController {
	return &FriendshipController{friendshipService: service}
}

func (c *FriendshipController) CreateFriendship(input model.CreateFriendshipInput) (*model.Friendship, error) {
	//return c.friendshipService.CreateFriendship(ctx, input)
	return nil, nil
}

func (c *FriendshipController) AcceptFriendshipRequest(input model.AcceptFriendshipInput) (*model.Friendship, error) {
	//return c.friendshipService.AcceptFriendshipRequest(ctx, input)
	return nil, nil
}

func (c *FriendshipController) RemoveFriendshipRequest(input model.RemoveFriendshipInput) (bool, error) {
	//return c.friendshipService.RemoveFriendshipRequest(ctx, input)
	return true, nil
}

func (c *FriendshipController) FriendshipRequestList(filter *model.FriendshipFilter) (*model.PaginatedFriendshipList, error) {
	//return c.friendshipService.FriendshipRequestList(ctx, filter)
	return nil, nil
}
