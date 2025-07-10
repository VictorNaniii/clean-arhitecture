package user

import (
	"context"
	"github.com/VictorNaniii/clean-arhitecture/internal/model"
	desc "github.com/VictorNaniii/clean-arhitecture/pkg/user_v1"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateUserRequest) (desc.CreateUserResponse, error) {
	uuid, err := i.userService.Create(ctx, *model.UserInfo)
}
