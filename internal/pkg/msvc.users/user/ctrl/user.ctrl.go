package ctrl

import (
	"app/internal/core/generic/middleware"
	"app/internal/core/generic/utils"
	"app/internal/core/graph/model"
	pbUtils "app/internal/core/grpc/generated/generic/utils"
	user "app/internal/core/grpc/generated/lotof.hub.msvc.users/user"
	"app/internal/pkg/msvc.users/user/svc"
	"context"
	"errors"
	"log"
)

type UserController struct {
	userService *svc.UserService
}

func NewUserController(service *svc.UserService) *UserController {
	return &UserController{userService: service}
}

// Users retrieves a list of users and converts raw data to a paginated model.
func (c *UserController) Users(ctx context.Context, filter *model.DefaultFilterInput) (*model.PaginatedUserList, error) {
	request := &user.GetUsersRequest{
		Search:     "",
		Pagination: &pbUtils.Pagination{},
		Sort:       &pbUtils.Sort{},
	}

	if filter.Search != nil {
		request.Search = *filter.Search
	}
	if filter.Pagination != nil {
		request.Pagination = &pbUtils.Pagination{
			Page:   int32(*filter.Pagination.Page),
			Length: utils.PaginationLengthToInt(*filter.Pagination.Length),
		}
	}
	if filter.Sort != nil {
		request.Sort = &pbUtils.Sort{
			Field:     *filter.Sort.Field,
			Direction: utils.SortByEnumToString(filter.Sort.By),
		}
	}

	response, err := c.userService.FindAllUsers(request)
	if err != nil {
		log.Printf("Error fetching users: %v", err)
		return nil, err
	}

	var result []*model.User
	for _, u := range response.Users {
		result = append(result, &model.User{
			ID:       u.Id,
			Username: u.Username,
			Email:    u.Email,
		})
	}

	return &model.PaginatedUserList{
		Rows: result,
		Info: &model.PaginationInfo{
			Count: int(response.PaginationInfo.Count),
		},
	}, nil
}

// CreateUser handles creating a new user.
func (c *UserController) CreateUser(ctx context.Context, input model.UserInput) (*model.User, error) {
	request := &user.CreateUserRequest{
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password,
	}

	createUser, err := c.userService.CreateUser(request)
	if err != nil {
		log.Printf("Error creating createUser: %v", err)
		return nil, err
	}

	return &model.User{
		ID:       createUser.Id,
		Username: createUser.Username,
		Email:    createUser.Email,
	}, nil
}

// FindOneUser retrieves a single user by ID.
func (c *UserController) FindOneUser(ctx context.Context, id string) (*model.User, error) {
	oneUser, err := c.userService.FindOneUser(id)
	if err != nil {
		log.Printf("Error fetching oneUser: %v", err)
		return nil, err
	}

	return &model.User{
		ID:       oneUser.Id,
		Username: oneUser.Username,
		Email:    oneUser.Email,
	}, nil
}

// UpdateUser updates a user by ID.
func (c *UserController) UpdateUser(ctx context.Context, id string, input *model.UserInput) (*model.User, error) {
	request := &user.UpdateUserRequest{
		Id:       id,
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password,
	}

	updateUser, err := c.userService.UpdateUser(request)
	if err != nil {
		log.Printf("Error updating updateUser: %v", err)
		return nil, err
	}

	return &model.User{
		ID:       updateUser.Id,
		Username: updateUser.Username,
		Email:    updateUser.Email,
	}, nil
}

// Me retrieves the current user by decrypted token.
func (c *UserController) Me(ctx context.Context) (*model.User, error) {
	me, ok := ctx.Value(middleware.TokenContextKey).(*model.User)
	if !ok {
		return nil, errors.New("unable to fetch current user")
	}

	return me, nil
}
