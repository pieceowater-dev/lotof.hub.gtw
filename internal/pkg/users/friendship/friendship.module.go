package friendship

import (
	"app/internal/pkg/users/friendship/ctrl"
	"app/internal/pkg/users/friendship/svc"
)

type Module struct {
	API *ctrl.FriendshipController
}

func NewFriendshipModule() Module {
	service := svc.NewFriendshipService()
	controller := ctrl.NewFriendshipController(service)
	return Module{
		API: controller,
	}
}
