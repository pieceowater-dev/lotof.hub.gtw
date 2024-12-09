package ctrl

import (
	"app/internal/core/generic"
	"app/internal/core/graph/model"
	fr "app/internal/core/grpc/generated"
	"app/internal/pkg/msvc.users/friendship/svc"
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
	request := &fr.AcceptFriendshipInput{
		FriendshipId: input.FriendshipID,
	}

	acceptFriendship, err := c.friendshipService.AcceptFriendshipRequest(request)
	if err != nil {
		return nil, err
	}

	return &model.Friendship{
		ID:       acceptFriendship.Id,
		UserID:   acceptFriendship.User.Id,
		FriendID: acceptFriendship.Friend.Id,
		Status:   generic.IntToFriendshipStatus(int(acceptFriendship.Status)),
	}, nil
}

func (c *FriendshipController) RemoveFriendship(input model.RemoveFriendshipInput) (bool, error) {
	// Mapping the input to gRPC model
	request := &fr.RemoveFriendshipInput{
		FriendshipId: input.FriendshipID,
	}

	// Calling the service to remove the friendship
	err := c.friendshipService.RemoveFriend(request.FriendshipId)
	if err != nil {
		return false, err
	}

	// If successful, return true
	return true, nil
}

func (c *FriendshipController) FriendshipList(filter *model.FriendshipFilter) (*model.PaginatedFriendshipList, error) {
	// Mapping the filter to gRPC model for pagination, search, and sorting
	request := &fr.FriendshipFilter{
		UserId: filter.UserID,
		Status: uint32(generic.FriendshipStatusToInt(*filter.Status)),
	}

	// Calling the service to get the friendship list
	friendshipList, err := c.friendshipService.FriendshipList(request)
	if err != nil {
		log.Printf("Error fetching friendships: %v", err)
		return nil, err
	}

	// Mapping the response to GraphQL model
	var result []*model.Friendship
	for _, f := range friendshipList.Friendships {
		result = append(result, &model.Friendship{
			ID:       f.Id,
			UserID:   f.User.Id,
			FriendID: f.Friend.Id,
			Status:   generic.IntToFriendshipStatus(int(f.Status)),
		})
	}

	return &model.PaginatedFriendshipList{
		Rows: result,
		Info: &model.PaginationInfo{
			Count: int(friendshipList.PaginationInfo.Count),
		},
	}, nil
}
