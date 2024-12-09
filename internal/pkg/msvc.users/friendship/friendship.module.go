package friendship

import (
	"app/internal/pkg/msvc.users/friendship/ctrl"
	"app/internal/pkg/msvc.users/friendship/svc"
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
