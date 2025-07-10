package user

import (
	"github.com/VictorNaniii/clean-arhitecture/internal/service"
	desc "github.com/VictorNaniii/clean-arhitecture/pkg/user_v1"
)

type Implementation struct {
	desc.UnimplementedUserV1Server
	userService service.UserService
}
