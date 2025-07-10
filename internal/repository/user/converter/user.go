package converter

import (
	"clean-arhitecture/internal/model"
	repoModel "clean-arhitecture/internal/repository/user/model"
	"time"
)

func ToUserFromRepo(user *repoModel.User) *model.User {
	var updatedAt *time.Time
	if user.UpdatedAt.Valid {
		updatedAt = &user.UpdatedAt.Time
	}
	return &model.User{
		UUID:      user.UUID,
		Info:      ToUserInfoFromRepo(user.Info),
		CreatedAt: user.CreatedAt,
		UpdatedAt: updatedAt,
	}
}

func ToUserInfoFromRepo(info repoModel.UserInfo) model.UserInfo {
	return model.UserInfo{
		First_name: info.FirstName,
		Last_name:  info.LastName,
		Age:        info.Age,
	}
}

func ToUserInfoFromService(info *model.UserInfo) repoModel.UserInfo {
	return repoModel.UserInfo{
		FirstName: info.First_name,
		LastName:  info.Last_name,
		Age:       info.Age,
	}
}
