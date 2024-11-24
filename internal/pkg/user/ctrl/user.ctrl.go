package ctrl

import (
	"app/internal/core/graph/model"
	user "app/internal/core/grpc/generated"
	"app/internal/pkg/user/svc"
	"log"
)

type UserController struct {
	userService *svc.UserService
}

func NewUserController(service *svc.UserService) *UserController {
	return &UserController{userService: service}
}

// Users retrieves a list of users and converts raw data to a paginated model.
func (c *UserController) Users() (*model.PaginatedUserList, error) {
	// Get raw response from the service
	response, err := c.userService.FindAllUsers(&user.GetUsersRequest{
		Pagination: &user.Pagination{
			Page:   1,
			Length: 10,
		},
	})
	if err != nil {
		log.Printf("Error fetching users: %v", err)
		return nil, err
	}

	// Convert the raw response to PaginatedUserList
	users := make([]*model.User, len(response.Users))
	for i, u := range response.Users {
		users[i] = &model.User{
			ID:       u.Id,
			Username: u.Username,
			Email:    u.Email,
			// Map other fields as needed
		}
	}

	userList := &model.PaginatedUserList{
		Rows: users,
		Info: &model.PaginationInfo{
			Count: len(response.Users), // Count from raw response
		},
	}

	return userList, nil
}

// CreateUser handles creating a new user.
func (c *UserController) CreateUser(input *model.CreateUserInput) (*model.User, error) {
	// Prepare raw request data
	request := &user.CreateUserRequest{
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password, // Ensure password is handled securely
	}

	// Call service to create the createUser and get raw response
	createUser, err := c.userService.CreateUser(request)
	if err != nil {
		log.Printf("Error creating createUser: %v", err)
		return nil, err
	}

	// Convert raw createUser data to model
	return &model.User{
		ID:       createUser.Id,
		Username: createUser.Username,
		Email:    createUser.Email,
		// Map other fields as necessary
	}, nil
}

// FindOneUser retrieves a single user by ID.
func (c *UserController) FindOneUser(id string) (*model.User, error) {
	// Get raw response from service
	oneUser, err := c.userService.FindOneUser(id)
	if err != nil {
		log.Printf("Error fetching oneUser: %v", err)
		return nil, err
	}

	// Convert raw oneUser data to model
	return &model.User{
		ID:       oneUser.Id,
		Username: oneUser.Username,
		Email:    oneUser.Email,
		// Map other fields as necessary
	}, nil
}

// UpdateUser updates a user by ID.
func (c *UserController) UpdateUser(id string, input *model.UpdateUserInput) (*model.User, error) {
	// Prepare raw request data
	request := &user.UpdateUserRequest{
		Id:       id,
		Username: *input.Username,
		Email:    *input.Email,
		Password: *input.Password,
	}

	// Call service to update the updateUser and get raw response
	updateUser, err := c.userService.UpdateUser(request)
	if err != nil {
		log.Printf("Error updating updateUser: %v", err)
		return nil, err
	}

	// Convert raw updateUser data to model
	return &model.User{
		ID:       updateUser.Id,
		Username: updateUser.Username,
		Email:    updateUser.Email,
		// Map other fields as necessary
	}, nil
}

// DeleteUser deletes a user by ID.
func (c *UserController) DeleteUser(id string) (bool, error) {
	// Get raw response from service
	response, err := c.userService.DeleteUser(id)
	if err != nil {
		log.Printf("Error deleting user: %v", err)
		return false, err
	}

	// Return the success flag from the raw response
	return response.Success, nil
}
