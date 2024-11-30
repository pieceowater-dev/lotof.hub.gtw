package ctrl

import (
	"app/internal/core/graph/model"
	"app/internal/pkg/users/auth/svc"
)

type AuthController struct {
	authService *svc.AuthService
}

func NewAuthController(service *svc.AuthService) *AuthController {
	return &AuthController{authService: service}
}

func (c *AuthController) Login(input model.LoginRequest) (*model.AuthResponse, error) {
	//return c.authService.Login(ctx, input)
	return nil, nil
}

func (c *AuthController) Register(input model.RegisterRequest) (*model.AuthResponse, error) {
	//return c.authService.Register(ctx, input)
	return nil, nil
}
