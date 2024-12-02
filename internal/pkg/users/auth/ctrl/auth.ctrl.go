package ctrl

import (
	"app/internal/core/graph/model"
	auth "app/internal/core/grpc/generated"
	"app/internal/pkg/users/auth/svc"
	"log"
)

type AuthController struct {
	authService *svc.AuthService
}

func NewAuthController(service *svc.AuthService) *AuthController {
	return &AuthController{authService: service}
}

func (c *AuthController) Login(input model.LoginRequest) (*model.AuthResponse, error) {
	request := &auth.LoginRequest{
		Email:    input.Email,
		Password: input.Password,
	}

	login, err := c.authService.Login(request)
	if err != nil {
		log.Printf("Error login: %v", err)
		return nil, err
	}

	return &model.AuthResponse{
		Token: login.Token,
		User: &model.User{
			ID:       login.User.Id,
			Username: login.User.Username,
			Email:    login.User.Email,
		},
	}, nil
}

func (c *AuthController) Register(input model.RegisterRequest) (*model.AuthResponse, error) {
	request := &auth.RegisterRequest{
		Email:    input.Email,
		Username: input.Username,
		Password: input.Password,
	}

	register, err := c.authService.Register(request)
	if err != nil {
		log.Printf("Error login: %v", err)
		return nil, err
	}

	return &model.AuthResponse{
		Token: register.Token,
		User: &model.User{
			ID:       register.User.Id,
			Username: register.User.Username,
			Email:    register.User.Email,
		},
	}, nil
}
