package repository

import (
	"context"
	"github.com/VictorNaniii/clean-arhitecture/internal/model"
)

type UserRepository interface {
	Create(ctx context.Context, userUUID string, info *model.UserInfo) error
	Get(ctx context.Context, uuid string) (*model.UserInfo, error)
}
