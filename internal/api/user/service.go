package user

import (
	desc "../../../pkg/user_v1"
	"../../service"
)

type Implementation struct {
	desc.UnimplementedUserV1Server
	userService service.UserService
}
