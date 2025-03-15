package ctrl

import (
	"app/internal/core/generic/middleware"
	"app/internal/core/generic/utils"
	"app/internal/core/graph/model"
	fr "app/internal/core/grpc/generated/lotof.hub.msvc.users/friendship"
	usr "app/internal/core/grpc/generated/lotof.hub.msvc.users/user"
	"app/internal/pkg/msvc.users/friendship/svc"
	"context"
	"errors"

	//fr "app/internal/core/grpc/generated/lotof.hub.msvc.users/friendship"
	//"app/internal/pkg/msvc.users/friendship/svc"
	"log"
)

type FriendshipController struct {
	friendshipService *svc.FriendshipService
}

func NewFriendshipController(service *svc.FriendshipService) *FriendshipController {
	return &FriendshipController{friendshipService: service}
}

func (c *FriendshipController) CreateFriendship(ctx context.Context, input model.CreateFriendshipInput) (*model.Friendship, error) {
	request := &fr.CreateFriendshipInput{
		UserId: input.UserID,
		//FriendId: input.FriendID,
		Friend: &usr.User{
			Id:       "",
			Username: "",
			Email:    "",
		},
	}

	createFriendship, err := c.friendshipService.CreateFriendship(request)
	if err != nil {
		log.Printf("Error creating createFriendship: %v", err)
		return nil, err
	}

	return &model.Friendship{
		ID:     createFriendship.Id,
		UserID: createFriendship.User.Id,
		Friend: &model.User{
			ID:       createFriendship.Friend.Id,
			Username: createFriendship.Friend.Username,
			Email:    createFriendship.Friend.Email,
		},
		Status: utils.IntToFriendshipStatus(int(createFriendship.Status)),
	}, nil
}

func (c *FriendshipController) AcceptFriendshipRequest(ctx context.Context, input model.AcceptFriendshipInput) (*model.Friendship, error) {
	request := &fr.AcceptFriendshipInput{
		FriendshipId: input.FriendshipID,
	}

	acceptFriendship, err := c.friendshipService.AcceptFriendshipRequest(request)
	if err != nil {
		return nil, err
	}

	return &model.Friendship{
		ID:     acceptFriendship.Id,
		UserID: acceptFriendship.User.Id,
		Friend: &model.User{
			ID:       acceptFriendship.Friend.Id,
			Username: acceptFriendship.Friend.Username,
			Email:    acceptFriendship.Friend.Email,
		},
		Status: utils.IntToFriendshipStatus(int(acceptFriendship.Status)),
	}, nil
}

func (c *FriendshipController) RemoveFriendship(ctx context.Context, input model.RemoveFriendshipInput) (bool, error) {
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

func (c *FriendshipController) FriendshipList(ctx context.Context, filter *model.FriendshipFilter) (*model.PaginatedFriendshipList, error) {
	// Default to FriendshipStatusAccepted if Status is nil
	if filter.Status == nil {
		// Create a variable and take the address of it
		status := model.FriendshipStatusAccepted
		filter.Status = &status
	}

	// Mapping the filter to gRPC model for pagination, search, and sorting
	request := &fr.FriendshipFilter{
		UserId: filter.UserID,
		Status: uint32(utils.FriendshipStatusToInt(*filter.Status)),
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
			ID:     f.Id,
			UserID: f.User.Id,
			Friend: &model.User{
				ID:       f.Friend.Id,
				Username: f.Friend.Username,
				Email:    f.Friend.Email,
			},
			Status: utils.IntToFriendshipStatus(int(f.Status)),
		})
	}

	return &model.PaginatedFriendshipList{
		Rows: result,
		Info: &model.PaginationInfo{
			Count: int(friendshipList.PaginationInfo.Count),
		},
	}, nil
}

func (c *FriendshipController) MyFriends(ctx context.Context, status *model.FriendshipStatus) (*model.PaginatedFriendshipList, error) {
	// Default to FriendshipStatusAccepted if Status is nil
	if status == nil {
		// Create a variable and take the address of it
		s := model.FriendshipStatusAccepted
		status = &s
	}

	currUser, ok := ctx.Value(middleware.TokenContextKey).(*model.User)
	if !ok {
		return nil, errors.New("user not found in context")
	}

	// Mapping the filter to gRPC model for pagination, search, and sorting
	request := &fr.FriendshipFilter{
		UserId: currUser.ID,
		Status: uint32(utils.FriendshipStatusToInt(*status)),
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
			ID:     f.Id,
			UserID: f.User.Id,
			Friend: &model.User{
				ID:       f.Friend.Id,
				Username: f.Friend.Username,
				Email:    f.Friend.Email,
			},
			Status: utils.IntToFriendshipStatus(int(f.Status)),
		})
	}

	return &model.PaginatedFriendshipList{
		Rows: result,
		Info: &model.PaginationInfo{
			Count: int(friendshipList.PaginationInfo.Count),
		},
	}, nil
}
