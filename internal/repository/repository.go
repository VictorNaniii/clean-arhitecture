package repository

import (
	"../model"
	"context"
)

type UserRepository interface {
	Create(ctx context.Context, userUUID string, info *model.UserInfo) error
	Get(ctx context.Context, uuid string) (*model.UserInfo, error)
}
