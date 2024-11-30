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

func (c *FriendshipController) RemoveFriendship(input model.RemoveFriendshipInput) (bool, error) {
	//return c.friendshipService.RemoveFriendship(ctx, input)
	return true, nil
}

func (c *FriendshipController) FriendshipList(filter *model.FriendshipFilter) (*model.PaginatedFriendshipList, error) {
	//return c.friendshipService.FriendshipList(ctx, filter)
	return nil, nil
}
