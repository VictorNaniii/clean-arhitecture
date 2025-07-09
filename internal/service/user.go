package service

import (
	"../model"
	"context"
)

type UserService interface {
	Create(ctx context.Context, info *model.UserInfo) (string error)
	Get(ctx context.Context, uuid string) (*model.UserInfo, error)
}
