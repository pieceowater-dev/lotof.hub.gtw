package ctrl

import (
	"app/internal/core/generic"
	"app/internal/core/graph/model"
	fr "app/internal/core/grpc/generated"
	"app/internal/pkg/users/friendship/svc"
	"log"
)

type FriendshipController struct {
	friendshipService *svc.FriendshipService
}

func NewFriendshipController(service *svc.FriendshipService) *FriendshipController {
	return &FriendshipController{friendshipService: service}
}

func (c *FriendshipController) CreateFriendship(input model.CreateFriendshipInput) (*model.Friendship, error) {
	request := &fr.CreateFriendshipInput{
		UserId:   input.UserID,
		FriendId: input.FriendID,
	}

	createFriendship, err := c.friendshipService.CreateFriendship(request)
	if err != nil {
		log.Printf("Error creating createFriendship: %v", err)
		return nil, err
	}

	return &model.Friendship{
		ID:       createFriendship.Id,
		UserID:   createFriendship.User.Id,
		FriendID: createFriendship.Friend.Id,
		Status:   generic.IntToFriendshipStatus(int(createFriendship.Status)),
	}, nil
}

func (c *FriendshipController) AcceptFriendshipRequest(input model.AcceptFriendshipInput) (*model.Friendship, error) {
	//return c.friendshipService.AcceptFriendshipRequest(input)
	return nil, nil
}

func (c *FriendshipController) RemoveFriendship(input model.RemoveFriendshipInput) (bool, error) {
	//return c.friendshipService.RemoveFriendship(input)
	return true, nil
}

func (c *FriendshipController) FriendshipList(filter *model.FriendshipFilter) (*model.PaginatedFriendshipList, error) {
	//return c.friendshipService.FriendshipList(filter)
	return nil, nil
}
