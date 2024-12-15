package ctrl

import (
	"app/internal/core/graph/model"
	auth "app/internal/core/grpc/generated"
	ns "app/internal/pkg/msvc.namespaces/ns/svc"
	"app/internal/pkg/msvc.users/auth/svc"
	"context"
	"log"
)

type AuthController struct {
	authService *svc.AuthService
	nsService   *ns.NSService
}

func NewAuthController(service *svc.AuthService, namespaceService *ns.NSService) *AuthController {
	return &AuthController{authService: service, nsService: namespaceService}
}

func (c *AuthController) VerifyToken(t string) (bool, *model.User, error) {
	data, err := c.authService.VerifyToken(t)
	if err != nil {
		return false, nil, err
	}

	if data.Message != "" {
		log.Printf("Token message: %v", data.Message)
	}

	log.Printf("Current user: %+v", data.User.Username)

	return data.Valid, &model.User{
		ID:       data.User.Id,
		Username: data.User.Username,
		Email:    data.User.Email,
	}, nil
}

func (c *AuthController) Login(ctx context.Context, input model.LoginRequest) (*model.AuthResponse, error) {
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

func (c *AuthController) Register(ctx context.Context, input model.RegisterRequest) (*model.AuthResponse, error) {
	request := &auth.RegisterRequest{
		Email:    input.Email,
		Username: input.Username,
		Password: input.Password,
	}

	register, err := c.authService.Register(request)
	if err != nil {
		log.Printf("Error: %v", err)
		return nil, err
	}

	_, err = c.nsService.CreateNamespace(&auth.NamespaceRequest{
		Title:       register.User.Username,
		Slug:        register.User.Username,
		Description: "",
		//pass ns owner id also
	})
	if err != nil {
		log.Printf("Error: %v", err)
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
